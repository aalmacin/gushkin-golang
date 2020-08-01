package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/aalmacin/gushkin-golang/dataloaders"
	"github.com/aalmacin/gushkin-golang/graph/generated"
	"github.com/aalmacin/gushkin-golang/graph/model"
)

func (r *actionResolver) Activity(ctx context.Context, obj *model.Action) (*model.Activity, error) {
	activityLoader := dataloaders.GetActivityLoader(ctx)
	activity, err := activityLoader.Load(obj.ActivityID)
	if err != nil {
		fmt.Println("Error in action resolver: Activity", err)
	}
	return activity, err
}

func (r *activityResolver) Actions(ctx context.Context, obj *model.Activity) ([]*model.Action, error) {
	return r.ActionRepo.ActionsByActivityId(obj.ID)
}

func (r *mutationResolver) CreateWish(ctx context.Context, input model.NewWishInput) (*model.Wish, error) {
	return r.WishRepo.Create(input)
}

func (r *mutationResolver) UpdateWish(ctx context.Context, input model.UpdateWishInput) (*model.Wish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateActivity(ctx context.Context, input model.NewActivityInput) (*model.Activity, error) {
	return r.ActivityRepo.Create(input)
}

func (r *mutationResolver) PerformActivity(ctx context.Context, input model.PerformActivityInput) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Wishes(ctx context.Context, input *model.GetWishInput) ([]*model.Wish, error) {
	return r.WishRepo.Wishes()
}

func (r *queryResolver) Activities(ctx context.Context) ([]*model.Activity, error) {
	return r.ActivityRepo.Activities()
}

func (r *queryResolver) Actions(ctx context.Context, input *model.GetActionInput) ([]*model.Action, error) {
	if input != nil {
		return r.ActionRepo.ActionsWithOptions(input)
	}
	return r.ActionRepo.Actions()
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *actionResolver) ActionTimestamp(ctx context.Context, obj *model.Action) (*time.Time, error) {
	return &obj.ActionTimestamp, nil
}
