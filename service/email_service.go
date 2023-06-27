package service

import (
	"ginchat/common"
	"ginchat/models"
	"ginchat/sql"
)

type EmailService struct{}

var (
	templateMap = make(map[string]models.EmailTemplates)
)

var emailTemplateDao = &sql.EmailTemplateDao{}

func init() {
	common.Logger.Info("init email template")
	ls, err := emailTemplateDao.QueryAllTemplate()
	if err != nil {
		common.Logger.Error("init email template error", err)
	}
	for t := range ls {
		temp := ls[t]
		templateMap[temp.Name] = temp
	}
}
