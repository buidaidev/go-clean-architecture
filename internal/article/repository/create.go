package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (r *articleRepo) Create(ctx context.Context, data *domain.ArticleCreate) error {
	if err := r.db.
		Table(domain.Article{}.TableName()).
		Create(&data).Error; err != nil {
		return app.ErrorDB(err)
	}

	return nil
}
