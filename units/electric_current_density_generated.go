// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentDensityUnits enumeration
type ElectricCurrentDensityUnits string

const (
    
        // 
        ElectricCurrentDensityAmperePerSquareMeter ElectricCurrentDensityUnits = "AmperePerSquareMeter"
        // 
        ElectricCurrentDensityAmperePerSquareInch ElectricCurrentDensityUnits = "AmperePerSquareInch"
        // 
        ElectricCurrentDensityAmperePerSquareFoot ElectricCurrentDensityUnits = "AmperePerSquareFoot"
)

// ElectricCurrentDensityDto represents an ElectricCurrentDensity
type ElectricCurrentDensityDto struct {
	Value float64
	Unit  ElectricCurrentDensityUnits
}

// ElectricCurrentDensityDtoFactory struct to group related functions
type ElectricCurrentDensityDtoFactory struct{}

func (udf ElectricCurrentDensityDtoFactory) FromJSON(data []byte) (*ElectricCurrentDensityDto, error) {
	a := ElectricCurrentDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricCurrentDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricCurrentDensity struct
type ElectricCurrentDensity struct {
	value       float64
    
    amperes_per_square_meterLazy *float64 
    amperes_per_square_inchLazy *float64 
    amperes_per_square_footLazy *float64 
}

// ElectricCurrentDensityFactory struct to group related functions
type ElectricCurrentDensityFactory struct{}

func (uf ElectricCurrentDensityFactory) CreateElectricCurrentDensity(value float64, unit ElectricCurrentDensityUnits) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, unit)
}

func (uf ElectricCurrentDensityFactory) FromDto(dto ElectricCurrentDensityDto) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(dto.Value, dto.Unit)
}

func (uf ElectricCurrentDensityFactory) FromDtoJSON(data []byte) (*ElectricCurrentDensity, error) {
	unitDto, err := ElectricCurrentDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricCurrentDensityFactory{}.FromDto(*unitDto)
}


// FromAmperePerSquareMeter creates a new ElectricCurrentDensity instance from AmperePerSquareMeter.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareMeter(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareMeter)
}

// FromAmperePerSquareInch creates a new ElectricCurrentDensity instance from AmperePerSquareInch.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareInch(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareInch)
}

// FromAmperePerSquareFoot creates a new ElectricCurrentDensity instance from AmperePerSquareFoot.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareFoot(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareFoot)
}




// newElectricCurrentDensity creates a new ElectricCurrentDensity.
func newElectricCurrentDensity(value float64, fromUnit ElectricCurrentDensityUnits) (*ElectricCurrentDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrentDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrentDensity in AmperePerSquareMeter.
func (a *ElectricCurrentDensity) BaseValue() float64 {
	return a.value
}


// AmperePerSquareMeter returns the value in AmperePerSquareMeter.
func (a *ElectricCurrentDensity) AmperesPerSquareMeter() float64 {
	if a.amperes_per_square_meterLazy != nil {
		return *a.amperes_per_square_meterLazy
	}
	amperes_per_square_meter := a.convertFromBase(ElectricCurrentDensityAmperePerSquareMeter)
	a.amperes_per_square_meterLazy = &amperes_per_square_meter
	return amperes_per_square_meter
}

// AmperePerSquareInch returns the value in AmperePerSquareInch.
func (a *ElectricCurrentDensity) AmperesPerSquareInch() float64 {
	if a.amperes_per_square_inchLazy != nil {
		return *a.amperes_per_square_inchLazy
	}
	amperes_per_square_inch := a.convertFromBase(ElectricCurrentDensityAmperePerSquareInch)
	a.amperes_per_square_inchLazy = &amperes_per_square_inch
	return amperes_per_square_inch
}

// AmperePerSquareFoot returns the value in AmperePerSquareFoot.
func (a *ElectricCurrentDensity) AmperesPerSquareFoot() float64 {
	if a.amperes_per_square_footLazy != nil {
		return *a.amperes_per_square_footLazy
	}
	amperes_per_square_foot := a.convertFromBase(ElectricCurrentDensityAmperePerSquareFoot)
	a.amperes_per_square_footLazy = &amperes_per_square_foot
	return amperes_per_square_foot
}


// ToDto creates an ElectricCurrentDensityDto representation.
func (a *ElectricCurrentDensity) ToDto(holdInUnit *ElectricCurrentDensityUnits) ElectricCurrentDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentDensityAmperePerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricCurrentDensityDto representation.
func (a *ElectricCurrentDensity) ToDtoJSON(holdInUnit *ElectricCurrentDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricCurrentDensity to a specific unit value.
func (a *ElectricCurrentDensity) Convert(toUnit ElectricCurrentDensityUnits) float64 {
	switch toUnit { 
    case ElectricCurrentDensityAmperePerSquareMeter:
		return a.AmperesPerSquareMeter()
    case ElectricCurrentDensityAmperePerSquareInch:
		return a.AmperesPerSquareInch()
    case ElectricCurrentDensityAmperePerSquareFoot:
		return a.AmperesPerSquareFoot()
	default:
		return 0
	}
}

func (a *ElectricCurrentDensity) convertFromBase(toUnit ElectricCurrentDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return (value) 
	case ElectricCurrentDensityAmperePerSquareInch:
		return (value / 1.5500031000062000e3) 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return (value / 1.0763910416709722e1) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentDensity) convertToBase(value float64, fromUnit ElectricCurrentDensityUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return (value) 
	case ElectricCurrentDensityAmperePerSquareInch:
		return (value * 1.5500031000062000e3) 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return (value * 1.0763910416709722e1) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricCurrentDensity) String() string {
	return a.ToString(ElectricCurrentDensityAmperePerSquareMeter, 2)
}

// ToString formats the ElectricCurrentDensity to string.
// fractionalDigits -1 for not mention
func (a *ElectricCurrentDensity) ToString(unit ElectricCurrentDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCurrentDensity) getUnitAbbreviation(unit ElectricCurrentDensityUnits) string {
	switch unit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return "A/m²" 
	case ElectricCurrentDensityAmperePerSquareInch:
		return "A/in²" 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return "A/ft²" 
	default:
		return ""
	}
}

// Check if the given ElectricCurrentDensity are equals to the current ElectricCurrentDensity
func (a *ElectricCurrentDensity) Equals(other *ElectricCurrentDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricCurrentDensity are equals to the current ElectricCurrentDensity
func (a *ElectricCurrentDensity) CompareTo(other *ElectricCurrentDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricCurrentDensity to the current ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Add(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricCurrentDensity to the current ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Subtract(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricCurrentDensity to the current ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Multiply(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value * other.BaseValue()}
}

// Divide the given ElectricCurrentDensity to the current ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Divide(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value / other.BaseValue()}
}