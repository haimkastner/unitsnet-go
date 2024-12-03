// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PressureChangeRateUnits enumeration
type PressureChangeRateUnits string

const (
    
        // 
        PressureChangeRatePascalPerSecond PressureChangeRateUnits = "PascalPerSecond"
        // 
        PressureChangeRatePascalPerMinute PressureChangeRateUnits = "PascalPerMinute"
        // 
        PressureChangeRateMillimeterOfMercuryPerSecond PressureChangeRateUnits = "MillimeterOfMercuryPerSecond"
        // 
        PressureChangeRateAtmospherePerSecond PressureChangeRateUnits = "AtmospherePerSecond"
        // 
        PressureChangeRatePoundForcePerSquareInchPerSecond PressureChangeRateUnits = "PoundForcePerSquareInchPerSecond"
        // 
        PressureChangeRatePoundForcePerSquareInchPerMinute PressureChangeRateUnits = "PoundForcePerSquareInchPerMinute"
        // 
        PressureChangeRateBarPerSecond PressureChangeRateUnits = "BarPerSecond"
        // 
        PressureChangeRateBarPerMinute PressureChangeRateUnits = "BarPerMinute"
        // 
        PressureChangeRateKilopascalPerSecond PressureChangeRateUnits = "KilopascalPerSecond"
        // 
        PressureChangeRateMegapascalPerSecond PressureChangeRateUnits = "MegapascalPerSecond"
        // 
        PressureChangeRateKilopascalPerMinute PressureChangeRateUnits = "KilopascalPerMinute"
        // 
        PressureChangeRateMegapascalPerMinute PressureChangeRateUnits = "MegapascalPerMinute"
        // 
        PressureChangeRateKilopoundForcePerSquareInchPerSecond PressureChangeRateUnits = "KilopoundForcePerSquareInchPerSecond"
        // 
        PressureChangeRateMegapoundForcePerSquareInchPerSecond PressureChangeRateUnits = "MegapoundForcePerSquareInchPerSecond"
        // 
        PressureChangeRateKilopoundForcePerSquareInchPerMinute PressureChangeRateUnits = "KilopoundForcePerSquareInchPerMinute"
        // 
        PressureChangeRateMegapoundForcePerSquareInchPerMinute PressureChangeRateUnits = "MegapoundForcePerSquareInchPerMinute"
        // 
        PressureChangeRateMillibarPerSecond PressureChangeRateUnits = "MillibarPerSecond"
        // 
        PressureChangeRateMillibarPerMinute PressureChangeRateUnits = "MillibarPerMinute"
)

// PressureChangeRateDto represents an PressureChangeRate
type PressureChangeRateDto struct {
	Value float64
	Unit  PressureChangeRateUnits
}

// PressureChangeRateDtoFactory struct to group related functions
type PressureChangeRateDtoFactory struct{}

func (udf PressureChangeRateDtoFactory) FromJSON(data []byte) (*PressureChangeRateDto, error) {
	a := PressureChangeRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PressureChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// PressureChangeRate struct
type PressureChangeRate struct {
	value       float64
    
    pascals_per_secondLazy *float64 
    pascals_per_minuteLazy *float64 
    millimeters_of_mercury_per_secondLazy *float64 
    atmospheres_per_secondLazy *float64 
    pounds_force_per_square_inch_per_secondLazy *float64 
    pounds_force_per_square_inch_per_minuteLazy *float64 
    bars_per_secondLazy *float64 
    bars_per_minuteLazy *float64 
    kilopascals_per_secondLazy *float64 
    megapascals_per_secondLazy *float64 
    kilopascals_per_minuteLazy *float64 
    megapascals_per_minuteLazy *float64 
    kilopounds_force_per_square_inch_per_secondLazy *float64 
    megapounds_force_per_square_inch_per_secondLazy *float64 
    kilopounds_force_per_square_inch_per_minuteLazy *float64 
    megapounds_force_per_square_inch_per_minuteLazy *float64 
    millibars_per_secondLazy *float64 
    millibars_per_minuteLazy *float64 
}

// PressureChangeRateFactory struct to group related functions
type PressureChangeRateFactory struct{}

func (uf PressureChangeRateFactory) CreatePressureChangeRate(value float64, unit PressureChangeRateUnits) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, unit)
}

func (uf PressureChangeRateFactory) FromDto(dto PressureChangeRateDto) (*PressureChangeRate, error) {
	return newPressureChangeRate(dto.Value, dto.Unit)
}

func (uf PressureChangeRateFactory) FromDtoJSON(data []byte) (*PressureChangeRate, error) {
	unitDto, err := PressureChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PressureChangeRateFactory{}.FromDto(*unitDto)
}


// FromPascalPerSecond creates a new PressureChangeRate instance from PascalPerSecond.
func (uf PressureChangeRateFactory) FromPascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePascalPerSecond)
}

// FromPascalPerMinute creates a new PressureChangeRate instance from PascalPerMinute.
func (uf PressureChangeRateFactory) FromPascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePascalPerMinute)
}

// FromMillimeterOfMercuryPerSecond creates a new PressureChangeRate instance from MillimeterOfMercuryPerSecond.
func (uf PressureChangeRateFactory) FromMillimetersOfMercuryPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMillimeterOfMercuryPerSecond)
}

// FromAtmospherePerSecond creates a new PressureChangeRate instance from AtmospherePerSecond.
func (uf PressureChangeRateFactory) FromAtmospheresPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateAtmospherePerSecond)
}

// FromPoundForcePerSquareInchPerSecond creates a new PressureChangeRate instance from PoundForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromPoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePoundForcePerSquareInchPerSecond)
}

// FromPoundForcePerSquareInchPerMinute creates a new PressureChangeRate instance from PoundForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromPoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePoundForcePerSquareInchPerMinute)
}

// FromBarPerSecond creates a new PressureChangeRate instance from BarPerSecond.
func (uf PressureChangeRateFactory) FromBarsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateBarPerSecond)
}

// FromBarPerMinute creates a new PressureChangeRate instance from BarPerMinute.
func (uf PressureChangeRateFactory) FromBarsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateBarPerMinute)
}

// FromKilopascalPerSecond creates a new PressureChangeRate instance from KilopascalPerSecond.
func (uf PressureChangeRateFactory) FromKilopascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopascalPerSecond)
}

// FromMegapascalPerSecond creates a new PressureChangeRate instance from MegapascalPerSecond.
func (uf PressureChangeRateFactory) FromMegapascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapascalPerSecond)
}

// FromKilopascalPerMinute creates a new PressureChangeRate instance from KilopascalPerMinute.
func (uf PressureChangeRateFactory) FromKilopascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopascalPerMinute)
}

// FromMegapascalPerMinute creates a new PressureChangeRate instance from MegapascalPerMinute.
func (uf PressureChangeRateFactory) FromMegapascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapascalPerMinute)
}

// FromKilopoundForcePerSquareInchPerSecond creates a new PressureChangeRate instance from KilopoundForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromKilopoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopoundForcePerSquareInchPerSecond)
}

// FromMegapoundForcePerSquareInchPerSecond creates a new PressureChangeRate instance from MegapoundForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromMegapoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapoundForcePerSquareInchPerSecond)
}

// FromKilopoundForcePerSquareInchPerMinute creates a new PressureChangeRate instance from KilopoundForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromKilopoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopoundForcePerSquareInchPerMinute)
}

// FromMegapoundForcePerSquareInchPerMinute creates a new PressureChangeRate instance from MegapoundForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromMegapoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapoundForcePerSquareInchPerMinute)
}

// FromMillibarPerSecond creates a new PressureChangeRate instance from MillibarPerSecond.
func (uf PressureChangeRateFactory) FromMillibarsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMillibarPerSecond)
}

// FromMillibarPerMinute creates a new PressureChangeRate instance from MillibarPerMinute.
func (uf PressureChangeRateFactory) FromMillibarsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMillibarPerMinute)
}




// newPressureChangeRate creates a new PressureChangeRate.
func newPressureChangeRate(value float64, fromUnit PressureChangeRateUnits) (*PressureChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &PressureChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PressureChangeRate in PascalPerSecond.
func (a *PressureChangeRate) BaseValue() float64 {
	return a.value
}


// PascalPerSecond returns the value in PascalPerSecond.
func (a *PressureChangeRate) PascalsPerSecond() float64 {
	if a.pascals_per_secondLazy != nil {
		return *a.pascals_per_secondLazy
	}
	pascals_per_second := a.convertFromBase(PressureChangeRatePascalPerSecond)
	a.pascals_per_secondLazy = &pascals_per_second
	return pascals_per_second
}

// PascalPerMinute returns the value in PascalPerMinute.
func (a *PressureChangeRate) PascalsPerMinute() float64 {
	if a.pascals_per_minuteLazy != nil {
		return *a.pascals_per_minuteLazy
	}
	pascals_per_minute := a.convertFromBase(PressureChangeRatePascalPerMinute)
	a.pascals_per_minuteLazy = &pascals_per_minute
	return pascals_per_minute
}

// MillimeterOfMercuryPerSecond returns the value in MillimeterOfMercuryPerSecond.
func (a *PressureChangeRate) MillimetersOfMercuryPerSecond() float64 {
	if a.millimeters_of_mercury_per_secondLazy != nil {
		return *a.millimeters_of_mercury_per_secondLazy
	}
	millimeters_of_mercury_per_second := a.convertFromBase(PressureChangeRateMillimeterOfMercuryPerSecond)
	a.millimeters_of_mercury_per_secondLazy = &millimeters_of_mercury_per_second
	return millimeters_of_mercury_per_second
}

// AtmospherePerSecond returns the value in AtmospherePerSecond.
func (a *PressureChangeRate) AtmospheresPerSecond() float64 {
	if a.atmospheres_per_secondLazy != nil {
		return *a.atmospheres_per_secondLazy
	}
	atmospheres_per_second := a.convertFromBase(PressureChangeRateAtmospherePerSecond)
	a.atmospheres_per_secondLazy = &atmospheres_per_second
	return atmospheres_per_second
}

// PoundForcePerSquareInchPerSecond returns the value in PoundForcePerSquareInchPerSecond.
func (a *PressureChangeRate) PoundsForcePerSquareInchPerSecond() float64 {
	if a.pounds_force_per_square_inch_per_secondLazy != nil {
		return *a.pounds_force_per_square_inch_per_secondLazy
	}
	pounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRatePoundForcePerSquareInchPerSecond)
	a.pounds_force_per_square_inch_per_secondLazy = &pounds_force_per_square_inch_per_second
	return pounds_force_per_square_inch_per_second
}

// PoundForcePerSquareInchPerMinute returns the value in PoundForcePerSquareInchPerMinute.
func (a *PressureChangeRate) PoundsForcePerSquareInchPerMinute() float64 {
	if a.pounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.pounds_force_per_square_inch_per_minuteLazy
	}
	pounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRatePoundForcePerSquareInchPerMinute)
	a.pounds_force_per_square_inch_per_minuteLazy = &pounds_force_per_square_inch_per_minute
	return pounds_force_per_square_inch_per_minute
}

// BarPerSecond returns the value in BarPerSecond.
func (a *PressureChangeRate) BarsPerSecond() float64 {
	if a.bars_per_secondLazy != nil {
		return *a.bars_per_secondLazy
	}
	bars_per_second := a.convertFromBase(PressureChangeRateBarPerSecond)
	a.bars_per_secondLazy = &bars_per_second
	return bars_per_second
}

// BarPerMinute returns the value in BarPerMinute.
func (a *PressureChangeRate) BarsPerMinute() float64 {
	if a.bars_per_minuteLazy != nil {
		return *a.bars_per_minuteLazy
	}
	bars_per_minute := a.convertFromBase(PressureChangeRateBarPerMinute)
	a.bars_per_minuteLazy = &bars_per_minute
	return bars_per_minute
}

// KilopascalPerSecond returns the value in KilopascalPerSecond.
func (a *PressureChangeRate) KilopascalsPerSecond() float64 {
	if a.kilopascals_per_secondLazy != nil {
		return *a.kilopascals_per_secondLazy
	}
	kilopascals_per_second := a.convertFromBase(PressureChangeRateKilopascalPerSecond)
	a.kilopascals_per_secondLazy = &kilopascals_per_second
	return kilopascals_per_second
}

// MegapascalPerSecond returns the value in MegapascalPerSecond.
func (a *PressureChangeRate) MegapascalsPerSecond() float64 {
	if a.megapascals_per_secondLazy != nil {
		return *a.megapascals_per_secondLazy
	}
	megapascals_per_second := a.convertFromBase(PressureChangeRateMegapascalPerSecond)
	a.megapascals_per_secondLazy = &megapascals_per_second
	return megapascals_per_second
}

// KilopascalPerMinute returns the value in KilopascalPerMinute.
func (a *PressureChangeRate) KilopascalsPerMinute() float64 {
	if a.kilopascals_per_minuteLazy != nil {
		return *a.kilopascals_per_minuteLazy
	}
	kilopascals_per_minute := a.convertFromBase(PressureChangeRateKilopascalPerMinute)
	a.kilopascals_per_minuteLazy = &kilopascals_per_minute
	return kilopascals_per_minute
}

// MegapascalPerMinute returns the value in MegapascalPerMinute.
func (a *PressureChangeRate) MegapascalsPerMinute() float64 {
	if a.megapascals_per_minuteLazy != nil {
		return *a.megapascals_per_minuteLazy
	}
	megapascals_per_minute := a.convertFromBase(PressureChangeRateMegapascalPerMinute)
	a.megapascals_per_minuteLazy = &megapascals_per_minute
	return megapascals_per_minute
}

// KilopoundForcePerSquareInchPerSecond returns the value in KilopoundForcePerSquareInchPerSecond.
func (a *PressureChangeRate) KilopoundsForcePerSquareInchPerSecond() float64 {
	if a.kilopounds_force_per_square_inch_per_secondLazy != nil {
		return *a.kilopounds_force_per_square_inch_per_secondLazy
	}
	kilopounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRateKilopoundForcePerSquareInchPerSecond)
	a.kilopounds_force_per_square_inch_per_secondLazy = &kilopounds_force_per_square_inch_per_second
	return kilopounds_force_per_square_inch_per_second
}

// MegapoundForcePerSquareInchPerSecond returns the value in MegapoundForcePerSquareInchPerSecond.
func (a *PressureChangeRate) MegapoundsForcePerSquareInchPerSecond() float64 {
	if a.megapounds_force_per_square_inch_per_secondLazy != nil {
		return *a.megapounds_force_per_square_inch_per_secondLazy
	}
	megapounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRateMegapoundForcePerSquareInchPerSecond)
	a.megapounds_force_per_square_inch_per_secondLazy = &megapounds_force_per_square_inch_per_second
	return megapounds_force_per_square_inch_per_second
}

// KilopoundForcePerSquareInchPerMinute returns the value in KilopoundForcePerSquareInchPerMinute.
func (a *PressureChangeRate) KilopoundsForcePerSquareInchPerMinute() float64 {
	if a.kilopounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.kilopounds_force_per_square_inch_per_minuteLazy
	}
	kilopounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRateKilopoundForcePerSquareInchPerMinute)
	a.kilopounds_force_per_square_inch_per_minuteLazy = &kilopounds_force_per_square_inch_per_minute
	return kilopounds_force_per_square_inch_per_minute
}

// MegapoundForcePerSquareInchPerMinute returns the value in MegapoundForcePerSquareInchPerMinute.
func (a *PressureChangeRate) MegapoundsForcePerSquareInchPerMinute() float64 {
	if a.megapounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.megapounds_force_per_square_inch_per_minuteLazy
	}
	megapounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRateMegapoundForcePerSquareInchPerMinute)
	a.megapounds_force_per_square_inch_per_minuteLazy = &megapounds_force_per_square_inch_per_minute
	return megapounds_force_per_square_inch_per_minute
}

// MillibarPerSecond returns the value in MillibarPerSecond.
func (a *PressureChangeRate) MillibarsPerSecond() float64 {
	if a.millibars_per_secondLazy != nil {
		return *a.millibars_per_secondLazy
	}
	millibars_per_second := a.convertFromBase(PressureChangeRateMillibarPerSecond)
	a.millibars_per_secondLazy = &millibars_per_second
	return millibars_per_second
}

// MillibarPerMinute returns the value in MillibarPerMinute.
func (a *PressureChangeRate) MillibarsPerMinute() float64 {
	if a.millibars_per_minuteLazy != nil {
		return *a.millibars_per_minuteLazy
	}
	millibars_per_minute := a.convertFromBase(PressureChangeRateMillibarPerMinute)
	a.millibars_per_minuteLazy = &millibars_per_minute
	return millibars_per_minute
}


// ToDto creates an PressureChangeRateDto representation.
func (a *PressureChangeRate) ToDto(holdInUnit *PressureChangeRateUnits) PressureChangeRateDto {
	if holdInUnit == nil {
		defaultUnit := PressureChangeRatePascalPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return PressureChangeRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PressureChangeRateDto representation.
func (a *PressureChangeRate) ToDtoJSON(holdInUnit *PressureChangeRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts PressureChangeRate to a specific unit value.
func (a *PressureChangeRate) Convert(toUnit PressureChangeRateUnits) float64 {
	switch toUnit { 
    case PressureChangeRatePascalPerSecond:
		return a.PascalsPerSecond()
    case PressureChangeRatePascalPerMinute:
		return a.PascalsPerMinute()
    case PressureChangeRateMillimeterOfMercuryPerSecond:
		return a.MillimetersOfMercuryPerSecond()
    case PressureChangeRateAtmospherePerSecond:
		return a.AtmospheresPerSecond()
    case PressureChangeRatePoundForcePerSquareInchPerSecond:
		return a.PoundsForcePerSquareInchPerSecond()
    case PressureChangeRatePoundForcePerSquareInchPerMinute:
		return a.PoundsForcePerSquareInchPerMinute()
    case PressureChangeRateBarPerSecond:
		return a.BarsPerSecond()
    case PressureChangeRateBarPerMinute:
		return a.BarsPerMinute()
    case PressureChangeRateKilopascalPerSecond:
		return a.KilopascalsPerSecond()
    case PressureChangeRateMegapascalPerSecond:
		return a.MegapascalsPerSecond()
    case PressureChangeRateKilopascalPerMinute:
		return a.KilopascalsPerMinute()
    case PressureChangeRateMegapascalPerMinute:
		return a.MegapascalsPerMinute()
    case PressureChangeRateKilopoundForcePerSquareInchPerSecond:
		return a.KilopoundsForcePerSquareInchPerSecond()
    case PressureChangeRateMegapoundForcePerSquareInchPerSecond:
		return a.MegapoundsForcePerSquareInchPerSecond()
    case PressureChangeRateKilopoundForcePerSquareInchPerMinute:
		return a.KilopoundsForcePerSquareInchPerMinute()
    case PressureChangeRateMegapoundForcePerSquareInchPerMinute:
		return a.MegapoundsForcePerSquareInchPerMinute()
    case PressureChangeRateMillibarPerSecond:
		return a.MillibarsPerSecond()
    case PressureChangeRateMillibarPerMinute:
		return a.MillibarsPerMinute()
	default:
		return 0
	}
}

func (a *PressureChangeRate) convertFromBase(toUnit PressureChangeRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case PressureChangeRatePascalPerSecond:
		return (value) 
	case PressureChangeRatePascalPerMinute:
		return (value * 60) 
	case PressureChangeRateMillimeterOfMercuryPerSecond:
		return (value / 133.322) 
	case PressureChangeRateAtmospherePerSecond:
		return (value / (1.01325 * 1e5)) 
	case PressureChangeRatePoundForcePerSquareInchPerSecond:
		return (value / 6.894757293168361e3) 
	case PressureChangeRatePoundForcePerSquareInchPerMinute:
		return (value / 6.894757293168361e3 * 60) 
	case PressureChangeRateBarPerSecond:
		return (value / 1e5) 
	case PressureChangeRateBarPerMinute:
		return (value / 1e5 * 60) 
	case PressureChangeRateKilopascalPerSecond:
		return ((value) / 1000.0) 
	case PressureChangeRateMegapascalPerSecond:
		return ((value) / 1000000.0) 
	case PressureChangeRateKilopascalPerMinute:
		return ((value * 60) / 1000.0) 
	case PressureChangeRateMegapascalPerMinute:
		return ((value * 60) / 1000000.0) 
	case PressureChangeRateKilopoundForcePerSquareInchPerSecond:
		return ((value / 6.894757293168361e3) / 1000.0) 
	case PressureChangeRateMegapoundForcePerSquareInchPerSecond:
		return ((value / 6.894757293168361e3) / 1000000.0) 
	case PressureChangeRateKilopoundForcePerSquareInchPerMinute:
		return ((value / 6.894757293168361e3 * 60) / 1000.0) 
	case PressureChangeRateMegapoundForcePerSquareInchPerMinute:
		return ((value / 6.894757293168361e3 * 60) / 1000000.0) 
	case PressureChangeRateMillibarPerSecond:
		return ((value / 1e5) / 0.001) 
	case PressureChangeRateMillibarPerMinute:
		return ((value / 1e5 * 60) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *PressureChangeRate) convertToBase(value float64, fromUnit PressureChangeRateUnits) float64 {
	switch fromUnit { 
	case PressureChangeRatePascalPerSecond:
		return (value) 
	case PressureChangeRatePascalPerMinute:
		return (value / 60) 
	case PressureChangeRateMillimeterOfMercuryPerSecond:
		return (value * 133.322) 
	case PressureChangeRateAtmospherePerSecond:
		return (value * 1.01325 * 1e5) 
	case PressureChangeRatePoundForcePerSquareInchPerSecond:
		return (value * 6.894757293168361e3) 
	case PressureChangeRatePoundForcePerSquareInchPerMinute:
		return (value * 6.894757293168361e3 / 60) 
	case PressureChangeRateBarPerSecond:
		return (value * 1e5) 
	case PressureChangeRateBarPerMinute:
		return (value * 1e5 / 60) 
	case PressureChangeRateKilopascalPerSecond:
		return ((value) * 1000.0) 
	case PressureChangeRateMegapascalPerSecond:
		return ((value) * 1000000.0) 
	case PressureChangeRateKilopascalPerMinute:
		return ((value / 60) * 1000.0) 
	case PressureChangeRateMegapascalPerMinute:
		return ((value / 60) * 1000000.0) 
	case PressureChangeRateKilopoundForcePerSquareInchPerSecond:
		return ((value * 6.894757293168361e3) * 1000.0) 
	case PressureChangeRateMegapoundForcePerSquareInchPerSecond:
		return ((value * 6.894757293168361e3) * 1000000.0) 
	case PressureChangeRateKilopoundForcePerSquareInchPerMinute:
		return ((value * 6.894757293168361e3 / 60) * 1000.0) 
	case PressureChangeRateMegapoundForcePerSquareInchPerMinute:
		return ((value * 6.894757293168361e3 / 60) * 1000000.0) 
	case PressureChangeRateMillibarPerSecond:
		return ((value * 1e5) * 0.001) 
	case PressureChangeRateMillibarPerMinute:
		return ((value * 1e5 / 60) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a PressureChangeRate) String() string {
	return a.ToString(PressureChangeRatePascalPerSecond, 2)
}

// ToString formats the PressureChangeRate to string.
// fractionalDigits -1 for not mention
func (a *PressureChangeRate) ToString(unit PressureChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *PressureChangeRate) getUnitAbbreviation(unit PressureChangeRateUnits) string {
	switch unit { 
	case PressureChangeRatePascalPerSecond:
		return "Pa/s" 
	case PressureChangeRatePascalPerMinute:
		return "Pa/min" 
	case PressureChangeRateMillimeterOfMercuryPerSecond:
		return "mmHg/s" 
	case PressureChangeRateAtmospherePerSecond:
		return "atm/s" 
	case PressureChangeRatePoundForcePerSquareInchPerSecond:
		return "psi/s" 
	case PressureChangeRatePoundForcePerSquareInchPerMinute:
		return "psi/min" 
	case PressureChangeRateBarPerSecond:
		return "bar/s" 
	case PressureChangeRateBarPerMinute:
		return "bar/min" 
	case PressureChangeRateKilopascalPerSecond:
		return "kPa/s" 
	case PressureChangeRateMegapascalPerSecond:
		return "MPa/s" 
	case PressureChangeRateKilopascalPerMinute:
		return "kPa/min" 
	case PressureChangeRateMegapascalPerMinute:
		return "MPa/min" 
	case PressureChangeRateKilopoundForcePerSquareInchPerSecond:
		return "kpsi/s" 
	case PressureChangeRateMegapoundForcePerSquareInchPerSecond:
		return "Mpsi/s" 
	case PressureChangeRateKilopoundForcePerSquareInchPerMinute:
		return "kpsi/min" 
	case PressureChangeRateMegapoundForcePerSquareInchPerMinute:
		return "Mpsi/min" 
	case PressureChangeRateMillibarPerSecond:
		return "mbar/s" 
	case PressureChangeRateMillibarPerMinute:
		return "mbar/min" 
	default:
		return ""
	}
}

// Check if the given PressureChangeRate are equals to the current PressureChangeRate
func (a *PressureChangeRate) Equals(other *PressureChangeRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given PressureChangeRate are equals to the current PressureChangeRate
func (a *PressureChangeRate) CompareTo(other *PressureChangeRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given PressureChangeRate to the current PressureChangeRate.
func (a *PressureChangeRate) Add(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value + other.BaseValue()}
}

// Subtract the given PressureChangeRate to the current PressureChangeRate.
func (a *PressureChangeRate) Subtract(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value - other.BaseValue()}
}

// Multiply the given PressureChangeRate to the current PressureChangeRate.
func (a *PressureChangeRate) Multiply(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value * other.BaseValue()}
}

// Divide the given PressureChangeRate to the current PressureChangeRate.
func (a *PressureChangeRate) Divide(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value / other.BaseValue()}
}