package controllers

import (
	"aikido/forms"
	"aikido/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

var voteModel = new(models.VoteModel)
var voteForm = new(forms.VoteForm)

func (c *VoteController) Create(ctx *gin.Context) {
	var form forms.CreateVoteForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := voteForm.Create(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	poolID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if poolID == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = voteModel.Create(ctx, getUserID(ctx), poolID, form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *VoteController) Delete(ctx *gin.Context) {
	poolID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if poolID == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = voteModel.Delete(ctx, getUserID(ctx), poolID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
