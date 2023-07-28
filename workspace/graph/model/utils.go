package model

import "go.mongodb.org/mongo-driver/bson"




func MatchDetails(category Category, detailsBytes []byte) (CategoryDetails, error) {
  var err error
  switch category {
  case CategoryConsole: 
    var console *Console 
    
    err = bson.Unmarshal(detailsBytes, &console); if err != nil {
      return nil, err
    }

    return console, nil

  case CategoryAmplifier: 
    // var amplifier *Amplifier


  } 

  return nil, nil
}



