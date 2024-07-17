package messages

import "github.com/gofiber/fiber/v2"

func SuccessResponseSlice[T any](data []T) *fiber.Map {
	return &fiber.Map{
		"data":  data,
		"error": "",
	}
}

func SuccessResponse[T any](data T) *fiber.Map {
	return &fiber.Map{
		"data":  data,
		"error": "",
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  "",
		"error": err.Error(),
	}
}
