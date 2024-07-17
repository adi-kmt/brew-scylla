package domain

type ProductPacksDTO struct {
	StoreName    string   `json:"store_name"`
	PackName     string   `json:"pack_name"`
	Price        float64  `json:"price"`
	Description  string   `json:"description"`
	ProductItems []string `json:"product_items"`
	Prizes10     float64  `json:"prizes_10"`
	Prizes5      float64  `json:"prizes_5"`
	Prizes3      float64  `json:"prizes_3"`
}
