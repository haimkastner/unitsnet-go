// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RadiationEquivalentDoseUnits enumeration
type RadiationEquivalentDoseUnits string

const (
    
        // The sievert is a unit in the International System of Units (SI) intended to represent the stochastic health risk of ionizing radiation, which is defined as the probability of causing radiation-induced cancer and genetic damage.
        RadiationEquivalentDoseSievert RadiationEquivalentDoseUnits = "Sievert"
        // 
        RadiationEquivalentDoseRoentgenEquivalentMan RadiationEquivalentDoseUnits = "RoentgenEquivalentMan"
        // 
        RadiationEquivalentDoseNanosievert RadiationEquivalentDoseUnits = "Nanosievert"
        // 
        RadiationEquivalentDoseMicrosievert RadiationEquivalentDoseUnits = "Microsievert"
        // 
        RadiationEquivalentDoseMillisievert RadiationEquivalentDoseUnits = "Millisievert"
        // 
        RadiationEquivalentDoseMilliroentgenEquivalentMan RadiationEquivalentDoseUnits = "MilliroentgenEquivalentMan"
)

// RadiationEquivalentDoseDto represents an RadiationEquivalentDose
type RadiationEquivalentDoseDto struct {
	Value float64
	Unit  RadiationEquivalentDoseUnits
}

// RadiationEquivalentDoseDtoFactory struct to group related functions
type RadiationEquivalentDoseDtoFactory struct{}

func (udf RadiationEquivalentDoseDtoFactory) FromJSON(data []byte) (*RadiationEquivalentDoseDto, error) {
	a := RadiationEquivalentDoseDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RadiationEquivalentDoseDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RadiationEquivalentDose struct
type RadiationEquivalentDose struct {
	value       float64
    
    sievertsLazy *float64 
    roentgens_equivalent_manLazy *float64 
    nanosievertsLazy *float64 
    microsievertsLazy *float64 
    millisievertsLazy *float64 
    milliroentgens_equivalent_manLazy *float64 
}

// RadiationEquivalentDoseFactory struct to group related functions
type RadiationEquivalentDoseFactory struct{}

func (uf RadiationEquivalentDoseFactory) CreateRadiationEquivalentDose(value float64, unit RadiationEquivalentDoseUnits) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, unit)
}

func (uf RadiationEquivalentDoseFactory) FromDto(dto RadiationEquivalentDoseDto) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(dto.Value, dto.Unit)
}

func (uf RadiationEquivalentDoseFactory) FromDtoJSON(data []byte) (*RadiationEquivalentDose, error) {
	unitDto, err := RadiationEquivalentDoseDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RadiationEquivalentDoseFactory{}.FromDto(*unitDto)
}


// FromSievert creates a new RadiationEquivalentDose instance from Sievert.
func (uf RadiationEquivalentDoseFactory) FromSieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseSievert)
}

// FromRoentgenEquivalentMan creates a new RadiationEquivalentDose instance from RoentgenEquivalentMan.
func (uf RadiationEquivalentDoseFactory) FromRoentgensEquivalentMan(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseRoentgenEquivalentMan)
}

// FromNanosievert creates a new RadiationEquivalentDose instance from Nanosievert.
func (uf RadiationEquivalentDoseFactory) FromNanosieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseNanosievert)
}

// FromMicrosievert creates a new RadiationEquivalentDose instance from Microsievert.
func (uf RadiationEquivalentDoseFactory) FromMicrosieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMicrosievert)
}

// FromMillisievert creates a new RadiationEquivalentDose instance from Millisievert.
func (uf RadiationEquivalentDoseFactory) FromMillisieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMillisievert)
}

// FromMilliroentgenEquivalentMan creates a new RadiationEquivalentDose instance from MilliroentgenEquivalentMan.
func (uf RadiationEquivalentDoseFactory) FromMilliroentgensEquivalentMan(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMilliroentgenEquivalentMan)
}




// newRadiationEquivalentDose creates a new RadiationEquivalentDose.
func newRadiationEquivalentDose(value float64, fromUnit RadiationEquivalentDoseUnits) (*RadiationEquivalentDose, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RadiationEquivalentDose{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RadiationEquivalentDose in Sievert.
func (a *RadiationEquivalentDose) BaseValue() float64 {
	return a.value
}


// Sievert returns the value in Sievert.
func (a *RadiationEquivalentDose) Sieverts() float64 {
	if a.sievertsLazy != nil {
		return *a.sievertsLazy
	}
	sieverts := a.convertFromBase(RadiationEquivalentDoseSievert)
	a.sievertsLazy = &sieverts
	return sieverts
}

// RoentgenEquivalentMan returns the value in RoentgenEquivalentMan.
func (a *RadiationEquivalentDose) RoentgensEquivalentMan() float64 {
	if a.roentgens_equivalent_manLazy != nil {
		return *a.roentgens_equivalent_manLazy
	}
	roentgens_equivalent_man := a.convertFromBase(RadiationEquivalentDoseRoentgenEquivalentMan)
	a.roentgens_equivalent_manLazy = &roentgens_equivalent_man
	return roentgens_equivalent_man
}

// Nanosievert returns the value in Nanosievert.
func (a *RadiationEquivalentDose) Nanosieverts() float64 {
	if a.nanosievertsLazy != nil {
		return *a.nanosievertsLazy
	}
	nanosieverts := a.convertFromBase(RadiationEquivalentDoseNanosievert)
	a.nanosievertsLazy = &nanosieverts
	return nanosieverts
}

// Microsievert returns the value in Microsievert.
func (a *RadiationEquivalentDose) Microsieverts() float64 {
	if a.microsievertsLazy != nil {
		return *a.microsievertsLazy
	}
	microsieverts := a.convertFromBase(RadiationEquivalentDoseMicrosievert)
	a.microsievertsLazy = &microsieverts
	return microsieverts
}

// Millisievert returns the value in Millisievert.
func (a *RadiationEquivalentDose) Millisieverts() float64 {
	if a.millisievertsLazy != nil {
		return *a.millisievertsLazy
	}
	millisieverts := a.convertFromBase(RadiationEquivalentDoseMillisievert)
	a.millisievertsLazy = &millisieverts
	return millisieverts
}

// MilliroentgenEquivalentMan returns the value in MilliroentgenEquivalentMan.
func (a *RadiationEquivalentDose) MilliroentgensEquivalentMan() float64 {
	if a.milliroentgens_equivalent_manLazy != nil {
		return *a.milliroentgens_equivalent_manLazy
	}
	milliroentgens_equivalent_man := a.convertFromBase(RadiationEquivalentDoseMilliroentgenEquivalentMan)
	a.milliroentgens_equivalent_manLazy = &milliroentgens_equivalent_man
	return milliroentgens_equivalent_man
}


// ToDto creates an RadiationEquivalentDoseDto representation.
func (a *RadiationEquivalentDose) ToDto(holdInUnit *RadiationEquivalentDoseUnits) RadiationEquivalentDoseDto {
	if holdInUnit == nil {
		defaultUnit := RadiationEquivalentDoseSievert // Default value
		holdInUnit = &defaultUnit
	}

	return RadiationEquivalentDoseDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RadiationEquivalentDoseDto representation.
func (a *RadiationEquivalentDose) ToDtoJSON(holdInUnit *RadiationEquivalentDoseUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RadiationEquivalentDose to a specific unit value.
func (a *RadiationEquivalentDose) Convert(toUnit RadiationEquivalentDoseUnits) float64 {
	switch toUnit { 
    case RadiationEquivalentDoseSievert:
		return a.Sieverts()
    case RadiationEquivalentDoseRoentgenEquivalentMan:
		return a.RoentgensEquivalentMan()
    case RadiationEquivalentDoseNanosievert:
		return a.Nanosieverts()
    case RadiationEquivalentDoseMicrosievert:
		return a.Microsieverts()
    case RadiationEquivalentDoseMillisievert:
		return a.Millisieverts()
    case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return a.MilliroentgensEquivalentMan()
	default:
		return 0
	}
}

func (a *RadiationEquivalentDose) convertFromBase(toUnit RadiationEquivalentDoseUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadiationEquivalentDoseSievert:
		return (value) 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return (value * 100) 
	case RadiationEquivalentDoseNanosievert:
		return ((value) / 1e-09) 
	case RadiationEquivalentDoseMicrosievert:
		return ((value) / 1e-06) 
	case RadiationEquivalentDoseMillisievert:
		return ((value) / 0.001) 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return ((value * 100) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RadiationEquivalentDose) convertToBase(value float64, fromUnit RadiationEquivalentDoseUnits) float64 {
	switch fromUnit { 
	case RadiationEquivalentDoseSievert:
		return (value) 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return (value / 100) 
	case RadiationEquivalentDoseNanosievert:
		return ((value) * 1e-09) 
	case RadiationEquivalentDoseMicrosievert:
		return ((value) * 1e-06) 
	case RadiationEquivalentDoseMillisievert:
		return ((value) * 0.001) 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return ((value / 100) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RadiationEquivalentDose) String() string {
	return a.ToString(RadiationEquivalentDoseSievert, 2)
}

// ToString formats the RadiationEquivalentDose to string.
// fractionalDigits -1 for not mention
func (a *RadiationEquivalentDose) ToString(unit RadiationEquivalentDoseUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RadiationEquivalentDose) getUnitAbbreviation(unit RadiationEquivalentDoseUnits) string {
	switch unit { 
	case RadiationEquivalentDoseSievert:
		return "Sv" 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return "rem" 
	case RadiationEquivalentDoseNanosievert:
		return "nSv" 
	case RadiationEquivalentDoseMicrosievert:
		return "μSv" 
	case RadiationEquivalentDoseMillisievert:
		return "mSv" 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return "mrem" 
	default:
		return ""
	}
}

// Check if the given RadiationEquivalentDose are equals to the current RadiationEquivalentDose
func (a *RadiationEquivalentDose) Equals(other *RadiationEquivalentDose) bool {
	return a.value == other.BaseValue()
}

// Check if the given RadiationEquivalentDose are equals to the current RadiationEquivalentDose
func (a *RadiationEquivalentDose) CompareTo(other *RadiationEquivalentDose) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RadiationEquivalentDose to the current RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Add(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value + other.BaseValue()}
}

// Subtract the given RadiationEquivalentDose to the current RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Subtract(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value - other.BaseValue()}
}

// Multiply the given RadiationEquivalentDose to the current RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Multiply(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value * other.BaseValue()}
}

// Divide the given RadiationEquivalentDose to the current RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Divide(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value / other.BaseValue()}
}