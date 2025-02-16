// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialDcUnits defines various units of ElectricPotentialDc.
type ElectricPotentialDcUnits string

const (
    
        // 
        ElectricPotentialDcVoltDc ElectricPotentialDcUnits = "VoltDc"
        // 
        ElectricPotentialDcMicrovoltDc ElectricPotentialDcUnits = "MicrovoltDc"
        // 
        ElectricPotentialDcMillivoltDc ElectricPotentialDcUnits = "MillivoltDc"
        // 
        ElectricPotentialDcKilovoltDc ElectricPotentialDcUnits = "KilovoltDc"
        // 
        ElectricPotentialDcMegavoltDc ElectricPotentialDcUnits = "MegavoltDc"
)

// ElectricPotentialDcDto represents a ElectricPotentialDc measurement with a numerical value and its corresponding unit.
type ElectricPotentialDcDto struct {
    // Value is the numerical representation of the ElectricPotentialDc.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricPotentialDc, as defined in the ElectricPotentialDcUnits enumeration.
	Unit  ElectricPotentialDcUnits `json:"unit"`
}

// ElectricPotentialDcDtoFactory groups methods for creating and serializing ElectricPotentialDcDto objects.
type ElectricPotentialDcDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialDcDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricPotentialDcDtoFactory) FromJSON(data []byte) (*ElectricPotentialDcDto, error) {
	a := ElectricPotentialDcDto{}

    // Parse JSON into ElectricPotentialDcDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricPotentialDcDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricPotentialDcDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricPotentialDc represents a measurement in a of ElectricPotentialDc.
//
// The Electric Potential of a system known to use Direct Current.
type ElectricPotentialDc struct {
	// value is the base measurement stored internally.
	value       float64
    
    volts_dcLazy *float64 
    microvolts_dcLazy *float64 
    millivolts_dcLazy *float64 
    kilovolts_dcLazy *float64 
    megavolts_dcLazy *float64 
}

// ElectricPotentialDcFactory groups methods for creating ElectricPotentialDc instances.
type ElectricPotentialDcFactory struct{}

// CreateElectricPotentialDc creates a new ElectricPotentialDc instance from the given value and unit.
func (uf ElectricPotentialDcFactory) CreateElectricPotentialDc(value float64, unit ElectricPotentialDcUnits) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, unit)
}

// FromDto converts a ElectricPotentialDcDto to a ElectricPotentialDc instance.
func (uf ElectricPotentialDcFactory) FromDto(dto ElectricPotentialDcDto) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialDc instance.
func (uf ElectricPotentialDcFactory) FromDtoJSON(data []byte) (*ElectricPotentialDc, error) {
	unitDto, err := ElectricPotentialDcDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricPotentialDcDto from JSON: %w", err)
	}
	return ElectricPotentialDcFactory{}.FromDto(*unitDto)
}


// FromVoltsDc creates a new ElectricPotentialDc instance from a value in VoltsDc.
func (uf ElectricPotentialDcFactory) FromVoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcVoltDc)
}

// FromMicrovoltsDc creates a new ElectricPotentialDc instance from a value in MicrovoltsDc.
func (uf ElectricPotentialDcFactory) FromMicrovoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMicrovoltDc)
}

// FromMillivoltsDc creates a new ElectricPotentialDc instance from a value in MillivoltsDc.
func (uf ElectricPotentialDcFactory) FromMillivoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMillivoltDc)
}

// FromKilovoltsDc creates a new ElectricPotentialDc instance from a value in KilovoltsDc.
func (uf ElectricPotentialDcFactory) FromKilovoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcKilovoltDc)
}

// FromMegavoltsDc creates a new ElectricPotentialDc instance from a value in MegavoltsDc.
func (uf ElectricPotentialDcFactory) FromMegavoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMegavoltDc)
}


// newElectricPotentialDc creates a new ElectricPotentialDc.
func newElectricPotentialDc(value float64, fromUnit ElectricPotentialDcUnits) (*ElectricPotentialDc, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotentialDc{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialDc in VoltDc unit (the base/default quantity).
func (a *ElectricPotentialDc) BaseValue() float64 {
	return a.value
}


// VoltsDc returns the ElectricPotentialDc value in VoltsDc.
//
// 
func (a *ElectricPotentialDc) VoltsDc() float64 {
	if a.volts_dcLazy != nil {
		return *a.volts_dcLazy
	}
	volts_dc := a.convertFromBase(ElectricPotentialDcVoltDc)
	a.volts_dcLazy = &volts_dc
	return volts_dc
}

// MicrovoltsDc returns the ElectricPotentialDc value in MicrovoltsDc.
//
// 
func (a *ElectricPotentialDc) MicrovoltsDc() float64 {
	if a.microvolts_dcLazy != nil {
		return *a.microvolts_dcLazy
	}
	microvolts_dc := a.convertFromBase(ElectricPotentialDcMicrovoltDc)
	a.microvolts_dcLazy = &microvolts_dc
	return microvolts_dc
}

// MillivoltsDc returns the ElectricPotentialDc value in MillivoltsDc.
//
// 
func (a *ElectricPotentialDc) MillivoltsDc() float64 {
	if a.millivolts_dcLazy != nil {
		return *a.millivolts_dcLazy
	}
	millivolts_dc := a.convertFromBase(ElectricPotentialDcMillivoltDc)
	a.millivolts_dcLazy = &millivolts_dc
	return millivolts_dc
}

// KilovoltsDc returns the ElectricPotentialDc value in KilovoltsDc.
//
// 
func (a *ElectricPotentialDc) KilovoltsDc() float64 {
	if a.kilovolts_dcLazy != nil {
		return *a.kilovolts_dcLazy
	}
	kilovolts_dc := a.convertFromBase(ElectricPotentialDcKilovoltDc)
	a.kilovolts_dcLazy = &kilovolts_dc
	return kilovolts_dc
}

// MegavoltsDc returns the ElectricPotentialDc value in MegavoltsDc.
//
// 
func (a *ElectricPotentialDc) MegavoltsDc() float64 {
	if a.megavolts_dcLazy != nil {
		return *a.megavolts_dcLazy
	}
	megavolts_dc := a.convertFromBase(ElectricPotentialDcMegavoltDc)
	a.megavolts_dcLazy = &megavolts_dc
	return megavolts_dc
}


// ToDto creates a ElectricPotentialDcDto representation from the ElectricPotentialDc instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltDc by default.
func (a *ElectricPotentialDc) ToDto(holdInUnit *ElectricPotentialDcUnits) ElectricPotentialDcDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialDcVoltDc // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialDcDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricPotentialDc instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltDc by default.
func (a *ElectricPotentialDc) ToDtoJSON(holdInUnit *ElectricPotentialDcUnits) ([]byte, error) {
	// Convert to ElectricPotentialDcDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricPotentialDc to a specific unit value.
// The function uses the provided unit type (ElectricPotentialDcUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricPotentialDc) Convert(toUnit ElectricPotentialDcUnits) float64 {
	switch toUnit { 
    case ElectricPotentialDcVoltDc:
		return a.VoltsDc()
    case ElectricPotentialDcMicrovoltDc:
		return a.MicrovoltsDc()
    case ElectricPotentialDcMillivoltDc:
		return a.MillivoltsDc()
    case ElectricPotentialDcKilovoltDc:
		return a.KilovoltsDc()
    case ElectricPotentialDcMegavoltDc:
		return a.MegavoltsDc()
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialDc) convertFromBase(toUnit ElectricPotentialDcUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialDcVoltDc:
		return (value) 
	case ElectricPotentialDcMicrovoltDc:
		return ((value) / 1e-06) 
	case ElectricPotentialDcMillivoltDc:
		return ((value) / 0.001) 
	case ElectricPotentialDcKilovoltDc:
		return ((value) / 1000.0) 
	case ElectricPotentialDcMegavoltDc:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialDc) convertToBase(value float64, fromUnit ElectricPotentialDcUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialDcVoltDc:
		return (value) 
	case ElectricPotentialDcMicrovoltDc:
		return ((value) * 1e-06) 
	case ElectricPotentialDcMillivoltDc:
		return ((value) * 0.001) 
	case ElectricPotentialDcKilovoltDc:
		return ((value) * 1000.0) 
	case ElectricPotentialDcMegavoltDc:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricPotentialDc in the default unit (VoltDc),
// formatted to two decimal places.
func (a ElectricPotentialDc) String() string {
	return a.ToString(ElectricPotentialDcVoltDc, 2)
}

// ToString formats the ElectricPotentialDc value as a string with the specified unit and fractional digits.
// It converts the ElectricPotentialDc to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricPotentialDc value will be converted (e.g., VoltDc))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricPotentialDc with the unit abbreviation.
func (a *ElectricPotentialDc) ToString(unit ElectricPotentialDcUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricPotentialDcAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricPotentialDcAbbreviation(unit))
}

// Equals checks if the given ElectricPotentialDc is equal to the current ElectricPotentialDc.
//
// Parameters:
//    other: The ElectricPotentialDc to compare against.
//
// Returns:
//    bool: Returns true if both ElectricPotentialDc are equal, false otherwise.
func (a *ElectricPotentialDc) Equals(other *ElectricPotentialDc) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricPotentialDc with another ElectricPotentialDc.
// It returns -1 if the current ElectricPotentialDc is less than the other ElectricPotentialDc, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricPotentialDc to compare against.
//
// Returns:
//    int: -1 if the current ElectricPotentialDc is less, 1 if greater, and 0 if equal.
func (a *ElectricPotentialDc) CompareTo(other *ElectricPotentialDc) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricPotentialDc to the current ElectricPotentialDc and returns the result.
// The result is a new ElectricPotentialDc instance with the sum of the values.
//
// Parameters:
//    other: The ElectricPotentialDc to add to the current ElectricPotentialDc.
//
// Returns:
//    *ElectricPotentialDc: A new ElectricPotentialDc instance representing the sum of both ElectricPotentialDc.
func (a *ElectricPotentialDc) Add(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricPotentialDc from the current ElectricPotentialDc and returns the result.
// The result is a new ElectricPotentialDc instance with the difference of the values.
//
// Parameters:
//    other: The ElectricPotentialDc to subtract from the current ElectricPotentialDc.
//
// Returns:
//    *ElectricPotentialDc: A new ElectricPotentialDc instance representing the difference of both ElectricPotentialDc.
func (a *ElectricPotentialDc) Subtract(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricPotentialDc by the given ElectricPotentialDc and returns the result.
// The result is a new ElectricPotentialDc instance with the product of the values.
//
// Parameters:
//    other: The ElectricPotentialDc to multiply with the current ElectricPotentialDc.
//
// Returns:
//    *ElectricPotentialDc: A new ElectricPotentialDc instance representing the product of both ElectricPotentialDc.
func (a *ElectricPotentialDc) Multiply(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricPotentialDc by the given ElectricPotentialDc and returns the result.
// The result is a new ElectricPotentialDc instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricPotentialDc to divide the current ElectricPotentialDc by.
//
// Returns:
//    *ElectricPotentialDc: A new ElectricPotentialDc instance representing the quotient of both ElectricPotentialDc.
func (a *ElectricPotentialDc) Divide(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value / other.BaseValue()}
}

// GetElectricPotentialDcAbbreviation gets the unit abbreviation.
func GetElectricPotentialDcAbbreviation(unit ElectricPotentialDcUnits) string {
	switch unit { 
	case ElectricPotentialDcVoltDc:
		return "Vdc" 
	case ElectricPotentialDcMicrovoltDc:
		return "μVdc" 
	case ElectricPotentialDcMillivoltDc:
		return "mVdc" 
	case ElectricPotentialDcKilovoltDc:
		return "kVdc" 
	case ElectricPotentialDcMegavoltDc:
		return "MVdc" 
	default:
		return ""
	}
}