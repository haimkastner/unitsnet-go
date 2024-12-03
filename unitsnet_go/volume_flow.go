package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeFlowUnits enumeration
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

// VolumeFlowDto represents an VolumeFlow
type VolumeFlowDto struct {
	Value float64
	Unit  VolumeFlowUnits
}

// VolumeFlowDtoFactory struct to group related functions
type VolumeFlowDtoFactory struct{}

func (udf VolumeFlowDtoFactory) FromJSON(data []byte) (*VolumeFlowDto, error) {
	a := VolumeFlowDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumeFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VolumeFlow struct
type VolumeFlow struct {
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

// VolumeFlowFactory struct to group related functions
type VolumeFlowFactory struct{}

func (uf VolumeFlowFactory) CreateVolumeFlow(value float64, unit VolumeFlowUnits) (*VolumeFlow, error) {
	return newVolumeFlow(value, unit)
}

func (uf VolumeFlowFactory) FromDto(dto VolumeFlowDto) (*VolumeFlow, error) {
	return newVolumeFlow(dto.Value, dto.Unit)
}

func (uf VolumeFlowFactory) FromDtoJSON(data []byte) (*VolumeFlow, error) {
	unitDto, err := VolumeFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumeFlowFactory{}.FromDto(*unitDto)
}


// FromCubicMeterPerSecond creates a new VolumeFlow instance from CubicMeterPerSecond.
func (uf VolumeFlowFactory) FromCubicMetersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerSecond)
}

// FromCubicMeterPerMinute creates a new VolumeFlow instance from CubicMeterPerMinute.
func (uf VolumeFlowFactory) FromCubicMetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerMinute)
}

// FromCubicMeterPerHour creates a new VolumeFlow instance from CubicMeterPerHour.
func (uf VolumeFlowFactory) FromCubicMetersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerHour)
}

// FromCubicMeterPerDay creates a new VolumeFlow instance from CubicMeterPerDay.
func (uf VolumeFlowFactory) FromCubicMetersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMeterPerDay)
}

// FromCubicFootPerSecond creates a new VolumeFlow instance from CubicFootPerSecond.
func (uf VolumeFlowFactory) FromCubicFeetPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerSecond)
}

// FromCubicFootPerMinute creates a new VolumeFlow instance from CubicFootPerMinute.
func (uf VolumeFlowFactory) FromCubicFeetPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerMinute)
}

// FromCubicFootPerHour creates a new VolumeFlow instance from CubicFootPerHour.
func (uf VolumeFlowFactory) FromCubicFeetPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicFootPerHour)
}

// FromCubicYardPerSecond creates a new VolumeFlow instance from CubicYardPerSecond.
func (uf VolumeFlowFactory) FromCubicYardsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerSecond)
}

// FromCubicYardPerMinute creates a new VolumeFlow instance from CubicYardPerMinute.
func (uf VolumeFlowFactory) FromCubicYardsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerMinute)
}

// FromCubicYardPerHour creates a new VolumeFlow instance from CubicYardPerHour.
func (uf VolumeFlowFactory) FromCubicYardsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerHour)
}

// FromCubicYardPerDay creates a new VolumeFlow instance from CubicYardPerDay.
func (uf VolumeFlowFactory) FromCubicYardsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicYardPerDay)
}

// FromMillionUsGallonPerDay creates a new VolumeFlow instance from MillionUsGallonPerDay.
func (uf VolumeFlowFactory) FromMillionUsGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMillionUsGallonPerDay)
}

// FromUsGallonPerDay creates a new VolumeFlow instance from UsGallonPerDay.
func (uf VolumeFlowFactory) FromUsGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerDay)
}

// FromLiterPerSecond creates a new VolumeFlow instance from LiterPerSecond.
func (uf VolumeFlowFactory) FromLitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerSecond)
}

// FromLiterPerMinute creates a new VolumeFlow instance from LiterPerMinute.
func (uf VolumeFlowFactory) FromLitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerMinute)
}

// FromLiterPerHour creates a new VolumeFlow instance from LiterPerHour.
func (uf VolumeFlowFactory) FromLitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerHour)
}

// FromLiterPerDay creates a new VolumeFlow instance from LiterPerDay.
func (uf VolumeFlowFactory) FromLitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowLiterPerDay)
}

// FromUsGallonPerSecond creates a new VolumeFlow instance from UsGallonPerSecond.
func (uf VolumeFlowFactory) FromUsGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerSecond)
}

// FromUsGallonPerMinute creates a new VolumeFlow instance from UsGallonPerMinute.
func (uf VolumeFlowFactory) FromUsGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerMinute)
}

// FromUkGallonPerDay creates a new VolumeFlow instance from UkGallonPerDay.
func (uf VolumeFlowFactory) FromUkGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerDay)
}

// FromUkGallonPerHour creates a new VolumeFlow instance from UkGallonPerHour.
func (uf VolumeFlowFactory) FromUkGallonsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerHour)
}

// FromUkGallonPerMinute creates a new VolumeFlow instance from UkGallonPerMinute.
func (uf VolumeFlowFactory) FromUkGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerMinute)
}

// FromUkGallonPerSecond creates a new VolumeFlow instance from UkGallonPerSecond.
func (uf VolumeFlowFactory) FromUkGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUkGallonPerSecond)
}

// FromKilousGallonPerMinute creates a new VolumeFlow instance from KilousGallonPerMinute.
func (uf VolumeFlowFactory) FromKilousGallonsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKilousGallonPerMinute)
}

// FromUsGallonPerHour creates a new VolumeFlow instance from UsGallonPerHour.
func (uf VolumeFlowFactory) FromUsGallonsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowUsGallonPerHour)
}

// FromCubicDecimeterPerMinute creates a new VolumeFlow instance from CubicDecimeterPerMinute.
func (uf VolumeFlowFactory) FromCubicDecimetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicDecimeterPerMinute)
}

// FromOilBarrelPerDay creates a new VolumeFlow instance from OilBarrelPerDay.
func (uf VolumeFlowFactory) FromOilBarrelsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerDay)
}

// FromOilBarrelPerMinute creates a new VolumeFlow instance from OilBarrelPerMinute.
func (uf VolumeFlowFactory) FromOilBarrelsPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerMinute)
}

// FromOilBarrelPerHour creates a new VolumeFlow instance from OilBarrelPerHour.
func (uf VolumeFlowFactory) FromOilBarrelsPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerHour)
}

// FromOilBarrelPerSecond creates a new VolumeFlow instance from OilBarrelPerSecond.
func (uf VolumeFlowFactory) FromOilBarrelsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowOilBarrelPerSecond)
}

// FromCubicMillimeterPerSecond creates a new VolumeFlow instance from CubicMillimeterPerSecond.
func (uf VolumeFlowFactory) FromCubicMillimetersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicMillimeterPerSecond)
}

// FromAcreFootPerSecond creates a new VolumeFlow instance from AcreFootPerSecond.
func (uf VolumeFlowFactory) FromAcreFeetPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerSecond)
}

// FromAcreFootPerMinute creates a new VolumeFlow instance from AcreFootPerMinute.
func (uf VolumeFlowFactory) FromAcreFeetPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerMinute)
}

// FromAcreFootPerHour creates a new VolumeFlow instance from AcreFootPerHour.
func (uf VolumeFlowFactory) FromAcreFeetPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerHour)
}

// FromAcreFootPerDay creates a new VolumeFlow instance from AcreFootPerDay.
func (uf VolumeFlowFactory) FromAcreFeetPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowAcreFootPerDay)
}

// FromCubicCentimeterPerMinute creates a new VolumeFlow instance from CubicCentimeterPerMinute.
func (uf VolumeFlowFactory) FromCubicCentimetersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCubicCentimeterPerMinute)
}

// FromMegausGallonPerDay creates a new VolumeFlow instance from MegausGallonPerDay.
func (uf VolumeFlowFactory) FromMegausGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegausGallonPerDay)
}

// FromNanoliterPerSecond creates a new VolumeFlow instance from NanoliterPerSecond.
func (uf VolumeFlowFactory) FromNanolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerSecond)
}

// FromMicroliterPerSecond creates a new VolumeFlow instance from MicroliterPerSecond.
func (uf VolumeFlowFactory) FromMicrolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerSecond)
}

// FromMilliliterPerSecond creates a new VolumeFlow instance from MilliliterPerSecond.
func (uf VolumeFlowFactory) FromMillilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerSecond)
}

// FromCentiliterPerSecond creates a new VolumeFlow instance from CentiliterPerSecond.
func (uf VolumeFlowFactory) FromCentilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerSecond)
}

// FromDeciliterPerSecond creates a new VolumeFlow instance from DeciliterPerSecond.
func (uf VolumeFlowFactory) FromDecilitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerSecond)
}

// FromDecaliterPerSecond creates a new VolumeFlow instance from DecaliterPerSecond.
func (uf VolumeFlowFactory) FromDecalitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerSecond)
}

// FromHectoliterPerSecond creates a new VolumeFlow instance from HectoliterPerSecond.
func (uf VolumeFlowFactory) FromHectolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerSecond)
}

// FromKiloliterPerSecond creates a new VolumeFlow instance from KiloliterPerSecond.
func (uf VolumeFlowFactory) FromKilolitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerSecond)
}

// FromMegaliterPerSecond creates a new VolumeFlow instance from MegaliterPerSecond.
func (uf VolumeFlowFactory) FromMegalitersPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerSecond)
}

// FromNanoliterPerMinute creates a new VolumeFlow instance from NanoliterPerMinute.
func (uf VolumeFlowFactory) FromNanolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerMinute)
}

// FromMicroliterPerMinute creates a new VolumeFlow instance from MicroliterPerMinute.
func (uf VolumeFlowFactory) FromMicrolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerMinute)
}

// FromMilliliterPerMinute creates a new VolumeFlow instance from MilliliterPerMinute.
func (uf VolumeFlowFactory) FromMillilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerMinute)
}

// FromCentiliterPerMinute creates a new VolumeFlow instance from CentiliterPerMinute.
func (uf VolumeFlowFactory) FromCentilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerMinute)
}

// FromDeciliterPerMinute creates a new VolumeFlow instance from DeciliterPerMinute.
func (uf VolumeFlowFactory) FromDecilitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerMinute)
}

// FromDecaliterPerMinute creates a new VolumeFlow instance from DecaliterPerMinute.
func (uf VolumeFlowFactory) FromDecalitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerMinute)
}

// FromHectoliterPerMinute creates a new VolumeFlow instance from HectoliterPerMinute.
func (uf VolumeFlowFactory) FromHectolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerMinute)
}

// FromKiloliterPerMinute creates a new VolumeFlow instance from KiloliterPerMinute.
func (uf VolumeFlowFactory) FromKilolitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerMinute)
}

// FromMegaliterPerMinute creates a new VolumeFlow instance from MegaliterPerMinute.
func (uf VolumeFlowFactory) FromMegalitersPerMinute(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerMinute)
}

// FromNanoliterPerHour creates a new VolumeFlow instance from NanoliterPerHour.
func (uf VolumeFlowFactory) FromNanolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerHour)
}

// FromMicroliterPerHour creates a new VolumeFlow instance from MicroliterPerHour.
func (uf VolumeFlowFactory) FromMicrolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerHour)
}

// FromMilliliterPerHour creates a new VolumeFlow instance from MilliliterPerHour.
func (uf VolumeFlowFactory) FromMillilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerHour)
}

// FromCentiliterPerHour creates a new VolumeFlow instance from CentiliterPerHour.
func (uf VolumeFlowFactory) FromCentilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerHour)
}

// FromDeciliterPerHour creates a new VolumeFlow instance from DeciliterPerHour.
func (uf VolumeFlowFactory) FromDecilitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerHour)
}

// FromDecaliterPerHour creates a new VolumeFlow instance from DecaliterPerHour.
func (uf VolumeFlowFactory) FromDecalitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerHour)
}

// FromHectoliterPerHour creates a new VolumeFlow instance from HectoliterPerHour.
func (uf VolumeFlowFactory) FromHectolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerHour)
}

// FromKiloliterPerHour creates a new VolumeFlow instance from KiloliterPerHour.
func (uf VolumeFlowFactory) FromKilolitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerHour)
}

// FromMegaliterPerHour creates a new VolumeFlow instance from MegaliterPerHour.
func (uf VolumeFlowFactory) FromMegalitersPerHour(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerHour)
}

// FromNanoliterPerDay creates a new VolumeFlow instance from NanoliterPerDay.
func (uf VolumeFlowFactory) FromNanolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowNanoliterPerDay)
}

// FromMicroliterPerDay creates a new VolumeFlow instance from MicroliterPerDay.
func (uf VolumeFlowFactory) FromMicrolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMicroliterPerDay)
}

// FromMilliliterPerDay creates a new VolumeFlow instance from MilliliterPerDay.
func (uf VolumeFlowFactory) FromMillilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMilliliterPerDay)
}

// FromCentiliterPerDay creates a new VolumeFlow instance from CentiliterPerDay.
func (uf VolumeFlowFactory) FromCentilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowCentiliterPerDay)
}

// FromDeciliterPerDay creates a new VolumeFlow instance from DeciliterPerDay.
func (uf VolumeFlowFactory) FromDecilitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDeciliterPerDay)
}

// FromDecaliterPerDay creates a new VolumeFlow instance from DecaliterPerDay.
func (uf VolumeFlowFactory) FromDecalitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowDecaliterPerDay)
}

// FromHectoliterPerDay creates a new VolumeFlow instance from HectoliterPerDay.
func (uf VolumeFlowFactory) FromHectolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowHectoliterPerDay)
}

// FromKiloliterPerDay creates a new VolumeFlow instance from KiloliterPerDay.
func (uf VolumeFlowFactory) FromKilolitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowKiloliterPerDay)
}

// FromMegaliterPerDay creates a new VolumeFlow instance from MegaliterPerDay.
func (uf VolumeFlowFactory) FromMegalitersPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaliterPerDay)
}

// FromMegaukGallonPerDay creates a new VolumeFlow instance from MegaukGallonPerDay.
func (uf VolumeFlowFactory) FromMegaukGallonsPerDay(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaukGallonPerDay)
}

// FromMegaukGallonPerSecond creates a new VolumeFlow instance from MegaukGallonPerSecond.
func (uf VolumeFlowFactory) FromMegaukGallonsPerSecond(value float64) (*VolumeFlow, error) {
	return newVolumeFlow(value, VolumeFlowMegaukGallonPerSecond)
}




// newVolumeFlow creates a new VolumeFlow.
func newVolumeFlow(value float64, fromUnit VolumeFlowUnits) (*VolumeFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VolumeFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumeFlow in CubicMeterPerSecond.
func (a *VolumeFlow) BaseValue() float64 {
	return a.value
}


// CubicMeterPerSecond returns the value in CubicMeterPerSecond.
func (a *VolumeFlow) CubicMetersPerSecond() float64 {
	if a.cubic_meters_per_secondLazy != nil {
		return *a.cubic_meters_per_secondLazy
	}
	cubic_meters_per_second := a.convertFromBase(VolumeFlowCubicMeterPerSecond)
	a.cubic_meters_per_secondLazy = &cubic_meters_per_second
	return cubic_meters_per_second
}

// CubicMeterPerMinute returns the value in CubicMeterPerMinute.
func (a *VolumeFlow) CubicMetersPerMinute() float64 {
	if a.cubic_meters_per_minuteLazy != nil {
		return *a.cubic_meters_per_minuteLazy
	}
	cubic_meters_per_minute := a.convertFromBase(VolumeFlowCubicMeterPerMinute)
	a.cubic_meters_per_minuteLazy = &cubic_meters_per_minute
	return cubic_meters_per_minute
}

// CubicMeterPerHour returns the value in CubicMeterPerHour.
func (a *VolumeFlow) CubicMetersPerHour() float64 {
	if a.cubic_meters_per_hourLazy != nil {
		return *a.cubic_meters_per_hourLazy
	}
	cubic_meters_per_hour := a.convertFromBase(VolumeFlowCubicMeterPerHour)
	a.cubic_meters_per_hourLazy = &cubic_meters_per_hour
	return cubic_meters_per_hour
}

// CubicMeterPerDay returns the value in CubicMeterPerDay.
func (a *VolumeFlow) CubicMetersPerDay() float64 {
	if a.cubic_meters_per_dayLazy != nil {
		return *a.cubic_meters_per_dayLazy
	}
	cubic_meters_per_day := a.convertFromBase(VolumeFlowCubicMeterPerDay)
	a.cubic_meters_per_dayLazy = &cubic_meters_per_day
	return cubic_meters_per_day
}

// CubicFootPerSecond returns the value in CubicFootPerSecond.
func (a *VolumeFlow) CubicFeetPerSecond() float64 {
	if a.cubic_feet_per_secondLazy != nil {
		return *a.cubic_feet_per_secondLazy
	}
	cubic_feet_per_second := a.convertFromBase(VolumeFlowCubicFootPerSecond)
	a.cubic_feet_per_secondLazy = &cubic_feet_per_second
	return cubic_feet_per_second
}

// CubicFootPerMinute returns the value in CubicFootPerMinute.
func (a *VolumeFlow) CubicFeetPerMinute() float64 {
	if a.cubic_feet_per_minuteLazy != nil {
		return *a.cubic_feet_per_minuteLazy
	}
	cubic_feet_per_minute := a.convertFromBase(VolumeFlowCubicFootPerMinute)
	a.cubic_feet_per_minuteLazy = &cubic_feet_per_minute
	return cubic_feet_per_minute
}

// CubicFootPerHour returns the value in CubicFootPerHour.
func (a *VolumeFlow) CubicFeetPerHour() float64 {
	if a.cubic_feet_per_hourLazy != nil {
		return *a.cubic_feet_per_hourLazy
	}
	cubic_feet_per_hour := a.convertFromBase(VolumeFlowCubicFootPerHour)
	a.cubic_feet_per_hourLazy = &cubic_feet_per_hour
	return cubic_feet_per_hour
}

// CubicYardPerSecond returns the value in CubicYardPerSecond.
func (a *VolumeFlow) CubicYardsPerSecond() float64 {
	if a.cubic_yards_per_secondLazy != nil {
		return *a.cubic_yards_per_secondLazy
	}
	cubic_yards_per_second := a.convertFromBase(VolumeFlowCubicYardPerSecond)
	a.cubic_yards_per_secondLazy = &cubic_yards_per_second
	return cubic_yards_per_second
}

// CubicYardPerMinute returns the value in CubicYardPerMinute.
func (a *VolumeFlow) CubicYardsPerMinute() float64 {
	if a.cubic_yards_per_minuteLazy != nil {
		return *a.cubic_yards_per_minuteLazy
	}
	cubic_yards_per_minute := a.convertFromBase(VolumeFlowCubicYardPerMinute)
	a.cubic_yards_per_minuteLazy = &cubic_yards_per_minute
	return cubic_yards_per_minute
}

// CubicYardPerHour returns the value in CubicYardPerHour.
func (a *VolumeFlow) CubicYardsPerHour() float64 {
	if a.cubic_yards_per_hourLazy != nil {
		return *a.cubic_yards_per_hourLazy
	}
	cubic_yards_per_hour := a.convertFromBase(VolumeFlowCubicYardPerHour)
	a.cubic_yards_per_hourLazy = &cubic_yards_per_hour
	return cubic_yards_per_hour
}

// CubicYardPerDay returns the value in CubicYardPerDay.
func (a *VolumeFlow) CubicYardsPerDay() float64 {
	if a.cubic_yards_per_dayLazy != nil {
		return *a.cubic_yards_per_dayLazy
	}
	cubic_yards_per_day := a.convertFromBase(VolumeFlowCubicYardPerDay)
	a.cubic_yards_per_dayLazy = &cubic_yards_per_day
	return cubic_yards_per_day
}

// MillionUsGallonPerDay returns the value in MillionUsGallonPerDay.
func (a *VolumeFlow) MillionUsGallonsPerDay() float64 {
	if a.million_us_gallons_per_dayLazy != nil {
		return *a.million_us_gallons_per_dayLazy
	}
	million_us_gallons_per_day := a.convertFromBase(VolumeFlowMillionUsGallonPerDay)
	a.million_us_gallons_per_dayLazy = &million_us_gallons_per_day
	return million_us_gallons_per_day
}

// UsGallonPerDay returns the value in UsGallonPerDay.
func (a *VolumeFlow) UsGallonsPerDay() float64 {
	if a.us_gallons_per_dayLazy != nil {
		return *a.us_gallons_per_dayLazy
	}
	us_gallons_per_day := a.convertFromBase(VolumeFlowUsGallonPerDay)
	a.us_gallons_per_dayLazy = &us_gallons_per_day
	return us_gallons_per_day
}

// LiterPerSecond returns the value in LiterPerSecond.
func (a *VolumeFlow) LitersPerSecond() float64 {
	if a.liters_per_secondLazy != nil {
		return *a.liters_per_secondLazy
	}
	liters_per_second := a.convertFromBase(VolumeFlowLiterPerSecond)
	a.liters_per_secondLazy = &liters_per_second
	return liters_per_second
}

// LiterPerMinute returns the value in LiterPerMinute.
func (a *VolumeFlow) LitersPerMinute() float64 {
	if a.liters_per_minuteLazy != nil {
		return *a.liters_per_minuteLazy
	}
	liters_per_minute := a.convertFromBase(VolumeFlowLiterPerMinute)
	a.liters_per_minuteLazy = &liters_per_minute
	return liters_per_minute
}

// LiterPerHour returns the value in LiterPerHour.
func (a *VolumeFlow) LitersPerHour() float64 {
	if a.liters_per_hourLazy != nil {
		return *a.liters_per_hourLazy
	}
	liters_per_hour := a.convertFromBase(VolumeFlowLiterPerHour)
	a.liters_per_hourLazy = &liters_per_hour
	return liters_per_hour
}

// LiterPerDay returns the value in LiterPerDay.
func (a *VolumeFlow) LitersPerDay() float64 {
	if a.liters_per_dayLazy != nil {
		return *a.liters_per_dayLazy
	}
	liters_per_day := a.convertFromBase(VolumeFlowLiterPerDay)
	a.liters_per_dayLazy = &liters_per_day
	return liters_per_day
}

// UsGallonPerSecond returns the value in UsGallonPerSecond.
func (a *VolumeFlow) UsGallonsPerSecond() float64 {
	if a.us_gallons_per_secondLazy != nil {
		return *a.us_gallons_per_secondLazy
	}
	us_gallons_per_second := a.convertFromBase(VolumeFlowUsGallonPerSecond)
	a.us_gallons_per_secondLazy = &us_gallons_per_second
	return us_gallons_per_second
}

// UsGallonPerMinute returns the value in UsGallonPerMinute.
func (a *VolumeFlow) UsGallonsPerMinute() float64 {
	if a.us_gallons_per_minuteLazy != nil {
		return *a.us_gallons_per_minuteLazy
	}
	us_gallons_per_minute := a.convertFromBase(VolumeFlowUsGallonPerMinute)
	a.us_gallons_per_minuteLazy = &us_gallons_per_minute
	return us_gallons_per_minute
}

// UkGallonPerDay returns the value in UkGallonPerDay.
func (a *VolumeFlow) UkGallonsPerDay() float64 {
	if a.uk_gallons_per_dayLazy != nil {
		return *a.uk_gallons_per_dayLazy
	}
	uk_gallons_per_day := a.convertFromBase(VolumeFlowUkGallonPerDay)
	a.uk_gallons_per_dayLazy = &uk_gallons_per_day
	return uk_gallons_per_day
}

// UkGallonPerHour returns the value in UkGallonPerHour.
func (a *VolumeFlow) UkGallonsPerHour() float64 {
	if a.uk_gallons_per_hourLazy != nil {
		return *a.uk_gallons_per_hourLazy
	}
	uk_gallons_per_hour := a.convertFromBase(VolumeFlowUkGallonPerHour)
	a.uk_gallons_per_hourLazy = &uk_gallons_per_hour
	return uk_gallons_per_hour
}

// UkGallonPerMinute returns the value in UkGallonPerMinute.
func (a *VolumeFlow) UkGallonsPerMinute() float64 {
	if a.uk_gallons_per_minuteLazy != nil {
		return *a.uk_gallons_per_minuteLazy
	}
	uk_gallons_per_minute := a.convertFromBase(VolumeFlowUkGallonPerMinute)
	a.uk_gallons_per_minuteLazy = &uk_gallons_per_minute
	return uk_gallons_per_minute
}

// UkGallonPerSecond returns the value in UkGallonPerSecond.
func (a *VolumeFlow) UkGallonsPerSecond() float64 {
	if a.uk_gallons_per_secondLazy != nil {
		return *a.uk_gallons_per_secondLazy
	}
	uk_gallons_per_second := a.convertFromBase(VolumeFlowUkGallonPerSecond)
	a.uk_gallons_per_secondLazy = &uk_gallons_per_second
	return uk_gallons_per_second
}

// KilousGallonPerMinute returns the value in KilousGallonPerMinute.
func (a *VolumeFlow) KilousGallonsPerMinute() float64 {
	if a.kilous_gallons_per_minuteLazy != nil {
		return *a.kilous_gallons_per_minuteLazy
	}
	kilous_gallons_per_minute := a.convertFromBase(VolumeFlowKilousGallonPerMinute)
	a.kilous_gallons_per_minuteLazy = &kilous_gallons_per_minute
	return kilous_gallons_per_minute
}

// UsGallonPerHour returns the value in UsGallonPerHour.
func (a *VolumeFlow) UsGallonsPerHour() float64 {
	if a.us_gallons_per_hourLazy != nil {
		return *a.us_gallons_per_hourLazy
	}
	us_gallons_per_hour := a.convertFromBase(VolumeFlowUsGallonPerHour)
	a.us_gallons_per_hourLazy = &us_gallons_per_hour
	return us_gallons_per_hour
}

// CubicDecimeterPerMinute returns the value in CubicDecimeterPerMinute.
func (a *VolumeFlow) CubicDecimetersPerMinute() float64 {
	if a.cubic_decimeters_per_minuteLazy != nil {
		return *a.cubic_decimeters_per_minuteLazy
	}
	cubic_decimeters_per_minute := a.convertFromBase(VolumeFlowCubicDecimeterPerMinute)
	a.cubic_decimeters_per_minuteLazy = &cubic_decimeters_per_minute
	return cubic_decimeters_per_minute
}

// OilBarrelPerDay returns the value in OilBarrelPerDay.
func (a *VolumeFlow) OilBarrelsPerDay() float64 {
	if a.oil_barrels_per_dayLazy != nil {
		return *a.oil_barrels_per_dayLazy
	}
	oil_barrels_per_day := a.convertFromBase(VolumeFlowOilBarrelPerDay)
	a.oil_barrels_per_dayLazy = &oil_barrels_per_day
	return oil_barrels_per_day
}

// OilBarrelPerMinute returns the value in OilBarrelPerMinute.
func (a *VolumeFlow) OilBarrelsPerMinute() float64 {
	if a.oil_barrels_per_minuteLazy != nil {
		return *a.oil_barrels_per_minuteLazy
	}
	oil_barrels_per_minute := a.convertFromBase(VolumeFlowOilBarrelPerMinute)
	a.oil_barrels_per_minuteLazy = &oil_barrels_per_minute
	return oil_barrels_per_minute
}

// OilBarrelPerHour returns the value in OilBarrelPerHour.
func (a *VolumeFlow) OilBarrelsPerHour() float64 {
	if a.oil_barrels_per_hourLazy != nil {
		return *a.oil_barrels_per_hourLazy
	}
	oil_barrels_per_hour := a.convertFromBase(VolumeFlowOilBarrelPerHour)
	a.oil_barrels_per_hourLazy = &oil_barrels_per_hour
	return oil_barrels_per_hour
}

// OilBarrelPerSecond returns the value in OilBarrelPerSecond.
func (a *VolumeFlow) OilBarrelsPerSecond() float64 {
	if a.oil_barrels_per_secondLazy != nil {
		return *a.oil_barrels_per_secondLazy
	}
	oil_barrels_per_second := a.convertFromBase(VolumeFlowOilBarrelPerSecond)
	a.oil_barrels_per_secondLazy = &oil_barrels_per_second
	return oil_barrels_per_second
}

// CubicMillimeterPerSecond returns the value in CubicMillimeterPerSecond.
func (a *VolumeFlow) CubicMillimetersPerSecond() float64 {
	if a.cubic_millimeters_per_secondLazy != nil {
		return *a.cubic_millimeters_per_secondLazy
	}
	cubic_millimeters_per_second := a.convertFromBase(VolumeFlowCubicMillimeterPerSecond)
	a.cubic_millimeters_per_secondLazy = &cubic_millimeters_per_second
	return cubic_millimeters_per_second
}

// AcreFootPerSecond returns the value in AcreFootPerSecond.
func (a *VolumeFlow) AcreFeetPerSecond() float64 {
	if a.acre_feet_per_secondLazy != nil {
		return *a.acre_feet_per_secondLazy
	}
	acre_feet_per_second := a.convertFromBase(VolumeFlowAcreFootPerSecond)
	a.acre_feet_per_secondLazy = &acre_feet_per_second
	return acre_feet_per_second
}

// AcreFootPerMinute returns the value in AcreFootPerMinute.
func (a *VolumeFlow) AcreFeetPerMinute() float64 {
	if a.acre_feet_per_minuteLazy != nil {
		return *a.acre_feet_per_minuteLazy
	}
	acre_feet_per_minute := a.convertFromBase(VolumeFlowAcreFootPerMinute)
	a.acre_feet_per_minuteLazy = &acre_feet_per_minute
	return acre_feet_per_minute
}

// AcreFootPerHour returns the value in AcreFootPerHour.
func (a *VolumeFlow) AcreFeetPerHour() float64 {
	if a.acre_feet_per_hourLazy != nil {
		return *a.acre_feet_per_hourLazy
	}
	acre_feet_per_hour := a.convertFromBase(VolumeFlowAcreFootPerHour)
	a.acre_feet_per_hourLazy = &acre_feet_per_hour
	return acre_feet_per_hour
}

// AcreFootPerDay returns the value in AcreFootPerDay.
func (a *VolumeFlow) AcreFeetPerDay() float64 {
	if a.acre_feet_per_dayLazy != nil {
		return *a.acre_feet_per_dayLazy
	}
	acre_feet_per_day := a.convertFromBase(VolumeFlowAcreFootPerDay)
	a.acre_feet_per_dayLazy = &acre_feet_per_day
	return acre_feet_per_day
}

// CubicCentimeterPerMinute returns the value in CubicCentimeterPerMinute.
func (a *VolumeFlow) CubicCentimetersPerMinute() float64 {
	if a.cubic_centimeters_per_minuteLazy != nil {
		return *a.cubic_centimeters_per_minuteLazy
	}
	cubic_centimeters_per_minute := a.convertFromBase(VolumeFlowCubicCentimeterPerMinute)
	a.cubic_centimeters_per_minuteLazy = &cubic_centimeters_per_minute
	return cubic_centimeters_per_minute
}

// MegausGallonPerDay returns the value in MegausGallonPerDay.
func (a *VolumeFlow) MegausGallonsPerDay() float64 {
	if a.megaus_gallons_per_dayLazy != nil {
		return *a.megaus_gallons_per_dayLazy
	}
	megaus_gallons_per_day := a.convertFromBase(VolumeFlowMegausGallonPerDay)
	a.megaus_gallons_per_dayLazy = &megaus_gallons_per_day
	return megaus_gallons_per_day
}

// NanoliterPerSecond returns the value in NanoliterPerSecond.
func (a *VolumeFlow) NanolitersPerSecond() float64 {
	if a.nanoliters_per_secondLazy != nil {
		return *a.nanoliters_per_secondLazy
	}
	nanoliters_per_second := a.convertFromBase(VolumeFlowNanoliterPerSecond)
	a.nanoliters_per_secondLazy = &nanoliters_per_second
	return nanoliters_per_second
}

// MicroliterPerSecond returns the value in MicroliterPerSecond.
func (a *VolumeFlow) MicrolitersPerSecond() float64 {
	if a.microliters_per_secondLazy != nil {
		return *a.microliters_per_secondLazy
	}
	microliters_per_second := a.convertFromBase(VolumeFlowMicroliterPerSecond)
	a.microliters_per_secondLazy = &microliters_per_second
	return microliters_per_second
}

// MilliliterPerSecond returns the value in MilliliterPerSecond.
func (a *VolumeFlow) MillilitersPerSecond() float64 {
	if a.milliliters_per_secondLazy != nil {
		return *a.milliliters_per_secondLazy
	}
	milliliters_per_second := a.convertFromBase(VolumeFlowMilliliterPerSecond)
	a.milliliters_per_secondLazy = &milliliters_per_second
	return milliliters_per_second
}

// CentiliterPerSecond returns the value in CentiliterPerSecond.
func (a *VolumeFlow) CentilitersPerSecond() float64 {
	if a.centiliters_per_secondLazy != nil {
		return *a.centiliters_per_secondLazy
	}
	centiliters_per_second := a.convertFromBase(VolumeFlowCentiliterPerSecond)
	a.centiliters_per_secondLazy = &centiliters_per_second
	return centiliters_per_second
}

// DeciliterPerSecond returns the value in DeciliterPerSecond.
func (a *VolumeFlow) DecilitersPerSecond() float64 {
	if a.deciliters_per_secondLazy != nil {
		return *a.deciliters_per_secondLazy
	}
	deciliters_per_second := a.convertFromBase(VolumeFlowDeciliterPerSecond)
	a.deciliters_per_secondLazy = &deciliters_per_second
	return deciliters_per_second
}

// DecaliterPerSecond returns the value in DecaliterPerSecond.
func (a *VolumeFlow) DecalitersPerSecond() float64 {
	if a.decaliters_per_secondLazy != nil {
		return *a.decaliters_per_secondLazy
	}
	decaliters_per_second := a.convertFromBase(VolumeFlowDecaliterPerSecond)
	a.decaliters_per_secondLazy = &decaliters_per_second
	return decaliters_per_second
}

// HectoliterPerSecond returns the value in HectoliterPerSecond.
func (a *VolumeFlow) HectolitersPerSecond() float64 {
	if a.hectoliters_per_secondLazy != nil {
		return *a.hectoliters_per_secondLazy
	}
	hectoliters_per_second := a.convertFromBase(VolumeFlowHectoliterPerSecond)
	a.hectoliters_per_secondLazy = &hectoliters_per_second
	return hectoliters_per_second
}

// KiloliterPerSecond returns the value in KiloliterPerSecond.
func (a *VolumeFlow) KilolitersPerSecond() float64 {
	if a.kiloliters_per_secondLazy != nil {
		return *a.kiloliters_per_secondLazy
	}
	kiloliters_per_second := a.convertFromBase(VolumeFlowKiloliterPerSecond)
	a.kiloliters_per_secondLazy = &kiloliters_per_second
	return kiloliters_per_second
}

// MegaliterPerSecond returns the value in MegaliterPerSecond.
func (a *VolumeFlow) MegalitersPerSecond() float64 {
	if a.megaliters_per_secondLazy != nil {
		return *a.megaliters_per_secondLazy
	}
	megaliters_per_second := a.convertFromBase(VolumeFlowMegaliterPerSecond)
	a.megaliters_per_secondLazy = &megaliters_per_second
	return megaliters_per_second
}

// NanoliterPerMinute returns the value in NanoliterPerMinute.
func (a *VolumeFlow) NanolitersPerMinute() float64 {
	if a.nanoliters_per_minuteLazy != nil {
		return *a.nanoliters_per_minuteLazy
	}
	nanoliters_per_minute := a.convertFromBase(VolumeFlowNanoliterPerMinute)
	a.nanoliters_per_minuteLazy = &nanoliters_per_minute
	return nanoliters_per_minute
}

// MicroliterPerMinute returns the value in MicroliterPerMinute.
func (a *VolumeFlow) MicrolitersPerMinute() float64 {
	if a.microliters_per_minuteLazy != nil {
		return *a.microliters_per_minuteLazy
	}
	microliters_per_minute := a.convertFromBase(VolumeFlowMicroliterPerMinute)
	a.microliters_per_minuteLazy = &microliters_per_minute
	return microliters_per_minute
}

// MilliliterPerMinute returns the value in MilliliterPerMinute.
func (a *VolumeFlow) MillilitersPerMinute() float64 {
	if a.milliliters_per_minuteLazy != nil {
		return *a.milliliters_per_minuteLazy
	}
	milliliters_per_minute := a.convertFromBase(VolumeFlowMilliliterPerMinute)
	a.milliliters_per_minuteLazy = &milliliters_per_minute
	return milliliters_per_minute
}

// CentiliterPerMinute returns the value in CentiliterPerMinute.
func (a *VolumeFlow) CentilitersPerMinute() float64 {
	if a.centiliters_per_minuteLazy != nil {
		return *a.centiliters_per_minuteLazy
	}
	centiliters_per_minute := a.convertFromBase(VolumeFlowCentiliterPerMinute)
	a.centiliters_per_minuteLazy = &centiliters_per_minute
	return centiliters_per_minute
}

// DeciliterPerMinute returns the value in DeciliterPerMinute.
func (a *VolumeFlow) DecilitersPerMinute() float64 {
	if a.deciliters_per_minuteLazy != nil {
		return *a.deciliters_per_minuteLazy
	}
	deciliters_per_minute := a.convertFromBase(VolumeFlowDeciliterPerMinute)
	a.deciliters_per_minuteLazy = &deciliters_per_minute
	return deciliters_per_minute
}

// DecaliterPerMinute returns the value in DecaliterPerMinute.
func (a *VolumeFlow) DecalitersPerMinute() float64 {
	if a.decaliters_per_minuteLazy != nil {
		return *a.decaliters_per_minuteLazy
	}
	decaliters_per_minute := a.convertFromBase(VolumeFlowDecaliterPerMinute)
	a.decaliters_per_minuteLazy = &decaliters_per_minute
	return decaliters_per_minute
}

// HectoliterPerMinute returns the value in HectoliterPerMinute.
func (a *VolumeFlow) HectolitersPerMinute() float64 {
	if a.hectoliters_per_minuteLazy != nil {
		return *a.hectoliters_per_minuteLazy
	}
	hectoliters_per_minute := a.convertFromBase(VolumeFlowHectoliterPerMinute)
	a.hectoliters_per_minuteLazy = &hectoliters_per_minute
	return hectoliters_per_minute
}

// KiloliterPerMinute returns the value in KiloliterPerMinute.
func (a *VolumeFlow) KilolitersPerMinute() float64 {
	if a.kiloliters_per_minuteLazy != nil {
		return *a.kiloliters_per_minuteLazy
	}
	kiloliters_per_minute := a.convertFromBase(VolumeFlowKiloliterPerMinute)
	a.kiloliters_per_minuteLazy = &kiloliters_per_minute
	return kiloliters_per_minute
}

// MegaliterPerMinute returns the value in MegaliterPerMinute.
func (a *VolumeFlow) MegalitersPerMinute() float64 {
	if a.megaliters_per_minuteLazy != nil {
		return *a.megaliters_per_minuteLazy
	}
	megaliters_per_minute := a.convertFromBase(VolumeFlowMegaliterPerMinute)
	a.megaliters_per_minuteLazy = &megaliters_per_minute
	return megaliters_per_minute
}

// NanoliterPerHour returns the value in NanoliterPerHour.
func (a *VolumeFlow) NanolitersPerHour() float64 {
	if a.nanoliters_per_hourLazy != nil {
		return *a.nanoliters_per_hourLazy
	}
	nanoliters_per_hour := a.convertFromBase(VolumeFlowNanoliterPerHour)
	a.nanoliters_per_hourLazy = &nanoliters_per_hour
	return nanoliters_per_hour
}

// MicroliterPerHour returns the value in MicroliterPerHour.
func (a *VolumeFlow) MicrolitersPerHour() float64 {
	if a.microliters_per_hourLazy != nil {
		return *a.microliters_per_hourLazy
	}
	microliters_per_hour := a.convertFromBase(VolumeFlowMicroliterPerHour)
	a.microliters_per_hourLazy = &microliters_per_hour
	return microliters_per_hour
}

// MilliliterPerHour returns the value in MilliliterPerHour.
func (a *VolumeFlow) MillilitersPerHour() float64 {
	if a.milliliters_per_hourLazy != nil {
		return *a.milliliters_per_hourLazy
	}
	milliliters_per_hour := a.convertFromBase(VolumeFlowMilliliterPerHour)
	a.milliliters_per_hourLazy = &milliliters_per_hour
	return milliliters_per_hour
}

// CentiliterPerHour returns the value in CentiliterPerHour.
func (a *VolumeFlow) CentilitersPerHour() float64 {
	if a.centiliters_per_hourLazy != nil {
		return *a.centiliters_per_hourLazy
	}
	centiliters_per_hour := a.convertFromBase(VolumeFlowCentiliterPerHour)
	a.centiliters_per_hourLazy = &centiliters_per_hour
	return centiliters_per_hour
}

// DeciliterPerHour returns the value in DeciliterPerHour.
func (a *VolumeFlow) DecilitersPerHour() float64 {
	if a.deciliters_per_hourLazy != nil {
		return *a.deciliters_per_hourLazy
	}
	deciliters_per_hour := a.convertFromBase(VolumeFlowDeciliterPerHour)
	a.deciliters_per_hourLazy = &deciliters_per_hour
	return deciliters_per_hour
}

// DecaliterPerHour returns the value in DecaliterPerHour.
func (a *VolumeFlow) DecalitersPerHour() float64 {
	if a.decaliters_per_hourLazy != nil {
		return *a.decaliters_per_hourLazy
	}
	decaliters_per_hour := a.convertFromBase(VolumeFlowDecaliterPerHour)
	a.decaliters_per_hourLazy = &decaliters_per_hour
	return decaliters_per_hour
}

// HectoliterPerHour returns the value in HectoliterPerHour.
func (a *VolumeFlow) HectolitersPerHour() float64 {
	if a.hectoliters_per_hourLazy != nil {
		return *a.hectoliters_per_hourLazy
	}
	hectoliters_per_hour := a.convertFromBase(VolumeFlowHectoliterPerHour)
	a.hectoliters_per_hourLazy = &hectoliters_per_hour
	return hectoliters_per_hour
}

// KiloliterPerHour returns the value in KiloliterPerHour.
func (a *VolumeFlow) KilolitersPerHour() float64 {
	if a.kiloliters_per_hourLazy != nil {
		return *a.kiloliters_per_hourLazy
	}
	kiloliters_per_hour := a.convertFromBase(VolumeFlowKiloliterPerHour)
	a.kiloliters_per_hourLazy = &kiloliters_per_hour
	return kiloliters_per_hour
}

// MegaliterPerHour returns the value in MegaliterPerHour.
func (a *VolumeFlow) MegalitersPerHour() float64 {
	if a.megaliters_per_hourLazy != nil {
		return *a.megaliters_per_hourLazy
	}
	megaliters_per_hour := a.convertFromBase(VolumeFlowMegaliterPerHour)
	a.megaliters_per_hourLazy = &megaliters_per_hour
	return megaliters_per_hour
}

// NanoliterPerDay returns the value in NanoliterPerDay.
func (a *VolumeFlow) NanolitersPerDay() float64 {
	if a.nanoliters_per_dayLazy != nil {
		return *a.nanoliters_per_dayLazy
	}
	nanoliters_per_day := a.convertFromBase(VolumeFlowNanoliterPerDay)
	a.nanoliters_per_dayLazy = &nanoliters_per_day
	return nanoliters_per_day
}

// MicroliterPerDay returns the value in MicroliterPerDay.
func (a *VolumeFlow) MicrolitersPerDay() float64 {
	if a.microliters_per_dayLazy != nil {
		return *a.microliters_per_dayLazy
	}
	microliters_per_day := a.convertFromBase(VolumeFlowMicroliterPerDay)
	a.microliters_per_dayLazy = &microliters_per_day
	return microliters_per_day
}

// MilliliterPerDay returns the value in MilliliterPerDay.
func (a *VolumeFlow) MillilitersPerDay() float64 {
	if a.milliliters_per_dayLazy != nil {
		return *a.milliliters_per_dayLazy
	}
	milliliters_per_day := a.convertFromBase(VolumeFlowMilliliterPerDay)
	a.milliliters_per_dayLazy = &milliliters_per_day
	return milliliters_per_day
}

// CentiliterPerDay returns the value in CentiliterPerDay.
func (a *VolumeFlow) CentilitersPerDay() float64 {
	if a.centiliters_per_dayLazy != nil {
		return *a.centiliters_per_dayLazy
	}
	centiliters_per_day := a.convertFromBase(VolumeFlowCentiliterPerDay)
	a.centiliters_per_dayLazy = &centiliters_per_day
	return centiliters_per_day
}

// DeciliterPerDay returns the value in DeciliterPerDay.
func (a *VolumeFlow) DecilitersPerDay() float64 {
	if a.deciliters_per_dayLazy != nil {
		return *a.deciliters_per_dayLazy
	}
	deciliters_per_day := a.convertFromBase(VolumeFlowDeciliterPerDay)
	a.deciliters_per_dayLazy = &deciliters_per_day
	return deciliters_per_day
}

// DecaliterPerDay returns the value in DecaliterPerDay.
func (a *VolumeFlow) DecalitersPerDay() float64 {
	if a.decaliters_per_dayLazy != nil {
		return *a.decaliters_per_dayLazy
	}
	decaliters_per_day := a.convertFromBase(VolumeFlowDecaliterPerDay)
	a.decaliters_per_dayLazy = &decaliters_per_day
	return decaliters_per_day
}

// HectoliterPerDay returns the value in HectoliterPerDay.
func (a *VolumeFlow) HectolitersPerDay() float64 {
	if a.hectoliters_per_dayLazy != nil {
		return *a.hectoliters_per_dayLazy
	}
	hectoliters_per_day := a.convertFromBase(VolumeFlowHectoliterPerDay)
	a.hectoliters_per_dayLazy = &hectoliters_per_day
	return hectoliters_per_day
}

// KiloliterPerDay returns the value in KiloliterPerDay.
func (a *VolumeFlow) KilolitersPerDay() float64 {
	if a.kiloliters_per_dayLazy != nil {
		return *a.kiloliters_per_dayLazy
	}
	kiloliters_per_day := a.convertFromBase(VolumeFlowKiloliterPerDay)
	a.kiloliters_per_dayLazy = &kiloliters_per_day
	return kiloliters_per_day
}

// MegaliterPerDay returns the value in MegaliterPerDay.
func (a *VolumeFlow) MegalitersPerDay() float64 {
	if a.megaliters_per_dayLazy != nil {
		return *a.megaliters_per_dayLazy
	}
	megaliters_per_day := a.convertFromBase(VolumeFlowMegaliterPerDay)
	a.megaliters_per_dayLazy = &megaliters_per_day
	return megaliters_per_day
}

// MegaukGallonPerDay returns the value in MegaukGallonPerDay.
func (a *VolumeFlow) MegaukGallonsPerDay() float64 {
	if a.megauk_gallons_per_dayLazy != nil {
		return *a.megauk_gallons_per_dayLazy
	}
	megauk_gallons_per_day := a.convertFromBase(VolumeFlowMegaukGallonPerDay)
	a.megauk_gallons_per_dayLazy = &megauk_gallons_per_day
	return megauk_gallons_per_day
}

// MegaukGallonPerSecond returns the value in MegaukGallonPerSecond.
func (a *VolumeFlow) MegaukGallonsPerSecond() float64 {
	if a.megauk_gallons_per_secondLazy != nil {
		return *a.megauk_gallons_per_secondLazy
	}
	megauk_gallons_per_second := a.convertFromBase(VolumeFlowMegaukGallonPerSecond)
	a.megauk_gallons_per_secondLazy = &megauk_gallons_per_second
	return megauk_gallons_per_second
}


// ToDto creates an VolumeFlowDto representation.
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

// ToDtoJSON creates an VolumeFlowDto representation.
func (a *VolumeFlow) ToDtoJSON(holdInUnit *VolumeFlowUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VolumeFlow to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a VolumeFlow) String() string {
	return a.ToString(VolumeFlowCubicMeterPerSecond, 2)
}

// ToString formats the VolumeFlow to string.
// fractionalDigits -1 for not mention
func (a *VolumeFlow) ToString(unit VolumeFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VolumeFlow) getUnitAbbreviation(unit VolumeFlowUnits) string {
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

// Check if the given VolumeFlow are equals to the current VolumeFlow
func (a *VolumeFlow) Equals(other *VolumeFlow) bool {
	return a.value == other.BaseValue()
}

// Check if the given VolumeFlow are equals to the current VolumeFlow
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

// Add the given VolumeFlow to the current VolumeFlow.
func (a *VolumeFlow) Add(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value + other.BaseValue()}
}

// Subtract the given VolumeFlow to the current VolumeFlow.
func (a *VolumeFlow) Subtract(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value - other.BaseValue()}
}

// Multiply the given VolumeFlow to the current VolumeFlow.
func (a *VolumeFlow) Multiply(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value * other.BaseValue()}
}

// Divide the given VolumeFlow to the current VolumeFlow.
func (a *VolumeFlow) Divide(other *VolumeFlow) *VolumeFlow {
	return &VolumeFlow{value: a.value / other.BaseValue()}
}