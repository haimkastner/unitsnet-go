// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// CompressibilityUnits defines various units of Compressibility.
type CompressibilityUnits string

const (
    
        // 
        CompressibilityInversePascal CompressibilityUnits = "InversePascal"
        // 
        CompressibilityInverseKilopascal CompressibilityUnits = "InverseKilopascal"
        // 
        CompressibilityInverseMegapascal CompressibilityUnits = "InverseMegapascal"
        // 
        CompressibilityInverseAtmosphere CompressibilityUnits = "InverseAtmosphere"
        // 
        CompressibilityInverseMillibar CompressibilityUnits = "InverseMillibar"
        // 
        CompressibilityInverseBar CompressibilityUnits = "InverseBar"
        // 
        CompressibilityInversePoundForcePerSquareInch CompressibilityUnits = "InversePoundForcePerSquareInch"
)

// CompressibilityDto represents a Compressibility measurement with a numerical value and its corresponding unit.
type CompressibilityDto struct {
    // Value is the numerical representation of the Compressibility.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Compressibility, as defined in the CompressibilityUnits enumeration.
	Unit  CompressibilityUnits `json:"unit"`
}

// CompressibilityDtoFactory groups methods for creating and serializing CompressibilityDto objects.
type CompressibilityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a CompressibilityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf CompressibilityDtoFactory) FromJSON(data []byte) (*CompressibilityDto, error) {
	a := CompressibilityDto{}

    // Parse JSON into CompressibilityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a CompressibilityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a CompressibilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Compressibility represents a measurement in a of Compressibility.
//
// None
type Compressibility struct {
	// value is the base measurement stored internally.
	value       float64
    
    inverse_pascalsLazy *float64 
    inverse_kilopascalsLazy *float64 
    inverse_megapascalsLazy *float64 
    inverse_atmospheresLazy *float64 
    inverse_millibarsLazy *float64 
    inverse_barsLazy *float64 
    inverse_pounds_force_per_square_inchLazy *float64 
}

// CompressibilityFactory groups methods for creating Compressibility instances.
type CompressibilityFactory struct{}

// CreateCompressibility creates a new Compressibility instance from the given value and unit.
func (uf CompressibilityFactory) CreateCompressibility(value float64, unit CompressibilityUnits) (*Compressibility, error) {
	return newCompressibility(value, unit)
}

// FromDto converts a CompressibilityDto to a Compressibility instance.
func (uf CompressibilityFactory) FromDto(dto CompressibilityDto) (*Compressibility, error) {
	return newCompressibility(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Compressibility instance.
func (uf CompressibilityFactory) FromDtoJSON(data []byte) (*Compressibility, error) {
	unitDto, err := CompressibilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CompressibilityDto from JSON: %w", err)
	}
	return CompressibilityFactory{}.FromDto(*unitDto)
}


// FromInversePascals creates a new Compressibility instance from a value in InversePascals.
func (uf CompressibilityFactory) FromInversePascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInversePascal)
}

// FromInverseKilopascals creates a new Compressibility instance from a value in InverseKilopascals.
func (uf CompressibilityFactory) FromInverseKilopascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseKilopascal)
}

// FromInverseMegapascals creates a new Compressibility instance from a value in InverseMegapascals.
func (uf CompressibilityFactory) FromInverseMegapascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseMegapascal)
}

// FromInverseAtmospheres creates a new Compressibility instance from a value in InverseAtmospheres.
func (uf CompressibilityFactory) FromInverseAtmospheres(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseAtmosphere)
}

// FromInverseMillibars creates a new Compressibility instance from a value in InverseMillibars.
func (uf CompressibilityFactory) FromInverseMillibars(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseMillibar)
}

// FromInverseBars creates a new Compressibility instance from a value in InverseBars.
func (uf CompressibilityFactory) FromInverseBars(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseBar)
}

// FromInversePoundsForcePerSquareInch creates a new Compressibility instance from a value in InversePoundsForcePerSquareInch.
func (uf CompressibilityFactory) FromInversePoundsForcePerSquareInch(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInversePoundForcePerSquareInch)
}


// newCompressibility creates a new Compressibility.
func newCompressibility(value float64, fromUnit CompressibilityUnits) (*Compressibility, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Compressibility{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Compressibility in InversePascal unit (the base/default quantity).
func (a *Compressibility) BaseValue() float64 {
	return a.value
}


// InversePascals returns the Compressibility value in InversePascals.
//
// 
func (a *Compressibility) InversePascals() float64 {
	if a.inverse_pascalsLazy != nil {
		return *a.inverse_pascalsLazy
	}
	inverse_pascals := a.convertFromBase(CompressibilityInversePascal)
	a.inverse_pascalsLazy = &inverse_pascals
	return inverse_pascals
}

// InverseKilopascals returns the Compressibility value in InverseKilopascals.
//
// 
func (a *Compressibility) InverseKilopascals() float64 {
	if a.inverse_kilopascalsLazy != nil {
		return *a.inverse_kilopascalsLazy
	}
	inverse_kilopascals := a.convertFromBase(CompressibilityInverseKilopascal)
	a.inverse_kilopascalsLazy = &inverse_kilopascals
	return inverse_kilopascals
}

// InverseMegapascals returns the Compressibility value in InverseMegapascals.
//
// 
func (a *Compressibility) InverseMegapascals() float64 {
	if a.inverse_megapascalsLazy != nil {
		return *a.inverse_megapascalsLazy
	}
	inverse_megapascals := a.convertFromBase(CompressibilityInverseMegapascal)
	a.inverse_megapascalsLazy = &inverse_megapascals
	return inverse_megapascals
}

// InverseAtmospheres returns the Compressibility value in InverseAtmospheres.
//
// 
func (a *Compressibility) InverseAtmospheres() float64 {
	if a.inverse_atmospheresLazy != nil {
		return *a.inverse_atmospheresLazy
	}
	inverse_atmospheres := a.convertFromBase(CompressibilityInverseAtmosphere)
	a.inverse_atmospheresLazy = &inverse_atmospheres
	return inverse_atmospheres
}

// InverseMillibars returns the Compressibility value in InverseMillibars.
//
// 
func (a *Compressibility) InverseMillibars() float64 {
	if a.inverse_millibarsLazy != nil {
		return *a.inverse_millibarsLazy
	}
	inverse_millibars := a.convertFromBase(CompressibilityInverseMillibar)
	a.inverse_millibarsLazy = &inverse_millibars
	return inverse_millibars
}

// InverseBars returns the Compressibility value in InverseBars.
//
// 
func (a *Compressibility) InverseBars() float64 {
	if a.inverse_barsLazy != nil {
		return *a.inverse_barsLazy
	}
	inverse_bars := a.convertFromBase(CompressibilityInverseBar)
	a.inverse_barsLazy = &inverse_bars
	return inverse_bars
}

// InversePoundsForcePerSquareInch returns the Compressibility value in InversePoundsForcePerSquareInch.
//
// 
func (a *Compressibility) InversePoundsForcePerSquareInch() float64 {
	if a.inverse_pounds_force_per_square_inchLazy != nil {
		return *a.inverse_pounds_force_per_square_inchLazy
	}
	inverse_pounds_force_per_square_inch := a.convertFromBase(CompressibilityInversePoundForcePerSquareInch)
	a.inverse_pounds_force_per_square_inchLazy = &inverse_pounds_force_per_square_inch
	return inverse_pounds_force_per_square_inch
}


// ToDto creates a CompressibilityDto representation from the Compressibility instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InversePascal by default.
func (a *Compressibility) ToDto(holdInUnit *CompressibilityUnits) CompressibilityDto {
	if holdInUnit == nil {
		defaultUnit := CompressibilityInversePascal // Default value
		holdInUnit = &defaultUnit
	}

	return CompressibilityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Compressibility instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InversePascal by default.
func (a *Compressibility) ToDtoJSON(holdInUnit *CompressibilityUnits) ([]byte, error) {
	// Convert to CompressibilityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Compressibility to a specific unit value.
// The function uses the provided unit type (CompressibilityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Compressibility) Convert(toUnit CompressibilityUnits) float64 {
	switch toUnit { 
    case CompressibilityInversePascal:
		return a.InversePascals()
    case CompressibilityInverseKilopascal:
		return a.InverseKilopascals()
    case CompressibilityInverseMegapascal:
		return a.InverseMegapascals()
    case CompressibilityInverseAtmosphere:
		return a.InverseAtmospheres()
    case CompressibilityInverseMillibar:
		return a.InverseMillibars()
    case CompressibilityInverseBar:
		return a.InverseBars()
    case CompressibilityInversePoundForcePerSquareInch:
		return a.InversePoundsForcePerSquareInch()
	default:
		return math.NaN()
	}
}

func (a *Compressibility) convertFromBase(toUnit CompressibilityUnits) float64 {
    value := a.value
	switch toUnit { 
	case CompressibilityInversePascal:
		return (value) 
	case CompressibilityInverseKilopascal:
		return (value / 1e3) 
	case CompressibilityInverseMegapascal:
		return (value / 1e6) 
	case CompressibilityInverseAtmosphere:
		return (value / 101325) 
	case CompressibilityInverseMillibar:
		return (value / 100) 
	case CompressibilityInverseBar:
		return (value / 1e5) 
	case CompressibilityInversePoundForcePerSquareInch:
		return (value / 6.894757293168361e3) 
	default:
		return math.NaN()
	}
}

func (a *Compressibility) convertToBase(value float64, fromUnit CompressibilityUnits) float64 {
	switch fromUnit { 
	case CompressibilityInversePascal:
		return (value) 
	case CompressibilityInverseKilopascal:
		return (value * 1e3) 
	case CompressibilityInverseMegapascal:
		return (value * 1e6) 
	case CompressibilityInverseAtmosphere:
		return (value * 101325) 
	case CompressibilityInverseMillibar:
		return (value * 100) 
	case CompressibilityInverseBar:
		return (value * 1e5) 
	case CompressibilityInversePoundForcePerSquareInch:
		return (value * 6.894757293168361e3) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Compressibility in the default unit (InversePascal),
// formatted to two decimal places.
func (a Compressibility) String() string {
	return a.ToString(CompressibilityInversePascal, 2)
}

// ToString formats the Compressibility value as a string with the specified unit and fractional digits.
// It converts the Compressibility to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Compressibility value will be converted (e.g., InversePascal))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Compressibility with the unit abbreviation.
func (a *Compressibility) ToString(unit CompressibilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetCompressibilityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetCompressibilityAbbreviation(unit))
}

// Equals checks if the given Compressibility is equal to the current Compressibility.
//
// Parameters:
//    other: The Compressibility to compare against.
//
// Returns:
//    bool: Returns true if both Compressibility are equal, false otherwise.
func (a *Compressibility) Equals(other *Compressibility) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Compressibility with another Compressibility.
// It returns -1 if the current Compressibility is less than the other Compressibility, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Compressibility to compare against.
//
// Returns:
//    int: -1 if the current Compressibility is less, 1 if greater, and 0 if equal.
func (a *Compressibility) CompareTo(other *Compressibility) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Compressibility to the current Compressibility and returns the result.
// The result is a new Compressibility instance with the sum of the values.
//
// Parameters:
//    other: The Compressibility to add to the current Compressibility.
//
// Returns:
//    *Compressibility: A new Compressibility instance representing the sum of both Compressibility.
func (a *Compressibility) Add(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Compressibility from the current Compressibility and returns the result.
// The result is a new Compressibility instance with the difference of the values.
//
// Parameters:
//    other: The Compressibility to subtract from the current Compressibility.
//
// Returns:
//    *Compressibility: A new Compressibility instance representing the difference of both Compressibility.
func (a *Compressibility) Subtract(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Compressibility by the given Compressibility and returns the result.
// The result is a new Compressibility instance with the product of the values.
//
// Parameters:
//    other: The Compressibility to multiply with the current Compressibility.
//
// Returns:
//    *Compressibility: A new Compressibility instance representing the product of both Compressibility.
func (a *Compressibility) Multiply(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value * other.BaseValue()}
}

// Divide divides the current Compressibility by the given Compressibility and returns the result.
// The result is a new Compressibility instance with the quotient of the values.
//
// Parameters:
//    other: The Compressibility to divide the current Compressibility by.
//
// Returns:
//    *Compressibility: A new Compressibility instance representing the quotient of both Compressibility.
func (a *Compressibility) Divide(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value / other.BaseValue()}
}

// GetCompressibilityAbbreviation gets the unit abbreviation.
func GetCompressibilityAbbreviation(unit CompressibilityUnits) string {
	switch unit { 
	case CompressibilityInversePascal:
		return "Pa⁻¹" 
	case CompressibilityInverseKilopascal:
		return "kPa⁻¹" 
	case CompressibilityInverseMegapascal:
		return "MPa⁻¹" 
	case CompressibilityInverseAtmosphere:
		return "atm⁻¹" 
	case CompressibilityInverseMillibar:
		return "mbar⁻¹" 
	case CompressibilityInverseBar:
		return "bar⁻¹" 
	case CompressibilityInversePoundForcePerSquareInch:
		return "psi⁻¹" 
	default:
		return ""
	}
}