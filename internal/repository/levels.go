package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type Levels interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Levels, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.Levels, error)
	Create(ctx context.Context, levels model.Levels) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type levels struct {
	Db *gorm.DB
}

func NewLevels(db *gorm.DB) *levels {
	return &levels{db}
}

func (r *levels) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Levels, *dto.PaginationInfo, error) {

	var (
		levels []model.Levels
		count  int64
	)

	query := r.Db.WithContext(ctx).Model(&model.Levels{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&levels).Error

	return levels, dto.CheckInfoPagination(paginate, count), err
}

func (r *levels) FindByID(ctx context.Context, ID uint) (model.Levels, error) {

	var levels model.Levels
	err := r.Db.WithContext(ctx).Model(&levels).Where("id = ?", ID).First(&levels).Error

	return levels, err
}

func (r *levels) Create(ctx context.Context, levels model.Levels) error {

	return r.Db.WithContext(ctx).Model(&model.Levels{}).Create(&levels).Error
}

func (r *levels) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.Levels{}).Updates(data).Error
	return err
}

func (r *levels) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.Levels{}).Error
	return err
}
