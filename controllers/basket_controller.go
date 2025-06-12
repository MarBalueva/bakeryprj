package controllers

import (
	"bakeryapp/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddProductToBasket godoc
// @Summary Добавить товар в корзину клиента
// @Tags Basket
// @Accept json
// @Produce json
// @Param productInBasket body models.ProductInBasket true "Товар и количество для добавления"
// @Success 200 {object} map[string]string "товар добавлен в корзину"
// @Failure 400 {object} map[string]string "ошибка валидации"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /basket [post]
func AddProductToBasket(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.ProductInBasket

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "некорректные данные: " + err.Error()})
			return
		}

		var existing models.ProductInBasket
		err := db.Where("productid = ? AND clientid = ?", item.ProductID, item.ClientID).First(&existing).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка базы данных"})
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&item).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось добавить товар в корзину"})
				return
			}
		} else {
			existing.Count += item.Count
			if err := db.Save(&existing).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить количество товара"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "товар добавлен в корзину"})
	}
}

// RemoveProductFromBasket godoc
// @Summary Удалить товар из корзины клиента
// @Tags Basket
// @Produce json
// @Param clientId path int true "ID клиента"
// @Param productId path int true "ID товара"
// @Success 200 {object} map[string]string "товар удален из корзины"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /basket/{clientId}/{productId} [delete]
func RemoveProductFromBasket(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := c.Param("clientid")
		productId := c.Param("productid")

		if err := db.Where("productid = ? AND clientid = ?", productId, clientId).Delete(&models.ProductInBasket{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить товар из корзины"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "товар удален из корзины"})
	}
}

// GetBasketByClientID godoc
// @Summary Получить корзину клиента
// @Tags Basket
// @Produce json
// @Param clientId path int true "ID клиента"
// @Success 200 {array} models.ProductInBasket "Список товаров в корзине"
// @Failure 404 {object} map[string]string "корзина не найдена"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /basket/{clientId} [get]
func GetBasketByClientID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := c.Param("clientId")

		var items []models.ProductInBasket
		err := db.Where("clientid = ?", clientId).Find(&items).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при получении корзины"})
			return
		}
		if len(items) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "корзина не найдена"})
			return
		}

		c.JSON(http.StatusOK, items)
	}
}
