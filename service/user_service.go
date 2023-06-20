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

var ApiResponse = &models.ApiResponse{}

// log
var Log = &util.Log{}

func (u *UserService) Register(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		Log.Debugf("user register param: %+v\n", user)
		err = userDao.Create(&user)
		if err == nil {
			c.JSON(http.StatusOK, ApiResponse.SuccessDefault())
		} else {
			c.JSON(http.StatusOK, ApiResponse.FailWithMessage(err.Error()))
		}
	} else {
		Log.Errorln("参数解析错误：", err)
	}

}

func (u *UserService) Login(c *gin.Context) {
	var user query.UserQuery
	if err := c.ShouldBind(&user); err == nil {
		Log.Debugf("user login param: %+v\n", user)

	} else {
		Log.Errorln("参数解析错误：", err)
	}
}
