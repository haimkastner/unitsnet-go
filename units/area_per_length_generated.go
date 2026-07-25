// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AreaPerLengthUnits defines various units of AreaPerLength.
type AreaPerLengthUnits string

const (
    
        // 
        AreaPerLengthSquareMeterPerMeter AreaPerLengthUnits = "SquareMeterPerMeter"
        // 
        AreaPerLengthSquareCentimeterPerMeter AreaPerLengthUnits = "SquareCentimeterPerMeter"
        // 
        AreaPerLengthSquareMillimeterPerMeter AreaPerLengthUnits = "SquareMillimeterPerMeter"
        // 
        AreaPerLengthSquareInchPerFoot AreaPerLengthUnits = "SquareInchPerFoot"
        // 
        AreaPerLengthSquareInchPerInch AreaPerLengthUnits = "SquareInchPerInch"
        // 
        AreaPerLengthSquareFootPerFoot AreaPerLengthUnits = "SquareFootPerFoot"
)

var internalAreaPerLengthUnitsMap = map[AreaPerLengthUnits]bool{
	
	AreaPerLengthSquareMeterPerMeter: true,
	AreaPerLengthSquareCentimeterPerMeter: true,
	AreaPerLengthSquareMillimeterPerMeter: true,
	AreaPerLengthSquareInchPerFoot: true,
	AreaPerLengthSquareInchPerInch: true,
	AreaPerLengthSquareFootPerFoot: true,
}

// AreaPerLengthDto represents a AreaPerLength measurement with a numerical value and its corresponding unit.
type AreaPerLengthDto struct {
    // Value is the numerical representation of the AreaPerLength.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the AreaPerLength, as defined in the AreaPerLengthUnits enumeration.
	Unit  AreaPerLengthUnits `json:"unit" validate:"required,oneof=SquareMeterPerMeter SquareCentimeterPerMeter SquareMillimeterPerMeter SquareInchPerFoot SquareInchPerInch SquareFootPerFoot"`
}

// AreaPerLengthDtoFactory groups methods for creating and serializing AreaPerLengthDto objects.
type AreaPerLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AreaPerLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AreaPerLengthDtoFactory) FromJSON(data []byte) (*AreaPerLengthDto, error) {
	a := AreaPerLengthDto{}

    // Parse JSON into AreaPerLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AreaPerLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AreaPerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// AreaPerLength represents a measurement in a of AreaPerLength.
//
// The magnitude of area per unit length, typically used in structural engineering to specify distributed reinforcement.
type AreaPerLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    square_meters_per_meterLazy *float64 
    square_centimeters_per_meterLazy *float64 
    square_millimeters_per_meterLazy *float64 
    square_inches_per_footLazy *float64 
    square_inches_per_inchLazy *float64 
    square_feet_per_footLazy *float64 
}

// AreaPerLengthFactory groups methods for creating AreaPerLength instances.
type AreaPerLengthFactory struct{}

// CreateAreaPerLength creates a new AreaPerLength instance from the given value and unit.
func (uf AreaPerLengthFactory) CreateAreaPerLength(value float64, unit AreaPerLengthUnits) (*AreaPerLength, error) {
	return newAreaPerLength(value, unit)
}

// FromDto converts a AreaPerLengthDto to a AreaPerLength instance.
func (uf AreaPerLengthFactory) FromDto(dto AreaPerLengthDto) (*AreaPerLength, error) {
	return newAreaPerLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AreaPerLength instance.
func (uf AreaPerLengthFactory) FromDtoJSON(data []byte) (*AreaPerLength, error) {
	unitDto, err := AreaPerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AreaPerLengthDto from JSON: %w", err)
	}
	return AreaPerLengthFactory{}.FromDto(*unitDto)
}


// FromSquareMetersPerMeter creates a new AreaPerLength instance from a value in SquareMetersPerMeter.
func (uf AreaPerLengthFactory) FromSquareMetersPerMeter(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareMeterPerMeter)
}

// FromSquareCentimetersPerMeter creates a new AreaPerLength instance from a value in SquareCentimetersPerMeter.
func (uf AreaPerLengthFactory) FromSquareCentimetersPerMeter(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareCentimeterPerMeter)
}

// FromSquareMillimetersPerMeter creates a new AreaPerLength instance from a value in SquareMillimetersPerMeter.
func (uf AreaPerLengthFactory) FromSquareMillimetersPerMeter(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareMillimeterPerMeter)
}

// FromSquareInchesPerFoot creates a new AreaPerLength instance from a value in SquareInchesPerFoot.
func (uf AreaPerLengthFactory) FromSquareInchesPerFoot(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareInchPerFoot)
}

// FromSquareInchesPerInch creates a new AreaPerLength instance from a value in SquareInchesPerInch.
func (uf AreaPerLengthFactory) FromSquareInchesPerInch(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareInchPerInch)
}

// FromSquareFeetPerFoot creates a new AreaPerLength instance from a value in SquareFeetPerFoot.
func (uf AreaPerLengthFactory) FromSquareFeetPerFoot(value float64) (*AreaPerLength, error) {
	return newAreaPerLength(value, AreaPerLengthSquareFootPerFoot)
}


// newAreaPerLength creates a new AreaPerLength.
func newAreaPerLength(value float64, fromUnit AreaPerLengthUnits) (*AreaPerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalAreaPerLengthUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in AreaPerLengthUnits", fromUnit)
	}
	a := &AreaPerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AreaPerLength in SquareMeterPerMeter unit (the base/default quantity).
func (a *AreaPerLength) BaseValue() float64 {
	return a.value
}


// SquareMetersPerMeter returns the AreaPerLength value in SquareMetersPerMeter.
//
// 
func (a *AreaPerLength) SquareMetersPerMeter() float64 {
	if a.square_meters_per_meterLazy != nil {
		return *a.square_meters_per_meterLazy
	}
	square_meters_per_meter := a.convertFromBase(AreaPerLengthSquareMeterPerMeter)
	a.square_meters_per_meterLazy = &square_meters_per_meter
	return square_meters_per_meter
}

// SquareCentimetersPerMeter returns the AreaPerLength value in SquareCentimetersPerMeter.
//
// 
func (a *AreaPerLength) SquareCentimetersPerMeter() float64 {
	if a.square_centimeters_per_meterLazy != nil {
		return *a.square_centimeters_per_meterLazy
	}
	square_centimeters_per_meter := a.convertFromBase(AreaPerLengthSquareCentimeterPerMeter)
	a.square_centimeters_per_meterLazy = &square_centimeters_per_meter
	return square_centimeters_per_meter
}

// SquareMillimetersPerMeter returns the AreaPerLength value in SquareMillimetersPerMeter.
//
// 
func (a *AreaPerLength) SquareMillimetersPerMeter() float64 {
	if a.square_millimeters_per_meterLazy != nil {
		return *a.square_millimeters_per_meterLazy
	}
	square_millimeters_per_meter := a.convertFromBase(AreaPerLengthSquareMillimeterPerMeter)
	a.square_millimeters_per_meterLazy = &square_millimeters_per_meter
	return square_millimeters_per_meter
}

// SquareInchesPerFoot returns the AreaPerLength value in SquareInchesPerFoot.
//
// 
func (a *AreaPerLength) SquareInchesPerFoot() float64 {
	if a.square_inches_per_footLazy != nil {
		return *a.square_inches_per_footLazy
	}
	square_inches_per_foot := a.convertFromBase(AreaPerLengthSquareInchPerFoot)
	a.square_inches_per_footLazy = &square_inches_per_foot
	return square_inches_per_foot
}

// SquareInchesPerInch returns the AreaPerLength value in SquareInchesPerInch.
//
// 
func (a *AreaPerLength) SquareInchesPerInch() float64 {
	if a.square_inches_per_inchLazy != nil {
		return *a.square_inches_per_inchLazy
	}
	square_inches_per_inch := a.convertFromBase(AreaPerLengthSquareInchPerInch)
	a.square_inches_per_inchLazy = &square_inches_per_inch
	return square_inches_per_inch
}

// SquareFeetPerFoot returns the AreaPerLength value in SquareFeetPerFoot.
//
// 
func (a *AreaPerLength) SquareFeetPerFoot() float64 {
	if a.square_feet_per_footLazy != nil {
		return *a.square_feet_per_footLazy
	}
	square_feet_per_foot := a.convertFromBase(AreaPerLengthSquareFootPerFoot)
	a.square_feet_per_footLazy = &square_feet_per_foot
	return square_feet_per_foot
}


// ToDto creates a AreaPerLengthDto representation from the AreaPerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterPerMeter by default.
func (a *AreaPerLength) ToDto(holdInUnit *AreaPerLengthUnits) AreaPerLengthDto {
	if holdInUnit == nil {
		defaultUnit := AreaPerLengthSquareMeterPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return AreaPerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AreaPerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterPerMeter by default.
func (a *AreaPerLength) ToDtoJSON(holdInUnit *AreaPerLengthUnits) ([]byte, error) {
	// Convert to AreaPerLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AreaPerLength to a specific unit value.
// The function uses the provided unit type (AreaPerLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AreaPerLength) Convert(toUnit AreaPerLengthUnits) float64 {
	switch toUnit { 
    case AreaPerLengthSquareMeterPerMeter:
		return a.SquareMetersPerMeter()
    case AreaPerLengthSquareCentimeterPerMeter:
		return a.SquareCentimetersPerMeter()
    case AreaPerLengthSquareMillimeterPerMeter:
		return a.SquareMillimetersPerMeter()
    case AreaPerLengthSquareInchPerFoot:
		return a.SquareInchesPerFoot()
    case AreaPerLengthSquareInchPerInch:
		return a.SquareInchesPerInch()
    case AreaPerLengthSquareFootPerFoot:
		return a.SquareFeetPerFoot()
	default:
		return math.NaN()
	}
}

func (a *AreaPerLength) convertFromBase(toUnit AreaPerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaPerLengthSquareMeterPerMeter:
		return (value) 
	case AreaPerLengthSquareCentimeterPerMeter:
		return (value / 1e-4) 
	case AreaPerLengthSquareMillimeterPerMeter:
		return (value / 1e-6) 
	case AreaPerLengthSquareInchPerFoot:
		return (value * 0.3048 / 0.00064516) 
	case AreaPerLengthSquareInchPerInch:
		return (value * 0.0254 / 0.00064516) 
	case AreaPerLengthSquareFootPerFoot:
		return (value * 0.3048 / 0.09290304) 
	default:
		return math.NaN()
	}
}

func (a *AreaPerLength) convertToBase(value float64, fromUnit AreaPerLengthUnits) float64 {
	switch fromUnit { 
	case AreaPerLengthSquareMeterPerMeter:
		return (value) 
	case AreaPerLengthSquareCentimeterPerMeter:
		return (value * 1e-4) 
	case AreaPerLengthSquareMillimeterPerMeter:
		return (value * 1e-6) 
	case AreaPerLengthSquareInchPerFoot:
		return (value * 0.00064516 / 0.3048) 
	case AreaPerLengthSquareInchPerInch:
		return (value * 0.00064516 / 0.0254) 
	case AreaPerLengthSquareFootPerFoot:
		return (value * 0.09290304 / 0.3048) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AreaPerLength in the default unit (SquareMeterPerMeter),
// formatted to two decimal places.
func (a AreaPerLength) String() string {
	return a.ToString(AreaPerLengthSquareMeterPerMeter, 2)
}

// ToString formats the AreaPerLength value as a string with the specified unit and fractional digits.
// It converts the AreaPerLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AreaPerLength value will be converted (e.g., SquareMeterPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AreaPerLength with the unit abbreviation.
func (a *AreaPerLength) ToString(unit AreaPerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAreaPerLengthAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAreaPerLengthAbbreviation(unit))
}

// Equals checks if the given AreaPerLength is equal to the current AreaPerLength.
//
// Parameters:
//    other: The AreaPerLength to compare against.
//
// Returns:
//    bool: Returns true if both AreaPerLength are equal, false otherwise.
func (a *AreaPerLength) Equals(other *AreaPerLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AreaPerLength with another AreaPerLength.
// It returns -1 if the current AreaPerLength is less than the other AreaPerLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AreaPerLength to compare against.
//
// Returns:
//    int: -1 if the current AreaPerLength is less, 1 if greater, and 0 if equal.
func (a *AreaPerLength) CompareTo(other *AreaPerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AreaPerLength to the current AreaPerLength and returns the result.
// The result is a new AreaPerLength instance with the sum of the values.
//
// Parameters:
//    other: The AreaPerLength to add to the current AreaPerLength.
//
// Returns:
//    *AreaPerLength: A new AreaPerLength instance representing the sum of both AreaPerLength.
func (a *AreaPerLength) Add(other *AreaPerLength) *AreaPerLength {
	return &AreaPerLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AreaPerLength from the current AreaPerLength and returns the result.
// The result is a new AreaPerLength instance with the difference of the values.
//
// Parameters:
//    other: The AreaPerLength to subtract from the current AreaPerLength.
//
// Returns:
//    *AreaPerLength: A new AreaPerLength instance representing the difference of both AreaPerLength.
func (a *AreaPerLength) Subtract(other *AreaPerLength) *AreaPerLength {
	return &AreaPerLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AreaPerLength by the given AreaPerLength and returns the result.
// The result is a new AreaPerLength instance with the product of the values.
//
// Parameters:
//    other: The AreaPerLength to multiply with the current AreaPerLength.
//
// Returns:
//    *AreaPerLength: A new AreaPerLength instance representing the product of both AreaPerLength.
func (a *AreaPerLength) Multiply(other *AreaPerLength) *AreaPerLength {
	return &AreaPerLength{value: a.value * other.BaseValue()}
}

// Divide divides the current AreaPerLength by the given AreaPerLength and returns the result.
// The result is a new AreaPerLength instance with the quotient of the values.
//
// Parameters:
//    other: The AreaPerLength to divide the current AreaPerLength by.
//
// Returns:
//    *AreaPerLength: A new AreaPerLength instance representing the quotient of both AreaPerLength.
func (a *AreaPerLength) Divide(other *AreaPerLength) *AreaPerLength {
	return &AreaPerLength{value: a.value / other.BaseValue()}
}

// GetAreaPerLengthAbbreviation gets the unit abbreviation.
func GetAreaPerLengthAbbreviation(unit AreaPerLengthUnits) string {
	switch unit { 
	case AreaPerLengthSquareMeterPerMeter:
		return "m²/m" 
	case AreaPerLengthSquareCentimeterPerMeter:
		return "cm²/m" 
	case AreaPerLengthSquareMillimeterPerMeter:
		return "mm²/m" 
	case AreaPerLengthSquareInchPerFoot:
		return "in²/ft" 
	case AreaPerLengthSquareInchPerInch:
		return "in²/in" 
	case AreaPerLengthSquareFootPerFoot:
		return "ft²/ft" 
	default:
		return ""
	}
}