// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureUnits defines various units of Temperature.
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

// TemperatureDto represents a Temperature measurement with a numerical value and its corresponding unit.
type TemperatureDto struct {
    // Value is the numerical representation of the Temperature.
	Value float64
    // Unit specifies the unit of measurement for the Temperature, as defined in the TemperatureUnits enumeration.
	Unit  TemperatureUnits
}

// TemperatureDtoFactory groups methods for creating and serializing TemperatureDto objects.
type TemperatureDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TemperatureDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TemperatureDtoFactory) FromJSON(data []byte) (*TemperatureDto, error) {
	a := TemperatureDto{}

    // Parse JSON into TemperatureDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a TemperatureDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TemperatureDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Temperature represents a measurement in a of Temperature.
//
// A temperature is a numerical measure of hot or cold. Its measurement is by detection of heat radiation or particle velocity or kinetic energy, or by the bulk behavior of a thermometric material. It may be calibrated in any of various temperature scales, Celsius, Fahrenheit, Kelvin, etc. The fundamental physical definition of temperature is provided by thermodynamics.
type Temperature struct {
	// value is the base measurement stored internally.
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

// TemperatureFactory groups methods for creating Temperature instances.
type TemperatureFactory struct{}

// CreateTemperature creates a new Temperature instance from the given value and unit.
func (uf TemperatureFactory) CreateTemperature(value float64, unit TemperatureUnits) (*Temperature, error) {
	return newTemperature(value, unit)
}

// FromDto converts a TemperatureDto to a Temperature instance.
func (uf TemperatureFactory) FromDto(dto TemperatureDto) (*Temperature, error) {
	return newTemperature(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Temperature instance.
func (uf TemperatureFactory) FromDtoJSON(data []byte) (*Temperature, error) {
	unitDto, err := TemperatureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TemperatureDto from JSON: %w", err)
	}
	return TemperatureFactory{}.FromDto(*unitDto)
}


// FromKelvins creates a new Temperature instance from a value in Kelvins.
func (uf TemperatureFactory) FromKelvins(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureKelvin)
}

// FromDegreesCelsius creates a new Temperature instance from a value in DegreesCelsius.
func (uf TemperatureFactory) FromDegreesCelsius(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeCelsius)
}

// FromMillidegreesCelsius creates a new Temperature instance from a value in MillidegreesCelsius.
func (uf TemperatureFactory) FromMillidegreesCelsius(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureMillidegreeCelsius)
}

// FromDegreesDelisle creates a new Temperature instance from a value in DegreesDelisle.
func (uf TemperatureFactory) FromDegreesDelisle(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeDelisle)
}

// FromDegreesFahrenheit creates a new Temperature instance from a value in DegreesFahrenheit.
func (uf TemperatureFactory) FromDegreesFahrenheit(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeFahrenheit)
}

// FromDegreesNewton creates a new Temperature instance from a value in DegreesNewton.
func (uf TemperatureFactory) FromDegreesNewton(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeNewton)
}

// FromDegreesRankine creates a new Temperature instance from a value in DegreesRankine.
func (uf TemperatureFactory) FromDegreesRankine(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeRankine)
}

// FromDegreesReaumur creates a new Temperature instance from a value in DegreesReaumur.
func (uf TemperatureFactory) FromDegreesReaumur(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeReaumur)
}

// FromDegreesRoemer creates a new Temperature instance from a value in DegreesRoemer.
func (uf TemperatureFactory) FromDegreesRoemer(value float64) (*Temperature, error) {
	return newTemperature(value, TemperatureDegreeRoemer)
}

// FromSolarTemperatures creates a new Temperature instance from a value in SolarTemperatures.
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

// BaseValue returns the base value of Temperature in Kelvin unit (the base/default quantity).
func (a *Temperature) BaseValue() float64 {
	return a.value
}


// Kelvins returns the Temperature value in Kelvins.
//
// 
func (a *Temperature) Kelvins() float64 {
	if a.kelvinsLazy != nil {
		return *a.kelvinsLazy
	}
	kelvins := a.convertFromBase(TemperatureKelvin)
	a.kelvinsLazy = &kelvins
	return kelvins
}

// DegreesCelsius returns the Temperature value in DegreesCelsius.
//
// 
func (a *Temperature) DegreesCelsius() float64 {
	if a.degrees_celsiusLazy != nil {
		return *a.degrees_celsiusLazy
	}
	degrees_celsius := a.convertFromBase(TemperatureDegreeCelsius)
	a.degrees_celsiusLazy = &degrees_celsius
	return degrees_celsius
}

// MillidegreesCelsius returns the Temperature value in MillidegreesCelsius.
//
// 
func (a *Temperature) MillidegreesCelsius() float64 {
	if a.millidegrees_celsiusLazy != nil {
		return *a.millidegrees_celsiusLazy
	}
	millidegrees_celsius := a.convertFromBase(TemperatureMillidegreeCelsius)
	a.millidegrees_celsiusLazy = &millidegrees_celsius
	return millidegrees_celsius
}

// DegreesDelisle returns the Temperature value in DegreesDelisle.
//
// 
func (a *Temperature) DegreesDelisle() float64 {
	if a.degrees_delisleLazy != nil {
		return *a.degrees_delisleLazy
	}
	degrees_delisle := a.convertFromBase(TemperatureDegreeDelisle)
	a.degrees_delisleLazy = &degrees_delisle
	return degrees_delisle
}

// DegreesFahrenheit returns the Temperature value in DegreesFahrenheit.
//
// 
func (a *Temperature) DegreesFahrenheit() float64 {
	if a.degrees_fahrenheitLazy != nil {
		return *a.degrees_fahrenheitLazy
	}
	degrees_fahrenheit := a.convertFromBase(TemperatureDegreeFahrenheit)
	a.degrees_fahrenheitLazy = &degrees_fahrenheit
	return degrees_fahrenheit
}

// DegreesNewton returns the Temperature value in DegreesNewton.
//
// 
func (a *Temperature) DegreesNewton() float64 {
	if a.degrees_newtonLazy != nil {
		return *a.degrees_newtonLazy
	}
	degrees_newton := a.convertFromBase(TemperatureDegreeNewton)
	a.degrees_newtonLazy = &degrees_newton
	return degrees_newton
}

// DegreesRankine returns the Temperature value in DegreesRankine.
//
// 
func (a *Temperature) DegreesRankine() float64 {
	if a.degrees_rankineLazy != nil {
		return *a.degrees_rankineLazy
	}
	degrees_rankine := a.convertFromBase(TemperatureDegreeRankine)
	a.degrees_rankineLazy = &degrees_rankine
	return degrees_rankine
}

// DegreesReaumur returns the Temperature value in DegreesReaumur.
//
// 
func (a *Temperature) DegreesReaumur() float64 {
	if a.degrees_reaumurLazy != nil {
		return *a.degrees_reaumurLazy
	}
	degrees_reaumur := a.convertFromBase(TemperatureDegreeReaumur)
	a.degrees_reaumurLazy = &degrees_reaumur
	return degrees_reaumur
}

// DegreesRoemer returns the Temperature value in DegreesRoemer.
//
// 
func (a *Temperature) DegreesRoemer() float64 {
	if a.degrees_roemerLazy != nil {
		return *a.degrees_roemerLazy
	}
	degrees_roemer := a.convertFromBase(TemperatureDegreeRoemer)
	a.degrees_roemerLazy = &degrees_roemer
	return degrees_roemer
}

// SolarTemperatures returns the Temperature value in SolarTemperatures.
//
// 
func (a *Temperature) SolarTemperatures() float64 {
	if a.solar_temperaturesLazy != nil {
		return *a.solar_temperaturesLazy
	}
	solar_temperatures := a.convertFromBase(TemperatureSolarTemperature)
	a.solar_temperaturesLazy = &solar_temperatures
	return solar_temperatures
}


// ToDto creates a TemperatureDto representation from the Temperature instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kelvin by default.
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

// ToDtoJSON creates a JSON representation of the Temperature instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kelvin by default.
func (a *Temperature) ToDtoJSON(holdInUnit *TemperatureUnits) ([]byte, error) {
	// Convert to TemperatureDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Temperature to a specific unit value.
// The function uses the provided unit type (TemperatureUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Temperature in the default unit (Kelvin),
// formatted to two decimal places.
func (a Temperature) String() string {
	return a.ToString(TemperatureKelvin, 2)
}

// ToString formats the Temperature value as a string with the specified unit and fractional digits.
// It converts the Temperature to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Temperature value will be converted (e.g., Kelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Temperature with the unit abbreviation.
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

// Equals checks if the given Temperature is equal to the current Temperature.
//
// Parameters:
//    other: The Temperature to compare against.
//
// Returns:
//    bool: Returns true if both Temperature are equal, false otherwise.
func (a *Temperature) Equals(other *Temperature) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Temperature with another Temperature.
// It returns -1 if the current Temperature is less than the other Temperature, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Temperature to compare against.
//
// Returns:
//    int: -1 if the current Temperature is less, 1 if greater, and 0 if equal.
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

// Add adds the given Temperature to the current Temperature and returns the result.
// The result is a new Temperature instance with the sum of the values.
//
// Parameters:
//    other: The Temperature to add to the current Temperature.
//
// Returns:
//    *Temperature: A new Temperature instance representing the sum of both Temperature.
func (a *Temperature) Add(other *Temperature) *Temperature {
	return &Temperature{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Temperature from the current Temperature and returns the result.
// The result is a new Temperature instance with the difference of the values.
//
// Parameters:
//    other: The Temperature to subtract from the current Temperature.
//
// Returns:
//    *Temperature: A new Temperature instance representing the difference of both Temperature.
func (a *Temperature) Subtract(other *Temperature) *Temperature {
	return &Temperature{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Temperature by the given Temperature and returns the result.
// The result is a new Temperature instance with the product of the values.
//
// Parameters:
//    other: The Temperature to multiply with the current Temperature.
//
// Returns:
//    *Temperature: A new Temperature instance representing the product of both Temperature.
func (a *Temperature) Multiply(other *Temperature) *Temperature {
	return &Temperature{value: a.value * other.BaseValue()}
}

// Divide divides the current Temperature by the given Temperature and returns the result.
// The result is a new Temperature instance with the quotient of the values.
//
// Parameters:
//    other: The Temperature to divide the current Temperature by.
//
// Returns:
//    *Temperature: A new Temperature instance representing the quotient of both Temperature.
func (a *Temperature) Divide(other *Temperature) *Temperature {
	return &Temperature{value: a.value / other.BaseValue()}
}