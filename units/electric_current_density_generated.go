// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentDensityUnits defines various units of ElectricCurrentDensity.
type ElectricCurrentDensityUnits string

const (
    
        // 
        ElectricCurrentDensityAmperePerSquareMeter ElectricCurrentDensityUnits = "AmperePerSquareMeter"
        // 
        ElectricCurrentDensityAmperePerSquareInch ElectricCurrentDensityUnits = "AmperePerSquareInch"
        // 
        ElectricCurrentDensityAmperePerSquareFoot ElectricCurrentDensityUnits = "AmperePerSquareFoot"
)

// ElectricCurrentDensityDto represents a ElectricCurrentDensity measurement with a numerical value and its corresponding unit.
type ElectricCurrentDensityDto struct {
    // Value is the numerical representation of the ElectricCurrentDensity.
	Value float64
    // Unit specifies the unit of measurement for the ElectricCurrentDensity, as defined in the ElectricCurrentDensityUnits enumeration.
	Unit  ElectricCurrentDensityUnits
}

// ElectricCurrentDensityDtoFactory groups methods for creating and serializing ElectricCurrentDensityDto objects.
type ElectricCurrentDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrentDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricCurrentDensityDtoFactory) FromJSON(data []byte) (*ElectricCurrentDensityDto, error) {
	a := ElectricCurrentDensityDto{}

    // Parse JSON into ElectricCurrentDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricCurrentDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricCurrentDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ElectricCurrentDensity represents a measurement in a of ElectricCurrentDensity.
//
// In electromagnetism, current density is the electric current per unit area of cross section.
type ElectricCurrentDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    amperes_per_square_meterLazy *float64 
    amperes_per_square_inchLazy *float64 
    amperes_per_square_footLazy *float64 
}

// ElectricCurrentDensityFactory groups methods for creating ElectricCurrentDensity instances.
type ElectricCurrentDensityFactory struct{}

// CreateElectricCurrentDensity creates a new ElectricCurrentDensity instance from the given value and unit.
func (uf ElectricCurrentDensityFactory) CreateElectricCurrentDensity(value float64, unit ElectricCurrentDensityUnits) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, unit)
}

// FromDto converts a ElectricCurrentDensityDto to a ElectricCurrentDensity instance.
func (uf ElectricCurrentDensityFactory) FromDto(dto ElectricCurrentDensityDto) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrentDensity instance.
func (uf ElectricCurrentDensityFactory) FromDtoJSON(data []byte) (*ElectricCurrentDensity, error) {
	unitDto, err := ElectricCurrentDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricCurrentDensityDto from JSON: %w", err)
	}
	return ElectricCurrentDensityFactory{}.FromDto(*unitDto)
}


// FromAmperesPerSquareMeter creates a new ElectricCurrentDensity instance from a value in AmperesPerSquareMeter.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareMeter(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareMeter)
}

// FromAmperesPerSquareInch creates a new ElectricCurrentDensity instance from a value in AmperesPerSquareInch.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareInch(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareInch)
}

// FromAmperesPerSquareFoot creates a new ElectricCurrentDensity instance from a value in AmperesPerSquareFoot.
func (uf ElectricCurrentDensityFactory) FromAmperesPerSquareFoot(value float64) (*ElectricCurrentDensity, error) {
	return newElectricCurrentDensity(value, ElectricCurrentDensityAmperePerSquareFoot)
}


// newElectricCurrentDensity creates a new ElectricCurrentDensity.
func newElectricCurrentDensity(value float64, fromUnit ElectricCurrentDensityUnits) (*ElectricCurrentDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrentDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrentDensity in AmperePerSquareMeter unit (the base/default quantity).
func (a *ElectricCurrentDensity) BaseValue() float64 {
	return a.value
}


// AmperesPerSquareMeter returns the ElectricCurrentDensity value in AmperesPerSquareMeter.
//
// 
func (a *ElectricCurrentDensity) AmperesPerSquareMeter() float64 {
	if a.amperes_per_square_meterLazy != nil {
		return *a.amperes_per_square_meterLazy
	}
	amperes_per_square_meter := a.convertFromBase(ElectricCurrentDensityAmperePerSquareMeter)
	a.amperes_per_square_meterLazy = &amperes_per_square_meter
	return amperes_per_square_meter
}

// AmperesPerSquareInch returns the ElectricCurrentDensity value in AmperesPerSquareInch.
//
// 
func (a *ElectricCurrentDensity) AmperesPerSquareInch() float64 {
	if a.amperes_per_square_inchLazy != nil {
		return *a.amperes_per_square_inchLazy
	}
	amperes_per_square_inch := a.convertFromBase(ElectricCurrentDensityAmperePerSquareInch)
	a.amperes_per_square_inchLazy = &amperes_per_square_inch
	return amperes_per_square_inch
}

// AmperesPerSquareFoot returns the ElectricCurrentDensity value in AmperesPerSquareFoot.
//
// 
func (a *ElectricCurrentDensity) AmperesPerSquareFoot() float64 {
	if a.amperes_per_square_footLazy != nil {
		return *a.amperes_per_square_footLazy
	}
	amperes_per_square_foot := a.convertFromBase(ElectricCurrentDensityAmperePerSquareFoot)
	a.amperes_per_square_footLazy = &amperes_per_square_foot
	return amperes_per_square_foot
}


// ToDto creates a ElectricCurrentDensityDto representation from the ElectricCurrentDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerSquareMeter by default.
func (a *ElectricCurrentDensity) ToDto(holdInUnit *ElectricCurrentDensityUnits) ElectricCurrentDensityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentDensityAmperePerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricCurrentDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by AmperePerSquareMeter by default.
func (a *ElectricCurrentDensity) ToDtoJSON(holdInUnit *ElectricCurrentDensityUnits) ([]byte, error) {
	// Convert to ElectricCurrentDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricCurrentDensity to a specific unit value.
// The function uses the provided unit type (ElectricCurrentDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricCurrentDensity) Convert(toUnit ElectricCurrentDensityUnits) float64 {
	switch toUnit { 
    case ElectricCurrentDensityAmperePerSquareMeter:
		return a.AmperesPerSquareMeter()
    case ElectricCurrentDensityAmperePerSquareInch:
		return a.AmperesPerSquareInch()
    case ElectricCurrentDensityAmperePerSquareFoot:
		return a.AmperesPerSquareFoot()
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentDensity) convertFromBase(toUnit ElectricCurrentDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return (value) 
	case ElectricCurrentDensityAmperePerSquareInch:
		return (value / 1.5500031000062000e3) 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return (value / 1.0763910416709722e1) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentDensity) convertToBase(value float64, fromUnit ElectricCurrentDensityUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return (value) 
	case ElectricCurrentDensityAmperePerSquareInch:
		return (value * 1.5500031000062000e3) 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return (value * 1.0763910416709722e1) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricCurrentDensity in the default unit (AmperePerSquareMeter),
// formatted to two decimal places.
func (a ElectricCurrentDensity) String() string {
	return a.ToString(ElectricCurrentDensityAmperePerSquareMeter, 2)
}

// ToString formats the ElectricCurrentDensity value as a string with the specified unit and fractional digits.
// It converts the ElectricCurrentDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricCurrentDensity value will be converted (e.g., AmperePerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricCurrentDensity with the unit abbreviation.
func (a *ElectricCurrentDensity) ToString(unit ElectricCurrentDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCurrentDensity) getUnitAbbreviation(unit ElectricCurrentDensityUnits) string {
	switch unit { 
	case ElectricCurrentDensityAmperePerSquareMeter:
		return "A/m²" 
	case ElectricCurrentDensityAmperePerSquareInch:
		return "A/in²" 
	case ElectricCurrentDensityAmperePerSquareFoot:
		return "A/ft²" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricCurrentDensity is equal to the current ElectricCurrentDensity.
//
// Parameters:
//    other: The ElectricCurrentDensity to compare against.
//
// Returns:
//    bool: Returns true if both ElectricCurrentDensity are equal, false otherwise.
func (a *ElectricCurrentDensity) Equals(other *ElectricCurrentDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricCurrentDensity with another ElectricCurrentDensity.
// It returns -1 if the current ElectricCurrentDensity is less than the other ElectricCurrentDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricCurrentDensity to compare against.
//
// Returns:
//    int: -1 if the current ElectricCurrentDensity is less, 1 if greater, and 0 if equal.
func (a *ElectricCurrentDensity) CompareTo(other *ElectricCurrentDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricCurrentDensity to the current ElectricCurrentDensity and returns the result.
// The result is a new ElectricCurrentDensity instance with the sum of the values.
//
// Parameters:
//    other: The ElectricCurrentDensity to add to the current ElectricCurrentDensity.
//
// Returns:
//    *ElectricCurrentDensity: A new ElectricCurrentDensity instance representing the sum of both ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Add(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricCurrentDensity from the current ElectricCurrentDensity and returns the result.
// The result is a new ElectricCurrentDensity instance with the difference of the values.
//
// Parameters:
//    other: The ElectricCurrentDensity to subtract from the current ElectricCurrentDensity.
//
// Returns:
//    *ElectricCurrentDensity: A new ElectricCurrentDensity instance representing the difference of both ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Subtract(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricCurrentDensity by the given ElectricCurrentDensity and returns the result.
// The result is a new ElectricCurrentDensity instance with the product of the values.
//
// Parameters:
//    other: The ElectricCurrentDensity to multiply with the current ElectricCurrentDensity.
//
// Returns:
//    *ElectricCurrentDensity: A new ElectricCurrentDensity instance representing the product of both ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Multiply(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricCurrentDensity by the given ElectricCurrentDensity and returns the result.
// The result is a new ElectricCurrentDensity instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricCurrentDensity to divide the current ElectricCurrentDensity by.
//
// Returns:
//    *ElectricCurrentDensity: A new ElectricCurrentDensity instance representing the quotient of both ElectricCurrentDensity.
func (a *ElectricCurrentDensity) Divide(other *ElectricCurrentDensity) *ElectricCurrentDensity {
	return &ElectricCurrentDensity{value: a.value / other.BaseValue()}
}