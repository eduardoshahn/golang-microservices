package main

import (
	"github.com/eduardoshahn/golang-microservices/cmd/api/controllers"
	"github.com/eduardoshahn/golang-microservices/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func (ctx *gin.Context)  {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func (ctx *gin.Context)  {
		controllers.ListCategories(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.DELETE("/", func(ctx *gin.Context) {
		controllers.DeleteCategory(ctx, inMemoryCategoryRepository)
	})
}