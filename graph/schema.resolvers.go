package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/alfiancikoa/graphql-mongodb/graph/generated"
	"github.com/alfiancikoa/graphql-mongodb/graph/model"
	"github.com/alfiancikoa/graphql-mongodb/repository"
)

func (r *mutationResolver) AddMovie(ctx context.Context, input model.InputMovie) (*model.Movie, error) {
	var stars []*model.Star
	for _, star := range input.Stars {
		stars = append(stars, &model.Star{
			Name: star.Name,
		})
	}

	movie := &model.Movie{
		Title: input.Title,
		Year:  input.Year,
		Stars: stars,
	}
	if err := movieRepo.Save(movie); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	return movie, nil
}

func (r *mutationResolver) UpdateMovie(ctx context.Context, id int, input model.InputMovie) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMovie(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	data, err := movieRepo.GetAll()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	return data, nil
}

func (r *queryResolver) Movie(ctx context.Context, id int) (*model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var movieRepo repository.MovieRepository = repository.New()
