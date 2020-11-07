package post

import (
	"mime/multipart"

	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/models/base"
	"github.com/divisi-developer-poros/poros-web-backend/models/postimage"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
)

// PostInterface ... Post interface declaration
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

// List ... Get all posts from DB
func (t *Post) List() (*[]Post, error) {
	var posts []Post
	if err := connection.Preload("User").Preload("PostType").Preload("PostImages").Preload("Tags").Find(&posts).Error; err != nil {
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

// Get ... Get single post from DB
func (t *Post) Get(id uint) (*Post, error) {
	var post Post
	if err := connection.Where("id = ?", id).Preload("PostType").Preload("PostImages").Preload("User").Preload("Tags").First(&post).Error; err != nil {
		return nil, err
	}
	post.User.Password = ""
	return &post, nil
}

// Create ... Create single post to DB
func (t *Post) Create(post *Post, imagesBlob []*multipart.FileHeader) (*Post, error) {
	if err := connection.Create(post).Error; err != nil {
		return nil, err
	}

	imagesName, err := storeImages(imagesBlob)
	if err != nil {
		return nil, err
	}

	if err := appendImages(post, imagesName); err != nil {
		return nil, err
	}

	return post, nil
}

// Update ... Update single post from DB
func (t *Post) Update(post *Post, id int, imagesBlob []*multipart.FileHeader) (*Post, error) {
	oldPost, err := t.Get(uint(id))
	if err != nil {
		return nil, err
	}

	if err := connection.Model(&oldPost).Updates(&post).Error; err != nil {
		return nil, err
	}

	*post = *oldPost
	imagesName, err := storeImages(imagesBlob)
	if err != nil {
		return nil, err
	}

	// If there's no image uploaded then update finished
	if len(imagesName) == 0 {
		return post, nil
	}

	if err := t.deletePostImages(post); err != nil {
		return nil, err
	}

	if err := appendImages(post, imagesName); err != nil {
		return nil, err
	}

	return post, nil
}

// Delete ... Delete single post from DB
func (t *Post) Delete(id uint) error {
	p, err := t.Get(id)
	if err != nil {
		return err
	}

	if err = t.deletePostImages(p); err != nil {
		return err
	}

	if err = connection.Delete(&Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

// AttachTags attach tags to corresponding post
func (t *Post) AttachTags(id int, tags *[]base.Tag) (p *Post, err error) {
	p = &Post{}
	p.ID = uint(id)
	if err := connection.Model(&p).Association("Tags").Append(tags); err != nil {
		return nil, err
	}

	p, err = t.Get(uint(id))
	if err != nil {
		return nil, err
	}
	return p, nil
}

// DetachTags detach tags from corresponding post
func (t *Post) DetachTags(id int, tags *[]base.Tag) (*Post, error) {

	p, err := t.Get(uint(id))
	if err != nil {
		return nil, err
	}

	if err := connection.Model(&p).Association("Tags").Delete(tags); err != nil {
		return nil, err
	}
	return p, nil
}

func appendImages(post *Post, images []string) error {
	var postImages []postimage.PostImage
	for _, image := range images {
		var postImage = postimage.PostImage{
			Image: image,
		}
		postImages = append(postImages, postImage)
	}
	connection.Model(post).Association("PostImages").Append(postImages)
	return nil
}

func (t *Post) deletePostImages(post *Post) error {
	for _, postImage := range post.PostImages {
		path := config.AssetPostsImages + postImage.Image
		if err := storage.RemoveFile(path); err != nil {
			return err
		}
		if err := connection.Delete(postImage).Error; err != nil {
			return err
		}
	}
	post.PostImages = []postimage.PostImage{}
	return nil
}

func storeImages(imagesBlob []*multipart.FileHeader) (filenames []string, err error) {
	if len(imagesBlob) == 0 {
		return filenames, nil
	}

	filenames, err = storage.StoreFilesBlob(imagesBlob, config.AssetPostsImages)
	if err != nil {
		return filenames, err
	}
	return filenames, nil
}
