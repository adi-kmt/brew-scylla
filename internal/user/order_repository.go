package user

import (
	"time"

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

func (repo *OrderRepository) AddProductToCart(userId, orderId, storeName, productName string, quantity int64, productPrice float64, orderTimestamp time.Time, orderStatus string, orderSubTotal, orderTotal float64) *messages.AppError {
	orderModel := struct {
		UserId             string  `json:"user_id"`
		OrderId            string  `json:"order_id"`
		ProductName        string  `json:"product_name"`
		ProductQuantity    int64   `json:"product_quantity"`
		ProductPrice       float64 `json:"product_price"`
		OrderStatus        string  `json:"order_status"`
		OrderTimestamp     string  `json:"order_timestamp"`
		OrderSubTotal      float64 `json:"order_sub_total"`
		DiscountPercentage float64 `json:"discount_percentage"`
		OrderTotal         float64 `json:"order_total"`
		PackName           string  `json:"pack_name"`
		IsPack             bool    `json:"is_pack"`
	}{
		UserId:             userId,
		OrderId:            orderId,
		ProductName:        productName,
		ProductQuantity:    quantity,
		ProductPrice:       productPrice,
		OrderStatus:        orderStatus,
		OrderTimestamp:     orderTimestamp.Format(time.RFC3339),
		OrderSubTotal:      orderSubTotal,
		DiscountPercentage: 0,
		OrderTotal:         orderTotal,
		PackName:           "",
		IsPack:             false,
	}

	err := db.GetOrderDetailsByIDTable.InsertQuery(repo.session).BindStruct(orderModel).ExecRelease()
	if err != nil {
		return messages.InternalServerError("Unable to add product to cart")
	}
	return nil
}

func (repo *OrderRepository) GetCouponsByStore(storeName string) ([]entities.CouponCodeEntity, *messages.AppError) {
	var coupons []entities.CouponCodeEntity
	err := db.GetCuponCodeByStoreIdTable.SelectQuery(repo.session).Bind(storeName).Select(&coupons)
	if err != nil {
		return nil, messages.InternalServerError("Unable to get coupons")
	}
	return coupons, nil
}

func (repo *OrderRepository) AddOrderToUser(orderEntity entities.OrderEntity) *messages.AppError {
	err := db.GetOrdersByUserIDTable.InsertQuery(repo.session).BindStruct(orderEntity).ExecRelease()
	if err != nil {
		return messages.InternalServerError("Unable to add order")
	}
	return nil
}

func (repo *OrderRepository) UpdateOrderDetailsByUserAndOrderId(userId, orderId string, orderDetails *entities.OrderDetailsEntity) *messages.AppError {
	err := db.GetOrdersByUserIDTable.UpdateQuery(repo.session).BindStruct(orderDetails).Bind(userId, orderId).ExecRelease()
	if err != nil {
		return messages.InternalServerError("Unable to update order details")
	}
	return nil
}
