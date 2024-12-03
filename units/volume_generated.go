// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeUnits enumeration
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

// VolumeDto represents an Volume
type VolumeDto struct {
	Value float64
	Unit  VolumeUnits
}

// VolumeDtoFactory struct to group related functions
type VolumeDtoFactory struct{}

func (udf VolumeDtoFactory) FromJSON(data []byte) (*VolumeDto, error) {
	a := VolumeDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumeDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Volume struct
type Volume struct {
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

// VolumeFactory struct to group related functions
type VolumeFactory struct{}

func (uf VolumeFactory) CreateVolume(value float64, unit VolumeUnits) (*Volume, error) {
	return newVolume(value, unit)
}

func (uf VolumeFactory) FromDto(dto VolumeDto) (*Volume, error) {
	return newVolume(dto.Value, dto.Unit)
}

func (uf VolumeFactory) FromDtoJSON(data []byte) (*Volume, error) {
	unitDto, err := VolumeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumeFactory{}.FromDto(*unitDto)
}


// FromLiter creates a new Volume instance from Liter.
func (uf VolumeFactory) FromLiters(value float64) (*Volume, error) {
	return newVolume(value, VolumeLiter)
}

// FromCubicMeter creates a new Volume instance from CubicMeter.
func (uf VolumeFactory) FromCubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMeter)
}

// FromCubicKilometer creates a new Volume instance from CubicKilometer.
func (uf VolumeFactory) FromCubicKilometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicKilometer)
}

// FromCubicHectometer creates a new Volume instance from CubicHectometer.
func (uf VolumeFactory) FromCubicHectometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicHectometer)
}

// FromCubicDecimeter creates a new Volume instance from CubicDecimeter.
func (uf VolumeFactory) FromCubicDecimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicDecimeter)
}

// FromCubicCentimeter creates a new Volume instance from CubicCentimeter.
func (uf VolumeFactory) FromCubicCentimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicCentimeter)
}

// FromCubicMillimeter creates a new Volume instance from CubicMillimeter.
func (uf VolumeFactory) FromCubicMillimeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMillimeter)
}

// FromCubicMicrometer creates a new Volume instance from CubicMicrometer.
func (uf VolumeFactory) FromCubicMicrometers(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMicrometer)
}

// FromCubicMile creates a new Volume instance from CubicMile.
func (uf VolumeFactory) FromCubicMiles(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicMile)
}

// FromCubicYard creates a new Volume instance from CubicYard.
func (uf VolumeFactory) FromCubicYards(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicYard)
}

// FromCubicFoot creates a new Volume instance from CubicFoot.
func (uf VolumeFactory) FromCubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicFoot)
}

// FromCubicInch creates a new Volume instance from CubicInch.
func (uf VolumeFactory) FromCubicInches(value float64) (*Volume, error) {
	return newVolume(value, VolumeCubicInch)
}

// FromImperialGallon creates a new Volume instance from ImperialGallon.
func (uf VolumeFactory) FromImperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialGallon)
}

// FromImperialOunce creates a new Volume instance from ImperialOunce.
func (uf VolumeFactory) FromImperialOunces(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialOunce)
}

// FromUsGallon creates a new Volume instance from UsGallon.
func (uf VolumeFactory) FromUsGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsGallon)
}

// FromUsOunce creates a new Volume instance from UsOunce.
func (uf VolumeFactory) FromUsOunces(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsOunce)
}

// FromUsTablespoon creates a new Volume instance from UsTablespoon.
func (uf VolumeFactory) FromUsTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsTablespoon)
}

// FromAuTablespoon creates a new Volume instance from AuTablespoon.
func (uf VolumeFactory) FromAuTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeAuTablespoon)
}

// FromUkTablespoon creates a new Volume instance from UkTablespoon.
func (uf VolumeFactory) FromUkTablespoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUkTablespoon)
}

// FromMetricTeaspoon creates a new Volume instance from MetricTeaspoon.
func (uf VolumeFactory) FromMetricTeaspoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMetricTeaspoon)
}

// FromUsTeaspoon creates a new Volume instance from UsTeaspoon.
func (uf VolumeFactory) FromUsTeaspoons(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsTeaspoon)
}

// FromMetricCup creates a new Volume instance from MetricCup.
func (uf VolumeFactory) FromMetricCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeMetricCup)
}

// FromUsCustomaryCup creates a new Volume instance from UsCustomaryCup.
func (uf VolumeFactory) FromUsCustomaryCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsCustomaryCup)
}

// FromUsLegalCup creates a new Volume instance from UsLegalCup.
func (uf VolumeFactory) FromUsLegalCups(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsLegalCup)
}

// FromOilBarrel creates a new Volume instance from OilBarrel.
func (uf VolumeFactory) FromOilBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeOilBarrel)
}

// FromUsBeerBarrel creates a new Volume instance from UsBeerBarrel.
func (uf VolumeFactory) FromUsBeerBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsBeerBarrel)
}

// FromImperialBeerBarrel creates a new Volume instance from ImperialBeerBarrel.
func (uf VolumeFactory) FromImperialBeerBarrels(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialBeerBarrel)
}

// FromUsQuart creates a new Volume instance from UsQuart.
func (uf VolumeFactory) FromUsQuarts(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsQuart)
}

// FromImperialQuart creates a new Volume instance from ImperialQuart.
func (uf VolumeFactory) FromImperialQuarts(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialQuart)
}

// FromUsPint creates a new Volume instance from UsPint.
func (uf VolumeFactory) FromUsPints(value float64) (*Volume, error) {
	return newVolume(value, VolumeUsPint)
}

// FromAcreFoot creates a new Volume instance from AcreFoot.
func (uf VolumeFactory) FromAcreFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeAcreFoot)
}

// FromImperialPint creates a new Volume instance from ImperialPint.
func (uf VolumeFactory) FromImperialPints(value float64) (*Volume, error) {
	return newVolume(value, VolumeImperialPint)
}

// FromBoardFoot creates a new Volume instance from BoardFoot.
func (uf VolumeFactory) FromBoardFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeBoardFoot)
}

// FromNanoliter creates a new Volume instance from Nanoliter.
func (uf VolumeFactory) FromNanoliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeNanoliter)
}

// FromMicroliter creates a new Volume instance from Microliter.
func (uf VolumeFactory) FromMicroliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMicroliter)
}

// FromMilliliter creates a new Volume instance from Milliliter.
func (uf VolumeFactory) FromMilliliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMilliliter)
}

// FromCentiliter creates a new Volume instance from Centiliter.
func (uf VolumeFactory) FromCentiliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeCentiliter)
}

// FromDeciliter creates a new Volume instance from Deciliter.
func (uf VolumeFactory) FromDeciliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeDeciliter)
}

// FromDecaliter creates a new Volume instance from Decaliter.
func (uf VolumeFactory) FromDecaliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeDecaliter)
}

// FromHectoliter creates a new Volume instance from Hectoliter.
func (uf VolumeFactory) FromHectoliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectoliter)
}

// FromKiloliter creates a new Volume instance from Kiloliter.
func (uf VolumeFactory) FromKiloliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeKiloliter)
}

// FromMegaliter creates a new Volume instance from Megaliter.
func (uf VolumeFactory) FromMegaliters(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegaliter)
}

// FromHectocubicMeter creates a new Volume instance from HectocubicMeter.
func (uf VolumeFactory) FromHectocubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectocubicMeter)
}

// FromKilocubicMeter creates a new Volume instance from KilocubicMeter.
func (uf VolumeFactory) FromKilocubicMeters(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilocubicMeter)
}

// FromHectocubicFoot creates a new Volume instance from HectocubicFoot.
func (uf VolumeFactory) FromHectocubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectocubicFoot)
}

// FromKilocubicFoot creates a new Volume instance from KilocubicFoot.
func (uf VolumeFactory) FromKilocubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilocubicFoot)
}

// FromMegacubicFoot creates a new Volume instance from MegacubicFoot.
func (uf VolumeFactory) FromMegacubicFeet(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegacubicFoot)
}

// FromKiloimperialGallon creates a new Volume instance from KiloimperialGallon.
func (uf VolumeFactory) FromKiloimperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeKiloimperialGallon)
}

// FromMegaimperialGallon creates a new Volume instance from MegaimperialGallon.
func (uf VolumeFactory) FromMegaimperialGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegaimperialGallon)
}

// FromDecausGallon creates a new Volume instance from DecausGallon.
func (uf VolumeFactory) FromDecausGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeDecausGallon)
}

// FromDeciusGallon creates a new Volume instance from DeciusGallon.
func (uf VolumeFactory) FromDeciusGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeDeciusGallon)
}

// FromHectousGallon creates a new Volume instance from HectousGallon.
func (uf VolumeFactory) FromHectousGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeHectousGallon)
}

// FromKilousGallon creates a new Volume instance from KilousGallon.
func (uf VolumeFactory) FromKilousGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeKilousGallon)
}

// FromMegausGallon creates a new Volume instance from MegausGallon.
func (uf VolumeFactory) FromMegausGallons(value float64) (*Volume, error) {
	return newVolume(value, VolumeMegausGallon)
}




// newVolume creates a new Volume.
func newVolume(value float64, fromUnit VolumeUnits) (*Volume, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Volume{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Volume in CubicMeter.
func (a *Volume) BaseValue() float64 {
	return a.value
}


// Liter returns the value in Liter.
func (a *Volume) Liters() float64 {
	if a.litersLazy != nil {
		return *a.litersLazy
	}
	liters := a.convertFromBase(VolumeLiter)
	a.litersLazy = &liters
	return liters
}

// CubicMeter returns the value in CubicMeter.
func (a *Volume) CubicMeters() float64 {
	if a.cubic_metersLazy != nil {
		return *a.cubic_metersLazy
	}
	cubic_meters := a.convertFromBase(VolumeCubicMeter)
	a.cubic_metersLazy = &cubic_meters
	return cubic_meters
}

// CubicKilometer returns the value in CubicKilometer.
func (a *Volume) CubicKilometers() float64 {
	if a.cubic_kilometersLazy != nil {
		return *a.cubic_kilometersLazy
	}
	cubic_kilometers := a.convertFromBase(VolumeCubicKilometer)
	a.cubic_kilometersLazy = &cubic_kilometers
	return cubic_kilometers
}

// CubicHectometer returns the value in CubicHectometer.
func (a *Volume) CubicHectometers() float64 {
	if a.cubic_hectometersLazy != nil {
		return *a.cubic_hectometersLazy
	}
	cubic_hectometers := a.convertFromBase(VolumeCubicHectometer)
	a.cubic_hectometersLazy = &cubic_hectometers
	return cubic_hectometers
}

// CubicDecimeter returns the value in CubicDecimeter.
func (a *Volume) CubicDecimeters() float64 {
	if a.cubic_decimetersLazy != nil {
		return *a.cubic_decimetersLazy
	}
	cubic_decimeters := a.convertFromBase(VolumeCubicDecimeter)
	a.cubic_decimetersLazy = &cubic_decimeters
	return cubic_decimeters
}

// CubicCentimeter returns the value in CubicCentimeter.
func (a *Volume) CubicCentimeters() float64 {
	if a.cubic_centimetersLazy != nil {
		return *a.cubic_centimetersLazy
	}
	cubic_centimeters := a.convertFromBase(VolumeCubicCentimeter)
	a.cubic_centimetersLazy = &cubic_centimeters
	return cubic_centimeters
}

// CubicMillimeter returns the value in CubicMillimeter.
func (a *Volume) CubicMillimeters() float64 {
	if a.cubic_millimetersLazy != nil {
		return *a.cubic_millimetersLazy
	}
	cubic_millimeters := a.convertFromBase(VolumeCubicMillimeter)
	a.cubic_millimetersLazy = &cubic_millimeters
	return cubic_millimeters
}

// CubicMicrometer returns the value in CubicMicrometer.
func (a *Volume) CubicMicrometers() float64 {
	if a.cubic_micrometersLazy != nil {
		return *a.cubic_micrometersLazy
	}
	cubic_micrometers := a.convertFromBase(VolumeCubicMicrometer)
	a.cubic_micrometersLazy = &cubic_micrometers
	return cubic_micrometers
}

// CubicMile returns the value in CubicMile.
func (a *Volume) CubicMiles() float64 {
	if a.cubic_milesLazy != nil {
		return *a.cubic_milesLazy
	}
	cubic_miles := a.convertFromBase(VolumeCubicMile)
	a.cubic_milesLazy = &cubic_miles
	return cubic_miles
}

// CubicYard returns the value in CubicYard.
func (a *Volume) CubicYards() float64 {
	if a.cubic_yardsLazy != nil {
		return *a.cubic_yardsLazy
	}
	cubic_yards := a.convertFromBase(VolumeCubicYard)
	a.cubic_yardsLazy = &cubic_yards
	return cubic_yards
}

// CubicFoot returns the value in CubicFoot.
func (a *Volume) CubicFeet() float64 {
	if a.cubic_feetLazy != nil {
		return *a.cubic_feetLazy
	}
	cubic_feet := a.convertFromBase(VolumeCubicFoot)
	a.cubic_feetLazy = &cubic_feet
	return cubic_feet
}

// CubicInch returns the value in CubicInch.
func (a *Volume) CubicInches() float64 {
	if a.cubic_inchesLazy != nil {
		return *a.cubic_inchesLazy
	}
	cubic_inches := a.convertFromBase(VolumeCubicInch)
	a.cubic_inchesLazy = &cubic_inches
	return cubic_inches
}

// ImperialGallon returns the value in ImperialGallon.
func (a *Volume) ImperialGallons() float64 {
	if a.imperial_gallonsLazy != nil {
		return *a.imperial_gallonsLazy
	}
	imperial_gallons := a.convertFromBase(VolumeImperialGallon)
	a.imperial_gallonsLazy = &imperial_gallons
	return imperial_gallons
}

// ImperialOunce returns the value in ImperialOunce.
func (a *Volume) ImperialOunces() float64 {
	if a.imperial_ouncesLazy != nil {
		return *a.imperial_ouncesLazy
	}
	imperial_ounces := a.convertFromBase(VolumeImperialOunce)
	a.imperial_ouncesLazy = &imperial_ounces
	return imperial_ounces
}

// UsGallon returns the value in UsGallon.
func (a *Volume) UsGallons() float64 {
	if a.us_gallonsLazy != nil {
		return *a.us_gallonsLazy
	}
	us_gallons := a.convertFromBase(VolumeUsGallon)
	a.us_gallonsLazy = &us_gallons
	return us_gallons
}

// UsOunce returns the value in UsOunce.
func (a *Volume) UsOunces() float64 {
	if a.us_ouncesLazy != nil {
		return *a.us_ouncesLazy
	}
	us_ounces := a.convertFromBase(VolumeUsOunce)
	a.us_ouncesLazy = &us_ounces
	return us_ounces
}

// UsTablespoon returns the value in UsTablespoon.
func (a *Volume) UsTablespoons() float64 {
	if a.us_tablespoonsLazy != nil {
		return *a.us_tablespoonsLazy
	}
	us_tablespoons := a.convertFromBase(VolumeUsTablespoon)
	a.us_tablespoonsLazy = &us_tablespoons
	return us_tablespoons
}

// AuTablespoon returns the value in AuTablespoon.
func (a *Volume) AuTablespoons() float64 {
	if a.au_tablespoonsLazy != nil {
		return *a.au_tablespoonsLazy
	}
	au_tablespoons := a.convertFromBase(VolumeAuTablespoon)
	a.au_tablespoonsLazy = &au_tablespoons
	return au_tablespoons
}

// UkTablespoon returns the value in UkTablespoon.
func (a *Volume) UkTablespoons() float64 {
	if a.uk_tablespoonsLazy != nil {
		return *a.uk_tablespoonsLazy
	}
	uk_tablespoons := a.convertFromBase(VolumeUkTablespoon)
	a.uk_tablespoonsLazy = &uk_tablespoons
	return uk_tablespoons
}

// MetricTeaspoon returns the value in MetricTeaspoon.
func (a *Volume) MetricTeaspoons() float64 {
	if a.metric_teaspoonsLazy != nil {
		return *a.metric_teaspoonsLazy
	}
	metric_teaspoons := a.convertFromBase(VolumeMetricTeaspoon)
	a.metric_teaspoonsLazy = &metric_teaspoons
	return metric_teaspoons
}

// UsTeaspoon returns the value in UsTeaspoon.
func (a *Volume) UsTeaspoons() float64 {
	if a.us_teaspoonsLazy != nil {
		return *a.us_teaspoonsLazy
	}
	us_teaspoons := a.convertFromBase(VolumeUsTeaspoon)
	a.us_teaspoonsLazy = &us_teaspoons
	return us_teaspoons
}

// MetricCup returns the value in MetricCup.
func (a *Volume) MetricCups() float64 {
	if a.metric_cupsLazy != nil {
		return *a.metric_cupsLazy
	}
	metric_cups := a.convertFromBase(VolumeMetricCup)
	a.metric_cupsLazy = &metric_cups
	return metric_cups
}

// UsCustomaryCup returns the value in UsCustomaryCup.
func (a *Volume) UsCustomaryCups() float64 {
	if a.us_customary_cupsLazy != nil {
		return *a.us_customary_cupsLazy
	}
	us_customary_cups := a.convertFromBase(VolumeUsCustomaryCup)
	a.us_customary_cupsLazy = &us_customary_cups
	return us_customary_cups
}

// UsLegalCup returns the value in UsLegalCup.
func (a *Volume) UsLegalCups() float64 {
	if a.us_legal_cupsLazy != nil {
		return *a.us_legal_cupsLazy
	}
	us_legal_cups := a.convertFromBase(VolumeUsLegalCup)
	a.us_legal_cupsLazy = &us_legal_cups
	return us_legal_cups
}

// OilBarrel returns the value in OilBarrel.
func (a *Volume) OilBarrels() float64 {
	if a.oil_barrelsLazy != nil {
		return *a.oil_barrelsLazy
	}
	oil_barrels := a.convertFromBase(VolumeOilBarrel)
	a.oil_barrelsLazy = &oil_barrels
	return oil_barrels
}

// UsBeerBarrel returns the value in UsBeerBarrel.
func (a *Volume) UsBeerBarrels() float64 {
	if a.us_beer_barrelsLazy != nil {
		return *a.us_beer_barrelsLazy
	}
	us_beer_barrels := a.convertFromBase(VolumeUsBeerBarrel)
	a.us_beer_barrelsLazy = &us_beer_barrels
	return us_beer_barrels
}

// ImperialBeerBarrel returns the value in ImperialBeerBarrel.
func (a *Volume) ImperialBeerBarrels() float64 {
	if a.imperial_beer_barrelsLazy != nil {
		return *a.imperial_beer_barrelsLazy
	}
	imperial_beer_barrels := a.convertFromBase(VolumeImperialBeerBarrel)
	a.imperial_beer_barrelsLazy = &imperial_beer_barrels
	return imperial_beer_barrels
}

// UsQuart returns the value in UsQuart.
func (a *Volume) UsQuarts() float64 {
	if a.us_quartsLazy != nil {
		return *a.us_quartsLazy
	}
	us_quarts := a.convertFromBase(VolumeUsQuart)
	a.us_quartsLazy = &us_quarts
	return us_quarts
}

// ImperialQuart returns the value in ImperialQuart.
func (a *Volume) ImperialQuarts() float64 {
	if a.imperial_quartsLazy != nil {
		return *a.imperial_quartsLazy
	}
	imperial_quarts := a.convertFromBase(VolumeImperialQuart)
	a.imperial_quartsLazy = &imperial_quarts
	return imperial_quarts
}

// UsPint returns the value in UsPint.
func (a *Volume) UsPints() float64 {
	if a.us_pintsLazy != nil {
		return *a.us_pintsLazy
	}
	us_pints := a.convertFromBase(VolumeUsPint)
	a.us_pintsLazy = &us_pints
	return us_pints
}

// AcreFoot returns the value in AcreFoot.
func (a *Volume) AcreFeet() float64 {
	if a.acre_feetLazy != nil {
		return *a.acre_feetLazy
	}
	acre_feet := a.convertFromBase(VolumeAcreFoot)
	a.acre_feetLazy = &acre_feet
	return acre_feet
}

// ImperialPint returns the value in ImperialPint.
func (a *Volume) ImperialPints() float64 {
	if a.imperial_pintsLazy != nil {
		return *a.imperial_pintsLazy
	}
	imperial_pints := a.convertFromBase(VolumeImperialPint)
	a.imperial_pintsLazy = &imperial_pints
	return imperial_pints
}

// BoardFoot returns the value in BoardFoot.
func (a *Volume) BoardFeet() float64 {
	if a.board_feetLazy != nil {
		return *a.board_feetLazy
	}
	board_feet := a.convertFromBase(VolumeBoardFoot)
	a.board_feetLazy = &board_feet
	return board_feet
}

// Nanoliter returns the value in Nanoliter.
func (a *Volume) Nanoliters() float64 {
	if a.nanolitersLazy != nil {
		return *a.nanolitersLazy
	}
	nanoliters := a.convertFromBase(VolumeNanoliter)
	a.nanolitersLazy = &nanoliters
	return nanoliters
}

// Microliter returns the value in Microliter.
func (a *Volume) Microliters() float64 {
	if a.microlitersLazy != nil {
		return *a.microlitersLazy
	}
	microliters := a.convertFromBase(VolumeMicroliter)
	a.microlitersLazy = &microliters
	return microliters
}

// Milliliter returns the value in Milliliter.
func (a *Volume) Milliliters() float64 {
	if a.millilitersLazy != nil {
		return *a.millilitersLazy
	}
	milliliters := a.convertFromBase(VolumeMilliliter)
	a.millilitersLazy = &milliliters
	return milliliters
}

// Centiliter returns the value in Centiliter.
func (a *Volume) Centiliters() float64 {
	if a.centilitersLazy != nil {
		return *a.centilitersLazy
	}
	centiliters := a.convertFromBase(VolumeCentiliter)
	a.centilitersLazy = &centiliters
	return centiliters
}

// Deciliter returns the value in Deciliter.
func (a *Volume) Deciliters() float64 {
	if a.decilitersLazy != nil {
		return *a.decilitersLazy
	}
	deciliters := a.convertFromBase(VolumeDeciliter)
	a.decilitersLazy = &deciliters
	return deciliters
}

// Decaliter returns the value in Decaliter.
func (a *Volume) Decaliters() float64 {
	if a.decalitersLazy != nil {
		return *a.decalitersLazy
	}
	decaliters := a.convertFromBase(VolumeDecaliter)
	a.decalitersLazy = &decaliters
	return decaliters
}

// Hectoliter returns the value in Hectoliter.
func (a *Volume) Hectoliters() float64 {
	if a.hectolitersLazy != nil {
		return *a.hectolitersLazy
	}
	hectoliters := a.convertFromBase(VolumeHectoliter)
	a.hectolitersLazy = &hectoliters
	return hectoliters
}

// Kiloliter returns the value in Kiloliter.
func (a *Volume) Kiloliters() float64 {
	if a.kilolitersLazy != nil {
		return *a.kilolitersLazy
	}
	kiloliters := a.convertFromBase(VolumeKiloliter)
	a.kilolitersLazy = &kiloliters
	return kiloliters
}

// Megaliter returns the value in Megaliter.
func (a *Volume) Megaliters() float64 {
	if a.megalitersLazy != nil {
		return *a.megalitersLazy
	}
	megaliters := a.convertFromBase(VolumeMegaliter)
	a.megalitersLazy = &megaliters
	return megaliters
}

// HectocubicMeter returns the value in HectocubicMeter.
func (a *Volume) HectocubicMeters() float64 {
	if a.hectocubic_metersLazy != nil {
		return *a.hectocubic_metersLazy
	}
	hectocubic_meters := a.convertFromBase(VolumeHectocubicMeter)
	a.hectocubic_metersLazy = &hectocubic_meters
	return hectocubic_meters
}

// KilocubicMeter returns the value in KilocubicMeter.
func (a *Volume) KilocubicMeters() float64 {
	if a.kilocubic_metersLazy != nil {
		return *a.kilocubic_metersLazy
	}
	kilocubic_meters := a.convertFromBase(VolumeKilocubicMeter)
	a.kilocubic_metersLazy = &kilocubic_meters
	return kilocubic_meters
}

// HectocubicFoot returns the value in HectocubicFoot.
func (a *Volume) HectocubicFeet() float64 {
	if a.hectocubic_feetLazy != nil {
		return *a.hectocubic_feetLazy
	}
	hectocubic_feet := a.convertFromBase(VolumeHectocubicFoot)
	a.hectocubic_feetLazy = &hectocubic_feet
	return hectocubic_feet
}

// KilocubicFoot returns the value in KilocubicFoot.
func (a *Volume) KilocubicFeet() float64 {
	if a.kilocubic_feetLazy != nil {
		return *a.kilocubic_feetLazy
	}
	kilocubic_feet := a.convertFromBase(VolumeKilocubicFoot)
	a.kilocubic_feetLazy = &kilocubic_feet
	return kilocubic_feet
}

// MegacubicFoot returns the value in MegacubicFoot.
func (a *Volume) MegacubicFeet() float64 {
	if a.megacubic_feetLazy != nil {
		return *a.megacubic_feetLazy
	}
	megacubic_feet := a.convertFromBase(VolumeMegacubicFoot)
	a.megacubic_feetLazy = &megacubic_feet
	return megacubic_feet
}

// KiloimperialGallon returns the value in KiloimperialGallon.
func (a *Volume) KiloimperialGallons() float64 {
	if a.kiloimperial_gallonsLazy != nil {
		return *a.kiloimperial_gallonsLazy
	}
	kiloimperial_gallons := a.convertFromBase(VolumeKiloimperialGallon)
	a.kiloimperial_gallonsLazy = &kiloimperial_gallons
	return kiloimperial_gallons
}

// MegaimperialGallon returns the value in MegaimperialGallon.
func (a *Volume) MegaimperialGallons() float64 {
	if a.megaimperial_gallonsLazy != nil {
		return *a.megaimperial_gallonsLazy
	}
	megaimperial_gallons := a.convertFromBase(VolumeMegaimperialGallon)
	a.megaimperial_gallonsLazy = &megaimperial_gallons
	return megaimperial_gallons
}

// DecausGallon returns the value in DecausGallon.
func (a *Volume) DecausGallons() float64 {
	if a.decaus_gallonsLazy != nil {
		return *a.decaus_gallonsLazy
	}
	decaus_gallons := a.convertFromBase(VolumeDecausGallon)
	a.decaus_gallonsLazy = &decaus_gallons
	return decaus_gallons
}

// DeciusGallon returns the value in DeciusGallon.
func (a *Volume) DeciusGallons() float64 {
	if a.decius_gallonsLazy != nil {
		return *a.decius_gallonsLazy
	}
	decius_gallons := a.convertFromBase(VolumeDeciusGallon)
	a.decius_gallonsLazy = &decius_gallons
	return decius_gallons
}

// HectousGallon returns the value in HectousGallon.
func (a *Volume) HectousGallons() float64 {
	if a.hectous_gallonsLazy != nil {
		return *a.hectous_gallonsLazy
	}
	hectous_gallons := a.convertFromBase(VolumeHectousGallon)
	a.hectous_gallonsLazy = &hectous_gallons
	return hectous_gallons
}

// KilousGallon returns the value in KilousGallon.
func (a *Volume) KilousGallons() float64 {
	if a.kilous_gallonsLazy != nil {
		return *a.kilous_gallonsLazy
	}
	kilous_gallons := a.convertFromBase(VolumeKilousGallon)
	a.kilous_gallonsLazy = &kilous_gallons
	return kilous_gallons
}

// MegausGallon returns the value in MegausGallon.
func (a *Volume) MegausGallons() float64 {
	if a.megaus_gallonsLazy != nil {
		return *a.megaus_gallonsLazy
	}
	megaus_gallons := a.convertFromBase(VolumeMegausGallon)
	a.megaus_gallonsLazy = &megaus_gallons
	return megaus_gallons
}


// ToDto creates an VolumeDto representation.
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

// ToDtoJSON creates an VolumeDto representation.
func (a *Volume) ToDtoJSON(holdInUnit *VolumeUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Volume to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Volume) String() string {
	return a.ToString(VolumeCubicMeter, 2)
}

// ToString formats the Volume to string.
// fractionalDigits -1 for not mention
func (a *Volume) ToString(unit VolumeUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Volume) getUnitAbbreviation(unit VolumeUnits) string {
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

// Check if the given Volume are equals to the current Volume
func (a *Volume) Equals(other *Volume) bool {
	return a.value == other.BaseValue()
}

// Check if the given Volume are equals to the current Volume
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

// Add the given Volume to the current Volume.
func (a *Volume) Add(other *Volume) *Volume {
	return &Volume{value: a.value + other.BaseValue()}
}

// Subtract the given Volume to the current Volume.
func (a *Volume) Subtract(other *Volume) *Volume {
	return &Volume{value: a.value - other.BaseValue()}
}

// Multiply the given Volume to the current Volume.
func (a *Volume) Multiply(other *Volume) *Volume {
	return &Volume{value: a.value * other.BaseValue()}
}

// Divide the given Volume to the current Volume.
func (a *Volume) Divide(other *Volume) *Volume {
	return &Volume{value: a.value / other.BaseValue()}
}