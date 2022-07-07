package schools

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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Schools], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Schools, error)
	Create(ctx context.Context, payload *dto.CreateSchoolsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateSchoolsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Schools, error)
}

type service struct {
	SchoolsRepository repository.Schools
}

func NewService(f *factory.Factory) Service {
	return &service{
		SchoolsRepository: f.SchoolsRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Schools], error) {

	Books, info, err := s.SchoolsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Schools])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Schools, error) {

	data, err := s.SchoolsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateSchoolsRequest) (string, error) {

	var schools = model.Schools{
		Name:     payload.Name,
		Address:  payload.Address,
		LevelsID: payload.LevelsID,
	}

	err := s.SchoolsRepository.Create(ctx, schools)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateSchoolsRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Address != nil {
		data["address"] = payload.Address
	}
	if payload.LevelsID != nil {
		data["levels_id"] = payload.LevelsID
	}

	err := s.SchoolsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Schools, error) {
	data, err := s.SchoolsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.SchoolsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
