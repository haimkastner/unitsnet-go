// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// CapacitanceUnits defines various units of Capacitance.
type CapacitanceUnits string

const (
    
        // 
        CapacitanceFarad CapacitanceUnits = "Farad"
        // 
        CapacitancePicofarad CapacitanceUnits = "Picofarad"
        // 
        CapacitanceNanofarad CapacitanceUnits = "Nanofarad"
        // 
        CapacitanceMicrofarad CapacitanceUnits = "Microfarad"
        // 
        CapacitanceMillifarad CapacitanceUnits = "Millifarad"
        // 
        CapacitanceKilofarad CapacitanceUnits = "Kilofarad"
        // 
        CapacitanceMegafarad CapacitanceUnits = "Megafarad"
)

// CapacitanceDto represents a Capacitance measurement with a numerical value and its corresponding unit.
type CapacitanceDto struct {
    // Value is the numerical representation of the Capacitance.
	Value float64
    // Unit specifies the unit of measurement for the Capacitance, as defined in the CapacitanceUnits enumeration.
	Unit  CapacitanceUnits
}

// CapacitanceDtoFactory groups methods for creating and serializing CapacitanceDto objects.
type CapacitanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a CapacitanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf CapacitanceDtoFactory) FromJSON(data []byte) (*CapacitanceDto, error) {
	a := CapacitanceDto{}

    // Parse JSON into CapacitanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a CapacitanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a CapacitanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Capacitance represents a measurement in a of Capacitance.
//
// Capacitance is the ability of a body to store an electric charge.
type Capacitance struct {
	// value is the base measurement stored internally.
	value       float64
    
    faradsLazy *float64 
    picofaradsLazy *float64 
    nanofaradsLazy *float64 
    microfaradsLazy *float64 
    millifaradsLazy *float64 
    kilofaradsLazy *float64 
    megafaradsLazy *float64 
}

// CapacitanceFactory groups methods for creating Capacitance instances.
type CapacitanceFactory struct{}

// CreateCapacitance creates a new Capacitance instance from the given value and unit.
func (uf CapacitanceFactory) CreateCapacitance(value float64, unit CapacitanceUnits) (*Capacitance, error) {
	return newCapacitance(value, unit)
}

// FromDto converts a CapacitanceDto to a Capacitance instance.
func (uf CapacitanceFactory) FromDto(dto CapacitanceDto) (*Capacitance, error) {
	return newCapacitance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Capacitance instance.
func (uf CapacitanceFactory) FromDtoJSON(data []byte) (*Capacitance, error) {
	unitDto, err := CapacitanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CapacitanceDto from JSON: %w", err)
	}
	return CapacitanceFactory{}.FromDto(*unitDto)
}


// FromFarads creates a new Capacitance instance from a value in Farads.
func (uf CapacitanceFactory) FromFarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceFarad)
}

// FromPicofarads creates a new Capacitance instance from a value in Picofarads.
func (uf CapacitanceFactory) FromPicofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitancePicofarad)
}

// FromNanofarads creates a new Capacitance instance from a value in Nanofarads.
func (uf CapacitanceFactory) FromNanofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceNanofarad)
}

// FromMicrofarads creates a new Capacitance instance from a value in Microfarads.
func (uf CapacitanceFactory) FromMicrofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMicrofarad)
}

// FromMillifarads creates a new Capacitance instance from a value in Millifarads.
func (uf CapacitanceFactory) FromMillifarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMillifarad)
}

// FromKilofarads creates a new Capacitance instance from a value in Kilofarads.
func (uf CapacitanceFactory) FromKilofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceKilofarad)
}

// FromMegafarads creates a new Capacitance instance from a value in Megafarads.
func (uf CapacitanceFactory) FromMegafarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMegafarad)
}


// newCapacitance creates a new Capacitance.
func newCapacitance(value float64, fromUnit CapacitanceUnits) (*Capacitance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Capacitance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Capacitance in Farad unit (the base/default quantity).
func (a *Capacitance) BaseValue() float64 {
	return a.value
}


// Farads returns the Capacitance value in Farads.
//
// 
func (a *Capacitance) Farads() float64 {
	if a.faradsLazy != nil {
		return *a.faradsLazy
	}
	farads := a.convertFromBase(CapacitanceFarad)
	a.faradsLazy = &farads
	return farads
}

// Picofarads returns the Capacitance value in Picofarads.
//
// 
func (a *Capacitance) Picofarads() float64 {
	if a.picofaradsLazy != nil {
		return *a.picofaradsLazy
	}
	picofarads := a.convertFromBase(CapacitancePicofarad)
	a.picofaradsLazy = &picofarads
	return picofarads
}

// Nanofarads returns the Capacitance value in Nanofarads.
//
// 
func (a *Capacitance) Nanofarads() float64 {
	if a.nanofaradsLazy != nil {
		return *a.nanofaradsLazy
	}
	nanofarads := a.convertFromBase(CapacitanceNanofarad)
	a.nanofaradsLazy = &nanofarads
	return nanofarads
}

// Microfarads returns the Capacitance value in Microfarads.
//
// 
func (a *Capacitance) Microfarads() float64 {
	if a.microfaradsLazy != nil {
		return *a.microfaradsLazy
	}
	microfarads := a.convertFromBase(CapacitanceMicrofarad)
	a.microfaradsLazy = &microfarads
	return microfarads
}

// Millifarads returns the Capacitance value in Millifarads.
//
// 
func (a *Capacitance) Millifarads() float64 {
	if a.millifaradsLazy != nil {
		return *a.millifaradsLazy
	}
	millifarads := a.convertFromBase(CapacitanceMillifarad)
	a.millifaradsLazy = &millifarads
	return millifarads
}

// Kilofarads returns the Capacitance value in Kilofarads.
//
// 
func (a *Capacitance) Kilofarads() float64 {
	if a.kilofaradsLazy != nil {
		return *a.kilofaradsLazy
	}
	kilofarads := a.convertFromBase(CapacitanceKilofarad)
	a.kilofaradsLazy = &kilofarads
	return kilofarads
}

// Megafarads returns the Capacitance value in Megafarads.
//
// 
func (a *Capacitance) Megafarads() float64 {
	if a.megafaradsLazy != nil {
		return *a.megafaradsLazy
	}
	megafarads := a.convertFromBase(CapacitanceMegafarad)
	a.megafaradsLazy = &megafarads
	return megafarads
}


// ToDto creates a CapacitanceDto representation from the Capacitance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Farad by default.
func (a *Capacitance) ToDto(holdInUnit *CapacitanceUnits) CapacitanceDto {
	if holdInUnit == nil {
		defaultUnit := CapacitanceFarad // Default value
		holdInUnit = &defaultUnit
	}

	return CapacitanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Capacitance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Farad by default.
func (a *Capacitance) ToDtoJSON(holdInUnit *CapacitanceUnits) ([]byte, error) {
	// Convert to CapacitanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Capacitance to a specific unit value.
// The function uses the provided unit type (CapacitanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Capacitance) Convert(toUnit CapacitanceUnits) float64 {
	switch toUnit { 
    case CapacitanceFarad:
		return a.Farads()
    case CapacitancePicofarad:
		return a.Picofarads()
    case CapacitanceNanofarad:
		return a.Nanofarads()
    case CapacitanceMicrofarad:
		return a.Microfarads()
    case CapacitanceMillifarad:
		return a.Millifarads()
    case CapacitanceKilofarad:
		return a.Kilofarads()
    case CapacitanceMegafarad:
		return a.Megafarads()
	default:
		return math.NaN()
	}
}

func (a *Capacitance) convertFromBase(toUnit CapacitanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case CapacitanceFarad:
		return (value) 
	case CapacitancePicofarad:
		return ((value) / 1e-12) 
	case CapacitanceNanofarad:
		return ((value) / 1e-09) 
	case CapacitanceMicrofarad:
		return ((value) / 1e-06) 
	case CapacitanceMillifarad:
		return ((value) / 0.001) 
	case CapacitanceKilofarad:
		return ((value) / 1000.0) 
	case CapacitanceMegafarad:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Capacitance) convertToBase(value float64, fromUnit CapacitanceUnits) float64 {
	switch fromUnit { 
	case CapacitanceFarad:
		return (value) 
	case CapacitancePicofarad:
		return ((value) * 1e-12) 
	case CapacitanceNanofarad:
		return ((value) * 1e-09) 
	case CapacitanceMicrofarad:
		return ((value) * 1e-06) 
	case CapacitanceMillifarad:
		return ((value) * 0.001) 
	case CapacitanceKilofarad:
		return ((value) * 1000.0) 
	case CapacitanceMegafarad:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Capacitance in the default unit (Farad),
// formatted to two decimal places.
func (a Capacitance) String() string {
	return a.ToString(CapacitanceFarad, 2)
}

// ToString formats the Capacitance value as a string with the specified unit and fractional digits.
// It converts the Capacitance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Capacitance value will be converted (e.g., Farad))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Capacitance with the unit abbreviation.
func (a *Capacitance) ToString(unit CapacitanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Capacitance) getUnitAbbreviation(unit CapacitanceUnits) string {
	switch unit { 
	case CapacitanceFarad:
		return "F" 
	case CapacitancePicofarad:
		return "pF" 
	case CapacitanceNanofarad:
		return "nF" 
	case CapacitanceMicrofarad:
		return "μF" 
	case CapacitanceMillifarad:
		return "mF" 
	case CapacitanceKilofarad:
		return "kF" 
	case CapacitanceMegafarad:
		return "MF" 
	default:
		return ""
	}
}

// Equals checks if the given Capacitance is equal to the current Capacitance.
//
// Parameters:
//    other: The Capacitance to compare against.
//
// Returns:
//    bool: Returns true if both Capacitance are equal, false otherwise.
func (a *Capacitance) Equals(other *Capacitance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Capacitance with another Capacitance.
// It returns -1 if the current Capacitance is less than the other Capacitance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Capacitance to compare against.
//
// Returns:
//    int: -1 if the current Capacitance is less, 1 if greater, and 0 if equal.
func (a *Capacitance) CompareTo(other *Capacitance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Capacitance to the current Capacitance and returns the result.
// The result is a new Capacitance instance with the sum of the values.
//
// Parameters:
//    other: The Capacitance to add to the current Capacitance.
//
// Returns:
//    *Capacitance: A new Capacitance instance representing the sum of both Capacitance.
func (a *Capacitance) Add(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Capacitance from the current Capacitance and returns the result.
// The result is a new Capacitance instance with the difference of the values.
//
// Parameters:
//    other: The Capacitance to subtract from the current Capacitance.
//
// Returns:
//    *Capacitance: A new Capacitance instance representing the difference of both Capacitance.
func (a *Capacitance) Subtract(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Capacitance by the given Capacitance and returns the result.
// The result is a new Capacitance instance with the product of the values.
//
// Parameters:
//    other: The Capacitance to multiply with the current Capacitance.
//
// Returns:
//    *Capacitance: A new Capacitance instance representing the product of both Capacitance.
func (a *Capacitance) Multiply(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value * other.BaseValue()}
}

// Divide divides the current Capacitance by the given Capacitance and returns the result.
// The result is a new Capacitance instance with the quotient of the values.
//
// Parameters:
//    other: The Capacitance to divide the current Capacitance by.
//
// Returns:
//    *Capacitance: A new Capacitance instance representing the quotient of both Capacitance.
func (a *Capacitance) Divide(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value / other.BaseValue()}
}