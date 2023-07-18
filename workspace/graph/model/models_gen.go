// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Dimension struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Item struct {
	ID           string     `json:"id"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	Cost         float64    `json:"cost"`
	Model        string     `json:"model"`
	Weight       float64    `json:"weight"`
	Manufacturer string     `json:"manufacturer"`
	Category     Category   `json:"category"`
	Notes        *string    `json:"notes,omitempty"`
	Dimensions   *Dimension `json:"dimensions,omitempty"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Category string

const (
	CategoryConsole     Category = "CONSOLE"
	CategoryProcessor   Category = "PROCESSOR"
	CategoryMonitoring  Category = "MONITORING"
	CategorySpeaker     Category = "SPEAKER"
	CategoryAmplifier   Category = "AMPLIFIER"
	CategoryComputer    Category = "COMPUTER"
	CategoryNetwork     Category = "NETWORK"
	CategoryRadio       Category = "RADIO"
	CategoryMicrophones Category = "MICROPHONES"
	CategorySpkHardware Category = "SPK_HARDWARE"
	CategoryGeneric     Category = "GENERIC"
)

var AllCategory = []Category{
	CategoryConsole,
	CategoryProcessor,
	CategoryMonitoring,
	CategorySpeaker,
	CategoryAmplifier,
	CategoryComputer,
	CategoryNetwork,
	CategoryRadio,
	CategoryMicrophones,
	CategorySpkHardware,
	CategoryGeneric,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryConsole, CategoryProcessor, CategoryMonitoring, CategorySpeaker, CategoryAmplifier, CategoryComputer, CategoryNetwork, CategoryRadio, CategoryMicrophones, CategorySpkHardware, CategoryGeneric:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
