package domain

import "github.com/adi-kmt/brew-scylla/internal/common/messages"

type ProductQueryPort interface {
	GetAllStores(city string) ([]StoreDTO, *messages.AppError)
	GetProductCollections(storeName string, isFeatured bool) ([]ProductCollectionDTO, *messages.AppError)
	GetProductsByStore(storeName string) ([]ProductDTO, *messages.AppError)
	GetProductPacksByStore(storeName string) ([]ProductPacksDTO, *messages.AppError)
	GetProductsDetailsByStore(storeName, productName string) ([]ProductDetails, *messages.AppError)
}

type ProductSearchPort interface {
	SearchProducts(productQuery, storeName, productCollection string) ([]ProductDTO, *messages.AppError)
}

type ProductPort interface {
	ProductQueryPort
	ProductSearchPort
}
