// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LevelUnits defines various units of Level.
type LevelUnits string

const (
    
        // 
        LevelDecibel LevelUnits = "Decibel"
        // 
        LevelNeper LevelUnits = "Neper"
)

// LevelDto represents a Level measurement with a numerical value and its corresponding unit.
type LevelDto struct {
    // Value is the numerical representation of the Level.
	Value float64
    // Unit specifies the unit of measurement for the Level, as defined in the LevelUnits enumeration.
	Unit  LevelUnits
}

// LevelDtoFactory groups methods for creating and serializing LevelDto objects.
type LevelDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LevelDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LevelDtoFactory) FromJSON(data []byte) (*LevelDto, error) {
	a := LevelDto{}

    // Parse JSON into LevelDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LevelDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LevelDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Level represents a measurement in a of Level.
//
// Level is the logarithm of the ratio of a quantity Q to a reference value of that quantity, Q₀, expressed in dimensionless units.
type Level struct {
	// value is the base measurement stored internally.
	value       float64
    
    decibelsLazy *float64 
    nepersLazy *float64 
}

// LevelFactory groups methods for creating Level instances.
type LevelFactory struct{}

// CreateLevel creates a new Level instance from the given value and unit.
func (uf LevelFactory) CreateLevel(value float64, unit LevelUnits) (*Level, error) {
	return newLevel(value, unit)
}

// FromDto converts a LevelDto to a Level instance.
func (uf LevelFactory) FromDto(dto LevelDto) (*Level, error) {
	return newLevel(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Level instance.
func (uf LevelFactory) FromDtoJSON(data []byte) (*Level, error) {
	unitDto, err := LevelDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LevelDto from JSON: %w", err)
	}
	return LevelFactory{}.FromDto(*unitDto)
}


// FromDecibels creates a new Level instance from a value in Decibels.
func (uf LevelFactory) FromDecibels(value float64) (*Level, error) {
	return newLevel(value, LevelDecibel)
}

// FromNepers creates a new Level instance from a value in Nepers.
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

// BaseValue returns the base value of Level in Decibel unit (the base/default quantity).
func (a *Level) BaseValue() float64 {
	return a.value
}


// Decibels returns the Level value in Decibels.
//
// 
func (a *Level) Decibels() float64 {
	if a.decibelsLazy != nil {
		return *a.decibelsLazy
	}
	decibels := a.convertFromBase(LevelDecibel)
	a.decibelsLazy = &decibels
	return decibels
}

// Nepers returns the Level value in Nepers.
//
// 
func (a *Level) Nepers() float64 {
	if a.nepersLazy != nil {
		return *a.nepersLazy
	}
	nepers := a.convertFromBase(LevelNeper)
	a.nepersLazy = &nepers
	return nepers
}


// ToDto creates a LevelDto representation from the Level instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Decibel by default.
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

// ToDtoJSON creates a JSON representation of the Level instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Decibel by default.
func (a *Level) ToDtoJSON(holdInUnit *LevelUnits) ([]byte, error) {
	// Convert to LevelDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Level to a specific unit value.
// The function uses the provided unit type (LevelUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Level) Convert(toUnit LevelUnits) float64 {
	switch toUnit { 
    case LevelDecibel:
		return a.Decibels()
    case LevelNeper:
		return a.Nepers()
	default:
		return math.NaN()
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

// String returns a string representation of the Level in the default unit (Decibel),
// formatted to two decimal places.
func (a Level) String() string {
	return a.ToString(LevelDecibel, 2)
}

// ToString formats the Level value as a string with the specified unit and fractional digits.
// It converts the Level to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Level value will be converted (e.g., Decibel))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Level with the unit abbreviation.
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

// Equals checks if the given Level is equal to the current Level.
//
// Parameters:
//    other: The Level to compare against.
//
// Returns:
//    bool: Returns true if both Level are equal, false otherwise.
func (a *Level) Equals(other *Level) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Level with another Level.
// It returns -1 if the current Level is less than the other Level, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Level to compare against.
//
// Returns:
//    int: -1 if the current Level is less, 1 if greater, and 0 if equal.
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

// Add adds the given Level to the current Level and returns the result.
// The result is a new Level instance with the sum of the values.
//
// Parameters:
//    other: The Level to add to the current Level.
//
// Returns:
//    *Level: A new Level instance representing the sum of both Level.
func (a *Level) Add(other *Level) *Level {
	return &Level{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Level from the current Level and returns the result.
// The result is a new Level instance with the difference of the values.
//
// Parameters:
//    other: The Level to subtract from the current Level.
//
// Returns:
//    *Level: A new Level instance representing the difference of both Level.
func (a *Level) Subtract(other *Level) *Level {
	return &Level{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Level by the given Level and returns the result.
// The result is a new Level instance with the product of the values.
//
// Parameters:
//    other: The Level to multiply with the current Level.
//
// Returns:
//    *Level: A new Level instance representing the product of both Level.
func (a *Level) Multiply(other *Level) *Level {
	return &Level{value: a.value * other.BaseValue()}
}

// Divide divides the current Level by the given Level and returns the result.
// The result is a new Level instance with the quotient of the values.
//
// Parameters:
//    other: The Level to divide the current Level by.
//
// Returns:
//    *Level: A new Level instance representing the quotient of both Level.
func (a *Level) Divide(other *Level) *Level {
	return &Level{value: a.value / other.BaseValue()}
}