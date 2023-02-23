package articlerepo

import (
	"context"
	"go-clean-architecture/app"
	"go-clean-architecture/domain"
)

func (r *articleRepo) Delete(
	ctx context.Context,
	id int,
) error {

	if err := r.db.
		Table(domain.Article{}.TableName()).
		Where("id =? ", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return app.ErrorDB(err)
	}

	return nil
}
