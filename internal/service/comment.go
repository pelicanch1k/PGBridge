package service

import (
	"PGBridge/internal/dto/commentDTO"
	"PGBridge/internal/repository"
	"encoding/json"
)

// CommentService отвечает за бизнес-логику, связанную с комментариями
type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(service repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: service}
}

// GetComment возвращает комментарий по ID или все комментарии, если ID = 0
func (s *CommentService) GetComment(dto commentDTO.GetCommentRequestDTO) ([]byte, error) {
	result, err := s.commentRepository.GetComment(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// UpdateComment обновляет текст комментария
func (s *CommentService) UpdateComment(dto commentDTO.UpdateCommentRequestDTO) ([]byte, error) {
	result, err := s.commentRepository.UpdateComment(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// DeleteComment удаляет комментарий по ID
func (s *CommentService) DeleteComment(dto commentDTO.DeleteCommentRequestDTO) ([]byte, error) {
	result, err := s.commentRepository.DeleteComment(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// CreateComment создает новый комментарий от пользователя
func (s *CommentService) CreateComment(dto commentDTO.CreateCommentRequestDTO) (int, error) {
	id, err := s.commentRepository.CreateComment(dto)
	if err != nil {
		return 0, err
	}
	
	return id, nil
}
