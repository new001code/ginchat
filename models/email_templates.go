package models

type EmailTemplates struct {
	Base     Base `gorm:"embedded"`
	Name     string
	IsInner  string
	Template string
}
