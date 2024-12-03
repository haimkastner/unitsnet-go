package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ThermalConductivityUnits enumeration
type ThermalConductivityUnits string

const (
    
        // 
        ThermalConductivityWattPerMeterKelvin ThermalConductivityUnits = "WattPerMeterKelvin"
        // 
        ThermalConductivityBtuPerHourFootFahrenheit ThermalConductivityUnits = "BtuPerHourFootFahrenheit"
)

// ThermalConductivityDto represents an ThermalConductivity
type ThermalConductivityDto struct {
	Value float64
	Unit  ThermalConductivityUnits
}

// ThermalConductivityDtoFactory struct to group related functions
type ThermalConductivityDtoFactory struct{}

func (udf ThermalConductivityDtoFactory) FromJSON(data []byte) (*ThermalConductivityDto, error) {
	a := ThermalConductivityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ThermalConductivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ThermalConductivity struct
type ThermalConductivity struct {
	value       float64
    
    watts_per_meter_kelvinLazy *float64 
    btus_per_hour_foot_fahrenheitLazy *float64 
}

// ThermalConductivityFactory struct to group related functions
type ThermalConductivityFactory struct{}

func (uf ThermalConductivityFactory) CreateThermalConductivity(value float64, unit ThermalConductivityUnits) (*ThermalConductivity, error) {
	return newThermalConductivity(value, unit)
}

func (uf ThermalConductivityFactory) FromDto(dto ThermalConductivityDto) (*ThermalConductivity, error) {
	return newThermalConductivity(dto.Value, dto.Unit)
}

func (uf ThermalConductivityFactory) FromDtoJSON(data []byte) (*ThermalConductivity, error) {
	unitDto, err := ThermalConductivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ThermalConductivityFactory{}.FromDto(*unitDto)
}


// FromWattPerMeterKelvin creates a new ThermalConductivity instance from WattPerMeterKelvin.
func (uf ThermalConductivityFactory) FromWattsPerMeterKelvin(value float64) (*ThermalConductivity, error) {
	return newThermalConductivity(value, ThermalConductivityWattPerMeterKelvin)
}

// FromBtuPerHourFootFahrenheit creates a new ThermalConductivity instance from BtuPerHourFootFahrenheit.
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

// BaseValue returns the base value of ThermalConductivity in WattPerMeterKelvin.
func (a *ThermalConductivity) BaseValue() float64 {
	return a.value
}


// WattPerMeterKelvin returns the value in WattPerMeterKelvin.
func (a *ThermalConductivity) WattsPerMeterKelvin() float64 {
	if a.watts_per_meter_kelvinLazy != nil {
		return *a.watts_per_meter_kelvinLazy
	}
	watts_per_meter_kelvin := a.convertFromBase(ThermalConductivityWattPerMeterKelvin)
	a.watts_per_meter_kelvinLazy = &watts_per_meter_kelvin
	return watts_per_meter_kelvin
}

// BtuPerHourFootFahrenheit returns the value in BtuPerHourFootFahrenheit.
func (a *ThermalConductivity) BtusPerHourFootFahrenheit() float64 {
	if a.btus_per_hour_foot_fahrenheitLazy != nil {
		return *a.btus_per_hour_foot_fahrenheitLazy
	}
	btus_per_hour_foot_fahrenheit := a.convertFromBase(ThermalConductivityBtuPerHourFootFahrenheit)
	a.btus_per_hour_foot_fahrenheitLazy = &btus_per_hour_foot_fahrenheit
	return btus_per_hour_foot_fahrenheit
}


// ToDto creates an ThermalConductivityDto representation.
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

// ToDtoJSON creates an ThermalConductivityDto representation.
func (a *ThermalConductivity) ToDtoJSON(holdInUnit *ThermalConductivityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ThermalConductivity to a specific unit value.
func (a *ThermalConductivity) Convert(toUnit ThermalConductivityUnits) float64 {
	switch toUnit { 
    case ThermalConductivityWattPerMeterKelvin:
		return a.WattsPerMeterKelvin()
    case ThermalConductivityBtuPerHourFootFahrenheit:
		return a.BtusPerHourFootFahrenheit()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a ThermalConductivity) String() string {
	return a.ToString(ThermalConductivityWattPerMeterKelvin, 2)
}

// ToString formats the ThermalConductivity to string.
// fractionalDigits -1 for not mention
func (a *ThermalConductivity) ToString(unit ThermalConductivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ThermalConductivity) getUnitAbbreviation(unit ThermalConductivityUnits) string {
	switch unit { 
	case ThermalConductivityWattPerMeterKelvin:
		return "W/m·K" 
	case ThermalConductivityBtuPerHourFootFahrenheit:
		return "BTU/h·ft·°F" 
	default:
		return ""
	}
}

// Check if the given ThermalConductivity are equals to the current ThermalConductivity
func (a *ThermalConductivity) Equals(other *ThermalConductivity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ThermalConductivity are equals to the current ThermalConductivity
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

// Add the given ThermalConductivity to the current ThermalConductivity.
func (a *ThermalConductivity) Add(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value + other.BaseValue()}
}

// Subtract the given ThermalConductivity to the current ThermalConductivity.
func (a *ThermalConductivity) Subtract(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value - other.BaseValue()}
}

// Multiply the given ThermalConductivity to the current ThermalConductivity.
func (a *ThermalConductivity) Multiply(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value * other.BaseValue()}
}

// Divide the given ThermalConductivity to the current ThermalConductivity.
func (a *ThermalConductivity) Divide(other *ThermalConductivity) *ThermalConductivity {
	return &ThermalConductivity{value: a.value / other.BaseValue()}
}