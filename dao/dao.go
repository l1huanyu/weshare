package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	_DB_TYPE     = "postgres"
	_DB_HOST     = "weshare_pg_test"
	_DB_PORT     = "5432"
	_DB_USER     = "l1huanyu"
	_DB_PASSWORD = "Li960127!"
	_DB_NAME     = "postgres"
)

var db *gorm.DB
var _PostsCount int
var _PublishersPostsCount int

func init() {
	var err error
	db, err = gorm.Open(_DB_TYPE, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", _DB_HOST, _DB_PORT, _DB_USER, _DB_PASSWORD, _DB_NAME))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database, err = %s", err.Error()))
	}
	db.AutoMigrate(new(Post))
}

func CloseDB() {
	db.Close()
}
