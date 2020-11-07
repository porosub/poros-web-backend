package usertype

// UserType ... User Type Declaration
type UserType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name" binding:"required"`
}

// TableName ... User Type Table Name
func (b *UserType) TableName() string {
	return "user_type"
}
