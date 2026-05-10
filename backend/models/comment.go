package models

import (
	"time"
)

// Model untuk komentar pada artikel
type Comment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ArticleID      uint      `json:"article_id" gorm:"not null"`
	CommenterName  string    `json:"commenter_name" gorm:"type:varchar(255);not null"`
	CommenterEmail string    `json:"commenter_email" gorm:"type:varchar(255);not null"`
	Content        string    `json:"content" gorm:"type:text;not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Relasi
	Article Article `json:"article,omitempty" gorm:"foreignKey:ArticleID"`
}
