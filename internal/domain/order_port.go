package domain

import (
	"time"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/domain/entities"
)

type OrderPort interface {
	GetOrdersByUserId(userId string) ([]entities.OrderEntity, *messages.AppError)
	GetOrderDetailsByUserAndOrderId(userId, orderId string) (*entities.OrderDetailsEntity, *messages.AppError)
	UpdateOrderDetailsByUserAndOrderId(userId, orderId string, orderDetails *entities.OrderDetailsEntity) *messages.AppError
	AddProductToCart(userId, orderId, storeName, productName string, quantity int64, productPrice float64, orderTimestamp time.Time, orderStatus string, orderSubTotal, orderTotal float64, isPack bool, packName string) *messages.AppError
	AddOrderToUser(orderEntity entities.OrderEntity) *messages.AppError
	GetCouponsByStore(storeName string) ([]entities.CouponCodeEntity, *messages.AppError)
	AddPackRedemptionByUser(packRedemptionEntity entities.PackRedemptionEntity) *messages.AppError
}
