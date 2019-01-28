package controllers

import (
	"../models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewIncomingProductController(db *gorm.DB) *IncomingProductController {
	return &IncomingProductController{DB: db}
}

type IncomingProductController struct {
	DB *gorm.DB
}

func (p IncomingProductController) Create(c *gin.Context) {
	var product models.IncomingProduct

	c.BindJSON(&product)
	
	p.DB.Create(&product)
	c.JSON(200, product)
}
