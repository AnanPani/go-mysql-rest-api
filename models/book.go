package models

// Book represents a book in the database.
type Book struct {
	ID     int    `json:"id"`
	BookName  string `json:"bookname"`
	Author string `json:"author"`
}