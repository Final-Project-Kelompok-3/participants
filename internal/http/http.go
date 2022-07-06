package http

import (
	"github.com/labstack/echo/v4"

	"github.com/Final-Project-Kelompok-3/participants/internal/app/participants"

	"github.com/Final-Project-Kelompok-3/participants/internal/factory"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {

	// role.NewHandler(f).Route(e.Group("/roles"))
	// user.NewHandler(f).Route(e.Group("/users"))
	participants.NewHandler(f).Route(e.Group("/participants"))
}
