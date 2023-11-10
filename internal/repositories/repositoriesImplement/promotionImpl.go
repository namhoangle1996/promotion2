package repositoriesImplement

import (
	"context"
	"gorm.io/gorm"
	"kk/internal/repositories/repositoriesInterface"
)

type PromotionRepo struct {
	db *gorm.DB
}

func (p PromotionRepo) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionRepo) Get(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionRepo) Update(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionRepo) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewPromotionRepo(db *gorm.DB) repositoriesInterface.PromotionInterface {
	return &PromotionRepo{
		db: db,
	}
}
