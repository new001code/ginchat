package service

import (
	"ginchat/common"
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/sql"
	"ginchat/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

// userDao
var userDao = &sql.UserDao{}

var ApiResponse = &models.ApiResponse{}

var EmailUtil = &util.EmailUtil{}
var RedisUtil = &util.RedisUtil{}

func (u *UserService) Register(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		common.Logger.Debugf("user register param: %+v", user)
		RedisUtil.SetString("email", user.Email, 30*time.Second)
		err = userDao.Create(&user)
		if err == nil {
			c.JSON(http.StatusOK, ApiResponse.SuccessDefault())
		} else {
			c.JSON(http.StatusOK, ApiResponse.FailWithMessage(err.Error()))
		}
	} else {
		common.Logger.Error("参数解析错误：", err)
	}

}

func (u *UserService) Login(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		common.Logger.Debugf("user login param: %+v", user)

	} else {
		common.Logger.Error("参数解析错误：", err)
	}
}

func (u *UserService) LoginCheckCode(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		common.Logger.Debugf("user email: %s", user.Email)
		EmailUtil.SendText([]string{user.Email}, "checkCode", "this is checkcode")
	} else {
		common.Logger.Error("参数解析错误：", err)
	}
}

func (u *UserService) GetPublicKey(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse.SuccessWithData(util.RSAPrivateKey.PublicKey))
}
