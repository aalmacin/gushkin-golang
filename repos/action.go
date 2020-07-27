package repos

import (
	"fmt"

	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type ActionRepo struct {
	DB *pg.DB
}

func (r *ActionRepo) actionsByField(field string, value interface{}) ([]*model.Action, error) {
	var activityActions []*model.Action
	err := r.DB.Model(&activityActions).Where(fmt.Sprintf("%s = ?", field), value).Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, nil
}

func (r *ActionRepo) ActionsByActivityId(id string) ([]*model.Action, error) {
	return r.actionsByField("activity_id", id)
}
