package model

import "time"

type PromotionUsage struct {
	BaseModel
	PromotionCode string    `json:"promotionCode" gorm:"column:promotionCode"`
	CustomerID    int       `json:"customerID" gorm:"column:customerID"`
	TransactionID int       `json:"transactionID" gorm:"column:transactionID"`
	UsageDate     time.Time `json:"usageDate;default:CURRENT_TIMESTAMP" gorm:"column:usageDate"`
}

func (m *PromotionUsage) TableName() string {
	return "promotionUsage"
}
