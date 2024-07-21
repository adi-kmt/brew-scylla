package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type ProductQueryPort interface {
	GetAllStores(city string) ([]entities.StoreEntity, *messages.AppError)
	GetProductCollections(storeName string, isFeatured bool) ([]entities.ProductCollectionEntity, *messages.AppError)
	GetProductsByStore(storeName string) ([]entities.ProductEntity, *messages.AppError)
	GetProductPacksByStore(storeName string) ([]entities.ProductPacksEntity, *messages.AppError)
	GetProductPackByStoreAndPackName(storeName, packName string) (*entities.ProductPacksEntity, *messages.AppError)
	GetProductsDetailsByStore(storeName, productName string) (*entities.ProductDetailsEntity, *messages.AppError)
}

type ProductSearchPort interface {
	SearchProducts(productQuery, storeName, productCollection string) ([]entities.ProductEntity, *messages.AppError)
}

type ProductPort interface {
	ProductQueryPort
	ProductSearchPort
}
