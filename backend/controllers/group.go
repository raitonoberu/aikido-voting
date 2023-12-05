package controllers

import (
	"aikido/forms"
	"aikido/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupController struct{}

var userGroupModel = new(models.UserGroupModel)
var groupModel = new(models.GroupModel)
var groupForm = new(forms.GroupForm)

func (c *GroupController) All(ctx *gin.Context) {
	groups, err := groupModel.All(ctx)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, groups)
}

func (c *GroupController) Create(ctx *gin.Context) {
	var form forms.CreateGroupForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := groupForm.Create(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	groupID, err := groupModel.Create(ctx, getUserID(ctx), form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"id": groupID,
	})
}

func (c *GroupController) Update(ctx *gin.Context) {
	var form forms.UpdateGroupForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := groupForm.Update(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = groupModel.Update(ctx, id, getUserID(ctx), form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *GroupController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = groupModel.Delete(ctx, getUserID(ctx), id)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *GroupController) Users(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	users, err := userModel.ByGroup(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (c *GroupController) Add(ctx *gin.Context) {
	var form forms.AddUserForm
	if err := ctx.ShouldBind(&form); err != nil {
		msg := groupForm.AddUser(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": msg})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = userGroupModel.Add(ctx, getUserID(ctx), id, form)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}

func (c *GroupController) Remove(ctx *gin.Context) {
	groupID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if groupID == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userID == 0 || err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid parameter"})
		return
	}

	err = userGroupModel.Remove(ctx, getUserID(ctx), userID, groupID)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(200)
}
