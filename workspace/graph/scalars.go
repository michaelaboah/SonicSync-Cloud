package graph 

import (
	"errors"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarshalObjectID(id primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, `"`+id.Hex()+`"`)
	})
}

func UnmarshalObjectID(v interface{}) (primitive.ObjectID, error) {
	str, ok := v.(string)
	if !ok {
		return primitive.ObjectID{}, errors.New("ObjectID must be a string")
	}
	id, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id, nil
}
