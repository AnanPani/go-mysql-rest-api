package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// แทนที่ด้วยข้อมูลการเชื่อมต่อ MySQL ของคุณ
	// เช่น "user:password@tcp(127.0.0.1:3306)/dbname"
	dbUser := "root"      // ผู้ใช้ MySQL
	dbPassword := "root" // รหัสผ่าน MySQL ของคุณ
	dbHost := "localhost" // หรือ localhost
	dbPort := "3306"
	dbName := "go_api_db" // ชื่อฐานข้อมูลที่เราสร้างเมื่อกี้

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("ไม่สามารถเปิดฐานข้อมูล : %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("การเชื่อมต่อฐานข้อมูลผิดพลาด: %v", err)
	}

	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จ!")
}