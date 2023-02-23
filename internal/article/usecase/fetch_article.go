package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (us *articleUseCase) FetchArticleus(
	ctx context.Context,
	filter *domain.ArticleFilter,
	paging *app.Paging,
) ([]domain.Article, error) {
	result, err := us.articleRepo.Fetch(ctx, filter, paging)

	if err != nil {
		return result, app.ErrorCanNotFetchEntity(domain.ArticleEntityName, err)
	}

	return result, nil
}
