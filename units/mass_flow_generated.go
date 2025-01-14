// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassFlowUnits defines various units of MassFlow.
type MassFlowUnits string

const (
    
        // 
        MassFlowGramPerSecond MassFlowUnits = "GramPerSecond"
        // 
        MassFlowGramPerDay MassFlowUnits = "GramPerDay"
        // 
        MassFlowGramPerHour MassFlowUnits = "GramPerHour"
        // 
        MassFlowKilogramPerHour MassFlowUnits = "KilogramPerHour"
        // 
        MassFlowKilogramPerMinute MassFlowUnits = "KilogramPerMinute"
        // 
        MassFlowTonnePerHour MassFlowUnits = "TonnePerHour"
        // 
        MassFlowPoundPerDay MassFlowUnits = "PoundPerDay"
        // 
        MassFlowPoundPerHour MassFlowUnits = "PoundPerHour"
        // 
        MassFlowPoundPerMinute MassFlowUnits = "PoundPerMinute"
        // 
        MassFlowPoundPerSecond MassFlowUnits = "PoundPerSecond"
        // 
        MassFlowTonnePerDay MassFlowUnits = "TonnePerDay"
        // 
        MassFlowShortTonPerHour MassFlowUnits = "ShortTonPerHour"
        // 
        MassFlowNanogramPerSecond MassFlowUnits = "NanogramPerSecond"
        // 
        MassFlowMicrogramPerSecond MassFlowUnits = "MicrogramPerSecond"
        // 
        MassFlowMilligramPerSecond MassFlowUnits = "MilligramPerSecond"
        // 
        MassFlowCentigramPerSecond MassFlowUnits = "CentigramPerSecond"
        // 
        MassFlowDecigramPerSecond MassFlowUnits = "DecigramPerSecond"
        // 
        MassFlowDecagramPerSecond MassFlowUnits = "DecagramPerSecond"
        // 
        MassFlowHectogramPerSecond MassFlowUnits = "HectogramPerSecond"
        // 
        MassFlowKilogramPerSecond MassFlowUnits = "KilogramPerSecond"
        // 
        MassFlowNanogramPerDay MassFlowUnits = "NanogramPerDay"
        // 
        MassFlowMicrogramPerDay MassFlowUnits = "MicrogramPerDay"
        // 
        MassFlowMilligramPerDay MassFlowUnits = "MilligramPerDay"
        // 
        MassFlowCentigramPerDay MassFlowUnits = "CentigramPerDay"
        // 
        MassFlowDecigramPerDay MassFlowUnits = "DecigramPerDay"
        // 
        MassFlowDecagramPerDay MassFlowUnits = "DecagramPerDay"
        // 
        MassFlowHectogramPerDay MassFlowUnits = "HectogramPerDay"
        // 
        MassFlowKilogramPerDay MassFlowUnits = "KilogramPerDay"
        // 
        MassFlowMegagramPerDay MassFlowUnits = "MegagramPerDay"
        // 
        MassFlowMegapoundPerDay MassFlowUnits = "MegapoundPerDay"
        // 
        MassFlowMegapoundPerHour MassFlowUnits = "MegapoundPerHour"
        // 
        MassFlowMegapoundPerMinute MassFlowUnits = "MegapoundPerMinute"
        // 
        MassFlowMegapoundPerSecond MassFlowUnits = "MegapoundPerSecond"
)

// MassFlowDto represents a MassFlow measurement with a numerical value and its corresponding unit.
type MassFlowDto struct {
    // Value is the numerical representation of the MassFlow.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MassFlow, as defined in the MassFlowUnits enumeration.
	Unit  MassFlowUnits `json:"unit"`
}

// MassFlowDtoFactory groups methods for creating and serializing MassFlowDto objects.
type MassFlowDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassFlowDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassFlowDtoFactory) FromJSON(data []byte) (*MassFlowDto, error) {
	a := MassFlowDto{}

    // Parse JSON into MassFlowDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MassFlowDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MassFlow represents a measurement in a of MassFlow.
//
// Mass flow is the ratio of the mass change to the time during which the change occurred (value of mass changes per unit time).
type MassFlow struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_secondLazy *float64 
    grams_per_dayLazy *float64 
    grams_per_hourLazy *float64 
    kilograms_per_hourLazy *float64 
    kilograms_per_minuteLazy *float64 
    tonnes_per_hourLazy *float64 
    pounds_per_dayLazy *float64 
    pounds_per_hourLazy *float64 
    pounds_per_minuteLazy *float64 
    pounds_per_secondLazy *float64 
    tonnes_per_dayLazy *float64 
    short_tons_per_hourLazy *float64 
    nanograms_per_secondLazy *float64 
    micrograms_per_secondLazy *float64 
    milligrams_per_secondLazy *float64 
    centigrams_per_secondLazy *float64 
    decigrams_per_secondLazy *float64 
    decagrams_per_secondLazy *float64 
    hectograms_per_secondLazy *float64 
    kilograms_per_secondLazy *float64 
    nanograms_per_dayLazy *float64 
    micrograms_per_dayLazy *float64 
    milligrams_per_dayLazy *float64 
    centigrams_per_dayLazy *float64 
    decigrams_per_dayLazy *float64 
    decagrams_per_dayLazy *float64 
    hectograms_per_dayLazy *float64 
    kilograms_per_dayLazy *float64 
    megagrams_per_dayLazy *float64 
    megapounds_per_dayLazy *float64 
    megapounds_per_hourLazy *float64 
    megapounds_per_minuteLazy *float64 
    megapounds_per_secondLazy *float64 
}

// MassFlowFactory groups methods for creating MassFlow instances.
type MassFlowFactory struct{}

// CreateMassFlow creates a new MassFlow instance from the given value and unit.
func (uf MassFlowFactory) CreateMassFlow(value float64, unit MassFlowUnits) (*MassFlow, error) {
	return newMassFlow(value, unit)
}

// FromDto converts a MassFlowDto to a MassFlow instance.
func (uf MassFlowFactory) FromDto(dto MassFlowDto) (*MassFlow, error) {
	return newMassFlow(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MassFlow instance.
func (uf MassFlowFactory) FromDtoJSON(data []byte) (*MassFlow, error) {
	unitDto, err := MassFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassFlowDto from JSON: %w", err)
	}
	return MassFlowFactory{}.FromDto(*unitDto)
}


// FromGramsPerSecond creates a new MassFlow instance from a value in GramsPerSecond.
func (uf MassFlowFactory) FromGramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerSecond)
}

// FromGramsPerDay creates a new MassFlow instance from a value in GramsPerDay.
func (uf MassFlowFactory) FromGramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerDay)
}

// FromGramsPerHour creates a new MassFlow instance from a value in GramsPerHour.
func (uf MassFlowFactory) FromGramsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerHour)
}

// FromKilogramsPerHour creates a new MassFlow instance from a value in KilogramsPerHour.
func (uf MassFlowFactory) FromKilogramsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerHour)
}

// FromKilogramsPerMinute creates a new MassFlow instance from a value in KilogramsPerMinute.
func (uf MassFlowFactory) FromKilogramsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerMinute)
}

// FromTonnesPerHour creates a new MassFlow instance from a value in TonnesPerHour.
func (uf MassFlowFactory) FromTonnesPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowTonnePerHour)
}

// FromPoundsPerDay creates a new MassFlow instance from a value in PoundsPerDay.
func (uf MassFlowFactory) FromPoundsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerDay)
}

// FromPoundsPerHour creates a new MassFlow instance from a value in PoundsPerHour.
func (uf MassFlowFactory) FromPoundsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerHour)
}

// FromPoundsPerMinute creates a new MassFlow instance from a value in PoundsPerMinute.
func (uf MassFlowFactory) FromPoundsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerMinute)
}

// FromPoundsPerSecond creates a new MassFlow instance from a value in PoundsPerSecond.
func (uf MassFlowFactory) FromPoundsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerSecond)
}

// FromTonnesPerDay creates a new MassFlow instance from a value in TonnesPerDay.
func (uf MassFlowFactory) FromTonnesPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowTonnePerDay)
}

// FromShortTonsPerHour creates a new MassFlow instance from a value in ShortTonsPerHour.
func (uf MassFlowFactory) FromShortTonsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowShortTonPerHour)
}

// FromNanogramsPerSecond creates a new MassFlow instance from a value in NanogramsPerSecond.
func (uf MassFlowFactory) FromNanogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowNanogramPerSecond)
}

// FromMicrogramsPerSecond creates a new MassFlow instance from a value in MicrogramsPerSecond.
func (uf MassFlowFactory) FromMicrogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMicrogramPerSecond)
}

// FromMilligramsPerSecond creates a new MassFlow instance from a value in MilligramsPerSecond.
func (uf MassFlowFactory) FromMilligramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMilligramPerSecond)
}

// FromCentigramsPerSecond creates a new MassFlow instance from a value in CentigramsPerSecond.
func (uf MassFlowFactory) FromCentigramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowCentigramPerSecond)
}

// FromDecigramsPerSecond creates a new MassFlow instance from a value in DecigramsPerSecond.
func (uf MassFlowFactory) FromDecigramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecigramPerSecond)
}

// FromDecagramsPerSecond creates a new MassFlow instance from a value in DecagramsPerSecond.
func (uf MassFlowFactory) FromDecagramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecagramPerSecond)
}

// FromHectogramsPerSecond creates a new MassFlow instance from a value in HectogramsPerSecond.
func (uf MassFlowFactory) FromHectogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowHectogramPerSecond)
}

// FromKilogramsPerSecond creates a new MassFlow instance from a value in KilogramsPerSecond.
func (uf MassFlowFactory) FromKilogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerSecond)
}

// FromNanogramsPerDay creates a new MassFlow instance from a value in NanogramsPerDay.
func (uf MassFlowFactory) FromNanogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowNanogramPerDay)
}

// FromMicrogramsPerDay creates a new MassFlow instance from a value in MicrogramsPerDay.
func (uf MassFlowFactory) FromMicrogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMicrogramPerDay)
}

// FromMilligramsPerDay creates a new MassFlow instance from a value in MilligramsPerDay.
func (uf MassFlowFactory) FromMilligramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMilligramPerDay)
}

// FromCentigramsPerDay creates a new MassFlow instance from a value in CentigramsPerDay.
func (uf MassFlowFactory) FromCentigramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowCentigramPerDay)
}

// FromDecigramsPerDay creates a new MassFlow instance from a value in DecigramsPerDay.
func (uf MassFlowFactory) FromDecigramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecigramPerDay)
}

// FromDecagramsPerDay creates a new MassFlow instance from a value in DecagramsPerDay.
func (uf MassFlowFactory) FromDecagramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecagramPerDay)
}

// FromHectogramsPerDay creates a new MassFlow instance from a value in HectogramsPerDay.
func (uf MassFlowFactory) FromHectogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowHectogramPerDay)
}

// FromKilogramsPerDay creates a new MassFlow instance from a value in KilogramsPerDay.
func (uf MassFlowFactory) FromKilogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerDay)
}

// FromMegagramsPerDay creates a new MassFlow instance from a value in MegagramsPerDay.
func (uf MassFlowFactory) FromMegagramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegagramPerDay)
}

// FromMegapoundsPerDay creates a new MassFlow instance from a value in MegapoundsPerDay.
func (uf MassFlowFactory) FromMegapoundsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerDay)
}

// FromMegapoundsPerHour creates a new MassFlow instance from a value in MegapoundsPerHour.
func (uf MassFlowFactory) FromMegapoundsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerHour)
}

// FromMegapoundsPerMinute creates a new MassFlow instance from a value in MegapoundsPerMinute.
func (uf MassFlowFactory) FromMegapoundsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerMinute)
}

// FromMegapoundsPerSecond creates a new MassFlow instance from a value in MegapoundsPerSecond.
func (uf MassFlowFactory) FromMegapoundsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerSecond)
}


// newMassFlow creates a new MassFlow.
func newMassFlow(value float64, fromUnit MassFlowUnits) (*MassFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MassFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassFlow in GramPerSecond unit (the base/default quantity).
func (a *MassFlow) BaseValue() float64 {
	return a.value
}


// GramsPerSecond returns the MassFlow value in GramsPerSecond.
//
// 
func (a *MassFlow) GramsPerSecond() float64 {
	if a.grams_per_secondLazy != nil {
		return *a.grams_per_secondLazy
	}
	grams_per_second := a.convertFromBase(MassFlowGramPerSecond)
	a.grams_per_secondLazy = &grams_per_second
	return grams_per_second
}

// GramsPerDay returns the MassFlow value in GramsPerDay.
//
// 
func (a *MassFlow) GramsPerDay() float64 {
	if a.grams_per_dayLazy != nil {
		return *a.grams_per_dayLazy
	}
	grams_per_day := a.convertFromBase(MassFlowGramPerDay)
	a.grams_per_dayLazy = &grams_per_day
	return grams_per_day
}

// GramsPerHour returns the MassFlow value in GramsPerHour.
//
// 
func (a *MassFlow) GramsPerHour() float64 {
	if a.grams_per_hourLazy != nil {
		return *a.grams_per_hourLazy
	}
	grams_per_hour := a.convertFromBase(MassFlowGramPerHour)
	a.grams_per_hourLazy = &grams_per_hour
	return grams_per_hour
}

// KilogramsPerHour returns the MassFlow value in KilogramsPerHour.
//
// 
func (a *MassFlow) KilogramsPerHour() float64 {
	if a.kilograms_per_hourLazy != nil {
		return *a.kilograms_per_hourLazy
	}
	kilograms_per_hour := a.convertFromBase(MassFlowKilogramPerHour)
	a.kilograms_per_hourLazy = &kilograms_per_hour
	return kilograms_per_hour
}

// KilogramsPerMinute returns the MassFlow value in KilogramsPerMinute.
//
// 
func (a *MassFlow) KilogramsPerMinute() float64 {
	if a.kilograms_per_minuteLazy != nil {
		return *a.kilograms_per_minuteLazy
	}
	kilograms_per_minute := a.convertFromBase(MassFlowKilogramPerMinute)
	a.kilograms_per_minuteLazy = &kilograms_per_minute
	return kilograms_per_minute
}

// TonnesPerHour returns the MassFlow value in TonnesPerHour.
//
// 
func (a *MassFlow) TonnesPerHour() float64 {
	if a.tonnes_per_hourLazy != nil {
		return *a.tonnes_per_hourLazy
	}
	tonnes_per_hour := a.convertFromBase(MassFlowTonnePerHour)
	a.tonnes_per_hourLazy = &tonnes_per_hour
	return tonnes_per_hour
}

// PoundsPerDay returns the MassFlow value in PoundsPerDay.
//
// 
func (a *MassFlow) PoundsPerDay() float64 {
	if a.pounds_per_dayLazy != nil {
		return *a.pounds_per_dayLazy
	}
	pounds_per_day := a.convertFromBase(MassFlowPoundPerDay)
	a.pounds_per_dayLazy = &pounds_per_day
	return pounds_per_day
}

// PoundsPerHour returns the MassFlow value in PoundsPerHour.
//
// 
func (a *MassFlow) PoundsPerHour() float64 {
	if a.pounds_per_hourLazy != nil {
		return *a.pounds_per_hourLazy
	}
	pounds_per_hour := a.convertFromBase(MassFlowPoundPerHour)
	a.pounds_per_hourLazy = &pounds_per_hour
	return pounds_per_hour
}

// PoundsPerMinute returns the MassFlow value in PoundsPerMinute.
//
// 
func (a *MassFlow) PoundsPerMinute() float64 {
	if a.pounds_per_minuteLazy != nil {
		return *a.pounds_per_minuteLazy
	}
	pounds_per_minute := a.convertFromBase(MassFlowPoundPerMinute)
	a.pounds_per_minuteLazy = &pounds_per_minute
	return pounds_per_minute
}

// PoundsPerSecond returns the MassFlow value in PoundsPerSecond.
//
// 
func (a *MassFlow) PoundsPerSecond() float64 {
	if a.pounds_per_secondLazy != nil {
		return *a.pounds_per_secondLazy
	}
	pounds_per_second := a.convertFromBase(MassFlowPoundPerSecond)
	a.pounds_per_secondLazy = &pounds_per_second
	return pounds_per_second
}

// TonnesPerDay returns the MassFlow value in TonnesPerDay.
//
// 
func (a *MassFlow) TonnesPerDay() float64 {
	if a.tonnes_per_dayLazy != nil {
		return *a.tonnes_per_dayLazy
	}
	tonnes_per_day := a.convertFromBase(MassFlowTonnePerDay)
	a.tonnes_per_dayLazy = &tonnes_per_day
	return tonnes_per_day
}

// ShortTonsPerHour returns the MassFlow value in ShortTonsPerHour.
//
// 
func (a *MassFlow) ShortTonsPerHour() float64 {
	if a.short_tons_per_hourLazy != nil {
		return *a.short_tons_per_hourLazy
	}
	short_tons_per_hour := a.convertFromBase(MassFlowShortTonPerHour)
	a.short_tons_per_hourLazy = &short_tons_per_hour
	return short_tons_per_hour
}

// NanogramsPerSecond returns the MassFlow value in NanogramsPerSecond.
//
// 
func (a *MassFlow) NanogramsPerSecond() float64 {
	if a.nanograms_per_secondLazy != nil {
		return *a.nanograms_per_secondLazy
	}
	nanograms_per_second := a.convertFromBase(MassFlowNanogramPerSecond)
	a.nanograms_per_secondLazy = &nanograms_per_second
	return nanograms_per_second
}

// MicrogramsPerSecond returns the MassFlow value in MicrogramsPerSecond.
//
// 
func (a *MassFlow) MicrogramsPerSecond() float64 {
	if a.micrograms_per_secondLazy != nil {
		return *a.micrograms_per_secondLazy
	}
	micrograms_per_second := a.convertFromBase(MassFlowMicrogramPerSecond)
	a.micrograms_per_secondLazy = &micrograms_per_second
	return micrograms_per_second
}

// MilligramsPerSecond returns the MassFlow value in MilligramsPerSecond.
//
// 
func (a *MassFlow) MilligramsPerSecond() float64 {
	if a.milligrams_per_secondLazy != nil {
		return *a.milligrams_per_secondLazy
	}
	milligrams_per_second := a.convertFromBase(MassFlowMilligramPerSecond)
	a.milligrams_per_secondLazy = &milligrams_per_second
	return milligrams_per_second
}

// CentigramsPerSecond returns the MassFlow value in CentigramsPerSecond.
//
// 
func (a *MassFlow) CentigramsPerSecond() float64 {
	if a.centigrams_per_secondLazy != nil {
		return *a.centigrams_per_secondLazy
	}
	centigrams_per_second := a.convertFromBase(MassFlowCentigramPerSecond)
	a.centigrams_per_secondLazy = &centigrams_per_second
	return centigrams_per_second
}

// DecigramsPerSecond returns the MassFlow value in DecigramsPerSecond.
//
// 
func (a *MassFlow) DecigramsPerSecond() float64 {
	if a.decigrams_per_secondLazy != nil {
		return *a.decigrams_per_secondLazy
	}
	decigrams_per_second := a.convertFromBase(MassFlowDecigramPerSecond)
	a.decigrams_per_secondLazy = &decigrams_per_second
	return decigrams_per_second
}

// DecagramsPerSecond returns the MassFlow value in DecagramsPerSecond.
//
// 
func (a *MassFlow) DecagramsPerSecond() float64 {
	if a.decagrams_per_secondLazy != nil {
		return *a.decagrams_per_secondLazy
	}
	decagrams_per_second := a.convertFromBase(MassFlowDecagramPerSecond)
	a.decagrams_per_secondLazy = &decagrams_per_second
	return decagrams_per_second
}

// HectogramsPerSecond returns the MassFlow value in HectogramsPerSecond.
//
// 
func (a *MassFlow) HectogramsPerSecond() float64 {
	if a.hectograms_per_secondLazy != nil {
		return *a.hectograms_per_secondLazy
	}
	hectograms_per_second := a.convertFromBase(MassFlowHectogramPerSecond)
	a.hectograms_per_secondLazy = &hectograms_per_second
	return hectograms_per_second
}

// KilogramsPerSecond returns the MassFlow value in KilogramsPerSecond.
//
// 
func (a *MassFlow) KilogramsPerSecond() float64 {
	if a.kilograms_per_secondLazy != nil {
		return *a.kilograms_per_secondLazy
	}
	kilograms_per_second := a.convertFromBase(MassFlowKilogramPerSecond)
	a.kilograms_per_secondLazy = &kilograms_per_second
	return kilograms_per_second
}

// NanogramsPerDay returns the MassFlow value in NanogramsPerDay.
//
// 
func (a *MassFlow) NanogramsPerDay() float64 {
	if a.nanograms_per_dayLazy != nil {
		return *a.nanograms_per_dayLazy
	}
	nanograms_per_day := a.convertFromBase(MassFlowNanogramPerDay)
	a.nanograms_per_dayLazy = &nanograms_per_day
	return nanograms_per_day
}

// MicrogramsPerDay returns the MassFlow value in MicrogramsPerDay.
//
// 
func (a *MassFlow) MicrogramsPerDay() float64 {
	if a.micrograms_per_dayLazy != nil {
		return *a.micrograms_per_dayLazy
	}
	micrograms_per_day := a.convertFromBase(MassFlowMicrogramPerDay)
	a.micrograms_per_dayLazy = &micrograms_per_day
	return micrograms_per_day
}

// MilligramsPerDay returns the MassFlow value in MilligramsPerDay.
//
// 
func (a *MassFlow) MilligramsPerDay() float64 {
	if a.milligrams_per_dayLazy != nil {
		return *a.milligrams_per_dayLazy
	}
	milligrams_per_day := a.convertFromBase(MassFlowMilligramPerDay)
	a.milligrams_per_dayLazy = &milligrams_per_day
	return milligrams_per_day
}

// CentigramsPerDay returns the MassFlow value in CentigramsPerDay.
//
// 
func (a *MassFlow) CentigramsPerDay() float64 {
	if a.centigrams_per_dayLazy != nil {
		return *a.centigrams_per_dayLazy
	}
	centigrams_per_day := a.convertFromBase(MassFlowCentigramPerDay)
	a.centigrams_per_dayLazy = &centigrams_per_day
	return centigrams_per_day
}

// DecigramsPerDay returns the MassFlow value in DecigramsPerDay.
//
// 
func (a *MassFlow) DecigramsPerDay() float64 {
	if a.decigrams_per_dayLazy != nil {
		return *a.decigrams_per_dayLazy
	}
	decigrams_per_day := a.convertFromBase(MassFlowDecigramPerDay)
	a.decigrams_per_dayLazy = &decigrams_per_day
	return decigrams_per_day
}

// DecagramsPerDay returns the MassFlow value in DecagramsPerDay.
//
// 
func (a *MassFlow) DecagramsPerDay() float64 {
	if a.decagrams_per_dayLazy != nil {
		return *a.decagrams_per_dayLazy
	}
	decagrams_per_day := a.convertFromBase(MassFlowDecagramPerDay)
	a.decagrams_per_dayLazy = &decagrams_per_day
	return decagrams_per_day
}

// HectogramsPerDay returns the MassFlow value in HectogramsPerDay.
//
// 
func (a *MassFlow) HectogramsPerDay() float64 {
	if a.hectograms_per_dayLazy != nil {
		return *a.hectograms_per_dayLazy
	}
	hectograms_per_day := a.convertFromBase(MassFlowHectogramPerDay)
	a.hectograms_per_dayLazy = &hectograms_per_day
	return hectograms_per_day
}

// KilogramsPerDay returns the MassFlow value in KilogramsPerDay.
//
// 
func (a *MassFlow) KilogramsPerDay() float64 {
	if a.kilograms_per_dayLazy != nil {
		return *a.kilograms_per_dayLazy
	}
	kilograms_per_day := a.convertFromBase(MassFlowKilogramPerDay)
	a.kilograms_per_dayLazy = &kilograms_per_day
	return kilograms_per_day
}

// MegagramsPerDay returns the MassFlow value in MegagramsPerDay.
//
// 
func (a *MassFlow) MegagramsPerDay() float64 {
	if a.megagrams_per_dayLazy != nil {
		return *a.megagrams_per_dayLazy
	}
	megagrams_per_day := a.convertFromBase(MassFlowMegagramPerDay)
	a.megagrams_per_dayLazy = &megagrams_per_day
	return megagrams_per_day
}

// MegapoundsPerDay returns the MassFlow value in MegapoundsPerDay.
//
// 
func (a *MassFlow) MegapoundsPerDay() float64 {
	if a.megapounds_per_dayLazy != nil {
		return *a.megapounds_per_dayLazy
	}
	megapounds_per_day := a.convertFromBase(MassFlowMegapoundPerDay)
	a.megapounds_per_dayLazy = &megapounds_per_day
	return megapounds_per_day
}

// MegapoundsPerHour returns the MassFlow value in MegapoundsPerHour.
//
// 
func (a *MassFlow) MegapoundsPerHour() float64 {
	if a.megapounds_per_hourLazy != nil {
		return *a.megapounds_per_hourLazy
	}
	megapounds_per_hour := a.convertFromBase(MassFlowMegapoundPerHour)
	a.megapounds_per_hourLazy = &megapounds_per_hour
	return megapounds_per_hour
}

// MegapoundsPerMinute returns the MassFlow value in MegapoundsPerMinute.
//
// 
func (a *MassFlow) MegapoundsPerMinute() float64 {
	if a.megapounds_per_minuteLazy != nil {
		return *a.megapounds_per_minuteLazy
	}
	megapounds_per_minute := a.convertFromBase(MassFlowMegapoundPerMinute)
	a.megapounds_per_minuteLazy = &megapounds_per_minute
	return megapounds_per_minute
}

// MegapoundsPerSecond returns the MassFlow value in MegapoundsPerSecond.
//
// 
func (a *MassFlow) MegapoundsPerSecond() float64 {
	if a.megapounds_per_secondLazy != nil {
		return *a.megapounds_per_secondLazy
	}
	megapounds_per_second := a.convertFromBase(MassFlowMegapoundPerSecond)
	a.megapounds_per_secondLazy = &megapounds_per_second
	return megapounds_per_second
}


// ToDto creates a MassFlowDto representation from the MassFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GramPerSecond by default.
func (a *MassFlow) ToDto(holdInUnit *MassFlowUnits) MassFlowDto {
	if holdInUnit == nil {
		defaultUnit := MassFlowGramPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return MassFlowDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MassFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GramPerSecond by default.
func (a *MassFlow) ToDtoJSON(holdInUnit *MassFlowUnits) ([]byte, error) {
	// Convert to MassFlowDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MassFlow to a specific unit value.
// The function uses the provided unit type (MassFlowUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MassFlow) Convert(toUnit MassFlowUnits) float64 {
	switch toUnit { 
    case MassFlowGramPerSecond:
		return a.GramsPerSecond()
    case MassFlowGramPerDay:
		return a.GramsPerDay()
    case MassFlowGramPerHour:
		return a.GramsPerHour()
    case MassFlowKilogramPerHour:
		return a.KilogramsPerHour()
    case MassFlowKilogramPerMinute:
		return a.KilogramsPerMinute()
    case MassFlowTonnePerHour:
		return a.TonnesPerHour()
    case MassFlowPoundPerDay:
		return a.PoundsPerDay()
    case MassFlowPoundPerHour:
		return a.PoundsPerHour()
    case MassFlowPoundPerMinute:
		return a.PoundsPerMinute()
    case MassFlowPoundPerSecond:
		return a.PoundsPerSecond()
    case MassFlowTonnePerDay:
		return a.TonnesPerDay()
    case MassFlowShortTonPerHour:
		return a.ShortTonsPerHour()
    case MassFlowNanogramPerSecond:
		return a.NanogramsPerSecond()
    case MassFlowMicrogramPerSecond:
		return a.MicrogramsPerSecond()
    case MassFlowMilligramPerSecond:
		return a.MilligramsPerSecond()
    case MassFlowCentigramPerSecond:
		return a.CentigramsPerSecond()
    case MassFlowDecigramPerSecond:
		return a.DecigramsPerSecond()
    case MassFlowDecagramPerSecond:
		return a.DecagramsPerSecond()
    case MassFlowHectogramPerSecond:
		return a.HectogramsPerSecond()
    case MassFlowKilogramPerSecond:
		return a.KilogramsPerSecond()
    case MassFlowNanogramPerDay:
		return a.NanogramsPerDay()
    case MassFlowMicrogramPerDay:
		return a.MicrogramsPerDay()
    case MassFlowMilligramPerDay:
		return a.MilligramsPerDay()
    case MassFlowCentigramPerDay:
		return a.CentigramsPerDay()
    case MassFlowDecigramPerDay:
		return a.DecigramsPerDay()
    case MassFlowDecagramPerDay:
		return a.DecagramsPerDay()
    case MassFlowHectogramPerDay:
		return a.HectogramsPerDay()
    case MassFlowKilogramPerDay:
		return a.KilogramsPerDay()
    case MassFlowMegagramPerDay:
		return a.MegagramsPerDay()
    case MassFlowMegapoundPerDay:
		return a.MegapoundsPerDay()
    case MassFlowMegapoundPerHour:
		return a.MegapoundsPerHour()
    case MassFlowMegapoundPerMinute:
		return a.MegapoundsPerMinute()
    case MassFlowMegapoundPerSecond:
		return a.MegapoundsPerSecond()
	default:
		return math.NaN()
	}
}

func (a *MassFlow) convertFromBase(toUnit MassFlowUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassFlowGramPerSecond:
		return (value) 
	case MassFlowGramPerDay:
		return (value * 86400) 
	case MassFlowGramPerHour:
		return (value * 3600) 
	case MassFlowKilogramPerHour:
		return (value * 3.6) 
	case MassFlowKilogramPerMinute:
		return (value * 0.06) 
	case MassFlowTonnePerHour:
		return (value * 3.6 / 1000) 
	case MassFlowPoundPerDay:
		return (value * 190.47936) 
	case MassFlowPoundPerHour:
		return (value * 7.93664) 
	case MassFlowPoundPerMinute:
		return (value * 0.132277) 
	case MassFlowPoundPerSecond:
		return (value / 453.59237) 
	case MassFlowTonnePerDay:
		return (value * 0.0864000) 
	case MassFlowShortTonPerHour:
		return (value / 251.9957611) 
	case MassFlowNanogramPerSecond:
		return ((value) / 1e-09) 
	case MassFlowMicrogramPerSecond:
		return ((value) / 1e-06) 
	case MassFlowMilligramPerSecond:
		return ((value) / 0.001) 
	case MassFlowCentigramPerSecond:
		return ((value) / 0.01) 
	case MassFlowDecigramPerSecond:
		return ((value) / 0.1) 
	case MassFlowDecagramPerSecond:
		return ((value) / 10.0) 
	case MassFlowHectogramPerSecond:
		return ((value) / 100.0) 
	case MassFlowKilogramPerSecond:
		return ((value) / 1000.0) 
	case MassFlowNanogramPerDay:
		return ((value * 86400) / 1e-09) 
	case MassFlowMicrogramPerDay:
		return ((value * 86400) / 1e-06) 
	case MassFlowMilligramPerDay:
		return ((value * 86400) / 0.001) 
	case MassFlowCentigramPerDay:
		return ((value * 86400) / 0.01) 
	case MassFlowDecigramPerDay:
		return ((value * 86400) / 0.1) 
	case MassFlowDecagramPerDay:
		return ((value * 86400) / 10.0) 
	case MassFlowHectogramPerDay:
		return ((value * 86400) / 100.0) 
	case MassFlowKilogramPerDay:
		return ((value * 86400) / 1000.0) 
	case MassFlowMegagramPerDay:
		return ((value * 86400) / 1000000.0) 
	case MassFlowMegapoundPerDay:
		return ((value * 190.47936) / 1000000.0) 
	case MassFlowMegapoundPerHour:
		return ((value * 7.93664) / 1000000.0) 
	case MassFlowMegapoundPerMinute:
		return ((value * 0.132277) / 1000000.0) 
	case MassFlowMegapoundPerSecond:
		return ((value / 453.59237) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *MassFlow) convertToBase(value float64, fromUnit MassFlowUnits) float64 {
	switch fromUnit { 
	case MassFlowGramPerSecond:
		return (value) 
	case MassFlowGramPerDay:
		return (value / 86400) 
	case MassFlowGramPerHour:
		return (value / 3600) 
	case MassFlowKilogramPerHour:
		return (value / 3.6) 
	case MassFlowKilogramPerMinute:
		return (value / 0.06) 
	case MassFlowTonnePerHour:
		return (1000 * value / 3.6) 
	case MassFlowPoundPerDay:
		return (value / 190.47936) 
	case MassFlowPoundPerHour:
		return (value / 7.93664) 
	case MassFlowPoundPerMinute:
		return (value / 0.132277) 
	case MassFlowPoundPerSecond:
		return (value * 453.59237) 
	case MassFlowTonnePerDay:
		return (value / 0.0864000) 
	case MassFlowShortTonPerHour:
		return (value * 251.9957611) 
	case MassFlowNanogramPerSecond:
		return ((value) * 1e-09) 
	case MassFlowMicrogramPerSecond:
		return ((value) * 1e-06) 
	case MassFlowMilligramPerSecond:
		return ((value) * 0.001) 
	case MassFlowCentigramPerSecond:
		return ((value) * 0.01) 
	case MassFlowDecigramPerSecond:
		return ((value) * 0.1) 
	case MassFlowDecagramPerSecond:
		return ((value) * 10.0) 
	case MassFlowHectogramPerSecond:
		return ((value) * 100.0) 
	case MassFlowKilogramPerSecond:
		return ((value) * 1000.0) 
	case MassFlowNanogramPerDay:
		return ((value / 86400) * 1e-09) 
	case MassFlowMicrogramPerDay:
		return ((value / 86400) * 1e-06) 
	case MassFlowMilligramPerDay:
		return ((value / 86400) * 0.001) 
	case MassFlowCentigramPerDay:
		return ((value / 86400) * 0.01) 
	case MassFlowDecigramPerDay:
		return ((value / 86400) * 0.1) 
	case MassFlowDecagramPerDay:
		return ((value / 86400) * 10.0) 
	case MassFlowHectogramPerDay:
		return ((value / 86400) * 100.0) 
	case MassFlowKilogramPerDay:
		return ((value / 86400) * 1000.0) 
	case MassFlowMegagramPerDay:
		return ((value / 86400) * 1000000.0) 
	case MassFlowMegapoundPerDay:
		return ((value / 190.47936) * 1000000.0) 
	case MassFlowMegapoundPerHour:
		return ((value / 7.93664) * 1000000.0) 
	case MassFlowMegapoundPerMinute:
		return ((value / 0.132277) * 1000000.0) 
	case MassFlowMegapoundPerSecond:
		return ((value * 453.59237) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MassFlow in the default unit (GramPerSecond),
// formatted to two decimal places.
func (a MassFlow) String() string {
	return a.ToString(MassFlowGramPerSecond, 2)
}

// ToString formats the MassFlow value as a string with the specified unit and fractional digits.
// It converts the MassFlow to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MassFlow value will be converted (e.g., GramPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MassFlow with the unit abbreviation.
func (a *MassFlow) ToString(unit MassFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MassFlow) getUnitAbbreviation(unit MassFlowUnits) string {
	switch unit { 
	case MassFlowGramPerSecond:
		return "g/s" 
	case MassFlowGramPerDay:
		return "g/d" 
	case MassFlowGramPerHour:
		return "g/h" 
	case MassFlowKilogramPerHour:
		return "kg/h" 
	case MassFlowKilogramPerMinute:
		return "kg/min" 
	case MassFlowTonnePerHour:
		return "t/h" 
	case MassFlowPoundPerDay:
		return "lb/d" 
	case MassFlowPoundPerHour:
		return "lb/h" 
	case MassFlowPoundPerMinute:
		return "lb/min" 
	case MassFlowPoundPerSecond:
		return "lb/s" 
	case MassFlowTonnePerDay:
		return "t/d" 
	case MassFlowShortTonPerHour:
		return "short tn/h" 
	case MassFlowNanogramPerSecond:
		return "ng/s" 
	case MassFlowMicrogramPerSecond:
		return "μg/s" 
	case MassFlowMilligramPerSecond:
		return "mg/s" 
	case MassFlowCentigramPerSecond:
		return "cg/s" 
	case MassFlowDecigramPerSecond:
		return "dg/s" 
	case MassFlowDecagramPerSecond:
		return "dag/s" 
	case MassFlowHectogramPerSecond:
		return "hg/s" 
	case MassFlowKilogramPerSecond:
		return "kg/s" 
	case MassFlowNanogramPerDay:
		return "ng/d" 
	case MassFlowMicrogramPerDay:
		return "μg/d" 
	case MassFlowMilligramPerDay:
		return "mg/d" 
	case MassFlowCentigramPerDay:
		return "cg/d" 
	case MassFlowDecigramPerDay:
		return "dg/d" 
	case MassFlowDecagramPerDay:
		return "dag/d" 
	case MassFlowHectogramPerDay:
		return "hg/d" 
	case MassFlowKilogramPerDay:
		return "kg/d" 
	case MassFlowMegagramPerDay:
		return "Mg/d" 
	case MassFlowMegapoundPerDay:
		return "Mlb/d" 
	case MassFlowMegapoundPerHour:
		return "Mlb/h" 
	case MassFlowMegapoundPerMinute:
		return "Mlb/min" 
	case MassFlowMegapoundPerSecond:
		return "Mlb/s" 
	default:
		return ""
	}
}

// Equals checks if the given MassFlow is equal to the current MassFlow.
//
// Parameters:
//    other: The MassFlow to compare against.
//
// Returns:
//    bool: Returns true if both MassFlow are equal, false otherwise.
func (a *MassFlow) Equals(other *MassFlow) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MassFlow with another MassFlow.
// It returns -1 if the current MassFlow is less than the other MassFlow, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MassFlow to compare against.
//
// Returns:
//    int: -1 if the current MassFlow is less, 1 if greater, and 0 if equal.
func (a *MassFlow) CompareTo(other *MassFlow) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MassFlow to the current MassFlow and returns the result.
// The result is a new MassFlow instance with the sum of the values.
//
// Parameters:
//    other: The MassFlow to add to the current MassFlow.
//
// Returns:
//    *MassFlow: A new MassFlow instance representing the sum of both MassFlow.
func (a *MassFlow) Add(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MassFlow from the current MassFlow and returns the result.
// The result is a new MassFlow instance with the difference of the values.
//
// Parameters:
//    other: The MassFlow to subtract from the current MassFlow.
//
// Returns:
//    *MassFlow: A new MassFlow instance representing the difference of both MassFlow.
func (a *MassFlow) Subtract(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MassFlow by the given MassFlow and returns the result.
// The result is a new MassFlow instance with the product of the values.
//
// Parameters:
//    other: The MassFlow to multiply with the current MassFlow.
//
// Returns:
//    *MassFlow: A new MassFlow instance representing the product of both MassFlow.
func (a *MassFlow) Multiply(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value * other.BaseValue()}
}

// Divide divides the current MassFlow by the given MassFlow and returns the result.
// The result is a new MassFlow instance with the quotient of the values.
//
// Parameters:
//    other: The MassFlow to divide the current MassFlow by.
//
// Returns:
//    *MassFlow: A new MassFlow instance representing the quotient of both MassFlow.
func (a *MassFlow) Divide(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value / other.BaseValue()}
}