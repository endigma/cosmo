package subgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/products/subgraph/generated"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/products/subgraph/model"
)

// URL is the resolver for the url field.
func (r *documentationResolver) URL(ctx context.Context, obj *model.Documentation, product model.ProductName) (string, error) {
	return documentationURL(product), nil
}

// Urls is the resolver for the urls field.
func (r *documentationResolver) Urls(ctx context.Context, obj *model.Documentation, products []model.ProductName) ([]string, error) {
	urls := make([]string, 0, 8)
	for _, product := range products {
		urls = append(urls, documentationURL(product))
	}
	return urls, nil
}

// AddFact is the resolver for the addFact field.
func (r *mutationResolver) AddFact(ctx context.Context, fact model.TopSecretFactInput) (model.TopSecretFact, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	switch fact.FactType {
	case model.TopSecretFactTypeDirective:
		fact := model.DirectiveFact{
			Title:       fact.Title,
			Description: fact.Description,
			FactType:    &topSecretFactTypeDirective,
		}
		topSecretFederationFacts = append(topSecretFederationFacts, fact)
		return fact, nil
	case model.TopSecretFactTypeEntity:
		fact := model.EntityFact{
			Title:       fact.Title,
			Description: fact.Description,
			FactType:    &topSecretFactTypeEntity,
		}
		topSecretFederationFacts = append(topSecretFederationFacts, fact)
		return fact, nil
	case model.TopSecretFactTypeMiscellaneous:
		fact := model.MiscellaneousFact{
			Title:       fact.Title,
			Description: fact.Description,
			FactType:    &topSecretFactTypeMiscellaneous,
		}
		topSecretFederationFacts = append(topSecretFederationFacts, fact)
		return fact, nil
	default:
		return nil, errors.New("unknown fact type")
	}
}

// ProductTypes is the resolver for the productTypes field.
func (r *queriesResolver) ProductTypes(ctx context.Context) ([]model.Products, error) {
	fmt.Println("error resolving field employees ----------------------------------------")

	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: fmt.Sprintf("error resolving field %s", "employees"),
		Extensions: map[string]interface{}{
			"code": "ERROR_CODE",
			"foo":  "bar",
		},
	})

	return nil, nil
}

// TopSecretFederationFacts is the resolver for the topSecretFederationFacts field.
func (r *queriesResolver) TopSecretFederationFacts(ctx context.Context) ([]model.TopSecretFact, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return topSecretFederationFacts, nil
}

// FactTypes is the resolver for the factTypes field.
func (r *queriesResolver) FactTypes(ctx context.Context) ([]model.TopSecretFactType, error) {
	return model.AllTopSecretFactType, nil
}

// Documentation returns generated.DocumentationResolver implementation.
func (r *Resolver) Documentation() generated.DocumentationResolver { return &documentationResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Queries returns generated.QueriesResolver implementation.
func (r *Resolver) Queries() generated.QueriesResolver { return &queriesResolver{r} }

type documentationResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queriesResolver struct{ *Resolver }
