package db

import (
	"github.com/scylladb/gocqlx/v3/table"
)

var GetUserDetailsByIDTable = table.New(getUserDetailsByIDMetadata)

var GetAllStoreByCityTable = table.New(getAllStoreByCityMetadata)

var GetCuponCodeByStoreIdTable = table.New(getCuponCodeByStoreIdMetadata)

var GetCollectionsByStoreIdTable = table.New(getCollectionsByStoreIdMetadata)

var GetProductsPacksByStoreIdTable = table.New(getProductsPacksByStoreIdMetadata)

var GetProductByStoreIDTable = table.New(getProductsByStoreIdMetadata)

var GetProductsDetailsByProductNameTable = table.New(getProductsDetailsByProductNameMetadata)

var GetOrdersByUserIDTable = table.New(getOrdersByUserIDMetadata)

var GetOrderDetailsByIDTable = table.New(getOrderDetailsByIDMetadata)

var PackRedemptionByUseridTable = table.New(packRedemptionByUseridMetadata)

var getUserDetailsByIDMetadata = table.Metadata{
	Name:    "get_user_details_by_id",
	Columns: []string{"username", "phone_no", "coins", "coupons_used"},
	PartKey: []string{"username"},
	SortKey: []string{},
}

var getAllStoreByCityMetadata = table.Metadata{
	Name:    "get_all_store_by_city",
	Columns: []string{"city", "store_name", "location", "closes_at"},
	PartKey: []string{"city"},
	SortKey: []string{"store_name"},
}

var getCuponCodeByStoreIdMetadata = table.Metadata{
	Name:    "cupon_code_by_store_id",
	Columns: []string{"store_name", "cupon_code", "expiry_date", "cupon_image"},
	PartKey: []string{"store_name"},
	SortKey: []string{"expiry_date", "cupon_code"},
}

var getCollectionsByStoreIdMetadata = table.Metadata{
	Name:    "get_collections_by_store_id",
	Columns: []string{"store_name", "is_featured", "collection_name"},
	PartKey: []string{"store_name"},
	SortKey: []string{"is_featured"},
}

var getProductsPacksByStoreIdMetadata = table.Metadata{
	Name:    "products_packs_by_store_id",
	Columns: []string{"store_name", "pack_name", "price", "description", "product_items", "prizes_10", "prizes_5", "prizes_3"},
	PartKey: []string{"store_name"},
	SortKey: []string{"pack_name"},
}

var getProductsByStoreIdMetadata = table.Metadata{
	Name:    "product_by_store_id",
	Columns: []string{"store_name", "product_name", "product_image", "price", "category", "collection_name"},
	PartKey: []string{"store_name"},
	SortKey: []string{"product_name"},
}

var getProductsDetailsByProductNameMetadata = table.Metadata{
	Name:    "get_products_details_by_product_name",
	Columns: []string{"store_name", "product_name", "product_image", "price", "category", "k_cal", "description", "ingredients"},
	PartKey: []string{"store_name"},
	SortKey: []string{"product_name"},
}

var getOrdersByUserIDMetadata = table.Metadata{
	Name:    "get_orders_by_user_id",
	Columns: []string{"username", "order_id", "order_time", "order_status"},
	PartKey: []string{"username"},
	SortKey: []string{"order_time", "order_id"},
}

var getOrderDetailsByIDMetadata = table.Metadata{
	Name:    "get_order_details_by_id",
	Columns: []string{"username", "order_id", "product_name", "product_quantity", "product_price", "order_status", "order_timestamp", "order_sub_total", "discount_percentage", "order_total", "pack_name", "is_pack"},
	PartKey: []string{"username", "order_id"},
	SortKey: []string{"product_name"},
}

var packRedemptionByUseridMetadata = table.Metadata{
	Name:    "pack_redemption_by_userid",
	Columns: []string{"username", "store_name", "order_id", "pack_name", "order_items_remaining", "expiry_timestamp"},
	PartKey: []string{"username"},
	SortKey: []string{"store_name", "pack_name"},
}
