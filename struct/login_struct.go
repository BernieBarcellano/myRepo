package structs

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Sign_up struct {
	FullName         string `json:"fullname"`
	Student_Number   string `json:"student_number"`
	Course           string `json:"course"`
	Year_Level       string `json:"year_level"`
	Email_Address    string `json:"email_address"`
	Password         string `json:"password"`
	Confirm_Password string `json:"confirm_password"`
}
