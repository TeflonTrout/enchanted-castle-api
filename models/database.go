package models

type Json interface{}

type Database struct {
	Public struct {
		Tables struct {
			Cards struct {
				Row struct {
					Abilities  Json
					Artist     string
					BodyText   string
					CardNumber int
					Color      string
					FlavorText string
					ID         int
					ImageURLs  Json
					InkCost    string
					Inkable    bool
					LoreValue  string
					Name       string
					Rarity     string
					SeriesID   int
					Set        string
					SetCode    string
					SetID      int
					Strength   string
					Subtitle   string
					Subtypes   []string
					Traits     []string
					Type       []string
					Willpower  string
				}
				Insert struct {
					Abilities  Json
					Artist     string
					BodyText   string
					CardNumber int
					Color      string
					FlavorText string
					ID         int
					ImageURLs  Json
					InkCost    string
					Inkable    bool
					LoreValue  string
					Name       string
					Rarity     string
					SeriesID   int
					Set        string
					SetCode    string
					SetID      int
					Strength   string
					Subtitle   string
					Subtypes   []string
					Traits     []string
					Type       []string
					Willpower  string
				}
			}
		}
	}
}
