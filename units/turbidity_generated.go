// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// TurbidityUnits defines various units of Turbidity.
type TurbidityUnits string

const (
    
        // 
        TurbidityNTU TurbidityUnits = "NTU"
)

// TurbidityDto represents a Turbidity measurement with a numerical value and its corresponding unit.
type TurbidityDto struct {
    // Value is the numerical representation of the Turbidity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Turbidity, as defined in the TurbidityUnits enumeration.
	Unit  TurbidityUnits `json:"unit"`
}

// TurbidityDtoFactory groups methods for creating and serializing TurbidityDto objects.
type TurbidityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TurbidityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TurbidityDtoFactory) FromJSON(data []byte) (*TurbidityDto, error) {
	a := TurbidityDto{}

    // Parse JSON into TurbidityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a TurbidityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TurbidityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Turbidity represents a measurement in a of Turbidity.
//
// Turbidity is the cloudiness or haziness of a fluid caused by large numbers of individual particles that are generally invisible to the naked eye, similar to smoke in air. The measurement of turbidity is a key test of water quality.
type Turbidity struct {
	// value is the base measurement stored internally.
	value       float64
    
    ntuLazy *float64 
}

// TurbidityFactory groups methods for creating Turbidity instances.
type TurbidityFactory struct{}

// CreateTurbidity creates a new Turbidity instance from the given value and unit.
func (uf TurbidityFactory) CreateTurbidity(value float64, unit TurbidityUnits) (*Turbidity, error) {
	return newTurbidity(value, unit)
}

// FromDto converts a TurbidityDto to a Turbidity instance.
func (uf TurbidityFactory) FromDto(dto TurbidityDto) (*Turbidity, error) {
	return newTurbidity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Turbidity instance.
func (uf TurbidityFactory) FromDtoJSON(data []byte) (*Turbidity, error) {
	unitDto, err := TurbidityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TurbidityDto from JSON: %w", err)
	}
	return TurbidityFactory{}.FromDto(*unitDto)
}


// FromNTU creates a new Turbidity instance from a value in NTU.
func (uf TurbidityFactory) FromNTU(value float64) (*Turbidity, error) {
	return newTurbidity(value, TurbidityNTU)
}


// newTurbidity creates a new Turbidity.
func newTurbidity(value float64, fromUnit TurbidityUnits) (*Turbidity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Turbidity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Turbidity in NTU unit (the base/default quantity).
func (a *Turbidity) BaseValue() float64 {
	return a.value
}


// NTU returns the Turbidity value in NTU.
//
// 
func (a *Turbidity) NTU() float64 {
	if a.ntuLazy != nil {
		return *a.ntuLazy
	}
	ntu := a.convertFromBase(TurbidityNTU)
	a.ntuLazy = &ntu
	return ntu
}


// ToDto creates a TurbidityDto representation from the Turbidity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NTU by default.
func (a *Turbidity) ToDto(holdInUnit *TurbidityUnits) TurbidityDto {
	if holdInUnit == nil {
		defaultUnit := TurbidityNTU // Default value
		holdInUnit = &defaultUnit
	}

	return TurbidityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Turbidity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NTU by default.
func (a *Turbidity) ToDtoJSON(holdInUnit *TurbidityUnits) ([]byte, error) {
	// Convert to TurbidityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Turbidity to a specific unit value.
// The function uses the provided unit type (TurbidityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Turbidity) Convert(toUnit TurbidityUnits) float64 {
	switch toUnit { 
    case TurbidityNTU:
		return a.NTU()
	default:
		return math.NaN()
	}
}

func (a *Turbidity) convertFromBase(toUnit TurbidityUnits) float64 {
    value := a.value
	switch toUnit { 
	case TurbidityNTU:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Turbidity) convertToBase(value float64, fromUnit TurbidityUnits) float64 {
	switch fromUnit { 
	case TurbidityNTU:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Turbidity in the default unit (NTU),
// formatted to two decimal places.
func (a Turbidity) String() string {
	return a.ToString(TurbidityNTU, 2)
}

// ToString formats the Turbidity value as a string with the specified unit and fractional digits.
// It converts the Turbidity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Turbidity value will be converted (e.g., NTU))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Turbidity with the unit abbreviation.
func (a *Turbidity) ToString(unit TurbidityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetTurbidityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetTurbidityAbbreviation(unit))
}

// Equals checks if the given Turbidity is equal to the current Turbidity.
//
// Parameters:
//    other: The Turbidity to compare against.
//
// Returns:
//    bool: Returns true if both Turbidity are equal, false otherwise.
func (a *Turbidity) Equals(other *Turbidity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Turbidity with another Turbidity.
// It returns -1 if the current Turbidity is less than the other Turbidity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Turbidity to compare against.
//
// Returns:
//    int: -1 if the current Turbidity is less, 1 if greater, and 0 if equal.
func (a *Turbidity) CompareTo(other *Turbidity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Turbidity to the current Turbidity and returns the result.
// The result is a new Turbidity instance with the sum of the values.
//
// Parameters:
//    other: The Turbidity to add to the current Turbidity.
//
// Returns:
//    *Turbidity: A new Turbidity instance representing the sum of both Turbidity.
func (a *Turbidity) Add(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Turbidity from the current Turbidity and returns the result.
// The result is a new Turbidity instance with the difference of the values.
//
// Parameters:
//    other: The Turbidity to subtract from the current Turbidity.
//
// Returns:
//    *Turbidity: A new Turbidity instance representing the difference of both Turbidity.
func (a *Turbidity) Subtract(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Turbidity by the given Turbidity and returns the result.
// The result is a new Turbidity instance with the product of the values.
//
// Parameters:
//    other: The Turbidity to multiply with the current Turbidity.
//
// Returns:
//    *Turbidity: A new Turbidity instance representing the product of both Turbidity.
func (a *Turbidity) Multiply(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value * other.BaseValue()}
}

// Divide divides the current Turbidity by the given Turbidity and returns the result.
// The result is a new Turbidity instance with the quotient of the values.
//
// Parameters:
//    other: The Turbidity to divide the current Turbidity by.
//
// Returns:
//    *Turbidity: A new Turbidity instance representing the quotient of both Turbidity.
func (a *Turbidity) Divide(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value / other.BaseValue()}
}

// GetTurbidityAbbreviation gets the unit abbreviation.
func GetTurbidityAbbreviation(unit TurbidityUnits) string {
	switch unit { 
	case TurbidityNTU:
		return "NTU" 
	default:
		return ""
	}
}