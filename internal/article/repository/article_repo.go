package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"

	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(ctx context.Context, data *domain.ArticleCreate) error
	Fetch(ctx context.Context, filter *domain.ArticleFilter, paging *app.Paging) ([]domain.Article, error)
	GetByCondition(ctx context.Context, condition map[string]interface{}) (*domain.Article, error)
	Update(ctx context.Context, data *domain.ArticleUpdate, id int) error
	Delete(ctx context.Context, id int) error
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}
