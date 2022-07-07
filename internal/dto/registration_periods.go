package dto

import "time"

type CreateRegistrationPeriodsRequest struct {
	SchoolsID            int `json:"schools_id" validate:"required"`
	RegistrationPricesID int `json:"registration_prices_id" validate:"required"`

	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

type UpdateRegistrationPeriodsRequest struct {
	SchoolsID            *int `json:"schools_id" validate:"required"`
	RegistrationPricesID *int `json:"registration_prices_id" validate:"required"`

	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
}

type RegistrationPeriodsResponse struct {
	ID                   int    `json:"id"`
	SchoolsID            int    `json:"schools_id" validate:"required"`
	RegistrationPricesID int    `json:"registration_prices_id" validate:"required"`
	StartDate            string `json:"start_date" validate:"required"`
	EndDate              string `json:"end_date" validate:"required"`
	Description          string `json:"description" validate:"required"`
}
