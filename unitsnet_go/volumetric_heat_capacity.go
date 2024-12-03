package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumetricHeatCapacityUnits enumeration
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

// VolumetricHeatCapacityDto represents an VolumetricHeatCapacity
type VolumetricHeatCapacityDto struct {
	Value float64
	Unit  VolumetricHeatCapacityUnits
}

// VolumetricHeatCapacityDtoFactory struct to group related functions
type VolumetricHeatCapacityDtoFactory struct{}

func (udf VolumetricHeatCapacityDtoFactory) FromJSON(data []byte) (*VolumetricHeatCapacityDto, error) {
	a := VolumetricHeatCapacityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumetricHeatCapacityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VolumetricHeatCapacity struct
type VolumetricHeatCapacity struct {
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

// VolumetricHeatCapacityFactory struct to group related functions
type VolumetricHeatCapacityFactory struct{}

func (uf VolumetricHeatCapacityFactory) CreateVolumetricHeatCapacity(value float64, unit VolumetricHeatCapacityUnits) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, unit)
}

func (uf VolumetricHeatCapacityFactory) FromDto(dto VolumetricHeatCapacityDto) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(dto.Value, dto.Unit)
}

func (uf VolumetricHeatCapacityFactory) FromDtoJSON(data []byte) (*VolumetricHeatCapacity, error) {
	unitDto, err := VolumetricHeatCapacityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumetricHeatCapacityFactory{}.FromDto(*unitDto)
}


// FromJoulePerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from JoulePerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromJoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityJoulePerCubicMeterKelvin)
}

// FromJoulePerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from JoulePerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromJoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
}

// FromCaloriePerCubicCentimeterDegreeCelsius creates a new VolumetricHeatCapacity instance from CaloriePerCubicCentimeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromCaloriesPerCubicCentimeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
}

// FromBtuPerCubicFootDegreeFahrenheit creates a new VolumetricHeatCapacity instance from BtuPerCubicFootDegreeFahrenheit.
func (uf VolumetricHeatCapacityFactory) FromBtusPerCubicFootDegreeFahrenheit(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
}

// FromKilojoulePerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from KilojoulePerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromKilojoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
}

// FromMegajoulePerCubicMeterKelvin creates a new VolumetricHeatCapacity instance from MegajoulePerCubicMeterKelvin.
func (uf VolumetricHeatCapacityFactory) FromMegajoulesPerCubicMeterKelvin(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
}

// FromKilojoulePerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from KilojoulePerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromKilojoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
}

// FromMegajoulePerCubicMeterDegreeCelsius creates a new VolumetricHeatCapacity instance from MegajoulePerCubicMeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromMegajoulesPerCubicMeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
}

// FromKilocaloriePerCubicCentimeterDegreeCelsius creates a new VolumetricHeatCapacity instance from KilocaloriePerCubicCentimeterDegreeCelsius.
func (uf VolumetricHeatCapacityFactory) FromKilocaloriesPerCubicCentimeterDegreeCelsius(value float64) (*VolumetricHeatCapacity, error) {
	return newVolumetricHeatCapacity(value, VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
}




// newVolumetricHeatCapacity creates a new VolumetricHeatCapacity.
func newVolumetricHeatCapacity(value float64, fromUnit VolumetricHeatCapacityUnits) (*VolumetricHeatCapacity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VolumetricHeatCapacity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumetricHeatCapacity in JoulePerCubicMeterKelvin.
func (a *VolumetricHeatCapacity) BaseValue() float64 {
	return a.value
}


// JoulePerCubicMeterKelvin returns the value in JoulePerCubicMeterKelvin.
func (a *VolumetricHeatCapacity) JoulesPerCubicMeterKelvin() float64 {
	if a.joules_per_cubic_meter_kelvinLazy != nil {
		return *a.joules_per_cubic_meter_kelvinLazy
	}
	joules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityJoulePerCubicMeterKelvin)
	a.joules_per_cubic_meter_kelvinLazy = &joules_per_cubic_meter_kelvin
	return joules_per_cubic_meter_kelvin
}

// JoulePerCubicMeterDegreeCelsius returns the value in JoulePerCubicMeterDegreeCelsius.
func (a *VolumetricHeatCapacity) JoulesPerCubicMeterDegreeCelsius() float64 {
	if a.joules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.joules_per_cubic_meter_degree_celsiusLazy
	}
	joules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityJoulePerCubicMeterDegreeCelsius)
	a.joules_per_cubic_meter_degree_celsiusLazy = &joules_per_cubic_meter_degree_celsius
	return joules_per_cubic_meter_degree_celsius
}

// CaloriePerCubicCentimeterDegreeCelsius returns the value in CaloriePerCubicCentimeterDegreeCelsius.
func (a *VolumetricHeatCapacity) CaloriesPerCubicCentimeterDegreeCelsius() float64 {
	if a.calories_per_cubic_centimeter_degree_celsiusLazy != nil {
		return *a.calories_per_cubic_centimeter_degree_celsiusLazy
	}
	calories_per_cubic_centimeter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityCaloriePerCubicCentimeterDegreeCelsius)
	a.calories_per_cubic_centimeter_degree_celsiusLazy = &calories_per_cubic_centimeter_degree_celsius
	return calories_per_cubic_centimeter_degree_celsius
}

// BtuPerCubicFootDegreeFahrenheit returns the value in BtuPerCubicFootDegreeFahrenheit.
func (a *VolumetricHeatCapacity) BtusPerCubicFootDegreeFahrenheit() float64 {
	if a.btus_per_cubic_foot_degree_fahrenheitLazy != nil {
		return *a.btus_per_cubic_foot_degree_fahrenheitLazy
	}
	btus_per_cubic_foot_degree_fahrenheit := a.convertFromBase(VolumetricHeatCapacityBtuPerCubicFootDegreeFahrenheit)
	a.btus_per_cubic_foot_degree_fahrenheitLazy = &btus_per_cubic_foot_degree_fahrenheit
	return btus_per_cubic_foot_degree_fahrenheit
}

// KilojoulePerCubicMeterKelvin returns the value in KilojoulePerCubicMeterKelvin.
func (a *VolumetricHeatCapacity) KilojoulesPerCubicMeterKelvin() float64 {
	if a.kilojoules_per_cubic_meter_kelvinLazy != nil {
		return *a.kilojoules_per_cubic_meter_kelvinLazy
	}
	kilojoules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityKilojoulePerCubicMeterKelvin)
	a.kilojoules_per_cubic_meter_kelvinLazy = &kilojoules_per_cubic_meter_kelvin
	return kilojoules_per_cubic_meter_kelvin
}

// MegajoulePerCubicMeterKelvin returns the value in MegajoulePerCubicMeterKelvin.
func (a *VolumetricHeatCapacity) MegajoulesPerCubicMeterKelvin() float64 {
	if a.megajoules_per_cubic_meter_kelvinLazy != nil {
		return *a.megajoules_per_cubic_meter_kelvinLazy
	}
	megajoules_per_cubic_meter_kelvin := a.convertFromBase(VolumetricHeatCapacityMegajoulePerCubicMeterKelvin)
	a.megajoules_per_cubic_meter_kelvinLazy = &megajoules_per_cubic_meter_kelvin
	return megajoules_per_cubic_meter_kelvin
}

// KilojoulePerCubicMeterDegreeCelsius returns the value in KilojoulePerCubicMeterDegreeCelsius.
func (a *VolumetricHeatCapacity) KilojoulesPerCubicMeterDegreeCelsius() float64 {
	if a.kilojoules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.kilojoules_per_cubic_meter_degree_celsiusLazy
	}
	kilojoules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityKilojoulePerCubicMeterDegreeCelsius)
	a.kilojoules_per_cubic_meter_degree_celsiusLazy = &kilojoules_per_cubic_meter_degree_celsius
	return kilojoules_per_cubic_meter_degree_celsius
}

// MegajoulePerCubicMeterDegreeCelsius returns the value in MegajoulePerCubicMeterDegreeCelsius.
func (a *VolumetricHeatCapacity) MegajoulesPerCubicMeterDegreeCelsius() float64 {
	if a.megajoules_per_cubic_meter_degree_celsiusLazy != nil {
		return *a.megajoules_per_cubic_meter_degree_celsiusLazy
	}
	megajoules_per_cubic_meter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityMegajoulePerCubicMeterDegreeCelsius)
	a.megajoules_per_cubic_meter_degree_celsiusLazy = &megajoules_per_cubic_meter_degree_celsius
	return megajoules_per_cubic_meter_degree_celsius
}

// KilocaloriePerCubicCentimeterDegreeCelsius returns the value in KilocaloriePerCubicCentimeterDegreeCelsius.
func (a *VolumetricHeatCapacity) KilocaloriesPerCubicCentimeterDegreeCelsius() float64 {
	if a.kilocalories_per_cubic_centimeter_degree_celsiusLazy != nil {
		return *a.kilocalories_per_cubic_centimeter_degree_celsiusLazy
	}
	kilocalories_per_cubic_centimeter_degree_celsius := a.convertFromBase(VolumetricHeatCapacityKilocaloriePerCubicCentimeterDegreeCelsius)
	a.kilocalories_per_cubic_centimeter_degree_celsiusLazy = &kilocalories_per_cubic_centimeter_degree_celsius
	return kilocalories_per_cubic_centimeter_degree_celsius
}


// ToDto creates an VolumetricHeatCapacityDto representation.
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

// ToDtoJSON creates an VolumetricHeatCapacityDto representation.
func (a *VolumetricHeatCapacity) ToDtoJSON(holdInUnit *VolumetricHeatCapacityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VolumetricHeatCapacity to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a VolumetricHeatCapacity) String() string {
	return a.ToString(VolumetricHeatCapacityJoulePerCubicMeterKelvin, 2)
}

// ToString formats the VolumetricHeatCapacity to string.
// fractionalDigits -1 for not mention
func (a *VolumetricHeatCapacity) ToString(unit VolumetricHeatCapacityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VolumetricHeatCapacity) getUnitAbbreviation(unit VolumetricHeatCapacityUnits) string {
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

// Check if the given VolumetricHeatCapacity are equals to the current VolumetricHeatCapacity
func (a *VolumetricHeatCapacity) Equals(other *VolumetricHeatCapacity) bool {
	return a.value == other.BaseValue()
}

// Check if the given VolumetricHeatCapacity are equals to the current VolumetricHeatCapacity
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

// Add the given VolumetricHeatCapacity to the current VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Add(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value + other.BaseValue()}
}

// Subtract the given VolumetricHeatCapacity to the current VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Subtract(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value - other.BaseValue()}
}

// Multiply the given VolumetricHeatCapacity to the current VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Multiply(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value * other.BaseValue()}
}

// Divide the given VolumetricHeatCapacity to the current VolumetricHeatCapacity.
func (a *VolumetricHeatCapacity) Divide(other *VolumetricHeatCapacity) *VolumetricHeatCapacity {
	return &VolumetricHeatCapacity{value: a.value / other.BaseValue()}
}