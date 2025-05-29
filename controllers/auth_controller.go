package controllers

import (
	"bakeryapp/models"
	"bakeryapp/utils"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Login      string  `json:"login"`
	Password   string  `json:"password"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Email      string  `json:"email"`
	Phone      *string `json:"phone"`
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "Данные для регистрации"
// @Success 201 {object} map[string]string "user registered"
// @Failure 400 {object} map[string]string "ошибка валидации"
// @Failure 500 {object} map[string]string "не удалось создать пользователя"
// @Router /register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existing models.Appuser
		if err := db.Where("login = ?", req.Login).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "логин уже используется"})
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		hashedPwd, _ := utils.HashPassword(req.Password)

		client := models.Client{
			Name:        req.Name,
			Surname:     req.Surname,
			Patronymic:  req.Patronymic,
			Email:       req.Email,
			PhoneNumber: req.Phone,
		}
		if err := db.Create(&client).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать клиента"})
			return
		}

		user := models.Appuser{
			Login:      req.Login,
			Password:   hashedPwd,
			ClientId:   &client.ID,
			CreateDate: time.Now(),
			IsActive:   true,
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать пользователя"})
			return
		}

		accessSQL := `INSERT INTO user_accesses (userid, groupid) VALUES (?, ?)`
		if err := db.Exec(accessSQL, user.ID, 1).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось назначить группу доступа"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "регистрация успешна"})
	}
}

// Login godoc
// @Summary Авторизация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Данные входа"
// @Success 200 {object} map[string]string "token"
// @Failure 401 {object} map[string]string "unauthorized"
// @Router /login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.Appuser
		if err := db.Where(`login = ? AND "isActive" = true`, req.Login).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token, _ := utils.GenerateToken(user.ID)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
