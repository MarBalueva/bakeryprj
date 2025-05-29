package controllers

import (
	"bakeryapp/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts godoc
// @Summary Получить список продуктов
// @Description Получить список продуктов, с возможной фильтрацией по категории
// @Tags Products
// @Accept json
// @Produce json
// @Param categoryid query int false "ID категории для фильтрации"
// @Success 200 {array} models.Product "Список продуктов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /products [get]
// @Security BearerAuth
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		query := db.Where("isdeleted = false")

		if categoryIdStr := c.Query("categoryid"); categoryIdStr != "" {
			if categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64); err == nil && categoryId > 0 {
				query = query.Where("categoryid = ?", categoryId)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный categoryid"})
				return
			}
		}

		if err := query.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}

// GetProductByID godoc
// @Summary Получение продукта по ID
// @Tags Products
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.Where("id = ? AND isdeleted = false", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "продукт не найден"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// CreateProduct godoc
// @Summary Добавить продукт (только для admin и manager)
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.Product true "Новый продукт"
// @Success 201 {object} models.Product
// @Failure 403 {object} map[string]string "access denied"
// @Failure 401 {object} map[string]string "unauthorized"
// @Router /products/ [post]
func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		priceEntry := models.PriceHistory{
			ProductID: product.ID,
			Cost:      product.Cost,
			StartDate: time.Now(),
			EndDate:   nil,
		}
		if err := db.Create(&priceEntry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "продукт создан, но обновление цены не удалось"})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

// UpdateProduct godoc
// @Summary Обновление продукта
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param product body models.Product true "Обновленные данные продукта"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [put]
// @Security BearerAuth
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		id := c.Param("id")

		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "продукт не найден"})
			return
		}

		oldCost := product.Cost

		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fieldsToCheck := []string{
			"categoryid",
			"proteins",
			"fats",
			"carbohydrates",
			"calories",
			"unweight",
			"weight",
			"cost",
		}

		for _, field := range fieldsToCheck {
			if val, exists := input[field]; exists {
				switch v := val.(type) {
				case float64:
					if v <= 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": field + " должно быть больше 0"})
						return
					}
				case int:
					if v <= 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": field + " должно быть больше 0"})
						return
					}
				case int64:
					if v <= 0 {
						c.JSON(http.StatusBadRequest, gin.H{"error": field + " должно быть больше 0"})
						return
					}
				default:
					c.JSON(http.StatusBadRequest, gin.H{"error": "неверный тип для поля " + field})
					return
				}
			}
		}

		if costVal, ok := input["cost"].(float64); ok && costVal != oldCost {
			db.Model(&models.PriceHistory{}).
				Where("productid = ? AND enddate IS NULL", product.ID).
				Update("enddate", time.Now())

			priceEntry := models.PriceHistory{
				ProductID: product.ID,
				Cost:      costVal,
				StartDate: time.Now(),
			}
			db.Create(&priceEntry)
		}

		if err := db.Model(&product).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить продукт"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// DeleteProduct godoc
// @Summary Удаление продукта
// @Tags Products
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
// @Security BearerAuth
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.First(&product, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "продукт не найден"})
			return
		}
		if !product.IsDeleted {
			c.JSON(http.StatusBadRequest, gin.H{"error": "продукт уже удалён"})
			return
		}

		product.IsDeleted = true
		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось пометить продукт как удаленный"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "продукт помечен как удаленный"})
	}
}
