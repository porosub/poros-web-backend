package user_type

type User_Type struct {
	Id	uint	`gorm:"primaryKey" json:"id" binding:"required"`
	Name string `gorm:"not null" json:"name" binding:"required"`
}

func (b *User_Type) TableName() string {
	return "user_type"
}