package repositories

import "github.com/eduardoshahn/golang-microservices/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
	Delete(category *entities.Category) error
}