package handlers

import (
	"encoding/json"
	"go-mysql-rest-api/database"
	"go-mysql-rest-api/models"
	"net/http"
	"time"

	"go-mysql-rest-api/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"database/sql"
)

// claims แทนข้อมูลอ้างสิทธิ์ (claims) ของ JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// RegisterUser จัดการคำขอ POST เพื่อสมัครสมาชิกผู้ใช้ใหม่
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// แฮชรหัสผ่าน
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// เพิ่มผู้ใช้ลงในฐานข้อมูล
	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		http.Error(w, "Username already exists or database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUser จัดการคำขอ POST เพื่อเข้าสู่ระบบผู้ใช้และออก JWT
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ดึงข้อมูลผู้ใช้จากฐานข้อมูล
	var user models.User
	row := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", creds.Username)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// เปรียบเทียบรหัสผ่านที่ผู้ใช้กรอกกับรหัสผ่านที่ถูกแฮชแล้ว
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	expirationTime := time.Now().Add(5 * time.Minute) // โทเค็นมีอายุใช้งาน 5 นาที
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTSecretKey) // ใช้คีย์ลับจากการตั้งค่า (config)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}