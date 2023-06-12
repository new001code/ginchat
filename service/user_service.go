package service

import (
	"encoding/json"
	query "ginchat/models/query"
	"ginchat/sql"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

// userDao
var userDao = &sql.UserDao{}

func (u *UserService) Register(c *gin.Context) {
	data, _ := c.GetRawData()
	var user query.UserQuery
	_ = json.Unmarshal(data, &user)
	userDao.Create(&user)

	c.JSON(200, gin.H{
		"message": "user register",
	})
}
