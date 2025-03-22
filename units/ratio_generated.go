// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RatioUnits defines various units of Ratio.
type RatioUnits string

const (
    
        // 
        RatioDecimalFraction RatioUnits = "DecimalFraction"
        // 
        RatioPercent RatioUnits = "Percent"
        // 
        RatioPartPerThousand RatioUnits = "PartPerThousand"
        // 
        RatioPartPerMillion RatioUnits = "PartPerMillion"
        // 
        RatioPartPerBillion RatioUnits = "PartPerBillion"
        // 
        RatioPartPerTrillion RatioUnits = "PartPerTrillion"
)

var internalRatioUnitsMap = map[RatioUnits]bool{
	
	RatioDecimalFraction: true,
	RatioPercent: true,
	RatioPartPerThousand: true,
	RatioPartPerMillion: true,
	RatioPartPerBillion: true,
	RatioPartPerTrillion: true,
}

// RatioDto represents a Ratio measurement with a numerical value and its corresponding unit.
type RatioDto struct {
    // Value is the numerical representation of the Ratio.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Ratio, as defined in the RatioUnits enumeration.
	Unit  RatioUnits `json:"unit" validate:"required,oneof=DecimalFraction,Percent,PartPerThousand,PartPerMillion,PartPerBillion,PartPerTrillion"`
}

// RatioDtoFactory groups methods for creating and serializing RatioDto objects.
type RatioDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RatioDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RatioDtoFactory) FromJSON(data []byte) (*RatioDto, error) {
	a := RatioDto{}

    // Parse JSON into RatioDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a RatioDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Ratio represents a measurement in a of Ratio.
//
// In mathematics, a ratio is a relationship between two numbers of the same kind (e.g., objects, persons, students, spoonfuls, units of whatever identical dimension), usually expressed as "a to b" or a:b, sometimes expressed arithmetically as a dimensionless quotient of the two that explicitly indicates how many times the first number contains the second (not necessarily an integer).
type Ratio struct {
	// value is the base measurement stored internally.
	value       float64
    
    decimal_fractionsLazy *float64 
    percentLazy *float64 
    parts_per_thousandLazy *float64 
    parts_per_millionLazy *float64 
    parts_per_billionLazy *float64 
    parts_per_trillionLazy *float64 
}

// RatioFactory groups methods for creating Ratio instances.
type RatioFactory struct{}

// CreateRatio creates a new Ratio instance from the given value and unit.
func (uf RatioFactory) CreateRatio(value float64, unit RatioUnits) (*Ratio, error) {
	return newRatio(value, unit)
}

// FromDto converts a RatioDto to a Ratio instance.
func (uf RatioFactory) FromDto(dto RatioDto) (*Ratio, error) {
	return newRatio(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Ratio instance.
func (uf RatioFactory) FromDtoJSON(data []byte) (*Ratio, error) {
	unitDto, err := RatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RatioDto from JSON: %w", err)
	}
	return RatioFactory{}.FromDto(*unitDto)
}


// FromDecimalFractions creates a new Ratio instance from a value in DecimalFractions.
func (uf RatioFactory) FromDecimalFractions(value float64) (*Ratio, error) {
	return newRatio(value, RatioDecimalFraction)
}

// FromPercent creates a new Ratio instance from a value in Percent.
func (uf RatioFactory) FromPercent(value float64) (*Ratio, error) {
	return newRatio(value, RatioPercent)
}

// FromPartsPerThousand creates a new Ratio instance from a value in PartsPerThousand.
func (uf RatioFactory) FromPartsPerThousand(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerThousand)
}

// FromPartsPerMillion creates a new Ratio instance from a value in PartsPerMillion.
func (uf RatioFactory) FromPartsPerMillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerMillion)
}

// FromPartsPerBillion creates a new Ratio instance from a value in PartsPerBillion.
func (uf RatioFactory) FromPartsPerBillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerBillion)
}

// FromPartsPerTrillion creates a new Ratio instance from a value in PartsPerTrillion.
func (uf RatioFactory) FromPartsPerTrillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerTrillion)
}


// newRatio creates a new Ratio.
func newRatio(value float64, fromUnit RatioUnits) (*Ratio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalRatioUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in RatioUnits", fromUnit)
	}
	a := &Ratio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Ratio in DecimalFraction unit (the base/default quantity).
func (a *Ratio) BaseValue() float64 {
	return a.value
}


// DecimalFractions returns the Ratio value in DecimalFractions.
//
// 
func (a *Ratio) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(RatioDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// Percent returns the Ratio value in Percent.
//
// 
func (a *Ratio) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(RatioPercent)
	a.percentLazy = &percent
	return percent
}

// PartsPerThousand returns the Ratio value in PartsPerThousand.
//
// 
func (a *Ratio) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(RatioPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartsPerMillion returns the Ratio value in PartsPerMillion.
//
// 
func (a *Ratio) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(RatioPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartsPerBillion returns the Ratio value in PartsPerBillion.
//
// 
func (a *Ratio) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(RatioPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartsPerTrillion returns the Ratio value in PartsPerTrillion.
//
// 
func (a *Ratio) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(RatioPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}


// ToDto creates a RatioDto representation from the Ratio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *Ratio) ToDto(holdInUnit *RatioUnits) RatioDto {
	if holdInUnit == nil {
		defaultUnit := RatioDecimalFraction // Default value
		holdInUnit = &defaultUnit
	}

	return RatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Ratio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *Ratio) ToDtoJSON(holdInUnit *RatioUnits) ([]byte, error) {
	// Convert to RatioDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Ratio to a specific unit value.
// The function uses the provided unit type (RatioUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Ratio) Convert(toUnit RatioUnits) float64 {
	switch toUnit { 
    case RatioDecimalFraction:
		return a.DecimalFractions()
    case RatioPercent:
		return a.Percent()
    case RatioPartPerThousand:
		return a.PartsPerThousand()
    case RatioPartPerMillion:
		return a.PartsPerMillion()
    case RatioPartPerBillion:
		return a.PartsPerBillion()
    case RatioPartPerTrillion:
		return a.PartsPerTrillion()
	default:
		return math.NaN()
	}
}

func (a *Ratio) convertFromBase(toUnit RatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case RatioDecimalFraction:
		return (value) 
	case RatioPercent:
		return (value * 1e2) 
	case RatioPartPerThousand:
		return (value * 1e3) 
	case RatioPartPerMillion:
		return (value * 1e6) 
	case RatioPartPerBillion:
		return (value * 1e9) 
	case RatioPartPerTrillion:
		return (value * 1e12) 
	default:
		return math.NaN()
	}
}

func (a *Ratio) convertToBase(value float64, fromUnit RatioUnits) float64 {
	switch fromUnit { 
	case RatioDecimalFraction:
		return (value) 
	case RatioPercent:
		return (value / 1e2) 
	case RatioPartPerThousand:
		return (value / 1e3) 
	case RatioPartPerMillion:
		return (value / 1e6) 
	case RatioPartPerBillion:
		return (value / 1e9) 
	case RatioPartPerTrillion:
		return (value / 1e12) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Ratio in the default unit (DecimalFraction),
// formatted to two decimal places.
func (a Ratio) String() string {
	return a.ToString(RatioDecimalFraction, 2)
}

// ToString formats the Ratio value as a string with the specified unit and fractional digits.
// It converts the Ratio to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Ratio value will be converted (e.g., DecimalFraction))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Ratio with the unit abbreviation.
func (a *Ratio) ToString(unit RatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRatioAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRatioAbbreviation(unit))
}

// Equals checks if the given Ratio is equal to the current Ratio.
//
// Parameters:
//    other: The Ratio to compare against.
//
// Returns:
//    bool: Returns true if both Ratio are equal, false otherwise.
func (a *Ratio) Equals(other *Ratio) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Ratio with another Ratio.
// It returns -1 if the current Ratio is less than the other Ratio, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Ratio to compare against.
//
// Returns:
//    int: -1 if the current Ratio is less, 1 if greater, and 0 if equal.
func (a *Ratio) CompareTo(other *Ratio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Ratio to the current Ratio and returns the result.
// The result is a new Ratio instance with the sum of the values.
//
// Parameters:
//    other: The Ratio to add to the current Ratio.
//
// Returns:
//    *Ratio: A new Ratio instance representing the sum of both Ratio.
func (a *Ratio) Add(other *Ratio) *Ratio {
	return &Ratio{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Ratio from the current Ratio and returns the result.
// The result is a new Ratio instance with the difference of the values.
//
// Parameters:
//    other: The Ratio to subtract from the current Ratio.
//
// Returns:
//    *Ratio: A new Ratio instance representing the difference of both Ratio.
func (a *Ratio) Subtract(other *Ratio) *Ratio {
	return &Ratio{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Ratio by the given Ratio and returns the result.
// The result is a new Ratio instance with the product of the values.
//
// Parameters:
//    other: The Ratio to multiply with the current Ratio.
//
// Returns:
//    *Ratio: A new Ratio instance representing the product of both Ratio.
func (a *Ratio) Multiply(other *Ratio) *Ratio {
	return &Ratio{value: a.value * other.BaseValue()}
}

// Divide divides the current Ratio by the given Ratio and returns the result.
// The result is a new Ratio instance with the quotient of the values.
//
// Parameters:
//    other: The Ratio to divide the current Ratio by.
//
// Returns:
//    *Ratio: A new Ratio instance representing the quotient of both Ratio.
func (a *Ratio) Divide(other *Ratio) *Ratio {
	return &Ratio{value: a.value / other.BaseValue()}
}

// GetRatioAbbreviation gets the unit abbreviation.
func GetRatioAbbreviation(unit RatioUnits) string {
	switch unit { 
	case RatioDecimalFraction:
		return "" 
	case RatioPercent:
		return "%" 
	case RatioPartPerThousand:
		return "‰" 
	case RatioPartPerMillion:
		return "ppm" 
	case RatioPartPerBillion:
		return "ppb" 
	case RatioPartPerTrillion:
		return "ppt" 
	default:
		return ""
	}
}