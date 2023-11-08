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
	Action      string `json:"action"`
	Artist      string `json:"artist"`
	Attack      int    `json:"attack"`
	CardSetCode string `json:"card_set_code"`
	CardSetID   int    `json:"card_set_id"`
	Color       string `json:"color"`
	ColorID     int    `json:"color_id"`
	Edition     []struct {
		Code string `json:"code"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"edition"`
	Flavor    string `json:"flavor"`
	Franchise struct {
		FranchiseCode string `json:"franchise_code"`
		FranchiseID   int    `json:"franchise_id"`
		FranchiseName string `json:"franchise_name"`
	} `json:"franchise"`
	ID        int    `json:"id"`
	Image     string `json:"image"`
	InkCost   int    `json:"ink_cost"`
	Inkable   bool   `json:"inkable"`
	Language  string `json:"language"`
	Languages struct {
		DE struct {
			Action   string `json:"action"`
			CardID   int    `json:"card_id"`
			Flavor   string `json:"flavor"`
			Image    string `json:"image"`
			Language string `json:"language"`
			Name     string `json:"name"`
			Title    string `json:"title"`
		} `json:"DE"`
		EN struct {
			Action   string `json:"action"`
			CardID   int    `json:"card_id"`
			Flavor   string `json:"flavor"`
			Image    string `json:"image"`
			Language string `json:"language"`
			Name     string `json:"name"`
			Title    string `json:"title"`
		} `json:"EN"`
		FR struct {
			Action   string `json:"action"`
			CardID   int    `json:"card_id"`
			Flavor   string `json:"flavor"`
			Image    string `json:"image"`
			Language string `json:"language"`
			Name     string `json:"name"`
			Title    string `json:"title"`
		} `json:"FR"`
	} `json:"languages"`
	Lore          int    `json:"lore"`
	Name          string `json:"name"`
	Number        int    `json:"number"`
	NumberInSet   string `json:"number_in_set"`
	Rarity        string `json:"rarity"`
	TextSeparator string `json:"text_separator"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Willpower     int    `json:"willpower"`
}
