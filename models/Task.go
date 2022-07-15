package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(100);not null;unique_index" json:"title"`
	Description string         `json:"description" json:"description"`
	Done        bool           `gorm:"default:false" json:"done"`
	UserID      uint           `json:"user_id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}
