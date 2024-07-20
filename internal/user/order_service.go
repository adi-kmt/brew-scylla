package user

import (
	"time"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
	"github.com/adi-kmt/brew-scylla/internal/utils"
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

func (s *OrderService) CheckoutCart(userId, orderId, storeName string, coins int64, couponCode string) (string, *messages.AppError) {
	user, err0 := s.userPort.GetUserDetailsByID(userId)
	if err0 != nil {
		return "", messages.BadRequest("User not found")
	}
	if coins > int64(user.Coins) || coins < 0 {
		return "", messages.BadRequest("Invalid number of coins")
	}

	couponCodeEntityList, err1 := s.orderPort.GetCouponsByStore(storeName)
	couponCodeList := utils.GetFieldSliceFromEntitySlice[entities.CouponCodeEntity](couponCodeEntityList, "CouponCode")
	if err1 != nil {
		return "", messages.BadRequest("Coupon not found")
	}
	if couponCode != "" {
		if utils.Contains[string](user.CouponsUsed, couponCode) {
			return "", messages.BadRequest("Coupon already used")
		} else if !utils.Contains[string](couponCodeList, couponCode) {
			return "", messages.BadRequest("Coupon code entered for wrong store")
		} else {
			user.CouponsUsed = append(user.CouponsUsed, couponCode)
		}
	}
	order := entities.OrderEntity{
		Username:    userId,
		OrderID:     orderId,
		OrderStatus: "Pending",
		OrderTime:   time.Now(),
	}
	err2 := s.orderPort.AddOrderToUser(order)
	if err2 != nil {
		return "", messages.InternalServerError("Unable to add order")
	}

	orderDetails, err3 := s.orderPort.GetOrderDetailsByUserAndOrderId(userId, orderId)
	if err3 != nil {
		return "", messages.BadRequest("Invalid order id")
	}
	if couponCode != "" {
		coupon, err5 := utils.GetEntityThatMatchesInSlice[entities.CouponCodeEntity](couponCodeEntityList, "CouponCode", couponCode)
		if err5 != nil {
			return "", messages.BadRequest("Coupon not found")
		}

		orderDetails.DiscountPercentage = coupon.Discount
		orderDetails.OrderTotal = orderDetails.OrderTotal * (100 - orderDetails.DiscountPercentage) / 100
	} else if coins > 0 {
		orderDetails.DiscountPercentage = float64(coins) * 0.3
		orderDetails.OrderTotal = orderDetails.OrderTotal * (100 - orderDetails.DiscountPercentage) / 100
	}

	orderDetails.OrderStatus = "Pending"

	err4 := s.orderPort.UpdateOrderDetailsByUserAndOrderId(userId, orderId, orderDetails)
	if err4 != nil {
		return "", messages.InternalServerError("Unable to update order")
	}

	if coins == 0 {
		user.Coins = user.Coins + int(0.3*float64(user.Coins))
		err3 := s.userPort.UpdateUserDetails(userId, user)
		if err3 != nil {
			return "", messages.InternalServerError("Unable to update user")
		}
		return "Order placed sucesfully", nil
	} else {
		user.Coins = user.Coins - int(coins)
		user.CouponsUsed = append(user.CouponsUsed, couponCode)
		err3 := s.userPort.UpdateUserDetails(userId, user)
		if err3 != nil {
			return "", messages.InternalServerError("Unable to update user")
		}
		return "Order placed sucesfully", nil
	}
}
