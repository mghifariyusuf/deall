package userhttp

type insertRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	RoleID      string `json:"role_id"`
	CreatedBy   string `json:"created_by"`
}

type updateRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	RoleID      string `json:"role_id"`
	UpdatedBy   string `json:"updated_by"`
}

type deleteRequest struct {
	DeletedBy string `json:"deleted_by"`
}
