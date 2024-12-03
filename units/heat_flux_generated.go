// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// HeatFluxUnits enumeration
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

// HeatFluxDto represents an HeatFlux
type HeatFluxDto struct {
	Value float64
	Unit  HeatFluxUnits
}

// HeatFluxDtoFactory struct to group related functions
type HeatFluxDtoFactory struct{}

func (udf HeatFluxDtoFactory) FromJSON(data []byte) (*HeatFluxDto, error) {
	a := HeatFluxDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a HeatFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// HeatFlux struct
type HeatFlux struct {
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

// HeatFluxFactory struct to group related functions
type HeatFluxFactory struct{}

func (uf HeatFluxFactory) CreateHeatFlux(value float64, unit HeatFluxUnits) (*HeatFlux, error) {
	return newHeatFlux(value, unit)
}

func (uf HeatFluxFactory) FromDto(dto HeatFluxDto) (*HeatFlux, error) {
	return newHeatFlux(dto.Value, dto.Unit)
}

func (uf HeatFluxFactory) FromDtoJSON(data []byte) (*HeatFlux, error) {
	unitDto, err := HeatFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return HeatFluxFactory{}.FromDto(*unitDto)
}


// FromWattPerSquareMeter creates a new HeatFlux instance from WattPerSquareMeter.
func (uf HeatFluxFactory) FromWattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareMeter)
}

// FromWattPerSquareInch creates a new HeatFlux instance from WattPerSquareInch.
func (uf HeatFluxFactory) FromWattsPerSquareInch(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareInch)
}

// FromWattPerSquareFoot creates a new HeatFlux instance from WattPerSquareFoot.
func (uf HeatFluxFactory) FromWattsPerSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxWattPerSquareFoot)
}

// FromBtuPerSecondSquareInch creates a new HeatFlux instance from BtuPerSecondSquareInch.
func (uf HeatFluxFactory) FromBtusPerSecondSquareInch(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerSecondSquareInch)
}

// FromBtuPerSecondSquareFoot creates a new HeatFlux instance from BtuPerSecondSquareFoot.
func (uf HeatFluxFactory) FromBtusPerSecondSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerSecondSquareFoot)
}

// FromBtuPerMinuteSquareFoot creates a new HeatFlux instance from BtuPerMinuteSquareFoot.
func (uf HeatFluxFactory) FromBtusPerMinuteSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerMinuteSquareFoot)
}

// FromBtuPerHourSquareFoot creates a new HeatFlux instance from BtuPerHourSquareFoot.
func (uf HeatFluxFactory) FromBtusPerHourSquareFoot(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxBtuPerHourSquareFoot)
}

// FromCaloriePerSecondSquareCentimeter creates a new HeatFlux instance from CaloriePerSecondSquareCentimeter.
func (uf HeatFluxFactory) FromCaloriesPerSecondSquareCentimeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxCaloriePerSecondSquareCentimeter)
}

// FromKilocaloriePerHourSquareMeter creates a new HeatFlux instance from KilocaloriePerHourSquareMeter.
func (uf HeatFluxFactory) FromKilocaloriesPerHourSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxKilocaloriePerHourSquareMeter)
}

// FromPoundForcePerFootSecond creates a new HeatFlux instance from PoundForcePerFootSecond.
func (uf HeatFluxFactory) FromPoundsForcePerFootSecond(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxPoundForcePerFootSecond)
}

// FromPoundPerSecondCubed creates a new HeatFlux instance from PoundPerSecondCubed.
func (uf HeatFluxFactory) FromPoundsPerSecondCubed(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxPoundPerSecondCubed)
}

// FromNanowattPerSquareMeter creates a new HeatFlux instance from NanowattPerSquareMeter.
func (uf HeatFluxFactory) FromNanowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxNanowattPerSquareMeter)
}

// FromMicrowattPerSquareMeter creates a new HeatFlux instance from MicrowattPerSquareMeter.
func (uf HeatFluxFactory) FromMicrowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxMicrowattPerSquareMeter)
}

// FromMilliwattPerSquareMeter creates a new HeatFlux instance from MilliwattPerSquareMeter.
func (uf HeatFluxFactory) FromMilliwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxMilliwattPerSquareMeter)
}

// FromCentiwattPerSquareMeter creates a new HeatFlux instance from CentiwattPerSquareMeter.
func (uf HeatFluxFactory) FromCentiwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxCentiwattPerSquareMeter)
}

// FromDeciwattPerSquareMeter creates a new HeatFlux instance from DeciwattPerSquareMeter.
func (uf HeatFluxFactory) FromDeciwattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxDeciwattPerSquareMeter)
}

// FromKilowattPerSquareMeter creates a new HeatFlux instance from KilowattPerSquareMeter.
func (uf HeatFluxFactory) FromKilowattsPerSquareMeter(value float64) (*HeatFlux, error) {
	return newHeatFlux(value, HeatFluxKilowattPerSquareMeter)
}

// FromKilocaloriePerSecondSquareCentimeter creates a new HeatFlux instance from KilocaloriePerSecondSquareCentimeter.
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

// BaseValue returns the base value of HeatFlux in WattPerSquareMeter.
func (a *HeatFlux) BaseValue() float64 {
	return a.value
}


// WattPerSquareMeter returns the value in WattPerSquareMeter.
func (a *HeatFlux) WattsPerSquareMeter() float64 {
	if a.watts_per_square_meterLazy != nil {
		return *a.watts_per_square_meterLazy
	}
	watts_per_square_meter := a.convertFromBase(HeatFluxWattPerSquareMeter)
	a.watts_per_square_meterLazy = &watts_per_square_meter
	return watts_per_square_meter
}

// WattPerSquareInch returns the value in WattPerSquareInch.
func (a *HeatFlux) WattsPerSquareInch() float64 {
	if a.watts_per_square_inchLazy != nil {
		return *a.watts_per_square_inchLazy
	}
	watts_per_square_inch := a.convertFromBase(HeatFluxWattPerSquareInch)
	a.watts_per_square_inchLazy = &watts_per_square_inch
	return watts_per_square_inch
}

// WattPerSquareFoot returns the value in WattPerSquareFoot.
func (a *HeatFlux) WattsPerSquareFoot() float64 {
	if a.watts_per_square_footLazy != nil {
		return *a.watts_per_square_footLazy
	}
	watts_per_square_foot := a.convertFromBase(HeatFluxWattPerSquareFoot)
	a.watts_per_square_footLazy = &watts_per_square_foot
	return watts_per_square_foot
}

// BtuPerSecondSquareInch returns the value in BtuPerSecondSquareInch.
func (a *HeatFlux) BtusPerSecondSquareInch() float64 {
	if a.btus_per_second_square_inchLazy != nil {
		return *a.btus_per_second_square_inchLazy
	}
	btus_per_second_square_inch := a.convertFromBase(HeatFluxBtuPerSecondSquareInch)
	a.btus_per_second_square_inchLazy = &btus_per_second_square_inch
	return btus_per_second_square_inch
}

// BtuPerSecondSquareFoot returns the value in BtuPerSecondSquareFoot.
func (a *HeatFlux) BtusPerSecondSquareFoot() float64 {
	if a.btus_per_second_square_footLazy != nil {
		return *a.btus_per_second_square_footLazy
	}
	btus_per_second_square_foot := a.convertFromBase(HeatFluxBtuPerSecondSquareFoot)
	a.btus_per_second_square_footLazy = &btus_per_second_square_foot
	return btus_per_second_square_foot
}

// BtuPerMinuteSquareFoot returns the value in BtuPerMinuteSquareFoot.
func (a *HeatFlux) BtusPerMinuteSquareFoot() float64 {
	if a.btus_per_minute_square_footLazy != nil {
		return *a.btus_per_minute_square_footLazy
	}
	btus_per_minute_square_foot := a.convertFromBase(HeatFluxBtuPerMinuteSquareFoot)
	a.btus_per_minute_square_footLazy = &btus_per_minute_square_foot
	return btus_per_minute_square_foot
}

// BtuPerHourSquareFoot returns the value in BtuPerHourSquareFoot.
func (a *HeatFlux) BtusPerHourSquareFoot() float64 {
	if a.btus_per_hour_square_footLazy != nil {
		return *a.btus_per_hour_square_footLazy
	}
	btus_per_hour_square_foot := a.convertFromBase(HeatFluxBtuPerHourSquareFoot)
	a.btus_per_hour_square_footLazy = &btus_per_hour_square_foot
	return btus_per_hour_square_foot
}

// CaloriePerSecondSquareCentimeter returns the value in CaloriePerSecondSquareCentimeter.
func (a *HeatFlux) CaloriesPerSecondSquareCentimeter() float64 {
	if a.calories_per_second_square_centimeterLazy != nil {
		return *a.calories_per_second_square_centimeterLazy
	}
	calories_per_second_square_centimeter := a.convertFromBase(HeatFluxCaloriePerSecondSquareCentimeter)
	a.calories_per_second_square_centimeterLazy = &calories_per_second_square_centimeter
	return calories_per_second_square_centimeter
}

// KilocaloriePerHourSquareMeter returns the value in KilocaloriePerHourSquareMeter.
func (a *HeatFlux) KilocaloriesPerHourSquareMeter() float64 {
	if a.kilocalories_per_hour_square_meterLazy != nil {
		return *a.kilocalories_per_hour_square_meterLazy
	}
	kilocalories_per_hour_square_meter := a.convertFromBase(HeatFluxKilocaloriePerHourSquareMeter)
	a.kilocalories_per_hour_square_meterLazy = &kilocalories_per_hour_square_meter
	return kilocalories_per_hour_square_meter
}

// PoundForcePerFootSecond returns the value in PoundForcePerFootSecond.
func (a *HeatFlux) PoundsForcePerFootSecond() float64 {
	if a.pounds_force_per_foot_secondLazy != nil {
		return *a.pounds_force_per_foot_secondLazy
	}
	pounds_force_per_foot_second := a.convertFromBase(HeatFluxPoundForcePerFootSecond)
	a.pounds_force_per_foot_secondLazy = &pounds_force_per_foot_second
	return pounds_force_per_foot_second
}

// PoundPerSecondCubed returns the value in PoundPerSecondCubed.
func (a *HeatFlux) PoundsPerSecondCubed() float64 {
	if a.pounds_per_second_cubedLazy != nil {
		return *a.pounds_per_second_cubedLazy
	}
	pounds_per_second_cubed := a.convertFromBase(HeatFluxPoundPerSecondCubed)
	a.pounds_per_second_cubedLazy = &pounds_per_second_cubed
	return pounds_per_second_cubed
}

// NanowattPerSquareMeter returns the value in NanowattPerSquareMeter.
func (a *HeatFlux) NanowattsPerSquareMeter() float64 {
	if a.nanowatts_per_square_meterLazy != nil {
		return *a.nanowatts_per_square_meterLazy
	}
	nanowatts_per_square_meter := a.convertFromBase(HeatFluxNanowattPerSquareMeter)
	a.nanowatts_per_square_meterLazy = &nanowatts_per_square_meter
	return nanowatts_per_square_meter
}

// MicrowattPerSquareMeter returns the value in MicrowattPerSquareMeter.
func (a *HeatFlux) MicrowattsPerSquareMeter() float64 {
	if a.microwatts_per_square_meterLazy != nil {
		return *a.microwatts_per_square_meterLazy
	}
	microwatts_per_square_meter := a.convertFromBase(HeatFluxMicrowattPerSquareMeter)
	a.microwatts_per_square_meterLazy = &microwatts_per_square_meter
	return microwatts_per_square_meter
}

// MilliwattPerSquareMeter returns the value in MilliwattPerSquareMeter.
func (a *HeatFlux) MilliwattsPerSquareMeter() float64 {
	if a.milliwatts_per_square_meterLazy != nil {
		return *a.milliwatts_per_square_meterLazy
	}
	milliwatts_per_square_meter := a.convertFromBase(HeatFluxMilliwattPerSquareMeter)
	a.milliwatts_per_square_meterLazy = &milliwatts_per_square_meter
	return milliwatts_per_square_meter
}

// CentiwattPerSquareMeter returns the value in CentiwattPerSquareMeter.
func (a *HeatFlux) CentiwattsPerSquareMeter() float64 {
	if a.centiwatts_per_square_meterLazy != nil {
		return *a.centiwatts_per_square_meterLazy
	}
	centiwatts_per_square_meter := a.convertFromBase(HeatFluxCentiwattPerSquareMeter)
	a.centiwatts_per_square_meterLazy = &centiwatts_per_square_meter
	return centiwatts_per_square_meter
}

// DeciwattPerSquareMeter returns the value in DeciwattPerSquareMeter.
func (a *HeatFlux) DeciwattsPerSquareMeter() float64 {
	if a.deciwatts_per_square_meterLazy != nil {
		return *a.deciwatts_per_square_meterLazy
	}
	deciwatts_per_square_meter := a.convertFromBase(HeatFluxDeciwattPerSquareMeter)
	a.deciwatts_per_square_meterLazy = &deciwatts_per_square_meter
	return deciwatts_per_square_meter
}

// KilowattPerSquareMeter returns the value in KilowattPerSquareMeter.
func (a *HeatFlux) KilowattsPerSquareMeter() float64 {
	if a.kilowatts_per_square_meterLazy != nil {
		return *a.kilowatts_per_square_meterLazy
	}
	kilowatts_per_square_meter := a.convertFromBase(HeatFluxKilowattPerSquareMeter)
	a.kilowatts_per_square_meterLazy = &kilowatts_per_square_meter
	return kilowatts_per_square_meter
}

// KilocaloriePerSecondSquareCentimeter returns the value in KilocaloriePerSecondSquareCentimeter.
func (a *HeatFlux) KilocaloriesPerSecondSquareCentimeter() float64 {
	if a.kilocalories_per_second_square_centimeterLazy != nil {
		return *a.kilocalories_per_second_square_centimeterLazy
	}
	kilocalories_per_second_square_centimeter := a.convertFromBase(HeatFluxKilocaloriePerSecondSquareCentimeter)
	a.kilocalories_per_second_square_centimeterLazy = &kilocalories_per_second_square_centimeter
	return kilocalories_per_second_square_centimeter
}


// ToDto creates an HeatFluxDto representation.
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

// ToDtoJSON creates an HeatFluxDto representation.
func (a *HeatFlux) ToDtoJSON(holdInUnit *HeatFluxUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts HeatFlux to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a HeatFlux) String() string {
	return a.ToString(HeatFluxWattPerSquareMeter, 2)
}

// ToString formats the HeatFlux to string.
// fractionalDigits -1 for not mention
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

// Check if the given HeatFlux are equals to the current HeatFlux
func (a *HeatFlux) Equals(other *HeatFlux) bool {
	return a.value == other.BaseValue()
}

// Check if the given HeatFlux are equals to the current HeatFlux
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

// Add the given HeatFlux to the current HeatFlux.
func (a *HeatFlux) Add(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value + other.BaseValue()}
}

// Subtract the given HeatFlux to the current HeatFlux.
func (a *HeatFlux) Subtract(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value - other.BaseValue()}
}

// Multiply the given HeatFlux to the current HeatFlux.
func (a *HeatFlux) Multiply(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value * other.BaseValue()}
}

// Divide the given HeatFlux to the current HeatFlux.
func (a *HeatFlux) Divide(other *HeatFlux) *HeatFlux {
	return &HeatFlux{value: a.value / other.BaseValue()}
}