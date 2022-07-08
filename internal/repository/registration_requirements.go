package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type RegistrationRequirements interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationRequirements, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.RegistrationRequirements, error)
	Create(ctx context.Context, registrationrequirements model.RegistrationRequirements) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type registrationrequirements struct {
	Db *gorm.DB
}

func NewRegistrationRequirements(db *gorm.DB) *registrationrequirements {
	return &registrationrequirements{db}
}

func (r *registrationrequirements) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.RegistrationRequirements, *dto.PaginationInfo, error) {

	var (
		registrationrequirements []model.RegistrationRequirements
		count                    int64
	)

	query := r.Db.WithContext(ctx).Model(&model.RegistrationRequirements{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&registrationrequirements).Error

	return registrationrequirements, dto.CheckInfoPagination(paginate, count), err
}

func (r *registrationrequirements) FindByID(ctx context.Context, ID uint) (model.RegistrationRequirements, error) {

	var registrationrequirements model.RegistrationRequirements
	err := r.Db.WithContext(ctx).Model(&registrationrequirements).Where("id = ?", ID).First(&registrationrequirements).Error

	return registrationrequirements, err
}

func (r *registrationrequirements) Create(ctx context.Context, registrationrequirements model.RegistrationRequirements) error {

	return r.Db.WithContext(ctx).Model(&model.RegistrationRequirements{}).Create(&registrationrequirements).Error
}

func (r *registrationrequirements) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.RegistrationRequirements{}).Updates(data).Error
	return err
}

func (r *registrationrequirements) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.RegistrationRequirements{}).Error
	return err
}
