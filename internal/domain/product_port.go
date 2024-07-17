package domain

import "github.com/adi-kmt/brew-scylla/internal/common/messages"

type ProductPort interface {
	GetAllStores(city string) ([]StoreDTO, *messages.AppError)
	GetProductCollections(storeName string, isFeatured bool) ([]ProductCollectionDTO, *messages.AppError)
	GetProductsByStore(storeName string) ([]ProductDTO, *messages.AppError)
	GetProductPacksByStore(storeName string) ([]ProductPacksDTO, *messages.AppError)
	GetProductsDetailsByStore(storeName, productName string) ([]ProductDetails, *messages.AppError)
	SearchProducts(productQuery, storeID, productCollection string) ([]ProductDTO, *messages.AppError)
}
