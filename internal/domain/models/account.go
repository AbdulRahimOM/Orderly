package models

const (
	SuperAdmin_TableName = "super_admin"
)

type SuperAdmin struct {
	ID             int    `gorm:"column:id;primaryKey"`
	Username       string `gorm:"column:username;unique"`
	HashedPassword string `gorm:"column:hashed_password"`
}

func (SuperAdmin) TableName() string {
	return SuperAdmin_TableName
}
