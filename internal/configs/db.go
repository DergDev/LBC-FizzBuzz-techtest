package configs

import (
	"fmt"
	"log"
	"os"
	"techtest/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func init() {
	loadDBClient(true)
}

//Connection of the client to the DB
func loadDBClient(fillDB bool) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PORT"),
	)

	for {
		log.Println("Connecting to DB...")
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Println("Connection failed, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		if fillDB {
			for {
				err = db.AutoMigrate(&models.Statistics{})
				if err != nil {
					log.Println("Migration failed, retrying in 5 seconds...")
					time.Sleep(5 * time.Second)
					continue
				}
				DBClient = db
				break
			}
		}

		break
	}

	log.Println("Connexion to DB successful")
	return nil
}

//Publicly accessible client
func GetDBClient() *gorm.DB {
	db, err := DBClient.DB()
	if err != nil {
		return nil
	}
	if db.Ping() != nil {
		loadDBClient(false)
		return DBClient
	} else {
		return DBClient
	}
}
