// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ApparentPowerUnits enumeration
type ApparentPowerUnits string

const (
    
        // 
        ApparentPowerVoltampere ApparentPowerUnits = "Voltampere"
        // 
        ApparentPowerMicrovoltampere ApparentPowerUnits = "Microvoltampere"
        // 
        ApparentPowerMillivoltampere ApparentPowerUnits = "Millivoltampere"
        // 
        ApparentPowerKilovoltampere ApparentPowerUnits = "Kilovoltampere"
        // 
        ApparentPowerMegavoltampere ApparentPowerUnits = "Megavoltampere"
        // 
        ApparentPowerGigavoltampere ApparentPowerUnits = "Gigavoltampere"
)

// ApparentPowerDto represents an ApparentPower
type ApparentPowerDto struct {
	Value float64
	Unit  ApparentPowerUnits
}

// ApparentPowerDtoFactory struct to group related functions
type ApparentPowerDtoFactory struct{}

func (udf ApparentPowerDtoFactory) FromJSON(data []byte) (*ApparentPowerDto, error) {
	a := ApparentPowerDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ApparentPowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ApparentPower struct
type ApparentPower struct {
	value       float64
    
    voltamperesLazy *float64 
    microvoltamperesLazy *float64 
    millivoltamperesLazy *float64 
    kilovoltamperesLazy *float64 
    megavoltamperesLazy *float64 
    gigavoltamperesLazy *float64 
}

// ApparentPowerFactory struct to group related functions
type ApparentPowerFactory struct{}

func (uf ApparentPowerFactory) CreateApparentPower(value float64, unit ApparentPowerUnits) (*ApparentPower, error) {
	return newApparentPower(value, unit)
}

func (uf ApparentPowerFactory) FromDto(dto ApparentPowerDto) (*ApparentPower, error) {
	return newApparentPower(dto.Value, dto.Unit)
}

func (uf ApparentPowerFactory) FromDtoJSON(data []byte) (*ApparentPower, error) {
	unitDto, err := ApparentPowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ApparentPowerFactory{}.FromDto(*unitDto)
}


// FromVoltampere creates a new ApparentPower instance from Voltampere.
func (uf ApparentPowerFactory) FromVoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerVoltampere)
}

// FromMicrovoltampere creates a new ApparentPower instance from Microvoltampere.
func (uf ApparentPowerFactory) FromMicrovoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMicrovoltampere)
}

// FromMillivoltampere creates a new ApparentPower instance from Millivoltampere.
func (uf ApparentPowerFactory) FromMillivoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMillivoltampere)
}

// FromKilovoltampere creates a new ApparentPower instance from Kilovoltampere.
func (uf ApparentPowerFactory) FromKilovoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerKilovoltampere)
}

// FromMegavoltampere creates a new ApparentPower instance from Megavoltampere.
func (uf ApparentPowerFactory) FromMegavoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMegavoltampere)
}

// FromGigavoltampere creates a new ApparentPower instance from Gigavoltampere.
func (uf ApparentPowerFactory) FromGigavoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerGigavoltampere)
}




// newApparentPower creates a new ApparentPower.
func newApparentPower(value float64, fromUnit ApparentPowerUnits) (*ApparentPower, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ApparentPower{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ApparentPower in Voltampere.
func (a *ApparentPower) BaseValue() float64 {
	return a.value
}


// Voltampere returns the value in Voltampere.
func (a *ApparentPower) Voltamperes() float64 {
	if a.voltamperesLazy != nil {
		return *a.voltamperesLazy
	}
	voltamperes := a.convertFromBase(ApparentPowerVoltampere)
	a.voltamperesLazy = &voltamperes
	return voltamperes
}

// Microvoltampere returns the value in Microvoltampere.
func (a *ApparentPower) Microvoltamperes() float64 {
	if a.microvoltamperesLazy != nil {
		return *a.microvoltamperesLazy
	}
	microvoltamperes := a.convertFromBase(ApparentPowerMicrovoltampere)
	a.microvoltamperesLazy = &microvoltamperes
	return microvoltamperes
}

// Millivoltampere returns the value in Millivoltampere.
func (a *ApparentPower) Millivoltamperes() float64 {
	if a.millivoltamperesLazy != nil {
		return *a.millivoltamperesLazy
	}
	millivoltamperes := a.convertFromBase(ApparentPowerMillivoltampere)
	a.millivoltamperesLazy = &millivoltamperes
	return millivoltamperes
}

// Kilovoltampere returns the value in Kilovoltampere.
func (a *ApparentPower) Kilovoltamperes() float64 {
	if a.kilovoltamperesLazy != nil {
		return *a.kilovoltamperesLazy
	}
	kilovoltamperes := a.convertFromBase(ApparentPowerKilovoltampere)
	a.kilovoltamperesLazy = &kilovoltamperes
	return kilovoltamperes
}

// Megavoltampere returns the value in Megavoltampere.
func (a *ApparentPower) Megavoltamperes() float64 {
	if a.megavoltamperesLazy != nil {
		return *a.megavoltamperesLazy
	}
	megavoltamperes := a.convertFromBase(ApparentPowerMegavoltampere)
	a.megavoltamperesLazy = &megavoltamperes
	return megavoltamperes
}

// Gigavoltampere returns the value in Gigavoltampere.
func (a *ApparentPower) Gigavoltamperes() float64 {
	if a.gigavoltamperesLazy != nil {
		return *a.gigavoltamperesLazy
	}
	gigavoltamperes := a.convertFromBase(ApparentPowerGigavoltampere)
	a.gigavoltamperesLazy = &gigavoltamperes
	return gigavoltamperes
}


// ToDto creates an ApparentPowerDto representation.
func (a *ApparentPower) ToDto(holdInUnit *ApparentPowerUnits) ApparentPowerDto {
	if holdInUnit == nil {
		defaultUnit := ApparentPowerVoltampere // Default value
		holdInUnit = &defaultUnit
	}

	return ApparentPowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ApparentPowerDto representation.
func (a *ApparentPower) ToDtoJSON(holdInUnit *ApparentPowerUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ApparentPower to a specific unit value.
func (a *ApparentPower) Convert(toUnit ApparentPowerUnits) float64 {
	switch toUnit { 
    case ApparentPowerVoltampere:
		return a.Voltamperes()
    case ApparentPowerMicrovoltampere:
		return a.Microvoltamperes()
    case ApparentPowerMillivoltampere:
		return a.Millivoltamperes()
    case ApparentPowerKilovoltampere:
		return a.Kilovoltamperes()
    case ApparentPowerMegavoltampere:
		return a.Megavoltamperes()
    case ApparentPowerGigavoltampere:
		return a.Gigavoltamperes()
	default:
		return 0
	}
}

func (a *ApparentPower) convertFromBase(toUnit ApparentPowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case ApparentPowerVoltampere:
		return (value) 
	case ApparentPowerMicrovoltampere:
		return ((value) / 1e-06) 
	case ApparentPowerMillivoltampere:
		return ((value) / 0.001) 
	case ApparentPowerKilovoltampere:
		return ((value) / 1000.0) 
	case ApparentPowerMegavoltampere:
		return ((value) / 1000000.0) 
	case ApparentPowerGigavoltampere:
		return ((value) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ApparentPower) convertToBase(value float64, fromUnit ApparentPowerUnits) float64 {
	switch fromUnit { 
	case ApparentPowerVoltampere:
		return (value) 
	case ApparentPowerMicrovoltampere:
		return ((value) * 1e-06) 
	case ApparentPowerMillivoltampere:
		return ((value) * 0.001) 
	case ApparentPowerKilovoltampere:
		return ((value) * 1000.0) 
	case ApparentPowerMegavoltampere:
		return ((value) * 1000000.0) 
	case ApparentPowerGigavoltampere:
		return ((value) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ApparentPower) String() string {
	return a.ToString(ApparentPowerVoltampere, 2)
}

// ToString formats the ApparentPower to string.
// fractionalDigits -1 for not mention
func (a *ApparentPower) ToString(unit ApparentPowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ApparentPower) getUnitAbbreviation(unit ApparentPowerUnits) string {
	switch unit { 
	case ApparentPowerVoltampere:
		return "VA" 
	case ApparentPowerMicrovoltampere:
		return "μVA" 
	case ApparentPowerMillivoltampere:
		return "mVA" 
	case ApparentPowerKilovoltampere:
		return "kVA" 
	case ApparentPowerMegavoltampere:
		return "MVA" 
	case ApparentPowerGigavoltampere:
		return "GVA" 
	default:
		return ""
	}
}

// Check if the given ApparentPower are equals to the current ApparentPower
func (a *ApparentPower) Equals(other *ApparentPower) bool {
	return a.value == other.BaseValue()
}

// Check if the given ApparentPower are equals to the current ApparentPower
func (a *ApparentPower) CompareTo(other *ApparentPower) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ApparentPower to the current ApparentPower.
func (a *ApparentPower) Add(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value + other.BaseValue()}
}

// Subtract the given ApparentPower to the current ApparentPower.
func (a *ApparentPower) Subtract(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value - other.BaseValue()}
}

// Multiply the given ApparentPower to the current ApparentPower.
func (a *ApparentPower) Multiply(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value * other.BaseValue()}
}

// Divide the given ApparentPower to the current ApparentPower.
func (a *ApparentPower) Divide(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value / other.BaseValue()}
}