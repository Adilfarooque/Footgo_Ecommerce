package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Fristname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
