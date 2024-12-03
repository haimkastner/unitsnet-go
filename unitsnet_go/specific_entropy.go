package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificEntropyUnits enumeration
type SpecificEntropyUnits string

const (
    
        // 
        SpecificEntropyJoulePerKilogramKelvin SpecificEntropyUnits = "JoulePerKilogramKelvin"
        // 
        SpecificEntropyJoulePerKilogramDegreeCelsius SpecificEntropyUnits = "JoulePerKilogramDegreeCelsius"
        // 
        SpecificEntropyCaloriePerGramKelvin SpecificEntropyUnits = "CaloriePerGramKelvin"
        // 
        SpecificEntropyBtuPerPoundFahrenheit SpecificEntropyUnits = "BtuPerPoundFahrenheit"
        // 
        SpecificEntropyKilojoulePerKilogramKelvin SpecificEntropyUnits = "KilojoulePerKilogramKelvin"
        // 
        SpecificEntropyMegajoulePerKilogramKelvin SpecificEntropyUnits = "MegajoulePerKilogramKelvin"
        // 
        SpecificEntropyKilojoulePerKilogramDegreeCelsius SpecificEntropyUnits = "KilojoulePerKilogramDegreeCelsius"
        // 
        SpecificEntropyMegajoulePerKilogramDegreeCelsius SpecificEntropyUnits = "MegajoulePerKilogramDegreeCelsius"
        // 
        SpecificEntropyKilocaloriePerGramKelvin SpecificEntropyUnits = "KilocaloriePerGramKelvin"
)

// SpecificEntropyDto represents an SpecificEntropy
type SpecificEntropyDto struct {
	Value float64
	Unit  SpecificEntropyUnits
}

// SpecificEntropyDtoFactory struct to group related functions
type SpecificEntropyDtoFactory struct{}

func (udf SpecificEntropyDtoFactory) FromJSON(data []byte) (*SpecificEntropyDto, error) {
	a := SpecificEntropyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpecificEntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SpecificEntropy struct
type SpecificEntropy struct {
	value       float64
    
    joules_per_kilogram_kelvinLazy *float64 
    joules_per_kilogram_degree_celsiusLazy *float64 
    calories_per_gram_kelvinLazy *float64 
    btus_per_pound_fahrenheitLazy *float64 
    kilojoules_per_kilogram_kelvinLazy *float64 
    megajoules_per_kilogram_kelvinLazy *float64 
    kilojoules_per_kilogram_degree_celsiusLazy *float64 
    megajoules_per_kilogram_degree_celsiusLazy *float64 
    kilocalories_per_gram_kelvinLazy *float64 
}

// SpecificEntropyFactory struct to group related functions
type SpecificEntropyFactory struct{}

func (uf SpecificEntropyFactory) CreateSpecificEntropy(value float64, unit SpecificEntropyUnits) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, unit)
}

func (uf SpecificEntropyFactory) FromDto(dto SpecificEntropyDto) (*SpecificEntropy, error) {
	return newSpecificEntropy(dto.Value, dto.Unit)
}

func (uf SpecificEntropyFactory) FromDtoJSON(data []byte) (*SpecificEntropy, error) {
	unitDto, err := SpecificEntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpecificEntropyFactory{}.FromDto(*unitDto)
}


// FromJoulePerKilogramKelvin creates a new SpecificEntropy instance from JoulePerKilogramKelvin.
func (uf SpecificEntropyFactory) FromJoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyJoulePerKilogramKelvin)
}

// FromJoulePerKilogramDegreeCelsius creates a new SpecificEntropy instance from JoulePerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromJoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyJoulePerKilogramDegreeCelsius)
}

// FromCaloriePerGramKelvin creates a new SpecificEntropy instance from CaloriePerGramKelvin.
func (uf SpecificEntropyFactory) FromCaloriesPerGramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyCaloriePerGramKelvin)
}

// FromBtuPerPoundFahrenheit creates a new SpecificEntropy instance from BtuPerPoundFahrenheit.
func (uf SpecificEntropyFactory) FromBtusPerPoundFahrenheit(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyBtuPerPoundFahrenheit)
}

// FromKilojoulePerKilogramKelvin creates a new SpecificEntropy instance from KilojoulePerKilogramKelvin.
func (uf SpecificEntropyFactory) FromKilojoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilojoulePerKilogramKelvin)
}

// FromMegajoulePerKilogramKelvin creates a new SpecificEntropy instance from MegajoulePerKilogramKelvin.
func (uf SpecificEntropyFactory) FromMegajoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyMegajoulePerKilogramKelvin)
}

// FromKilojoulePerKilogramDegreeCelsius creates a new SpecificEntropy instance from KilojoulePerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromKilojoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilojoulePerKilogramDegreeCelsius)
}

// FromMegajoulePerKilogramDegreeCelsius creates a new SpecificEntropy instance from MegajoulePerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromMegajoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyMegajoulePerKilogramDegreeCelsius)
}

// FromKilocaloriePerGramKelvin creates a new SpecificEntropy instance from KilocaloriePerGramKelvin.
func (uf SpecificEntropyFactory) FromKilocaloriesPerGramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilocaloriePerGramKelvin)
}




// newSpecificEntropy creates a new SpecificEntropy.
func newSpecificEntropy(value float64, fromUnit SpecificEntropyUnits) (*SpecificEntropy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificEntropy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificEntropy in JoulePerKilogramKelvin.
func (a *SpecificEntropy) BaseValue() float64 {
	return a.value
}


// JoulePerKilogramKelvin returns the value in JoulePerKilogramKelvin.
func (a *SpecificEntropy) JoulesPerKilogramKelvin() float64 {
	if a.joules_per_kilogram_kelvinLazy != nil {
		return *a.joules_per_kilogram_kelvinLazy
	}
	joules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyJoulePerKilogramKelvin)
	a.joules_per_kilogram_kelvinLazy = &joules_per_kilogram_kelvin
	return joules_per_kilogram_kelvin
}

// JoulePerKilogramDegreeCelsius returns the value in JoulePerKilogramDegreeCelsius.
func (a *SpecificEntropy) JoulesPerKilogramDegreeCelsius() float64 {
	if a.joules_per_kilogram_degree_celsiusLazy != nil {
		return *a.joules_per_kilogram_degree_celsiusLazy
	}
	joules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyJoulePerKilogramDegreeCelsius)
	a.joules_per_kilogram_degree_celsiusLazy = &joules_per_kilogram_degree_celsius
	return joules_per_kilogram_degree_celsius
}

// CaloriePerGramKelvin returns the value in CaloriePerGramKelvin.
func (a *SpecificEntropy) CaloriesPerGramKelvin() float64 {
	if a.calories_per_gram_kelvinLazy != nil {
		return *a.calories_per_gram_kelvinLazy
	}
	calories_per_gram_kelvin := a.convertFromBase(SpecificEntropyCaloriePerGramKelvin)
	a.calories_per_gram_kelvinLazy = &calories_per_gram_kelvin
	return calories_per_gram_kelvin
}

// BtuPerPoundFahrenheit returns the value in BtuPerPoundFahrenheit.
func (a *SpecificEntropy) BtusPerPoundFahrenheit() float64 {
	if a.btus_per_pound_fahrenheitLazy != nil {
		return *a.btus_per_pound_fahrenheitLazy
	}
	btus_per_pound_fahrenheit := a.convertFromBase(SpecificEntropyBtuPerPoundFahrenheit)
	a.btus_per_pound_fahrenheitLazy = &btus_per_pound_fahrenheit
	return btus_per_pound_fahrenheit
}

// KilojoulePerKilogramKelvin returns the value in KilojoulePerKilogramKelvin.
func (a *SpecificEntropy) KilojoulesPerKilogramKelvin() float64 {
	if a.kilojoules_per_kilogram_kelvinLazy != nil {
		return *a.kilojoules_per_kilogram_kelvinLazy
	}
	kilojoules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyKilojoulePerKilogramKelvin)
	a.kilojoules_per_kilogram_kelvinLazy = &kilojoules_per_kilogram_kelvin
	return kilojoules_per_kilogram_kelvin
}

// MegajoulePerKilogramKelvin returns the value in MegajoulePerKilogramKelvin.
func (a *SpecificEntropy) MegajoulesPerKilogramKelvin() float64 {
	if a.megajoules_per_kilogram_kelvinLazy != nil {
		return *a.megajoules_per_kilogram_kelvinLazy
	}
	megajoules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyMegajoulePerKilogramKelvin)
	a.megajoules_per_kilogram_kelvinLazy = &megajoules_per_kilogram_kelvin
	return megajoules_per_kilogram_kelvin
}

// KilojoulePerKilogramDegreeCelsius returns the value in KilojoulePerKilogramDegreeCelsius.
func (a *SpecificEntropy) KilojoulesPerKilogramDegreeCelsius() float64 {
	if a.kilojoules_per_kilogram_degree_celsiusLazy != nil {
		return *a.kilojoules_per_kilogram_degree_celsiusLazy
	}
	kilojoules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyKilojoulePerKilogramDegreeCelsius)
	a.kilojoules_per_kilogram_degree_celsiusLazy = &kilojoules_per_kilogram_degree_celsius
	return kilojoules_per_kilogram_degree_celsius
}

// MegajoulePerKilogramDegreeCelsius returns the value in MegajoulePerKilogramDegreeCelsius.
func (a *SpecificEntropy) MegajoulesPerKilogramDegreeCelsius() float64 {
	if a.megajoules_per_kilogram_degree_celsiusLazy != nil {
		return *a.megajoules_per_kilogram_degree_celsiusLazy
	}
	megajoules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyMegajoulePerKilogramDegreeCelsius)
	a.megajoules_per_kilogram_degree_celsiusLazy = &megajoules_per_kilogram_degree_celsius
	return megajoules_per_kilogram_degree_celsius
}

// KilocaloriePerGramKelvin returns the value in KilocaloriePerGramKelvin.
func (a *SpecificEntropy) KilocaloriesPerGramKelvin() float64 {
	if a.kilocalories_per_gram_kelvinLazy != nil {
		return *a.kilocalories_per_gram_kelvinLazy
	}
	kilocalories_per_gram_kelvin := a.convertFromBase(SpecificEntropyKilocaloriePerGramKelvin)
	a.kilocalories_per_gram_kelvinLazy = &kilocalories_per_gram_kelvin
	return kilocalories_per_gram_kelvin
}


// ToDto creates an SpecificEntropyDto representation.
func (a *SpecificEntropy) ToDto(holdInUnit *SpecificEntropyUnits) SpecificEntropyDto {
	if holdInUnit == nil {
		defaultUnit := SpecificEntropyJoulePerKilogramKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificEntropyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an SpecificEntropyDto representation.
func (a *SpecificEntropy) ToDtoJSON(holdInUnit *SpecificEntropyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SpecificEntropy to a specific unit value.
func (a *SpecificEntropy) Convert(toUnit SpecificEntropyUnits) float64 {
	switch toUnit { 
    case SpecificEntropyJoulePerKilogramKelvin:
		return a.JoulesPerKilogramKelvin()
    case SpecificEntropyJoulePerKilogramDegreeCelsius:
		return a.JoulesPerKilogramDegreeCelsius()
    case SpecificEntropyCaloriePerGramKelvin:
		return a.CaloriesPerGramKelvin()
    case SpecificEntropyBtuPerPoundFahrenheit:
		return a.BtusPerPoundFahrenheit()
    case SpecificEntropyKilojoulePerKilogramKelvin:
		return a.KilojoulesPerKilogramKelvin()
    case SpecificEntropyMegajoulePerKilogramKelvin:
		return a.MegajoulesPerKilogramKelvin()
    case SpecificEntropyKilojoulePerKilogramDegreeCelsius:
		return a.KilojoulesPerKilogramDegreeCelsius()
    case SpecificEntropyMegajoulePerKilogramDegreeCelsius:
		return a.MegajoulesPerKilogramDegreeCelsius()
    case SpecificEntropyKilocaloriePerGramKelvin:
		return a.KilocaloriesPerGramKelvin()
	default:
		return 0
	}
}

func (a *SpecificEntropy) convertFromBase(toUnit SpecificEntropyUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificEntropyJoulePerKilogramKelvin:
		return (value) 
	case SpecificEntropyJoulePerKilogramDegreeCelsius:
		return (value) 
	case SpecificEntropyCaloriePerGramKelvin:
		return (value / 4.184e3) 
	case SpecificEntropyBtuPerPoundFahrenheit:
		return (value / 4.1868e3) 
	case SpecificEntropyKilojoulePerKilogramKelvin:
		return ((value) / 1000.0) 
	case SpecificEntropyMegajoulePerKilogramKelvin:
		return ((value) / 1000000.0) 
	case SpecificEntropyKilojoulePerKilogramDegreeCelsius:
		return ((value) / 1000.0) 
	case SpecificEntropyMegajoulePerKilogramDegreeCelsius:
		return ((value) / 1000000.0) 
	case SpecificEntropyKilocaloriePerGramKelvin:
		return ((value / 4.184e3) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *SpecificEntropy) convertToBase(value float64, fromUnit SpecificEntropyUnits) float64 {
	switch fromUnit { 
	case SpecificEntropyJoulePerKilogramKelvin:
		return (value) 
	case SpecificEntropyJoulePerKilogramDegreeCelsius:
		return (value) 
	case SpecificEntropyCaloriePerGramKelvin:
		return (value * 4.184e3) 
	case SpecificEntropyBtuPerPoundFahrenheit:
		return (value * 4.1868e3) 
	case SpecificEntropyKilojoulePerKilogramKelvin:
		return ((value) * 1000.0) 
	case SpecificEntropyMegajoulePerKilogramKelvin:
		return ((value) * 1000000.0) 
	case SpecificEntropyKilojoulePerKilogramDegreeCelsius:
		return ((value) * 1000.0) 
	case SpecificEntropyMegajoulePerKilogramDegreeCelsius:
		return ((value) * 1000000.0) 
	case SpecificEntropyKilocaloriePerGramKelvin:
		return ((value * 4.184e3) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a SpecificEntropy) String() string {
	return a.ToString(SpecificEntropyJoulePerKilogramKelvin, 2)
}

// ToString formats the SpecificEntropy to string.
// fractionalDigits -1 for not mention
func (a *SpecificEntropy) ToString(unit SpecificEntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SpecificEntropy) getUnitAbbreviation(unit SpecificEntropyUnits) string {
	switch unit { 
	case SpecificEntropyJoulePerKilogramKelvin:
		return "J/kg.K" 
	case SpecificEntropyJoulePerKilogramDegreeCelsius:
		return "J/kg.C" 
	case SpecificEntropyCaloriePerGramKelvin:
		return "cal/g.K" 
	case SpecificEntropyBtuPerPoundFahrenheit:
		return "BTU/lb·°F" 
	case SpecificEntropyKilojoulePerKilogramKelvin:
		return "kJ/kg.K" 
	case SpecificEntropyMegajoulePerKilogramKelvin:
		return "MJ/kg.K" 
	case SpecificEntropyKilojoulePerKilogramDegreeCelsius:
		return "kJ/kg.C" 
	case SpecificEntropyMegajoulePerKilogramDegreeCelsius:
		return "MJ/kg.C" 
	case SpecificEntropyKilocaloriePerGramKelvin:
		return "kcal/g.K" 
	default:
		return ""
	}
}

// Check if the given SpecificEntropy are equals to the current SpecificEntropy
func (a *SpecificEntropy) Equals(other *SpecificEntropy) bool {
	return a.value == other.BaseValue()
}

// Check if the given SpecificEntropy are equals to the current SpecificEntropy
func (a *SpecificEntropy) CompareTo(other *SpecificEntropy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given SpecificEntropy to the current SpecificEntropy.
func (a *SpecificEntropy) Add(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value + other.BaseValue()}
}

// Subtract the given SpecificEntropy to the current SpecificEntropy.
func (a *SpecificEntropy) Subtract(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value - other.BaseValue()}
}

// Multiply the given SpecificEntropy to the current SpecificEntropy.
func (a *SpecificEntropy) Multiply(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value * other.BaseValue()}
}

// Divide the given SpecificEntropy to the current SpecificEntropy.
func (a *SpecificEntropy) Divide(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value / other.BaseValue()}
}