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

func checkUserID(userID string) error {
	if userID == "" {
		fmt.Println("UserID not found")
		return fmt.Errorf("Something went wrong")
	}
	return nil
}

func (r *actionResolver) Activity(ctx context.Context, obj *model.Action) (*model.Activity, error) {
	activityLoader := dataloaders.GetActivityLoader(ctx)
	activity, err := activityLoader.Load(obj.ActivityID)
	if err != nil {
		fmt.Println("Error in action resolver: Activity", err)
	}
	return activity, err
}

func (r *activityResolver) Actions(ctx context.Context, obj *model.Activity) ([]*model.Action, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	return r.ActionRepo.ActionsByActivityId(obj.ID, r.UserID)
}

func (r *mutationResolver) CreateWish(ctx context.Context, input model.NewWishInput) (*model.Wish, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	return r.WishRepo.Create(input, r.UserID)
}

func (r *mutationResolver) UpdateWish(ctx context.Context, input model.UpdateWishInput) (*model.Wish, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	return r.WishRepo.Update(input, r.UserID)
}

func (r *mutationResolver) CreateActivity(ctx context.Context, input model.NewActivityInput) (*model.Activity, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}

	return r.ActivityRepo.Create(input, r.UserID)
}

func (r *mutationResolver) PerformActivity(ctx context.Context, input model.PerformActivityInput) (*model.Action, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}

	_, err = r.ActivityRepo.ActivityById(fmt.Sprintf("%v", input.ActivityID), r.UserID)

	if err != nil {
		return nil, fmt.Errorf("Invalid Activity")
	}
	return r.ActionRepo.Create(input, r.UserID)
}

func (r *queryResolver) Wishes(ctx context.Context, input *model.GetWishInput) ([]*model.Wish, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	return r.WishRepo.Wishes(r.UserID)
}

func (r *queryResolver) Activities(ctx context.Context) ([]*model.Activity, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	return r.ActivityRepo.Activities(r.UserID)
}

func (r *queryResolver) Actions(ctx context.Context, input *model.GetActionInput) ([]*model.Action, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
	if input != nil {
		return r.ActionRepo.ActionsWithOptions(input, r.UserID)
	}
	return r.ActionRepo.Actions(r.UserID)
}

func (r *queryResolver) CurrentFunds(ctx context.Context) (int, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return 0, err
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ActionCount(ctx context.Context) ([]*model.ActionCount, error) {
	err := checkUserID(r.UserID)
	if err != nil {
		return nil, err
	}
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
