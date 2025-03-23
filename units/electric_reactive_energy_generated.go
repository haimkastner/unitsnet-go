// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricReactiveEnergyUnits defines various units of ElectricReactiveEnergy.
type ElectricReactiveEnergyUnits string

const (
    
        // 
        ElectricReactiveEnergyVoltampereReactiveHour ElectricReactiveEnergyUnits = "VoltampereReactiveHour"
        // 
        ElectricReactiveEnergyKilovoltampereReactiveHour ElectricReactiveEnergyUnits = "KilovoltampereReactiveHour"
        // 
        ElectricReactiveEnergyMegavoltampereReactiveHour ElectricReactiveEnergyUnits = "MegavoltampereReactiveHour"
)

var internalElectricReactiveEnergyUnitsMap = map[ElectricReactiveEnergyUnits]bool{
	
	ElectricReactiveEnergyVoltampereReactiveHour: true,
	ElectricReactiveEnergyKilovoltampereReactiveHour: true,
	ElectricReactiveEnergyMegavoltampereReactiveHour: true,
}

// ElectricReactiveEnergyDto represents a ElectricReactiveEnergy measurement with a numerical value and its corresponding unit.
type ElectricReactiveEnergyDto struct {
    // Value is the numerical representation of the ElectricReactiveEnergy.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricReactiveEnergy, as defined in the ElectricReactiveEnergyUnits enumeration.
	Unit  ElectricReactiveEnergyUnits `json:"unit" validate:"required,oneof=VoltampereReactiveHour KilovoltampereReactiveHour MegavoltampereReactiveHour"`
}

// ElectricReactiveEnergyDtoFactory groups methods for creating and serializing ElectricReactiveEnergyDto objects.
type ElectricReactiveEnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactiveEnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricReactiveEnergyDtoFactory) FromJSON(data []byte) (*ElectricReactiveEnergyDto, error) {
	a := ElectricReactiveEnergyDto{}

    // Parse JSON into ElectricReactiveEnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricReactiveEnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricReactiveEnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricReactiveEnergy represents a measurement in a of ElectricReactiveEnergy.
//
// The volt-ampere reactive hour (expressed as varh) is the reactive power of one Volt-ampere reactive produced in one hour.
type ElectricReactiveEnergy struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltampere_reactive_hoursLazy *float64 
    kilovoltampere_reactive_hoursLazy *float64 
    megavoltampere_reactive_hoursLazy *float64 
}

// ElectricReactiveEnergyFactory groups methods for creating ElectricReactiveEnergy instances.
type ElectricReactiveEnergyFactory struct{}

// CreateElectricReactiveEnergy creates a new ElectricReactiveEnergy instance from the given value and unit.
func (uf ElectricReactiveEnergyFactory) CreateElectricReactiveEnergy(value float64, unit ElectricReactiveEnergyUnits) (*ElectricReactiveEnergy, error) {
	return newElectricReactiveEnergy(value, unit)
}

// FromDto converts a ElectricReactiveEnergyDto to a ElectricReactiveEnergy instance.
func (uf ElectricReactiveEnergyFactory) FromDto(dto ElectricReactiveEnergyDto) (*ElectricReactiveEnergy, error) {
	return newElectricReactiveEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactiveEnergy instance.
func (uf ElectricReactiveEnergyFactory) FromDtoJSON(data []byte) (*ElectricReactiveEnergy, error) {
	unitDto, err := ElectricReactiveEnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricReactiveEnergyDto from JSON: %w", err)
	}
	return ElectricReactiveEnergyFactory{}.FromDto(*unitDto)
}


// FromVoltampereReactiveHours creates a new ElectricReactiveEnergy instance from a value in VoltampereReactiveHours.
func (uf ElectricReactiveEnergyFactory) FromVoltampereReactiveHours(value float64) (*ElectricReactiveEnergy, error) {
	return newElectricReactiveEnergy(value, ElectricReactiveEnergyVoltampereReactiveHour)
}

// FromKilovoltampereReactiveHours creates a new ElectricReactiveEnergy instance from a value in KilovoltampereReactiveHours.
func (uf ElectricReactiveEnergyFactory) FromKilovoltampereReactiveHours(value float64) (*ElectricReactiveEnergy, error) {
	return newElectricReactiveEnergy(value, ElectricReactiveEnergyKilovoltampereReactiveHour)
}

// FromMegavoltampereReactiveHours creates a new ElectricReactiveEnergy instance from a value in MegavoltampereReactiveHours.
func (uf ElectricReactiveEnergyFactory) FromMegavoltampereReactiveHours(value float64) (*ElectricReactiveEnergy, error) {
	return newElectricReactiveEnergy(value, ElectricReactiveEnergyMegavoltampereReactiveHour)
}


// newElectricReactiveEnergy creates a new ElectricReactiveEnergy.
func newElectricReactiveEnergy(value float64, fromUnit ElectricReactiveEnergyUnits) (*ElectricReactiveEnergy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricReactiveEnergyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricReactiveEnergyUnits", fromUnit)
	}
	a := &ElectricReactiveEnergy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricReactiveEnergy in VoltampereReactiveHour unit (the base/default quantity).
func (a *ElectricReactiveEnergy) BaseValue() float64 {
	return a.value
}


// VoltampereReactiveHours returns the ElectricReactiveEnergy value in VoltampereReactiveHours.
//
// 
func (a *ElectricReactiveEnergy) VoltampereReactiveHours() float64 {
	if a.voltampere_reactive_hoursLazy != nil {
		return *a.voltampere_reactive_hoursLazy
	}
	voltampere_reactive_hours := a.convertFromBase(ElectricReactiveEnergyVoltampereReactiveHour)
	a.voltampere_reactive_hoursLazy = &voltampere_reactive_hours
	return voltampere_reactive_hours
}

// KilovoltampereReactiveHours returns the ElectricReactiveEnergy value in KilovoltampereReactiveHours.
//
// 
func (a *ElectricReactiveEnergy) KilovoltampereReactiveHours() float64 {
	if a.kilovoltampere_reactive_hoursLazy != nil {
		return *a.kilovoltampere_reactive_hoursLazy
	}
	kilovoltampere_reactive_hours := a.convertFromBase(ElectricReactiveEnergyKilovoltampereReactiveHour)
	a.kilovoltampere_reactive_hoursLazy = &kilovoltampere_reactive_hours
	return kilovoltampere_reactive_hours
}

// MegavoltampereReactiveHours returns the ElectricReactiveEnergy value in MegavoltampereReactiveHours.
//
// 
func (a *ElectricReactiveEnergy) MegavoltampereReactiveHours() float64 {
	if a.megavoltampere_reactive_hoursLazy != nil {
		return *a.megavoltampere_reactive_hoursLazy
	}
	megavoltampere_reactive_hours := a.convertFromBase(ElectricReactiveEnergyMegavoltampereReactiveHour)
	a.megavoltampere_reactive_hoursLazy = &megavoltampere_reactive_hours
	return megavoltampere_reactive_hours
}


// ToDto creates a ElectricReactiveEnergyDto representation from the ElectricReactiveEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactiveHour by default.
func (a *ElectricReactiveEnergy) ToDto(holdInUnit *ElectricReactiveEnergyUnits) ElectricReactiveEnergyDto {
	if holdInUnit == nil {
		defaultUnit := ElectricReactiveEnergyVoltampereReactiveHour // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricReactiveEnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricReactiveEnergy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactiveHour by default.
func (a *ElectricReactiveEnergy) ToDtoJSON(holdInUnit *ElectricReactiveEnergyUnits) ([]byte, error) {
	// Convert to ElectricReactiveEnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricReactiveEnergy to a specific unit value.
// The function uses the provided unit type (ElectricReactiveEnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricReactiveEnergy) Convert(toUnit ElectricReactiveEnergyUnits) float64 {
	switch toUnit { 
    case ElectricReactiveEnergyVoltampereReactiveHour:
		return a.VoltampereReactiveHours()
    case ElectricReactiveEnergyKilovoltampereReactiveHour:
		return a.KilovoltampereReactiveHours()
    case ElectricReactiveEnergyMegavoltampereReactiveHour:
		return a.MegavoltampereReactiveHours()
	default:
		return math.NaN()
	}
}

func (a *ElectricReactiveEnergy) convertFromBase(toUnit ElectricReactiveEnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ElectricReactiveEnergyKilovoltampereReactiveHour:
		return ((value) / 1000.0) 
	case ElectricReactiveEnergyMegavoltampereReactiveHour:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricReactiveEnergy) convertToBase(value float64, fromUnit ElectricReactiveEnergyUnits) float64 {
	switch fromUnit { 
	case ElectricReactiveEnergyVoltampereReactiveHour:
		return (value) 
	case ElectricReactiveEnergyKilovoltampereReactiveHour:
		return ((value) * 1000.0) 
	case ElectricReactiveEnergyMegavoltampereReactiveHour:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricReactiveEnergy in the default unit (VoltampereReactiveHour),
// formatted to two decimal places.
func (a ElectricReactiveEnergy) String() string {
	return a.ToString(ElectricReactiveEnergyVoltampereReactiveHour, 2)
}

// ToString formats the ElectricReactiveEnergy value as a string with the specified unit and fractional digits.
// It converts the ElectricReactiveEnergy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricReactiveEnergy value will be converted (e.g., VoltampereReactiveHour))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricReactiveEnergy with the unit abbreviation.
func (a *ElectricReactiveEnergy) ToString(unit ElectricReactiveEnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricReactiveEnergyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricReactiveEnergyAbbreviation(unit))
}

// Equals checks if the given ElectricReactiveEnergy is equal to the current ElectricReactiveEnergy.
//
// Parameters:
//    other: The ElectricReactiveEnergy to compare against.
//
// Returns:
//    bool: Returns true if both ElectricReactiveEnergy are equal, false otherwise.
func (a *ElectricReactiveEnergy) Equals(other *ElectricReactiveEnergy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricReactiveEnergy with another ElectricReactiveEnergy.
// It returns -1 if the current ElectricReactiveEnergy is less than the other ElectricReactiveEnergy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricReactiveEnergy to compare against.
//
// Returns:
//    int: -1 if the current ElectricReactiveEnergy is less, 1 if greater, and 0 if equal.
func (a *ElectricReactiveEnergy) CompareTo(other *ElectricReactiveEnergy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricReactiveEnergy to the current ElectricReactiveEnergy and returns the result.
// The result is a new ElectricReactiveEnergy instance with the sum of the values.
//
// Parameters:
//    other: The ElectricReactiveEnergy to add to the current ElectricReactiveEnergy.
//
// Returns:
//    *ElectricReactiveEnergy: A new ElectricReactiveEnergy instance representing the sum of both ElectricReactiveEnergy.
func (a *ElectricReactiveEnergy) Add(other *ElectricReactiveEnergy) *ElectricReactiveEnergy {
	return &ElectricReactiveEnergy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricReactiveEnergy from the current ElectricReactiveEnergy and returns the result.
// The result is a new ElectricReactiveEnergy instance with the difference of the values.
//
// Parameters:
//    other: The ElectricReactiveEnergy to subtract from the current ElectricReactiveEnergy.
//
// Returns:
//    *ElectricReactiveEnergy: A new ElectricReactiveEnergy instance representing the difference of both ElectricReactiveEnergy.
func (a *ElectricReactiveEnergy) Subtract(other *ElectricReactiveEnergy) *ElectricReactiveEnergy {
	return &ElectricReactiveEnergy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricReactiveEnergy by the given ElectricReactiveEnergy and returns the result.
// The result is a new ElectricReactiveEnergy instance with the product of the values.
//
// Parameters:
//    other: The ElectricReactiveEnergy to multiply with the current ElectricReactiveEnergy.
//
// Returns:
//    *ElectricReactiveEnergy: A new ElectricReactiveEnergy instance representing the product of both ElectricReactiveEnergy.
func (a *ElectricReactiveEnergy) Multiply(other *ElectricReactiveEnergy) *ElectricReactiveEnergy {
	return &ElectricReactiveEnergy{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricReactiveEnergy by the given ElectricReactiveEnergy and returns the result.
// The result is a new ElectricReactiveEnergy instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricReactiveEnergy to divide the current ElectricReactiveEnergy by.
//
// Returns:
//    *ElectricReactiveEnergy: A new ElectricReactiveEnergy instance representing the quotient of both ElectricReactiveEnergy.
func (a *ElectricReactiveEnergy) Divide(other *ElectricReactiveEnergy) *ElectricReactiveEnergy {
	return &ElectricReactiveEnergy{value: a.value / other.BaseValue()}
}

// GetElectricReactiveEnergyAbbreviation gets the unit abbreviation.
func GetElectricReactiveEnergyAbbreviation(unit ElectricReactiveEnergyUnits) string {
	switch unit { 
	case ElectricReactiveEnergyVoltampereReactiveHour:
		return "varh" 
	case ElectricReactiveEnergyKilovoltampereReactiveHour:
		return "kvarh" 
	case ElectricReactiveEnergyMegavoltampereReactiveHour:
		return "Mvarh" 
	default:
		return ""
	}
}