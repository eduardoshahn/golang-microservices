package repositories

import "github.com/eduardoshahn/golang-microservices/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)

	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}

// Remove um elemento do Slice (Sequência de elementos do tipo []*entities.Category)
func removeByIndex(slice []*entities.Category, index int) []*entities.Category {
    if index < 0 || index >= len(slice) {
        return slice // Retorna o slice original se o índice for inválido
    }
	// Aqui ele usa a função append para concatenar dois slices, um com os elementos antes do excluído e outro depois
    return append(slice[:index], slice[index+1:]...)
}

func (r *inMemoryCategoryRepository) Delete(category *entities.Category) error {
	for i := len(r.db) - 1; i >= 0; i-- {
		if r.db[i].Name == category.Name{
			r.db = removeByIndex(r.db, i)
			return nil
		}
	}
	return nil
}

