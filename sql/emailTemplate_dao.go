package sql

import (
	"ginchat/models"
	"ginchat/util"
)

type EmailTemplateDao struct{}

func (et *EmailTemplateDao) QueryAllTemplate() ([]models.EmailTemplates, error) {
	db := util.GetDB()
	tm := []models.EmailTemplates{}
	if err := db.Table("email_templates").Where("status = '1'").Find(&tm).Error; err != nil {
		return nil, err
	}
	return tm, nil
}
