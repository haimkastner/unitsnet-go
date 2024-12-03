// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LevelUnits enumeration
type LevelUnits string

const (
    
        // 
        LevelDecibel LevelUnits = "Decibel"
        // 
        LevelNeper LevelUnits = "Neper"
)

// LevelDto represents an Level
type LevelDto struct {
	Value float64
	Unit  LevelUnits
}

// LevelDtoFactory struct to group related functions
type LevelDtoFactory struct{}

func (udf LevelDtoFactory) FromJSON(data []byte) (*LevelDto, error) {
	a := LevelDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LevelDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Level struct
type Level struct {
	value       float64
    
    decibelsLazy *float64 
    nepersLazy *float64 
}

// LevelFactory struct to group related functions
type LevelFactory struct{}

func (uf LevelFactory) CreateLevel(value float64, unit LevelUnits) (*Level, error) {
	return newLevel(value, unit)
}

func (uf LevelFactory) FromDto(dto LevelDto) (*Level, error) {
	return newLevel(dto.Value, dto.Unit)
}

func (uf LevelFactory) FromDtoJSON(data []byte) (*Level, error) {
	unitDto, err := LevelDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LevelFactory{}.FromDto(*unitDto)
}


// FromDecibel creates a new Level instance from Decibel.
func (uf LevelFactory) FromDecibels(value float64) (*Level, error) {
	return newLevel(value, LevelDecibel)
}

// FromNeper creates a new Level instance from Neper.
func (uf LevelFactory) FromNepers(value float64) (*Level, error) {
	return newLevel(value, LevelNeper)
}




// newLevel creates a new Level.
func newLevel(value float64, fromUnit LevelUnits) (*Level, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Level{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Level in Decibel.
func (a *Level) BaseValue() float64 {
	return a.value
}


// Decibel returns the value in Decibel.
func (a *Level) Decibels() float64 {
	if a.decibelsLazy != nil {
		return *a.decibelsLazy
	}
	decibels := a.convertFromBase(LevelDecibel)
	a.decibelsLazy = &decibels
	return decibels
}

// Neper returns the value in Neper.
func (a *Level) Nepers() float64 {
	if a.nepersLazy != nil {
		return *a.nepersLazy
	}
	nepers := a.convertFromBase(LevelNeper)
	a.nepersLazy = &nepers
	return nepers
}


// ToDto creates an LevelDto representation.
func (a *Level) ToDto(holdInUnit *LevelUnits) LevelDto {
	if holdInUnit == nil {
		defaultUnit := LevelDecibel // Default value
		holdInUnit = &defaultUnit
	}

	return LevelDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LevelDto representation.
func (a *Level) ToDtoJSON(holdInUnit *LevelUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Level to a specific unit value.
func (a *Level) Convert(toUnit LevelUnits) float64 {
	switch toUnit { 
    case LevelDecibel:
		return a.Decibels()
    case LevelNeper:
		return a.Nepers()
	default:
		return 0
	}
}

func (a *Level) convertFromBase(toUnit LevelUnits) float64 {
    value := a.value
	switch toUnit { 
	case LevelDecibel:
		return (value) 
	case LevelNeper:
		return (0.115129254 * value) 
	default:
		return math.NaN()
	}
}

func (a *Level) convertToBase(value float64, fromUnit LevelUnits) float64 {
	switch fromUnit { 
	case LevelDecibel:
		return (value) 
	case LevelNeper:
		return ((1 / 0.115129254) * value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Level) String() string {
	return a.ToString(LevelDecibel, 2)
}

// ToString formats the Level to string.
// fractionalDigits -1 for not mention
func (a *Level) ToString(unit LevelUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Level) getUnitAbbreviation(unit LevelUnits) string {
	switch unit { 
	case LevelDecibel:
		return "dB" 
	case LevelNeper:
		return "Np" 
	default:
		return ""
	}
}

// Check if the given Level are equals to the current Level
func (a *Level) Equals(other *Level) bool {
	return a.value == other.BaseValue()
}

// Check if the given Level are equals to the current Level
func (a *Level) CompareTo(other *Level) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Level to the current Level.
func (a *Level) Add(other *Level) *Level {
	return &Level{value: a.value + other.BaseValue()}
}

// Subtract the given Level to the current Level.
func (a *Level) Subtract(other *Level) *Level {
	return &Level{value: a.value - other.BaseValue()}
}

// Multiply the given Level to the current Level.
func (a *Level) Multiply(other *Level) *Level {
	return &Level{value: a.value * other.BaseValue()}
}

// Divide the given Level to the current Level.
func (a *Level) Divide(other *Level) *Level {
	return &Level{value: a.value / other.BaseValue()}
}