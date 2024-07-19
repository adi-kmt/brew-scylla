package entities

import (
	"time"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
)

type OrderStatus string

type OrderEntity struct {
	Username    string      `json:"username"`
	OrderID     string      `json:"order_id"`
	OrderTime   time.Time   `json:"order_time"`
	OrderStatus OrderStatus `json:"order_status"`
}

const (
	OrderStatusPending  OrderStatus = "PENDING"
	OrderStatusApproved OrderStatus = "APPROVED"
	OrderStatusRejected OrderStatus = "REJECTED"
	OrderStatusDone     OrderStatus = "DONE"
)

func GetOrderStatus(status string) (OrderStatus, error) {
	switch status {
	case "PENDING":
		return OrderStatusPending, nil
	case "APPROVED":
		return OrderStatusApproved, nil
	case "REJECTED":
		return OrderStatusRejected, nil
	case "DONE":
		return OrderStatusDone, nil
	}
	return "", messages.BadRequest("Invalid order status")
}

func (status OrderStatus) String() (string, error) {
	switch status {
	case OrderStatusPending:
		return "PENDING", nil
	case OrderStatusApproved:
		return "APPROVED", nil
	case OrderStatusRejected:
		return "REJECTED", nil
	case OrderStatusDone:
		return "DONE", nil
	}
	return "", messages.BadRequest("Invalid order status")
}
