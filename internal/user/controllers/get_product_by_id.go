package user_controllers

import (
	"net/url"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

func getProductById(s *user.ProductService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		productName := c.Params("productId")
		productNameUnescaped, err0 := url.QueryUnescape(productName)
		if err0 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest(err0.Error()))
		}
		storeName := c.Params("storeName")
		product, err := s.GetProductDetailsByStore(storeName, productNameUnescaped)
		if err != nil {
			return c.Status(err.Code).JSON(messages.BadRequest(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(product))
	}
}
