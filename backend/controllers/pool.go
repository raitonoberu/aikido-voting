package controllers

import (
	"aikido/forms"
	"aikido/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PoolController struct{}

var poolModel = new(models.PoolModel)
var poolForm = new(forms.PoolForm)

func (c *PoolController) Create(ctx *gin.Context) {
	var form forms.CreatePoolForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := poolForm.Create(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	poolID, err := poolModel.Create(ctx, getUserID(ctx), form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id": poolID,
	})
}

func (c *PoolController) Available(ctx *gin.Context) {
	pools, err := poolModel.Available(ctx, getUserID(ctx))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, pools)
}

func (c *PoolController) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	pool, err := poolModel.Get(ctx, getUserID(ctx), id)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, pool)
}

func (c *PoolController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = poolModel.Delete(ctx, getUserID(ctx), id)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
