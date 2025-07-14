package config

import (
	"os"
)

// JWTSecretKey คือคีย์ลับที่ใช้สำหรับเซ็น JWTs
// ในแอปพลิเคชันจริง ควรโหลดค่าจาก environment variables
// หรือระบบจัดการการตั้งค่าที่ปลอดภัย
var JWTSecretKey = []byte("super-secret-jwt-key") // เปลี่ยนเป็นคีย์ที่ซับซ้อนและปลอดภัย!

func LoadConfig() {
	// ตัวอย่าง: โหลดค่าจาก environment variable หากมีอยู่
	if os.Getenv("JWT_SECRET_KEY") != "" {
		JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	}
}