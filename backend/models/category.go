package models

import (
	"time"
)

// Model untuk kategori artikel blog
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(255);uniqueIndex;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Color       string    `json:"color" gorm:"type:varchar(7);default:'#3B82F6'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relasi: Satu kategori memiliki banyak artikel
	Articles []Article `json:"articles,omitempty" gorm:"foreignKey:CategoryID"`
}
