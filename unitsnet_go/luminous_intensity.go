package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LuminousIntensityUnits enumeration
type LuminousIntensityUnits string

const (
    
        // 
        LuminousIntensityCandela LuminousIntensityUnits = "Candela"
)

// LuminousIntensityDto represents an LuminousIntensity
type LuminousIntensityDto struct {
	Value float64
	Unit  LuminousIntensityUnits
}

// LuminousIntensityDtoFactory struct to group related functions
type LuminousIntensityDtoFactory struct{}

func (udf LuminousIntensityDtoFactory) FromJSON(data []byte) (*LuminousIntensityDto, error) {
	a := LuminousIntensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LuminousIntensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// LuminousIntensity struct
type LuminousIntensity struct {
	value       float64
    
    candelaLazy *float64 
}

// LuminousIntensityFactory struct to group related functions
type LuminousIntensityFactory struct{}

func (uf LuminousIntensityFactory) CreateLuminousIntensity(value float64, unit LuminousIntensityUnits) (*LuminousIntensity, error) {
	return newLuminousIntensity(value, unit)
}

func (uf LuminousIntensityFactory) FromDto(dto LuminousIntensityDto) (*LuminousIntensity, error) {
	return newLuminousIntensity(dto.Value, dto.Unit)
}

func (uf LuminousIntensityFactory) FromDtoJSON(data []byte) (*LuminousIntensity, error) {
	unitDto, err := LuminousIntensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LuminousIntensityFactory{}.FromDto(*unitDto)
}


// FromCandela creates a new LuminousIntensity instance from Candela.
func (uf LuminousIntensityFactory) FromCandela(value float64) (*LuminousIntensity, error) {
	return newLuminousIntensity(value, LuminousIntensityCandela)
}




// newLuminousIntensity creates a new LuminousIntensity.
func newLuminousIntensity(value float64, fromUnit LuminousIntensityUnits) (*LuminousIntensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LuminousIntensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LuminousIntensity in Candela.
func (a *LuminousIntensity) BaseValue() float64 {
	return a.value
}


// Candela returns the value in Candela.
func (a *LuminousIntensity) Candela() float64 {
	if a.candelaLazy != nil {
		return *a.candelaLazy
	}
	candela := a.convertFromBase(LuminousIntensityCandela)
	a.candelaLazy = &candela
	return candela
}


// ToDto creates an LuminousIntensityDto representation.
func (a *LuminousIntensity) ToDto(holdInUnit *LuminousIntensityUnits) LuminousIntensityDto {
	if holdInUnit == nil {
		defaultUnit := LuminousIntensityCandela // Default value
		holdInUnit = &defaultUnit
	}

	return LuminousIntensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LuminousIntensityDto representation.
func (a *LuminousIntensity) ToDtoJSON(holdInUnit *LuminousIntensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts LuminousIntensity to a specific unit value.
func (a *LuminousIntensity) Convert(toUnit LuminousIntensityUnits) float64 {
	switch toUnit { 
    case LuminousIntensityCandela:
		return a.Candela()
	default:
		return 0
	}
}

func (a *LuminousIntensity) convertFromBase(toUnit LuminousIntensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminousIntensityCandela:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *LuminousIntensity) convertToBase(value float64, fromUnit LuminousIntensityUnits) float64 {
	switch fromUnit { 
	case LuminousIntensityCandela:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a LuminousIntensity) String() string {
	return a.ToString(LuminousIntensityCandela, 2)
}

// ToString formats the LuminousIntensity to string.
// fractionalDigits -1 for not mention
func (a *LuminousIntensity) ToString(unit LuminousIntensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LuminousIntensity) getUnitAbbreviation(unit LuminousIntensityUnits) string {
	switch unit { 
	case LuminousIntensityCandela:
		return "cd" 
	default:
		return ""
	}
}

// Check if the given LuminousIntensity are equals to the current LuminousIntensity
func (a *LuminousIntensity) Equals(other *LuminousIntensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given LuminousIntensity are equals to the current LuminousIntensity
func (a *LuminousIntensity) CompareTo(other *LuminousIntensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given LuminousIntensity to the current LuminousIntensity.
func (a *LuminousIntensity) Add(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value + other.BaseValue()}
}

// Subtract the given LuminousIntensity to the current LuminousIntensity.
func (a *LuminousIntensity) Subtract(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value - other.BaseValue()}
}

// Multiply the given LuminousIntensity to the current LuminousIntensity.
func (a *LuminousIntensity) Multiply(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value * other.BaseValue()}
}

// Divide the given LuminousIntensity to the current LuminousIntensity.
func (a *LuminousIntensity) Divide(other *LuminousIntensity) *LuminousIntensity {
	return &LuminousIntensity{value: a.value / other.BaseValue()}
}