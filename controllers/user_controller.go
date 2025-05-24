package controllers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"bakeryapp/models"
	"bakeryapp/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUser godoc
// @Summary Создать пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.Appuser true "Новый пользователь"
// @Success 201 {object} models.Appuser
// @Failure 400 {object} map[string]string
// @Router /users [post]
// @Security BearerAuth
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Appuser
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Проверка логина на уникальность
		var existing models.Appuser
		if err := db.Where("login = ?", req.Login).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "login already in use"})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// Обработка null-значений
		if req.ClientId != nil && *req.ClientId == 0 {
			req.ClientId = nil
		}
		if req.EmpId != nil && *req.EmpId == 0 {
			req.EmpId = nil
		}

		// Хэширование пароля
		hashedPwd, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}

		// Создание пользователя
		user := models.Appuser{
			Login:      req.Login,
			Password:   hashedPwd,
			EmpId:      req.EmpId,
			ClientId:   req.ClientId,
			CreateDate: time.Now(),
			IsActive:   true,
		}

		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				c.JSON(http.StatusConflict, gin.H{"error": "login already in use"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create user"})
			}
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user created"})
	}
}

// UpdateUser godoc
// @Summary Обновить пользователя
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.Appuser true "Обновленные данные"
// @Success 200 {object} models.Appuser
// @Failure 400,404 {object} map[string]string
// @Router /users/{id} [put]
// @Security BearerAuth
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Appuser
		id := c.Param("id")

		// Найти пользователя
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
			return
		}

		// Получить входные данные как map
		var input map[string]interface{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Удалить поле id, если оно передано
		delete(input, "id")

		// Обработка пароля: хэшируем, если передан
		if pwdRaw, ok := input["password"].(string); ok && pwdRaw != "" {
			hashedPwd, err := utils.HashPassword(pwdRaw)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при хэшировании пароля"})
				return
			}
			input["password"] = hashedPwd
		}

		// Обработка clientId и empId: 0 → nil
		if clientID, ok := input["clientId"].(float64); ok && clientID == 0 {
			input["clientId"] = nil
		}
		if empID, ok := input["empId"].(float64); ok && empID == 0 {
			input["empId"] = nil
		}

		// Обновление
		if err := db.Model(&user).Updates(input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить пользователя"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Tags Users
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
// @Security BearerAuth
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Appuser
		id := c.Param("id")

		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
			return
		}

		user.IsDeleted = true
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось пометить пользователя как удаленного"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "пользователь помечен как удаленный"})
	}
}

// GetAllUsers godoc
// @Summary Получить всех пользователей (кроме удалённых)
// @Tags Users
// @Produce json
// @Success 200 {array} models.Appuser
// @Router /users [get]
// @Security BearerAuth
func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.Appuser
		if err := db.Where("is_deleted = false").Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при получении пользователей"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUserByID godoc
// @Summary Получить пользователя по ID
// @Tags Users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.Appuser
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
// @Security BearerAuth
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Appuser
		id := c.Param("id")

		if err := db.Where("id = ? AND is_deleted = false", id).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
