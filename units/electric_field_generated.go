// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricFieldUnits defines various units of ElectricField.
type ElectricFieldUnits string

const (
    
        // 
        ElectricFieldVoltPerMeter ElectricFieldUnits = "VoltPerMeter"
)

// ElectricFieldDto represents a ElectricField measurement with a numerical value and its corresponding unit.
type ElectricFieldDto struct {
    // Value is the numerical representation of the ElectricField.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricField, as defined in the ElectricFieldUnits enumeration.
	Unit  ElectricFieldUnits `json:"unit"`
}

// ElectricFieldDtoFactory groups methods for creating and serializing ElectricFieldDto objects.
type ElectricFieldDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricFieldDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricFieldDtoFactory) FromJSON(data []byte) (*ElectricFieldDto, error) {
	a := ElectricFieldDto{}

    // Parse JSON into ElectricFieldDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricFieldDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricFieldDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricField represents a measurement in a of ElectricField.
//
// An electric field is a force field that surrounds electric charges that attracts or repels other electric charges.
type ElectricField struct {
	// value is the base measurement stored internally.
	value       float64
    
    volts_per_meterLazy *float64 
}

// ElectricFieldFactory groups methods for creating ElectricField instances.
type ElectricFieldFactory struct{}

// CreateElectricField creates a new ElectricField instance from the given value and unit.
func (uf ElectricFieldFactory) CreateElectricField(value float64, unit ElectricFieldUnits) (*ElectricField, error) {
	return newElectricField(value, unit)
}

// FromDto converts a ElectricFieldDto to a ElectricField instance.
func (uf ElectricFieldFactory) FromDto(dto ElectricFieldDto) (*ElectricField, error) {
	return newElectricField(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricField instance.
func (uf ElectricFieldFactory) FromDtoJSON(data []byte) (*ElectricField, error) {
	unitDto, err := ElectricFieldDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricFieldDto from JSON: %w", err)
	}
	return ElectricFieldFactory{}.FromDto(*unitDto)
}


// FromVoltsPerMeter creates a new ElectricField instance from a value in VoltsPerMeter.
func (uf ElectricFieldFactory) FromVoltsPerMeter(value float64) (*ElectricField, error) {
	return newElectricField(value, ElectricFieldVoltPerMeter)
}


// newElectricField creates a new ElectricField.
func newElectricField(value float64, fromUnit ElectricFieldUnits) (*ElectricField, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricField{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricField in VoltPerMeter unit (the base/default quantity).
func (a *ElectricField) BaseValue() float64 {
	return a.value
}


// VoltsPerMeter returns the ElectricField value in VoltsPerMeter.
//
// 
func (a *ElectricField) VoltsPerMeter() float64 {
	if a.volts_per_meterLazy != nil {
		return *a.volts_per_meterLazy
	}
	volts_per_meter := a.convertFromBase(ElectricFieldVoltPerMeter)
	a.volts_per_meterLazy = &volts_per_meter
	return volts_per_meter
}


// ToDto creates a ElectricFieldDto representation from the ElectricField instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltPerMeter by default.
func (a *ElectricField) ToDto(holdInUnit *ElectricFieldUnits) ElectricFieldDto {
	if holdInUnit == nil {
		defaultUnit := ElectricFieldVoltPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricFieldDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricField instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltPerMeter by default.
func (a *ElectricField) ToDtoJSON(holdInUnit *ElectricFieldUnits) ([]byte, error) {
	// Convert to ElectricFieldDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricField to a specific unit value.
// The function uses the provided unit type (ElectricFieldUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricField) Convert(toUnit ElectricFieldUnits) float64 {
	switch toUnit { 
    case ElectricFieldVoltPerMeter:
		return a.VoltsPerMeter()
	default:
		return math.NaN()
	}
}

func (a *ElectricField) convertFromBase(toUnit ElectricFieldUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricFieldVoltPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *ElectricField) convertToBase(value float64, fromUnit ElectricFieldUnits) float64 {
	switch fromUnit { 
	case ElectricFieldVoltPerMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricField in the default unit (VoltPerMeter),
// formatted to two decimal places.
func (a ElectricField) String() string {
	return a.ToString(ElectricFieldVoltPerMeter, 2)
}

// ToString formats the ElectricField value as a string with the specified unit and fractional digits.
// It converts the ElectricField to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricField value will be converted (e.g., VoltPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricField with the unit abbreviation.
func (a *ElectricField) ToString(unit ElectricFieldUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricField) getUnitAbbreviation(unit ElectricFieldUnits) string {
	switch unit { 
	case ElectricFieldVoltPerMeter:
		return "V/m" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricField is equal to the current ElectricField.
//
// Parameters:
//    other: The ElectricField to compare against.
//
// Returns:
//    bool: Returns true if both ElectricField are equal, false otherwise.
func (a *ElectricField) Equals(other *ElectricField) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricField with another ElectricField.
// It returns -1 if the current ElectricField is less than the other ElectricField, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricField to compare against.
//
// Returns:
//    int: -1 if the current ElectricField is less, 1 if greater, and 0 if equal.
func (a *ElectricField) CompareTo(other *ElectricField) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricField to the current ElectricField and returns the result.
// The result is a new ElectricField instance with the sum of the values.
//
// Parameters:
//    other: The ElectricField to add to the current ElectricField.
//
// Returns:
//    *ElectricField: A new ElectricField instance representing the sum of both ElectricField.
func (a *ElectricField) Add(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricField from the current ElectricField and returns the result.
// The result is a new ElectricField instance with the difference of the values.
//
// Parameters:
//    other: The ElectricField to subtract from the current ElectricField.
//
// Returns:
//    *ElectricField: A new ElectricField instance representing the difference of both ElectricField.
func (a *ElectricField) Subtract(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricField by the given ElectricField and returns the result.
// The result is a new ElectricField instance with the product of the values.
//
// Parameters:
//    other: The ElectricField to multiply with the current ElectricField.
//
// Returns:
//    *ElectricField: A new ElectricField instance representing the product of both ElectricField.
func (a *ElectricField) Multiply(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricField by the given ElectricField and returns the result.
// The result is a new ElectricField instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricField to divide the current ElectricField by.
//
// Returns:
//    *ElectricField: A new ElectricField instance representing the quotient of both ElectricField.
func (a *ElectricField) Divide(other *ElectricField) *ElectricField {
	return &ElectricField{value: a.value / other.BaseValue()}
}