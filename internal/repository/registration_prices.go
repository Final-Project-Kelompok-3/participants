package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type RegistrationPrices interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationPrices, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.RegistrationPrices, error)
	Create(ctx context.Context, registrationprices model.RegistrationPrices) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type registrationprices struct {
	Db *gorm.DB
}

func NewRegistrationPrices(db *gorm.DB) *registrationprices {
	return &registrationprices{db}
}

func (r *registrationprices) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationPrices, *dto.PaginationInfo, error) {

	var (
		registrationprices []model.RegistrationPrices
		count              int64
	)

	query := r.Db.WithContext(ctx).Model(&model.RegistrationPrices{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&registrationprices).Error

	return registrationprices, dto.CheckInfoPagination(paginate, count), err
}

func (r *registrationprices) FindByID(ctx context.Context, ID uint) (model.RegistrationPrices, error) {

	var registrationprices model.RegistrationPrices
	err := r.Db.WithContext(ctx).Model(&registrationprices).Where("id = ?", ID).First(&registrationprices).Error

	return registrationprices, err
}

func (r *registrationprices) Create(ctx context.Context, registrationprices model.RegistrationPrices) error {

	return r.Db.WithContext(ctx).Model(&model.RegistrationPrices{}).Create(&registrationprices).Error
}

func (r *registrationprices) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.RegistrationPrices{}).Updates(data).Error
	return err
}

func (r *registrationprices) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.RegistrationPrices{}).Error
	return err
}
