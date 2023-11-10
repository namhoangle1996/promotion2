package usecaseImple

import (
	"context"
	"kk/internal/model"
	"kk/internal/repositories/repositoriesInterface"
	"kk/internal/usecase/usecaseInterface"
)

type PromotionUsecase struct {
	promotionRepo repositoriesInterface.PromotionInterface
}

func (p PromotionUsecase) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionUsecase) Get(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionUsecase) Update(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionUsecase) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p PromotionUsecase) GetVoucherByCode(ctx context.Context, vcCode string) (*model.Promotion, error) {
	//TODO implement me
	panic("implement me")
}

func NewPromotionUsecase(rp repositoriesInterface.PromotionInterface) usecaseInterface.PromotionServiceInterface {
	return &PromotionUsecase{promotionRepo: rp}
}
