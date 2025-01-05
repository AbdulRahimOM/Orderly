package db

import (
	"fmt"
	"log"
	"orderly/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = GetConnectionToDB()

func init() {
	if resetDBFlagCalled() { //if resetdb flag is set or INIT_DB_EMPTY env is set to true, then clear the database
		if err := ClearDB(); err != nil {
			log.Fatal("Couldn't clear the database. Error:", err)
		}
	}

	InitDB() //initialize the public database (create tables and seed super admin)
}

func InitDB() { //is called in init function

	if config.Configs.Dev_AutoMigrateDbOnStart {
		migratePublicTables(DB)
	}
	initiateSuperAdmin(DB)
}

func GetConnectionToDB() *gorm.DB {
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Couldn't connect to the database. Error:", err)
	}

	return db
}

// ConnectToDB connects to the database and returns the connection.
// If the database does not exist, it creates it and returns the connection.
func connectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Configs.PostgresConn.DbHost,
		config.Configs.PostgresConn.DbUser,
		config.Configs.PostgresConn.DbPassword,
		config.Configs.PostgresConn.DbName,
		config.Configs.PostgresConn.DbPort,
		config.Configs.PostgresConn.DbSslMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to DB. Error:", err)
		return nil, err
	}
	return db, nil
}
