package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"schema-resolver-query/graph/generated"
	"schema-resolver-query/graph/model"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
