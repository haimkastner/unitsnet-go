// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// JerkUnits enumeration
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

// JerkDto represents an Jerk
type JerkDto struct {
	Value float64
	Unit  JerkUnits
}

// JerkDtoFactory struct to group related functions
type JerkDtoFactory struct{}

func (udf JerkDtoFactory) FromJSON(data []byte) (*JerkDto, error) {
	a := JerkDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a JerkDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Jerk struct
type Jerk struct {
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

// JerkFactory struct to group related functions
type JerkFactory struct{}

func (uf JerkFactory) CreateJerk(value float64, unit JerkUnits) (*Jerk, error) {
	return newJerk(value, unit)
}

func (uf JerkFactory) FromDto(dto JerkDto) (*Jerk, error) {
	return newJerk(dto.Value, dto.Unit)
}

func (uf JerkFactory) FromDtoJSON(data []byte) (*Jerk, error) {
	unitDto, err := JerkDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return JerkFactory{}.FromDto(*unitDto)
}


// FromMeterPerSecondCubed creates a new Jerk instance from MeterPerSecondCubed.
func (uf JerkFactory) FromMetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMeterPerSecondCubed)
}

// FromInchPerSecondCubed creates a new Jerk instance from InchPerSecondCubed.
func (uf JerkFactory) FromInchesPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkInchPerSecondCubed)
}

// FromFootPerSecondCubed creates a new Jerk instance from FootPerSecondCubed.
func (uf JerkFactory) FromFeetPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkFootPerSecondCubed)
}

// FromStandardGravitiesPerSecond creates a new Jerk instance from StandardGravitiesPerSecond.
func (uf JerkFactory) FromStandardGravitiesPerSecond(value float64) (*Jerk, error) {
	return newJerk(value, JerkStandardGravitiesPerSecond)
}

// FromNanometerPerSecondCubed creates a new Jerk instance from NanometerPerSecondCubed.
func (uf JerkFactory) FromNanometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkNanometerPerSecondCubed)
}

// FromMicrometerPerSecondCubed creates a new Jerk instance from MicrometerPerSecondCubed.
func (uf JerkFactory) FromMicrometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMicrometerPerSecondCubed)
}

// FromMillimeterPerSecondCubed creates a new Jerk instance from MillimeterPerSecondCubed.
func (uf JerkFactory) FromMillimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkMillimeterPerSecondCubed)
}

// FromCentimeterPerSecondCubed creates a new Jerk instance from CentimeterPerSecondCubed.
func (uf JerkFactory) FromCentimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkCentimeterPerSecondCubed)
}

// FromDecimeterPerSecondCubed creates a new Jerk instance from DecimeterPerSecondCubed.
func (uf JerkFactory) FromDecimetersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkDecimeterPerSecondCubed)
}

// FromKilometerPerSecondCubed creates a new Jerk instance from KilometerPerSecondCubed.
func (uf JerkFactory) FromKilometersPerSecondCubed(value float64) (*Jerk, error) {
	return newJerk(value, JerkKilometerPerSecondCubed)
}

// FromMillistandardGravitiesPerSecond creates a new Jerk instance from MillistandardGravitiesPerSecond.
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

// BaseValue returns the base value of Jerk in MeterPerSecondCubed.
func (a *Jerk) BaseValue() float64 {
	return a.value
}


// MeterPerSecondCubed returns the value in MeterPerSecondCubed.
func (a *Jerk) MetersPerSecondCubed() float64 {
	if a.meters_per_second_cubedLazy != nil {
		return *a.meters_per_second_cubedLazy
	}
	meters_per_second_cubed := a.convertFromBase(JerkMeterPerSecondCubed)
	a.meters_per_second_cubedLazy = &meters_per_second_cubed
	return meters_per_second_cubed
}

// InchPerSecondCubed returns the value in InchPerSecondCubed.
func (a *Jerk) InchesPerSecondCubed() float64 {
	if a.inches_per_second_cubedLazy != nil {
		return *a.inches_per_second_cubedLazy
	}
	inches_per_second_cubed := a.convertFromBase(JerkInchPerSecondCubed)
	a.inches_per_second_cubedLazy = &inches_per_second_cubed
	return inches_per_second_cubed
}

// FootPerSecondCubed returns the value in FootPerSecondCubed.
func (a *Jerk) FeetPerSecondCubed() float64 {
	if a.feet_per_second_cubedLazy != nil {
		return *a.feet_per_second_cubedLazy
	}
	feet_per_second_cubed := a.convertFromBase(JerkFootPerSecondCubed)
	a.feet_per_second_cubedLazy = &feet_per_second_cubed
	return feet_per_second_cubed
}

// StandardGravitiesPerSecond returns the value in StandardGravitiesPerSecond.
func (a *Jerk) StandardGravitiesPerSecond() float64 {
	if a.standard_gravities_per_secondLazy != nil {
		return *a.standard_gravities_per_secondLazy
	}
	standard_gravities_per_second := a.convertFromBase(JerkStandardGravitiesPerSecond)
	a.standard_gravities_per_secondLazy = &standard_gravities_per_second
	return standard_gravities_per_second
}

// NanometerPerSecondCubed returns the value in NanometerPerSecondCubed.
func (a *Jerk) NanometersPerSecondCubed() float64 {
	if a.nanometers_per_second_cubedLazy != nil {
		return *a.nanometers_per_second_cubedLazy
	}
	nanometers_per_second_cubed := a.convertFromBase(JerkNanometerPerSecondCubed)
	a.nanometers_per_second_cubedLazy = &nanometers_per_second_cubed
	return nanometers_per_second_cubed
}

// MicrometerPerSecondCubed returns the value in MicrometerPerSecondCubed.
func (a *Jerk) MicrometersPerSecondCubed() float64 {
	if a.micrometers_per_second_cubedLazy != nil {
		return *a.micrometers_per_second_cubedLazy
	}
	micrometers_per_second_cubed := a.convertFromBase(JerkMicrometerPerSecondCubed)
	a.micrometers_per_second_cubedLazy = &micrometers_per_second_cubed
	return micrometers_per_second_cubed
}

// MillimeterPerSecondCubed returns the value in MillimeterPerSecondCubed.
func (a *Jerk) MillimetersPerSecondCubed() float64 {
	if a.millimeters_per_second_cubedLazy != nil {
		return *a.millimeters_per_second_cubedLazy
	}
	millimeters_per_second_cubed := a.convertFromBase(JerkMillimeterPerSecondCubed)
	a.millimeters_per_second_cubedLazy = &millimeters_per_second_cubed
	return millimeters_per_second_cubed
}

// CentimeterPerSecondCubed returns the value in CentimeterPerSecondCubed.
func (a *Jerk) CentimetersPerSecondCubed() float64 {
	if a.centimeters_per_second_cubedLazy != nil {
		return *a.centimeters_per_second_cubedLazy
	}
	centimeters_per_second_cubed := a.convertFromBase(JerkCentimeterPerSecondCubed)
	a.centimeters_per_second_cubedLazy = &centimeters_per_second_cubed
	return centimeters_per_second_cubed
}

// DecimeterPerSecondCubed returns the value in DecimeterPerSecondCubed.
func (a *Jerk) DecimetersPerSecondCubed() float64 {
	if a.decimeters_per_second_cubedLazy != nil {
		return *a.decimeters_per_second_cubedLazy
	}
	decimeters_per_second_cubed := a.convertFromBase(JerkDecimeterPerSecondCubed)
	a.decimeters_per_second_cubedLazy = &decimeters_per_second_cubed
	return decimeters_per_second_cubed
}

// KilometerPerSecondCubed returns the value in KilometerPerSecondCubed.
func (a *Jerk) KilometersPerSecondCubed() float64 {
	if a.kilometers_per_second_cubedLazy != nil {
		return *a.kilometers_per_second_cubedLazy
	}
	kilometers_per_second_cubed := a.convertFromBase(JerkKilometerPerSecondCubed)
	a.kilometers_per_second_cubedLazy = &kilometers_per_second_cubed
	return kilometers_per_second_cubed
}

// MillistandardGravitiesPerSecond returns the value in MillistandardGravitiesPerSecond.
func (a *Jerk) MillistandardGravitiesPerSecond() float64 {
	if a.millistandard_gravities_per_secondLazy != nil {
		return *a.millistandard_gravities_per_secondLazy
	}
	millistandard_gravities_per_second := a.convertFromBase(JerkMillistandardGravitiesPerSecond)
	a.millistandard_gravities_per_secondLazy = &millistandard_gravities_per_second
	return millistandard_gravities_per_second
}


// ToDto creates an JerkDto representation.
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

// ToDtoJSON creates an JerkDto representation.
func (a *Jerk) ToDtoJSON(holdInUnit *JerkUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Jerk to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Jerk) String() string {
	return a.ToString(JerkMeterPerSecondCubed, 2)
}

// ToString formats the Jerk to string.
// fractionalDigits -1 for not mention
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

// Check if the given Jerk are equals to the current Jerk
func (a *Jerk) Equals(other *Jerk) bool {
	return a.value == other.BaseValue()
}

// Check if the given Jerk are equals to the current Jerk
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

// Add the given Jerk to the current Jerk.
func (a *Jerk) Add(other *Jerk) *Jerk {
	return &Jerk{value: a.value + other.BaseValue()}
}

// Subtract the given Jerk to the current Jerk.
func (a *Jerk) Subtract(other *Jerk) *Jerk {
	return &Jerk{value: a.value - other.BaseValue()}
}

// Multiply the given Jerk to the current Jerk.
func (a *Jerk) Multiply(other *Jerk) *Jerk {
	return &Jerk{value: a.value * other.BaseValue()}
}

// Divide the given Jerk to the current Jerk.
func (a *Jerk) Divide(other *Jerk) *Jerk {
	return &Jerk{value: a.value / other.BaseValue()}
}