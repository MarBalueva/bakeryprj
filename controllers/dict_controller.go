package controllers

import (
	"bakeryapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getDictModel(dictType string) interface{} {
	switch dictType {
	case "position":
		return &models.Position{}
	case "accessGroup":
		return &models.AccessGroup{}
	case "subcategoryProduct":
		return &models.SubcategoryProduct{}
	case "statuses":
		return &models.Status{}
	case "paymentType":
		return &models.PaymentType{}
	default:
		return nil
	}
}

// GetAllDictItems godoc
// @Summary Получить элементы справочника
// @Description Возвращает все не удалённые элементы справочника (isDeleted = false)
// @Tags Dictionary
// @Param dict path string true "Тип справочника (например: position, accessGroup, subcategoryProduct, statuses, paymentType)"
// @Produce json
// @Success 200 {array} models.DictionaryItem
// @Failure 400 {object} map[string]string "неверный тип справочника"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /dict/{dict} [get]
// @Security BearerAuth
func GetAllDictItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dictType := c.Param("dict")
		model := getDictModel(dictType)
		if model == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный тип справочника"})
			return
		}

		slice := []map[string]interface{}{}
		if err := db.Model(model).Where("is_deleted = false").Find(&slice).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, slice)
	}
}

// CreateDictItem godoc
// @Summary Создать элемент справочника
// @Description Добавляет новый элемент в указанный справочник
// @Tags Dictionary
// @Param dict path string true "Тип справочника"
// @Accept json
// @Produce json
// @Param item body models.DictionaryItem true "Элемент справочника"
// @Success 201 {object} models.DictionaryItem
// @Failure 400 {object} map[string]string "неверный тип справочника"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /dict/{dict} [post]
// @Security BearerAuth
func CreateDictItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dictType := c.Param("dict")

		switch dictType {
		case "position":
			var item models.Position
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, item)

		case "accessGroup":
			var item models.AccessGroup
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, item)

		case "subcategoryProduct":
			var item models.SubcategoryProduct
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, item)

		case "statuses":
			var item models.Status
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, item)

		case "paymentType":
			var item models.PaymentType
			if err := c.ShouldBindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, item)

		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный тип справочника"})
		}
	}
}

// DeleteDictItem godoc
// @Summary Удалить элемент справочника
// @Description Помечает элемент как удалённый (isDeleted = true)
// @Tags Dictionary
// @Param dict path string true "Тип справочника"
// @Param id path int true "ID элемента"
// @Success 204 {object} map[string]string "элемент справочника удалён"
// @Failure 400 {object} map[string]string "неверный тип справочника"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /dict/{dict}/{id} [delete]
// @Security BearerAuth
func DeleteDictItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		dictType := c.Param("dict")
		idStr := c.Param("id")
		model := getDictModel(dictType)
		if model == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный тип справочника"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
			return
		}

		result := db.Model(model).Where("id = ?", id).First(model)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "элемент не найден"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}

		if err := db.Model(model).Where("id = ?", id).Update("is_deleted", true).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "элемент успешно помечен как удалённый", "id": id})
	}
}
