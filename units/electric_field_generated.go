// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricFieldUnits enumeration
type ElectricFieldUnits string

const (
    
        // 
        ElectricFieldVoltPerMeter ElectricFieldUnits = "VoltPerMeter"
)

// ElectricFieldDto represents an ElectricField
type ElectricFieldDto struct {
	Value float64
	Unit  ElectricFieldUnits
}

// ElectricFieldDtoFactory struct to group related functions
type ElectricFieldDtoFactory struct{}

func (udf ElectricFieldDtoFactory) FromJSON(data []byte) (*ElectricFieldDto, error) {
	a := ElectricFieldDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricFieldDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricField struct
type ElectricField struct {
	value       float64
    
    volts_per_meterLazy *float64 
}

// ElectricFieldFactory struct to group related functions
type ElectricFieldFactory struct{}

func (uf ElectricFieldFactory) CreateElectricField(value float64, unit ElectricFieldUnits) (*ElectricField, error) {
	return newElectricField(value, unit)
}

func (uf ElectricFieldFactory) FromDto(dto ElectricFieldDto) (*ElectricField, error) {
	return newElectricField(dto.Value, dto.Unit)
}

func (uf ElectricFieldFactory) FromDtoJSON(data []byte) (*ElectricField, error) {
	unitDto, err := ElectricFieldDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricFieldFactory{}.FromDto(*unitDto)
}


// FromVoltPerMeter creates a new ElectricField instance from VoltPerMeter.
func (uf ElectricFieldFactory) FromVoltsPerMeter(value float64) (*ElectricField, error) {
	return newElectricField(value, ElectricFieldVoltPerMeter)
}




// newElectricField creates a new ElectricField.
func newElectricField(value float64, fromUnit ElectricFieldUnits) (*ElectricField, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricField{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricField in VoltPerMeter.
func (a *ElectricField) BaseValue() float64 {
	return a.value
}


// VoltPerMeter returns the value in VoltPerMeter.
func (a *ElectricField) VoltsPerMeter() float64 {
	if a.volts_per_meterLazy != nil {
		return *a.volts_per_meterLazy
	}
	volts_per_meter := a.convertFromBase(ElectricFieldVoltPerMeter)
	a.volts_per_meterLazy = &volts_per_meter
	return volts_per_meter
}


// ToDto creates an ElectricFieldDto representation.
func (a *ElectricField) ToDto(holdInUnit *ElectricFieldUnits) ElectricFieldDto {
	if holdInUnit == nil {
		defaultUnit := ElectricFieldVoltPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricFieldDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricFieldDto representation.
func (a *ElectricField) ToDtoJSON(holdInUnit *ElectricFieldUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricField to a specific unit value.
func (a *ElectricField) Convert(toUnit ElectricFieldUnits) float64 {
	switch toUnit { 
    case ElectricFieldVoltPerMeter:
		return a.VoltsPerMeter()
	default:
		return 0
	}
}

func (a *ElectricField) convertFromBase(toUnit ElectricFieldUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricFieldVoltPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *ElectricField) convertToBase(value float64, fromUnit ElectricFieldUnits) float64 {
	switch fromUnit { 
	case ElectricFieldVoltPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricField) String() string {
	return a.ToString(ElectricFieldVoltPerMeter, 2)
}

// ToString formats the ElectricField to string.
// fractionalDigits -1 for not mention
func (a *ElectricField) ToString(unit ElectricFieldUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricField) getUnitAbbreviation(unit ElectricFieldUnits) string {
	switch unit { 
	case ElectricFieldVoltPerMeter:
		return "V/m" 
	default:
		return ""
	}
}

// Check if the given ElectricField are equals to the current ElectricField
func (a *ElectricField) Equals(other *ElectricField) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricField are equals to the current ElectricField
func (a *ElectricField) CompareTo(other *ElectricField) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricField to the current ElectricField.
func (a *ElectricField) Add(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricField to the current ElectricField.
func (a *ElectricField) Subtract(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricField to the current ElectricField.
func (a *ElectricField) Multiply(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value * other.BaseValue()}
}

// Divide the given ElectricField to the current ElectricField.
func (a *ElectricField) Divide(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value / other.BaseValue()}
}