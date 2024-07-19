package user

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type ProductService struct {
	productPort domain.ProductPort
}

func NewProductService(productPort domain.ProductPort) *ProductService {
	return &ProductService{
		productPort: productPort,
	}
}

func (s *ProductService) GetFeaturedProductCollections(storeName string) ([]entities.ProductCollectionEntity, *messages.AppError) {
	return s.productPort.GetProductCollections(storeName, true)
}

func (s *ProductService) GetAllProductCollections(storeName string) ([]entities.ProductCollectionEntity, *messages.AppError) {
	return s.productPort.GetProductCollections(storeName, false)
}

func (s *ProductService) GetProductsByStore(storeName string) ([]entities.ProductEntity, *messages.AppError) {
	return s.productPort.GetProductsByStore(storeName)
}

func (s *ProductService) GetProductDetailsByStore(storeName, productName string) ([]entities.ProductDetailsEntity, *messages.AppError) {
	return s.productPort.GetProductsDetailsByStore(storeName, productName)
}

func (s *ProductService) GetProductPacks(storeName string) ([]entities.ProductPacksEntity, *messages.AppError) {
	return s.productPort.GetProductPacksByStore(storeName)
}

func (s *ProductService) SearchProducts(productQuery, storeName, productCollection string) ([]entities.ProductEntity, *messages.AppError) {
	return s.productPort.SearchProducts(productQuery, storeName, productCollection)
}

func (s *ProductService) GetAllStores(city string) ([]entities.StoreEntity, *messages.AppError) {
	return s.productPort.GetAllStores(city)
}
