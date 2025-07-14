ใช้คำสั่ง
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

รันและทดสอบ API
รัน Go Application:
เปิด Terminal/Command Prompt ในโฟลเดอร์ go-mysql-rest-api แล้วรัน:

go run main.go

	 ตัวอย่างการทดสอบ API
	 คุณสามารถใช้ Postman หรือ curl เพื่อทดสอบ API ที่สร้างขึ้นได้
	
	 ตัวอย่างลงทะเบียนผู้ใช้ใหม่:
	 Method: POST

	 URL: http://localhost:8080/register

	 Headers: Content-Type: application/json

	 Body (raw JSON):

	 JSON

	 {
	     "username": "testuser",
	     "password": "password123"
	 }
	 Expected Response: 201 Created พร้อม { "message": "User registered successfully" }


	 ตัวอย่างเข้าสู่ระบบผู้ใช้:
	 	Method: POST

	 URL: http://localhost:8080/login

	 Headers: Content-Type: application/json

	 Body (raw JSON):

	 JSON

	 {
	     "username": "testuser",
	     "password": "password123"
	 }
	 จะได้ Token
	 	{
	     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNjg1NTEyMzQ1LCJpYXQiOjE2ODU1MTE3NDUsIm5iZiI6MTY4NTUxMTc0NX0...."
	 }


	  ตัวอย่างดึงข้อมูลหนังสือ:
	 	ก่อนที่จะดึงข้อมูลหนังสือ คุณต้องเข้าสู่ระบบเพื่อรับ JWT Token
	 	Method: GET

	 URL: http://localhost:8080/api/books

	 Headers:

	 Content-Type: application/json

	 Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)

	 Expected Response: 200 OK พร้อม JSON array ของหนังสือ


	 ตัวอย่างการทดสอบ API
	 คุณสามารถใช้ Postman หรือ curl เพื่อทดสอบ API ที่สร้างขึ้นได้

	 	สร้างหนังสือใหม่:
	 POST http://localhost:8080/books
	 Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	 Body (raw JSON):

	 JSON

	 {
	     "bookname": "The Hitchhiker's Guide to the Galaxy",
	     "author": "Douglas Adams",
	 }


	 อัพเดทหนังสือ (ตัวอย่าง, ID 1):
	 PUT http://localhost:8080/books/1
	 Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	 Body (raw JSON):

	 JSON

	 {
	    "title": "The Lord of the Rings (Updated)",
	    "author": "J.R.R. Tolkien",
	     "isbn": "978-0618052163"
	}

	ลบหนังสือ (ตัวอย่าง, ID 1):
	Headers: Content-Type: application/json // Authorization: Bearer <your_jwt_token_from_login_step> (แทนที่ <your_jwt_token_from_login_step> ด้วย JWT ที่คุณได้รับ)
	DELETE http://localhost:8080/books/1
