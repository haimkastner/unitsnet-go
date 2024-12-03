// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ApparentEnergyUnits enumeration
type ApparentEnergyUnits string

const (
    
        // 
        ApparentEnergyVoltampereHour ApparentEnergyUnits = "VoltampereHour"
        // 
        ApparentEnergyKilovoltampereHour ApparentEnergyUnits = "KilovoltampereHour"
        // 
        ApparentEnergyMegavoltampereHour ApparentEnergyUnits = "MegavoltampereHour"
)

// ApparentEnergyDto represents an ApparentEnergy
type ApparentEnergyDto struct {
	Value float64
	Unit  ApparentEnergyUnits
}

// ApparentEnergyDtoFactory struct to group related functions
type ApparentEnergyDtoFactory struct{}

func (udf ApparentEnergyDtoFactory) FromJSON(data []byte) (*ApparentEnergyDto, error) {
	a := ApparentEnergyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ApparentEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ApparentEnergy struct
type ApparentEnergy struct {
	value       float64
    
    voltampere_hoursLazy *float64 
    kilovoltampere_hoursLazy *float64 
    megavoltampere_hoursLazy *float64 
}

// ApparentEnergyFactory struct to group related functions
type ApparentEnergyFactory struct{}

func (uf ApparentEnergyFactory) CreateApparentEnergy(value float64, unit ApparentEnergyUnits) (*ApparentEnergy, error) {
	return newApparentEnergy(value, unit)
}

func (uf ApparentEnergyFactory) FromDto(dto ApparentEnergyDto) (*ApparentEnergy, error) {
	return newApparentEnergy(dto.Value, dto.Unit)
}

func (uf ApparentEnergyFactory) FromDtoJSON(data []byte) (*ApparentEnergy, error) {
	unitDto, err := ApparentEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ApparentEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereHour creates a new ApparentEnergy instance from VoltampereHour.
func (uf ApparentEnergyFactory) FromVoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyVoltampereHour)
}

// FromKilovoltampereHour creates a new ApparentEnergy instance from KilovoltampereHour.
func (uf ApparentEnergyFactory) FromKilovoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyKilovoltampereHour)
}

// FromMegavoltampereHour creates a new ApparentEnergy instance from MegavoltampereHour.
func (uf ApparentEnergyFactory) FromMegavoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyMegavoltampereHour)
}




// newApparentEnergy creates a new ApparentEnergy.
func newApparentEnergy(value float64, fromUnit ApparentEnergyUnits) (*ApparentEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ApparentEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ApparentEnergy in VoltampereHour.
func (a *ApparentEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereHour returns the value in VoltampereHour.
func (a *ApparentEnergy) VoltampereHours() float64 {
	if a.voltampere_hoursLazy != nil {
		return *a.voltampere_hoursLazy
	}
	voltampere_hours := a.convertFromBase(ApparentEnergyVoltampereHour)
	a.voltampere_hoursLazy = &voltampere_hours
	return voltampere_hours
}

// KilovoltampereHour returns the value in KilovoltampereHour.
func (a *ApparentEnergy) KilovoltampereHours() float64 {
	if a.kilovoltampere_hoursLazy != nil {
		return *a.kilovoltampere_hoursLazy
	}
	kilovoltampere_hours := a.convertFromBase(ApparentEnergyKilovoltampereHour)
	a.kilovoltampere_hoursLazy = &kilovoltampere_hours
	return kilovoltampere_hours
}

// MegavoltampereHour returns the value in MegavoltampereHour.
func (a *ApparentEnergy) MegavoltampereHours() float64 {
	if a.megavoltampere_hoursLazy != nil {
		return *a.megavoltampere_hoursLazy
	}
	megavoltampere_hours := a.convertFromBase(ApparentEnergyMegavoltampereHour)
	a.megavoltampere_hoursLazy = &megavoltampere_hours
	return megavoltampere_hours
}


// ToDto creates an ApparentEnergyDto representation.
func (a *ApparentEnergy) ToDto(holdInUnit *ApparentEnergyUnits) ApparentEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ApparentEnergyVoltampereHour // Default value
		holdInUnit = &defaultUnit
	}

	return ApparentEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ApparentEnergyDto representation.
func (a *ApparentEnergy) ToDtoJSON(holdInUnit *ApparentEnergyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ApparentEnergy to a specific unit value.
func (a *ApparentEnergy) Convert(toUnit ApparentEnergyUnits) float64 {
	switch toUnit { 
    case ApparentEnergyVoltampereHour:
		return a.VoltampereHours()
    case ApparentEnergyKilovoltampereHour:
		return a.KilovoltampereHours()
    case ApparentEnergyMegavoltampereHour:
		return a.MegavoltampereHours()
	default:
		return 0
	}
}

func (a *ApparentEnergy) convertFromBase(toUnit ApparentEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ApparentEnergyVoltampereHour:
		return (value) 
	case ApparentEnergyKilovoltampereHour:
		return ((value) / 1000.0) 
	case ApparentEnergyMegavoltampereHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ApparentEnergy) convertToBase(value float64, fromUnit ApparentEnergyUnits) float64 {
	switch fromUnit { 
	case ApparentEnergyVoltampereHour:
		return (value) 
	case ApparentEnergyKilovoltampereHour:
		return ((value) * 1000.0) 
	case ApparentEnergyMegavoltampereHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ApparentEnergy) String() string {
	return a.ToString(ApparentEnergyVoltampereHour, 2)
}

// ToString formats the ApparentEnergy to string.
// fractionalDigits -1 for not mention
func (a *ApparentEnergy) ToString(unit ApparentEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ApparentEnergy) getUnitAbbreviation(unit ApparentEnergyUnits) string {
	switch unit { 
	case ApparentEnergyVoltampereHour:
		return "VAh" 
	case ApparentEnergyKilovoltampereHour:
		return "kVAh" 
	case ApparentEnergyMegavoltampereHour:
		return "MVAh" 
	default:
		return ""
	}
}

// Check if the given ApparentEnergy are equals to the current ApparentEnergy
func (a *ApparentEnergy) Equals(other *ApparentEnergy) bool {
	return a.value == other.BaseValue()
}

// Check if the given ApparentEnergy are equals to the current ApparentEnergy
func (a *ApparentEnergy) CompareTo(other *ApparentEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ApparentEnergy to the current ApparentEnergy.
func (a *ApparentEnergy) Add(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value + other.BaseValue()}
}

// Subtract the given ApparentEnergy to the current ApparentEnergy.
func (a *ApparentEnergy) Subtract(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value - other.BaseValue()}
}

// Multiply the given ApparentEnergy to the current ApparentEnergy.
func (a *ApparentEnergy) Multiply(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value * other.BaseValue()}
}

// Divide the given ApparentEnergy to the current ApparentEnergy.
func (a *ApparentEnergy) Divide(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value / other.BaseValue()}
}