package model

type Book struct {
	ID          int    `json:"book_id"`
	Title       string `json:"book_title"`
	Author      string `json:"book_author"`
	Description string `json:"description"`
}
