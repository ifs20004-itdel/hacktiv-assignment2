package models

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `gorm:"not_null;type:varchar(191)" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"items"`
}
