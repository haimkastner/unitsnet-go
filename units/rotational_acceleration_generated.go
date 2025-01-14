// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalAccelerationUnits defines various units of RotationalAcceleration.
type RotationalAccelerationUnits string

const (
    
        // 
        RotationalAccelerationRadianPerSecondSquared RotationalAccelerationUnits = "RadianPerSecondSquared"
        // 
        RotationalAccelerationDegreePerSecondSquared RotationalAccelerationUnits = "DegreePerSecondSquared"
        // 
        RotationalAccelerationRevolutionPerMinutePerSecond RotationalAccelerationUnits = "RevolutionPerMinutePerSecond"
        // 
        RotationalAccelerationRevolutionPerSecondSquared RotationalAccelerationUnits = "RevolutionPerSecondSquared"
)

// RotationalAccelerationDto represents a RotationalAcceleration measurement with a numerical value and its corresponding unit.
type RotationalAccelerationDto struct {
    // Value is the numerical representation of the RotationalAcceleration.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RotationalAcceleration, as defined in the RotationalAccelerationUnits enumeration.
	Unit  RotationalAccelerationUnits `json:"unit"`
}

// RotationalAccelerationDtoFactory groups methods for creating and serializing RotationalAccelerationDto objects.
type RotationalAccelerationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RotationalAccelerationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RotationalAccelerationDtoFactory) FromJSON(data []byte) (*RotationalAccelerationDto, error) {
	a := RotationalAccelerationDto{}

    // Parse JSON into RotationalAccelerationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RotationalAccelerationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RotationalAccelerationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RotationalAcceleration represents a measurement in a of RotationalAcceleration.
//
// Angular acceleration is the rate of change of rotational speed.
type RotationalAcceleration struct {
	// value is the base measurement stored internally.
	value       float64
    
    radians_per_second_squaredLazy *float64 
    degrees_per_second_squaredLazy *float64 
    revolutions_per_minute_per_secondLazy *float64 
    revolutions_per_second_squaredLazy *float64 
}

// RotationalAccelerationFactory groups methods for creating RotationalAcceleration instances.
type RotationalAccelerationFactory struct{}

// CreateRotationalAcceleration creates a new RotationalAcceleration instance from the given value and unit.
func (uf RotationalAccelerationFactory) CreateRotationalAcceleration(value float64, unit RotationalAccelerationUnits) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, unit)
}

// FromDto converts a RotationalAccelerationDto to a RotationalAcceleration instance.
func (uf RotationalAccelerationFactory) FromDto(dto RotationalAccelerationDto) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RotationalAcceleration instance.
func (uf RotationalAccelerationFactory) FromDtoJSON(data []byte) (*RotationalAcceleration, error) {
	unitDto, err := RotationalAccelerationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RotationalAccelerationDto from JSON: %w", err)
	}
	return RotationalAccelerationFactory{}.FromDto(*unitDto)
}


// FromRadiansPerSecondSquared creates a new RotationalAcceleration instance from a value in RadiansPerSecondSquared.
func (uf RotationalAccelerationFactory) FromRadiansPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRadianPerSecondSquared)
}

// FromDegreesPerSecondSquared creates a new RotationalAcceleration instance from a value in DegreesPerSecondSquared.
func (uf RotationalAccelerationFactory) FromDegreesPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationDegreePerSecondSquared)
}

// FromRevolutionsPerMinutePerSecond creates a new RotationalAcceleration instance from a value in RevolutionsPerMinutePerSecond.
func (uf RotationalAccelerationFactory) FromRevolutionsPerMinutePerSecond(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRevolutionPerMinutePerSecond)
}

// FromRevolutionsPerSecondSquared creates a new RotationalAcceleration instance from a value in RevolutionsPerSecondSquared.
func (uf RotationalAccelerationFactory) FromRevolutionsPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRevolutionPerSecondSquared)
}


// newRotationalAcceleration creates a new RotationalAcceleration.
func newRotationalAcceleration(value float64, fromUnit RotationalAccelerationUnits) (*RotationalAcceleration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalAcceleration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalAcceleration in RadianPerSecondSquared unit (the base/default quantity).
func (a *RotationalAcceleration) BaseValue() float64 {
	return a.value
}


// RadiansPerSecondSquared returns the RotationalAcceleration value in RadiansPerSecondSquared.
//
// 
func (a *RotationalAcceleration) RadiansPerSecondSquared() float64 {
	if a.radians_per_second_squaredLazy != nil {
		return *a.radians_per_second_squaredLazy
	}
	radians_per_second_squared := a.convertFromBase(RotationalAccelerationRadianPerSecondSquared)
	a.radians_per_second_squaredLazy = &radians_per_second_squared
	return radians_per_second_squared
}

// DegreesPerSecondSquared returns the RotationalAcceleration value in DegreesPerSecondSquared.
//
// 
func (a *RotationalAcceleration) DegreesPerSecondSquared() float64 {
	if a.degrees_per_second_squaredLazy != nil {
		return *a.degrees_per_second_squaredLazy
	}
	degrees_per_second_squared := a.convertFromBase(RotationalAccelerationDegreePerSecondSquared)
	a.degrees_per_second_squaredLazy = &degrees_per_second_squared
	return degrees_per_second_squared
}

// RevolutionsPerMinutePerSecond returns the RotationalAcceleration value in RevolutionsPerMinutePerSecond.
//
// 
func (a *RotationalAcceleration) RevolutionsPerMinutePerSecond() float64 {
	if a.revolutions_per_minute_per_secondLazy != nil {
		return *a.revolutions_per_minute_per_secondLazy
	}
	revolutions_per_minute_per_second := a.convertFromBase(RotationalAccelerationRevolutionPerMinutePerSecond)
	a.revolutions_per_minute_per_secondLazy = &revolutions_per_minute_per_second
	return revolutions_per_minute_per_second
}

// RevolutionsPerSecondSquared returns the RotationalAcceleration value in RevolutionsPerSecondSquared.
//
// 
func (a *RotationalAcceleration) RevolutionsPerSecondSquared() float64 {
	if a.revolutions_per_second_squaredLazy != nil {
		return *a.revolutions_per_second_squaredLazy
	}
	revolutions_per_second_squared := a.convertFromBase(RotationalAccelerationRevolutionPerSecondSquared)
	a.revolutions_per_second_squaredLazy = &revolutions_per_second_squared
	return revolutions_per_second_squared
}


// ToDto creates a RotationalAccelerationDto representation from the RotationalAcceleration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by RadianPerSecondSquared by default.
func (a *RotationalAcceleration) ToDto(holdInUnit *RotationalAccelerationUnits) RotationalAccelerationDto {
	if holdInUnit == nil {
		defaultUnit := RotationalAccelerationRadianPerSecondSquared // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalAccelerationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RotationalAcceleration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by RadianPerSecondSquared by default.
func (a *RotationalAcceleration) ToDtoJSON(holdInUnit *RotationalAccelerationUnits) ([]byte, error) {
	// Convert to RotationalAccelerationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RotationalAcceleration to a specific unit value.
// The function uses the provided unit type (RotationalAccelerationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RotationalAcceleration) Convert(toUnit RotationalAccelerationUnits) float64 {
	switch toUnit { 
    case RotationalAccelerationRadianPerSecondSquared:
		return a.RadiansPerSecondSquared()
    case RotationalAccelerationDegreePerSecondSquared:
		return a.DegreesPerSecondSquared()
    case RotationalAccelerationRevolutionPerMinutePerSecond:
		return a.RevolutionsPerMinutePerSecond()
    case RotationalAccelerationRevolutionPerSecondSquared:
		return a.RevolutionsPerSecondSquared()
	default:
		return math.NaN()
	}
}

func (a *RotationalAcceleration) convertFromBase(toUnit RotationalAccelerationUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return (value) 
	case RotationalAccelerationDegreePerSecondSquared:
		return ((180 / math.Pi) * value) 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return ((60 / (2 * math.Pi)) * value) 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return ((1 / (2 * math.Pi)) * value) 
	default:
		return math.NaN()
	}
}

func (a *RotationalAcceleration) convertToBase(value float64, fromUnit RotationalAccelerationUnits) float64 {
	switch fromUnit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return (value) 
	case RotationalAccelerationDegreePerSecondSquared:
		return ((math.Pi / 180) * value) 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return (((2 * math.Pi) / 60) * value) 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return ((2 * math.Pi) * value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RotationalAcceleration in the default unit (RadianPerSecondSquared),
// formatted to two decimal places.
func (a RotationalAcceleration) String() string {
	return a.ToString(RotationalAccelerationRadianPerSecondSquared, 2)
}

// ToString formats the RotationalAcceleration value as a string with the specified unit and fractional digits.
// It converts the RotationalAcceleration to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RotationalAcceleration value will be converted (e.g., RadianPerSecondSquared))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RotationalAcceleration with the unit abbreviation.
func (a *RotationalAcceleration) ToString(unit RotationalAccelerationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalAcceleration) getUnitAbbreviation(unit RotationalAccelerationUnits) string {
	switch unit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return "rad/s²" 
	case RotationalAccelerationDegreePerSecondSquared:
		return "°/s²" 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return "rpm/s" 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return "r/s²" 
	default:
		return ""
	}
}

// Equals checks if the given RotationalAcceleration is equal to the current RotationalAcceleration.
//
// Parameters:
//    other: The RotationalAcceleration to compare against.
//
// Returns:
//    bool: Returns true if both RotationalAcceleration are equal, false otherwise.
func (a *RotationalAcceleration) Equals(other *RotationalAcceleration) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RotationalAcceleration with another RotationalAcceleration.
// It returns -1 if the current RotationalAcceleration is less than the other RotationalAcceleration, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RotationalAcceleration to compare against.
//
// Returns:
//    int: -1 if the current RotationalAcceleration is less, 1 if greater, and 0 if equal.
func (a *RotationalAcceleration) CompareTo(other *RotationalAcceleration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RotationalAcceleration to the current RotationalAcceleration and returns the result.
// The result is a new RotationalAcceleration instance with the sum of the values.
//
// Parameters:
//    other: The RotationalAcceleration to add to the current RotationalAcceleration.
//
// Returns:
//    *RotationalAcceleration: A new RotationalAcceleration instance representing the sum of both RotationalAcceleration.
func (a *RotationalAcceleration) Add(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RotationalAcceleration from the current RotationalAcceleration and returns the result.
// The result is a new RotationalAcceleration instance with the difference of the values.
//
// Parameters:
//    other: The RotationalAcceleration to subtract from the current RotationalAcceleration.
//
// Returns:
//    *RotationalAcceleration: A new RotationalAcceleration instance representing the difference of both RotationalAcceleration.
func (a *RotationalAcceleration) Subtract(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RotationalAcceleration by the given RotationalAcceleration and returns the result.
// The result is a new RotationalAcceleration instance with the product of the values.
//
// Parameters:
//    other: The RotationalAcceleration to multiply with the current RotationalAcceleration.
//
// Returns:
//    *RotationalAcceleration: A new RotationalAcceleration instance representing the product of both RotationalAcceleration.
func (a *RotationalAcceleration) Multiply(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value * other.BaseValue()}
}

// Divide divides the current RotationalAcceleration by the given RotationalAcceleration and returns the result.
// The result is a new RotationalAcceleration instance with the quotient of the values.
//
// Parameters:
//    other: The RotationalAcceleration to divide the current RotationalAcceleration by.
//
// Returns:
//    *RotationalAcceleration: A new RotationalAcceleration instance representing the quotient of both RotationalAcceleration.
func (a *RotationalAcceleration) Divide(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value / other.BaseValue()}
}