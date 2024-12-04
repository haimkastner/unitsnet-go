// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// HeatFluxUnits defines various units of HeatFlux.
type HeatFluxUnits string

const (
    
        // 
        HeatFluxWattPerSquareMeter HeatFluxUnits = "WattPerSquareMeter"
        // 
        HeatFluxWattPerSquareInch HeatFluxUnits = "WattPerSquareInch"
        // 
        HeatFluxWattPerSquareFoot HeatFluxUnits = "WattPerSquareFoot"
        // 
        HeatFluxBtuPerSecondSquareInch HeatFluxUnits = "BtuPerSecondSquareInch"
        // 
        HeatFluxBtuPerSecondSquareFoot HeatFluxUnits = "BtuPerSecondSquareFoot"
        // 
        HeatFluxBtuPerMinuteSquareFoot HeatFluxUnits = "BtuPerMinuteSquareFoot"
        // 
        HeatFluxBtuPerHourSquareFoot HeatFluxUnits = "BtuPerHourSquareFoot"
        // 
        HeatFluxCaloriePerSecondSquareCentimeter HeatFluxUnits = "CaloriePerSecondSquareCentimeter"
        // 
        HeatFluxKilocaloriePerHourSquareMeter HeatFluxUnits = "KilocaloriePerHourSquareMeter"
        // 
        HeatFluxPoundForcePerFootSecond HeatFluxUnits = "PoundForcePerFootSecond"
        // 
        HeatFluxPoundPerSecondCubed HeatFluxUnits = "PoundPerSecondCubed"
        // 
        HeatFluxNanowattPerSquareMeter HeatFluxUnits = "NanowattPerSquareMeter"
        // 
        HeatFluxMicrowattPerSquareMeter HeatFluxUnits = "MicrowattPerSquareMeter"
        // 
        HeatFluxMilliwattPerSquareMeter HeatFluxUnits = "MilliwattPerSquareMeter"
        // 
        HeatFluxCentiwattPerSquareMeter HeatFluxUnits = "CentiwattPerSquareMeter"
        // 
        HeatFluxDeciwattPerSquareMeter HeatFluxUnits = "DeciwattPerSquareMeter"
        // 
        HeatFluxKilowattPerSquareMeter HeatFluxUnits = "KilowattPerSquareMeter"
        // 
        HeatFluxKilocaloriePerSecondSquareCentimeter HeatFluxUnits = "KilocaloriePerSecondSquareCentimeter"
)

// HeatFluxDto represents a HeatFlux measurement with a numerical value and its corresponding unit.
type HeatFluxDto struct {
    // Value is the numerical representation of the HeatFlux.
	Value float64
    // Unit specifies the unit of measurement for the HeatFlux, as defined in the HeatFluxUnits enumeration.
	Unit  HeatFluxUnits
}

// HeatFluxDtoFactory groups methods for creating and serializing HeatFluxDto objects.
type HeatFluxDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a HeatFluxDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf HeatFluxDtoFactory) FromJSON(data []byte) (*HeatFluxDto, error) {
	a := HeatFluxDto{}

    // Parse JSON into HeatFluxDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a HeatFluxDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a HeatFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// HeatFlux represents a measurement in a of HeatFlux.
//
// Heat flux is the flow of energy per unit of area per unit of time
type HeatFlux struct {
	// value is the base measurement stored internally.
	value       float64
    
    watts_per_square_meterLazy *float64 
    watts_per_square_inchLazy *float64 
    watts_per_square_footLazy *float64 
    btus_per_second_square_inchLazy *float64 
    btus_per_second_square_footLazy *float64 
    btus_per_minute_square_footLazy *float64 
    btus_per_hour_square_footLazy *float64 
    calories_per_second_square_centimeterLazy *float64 
    kilocalories_per_hour_square_meterLazy *float64 
    pounds_force_per_foot_secondLazy *float64 
    pounds_per_second_cubedLazy *float64 
    nanowatts_per_square_meterLazy *float64 
    microwatts_per_square_meterLazy *float64 
    milliwatts_per_square_meterLazy *float64 
    centiwatts_per_square_meterLazy *float64 
    deciwatts_per_square_meterLazy *float64 
    kilowatts_per_square_meterLazy *float64 
    kilocalories_per_second_square_centimeterLazy *float64 
}

// HeatFluxFactory groups methods for creating HeatFlux instances.
type HeatFluxFactory struct{}

// CreateHeatFlux creates a new HeatFlux instance from the given value and unit.
func (uf HeatFluxFactory) CreateHeatFlux(value float64, unit HeatFluxUnits) (*HeatFlux, error) {
	return newHeatFlux(value, unit)
}

// FromDto converts a HeatFluxDto to a HeatFlux instance.
func (uf HeatFluxFactory) FromDto(dto HeatFluxDto) (*HeatFlux, error) {
	return newHeatFlux(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a HeatFlux instance.
func (uf HeatFluxFactory) FromDtoJSON(data []byte) (*HeatFlux, error) {
	unitDto, err := HeatFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HeatFluxDto from JSON: %w", err)
	}
	return HeatFluxFactory{}.FromDto(*unitDto)
}


// FromWattsPerSquareMeter creates a new HeatFlux instance from a value in WattsPerSquareMeter.
func (uf HeatFluxFactory) FromWattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareMeter)
}

// FromWattsPerSquareInch creates a new HeatFlux instance from a value in WattsPerSquareInch.
func (uf HeatFluxFactory) FromWattsPerSquareInch(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareInch)
}

// FromWattsPerSquareFoot creates a new HeatFlux instance from a value in WattsPerSquareFoot.
func (uf HeatFluxFactory) FromWattsPerSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareFoot)
}

// FromBtusPerSecondSquareInch creates a new HeatFlux instance from a value in BtusPerSecondSquareInch.
func (uf HeatFluxFactory) FromBtusPerSecondSquareInch(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerSecondSquareInch)
}

// FromBtusPerSecondSquareFoot creates a new HeatFlux instance from a value in BtusPerSecondSquareFoot.
func (uf HeatFluxFactory) FromBtusPerSecondSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerSecondSquareFoot)
}

// FromBtusPerMinuteSquareFoot creates a new HeatFlux instance from a value in BtusPerMinuteSquareFoot.
func (uf HeatFluxFactory) FromBtusPerMinuteSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerMinuteSquareFoot)
}

// FromBtusPerHourSquareFoot creates a new HeatFlux instance from a value in BtusPerHourSquareFoot.
func (uf HeatFluxFactory) FromBtusPerHourSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerHourSquareFoot)
}

// FromCaloriesPerSecondSquareCentimeter creates a new HeatFlux instance from a value in CaloriesPerSecondSquareCentimeter.
func (uf HeatFluxFactory) FromCaloriesPerSecondSquareCentimeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxCaloriePerSecondSquareCentimeter)
}

// FromKilocaloriesPerHourSquareMeter creates a new HeatFlux instance from a value in KilocaloriesPerHourSquareMeter.
func (uf HeatFluxFactory) FromKilocaloriesPerHourSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxKilocaloriePerHourSquareMeter)
}

// FromPoundsForcePerFootSecond creates a new HeatFlux instance from a value in PoundsForcePerFootSecond.
func (uf HeatFluxFactory) FromPoundsForcePerFootSecond(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxPoundForcePerFootSecond)
}

// FromPoundsPerSecondCubed creates a new HeatFlux instance from a value in PoundsPerSecondCubed.
func (uf HeatFluxFactory) FromPoundsPerSecondCubed(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxPoundPerSecondCubed)
}

// FromNanowattsPerSquareMeter creates a new HeatFlux instance from a value in NanowattsPerSquareMeter.
func (uf HeatFluxFactory) FromNanowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxNanowattPerSquareMeter)
}

// FromMicrowattsPerSquareMeter creates a new HeatFlux instance from a value in MicrowattsPerSquareMeter.
func (uf HeatFluxFactory) FromMicrowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxMicrowattPerSquareMeter)
}

// FromMilliwattsPerSquareMeter creates a new HeatFlux instance from a value in MilliwattsPerSquareMeter.
func (uf HeatFluxFactory) FromMilliwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxMilliwattPerSquareMeter)
}

// FromCentiwattsPerSquareMeter creates a new HeatFlux instance from a value in CentiwattsPerSquareMeter.
func (uf HeatFluxFactory) FromCentiwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxCentiwattPerSquareMeter)
}

// FromDeciwattsPerSquareMeter creates a new HeatFlux instance from a value in DeciwattsPerSquareMeter.
func (uf HeatFluxFactory) FromDeciwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxDeciwattPerSquareMeter)
}

// FromKilowattsPerSquareMeter creates a new HeatFlux instance from a value in KilowattsPerSquareMeter.
func (uf HeatFluxFactory) FromKilowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxKilowattPerSquareMeter)
}

// FromKilocaloriesPerSecondSquareCentimeter creates a new HeatFlux instance from a value in KilocaloriesPerSecondSquareCentimeter.
func (uf HeatFluxFactory) FromKilocaloriesPerSecondSquareCentimeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxKilocaloriePerSecondSquareCentimeter)
}


// newHeatFlux creates a new HeatFlux.
func newHeatFlux(value float64, fromUnit HeatFluxUnits) (*HeatFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &HeatFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of HeatFlux in WattPerSquareMeter unit (the base/default quantity).
func (a *HeatFlux) BaseValue() float64 {
	return a.value
}


// WattsPerSquareMeter returns the HeatFlux value in WattsPerSquareMeter.
//
// 
func (a *HeatFlux) WattsPerSquareMeter() float64 {
	if a.watts_per_square_meterLazy != nil {
		return *a.watts_per_square_meterLazy
	}
	watts_per_square_meter := a.convertFromBase(HeatFluxWattPerSquareMeter)
	a.watts_per_square_meterLazy = &watts_per_square_meter
	return watts_per_square_meter
}

// WattsPerSquareInch returns the HeatFlux value in WattsPerSquareInch.
//
// 
func (a *HeatFlux) WattsPerSquareInch() float64 {
	if a.watts_per_square_inchLazy != nil {
		return *a.watts_per_square_inchLazy
	}
	watts_per_square_inch := a.convertFromBase(HeatFluxWattPerSquareInch)
	a.watts_per_square_inchLazy = &watts_per_square_inch
	return watts_per_square_inch
}

// WattsPerSquareFoot returns the HeatFlux value in WattsPerSquareFoot.
//
// 
func (a *HeatFlux) WattsPerSquareFoot() float64 {
	if a.watts_per_square_footLazy != nil {
		return *a.watts_per_square_footLazy
	}
	watts_per_square_foot := a.convertFromBase(HeatFluxWattPerSquareFoot)
	a.watts_per_square_footLazy = &watts_per_square_foot
	return watts_per_square_foot
}

// BtusPerSecondSquareInch returns the HeatFlux value in BtusPerSecondSquareInch.
//
// 
func (a *HeatFlux) BtusPerSecondSquareInch() float64 {
	if a.btus_per_second_square_inchLazy != nil {
		return *a.btus_per_second_square_inchLazy
	}
	btus_per_second_square_inch := a.convertFromBase(HeatFluxBtuPerSecondSquareInch)
	a.btus_per_second_square_inchLazy = &btus_per_second_square_inch
	return btus_per_second_square_inch
}

// BtusPerSecondSquareFoot returns the HeatFlux value in BtusPerSecondSquareFoot.
//
// 
func (a *HeatFlux) BtusPerSecondSquareFoot() float64 {
	if a.btus_per_second_square_footLazy != nil {
		return *a.btus_per_second_square_footLazy
	}
	btus_per_second_square_foot := a.convertFromBase(HeatFluxBtuPerSecondSquareFoot)
	a.btus_per_second_square_footLazy = &btus_per_second_square_foot
	return btus_per_second_square_foot
}

// BtusPerMinuteSquareFoot returns the HeatFlux value in BtusPerMinuteSquareFoot.
//
// 
func (a *HeatFlux) BtusPerMinuteSquareFoot() float64 {
	if a.btus_per_minute_square_footLazy != nil {
		return *a.btus_per_minute_square_footLazy
	}
	btus_per_minute_square_foot := a.convertFromBase(HeatFluxBtuPerMinuteSquareFoot)
	a.btus_per_minute_square_footLazy = &btus_per_minute_square_foot
	return btus_per_minute_square_foot
}

// BtusPerHourSquareFoot returns the HeatFlux value in BtusPerHourSquareFoot.
//
// 
func (a *HeatFlux) BtusPerHourSquareFoot() float64 {
	if a.btus_per_hour_square_footLazy != nil {
		return *a.btus_per_hour_square_footLazy
	}
	btus_per_hour_square_foot := a.convertFromBase(HeatFluxBtuPerHourSquareFoot)
	a.btus_per_hour_square_footLazy = &btus_per_hour_square_foot
	return btus_per_hour_square_foot
}

// CaloriesPerSecondSquareCentimeter returns the HeatFlux value in CaloriesPerSecondSquareCentimeter.
//
// 
func (a *HeatFlux) CaloriesPerSecondSquareCentimeter() float64 {
	if a.calories_per_second_square_centimeterLazy != nil {
		return *a.calories_per_second_square_centimeterLazy
	}
	calories_per_second_square_centimeter := a.convertFromBase(HeatFluxCaloriePerSecondSquareCentimeter)
	a.calories_per_second_square_centimeterLazy = &calories_per_second_square_centimeter
	return calories_per_second_square_centimeter
}

// KilocaloriesPerHourSquareMeter returns the HeatFlux value in KilocaloriesPerHourSquareMeter.
//
// 
func (a *HeatFlux) KilocaloriesPerHourSquareMeter() float64 {
	if a.kilocalories_per_hour_square_meterLazy != nil {
		return *a.kilocalories_per_hour_square_meterLazy
	}
	kilocalories_per_hour_square_meter := a.convertFromBase(HeatFluxKilocaloriePerHourSquareMeter)
	a.kilocalories_per_hour_square_meterLazy = &kilocalories_per_hour_square_meter
	return kilocalories_per_hour_square_meter
}

// PoundsForcePerFootSecond returns the HeatFlux value in PoundsForcePerFootSecond.
//
// 
func (a *HeatFlux) PoundsForcePerFootSecond() float64 {
	if a.pounds_force_per_foot_secondLazy != nil {
		return *a.pounds_force_per_foot_secondLazy
	}
	pounds_force_per_foot_second := a.convertFromBase(HeatFluxPoundForcePerFootSecond)
	a.pounds_force_per_foot_secondLazy = &pounds_force_per_foot_second
	return pounds_force_per_foot_second
}

// PoundsPerSecondCubed returns the HeatFlux value in PoundsPerSecondCubed.
//
// 
func (a *HeatFlux) PoundsPerSecondCubed() float64 {
	if a.pounds_per_second_cubedLazy != nil {
		return *a.pounds_per_second_cubedLazy
	}
	pounds_per_second_cubed := a.convertFromBase(HeatFluxPoundPerSecondCubed)
	a.pounds_per_second_cubedLazy = &pounds_per_second_cubed
	return pounds_per_second_cubed
}

// NanowattsPerSquareMeter returns the HeatFlux value in NanowattsPerSquareMeter.
//
// 
func (a *HeatFlux) NanowattsPerSquareMeter() float64 {
	if a.nanowatts_per_square_meterLazy != nil {
		return *a.nanowatts_per_square_meterLazy
	}
	nanowatts_per_square_meter := a.convertFromBase(HeatFluxNanowattPerSquareMeter)
	a.nanowatts_per_square_meterLazy = &nanowatts_per_square_meter
	return nanowatts_per_square_meter
}

// MicrowattsPerSquareMeter returns the HeatFlux value in MicrowattsPerSquareMeter.
//
// 
func (a *HeatFlux) MicrowattsPerSquareMeter() float64 {
	if a.microwatts_per_square_meterLazy != nil {
		return *a.microwatts_per_square_meterLazy
	}
	microwatts_per_square_meter := a.convertFromBase(HeatFluxMicrowattPerSquareMeter)
	a.microwatts_per_square_meterLazy = &microwatts_per_square_meter
	return microwatts_per_square_meter
}

// MilliwattsPerSquareMeter returns the HeatFlux value in MilliwattsPerSquareMeter.
//
// 
func (a *HeatFlux) MilliwattsPerSquareMeter() float64 {
	if a.milliwatts_per_square_meterLazy != nil {
		return *a.milliwatts_per_square_meterLazy
	}
	milliwatts_per_square_meter := a.convertFromBase(HeatFluxMilliwattPerSquareMeter)
	a.milliwatts_per_square_meterLazy = &milliwatts_per_square_meter
	return milliwatts_per_square_meter
}

// CentiwattsPerSquareMeter returns the HeatFlux value in CentiwattsPerSquareMeter.
//
// 
func (a *HeatFlux) CentiwattsPerSquareMeter() float64 {
	if a.centiwatts_per_square_meterLazy != nil {
		return *a.centiwatts_per_square_meterLazy
	}
	centiwatts_per_square_meter := a.convertFromBase(HeatFluxCentiwattPerSquareMeter)
	a.centiwatts_per_square_meterLazy = &centiwatts_per_square_meter
	return centiwatts_per_square_meter
}

// DeciwattsPerSquareMeter returns the HeatFlux value in DeciwattsPerSquareMeter.
//
// 
func (a *HeatFlux) DeciwattsPerSquareMeter() float64 {
	if a.deciwatts_per_square_meterLazy != nil {
		return *a.deciwatts_per_square_meterLazy
	}
	deciwatts_per_square_meter := a.convertFromBase(HeatFluxDeciwattPerSquareMeter)
	a.deciwatts_per_square_meterLazy = &deciwatts_per_square_meter
	return deciwatts_per_square_meter
}

// KilowattsPerSquareMeter returns the HeatFlux value in KilowattsPerSquareMeter.
//
// 
func (a *HeatFlux) KilowattsPerSquareMeter() float64 {
	if a.kilowatts_per_square_meterLazy != nil {
		return *a.kilowatts_per_square_meterLazy
	}
	kilowatts_per_square_meter := a.convertFromBase(HeatFluxKilowattPerSquareMeter)
	a.kilowatts_per_square_meterLazy = &kilowatts_per_square_meter
	return kilowatts_per_square_meter
}

// KilocaloriesPerSecondSquareCentimeter returns the HeatFlux value in KilocaloriesPerSecondSquareCentimeter.
//
// 
func (a *HeatFlux) KilocaloriesPerSecondSquareCentimeter() float64 {
	if a.kilocalories_per_second_square_centimeterLazy != nil {
		return *a.kilocalories_per_second_square_centimeterLazy
	}
	kilocalories_per_second_square_centimeter := a.convertFromBase(HeatFluxKilocaloriePerSecondSquareCentimeter)
	a.kilocalories_per_second_square_centimeterLazy = &kilocalories_per_second_square_centimeter
	return kilocalories_per_second_square_centimeter
}


// ToDto creates a HeatFluxDto representation from the HeatFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeter by default.
func (a *HeatFlux) ToDto(holdInUnit *HeatFluxUnits) HeatFluxDto {
	if holdInUnit == nil {
		defaultUnit := HeatFluxWattPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return HeatFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the HeatFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by WattPerSquareMeter by default.
func (a *HeatFlux) ToDtoJSON(holdInUnit *HeatFluxUnits) ([]byte, error) {
	// Convert to HeatFluxDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a HeatFlux to a specific unit value.
// The function uses the provided unit type (HeatFluxUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *HeatFlux) Convert(toUnit HeatFluxUnits) float64 {
	switch toUnit { 
    case HeatFluxWattPerSquareMeter:
		return a.WattsPerSquareMeter()
    case HeatFluxWattPerSquareInch:
		return a.WattsPerSquareInch()
    case HeatFluxWattPerSquareFoot:
		return a.WattsPerSquareFoot()
    case HeatFluxBtuPerSecondSquareInch:
		return a.BtusPerSecondSquareInch()
    case HeatFluxBtuPerSecondSquareFoot:
		return a.BtusPerSecondSquareFoot()
    case HeatFluxBtuPerMinuteSquareFoot:
		return a.BtusPerMinuteSquareFoot()
    case HeatFluxBtuPerHourSquareFoot:
		return a.BtusPerHourSquareFoot()
    case HeatFluxCaloriePerSecondSquareCentimeter:
		return a.CaloriesPerSecondSquareCentimeter()
    case HeatFluxKilocaloriePerHourSquareMeter:
		return a.KilocaloriesPerHourSquareMeter()
    case HeatFluxPoundForcePerFootSecond:
		return a.PoundsForcePerFootSecond()
    case HeatFluxPoundPerSecondCubed:
		return a.PoundsPerSecondCubed()
    case HeatFluxNanowattPerSquareMeter:
		return a.NanowattsPerSquareMeter()
    case HeatFluxMicrowattPerSquareMeter:
		return a.MicrowattsPerSquareMeter()
    case HeatFluxMilliwattPerSquareMeter:
		return a.MilliwattsPerSquareMeter()
    case HeatFluxCentiwattPerSquareMeter:
		return a.CentiwattsPerSquareMeter()
    case HeatFluxDeciwattPerSquareMeter:
		return a.DeciwattsPerSquareMeter()
    case HeatFluxKilowattPerSquareMeter:
		return a.KilowattsPerSquareMeter()
    case HeatFluxKilocaloriePerSecondSquareCentimeter:
		return a.KilocaloriesPerSecondSquareCentimeter()
	default:
		return math.NaN()
	}
}

func (a *HeatFlux) convertFromBase(toUnit HeatFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case HeatFluxWattPerSquareMeter:
		return (value) 
	case HeatFluxWattPerSquareInch:
		return (value / 1.5500031e3) 
	case HeatFluxWattPerSquareFoot:
		return (value / 1.07639e1) 
	case HeatFluxBtuPerSecondSquareInch:
		return (value / 1.63533984e6) 
	case HeatFluxBtuPerSecondSquareFoot:
		return (value / 1.13565267e4) 
	case HeatFluxBtuPerMinuteSquareFoot:
		return (value / 1.89275445e2) 
	case HeatFluxBtuPerHourSquareFoot:
		return (value / 3.15459075) 
	case HeatFluxCaloriePerSecondSquareCentimeter:
		return (value / 4.1868e4) 
	case HeatFluxKilocaloriePerHourSquareMeter:
		return (value / 1.163) 
	case HeatFluxPoundForcePerFootSecond:
		return (value / 1.459390293720636e1) 
	case HeatFluxPoundPerSecondCubed:
		return (value / 4.5359237e-1) 
	case HeatFluxNanowattPerSquareMeter:
		return ((value) / 1e-09) 
	case HeatFluxMicrowattPerSquareMeter:
		return ((value) / 1e-06) 
	case HeatFluxMilliwattPerSquareMeter:
		return ((value) / 0.001) 
	case HeatFluxCentiwattPerSquareMeter:
		return ((value) / 0.01) 
	case HeatFluxDeciwattPerSquareMeter:
		return ((value) / 0.1) 
	case HeatFluxKilowattPerSquareMeter:
		return ((value) / 1000.0) 
	case HeatFluxKilocaloriePerSecondSquareCentimeter:
		return ((value / 4.1868e4) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *HeatFlux) convertToBase(value float64, fromUnit HeatFluxUnits) float64 {
	switch fromUnit { 
	case HeatFluxWattPerSquareMeter:
		return (value) 
	case HeatFluxWattPerSquareInch:
		return (value * 1.5500031e3) 
	case HeatFluxWattPerSquareFoot:
		return (value * 1.07639e1) 
	case HeatFluxBtuPerSecondSquareInch:
		return (value * 1.63533984e6) 
	case HeatFluxBtuPerSecondSquareFoot:
		return (value * 1.13565267e4) 
	case HeatFluxBtuPerMinuteSquareFoot:
		return (value * 1.89275445e2) 
	case HeatFluxBtuPerHourSquareFoot:
		return (value * 3.15459075) 
	case HeatFluxCaloriePerSecondSquareCentimeter:
		return (value * 4.1868e4) 
	case HeatFluxKilocaloriePerHourSquareMeter:
		return (value * 1.163) 
	case HeatFluxPoundForcePerFootSecond:
		return (value * 1.459390293720636e1) 
	case HeatFluxPoundPerSecondCubed:
		return (value * 4.5359237e-1) 
	case HeatFluxNanowattPerSquareMeter:
		return ((value) * 1e-09) 
	case HeatFluxMicrowattPerSquareMeter:
		return ((value) * 1e-06) 
	case HeatFluxMilliwattPerSquareMeter:
		return ((value) * 0.001) 
	case HeatFluxCentiwattPerSquareMeter:
		return ((value) * 0.01) 
	case HeatFluxDeciwattPerSquareMeter:
		return ((value) * 0.1) 
	case HeatFluxKilowattPerSquareMeter:
		return ((value) * 1000.0) 
	case HeatFluxKilocaloriePerSecondSquareCentimeter:
		return ((value * 4.1868e4) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the HeatFlux in the default unit (WattPerSquareMeter),
// formatted to two decimal places.
func (a HeatFlux) String() string {
	return a.ToString(HeatFluxWattPerSquareMeter, 2)
}

// ToString formats the HeatFlux value as a string with the specified unit and fractional digits.
// It converts the HeatFlux to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the HeatFlux value will be converted (e.g., WattPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the HeatFlux with the unit abbreviation.
func (a *HeatFlux) ToString(unit HeatFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *HeatFlux) getUnitAbbreviation(unit HeatFluxUnits) string {
	switch unit { 
	case HeatFluxWattPerSquareMeter:
		return "W/m²" 
	case HeatFluxWattPerSquareInch:
		return "W/in²" 
	case HeatFluxWattPerSquareFoot:
		return "W/ft²" 
	case HeatFluxBtuPerSecondSquareInch:
		return "BTU/s·in²" 
	case HeatFluxBtuPerSecondSquareFoot:
		return "BTU/s·ft²" 
	case HeatFluxBtuPerMinuteSquareFoot:
		return "BTU/min·ft²" 
	case HeatFluxBtuPerHourSquareFoot:
		return "BTU/h·ft²" 
	case HeatFluxCaloriePerSecondSquareCentimeter:
		return "cal/s·cm²" 
	case HeatFluxKilocaloriePerHourSquareMeter:
		return "kcal/h·m²" 
	case HeatFluxPoundForcePerFootSecond:
		return "lbf/(ft·s)" 
	case HeatFluxPoundPerSecondCubed:
		return "lb/s³" 
	case HeatFluxNanowattPerSquareMeter:
		return "nW/m²" 
	case HeatFluxMicrowattPerSquareMeter:
		return "μW/m²" 
	case HeatFluxMilliwattPerSquareMeter:
		return "mW/m²" 
	case HeatFluxCentiwattPerSquareMeter:
		return "cW/m²" 
	case HeatFluxDeciwattPerSquareMeter:
		return "dW/m²" 
	case HeatFluxKilowattPerSquareMeter:
		return "kW/m²" 
	case HeatFluxKilocaloriePerSecondSquareCentimeter:
		return "kcal/s·cm²" 
	default:
		return ""
	}
}

// Equals checks if the given HeatFlux is equal to the current HeatFlux.
//
// Parameters:
//    other: The HeatFlux to compare against.
//
// Returns:
//    bool: Returns true if both HeatFlux are equal, false otherwise.
func (a *HeatFlux) Equals(other *HeatFlux) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current HeatFlux with another HeatFlux.
// It returns -1 if the current HeatFlux is less than the other HeatFlux, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The HeatFlux to compare against.
//
// Returns:
//    int: -1 if the current HeatFlux is less, 1 if greater, and 0 if equal.
func (a *HeatFlux) CompareTo(other *HeatFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given HeatFlux to the current HeatFlux and returns the result.
// The result is a new HeatFlux instance with the sum of the values.
//
// Parameters:
//    other: The HeatFlux to add to the current HeatFlux.
//
// Returns:
//    *HeatFlux: A new HeatFlux instance representing the sum of both HeatFlux.
func (a *HeatFlux) Add(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given HeatFlux from the current HeatFlux and returns the result.
// The result is a new HeatFlux instance with the difference of the values.
//
// Parameters:
//    other: The HeatFlux to subtract from the current HeatFlux.
//
// Returns:
//    *HeatFlux: A new HeatFlux instance representing the difference of both HeatFlux.
func (a *HeatFlux) Subtract(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current HeatFlux by the given HeatFlux and returns the result.
// The result is a new HeatFlux instance with the product of the values.
//
// Parameters:
//    other: The HeatFlux to multiply with the current HeatFlux.
//
// Returns:
//    *HeatFlux: A new HeatFlux instance representing the product of both HeatFlux.
func (a *HeatFlux) Multiply(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value * other.BaseValue()}
}

// Divide divides the current HeatFlux by the given HeatFlux and returns the result.
// The result is a new HeatFlux instance with the quotient of the values.
//
// Parameters:
//    other: The HeatFlux to divide the current HeatFlux by.
//
// Returns:
//    *HeatFlux: A new HeatFlux instance representing the quotient of both HeatFlux.
func (a *HeatFlux) Divide(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value / other.BaseValue()}
}