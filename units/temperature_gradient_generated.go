// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureGradientUnits defines various units of TemperatureGradient.
type TemperatureGradientUnits string

const (
    
        // 
        TemperatureGradientKelvinPerMeter TemperatureGradientUnits = "KelvinPerMeter"
        // 
        TemperatureGradientDegreeCelsiusPerMeter TemperatureGradientUnits = "DegreeCelsiusPerMeter"
        // 
        TemperatureGradientDegreeFahrenheitPerFoot TemperatureGradientUnits = "DegreeFahrenheitPerFoot"
        // 
        TemperatureGradientDegreeCelsiusPerKilometer TemperatureGradientUnits = "DegreeCelsiusPerKilometer"
)

// TemperatureGradientDto represents a TemperatureGradient measurement with a numerical value and its corresponding unit.
type TemperatureGradientDto struct {
    // Value is the numerical representation of the TemperatureGradient.
	Value float64
    // Unit specifies the unit of measurement for the TemperatureGradient, as defined in the TemperatureGradientUnits enumeration.
	Unit  TemperatureGradientUnits
}

// TemperatureGradientDtoFactory groups methods for creating and serializing TemperatureGradientDto objects.
type TemperatureGradientDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TemperatureGradientDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TemperatureGradientDtoFactory) FromJSON(data []byte) (*TemperatureGradientDto, error) {
	a := TemperatureGradientDto{}

    // Parse JSON into TemperatureGradientDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a TemperatureGradientDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TemperatureGradientDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// TemperatureGradient represents a measurement in a of TemperatureGradient.
//
// None
type TemperatureGradient struct {
	// value is the base measurement stored internally.
	value       float64
    
    kelvins_per_meterLazy *float64 
    degrees_celcius_per_meterLazy *float64 
    degrees_fahrenheit_per_footLazy *float64 
    degrees_celcius_per_kilometerLazy *float64 
}

// TemperatureGradientFactory groups methods for creating TemperatureGradient instances.
type TemperatureGradientFactory struct{}

// CreateTemperatureGradient creates a new TemperatureGradient instance from the given value and unit.
func (uf TemperatureGradientFactory) CreateTemperatureGradient(value float64, unit TemperatureGradientUnits) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, unit)
}

// FromDto converts a TemperatureGradientDto to a TemperatureGradient instance.
func (uf TemperatureGradientFactory) FromDto(dto TemperatureGradientDto) (*TemperatureGradient, error) {
	return newTemperatureGradient(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a TemperatureGradient instance.
func (uf TemperatureGradientFactory) FromDtoJSON(data []byte) (*TemperatureGradient, error) {
	unitDto, err := TemperatureGradientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TemperatureGradientDto from JSON: %w", err)
	}
	return TemperatureGradientFactory{}.FromDto(*unitDto)
}


// FromKelvinsPerMeter creates a new TemperatureGradient instance from a value in KelvinsPerMeter.
func (uf TemperatureGradientFactory) FromKelvinsPerMeter(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientKelvinPerMeter)
}

// FromDegreesCelciusPerMeter creates a new TemperatureGradient instance from a value in DegreesCelciusPerMeter.
func (uf TemperatureGradientFactory) FromDegreesCelciusPerMeter(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeCelsiusPerMeter)
}

// FromDegreesFahrenheitPerFoot creates a new TemperatureGradient instance from a value in DegreesFahrenheitPerFoot.
func (uf TemperatureGradientFactory) FromDegreesFahrenheitPerFoot(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeFahrenheitPerFoot)
}

// FromDegreesCelciusPerKilometer creates a new TemperatureGradient instance from a value in DegreesCelciusPerKilometer.
func (uf TemperatureGradientFactory) FromDegreesCelciusPerKilometer(value float64) (*TemperatureGradient, error) {
	return newTemperatureGradient(value, TemperatureGradientDegreeCelsiusPerKilometer)
}


// newTemperatureGradient creates a new TemperatureGradient.
func newTemperatureGradient(value float64, fromUnit TemperatureGradientUnits) (*TemperatureGradient, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TemperatureGradient{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureGradient in KelvinPerMeter unit (the base/default quantity).
func (a *TemperatureGradient) BaseValue() float64 {
	return a.value
}


// KelvinsPerMeter returns the TemperatureGradient value in KelvinsPerMeter.
//
// 
func (a *TemperatureGradient) KelvinsPerMeter() float64 {
	if a.kelvins_per_meterLazy != nil {
		return *a.kelvins_per_meterLazy
	}
	kelvins_per_meter := a.convertFromBase(TemperatureGradientKelvinPerMeter)
	a.kelvins_per_meterLazy = &kelvins_per_meter
	return kelvins_per_meter
}

// DegreesCelciusPerMeter returns the TemperatureGradient value in DegreesCelciusPerMeter.
//
// 
func (a *TemperatureGradient) DegreesCelciusPerMeter() float64 {
	if a.degrees_celcius_per_meterLazy != nil {
		return *a.degrees_celcius_per_meterLazy
	}
	degrees_celcius_per_meter := a.convertFromBase(TemperatureGradientDegreeCelsiusPerMeter)
	a.degrees_celcius_per_meterLazy = &degrees_celcius_per_meter
	return degrees_celcius_per_meter
}

// DegreesFahrenheitPerFoot returns the TemperatureGradient value in DegreesFahrenheitPerFoot.
//
// 
func (a *TemperatureGradient) DegreesFahrenheitPerFoot() float64 {
	if a.degrees_fahrenheit_per_footLazy != nil {
		return *a.degrees_fahrenheit_per_footLazy
	}
	degrees_fahrenheit_per_foot := a.convertFromBase(TemperatureGradientDegreeFahrenheitPerFoot)
	a.degrees_fahrenheit_per_footLazy = &degrees_fahrenheit_per_foot
	return degrees_fahrenheit_per_foot
}

// DegreesCelciusPerKilometer returns the TemperatureGradient value in DegreesCelciusPerKilometer.
//
// 
func (a *TemperatureGradient) DegreesCelciusPerKilometer() float64 {
	if a.degrees_celcius_per_kilometerLazy != nil {
		return *a.degrees_celcius_per_kilometerLazy
	}
	degrees_celcius_per_kilometer := a.convertFromBase(TemperatureGradientDegreeCelsiusPerKilometer)
	a.degrees_celcius_per_kilometerLazy = &degrees_celcius_per_kilometer
	return degrees_celcius_per_kilometer
}


// ToDto creates a TemperatureGradientDto representation from the TemperatureGradient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KelvinPerMeter by default.
func (a *TemperatureGradient) ToDto(holdInUnit *TemperatureGradientUnits) TemperatureGradientDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureGradientKelvinPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureGradientDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the TemperatureGradient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KelvinPerMeter by default.
func (a *TemperatureGradient) ToDtoJSON(holdInUnit *TemperatureGradientUnits) ([]byte, error) {
	// Convert to TemperatureGradientDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a TemperatureGradient to a specific unit value.
// The function uses the provided unit type (TemperatureGradientUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *TemperatureGradient) Convert(toUnit TemperatureGradientUnits) float64 {
	switch toUnit { 
    case TemperatureGradientKelvinPerMeter:
		return a.KelvinsPerMeter()
    case TemperatureGradientDegreeCelsiusPerMeter:
		return a.DegreesCelciusPerMeter()
    case TemperatureGradientDegreeFahrenheitPerFoot:
		return a.DegreesFahrenheitPerFoot()
    case TemperatureGradientDegreeCelsiusPerKilometer:
		return a.DegreesCelciusPerKilometer()
	default:
		return math.NaN()
	}
}

func (a *TemperatureGradient) convertFromBase(toUnit TemperatureGradientUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureGradientKelvinPerMeter:
		return (value) 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return (value) 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return ((value * 0.3048) * 9 / 5) 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return (value * 1e3) 
	default:
		return math.NaN()
	}
}

func (a *TemperatureGradient) convertToBase(value float64, fromUnit TemperatureGradientUnits) float64 {
	switch fromUnit { 
	case TemperatureGradientKelvinPerMeter:
		return (value) 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return (value) 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return ((value / 0.3048) * 5 / 9) 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return (value / 1e3) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the TemperatureGradient in the default unit (KelvinPerMeter),
// formatted to two decimal places.
func (a TemperatureGradient) String() string {
	return a.ToString(TemperatureGradientKelvinPerMeter, 2)
}

// ToString formats the TemperatureGradient value as a string with the specified unit and fractional digits.
// It converts the TemperatureGradient to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the TemperatureGradient value will be converted (e.g., KelvinPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the TemperatureGradient with the unit abbreviation.
func (a *TemperatureGradient) ToString(unit TemperatureGradientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TemperatureGradient) getUnitAbbreviation(unit TemperatureGradientUnits) string {
	switch unit { 
	case TemperatureGradientKelvinPerMeter:
		return "∆°K/m" 
	case TemperatureGradientDegreeCelsiusPerMeter:
		return "∆°C/m" 
	case TemperatureGradientDegreeFahrenheitPerFoot:
		return "∆°F/ft" 
	case TemperatureGradientDegreeCelsiusPerKilometer:
		return "∆°C/km" 
	default:
		return ""
	}
}

// Equals checks if the given TemperatureGradient is equal to the current TemperatureGradient.
//
// Parameters:
//    other: The TemperatureGradient to compare against.
//
// Returns:
//    bool: Returns true if both TemperatureGradient are equal, false otherwise.
func (a *TemperatureGradient) Equals(other *TemperatureGradient) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current TemperatureGradient with another TemperatureGradient.
// It returns -1 if the current TemperatureGradient is less than the other TemperatureGradient, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The TemperatureGradient to compare against.
//
// Returns:
//    int: -1 if the current TemperatureGradient is less, 1 if greater, and 0 if equal.
func (a *TemperatureGradient) CompareTo(other *TemperatureGradient) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given TemperatureGradient to the current TemperatureGradient and returns the result.
// The result is a new TemperatureGradient instance with the sum of the values.
//
// Parameters:
//    other: The TemperatureGradient to add to the current TemperatureGradient.
//
// Returns:
//    *TemperatureGradient: A new TemperatureGradient instance representing the sum of both TemperatureGradient.
func (a *TemperatureGradient) Add(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given TemperatureGradient from the current TemperatureGradient and returns the result.
// The result is a new TemperatureGradient instance with the difference of the values.
//
// Parameters:
//    other: The TemperatureGradient to subtract from the current TemperatureGradient.
//
// Returns:
//    *TemperatureGradient: A new TemperatureGradient instance representing the difference of both TemperatureGradient.
func (a *TemperatureGradient) Subtract(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current TemperatureGradient by the given TemperatureGradient and returns the result.
// The result is a new TemperatureGradient instance with the product of the values.
//
// Parameters:
//    other: The TemperatureGradient to multiply with the current TemperatureGradient.
//
// Returns:
//    *TemperatureGradient: A new TemperatureGradient instance representing the product of both TemperatureGradient.
func (a *TemperatureGradient) Multiply(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value * other.BaseValue()}
}

// Divide divides the current TemperatureGradient by the given TemperatureGradient and returns the result.
// The result is a new TemperatureGradient instance with the quotient of the values.
//
// Parameters:
//    other: The TemperatureGradient to divide the current TemperatureGradient by.
//
// Returns:
//    *TemperatureGradient: A new TemperatureGradient instance representing the quotient of both TemperatureGradient.
func (a *TemperatureGradient) Divide(other *TemperatureGradient) *TemperatureGradient {
	return &TemperatureGradient{value: a.value / other.BaseValue()}
}