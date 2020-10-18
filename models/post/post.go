package post

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/models/postimage"
)

type PostInterface interface {
	List() (*[]Post, error)
	Get(id int) (Post, error)
	Create(post *Post) (*Post, error)
	Update(post *Post) (*Post, error)
	Delete(id int) error
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

	// Create new empty image
	var postImageModel postimage.PostImage
	var err error
	postImage := &postimage.PostImage{
		ID:    post.ID,
		Image: "",
	}
	postImage, err = postImageModel.Create(postImage)
	if err != nil {
		return nil, err
	}
	post.PostImage = *postImage
	return post, nil
}

func (t *Post) Update(post *Post) (*Post, error) {
	if err := connection.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (t *Post) Delete(id uint) error {
	if err := connection.Delete(&Post{}, id).Error; err != nil {
		return err
	}
	var postImageModel postimage.PostImage
	if err := postImageModel.Delete(id); err != nil {
		return err
	}
	return nil
}
