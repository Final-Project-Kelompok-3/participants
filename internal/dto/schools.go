package dto

type CreateSchoolsRequest struct {
	LevelsID int    `json:"levels_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type UpdateSchoolsRequest struct {
	LevelsID int    `json:"levels_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type SchoolsResponse struct {
	ID       int    `json:"id"`
	LevelsID int    `json:"levels_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
