// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PermeabilityUnits enumeration
type PermeabilityUnits string

const (
    
        // 
        PermeabilityHenryPerMeter PermeabilityUnits = "HenryPerMeter"
)

// PermeabilityDto represents an Permeability
type PermeabilityDto struct {
	Value float64
	Unit  PermeabilityUnits
}

// PermeabilityDtoFactory struct to group related functions
type PermeabilityDtoFactory struct{}

func (udf PermeabilityDtoFactory) FromJSON(data []byte) (*PermeabilityDto, error) {
	a := PermeabilityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PermeabilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Permeability struct
type Permeability struct {
	value       float64
    
    henries_per_meterLazy *float64 
}

// PermeabilityFactory struct to group related functions
type PermeabilityFactory struct{}

func (uf PermeabilityFactory) CreatePermeability(value float64, unit PermeabilityUnits) (*Permeability, error) {
	return newPermeability(value, unit)
}

func (uf PermeabilityFactory) FromDto(dto PermeabilityDto) (*Permeability, error) {
	return newPermeability(dto.Value, dto.Unit)
}

func (uf PermeabilityFactory) FromDtoJSON(data []byte) (*Permeability, error) {
	unitDto, err := PermeabilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PermeabilityFactory{}.FromDto(*unitDto)
}


// FromHenryPerMeter creates a new Permeability instance from HenryPerMeter.
func (uf PermeabilityFactory) FromHenriesPerMeter(value float64) (*Permeability, error) {
	return newPermeability(value, PermeabilityHenryPerMeter)
}




// newPermeability creates a new Permeability.
func newPermeability(value float64, fromUnit PermeabilityUnits) (*Permeability, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Permeability{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Permeability in HenryPerMeter.
func (a *Permeability) BaseValue() float64 {
	return a.value
}


// HenryPerMeter returns the value in HenryPerMeter.
func (a *Permeability) HenriesPerMeter() float64 {
	if a.henries_per_meterLazy != nil {
		return *a.henries_per_meterLazy
	}
	henries_per_meter := a.convertFromBase(PermeabilityHenryPerMeter)
	a.henries_per_meterLazy = &henries_per_meter
	return henries_per_meter
}


// ToDto creates an PermeabilityDto representation.
func (a *Permeability) ToDto(holdInUnit *PermeabilityUnits) PermeabilityDto {
	if holdInUnit == nil {
		defaultUnit := PermeabilityHenryPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PermeabilityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PermeabilityDto representation.
func (a *Permeability) ToDtoJSON(holdInUnit *PermeabilityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Permeability to a specific unit value.
func (a *Permeability) Convert(toUnit PermeabilityUnits) float64 {
	switch toUnit { 
    case PermeabilityHenryPerMeter:
		return a.HenriesPerMeter()
	default:
		return 0
	}
}

func (a *Permeability) convertFromBase(toUnit PermeabilityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PermeabilityHenryPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Permeability) convertToBase(value float64, fromUnit PermeabilityUnits) float64 {
	switch fromUnit { 
	case PermeabilityHenryPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Permeability) String() string {
	return a.ToString(PermeabilityHenryPerMeter, 2)
}

// ToString formats the Permeability to string.
// fractionalDigits -1 for not mention
func (a *Permeability) ToString(unit PermeabilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Permeability) getUnitAbbreviation(unit PermeabilityUnits) string {
	switch unit { 
	case PermeabilityHenryPerMeter:
		return "H/m" 
	default:
		return ""
	}
}

// Check if the given Permeability are equals to the current Permeability
func (a *Permeability) Equals(other *Permeability) bool {
	return a.value == other.BaseValue()
}

// Check if the given Permeability are equals to the current Permeability
func (a *Permeability) CompareTo(other *Permeability) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Permeability to the current Permeability.
func (a *Permeability) Add(other *Permeability) *Permeability {
	return &Permeability{value: a.value + other.BaseValue()}
}

// Subtract the given Permeability to the current Permeability.
func (a *Permeability) Subtract(other *Permeability) *Permeability {
	return &Permeability{value: a.value - other.BaseValue()}
}

// Multiply the given Permeability to the current Permeability.
func (a *Permeability) Multiply(other *Permeability) *Permeability {
	return &Permeability{value: a.value * other.BaseValue()}
}

// Divide the given Permeability to the current Permeability.
func (a *Permeability) Divide(other *Permeability) *Permeability {
	return &Permeability{value: a.value / other.BaseValue()}
}