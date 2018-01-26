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

type OutProduct struct {
	TrxTime time.Time 	string 	`json:"dateTime"`
	ProductId          	uint	`json:"productId"`
	TotalOutProduct     uint 	`json:"totalOutProduct"`
	SellPrice 			uint	`json:"sellPrice"`
	TotalPrice 			uint	`json:"totalPrice"`
	Invoice 			string	`json:"invoice"`
	Notes 				uint   	`json:"notes"`
	IsRemoved	 		bool	`json:"isRemoved"`
	gorm.Model
}
