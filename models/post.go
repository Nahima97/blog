package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
