package articleus

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (us *articleUseCase) DeleteArticleus(ctx context.Context, id int) error {
	_, err := us.articleRepo.GetByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return app.ErrorCanNotGetEntity(domain.ArticleEntityName, err)
	}

	if err := us.articleRepo.Delete(ctx, id); err != nil {
		return app.ErrworCanNotDeleteEntity(domain.ArticleEntityName, err)
	}

	return nil
}
