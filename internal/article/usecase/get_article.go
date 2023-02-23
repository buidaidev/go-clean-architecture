package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (us *articleUseCase) GetArticleus(ctx context.Context, condition map[string]interface{}) (*domain.Article, error) {
	data, err := us.articleRepo.GetByCondition(ctx, condition)

	if err != nil {
		return nil, app.ErrorCanNotGetEntity(domain.ArticleEntityName, err)
	}

	if data.Status == 0 {
		return nil, app.ErrorEntityDeleted(domain.ArticleEntityName, err)
	}

	return data, nil
}
