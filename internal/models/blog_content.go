package models

import "time"

// blog Content struct
type Blog_Content struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsActive  bool      `gorm:"default:false;column:is_active" json:"-"`
	CreatedBy int       `gorm:"not null" json:"created_by"`
	UpdatedBy int       `gorm:"not null" json:"updated_by"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
