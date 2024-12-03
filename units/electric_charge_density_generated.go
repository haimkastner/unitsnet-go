// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricChargeDensityUnits enumeration
type ElectricChargeDensityUnits string

const (
    
        // 
        ElectricChargeDensityCoulombPerCubicMeter ElectricChargeDensityUnits = "CoulombPerCubicMeter"
)

// ElectricChargeDensityDto represents an ElectricChargeDensity
type ElectricChargeDensityDto struct {
	Value float64
	Unit  ElectricChargeDensityUnits
}

// ElectricChargeDensityDtoFactory struct to group related functions
type ElectricChargeDensityDtoFactory struct{}

func (udf ElectricChargeDensityDtoFactory) FromJSON(data []byte) (*ElectricChargeDensityDto, error) {
	a := ElectricChargeDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricChargeDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricChargeDensity struct
type ElectricChargeDensity struct {
	value       float64
    
    coulombs_per_cubic_meterLazy *float64 
}

// ElectricChargeDensityFactory struct to group related functions
type ElectricChargeDensityFactory struct{}

func (uf ElectricChargeDensityFactory) CreateElectricChargeDensity(value float64, unit ElectricChargeDensityUnits) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(value, unit)
}

func (uf ElectricChargeDensityFactory) FromDto(dto ElectricChargeDensityDto) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(dto.Value, dto.Unit)
}

func (uf ElectricChargeDensityFactory) FromDtoJSON(data []byte) (*ElectricChargeDensity, error) {
	unitDto, err := ElectricChargeDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricChargeDensityFactory{}.FromDto(*unitDto)
}


// FromCoulombPerCubicMeter creates a new ElectricChargeDensity instance from CoulombPerCubicMeter.
func (uf ElectricChargeDensityFactory) FromCoulombsPerCubicMeter(value float64) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(value, ElectricChargeDensityCoulombPerCubicMeter)
}




// newElectricChargeDensity creates a new ElectricChargeDensity.
func newElectricChargeDensity(value float64, fromUnit ElectricChargeDensityUnits) (*ElectricChargeDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricChargeDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricChargeDensity in CoulombPerCubicMeter.
func (a *ElectricChargeDensity) BaseValue() float64 {
	return a.value
}


// CoulombPerCubicMeter returns the value in CoulombPerCubicMeter.
func (a *ElectricChargeDensity) CoulombsPerCubicMeter() float64 {
	if a.coulombs_per_cubic_meterLazy != nil {
		return *a.coulombs_per_cubic_meterLazy
	}
	coulombs_per_cubic_meter := a.convertFromBase(ElectricChargeDensityCoulombPerCubicMeter)
	a.coulombs_per_cubic_meterLazy = &coulombs_per_cubic_meter
	return coulombs_per_cubic_meter
}


// ToDto creates an ElectricChargeDensityDto representation.
func (a *ElectricChargeDensity) ToDto(holdInUnit *ElectricChargeDensityUnits) ElectricChargeDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricChargeDensityCoulombPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricChargeDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricChargeDensityDto representation.
func (a *ElectricChargeDensity) ToDtoJSON(holdInUnit *ElectricChargeDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricChargeDensity to a specific unit value.
func (a *ElectricChargeDensity) Convert(toUnit ElectricChargeDensityUnits) float64 {
	switch toUnit { 
    case ElectricChargeDensityCoulombPerCubicMeter:
		return a.CoulombsPerCubicMeter()
	default:
		return 0
	}
}

func (a *ElectricChargeDensity) convertFromBase(toUnit ElectricChargeDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *ElectricChargeDensity) convertToBase(value float64, fromUnit ElectricChargeDensityUnits) float64 {
	switch fromUnit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricChargeDensity) String() string {
	return a.ToString(ElectricChargeDensityCoulombPerCubicMeter, 2)
}

// ToString formats the ElectricChargeDensity to string.
// fractionalDigits -1 for not mention
func (a *ElectricChargeDensity) ToString(unit ElectricChargeDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricChargeDensity) getUnitAbbreviation(unit ElectricChargeDensityUnits) string {
	switch unit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return "C/m³" 
	default:
		return ""
	}
}

// Check if the given ElectricChargeDensity are equals to the current ElectricChargeDensity
func (a *ElectricChargeDensity) Equals(other *ElectricChargeDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricChargeDensity are equals to the current ElectricChargeDensity
func (a *ElectricChargeDensity) CompareTo(other *ElectricChargeDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricChargeDensity to the current ElectricChargeDensity.
func (a *ElectricChargeDensity) Add(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricChargeDensity to the current ElectricChargeDensity.
func (a *ElectricChargeDensity) Subtract(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricChargeDensity to the current ElectricChargeDensity.
func (a *ElectricChargeDensity) Multiply(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value * other.BaseValue()}
}

// Divide the given ElectricChargeDensity to the current ElectricChargeDensity.
func (a *ElectricChargeDensity) Divide(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value / other.BaseValue()}
}