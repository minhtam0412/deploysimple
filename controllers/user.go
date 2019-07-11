package controllers

import (
	"deploysimple/dto"
	"deploysimple/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetAllUserRequest(ctx *gin.Context) {
	var lstUser []dto.User
	lstUser, err := user.GetAllUser()
	if err != nil {
		ctx.AbortWithStatusJSON(500, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	ctx.AbortWithStatusJSON(200, dto.Result{ErrorMessage: "", Data: lstUser})
}

func Ping(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(200, dto.Result{ErrorMessage: "", Data: "pong"})
}

func HandleAddUserRequest(ctx *gin.Context) {
	var objUser dto.User
	if err := ctx.ShouldBindJSON(&objUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	newID, err := user.AddNewUser(objUser);
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, dto.Result{ErrorMessage: "", Data: newID})

}

func HandleGetUserByKey(ctx *gin.Context) {
	var keyword = ctx.Param("id")
	var rsl dto.User
	rsl, err := user.GetUserByKey(keyword)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: rsl})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, dto.Result{Data: rsl, ErrorMessage: ""})
}

func HandleUpdateUserRequest(ctx *gin.Context) {
	var objUser dto.User
	if err := ctx.ShouldBindJSON(&objUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	err := user.UpdateUser(objUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: "", Data: "Update success!"})
}

func HandleDeleteUserRequest(ctx *gin.Context) {
	var id = ctx.Query("id")
	var deleteduser = ctx.Query("deleteduser")
	err := user.DeleteUser(id, deleteduser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{ErrorMessage: err.Error(), Data: nil})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, dto.Result{ErrorMessage: "", Data: "Delete success"})
}
