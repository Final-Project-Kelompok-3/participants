package participants

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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationRequirements], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationRequirements, error)
	Create(ctx context.Context, payload *dto.CreateRegistrationRequirementsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationRequirementsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.RegistrationRequirements, error)
}

type service struct {
	RegistrationRequirementsRepository repository.RegistrationRequirements
}

func NewService(f *factory.Factory) Service {
	return &service{
		RegistrationRequirementsRepository: f.RegistrationRequirementsRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationRequirements], error) {

	Books, info, err := s.RegistrationRequirementsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.RegistrationRequirements])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationRequirements, error) {

	data, err := s.RegistrationRequirementsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateRegistrationRequirementsRequest) (string, error) {

	var registration_requirements = model.RegistrationRequirements{
		FileRequirementsID:    payload.FileRequirementsID,
		RegistrationPeriodsID: payload.RegistrationPeriodsID,
	}

	err := s.RegistrationRequirementsRepository.Create(ctx, registration_requirements)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationRequirementsRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.FileRequirementsID != nil {
		data["file_requirements_id"] = payload.FileRequirementsID
	}
	if payload.RegistrationPeriodsID != nil {
		data["registration_requirements_id"] = payload.RegistrationPeriodsID
	}

	err := s.RegistrationRequirementsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.RegistrationRequirements, error) {
	data, err := s.RegistrationRequirementsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.RegistrationRequirementsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
