package models

import (
	"time"
)

type Post struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null;column:title" json:"title"`
	Content     string    `gorm:"type:text;not null;column:content" json:"content"`
	Category    string    `gorm:"type:varchar(100);not null;column:category" json:"category"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null;column:status" json:"status"`
}

func (Post) TableName() string {
	return "posts"
}
