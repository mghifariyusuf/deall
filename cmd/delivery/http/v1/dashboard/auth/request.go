package http

type loginRequest struct {
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password" validate:"required,min=5"`
}
