// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ThermalResistanceUnits enumeration
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

// ThermalResistanceDto represents an ThermalResistance
type ThermalResistanceDto struct {
	Value float64
	Unit  ThermalResistanceUnits
}

// ThermalResistanceDtoFactory struct to group related functions
type ThermalResistanceDtoFactory struct{}

func (udf ThermalResistanceDtoFactory) FromJSON(data []byte) (*ThermalResistanceDto, error) {
	a := ThermalResistanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ThermalResistanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ThermalResistance struct
type ThermalResistance struct {
	value       float64
    
    square_meter_kelvins_per_kilowattLazy *float64 
    square_meter_kelvins_per_wattLazy *float64 
    square_meter_degrees_celsius_per_wattLazy *float64 
    square_centimeter_kelvins_per_wattLazy *float64 
    square_centimeter_hour_degrees_celsius_per_kilocalorieLazy *float64 
    hour_square_feet_degrees_fahrenheit_per_btuLazy *float64 
}

// ThermalResistanceFactory struct to group related functions
type ThermalResistanceFactory struct{}

func (uf ThermalResistanceFactory) CreateThermalResistance(value float64, unit ThermalResistanceUnits) (*ThermalResistance, error) {
	return newThermalResistance(value, unit)
}

func (uf ThermalResistanceFactory) FromDto(dto ThermalResistanceDto) (*ThermalResistance, error) {
	return newThermalResistance(dto.Value, dto.Unit)
}

func (uf ThermalResistanceFactory) FromDtoJSON(data []byte) (*ThermalResistance, error) {
	unitDto, err := ThermalResistanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ThermalResistanceFactory{}.FromDto(*unitDto)
}


// FromSquareMeterKelvinPerKilowatt creates a new ThermalResistance instance from SquareMeterKelvinPerKilowatt.
func (uf ThermalResistanceFactory) FromSquareMeterKelvinsPerKilowatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterKelvinPerKilowatt)
}

// FromSquareMeterKelvinPerWatt creates a new ThermalResistance instance from SquareMeterKelvinPerWatt.
func (uf ThermalResistanceFactory) FromSquareMeterKelvinsPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterKelvinPerWatt)
}

// FromSquareMeterDegreeCelsiusPerWatt creates a new ThermalResistance instance from SquareMeterDegreeCelsiusPerWatt.
func (uf ThermalResistanceFactory) FromSquareMeterDegreesCelsiusPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
}

// FromSquareCentimeterKelvinPerWatt creates a new ThermalResistance instance from SquareCentimeterKelvinPerWatt.
func (uf ThermalResistanceFactory) FromSquareCentimeterKelvinsPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareCentimeterKelvinPerWatt)
}

// FromSquareCentimeterHourDegreeCelsiusPerKilocalorie creates a new ThermalResistance instance from SquareCentimeterHourDegreeCelsiusPerKilocalorie.
func (uf ThermalResistanceFactory) FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
}

// FromHourSquareFeetDegreeFahrenheitPerBtu creates a new ThermalResistance instance from HourSquareFeetDegreeFahrenheitPerBtu.
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

// BaseValue returns the base value of ThermalResistance in SquareMeterKelvinPerKilowatt.
func (a *ThermalResistance) BaseValue() float64 {
	return a.value
}


// SquareMeterKelvinPerKilowatt returns the value in SquareMeterKelvinPerKilowatt.
func (a *ThermalResistance) SquareMeterKelvinsPerKilowatt() float64 {
	if a.square_meter_kelvins_per_kilowattLazy != nil {
		return *a.square_meter_kelvins_per_kilowattLazy
	}
	square_meter_kelvins_per_kilowatt := a.convertFromBase(ThermalResistanceSquareMeterKelvinPerKilowatt)
	a.square_meter_kelvins_per_kilowattLazy = &square_meter_kelvins_per_kilowatt
	return square_meter_kelvins_per_kilowatt
}

// SquareMeterKelvinPerWatt returns the value in SquareMeterKelvinPerWatt.
func (a *ThermalResistance) SquareMeterKelvinsPerWatt() float64 {
	if a.square_meter_kelvins_per_wattLazy != nil {
		return *a.square_meter_kelvins_per_wattLazy
	}
	square_meter_kelvins_per_watt := a.convertFromBase(ThermalResistanceSquareMeterKelvinPerWatt)
	a.square_meter_kelvins_per_wattLazy = &square_meter_kelvins_per_watt
	return square_meter_kelvins_per_watt
}

// SquareMeterDegreeCelsiusPerWatt returns the value in SquareMeterDegreeCelsiusPerWatt.
func (a *ThermalResistance) SquareMeterDegreesCelsiusPerWatt() float64 {
	if a.square_meter_degrees_celsius_per_wattLazy != nil {
		return *a.square_meter_degrees_celsius_per_wattLazy
	}
	square_meter_degrees_celsius_per_watt := a.convertFromBase(ThermalResistanceSquareMeterDegreeCelsiusPerWatt)
	a.square_meter_degrees_celsius_per_wattLazy = &square_meter_degrees_celsius_per_watt
	return square_meter_degrees_celsius_per_watt
}

// SquareCentimeterKelvinPerWatt returns the value in SquareCentimeterKelvinPerWatt.
func (a *ThermalResistance) SquareCentimeterKelvinsPerWatt() float64 {
	if a.square_centimeter_kelvins_per_wattLazy != nil {
		return *a.square_centimeter_kelvins_per_wattLazy
	}
	square_centimeter_kelvins_per_watt := a.convertFromBase(ThermalResistanceSquareCentimeterKelvinPerWatt)
	a.square_centimeter_kelvins_per_wattLazy = &square_centimeter_kelvins_per_watt
	return square_centimeter_kelvins_per_watt
}

// SquareCentimeterHourDegreeCelsiusPerKilocalorie returns the value in SquareCentimeterHourDegreeCelsiusPerKilocalorie.
func (a *ThermalResistance) SquareCentimeterHourDegreesCelsiusPerKilocalorie() float64 {
	if a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy != nil {
		return *a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy
	}
	square_centimeter_hour_degrees_celsius_per_kilocalorie := a.convertFromBase(ThermalResistanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
	a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy = &square_centimeter_hour_degrees_celsius_per_kilocalorie
	return square_centimeter_hour_degrees_celsius_per_kilocalorie
}

// HourSquareFeetDegreeFahrenheitPerBtu returns the value in HourSquareFeetDegreeFahrenheitPerBtu.
func (a *ThermalResistance) HourSquareFeetDegreesFahrenheitPerBtu() float64 {
	if a.hour_square_feet_degrees_fahrenheit_per_btuLazy != nil {
		return *a.hour_square_feet_degrees_fahrenheit_per_btuLazy
	}
	hour_square_feet_degrees_fahrenheit_per_btu := a.convertFromBase(ThermalResistanceHourSquareFeetDegreeFahrenheitPerBtu)
	a.hour_square_feet_degrees_fahrenheit_per_btuLazy = &hour_square_feet_degrees_fahrenheit_per_btu
	return hour_square_feet_degrees_fahrenheit_per_btu
}


// ToDto creates an ThermalResistanceDto representation.
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

// ToDtoJSON creates an ThermalResistanceDto representation.
func (a *ThermalResistance) ToDtoJSON(holdInUnit *ThermalResistanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ThermalResistance to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ThermalResistance) String() string {
	return a.ToString(ThermalResistanceSquareMeterKelvinPerKilowatt, 2)
}

// ToString formats the ThermalResistance to string.
// fractionalDigits -1 for not mention
func (a *ThermalResistance) ToString(unit ThermalResistanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ThermalResistance) getUnitAbbreviation(unit ThermalResistanceUnits) string {
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

// Check if the given ThermalResistance are equals to the current ThermalResistance
func (a *ThermalResistance) Equals(other *ThermalResistance) bool {
	return a.value == other.BaseValue()
}

// Check if the given ThermalResistance are equals to the current ThermalResistance
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

// Add the given ThermalResistance to the current ThermalResistance.
func (a *ThermalResistance) Add(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value + other.BaseValue()}
}

// Subtract the given ThermalResistance to the current ThermalResistance.
func (a *ThermalResistance) Subtract(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value - other.BaseValue()}
}

// Multiply the given ThermalResistance to the current ThermalResistance.
func (a *ThermalResistance) Multiply(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value * other.BaseValue()}
}

// Divide the given ThermalResistance to the current ThermalResistance.
func (a *ThermalResistance) Divide(other *ThermalResistance) *ThermalResistance {
	return &ThermalResistance{value: a.value / other.BaseValue()}
}