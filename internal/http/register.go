package http

import (
	"github.com/ptaas-tool/gateway/internal/config"
	"github.com/ptaas-tool/gateway/internal/http/controller"
	"github.com/ptaas-tool/gateway/internal/http/controller/handler"
	"github.com/ptaas-tool/gateway/internal/http/middleware"
	"github.com/ptaas-tool/gateway/internal/utils/jwt"
	"github.com/ptaas-tool/gateway/pkg/client"

	"github.com/ptaas-tool/base-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Register struct {
	Config          config.Config
	ModelsInterface *models.Interface
}

func (r Register) Create(app *fiber.App) {
	// create new jwt authenticator
	authenticator := jwt.New(r.Config.JWT)

	// create an error handler for http service
	errHandler := handler.ErrorHandler{DevMode: r.Config.HTTP.DevMode}

	// create middleware
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
		ErrHandler:       errHandler,
	}

	// create controller
	ctl := controller.Controller{
		Config:           r.Config,
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		ErrHandler:       errHandler,
		Client:           client.NewClient(),
		Metrics: &controller.Metrics{
			SuccessfulRequests: 0,
			FailedRequests:     0,
			TotalDownloads:     0,
			TotalExecutes:      0,
		},
	}

	// health, logger and metrics
	app.Get("/metrics", ctl.MetricsHandler)
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// user login endpoint
	app.Post("/login", ctl.Login)

	// add auth middleware
	auth := app.Use(mid.Auth)

	// live tracking of project
	auth.Get("/live-tracking/:project_id", ctl.GetTracksList) // last id in query param for filter

	// users crud
	users := auth.Group("/users")
	users.Get("/", ctl.GetUsersList)
	users.Post("/", ctl.CreateUser)
	users.Delete("/:id", ctl.DeleteUser)

	// projects crud
	projects := auth.Group("/projects")
	projects.Get("/", ctl.GetProjectsList)
	projects.Post("/", ctl.CreateProject)
	projects.Get("/:id", ctl.GetProject)
	projects.Post("/:id", ctl.ExecuteProject)
	projects.Delete("/:id", ctl.DeleteProject)
	projects.Get("/:id/documents/:document_id", ctl.DownloadProjectDocument)
	projects.Post("/:id/documents/:document_id", ctl.RerunDocument)
}
