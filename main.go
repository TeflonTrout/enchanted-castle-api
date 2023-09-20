package main

import (
	"enchanted-castle-go/controllers"
	"enchanted-castle-go/initalizers"

	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func init() {
	initalizers.ConnectToSupabase()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	router := gin.Default()
	router.Use(cors.Default())

	// ROUTES
	router.GET("/all", controllers.GetAllCards(supabase))
	router.GET("/cards/:setCode", controllers.GetAllCardsBySetCode(supabase))
	router.GET("/cards", controllers.GetAllCardsByColor(supabase))
	router.Run(":9090")
}
