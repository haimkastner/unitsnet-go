// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ScalarUnits defines various units of Scalar.
type ScalarUnits string

const (
    
        // 
        ScalarAmount ScalarUnits = "Amount"
)

// ScalarDto represents a Scalar measurement with a numerical value and its corresponding unit.
type ScalarDto struct {
    // Value is the numerical representation of the Scalar.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Scalar, as defined in the ScalarUnits enumeration.
	Unit  ScalarUnits `json:"unit"`
}

// ScalarDtoFactory groups methods for creating and serializing ScalarDto objects.
type ScalarDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ScalarDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ScalarDtoFactory) FromJSON(data []byte) (*ScalarDto, error) {
	a := ScalarDto{}

    // Parse JSON into ScalarDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ScalarDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ScalarDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Scalar represents a measurement in a of Scalar.
//
// A way of representing a number of items.
type Scalar struct {
	// value is the base measurement stored internally.
	value       float64
    
    amountLazy *float64 
}

// ScalarFactory groups methods for creating Scalar instances.
type ScalarFactory struct{}

// CreateScalar creates a new Scalar instance from the given value and unit.
func (uf ScalarFactory) CreateScalar(value float64, unit ScalarUnits) (*Scalar, error) {
	return newScalar(value, unit)
}

// FromDto converts a ScalarDto to a Scalar instance.
func (uf ScalarFactory) FromDto(dto ScalarDto) (*Scalar, error) {
	return newScalar(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Scalar instance.
func (uf ScalarFactory) FromDtoJSON(data []byte) (*Scalar, error) {
	unitDto, err := ScalarDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ScalarDto from JSON: %w", err)
	}
	return ScalarFactory{}.FromDto(*unitDto)
}


// FromAmount creates a new Scalar instance from a value in Amount.
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

// BaseValue returns the base value of Scalar in Amount unit (the base/default quantity).
func (a *Scalar) BaseValue() float64 {
	return a.value
}


// Amount returns the Scalar value in Amount.
//
// 
func (a *Scalar) Amount() float64 {
	if a.amountLazy != nil {
		return *a.amountLazy
	}
	amount := a.convertFromBase(ScalarAmount)
	a.amountLazy = &amount
	return amount
}


// ToDto creates a ScalarDto representation from the Scalar instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Amount by default.
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

// ToDtoJSON creates a JSON representation of the Scalar instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Amount by default.
func (a *Scalar) ToDtoJSON(holdInUnit *ScalarUnits) ([]byte, error) {
	// Convert to ScalarDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Scalar to a specific unit value.
// The function uses the provided unit type (ScalarUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Scalar) Convert(toUnit ScalarUnits) float64 {
	switch toUnit { 
    case ScalarAmount:
		return a.Amount()
	default:
		return math.NaN()
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

// String returns a string representation of the Scalar in the default unit (Amount),
// formatted to two decimal places.
func (a Scalar) String() string {
	return a.ToString(ScalarAmount, 2)
}

// ToString formats the Scalar value as a string with the specified unit and fractional digits.
// It converts the Scalar to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Scalar value will be converted (e.g., Amount))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Scalar with the unit abbreviation.
func (a *Scalar) ToString(unit ScalarUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetScalarAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetScalarAbbreviation(unit))
}

// Equals checks if the given Scalar is equal to the current Scalar.
//
// Parameters:
//    other: The Scalar to compare against.
//
// Returns:
//    bool: Returns true if both Scalar are equal, false otherwise.
func (a *Scalar) Equals(other *Scalar) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Scalar with another Scalar.
// It returns -1 if the current Scalar is less than the other Scalar, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Scalar to compare against.
//
// Returns:
//    int: -1 if the current Scalar is less, 1 if greater, and 0 if equal.
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

// Add adds the given Scalar to the current Scalar and returns the result.
// The result is a new Scalar instance with the sum of the values.
//
// Parameters:
//    other: The Scalar to add to the current Scalar.
//
// Returns:
//    *Scalar: A new Scalar instance representing the sum of both Scalar.
func (a *Scalar) Add(other *Scalar) *Scalar {
	return &Scalar{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Scalar from the current Scalar and returns the result.
// The result is a new Scalar instance with the difference of the values.
//
// Parameters:
//    other: The Scalar to subtract from the current Scalar.
//
// Returns:
//    *Scalar: A new Scalar instance representing the difference of both Scalar.
func (a *Scalar) Subtract(other *Scalar) *Scalar {
	return &Scalar{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Scalar by the given Scalar and returns the result.
// The result is a new Scalar instance with the product of the values.
//
// Parameters:
//    other: The Scalar to multiply with the current Scalar.
//
// Returns:
//    *Scalar: A new Scalar instance representing the product of both Scalar.
func (a *Scalar) Multiply(other *Scalar) *Scalar {
	return &Scalar{value: a.value * other.BaseValue()}
}

// Divide divides the current Scalar by the given Scalar and returns the result.
// The result is a new Scalar instance with the quotient of the values.
//
// Parameters:
//    other: The Scalar to divide the current Scalar by.
//
// Returns:
//    *Scalar: A new Scalar instance representing the quotient of both Scalar.
func (a *Scalar) Divide(other *Scalar) *Scalar {
	return &Scalar{value: a.value / other.BaseValue()}
}

// GetScalarAbbreviation gets the unit abbreviation.
func GetScalarAbbreviation(unit ScalarUnits) string {
	switch unit { 
	case ScalarAmount:
		return "" 
	default:
		return ""
	}
}