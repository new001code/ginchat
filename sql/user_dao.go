package sql

import (
	"database/sql"
	"errors"
	"ginchat/models"
	query "ginchat/models/query"
	"ginchat/util"

	// "ginchat/util"
	"time"
)

type UserDao struct{}

func (d *UserDao) Create(param *query.UserQuery) error {
	if param.Username == "" {
		return errors.New("username is empty")
	}
	if param.Password == "" {
		return errors.New("password is empty")
	}
	if param.Email == "" {
		return errors.New("email is empty")
	}
	if param.CheckCode == "" {
		return errors.New("checkCode is empty")
	}
	db := util.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var count int64
	tx.Model(&models.Users{}).Where("username = ?", param.Username).Count(&count)

	if count > 0 {
		return errors.New("this username has exited")
	}
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

func (d *UserDao) Find(param *query.UserQuery) error {
	if param.Username == "" {
		return errors.New("username is empty")
	}
	if param.Password == "" {
		return errors.New("password is empty")
	}
	db := util.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return tx.Commit().Error

}
