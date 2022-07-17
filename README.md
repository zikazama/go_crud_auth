# Example API Gin Go With Auth

Route

- POST /login (json data with username and password attribute)
- GET /student
- GET /student/:studentId
- POST /student (json data with student_id, student_name, student_age, student_address and student_phone_no attribute)
- PUT /student/:studentId (json data with student_name, student_age, student_address and student_phone_no attribute)
- DELETE /student/:studentId

# How to run

1. Copy .example.env to .env
2. Setup db config in .env
3. go run main.go

# Credential

- Username : admin
- Password : Password123!
