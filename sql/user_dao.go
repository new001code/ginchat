package sql

import (
	"database/sql"
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/util"
	"time"
)

type UserDao struct{}

func (d *UserDao) Create(param *query.UserQuery) {
	if param.Username == "" {

	}
	db := util.GetDB()
	tx := db.Begin()
	uid := util.GetUUID()
	b := models.Base{
		ID:         uid,
		Status:     sql.NullString{String: "1", Valid: true},
		CreatorId:  sql.NullInt64{Int64: int64(uid), Valid: true},
		CreateTime: sql.NullTime{Time: time.Now(), Valid: true},
		ModifierId: sql.NullInt64{},
		ModifyTime: sql.NullTime{},
	}
	u := models.Users{
		Base:     b,
		Username: param.Username,
		Password: util.PasswordEncode(param.Password),
	}
	db.Create(&u)
}
