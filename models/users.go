package models

type Users struct {
	Base     Base `gorm:"embedded"`
	Username string
	Password string
}
