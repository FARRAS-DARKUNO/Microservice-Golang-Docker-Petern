package models

import (
	"time"

	"gorm.io/gorm"
)

// BookManagement adalah model untuk tabel buku
type BookManagement struct {
	gorm.Model
	BookID      uint     `json:"book_id" gorm:"primaryKey;autoIncrement"`
	Title       string   `json:"title"`
	AuthorID    uint     `json:"author_id"`
	CategoryID  uint     `json:"category_id"`
	Description string   `json:"description"`
	// Author      Author   `gorm:"foreignKey:AuthorID"`
	// Category    Category `gorm:"foreignKey:CategoryID"`
}

// BookStock adalah model untuk manajemen stok buku
type BookStock struct {
	gorm.Model
	BookStockID uint           `json:"book_stock_management_id" gorm:"primaryKey;autoIncrement"`
	Stock       int            `json:"stock"`
	BookID      uint           `json:"book_id"`
	Book        BookManagement `gorm:"foreignKey:BookID"`
}

// BorrowBook adalah model untuk peminjaman buku
type BorrowBook struct {
	gorm.Model
	BorrowID    uint      `json:"borrow_id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id"`
	BookID      uint      `json:"book_id"`
	DueDate     time.Time `json:"due_date"`
	Returned    bool      `json:"returned"`
	TotalBorrow int       `json:"total_borrow"`
	// User        User      `gorm:"foreignKey:UserID"`
}
