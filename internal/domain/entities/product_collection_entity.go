package entities

type ProductCollectionEntity struct {
	StoreName      string `json:"store_name"`
	IsFeatured     bool   `json:"is_featured"`
	CollectionName string `json:"collection_name"`
}
