package controller

import (
	"PGBridge/internal/dto/commentDTO"
	"PGBridge/internal/dto/userDTO"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) DeleteUser(contentGin *gin.Context) {
	// Получаем UserID из URL
	userID, err := strconv.Atoi(contentGin.Param("id"))
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Создаем DTO
	dto := userDTO.DeleteUserRequestDTO{
		UserID: userID,
	}

	c.service.UserService.DeleteUser(dto)
}

func (c *Controller) DeleteComment(contentGin *gin.Context){
	// Получаем UserID из URL
	commentID, err := strconv.Atoi(contentGin.Param("id"))
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Создаем DTO
	dto := commentDTO.DeleteCommentRequestDTO{
		CommentID: commentID,
	}

	c.service.CommentService.DeleteComment(dto)
}