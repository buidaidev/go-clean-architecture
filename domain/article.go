package domain

import (
	"go-clean-architecture/pkg/db/mysql"
)

const ArticleEntityName = "Article"

type Article struct {
	mysql.Model
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
	Status  int    `json:"-" gorm:"column:status;"`
}

type ArticleCreate struct {
	mysql.Model
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

type ArticleUpdate struct {
	Title   *string `json:"title" gorm:"column:title;"`
	Content *string `json:"content" gorm:"column:content;"`
}

type ArticleFilter struct {
	Status []int `json:"-"`
}

func (Article) TableName() string { return "articles" }
