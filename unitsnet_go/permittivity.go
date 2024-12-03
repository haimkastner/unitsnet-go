package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PermittivityUnits enumeration
type PermittivityUnits string

const (
    
        // 
        PermittivityFaradPerMeter PermittivityUnits = "FaradPerMeter"
)

// PermittivityDto represents an Permittivity
type PermittivityDto struct {
	Value float64
	Unit  PermittivityUnits
}

// PermittivityDtoFactory struct to group related functions
type PermittivityDtoFactory struct{}

func (udf PermittivityDtoFactory) FromJSON(data []byte) (*PermittivityDto, error) {
	a := PermittivityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PermittivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Permittivity struct
type Permittivity struct {
	value       float64
    
    farads_per_meterLazy *float64 
}

// PermittivityFactory struct to group related functions
type PermittivityFactory struct{}

func (uf PermittivityFactory) CreatePermittivity(value float64, unit PermittivityUnits) (*Permittivity, error) {
	return newPermittivity(value, unit)
}

func (uf PermittivityFactory) FromDto(dto PermittivityDto) (*Permittivity, error) {
	return newPermittivity(dto.Value, dto.Unit)
}

func (uf PermittivityFactory) FromDtoJSON(data []byte) (*Permittivity, error) {
	unitDto, err := PermittivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PermittivityFactory{}.FromDto(*unitDto)
}


// FromFaradPerMeter creates a new Permittivity instance from FaradPerMeter.
func (uf PermittivityFactory) FromFaradsPerMeter(value float64) (*Permittivity, error) {
	return newPermittivity(value, PermittivityFaradPerMeter)
}




// newPermittivity creates a new Permittivity.
func newPermittivity(value float64, fromUnit PermittivityUnits) (*Permittivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Permittivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Permittivity in FaradPerMeter.
func (a *Permittivity) BaseValue() float64 {
	return a.value
}


// FaradPerMeter returns the value in FaradPerMeter.
func (a *Permittivity) FaradsPerMeter() float64 {
	if a.farads_per_meterLazy != nil {
		return *a.farads_per_meterLazy
	}
	farads_per_meter := a.convertFromBase(PermittivityFaradPerMeter)
	a.farads_per_meterLazy = &farads_per_meter
	return farads_per_meter
}


// ToDto creates an PermittivityDto representation.
func (a *Permittivity) ToDto(holdInUnit *PermittivityUnits) PermittivityDto {
	if holdInUnit == nil {
		defaultUnit := PermittivityFaradPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PermittivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PermittivityDto representation.
func (a *Permittivity) ToDtoJSON(holdInUnit *PermittivityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Permittivity to a specific unit value.
func (a *Permittivity) Convert(toUnit PermittivityUnits) float64 {
	switch toUnit { 
    case PermittivityFaradPerMeter:
		return a.FaradsPerMeter()
	default:
		return 0
	}
}

func (a *Permittivity) convertFromBase(toUnit PermittivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PermittivityFaradPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Permittivity) convertToBase(value float64, fromUnit PermittivityUnits) float64 {
	switch fromUnit { 
	case PermittivityFaradPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Permittivity) String() string {
	return a.ToString(PermittivityFaradPerMeter, 2)
}

// ToString formats the Permittivity to string.
// fractionalDigits -1 for not mention
func (a *Permittivity) ToString(unit PermittivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Permittivity) getUnitAbbreviation(unit PermittivityUnits) string {
	switch unit { 
	case PermittivityFaradPerMeter:
		return "F/m" 
	default:
		return ""
	}
}

// Check if the given Permittivity are equals to the current Permittivity
func (a *Permittivity) Equals(other *Permittivity) bool {
	return a.value == other.BaseValue()
}

// Check if the given Permittivity are equals to the current Permittivity
func (a *Permittivity) CompareTo(other *Permittivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Permittivity to the current Permittivity.
func (a *Permittivity) Add(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value + other.BaseValue()}
}

// Subtract the given Permittivity to the current Permittivity.
func (a *Permittivity) Subtract(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value - other.BaseValue()}
}

// Multiply the given Permittivity to the current Permittivity.
func (a *Permittivity) Multiply(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value * other.BaseValue()}
}

// Divide the given Permittivity to the current Permittivity.
func (a *Permittivity) Divide(other *Permittivity) *Permittivity {
	return &Permittivity{value: a.value / other.BaseValue()}
}