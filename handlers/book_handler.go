package handlers

import (
	"encoding/json"
	"fmt"
	"go-mysql-rest-api/database"
	"go-mysql-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"database/sql"
)

// GetBooks จัดการคำขอ GET สำหรับหนังสือทั้งหมด
func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, bookname, author FROM book")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.BookName, &book.Author); err != nil {
			log.Printf("เกิดปัญหา : %v", err)
			continue // ข้ามแถวที่มีปัญหา
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook จัดการคำขอ GET สำหรับหนังสือหนึ่งเล่มโดยใช้รหัสประจำตัว (ID)
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	row := database.DB.QueryRow("SELECT id, bookname, author FROM book WHERE id = ?", id)
	err = row.Scan(&book.ID, &book.BookName, &book.Author)
	if err == sql.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// CreateBook handles POST requests to create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("INSERT INTO books (bookname, author) VALUES (?, ?)", book.BookName, book.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("แจ้งเตือน: ไม่สามารถรับค่า ID ที่แทรกล่าสุดได้: %v", err)
	}
	book.ID = int(id) // ตั้งค่า ID ที่สร้างขึ้นโดย DB ให้กับ object

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // ส่งสถานะ 201 Created
	json.NewEncoder(w).Encode(book)
}

// UpdateBook จัดการคำขอ PUT เพื่ออัปเดตข้อมูลหนังสือที่มีอยู่
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ตรวจสอบให้แน่ใจว่า ID ที่ส่งมาใน URL ตรงกับ ID ใน body (ถ้ามี)
	// แต่สำหรับ PUT, ID ใน URL เป็นตัวหลัก
	if book.ID != 0 && book.ID != id {
		log.Printf("แจ้งเตือน: ID ของ URL (%d) ไม่ตรงกับ body (%d). ใช้ URL ID.", id, book.ID)
	}

	_, err = database.DB.Exec("UPDATE book SET bookname = ?, author = ? WHERE id = ?", book.BookName, book.Author, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ดึงข้อมูลล่าสุดกลับไปให้ client
	row := database.DB.QueryRow("SELECT id, bookname, author FROM book WHERE id = ?", id)
	err = row.Scan(&book.ID, &book.BookName, &book.Author)
	if err == sql.ErrNoRows {
		http.Error(w, "ไม่พบหนังสือที่จะอัพเดท)", http.StatusInternalServerError)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// DeleteBook จัดการคำขอ DELETE เพื่อทำการลบหนังสือ
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ไม่พบ ID หนังสือ", http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("แจ้งเตือน: ไม่สามารถรับจำนวนแถวที่ได้รับผลกระทบได้: %v", err)
	}

	if rowsAffected == 0 {
		http.Error(w, "ไม่พบหนังสือ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // ส่งสถานะ 204 No Content สำหรับการลบสำเร็จ
	fmt.Fprintf(w, "Book with ID %d deleted successfully", id) // หรือจะส่งแค่ 204 ก็ได้
}