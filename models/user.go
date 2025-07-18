package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password"`
	Posts    []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
}
