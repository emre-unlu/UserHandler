package controllers

import (
	"github.com/emre-unlu/GinTest/internal/config"
	"github.com/emre-unlu/GinTest/internal/dtos"
	"github.com/emre-unlu/GinTest/internal/services"
	"github.com/emre-unlu/GinTest/internal/utils"
	"github.com/emre-unlu/GinTest/pkg/jwt"
	"github.com/gin-gonic/gin"
	myjwt "github.com/golang-jwt/jwt/v4"
	"net/http"
)

var authService *services.AuthService

func InitilizeAuthController(service *services.AuthService) {
	authService = service
}

func Login(c *gin.Context) {
	var loginDTO dtos.LoginDto
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	dtoToGenerate, err := authService.Login(loginDTO.Email, loginDTO.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	accessToken, refreshToken, err := jwt.GenerateJWT(dtoToGenerate.Email, dtoToGenerate.Id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
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
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	claims := jwt.Claims{}
	myjwt.ParseWithClaims(refreshTokenDTO.RefreshToken, claims, func(token *myjwt.Token) (interface{}, error) {
		return config.JWTSecretKey, nil
	})

	newAccessToken, _, err := jwt.GenerateJWT(claims.Email, claims.Id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})

}
