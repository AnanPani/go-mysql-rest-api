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
}