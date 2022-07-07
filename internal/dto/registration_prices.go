package dto

type CreateRegistrationPricesRequest struct {
	RegistrationFee int `json:"registration_fee" validate:"required"`
	SystemChange    int `json:"system_change" validate:"required"`
	Total           int `json:"total" validate:"required"`
}

type UpdateRegistrationPricesRequest struct {
	RegistrationFee int `json:"registration_fee" validate:"required"`
	SystemChange    int `json:"system_change" validate:"required"`
	Total           int `json:"total" validate:"required"`
}

type RegistrationPricesResponse struct {
	ID              int `json:"id"`
	RegistrationFee int `json:"registration_fee" validate:"required"`
	SystemChange    int `json:"system_change" validate:"required"`
	Total           int `json:"total" validate:"required"`
}
