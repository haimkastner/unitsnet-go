// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ThermalConductivityUnits defines various units of ThermalConductivity.
type ThermalConductivityUnits string

const (
    
        // 
        ThermalConductivityWattPerMeterKelvin ThermalConductivityUnits = "WattPerMeterKelvin"
        // 
        ThermalConductivityBtuPerHourFootFahrenheit ThermalConductivityUnits = "BtuPerHourFootFahrenheit"
)

// ThermalConductivityDto represents a ThermalConductivity measurement with a numerical value and its corresponding unit.
type ThermalConductivityDto struct {
    // Value is the numerical representation of the ThermalConductivity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ThermalConductivity, as defined in the ThermalConductivityUnits enumeration.
	Unit  ThermalConductivityUnits `json:"unit"`
}

// ThermalConductivityDtoFactory groups methods for creating and serializing ThermalConductivityDto objects.
type ThermalConductivityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ThermalConductivityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ThermalConductivityDtoFactory) FromJSON(data []byte) (*ThermalConductivityDto, error) {
	a := ThermalConductivityDto{}

    // Parse JSON into ThermalConductivityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ThermalConductivityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ThermalConductivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ThermalConductivity represents a measurement in a of ThermalConductivity.
//
// Thermal conductivity is the property of a material to conduct heat.
type ThermalConductivity struct {
	// value is the base measurement stored internally.
	value       float64
    
    watts_per_meter_kelvinLazy *float64 
    btus_per_hour_foot_fahrenheitLazy *float64 
}

// ThermalConductivityFactory groups methods for creating ThermalConductivity instances.
type ThermalConductivityFactory struct{}

// CreateThermalConductivity creates a new ThermalConductivity instance from the given value and unit.
func (uf ThermalConductivityFactory) CreateThermalConductivity(value float64, unit ThermalConductivityUnits) (*ThermalConductivity, error) {
	return newThermalConductivity(value, unit)
}

// FromDto converts a ThermalConductivityDto to a ThermalConductivity instance.
func (uf ThermalConductivityFactory) FromDto(dto ThermalConductivityDto) (*ThermalConductivity, error) {
	return newThermalConductivity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ThermalConductivity instance.
func (uf ThermalConductivityFactory) FromDtoJSON(data []byte) (*ThermalConductivity, error) {
	unitDto, err := ThermalConductivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ThermalConductivityDto from JSON: %w", err)
	}
	return ThermalConductivityFactory{}.FromDto(*unitDto)
}


// FromWattsPerMeterKelvin creates a new ThermalConductivity instance from a value in WattsPerMeterKelvin.
func (uf ThermalConductivityFactory) FromWattsPerMeterKelvin(value float64) (*ThermalConductivity, error) {
	return newThermalConductivity(value, ThermalConductivityWattPerMeterKelvin)
}

// FromBtusPerHourFootFahrenheit creates a new ThermalConductivity instance from a value in BtusPerHourFootFahrenheit.
func (uf ThermalConductivityFactory) FromBtusPerHourFootFahrenheit(value float64) (*ThermalConductivity, error) {
	return newThermalConductivity(value, ThermalConductivityBtuPerHourFootFahrenheit)
}


// newThermalConductivity creates a new ThermalConductivity.
func newThermalConductivity(value float64, fromUnit ThermalConductivityUnits) (*ThermalConductivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ThermalConductivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ThermalConductivity in WattPerMeterKelvin unit (the base/default quantity).
func (a *ThermalConductivity) BaseValue() float64 {
	return a.value
}


// WattsPerMeterKelvin returns the ThermalConductivity value in WattsPerMeterKelvin.
//
// 
func (a *ThermalConductivity) WattsPerMeterKelvin() float64 {
	if a.watts_per_meter_kelvinLazy != nil {
		return *a.watts_per_meter_kelvinLazy
	}
	watts_per_meter_kelvin := a.convertFromBase(ThermalConductivityWattPerMeterKelvin)
	a.watts_per_meter_kelvinLazy = &watts_per_meter_kelvin
	return watts_per_meter_kelvin
}

// BtusPerHourFootFahrenheit returns the ThermalConductivity value in BtusPerHourFootFahrenheit.
//
// 
func (a *ThermalConductivity) BtusPerHourFootFahrenheit() float64 {
	if a.btus_per_hour_foot_fahrenheitLazy != nil {
		return *a.btus_per_hour_foot_fahrenheitLazy
	}
	btus_per_hour_foot_fahrenheit := a.convertFromBase(ThermalConductivityBtuPerHourFootFahrenheit)
	a.btus_per_hour_foot_fahrenheitLazy = &btus_per_hour_foot_fahrenheit
	return btus_per_hour_foot_fahrenheit
}


// ToDto creates a ThermalConductivityDto representation from the ThermalConductivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerMeterKelvin by default.
func (a *ThermalConductivity) ToDto(holdInUnit *ThermalConductivityUnits) ThermalConductivityDto {
	if holdInUnit == nil {
		defaultUnit := ThermalConductivityWattPerMeterKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return ThermalConductivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ThermalConductivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerMeterKelvin by default.
func (a *ThermalConductivity) ToDtoJSON(holdInUnit *ThermalConductivityUnits) ([]byte, error) {
	// Convert to ThermalConductivityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ThermalConductivity to a specific unit value.
// The function uses the provided unit type (ThermalConductivityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ThermalConductivity) Convert(toUnit ThermalConductivityUnits) float64 {
	switch toUnit { 
    case ThermalConductivityWattPerMeterKelvin:
		return a.WattsPerMeterKelvin()
    case ThermalConductivityBtuPerHourFootFahrenheit:
		return a.BtusPerHourFootFahrenheit()
	default:
		return math.NaN()
	}
}

func (a *ThermalConductivity) convertFromBase(toUnit ThermalConductivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ThermalConductivityWattPerMeterKelvin:
		return (value) 
	case ThermalConductivityBtuPerHourFootFahrenheit:
		return (value / 1.73073467) 
	default:
		return math.NaN()
	}
}

func (a *ThermalConductivity) convertToBase(value float64, fromUnit ThermalConductivityUnits) float64 {
	switch fromUnit { 
	case ThermalConductivityWattPerMeterKelvin:
		return (value) 
	case ThermalConductivityBtuPerHourFootFahrenheit:
		return (value * 1.73073467) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ThermalConductivity in the default unit (WattPerMeterKelvin),
// formatted to two decimal places.
func (a ThermalConductivity) String() string {
	return a.ToString(ThermalConductivityWattPerMeterKelvin, 2)
}

// ToString formats the ThermalConductivity value as a string with the specified unit and fractional digits.
// It converts the ThermalConductivity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ThermalConductivity value will be converted (e.g., WattPerMeterKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ThermalConductivity with the unit abbreviation.
func (a *ThermalConductivity) ToString(unit ThermalConductivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetThermalConductivityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetThermalConductivityAbbreviation(unit))
}

// Equals checks if the given ThermalConductivity is equal to the current ThermalConductivity.
//
// Parameters:
//    other: The ThermalConductivity to compare against.
//
// Returns:
//    bool: Returns true if both ThermalConductivity are equal, false otherwise.
func (a *ThermalConductivity) Equals(other *ThermalConductivity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ThermalConductivity with another ThermalConductivity.
// It returns -1 if the current ThermalConductivity is less than the other ThermalConductivity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ThermalConductivity to compare against.
//
// Returns:
//    int: -1 if the current ThermalConductivity is less, 1 if greater, and 0 if equal.
func (a *ThermalConductivity) CompareTo(other *ThermalConductivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ThermalConductivity to the current ThermalConductivity and returns the result.
// The result is a new ThermalConductivity instance with the sum of the values.
//
// Parameters:
//    other: The ThermalConductivity to add to the current ThermalConductivity.
//
// Returns:
//    *ThermalConductivity: A new ThermalConductivity instance representing the sum of both ThermalConductivity.
func (a *ThermalConductivity) Add(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ThermalConductivity from the current ThermalConductivity and returns the result.
// The result is a new ThermalConductivity instance with the difference of the values.
//
// Parameters:
//    other: The ThermalConductivity to subtract from the current ThermalConductivity.
//
// Returns:
//    *ThermalConductivity: A new ThermalConductivity instance representing the difference of both ThermalConductivity.
func (a *ThermalConductivity) Subtract(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ThermalConductivity by the given ThermalConductivity and returns the result.
// The result is a new ThermalConductivity instance with the product of the values.
//
// Parameters:
//    other: The ThermalConductivity to multiply with the current ThermalConductivity.
//
// Returns:
//    *ThermalConductivity: A new ThermalConductivity instance representing the product of both ThermalConductivity.
func (a *ThermalConductivity) Multiply(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value * other.BaseValue()}
}

// Divide divides the current ThermalConductivity by the given ThermalConductivity and returns the result.
// The result is a new ThermalConductivity instance with the quotient of the values.
//
// Parameters:
//    other: The ThermalConductivity to divide the current ThermalConductivity by.
//
// Returns:
//    *ThermalConductivity: A new ThermalConductivity instance representing the quotient of both ThermalConductivity.
func (a *ThermalConductivity) Divide(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value / other.BaseValue()}
}

// GetThermalConductivityAbbreviation gets the unit abbreviation.
func GetThermalConductivityAbbreviation(unit ThermalConductivityUnits) string {
	switch unit { 
	case ThermalConductivityWattPerMeterKelvin:
		return "W/m·K" 
	case ThermalConductivityBtuPerHourFootFahrenheit:
		return "BTU/h·ft·°F" 
	default:
		return ""
	}
}