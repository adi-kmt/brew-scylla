package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type OrderPort interface {
	GetOrdersByUserId(userId string) ([]entities.OrderEntity, *messages.AppError)
	GetOrderDetailsByUserAndOrderId(userId, orderId string) (*entities.OrderDetailsEntity, *messages.AppError)
	AddProductToCart(orderId, storeName, productName string, quantity int64) *messages.AppError
	CheckoutCart(orderId, storeName string, coins int64) *messages.AppError
}
