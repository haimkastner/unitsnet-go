// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PressureChangeRateUnits defines various units of PressureChangeRate.
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

// PressureChangeRateDto represents a PressureChangeRate measurement with a numerical value and its corresponding unit.
type PressureChangeRateDto struct {
    // Value is the numerical representation of the PressureChangeRate.
	Value float64
    // Unit specifies the unit of measurement for the PressureChangeRate, as defined in the PressureChangeRateUnits enumeration.
	Unit  PressureChangeRateUnits
}

// PressureChangeRateDtoFactory groups methods for creating and serializing PressureChangeRateDto objects.
type PressureChangeRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PressureChangeRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PressureChangeRateDtoFactory) FromJSON(data []byte) (*PressureChangeRateDto, error) {
	a := PressureChangeRateDto{}

    // Parse JSON into PressureChangeRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a PressureChangeRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PressureChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// PressureChangeRate represents a measurement in a of PressureChangeRate.
//
// Pressure change rate is the ratio of the pressure change to the time during which the change occurred (value of pressure changes per unit time).
type PressureChangeRate struct {
	// value is the base measurement stored internally.
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

// PressureChangeRateFactory groups methods for creating PressureChangeRate instances.
type PressureChangeRateFactory struct{}

// CreatePressureChangeRate creates a new PressureChangeRate instance from the given value and unit.
func (uf PressureChangeRateFactory) CreatePressureChangeRate(value float64, unit PressureChangeRateUnits) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, unit)
}

// FromDto converts a PressureChangeRateDto to a PressureChangeRate instance.
func (uf PressureChangeRateFactory) FromDto(dto PressureChangeRateDto) (*PressureChangeRate, error) {
	return newPressureChangeRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a PressureChangeRate instance.
func (uf PressureChangeRateFactory) FromDtoJSON(data []byte) (*PressureChangeRate, error) {
	unitDto, err := PressureChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PressureChangeRateDto from JSON: %w", err)
	}
	return PressureChangeRateFactory{}.FromDto(*unitDto)
}


// FromPascalsPerSecond creates a new PressureChangeRate instance from a value in PascalsPerSecond.
func (uf PressureChangeRateFactory) FromPascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePascalPerSecond)
}

// FromPascalsPerMinute creates a new PressureChangeRate instance from a value in PascalsPerMinute.
func (uf PressureChangeRateFactory) FromPascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePascalPerMinute)
}

// FromMillimetersOfMercuryPerSecond creates a new PressureChangeRate instance from a value in MillimetersOfMercuryPerSecond.
func (uf PressureChangeRateFactory) FromMillimetersOfMercuryPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMillimeterOfMercuryPerSecond)
}

// FromAtmospheresPerSecond creates a new PressureChangeRate instance from a value in AtmospheresPerSecond.
func (uf PressureChangeRateFactory) FromAtmospheresPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateAtmospherePerSecond)
}

// FromPoundsForcePerSquareInchPerSecond creates a new PressureChangeRate instance from a value in PoundsForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromPoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePoundForcePerSquareInchPerSecond)
}

// FromPoundsForcePerSquareInchPerMinute creates a new PressureChangeRate instance from a value in PoundsForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromPoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRatePoundForcePerSquareInchPerMinute)
}

// FromBarsPerSecond creates a new PressureChangeRate instance from a value in BarsPerSecond.
func (uf PressureChangeRateFactory) FromBarsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateBarPerSecond)
}

// FromBarsPerMinute creates a new PressureChangeRate instance from a value in BarsPerMinute.
func (uf PressureChangeRateFactory) FromBarsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateBarPerMinute)
}

// FromKilopascalsPerSecond creates a new PressureChangeRate instance from a value in KilopascalsPerSecond.
func (uf PressureChangeRateFactory) FromKilopascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopascalPerSecond)
}

// FromMegapascalsPerSecond creates a new PressureChangeRate instance from a value in MegapascalsPerSecond.
func (uf PressureChangeRateFactory) FromMegapascalsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapascalPerSecond)
}

// FromKilopascalsPerMinute creates a new PressureChangeRate instance from a value in KilopascalsPerMinute.
func (uf PressureChangeRateFactory) FromKilopascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopascalPerMinute)
}

// FromMegapascalsPerMinute creates a new PressureChangeRate instance from a value in MegapascalsPerMinute.
func (uf PressureChangeRateFactory) FromMegapascalsPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapascalPerMinute)
}

// FromKilopoundsForcePerSquareInchPerSecond creates a new PressureChangeRate instance from a value in KilopoundsForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromKilopoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopoundForcePerSquareInchPerSecond)
}

// FromMegapoundsForcePerSquareInchPerSecond creates a new PressureChangeRate instance from a value in MegapoundsForcePerSquareInchPerSecond.
func (uf PressureChangeRateFactory) FromMegapoundsForcePerSquareInchPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapoundForcePerSquareInchPerSecond)
}

// FromKilopoundsForcePerSquareInchPerMinute creates a new PressureChangeRate instance from a value in KilopoundsForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromKilopoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateKilopoundForcePerSquareInchPerMinute)
}

// FromMegapoundsForcePerSquareInchPerMinute creates a new PressureChangeRate instance from a value in MegapoundsForcePerSquareInchPerMinute.
func (uf PressureChangeRateFactory) FromMegapoundsForcePerSquareInchPerMinute(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMegapoundForcePerSquareInchPerMinute)
}

// FromMillibarsPerSecond creates a new PressureChangeRate instance from a value in MillibarsPerSecond.
func (uf PressureChangeRateFactory) FromMillibarsPerSecond(value float64) (*PressureChangeRate, error) {
	return newPressureChangeRate(value, PressureChangeRateMillibarPerSecond)
}

// FromMillibarsPerMinute creates a new PressureChangeRate instance from a value in MillibarsPerMinute.
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

// BaseValue returns the base value of PressureChangeRate in PascalPerSecond unit (the base/default quantity).
func (a *PressureChangeRate) BaseValue() float64 {
	return a.value
}


// PascalsPerSecond returns the PressureChangeRate value in PascalsPerSecond.
//
// 
func (a *PressureChangeRate) PascalsPerSecond() float64 {
	if a.pascals_per_secondLazy != nil {
		return *a.pascals_per_secondLazy
	}
	pascals_per_second := a.convertFromBase(PressureChangeRatePascalPerSecond)
	a.pascals_per_secondLazy = &pascals_per_second
	return pascals_per_second
}

// PascalsPerMinute returns the PressureChangeRate value in PascalsPerMinute.
//
// 
func (a *PressureChangeRate) PascalsPerMinute() float64 {
	if a.pascals_per_minuteLazy != nil {
		return *a.pascals_per_minuteLazy
	}
	pascals_per_minute := a.convertFromBase(PressureChangeRatePascalPerMinute)
	a.pascals_per_minuteLazy = &pascals_per_minute
	return pascals_per_minute
}

// MillimetersOfMercuryPerSecond returns the PressureChangeRate value in MillimetersOfMercuryPerSecond.
//
// 
func (a *PressureChangeRate) MillimetersOfMercuryPerSecond() float64 {
	if a.millimeters_of_mercury_per_secondLazy != nil {
		return *a.millimeters_of_mercury_per_secondLazy
	}
	millimeters_of_mercury_per_second := a.convertFromBase(PressureChangeRateMillimeterOfMercuryPerSecond)
	a.millimeters_of_mercury_per_secondLazy = &millimeters_of_mercury_per_second
	return millimeters_of_mercury_per_second
}

// AtmospheresPerSecond returns the PressureChangeRate value in AtmospheresPerSecond.
//
// 
func (a *PressureChangeRate) AtmospheresPerSecond() float64 {
	if a.atmospheres_per_secondLazy != nil {
		return *a.atmospheres_per_secondLazy
	}
	atmospheres_per_second := a.convertFromBase(PressureChangeRateAtmospherePerSecond)
	a.atmospheres_per_secondLazy = &atmospheres_per_second
	return atmospheres_per_second
}

// PoundsForcePerSquareInchPerSecond returns the PressureChangeRate value in PoundsForcePerSquareInchPerSecond.
//
// 
func (a *PressureChangeRate) PoundsForcePerSquareInchPerSecond() float64 {
	if a.pounds_force_per_square_inch_per_secondLazy != nil {
		return *a.pounds_force_per_square_inch_per_secondLazy
	}
	pounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRatePoundForcePerSquareInchPerSecond)
	a.pounds_force_per_square_inch_per_secondLazy = &pounds_force_per_square_inch_per_second
	return pounds_force_per_square_inch_per_second
}

// PoundsForcePerSquareInchPerMinute returns the PressureChangeRate value in PoundsForcePerSquareInchPerMinute.
//
// 
func (a *PressureChangeRate) PoundsForcePerSquareInchPerMinute() float64 {
	if a.pounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.pounds_force_per_square_inch_per_minuteLazy
	}
	pounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRatePoundForcePerSquareInchPerMinute)
	a.pounds_force_per_square_inch_per_minuteLazy = &pounds_force_per_square_inch_per_minute
	return pounds_force_per_square_inch_per_minute
}

// BarsPerSecond returns the PressureChangeRate value in BarsPerSecond.
//
// 
func (a *PressureChangeRate) BarsPerSecond() float64 {
	if a.bars_per_secondLazy != nil {
		return *a.bars_per_secondLazy
	}
	bars_per_second := a.convertFromBase(PressureChangeRateBarPerSecond)
	a.bars_per_secondLazy = &bars_per_second
	return bars_per_second
}

// BarsPerMinute returns the PressureChangeRate value in BarsPerMinute.
//
// 
func (a *PressureChangeRate) BarsPerMinute() float64 {
	if a.bars_per_minuteLazy != nil {
		return *a.bars_per_minuteLazy
	}
	bars_per_minute := a.convertFromBase(PressureChangeRateBarPerMinute)
	a.bars_per_minuteLazy = &bars_per_minute
	return bars_per_minute
}

// KilopascalsPerSecond returns the PressureChangeRate value in KilopascalsPerSecond.
//
// 
func (a *PressureChangeRate) KilopascalsPerSecond() float64 {
	if a.kilopascals_per_secondLazy != nil {
		return *a.kilopascals_per_secondLazy
	}
	kilopascals_per_second := a.convertFromBase(PressureChangeRateKilopascalPerSecond)
	a.kilopascals_per_secondLazy = &kilopascals_per_second
	return kilopascals_per_second
}

// MegapascalsPerSecond returns the PressureChangeRate value in MegapascalsPerSecond.
//
// 
func (a *PressureChangeRate) MegapascalsPerSecond() float64 {
	if a.megapascals_per_secondLazy != nil {
		return *a.megapascals_per_secondLazy
	}
	megapascals_per_second := a.convertFromBase(PressureChangeRateMegapascalPerSecond)
	a.megapascals_per_secondLazy = &megapascals_per_second
	return megapascals_per_second
}

// KilopascalsPerMinute returns the PressureChangeRate value in KilopascalsPerMinute.
//
// 
func (a *PressureChangeRate) KilopascalsPerMinute() float64 {
	if a.kilopascals_per_minuteLazy != nil {
		return *a.kilopascals_per_minuteLazy
	}
	kilopascals_per_minute := a.convertFromBase(PressureChangeRateKilopascalPerMinute)
	a.kilopascals_per_minuteLazy = &kilopascals_per_minute
	return kilopascals_per_minute
}

// MegapascalsPerMinute returns the PressureChangeRate value in MegapascalsPerMinute.
//
// 
func (a *PressureChangeRate) MegapascalsPerMinute() float64 {
	if a.megapascals_per_minuteLazy != nil {
		return *a.megapascals_per_minuteLazy
	}
	megapascals_per_minute := a.convertFromBase(PressureChangeRateMegapascalPerMinute)
	a.megapascals_per_minuteLazy = &megapascals_per_minute
	return megapascals_per_minute
}

// KilopoundsForcePerSquareInchPerSecond returns the PressureChangeRate value in KilopoundsForcePerSquareInchPerSecond.
//
// 
func (a *PressureChangeRate) KilopoundsForcePerSquareInchPerSecond() float64 {
	if a.kilopounds_force_per_square_inch_per_secondLazy != nil {
		return *a.kilopounds_force_per_square_inch_per_secondLazy
	}
	kilopounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRateKilopoundForcePerSquareInchPerSecond)
	a.kilopounds_force_per_square_inch_per_secondLazy = &kilopounds_force_per_square_inch_per_second
	return kilopounds_force_per_square_inch_per_second
}

// MegapoundsForcePerSquareInchPerSecond returns the PressureChangeRate value in MegapoundsForcePerSquareInchPerSecond.
//
// 
func (a *PressureChangeRate) MegapoundsForcePerSquareInchPerSecond() float64 {
	if a.megapounds_force_per_square_inch_per_secondLazy != nil {
		return *a.megapounds_force_per_square_inch_per_secondLazy
	}
	megapounds_force_per_square_inch_per_second := a.convertFromBase(PressureChangeRateMegapoundForcePerSquareInchPerSecond)
	a.megapounds_force_per_square_inch_per_secondLazy = &megapounds_force_per_square_inch_per_second
	return megapounds_force_per_square_inch_per_second
}

// KilopoundsForcePerSquareInchPerMinute returns the PressureChangeRate value in KilopoundsForcePerSquareInchPerMinute.
//
// 
func (a *PressureChangeRate) KilopoundsForcePerSquareInchPerMinute() float64 {
	if a.kilopounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.kilopounds_force_per_square_inch_per_minuteLazy
	}
	kilopounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRateKilopoundForcePerSquareInchPerMinute)
	a.kilopounds_force_per_square_inch_per_minuteLazy = &kilopounds_force_per_square_inch_per_minute
	return kilopounds_force_per_square_inch_per_minute
}

// MegapoundsForcePerSquareInchPerMinute returns the PressureChangeRate value in MegapoundsForcePerSquareInchPerMinute.
//
// 
func (a *PressureChangeRate) MegapoundsForcePerSquareInchPerMinute() float64 {
	if a.megapounds_force_per_square_inch_per_minuteLazy != nil {
		return *a.megapounds_force_per_square_inch_per_minuteLazy
	}
	megapounds_force_per_square_inch_per_minute := a.convertFromBase(PressureChangeRateMegapoundForcePerSquareInchPerMinute)
	a.megapounds_force_per_square_inch_per_minuteLazy = &megapounds_force_per_square_inch_per_minute
	return megapounds_force_per_square_inch_per_minute
}

// MillibarsPerSecond returns the PressureChangeRate value in MillibarsPerSecond.
//
// 
func (a *PressureChangeRate) MillibarsPerSecond() float64 {
	if a.millibars_per_secondLazy != nil {
		return *a.millibars_per_secondLazy
	}
	millibars_per_second := a.convertFromBase(PressureChangeRateMillibarPerSecond)
	a.millibars_per_secondLazy = &millibars_per_second
	return millibars_per_second
}

// MillibarsPerMinute returns the PressureChangeRate value in MillibarsPerMinute.
//
// 
func (a *PressureChangeRate) MillibarsPerMinute() float64 {
	if a.millibars_per_minuteLazy != nil {
		return *a.millibars_per_minuteLazy
	}
	millibars_per_minute := a.convertFromBase(PressureChangeRateMillibarPerMinute)
	a.millibars_per_minuteLazy = &millibars_per_minute
	return millibars_per_minute
}


// ToDto creates a PressureChangeRateDto representation from the PressureChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalPerSecond by default.
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

// ToDtoJSON creates a JSON representation of the PressureChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by PascalPerSecond by default.
func (a *PressureChangeRate) ToDtoJSON(holdInUnit *PressureChangeRateUnits) ([]byte, error) {
	// Convert to PressureChangeRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a PressureChangeRate to a specific unit value.
// The function uses the provided unit type (PressureChangeRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the PressureChangeRate in the default unit (PascalPerSecond),
// formatted to two decimal places.
func (a PressureChangeRate) String() string {
	return a.ToString(PressureChangeRatePascalPerSecond, 2)
}

// ToString formats the PressureChangeRate value as a string with the specified unit and fractional digits.
// It converts the PressureChangeRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the PressureChangeRate value will be converted (e.g., PascalPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the PressureChangeRate with the unit abbreviation.
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

// Equals checks if the given PressureChangeRate is equal to the current PressureChangeRate.
//
// Parameters:
//    other: The PressureChangeRate to compare against.
//
// Returns:
//    bool: Returns true if both PressureChangeRate are equal, false otherwise.
func (a *PressureChangeRate) Equals(other *PressureChangeRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current PressureChangeRate with another PressureChangeRate.
// It returns -1 if the current PressureChangeRate is less than the other PressureChangeRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The PressureChangeRate to compare against.
//
// Returns:
//    int: -1 if the current PressureChangeRate is less, 1 if greater, and 0 if equal.
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

// Add adds the given PressureChangeRate to the current PressureChangeRate and returns the result.
// The result is a new PressureChangeRate instance with the sum of the values.
//
// Parameters:
//    other: The PressureChangeRate to add to the current PressureChangeRate.
//
// Returns:
//    *PressureChangeRate: A new PressureChangeRate instance representing the sum of both PressureChangeRate.
func (a *PressureChangeRate) Add(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given PressureChangeRate from the current PressureChangeRate and returns the result.
// The result is a new PressureChangeRate instance with the difference of the values.
//
// Parameters:
//    other: The PressureChangeRate to subtract from the current PressureChangeRate.
//
// Returns:
//    *PressureChangeRate: A new PressureChangeRate instance representing the difference of both PressureChangeRate.
func (a *PressureChangeRate) Subtract(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current PressureChangeRate by the given PressureChangeRate and returns the result.
// The result is a new PressureChangeRate instance with the product of the values.
//
// Parameters:
//    other: The PressureChangeRate to multiply with the current PressureChangeRate.
//
// Returns:
//    *PressureChangeRate: A new PressureChangeRate instance representing the product of both PressureChangeRate.
func (a *PressureChangeRate) Multiply(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value * other.BaseValue()}
}

// Divide divides the current PressureChangeRate by the given PressureChangeRate and returns the result.
// The result is a new PressureChangeRate instance with the quotient of the values.
//
// Parameters:
//    other: The PressureChangeRate to divide the current PressureChangeRate by.
//
// Returns:
//    *PressureChangeRate: A new PressureChangeRate instance representing the quotient of both PressureChangeRate.
func (a *PressureChangeRate) Divide(other *PressureChangeRate) *PressureChangeRate {
	return &PressureChangeRate{value: a.value / other.BaseValue()}
}