package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type Participants interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Participants, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Participants, error)
	Create(ctx context.Context, participants model.Participants) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type participants struct {
	Db *gorm.DB
}

func NewParticipants(db *gorm.DB) *participants {
	return &participants{db}
}

func (r *participants) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Participants, *dto.PaginationInfo, error) {

	var (
		participants []model.Participants
		count        int64
	)

	query := r.Db.WithContext(ctx).Model(&model.Participants{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&participants).Error

	return participants, dto.CheckInfoPagination(paginate, count), err
}

func (r *participants) FindByID(ctx context.Context, ID uint) (model.Participants, error) {

	var participants model.Participants
	err := r.Db.WithContext(ctx).Model(&participants).Where("id = ?", ID).First(&participants).Error

	return participants, err
}

func (r *participants) Create(ctx context.Context, participants model.Participants) error {

	return r.Db.WithContext(ctx).Model(&model.Participants{}).Create(&participants).Error
}

func (r *participants) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Participants{}).Updates(data).Error
	return err
}

func (r *participants) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Participants{}).Error
	return err
}
