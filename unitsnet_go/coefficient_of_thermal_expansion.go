package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// CoefficientOfThermalExpansionUnits enumeration
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

// CoefficientOfThermalExpansionDto represents an CoefficientOfThermalExpansion
type CoefficientOfThermalExpansionDto struct {
	Value float64
	Unit  CoefficientOfThermalExpansionUnits
}

// CoefficientOfThermalExpansionDtoFactory struct to group related functions
type CoefficientOfThermalExpansionDtoFactory struct{}

func (udf CoefficientOfThermalExpansionDtoFactory) FromJSON(data []byte) (*CoefficientOfThermalExpansionDto, error) {
	a := CoefficientOfThermalExpansionDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a CoefficientOfThermalExpansionDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// CoefficientOfThermalExpansion struct
type CoefficientOfThermalExpansion struct {
	value       float64
    
    per_kelvinLazy *float64 
    per_degree_celsiusLazy *float64 
    per_degree_fahrenheitLazy *float64 
    ppm_per_kelvinLazy *float64 
    ppm_per_degree_celsiusLazy *float64 
    ppm_per_degree_fahrenheitLazy *float64 
}

// CoefficientOfThermalExpansionFactory struct to group related functions
type CoefficientOfThermalExpansionFactory struct{}

func (uf CoefficientOfThermalExpansionFactory) CreateCoefficientOfThermalExpansion(value float64, unit CoefficientOfThermalExpansionUnits) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, unit)
}

func (uf CoefficientOfThermalExpansionFactory) FromDto(dto CoefficientOfThermalExpansionDto) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(dto.Value, dto.Unit)
}

func (uf CoefficientOfThermalExpansionFactory) FromDtoJSON(data []byte) (*CoefficientOfThermalExpansion, error) {
	unitDto, err := CoefficientOfThermalExpansionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return CoefficientOfThermalExpansionFactory{}.FromDto(*unitDto)
}


// FromPerKelvin creates a new CoefficientOfThermalExpansion instance from PerKelvin.
func (uf CoefficientOfThermalExpansionFactory) FromPerKelvin(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerKelvin)
}

// FromPerDegreeCelsius creates a new CoefficientOfThermalExpansion instance from PerDegreeCelsius.
func (uf CoefficientOfThermalExpansionFactory) FromPerDegreeCelsius(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerDegreeCelsius)
}

// FromPerDegreeFahrenheit creates a new CoefficientOfThermalExpansion instance from PerDegreeFahrenheit.
func (uf CoefficientOfThermalExpansionFactory) FromPerDegreeFahrenheit(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPerDegreeFahrenheit)
}

// FromPpmPerKelvin creates a new CoefficientOfThermalExpansion instance from PpmPerKelvin.
func (uf CoefficientOfThermalExpansionFactory) FromPpmPerKelvin(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPpmPerKelvin)
}

// FromPpmPerDegreeCelsius creates a new CoefficientOfThermalExpansion instance from PpmPerDegreeCelsius.
func (uf CoefficientOfThermalExpansionFactory) FromPpmPerDegreeCelsius(value float64) (*CoefficientOfThermalExpansion, error) {
	return newCoefficientOfThermalExpansion(value, CoefficientOfThermalExpansionPpmPerDegreeCelsius)
}

// FromPpmPerDegreeFahrenheit creates a new CoefficientOfThermalExpansion instance from PpmPerDegreeFahrenheit.
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

// BaseValue returns the base value of CoefficientOfThermalExpansion in PerKelvin.
func (a *CoefficientOfThermalExpansion) BaseValue() float64 {
	return a.value
}


// PerKelvin returns the value in PerKelvin.
func (a *CoefficientOfThermalExpansion) PerKelvin() float64 {
	if a.per_kelvinLazy != nil {
		return *a.per_kelvinLazy
	}
	per_kelvin := a.convertFromBase(CoefficientOfThermalExpansionPerKelvin)
	a.per_kelvinLazy = &per_kelvin
	return per_kelvin
}

// PerDegreeCelsius returns the value in PerDegreeCelsius.
func (a *CoefficientOfThermalExpansion) PerDegreeCelsius() float64 {
	if a.per_degree_celsiusLazy != nil {
		return *a.per_degree_celsiusLazy
	}
	per_degree_celsius := a.convertFromBase(CoefficientOfThermalExpansionPerDegreeCelsius)
	a.per_degree_celsiusLazy = &per_degree_celsius
	return per_degree_celsius
}

// PerDegreeFahrenheit returns the value in PerDegreeFahrenheit.
func (a *CoefficientOfThermalExpansion) PerDegreeFahrenheit() float64 {
	if a.per_degree_fahrenheitLazy != nil {
		return *a.per_degree_fahrenheitLazy
	}
	per_degree_fahrenheit := a.convertFromBase(CoefficientOfThermalExpansionPerDegreeFahrenheit)
	a.per_degree_fahrenheitLazy = &per_degree_fahrenheit
	return per_degree_fahrenheit
}

// PpmPerKelvin returns the value in PpmPerKelvin.
func (a *CoefficientOfThermalExpansion) PpmPerKelvin() float64 {
	if a.ppm_per_kelvinLazy != nil {
		return *a.ppm_per_kelvinLazy
	}
	ppm_per_kelvin := a.convertFromBase(CoefficientOfThermalExpansionPpmPerKelvin)
	a.ppm_per_kelvinLazy = &ppm_per_kelvin
	return ppm_per_kelvin
}

// PpmPerDegreeCelsius returns the value in PpmPerDegreeCelsius.
func (a *CoefficientOfThermalExpansion) PpmPerDegreeCelsius() float64 {
	if a.ppm_per_degree_celsiusLazy != nil {
		return *a.ppm_per_degree_celsiusLazy
	}
	ppm_per_degree_celsius := a.convertFromBase(CoefficientOfThermalExpansionPpmPerDegreeCelsius)
	a.ppm_per_degree_celsiusLazy = &ppm_per_degree_celsius
	return ppm_per_degree_celsius
}

// PpmPerDegreeFahrenheit returns the value in PpmPerDegreeFahrenheit.
func (a *CoefficientOfThermalExpansion) PpmPerDegreeFahrenheit() float64 {
	if a.ppm_per_degree_fahrenheitLazy != nil {
		return *a.ppm_per_degree_fahrenheitLazy
	}
	ppm_per_degree_fahrenheit := a.convertFromBase(CoefficientOfThermalExpansionPpmPerDegreeFahrenheit)
	a.ppm_per_degree_fahrenheitLazy = &ppm_per_degree_fahrenheit
	return ppm_per_degree_fahrenheit
}


// ToDto creates an CoefficientOfThermalExpansionDto representation.
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

// ToDtoJSON creates an CoefficientOfThermalExpansionDto representation.
func (a *CoefficientOfThermalExpansion) ToDtoJSON(holdInUnit *CoefficientOfThermalExpansionUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts CoefficientOfThermalExpansion to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a CoefficientOfThermalExpansion) String() string {
	return a.ToString(CoefficientOfThermalExpansionPerKelvin, 2)
}

// ToString formats the CoefficientOfThermalExpansion to string.
// fractionalDigits -1 for not mention
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

// Check if the given CoefficientOfThermalExpansion are equals to the current CoefficientOfThermalExpansion
func (a *CoefficientOfThermalExpansion) Equals(other *CoefficientOfThermalExpansion) bool {
	return a.value == other.BaseValue()
}

// Check if the given CoefficientOfThermalExpansion are equals to the current CoefficientOfThermalExpansion
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

// Add the given CoefficientOfThermalExpansion to the current CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Add(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value + other.BaseValue()}
}

// Subtract the given CoefficientOfThermalExpansion to the current CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Subtract(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value - other.BaseValue()}
}

// Multiply the given CoefficientOfThermalExpansion to the current CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Multiply(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value * other.BaseValue()}
}

// Divide the given CoefficientOfThermalExpansion to the current CoefficientOfThermalExpansion.
func (a *CoefficientOfThermalExpansion) Divide(other *CoefficientOfThermalExpansion) *CoefficientOfThermalExpansion {
	return &CoefficientOfThermalExpansion{value: a.value / other.BaseValue()}
}