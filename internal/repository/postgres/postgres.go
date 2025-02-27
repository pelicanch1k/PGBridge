package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	_ "github.com/lib/pq"

	"PGBridge/internal/config"
)

// InitDB инициализирует соединение с базой данных и выполняет sample.sql
func NewPostgresRepository() (*sqlx.DB, error) {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Выполнение sample.sql для инициализации базы данных
	file, err := os.ReadFile("sample.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read sample.sql: %v", err)
	}

	_, err = db.Exec(string(file))
	if err != nil {
		return nil, fmt.Errorf("failed to execute sample.sql: %v", err)
	}

	log.Println("Database initialized successfully")

	return db, nil
}
