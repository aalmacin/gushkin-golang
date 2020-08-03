package repos

import (
	"errors"
	"fmt"

	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type ActivityRepo struct {
	DB *pg.DB
}

func (r *ActivityRepo) Activities(userID string) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := r.DB.Model(&activities).Where("user_id = ?", userID).Order("id").Select()
	if err != nil {
		fmt.Println("repos.Activity Activities() error: ", activities, err)
		return nil, errors.New("Something went wrong")
	}

	return activities, nil
}

func (r *ActivityRepo) ActivityById(id string, userID string) (*model.Activity, error) {
	var activity model.Activity
	err := r.DB.Model(&activity).Where("id = ? AND user_id = ?", id, userID).First()
	if err != nil {
		fmt.Println("repos.Activity ActivityById() error: ", activity, err)
		return nil, errors.New("Something went wrong")
	}

	return &activity, nil
}

func (r *ActivityRepo) Create(input model.NewActivityInput, userID string) (*model.Activity, error) {
	activity := &model.Activity{
		Description: input.Description,
		FundAmt:     input.FundAmt,
		Positive:    *input.Positive,
		UserID:      userID,
	}
	_, err := r.DB.Model(activity).Returning("*").Insert()

	return activity, err
}
