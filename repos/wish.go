package repos

import (
	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type Wish struct {
	DB *pg.DB
}

func (r *Wish) Wishs() ([]*model.Wish, error) {
	var wishItems []*model.Wish
	err := r.DB.Model(&wishItems).Select()
	if err != nil {
		return nil, err
	}

	return wishItems, nil
}
