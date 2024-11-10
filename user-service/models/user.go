package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleCode int    `json:"role_code"`  // 1 for Admin, 2 for Pengunjung
	Role     string `json:"role"`       // Admin or Pengunjung
}
