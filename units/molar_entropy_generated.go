// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolarEntropyUnits enumeration
type MolarEntropyUnits string

const (
    
        // 
        MolarEntropyJoulePerMoleKelvin MolarEntropyUnits = "JoulePerMoleKelvin"
        // 
        MolarEntropyKilojoulePerMoleKelvin MolarEntropyUnits = "KilojoulePerMoleKelvin"
        // 
        MolarEntropyMegajoulePerMoleKelvin MolarEntropyUnits = "MegajoulePerMoleKelvin"
)

// MolarEntropyDto represents an MolarEntropy
type MolarEntropyDto struct {
	Value float64
	Unit  MolarEntropyUnits
}

// MolarEntropyDtoFactory struct to group related functions
type MolarEntropyDtoFactory struct{}

func (udf MolarEntropyDtoFactory) FromJSON(data []byte) (*MolarEntropyDto, error) {
	a := MolarEntropyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolarEntropyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MolarEntropy struct
type MolarEntropy struct {
	value       float64
    
    joules_per_mole_kelvinLazy *float64 
    kilojoules_per_mole_kelvinLazy *float64 
    megajoules_per_mole_kelvinLazy *float64 
}

// MolarEntropyFactory struct to group related functions
type MolarEntropyFactory struct{}

func (uf MolarEntropyFactory) CreateMolarEntropy(value float64, unit MolarEntropyUnits) (*MolarEntropy, error) {
	return newMolarEntropy(value, unit)
}

func (uf MolarEntropyFactory) FromDto(dto MolarEntropyDto) (*MolarEntropy, error) {
	return newMolarEntropy(dto.Value, dto.Unit)
}

func (uf MolarEntropyFactory) FromDtoJSON(data []byte) (*MolarEntropy, error) {
	unitDto, err := MolarEntropyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolarEntropyFactory{}.FromDto(*unitDto)
}


// FromJoulePerMoleKelvin creates a new MolarEntropy instance from JoulePerMoleKelvin.
func (uf MolarEntropyFactory) FromJoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyJoulePerMoleKelvin)
}

// FromKilojoulePerMoleKelvin creates a new MolarEntropy instance from KilojoulePerMoleKelvin.
func (uf MolarEntropyFactory) FromKilojoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyKilojoulePerMoleKelvin)
}

// FromMegajoulePerMoleKelvin creates a new MolarEntropy instance from MegajoulePerMoleKelvin.
func (uf MolarEntropyFactory) FromMegajoulesPerMoleKelvin(value float64) (*MolarEntropy, error) {
	return newMolarEntropy(value, MolarEntropyMegajoulePerMoleKelvin)
}




// newMolarEntropy creates a new MolarEntropy.
func newMolarEntropy(value float64, fromUnit MolarEntropyUnits) (*MolarEntropy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MolarEntropy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarEntropy in JoulePerMoleKelvin.
func (a *MolarEntropy) BaseValue() float64 {
	return a.value
}


// JoulePerMoleKelvin returns the value in JoulePerMoleKelvin.
func (a *MolarEntropy) JoulesPerMoleKelvin() float64 {
	if a.joules_per_mole_kelvinLazy != nil {
		return *a.joules_per_mole_kelvinLazy
	}
	joules_per_mole_kelvin := a.convertFromBase(MolarEntropyJoulePerMoleKelvin)
	a.joules_per_mole_kelvinLazy = &joules_per_mole_kelvin
	return joules_per_mole_kelvin
}

// KilojoulePerMoleKelvin returns the value in KilojoulePerMoleKelvin.
func (a *MolarEntropy) KilojoulesPerMoleKelvin() float64 {
	if a.kilojoules_per_mole_kelvinLazy != nil {
		return *a.kilojoules_per_mole_kelvinLazy
	}
	kilojoules_per_mole_kelvin := a.convertFromBase(MolarEntropyKilojoulePerMoleKelvin)
	a.kilojoules_per_mole_kelvinLazy = &kilojoules_per_mole_kelvin
	return kilojoules_per_mole_kelvin
}

// MegajoulePerMoleKelvin returns the value in MegajoulePerMoleKelvin.
func (a *MolarEntropy) MegajoulesPerMoleKelvin() float64 {
	if a.megajoules_per_mole_kelvinLazy != nil {
		return *a.megajoules_per_mole_kelvinLazy
	}
	megajoules_per_mole_kelvin := a.convertFromBase(MolarEntropyMegajoulePerMoleKelvin)
	a.megajoules_per_mole_kelvinLazy = &megajoules_per_mole_kelvin
	return megajoules_per_mole_kelvin
}


// ToDto creates an MolarEntropyDto representation.
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

// ToDtoJSON creates an MolarEntropyDto representation.
func (a *MolarEntropy) ToDtoJSON(holdInUnit *MolarEntropyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MolarEntropy to a specific unit value.
func (a *MolarEntropy) Convert(toUnit MolarEntropyUnits) float64 {
	switch toUnit { 
    case MolarEntropyJoulePerMoleKelvin:
		return a.JoulesPerMoleKelvin()
    case MolarEntropyKilojoulePerMoleKelvin:
		return a.KilojoulesPerMoleKelvin()
    case MolarEntropyMegajoulePerMoleKelvin:
		return a.MegajoulesPerMoleKelvin()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a MolarEntropy) String() string {
	return a.ToString(MolarEntropyJoulePerMoleKelvin, 2)
}

// ToString formats the MolarEntropy to string.
// fractionalDigits -1 for not mention
func (a *MolarEntropy) ToString(unit MolarEntropyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MolarEntropy) getUnitAbbreviation(unit MolarEntropyUnits) string {
	switch unit { 
	case MolarEntropyJoulePerMoleKelvin:
		return "J/(mol*K)" 
	case MolarEntropyKilojoulePerMoleKelvin:
		return "kJ/(mol*K)" 
	case MolarEntropyMegajoulePerMoleKelvin:
		return "MJ/(mol*K)" 
	default:
		return ""
	}
}

// Check if the given MolarEntropy are equals to the current MolarEntropy
func (a *MolarEntropy) Equals(other *MolarEntropy) bool {
	return a.value == other.BaseValue()
}

// Check if the given MolarEntropy are equals to the current MolarEntropy
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

// Add the given MolarEntropy to the current MolarEntropy.
func (a *MolarEntropy) Add(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value + other.BaseValue()}
}

// Subtract the given MolarEntropy to the current MolarEntropy.
func (a *MolarEntropy) Subtract(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value - other.BaseValue()}
}

// Multiply the given MolarEntropy to the current MolarEntropy.
func (a *MolarEntropy) Multiply(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value * other.BaseValue()}
}

// Divide the given MolarEntropy to the current MolarEntropy.
func (a *MolarEntropy) Divide(other *MolarEntropy) *MolarEntropy {
	return &MolarEntropy{value: a.value / other.BaseValue()}
}