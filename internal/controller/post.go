package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"PGBridge/internal/dto/commentDTO"
	"PGBridge/internal/dto/userDTO"
)

func (c *Controller) CreateComment(contentGin *gin.Context) {
	// Получаем UserID из URL
	userID, err := strconv.Atoi(contentGin.Param("id"))
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Маппим тело запроса в Params
	var params map[string]interface{}
	if err := contentGin.BindJSON(&params); err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Создаем DTO
	dto := commentDTO.CreateCommentRequestDTO{
		UserID: userID,
		Params: params,
	}

	id, err := c.service.CommentService.CreateComment(dto)
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логика обработки
	contentGin.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"params":  dto.Params,
	})
}

func (c *Controller) CreateUser(contentGin *gin.Context) {
	// Маппим тело запроса в Params
	var params map[string]interface{}
	if err := contentGin.BindJSON(&params); err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Создаем DTO
	dto := userDTO.CreateUserRequestDTO{
		Params: params,
	}

	id, err := c.service.UserService.CreateUser(dto)
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логика обработки
	contentGin.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"params":  dto.Params,
	})
}
