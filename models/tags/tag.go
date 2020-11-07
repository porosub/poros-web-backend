package tags

import (
	"errors"

	"github.com/divisi-developer-poros/poros-web-backend/config"
)

// TagInterface ... Tag interface declaration
type TagInterface interface {
	FetchTags() ([]Tag, error)
	FetchTagByID(id int) (*Tag, error)
	CreateTag(newTag *Tag) error
	UpdateTagByID(updatedTag *Tag) error
	DeleteTagByID(id int) error
}

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

// FetchTags ... Get all tags from DB
func (t *Tag) FetchTags() (*[]Tag, error) {
	var tags []Tag
	if err := connection.Find(&tags).Error; err != nil {
		return nil, err
	}
	return &tags, nil
}

// FetchTagByID ... Get single tag from Db
func (t *Tag) FetchTagByID(id int) (*Tag, error) {
	var tag Tag
	if err := connection.Where("id = ?", id).Find(&tag).Error; err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, errors.New("data not found")
	}
	return &tag, nil
}

// CreateTag ... Create single tag to DB
func (t *Tag) CreateTag(newTag *Tag) (*Tag, error) {
	result := connection.Create(&newTag)
	if result.Error != nil {
		return nil, result.Error
	}
	return newTag, nil
}

// UpdateTagByID ... Update single tag from DB
func (t *Tag) UpdateTagByID(updatedTag *Tag) (*Tag, error) {
	// if data exists
	if _, err := t.FetchTagByID(updatedTag.ID); err != nil {
		return nil, err
	}

	result := connection.Save(&updatedTag)
	if result.Error != nil {
		return nil, result.Error
	}
	return updatedTag, nil
}

// DeleteTagByID ... Delete single tag from DB
func (t *Tag) DeleteTagByID(id int) error {
	// if data exists
	if _, err := t.FetchTagByID(id); err != nil {
		return err
	}

	result := connection.Delete(&Tag{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
