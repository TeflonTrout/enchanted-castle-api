package main

import (
	"enchanted-castle-go/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
)

func pingHealthEndpoint() {
	fmt.Println("Pinging!")
	url := "https://enchanted-castle-server.onrender.com/health"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Ping to %s - Status: %s\n", url, resp.Status)
}

func runCronJobs() {
	fmt.Println("John Doe")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	password := os.Getenv("SUPABASE_PASSWORD")
	host := os.Getenv("SUPABASE_HOST")
	user := os.Getenv("SUPABASE_USER")
	dbName := os.Getenv("SUPABASE_NAME")
	dbPort := os.Getenv("SUPABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbName, password, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	router := gin.Default()
	router.Use(cors.Default())
	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	// CREATE GO ROUTINE TO PING /HEALTH ENDPOINT EVERY 10 MINUTES
	go func() {
		c := cron.New()
		c.AddFunc("@every 10m", func() {
			url := "https://enchanted-castle-server.onrender.com/health"
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Ping to %s - Status: %s\n", url, resp.Status)
		})
		c.Start()
	}()

	// ROUTES
	// HEALTH CHECK
	router.GET("/health", controllers.HealthCheck(db))
	// GET ROUTES
	router.GET("/all", controllers.GetAllCards(supabase, db))
	router.GET("/search", controllers.GetCardsByAdvanceSearch(supabase, db))
	router.GET("/cards/:setCode", controllers.GetCardsBySetCode(supabase))
	router.GET("/cards/:setCode/:cardNumber", controllers.GetSingleCardInSet(supabase, db))
	router.GET("/products", controllers.GetAllProducts(supabase))
	router.GET("/products/:setCode", controllers.GetProductsBySetCode(supabase))

	serverErr := router.Run(":" + port)
	if serverErr != nil {
		log.Panicf("error: %s", err)
	}

}
