package user

import (
	"strconv"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/db"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
)

type OrderRepository struct {
	session gocqlx.Session
}

func NewOrderRepository(session gocqlx.Session) *OrderRepository {
	return &OrderRepository{
		session: session,
	}
}

func (repo *OrderRepository) GetOrderDetailsByUserAndOrderId(userId, orderId string) (*entities.OrderDetailsEntity, *messages.AppError) {
	var order entities.OrderDetailsEntity
	err := repo.session.Query(
		qb.Select(db.GetOrderDetailsByIDTable.Name()).
			Columns(db.GetOrderDetailsByIDTable.Metadata().Columns...).
			Where(db.GetOrderDetailsByIDTable.PrimaryKeyCmp()[:2]...).ToCql(),
	).Bind(userId, orderId).SelectRelease(&order)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get order details")
	}
	return &order, nil
}

func (repo *OrderRepository) GetOrdersByUserId(userId string) ([]entities.OrderEntity, *messages.AppError) {
	var orders []entities.OrderEntity
	err := db.GetOrdersByUserIDTable.SelectQuery(repo.session).Bind(userId).Select(&orders)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get orders")
	}
	return orders, nil
}

func (repo *OrderRepository) AddProductToCart(orderId, storeName, productName string, quantity int64) *messages.AppError {
	//TODO wrong
	if err := repo.session.Query("INSERT INTO order_items (order_id, store_name, product_name, quantity) VALUES (?, ?, ?, ?)", []string{orderId, storeName, productName, strconv.Itoa(int(quantity))}).Exec(); err != nil {
		return messages.InternalServerError("Unable to add product to cart")
	}
	return nil
}

func (repo *OrderRepository) CheckoutCart(orderId, storeName string, coins int64) *messages.AppError {
	//TODO wrong
	if err := repo.session.Query("UPDATE orders SET store_name = ?, coins = ? WHERE order_id = ?", []string{storeName, strconv.Itoa(int(coins)), orderId}).Exec(); err != nil {
		return messages.InternalServerError("Unable to checkout cart")
	}
	return nil
}
