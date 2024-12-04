// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolarEnergyUnits defines various units of MolarEnergy.
type MolarEnergyUnits string

const (
    
        // 
        MolarEnergyJoulePerMole MolarEnergyUnits = "JoulePerMole"
        // 
        MolarEnergyKilojoulePerMole MolarEnergyUnits = "KilojoulePerMole"
        // 
        MolarEnergyMegajoulePerMole MolarEnergyUnits = "MegajoulePerMole"
)

// MolarEnergyDto represents a MolarEnergy measurement with a numerical value and its corresponding unit.
type MolarEnergyDto struct {
    // Value is the numerical representation of the MolarEnergy.
	Value float64
    // Unit specifies the unit of measurement for the MolarEnergy, as defined in the MolarEnergyUnits enumeration.
	Unit  MolarEnergyUnits
}

// MolarEnergyDtoFactory groups methods for creating and serializing MolarEnergyDto objects.
type MolarEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolarEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolarEnergyDtoFactory) FromJSON(data []byte) (*MolarEnergyDto, error) {
	a := MolarEnergyDto{}

    // Parse JSON into MolarEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MolarEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolarEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// MolarEnergy represents a measurement in a of MolarEnergy.
//
// Molar energy is the amount of energy stored in 1 mole of a substance.
type MolarEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_moleLazy *float64 
    kilojoules_per_moleLazy *float64 
    megajoules_per_moleLazy *float64 
}

// MolarEnergyFactory groups methods for creating MolarEnergy instances.
type MolarEnergyFactory struct{}

// CreateMolarEnergy creates a new MolarEnergy instance from the given value and unit.
func (uf MolarEnergyFactory) CreateMolarEnergy(value float64, unit MolarEnergyUnits) (*MolarEnergy, error) {
	return newMolarEnergy(value, unit)
}

// FromDto converts a MolarEnergyDto to a MolarEnergy instance.
func (uf MolarEnergyFactory) FromDto(dto MolarEnergyDto) (*MolarEnergy, error) {
	return newMolarEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MolarEnergy instance.
func (uf MolarEnergyFactory) FromDtoJSON(data []byte) (*MolarEnergy, error) {
	unitDto, err := MolarEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolarEnergyDto from JSON: %w", err)
	}
	return MolarEnergyFactory{}.FromDto(*unitDto)
}


// FromJoulesPerMole creates a new MolarEnergy instance from a value in JoulesPerMole.
func (uf MolarEnergyFactory) FromJoulesPerMole(value float64) (*MolarEnergy, error) {
	return newMolarEnergy(value, MolarEnergyJoulePerMole)
}

// FromKilojoulesPerMole creates a new MolarEnergy instance from a value in KilojoulesPerMole.
func (uf MolarEnergyFactory) FromKilojoulesPerMole(value float64) (*MolarEnergy, error) {
	return newMolarEnergy(value, MolarEnergyKilojoulePerMole)
}

// FromMegajoulesPerMole creates a new MolarEnergy instance from a value in MegajoulesPerMole.
func (uf MolarEnergyFactory) FromMegajoulesPerMole(value float64) (*MolarEnergy, error) {
	return newMolarEnergy(value, MolarEnergyMegajoulePerMole)
}


// newMolarEnergy creates a new MolarEnergy.
func newMolarEnergy(value float64, fromUnit MolarEnergyUnits) (*MolarEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MolarEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarEnergy in JoulePerMole unit (the base/default quantity).
func (a *MolarEnergy) BaseValue() float64 {
	return a.value
}


// JoulesPerMole returns the MolarEnergy value in JoulesPerMole.
//
// 
func (a *MolarEnergy) JoulesPerMole() float64 {
	if a.joules_per_moleLazy != nil {
		return *a.joules_per_moleLazy
	}
	joules_per_mole := a.convertFromBase(MolarEnergyJoulePerMole)
	a.joules_per_moleLazy = &joules_per_mole
	return joules_per_mole
}

// KilojoulesPerMole returns the MolarEnergy value in KilojoulesPerMole.
//
// 
func (a *MolarEnergy) KilojoulesPerMole() float64 {
	if a.kilojoules_per_moleLazy != nil {
		return *a.kilojoules_per_moleLazy
	}
	kilojoules_per_mole := a.convertFromBase(MolarEnergyKilojoulePerMole)
	a.kilojoules_per_moleLazy = &kilojoules_per_mole
	return kilojoules_per_mole
}

// MegajoulesPerMole returns the MolarEnergy value in MegajoulesPerMole.
//
// 
func (a *MolarEnergy) MegajoulesPerMole() float64 {
	if a.megajoules_per_moleLazy != nil {
		return *a.megajoules_per_moleLazy
	}
	megajoules_per_mole := a.convertFromBase(MolarEnergyMegajoulePerMole)
	a.megajoules_per_moleLazy = &megajoules_per_mole
	return megajoules_per_mole
}


// ToDto creates a MolarEnergyDto representation from the MolarEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerMole by default.
func (a *MolarEnergy) ToDto(holdInUnit *MolarEnergyUnits) MolarEnergyDto {
	if holdInUnit == nil {
		defaultUnit := MolarEnergyJoulePerMole // Default value
		holdInUnit = &defaultUnit
	}

	return MolarEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MolarEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerMole by default.
func (a *MolarEnergy) ToDtoJSON(holdInUnit *MolarEnergyUnits) ([]byte, error) {
	// Convert to MolarEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MolarEnergy to a specific unit value.
// The function uses the provided unit type (MolarEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MolarEnergy) Convert(toUnit MolarEnergyUnits) float64 {
	switch toUnit { 
    case MolarEnergyJoulePerMole:
		return a.JoulesPerMole()
    case MolarEnergyKilojoulePerMole:
		return a.KilojoulesPerMole()
    case MolarEnergyMegajoulePerMole:
		return a.MegajoulesPerMole()
	default:
		return math.NaN()
	}
}

func (a *MolarEnergy) convertFromBase(toUnit MolarEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarEnergyJoulePerMole:
		return (value) 
	case MolarEnergyKilojoulePerMole:
		return ((value) / 1000.0) 
	case MolarEnergyMegajoulePerMole:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *MolarEnergy) convertToBase(value float64, fromUnit MolarEnergyUnits) float64 {
	switch fromUnit { 
	case MolarEnergyJoulePerMole:
		return (value) 
	case MolarEnergyKilojoulePerMole:
		return ((value) * 1000.0) 
	case MolarEnergyMegajoulePerMole:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MolarEnergy in the default unit (JoulePerMole),
// formatted to two decimal places.
func (a MolarEnergy) String() string {
	return a.ToString(MolarEnergyJoulePerMole, 2)
}

// ToString formats the MolarEnergy value as a string with the specified unit and fractional digits.
// It converts the MolarEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MolarEnergy value will be converted (e.g., JoulePerMole))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MolarEnergy with the unit abbreviation.
func (a *MolarEnergy) ToString(unit MolarEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MolarEnergy) getUnitAbbreviation(unit MolarEnergyUnits) string {
	switch unit { 
	case MolarEnergyJoulePerMole:
		return "J/mol" 
	case MolarEnergyKilojoulePerMole:
		return "kJ/mol" 
	case MolarEnergyMegajoulePerMole:
		return "MJ/mol" 
	default:
		return ""
	}
}

// Equals checks if the given MolarEnergy is equal to the current MolarEnergy.
//
// Parameters:
//    other: The MolarEnergy to compare against.
//
// Returns:
//    bool: Returns true if both MolarEnergy are equal, false otherwise.
func (a *MolarEnergy) Equals(other *MolarEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MolarEnergy with another MolarEnergy.
// It returns -1 if the current MolarEnergy is less than the other MolarEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MolarEnergy to compare against.
//
// Returns:
//    int: -1 if the current MolarEnergy is less, 1 if greater, and 0 if equal.
func (a *MolarEnergy) CompareTo(other *MolarEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MolarEnergy to the current MolarEnergy and returns the result.
// The result is a new MolarEnergy instance with the sum of the values.
//
// Parameters:
//    other: The MolarEnergy to add to the current MolarEnergy.
//
// Returns:
//    *MolarEnergy: A new MolarEnergy instance representing the sum of both MolarEnergy.
func (a *MolarEnergy) Add(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MolarEnergy from the current MolarEnergy and returns the result.
// The result is a new MolarEnergy instance with the difference of the values.
//
// Parameters:
//    other: The MolarEnergy to subtract from the current MolarEnergy.
//
// Returns:
//    *MolarEnergy: A new MolarEnergy instance representing the difference of both MolarEnergy.
func (a *MolarEnergy) Subtract(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MolarEnergy by the given MolarEnergy and returns the result.
// The result is a new MolarEnergy instance with the product of the values.
//
// Parameters:
//    other: The MolarEnergy to multiply with the current MolarEnergy.
//
// Returns:
//    *MolarEnergy: A new MolarEnergy instance representing the product of both MolarEnergy.
func (a *MolarEnergy) Multiply(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current MolarEnergy by the given MolarEnergy and returns the result.
// The result is a new MolarEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The MolarEnergy to divide the current MolarEnergy by.
//
// Returns:
//    *MolarEnergy: A new MolarEnergy instance representing the quotient of both MolarEnergy.
func (a *MolarEnergy) Divide(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value / other.BaseValue()}
}