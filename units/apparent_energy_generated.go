// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ApparentEnergyUnits defines various units of ApparentEnergy.
type ApparentEnergyUnits string

const (
    
        // 
        ApparentEnergyVoltampereHour ApparentEnergyUnits = "VoltampereHour"
        // 
        ApparentEnergyKilovoltampereHour ApparentEnergyUnits = "KilovoltampereHour"
        // 
        ApparentEnergyMegavoltampereHour ApparentEnergyUnits = "MegavoltampereHour"
)

var internalApparentEnergyUnitsMap = map[ApparentEnergyUnits]bool{
	
	ApparentEnergyVoltampereHour: true,
	ApparentEnergyKilovoltampereHour: true,
	ApparentEnergyMegavoltampereHour: true,
}

// ApparentEnergyDto represents a ApparentEnergy measurement with a numerical value and its corresponding unit.
type ApparentEnergyDto struct {
    // Value is the numerical representation of the ApparentEnergy.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ApparentEnergy, as defined in the ApparentEnergyUnits enumeration.
	Unit  ApparentEnergyUnits `json:"unit" validate:"required,oneof=VoltampereHour KilovoltampereHour MegavoltampereHour"`
}

// ApparentEnergyDtoFactory groups methods for creating and serializing ApparentEnergyDto objects.
type ApparentEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ApparentEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ApparentEnergyDtoFactory) FromJSON(data []byte) (*ApparentEnergyDto, error) {
	a := ApparentEnergyDto{}

    // Parse JSON into ApparentEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ApparentEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ApparentEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ApparentEnergy represents a measurement in a of ApparentEnergy.
//
// A unit for expressing the integral of apparent power over time, equal to the product of 1 volt-ampere and 1 hour, or to 3600 joules.
type ApparentEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltampere_hoursLazy *float64 
    kilovoltampere_hoursLazy *float64 
    megavoltampere_hoursLazy *float64 
}

// ApparentEnergyFactory groups methods for creating ApparentEnergy instances.
type ApparentEnergyFactory struct{}

// CreateApparentEnergy creates a new ApparentEnergy instance from the given value and unit.
func (uf ApparentEnergyFactory) CreateApparentEnergy(value float64, unit ApparentEnergyUnits) (*ApparentEnergy, error) {
	return newApparentEnergy(value, unit)
}

// FromDto converts a ApparentEnergyDto to a ApparentEnergy instance.
func (uf ApparentEnergyFactory) FromDto(dto ApparentEnergyDto) (*ApparentEnergy, error) {
	return newApparentEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ApparentEnergy instance.
func (uf ApparentEnergyFactory) FromDtoJSON(data []byte) (*ApparentEnergy, error) {
	unitDto, err := ApparentEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ApparentEnergyDto from JSON: %w", err)
	}
	return ApparentEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereHours creates a new ApparentEnergy instance from a value in VoltampereHours.
func (uf ApparentEnergyFactory) FromVoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyVoltampereHour)
}

// FromKilovoltampereHours creates a new ApparentEnergy instance from a value in KilovoltampereHours.
func (uf ApparentEnergyFactory) FromKilovoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyKilovoltampereHour)
}

// FromMegavoltampereHours creates a new ApparentEnergy instance from a value in MegavoltampereHours.
func (uf ApparentEnergyFactory) FromMegavoltampereHours(value float64) (*ApparentEnergy, error) {
	return newApparentEnergy(value, ApparentEnergyMegavoltampereHour)
}


// newApparentEnergy creates a new ApparentEnergy.
func newApparentEnergy(value float64, fromUnit ApparentEnergyUnits) (*ApparentEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalApparentEnergyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ApparentEnergyUnits", fromUnit)
	}
	a := &ApparentEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ApparentEnergy in VoltampereHour unit (the base/default quantity).
func (a *ApparentEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereHours returns the ApparentEnergy value in VoltampereHours.
//
// 
func (a *ApparentEnergy) VoltampereHours() float64 {
	if a.voltampere_hoursLazy != nil {
		return *a.voltampere_hoursLazy
	}
	voltampere_hours := a.convertFromBase(ApparentEnergyVoltampereHour)
	a.voltampere_hoursLazy = &voltampere_hours
	return voltampere_hours
}

// KilovoltampereHours returns the ApparentEnergy value in KilovoltampereHours.
//
// 
func (a *ApparentEnergy) KilovoltampereHours() float64 {
	if a.kilovoltampere_hoursLazy != nil {
		return *a.kilovoltampere_hoursLazy
	}
	kilovoltampere_hours := a.convertFromBase(ApparentEnergyKilovoltampereHour)
	a.kilovoltampere_hoursLazy = &kilovoltampere_hours
	return kilovoltampere_hours
}

// MegavoltampereHours returns the ApparentEnergy value in MegavoltampereHours.
//
// 
func (a *ApparentEnergy) MegavoltampereHours() float64 {
	if a.megavoltampere_hoursLazy != nil {
		return *a.megavoltampere_hoursLazy
	}
	megavoltampere_hours := a.convertFromBase(ApparentEnergyMegavoltampereHour)
	a.megavoltampere_hoursLazy = &megavoltampere_hours
	return megavoltampere_hours
}


// ToDto creates a ApparentEnergyDto representation from the ApparentEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereHour by default.
func (a *ApparentEnergy) ToDto(holdInUnit *ApparentEnergyUnits) ApparentEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ApparentEnergyVoltampereHour // Default value
		holdInUnit = &defaultUnit
	}

	return ApparentEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ApparentEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereHour by default.
func (a *ApparentEnergy) ToDtoJSON(holdInUnit *ApparentEnergyUnits) ([]byte, error) {
	// Convert to ApparentEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ApparentEnergy to a specific unit value.
// The function uses the provided unit type (ApparentEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ApparentEnergy) Convert(toUnit ApparentEnergyUnits) float64 {
	switch toUnit { 
    case ApparentEnergyVoltampereHour:
		return a.VoltampereHours()
    case ApparentEnergyKilovoltampereHour:
		return a.KilovoltampereHours()
    case ApparentEnergyMegavoltampereHour:
		return a.MegavoltampereHours()
	default:
		return math.NaN()
	}
}

func (a *ApparentEnergy) convertFromBase(toUnit ApparentEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ApparentEnergyVoltampereHour:
		return (value) 
	case ApparentEnergyKilovoltampereHour:
		return ((value) / 1000.0) 
	case ApparentEnergyMegavoltampereHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ApparentEnergy) convertToBase(value float64, fromUnit ApparentEnergyUnits) float64 {
	switch fromUnit { 
	case ApparentEnergyVoltampereHour:
		return (value) 
	case ApparentEnergyKilovoltampereHour:
		return ((value) * 1000.0) 
	case ApparentEnergyMegavoltampereHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ApparentEnergy in the default unit (VoltampereHour),
// formatted to two decimal places.
func (a ApparentEnergy) String() string {
	return a.ToString(ApparentEnergyVoltampereHour, 2)
}

// ToString formats the ApparentEnergy value as a string with the specified unit and fractional digits.
// It converts the ApparentEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ApparentEnergy value will be converted (e.g., VoltampereHour))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ApparentEnergy with the unit abbreviation.
func (a *ApparentEnergy) ToString(unit ApparentEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetApparentEnergyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetApparentEnergyAbbreviation(unit))
}

// Equals checks if the given ApparentEnergy is equal to the current ApparentEnergy.
//
// Parameters:
//    other: The ApparentEnergy to compare against.
//
// Returns:
//    bool: Returns true if both ApparentEnergy are equal, false otherwise.
func (a *ApparentEnergy) Equals(other *ApparentEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ApparentEnergy with another ApparentEnergy.
// It returns -1 if the current ApparentEnergy is less than the other ApparentEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ApparentEnergy to compare against.
//
// Returns:
//    int: -1 if the current ApparentEnergy is less, 1 if greater, and 0 if equal.
func (a *ApparentEnergy) CompareTo(other *ApparentEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ApparentEnergy to the current ApparentEnergy and returns the result.
// The result is a new ApparentEnergy instance with the sum of the values.
//
// Parameters:
//    other: The ApparentEnergy to add to the current ApparentEnergy.
//
// Returns:
//    *ApparentEnergy: A new ApparentEnergy instance representing the sum of both ApparentEnergy.
func (a *ApparentEnergy) Add(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ApparentEnergy from the current ApparentEnergy and returns the result.
// The result is a new ApparentEnergy instance with the difference of the values.
//
// Parameters:
//    other: The ApparentEnergy to subtract from the current ApparentEnergy.
//
// Returns:
//    *ApparentEnergy: A new ApparentEnergy instance representing the difference of both ApparentEnergy.
func (a *ApparentEnergy) Subtract(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ApparentEnergy by the given ApparentEnergy and returns the result.
// The result is a new ApparentEnergy instance with the product of the values.
//
// Parameters:
//    other: The ApparentEnergy to multiply with the current ApparentEnergy.
//
// Returns:
//    *ApparentEnergy: A new ApparentEnergy instance representing the product of both ApparentEnergy.
func (a *ApparentEnergy) Multiply(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current ApparentEnergy by the given ApparentEnergy and returns the result.
// The result is a new ApparentEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The ApparentEnergy to divide the current ApparentEnergy by.
//
// Returns:
//    *ApparentEnergy: A new ApparentEnergy instance representing the quotient of both ApparentEnergy.
func (a *ApparentEnergy) Divide(other *ApparentEnergy) *ApparentEnergy {
	return &ApparentEnergy{value: a.value / other.BaseValue()}
}

// GetApparentEnergyAbbreviation gets the unit abbreviation.
func GetApparentEnergyAbbreviation(unit ApparentEnergyUnits) string {
	switch unit { 
	case ApparentEnergyVoltampereHour:
		return "VAh" 
	case ApparentEnergyKilovoltampereHour:
		return "kVAh" 
	case ApparentEnergyMegavoltampereHour:
		return "MVAh" 
	default:
		return ""
	}
}