package usecases

import (
	"github.com/eduardoshahn/golang-microservices/internal/entities"
	"github.com/eduardoshahn/golang-microservices/internal/repositories"
)

type listCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoryUseCase(repository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}