package user

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
)

type OrderService struct {
	orderPort domain.OrderPort
}

func NewOrderService(orderPort domain.OrderPort) *OrderService {
	return &OrderService{
		orderPort: orderPort,
	}
}

func (s *OrderService) GetOrders(userId string) ([]domain.OrderDTO, *messages.AppError) {
	return s.orderPort.GetOrdersByUserId(userId)
}

func (s *OrderService) GetOrderDetails(userId, orderId string) (domain.OrderDetailsDTO, *messages.AppError) {
	return s.orderPort.GetOrderDetailsByUserAndOrderId(userId, orderId)
}

func (s *OrderService) AddProductToCart(orderId, productName, storeName string, quantity int64) *messages.AppError {
	return s.orderPort.AddProductToCart(orderId, productName, storeName, quantity)
}

func (s *OrderService) CheckoutCart(orderId, storeName string, coins int64) *messages.AppError {
	return s.orderPort.CheckoutCart(orderId, storeName, coins)
}
