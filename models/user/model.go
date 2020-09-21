package user

type User struct {
	Id           uint		`gorm:"primaryKey" json:"id"`
	Image        string		`gorm:"not null" json:"image"`
	Username     string		`gorm:"not null" json:"username"`
	Password     string		`gorm:"not null" json:"password"`
	Full_name    string		`gorm:"not null" json:"full_name"`
	User_type_id uint		`gorm:"not null" gorm:"index" json:"user_type_id"`
}

func (b *User) TableName() string {
	return "users"
}