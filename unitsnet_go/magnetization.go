package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MagnetizationUnits enumeration
type MagnetizationUnits string

const (
    
        // 
        MagnetizationAmperePerMeter MagnetizationUnits = "AmperePerMeter"
)

// MagnetizationDto represents an Magnetization
type MagnetizationDto struct {
	Value float64
	Unit  MagnetizationUnits
}

// MagnetizationDtoFactory struct to group related functions
type MagnetizationDtoFactory struct{}

func (udf MagnetizationDtoFactory) FromJSON(data []byte) (*MagnetizationDto, error) {
	a := MagnetizationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MagnetizationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Magnetization struct
type Magnetization struct {
	value       float64
    
    amperes_per_meterLazy *float64 
}

// MagnetizationFactory struct to group related functions
type MagnetizationFactory struct{}

func (uf MagnetizationFactory) CreateMagnetization(value float64, unit MagnetizationUnits) (*Magnetization, error) {
	return newMagnetization(value, unit)
}

func (uf MagnetizationFactory) FromDto(dto MagnetizationDto) (*Magnetization, error) {
	return newMagnetization(dto.Value, dto.Unit)
}

func (uf MagnetizationFactory) FromDtoJSON(data []byte) (*Magnetization, error) {
	unitDto, err := MagnetizationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MagnetizationFactory{}.FromDto(*unitDto)
}


// FromAmperePerMeter creates a new Magnetization instance from AmperePerMeter.
func (uf MagnetizationFactory) FromAmperesPerMeter(value float64) (*Magnetization, error) {
	return newMagnetization(value, MagnetizationAmperePerMeter)
}




// newMagnetization creates a new Magnetization.
func newMagnetization(value float64, fromUnit MagnetizationUnits) (*Magnetization, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Magnetization{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Magnetization in AmperePerMeter.
func (a *Magnetization) BaseValue() float64 {
	return a.value
}


// AmperePerMeter returns the value in AmperePerMeter.
func (a *Magnetization) AmperesPerMeter() float64 {
	if a.amperes_per_meterLazy != nil {
		return *a.amperes_per_meterLazy
	}
	amperes_per_meter := a.convertFromBase(MagnetizationAmperePerMeter)
	a.amperes_per_meterLazy = &amperes_per_meter
	return amperes_per_meter
}


// ToDto creates an MagnetizationDto representation.
func (a *Magnetization) ToDto(holdInUnit *MagnetizationUnits) MagnetizationDto {
	if holdInUnit == nil {
		defaultUnit := MagnetizationAmperePerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MagnetizationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MagnetizationDto representation.
func (a *Magnetization) ToDtoJSON(holdInUnit *MagnetizationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Magnetization to a specific unit value.
func (a *Magnetization) Convert(toUnit MagnetizationUnits) float64 {
	switch toUnit { 
    case MagnetizationAmperePerMeter:
		return a.AmperesPerMeter()
	default:
		return 0
	}
}

func (a *Magnetization) convertFromBase(toUnit MagnetizationUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagnetizationAmperePerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Magnetization) convertToBase(value float64, fromUnit MagnetizationUnits) float64 {
	switch fromUnit { 
	case MagnetizationAmperePerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Magnetization) String() string {
	return a.ToString(MagnetizationAmperePerMeter, 2)
}

// ToString formats the Magnetization to string.
// fractionalDigits -1 for not mention
func (a *Magnetization) ToString(unit MagnetizationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Magnetization) getUnitAbbreviation(unit MagnetizationUnits) string {
	switch unit { 
	case MagnetizationAmperePerMeter:
		return "A/m" 
	default:
		return ""
	}
}

// Check if the given Magnetization are equals to the current Magnetization
func (a *Magnetization) Equals(other *Magnetization) bool {
	return a.value == other.BaseValue()
}

// Check if the given Magnetization are equals to the current Magnetization
func (a *Magnetization) CompareTo(other *Magnetization) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Magnetization to the current Magnetization.
func (a *Magnetization) Add(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value + other.BaseValue()}
}

// Subtract the given Magnetization to the current Magnetization.
func (a *Magnetization) Subtract(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value - other.BaseValue()}
}

// Multiply the given Magnetization to the current Magnetization.
func (a *Magnetization) Multiply(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value * other.BaseValue()}
}

// Divide the given Magnetization to the current Magnetization.
func (a *Magnetization) Divide(other *Magnetization) *Magnetization {
	return &Magnetization{value: a.value / other.BaseValue()}
}