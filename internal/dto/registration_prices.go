package dto

type CreateRegistrationPricesRequest struct {
	RegistrationFee int64 `json:"registration_fee" validate:"required"`
	SystemChange    int64 `json:"system_change" validate:"required"`
	Total           int64 `json:"total" validate:"required"`
}

type UpdateRegistrationPricesRequest struct {
	RegistrationFee *int64 `json:"registration_fee" validate:"required"`
	SystemChange    *int64 `json:"system_change" validate:"required"`
	Total           *int64 `json:"total" validate:"required"`
}

type RegistrationPricesResponse struct {
	ID              int   `json:"id"`
	RegistrationFee int64 `json:"registration_fee" validate:"required"`
	SystemChange    int64 `json:"system_change" validate:"required"`
	Total           int64 `json:"total" validate:"required"`
}
