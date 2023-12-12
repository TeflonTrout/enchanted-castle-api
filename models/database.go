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
	CardUID     string   `json:"card_uid"`
	Name        string   `json:"name"`
	Subname     string   `json:"subname"`
	InkCost     int      `json:"ink_cost"`
	Inkable     bool     `json:"inkable"`
	Attack      int      `json:"attack"`
	Willpower   int      `json:"willpower"`
	ColorID     int      `json:"color_id"`
	Color       string   `json:"color"`
	Type        string   `json:"type"`
	Abilities   []string `json:"abilities"`
	BodyText    []string `json:"body_text"`
	Flavor      string   `json:"flavor"`
	Lore        string   `json:"lore"`
	Artist      string   `json:"artist"`
	SetID       int      `json:"set_id"`
	SetCode     string   `json:"set_code"`
	Number      int      `json:"number"`
	NumberInSet string   `json:"number_in_set"`
	Rarity      string   `json:"rarity"`
	Image       string   `json:"image"`
	Subtypes    []string `json:"subtypes"`
	Franchise   struct {
		FranchiseID   interface{} `json:"franchise_id"`
		FranchiseCode string      `json:"franchise_code"`
		FranchiseName string      `json:"franchise_name"`
	} `json:"franchise"`
	TextSeparator string `json:"text_separator"`
	ID            int    `json:"id"`
}
