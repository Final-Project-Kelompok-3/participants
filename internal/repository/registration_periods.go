package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type RegistrationPeriods interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationPeriods, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.RegistrationPeriods, error)
	Create(ctx context.Context, registrationperiods model.RegistrationPeriods) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type registrationperiods struct {
	Db *gorm.DB
}

func NewRegistrationPeriods(db *gorm.DB) *registrationperiods {
	return &registrationperiods{db}
}

func (r *registrationperiods) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationPeriods, *dto.PaginationInfo, error) {

	var (
		registrationperiods []model.RegistrationPeriods
		count               int64
	)

	query := r.Db.WithContext(ctx).Model(&model.RegistrationPeriods{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&registrationperiods).Error

	return registrationperiods, dto.CheckInfoPagination(paginate, count), err
}

func (r *registrationperiods) FindByID(ctx context.Context, ID uint) (model.RegistrationPeriods, error) {

	var registrationperiods model.RegistrationPeriods
	err := r.Db.WithContext(ctx).Model(&registrationperiods).Where("id = ?", ID).First(&registrationperiods).Error

	return registrationperiods, err
}

func (r *registrationperiods) Create(ctx context.Context, registrationperiods model.RegistrationPeriods) error {

	return r.Db.WithContext(ctx).Model(&model.RegistrationPeriods{}).Create(&registrationperiods).Error
}

func (r *registrationperiods) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.RegistrationPeriods{}).Updates(data).Error
	return err
}

func (r *registrationperiods) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.RegistrationPeriods{}).Error
	return err
}
