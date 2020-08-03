package repos

import (
	"fmt"

	"github.com/aalmacin/gushkin-golang/graph/model"
	"github.com/go-pg/pg"
)

type ActionRepo struct {
	DB *pg.DB
}

func (r *ActionRepo) actionsByField(field string, value interface{}, userID string) ([]*model.Action, error) {
	var activityActions []*model.Action
	err := r.DB.Model(&activityActions).Where(fmt.Sprintf("%s = ? AND user_id = ?", field), value, userID).Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, nil
}

func (r *ActionRepo) ActionsByActivityId(id string, userId string) ([]*model.Action, error) {
	return r.actionsByField("activity_id", id, userId)
}

func (r *ActionRepo) ActionsWithOptions(options *model.GetActionInput, userID string) ([]*model.Action, error) {
	var activityActions []*model.Action
	model := r.DB.Model(&activityActions)

	model.Where("user_id = ?", userID)

	if *options.Today {
		model.Where(fmt.Sprintf("action_timestamp > current_date"))
	}

	err := model.Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, err
}

func (r *ActionRepo) Actions(userID string) ([]*model.Action, error) {
	var activityActions []*model.Action
	model := r.DB.Model(&activityActions).Where("user_id = ?", userID)

	err := model.Select()
	if err != nil {
		fmt.Println("repos.ActionRepo", err)
		return nil, err
	}

	return activityActions, err
}

func (r *ActionRepo) Create(input model.PerformActivityInput, userID string) (*model.Action, error) {
	action := &model.Action{
		ActivityID: fmt.Sprintf("%v", input.ActivityID),
		UserID:     userID,
	}
	_, err := r.DB.Model(action).Returning("*").Insert()

	return action, err
}
