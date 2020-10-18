package postimage

type PostImage struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Image string `json:"image" gorm:"type:longtext null;"`
}
