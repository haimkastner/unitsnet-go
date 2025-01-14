// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ReactiveEnergyUnits defines various units of ReactiveEnergy.
type ReactiveEnergyUnits string

const (
    
        // 
        ReactiveEnergyVoltampereReactiveHour ReactiveEnergyUnits = "VoltampereReactiveHour"
        // 
        ReactiveEnergyKilovoltampereReactiveHour ReactiveEnergyUnits = "KilovoltampereReactiveHour"
        // 
        ReactiveEnergyMegavoltampereReactiveHour ReactiveEnergyUnits = "MegavoltampereReactiveHour"
)

// ReactiveEnergyDto represents a ReactiveEnergy measurement with a numerical value and its corresponding unit.
type ReactiveEnergyDto struct {
    // Value is the numerical representation of the ReactiveEnergy.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ReactiveEnergy, as defined in the ReactiveEnergyUnits enumeration.
	Unit  ReactiveEnergyUnits `json:"unit"`
}

// ReactiveEnergyDtoFactory groups methods for creating and serializing ReactiveEnergyDto objects.
type ReactiveEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ReactiveEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ReactiveEnergyDtoFactory) FromJSON(data []byte) (*ReactiveEnergyDto, error) {
	a := ReactiveEnergyDto{}

    // Parse JSON into ReactiveEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ReactiveEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ReactiveEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ReactiveEnergy represents a measurement in a of ReactiveEnergy.
//
// The Volt-ampere reactive hour (expressed as varh) is the reactive power of one Volt-ampere reactive produced in one hour.
type ReactiveEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltampere_reactive_hoursLazy *float64 
    kilovoltampere_reactive_hoursLazy *float64 
    megavoltampere_reactive_hoursLazy *float64 
}

// ReactiveEnergyFactory groups methods for creating ReactiveEnergy instances.
type ReactiveEnergyFactory struct{}

// CreateReactiveEnergy creates a new ReactiveEnergy instance from the given value and unit.
func (uf ReactiveEnergyFactory) CreateReactiveEnergy(value float64, unit ReactiveEnergyUnits) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, unit)
}

// FromDto converts a ReactiveEnergyDto to a ReactiveEnergy instance.
func (uf ReactiveEnergyFactory) FromDto(dto ReactiveEnergyDto) (*ReactiveEnergy, error) {
	return newReactiveEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ReactiveEnergy instance.
func (uf ReactiveEnergyFactory) FromDtoJSON(data []byte) (*ReactiveEnergy, error) {
	unitDto, err := ReactiveEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ReactiveEnergyDto from JSON: %w", err)
	}
	return ReactiveEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereReactiveHours creates a new ReactiveEnergy instance from a value in VoltampereReactiveHours.
func (uf ReactiveEnergyFactory) FromVoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyVoltampereReactiveHour)
}

// FromKilovoltampereReactiveHours creates a new ReactiveEnergy instance from a value in KilovoltampereReactiveHours.
func (uf ReactiveEnergyFactory) FromKilovoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyKilovoltampereReactiveHour)
}

// FromMegavoltampereReactiveHours creates a new ReactiveEnergy instance from a value in MegavoltampereReactiveHours.
func (uf ReactiveEnergyFactory) FromMegavoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyMegavoltampereReactiveHour)
}


// newReactiveEnergy creates a new ReactiveEnergy.
func newReactiveEnergy(value float64, fromUnit ReactiveEnergyUnits) (*ReactiveEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ReactiveEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReactiveEnergy in VoltampereReactiveHour unit (the base/default quantity).
func (a *ReactiveEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereReactiveHours returns the ReactiveEnergy value in VoltampereReactiveHours.
//
// 
func (a *ReactiveEnergy) VoltampereReactiveHours() float64 {
	if a.voltampere_reactive_hoursLazy != nil {
		return *a.voltampere_reactive_hoursLazy
	}
	voltampere_reactive_hours := a.convertFromBase(ReactiveEnergyVoltampereReactiveHour)
	a.voltampere_reactive_hoursLazy = &voltampere_reactive_hours
	return voltampere_reactive_hours
}

// KilovoltampereReactiveHours returns the ReactiveEnergy value in KilovoltampereReactiveHours.
//
// 
func (a *ReactiveEnergy) KilovoltampereReactiveHours() float64 {
	if a.kilovoltampere_reactive_hoursLazy != nil {
		return *a.kilovoltampere_reactive_hoursLazy
	}
	kilovoltampere_reactive_hours := a.convertFromBase(ReactiveEnergyKilovoltampereReactiveHour)
	a.kilovoltampere_reactive_hoursLazy = &kilovoltampere_reactive_hours
	return kilovoltampere_reactive_hours
}

// MegavoltampereReactiveHours returns the ReactiveEnergy value in MegavoltampereReactiveHours.
//
// 
func (a *ReactiveEnergy) MegavoltampereReactiveHours() float64 {
	if a.megavoltampere_reactive_hoursLazy != nil {
		return *a.megavoltampere_reactive_hoursLazy
	}
	megavoltampere_reactive_hours := a.convertFromBase(ReactiveEnergyMegavoltampereReactiveHour)
	a.megavoltampere_reactive_hoursLazy = &megavoltampere_reactive_hours
	return megavoltampere_reactive_hours
}


// ToDto creates a ReactiveEnergyDto representation from the ReactiveEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactiveHour by default.
func (a *ReactiveEnergy) ToDto(holdInUnit *ReactiveEnergyUnits) ReactiveEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ReactiveEnergyVoltampereReactiveHour // Default value
		holdInUnit = &defaultUnit
	}

	return ReactiveEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ReactiveEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactiveHour by default.
func (a *ReactiveEnergy) ToDtoJSON(holdInUnit *ReactiveEnergyUnits) ([]byte, error) {
	// Convert to ReactiveEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ReactiveEnergy to a specific unit value.
// The function uses the provided unit type (ReactiveEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ReactiveEnergy) Convert(toUnit ReactiveEnergyUnits) float64 {
	switch toUnit { 
    case ReactiveEnergyVoltampereReactiveHour:
		return a.VoltampereReactiveHours()
    case ReactiveEnergyKilovoltampereReactiveHour:
		return a.KilovoltampereReactiveHours()
    case ReactiveEnergyMegavoltampereReactiveHour:
		return a.MegavoltampereReactiveHours()
	default:
		return math.NaN()
	}
}

func (a *ReactiveEnergy) convertFromBase(toUnit ReactiveEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return ((value) / 1000.0) 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ReactiveEnergy) convertToBase(value float64, fromUnit ReactiveEnergyUnits) float64 {
	switch fromUnit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return ((value) * 1000.0) 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ReactiveEnergy in the default unit (VoltampereReactiveHour),
// formatted to two decimal places.
func (a ReactiveEnergy) String() string {
	return a.ToString(ReactiveEnergyVoltampereReactiveHour, 2)
}

// ToString formats the ReactiveEnergy value as a string with the specified unit and fractional digits.
// It converts the ReactiveEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ReactiveEnergy value will be converted (e.g., VoltampereReactiveHour))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ReactiveEnergy with the unit abbreviation.
func (a *ReactiveEnergy) ToString(unit ReactiveEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ReactiveEnergy) getUnitAbbreviation(unit ReactiveEnergyUnits) string {
	switch unit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return "varh" 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return "kvarh" 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return "Mvarh" 
	default:
		return ""
	}
}

// Equals checks if the given ReactiveEnergy is equal to the current ReactiveEnergy.
//
// Parameters:
//    other: The ReactiveEnergy to compare against.
//
// Returns:
//    bool: Returns true if both ReactiveEnergy are equal, false otherwise.
func (a *ReactiveEnergy) Equals(other *ReactiveEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ReactiveEnergy with another ReactiveEnergy.
// It returns -1 if the current ReactiveEnergy is less than the other ReactiveEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ReactiveEnergy to compare against.
//
// Returns:
//    int: -1 if the current ReactiveEnergy is less, 1 if greater, and 0 if equal.
func (a *ReactiveEnergy) CompareTo(other *ReactiveEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ReactiveEnergy to the current ReactiveEnergy and returns the result.
// The result is a new ReactiveEnergy instance with the sum of the values.
//
// Parameters:
//    other: The ReactiveEnergy to add to the current ReactiveEnergy.
//
// Returns:
//    *ReactiveEnergy: A new ReactiveEnergy instance representing the sum of both ReactiveEnergy.
func (a *ReactiveEnergy) Add(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ReactiveEnergy from the current ReactiveEnergy and returns the result.
// The result is a new ReactiveEnergy instance with the difference of the values.
//
// Parameters:
//    other: The ReactiveEnergy to subtract from the current ReactiveEnergy.
//
// Returns:
//    *ReactiveEnergy: A new ReactiveEnergy instance representing the difference of both ReactiveEnergy.
func (a *ReactiveEnergy) Subtract(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ReactiveEnergy by the given ReactiveEnergy and returns the result.
// The result is a new ReactiveEnergy instance with the product of the values.
//
// Parameters:
//    other: The ReactiveEnergy to multiply with the current ReactiveEnergy.
//
// Returns:
//    *ReactiveEnergy: A new ReactiveEnergy instance representing the product of both ReactiveEnergy.
func (a *ReactiveEnergy) Multiply(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current ReactiveEnergy by the given ReactiveEnergy and returns the result.
// The result is a new ReactiveEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The ReactiveEnergy to divide the current ReactiveEnergy by.
//
// Returns:
//    *ReactiveEnergy: A new ReactiveEnergy instance representing the quotient of both ReactiveEnergy.
func (a *ReactiveEnergy) Divide(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value / other.BaseValue()}
}