package cmd

import (
	"fmt"
	"log"

	"github.com/ptaas-tool/gateway/internal/config"
	"github.com/ptaas-tool/gateway/internal/http"

	"github.com/ptaas-tool/base-api/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// API command is used to start API server
type API struct {
	Cfg config.Config
	Db  *gorm.DB
}

func (a API) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "build and start apt api server",
		Run: func(_ *cobra.Command, _ []string) {
			a.main()
		},
	}
}

func (a API) main() {
	// create new models interface
	modelsInstance := models.New(a.Db)

	// creating a new fiber app
	app := fiber.New()

	// use cors middleware for our application
	app.Use(cors.New())

	// register http
	http.Register{
		Config:          a.Cfg,
		ModelsInterface: modelsInstance,
	}.Create(app)

	// starting app on choosing port
	if er := app.Listen(fmt.Sprintf(":%d", a.Cfg.HTTP.Port)); er != nil {
		log.Fatalf("[api] failed to start api server error=%v", er)
	}
}
