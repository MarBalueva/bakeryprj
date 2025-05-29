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
// @Security BearerAuth
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

		var order models.Order
		if err := db.First(&order, req.OrderID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "заказ не найден"})
			return
		}

		if req.Sum < order.SumOrder {
			c.JSON(http.StatusBadRequest, gin.H{"error": "оплаченная сумма меньше суммы заказа"})
			return
		}

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
// @Security BearerAuth
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
		orderId := c.Param("orderid")

		userIdVal, exists := c.Get("userid")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "не авторизован"})
			return
		}
		userId := userIdVal.(int64)

		var client models.Client
		if err := db.Where("user_id = ? AND isdeleted = false", userId).First(&client).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "клиент не найден"})
			return
		}

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

// @Summary Получить статус платежа
// @Description Получает поле statusid платежа по ID
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path int true "ID платежа"
// @Success 200 {object} models.StatusResponse
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/{id}/status [GET]
// @Security BearerAuth
func GetPaymentStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payment models.Payment

		if err := db.Select("statusid").First(&payment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "платеж не найден"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"statusid": payment.StatusId})
	}
}

// @Summary Обновить статус платежа
// @Description Обновляет поле statusid платежа по ID
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path int true "ID платежа"
// @Param status body models.StatusInput true "Новый статус"
// @Success 200 {object} models.Payment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/{id}/status [PUT]
// @Security BearerAuth
func UpdatePaymentStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payment models.Payment

		if err := db.First(&payment, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "платеж не найден"})
			return
		}

		var input models.StatusInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payment.StatusId = input.StatusId
		if err := db.Save(&payment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить статус платежа"})
			return
		}

		c.JSON(http.StatusOK, payment)
	}
}
