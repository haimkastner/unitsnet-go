package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ScalarUnits enumeration
type ScalarUnits string

const (
    
        // 
        ScalarAmount ScalarUnits = "Amount"
)

// ScalarDto represents an Scalar
type ScalarDto struct {
	Value float64
	Unit  ScalarUnits
}

// ScalarDtoFactory struct to group related functions
type ScalarDtoFactory struct{}

func (udf ScalarDtoFactory) FromJSON(data []byte) (*ScalarDto, error) {
	a := ScalarDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ScalarDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Scalar struct
type Scalar struct {
	value       float64
    
    amountLazy *float64 
}

// ScalarFactory struct to group related functions
type ScalarFactory struct{}

func (uf ScalarFactory) CreateScalar(value float64, unit ScalarUnits) (*Scalar, error) {
	return newScalar(value, unit)
}

func (uf ScalarFactory) FromDto(dto ScalarDto) (*Scalar, error) {
	return newScalar(dto.Value, dto.Unit)
}

func (uf ScalarFactory) FromDtoJSON(data []byte) (*Scalar, error) {
	unitDto, err := ScalarDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ScalarFactory{}.FromDto(*unitDto)
}


// FromAmount creates a new Scalar instance from Amount.
func (uf ScalarFactory) FromAmount(value float64) (*Scalar, error) {
	return newScalar(value, ScalarAmount)
}




// newScalar creates a new Scalar.
func newScalar(value float64, fromUnit ScalarUnits) (*Scalar, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Scalar{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Scalar in Amount.
func (a *Scalar) BaseValue() float64 {
	return a.value
}


// Amount returns the value in Amount.
func (a *Scalar) Amount() float64 {
	if a.amountLazy != nil {
		return *a.amountLazy
	}
	amount := a.convertFromBase(ScalarAmount)
	a.amountLazy = &amount
	return amount
}


// ToDto creates an ScalarDto representation.
func (a *Scalar) ToDto(holdInUnit *ScalarUnits) ScalarDto {
	if holdInUnit == nil {
		defaultUnit := ScalarAmount // Default value
		holdInUnit = &defaultUnit
	}

	return ScalarDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ScalarDto representation.
func (a *Scalar) ToDtoJSON(holdInUnit *ScalarUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Scalar to a specific unit value.
func (a *Scalar) Convert(toUnit ScalarUnits) float64 {
	switch toUnit { 
    case ScalarAmount:
		return a.Amount()
	default:
		return 0
	}
}

func (a *Scalar) convertFromBase(toUnit ScalarUnits) float64 {
    value := a.value
	switch toUnit { 
	case ScalarAmount:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *Scalar) convertToBase(value float64, fromUnit ScalarUnits) float64 {
	switch fromUnit { 
	case ScalarAmount:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Scalar) String() string {
	return a.ToString(ScalarAmount, 2)
}

// ToString formats the Scalar to string.
// fractionalDigits -1 for not mention
func (a *Scalar) ToString(unit ScalarUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Scalar) getUnitAbbreviation(unit ScalarUnits) string {
	switch unit { 
	case ScalarAmount:
		return "" 
	default:
		return ""
	}
}

// Check if the given Scalar are equals to the current Scalar
func (a *Scalar) Equals(other *Scalar) bool {
	return a.value == other.BaseValue()
}

// Check if the given Scalar are equals to the current Scalar
func (a *Scalar) CompareTo(other *Scalar) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Scalar to the current Scalar.
func (a *Scalar) Add(other *Scalar) *Scalar {
	return &Scalar{value: a.value + other.BaseValue()}
}

// Subtract the given Scalar to the current Scalar.
func (a *Scalar) Subtract(other *Scalar) *Scalar {
	return &Scalar{value: a.value - other.BaseValue()}
}

// Multiply the given Scalar to the current Scalar.
func (a *Scalar) Multiply(other *Scalar) *Scalar {
	return &Scalar{value: a.value * other.BaseValue()}
}

// Divide the given Scalar to the current Scalar.
func (a *Scalar) Divide(other *Scalar) *Scalar {
	return &Scalar{value: a.value / other.BaseValue()}
}