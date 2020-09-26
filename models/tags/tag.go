package tags

import (
	"errors"
	"github.com/divisi-developer-poros/poros-web-backend/config"
)

type TagInterface interface {
	FetchTags() ([]Tag, error)
	FetchTagByID(id int) (*Tag, error)
	CreateTag(newTag *Tag) error
	UpdateTagByID(updatedTag *Tag) error
	DeleteTagByID(id int) error
}

var connection, err = config.MysqlConn()

func (t *Tag) FetchTags() (*[]Tag, error) {
	if err != nil {
		return nil, err
	}

	var tags []Tag
	if err := connection.Find(&tags).Error; err != nil {
		return nil, err
	}
	return &tags, nil
}

func (t *Tag) FetchTagByID(id int) (*Tag, error) {
	if err != nil {
		return nil, err
	}

	var tag Tag
	if err := connection.Where("id = ?", id).Find(&tag).Error; err != nil {
		return nil, err
	}

	if tag.ID == 0 {
		return nil, errors.New("data not found")
	}
	return &tag, nil
}

func (t *Tag) CreateTag(newTag *Tag) error {
	if err != nil {
		return err
	}

	if err := connection.Create(&newTag).Error; err != nil {
		return err
	}
	return nil
}

func (t *Tag) UpdateTagByID(updatedTag *Tag) error {
	if err != nil {
		return err
	}

	// if data exists
	if _, err := t.FetchTagByID(updatedTag.ID); err != nil {
		return err
	}

	if err := connection.Save(&updatedTag).Error; err != nil {
		return err
	}
	return nil
}

func (t *Tag) DeleteTagByID(id int) error {
	if err != nil {
		return err
	}

	if err := connection.Delete(&Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}
