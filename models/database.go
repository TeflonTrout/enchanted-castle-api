package models

import "github.com/lib/pq"

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
	CardUID       string         `json:"uid" gorm:"type:text"`
	Name          string         `json:"name"`
	Subname       string         `json:"subname"`
	InkCost       int            `json:"ink_cost"`
	Inkable       bool           `json:"inkable"`
	Attack        int            `json:"attack"`
	Willpower     int            `json:"willpower"`
	ColorID       int            `json:"color_id"`
	Color         string         `json:"color"`
	Type          string         `json:"type"`
	Abilities     []string       `json:"abilities" gorm:"serializer:json"`
	BodyText      []string       `json:"body_text" gorm:"serializer:json"`
	Flavor        string         `json:"flavor"`
	Lore          int            `json:"lore"`
	Artist        string         `json:"artist"`
	SetID         int            `json:"set_id"`
	SetCode       string         `json:"set_code"`
	Number        int            `json:"number"`
	NumberInSet   string         `json:"number_in_set"`
	Rarity        string         `json:"rarity"`
	Image         string         `json:"image"`
	Subtypes      pq.StringArray `json:"subtypes" gorm:"type:text[]"`
	Franchise     Franchise      `json:"franchise" gorm:"serializer:json"`
	ID            int            `json:"id"`
	TextSeparator string         `json:"text_separator"`
}

type Franchise struct {
	FranchiseID   int    `json:"franchise_id"`
	FranchiseCode string `json:"franchise_code"`
	FranchiseName string `json:"franchise_name"`
}

type SetData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Card_count int    `json:"card_count"`
	Set_code   string `json:"set_code"`
}
