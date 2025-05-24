package controllers

import (
	"bakeryapp/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateDocument godoc
// @Summary Создать документ
// @Tags Documents
// @Accept json
// @Produce json
// @Param document body models.Document true "Новый документ"
// @Success 201 {object} models.Document
// @Failure 400,500 {object} map[string]string
// @Router /documents [post]
// CreateDocument godoc
// @Summary Создать документ
// @Tags Documents
// @Accept json
// @Produce json
// @Param document body models.Document true "Новый документ"
// @Success 201 {object} models.Document
// @Failure 400,500 {object} map[string]string
// @Router /documents [post]
func CreateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var document models.Document
		if err := c.ShouldBindJSON(&document); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		document.ID = 0 // Не задаём ID, база сама сгенерирует
		if document.CreateDate.IsZero() {
			document.CreateDate = time.Now()
		}

		if err := db.Create(&document).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать документ"})
			return
		}
		c.JSON(http.StatusCreated, document)
	}
}

// GetAllDocuments godoc
// @Summary Получить все документы
// @Tags Documents
// @Produce json
// @Success 200 {array} models.Document
// @Failure 500 {object} map[string]string
// @Router /documents [get]
func GetAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var docs []models.Document
		if err := db.Find(&docs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить документы"})
			return
		}
		c.JSON(http.StatusOK, docs)
	}
}

// GetDocumentByID godoc
// @Summary Получить документ по ID
// @Tags Documents
// @Produce json
// @Param id path int true "ID документа"
// @Success 200 {object} models.Document
// @Failure 404,500 {object} map[string]string
// @Router /documents/{id} [get]
func GetDocumentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var doc models.Document
		if err := db.First(&doc, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "документ не найден"})
			return
		}
		c.JSON(http.StatusOK, doc)
	}
}

// UpdateDocument godoc
// @Summary Обновить документ по ID
// @Tags Documents
// @Accept json
// @Produce json
// @Param id path int true "ID документа"
// @Param document body models.Document true "Обновлённый документ"
// @Success 200 {object} models.Document
// @Failure 400,404,500 {object} map[string]string
// @Router /documents/{id} [put]
func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
			return
		}

		var document models.Document
		if err := db.First(&document, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "документ не найден"})
			return
		}

		var input models.Document
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Не обновляем поле ID
		input.ID = document.ID

		if err := db.Model(&document).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить документ"})
			return
		}

		c.JSON(http.StatusOK, document)
	}
}

// DeleteDocument godoc
// @Summary Удалить документ
// @Tags Documents
// @Produce json
// @Param id path int true "ID документа"
// @Success 200 {object} map[string]string
// @Failure 404,500 {object} map[string]string
// @Router /documents/{id} [delete]
func DeleteDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Document{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить документ"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "документ удален"})
	}
}
