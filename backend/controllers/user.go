package controllers

import (
	"aikido/forms"
	"aikido/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)
var userForm = new(forms.UserForm)

func (c *UserController) Login(ctx *gin.Context) {
	var form forms.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := userForm.Login(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	user, token, err := userModel.Login(ctx, form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id":    user.ID,
		"token": token,
	})
}

func (c *UserController) Register(ctx *gin.Context) {
	var form forms.RegisterForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := userForm.Register(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	user, token, err := userModel.Register(ctx, form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id":    user.ID,
		"token": token,
	})
}

func (c *UserController) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	user, err := userModel.Get(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	err := userModel.Delete(ctx, getUserID(ctx))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// TODO: LOGOUT USER!!

	ctx.Status(200)
}
