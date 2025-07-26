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
        ThermalResistanceKelvinPerWatt ThermalResistanceUnits = "KelvinPerWatt"
        // 
        ThermalResistanceDegreeCelsiusPerWatt ThermalResistanceUnits = "DegreeCelsiusPerWatt"
)

var internalThermalResistanceUnitsMap = map[ThermalResistanceUnits]bool{
	
	ThermalResistanceKelvinPerWatt: true,
	ThermalResistanceDegreeCelsiusPerWatt: true,
}

// ThermalResistanceDto represents a ThermalResistance measurement with a numerical value and its corresponding unit.
type ThermalResistanceDto struct {
    // Value is the numerical representation of the ThermalResistance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ThermalResistance, as defined in the ThermalResistanceUnits enumeration.
	Unit  ThermalResistanceUnits `json:"unit" validate:"required,oneof=KelvinPerWatt DegreeCelsiusPerWatt"`
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

	if a.Unit == "" {
		return nil, errors.New("unit is required")
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
// Thermal resistance (R) measures the opposition to the heat current in a material or system. It is measured in units of kelvins per watt (K/W) and indicates how much temperature difference (in kelvins) is required to transfer a unit of heat current (in watts) through the material or object. It is essential to optimize the building insulation, evaluate the efficiency of electronic devices, and enhance the performance of heat sinks in various applications.
type ThermalResistance struct {
	// value is the base measurement stored internally.
	value       float64
    
    kelvins_per_wattLazy *float64 
    degrees_celsius_per_wattLazy *float64 
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


// FromKelvinsPerWatt creates a new ThermalResistance instance from a value in KelvinsPerWatt.
func (uf ThermalResistanceFactory) FromKelvinsPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceKelvinPerWatt)
}

// FromDegreesCelsiusPerWatt creates a new ThermalResistance instance from a value in DegreesCelsiusPerWatt.
func (uf ThermalResistanceFactory) FromDegreesCelsiusPerWatt(value float64) (*ThermalResistance, error) {
	return newThermalResistance(value, ThermalResistanceDegreeCelsiusPerWatt)
}


// newThermalResistance creates a new ThermalResistance.
func newThermalResistance(value float64, fromUnit ThermalResistanceUnits) (*ThermalResistance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalThermalResistanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ThermalResistanceUnits", fromUnit)
	}
	a := &ThermalResistance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ThermalResistance in KelvinPerWatt unit (the base/default quantity).
func (a *ThermalResistance) BaseValue() float64 {
	return a.value
}


// KelvinsPerWatt returns the ThermalResistance value in KelvinsPerWatt.
//
// 
func (a *ThermalResistance) KelvinsPerWatt() float64 {
	if a.kelvins_per_wattLazy != nil {
		return *a.kelvins_per_wattLazy
	}
	kelvins_per_watt := a.convertFromBase(ThermalResistanceKelvinPerWatt)
	a.kelvins_per_wattLazy = &kelvins_per_watt
	return kelvins_per_watt
}

// DegreesCelsiusPerWatt returns the ThermalResistance value in DegreesCelsiusPerWatt.
//
// 
func (a *ThermalResistance) DegreesCelsiusPerWatt() float64 {
	if a.degrees_celsius_per_wattLazy != nil {
		return *a.degrees_celsius_per_wattLazy
	}
	degrees_celsius_per_watt := a.convertFromBase(ThermalResistanceDegreeCelsiusPerWatt)
	a.degrees_celsius_per_wattLazy = &degrees_celsius_per_watt
	return degrees_celsius_per_watt
}


// ToDto creates a ThermalResistanceDto representation from the ThermalResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KelvinPerWatt by default.
func (a *ThermalResistance) ToDto(holdInUnit *ThermalResistanceUnits) ThermalResistanceDto {
	if holdInUnit == nil {
		defaultUnit := ThermalResistanceKelvinPerWatt // Default value
		holdInUnit = &defaultUnit
	}

	return ThermalResistanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ThermalResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KelvinPerWatt by default.
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
    case ThermalResistanceKelvinPerWatt:
		return a.KelvinsPerWatt()
    case ThermalResistanceDegreeCelsiusPerWatt:
		return a.DegreesCelsiusPerWatt()
	default:
		return math.NaN()
	}
}

func (a *ThermalResistance) convertFromBase(toUnit ThermalResistanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ThermalResistanceKelvinPerWatt:
		return (value) 
	case ThermalResistanceDegreeCelsiusPerWatt:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *ThermalResistance) convertToBase(value float64, fromUnit ThermalResistanceUnits) float64 {
	switch fromUnit { 
	case ThermalResistanceKelvinPerWatt:
		return (value) 
	case ThermalResistanceDegreeCelsiusPerWatt:
		return (value) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ThermalResistance in the default unit (KelvinPerWatt),
// formatted to two decimal places.
func (a ThermalResistance) String() string {
	return a.ToString(ThermalResistanceKelvinPerWatt, 2)
}

// ToString formats the ThermalResistance value as a string with the specified unit and fractional digits.
// It converts the ThermalResistance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ThermalResistance value will be converted (e.g., KelvinPerWatt))
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
	case ThermalResistanceKelvinPerWatt:
		return "K/W" 
	case ThermalResistanceDegreeCelsiusPerWatt:
		return "°C/W" 
	default:
		return ""
	}
}