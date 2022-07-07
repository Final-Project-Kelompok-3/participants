package dto

type CreateFileRequirementsRequest struct {
	Name          string `json:"name"`
	FileExtension string `json:"file_extension"`
}

type UpdateFileRequirementsRequest struct {
	Name          *string `json:"name"`
	FileExtension *string `json:"file_extension"`
}

type FileRequirementsResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FileExtension string `json:"file_extension"`
}
