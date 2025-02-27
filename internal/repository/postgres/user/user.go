package user

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"

	"PGBridge/internal/dto/userDTO"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// CreateUser создает нового пользователя
func (m *UserPostgres) CreateUser(dto userDTO.CreateUserRequestDTO) (int, error) {
	// Преобразуем Params в JSON
	paramsJSON, err := json.Marshal(dto.Params)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal params: %v", err)
	}

	// Выполняем SQL-запрос и сканируем результат (ID)
	var userID int
	query := "SELECT id FROM test.user_ins($1)"
	err = m.db.QueryRow(query, paramsJSON).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}

	// Возвращаем ID созданного пользователя
	return userID, nil
}

// GetUser возвращает данные пользователя по ID или всех пользователей, если ID = 0
func (m *UserPostgres) GetUser(dto userDTO.GetUserRequestDTO) (interface{}, error) {
	var result interface{}
	query := "SELECT * FROM test.user_get($1)"
	err := m.db.QueryRow(query, dto.UserID).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get userDTO: %v", err)
	}
	return result, nil
}

// UpdateUser обновляет данные пользователя
func (m *UserPostgres) UpdateUser(dto userDTO.UpdateUserRequestDTO) (interface{}, error) {
	var result interface{}
	paramsJSON, err := json.Marshal(dto.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}

	query := "SELECT * FROM test.user_upd($1, $2)"
	err = m.db.QueryRow(query, dto.UserID, paramsJSON).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to update userDTO: %v", err)
	}
	return result, nil
}

// DeleteUser удаляет пользователя по ID
func (m *UserPostgres) DeleteUser(dto userDTO.DeleteUserRequestDTO) (interface{}, error) {
	var result interface{}
	query := "SELECT * FROM test.user_del($1)"
	err := m.db.QueryRow(query, dto.UserID).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to delete userDTO: %v", err)
	}
	return result, nil
}
