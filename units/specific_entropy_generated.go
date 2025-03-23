// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificEntropyUnits defines various units of SpecificEntropy.
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

var internalSpecificEntropyUnitsMap = map[SpecificEntropyUnits]bool{
	
	SpecificEntropyJoulePerKilogramKelvin: true,
	SpecificEntropyJoulePerKilogramDegreeCelsius: true,
	SpecificEntropyCaloriePerGramKelvin: true,
	SpecificEntropyBtuPerPoundFahrenheit: true,
	SpecificEntropyKilojoulePerKilogramKelvin: true,
	SpecificEntropyMegajoulePerKilogramKelvin: true,
	SpecificEntropyKilojoulePerKilogramDegreeCelsius: true,
	SpecificEntropyMegajoulePerKilogramDegreeCelsius: true,
	SpecificEntropyKilocaloriePerGramKelvin: true,
}

// SpecificEntropyDto represents a SpecificEntropy measurement with a numerical value and its corresponding unit.
type SpecificEntropyDto struct {
    // Value is the numerical representation of the SpecificEntropy.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the SpecificEntropy, as defined in the SpecificEntropyUnits enumeration.
	Unit  SpecificEntropyUnits `json:"unit" validate:"required,oneof=JoulePerKilogramKelvin JoulePerKilogramDegreeCelsius CaloriePerGramKelvin BtuPerPoundFahrenheit KilojoulePerKilogramKelvin MegajoulePerKilogramKelvin KilojoulePerKilogramDegreeCelsius MegajoulePerKilogramDegreeCelsius KilocaloriePerGramKelvin"`
}

// SpecificEntropyDtoFactory groups methods for creating and serializing SpecificEntropyDto objects.
type SpecificEntropyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpecificEntropyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpecificEntropyDtoFactory) FromJSON(data []byte) (*SpecificEntropyDto, error) {
	a := SpecificEntropyDto{}

    // Parse JSON into SpecificEntropyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a SpecificEntropyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpecificEntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SpecificEntropy represents a measurement in a of SpecificEntropy.
//
// Specific entropy is an amount of energy required to raise temperature of a substance by 1 Kelvin per unit mass.
type SpecificEntropy struct {
	// value is the base measurement stored internally.
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

// SpecificEntropyFactory groups methods for creating SpecificEntropy instances.
type SpecificEntropyFactory struct{}

// CreateSpecificEntropy creates a new SpecificEntropy instance from the given value and unit.
func (uf SpecificEntropyFactory) CreateSpecificEntropy(value float64, unit SpecificEntropyUnits) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, unit)
}

// FromDto converts a SpecificEntropyDto to a SpecificEntropy instance.
func (uf SpecificEntropyFactory) FromDto(dto SpecificEntropyDto) (*SpecificEntropy, error) {
	return newSpecificEntropy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SpecificEntropy instance.
func (uf SpecificEntropyFactory) FromDtoJSON(data []byte) (*SpecificEntropy, error) {
	unitDto, err := SpecificEntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpecificEntropyDto from JSON: %w", err)
	}
	return SpecificEntropyFactory{}.FromDto(*unitDto)
}


// FromJoulesPerKilogramKelvin creates a new SpecificEntropy instance from a value in JoulesPerKilogramKelvin.
func (uf SpecificEntropyFactory) FromJoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyJoulePerKilogramKelvin)
}

// FromJoulesPerKilogramDegreeCelsius creates a new SpecificEntropy instance from a value in JoulesPerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromJoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyJoulePerKilogramDegreeCelsius)
}

// FromCaloriesPerGramKelvin creates a new SpecificEntropy instance from a value in CaloriesPerGramKelvin.
func (uf SpecificEntropyFactory) FromCaloriesPerGramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyCaloriePerGramKelvin)
}

// FromBtusPerPoundFahrenheit creates a new SpecificEntropy instance from a value in BtusPerPoundFahrenheit.
func (uf SpecificEntropyFactory) FromBtusPerPoundFahrenheit(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyBtuPerPoundFahrenheit)
}

// FromKilojoulesPerKilogramKelvin creates a new SpecificEntropy instance from a value in KilojoulesPerKilogramKelvin.
func (uf SpecificEntropyFactory) FromKilojoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilojoulePerKilogramKelvin)
}

// FromMegajoulesPerKilogramKelvin creates a new SpecificEntropy instance from a value in MegajoulesPerKilogramKelvin.
func (uf SpecificEntropyFactory) FromMegajoulesPerKilogramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyMegajoulePerKilogramKelvin)
}

// FromKilojoulesPerKilogramDegreeCelsius creates a new SpecificEntropy instance from a value in KilojoulesPerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromKilojoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilojoulePerKilogramDegreeCelsius)
}

// FromMegajoulesPerKilogramDegreeCelsius creates a new SpecificEntropy instance from a value in MegajoulesPerKilogramDegreeCelsius.
func (uf SpecificEntropyFactory) FromMegajoulesPerKilogramDegreeCelsius(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyMegajoulePerKilogramDegreeCelsius)
}

// FromKilocaloriesPerGramKelvin creates a new SpecificEntropy instance from a value in KilocaloriesPerGramKelvin.
func (uf SpecificEntropyFactory) FromKilocaloriesPerGramKelvin(value float64) (*SpecificEntropy, error) {
	return newSpecificEntropy(value, SpecificEntropyKilocaloriePerGramKelvin)
}


// newSpecificEntropy creates a new SpecificEntropy.
func newSpecificEntropy(value float64, fromUnit SpecificEntropyUnits) (*SpecificEntropy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalSpecificEntropyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in SpecificEntropyUnits", fromUnit)
	}
	a := &SpecificEntropy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificEntropy in JoulePerKilogramKelvin unit (the base/default quantity).
func (a *SpecificEntropy) BaseValue() float64 {
	return a.value
}


// JoulesPerKilogramKelvin returns the SpecificEntropy value in JoulesPerKilogramKelvin.
//
// 
func (a *SpecificEntropy) JoulesPerKilogramKelvin() float64 {
	if a.joules_per_kilogram_kelvinLazy != nil {
		return *a.joules_per_kilogram_kelvinLazy
	}
	joules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyJoulePerKilogramKelvin)
	a.joules_per_kilogram_kelvinLazy = &joules_per_kilogram_kelvin
	return joules_per_kilogram_kelvin
}

// JoulesPerKilogramDegreeCelsius returns the SpecificEntropy value in JoulesPerKilogramDegreeCelsius.
//
// 
func (a *SpecificEntropy) JoulesPerKilogramDegreeCelsius() float64 {
	if a.joules_per_kilogram_degree_celsiusLazy != nil {
		return *a.joules_per_kilogram_degree_celsiusLazy
	}
	joules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyJoulePerKilogramDegreeCelsius)
	a.joules_per_kilogram_degree_celsiusLazy = &joules_per_kilogram_degree_celsius
	return joules_per_kilogram_degree_celsius
}

// CaloriesPerGramKelvin returns the SpecificEntropy value in CaloriesPerGramKelvin.
//
// 
func (a *SpecificEntropy) CaloriesPerGramKelvin() float64 {
	if a.calories_per_gram_kelvinLazy != nil {
		return *a.calories_per_gram_kelvinLazy
	}
	calories_per_gram_kelvin := a.convertFromBase(SpecificEntropyCaloriePerGramKelvin)
	a.calories_per_gram_kelvinLazy = &calories_per_gram_kelvin
	return calories_per_gram_kelvin
}

// BtusPerPoundFahrenheit returns the SpecificEntropy value in BtusPerPoundFahrenheit.
//
// 
func (a *SpecificEntropy) BtusPerPoundFahrenheit() float64 {
	if a.btus_per_pound_fahrenheitLazy != nil {
		return *a.btus_per_pound_fahrenheitLazy
	}
	btus_per_pound_fahrenheit := a.convertFromBase(SpecificEntropyBtuPerPoundFahrenheit)
	a.btus_per_pound_fahrenheitLazy = &btus_per_pound_fahrenheit
	return btus_per_pound_fahrenheit
}

// KilojoulesPerKilogramKelvin returns the SpecificEntropy value in KilojoulesPerKilogramKelvin.
//
// 
func (a *SpecificEntropy) KilojoulesPerKilogramKelvin() float64 {
	if a.kilojoules_per_kilogram_kelvinLazy != nil {
		return *a.kilojoules_per_kilogram_kelvinLazy
	}
	kilojoules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyKilojoulePerKilogramKelvin)
	a.kilojoules_per_kilogram_kelvinLazy = &kilojoules_per_kilogram_kelvin
	return kilojoules_per_kilogram_kelvin
}

// MegajoulesPerKilogramKelvin returns the SpecificEntropy value in MegajoulesPerKilogramKelvin.
//
// 
func (a *SpecificEntropy) MegajoulesPerKilogramKelvin() float64 {
	if a.megajoules_per_kilogram_kelvinLazy != nil {
		return *a.megajoules_per_kilogram_kelvinLazy
	}
	megajoules_per_kilogram_kelvin := a.convertFromBase(SpecificEntropyMegajoulePerKilogramKelvin)
	a.megajoules_per_kilogram_kelvinLazy = &megajoules_per_kilogram_kelvin
	return megajoules_per_kilogram_kelvin
}

// KilojoulesPerKilogramDegreeCelsius returns the SpecificEntropy value in KilojoulesPerKilogramDegreeCelsius.
//
// 
func (a *SpecificEntropy) KilojoulesPerKilogramDegreeCelsius() float64 {
	if a.kilojoules_per_kilogram_degree_celsiusLazy != nil {
		return *a.kilojoules_per_kilogram_degree_celsiusLazy
	}
	kilojoules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyKilojoulePerKilogramDegreeCelsius)
	a.kilojoules_per_kilogram_degree_celsiusLazy = &kilojoules_per_kilogram_degree_celsius
	return kilojoules_per_kilogram_degree_celsius
}

// MegajoulesPerKilogramDegreeCelsius returns the SpecificEntropy value in MegajoulesPerKilogramDegreeCelsius.
//
// 
func (a *SpecificEntropy) MegajoulesPerKilogramDegreeCelsius() float64 {
	if a.megajoules_per_kilogram_degree_celsiusLazy != nil {
		return *a.megajoules_per_kilogram_degree_celsiusLazy
	}
	megajoules_per_kilogram_degree_celsius := a.convertFromBase(SpecificEntropyMegajoulePerKilogramDegreeCelsius)
	a.megajoules_per_kilogram_degree_celsiusLazy = &megajoules_per_kilogram_degree_celsius
	return megajoules_per_kilogram_degree_celsius
}

// KilocaloriesPerGramKelvin returns the SpecificEntropy value in KilocaloriesPerGramKelvin.
//
// 
func (a *SpecificEntropy) KilocaloriesPerGramKelvin() float64 {
	if a.kilocalories_per_gram_kelvinLazy != nil {
		return *a.kilocalories_per_gram_kelvinLazy
	}
	kilocalories_per_gram_kelvin := a.convertFromBase(SpecificEntropyKilocaloriePerGramKelvin)
	a.kilocalories_per_gram_kelvinLazy = &kilocalories_per_gram_kelvin
	return kilocalories_per_gram_kelvin
}


// ToDto creates a SpecificEntropyDto representation from the SpecificEntropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKilogramKelvin by default.
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

// ToDtoJSON creates a JSON representation of the SpecificEntropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerKilogramKelvin by default.
func (a *SpecificEntropy) ToDtoJSON(holdInUnit *SpecificEntropyUnits) ([]byte, error) {
	// Convert to SpecificEntropyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SpecificEntropy to a specific unit value.
// The function uses the provided unit type (SpecificEntropyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the SpecificEntropy in the default unit (JoulePerKilogramKelvin),
// formatted to two decimal places.
func (a SpecificEntropy) String() string {
	return a.ToString(SpecificEntropyJoulePerKilogramKelvin, 2)
}

// ToString formats the SpecificEntropy value as a string with the specified unit and fractional digits.
// It converts the SpecificEntropy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SpecificEntropy value will be converted (e.g., JoulePerKilogramKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SpecificEntropy with the unit abbreviation.
func (a *SpecificEntropy) ToString(unit SpecificEntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSpecificEntropyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSpecificEntropyAbbreviation(unit))
}

// Equals checks if the given SpecificEntropy is equal to the current SpecificEntropy.
//
// Parameters:
//    other: The SpecificEntropy to compare against.
//
// Returns:
//    bool: Returns true if both SpecificEntropy are equal, false otherwise.
func (a *SpecificEntropy) Equals(other *SpecificEntropy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SpecificEntropy with another SpecificEntropy.
// It returns -1 if the current SpecificEntropy is less than the other SpecificEntropy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SpecificEntropy to compare against.
//
// Returns:
//    int: -1 if the current SpecificEntropy is less, 1 if greater, and 0 if equal.
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

// Add adds the given SpecificEntropy to the current SpecificEntropy and returns the result.
// The result is a new SpecificEntropy instance with the sum of the values.
//
// Parameters:
//    other: The SpecificEntropy to add to the current SpecificEntropy.
//
// Returns:
//    *SpecificEntropy: A new SpecificEntropy instance representing the sum of both SpecificEntropy.
func (a *SpecificEntropy) Add(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SpecificEntropy from the current SpecificEntropy and returns the result.
// The result is a new SpecificEntropy instance with the difference of the values.
//
// Parameters:
//    other: The SpecificEntropy to subtract from the current SpecificEntropy.
//
// Returns:
//    *SpecificEntropy: A new SpecificEntropy instance representing the difference of both SpecificEntropy.
func (a *SpecificEntropy) Subtract(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SpecificEntropy by the given SpecificEntropy and returns the result.
// The result is a new SpecificEntropy instance with the product of the values.
//
// Parameters:
//    other: The SpecificEntropy to multiply with the current SpecificEntropy.
//
// Returns:
//    *SpecificEntropy: A new SpecificEntropy instance representing the product of both SpecificEntropy.
func (a *SpecificEntropy) Multiply(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value * other.BaseValue()}
}

// Divide divides the current SpecificEntropy by the given SpecificEntropy and returns the result.
// The result is a new SpecificEntropy instance with the quotient of the values.
//
// Parameters:
//    other: The SpecificEntropy to divide the current SpecificEntropy by.
//
// Returns:
//    *SpecificEntropy: A new SpecificEntropy instance representing the quotient of both SpecificEntropy.
func (a *SpecificEntropy) Divide(other *SpecificEntropy) *SpecificEntropy {
	return &SpecificEntropy{value: a.value / other.BaseValue()}
}

// GetSpecificEntropyAbbreviation gets the unit abbreviation.
func GetSpecificEntropyAbbreviation(unit SpecificEntropyUnits) string {
	switch unit { 
	case SpecificEntropyJoulePerKilogramKelvin:
		return "J/kg·K" 
	case SpecificEntropyJoulePerKilogramDegreeCelsius:
		return "J/kg·°C" 
	case SpecificEntropyCaloriePerGramKelvin:
		return "cal/g·K" 
	case SpecificEntropyBtuPerPoundFahrenheit:
		return "BTU/lb·°F" 
	case SpecificEntropyKilojoulePerKilogramKelvin:
		return "kJ/kg·K" 
	case SpecificEntropyMegajoulePerKilogramKelvin:
		return "MJ/kg·K" 
	case SpecificEntropyKilojoulePerKilogramDegreeCelsius:
		return "kJ/kg·°C" 
	case SpecificEntropyMegajoulePerKilogramDegreeCelsius:
		return "MJ/kg·°C" 
	case SpecificEntropyKilocaloriePerGramKelvin:
		return "kcal/g·K" 
	default:
		return ""
	}
}