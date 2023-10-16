package models

type Item struct {
	Id        string `gorm:"primaryKey;not null;size:255"`
	ItemName  string `gorm:"size:255"`
	ItemPrice  float64 
	Amount     int
}

type ItemRequst struct {
	ItemName string `json:"item_name"`
	ItemPrice float64 `json:"item_price"`
	Amount int `json:"amount"`
}