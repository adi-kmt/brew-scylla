package user

import (
	"time"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/google/uuid"
)

type OrderService struct {
	orderPort   domain.OrderPort
	productPort domain.ProductPort
	userPort    domain.UserPort
}

func NewOrderService(orderPort domain.OrderPort, productPort domain.ProductPort, userPort domain.UserPort) *OrderService {
	return &OrderService{
		orderPort:   orderPort,
		productPort: productPort,
		userPort:    userPort,
	}
}

func (s *OrderService) GetOrders(userId string) ([]entities.OrderEntity, *messages.AppError) {
	return s.orderPort.GetOrdersByUserId(userId)
}

func (s *OrderService) GetOrderDetails(userId, orderId string) (*entities.OrderDetailsEntity, *messages.AppError) {
	return s.orderPort.GetOrderDetailsByUserAndOrderId(userId, orderId)
}

func (s *OrderService) AddProductToCart(userId, orderId, productName, storeName string, quantity int64, price float64) *messages.AppError {
	var productPrice float64 = 0
	var orderTimestamp time.Time
	var orderStatus string
	var orderTotal float64 = 0

	product, err0 := s.productPort.GetProductsDetailsByStore(storeName, productName)
	for _, priceMapValue := range product.Price {
		if priceMapValue == price {
			productPrice = price
		}
	}

	if productPrice == 0 {
		return messages.BadRequest("Price not found, Wrong product price")
	}
	if err0 != nil {
		return messages.BadRequest("Product not found")
	}
	if orderId == "" {
		newUUID, err := uuid.NewV7()
		if err != nil {
			return messages.InternalServerError("Unable to generate order id")
		}
		orderId = newUUID.String()
		orderTimestamp = time.Now()
		orderStatus = "initial"
		orderTotal = productPrice * float64(quantity)
	} else {
		order, err := s.orderPort.GetOrderDetailsByUserAndOrderId(userId, orderId)
		if err != nil {
			return messages.BadRequest("Invalid order id")
		}
		orderTimestamp = order.OrderTimestamp
		orderStatus = order.OrderStatus
		orderTotal = order.OrderTotal + productPrice*float64(quantity)
	}
	return s.orderPort.AddProductToCart(userId, orderId, productName, storeName, quantity, productPrice, orderTimestamp, orderStatus, orderTotal, orderTotal)
}

func (s *OrderService) CheckoutCart(userId, orderId, storeName string, coins int64, couponCode string) *messages.AppError {
	/*
		1. Check if the coins exist in user
		2. Check if coupon is valid and user user hasn't used it
		3. Add to cart
		4. If coins not used, then award coins to user
		5. If coupon used, then add used coupons to user
	*/
	return s.orderPort.CheckoutCart(userId, orderId, storeName, coins)
}
