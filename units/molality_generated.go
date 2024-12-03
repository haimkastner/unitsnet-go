// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolalityUnits enumeration
type MolalityUnits string

const (
    
        // 
        MolalityMolePerKilogram MolalityUnits = "MolePerKilogram"
        // 
        MolalityMolePerGram MolalityUnits = "MolePerGram"
        // 
        MolalityMillimolePerKilogram MolalityUnits = "MillimolePerKilogram"
)

// MolalityDto represents an Molality
type MolalityDto struct {
	Value float64
	Unit  MolalityUnits
}

// MolalityDtoFactory struct to group related functions
type MolalityDtoFactory struct{}

func (udf MolalityDtoFactory) FromJSON(data []byte) (*MolalityDto, error) {
	a := MolalityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolalityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Molality struct
type Molality struct {
	value       float64
    
    moles_per_kilogramLazy *float64 
    moles_per_gramLazy *float64 
    millimoles_per_kilogramLazy *float64 
}

// MolalityFactory struct to group related functions
type MolalityFactory struct{}

func (uf MolalityFactory) CreateMolality(value float64, unit MolalityUnits) (*Molality, error) {
	return newMolality(value, unit)
}

func (uf MolalityFactory) FromDto(dto MolalityDto) (*Molality, error) {
	return newMolality(dto.Value, dto.Unit)
}

func (uf MolalityFactory) FromDtoJSON(data []byte) (*Molality, error) {
	unitDto, err := MolalityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolalityFactory{}.FromDto(*unitDto)
}


// FromMolePerKilogram creates a new Molality instance from MolePerKilogram.
func (uf MolalityFactory) FromMolesPerKilogram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMolePerKilogram)
}

// FromMolePerGram creates a new Molality instance from MolePerGram.
func (uf MolalityFactory) FromMolesPerGram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMolePerGram)
}

// FromMillimolePerKilogram creates a new Molality instance from MillimolePerKilogram.
func (uf MolalityFactory) FromMillimolesPerKilogram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMillimolePerKilogram)
}




// newMolality creates a new Molality.
func newMolality(value float64, fromUnit MolalityUnits) (*Molality, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Molality{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Molality in MolePerKilogram.
func (a *Molality) BaseValue() float64 {
	return a.value
}


// MolePerKilogram returns the value in MolePerKilogram.
func (a *Molality) MolesPerKilogram() float64 {
	if a.moles_per_kilogramLazy != nil {
		return *a.moles_per_kilogramLazy
	}
	moles_per_kilogram := a.convertFromBase(MolalityMolePerKilogram)
	a.moles_per_kilogramLazy = &moles_per_kilogram
	return moles_per_kilogram
}

// MolePerGram returns the value in MolePerGram.
func (a *Molality) MolesPerGram() float64 {
	if a.moles_per_gramLazy != nil {
		return *a.moles_per_gramLazy
	}
	moles_per_gram := a.convertFromBase(MolalityMolePerGram)
	a.moles_per_gramLazy = &moles_per_gram
	return moles_per_gram
}

// MillimolePerKilogram returns the value in MillimolePerKilogram.
func (a *Molality) MillimolesPerKilogram() float64 {
	if a.millimoles_per_kilogramLazy != nil {
		return *a.millimoles_per_kilogramLazy
	}
	millimoles_per_kilogram := a.convertFromBase(MolalityMillimolePerKilogram)
	a.millimoles_per_kilogramLazy = &millimoles_per_kilogram
	return millimoles_per_kilogram
}


// ToDto creates an MolalityDto representation.
func (a *Molality) ToDto(holdInUnit *MolalityUnits) MolalityDto {
	if holdInUnit == nil {
		defaultUnit := MolalityMolePerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return MolalityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MolalityDto representation.
func (a *Molality) ToDtoJSON(holdInUnit *MolalityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Molality to a specific unit value.
func (a *Molality) Convert(toUnit MolalityUnits) float64 {
	switch toUnit { 
    case MolalityMolePerKilogram:
		return a.MolesPerKilogram()
    case MolalityMolePerGram:
		return a.MolesPerGram()
    case MolalityMillimolePerKilogram:
		return a.MillimolesPerKilogram()
	default:
		return 0
	}
}

func (a *Molality) convertFromBase(toUnit MolalityUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolalityMolePerKilogram:
		return (value) 
	case MolalityMolePerGram:
		return (value * 1e-3) 
	case MolalityMillimolePerKilogram:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Molality) convertToBase(value float64, fromUnit MolalityUnits) float64 {
	switch fromUnit { 
	case MolalityMolePerKilogram:
		return (value) 
	case MolalityMolePerGram:
		return (value / 1e-3) 
	case MolalityMillimolePerKilogram:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Molality) String() string {
	return a.ToString(MolalityMolePerKilogram, 2)
}

// ToString formats the Molality to string.
// fractionalDigits -1 for not mention
func (a *Molality) ToString(unit MolalityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Molality) getUnitAbbreviation(unit MolalityUnits) string {
	switch unit { 
	case MolalityMolePerKilogram:
		return "mol/kg" 
	case MolalityMolePerGram:
		return "mol/g" 
	case MolalityMillimolePerKilogram:
		return "mmol/kg" 
	default:
		return ""
	}
}

// Check if the given Molality are equals to the current Molality
func (a *Molality) Equals(other *Molality) bool {
	return a.value == other.BaseValue()
}

// Check if the given Molality are equals to the current Molality
func (a *Molality) CompareTo(other *Molality) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Molality to the current Molality.
func (a *Molality) Add(other *Molality) *Molality {
	return &Molality{value: a.value + other.BaseValue()}
}

// Subtract the given Molality to the current Molality.
func (a *Molality) Subtract(other *Molality) *Molality {
	return &Molality{value: a.value - other.BaseValue()}
}

// Multiply the given Molality to the current Molality.
func (a *Molality) Multiply(other *Molality) *Molality {
	return &Molality{value: a.value * other.BaseValue()}
}

// Divide the given Molality to the current Molality.
func (a *Molality) Divide(other *Molality) *Molality {
	return &Molality{value: a.value / other.BaseValue()}
}