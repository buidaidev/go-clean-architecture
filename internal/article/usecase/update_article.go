package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (us *articleUseCase) UpdateArticleus(ctx context.Context, data *domain.ArticleUpdate, id int) error {
	_, err := us.articleRepo.GetByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return app.ErrorCanNotGetEntity(domain.ArticleEntityName, err)
	}

	if err := us.articleRepo.Update(ctx, data, id); err != nil {
		return app.ErrorCanNotUpdateEntity(domain.ArticleEntityName, err)
	}

	return nil
}
