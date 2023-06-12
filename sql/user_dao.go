package sql

import (
	"database/sql"
	"errors"
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/util"
	"time"
)

type UserDao struct{}

func (d *UserDao) Create(param *query.UserQuery) error {
	if param.Username == "" {
		return errors.New("username is null")
	}
	if param.Password == "" {
		return errors.New("password is null")
	}
	db := util.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var count uint64
	tx.Model(&model.Users{}).Where("username = ?", param.Username).Count(&count)
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
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
