package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        uint 	    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Por bons hábitos, nunca será instaciado diretamente da Struct mas sim de uma função específica.
// NewCategory recebe o nome e retorna a categoria feita dentro dessa alocação da memória.
func NewCategory(name string) (*Category, error) {
	category := &Category{
		// ID o bd será responsável
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

//Método de validação para a Struct Category
func (c *Category) IsValid() error {
	if(len(c.Name) < 5){
		return fmt.Errorf("name must be greater then 5 characters, got %d", len(c.Name))
	}
	return nil
}