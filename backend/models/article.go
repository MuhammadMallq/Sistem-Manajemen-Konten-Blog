package models

import (
	"time"
)

// Model untuk artikel blog
type Article struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"type:varchar(500);not null"`
	Slug        string     `json:"slug" gorm:"type:varchar(500);uniqueIndex;not null"`
	Content     string     `json:"content" gorm:"type:text;not null"`
	Excerpt     string     `json:"excerpt" gorm:"type:text"`
	CoverImage  string     `json:"cover_image" gorm:"type:varchar(500)"`
	Status      string     `json:"status" gorm:"type:varchar(20);default:'draft'"` // draft, published
	AuthorID    uint       `json:"author_id" gorm:"not null"`
	CategoryID  uint       `json:"category_id" gorm:"not null"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relasi
	Author   Author    `json:"author" gorm:"foreignKey:AuthorID"`
	Category Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:ArticleID"`
}
