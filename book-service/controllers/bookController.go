package controllers

import (
	"github.com/FARRAS-DARKUNO/library-management/book-service/config"
	"github.com/FARRAS-DARKUNO/library-management/book-service/models"
	"github.com/FARRAS-DARKUNO/library-management/book-service/service"
	"github.com/gofiber/fiber/v2"
)

// CreateBook menangani pembuatan buku baru
func CreateBook(c *fiber.Ctx) error {
    var book models.BookManagement
    if err := c.BodyParser(&book); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
    }

    // Validasi apakah author ada
    if _, err := service.GetAuthorByID(book.AuthorID); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid Author ID")
    }

    // Validasi apakah category ada
    if _, err := service.GetCategoryByID(book.CategoryID); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid Category ID")
    }

    if err := config.DB.Create(&book).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to create book")
    }

    return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBooks menangani permintaan untuk mengambil semua buku
func GetBooks(c *fiber.Ctx) error {
	var books []models.BookManagement
	if err := config.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve books")
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

// GetBookByID menangani permintaan untuk mengambil buku berdasarkan ID
func GetBookByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.BookManagement

    // Cek apakah buku ditemukan
    if err := config.DB.First(&book, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).SendString("Book not found")
    }

    // Mengambil data Author dari service
    author, err := service.GetAuthorByID(book.AuthorID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve author data")
    }

    // Mengambil data Category dari service
    category, err := service.GetCategoryByID(book.CategoryID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve category data")
    }

    // Menyusun response dengan data tambahan
    response := map[string]interface{}{
        "book": book,
        "author": author,
        "category": category,
    }

    return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateBook menangani permintaan untuk memperbarui buku
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.BookManagement
	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Book not found")
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update book")
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

// DeleteBook menangani permintaan untuk menghapus buku
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.BookManagement{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete book")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// CreateBookStock menangani pembuatan stok buku baru
func CreateBookStock(c *fiber.Ctx) error {
	var bookStock models.BookStock
	if err := c.BodyParser(&bookStock); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	if err := config.DB.Create(&bookStock).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create book stock")
	}
	return c.Status(fiber.StatusCreated).JSON(bookStock)
}

// CreateBorrowBook menangani pembuatan peminjaman buku
func CreateBorrowBook(c *fiber.Ctx) error {
	var borrow models.BorrowBook
	if err := c.BodyParser(&borrow); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	if err := config.DB.Create(&borrow).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create borrow record")
	}
	return c.Status(fiber.StatusCreated).JSON(borrow)
}


// ReturnBook menangani pengembalian buku
func ReturnBook(c *fiber.Ctx) error {
	// Mendapatkan ID peminjaman dari parameter URL
	borrowID := c.Params("id")
	var borrow models.BorrowBook

	// Cari record peminjaman berdasarkan borrowID
	if err := config.DB.First(&borrow, borrowID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Borrow record not found")
	}

	// Periksa apakah buku sudah dikembalikan
	if borrow.Returned {
		return c.Status(fiber.StatusConflict).SendString("Book already returned")
	}

	// Update status pengembalian menjadi true
	borrow.Returned = true

	// Mengurangi stok buku jika sudah dikembalikan
	var bookStock models.BookStock
	if err := config.DB.First(&bookStock, "book_id = ?", borrow.BookID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Book stock not found")
	}
	bookStock.Stock++

	// Simpan perubahan pada peminjaman dan stok buku
	if err := config.DB.Save(&borrow).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update borrow record")
	}
	if err := config.DB.Save(&bookStock).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update book stock")
	}

	// Kirimkan respons sukses
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book returned successfully",
		"borrow":  borrow,
		"bookStock": bookStock,
	})
}
