// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RelativeHumidityUnits defines various units of RelativeHumidity.
type RelativeHumidityUnits string

const (
    
        // 
        RelativeHumidityPercent RelativeHumidityUnits = "Percent"
)

// RelativeHumidityDto represents a RelativeHumidity measurement with a numerical value and its corresponding unit.
type RelativeHumidityDto struct {
    // Value is the numerical representation of the RelativeHumidity.
	Value float64
    // Unit specifies the unit of measurement for the RelativeHumidity, as defined in the RelativeHumidityUnits enumeration.
	Unit  RelativeHumidityUnits
}

// RelativeHumidityDtoFactory groups methods for creating and serializing RelativeHumidityDto objects.
type RelativeHumidityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RelativeHumidityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RelativeHumidityDtoFactory) FromJSON(data []byte) (*RelativeHumidityDto, error) {
	a := RelativeHumidityDto{}

    // Parse JSON into RelativeHumidityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RelativeHumidityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RelativeHumidityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// RelativeHumidity represents a measurement in a of RelativeHumidity.
//
// Relative humidity is a ratio of the actual water vapor present in the air to the maximum water vapor in the air at the given temperature.
type RelativeHumidity struct {
	// value is the base measurement stored internally.
	value       float64
    
    percentLazy *float64 
}

// RelativeHumidityFactory groups methods for creating RelativeHumidity instances.
type RelativeHumidityFactory struct{}

// CreateRelativeHumidity creates a new RelativeHumidity instance from the given value and unit.
func (uf RelativeHumidityFactory) CreateRelativeHumidity(value float64, unit RelativeHumidityUnits) (*RelativeHumidity, error) {
	return newRelativeHumidity(value, unit)
}

// FromDto converts a RelativeHumidityDto to a RelativeHumidity instance.
func (uf RelativeHumidityFactory) FromDto(dto RelativeHumidityDto) (*RelativeHumidity, error) {
	return newRelativeHumidity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RelativeHumidity instance.
func (uf RelativeHumidityFactory) FromDtoJSON(data []byte) (*RelativeHumidity, error) {
	unitDto, err := RelativeHumidityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RelativeHumidityDto from JSON: %w", err)
	}
	return RelativeHumidityFactory{}.FromDto(*unitDto)
}


// FromPercent creates a new RelativeHumidity instance from a value in Percent.
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

// BaseValue returns the base value of RelativeHumidity in Percent unit (the base/default quantity).
func (a *RelativeHumidity) BaseValue() float64 {
	return a.value
}


// Percent returns the RelativeHumidity value in Percent.
//
// 
func (a *RelativeHumidity) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(RelativeHumidityPercent)
	a.percentLazy = &percent
	return percent
}


// ToDto creates a RelativeHumidityDto representation from the RelativeHumidity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Percent by default.
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

// ToDtoJSON creates a JSON representation of the RelativeHumidity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Percent by default.
func (a *RelativeHumidity) ToDtoJSON(holdInUnit *RelativeHumidityUnits) ([]byte, error) {
	// Convert to RelativeHumidityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RelativeHumidity to a specific unit value.
// The function uses the provided unit type (RelativeHumidityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RelativeHumidity) Convert(toUnit RelativeHumidityUnits) float64 {
	switch toUnit { 
    case RelativeHumidityPercent:
		return a.Percent()
	default:
		return math.NaN()
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

// String returns a string representation of the RelativeHumidity in the default unit (Percent),
// formatted to two decimal places.
func (a RelativeHumidity) String() string {
	return a.ToString(RelativeHumidityPercent, 2)
}

// ToString formats the RelativeHumidity value as a string with the specified unit and fractional digits.
// It converts the RelativeHumidity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RelativeHumidity value will be converted (e.g., Percent))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RelativeHumidity with the unit abbreviation.
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

// Equals checks if the given RelativeHumidity is equal to the current RelativeHumidity.
//
// Parameters:
//    other: The RelativeHumidity to compare against.
//
// Returns:
//    bool: Returns true if both RelativeHumidity are equal, false otherwise.
func (a *RelativeHumidity) Equals(other *RelativeHumidity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RelativeHumidity with another RelativeHumidity.
// It returns -1 if the current RelativeHumidity is less than the other RelativeHumidity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RelativeHumidity to compare against.
//
// Returns:
//    int: -1 if the current RelativeHumidity is less, 1 if greater, and 0 if equal.
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

// Add adds the given RelativeHumidity to the current RelativeHumidity and returns the result.
// The result is a new RelativeHumidity instance with the sum of the values.
//
// Parameters:
//    other: The RelativeHumidity to add to the current RelativeHumidity.
//
// Returns:
//    *RelativeHumidity: A new RelativeHumidity instance representing the sum of both RelativeHumidity.
func (a *RelativeHumidity) Add(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RelativeHumidity from the current RelativeHumidity and returns the result.
// The result is a new RelativeHumidity instance with the difference of the values.
//
// Parameters:
//    other: The RelativeHumidity to subtract from the current RelativeHumidity.
//
// Returns:
//    *RelativeHumidity: A new RelativeHumidity instance representing the difference of both RelativeHumidity.
func (a *RelativeHumidity) Subtract(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RelativeHumidity by the given RelativeHumidity and returns the result.
// The result is a new RelativeHumidity instance with the product of the values.
//
// Parameters:
//    other: The RelativeHumidity to multiply with the current RelativeHumidity.
//
// Returns:
//    *RelativeHumidity: A new RelativeHumidity instance representing the product of both RelativeHumidity.
func (a *RelativeHumidity) Multiply(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value * other.BaseValue()}
}

// Divide divides the current RelativeHumidity by the given RelativeHumidity and returns the result.
// The result is a new RelativeHumidity instance with the quotient of the values.
//
// Parameters:
//    other: The RelativeHumidity to divide the current RelativeHumidity by.
//
// Returns:
//    *RelativeHumidity: A new RelativeHumidity instance representing the quotient of both RelativeHumidity.
func (a *RelativeHumidity) Divide(other *RelativeHumidity) *RelativeHumidity {
	return &RelativeHumidity{value: a.value / other.BaseValue()}
}