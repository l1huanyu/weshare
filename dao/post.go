package dao

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	CREATED = iota
	SET_TYPE
	SET_NAME
	SET_SOURCE
	SUCCEED
)

type Post struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Source      string
	Description string
	Publisher   string `gorm:"not null"`
	Type        int    `gorm:"index:type_index"`
	Like        int
	Version     int
	State       int
}

func (p *Post) Display() string {
	return fmt.Sprintf("No. %d\n《%s》\n获取方式：%s\n描述：%s", p.ID, p.Name, p.Source, p.Description)
}

func (p *Post) Create() error {
	if db.NewRecord(p) {
		return db.Create(p).Error
	}
	return errors.New("Not New Record")
}

func (p *Post) Update() error {
	return db.Save(p).Error
}

func QueryPostByType(tp int) (*Post, error) {
	p := new(Post)
	if err := db.Raw(fmt.Sprintf("select * from posts where type = %d and state = %d order by random() limit 1", tp, SUCCEED)).Scan(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func QueryPostRandomly() (*Post, error) {
	p := new(Post)
	if err := db.Raw(fmt.Sprintf("select * from posts where state = %d order by random() limit 1", SUCCEED)).Scan(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func QueryUnfinishedPost() (*Post, error) {
	p := new(Post)
	if err := db.Where("state < ?", SUCCEED).First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func CountPosts() int {
	count := 0
	if err := db.Table("posts").Where("state = ?", SUCCEED).Count(&count).Error; err == nil {
		_PostsCount = count
	}
	return _PostsCount
}

func CountPostsByPublisher(publisher string) int {
	count := 0
	if err := db.Model(&Post{}).Where("state = ?", SUCCEED).Where("publisher = ?", publisher).Count(&count).Error; err == nil {
		_PublishersPostsCount = count
	}
	return _PublishersPostsCount
}
