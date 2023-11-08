package model

import "time"

type Promotion struct {
	BaseModel
	PromotionCode       string                    `json:"promotionCode" gorm:"varchar(20);uniqueIndex;not null"`
	Description         string                    `json:"description" `
	StartDate           time.Time                 `json:"startDate" gorm:"not null"`
	Status              PromotionStatus           `json:"status" gorm:"not null;VARCHAR(20);default:INACTIVE"`
	ProcessingStatus    PromotionProcessingStatus `json:"processingStatus" gorm:"not null;VARCHAR(20);  "`
	MinTransactionValue float64                   `json:"minTransactionValue" gorm:"default:0"`
	MaxTransactionValue float64                   `json:"maxTransactionValue" gorm:"null"`
	MaxUsagePerCustomer int                       `json:"maxUsagePerCustomer" gorm:"null"`
	Total               int                       `json:"total" gorm:"not null"`
	RemainingQuantity   int                       `json:"remainingQuantity" `
	MinDiscountValue    float64                   `json:"minDiscountValue" gorm:"default:0"`
	MaxDiscountValue    float64                   `json:"maxDiscountValue"`

	EndDate       time.Time    `json:"end_date" gorm:"not null"`
	DiscountType  DiscountType `json:"discount_type" gorm:"not null;default:FIXED_VALUE"`
	DiscountValue float64      `json:"discount_value" gorm:"not null"`
	RejectReason  string       `json:"rejectReason"  `
}

func (m *Promotion) TableName() string {
	return "promotion"
}

type DiscountType string
type PromotionStatus string
type PromotionProcessingStatus string

const (
	DiscountTypeFixedValue DiscountType = "FIXED_VALUE"
	DiscountTypePercentage DiscountType = "PERCENTAGE"
)

const (
	PromotionStatusActive   PromotionStatus = "ACTIVE"   // Đang hoạt động
	PromotionStatusInActive PromotionStatus = "INACTIVE" // Không hoạt động
	PromotionStatusInPaused PromotionStatus = "PAUSED"   // Tạm dừng
	PromotionStatusFinished PromotionStatus = "FINISHED" // Kết thúc
)
