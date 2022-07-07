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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationPrices], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationPrices, error)
	Create(ctx context.Context, payload *dto.CreateRegistrationPricesRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationPricesRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.RegistrationPrices, error)
}

type service struct {
	RegistrationPricesRepository repository.RegistrationPrices
}

func NewService(f *factory.Factory) Service {
	return &service{
		RegistrationPricesRepository: f.RegistrationPricesRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.RegistrationPrices], error) {

	Books, info, err := s.RegistrationPricesRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.RegistrationPrices])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.RegistrationPrices, error) {

	data, err := s.RegistrationPricesRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateRegistrationPricesRequest) (string, error) {

	var registration_prices = model.RegistrationPrices{
		RegistrationFee: payload.RegistrationFee,
		SystemChange:    payload.SystemChange,
		Total:           payload.Total,
	}

	err := s.RegistrationPricesRepository.Create(ctx, registration_prices)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateRegistrationPricesRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.RegistrationFee != nil {
		data["registration_fee"] = payload.RegistrationFee
	}
	if payload.SystemChange != nil {
		data["system_change"] = payload.SystemChange
	}
	if payload.Total != nil {
		data["total"] = payload.Total
	}

	err := s.RegistrationPricesRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.RegistrationPrices, error) {
	data, err := s.RegistrationPricesRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.RegistrationPricesRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
