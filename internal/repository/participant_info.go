package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type ParticipantsInfo interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.ParticipantInfo, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.ParticipantInfo, error)
	Create(ctx context.Context, participantinfo model.ParticipantInfo) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type participantinfo struct {
	Db *gorm.DB
}

func NewParticipantInfo(db *gorm.DB) *participantinfo {
	return &participantinfo{db}
}

func (r *participantinfo) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.ParticipantInfo, *dto.PaginationInfo, error) {

	var (
		participantinfo []model.ParticipantInfo
		count           int64
	)

	query := r.Db.WithContext(ctx).Model(&model.ParticipantInfo{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&participantinfo).Error

	return participantinfo, dto.CheckInfoPagination(paginate, count), err
}

func (r *participantinfo) FindByID(ctx context.Context, ID uint) (model.ParticipantInfo, error) {

	var participantinfo model.ParticipantInfo
	err := r.Db.WithContext(ctx).Model(&participantinfo).Where("id = ?", ID).First(&participantinfo).Error

	return participantinfo, err
}

func (r *participantinfo) Create(ctx context.Context, participantinfo model.ParticipantInfo) error {

	return r.Db.WithContext(ctx).Model(&model.ParticipantInfo{}).Create(&participantinfo).Error
}

func (r *participantinfo) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.ParticipantInfo{}).Updates(data).Error
	return err
}

func (r *participantinfo) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.ParticipantInfo{}).Error
	return err
}
