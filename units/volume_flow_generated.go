// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeFlowUnits defines various units of VolumeFlow.
type VolumeFlowUnits string

const (
    
        // 
        VolumeFlowCubicMeterPerSecond VolumeFlowUnits = "CubicMeterPerSecond"
        // 
        VolumeFlowCubicMeterPerMinute VolumeFlowUnits = "CubicMeterPerMinute"
        // 
        VolumeFlowCubicMeterPerHour VolumeFlowUnits = "CubicMeterPerHour"
        // 
        VolumeFlowCubicMeterPerDay VolumeFlowUnits = "CubicMeterPerDay"
        // 
        VolumeFlowCubicFootPerSecond VolumeFlowUnits = "CubicFootPerSecond"
        // 
        VolumeFlowCubicFootPerMinute VolumeFlowUnits = "CubicFootPerMinute"
        // 
        VolumeFlowCubicFootPerHour VolumeFlowUnits = "CubicFootPerHour"
        // 
        VolumeFlowCubicYardPerSecond VolumeFlowUnits = "CubicYardPerSecond"
        // 
        VolumeFlowCubicYardPerMinute VolumeFlowUnits = "CubicYardPerMinute"
        // 
        VolumeFlowCubicYardPerHour VolumeFlowUnits = "CubicYardPerHour"
        // 
        VolumeFlowCubicYardPerDay VolumeFlowUnits = "CubicYardPerDay"
        // 
        VolumeFlowMillionUsGallonPerDay VolumeFlowUnits = "MillionUsGallonPerDay"
        // 
        VolumeFlowUsGallonPerDay VolumeFlowUnits = "UsGallonPerDay"
        // 
        VolumeFlowLiterPerSecond VolumeFlowUnits = "LiterPerSecond"
        // 
        VolumeFlowLiterPerMinute VolumeFlowUnits = "LiterPerMinute"
        // 
        VolumeFlowLiterPerHour VolumeFlowUnits = "LiterPerHour"
        // 
        VolumeFlowLiterPerDay VolumeFlowUnits = "LiterPerDay"
        // 
        VolumeFlowUsGallonPerSecond VolumeFlowUnits = "UsGallonPerSecond"
        // 
        VolumeFlowUsGallonPerMinute VolumeFlowUnits = "UsGallonPerMinute"
        // 
        VolumeFlowUkGallonPerDay VolumeFlowUnits = "UkGallonPerDay"
        // 
        VolumeFlowUkGallonPerHour VolumeFlowUnits = "UkGallonPerHour"
        // 
        VolumeFlowUkGallonPerMinute VolumeFlowUnits = "UkGallonPerMinute"
        // 
        VolumeFlowUkGallonPerSecond VolumeFlowUnits = "UkGallonPerSecond"
        // 
        VolumeFlowKilousGallonPerMinute VolumeFlowUnits = "KilousGallonPerMinute"
        // 
        VolumeFlowUsGallonPerHour VolumeFlowUnits = "UsGallonPerHour"
        // 
        VolumeFlowCubicDecimeterPerMinute VolumeFlowUnits = "CubicDecimeterPerMinute"
        // 
        VolumeFlowOilBarrelPerDay VolumeFlowUnits = "OilBarrelPerDay"
        // 
        VolumeFlowOilBarrelPerMinute VolumeFlowUnits = "OilBarrelPerMinute"
        // 
        VolumeFlowOilBarrelPerHour VolumeFlowUnits = "OilBarrelPerHour"
        // 
        VolumeFlowOilBarrelPerSecond VolumeFlowUnits = "OilBarrelPerSecond"
        // 
        VolumeFlowCubicMillimeterPerSecond VolumeFlowUnits = "CubicMillimeterPerSecond"
        // 
        VolumeFlowAcreFootPerSecond VolumeFlowUnits = "AcreFootPerSecond"
        // 
        VolumeFlowAcreFootPerMinute VolumeFlowUnits = "AcreFootPerMinute"
        // 
        VolumeFlowAcreFootPerHour VolumeFlowUnits = "AcreFootPerHour"
        // 
        VolumeFlowAcreFootPerDay VolumeFlowUnits = "AcreFootPerDay"
        // 
        VolumeFlowCubicCentimeterPerMinute VolumeFlowUnits = "CubicCentimeterPerMinute"
        // 
        VolumeFlowMegausGallonPerDay VolumeFlowUnits = "MegausGallonPerDay"
        // 
        VolumeFlowNanoliterPerSecond VolumeFlowUnits = "NanoliterPerSecond"
        // 
        VolumeFlowMicroliterPerSecond VolumeFlowUnits = "MicroliterPerSecond"
        // 
        VolumeFlowMilliliterPerSecond VolumeFlowUnits = "MilliliterPerSecond"
        // 
        VolumeFlowCentiliterPerSecond VolumeFlowUnits = "CentiliterPerSecond"
        // 
        VolumeFlowDeciliterPerSecond VolumeFlowUnits = "DeciliterPerSecond"
        // 
        VolumeFlowDecaliterPerSecond VolumeFlowUnits = "DecaliterPerSecond"
        // 
        VolumeFlowHectoliterPerSecond VolumeFlowUnits = "HectoliterPerSecond"
        // 
        VolumeFlowKiloliterPerSecond VolumeFlowUnits = "KiloliterPerSecond"
        // 
        VolumeFlowMegaliterPerSecond VolumeFlowUnits = "MegaliterPerSecond"
        // 
        VolumeFlowNanoliterPerMinute VolumeFlowUnits = "NanoliterPerMinute"
        // 
        VolumeFlowMicroliterPerMinute VolumeFlowUnits = "MicroliterPerMinute"
        // 
        VolumeFlowMilliliterPerMinute VolumeFlowUnits = "MilliliterPerMinute"
        // 
        VolumeFlowCentiliterPerMinute VolumeFlowUnits = "CentiliterPerMinute"
        // 
        VolumeFlowDeciliterPerMinute VolumeFlowUnits = "DeciliterPerMinute"
        // 
        VolumeFlowDecaliterPerMinute VolumeFlowUnits = "DecaliterPerMinute"
        // 
        VolumeFlowHectoliterPerMinute VolumeFlowUnits = "HectoliterPerMinute"
        // 
        VolumeFlowKiloliterPerMinute VolumeFlowUnits = "KiloliterPerMinute"
        // 
        VolumeFlowMegaliterPerMinute VolumeFlowUnits = "MegaliterPerMinute"
        // 
        VolumeFlowNanoliterPerHour VolumeFlowUnits = "NanoliterPerHour"
        // 
        VolumeFlowMicroliterPerHour VolumeFlowUnits = "MicroliterPerHour"
        // 
        VolumeFlowMilliliterPerHour VolumeFlowUnits = "MilliliterPerHour"
        // 
        VolumeFlowCentiliterPerHour VolumeFlowUnits = "CentiliterPerHour"
        // 
        VolumeFlowDeciliterPerHour VolumeFlowUnits = "DeciliterPerHour"
        // 
        VolumeFlowDecaliterPerHour VolumeFlowUnits = "DecaliterPerHour"
        // 
        VolumeFlowHectoliterPerHour VolumeFlowUnits = "HectoliterPerHour"
        // 
        VolumeFlowKiloliterPerHour VolumeFlowUnits = "KiloliterPerHour"
        // 
        VolumeFlowMegaliterPerHour VolumeFlowUnits = "MegaliterPerHour"
        // 
        VolumeFlowNanoliterPerDay VolumeFlowUnits = "NanoliterPerDay"
        // 
        VolumeFlowMicroliterPerDay VolumeFlowUnits = "MicroliterPerDay"
        // 
        VolumeFlowMilliliterPerDay VolumeFlowUnits = "MilliliterPerDay"
        // 
        VolumeFlowCentiliterPerDay VolumeFlowUnits = "CentiliterPerDay"
        // 
        VolumeFlowDeciliterPerDay VolumeFlowUnits = "DeciliterPerDay"
        // 
        VolumeFlowDecaliterPerDay VolumeFlowUnits = "DecaliterPerDay"
        // 
        VolumeFlowHectoliterPerDay VolumeFlowUnits = "HectoliterPerDay"
        // 
        VolumeFlowKiloliterPerDay VolumeFlowUnits = "KiloliterPerDay"
        // 
        VolumeFlowMegaliterPerDay VolumeFlowUnits = "MegaliterPerDay"
        // 
        VolumeFlowMegaukGallonPerDay VolumeFlowUnits = "MegaukGallonPerDay"
        // 
        VolumeFlowMegaukGallonPerSecond VolumeFlowUnits = "MegaukGallonPerSecond"
)

var internalVolumeFlowUnitsMap = map[VolumeFlowUnits]bool{
	
	VolumeFlowCubicMeterPerSecond: true,
	VolumeFlowCubicMeterPerMinute: true,
	VolumeFlowCubicMeterPerHour: true,
	VolumeFlowCubicMeterPerDay: true,
	VolumeFlowCubicFootPerSecond: true,
	VolumeFlowCubicFootPerMinute: true,
	VolumeFlowCubicFootPerHour: true,
	VolumeFlowCubicYardPerSecond: true,
	VolumeFlowCubicYardPerMinute: true,
	VolumeFlowCubicYardPerHour: true,
	VolumeFlowCubicYardPerDay: true,
	VolumeFlowMillionUsGallonPerDay: true,
	VolumeFlowUsGallonPerDay: true,
	VolumeFlowLiterPerSecond: true,
	VolumeFlowLiterPerMinute: true,
	VolumeFlowLiterPerHour: true,
	VolumeFlowLiterPerDay: true,
	VolumeFlowUsGallonPerSecond: true,
	VolumeFlowUsGallonPerMinute: true,
	VolumeFlowUkGallonPerDay: true,
	VolumeFlowUkGallonPerHour: true,
	VolumeFlowUkGallonPerMinute: true,
	VolumeFlowUkGallonPerSecond: true,
	VolumeFlowKilousGallonPerMinute: true,
	VolumeFlowUsGallonPerHour: true,
	VolumeFlowCubicDecimeterPerMinute: true,
	VolumeFlowOilBarrelPerDay: true,
	VolumeFlowOilBarrelPerMinute: true,
	VolumeFlowOilBarrelPerHour: true,
	VolumeFlowOilBarrelPerSecond: true,
	VolumeFlowCubicMillimeterPerSecond: true,
	VolumeFlowAcreFootPerSecond: true,
	VolumeFlowAcreFootPerMinute: true,
	VolumeFlowAcreFootPerHour: true,
	VolumeFlowAcreFootPerDay: true,
	VolumeFlowCubicCentimeterPerMinute: true,
	VolumeFlowMegausGallonPerDay: true,
	VolumeFlowNanoliterPerSecond: true,
	VolumeFlowMicroliterPerSecond: true,
	VolumeFlowMilliliterPerSecond: true,
	VolumeFlowCentiliterPerSecond: true,
	VolumeFlowDeciliterPerSecond: true,
	VolumeFlowDecaliterPerSecond: true,
	VolumeFlowHectoliterPerSecond: true,
	VolumeFlowKiloliterPerSecond: true,
	VolumeFlowMegaliterPerSecond: true,
	VolumeFlowNanoliterPerMinute: true,
	VolumeFlowMicroliterPerMinute: true,
	VolumeFlowMilliliterPerMinute: true,
	VolumeFlowCentiliterPerMinute: true,
	VolumeFlowDeciliterPerMinute: true,
	VolumeFlowDecaliterPerMinute: true,
	VolumeFlowHectoliterPerMinute: true,
	VolumeFlowKiloliterPerMinute: true,
	VolumeFlowMegaliterPerMinute: true,
	VolumeFlowNanoliterPerHour: true,
	VolumeFlowMicroliterPerHour: true,
	VolumeFlowMilliliterPerHour: true,
	VolumeFlowCentiliterPerHour: true,
	VolumeFlowDeciliterPerHour: true,
	VolumeFlowDecaliterPerHour: true,
	VolumeFlowHectoliterPerHour: true,
	VolumeFlowKiloliterPerHour: true,
	VolumeFlowMegaliterPerHour: true,
	VolumeFlowNanoliterPerDay: true,
	VolumeFlowMicroliterPerDay: true,
	VolumeFlowMilliliterPerDay: true,
	VolumeFlowCentiliterPerDay: true,
	VolumeFlowDeciliterPerDay: true,
	VolumeFlowDecaliterPerDay: true,
	VolumeFlowHectoliterPerDay: true,
	VolumeFlowKiloliterPerDay: true,
	VolumeFlowMegaliterPerDay: true,
	VolumeFlowMegaukGallonPerDay: true,
	VolumeFlowMegaukGallonPerSecond: true,
}

// VolumeFlowDto represents a VolumeFlow measurement with a numerical value and its corresponding unit.
type VolumeFlowDto struct {
    // Value is the numerical representation of the VolumeFlow.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the VolumeFlow, as defined in the VolumeFlowUnits enumeration.
	Unit  VolumeFlowUnits `json:"unit" validate:"required,oneof=CubicMeterPerSecond CubicMeterPerMinute CubicMeterPerHour CubicMeterPerDay CubicFootPerSecond CubicFootPerMinute CubicFootPerHour CubicYardPerSecond CubicYardPerMinute CubicYardPerHour CubicYardPerDay MillionUsGallonPerDay UsGallonPerDay LiterPerSecond LiterPerMinute LiterPerHour LiterPerDay UsGallonPerSecond UsGallonPerMinute UkGallonPerDay UkGallonPerHour UkGallonPerMinute UkGallonPerSecond KilousGallonPerMinute UsGallonPerHour CubicDecimeterPerMinute OilBarrelPerDay OilBarrelPerMinute OilBarrelPerHour OilBarrelPerSecond CubicMillimeterPerSecond AcreFootPerSecond AcreFootPerMinute AcreFootPerHour AcreFootPerDay CubicCentimeterPerMinute MegausGallonPerDay NanoliterPerSecond MicroliterPerSecond MilliliterPerSecond CentiliterPerSecond DeciliterPerSecond DecaliterPerSecond HectoliterPerSecond KiloliterPerSecond MegaliterPerSecond NanoliterPerMinute MicroliterPerMinute MilliliterPerMinute CentiliterPerMinute DeciliterPerMinute DecaliterPerMinute HectoliterPerMinute KiloliterPerMinute MegaliterPerMinute NanoliterPerHour MicroliterPerHour MilliliterPerHour CentiliterPerHour DeciliterPerHour DecaliterPerHour HectoliterPerHour KiloliterPerHour MegaliterPerHour NanoliterPerDay MicroliterPerDay MilliliterPerDay CentiliterPerDay DeciliterPerDay DecaliterPerDay HectoliterPerDay KiloliterPerDay MegaliterPerDay MegaukGallonPerDay MegaukGallonPerSecond"`
}

// VolumeFlowDtoFactory groups methods for creating and serializing VolumeFlowDto objects.
type VolumeFlowDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumeFlowDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumeFlowDtoFactory) FromJSON(data []byte) (*VolumeFlowDto, error) {
	a := VolumeFlowDto{}

    // Parse JSON into VolumeFlowDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a VolumeFlowDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumeFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// VolumeFlow represents a measurement in a of VolumeFlow.
//
// In physics and engineering, in particular fluid dynamics and hydrometry, the volumetric flow rate, (also known as volume flow rate, rate of fluid flow or volume velocity) is the volume of fluid which passes through a given surface per unit time. The SI unit is m³/s (cubic meters per second). In US Customary Units and British Imperial Units, volumetric flow rate is often expressed as ft³/s (cubic feet per second). It is usually represented by the symbol Q.
type VolumeFlow struct {
	// value is the base measurement stored internally.
	value       float64
    
    cubic_meters_per_secondLazy *float64 
    cubic_meters_per_minuteLazy *float64 
    cubic_meters_per_hourLazy *float64 
    cubic_meters_per_dayLazy *float64 
    cubic_feet_per_secondLazy *float64 
    cubic_feet_per_minuteLazy *float64 
    cubic_feet_per_hourLazy *float64 
    cubic_yards_per_secondLazy *float64 
    cubic_yards_per_minuteLazy *float64 
    cubic_yards_per_hourLazy *float64 
    cubic_yards_per_dayLazy *float64 
    million_us_gallons_per_dayLazy *float64 
    us_gallons_per_dayLazy *float64 
    liters_per_secondLazy *float64 
    liters_per_minuteLazy *float64 
    liters_per_hourLazy *float64 
    liters_per_dayLazy *float64 
    us_gallons_per_secondLazy *float64 
    us_gallons_per_minuteLazy *float64 
    uk_gallons_per_dayLazy *float64 
    uk_gallons_per_hourLazy *float64 
    uk_gallons_per_minuteLazy *float64 
    uk_gallons_per_secondLazy *float64 
    kilous_gallons_per_minuteLazy *float64 
    us_gallons_per_hourLazy *float64 
    cubic_decimeters_per_minuteLazy *float64 
    oil_barrels_per_dayLazy *float64 
    oil_barrels_per_minuteLazy *float64 
    oil_barrels_per_hourLazy *float64 
    oil_barrels_per_secondLazy *float64 
    cubic_millimeters_per_secondLazy *float64 
    acre_feet_per_secondLazy *float64 
    acre_feet_per_minuteLazy *float64 
    acre_feet_per_hourLazy *float64 
    acre_feet_per_dayLazy *float64 
    cubic_centimeters_per_minuteLazy *float64 
    megaus_gallons_per_dayLazy *float64 
    nanoliters_per_secondLazy *float64 
    microliters_per_secondLazy *float64 
    milliliters_per_secondLazy *float64 
    centiliters_per_secondLazy *float64 
    deciliters_per_secondLazy *float64 
    decaliters_per_secondLazy *float64 
    hectoliters_per_secondLazy *float64 
    kiloliters_per_secondLazy *float64 
    megaliters_per_secondLazy *float64 
    nanoliters_per_minuteLazy *float64 
    microliters_per_minuteLazy *float64 
    milliliters_per_minuteLazy *float64 
    centiliters_per_minuteLazy *float64 
    deciliters_per_minuteLazy *float64 
    decaliters_per_minuteLazy *float64 
    hectoliters_per_minuteLazy *float64 
    kiloliters_per_minuteLazy *float64 
    megaliters_per_minuteLazy *float64 
    nanoliters_per_hourLazy *float64 
    microliters_per_hourLazy *float64 
    milliliters_per_hourLazy *float64 
    centiliters_per_hourLazy *float64 
    deciliters_per_hourLazy *float64 
    decaliters_per_hourLazy *float64 
    hectoliters_per_hourLazy *float64 
    kiloliters_per_hourLazy *float64 
    megaliters_per_hourLazy *float64 
    nanoliters_per_dayLazy *float64 
    microliters_per_dayLazy *float64 
    milliliters_per_dayLazy *float64 
    centiliters_per_dayLazy *float64 
    deciliters_per_dayLazy *float64 
    decaliters_per_dayLazy *float64 
    hectoliters_per_dayLazy *float64 
    kiloliters_per_dayLazy *float64 
    megaliters_per_dayLazy *float64 
    megauk_gallons_per_dayLazy *float64 
    megauk_gallons_per_secondLazy *float64 
}

// VolumeFlowFactory groups methods for creating VolumeFlow instances.
type VolumeFlowFactory struct{}

// CreateVolumeFlow creates a new VolumeFlow instance from the given value and unit.
func (uf VolumeFlowFactory) CreateVolumeFlow(value float64, unit VolumeFlowUnits) (*VolumeFlow, error) {
	return newVolumeFlow(value, unit)
}

// FromDto converts a VolumeFlowDto to a VolumeFlow instance.
func (uf VolumeFlowFactory) FromDto(dto VolumeFlowDto) (*VolumeFlow, error) {
	return newVolumeFlow(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VolumeFlow instance.
func (uf VolumeFlowFactory) FromDtoJSON(data []byte) (*VolumeFlow, error) {
	unitDto, err := VolumeFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumeFlowDto from JSON: %w", err)
	}
	return VolumeFlowFactory{}.FromDto(*unitDto)
}


// FromCubicMetersPerSecond creates a new VolumeFlow instance from a value in CubicMetersPerSecond.
func (uf VolumeFlowFactory) FromCubicMetersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerSecond)
}

// FromCubicMetersPerMinute creates a new VolumeFlow instance from a value in CubicMetersPerMinute.
func (uf VolumeFlowFactory) FromCubicMetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerMinute)
}

// FromCubicMetersPerHour creates a new VolumeFlow instance from a value in CubicMetersPerHour.
func (uf VolumeFlowFactory) FromCubicMetersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerHour)
}

// FromCubicMetersPerDay creates a new VolumeFlow instance from a value in CubicMetersPerDay.
func (uf VolumeFlowFactory) FromCubicMetersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerDay)
}

// FromCubicFeetPerSecond creates a new VolumeFlow instance from a value in CubicFeetPerSecond.
func (uf VolumeFlowFactory) FromCubicFeetPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerSecond)
}

// FromCubicFeetPerMinute creates a new VolumeFlow instance from a value in CubicFeetPerMinute.
func (uf VolumeFlowFactory) FromCubicFeetPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerMinute)
}

// FromCubicFeetPerHour creates a new VolumeFlow instance from a value in CubicFeetPerHour.
func (uf VolumeFlowFactory) FromCubicFeetPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerHour)
}

// FromCubicYardsPerSecond creates a new VolumeFlow instance from a value in CubicYardsPerSecond.
func (uf VolumeFlowFactory) FromCubicYardsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerSecond)
}

// FromCubicYardsPerMinute creates a new VolumeFlow instance from a value in CubicYardsPerMinute.
func (uf VolumeFlowFactory) FromCubicYardsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerMinute)
}

// FromCubicYardsPerHour creates a new VolumeFlow instance from a value in CubicYardsPerHour.
func (uf VolumeFlowFactory) FromCubicYardsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerHour)
}

// FromCubicYardsPerDay creates a new VolumeFlow instance from a value in CubicYardsPerDay.
func (uf VolumeFlowFactory) FromCubicYardsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerDay)
}

// FromMillionUsGallonsPerDay creates a new VolumeFlow instance from a value in MillionUsGallonsPerDay.
func (uf VolumeFlowFactory) FromMillionUsGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMillionUsGallonPerDay)
}

// FromUsGallonsPerDay creates a new VolumeFlow instance from a value in UsGallonsPerDay.
func (uf VolumeFlowFactory) FromUsGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerDay)
}

// FromLitersPerSecond creates a new VolumeFlow instance from a value in LitersPerSecond.
func (uf VolumeFlowFactory) FromLitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerSecond)
}

// FromLitersPerMinute creates a new VolumeFlow instance from a value in LitersPerMinute.
func (uf VolumeFlowFactory) FromLitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerMinute)
}

// FromLitersPerHour creates a new VolumeFlow instance from a value in LitersPerHour.
func (uf VolumeFlowFactory) FromLitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerHour)
}

// FromLitersPerDay creates a new VolumeFlow instance from a value in LitersPerDay.
func (uf VolumeFlowFactory) FromLitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerDay)
}

// FromUsGallonsPerSecond creates a new VolumeFlow instance from a value in UsGallonsPerSecond.
func (uf VolumeFlowFactory) FromUsGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerSecond)
}

// FromUsGallonsPerMinute creates a new VolumeFlow instance from a value in UsGallonsPerMinute.
func (uf VolumeFlowFactory) FromUsGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerMinute)
}

// FromUkGallonsPerDay creates a new VolumeFlow instance from a value in UkGallonsPerDay.
func (uf VolumeFlowFactory) FromUkGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerDay)
}

// FromUkGallonsPerHour creates a new VolumeFlow instance from a value in UkGallonsPerHour.
func (uf VolumeFlowFactory) FromUkGallonsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerHour)
}

// FromUkGallonsPerMinute creates a new VolumeFlow instance from a value in UkGallonsPerMinute.
func (uf VolumeFlowFactory) FromUkGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerMinute)
}

// FromUkGallonsPerSecond creates a new VolumeFlow instance from a value in UkGallonsPerSecond.
func (uf VolumeFlowFactory) FromUkGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerSecond)
}

// FromKilousGallonsPerMinute creates a new VolumeFlow instance from a value in KilousGallonsPerMinute.
func (uf VolumeFlowFactory) FromKilousGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKilousGallonPerMinute)
}

// FromUsGallonsPerHour creates a new VolumeFlow instance from a value in UsGallonsPerHour.
func (uf VolumeFlowFactory) FromUsGallonsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerHour)
}

// FromCubicDecimetersPerMinute creates a new VolumeFlow instance from a value in CubicDecimetersPerMinute.
func (uf VolumeFlowFactory) FromCubicDecimetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicDecimeterPerMinute)
}

// FromOilBarrelsPerDay creates a new VolumeFlow instance from a value in OilBarrelsPerDay.
func (uf VolumeFlowFactory) FromOilBarrelsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerDay)
}

// FromOilBarrelsPerMinute creates a new VolumeFlow instance from a value in OilBarrelsPerMinute.
func (uf VolumeFlowFactory) FromOilBarrelsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerMinute)
}

// FromOilBarrelsPerHour creates a new VolumeFlow instance from a value in OilBarrelsPerHour.
func (uf VolumeFlowFactory) FromOilBarrelsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerHour)
}

// FromOilBarrelsPerSecond creates a new VolumeFlow instance from a value in OilBarrelsPerSecond.
func (uf VolumeFlowFactory) FromOilBarrelsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerSecond)
}

// FromCubicMillimetersPerSecond creates a new VolumeFlow instance from a value in CubicMillimetersPerSecond.
func (uf VolumeFlowFactory) FromCubicMillimetersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMillimeterPerSecond)
}

// FromAcreFeetPerSecond creates a new VolumeFlow instance from a value in AcreFeetPerSecond.
func (uf VolumeFlowFactory) FromAcreFeetPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerSecond)
}

// FromAcreFeetPerMinute creates a new VolumeFlow instance from a value in AcreFeetPerMinute.
func (uf VolumeFlowFactory) FromAcreFeetPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerMinute)
}

// FromAcreFeetPerHour creates a new VolumeFlow instance from a value in AcreFeetPerHour.
func (uf VolumeFlowFactory) FromAcreFeetPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerHour)
}

// FromAcreFeetPerDay creates a new VolumeFlow instance from a value in AcreFeetPerDay.
func (uf VolumeFlowFactory) FromAcreFeetPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerDay)
}

// FromCubicCentimetersPerMinute creates a new VolumeFlow instance from a value in CubicCentimetersPerMinute.
func (uf VolumeFlowFactory) FromCubicCentimetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicCentimeterPerMinute)
}

// FromMegausGallonsPerDay creates a new VolumeFlow instance from a value in MegausGallonsPerDay.
func (uf VolumeFlowFactory) FromMegausGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegausGallonPerDay)
}

// FromNanolitersPerSecond creates a new VolumeFlow instance from a value in NanolitersPerSecond.
func (uf VolumeFlowFactory) FromNanolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerSecond)
}

// FromMicrolitersPerSecond creates a new VolumeFlow instance from a value in MicrolitersPerSecond.
func (uf VolumeFlowFactory) FromMicrolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerSecond)
}

// FromMillilitersPerSecond creates a new VolumeFlow instance from a value in MillilitersPerSecond.
func (uf VolumeFlowFactory) FromMillilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerSecond)
}

// FromCentilitersPerSecond creates a new VolumeFlow instance from a value in CentilitersPerSecond.
func (uf VolumeFlowFactory) FromCentilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerSecond)
}

// FromDecilitersPerSecond creates a new VolumeFlow instance from a value in DecilitersPerSecond.
func (uf VolumeFlowFactory) FromDecilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerSecond)
}

// FromDecalitersPerSecond creates a new VolumeFlow instance from a value in DecalitersPerSecond.
func (uf VolumeFlowFactory) FromDecalitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerSecond)
}

// FromHectolitersPerSecond creates a new VolumeFlow instance from a value in HectolitersPerSecond.
func (uf VolumeFlowFactory) FromHectolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerSecond)
}

// FromKilolitersPerSecond creates a new VolumeFlow instance from a value in KilolitersPerSecond.
func (uf VolumeFlowFactory) FromKilolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerSecond)
}

// FromMegalitersPerSecond creates a new VolumeFlow instance from a value in MegalitersPerSecond.
func (uf VolumeFlowFactory) FromMegalitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerSecond)
}

// FromNanolitersPerMinute creates a new VolumeFlow instance from a value in NanolitersPerMinute.
func (uf VolumeFlowFactory) FromNanolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerMinute)
}

// FromMicrolitersPerMinute creates a new VolumeFlow instance from a value in MicrolitersPerMinute.
func (uf VolumeFlowFactory) FromMicrolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerMinute)
}

// FromMillilitersPerMinute creates a new VolumeFlow instance from a value in MillilitersPerMinute.
func (uf VolumeFlowFactory) FromMillilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerMinute)
}

// FromCentilitersPerMinute creates a new VolumeFlow instance from a value in CentilitersPerMinute.
func (uf VolumeFlowFactory) FromCentilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerMinute)
}

// FromDecilitersPerMinute creates a new VolumeFlow instance from a value in DecilitersPerMinute.
func (uf VolumeFlowFactory) FromDecilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerMinute)
}

// FromDecalitersPerMinute creates a new VolumeFlow instance from a value in DecalitersPerMinute.
func (uf VolumeFlowFactory) FromDecalitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerMinute)
}

// FromHectolitersPerMinute creates a new VolumeFlow instance from a value in HectolitersPerMinute.
func (uf VolumeFlowFactory) FromHectolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerMinute)
}

// FromKilolitersPerMinute creates a new VolumeFlow instance from a value in KilolitersPerMinute.
func (uf VolumeFlowFactory) FromKilolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerMinute)
}

// FromMegalitersPerMinute creates a new VolumeFlow instance from a value in MegalitersPerMinute.
func (uf VolumeFlowFactory) FromMegalitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerMinute)
}

// FromNanolitersPerHour creates a new VolumeFlow instance from a value in NanolitersPerHour.
func (uf VolumeFlowFactory) FromNanolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerHour)
}

// FromMicrolitersPerHour creates a new VolumeFlow instance from a value in MicrolitersPerHour.
func (uf VolumeFlowFactory) FromMicrolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerHour)
}

// FromMillilitersPerHour creates a new VolumeFlow instance from a value in MillilitersPerHour.
func (uf VolumeFlowFactory) FromMillilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerHour)
}

// FromCentilitersPerHour creates a new VolumeFlow instance from a value in CentilitersPerHour.
func (uf VolumeFlowFactory) FromCentilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerHour)
}

// FromDecilitersPerHour creates a new VolumeFlow instance from a value in DecilitersPerHour.
func (uf VolumeFlowFactory) FromDecilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerHour)
}

// FromDecalitersPerHour creates a new VolumeFlow instance from a value in DecalitersPerHour.
func (uf VolumeFlowFactory) FromDecalitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerHour)
}

// FromHectolitersPerHour creates a new VolumeFlow instance from a value in HectolitersPerHour.
func (uf VolumeFlowFactory) FromHectolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerHour)
}

// FromKilolitersPerHour creates a new VolumeFlow instance from a value in KilolitersPerHour.
func (uf VolumeFlowFactory) FromKilolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerHour)
}

// FromMegalitersPerHour creates a new VolumeFlow instance from a value in MegalitersPerHour.
func (uf VolumeFlowFactory) FromMegalitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerHour)
}

// FromNanolitersPerDay creates a new VolumeFlow instance from a value in NanolitersPerDay.
func (uf VolumeFlowFactory) FromNanolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerDay)
}

// FromMicrolitersPerDay creates a new VolumeFlow instance from a value in MicrolitersPerDay.
func (uf VolumeFlowFactory) FromMicrolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerDay)
}

// FromMillilitersPerDay creates a new VolumeFlow instance from a value in MillilitersPerDay.
func (uf VolumeFlowFactory) FromMillilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerDay)
}

// FromCentilitersPerDay creates a new VolumeFlow instance from a value in CentilitersPerDay.
func (uf VolumeFlowFactory) FromCentilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerDay)
}

// FromDecilitersPerDay creates a new VolumeFlow instance from a value in DecilitersPerDay.
func (uf VolumeFlowFactory) FromDecilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerDay)
}

// FromDecalitersPerDay creates a new VolumeFlow instance from a value in DecalitersPerDay.
func (uf VolumeFlowFactory) FromDecalitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerDay)
}

// FromHectolitersPerDay creates a new VolumeFlow instance from a value in HectolitersPerDay.
func (uf VolumeFlowFactory) FromHectolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerDay)
}

// FromKilolitersPerDay creates a new VolumeFlow instance from a value in KilolitersPerDay.
func (uf VolumeFlowFactory) FromKilolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerDay)
}

// FromMegalitersPerDay creates a new VolumeFlow instance from a value in MegalitersPerDay.
func (uf VolumeFlowFactory) FromMegalitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerDay)
}

// FromMegaukGallonsPerDay creates a new VolumeFlow instance from a value in MegaukGallonsPerDay.
func (uf VolumeFlowFactory) FromMegaukGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaukGallonPerDay)
}

// FromMegaukGallonsPerSecond creates a new VolumeFlow instance from a value in MegaukGallonsPerSecond.
func (uf VolumeFlowFactory) FromMegaukGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaukGallonPerSecond)
}


// newVolumeFlow creates a new VolumeFlow.
func newVolumeFlow(value float64, fromUnit VolumeFlowUnits) (*VolumeFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalVolumeFlowUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in VolumeFlowUnits", fromUnit)
	}
	a := &VolumeFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumeFlow in CubicMeterPerSecond unit (the base/default quantity).
func (a *VolumeFlow) BaseValue() float64 {
	return a.value
}


// CubicMetersPerSecond returns the VolumeFlow value in CubicMetersPerSecond.
//
// 
func (a *VolumeFlow) CubicMetersPerSecond() float64 {
	if a.cubic_meters_per_secondLazy != nil {
		return *a.cubic_meters_per_secondLazy
	}
	cubic_meters_per_second := a.convertFromBase(VolumeFlowCubicMeterPerSecond)
	a.cubic_meters_per_secondLazy = &cubic_meters_per_second
	return cubic_meters_per_second
}

// CubicMetersPerMinute returns the VolumeFlow value in CubicMetersPerMinute.
//
// 
func (a *VolumeFlow) CubicMetersPerMinute() float64 {
	if a.cubic_meters_per_minuteLazy != nil {
		return *a.cubic_meters_per_minuteLazy
	}
	cubic_meters_per_minute := a.convertFromBase(VolumeFlowCubicMeterPerMinute)
	a.cubic_meters_per_minuteLazy = &cubic_meters_per_minute
	return cubic_meters_per_minute
}

// CubicMetersPerHour returns the VolumeFlow value in CubicMetersPerHour.
//
// 
func (a *VolumeFlow) CubicMetersPerHour() float64 {
	if a.cubic_meters_per_hourLazy != nil {
		return *a.cubic_meters_per_hourLazy
	}
	cubic_meters_per_hour := a.convertFromBase(VolumeFlowCubicMeterPerHour)
	a.cubic_meters_per_hourLazy = &cubic_meters_per_hour
	return cubic_meters_per_hour
}

// CubicMetersPerDay returns the VolumeFlow value in CubicMetersPerDay.
//
// 
func (a *VolumeFlow) CubicMetersPerDay() float64 {
	if a.cubic_meters_per_dayLazy != nil {
		return *a.cubic_meters_per_dayLazy
	}
	cubic_meters_per_day := a.convertFromBase(VolumeFlowCubicMeterPerDay)
	a.cubic_meters_per_dayLazy = &cubic_meters_per_day
	return cubic_meters_per_day
}

// CubicFeetPerSecond returns the VolumeFlow value in CubicFeetPerSecond.
//
// 
func (a *VolumeFlow) CubicFeetPerSecond() float64 {
	if a.cubic_feet_per_secondLazy != nil {
		return *a.cubic_feet_per_secondLazy
	}
	cubic_feet_per_second := a.convertFromBase(VolumeFlowCubicFootPerSecond)
	a.cubic_feet_per_secondLazy = &cubic_feet_per_second
	return cubic_feet_per_second
}

// CubicFeetPerMinute returns the VolumeFlow value in CubicFeetPerMinute.
//
// 
func (a *VolumeFlow) CubicFeetPerMinute() float64 {
	if a.cubic_feet_per_minuteLazy != nil {
		return *a.cubic_feet_per_minuteLazy
	}
	cubic_feet_per_minute := a.convertFromBase(VolumeFlowCubicFootPerMinute)
	a.cubic_feet_per_minuteLazy = &cubic_feet_per_minute
	return cubic_feet_per_minute
}

// CubicFeetPerHour returns the VolumeFlow value in CubicFeetPerHour.
//
// 
func (a *VolumeFlow) CubicFeetPerHour() float64 {
	if a.cubic_feet_per_hourLazy != nil {
		return *a.cubic_feet_per_hourLazy
	}
	cubic_feet_per_hour := a.convertFromBase(VolumeFlowCubicFootPerHour)
	a.cubic_feet_per_hourLazy = &cubic_feet_per_hour
	return cubic_feet_per_hour
}

// CubicYardsPerSecond returns the VolumeFlow value in CubicYardsPerSecond.
//
// 
func (a *VolumeFlow) CubicYardsPerSecond() float64 {
	if a.cubic_yards_per_secondLazy != nil {
		return *a.cubic_yards_per_secondLazy
	}
	cubic_yards_per_second := a.convertFromBase(VolumeFlowCubicYardPerSecond)
	a.cubic_yards_per_secondLazy = &cubic_yards_per_second
	return cubic_yards_per_second
}

// CubicYardsPerMinute returns the VolumeFlow value in CubicYardsPerMinute.
//
// 
func (a *VolumeFlow) CubicYardsPerMinute() float64 {
	if a.cubic_yards_per_minuteLazy != nil {
		return *a.cubic_yards_per_minuteLazy
	}
	cubic_yards_per_minute := a.convertFromBase(VolumeFlowCubicYardPerMinute)
	a.cubic_yards_per_minuteLazy = &cubic_yards_per_minute
	return cubic_yards_per_minute
}

// CubicYardsPerHour returns the VolumeFlow value in CubicYardsPerHour.
//
// 
func (a *VolumeFlow) CubicYardsPerHour() float64 {
	if a.cubic_yards_per_hourLazy != nil {
		return *a.cubic_yards_per_hourLazy
	}
	cubic_yards_per_hour := a.convertFromBase(VolumeFlowCubicYardPerHour)
	a.cubic_yards_per_hourLazy = &cubic_yards_per_hour
	return cubic_yards_per_hour
}

// CubicYardsPerDay returns the VolumeFlow value in CubicYardsPerDay.
//
// 
func (a *VolumeFlow) CubicYardsPerDay() float64 {
	if a.cubic_yards_per_dayLazy != nil {
		return *a.cubic_yards_per_dayLazy
	}
	cubic_yards_per_day := a.convertFromBase(VolumeFlowCubicYardPerDay)
	a.cubic_yards_per_dayLazy = &cubic_yards_per_day
	return cubic_yards_per_day
}

// MillionUsGallonsPerDay returns the VolumeFlow value in MillionUsGallonsPerDay.
//
// 
func (a *VolumeFlow) MillionUsGallonsPerDay() float64 {
	if a.million_us_gallons_per_dayLazy != nil {
		return *a.million_us_gallons_per_dayLazy
	}
	million_us_gallons_per_day := a.convertFromBase(VolumeFlowMillionUsGallonPerDay)
	a.million_us_gallons_per_dayLazy = &million_us_gallons_per_day
	return million_us_gallons_per_day
}

// UsGallonsPerDay returns the VolumeFlow value in UsGallonsPerDay.
//
// 
func (a *VolumeFlow) UsGallonsPerDay() float64 {
	if a.us_gallons_per_dayLazy != nil {
		return *a.us_gallons_per_dayLazy
	}
	us_gallons_per_day := a.convertFromBase(VolumeFlowUsGallonPerDay)
	a.us_gallons_per_dayLazy = &us_gallons_per_day
	return us_gallons_per_day
}

// LitersPerSecond returns the VolumeFlow value in LitersPerSecond.
//
// 
func (a *VolumeFlow) LitersPerSecond() float64 {
	if a.liters_per_secondLazy != nil {
		return *a.liters_per_secondLazy
	}
	liters_per_second := a.convertFromBase(VolumeFlowLiterPerSecond)
	a.liters_per_secondLazy = &liters_per_second
	return liters_per_second
}

// LitersPerMinute returns the VolumeFlow value in LitersPerMinute.
//
// 
func (a *VolumeFlow) LitersPerMinute() float64 {
	if a.liters_per_minuteLazy != nil {
		return *a.liters_per_minuteLazy
	}
	liters_per_minute := a.convertFromBase(VolumeFlowLiterPerMinute)
	a.liters_per_minuteLazy = &liters_per_minute
	return liters_per_minute
}

// LitersPerHour returns the VolumeFlow value in LitersPerHour.
//
// 
func (a *VolumeFlow) LitersPerHour() float64 {
	if a.liters_per_hourLazy != nil {
		return *a.liters_per_hourLazy
	}
	liters_per_hour := a.convertFromBase(VolumeFlowLiterPerHour)
	a.liters_per_hourLazy = &liters_per_hour
	return liters_per_hour
}

// LitersPerDay returns the VolumeFlow value in LitersPerDay.
//
// 
func (a *VolumeFlow) LitersPerDay() float64 {
	if a.liters_per_dayLazy != nil {
		return *a.liters_per_dayLazy
	}
	liters_per_day := a.convertFromBase(VolumeFlowLiterPerDay)
	a.liters_per_dayLazy = &liters_per_day
	return liters_per_day
}

// UsGallonsPerSecond returns the VolumeFlow value in UsGallonsPerSecond.
//
// 
func (a *VolumeFlow) UsGallonsPerSecond() float64 {
	if a.us_gallons_per_secondLazy != nil {
		return *a.us_gallons_per_secondLazy
	}
	us_gallons_per_second := a.convertFromBase(VolumeFlowUsGallonPerSecond)
	a.us_gallons_per_secondLazy = &us_gallons_per_second
	return us_gallons_per_second
}

// UsGallonsPerMinute returns the VolumeFlow value in UsGallonsPerMinute.
//
// 
func (a *VolumeFlow) UsGallonsPerMinute() float64 {
	if a.us_gallons_per_minuteLazy != nil {
		return *a.us_gallons_per_minuteLazy
	}
	us_gallons_per_minute := a.convertFromBase(VolumeFlowUsGallonPerMinute)
	a.us_gallons_per_minuteLazy = &us_gallons_per_minute
	return us_gallons_per_minute
}

// UkGallonsPerDay returns the VolumeFlow value in UkGallonsPerDay.
//
// 
func (a *VolumeFlow) UkGallonsPerDay() float64 {
	if a.uk_gallons_per_dayLazy != nil {
		return *a.uk_gallons_per_dayLazy
	}
	uk_gallons_per_day := a.convertFromBase(VolumeFlowUkGallonPerDay)
	a.uk_gallons_per_dayLazy = &uk_gallons_per_day
	return uk_gallons_per_day
}

// UkGallonsPerHour returns the VolumeFlow value in UkGallonsPerHour.
//
// 
func (a *VolumeFlow) UkGallonsPerHour() float64 {
	if a.uk_gallons_per_hourLazy != nil {
		return *a.uk_gallons_per_hourLazy
	}
	uk_gallons_per_hour := a.convertFromBase(VolumeFlowUkGallonPerHour)
	a.uk_gallons_per_hourLazy = &uk_gallons_per_hour
	return uk_gallons_per_hour
}

// UkGallonsPerMinute returns the VolumeFlow value in UkGallonsPerMinute.
//
// 
func (a *VolumeFlow) UkGallonsPerMinute() float64 {
	if a.uk_gallons_per_minuteLazy != nil {
		return *a.uk_gallons_per_minuteLazy
	}
	uk_gallons_per_minute := a.convertFromBase(VolumeFlowUkGallonPerMinute)
	a.uk_gallons_per_minuteLazy = &uk_gallons_per_minute
	return uk_gallons_per_minute
}

// UkGallonsPerSecond returns the VolumeFlow value in UkGallonsPerSecond.
//
// 
func (a *VolumeFlow) UkGallonsPerSecond() float64 {
	if a.uk_gallons_per_secondLazy != nil {
		return *a.uk_gallons_per_secondLazy
	}
	uk_gallons_per_second := a.convertFromBase(VolumeFlowUkGallonPerSecond)
	a.uk_gallons_per_secondLazy = &uk_gallons_per_second
	return uk_gallons_per_second
}

// KilousGallonsPerMinute returns the VolumeFlow value in KilousGallonsPerMinute.
//
// 
func (a *VolumeFlow) KilousGallonsPerMinute() float64 {
	if a.kilous_gallons_per_minuteLazy != nil {
		return *a.kilous_gallons_per_minuteLazy
	}
	kilous_gallons_per_minute := a.convertFromBase(VolumeFlowKilousGallonPerMinute)
	a.kilous_gallons_per_minuteLazy = &kilous_gallons_per_minute
	return kilous_gallons_per_minute
}

// UsGallonsPerHour returns the VolumeFlow value in UsGallonsPerHour.
//
// 
func (a *VolumeFlow) UsGallonsPerHour() float64 {
	if a.us_gallons_per_hourLazy != nil {
		return *a.us_gallons_per_hourLazy
	}
	us_gallons_per_hour := a.convertFromBase(VolumeFlowUsGallonPerHour)
	a.us_gallons_per_hourLazy = &us_gallons_per_hour
	return us_gallons_per_hour
}

// CubicDecimetersPerMinute returns the VolumeFlow value in CubicDecimetersPerMinute.
//
// 
func (a *VolumeFlow) CubicDecimetersPerMinute() float64 {
	if a.cubic_decimeters_per_minuteLazy != nil {
		return *a.cubic_decimeters_per_minuteLazy
	}
	cubic_decimeters_per_minute := a.convertFromBase(VolumeFlowCubicDecimeterPerMinute)
	a.cubic_decimeters_per_minuteLazy = &cubic_decimeters_per_minute
	return cubic_decimeters_per_minute
}

// OilBarrelsPerDay returns the VolumeFlow value in OilBarrelsPerDay.
//
// 
func (a *VolumeFlow) OilBarrelsPerDay() float64 {
	if a.oil_barrels_per_dayLazy != nil {
		return *a.oil_barrels_per_dayLazy
	}
	oil_barrels_per_day := a.convertFromBase(VolumeFlowOilBarrelPerDay)
	a.oil_barrels_per_dayLazy = &oil_barrels_per_day
	return oil_barrels_per_day
}

// OilBarrelsPerMinute returns the VolumeFlow value in OilBarrelsPerMinute.
//
// 
func (a *VolumeFlow) OilBarrelsPerMinute() float64 {
	if a.oil_barrels_per_minuteLazy != nil {
		return *a.oil_barrels_per_minuteLazy
	}
	oil_barrels_per_minute := a.convertFromBase(VolumeFlowOilBarrelPerMinute)
	a.oil_barrels_per_minuteLazy = &oil_barrels_per_minute
	return oil_barrels_per_minute
}

// OilBarrelsPerHour returns the VolumeFlow value in OilBarrelsPerHour.
//
// 
func (a *VolumeFlow) OilBarrelsPerHour() float64 {
	if a.oil_barrels_per_hourLazy != nil {
		return *a.oil_barrels_per_hourLazy
	}
	oil_barrels_per_hour := a.convertFromBase(VolumeFlowOilBarrelPerHour)
	a.oil_barrels_per_hourLazy = &oil_barrels_per_hour
	return oil_barrels_per_hour
}

// OilBarrelsPerSecond returns the VolumeFlow value in OilBarrelsPerSecond.
//
// 
func (a *VolumeFlow) OilBarrelsPerSecond() float64 {
	if a.oil_barrels_per_secondLazy != nil {
		return *a.oil_barrels_per_secondLazy
	}
	oil_barrels_per_second := a.convertFromBase(VolumeFlowOilBarrelPerSecond)
	a.oil_barrels_per_secondLazy = &oil_barrels_per_second
	return oil_barrels_per_second
}

// CubicMillimetersPerSecond returns the VolumeFlow value in CubicMillimetersPerSecond.
//
// 
func (a *VolumeFlow) CubicMillimetersPerSecond() float64 {
	if a.cubic_millimeters_per_secondLazy != nil {
		return *a.cubic_millimeters_per_secondLazy
	}
	cubic_millimeters_per_second := a.convertFromBase(VolumeFlowCubicMillimeterPerSecond)
	a.cubic_millimeters_per_secondLazy = &cubic_millimeters_per_second
	return cubic_millimeters_per_second
}

// AcreFeetPerSecond returns the VolumeFlow value in AcreFeetPerSecond.
//
// 
func (a *VolumeFlow) AcreFeetPerSecond() float64 {
	if a.acre_feet_per_secondLazy != nil {
		return *a.acre_feet_per_secondLazy
	}
	acre_feet_per_second := a.convertFromBase(VolumeFlowAcreFootPerSecond)
	a.acre_feet_per_secondLazy = &acre_feet_per_second
	return acre_feet_per_second
}

// AcreFeetPerMinute returns the VolumeFlow value in AcreFeetPerMinute.
//
// 
func (a *VolumeFlow) AcreFeetPerMinute() float64 {
	if a.acre_feet_per_minuteLazy != nil {
		return *a.acre_feet_per_minuteLazy
	}
	acre_feet_per_minute := a.convertFromBase(VolumeFlowAcreFootPerMinute)
	a.acre_feet_per_minuteLazy = &acre_feet_per_minute
	return acre_feet_per_minute
}

// AcreFeetPerHour returns the VolumeFlow value in AcreFeetPerHour.
//
// 
func (a *VolumeFlow) AcreFeetPerHour() float64 {
	if a.acre_feet_per_hourLazy != nil {
		return *a.acre_feet_per_hourLazy
	}
	acre_feet_per_hour := a.convertFromBase(VolumeFlowAcreFootPerHour)
	a.acre_feet_per_hourLazy = &acre_feet_per_hour
	return acre_feet_per_hour
}

// AcreFeetPerDay returns the VolumeFlow value in AcreFeetPerDay.
//
// 
func (a *VolumeFlow) AcreFeetPerDay() float64 {
	if a.acre_feet_per_dayLazy != nil {
		return *a.acre_feet_per_dayLazy
	}
	acre_feet_per_day := a.convertFromBase(VolumeFlowAcreFootPerDay)
	a.acre_feet_per_dayLazy = &acre_feet_per_day
	return acre_feet_per_day
}

// CubicCentimetersPerMinute returns the VolumeFlow value in CubicCentimetersPerMinute.
//
// 
func (a *VolumeFlow) CubicCentimetersPerMinute() float64 {
	if a.cubic_centimeters_per_minuteLazy != nil {
		return *a.cubic_centimeters_per_minuteLazy
	}
	cubic_centimeters_per_minute := a.convertFromBase(VolumeFlowCubicCentimeterPerMinute)
	a.cubic_centimeters_per_minuteLazy = &cubic_centimeters_per_minute
	return cubic_centimeters_per_minute
}

// MegausGallonsPerDay returns the VolumeFlow value in MegausGallonsPerDay.
//
// 
func (a *VolumeFlow) MegausGallonsPerDay() float64 {
	if a.megaus_gallons_per_dayLazy != nil {
		return *a.megaus_gallons_per_dayLazy
	}
	megaus_gallons_per_day := a.convertFromBase(VolumeFlowMegausGallonPerDay)
	a.megaus_gallons_per_dayLazy = &megaus_gallons_per_day
	return megaus_gallons_per_day
}

// NanolitersPerSecond returns the VolumeFlow value in NanolitersPerSecond.
//
// 
func (a *VolumeFlow) NanolitersPerSecond() float64 {
	if a.nanoliters_per_secondLazy != nil {
		return *a.nanoliters_per_secondLazy
	}
	nanoliters_per_second := a.convertFromBase(VolumeFlowNanoliterPerSecond)
	a.nanoliters_per_secondLazy = &nanoliters_per_second
	return nanoliters_per_second
}

// MicrolitersPerSecond returns the VolumeFlow value in MicrolitersPerSecond.
//
// 
func (a *VolumeFlow) MicrolitersPerSecond() float64 {
	if a.microliters_per_secondLazy != nil {
		return *a.microliters_per_secondLazy
	}
	microliters_per_second := a.convertFromBase(VolumeFlowMicroliterPerSecond)
	a.microliters_per_secondLazy = &microliters_per_second
	return microliters_per_second
}

// MillilitersPerSecond returns the VolumeFlow value in MillilitersPerSecond.
//
// 
func (a *VolumeFlow) MillilitersPerSecond() float64 {
	if a.milliliters_per_secondLazy != nil {
		return *a.milliliters_per_secondLazy
	}
	milliliters_per_second := a.convertFromBase(VolumeFlowMilliliterPerSecond)
	a.milliliters_per_secondLazy = &milliliters_per_second
	return milliliters_per_second
}

// CentilitersPerSecond returns the VolumeFlow value in CentilitersPerSecond.
//
// 
func (a *VolumeFlow) CentilitersPerSecond() float64 {
	if a.centiliters_per_secondLazy != nil {
		return *a.centiliters_per_secondLazy
	}
	centiliters_per_second := a.convertFromBase(VolumeFlowCentiliterPerSecond)
	a.centiliters_per_secondLazy = &centiliters_per_second
	return centiliters_per_second
}

// DecilitersPerSecond returns the VolumeFlow value in DecilitersPerSecond.
//
// 
func (a *VolumeFlow) DecilitersPerSecond() float64 {
	if a.deciliters_per_secondLazy != nil {
		return *a.deciliters_per_secondLazy
	}
	deciliters_per_second := a.convertFromBase(VolumeFlowDeciliterPerSecond)
	a.deciliters_per_secondLazy = &deciliters_per_second
	return deciliters_per_second
}

// DecalitersPerSecond returns the VolumeFlow value in DecalitersPerSecond.
//
// 
func (a *VolumeFlow) DecalitersPerSecond() float64 {
	if a.decaliters_per_secondLazy != nil {
		return *a.decaliters_per_secondLazy
	}
	decaliters_per_second := a.convertFromBase(VolumeFlowDecaliterPerSecond)
	a.decaliters_per_secondLazy = &decaliters_per_second
	return decaliters_per_second
}

// HectolitersPerSecond returns the VolumeFlow value in HectolitersPerSecond.
//
// 
func (a *VolumeFlow) HectolitersPerSecond() float64 {
	if a.hectoliters_per_secondLazy != nil {
		return *a.hectoliters_per_secondLazy
	}
	hectoliters_per_second := a.convertFromBase(VolumeFlowHectoliterPerSecond)
	a.hectoliters_per_secondLazy = &hectoliters_per_second
	return hectoliters_per_second
}

// KilolitersPerSecond returns the VolumeFlow value in KilolitersPerSecond.
//
// 
func (a *VolumeFlow) KilolitersPerSecond() float64 {
	if a.kiloliters_per_secondLazy != nil {
		return *a.kiloliters_per_secondLazy
	}
	kiloliters_per_second := a.convertFromBase(VolumeFlowKiloliterPerSecond)
	a.kiloliters_per_secondLazy = &kiloliters_per_second
	return kiloliters_per_second
}

// MegalitersPerSecond returns the VolumeFlow value in MegalitersPerSecond.
//
// 
func (a *VolumeFlow) MegalitersPerSecond() float64 {
	if a.megaliters_per_secondLazy != nil {
		return *a.megaliters_per_secondLazy
	}
	megaliters_per_second := a.convertFromBase(VolumeFlowMegaliterPerSecond)
	a.megaliters_per_secondLazy = &megaliters_per_second
	return megaliters_per_second
}

// NanolitersPerMinute returns the VolumeFlow value in NanolitersPerMinute.
//
// 
func (a *VolumeFlow) NanolitersPerMinute() float64 {
	if a.nanoliters_per_minuteLazy != nil {
		return *a.nanoliters_per_minuteLazy
	}
	nanoliters_per_minute := a.convertFromBase(VolumeFlowNanoliterPerMinute)
	a.nanoliters_per_minuteLazy = &nanoliters_per_minute
	return nanoliters_per_minute
}

// MicrolitersPerMinute returns the VolumeFlow value in MicrolitersPerMinute.
//
// 
func (a *VolumeFlow) MicrolitersPerMinute() float64 {
	if a.microliters_per_minuteLazy != nil {
		return *a.microliters_per_minuteLazy
	}
	microliters_per_minute := a.convertFromBase(VolumeFlowMicroliterPerMinute)
	a.microliters_per_minuteLazy = &microliters_per_minute
	return microliters_per_minute
}

// MillilitersPerMinute returns the VolumeFlow value in MillilitersPerMinute.
//
// 
func (a *VolumeFlow) MillilitersPerMinute() float64 {
	if a.milliliters_per_minuteLazy != nil {
		return *a.milliliters_per_minuteLazy
	}
	milliliters_per_minute := a.convertFromBase(VolumeFlowMilliliterPerMinute)
	a.milliliters_per_minuteLazy = &milliliters_per_minute
	return milliliters_per_minute
}

// CentilitersPerMinute returns the VolumeFlow value in CentilitersPerMinute.
//
// 
func (a *VolumeFlow) CentilitersPerMinute() float64 {
	if a.centiliters_per_minuteLazy != nil {
		return *a.centiliters_per_minuteLazy
	}
	centiliters_per_minute := a.convertFromBase(VolumeFlowCentiliterPerMinute)
	a.centiliters_per_minuteLazy = &centiliters_per_minute
	return centiliters_per_minute
}

// DecilitersPerMinute returns the VolumeFlow value in DecilitersPerMinute.
//
// 
func (a *VolumeFlow) DecilitersPerMinute() float64 {
	if a.deciliters_per_minuteLazy != nil {
		return *a.deciliters_per_minuteLazy
	}
	deciliters_per_minute := a.convertFromBase(VolumeFlowDeciliterPerMinute)
	a.deciliters_per_minuteLazy = &deciliters_per_minute
	return deciliters_per_minute
}

// DecalitersPerMinute returns the VolumeFlow value in DecalitersPerMinute.
//
// 
func (a *VolumeFlow) DecalitersPerMinute() float64 {
	if a.decaliters_per_minuteLazy != nil {
		return *a.decaliters_per_minuteLazy
	}
	decaliters_per_minute := a.convertFromBase(VolumeFlowDecaliterPerMinute)
	a.decaliters_per_minuteLazy = &decaliters_per_minute
	return decaliters_per_minute
}

// HectolitersPerMinute returns the VolumeFlow value in HectolitersPerMinute.
//
// 
func (a *VolumeFlow) HectolitersPerMinute() float64 {
	if a.hectoliters_per_minuteLazy != nil {
		return *a.hectoliters_per_minuteLazy
	}
	hectoliters_per_minute := a.convertFromBase(VolumeFlowHectoliterPerMinute)
	a.hectoliters_per_minuteLazy = &hectoliters_per_minute
	return hectoliters_per_minute
}

// KilolitersPerMinute returns the VolumeFlow value in KilolitersPerMinute.
//
// 
func (a *VolumeFlow) KilolitersPerMinute() float64 {
	if a.kiloliters_per_minuteLazy != nil {
		return *a.kiloliters_per_minuteLazy
	}
	kiloliters_per_minute := a.convertFromBase(VolumeFlowKiloliterPerMinute)
	a.kiloliters_per_minuteLazy = &kiloliters_per_minute
	return kiloliters_per_minute
}

// MegalitersPerMinute returns the VolumeFlow value in MegalitersPerMinute.
//
// 
func (a *VolumeFlow) MegalitersPerMinute() float64 {
	if a.megaliters_per_minuteLazy != nil {
		return *a.megaliters_per_minuteLazy
	}
	megaliters_per_minute := a.convertFromBase(VolumeFlowMegaliterPerMinute)
	a.megaliters_per_minuteLazy = &megaliters_per_minute
	return megaliters_per_minute
}

// NanolitersPerHour returns the VolumeFlow value in NanolitersPerHour.
//
// 
func (a *VolumeFlow) NanolitersPerHour() float64 {
	if a.nanoliters_per_hourLazy != nil {
		return *a.nanoliters_per_hourLazy
	}
	nanoliters_per_hour := a.convertFromBase(VolumeFlowNanoliterPerHour)
	a.nanoliters_per_hourLazy = &nanoliters_per_hour
	return nanoliters_per_hour
}

// MicrolitersPerHour returns the VolumeFlow value in MicrolitersPerHour.
//
// 
func (a *VolumeFlow) MicrolitersPerHour() float64 {
	if a.microliters_per_hourLazy != nil {
		return *a.microliters_per_hourLazy
	}
	microliters_per_hour := a.convertFromBase(VolumeFlowMicroliterPerHour)
	a.microliters_per_hourLazy = &microliters_per_hour
	return microliters_per_hour
}

// MillilitersPerHour returns the VolumeFlow value in MillilitersPerHour.
//
// 
func (a *VolumeFlow) MillilitersPerHour() float64 {
	if a.milliliters_per_hourLazy != nil {
		return *a.milliliters_per_hourLazy
	}
	milliliters_per_hour := a.convertFromBase(VolumeFlowMilliliterPerHour)
	a.milliliters_per_hourLazy = &milliliters_per_hour
	return milliliters_per_hour
}

// CentilitersPerHour returns the VolumeFlow value in CentilitersPerHour.
//
// 
func (a *VolumeFlow) CentilitersPerHour() float64 {
	if a.centiliters_per_hourLazy != nil {
		return *a.centiliters_per_hourLazy
	}
	centiliters_per_hour := a.convertFromBase(VolumeFlowCentiliterPerHour)
	a.centiliters_per_hourLazy = &centiliters_per_hour
	return centiliters_per_hour
}

// DecilitersPerHour returns the VolumeFlow value in DecilitersPerHour.
//
// 
func (a *VolumeFlow) DecilitersPerHour() float64 {
	if a.deciliters_per_hourLazy != nil {
		return *a.deciliters_per_hourLazy
	}
	deciliters_per_hour := a.convertFromBase(VolumeFlowDeciliterPerHour)
	a.deciliters_per_hourLazy = &deciliters_per_hour
	return deciliters_per_hour
}

// DecalitersPerHour returns the VolumeFlow value in DecalitersPerHour.
//
// 
func (a *VolumeFlow) DecalitersPerHour() float64 {
	if a.decaliters_per_hourLazy != nil {
		return *a.decaliters_per_hourLazy
	}
	decaliters_per_hour := a.convertFromBase(VolumeFlowDecaliterPerHour)
	a.decaliters_per_hourLazy = &decaliters_per_hour
	return decaliters_per_hour
}

// HectolitersPerHour returns the VolumeFlow value in HectolitersPerHour.
//
// 
func (a *VolumeFlow) HectolitersPerHour() float64 {
	if a.hectoliters_per_hourLazy != nil {
		return *a.hectoliters_per_hourLazy
	}
	hectoliters_per_hour := a.convertFromBase(VolumeFlowHectoliterPerHour)
	a.hectoliters_per_hourLazy = &hectoliters_per_hour
	return hectoliters_per_hour
}

// KilolitersPerHour returns the VolumeFlow value in KilolitersPerHour.
//
// 
func (a *VolumeFlow) KilolitersPerHour() float64 {
	if a.kiloliters_per_hourLazy != nil {
		return *a.kiloliters_per_hourLazy
	}
	kiloliters_per_hour := a.convertFromBase(VolumeFlowKiloliterPerHour)
	a.kiloliters_per_hourLazy = &kiloliters_per_hour
	return kiloliters_per_hour
}

// MegalitersPerHour returns the VolumeFlow value in MegalitersPerHour.
//
// 
func (a *VolumeFlow) MegalitersPerHour() float64 {
	if a.megaliters_per_hourLazy != nil {
		return *a.megaliters_per_hourLazy
	}
	megaliters_per_hour := a.convertFromBase(VolumeFlowMegaliterPerHour)
	a.megaliters_per_hourLazy = &megaliters_per_hour
	return megaliters_per_hour
}

// NanolitersPerDay returns the VolumeFlow value in NanolitersPerDay.
//
// 
func (a *VolumeFlow) NanolitersPerDay() float64 {
	if a.nanoliters_per_dayLazy != nil {
		return *a.nanoliters_per_dayLazy
	}
	nanoliters_per_day := a.convertFromBase(VolumeFlowNanoliterPerDay)
	a.nanoliters_per_dayLazy = &nanoliters_per_day
	return nanoliters_per_day
}

// MicrolitersPerDay returns the VolumeFlow value in MicrolitersPerDay.
//
// 
func (a *VolumeFlow) MicrolitersPerDay() float64 {
	if a.microliters_per_dayLazy != nil {
		return *a.microliters_per_dayLazy
	}
	microliters_per_day := a.convertFromBase(VolumeFlowMicroliterPerDay)
	a.microliters_per_dayLazy = &microliters_per_day
	return microliters_per_day
}

// MillilitersPerDay returns the VolumeFlow value in MillilitersPerDay.
//
// 
func (a *VolumeFlow) MillilitersPerDay() float64 {
	if a.milliliters_per_dayLazy != nil {
		return *a.milliliters_per_dayLazy
	}
	milliliters_per_day := a.convertFromBase(VolumeFlowMilliliterPerDay)
	a.milliliters_per_dayLazy = &milliliters_per_day
	return milliliters_per_day
}

// CentilitersPerDay returns the VolumeFlow value in CentilitersPerDay.
//
// 
func (a *VolumeFlow) CentilitersPerDay() float64 {
	if a.centiliters_per_dayLazy != nil {
		return *a.centiliters_per_dayLazy
	}
	centiliters_per_day := a.convertFromBase(VolumeFlowCentiliterPerDay)
	a.centiliters_per_dayLazy = &centiliters_per_day
	return centiliters_per_day
}

// DecilitersPerDay returns the VolumeFlow value in DecilitersPerDay.
//
// 
func (a *VolumeFlow) DecilitersPerDay() float64 {
	if a.deciliters_per_dayLazy != nil {
		return *a.deciliters_per_dayLazy
	}
	deciliters_per_day := a.convertFromBase(VolumeFlowDeciliterPerDay)
	a.deciliters_per_dayLazy = &deciliters_per_day
	return deciliters_per_day
}

// DecalitersPerDay returns the VolumeFlow value in DecalitersPerDay.
//
// 
func (a *VolumeFlow) DecalitersPerDay() float64 {
	if a.decaliters_per_dayLazy != nil {
		return *a.decaliters_per_dayLazy
	}
	decaliters_per_day := a.convertFromBase(VolumeFlowDecaliterPerDay)
	a.decaliters_per_dayLazy = &decaliters_per_day
	return decaliters_per_day
}

// HectolitersPerDay returns the VolumeFlow value in HectolitersPerDay.
//
// 
func (a *VolumeFlow) HectolitersPerDay() float64 {
	if a.hectoliters_per_dayLazy != nil {
		return *a.hectoliters_per_dayLazy
	}
	hectoliters_per_day := a.convertFromBase(VolumeFlowHectoliterPerDay)
	a.hectoliters_per_dayLazy = &hectoliters_per_day
	return hectoliters_per_day
}

// KilolitersPerDay returns the VolumeFlow value in KilolitersPerDay.
//
// 
func (a *VolumeFlow) KilolitersPerDay() float64 {
	if a.kiloliters_per_dayLazy != nil {
		return *a.kiloliters_per_dayLazy
	}
	kiloliters_per_day := a.convertFromBase(VolumeFlowKiloliterPerDay)
	a.kiloliters_per_dayLazy = &kiloliters_per_day
	return kiloliters_per_day
}

// MegalitersPerDay returns the VolumeFlow value in MegalitersPerDay.
//
// 
func (a *VolumeFlow) MegalitersPerDay() float64 {
	if a.megaliters_per_dayLazy != nil {
		return *a.megaliters_per_dayLazy
	}
	megaliters_per_day := a.convertFromBase(VolumeFlowMegaliterPerDay)
	a.megaliters_per_dayLazy = &megaliters_per_day
	return megaliters_per_day
}

// MegaukGallonsPerDay returns the VolumeFlow value in MegaukGallonsPerDay.
//
// 
func (a *VolumeFlow) MegaukGallonsPerDay() float64 {
	if a.megauk_gallons_per_dayLazy != nil {
		return *a.megauk_gallons_per_dayLazy
	}
	megauk_gallons_per_day := a.convertFromBase(VolumeFlowMegaukGallonPerDay)
	a.megauk_gallons_per_dayLazy = &megauk_gallons_per_day
	return megauk_gallons_per_day
}

// MegaukGallonsPerSecond returns the VolumeFlow value in MegaukGallonsPerSecond.
//
// 
func (a *VolumeFlow) MegaukGallonsPerSecond() float64 {
	if a.megauk_gallons_per_secondLazy != nil {
		return *a.megauk_gallons_per_secondLazy
	}
	megauk_gallons_per_second := a.convertFromBase(VolumeFlowMegaukGallonPerSecond)
	a.megauk_gallons_per_secondLazy = &megauk_gallons_per_second
	return megauk_gallons_per_second
}


// ToDto creates a VolumeFlowDto representation from the VolumeFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerSecond by default.
func (a *VolumeFlow) ToDto(holdInUnit *VolumeFlowUnits) VolumeFlowDto {
	if holdInUnit == nil {
		defaultUnit := VolumeFlowCubicMeterPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return VolumeFlowDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VolumeFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerSecond by default.
func (a *VolumeFlow) ToDtoJSON(holdInUnit *VolumeFlowUnits) ([]byte, error) {
	// Convert to VolumeFlowDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VolumeFlow to a specific unit value.
// The function uses the provided unit type (VolumeFlowUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VolumeFlow) Convert(toUnit VolumeFlowUnits) float64 {
	switch toUnit { 
    case VolumeFlowCubicMeterPerSecond:
		return a.CubicMetersPerSecond()
    case VolumeFlowCubicMeterPerMinute:
		return a.CubicMetersPerMinute()
    case VolumeFlowCubicMeterPerHour:
		return a.CubicMetersPerHour()
    case VolumeFlowCubicMeterPerDay:
		return a.CubicMetersPerDay()
    case VolumeFlowCubicFootPerSecond:
		return a.CubicFeetPerSecond()
    case VolumeFlowCubicFootPerMinute:
		return a.CubicFeetPerMinute()
    case VolumeFlowCubicFootPerHour:
		return a.CubicFeetPerHour()
    case VolumeFlowCubicYardPerSecond:
		return a.CubicYardsPerSecond()
    case VolumeFlowCubicYardPerMinute:
		return a.CubicYardsPerMinute()
    case VolumeFlowCubicYardPerHour:
		return a.CubicYardsPerHour()
    case VolumeFlowCubicYardPerDay:
		return a.CubicYardsPerDay()
    case VolumeFlowMillionUsGallonPerDay:
		return a.MillionUsGallonsPerDay()
    case VolumeFlowUsGallonPerDay:
		return a.UsGallonsPerDay()
    case VolumeFlowLiterPerSecond:
		return a.LitersPerSecond()
    case VolumeFlowLiterPerMinute:
		return a.LitersPerMinute()
    case VolumeFlowLiterPerHour:
		return a.LitersPerHour()
    case VolumeFlowLiterPerDay:
		return a.LitersPerDay()
    case VolumeFlowUsGallonPerSecond:
		return a.UsGallonsPerSecond()
    case VolumeFlowUsGallonPerMinute:
		return a.UsGallonsPerMinute()
    case VolumeFlowUkGallonPerDay:
		return a.UkGallonsPerDay()
    case VolumeFlowUkGallonPerHour:
		return a.UkGallonsPerHour()
    case VolumeFlowUkGallonPerMinute:
		return a.UkGallonsPerMinute()
    case VolumeFlowUkGallonPerSecond:
		return a.UkGallonsPerSecond()
    case VolumeFlowKilousGallonPerMinute:
		return a.KilousGallonsPerMinute()
    case VolumeFlowUsGallonPerHour:
		return a.UsGallonsPerHour()
    case VolumeFlowCubicDecimeterPerMinute:
		return a.CubicDecimetersPerMinute()
    case VolumeFlowOilBarrelPerDay:
		return a.OilBarrelsPerDay()
    case VolumeFlowOilBarrelPerMinute:
		return a.OilBarrelsPerMinute()
    case VolumeFlowOilBarrelPerHour:
		return a.OilBarrelsPerHour()
    case VolumeFlowOilBarrelPerSecond:
		return a.OilBarrelsPerSecond()
    case VolumeFlowCubicMillimeterPerSecond:
		return a.CubicMillimetersPerSecond()
    case VolumeFlowAcreFootPerSecond:
		return a.AcreFeetPerSecond()
    case VolumeFlowAcreFootPerMinute:
		return a.AcreFeetPerMinute()
    case VolumeFlowAcreFootPerHour:
		return a.AcreFeetPerHour()
    case VolumeFlowAcreFootPerDay:
		return a.AcreFeetPerDay()
    case VolumeFlowCubicCentimeterPerMinute:
		return a.CubicCentimetersPerMinute()
    case VolumeFlowMegausGallonPerDay:
		return a.MegausGallonsPerDay()
    case VolumeFlowNanoliterPerSecond:
		return a.NanolitersPerSecond()
    case VolumeFlowMicroliterPerSecond:
		return a.MicrolitersPerSecond()
    case VolumeFlowMilliliterPerSecond:
		return a.MillilitersPerSecond()
    case VolumeFlowCentiliterPerSecond:
		return a.CentilitersPerSecond()
    case VolumeFlowDeciliterPerSecond:
		return a.DecilitersPerSecond()
    case VolumeFlowDecaliterPerSecond:
		return a.DecalitersPerSecond()
    case VolumeFlowHectoliterPerSecond:
		return a.HectolitersPerSecond()
    case VolumeFlowKiloliterPerSecond:
		return a.KilolitersPerSecond()
    case VolumeFlowMegaliterPerSecond:
		return a.MegalitersPerSecond()
    case VolumeFlowNanoliterPerMinute:
		return a.NanolitersPerMinute()
    case VolumeFlowMicroliterPerMinute:
		return a.MicrolitersPerMinute()
    case VolumeFlowMilliliterPerMinute:
		return a.MillilitersPerMinute()
    case VolumeFlowCentiliterPerMinute:
		return a.CentilitersPerMinute()
    case VolumeFlowDeciliterPerMinute:
		return a.DecilitersPerMinute()
    case VolumeFlowDecaliterPerMinute:
		return a.DecalitersPerMinute()
    case VolumeFlowHectoliterPerMinute:
		return a.HectolitersPerMinute()
    case VolumeFlowKiloliterPerMinute:
		return a.KilolitersPerMinute()
    case VolumeFlowMegaliterPerMinute:
		return a.MegalitersPerMinute()
    case VolumeFlowNanoliterPerHour:
		return a.NanolitersPerHour()
    case VolumeFlowMicroliterPerHour:
		return a.MicrolitersPerHour()
    case VolumeFlowMilliliterPerHour:
		return a.MillilitersPerHour()
    case VolumeFlowCentiliterPerHour:
		return a.CentilitersPerHour()
    case VolumeFlowDeciliterPerHour:
		return a.DecilitersPerHour()
    case VolumeFlowDecaliterPerHour:
		return a.DecalitersPerHour()
    case VolumeFlowHectoliterPerHour:
		return a.HectolitersPerHour()
    case VolumeFlowKiloliterPerHour:
		return a.KilolitersPerHour()
    case VolumeFlowMegaliterPerHour:
		return a.MegalitersPerHour()
    case VolumeFlowNanoliterPerDay:
		return a.NanolitersPerDay()
    case VolumeFlowMicroliterPerDay:
		return a.MicrolitersPerDay()
    case VolumeFlowMilliliterPerDay:
		return a.MillilitersPerDay()
    case VolumeFlowCentiliterPerDay:
		return a.CentilitersPerDay()
    case VolumeFlowDeciliterPerDay:
		return a.DecilitersPerDay()
    case VolumeFlowDecaliterPerDay:
		return a.DecalitersPerDay()
    case VolumeFlowHectoliterPerDay:
		return a.HectolitersPerDay()
    case VolumeFlowKiloliterPerDay:
		return a.KilolitersPerDay()
    case VolumeFlowMegaliterPerDay:
		return a.MegalitersPerDay()
    case VolumeFlowMegaukGallonPerDay:
		return a.MegaukGallonsPerDay()
    case VolumeFlowMegaukGallonPerSecond:
		return a.MegaukGallonsPerSecond()
	default:
		return math.NaN()
	}
}

func (a *VolumeFlow) convertFromBase(toUnit VolumeFlowUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumeFlowCubicMeterPerSecond:
		return (value) 
	case VolumeFlowCubicMeterPerMinute:
		return (value * 60) 
	case VolumeFlowCubicMeterPerHour:
		return (value * 3600) 
	case VolumeFlowCubicMeterPerDay:
		return (value * 86400) 
	case VolumeFlowCubicFootPerSecond:
		return (value * 35.314666721) 
	case VolumeFlowCubicFootPerMinute:
		return (value * 2118.88000326) 
	case VolumeFlowCubicFootPerHour:
		return (value / 7.8657907199999087346816086183876e-6) 
	case VolumeFlowCubicYardPerSecond:
		return (value / 0.764554857984) 
	case VolumeFlowCubicYardPerMinute:
		return (value / 0.0127425809664) 
	case VolumeFlowCubicYardPerHour:
		return (value / 2.1237634944e-4) 
	case VolumeFlowCubicYardPerDay:
		return (value * 113007) 
	case VolumeFlowMillionUsGallonPerDay:
		return (value * 22.824465227) 
	case VolumeFlowUsGallonPerDay:
		return (value * 22824465.227) 
	case VolumeFlowLiterPerSecond:
		return (value * 1000) 
	case VolumeFlowLiterPerMinute:
		return (value * 60000.00000) 
	case VolumeFlowLiterPerHour:
		return (value * 3600000.000) 
	case VolumeFlowLiterPerDay:
		return (value * 86400000) 
	case VolumeFlowUsGallonPerSecond:
		return (value * 264.1720523581484) 
	case VolumeFlowUsGallonPerMinute:
		return (value * 15850.323141489) 
	case VolumeFlowUkGallonPerDay:
		return (value * 19005304) 
	case VolumeFlowUkGallonPerHour:
		return (value * 791887.667) 
	case VolumeFlowUkGallonPerMinute:
		return (value * 13198.2) 
	case VolumeFlowUkGallonPerSecond:
		return (value * 219.969) 
	case VolumeFlowKilousGallonPerMinute:
		return (value * 15.850323141489) 
	case VolumeFlowUsGallonPerHour:
		return (value * 951019.38848933424) 
	case VolumeFlowCubicDecimeterPerMinute:
		return (value * 60000.00000) 
	case VolumeFlowOilBarrelPerDay:
		return (value / 1.8401307283333333333333333333333e-6) 
	case VolumeFlowOilBarrelPerMinute:
		return (value / 2.64978825e-3) 
	case VolumeFlowOilBarrelPerHour:
		return (value / 4.41631375e-5) 
	case VolumeFlowOilBarrelPerSecond:
		return (value * 6.28981) 
	case VolumeFlowCubicMillimeterPerSecond:
		return (value / 1e-9) 
	case VolumeFlowAcreFootPerSecond:
		return (value * 0.000810713194) 
	case VolumeFlowAcreFootPerMinute:
		return (value * 0.0486427916) 
	case VolumeFlowAcreFootPerHour:
		return (value * 2.91857) 
	case VolumeFlowAcreFootPerDay:
		return (value * 70.0457) 
	case VolumeFlowCubicCentimeterPerMinute:
		return (value / 1.6666666666667e-8) 
	case VolumeFlowMegausGallonPerDay:
		return ((value * 22824465.227) / 1000000.0) 
	case VolumeFlowNanoliterPerSecond:
		return ((value * 1000) / 1e-09) 
	case VolumeFlowMicroliterPerSecond:
		return ((value * 1000) / 1e-06) 
	case VolumeFlowMilliliterPerSecond:
		return ((value * 1000) / 0.001) 
	case VolumeFlowCentiliterPerSecond:
		return ((value * 1000) / 0.01) 
	case VolumeFlowDeciliterPerSecond:
		return ((value * 1000) / 0.1) 
	case VolumeFlowDecaliterPerSecond:
		return ((value * 1000) / 10.0) 
	case VolumeFlowHectoliterPerSecond:
		return ((value * 1000) / 100.0) 
	case VolumeFlowKiloliterPerSecond:
		return ((value * 1000) / 1000.0) 
	case VolumeFlowMegaliterPerSecond:
		return ((value * 1000) / 1000000.0) 
	case VolumeFlowNanoliterPerMinute:
		return ((value * 60000.00000) / 1e-09) 
	case VolumeFlowMicroliterPerMinute:
		return ((value * 60000.00000) / 1e-06) 
	case VolumeFlowMilliliterPerMinute:
		return ((value * 60000.00000) / 0.001) 
	case VolumeFlowCentiliterPerMinute:
		return ((value * 60000.00000) / 0.01) 
	case VolumeFlowDeciliterPerMinute:
		return ((value * 60000.00000) / 0.1) 
	case VolumeFlowDecaliterPerMinute:
		return ((value * 60000.00000) / 10.0) 
	case VolumeFlowHectoliterPerMinute:
		return ((value * 60000.00000) / 100.0) 
	case VolumeFlowKiloliterPerMinute:
		return ((value * 60000.00000) / 1000.0) 
	case VolumeFlowMegaliterPerMinute:
		return ((value * 60000.00000) / 1000000.0) 
	case VolumeFlowNanoliterPerHour:
		return ((value * 3600000.000) / 1e-09) 
	case VolumeFlowMicroliterPerHour:
		return ((value * 3600000.000) / 1e-06) 
	case VolumeFlowMilliliterPerHour:
		return ((value * 3600000.000) / 0.001) 
	case VolumeFlowCentiliterPerHour:
		return ((value * 3600000.000) / 0.01) 
	case VolumeFlowDeciliterPerHour:
		return ((value * 3600000.000) / 0.1) 
	case VolumeFlowDecaliterPerHour:
		return ((value * 3600000.000) / 10.0) 
	case VolumeFlowHectoliterPerHour:
		return ((value * 3600000.000) / 100.0) 
	case VolumeFlowKiloliterPerHour:
		return ((value * 3600000.000) / 1000.0) 
	case VolumeFlowMegaliterPerHour:
		return ((value * 3600000.000) / 1000000.0) 
	case VolumeFlowNanoliterPerDay:
		return ((value * 86400000) / 1e-09) 
	case VolumeFlowMicroliterPerDay:
		return ((value * 86400000) / 1e-06) 
	case VolumeFlowMilliliterPerDay:
		return ((value * 86400000) / 0.001) 
	case VolumeFlowCentiliterPerDay:
		return ((value * 86400000) / 0.01) 
	case VolumeFlowDeciliterPerDay:
		return ((value * 86400000) / 0.1) 
	case VolumeFlowDecaliterPerDay:
		return ((value * 86400000) / 10.0) 
	case VolumeFlowHectoliterPerDay:
		return ((value * 86400000) / 100.0) 
	case VolumeFlowKiloliterPerDay:
		return ((value * 86400000) / 1000.0) 
	case VolumeFlowMegaliterPerDay:
		return ((value * 86400000) / 1000000.0) 
	case VolumeFlowMegaukGallonPerDay:
		return ((value * 19005304) / 1000000.0) 
	case VolumeFlowMegaukGallonPerSecond:
		return ((value * 219.969) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *VolumeFlow) convertToBase(value float64, fromUnit VolumeFlowUnits) float64 {
	switch fromUnit { 
	case VolumeFlowCubicMeterPerSecond:
		return (value) 
	case VolumeFlowCubicMeterPerMinute:
		return (value / 60) 
	case VolumeFlowCubicMeterPerHour:
		return (value / 3600) 
	case VolumeFlowCubicMeterPerDay:
		return (value / 86400) 
	case VolumeFlowCubicFootPerSecond:
		return (value / 35.314666721) 
	case VolumeFlowCubicFootPerMinute:
		return (value / 2118.88000326) 
	case VolumeFlowCubicFootPerHour:
		return (value * 7.8657907199999087346816086183876e-6) 
	case VolumeFlowCubicYardPerSecond:
		return (value * 0.764554857984) 
	case VolumeFlowCubicYardPerMinute:
		return (value * 0.0127425809664) 
	case VolumeFlowCubicYardPerHour:
		return (value * 2.1237634944e-4) 
	case VolumeFlowCubicYardPerDay:
		return (value / 113007) 
	case VolumeFlowMillionUsGallonPerDay:
		return (value / 22.824465227) 
	case VolumeFlowUsGallonPerDay:
		return (value / 22824465.227) 
	case VolumeFlowLiterPerSecond:
		return (value / 1000) 
	case VolumeFlowLiterPerMinute:
		return (value / 60000.00000) 
	case VolumeFlowLiterPerHour:
		return (value / 3600000.000) 
	case VolumeFlowLiterPerDay:
		return (value / 86400000) 
	case VolumeFlowUsGallonPerSecond:
		return (value / 264.1720523581484) 
	case VolumeFlowUsGallonPerMinute:
		return (value / 15850.323141489) 
	case VolumeFlowUkGallonPerDay:
		return (value / 19005304) 
	case VolumeFlowUkGallonPerHour:
		return (value / 791887.667) 
	case VolumeFlowUkGallonPerMinute:
		return (value / 13198.2) 
	case VolumeFlowUkGallonPerSecond:
		return (value / 219.969) 
	case VolumeFlowKilousGallonPerMinute:
		return (value / 15.850323141489) 
	case VolumeFlowUsGallonPerHour:
		return (value / 951019.38848933424) 
	case VolumeFlowCubicDecimeterPerMinute:
		return (value / 60000.00000) 
	case VolumeFlowOilBarrelPerDay:
		return (value * 1.8401307283333333333333333333333e-6) 
	case VolumeFlowOilBarrelPerMinute:
		return (value * 2.64978825e-3) 
	case VolumeFlowOilBarrelPerHour:
		return (value * 4.41631375e-5) 
	case VolumeFlowOilBarrelPerSecond:
		return (value / 6.28981) 
	case VolumeFlowCubicMillimeterPerSecond:
		return (value * 1e-9) 
	case VolumeFlowAcreFootPerSecond:
		return (value / 0.000810713194) 
	case VolumeFlowAcreFootPerMinute:
		return (value / 0.0486427916) 
	case VolumeFlowAcreFootPerHour:
		return (value / 2.91857) 
	case VolumeFlowAcreFootPerDay:
		return (value / 70.0457) 
	case VolumeFlowCubicCentimeterPerMinute:
		return (value * 1.6666666666667e-8) 
	case VolumeFlowMegausGallonPerDay:
		return ((value / 22824465.227) * 1000000.0) 
	case VolumeFlowNanoliterPerSecond:
		return ((value / 1000) * 1e-09) 
	case VolumeFlowMicroliterPerSecond:
		return ((value / 1000) * 1e-06) 
	case VolumeFlowMilliliterPerSecond:
		return ((value / 1000) * 0.001) 
	case VolumeFlowCentiliterPerSecond:
		return ((value / 1000) * 0.01) 
	case VolumeFlowDeciliterPerSecond:
		return ((value / 1000) * 0.1) 
	case VolumeFlowDecaliterPerSecond:
		return ((value / 1000) * 10.0) 
	case VolumeFlowHectoliterPerSecond:
		return ((value / 1000) * 100.0) 
	case VolumeFlowKiloliterPerSecond:
		return ((value / 1000) * 1000.0) 
	case VolumeFlowMegaliterPerSecond:
		return ((value / 1000) * 1000000.0) 
	case VolumeFlowNanoliterPerMinute:
		return ((value / 60000.00000) * 1e-09) 
	case VolumeFlowMicroliterPerMinute:
		return ((value / 60000.00000) * 1e-06) 
	case VolumeFlowMilliliterPerMinute:
		return ((value / 60000.00000) * 0.001) 
	case VolumeFlowCentiliterPerMinute:
		return ((value / 60000.00000) * 0.01) 
	case VolumeFlowDeciliterPerMinute:
		return ((value / 60000.00000) * 0.1) 
	case VolumeFlowDecaliterPerMinute:
		return ((value / 60000.00000) * 10.0) 
	case VolumeFlowHectoliterPerMinute:
		return ((value / 60000.00000) * 100.0) 
	case VolumeFlowKiloliterPerMinute:
		return ((value / 60000.00000) * 1000.0) 
	case VolumeFlowMegaliterPerMinute:
		return ((value / 60000.00000) * 1000000.0) 
	case VolumeFlowNanoliterPerHour:
		return ((value / 3600000.000) * 1e-09) 
	case VolumeFlowMicroliterPerHour:
		return ((value / 3600000.000) * 1e-06) 
	case VolumeFlowMilliliterPerHour:
		return ((value / 3600000.000) * 0.001) 
	case VolumeFlowCentiliterPerHour:
		return ((value / 3600000.000) * 0.01) 
	case VolumeFlowDeciliterPerHour:
		return ((value / 3600000.000) * 0.1) 
	case VolumeFlowDecaliterPerHour:
		return ((value / 3600000.000) * 10.0) 
	case VolumeFlowHectoliterPerHour:
		return ((value / 3600000.000) * 100.0) 
	case VolumeFlowKiloliterPerHour:
		return ((value / 3600000.000) * 1000.0) 
	case VolumeFlowMegaliterPerHour:
		return ((value / 3600000.000) * 1000000.0) 
	case VolumeFlowNanoliterPerDay:
		return ((value / 86400000) * 1e-09) 
	case VolumeFlowMicroliterPerDay:
		return ((value / 86400000) * 1e-06) 
	case VolumeFlowMilliliterPerDay:
		return ((value / 86400000) * 0.001) 
	case VolumeFlowCentiliterPerDay:
		return ((value / 86400000) * 0.01) 
	case VolumeFlowDeciliterPerDay:
		return ((value / 86400000) * 0.1) 
	case VolumeFlowDecaliterPerDay:
		return ((value / 86400000) * 10.0) 
	case VolumeFlowHectoliterPerDay:
		return ((value / 86400000) * 100.0) 
	case VolumeFlowKiloliterPerDay:
		return ((value / 86400000) * 1000.0) 
	case VolumeFlowMegaliterPerDay:
		return ((value / 86400000) * 1000000.0) 
	case VolumeFlowMegaukGallonPerDay:
		return ((value / 19005304) * 1000000.0) 
	case VolumeFlowMegaukGallonPerSecond:
		return ((value / 219.969) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VolumeFlow in the default unit (CubicMeterPerSecond),
// formatted to two decimal places.
func (a VolumeFlow) String() string {
	return a.ToString(VolumeFlowCubicMeterPerSecond, 2)
}

// ToString formats the VolumeFlow value as a string with the specified unit and fractional digits.
// It converts the VolumeFlow to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VolumeFlow value will be converted (e.g., CubicMeterPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VolumeFlow with the unit abbreviation.
func (a *VolumeFlow) ToString(unit VolumeFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetVolumeFlowAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetVolumeFlowAbbreviation(unit))
}

// Equals checks if the given VolumeFlow is equal to the current VolumeFlow.
//
// Parameters:
//    other: The VolumeFlow to compare against.
//
// Returns:
//    bool: Returns true if both VolumeFlow are equal, false otherwise.
func (a *VolumeFlow) Equals(other *VolumeFlow) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VolumeFlow with another VolumeFlow.
// It returns -1 if the current VolumeFlow is less than the other VolumeFlow, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VolumeFlow to compare against.
//
// Returns:
//    int: -1 if the current VolumeFlow is less, 1 if greater, and 0 if equal.
func (a *VolumeFlow) CompareTo(other *VolumeFlow) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VolumeFlow to the current VolumeFlow and returns the result.
// The result is a new VolumeFlow instance with the sum of the values.
//
// Parameters:
//    other: The VolumeFlow to add to the current VolumeFlow.
//
// Returns:
//    *VolumeFlow: A new VolumeFlow instance representing the sum of both VolumeFlow.
func (a *VolumeFlow) Add(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VolumeFlow from the current VolumeFlow and returns the result.
// The result is a new VolumeFlow instance with the difference of the values.
//
// Parameters:
//    other: The VolumeFlow to subtract from the current VolumeFlow.
//
// Returns:
//    *VolumeFlow: A new VolumeFlow instance representing the difference of both VolumeFlow.
func (a *VolumeFlow) Subtract(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VolumeFlow by the given VolumeFlow and returns the result.
// The result is a new VolumeFlow instance with the product of the values.
//
// Parameters:
//    other: The VolumeFlow to multiply with the current VolumeFlow.
//
// Returns:
//    *VolumeFlow: A new VolumeFlow instance representing the product of both VolumeFlow.
func (a *VolumeFlow) Multiply(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value * other.BaseValue()}
}

// Divide divides the current VolumeFlow by the given VolumeFlow and returns the result.
// The result is a new VolumeFlow instance with the quotient of the values.
//
// Parameters:
//    other: The VolumeFlow to divide the current VolumeFlow by.
//
// Returns:
//    *VolumeFlow: A new VolumeFlow instance representing the quotient of both VolumeFlow.
func (a *VolumeFlow) Divide(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value / other.BaseValue()}
}

// GetVolumeFlowAbbreviation gets the unit abbreviation.
func GetVolumeFlowAbbreviation(unit VolumeFlowUnits) string {
	switch unit { 
	case VolumeFlowCubicMeterPerSecond:
		return "m³/s" 
	case VolumeFlowCubicMeterPerMinute:
		return "m³/min" 
	case VolumeFlowCubicMeterPerHour:
		return "m³/h" 
	case VolumeFlowCubicMeterPerDay:
		return "m³/d" 
	case VolumeFlowCubicFootPerSecond:
		return "ft³/s" 
	case VolumeFlowCubicFootPerMinute:
		return "ft³/min" 
	case VolumeFlowCubicFootPerHour:
		return "ft³/h" 
	case VolumeFlowCubicYardPerSecond:
		return "yd³/s" 
	case VolumeFlowCubicYardPerMinute:
		return "yd³/min" 
	case VolumeFlowCubicYardPerHour:
		return "yd³/h" 
	case VolumeFlowCubicYardPerDay:
		return "cy/day" 
	case VolumeFlowMillionUsGallonPerDay:
		return "MGD" 
	case VolumeFlowUsGallonPerDay:
		return "gpd" 
	case VolumeFlowLiterPerSecond:
		return "L/s" 
	case VolumeFlowLiterPerMinute:
		return "L/min" 
	case VolumeFlowLiterPerHour:
		return "L/h" 
	case VolumeFlowLiterPerDay:
		return "l/day" 
	case VolumeFlowUsGallonPerSecond:
		return "gal (U.S.)/s" 
	case VolumeFlowUsGallonPerMinute:
		return "gal (U.S.)/min" 
	case VolumeFlowUkGallonPerDay:
		return "gal (U. K.)/d" 
	case VolumeFlowUkGallonPerHour:
		return "gal (imp.)/h" 
	case VolumeFlowUkGallonPerMinute:
		return "gal (imp.)/min" 
	case VolumeFlowUkGallonPerSecond:
		return "gal (imp.)/s" 
	case VolumeFlowKilousGallonPerMinute:
		return "kgal (U.S.)/min" 
	case VolumeFlowUsGallonPerHour:
		return "gal (U.S.)/h" 
	case VolumeFlowCubicDecimeterPerMinute:
		return "dm³/min" 
	case VolumeFlowOilBarrelPerDay:
		return "bbl/d" 
	case VolumeFlowOilBarrelPerMinute:
		return "bbl/min" 
	case VolumeFlowOilBarrelPerHour:
		return "bbl/hr" 
	case VolumeFlowOilBarrelPerSecond:
		return "bbl/s" 
	case VolumeFlowCubicMillimeterPerSecond:
		return "mm³/s" 
	case VolumeFlowAcreFootPerSecond:
		return "af/s" 
	case VolumeFlowAcreFootPerMinute:
		return "af/m" 
	case VolumeFlowAcreFootPerHour:
		return "af/h" 
	case VolumeFlowAcreFootPerDay:
		return "af/d" 
	case VolumeFlowCubicCentimeterPerMinute:
		return "cm³/min" 
	case VolumeFlowMegausGallonPerDay:
		return "Mgpd" 
	case VolumeFlowNanoliterPerSecond:
		return "nL/s" 
	case VolumeFlowMicroliterPerSecond:
		return "μL/s" 
	case VolumeFlowMilliliterPerSecond:
		return "mL/s" 
	case VolumeFlowCentiliterPerSecond:
		return "cL/s" 
	case VolumeFlowDeciliterPerSecond:
		return "dL/s" 
	case VolumeFlowDecaliterPerSecond:
		return "daL/s" 
	case VolumeFlowHectoliterPerSecond:
		return "hL/s" 
	case VolumeFlowKiloliterPerSecond:
		return "kL/s" 
	case VolumeFlowMegaliterPerSecond:
		return "ML/s" 
	case VolumeFlowNanoliterPerMinute:
		return "nL/min" 
	case VolumeFlowMicroliterPerMinute:
		return "μL/min" 
	case VolumeFlowMilliliterPerMinute:
		return "mL/min" 
	case VolumeFlowCentiliterPerMinute:
		return "cL/min" 
	case VolumeFlowDeciliterPerMinute:
		return "dL/min" 
	case VolumeFlowDecaliterPerMinute:
		return "daL/min" 
	case VolumeFlowHectoliterPerMinute:
		return "hL/min" 
	case VolumeFlowKiloliterPerMinute:
		return "kL/min" 
	case VolumeFlowMegaliterPerMinute:
		return "ML/min" 
	case VolumeFlowNanoliterPerHour:
		return "nL/h" 
	case VolumeFlowMicroliterPerHour:
		return "μL/h" 
	case VolumeFlowMilliliterPerHour:
		return "mL/h" 
	case VolumeFlowCentiliterPerHour:
		return "cL/h" 
	case VolumeFlowDeciliterPerHour:
		return "dL/h" 
	case VolumeFlowDecaliterPerHour:
		return "daL/h" 
	case VolumeFlowHectoliterPerHour:
		return "hL/h" 
	case VolumeFlowKiloliterPerHour:
		return "kL/h" 
	case VolumeFlowMegaliterPerHour:
		return "ML/h" 
	case VolumeFlowNanoliterPerDay:
		return "nl/day" 
	case VolumeFlowMicroliterPerDay:
		return "μl/day" 
	case VolumeFlowMilliliterPerDay:
		return "ml/day" 
	case VolumeFlowCentiliterPerDay:
		return "cl/day" 
	case VolumeFlowDeciliterPerDay:
		return "dl/day" 
	case VolumeFlowDecaliterPerDay:
		return "dal/day" 
	case VolumeFlowHectoliterPerDay:
		return "hl/day" 
	case VolumeFlowKiloliterPerDay:
		return "kl/day" 
	case VolumeFlowMegaliterPerDay:
		return "Ml/day" 
	case VolumeFlowMegaukGallonPerDay:
		return "Mgal (U. K.)/d" 
	case VolumeFlowMegaukGallonPerSecond:
		return "Mgal (imp.)/s" 
	default:
		return ""
	}
}