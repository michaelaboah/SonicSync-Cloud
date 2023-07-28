package model_test

import (
	"testing"

	"github.com/michaelaboah/sonic-sync-cloud/graph/model"
	// "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMatchDetails(t *testing.T) {

  consoleDetails := model.Console{
  	TotalInputs:        0,
  	TotalOutputs:       0,
  	TotalBusses:        0,
  	PhysicalInputs:     0,
  	PhysicalOutputs:    0,
  	AuxInputs:          0,
  	PhysicalAuxInputs:  0,
  	PhantomPowerInputs: 0,
  	Faders:             0,
  	Motorized:          false,
  	Midi:               "",
  	ProtocolInputs:     0,
  	SignalProtocol:     "",
  	CanExpand:          false,
  	MaxSampleRate:      "",
  	Power:              &model.Power{},
  }

  bytes, err := bson.Marshal(consoleDetails); if err != nil {
    t.Fatal(err)
  }


  deets, err := model.MatchDetails(model.CategoryConsole, bytes); if err != nil {
    t.Fatal(err)
  }

  t.Log(deets)
  // assert.Equal(t, consoleDetails.TotalInputs, deets, "Simple")
  
} 
