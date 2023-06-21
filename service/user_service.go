package service

import (
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/sql"
	"ginchat/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

// userDao
var userDao = &sql.UserDao{}

var MyLog = &util.MyLog{}

var ApiResponse = &models.ApiResponse{}

func (u *UserService) Register(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		MyLog.Errorf("user register param: %+v\n", user)
		err = userDao.Create(&user)
		if err == nil {
			c.JSON(http.StatusOK, ApiResponse.SuccessDefault())
		} else {
			c.JSON(http.StatusOK, ApiResponse.FailWithMessage(err.Error()))
		}
	} else {
		MyLog.Errorln("参数解析错误：", err)
	}

}

func (u *UserService) Login(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		MyLog.Debugf("user login param: %+v\n", user)

	} else {
		MyLog.Errorln("参数解析错误：", err)
	}
}

func (u *UserService) LoginCheckCode(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		MyLog.Debugf("user email: %s\n", user.Email)
	} else {
		MyLog.Errorln("参数解析错误：", err)
	}
}
