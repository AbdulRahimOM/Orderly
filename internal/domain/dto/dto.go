package dto

type Credentials struct {
	ID             int    `json:"id" gorm:"column:id"`
	Username       string `json:"username" gorm:"column:username"`
	HashedPassword string `json:"-" gorm:"column:hashed_password"`
}
