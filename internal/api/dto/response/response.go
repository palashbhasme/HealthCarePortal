package response

import "time"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type PatientResponse struct {
	ID             uint      `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          *string   `json:"email,omitempty"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	MedicalHistory string    `json:"medical_history"`
	DOB            time.Time `json:"dob"`
}
