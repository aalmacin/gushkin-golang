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

func (r *ActivityRepo) Activities() ([]*model.Activity, error) {
	var activities []*model.Activity
	err := r.DB.Model(&activities).Order("id").Select()
	if err != nil {
		fmt.Println("repos.Activity Activities() error: ", activities, err)
		return nil, errors.New("Something went wrong")
	}

	return activities, nil
}

func (r *ActivityRepo) ActivityById(id string) (*model.Activity, error) {
	var activity model.Activity
	err := r.DB.Model(&activity).Where("id = ?", id).First()
	if err != nil {
		fmt.Println("repos.Activity ActivityById() error: ", activity, err)
		return nil, errors.New("Something went wrong")
	}

	return &activity, nil
}
