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

// Fungsi untuk menambahkan new Movie
func (r *mutationResolver) AddMovie(ctx context.Context, input model.InputMovie) (*model.Movie, error) {
	// variable stars digunakan untuk menampung data stars dari input-an
	var stars []*model.Star
	// simpan seluruh data star ke dalam variable stars
	for _, star := range input.Stars {
		stars = append(stars, &model.Star{
			Name: star.Name,
		})
	}
	// variable movie digunakan untuk menampung data movie dari input-an
	movie := &model.Movie{
		ID:    input.ID,
		Title: input.Title,
		Year:  input.Year,
		Stars: stars,
	}
	// data disimpan ke dalam database, jika error maka tampilkan pesan error
	if err := movieRepo.Save(movie); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	// kembalikan data movie
	return movie, nil
}

// Fungsi untuk edit data movie
func (r *mutationResolver) UpdateMovie(ctx context.Context, id int, input model.InputMovie) (*model.Movie, error) {
	// variable stars digunakan untuk menampung data stars dari input-an
	var stars []*model.Star
	// simpan seluruh data star yang baru ke dalam variable stars
	for _, star := range input.Stars {
		stars = append(stars, &model.Star{
			Name: star.Name,
		})
	}
	// variable newMovie digunakan untuk menampung data movie dari input-an
	newMovie := &model.Movie{
		ID:    input.ID,
		Title: input.Title,
		Year:  input.Year,
		Stars: stars,
	}
	// edit data yang ada pada database dengan data yang baru
	data, err := movieRepo.Edit(id, newMovie)
	// jika terjadi error maka kembalikan pesan error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	// kembbalikan data new movie
	return data, nil
}

func (r *mutationResolver) DeleteMovie(ctx context.Context, id int) (string, error) {
	numberOfDel, err := movieRepo.Delete(id)
	if err != nil {
		return "false", fmt.Errorf("internal server error")
	}
	return fmt.Sprintf("Number of documents deleted: %d", numberOfDel), nil
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
	movie, err := movieRepo.Get(id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}
	return movie, nil
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
