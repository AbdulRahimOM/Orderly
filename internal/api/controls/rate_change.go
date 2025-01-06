package controls

import (
	"fmt"
	"math"
	"orderly/internal/infrastructure/db"

	"github.com/gofiber/fiber/v2/log"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type ProductDemand struct {
	ID                 int     `json:"id" gorm:"column:id"`
	CurrentStock       int     `json:"current_stock" gorm:"column:current_stock"`
	OptimalStock       int     `json:"optimal_stock" gorm:"column:optimal_stock"`
	CurrentSalePrice   float64 `json:"current_sale_price" gorm:"column:current_sale_price"`
	MaxSalePrice       float64 `json:"max_sale_price" gorm:"column:max_sale_price"`
	MinSalePrice       float64 `json:"min_sale_price" gorm:"column:min_sale_price"`
	BasePrice          float64 `json:"base_price" gorm:"column:base_price"`
	LastDayOrdersCount int     `json:"last_day_orders_count" gorm:"column:last_day_orders_count"`
	NewPrice           float64 `json:"new_price" gorm:"-"`
}

func ScheduleMidNightRateChangeOperation() {
	c := cron.New()

	_, err := c.AddFunc("34 20 * * *", rateChangeOperation)
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	fmt.Println("Rate Change scheduled")
	c.Start()
}

func rateChangeOperation() {
	var err error

	fmt.Println("Rate Change started")
	defer func() {
		if err != nil {
			fmt.Println("Error happened in Rate Change: ", err)
			fmt.Println("Terminating Rate Change... Rate Change ended")
		} else {
			fmt.Println("Rate Change completed successfully")
		}

	}()
	db := db.DB
	productsInfo := []ProductDemand{}
	err = db.Raw(`
	SELECT
	    p.id, 
	    p.current_stock, 
	    p.optimal_stock,
	    p.current_sale_price, 
	    p.max_sale_price, 
	    p.min_sale_price, 
	    p.base_price,
	    COALESCE(SUM(op.quantity), 0) AS last_day_orders_count
	FROM 
	    products p
	LEFT JOIN 
	    order_products op ON op.product_id = p.id
	LEFT JOIN 
	    orders o ON o.id = op.order_id AND o.order_time >= NOW() - INTERVAL '1 day'
	WHERE 
		p.deleted_at IS NULL
	GROUP BY 
	    p.id
	`).Scan(&productsInfo).Error
	if err != nil {
		log.Errorf(`Error in fetching products: %v`, err)
		return
	}

	const (

		//supply based
		inflationFor0_5xStock          = 10
		maxAllowedSupplyBasedInflation = 15
		deflationfactorFor2xStock      = 10
		minAllowedSupplyBasedDeflation = 15

		//demandBased
		inflationFactorFor2xSalesWrtExpected  = 15
		maxAllowedDemandBasedInflation        = 25
		deflationfactorFor0_5SalesWrtExpected = 10
		minAllowedDemandBasedDeflation        = 15

		//Net Limits
		netAllowedInfation  = 30
		netAllowedDeflation = 25
	)

	for i, product := range productsInfo {
		stockRatio := float64(product.CurrentStock) / float64(product.OptimalStock)
		supplyBasedChange := 100 + (math.Log2(1/stockRatio))*10 //percentage
		if supplyBasedChange >= 100 {
			supplyBasedChange = min(supplyBasedChange, maxAllowedSupplyBasedInflation)
		} else {
			supplyBasedChange = max(supplyBasedChange, minAllowedSupplyBasedDeflation)
		}

		salesRatio := float64(product.LastDayOrdersCount) * 7 / float64(product.OptimalStock)
		demandBasedChange := 100 + (math.Log2(salesRatio))*10 //percentage
		if demandBasedChange >= 100 {
			demandBasedChange = min(demandBasedChange, maxAllowedDemandBasedInflation)
		} else {
			demandBasedChange = max(demandBasedChange, minAllowedDemandBasedDeflation)
		}

		netChange := (supplyBasedChange + demandBasedChange) / 2
		if netChange >= 100 {
			netChange = min(netChange, netAllowedInfation)
		} else {
			netChange = max(netChange, netAllowedDeflation)
		}

		newPrice := product.CurrentSalePrice * (netChange / 100)
		if newPrice > product.MaxSalePrice {
			newPrice = product.MaxSalePrice
		} else if newPrice < product.MinSalePrice {
			newPrice = product.MinSalePrice
		}

		productsInfo[i].NewPrice = newPrice
	}

	tx := db.Begin()
	err = tx.Error
	if err != nil {
		log.Errorf("failed to start transaction: %w", tx.Error)
		return
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback().Error; err != nil && err != gorm.ErrInvalidTransaction {
				log.Error("rollback failed: %v\n", err)
			}
		}
	}()

	//update prices
	for _, product := range productsInfo {
		err := tx.Exec(`
		UPDATE products
		SET current_sale_price = ?
		WHERE id = ?
		`, product.NewPrice, product.ID).Error
		if err != nil {
			log.Errorf(`Error in updating price for product: %v`, err)
			return
		}
	}

	err = tx.Commit().Error
	if err != nil {
		log.Errorf("error committing transaction: %v", err)
	}

}
