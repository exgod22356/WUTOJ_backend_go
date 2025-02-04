package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdateUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "比赛期间无法修改个人信息", 0))
	return
	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	var userJson model.User
	userModel := model.User{}
	userValidate := validate.UserValidate

	if err := c.ShouldBindUri(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userJson.UserID = GetUserIdFromSession(c)

	log.Print(userJson)
	userMap := helper.Struct2Map(userJson)

	if res, err := userValidate.ValidateMap(userMap, "update_info"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res = userModel.EditUserByID(userJson.UserID, userJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func SearchUser(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	userJson := struct {
		Param string `json:"param" form:"param" uri:"param"`
	}{}
	userModel := model.User{}

	if err := c.ShouldBindUri(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res = userModel.SearchUser(userJson.Param)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
