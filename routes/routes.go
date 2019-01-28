package routes

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, pc *controllers.ProductController) {
	// r.POST("/product", pc.Create)
	product := r.Group("/api/product")
	{
		product.POST("/", ip.CreateNewProduct)
		product.GET("/", ip.FetchAllProduct)
		product.GET("/:id", ip.FetchOneProduct)
		product.PUT("/:id", ip.UpdateProduct)
		product.DELETE("/:id", ip.DeleteProduct)
	}
}

func IncomingProductRoutes(r *gin.Engine, ip *controllers.IncomingProductController) {
	incoming := r.Group("/api/incoming")
	{
		incoming.POST("/", ip.Create)
		incoming.GET("/", ip.FetchAll)
		incoming.GET("/:id", ip.FetchOne)
		incoming.PUT("/:id", ip.Update)
		incoming.DELETE("/:id", ip.Delete)
	}
}

func OutProductRoutes(r *gin.Engine, op *controllers.OutProductController) {
	out := r.Group("/api/out")
	{
		out.POST("/", ip.CreateNewOutProduct)
		out.GET("/", ip.FetchAllOutProduct)
		out.GET("/:id", ip.FetchOneOutProduct)
		out.PUT("/:id", ip.UpdateOutProduct)
		out.DELETE("/:id", ip.DeleteOutProduct)
	}
}
