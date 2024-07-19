package authflow

import (
	"github.com/adi-kmt/brew-scylla/internal/common/messages"
	"github.com/gofiber/fiber/v2"
)

type loginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerDto struct {
	PhoneNo int64 `json:"phoneNo"`
}

func AuthHandler(router fiber.Router, s *authService) {
	router.Post("/login", func(c *fiber.Ctx) error {
		request := new(loginDto)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest("Invalid request body"))
		}

		//TODO return token
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse("Login successful"))
	})
	router.Post("/register", func(c *fiber.Ctx) error {
		request := new(registerDto)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(messages.BadRequest("Invalid request body"))
		}

		err := s.Register(request.PhoneNo)
		if err != nil {
			return c.Status(err.Code).JSON(messages.InternalServerError(err.Error()))
		}
		//TODO return token
		return c.Status(fiber.StatusOK).JSON(messages.SuccessResponse("Registration successful"))
	})
}
