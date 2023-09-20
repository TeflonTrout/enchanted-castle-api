package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func GetAllCards(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var results any
		err := supabase.DB.From("cards").Select("*").Execute(&results)
		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"Results": results,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetAllCardsBySetCode(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		setCode := context.Param("setCode")

		var results any
		err := supabase.DB.From("cards").Select("*").Eq("set_code", setCode).Execute(&results)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"Results": results,
		})
	}
	return gin.HandlerFunc(fn)
}

func GetAllCardsByColor(supabase *supa.Client) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		array := context.QueryArray("colors")

		var results any
		err := supabase.DB.From("cards").Select("*").In("color", array).Execute(&results)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"Query Array": array,
			"Results":     results,
		})
	}
	return gin.HandlerFunc(fn)
}
