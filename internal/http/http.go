package http

import (
	"github.com/labstack/echo/v4"

	"github.com/Final-Project-Kelompok-3/participants/internal/app/file_requirements"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/levels"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/participant_info"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/participants"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/registration_periods"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/registration_prices"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/registration_requirements"
	"github.com/Final-Project-Kelompok-3/participants/internal/app/schools"

	"github.com/Final-Project-Kelompok-3/participants/internal/factory"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	file_requirements.NewHandler(f).Route(e.Group("/file_requirements"))
	levels.NewHandler(f).Route(e.Group("/levels"))
	participant_info.NewHandler(f).Route(e.Group("/participant_info"))
	participants.NewHandler(f).Route(e.Group("/participants"))
	registration_periods.NewHandler(f).Route(e.Group("/registration_periods"))
	registration_prices.NewHandler(f).Route(e.Group("/registration_prices"))
	registration_requirements.NewHandler(f).Route(e.Group("/registration_requirements"))
	schools.NewHandler(f).Route(e.Group("/schools"))
}
