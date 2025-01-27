package database

import (
	"admin_shop/pkg/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	var database *gorm.DB
	var err error
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "docker", "password", "127.0.0.1", "5435", "go_app_dev")

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

		if err == nil {
			break
		} else {
			log.Printf("Attemp %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}

	if database != nil {
		err = database.AutoMigrate(&models.User{})
		if err != nil {
			log.Printf("user is not migrate: %s", err)
		}
		err = database.AutoMigrate(&models.Shop{})
		if err != nil {
			log.Printf("shop is not migrate: %s", err)
		}
		err = database.AutoMigrate(&models.Product{})
		if err != nil {
			log.Printf("product is not migrate:%s", err)
		}

		err = database.AutoMigrate(&models.Order{})
		if err != nil {
			log.Printf("product is not migrate:%s", err)
		}
		err = database.AutoMigrate(&models.OrderItem{})
		if err != nil {
			log.Printf("product is not migrate:%s", err)
		}
	}

	// person1 := models.Person{
	// 	Name:  "Igor",
	// 	Age:   30,
	// 	Adres: "Chisinau",
	// 	Prof:  "Sofer",
	// }

	// database.Create(person1)

	return database
}
