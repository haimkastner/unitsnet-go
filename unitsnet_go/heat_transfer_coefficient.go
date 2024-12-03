package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// HeatTransferCoefficientUnits enumeration
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

// HeatTransferCoefficientDto represents an HeatTransferCoefficient
type HeatTransferCoefficientDto struct {
	Value float64
	Unit  HeatTransferCoefficientUnits
}

// HeatTransferCoefficientDtoFactory struct to group related functions
type HeatTransferCoefficientDtoFactory struct{}

func (udf HeatTransferCoefficientDtoFactory) FromJSON(data []byte) (*HeatTransferCoefficientDto, error) {
	a := HeatTransferCoefficientDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a HeatTransferCoefficientDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// HeatTransferCoefficient struct
type HeatTransferCoefficient struct {
	value       float64
    
    watts_per_square_meter_kelvinLazy *float64 
    watts_per_square_meter_celsiusLazy *float64 
    btus_per_hour_square_foot_degree_fahrenheitLazy *float64 
    calories_per_hour_square_meter_degree_celsiusLazy *float64 
    kilocalories_per_hour_square_meter_degree_celsiusLazy *float64 
}

// HeatTransferCoefficientFactory struct to group related functions
type HeatTransferCoefficientFactory struct{}

func (uf HeatTransferCoefficientFactory) CreateHeatTransferCoefficient(value float64, unit HeatTransferCoefficientUnits) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, unit)
}

func (uf HeatTransferCoefficientFactory) FromDto(dto HeatTransferCoefficientDto) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(dto.Value, dto.Unit)
}

func (uf HeatTransferCoefficientFactory) FromDtoJSON(data []byte) (*HeatTransferCoefficient, error) {
	unitDto, err := HeatTransferCoefficientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return HeatTransferCoefficientFactory{}.FromDto(*unitDto)
}


// FromWattPerSquareMeterKelvin creates a new HeatTransferCoefficient instance from WattPerSquareMeterKelvin.
func (uf HeatTransferCoefficientFactory) FromWattsPerSquareMeterKelvin(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientWattPerSquareMeterKelvin)
}

// FromWattPerSquareMeterCelsius creates a new HeatTransferCoefficient instance from WattPerSquareMeterCelsius.
func (uf HeatTransferCoefficientFactory) FromWattsPerSquareMeterCelsius(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientWattPerSquareMeterCelsius)
}

// FromBtuPerHourSquareFootDegreeFahrenheit creates a new HeatTransferCoefficient instance from BtuPerHourSquareFootDegreeFahrenheit.
func (uf HeatTransferCoefficientFactory) FromBtusPerHourSquareFootDegreeFahrenheit(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
}

// FromCaloriePerHourSquareMeterDegreeCelsius creates a new HeatTransferCoefficient instance from CaloriePerHourSquareMeterDegreeCelsius.
func (uf HeatTransferCoefficientFactory) FromCaloriesPerHourSquareMeterDegreeCelsius(value float64) (*HeatTransferCoefficient, error) {
	return newHeatTransferCoefficient(value, HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
}

// FromKilocaloriePerHourSquareMeterDegreeCelsius creates a new HeatTransferCoefficient instance from KilocaloriePerHourSquareMeterDegreeCelsius.
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

// BaseValue returns the base value of HeatTransferCoefficient in WattPerSquareMeterKelvin.
func (a *HeatTransferCoefficient) BaseValue() float64 {
	return a.value
}


// WattPerSquareMeterKelvin returns the value in WattPerSquareMeterKelvin.
func (a *HeatTransferCoefficient) WattsPerSquareMeterKelvin() float64 {
	if a.watts_per_square_meter_kelvinLazy != nil {
		return *a.watts_per_square_meter_kelvinLazy
	}
	watts_per_square_meter_kelvin := a.convertFromBase(HeatTransferCoefficientWattPerSquareMeterKelvin)
	a.watts_per_square_meter_kelvinLazy = &watts_per_square_meter_kelvin
	return watts_per_square_meter_kelvin
}

// WattPerSquareMeterCelsius returns the value in WattPerSquareMeterCelsius.
func (a *HeatTransferCoefficient) WattsPerSquareMeterCelsius() float64 {
	if a.watts_per_square_meter_celsiusLazy != nil {
		return *a.watts_per_square_meter_celsiusLazy
	}
	watts_per_square_meter_celsius := a.convertFromBase(HeatTransferCoefficientWattPerSquareMeterCelsius)
	a.watts_per_square_meter_celsiusLazy = &watts_per_square_meter_celsius
	return watts_per_square_meter_celsius
}

// BtuPerHourSquareFootDegreeFahrenheit returns the value in BtuPerHourSquareFootDegreeFahrenheit.
func (a *HeatTransferCoefficient) BtusPerHourSquareFootDegreeFahrenheit() float64 {
	if a.btus_per_hour_square_foot_degree_fahrenheitLazy != nil {
		return *a.btus_per_hour_square_foot_degree_fahrenheitLazy
	}
	btus_per_hour_square_foot_degree_fahrenheit := a.convertFromBase(HeatTransferCoefficientBtuPerHourSquareFootDegreeFahrenheit)
	a.btus_per_hour_square_foot_degree_fahrenheitLazy = &btus_per_hour_square_foot_degree_fahrenheit
	return btus_per_hour_square_foot_degree_fahrenheit
}

// CaloriePerHourSquareMeterDegreeCelsius returns the value in CaloriePerHourSquareMeterDegreeCelsius.
func (a *HeatTransferCoefficient) CaloriesPerHourSquareMeterDegreeCelsius() float64 {
	if a.calories_per_hour_square_meter_degree_celsiusLazy != nil {
		return *a.calories_per_hour_square_meter_degree_celsiusLazy
	}
	calories_per_hour_square_meter_degree_celsius := a.convertFromBase(HeatTransferCoefficientCaloriePerHourSquareMeterDegreeCelsius)
	a.calories_per_hour_square_meter_degree_celsiusLazy = &calories_per_hour_square_meter_degree_celsius
	return calories_per_hour_square_meter_degree_celsius
}

// KilocaloriePerHourSquareMeterDegreeCelsius returns the value in KilocaloriePerHourSquareMeterDegreeCelsius.
func (a *HeatTransferCoefficient) KilocaloriesPerHourSquareMeterDegreeCelsius() float64 {
	if a.kilocalories_per_hour_square_meter_degree_celsiusLazy != nil {
		return *a.kilocalories_per_hour_square_meter_degree_celsiusLazy
	}
	kilocalories_per_hour_square_meter_degree_celsius := a.convertFromBase(HeatTransferCoefficientKilocaloriePerHourSquareMeterDegreeCelsius)
	a.kilocalories_per_hour_square_meter_degree_celsiusLazy = &kilocalories_per_hour_square_meter_degree_celsius
	return kilocalories_per_hour_square_meter_degree_celsius
}


// ToDto creates an HeatTransferCoefficientDto representation.
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

// ToDtoJSON creates an HeatTransferCoefficientDto representation.
func (a *HeatTransferCoefficient) ToDtoJSON(holdInUnit *HeatTransferCoefficientUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts HeatTransferCoefficient to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a HeatTransferCoefficient) String() string {
	return a.ToString(HeatTransferCoefficientWattPerSquareMeterKelvin, 2)
}

// ToString formats the HeatTransferCoefficient to string.
// fractionalDigits -1 for not mention
func (a *HeatTransferCoefficient) ToString(unit HeatTransferCoefficientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *HeatTransferCoefficient) getUnitAbbreviation(unit HeatTransferCoefficientUnits) string {
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

// Check if the given HeatTransferCoefficient are equals to the current HeatTransferCoefficient
func (a *HeatTransferCoefficient) Equals(other *HeatTransferCoefficient) bool {
	return a.value == other.BaseValue()
}

// Check if the given HeatTransferCoefficient are equals to the current HeatTransferCoefficient
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

// Add the given HeatTransferCoefficient to the current HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Add(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value + other.BaseValue()}
}

// Subtract the given HeatTransferCoefficient to the current HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Subtract(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value - other.BaseValue()}
}

// Multiply the given HeatTransferCoefficient to the current HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Multiply(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value * other.BaseValue()}
}

// Divide the given HeatTransferCoefficient to the current HeatTransferCoefficient.
func (a *HeatTransferCoefficient) Divide(other *HeatTransferCoefficient) *HeatTransferCoefficient {
	return &HeatTransferCoefficient{value: a.value / other.BaseValue()}
}