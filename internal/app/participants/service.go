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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Participants], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Participants, error)
	Create(ctx context.Context, payload *dto.CreateRegistrationParticipantsRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateParticipantsRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.Participants, error)
}

type service struct {
	ParticipantsRepository repository.Participants
}

func NewService(f *factory.Factory) Service {
	return &service{
		ParticipantsRepository: f.ParticipantsRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Participants], error) {

	Books, info, err := s.ParticipantsRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Participants])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Participants, error) {

	data, err := s.ParticipantsRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateRegistrationParticipantsRequest) (string, error) {

	var participants = model.Participants{
		Name:             payload.Name,
		Address:          payload.Address,
		NISN:             payload.NISN,
		FinalReportScore: payload.FinalReportScore,
		Email:            payload.Email,
		FileRequirement:  payload.FileRequirement,
	}

	err := s.ParticipantsRepository.Create(ctx, participants)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateParticipantsRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["title"] = payload.Name
	}
	if payload.NISN != nil {
		data["NISN"] = payload.NISN
	}

	err := s.ParticipantsRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.Participants, error) {
	data, err := s.ParticipantsRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.ParticipantsRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
