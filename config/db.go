package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver

	"go_crud_auth/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// kita buat koneksi dgn db posgres
func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Kita buka koneksi ke db
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Konek ke Db!")
	// return the connection
	return db
}

type NullString struct {
	sql.NullString
}

var DB *gorm.DB

func InitDB() {

	// load .env file
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}

	var err error
	// DB, err = gorm.Open("mysql", "root:@/egommerce?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic("Gagal konek ke db")
	}
	// defer DB.Close()

	Migrate()
}

func Migrate() {

	// migrasi
	DB.AutoMigrate(&models.Student{})

	data := models.Student{}
	if DB.Where("student_id = ?", 1).Find(&data).RecordNotFound() {
		seederUser()
	}
}

func seederUser() {
	data := models.Student{
		Student_id:       1,
		Student_name:     "Fauzi",
		Student_age:      17,
		Student_address:  "Semarang",
		Student_phone_no: "082212341234",
	}
	DB.Create(&data)
}
