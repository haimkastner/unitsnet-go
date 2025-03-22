// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AreaDensityUnits defines various units of AreaDensity.
type AreaDensityUnits string

const (
    
        // 
        AreaDensityKilogramPerSquareMeter AreaDensityUnits = "KilogramPerSquareMeter"
        // Also known as grammage for paper industry. In fiber industry used with abbreviation 'gsm'.
        AreaDensityGramPerSquareMeter AreaDensityUnits = "GramPerSquareMeter"
        // 
        AreaDensityMilligramPerSquareMeter AreaDensityUnits = "MilligramPerSquareMeter"
)

var internalAreaDensityUnitsMap = map[AreaDensityUnits]bool{
	
	AreaDensityKilogramPerSquareMeter: true,
	AreaDensityGramPerSquareMeter: true,
	AreaDensityMilligramPerSquareMeter: true,
}

// AreaDensityDto represents a AreaDensity measurement with a numerical value and its corresponding unit.
type AreaDensityDto struct {
    // Value is the numerical representation of the AreaDensity.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the AreaDensity, as defined in the AreaDensityUnits enumeration.
	Unit  AreaDensityUnits `json:"unit" validate:"required,oneof=KilogramPerSquareMeter,GramPerSquareMeter,MilligramPerSquareMeter"`
}

// AreaDensityDtoFactory groups methods for creating and serializing AreaDensityDto objects.
type AreaDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AreaDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AreaDensityDtoFactory) FromJSON(data []byte) (*AreaDensityDto, error) {
	a := AreaDensityDto{}

    // Parse JSON into AreaDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AreaDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AreaDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// AreaDensity represents a measurement in a of AreaDensity.
//
// The area density of a two-dimensional object is calculated as the mass per unit area. For paper this is also called grammage.
type AreaDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    kilograms_per_square_meterLazy *float64 
    grams_per_square_meterLazy *float64 
    milligrams_per_square_meterLazy *float64 
}

// AreaDensityFactory groups methods for creating AreaDensity instances.
type AreaDensityFactory struct{}

// CreateAreaDensity creates a new AreaDensity instance from the given value and unit.
func (uf AreaDensityFactory) CreateAreaDensity(value float64, unit AreaDensityUnits) (*AreaDensity, error) {
	return newAreaDensity(value, unit)
}

// FromDto converts a AreaDensityDto to a AreaDensity instance.
func (uf AreaDensityFactory) FromDto(dto AreaDensityDto) (*AreaDensity, error) {
	return newAreaDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AreaDensity instance.
func (uf AreaDensityFactory) FromDtoJSON(data []byte) (*AreaDensity, error) {
	unitDto, err := AreaDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AreaDensityDto from JSON: %w", err)
	}
	return AreaDensityFactory{}.FromDto(*unitDto)
}


// FromKilogramsPerSquareMeter creates a new AreaDensity instance from a value in KilogramsPerSquareMeter.
func (uf AreaDensityFactory) FromKilogramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityKilogramPerSquareMeter)
}

// FromGramsPerSquareMeter creates a new AreaDensity instance from a value in GramsPerSquareMeter.
func (uf AreaDensityFactory) FromGramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityGramPerSquareMeter)
}

// FromMilligramsPerSquareMeter creates a new AreaDensity instance from a value in MilligramsPerSquareMeter.
func (uf AreaDensityFactory) FromMilligramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityMilligramPerSquareMeter)
}


// newAreaDensity creates a new AreaDensity.
func newAreaDensity(value float64, fromUnit AreaDensityUnits) (*AreaDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalAreaDensityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in AreaDensityUnits", fromUnit)
	}
	a := &AreaDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AreaDensity in KilogramPerSquareMeter unit (the base/default quantity).
func (a *AreaDensity) BaseValue() float64 {
	return a.value
}


// KilogramsPerSquareMeter returns the AreaDensity value in KilogramsPerSquareMeter.
//
// 
func (a *AreaDensity) KilogramsPerSquareMeter() float64 {
	if a.kilograms_per_square_meterLazy != nil {
		return *a.kilograms_per_square_meterLazy
	}
	kilograms_per_square_meter := a.convertFromBase(AreaDensityKilogramPerSquareMeter)
	a.kilograms_per_square_meterLazy = &kilograms_per_square_meter
	return kilograms_per_square_meter
}

// GramsPerSquareMeter returns the AreaDensity value in GramsPerSquareMeter.
//
// Also known as grammage for paper industry. In fiber industry used with abbreviation 'gsm'.
func (a *AreaDensity) GramsPerSquareMeter() float64 {
	if a.grams_per_square_meterLazy != nil {
		return *a.grams_per_square_meterLazy
	}
	grams_per_square_meter := a.convertFromBase(AreaDensityGramPerSquareMeter)
	a.grams_per_square_meterLazy = &grams_per_square_meter
	return grams_per_square_meter
}

// MilligramsPerSquareMeter returns the AreaDensity value in MilligramsPerSquareMeter.
//
// 
func (a *AreaDensity) MilligramsPerSquareMeter() float64 {
	if a.milligrams_per_square_meterLazy != nil {
		return *a.milligrams_per_square_meterLazy
	}
	milligrams_per_square_meter := a.convertFromBase(AreaDensityMilligramPerSquareMeter)
	a.milligrams_per_square_meterLazy = &milligrams_per_square_meter
	return milligrams_per_square_meter
}


// ToDto creates a AreaDensityDto representation from the AreaDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerSquareMeter by default.
func (a *AreaDensity) ToDto(holdInUnit *AreaDensityUnits) AreaDensityDto {
	if holdInUnit == nil {
		defaultUnit := AreaDensityKilogramPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return AreaDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AreaDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerSquareMeter by default.
func (a *AreaDensity) ToDtoJSON(holdInUnit *AreaDensityUnits) ([]byte, error) {
	// Convert to AreaDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AreaDensity to a specific unit value.
// The function uses the provided unit type (AreaDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AreaDensity) Convert(toUnit AreaDensityUnits) float64 {
	switch toUnit { 
    case AreaDensityKilogramPerSquareMeter:
		return a.KilogramsPerSquareMeter()
    case AreaDensityGramPerSquareMeter:
		return a.GramsPerSquareMeter()
    case AreaDensityMilligramPerSquareMeter:
		return a.MilligramsPerSquareMeter()
	default:
		return math.NaN()
	}
}

func (a *AreaDensity) convertFromBase(toUnit AreaDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaDensityKilogramPerSquareMeter:
		return (value) 
	case AreaDensityGramPerSquareMeter:
		return (value * 1000) 
	case AreaDensityMilligramPerSquareMeter:
		return (value * 1000000) 
	default:
		return math.NaN()
	}
}

func (a *AreaDensity) convertToBase(value float64, fromUnit AreaDensityUnits) float64 {
	switch fromUnit { 
	case AreaDensityKilogramPerSquareMeter:
		return (value) 
	case AreaDensityGramPerSquareMeter:
		return (value / 1000) 
	case AreaDensityMilligramPerSquareMeter:
		return (value / 1000000) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AreaDensity in the default unit (KilogramPerSquareMeter),
// formatted to two decimal places.
func (a AreaDensity) String() string {
	return a.ToString(AreaDensityKilogramPerSquareMeter, 2)
}

// ToString formats the AreaDensity value as a string with the specified unit and fractional digits.
// It converts the AreaDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AreaDensity value will be converted (e.g., KilogramPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AreaDensity with the unit abbreviation.
func (a *AreaDensity) ToString(unit AreaDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAreaDensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAreaDensityAbbreviation(unit))
}

// Equals checks if the given AreaDensity is equal to the current AreaDensity.
//
// Parameters:
//    other: The AreaDensity to compare against.
//
// Returns:
//    bool: Returns true if both AreaDensity are equal, false otherwise.
func (a *AreaDensity) Equals(other *AreaDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AreaDensity with another AreaDensity.
// It returns -1 if the current AreaDensity is less than the other AreaDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AreaDensity to compare against.
//
// Returns:
//    int: -1 if the current AreaDensity is less, 1 if greater, and 0 if equal.
func (a *AreaDensity) CompareTo(other *AreaDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AreaDensity to the current AreaDensity and returns the result.
// The result is a new AreaDensity instance with the sum of the values.
//
// Parameters:
//    other: The AreaDensity to add to the current AreaDensity.
//
// Returns:
//    *AreaDensity: A new AreaDensity instance representing the sum of both AreaDensity.
func (a *AreaDensity) Add(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AreaDensity from the current AreaDensity and returns the result.
// The result is a new AreaDensity instance with the difference of the values.
//
// Parameters:
//    other: The AreaDensity to subtract from the current AreaDensity.
//
// Returns:
//    *AreaDensity: A new AreaDensity instance representing the difference of both AreaDensity.
func (a *AreaDensity) Subtract(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AreaDensity by the given AreaDensity and returns the result.
// The result is a new AreaDensity instance with the product of the values.
//
// Parameters:
//    other: The AreaDensity to multiply with the current AreaDensity.
//
// Returns:
//    *AreaDensity: A new AreaDensity instance representing the product of both AreaDensity.
func (a *AreaDensity) Multiply(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current AreaDensity by the given AreaDensity and returns the result.
// The result is a new AreaDensity instance with the quotient of the values.
//
// Parameters:
//    other: The AreaDensity to divide the current AreaDensity by.
//
// Returns:
//    *AreaDensity: A new AreaDensity instance representing the quotient of both AreaDensity.
func (a *AreaDensity) Divide(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value / other.BaseValue()}
}

// GetAreaDensityAbbreviation gets the unit abbreviation.
func GetAreaDensityAbbreviation(unit AreaDensityUnits) string {
	switch unit { 
	case AreaDensityKilogramPerSquareMeter:
		return "kg/m²" 
	case AreaDensityGramPerSquareMeter:
		return "g/m²" 
	case AreaDensityMilligramPerSquareMeter:
		return "mg/m²" 
	default:
		return ""
	}
}