// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AccelerationUnits defines various units of Acceleration.
type AccelerationUnits string

const (
    
        // 
        AccelerationMeterPerSecondSquared AccelerationUnits = "MeterPerSecondSquared"
        // 
        AccelerationInchPerSecondSquared AccelerationUnits = "InchPerSecondSquared"
        // 
        AccelerationFootPerSecondSquared AccelerationUnits = "FootPerSecondSquared"
        // 
        AccelerationKnotPerSecond AccelerationUnits = "KnotPerSecond"
        // 
        AccelerationKnotPerMinute AccelerationUnits = "KnotPerMinute"
        // 
        AccelerationKnotPerHour AccelerationUnits = "KnotPerHour"
        // 
        AccelerationStandardGravity AccelerationUnits = "StandardGravity"
        // 
        AccelerationNanometerPerSecondSquared AccelerationUnits = "NanometerPerSecondSquared"
        // 
        AccelerationMicrometerPerSecondSquared AccelerationUnits = "MicrometerPerSecondSquared"
        // 
        AccelerationMillimeterPerSecondSquared AccelerationUnits = "MillimeterPerSecondSquared"
        // 
        AccelerationCentimeterPerSecondSquared AccelerationUnits = "CentimeterPerSecondSquared"
        // 
        AccelerationDecimeterPerSecondSquared AccelerationUnits = "DecimeterPerSecondSquared"
        // 
        AccelerationKilometerPerSecondSquared AccelerationUnits = "KilometerPerSecondSquared"
        // 
        AccelerationMillistandardGravity AccelerationUnits = "MillistandardGravity"
)

// AccelerationDto represents a Acceleration measurement with a numerical value and its corresponding unit.
type AccelerationDto struct {
    // Value is the numerical representation of the Acceleration.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Acceleration, as defined in the AccelerationUnits enumeration.
	Unit  AccelerationUnits `json:"unit"`
}

// AccelerationDtoFactory groups methods for creating and serializing AccelerationDto objects.
type AccelerationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AccelerationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AccelerationDtoFactory) FromJSON(data []byte) (*AccelerationDto, error) {
	a := AccelerationDto{}

    // Parse JSON into AccelerationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a AccelerationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AccelerationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Acceleration represents a measurement in a of Acceleration.
//
// Acceleration, in physics, is the rate at which the velocity of an object changes over time. An object's acceleration is the net result of any and all forces acting on the object, as described by Newton's Second Law. The SI unit for acceleration is the Meter per second squared (m/s²). Accelerations are vector quantities (they have magnitude and direction) and add according to the parallelogram law. As a vector, the calculated net force is equal to the product of the object's mass (a scalar quantity) and the acceleration.
type Acceleration struct {
	// value is the base measurement stored internally.
	value       float64
    
    meters_per_second_squaredLazy *float64 
    inches_per_second_squaredLazy *float64 
    feet_per_second_squaredLazy *float64 
    knots_per_secondLazy *float64 
    knots_per_minuteLazy *float64 
    knots_per_hourLazy *float64 
    standard_gravityLazy *float64 
    nanometers_per_second_squaredLazy *float64 
    micrometers_per_second_squaredLazy *float64 
    millimeters_per_second_squaredLazy *float64 
    centimeters_per_second_squaredLazy *float64 
    decimeters_per_second_squaredLazy *float64 
    kilometers_per_second_squaredLazy *float64 
    millistandard_gravityLazy *float64 
}

// AccelerationFactory groups methods for creating Acceleration instances.
type AccelerationFactory struct{}

// CreateAcceleration creates a new Acceleration instance from the given value and unit.
func (uf AccelerationFactory) CreateAcceleration(value float64, unit AccelerationUnits) (*Acceleration, error) {
	return newAcceleration(value, unit)
}

// FromDto converts a AccelerationDto to a Acceleration instance.
func (uf AccelerationFactory) FromDto(dto AccelerationDto) (*Acceleration, error) {
	return newAcceleration(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Acceleration instance.
func (uf AccelerationFactory) FromDtoJSON(data []byte) (*Acceleration, error) {
	unitDto, err := AccelerationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AccelerationDto from JSON: %w", err)
	}
	return AccelerationFactory{}.FromDto(*unitDto)
}


// FromMetersPerSecondSquared creates a new Acceleration instance from a value in MetersPerSecondSquared.
func (uf AccelerationFactory) FromMetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMeterPerSecondSquared)
}

// FromInchesPerSecondSquared creates a new Acceleration instance from a value in InchesPerSecondSquared.
func (uf AccelerationFactory) FromInchesPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationInchPerSecondSquared)
}

// FromFeetPerSecondSquared creates a new Acceleration instance from a value in FeetPerSecondSquared.
func (uf AccelerationFactory) FromFeetPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationFootPerSecondSquared)
}

// FromKnotsPerSecond creates a new Acceleration instance from a value in KnotsPerSecond.
func (uf AccelerationFactory) FromKnotsPerSecond(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerSecond)
}

// FromKnotsPerMinute creates a new Acceleration instance from a value in KnotsPerMinute.
func (uf AccelerationFactory) FromKnotsPerMinute(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerMinute)
}

// FromKnotsPerHour creates a new Acceleration instance from a value in KnotsPerHour.
func (uf AccelerationFactory) FromKnotsPerHour(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerHour)
}

// FromStandardGravity creates a new Acceleration instance from a value in StandardGravity.
func (uf AccelerationFactory) FromStandardGravity(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationStandardGravity)
}

// FromNanometersPerSecondSquared creates a new Acceleration instance from a value in NanometersPerSecondSquared.
func (uf AccelerationFactory) FromNanometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationNanometerPerSecondSquared)
}

// FromMicrometersPerSecondSquared creates a new Acceleration instance from a value in MicrometersPerSecondSquared.
func (uf AccelerationFactory) FromMicrometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMicrometerPerSecondSquared)
}

// FromMillimetersPerSecondSquared creates a new Acceleration instance from a value in MillimetersPerSecondSquared.
func (uf AccelerationFactory) FromMillimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMillimeterPerSecondSquared)
}

// FromCentimetersPerSecondSquared creates a new Acceleration instance from a value in CentimetersPerSecondSquared.
func (uf AccelerationFactory) FromCentimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationCentimeterPerSecondSquared)
}

// FromDecimetersPerSecondSquared creates a new Acceleration instance from a value in DecimetersPerSecondSquared.
func (uf AccelerationFactory) FromDecimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationDecimeterPerSecondSquared)
}

// FromKilometersPerSecondSquared creates a new Acceleration instance from a value in KilometersPerSecondSquared.
func (uf AccelerationFactory) FromKilometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKilometerPerSecondSquared)
}

// FromMillistandardGravity creates a new Acceleration instance from a value in MillistandardGravity.
func (uf AccelerationFactory) FromMillistandardGravity(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMillistandardGravity)
}


// newAcceleration creates a new Acceleration.
func newAcceleration(value float64, fromUnit AccelerationUnits) (*Acceleration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Acceleration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Acceleration in MeterPerSecondSquared unit (the base/default quantity).
func (a *Acceleration) BaseValue() float64 {
	return a.value
}


// MetersPerSecondSquared returns the Acceleration value in MetersPerSecondSquared.
//
// 
func (a *Acceleration) MetersPerSecondSquared() float64 {
	if a.meters_per_second_squaredLazy != nil {
		return *a.meters_per_second_squaredLazy
	}
	meters_per_second_squared := a.convertFromBase(AccelerationMeterPerSecondSquared)
	a.meters_per_second_squaredLazy = &meters_per_second_squared
	return meters_per_second_squared
}

// InchesPerSecondSquared returns the Acceleration value in InchesPerSecondSquared.
//
// 
func (a *Acceleration) InchesPerSecondSquared() float64 {
	if a.inches_per_second_squaredLazy != nil {
		return *a.inches_per_second_squaredLazy
	}
	inches_per_second_squared := a.convertFromBase(AccelerationInchPerSecondSquared)
	a.inches_per_second_squaredLazy = &inches_per_second_squared
	return inches_per_second_squared
}

// FeetPerSecondSquared returns the Acceleration value in FeetPerSecondSquared.
//
// 
func (a *Acceleration) FeetPerSecondSquared() float64 {
	if a.feet_per_second_squaredLazy != nil {
		return *a.feet_per_second_squaredLazy
	}
	feet_per_second_squared := a.convertFromBase(AccelerationFootPerSecondSquared)
	a.feet_per_second_squaredLazy = &feet_per_second_squared
	return feet_per_second_squared
}

// KnotsPerSecond returns the Acceleration value in KnotsPerSecond.
//
// 
func (a *Acceleration) KnotsPerSecond() float64 {
	if a.knots_per_secondLazy != nil {
		return *a.knots_per_secondLazy
	}
	knots_per_second := a.convertFromBase(AccelerationKnotPerSecond)
	a.knots_per_secondLazy = &knots_per_second
	return knots_per_second
}

// KnotsPerMinute returns the Acceleration value in KnotsPerMinute.
//
// 
func (a *Acceleration) KnotsPerMinute() float64 {
	if a.knots_per_minuteLazy != nil {
		return *a.knots_per_minuteLazy
	}
	knots_per_minute := a.convertFromBase(AccelerationKnotPerMinute)
	a.knots_per_minuteLazy = &knots_per_minute
	return knots_per_minute
}

// KnotsPerHour returns the Acceleration value in KnotsPerHour.
//
// 
func (a *Acceleration) KnotsPerHour() float64 {
	if a.knots_per_hourLazy != nil {
		return *a.knots_per_hourLazy
	}
	knots_per_hour := a.convertFromBase(AccelerationKnotPerHour)
	a.knots_per_hourLazy = &knots_per_hour
	return knots_per_hour
}

// StandardGravity returns the Acceleration value in StandardGravity.
//
// 
func (a *Acceleration) StandardGravity() float64 {
	if a.standard_gravityLazy != nil {
		return *a.standard_gravityLazy
	}
	standard_gravity := a.convertFromBase(AccelerationStandardGravity)
	a.standard_gravityLazy = &standard_gravity
	return standard_gravity
}

// NanometersPerSecondSquared returns the Acceleration value in NanometersPerSecondSquared.
//
// 
func (a *Acceleration) NanometersPerSecondSquared() float64 {
	if a.nanometers_per_second_squaredLazy != nil {
		return *a.nanometers_per_second_squaredLazy
	}
	nanometers_per_second_squared := a.convertFromBase(AccelerationNanometerPerSecondSquared)
	a.nanometers_per_second_squaredLazy = &nanometers_per_second_squared
	return nanometers_per_second_squared
}

// MicrometersPerSecondSquared returns the Acceleration value in MicrometersPerSecondSquared.
//
// 
func (a *Acceleration) MicrometersPerSecondSquared() float64 {
	if a.micrometers_per_second_squaredLazy != nil {
		return *a.micrometers_per_second_squaredLazy
	}
	micrometers_per_second_squared := a.convertFromBase(AccelerationMicrometerPerSecondSquared)
	a.micrometers_per_second_squaredLazy = &micrometers_per_second_squared
	return micrometers_per_second_squared
}

// MillimetersPerSecondSquared returns the Acceleration value in MillimetersPerSecondSquared.
//
// 
func (a *Acceleration) MillimetersPerSecondSquared() float64 {
	if a.millimeters_per_second_squaredLazy != nil {
		return *a.millimeters_per_second_squaredLazy
	}
	millimeters_per_second_squared := a.convertFromBase(AccelerationMillimeterPerSecondSquared)
	a.millimeters_per_second_squaredLazy = &millimeters_per_second_squared
	return millimeters_per_second_squared
}

// CentimetersPerSecondSquared returns the Acceleration value in CentimetersPerSecondSquared.
//
// 
func (a *Acceleration) CentimetersPerSecondSquared() float64 {
	if a.centimeters_per_second_squaredLazy != nil {
		return *a.centimeters_per_second_squaredLazy
	}
	centimeters_per_second_squared := a.convertFromBase(AccelerationCentimeterPerSecondSquared)
	a.centimeters_per_second_squaredLazy = &centimeters_per_second_squared
	return centimeters_per_second_squared
}

// DecimetersPerSecondSquared returns the Acceleration value in DecimetersPerSecondSquared.
//
// 
func (a *Acceleration) DecimetersPerSecondSquared() float64 {
	if a.decimeters_per_second_squaredLazy != nil {
		return *a.decimeters_per_second_squaredLazy
	}
	decimeters_per_second_squared := a.convertFromBase(AccelerationDecimeterPerSecondSquared)
	a.decimeters_per_second_squaredLazy = &decimeters_per_second_squared
	return decimeters_per_second_squared
}

// KilometersPerSecondSquared returns the Acceleration value in KilometersPerSecondSquared.
//
// 
func (a *Acceleration) KilometersPerSecondSquared() float64 {
	if a.kilometers_per_second_squaredLazy != nil {
		return *a.kilometers_per_second_squaredLazy
	}
	kilometers_per_second_squared := a.convertFromBase(AccelerationKilometerPerSecondSquared)
	a.kilometers_per_second_squaredLazy = &kilometers_per_second_squared
	return kilometers_per_second_squared
}

// MillistandardGravity returns the Acceleration value in MillistandardGravity.
//
// 
func (a *Acceleration) MillistandardGravity() float64 {
	if a.millistandard_gravityLazy != nil {
		return *a.millistandard_gravityLazy
	}
	millistandard_gravity := a.convertFromBase(AccelerationMillistandardGravity)
	a.millistandard_gravityLazy = &millistandard_gravity
	return millistandard_gravity
}


// ToDto creates a AccelerationDto representation from the Acceleration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecondSquared by default.
func (a *Acceleration) ToDto(holdInUnit *AccelerationUnits) AccelerationDto {
	if holdInUnit == nil {
		defaultUnit := AccelerationMeterPerSecondSquared // Default value
		holdInUnit = &defaultUnit
	}

	return AccelerationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Acceleration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecondSquared by default.
func (a *Acceleration) ToDtoJSON(holdInUnit *AccelerationUnits) ([]byte, error) {
	// Convert to AccelerationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Acceleration to a specific unit value.
// The function uses the provided unit type (AccelerationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Acceleration) Convert(toUnit AccelerationUnits) float64 {
	switch toUnit { 
    case AccelerationMeterPerSecondSquared:
		return a.MetersPerSecondSquared()
    case AccelerationInchPerSecondSquared:
		return a.InchesPerSecondSquared()
    case AccelerationFootPerSecondSquared:
		return a.FeetPerSecondSquared()
    case AccelerationKnotPerSecond:
		return a.KnotsPerSecond()
    case AccelerationKnotPerMinute:
		return a.KnotsPerMinute()
    case AccelerationKnotPerHour:
		return a.KnotsPerHour()
    case AccelerationStandardGravity:
		return a.StandardGravity()
    case AccelerationNanometerPerSecondSquared:
		return a.NanometersPerSecondSquared()
    case AccelerationMicrometerPerSecondSquared:
		return a.MicrometersPerSecondSquared()
    case AccelerationMillimeterPerSecondSquared:
		return a.MillimetersPerSecondSquared()
    case AccelerationCentimeterPerSecondSquared:
		return a.CentimetersPerSecondSquared()
    case AccelerationDecimeterPerSecondSquared:
		return a.DecimetersPerSecondSquared()
    case AccelerationKilometerPerSecondSquared:
		return a.KilometersPerSecondSquared()
    case AccelerationMillistandardGravity:
		return a.MillistandardGravity()
	default:
		return math.NaN()
	}
}

func (a *Acceleration) convertFromBase(toUnit AccelerationUnits) float64 {
    value := a.value
	switch toUnit { 
	case AccelerationMeterPerSecondSquared:
		return (value) 
	case AccelerationInchPerSecondSquared:
		return (value / 0.0254) 
	case AccelerationFootPerSecondSquared:
		return (value / 0.304800) 
	case AccelerationKnotPerSecond:
		return (value / 0.5144444444444) 
	case AccelerationKnotPerMinute:
		return (value / 0.5144444444444 * 60) 
	case AccelerationKnotPerHour:
		return (value / 0.5144444444444 * 3600) 
	case AccelerationStandardGravity:
		return (value / 9.80665) 
	case AccelerationNanometerPerSecondSquared:
		return ((value) / 1e-09) 
	case AccelerationMicrometerPerSecondSquared:
		return ((value) / 1e-06) 
	case AccelerationMillimeterPerSecondSquared:
		return ((value) / 0.001) 
	case AccelerationCentimeterPerSecondSquared:
		return ((value) / 0.01) 
	case AccelerationDecimeterPerSecondSquared:
		return ((value) / 0.1) 
	case AccelerationKilometerPerSecondSquared:
		return ((value) / 1000.0) 
	case AccelerationMillistandardGravity:
		return ((value / 9.80665) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Acceleration) convertToBase(value float64, fromUnit AccelerationUnits) float64 {
	switch fromUnit { 
	case AccelerationMeterPerSecondSquared:
		return (value) 
	case AccelerationInchPerSecondSquared:
		return (value * 0.0254) 
	case AccelerationFootPerSecondSquared:
		return (value * 0.304800) 
	case AccelerationKnotPerSecond:
		return (value * 0.5144444444444) 
	case AccelerationKnotPerMinute:
		return (value * 0.5144444444444 / 60) 
	case AccelerationKnotPerHour:
		return (value * 0.5144444444444 / 3600) 
	case AccelerationStandardGravity:
		return (value * 9.80665) 
	case AccelerationNanometerPerSecondSquared:
		return ((value) * 1e-09) 
	case AccelerationMicrometerPerSecondSquared:
		return ((value) * 1e-06) 
	case AccelerationMillimeterPerSecondSquared:
		return ((value) * 0.001) 
	case AccelerationCentimeterPerSecondSquared:
		return ((value) * 0.01) 
	case AccelerationDecimeterPerSecondSquared:
		return ((value) * 0.1) 
	case AccelerationKilometerPerSecondSquared:
		return ((value) * 1000.0) 
	case AccelerationMillistandardGravity:
		return ((value * 9.80665) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Acceleration in the default unit (MeterPerSecondSquared),
// formatted to two decimal places.
func (a Acceleration) String() string {
	return a.ToString(AccelerationMeterPerSecondSquared, 2)
}

// ToString formats the Acceleration value as a string with the specified unit and fractional digits.
// It converts the Acceleration to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Acceleration value will be converted (e.g., MeterPerSecondSquared))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Acceleration with the unit abbreviation.
func (a *Acceleration) ToString(unit AccelerationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAccelerationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAccelerationAbbreviation(unit))
}

// Equals checks if the given Acceleration is equal to the current Acceleration.
//
// Parameters:
//    other: The Acceleration to compare against.
//
// Returns:
//    bool: Returns true if both Acceleration are equal, false otherwise.
func (a *Acceleration) Equals(other *Acceleration) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Acceleration with another Acceleration.
// It returns -1 if the current Acceleration is less than the other Acceleration, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Acceleration to compare against.
//
// Returns:
//    int: -1 if the current Acceleration is less, 1 if greater, and 0 if equal.
func (a *Acceleration) CompareTo(other *Acceleration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Acceleration to the current Acceleration and returns the result.
// The result is a new Acceleration instance with the sum of the values.
//
// Parameters:
//    other: The Acceleration to add to the current Acceleration.
//
// Returns:
//    *Acceleration: A new Acceleration instance representing the sum of both Acceleration.
func (a *Acceleration) Add(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Acceleration from the current Acceleration and returns the result.
// The result is a new Acceleration instance with the difference of the values.
//
// Parameters:
//    other: The Acceleration to subtract from the current Acceleration.
//
// Returns:
//    *Acceleration: A new Acceleration instance representing the difference of both Acceleration.
func (a *Acceleration) Subtract(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Acceleration by the given Acceleration and returns the result.
// The result is a new Acceleration instance with the product of the values.
//
// Parameters:
//    other: The Acceleration to multiply with the current Acceleration.
//
// Returns:
//    *Acceleration: A new Acceleration instance representing the product of both Acceleration.
func (a *Acceleration) Multiply(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value * other.BaseValue()}
}

// Divide divides the current Acceleration by the given Acceleration and returns the result.
// The result is a new Acceleration instance with the quotient of the values.
//
// Parameters:
//    other: The Acceleration to divide the current Acceleration by.
//
// Returns:
//    *Acceleration: A new Acceleration instance representing the quotient of both Acceleration.
func (a *Acceleration) Divide(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value / other.BaseValue()}
}

// GetAccelerationAbbreviation gets the unit abbreviation.
func GetAccelerationAbbreviation(unit AccelerationUnits) string {
	switch unit { 
	case AccelerationMeterPerSecondSquared:
		return "m/s²" 
	case AccelerationInchPerSecondSquared:
		return "in/s²" 
	case AccelerationFootPerSecondSquared:
		return "ft/s²" 
	case AccelerationKnotPerSecond:
		return "kn/s" 
	case AccelerationKnotPerMinute:
		return "kn/min" 
	case AccelerationKnotPerHour:
		return "kn/h" 
	case AccelerationStandardGravity:
		return "g" 
	case AccelerationNanometerPerSecondSquared:
		return "nm/s²" 
	case AccelerationMicrometerPerSecondSquared:
		return "μm/s²" 
	case AccelerationMillimeterPerSecondSquared:
		return "mm/s²" 
	case AccelerationCentimeterPerSecondSquared:
		return "cm/s²" 
	case AccelerationDecimeterPerSecondSquared:
		return "dm/s²" 
	case AccelerationKilometerPerSecondSquared:
		return "km/s²" 
	case AccelerationMillistandardGravity:
		return "mg" 
	default:
		return ""
	}
}