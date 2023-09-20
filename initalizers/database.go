package initalizers

import (
	"fmt"
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectToSupabase() {
	password := os.Getenv("SUPABASE_PASSWORD")
	dsn := fmt.Sprintf("user=postgres password=%s host=db.dleeovrjibapyjgzvojl.supabase.co port=5432 dbname=postgres", password)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
}
