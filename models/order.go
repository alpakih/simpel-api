package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string      `gorm:"type:varchar(60);column:customer_name" json:"customerName"`
	OrderAt      time.Time   `gorm:"column:order_at" json:"orderedAt"`
	OrderItem    []OrderItem `gorm:"foreignKey:OrderRef;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items"`
}

func (r Order) TableName() string {
	return "orders"
}
