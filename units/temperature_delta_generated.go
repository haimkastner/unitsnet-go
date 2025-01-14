// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureDeltaUnits defines various units of TemperatureDelta.
type TemperatureDeltaUnits string

const (
    
        // 
        TemperatureDeltaKelvin TemperatureDeltaUnits = "Kelvin"
        // 
        TemperatureDeltaDegreeCelsius TemperatureDeltaUnits = "DegreeCelsius"
        // 
        TemperatureDeltaDegreeDelisle TemperatureDeltaUnits = "DegreeDelisle"
        // 
        TemperatureDeltaDegreeFahrenheit TemperatureDeltaUnits = "DegreeFahrenheit"
        // 
        TemperatureDeltaDegreeNewton TemperatureDeltaUnits = "DegreeNewton"
        // 
        TemperatureDeltaDegreeRankine TemperatureDeltaUnits = "DegreeRankine"
        // 
        TemperatureDeltaDegreeReaumur TemperatureDeltaUnits = "DegreeReaumur"
        // 
        TemperatureDeltaDegreeRoemer TemperatureDeltaUnits = "DegreeRoemer"
        // 
        TemperatureDeltaMillidegreeCelsius TemperatureDeltaUnits = "MillidegreeCelsius"
)

// TemperatureDeltaDto represents a TemperatureDelta measurement with a numerical value and its corresponding unit.
type TemperatureDeltaDto struct {
    // Value is the numerical representation of the TemperatureDelta.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the TemperatureDelta, as defined in the TemperatureDeltaUnits enumeration.
	Unit  TemperatureDeltaUnits `json:"unit"`
}

// TemperatureDeltaDtoFactory groups methods for creating and serializing TemperatureDeltaDto objects.
type TemperatureDeltaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TemperatureDeltaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TemperatureDeltaDtoFactory) FromJSON(data []byte) (*TemperatureDeltaDto, error) {
	a := TemperatureDeltaDto{}

    // Parse JSON into TemperatureDeltaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a TemperatureDeltaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TemperatureDeltaDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// TemperatureDelta represents a measurement in a of TemperatureDelta.
//
// Difference between two temperatures. The conversions are different than for Temperature.
type TemperatureDelta struct {
	// value is the base measurement stored internally.
	value       float64
    
    kelvinsLazy *float64 
    degrees_celsiusLazy *float64 
    degrees_delisleLazy *float64 
    degrees_fahrenheitLazy *float64 
    degrees_newtonLazy *float64 
    degrees_rankineLazy *float64 
    degrees_reaumurLazy *float64 
    degrees_roemerLazy *float64 
    millidegrees_celsiusLazy *float64 
}

// TemperatureDeltaFactory groups methods for creating TemperatureDelta instances.
type TemperatureDeltaFactory struct{}

// CreateTemperatureDelta creates a new TemperatureDelta instance from the given value and unit.
func (uf TemperatureDeltaFactory) CreateTemperatureDelta(value float64, unit TemperatureDeltaUnits) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, unit)
}

// FromDto converts a TemperatureDeltaDto to a TemperatureDelta instance.
func (uf TemperatureDeltaFactory) FromDto(dto TemperatureDeltaDto) (*TemperatureDelta, error) {
	return newTemperatureDelta(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a TemperatureDelta instance.
func (uf TemperatureDeltaFactory) FromDtoJSON(data []byte) (*TemperatureDelta, error) {
	unitDto, err := TemperatureDeltaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TemperatureDeltaDto from JSON: %w", err)
	}
	return TemperatureDeltaFactory{}.FromDto(*unitDto)
}


// FromKelvins creates a new TemperatureDelta instance from a value in Kelvins.
func (uf TemperatureDeltaFactory) FromKelvins(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaKelvin)
}

// FromDegreesCelsius creates a new TemperatureDelta instance from a value in DegreesCelsius.
func (uf TemperatureDeltaFactory) FromDegreesCelsius(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeCelsius)
}

// FromDegreesDelisle creates a new TemperatureDelta instance from a value in DegreesDelisle.
func (uf TemperatureDeltaFactory) FromDegreesDelisle(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeDelisle)
}

// FromDegreesFahrenheit creates a new TemperatureDelta instance from a value in DegreesFahrenheit.
func (uf TemperatureDeltaFactory) FromDegreesFahrenheit(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeFahrenheit)
}

// FromDegreesNewton creates a new TemperatureDelta instance from a value in DegreesNewton.
func (uf TemperatureDeltaFactory) FromDegreesNewton(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeNewton)
}

// FromDegreesRankine creates a new TemperatureDelta instance from a value in DegreesRankine.
func (uf TemperatureDeltaFactory) FromDegreesRankine(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeRankine)
}

// FromDegreesReaumur creates a new TemperatureDelta instance from a value in DegreesReaumur.
func (uf TemperatureDeltaFactory) FromDegreesReaumur(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeReaumur)
}

// FromDegreesRoemer creates a new TemperatureDelta instance from a value in DegreesRoemer.
func (uf TemperatureDeltaFactory) FromDegreesRoemer(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeRoemer)
}

// FromMillidegreesCelsius creates a new TemperatureDelta instance from a value in MillidegreesCelsius.
func (uf TemperatureDeltaFactory) FromMillidegreesCelsius(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaMillidegreeCelsius)
}


// newTemperatureDelta creates a new TemperatureDelta.
func newTemperatureDelta(value float64, fromUnit TemperatureDeltaUnits) (*TemperatureDelta, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TemperatureDelta{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureDelta in Kelvin unit (the base/default quantity).
func (a *TemperatureDelta) BaseValue() float64 {
	return a.value
}


// Kelvins returns the TemperatureDelta value in Kelvins.
//
// 
func (a *TemperatureDelta) Kelvins() float64 {
	if a.kelvinsLazy != nil {
		return *a.kelvinsLazy
	}
	kelvins := a.convertFromBase(TemperatureDeltaKelvin)
	a.kelvinsLazy = &kelvins
	return kelvins
}

// DegreesCelsius returns the TemperatureDelta value in DegreesCelsius.
//
// 
func (a *TemperatureDelta) DegreesCelsius() float64 {
	if a.degrees_celsiusLazy != nil {
		return *a.degrees_celsiusLazy
	}
	degrees_celsius := a.convertFromBase(TemperatureDeltaDegreeCelsius)
	a.degrees_celsiusLazy = &degrees_celsius
	return degrees_celsius
}

// DegreesDelisle returns the TemperatureDelta value in DegreesDelisle.
//
// 
func (a *TemperatureDelta) DegreesDelisle() float64 {
	if a.degrees_delisleLazy != nil {
		return *a.degrees_delisleLazy
	}
	degrees_delisle := a.convertFromBase(TemperatureDeltaDegreeDelisle)
	a.degrees_delisleLazy = &degrees_delisle
	return degrees_delisle
}

// DegreesFahrenheit returns the TemperatureDelta value in DegreesFahrenheit.
//
// 
func (a *TemperatureDelta) DegreesFahrenheit() float64 {
	if a.degrees_fahrenheitLazy != nil {
		return *a.degrees_fahrenheitLazy
	}
	degrees_fahrenheit := a.convertFromBase(TemperatureDeltaDegreeFahrenheit)
	a.degrees_fahrenheitLazy = &degrees_fahrenheit
	return degrees_fahrenheit
}

// DegreesNewton returns the TemperatureDelta value in DegreesNewton.
//
// 
func (a *TemperatureDelta) DegreesNewton() float64 {
	if a.degrees_newtonLazy != nil {
		return *a.degrees_newtonLazy
	}
	degrees_newton := a.convertFromBase(TemperatureDeltaDegreeNewton)
	a.degrees_newtonLazy = &degrees_newton
	return degrees_newton
}

// DegreesRankine returns the TemperatureDelta value in DegreesRankine.
//
// 
func (a *TemperatureDelta) DegreesRankine() float64 {
	if a.degrees_rankineLazy != nil {
		return *a.degrees_rankineLazy
	}
	degrees_rankine := a.convertFromBase(TemperatureDeltaDegreeRankine)
	a.degrees_rankineLazy = &degrees_rankine
	return degrees_rankine
}

// DegreesReaumur returns the TemperatureDelta value in DegreesReaumur.
//
// 
func (a *TemperatureDelta) DegreesReaumur() float64 {
	if a.degrees_reaumurLazy != nil {
		return *a.degrees_reaumurLazy
	}
	degrees_reaumur := a.convertFromBase(TemperatureDeltaDegreeReaumur)
	a.degrees_reaumurLazy = &degrees_reaumur
	return degrees_reaumur
}

// DegreesRoemer returns the TemperatureDelta value in DegreesRoemer.
//
// 
func (a *TemperatureDelta) DegreesRoemer() float64 {
	if a.degrees_roemerLazy != nil {
		return *a.degrees_roemerLazy
	}
	degrees_roemer := a.convertFromBase(TemperatureDeltaDegreeRoemer)
	a.degrees_roemerLazy = &degrees_roemer
	return degrees_roemer
}

// MillidegreesCelsius returns the TemperatureDelta value in MillidegreesCelsius.
//
// 
func (a *TemperatureDelta) MillidegreesCelsius() float64 {
	if a.millidegrees_celsiusLazy != nil {
		return *a.millidegrees_celsiusLazy
	}
	millidegrees_celsius := a.convertFromBase(TemperatureDeltaMillidegreeCelsius)
	a.millidegrees_celsiusLazy = &millidegrees_celsius
	return millidegrees_celsius
}


// ToDto creates a TemperatureDeltaDto representation from the TemperatureDelta instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kelvin by default.
func (a *TemperatureDelta) ToDto(holdInUnit *TemperatureDeltaUnits) TemperatureDeltaDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureDeltaKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureDeltaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the TemperatureDelta instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kelvin by default.
func (a *TemperatureDelta) ToDtoJSON(holdInUnit *TemperatureDeltaUnits) ([]byte, error) {
	// Convert to TemperatureDeltaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a TemperatureDelta to a specific unit value.
// The function uses the provided unit type (TemperatureDeltaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *TemperatureDelta) Convert(toUnit TemperatureDeltaUnits) float64 {
	switch toUnit { 
    case TemperatureDeltaKelvin:
		return a.Kelvins()
    case TemperatureDeltaDegreeCelsius:
		return a.DegreesCelsius()
    case TemperatureDeltaDegreeDelisle:
		return a.DegreesDelisle()
    case TemperatureDeltaDegreeFahrenheit:
		return a.DegreesFahrenheit()
    case TemperatureDeltaDegreeNewton:
		return a.DegreesNewton()
    case TemperatureDeltaDegreeRankine:
		return a.DegreesRankine()
    case TemperatureDeltaDegreeReaumur:
		return a.DegreesReaumur()
    case TemperatureDeltaDegreeRoemer:
		return a.DegreesRoemer()
    case TemperatureDeltaMillidegreeCelsius:
		return a.MillidegreesCelsius()
	default:
		return math.NaN()
	}
}

func (a *TemperatureDelta) convertFromBase(toUnit TemperatureDeltaUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureDeltaKelvin:
		return (value) 
	case TemperatureDeltaDegreeCelsius:
		return (value) 
	case TemperatureDeltaDegreeDelisle:
		return (value * -3 / 2) 
	case TemperatureDeltaDegreeFahrenheit:
		return (value * 9 / 5) 
	case TemperatureDeltaDegreeNewton:
		return (value * 33 / 100) 
	case TemperatureDeltaDegreeRankine:
		return (value * 9 / 5) 
	case TemperatureDeltaDegreeReaumur:
		return (value * 4 / 5) 
	case TemperatureDeltaDegreeRoemer:
		return (value * 21 / 40) 
	case TemperatureDeltaMillidegreeCelsius:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *TemperatureDelta) convertToBase(value float64, fromUnit TemperatureDeltaUnits) float64 {
	switch fromUnit { 
	case TemperatureDeltaKelvin:
		return (value) 
	case TemperatureDeltaDegreeCelsius:
		return (value) 
	case TemperatureDeltaDegreeDelisle:
		return (value * -2 / 3) 
	case TemperatureDeltaDegreeFahrenheit:
		return (value * 5 / 9) 
	case TemperatureDeltaDegreeNewton:
		return (value * 100 / 33) 
	case TemperatureDeltaDegreeRankine:
		return (value * 5 / 9) 
	case TemperatureDeltaDegreeReaumur:
		return (value * 5 / 4) 
	case TemperatureDeltaDegreeRoemer:
		return (value * 40 / 21) 
	case TemperatureDeltaMillidegreeCelsius:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the TemperatureDelta in the default unit (Kelvin),
// formatted to two decimal places.
func (a TemperatureDelta) String() string {
	return a.ToString(TemperatureDeltaKelvin, 2)
}

// ToString formats the TemperatureDelta value as a string with the specified unit and fractional digits.
// It converts the TemperatureDelta to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the TemperatureDelta value will be converted (e.g., Kelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the TemperatureDelta with the unit abbreviation.
func (a *TemperatureDelta) ToString(unit TemperatureDeltaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TemperatureDelta) getUnitAbbreviation(unit TemperatureDeltaUnits) string {
	switch unit { 
	case TemperatureDeltaKelvin:
		return "∆K" 
	case TemperatureDeltaDegreeCelsius:
		return "∆°C" 
	case TemperatureDeltaDegreeDelisle:
		return "∆°De" 
	case TemperatureDeltaDegreeFahrenheit:
		return "∆°F" 
	case TemperatureDeltaDegreeNewton:
		return "∆°N" 
	case TemperatureDeltaDegreeRankine:
		return "∆°R" 
	case TemperatureDeltaDegreeReaumur:
		return "∆°Ré" 
	case TemperatureDeltaDegreeRoemer:
		return "∆°Rø" 
	case TemperatureDeltaMillidegreeCelsius:
		return "m∆°C" 
	default:
		return ""
	}
}

// Equals checks if the given TemperatureDelta is equal to the current TemperatureDelta.
//
// Parameters:
//    other: The TemperatureDelta to compare against.
//
// Returns:
//    bool: Returns true if both TemperatureDelta are equal, false otherwise.
func (a *TemperatureDelta) Equals(other *TemperatureDelta) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current TemperatureDelta with another TemperatureDelta.
// It returns -1 if the current TemperatureDelta is less than the other TemperatureDelta, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The TemperatureDelta to compare against.
//
// Returns:
//    int: -1 if the current TemperatureDelta is less, 1 if greater, and 0 if equal.
func (a *TemperatureDelta) CompareTo(other *TemperatureDelta) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given TemperatureDelta to the current TemperatureDelta and returns the result.
// The result is a new TemperatureDelta instance with the sum of the values.
//
// Parameters:
//    other: The TemperatureDelta to add to the current TemperatureDelta.
//
// Returns:
//    *TemperatureDelta: A new TemperatureDelta instance representing the sum of both TemperatureDelta.
func (a *TemperatureDelta) Add(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given TemperatureDelta from the current TemperatureDelta and returns the result.
// The result is a new TemperatureDelta instance with the difference of the values.
//
// Parameters:
//    other: The TemperatureDelta to subtract from the current TemperatureDelta.
//
// Returns:
//    *TemperatureDelta: A new TemperatureDelta instance representing the difference of both TemperatureDelta.
func (a *TemperatureDelta) Subtract(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current TemperatureDelta by the given TemperatureDelta and returns the result.
// The result is a new TemperatureDelta instance with the product of the values.
//
// Parameters:
//    other: The TemperatureDelta to multiply with the current TemperatureDelta.
//
// Returns:
//    *TemperatureDelta: A new TemperatureDelta instance representing the product of both TemperatureDelta.
func (a *TemperatureDelta) Multiply(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value * other.BaseValue()}
}

// Divide divides the current TemperatureDelta by the given TemperatureDelta and returns the result.
// The result is a new TemperatureDelta instance with the quotient of the values.
//
// Parameters:
//    other: The TemperatureDelta to divide the current TemperatureDelta by.
//
// Returns:
//    *TemperatureDelta: A new TemperatureDelta instance representing the quotient of both TemperatureDelta.
func (a *TemperatureDelta) Divide(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value / other.BaseValue()}
}