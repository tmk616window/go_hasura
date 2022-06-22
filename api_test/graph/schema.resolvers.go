package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api_test/graph/generated"
	"api_test/graph/model"
	"api_test/graph/models"
	db "api_test/models"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	todo := &models.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	}
	db := db.DbConnect()
	result := db.Create(&todo)
	fmt.Println(result)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*models.Todo, error) {
	db := db.DbConnect()
	todo := models.Todo{}
	db.First(&todo, input.ID)
	todo.Text = input.Text
	db.Save(&todo)
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	db := db.DbConnect()
	db.Find(&todos)
	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	db := db.DbConnect()
	db.Preload("Todos").Find(&users)
	return users, nil
}

func (r *queryResolver) Todo(ctx context.Context, input int) (*models.Todo, error) {
	var todo models.Todo
	db := db.DbConnect()
	db.Preload("User").First(&todo, input)
	return &todo, nil
}

func (r *todoResolver) Done(ctx context.Context, obj *models.Todo) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*model.User, error) {
	var user model.User
	db := db.DbConnect()
	db.First(&user, obj.UserID)
	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *todoResolver) UserID(ctx context.Context, obj *models.Todo) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
