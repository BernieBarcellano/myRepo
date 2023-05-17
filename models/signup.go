package models

import "gorm.io/gorm"

type Sign_up struct {
	ID               uint   `gorm:"primary key;autoIncrement" json:"id"`
	FullName         string `json:"fullname"`
	Student_Number   string `json:"student_number"`
	Course           string `json:"course"`
	Year_Level       string `json:"year_level"`
	Email_Address    string `json:"email_address"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm_password"`
}

func MigrateSign_up(db *gorm.DB) error {
	err := db.AutoMigrate(&Sign_up{})
	return err
}
