package user_controllers

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/adi-kmt/brew-scylla/internal/user"
	user_dtos "github.com/adi-kmt/brew-scylla/internal/user/dtos"
	"github.com/gofiber/fiber/v2"
)

func getProductsByStore(s *user.ProductService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var products []user_dtos.ProductsWithReviewsDTO
		var err *messages.AppError
		storeId := c.Params("storeId")
		productCollection := c.Query("productCollection")
		if productCollection != "" {
			products, err = s.GetProductsByProductCollection(productCollection, storeId)
		} else {
			products, err = s.GetProductsByStore(storeId)
		}
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponseSlice(products))
	}
}
