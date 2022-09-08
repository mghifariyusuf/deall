package entity

import "time"

// User .
type User struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Salt        *string    `json:"-"`
	RoleID      string     `json:"role_id"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at"`
	UpdatedBy   *string    `json:"updated_by,omitempty"`
}
