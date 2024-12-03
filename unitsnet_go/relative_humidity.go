package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RelativeHumidityUnits enumeration
type RelativeHumidityUnits string

const (
    
        // 
        RelativeHumidityPercent RelativeHumidityUnits = "Percent"
)

// RelativeHumidityDto represents an RelativeHumidity
type RelativeHumidityDto struct {
	Value float64
	Unit  RelativeHumidityUnits
}

// RelativeHumidityDtoFactory struct to group related functions
type RelativeHumidityDtoFactory struct{}

func (udf RelativeHumidityDtoFactory) FromJSON(data []byte) (*RelativeHumidityDto, error) {
	a := RelativeHumidityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RelativeHumidityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RelativeHumidity struct
type RelativeHumidity struct {
	value       float64
    
    percentLazy *float64 
}

// RelativeHumidityFactory struct to group related functions
type RelativeHumidityFactory struct{}

func (uf RelativeHumidityFactory) CreateRelativeHumidity(value float64, unit RelativeHumidityUnits) (*RelativeHumidity, error) {
	return newRelativeHumidity(value, unit)
}

func (uf RelativeHumidityFactory) FromDto(dto RelativeHumidityDto) (*RelativeHumidity, error) {
	return newRelativeHumidity(dto.Value, dto.Unit)
}

func (uf RelativeHumidityFactory) FromDtoJSON(data []byte) (*RelativeHumidity, error) {
	unitDto, err := RelativeHumidityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RelativeHumidityFactory{}.FromDto(*unitDto)
}


// FromPercent creates a new RelativeHumidity instance from Percent.
func (uf RelativeHumidityFactory) FromPercent(value float64) (*RelativeHumidity, error) {
	return newRelativeHumidity(value, RelativeHumidityPercent)
}




// newRelativeHumidity creates a new RelativeHumidity.
func newRelativeHumidity(value float64, fromUnit RelativeHumidityUnits) (*RelativeHumidity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RelativeHumidity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RelativeHumidity in Percent.
func (a *RelativeHumidity) BaseValue() float64 {
	return a.value
}


// Percent returns the value in Percent.
func (a *RelativeHumidity) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(RelativeHumidityPercent)
	a.percentLazy = &percent
	return percent
}


// ToDto creates an RelativeHumidityDto representation.
func (a *RelativeHumidity) ToDto(holdInUnit *RelativeHumidityUnits) RelativeHumidityDto {
	if holdInUnit == nil {
		defaultUnit := RelativeHumidityPercent // Default value
		holdInUnit = &defaultUnit
	}

	return RelativeHumidityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RelativeHumidityDto representation.
func (a *RelativeHumidity) ToDtoJSON(holdInUnit *RelativeHumidityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RelativeHumidity to a specific unit value.
func (a *RelativeHumidity) Convert(toUnit RelativeHumidityUnits) float64 {
	switch toUnit { 
    case RelativeHumidityPercent:
		return a.Percent()
	default:
		return 0
	}
}

func (a *RelativeHumidity) convertFromBase(toUnit RelativeHumidityUnits) float64 {
    value := a.value
	switch toUnit { 
	case RelativeHumidityPercent:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *RelativeHumidity) convertToBase(value float64, fromUnit RelativeHumidityUnits) float64 {
	switch fromUnit { 
	case RelativeHumidityPercent:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RelativeHumidity) String() string {
	return a.ToString(RelativeHumidityPercent, 2)
}

// ToString formats the RelativeHumidity to string.
// fractionalDigits -1 for not mention
func (a *RelativeHumidity) ToString(unit RelativeHumidityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RelativeHumidity) getUnitAbbreviation(unit RelativeHumidityUnits) string {
	switch unit { 
	case RelativeHumidityPercent:
		return "%RH" 
	default:
		return ""
	}
}

// Check if the given RelativeHumidity are equals to the current RelativeHumidity
func (a *RelativeHumidity) Equals(other *RelativeHumidity) bool {
	return a.value == other.BaseValue()
}

// Check if the given RelativeHumidity are equals to the current RelativeHumidity
func (a *RelativeHumidity) CompareTo(other *RelativeHumidity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RelativeHumidity to the current RelativeHumidity.
func (a *RelativeHumidity) Add(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value + other.BaseValue()}
}

// Subtract the given RelativeHumidity to the current RelativeHumidity.
func (a *RelativeHumidity) Subtract(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value - other.BaseValue()}
}

// Multiply the given RelativeHumidity to the current RelativeHumidity.
func (a *RelativeHumidity) Multiply(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value * other.BaseValue()}
}

// Divide the given RelativeHumidity to the current RelativeHumidity.
func (a *RelativeHumidity) Divide(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value / other.BaseValue()}
}