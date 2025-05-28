package controllers

import (
	"bakeryapp/models"
	"encoding/json"
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
// @Security BearerAuth
func CreateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var raw map[string]interface{}
		if err := c.BindJSON(&raw); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат JSON"})
			return
		}

		if val, ok := raw["createDate"]; ok && val == "string" {
			delete(raw, "createDate")
		}
		if val, ok := raw["endDate"]; ok && val == "string" {
			delete(raw, "endDate")
		}

		cleaned, err := json.Marshal(raw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка обработки данных"})
			return
		}

		var document models.Document
		if err := json.Unmarshal(cleaned, &document); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось распознать документ"})
			return
		}

		if document.StartDate.IsZero() {
			now := time.Now()
			document.StartDate = now
		}

		if document.ProductID != nil && *document.ProductID <= 0 {
			document.ProductID = nil
		}
		if document.OrderID != nil && *document.OrderID <= 0 {
			document.OrderID = nil
		}

		document.ID = 0

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
// @Security BearerAuth
func GetAllDocuments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var docs []models.Document
		if err := db.Where("isdeleted = ?", false).Find(&docs).Error; err != nil {
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
// @Security BearerAuth
func GetDocumentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var doc models.Document
		if err := db.Where("id = ? AND isdeleted = ?", id, false).First(&doc).Error; err != nil {
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
// @Security BearerAuth
func UpdateDocument(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
			return
		}

		var existing models.Document
		if err := db.First(&existing, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "документ не найден"})
			return
		}

		var raw map[string]interface{}
		if err := c.BindJSON(&raw); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат JSON"})
			return
		}

		if val, ok := raw["createDate"]; ok && val == "string" {
			delete(raw, "createDate")
		}
		if val, ok := raw["endDate"]; ok && val == "string" {
			delete(raw, "endDate")
		}

		cleaned, err := json.Marshal(raw)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка обработки данных"})
			return
		}

		var input models.Document
		if err := json.Unmarshal(cleaned, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось распознать документ"})
			return
		}

		if input.StartDate.IsZero() {
			now := time.Now()
			input.StartDate = now
		}

		if input.ProductID != nil && *input.ProductID <= 0 {
			input.ProductID = nil
		}
		if input.OrderID != nil && *input.OrderID <= 0 {
			input.OrderID = nil
		}

		input.ID = existing.ID

		if err := db.Model(&existing).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить документ"})
			return
		}

		c.JSON(http.StatusOK, existing)
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
// @Security BearerAuth
func DeleteDocument(db *gorm.DB) gin.HandlerFunc {
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

		if !document.Status {
			c.JSON(http.StatusBadRequest, gin.H{"error": "документ уже удалён"})
			return
		}

		if err := db.Model(&document).Update("status", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить документ"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "документ удалён"})
	}
}

// SubmitDocument godoc
// @Summary Подписать документ
// @Description Устанавливает поле issubmit = true для документа, если он существует и не удалён (status = true)
// @Tags Documents
// @Produce json
// @Param id path int true "ID документа"
// @Success 200 {object} map[string]string "документ подписан"
// @Failure 400 {object} map[string]string "неверный ID или попытка подписать удалённый документ"
// @Failure 404 {object} map[string]string "документ не найден"
// @Failure 500 {object} map[string]string "ошибка сервера при подписании документа"
// @Router /documents/{id}/submit [put]
// @Security BearerAuth
func SubmitDocument(db *gorm.DB) gin.HandlerFunc {
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

		if !document.Status {
			c.JSON(http.StatusBadRequest, gin.H{"error": "удалённый документ нельзя подписать"})
			return
		}

		if err := db.Model(&document).Update("issubmit", true).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось подписать документ"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "документ подписан"})
	}
}
