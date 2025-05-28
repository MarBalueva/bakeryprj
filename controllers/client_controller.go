package controllers

import (
	"bakeryapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllClients godoc
// @Summary Получить всех клиентов
// @Tags Client
// @Accept json
// @Produce json
// @Success 200 {array} models.Client
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /clients [get]
// @Security BearerAuth
func GetAllClients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var clients []models.Client
		if err := db.Where("isdeleted = ?", false).Find(&clients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить клиентов"})
			return
		}
		c.JSON(http.StatusOK, clients)
	}
}

// GetClientByID godoc
// @Summary Получить клиента по ID
// @Tags Client
// @Produce json
// @Param id path int true "ID клиента"
// @Success 200 {object} models.Client
// @Failure 404 {object} map[string]string "клиент не найден"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /clients/{id} [get]
// @Security BearerAuth
func GetClientByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var client models.Client
		if err := db.Where("id = ? AND isdeleted = ?", id, false).First(&client).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "клиент не найден"})
			return
		}
		c.JSON(http.StatusOK, client)
	}
}

// CreateClient godoc
// @Summary Создать нового клиента
// @Tags Client
// @Accept json
// @Produce json
// @Param client body models.Client true "Данные клиента"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]string "ошибка валидации"
// @Failure 500 {object} map[string]string "не удалось создать клиента"
// @Router /clients [post]
// @Security BearerAuth
func CreateClient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var client models.Client
		if err := c.ShouldBindJSON(&client); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		client.ID = 0
		client.IsDeleted = false
		if err := db.Create(&client).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать клиента"})
			return
		}
		c.JSON(http.StatusCreated, client)
	}
}

// UpdateClient godoc
// @Summary Обновить данные клиента
// @Tags Client
// @Accept json
// @Produce json
// @Param id path int true "ID клиента"
// @Param client body models.Client true "Обновлённые данные клиента"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]string "ошибка валидации"
// @Failure 404 {object} map[string]string "клиент не найден"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /clients/{id} [put]
// @Security BearerAuth
func UpdateClient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var client models.Client
		id := c.Param("id")

		if err := db.Where("id = ? AND isdeleted = ?", id, false).First(&client).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "клиент не найден"})
			return
		}

		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		delete(input, "id")
		delete(input, "isdeleted")

		if err := db.Model(&client).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить клиента"})
			return
		}

		c.JSON(http.StatusOK, client)
	}
}

// DeleteClient godoc
// @Summary Удалить клиента
// @Tags Client
// @Produce json
// @Param id path int true "ID клиента"
// @Success 200 {object} map[string]string "клиент удалён"
// @Failure 404 {object} map[string]string "клиент не найден"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /clients/{id} [delete]
// @Security BearerAuth
func DeleteClient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var client models.Client
		if err := db.Where("id = ? AND isdeleted = ?", id, false).First(&client).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "клиент не найден"})
			return
		}

		if !client.IsDeleted {
			c.JSON(http.StatusBadRequest, gin.H{"error": "клиент уже удалён"})
			return
		}

		client.IsDeleted = true
		if err := db.Save(&client).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить клиента"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "клиент удален"})
	}
}
