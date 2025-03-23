// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// WarpingMomentOfInertiaUnits defines various units of WarpingMomentOfInertia.
type WarpingMomentOfInertiaUnits string

const (
    
        // 
        WarpingMomentOfInertiaMeterToTheSixth WarpingMomentOfInertiaUnits = "MeterToTheSixth"
        // 
        WarpingMomentOfInertiaDecimeterToTheSixth WarpingMomentOfInertiaUnits = "DecimeterToTheSixth"
        // 
        WarpingMomentOfInertiaCentimeterToTheSixth WarpingMomentOfInertiaUnits = "CentimeterToTheSixth"
        // 
        WarpingMomentOfInertiaMillimeterToTheSixth WarpingMomentOfInertiaUnits = "MillimeterToTheSixth"
        // 
        WarpingMomentOfInertiaFootToTheSixth WarpingMomentOfInertiaUnits = "FootToTheSixth"
        // 
        WarpingMomentOfInertiaInchToTheSixth WarpingMomentOfInertiaUnits = "InchToTheSixth"
)

var internalWarpingMomentOfInertiaUnitsMap = map[WarpingMomentOfInertiaUnits]bool{
	
	WarpingMomentOfInertiaMeterToTheSixth: true,
	WarpingMomentOfInertiaDecimeterToTheSixth: true,
	WarpingMomentOfInertiaCentimeterToTheSixth: true,
	WarpingMomentOfInertiaMillimeterToTheSixth: true,
	WarpingMomentOfInertiaFootToTheSixth: true,
	WarpingMomentOfInertiaInchToTheSixth: true,
}

// WarpingMomentOfInertiaDto represents a WarpingMomentOfInertia measurement with a numerical value and its corresponding unit.
type WarpingMomentOfInertiaDto struct {
    // Value is the numerical representation of the WarpingMomentOfInertia.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the WarpingMomentOfInertia, as defined in the WarpingMomentOfInertiaUnits enumeration.
	Unit  WarpingMomentOfInertiaUnits `json:"unit" validate:"required,oneof=MeterToTheSixth DecimeterToTheSixth CentimeterToTheSixth MillimeterToTheSixth FootToTheSixth InchToTheSixth"`
}

// WarpingMomentOfInertiaDtoFactory groups methods for creating and serializing WarpingMomentOfInertiaDto objects.
type WarpingMomentOfInertiaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a WarpingMomentOfInertiaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf WarpingMomentOfInertiaDtoFactory) FromJSON(data []byte) (*WarpingMomentOfInertiaDto, error) {
	a := WarpingMomentOfInertiaDto{}

    // Parse JSON into WarpingMomentOfInertiaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a WarpingMomentOfInertiaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a WarpingMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// WarpingMomentOfInertia represents a measurement in a of WarpingMomentOfInertia.
//
// A geometric property of an area that is used to determine the warping stress.
type WarpingMomentOfInertia struct {
	// value is the base measurement stored internally.
	value       float64
    
    meters_to_the_sixthLazy *float64 
    decimeters_to_the_sixthLazy *float64 
    centimeters_to_the_sixthLazy *float64 
    millimeters_to_the_sixthLazy *float64 
    feet_to_the_sixthLazy *float64 
    inches_to_the_sixthLazy *float64 
}

// WarpingMomentOfInertiaFactory groups methods for creating WarpingMomentOfInertia instances.
type WarpingMomentOfInertiaFactory struct{}

// CreateWarpingMomentOfInertia creates a new WarpingMomentOfInertia instance from the given value and unit.
func (uf WarpingMomentOfInertiaFactory) CreateWarpingMomentOfInertia(value float64, unit WarpingMomentOfInertiaUnits) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, unit)
}

// FromDto converts a WarpingMomentOfInertiaDto to a WarpingMomentOfInertia instance.
func (uf WarpingMomentOfInertiaFactory) FromDto(dto WarpingMomentOfInertiaDto) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a WarpingMomentOfInertia instance.
func (uf WarpingMomentOfInertiaFactory) FromDtoJSON(data []byte) (*WarpingMomentOfInertia, error) {
	unitDto, err := WarpingMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse WarpingMomentOfInertiaDto from JSON: %w", err)
	}
	return WarpingMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromMetersToTheSixth creates a new WarpingMomentOfInertia instance from a value in MetersToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromMetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaMeterToTheSixth)
}

// FromDecimetersToTheSixth creates a new WarpingMomentOfInertia instance from a value in DecimetersToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromDecimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaDecimeterToTheSixth)
}

// FromCentimetersToTheSixth creates a new WarpingMomentOfInertia instance from a value in CentimetersToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromCentimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaCentimeterToTheSixth)
}

// FromMillimetersToTheSixth creates a new WarpingMomentOfInertia instance from a value in MillimetersToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromMillimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaMillimeterToTheSixth)
}

// FromFeetToTheSixth creates a new WarpingMomentOfInertia instance from a value in FeetToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromFeetToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaFootToTheSixth)
}

// FromInchesToTheSixth creates a new WarpingMomentOfInertia instance from a value in InchesToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromInchesToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaInchToTheSixth)
}


// newWarpingMomentOfInertia creates a new WarpingMomentOfInertia.
func newWarpingMomentOfInertia(value float64, fromUnit WarpingMomentOfInertiaUnits) (*WarpingMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalWarpingMomentOfInertiaUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in WarpingMomentOfInertiaUnits", fromUnit)
	}
	a := &WarpingMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of WarpingMomentOfInertia in MeterToTheSixth unit (the base/default quantity).
func (a *WarpingMomentOfInertia) BaseValue() float64 {
	return a.value
}


// MetersToTheSixth returns the WarpingMomentOfInertia value in MetersToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) MetersToTheSixth() float64 {
	if a.meters_to_the_sixthLazy != nil {
		return *a.meters_to_the_sixthLazy
	}
	meters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaMeterToTheSixth)
	a.meters_to_the_sixthLazy = &meters_to_the_sixth
	return meters_to_the_sixth
}

// DecimetersToTheSixth returns the WarpingMomentOfInertia value in DecimetersToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) DecimetersToTheSixth() float64 {
	if a.decimeters_to_the_sixthLazy != nil {
		return *a.decimeters_to_the_sixthLazy
	}
	decimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaDecimeterToTheSixth)
	a.decimeters_to_the_sixthLazy = &decimeters_to_the_sixth
	return decimeters_to_the_sixth
}

// CentimetersToTheSixth returns the WarpingMomentOfInertia value in CentimetersToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) CentimetersToTheSixth() float64 {
	if a.centimeters_to_the_sixthLazy != nil {
		return *a.centimeters_to_the_sixthLazy
	}
	centimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaCentimeterToTheSixth)
	a.centimeters_to_the_sixthLazy = &centimeters_to_the_sixth
	return centimeters_to_the_sixth
}

// MillimetersToTheSixth returns the WarpingMomentOfInertia value in MillimetersToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) MillimetersToTheSixth() float64 {
	if a.millimeters_to_the_sixthLazy != nil {
		return *a.millimeters_to_the_sixthLazy
	}
	millimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaMillimeterToTheSixth)
	a.millimeters_to_the_sixthLazy = &millimeters_to_the_sixth
	return millimeters_to_the_sixth
}

// FeetToTheSixth returns the WarpingMomentOfInertia value in FeetToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) FeetToTheSixth() float64 {
	if a.feet_to_the_sixthLazy != nil {
		return *a.feet_to_the_sixthLazy
	}
	feet_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaFootToTheSixth)
	a.feet_to_the_sixthLazy = &feet_to_the_sixth
	return feet_to_the_sixth
}

// InchesToTheSixth returns the WarpingMomentOfInertia value in InchesToTheSixth.
//
// 
func (a *WarpingMomentOfInertia) InchesToTheSixth() float64 {
	if a.inches_to_the_sixthLazy != nil {
		return *a.inches_to_the_sixthLazy
	}
	inches_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaInchToTheSixth)
	a.inches_to_the_sixthLazy = &inches_to_the_sixth
	return inches_to_the_sixth
}


// ToDto creates a WarpingMomentOfInertiaDto representation from the WarpingMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterToTheSixth by default.
func (a *WarpingMomentOfInertia) ToDto(holdInUnit *WarpingMomentOfInertiaUnits) WarpingMomentOfInertiaDto {
	if holdInUnit == nil {
		defaultUnit := WarpingMomentOfInertiaMeterToTheSixth // Default value
		holdInUnit = &defaultUnit
	}

	return WarpingMomentOfInertiaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the WarpingMomentOfInertia instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterToTheSixth by default.
func (a *WarpingMomentOfInertia) ToDtoJSON(holdInUnit *WarpingMomentOfInertiaUnits) ([]byte, error) {
	// Convert to WarpingMomentOfInertiaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a WarpingMomentOfInertia to a specific unit value.
// The function uses the provided unit type (WarpingMomentOfInertiaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *WarpingMomentOfInertia) Convert(toUnit WarpingMomentOfInertiaUnits) float64 {
	switch toUnit { 
    case WarpingMomentOfInertiaMeterToTheSixth:
		return a.MetersToTheSixth()
    case WarpingMomentOfInertiaDecimeterToTheSixth:
		return a.DecimetersToTheSixth()
    case WarpingMomentOfInertiaCentimeterToTheSixth:
		return a.CentimetersToTheSixth()
    case WarpingMomentOfInertiaMillimeterToTheSixth:
		return a.MillimetersToTheSixth()
    case WarpingMomentOfInertiaFootToTheSixth:
		return a.FeetToTheSixth()
    case WarpingMomentOfInertiaInchToTheSixth:
		return a.InchesToTheSixth()
	default:
		return math.NaN()
	}
}

func (a *WarpingMomentOfInertia) convertFromBase(toUnit WarpingMomentOfInertiaUnits) float64 {
    value := a.value
	switch toUnit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return (value) 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return (value * 1e6) 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return (value * 1e12) 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return (value * 1e18) 
	case WarpingMomentOfInertiaFootToTheSixth:
		return (value / math.Pow(0.3048, 6)) 
	case WarpingMomentOfInertiaInchToTheSixth:
		return (value / math.Pow(2.54e-2, 6)) 
	default:
		return math.NaN()
	}
}

func (a *WarpingMomentOfInertia) convertToBase(value float64, fromUnit WarpingMomentOfInertiaUnits) float64 {
	switch fromUnit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return (value) 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return (value / 1e6) 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return (value / 1e12) 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return (value / 1e18) 
	case WarpingMomentOfInertiaFootToTheSixth:
		return (value * math.Pow(0.3048, 6)) 
	case WarpingMomentOfInertiaInchToTheSixth:
		return (value * math.Pow(2.54e-2, 6)) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the WarpingMomentOfInertia in the default unit (MeterToTheSixth),
// formatted to two decimal places.
func (a WarpingMomentOfInertia) String() string {
	return a.ToString(WarpingMomentOfInertiaMeterToTheSixth, 2)
}

// ToString formats the WarpingMomentOfInertia value as a string with the specified unit and fractional digits.
// It converts the WarpingMomentOfInertia to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the WarpingMomentOfInertia value will be converted (e.g., MeterToTheSixth))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the WarpingMomentOfInertia with the unit abbreviation.
func (a *WarpingMomentOfInertia) ToString(unit WarpingMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetWarpingMomentOfInertiaAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetWarpingMomentOfInertiaAbbreviation(unit))
}

// Equals checks if the given WarpingMomentOfInertia is equal to the current WarpingMomentOfInertia.
//
// Parameters:
//    other: The WarpingMomentOfInertia to compare against.
//
// Returns:
//    bool: Returns true if both WarpingMomentOfInertia are equal, false otherwise.
func (a *WarpingMomentOfInertia) Equals(other *WarpingMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current WarpingMomentOfInertia with another WarpingMomentOfInertia.
// It returns -1 if the current WarpingMomentOfInertia is less than the other WarpingMomentOfInertia, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The WarpingMomentOfInertia to compare against.
//
// Returns:
//    int: -1 if the current WarpingMomentOfInertia is less, 1 if greater, and 0 if equal.
func (a *WarpingMomentOfInertia) CompareTo(other *WarpingMomentOfInertia) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given WarpingMomentOfInertia to the current WarpingMomentOfInertia and returns the result.
// The result is a new WarpingMomentOfInertia instance with the sum of the values.
//
// Parameters:
//    other: The WarpingMomentOfInertia to add to the current WarpingMomentOfInertia.
//
// Returns:
//    *WarpingMomentOfInertia: A new WarpingMomentOfInertia instance representing the sum of both WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Add(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given WarpingMomentOfInertia from the current WarpingMomentOfInertia and returns the result.
// The result is a new WarpingMomentOfInertia instance with the difference of the values.
//
// Parameters:
//    other: The WarpingMomentOfInertia to subtract from the current WarpingMomentOfInertia.
//
// Returns:
//    *WarpingMomentOfInertia: A new WarpingMomentOfInertia instance representing the difference of both WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Subtract(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current WarpingMomentOfInertia by the given WarpingMomentOfInertia and returns the result.
// The result is a new WarpingMomentOfInertia instance with the product of the values.
//
// Parameters:
//    other: The WarpingMomentOfInertia to multiply with the current WarpingMomentOfInertia.
//
// Returns:
//    *WarpingMomentOfInertia: A new WarpingMomentOfInertia instance representing the product of both WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Multiply(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide divides the current WarpingMomentOfInertia by the given WarpingMomentOfInertia and returns the result.
// The result is a new WarpingMomentOfInertia instance with the quotient of the values.
//
// Parameters:
//    other: The WarpingMomentOfInertia to divide the current WarpingMomentOfInertia by.
//
// Returns:
//    *WarpingMomentOfInertia: A new WarpingMomentOfInertia instance representing the quotient of both WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Divide(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value / other.BaseValue()}
}

// GetWarpingMomentOfInertiaAbbreviation gets the unit abbreviation.
func GetWarpingMomentOfInertiaAbbreviation(unit WarpingMomentOfInertiaUnits) string {
	switch unit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return "m⁶" 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return "dm⁶" 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return "cm⁶" 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return "mm⁶" 
	case WarpingMomentOfInertiaFootToTheSixth:
		return "ft⁶" 
	case WarpingMomentOfInertiaInchToTheSixth:
		return "in⁶" 
	default:
		return ""
	}
}