package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureGradientUnits enumeration
type TemperatureGradientUnits string

const (
    
        // 
        TemperatureGradientKelvinPerMeter TemperatureGradientUnits = "KelvinPerMeter"
        // 
        TemperatureGradientDegreeCelsiusPerMeter TemperatureGradientUnits = "DegreeCelsiusPerMeter"
        // 
        TemperatureGradientDegreeFahrenheitPerFoot TemperatureGradientUnits = "DegreeFahrenheitPerFoot"
        // 
        TemperatureGradientDegreeCelsiusPerKilometer TemperatureGradientUnits = "DegreeCelsiusPerKilometer"
)

// TemperatureGradientDto represents an TemperatureGradient
type TemperatureGradientDto struct {
	Value float64
	Unit  TemperatureGradientUnits
}

// TemperatureGradientDtoFactory struct to group related functions
type TemperatureGradientDtoFactory struct{}

func (udf TemperatureGradientDtoFactory) FromJSON(data []byte) (*TemperatureGradientDto, error) {
	a := TemperatureGradientDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TemperatureGradientDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// TemperatureGradient struct
type TemperatureGradient struct {
	value       float64
    
    kelvins_per_meterLazy *float64 
    degrees_celcius_per_meterLazy *float64 
    degrees_fahrenheit_per_footLazy *float64 
    degrees_celcius_per_kilometerLazy *float64 
}

// TemperatureGradientFactory struct to group related functions
type TemperatureGradientFactory struct{}

func (uf TemperatureGradientFactory) CreateTemperatureGradient(value float64, unit TemperatureGradientUnits) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, unit)
}

func (uf TemperatureGradientFactory) FromDto(dto TemperatureGradientDto) (*TemperatureGradient, error) {
	return newTemperatureGradient(dto.Value, dto.Unit)
}

func (uf TemperatureGradientFactory) FromDtoJSON(data []byte) (*TemperatureGradient, error) {
	unitDto, err := TemperatureGradientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TemperatureGradientFactory{}.FromDto(*unitDto)
}


// FromKelvinPerMeter creates a new TemperatureGradient instance from KelvinPerMeter.
func (uf TemperatureGradientFactory) FromKelvinsPerMeter(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientKelvinPerMeter)
}

// FromDegreeCelsiusPerMeter creates a new TemperatureGradient instance from DegreeCelsiusPerMeter.
func (uf TemperatureGradientFactory) FromDegreesCelciusPerMeter(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeCelsiusPerMeter)
}

// FromDegreeFahrenheitPerFoot creates a new TemperatureGradient instance from DegreeFahrenheitPerFoot.
func (uf TemperatureGradientFactory) FromDegreesFahrenheitPerFoot(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeFahrenheitPerFoot)
}

// FromDegreeCelsiusPerKilometer creates a new TemperatureGradient instance from DegreeCelsiusPerKilometer.
func (uf TemperatureGradientFactory) FromDegreesCelciusPerKilometer(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeCelsiusPerKilometer)
}




// newTemperatureGradient creates a new TemperatureGradient.
func newTemperatureGradient(value float64, fromUnit TemperatureGradientUnits) (*TemperatureGradient, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TemperatureGradient{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureGradient in KelvinPerMeter.
func (a *TemperatureGradient) BaseValue() float64 {
	return a.value
}


// KelvinPerMeter returns the value in KelvinPerMeter.
func (a *TemperatureGradient) KelvinsPerMeter() float64 {
	if a.kelvins_per_meterLazy != nil {
		return *a.kelvins_per_meterLazy
	}
	kelvins_per_meter := a.convertFromBase(TemperatureGradientKelvinPerMeter)
	a.kelvins_per_meterLazy = &kelvins_per_meter
	return kelvins_per_meter
}

// DegreeCelsiusPerMeter returns the value in DegreeCelsiusPerMeter.
func (a *TemperatureGradient) DegreesCelciusPerMeter() float64 {
	if a.degrees_celcius_per_meterLazy != nil {
		return *a.degrees_celcius_per_meterLazy
	}
	degrees_celcius_per_meter := a.convertFromBase(TemperatureGradientDegreeCelsiusPerMeter)
	a.degrees_celcius_per_meterLazy = &degrees_celcius_per_meter
	return degrees_celcius_per_meter
}

// DegreeFahrenheitPerFoot returns the value in DegreeFahrenheitPerFoot.
func (a *TemperatureGradient) DegreesFahrenheitPerFoot() float64 {
	if a.degrees_fahrenheit_per_footLazy != nil {
		return *a.degrees_fahrenheit_per_footLazy
	}
	degrees_fahrenheit_per_foot := a.convertFromBase(TemperatureGradientDegreeFahrenheitPerFoot)
	a.degrees_fahrenheit_per_footLazy = &degrees_fahrenheit_per_foot
	return degrees_fahrenheit_per_foot
}

// DegreeCelsiusPerKilometer returns the value in DegreeCelsiusPerKilometer.
func (a *TemperatureGradient) DegreesCelciusPerKilometer() float64 {
	if a.degrees_celcius_per_kilometerLazy != nil {
		return *a.degrees_celcius_per_kilometerLazy
	}
	degrees_celcius_per_kilometer := a.convertFromBase(TemperatureGradientDegreeCelsiusPerKilometer)
	a.degrees_celcius_per_kilometerLazy = &degrees_celcius_per_kilometer
	return degrees_celcius_per_kilometer
}


// ToDto creates an TemperatureGradientDto representation.
func (a *TemperatureGradient) ToDto(holdInUnit *TemperatureGradientUnits) TemperatureGradientDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureGradientKelvinPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureGradientDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an TemperatureGradientDto representation.
func (a *TemperatureGradient) ToDtoJSON(holdInUnit *TemperatureGradientUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts TemperatureGradient to a specific unit value.
func (a *TemperatureGradient) Convert(toUnit TemperatureGradientUnits) float64 {
	switch toUnit { 
    case TemperatureGradientKelvinPerMeter:
		return a.KelvinsPerMeter()
    case TemperatureGradientDegreeCelsiusPerMeter:
		return a.DegreesCelciusPerMeter()
    case TemperatureGradientDegreeFahrenheitPerFoot:
		return a.DegreesFahrenheitPerFoot()
    case TemperatureGradientDegreeCelsiusPerKilometer:
		return a.DegreesCelciusPerKilometer()
	default:
		return 0
	}
}

func (a *TemperatureGradient) convertFromBase(toUnit TemperatureGradientUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureGradientKelvinPerMeter:
		return (value) 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return (value) 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return ((value * 0.3048) * 9 / 5) 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return (value * 1e3) 
	default:
		return math.NaN()
	}
}

func (a *TemperatureGradient) convertToBase(value float64, fromUnit TemperatureGradientUnits) float64 {
	switch fromUnit { 
	case TemperatureGradientKelvinPerMeter:
		return (value) 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return (value) 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return ((value / 0.3048) * 5 / 9) 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return (value / 1e3) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a TemperatureGradient) String() string {
	return a.ToString(TemperatureGradientKelvinPerMeter, 2)
}

// ToString formats the TemperatureGradient to string.
// fractionalDigits -1 for not mention
func (a *TemperatureGradient) ToString(unit TemperatureGradientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TemperatureGradient) getUnitAbbreviation(unit TemperatureGradientUnits) string {
	switch unit { 
	case TemperatureGradientKelvinPerMeter:
		return "∆°K/m" 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return "∆°C/m" 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return "∆°F/ft" 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return "∆°C/km" 
	default:
		return ""
	}
}

// Check if the given TemperatureGradient are equals to the current TemperatureGradient
func (a *TemperatureGradient) Equals(other *TemperatureGradient) bool {
	return a.value == other.BaseValue()
}

// Check if the given TemperatureGradient are equals to the current TemperatureGradient
func (a *TemperatureGradient) CompareTo(other *TemperatureGradient) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given TemperatureGradient to the current TemperatureGradient.
func (a *TemperatureGradient) Add(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value + other.BaseValue()}
}

// Subtract the given TemperatureGradient to the current TemperatureGradient.
func (a *TemperatureGradient) Subtract(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value - other.BaseValue()}
}

// Multiply the given TemperatureGradient to the current TemperatureGradient.
func (a *TemperatureGradient) Multiply(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value * other.BaseValue()}
}

// Divide the given TemperatureGradient to the current TemperatureGradient.
func (a *TemperatureGradient) Divide(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value / other.BaseValue()}
}