package database

import (
	"log"
	"os"

	"github.com/joshuaikechukwuodoh/fiberApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	// Neon PostgreSQL DSN
	dsn := "host=ep-twilight-frost-ahq4qcka-pooler.c-3.us-east-1.aws.neon.tech user=neondb_owner password=npg_xrWvVA7QX8Hj dbname=neondb port=5432 sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{
		Db: db,
	}
}
