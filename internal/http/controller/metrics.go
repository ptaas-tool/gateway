package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// MetricsHandler returns cluster metrics
func (c Controller) MetricsHandler(ctx *fiber.Ctx) error {
	users, err := c.Models.Users.GetAll()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[metrics] failed to get users error=%w", err),
			MessageFailedEntityList,
		)
	}

	projects, err := c.Models.Projects.GetAll()
	if err != nil {
		return c.ErrHandler.ErrDatabase(
			ctx,
			fmt.Errorf("[metrics] failed to get projects error=%w", err),
			MessageFailedEntityList,
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"users":    len(users),
		"projects": len(projects),
		"core":     c.Config.HTTP.Core,
		"ftp":      c.Config.FTP.Host,
		"jwt":      c.Config.JWT.ExpireTime,
		"mysql":    fmt.Sprintf("%s:%d", c.Config.MySQL.Host, c.Config.MySQL.Port),
		"metrics":  c.Metrics,
	})
}
