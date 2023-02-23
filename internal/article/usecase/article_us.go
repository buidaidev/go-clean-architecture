package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
	articlerepo "go-clean-architecture/internal/article/repository"
)

type ArticleUseCase interface {
	CreateArticle(ctx context.Context, data *domain.ArticleCreate) error
	FetchArticleus(ctx context.Context, filter *domain.ArticleFilter, paging *app.Paging) ([]domain.Article, error)
	GetArticleus(ctx context.Context, condition map[string]interface{}) (*domain.Article, error)
	UpdateArticleus(ctx context.Context, data *domain.ArticleUpdate, id int) error
	DeleteArticleus(ctx context.Context, id int) error
}

type articleUseCase struct {
	articleRepo articlerepo.ArticleRepo
}

func NewArticleUseCase(articleRepo articlerepo.ArticleRepo) ArticleUseCase {
	return &articleUseCase{articleRepo: articleRepo}
}
