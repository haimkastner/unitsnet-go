// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PermeabilityUnits defines various units of Permeability.
type PermeabilityUnits string

const (
    
        // 
        PermeabilityHenryPerMeter PermeabilityUnits = "HenryPerMeter"
)

// PermeabilityDto represents a Permeability measurement with a numerical value and its corresponding unit.
type PermeabilityDto struct {
    // Value is the numerical representation of the Permeability.
	Value float64
    // Unit specifies the unit of measurement for the Permeability, as defined in the PermeabilityUnits enumeration.
	Unit  PermeabilityUnits
}

// PermeabilityDtoFactory groups methods for creating and serializing PermeabilityDto objects.
type PermeabilityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PermeabilityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PermeabilityDtoFactory) FromJSON(data []byte) (*PermeabilityDto, error) {
	a := PermeabilityDto{}

    // Parse JSON into PermeabilityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a PermeabilityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PermeabilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Permeability represents a measurement in a of Permeability.
//
// In electromagnetism, permeability is the measure of the ability of a material to support the formation of a magnetic field within itself.
type Permeability struct {
	// value is the base measurement stored internally.
	value       float64
    
    henries_per_meterLazy *float64 
}

// PermeabilityFactory groups methods for creating Permeability instances.
type PermeabilityFactory struct{}

// CreatePermeability creates a new Permeability instance from the given value and unit.
func (uf PermeabilityFactory) CreatePermeability(value float64, unit PermeabilityUnits) (*Permeability, error) {
	return newPermeability(value, unit)
}

// FromDto converts a PermeabilityDto to a Permeability instance.
func (uf PermeabilityFactory) FromDto(dto PermeabilityDto) (*Permeability, error) {
	return newPermeability(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Permeability instance.
func (uf PermeabilityFactory) FromDtoJSON(data []byte) (*Permeability, error) {
	unitDto, err := PermeabilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PermeabilityDto from JSON: %w", err)
	}
	return PermeabilityFactory{}.FromDto(*unitDto)
}


// FromHenriesPerMeter creates a new Permeability instance from a value in HenriesPerMeter.
func (uf PermeabilityFactory) FromHenriesPerMeter(value float64) (*Permeability, error) {
	return newPermeability(value, PermeabilityHenryPerMeter)
}


// newPermeability creates a new Permeability.
func newPermeability(value float64, fromUnit PermeabilityUnits) (*Permeability, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Permeability{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Permeability in HenryPerMeter unit (the base/default quantity).
func (a *Permeability) BaseValue() float64 {
	return a.value
}


// HenriesPerMeter returns the Permeability value in HenriesPerMeter.
//
// 
func (a *Permeability) HenriesPerMeter() float64 {
	if a.henries_per_meterLazy != nil {
		return *a.henries_per_meterLazy
	}
	henries_per_meter := a.convertFromBase(PermeabilityHenryPerMeter)
	a.henries_per_meterLazy = &henries_per_meter
	return henries_per_meter
}


// ToDto creates a PermeabilityDto representation from the Permeability instance.
//
// If the provided holdInUnit is nil, the value will be repesented by HenryPerMeter by default.
func (a *Permeability) ToDto(holdInUnit *PermeabilityUnits) PermeabilityDto {
	if holdInUnit == nil {
		defaultUnit := PermeabilityHenryPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PermeabilityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Permeability instance.
//
// If the provided holdInUnit is nil, the value will be repesented by HenryPerMeter by default.
func (a *Permeability) ToDtoJSON(holdInUnit *PermeabilityUnits) ([]byte, error) {
	// Convert to PermeabilityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Permeability to a specific unit value.
// The function uses the provided unit type (PermeabilityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Permeability) Convert(toUnit PermeabilityUnits) float64 {
	switch toUnit { 
    case PermeabilityHenryPerMeter:
		return a.HenriesPerMeter()
	default:
		return math.NaN()
	}
}

func (a *Permeability) convertFromBase(toUnit PermeabilityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PermeabilityHenryPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Permeability) convertToBase(value float64, fromUnit PermeabilityUnits) float64 {
	switch fromUnit { 
	case PermeabilityHenryPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Permeability in the default unit (HenryPerMeter),
// formatted to two decimal places.
func (a Permeability) String() string {
	return a.ToString(PermeabilityHenryPerMeter, 2)
}

// ToString formats the Permeability value as a string with the specified unit and fractional digits.
// It converts the Permeability to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Permeability value will be converted (e.g., HenryPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Permeability with the unit abbreviation.
func (a *Permeability) ToString(unit PermeabilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Permeability) getUnitAbbreviation(unit PermeabilityUnits) string {
	switch unit { 
	case PermeabilityHenryPerMeter:
		return "H/m" 
	default:
		return ""
	}
}

// Equals checks if the given Permeability is equal to the current Permeability.
//
// Parameters:
//    other: The Permeability to compare against.
//
// Returns:
//    bool: Returns true if both Permeability are equal, false otherwise.
func (a *Permeability) Equals(other *Permeability) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Permeability with another Permeability.
// It returns -1 if the current Permeability is less than the other Permeability, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Permeability to compare against.
//
// Returns:
//    int: -1 if the current Permeability is less, 1 if greater, and 0 if equal.
func (a *Permeability) CompareTo(other *Permeability) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Permeability to the current Permeability and returns the result.
// The result is a new Permeability instance with the sum of the values.
//
// Parameters:
//    other: The Permeability to add to the current Permeability.
//
// Returns:
//    *Permeability: A new Permeability instance representing the sum of both Permeability.
func (a *Permeability) Add(other *Permeability) *Permeability {
	return &Permeability{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Permeability from the current Permeability and returns the result.
// The result is a new Permeability instance with the difference of the values.
//
// Parameters:
//    other: The Permeability to subtract from the current Permeability.
//
// Returns:
//    *Permeability: A new Permeability instance representing the difference of both Permeability.
func (a *Permeability) Subtract(other *Permeability) *Permeability {
	return &Permeability{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Permeability by the given Permeability and returns the result.
// The result is a new Permeability instance with the product of the values.
//
// Parameters:
//    other: The Permeability to multiply with the current Permeability.
//
// Returns:
//    *Permeability: A new Permeability instance representing the product of both Permeability.
func (a *Permeability) Multiply(other *Permeability) *Permeability {
	return &Permeability{value: a.value * other.BaseValue()}
}

// Divide divides the current Permeability by the given Permeability and returns the result.
// The result is a new Permeability instance with the quotient of the values.
//
// Parameters:
//    other: The Permeability to divide the current Permeability by.
//
// Returns:
//    *Permeability: A new Permeability instance representing the quotient of both Permeability.
func (a *Permeability) Divide(other *Permeability) *Permeability {
	return &Permeability{value: a.value / other.BaseValue()}
}