package service

import (
	"encoding/json"
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

// userDao
var userDao = &sql.UserDao{}

var ApiResponse = &models.ApiResponse{}

func (u *UserService) Register(c *gin.Context) {
	data, _ := c.GetRawData()
	var user query.UserQuery
	_ = json.Unmarshal(data, &user)
	err := userDao.Create(&user)
	if err == nil {
		c.JSON(http.StatusOK, ApiResponse.SuccessDefault())
	} else {
		c.JSON(http.StatusOK, ApiResponse.FailWithMessage(err.Error()))
	}

}
