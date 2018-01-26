package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type IncomingProduct struct {
	TrxTime time.Time 		string `json:"dateTime"`
	Sku          	  		string `json:"sku"`
	Name         	  		string `json:"name"`
	TotalProductOrder 		uint   `json:"totalProductOrder"`
	TotalProductReceived 	uint   `json:"totalProductReceived"`
	BuyPrice 				uint   `json:"buyPrice"`
	TotalPrice 				uint   `json:"totalPrice"`
	Invoice 				uint   `json:"invoice"`
	Notes 					uint   `json:"notes"`
	IsRemoved	 			bool   `json:"isRemoved"`
	gorm.Model
}
