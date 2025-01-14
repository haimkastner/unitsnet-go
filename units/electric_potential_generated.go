// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialUnits defines various units of ElectricPotential.
type ElectricPotentialUnits string

const (
    
        // 
        ElectricPotentialVolt ElectricPotentialUnits = "Volt"
        // 
        ElectricPotentialNanovolt ElectricPotentialUnits = "Nanovolt"
        // 
        ElectricPotentialMicrovolt ElectricPotentialUnits = "Microvolt"
        // 
        ElectricPotentialMillivolt ElectricPotentialUnits = "Millivolt"
        // 
        ElectricPotentialKilovolt ElectricPotentialUnits = "Kilovolt"
        // 
        ElectricPotentialMegavolt ElectricPotentialUnits = "Megavolt"
)

// ElectricPotentialDto represents a ElectricPotential measurement with a numerical value and its corresponding unit.
type ElectricPotentialDto struct {
    // Value is the numerical representation of the ElectricPotential.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricPotential, as defined in the ElectricPotentialUnits enumeration.
	Unit  ElectricPotentialUnits `json:"unit"`
}

// ElectricPotentialDtoFactory groups methods for creating and serializing ElectricPotentialDto objects.
type ElectricPotentialDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricPotentialDtoFactory) FromJSON(data []byte) (*ElectricPotentialDto, error) {
	a := ElectricPotentialDto{}

    // Parse JSON into ElectricPotentialDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricPotentialDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricPotentialDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricPotential represents a measurement in a of ElectricPotential.
//
// In classical electromagnetism, the electric potential (a scalar quantity denoted by Φ, ΦE or V and also called the electric field potential or the electrostatic potential) at a point is the amount of electric potential energy that a unitary point charge would have when located at that point.
type ElectricPotential struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltsLazy *float64 
    nanovoltsLazy *float64 
    microvoltsLazy *float64 
    millivoltsLazy *float64 
    kilovoltsLazy *float64 
    megavoltsLazy *float64 
}

// ElectricPotentialFactory groups methods for creating ElectricPotential instances.
type ElectricPotentialFactory struct{}

// CreateElectricPotential creates a new ElectricPotential instance from the given value and unit.
func (uf ElectricPotentialFactory) CreateElectricPotential(value float64, unit ElectricPotentialUnits) (*ElectricPotential, error) {
	return newElectricPotential(value, unit)
}

// FromDto converts a ElectricPotentialDto to a ElectricPotential instance.
func (uf ElectricPotentialFactory) FromDto(dto ElectricPotentialDto) (*ElectricPotential, error) {
	return newElectricPotential(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotential instance.
func (uf ElectricPotentialFactory) FromDtoJSON(data []byte) (*ElectricPotential, error) {
	unitDto, err := ElectricPotentialDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricPotentialDto from JSON: %w", err)
	}
	return ElectricPotentialFactory{}.FromDto(*unitDto)
}


// FromVolts creates a new ElectricPotential instance from a value in Volts.
func (uf ElectricPotentialFactory) FromVolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialVolt)
}

// FromNanovolts creates a new ElectricPotential instance from a value in Nanovolts.
func (uf ElectricPotentialFactory) FromNanovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialNanovolt)
}

// FromMicrovolts creates a new ElectricPotential instance from a value in Microvolts.
func (uf ElectricPotentialFactory) FromMicrovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMicrovolt)
}

// FromMillivolts creates a new ElectricPotential instance from a value in Millivolts.
func (uf ElectricPotentialFactory) FromMillivolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMillivolt)
}

// FromKilovolts creates a new ElectricPotential instance from a value in Kilovolts.
func (uf ElectricPotentialFactory) FromKilovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialKilovolt)
}

// FromMegavolts creates a new ElectricPotential instance from a value in Megavolts.
func (uf ElectricPotentialFactory) FromMegavolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMegavolt)
}


// newElectricPotential creates a new ElectricPotential.
func newElectricPotential(value float64, fromUnit ElectricPotentialUnits) (*ElectricPotential, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotential{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotential in Volt unit (the base/default quantity).
func (a *ElectricPotential) BaseValue() float64 {
	return a.value
}


// Volts returns the ElectricPotential value in Volts.
//
// 
func (a *ElectricPotential) Volts() float64 {
	if a.voltsLazy != nil {
		return *a.voltsLazy
	}
	volts := a.convertFromBase(ElectricPotentialVolt)
	a.voltsLazy = &volts
	return volts
}

// Nanovolts returns the ElectricPotential value in Nanovolts.
//
// 
func (a *ElectricPotential) Nanovolts() float64 {
	if a.nanovoltsLazy != nil {
		return *a.nanovoltsLazy
	}
	nanovolts := a.convertFromBase(ElectricPotentialNanovolt)
	a.nanovoltsLazy = &nanovolts
	return nanovolts
}

// Microvolts returns the ElectricPotential value in Microvolts.
//
// 
func (a *ElectricPotential) Microvolts() float64 {
	if a.microvoltsLazy != nil {
		return *a.microvoltsLazy
	}
	microvolts := a.convertFromBase(ElectricPotentialMicrovolt)
	a.microvoltsLazy = &microvolts
	return microvolts
}

// Millivolts returns the ElectricPotential value in Millivolts.
//
// 
func (a *ElectricPotential) Millivolts() float64 {
	if a.millivoltsLazy != nil {
		return *a.millivoltsLazy
	}
	millivolts := a.convertFromBase(ElectricPotentialMillivolt)
	a.millivoltsLazy = &millivolts
	return millivolts
}

// Kilovolts returns the ElectricPotential value in Kilovolts.
//
// 
func (a *ElectricPotential) Kilovolts() float64 {
	if a.kilovoltsLazy != nil {
		return *a.kilovoltsLazy
	}
	kilovolts := a.convertFromBase(ElectricPotentialKilovolt)
	a.kilovoltsLazy = &kilovolts
	return kilovolts
}

// Megavolts returns the ElectricPotential value in Megavolts.
//
// 
func (a *ElectricPotential) Megavolts() float64 {
	if a.megavoltsLazy != nil {
		return *a.megavoltsLazy
	}
	megavolts := a.convertFromBase(ElectricPotentialMegavolt)
	a.megavoltsLazy = &megavolts
	return megavolts
}


// ToDto creates a ElectricPotentialDto representation from the ElectricPotential instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Volt by default.
func (a *ElectricPotential) ToDto(holdInUnit *ElectricPotentialUnits) ElectricPotentialDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialVolt // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricPotential instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Volt by default.
func (a *ElectricPotential) ToDtoJSON(holdInUnit *ElectricPotentialUnits) ([]byte, error) {
	// Convert to ElectricPotentialDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricPotential to a specific unit value.
// The function uses the provided unit type (ElectricPotentialUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricPotential) Convert(toUnit ElectricPotentialUnits) float64 {
	switch toUnit { 
    case ElectricPotentialVolt:
		return a.Volts()
    case ElectricPotentialNanovolt:
		return a.Nanovolts()
    case ElectricPotentialMicrovolt:
		return a.Microvolts()
    case ElectricPotentialMillivolt:
		return a.Millivolts()
    case ElectricPotentialKilovolt:
		return a.Kilovolts()
    case ElectricPotentialMegavolt:
		return a.Megavolts()
	default:
		return math.NaN()
	}
}

func (a *ElectricPotential) convertFromBase(toUnit ElectricPotentialUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialVolt:
		return (value) 
	case ElectricPotentialNanovolt:
		return ((value) / 1e-09) 
	case ElectricPotentialMicrovolt:
		return ((value) / 1e-06) 
	case ElectricPotentialMillivolt:
		return ((value) / 0.001) 
	case ElectricPotentialKilovolt:
		return ((value) / 1000.0) 
	case ElectricPotentialMegavolt:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotential) convertToBase(value float64, fromUnit ElectricPotentialUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialVolt:
		return (value) 
	case ElectricPotentialNanovolt:
		return ((value) * 1e-09) 
	case ElectricPotentialMicrovolt:
		return ((value) * 1e-06) 
	case ElectricPotentialMillivolt:
		return ((value) * 0.001) 
	case ElectricPotentialKilovolt:
		return ((value) * 1000.0) 
	case ElectricPotentialMegavolt:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricPotential in the default unit (Volt),
// formatted to two decimal places.
func (a ElectricPotential) String() string {
	return a.ToString(ElectricPotentialVolt, 2)
}

// ToString formats the ElectricPotential value as a string with the specified unit and fractional digits.
// It converts the ElectricPotential to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricPotential value will be converted (e.g., Volt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricPotential with the unit abbreviation.
func (a *ElectricPotential) ToString(unit ElectricPotentialUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricPotential) getUnitAbbreviation(unit ElectricPotentialUnits) string {
	switch unit { 
	case ElectricPotentialVolt:
		return "V" 
	case ElectricPotentialNanovolt:
		return "nV" 
	case ElectricPotentialMicrovolt:
		return "μV" 
	case ElectricPotentialMillivolt:
		return "mV" 
	case ElectricPotentialKilovolt:
		return "kV" 
	case ElectricPotentialMegavolt:
		return "MV" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricPotential is equal to the current ElectricPotential.
//
// Parameters:
//    other: The ElectricPotential to compare against.
//
// Returns:
//    bool: Returns true if both ElectricPotential are equal, false otherwise.
func (a *ElectricPotential) Equals(other *ElectricPotential) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricPotential with another ElectricPotential.
// It returns -1 if the current ElectricPotential is less than the other ElectricPotential, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricPotential to compare against.
//
// Returns:
//    int: -1 if the current ElectricPotential is less, 1 if greater, and 0 if equal.
func (a *ElectricPotential) CompareTo(other *ElectricPotential) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricPotential to the current ElectricPotential and returns the result.
// The result is a new ElectricPotential instance with the sum of the values.
//
// Parameters:
//    other: The ElectricPotential to add to the current ElectricPotential.
//
// Returns:
//    *ElectricPotential: A new ElectricPotential instance representing the sum of both ElectricPotential.
func (a *ElectricPotential) Add(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricPotential from the current ElectricPotential and returns the result.
// The result is a new ElectricPotential instance with the difference of the values.
//
// Parameters:
//    other: The ElectricPotential to subtract from the current ElectricPotential.
//
// Returns:
//    *ElectricPotential: A new ElectricPotential instance representing the difference of both ElectricPotential.
func (a *ElectricPotential) Subtract(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricPotential by the given ElectricPotential and returns the result.
// The result is a new ElectricPotential instance with the product of the values.
//
// Parameters:
//    other: The ElectricPotential to multiply with the current ElectricPotential.
//
// Returns:
//    *ElectricPotential: A new ElectricPotential instance representing the product of both ElectricPotential.
func (a *ElectricPotential) Multiply(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricPotential by the given ElectricPotential and returns the result.
// The result is a new ElectricPotential instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricPotential to divide the current ElectricPotential by.
//
// Returns:
//    *ElectricPotential: A new ElectricPotential instance representing the quotient of both ElectricPotential.
func (a *ElectricPotential) Divide(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value / other.BaseValue()}
}