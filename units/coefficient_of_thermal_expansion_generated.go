// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// CoefficientOfThermalExpansionUnits defines various units of CoefficientOfThermalExpansion.
type CoefficientOfThermalExpansionUnits string

const (
    
        // 
        CoefficientOfThermalExpansionPerKelvin CoefficientOfThermalExpansionUnits = "PerKelvin"
        // 
        CoefficientOfThermalExpansionPerDegreeCelsius CoefficientOfThermalExpansionUnits = "PerDegreeCelsius"
        // 
        CoefficientOfThermalExpansionPerDegreeFahrenheit CoefficientOfThermalExpansionUnits = "PerDegreeFahrenheit"
        // 
        CoefficientOfThermalExpansionPpmPerKelvin CoefficientOfThermalExpansionUnits = "PpmPerKelvin"
        // 
        CoefficientOfThermalExpansionPpmPerDegreeCelsius CoefficientOfThermalExpansionUnits = "PpmPerDegreeCelsius"
        // 
        CoefficientOfThermalExpansionPpmPerDegreeFahrenheit CoefficientOfThermalExpansionUnits = "PpmPerDegreeFahrenheit"
)

// CoefficientOfThermalExpansionDto represents a CoefficientOfThermalExpansion measurement with a numerical value and its corresponding unit.
type CoefficientOfThermalExpansionDto struct {
    // Value is the numerical representation of the CoefficientOfThermalExpansion.
	Value float64
    // Unit specifies the unit of measurement for the CoefficientOfThermalExpansion, as defined in the CoefficientOfThermalExpansionUnits enumeration.
	Unit  CoefficientOfThermalExpansionUnits
}

// CoefficientOfThermalExpansionDtoFactory groups methods for creating and serializing CoefficientOfThermalExpansionDto objects.
type CoefficientOfThermalExpansionDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a CoefficientOfThermalExpansionDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf CoefficientOfThermalExpansionDtoFactory) FromJSON(data []byte) (*CoefficientOfThermalExpansionDto, error) {
	a := CoefficientOfThermalExpansionDto{}

    // Parse JSON into CoefficientOfThermalExpansionDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a CoefficientOfThermalExpansionDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a CoefficientOfThermalExpansionDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// CoefficientOfThermalExpansion represents a measurement in a of CoefficientOfThermalExpansion.
//
// A unit that represents a fractional change in size in response to a change in temperature.
type CoefficientOfThermalExpansion struct {
	// value is the base measurement stored internally.
	value       float64
    
    per_kelvinLazy *float64 
    per_degree_celsiusLazy *float64 
    per_degree_fahrenheitLazy *float64 
    ppm_per_kelvinLazy *float64 
    ppm_per_degree_celsiusLazy *float64 
    ppm_per_degree_fahrenheitLazy *float64 
}

// CoefficientOfThermalExpansionFactory groups methods for creating CoefficientOfThermalExpansion instances.
type CoefficientOfThermalExpansionFactory struct{}

// CreateCoefficientOfThermalExpansion creates a new CoefficientOfThermalExpansion instance from the given value and unit.
func (uf CoefficientOfThermalExpansionFactory) CreateCoefficientOfThermalExpansion(value float64, unit CoefficientOfThermalExpansionUnits) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, unit)
}

// FromDto converts a CoefficientOfThermalExpansionDto to a CoefficientOfThermalExpansion instance.
func (uf CoefficientOfThermalExpansionFactory) FromDto(dto CoefficientOfThermalExpansionDto) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a CoefficientOfThermalExpansion instance.
func (uf CoefficientOfThermalExpansionFactory) FromDtoJSON(data []byte) (*CoefficientOfThermalExpansion, error) {
	unitDto, err := CoefficientOfThermalExpansionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CoefficientOfThermalExpansionDto from JSON: %w", err)
	}
	return CoefficientOfThermalExpansionFactory{}.FromDto(*unitDto)
}


// FromPerKelvin creates a new CoefficientOfThermalExpansion instance from a value in PerKelvin.
func (uf CoefficientOfThermalExpansionFactory) FromPerKelvin(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerKelvin)
}

// FromPerDegreeCelsius creates a new CoefficientOfThermalExpansion instance from a value in PerDegreeCelsius.
func (uf CoefficientOfThermalExpansionFactory) FromPerDegreeCelsius(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerDegreeCelsius)
}

// FromPerDegreeFahrenheit creates a new CoefficientOfThermalExpansion instance from a value in PerDegreeFahrenheit.
func (uf CoefficientOfThermalExpansionFactory) FromPerDegreeFahrenheit(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerDegreeFahrenheit)
}

// FromPpmPerKelvin creates a new CoefficientOfThermalExpansion instance from a value in PpmPerKelvin.
func (uf CoefficientOfThermalExpansionFactory) FromPpmPerKelvin(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPpmPerKelvin)
}

// FromPpmPerDegreeCelsius creates a new CoefficientOfThermalExpansion instance from a value in PpmPerDegreeCelsius.
func (uf CoefficientOfThermalExpansionFactory) FromPpmPerDegreeCelsius(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPpmPerDegreeCelsius)
}

// FromPpmPerDegreeFahrenheit creates a new CoefficientOfThermalExpansion instance from a value in PpmPerDegreeFahrenheit.
func (uf CoefficientOfThermalExpansionFactory) FromPpmPerDegreeFahrenheit(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
}


// newCoefficientOfThermalExpansion creates a new CoefficientOfThermalExpansion.
func newCoefficientOfThermalExpansion(value float64, fromUnit CoefficientOfThermalExpansionUnits) (*CoefficientOfThermalExpansion, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &CoefficientOfThermalExpansion{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of CoefficientOfThermalExpansion in PerKelvin unit (the base/default quantity).
func (a *CoefficientOfThermalExpansion) BaseValue() float64 {
	return a.value
}


// PerKelvin returns the CoefficientOfThermalExpansion value in PerKelvin.
//
// 
func (a *CoefficientOfThermalExpansion) PerKelvin() float64 {
	if a.per_kelvinLazy != nil {
		return *a.per_kelvinLazy
	}
	per_kelvin := a.convertFromBase(CoefficientOfThermalExpansionPerKelvin)
	a.per_kelvinLazy = &per_kelvin
	return per_kelvin
}

// PerDegreeCelsius returns the CoefficientOfThermalExpansion value in PerDegreeCelsius.
//
// 
func (a *CoefficientOfThermalExpansion) PerDegreeCelsius() float64 {
	if a.per_degree_celsiusLazy != nil {
		return *a.per_degree_celsiusLazy
	}
	per_degree_celsius := a.convertFromBase(CoefficientOfThermalExpansionPerDegreeCelsius)
	a.per_degree_celsiusLazy = &per_degree_celsius
	return per_degree_celsius
}

// PerDegreeFahrenheit returns the CoefficientOfThermalExpansion value in PerDegreeFahrenheit.
//
// 
func (a *CoefficientOfThermalExpansion) PerDegreeFahrenheit() float64 {
	if a.per_degree_fahrenheitLazy != nil {
		return *a.per_degree_fahrenheitLazy
	}
	per_degree_fahrenheit := a.convertFromBase(CoefficientOfThermalExpansionPerDegreeFahrenheit)
	a.per_degree_fahrenheitLazy = &per_degree_fahrenheit
	return per_degree_fahrenheit
}

// PpmPerKelvin returns the CoefficientOfThermalExpansion value in PpmPerKelvin.
//
// 
func (a *CoefficientOfThermalExpansion) PpmPerKelvin() float64 {
	if a.ppm_per_kelvinLazy != nil {
		return *a.ppm_per_kelvinLazy
	}
	ppm_per_kelvin := a.convertFromBase(CoefficientOfThermalExpansionPpmPerKelvin)
	a.ppm_per_kelvinLazy = &ppm_per_kelvin
	return ppm_per_kelvin
}

// PpmPerDegreeCelsius returns the CoefficientOfThermalExpansion value in PpmPerDegreeCelsius.
//
// 
func (a *CoefficientOfThermalExpansion) PpmPerDegreeCelsius() float64 {
	if a.ppm_per_degree_celsiusLazy != nil {
		return *a.ppm_per_degree_celsiusLazy
	}
	ppm_per_degree_celsius := a.convertFromBase(CoefficientOfThermalExpansionPpmPerDegreeCelsius)
	a.ppm_per_degree_celsiusLazy = &ppm_per_degree_celsius
	return ppm_per_degree_celsius
}

// PpmPerDegreeFahrenheit returns the CoefficientOfThermalExpansion value in PpmPerDegreeFahrenheit.
//
// 
func (a *CoefficientOfThermalExpansion) PpmPerDegreeFahrenheit() float64 {
	if a.ppm_per_degree_fahrenheitLazy != nil {
		return *a.ppm_per_degree_fahrenheitLazy
	}
	ppm_per_degree_fahrenheit := a.convertFromBase(CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
	a.ppm_per_degree_fahrenheitLazy = &ppm_per_degree_fahrenheit
	return ppm_per_degree_fahrenheit
}


// ToDto creates a CoefficientOfThermalExpansionDto representation from the CoefficientOfThermalExpansion instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PerKelvin by default.
func (a *CoefficientOfThermalExpansion) ToDto(holdInUnit *CoefficientOfThermalExpansionUnits) CoefficientOfThermalExpansionDto {
	if holdInUnit == nil {
		defaultUnit := CoefficientOfThermalExpansionPerKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return CoefficientOfThermalExpansionDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the CoefficientOfThermalExpansion instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PerKelvin by default.
func (a *CoefficientOfThermalExpansion) ToDtoJSON(holdInUnit *CoefficientOfThermalExpansionUnits) ([]byte, error) {
	// Convert to CoefficientOfThermalExpansionDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a CoefficientOfThermalExpansion to a specific unit value.
// The function uses the provided unit type (CoefficientOfThermalExpansionUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *CoefficientOfThermalExpansion) Convert(toUnit CoefficientOfThermalExpansionUnits) float64 {
	switch toUnit { 
    case CoefficientOfThermalExpansionPerKelvin:
		return a.PerKelvin()
    case CoefficientOfThermalExpansionPerDegreeCelsius:
		return a.PerDegreeCelsius()
    case CoefficientOfThermalExpansionPerDegreeFahrenheit:
		return a.PerDegreeFahrenheit()
    case CoefficientOfThermalExpansionPpmPerKelvin:
		return a.PpmPerKelvin()
    case CoefficientOfThermalExpansionPpmPerDegreeCelsius:
		return a.PpmPerDegreeCelsius()
    case CoefficientOfThermalExpansionPpmPerDegreeFahrenheit:
		return a.PpmPerDegreeFahrenheit()
	default:
		return math.NaN()
	}
}

func (a *CoefficientOfThermalExpansion) convertFromBase(toUnit CoefficientOfThermalExpansionUnits) float64 {
    value := a.value
	switch toUnit { 
	case CoefficientOfThermalExpansionPerKelvin:
		return (value) 
	case CoefficientOfThermalExpansionPerDegreeCelsius:
		return (value) 
	case CoefficientOfThermalExpansionPerDegreeFahrenheit:
		return (value * 5 / 9) 
	case CoefficientOfThermalExpansionPpmPerKelvin:
		return (value * 1e6) 
	case CoefficientOfThermalExpansionPpmPerDegreeCelsius:
		return (value * 1e6) 
	case CoefficientOfThermalExpansionPpmPerDegreeFahrenheit:
		return (value * 5e6 / 9) 
	default:
		return math.NaN()
	}
}

func (a *CoefficientOfThermalExpansion) convertToBase(value float64, fromUnit CoefficientOfThermalExpansionUnits) float64 {
	switch fromUnit { 
	case CoefficientOfThermalExpansionPerKelvin:
		return (value) 
	case CoefficientOfThermalExpansionPerDegreeCelsius:
		return (value) 
	case CoefficientOfThermalExpansionPerDegreeFahrenheit:
		return (value * 9 / 5) 
	case CoefficientOfThermalExpansionPpmPerKelvin:
		return (value / 1e6) 
	case CoefficientOfThermalExpansionPpmPerDegreeCelsius:
		return (value / 1e6) 
	case CoefficientOfThermalExpansionPpmPerDegreeFahrenheit:
		return (value * 9 / 5e6) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the CoefficientOfThermalExpansion in the default unit (PerKelvin),
// formatted to two decimal places.
func (a CoefficientOfThermalExpansion) String() string {
	return a.ToString(CoefficientOfThermalExpansionPerKelvin, 2)
}

// ToString formats the CoefficientOfThermalExpansion value as a string with the specified unit and fractional digits.
// It converts the CoefficientOfThermalExpansion to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the CoefficientOfThermalExpansion value will be converted (e.g., PerKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the CoefficientOfThermalExpansion with the unit abbreviation.
func (a *CoefficientOfThermalExpansion) ToString(unit CoefficientOfThermalExpansionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *CoefficientOfThermalExpansion) getUnitAbbreviation(unit CoefficientOfThermalExpansionUnits) string {
	switch unit { 
	case CoefficientOfThermalExpansionPerKelvin:
		return "K⁻¹" 
	case CoefficientOfThermalExpansionPerDegreeCelsius:
		return "°C⁻¹" 
	case CoefficientOfThermalExpansionPerDegreeFahrenheit:
		return "°F⁻¹" 
	case CoefficientOfThermalExpansionPpmPerKelvin:
		return "ppm/K" 
	case CoefficientOfThermalExpansionPpmPerDegreeCelsius:
		return "ppm/°C" 
	case CoefficientOfThermalExpansionPpmPerDegreeFahrenheit:
		return "ppm/°F" 
	default:
		return ""
	}
}

// Equals checks if the given CoefficientOfThermalExpansion is equal to the current CoefficientOfThermalExpansion.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to compare against.
//
// Returns:
//    bool: Returns true if both CoefficientOfThermalExpansion are equal, false otherwise.
func (a *CoefficientOfThermalExpansion) Equals(other *CoefficientOfThermalExpansion) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current CoefficientOfThermalExpansion with another CoefficientOfThermalExpansion.
// It returns -1 if the current CoefficientOfThermalExpansion is less than the other CoefficientOfThermalExpansion, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to compare against.
//
// Returns:
//    int: -1 if the current CoefficientOfThermalExpansion is less, 1 if greater, and 0 if equal.
func (a *CoefficientOfThermalExpansion) CompareTo(other *CoefficientOfThermalExpansion) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given CoefficientOfThermalExpansion to the current CoefficientOfThermalExpansion and returns the result.
// The result is a new CoefficientOfThermalExpansion instance with the sum of the values.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to add to the current CoefficientOfThermalExpansion.
//
// Returns:
//    *CoefficientOfThermalExpansion: A new CoefficientOfThermalExpansion instance representing the sum of both CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Add(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given CoefficientOfThermalExpansion from the current CoefficientOfThermalExpansion and returns the result.
// The result is a new CoefficientOfThermalExpansion instance with the difference of the values.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to subtract from the current CoefficientOfThermalExpansion.
//
// Returns:
//    *CoefficientOfThermalExpansion: A new CoefficientOfThermalExpansion instance representing the difference of both CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Subtract(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current CoefficientOfThermalExpansion by the given CoefficientOfThermalExpansion and returns the result.
// The result is a new CoefficientOfThermalExpansion instance with the product of the values.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to multiply with the current CoefficientOfThermalExpansion.
//
// Returns:
//    *CoefficientOfThermalExpansion: A new CoefficientOfThermalExpansion instance representing the product of both CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Multiply(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value * other.BaseValue()}
}

// Divide divides the current CoefficientOfThermalExpansion by the given CoefficientOfThermalExpansion and returns the result.
// The result is a new CoefficientOfThermalExpansion instance with the quotient of the values.
//
// Parameters:
//    other: The CoefficientOfThermalExpansion to divide the current CoefficientOfThermalExpansion by.
//
// Returns:
//    *CoefficientOfThermalExpansion: A new CoefficientOfThermalExpansion instance representing the quotient of both CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Divide(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value / other.BaseValue()}
}