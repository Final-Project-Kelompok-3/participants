package participant_info

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
	FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.ParticipantInfo], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.ParticipantInfo, error)
	Create(ctx context.Context, payload *dto.CreateParticipantInfoRequest) (string, error)
	Update(ctx context.Context, ID uint, payload *dto.UpdateParticipantInfoRequest) (string, error)
	Delete(ctx context.Context, ID uint) (*model.ParticipantInfo, error)
}

type service struct {
	ParticipantsInfoRepository repository.ParticipantsInfo
}

func NewService(f *factory.Factory) Service {
	return &service{
		ParticipantsInfoRepository: f.ParticipantsInfoRepository,
	}
}

func (s *service) FindAll(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.ParticipantInfo], error) {

	Books, info, err := s.ParticipantsInfoRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.ParticipantInfo])
	result.Datas = Books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.ParticipantInfo, error) {

	data, err := s.ParticipantsInfoRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Create(ctx context.Context, payload *dto.CreateParticipantInfoRequest) (string, error) {

	var participant_info = model.ParticipantInfo{
		NISN:             payload.NISN,
		FinalReportScore: payload.FinalReportScore,
		Email:            payload.Email,
		FileRequirement:  payload.FileRequirement,
	}

	err := s.ParticipantsInfoRepository.Create(ctx, participant_info)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdateParticipantInfoRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.NISN != nil {
		data["nisn"] = payload.NISN
	}
	if payload.FinalReportScore != nil {
		data["final_report_score"] = payload.FinalReportScore
	}
	if payload.Email != nil {
		data["email"] = payload.Email
	}
	if payload.FileRequirement != nil {
		data["file_requirement"] = payload.FileRequirement
	}

	err := s.ParticipantsInfoRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID uint) (*model.ParticipantInfo, error) {
	data, err := s.ParticipantsInfoRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.ParticipantsInfoRepository.Delete(ctx, ID)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}
