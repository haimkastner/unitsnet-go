// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassFlowUnits enumeration
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

// MassFlowDto represents an MassFlow
type MassFlowDto struct {
	Value float64
	Unit  MassFlowUnits
}

// MassFlowDtoFactory struct to group related functions
type MassFlowDtoFactory struct{}

func (udf MassFlowDtoFactory) FromJSON(data []byte) (*MassFlowDto, error) {
	a := MassFlowDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MassFlow struct
type MassFlow struct {
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

// MassFlowFactory struct to group related functions
type MassFlowFactory struct{}

func (uf MassFlowFactory) CreateMassFlow(value float64, unit MassFlowUnits) (*MassFlow, error) {
	return newMassFlow(value, unit)
}

func (uf MassFlowFactory) FromDto(dto MassFlowDto) (*MassFlow, error) {
	return newMassFlow(dto.Value, dto.Unit)
}

func (uf MassFlowFactory) FromDtoJSON(data []byte) (*MassFlow, error) {
	unitDto, err := MassFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassFlowFactory{}.FromDto(*unitDto)
}


// FromGramPerSecond creates a new MassFlow instance from GramPerSecond.
func (uf MassFlowFactory) FromGramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerSecond)
}

// FromGramPerDay creates a new MassFlow instance from GramPerDay.
func (uf MassFlowFactory) FromGramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerDay)
}

// FromGramPerHour creates a new MassFlow instance from GramPerHour.
func (uf MassFlowFactory) FromGramsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowGramPerHour)
}

// FromKilogramPerHour creates a new MassFlow instance from KilogramPerHour.
func (uf MassFlowFactory) FromKilogramsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerHour)
}

// FromKilogramPerMinute creates a new MassFlow instance from KilogramPerMinute.
func (uf MassFlowFactory) FromKilogramsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerMinute)
}

// FromTonnePerHour creates a new MassFlow instance from TonnePerHour.
func (uf MassFlowFactory) FromTonnesPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowTonnePerHour)
}

// FromPoundPerDay creates a new MassFlow instance from PoundPerDay.
func (uf MassFlowFactory) FromPoundsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerDay)
}

// FromPoundPerHour creates a new MassFlow instance from PoundPerHour.
func (uf MassFlowFactory) FromPoundsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerHour)
}

// FromPoundPerMinute creates a new MassFlow instance from PoundPerMinute.
func (uf MassFlowFactory) FromPoundsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerMinute)
}

// FromPoundPerSecond creates a new MassFlow instance from PoundPerSecond.
func (uf MassFlowFactory) FromPoundsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowPoundPerSecond)
}

// FromTonnePerDay creates a new MassFlow instance from TonnePerDay.
func (uf MassFlowFactory) FromTonnesPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowTonnePerDay)
}

// FromShortTonPerHour creates a new MassFlow instance from ShortTonPerHour.
func (uf MassFlowFactory) FromShortTonsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowShortTonPerHour)
}

// FromNanogramPerSecond creates a new MassFlow instance from NanogramPerSecond.
func (uf MassFlowFactory) FromNanogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowNanogramPerSecond)
}

// FromMicrogramPerSecond creates a new MassFlow instance from MicrogramPerSecond.
func (uf MassFlowFactory) FromMicrogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMicrogramPerSecond)
}

// FromMilligramPerSecond creates a new MassFlow instance from MilligramPerSecond.
func (uf MassFlowFactory) FromMilligramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMilligramPerSecond)
}

// FromCentigramPerSecond creates a new MassFlow instance from CentigramPerSecond.
func (uf MassFlowFactory) FromCentigramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowCentigramPerSecond)
}

// FromDecigramPerSecond creates a new MassFlow instance from DecigramPerSecond.
func (uf MassFlowFactory) FromDecigramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecigramPerSecond)
}

// FromDecagramPerSecond creates a new MassFlow instance from DecagramPerSecond.
func (uf MassFlowFactory) FromDecagramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecagramPerSecond)
}

// FromHectogramPerSecond creates a new MassFlow instance from HectogramPerSecond.
func (uf MassFlowFactory) FromHectogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowHectogramPerSecond)
}

// FromKilogramPerSecond creates a new MassFlow instance from KilogramPerSecond.
func (uf MassFlowFactory) FromKilogramsPerSecond(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerSecond)
}

// FromNanogramPerDay creates a new MassFlow instance from NanogramPerDay.
func (uf MassFlowFactory) FromNanogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowNanogramPerDay)
}

// FromMicrogramPerDay creates a new MassFlow instance from MicrogramPerDay.
func (uf MassFlowFactory) FromMicrogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMicrogramPerDay)
}

// FromMilligramPerDay creates a new MassFlow instance from MilligramPerDay.
func (uf MassFlowFactory) FromMilligramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMilligramPerDay)
}

// FromCentigramPerDay creates a new MassFlow instance from CentigramPerDay.
func (uf MassFlowFactory) FromCentigramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowCentigramPerDay)
}

// FromDecigramPerDay creates a new MassFlow instance from DecigramPerDay.
func (uf MassFlowFactory) FromDecigramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecigramPerDay)
}

// FromDecagramPerDay creates a new MassFlow instance from DecagramPerDay.
func (uf MassFlowFactory) FromDecagramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowDecagramPerDay)
}

// FromHectogramPerDay creates a new MassFlow instance from HectogramPerDay.
func (uf MassFlowFactory) FromHectogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowHectogramPerDay)
}

// FromKilogramPerDay creates a new MassFlow instance from KilogramPerDay.
func (uf MassFlowFactory) FromKilogramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowKilogramPerDay)
}

// FromMegagramPerDay creates a new MassFlow instance from MegagramPerDay.
func (uf MassFlowFactory) FromMegagramsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegagramPerDay)
}

// FromMegapoundPerDay creates a new MassFlow instance from MegapoundPerDay.
func (uf MassFlowFactory) FromMegapoundsPerDay(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerDay)
}

// FromMegapoundPerHour creates a new MassFlow instance from MegapoundPerHour.
func (uf MassFlowFactory) FromMegapoundsPerHour(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerHour)
}

// FromMegapoundPerMinute creates a new MassFlow instance from MegapoundPerMinute.
func (uf MassFlowFactory) FromMegapoundsPerMinute(value float64) (*MassFlow, error) {
	return newMassFlow(value, MassFlowMegapoundPerMinute)
}

// FromMegapoundPerSecond creates a new MassFlow instance from MegapoundPerSecond.
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

// BaseValue returns the base value of MassFlow in GramPerSecond.
func (a *MassFlow) BaseValue() float64 {
	return a.value
}


// GramPerSecond returns the value in GramPerSecond.
func (a *MassFlow) GramsPerSecond() float64 {
	if a.grams_per_secondLazy != nil {
		return *a.grams_per_secondLazy
	}
	grams_per_second := a.convertFromBase(MassFlowGramPerSecond)
	a.grams_per_secondLazy = &grams_per_second
	return grams_per_second
}

// GramPerDay returns the value in GramPerDay.
func (a *MassFlow) GramsPerDay() float64 {
	if a.grams_per_dayLazy != nil {
		return *a.grams_per_dayLazy
	}
	grams_per_day := a.convertFromBase(MassFlowGramPerDay)
	a.grams_per_dayLazy = &grams_per_day
	return grams_per_day
}

// GramPerHour returns the value in GramPerHour.
func (a *MassFlow) GramsPerHour() float64 {
	if a.grams_per_hourLazy != nil {
		return *a.grams_per_hourLazy
	}
	grams_per_hour := a.convertFromBase(MassFlowGramPerHour)
	a.grams_per_hourLazy = &grams_per_hour
	return grams_per_hour
}

// KilogramPerHour returns the value in KilogramPerHour.
func (a *MassFlow) KilogramsPerHour() float64 {
	if a.kilograms_per_hourLazy != nil {
		return *a.kilograms_per_hourLazy
	}
	kilograms_per_hour := a.convertFromBase(MassFlowKilogramPerHour)
	a.kilograms_per_hourLazy = &kilograms_per_hour
	return kilograms_per_hour
}

// KilogramPerMinute returns the value in KilogramPerMinute.
func (a *MassFlow) KilogramsPerMinute() float64 {
	if a.kilograms_per_minuteLazy != nil {
		return *a.kilograms_per_minuteLazy
	}
	kilograms_per_minute := a.convertFromBase(MassFlowKilogramPerMinute)
	a.kilograms_per_minuteLazy = &kilograms_per_minute
	return kilograms_per_minute
}

// TonnePerHour returns the value in TonnePerHour.
func (a *MassFlow) TonnesPerHour() float64 {
	if a.tonnes_per_hourLazy != nil {
		return *a.tonnes_per_hourLazy
	}
	tonnes_per_hour := a.convertFromBase(MassFlowTonnePerHour)
	a.tonnes_per_hourLazy = &tonnes_per_hour
	return tonnes_per_hour
}

// PoundPerDay returns the value in PoundPerDay.
func (a *MassFlow) PoundsPerDay() float64 {
	if a.pounds_per_dayLazy != nil {
		return *a.pounds_per_dayLazy
	}
	pounds_per_day := a.convertFromBase(MassFlowPoundPerDay)
	a.pounds_per_dayLazy = &pounds_per_day
	return pounds_per_day
}

// PoundPerHour returns the value in PoundPerHour.
func (a *MassFlow) PoundsPerHour() float64 {
	if a.pounds_per_hourLazy != nil {
		return *a.pounds_per_hourLazy
	}
	pounds_per_hour := a.convertFromBase(MassFlowPoundPerHour)
	a.pounds_per_hourLazy = &pounds_per_hour
	return pounds_per_hour
}

// PoundPerMinute returns the value in PoundPerMinute.
func (a *MassFlow) PoundsPerMinute() float64 {
	if a.pounds_per_minuteLazy != nil {
		return *a.pounds_per_minuteLazy
	}
	pounds_per_minute := a.convertFromBase(MassFlowPoundPerMinute)
	a.pounds_per_minuteLazy = &pounds_per_minute
	return pounds_per_minute
}

// PoundPerSecond returns the value in PoundPerSecond.
func (a *MassFlow) PoundsPerSecond() float64 {
	if a.pounds_per_secondLazy != nil {
		return *a.pounds_per_secondLazy
	}
	pounds_per_second := a.convertFromBase(MassFlowPoundPerSecond)
	a.pounds_per_secondLazy = &pounds_per_second
	return pounds_per_second
}

// TonnePerDay returns the value in TonnePerDay.
func (a *MassFlow) TonnesPerDay() float64 {
	if a.tonnes_per_dayLazy != nil {
		return *a.tonnes_per_dayLazy
	}
	tonnes_per_day := a.convertFromBase(MassFlowTonnePerDay)
	a.tonnes_per_dayLazy = &tonnes_per_day
	return tonnes_per_day
}

// ShortTonPerHour returns the value in ShortTonPerHour.
func (a *MassFlow) ShortTonsPerHour() float64 {
	if a.short_tons_per_hourLazy != nil {
		return *a.short_tons_per_hourLazy
	}
	short_tons_per_hour := a.convertFromBase(MassFlowShortTonPerHour)
	a.short_tons_per_hourLazy = &short_tons_per_hour
	return short_tons_per_hour
}

// NanogramPerSecond returns the value in NanogramPerSecond.
func (a *MassFlow) NanogramsPerSecond() float64 {
	if a.nanograms_per_secondLazy != nil {
		return *a.nanograms_per_secondLazy
	}
	nanograms_per_second := a.convertFromBase(MassFlowNanogramPerSecond)
	a.nanograms_per_secondLazy = &nanograms_per_second
	return nanograms_per_second
}

// MicrogramPerSecond returns the value in MicrogramPerSecond.
func (a *MassFlow) MicrogramsPerSecond() float64 {
	if a.micrograms_per_secondLazy != nil {
		return *a.micrograms_per_secondLazy
	}
	micrograms_per_second := a.convertFromBase(MassFlowMicrogramPerSecond)
	a.micrograms_per_secondLazy = &micrograms_per_second
	return micrograms_per_second
}

// MilligramPerSecond returns the value in MilligramPerSecond.
func (a *MassFlow) MilligramsPerSecond() float64 {
	if a.milligrams_per_secondLazy != nil {
		return *a.milligrams_per_secondLazy
	}
	milligrams_per_second := a.convertFromBase(MassFlowMilligramPerSecond)
	a.milligrams_per_secondLazy = &milligrams_per_second
	return milligrams_per_second
}

// CentigramPerSecond returns the value in CentigramPerSecond.
func (a *MassFlow) CentigramsPerSecond() float64 {
	if a.centigrams_per_secondLazy != nil {
		return *a.centigrams_per_secondLazy
	}
	centigrams_per_second := a.convertFromBase(MassFlowCentigramPerSecond)
	a.centigrams_per_secondLazy = &centigrams_per_second
	return centigrams_per_second
}

// DecigramPerSecond returns the value in DecigramPerSecond.
func (a *MassFlow) DecigramsPerSecond() float64 {
	if a.decigrams_per_secondLazy != nil {
		return *a.decigrams_per_secondLazy
	}
	decigrams_per_second := a.convertFromBase(MassFlowDecigramPerSecond)
	a.decigrams_per_secondLazy = &decigrams_per_second
	return decigrams_per_second
}

// DecagramPerSecond returns the value in DecagramPerSecond.
func (a *MassFlow) DecagramsPerSecond() float64 {
	if a.decagrams_per_secondLazy != nil {
		return *a.decagrams_per_secondLazy
	}
	decagrams_per_second := a.convertFromBase(MassFlowDecagramPerSecond)
	a.decagrams_per_secondLazy = &decagrams_per_second
	return decagrams_per_second
}

// HectogramPerSecond returns the value in HectogramPerSecond.
func (a *MassFlow) HectogramsPerSecond() float64 {
	if a.hectograms_per_secondLazy != nil {
		return *a.hectograms_per_secondLazy
	}
	hectograms_per_second := a.convertFromBase(MassFlowHectogramPerSecond)
	a.hectograms_per_secondLazy = &hectograms_per_second
	return hectograms_per_second
}

// KilogramPerSecond returns the value in KilogramPerSecond.
func (a *MassFlow) KilogramsPerSecond() float64 {
	if a.kilograms_per_secondLazy != nil {
		return *a.kilograms_per_secondLazy
	}
	kilograms_per_second := a.convertFromBase(MassFlowKilogramPerSecond)
	a.kilograms_per_secondLazy = &kilograms_per_second
	return kilograms_per_second
}

// NanogramPerDay returns the value in NanogramPerDay.
func (a *MassFlow) NanogramsPerDay() float64 {
	if a.nanograms_per_dayLazy != nil {
		return *a.nanograms_per_dayLazy
	}
	nanograms_per_day := a.convertFromBase(MassFlowNanogramPerDay)
	a.nanograms_per_dayLazy = &nanograms_per_day
	return nanograms_per_day
}

// MicrogramPerDay returns the value in MicrogramPerDay.
func (a *MassFlow) MicrogramsPerDay() float64 {
	if a.micrograms_per_dayLazy != nil {
		return *a.micrograms_per_dayLazy
	}
	micrograms_per_day := a.convertFromBase(MassFlowMicrogramPerDay)
	a.micrograms_per_dayLazy = &micrograms_per_day
	return micrograms_per_day
}

// MilligramPerDay returns the value in MilligramPerDay.
func (a *MassFlow) MilligramsPerDay() float64 {
	if a.milligrams_per_dayLazy != nil {
		return *a.milligrams_per_dayLazy
	}
	milligrams_per_day := a.convertFromBase(MassFlowMilligramPerDay)
	a.milligrams_per_dayLazy = &milligrams_per_day
	return milligrams_per_day
}

// CentigramPerDay returns the value in CentigramPerDay.
func (a *MassFlow) CentigramsPerDay() float64 {
	if a.centigrams_per_dayLazy != nil {
		return *a.centigrams_per_dayLazy
	}
	centigrams_per_day := a.convertFromBase(MassFlowCentigramPerDay)
	a.centigrams_per_dayLazy = &centigrams_per_day
	return centigrams_per_day
}

// DecigramPerDay returns the value in DecigramPerDay.
func (a *MassFlow) DecigramsPerDay() float64 {
	if a.decigrams_per_dayLazy != nil {
		return *a.decigrams_per_dayLazy
	}
	decigrams_per_day := a.convertFromBase(MassFlowDecigramPerDay)
	a.decigrams_per_dayLazy = &decigrams_per_day
	return decigrams_per_day
}

// DecagramPerDay returns the value in DecagramPerDay.
func (a *MassFlow) DecagramsPerDay() float64 {
	if a.decagrams_per_dayLazy != nil {
		return *a.decagrams_per_dayLazy
	}
	decagrams_per_day := a.convertFromBase(MassFlowDecagramPerDay)
	a.decagrams_per_dayLazy = &decagrams_per_day
	return decagrams_per_day
}

// HectogramPerDay returns the value in HectogramPerDay.
func (a *MassFlow) HectogramsPerDay() float64 {
	if a.hectograms_per_dayLazy != nil {
		return *a.hectograms_per_dayLazy
	}
	hectograms_per_day := a.convertFromBase(MassFlowHectogramPerDay)
	a.hectograms_per_dayLazy = &hectograms_per_day
	return hectograms_per_day
}

// KilogramPerDay returns the value in KilogramPerDay.
func (a *MassFlow) KilogramsPerDay() float64 {
	if a.kilograms_per_dayLazy != nil {
		return *a.kilograms_per_dayLazy
	}
	kilograms_per_day := a.convertFromBase(MassFlowKilogramPerDay)
	a.kilograms_per_dayLazy = &kilograms_per_day
	return kilograms_per_day
}

// MegagramPerDay returns the value in MegagramPerDay.
func (a *MassFlow) MegagramsPerDay() float64 {
	if a.megagrams_per_dayLazy != nil {
		return *a.megagrams_per_dayLazy
	}
	megagrams_per_day := a.convertFromBase(MassFlowMegagramPerDay)
	a.megagrams_per_dayLazy = &megagrams_per_day
	return megagrams_per_day
}

// MegapoundPerDay returns the value in MegapoundPerDay.
func (a *MassFlow) MegapoundsPerDay() float64 {
	if a.megapounds_per_dayLazy != nil {
		return *a.megapounds_per_dayLazy
	}
	megapounds_per_day := a.convertFromBase(MassFlowMegapoundPerDay)
	a.megapounds_per_dayLazy = &megapounds_per_day
	return megapounds_per_day
}

// MegapoundPerHour returns the value in MegapoundPerHour.
func (a *MassFlow) MegapoundsPerHour() float64 {
	if a.megapounds_per_hourLazy != nil {
		return *a.megapounds_per_hourLazy
	}
	megapounds_per_hour := a.convertFromBase(MassFlowMegapoundPerHour)
	a.megapounds_per_hourLazy = &megapounds_per_hour
	return megapounds_per_hour
}

// MegapoundPerMinute returns the value in MegapoundPerMinute.
func (a *MassFlow) MegapoundsPerMinute() float64 {
	if a.megapounds_per_minuteLazy != nil {
		return *a.megapounds_per_minuteLazy
	}
	megapounds_per_minute := a.convertFromBase(MassFlowMegapoundPerMinute)
	a.megapounds_per_minuteLazy = &megapounds_per_minute
	return megapounds_per_minute
}

// MegapoundPerSecond returns the value in MegapoundPerSecond.
func (a *MassFlow) MegapoundsPerSecond() float64 {
	if a.megapounds_per_secondLazy != nil {
		return *a.megapounds_per_secondLazy
	}
	megapounds_per_second := a.convertFromBase(MassFlowMegapoundPerSecond)
	a.megapounds_per_secondLazy = &megapounds_per_second
	return megapounds_per_second
}


// ToDto creates an MassFlowDto representation.
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

// ToDtoJSON creates an MassFlowDto representation.
func (a *MassFlow) ToDtoJSON(holdInUnit *MassFlowUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MassFlow to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a MassFlow) String() string {
	return a.ToString(MassFlowGramPerSecond, 2)
}

// ToString formats the MassFlow to string.
// fractionalDigits -1 for not mention
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

// Check if the given MassFlow are equals to the current MassFlow
func (a *MassFlow) Equals(other *MassFlow) bool {
	return a.value == other.BaseValue()
}

// Check if the given MassFlow are equals to the current MassFlow
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

// Add the given MassFlow to the current MassFlow.
func (a *MassFlow) Add(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value + other.BaseValue()}
}

// Subtract the given MassFlow to the current MassFlow.
func (a *MassFlow) Subtract(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value - other.BaseValue()}
}

// Multiply the given MassFlow to the current MassFlow.
func (a *MassFlow) Multiply(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value * other.BaseValue()}
}

// Divide the given MassFlow to the current MassFlow.
func (a *MassFlow) Divide(other *MassFlow) *MassFlow {
	return &MassFlow{value: a.value / other.BaseValue()}
}