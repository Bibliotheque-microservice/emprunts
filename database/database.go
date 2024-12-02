package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Bibliotheque-microservice/emprunts/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),        // Nom d'utilisateur
		os.Getenv("DB_PASSWORD"),    // Mot de passe
		os.Getenv("DB_NAME"),        // Nom de la base
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.Emprunt{}, &models.Penalite{})

	DB = Dbinstance{
		Db: db,
	}
}
