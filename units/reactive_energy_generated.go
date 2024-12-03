// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ReactiveEnergyUnits enumeration
type ReactiveEnergyUnits string

const (
    
        // 
        ReactiveEnergyVoltampereReactiveHour ReactiveEnergyUnits = "VoltampereReactiveHour"
        // 
        ReactiveEnergyKilovoltampereReactiveHour ReactiveEnergyUnits = "KilovoltampereReactiveHour"
        // 
        ReactiveEnergyMegavoltampereReactiveHour ReactiveEnergyUnits = "MegavoltampereReactiveHour"
)

// ReactiveEnergyDto represents an ReactiveEnergy
type ReactiveEnergyDto struct {
	Value float64
	Unit  ReactiveEnergyUnits
}

// ReactiveEnergyDtoFactory struct to group related functions
type ReactiveEnergyDtoFactory struct{}

func (udf ReactiveEnergyDtoFactory) FromJSON(data []byte) (*ReactiveEnergyDto, error) {
	a := ReactiveEnergyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ReactiveEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ReactiveEnergy struct
type ReactiveEnergy struct {
	value       float64
    
    voltampere_reactive_hoursLazy *float64 
    kilovoltampere_reactive_hoursLazy *float64 
    megavoltampere_reactive_hoursLazy *float64 
}

// ReactiveEnergyFactory struct to group related functions
type ReactiveEnergyFactory struct{}

func (uf ReactiveEnergyFactory) CreateReactiveEnergy(value float64, unit ReactiveEnergyUnits) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, unit)
}

func (uf ReactiveEnergyFactory) FromDto(dto ReactiveEnergyDto) (*ReactiveEnergy, error) {
	return newReactiveEnergy(dto.Value, dto.Unit)
}

func (uf ReactiveEnergyFactory) FromDtoJSON(data []byte) (*ReactiveEnergy, error) {
	unitDto, err := ReactiveEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ReactiveEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereReactiveHour creates a new ReactiveEnergy instance from VoltampereReactiveHour.
func (uf ReactiveEnergyFactory) FromVoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyVoltampereReactiveHour)
}

// FromKilovoltampereReactiveHour creates a new ReactiveEnergy instance from KilovoltampereReactiveHour.
func (uf ReactiveEnergyFactory) FromKilovoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyKilovoltampereReactiveHour)
}

// FromMegavoltampereReactiveHour creates a new ReactiveEnergy instance from MegavoltampereReactiveHour.
func (uf ReactiveEnergyFactory) FromMegavoltampereReactiveHours(value float64) (*ReactiveEnergy, error) {
	return newReactiveEnergy(value, ReactiveEnergyMegavoltampereReactiveHour)
}




// newReactiveEnergy creates a new ReactiveEnergy.
func newReactiveEnergy(value float64, fromUnit ReactiveEnergyUnits) (*ReactiveEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ReactiveEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReactiveEnergy in VoltampereReactiveHour.
func (a *ReactiveEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereReactiveHour returns the value in VoltampereReactiveHour.
func (a *ReactiveEnergy) VoltampereReactiveHours() float64 {
	if a.voltampere_reactive_hoursLazy != nil {
		return *a.voltampere_reactive_hoursLazy
	}
	voltampere_reactive_hours := a.convertFromBase(ReactiveEnergyVoltampereReactiveHour)
	a.voltampere_reactive_hoursLazy = &voltampere_reactive_hours
	return voltampere_reactive_hours
}

// KilovoltampereReactiveHour returns the value in KilovoltampereReactiveHour.
func (a *ReactiveEnergy) KilovoltampereReactiveHours() float64 {
	if a.kilovoltampere_reactive_hoursLazy != nil {
		return *a.kilovoltampere_reactive_hoursLazy
	}
	kilovoltampere_reactive_hours := a.convertFromBase(ReactiveEnergyKilovoltampereReactiveHour)
	a.kilovoltampere_reactive_hoursLazy = &kilovoltampere_reactive_hours
	return kilovoltampere_reactive_hours
}

// MegavoltampereReactiveHour returns the value in MegavoltampereReactiveHour.
func (a *ReactiveEnergy) MegavoltampereReactiveHours() float64 {
	if a.megavoltampere_reactive_hoursLazy != nil {
		return *a.megavoltampere_reactive_hoursLazy
	}
	megavoltampere_reactive_hours := a.convertFromBase(ReactiveEnergyMegavoltampereReactiveHour)
	a.megavoltampere_reactive_hoursLazy = &megavoltampere_reactive_hours
	return megavoltampere_reactive_hours
}


// ToDto creates an ReactiveEnergyDto representation.
func (a *ReactiveEnergy) ToDto(holdInUnit *ReactiveEnergyUnits) ReactiveEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ReactiveEnergyVoltampereReactiveHour // Default value
		holdInUnit = &defaultUnit
	}

	return ReactiveEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ReactiveEnergyDto representation.
func (a *ReactiveEnergy) ToDtoJSON(holdInUnit *ReactiveEnergyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ReactiveEnergy to a specific unit value.
func (a *ReactiveEnergy) Convert(toUnit ReactiveEnergyUnits) float64 {
	switch toUnit { 
    case ReactiveEnergyVoltampereReactiveHour:
		return a.VoltampereReactiveHours()
    case ReactiveEnergyKilovoltampereReactiveHour:
		return a.KilovoltampereReactiveHours()
    case ReactiveEnergyMegavoltampereReactiveHour:
		return a.MegavoltampereReactiveHours()
	default:
		return 0
	}
}

func (a *ReactiveEnergy) convertFromBase(toUnit ReactiveEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return ((value) / 1000.0) 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ReactiveEnergy) convertToBase(value float64, fromUnit ReactiveEnergyUnits) float64 {
	switch fromUnit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return ((value) * 1000.0) 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ReactiveEnergy) String() string {
	return a.ToString(ReactiveEnergyVoltampereReactiveHour, 2)
}

// ToString formats the ReactiveEnergy to string.
// fractionalDigits -1 for not mention
func (a *ReactiveEnergy) ToString(unit ReactiveEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ReactiveEnergy) getUnitAbbreviation(unit ReactiveEnergyUnits) string {
	switch unit { 
	case ReactiveEnergyVoltampereReactiveHour:
		return "varh" 
	case ReactiveEnergyKilovoltampereReactiveHour:
		return "kvarh" 
	case ReactiveEnergyMegavoltampereReactiveHour:
		return "Mvarh" 
	default:
		return ""
	}
}

// Check if the given ReactiveEnergy are equals to the current ReactiveEnergy
func (a *ReactiveEnergy) Equals(other *ReactiveEnergy) bool {
	return a.value == other.BaseValue()
}

// Check if the given ReactiveEnergy are equals to the current ReactiveEnergy
func (a *ReactiveEnergy) CompareTo(other *ReactiveEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ReactiveEnergy to the current ReactiveEnergy.
func (a *ReactiveEnergy) Add(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value + other.BaseValue()}
}

// Subtract the given ReactiveEnergy to the current ReactiveEnergy.
func (a *ReactiveEnergy) Subtract(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value - other.BaseValue()}
}

// Multiply the given ReactiveEnergy to the current ReactiveEnergy.
func (a *ReactiveEnergy) Multiply(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value * other.BaseValue()}
}

// Divide the given ReactiveEnergy to the current ReactiveEnergy.
func (a *ReactiveEnergy) Divide(other *ReactiveEnergy) *ReactiveEnergy {
	return &ReactiveEnergy{value: a.value / other.BaseValue()}
}