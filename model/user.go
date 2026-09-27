package model

type User struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"id"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Username string `gorm:"column:username;unique" json:"username"`
	Password string `gorm:"column:password" json:"-"`
	Role     string `gorm:"column:role;default:user" json:"role"`
}

func (User) TableName() string {
	return "users"
}
