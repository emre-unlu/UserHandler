package controllers

import (
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authService *services.AuthService

func InitilizeAuthController(service *services.AuthService) {
	authService = service
}

func Login(c *gin.Context) {
	var loginDTO dtos.LoginDto
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoToGenerate, err := authService.Login(loginDTO.Email, loginDTO.Password)
	accessToken, refreshToken, err := jwt.GenerateJWT(dtoToGenerate.Email, dtoToGenerate.Id)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to log in. Please try again later."})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func RefreshToken(c *gin.Context) {
	var refreshTokenDTO dtos.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshTokenDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := jwt.ValidateJWT(refreshTokenDTO.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	newAccessToken, _, err := jwt.GenerateJWT(claims.Email, claims.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to refresh token. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})

}
