package service

import (
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
		util.DebugLogger.Printf("user register param: %+v\n", user)
		RedisUtil.SetString("email", user.Email, 30*time.Second)
		err = userDao.Create(&user)
		if err == nil {
			c.JSON(http.StatusOK, ApiResponse.SuccessDefault())
		} else {
			c.JSON(http.StatusOK, ApiResponse.FailWithMessage(err.Error()))
		}
	} else {
		util.ErrorLogger.Println("参数解析错误：", err)
	}

}

func (u *UserService) Login(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		util.DebugLogger.Printf("user login param: %+v\n", user)

	} else {
		util.ErrorLogger.Println("参数解析错误：", err)
	}
}

func (u *UserService) LoginCheckCode(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		util.DebugLogger.Printf("user email: %s\n", user.Email)
		EmailUtil.SendText([]string{user.Email}, "checkCode", "this is checkcode")
	} else {
		util.ErrorLogger.Println("参数解析错误：", err)
	}
}

func (u *UserService) GetPublicKey(c *gin.Context) {
	c.JSON(http.StatusOK, ApiResponse.SuccessWithData(util.RSAPrivateKey.PublicKey))
}
