package comment

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"

	"PGBridge/internal/dto/commentDTO"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

// CreateComment создает новый комментарий от пользователя
func (m *CommentPostgres) CreateComment(dto commentDTO.CreateCommentRequestDTO) (int, error) {
	paramsJSON, err := json.Marshal(dto.Params)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal params: %v", err)
	}

	var id int
	query := "SELECT * FROM test.user_comment_ins($1, $2)"
	err = m.db.QueryRow(query, dto.UserID, paramsJSON).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create commentDTO: %v", err)
	}
	return id, nil
}

// GetComment возвращает комментарий по ID или все комментарии, если ID = 0
func (m *CommentPostgres) GetComment(dto commentDTO.GetCommentRequestDTO) (interface{}, error) {
	var result interface{}
	query := "SELECT * FROM test.comment_get($1)"
	err := m.db.QueryRow(query, dto.CommentID).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get commentDTO: %v", err)
	}
	return result, nil
}

// UpdateComment обновляет текст комментария
func (m *CommentPostgres) UpdateComment(dto commentDTO.UpdateCommentRequestDTO) (interface{}, error) {
	var result interface{}
	paramsJSON, err := json.Marshal(dto.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}

	query := "SELECT * FROM test.comment_upd($1, $2)"
	err = m.db.QueryRow(query, dto.CommentID, paramsJSON).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to update commentDTO: %v", err)
	}
	return result, nil
}

// DeleteComment удаляет комментарий по ID
func (m *CommentPostgres) DeleteComment(dto commentDTO.DeleteCommentRequestDTO) (interface{}, error) {
	var result interface{}
	query := "SELECT * FROM test.comment_del($1)"
	err := m.db.QueryRow(query, dto.CommentID).Scan(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to delete commentDTO: %v", err)
	}
	return result, nil
}
