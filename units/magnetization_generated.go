// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MagnetizationUnits defines various units of Magnetization.
type MagnetizationUnits string

const (
    
        // 
        MagnetizationAmperePerMeter MagnetizationUnits = "AmperePerMeter"
)

var internalMagnetizationUnitsMap = map[MagnetizationUnits]bool{
	
	MagnetizationAmperePerMeter: true,
}

// MagnetizationDto represents a Magnetization measurement with a numerical value and its corresponding unit.
type MagnetizationDto struct {
    // Value is the numerical representation of the Magnetization.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Magnetization, as defined in the MagnetizationUnits enumeration.
	Unit  MagnetizationUnits `json:"unit" validate:"required,oneof=AmperePerMeter"`
}

// MagnetizationDtoFactory groups methods for creating and serializing MagnetizationDto objects.
type MagnetizationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MagnetizationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MagnetizationDtoFactory) FromJSON(data []byte) (*MagnetizationDto, error) {
	a := MagnetizationDto{}

    // Parse JSON into MagnetizationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MagnetizationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MagnetizationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Magnetization represents a measurement in a of Magnetization.
//
// In classical electromagnetism, magnetization is the vector field that expresses the density of permanent or induced magnetic dipole moments in a magnetic material.
type Magnetization struct {
	// value is the base measurement stored internally.
	value       float64
    
    amperes_per_meterLazy *float64 
}

// MagnetizationFactory groups methods for creating Magnetization instances.
type MagnetizationFactory struct{}

// CreateMagnetization creates a new Magnetization instance from the given value and unit.
func (uf MagnetizationFactory) CreateMagnetization(value float64, unit MagnetizationUnits) (*Magnetization, error) {
	return newMagnetization(value, unit)
}

// FromDto converts a MagnetizationDto to a Magnetization instance.
func (uf MagnetizationFactory) FromDto(dto MagnetizationDto) (*Magnetization, error) {
	return newMagnetization(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Magnetization instance.
func (uf MagnetizationFactory) FromDtoJSON(data []byte) (*Magnetization, error) {
	unitDto, err := MagnetizationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MagnetizationDto from JSON: %w", err)
	}
	return MagnetizationFactory{}.FromDto(*unitDto)
}


// FromAmperesPerMeter creates a new Magnetization instance from a value in AmperesPerMeter.
func (uf MagnetizationFactory) FromAmperesPerMeter(value float64) (*Magnetization, error) {
	return newMagnetization(value, MagnetizationAmperePerMeter)
}


// newMagnetization creates a new Magnetization.
func newMagnetization(value float64, fromUnit MagnetizationUnits) (*Magnetization, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMagnetizationUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MagnetizationUnits", fromUnit)
	}
	a := &Magnetization{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Magnetization in AmperePerMeter unit (the base/default quantity).
func (a *Magnetization) BaseValue() float64 {
	return a.value
}


// AmperesPerMeter returns the Magnetization value in AmperesPerMeter.
//
// 
func (a *Magnetization) AmperesPerMeter() float64 {
	if a.amperes_per_meterLazy != nil {
		return *a.amperes_per_meterLazy
	}
	amperes_per_meter := a.convertFromBase(MagnetizationAmperePerMeter)
	a.amperes_per_meterLazy = &amperes_per_meter
	return amperes_per_meter
}


// ToDto creates a MagnetizationDto representation from the Magnetization instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerMeter by default.
func (a *Magnetization) ToDto(holdInUnit *MagnetizationUnits) MagnetizationDto {
	if holdInUnit == nil {
		defaultUnit := MagnetizationAmperePerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MagnetizationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Magnetization instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerMeter by default.
func (a *Magnetization) ToDtoJSON(holdInUnit *MagnetizationUnits) ([]byte, error) {
	// Convert to MagnetizationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Magnetization to a specific unit value.
// The function uses the provided unit type (MagnetizationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Magnetization) Convert(toUnit MagnetizationUnits) float64 {
	switch toUnit { 
    case MagnetizationAmperePerMeter:
		return a.AmperesPerMeter()
	default:
		return math.NaN()
	}
}

func (a *Magnetization) convertFromBase(toUnit MagnetizationUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagnetizationAmperePerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Magnetization) convertToBase(value float64, fromUnit MagnetizationUnits) float64 {
	switch fromUnit { 
	case MagnetizationAmperePerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Magnetization in the default unit (AmperePerMeter),
// formatted to two decimal places.
func (a Magnetization) String() string {
	return a.ToString(MagnetizationAmperePerMeter, 2)
}

// ToString formats the Magnetization value as a string with the specified unit and fractional digits.
// It converts the Magnetization to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Magnetization value will be converted (e.g., AmperePerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Magnetization with the unit abbreviation.
func (a *Magnetization) ToString(unit MagnetizationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMagnetizationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMagnetizationAbbreviation(unit))
}

// Equals checks if the given Magnetization is equal to the current Magnetization.
//
// Parameters:
//    other: The Magnetization to compare against.
//
// Returns:
//    bool: Returns true if both Magnetization are equal, false otherwise.
func (a *Magnetization) Equals(other *Magnetization) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Magnetization with another Magnetization.
// It returns -1 if the current Magnetization is less than the other Magnetization, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Magnetization to compare against.
//
// Returns:
//    int: -1 if the current Magnetization is less, 1 if greater, and 0 if equal.
func (a *Magnetization) CompareTo(other *Magnetization) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Magnetization to the current Magnetization and returns the result.
// The result is a new Magnetization instance with the sum of the values.
//
// Parameters:
//    other: The Magnetization to add to the current Magnetization.
//
// Returns:
//    *Magnetization: A new Magnetization instance representing the sum of both Magnetization.
func (a *Magnetization) Add(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Magnetization from the current Magnetization and returns the result.
// The result is a new Magnetization instance with the difference of the values.
//
// Parameters:
//    other: The Magnetization to subtract from the current Magnetization.
//
// Returns:
//    *Magnetization: A new Magnetization instance representing the difference of both Magnetization.
func (a *Magnetization) Subtract(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Magnetization by the given Magnetization and returns the result.
// The result is a new Magnetization instance with the product of the values.
//
// Parameters:
//    other: The Magnetization to multiply with the current Magnetization.
//
// Returns:
//    *Magnetization: A new Magnetization instance representing the product of both Magnetization.
func (a *Magnetization) Multiply(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value * other.BaseValue()}
}

// Divide divides the current Magnetization by the given Magnetization and returns the result.
// The result is a new Magnetization instance with the quotient of the values.
//
// Parameters:
//    other: The Magnetization to divide the current Magnetization by.
//
// Returns:
//    *Magnetization: A new Magnetization instance representing the quotient of both Magnetization.
func (a *Magnetization) Divide(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value / other.BaseValue()}
}

// GetMagnetizationAbbreviation gets the unit abbreviation.
func GetMagnetizationAbbreviation(unit MagnetizationUnits) string {
	switch unit { 
	case MagnetizationAmperePerMeter:
		return "A/m" 
	default:
		return ""
	}
}