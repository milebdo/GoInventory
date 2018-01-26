package controllers

import (
	"../models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

type ProductController struct {
	DB *gorm.DB
}

func (p ProductController) Create(c *gin.Context) {
	var product models.Product

	c.BindJSON(&product)
	
	p.DB.Create(&product)
	c.JSON(200, product)
}
