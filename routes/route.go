package routes

import (
	"Book-GOAPI/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	books := r.Group("/books")
	{
		books.GET("", controller.GetBooks)
		books.GET("/:id", controller.GetBook)
		books.POST("", controller.CreateBook)
		books.PUT("/:id", controller.UpdateBook)
		books.PATCH("/:id", controller.PatchBook)
		books.DELETE("/:id", controller.DeleteBook)
	}

	return r
}
