package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	FiberFailedParamError = func(paramName string, paramString string) error {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Failed to parse param %s with value %s", paramName, paramString))
	}

	FiberFailedParamIntError = func(paramName string, paramInt int) error {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Failed to parse param %s with value %d", paramName, paramInt))
	}

	FiberFailedBodyParseError = func(e error) error {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Failed to parse body: %s", e.Error()))
	}
)

func FiberError(code int, message string) *fiber.Error {
	return fiber.NewError(code, message)
}
