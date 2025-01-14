// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalStiffnessPerLengthUnits defines various units of RotationalStiffnessPerLength.
type RotationalStiffnessPerLengthUnits string

const (
    
        // 
        RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "NewtonMeterPerRadianPerMeter"
        // 
        RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot RotationalStiffnessPerLengthUnits = "PoundForceFootPerDegreesPerFoot"
        // 
        RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot RotationalStiffnessPerLengthUnits = "KilopoundForceFootPerDegreesPerFoot"
        // 
        RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "KilonewtonMeterPerRadianPerMeter"
        // 
        RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "MeganewtonMeterPerRadianPerMeter"
)

// RotationalStiffnessPerLengthDto represents a RotationalStiffnessPerLength measurement with a numerical value and its corresponding unit.
type RotationalStiffnessPerLengthDto struct {
    // Value is the numerical representation of the RotationalStiffnessPerLength.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RotationalStiffnessPerLength, as defined in the RotationalStiffnessPerLengthUnits enumeration.
	Unit  RotationalStiffnessPerLengthUnits `json:"unit"`
}

// RotationalStiffnessPerLengthDtoFactory groups methods for creating and serializing RotationalStiffnessPerLengthDto objects.
type RotationalStiffnessPerLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RotationalStiffnessPerLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RotationalStiffnessPerLengthDtoFactory) FromJSON(data []byte) (*RotationalStiffnessPerLengthDto, error) {
	a := RotationalStiffnessPerLengthDto{}

    // Parse JSON into RotationalStiffnessPerLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RotationalStiffnessPerLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RotationalStiffnessPerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RotationalStiffnessPerLength represents a measurement in a of RotationalStiffnessPerLength.
//
// https://en.wikipedia.org/wiki/Stiffness#Rotational_stiffness
type RotationalStiffnessPerLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    newton_meters_per_radian_per_meterLazy *float64 
    pound_force_feet_per_degrees_per_feetLazy *float64 
    kilopound_force_feet_per_degrees_per_feetLazy *float64 
    kilonewton_meters_per_radian_per_meterLazy *float64 
    meganewton_meters_per_radian_per_meterLazy *float64 
}

// RotationalStiffnessPerLengthFactory groups methods for creating RotationalStiffnessPerLength instances.
type RotationalStiffnessPerLengthFactory struct{}

// CreateRotationalStiffnessPerLength creates a new RotationalStiffnessPerLength instance from the given value and unit.
func (uf RotationalStiffnessPerLengthFactory) CreateRotationalStiffnessPerLength(value float64, unit RotationalStiffnessPerLengthUnits) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, unit)
}

// FromDto converts a RotationalStiffnessPerLengthDto to a RotationalStiffnessPerLength instance.
func (uf RotationalStiffnessPerLengthFactory) FromDto(dto RotationalStiffnessPerLengthDto) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RotationalStiffnessPerLength instance.
func (uf RotationalStiffnessPerLengthFactory) FromDtoJSON(data []byte) (*RotationalStiffnessPerLength, error) {
	unitDto, err := RotationalStiffnessPerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RotationalStiffnessPerLengthDto from JSON: %w", err)
	}
	return RotationalStiffnessPerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonMetersPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from a value in NewtonMetersPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromNewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
}

// FromPoundForceFeetPerDegreesPerFeet creates a new RotationalStiffnessPerLength instance from a value in PoundForceFeetPerDegreesPerFeet.
func (uf RotationalStiffnessPerLengthFactory) FromPoundForceFeetPerDegreesPerFeet(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
}

// FromKilopoundForceFeetPerDegreesPerFeet creates a new RotationalStiffnessPerLength instance from a value in KilopoundForceFeetPerDegreesPerFeet.
func (uf RotationalStiffnessPerLengthFactory) FromKilopoundForceFeetPerDegreesPerFeet(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
}

// FromKilonewtonMetersPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from a value in KilonewtonMetersPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromKilonewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
}

// FromMeganewtonMetersPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from a value in MeganewtonMetersPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromMeganewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
}


// newRotationalStiffnessPerLength creates a new RotationalStiffnessPerLength.
func newRotationalStiffnessPerLength(value float64, fromUnit RotationalStiffnessPerLengthUnits) (*RotationalStiffnessPerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalStiffnessPerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalStiffnessPerLength in NewtonMeterPerRadianPerMeter unit (the base/default quantity).
func (a *RotationalStiffnessPerLength) BaseValue() float64 {
	return a.value
}


// NewtonMetersPerRadianPerMeter returns the RotationalStiffnessPerLength value in NewtonMetersPerRadianPerMeter.
//
// 
func (a *RotationalStiffnessPerLength) NewtonMetersPerRadianPerMeter() float64 {
	if a.newton_meters_per_radian_per_meterLazy != nil {
		return *a.newton_meters_per_radian_per_meterLazy
	}
	newton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a.newton_meters_per_radian_per_meterLazy = &newton_meters_per_radian_per_meter
	return newton_meters_per_radian_per_meter
}

// PoundForceFeetPerDegreesPerFeet returns the RotationalStiffnessPerLength value in PoundForceFeetPerDegreesPerFeet.
//
// 
func (a *RotationalStiffnessPerLength) PoundForceFeetPerDegreesPerFeet() float64 {
	if a.pound_force_feet_per_degrees_per_feetLazy != nil {
		return *a.pound_force_feet_per_degrees_per_feetLazy
	}
	pound_force_feet_per_degrees_per_feet := a.convertFromBase(RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
	a.pound_force_feet_per_degrees_per_feetLazy = &pound_force_feet_per_degrees_per_feet
	return pound_force_feet_per_degrees_per_feet
}

// KilopoundForceFeetPerDegreesPerFeet returns the RotationalStiffnessPerLength value in KilopoundForceFeetPerDegreesPerFeet.
//
// 
func (a *RotationalStiffnessPerLength) KilopoundForceFeetPerDegreesPerFeet() float64 {
	if a.kilopound_force_feet_per_degrees_per_feetLazy != nil {
		return *a.kilopound_force_feet_per_degrees_per_feetLazy
	}
	kilopound_force_feet_per_degrees_per_feet := a.convertFromBase(RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
	a.kilopound_force_feet_per_degrees_per_feetLazy = &kilopound_force_feet_per_degrees_per_feet
	return kilopound_force_feet_per_degrees_per_feet
}

// KilonewtonMetersPerRadianPerMeter returns the RotationalStiffnessPerLength value in KilonewtonMetersPerRadianPerMeter.
//
// 
func (a *RotationalStiffnessPerLength) KilonewtonMetersPerRadianPerMeter() float64 {
	if a.kilonewton_meters_per_radian_per_meterLazy != nil {
		return *a.kilonewton_meters_per_radian_per_meterLazy
	}
	kilonewton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
	a.kilonewton_meters_per_radian_per_meterLazy = &kilonewton_meters_per_radian_per_meter
	return kilonewton_meters_per_radian_per_meter
}

// MeganewtonMetersPerRadianPerMeter returns the RotationalStiffnessPerLength value in MeganewtonMetersPerRadianPerMeter.
//
// 
func (a *RotationalStiffnessPerLength) MeganewtonMetersPerRadianPerMeter() float64 {
	if a.meganewton_meters_per_radian_per_meterLazy != nil {
		return *a.meganewton_meters_per_radian_per_meterLazy
	}
	meganewton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
	a.meganewton_meters_per_radian_per_meterLazy = &meganewton_meters_per_radian_per_meter
	return meganewton_meters_per_radian_per_meter
}


// ToDto creates a RotationalStiffnessPerLengthDto representation from the RotationalStiffnessPerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerRadianPerMeter by default.
func (a *RotationalStiffnessPerLength) ToDto(holdInUnit *RotationalStiffnessPerLengthUnits) RotationalStiffnessPerLengthDto {
	if holdInUnit == nil {
		defaultUnit := RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalStiffnessPerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RotationalStiffnessPerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerRadianPerMeter by default.
func (a *RotationalStiffnessPerLength) ToDtoJSON(holdInUnit *RotationalStiffnessPerLengthUnits) ([]byte, error) {
	// Convert to RotationalStiffnessPerLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RotationalStiffnessPerLength to a specific unit value.
// The function uses the provided unit type (RotationalStiffnessPerLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RotationalStiffnessPerLength) Convert(toUnit RotationalStiffnessPerLengthUnits) float64 {
	switch toUnit { 
    case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return a.NewtonMetersPerRadianPerMeter()
    case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return a.PoundForceFeetPerDegreesPerFeet()
    case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return a.KilopoundForceFeetPerDegreesPerFeet()
    case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return a.KilonewtonMetersPerRadianPerMeter()
    case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return a.MeganewtonMetersPerRadianPerMeter()
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffnessPerLength) convertFromBase(toUnit RotationalStiffnessPerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return (value) 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return (value / 254.864324570) 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return (value / 254864.324570) 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return ((value) / 1000.0) 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffnessPerLength) convertToBase(value float64, fromUnit RotationalStiffnessPerLengthUnits) float64 {
	switch fromUnit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return (value) 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return (value * 254.864324570) 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return (value * 254864.324570) 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return ((value) * 1000.0) 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RotationalStiffnessPerLength in the default unit (NewtonMeterPerRadianPerMeter),
// formatted to two decimal places.
func (a RotationalStiffnessPerLength) String() string {
	return a.ToString(RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, 2)
}

// ToString formats the RotationalStiffnessPerLength value as a string with the specified unit and fractional digits.
// It converts the RotationalStiffnessPerLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RotationalStiffnessPerLength value will be converted (e.g., NewtonMeterPerRadianPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RotationalStiffnessPerLength with the unit abbreviation.
func (a *RotationalStiffnessPerLength) ToString(unit RotationalStiffnessPerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalStiffnessPerLength) getUnitAbbreviation(unit RotationalStiffnessPerLengthUnits) string {
	switch unit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return "N·m/rad/m" 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return "lbf·ft/deg/ft" 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return "kipf·ft/°/ft" 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return "kN·m/rad/m" 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return "MN·m/rad/m" 
	default:
		return ""
	}
}

// Equals checks if the given RotationalStiffnessPerLength is equal to the current RotationalStiffnessPerLength.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to compare against.
//
// Returns:
//    bool: Returns true if both RotationalStiffnessPerLength are equal, false otherwise.
func (a *RotationalStiffnessPerLength) Equals(other *RotationalStiffnessPerLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RotationalStiffnessPerLength with another RotationalStiffnessPerLength.
// It returns -1 if the current RotationalStiffnessPerLength is less than the other RotationalStiffnessPerLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to compare against.
//
// Returns:
//    int: -1 if the current RotationalStiffnessPerLength is less, 1 if greater, and 0 if equal.
func (a *RotationalStiffnessPerLength) CompareTo(other *RotationalStiffnessPerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RotationalStiffnessPerLength to the current RotationalStiffnessPerLength and returns the result.
// The result is a new RotationalStiffnessPerLength instance with the sum of the values.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to add to the current RotationalStiffnessPerLength.
//
// Returns:
//    *RotationalStiffnessPerLength: A new RotationalStiffnessPerLength instance representing the sum of both RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Add(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RotationalStiffnessPerLength from the current RotationalStiffnessPerLength and returns the result.
// The result is a new RotationalStiffnessPerLength instance with the difference of the values.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to subtract from the current RotationalStiffnessPerLength.
//
// Returns:
//    *RotationalStiffnessPerLength: A new RotationalStiffnessPerLength instance representing the difference of both RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Subtract(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RotationalStiffnessPerLength by the given RotationalStiffnessPerLength and returns the result.
// The result is a new RotationalStiffnessPerLength instance with the product of the values.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to multiply with the current RotationalStiffnessPerLength.
//
// Returns:
//    *RotationalStiffnessPerLength: A new RotationalStiffnessPerLength instance representing the product of both RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Multiply(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value * other.BaseValue()}
}

// Divide divides the current RotationalStiffnessPerLength by the given RotationalStiffnessPerLength and returns the result.
// The result is a new RotationalStiffnessPerLength instance with the quotient of the values.
//
// Parameters:
//    other: The RotationalStiffnessPerLength to divide the current RotationalStiffnessPerLength by.
//
// Returns:
//    *RotationalStiffnessPerLength: A new RotationalStiffnessPerLength instance representing the quotient of both RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Divide(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value / other.BaseValue()}
}