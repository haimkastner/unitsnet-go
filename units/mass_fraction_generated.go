// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassFractionUnits defines various units of MassFraction.
type MassFractionUnits string

const (
    
        // 
        MassFractionDecimalFraction MassFractionUnits = "DecimalFraction"
        // 
        MassFractionGramPerGram MassFractionUnits = "GramPerGram"
        // 
        MassFractionGramPerKilogram MassFractionUnits = "GramPerKilogram"
        // 
        MassFractionPercent MassFractionUnits = "Percent"
        // 
        MassFractionPartPerThousand MassFractionUnits = "PartPerThousand"
        // 
        MassFractionPartPerMillion MassFractionUnits = "PartPerMillion"
        // 
        MassFractionPartPerBillion MassFractionUnits = "PartPerBillion"
        // 
        MassFractionPartPerTrillion MassFractionUnits = "PartPerTrillion"
        // 
        MassFractionNanogramPerGram MassFractionUnits = "NanogramPerGram"
        // 
        MassFractionMicrogramPerGram MassFractionUnits = "MicrogramPerGram"
        // 
        MassFractionMilligramPerGram MassFractionUnits = "MilligramPerGram"
        // 
        MassFractionCentigramPerGram MassFractionUnits = "CentigramPerGram"
        // 
        MassFractionDecigramPerGram MassFractionUnits = "DecigramPerGram"
        // 
        MassFractionDecagramPerGram MassFractionUnits = "DecagramPerGram"
        // 
        MassFractionHectogramPerGram MassFractionUnits = "HectogramPerGram"
        // 
        MassFractionKilogramPerGram MassFractionUnits = "KilogramPerGram"
        // 
        MassFractionNanogramPerKilogram MassFractionUnits = "NanogramPerKilogram"
        // 
        MassFractionMicrogramPerKilogram MassFractionUnits = "MicrogramPerKilogram"
        // 
        MassFractionMilligramPerKilogram MassFractionUnits = "MilligramPerKilogram"
        // 
        MassFractionCentigramPerKilogram MassFractionUnits = "CentigramPerKilogram"
        // 
        MassFractionDecigramPerKilogram MassFractionUnits = "DecigramPerKilogram"
        // 
        MassFractionDecagramPerKilogram MassFractionUnits = "DecagramPerKilogram"
        // 
        MassFractionHectogramPerKilogram MassFractionUnits = "HectogramPerKilogram"
        // 
        MassFractionKilogramPerKilogram MassFractionUnits = "KilogramPerKilogram"
)

// MassFractionDto represents a MassFraction measurement with a numerical value and its corresponding unit.
type MassFractionDto struct {
    // Value is the numerical representation of the MassFraction.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MassFraction, as defined in the MassFractionUnits enumeration.
	Unit  MassFractionUnits `json:"unit"`
}

// MassFractionDtoFactory groups methods for creating and serializing MassFractionDto objects.
type MassFractionDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassFractionDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassFractionDtoFactory) FromJSON(data []byte) (*MassFractionDto, error) {
	a := MassFractionDto{}

    // Parse JSON into MassFractionDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MassFractionDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassFractionDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MassFraction represents a measurement in a of MassFraction.
//
// The mass fraction is defined as the mass of a constituent divided by the total mass of the mixture.
type MassFraction struct {
	// value is the base measurement stored internally.
	value       float64
    
    decimal_fractionsLazy *float64 
    grams_per_gramLazy *float64 
    grams_per_kilogramLazy *float64 
    percentLazy *float64 
    parts_per_thousandLazy *float64 
    parts_per_millionLazy *float64 
    parts_per_billionLazy *float64 
    parts_per_trillionLazy *float64 
    nanograms_per_gramLazy *float64 
    micrograms_per_gramLazy *float64 
    milligrams_per_gramLazy *float64 
    centigrams_per_gramLazy *float64 
    decigrams_per_gramLazy *float64 
    decagrams_per_gramLazy *float64 
    hectograms_per_gramLazy *float64 
    kilograms_per_gramLazy *float64 
    nanograms_per_kilogramLazy *float64 
    micrograms_per_kilogramLazy *float64 
    milligrams_per_kilogramLazy *float64 
    centigrams_per_kilogramLazy *float64 
    decigrams_per_kilogramLazy *float64 
    decagrams_per_kilogramLazy *float64 
    hectograms_per_kilogramLazy *float64 
    kilograms_per_kilogramLazy *float64 
}

// MassFractionFactory groups methods for creating MassFraction instances.
type MassFractionFactory struct{}

// CreateMassFraction creates a new MassFraction instance from the given value and unit.
func (uf MassFractionFactory) CreateMassFraction(value float64, unit MassFractionUnits) (*MassFraction, error) {
	return newMassFraction(value, unit)
}

// FromDto converts a MassFractionDto to a MassFraction instance.
func (uf MassFractionFactory) FromDto(dto MassFractionDto) (*MassFraction, error) {
	return newMassFraction(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MassFraction instance.
func (uf MassFractionFactory) FromDtoJSON(data []byte) (*MassFraction, error) {
	unitDto, err := MassFractionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassFractionDto from JSON: %w", err)
	}
	return MassFractionFactory{}.FromDto(*unitDto)
}


// FromDecimalFractions creates a new MassFraction instance from a value in DecimalFractions.
func (uf MassFractionFactory) FromDecimalFractions(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecimalFraction)
}

// FromGramsPerGram creates a new MassFraction instance from a value in GramsPerGram.
func (uf MassFractionFactory) FromGramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionGramPerGram)
}

// FromGramsPerKilogram creates a new MassFraction instance from a value in GramsPerKilogram.
func (uf MassFractionFactory) FromGramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionGramPerKilogram)
}

// FromPercent creates a new MassFraction instance from a value in Percent.
func (uf MassFractionFactory) FromPercent(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPercent)
}

// FromPartsPerThousand creates a new MassFraction instance from a value in PartsPerThousand.
func (uf MassFractionFactory) FromPartsPerThousand(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerThousand)
}

// FromPartsPerMillion creates a new MassFraction instance from a value in PartsPerMillion.
func (uf MassFractionFactory) FromPartsPerMillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerMillion)
}

// FromPartsPerBillion creates a new MassFraction instance from a value in PartsPerBillion.
func (uf MassFractionFactory) FromPartsPerBillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerBillion)
}

// FromPartsPerTrillion creates a new MassFraction instance from a value in PartsPerTrillion.
func (uf MassFractionFactory) FromPartsPerTrillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerTrillion)
}

// FromNanogramsPerGram creates a new MassFraction instance from a value in NanogramsPerGram.
func (uf MassFractionFactory) FromNanogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionNanogramPerGram)
}

// FromMicrogramsPerGram creates a new MassFraction instance from a value in MicrogramsPerGram.
func (uf MassFractionFactory) FromMicrogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMicrogramPerGram)
}

// FromMilligramsPerGram creates a new MassFraction instance from a value in MilligramsPerGram.
func (uf MassFractionFactory) FromMilligramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMilligramPerGram)
}

// FromCentigramsPerGram creates a new MassFraction instance from a value in CentigramsPerGram.
func (uf MassFractionFactory) FromCentigramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionCentigramPerGram)
}

// FromDecigramsPerGram creates a new MassFraction instance from a value in DecigramsPerGram.
func (uf MassFractionFactory) FromDecigramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecigramPerGram)
}

// FromDecagramsPerGram creates a new MassFraction instance from a value in DecagramsPerGram.
func (uf MassFractionFactory) FromDecagramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecagramPerGram)
}

// FromHectogramsPerGram creates a new MassFraction instance from a value in HectogramsPerGram.
func (uf MassFractionFactory) FromHectogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionHectogramPerGram)
}

// FromKilogramsPerGram creates a new MassFraction instance from a value in KilogramsPerGram.
func (uf MassFractionFactory) FromKilogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionKilogramPerGram)
}

// FromNanogramsPerKilogram creates a new MassFraction instance from a value in NanogramsPerKilogram.
func (uf MassFractionFactory) FromNanogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionNanogramPerKilogram)
}

// FromMicrogramsPerKilogram creates a new MassFraction instance from a value in MicrogramsPerKilogram.
func (uf MassFractionFactory) FromMicrogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMicrogramPerKilogram)
}

// FromMilligramsPerKilogram creates a new MassFraction instance from a value in MilligramsPerKilogram.
func (uf MassFractionFactory) FromMilligramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMilligramPerKilogram)
}

// FromCentigramsPerKilogram creates a new MassFraction instance from a value in CentigramsPerKilogram.
func (uf MassFractionFactory) FromCentigramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionCentigramPerKilogram)
}

// FromDecigramsPerKilogram creates a new MassFraction instance from a value in DecigramsPerKilogram.
func (uf MassFractionFactory) FromDecigramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecigramPerKilogram)
}

// FromDecagramsPerKilogram creates a new MassFraction instance from a value in DecagramsPerKilogram.
func (uf MassFractionFactory) FromDecagramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecagramPerKilogram)
}

// FromHectogramsPerKilogram creates a new MassFraction instance from a value in HectogramsPerKilogram.
func (uf MassFractionFactory) FromHectogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionHectogramPerKilogram)
}

// FromKilogramsPerKilogram creates a new MassFraction instance from a value in KilogramsPerKilogram.
func (uf MassFractionFactory) FromKilogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionKilogramPerKilogram)
}


// newMassFraction creates a new MassFraction.
func newMassFraction(value float64, fromUnit MassFractionUnits) (*MassFraction, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MassFraction{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassFraction in DecimalFraction unit (the base/default quantity).
func (a *MassFraction) BaseValue() float64 {
	return a.value
}


// DecimalFractions returns the MassFraction value in DecimalFractions.
//
// 
func (a *MassFraction) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(MassFractionDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// GramsPerGram returns the MassFraction value in GramsPerGram.
//
// 
func (a *MassFraction) GramsPerGram() float64 {
	if a.grams_per_gramLazy != nil {
		return *a.grams_per_gramLazy
	}
	grams_per_gram := a.convertFromBase(MassFractionGramPerGram)
	a.grams_per_gramLazy = &grams_per_gram
	return grams_per_gram
}

// GramsPerKilogram returns the MassFraction value in GramsPerKilogram.
//
// 
func (a *MassFraction) GramsPerKilogram() float64 {
	if a.grams_per_kilogramLazy != nil {
		return *a.grams_per_kilogramLazy
	}
	grams_per_kilogram := a.convertFromBase(MassFractionGramPerKilogram)
	a.grams_per_kilogramLazy = &grams_per_kilogram
	return grams_per_kilogram
}

// Percent returns the MassFraction value in Percent.
//
// 
func (a *MassFraction) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(MassFractionPercent)
	a.percentLazy = &percent
	return percent
}

// PartsPerThousand returns the MassFraction value in PartsPerThousand.
//
// 
func (a *MassFraction) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(MassFractionPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartsPerMillion returns the MassFraction value in PartsPerMillion.
//
// 
func (a *MassFraction) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(MassFractionPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartsPerBillion returns the MassFraction value in PartsPerBillion.
//
// 
func (a *MassFraction) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(MassFractionPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartsPerTrillion returns the MassFraction value in PartsPerTrillion.
//
// 
func (a *MassFraction) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(MassFractionPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}

// NanogramsPerGram returns the MassFraction value in NanogramsPerGram.
//
// 
func (a *MassFraction) NanogramsPerGram() float64 {
	if a.nanograms_per_gramLazy != nil {
		return *a.nanograms_per_gramLazy
	}
	nanograms_per_gram := a.convertFromBase(MassFractionNanogramPerGram)
	a.nanograms_per_gramLazy = &nanograms_per_gram
	return nanograms_per_gram
}

// MicrogramsPerGram returns the MassFraction value in MicrogramsPerGram.
//
// 
func (a *MassFraction) MicrogramsPerGram() float64 {
	if a.micrograms_per_gramLazy != nil {
		return *a.micrograms_per_gramLazy
	}
	micrograms_per_gram := a.convertFromBase(MassFractionMicrogramPerGram)
	a.micrograms_per_gramLazy = &micrograms_per_gram
	return micrograms_per_gram
}

// MilligramsPerGram returns the MassFraction value in MilligramsPerGram.
//
// 
func (a *MassFraction) MilligramsPerGram() float64 {
	if a.milligrams_per_gramLazy != nil {
		return *a.milligrams_per_gramLazy
	}
	milligrams_per_gram := a.convertFromBase(MassFractionMilligramPerGram)
	a.milligrams_per_gramLazy = &milligrams_per_gram
	return milligrams_per_gram
}

// CentigramsPerGram returns the MassFraction value in CentigramsPerGram.
//
// 
func (a *MassFraction) CentigramsPerGram() float64 {
	if a.centigrams_per_gramLazy != nil {
		return *a.centigrams_per_gramLazy
	}
	centigrams_per_gram := a.convertFromBase(MassFractionCentigramPerGram)
	a.centigrams_per_gramLazy = &centigrams_per_gram
	return centigrams_per_gram
}

// DecigramsPerGram returns the MassFraction value in DecigramsPerGram.
//
// 
func (a *MassFraction) DecigramsPerGram() float64 {
	if a.decigrams_per_gramLazy != nil {
		return *a.decigrams_per_gramLazy
	}
	decigrams_per_gram := a.convertFromBase(MassFractionDecigramPerGram)
	a.decigrams_per_gramLazy = &decigrams_per_gram
	return decigrams_per_gram
}

// DecagramsPerGram returns the MassFraction value in DecagramsPerGram.
//
// 
func (a *MassFraction) DecagramsPerGram() float64 {
	if a.decagrams_per_gramLazy != nil {
		return *a.decagrams_per_gramLazy
	}
	decagrams_per_gram := a.convertFromBase(MassFractionDecagramPerGram)
	a.decagrams_per_gramLazy = &decagrams_per_gram
	return decagrams_per_gram
}

// HectogramsPerGram returns the MassFraction value in HectogramsPerGram.
//
// 
func (a *MassFraction) HectogramsPerGram() float64 {
	if a.hectograms_per_gramLazy != nil {
		return *a.hectograms_per_gramLazy
	}
	hectograms_per_gram := a.convertFromBase(MassFractionHectogramPerGram)
	a.hectograms_per_gramLazy = &hectograms_per_gram
	return hectograms_per_gram
}

// KilogramsPerGram returns the MassFraction value in KilogramsPerGram.
//
// 
func (a *MassFraction) KilogramsPerGram() float64 {
	if a.kilograms_per_gramLazy != nil {
		return *a.kilograms_per_gramLazy
	}
	kilograms_per_gram := a.convertFromBase(MassFractionKilogramPerGram)
	a.kilograms_per_gramLazy = &kilograms_per_gram
	return kilograms_per_gram
}

// NanogramsPerKilogram returns the MassFraction value in NanogramsPerKilogram.
//
// 
func (a *MassFraction) NanogramsPerKilogram() float64 {
	if a.nanograms_per_kilogramLazy != nil {
		return *a.nanograms_per_kilogramLazy
	}
	nanograms_per_kilogram := a.convertFromBase(MassFractionNanogramPerKilogram)
	a.nanograms_per_kilogramLazy = &nanograms_per_kilogram
	return nanograms_per_kilogram
}

// MicrogramsPerKilogram returns the MassFraction value in MicrogramsPerKilogram.
//
// 
func (a *MassFraction) MicrogramsPerKilogram() float64 {
	if a.micrograms_per_kilogramLazy != nil {
		return *a.micrograms_per_kilogramLazy
	}
	micrograms_per_kilogram := a.convertFromBase(MassFractionMicrogramPerKilogram)
	a.micrograms_per_kilogramLazy = &micrograms_per_kilogram
	return micrograms_per_kilogram
}

// MilligramsPerKilogram returns the MassFraction value in MilligramsPerKilogram.
//
// 
func (a *MassFraction) MilligramsPerKilogram() float64 {
	if a.milligrams_per_kilogramLazy != nil {
		return *a.milligrams_per_kilogramLazy
	}
	milligrams_per_kilogram := a.convertFromBase(MassFractionMilligramPerKilogram)
	a.milligrams_per_kilogramLazy = &milligrams_per_kilogram
	return milligrams_per_kilogram
}

// CentigramsPerKilogram returns the MassFraction value in CentigramsPerKilogram.
//
// 
func (a *MassFraction) CentigramsPerKilogram() float64 {
	if a.centigrams_per_kilogramLazy != nil {
		return *a.centigrams_per_kilogramLazy
	}
	centigrams_per_kilogram := a.convertFromBase(MassFractionCentigramPerKilogram)
	a.centigrams_per_kilogramLazy = &centigrams_per_kilogram
	return centigrams_per_kilogram
}

// DecigramsPerKilogram returns the MassFraction value in DecigramsPerKilogram.
//
// 
func (a *MassFraction) DecigramsPerKilogram() float64 {
	if a.decigrams_per_kilogramLazy != nil {
		return *a.decigrams_per_kilogramLazy
	}
	decigrams_per_kilogram := a.convertFromBase(MassFractionDecigramPerKilogram)
	a.decigrams_per_kilogramLazy = &decigrams_per_kilogram
	return decigrams_per_kilogram
}

// DecagramsPerKilogram returns the MassFraction value in DecagramsPerKilogram.
//
// 
func (a *MassFraction) DecagramsPerKilogram() float64 {
	if a.decagrams_per_kilogramLazy != nil {
		return *a.decagrams_per_kilogramLazy
	}
	decagrams_per_kilogram := a.convertFromBase(MassFractionDecagramPerKilogram)
	a.decagrams_per_kilogramLazy = &decagrams_per_kilogram
	return decagrams_per_kilogram
}

// HectogramsPerKilogram returns the MassFraction value in HectogramsPerKilogram.
//
// 
func (a *MassFraction) HectogramsPerKilogram() float64 {
	if a.hectograms_per_kilogramLazy != nil {
		return *a.hectograms_per_kilogramLazy
	}
	hectograms_per_kilogram := a.convertFromBase(MassFractionHectogramPerKilogram)
	a.hectograms_per_kilogramLazy = &hectograms_per_kilogram
	return hectograms_per_kilogram
}

// KilogramsPerKilogram returns the MassFraction value in KilogramsPerKilogram.
//
// 
func (a *MassFraction) KilogramsPerKilogram() float64 {
	if a.kilograms_per_kilogramLazy != nil {
		return *a.kilograms_per_kilogramLazy
	}
	kilograms_per_kilogram := a.convertFromBase(MassFractionKilogramPerKilogram)
	a.kilograms_per_kilogramLazy = &kilograms_per_kilogram
	return kilograms_per_kilogram
}


// ToDto creates a MassFractionDto representation from the MassFraction instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *MassFraction) ToDto(holdInUnit *MassFractionUnits) MassFractionDto {
	if holdInUnit == nil {
		defaultUnit := MassFractionDecimalFraction // Default value
		holdInUnit = &defaultUnit
	}

	return MassFractionDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MassFraction instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *MassFraction) ToDtoJSON(holdInUnit *MassFractionUnits) ([]byte, error) {
	// Convert to MassFractionDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MassFraction to a specific unit value.
// The function uses the provided unit type (MassFractionUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MassFraction) Convert(toUnit MassFractionUnits) float64 {
	switch toUnit { 
    case MassFractionDecimalFraction:
		return a.DecimalFractions()
    case MassFractionGramPerGram:
		return a.GramsPerGram()
    case MassFractionGramPerKilogram:
		return a.GramsPerKilogram()
    case MassFractionPercent:
		return a.Percent()
    case MassFractionPartPerThousand:
		return a.PartsPerThousand()
    case MassFractionPartPerMillion:
		return a.PartsPerMillion()
    case MassFractionPartPerBillion:
		return a.PartsPerBillion()
    case MassFractionPartPerTrillion:
		return a.PartsPerTrillion()
    case MassFractionNanogramPerGram:
		return a.NanogramsPerGram()
    case MassFractionMicrogramPerGram:
		return a.MicrogramsPerGram()
    case MassFractionMilligramPerGram:
		return a.MilligramsPerGram()
    case MassFractionCentigramPerGram:
		return a.CentigramsPerGram()
    case MassFractionDecigramPerGram:
		return a.DecigramsPerGram()
    case MassFractionDecagramPerGram:
		return a.DecagramsPerGram()
    case MassFractionHectogramPerGram:
		return a.HectogramsPerGram()
    case MassFractionKilogramPerGram:
		return a.KilogramsPerGram()
    case MassFractionNanogramPerKilogram:
		return a.NanogramsPerKilogram()
    case MassFractionMicrogramPerKilogram:
		return a.MicrogramsPerKilogram()
    case MassFractionMilligramPerKilogram:
		return a.MilligramsPerKilogram()
    case MassFractionCentigramPerKilogram:
		return a.CentigramsPerKilogram()
    case MassFractionDecigramPerKilogram:
		return a.DecigramsPerKilogram()
    case MassFractionDecagramPerKilogram:
		return a.DecagramsPerKilogram()
    case MassFractionHectogramPerKilogram:
		return a.HectogramsPerKilogram()
    case MassFractionKilogramPerKilogram:
		return a.KilogramsPerKilogram()
	default:
		return math.NaN()
	}
}

func (a *MassFraction) convertFromBase(toUnit MassFractionUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassFractionDecimalFraction:
		return (value) 
	case MassFractionGramPerGram:
		return (value) 
	case MassFractionGramPerKilogram:
		return (value * 1e3) 
	case MassFractionPercent:
		return (value * 1e2) 
	case MassFractionPartPerThousand:
		return (value * 1e3) 
	case MassFractionPartPerMillion:
		return (value * 1e6) 
	case MassFractionPartPerBillion:
		return (value * 1e9) 
	case MassFractionPartPerTrillion:
		return (value * 1e12) 
	case MassFractionNanogramPerGram:
		return ((value) / 1e-09) 
	case MassFractionMicrogramPerGram:
		return ((value) / 1e-06) 
	case MassFractionMilligramPerGram:
		return ((value) / 0.001) 
	case MassFractionCentigramPerGram:
		return ((value) / 0.01) 
	case MassFractionDecigramPerGram:
		return ((value) / 0.1) 
	case MassFractionDecagramPerGram:
		return ((value) / 10.0) 
	case MassFractionHectogramPerGram:
		return ((value) / 100.0) 
	case MassFractionKilogramPerGram:
		return ((value) / 1000.0) 
	case MassFractionNanogramPerKilogram:
		return ((value * 1e3) / 1e-09) 
	case MassFractionMicrogramPerKilogram:
		return ((value * 1e3) / 1e-06) 
	case MassFractionMilligramPerKilogram:
		return ((value * 1e3) / 0.001) 
	case MassFractionCentigramPerKilogram:
		return ((value * 1e3) / 0.01) 
	case MassFractionDecigramPerKilogram:
		return ((value * 1e3) / 0.1) 
	case MassFractionDecagramPerKilogram:
		return ((value * 1e3) / 10.0) 
	case MassFractionHectogramPerKilogram:
		return ((value * 1e3) / 100.0) 
	case MassFractionKilogramPerKilogram:
		return ((value * 1e3) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *MassFraction) convertToBase(value float64, fromUnit MassFractionUnits) float64 {
	switch fromUnit { 
	case MassFractionDecimalFraction:
		return (value) 
	case MassFractionGramPerGram:
		return (value) 
	case MassFractionGramPerKilogram:
		return (value / 1e3) 
	case MassFractionPercent:
		return (value / 1e2) 
	case MassFractionPartPerThousand:
		return (value / 1e3) 
	case MassFractionPartPerMillion:
		return (value / 1e6) 
	case MassFractionPartPerBillion:
		return (value / 1e9) 
	case MassFractionPartPerTrillion:
		return (value / 1e12) 
	case MassFractionNanogramPerGram:
		return ((value) * 1e-09) 
	case MassFractionMicrogramPerGram:
		return ((value) * 1e-06) 
	case MassFractionMilligramPerGram:
		return ((value) * 0.001) 
	case MassFractionCentigramPerGram:
		return ((value) * 0.01) 
	case MassFractionDecigramPerGram:
		return ((value) * 0.1) 
	case MassFractionDecagramPerGram:
		return ((value) * 10.0) 
	case MassFractionHectogramPerGram:
		return ((value) * 100.0) 
	case MassFractionKilogramPerGram:
		return ((value) * 1000.0) 
	case MassFractionNanogramPerKilogram:
		return ((value / 1e3) * 1e-09) 
	case MassFractionMicrogramPerKilogram:
		return ((value / 1e3) * 1e-06) 
	case MassFractionMilligramPerKilogram:
		return ((value / 1e3) * 0.001) 
	case MassFractionCentigramPerKilogram:
		return ((value / 1e3) * 0.01) 
	case MassFractionDecigramPerKilogram:
		return ((value / 1e3) * 0.1) 
	case MassFractionDecagramPerKilogram:
		return ((value / 1e3) * 10.0) 
	case MassFractionHectogramPerKilogram:
		return ((value / 1e3) * 100.0) 
	case MassFractionKilogramPerKilogram:
		return ((value / 1e3) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MassFraction in the default unit (DecimalFraction),
// formatted to two decimal places.
func (a MassFraction) String() string {
	return a.ToString(MassFractionDecimalFraction, 2)
}

// ToString formats the MassFraction value as a string with the specified unit and fractional digits.
// It converts the MassFraction to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MassFraction value will be converted (e.g., DecimalFraction))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MassFraction with the unit abbreviation.
func (a *MassFraction) ToString(unit MassFractionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMassFractionAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMassFractionAbbreviation(unit))
}

// Equals checks if the given MassFraction is equal to the current MassFraction.
//
// Parameters:
//    other: The MassFraction to compare against.
//
// Returns:
//    bool: Returns true if both MassFraction are equal, false otherwise.
func (a *MassFraction) Equals(other *MassFraction) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MassFraction with another MassFraction.
// It returns -1 if the current MassFraction is less than the other MassFraction, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MassFraction to compare against.
//
// Returns:
//    int: -1 if the current MassFraction is less, 1 if greater, and 0 if equal.
func (a *MassFraction) CompareTo(other *MassFraction) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MassFraction to the current MassFraction and returns the result.
// The result is a new MassFraction instance with the sum of the values.
//
// Parameters:
//    other: The MassFraction to add to the current MassFraction.
//
// Returns:
//    *MassFraction: A new MassFraction instance representing the sum of both MassFraction.
func (a *MassFraction) Add(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MassFraction from the current MassFraction and returns the result.
// The result is a new MassFraction instance with the difference of the values.
//
// Parameters:
//    other: The MassFraction to subtract from the current MassFraction.
//
// Returns:
//    *MassFraction: A new MassFraction instance representing the difference of both MassFraction.
func (a *MassFraction) Subtract(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MassFraction by the given MassFraction and returns the result.
// The result is a new MassFraction instance with the product of the values.
//
// Parameters:
//    other: The MassFraction to multiply with the current MassFraction.
//
// Returns:
//    *MassFraction: A new MassFraction instance representing the product of both MassFraction.
func (a *MassFraction) Multiply(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value * other.BaseValue()}
}

// Divide divides the current MassFraction by the given MassFraction and returns the result.
// The result is a new MassFraction instance with the quotient of the values.
//
// Parameters:
//    other: The MassFraction to divide the current MassFraction by.
//
// Returns:
//    *MassFraction: A new MassFraction instance representing the quotient of both MassFraction.
func (a *MassFraction) Divide(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value / other.BaseValue()}
}

// GetMassFractionAbbreviation gets the unit abbreviation.
func GetMassFractionAbbreviation(unit MassFractionUnits) string {
	switch unit { 
	case MassFractionDecimalFraction:
		return "" 
	case MassFractionGramPerGram:
		return "g/g" 
	case MassFractionGramPerKilogram:
		return "g/kg" 
	case MassFractionPercent:
		return "%" 
	case MassFractionPartPerThousand:
		return "‰" 
	case MassFractionPartPerMillion:
		return "ppm" 
	case MassFractionPartPerBillion:
		return "ppb" 
	case MassFractionPartPerTrillion:
		return "ppt" 
	case MassFractionNanogramPerGram:
		return "ng/g" 
	case MassFractionMicrogramPerGram:
		return "μg/g" 
	case MassFractionMilligramPerGram:
		return "mg/g" 
	case MassFractionCentigramPerGram:
		return "cg/g" 
	case MassFractionDecigramPerGram:
		return "dg/g" 
	case MassFractionDecagramPerGram:
		return "dag/g" 
	case MassFractionHectogramPerGram:
		return "hg/g" 
	case MassFractionKilogramPerGram:
		return "kg/g" 
	case MassFractionNanogramPerKilogram:
		return "ng/kg" 
	case MassFractionMicrogramPerKilogram:
		return "μg/kg" 
	case MassFractionMilligramPerKilogram:
		return "mg/kg" 
	case MassFractionCentigramPerKilogram:
		return "cg/kg" 
	case MassFractionDecigramPerKilogram:
		return "dg/kg" 
	case MassFractionDecagramPerKilogram:
		return "dag/kg" 
	case MassFractionHectogramPerKilogram:
		return "hg/kg" 
	case MassFractionKilogramPerKilogram:
		return "kg/kg" 
	default:
		return ""
	}
}