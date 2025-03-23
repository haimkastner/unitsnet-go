// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolarEntropyUnits defines various units of MolarEntropy.
type MolarEntropyUnits string

const (
    
        // 
        MolarEntropyJoulePerMoleKelvin MolarEntropyUnits = "JoulePerMoleKelvin"
        // 
        MolarEntropyKilojoulePerMoleKelvin MolarEntropyUnits = "KilojoulePerMoleKelvin"
        // 
        MolarEntropyMegajoulePerMoleKelvin MolarEntropyUnits = "MegajoulePerMoleKelvin"
)

var internalMolarEntropyUnitsMap = map[MolarEntropyUnits]bool{
	
	MolarEntropyJoulePerMoleKelvin: true,
	MolarEntropyKilojoulePerMoleKelvin: true,
	MolarEntropyMegajoulePerMoleKelvin: true,
}

// MolarEntropyDto represents a MolarEntropy measurement with a numerical value and its corresponding unit.
type MolarEntropyDto struct {
    // Value is the numerical representation of the MolarEntropy.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the MolarEntropy, as defined in the MolarEntropyUnits enumeration.
	Unit  MolarEntropyUnits `json:"unit" validate:"required,oneof=JoulePerMoleKelvin KilojoulePerMoleKelvin MegajoulePerMoleKelvin"`
}

// MolarEntropyDtoFactory groups methods for creating and serializing MolarEntropyDto objects.
type MolarEntropyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolarEntropyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolarEntropyDtoFactory) FromJSON(data []byte) (*MolarEntropyDto, error) {
	a := MolarEntropyDto{}

    // Parse JSON into MolarEntropyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MolarEntropyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolarEntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MolarEntropy represents a measurement in a of MolarEntropy.
//
// Molar entropy is amount of energy required to increase temperature of 1 mole substance by 1 Kelvin.
type MolarEntropy struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_mole_kelvinLazy *float64 
    kilojoules_per_mole_kelvinLazy *float64 
    megajoules_per_mole_kelvinLazy *float64 
}

// MolarEntropyFactory groups methods for creating MolarEntropy instances.
type MolarEntropyFactory struct{}

// CreateMolarEntropy creates a new MolarEntropy instance from the given value and unit.
func (uf MolarEntropyFactory) CreateMolarEntropy(value float64, unit MolarEntropyUnits) (*MolarEntropy, error) {
	return newMolarEntropy(value, unit)
}

// FromDto converts a MolarEntropyDto to a MolarEntropy instance.
func (uf MolarEntropyFactory) FromDto(dto MolarEntropyDto) (*MolarEntropy, error) {
	return newMolarEntropy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MolarEntropy instance.
func (uf MolarEntropyFactory) FromDtoJSON(data []byte) (*MolarEntropy, error) {
	unitDto, err := MolarEntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolarEntropyDto from JSON: %w", err)
	}
	return MolarEntropyFactory{}.FromDto(*unitDto)
}


// FromJoulesPerMoleKelvin creates a new MolarEntropy instance from a value in JoulesPerMoleKelvin.
func (uf MolarEntropyFactory) FromJoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyJoulePerMoleKelvin)
}

// FromKilojoulesPerMoleKelvin creates a new MolarEntropy instance from a value in KilojoulesPerMoleKelvin.
func (uf MolarEntropyFactory) FromKilojoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyKilojoulePerMoleKelvin)
}

// FromMegajoulesPerMoleKelvin creates a new MolarEntropy instance from a value in MegajoulesPerMoleKelvin.
func (uf MolarEntropyFactory) FromMegajoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyMegajoulePerMoleKelvin)
}


// newMolarEntropy creates a new MolarEntropy.
func newMolarEntropy(value float64, fromUnit MolarEntropyUnits) (*MolarEntropy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMolarEntropyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MolarEntropyUnits", fromUnit)
	}
	a := &MolarEntropy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarEntropy in JoulePerMoleKelvin unit (the base/default quantity).
func (a *MolarEntropy) BaseValue() float64 {
	return a.value
}


// JoulesPerMoleKelvin returns the MolarEntropy value in JoulesPerMoleKelvin.
//
// 
func (a *MolarEntropy) JoulesPerMoleKelvin() float64 {
	if a.joules_per_mole_kelvinLazy != nil {
		return *a.joules_per_mole_kelvinLazy
	}
	joules_per_mole_kelvin := a.convertFromBase(MolarEntropyJoulePerMoleKelvin)
	a.joules_per_mole_kelvinLazy = &joules_per_mole_kelvin
	return joules_per_mole_kelvin
}

// KilojoulesPerMoleKelvin returns the MolarEntropy value in KilojoulesPerMoleKelvin.
//
// 
func (a *MolarEntropy) KilojoulesPerMoleKelvin() float64 {
	if a.kilojoules_per_mole_kelvinLazy != nil {
		return *a.kilojoules_per_mole_kelvinLazy
	}
	kilojoules_per_mole_kelvin := a.convertFromBase(MolarEntropyKilojoulePerMoleKelvin)
	a.kilojoules_per_mole_kelvinLazy = &kilojoules_per_mole_kelvin
	return kilojoules_per_mole_kelvin
}

// MegajoulesPerMoleKelvin returns the MolarEntropy value in MegajoulesPerMoleKelvin.
//
// 
func (a *MolarEntropy) MegajoulesPerMoleKelvin() float64 {
	if a.megajoules_per_mole_kelvinLazy != nil {
		return *a.megajoules_per_mole_kelvinLazy
	}
	megajoules_per_mole_kelvin := a.convertFromBase(MolarEntropyMegajoulePerMoleKelvin)
	a.megajoules_per_mole_kelvinLazy = &megajoules_per_mole_kelvin
	return megajoules_per_mole_kelvin
}


// ToDto creates a MolarEntropyDto representation from the MolarEntropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerMoleKelvin by default.
func (a *MolarEntropy) ToDto(holdInUnit *MolarEntropyUnits) MolarEntropyDto {
	if holdInUnit == nil {
		defaultUnit := MolarEntropyJoulePerMoleKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return MolarEntropyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MolarEntropy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerMoleKelvin by default.
func (a *MolarEntropy) ToDtoJSON(holdInUnit *MolarEntropyUnits) ([]byte, error) {
	// Convert to MolarEntropyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MolarEntropy to a specific unit value.
// The function uses the provided unit type (MolarEntropyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MolarEntropy) Convert(toUnit MolarEntropyUnits) float64 {
	switch toUnit { 
    case MolarEntropyJoulePerMoleKelvin:
		return a.JoulesPerMoleKelvin()
    case MolarEntropyKilojoulePerMoleKelvin:
		return a.KilojoulesPerMoleKelvin()
    case MolarEntropyMegajoulePerMoleKelvin:
		return a.MegajoulesPerMoleKelvin()
	default:
		return math.NaN()
	}
}

func (a *MolarEntropy) convertFromBase(toUnit MolarEntropyUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarEntropyJoulePerMoleKelvin:
		return (value) 
	case MolarEntropyKilojoulePerMoleKelvin:
		return ((value) / 1000.0) 
	case MolarEntropyMegajoulePerMoleKelvin:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *MolarEntropy) convertToBase(value float64, fromUnit MolarEntropyUnits) float64 {
	switch fromUnit { 
	case MolarEntropyJoulePerMoleKelvin:
		return (value) 
	case MolarEntropyKilojoulePerMoleKelvin:
		return ((value) * 1000.0) 
	case MolarEntropyMegajoulePerMoleKelvin:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MolarEntropy in the default unit (JoulePerMoleKelvin),
// formatted to two decimal places.
func (a MolarEntropy) String() string {
	return a.ToString(MolarEntropyJoulePerMoleKelvin, 2)
}

// ToString formats the MolarEntropy value as a string with the specified unit and fractional digits.
// It converts the MolarEntropy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MolarEntropy value will be converted (e.g., JoulePerMoleKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MolarEntropy with the unit abbreviation.
func (a *MolarEntropy) ToString(unit MolarEntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMolarEntropyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMolarEntropyAbbreviation(unit))
}

// Equals checks if the given MolarEntropy is equal to the current MolarEntropy.
//
// Parameters:
//    other: The MolarEntropy to compare against.
//
// Returns:
//    bool: Returns true if both MolarEntropy are equal, false otherwise.
func (a *MolarEntropy) Equals(other *MolarEntropy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MolarEntropy with another MolarEntropy.
// It returns -1 if the current MolarEntropy is less than the other MolarEntropy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MolarEntropy to compare against.
//
// Returns:
//    int: -1 if the current MolarEntropy is less, 1 if greater, and 0 if equal.
func (a *MolarEntropy) CompareTo(other *MolarEntropy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MolarEntropy to the current MolarEntropy and returns the result.
// The result is a new MolarEntropy instance with the sum of the values.
//
// Parameters:
//    other: The MolarEntropy to add to the current MolarEntropy.
//
// Returns:
//    *MolarEntropy: A new MolarEntropy instance representing the sum of both MolarEntropy.
func (a *MolarEntropy) Add(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MolarEntropy from the current MolarEntropy and returns the result.
// The result is a new MolarEntropy instance with the difference of the values.
//
// Parameters:
//    other: The MolarEntropy to subtract from the current MolarEntropy.
//
// Returns:
//    *MolarEntropy: A new MolarEntropy instance representing the difference of both MolarEntropy.
func (a *MolarEntropy) Subtract(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MolarEntropy by the given MolarEntropy and returns the result.
// The result is a new MolarEntropy instance with the product of the values.
//
// Parameters:
//    other: The MolarEntropy to multiply with the current MolarEntropy.
//
// Returns:
//    *MolarEntropy: A new MolarEntropy instance representing the product of both MolarEntropy.
func (a *MolarEntropy) Multiply(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value * other.BaseValue()}
}

// Divide divides the current MolarEntropy by the given MolarEntropy and returns the result.
// The result is a new MolarEntropy instance with the quotient of the values.
//
// Parameters:
//    other: The MolarEntropy to divide the current MolarEntropy by.
//
// Returns:
//    *MolarEntropy: A new MolarEntropy instance representing the quotient of both MolarEntropy.
func (a *MolarEntropy) Divide(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value / other.BaseValue()}
}

// GetMolarEntropyAbbreviation gets the unit abbreviation.
func GetMolarEntropyAbbreviation(unit MolarEntropyUnits) string {
	switch unit { 
	case MolarEntropyJoulePerMoleKelvin:
		return "J/(mol·K)" 
	case MolarEntropyKilojoulePerMoleKelvin:
		return "kJ/(mol·K)" 
	case MolarEntropyMegajoulePerMoleKelvin:
		return "MJ/(mol·K)" 
	default:
		return ""
	}
}