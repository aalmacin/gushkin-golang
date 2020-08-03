package repos

import (
	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type WishRepo struct {
	DB *pg.DB
}

func (r *WishRepo) Wishes(userID string) ([]*model.Wish, error) {
	var wishes []*model.Wish
	err := r.DB.Model(&wishes).Where("user_id = ?", userID).Select()
	if err != nil {
		return nil, err
	}

	return wishes, nil
}

func (r *WishRepo) Create(input model.NewWishInput, userID string) (*model.Wish, error) {
	wish := &model.Wish{
		Description: input.Description,
		Price:       input.Price,
		Priority:    input.Priority.String(),
		Source:      input.Source,
		Status:      input.Status.String(),
		UserID:      userID,
	}
	_, err := r.DB.Model(wish).Returning("*").Insert()

	return wish, err
}

func (r *WishRepo) Update(input model.UpdateWishInput, userID string) (*model.Wish, error) {
	wish := &model.Wish{
		Description: *input.Description,
		Price:       *input.Price,
		Priority:    input.Priority.String(),
		Source:      input.Source,
		Status:      input.Status.String(),
		UserID:      userID,
	}
	_, err := r.DB.Model(wish).Where("id = ?", input.ID).Returning("*").Update()

	return wish, err
}
