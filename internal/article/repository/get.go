package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"

	"gorm.io/gorm"
)

func (r *articleRepo) GetByCondition(ctx context.Context, condition map[string]interface{}) (*domain.Article, error) {
	var data domain.Article

	if err := r.db.
		Where(condition).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, app.ErrorEntityNotFound(domain.ArticleEntityName, err)
		}
		return nil, app.ErrorDB(err)
	}

	return &data, nil
}
