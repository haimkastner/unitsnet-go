// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeUnits defines various units of Volume.
type VolumeUnits string

const (
    
        // 
        VolumeLiter VolumeUnits = "Liter"
        // 
        VolumeCubicMeter VolumeUnits = "CubicMeter"
        // 
        VolumeCubicKilometer VolumeUnits = "CubicKilometer"
        // 
        VolumeCubicHectometer VolumeUnits = "CubicHectometer"
        // 
        VolumeCubicDecimeter VolumeUnits = "CubicDecimeter"
        // 
        VolumeCubicCentimeter VolumeUnits = "CubicCentimeter"
        // 
        VolumeCubicMillimeter VolumeUnits = "CubicMillimeter"
        // 
        VolumeCubicMicrometer VolumeUnits = "CubicMicrometer"
        // 
        VolumeCubicMile VolumeUnits = "CubicMile"
        // 
        VolumeCubicYard VolumeUnits = "CubicYard"
        // 
        VolumeCubicFoot VolumeUnits = "CubicFoot"
        // 
        VolumeCubicInch VolumeUnits = "CubicInch"
        // The British imperial gallon (frequently called simply "gallon") is defined as exactly 4.54609 litres.
        VolumeImperialGallon VolumeUnits = "ImperialGallon"
        // 
        VolumeImperialOunce VolumeUnits = "ImperialOunce"
        // The US liquid gallon (frequently called simply "gallon") is legally defined as 231 cubic inches, which is exactly 3.785411784 litres.
        VolumeUsGallon VolumeUnits = "UsGallon"
        // 
        VolumeUsOunce VolumeUnits = "UsOunce"
        // 
        VolumeUsTablespoon VolumeUnits = "UsTablespoon"
        // 
        VolumeAuTablespoon VolumeUnits = "AuTablespoon"
        // 
        VolumeUkTablespoon VolumeUnits = "UkTablespoon"
        // 
        VolumeMetricTeaspoon VolumeUnits = "MetricTeaspoon"
        // 
        VolumeUsTeaspoon VolumeUnits = "UsTeaspoon"
        // 
        VolumeMetricCup VolumeUnits = "MetricCup"
        // 
        VolumeUsCustomaryCup VolumeUnits = "UsCustomaryCup"
        // 
        VolumeUsLegalCup VolumeUnits = "UsLegalCup"
        // 
        VolumeOilBarrel VolumeUnits = "OilBarrel"
        // 
        VolumeUsBeerBarrel VolumeUnits = "UsBeerBarrel"
        // 
        VolumeImperialBeerBarrel VolumeUnits = "ImperialBeerBarrel"
        // 
        VolumeUsQuart VolumeUnits = "UsQuart"
        // 
        VolumeImperialQuart VolumeUnits = "ImperialQuart"
        // 
        VolumeUsPint VolumeUnits = "UsPint"
        // 
        VolumeAcreFoot VolumeUnits = "AcreFoot"
        // 
        VolumeImperialPint VolumeUnits = "ImperialPint"
        // 
        VolumeBoardFoot VolumeUnits = "BoardFoot"
        // 
        VolumeNanoliter VolumeUnits = "Nanoliter"
        // 
        VolumeMicroliter VolumeUnits = "Microliter"
        // 
        VolumeMilliliter VolumeUnits = "Milliliter"
        // 
        VolumeCentiliter VolumeUnits = "Centiliter"
        // 
        VolumeDeciliter VolumeUnits = "Deciliter"
        // 
        VolumeDecaliter VolumeUnits = "Decaliter"
        // 
        VolumeHectoliter VolumeUnits = "Hectoliter"
        // 
        VolumeKiloliter VolumeUnits = "Kiloliter"
        // 
        VolumeMegaliter VolumeUnits = "Megaliter"
        // 
        VolumeHectocubicMeter VolumeUnits = "HectocubicMeter"
        // 
        VolumeKilocubicMeter VolumeUnits = "KilocubicMeter"
        // 
        VolumeHectocubicFoot VolumeUnits = "HectocubicFoot"
        // 
        VolumeKilocubicFoot VolumeUnits = "KilocubicFoot"
        // 
        VolumeMegacubicFoot VolumeUnits = "MegacubicFoot"
        // 
        VolumeKiloimperialGallon VolumeUnits = "KiloimperialGallon"
        // 
        VolumeMegaimperialGallon VolumeUnits = "MegaimperialGallon"
        // 
        VolumeDecausGallon VolumeUnits = "DecausGallon"
        // 
        VolumeDeciusGallon VolumeUnits = "DeciusGallon"
        // 
        VolumeHectousGallon VolumeUnits = "HectousGallon"
        // 
        VolumeKilousGallon VolumeUnits = "KilousGallon"
        // 
        VolumeMegausGallon VolumeUnits = "MegausGallon"
)

var internalVolumeUnitsMap = map[VolumeUnits]bool{
	
	VolumeLiter: true,
	VolumeCubicMeter: true,
	VolumeCubicKilometer: true,
	VolumeCubicHectometer: true,
	VolumeCubicDecimeter: true,
	VolumeCubicCentimeter: true,
	VolumeCubicMillimeter: true,
	VolumeCubicMicrometer: true,
	VolumeCubicMile: true,
	VolumeCubicYard: true,
	VolumeCubicFoot: true,
	VolumeCubicInch: true,
	VolumeImperialGallon: true,
	VolumeImperialOunce: true,
	VolumeUsGallon: true,
	VolumeUsOunce: true,
	VolumeUsTablespoon: true,
	VolumeAuTablespoon: true,
	VolumeUkTablespoon: true,
	VolumeMetricTeaspoon: true,
	VolumeUsTeaspoon: true,
	VolumeMetricCup: true,
	VolumeUsCustomaryCup: true,
	VolumeUsLegalCup: true,
	VolumeOilBarrel: true,
	VolumeUsBeerBarrel: true,
	VolumeImperialBeerBarrel: true,
	VolumeUsQuart: true,
	VolumeImperialQuart: true,
	VolumeUsPint: true,
	VolumeAcreFoot: true,
	VolumeImperialPint: true,
	VolumeBoardFoot: true,
	VolumeNanoliter: true,
	VolumeMicroliter: true,
	VolumeMilliliter: true,
	VolumeCentiliter: true,
	VolumeDeciliter: true,
	VolumeDecaliter: true,
	VolumeHectoliter: true,
	VolumeKiloliter: true,
	VolumeMegaliter: true,
	VolumeHectocubicMeter: true,
	VolumeKilocubicMeter: true,
	VolumeHectocubicFoot: true,
	VolumeKilocubicFoot: true,
	VolumeMegacubicFoot: true,
	VolumeKiloimperialGallon: true,
	VolumeMegaimperialGallon: true,
	VolumeDecausGallon: true,
	VolumeDeciusGallon: true,
	VolumeHectousGallon: true,
	VolumeKilousGallon: true,
	VolumeMegausGallon: true,
}

// VolumeDto represents a Volume measurement with a numerical value and its corresponding unit.
type VolumeDto struct {
    // Value is the numerical representation of the Volume.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Volume, as defined in the VolumeUnits enumeration.
	Unit  VolumeUnits `json:"unit" validate:"required,oneof=Liter,CubicMeter,CubicKilometer,CubicHectometer,CubicDecimeter,CubicCentimeter,CubicMillimeter,CubicMicrometer,CubicMile,CubicYard,CubicFoot,CubicInch,ImperialGallon,ImperialOunce,UsGallon,UsOunce,UsTablespoon,AuTablespoon,UkTablespoon,MetricTeaspoon,UsTeaspoon,MetricCup,UsCustomaryCup,UsLegalCup,OilBarrel,UsBeerBarrel,ImperialBeerBarrel,UsQuart,ImperialQuart,UsPint,AcreFoot,ImperialPint,BoardFoot,Nanoliter,Microliter,Milliliter,Centiliter,Deciliter,Decaliter,Hectoliter,Kiloliter,Megaliter,HectocubicMeter,KilocubicMeter,HectocubicFoot,KilocubicFoot,MegacubicFoot,KiloimperialGallon,MegaimperialGallon,DecausGallon,DeciusGallon,HectousGallon,KilousGallon,MegausGallon"`
}

// VolumeDtoFactory groups methods for creating and serializing VolumeDto objects.
type VolumeDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumeDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumeDtoFactory) FromJSON(data []byte) (*VolumeDto, error) {
	a := VolumeDto{}

    // Parse JSON into VolumeDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a VolumeDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumeDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Volume represents a measurement in a of Volume.
//
// Volume is the quantity of three-dimensional space enclosed by some closed boundary, for example, the space that a substance (solid, liquid, gas, or plasma) or shape occupies or contains.[1] Volume is often quantified numerically using the SI derived unit, the cubic metre. The volume of a container is generally understood to be the capacity of the container, i. e. the amount of fluid (gas or liquid) that the container could hold, rather than the amount of space the container itself displaces.
type Volume struct {
	// value is the base measurement stored internally.
	value       float64
    
    litersLazy *float64 
    cubic_metersLazy *float64 
    cubic_kilometersLazy *float64 
    cubic_hectometersLazy *float64 
    cubic_decimetersLazy *float64 
    cubic_centimetersLazy *float64 
    cubic_millimetersLazy *float64 
    cubic_micrometersLazy *float64 
    cubic_milesLazy *float64 
    cubic_yardsLazy *float64 
    cubic_feetLazy *float64 
    cubic_inchesLazy *float64 
    imperial_gallonsLazy *float64 
    imperial_ouncesLazy *float64 
    us_gallonsLazy *float64 
    us_ouncesLazy *float64 
    us_tablespoonsLazy *float64 
    au_tablespoonsLazy *float64 
    uk_tablespoonsLazy *float64 
    metric_teaspoonsLazy *float64 
    us_teaspoonsLazy *float64 
    metric_cupsLazy *float64 
    us_customary_cupsLazy *float64 
    us_legal_cupsLazy *float64 
    oil_barrelsLazy *float64 
    us_beer_barrelsLazy *float64 
    imperial_beer_barrelsLazy *float64 
    us_quartsLazy *float64 
    imperial_quartsLazy *float64 
    us_pintsLazy *float64 
    acre_feetLazy *float64 
    imperial_pintsLazy *float64 
    board_feetLazy *float64 
    nanolitersLazy *float64 
    microlitersLazy *float64 
    millilitersLazy *float64 
    centilitersLazy *float64 
    decilitersLazy *float64 
    decalitersLazy *float64 
    hectolitersLazy *float64 
    kilolitersLazy *float64 
    megalitersLazy *float64 
    hectocubic_metersLazy *float64 
    kilocubic_metersLazy *float64 
    hectocubic_feetLazy *float64 
    kilocubic_feetLazy *float64 
    megacubic_feetLazy *float64 
    kiloimperial_gallonsLazy *float64 
    megaimperial_gallonsLazy *float64 
    decaus_gallonsLazy *float64 
    decius_gallonsLazy *float64 
    hectous_gallonsLazy *float64 
    kilous_gallonsLazy *float64 
    megaus_gallonsLazy *float64 
}

// VolumeFactory groups methods for creating Volume instances.
type VolumeFactory struct{}

// CreateVolume creates a new Volume instance from the given value and unit.
func (uf VolumeFactory) CreateVolume(value float64, unit VolumeUnits) (*Volume, error) {
	return newVolume(value, unit)
}

// FromDto converts a VolumeDto to a Volume instance.
func (uf VolumeFactory) FromDto(dto VolumeDto) (*Volume, error) {
	return newVolume(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Volume instance.
func (uf VolumeFactory) FromDtoJSON(data []byte) (*Volume, error) {
	unitDto, err := VolumeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumeDto from JSON: %w", err)
	}
	return VolumeFactory{}.FromDto(*unitDto)
}


// FromLiters creates a new Volume instance from a value in Liters.
func (uf VolumeFactory) FromLiters(value float64) (*Volume, error) {
	return newVolume(value, VolumeLiter)
}

// FromCubicMeters creates a new Volume instance from a value in CubicMeters.
func (uf VolumeFactory) FromCubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMeter)
}

// FromCubicKilometers creates a new Volume instance from a value in CubicKilometers.
func (uf VolumeFactory) FromCubicKilometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicKilometer)
}

// FromCubicHectometers creates a new Volume instance from a value in CubicHectometers.
func (uf VolumeFactory) FromCubicHectometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicHectometer)
}

// FromCubicDecimeters creates a new Volume instance from a value in CubicDecimeters.
func (uf VolumeFactory) FromCubicDecimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicDecimeter)
}

// FromCubicCentimeters creates a new Volume instance from a value in CubicCentimeters.
func (uf VolumeFactory) FromCubicCentimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicCentimeter)
}

// FromCubicMillimeters creates a new Volume instance from a value in CubicMillimeters.
func (uf VolumeFactory) FromCubicMillimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMillimeter)
}

// FromCubicMicrometers creates a new Volume instance from a value in CubicMicrometers.
func (uf VolumeFactory) FromCubicMicrometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMicrometer)
}

// FromCubicMiles creates a new Volume instance from a value in CubicMiles.
func (uf VolumeFactory) FromCubicMiles(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMile)
}

// FromCubicYards creates a new Volume instance from a value in CubicYards.
func (uf VolumeFactory) FromCubicYards(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicYard)
}

// FromCubicFeet creates a new Volume instance from a value in CubicFeet.
func (uf VolumeFactory) FromCubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicFoot)
}

// FromCubicInches creates a new Volume instance from a value in CubicInches.
func (uf VolumeFactory) FromCubicInches(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicInch)
}

// FromImperialGallons creates a new Volume instance from a value in ImperialGallons.
func (uf VolumeFactory) FromImperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialGallon)
}

// FromImperialOunces creates a new Volume instance from a value in ImperialOunces.
func (uf VolumeFactory) FromImperialOunces(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialOunce)
}

// FromUsGallons creates a new Volume instance from a value in UsGallons.
func (uf VolumeFactory) FromUsGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsGallon)
}

// FromUsOunces creates a new Volume instance from a value in UsOunces.
func (uf VolumeFactory) FromUsOunces(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsOunce)
}

// FromUsTablespoons creates a new Volume instance from a value in UsTablespoons.
func (uf VolumeFactory) FromUsTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsTablespoon)
}

// FromAuTablespoons creates a new Volume instance from a value in AuTablespoons.
func (uf VolumeFactory) FromAuTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeAuTablespoon)
}

// FromUkTablespoons creates a new Volume instance from a value in UkTablespoons.
func (uf VolumeFactory) FromUkTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUkTablespoon)
}

// FromMetricTeaspoons creates a new Volume instance from a value in MetricTeaspoons.
func (uf VolumeFactory) FromMetricTeaspoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMetricTeaspoon)
}

// FromUsTeaspoons creates a new Volume instance from a value in UsTeaspoons.
func (uf VolumeFactory) FromUsTeaspoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsTeaspoon)
}

// FromMetricCups creates a new Volume instance from a value in MetricCups.
func (uf VolumeFactory) FromMetricCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeMetricCup)
}

// FromUsCustomaryCups creates a new Volume instance from a value in UsCustomaryCups.
func (uf VolumeFactory) FromUsCustomaryCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsCustomaryCup)
}

// FromUsLegalCups creates a new Volume instance from a value in UsLegalCups.
func (uf VolumeFactory) FromUsLegalCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsLegalCup)
}

// FromOilBarrels creates a new Volume instance from a value in OilBarrels.
func (uf VolumeFactory) FromOilBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeOilBarrel)
}

// FromUsBeerBarrels creates a new Volume instance from a value in UsBeerBarrels.
func (uf VolumeFactory) FromUsBeerBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsBeerBarrel)
}

// FromImperialBeerBarrels creates a new Volume instance from a value in ImperialBeerBarrels.
func (uf VolumeFactory) FromImperialBeerBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialBeerBarrel)
}

// FromUsQuarts creates a new Volume instance from a value in UsQuarts.
func (uf VolumeFactory) FromUsQuarts(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsQuart)
}

// FromImperialQuarts creates a new Volume instance from a value in ImperialQuarts.
func (uf VolumeFactory) FromImperialQuarts(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialQuart)
}

// FromUsPints creates a new Volume instance from a value in UsPints.
func (uf VolumeFactory) FromUsPints(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsPint)
}

// FromAcreFeet creates a new Volume instance from a value in AcreFeet.
func (uf VolumeFactory) FromAcreFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeAcreFoot)
}

// FromImperialPints creates a new Volume instance from a value in ImperialPints.
func (uf VolumeFactory) FromImperialPints(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialPint)
}

// FromBoardFeet creates a new Volume instance from a value in BoardFeet.
func (uf VolumeFactory) FromBoardFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeBoardFoot)
}

// FromNanoliters creates a new Volume instance from a value in Nanoliters.
func (uf VolumeFactory) FromNanoliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeNanoliter)
}

// FromMicroliters creates a new Volume instance from a value in Microliters.
func (uf VolumeFactory) FromMicroliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMicroliter)
}

// FromMilliliters creates a new Volume instance from a value in Milliliters.
func (uf VolumeFactory) FromMilliliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMilliliter)
}

// FromCentiliters creates a new Volume instance from a value in Centiliters.
func (uf VolumeFactory) FromCentiliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCentiliter)
}

// FromDeciliters creates a new Volume instance from a value in Deciliters.
func (uf VolumeFactory) FromDeciliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeDeciliter)
}

// FromDecaliters creates a new Volume instance from a value in Decaliters.
func (uf VolumeFactory) FromDecaliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeDecaliter)
}

// FromHectoliters creates a new Volume instance from a value in Hectoliters.
func (uf VolumeFactory) FromHectoliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectoliter)
}

// FromKiloliters creates a new Volume instance from a value in Kiloliters.
func (uf VolumeFactory) FromKiloliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeKiloliter)
}

// FromMegaliters creates a new Volume instance from a value in Megaliters.
func (uf VolumeFactory) FromMegaliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegaliter)
}

// FromHectocubicMeters creates a new Volume instance from a value in HectocubicMeters.
func (uf VolumeFactory) FromHectocubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectocubicMeter)
}

// FromKilocubicMeters creates a new Volume instance from a value in KilocubicMeters.
func (uf VolumeFactory) FromKilocubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilocubicMeter)
}

// FromHectocubicFeet creates a new Volume instance from a value in HectocubicFeet.
func (uf VolumeFactory) FromHectocubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectocubicFoot)
}

// FromKilocubicFeet creates a new Volume instance from a value in KilocubicFeet.
func (uf VolumeFactory) FromKilocubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilocubicFoot)
}

// FromMegacubicFeet creates a new Volume instance from a value in MegacubicFeet.
func (uf VolumeFactory) FromMegacubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegacubicFoot)
}

// FromKiloimperialGallons creates a new Volume instance from a value in KiloimperialGallons.
func (uf VolumeFactory) FromKiloimperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeKiloimperialGallon)
}

// FromMegaimperialGallons creates a new Volume instance from a value in MegaimperialGallons.
func (uf VolumeFactory) FromMegaimperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegaimperialGallon)
}

// FromDecausGallons creates a new Volume instance from a value in DecausGallons.
func (uf VolumeFactory) FromDecausGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeDecausGallon)
}

// FromDeciusGallons creates a new Volume instance from a value in DeciusGallons.
func (uf VolumeFactory) FromDeciusGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeDeciusGallon)
}

// FromHectousGallons creates a new Volume instance from a value in HectousGallons.
func (uf VolumeFactory) FromHectousGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectousGallon)
}

// FromKilousGallons creates a new Volume instance from a value in KilousGallons.
func (uf VolumeFactory) FromKilousGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilousGallon)
}

// FromMegausGallons creates a new Volume instance from a value in MegausGallons.
func (uf VolumeFactory) FromMegausGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegausGallon)
}


// newVolume creates a new Volume.
func newVolume(value float64, fromUnit VolumeUnits) (*Volume, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalVolumeUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in VolumeUnits", fromUnit)
	}
	a := &Volume{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Volume in CubicMeter unit (the base/default quantity).
func (a *Volume) BaseValue() float64 {
	return a.value
}


// Liters returns the Volume value in Liters.
//
// 
func (a *Volume) Liters() float64 {
	if a.litersLazy != nil {
		return *a.litersLazy
	}
	liters := a.convertFromBase(VolumeLiter)
	a.litersLazy = &liters
	return liters
}

// CubicMeters returns the Volume value in CubicMeters.
//
// 
func (a *Volume) CubicMeters() float64 {
	if a.cubic_metersLazy != nil {
		return *a.cubic_metersLazy
	}
	cubic_meters := a.convertFromBase(VolumeCubicMeter)
	a.cubic_metersLazy = &cubic_meters
	return cubic_meters
}

// CubicKilometers returns the Volume value in CubicKilometers.
//
// 
func (a *Volume) CubicKilometers() float64 {
	if a.cubic_kilometersLazy != nil {
		return *a.cubic_kilometersLazy
	}
	cubic_kilometers := a.convertFromBase(VolumeCubicKilometer)
	a.cubic_kilometersLazy = &cubic_kilometers
	return cubic_kilometers
}

// CubicHectometers returns the Volume value in CubicHectometers.
//
// 
func (a *Volume) CubicHectometers() float64 {
	if a.cubic_hectometersLazy != nil {
		return *a.cubic_hectometersLazy
	}
	cubic_hectometers := a.convertFromBase(VolumeCubicHectometer)
	a.cubic_hectometersLazy = &cubic_hectometers
	return cubic_hectometers
}

// CubicDecimeters returns the Volume value in CubicDecimeters.
//
// 
func (a *Volume) CubicDecimeters() float64 {
	if a.cubic_decimetersLazy != nil {
		return *a.cubic_decimetersLazy
	}
	cubic_decimeters := a.convertFromBase(VolumeCubicDecimeter)
	a.cubic_decimetersLazy = &cubic_decimeters
	return cubic_decimeters
}

// CubicCentimeters returns the Volume value in CubicCentimeters.
//
// 
func (a *Volume) CubicCentimeters() float64 {
	if a.cubic_centimetersLazy != nil {
		return *a.cubic_centimetersLazy
	}
	cubic_centimeters := a.convertFromBase(VolumeCubicCentimeter)
	a.cubic_centimetersLazy = &cubic_centimeters
	return cubic_centimeters
}

// CubicMillimeters returns the Volume value in CubicMillimeters.
//
// 
func (a *Volume) CubicMillimeters() float64 {
	if a.cubic_millimetersLazy != nil {
		return *a.cubic_millimetersLazy
	}
	cubic_millimeters := a.convertFromBase(VolumeCubicMillimeter)
	a.cubic_millimetersLazy = &cubic_millimeters
	return cubic_millimeters
}

// CubicMicrometers returns the Volume value in CubicMicrometers.
//
// 
func (a *Volume) CubicMicrometers() float64 {
	if a.cubic_micrometersLazy != nil {
		return *a.cubic_micrometersLazy
	}
	cubic_micrometers := a.convertFromBase(VolumeCubicMicrometer)
	a.cubic_micrometersLazy = &cubic_micrometers
	return cubic_micrometers
}

// CubicMiles returns the Volume value in CubicMiles.
//
// 
func (a *Volume) CubicMiles() float64 {
	if a.cubic_milesLazy != nil {
		return *a.cubic_milesLazy
	}
	cubic_miles := a.convertFromBase(VolumeCubicMile)
	a.cubic_milesLazy = &cubic_miles
	return cubic_miles
}

// CubicYards returns the Volume value in CubicYards.
//
// 
func (a *Volume) CubicYards() float64 {
	if a.cubic_yardsLazy != nil {
		return *a.cubic_yardsLazy
	}
	cubic_yards := a.convertFromBase(VolumeCubicYard)
	a.cubic_yardsLazy = &cubic_yards
	return cubic_yards
}

// CubicFeet returns the Volume value in CubicFeet.
//
// 
func (a *Volume) CubicFeet() float64 {
	if a.cubic_feetLazy != nil {
		return *a.cubic_feetLazy
	}
	cubic_feet := a.convertFromBase(VolumeCubicFoot)
	a.cubic_feetLazy = &cubic_feet
	return cubic_feet
}

// CubicInches returns the Volume value in CubicInches.
//
// 
func (a *Volume) CubicInches() float64 {
	if a.cubic_inchesLazy != nil {
		return *a.cubic_inchesLazy
	}
	cubic_inches := a.convertFromBase(VolumeCubicInch)
	a.cubic_inchesLazy = &cubic_inches
	return cubic_inches
}

// ImperialGallons returns the Volume value in ImperialGallons.
//
// The British imperial gallon (frequently called simply "gallon") is defined as exactly 4.54609 litres.
func (a *Volume) ImperialGallons() float64 {
	if a.imperial_gallonsLazy != nil {
		return *a.imperial_gallonsLazy
	}
	imperial_gallons := a.convertFromBase(VolumeImperialGallon)
	a.imperial_gallonsLazy = &imperial_gallons
	return imperial_gallons
}

// ImperialOunces returns the Volume value in ImperialOunces.
//
// 
func (a *Volume) ImperialOunces() float64 {
	if a.imperial_ouncesLazy != nil {
		return *a.imperial_ouncesLazy
	}
	imperial_ounces := a.convertFromBase(VolumeImperialOunce)
	a.imperial_ouncesLazy = &imperial_ounces
	return imperial_ounces
}

// UsGallons returns the Volume value in UsGallons.
//
// The US liquid gallon (frequently called simply "gallon") is legally defined as 231 cubic inches, which is exactly 3.785411784 litres.
func (a *Volume) UsGallons() float64 {
	if a.us_gallonsLazy != nil {
		return *a.us_gallonsLazy
	}
	us_gallons := a.convertFromBase(VolumeUsGallon)
	a.us_gallonsLazy = &us_gallons
	return us_gallons
}

// UsOunces returns the Volume value in UsOunces.
//
// 
func (a *Volume) UsOunces() float64 {
	if a.us_ouncesLazy != nil {
		return *a.us_ouncesLazy
	}
	us_ounces := a.convertFromBase(VolumeUsOunce)
	a.us_ouncesLazy = &us_ounces
	return us_ounces
}

// UsTablespoons returns the Volume value in UsTablespoons.
//
// 
func (a *Volume) UsTablespoons() float64 {
	if a.us_tablespoonsLazy != nil {
		return *a.us_tablespoonsLazy
	}
	us_tablespoons := a.convertFromBase(VolumeUsTablespoon)
	a.us_tablespoonsLazy = &us_tablespoons
	return us_tablespoons
}

// AuTablespoons returns the Volume value in AuTablespoons.
//
// 
func (a *Volume) AuTablespoons() float64 {
	if a.au_tablespoonsLazy != nil {
		return *a.au_tablespoonsLazy
	}
	au_tablespoons := a.convertFromBase(VolumeAuTablespoon)
	a.au_tablespoonsLazy = &au_tablespoons
	return au_tablespoons
}

// UkTablespoons returns the Volume value in UkTablespoons.
//
// 
func (a *Volume) UkTablespoons() float64 {
	if a.uk_tablespoonsLazy != nil {
		return *a.uk_tablespoonsLazy
	}
	uk_tablespoons := a.convertFromBase(VolumeUkTablespoon)
	a.uk_tablespoonsLazy = &uk_tablespoons
	return uk_tablespoons
}

// MetricTeaspoons returns the Volume value in MetricTeaspoons.
//
// 
func (a *Volume) MetricTeaspoons() float64 {
	if a.metric_teaspoonsLazy != nil {
		return *a.metric_teaspoonsLazy
	}
	metric_teaspoons := a.convertFromBase(VolumeMetricTeaspoon)
	a.metric_teaspoonsLazy = &metric_teaspoons
	return metric_teaspoons
}

// UsTeaspoons returns the Volume value in UsTeaspoons.
//
// 
func (a *Volume) UsTeaspoons() float64 {
	if a.us_teaspoonsLazy != nil {
		return *a.us_teaspoonsLazy
	}
	us_teaspoons := a.convertFromBase(VolumeUsTeaspoon)
	a.us_teaspoonsLazy = &us_teaspoons
	return us_teaspoons
}

// MetricCups returns the Volume value in MetricCups.
//
// 
func (a *Volume) MetricCups() float64 {
	if a.metric_cupsLazy != nil {
		return *a.metric_cupsLazy
	}
	metric_cups := a.convertFromBase(VolumeMetricCup)
	a.metric_cupsLazy = &metric_cups
	return metric_cups
}

// UsCustomaryCups returns the Volume value in UsCustomaryCups.
//
// 
func (a *Volume) UsCustomaryCups() float64 {
	if a.us_customary_cupsLazy != nil {
		return *a.us_customary_cupsLazy
	}
	us_customary_cups := a.convertFromBase(VolumeUsCustomaryCup)
	a.us_customary_cupsLazy = &us_customary_cups
	return us_customary_cups
}

// UsLegalCups returns the Volume value in UsLegalCups.
//
// 
func (a *Volume) UsLegalCups() float64 {
	if a.us_legal_cupsLazy != nil {
		return *a.us_legal_cupsLazy
	}
	us_legal_cups := a.convertFromBase(VolumeUsLegalCup)
	a.us_legal_cupsLazy = &us_legal_cups
	return us_legal_cups
}

// OilBarrels returns the Volume value in OilBarrels.
//
// 
func (a *Volume) OilBarrels() float64 {
	if a.oil_barrelsLazy != nil {
		return *a.oil_barrelsLazy
	}
	oil_barrels := a.convertFromBase(VolumeOilBarrel)
	a.oil_barrelsLazy = &oil_barrels
	return oil_barrels
}

// UsBeerBarrels returns the Volume value in UsBeerBarrels.
//
// 
func (a *Volume) UsBeerBarrels() float64 {
	if a.us_beer_barrelsLazy != nil {
		return *a.us_beer_barrelsLazy
	}
	us_beer_barrels := a.convertFromBase(VolumeUsBeerBarrel)
	a.us_beer_barrelsLazy = &us_beer_barrels
	return us_beer_barrels
}

// ImperialBeerBarrels returns the Volume value in ImperialBeerBarrels.
//
// 
func (a *Volume) ImperialBeerBarrels() float64 {
	if a.imperial_beer_barrelsLazy != nil {
		return *a.imperial_beer_barrelsLazy
	}
	imperial_beer_barrels := a.convertFromBase(VolumeImperialBeerBarrel)
	a.imperial_beer_barrelsLazy = &imperial_beer_barrels
	return imperial_beer_barrels
}

// UsQuarts returns the Volume value in UsQuarts.
//
// 
func (a *Volume) UsQuarts() float64 {
	if a.us_quartsLazy != nil {
		return *a.us_quartsLazy
	}
	us_quarts := a.convertFromBase(VolumeUsQuart)
	a.us_quartsLazy = &us_quarts
	return us_quarts
}

// ImperialQuarts returns the Volume value in ImperialQuarts.
//
// 
func (a *Volume) ImperialQuarts() float64 {
	if a.imperial_quartsLazy != nil {
		return *a.imperial_quartsLazy
	}
	imperial_quarts := a.convertFromBase(VolumeImperialQuart)
	a.imperial_quartsLazy = &imperial_quarts
	return imperial_quarts
}

// UsPints returns the Volume value in UsPints.
//
// 
func (a *Volume) UsPints() float64 {
	if a.us_pintsLazy != nil {
		return *a.us_pintsLazy
	}
	us_pints := a.convertFromBase(VolumeUsPint)
	a.us_pintsLazy = &us_pints
	return us_pints
}

// AcreFeet returns the Volume value in AcreFeet.
//
// 
func (a *Volume) AcreFeet() float64 {
	if a.acre_feetLazy != nil {
		return *a.acre_feetLazy
	}
	acre_feet := a.convertFromBase(VolumeAcreFoot)
	a.acre_feetLazy = &acre_feet
	return acre_feet
}

// ImperialPints returns the Volume value in ImperialPints.
//
// 
func (a *Volume) ImperialPints() float64 {
	if a.imperial_pintsLazy != nil {
		return *a.imperial_pintsLazy
	}
	imperial_pints := a.convertFromBase(VolumeImperialPint)
	a.imperial_pintsLazy = &imperial_pints
	return imperial_pints
}

// BoardFeet returns the Volume value in BoardFeet.
//
// 
func (a *Volume) BoardFeet() float64 {
	if a.board_feetLazy != nil {
		return *a.board_feetLazy
	}
	board_feet := a.convertFromBase(VolumeBoardFoot)
	a.board_feetLazy = &board_feet
	return board_feet
}

// Nanoliters returns the Volume value in Nanoliters.
//
// 
func (a *Volume) Nanoliters() float64 {
	if a.nanolitersLazy != nil {
		return *a.nanolitersLazy
	}
	nanoliters := a.convertFromBase(VolumeNanoliter)
	a.nanolitersLazy = &nanoliters
	return nanoliters
}

// Microliters returns the Volume value in Microliters.
//
// 
func (a *Volume) Microliters() float64 {
	if a.microlitersLazy != nil {
		return *a.microlitersLazy
	}
	microliters := a.convertFromBase(VolumeMicroliter)
	a.microlitersLazy = &microliters
	return microliters
}

// Milliliters returns the Volume value in Milliliters.
//
// 
func (a *Volume) Milliliters() float64 {
	if a.millilitersLazy != nil {
		return *a.millilitersLazy
	}
	milliliters := a.convertFromBase(VolumeMilliliter)
	a.millilitersLazy = &milliliters
	return milliliters
}

// Centiliters returns the Volume value in Centiliters.
//
// 
func (a *Volume) Centiliters() float64 {
	if a.centilitersLazy != nil {
		return *a.centilitersLazy
	}
	centiliters := a.convertFromBase(VolumeCentiliter)
	a.centilitersLazy = &centiliters
	return centiliters
}

// Deciliters returns the Volume value in Deciliters.
//
// 
func (a *Volume) Deciliters() float64 {
	if a.decilitersLazy != nil {
		return *a.decilitersLazy
	}
	deciliters := a.convertFromBase(VolumeDeciliter)
	a.decilitersLazy = &deciliters
	return deciliters
}

// Decaliters returns the Volume value in Decaliters.
//
// 
func (a *Volume) Decaliters() float64 {
	if a.decalitersLazy != nil {
		return *a.decalitersLazy
	}
	decaliters := a.convertFromBase(VolumeDecaliter)
	a.decalitersLazy = &decaliters
	return decaliters
}

// Hectoliters returns the Volume value in Hectoliters.
//
// 
func (a *Volume) Hectoliters() float64 {
	if a.hectolitersLazy != nil {
		return *a.hectolitersLazy
	}
	hectoliters := a.convertFromBase(VolumeHectoliter)
	a.hectolitersLazy = &hectoliters
	return hectoliters
}

// Kiloliters returns the Volume value in Kiloliters.
//
// 
func (a *Volume) Kiloliters() float64 {
	if a.kilolitersLazy != nil {
		return *a.kilolitersLazy
	}
	kiloliters := a.convertFromBase(VolumeKiloliter)
	a.kilolitersLazy = &kiloliters
	return kiloliters
}

// Megaliters returns the Volume value in Megaliters.
//
// 
func (a *Volume) Megaliters() float64 {
	if a.megalitersLazy != nil {
		return *a.megalitersLazy
	}
	megaliters := a.convertFromBase(VolumeMegaliter)
	a.megalitersLazy = &megaliters
	return megaliters
}

// HectocubicMeters returns the Volume value in HectocubicMeters.
//
// 
func (a *Volume) HectocubicMeters() float64 {
	if a.hectocubic_metersLazy != nil {
		return *a.hectocubic_metersLazy
	}
	hectocubic_meters := a.convertFromBase(VolumeHectocubicMeter)
	a.hectocubic_metersLazy = &hectocubic_meters
	return hectocubic_meters
}

// KilocubicMeters returns the Volume value in KilocubicMeters.
//
// 
func (a *Volume) KilocubicMeters() float64 {
	if a.kilocubic_metersLazy != nil {
		return *a.kilocubic_metersLazy
	}
	kilocubic_meters := a.convertFromBase(VolumeKilocubicMeter)
	a.kilocubic_metersLazy = &kilocubic_meters
	return kilocubic_meters
}

// HectocubicFeet returns the Volume value in HectocubicFeet.
//
// 
func (a *Volume) HectocubicFeet() float64 {
	if a.hectocubic_feetLazy != nil {
		return *a.hectocubic_feetLazy
	}
	hectocubic_feet := a.convertFromBase(VolumeHectocubicFoot)
	a.hectocubic_feetLazy = &hectocubic_feet
	return hectocubic_feet
}

// KilocubicFeet returns the Volume value in KilocubicFeet.
//
// 
func (a *Volume) KilocubicFeet() float64 {
	if a.kilocubic_feetLazy != nil {
		return *a.kilocubic_feetLazy
	}
	kilocubic_feet := a.convertFromBase(VolumeKilocubicFoot)
	a.kilocubic_feetLazy = &kilocubic_feet
	return kilocubic_feet
}

// MegacubicFeet returns the Volume value in MegacubicFeet.
//
// 
func (a *Volume) MegacubicFeet() float64 {
	if a.megacubic_feetLazy != nil {
		return *a.megacubic_feetLazy
	}
	megacubic_feet := a.convertFromBase(VolumeMegacubicFoot)
	a.megacubic_feetLazy = &megacubic_feet
	return megacubic_feet
}

// KiloimperialGallons returns the Volume value in KiloimperialGallons.
//
// 
func (a *Volume) KiloimperialGallons() float64 {
	if a.kiloimperial_gallonsLazy != nil {
		return *a.kiloimperial_gallonsLazy
	}
	kiloimperial_gallons := a.convertFromBase(VolumeKiloimperialGallon)
	a.kiloimperial_gallonsLazy = &kiloimperial_gallons
	return kiloimperial_gallons
}

// MegaimperialGallons returns the Volume value in MegaimperialGallons.
//
// 
func (a *Volume) MegaimperialGallons() float64 {
	if a.megaimperial_gallonsLazy != nil {
		return *a.megaimperial_gallonsLazy
	}
	megaimperial_gallons := a.convertFromBase(VolumeMegaimperialGallon)
	a.megaimperial_gallonsLazy = &megaimperial_gallons
	return megaimperial_gallons
}

// DecausGallons returns the Volume value in DecausGallons.
//
// 
func (a *Volume) DecausGallons() float64 {
	if a.decaus_gallonsLazy != nil {
		return *a.decaus_gallonsLazy
	}
	decaus_gallons := a.convertFromBase(VolumeDecausGallon)
	a.decaus_gallonsLazy = &decaus_gallons
	return decaus_gallons
}

// DeciusGallons returns the Volume value in DeciusGallons.
//
// 
func (a *Volume) DeciusGallons() float64 {
	if a.decius_gallonsLazy != nil {
		return *a.decius_gallonsLazy
	}
	decius_gallons := a.convertFromBase(VolumeDeciusGallon)
	a.decius_gallonsLazy = &decius_gallons
	return decius_gallons
}

// HectousGallons returns the Volume value in HectousGallons.
//
// 
func (a *Volume) HectousGallons() float64 {
	if a.hectous_gallonsLazy != nil {
		return *a.hectous_gallonsLazy
	}
	hectous_gallons := a.convertFromBase(VolumeHectousGallon)
	a.hectous_gallonsLazy = &hectous_gallons
	return hectous_gallons
}

// KilousGallons returns the Volume value in KilousGallons.
//
// 
func (a *Volume) KilousGallons() float64 {
	if a.kilous_gallonsLazy != nil {
		return *a.kilous_gallonsLazy
	}
	kilous_gallons := a.convertFromBase(VolumeKilousGallon)
	a.kilous_gallonsLazy = &kilous_gallons
	return kilous_gallons
}

// MegausGallons returns the Volume value in MegausGallons.
//
// 
func (a *Volume) MegausGallons() float64 {
	if a.megaus_gallonsLazy != nil {
		return *a.megaus_gallonsLazy
	}
	megaus_gallons := a.convertFromBase(VolumeMegausGallon)
	a.megaus_gallonsLazy = &megaus_gallons
	return megaus_gallons
}


// ToDto creates a VolumeDto representation from the Volume instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeter by default.
func (a *Volume) ToDto(holdInUnit *VolumeUnits) VolumeDto {
	if holdInUnit == nil {
		defaultUnit := VolumeCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return VolumeDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Volume instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeter by default.
func (a *Volume) ToDtoJSON(holdInUnit *VolumeUnits) ([]byte, error) {
	// Convert to VolumeDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Volume to a specific unit value.
// The function uses the provided unit type (VolumeUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Volume) Convert(toUnit VolumeUnits) float64 {
	switch toUnit { 
    case VolumeLiter:
		return a.Liters()
    case VolumeCubicMeter:
		return a.CubicMeters()
    case VolumeCubicKilometer:
		return a.CubicKilometers()
    case VolumeCubicHectometer:
		return a.CubicHectometers()
    case VolumeCubicDecimeter:
		return a.CubicDecimeters()
    case VolumeCubicCentimeter:
		return a.CubicCentimeters()
    case VolumeCubicMillimeter:
		return a.CubicMillimeters()
    case VolumeCubicMicrometer:
		return a.CubicMicrometers()
    case VolumeCubicMile:
		return a.CubicMiles()
    case VolumeCubicYard:
		return a.CubicYards()
    case VolumeCubicFoot:
		return a.CubicFeet()
    case VolumeCubicInch:
		return a.CubicInches()
    case VolumeImperialGallon:
		return a.ImperialGallons()
    case VolumeImperialOunce:
		return a.ImperialOunces()
    case VolumeUsGallon:
		return a.UsGallons()
    case VolumeUsOunce:
		return a.UsOunces()
    case VolumeUsTablespoon:
		return a.UsTablespoons()
    case VolumeAuTablespoon:
		return a.AuTablespoons()
    case VolumeUkTablespoon:
		return a.UkTablespoons()
    case VolumeMetricTeaspoon:
		return a.MetricTeaspoons()
    case VolumeUsTeaspoon:
		return a.UsTeaspoons()
    case VolumeMetricCup:
		return a.MetricCups()
    case VolumeUsCustomaryCup:
		return a.UsCustomaryCups()
    case VolumeUsLegalCup:
		return a.UsLegalCups()
    case VolumeOilBarrel:
		return a.OilBarrels()
    case VolumeUsBeerBarrel:
		return a.UsBeerBarrels()
    case VolumeImperialBeerBarrel:
		return a.ImperialBeerBarrels()
    case VolumeUsQuart:
		return a.UsQuarts()
    case VolumeImperialQuart:
		return a.ImperialQuarts()
    case VolumeUsPint:
		return a.UsPints()
    case VolumeAcreFoot:
		return a.AcreFeet()
    case VolumeImperialPint:
		return a.ImperialPints()
    case VolumeBoardFoot:
		return a.BoardFeet()
    case VolumeNanoliter:
		return a.Nanoliters()
    case VolumeMicroliter:
		return a.Microliters()
    case VolumeMilliliter:
		return a.Milliliters()
    case VolumeCentiliter:
		return a.Centiliters()
    case VolumeDeciliter:
		return a.Deciliters()
    case VolumeDecaliter:
		return a.Decaliters()
    case VolumeHectoliter:
		return a.Hectoliters()
    case VolumeKiloliter:
		return a.Kiloliters()
    case VolumeMegaliter:
		return a.Megaliters()
    case VolumeHectocubicMeter:
		return a.HectocubicMeters()
    case VolumeKilocubicMeter:
		return a.KilocubicMeters()
    case VolumeHectocubicFoot:
		return a.HectocubicFeet()
    case VolumeKilocubicFoot:
		return a.KilocubicFeet()
    case VolumeMegacubicFoot:
		return a.MegacubicFeet()
    case VolumeKiloimperialGallon:
		return a.KiloimperialGallons()
    case VolumeMegaimperialGallon:
		return a.MegaimperialGallons()
    case VolumeDecausGallon:
		return a.DecausGallons()
    case VolumeDeciusGallon:
		return a.DeciusGallons()
    case VolumeHectousGallon:
		return a.HectousGallons()
    case VolumeKilousGallon:
		return a.KilousGallons()
    case VolumeMegausGallon:
		return a.MegausGallons()
	default:
		return math.NaN()
	}
}

func (a *Volume) convertFromBase(toUnit VolumeUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumeLiter:
		return (value * 1e3) 
	case VolumeCubicMeter:
		return (value) 
	case VolumeCubicKilometer:
		return (value / 1e9) 
	case VolumeCubicHectometer:
		return (value / 1e6) 
	case VolumeCubicDecimeter:
		return (value * 1e3) 
	case VolumeCubicCentimeter:
		return (value * 1e6) 
	case VolumeCubicMillimeter:
		return (value * 1e9) 
	case VolumeCubicMicrometer:
		return (value * 1e18) 
	case VolumeCubicMile:
		return (value / 4.16818182544058e9) 
	case VolumeCubicYard:
		return (value / 0.764554858) 
	case VolumeCubicFoot:
		return (value / 2.8316846592e-2) 
	case VolumeCubicInch:
		return (value / 1.6387064e-5) 
	case VolumeImperialGallon:
		return (value / 0.00454609) 
	case VolumeImperialOunce:
		return (value / 2.8413062499962901241875439064617e-5) 
	case VolumeUsGallon:
		return (value / 0.003785411784) 
	case VolumeUsOunce:
		return (value / 2.957352956253760505068307980135e-5) 
	case VolumeUsTablespoon:
		return (value / 1.478676478125e-5) 
	case VolumeAuTablespoon:
		return (value / 2e-5) 
	case VolumeUkTablespoon:
		return (value / 1.5e-5) 
	case VolumeMetricTeaspoon:
		return (value / 0.5e-5) 
	case VolumeUsTeaspoon:
		return (value / 4.92892159375e-6) 
	case VolumeMetricCup:
		return (value / 0.00025) 
	case VolumeUsCustomaryCup:
		return (value / 0.0002365882365) 
	case VolumeUsLegalCup:
		return (value / 0.00024) 
	case VolumeOilBarrel:
		return (value / 0.158987294928) 
	case VolumeUsBeerBarrel:
		return (value / 0.1173477658) 
	case VolumeImperialBeerBarrel:
		return (value / 0.16365924) 
	case VolumeUsQuart:
		return (value / 9.46352946e-4) 
	case VolumeImperialQuart:
		return (value / 1.1365225e-3) 
	case VolumeUsPint:
		return (value / 4.73176473e-4) 
	case VolumeAcreFoot:
		return (value * 0.000810714) 
	case VolumeImperialPint:
		return (value / 5.6826125e-4) 
	case VolumeBoardFoot:
		return (value / 2.3597372158e-3) 
	case VolumeNanoliter:
		return ((value * 1e3) / 1e-09) 
	case VolumeMicroliter:
		return ((value * 1e3) / 1e-06) 
	case VolumeMilliliter:
		return ((value * 1e3) / 0.001) 
	case VolumeCentiliter:
		return ((value * 1e3) / 0.01) 
	case VolumeDeciliter:
		return ((value * 1e3) / 0.1) 
	case VolumeDecaliter:
		return ((value * 1e3) / 10.0) 
	case VolumeHectoliter:
		return ((value * 1e3) / 100.0) 
	case VolumeKiloliter:
		return ((value * 1e3) / 1000.0) 
	case VolumeMegaliter:
		return ((value * 1e3) / 1000000.0) 
	case VolumeHectocubicMeter:
		return ((value) / 100.0) 
	case VolumeKilocubicMeter:
		return ((value) / 1000.0) 
	case VolumeHectocubicFoot:
		return ((value / 2.8316846592e-2) / 100.0) 
	case VolumeKilocubicFoot:
		return ((value / 2.8316846592e-2) / 1000.0) 
	case VolumeMegacubicFoot:
		return ((value / 2.8316846592e-2) / 1000000.0) 
	case VolumeKiloimperialGallon:
		return ((value / 0.00454609) / 1000.0) 
	case VolumeMegaimperialGallon:
		return ((value / 0.00454609) / 1000000.0) 
	case VolumeDecausGallon:
		return ((value / 0.003785411784) / 10.0) 
	case VolumeDeciusGallon:
		return ((value / 0.003785411784) / 0.1) 
	case VolumeHectousGallon:
		return ((value / 0.003785411784) / 100.0) 
	case VolumeKilousGallon:
		return ((value / 0.003785411784) / 1000.0) 
	case VolumeMegausGallon:
		return ((value / 0.003785411784) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Volume) convertToBase(value float64, fromUnit VolumeUnits) float64 {
	switch fromUnit { 
	case VolumeLiter:
		return (value / 1e3) 
	case VolumeCubicMeter:
		return (value) 
	case VolumeCubicKilometer:
		return (value * 1e9) 
	case VolumeCubicHectometer:
		return (value * 1e6) 
	case VolumeCubicDecimeter:
		return (value / 1e3) 
	case VolumeCubicCentimeter:
		return (value / 1e6) 
	case VolumeCubicMillimeter:
		return (value / 1e9) 
	case VolumeCubicMicrometer:
		return (value / 1e18) 
	case VolumeCubicMile:
		return (value * 4.16818182544058e9) 
	case VolumeCubicYard:
		return (value * 0.764554858) 
	case VolumeCubicFoot:
		return (value * 2.8316846592e-2) 
	case VolumeCubicInch:
		return (value * 1.6387064e-5) 
	case VolumeImperialGallon:
		return (value * 0.00454609) 
	case VolumeImperialOunce:
		return (value * 2.8413062499962901241875439064617e-5) 
	case VolumeUsGallon:
		return (value * 0.003785411784) 
	case VolumeUsOunce:
		return (value * 2.957352956253760505068307980135e-5) 
	case VolumeUsTablespoon:
		return (value * 1.478676478125e-5) 
	case VolumeAuTablespoon:
		return (value * 2e-5) 
	case VolumeUkTablespoon:
		return (value * 1.5e-5) 
	case VolumeMetricTeaspoon:
		return (value * 0.5e-5) 
	case VolumeUsTeaspoon:
		return (value * 4.92892159375e-6) 
	case VolumeMetricCup:
		return (value * 0.00025) 
	case VolumeUsCustomaryCup:
		return (value * 0.0002365882365) 
	case VolumeUsLegalCup:
		return (value * 0.00024) 
	case VolumeOilBarrel:
		return (value * 0.158987294928) 
	case VolumeUsBeerBarrel:
		return (value * 0.1173477658) 
	case VolumeImperialBeerBarrel:
		return (value * 0.16365924) 
	case VolumeUsQuart:
		return (value * 9.46352946e-4) 
	case VolumeImperialQuart:
		return (value * 1.1365225e-3) 
	case VolumeUsPint:
		return (value * 4.73176473e-4) 
	case VolumeAcreFoot:
		return (value / 0.000810714) 
	case VolumeImperialPint:
		return (value * 5.6826125e-4) 
	case VolumeBoardFoot:
		return (value * 2.3597372158e-3) 
	case VolumeNanoliter:
		return ((value / 1e3) * 1e-09) 
	case VolumeMicroliter:
		return ((value / 1e3) * 1e-06) 
	case VolumeMilliliter:
		return ((value / 1e3) * 0.001) 
	case VolumeCentiliter:
		return ((value / 1e3) * 0.01) 
	case VolumeDeciliter:
		return ((value / 1e3) * 0.1) 
	case VolumeDecaliter:
		return ((value / 1e3) * 10.0) 
	case VolumeHectoliter:
		return ((value / 1e3) * 100.0) 
	case VolumeKiloliter:
		return ((value / 1e3) * 1000.0) 
	case VolumeMegaliter:
		return ((value / 1e3) * 1000000.0) 
	case VolumeHectocubicMeter:
		return ((value) * 100.0) 
	case VolumeKilocubicMeter:
		return ((value) * 1000.0) 
	case VolumeHectocubicFoot:
		return ((value * 2.8316846592e-2) * 100.0) 
	case VolumeKilocubicFoot:
		return ((value * 2.8316846592e-2) * 1000.0) 
	case VolumeMegacubicFoot:
		return ((value * 2.8316846592e-2) * 1000000.0) 
	case VolumeKiloimperialGallon:
		return ((value * 0.00454609) * 1000.0) 
	case VolumeMegaimperialGallon:
		return ((value * 0.00454609) * 1000000.0) 
	case VolumeDecausGallon:
		return ((value * 0.003785411784) * 10.0) 
	case VolumeDeciusGallon:
		return ((value * 0.003785411784) * 0.1) 
	case VolumeHectousGallon:
		return ((value * 0.003785411784) * 100.0) 
	case VolumeKilousGallon:
		return ((value * 0.003785411784) * 1000.0) 
	case VolumeMegausGallon:
		return ((value * 0.003785411784) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Volume in the default unit (CubicMeter),
// formatted to two decimal places.
func (a Volume) String() string {
	return a.ToString(VolumeCubicMeter, 2)
}

// ToString formats the Volume value as a string with the specified unit and fractional digits.
// It converts the Volume to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Volume value will be converted (e.g., CubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Volume with the unit abbreviation.
func (a *Volume) ToString(unit VolumeUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetVolumeAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetVolumeAbbreviation(unit))
}

// Equals checks if the given Volume is equal to the current Volume.
//
// Parameters:
//    other: The Volume to compare against.
//
// Returns:
//    bool: Returns true if both Volume are equal, false otherwise.
func (a *Volume) Equals(other *Volume) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Volume with another Volume.
// It returns -1 if the current Volume is less than the other Volume, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Volume to compare against.
//
// Returns:
//    int: -1 if the current Volume is less, 1 if greater, and 0 if equal.
func (a *Volume) CompareTo(other *Volume) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Volume to the current Volume and returns the result.
// The result is a new Volume instance with the sum of the values.
//
// Parameters:
//    other: The Volume to add to the current Volume.
//
// Returns:
//    *Volume: A new Volume instance representing the sum of both Volume.
func (a *Volume) Add(other *Volume) *Volume {
	return &Volume{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Volume from the current Volume and returns the result.
// The result is a new Volume instance with the difference of the values.
//
// Parameters:
//    other: The Volume to subtract from the current Volume.
//
// Returns:
//    *Volume: A new Volume instance representing the difference of both Volume.
func (a *Volume) Subtract(other *Volume) *Volume {
	return &Volume{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Volume by the given Volume and returns the result.
// The result is a new Volume instance with the product of the values.
//
// Parameters:
//    other: The Volume to multiply with the current Volume.
//
// Returns:
//    *Volume: A new Volume instance representing the product of both Volume.
func (a *Volume) Multiply(other *Volume) *Volume {
	return &Volume{value: a.value * other.BaseValue()}
}

// Divide divides the current Volume by the given Volume and returns the result.
// The result is a new Volume instance with the quotient of the values.
//
// Parameters:
//    other: The Volume to divide the current Volume by.
//
// Returns:
//    *Volume: A new Volume instance representing the quotient of both Volume.
func (a *Volume) Divide(other *Volume) *Volume {
	return &Volume{value: a.value / other.BaseValue()}
}

// GetVolumeAbbreviation gets the unit abbreviation.
func GetVolumeAbbreviation(unit VolumeUnits) string {
	switch unit { 
	case VolumeLiter:
		return "l" 
	case VolumeCubicMeter:
		return "m³" 
	case VolumeCubicKilometer:
		return "km³" 
	case VolumeCubicHectometer:
		return "hm³" 
	case VolumeCubicDecimeter:
		return "dm³" 
	case VolumeCubicCentimeter:
		return "cm³" 
	case VolumeCubicMillimeter:
		return "mm³" 
	case VolumeCubicMicrometer:
		return "µm³" 
	case VolumeCubicMile:
		return "mi³" 
	case VolumeCubicYard:
		return "yd³" 
	case VolumeCubicFoot:
		return "ft³" 
	case VolumeCubicInch:
		return "in³" 
	case VolumeImperialGallon:
		return "gal (imp.)" 
	case VolumeImperialOunce:
		return "oz (imp.)" 
	case VolumeUsGallon:
		return "gal (U.S.)" 
	case VolumeUsOunce:
		return "oz (U.S.)" 
	case VolumeUsTablespoon:
		return "" 
	case VolumeAuTablespoon:
		return "" 
	case VolumeUkTablespoon:
		return "" 
	case VolumeMetricTeaspoon:
		return "tsp" 
	case VolumeUsTeaspoon:
		return "" 
	case VolumeMetricCup:
		return "" 
	case VolumeUsCustomaryCup:
		return "" 
	case VolumeUsLegalCup:
		return "" 
	case VolumeOilBarrel:
		return "bbl" 
	case VolumeUsBeerBarrel:
		return "bl (U.S.)" 
	case VolumeImperialBeerBarrel:
		return "bl (imp.)" 
	case VolumeUsQuart:
		return "qt (U.S.)" 
	case VolumeImperialQuart:
		return "qt (imp.)" 
	case VolumeUsPint:
		return "pt (U.S.)" 
	case VolumeAcreFoot:
		return "ac-ft" 
	case VolumeImperialPint:
		return "pt (imp.)" 
	case VolumeBoardFoot:
		return "bf" 
	case VolumeNanoliter:
		return "nl" 
	case VolumeMicroliter:
		return "μl" 
	case VolumeMilliliter:
		return "ml" 
	case VolumeCentiliter:
		return "cl" 
	case VolumeDeciliter:
		return "dl" 
	case VolumeDecaliter:
		return "dal" 
	case VolumeHectoliter:
		return "hl" 
	case VolumeKiloliter:
		return "kl" 
	case VolumeMegaliter:
		return "Ml" 
	case VolumeHectocubicMeter:
		return "hm³" 
	case VolumeKilocubicMeter:
		return "km³" 
	case VolumeHectocubicFoot:
		return "hft³" 
	case VolumeKilocubicFoot:
		return "kft³" 
	case VolumeMegacubicFoot:
		return "Mft³" 
	case VolumeKiloimperialGallon:
		return "kgal (imp.)" 
	case VolumeMegaimperialGallon:
		return "Mgal (imp.)" 
	case VolumeDecausGallon:
		return "dagal (U.S.)" 
	case VolumeDeciusGallon:
		return "dgal (U.S.)" 
	case VolumeHectousGallon:
		return "hgal (U.S.)" 
	case VolumeKilousGallon:
		return "kgal (U.S.)" 
	case VolumeMegausGallon:
		return "Mgal (U.S.)" 
	default:
		return ""
	}
}