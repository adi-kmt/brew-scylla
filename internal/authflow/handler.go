package authflow

import (
	"net/url"

	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/gofiber/fiber/v2"
)

type registerDto struct {
	PhoneNo int64 `json:"phone_no"`
}

func AuthHandler(router fiber.Router, s *AuthService) {
	router.Post("/register", func(c *fiber.Ctx) error {
		request := new(registerDto)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest("Invalid request body"))
		}

		userId, err := s.Register(request.PhoneNo)
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse("Registration successful, UserID is " + userId))
	})
	router.Get("/user/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		userIdUnescaped, err0 := url.QueryUnescape(userId)
		if err0 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest(err0.Error()))
		}
		user, err := s.GetUserDetailsByID(userIdUnescaped)
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse(user))
	})
}
