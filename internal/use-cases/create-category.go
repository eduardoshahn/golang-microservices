package usecases

import (
	"github.com/eduardoshahn/golang-microservices/internal/entities"
	"github.com/eduardoshahn/golang-microservices/internal/repositories"
)

//Irá receber uma instacia do banco de dados, e não pode ser acessivel externamente
type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

//Função usada para retornar a locação do createCategoryUseCase, para não ser acessível a ele diretamente
func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

//Método para ser executado a use-case
func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}