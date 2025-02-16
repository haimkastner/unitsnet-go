// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricChargeDensityUnits defines various units of ElectricChargeDensity.
type ElectricChargeDensityUnits string

const (
    
        // 
        ElectricChargeDensityCoulombPerCubicMeter ElectricChargeDensityUnits = "CoulombPerCubicMeter"
)

// ElectricChargeDensityDto represents a ElectricChargeDensity measurement with a numerical value and its corresponding unit.
type ElectricChargeDensityDto struct {
    // Value is the numerical representation of the ElectricChargeDensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricChargeDensity, as defined in the ElectricChargeDensityUnits enumeration.
	Unit  ElectricChargeDensityUnits `json:"unit"`
}

// ElectricChargeDensityDtoFactory groups methods for creating and serializing ElectricChargeDensityDto objects.
type ElectricChargeDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricChargeDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricChargeDensityDtoFactory) FromJSON(data []byte) (*ElectricChargeDensityDto, error) {
	a := ElectricChargeDensityDto{}

    // Parse JSON into ElectricChargeDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricChargeDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricChargeDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricChargeDensity represents a measurement in a of ElectricChargeDensity.
//
// In electromagnetism, charge density is a measure of the amount of electric charge per volume.
type ElectricChargeDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    coulombs_per_cubic_meterLazy *float64 
}

// ElectricChargeDensityFactory groups methods for creating ElectricChargeDensity instances.
type ElectricChargeDensityFactory struct{}

// CreateElectricChargeDensity creates a new ElectricChargeDensity instance from the given value and unit.
func (uf ElectricChargeDensityFactory) CreateElectricChargeDensity(value float64, unit ElectricChargeDensityUnits) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(value, unit)
}

// FromDto converts a ElectricChargeDensityDto to a ElectricChargeDensity instance.
func (uf ElectricChargeDensityFactory) FromDto(dto ElectricChargeDensityDto) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricChargeDensity instance.
func (uf ElectricChargeDensityFactory) FromDtoJSON(data []byte) (*ElectricChargeDensity, error) {
	unitDto, err := ElectricChargeDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricChargeDensityDto from JSON: %w", err)
	}
	return ElectricChargeDensityFactory{}.FromDto(*unitDto)
}


// FromCoulombsPerCubicMeter creates a new ElectricChargeDensity instance from a value in CoulombsPerCubicMeter.
func (uf ElectricChargeDensityFactory) FromCoulombsPerCubicMeter(value float64) (*ElectricChargeDensity, error) {
	return newElectricChargeDensity(value, ElectricChargeDensityCoulombPerCubicMeter)
}


// newElectricChargeDensity creates a new ElectricChargeDensity.
func newElectricChargeDensity(value float64, fromUnit ElectricChargeDensityUnits) (*ElectricChargeDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricChargeDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricChargeDensity in CoulombPerCubicMeter unit (the base/default quantity).
func (a *ElectricChargeDensity) BaseValue() float64 {
	return a.value
}


// CoulombsPerCubicMeter returns the ElectricChargeDensity value in CoulombsPerCubicMeter.
//
// 
func (a *ElectricChargeDensity) CoulombsPerCubicMeter() float64 {
	if a.coulombs_per_cubic_meterLazy != nil {
		return *a.coulombs_per_cubic_meterLazy
	}
	coulombs_per_cubic_meter := a.convertFromBase(ElectricChargeDensityCoulombPerCubicMeter)
	a.coulombs_per_cubic_meterLazy = &coulombs_per_cubic_meter
	return coulombs_per_cubic_meter
}


// ToDto creates a ElectricChargeDensityDto representation from the ElectricChargeDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerCubicMeter by default.
func (a *ElectricChargeDensity) ToDto(holdInUnit *ElectricChargeDensityUnits) ElectricChargeDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricChargeDensityCoulombPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricChargeDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricChargeDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerCubicMeter by default.
func (a *ElectricChargeDensity) ToDtoJSON(holdInUnit *ElectricChargeDensityUnits) ([]byte, error) {
	// Convert to ElectricChargeDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricChargeDensity to a specific unit value.
// The function uses the provided unit type (ElectricChargeDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricChargeDensity) Convert(toUnit ElectricChargeDensityUnits) float64 {
	switch toUnit { 
    case ElectricChargeDensityCoulombPerCubicMeter:
		return a.CoulombsPerCubicMeter()
	default:
		return math.NaN()
	}
}

func (a *ElectricChargeDensity) convertFromBase(toUnit ElectricChargeDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *ElectricChargeDensity) convertToBase(value float64, fromUnit ElectricChargeDensityUnits) float64 {
	switch fromUnit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricChargeDensity in the default unit (CoulombPerCubicMeter),
// formatted to two decimal places.
func (a ElectricChargeDensity) String() string {
	return a.ToString(ElectricChargeDensityCoulombPerCubicMeter, 2)
}

// ToString formats the ElectricChargeDensity value as a string with the specified unit and fractional digits.
// It converts the ElectricChargeDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricChargeDensity value will be converted (e.g., CoulombPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricChargeDensity with the unit abbreviation.
func (a *ElectricChargeDensity) ToString(unit ElectricChargeDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricChargeDensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricChargeDensityAbbreviation(unit))
}

// Equals checks if the given ElectricChargeDensity is equal to the current ElectricChargeDensity.
//
// Parameters:
//    other: The ElectricChargeDensity to compare against.
//
// Returns:
//    bool: Returns true if both ElectricChargeDensity are equal, false otherwise.
func (a *ElectricChargeDensity) Equals(other *ElectricChargeDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricChargeDensity with another ElectricChargeDensity.
// It returns -1 if the current ElectricChargeDensity is less than the other ElectricChargeDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricChargeDensity to compare against.
//
// Returns:
//    int: -1 if the current ElectricChargeDensity is less, 1 if greater, and 0 if equal.
func (a *ElectricChargeDensity) CompareTo(other *ElectricChargeDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricChargeDensity to the current ElectricChargeDensity and returns the result.
// The result is a new ElectricChargeDensity instance with the sum of the values.
//
// Parameters:
//    other: The ElectricChargeDensity to add to the current ElectricChargeDensity.
//
// Returns:
//    *ElectricChargeDensity: A new ElectricChargeDensity instance representing the sum of both ElectricChargeDensity.
func (a *ElectricChargeDensity) Add(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricChargeDensity from the current ElectricChargeDensity and returns the result.
// The result is a new ElectricChargeDensity instance with the difference of the values.
//
// Parameters:
//    other: The ElectricChargeDensity to subtract from the current ElectricChargeDensity.
//
// Returns:
//    *ElectricChargeDensity: A new ElectricChargeDensity instance representing the difference of both ElectricChargeDensity.
func (a *ElectricChargeDensity) Subtract(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricChargeDensity by the given ElectricChargeDensity and returns the result.
// The result is a new ElectricChargeDensity instance with the product of the values.
//
// Parameters:
//    other: The ElectricChargeDensity to multiply with the current ElectricChargeDensity.
//
// Returns:
//    *ElectricChargeDensity: A new ElectricChargeDensity instance representing the product of both ElectricChargeDensity.
func (a *ElectricChargeDensity) Multiply(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricChargeDensity by the given ElectricChargeDensity and returns the result.
// The result is a new ElectricChargeDensity instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricChargeDensity to divide the current ElectricChargeDensity by.
//
// Returns:
//    *ElectricChargeDensity: A new ElectricChargeDensity instance representing the quotient of both ElectricChargeDensity.
func (a *ElectricChargeDensity) Divide(other *ElectricChargeDensity) *ElectricChargeDensity {
	return &ElectricChargeDensity{value: a.value / other.BaseValue()}
}

// GetElectricChargeDensityAbbreviation gets the unit abbreviation.
func GetElectricChargeDensityAbbreviation(unit ElectricChargeDensityUnits) string {
	switch unit { 
	case ElectricChargeDensityCoulombPerCubicMeter:
		return "C/m³" 
	default:
		return ""
	}
}