package controllers

import "github.com/gofiber/fiber/v2"

func CreateFlight(c *fiber.Ctx) error {
	_, err := c.Write([]byte("Tudo ok"))
	return err
}
