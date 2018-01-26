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

type Product struct {
	Sku          string `json:"sku"`
	Name         string `json:"name"`
	TotalProduct uint   `json:"totalProduct"`
	IsRemoved	 bool	`json:"isRemoved"`
	gorm.Model
}
