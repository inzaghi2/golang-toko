package handlers

import (
	"golang-toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}
		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"message": "Invalid credentials"})
			return
		}
		//pastikan untuk memeriksa kata sandi yang benar disini

		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		c.JSON(200, gin.H{"message": "Login successful", "token": token})

	}

}
