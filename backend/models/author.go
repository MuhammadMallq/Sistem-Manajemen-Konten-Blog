package models

import (
	"time"
)

//  Model untuk data penulis artikel blog
type Author struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Bio       string    `json:"bio" gorm:"type:text"`
	Avatar    string    `json:"avatar" gorm:"type:varchar(500)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi: Satu penulis memiliki banyak artikel
	Articles []Article `json:"articles,omitempty" gorm:"foreignKey:AuthorID"`
}
