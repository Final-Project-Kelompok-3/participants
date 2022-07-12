package dto

import "time"

type CreateRegistrationPeriodsRequest struct {
	SchoolsID            uint `json:"schools_id" validate:"required"`
	RegistrationPricesID uint `json:"registration_prices_id" validate:"required"`

	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateRegistrationPeriodsRequest struct {
	SchoolsID            *uint `json:"schools_id"`
	RegistrationPricesID *uint `json:"registration_prices_id"`

	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
}

type RegistrationPeriodsResponse struct {
	ID                   uint      `json:"id"`
	SchoolsID            uint      `json:"schools_id" validate:"required"`
	RegistrationPricesID uint      `json:"registration_prices_id" validate:"required"`
	StartDate            time.Time `json:"start_date" validate:"required"`
	EndDate              time.Time `json:"end_date" validate:"required"`
	Description          string    `json:"description" validate:"required"`
}
