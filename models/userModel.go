package models

type User struct {
	ID           uint   `gorm:"primary key" json:"id"`
	FirstName    string `json:"first_name" validat:"required, min=2,max=100"`
	LastName     string `json:"last_name" validat:"required, min=2,max=100"`
	Password     string `json:"password" validate:"required, min=6"`
	Email        string `json:"email" validate:"email, required"`
	Phone        string `json:"phone" validate:"required"`
	Token        string `json:"token"`
	UserType     string `json:"user_type" validate:"required, eq==ADMIN|USER"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	UserID       string `json:"user_id"`
}
