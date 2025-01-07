package main

import (
	"fmt"
	"orderly/internal/infrastructure/db"
)

func main() {
	fmt.Println("Migrating the database... Seeding the super admin if not exists...")
	db.GetConnectionToDB()
	db.InitDB() //initialize the public database (create tables and seed super admin)
}
