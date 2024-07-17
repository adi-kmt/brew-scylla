package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

type CheckoutOrderRequest struct {
	OrderID string `json:"order_id"`
	Coins   int64  `json:"coins"`
}

func checkoutOrder(s *user.OrderService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userId := ""

		request := new(CheckoutOrderRequest)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest("Invalid request body"))
		}
		err := s.CheckoutCart(userId, request.OrderID, request.Coins)
		if err != nil {
			return c.Status(err.Code).JSON(messages.BadRequest(err.Error()))
		}
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse("Order checked out successfully!!"))
	}
}
