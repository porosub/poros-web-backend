package postimage

// PostImage model
type PostImage struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement:false;"`
	Image string `json:"image" gorm:"primaryKey;autoIncrement:false"`
}
