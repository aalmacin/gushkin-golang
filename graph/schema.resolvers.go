package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aalmacin/gushkin-golang/graph/generated"
	"github.com/aalmacin/gushkin-golang/graph/model"
)

func (r *mutationResolver) CreateWishItem(ctx context.Context, input model.NewWishItemInput) (*model.WishItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateWishItem(ctx context.Context, input model.UpdateWishItemInput) (*model.WishItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateActivity(ctx context.Context, input model.CreateActivityInput) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PerformActivity(ctx context.Context, input model.PerformActivityInput) (*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetWishItemsForUser(ctx context.Context, input model.GetWishItemsForUser) ([]*model.WishItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetActivitiesForUser(ctx context.Context, userID string) ([]*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTodaysActivities(ctx context.Context, userID string) ([]*model.Activity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCurrentFunds(ctx context.Context, userID string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetActivityActionCount(ctx context.Context, userID string) ([]*model.ActivityActionCount, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
