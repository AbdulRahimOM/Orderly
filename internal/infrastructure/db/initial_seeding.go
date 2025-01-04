package db

import (
	"fmt"
	"log"
	"orderly/internal/domain/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func initiateSuperAdmin(db *gorm.DB) {
	//check if super admin exists
	var superAdmin models.SuperAdmin
	err := db.First(&superAdmin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Super admin doesn't exist. Creating super admin...")

			//check if super admin username and password are provided
			if superAdminUsername == "" || superAdminPassword == "" {
				log.Fatal("Super admin username and password are required (for initial setup). But, either of them or both are not provided in environment.")
			}
			encryptedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(superAdminPassword), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Couldn't encrypt super admin password. Error:", err)
			}

			//create super admin
			superAdmin = models.SuperAdmin{
				Username:       superAdminUsername,
				HashedPassword: string(encryptedPasswordByte),
			}
			err = db.Create(&superAdmin).Error
			if err != nil {
				log.Fatal("Couldn't create super admin. Error:", err)
			}
			fmt.Println("Super admin created successfully")
		} else {
			log.Fatal("Couldn't get super admin. Error:", err)
		}
	} else {
		fmt.Println("Super admin already exists")
	}
}
