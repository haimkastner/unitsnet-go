// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PowerRatioUnits defines various units of PowerRatio.
type PowerRatioUnits string

const (
    
        // 
        PowerRatioDecibelWatt PowerRatioUnits = "DecibelWatt"
        // 
        PowerRatioDecibelMilliwatt PowerRatioUnits = "DecibelMilliwatt"
)

// PowerRatioDto represents a PowerRatio measurement with a numerical value and its corresponding unit.
type PowerRatioDto struct {
    // Value is the numerical representation of the PowerRatio.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the PowerRatio, as defined in the PowerRatioUnits enumeration.
	Unit  PowerRatioUnits `json:"unit"`
}

// PowerRatioDtoFactory groups methods for creating and serializing PowerRatioDto objects.
type PowerRatioDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PowerRatioDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PowerRatioDtoFactory) FromJSON(data []byte) (*PowerRatioDto, error) {
	a := PowerRatioDto{}

    // Parse JSON into PowerRatioDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a PowerRatioDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PowerRatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// PowerRatio represents a measurement in a of PowerRatio.
//
// The strength of a signal expressed in decibels (dB) relative to one watt.
type PowerRatio struct {
	// value is the base measurement stored internally.
	value       float64
    
    decibel_wattsLazy *float64 
    decibel_milliwattsLazy *float64 
}

// PowerRatioFactory groups methods for creating PowerRatio instances.
type PowerRatioFactory struct{}

// CreatePowerRatio creates a new PowerRatio instance from the given value and unit.
func (uf PowerRatioFactory) CreatePowerRatio(value float64, unit PowerRatioUnits) (*PowerRatio, error) {
	return newPowerRatio(value, unit)
}

// FromDto converts a PowerRatioDto to a PowerRatio instance.
func (uf PowerRatioFactory) FromDto(dto PowerRatioDto) (*PowerRatio, error) {
	return newPowerRatio(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a PowerRatio instance.
func (uf PowerRatioFactory) FromDtoJSON(data []byte) (*PowerRatio, error) {
	unitDto, err := PowerRatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PowerRatioDto from JSON: %w", err)
	}
	return PowerRatioFactory{}.FromDto(*unitDto)
}


// FromDecibelWatts creates a new PowerRatio instance from a value in DecibelWatts.
func (uf PowerRatioFactory) FromDecibelWatts(value float64) (*PowerRatio, error) {
	return newPowerRatio(value, PowerRatioDecibelWatt)
}

// FromDecibelMilliwatts creates a new PowerRatio instance from a value in DecibelMilliwatts.
func (uf PowerRatioFactory) FromDecibelMilliwatts(value float64) (*PowerRatio, error) {
	return newPowerRatio(value, PowerRatioDecibelMilliwatt)
}


// newPowerRatio creates a new PowerRatio.
func newPowerRatio(value float64, fromUnit PowerRatioUnits) (*PowerRatio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &PowerRatio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PowerRatio in DecibelWatt unit (the base/default quantity).
func (a *PowerRatio) BaseValue() float64 {
	return a.value
}


// DecibelWatts returns the PowerRatio value in DecibelWatts.
//
// 
func (a *PowerRatio) DecibelWatts() float64 {
	if a.decibel_wattsLazy != nil {
		return *a.decibel_wattsLazy
	}
	decibel_watts := a.convertFromBase(PowerRatioDecibelWatt)
	a.decibel_wattsLazy = &decibel_watts
	return decibel_watts
}

// DecibelMilliwatts returns the PowerRatio value in DecibelMilliwatts.
//
// 
func (a *PowerRatio) DecibelMilliwatts() float64 {
	if a.decibel_milliwattsLazy != nil {
		return *a.decibel_milliwattsLazy
	}
	decibel_milliwatts := a.convertFromBase(PowerRatioDecibelMilliwatt)
	a.decibel_milliwattsLazy = &decibel_milliwatts
	return decibel_milliwatts
}


// ToDto creates a PowerRatioDto representation from the PowerRatio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecibelWatt by default.
func (a *PowerRatio) ToDto(holdInUnit *PowerRatioUnits) PowerRatioDto {
	if holdInUnit == nil {
		defaultUnit := PowerRatioDecibelWatt // Default value
		holdInUnit = &defaultUnit
	}

	return PowerRatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the PowerRatio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecibelWatt by default.
func (a *PowerRatio) ToDtoJSON(holdInUnit *PowerRatioUnits) ([]byte, error) {
	// Convert to PowerRatioDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a PowerRatio to a specific unit value.
// The function uses the provided unit type (PowerRatioUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *PowerRatio) Convert(toUnit PowerRatioUnits) float64 {
	switch toUnit { 
    case PowerRatioDecibelWatt:
		return a.DecibelWatts()
    case PowerRatioDecibelMilliwatt:
		return a.DecibelMilliwatts()
	default:
		return math.NaN()
	}
}

func (a *PowerRatio) convertFromBase(toUnit PowerRatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case PowerRatioDecibelWatt:
		return (value) 
	case PowerRatioDecibelMilliwatt:
		return (value + 30) 
	default:
		return math.NaN()
	}
}

func (a *PowerRatio) convertToBase(value float64, fromUnit PowerRatioUnits) float64 {
	switch fromUnit { 
	case PowerRatioDecibelWatt:
		return (value) 
	case PowerRatioDecibelMilliwatt:
		return (value - 30) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the PowerRatio in the default unit (DecibelWatt),
// formatted to two decimal places.
func (a PowerRatio) String() string {
	return a.ToString(PowerRatioDecibelWatt, 2)
}

// ToString formats the PowerRatio value as a string with the specified unit and fractional digits.
// It converts the PowerRatio to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the PowerRatio value will be converted (e.g., DecibelWatt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the PowerRatio with the unit abbreviation.
func (a *PowerRatio) ToString(unit PowerRatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *PowerRatio) getUnitAbbreviation(unit PowerRatioUnits) string {
	switch unit { 
	case PowerRatioDecibelWatt:
		return "dBW" 
	case PowerRatioDecibelMilliwatt:
		return "dBmW" 
	default:
		return ""
	}
}

// Equals checks if the given PowerRatio is equal to the current PowerRatio.
//
// Parameters:
//    other: The PowerRatio to compare against.
//
// Returns:
//    bool: Returns true if both PowerRatio are equal, false otherwise.
func (a *PowerRatio) Equals(other *PowerRatio) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current PowerRatio with another PowerRatio.
// It returns -1 if the current PowerRatio is less than the other PowerRatio, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The PowerRatio to compare against.
//
// Returns:
//    int: -1 if the current PowerRatio is less, 1 if greater, and 0 if equal.
func (a *PowerRatio) CompareTo(other *PowerRatio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given PowerRatio to the current PowerRatio and returns the result.
// The result is a new PowerRatio instance with the sum of the values.
//
// Parameters:
//    other: The PowerRatio to add to the current PowerRatio.
//
// Returns:
//    *PowerRatio: A new PowerRatio instance representing the sum of both PowerRatio.
func (a *PowerRatio) Add(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given PowerRatio from the current PowerRatio and returns the result.
// The result is a new PowerRatio instance with the difference of the values.
//
// Parameters:
//    other: The PowerRatio to subtract from the current PowerRatio.
//
// Returns:
//    *PowerRatio: A new PowerRatio instance representing the difference of both PowerRatio.
func (a *PowerRatio) Subtract(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current PowerRatio by the given PowerRatio and returns the result.
// The result is a new PowerRatio instance with the product of the values.
//
// Parameters:
//    other: The PowerRatio to multiply with the current PowerRatio.
//
// Returns:
//    *PowerRatio: A new PowerRatio instance representing the product of both PowerRatio.
func (a *PowerRatio) Multiply(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value * other.BaseValue()}
}

// Divide divides the current PowerRatio by the given PowerRatio and returns the result.
// The result is a new PowerRatio instance with the quotient of the values.
//
// Parameters:
//    other: The PowerRatio to divide the current PowerRatio by.
//
// Returns:
//    *PowerRatio: A new PowerRatio instance representing the quotient of both PowerRatio.
func (a *PowerRatio) Divide(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value / other.BaseValue()}
}