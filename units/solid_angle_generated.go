// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SolidAngleUnits defines various units of SolidAngle.
type SolidAngleUnits string

const (
    
        // 
        SolidAngleSteradian SolidAngleUnits = "Steradian"
)

// SolidAngleDto represents a SolidAngle measurement with a numerical value and its corresponding unit.
type SolidAngleDto struct {
    // Value is the numerical representation of the SolidAngle.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the SolidAngle, as defined in the SolidAngleUnits enumeration.
	Unit  SolidAngleUnits `json:"unit"`
}

// SolidAngleDtoFactory groups methods for creating and serializing SolidAngleDto objects.
type SolidAngleDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SolidAngleDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SolidAngleDtoFactory) FromJSON(data []byte) (*SolidAngleDto, error) {
	a := SolidAngleDto{}

    // Parse JSON into SolidAngleDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a SolidAngleDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SolidAngleDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SolidAngle represents a measurement in a of SolidAngle.
//
// In geometry, a solid angle is the two-dimensional angle in three-dimensional space that an object subtends at a point.
type SolidAngle struct {
	// value is the base measurement stored internally.
	value       float64
    
    steradiansLazy *float64 
}

// SolidAngleFactory groups methods for creating SolidAngle instances.
type SolidAngleFactory struct{}

// CreateSolidAngle creates a new SolidAngle instance from the given value and unit.
func (uf SolidAngleFactory) CreateSolidAngle(value float64, unit SolidAngleUnits) (*SolidAngle, error) {
	return newSolidAngle(value, unit)
}

// FromDto converts a SolidAngleDto to a SolidAngle instance.
func (uf SolidAngleFactory) FromDto(dto SolidAngleDto) (*SolidAngle, error) {
	return newSolidAngle(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SolidAngle instance.
func (uf SolidAngleFactory) FromDtoJSON(data []byte) (*SolidAngle, error) {
	unitDto, err := SolidAngleDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SolidAngleDto from JSON: %w", err)
	}
	return SolidAngleFactory{}.FromDto(*unitDto)
}


// FromSteradians creates a new SolidAngle instance from a value in Steradians.
func (uf SolidAngleFactory) FromSteradians(value float64) (*SolidAngle, error) {
	return newSolidAngle(value, SolidAngleSteradian)
}


// newSolidAngle creates a new SolidAngle.
func newSolidAngle(value float64, fromUnit SolidAngleUnits) (*SolidAngle, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SolidAngle{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SolidAngle in Steradian unit (the base/default quantity).
func (a *SolidAngle) BaseValue() float64 {
	return a.value
}


// Steradians returns the SolidAngle value in Steradians.
//
// 
func (a *SolidAngle) Steradians() float64 {
	if a.steradiansLazy != nil {
		return *a.steradiansLazy
	}
	steradians := a.convertFromBase(SolidAngleSteradian)
	a.steradiansLazy = &steradians
	return steradians
}


// ToDto creates a SolidAngleDto representation from the SolidAngle instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Steradian by default.
func (a *SolidAngle) ToDto(holdInUnit *SolidAngleUnits) SolidAngleDto {
	if holdInUnit == nil {
		defaultUnit := SolidAngleSteradian // Default value
		holdInUnit = &defaultUnit
	}

	return SolidAngleDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the SolidAngle instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Steradian by default.
func (a *SolidAngle) ToDtoJSON(holdInUnit *SolidAngleUnits) ([]byte, error) {
	// Convert to SolidAngleDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SolidAngle to a specific unit value.
// The function uses the provided unit type (SolidAngleUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *SolidAngle) Convert(toUnit SolidAngleUnits) float64 {
	switch toUnit { 
    case SolidAngleSteradian:
		return a.Steradians()
	default:
		return math.NaN()
	}
}

func (a *SolidAngle) convertFromBase(toUnit SolidAngleUnits) float64 {
    value := a.value
	switch toUnit { 
	case SolidAngleSteradian:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *SolidAngle) convertToBase(value float64, fromUnit SolidAngleUnits) float64 {
	switch fromUnit { 
	case SolidAngleSteradian:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the SolidAngle in the default unit (Steradian),
// formatted to two decimal places.
func (a SolidAngle) String() string {
	return a.ToString(SolidAngleSteradian, 2)
}

// ToString formats the SolidAngle value as a string with the specified unit and fractional digits.
// It converts the SolidAngle to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SolidAngle value will be converted (e.g., Steradian))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SolidAngle with the unit abbreviation.
func (a *SolidAngle) ToString(unit SolidAngleUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSolidAngleAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSolidAngleAbbreviation(unit))
}

// Equals checks if the given SolidAngle is equal to the current SolidAngle.
//
// Parameters:
//    other: The SolidAngle to compare against.
//
// Returns:
//    bool: Returns true if both SolidAngle are equal, false otherwise.
func (a *SolidAngle) Equals(other *SolidAngle) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SolidAngle with another SolidAngle.
// It returns -1 if the current SolidAngle is less than the other SolidAngle, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SolidAngle to compare against.
//
// Returns:
//    int: -1 if the current SolidAngle is less, 1 if greater, and 0 if equal.
func (a *SolidAngle) CompareTo(other *SolidAngle) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given SolidAngle to the current SolidAngle and returns the result.
// The result is a new SolidAngle instance with the sum of the values.
//
// Parameters:
//    other: The SolidAngle to add to the current SolidAngle.
//
// Returns:
//    *SolidAngle: A new SolidAngle instance representing the sum of both SolidAngle.
func (a *SolidAngle) Add(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SolidAngle from the current SolidAngle and returns the result.
// The result is a new SolidAngle instance with the difference of the values.
//
// Parameters:
//    other: The SolidAngle to subtract from the current SolidAngle.
//
// Returns:
//    *SolidAngle: A new SolidAngle instance representing the difference of both SolidAngle.
func (a *SolidAngle) Subtract(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SolidAngle by the given SolidAngle and returns the result.
// The result is a new SolidAngle instance with the product of the values.
//
// Parameters:
//    other: The SolidAngle to multiply with the current SolidAngle.
//
// Returns:
//    *SolidAngle: A new SolidAngle instance representing the product of both SolidAngle.
func (a *SolidAngle) Multiply(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value * other.BaseValue()}
}

// Divide divides the current SolidAngle by the given SolidAngle and returns the result.
// The result is a new SolidAngle instance with the quotient of the values.
//
// Parameters:
//    other: The SolidAngle to divide the current SolidAngle by.
//
// Returns:
//    *SolidAngle: A new SolidAngle instance representing the quotient of both SolidAngle.
func (a *SolidAngle) Divide(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value / other.BaseValue()}
}

// GetSolidAngleAbbreviation gets the unit abbreviation.
func GetSolidAngleAbbreviation(unit SolidAngleUnits) string {
	switch unit { 
	case SolidAngleSteradian:
		return "sr" 
	default:
		return ""
	}
}