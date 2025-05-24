package controllers

import (
	"net/http"
	"time"

	"bakeryapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateEmployee godoc
// @Summary Создать сотрудника
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body models.Employee true "Новый сотрудник"
// @Success 201 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Router /employees [post]
// @Security BearerAuth
func CreateEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Employee
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Обнуляем ID, чтобы автоинкремент сработал
		req.ID = 0

		// Установить текущую дату, если не передан StartDate
		if req.StartDate.IsZero() {
			req.StartDate = time.Now()
		}

		// Обнулить EndDate, если она нулевая
		if req.EndDate != nil && req.EndDate.IsZero() {
			req.EndDate = nil
		}

		if err := db.Create(&req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать сотрудника"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "employee created", "employee": req})
	}
}

// UpdateEmployee godoc
// @Summary Обновить сотрудника
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Param employee body models.Employee true "Обновленные данные"
// @Success 200 {object} models.Employee
// @Failure 400,404 {object} map[string]string
// @Router /employees/{id} [put]
// @Security BearerAuth
func UpdateEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")

		// Поиск сотрудника по ID
		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}

		// Считывание входных данных в map, чтобы исключить обновление ID
		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Удаляем поле id из входных данных, если оно передано
		delete(input, "id")

		// Если endDate передано как пустое или нулевое значение, зануляем
		if endDate, ok := input["enddate"]; ok {
			if endDate == nil || endDate == "" {
				input["enddate"] = nil
			}
		}

		// Обновляем только разрешённые поля
		if err := db.Model(&emp).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить сотрудника"})
			return
		}

		c.JSON(http.StatusOK, emp)
	}
}

// DeleteEmployee godoc
// @Summary Удалить сотрудника
// @Tags Employees
// @Param id path int true "ID сотрудника"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /employees/{id} [delete]
// @Security BearerAuth
func DeleteEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")

		// Поиск сотрудника
		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}

		// Установка даты завершения работы
		now := time.Now()
		if err := db.Model(&emp).Update("enddate", now).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось завершить сотрудника"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "сотрудник помечен как завершивший работу"})
	}
}

// GetAllEmployees godoc
// @Summary Получить всех сотрудников
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Router /employees [get]
// @Security BearerAuth
func GetAllEmployees(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emps []models.Employee
		db.Find(&emps)
		c.JSON(http.StatusOK, emps)
	}
}

// GetEmployeeById godoc
// @Summary Получить сотрудника по ID
// @Tags Employees
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} models.Employee
// @Failure 404 {object} map[string]string
// @Router /employees/{id} [get]
// @Security BearerAuth
func GetEmployeeById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")
		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}
		c.JSON(http.StatusOK, emp)
	}
}
