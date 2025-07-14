package middleware

import (
	"context"
	"fmt"
	"go-mysql-rest-api/config"
	"go-mysql-rest-api/handlers"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware เป็นมิดเดิลแวร์สำหรับป้องกันเส้นทาง (routes) ด้วยการยืนยันตัวตนแบบ JWT
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// รูปแบบที่คาดหวัง: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims := &handlers.Claims{} // ใช้ struct Claims จากแพ็กเกจ handlers

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.JWTSecretKey, nil // ใช้คีย์ลับจากไฟล์ config
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// หากโทเค็นถูกต้อง ให้ส่งคำขอไปยัง handler ถัดไป
		// หรืออาจเพิ่มข้อมูลผู้ใช้ลงใน context ของคำขอก็ได้
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUsernameFromContext เป็นฟังก์ชันช่วยสำหรับดึงชื่อผู้ใช้จาก context
func GetUsernameFromContext(r *http.Request) (string, bool) {
    username, ok := r.Context().Value("username").(string)
    return username, ok
}