package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

func getOrdersByUserId(o *user.OrderService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orders, err := o.GetOrders(userId)
		if err != nil {
			return c.Status(err.Code).JSON(messages.BadRequest(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(orders))
	}
}
