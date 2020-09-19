package user

type User struct {
	Id           uint		`gorm:"primaryKey" json:"id"`
	Image        string		`json:"image"`
	Username     string		`json:"username"`
	Password     string		`json:"password"`
	Full_name    string		`json:"full_name"`
	User_type_id uint		`gorm:"index" json:"user_type_id"`
}

func (b *User) TableName() string {
	return "users"
}