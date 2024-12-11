package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "go-final-project/models"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
    if DB == nil {
        // Load environment variables
        err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }

        // Get environment variables
        dbUser := os.Getenv("DB_USER")
        dbPassword := os.Getenv("DB_PASSWORD")
        dbHost := os.Getenv("DB_HOST")
        dbPort := os.Getenv("DB_PORT")
        dbName := os.Getenv("DB_NAME")

        // Create the DSN
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            dbUser, dbPassword, dbHost, dbPort, dbName)

        // Connect to the database
        DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatalf("Error connecting to the database: %v", err)
        }

        fmt.Println("Connected to the database successfully")

        // Drop all tables before running migrations
        err = DB.Migrator().DropTable(&models.Barista{}, &models.Category{}, &models.Product{}, &models.Review{})
        if err != nil {
            log.Fatalf("Error dropping tables: %v", err)
        }
        fmt.Println("Dropped existing tables")

        // Run migrations
        err = DB.AutoMigrate(
            &models.Barista{},
            &models.Category{},
            &models.Product{},
            &models.Review{},
        )
        if err != nil {
            log.Fatalf("Error during migration: %v", err)
        }
        fmt.Println("Database migration completed")
    }

    return DB
}