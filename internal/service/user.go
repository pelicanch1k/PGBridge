package service

import (
	"PGBridge/internal/dto/userDTO"
	"PGBridge/internal/repository"
	"encoding/json"
)

// UserService отвечает за бизнес-логику, связанную с пользователями
type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(service repository.UserRepository) *UserService {
	return &UserService{userRepository: service}
}

// GetUser возвращает данные пользователя по ID
func (s *UserService) GetUser(dto userDTO.GetUserRequestDTO) ([]byte, error) {
	result, err := s.userRepository.GetUser(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(dto userDTO.CreateUserRequestDTO) (int, error) {
	id, err := s.userRepository.CreateUser(dto)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateUser обновляет данные пользователя
func (s *UserService) UpdateUser(dto userDTO.UpdateUserRequestDTO) ([]byte, error) {
	result, err := s.userRepository.UpdateUser(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}

// DeleteUser удаляет пользователя по ID
func (s *UserService) DeleteUser(dto userDTO.DeleteUserRequestDTO) ([]byte, error) {
	result, err := s.userRepository.DeleteUser(dto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(result)
}
