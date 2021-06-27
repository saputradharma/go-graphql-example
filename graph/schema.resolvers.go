package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/saputradharma/go-graphql-example/graph/generated"
	"github.com/saputradharma/go-graphql-example/graph/model"
	"github.com/saputradharma/go-graphql-example/internal/auth"
	"github.com/saputradharma/go-graphql-example/internal/links"
	"github.com/saputradharma/go-graphql-example/internal/pkg/jwt"
	"github.com/saputradharma/go-graphql-example/internal/users"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link

	user := auth.ForContext(ctx)

	if user != nil {
		log.Println("User unauthenticated.")
		return &model.Link{}, fmt.Errorf("Access Denied.")
	}

	link.Address = input.Address
	link.Title = input.Title
	link.User = user

	id := links.CreateLink(link)

	graphqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}

	return &model.Link{
		ID:      strconv.FormatInt(id, 10),
		Title:   link.Title,
		Address: link.Address,
		User:    graphqlUser,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User

	user.Username = input.Username
	user.Password = input.Password
	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User

	user.Username = input.Username
	user.Password = input.Password

	valid := user.Authenticate()

	if !valid {
		return "", &users.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)

	if err != nil {
		return "", fmt.Errorf("Access Denied.")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var result []*model.Link

	dbLinks := links.GetAll()

	for _, link := range dbLinks {

		grahpqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}

		result = append(
			result,
			&model.Link{
				ID:      link.ID,
				Title:   link.Title,
				Address: link.Address,
				User:    grahpqlUser,
			})
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
