package main

import (
	"enchanted-castle-go/controllers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

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
	// GET ROUTES
	router.GET("/all", controllers.GetAllCards(supabase))
	router.GET("/search", controllers.GetCardsByAdvanceSearch(supabase))
	router.GET("/cards/:setCode", controllers.GetCardsBySetCode(supabase))
	router.GET("/cards/:setCode/:cardNumber", controllers.GetSingleCardInSet(supabase))
	router.GET("/products", controllers.GetAllProducts(supabase))
	router.GET("/products/:setCode", controllers.GetProductsBySetCode(supabase))

	router.Run(":9090")
}
