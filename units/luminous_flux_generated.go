// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LuminousFluxUnits defines various units of LuminousFlux.
type LuminousFluxUnits string

const (
    
        // 
        LuminousFluxLumen LuminousFluxUnits = "Lumen"
)

// LuminousFluxDto represents a LuminousFlux measurement with a numerical value and its corresponding unit.
type LuminousFluxDto struct {
    // Value is the numerical representation of the LuminousFlux.
	Value float64
    // Unit specifies the unit of measurement for the LuminousFlux, as defined in the LuminousFluxUnits enumeration.
	Unit  LuminousFluxUnits
}

// LuminousFluxDtoFactory groups methods for creating and serializing LuminousFluxDto objects.
type LuminousFluxDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LuminousFluxDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LuminousFluxDtoFactory) FromJSON(data []byte) (*LuminousFluxDto, error) {
	a := LuminousFluxDto{}

    // Parse JSON into LuminousFluxDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LuminousFluxDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LuminousFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// LuminousFlux represents a measurement in a of LuminousFlux.
//
// In photometry, luminous flux or luminous power is the measure of the perceived power of light.
type LuminousFlux struct {
	// value is the base measurement stored internally.
	value       float64
    
    lumensLazy *float64 
}

// LuminousFluxFactory groups methods for creating LuminousFlux instances.
type LuminousFluxFactory struct{}

// CreateLuminousFlux creates a new LuminousFlux instance from the given value and unit.
func (uf LuminousFluxFactory) CreateLuminousFlux(value float64, unit LuminousFluxUnits) (*LuminousFlux, error) {
	return newLuminousFlux(value, unit)
}

// FromDto converts a LuminousFluxDto to a LuminousFlux instance.
func (uf LuminousFluxFactory) FromDto(dto LuminousFluxDto) (*LuminousFlux, error) {
	return newLuminousFlux(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a LuminousFlux instance.
func (uf LuminousFluxFactory) FromDtoJSON(data []byte) (*LuminousFlux, error) {
	unitDto, err := LuminousFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LuminousFluxDto from JSON: %w", err)
	}
	return LuminousFluxFactory{}.FromDto(*unitDto)
}


// FromLumens creates a new LuminousFlux instance from a value in Lumens.
func (uf LuminousFluxFactory) FromLumens(value float64) (*LuminousFlux, error) {
	return newLuminousFlux(value, LuminousFluxLumen)
}


// newLuminousFlux creates a new LuminousFlux.
func newLuminousFlux(value float64, fromUnit LuminousFluxUnits) (*LuminousFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LuminousFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LuminousFlux in Lumen unit (the base/default quantity).
func (a *LuminousFlux) BaseValue() float64 {
	return a.value
}


// Lumens returns the LuminousFlux value in Lumens.
//
// 
func (a *LuminousFlux) Lumens() float64 {
	if a.lumensLazy != nil {
		return *a.lumensLazy
	}
	lumens := a.convertFromBase(LuminousFluxLumen)
	a.lumensLazy = &lumens
	return lumens
}


// ToDto creates a LuminousFluxDto representation from the LuminousFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Lumen by default.
func (a *LuminousFlux) ToDto(holdInUnit *LuminousFluxUnits) LuminousFluxDto {
	if holdInUnit == nil {
		defaultUnit := LuminousFluxLumen // Default value
		holdInUnit = &defaultUnit
	}

	return LuminousFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the LuminousFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Lumen by default.
func (a *LuminousFlux) ToDtoJSON(holdInUnit *LuminousFluxUnits) ([]byte, error) {
	// Convert to LuminousFluxDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a LuminousFlux to a specific unit value.
// The function uses the provided unit type (LuminousFluxUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *LuminousFlux) Convert(toUnit LuminousFluxUnits) float64 {
	switch toUnit { 
    case LuminousFluxLumen:
		return a.Lumens()
	default:
		return math.NaN()
	}
}

func (a *LuminousFlux) convertFromBase(toUnit LuminousFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminousFluxLumen:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *LuminousFlux) convertToBase(value float64, fromUnit LuminousFluxUnits) float64 {
	switch fromUnit { 
	case LuminousFluxLumen:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the LuminousFlux in the default unit (Lumen),
// formatted to two decimal places.
func (a LuminousFlux) String() string {
	return a.ToString(LuminousFluxLumen, 2)
}

// ToString formats the LuminousFlux value as a string with the specified unit and fractional digits.
// It converts the LuminousFlux to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the LuminousFlux value will be converted (e.g., Lumen))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the LuminousFlux with the unit abbreviation.
func (a *LuminousFlux) ToString(unit LuminousFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LuminousFlux) getUnitAbbreviation(unit LuminousFluxUnits) string {
	switch unit { 
	case LuminousFluxLumen:
		return "lm" 
	default:
		return ""
	}
}

// Equals checks if the given LuminousFlux is equal to the current LuminousFlux.
//
// Parameters:
//    other: The LuminousFlux to compare against.
//
// Returns:
//    bool: Returns true if both LuminousFlux are equal, false otherwise.
func (a *LuminousFlux) Equals(other *LuminousFlux) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current LuminousFlux with another LuminousFlux.
// It returns -1 if the current LuminousFlux is less than the other LuminousFlux, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The LuminousFlux to compare against.
//
// Returns:
//    int: -1 if the current LuminousFlux is less, 1 if greater, and 0 if equal.
func (a *LuminousFlux) CompareTo(other *LuminousFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given LuminousFlux to the current LuminousFlux and returns the result.
// The result is a new LuminousFlux instance with the sum of the values.
//
// Parameters:
//    other: The LuminousFlux to add to the current LuminousFlux.
//
// Returns:
//    *LuminousFlux: A new LuminousFlux instance representing the sum of both LuminousFlux.
func (a *LuminousFlux) Add(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given LuminousFlux from the current LuminousFlux and returns the result.
// The result is a new LuminousFlux instance with the difference of the values.
//
// Parameters:
//    other: The LuminousFlux to subtract from the current LuminousFlux.
//
// Returns:
//    *LuminousFlux: A new LuminousFlux instance representing the difference of both LuminousFlux.
func (a *LuminousFlux) Subtract(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current LuminousFlux by the given LuminousFlux and returns the result.
// The result is a new LuminousFlux instance with the product of the values.
//
// Parameters:
//    other: The LuminousFlux to multiply with the current LuminousFlux.
//
// Returns:
//    *LuminousFlux: A new LuminousFlux instance representing the product of both LuminousFlux.
func (a *LuminousFlux) Multiply(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value * other.BaseValue()}
}

// Divide divides the current LuminousFlux by the given LuminousFlux and returns the result.
// The result is a new LuminousFlux instance with the quotient of the values.
//
// Parameters:
//    other: The LuminousFlux to divide the current LuminousFlux by.
//
// Returns:
//    *LuminousFlux: A new LuminousFlux instance representing the quotient of both LuminousFlux.
func (a *LuminousFlux) Divide(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value / other.BaseValue()}
}