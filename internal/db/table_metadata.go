package db

import (
	"github.com/scylladb/gocqlx/v3/table"
)

var GetUserDetailsByIDTable = table.New(GetUserDetailsByIDMetadata)

var GetAllStoreByCityTable = table.New(GetAllStoreByCityMetadata)

var GetCuponCodeByStoreIdTable = table.New(GetCuponCodeByStoreIdMetadata)

var GetCollectionsByStoreIdTable = table.New(GetCollectionsByStoreIdMetadata)

var GetProductsPacksByStoreIdTable = table.New(GetProductsPacksByStoreIdMetadata)

var GetProductByStoreIDTable = table.New(GetProductsByStoreIdMetadata)

var GetProductsDetailsByProductNameTable = table.New(GetProductsDetailsByProductNameMetadata)

var GetOrdersByUserIDTable = table.New(GetOrdersByUserIDMetadata)

var GetOrderDetailsByIDTable = table.New(GetOrderDetailsByIDMetadata)

var PackRedemptionByUseridTable = table.New(PackRedemptionByUseridMetadata)

// Define table metadata for GetUserDetailsByID
var GetUserDetailsByIDMetadata = table.Metadata{
	Name:    "get_user_details_by_id",
	Columns: []string{"username", "phone_no", "coins", "coupons_used"},
	PartKey: []string{"username"},
	SortKey: []string{},
}

// Define table metadata for GetAllStoreByCity
var GetAllStoreByCityMetadata = table.Metadata{
	Name:    "get_all_store_by_city",
	Columns: []string{"city", "store_name", "location", "closes_at"},
	PartKey: []string{"city"},
	SortKey: []string{"store_name"},
}

// Define table metadata for CuponCodeByStoreId
var GetCuponCodeByStoreIdMetadata = table.Metadata{
	Name:    "cupon_code_by_store_id",
	Columns: []string{"store_name", "cupon_code", "expiry_date", "cupon_image"},
	PartKey: []string{"store_name"},
	SortKey: []string{"expiry_date", "cupon_code"},
}

// Define table metadata for GetCollectionsByStoreId
var GetCollectionsByStoreIdMetadata = table.Metadata{
	Name:    "get_collections_by_store_id",
	Columns: []string{"store_name", "is_featured", "collection_name"},
	PartKey: []string{"store_name"},
	SortKey: []string{"is_featured"},
}

// Define table metadata for ProductsPacksByStoreId
var GetProductsPacksByStoreIdMetadata = table.Metadata{
	Name:    "products_packs_by_store_id",
	Columns: []string{"store_name", "pack_name", "price", "description", "product_items", "prizes_10", "prizes_5", "prizes_3"},
	PartKey: []string{"store_name"},
	SortKey: []string{"pack_name"},
}

// Define table metadata for ProductByStoreID
var GetProductsByStoreIdMetadata = table.Metadata{
	Name:    "product_by_store_id",
	Columns: []string{"store_name", "product_name", "product_image", "price", "category", "collection_name"},
	PartKey: []string{"store_name"},
	SortKey: []string{"product_name"},
}

// Define table metadata for GetProductsDetailsByProductName
var GetProductsDetailsByProductNameMetadata = table.Metadata{
	Name:    "get_products_details_by_product_name",
	Columns: []string{"store_name", "product_name", "product_image", "price", "category", "k_cal", "description", "ingredients"},
	PartKey: []string{"store_name"},
	SortKey: []string{"product_name"},
}

// Define table metadata for GetOrdersByUserID
var GetOrdersByUserIDMetadata = table.Metadata{
	Name:    "get_orders_by_user_id",
	Columns: []string{"username", "order_id", "order_time", "order_status"},
	PartKey: []string{"username"},
	SortKey: []string{"order_time", "order_id"},
}

// Define table metadata for GetOrderDetailsByID
var GetOrderDetailsByIDMetadata = table.Metadata{
	Name:    "get_order_details_by_id",
	Columns: []string{"username", "order_id", "product_name", "product_quantity", "product_price", "order_status", "order_timestamp", "order_sub_total", "discount_percentage", "order_total", "pack_name", "is_pack"},
	PartKey: []string{"username", "order_id"},
	SortKey: []string{"product_name"},
}

// Define table metadata for PackRedemptionByUserid
var PackRedemptionByUseridMetadata = table.Metadata{
	Name:    "pack_redemption_by_userid",
	Columns: []string{"username", "store_name", "order_id", "pack_name", "order_items_remaining", "expiry_timestamp"},
	PartKey: []string{"username"},
	SortKey: []string{"store_name", "pack_name"},
}
