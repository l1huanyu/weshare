package dao

import (
	"github.com/jinzhu/gorm"
)

type LikeLink struct {
	gorm.Model
	PostID uint
	UserID string
}

func Like(userID string, postID uint) {
	if !IsLike(userID, postID) {
		l := new(LikeLink)
		l.UserID = userID
		l.PostID = postID
		db.Create(l)
	}
}

func IsLike(userID string, postID uint) bool {
	l := new(LikeLink)
	err := db.Model(new(LikeLink)).Where("user_id = ?", userID).Where("post_id = ?", postID).First(l).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	return true
}
