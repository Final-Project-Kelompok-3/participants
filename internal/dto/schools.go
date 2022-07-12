package dto

type CreateSchoolsRequest struct {
	LevelsID uint   `json:"levels_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type UpdateSchoolsRequest struct {
	LevelsID *uint   `json:"levels_id"`
	Name     *string `json:"name"`
	Address  *string `json:"address"`
}

type SchoolsResponse struct {
	ID       uint   `json:"id"`
	LevelsID uint   `json:"levels_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
