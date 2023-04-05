package cmd

import (
	handler "auth_service/internal/handler/http"
	"auth_service/internal/middleware"

	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.New()

	fileAdapter := fileadapter.NewAdapter("config/basic_policy.csv")

	authorized := r.Group("/")
	authorized.Use(gin.Logger())
	authorized.Use(gin.Recovery())
	{
		authorized.POST("/api/book/create", middleware.Authorize("book", "write", fileAdapter), handler.CreateBook)
		authorized.GET("/api/book/:id", middleware.Authorize("book", "read", fileAdapter), handler.GetBook)
	}

	return r
}
