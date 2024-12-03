// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolarEnergyUnits enumeration
type MolarEnergyUnits string

const (
    
        // 
        MolarEnergyJoulePerMole MolarEnergyUnits = "JoulePerMole"
        // 
        MolarEnergyKilojoulePerMole MolarEnergyUnits = "KilojoulePerMole"
        // 
        MolarEnergyMegajoulePerMole MolarEnergyUnits = "MegajoulePerMole"
)

// MolarEnergyDto represents an MolarEnergy
type MolarEnergyDto struct {
	Value float64
	Unit  MolarEnergyUnits
}

// MolarEnergyDtoFactory struct to group related functions
type MolarEnergyDtoFactory struct{}

func (udf MolarEnergyDtoFactory) FromJSON(data []byte) (*MolarEnergyDto, error) {
	a := MolarEnergyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolarEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MolarEnergy struct
type MolarEnergy struct {
	value       float64
    
    joules_per_moleLazy *float64 
    kilojoules_per_moleLazy *float64 
    megajoules_per_moleLazy *float64 
}

// MolarEnergyFactory struct to group related functions
type MolarEnergyFactory struct{}

func (uf MolarEnergyFactory) CreateMolarEnergy(value float64, unit MolarEnergyUnits) (*MolarEnergy, error) {
	return newMolarEnergy(value, unit)
}

func (uf MolarEnergyFactory) FromDto(dto MolarEnergyDto) (*MolarEnergy, error) {
	return newMolarEnergy(dto.Value, dto.Unit)
}

func (uf MolarEnergyFactory) FromDtoJSON(data []byte) (*MolarEnergy, error) {
	unitDto, err := MolarEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolarEnergyFactory{}.FromDto(*unitDto)
}


// FromJoulePerMole creates a new MolarEnergy instance from JoulePerMole.
func (uf MolarEnergyFactory) FromJoulesPerMole(value float64) (*MolarEnergy, error) {
	return newMolarEnergy(value, MolarEnergyJoulePerMole)
}

// FromKilojoulePerMole creates a new MolarEnergy instance from KilojoulePerMole.
func (uf MolarEnergyFactory) FromKilojoulesPerMole(value float64) (*MolarEnergy, error) {
	return newMolarEnergy(value, MolarEnergyKilojoulePerMole)
}

// FromMegajoulePerMole creates a new MolarEnergy instance from MegajoulePerMole.
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

// BaseValue returns the base value of MolarEnergy in JoulePerMole.
func (a *MolarEnergy) BaseValue() float64 {
	return a.value
}


// JoulePerMole returns the value in JoulePerMole.
func (a *MolarEnergy) JoulesPerMole() float64 {
	if a.joules_per_moleLazy != nil {
		return *a.joules_per_moleLazy
	}
	joules_per_mole := a.convertFromBase(MolarEnergyJoulePerMole)
	a.joules_per_moleLazy = &joules_per_mole
	return joules_per_mole
}

// KilojoulePerMole returns the value in KilojoulePerMole.
func (a *MolarEnergy) KilojoulesPerMole() float64 {
	if a.kilojoules_per_moleLazy != nil {
		return *a.kilojoules_per_moleLazy
	}
	kilojoules_per_mole := a.convertFromBase(MolarEnergyKilojoulePerMole)
	a.kilojoules_per_moleLazy = &kilojoules_per_mole
	return kilojoules_per_mole
}

// MegajoulePerMole returns the value in MegajoulePerMole.
func (a *MolarEnergy) MegajoulesPerMole() float64 {
	if a.megajoules_per_moleLazy != nil {
		return *a.megajoules_per_moleLazy
	}
	megajoules_per_mole := a.convertFromBase(MolarEnergyMegajoulePerMole)
	a.megajoules_per_moleLazy = &megajoules_per_mole
	return megajoules_per_mole
}


// ToDto creates an MolarEnergyDto representation.
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

// ToDtoJSON creates an MolarEnergyDto representation.
func (a *MolarEnergy) ToDtoJSON(holdInUnit *MolarEnergyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MolarEnergy to a specific unit value.
func (a *MolarEnergy) Convert(toUnit MolarEnergyUnits) float64 {
	switch toUnit { 
    case MolarEnergyJoulePerMole:
		return a.JoulesPerMole()
    case MolarEnergyKilojoulePerMole:
		return a.KilojoulesPerMole()
    case MolarEnergyMegajoulePerMole:
		return a.MegajoulesPerMole()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a MolarEnergy) String() string {
	return a.ToString(MolarEnergyJoulePerMole, 2)
}

// ToString formats the MolarEnergy to string.
// fractionalDigits -1 for not mention
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

// Check if the given MolarEnergy are equals to the current MolarEnergy
func (a *MolarEnergy) Equals(other *MolarEnergy) bool {
	return a.value == other.BaseValue()
}

// Check if the given MolarEnergy are equals to the current MolarEnergy
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

// Add the given MolarEnergy to the current MolarEnergy.
func (a *MolarEnergy) Add(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value + other.BaseValue()}
}

// Subtract the given MolarEnergy to the current MolarEnergy.
func (a *MolarEnergy) Subtract(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value - other.BaseValue()}
}

// Multiply the given MolarEnergy to the current MolarEnergy.
func (a *MolarEnergy) Multiply(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value * other.BaseValue()}
}

// Divide the given MolarEnergy to the current MolarEnergy.
func (a *MolarEnergy) Divide(other *MolarEnergy) *MolarEnergy {
	return &MolarEnergy{value: a.value / other.BaseValue()}
}