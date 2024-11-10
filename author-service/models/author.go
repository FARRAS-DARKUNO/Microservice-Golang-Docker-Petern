package models

import "gorm.io/gorm"

// Author adalah model untuk penulis buku
type Author struct {
	gorm.Model
	AuthorID uint   `json:"author_id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}
