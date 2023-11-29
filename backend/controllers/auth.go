package controllers

import (
	"aikido/models"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authModel = new(models.AuthModel)

func getUserID(c *gin.Context) (userID int64) {
	return c.MustGet("userID").(int64)
}

func (c AuthController) Middleware(ctx *gin.Context) {
	token := authModel.ExtractToken(ctx.Request)
	if token == "" {
		ctx.JSON(401, gin.H{
			"error": "unauthorized",
		})
		ctx.Abort()
		return
	}

	claims, err := authModel.VerifyToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": "couldn't verify token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userID", claims.ID)
}
