package model

import "time"

type PromotionUsage struct {
	BaseModel
	PromotionId   string    `json:"promotionId" gorm:"column:promotionId;not null"`
	CustomerID    int       `json:"customerID" gorm:"column:customerID"`
	TransactionID int       `json:"transactionID" gorm:"column:transactionID;not null"`
	UsageDate     time.Time `json:"usageDate;default:CURRENT_TIMESTAMP" gorm:"column:usageDate"`
}

func (m *PromotionUsage) TableName() string {
	return "promotionUsage"
}
