// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// JerkUnits defines various units of Jerk.
type JerkUnits string

const (
    
        // 
        JerkMeterPerSecondCubed JerkUnits = "MeterPerSecondCubed"
        // 
        JerkInchPerSecondCubed JerkUnits = "InchPerSecondCubed"
        // 
        JerkFootPerSecondCubed JerkUnits = "FootPerSecondCubed"
        // 
        JerkStandardGravitiesPerSecond JerkUnits = "StandardGravitiesPerSecond"
        // 
        JerkNanometerPerSecondCubed JerkUnits = "NanometerPerSecondCubed"
        // 
        JerkMicrometerPerSecondCubed JerkUnits = "MicrometerPerSecondCubed"
        // 
        JerkMillimeterPerSecondCubed JerkUnits = "MillimeterPerSecondCubed"
        // 
        JerkCentimeterPerSecondCubed JerkUnits = "CentimeterPerSecondCubed"
        // 
        JerkDecimeterPerSecondCubed JerkUnits = "DecimeterPerSecondCubed"
        // 
        JerkKilometerPerSecondCubed JerkUnits = "KilometerPerSecondCubed"
        // 
        JerkMillistandardGravitiesPerSecond JerkUnits = "MillistandardGravitiesPerSecond"
)

// JerkDto represents a Jerk measurement with a numerical value and its corresponding unit.
type JerkDto struct {
    // Value is the numerical representation of the Jerk.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Jerk, as defined in the JerkUnits enumeration.
	Unit  JerkUnits `json:"unit"`
}

// JerkDtoFactory groups methods for creating and serializing JerkDto objects.
type JerkDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a JerkDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf JerkDtoFactory) FromJSON(data []byte) (*JerkDto, error) {
	a := JerkDto{}

    // Parse JSON into JerkDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a JerkDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a JerkDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Jerk represents a measurement in a of Jerk.
//
// None
type Jerk struct {
	// value is the base measurement stored internally.
	value       float64
    
    meters_per_second_cubedLazy *float64 
    inches_per_second_cubedLazy *float64 
    feet_per_second_cubedLazy *float64 
    standard_gravities_per_secondLazy *float64 
    nanometers_per_second_cubedLazy *float64 
    micrometers_per_second_cubedLazy *float64 
    millimeters_per_second_cubedLazy *float64 
    centimeters_per_second_cubedLazy *float64 
    decimeters_per_second_cubedLazy *float64 
    kilometers_per_second_cubedLazy *float64 
    millistandard_gravities_per_secondLazy *float64 
}

// JerkFactory groups methods for creating Jerk instances.
type JerkFactory struct{}

// CreateJerk creates a new Jerk instance from the given value and unit.
func (uf JerkFactory) CreateJerk(value float64, unit JerkUnits) (*Jerk, error) {
	return newJerk(value, unit)
}

// FromDto converts a JerkDto to a Jerk instance.
func (uf JerkFactory) FromDto(dto JerkDto) (*Jerk, error) {
	return newJerk(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Jerk instance.
func (uf JerkFactory) FromDtoJSON(data []byte) (*Jerk, error) {
	unitDto, err := JerkDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JerkDto from JSON: %w", err)
	}
	return JerkFactory{}.FromDto(*unitDto)
}


// FromMetersPerSecondCubed creates a new Jerk instance from a value in MetersPerSecondCubed.
func (uf JerkFactory) FromMetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMeterPerSecondCubed)
}

// FromInchesPerSecondCubed creates a new Jerk instance from a value in InchesPerSecondCubed.
func (uf JerkFactory) FromInchesPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkInchPerSecondCubed)
}

// FromFeetPerSecondCubed creates a new Jerk instance from a value in FeetPerSecondCubed.
func (uf JerkFactory) FromFeetPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkFootPerSecondCubed)
}

// FromStandardGravitiesPerSecond creates a new Jerk instance from a value in StandardGravitiesPerSecond.
func (uf JerkFactory) FromStandardGravitiesPerSecond(value float64) (*Jerk, error) {
	return newJerk(value, JerkStandardGravitiesPerSecond)
}

// FromNanometersPerSecondCubed creates a new Jerk instance from a value in NanometersPerSecondCubed.
func (uf JerkFactory) FromNanometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkNanometerPerSecondCubed)
}

// FromMicrometersPerSecondCubed creates a new Jerk instance from a value in MicrometersPerSecondCubed.
func (uf JerkFactory) FromMicrometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMicrometerPerSecondCubed)
}

// FromMillimetersPerSecondCubed creates a new Jerk instance from a value in MillimetersPerSecondCubed.
func (uf JerkFactory) FromMillimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMillimeterPerSecondCubed)
}

// FromCentimetersPerSecondCubed creates a new Jerk instance from a value in CentimetersPerSecondCubed.
func (uf JerkFactory) FromCentimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkCentimeterPerSecondCubed)
}

// FromDecimetersPerSecondCubed creates a new Jerk instance from a value in DecimetersPerSecondCubed.
func (uf JerkFactory) FromDecimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkDecimeterPerSecondCubed)
}

// FromKilometersPerSecondCubed creates a new Jerk instance from a value in KilometersPerSecondCubed.
func (uf JerkFactory) FromKilometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkKilometerPerSecondCubed)
}

// FromMillistandardGravitiesPerSecond creates a new Jerk instance from a value in MillistandardGravitiesPerSecond.
func (uf JerkFactory) FromMillistandardGravitiesPerSecond(value float64) (*Jerk, error) {
	return newJerk(value, JerkMillistandardGravitiesPerSecond)
}


// newJerk creates a new Jerk.
func newJerk(value float64, fromUnit JerkUnits) (*Jerk, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Jerk{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Jerk in MeterPerSecondCubed unit (the base/default quantity).
func (a *Jerk) BaseValue() float64 {
	return a.value
}


// MetersPerSecondCubed returns the Jerk value in MetersPerSecondCubed.
//
// 
func (a *Jerk) MetersPerSecondCubed() float64 {
	if a.meters_per_second_cubedLazy != nil {
		return *a.meters_per_second_cubedLazy
	}
	meters_per_second_cubed := a.convertFromBase(JerkMeterPerSecondCubed)
	a.meters_per_second_cubedLazy = &meters_per_second_cubed
	return meters_per_second_cubed
}

// InchesPerSecondCubed returns the Jerk value in InchesPerSecondCubed.
//
// 
func (a *Jerk) InchesPerSecondCubed() float64 {
	if a.inches_per_second_cubedLazy != nil {
		return *a.inches_per_second_cubedLazy
	}
	inches_per_second_cubed := a.convertFromBase(JerkInchPerSecondCubed)
	a.inches_per_second_cubedLazy = &inches_per_second_cubed
	return inches_per_second_cubed
}

// FeetPerSecondCubed returns the Jerk value in FeetPerSecondCubed.
//
// 
func (a *Jerk) FeetPerSecondCubed() float64 {
	if a.feet_per_second_cubedLazy != nil {
		return *a.feet_per_second_cubedLazy
	}
	feet_per_second_cubed := a.convertFromBase(JerkFootPerSecondCubed)
	a.feet_per_second_cubedLazy = &feet_per_second_cubed
	return feet_per_second_cubed
}

// StandardGravitiesPerSecond returns the Jerk value in StandardGravitiesPerSecond.
//
// 
func (a *Jerk) StandardGravitiesPerSecond() float64 {
	if a.standard_gravities_per_secondLazy != nil {
		return *a.standard_gravities_per_secondLazy
	}
	standard_gravities_per_second := a.convertFromBase(JerkStandardGravitiesPerSecond)
	a.standard_gravities_per_secondLazy = &standard_gravities_per_second
	return standard_gravities_per_second
}

// NanometersPerSecondCubed returns the Jerk value in NanometersPerSecondCubed.
//
// 
func (a *Jerk) NanometersPerSecondCubed() float64 {
	if a.nanometers_per_second_cubedLazy != nil {
		return *a.nanometers_per_second_cubedLazy
	}
	nanometers_per_second_cubed := a.convertFromBase(JerkNanometerPerSecondCubed)
	a.nanometers_per_second_cubedLazy = &nanometers_per_second_cubed
	return nanometers_per_second_cubed
}

// MicrometersPerSecondCubed returns the Jerk value in MicrometersPerSecondCubed.
//
// 
func (a *Jerk) MicrometersPerSecondCubed() float64 {
	if a.micrometers_per_second_cubedLazy != nil {
		return *a.micrometers_per_second_cubedLazy
	}
	micrometers_per_second_cubed := a.convertFromBase(JerkMicrometerPerSecondCubed)
	a.micrometers_per_second_cubedLazy = &micrometers_per_second_cubed
	return micrometers_per_second_cubed
}

// MillimetersPerSecondCubed returns the Jerk value in MillimetersPerSecondCubed.
//
// 
func (a *Jerk) MillimetersPerSecondCubed() float64 {
	if a.millimeters_per_second_cubedLazy != nil {
		return *a.millimeters_per_second_cubedLazy
	}
	millimeters_per_second_cubed := a.convertFromBase(JerkMillimeterPerSecondCubed)
	a.millimeters_per_second_cubedLazy = &millimeters_per_second_cubed
	return millimeters_per_second_cubed
}

// CentimetersPerSecondCubed returns the Jerk value in CentimetersPerSecondCubed.
//
// 
func (a *Jerk) CentimetersPerSecondCubed() float64 {
	if a.centimeters_per_second_cubedLazy != nil {
		return *a.centimeters_per_second_cubedLazy
	}
	centimeters_per_second_cubed := a.convertFromBase(JerkCentimeterPerSecondCubed)
	a.centimeters_per_second_cubedLazy = &centimeters_per_second_cubed
	return centimeters_per_second_cubed
}

// DecimetersPerSecondCubed returns the Jerk value in DecimetersPerSecondCubed.
//
// 
func (a *Jerk) DecimetersPerSecondCubed() float64 {
	if a.decimeters_per_second_cubedLazy != nil {
		return *a.decimeters_per_second_cubedLazy
	}
	decimeters_per_second_cubed := a.convertFromBase(JerkDecimeterPerSecondCubed)
	a.decimeters_per_second_cubedLazy = &decimeters_per_second_cubed
	return decimeters_per_second_cubed
}

// KilometersPerSecondCubed returns the Jerk value in KilometersPerSecondCubed.
//
// 
func (a *Jerk) KilometersPerSecondCubed() float64 {
	if a.kilometers_per_second_cubedLazy != nil {
		return *a.kilometers_per_second_cubedLazy
	}
	kilometers_per_second_cubed := a.convertFromBase(JerkKilometerPerSecondCubed)
	a.kilometers_per_second_cubedLazy = &kilometers_per_second_cubed
	return kilometers_per_second_cubed
}

// MillistandardGravitiesPerSecond returns the Jerk value in MillistandardGravitiesPerSecond.
//
// 
func (a *Jerk) MillistandardGravitiesPerSecond() float64 {
	if a.millistandard_gravities_per_secondLazy != nil {
		return *a.millistandard_gravities_per_secondLazy
	}
	millistandard_gravities_per_second := a.convertFromBase(JerkMillistandardGravitiesPerSecond)
	a.millistandard_gravities_per_secondLazy = &millistandard_gravities_per_second
	return millistandard_gravities_per_second
}


// ToDto creates a JerkDto representation from the Jerk instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecondCubed by default.
func (a *Jerk) ToDto(holdInUnit *JerkUnits) JerkDto {
	if holdInUnit == nil {
		defaultUnit := JerkMeterPerSecondCubed // Default value
		holdInUnit = &defaultUnit
	}

	return JerkDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Jerk instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecondCubed by default.
func (a *Jerk) ToDtoJSON(holdInUnit *JerkUnits) ([]byte, error) {
	// Convert to JerkDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Jerk to a specific unit value.
// The function uses the provided unit type (JerkUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Jerk) Convert(toUnit JerkUnits) float64 {
	switch toUnit { 
    case JerkMeterPerSecondCubed:
		return a.MetersPerSecondCubed()
    case JerkInchPerSecondCubed:
		return a.InchesPerSecondCubed()
    case JerkFootPerSecondCubed:
		return a.FeetPerSecondCubed()
    case JerkStandardGravitiesPerSecond:
		return a.StandardGravitiesPerSecond()
    case JerkNanometerPerSecondCubed:
		return a.NanometersPerSecondCubed()
    case JerkMicrometerPerSecondCubed:
		return a.MicrometersPerSecondCubed()
    case JerkMillimeterPerSecondCubed:
		return a.MillimetersPerSecondCubed()
    case JerkCentimeterPerSecondCubed:
		return a.CentimetersPerSecondCubed()
    case JerkDecimeterPerSecondCubed:
		return a.DecimetersPerSecondCubed()
    case JerkKilometerPerSecondCubed:
		return a.KilometersPerSecondCubed()
    case JerkMillistandardGravitiesPerSecond:
		return a.MillistandardGravitiesPerSecond()
	default:
		return math.NaN()
	}
}

func (a *Jerk) convertFromBase(toUnit JerkUnits) float64 {
    value := a.value
	switch toUnit { 
	case JerkMeterPerSecondCubed:
		return (value) 
	case JerkInchPerSecondCubed:
		return (value / 0.0254) 
	case JerkFootPerSecondCubed:
		return (value / 0.304800) 
	case JerkStandardGravitiesPerSecond:
		return (value / 9.80665) 
	case JerkNanometerPerSecondCubed:
		return ((value) / 1e-09) 
	case JerkMicrometerPerSecondCubed:
		return ((value) / 1e-06) 
	case JerkMillimeterPerSecondCubed:
		return ((value) / 0.001) 
	case JerkCentimeterPerSecondCubed:
		return ((value) / 0.01) 
	case JerkDecimeterPerSecondCubed:
		return ((value) / 0.1) 
	case JerkKilometerPerSecondCubed:
		return ((value) / 1000.0) 
	case JerkMillistandardGravitiesPerSecond:
		return ((value / 9.80665) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Jerk) convertToBase(value float64, fromUnit JerkUnits) float64 {
	switch fromUnit { 
	case JerkMeterPerSecondCubed:
		return (value) 
	case JerkInchPerSecondCubed:
		return (value * 0.0254) 
	case JerkFootPerSecondCubed:
		return (value * 0.304800) 
	case JerkStandardGravitiesPerSecond:
		return (value * 9.80665) 
	case JerkNanometerPerSecondCubed:
		return ((value) * 1e-09) 
	case JerkMicrometerPerSecondCubed:
		return ((value) * 1e-06) 
	case JerkMillimeterPerSecondCubed:
		return ((value) * 0.001) 
	case JerkCentimeterPerSecondCubed:
		return ((value) * 0.01) 
	case JerkDecimeterPerSecondCubed:
		return ((value) * 0.1) 
	case JerkKilometerPerSecondCubed:
		return ((value) * 1000.0) 
	case JerkMillistandardGravitiesPerSecond:
		return ((value * 9.80665) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Jerk in the default unit (MeterPerSecondCubed),
// formatted to two decimal places.
func (a Jerk) String() string {
	return a.ToString(JerkMeterPerSecondCubed, 2)
}

// ToString formats the Jerk value as a string with the specified unit and fractional digits.
// It converts the Jerk to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Jerk value will be converted (e.g., MeterPerSecondCubed))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Jerk with the unit abbreviation.
func (a *Jerk) ToString(unit JerkUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Jerk) getUnitAbbreviation(unit JerkUnits) string {
	switch unit { 
	case JerkMeterPerSecondCubed:
		return "m/s³" 
	case JerkInchPerSecondCubed:
		return "in/s³" 
	case JerkFootPerSecondCubed:
		return "ft/s³" 
	case JerkStandardGravitiesPerSecond:
		return "g/s" 
	case JerkNanometerPerSecondCubed:
		return "nm/s³" 
	case JerkMicrometerPerSecondCubed:
		return "μm/s³" 
	case JerkMillimeterPerSecondCubed:
		return "mm/s³" 
	case JerkCentimeterPerSecondCubed:
		return "cm/s³" 
	case JerkDecimeterPerSecondCubed:
		return "dm/s³" 
	case JerkKilometerPerSecondCubed:
		return "km/s³" 
	case JerkMillistandardGravitiesPerSecond:
		return "mg/s" 
	default:
		return ""
	}
}

// Equals checks if the given Jerk is equal to the current Jerk.
//
// Parameters:
//    other: The Jerk to compare against.
//
// Returns:
//    bool: Returns true if both Jerk are equal, false otherwise.
func (a *Jerk) Equals(other *Jerk) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Jerk with another Jerk.
// It returns -1 if the current Jerk is less than the other Jerk, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Jerk to compare against.
//
// Returns:
//    int: -1 if the current Jerk is less, 1 if greater, and 0 if equal.
func (a *Jerk) CompareTo(other *Jerk) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Jerk to the current Jerk and returns the result.
// The result is a new Jerk instance with the sum of the values.
//
// Parameters:
//    other: The Jerk to add to the current Jerk.
//
// Returns:
//    *Jerk: A new Jerk instance representing the sum of both Jerk.
func (a *Jerk) Add(other *Jerk) *Jerk {
	return &Jerk{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Jerk from the current Jerk and returns the result.
// The result is a new Jerk instance with the difference of the values.
//
// Parameters:
//    other: The Jerk to subtract from the current Jerk.
//
// Returns:
//    *Jerk: A new Jerk instance representing the difference of both Jerk.
func (a *Jerk) Subtract(other *Jerk) *Jerk {
	return &Jerk{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Jerk by the given Jerk and returns the result.
// The result is a new Jerk instance with the product of the values.
//
// Parameters:
//    other: The Jerk to multiply with the current Jerk.
//
// Returns:
//    *Jerk: A new Jerk instance representing the product of both Jerk.
func (a *Jerk) Multiply(other *Jerk) *Jerk {
	return &Jerk{value: a.value * other.BaseValue()}
}

// Divide divides the current Jerk by the given Jerk and returns the result.
// The result is a new Jerk instance with the quotient of the values.
//
// Parameters:
//    other: The Jerk to divide the current Jerk by.
//
// Returns:
//    *Jerk: A new Jerk instance representing the quotient of both Jerk.
func (a *Jerk) Divide(other *Jerk) *Jerk {
	return &Jerk{value: a.value / other.BaseValue()}
}