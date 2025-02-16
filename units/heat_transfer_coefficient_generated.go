// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// HeatTransferCoefficientUnits defines various units of HeatTransferCoefficient.
type HeatTransferCoefficientUnits string

const (
    
        // 
        HeatTransferCoefficientWattPerSquareMeterKelvin HeatTransferCoefficientUnits = "WattPerSquareMeterKelvin"
        // 
        HeatTransferCoefficientWattPerSquareMeterCelsius HeatTransferCoefficientUnits = "WattPerSquareMeterCelsius"
        // 
        HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit HeatTransferCoefficientUnits = "BtuPerHourSquareFootDegreeFahrenheit"
        // 
        HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius HeatTransferCoefficientUnits = "CaloriePerHourSquareMeterDegreeCelsius"
        // 
        HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius HeatTransferCoefficientUnits = "KilocaloriePerHourSquareMeterDegreeCelsius"
)

// HeatTransferCoefficientDto represents a HeatTransferCoefficient measurement with a numerical value and its corresponding unit.
type HeatTransferCoefficientDto struct {
    // Value is the numerical representation of the HeatTransferCoefficient.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the HeatTransferCoefficient, as defined in the HeatTransferCoefficientUnits enumeration.
	Unit  HeatTransferCoefficientUnits `json:"unit"`
}

// HeatTransferCoefficientDtoFactory groups methods for creating and serializing HeatTransferCoefficientDto objects.
type HeatTransferCoefficientDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a HeatTransferCoefficientDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf HeatTransferCoefficientDtoFactory) FromJSON(data []byte) (*HeatTransferCoefficientDto, error) {
	a := HeatTransferCoefficientDto{}

    // Parse JSON into HeatTransferCoefficientDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a HeatTransferCoefficientDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a HeatTransferCoefficientDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// HeatTransferCoefficient represents a measurement in a of HeatTransferCoefficient.
//
// The heat transfer coefficient or film coefficient, or film effectiveness, in thermodynamics and in mechanics is the proportionality constant between the heat flux and the thermodynamic driving force for the flow of heat (i.e., the temperature difference, ΔT)
type HeatTransferCoefficient struct {
	// value is the base measurement stored internally.
	value       float64
    
    watts_per_square_meter_kelvinLazy *float64 
    watts_per_square_meter_celsiusLazy *float64 
    btus_per_hour_square_foot_degree_fahrenheitLazy *float64 
    calories_per_hour_square_meter_degree_celsiusLazy *float64 
    kilocalories_per_hour_square_meter_degree_celsiusLazy *float64 
}

// HeatTransferCoefficientFactory groups methods for creating HeatTransferCoefficient instances.
type HeatTransferCoefficientFactory struct{}

// CreateHeatTransferCoefficient creates a new HeatTransferCoefficient instance from the given value and unit.
func (uf HeatTransferCoefficientFactory) CreateHeatTransferCoefficient(value float64, unit HeatTransferCoefficientUnits) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, unit)
}

// FromDto converts a HeatTransferCoefficientDto to a HeatTransferCoefficient instance.
func (uf HeatTransferCoefficientFactory) FromDto(dto HeatTransferCoefficientDto) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a HeatTransferCoefficient instance.
func (uf HeatTransferCoefficientFactory) FromDtoJSON(data []byte) (*HeatTransferCoefficient, error) {
	unitDto, err := HeatTransferCoefficientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HeatTransferCoefficientDto from JSON: %w", err)
	}
	return HeatTransferCoefficientFactory{}.FromDto(*unitDto)
}


// FromWattsPerSquareMeterKelvin creates a new HeatTransferCoefficient instance from a value in WattsPerSquareMeterKelvin.
func (uf HeatTransferCoefficientFactory) FromWattsPerSquareMeterKelvin(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientWattPerSquareMeterKelvin)
}

// FromWattsPerSquareMeterCelsius creates a new HeatTransferCoefficient instance from a value in WattsPerSquareMeterCelsius.
func (uf HeatTransferCoefficientFactory) FromWattsPerSquareMeterCelsius(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientWattPerSquareMeterCelsius)
}

// FromBtusPerHourSquareFootDegreeFahrenheit creates a new HeatTransferCoefficient instance from a value in BtusPerHourSquareFootDegreeFahrenheit.
func (uf HeatTransferCoefficientFactory) FromBtusPerHourSquareFootDegreeFahrenheit(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
}

// FromCaloriesPerHourSquareMeterDegreeCelsius creates a new HeatTransferCoefficient instance from a value in CaloriesPerHourSquareMeterDegreeCelsius.
func (uf HeatTransferCoefficientFactory) FromCaloriesPerHourSquareMeterDegreeCelsius(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
}

// FromKilocaloriesPerHourSquareMeterDegreeCelsius creates a new HeatTransferCoefficient instance from a value in KilocaloriesPerHourSquareMeterDegreeCelsius.
func (uf HeatTransferCoefficientFactory) FromKilocaloriesPerHourSquareMeterDegreeCelsius(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
}


// newHeatTransferCoefficient creates a new HeatTransferCoefficient.
func newHeatTransferCoefficient(value float64, fromUnit HeatTransferCoefficientUnits) (*HeatTransferCoefficient, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &HeatTransferCoefficient{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of HeatTransferCoefficient in WattPerSquareMeterKelvin unit (the base/default quantity).
func (a *HeatTransferCoefficient) BaseValue() float64 {
	return a.value
}


// WattsPerSquareMeterKelvin returns the HeatTransferCoefficient value in WattsPerSquareMeterKelvin.
//
// 
func (a *HeatTransferCoefficient) WattsPerSquareMeterKelvin() float64 {
	if a.watts_per_square_meter_kelvinLazy != nil {
		return *a.watts_per_square_meter_kelvinLazy
	}
	watts_per_square_meter_kelvin := a.convertFromBase(HeatTransferCoefficientWattPerSquareMeterKelvin)
	a.watts_per_square_meter_kelvinLazy = &watts_per_square_meter_kelvin
	return watts_per_square_meter_kelvin
}

// WattsPerSquareMeterCelsius returns the HeatTransferCoefficient value in WattsPerSquareMeterCelsius.
//
// 
func (a *HeatTransferCoefficient) WattsPerSquareMeterCelsius() float64 {
	if a.watts_per_square_meter_celsiusLazy != nil {
		return *a.watts_per_square_meter_celsiusLazy
	}
	watts_per_square_meter_celsius := a.convertFromBase(HeatTransferCoefficientWattPerSquareMeterCelsius)
	a.watts_per_square_meter_celsiusLazy = &watts_per_square_meter_celsius
	return watts_per_square_meter_celsius
}

// BtusPerHourSquareFootDegreeFahrenheit returns the HeatTransferCoefficient value in BtusPerHourSquareFootDegreeFahrenheit.
//
// 
func (a *HeatTransferCoefficient) BtusPerHourSquareFootDegreeFahrenheit() float64 {
	if a.btus_per_hour_square_foot_degree_fahrenheitLazy != nil {
		return *a.btus_per_hour_square_foot_degree_fahrenheitLazy
	}
	btus_per_hour_square_foot_degree_fahrenheit := a.convertFromBase(HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
	a.btus_per_hour_square_foot_degree_fahrenheitLazy = &btus_per_hour_square_foot_degree_fahrenheit
	return btus_per_hour_square_foot_degree_fahrenheit
}

// CaloriesPerHourSquareMeterDegreeCelsius returns the HeatTransferCoefficient value in CaloriesPerHourSquareMeterDegreeCelsius.
//
// 
func (a *HeatTransferCoefficient) CaloriesPerHourSquareMeterDegreeCelsius() float64 {
	if a.calories_per_hour_square_meter_degree_celsiusLazy != nil {
		return *a.calories_per_hour_square_meter_degree_celsiusLazy
	}
	calories_per_hour_square_meter_degree_celsius := a.convertFromBase(HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
	a.calories_per_hour_square_meter_degree_celsiusLazy = &calories_per_hour_square_meter_degree_celsius
	return calories_per_hour_square_meter_degree_celsius
}

// KilocaloriesPerHourSquareMeterDegreeCelsius returns the HeatTransferCoefficient value in KilocaloriesPerHourSquareMeterDegreeCelsius.
//
// 
func (a *HeatTransferCoefficient) KilocaloriesPerHourSquareMeterDegreeCelsius() float64 {
	if a.kilocalories_per_hour_square_meter_degree_celsiusLazy != nil {
		return *a.kilocalories_per_hour_square_meter_degree_celsiusLazy
	}
	kilocalories_per_hour_square_meter_degree_celsius := a.convertFromBase(HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
	a.kilocalories_per_hour_square_meter_degree_celsiusLazy = &kilocalories_per_hour_square_meter_degree_celsius
	return kilocalories_per_hour_square_meter_degree_celsius
}


// ToDto creates a HeatTransferCoefficientDto representation from the HeatTransferCoefficient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeterKelvin by default.
func (a *HeatTransferCoefficient) ToDto(holdInUnit *HeatTransferCoefficientUnits) HeatTransferCoefficientDto {
	if holdInUnit == nil {
		defaultUnit := HeatTransferCoefficientWattPerSquareMeterKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return HeatTransferCoefficientDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the HeatTransferCoefficient instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeterKelvin by default.
func (a *HeatTransferCoefficient) ToDtoJSON(holdInUnit *HeatTransferCoefficientUnits) ([]byte, error) {
	// Convert to HeatTransferCoefficientDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a HeatTransferCoefficient to a specific unit value.
// The function uses the provided unit type (HeatTransferCoefficientUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *HeatTransferCoefficient) Convert(toUnit HeatTransferCoefficientUnits) float64 {
	switch toUnit { 
    case HeatTransferCoefficientWattPerSquareMeterKelvin:
		return a.WattsPerSquareMeterKelvin()
    case HeatTransferCoefficientWattPerSquareMeterCelsius:
		return a.WattsPerSquareMeterCelsius()
    case HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit:
		return a.BtusPerHourSquareFootDegreeFahrenheit()
    case HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius:
		return a.CaloriesPerHourSquareMeterDegreeCelsius()
    case HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius:
		return a.KilocaloriesPerHourSquareMeterDegreeCelsius()
	default:
		return math.NaN()
	}
}

func (a *HeatTransferCoefficient) convertFromBase(toUnit HeatTransferCoefficientUnits) float64 {
    value := a.value
	switch toUnit { 
	case HeatTransferCoefficientWattPerSquareMeterKelvin:
		return (value) 
	case HeatTransferCoefficientWattPerSquareMeterCelsius:
		return (value) 
	case HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit:
		return (value / 5.6782633411134878) 
	case HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius:
		return ((value / 4.1868) * 3600) 
	case HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius:
		return (((value / 4.1868) * 3600) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *HeatTransferCoefficient) convertToBase(value float64, fromUnit HeatTransferCoefficientUnits) float64 {
	switch fromUnit { 
	case HeatTransferCoefficientWattPerSquareMeterKelvin:
		return (value) 
	case HeatTransferCoefficientWattPerSquareMeterCelsius:
		return (value) 
	case HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit:
		return (value * 5.6782633411134878) 
	case HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius:
		return ((value * 4.1868) / 3600) 
	case HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius:
		return (((value * 4.1868) / 3600) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the HeatTransferCoefficient in the default unit (WattPerSquareMeterKelvin),
// formatted to two decimal places.
func (a HeatTransferCoefficient) String() string {
	return a.ToString(HeatTransferCoefficientWattPerSquareMeterKelvin, 2)
}

// ToString formats the HeatTransferCoefficient value as a string with the specified unit and fractional digits.
// It converts the HeatTransferCoefficient to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the HeatTransferCoefficient value will be converted (e.g., WattPerSquareMeterKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the HeatTransferCoefficient with the unit abbreviation.
func (a *HeatTransferCoefficient) ToString(unit HeatTransferCoefficientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetHeatTransferCoefficientAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetHeatTransferCoefficientAbbreviation(unit))
}

// Equals checks if the given HeatTransferCoefficient is equal to the current HeatTransferCoefficient.
//
// Parameters:
//    other: The HeatTransferCoefficient to compare against.
//
// Returns:
//    bool: Returns true if both HeatTransferCoefficient are equal, false otherwise.
func (a *HeatTransferCoefficient) Equals(other *HeatTransferCoefficient) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current HeatTransferCoefficient with another HeatTransferCoefficient.
// It returns -1 if the current HeatTransferCoefficient is less than the other HeatTransferCoefficient, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The HeatTransferCoefficient to compare against.
//
// Returns:
//    int: -1 if the current HeatTransferCoefficient is less, 1 if greater, and 0 if equal.
func (a *HeatTransferCoefficient) CompareTo(other *HeatTransferCoefficient) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given HeatTransferCoefficient to the current HeatTransferCoefficient and returns the result.
// The result is a new HeatTransferCoefficient instance with the sum of the values.
//
// Parameters:
//    other: The HeatTransferCoefficient to add to the current HeatTransferCoefficient.
//
// Returns:
//    *HeatTransferCoefficient: A new HeatTransferCoefficient instance representing the sum of both HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Add(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given HeatTransferCoefficient from the current HeatTransferCoefficient and returns the result.
// The result is a new HeatTransferCoefficient instance with the difference of the values.
//
// Parameters:
//    other: The HeatTransferCoefficient to subtract from the current HeatTransferCoefficient.
//
// Returns:
//    *HeatTransferCoefficient: A new HeatTransferCoefficient instance representing the difference of both HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Subtract(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current HeatTransferCoefficient by the given HeatTransferCoefficient and returns the result.
// The result is a new HeatTransferCoefficient instance with the product of the values.
//
// Parameters:
//    other: The HeatTransferCoefficient to multiply with the current HeatTransferCoefficient.
//
// Returns:
//    *HeatTransferCoefficient: A new HeatTransferCoefficient instance representing the product of both HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Multiply(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value * other.BaseValue()}
}

// Divide divides the current HeatTransferCoefficient by the given HeatTransferCoefficient and returns the result.
// The result is a new HeatTransferCoefficient instance with the quotient of the values.
//
// Parameters:
//    other: The HeatTransferCoefficient to divide the current HeatTransferCoefficient by.
//
// Returns:
//    *HeatTransferCoefficient: A new HeatTransferCoefficient instance representing the quotient of both HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Divide(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value / other.BaseValue()}
}

// GetHeatTransferCoefficientAbbreviation gets the unit abbreviation.
func GetHeatTransferCoefficientAbbreviation(unit HeatTransferCoefficientUnits) string {
	switch unit { 
	case HeatTransferCoefficientWattPerSquareMeterKelvin:
		return "W/m²·K" 
	case HeatTransferCoefficientWattPerSquareMeterCelsius:
		return "W/m²·°C" 
	case HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit:
		return "Btu/h·ft²·°F" 
	case HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius:
		return "kcal/h·m²·°C" 
	case HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius:
		return "kkcal/h·m²·°C" 
	default:
		return ""
	}
}