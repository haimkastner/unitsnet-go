// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LeakRateUnits enumeration
type LeakRateUnits string

const (
    
        // 
        LeakRatePascalCubicMeterPerSecond LeakRateUnits = "PascalCubicMeterPerSecond"
        // 
        LeakRateMillibarLiterPerSecond LeakRateUnits = "MillibarLiterPerSecond"
        // 
        LeakRateTorrLiterPerSecond LeakRateUnits = "TorrLiterPerSecond"
)

// LeakRateDto represents an LeakRate
type LeakRateDto struct {
	Value float64
	Unit  LeakRateUnits
}

// LeakRateDtoFactory struct to group related functions
type LeakRateDtoFactory struct{}

func (udf LeakRateDtoFactory) FromJSON(data []byte) (*LeakRateDto, error) {
	a := LeakRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LeakRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// LeakRate struct
type LeakRate struct {
	value       float64
    
    pascal_cubic_meters_per_secondLazy *float64 
    millibar_liters_per_secondLazy *float64 
    torr_liters_per_secondLazy *float64 
}

// LeakRateFactory struct to group related functions
type LeakRateFactory struct{}

func (uf LeakRateFactory) CreateLeakRate(value float64, unit LeakRateUnits) (*LeakRate, error) {
	return newLeakRate(value, unit)
}

func (uf LeakRateFactory) FromDto(dto LeakRateDto) (*LeakRate, error) {
	return newLeakRate(dto.Value, dto.Unit)
}

func (uf LeakRateFactory) FromDtoJSON(data []byte) (*LeakRate, error) {
	unitDto, err := LeakRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LeakRateFactory{}.FromDto(*unitDto)
}


// FromPascalCubicMeterPerSecond creates a new LeakRate instance from PascalCubicMeterPerSecond.
func (uf LeakRateFactory) FromPascalCubicMetersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRatePascalCubicMeterPerSecond)
}

// FromMillibarLiterPerSecond creates a new LeakRate instance from MillibarLiterPerSecond.
func (uf LeakRateFactory) FromMillibarLitersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRateMillibarLiterPerSecond)
}

// FromTorrLiterPerSecond creates a new LeakRate instance from TorrLiterPerSecond.
func (uf LeakRateFactory) FromTorrLitersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRateTorrLiterPerSecond)
}




// newLeakRate creates a new LeakRate.
func newLeakRate(value float64, fromUnit LeakRateUnits) (*LeakRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LeakRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LeakRate in PascalCubicMeterPerSecond.
func (a *LeakRate) BaseValue() float64 {
	return a.value
}


// PascalCubicMeterPerSecond returns the value in PascalCubicMeterPerSecond.
func (a *LeakRate) PascalCubicMetersPerSecond() float64 {
	if a.pascal_cubic_meters_per_secondLazy != nil {
		return *a.pascal_cubic_meters_per_secondLazy
	}
	pascal_cubic_meters_per_second := a.convertFromBase(LeakRatePascalCubicMeterPerSecond)
	a.pascal_cubic_meters_per_secondLazy = &pascal_cubic_meters_per_second
	return pascal_cubic_meters_per_second
}

// MillibarLiterPerSecond returns the value in MillibarLiterPerSecond.
func (a *LeakRate) MillibarLitersPerSecond() float64 {
	if a.millibar_liters_per_secondLazy != nil {
		return *a.millibar_liters_per_secondLazy
	}
	millibar_liters_per_second := a.convertFromBase(LeakRateMillibarLiterPerSecond)
	a.millibar_liters_per_secondLazy = &millibar_liters_per_second
	return millibar_liters_per_second
}

// TorrLiterPerSecond returns the value in TorrLiterPerSecond.
func (a *LeakRate) TorrLitersPerSecond() float64 {
	if a.torr_liters_per_secondLazy != nil {
		return *a.torr_liters_per_secondLazy
	}
	torr_liters_per_second := a.convertFromBase(LeakRateTorrLiterPerSecond)
	a.torr_liters_per_secondLazy = &torr_liters_per_second
	return torr_liters_per_second
}


// ToDto creates an LeakRateDto representation.
func (a *LeakRate) ToDto(holdInUnit *LeakRateUnits) LeakRateDto {
	if holdInUnit == nil {
		defaultUnit := LeakRatePascalCubicMeterPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return LeakRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LeakRateDto representation.
func (a *LeakRate) ToDtoJSON(holdInUnit *LeakRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts LeakRate to a specific unit value.
func (a *LeakRate) Convert(toUnit LeakRateUnits) float64 {
	switch toUnit { 
    case LeakRatePascalCubicMeterPerSecond:
		return a.PascalCubicMetersPerSecond()
    case LeakRateMillibarLiterPerSecond:
		return a.MillibarLitersPerSecond()
    case LeakRateTorrLiterPerSecond:
		return a.TorrLitersPerSecond()
	default:
		return 0
	}
}

func (a *LeakRate) convertFromBase(toUnit LeakRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case LeakRatePascalCubicMeterPerSecond:
		return (value) 
	case LeakRateMillibarLiterPerSecond:
		return (value * 10) 
	case LeakRateTorrLiterPerSecond:
		return (value * 7.5) 
	default:
		return math.NaN()
	}
}

func (a *LeakRate) convertToBase(value float64, fromUnit LeakRateUnits) float64 {
	switch fromUnit { 
	case LeakRatePascalCubicMeterPerSecond:
		return (value) 
	case LeakRateMillibarLiterPerSecond:
		return (value / 10) 
	case LeakRateTorrLiterPerSecond:
		return (value / 7.5) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a LeakRate) String() string {
	return a.ToString(LeakRatePascalCubicMeterPerSecond, 2)
}

// ToString formats the LeakRate to string.
// fractionalDigits -1 for not mention
func (a *LeakRate) ToString(unit LeakRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LeakRate) getUnitAbbreviation(unit LeakRateUnits) string {
	switch unit { 
	case LeakRatePascalCubicMeterPerSecond:
		return "Pa·m³/s" 
	case LeakRateMillibarLiterPerSecond:
		return "mbar·l/s" 
	case LeakRateTorrLiterPerSecond:
		return "Torr·l/s" 
	default:
		return ""
	}
}

// Check if the given LeakRate are equals to the current LeakRate
func (a *LeakRate) Equals(other *LeakRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given LeakRate are equals to the current LeakRate
func (a *LeakRate) CompareTo(other *LeakRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given LeakRate to the current LeakRate.
func (a *LeakRate) Add(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value + other.BaseValue()}
}

// Subtract the given LeakRate to the current LeakRate.
func (a *LeakRate) Subtract(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value - other.BaseValue()}
}

// Multiply the given LeakRate to the current LeakRate.
func (a *LeakRate) Multiply(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value * other.BaseValue()}
}

// Divide the given LeakRate to the current LeakRate.
func (a *LeakRate) Divide(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value / other.BaseValue()}
}