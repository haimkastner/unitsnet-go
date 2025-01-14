// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RatioChangeRateUnits defines various units of RatioChangeRate.
type RatioChangeRateUnits string

const (
    
        // 
        RatioChangeRatePercentPerSecond RatioChangeRateUnits = "PercentPerSecond"
        // 
        RatioChangeRateDecimalFractionPerSecond RatioChangeRateUnits = "DecimalFractionPerSecond"
)

// RatioChangeRateDto represents a RatioChangeRate measurement with a numerical value and its corresponding unit.
type RatioChangeRateDto struct {
    // Value is the numerical representation of the RatioChangeRate.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RatioChangeRate, as defined in the RatioChangeRateUnits enumeration.
	Unit  RatioChangeRateUnits `json:"unit"`
}

// RatioChangeRateDtoFactory groups methods for creating and serializing RatioChangeRateDto objects.
type RatioChangeRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RatioChangeRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RatioChangeRateDtoFactory) FromJSON(data []byte) (*RatioChangeRateDto, error) {
	a := RatioChangeRateDto{}

    // Parse JSON into RatioChangeRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RatioChangeRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RatioChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RatioChangeRate represents a measurement in a of RatioChangeRate.
//
// The change in ratio per unit of time.
type RatioChangeRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    percents_per_secondLazy *float64 
    decimal_fractions_per_secondLazy *float64 
}

// RatioChangeRateFactory groups methods for creating RatioChangeRate instances.
type RatioChangeRateFactory struct{}

// CreateRatioChangeRate creates a new RatioChangeRate instance from the given value and unit.
func (uf RatioChangeRateFactory) CreateRatioChangeRate(value float64, unit RatioChangeRateUnits) (*RatioChangeRate, error) {
	return newRatioChangeRate(value, unit)
}

// FromDto converts a RatioChangeRateDto to a RatioChangeRate instance.
func (uf RatioChangeRateFactory) FromDto(dto RatioChangeRateDto) (*RatioChangeRate, error) {
	return newRatioChangeRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RatioChangeRate instance.
func (uf RatioChangeRateFactory) FromDtoJSON(data []byte) (*RatioChangeRate, error) {
	unitDto, err := RatioChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RatioChangeRateDto from JSON: %w", err)
	}
	return RatioChangeRateFactory{}.FromDto(*unitDto)
}


// FromPercentsPerSecond creates a new RatioChangeRate instance from a value in PercentsPerSecond.
func (uf RatioChangeRateFactory) FromPercentsPerSecond(value float64) (*RatioChangeRate, error) {
	return newRatioChangeRate(value, RatioChangeRatePercentPerSecond)
}

// FromDecimalFractionsPerSecond creates a new RatioChangeRate instance from a value in DecimalFractionsPerSecond.
func (uf RatioChangeRateFactory) FromDecimalFractionsPerSecond(value float64) (*RatioChangeRate, error) {
	return newRatioChangeRate(value, RatioChangeRateDecimalFractionPerSecond)
}


// newRatioChangeRate creates a new RatioChangeRate.
func newRatioChangeRate(value float64, fromUnit RatioChangeRateUnits) (*RatioChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RatioChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RatioChangeRate in DecimalFractionPerSecond unit (the base/default quantity).
func (a *RatioChangeRate) BaseValue() float64 {
	return a.value
}


// PercentsPerSecond returns the RatioChangeRate value in PercentsPerSecond.
//
// 
func (a *RatioChangeRate) PercentsPerSecond() float64 {
	if a.percents_per_secondLazy != nil {
		return *a.percents_per_secondLazy
	}
	percents_per_second := a.convertFromBase(RatioChangeRatePercentPerSecond)
	a.percents_per_secondLazy = &percents_per_second
	return percents_per_second
}

// DecimalFractionsPerSecond returns the RatioChangeRate value in DecimalFractionsPerSecond.
//
// 
func (a *RatioChangeRate) DecimalFractionsPerSecond() float64 {
	if a.decimal_fractions_per_secondLazy != nil {
		return *a.decimal_fractions_per_secondLazy
	}
	decimal_fractions_per_second := a.convertFromBase(RatioChangeRateDecimalFractionPerSecond)
	a.decimal_fractions_per_secondLazy = &decimal_fractions_per_second
	return decimal_fractions_per_second
}


// ToDto creates a RatioChangeRateDto representation from the RatioChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFractionPerSecond by default.
func (a *RatioChangeRate) ToDto(holdInUnit *RatioChangeRateUnits) RatioChangeRateDto {
	if holdInUnit == nil {
		defaultUnit := RatioChangeRateDecimalFractionPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return RatioChangeRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RatioChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFractionPerSecond by default.
func (a *RatioChangeRate) ToDtoJSON(holdInUnit *RatioChangeRateUnits) ([]byte, error) {
	// Convert to RatioChangeRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RatioChangeRate to a specific unit value.
// The function uses the provided unit type (RatioChangeRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RatioChangeRate) Convert(toUnit RatioChangeRateUnits) float64 {
	switch toUnit { 
    case RatioChangeRatePercentPerSecond:
		return a.PercentsPerSecond()
    case RatioChangeRateDecimalFractionPerSecond:
		return a.DecimalFractionsPerSecond()
	default:
		return math.NaN()
	}
}

func (a *RatioChangeRate) convertFromBase(toUnit RatioChangeRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case RatioChangeRatePercentPerSecond:
		return (value * 1e2) 
	case RatioChangeRateDecimalFractionPerSecond:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *RatioChangeRate) convertToBase(value float64, fromUnit RatioChangeRateUnits) float64 {
	switch fromUnit { 
	case RatioChangeRatePercentPerSecond:
		return (value / 1e2) 
	case RatioChangeRateDecimalFractionPerSecond:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RatioChangeRate in the default unit (DecimalFractionPerSecond),
// formatted to two decimal places.
func (a RatioChangeRate) String() string {
	return a.ToString(RatioChangeRateDecimalFractionPerSecond, 2)
}

// ToString formats the RatioChangeRate value as a string with the specified unit and fractional digits.
// It converts the RatioChangeRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RatioChangeRate value will be converted (e.g., DecimalFractionPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RatioChangeRate with the unit abbreviation.
func (a *RatioChangeRate) ToString(unit RatioChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RatioChangeRate) getUnitAbbreviation(unit RatioChangeRateUnits) string {
	switch unit { 
	case RatioChangeRatePercentPerSecond:
		return "%/s" 
	case RatioChangeRateDecimalFractionPerSecond:
		return "/s" 
	default:
		return ""
	}
}

// Equals checks if the given RatioChangeRate is equal to the current RatioChangeRate.
//
// Parameters:
//    other: The RatioChangeRate to compare against.
//
// Returns:
//    bool: Returns true if both RatioChangeRate are equal, false otherwise.
func (a *RatioChangeRate) Equals(other *RatioChangeRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RatioChangeRate with another RatioChangeRate.
// It returns -1 if the current RatioChangeRate is less than the other RatioChangeRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RatioChangeRate to compare against.
//
// Returns:
//    int: -1 if the current RatioChangeRate is less, 1 if greater, and 0 if equal.
func (a *RatioChangeRate) CompareTo(other *RatioChangeRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RatioChangeRate to the current RatioChangeRate and returns the result.
// The result is a new RatioChangeRate instance with the sum of the values.
//
// Parameters:
//    other: The RatioChangeRate to add to the current RatioChangeRate.
//
// Returns:
//    *RatioChangeRate: A new RatioChangeRate instance representing the sum of both RatioChangeRate.
func (a *RatioChangeRate) Add(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RatioChangeRate from the current RatioChangeRate and returns the result.
// The result is a new RatioChangeRate instance with the difference of the values.
//
// Parameters:
//    other: The RatioChangeRate to subtract from the current RatioChangeRate.
//
// Returns:
//    *RatioChangeRate: A new RatioChangeRate instance representing the difference of both RatioChangeRate.
func (a *RatioChangeRate) Subtract(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RatioChangeRate by the given RatioChangeRate and returns the result.
// The result is a new RatioChangeRate instance with the product of the values.
//
// Parameters:
//    other: The RatioChangeRate to multiply with the current RatioChangeRate.
//
// Returns:
//    *RatioChangeRate: A new RatioChangeRate instance representing the product of both RatioChangeRate.
func (a *RatioChangeRate) Multiply(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value * other.BaseValue()}
}

// Divide divides the current RatioChangeRate by the given RatioChangeRate and returns the result.
// The result is a new RatioChangeRate instance with the quotient of the values.
//
// Parameters:
//    other: The RatioChangeRate to divide the current RatioChangeRate by.
//
// Returns:
//    *RatioChangeRate: A new RatioChangeRate instance representing the quotient of both RatioChangeRate.
func (a *RatioChangeRate) Divide(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value / other.BaseValue()}
}