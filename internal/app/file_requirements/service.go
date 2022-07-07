package file_requirements

import (
	"context"

	"github.com/Final-Project-Kelompok-3/participants/pkg/constant"
	res "github.com/Final-Project-Kelompok-3/participants/pkg/util/response"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/factory"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"github.com/Final-Project-Kelompok-3/participants/internal/repository"
)

type Service interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.FileRequirements], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.FileRequirements, error)
	Create(ctx context.Context, payload *dto.CreateFileRequirementsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateFileRequirementsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.FileRequirements, error)
}

type service struct {
	FileRequirementsRepository repository.FileRequirements
}

func NewService(f *factory.Factory) Service {
	return &service{
		FileRequirementsRepository: f.FileRequirementsRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.FileRequirements], error) {

	Books, info, err := s.FileRequirementsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.FileRequirements])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.FileRequirements, error) {

	data, err := s.FileRequirementsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateFileRequirementsRequest) (string, error) {

	var file_requirements = model.FileRequirements{
		Name:          payload.Name,
		FileExtension: payload.FileExtension,
	}

	err := s.FileRequirementsRepository.Create(ctx, file_requirements)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateFileRequirementsRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.FileExtension != nil {
		data["file_extension"] = payload.FileExtension
	}

	err := s.FileRequirementsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.FileRequirements, error) {
	data, err := s.FileRequirementsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.FileRequirementsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
