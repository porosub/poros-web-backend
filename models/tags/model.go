package tags

type Tag struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
}

func (t *Tag) TableName() string {
	return "tags"
}