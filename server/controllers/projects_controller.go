package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/queries"
)

func HandleGetProjects(c *fiber.Ctx) error {
	return c.JSON(queries.GetAllProjects())
}
