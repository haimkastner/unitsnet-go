// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PowerRatioUnits enumeration
type PowerRatioUnits string

const (
    
        // 
        PowerRatioDecibelWatt PowerRatioUnits = "DecibelWatt"
        // 
        PowerRatioDecibelMilliwatt PowerRatioUnits = "DecibelMilliwatt"
)

// PowerRatioDto represents an PowerRatio
type PowerRatioDto struct {
	Value float64
	Unit  PowerRatioUnits
}

// PowerRatioDtoFactory struct to group related functions
type PowerRatioDtoFactory struct{}

func (udf PowerRatioDtoFactory) FromJSON(data []byte) (*PowerRatioDto, error) {
	a := PowerRatioDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PowerRatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// PowerRatio struct
type PowerRatio struct {
	value       float64
    
    decibel_wattsLazy *float64 
    decibel_milliwattsLazy *float64 
}

// PowerRatioFactory struct to group related functions
type PowerRatioFactory struct{}

func (uf PowerRatioFactory) CreatePowerRatio(value float64, unit PowerRatioUnits) (*PowerRatio, error) {
	return newPowerRatio(value, unit)
}

func (uf PowerRatioFactory) FromDto(dto PowerRatioDto) (*PowerRatio, error) {
	return newPowerRatio(dto.Value, dto.Unit)
}

func (uf PowerRatioFactory) FromDtoJSON(data []byte) (*PowerRatio, error) {
	unitDto, err := PowerRatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PowerRatioFactory{}.FromDto(*unitDto)
}


// FromDecibelWatt creates a new PowerRatio instance from DecibelWatt.
func (uf PowerRatioFactory) FromDecibelWatts(value float64) (*PowerRatio, error) {
	return newPowerRatio(value, PowerRatioDecibelWatt)
}

// FromDecibelMilliwatt creates a new PowerRatio instance from DecibelMilliwatt.
func (uf PowerRatioFactory) FromDecibelMilliwatts(value float64) (*PowerRatio, error) {
	return newPowerRatio(value, PowerRatioDecibelMilliwatt)
}




// newPowerRatio creates a new PowerRatio.
func newPowerRatio(value float64, fromUnit PowerRatioUnits) (*PowerRatio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &PowerRatio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PowerRatio in DecibelWatt.
func (a *PowerRatio) BaseValue() float64 {
	return a.value
}


// DecibelWatt returns the value in DecibelWatt.
func (a *PowerRatio) DecibelWatts() float64 {
	if a.decibel_wattsLazy != nil {
		return *a.decibel_wattsLazy
	}
	decibel_watts := a.convertFromBase(PowerRatioDecibelWatt)
	a.decibel_wattsLazy = &decibel_watts
	return decibel_watts
}

// DecibelMilliwatt returns the value in DecibelMilliwatt.
func (a *PowerRatio) DecibelMilliwatts() float64 {
	if a.decibel_milliwattsLazy != nil {
		return *a.decibel_milliwattsLazy
	}
	decibel_milliwatts := a.convertFromBase(PowerRatioDecibelMilliwatt)
	a.decibel_milliwattsLazy = &decibel_milliwatts
	return decibel_milliwatts
}


// ToDto creates an PowerRatioDto representation.
func (a *PowerRatio) ToDto(holdInUnit *PowerRatioUnits) PowerRatioDto {
	if holdInUnit == nil {
		defaultUnit := PowerRatioDecibelWatt // Default value
		holdInUnit = &defaultUnit
	}

	return PowerRatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PowerRatioDto representation.
func (a *PowerRatio) ToDtoJSON(holdInUnit *PowerRatioUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts PowerRatio to a specific unit value.
func (a *PowerRatio) Convert(toUnit PowerRatioUnits) float64 {
	switch toUnit { 
    case PowerRatioDecibelWatt:
		return a.DecibelWatts()
    case PowerRatioDecibelMilliwatt:
		return a.DecibelMilliwatts()
	default:
		return 0
	}
}

func (a *PowerRatio) convertFromBase(toUnit PowerRatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case PowerRatioDecibelWatt:
		return (value) 
	case PowerRatioDecibelMilliwatt:
		return (value + 30) 
	default:
		return math.NaN()
	}
}

func (a *PowerRatio) convertToBase(value float64, fromUnit PowerRatioUnits) float64 {
	switch fromUnit { 
	case PowerRatioDecibelWatt:
		return (value) 
	case PowerRatioDecibelMilliwatt:
		return (value - 30) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a PowerRatio) String() string {
	return a.ToString(PowerRatioDecibelWatt, 2)
}

// ToString formats the PowerRatio to string.
// fractionalDigits -1 for not mention
func (a *PowerRatio) ToString(unit PowerRatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *PowerRatio) getUnitAbbreviation(unit PowerRatioUnits) string {
	switch unit { 
	case PowerRatioDecibelWatt:
		return "dBW" 
	case PowerRatioDecibelMilliwatt:
		return "dBmW" 
	default:
		return ""
	}
}

// Check if the given PowerRatio are equals to the current PowerRatio
func (a *PowerRatio) Equals(other *PowerRatio) bool {
	return a.value == other.BaseValue()
}

// Check if the given PowerRatio are equals to the current PowerRatio
func (a *PowerRatio) CompareTo(other *PowerRatio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given PowerRatio to the current PowerRatio.
func (a *PowerRatio) Add(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value + other.BaseValue()}
}

// Subtract the given PowerRatio to the current PowerRatio.
func (a *PowerRatio) Subtract(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value - other.BaseValue()}
}

// Multiply the given PowerRatio to the current PowerRatio.
func (a *PowerRatio) Multiply(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value * other.BaseValue()}
}

// Divide the given PowerRatio to the current PowerRatio.
func (a *PowerRatio) Divide(other *PowerRatio) *PowerRatio {
	return &PowerRatio{value: a.value / other.BaseValue()}
}