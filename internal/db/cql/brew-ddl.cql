-- GetUserDetailsByID table
CREATE TABLE get_user_details_by_id (
    username text PRIMARY KEY,
    phone_no bigint,
    coins int,
    coupons_used set<text>
);

-- GetAllStoreByCity table
CREATE TABLE get_all_store_by_city (
    city text,
    store_name text,
    location frozen<map<text, text>>,
    closes_at timestamp,
    PRIMARY KEY ((city), store_name)
) WITH CLUSTERING ORDER BY (store_name ASC);

-- CuponCodeByStoreId table (with TTL)
CREATE TABLE cupon_code_by_store_id (
    store_name text,
    cupon_code text,
    expiry_date timestamp,
    cupon_image text,
    PRIMARY KEY ((store_name), expiry_date, cupon_code)
) WITH CLUSTERING ORDER BY (expiry_date ASC);

-- GetCollectionsByStoreId table
CREATE TABLE get_collections_by_store_id (
    store_name text,
    is_featured boolean,
    collection_name text,
    PRIMARY KEY ((store_name), is_featured)
);

-- ProductsPacksByStoreId table
CREATE TABLE products_packs_by_store_id (
    store_name text,
    pack_name text,
    price float,
    description text,
    product_items frozen<list<text>>,
    prizes_10 float,
    prizes_5 float,
    prizes_3 float,
    PRIMARY KEY ((store_name), pack_name)
);

-- ProductByStoreID table
CREATE TABLE product_by_store_id (
    store_name text,
    product_name text,
    product_image text,
    price map<text, float>,
    category text,
    collection_name text,
    PRIMARY KEY (store_name, product_name)
) WITH CLUSTERING ORDER BY (product_name ASC);

-- GetProductsDetailsByProductName table
CREATE TABLE get_products_details_by_product_name (
    store_name text,
    product_name text,
    product_image text,
    price map<text, float>,
    category text,
    k_cal int,
    description text,
    ingredients frozen<list<text>>,
    PRIMARY KEY ((store_name), product_name)
) WITH CLUSTERING ORDER BY (product_name ASC);

-- GetOrdersByUserID table
CREATE TABLE get_orders_by_user_id (
    username text,
    order_id text,
    order_time timestamp,
    order_status text,
    PRIMARY KEY ((username), order_time, order_id)
) WITH CLUSTERING ORDER BY (order_time DESC);

-- GetOrderDetailsByID table
CREATE TABLE get_order_details_by_id (
    username text,
    order_id text,
    product_name text,
    product_quantity int,
    product_price float,
    order_status text STATIC,
    order_timestamp timestamp STATIC,
    order_sub_total float STATIC,
    discount_percentage float STATIC,
    order_total float STATIC,
    pack_name text STATIC,
    is_pack boolean STATIC,
    PRIMARY KEY ((username, order_id), product_name)
) WITH CLUSTERING ORDER BY (product_name ASC);

-- PackRedemptionByUserid table
CREATE TABLE pack_redemption_by_userid (
    username text,
    store_name text,
    order_id text,
    pack_name text,
    order_items_remaining int,
    expiry_timestamp timestamp,
    PRIMARY KEY ((username), store_name, pack_name)
) WITH CLUSTERING ORDER BY (store_name ASC);
