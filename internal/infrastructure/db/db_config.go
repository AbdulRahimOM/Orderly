package db

import (
	"fmt"
	"log"
	"orderly/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
