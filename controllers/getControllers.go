package controllers

import (
	"enchanted-castle-go/models"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"slices"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	"gorm.io/gorm"
)

var validSetCodes = []string{"TFC", "RFB"}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "Server Online",
	})
}

// RETURN ALL CARDS IN DATABASE
func GetAllCards(supabase *supa.Client, db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []models.Card
		// var cards []models.Card
		db.Model(&models.Card{}).Table("all_cards").Find(&results)

		// err := supabase.DB.From("all_cards").Select("*").Execute(&results)
		// if err != nil {
		// panic(err)
		// }

		// Paginate the data based on query parameters (e.g., page and itemsPerPage)
		page := 1
		itemsPerPage := 20

		// Extract page and itemsPerPage from query parameters if provided
		if sortParam := context.Request.URL.Query().Get("sort"); sortParam != "" {
			if sortParam == "alphabetical" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Name < results[j].Name
				})
			}
			if sortParam == "cardNumber" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Number < results[j].Number
				})
			}
			if sortParam == "attack" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Attack > results[j].Attack
				})
			}
			if sortParam == "willpower" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Willpower > results[j].Willpower
				})
			}
			if sortParam == "lore" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Lore > results[j].Lore
				})
			}
		}

		// Extract page and itemsPerPage from query parameters if provided
		if pageParam := context.Request.URL.Query().Get("page"); pageParam != "" {
			page, _ = strconv.Atoi(pageParam)
		}

		if itemsPerPageParam := context.Request.URL.Query().Get("limit"); itemsPerPageParam != "" {
			itemsPerPage, _ = strconv.Atoi(itemsPerPageParam)
		}

		// Calculate total pages
		totalPages := (len(results) + itemsPerPage - 1) / itemsPerPage

		// Paginate the data based on the requested page
		startIndex := (page - 1) * itemsPerPage
		endIndex := startIndex + itemsPerPage

		// Ensure endIndex is within bounds
		if endIndex > len(results) {
			endIndex = len(results)
		}

		// Extract the items for the current page
		paginatedItems := results[startIndex:endIndex]

		context.JSON(http.StatusOK, gin.H{
			"limit":      len(paginatedItems),
			"page":       page,
			"totalPages": totalPages,
			"data":       paginatedItems,
		})
	}
	return gin.HandlerFunc(fn)
}

// ADVANCED SEARCH FUNCTION FOR CARDS
func GetCardsByAdvanceSearch(supabase *supa.Client, db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		sets, isSets := context.GetQueryArray("setCode")
		colors, isColors := context.GetQueryArray("color")
		inkable, isInkable := context.GetQueryArray("inkable")
		inkCost, isInkCost := context.GetQueryArray("inkCost")
		loreValue, isLoreValue := context.GetQueryArray("loreValue")
		rarity, isRarity := context.GetQueryArray("rarity")
		name, isName := context.GetQuery("name")
		franchiseCode, isFranchiseCode := context.GetQueryArray("franchiseCode")
		// bodyText, isBodyText := context.GetQueryArray("bodyText")

		var results []models.Card
		queryDB := db.Model(&models.Card{}).Table("all_cards")

		if isColors {
			queryDB.Where("color IN ?", colors)
		}
		if isSets {
			queryDB.Where("set_code IN ?", sets)
		}
		if isInkable {
			queryDB.Where("inkable IN ?", inkable)
		}
		if isInkCost {
			queryDB.Where("ink_cost IN ?", inkCost)
		}
		if isLoreValue {
			queryDB.Where("lore_value IN ?", loreValue)
		}
		if isRarity {
			queryDB.Where("rarity IN ?", rarity)
		}
		if isName {
			query := fmt.Sprintf("SELECT * FROM all_cards WHERE name ILIKE '%%%s%%' OR subname ILIKE '%%%s%%';", name, name)
			queryDB.Raw(query)
		}
		if isFranchiseCode {
			queryDB.Where("franchise->>'franchise_code' IN ?", franchiseCode)
		}

		queryDB.Scan(&results)

		// Paginate the data based on query parameters (e.g., page and itemsPerPage)
		page := 1
		itemsPerPage := 20

		// Extract page and itemsPerPage from query parameters if provided
		if sortParam := context.Request.URL.Query().Get("sort"); sortParam != "" {
			if sortParam == "alphabetical" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Name < results[j].Name
				})
			}
			if sortParam == "cardNumber" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Number < results[j].Number
				})
			}
			if sortParam == "attack" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Attack > results[j].Attack
				})
			}
			if sortParam == "willpower" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Willpower > results[j].Willpower
				})
			}
			if sortParam == "lore" {
				sort.Slice(results, func(i, j int) bool {
					return results[i].Lore > results[j].Lore
				})
			}
		}

		// Extract page and itemsPerPage from query parameters if provided
		if pageParam := context.Request.URL.Query().Get("page"); pageParam != "" {
			page, _ = strconv.Atoi(pageParam)
		}

		if itemsPerPageParam := context.Request.URL.Query().Get("limit"); itemsPerPageParam != "" {
			itemsPerPage, _ = strconv.Atoi(itemsPerPageParam)
		}

		// Calculate total pages
		totalPages := (len(results) + itemsPerPage - 1) / itemsPerPage

		// Paginate the data based on the requested page
		startIndex := (page - 1) * itemsPerPage
		endIndex := startIndex + itemsPerPage

		// Ensure endIndex is within bounds
		if endIndex > len(results) {
			endIndex = len(results)
		}

		// Extract the items for the current page
		paginatedItems := results[startIndex:endIndex]

		context.JSON(http.StatusOK, gin.H{
			"limit":      len(paginatedItems),
			"page":       page,
			"totalPages": totalPages,
			"data":       paginatedItems,
		})
	}
	return gin.HandlerFunc(fn)
}

// RETURN ALL CARDS IN A SET
func GetCardsBySetCode(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []models.Card
		var setResults []any

		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)

		// CHECK IF SET CODE IS A VALID SET CODE
		if slices.Contains(validSetCodes, upperSet) {
			err := supabase.DB.From("all_cards").Select("*").Eq("set_code", upperSet).Execute(&results)
			setErr := supabase.DB.From("card_sets").Select("*").Eq("set_code", upperSet).Execute(&setResults)
			if err != nil {
				panic(err)
			}
			if setErr != nil {
				panic(err)
			}

			context.JSON(http.StatusOK, gin.H{
				"length":  len(results),
				"data":    results,
				"setData": setResults[0],
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide a valid set code.",
			})
		}
	}
	return gin.HandlerFunc(fn)
}

// RETURN SINGLE CARD
func GetSingleCardInSet(supabase *supa.Client, db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var result models.Card

		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)
		cardNumber := context.Param("cardNumber")

		db.Model(&models.Card{}).Table("all_cards").Where("set_code = ?", upperSet).Where("number = ?", cardNumber).First(&result)

		context.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
	return gin.HandlerFunc(fn)
}

// RETURN ALL CARD PRODUCTS
func GetAllProducts(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []any

		context.JSON(http.StatusOK, gin.H{
			"data":   results,
			"length": len(results),
		})
	}
	return gin.HandlerFunc(fn)
}

// SEARCH FOR PRODUCTS BY SET CODE
func GetProductsBySetCode(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []any
		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)

		if slices.Contains(validSetCodes, upperSet) {
			context.JSON(http.StatusOK, gin.H{
				"data":   results,
				"length": len(results),
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide a valid set code.",
			})
		}
	}
	return gin.HandlerFunc(fn)
}
