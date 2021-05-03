package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ItemCode    string `gorm:"type:varchar(60);column:item_code" json:"itemCode"`
	Description string `gorm:"type:varchar(80);column:description" json:"description"`
	Quantity    int    `gorm:"column:quantity" json:"quantity"`
	OrderRef    uint   `gorm:"column:order_id" json:"orderId"`
}

func (r OrderItem) TableName() string {
	return "items"
}
