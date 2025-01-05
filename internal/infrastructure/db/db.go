package db

import (
	"fmt"
	"log"
	"orderly/internal/domain/models"

	"gorm.io/gorm"
)

type MigrateTable interface {
	TableName() string
}

type postTableCreation interface {
	PostTableCreation(tx *gorm.DB) error
}

var (
	publicTables = []MigrateTable{
		&models.SuperAdmin{},
		&models.Admin{},
		&models.User{},
		&models.Address{},
		&models.AdminPrivilege{},

		&models.CartItem{},
		&models.IncomingTransaction{},
		&models.Order{},
		&models.OrderProduct{},
		&models.Product{},
		&models.ProductCategory{},
		&models.ProductRating{},
		&models.ProductReview{},
		&models.Return{},
		&models.RefundTransaction{},
	}
)

func migratePublicTables(db *gorm.DB) {
	for _, table := range publicTables {
		err := db.AutoMigrate(table)
		if err != nil {
			log.Fatalf("error migrating %s table: %v", table.TableName(), err)
		}

		if implementedTable, ok := table.(postTableCreation); ok {
			err = implementedTable.PostTableCreation(db)
			if err != nil {
				log.Fatalf("error running post-table-creation for %s: %v", table.TableName(), err)
			}
		}
	}

	fmt.Println("Migrated tables successfully")
}
