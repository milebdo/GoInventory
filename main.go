package main

import (
	"./controllers"
	"./routes"
	"./models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "./goinventory.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Product{}, $models.IncomingProduct{})

	product := controllers.NewProductController(db)
	r := gin.Default()
	routes.ProductRoutes(r, product)

	r.Run(":8000")
}
