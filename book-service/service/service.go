package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Author struct {
	AuthorID uint   `json:"author_id"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

type Category struct {
	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleCode int    `json:"role_code"`
	Role     string `json:"role"`
}

// Perbaiki tipe return di sini menjadi *User, bukan *Author
func GetUserByID(userID uint) (*User, error) {
	url := fmt.Sprintf("http://user-service-url/api/users/%d", userID) 
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user: status %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAuthorByID(authorID uint) (*Author, error) {
	url := fmt.Sprintf("http://localhost:3002/authors%d", authorID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get author: status %d", resp.StatusCode)
	}

	var author Author
	if err := json.NewDecoder(resp.Body).Decode(&author); err != nil {
		return nil, err
	}

	return &author, nil
}

func GetCategoryByID(categoryID uint) (*Category, error) {
	url := fmt.Sprintf("http://localhost:3003/categories%d", categoryID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get category: status %d", resp.StatusCode)
	}

	var category Category
	if err := json.NewDecoder(resp.Body).Decode(&category); err != nil {
		return nil, err
	}

	return &category, nil
}
