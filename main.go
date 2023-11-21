package main

import (
	"enchanted-castle-go/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func initCronJob() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(12).Minutes().Do(func() {
		res, err := http.Get("https://enchanted-castle-server.onrender.com/health")
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %s\n", res.Status)
	})

	s.StartBlocking()
}

func main() {
	initCronJob()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	router := gin.Default()
	router.Use(cors.Default())
	port := os.Getenv("PORT")

	// ROUTES
	// HEALTH CHECK
	router.GET("/health", controllers.HealthCheck)
	// GET ROUTES
	router.GET("/all", controllers.GetAllCards(supabase))
	router.GET("/search", controllers.GetCardsByAdvanceSearch(supabase))
	router.GET("/cards/:setCode", controllers.GetCardsBySetCode(supabase))
	router.GET("/cards/:setCode/:cardNumber", controllers.GetSingleCardInSet(supabase))
	router.GET("/products", controllers.GetAllProducts(supabase))
	router.GET("/products/:setCode", controllers.GetProductsBySetCode(supabase))

	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
