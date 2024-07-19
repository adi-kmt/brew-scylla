package user_controllers

import (
	"net/url"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

func getFeaturedProductCollections(s *user.ProductService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		storeId := c.Params("storeId")
		storeIdUnescaped, err0 := url.QueryUnescape(storeId)
		if err0 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest(err0.Error()))
		}
		collections, err := s.GetFeaturedProductCollections(storeIdUnescaped)
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(collections))
	}
}
