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

var validSetCodes = []string{"P1", "TFC", "RFB", "ITI"}

func HealthCheck(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var results []models.Card
		db.Model(&models.Card{}).Table("all_cards").Find(&results)

		c.JSON(http.StatusOK, gin.H{
			"health": "Server Online",
		})
	}
	return gin.HandlerFunc(fn)
}

// RETURN ALL CARDS IN DATABASE
func GetAllCards(db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []models.Card
		db.Model(&models.Card{}).Table("all_cards").Find(&results)

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
func GetCardsByAdvanceSearch(db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		sets, isSets := context.GetQueryArray("setCode")
		colors, isColors := context.GetQuery("color")
		inkable, isInkable := context.GetQueryArray("inkable")
		inkCost, isInkCost := context.GetQueryArray("inkCost")
		loreValue, isLoreValue := context.GetQueryArray("loreValue")
		rarity, isRarity := context.GetQuery("rarity")
		name, isName := context.GetQuery("name")
		franchiseCode, isFranchiseCode := context.GetQueryArray("franchiseCode")
		bodyText, isBodyText := context.GetQuery("bodyText")

		var results []models.Card
		queryDB := db.Model(&models.Card{}).Table("all_cards")

		fmt.Println(isRarity)
		if isColors {
			// SPLIT COLORS STRING INTO ARRAY OF VALUES
			colorsArray := strings.Split(colors, ",")
			queryDB.Where("color IN ?", colorsArray)
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
			queryDB.Where(fmt.Sprintf("rarity::text ILIKE '%%%s%%'", rarity))
		}
		if isName {
			queryDB.Where(fmt.Sprintf("name ILIKE '%%%s%%' OR subname ILIKE '%%%s%%'", name, name))
		}
		if isFranchiseCode {
			queryDB.Where("franchise->>'franchise_code' IN ?", franchiseCode)
		}
		if isBodyText {
			queryDB.Where(fmt.Sprintf("EXISTS (SELECT 1 FROM jsonb_array_elements_text(body_text) AS elem WHERE elem ILIKE '%%%s%%');", bodyText))
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
func GetCardsBySetCode(db *gorm.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []models.Card
		var setResults models.SetData

		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)

		err := db.Model(&models.Card{}).Table("all_cards").Where("set_code = ?", upperSet).Scan(&results)
		setErr := db.Model(&models.SetData{}).Table("card_sets").Where("set_code = ?", upperSet).Scan(&setResults)

		if err != nil {
			fmt.Println("ERR HERE")
		}
		if setErr != nil {
			fmt.Println("ERR HERE2")
		}

		// CHECK IF SET CODE IS A VALID SET CODE
		if slices.Contains(validSetCodes, upperSet) {
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

			context.JSON(http.StatusOK, gin.H{
				"length":  len(results),
				"data":    results,
				"setData": setResults,
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

		fmt.Println(context.Params)
		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)
		cardNumber := context.Param("cardNumber")

		db.Model(&models.Card{}).Table("all_cards").Where("set_code = ?", upperSet).Where("number = ?", cardNumber).First(&result)

		fmt.Println(result)
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
