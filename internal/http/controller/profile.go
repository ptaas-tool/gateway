package controller

import (
	"fmt"

	"github.com/apt-tool/apt-gateway/internal/http/request"
	"github.com/apt-tool/apt-gateway/internal/http/response"

	"github.com/apt-tool/apt-core/pkg/models/user"

	"github.com/gofiber/fiber/v2"
)

// GetProfile profile
func (c Controller) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(response.UserResponse{}.DTO(ctx.Locals("user").(*user.User)))
}

// UpdateProfile information
func (c Controller) UpdateProfile(ctx *fiber.Ctx) error {
	u := ctx.Locals("user").(*user.User)

	req := new(request.UserProfileRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.user.Update] failed to parse body error=%w", err))
	}

	u.Username = req.Name

	if len(req.Pass) > 0 {
		u.Password = req.Pass
	}

	if er := c.Models.Users.Update(u.ID, u); er != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.user.Update] failed to update user error=%w", er))
	}

	return ctx.SendStatus(fiber.StatusOK)
}