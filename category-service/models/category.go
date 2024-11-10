package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryID uint   `json:"category_id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Description string `json:"description"`
}