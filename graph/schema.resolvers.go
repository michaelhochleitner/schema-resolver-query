package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"schema-resolver-query/graph/generated"
	"schema-resolver-query/graph/model"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, param string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
