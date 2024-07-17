package domain

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
)

type OrderRepository interface {
	GetOrdersByUserId(userName string) ([]OrderDTO, *messages.AppError)
	GetOrderDetailsByUserAndOrderId(userName, orderId string) (OrderDetailsDTO, *messages.AppError)
	AddProductToCart(orderId, storeId, productId string, quantity int64) *messages.AppError
	CheckoutCart(orderId, storeId string, coins int64) *messages.AppError
}
