// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VitaminAUnits defines various units of VitaminA.
type VitaminAUnits string

const (
    
        // 
        VitaminAInternationalUnit VitaminAUnits = "InternationalUnit"
)

// VitaminADto represents a VitaminA measurement with a numerical value and its corresponding unit.
type VitaminADto struct {
    // Value is the numerical representation of the VitaminA.
	Value float64
    // Unit specifies the unit of measurement for the VitaminA, as defined in the VitaminAUnits enumeration.
	Unit  VitaminAUnits
}

// VitaminADtoFactory groups methods for creating and serializing VitaminADto objects.
type VitaminADtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VitaminADto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VitaminADtoFactory) FromJSON(data []byte) (*VitaminADto, error) {
	a := VitaminADto{}

    // Parse JSON into VitaminADto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a VitaminADto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VitaminADto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// VitaminA represents a measurement in a of VitaminA.
//
// Vitamin A: 1 IU is the biological equivalent of 0.3 µg retinol, or of 0.6 µg beta-carotene.
type VitaminA struct {
	// value is the base measurement stored internally.
	value       float64
    
    international_unitsLazy *float64 
}

// VitaminAFactory groups methods for creating VitaminA instances.
type VitaminAFactory struct{}

// CreateVitaminA creates a new VitaminA instance from the given value and unit.
func (uf VitaminAFactory) CreateVitaminA(value float64, unit VitaminAUnits) (*VitaminA, error) {
	return newVitaminA(value, unit)
}

// FromDto converts a VitaminADto to a VitaminA instance.
func (uf VitaminAFactory) FromDto(dto VitaminADto) (*VitaminA, error) {
	return newVitaminA(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VitaminA instance.
func (uf VitaminAFactory) FromDtoJSON(data []byte) (*VitaminA, error) {
	unitDto, err := VitaminADtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VitaminADto from JSON: %w", err)
	}
	return VitaminAFactory{}.FromDto(*unitDto)
}


// FromInternationalUnits creates a new VitaminA instance from a value in InternationalUnits.
func (uf VitaminAFactory) FromInternationalUnits(value float64) (*VitaminA, error) {
	return newVitaminA(value, VitaminAInternationalUnit)
}


// newVitaminA creates a new VitaminA.
func newVitaminA(value float64, fromUnit VitaminAUnits) (*VitaminA, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VitaminA{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VitaminA in InternationalUnit unit (the base/default quantity).
func (a *VitaminA) BaseValue() float64 {
	return a.value
}


// InternationalUnits returns the VitaminA value in InternationalUnits.
//
// 
func (a *VitaminA) InternationalUnits() float64 {
	if a.international_unitsLazy != nil {
		return *a.international_unitsLazy
	}
	international_units := a.convertFromBase(VitaminAInternationalUnit)
	a.international_unitsLazy = &international_units
	return international_units
}


// ToDto creates a VitaminADto representation from the VitaminA instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InternationalUnit by default.
func (a *VitaminA) ToDto(holdInUnit *VitaminAUnits) VitaminADto {
	if holdInUnit == nil {
		defaultUnit := VitaminAInternationalUnit // Default value
		holdInUnit = &defaultUnit
	}

	return VitaminADto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VitaminA instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InternationalUnit by default.
func (a *VitaminA) ToDtoJSON(holdInUnit *VitaminAUnits) ([]byte, error) {
	// Convert to VitaminADto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VitaminA to a specific unit value.
// The function uses the provided unit type (VitaminAUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VitaminA) Convert(toUnit VitaminAUnits) float64 {
	switch toUnit { 
    case VitaminAInternationalUnit:
		return a.InternationalUnits()
	default:
		return math.NaN()
	}
}

func (a *VitaminA) convertFromBase(toUnit VitaminAUnits) float64 {
    value := a.value
	switch toUnit { 
	case VitaminAInternationalUnit:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *VitaminA) convertToBase(value float64, fromUnit VitaminAUnits) float64 {
	switch fromUnit { 
	case VitaminAInternationalUnit:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VitaminA in the default unit (InternationalUnit),
// formatted to two decimal places.
func (a VitaminA) String() string {
	return a.ToString(VitaminAInternationalUnit, 2)
}

// ToString formats the VitaminA value as a string with the specified unit and fractional digits.
// It converts the VitaminA to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VitaminA value will be converted (e.g., InternationalUnit))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VitaminA with the unit abbreviation.
func (a *VitaminA) ToString(unit VitaminAUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VitaminA) getUnitAbbreviation(unit VitaminAUnits) string {
	switch unit { 
	case VitaminAInternationalUnit:
		return "IU" 
	default:
		return ""
	}
}

// Equals checks if the given VitaminA is equal to the current VitaminA.
//
// Parameters:
//    other: The VitaminA to compare against.
//
// Returns:
//    bool: Returns true if both VitaminA are equal, false otherwise.
func (a *VitaminA) Equals(other *VitaminA) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VitaminA with another VitaminA.
// It returns -1 if the current VitaminA is less than the other VitaminA, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VitaminA to compare against.
//
// Returns:
//    int: -1 if the current VitaminA is less, 1 if greater, and 0 if equal.
func (a *VitaminA) CompareTo(other *VitaminA) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VitaminA to the current VitaminA and returns the result.
// The result is a new VitaminA instance with the sum of the values.
//
// Parameters:
//    other: The VitaminA to add to the current VitaminA.
//
// Returns:
//    *VitaminA: A new VitaminA instance representing the sum of both VitaminA.
func (a *VitaminA) Add(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VitaminA from the current VitaminA and returns the result.
// The result is a new VitaminA instance with the difference of the values.
//
// Parameters:
//    other: The VitaminA to subtract from the current VitaminA.
//
// Returns:
//    *VitaminA: A new VitaminA instance representing the difference of both VitaminA.
func (a *VitaminA) Subtract(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VitaminA by the given VitaminA and returns the result.
// The result is a new VitaminA instance with the product of the values.
//
// Parameters:
//    other: The VitaminA to multiply with the current VitaminA.
//
// Returns:
//    *VitaminA: A new VitaminA instance representing the product of both VitaminA.
func (a *VitaminA) Multiply(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value * other.BaseValue()}
}

// Divide divides the current VitaminA by the given VitaminA and returns the result.
// The result is a new VitaminA instance with the quotient of the values.
//
// Parameters:
//    other: The VitaminA to divide the current VitaminA by.
//
// Returns:
//    *VitaminA: A new VitaminA instance representing the quotient of both VitaminA.
func (a *VitaminA) Divide(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value / other.BaseValue()}
}