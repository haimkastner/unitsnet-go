// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ThermalInsulanceUnits defines various units of ThermalInsulance.
type ThermalInsulanceUnits string

const (
    
        // 
        ThermalInsulanceSquareMeterKelvinPerKilowatt ThermalInsulanceUnits = "SquareMeterKelvinPerKilowatt"
        // 
        ThermalInsulanceSquareMeterKelvinPerWatt ThermalInsulanceUnits = "SquareMeterKelvinPerWatt"
        // 
        ThermalInsulanceSquareMeterDegreeCelsiusPerWatt ThermalInsulanceUnits = "SquareMeterDegreeCelsiusPerWatt"
        // 
        ThermalInsulanceSquareCentimeterKelvinPerWatt ThermalInsulanceUnits = "SquareCentimeterKelvinPerWatt"
        // 
        ThermalInsulanceSquareMillimeterKelvinPerWatt ThermalInsulanceUnits = "SquareMillimeterKelvinPerWatt"
        // 
        ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie ThermalInsulanceUnits = "SquareCentimeterHourDegreeCelsiusPerKilocalorie"
        // 
        ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu ThermalInsulanceUnits = "HourSquareFeetDegreeFahrenheitPerBtu"
)

var internalThermalInsulanceUnitsMap = map[ThermalInsulanceUnits]bool{
	
	ThermalInsulanceSquareMeterKelvinPerKilowatt: true,
	ThermalInsulanceSquareMeterKelvinPerWatt: true,
	ThermalInsulanceSquareMeterDegreeCelsiusPerWatt: true,
	ThermalInsulanceSquareCentimeterKelvinPerWatt: true,
	ThermalInsulanceSquareMillimeterKelvinPerWatt: true,
	ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie: true,
	ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu: true,
}

// ThermalInsulanceDto represents a ThermalInsulance measurement with a numerical value and its corresponding unit.
type ThermalInsulanceDto struct {
    // Value is the numerical representation of the ThermalInsulance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ThermalInsulance, as defined in the ThermalInsulanceUnits enumeration.
	Unit  ThermalInsulanceUnits `json:"unit" validate:"required,oneof=SquareMeterKelvinPerKilowatt SquareMeterKelvinPerWatt SquareMeterDegreeCelsiusPerWatt SquareCentimeterKelvinPerWatt SquareMillimeterKelvinPerWatt SquareCentimeterHourDegreeCelsiusPerKilocalorie HourSquareFeetDegreeFahrenheitPerBtu"`
}

// ThermalInsulanceDtoFactory groups methods for creating and serializing ThermalInsulanceDto objects.
type ThermalInsulanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ThermalInsulanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ThermalInsulanceDtoFactory) FromJSON(data []byte) (*ThermalInsulanceDto, error) {
	a := ThermalInsulanceDto{}

    // Parse JSON into ThermalInsulanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ThermalInsulanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ThermalInsulanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ThermalInsulance represents a measurement in a of ThermalInsulance.
//
// Thermal insulance (R-value) is a measure of a material's resistance to the heat current. It quantifies how effectively a material can resist the transfer of heat through conduction, convection, and radiation. It has the units square metre kelvins per watt (m2⋅K/W) in SI units or square foot degree Fahrenheit–hours per British thermal unit (ft2⋅°F⋅h/Btu) in imperial units. The higher the thermal insulance, the better a material insulates against heat transfer. It is commonly used in construction to assess the insulation properties of materials such as walls, roofs, and insulation products.
type ThermalInsulance struct {
	// value is the base measurement stored internally.
	value       float64
    
    square_meter_kelvins_per_kilowattLazy *float64 
    square_meter_kelvins_per_wattLazy *float64 
    square_meter_degrees_celsius_per_wattLazy *float64 
    square_centimeter_kelvins_per_wattLazy *float64 
    square_millimeter_kelvins_per_wattLazy *float64 
    square_centimeter_hour_degrees_celsius_per_kilocalorieLazy *float64 
    hour_square_feet_degrees_fahrenheit_per_btuLazy *float64 
}

// ThermalInsulanceFactory groups methods for creating ThermalInsulance instances.
type ThermalInsulanceFactory struct{}

// CreateThermalInsulance creates a new ThermalInsulance instance from the given value and unit.
func (uf ThermalInsulanceFactory) CreateThermalInsulance(value float64, unit ThermalInsulanceUnits) (*ThermalInsulance, error) {
	return newThermalInsulance(value, unit)
}

// FromDto converts a ThermalInsulanceDto to a ThermalInsulance instance.
func (uf ThermalInsulanceFactory) FromDto(dto ThermalInsulanceDto) (*ThermalInsulance, error) {
	return newThermalInsulance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ThermalInsulance instance.
func (uf ThermalInsulanceFactory) FromDtoJSON(data []byte) (*ThermalInsulance, error) {
	unitDto, err := ThermalInsulanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ThermalInsulanceDto from JSON: %w", err)
	}
	return ThermalInsulanceFactory{}.FromDto(*unitDto)
}


// FromSquareMeterKelvinsPerKilowatt creates a new ThermalInsulance instance from a value in SquareMeterKelvinsPerKilowatt.
func (uf ThermalInsulanceFactory) FromSquareMeterKelvinsPerKilowatt(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareMeterKelvinPerKilowatt)
}

// FromSquareMeterKelvinsPerWatt creates a new ThermalInsulance instance from a value in SquareMeterKelvinsPerWatt.
func (uf ThermalInsulanceFactory) FromSquareMeterKelvinsPerWatt(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareMeterKelvinPerWatt)
}

// FromSquareMeterDegreesCelsiusPerWatt creates a new ThermalInsulance instance from a value in SquareMeterDegreesCelsiusPerWatt.
func (uf ThermalInsulanceFactory) FromSquareMeterDegreesCelsiusPerWatt(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
}

// FromSquareCentimeterKelvinsPerWatt creates a new ThermalInsulance instance from a value in SquareCentimeterKelvinsPerWatt.
func (uf ThermalInsulanceFactory) FromSquareCentimeterKelvinsPerWatt(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareCentimeterKelvinPerWatt)
}

// FromSquareMillimeterKelvinsPerWatt creates a new ThermalInsulance instance from a value in SquareMillimeterKelvinsPerWatt.
func (uf ThermalInsulanceFactory) FromSquareMillimeterKelvinsPerWatt(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareMillimeterKelvinPerWatt)
}

// FromSquareCentimeterHourDegreesCelsiusPerKilocalorie creates a new ThermalInsulance instance from a value in SquareCentimeterHourDegreesCelsiusPerKilocalorie.
func (uf ThermalInsulanceFactory) FromSquareCentimeterHourDegreesCelsiusPerKilocalorie(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
}

// FromHourSquareFeetDegreesFahrenheitPerBtu creates a new ThermalInsulance instance from a value in HourSquareFeetDegreesFahrenheitPerBtu.
func (uf ThermalInsulanceFactory) FromHourSquareFeetDegreesFahrenheitPerBtu(value float64) (*ThermalInsulance, error) {
	return newThermalInsulance(value, ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
}


// newThermalInsulance creates a new ThermalInsulance.
func newThermalInsulance(value float64, fromUnit ThermalInsulanceUnits) (*ThermalInsulance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalThermalInsulanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ThermalInsulanceUnits", fromUnit)
	}
	a := &ThermalInsulance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ThermalInsulance in SquareMeterKelvinPerKilowatt unit (the base/default quantity).
func (a *ThermalInsulance) BaseValue() float64 {
	return a.value
}


// SquareMeterKelvinsPerKilowatt returns the ThermalInsulance value in SquareMeterKelvinsPerKilowatt.
//
// 
func (a *ThermalInsulance) SquareMeterKelvinsPerKilowatt() float64 {
	if a.square_meter_kelvins_per_kilowattLazy != nil {
		return *a.square_meter_kelvins_per_kilowattLazy
	}
	square_meter_kelvins_per_kilowatt := a.convertFromBase(ThermalInsulanceSquareMeterKelvinPerKilowatt)
	a.square_meter_kelvins_per_kilowattLazy = &square_meter_kelvins_per_kilowatt
	return square_meter_kelvins_per_kilowatt
}

// SquareMeterKelvinsPerWatt returns the ThermalInsulance value in SquareMeterKelvinsPerWatt.
//
// 
func (a *ThermalInsulance) SquareMeterKelvinsPerWatt() float64 {
	if a.square_meter_kelvins_per_wattLazy != nil {
		return *a.square_meter_kelvins_per_wattLazy
	}
	square_meter_kelvins_per_watt := a.convertFromBase(ThermalInsulanceSquareMeterKelvinPerWatt)
	a.square_meter_kelvins_per_wattLazy = &square_meter_kelvins_per_watt
	return square_meter_kelvins_per_watt
}

// SquareMeterDegreesCelsiusPerWatt returns the ThermalInsulance value in SquareMeterDegreesCelsiusPerWatt.
//
// 
func (a *ThermalInsulance) SquareMeterDegreesCelsiusPerWatt() float64 {
	if a.square_meter_degrees_celsius_per_wattLazy != nil {
		return *a.square_meter_degrees_celsius_per_wattLazy
	}
	square_meter_degrees_celsius_per_watt := a.convertFromBase(ThermalInsulanceSquareMeterDegreeCelsiusPerWatt)
	a.square_meter_degrees_celsius_per_wattLazy = &square_meter_degrees_celsius_per_watt
	return square_meter_degrees_celsius_per_watt
}

// SquareCentimeterKelvinsPerWatt returns the ThermalInsulance value in SquareCentimeterKelvinsPerWatt.
//
// 
func (a *ThermalInsulance) SquareCentimeterKelvinsPerWatt() float64 {
	if a.square_centimeter_kelvins_per_wattLazy != nil {
		return *a.square_centimeter_kelvins_per_wattLazy
	}
	square_centimeter_kelvins_per_watt := a.convertFromBase(ThermalInsulanceSquareCentimeterKelvinPerWatt)
	a.square_centimeter_kelvins_per_wattLazy = &square_centimeter_kelvins_per_watt
	return square_centimeter_kelvins_per_watt
}

// SquareMillimeterKelvinsPerWatt returns the ThermalInsulance value in SquareMillimeterKelvinsPerWatt.
//
// 
func (a *ThermalInsulance) SquareMillimeterKelvinsPerWatt() float64 {
	if a.square_millimeter_kelvins_per_wattLazy != nil {
		return *a.square_millimeter_kelvins_per_wattLazy
	}
	square_millimeter_kelvins_per_watt := a.convertFromBase(ThermalInsulanceSquareMillimeterKelvinPerWatt)
	a.square_millimeter_kelvins_per_wattLazy = &square_millimeter_kelvins_per_watt
	return square_millimeter_kelvins_per_watt
}

// SquareCentimeterHourDegreesCelsiusPerKilocalorie returns the ThermalInsulance value in SquareCentimeterHourDegreesCelsiusPerKilocalorie.
//
// 
func (a *ThermalInsulance) SquareCentimeterHourDegreesCelsiusPerKilocalorie() float64 {
	if a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy != nil {
		return *a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy
	}
	square_centimeter_hour_degrees_celsius_per_kilocalorie := a.convertFromBase(ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie)
	a.square_centimeter_hour_degrees_celsius_per_kilocalorieLazy = &square_centimeter_hour_degrees_celsius_per_kilocalorie
	return square_centimeter_hour_degrees_celsius_per_kilocalorie
}

// HourSquareFeetDegreesFahrenheitPerBtu returns the ThermalInsulance value in HourSquareFeetDegreesFahrenheitPerBtu.
//
// 
func (a *ThermalInsulance) HourSquareFeetDegreesFahrenheitPerBtu() float64 {
	if a.hour_square_feet_degrees_fahrenheit_per_btuLazy != nil {
		return *a.hour_square_feet_degrees_fahrenheit_per_btuLazy
	}
	hour_square_feet_degrees_fahrenheit_per_btu := a.convertFromBase(ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu)
	a.hour_square_feet_degrees_fahrenheit_per_btuLazy = &hour_square_feet_degrees_fahrenheit_per_btu
	return hour_square_feet_degrees_fahrenheit_per_btu
}


// ToDto creates a ThermalInsulanceDto representation from the ThermalInsulance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterKelvinPerKilowatt by default.
func (a *ThermalInsulance) ToDto(holdInUnit *ThermalInsulanceUnits) ThermalInsulanceDto {
	if holdInUnit == nil {
		defaultUnit := ThermalInsulanceSquareMeterKelvinPerKilowatt // Default value
		holdInUnit = &defaultUnit
	}

	return ThermalInsulanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ThermalInsulance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterKelvinPerKilowatt by default.
func (a *ThermalInsulance) ToDtoJSON(holdInUnit *ThermalInsulanceUnits) ([]byte, error) {
	// Convert to ThermalInsulanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ThermalInsulance to a specific unit value.
// The function uses the provided unit type (ThermalInsulanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ThermalInsulance) Convert(toUnit ThermalInsulanceUnits) float64 {
	switch toUnit { 
    case ThermalInsulanceSquareMeterKelvinPerKilowatt:
		return a.SquareMeterKelvinsPerKilowatt()
    case ThermalInsulanceSquareMeterKelvinPerWatt:
		return a.SquareMeterKelvinsPerWatt()
    case ThermalInsulanceSquareMeterDegreeCelsiusPerWatt:
		return a.SquareMeterDegreesCelsiusPerWatt()
    case ThermalInsulanceSquareCentimeterKelvinPerWatt:
		return a.SquareCentimeterKelvinsPerWatt()
    case ThermalInsulanceSquareMillimeterKelvinPerWatt:
		return a.SquareMillimeterKelvinsPerWatt()
    case ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return a.SquareCentimeterHourDegreesCelsiusPerKilocalorie()
    case ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu:
		return a.HourSquareFeetDegreesFahrenheitPerBtu()
	default:
		return math.NaN()
	}
}

func (a *ThermalInsulance) convertFromBase(toUnit ThermalInsulanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ThermalInsulanceSquareMeterKelvinPerKilowatt:
		return (value) 
	case ThermalInsulanceSquareMeterKelvinPerWatt:
		return (value / 1000) 
	case ThermalInsulanceSquareMeterDegreeCelsiusPerWatt:
		return (value / 1000.0) 
	case ThermalInsulanceSquareCentimeterKelvinPerWatt:
		return (value / 0.1) 
	case ThermalInsulanceSquareMillimeterKelvinPerWatt:
		return (value / 0.001) 
	case ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return (value * 4.184 / (0.0001 * 3600)) 
	case ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu:
		return (value * (1055.05585262 * 1.8) / (1000 * 0.3048 * 0.3048 * 3600)) 
	default:
		return math.NaN()
	}
}

func (a *ThermalInsulance) convertToBase(value float64, fromUnit ThermalInsulanceUnits) float64 {
	switch fromUnit { 
	case ThermalInsulanceSquareMeterKelvinPerKilowatt:
		return (value) 
	case ThermalInsulanceSquareMeterKelvinPerWatt:
		return (value * 1000) 
	case ThermalInsulanceSquareMeterDegreeCelsiusPerWatt:
		return (value * 1000.0) 
	case ThermalInsulanceSquareCentimeterKelvinPerWatt:
		return (value * 0.1) 
	case ThermalInsulanceSquareMillimeterKelvinPerWatt:
		return (value * 0.001) 
	case ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return (value * (0.0001 * 3600) / 4.184) 
	case ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu:
		return (value * (1000 * 0.3048 * 0.3048 * 3600) / (1055.05585262 * 1.8)) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ThermalInsulance in the default unit (SquareMeterKelvinPerKilowatt),
// formatted to two decimal places.
func (a ThermalInsulance) String() string {
	return a.ToString(ThermalInsulanceSquareMeterKelvinPerKilowatt, 2)
}

// ToString formats the ThermalInsulance value as a string with the specified unit and fractional digits.
// It converts the ThermalInsulance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ThermalInsulance value will be converted (e.g., SquareMeterKelvinPerKilowatt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ThermalInsulance with the unit abbreviation.
func (a *ThermalInsulance) ToString(unit ThermalInsulanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetThermalInsulanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetThermalInsulanceAbbreviation(unit))
}

// Equals checks if the given ThermalInsulance is equal to the current ThermalInsulance.
//
// Parameters:
//    other: The ThermalInsulance to compare against.
//
// Returns:
//    bool: Returns true if both ThermalInsulance are equal, false otherwise.
func (a *ThermalInsulance) Equals(other *ThermalInsulance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ThermalInsulance with another ThermalInsulance.
// It returns -1 if the current ThermalInsulance is less than the other ThermalInsulance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ThermalInsulance to compare against.
//
// Returns:
//    int: -1 if the current ThermalInsulance is less, 1 if greater, and 0 if equal.
func (a *ThermalInsulance) CompareTo(other *ThermalInsulance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ThermalInsulance to the current ThermalInsulance and returns the result.
// The result is a new ThermalInsulance instance with the sum of the values.
//
// Parameters:
//    other: The ThermalInsulance to add to the current ThermalInsulance.
//
// Returns:
//    *ThermalInsulance: A new ThermalInsulance instance representing the sum of both ThermalInsulance.
func (a *ThermalInsulance) Add(other *ThermalInsulance) *ThermalInsulance {
	return &ThermalInsulance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ThermalInsulance from the current ThermalInsulance and returns the result.
// The result is a new ThermalInsulance instance with the difference of the values.
//
// Parameters:
//    other: The ThermalInsulance to subtract from the current ThermalInsulance.
//
// Returns:
//    *ThermalInsulance: A new ThermalInsulance instance representing the difference of both ThermalInsulance.
func (a *ThermalInsulance) Subtract(other *ThermalInsulance) *ThermalInsulance {
	return &ThermalInsulance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ThermalInsulance by the given ThermalInsulance and returns the result.
// The result is a new ThermalInsulance instance with the product of the values.
//
// Parameters:
//    other: The ThermalInsulance to multiply with the current ThermalInsulance.
//
// Returns:
//    *ThermalInsulance: A new ThermalInsulance instance representing the product of both ThermalInsulance.
func (a *ThermalInsulance) Multiply(other *ThermalInsulance) *ThermalInsulance {
	return &ThermalInsulance{value: a.value * other.BaseValue()}
}

// Divide divides the current ThermalInsulance by the given ThermalInsulance and returns the result.
// The result is a new ThermalInsulance instance with the quotient of the values.
//
// Parameters:
//    other: The ThermalInsulance to divide the current ThermalInsulance by.
//
// Returns:
//    *ThermalInsulance: A new ThermalInsulance instance representing the quotient of both ThermalInsulance.
func (a *ThermalInsulance) Divide(other *ThermalInsulance) *ThermalInsulance {
	return &ThermalInsulance{value: a.value / other.BaseValue()}
}

// GetThermalInsulanceAbbreviation gets the unit abbreviation.
func GetThermalInsulanceAbbreviation(unit ThermalInsulanceUnits) string {
	switch unit { 
	case ThermalInsulanceSquareMeterKelvinPerKilowatt:
		return "m²K/kW" 
	case ThermalInsulanceSquareMeterKelvinPerWatt:
		return "m²K/W" 
	case ThermalInsulanceSquareMeterDegreeCelsiusPerWatt:
		return "m²°C/W" 
	case ThermalInsulanceSquareCentimeterKelvinPerWatt:
		return "cm²K/W" 
	case ThermalInsulanceSquareMillimeterKelvinPerWatt:
		return "mm²K/W" 
	case ThermalInsulanceSquareCentimeterHourDegreeCelsiusPerKilocalorie:
		return "cm²Hr°C/kcal" 
	case ThermalInsulanceHourSquareFeetDegreeFahrenheitPerBtu:
		return "Hrft²°F/Btu" 
	default:
		return ""
	}
}