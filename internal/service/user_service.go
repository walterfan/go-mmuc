package service

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/walterfan/go-mmuc/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"

)

type UserService struct {
	db *gorm.DB
}

func NewUserService(username string, password string, host string, port int, database string) (*UserService, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("failed to connect to database: %v", err)
		return nil, err
	}
	db.AutoMigrate(&domain.User{})
	return &UserService{db: db}, nil
}

func (us *UserService) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.New()
	us.db.Create(&user)
	c.JSON(http.StatusOK, user)
}

func (us *UserService) ListUsers(c *gin.Context) {
	var users []domain.User
	us.db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func (us *UserService) GetUser(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	if err := us.db.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (us *UserService) UpdateUser(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	if err := us.db.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	us.db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func (us *UserService) DeleteUser(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	us.db.Delete(&user, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
