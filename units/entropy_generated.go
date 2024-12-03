// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// EntropyUnits enumeration
type EntropyUnits string

const (
    
        // 
        EntropyJoulePerKelvin EntropyUnits = "JoulePerKelvin"
        // 
        EntropyCaloriePerKelvin EntropyUnits = "CaloriePerKelvin"
        // 
        EntropyJoulePerDegreeCelsius EntropyUnits = "JoulePerDegreeCelsius"
        // 
        EntropyKilojoulePerKelvin EntropyUnits = "KilojoulePerKelvin"
        // 
        EntropyMegajoulePerKelvin EntropyUnits = "MegajoulePerKelvin"
        // 
        EntropyKilocaloriePerKelvin EntropyUnits = "KilocaloriePerKelvin"
        // 
        EntropyKilojoulePerDegreeCelsius EntropyUnits = "KilojoulePerDegreeCelsius"
)

// EntropyDto represents an Entropy
type EntropyDto struct {
	Value float64
	Unit  EntropyUnits
}

// EntropyDtoFactory struct to group related functions
type EntropyDtoFactory struct{}

func (udf EntropyDtoFactory) FromJSON(data []byte) (*EntropyDto, error) {
	a := EntropyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a EntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Entropy struct
type Entropy struct {
	value       float64
    
    joules_per_kelvinLazy *float64 
    calories_per_kelvinLazy *float64 
    joules_per_degree_celsiusLazy *float64 
    kilojoules_per_kelvinLazy *float64 
    megajoules_per_kelvinLazy *float64 
    kilocalories_per_kelvinLazy *float64 
    kilojoules_per_degree_celsiusLazy *float64 
}

// EntropyFactory struct to group related functions
type EntropyFactory struct{}

func (uf EntropyFactory) CreateEntropy(value float64, unit EntropyUnits) (*Entropy, error) {
	return newEntropy(value, unit)
}

func (uf EntropyFactory) FromDto(dto EntropyDto) (*Entropy, error) {
	return newEntropy(dto.Value, dto.Unit)
}

func (uf EntropyFactory) FromDtoJSON(data []byte) (*Entropy, error) {
	unitDto, err := EntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return EntropyFactory{}.FromDto(*unitDto)
}


// FromJoulePerKelvin creates a new Entropy instance from JoulePerKelvin.
func (uf EntropyFactory) FromJoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyJoulePerKelvin)
}

// FromCaloriePerKelvin creates a new Entropy instance from CaloriePerKelvin.
func (uf EntropyFactory) FromCaloriesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyCaloriePerKelvin)
}

// FromJoulePerDegreeCelsius creates a new Entropy instance from JoulePerDegreeCelsius.
func (uf EntropyFactory) FromJoulesPerDegreeCelsius(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyJoulePerDegreeCelsius)
}

// FromKilojoulePerKelvin creates a new Entropy instance from KilojoulePerKelvin.
func (uf EntropyFactory) FromKilojoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyKilojoulePerKelvin)
}

// FromMegajoulePerKelvin creates a new Entropy instance from MegajoulePerKelvin.
func (uf EntropyFactory) FromMegajoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyMegajoulePerKelvin)
}

// FromKilocaloriePerKelvin creates a new Entropy instance from KilocaloriePerKelvin.
func (uf EntropyFactory) FromKilocaloriesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyKilocaloriePerKelvin)
}

// FromKilojoulePerDegreeCelsius creates a new Entropy instance from KilojoulePerDegreeCelsius.
func (uf EntropyFactory) FromKilojoulesPerDegreeCelsius(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyKilojoulePerDegreeCelsius)
}




// newEntropy creates a new Entropy.
func newEntropy(value float64, fromUnit EntropyUnits) (*Entropy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Entropy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Entropy in JoulePerKelvin.
func (a *Entropy) BaseValue() float64 {
	return a.value
}


// JoulePerKelvin returns the value in JoulePerKelvin.
func (a *Entropy) JoulesPerKelvin() float64 {
	if a.joules_per_kelvinLazy != nil {
		return *a.joules_per_kelvinLazy
	}
	joules_per_kelvin := a.convertFromBase(EntropyJoulePerKelvin)
	a.joules_per_kelvinLazy = &joules_per_kelvin
	return joules_per_kelvin
}

// CaloriePerKelvin returns the value in CaloriePerKelvin.
func (a *Entropy) CaloriesPerKelvin() float64 {
	if a.calories_per_kelvinLazy != nil {
		return *a.calories_per_kelvinLazy
	}
	calories_per_kelvin := a.convertFromBase(EntropyCaloriePerKelvin)
	a.calories_per_kelvinLazy = &calories_per_kelvin
	return calories_per_kelvin
}

// JoulePerDegreeCelsius returns the value in JoulePerDegreeCelsius.
func (a *Entropy) JoulesPerDegreeCelsius() float64 {
	if a.joules_per_degree_celsiusLazy != nil {
		return *a.joules_per_degree_celsiusLazy
	}
	joules_per_degree_celsius := a.convertFromBase(EntropyJoulePerDegreeCelsius)
	a.joules_per_degree_celsiusLazy = &joules_per_degree_celsius
	return joules_per_degree_celsius
}

// KilojoulePerKelvin returns the value in KilojoulePerKelvin.
func (a *Entropy) KilojoulesPerKelvin() float64 {
	if a.kilojoules_per_kelvinLazy != nil {
		return *a.kilojoules_per_kelvinLazy
	}
	kilojoules_per_kelvin := a.convertFromBase(EntropyKilojoulePerKelvin)
	a.kilojoules_per_kelvinLazy = &kilojoules_per_kelvin
	return kilojoules_per_kelvin
}

// MegajoulePerKelvin returns the value in MegajoulePerKelvin.
func (a *Entropy) MegajoulesPerKelvin() float64 {
	if a.megajoules_per_kelvinLazy != nil {
		return *a.megajoules_per_kelvinLazy
	}
	megajoules_per_kelvin := a.convertFromBase(EntropyMegajoulePerKelvin)
	a.megajoules_per_kelvinLazy = &megajoules_per_kelvin
	return megajoules_per_kelvin
}

// KilocaloriePerKelvin returns the value in KilocaloriePerKelvin.
func (a *Entropy) KilocaloriesPerKelvin() float64 {
	if a.kilocalories_per_kelvinLazy != nil {
		return *a.kilocalories_per_kelvinLazy
	}
	kilocalories_per_kelvin := a.convertFromBase(EntropyKilocaloriePerKelvin)
	a.kilocalories_per_kelvinLazy = &kilocalories_per_kelvin
	return kilocalories_per_kelvin
}

// KilojoulePerDegreeCelsius returns the value in KilojoulePerDegreeCelsius.
func (a *Entropy) KilojoulesPerDegreeCelsius() float64 {
	if a.kilojoules_per_degree_celsiusLazy != nil {
		return *a.kilojoules_per_degree_celsiusLazy
	}
	kilojoules_per_degree_celsius := a.convertFromBase(EntropyKilojoulePerDegreeCelsius)
	a.kilojoules_per_degree_celsiusLazy = &kilojoules_per_degree_celsius
	return kilojoules_per_degree_celsius
}


// ToDto creates an EntropyDto representation.
func (a *Entropy) ToDto(holdInUnit *EntropyUnits) EntropyDto {
	if holdInUnit == nil {
		defaultUnit := EntropyJoulePerKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return EntropyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an EntropyDto representation.
func (a *Entropy) ToDtoJSON(holdInUnit *EntropyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Entropy to a specific unit value.
func (a *Entropy) Convert(toUnit EntropyUnits) float64 {
	switch toUnit { 
    case EntropyJoulePerKelvin:
		return a.JoulesPerKelvin()
    case EntropyCaloriePerKelvin:
		return a.CaloriesPerKelvin()
    case EntropyJoulePerDegreeCelsius:
		return a.JoulesPerDegreeCelsius()
    case EntropyKilojoulePerKelvin:
		return a.KilojoulesPerKelvin()
    case EntropyMegajoulePerKelvin:
		return a.MegajoulesPerKelvin()
    case EntropyKilocaloriePerKelvin:
		return a.KilocaloriesPerKelvin()
    case EntropyKilojoulePerDegreeCelsius:
		return a.KilojoulesPerDegreeCelsius()
	default:
		return 0
	}
}

func (a *Entropy) convertFromBase(toUnit EntropyUnits) float64 {
    value := a.value
	switch toUnit { 
	case EntropyJoulePerKelvin:
		return (value) 
	case EntropyCaloriePerKelvin:
		return (value / 4.184) 
	case EntropyJoulePerDegreeCelsius:
		return (value) 
	case EntropyKilojoulePerKelvin:
		return ((value) / 1000.0) 
	case EntropyMegajoulePerKelvin:
		return ((value) / 1000000.0) 
	case EntropyKilocaloriePerKelvin:
		return ((value / 4.184) / 1000.0) 
	case EntropyKilojoulePerDegreeCelsius:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Entropy) convertToBase(value float64, fromUnit EntropyUnits) float64 {
	switch fromUnit { 
	case EntropyJoulePerKelvin:
		return (value) 
	case EntropyCaloriePerKelvin:
		return (value * 4.184) 
	case EntropyJoulePerDegreeCelsius:
		return (value) 
	case EntropyKilojoulePerKelvin:
		return ((value) * 1000.0) 
	case EntropyMegajoulePerKelvin:
		return ((value) * 1000000.0) 
	case EntropyKilocaloriePerKelvin:
		return ((value * 4.184) * 1000.0) 
	case EntropyKilojoulePerDegreeCelsius:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Entropy) String() string {
	return a.ToString(EntropyJoulePerKelvin, 2)
}

// ToString formats the Entropy to string.
// fractionalDigits -1 for not mention
func (a *Entropy) ToString(unit EntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Entropy) getUnitAbbreviation(unit EntropyUnits) string {
	switch unit { 
	case EntropyJoulePerKelvin:
		return "J/K" 
	case EntropyCaloriePerKelvin:
		return "cal/K" 
	case EntropyJoulePerDegreeCelsius:
		return "J/C" 
	case EntropyKilojoulePerKelvin:
		return "kJ/K" 
	case EntropyMegajoulePerKelvin:
		return "MJ/K" 
	case EntropyKilocaloriePerKelvin:
		return "kcal/K" 
	case EntropyKilojoulePerDegreeCelsius:
		return "kJ/C" 
	default:
		return ""
	}
}

// Check if the given Entropy are equals to the current Entropy
func (a *Entropy) Equals(other *Entropy) bool {
	return a.value == other.BaseValue()
}

// Check if the given Entropy are equals to the current Entropy
func (a *Entropy) CompareTo(other *Entropy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Entropy to the current Entropy.
func (a *Entropy) Add(other *Entropy) *Entropy {
	return &Entropy{value: a.value + other.BaseValue()}
}

// Subtract the given Entropy to the current Entropy.
func (a *Entropy) Subtract(other *Entropy) *Entropy {
	return &Entropy{value: a.value - other.BaseValue()}
}

// Multiply the given Entropy to the current Entropy.
func (a *Entropy) Multiply(other *Entropy) *Entropy {
	return &Entropy{value: a.value * other.BaseValue()}
}

// Divide the given Entropy to the current Entropy.
func (a *Entropy) Divide(other *Entropy) *Entropy {
	return &Entropy{value: a.value / other.BaseValue()}
}