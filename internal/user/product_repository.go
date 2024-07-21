package user

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
)

type ProductRepository struct {
	session gocqlx.Session
}

func NewProductRepository(session gocqlx.Session) *ProductRepository {
	return &ProductRepository{
		session: session,
	}
}

func (repo *ProductRepository) GetAllStores(city string) ([]entities.StoreEntity, *messages.AppError) {
	var stores []entities.StoreEntity
	err := db.GetAllStoreByCityTable.SelectQuery(repo.session).Bind(city).Select(&stores)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get all stores present")
	}
	return stores, nil
}

func (repo *ProductRepository) GetProductCollections(storeName string, isFeatured bool) ([]entities.ProductCollectionEntity, *messages.AppError) {
	var productCollections []entities.ProductCollectionEntity
	err := repo.session.Query(
		qb.Select(db.GetCollectionsByStoreIdTable.Name()).
			Columns(db.GetCollectionsByStoreIdTable.Metadata().Columns...).
			Where(db.GetCollectionsByStoreIdTable.PrimaryKeyCmp()[:2]...).ToCql(),
	).Bind(storeName, isFeatured).SelectRelease(&productCollections)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get product collections")
	}
	return productCollections, nil
}

func (repo *ProductRepository) GetProductsByStore(storeName string) ([]entities.ProductEntity, *messages.AppError) {
	var products []entities.ProductEntity
	err := db.GetProductByStoreIDTable.SelectQuery(repo.session).Bind(storeName).Select(&products)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get products")
	}
	return products, nil
}

func (repo *ProductRepository) GetProductsDetailsByStore(storeName, productName string) (*entities.ProductDetailsEntity, *messages.AppError) {
	var productDetails entities.ProductDetailsEntity
	err := repo.session.Query(
		qb.Select(db.GetProductsDetailsByProductNameTable.Name()).
			Columns(db.GetProductsDetailsByProductNameTable.Metadata().Columns...).
			Where(db.GetProductsDetailsByProductNameTable.PrimaryKeyCmp()[:2]...).ToCql(),
	).Bind(storeName, productName).SelectRelease(&productDetails)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get product details")
	}
	return &productDetails, nil
}

func (repo *ProductRepository) GetProductPacksByStore(storeName string) ([]entities.ProductPacksEntity, *messages.AppError) {
	var productPacks []entities.ProductPacksEntity
	err := db.GetProductsPacksByStoreIdTable.SelectQuery(repo.session).Bind(storeName).Select(&productPacks)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get product packs")
	}
	return productPacks, nil
}

func (repo *ProductRepository) SearchProducts(productQuery, storeName, productCollection string) ([]entities.ProductEntity, *messages.AppError) {
	return nil, nil
}

func (repo *ProductRepository) GetProductPackByStoreAndPackName(storeName, packName string) (*entities.ProductPacksEntity, *messages.AppError) {
	return nil, nil
}
