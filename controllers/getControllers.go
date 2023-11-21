package controllers

import (
	"enchanted-castle-go/models"
	"net/http"
	"strings"

	"slices"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

var validSetCodes = []string{"TFC", "RFB"}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "Server Online",
	})
}

// RETURN ALL CARDS IN DATABASE
func GetAllCards(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results []models.Card
		err := supabase.DB.From("all_cards").Select("*").Execute(&results)
		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"length": len(results),
			"data":   results,
		})
	}
	return gin.HandlerFunc(fn)
}

// ADVANCED SEARCH FUNCTION FOR CARDS
func GetCardsByAdvanceSearch(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		sets, isSets := context.GetQueryArray("setCode")
		colors, isColors := context.GetQueryArray("color")
		inkable, isInkable := context.GetQueryArray("inkable")
		inkCost, isInkCost := context.GetQueryArray("inkCost")
		loreValue, isLoreValue := context.GetQueryArray("loreValue")
		rarity, isRarity := context.GetQueryArray("rarity")
		name, isName := context.GetQueryArray("name")
		// bodyText, isBodyText := context.GetQueryArray("bodyText")

		var results []models.Card

		allCards := supabase.DB.From("all_cards").Select("*")

		if isColors {
			allCards.In("color", colors)
		}
		if isSets {
			allCards.In("set_code", sets)
		}
		if isInkable {
			allCards.In("inkable", inkable)
		}
		if isInkCost {
			allCards.In("ink_cost", inkCost)
		}
		if isLoreValue {
			allCards.In("lore_value", loreValue)
		}
		if isRarity {
			allCards.In("rarity", rarity)
		}
		if isName {
			allCards.In("name", name)
		}

		err := allCards.Execute(&results)

		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{
			"length": len(results),
			"data":   results,
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
			err := supabase.DB.From("all_cards").Select("*").Eq("card_set_code", upperSet).Execute(&results)
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
func GetSingleCardInSet(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var result any

		set := context.Param("setCode")
		upperSet := strings.ToUpper(set)
		cardNumber := context.Param("cardNumber")

		err := supabase.DB.From("all_cards").Select("*").Single().Eq("card_set_code", upperSet).Eq("number", cardNumber).Execute(&result)

		if err != nil {
			panic(err)
		}

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
