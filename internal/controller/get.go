package controller

import (
	"PGBridge/internal/dto/userDTO"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Получить данные по пользователю по ID
func (c *Controller) GetUserByID(contentGin *gin.Context) {
	// Получаем UserID из URL
	userID, err := strconv.Atoi(contentGin.Param("id"))
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Создаем DTO
	dto := userDTO.GetUserRequestDTO{
		UserID: userID,
	}

	userJson, err := c.service.UserService.GetUser(dto)
	if err != nil {
		contentGin.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentGin.JSON(http.StatusOK, userJson)
}

// Получить комментарий с ID
func (c *Controller) GetCommentByID(contentGin *gin.Context) {
	fmt.Println(contentGin.Request.URL.Path)
}

// Получить комментарий с ID от пользователя
func (c *Controller) GetCommentByUserAndID(contentGin *gin.Context) {
}

// Получить все комментарии пользователя по ID
func (c *Controller) GetListComment(contentGin *gin.Context) {
}
