// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PermittivityUnits defines various units of Permittivity.
type PermittivityUnits string

const (
    
        // 
        PermittivityFaradPerMeter PermittivityUnits = "FaradPerMeter"
)

var internalPermittivityUnitsMap = map[PermittivityUnits]bool{
	
	PermittivityFaradPerMeter: true,
}

// PermittivityDto represents a Permittivity measurement with a numerical value and its corresponding unit.
type PermittivityDto struct {
    // Value is the numerical representation of the Permittivity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Permittivity, as defined in the PermittivityUnits enumeration.
	Unit  PermittivityUnits `json:"unit" validate:"required,oneof=FaradPerMeter"`
}

// PermittivityDtoFactory groups methods for creating and serializing PermittivityDto objects.
type PermittivityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PermittivityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PermittivityDtoFactory) FromJSON(data []byte) (*PermittivityDto, error) {
	a := PermittivityDto{}

    // Parse JSON into PermittivityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a PermittivityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PermittivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Permittivity represents a measurement in a of Permittivity.
//
// In electromagnetism, permittivity is the measure of resistance that is encountered when forming an electric field in a particular medium.
type Permittivity struct {
	// value is the base measurement stored internally.
	value       float64
    
    farads_per_meterLazy *float64 
}

// PermittivityFactory groups methods for creating Permittivity instances.
type PermittivityFactory struct{}

// CreatePermittivity creates a new Permittivity instance from the given value and unit.
func (uf PermittivityFactory) CreatePermittivity(value float64, unit PermittivityUnits) (*Permittivity, error) {
	return newPermittivity(value, unit)
}

// FromDto converts a PermittivityDto to a Permittivity instance.
func (uf PermittivityFactory) FromDto(dto PermittivityDto) (*Permittivity, error) {
	return newPermittivity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Permittivity instance.
func (uf PermittivityFactory) FromDtoJSON(data []byte) (*Permittivity, error) {
	unitDto, err := PermittivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PermittivityDto from JSON: %w", err)
	}
	return PermittivityFactory{}.FromDto(*unitDto)
}


// FromFaradsPerMeter creates a new Permittivity instance from a value in FaradsPerMeter.
func (uf PermittivityFactory) FromFaradsPerMeter(value float64) (*Permittivity, error) {
	return newPermittivity(value, PermittivityFaradPerMeter)
}


// newPermittivity creates a new Permittivity.
func newPermittivity(value float64, fromUnit PermittivityUnits) (*Permittivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalPermittivityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in PermittivityUnits", fromUnit)
	}
	a := &Permittivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Permittivity in FaradPerMeter unit (the base/default quantity).
func (a *Permittivity) BaseValue() float64 {
	return a.value
}


// FaradsPerMeter returns the Permittivity value in FaradsPerMeter.
//
// 
func (a *Permittivity) FaradsPerMeter() float64 {
	if a.farads_per_meterLazy != nil {
		return *a.farads_per_meterLazy
	}
	farads_per_meter := a.convertFromBase(PermittivityFaradPerMeter)
	a.farads_per_meterLazy = &farads_per_meter
	return farads_per_meter
}


// ToDto creates a PermittivityDto representation from the Permittivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by FaradPerMeter by default.
func (a *Permittivity) ToDto(holdInUnit *PermittivityUnits) PermittivityDto {
	if holdInUnit == nil {
		defaultUnit := PermittivityFaradPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PermittivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Permittivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by FaradPerMeter by default.
func (a *Permittivity) ToDtoJSON(holdInUnit *PermittivityUnits) ([]byte, error) {
	// Convert to PermittivityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Permittivity to a specific unit value.
// The function uses the provided unit type (PermittivityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Permittivity) Convert(toUnit PermittivityUnits) float64 {
	switch toUnit { 
    case PermittivityFaradPerMeter:
		return a.FaradsPerMeter()
	default:
		return math.NaN()
	}
}

func (a *Permittivity) convertFromBase(toUnit PermittivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PermittivityFaradPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Permittivity) convertToBase(value float64, fromUnit PermittivityUnits) float64 {
	switch fromUnit { 
	case PermittivityFaradPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Permittivity in the default unit (FaradPerMeter),
// formatted to two decimal places.
func (a Permittivity) String() string {
	return a.ToString(PermittivityFaradPerMeter, 2)
}

// ToString formats the Permittivity value as a string with the specified unit and fractional digits.
// It converts the Permittivity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Permittivity value will be converted (e.g., FaradPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Permittivity with the unit abbreviation.
func (a *Permittivity) ToString(unit PermittivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetPermittivityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetPermittivityAbbreviation(unit))
}

// Equals checks if the given Permittivity is equal to the current Permittivity.
//
// Parameters:
//    other: The Permittivity to compare against.
//
// Returns:
//    bool: Returns true if both Permittivity are equal, false otherwise.
func (a *Permittivity) Equals(other *Permittivity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Permittivity with another Permittivity.
// It returns -1 if the current Permittivity is less than the other Permittivity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Permittivity to compare against.
//
// Returns:
//    int: -1 if the current Permittivity is less, 1 if greater, and 0 if equal.
func (a *Permittivity) CompareTo(other *Permittivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Permittivity to the current Permittivity and returns the result.
// The result is a new Permittivity instance with the sum of the values.
//
// Parameters:
//    other: The Permittivity to add to the current Permittivity.
//
// Returns:
//    *Permittivity: A new Permittivity instance representing the sum of both Permittivity.
func (a *Permittivity) Add(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Permittivity from the current Permittivity and returns the result.
// The result is a new Permittivity instance with the difference of the values.
//
// Parameters:
//    other: The Permittivity to subtract from the current Permittivity.
//
// Returns:
//    *Permittivity: A new Permittivity instance representing the difference of both Permittivity.
func (a *Permittivity) Subtract(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Permittivity by the given Permittivity and returns the result.
// The result is a new Permittivity instance with the product of the values.
//
// Parameters:
//    other: The Permittivity to multiply with the current Permittivity.
//
// Returns:
//    *Permittivity: A new Permittivity instance representing the product of both Permittivity.
func (a *Permittivity) Multiply(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value * other.BaseValue()}
}

// Divide divides the current Permittivity by the given Permittivity and returns the result.
// The result is a new Permittivity instance with the quotient of the values.
//
// Parameters:
//    other: The Permittivity to divide the current Permittivity by.
//
// Returns:
//    *Permittivity: A new Permittivity instance representing the quotient of both Permittivity.
func (a *Permittivity) Divide(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value / other.BaseValue()}
}

// GetPermittivityAbbreviation gets the unit abbreviation.
func GetPermittivityAbbreviation(unit PermittivityUnits) string {
	switch unit { 
	case PermittivityFaradPerMeter:
		return "F/m" 
	default:
		return ""
	}
}