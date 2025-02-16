// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialAcUnits defines various units of ElectricPotentialAc.
type ElectricPotentialAcUnits string

const (
    
        // 
        ElectricPotentialAcVoltAc ElectricPotentialAcUnits = "VoltAc"
        // 
        ElectricPotentialAcMicrovoltAc ElectricPotentialAcUnits = "MicrovoltAc"
        // 
        ElectricPotentialAcMillivoltAc ElectricPotentialAcUnits = "MillivoltAc"
        // 
        ElectricPotentialAcKilovoltAc ElectricPotentialAcUnits = "KilovoltAc"
        // 
        ElectricPotentialAcMegavoltAc ElectricPotentialAcUnits = "MegavoltAc"
)

// ElectricPotentialAcDto represents a ElectricPotentialAc measurement with a numerical value and its corresponding unit.
type ElectricPotentialAcDto struct {
    // Value is the numerical representation of the ElectricPotentialAc.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricPotentialAc, as defined in the ElectricPotentialAcUnits enumeration.
	Unit  ElectricPotentialAcUnits `json:"unit"`
}

// ElectricPotentialAcDtoFactory groups methods for creating and serializing ElectricPotentialAcDto objects.
type ElectricPotentialAcDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialAcDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricPotentialAcDtoFactory) FromJSON(data []byte) (*ElectricPotentialAcDto, error) {
	a := ElectricPotentialAcDto{}

    // Parse JSON into ElectricPotentialAcDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricPotentialAcDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricPotentialAcDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricPotentialAc represents a measurement in a of ElectricPotentialAc.
//
// The Electric Potential of a system known to use Alternating Current.
type ElectricPotentialAc struct {
	// value is the base measurement stored internally.
	value       float64
    
    volts_acLazy *float64 
    microvolts_acLazy *float64 
    millivolts_acLazy *float64 
    kilovolts_acLazy *float64 
    megavolts_acLazy *float64 
}

// ElectricPotentialAcFactory groups methods for creating ElectricPotentialAc instances.
type ElectricPotentialAcFactory struct{}

// CreateElectricPotentialAc creates a new ElectricPotentialAc instance from the given value and unit.
func (uf ElectricPotentialAcFactory) CreateElectricPotentialAc(value float64, unit ElectricPotentialAcUnits) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, unit)
}

// FromDto converts a ElectricPotentialAcDto to a ElectricPotentialAc instance.
func (uf ElectricPotentialAcFactory) FromDto(dto ElectricPotentialAcDto) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialAc instance.
func (uf ElectricPotentialAcFactory) FromDtoJSON(data []byte) (*ElectricPotentialAc, error) {
	unitDto, err := ElectricPotentialAcDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricPotentialAcDto from JSON: %w", err)
	}
	return ElectricPotentialAcFactory{}.FromDto(*unitDto)
}


// FromVoltsAc creates a new ElectricPotentialAc instance from a value in VoltsAc.
func (uf ElectricPotentialAcFactory) FromVoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcVoltAc)
}

// FromMicrovoltsAc creates a new ElectricPotentialAc instance from a value in MicrovoltsAc.
func (uf ElectricPotentialAcFactory) FromMicrovoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMicrovoltAc)
}

// FromMillivoltsAc creates a new ElectricPotentialAc instance from a value in MillivoltsAc.
func (uf ElectricPotentialAcFactory) FromMillivoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMillivoltAc)
}

// FromKilovoltsAc creates a new ElectricPotentialAc instance from a value in KilovoltsAc.
func (uf ElectricPotentialAcFactory) FromKilovoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcKilovoltAc)
}

// FromMegavoltsAc creates a new ElectricPotentialAc instance from a value in MegavoltsAc.
func (uf ElectricPotentialAcFactory) FromMegavoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMegavoltAc)
}


// newElectricPotentialAc creates a new ElectricPotentialAc.
func newElectricPotentialAc(value float64, fromUnit ElectricPotentialAcUnits) (*ElectricPotentialAc, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotentialAc{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialAc in VoltAc unit (the base/default quantity).
func (a *ElectricPotentialAc) BaseValue() float64 {
	return a.value
}


// VoltsAc returns the ElectricPotentialAc value in VoltsAc.
//
// 
func (a *ElectricPotentialAc) VoltsAc() float64 {
	if a.volts_acLazy != nil {
		return *a.volts_acLazy
	}
	volts_ac := a.convertFromBase(ElectricPotentialAcVoltAc)
	a.volts_acLazy = &volts_ac
	return volts_ac
}

// MicrovoltsAc returns the ElectricPotentialAc value in MicrovoltsAc.
//
// 
func (a *ElectricPotentialAc) MicrovoltsAc() float64 {
	if a.microvolts_acLazy != nil {
		return *a.microvolts_acLazy
	}
	microvolts_ac := a.convertFromBase(ElectricPotentialAcMicrovoltAc)
	a.microvolts_acLazy = &microvolts_ac
	return microvolts_ac
}

// MillivoltsAc returns the ElectricPotentialAc value in MillivoltsAc.
//
// 
func (a *ElectricPotentialAc) MillivoltsAc() float64 {
	if a.millivolts_acLazy != nil {
		return *a.millivolts_acLazy
	}
	millivolts_ac := a.convertFromBase(ElectricPotentialAcMillivoltAc)
	a.millivolts_acLazy = &millivolts_ac
	return millivolts_ac
}

// KilovoltsAc returns the ElectricPotentialAc value in KilovoltsAc.
//
// 
func (a *ElectricPotentialAc) KilovoltsAc() float64 {
	if a.kilovolts_acLazy != nil {
		return *a.kilovolts_acLazy
	}
	kilovolts_ac := a.convertFromBase(ElectricPotentialAcKilovoltAc)
	a.kilovolts_acLazy = &kilovolts_ac
	return kilovolts_ac
}

// MegavoltsAc returns the ElectricPotentialAc value in MegavoltsAc.
//
// 
func (a *ElectricPotentialAc) MegavoltsAc() float64 {
	if a.megavolts_acLazy != nil {
		return *a.megavolts_acLazy
	}
	megavolts_ac := a.convertFromBase(ElectricPotentialAcMegavoltAc)
	a.megavolts_acLazy = &megavolts_ac
	return megavolts_ac
}


// ToDto creates a ElectricPotentialAcDto representation from the ElectricPotentialAc instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltAc by default.
func (a *ElectricPotentialAc) ToDto(holdInUnit *ElectricPotentialAcUnits) ElectricPotentialAcDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialAcVoltAc // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialAcDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricPotentialAc instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltAc by default.
func (a *ElectricPotentialAc) ToDtoJSON(holdInUnit *ElectricPotentialAcUnits) ([]byte, error) {
	// Convert to ElectricPotentialAcDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricPotentialAc to a specific unit value.
// The function uses the provided unit type (ElectricPotentialAcUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricPotentialAc) Convert(toUnit ElectricPotentialAcUnits) float64 {
	switch toUnit { 
    case ElectricPotentialAcVoltAc:
		return a.VoltsAc()
    case ElectricPotentialAcMicrovoltAc:
		return a.MicrovoltsAc()
    case ElectricPotentialAcMillivoltAc:
		return a.MillivoltsAc()
    case ElectricPotentialAcKilovoltAc:
		return a.KilovoltsAc()
    case ElectricPotentialAcMegavoltAc:
		return a.MegavoltsAc()
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialAc) convertFromBase(toUnit ElectricPotentialAcUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialAcVoltAc:
		return (value) 
	case ElectricPotentialAcMicrovoltAc:
		return ((value) / 1e-06) 
	case ElectricPotentialAcMillivoltAc:
		return ((value) / 0.001) 
	case ElectricPotentialAcKilovoltAc:
		return ((value) / 1000.0) 
	case ElectricPotentialAcMegavoltAc:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialAc) convertToBase(value float64, fromUnit ElectricPotentialAcUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialAcVoltAc:
		return (value) 
	case ElectricPotentialAcMicrovoltAc:
		return ((value) * 1e-06) 
	case ElectricPotentialAcMillivoltAc:
		return ((value) * 0.001) 
	case ElectricPotentialAcKilovoltAc:
		return ((value) * 1000.0) 
	case ElectricPotentialAcMegavoltAc:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricPotentialAc in the default unit (VoltAc),
// formatted to two decimal places.
func (a ElectricPotentialAc) String() string {
	return a.ToString(ElectricPotentialAcVoltAc, 2)
}

// ToString formats the ElectricPotentialAc value as a string with the specified unit and fractional digits.
// It converts the ElectricPotentialAc to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricPotentialAc value will be converted (e.g., VoltAc))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricPotentialAc with the unit abbreviation.
func (a *ElectricPotentialAc) ToString(unit ElectricPotentialAcUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricPotentialAcAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricPotentialAcAbbreviation(unit))
}

// Equals checks if the given ElectricPotentialAc is equal to the current ElectricPotentialAc.
//
// Parameters:
//    other: The ElectricPotentialAc to compare against.
//
// Returns:
//    bool: Returns true if both ElectricPotentialAc are equal, false otherwise.
func (a *ElectricPotentialAc) Equals(other *ElectricPotentialAc) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricPotentialAc with another ElectricPotentialAc.
// It returns -1 if the current ElectricPotentialAc is less than the other ElectricPotentialAc, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricPotentialAc to compare against.
//
// Returns:
//    int: -1 if the current ElectricPotentialAc is less, 1 if greater, and 0 if equal.
func (a *ElectricPotentialAc) CompareTo(other *ElectricPotentialAc) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricPotentialAc to the current ElectricPotentialAc and returns the result.
// The result is a new ElectricPotentialAc instance with the sum of the values.
//
// Parameters:
//    other: The ElectricPotentialAc to add to the current ElectricPotentialAc.
//
// Returns:
//    *ElectricPotentialAc: A new ElectricPotentialAc instance representing the sum of both ElectricPotentialAc.
func (a *ElectricPotentialAc) Add(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricPotentialAc from the current ElectricPotentialAc and returns the result.
// The result is a new ElectricPotentialAc instance with the difference of the values.
//
// Parameters:
//    other: The ElectricPotentialAc to subtract from the current ElectricPotentialAc.
//
// Returns:
//    *ElectricPotentialAc: A new ElectricPotentialAc instance representing the difference of both ElectricPotentialAc.
func (a *ElectricPotentialAc) Subtract(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricPotentialAc by the given ElectricPotentialAc and returns the result.
// The result is a new ElectricPotentialAc instance with the product of the values.
//
// Parameters:
//    other: The ElectricPotentialAc to multiply with the current ElectricPotentialAc.
//
// Returns:
//    *ElectricPotentialAc: A new ElectricPotentialAc instance representing the product of both ElectricPotentialAc.
func (a *ElectricPotentialAc) Multiply(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricPotentialAc by the given ElectricPotentialAc and returns the result.
// The result is a new ElectricPotentialAc instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricPotentialAc to divide the current ElectricPotentialAc by.
//
// Returns:
//    *ElectricPotentialAc: A new ElectricPotentialAc instance representing the quotient of both ElectricPotentialAc.
func (a *ElectricPotentialAc) Divide(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value / other.BaseValue()}
}

// GetElectricPotentialAcAbbreviation gets the unit abbreviation.
func GetElectricPotentialAcAbbreviation(unit ElectricPotentialAcUnits) string {
	switch unit { 
	case ElectricPotentialAcVoltAc:
		return "Vac" 
	case ElectricPotentialAcMicrovoltAc:
		return "μVac" 
	case ElectricPotentialAcMillivoltAc:
		return "mVac" 
	case ElectricPotentialAcKilovoltAc:
		return "kVac" 
	case ElectricPotentialAcMegavoltAc:
		return "MVac" 
	default:
		return ""
	}
}