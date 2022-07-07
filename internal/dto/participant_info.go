package dto

type CreateParticipantInfoRequest struct {
	NISN             string `json:"nisn" validate:"required"`
	FinalReportScore int    `json:"final_report_score" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	FileRequirement  string `json:"file_requirement" validate:"required"`
}

type UpdateParticipantInfoRequest struct {
	NISN             *string `json:"nisn"`
	FinalReportScore *string `json:"final_report_score"`
	Email            *string `json:"email" validate:"email"`
	FileRequirement  *string `json:"file_requirement"`
}

type ParticipantInfoResponse struct {
	ID               int    `json:"id"`
	NISN             string `json:"nisn"`
	FinalReportScore string `json:"final_report_score"`
	Email            string `json:"email" validate:"email"`
	FileRequirement  string `json:"file_requirement"`
}
