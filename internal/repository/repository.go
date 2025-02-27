package repository

import (
	"PGBridge/internal/dto/commentDTO"
	"PGBridge/internal/dto/userDTO"
	"PGBridge/internal/repository/postgres/comment"
	"PGBridge/internal/repository/postgres/user"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(dto userDTO.CreateUserRequestDTO) (int, error)
	GetUser(dto userDTO.GetUserRequestDTO) (interface{}, error)
	UpdateUser(dto userDTO.UpdateUserRequestDTO) (interface{}, error)
	DeleteUser(dto userDTO.DeleteUserRequestDTO) (interface{}, error)
}

type CommentRepository interface {
	CreateComment(dto commentDTO.CreateCommentRequestDTO) (int, error)
	GetComment(dto commentDTO.GetCommentRequestDTO) (interface{}, error)
	UpdateComment(dto commentDTO.UpdateCommentRequestDTO) (interface{}, error)
	DeleteComment(dto commentDTO.DeleteCommentRequestDTO) (interface{}, error)
}

type Repository struct {
	UserRepository
	CommentRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		user.NewUserPostgres(db),
		comment.NewCommentPostgres(db),
	}
}
