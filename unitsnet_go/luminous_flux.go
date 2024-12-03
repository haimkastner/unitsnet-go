package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LuminousFluxUnits enumeration
type LuminousFluxUnits string

const (
    
        // 
        LuminousFluxLumen LuminousFluxUnits = "Lumen"
)

// LuminousFluxDto represents an LuminousFlux
type LuminousFluxDto struct {
	Value float64
	Unit  LuminousFluxUnits
}

// LuminousFluxDtoFactory struct to group related functions
type LuminousFluxDtoFactory struct{}

func (udf LuminousFluxDtoFactory) FromJSON(data []byte) (*LuminousFluxDto, error) {
	a := LuminousFluxDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LuminousFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// LuminousFlux struct
type LuminousFlux struct {
	value       float64
    
    lumensLazy *float64 
}

// LuminousFluxFactory struct to group related functions
type LuminousFluxFactory struct{}

func (uf LuminousFluxFactory) CreateLuminousFlux(value float64, unit LuminousFluxUnits) (*LuminousFlux, error) {
	return newLuminousFlux(value, unit)
}

func (uf LuminousFluxFactory) FromDto(dto LuminousFluxDto) (*LuminousFlux, error) {
	return newLuminousFlux(dto.Value, dto.Unit)
}

func (uf LuminousFluxFactory) FromDtoJSON(data []byte) (*LuminousFlux, error) {
	unitDto, err := LuminousFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LuminousFluxFactory{}.FromDto(*unitDto)
}


// FromLumen creates a new LuminousFlux instance from Lumen.
func (uf LuminousFluxFactory) FromLumens(value float64) (*LuminousFlux, error) {
	return newLuminousFlux(value, LuminousFluxLumen)
}




// newLuminousFlux creates a new LuminousFlux.
func newLuminousFlux(value float64, fromUnit LuminousFluxUnits) (*LuminousFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LuminousFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LuminousFlux in Lumen.
func (a *LuminousFlux) BaseValue() float64 {
	return a.value
}


// Lumen returns the value in Lumen.
func (a *LuminousFlux) Lumens() float64 {
	if a.lumensLazy != nil {
		return *a.lumensLazy
	}
	lumens := a.convertFromBase(LuminousFluxLumen)
	a.lumensLazy = &lumens
	return lumens
}


// ToDto creates an LuminousFluxDto representation.
func (a *LuminousFlux) ToDto(holdInUnit *LuminousFluxUnits) LuminousFluxDto {
	if holdInUnit == nil {
		defaultUnit := LuminousFluxLumen // Default value
		holdInUnit = &defaultUnit
	}

	return LuminousFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LuminousFluxDto representation.
func (a *LuminousFlux) ToDtoJSON(holdInUnit *LuminousFluxUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts LuminousFlux to a specific unit value.
func (a *LuminousFlux) Convert(toUnit LuminousFluxUnits) float64 {
	switch toUnit { 
    case LuminousFluxLumen:
		return a.Lumens()
	default:
		return 0
	}
}

func (a *LuminousFlux) convertFromBase(toUnit LuminousFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminousFluxLumen:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *LuminousFlux) convertToBase(value float64, fromUnit LuminousFluxUnits) float64 {
	switch fromUnit { 
	case LuminousFluxLumen:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a LuminousFlux) String() string {
	return a.ToString(LuminousFluxLumen, 2)
}

// ToString formats the LuminousFlux to string.
// fractionalDigits -1 for not mention
func (a *LuminousFlux) ToString(unit LuminousFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LuminousFlux) getUnitAbbreviation(unit LuminousFluxUnits) string {
	switch unit { 
	case LuminousFluxLumen:
		return "lm" 
	default:
		return ""
	}
}

// Check if the given LuminousFlux are equals to the current LuminousFlux
func (a *LuminousFlux) Equals(other *LuminousFlux) bool {
	return a.value == other.BaseValue()
}

// Check if the given LuminousFlux are equals to the current LuminousFlux
func (a *LuminousFlux) CompareTo(other *LuminousFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given LuminousFlux to the current LuminousFlux.
func (a *LuminousFlux) Add(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value + other.BaseValue()}
}

// Subtract the given LuminousFlux to the current LuminousFlux.
func (a *LuminousFlux) Subtract(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value - other.BaseValue()}
}

// Multiply the given LuminousFlux to the current LuminousFlux.
func (a *LuminousFlux) Multiply(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value * other.BaseValue()}
}

// Divide the given LuminousFlux to the current LuminousFlux.
func (a *LuminousFlux) Divide(other *LuminousFlux) *LuminousFlux {
	return &LuminousFlux{value: a.value / other.BaseValue()}
}