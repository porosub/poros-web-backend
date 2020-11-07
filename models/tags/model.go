package tags

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/base"
)

// Tag ... Tag model declaration
type Tag struct {
	base.Tag
}

// TableName ... Tag table name
func (t *Tag) TableName() string {
	return "tags"
}
