package usecaseInterface

import (
	"context"
	"kk/internal/model"
)

type PromotionServiceInterface interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	GetVoucherByCode(ctx context.Context, vcCode string) (*model.Promotion, error)
}
