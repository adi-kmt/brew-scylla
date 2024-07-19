package entities

import "github.com/adi-kmt/brew-scylla/internal/common/messages"

type Category string

type ProductEntity struct {
	StoreName      string             `json:"store_name"`
	ProductName    string             `json:"product_name"`
	ProductImage   string             `json:"product_image"`
	Price          map[string]float64 `json:"price"`
	Category       Category           `json:"category"`
	CollectionName string             `json:"collection_name"`
}

const (
	VEGETARIAN    Category = "VEGETARIAN"
	VEGAN         Category = "VEGAN"
	NonVegeterian Category = "NONVEGETERIAN"
)

func (category Category) String() (string, error) {
	switch category {
	case VEGETARIAN:
		return "VEGETARIAN", nil
	case VEGAN:
		return "VEGAN", nil
	case NonVegeterian:
		return "NONVEGETERIAN", nil
	}
	return "", messages.BadRequest("Invalid category")
}

func GetCategory(category string) (Category, error) {
	switch category {
	case "VEGETARIAN":
		return VEGETARIAN, nil
	case "VEGAN":
		return VEGAN, nil
	case "NONVEGETERIAN":
		return NonVegeterian, nil
	}
	return "", messages.BadRequest("Invalid category")
}
