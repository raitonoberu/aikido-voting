package controllers

import (
	"aikido/forms"
	"aikido/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatController struct{}

var chatModel = new(models.ChatModel)
var chatForm = new(forms.ChatForm)

func (c *ChatController) WriteMessage(ctx *gin.Context) {
	var form forms.WriteMessageForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := chatForm.WriteMessage(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = chatModel.WriteMessage(ctx, getUserID(ctx), id, form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *ChatController) ReadMessages(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	count, _ := strconv.Atoi(ctx.DefaultQuery("count", "20"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	messages, err := chatModel.ReadMessages(ctx, getUserID(ctx), id, count, offset)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, messages)
}
