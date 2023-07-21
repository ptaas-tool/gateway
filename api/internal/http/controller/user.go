package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/automated-pen-testing/api/internal/http/request"
	"github.com/automated-pen-testing/api/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

// UserRegister will create a new user into system.
func (c Controller) UserRegister(ctx *fiber.Ctx) error {
	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.user.Register] failed to parse body error=%w", err))
	}

	if err := req.Validate(); err != nil {
		return c.ErrHandler.ErrValidation(ctx, fmt.Errorf("[controller.user.Register] failed to validate request error=%w", err))
	}

	if err := c.Models.Users.Create(req.ToModel()); err != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.user.Register] failed to create user error=%w", err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// UserLogin logs in a user.
func (c Controller) UserLogin(ctx *fiber.Ctx) error {
	req := new(request.UserRegisterRequest)

	if err := ctx.BodyParser(req); err != nil {
		return c.ErrHandler.ErrBodyParser(ctx, fmt.Errorf("[controller.Loing] failed to parse body error=%w", err))
	}

	if err := req.Validate(); err != nil {
		return c.ErrHandler.ErrValidation(ctx, fmt.Errorf("[controller.user.Login] failed to validate request error=%w", err))
	}

	userTmp, err := c.Models.Users.Validate(req.Name, req.Pass)
	if err != nil {
		return c.ErrHandler.ErrRecordNotFound(ctx, fmt.Errorf("[controller.user.Login] username and password don't match error=%w", err))
	}

	token, etime, err := c.JWTAuthenticator.GenerateToken(userTmp.Username, userTmp.Role)
	if err != nil {
		return c.ErrHandler.ErrLogical(ctx, fmt.Errorf("[controller.Loing] failed to create token error=%w", err))
	}

	if er := c.RedisConnector.Set(userTmp.Username, strconv.Itoa(int(userTmp.Role)), etime.Sub(time.Now())); er != nil {
		return c.ErrHandler.ErrDatabase(ctx, fmt.Errorf("[controller.Loing] failed to save token error=%w", er))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.Token{
		Token: token,
	})
}
