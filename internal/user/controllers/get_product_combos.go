package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	"github.com/gofiber/fiber/v2"
)

func getProductPacks(s *user.ProductService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		storeId := c.Params("storeId")
		packs, err := s.GetProductPacks(storeId)
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(packs))
	}
}
