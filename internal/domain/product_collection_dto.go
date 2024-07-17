package domain

type ProductCollectionDTO struct {
	StoreName      string `json:"store_name"`
	IsFeatured     bool   `json:"is_featured"`
	CollectionName string `json:"collection_name"`
}
