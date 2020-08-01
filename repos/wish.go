package repos

import (
	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type WishRepo struct {
	DB *pg.DB
}

func (r *WishRepo) Wishes() ([]*model.Wish, error) {
	var wishes []*model.Wish
	err := r.DB.Model(&wishes).Select()
	if err != nil {
		return nil, err
	}

	return wishes, nil
}
