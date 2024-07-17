package user

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
)

type ProductService struct {
	productPort domain.ProductPort
}

func NewProductService(productPort domain.ProductPort) *ProductService {
	return &ProductService{
		productPort: productPort,
	}
}

func (s *ProductService) GetFeaturedProductCollections(storeName string) ([]domain.ProductCollectionDTO, *messages.AppError) {
	return s.productPort.GetProductCollections(storeName, true)
}

func (s *ProductService) GetAllProductCollections(storeName string) ([]domain.ProductCollectionDTO, *messages.AppError) {
	return s.productPort.GetProductCollections(storeName, false)
}

func (s *ProductService) GetProductsByStore(storeName string) ([]domain.ProductDTO, *messages.AppError) {
	return s.productPort.GetProductsByStore(storeName)
}

func (s *ProductService) GetProductPacks(storeName string) ([]domain.ProductPacksDTO, *messages.AppError) {
	return s.productPort.GetProductPacksByStore(storeName)
}

func (s *ProductService) SearchProducts(productQuery, storeName, productCollection string) ([]domain.ProductDTO, *messages.AppError) {
	return s.productPort.SearchProducts(productQuery, storeName, productCollection)
}

func (s *ProductService) GetAllStores(city string) ([]domain.StoreDTO, *messages.AppError) {
	return s.productPort.GetAllStores(city)
}
