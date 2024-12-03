// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MagneticFluxUnits enumeration
type MagneticFluxUnits string

const (
    
        // 
        MagneticFluxWeber MagneticFluxUnits = "Weber"
)

// MagneticFluxDto represents an MagneticFlux
type MagneticFluxDto struct {
	Value float64
	Unit  MagneticFluxUnits
}

// MagneticFluxDtoFactory struct to group related functions
type MagneticFluxDtoFactory struct{}

func (udf MagneticFluxDtoFactory) FromJSON(data []byte) (*MagneticFluxDto, error) {
	a := MagneticFluxDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MagneticFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MagneticFlux struct
type MagneticFlux struct {
	value       float64
    
    webersLazy *float64 
}

// MagneticFluxFactory struct to group related functions
type MagneticFluxFactory struct{}

func (uf MagneticFluxFactory) CreateMagneticFlux(value float64, unit MagneticFluxUnits) (*MagneticFlux, error) {
	return newMagneticFlux(value, unit)
}

func (uf MagneticFluxFactory) FromDto(dto MagneticFluxDto) (*MagneticFlux, error) {
	return newMagneticFlux(dto.Value, dto.Unit)
}

func (uf MagneticFluxFactory) FromDtoJSON(data []byte) (*MagneticFlux, error) {
	unitDto, err := MagneticFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MagneticFluxFactory{}.FromDto(*unitDto)
}


// FromWeber creates a new MagneticFlux instance from Weber.
func (uf MagneticFluxFactory) FromWebers(value float64) (*MagneticFlux, error) {
	return newMagneticFlux(value, MagneticFluxWeber)
}




// newMagneticFlux creates a new MagneticFlux.
func newMagneticFlux(value float64, fromUnit MagneticFluxUnits) (*MagneticFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MagneticFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MagneticFlux in Weber.
func (a *MagneticFlux) BaseValue() float64 {
	return a.value
}


// Weber returns the value in Weber.
func (a *MagneticFlux) Webers() float64 {
	if a.webersLazy != nil {
		return *a.webersLazy
	}
	webers := a.convertFromBase(MagneticFluxWeber)
	a.webersLazy = &webers
	return webers
}


// ToDto creates an MagneticFluxDto representation.
func (a *MagneticFlux) ToDto(holdInUnit *MagneticFluxUnits) MagneticFluxDto {
	if holdInUnit == nil {
		defaultUnit := MagneticFluxWeber // Default value
		holdInUnit = &defaultUnit
	}

	return MagneticFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MagneticFluxDto representation.
func (a *MagneticFlux) ToDtoJSON(holdInUnit *MagneticFluxUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MagneticFlux to a specific unit value.
func (a *MagneticFlux) Convert(toUnit MagneticFluxUnits) float64 {
	switch toUnit { 
    case MagneticFluxWeber:
		return a.Webers()
	default:
		return 0
	}
}

func (a *MagneticFlux) convertFromBase(toUnit MagneticFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagneticFluxWeber:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *MagneticFlux) convertToBase(value float64, fromUnit MagneticFluxUnits) float64 {
	switch fromUnit { 
	case MagneticFluxWeber:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a MagneticFlux) String() string {
	return a.ToString(MagneticFluxWeber, 2)
}

// ToString formats the MagneticFlux to string.
// fractionalDigits -1 for not mention
func (a *MagneticFlux) ToString(unit MagneticFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MagneticFlux) getUnitAbbreviation(unit MagneticFluxUnits) string {
	switch unit { 
	case MagneticFluxWeber:
		return "Wb" 
	default:
		return ""
	}
}

// Check if the given MagneticFlux are equals to the current MagneticFlux
func (a *MagneticFlux) Equals(other *MagneticFlux) bool {
	return a.value == other.BaseValue()
}

// Check if the given MagneticFlux are equals to the current MagneticFlux
func (a *MagneticFlux) CompareTo(other *MagneticFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given MagneticFlux to the current MagneticFlux.
func (a *MagneticFlux) Add(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value + other.BaseValue()}
}

// Subtract the given MagneticFlux to the current MagneticFlux.
func (a *MagneticFlux) Subtract(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value - other.BaseValue()}
}

// Multiply the given MagneticFlux to the current MagneticFlux.
func (a *MagneticFlux) Multiply(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value * other.BaseValue()}
}

// Divide the given MagneticFlux to the current MagneticFlux.
func (a *MagneticFlux) Divide(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value / other.BaseValue()}
}