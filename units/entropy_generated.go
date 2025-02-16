// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// EntropyUnits defines various units of Entropy.
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

// EntropyDto represents a Entropy measurement with a numerical value and its corresponding unit.
type EntropyDto struct {
    // Value is the numerical representation of the Entropy.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Entropy, as defined in the EntropyUnits enumeration.
	Unit  EntropyUnits `json:"unit"`
}

// EntropyDtoFactory groups methods for creating and serializing EntropyDto objects.
type EntropyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a EntropyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf EntropyDtoFactory) FromJSON(data []byte) (*EntropyDto, error) {
	a := EntropyDto{}

    // Parse JSON into EntropyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a EntropyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a EntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Entropy represents a measurement in a of Entropy.
//
// Entropy is an important concept in the branch of science known as thermodynamics. The idea of "irreversibility" is central to the understanding of entropy.  It is often said that entropy is an expression of the disorder, or randomness of a system, or of our lack of information about it. Entropy is an extensive property. It has the dimension of energy divided by temperature, which has a unit of joules per kelvin (J/K) in the International System of Units
type Entropy struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_kelvinLazy *float64 
    calories_per_kelvinLazy *float64 
    joules_per_degree_celsiusLazy *float64 
    kilojoules_per_kelvinLazy *float64 
    megajoules_per_kelvinLazy *float64 
    kilocalories_per_kelvinLazy *float64 
    kilojoules_per_degree_celsiusLazy *float64 
}

// EntropyFactory groups methods for creating Entropy instances.
type EntropyFactory struct{}

// CreateEntropy creates a new Entropy instance from the given value and unit.
func (uf EntropyFactory) CreateEntropy(value float64, unit EntropyUnits) (*Entropy, error) {
	return newEntropy(value, unit)
}

// FromDto converts a EntropyDto to a Entropy instance.
func (uf EntropyFactory) FromDto(dto EntropyDto) (*Entropy, error) {
	return newEntropy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Entropy instance.
func (uf EntropyFactory) FromDtoJSON(data []byte) (*Entropy, error) {
	unitDto, err := EntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EntropyDto from JSON: %w", err)
	}
	return EntropyFactory{}.FromDto(*unitDto)
}


// FromJoulesPerKelvin creates a new Entropy instance from a value in JoulesPerKelvin.
func (uf EntropyFactory) FromJoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyJoulePerKelvin)
}

// FromCaloriesPerKelvin creates a new Entropy instance from a value in CaloriesPerKelvin.
func (uf EntropyFactory) FromCaloriesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyCaloriePerKelvin)
}

// FromJoulesPerDegreeCelsius creates a new Entropy instance from a value in JoulesPerDegreeCelsius.
func (uf EntropyFactory) FromJoulesPerDegreeCelsius(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyJoulePerDegreeCelsius)
}

// FromKilojoulesPerKelvin creates a new Entropy instance from a value in KilojoulesPerKelvin.
func (uf EntropyFactory) FromKilojoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyKilojoulePerKelvin)
}

// FromMegajoulesPerKelvin creates a new Entropy instance from a value in MegajoulesPerKelvin.
func (uf EntropyFactory) FromMegajoulesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyMegajoulePerKelvin)
}

// FromKilocaloriesPerKelvin creates a new Entropy instance from a value in KilocaloriesPerKelvin.
func (uf EntropyFactory) FromKilocaloriesPerKelvin(value float64) (*Entropy, error) {
	return newEntropy(value, EntropyKilocaloriePerKelvin)
}

// FromKilojoulesPerDegreeCelsius creates a new Entropy instance from a value in KilojoulesPerDegreeCelsius.
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

// BaseValue returns the base value of Entropy in JoulePerKelvin unit (the base/default quantity).
func (a *Entropy) BaseValue() float64 {
	return a.value
}


// JoulesPerKelvin returns the Entropy value in JoulesPerKelvin.
//
// 
func (a *Entropy) JoulesPerKelvin() float64 {
	if a.joules_per_kelvinLazy != nil {
		return *a.joules_per_kelvinLazy
	}
	joules_per_kelvin := a.convertFromBase(EntropyJoulePerKelvin)
	a.joules_per_kelvinLazy = &joules_per_kelvin
	return joules_per_kelvin
}

// CaloriesPerKelvin returns the Entropy value in CaloriesPerKelvin.
//
// 
func (a *Entropy) CaloriesPerKelvin() float64 {
	if a.calories_per_kelvinLazy != nil {
		return *a.calories_per_kelvinLazy
	}
	calories_per_kelvin := a.convertFromBase(EntropyCaloriePerKelvin)
	a.calories_per_kelvinLazy = &calories_per_kelvin
	return calories_per_kelvin
}

// JoulesPerDegreeCelsius returns the Entropy value in JoulesPerDegreeCelsius.
//
// 
func (a *Entropy) JoulesPerDegreeCelsius() float64 {
	if a.joules_per_degree_celsiusLazy != nil {
		return *a.joules_per_degree_celsiusLazy
	}
	joules_per_degree_celsius := a.convertFromBase(EntropyJoulePerDegreeCelsius)
	a.joules_per_degree_celsiusLazy = &joules_per_degree_celsius
	return joules_per_degree_celsius
}

// KilojoulesPerKelvin returns the Entropy value in KilojoulesPerKelvin.
//
// 
func (a *Entropy) KilojoulesPerKelvin() float64 {
	if a.kilojoules_per_kelvinLazy != nil {
		return *a.kilojoules_per_kelvinLazy
	}
	kilojoules_per_kelvin := a.convertFromBase(EntropyKilojoulePerKelvin)
	a.kilojoules_per_kelvinLazy = &kilojoules_per_kelvin
	return kilojoules_per_kelvin
}

// MegajoulesPerKelvin returns the Entropy value in MegajoulesPerKelvin.
//
// 
func (a *Entropy) MegajoulesPerKelvin() float64 {
	if a.megajoules_per_kelvinLazy != nil {
		return *a.megajoules_per_kelvinLazy
	}
	megajoules_per_kelvin := a.convertFromBase(EntropyMegajoulePerKelvin)
	a.megajoules_per_kelvinLazy = &megajoules_per_kelvin
	return megajoules_per_kelvin
}

// KilocaloriesPerKelvin returns the Entropy value in KilocaloriesPerKelvin.
//
// 
func (a *Entropy) KilocaloriesPerKelvin() float64 {
	if a.kilocalories_per_kelvinLazy != nil {
		return *a.kilocalories_per_kelvinLazy
	}
	kilocalories_per_kelvin := a.convertFromBase(EntropyKilocaloriePerKelvin)
	a.kilocalories_per_kelvinLazy = &kilocalories_per_kelvin
	return kilocalories_per_kelvin
}

// KilojoulesPerDegreeCelsius returns the Entropy value in KilojoulesPerDegreeCelsius.
//
// 
func (a *Entropy) KilojoulesPerDegreeCelsius() float64 {
	if a.kilojoules_per_degree_celsiusLazy != nil {
		return *a.kilojoules_per_degree_celsiusLazy
	}
	kilojoules_per_degree_celsius := a.convertFromBase(EntropyKilojoulePerDegreeCelsius)
	a.kilojoules_per_degree_celsiusLazy = &kilojoules_per_degree_celsius
	return kilojoules_per_degree_celsius
}


// ToDto creates a EntropyDto representation from the Entropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKelvin by default.
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

// ToDtoJSON creates a JSON representation of the Entropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKelvin by default.
func (a *Entropy) ToDtoJSON(holdInUnit *EntropyUnits) ([]byte, error) {
	// Convert to EntropyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Entropy to a specific unit value.
// The function uses the provided unit type (EntropyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Entropy in the default unit (JoulePerKelvin),
// formatted to two decimal places.
func (a Entropy) String() string {
	return a.ToString(EntropyJoulePerKelvin, 2)
}

// ToString formats the Entropy value as a string with the specified unit and fractional digits.
// It converts the Entropy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Entropy value will be converted (e.g., JoulePerKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Entropy with the unit abbreviation.
func (a *Entropy) ToString(unit EntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetEntropyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetEntropyAbbreviation(unit))
}

// Equals checks if the given Entropy is equal to the current Entropy.
//
// Parameters:
//    other: The Entropy to compare against.
//
// Returns:
//    bool: Returns true if both Entropy are equal, false otherwise.
func (a *Entropy) Equals(other *Entropy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Entropy with another Entropy.
// It returns -1 if the current Entropy is less than the other Entropy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Entropy to compare against.
//
// Returns:
//    int: -1 if the current Entropy is less, 1 if greater, and 0 if equal.
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

// Add adds the given Entropy to the current Entropy and returns the result.
// The result is a new Entropy instance with the sum of the values.
//
// Parameters:
//    other: The Entropy to add to the current Entropy.
//
// Returns:
//    *Entropy: A new Entropy instance representing the sum of both Entropy.
func (a *Entropy) Add(other *Entropy) *Entropy {
	return &Entropy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Entropy from the current Entropy and returns the result.
// The result is a new Entropy instance with the difference of the values.
//
// Parameters:
//    other: The Entropy to subtract from the current Entropy.
//
// Returns:
//    *Entropy: A new Entropy instance representing the difference of both Entropy.
func (a *Entropy) Subtract(other *Entropy) *Entropy {
	return &Entropy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Entropy by the given Entropy and returns the result.
// The result is a new Entropy instance with the product of the values.
//
// Parameters:
//    other: The Entropy to multiply with the current Entropy.
//
// Returns:
//    *Entropy: A new Entropy instance representing the product of both Entropy.
func (a *Entropy) Multiply(other *Entropy) *Entropy {
	return &Entropy{value: a.value * other.BaseValue()}
}

// Divide divides the current Entropy by the given Entropy and returns the result.
// The result is a new Entropy instance with the quotient of the values.
//
// Parameters:
//    other: The Entropy to divide the current Entropy by.
//
// Returns:
//    *Entropy: A new Entropy instance representing the quotient of both Entropy.
func (a *Entropy) Divide(other *Entropy) *Entropy {
	return &Entropy{value: a.value / other.BaseValue()}
}

// GetEntropyAbbreviation gets the unit abbreviation.
func GetEntropyAbbreviation(unit EntropyUnits) string {
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