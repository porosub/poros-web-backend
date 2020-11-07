package postimage

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/utils/host"
)

// PostImage model
type PostImage struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement:false;"`
	Image string `json:"image" gorm:"primaryKey;autoIncrement:false"`
}

// LocalizedField ... Localized all field to corresponded host
func (p *PostImage) LocalizedField() {
	if p.Image != "" {
		p.Image = host.GetURL() + config.AssetPostsImages + p.Image
	}
}
