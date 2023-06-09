package models

// data model
type Login struct {
	Email    string `json:"email_address"`
	Password string `json:"password"`
}

type SignUp struct {
	ID               uint   `gorm:"primary key;autoIncrement" json:"id"`
	FullName         string `json:"full_name"`
	Student_Number   string `json:"student_number"`
	Course           string `json:"course"`
	Year_Level       string `json:"year_level"`
	Email_Address    string `json:"email_address"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm_password"`
}
type StudentDashboar struct {
	FullName       string `json:"full_name"`
	Student_Number int    `json:"student_number"`
	Course         string `json:"course"`
	Year           string `json:"year"`
	Contact_Number string `json:"contact_number"`
	Email          string `json:"email"`
}
type Attendance struct {
	ID           uint   `json:"id"`
	Full_Name    string `json:"full_name"`
	Subject      string `json:"subject"`
	Block_No     string `json:"block_no`
	Today_Date   string `json:"today_date`
	Current_Time string `json:"current_time`
}
