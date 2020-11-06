package tags

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/base"
)

type Tag struct {
	base.Tag
}

func (t *Tag) TableName() string {
	return "tags"
}
