// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RatioChangeRateUnits enumeration
type RatioChangeRateUnits string

const (
    
        // 
        RatioChangeRatePercentPerSecond RatioChangeRateUnits = "PercentPerSecond"
        // 
        RatioChangeRateDecimalFractionPerSecond RatioChangeRateUnits = "DecimalFractionPerSecond"
)

// RatioChangeRateDto represents an RatioChangeRate
type RatioChangeRateDto struct {
	Value float64
	Unit  RatioChangeRateUnits
}

// RatioChangeRateDtoFactory struct to group related functions
type RatioChangeRateDtoFactory struct{}

func (udf RatioChangeRateDtoFactory) FromJSON(data []byte) (*RatioChangeRateDto, error) {
	a := RatioChangeRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RatioChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RatioChangeRate struct
type RatioChangeRate struct {
	value       float64
    
    percents_per_secondLazy *float64 
    decimal_fractions_per_secondLazy *float64 
}

// RatioChangeRateFactory struct to group related functions
type RatioChangeRateFactory struct{}

func (uf RatioChangeRateFactory) CreateRatioChangeRate(value float64, unit RatioChangeRateUnits) (*RatioChangeRate, error) {
	return newRatioChangeRate(value, unit)
}

func (uf RatioChangeRateFactory) FromDto(dto RatioChangeRateDto) (*RatioChangeRate, error) {
	return newRatioChangeRate(dto.Value, dto.Unit)
}

func (uf RatioChangeRateFactory) FromDtoJSON(data []byte) (*RatioChangeRate, error) {
	unitDto, err := RatioChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RatioChangeRateFactory{}.FromDto(*unitDto)
}


// FromPercentPerSecond creates a new RatioChangeRate instance from PercentPerSecond.
func (uf RatioChangeRateFactory) FromPercentsPerSecond(value float64) (*RatioChangeRate, error) {
	return newRatioChangeRate(value, RatioChangeRatePercentPerSecond)
}

// FromDecimalFractionPerSecond creates a new RatioChangeRate instance from DecimalFractionPerSecond.
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

// BaseValue returns the base value of RatioChangeRate in DecimalFractionPerSecond.
func (a *RatioChangeRate) BaseValue() float64 {
	return a.value
}


// PercentPerSecond returns the value in PercentPerSecond.
func (a *RatioChangeRate) PercentsPerSecond() float64 {
	if a.percents_per_secondLazy != nil {
		return *a.percents_per_secondLazy
	}
	percents_per_second := a.convertFromBase(RatioChangeRatePercentPerSecond)
	a.percents_per_secondLazy = &percents_per_second
	return percents_per_second
}

// DecimalFractionPerSecond returns the value in DecimalFractionPerSecond.
func (a *RatioChangeRate) DecimalFractionsPerSecond() float64 {
	if a.decimal_fractions_per_secondLazy != nil {
		return *a.decimal_fractions_per_secondLazy
	}
	decimal_fractions_per_second := a.convertFromBase(RatioChangeRateDecimalFractionPerSecond)
	a.decimal_fractions_per_secondLazy = &decimal_fractions_per_second
	return decimal_fractions_per_second
}


// ToDto creates an RatioChangeRateDto representation.
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

// ToDtoJSON creates an RatioChangeRateDto representation.
func (a *RatioChangeRate) ToDtoJSON(holdInUnit *RatioChangeRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RatioChangeRate to a specific unit value.
func (a *RatioChangeRate) Convert(toUnit RatioChangeRateUnits) float64 {
	switch toUnit { 
    case RatioChangeRatePercentPerSecond:
		return a.PercentsPerSecond()
    case RatioChangeRateDecimalFractionPerSecond:
		return a.DecimalFractionsPerSecond()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a RatioChangeRate) String() string {
	return a.ToString(RatioChangeRateDecimalFractionPerSecond, 2)
}

// ToString formats the RatioChangeRate to string.
// fractionalDigits -1 for not mention
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

// Check if the given RatioChangeRate are equals to the current RatioChangeRate
func (a *RatioChangeRate) Equals(other *RatioChangeRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given RatioChangeRate are equals to the current RatioChangeRate
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

// Add the given RatioChangeRate to the current RatioChangeRate.
func (a *RatioChangeRate) Add(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value + other.BaseValue()}
}

// Subtract the given RatioChangeRate to the current RatioChangeRate.
func (a *RatioChangeRate) Subtract(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value - other.BaseValue()}
}

// Multiply the given RatioChangeRate to the current RatioChangeRate.
func (a *RatioChangeRate) Multiply(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value * other.BaseValue()}
}

// Divide the given RatioChangeRate to the current RatioChangeRate.
func (a *RatioChangeRate) Divide(other *RatioChangeRate) *RatioChangeRate {
	return &RatioChangeRate{value: a.value / other.BaseValue()}
}