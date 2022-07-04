package dto

type CreateRegistrationParticipantsRequest struct {
	Name             string `json:"name" validate:"required"`
	Address          string `json:"address" validate:"required"`
	NISN             string `json:"nisn" validate:"required"`
	FinalReportScore int    `json:"final_report_score" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	FileRequirement  string `json:"filerequirement" validate:"required"`
}

type UpdateParticipantsRequest struct {
	Name             *string `json:"name"`
	Address          *string `json:"address"`
	NISN             *string `json:"nisn"`
	FinalReportScore *string `json:"final_report_score"`
	Email            *string `json:"email" validate:"email"`
	FileRequirement  *string `json:"file_requirement"`
}

type ParticipantsResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	NISN             string `json:"nisn"`
	FinalReportScore string `json:"final_report_score"`
	Email            string `json:"email" validate:"email"`
	FileRequirement  string `json:"file_requirement"`
}
