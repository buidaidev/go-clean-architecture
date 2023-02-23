package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (r *articleRepo) Update(ctx context.Context, data *domain.ArticleUpdate, id int) error {
	if err := r.db.
		Table(domain.Article{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return app.ErrorDB(err)
	}

	return nil
}
