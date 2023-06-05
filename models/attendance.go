package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type StudentAttendance struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Subject     string `json:"subject"`
	BlockNo     string `json:"block_no"`
	TodayDate   string `json:"today_date"`
	CurrentTime string `json:"current_time"`
	Code        string `json:"code"`
}

func generateAttendanceCode() string {
	const codeLength = 6
	const codeCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, codeLength)
	for i := range code {
		code[i] = codeCharset[rand.Intn(len(codeCharset))]
	}

	return string(code)
}

func main() {
	attendance := StudentAttendance{
		ID:          1,
		FullName:    "John Doe",
		Subject:     "Mathematics",
		BlockNo:     "B-12",
		TodayDate:   "2023-06-05",
		CurrentTime: "09:30 AM",
		Code:        generateAttendanceCode(),
	}

	attendanceJSON, err := json.Marshal(attendance)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(attendanceJSON))
}
