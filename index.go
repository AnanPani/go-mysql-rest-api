package main

import (
	"fmt"
	"go-mysql-rest-api/config" // NEW: Import config
	"go-mysql-rest-api/database"
	"go-mysql-rest-api/handlers"
	"go-mysql-rest-api/middleware" // NEW: Import middleware
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// NEW: Load configurations
	config.LoadConfig()

	// 1. เชื่อมต่อฐานข้อมูล
	database.Connect()
	defer database.DB.Close() // ตรวจสอบให้แน่ใจว่าปิดการเชื่อมต่อเมื่อโปรแกรมจบ

	// 2. สร้าง Router
	router := mux.NewRouter()

	// 3. กำหนด API Endpoints

	// Public Routes (ไม่ต้อง Authenticate)
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Protected Routes (ต้อง Authenticate)
	// เราจะใช้ Subrouter เพื่อใช้ Middleware กับกลุ่มของ Routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware) // ใช้ AuthMiddleware กับทุก Route ใน Subrouter นี้

	apiRouter.HandleFunc("/book", handlers.GetBooks).Methods("GET")
	apiRouter.HandleFunc("/book/{id}", handlers.GetBook).Methods("GET")
	apiRouter.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	apiRouter.HandleFunc("/book/{id}", handlers.UpdateBook).Methods("PUT")
	apiRouter.HandleFunc("/book/{id}", handlers.DeleteBook).Methods("DELETE")

	// 4. รัน Server
	port := ":8080"
	fmt.Printf("เซิร์ฟเวอร์กำลังรันและรอรับคำขอบนพอร์ต %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}


	// ตัวอย่างการทดสอบ API
	// คุณสามารถใช้ Postman หรือ curl เพื่อทดสอบ API ที่สร้างขึ้นได้
	
	// ตัวอย่างลงทะเบียนผู้ใช้ใหม่:
	// Method: POST

	// URL: http://localhost:8080/register

	// Headers: Content-Type: application/json

	// Body (raw JSON):

	// JSON

	// {
	//     "username": "testuser",
	//     "password": "password123"
	// }
	// Expected Response: 201 Created พร้อม { "message": "User registered successfully" }


	// ตัวอย่างเข้าสู่ระบบผู้ใช้:
	// 	Method: POST

	// URL: http://localhost:8080/login

	// Headers: Content-Type: application/json

	// Body (raw JSON):

	// JSON

	// {
	//     "username": "testuser",
	//     "password": "password123"
	// }
	// จะได้ Token
	// 	{
	//     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNjg1NTEyMzQ1LCJpYXQiOjE2ODU1MTE3NDUsIm5iZiI6MTY4NTUxMTc0NX0...."
	// }


	// ตัวอย่างดึงข้อมูลหนังสือ:
	// 	ก่อนที่จะดึงข้อมูลหนังสือ คุณต้องเข้าสู่ระบบเพื่อรับ JWT Token
	// 	Method: GET

	// URL: http://localhost:8080/api/books

	// Headers:

	// Content-Type: application/json

	// Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)

	// Expected Response: 200 OK พร้อม JSON array ของหนังสือ


	// ตัวอย่างการทดสอบ API
	// คุณสามารถใช้ Postman หรือ curl เพื่อทดสอบ API ที่สร้างขึ้นได้

	// 	สร้างหนังสือใหม่:
	// POST http://localhost:8080/books
	// Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	// Body (raw JSON):

	// JSON

	// {
	//     "bookname": "The Hitchhiker's Guide to the Galaxy",
	//     "author": "Douglas Adams",
	// }


	// อัพเดทหนังสือ (ตัวอย่าง, ID 1):
	// PUT http://localhost:8080/books/1
	// Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	// Body (raw JSON):

	// JSON

	// {
	//     "title": "The Lord of the Rings (Updated)",
	//     "author": "J.R.R. Tolkien",
	//     "isbn": "978-0618052163"
	// }

	// ลบหนังสือ (ตัวอย่าง, ID 1):
	// Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	// DELETE http://localhost:8080/books/1
