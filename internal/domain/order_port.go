package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
)

type OrderPort interface {
	GetOrdersByUserId(userId string) ([]OrderDTO, *messages.AppError)
	GetOrderDetailsByUserAndOrderId(userId, orderId string) (OrderDetailsDTO, *messages.AppError)
	AddProductToCart(orderId, storeName, productName string, quantity int64) *messages.AppError
	CheckoutCart(orderId, storeName string, coins int64) *messages.AppError
}
