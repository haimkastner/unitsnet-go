// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LuminousIntensityUnits defines various units of LuminousIntensity.
type LuminousIntensityUnits string

const (
    
        // 
        LuminousIntensityCandela LuminousIntensityUnits = "Candela"
)

// LuminousIntensityDto represents a LuminousIntensity measurement with a numerical value and its corresponding unit.
type LuminousIntensityDto struct {
    // Value is the numerical representation of the LuminousIntensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the LuminousIntensity, as defined in the LuminousIntensityUnits enumeration.
	Unit  LuminousIntensityUnits `json:"unit"`
}

// LuminousIntensityDtoFactory groups methods for creating and serializing LuminousIntensityDto objects.
type LuminousIntensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LuminousIntensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LuminousIntensityDtoFactory) FromJSON(data []byte) (*LuminousIntensityDto, error) {
	a := LuminousIntensityDto{}

    // Parse JSON into LuminousIntensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LuminousIntensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LuminousIntensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// LuminousIntensity represents a measurement in a of LuminousIntensity.
//
// In photometry, luminous intensity is a measure of the wavelength-weighted power emitted by a light source in a particular direction per unit solid angle, based on the luminosity function, a standardized model of the sensitivity of the human eye.
type LuminousIntensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    candelaLazy *float64 
}

// LuminousIntensityFactory groups methods for creating LuminousIntensity instances.
type LuminousIntensityFactory struct{}

// CreateLuminousIntensity creates a new LuminousIntensity instance from the given value and unit.
func (uf LuminousIntensityFactory) CreateLuminousIntensity(value float64, unit LuminousIntensityUnits) (*LuminousIntensity, error) {
	return newLuminousIntensity(value, unit)
}

// FromDto converts a LuminousIntensityDto to a LuminousIntensity instance.
func (uf LuminousIntensityFactory) FromDto(dto LuminousIntensityDto) (*LuminousIntensity, error) {
	return newLuminousIntensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a LuminousIntensity instance.
func (uf LuminousIntensityFactory) FromDtoJSON(data []byte) (*LuminousIntensity, error) {
	unitDto, err := LuminousIntensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LuminousIntensityDto from JSON: %w", err)
	}
	return LuminousIntensityFactory{}.FromDto(*unitDto)
}


// FromCandela creates a new LuminousIntensity instance from a value in Candela.
func (uf LuminousIntensityFactory) FromCandela(value float64) (*LuminousIntensity, error) {
	return newLuminousIntensity(value, LuminousIntensityCandela)
}


// newLuminousIntensity creates a new LuminousIntensity.
func newLuminousIntensity(value float64, fromUnit LuminousIntensityUnits) (*LuminousIntensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LuminousIntensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LuminousIntensity in Candela unit (the base/default quantity).
func (a *LuminousIntensity) BaseValue() float64 {
	return a.value
}


// Candela returns the LuminousIntensity value in Candela.
//
// 
func (a *LuminousIntensity) Candela() float64 {
	if a.candelaLazy != nil {
		return *a.candelaLazy
	}
	candela := a.convertFromBase(LuminousIntensityCandela)
	a.candelaLazy = &candela
	return candela
}


// ToDto creates a LuminousIntensityDto representation from the LuminousIntensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Candela by default.
func (a *LuminousIntensity) ToDto(holdInUnit *LuminousIntensityUnits) LuminousIntensityDto {
	if holdInUnit == nil {
		defaultUnit := LuminousIntensityCandela // Default value
		holdInUnit = &defaultUnit
	}

	return LuminousIntensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the LuminousIntensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Candela by default.
func (a *LuminousIntensity) ToDtoJSON(holdInUnit *LuminousIntensityUnits) ([]byte, error) {
	// Convert to LuminousIntensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a LuminousIntensity to a specific unit value.
// The function uses the provided unit type (LuminousIntensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *LuminousIntensity) Convert(toUnit LuminousIntensityUnits) float64 {
	switch toUnit { 
    case LuminousIntensityCandela:
		return a.Candela()
	default:
		return math.NaN()
	}
}

func (a *LuminousIntensity) convertFromBase(toUnit LuminousIntensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminousIntensityCandela:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *LuminousIntensity) convertToBase(value float64, fromUnit LuminousIntensityUnits) float64 {
	switch fromUnit { 
	case LuminousIntensityCandela:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the LuminousIntensity in the default unit (Candela),
// formatted to two decimal places.
func (a LuminousIntensity) String() string {
	return a.ToString(LuminousIntensityCandela, 2)
}

// ToString formats the LuminousIntensity value as a string with the specified unit and fractional digits.
// It converts the LuminousIntensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the LuminousIntensity value will be converted (e.g., Candela))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the LuminousIntensity with the unit abbreviation.
func (a *LuminousIntensity) ToString(unit LuminousIntensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetLuminousIntensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetLuminousIntensityAbbreviation(unit))
}

// Equals checks if the given LuminousIntensity is equal to the current LuminousIntensity.
//
// Parameters:
//    other: The LuminousIntensity to compare against.
//
// Returns:
//    bool: Returns true if both LuminousIntensity are equal, false otherwise.
func (a *LuminousIntensity) Equals(other *LuminousIntensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current LuminousIntensity with another LuminousIntensity.
// It returns -1 if the current LuminousIntensity is less than the other LuminousIntensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The LuminousIntensity to compare against.
//
// Returns:
//    int: -1 if the current LuminousIntensity is less, 1 if greater, and 0 if equal.
func (a *LuminousIntensity) CompareTo(other *LuminousIntensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given LuminousIntensity to the current LuminousIntensity and returns the result.
// The result is a new LuminousIntensity instance with the sum of the values.
//
// Parameters:
//    other: The LuminousIntensity to add to the current LuminousIntensity.
//
// Returns:
//    *LuminousIntensity: A new LuminousIntensity instance representing the sum of both LuminousIntensity.
func (a *LuminousIntensity) Add(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given LuminousIntensity from the current LuminousIntensity and returns the result.
// The result is a new LuminousIntensity instance with the difference of the values.
//
// Parameters:
//    other: The LuminousIntensity to subtract from the current LuminousIntensity.
//
// Returns:
//    *LuminousIntensity: A new LuminousIntensity instance representing the difference of both LuminousIntensity.
func (a *LuminousIntensity) Subtract(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current LuminousIntensity by the given LuminousIntensity and returns the result.
// The result is a new LuminousIntensity instance with the product of the values.
//
// Parameters:
//    other: The LuminousIntensity to multiply with the current LuminousIntensity.
//
// Returns:
//    *LuminousIntensity: A new LuminousIntensity instance representing the product of both LuminousIntensity.
func (a *LuminousIntensity) Multiply(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value * other.BaseValue()}
}

// Divide divides the current LuminousIntensity by the given LuminousIntensity and returns the result.
// The result is a new LuminousIntensity instance with the quotient of the values.
//
// Parameters:
//    other: The LuminousIntensity to divide the current LuminousIntensity by.
//
// Returns:
//    *LuminousIntensity: A new LuminousIntensity instance representing the quotient of both LuminousIntensity.
func (a *LuminousIntensity) Divide(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value / other.BaseValue()}
}

// GetLuminousIntensityAbbreviation gets the unit abbreviation.
func GetLuminousIntensityAbbreviation(unit LuminousIntensityUnits) string {
	switch unit { 
	case LuminousIntensityCandela:
		return "cd" 
	default:
		return ""
	}
}