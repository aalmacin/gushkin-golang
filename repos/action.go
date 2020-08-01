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

func (r *ActionRepo) ActionsWithOptions(options *model.GetActionInput) ([]*model.Action, error) {
	var activityActions []*model.Action
	model := r.DB.Model(&activityActions)

	if *options.Today {
		model.Where(fmt.Sprintf("action_timestamp = current_date"))
	}

	err := model.Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, err
}

func (r *ActionRepo) Actions() ([]*model.Action, error) {
	var activityActions []*model.Action
	model := r.DB.Model(&activityActions)

	err := model.Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, err
}
