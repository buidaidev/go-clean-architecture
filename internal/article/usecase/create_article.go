package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (us *articleUseCase) CreateArticle(ctx context.Context, data *domain.ArticleCreate) error {
	if err := us.articleRepo.Create(ctx, data); err != nil {
		return app.ErrorCanNotCreateEntity(domain.ArticleEntityName, err)
	}

	return nil
}
