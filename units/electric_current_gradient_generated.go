// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentGradientUnits defines various units of ElectricCurrentGradient.
type ElectricCurrentGradientUnits string

const (
    
        // 
        ElectricCurrentGradientAmperePerSecond ElectricCurrentGradientUnits = "AmperePerSecond"
        // 
        ElectricCurrentGradientAmperePerMinute ElectricCurrentGradientUnits = "AmperePerMinute"
        // 
        ElectricCurrentGradientAmperePerMillisecond ElectricCurrentGradientUnits = "AmperePerMillisecond"
        // 
        ElectricCurrentGradientAmperePerMicrosecond ElectricCurrentGradientUnits = "AmperePerMicrosecond"
        // 
        ElectricCurrentGradientAmperePerNanosecond ElectricCurrentGradientUnits = "AmperePerNanosecond"
        // 
        ElectricCurrentGradientMilliamperePerSecond ElectricCurrentGradientUnits = "MilliamperePerSecond"
        // 
        ElectricCurrentGradientMilliamperePerMinute ElectricCurrentGradientUnits = "MilliamperePerMinute"
)

// ElectricCurrentGradientDto represents a ElectricCurrentGradient measurement with a numerical value and its corresponding unit.
type ElectricCurrentGradientDto struct {
    // Value is the numerical representation of the ElectricCurrentGradient.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricCurrentGradient, as defined in the ElectricCurrentGradientUnits enumeration.
	Unit  ElectricCurrentGradientUnits `json:"unit"`
}

// ElectricCurrentGradientDtoFactory groups methods for creating and serializing ElectricCurrentGradientDto objects.
type ElectricCurrentGradientDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrentGradientDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricCurrentGradientDtoFactory) FromJSON(data []byte) (*ElectricCurrentGradientDto, error) {
	a := ElectricCurrentGradientDto{}

    // Parse JSON into ElectricCurrentGradientDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricCurrentGradientDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricCurrentGradientDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricCurrentGradient represents a measurement in a of ElectricCurrentGradient.
//
// In electromagnetism, the current gradient describes how the current changes in time.
type ElectricCurrentGradient struct {
	// value is the base measurement stored internally.
	value       float64
    
    amperes_per_secondLazy *float64 
    amperes_per_minuteLazy *float64 
    amperes_per_millisecondLazy *float64 
    amperes_per_microsecondLazy *float64 
    amperes_per_nanosecondLazy *float64 
    milliamperes_per_secondLazy *float64 
    milliamperes_per_minuteLazy *float64 
}

// ElectricCurrentGradientFactory groups methods for creating ElectricCurrentGradient instances.
type ElectricCurrentGradientFactory struct{}

// CreateElectricCurrentGradient creates a new ElectricCurrentGradient instance from the given value and unit.
func (uf ElectricCurrentGradientFactory) CreateElectricCurrentGradient(value float64, unit ElectricCurrentGradientUnits) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, unit)
}

// FromDto converts a ElectricCurrentGradientDto to a ElectricCurrentGradient instance.
func (uf ElectricCurrentGradientFactory) FromDto(dto ElectricCurrentGradientDto) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrentGradient instance.
func (uf ElectricCurrentGradientFactory) FromDtoJSON(data []byte) (*ElectricCurrentGradient, error) {
	unitDto, err := ElectricCurrentGradientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricCurrentGradientDto from JSON: %w", err)
	}
	return ElectricCurrentGradientFactory{}.FromDto(*unitDto)
}


// FromAmperesPerSecond creates a new ElectricCurrentGradient instance from a value in AmperesPerSecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerSecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerSecond)
}

// FromAmperesPerMinute creates a new ElectricCurrentGradient instance from a value in AmperesPerMinute.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMinute(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMinute)
}

// FromAmperesPerMillisecond creates a new ElectricCurrentGradient instance from a value in AmperesPerMillisecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMillisecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMillisecond)
}

// FromAmperesPerMicrosecond creates a new ElectricCurrentGradient instance from a value in AmperesPerMicrosecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMicrosecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMicrosecond)
}

// FromAmperesPerNanosecond creates a new ElectricCurrentGradient instance from a value in AmperesPerNanosecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerNanosecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerNanosecond)
}

// FromMilliamperesPerSecond creates a new ElectricCurrentGradient instance from a value in MilliamperesPerSecond.
func (uf ElectricCurrentGradientFactory) FromMilliamperesPerSecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientMilliamperePerSecond)
}

// FromMilliamperesPerMinute creates a new ElectricCurrentGradient instance from a value in MilliamperesPerMinute.
func (uf ElectricCurrentGradientFactory) FromMilliamperesPerMinute(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientMilliamperePerMinute)
}


// newElectricCurrentGradient creates a new ElectricCurrentGradient.
func newElectricCurrentGradient(value float64, fromUnit ElectricCurrentGradientUnits) (*ElectricCurrentGradient, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrentGradient{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrentGradient in AmperePerSecond unit (the base/default quantity).
func (a *ElectricCurrentGradient) BaseValue() float64 {
	return a.value
}


// AmperesPerSecond returns the ElectricCurrentGradient value in AmperesPerSecond.
//
// 
func (a *ElectricCurrentGradient) AmperesPerSecond() float64 {
	if a.amperes_per_secondLazy != nil {
		return *a.amperes_per_secondLazy
	}
	amperes_per_second := a.convertFromBase(ElectricCurrentGradientAmperePerSecond)
	a.amperes_per_secondLazy = &amperes_per_second
	return amperes_per_second
}

// AmperesPerMinute returns the ElectricCurrentGradient value in AmperesPerMinute.
//
// 
func (a *ElectricCurrentGradient) AmperesPerMinute() float64 {
	if a.amperes_per_minuteLazy != nil {
		return *a.amperes_per_minuteLazy
	}
	amperes_per_minute := a.convertFromBase(ElectricCurrentGradientAmperePerMinute)
	a.amperes_per_minuteLazy = &amperes_per_minute
	return amperes_per_minute
}

// AmperesPerMillisecond returns the ElectricCurrentGradient value in AmperesPerMillisecond.
//
// 
func (a *ElectricCurrentGradient) AmperesPerMillisecond() float64 {
	if a.amperes_per_millisecondLazy != nil {
		return *a.amperes_per_millisecondLazy
	}
	amperes_per_millisecond := a.convertFromBase(ElectricCurrentGradientAmperePerMillisecond)
	a.amperes_per_millisecondLazy = &amperes_per_millisecond
	return amperes_per_millisecond
}

// AmperesPerMicrosecond returns the ElectricCurrentGradient value in AmperesPerMicrosecond.
//
// 
func (a *ElectricCurrentGradient) AmperesPerMicrosecond() float64 {
	if a.amperes_per_microsecondLazy != nil {
		return *a.amperes_per_microsecondLazy
	}
	amperes_per_microsecond := a.convertFromBase(ElectricCurrentGradientAmperePerMicrosecond)
	a.amperes_per_microsecondLazy = &amperes_per_microsecond
	return amperes_per_microsecond
}

// AmperesPerNanosecond returns the ElectricCurrentGradient value in AmperesPerNanosecond.
//
// 
func (a *ElectricCurrentGradient) AmperesPerNanosecond() float64 {
	if a.amperes_per_nanosecondLazy != nil {
		return *a.amperes_per_nanosecondLazy
	}
	amperes_per_nanosecond := a.convertFromBase(ElectricCurrentGradientAmperePerNanosecond)
	a.amperes_per_nanosecondLazy = &amperes_per_nanosecond
	return amperes_per_nanosecond
}

// MilliamperesPerSecond returns the ElectricCurrentGradient value in MilliamperesPerSecond.
//
// 
func (a *ElectricCurrentGradient) MilliamperesPerSecond() float64 {
	if a.milliamperes_per_secondLazy != nil {
		return *a.milliamperes_per_secondLazy
	}
	milliamperes_per_second := a.convertFromBase(ElectricCurrentGradientMilliamperePerSecond)
	a.milliamperes_per_secondLazy = &milliamperes_per_second
	return milliamperes_per_second
}

// MilliamperesPerMinute returns the ElectricCurrentGradient value in MilliamperesPerMinute.
//
// 
func (a *ElectricCurrentGradient) MilliamperesPerMinute() float64 {
	if a.milliamperes_per_minuteLazy != nil {
		return *a.milliamperes_per_minuteLazy
	}
	milliamperes_per_minute := a.convertFromBase(ElectricCurrentGradientMilliamperePerMinute)
	a.milliamperes_per_minuteLazy = &milliamperes_per_minute
	return milliamperes_per_minute
}


// ToDto creates a ElectricCurrentGradientDto representation from the ElectricCurrentGradient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerSecond by default.
func (a *ElectricCurrentGradient) ToDto(holdInUnit *ElectricCurrentGradientUnits) ElectricCurrentGradientDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentGradientAmperePerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentGradientDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricCurrentGradient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerSecond by default.
func (a *ElectricCurrentGradient) ToDtoJSON(holdInUnit *ElectricCurrentGradientUnits) ([]byte, error) {
	// Convert to ElectricCurrentGradientDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricCurrentGradient to a specific unit value.
// The function uses the provided unit type (ElectricCurrentGradientUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricCurrentGradient) Convert(toUnit ElectricCurrentGradientUnits) float64 {
	switch toUnit { 
    case ElectricCurrentGradientAmperePerSecond:
		return a.AmperesPerSecond()
    case ElectricCurrentGradientAmperePerMinute:
		return a.AmperesPerMinute()
    case ElectricCurrentGradientAmperePerMillisecond:
		return a.AmperesPerMillisecond()
    case ElectricCurrentGradientAmperePerMicrosecond:
		return a.AmperesPerMicrosecond()
    case ElectricCurrentGradientAmperePerNanosecond:
		return a.AmperesPerNanosecond()
    case ElectricCurrentGradientMilliamperePerSecond:
		return a.MilliamperesPerSecond()
    case ElectricCurrentGradientMilliamperePerMinute:
		return a.MilliamperesPerMinute()
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentGradient) convertFromBase(toUnit ElectricCurrentGradientUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentGradientAmperePerSecond:
		return (value) 
	case ElectricCurrentGradientAmperePerMinute:
		return (value * 60) 
	case ElectricCurrentGradientAmperePerMillisecond:
		return (value / 1e3) 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return (value / 1e6) 
	case ElectricCurrentGradientAmperePerNanosecond:
		return (value / 1e9) 
	case ElectricCurrentGradientMilliamperePerSecond:
		return ((value) / 0.001) 
	case ElectricCurrentGradientMilliamperePerMinute:
		return ((value * 60) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentGradient) convertToBase(value float64, fromUnit ElectricCurrentGradientUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentGradientAmperePerSecond:
		return (value) 
	case ElectricCurrentGradientAmperePerMinute:
		return (value / 60) 
	case ElectricCurrentGradientAmperePerMillisecond:
		return (value * 1e3) 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return (value * 1e6) 
	case ElectricCurrentGradientAmperePerNanosecond:
		return (value * 1e9) 
	case ElectricCurrentGradientMilliamperePerSecond:
		return ((value) * 0.001) 
	case ElectricCurrentGradientMilliamperePerMinute:
		return ((value / 60) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricCurrentGradient in the default unit (AmperePerSecond),
// formatted to two decimal places.
func (a ElectricCurrentGradient) String() string {
	return a.ToString(ElectricCurrentGradientAmperePerSecond, 2)
}

// ToString formats the ElectricCurrentGradient value as a string with the specified unit and fractional digits.
// It converts the ElectricCurrentGradient to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricCurrentGradient value will be converted (e.g., AmperePerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricCurrentGradient with the unit abbreviation.
func (a *ElectricCurrentGradient) ToString(unit ElectricCurrentGradientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricCurrentGradientAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricCurrentGradientAbbreviation(unit))
}

// Equals checks if the given ElectricCurrentGradient is equal to the current ElectricCurrentGradient.
//
// Parameters:
//    other: The ElectricCurrentGradient to compare against.
//
// Returns:
//    bool: Returns true if both ElectricCurrentGradient are equal, false otherwise.
func (a *ElectricCurrentGradient) Equals(other *ElectricCurrentGradient) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricCurrentGradient with another ElectricCurrentGradient.
// It returns -1 if the current ElectricCurrentGradient is less than the other ElectricCurrentGradient, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricCurrentGradient to compare against.
//
// Returns:
//    int: -1 if the current ElectricCurrentGradient is less, 1 if greater, and 0 if equal.
func (a *ElectricCurrentGradient) CompareTo(other *ElectricCurrentGradient) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricCurrentGradient to the current ElectricCurrentGradient and returns the result.
// The result is a new ElectricCurrentGradient instance with the sum of the values.
//
// Parameters:
//    other: The ElectricCurrentGradient to add to the current ElectricCurrentGradient.
//
// Returns:
//    *ElectricCurrentGradient: A new ElectricCurrentGradient instance representing the sum of both ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Add(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricCurrentGradient from the current ElectricCurrentGradient and returns the result.
// The result is a new ElectricCurrentGradient instance with the difference of the values.
//
// Parameters:
//    other: The ElectricCurrentGradient to subtract from the current ElectricCurrentGradient.
//
// Returns:
//    *ElectricCurrentGradient: A new ElectricCurrentGradient instance representing the difference of both ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Subtract(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricCurrentGradient by the given ElectricCurrentGradient and returns the result.
// The result is a new ElectricCurrentGradient instance with the product of the values.
//
// Parameters:
//    other: The ElectricCurrentGradient to multiply with the current ElectricCurrentGradient.
//
// Returns:
//    *ElectricCurrentGradient: A new ElectricCurrentGradient instance representing the product of both ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Multiply(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricCurrentGradient by the given ElectricCurrentGradient and returns the result.
// The result is a new ElectricCurrentGradient instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricCurrentGradient to divide the current ElectricCurrentGradient by.
//
// Returns:
//    *ElectricCurrentGradient: A new ElectricCurrentGradient instance representing the quotient of both ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Divide(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value / other.BaseValue()}
}

// GetElectricCurrentGradientAbbreviation gets the unit abbreviation.
func GetElectricCurrentGradientAbbreviation(unit ElectricCurrentGradientUnits) string {
	switch unit { 
	case ElectricCurrentGradientAmperePerSecond:
		return "A/s" 
	case ElectricCurrentGradientAmperePerMinute:
		return "A/min" 
	case ElectricCurrentGradientAmperePerMillisecond:
		return "A/ms" 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return "A/μs" 
	case ElectricCurrentGradientAmperePerNanosecond:
		return "A/ns" 
	case ElectricCurrentGradientMilliamperePerSecond:
		return "mA/s" 
	case ElectricCurrentGradientMilliamperePerMinute:
		return "mA/min" 
	default:
		return ""
	}
}