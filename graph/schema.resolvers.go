package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/aalmacin/gushkin-golang/graph/generated"
	"github.com/aalmacin/gushkin-golang/graph/model"
)

func (r *actionResolver) ActionTimestamp(ctx context.Context, obj *model.Action) (*time.Time, error) {
	return &obj.ActionTimestamp, nil
}

func (r *actionResolver) Activity(ctx context.Context, obj *model.Action) (*model.Activity, error) {
	return r.ActivityRepo.ActivityById(obj.ActivityID)
}

func (r *activityResolver) Actions(ctx context.Context, obj *model.Activity) ([]*model.Action, error) {
	return r.ActionRepo.ActionsByActivityId(obj.ID)
}

func (r *mutationResolver) CreateWish(ctx context.Context, input model.NewWishInput) (*model.Wish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWish(ctx context.Context, input model.UpdateWishInput) (*model.Wish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateActivity(ctx context.Context, input model.CreateActivityInput) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PerformActivity(ctx context.Context, input model.PerformActivityInput) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Wishes(ctx context.Context, input *model.GetWishInput) ([]*model.Wish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Activities(ctx context.Context) ([]*model.Activity, error) {
	return r.ActivityRepo.Activities()
}

func (r *queryResolver) TodaysActivities(ctx context.Context) ([]*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CurrentFunds(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ActionCount(ctx context.Context) ([]*model.ActionCount, error) {
	panic(fmt.Errorf("not implemented"))
}

// Action returns generated.ActionResolver implementation.
func (r *Resolver) Action() generated.ActionResolver { return &actionResolver{r} }

// Activity returns generated.ActivityResolver implementation.
func (r *Resolver) Activity() generated.ActivityResolver { return &activityResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type actionResolver struct{ *Resolver }
type activityResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
