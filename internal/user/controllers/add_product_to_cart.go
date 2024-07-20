package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

type AddProductToCartRequest struct {
	StoreId   string  `json:"storeId"`
	OrderId   string  `json:"orderId"`
	ProductId string  `json:"productId"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}

func addProductToCart(s *user.OrderService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userId := ""

		request := new(AddProductToCartRequest)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest("Invalid request body"))
		}
		err := s.AddProductToCart(userId, request.OrderId, request.ProductId, request.StoreId, request.Quantity, request.Price)
		if err != nil {
			return c.Status(err.Code).JSON(messages.BadRequest(err.Error()))
		}
		return c.Status(fiber.StatusCreated).JSON(messages.SuccessResponse("Product added to cart successfully!!"))
	}
}
