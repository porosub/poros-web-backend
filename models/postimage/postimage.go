package postimage

import "github.com/divisi-developer-poros/poros-web-backend/config"

type PostImageInterface interface {
	Get(id int) (PostImage, error)
	Update(postImage *PostImage) (*PostImage, error)
	Delete(id int) error
}

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

func (t *PostImage) Get(id uint) (*PostImage, error) {
	var postImage PostImage
	if err := connection.Where("id = ?", id).First(&postImage).Error; err != nil {
		return nil, err
	}
	return &postImage, nil
}

func (t *PostImage) Create(postImage *PostImage) (*PostImage, error) {
	if err := connection.Create(postImage).Error; err != nil {
		return nil, err
	}
	return postImage, nil
}

func (t *PostImage) Update(postImage *PostImage) (*PostImage, error) {
	if err := connection.Save(postImage).Error; err != nil {
		return nil, err
	}
	return postImage, nil
}

func (t *PostImage) Delete(id uint) error {
	if err := connection.Delete(&PostImage{}, id).Error; err != nil {
		return err
	}
	return nil
}
