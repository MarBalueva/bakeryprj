package controllers

import (
	"bakeryapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts godoc
// @Summary Получить список продуктов
// @Description Получить список продуктов
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Список продуктов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /products [get]
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if err := db.Where("is_deleted = false").Find(&products).Error; err != nil {
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
// @Security BearerAuth
func GetProductByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.Where("id = ? AND is_deleted = false", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
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
// @Security BearerAuth
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

		// Добавить запись в priceHistory
		priceEntry := models.PriceHistory{
			ProductID: product.ID,
			Cost:      product.Cost,
			StartDate: time.Now(),
			EndDate:   nil,
		}
		if err := db.Create(&priceEntry).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "product created, but failed to add price history"})
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

		// Найти продукт по ID
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		oldCost := product.Cost

		// Прочитать тело запроса в map
		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверка изменения цены
		if costVal, ok := input["cost"].(float64); ok && costVal != oldCost {
			// Закрыть старую запись
			db.Model(&models.PriceHistory{}).
				Where("productid = ? AND enddate IS NULL", product.ID).
				Update("enddate", time.Now())

			// Добавить новую запись
			priceEntry := models.PriceHistory{
				ProductID: product.ID,
				Cost:      costVal,
				StartDate: time.Now(),
			}
			db.Create(&priceEntry)
		}

		// Обновить поля
		if err := db.Model(&product).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
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
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		db.Model(&product).Update("is_deleted", true)
		c.JSON(http.StatusOK, gin.H{"message": "Product marked as deleted"})
	}
}
