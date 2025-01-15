package controllers

import (
	"net/http"

	"github.com/eduardoshahn/golang-microservices/internal/repositories"
	use_cases "github.com/eduardoshahn/golang-microservices/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type deleteCategoryInput struct{
	Name string `json:"name" binding:"required"`
}

func DeleteCategory(ctx *gin.Context, repository repositories.ICategoryRepository){
	var body deleteCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		return 
	}

	//USE_CASE
	// Cria uma nova instância da use case de delete
	useCase := use_cases.NewDeleteCategoryUseCase(repository)

	// Executa a exclusão da categoria
	err := useCase.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Retorna sucesso se a exclusão for bem-sucedida
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}