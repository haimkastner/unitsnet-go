// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureUnits enumeration
type TemperatureUnits string

const (
    
        // 
        TemperatureKelvin TemperatureUnits = "Kelvin"
        // 
        TemperatureDegreeCelsius TemperatureUnits = "DegreeCelsius"
        // 
        TemperatureMillidegreeCelsius TemperatureUnits = "MillidegreeCelsius"
        // 
        TemperatureDegreeDelisle TemperatureUnits = "DegreeDelisle"
        // 
        TemperatureDegreeFahrenheit TemperatureUnits = "DegreeFahrenheit"
        // 
        TemperatureDegreeNewton TemperatureUnits = "DegreeNewton"
        // 
        TemperatureDegreeRankine TemperatureUnits = "DegreeRankine"
        // 
        TemperatureDegreeReaumur TemperatureUnits = "DegreeReaumur"
        // 
        TemperatureDegreeRoemer TemperatureUnits = "DegreeRoemer"
        // 
        TemperatureSolarTemperature TemperatureUnits = "SolarTemperature"
)

// TemperatureDto represents an Temperature
type TemperatureDto struct {
	Value float64
	Unit  TemperatureUnits
}

// TemperatureDtoFactory struct to group related functions
type TemperatureDtoFactory struct{}

func (udf TemperatureDtoFactory) FromJSON(data []byte) (*TemperatureDto, error) {
	a := TemperatureDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TemperatureDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Temperature struct
type Temperature struct {
	value       float64
    
    kelvinsLazy *float64 
    degrees_celsiusLazy *float64 
    millidegrees_celsiusLazy *float64 
    degrees_delisleLazy *float64 
    degrees_fahrenheitLazy *float64 
    degrees_newtonLazy *float64 
    degrees_rankineLazy *float64 
    degrees_reaumurLazy *float64 
    degrees_roemerLazy *float64 
    solar_temperaturesLazy *float64 
}

// TemperatureFactory struct to group related functions
type TemperatureFactory struct{}

func (uf TemperatureFactory) CreateTemperature(value float64, unit TemperatureUnits) (*Temperature, error) {
	return newTemperature(value, unit)
}

func (uf TemperatureFactory) FromDto(dto TemperatureDto) (*Temperature, error) {
	return newTemperature(dto.Value, dto.Unit)
}

func (uf TemperatureFactory) FromDtoJSON(data []byte) (*Temperature, error) {
	unitDto, err := TemperatureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TemperatureFactory{}.FromDto(*unitDto)
}


// FromKelvin creates a new Temperature instance from Kelvin.
func (uf TemperatureFactory) FromKelvins(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureKelvin)
}

// FromDegreeCelsius creates a new Temperature instance from DegreeCelsius.
func (uf TemperatureFactory) FromDegreesCelsius(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeCelsius)
}

// FromMillidegreeCelsius creates a new Temperature instance from MillidegreeCelsius.
func (uf TemperatureFactory) FromMillidegreesCelsius(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureMillidegreeCelsius)
}

// FromDegreeDelisle creates a new Temperature instance from DegreeDelisle.
func (uf TemperatureFactory) FromDegreesDelisle(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeDelisle)
}

// FromDegreeFahrenheit creates a new Temperature instance from DegreeFahrenheit.
func (uf TemperatureFactory) FromDegreesFahrenheit(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeFahrenheit)
}

// FromDegreeNewton creates a new Temperature instance from DegreeNewton.
func (uf TemperatureFactory) FromDegreesNewton(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeNewton)
}

// FromDegreeRankine creates a new Temperature instance from DegreeRankine.
func (uf TemperatureFactory) FromDegreesRankine(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeRankine)
}

// FromDegreeReaumur creates a new Temperature instance from DegreeReaumur.
func (uf TemperatureFactory) FromDegreesReaumur(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeReaumur)
}

// FromDegreeRoemer creates a new Temperature instance from DegreeRoemer.
func (uf TemperatureFactory) FromDegreesRoemer(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeRoemer)
}

// FromSolarTemperature creates a new Temperature instance from SolarTemperature.
func (uf TemperatureFactory) FromSolarTemperatures(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureSolarTemperature)
}




// newTemperature creates a new Temperature.
func newTemperature(value float64, fromUnit TemperatureUnits) (*Temperature, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Temperature{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Temperature in Kelvin.
func (a *Temperature) BaseValue() float64 {
	return a.value
}


// Kelvin returns the value in Kelvin.
func (a *Temperature) Kelvins() float64 {
	if a.kelvinsLazy != nil {
		return *a.kelvinsLazy
	}
	kelvins := a.convertFromBase(TemperatureKelvin)
	a.kelvinsLazy = &kelvins
	return kelvins
}

// DegreeCelsius returns the value in DegreeCelsius.
func (a *Temperature) DegreesCelsius() float64 {
	if a.degrees_celsiusLazy != nil {
		return *a.degrees_celsiusLazy
	}
	degrees_celsius := a.convertFromBase(TemperatureDegreeCelsius)
	a.degrees_celsiusLazy = &degrees_celsius
	return degrees_celsius
}

// MillidegreeCelsius returns the value in MillidegreeCelsius.
func (a *Temperature) MillidegreesCelsius() float64 {
	if a.millidegrees_celsiusLazy != nil {
		return *a.millidegrees_celsiusLazy
	}
	millidegrees_celsius := a.convertFromBase(TemperatureMillidegreeCelsius)
	a.millidegrees_celsiusLazy = &millidegrees_celsius
	return millidegrees_celsius
}

// DegreeDelisle returns the value in DegreeDelisle.
func (a *Temperature) DegreesDelisle() float64 {
	if a.degrees_delisleLazy != nil {
		return *a.degrees_delisleLazy
	}
	degrees_delisle := a.convertFromBase(TemperatureDegreeDelisle)
	a.degrees_delisleLazy = &degrees_delisle
	return degrees_delisle
}

// DegreeFahrenheit returns the value in DegreeFahrenheit.
func (a *Temperature) DegreesFahrenheit() float64 {
	if a.degrees_fahrenheitLazy != nil {
		return *a.degrees_fahrenheitLazy
	}
	degrees_fahrenheit := a.convertFromBase(TemperatureDegreeFahrenheit)
	a.degrees_fahrenheitLazy = &degrees_fahrenheit
	return degrees_fahrenheit
}

// DegreeNewton returns the value in DegreeNewton.
func (a *Temperature) DegreesNewton() float64 {
	if a.degrees_newtonLazy != nil {
		return *a.degrees_newtonLazy
	}
	degrees_newton := a.convertFromBase(TemperatureDegreeNewton)
	a.degrees_newtonLazy = &degrees_newton
	return degrees_newton
}

// DegreeRankine returns the value in DegreeRankine.
func (a *Temperature) DegreesRankine() float64 {
	if a.degrees_rankineLazy != nil {
		return *a.degrees_rankineLazy
	}
	degrees_rankine := a.convertFromBase(TemperatureDegreeRankine)
	a.degrees_rankineLazy = &degrees_rankine
	return degrees_rankine
}

// DegreeReaumur returns the value in DegreeReaumur.
func (a *Temperature) DegreesReaumur() float64 {
	if a.degrees_reaumurLazy != nil {
		return *a.degrees_reaumurLazy
	}
	degrees_reaumur := a.convertFromBase(TemperatureDegreeReaumur)
	a.degrees_reaumurLazy = &degrees_reaumur
	return degrees_reaumur
}

// DegreeRoemer returns the value in DegreeRoemer.
func (a *Temperature) DegreesRoemer() float64 {
	if a.degrees_roemerLazy != nil {
		return *a.degrees_roemerLazy
	}
	degrees_roemer := a.convertFromBase(TemperatureDegreeRoemer)
	a.degrees_roemerLazy = &degrees_roemer
	return degrees_roemer
}

// SolarTemperature returns the value in SolarTemperature.
func (a *Temperature) SolarTemperatures() float64 {
	if a.solar_temperaturesLazy != nil {
		return *a.solar_temperaturesLazy
	}
	solar_temperatures := a.convertFromBase(TemperatureSolarTemperature)
	a.solar_temperaturesLazy = &solar_temperatures
	return solar_temperatures
}


// ToDto creates an TemperatureDto representation.
func (a *Temperature) ToDto(holdInUnit *TemperatureUnits) TemperatureDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an TemperatureDto representation.
func (a *Temperature) ToDtoJSON(holdInUnit *TemperatureUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Temperature to a specific unit value.
func (a *Temperature) Convert(toUnit TemperatureUnits) float64 {
	switch toUnit { 
    case TemperatureKelvin:
		return a.Kelvins()
    case TemperatureDegreeCelsius:
		return a.DegreesCelsius()
    case TemperatureMillidegreeCelsius:
		return a.MillidegreesCelsius()
    case TemperatureDegreeDelisle:
		return a.DegreesDelisle()
    case TemperatureDegreeFahrenheit:
		return a.DegreesFahrenheit()
    case TemperatureDegreeNewton:
		return a.DegreesNewton()
    case TemperatureDegreeRankine:
		return a.DegreesRankine()
    case TemperatureDegreeReaumur:
		return a.DegreesReaumur()
    case TemperatureDegreeRoemer:
		return a.DegreesRoemer()
    case TemperatureSolarTemperature:
		return a.SolarTemperatures()
	default:
		return 0
	}
}

func (a *Temperature) convertFromBase(toUnit TemperatureUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureKelvin:
		return (value) 
	case TemperatureDegreeCelsius:
		return (value - 273.15) 
	case TemperatureMillidegreeCelsius:
		return ((value - 273.15) * 1000) 
	case TemperatureDegreeDelisle:
		return ((value - 373.15) * -3 / 2) 
	case TemperatureDegreeFahrenheit:
		return ((value - 459.67 * 5 / 9) * 9 / 5) 
	case TemperatureDegreeNewton:
		return ((value - 273.15) * 33 / 100) 
	case TemperatureDegreeRankine:
		return (value * 9 / 5) 
	case TemperatureDegreeReaumur:
		return ((value - 273.15) * 4 / 5) 
	case TemperatureDegreeRoemer:
		return ((value - (273.15 - 7.5 * 40 / 21)) * 21 / 40) 
	case TemperatureSolarTemperature:
		return (value / 5778) 
	default:
		return math.NaN()
	}
}

func (a *Temperature) convertToBase(value float64, fromUnit TemperatureUnits) float64 {
	switch fromUnit { 
	case TemperatureKelvin:
		return (value) 
	case TemperatureDegreeCelsius:
		return (value + 273.15) 
	case TemperatureMillidegreeCelsius:
		return (value / 1000 + 273.15) 
	case TemperatureDegreeDelisle:
		return (value * -2 / 3 + 373.15) 
	case TemperatureDegreeFahrenheit:
		return (value * 5 / 9 + 459.67 * 5 / 9) 
	case TemperatureDegreeNewton:
		return (value * 100 / 33 + 273.15) 
	case TemperatureDegreeRankine:
		return (value * 5 / 9) 
	case TemperatureDegreeReaumur:
		return (value * 5 / 4 + 273.15) 
	case TemperatureDegreeRoemer:
		return (value * 40 / 21 + 273.15 - 7.5 * 40 / 21) 
	case TemperatureSolarTemperature:
		return (value * 5778) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Temperature) String() string {
	return a.ToString(TemperatureKelvin, 2)
}

// ToString formats the Temperature to string.
// fractionalDigits -1 for not mention
func (a *Temperature) ToString(unit TemperatureUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Temperature) getUnitAbbreviation(unit TemperatureUnits) string {
	switch unit { 
	case TemperatureKelvin:
		return "K" 
	case TemperatureDegreeCelsius:
		return "°C" 
	case TemperatureMillidegreeCelsius:
		return "m°C" 
	case TemperatureDegreeDelisle:
		return "°De" 
	case TemperatureDegreeFahrenheit:
		return "°F" 
	case TemperatureDegreeNewton:
		return "°N" 
	case TemperatureDegreeRankine:
		return "°R" 
	case TemperatureDegreeReaumur:
		return "°Ré" 
	case TemperatureDegreeRoemer:
		return "°Rø" 
	case TemperatureSolarTemperature:
		return "T⊙" 
	default:
		return ""
	}
}

// Check if the given Temperature are equals to the current Temperature
func (a *Temperature) Equals(other *Temperature) bool {
	return a.value == other.BaseValue()
}

// Check if the given Temperature are equals to the current Temperature
func (a *Temperature) CompareTo(other *Temperature) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Temperature to the current Temperature.
func (a *Temperature) Add(other *Temperature) *Temperature {
	return &Temperature{value: a.value + other.BaseValue()}
}

// Subtract the given Temperature to the current Temperature.
func (a *Temperature) Subtract(other *Temperature) *Temperature {
	return &Temperature{value: a.value - other.BaseValue()}
}

// Multiply the given Temperature to the current Temperature.
func (a *Temperature) Multiply(other *Temperature) *Temperature {
	return &Temperature{value: a.value * other.BaseValue()}
}

// Divide the given Temperature to the current Temperature.
func (a *Temperature) Divide(other *Temperature) *Temperature {
	return &Temperature{value: a.value / other.BaseValue()}
}