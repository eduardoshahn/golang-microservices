package usecases

import (
	"github.com/eduardoshahn/golang-microservices/internal/entities"
	"github.com/eduardoshahn/golang-microservices/internal/repositories"
)

// Estrutura para a use case de deletar uma categoria
type deleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

// Função para criar e retornar uma nova instância da use case de delete
func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{repository}
}

// Método para executar a exclusão da categoria
func (u *deleteCategoryUseCase) Execute(name string) error {
	// Tenta excluir a categoria do repositório
	category := &entities.Category{Name: name}

	err := u.repository.Delete(category)
	if err != nil {
		// Se ocorrer um erro, como "category not found", retornamos o erro
		return err
	}

	return nil
}