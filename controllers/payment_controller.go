package controllers

import (
	"bakeryapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePaymentRequest struct {
	OrderID   int64   `json:"orderid"`
	Sum       float64 `json:"sum"`
	PayTypeID int     `json:"paytypeid"`
}

// CreatePayment godoc
// @Summary Провести оплату по заказу
// @Tags Payments
// @Accept json
// @Produce json
// @Param payment body CreatePaymentRequest true "Информация о платеже"
// @Success 201 {object} models.Payment
// @Failure 400 {object} map[string]string "ошибка валидации или сумма меньше суммы заказа"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /payments [post]
func CreatePayment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreatePaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Найти заказ
		var order models.Order
		if err := db.First(&order, req.OrderID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "заказ не найден"})
			return
		}

		// Проверка суммы
		if req.Sum < order.SumOrder {
			c.JSON(http.StatusBadRequest, gin.H{"error": "оплаченная сумма меньше суммы заказа"})
			return
		}

		// Создать платеж
		payment := models.Payment{
			OrderID:   req.OrderID,
			Sum:       req.Sum,
			Date:      time.Now(),
			PayTypeID: req.PayTypeID,
		}
		if err := db.Create(&payment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать платеж"})
			return
		}

		// Обновить флаг оплаты заказа
		order.IsPay = true
		if err := db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить заказ"})
			return
		}

		c.JSON(http.StatusCreated, payment)
	}
}

// GetAllPayments godoc
// @Summary Получить все платежи
// @Tags Payments
// @Produce json
// @Success 200 {array} models.Payment
// @Failure 500 {object} map[string]string
// @Router /payments [get]
func GetAllPayments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payments []models.Payment
		if err := db.Find(&payments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить список платежей"})
			return
		}
		c.JSON(http.StatusOK, payments)
	}
}

// GetPaymentsByOrder godoc
// @Summary Получить платежи по ID заказа
// @Tags Payments
// @Produce json
// @Param orderId path int true "ID заказа"
// @Success 200 {array} models.Payment
// @Failure 404 {object} map[string]string "заказ не найден или нет платежей"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Router /payments/order/{orderId} [get]
// GetPaymentsByOrder godoc
// @Summary Получить платежи клиента по ID заказа
// @Tags Payments
// @Produce json
// @Param orderId path int true "ID заказа"
// @Success 200 {array} models.Payment
// @Failure 403 {object} map[string]string "доступ запрещен"
// @Failure 404 {object} map[string]string "платежи не найдены"
// @Failure 500 {object} map[string]string "внутренняя ошибка сервера"
// @Security Bearer
// @Router /payments/order/{orderId} [get]
func GetPaymentsByOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")

		// Получаем userId из middleware (устанавливается в JWT или другом механизме авторизации)
		userIdVal, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "не авторизован"})
			return
		}
		userId := userIdVal.(int64)

		// Получаем клиента по userId
		var client models.Client
		if err := db.Where("user_id = ? AND is_deleted = false", userId).First(&client).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "клиент не найден"})
			return
		}

		// Проверяем, что заказ принадлежит клиенту
		var order models.Order
		if err := db.Where("id = ? AND clientid = ?", orderId, client.ID).First(&order).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "нет доступа к заказу"})
			return
		}

		var payments []models.Payment
		if err := db.Where("orderid = ?", orderId).Find(&payments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при получении платежей"})
			return
		}
		if len(payments) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "платежи не найдены"})
			return
		}

		c.JSON(http.StatusOK, payments)
	}
}
