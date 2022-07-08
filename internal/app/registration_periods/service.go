package registration_periods

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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationPeriods], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationPeriods, error)
	Create(ctx context.Context, payload *dto.CreateRegistrationPeriodsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationPeriodsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.RegistrationPeriods, error)
}

type service struct {
	RegistrationPeriodsRepository repository.RegistrationPeriods
}

func NewService(f *factory.Factory) Service {
	return &service{
		RegistrationPeriodsRepository: f.RegistrationPeriodsRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationPeriods], error) {

	Books, info, err := s.RegistrationPeriodsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.RegistrationPeriods])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationPeriods, error) {

	data, err := s.RegistrationPeriodsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateRegistrationPeriodsRequest) (string, error) {

	var registration_periods = model.RegistrationPeriods{
		StartDate:            payload.StartDate,
		EndDate:              payload.EndDate,
		Description:          payload.Description,
		RegistrationPricesID: payload.RegistrationPricesID,
		SchoolsID:            payload.SchoolsID,
	}

	err := s.RegistrationPeriodsRepository.Create(ctx, registration_periods)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationPeriodsRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.StartDate != nil {
		data["start_date"] = payload.StartDate
	}
	if payload.EndDate != nil {
		data["end_date"] = payload.EndDate
	}
	if payload.Description != nil {
		data["description"] = payload.Description
	}
	if payload.RegistrationPricesID != nil {
		data["registration_prices_id"] = payload.RegistrationPricesID
	}
	if payload.SchoolsID != nil {
		data["schools"] = payload.SchoolsID
	}

	err := s.RegistrationPeriodsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.RegistrationPeriods, error) {
	data, err := s.RegistrationPeriodsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.RegistrationPeriodsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
