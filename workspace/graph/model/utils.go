package model

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)




func MatchDetails(category Category, detailsBytes []byte) (CategoryDetails, error) {
  var err error

  switch category {


    case CategoryConsole: 
      var console Console 
      
      err = bson.Unmarshal(detailsBytes, &console); if err != nil {
        return nil, err
      }

      return console, nil

    case CategoryComputer: 
      var computer Computer
      
      err = bson.Unmarshal(detailsBytes, &computer); if err != nil {
        return nil, err
      }

      return computer, nil

    default:
      return nil, errors.New("Unimplemnted Category" + category.String())
  } 

  return nil, nil
}



