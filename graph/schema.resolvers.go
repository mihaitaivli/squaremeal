package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mihaitaivli/squaremeal/graph/generated"
	"github.com/mihaitaivli/squaremeal/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUser) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMeal(ctx context.Context, input model.NewMeal) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditMeal(ctx context.Context, input model.EditMeal) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMeal(ctx context.Context, mealID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSession(ctx context.Context, input model.NewSession) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSession(ctx context.Context, sessionID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMealsToSession(ctx context.Context, input model.MealsToSessionMutation) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveMealsFromSession(ctx context.Context, input model.MealsToSessionMutation) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VoteUp(ctx context.Context, input model.VoteMutation) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VoteDown(ctx context.Context, input model.VoteMutation) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) InviteUserToSession(ctx context.Context, input *model.InviteUsersToSession) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meal(ctx context.Context) (*model.Meal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meals(ctx context.Context) ([]*model.Meal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Session(ctx context.Context) (*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
