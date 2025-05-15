package request

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=doctor receptionist"`
}

type PatientRequest struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	DOB            string `json:"dob" binding:"required,datetime=2006-01-02"`
	Email          string `json:"email" binding:"omitempty,email"`
	Gender         string `json:"gender" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Address        string `json:"address" binding:"required"`
	MedicalHistory string `json:"medical_history" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
