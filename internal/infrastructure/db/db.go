package db

import (
	"fmt"
	"log"
	"orderly/internal/domain/models"
	"orderly/internal/infrastructure/config"

	"gorm.io/gorm"
)

type MigrateTable interface {
	TableName() string
}

type postTableCreation interface {
	PostTableCreation(tx *gorm.DB) error
}

var (
	superAdminUsername = config.InitialData.SuperAdminUsername
	superAdminPassword = config.InitialData.SuperAdminPassword
	publicTables       = []MigrateTable{
		&models.SuperAdmin{},
	}
)

var PublicDB *gorm.DB = GetConnectionToPublicDB()

func init() {
	InitPublicDB() //initialize the public database (create tables and seed super admin)
}

func GetConnectionToPublicDB() *gorm.DB {
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Couldn't connect to the database. Error:", err)
	}

	return db
}

func InitPublicDB() { //is called in init function

	if config.Configs.Dev_AutoMigrateDbOnStart {
		fmt.Println("Auto migrating public tables...")
		migratePublicTables(PublicDB)
	}
	initiateSuperAdmin(PublicDB)
}

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
