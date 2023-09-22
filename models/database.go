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

type Card struct {
	Abilities struct {
	} `json:"abilities"`
	Artist     string `json:"artist"`
	BodyText   string `json:"body_text"`
	CardNumber int    `json:"card_number"`
	Color      string `json:"color"`
	FlavorText string `json:"flavor_text"`
	ID         int    `json:"id"`
	ImageUrls  struct {
		ArtCrop string `json:"art_crop"`
		Foil    string `json:"foil"`
		Large   string `json:"large"`
		Medium  string `json:"medium"`
		NoArt   string `json:"no_art"`
		Small   string `json:"small"`
	} `json:"image_urls"`
	InkCost   string `json:"ink_cost"`
	Inkable   bool   `json:"inkable"`
	LoreValue string `json:"lore_value"`
	Name      string `json:"name"`
	Rarity    string `json:"rarity"`
	SeriesID  any    `json:"series_id"`
	Set       string `json:"set"`
	SetCode   string `json:"set_code"`
	SetID     any    `json:"set_id"`
	Strength  string `json:"strength"`
	Subtitle  string `json:"subtitle"`
	Subtypes  any    `json:"subtypes"`
	Traits    any    `json:"traits"`
	Type      any    `json:"type"`
	Willpower string `json:"willpower"`
}
