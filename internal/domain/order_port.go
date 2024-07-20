package domain

import (
	"time"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type OrderPort interface {
	GetOrdersByUserId(userId string) ([]entities.OrderEntity, *messages.AppError)
	GetOrderDetailsByUserAndOrderId(userId, orderId string) (*entities.OrderDetailsEntity, *messages.AppError)
	AddProductToCart(userId, orderId, storeName, productName string, quantity int64, productPrice float64, orderTimestamp time.Time, orderStatus string, orderSubTotal, orderTotal float64) *messages.AppError
	CheckoutCart(userId, orderId, storeName string, coins int64) *messages.AppError
}
