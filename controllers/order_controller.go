package controllers

import (
	"bakeryapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder godoc
// @Summary Создать заказ
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Новый заказ"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string "ошибка запроса"
// @Failure 500 {object} map[string]string "ошибка сервера"
// @Router /orders [post]
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Не задаём ID — он будет выставлен автоматически БД
		order.ID = 0
		order.CreateDate = time.Now()
		order.IsPay = false

		// Получаем товары из корзины клиента
		var basketItems []models.ProductInBasket
		if err := db.Where("clientid = ?", order.ClientId).Find(&basketItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить корзину клиента"})
			return
		}
		if len(basketItems) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "корзина клиента пуста"})
			return
		}

		// Создаём заказ
		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать заказ"})
			return
		}

		// Формируем записи в product_in_order
		var orderItems []models.ProductInOrder
		for _, item := range basketItems {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить информацию о товаре"})
				return
			}

			orderItems = append(orderItems, models.ProductInOrder{
				ProductID: item.ProductID,
				OrderID:   order.ID,
				Count:     item.Count,
				Cost:      product.Cost,
			})
		}

		// Вставляем позиции в product_in_order
		if err := db.Create(&orderItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить товары заказа"})
			return
		}

		// Очищаем корзину
		if err := db.Where("clientid = ?", order.ClientId).Delete(&models.ProductInBasket{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось очистить корзину"})
			return
		}

		c.JSON(http.StatusCreated, order)
	}
}

// GetAllOrders godoc
// @Summary Получить все заказы
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func GetAllOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Order
		db.Find(&orders)
		c.JSON(http.StatusOK, orders)
	}
}

// GetOrderById godoc
// @Summary Получить заказ по ID
// @Tags Orders
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} models.Order
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func GetOrderById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		id := c.Param("id")
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// UpdateOrder godoc
// @Summary Обновить заказ
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "ID заказа"
// @Param order body models.Order true "Обновленные данные"
// @Success 200 {object} models.Order
// @Failure 400,404 {object} map[string]string
// @Router /orders/{id} [put]
func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		id := c.Param("id")

		// Найти заказ
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
			return
		}

		// Получить данные из тела запроса
		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Удалить поле ID, если оно пришло
		delete(input, "id")

		// Обновить заказ
		if err := db.Model(&order).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить заказ"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

// DeleteOrder godoc
// @Summary Удалить заказ
// @Tags Orders
// @Param id path int true "ID заказа"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [delete]
func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		id := c.Param("id")

		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
			return
		}

		order.IsDeleted = true
		if err := db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось пометить заказ как удалённый"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "заказ помечен как удалённый"})
	}
}
