package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AccelerationUnits enumeration
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

// AccelerationDto represents an Acceleration
type AccelerationDto struct {
	Value float64
	Unit  AccelerationUnits
}

// AccelerationDtoFactory struct to group related functions
type AccelerationDtoFactory struct{}

func (udf AccelerationDtoFactory) FromJSON(data []byte) (*AccelerationDto, error) {
	a := AccelerationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AccelerationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Acceleration struct
type Acceleration struct {
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

// AccelerationFactory struct to group related functions
type AccelerationFactory struct{}

func (uf AccelerationFactory) CreateAcceleration(value float64, unit AccelerationUnits) (*Acceleration, error) {
	return newAcceleration(value, unit)
}

func (uf AccelerationFactory) FromDto(dto AccelerationDto) (*Acceleration, error) {
	return newAcceleration(dto.Value, dto.Unit)
}

func (uf AccelerationFactory) FromDtoJSON(data []byte) (*Acceleration, error) {
	unitDto, err := AccelerationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AccelerationFactory{}.FromDto(*unitDto)
}


// FromMeterPerSecondSquared creates a new Acceleration instance from MeterPerSecondSquared.
func (uf AccelerationFactory) FromMetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMeterPerSecondSquared)
}

// FromInchPerSecondSquared creates a new Acceleration instance from InchPerSecondSquared.
func (uf AccelerationFactory) FromInchesPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationInchPerSecondSquared)
}

// FromFootPerSecondSquared creates a new Acceleration instance from FootPerSecondSquared.
func (uf AccelerationFactory) FromFeetPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationFootPerSecondSquared)
}

// FromKnotPerSecond creates a new Acceleration instance from KnotPerSecond.
func (uf AccelerationFactory) FromKnotsPerSecond(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerSecond)
}

// FromKnotPerMinute creates a new Acceleration instance from KnotPerMinute.
func (uf AccelerationFactory) FromKnotsPerMinute(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerMinute)
}

// FromKnotPerHour creates a new Acceleration instance from KnotPerHour.
func (uf AccelerationFactory) FromKnotsPerHour(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKnotPerHour)
}

// FromStandardGravity creates a new Acceleration instance from StandardGravity.
func (uf AccelerationFactory) FromStandardGravity(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationStandardGravity)
}

// FromNanometerPerSecondSquared creates a new Acceleration instance from NanometerPerSecondSquared.
func (uf AccelerationFactory) FromNanometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationNanometerPerSecondSquared)
}

// FromMicrometerPerSecondSquared creates a new Acceleration instance from MicrometerPerSecondSquared.
func (uf AccelerationFactory) FromMicrometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMicrometerPerSecondSquared)
}

// FromMillimeterPerSecondSquared creates a new Acceleration instance from MillimeterPerSecondSquared.
func (uf AccelerationFactory) FromMillimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationMillimeterPerSecondSquared)
}

// FromCentimeterPerSecondSquared creates a new Acceleration instance from CentimeterPerSecondSquared.
func (uf AccelerationFactory) FromCentimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationCentimeterPerSecondSquared)
}

// FromDecimeterPerSecondSquared creates a new Acceleration instance from DecimeterPerSecondSquared.
func (uf AccelerationFactory) FromDecimetersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationDecimeterPerSecondSquared)
}

// FromKilometerPerSecondSquared creates a new Acceleration instance from KilometerPerSecondSquared.
func (uf AccelerationFactory) FromKilometersPerSecondSquared(value float64) (*Acceleration, error) {
	return newAcceleration(value, AccelerationKilometerPerSecondSquared)
}

// FromMillistandardGravity creates a new Acceleration instance from MillistandardGravity.
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

// BaseValue returns the base value of Acceleration in MeterPerSecondSquared.
func (a *Acceleration) BaseValue() float64 {
	return a.value
}


// MeterPerSecondSquared returns the value in MeterPerSecondSquared.
func (a *Acceleration) MetersPerSecondSquared() float64 {
	if a.meters_per_second_squaredLazy != nil {
		return *a.meters_per_second_squaredLazy
	}
	meters_per_second_squared := a.convertFromBase(AccelerationMeterPerSecondSquared)
	a.meters_per_second_squaredLazy = &meters_per_second_squared
	return meters_per_second_squared
}

// InchPerSecondSquared returns the value in InchPerSecondSquared.
func (a *Acceleration) InchesPerSecondSquared() float64 {
	if a.inches_per_second_squaredLazy != nil {
		return *a.inches_per_second_squaredLazy
	}
	inches_per_second_squared := a.convertFromBase(AccelerationInchPerSecondSquared)
	a.inches_per_second_squaredLazy = &inches_per_second_squared
	return inches_per_second_squared
}

// FootPerSecondSquared returns the value in FootPerSecondSquared.
func (a *Acceleration) FeetPerSecondSquared() float64 {
	if a.feet_per_second_squaredLazy != nil {
		return *a.feet_per_second_squaredLazy
	}
	feet_per_second_squared := a.convertFromBase(AccelerationFootPerSecondSquared)
	a.feet_per_second_squaredLazy = &feet_per_second_squared
	return feet_per_second_squared
}

// KnotPerSecond returns the value in KnotPerSecond.
func (a *Acceleration) KnotsPerSecond() float64 {
	if a.knots_per_secondLazy != nil {
		return *a.knots_per_secondLazy
	}
	knots_per_second := a.convertFromBase(AccelerationKnotPerSecond)
	a.knots_per_secondLazy = &knots_per_second
	return knots_per_second
}

// KnotPerMinute returns the value in KnotPerMinute.
func (a *Acceleration) KnotsPerMinute() float64 {
	if a.knots_per_minuteLazy != nil {
		return *a.knots_per_minuteLazy
	}
	knots_per_minute := a.convertFromBase(AccelerationKnotPerMinute)
	a.knots_per_minuteLazy = &knots_per_minute
	return knots_per_minute
}

// KnotPerHour returns the value in KnotPerHour.
func (a *Acceleration) KnotsPerHour() float64 {
	if a.knots_per_hourLazy != nil {
		return *a.knots_per_hourLazy
	}
	knots_per_hour := a.convertFromBase(AccelerationKnotPerHour)
	a.knots_per_hourLazy = &knots_per_hour
	return knots_per_hour
}

// StandardGravity returns the value in StandardGravity.
func (a *Acceleration) StandardGravity() float64 {
	if a.standard_gravityLazy != nil {
		return *a.standard_gravityLazy
	}
	standard_gravity := a.convertFromBase(AccelerationStandardGravity)
	a.standard_gravityLazy = &standard_gravity
	return standard_gravity
}

// NanometerPerSecondSquared returns the value in NanometerPerSecondSquared.
func (a *Acceleration) NanometersPerSecondSquared() float64 {
	if a.nanometers_per_second_squaredLazy != nil {
		return *a.nanometers_per_second_squaredLazy
	}
	nanometers_per_second_squared := a.convertFromBase(AccelerationNanometerPerSecondSquared)
	a.nanometers_per_second_squaredLazy = &nanometers_per_second_squared
	return nanometers_per_second_squared
}

// MicrometerPerSecondSquared returns the value in MicrometerPerSecondSquared.
func (a *Acceleration) MicrometersPerSecondSquared() float64 {
	if a.micrometers_per_second_squaredLazy != nil {
		return *a.micrometers_per_second_squaredLazy
	}
	micrometers_per_second_squared := a.convertFromBase(AccelerationMicrometerPerSecondSquared)
	a.micrometers_per_second_squaredLazy = &micrometers_per_second_squared
	return micrometers_per_second_squared
}

// MillimeterPerSecondSquared returns the value in MillimeterPerSecondSquared.
func (a *Acceleration) MillimetersPerSecondSquared() float64 {
	if a.millimeters_per_second_squaredLazy != nil {
		return *a.millimeters_per_second_squaredLazy
	}
	millimeters_per_second_squared := a.convertFromBase(AccelerationMillimeterPerSecondSquared)
	a.millimeters_per_second_squaredLazy = &millimeters_per_second_squared
	return millimeters_per_second_squared
}

// CentimeterPerSecondSquared returns the value in CentimeterPerSecondSquared.
func (a *Acceleration) CentimetersPerSecondSquared() float64 {
	if a.centimeters_per_second_squaredLazy != nil {
		return *a.centimeters_per_second_squaredLazy
	}
	centimeters_per_second_squared := a.convertFromBase(AccelerationCentimeterPerSecondSquared)
	a.centimeters_per_second_squaredLazy = &centimeters_per_second_squared
	return centimeters_per_second_squared
}

// DecimeterPerSecondSquared returns the value in DecimeterPerSecondSquared.
func (a *Acceleration) DecimetersPerSecondSquared() float64 {
	if a.decimeters_per_second_squaredLazy != nil {
		return *a.decimeters_per_second_squaredLazy
	}
	decimeters_per_second_squared := a.convertFromBase(AccelerationDecimeterPerSecondSquared)
	a.decimeters_per_second_squaredLazy = &decimeters_per_second_squared
	return decimeters_per_second_squared
}

// KilometerPerSecondSquared returns the value in KilometerPerSecondSquared.
func (a *Acceleration) KilometersPerSecondSquared() float64 {
	if a.kilometers_per_second_squaredLazy != nil {
		return *a.kilometers_per_second_squaredLazy
	}
	kilometers_per_second_squared := a.convertFromBase(AccelerationKilometerPerSecondSquared)
	a.kilometers_per_second_squaredLazy = &kilometers_per_second_squared
	return kilometers_per_second_squared
}

// MillistandardGravity returns the value in MillistandardGravity.
func (a *Acceleration) MillistandardGravity() float64 {
	if a.millistandard_gravityLazy != nil {
		return *a.millistandard_gravityLazy
	}
	millistandard_gravity := a.convertFromBase(AccelerationMillistandardGravity)
	a.millistandard_gravityLazy = &millistandard_gravity
	return millistandard_gravity
}


// ToDto creates an AccelerationDto representation.
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

// ToDtoJSON creates an AccelerationDto representation.
func (a *Acceleration) ToDtoJSON(holdInUnit *AccelerationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Acceleration to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Acceleration) String() string {
	return a.ToString(AccelerationMeterPerSecondSquared, 2)
}

// ToString formats the Acceleration to string.
// fractionalDigits -1 for not mention
func (a *Acceleration) ToString(unit AccelerationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Acceleration) getUnitAbbreviation(unit AccelerationUnits) string {
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

// Check if the given Acceleration are equals to the current Acceleration
func (a *Acceleration) Equals(other *Acceleration) bool {
	return a.value == other.BaseValue()
}

// Check if the given Acceleration are equals to the current Acceleration
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

// Add the given Acceleration to the current Acceleration.
func (a *Acceleration) Add(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value + other.BaseValue()}
}

// Subtract the given Acceleration to the current Acceleration.
func (a *Acceleration) Subtract(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value - other.BaseValue()}
}

// Multiply the given Acceleration to the current Acceleration.
func (a *Acceleration) Multiply(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value * other.BaseValue()}
}

// Divide the given Acceleration to the current Acceleration.
func (a *Acceleration) Divide(other *Acceleration) *Acceleration {
	return &Acceleration{value: a.value / other.BaseValue()}
}