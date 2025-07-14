package main

import (
	"fmt"
	"go-mysql-rest-api/database"
	"go-mysql-rest-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 1. เชื่อมต่อฐานข้อมูล
	database.Connect()
	defer database.DB.Close() // ตรวจสอบให้แน่ใจว่าปิดการเชื่อมต่อเมื่อโปรแกรมจบ

	// 2. สร้าง Router
	router := mux.NewRouter()

	// 3. กำหนด API Endpoints
	router.HandleFunc("/book", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handlers.DeleteBook).Methods("DELETE")

	// 4. รัน Server
	port := ":8080"
	fmt.Printf("เซิร์ฟเวอร์กำลังรอรับคำขอบนพอร์ต %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))



// ตัวอย่างการทดสอบ API
// คุณสามารถใช้ Postman หรือ curl เพื่อทดสอบ API ที่สร้างขึ้นได้

// 	สร้างหนังสือใหม่:
// POST http://localhost:8080/books
// Headers: Content-Type: application/json
// Body (raw JSON):

// JSON

// {
//     "bookname": "The Hitchhiker's Guide to the Galaxy",
//     "author": "Douglas Adams",
// }


}

// อัพเดทหนังสือ (ตัวอย่าง, ID 1):
// PUT http://localhost:8080/books/1
// Headers: Content-Type: application/json
// Body (raw JSON):

// JSON

// {
//     "title": "The Lord of the Rings (Updated)",
//     "author": "J.R.R. Tolkien",
//     "isbn": "978-0618052163"
// }

// ลบหนังสือ (ตัวอย่าง, ID 1):
// DELETE http://localhost:8080/books/1