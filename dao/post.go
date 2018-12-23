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
	PENDING
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
	display := fmt.Sprintf("No.%d\n墙裂安利！《%s》\n类型：%sο(=•ω＜=)ρ⌒☆", p.ID, p.Name, p.GetType())
	if len(p.Source) != 0 {
		display = display + fmt.Sprintf("\n【伪】传送门：%s...", p.Source)
	}
	if len(p.Description) != 0 {
		display = display + fmt.Sprintf("\n安利理由：%sΣ(っ °Д °;)っ", p.Description)
	}
	likeCount := p.LikeCount()
	if likeCount != 0 {
		display = display + fmt.Sprintf("\n已有[%d]位客官点赞了这份安利👍~", likeCount)
	}
	return display
}

func (p *Post) GetType() string {
	t := ""
	switch p.Type {
	case 1:
		t = "电影"
	case 2:
		t = "电视剧"
	case 3:
		t = "游戏"
	case 4:
		t = "动漫"
	case 5:
		t = "小说"
	case 6:
		t = "漫画"
	case 7:
		t = "其他"
	}
	return t
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

func QueryPostByType(tp int, postIDs []uint) (*Post, error) {
	query := ""
	if len(postIDs) != 0 {
		for id := range postIDs {
			query = query + fmt.Sprintf("and where id != %d ", id)
		}
	}
	p := new(Post)
	if err := db.Raw(fmt.Sprintf("select * from posts where type = %d and state = %d %s order by random() limit 1", tp, SUCCEED, query)).Scan(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func QueryPostRandomly(postIDs []uint) (*Post, error) {
	query := ""
	if len(postIDs) != 0 {
		for id := range postIDs {
			query = query + fmt.Sprintf("and where id != %d ", id)
		}
	}
	p := new(Post)
	if err := db.Raw(fmt.Sprintf("select * from posts where state = %d %s order by random() limit 1", SUCCEED, query)).Scan(p).Error; err != nil {
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

func QueryPendingPosts(offset, limit int) ([]Post, error) {
	posts := make([]Post, 0)
	if err := db.Find(&posts, "state = ?", PENDING).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}
	return posts, nil
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

func (p *Post) LikeCount() int {
	count := 0
	db.Model(new(LikeLink)).Where("post_id = ?", p.ID).Count(&count)
	return count
}
