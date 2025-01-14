// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricSurfaceChargeDensityUnits defines various units of ElectricSurfaceChargeDensity.
type ElectricSurfaceChargeDensityUnits string

const (
    
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareMeter ElectricSurfaceChargeDensityUnits = "CoulombPerSquareMeter"
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareCentimeter ElectricSurfaceChargeDensityUnits = "CoulombPerSquareCentimeter"
        // 
        ElectricSurfaceChargeDensityCoulombPerSquareInch ElectricSurfaceChargeDensityUnits = "CoulombPerSquareInch"
)

// ElectricSurfaceChargeDensityDto represents a ElectricSurfaceChargeDensity measurement with a numerical value and its corresponding unit.
type ElectricSurfaceChargeDensityDto struct {
    // Value is the numerical representation of the ElectricSurfaceChargeDensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricSurfaceChargeDensity, as defined in the ElectricSurfaceChargeDensityUnits enumeration.
	Unit  ElectricSurfaceChargeDensityUnits `json:"unit"`
}

// ElectricSurfaceChargeDensityDtoFactory groups methods for creating and serializing ElectricSurfaceChargeDensityDto objects.
type ElectricSurfaceChargeDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricSurfaceChargeDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricSurfaceChargeDensityDtoFactory) FromJSON(data []byte) (*ElectricSurfaceChargeDensityDto, error) {
	a := ElectricSurfaceChargeDensityDto{}

    // Parse JSON into ElectricSurfaceChargeDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricSurfaceChargeDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricSurfaceChargeDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricSurfaceChargeDensity represents a measurement in a of ElectricSurfaceChargeDensity.
//
// In electromagnetism, surface charge density is a measure of the amount of electric charge per surface area.
type ElectricSurfaceChargeDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    coulombs_per_square_meterLazy *float64 
    coulombs_per_square_centimeterLazy *float64 
    coulombs_per_square_inchLazy *float64 
}

// ElectricSurfaceChargeDensityFactory groups methods for creating ElectricSurfaceChargeDensity instances.
type ElectricSurfaceChargeDensityFactory struct{}

// CreateElectricSurfaceChargeDensity creates a new ElectricSurfaceChargeDensity instance from the given value and unit.
func (uf ElectricSurfaceChargeDensityFactory) CreateElectricSurfaceChargeDensity(value float64, unit ElectricSurfaceChargeDensityUnits) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, unit)
}

// FromDto converts a ElectricSurfaceChargeDensityDto to a ElectricSurfaceChargeDensity instance.
func (uf ElectricSurfaceChargeDensityFactory) FromDto(dto ElectricSurfaceChargeDensityDto) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricSurfaceChargeDensity instance.
func (uf ElectricSurfaceChargeDensityFactory) FromDtoJSON(data []byte) (*ElectricSurfaceChargeDensity, error) {
	unitDto, err := ElectricSurfaceChargeDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricSurfaceChargeDensityDto from JSON: %w", err)
	}
	return ElectricSurfaceChargeDensityFactory{}.FromDto(*unitDto)
}


// FromCoulombsPerSquareMeter creates a new ElectricSurfaceChargeDensity instance from a value in CoulombsPerSquareMeter.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareMeter(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareMeter)
}

// FromCoulombsPerSquareCentimeter creates a new ElectricSurfaceChargeDensity instance from a value in CoulombsPerSquareCentimeter.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareCentimeter(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
}

// FromCoulombsPerSquareInch creates a new ElectricSurfaceChargeDensity instance from a value in CoulombsPerSquareInch.
func (uf ElectricSurfaceChargeDensityFactory) FromCoulombsPerSquareInch(value float64) (*ElectricSurfaceChargeDensity, error) {
	return newElectricSurfaceChargeDensity(value, ElectricSurfaceChargeDensityCoulombPerSquareInch)
}


// newElectricSurfaceChargeDensity creates a new ElectricSurfaceChargeDensity.
func newElectricSurfaceChargeDensity(value float64, fromUnit ElectricSurfaceChargeDensityUnits) (*ElectricSurfaceChargeDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricSurfaceChargeDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricSurfaceChargeDensity in CoulombPerSquareMeter unit (the base/default quantity).
func (a *ElectricSurfaceChargeDensity) BaseValue() float64 {
	return a.value
}


// CoulombsPerSquareMeter returns the ElectricSurfaceChargeDensity value in CoulombsPerSquareMeter.
//
// 
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareMeter() float64 {
	if a.coulombs_per_square_meterLazy != nil {
		return *a.coulombs_per_square_meterLazy
	}
	coulombs_per_square_meter := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareMeter)
	a.coulombs_per_square_meterLazy = &coulombs_per_square_meter
	return coulombs_per_square_meter
}

// CoulombsPerSquareCentimeter returns the ElectricSurfaceChargeDensity value in CoulombsPerSquareCentimeter.
//
// 
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareCentimeter() float64 {
	if a.coulombs_per_square_centimeterLazy != nil {
		return *a.coulombs_per_square_centimeterLazy
	}
	coulombs_per_square_centimeter := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareCentimeter)
	a.coulombs_per_square_centimeterLazy = &coulombs_per_square_centimeter
	return coulombs_per_square_centimeter
}

// CoulombsPerSquareInch returns the ElectricSurfaceChargeDensity value in CoulombsPerSquareInch.
//
// 
func (a *ElectricSurfaceChargeDensity) CoulombsPerSquareInch() float64 {
	if a.coulombs_per_square_inchLazy != nil {
		return *a.coulombs_per_square_inchLazy
	}
	coulombs_per_square_inch := a.convertFromBase(ElectricSurfaceChargeDensityCoulombPerSquareInch)
	a.coulombs_per_square_inchLazy = &coulombs_per_square_inch
	return coulombs_per_square_inch
}


// ToDto creates a ElectricSurfaceChargeDensityDto representation from the ElectricSurfaceChargeDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerSquareMeter by default.
func (a *ElectricSurfaceChargeDensity) ToDto(holdInUnit *ElectricSurfaceChargeDensityUnits) ElectricSurfaceChargeDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricSurfaceChargeDensityCoulombPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricSurfaceChargeDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricSurfaceChargeDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerSquareMeter by default.
func (a *ElectricSurfaceChargeDensity) ToDtoJSON(holdInUnit *ElectricSurfaceChargeDensityUnits) ([]byte, error) {
	// Convert to ElectricSurfaceChargeDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricSurfaceChargeDensity to a specific unit value.
// The function uses the provided unit type (ElectricSurfaceChargeDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricSurfaceChargeDensity) Convert(toUnit ElectricSurfaceChargeDensityUnits) float64 {
	switch toUnit { 
    case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return a.CoulombsPerSquareMeter()
    case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return a.CoulombsPerSquareCentimeter()
    case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return a.CoulombsPerSquareInch()
	default:
		return math.NaN()
	}
}

func (a *ElectricSurfaceChargeDensity) convertFromBase(toUnit ElectricSurfaceChargeDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return (value) 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return (value / 1.0e4) 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return (value / 1.5500031000062000e3) 
	default:
		return math.NaN()
	}
}

func (a *ElectricSurfaceChargeDensity) convertToBase(value float64, fromUnit ElectricSurfaceChargeDensityUnits) float64 {
	switch fromUnit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return (value) 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return (value * 1.0e4) 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return (value * 1.5500031000062000e3) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricSurfaceChargeDensity in the default unit (CoulombPerSquareMeter),
// formatted to two decimal places.
func (a ElectricSurfaceChargeDensity) String() string {
	return a.ToString(ElectricSurfaceChargeDensityCoulombPerSquareMeter, 2)
}

// ToString formats the ElectricSurfaceChargeDensity value as a string with the specified unit and fractional digits.
// It converts the ElectricSurfaceChargeDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricSurfaceChargeDensity value will be converted (e.g., CoulombPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricSurfaceChargeDensity with the unit abbreviation.
func (a *ElectricSurfaceChargeDensity) ToString(unit ElectricSurfaceChargeDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricSurfaceChargeDensity) getUnitAbbreviation(unit ElectricSurfaceChargeDensityUnits) string {
	switch unit { 
	case ElectricSurfaceChargeDensityCoulombPerSquareMeter:
		return "C/m²" 
	case ElectricSurfaceChargeDensityCoulombPerSquareCentimeter:
		return "C/cm²" 
	case ElectricSurfaceChargeDensityCoulombPerSquareInch:
		return "C/in²" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricSurfaceChargeDensity is equal to the current ElectricSurfaceChargeDensity.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to compare against.
//
// Returns:
//    bool: Returns true if both ElectricSurfaceChargeDensity are equal, false otherwise.
func (a *ElectricSurfaceChargeDensity) Equals(other *ElectricSurfaceChargeDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricSurfaceChargeDensity with another ElectricSurfaceChargeDensity.
// It returns -1 if the current ElectricSurfaceChargeDensity is less than the other ElectricSurfaceChargeDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to compare against.
//
// Returns:
//    int: -1 if the current ElectricSurfaceChargeDensity is less, 1 if greater, and 0 if equal.
func (a *ElectricSurfaceChargeDensity) CompareTo(other *ElectricSurfaceChargeDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricSurfaceChargeDensity to the current ElectricSurfaceChargeDensity and returns the result.
// The result is a new ElectricSurfaceChargeDensity instance with the sum of the values.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to add to the current ElectricSurfaceChargeDensity.
//
// Returns:
//    *ElectricSurfaceChargeDensity: A new ElectricSurfaceChargeDensity instance representing the sum of both ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Add(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricSurfaceChargeDensity from the current ElectricSurfaceChargeDensity and returns the result.
// The result is a new ElectricSurfaceChargeDensity instance with the difference of the values.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to subtract from the current ElectricSurfaceChargeDensity.
//
// Returns:
//    *ElectricSurfaceChargeDensity: A new ElectricSurfaceChargeDensity instance representing the difference of both ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Subtract(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricSurfaceChargeDensity by the given ElectricSurfaceChargeDensity and returns the result.
// The result is a new ElectricSurfaceChargeDensity instance with the product of the values.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to multiply with the current ElectricSurfaceChargeDensity.
//
// Returns:
//    *ElectricSurfaceChargeDensity: A new ElectricSurfaceChargeDensity instance representing the product of both ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Multiply(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricSurfaceChargeDensity by the given ElectricSurfaceChargeDensity and returns the result.
// The result is a new ElectricSurfaceChargeDensity instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricSurfaceChargeDensity to divide the current ElectricSurfaceChargeDensity by.
//
// Returns:
//    *ElectricSurfaceChargeDensity: A new ElectricSurfaceChargeDensity instance representing the quotient of both ElectricSurfaceChargeDensity.
func (a *ElectricSurfaceChargeDensity) Divide(other *ElectricSurfaceChargeDensity) *ElectricSurfaceChargeDensity {
	return &ElectricSurfaceChargeDensity{value: a.value / other.BaseValue()}
}