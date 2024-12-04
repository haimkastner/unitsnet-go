// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AreaMomentOfInertiaUnits defines various units of AreaMomentOfInertia.
type AreaMomentOfInertiaUnits string

const (
    
        // 
        AreaMomentOfInertiaMeterToTheFourth AreaMomentOfInertiaUnits = "MeterToTheFourth"
        // 
        AreaMomentOfInertiaDecimeterToTheFourth AreaMomentOfInertiaUnits = "DecimeterToTheFourth"
        // 
        AreaMomentOfInertiaCentimeterToTheFourth AreaMomentOfInertiaUnits = "CentimeterToTheFourth"
        // 
        AreaMomentOfInertiaMillimeterToTheFourth AreaMomentOfInertiaUnits = "MillimeterToTheFourth"
        // 
        AreaMomentOfInertiaFootToTheFourth AreaMomentOfInertiaUnits = "FootToTheFourth"
        // 
        AreaMomentOfInertiaInchToTheFourth AreaMomentOfInertiaUnits = "InchToTheFourth"
)

// AreaMomentOfInertiaDto represents a AreaMomentOfInertia measurement with a numerical value and its corresponding unit.
type AreaMomentOfInertiaDto struct {
    // Value is the numerical representation of the AreaMomentOfInertia.
	Value float64
    // Unit specifies the unit of measurement for the AreaMomentOfInertia, as defined in the AreaMomentOfInertiaUnits enumeration.
	Unit  AreaMomentOfInertiaUnits
}

// AreaMomentOfInertiaDtoFactory groups methods for creating and serializing AreaMomentOfInertiaDto objects.
type AreaMomentOfInertiaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AreaMomentOfInertiaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AreaMomentOfInertiaDtoFactory) FromJSON(data []byte) (*AreaMomentOfInertiaDto, error) {
	a := AreaMomentOfInertiaDto{}

    // Parse JSON into AreaMomentOfInertiaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a AreaMomentOfInertiaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AreaMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// AreaMomentOfInertia represents a measurement in a of AreaMomentOfInertia.
//
// A geometric property of an area that reflects how its points are distributed with regard to an axis.
type AreaMomentOfInertia struct {
	// value is the base measurement stored internally.
	value       float64
    
    meters_to_the_fourthLazy *float64 
    decimeters_to_the_fourthLazy *float64 
    centimeters_to_the_fourthLazy *float64 
    millimeters_to_the_fourthLazy *float64 
    feet_to_the_fourthLazy *float64 
    inches_to_the_fourthLazy *float64 
}

// AreaMomentOfInertiaFactory groups methods for creating AreaMomentOfInertia instances.
type AreaMomentOfInertiaFactory struct{}

// CreateAreaMomentOfInertia creates a new AreaMomentOfInertia instance from the given value and unit.
func (uf AreaMomentOfInertiaFactory) CreateAreaMomentOfInertia(value float64, unit AreaMomentOfInertiaUnits) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, unit)
}

// FromDto converts a AreaMomentOfInertiaDto to a AreaMomentOfInertia instance.
func (uf AreaMomentOfInertiaFactory) FromDto(dto AreaMomentOfInertiaDto) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AreaMomentOfInertia instance.
func (uf AreaMomentOfInertiaFactory) FromDtoJSON(data []byte) (*AreaMomentOfInertia, error) {
	unitDto, err := AreaMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AreaMomentOfInertiaDto from JSON: %w", err)
	}
	return AreaMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromMetersToTheFourth creates a new AreaMomentOfInertia instance from a value in MetersToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromMetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaMeterToTheFourth)
}

// FromDecimetersToTheFourth creates a new AreaMomentOfInertia instance from a value in DecimetersToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromDecimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaDecimeterToTheFourth)
}

// FromCentimetersToTheFourth creates a new AreaMomentOfInertia instance from a value in CentimetersToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromCentimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaCentimeterToTheFourth)
}

// FromMillimetersToTheFourth creates a new AreaMomentOfInertia instance from a value in MillimetersToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromMillimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaMillimeterToTheFourth)
}

// FromFeetToTheFourth creates a new AreaMomentOfInertia instance from a value in FeetToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromFeetToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaFootToTheFourth)
}

// FromInchesToTheFourth creates a new AreaMomentOfInertia instance from a value in InchesToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromInchesToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaInchToTheFourth)
}


// newAreaMomentOfInertia creates a new AreaMomentOfInertia.
func newAreaMomentOfInertia(value float64, fromUnit AreaMomentOfInertiaUnits) (*AreaMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AreaMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AreaMomentOfInertia in MeterToTheFourth unit (the base/default quantity).
func (a *AreaMomentOfInertia) BaseValue() float64 {
	return a.value
}


// MetersToTheFourth returns the AreaMomentOfInertia value in MetersToTheFourth.
//
// 
func (a *AreaMomentOfInertia) MetersToTheFourth() float64 {
	if a.meters_to_the_fourthLazy != nil {
		return *a.meters_to_the_fourthLazy
	}
	meters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaMeterToTheFourth)
	a.meters_to_the_fourthLazy = &meters_to_the_fourth
	return meters_to_the_fourth
}

// DecimetersToTheFourth returns the AreaMomentOfInertia value in DecimetersToTheFourth.
//
// 
func (a *AreaMomentOfInertia) DecimetersToTheFourth() float64 {
	if a.decimeters_to_the_fourthLazy != nil {
		return *a.decimeters_to_the_fourthLazy
	}
	decimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaDecimeterToTheFourth)
	a.decimeters_to_the_fourthLazy = &decimeters_to_the_fourth
	return decimeters_to_the_fourth
}

// CentimetersToTheFourth returns the AreaMomentOfInertia value in CentimetersToTheFourth.
//
// 
func (a *AreaMomentOfInertia) CentimetersToTheFourth() float64 {
	if a.centimeters_to_the_fourthLazy != nil {
		return *a.centimeters_to_the_fourthLazy
	}
	centimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaCentimeterToTheFourth)
	a.centimeters_to_the_fourthLazy = &centimeters_to_the_fourth
	return centimeters_to_the_fourth
}

// MillimetersToTheFourth returns the AreaMomentOfInertia value in MillimetersToTheFourth.
//
// 
func (a *AreaMomentOfInertia) MillimetersToTheFourth() float64 {
	if a.millimeters_to_the_fourthLazy != nil {
		return *a.millimeters_to_the_fourthLazy
	}
	millimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaMillimeterToTheFourth)
	a.millimeters_to_the_fourthLazy = &millimeters_to_the_fourth
	return millimeters_to_the_fourth
}

// FeetToTheFourth returns the AreaMomentOfInertia value in FeetToTheFourth.
//
// 
func (a *AreaMomentOfInertia) FeetToTheFourth() float64 {
	if a.feet_to_the_fourthLazy != nil {
		return *a.feet_to_the_fourthLazy
	}
	feet_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaFootToTheFourth)
	a.feet_to_the_fourthLazy = &feet_to_the_fourth
	return feet_to_the_fourth
}

// InchesToTheFourth returns the AreaMomentOfInertia value in InchesToTheFourth.
//
// 
func (a *AreaMomentOfInertia) InchesToTheFourth() float64 {
	if a.inches_to_the_fourthLazy != nil {
		return *a.inches_to_the_fourthLazy
	}
	inches_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaInchToTheFourth)
	a.inches_to_the_fourthLazy = &inches_to_the_fourth
	return inches_to_the_fourth
}


// ToDto creates a AreaMomentOfInertiaDto representation from the AreaMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterToTheFourth by default.
func (a *AreaMomentOfInertia) ToDto(holdInUnit *AreaMomentOfInertiaUnits) AreaMomentOfInertiaDto {
	if holdInUnit == nil {
		defaultUnit := AreaMomentOfInertiaMeterToTheFourth // Default value
		holdInUnit = &defaultUnit
	}

	return AreaMomentOfInertiaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AreaMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterToTheFourth by default.
func (a *AreaMomentOfInertia) ToDtoJSON(holdInUnit *AreaMomentOfInertiaUnits) ([]byte, error) {
	// Convert to AreaMomentOfInertiaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AreaMomentOfInertia to a specific unit value.
// The function uses the provided unit type (AreaMomentOfInertiaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AreaMomentOfInertia) Convert(toUnit AreaMomentOfInertiaUnits) float64 {
	switch toUnit { 
    case AreaMomentOfInertiaMeterToTheFourth:
		return a.MetersToTheFourth()
    case AreaMomentOfInertiaDecimeterToTheFourth:
		return a.DecimetersToTheFourth()
    case AreaMomentOfInertiaCentimeterToTheFourth:
		return a.CentimetersToTheFourth()
    case AreaMomentOfInertiaMillimeterToTheFourth:
		return a.MillimetersToTheFourth()
    case AreaMomentOfInertiaFootToTheFourth:
		return a.FeetToTheFourth()
    case AreaMomentOfInertiaInchToTheFourth:
		return a.InchesToTheFourth()
	default:
		return math.NaN()
	}
}

func (a *AreaMomentOfInertia) convertFromBase(toUnit AreaMomentOfInertiaUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return (value) 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return (value * 1e4) 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return (value * 1e8) 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return (value * 1e12) 
	case AreaMomentOfInertiaFootToTheFourth:
		return (value / math.Pow(0.3048, 4)) 
	case AreaMomentOfInertiaInchToTheFourth:
		return (value / math.Pow(2.54e-2, 4)) 
	default:
		return math.NaN()
	}
}

func (a *AreaMomentOfInertia) convertToBase(value float64, fromUnit AreaMomentOfInertiaUnits) float64 {
	switch fromUnit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return (value) 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return (value / 1e4) 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return (value / 1e8) 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return (value / 1e12) 
	case AreaMomentOfInertiaFootToTheFourth:
		return (value * math.Pow(0.3048, 4)) 
	case AreaMomentOfInertiaInchToTheFourth:
		return (value * math.Pow(2.54e-2, 4)) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AreaMomentOfInertia in the default unit (MeterToTheFourth),
// formatted to two decimal places.
func (a AreaMomentOfInertia) String() string {
	return a.ToString(AreaMomentOfInertiaMeterToTheFourth, 2)
}

// ToString formats the AreaMomentOfInertia value as a string with the specified unit and fractional digits.
// It converts the AreaMomentOfInertia to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AreaMomentOfInertia value will be converted (e.g., MeterToTheFourth))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AreaMomentOfInertia with the unit abbreviation.
func (a *AreaMomentOfInertia) ToString(unit AreaMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AreaMomentOfInertia) getUnitAbbreviation(unit AreaMomentOfInertiaUnits) string {
	switch unit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return "m⁴" 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return "dm⁴" 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return "cm⁴" 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return "mm⁴" 
	case AreaMomentOfInertiaFootToTheFourth:
		return "ft⁴" 
	case AreaMomentOfInertiaInchToTheFourth:
		return "in⁴" 
	default:
		return ""
	}
}

// Equals checks if the given AreaMomentOfInertia is equal to the current AreaMomentOfInertia.
//
// Parameters:
//    other: The AreaMomentOfInertia to compare against.
//
// Returns:
//    bool: Returns true if both AreaMomentOfInertia are equal, false otherwise.
func (a *AreaMomentOfInertia) Equals(other *AreaMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AreaMomentOfInertia with another AreaMomentOfInertia.
// It returns -1 if the current AreaMomentOfInertia is less than the other AreaMomentOfInertia, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AreaMomentOfInertia to compare against.
//
// Returns:
//    int: -1 if the current AreaMomentOfInertia is less, 1 if greater, and 0 if equal.
func (a *AreaMomentOfInertia) CompareTo(other *AreaMomentOfInertia) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AreaMomentOfInertia to the current AreaMomentOfInertia and returns the result.
// The result is a new AreaMomentOfInertia instance with the sum of the values.
//
// Parameters:
//    other: The AreaMomentOfInertia to add to the current AreaMomentOfInertia.
//
// Returns:
//    *AreaMomentOfInertia: A new AreaMomentOfInertia instance representing the sum of both AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Add(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AreaMomentOfInertia from the current AreaMomentOfInertia and returns the result.
// The result is a new AreaMomentOfInertia instance with the difference of the values.
//
// Parameters:
//    other: The AreaMomentOfInertia to subtract from the current AreaMomentOfInertia.
//
// Returns:
//    *AreaMomentOfInertia: A new AreaMomentOfInertia instance representing the difference of both AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Subtract(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AreaMomentOfInertia by the given AreaMomentOfInertia and returns the result.
// The result is a new AreaMomentOfInertia instance with the product of the values.
//
// Parameters:
//    other: The AreaMomentOfInertia to multiply with the current AreaMomentOfInertia.
//
// Returns:
//    *AreaMomentOfInertia: A new AreaMomentOfInertia instance representing the product of both AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Multiply(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide divides the current AreaMomentOfInertia by the given AreaMomentOfInertia and returns the result.
// The result is a new AreaMomentOfInertia instance with the quotient of the values.
//
// Parameters:
//    other: The AreaMomentOfInertia to divide the current AreaMomentOfInertia by.
//
// Returns:
//    *AreaMomentOfInertia: A new AreaMomentOfInertia instance representing the quotient of both AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Divide(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value / other.BaseValue()}
}