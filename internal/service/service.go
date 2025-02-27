package service

import "PGBridge/internal/repository"

type Service struct {
	CommentService *CommentService
	UserService    *UserService
}

func NewService(db *repository.Repository) *Service {
	return &Service{
		CommentService: NewCommentService(db),
		UserService:    NewUserService(db),
	}
}
