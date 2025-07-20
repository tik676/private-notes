package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	/*err = godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}*/
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	for i := 0; i < 10; i++ {
		DB, err = sql.Open("postgres", dsn)
		if err == nil && DB.Ping() == nil {
			log.Println("Успешное подключение к базе данных")
			return
		}
		log.Println("Ожидание базы данных...")
		time.Sleep(2 * time.Second)
	}

	log.Println("Не удалось подключится к базе данных:", err)
}
