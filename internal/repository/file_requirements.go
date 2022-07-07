package repository

import (
	"context"
	"strings"

	"github.com/Final-Project-Kelompok-3/participants/internal/dto"
	"github.com/Final-Project-Kelompok-3/participants/internal/model"

	"gorm.io/gorm"
)

type FileRequirements interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.FileRequirements, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.FileRequirements, error)
	Create(ctx context.Context, filerequirements model.FileRequirements) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type filerequirements struct {
	Db *gorm.DB
}

func NewFileRequirements(db *gorm.DB) *filerequirements {
	return &filerequirements{db}
}

func (r *filerequirements) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.FileRequirements, *dto.PaginationInfo, error) {

	var (
		filerequirements []model.FileRequirements
		count            int64
	)

	query := r.Db.WithContext(ctx).Model(&model.FileRequirements{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&filerequirements).Error

	return filerequirements, dto.CheckInfoPagination(paginate, count), err
}

func (r *filerequirements) FindByID(ctx context.Context, ID uint) (model.FileRequirements, error) {

	var filerequirements model.FileRequirements
	err := r.Db.WithContext(ctx).Model(&filerequirements).Where("id = ?", ID).First(&filerequirements).Error

	return filerequirements, err
}

func (r *filerequirements) Create(ctx context.Context, filerequirements model.FileRequirements) error {

	return r.Db.WithContext(ctx).Model(&model.Participants{}).Create(&filerequirements).Error
}

func (r *filerequirements) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.FileRequirements{}).Updates(data).Error
	return err
}

func (r *filerequirements) Delete(ctx context.Context, ID uint) error {
	err := r.Db.WithContext(ctx).Where("id = ?", ID).Delete(&model.FileRequirements{}).Error
	return err
}
