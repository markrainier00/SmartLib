package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// 🚀 STEP 1: Hanapin ang buong DATABASE_URL mula sa .env
	dsn := os.Getenv("DATABASE_URL")

	// 🚀 STEP 2: Fallback (Kung walang DATABASE_URL, gagamitin niya yung luma mong setup)
	if dsn == "" {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")

		// Note: Sa Supabase, kailangan madalas ay isama ang pooler connection.
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, dbname, port)
	}

	// 🚀 STEP 3: Connect to Supabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("🚨 Failed to connect to database. Check your .env file! Error: ", err)
	}

	fmt.Println("✅ Database connected successfully!")
	DB = db
}
