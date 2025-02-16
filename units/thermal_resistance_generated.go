// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ThermalResistanceUnits defines various units of ThermalResistance.
type ThermalResistanceUnits string

const (
    
        // 
        ThermalResistanceSquareMeterKelvinPerKilowatt ThermalResistanceUnits = "SquareMeterKelvinPerKilowatt"
        // 
        ThermalResistanceSquareMeterKelvinPerWatt ThermalResistanceUnits = "SquareMeterKelvinPerWatt"
        // 
        ThermalResistanceSquareMeterDegreeCelsiusPerWatt ThermalResistanceUnits = "SquareMeterDegreeCelsiusPerWatt"
        // 
        ThermalResistanceSquareCentimeterKelvinPerWatt ThermalResistanceUnits = "SquareCentimeterKelvinPerWatt"
        // 
        ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie ThermalResistanceUnits = "SquareCentimeterHourDegreeCelsiusPerKilocalorie"
        // 
        ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu ThermalResistanceUnits = "HourSquareFeetDegreeFahrenheitPerBtu"
)

// ThermalResistanceDto represents a ThermalResistance measurement with a numerical value and its corresponding unit.
type ThermalResistanceDto struct {
    // Value is the numerical representation of the ThermalResistance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ThermalResistance, as defined in the ThermalResistanceUnits enumeration.
	Unit  ThermalResistanceUnits `json:"unit"`
}

// ThermalResistanceDtoFactory groups methods for creating and serializing ThermalResistanceDto objects.
type ThermalResistanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ThermalResistanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ThermalResistanceDtoFactory) FromJSON(data []byte) (*ThermalResistanceDto, error) {
	a := ThermalResistanceDto{}

    // Parse JSON into ThermalResistanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ThermalResistanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ThermalResistanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ThermalResistance represents a measurement in a of ThermalResistance.
//
// Heat Transfer Coefficient or Thermal conductivity - indicates a materials ability to conduct heat.
type ThermalResistance struct {
	// value is the base measurement stored internally.
	value       float64
    
    square_meter_kelvins_per_kilowattLazy *float64 
    square_meter_kelvins_per_wattLazy *float64 
    square_meter_degrees_celsius_per_wattLazy *float64 
    square_centimeter_kelvins_per_wattLazy *float64 
    square_centimeter_hour_degrees_celsius_per_kilocalorieLazy *float64 
    hour_square_feet_degrees_fahrenheit_per_btuLazy *float64 
}

// ThermalResistanceFactory groups methods for creating ThermalResistance instances.
type ThermalResistanceFactory struct{}

// CreateThermalResistance creates a new ThermalResistance instance from the given value and unit.
func (uf ThermalResistanceFactory) CreateThermalResistance(value float64, unit ThermalResistanceUnits) (*ThermalResistance, error) {
	return newThermalResistance(value, unit)
}

// FromDto converts a ThermalResistanceDto to a ThermalResistance instance.
func (uf ThermalResistanceFactory) FromDto(dto ThermalResistanceDto) (*ThermalResistance, error) {
	return newThermalResistance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ThermalResistance instance.
func (uf ThermalResistanceFactory) FromDtoJSON(data []byte) (*ThermalResistance, error) {
	unitDto, err := ThermalResistanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ThermalResistanceDto from JSON: %w", err)
	}
	return ThermalResistanceFactory{}.FromDto(*unitDto)
}


// FromSquareMeterKelvinsPerKilowatt creates a new ThermalResistance instance from a value in SquareMeterKelvinsPerKilowatt.
func (uf ThermalResistanceFactory) FromSquareMeterKelvinsPerKilowatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterKelvinPerKilowatt)
}

// FromSquareMeterKelvinsPerWatt creates a new ThermalResistance instance from a value in SquareMeterKelvinsPerWatt.
func (uf ThermalResistanceFactory) FromSquareMeterKelvinsPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterKelvinPerWatt)
}

// FromSquareMeterDegreesCelsiusPerWatt creates a new ThermalResistance instance from a value in SquareMeterDegreesCelsiusPerWatt.
func (uf ThermalResistanceFactory) FromSquareMeterDegreesCelsiusPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
}

// FromSquareCentimeterKelvinsPerWatt creates a new ThermalResistance instance from a value in SquareCentimeterKelvinsPerWatt.
func (uf ThermalResistanceFactory) FromSquareCentimeterKelvinsPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareCentimeterKelvinPerWatt)
}

// FromSquareCentimeterHourDegreesCelsiusPerKilocalorie creates a new ThermalResistance instance from a value in SquareCentimeterHourDegreesCelsiusPerKilocalorie.
func (uf ThermalResistanceFactory) FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
}

// FromHourSquareFeetDegreesFahrenheitPerBtu creates a new ThermalResistance instance from a value in HourSquareFeetDegreesFahrenheitPerBtu.
func (uf ThermalResistanceFactory) FromHourSquareFeetDegreesFahrenheitPerBtu(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
}


// newThermalResistance creates a new ThermalResistance.
func newThermalResistance(value float64, fromUnit ThermalResistanceUnits) (*ThermalResistance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ThermalResistance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ThermalResistance in SquareMeterKelvinPerKilowatt unit (the base/default quantity).
func (a *ThermalResistance) BaseValue() float64 {
	return a.value
}


// SquareMeterKelvinsPerKilowatt returns the ThermalResistance value in SquareMeterKelvinsPerKilowatt.
//
// 
func (a *ThermalResistance) SquareMeterKelvinsPerKilowatt() float64 {
	if a.square_meter_kelvins_per_kilowattLazy != nil {
		return *a.square_meter_kelvins_per_kilowattLazy
	}
	square_meter_kelvins_per_kilowatt := a.convertFromBase(ThermalResistanceSquareMeterKelvinPerKilowatt)
	a.square_meter_kelvins_per_kilowattLazy = &square_meter_kelvins_per_kilowatt
	return square_meter_kelvins_per_kilowatt
}

// SquareMeterKelvinsPerWatt returns the ThermalResistance value in SquareMeterKelvinsPerWatt.
//
// 
func (a *ThermalResistance) SquareMeterKelvinsPerWatt() float64 {
	if a.square_meter_kelvins_per_wattLazy != nil {
		return *a.square_meter_kelvins_per_wattLazy
	}
	square_meter_kelvins_per_watt := a.convertFromBase(ThermalResistanceSquareMeterKelvinPerWatt)
	a.square_meter_kelvins_per_wattLazy = &square_meter_kelvins_per_watt
	return square_meter_kelvins_per_watt
}

// SquareMeterDegreesCelsiusPerWatt returns the ThermalResistance value in SquareMeterDegreesCelsiusPerWatt.
//
// 
func (a *ThermalResistance) SquareMeterDegreesCelsiusPerWatt() float64 {
	if a.square_meter_degrees_celsius_per_wattLazy != nil {
		return *a.square_meter_degrees_celsius_per_wattLazy
	}
	square_meter_degrees_celsius_per_watt := a.convertFromBase(ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
	a.square_meter_degrees_celsius_per_wattLazy = &square_meter_degrees_celsius_per_watt
	return square_meter_degrees_celsius_per_watt
}

// SquareCentimeterKelvinsPerWatt returns the ThermalResistance value in SquareCentimeterKelvinsPerWatt.
//
// 
func (a *ThermalResistance) SquareCentimeterKelvinsPerWatt() float64 {
	if a.square_centimeter_kelvins_per_wattLazy != nil {
		return *a.square_centimeter_kelvins_per_wattLazy
	}
	square_centimeter_kelvins_per_watt := a.convertFromBase(ThermalResistanceSquareCentimeterKelvinPerWatt)
	a.square_centimeter_kelvins_per_wattLazy = &square_centimeter_kelvins_per_watt
	return square_centimeter_kelvins_per_watt
}

// SquareCentimeterHourDegreesCelsiusPerKilocalorie returns the ThermalResistance value in SquareCentimeterHourDegreesCelsiusPerKilocalorie.
//
// 
func (a *ThermalResistance) SquareCentimeterHourDegreesCelsiusPerKilocalorie() float64 {
	if a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy != nil {
		return *a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy
	}
	square_centimeter_hour_degrees_celsius_per_kilocalorie := a.convertFromBase(ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
	a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy = &square_centimeter_hour_degrees_celsius_per_kilocalorie
	return square_centimeter_hour_degrees_celsius_per_kilocalorie
}

// HourSquareFeetDegreesFahrenheitPerBtu returns the ThermalResistance value in HourSquareFeetDegreesFahrenheitPerBtu.
//
// 
func (a *ThermalResistance) HourSquareFeetDegreesFahrenheitPerBtu() float64 {
	if a.hour_square_feet_degrees_fahrenheit_per_btuLazy != nil {
		return *a.hour_square_feet_degrees_fahrenheit_per_btuLazy
	}
	hour_square_feet_degrees_fahrenheit_per_btu := a.convertFromBase(ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
	a.hour_square_feet_degrees_fahrenheit_per_btuLazy = &hour_square_feet_degrees_fahrenheit_per_btu
	return hour_square_feet_degrees_fahrenheit_per_btu
}


// ToDto creates a ThermalResistanceDto representation from the ThermalResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterKelvinPerKilowatt by default.
func (a *ThermalResistance) ToDto(holdInUnit *ThermalResistanceUnits) ThermalResistanceDto {
	if holdInUnit == nil {
		defaultUnit := ThermalResistanceSquareMeterKelvinPerKilowatt // Default value
		holdInUnit = &defaultUnit
	}

	return ThermalResistanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ThermalResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterKelvinPerKilowatt by default.
func (a *ThermalResistance) ToDtoJSON(holdInUnit *ThermalResistanceUnits) ([]byte, error) {
	// Convert to ThermalResistanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ThermalResistance to a specific unit value.
// The function uses the provided unit type (ThermalResistanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ThermalResistance) Convert(toUnit ThermalResistanceUnits) float64 {
	switch toUnit { 
    case ThermalResistanceSquareMeterKelvinPerKilowatt:
		return a.SquareMeterKelvinsPerKilowatt()
    case ThermalResistanceSquareMeterKelvinPerWatt:
		return a.SquareMeterKelvinsPerWatt()
    case ThermalResistanceSquareMeterDegreeCelsiusPerWatt:
		return a.SquareMeterDegreesCelsiusPerWatt()
    case ThermalResistanceSquareCentimeterKelvinPerWatt:
		return a.SquareCentimeterKelvinsPerWatt()
    case ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
    case ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu:
		return a.HourSquareFeetDegreesFahrenheitPerBtu()
	default:
		return math.NaN()
	}
}

func (a *ThermalResistance) convertFromBase(toUnit ThermalResistanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ThermalResistanceSquareMeterKelvinPerKilowatt:
		return (value) 
	case ThermalResistanceSquareMeterKelvinPerWatt:
		return (value / 1000) 
	case ThermalResistanceSquareMeterDegreeCelsiusPerWatt:
		return (value / 1000.0) 
	case ThermalResistanceSquareCentimeterKelvinPerWatt:
		return (value / 0.1) 
	case ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return (value / 0.0859779507590433) 
	case ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu:
		return (value / 176.1121482159839) 
	default:
		return math.NaN()
	}
}

func (a *ThermalResistance) convertToBase(value float64, fromUnit ThermalResistanceUnits) float64 {
	switch fromUnit { 
	case ThermalResistanceSquareMeterKelvinPerKilowatt:
		return (value) 
	case ThermalResistanceSquareMeterKelvinPerWatt:
		return (value * 1000) 
	case ThermalResistanceSquareMeterDegreeCelsiusPerWatt:
		return (value * 1000.0) 
	case ThermalResistanceSquareCentimeterKelvinPerWatt:
		return (value * 0.1) 
	case ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return (value * 0.0859779507590433) 
	case ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu:
		return (value * 176.1121482159839) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ThermalResistance in the default unit (SquareMeterKelvinPerKilowatt),
// formatted to two decimal places.
func (a ThermalResistance) String() string {
	return a.ToString(ThermalResistanceSquareMeterKelvinPerKilowatt, 2)
}

// ToString formats the ThermalResistance value as a string with the specified unit and fractional digits.
// It converts the ThermalResistance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ThermalResistance value will be converted (e.g., SquareMeterKelvinPerKilowatt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ThermalResistance with the unit abbreviation.
func (a *ThermalResistance) ToString(unit ThermalResistanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetThermalResistanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetThermalResistanceAbbreviation(unit))
}

// Equals checks if the given ThermalResistance is equal to the current ThermalResistance.
//
// Parameters:
//    other: The ThermalResistance to compare against.
//
// Returns:
//    bool: Returns true if both ThermalResistance are equal, false otherwise.
func (a *ThermalResistance) Equals(other *ThermalResistance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ThermalResistance with another ThermalResistance.
// It returns -1 if the current ThermalResistance is less than the other ThermalResistance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ThermalResistance to compare against.
//
// Returns:
//    int: -1 if the current ThermalResistance is less, 1 if greater, and 0 if equal.
func (a *ThermalResistance) CompareTo(other *ThermalResistance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ThermalResistance to the current ThermalResistance and returns the result.
// The result is a new ThermalResistance instance with the sum of the values.
//
// Parameters:
//    other: The ThermalResistance to add to the current ThermalResistance.
//
// Returns:
//    *ThermalResistance: A new ThermalResistance instance representing the sum of both ThermalResistance.
func (a *ThermalResistance) Add(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ThermalResistance from the current ThermalResistance and returns the result.
// The result is a new ThermalResistance instance with the difference of the values.
//
// Parameters:
//    other: The ThermalResistance to subtract from the current ThermalResistance.
//
// Returns:
//    *ThermalResistance: A new ThermalResistance instance representing the difference of both ThermalResistance.
func (a *ThermalResistance) Subtract(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ThermalResistance by the given ThermalResistance and returns the result.
// The result is a new ThermalResistance instance with the product of the values.
//
// Parameters:
//    other: The ThermalResistance to multiply with the current ThermalResistance.
//
// Returns:
//    *ThermalResistance: A new ThermalResistance instance representing the product of both ThermalResistance.
func (a *ThermalResistance) Multiply(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value * other.BaseValue()}
}

// Divide divides the current ThermalResistance by the given ThermalResistance and returns the result.
// The result is a new ThermalResistance instance with the quotient of the values.
//
// Parameters:
//    other: The ThermalResistance to divide the current ThermalResistance by.
//
// Returns:
//    *ThermalResistance: A new ThermalResistance instance representing the quotient of both ThermalResistance.
func (a *ThermalResistance) Divide(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value / other.BaseValue()}
}

// GetThermalResistanceAbbreviation gets the unit abbreviation.
func GetThermalResistanceAbbreviation(unit ThermalResistanceUnits) string {
	switch unit { 
	case ThermalResistanceSquareMeterKelvinPerKilowatt:
		return "m²K/kW" 
	case ThermalResistanceSquareMeterKelvinPerWatt:
		return "m²K/W" 
	case ThermalResistanceSquareMeterDegreeCelsiusPerWatt:
		return "m²°C/W" 
	case ThermalResistanceSquareCentimeterKelvinPerWatt:
		return "cm²K/W" 
	case ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return "cm²Hr°C/kcal" 
	case ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu:
		return "Hrft²°F/Btu" 
	default:
		return ""
	}
}