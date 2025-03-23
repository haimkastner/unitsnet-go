// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VolumetricHeatCapacityUnits defines various units of VolumetricHeatCapacity.
type VolumetricHeatCapacityUnits string

const (
    
        // 
        VolumetricHeatCapacityJoulePerCubicMeterKelvin VolumetricHeatCapacityUnits = "JoulePerCubicMeterKelvin"
        // 
        VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius VolumetricHeatCapacityUnits = "JoulePerCubicMeterDegreeCelsius"
        // 
        VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius VolumetricHeatCapacityUnits = "CaloriePerCubicCentimeterDegreeCelsius"
        // 
        VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit VolumetricHeatCapacityUnits = "BtuPerCubicFootDegreeFahrenheit"
        // 
        VolumetricHeatCapacityKilojoulePerCubicMeterKelvin VolumetricHeatCapacityUnits = "KilojoulePerCubicMeterKelvin"
        // 
        VolumetricHeatCapacityMegajoulePerCubicMeterKelvin VolumetricHeatCapacityUnits = "MegajoulePerCubicMeterKelvin"
        // 
        VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius VolumetricHeatCapacityUnits = "KilojoulePerCubicMeterDegreeCelsius"
        // 
        VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius VolumetricHeatCapacityUnits = "MegajoulePerCubicMeterDegreeCelsius"
        // 
        VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius VolumetricHeatCapacityUnits = "KilocaloriePerCubicCentimeterDegreeCelsius"
)

var internalVolumetricHeatCapacityUnitsMap = map[VolumetricHeatCapacityUnits]bool{
	
	VolumetricHeatCapacityJoulePerCubicMeterKelvin: true,
	VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius: true,
	VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius: true,
	VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit: true,
	VolumetricHeatCapacityKilojoulePerCubicMeterKelvin: true,
	VolumetricHeatCapacityMegajoulePerCubicMeterKelvin: true,
	VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius: true,
	VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius: true,
	VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius: true,
}

// VolumetricHeatCapacityDto represents a VolumetricHeatCapacity measurement with a numerical value and its corresponding unit.
type VolumetricHeatCapacityDto struct {
    // Value is the numerical representation of the VolumetricHeatCapacity.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the VolumetricHeatCapacity, as defined in the VolumetricHeatCapacityUnits enumeration.
	Unit  VolumetricHeatCapacityUnits `json:"unit" validate:"required,oneof=JoulePerCubicMeterKelvin JoulePerCubicMeterDegreeCelsius CaloriePerCubicCentimeterDegreeCelsius BtuPerCubicFootDegreeFahrenheit KilojoulePerCubicMeterKelvin MegajoulePerCubicMeterKelvin KilojoulePerCubicMeterDegreeCelsius MegajoulePerCubicMeterDegreeCelsius KilocaloriePerCubicCentimeterDegreeCelsius"`
}

// VolumetricHeatCapacityDtoFactory groups methods for creating and serializing VolumetricHeatCapacityDto objects.
type VolumetricHeatCapacityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumetricHeatCapacityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumetricHeatCapacityDtoFactory) FromJSON(data []byte) (*VolumetricHeatCapacityDto, error) {
	a := VolumetricHeatCapacityDto{}

    // Parse JSON into VolumetricHeatCapacityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a VolumetricHeatCapacityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumetricHeatCapacityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// VolumetricHeatCapacity represents a measurement in a of VolumetricHeatCapacity.
//
// The volumetric heat capacity is the amount of energy that must be added, in the form of heat, to one unit of volume of the material in order to cause an increase of one unit in its temperature.
type VolumetricHeatCapacity struct {
	// value is the base measurement stored internally.
	value       float64
    
    joules_per_cubic_meter_kelvinLazy *float64 
    joules_per_cubic_meter_degree_celsiusLazy *float64 
    calories_per_cubic_centimeter_degree_celsiusLazy *float64 
    btus_per_cubic_foot_degree_fahrenheitLazy *float64 
    kilojoules_per_cubic_meter_kelvinLazy *float64 
    megajoules_per_cubic_meter_kelvinLazy *float64 
    kilojoules_per_cubic_meter_degree_celsiusLazy *float64 
    megajoules_per_cubic_meter_degree_celsiusLazy *float64 
    kilocalories_per_cubic_centimeter_degree_celsiusLazy *float64 
}

// VolumetricHeatCapacityFactory groups methods for creating VolumetricHeatCapacity instances.
type VolumetricHeatCapacityFactory struct{}

// CreateVolumetricHeatCapacity creates a new VolumetricHeatCapacity instance from the given value and unit.
func (uf VolumetricHeatCapacityFactory) CreateVolumetricHeatCapacity(value float64, unit VolumetricHeatCapacityUnits) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, unit)
}

// FromDto converts a VolumetricHeatCapacityDto to a VolumetricHeatCapacity instance.
func (uf VolumetricHeatCapacityFactory) FromDto(dto VolumetricHeatCapacityDto) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VolumetricHeatCapacity instance.
func (uf VolumetricHeatCapacityFactory) FromDtoJSON(data []byte) (*VolumetricHeatCapacity, error) {
	unitDto, err := VolumetricHeatCapacityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumetricHeatCapacityDto from JSON: %w", err)
	}
	return VolumetricHeatCapacityFactory{}.FromDto(*unitDto)
}


// FromJoulesPerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from a value in JoulesPerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromJoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityJoulePerCubicMeterKelvin)
}

// FromJoulesPerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from a value in JoulesPerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromJoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
}

// FromCaloriesPerCubicCentimeterDegreeCelsius creates a new VolumetricHeatCapacity instance from a value in CaloriesPerCubicCentimeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromCaloriesPerCubicCentimeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
}

// FromBtusPerCubicFootDegreeFahrenheit creates a new VolumetricHeatCapacity instance from a value in BtusPerCubicFootDegreeFahrenheit.
func (uf VolumetricHeatCapacityFactory) FromBtusPerCubicFootDegreeFahrenheit(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
}

// FromKilojoulesPerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from a value in KilojoulesPerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromKilojoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
}

// FromMegajoulesPerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from a value in MegajoulesPerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromMegajoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
}

// FromKilojoulesPerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from a value in KilojoulesPerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromKilojoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
}

// FromMegajoulesPerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from a value in MegajoulesPerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromMegajoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
}

// FromKilocaloriesPerCubicCentimeterDegreeCelsius creates a new VolumetricHeatCapacity instance from a value in KilocaloriesPerCubicCentimeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromKilocaloriesPerCubicCentimeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
}


// newVolumetricHeatCapacity creates a new VolumetricHeatCapacity.
func newVolumetricHeatCapacity(value float64, fromUnit VolumetricHeatCapacityUnits) (*VolumetricHeatCapacity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalVolumetricHeatCapacityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in VolumetricHeatCapacityUnits", fromUnit)
	}
	a := &VolumetricHeatCapacity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumetricHeatCapacity in JoulePerCubicMeterKelvin unit (the base/default quantity).
func (a *VolumetricHeatCapacity) BaseValue() float64 {
	return a.value
}


// JoulesPerCubicMeterKelvin returns the VolumetricHeatCapacity value in JoulesPerCubicMeterKelvin.
//
// 
func (a *VolumetricHeatCapacity) JoulesPerCubicMeterKelvin() float64 {
	if a.joules_per_cubic_meter_kelvinLazy != nil {
		return *a.joules_per_cubic_meter_kelvinLazy
	}
	joules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a.joules_per_cubic_meter_kelvinLazy = &joules_per_cubic_meter_kelvin
	return joules_per_cubic_meter_kelvin
}

// JoulesPerCubicMeterDegreeCelsius returns the VolumetricHeatCapacity value in JoulesPerCubicMeterDegreeCelsius.
//
// 
func (a *VolumetricHeatCapacity) JoulesPerCubicMeterDegreeCelsius() float64 {
	if a.joules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.joules_per_cubic_meter_degree_celsiusLazy
	}
	joules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
	a.joules_per_cubic_meter_degree_celsiusLazy = &joules_per_cubic_meter_degree_celsius
	return joules_per_cubic_meter_degree_celsius
}

// CaloriesPerCubicCentimeterDegreeCelsius returns the VolumetricHeatCapacity value in CaloriesPerCubicCentimeterDegreeCelsius.
//
// 
func (a *VolumetricHeatCapacity) CaloriesPerCubicCentimeterDegreeCelsius() float64 {
	if a.calories_per_cubic_centimeter_degree_celsiusLazy != nil {
		return *a.calories_per_cubic_centimeter_degree_celsiusLazy
	}
	calories_per_cubic_centimeter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
	a.calories_per_cubic_centimeter_degree_celsiusLazy = &calories_per_cubic_centimeter_degree_celsius
	return calories_per_cubic_centimeter_degree_celsius
}

// BtusPerCubicFootDegreeFahrenheit returns the VolumetricHeatCapacity value in BtusPerCubicFootDegreeFahrenheit.
//
// 
func (a *VolumetricHeatCapacity) BtusPerCubicFootDegreeFahrenheit() float64 {
	if a.btus_per_cubic_foot_degree_fahrenheitLazy != nil {
		return *a.btus_per_cubic_foot_degree_fahrenheitLazy
	}
	btus_per_cubic_foot_degree_fahrenheit := a.convertFromBase(VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
	a.btus_per_cubic_foot_degree_fahrenheitLazy = &btus_per_cubic_foot_degree_fahrenheit
	return btus_per_cubic_foot_degree_fahrenheit
}

// KilojoulesPerCubicMeterKelvin returns the VolumetricHeatCapacity value in KilojoulesPerCubicMeterKelvin.
//
// 
func (a *VolumetricHeatCapacity) KilojoulesPerCubicMeterKelvin() float64 {
	if a.kilojoules_per_cubic_meter_kelvinLazy != nil {
		return *a.kilojoules_per_cubic_meter_kelvinLazy
	}
	kilojoules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
	a.kilojoules_per_cubic_meter_kelvinLazy = &kilojoules_per_cubic_meter_kelvin
	return kilojoules_per_cubic_meter_kelvin
}

// MegajoulesPerCubicMeterKelvin returns the VolumetricHeatCapacity value in MegajoulesPerCubicMeterKelvin.
//
// 
func (a *VolumetricHeatCapacity) MegajoulesPerCubicMeterKelvin() float64 {
	if a.megajoules_per_cubic_meter_kelvinLazy != nil {
		return *a.megajoules_per_cubic_meter_kelvinLazy
	}
	megajoules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
	a.megajoules_per_cubic_meter_kelvinLazy = &megajoules_per_cubic_meter_kelvin
	return megajoules_per_cubic_meter_kelvin
}

// KilojoulesPerCubicMeterDegreeCelsius returns the VolumetricHeatCapacity value in KilojoulesPerCubicMeterDegreeCelsius.
//
// 
func (a *VolumetricHeatCapacity) KilojoulesPerCubicMeterDegreeCelsius() float64 {
	if a.kilojoules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.kilojoules_per_cubic_meter_degree_celsiusLazy
	}
	kilojoules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
	a.kilojoules_per_cubic_meter_degree_celsiusLazy = &kilojoules_per_cubic_meter_degree_celsius
	return kilojoules_per_cubic_meter_degree_celsius
}

// MegajoulesPerCubicMeterDegreeCelsius returns the VolumetricHeatCapacity value in MegajoulesPerCubicMeterDegreeCelsius.
//
// 
func (a *VolumetricHeatCapacity) MegajoulesPerCubicMeterDegreeCelsius() float64 {
	if a.megajoules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.megajoules_per_cubic_meter_degree_celsiusLazy
	}
	megajoules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
	a.megajoules_per_cubic_meter_degree_celsiusLazy = &megajoules_per_cubic_meter_degree_celsius
	return megajoules_per_cubic_meter_degree_celsius
}

// KilocaloriesPerCubicCentimeterDegreeCelsius returns the VolumetricHeatCapacity value in KilocaloriesPerCubicCentimeterDegreeCelsius.
//
// 
func (a *VolumetricHeatCapacity) KilocaloriesPerCubicCentimeterDegreeCelsius() float64 {
	if a.kilocalories_per_cubic_centimeter_degree_celsiusLazy != nil {
		return *a.kilocalories_per_cubic_centimeter_degree_celsiusLazy
	}
	kilocalories_per_cubic_centimeter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
	a.kilocalories_per_cubic_centimeter_degree_celsiusLazy = &kilocalories_per_cubic_centimeter_degree_celsius
	return kilocalories_per_cubic_centimeter_degree_celsius
}


// ToDto creates a VolumetricHeatCapacityDto representation from the VolumetricHeatCapacity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerCubicMeterKelvin by default.
func (a *VolumetricHeatCapacity) ToDto(holdInUnit *VolumetricHeatCapacityUnits) VolumetricHeatCapacityDto {
	if holdInUnit == nil {
		defaultUnit := VolumetricHeatCapacityJoulePerCubicMeterKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return VolumetricHeatCapacityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VolumetricHeatCapacity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by JoulePerCubicMeterKelvin by default.
func (a *VolumetricHeatCapacity) ToDtoJSON(holdInUnit *VolumetricHeatCapacityUnits) ([]byte, error) {
	// Convert to VolumetricHeatCapacityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VolumetricHeatCapacity to a specific unit value.
// The function uses the provided unit type (VolumetricHeatCapacityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VolumetricHeatCapacity) Convert(toUnit VolumetricHeatCapacityUnits) float64 {
	switch toUnit { 
    case VolumetricHeatCapacityJoulePerCubicMeterKelvin:
		return a.JoulesPerCubicMeterKelvin()
    case VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius:
		return a.JoulesPerCubicMeterDegreeCelsius()
    case VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius:
		return a.CaloriesPerCubicCentimeterDegreeCelsius()
    case VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit:
		return a.BtusPerCubicFootDegreeFahrenheit()
    case VolumetricHeatCapacityKilojoulePerCubicMeterKelvin:
		return a.KilojoulesPerCubicMeterKelvin()
    case VolumetricHeatCapacityMegajoulePerCubicMeterKelvin:
		return a.MegajoulesPerCubicMeterKelvin()
    case VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius:
		return a.KilojoulesPerCubicMeterDegreeCelsius()
    case VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius:
		return a.MegajoulesPerCubicMeterDegreeCelsius()
    case VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius:
		return a.KilocaloriesPerCubicCentimeterDegreeCelsius()
	default:
		return math.NaN()
	}
}

func (a *VolumetricHeatCapacity) convertFromBase(toUnit VolumetricHeatCapacityUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumetricHeatCapacityJoulePerCubicMeterKelvin:
		return (value) 
	case VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius:
		return (value) 
	case VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius:
		return (value * 2.388459e-7) 
	case VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit:
		return (value * 1.4910660e-5) 
	case VolumetricHeatCapacityKilojoulePerCubicMeterKelvin:
		return ((value) / 1000.0) 
	case VolumetricHeatCapacityMegajoulePerCubicMeterKelvin:
		return ((value) / 1000000.0) 
	case VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius:
		return ((value) / 1000.0) 
	case VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius:
		return ((value) / 1000000.0) 
	case VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius:
		return ((value * 2.388459e-7) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *VolumetricHeatCapacity) convertToBase(value float64, fromUnit VolumetricHeatCapacityUnits) float64 {
	switch fromUnit { 
	case VolumetricHeatCapacityJoulePerCubicMeterKelvin:
		return (value) 
	case VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius:
		return (value) 
	case VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius:
		return (value / 2.388459e-7) 
	case VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit:
		return (value / 1.4910660e-5) 
	case VolumetricHeatCapacityKilojoulePerCubicMeterKelvin:
		return ((value) * 1000.0) 
	case VolumetricHeatCapacityMegajoulePerCubicMeterKelvin:
		return ((value) * 1000000.0) 
	case VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius:
		return ((value) * 1000.0) 
	case VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius:
		return ((value) * 1000000.0) 
	case VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius:
		return ((value / 2.388459e-7) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VolumetricHeatCapacity in the default unit (JoulePerCubicMeterKelvin),
// formatted to two decimal places.
func (a VolumetricHeatCapacity) String() string {
	return a.ToString(VolumetricHeatCapacityJoulePerCubicMeterKelvin, 2)
}

// ToString formats the VolumetricHeatCapacity value as a string with the specified unit and fractional digits.
// It converts the VolumetricHeatCapacity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VolumetricHeatCapacity value will be converted (e.g., JoulePerCubicMeterKelvin))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VolumetricHeatCapacity with the unit abbreviation.
func (a *VolumetricHeatCapacity) ToString(unit VolumetricHeatCapacityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetVolumetricHeatCapacityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetVolumetricHeatCapacityAbbreviation(unit))
}

// Equals checks if the given VolumetricHeatCapacity is equal to the current VolumetricHeatCapacity.
//
// Parameters:
//    other: The VolumetricHeatCapacity to compare against.
//
// Returns:
//    bool: Returns true if both VolumetricHeatCapacity are equal, false otherwise.
func (a *VolumetricHeatCapacity) Equals(other *VolumetricHeatCapacity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VolumetricHeatCapacity with another VolumetricHeatCapacity.
// It returns -1 if the current VolumetricHeatCapacity is less than the other VolumetricHeatCapacity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VolumetricHeatCapacity to compare against.
//
// Returns:
//    int: -1 if the current VolumetricHeatCapacity is less, 1 if greater, and 0 if equal.
func (a *VolumetricHeatCapacity) CompareTo(other *VolumetricHeatCapacity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VolumetricHeatCapacity to the current VolumetricHeatCapacity and returns the result.
// The result is a new VolumetricHeatCapacity instance with the sum of the values.
//
// Parameters:
//    other: The VolumetricHeatCapacity to add to the current VolumetricHeatCapacity.
//
// Returns:
//    *VolumetricHeatCapacity: A new VolumetricHeatCapacity instance representing the sum of both VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Add(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VolumetricHeatCapacity from the current VolumetricHeatCapacity and returns the result.
// The result is a new VolumetricHeatCapacity instance with the difference of the values.
//
// Parameters:
//    other: The VolumetricHeatCapacity to subtract from the current VolumetricHeatCapacity.
//
// Returns:
//    *VolumetricHeatCapacity: A new VolumetricHeatCapacity instance representing the difference of both VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Subtract(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VolumetricHeatCapacity by the given VolumetricHeatCapacity and returns the result.
// The result is a new VolumetricHeatCapacity instance with the product of the values.
//
// Parameters:
//    other: The VolumetricHeatCapacity to multiply with the current VolumetricHeatCapacity.
//
// Returns:
//    *VolumetricHeatCapacity: A new VolumetricHeatCapacity instance representing the product of both VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Multiply(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value * other.BaseValue()}
}

// Divide divides the current VolumetricHeatCapacity by the given VolumetricHeatCapacity and returns the result.
// The result is a new VolumetricHeatCapacity instance with the quotient of the values.
//
// Parameters:
//    other: The VolumetricHeatCapacity to divide the current VolumetricHeatCapacity by.
//
// Returns:
//    *VolumetricHeatCapacity: A new VolumetricHeatCapacity instance representing the quotient of both VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Divide(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value / other.BaseValue()}
}

// GetVolumetricHeatCapacityAbbreviation gets the unit abbreviation.
func GetVolumetricHeatCapacityAbbreviation(unit VolumetricHeatCapacityUnits) string {
	switch unit { 
	case VolumetricHeatCapacityJoulePerCubicMeterKelvin:
		return "J/m³·K" 
	case VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius:
		return "J/m³·°C" 
	case VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius:
		return "cal/cm³·°C" 
	case VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit:
		return "BTU/ft³·°F" 
	case VolumetricHeatCapacityKilojoulePerCubicMeterKelvin:
		return "kJ/m³·K" 
	case VolumetricHeatCapacityMegajoulePerCubicMeterKelvin:
		return "MJ/m³·K" 
	case VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius:
		return "kJ/m³·°C" 
	case VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius:
		return "MJ/m³·°C" 
	case VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius:
		return "kcal/cm³·°C" 
	default:
		return ""
	}
}