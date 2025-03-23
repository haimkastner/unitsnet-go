// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LeakRateUnits defines various units of LeakRate.
type LeakRateUnits string

const (
    
        // 
        LeakRatePascalCubicMeterPerSecond LeakRateUnits = "PascalCubicMeterPerSecond"
        // 
        LeakRateMillibarLiterPerSecond LeakRateUnits = "MillibarLiterPerSecond"
        // 
        LeakRateTorrLiterPerSecond LeakRateUnits = "TorrLiterPerSecond"
)

var internalLeakRateUnitsMap = map[LeakRateUnits]bool{
	
	LeakRatePascalCubicMeterPerSecond: true,
	LeakRateMillibarLiterPerSecond: true,
	LeakRateTorrLiterPerSecond: true,
}

// LeakRateDto represents a LeakRate measurement with a numerical value and its corresponding unit.
type LeakRateDto struct {
    // Value is the numerical representation of the LeakRate.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the LeakRate, as defined in the LeakRateUnits enumeration.
	Unit  LeakRateUnits `json:"unit" validate:"required,oneof=PascalCubicMeterPerSecond MillibarLiterPerSecond TorrLiterPerSecond"`
}

// LeakRateDtoFactory groups methods for creating and serializing LeakRateDto objects.
type LeakRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LeakRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LeakRateDtoFactory) FromJSON(data []byte) (*LeakRateDto, error) {
	a := LeakRateDto{}

    // Parse JSON into LeakRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a LeakRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LeakRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// LeakRate represents a measurement in a of LeakRate.
//
// A leakage rate of QL = 1 Pa-m³/s is given when the pressure in a closed, evacuated container with a volume of 1 m³ rises by 1 Pa per second or when the pressure in the container drops by 1 Pa in the event of overpressure.
type LeakRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    pascal_cubic_meters_per_secondLazy *float64 
    millibar_liters_per_secondLazy *float64 
    torr_liters_per_secondLazy *float64 
}

// LeakRateFactory groups methods for creating LeakRate instances.
type LeakRateFactory struct{}

// CreateLeakRate creates a new LeakRate instance from the given value and unit.
func (uf LeakRateFactory) CreateLeakRate(value float64, unit LeakRateUnits) (*LeakRate, error) {
	return newLeakRate(value, unit)
}

// FromDto converts a LeakRateDto to a LeakRate instance.
func (uf LeakRateFactory) FromDto(dto LeakRateDto) (*LeakRate, error) {
	return newLeakRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a LeakRate instance.
func (uf LeakRateFactory) FromDtoJSON(data []byte) (*LeakRate, error) {
	unitDto, err := LeakRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LeakRateDto from JSON: %w", err)
	}
	return LeakRateFactory{}.FromDto(*unitDto)
}


// FromPascalCubicMetersPerSecond creates a new LeakRate instance from a value in PascalCubicMetersPerSecond.
func (uf LeakRateFactory) FromPascalCubicMetersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRatePascalCubicMeterPerSecond)
}

// FromMillibarLitersPerSecond creates a new LeakRate instance from a value in MillibarLitersPerSecond.
func (uf LeakRateFactory) FromMillibarLitersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRateMillibarLiterPerSecond)
}

// FromTorrLitersPerSecond creates a new LeakRate instance from a value in TorrLitersPerSecond.
func (uf LeakRateFactory) FromTorrLitersPerSecond(value float64) (*LeakRate, error) {
	return newLeakRate(value, LeakRateTorrLiterPerSecond)
}


// newLeakRate creates a new LeakRate.
func newLeakRate(value float64, fromUnit LeakRateUnits) (*LeakRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalLeakRateUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in LeakRateUnits", fromUnit)
	}
	a := &LeakRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LeakRate in PascalCubicMeterPerSecond unit (the base/default quantity).
func (a *LeakRate) BaseValue() float64 {
	return a.value
}


// PascalCubicMetersPerSecond returns the LeakRate value in PascalCubicMetersPerSecond.
//
// 
func (a *LeakRate) PascalCubicMetersPerSecond() float64 {
	if a.pascal_cubic_meters_per_secondLazy != nil {
		return *a.pascal_cubic_meters_per_secondLazy
	}
	pascal_cubic_meters_per_second := a.convertFromBase(LeakRatePascalCubicMeterPerSecond)
	a.pascal_cubic_meters_per_secondLazy = &pascal_cubic_meters_per_second
	return pascal_cubic_meters_per_second
}

// MillibarLitersPerSecond returns the LeakRate value in MillibarLitersPerSecond.
//
// 
func (a *LeakRate) MillibarLitersPerSecond() float64 {
	if a.millibar_liters_per_secondLazy != nil {
		return *a.millibar_liters_per_secondLazy
	}
	millibar_liters_per_second := a.convertFromBase(LeakRateMillibarLiterPerSecond)
	a.millibar_liters_per_secondLazy = &millibar_liters_per_second
	return millibar_liters_per_second
}

// TorrLitersPerSecond returns the LeakRate value in TorrLitersPerSecond.
//
// 
func (a *LeakRate) TorrLitersPerSecond() float64 {
	if a.torr_liters_per_secondLazy != nil {
		return *a.torr_liters_per_secondLazy
	}
	torr_liters_per_second := a.convertFromBase(LeakRateTorrLiterPerSecond)
	a.torr_liters_per_secondLazy = &torr_liters_per_second
	return torr_liters_per_second
}


// ToDto creates a LeakRateDto representation from the LeakRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalCubicMeterPerSecond by default.
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

// ToDtoJSON creates a JSON representation of the LeakRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalCubicMeterPerSecond by default.
func (a *LeakRate) ToDtoJSON(holdInUnit *LeakRateUnits) ([]byte, error) {
	// Convert to LeakRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a LeakRate to a specific unit value.
// The function uses the provided unit type (LeakRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *LeakRate) Convert(toUnit LeakRateUnits) float64 {
	switch toUnit { 
    case LeakRatePascalCubicMeterPerSecond:
		return a.PascalCubicMetersPerSecond()
    case LeakRateMillibarLiterPerSecond:
		return a.MillibarLitersPerSecond()
    case LeakRateTorrLiterPerSecond:
		return a.TorrLitersPerSecond()
	default:
		return math.NaN()
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

// String returns a string representation of the LeakRate in the default unit (PascalCubicMeterPerSecond),
// formatted to two decimal places.
func (a LeakRate) String() string {
	return a.ToString(LeakRatePascalCubicMeterPerSecond, 2)
}

// ToString formats the LeakRate value as a string with the specified unit and fractional digits.
// It converts the LeakRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the LeakRate value will be converted (e.g., PascalCubicMeterPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the LeakRate with the unit abbreviation.
func (a *LeakRate) ToString(unit LeakRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetLeakRateAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetLeakRateAbbreviation(unit))
}

// Equals checks if the given LeakRate is equal to the current LeakRate.
//
// Parameters:
//    other: The LeakRate to compare against.
//
// Returns:
//    bool: Returns true if both LeakRate are equal, false otherwise.
func (a *LeakRate) Equals(other *LeakRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current LeakRate with another LeakRate.
// It returns -1 if the current LeakRate is less than the other LeakRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The LeakRate to compare against.
//
// Returns:
//    int: -1 if the current LeakRate is less, 1 if greater, and 0 if equal.
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

// Add adds the given LeakRate to the current LeakRate and returns the result.
// The result is a new LeakRate instance with the sum of the values.
//
// Parameters:
//    other: The LeakRate to add to the current LeakRate.
//
// Returns:
//    *LeakRate: A new LeakRate instance representing the sum of both LeakRate.
func (a *LeakRate) Add(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given LeakRate from the current LeakRate and returns the result.
// The result is a new LeakRate instance with the difference of the values.
//
// Parameters:
//    other: The LeakRate to subtract from the current LeakRate.
//
// Returns:
//    *LeakRate: A new LeakRate instance representing the difference of both LeakRate.
func (a *LeakRate) Subtract(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current LeakRate by the given LeakRate and returns the result.
// The result is a new LeakRate instance with the product of the values.
//
// Parameters:
//    other: The LeakRate to multiply with the current LeakRate.
//
// Returns:
//    *LeakRate: A new LeakRate instance representing the product of both LeakRate.
func (a *LeakRate) Multiply(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value * other.BaseValue()}
}

// Divide divides the current LeakRate by the given LeakRate and returns the result.
// The result is a new LeakRate instance with the quotient of the values.
//
// Parameters:
//    other: The LeakRate to divide the current LeakRate by.
//
// Returns:
//    *LeakRate: A new LeakRate instance representing the quotient of both LeakRate.
func (a *LeakRate) Divide(other *LeakRate) *LeakRate {
	return &LeakRate{value: a.value / other.BaseValue()}
}

// GetLeakRateAbbreviation gets the unit abbreviation.
func GetLeakRateAbbreviation(unit LeakRateUnits) string {
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