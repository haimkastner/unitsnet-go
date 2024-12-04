// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MagneticFluxUnits defines various units of MagneticFlux.
type MagneticFluxUnits string

const (
    
        // 
        MagneticFluxWeber MagneticFluxUnits = "Weber"
)

// MagneticFluxDto represents a MagneticFlux measurement with a numerical value and its corresponding unit.
type MagneticFluxDto struct {
    // Value is the numerical representation of the MagneticFlux.
	Value float64
    // Unit specifies the unit of measurement for the MagneticFlux, as defined in the MagneticFluxUnits enumeration.
	Unit  MagneticFluxUnits
}

// MagneticFluxDtoFactory groups methods for creating and serializing MagneticFluxDto objects.
type MagneticFluxDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MagneticFluxDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MagneticFluxDtoFactory) FromJSON(data []byte) (*MagneticFluxDto, error) {
	a := MagneticFluxDto{}

    // Parse JSON into MagneticFluxDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MagneticFluxDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MagneticFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// MagneticFlux represents a measurement in a of MagneticFlux.
//
// In physics, specifically electromagnetism, the magnetic flux through a surface is the surface integral of the normal component of the magnetic field B passing through that surface.
type MagneticFlux struct {
	// value is the base measurement stored internally.
	value       float64
    
    webersLazy *float64 
}

// MagneticFluxFactory groups methods for creating MagneticFlux instances.
type MagneticFluxFactory struct{}

// CreateMagneticFlux creates a new MagneticFlux instance from the given value and unit.
func (uf MagneticFluxFactory) CreateMagneticFlux(value float64, unit MagneticFluxUnits) (*MagneticFlux, error) {
	return newMagneticFlux(value, unit)
}

// FromDto converts a MagneticFluxDto to a MagneticFlux instance.
func (uf MagneticFluxFactory) FromDto(dto MagneticFluxDto) (*MagneticFlux, error) {
	return newMagneticFlux(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MagneticFlux instance.
func (uf MagneticFluxFactory) FromDtoJSON(data []byte) (*MagneticFlux, error) {
	unitDto, err := MagneticFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MagneticFluxDto from JSON: %w", err)
	}
	return MagneticFluxFactory{}.FromDto(*unitDto)
}


// FromWebers creates a new MagneticFlux instance from a value in Webers.
func (uf MagneticFluxFactory) FromWebers(value float64) (*MagneticFlux, error) {
	return newMagneticFlux(value, MagneticFluxWeber)
}


// newMagneticFlux creates a new MagneticFlux.
func newMagneticFlux(value float64, fromUnit MagneticFluxUnits) (*MagneticFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MagneticFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MagneticFlux in Weber unit (the base/default quantity).
func (a *MagneticFlux) BaseValue() float64 {
	return a.value
}


// Webers returns the MagneticFlux value in Webers.
//
// 
func (a *MagneticFlux) Webers() float64 {
	if a.webersLazy != nil {
		return *a.webersLazy
	}
	webers := a.convertFromBase(MagneticFluxWeber)
	a.webersLazy = &webers
	return webers
}


// ToDto creates a MagneticFluxDto representation from the MagneticFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Weber by default.
func (a *MagneticFlux) ToDto(holdInUnit *MagneticFluxUnits) MagneticFluxDto {
	if holdInUnit == nil {
		defaultUnit := MagneticFluxWeber // Default value
		holdInUnit = &defaultUnit
	}

	return MagneticFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MagneticFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Weber by default.
func (a *MagneticFlux) ToDtoJSON(holdInUnit *MagneticFluxUnits) ([]byte, error) {
	// Convert to MagneticFluxDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MagneticFlux to a specific unit value.
// The function uses the provided unit type (MagneticFluxUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MagneticFlux) Convert(toUnit MagneticFluxUnits) float64 {
	switch toUnit { 
    case MagneticFluxWeber:
		return a.Webers()
	default:
		return math.NaN()
	}
}

func (a *MagneticFlux) convertFromBase(toUnit MagneticFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagneticFluxWeber:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *MagneticFlux) convertToBase(value float64, fromUnit MagneticFluxUnits) float64 {
	switch fromUnit { 
	case MagneticFluxWeber:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MagneticFlux in the default unit (Weber),
// formatted to two decimal places.
func (a MagneticFlux) String() string {
	return a.ToString(MagneticFluxWeber, 2)
}

// ToString formats the MagneticFlux value as a string with the specified unit and fractional digits.
// It converts the MagneticFlux to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MagneticFlux value will be converted (e.g., Weber))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MagneticFlux with the unit abbreviation.
func (a *MagneticFlux) ToString(unit MagneticFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MagneticFlux) getUnitAbbreviation(unit MagneticFluxUnits) string {
	switch unit { 
	case MagneticFluxWeber:
		return "Wb" 
	default:
		return ""
	}
}

// Equals checks if the given MagneticFlux is equal to the current MagneticFlux.
//
// Parameters:
//    other: The MagneticFlux to compare against.
//
// Returns:
//    bool: Returns true if both MagneticFlux are equal, false otherwise.
func (a *MagneticFlux) Equals(other *MagneticFlux) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MagneticFlux with another MagneticFlux.
// It returns -1 if the current MagneticFlux is less than the other MagneticFlux, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MagneticFlux to compare against.
//
// Returns:
//    int: -1 if the current MagneticFlux is less, 1 if greater, and 0 if equal.
func (a *MagneticFlux) CompareTo(other *MagneticFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MagneticFlux to the current MagneticFlux and returns the result.
// The result is a new MagneticFlux instance with the sum of the values.
//
// Parameters:
//    other: The MagneticFlux to add to the current MagneticFlux.
//
// Returns:
//    *MagneticFlux: A new MagneticFlux instance representing the sum of both MagneticFlux.
func (a *MagneticFlux) Add(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MagneticFlux from the current MagneticFlux and returns the result.
// The result is a new MagneticFlux instance with the difference of the values.
//
// Parameters:
//    other: The MagneticFlux to subtract from the current MagneticFlux.
//
// Returns:
//    *MagneticFlux: A new MagneticFlux instance representing the difference of both MagneticFlux.
func (a *MagneticFlux) Subtract(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MagneticFlux by the given MagneticFlux and returns the result.
// The result is a new MagneticFlux instance with the product of the values.
//
// Parameters:
//    other: The MagneticFlux to multiply with the current MagneticFlux.
//
// Returns:
//    *MagneticFlux: A new MagneticFlux instance representing the product of both MagneticFlux.
func (a *MagneticFlux) Multiply(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value * other.BaseValue()}
}

// Divide divides the current MagneticFlux by the given MagneticFlux and returns the result.
// The result is a new MagneticFlux instance with the quotient of the values.
//
// Parameters:
//    other: The MagneticFlux to divide the current MagneticFlux by.
//
// Returns:
//    *MagneticFlux: A new MagneticFlux instance representing the quotient of both MagneticFlux.
func (a *MagneticFlux) Divide(other *MagneticFlux) *MagneticFlux {
	return &MagneticFlux{value: a.value / other.BaseValue()}
}