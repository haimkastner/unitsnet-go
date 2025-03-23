// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricApparentEnergyUnits defines various units of ElectricApparentEnergy.
type ElectricApparentEnergyUnits string

const (
    
        // 
        ElectricApparentEnergyVoltampereHour ElectricApparentEnergyUnits = "VoltampereHour"
        // 
        ElectricApparentEnergyKilovoltampereHour ElectricApparentEnergyUnits = "KilovoltampereHour"
        // 
        ElectricApparentEnergyMegavoltampereHour ElectricApparentEnergyUnits = "MegavoltampereHour"
)

var internalElectricApparentEnergyUnitsMap = map[ElectricApparentEnergyUnits]bool{
	
	ElectricApparentEnergyVoltampereHour: true,
	ElectricApparentEnergyKilovoltampereHour: true,
	ElectricApparentEnergyMegavoltampereHour: true,
}

// ElectricApparentEnergyDto represents a ElectricApparentEnergy measurement with a numerical value and its corresponding unit.
type ElectricApparentEnergyDto struct {
    // Value is the numerical representation of the ElectricApparentEnergy.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricApparentEnergy, as defined in the ElectricApparentEnergyUnits enumeration.
	Unit  ElectricApparentEnergyUnits `json:"unit" validate:"required,oneof=VoltampereHour KilovoltampereHour MegavoltampereHour"`
}

// ElectricApparentEnergyDtoFactory groups methods for creating and serializing ElectricApparentEnergyDto objects.
type ElectricApparentEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricApparentEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricApparentEnergyDtoFactory) FromJSON(data []byte) (*ElectricApparentEnergyDto, error) {
	a := ElectricApparentEnergyDto{}

    // Parse JSON into ElectricApparentEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricApparentEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricApparentEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricApparentEnergy represents a measurement in a of ElectricApparentEnergy.
//
// A unit for expressing the integral of apparent power over time, equal to the product of 1 volt-ampere and 1 hour, or to 3600 joules.
type ElectricApparentEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltampere_hoursLazy *float64 
    kilovoltampere_hoursLazy *float64 
    megavoltampere_hoursLazy *float64 
}

// ElectricApparentEnergyFactory groups methods for creating ElectricApparentEnergy instances.
type ElectricApparentEnergyFactory struct{}

// CreateElectricApparentEnergy creates a new ElectricApparentEnergy instance from the given value and unit.
func (uf ElectricApparentEnergyFactory) CreateElectricApparentEnergy(value float64, unit ElectricApparentEnergyUnits) (*ElectricApparentEnergy, error) {
	return newElectricApparentEnergy(value, unit)
}

// FromDto converts a ElectricApparentEnergyDto to a ElectricApparentEnergy instance.
func (uf ElectricApparentEnergyFactory) FromDto(dto ElectricApparentEnergyDto) (*ElectricApparentEnergy, error) {
	return newElectricApparentEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricApparentEnergy instance.
func (uf ElectricApparentEnergyFactory) FromDtoJSON(data []byte) (*ElectricApparentEnergy, error) {
	unitDto, err := ElectricApparentEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricApparentEnergyDto from JSON: %w", err)
	}
	return ElectricApparentEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereHours creates a new ElectricApparentEnergy instance from a value in VoltampereHours.
func (uf ElectricApparentEnergyFactory) FromVoltampereHours(value float64) (*ElectricApparentEnergy, error) {
	return newElectricApparentEnergy(value, ElectricApparentEnergyVoltampereHour)
}

// FromKilovoltampereHours creates a new ElectricApparentEnergy instance from a value in KilovoltampereHours.
func (uf ElectricApparentEnergyFactory) FromKilovoltampereHours(value float64) (*ElectricApparentEnergy, error) {
	return newElectricApparentEnergy(value, ElectricApparentEnergyKilovoltampereHour)
}

// FromMegavoltampereHours creates a new ElectricApparentEnergy instance from a value in MegavoltampereHours.
func (uf ElectricApparentEnergyFactory) FromMegavoltampereHours(value float64) (*ElectricApparentEnergy, error) {
	return newElectricApparentEnergy(value, ElectricApparentEnergyMegavoltampereHour)
}


// newElectricApparentEnergy creates a new ElectricApparentEnergy.
func newElectricApparentEnergy(value float64, fromUnit ElectricApparentEnergyUnits) (*ElectricApparentEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricApparentEnergyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricApparentEnergyUnits", fromUnit)
	}
	a := &ElectricApparentEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricApparentEnergy in VoltampereHour unit (the base/default quantity).
func (a *ElectricApparentEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereHours returns the ElectricApparentEnergy value in VoltampereHours.
//
// 
func (a *ElectricApparentEnergy) VoltampereHours() float64 {
	if a.voltampere_hoursLazy != nil {
		return *a.voltampere_hoursLazy
	}
	voltampere_hours := a.convertFromBase(ElectricApparentEnergyVoltampereHour)
	a.voltampere_hoursLazy = &voltampere_hours
	return voltampere_hours
}

// KilovoltampereHours returns the ElectricApparentEnergy value in KilovoltampereHours.
//
// 
func (a *ElectricApparentEnergy) KilovoltampereHours() float64 {
	if a.kilovoltampere_hoursLazy != nil {
		return *a.kilovoltampere_hoursLazy
	}
	kilovoltampere_hours := a.convertFromBase(ElectricApparentEnergyKilovoltampereHour)
	a.kilovoltampere_hoursLazy = &kilovoltampere_hours
	return kilovoltampere_hours
}

// MegavoltampereHours returns the ElectricApparentEnergy value in MegavoltampereHours.
//
// 
func (a *ElectricApparentEnergy) MegavoltampereHours() float64 {
	if a.megavoltampere_hoursLazy != nil {
		return *a.megavoltampere_hoursLazy
	}
	megavoltampere_hours := a.convertFromBase(ElectricApparentEnergyMegavoltampereHour)
	a.megavoltampere_hoursLazy = &megavoltampere_hours
	return megavoltampere_hours
}


// ToDto creates a ElectricApparentEnergyDto representation from the ElectricApparentEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereHour by default.
func (a *ElectricApparentEnergy) ToDto(holdInUnit *ElectricApparentEnergyUnits) ElectricApparentEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ElectricApparentEnergyVoltampereHour // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricApparentEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricApparentEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereHour by default.
func (a *ElectricApparentEnergy) ToDtoJSON(holdInUnit *ElectricApparentEnergyUnits) ([]byte, error) {
	// Convert to ElectricApparentEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricApparentEnergy to a specific unit value.
// The function uses the provided unit type (ElectricApparentEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricApparentEnergy) Convert(toUnit ElectricApparentEnergyUnits) float64 {
	switch toUnit { 
    case ElectricApparentEnergyVoltampereHour:
		return a.VoltampereHours()
    case ElectricApparentEnergyKilovoltampereHour:
		return a.KilovoltampereHours()
    case ElectricApparentEnergyMegavoltampereHour:
		return a.MegavoltampereHours()
	default:
		return math.NaN()
	}
}

func (a *ElectricApparentEnergy) convertFromBase(toUnit ElectricApparentEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricApparentEnergyVoltampereHour:
		return (value) 
	case ElectricApparentEnergyKilovoltampereHour:
		return ((value) / 1000.0) 
	case ElectricApparentEnergyMegavoltampereHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricApparentEnergy) convertToBase(value float64, fromUnit ElectricApparentEnergyUnits) float64 {
	switch fromUnit { 
	case ElectricApparentEnergyVoltampereHour:
		return (value) 
	case ElectricApparentEnergyKilovoltampereHour:
		return ((value) * 1000.0) 
	case ElectricApparentEnergyMegavoltampereHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricApparentEnergy in the default unit (VoltampereHour),
// formatted to two decimal places.
func (a ElectricApparentEnergy) String() string {
	return a.ToString(ElectricApparentEnergyVoltampereHour, 2)
}

// ToString formats the ElectricApparentEnergy value as a string with the specified unit and fractional digits.
// It converts the ElectricApparentEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricApparentEnergy value will be converted (e.g., VoltampereHour))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricApparentEnergy with the unit abbreviation.
func (a *ElectricApparentEnergy) ToString(unit ElectricApparentEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricApparentEnergyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricApparentEnergyAbbreviation(unit))
}

// Equals checks if the given ElectricApparentEnergy is equal to the current ElectricApparentEnergy.
//
// Parameters:
//    other: The ElectricApparentEnergy to compare against.
//
// Returns:
//    bool: Returns true if both ElectricApparentEnergy are equal, false otherwise.
func (a *ElectricApparentEnergy) Equals(other *ElectricApparentEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricApparentEnergy with another ElectricApparentEnergy.
// It returns -1 if the current ElectricApparentEnergy is less than the other ElectricApparentEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricApparentEnergy to compare against.
//
// Returns:
//    int: -1 if the current ElectricApparentEnergy is less, 1 if greater, and 0 if equal.
func (a *ElectricApparentEnergy) CompareTo(other *ElectricApparentEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricApparentEnergy to the current ElectricApparentEnergy and returns the result.
// The result is a new ElectricApparentEnergy instance with the sum of the values.
//
// Parameters:
//    other: The ElectricApparentEnergy to add to the current ElectricApparentEnergy.
//
// Returns:
//    *ElectricApparentEnergy: A new ElectricApparentEnergy instance representing the sum of both ElectricApparentEnergy.
func (a *ElectricApparentEnergy) Add(other *ElectricApparentEnergy) *ElectricApparentEnergy {
	return &ElectricApparentEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricApparentEnergy from the current ElectricApparentEnergy and returns the result.
// The result is a new ElectricApparentEnergy instance with the difference of the values.
//
// Parameters:
//    other: The ElectricApparentEnergy to subtract from the current ElectricApparentEnergy.
//
// Returns:
//    *ElectricApparentEnergy: A new ElectricApparentEnergy instance representing the difference of both ElectricApparentEnergy.
func (a *ElectricApparentEnergy) Subtract(other *ElectricApparentEnergy) *ElectricApparentEnergy {
	return &ElectricApparentEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricApparentEnergy by the given ElectricApparentEnergy and returns the result.
// The result is a new ElectricApparentEnergy instance with the product of the values.
//
// Parameters:
//    other: The ElectricApparentEnergy to multiply with the current ElectricApparentEnergy.
//
// Returns:
//    *ElectricApparentEnergy: A new ElectricApparentEnergy instance representing the product of both ElectricApparentEnergy.
func (a *ElectricApparentEnergy) Multiply(other *ElectricApparentEnergy) *ElectricApparentEnergy {
	return &ElectricApparentEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricApparentEnergy by the given ElectricApparentEnergy and returns the result.
// The result is a new ElectricApparentEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricApparentEnergy to divide the current ElectricApparentEnergy by.
//
// Returns:
//    *ElectricApparentEnergy: A new ElectricApparentEnergy instance representing the quotient of both ElectricApparentEnergy.
func (a *ElectricApparentEnergy) Divide(other *ElectricApparentEnergy) *ElectricApparentEnergy {
	return &ElectricApparentEnergy{value: a.value / other.BaseValue()}
}

// GetElectricApparentEnergyAbbreviation gets the unit abbreviation.
func GetElectricApparentEnergyAbbreviation(unit ElectricApparentEnergyUnits) string {
	switch unit { 
	case ElectricApparentEnergyVoltampereHour:
		return "VAh" 
	case ElectricApparentEnergyKilovoltampereHour:
		return "kVAh" 
	case ElectricApparentEnergyMegavoltampereHour:
		return "MVAh" 
	default:
		return ""
	}
}