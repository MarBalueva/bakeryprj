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

		req.ID = 0

		if req.StartDate.IsZero() {
			req.StartDate = time.Now()
		}

		if req.EndDate != nil && req.EndDate.IsZero() {
			req.EndDate = nil
		}

		if err := db.Create(&req).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать сотрудника"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "сотрудник создан", "employee": req})
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

		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}

		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		delete(input, "id")

		if val, ok := input["jobpositionid"]; ok {
			switch v := val.(type) {
			case float64:
				if v <= 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "jobpositionid должно быть больше 0"})
					return
				}
			case int:
				if v <= 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "jobpositionid должно быть больше 0"})
					return
				}
			case int64:
				if v <= 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "jobpositionid должно быть больше 0"})
					return
				}
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "неверный тип для jobpositionid"})
				return
			}
		}

		if endDate, ok := input["enddate"]; ok {
			if endDate == nil || endDate == "" || endDate == "string" {
				input["enddate"] = nil
			}
		}

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
// @Failure 404,500 {object} map[string]string
// @Router /employees/{id} [delete]
// @Security BearerAuth
func DeleteEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")

		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}
		if !emp.IsDeleted {
			c.JSON(http.StatusBadRequest, gin.H{"error": "сотрудник уже удалён"})
			return
		}

		if err := db.Model(&emp).Update("isdeleted", true).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить сотрудника"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "сотрудник помечен как удаленный"})
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
		if err := db.Where("isdeleted = ?", false).Find(&emps).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить сотрудников"})
			return
		}
		c.JSON(http.StatusOK, emps)
	}
}

// GetEmployeeById godoc
// @Summary Получить сотрудника по ID
// @Tags Employees
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} models.Employee
// @Failure 404,500 {object} map[string]string
// @Router /employees/{id} [get]
// @Security BearerAuth
func GetEmployeeById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")
		if err := db.Where("id = ? AND isdeleted = ?", id, false).First(&emp).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}
		c.JSON(http.StatusOK, emp)
	}
}

// @Summary Уволить сотрудника
// @Description Устанавливает поле enddate текущей датой для сотрудника с заданным ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} models.Employee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /employees/{id}/fire [PUT]
// @Security BearerAuth
func FireEmployee(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var emp models.Employee
		id := c.Param("id")

		if err := db.First(&emp, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "сотрудник не найден"})
			return
		}

		now := time.Now()
		emp.EndDate = &now

		if err := db.Save(&emp).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось уволить сотрудника"})
			return
		}

		c.JSON(http.StatusOK, emp)
	}
}
