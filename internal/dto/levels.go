package dto

type CreateLevelsRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateLevelsRequest struct {
	Name *string `json:"name"`
}

type LevelsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
