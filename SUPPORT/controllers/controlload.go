// 登录控制器
package controllers

import (
	"log"
	"net/http"

	"SUPPORT/global"
	"SUPPORT/models"

	"SUPPORT/utils"

	"github.com/gin-gonic/gin"
)

// 注册
func AddUser(ctx *gin.Context) {
	var load models.Load //请求体与数据库模型绑定
	if err := ctx.ShouldBindJSON(&load); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"code":    6,
			"message": "invalid request body",
		})
		return
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(load.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    5,
			"message": "failed to hash password",
			"error":   err.Error(),
		})
		return
	}

	//身份验证
	load.Password = hashedPassword
	token, err := utils.GenerateJWT(load.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    4,
			"message": "failed to generate token",
			"error":   err.Error(),
		})
		return
	}

	//自动迁移，方便表的初始化和更新
	if err := global.Db.AutoMigrate(&load); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    3,
			"message": "failed to migrate table",
			"error":   err.Error(),
		})
		return
	}

	// 创建表
	result := global.Db.Create(&load)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    1,
			"message": "failed to create user",
			"error":   result.Error.Error(),
		})
		return
	} else if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"code":    2,
			"message": "user not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"code":    0,
			"message": "user created",
			"token":   token,
		})
	}
}

// 登录
func LoadUser(ctx *gin.Context) {
	var input struct {
		Userid   string `json:"userid"`
		Password string `json:"password"`
	}
	var load models.Load
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"code":    5,
			"message": "invalid request body",
		})
		return
	}

	err := global.Db.Where("userid = ?", input.Userid).First(&load).Error

	if err != nil && !utils.CheckPassword(input.Password, load.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"code":    3,
			"message": "invalid userid",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"code":    2,
			"message": "invalid userid",
		})
		return
	}

	if !utils.CheckPassword(input.Password, load.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"code":    1,
			"message": "invalid password",
		})
		return
	}

	token, err := utils.GenerateJWT(load.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    4,
			"message": "failed to generate token",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"code":    0,
		"message": "login success",
		"token":   token,
	})
}

// 注销
func DeleteUser(ctx *gin.Context) {
	Userid := ctx.Param("Userid")
	result := global.Db.Delete(&models.Load{}, "Userid = ?", Userid)
	log.Printf("Attempting to delete user with Userid: %s", Userid)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    1,
			"message": "failed to delete user",
			"error":   result.Error.Error(),
		})
		return
	} else if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"code":    2,
			"message": "user not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"code":    0,
			"message": "user deleted",
		})
	}
}

// 修改密码
func UpdateUser(ctx *gin.Context) {
	var input struct {
		Userid      string `json:"userid"`
		OldPassword string `json:"oldpassword"`
		NewPassword string `json:"newpassword"`
	}

	var load models.Load
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"code":    3,
			"message": "invalid request body",
		})
		return
	}

	if err := global.Db.Where("userid = ?", input.Userid).First(&load).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"code":    4,
			"message": "invalid userid",
		})
		return
	}

	if !utils.CheckPassword(input.OldPassword, load.Password) {
		log.Printf("Checking password: input: %s, hashed: %s", input.OldPassword, load.Password)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"code":    4,
			"message": "invalid old password",
		})
		return
	}
	log.Printf("Loaded User: %+v", load)

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    1,
			"message": "failed to hash password",
			"error":   err.Error(),
		})
		return
	}

	load.Password = hashedPassword
	result := global.Db.Save(&load)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"code":    1,
			"message": "failed to update user",
			"error":   result.Error.Error(),
		})
		return
	} else if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"code":    2,
			"message": "user not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"code":    0,
			"message": "user updated",
		})
	}
}
