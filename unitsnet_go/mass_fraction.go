package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassFractionUnits enumeration
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

// MassFractionDto represents an MassFraction
type MassFractionDto struct {
	Value float64
	Unit  MassFractionUnits
}

// MassFractionDtoFactory struct to group related functions
type MassFractionDtoFactory struct{}

func (udf MassFractionDtoFactory) FromJSON(data []byte) (*MassFractionDto, error) {
	a := MassFractionDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassFractionDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MassFraction struct
type MassFraction struct {
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

// MassFractionFactory struct to group related functions
type MassFractionFactory struct{}

func (uf MassFractionFactory) CreateMassFraction(value float64, unit MassFractionUnits) (*MassFraction, error) {
	return newMassFraction(value, unit)
}

func (uf MassFractionFactory) FromDto(dto MassFractionDto) (*MassFraction, error) {
	return newMassFraction(dto.Value, dto.Unit)
}

func (uf MassFractionFactory) FromDtoJSON(data []byte) (*MassFraction, error) {
	unitDto, err := MassFractionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassFractionFactory{}.FromDto(*unitDto)
}


// FromDecimalFraction creates a new MassFraction instance from DecimalFraction.
func (uf MassFractionFactory) FromDecimalFractions(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecimalFraction)
}

// FromGramPerGram creates a new MassFraction instance from GramPerGram.
func (uf MassFractionFactory) FromGramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionGramPerGram)
}

// FromGramPerKilogram creates a new MassFraction instance from GramPerKilogram.
func (uf MassFractionFactory) FromGramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionGramPerKilogram)
}

// FromPercent creates a new MassFraction instance from Percent.
func (uf MassFractionFactory) FromPercent(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPercent)
}

// FromPartPerThousand creates a new MassFraction instance from PartPerThousand.
func (uf MassFractionFactory) FromPartsPerThousand(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerThousand)
}

// FromPartPerMillion creates a new MassFraction instance from PartPerMillion.
func (uf MassFractionFactory) FromPartsPerMillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerMillion)
}

// FromPartPerBillion creates a new MassFraction instance from PartPerBillion.
func (uf MassFractionFactory) FromPartsPerBillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerBillion)
}

// FromPartPerTrillion creates a new MassFraction instance from PartPerTrillion.
func (uf MassFractionFactory) FromPartsPerTrillion(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionPartPerTrillion)
}

// FromNanogramPerGram creates a new MassFraction instance from NanogramPerGram.
func (uf MassFractionFactory) FromNanogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionNanogramPerGram)
}

// FromMicrogramPerGram creates a new MassFraction instance from MicrogramPerGram.
func (uf MassFractionFactory) FromMicrogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMicrogramPerGram)
}

// FromMilligramPerGram creates a new MassFraction instance from MilligramPerGram.
func (uf MassFractionFactory) FromMilligramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMilligramPerGram)
}

// FromCentigramPerGram creates a new MassFraction instance from CentigramPerGram.
func (uf MassFractionFactory) FromCentigramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionCentigramPerGram)
}

// FromDecigramPerGram creates a new MassFraction instance from DecigramPerGram.
func (uf MassFractionFactory) FromDecigramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecigramPerGram)
}

// FromDecagramPerGram creates a new MassFraction instance from DecagramPerGram.
func (uf MassFractionFactory) FromDecagramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecagramPerGram)
}

// FromHectogramPerGram creates a new MassFraction instance from HectogramPerGram.
func (uf MassFractionFactory) FromHectogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionHectogramPerGram)
}

// FromKilogramPerGram creates a new MassFraction instance from KilogramPerGram.
func (uf MassFractionFactory) FromKilogramsPerGram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionKilogramPerGram)
}

// FromNanogramPerKilogram creates a new MassFraction instance from NanogramPerKilogram.
func (uf MassFractionFactory) FromNanogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionNanogramPerKilogram)
}

// FromMicrogramPerKilogram creates a new MassFraction instance from MicrogramPerKilogram.
func (uf MassFractionFactory) FromMicrogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMicrogramPerKilogram)
}

// FromMilligramPerKilogram creates a new MassFraction instance from MilligramPerKilogram.
func (uf MassFractionFactory) FromMilligramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionMilligramPerKilogram)
}

// FromCentigramPerKilogram creates a new MassFraction instance from CentigramPerKilogram.
func (uf MassFractionFactory) FromCentigramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionCentigramPerKilogram)
}

// FromDecigramPerKilogram creates a new MassFraction instance from DecigramPerKilogram.
func (uf MassFractionFactory) FromDecigramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecigramPerKilogram)
}

// FromDecagramPerKilogram creates a new MassFraction instance from DecagramPerKilogram.
func (uf MassFractionFactory) FromDecagramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionDecagramPerKilogram)
}

// FromHectogramPerKilogram creates a new MassFraction instance from HectogramPerKilogram.
func (uf MassFractionFactory) FromHectogramsPerKilogram(value float64) (*MassFraction, error) {
	return newMassFraction(value, MassFractionHectogramPerKilogram)
}

// FromKilogramPerKilogram creates a new MassFraction instance from KilogramPerKilogram.
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

// BaseValue returns the base value of MassFraction in DecimalFraction.
func (a *MassFraction) BaseValue() float64 {
	return a.value
}


// DecimalFraction returns the value in DecimalFraction.
func (a *MassFraction) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(MassFractionDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// GramPerGram returns the value in GramPerGram.
func (a *MassFraction) GramsPerGram() float64 {
	if a.grams_per_gramLazy != nil {
		return *a.grams_per_gramLazy
	}
	grams_per_gram := a.convertFromBase(MassFractionGramPerGram)
	a.grams_per_gramLazy = &grams_per_gram
	return grams_per_gram
}

// GramPerKilogram returns the value in GramPerKilogram.
func (a *MassFraction) GramsPerKilogram() float64 {
	if a.grams_per_kilogramLazy != nil {
		return *a.grams_per_kilogramLazy
	}
	grams_per_kilogram := a.convertFromBase(MassFractionGramPerKilogram)
	a.grams_per_kilogramLazy = &grams_per_kilogram
	return grams_per_kilogram
}

// Percent returns the value in Percent.
func (a *MassFraction) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(MassFractionPercent)
	a.percentLazy = &percent
	return percent
}

// PartPerThousand returns the value in PartPerThousand.
func (a *MassFraction) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(MassFractionPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartPerMillion returns the value in PartPerMillion.
func (a *MassFraction) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(MassFractionPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartPerBillion returns the value in PartPerBillion.
func (a *MassFraction) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(MassFractionPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartPerTrillion returns the value in PartPerTrillion.
func (a *MassFraction) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(MassFractionPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}

// NanogramPerGram returns the value in NanogramPerGram.
func (a *MassFraction) NanogramsPerGram() float64 {
	if a.nanograms_per_gramLazy != nil {
		return *a.nanograms_per_gramLazy
	}
	nanograms_per_gram := a.convertFromBase(MassFractionNanogramPerGram)
	a.nanograms_per_gramLazy = &nanograms_per_gram
	return nanograms_per_gram
}

// MicrogramPerGram returns the value in MicrogramPerGram.
func (a *MassFraction) MicrogramsPerGram() float64 {
	if a.micrograms_per_gramLazy != nil {
		return *a.micrograms_per_gramLazy
	}
	micrograms_per_gram := a.convertFromBase(MassFractionMicrogramPerGram)
	a.micrograms_per_gramLazy = &micrograms_per_gram
	return micrograms_per_gram
}

// MilligramPerGram returns the value in MilligramPerGram.
func (a *MassFraction) MilligramsPerGram() float64 {
	if a.milligrams_per_gramLazy != nil {
		return *a.milligrams_per_gramLazy
	}
	milligrams_per_gram := a.convertFromBase(MassFractionMilligramPerGram)
	a.milligrams_per_gramLazy = &milligrams_per_gram
	return milligrams_per_gram
}

// CentigramPerGram returns the value in CentigramPerGram.
func (a *MassFraction) CentigramsPerGram() float64 {
	if a.centigrams_per_gramLazy != nil {
		return *a.centigrams_per_gramLazy
	}
	centigrams_per_gram := a.convertFromBase(MassFractionCentigramPerGram)
	a.centigrams_per_gramLazy = &centigrams_per_gram
	return centigrams_per_gram
}

// DecigramPerGram returns the value in DecigramPerGram.
func (a *MassFraction) DecigramsPerGram() float64 {
	if a.decigrams_per_gramLazy != nil {
		return *a.decigrams_per_gramLazy
	}
	decigrams_per_gram := a.convertFromBase(MassFractionDecigramPerGram)
	a.decigrams_per_gramLazy = &decigrams_per_gram
	return decigrams_per_gram
}

// DecagramPerGram returns the value in DecagramPerGram.
func (a *MassFraction) DecagramsPerGram() float64 {
	if a.decagrams_per_gramLazy != nil {
		return *a.decagrams_per_gramLazy
	}
	decagrams_per_gram := a.convertFromBase(MassFractionDecagramPerGram)
	a.decagrams_per_gramLazy = &decagrams_per_gram
	return decagrams_per_gram
}

// HectogramPerGram returns the value in HectogramPerGram.
func (a *MassFraction) HectogramsPerGram() float64 {
	if a.hectograms_per_gramLazy != nil {
		return *a.hectograms_per_gramLazy
	}
	hectograms_per_gram := a.convertFromBase(MassFractionHectogramPerGram)
	a.hectograms_per_gramLazy = &hectograms_per_gram
	return hectograms_per_gram
}

// KilogramPerGram returns the value in KilogramPerGram.
func (a *MassFraction) KilogramsPerGram() float64 {
	if a.kilograms_per_gramLazy != nil {
		return *a.kilograms_per_gramLazy
	}
	kilograms_per_gram := a.convertFromBase(MassFractionKilogramPerGram)
	a.kilograms_per_gramLazy = &kilograms_per_gram
	return kilograms_per_gram
}

// NanogramPerKilogram returns the value in NanogramPerKilogram.
func (a *MassFraction) NanogramsPerKilogram() float64 {
	if a.nanograms_per_kilogramLazy != nil {
		return *a.nanograms_per_kilogramLazy
	}
	nanograms_per_kilogram := a.convertFromBase(MassFractionNanogramPerKilogram)
	a.nanograms_per_kilogramLazy = &nanograms_per_kilogram
	return nanograms_per_kilogram
}

// MicrogramPerKilogram returns the value in MicrogramPerKilogram.
func (a *MassFraction) MicrogramsPerKilogram() float64 {
	if a.micrograms_per_kilogramLazy != nil {
		return *a.micrograms_per_kilogramLazy
	}
	micrograms_per_kilogram := a.convertFromBase(MassFractionMicrogramPerKilogram)
	a.micrograms_per_kilogramLazy = &micrograms_per_kilogram
	return micrograms_per_kilogram
}

// MilligramPerKilogram returns the value in MilligramPerKilogram.
func (a *MassFraction) MilligramsPerKilogram() float64 {
	if a.milligrams_per_kilogramLazy != nil {
		return *a.milligrams_per_kilogramLazy
	}
	milligrams_per_kilogram := a.convertFromBase(MassFractionMilligramPerKilogram)
	a.milligrams_per_kilogramLazy = &milligrams_per_kilogram
	return milligrams_per_kilogram
}

// CentigramPerKilogram returns the value in CentigramPerKilogram.
func (a *MassFraction) CentigramsPerKilogram() float64 {
	if a.centigrams_per_kilogramLazy != nil {
		return *a.centigrams_per_kilogramLazy
	}
	centigrams_per_kilogram := a.convertFromBase(MassFractionCentigramPerKilogram)
	a.centigrams_per_kilogramLazy = &centigrams_per_kilogram
	return centigrams_per_kilogram
}

// DecigramPerKilogram returns the value in DecigramPerKilogram.
func (a *MassFraction) DecigramsPerKilogram() float64 {
	if a.decigrams_per_kilogramLazy != nil {
		return *a.decigrams_per_kilogramLazy
	}
	decigrams_per_kilogram := a.convertFromBase(MassFractionDecigramPerKilogram)
	a.decigrams_per_kilogramLazy = &decigrams_per_kilogram
	return decigrams_per_kilogram
}

// DecagramPerKilogram returns the value in DecagramPerKilogram.
func (a *MassFraction) DecagramsPerKilogram() float64 {
	if a.decagrams_per_kilogramLazy != nil {
		return *a.decagrams_per_kilogramLazy
	}
	decagrams_per_kilogram := a.convertFromBase(MassFractionDecagramPerKilogram)
	a.decagrams_per_kilogramLazy = &decagrams_per_kilogram
	return decagrams_per_kilogram
}

// HectogramPerKilogram returns the value in HectogramPerKilogram.
func (a *MassFraction) HectogramsPerKilogram() float64 {
	if a.hectograms_per_kilogramLazy != nil {
		return *a.hectograms_per_kilogramLazy
	}
	hectograms_per_kilogram := a.convertFromBase(MassFractionHectogramPerKilogram)
	a.hectograms_per_kilogramLazy = &hectograms_per_kilogram
	return hectograms_per_kilogram
}

// KilogramPerKilogram returns the value in KilogramPerKilogram.
func (a *MassFraction) KilogramsPerKilogram() float64 {
	if a.kilograms_per_kilogramLazy != nil {
		return *a.kilograms_per_kilogramLazy
	}
	kilograms_per_kilogram := a.convertFromBase(MassFractionKilogramPerKilogram)
	a.kilograms_per_kilogramLazy = &kilograms_per_kilogram
	return kilograms_per_kilogram
}


// ToDto creates an MassFractionDto representation.
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

// ToDtoJSON creates an MassFractionDto representation.
func (a *MassFraction) ToDtoJSON(holdInUnit *MassFractionUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MassFraction to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a MassFraction) String() string {
	return a.ToString(MassFractionDecimalFraction, 2)
}

// ToString formats the MassFraction to string.
// fractionalDigits -1 for not mention
func (a *MassFraction) ToString(unit MassFractionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MassFraction) getUnitAbbreviation(unit MassFractionUnits) string {
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

// Check if the given MassFraction are equals to the current MassFraction
func (a *MassFraction) Equals(other *MassFraction) bool {
	return a.value == other.BaseValue()
}

// Check if the given MassFraction are equals to the current MassFraction
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

// Add the given MassFraction to the current MassFraction.
func (a *MassFraction) Add(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value + other.BaseValue()}
}

// Subtract the given MassFraction to the current MassFraction.
func (a *MassFraction) Subtract(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value - other.BaseValue()}
}

// Multiply the given MassFraction to the current MassFraction.
func (a *MassFraction) Multiply(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value * other.BaseValue()}
}

// Divide the given MassFraction to the current MassFraction.
func (a *MassFraction) Divide(other *MassFraction) *MassFraction {
	return &MassFraction{value: a.value / other.BaseValue()}
}