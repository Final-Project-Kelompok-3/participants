package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type Schools interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Schools, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Schools, error)
	Create(ctx context.Context, schools model.Schools) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type schools struct {
	Db *gorm.DB
}

func NewSchools(db *gorm.DB) *schools {
	return &schools{db}
}

func (r *schools) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Schools, *dto.PaginationInfo, error) {

	var (
		schools []model.Schools
		count   int64
	)

	query := r.Db.WithContext(ctx).Model(&model.Schools{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&schools).Error

	return schools, dto.CheckInfoPagination(paginate, count), err
}

func (r *schools) FindByID(ctx context.Context, ID uint) (model.Schools, error) {

	var schools model.Schools
	err := r.Db.WithContext(ctx).Model(&schools).Where("id = ?", ID).First(&schools).Error

	return schools, err
}

func (r *schools) Create(ctx context.Context, schools model.Schools) error {

	return r.Db.WithContext(ctx).Model(&model.Schools{}).Create(&schools).Error
}

func (r *schools) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Schools{}).Updates(data).Error
	return err
}

func (r *schools) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Schools{}).Error
	return err
}
