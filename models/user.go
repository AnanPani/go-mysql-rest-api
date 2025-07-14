package models

// User represents a user in the database.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // omitempty จะไม่รวมฟิลด์นี้เมื่อ encode เป็น JSON ถ้าค่าว่าง
}