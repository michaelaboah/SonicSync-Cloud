package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"
	"log"

	pkgDB "github.com/michaelaboah/sonic-sync-cloud/database"
	"github.com/michaelaboah/sonic-sync-cloud/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	users := r.DB.Database(pkgDB.ClientsDB).Collection(pkgDB.UserCol)

	user := &model.User{
		ID:    "0",
		Name:  input.Name,
		Email: input.Email,
	}

	users.InsertOne(ctx, user)
	return user, nil
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.ItemInput, details *model.CategoryDetailsInput) (*model.Item, error) {
	var cat_details model.CategoryDetails

	switch input.Category {
	case model.CategoryConsole:
		cat_details = details.ConsoleInput
	}

	items := r.DB.Database(pkgDB.EquipDB).Collection(pkgDB.ItemsCol)
	item := &model.Item{
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
		Cost:         input.Cost,
		Model:        input.Model,
		Weight:       input.Weight,
		Manufacturer: input.Manufacturer,
		Category:     input.Category,
		Details:      cat_details,
		Notes:        &input.Model,
		Dimensions:   (*model.Dimension)(input.Dimensions),
		PDFBlob:      input.PDFBlob,
	}
	items.InsertOne(ctx, item)

	return item, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	usersCollection := r.DB.Database(pkgDB.ClientsDB).Collection(pkgDB.UserCol)

	usersCursor, err := usersCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	fmt.Println(usersCursor)

	return nil, nil
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	itemsCollection := r.DB.Database(pkgDB.EquipDB).Collection(pkgDB.ItemsCol)
	itemsCursor, err := itemsCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var results []*model.Item

	for itemsCursor.Next(ctx) {
		var (
			doc  bson.M
			item *model.Item
			err  error
		)

		itemsCursor.Decode(&doc)

		// Should be a better way to do this without needing to marshal to itemBytes and back
		itemBytes, err := bson.Marshal(doc)
		if err != nil {
			log.Println("Error Marshaling BSON to bytes: ", err)
		}

		detailsBytes, _ := bson.Marshal(doc["details"])

		err = bson.Unmarshal(itemBytes, &item)
		if err != nil {
			log.Println(err)
		}

		details, err := model.MatchDetails(item.Category, detailsBytes)
		if err != nil {
			log.Println("Error Unmarshaling bytes", err)
		}

		item.Details = details

		results = append(results, item)

	}
	fmt.Println("Number of Items: ", len(results))
	return results, nil
}

// FindByModel is the resolver for the find_by_model field.
func (r *queryResolver) FindByModel(ctx context.Context, modelName string) (*model.Item, error) {
	itemsCollection := r.DB.Database(pkgDB.EquipDB).Collection(pkgDB.ItemsCol)
	itemResult := itemsCollection.FindOne(ctx, bson.M{"model": modelName})

	var doc bson.M
	var item *model.Item
	var err error

	itemResult.Decode(&doc)

	itemBytes, err := bson.Marshal(doc)
	if err != nil {
		log.Println("Error Marshaling BSON to bytes: ", err)
	}

	detailsBytes, err := bson.Marshal(doc["details"])
	if err != nil {
		log.Println("Error Marshal 'details' from mongo document: ", err)
	}

	err = bson.Unmarshal(itemBytes, &item)
	if err != nil {
		log.Println(err)
	}

	details, err := model.MatchDetails(item.Category, detailsBytes)
	if err != nil {
		log.Println("Error Unmarshaling bytes", err)
	}

	item.Details = details

	return item, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) FindByID(ctx context.Context) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: FindByID - find_by_id"))
}
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.ItemInput) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: UpdateItem - updateItem"))
}
