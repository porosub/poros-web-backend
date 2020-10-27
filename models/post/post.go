package post

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/models/postimage"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
	"gorm.io/gorm"
)

type PostInterface interface {
	List() (*[]Post, error)
	Get(id int) (Post, error)
	Create(post *Post) (*Post, error)
	Update(post *Post) (*Post, error)
	Delete(id int) error
	LinkImagesName(post *Post, images []string) error
	DeletePostImages(post *Post) error
}

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

func (t *Post) List() (*[]Post, error) {
	var posts []Post
	if err := connection.Preload("User").Preload("User.User_Type").Preload("PostType").Preload("PostImage").Find(&posts).Error; err != nil {
		return nil, err
	}

	// Clear user password output
	cleanPosts := []Post{}
	for _, post := range posts {
		post.User.Password = ""
		cleanPosts = append(cleanPosts, post)
	}
	return &cleanPosts, nil
}

func (t *Post) Get(id uint) (*Post, error) {
	var post Post
	if err := connection.Where("id = ?", id).Preload("User").Preload("User.User_Type").Preload("PostType").Preload("PostImage").First(&post).Error; err != nil {
		return nil, err
	}
	post.User.Password = ""
	return &post, nil
}

func (t *Post) Create(post *Post) (*Post, error) {
	if err := connection.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (t *Post) Update(post *Post) (*Post, error) {
	if err := connection.Session(&gorm.Session{FullSaveAssociations: true}).Omit("User", "PostType").Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (t *Post) Delete(id uint) error {
	p, err := t.Get(id)
	if err != nil {
		return err
	}

	if err = t.DeletePostImages(p); err != nil {
		return err
	}

	if err = connection.Delete(&Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

// LinkImagesName ... Link images name with post object
func (t *Post) LinkImagesName(post *Post, images []string) error {
	for _, image := range images {
		var postImage = postimage.PostImage{
			Image: image,
		}
		post.PostImage = append(post.PostImage, postImage)
	}
	return nil
}

func (t *Post) DeletePostImages(post *Post) error {
	for _, postImage := range post.PostImage {
		path := config.AssetPostsImages + postImage.Image
		if err := storage.RemoveFile(path); err != nil {
			return err
		}
		if err := connection.Delete(postImage).Error; err != nil {
			return err
		}
	}
	post.PostImage = []postimage.PostImage{}
	return nil
}
