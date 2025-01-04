package request

type SuperAdminSigninReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateAdminReq struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone" validate:"required,e164"`
	Designation string `json:"designation"`
}