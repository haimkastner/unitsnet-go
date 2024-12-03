package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TurbidityUnits enumeration
type TurbidityUnits string

const (
    
        // 
        TurbidityNTU TurbidityUnits = "NTU"
)

// TurbidityDto represents an Turbidity
type TurbidityDto struct {
	Value float64
	Unit  TurbidityUnits
}

// TurbidityDtoFactory struct to group related functions
type TurbidityDtoFactory struct{}

func (udf TurbidityDtoFactory) FromJSON(data []byte) (*TurbidityDto, error) {
	a := TurbidityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TurbidityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Turbidity struct
type Turbidity struct {
	value       float64
    
    ntuLazy *float64 
}

// TurbidityFactory struct to group related functions
type TurbidityFactory struct{}

func (uf TurbidityFactory) CreateTurbidity(value float64, unit TurbidityUnits) (*Turbidity, error) {
	return newTurbidity(value, unit)
}

func (uf TurbidityFactory) FromDto(dto TurbidityDto) (*Turbidity, error) {
	return newTurbidity(dto.Value, dto.Unit)
}

func (uf TurbidityFactory) FromDtoJSON(data []byte) (*Turbidity, error) {
	unitDto, err := TurbidityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TurbidityFactory{}.FromDto(*unitDto)
}


// FromNTU creates a new Turbidity instance from NTU.
func (uf TurbidityFactory) FromNTU(value float64) (*Turbidity, error) {
	return newTurbidity(value, TurbidityNTU)
}




// newTurbidity creates a new Turbidity.
func newTurbidity(value float64, fromUnit TurbidityUnits) (*Turbidity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Turbidity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Turbidity in NTU.
func (a *Turbidity) BaseValue() float64 {
	return a.value
}


// NTU returns the value in NTU.
func (a *Turbidity) NTU() float64 {
	if a.ntuLazy != nil {
		return *a.ntuLazy
	}
	ntu := a.convertFromBase(TurbidityNTU)
	a.ntuLazy = &ntu
	return ntu
}


// ToDto creates an TurbidityDto representation.
func (a *Turbidity) ToDto(holdInUnit *TurbidityUnits) TurbidityDto {
	if holdInUnit == nil {
		defaultUnit := TurbidityNTU // Default value
		holdInUnit = &defaultUnit
	}

	return TurbidityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an TurbidityDto representation.
func (a *Turbidity) ToDtoJSON(holdInUnit *TurbidityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Turbidity to a specific unit value.
func (a *Turbidity) Convert(toUnit TurbidityUnits) float64 {
	switch toUnit { 
    case TurbidityNTU:
		return a.NTU()
	default:
		return 0
	}
}

func (a *Turbidity) convertFromBase(toUnit TurbidityUnits) float64 {
    value := a.value
	switch toUnit { 
	case TurbidityNTU:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Turbidity) convertToBase(value float64, fromUnit TurbidityUnits) float64 {
	switch fromUnit { 
	case TurbidityNTU:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Turbidity) String() string {
	return a.ToString(TurbidityNTU, 2)
}

// ToString formats the Turbidity to string.
// fractionalDigits -1 for not mention
func (a *Turbidity) ToString(unit TurbidityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Turbidity) getUnitAbbreviation(unit TurbidityUnits) string {
	switch unit { 
	case TurbidityNTU:
		return "NTU" 
	default:
		return ""
	}
}

// Check if the given Turbidity are equals to the current Turbidity
func (a *Turbidity) Equals(other *Turbidity) bool {
	return a.value == other.BaseValue()
}

// Check if the given Turbidity are equals to the current Turbidity
func (a *Turbidity) CompareTo(other *Turbidity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Turbidity to the current Turbidity.
func (a *Turbidity) Add(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value + other.BaseValue()}
}

// Subtract the given Turbidity to the current Turbidity.
func (a *Turbidity) Subtract(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value - other.BaseValue()}
}

// Multiply the given Turbidity to the current Turbidity.
func (a *Turbidity) Multiply(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value * other.BaseValue()}
}

// Divide the given Turbidity to the current Turbidity.
func (a *Turbidity) Divide(other *Turbidity) *Turbidity {
	return &Turbidity{value: a.value / other.BaseValue()}
}