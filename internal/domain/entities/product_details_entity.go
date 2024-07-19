package entities

type ProductDetailsEntity struct {
	StoreName    string             `json:"store_name"`
	ProductName  string             `json:"product_name"`
	ProductImage string             `json:"product_image"`
	Price        map[string]float64 `json:"price"`
	Category     string             `json:"category"`
	KCal         int                `json:"k_cal"`
	Description  string             `json:"description"`
	Ingredients  []string           `json:"ingredients"`
}
