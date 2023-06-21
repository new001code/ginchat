package service

import (
	"ginchat/models"
	"ginchat/sql"
	"ginchat/util"
)

type EmailService struct{}

var (
	templateMap = make(map[string]models.EmailTemplates)
)

var emailTemplateDao = &sql.EmailTemplateDao{}

func init() {
	util.Logger.Println("init email template")
	ls, err := emailTemplateDao.QueryAllTemplate()
	if err != nil {
		util.ErrorLogger.Println("init email template error", err)
	}
	for t := range ls {
		temp := ls[t]
		templateMap[temp.Name] = temp
	}
}
