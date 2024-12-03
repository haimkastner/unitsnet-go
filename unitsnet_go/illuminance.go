package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// IlluminanceUnits enumeration
type IlluminanceUnits string

const (
    
        // 
        IlluminanceLux IlluminanceUnits = "Lux"
        // 
        IlluminanceMillilux IlluminanceUnits = "Millilux"
        // 
        IlluminanceKilolux IlluminanceUnits = "Kilolux"
        // 
        IlluminanceMegalux IlluminanceUnits = "Megalux"
)

// IlluminanceDto represents an Illuminance
type IlluminanceDto struct {
	Value float64
	Unit  IlluminanceUnits
}

// IlluminanceDtoFactory struct to group related functions
type IlluminanceDtoFactory struct{}

func (udf IlluminanceDtoFactory) FromJSON(data []byte) (*IlluminanceDto, error) {
	a := IlluminanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a IlluminanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Illuminance struct
type Illuminance struct {
	value       float64
    
    luxLazy *float64 
    milliluxLazy *float64 
    kiloluxLazy *float64 
    megaluxLazy *float64 
}

// IlluminanceFactory struct to group related functions
type IlluminanceFactory struct{}

func (uf IlluminanceFactory) CreateIlluminance(value float64, unit IlluminanceUnits) (*Illuminance, error) {
	return newIlluminance(value, unit)
}

func (uf IlluminanceFactory) FromDto(dto IlluminanceDto) (*Illuminance, error) {
	return newIlluminance(dto.Value, dto.Unit)
}

func (uf IlluminanceFactory) FromDtoJSON(data []byte) (*Illuminance, error) {
	unitDto, err := IlluminanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return IlluminanceFactory{}.FromDto(*unitDto)
}


// FromLux creates a new Illuminance instance from Lux.
func (uf IlluminanceFactory) FromLux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceLux)
}

// FromMillilux creates a new Illuminance instance from Millilux.
func (uf IlluminanceFactory) FromMillilux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceMillilux)
}

// FromKilolux creates a new Illuminance instance from Kilolux.
func (uf IlluminanceFactory) FromKilolux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceKilolux)
}

// FromMegalux creates a new Illuminance instance from Megalux.
func (uf IlluminanceFactory) FromMegalux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceMegalux)
}




// newIlluminance creates a new Illuminance.
func newIlluminance(value float64, fromUnit IlluminanceUnits) (*Illuminance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Illuminance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Illuminance in Lux.
func (a *Illuminance) BaseValue() float64 {
	return a.value
}


// Lux returns the value in Lux.
func (a *Illuminance) Lux() float64 {
	if a.luxLazy != nil {
		return *a.luxLazy
	}
	lux := a.convertFromBase(IlluminanceLux)
	a.luxLazy = &lux
	return lux
}

// Millilux returns the value in Millilux.
func (a *Illuminance) Millilux() float64 {
	if a.milliluxLazy != nil {
		return *a.milliluxLazy
	}
	millilux := a.convertFromBase(IlluminanceMillilux)
	a.milliluxLazy = &millilux
	return millilux
}

// Kilolux returns the value in Kilolux.
func (a *Illuminance) Kilolux() float64 {
	if a.kiloluxLazy != nil {
		return *a.kiloluxLazy
	}
	kilolux := a.convertFromBase(IlluminanceKilolux)
	a.kiloluxLazy = &kilolux
	return kilolux
}

// Megalux returns the value in Megalux.
func (a *Illuminance) Megalux() float64 {
	if a.megaluxLazy != nil {
		return *a.megaluxLazy
	}
	megalux := a.convertFromBase(IlluminanceMegalux)
	a.megaluxLazy = &megalux
	return megalux
}


// ToDto creates an IlluminanceDto representation.
func (a *Illuminance) ToDto(holdInUnit *IlluminanceUnits) IlluminanceDto {
	if holdInUnit == nil {
		defaultUnit := IlluminanceLux // Default value
		holdInUnit = &defaultUnit
	}

	return IlluminanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an IlluminanceDto representation.
func (a *Illuminance) ToDtoJSON(holdInUnit *IlluminanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Illuminance to a specific unit value.
func (a *Illuminance) Convert(toUnit IlluminanceUnits) float64 {
	switch toUnit { 
    case IlluminanceLux:
		return a.Lux()
    case IlluminanceMillilux:
		return a.Millilux()
    case IlluminanceKilolux:
		return a.Kilolux()
    case IlluminanceMegalux:
		return a.Megalux()
	default:
		return 0
	}
}

func (a *Illuminance) convertFromBase(toUnit IlluminanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case IlluminanceLux:
		return (value) 
	case IlluminanceMillilux:
		return ((value) / 0.001) 
	case IlluminanceKilolux:
		return ((value) / 1000.0) 
	case IlluminanceMegalux:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Illuminance) convertToBase(value float64, fromUnit IlluminanceUnits) float64 {
	switch fromUnit { 
	case IlluminanceLux:
		return (value) 
	case IlluminanceMillilux:
		return ((value) * 0.001) 
	case IlluminanceKilolux:
		return ((value) * 1000.0) 
	case IlluminanceMegalux:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Illuminance) String() string {
	return a.ToString(IlluminanceLux, 2)
}

// ToString formats the Illuminance to string.
// fractionalDigits -1 for not mention
func (a *Illuminance) ToString(unit IlluminanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Illuminance) getUnitAbbreviation(unit IlluminanceUnits) string {
	switch unit { 
	case IlluminanceLux:
		return "lx" 
	case IlluminanceMillilux:
		return "mlx" 
	case IlluminanceKilolux:
		return "klx" 
	case IlluminanceMegalux:
		return "Mlx" 
	default:
		return ""
	}
}

// Check if the given Illuminance are equals to the current Illuminance
func (a *Illuminance) Equals(other *Illuminance) bool {
	return a.value == other.BaseValue()
}

// Check if the given Illuminance are equals to the current Illuminance
func (a *Illuminance) CompareTo(other *Illuminance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Illuminance to the current Illuminance.
func (a *Illuminance) Add(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value + other.BaseValue()}
}

// Subtract the given Illuminance to the current Illuminance.
func (a *Illuminance) Subtract(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value - other.BaseValue()}
}

// Multiply the given Illuminance to the current Illuminance.
func (a *Illuminance) Multiply(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value * other.BaseValue()}
}

// Divide the given Illuminance to the current Illuminance.
func (a *Illuminance) Divide(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value / other.BaseValue()}
}