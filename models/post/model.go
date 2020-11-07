package post

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/base"
)

// Post ... post model declaration
type Post struct {
	base.Post
}

// TableName ... Post Table Name in DB
func (p *Post) TableName() string {
	return "posts"
}
