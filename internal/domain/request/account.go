package request



type SigninReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyOTPReq struct {
	OTP string `json:"otp" validate:"required"`
}

type UserSignupReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required,e164"`
}

type CreateAdminReq struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	Phone       string `json:"phone" validate:"required,e164"`
	Designation string `json:"designation"`
}

type UpdateAdminReq struct {
	Email       string `json:"email" validate:"omitempty,email"`
	Name        string `json:"name" validate:"omitempty"`
	Phone       string `json:"phone" validate:"omitempty,e164"`
	Designation string `json:"designation" validate:"omitempty"`
}

type UserAddressReq struct {
    House    string    `gorm:"column:house" json:"house"`
    Street1  string    `gorm:"column:street1" json:"street1"`
    Street2  string    `gorm:"column:street2" json:"street2"`
    City     string    `gorm:"column:city" json:"city"`
    State    string    `gorm:"column:state" json:"state"`
    Pincode  string    `gorm:"column:pincode" json:"pincode"`
    Landmark string    `gorm:"column:landmark" json:"landmark"`
    Country  string    `gorm:"column:country" json:"country"`
}
