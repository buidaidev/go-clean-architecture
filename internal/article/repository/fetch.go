package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (r *articleRepo) Fetch(
	ctx context.Context,
	filter *domain.ArticleFilter,
	paging *app.Paging,
) ([]domain.Article, error) {
	var result []domain.Article

	db := r.db.Table(domain.Article{}.TableName())

	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, app.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(offset)

	if err := r.db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, app.ErrorDB(err)
	}

	return result, nil
}
