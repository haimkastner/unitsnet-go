// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassConcentrationUnits defines various units of MassConcentration.
type MassConcentrationUnits string

const (
    
        // 
        MassConcentrationGramPerCubicMillimeter MassConcentrationUnits = "GramPerCubicMillimeter"
        // 
        MassConcentrationGramPerCubicCentimeter MassConcentrationUnits = "GramPerCubicCentimeter"
        // 
        MassConcentrationGramPerCubicMeter MassConcentrationUnits = "GramPerCubicMeter"
        // 
        MassConcentrationGramPerMicroliter MassConcentrationUnits = "GramPerMicroliter"
        // 
        MassConcentrationGramPerMilliliter MassConcentrationUnits = "GramPerMilliliter"
        // 
        MassConcentrationGramPerDeciliter MassConcentrationUnits = "GramPerDeciliter"
        // 
        MassConcentrationGramPerLiter MassConcentrationUnits = "GramPerLiter"
        // 
        MassConcentrationTonnePerCubicMillimeter MassConcentrationUnits = "TonnePerCubicMillimeter"
        // 
        MassConcentrationTonnePerCubicCentimeter MassConcentrationUnits = "TonnePerCubicCentimeter"
        // 
        MassConcentrationTonnePerCubicMeter MassConcentrationUnits = "TonnePerCubicMeter"
        // 
        MassConcentrationPoundPerCubicInch MassConcentrationUnits = "PoundPerCubicInch"
        // 
        MassConcentrationPoundPerCubicFoot MassConcentrationUnits = "PoundPerCubicFoot"
        // 
        MassConcentrationSlugPerCubicFoot MassConcentrationUnits = "SlugPerCubicFoot"
        // 
        MassConcentrationPoundPerUSGallon MassConcentrationUnits = "PoundPerUSGallon"
        // 
        MassConcentrationOuncePerUSGallon MassConcentrationUnits = "OuncePerUSGallon"
        // 
        MassConcentrationOuncePerImperialGallon MassConcentrationUnits = "OuncePerImperialGallon"
        // 
        MassConcentrationPoundPerImperialGallon MassConcentrationUnits = "PoundPerImperialGallon"
        // 
        MassConcentrationKilogramPerCubicMillimeter MassConcentrationUnits = "KilogramPerCubicMillimeter"
        // 
        MassConcentrationKilogramPerCubicCentimeter MassConcentrationUnits = "KilogramPerCubicCentimeter"
        // 
        MassConcentrationKilogramPerCubicMeter MassConcentrationUnits = "KilogramPerCubicMeter"
        // 
        MassConcentrationMilligramPerCubicMeter MassConcentrationUnits = "MilligramPerCubicMeter"
        // 
        MassConcentrationMicrogramPerCubicMeter MassConcentrationUnits = "MicrogramPerCubicMeter"
        // 
        MassConcentrationPicogramPerMicroliter MassConcentrationUnits = "PicogramPerMicroliter"
        // 
        MassConcentrationNanogramPerMicroliter MassConcentrationUnits = "NanogramPerMicroliter"
        // 
        MassConcentrationMicrogramPerMicroliter MassConcentrationUnits = "MicrogramPerMicroliter"
        // 
        MassConcentrationMilligramPerMicroliter MassConcentrationUnits = "MilligramPerMicroliter"
        // 
        MassConcentrationCentigramPerMicroliter MassConcentrationUnits = "CentigramPerMicroliter"
        // 
        MassConcentrationDecigramPerMicroliter MassConcentrationUnits = "DecigramPerMicroliter"
        // 
        MassConcentrationPicogramPerMilliliter MassConcentrationUnits = "PicogramPerMilliliter"
        // 
        MassConcentrationNanogramPerMilliliter MassConcentrationUnits = "NanogramPerMilliliter"
        // 
        MassConcentrationMicrogramPerMilliliter MassConcentrationUnits = "MicrogramPerMilliliter"
        // 
        MassConcentrationMilligramPerMilliliter MassConcentrationUnits = "MilligramPerMilliliter"
        // 
        MassConcentrationCentigramPerMilliliter MassConcentrationUnits = "CentigramPerMilliliter"
        // 
        MassConcentrationDecigramPerMilliliter MassConcentrationUnits = "DecigramPerMilliliter"
        // 
        MassConcentrationPicogramPerDeciliter MassConcentrationUnits = "PicogramPerDeciliter"
        // 
        MassConcentrationNanogramPerDeciliter MassConcentrationUnits = "NanogramPerDeciliter"
        // 
        MassConcentrationMicrogramPerDeciliter MassConcentrationUnits = "MicrogramPerDeciliter"
        // 
        MassConcentrationMilligramPerDeciliter MassConcentrationUnits = "MilligramPerDeciliter"
        // 
        MassConcentrationCentigramPerDeciliter MassConcentrationUnits = "CentigramPerDeciliter"
        // 
        MassConcentrationDecigramPerDeciliter MassConcentrationUnits = "DecigramPerDeciliter"
        // 
        MassConcentrationPicogramPerLiter MassConcentrationUnits = "PicogramPerLiter"
        // 
        MassConcentrationNanogramPerLiter MassConcentrationUnits = "NanogramPerLiter"
        // 
        MassConcentrationMicrogramPerLiter MassConcentrationUnits = "MicrogramPerLiter"
        // 
        MassConcentrationMilligramPerLiter MassConcentrationUnits = "MilligramPerLiter"
        // 
        MassConcentrationCentigramPerLiter MassConcentrationUnits = "CentigramPerLiter"
        // 
        MassConcentrationDecigramPerLiter MassConcentrationUnits = "DecigramPerLiter"
        // 
        MassConcentrationKilogramPerLiter MassConcentrationUnits = "KilogramPerLiter"
        // 
        MassConcentrationKilopoundPerCubicInch MassConcentrationUnits = "KilopoundPerCubicInch"
        // 
        MassConcentrationKilopoundPerCubicFoot MassConcentrationUnits = "KilopoundPerCubicFoot"
)

// MassConcentrationDto represents a MassConcentration measurement with a numerical value and its corresponding unit.
type MassConcentrationDto struct {
    // Value is the numerical representation of the MassConcentration.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MassConcentration, as defined in the MassConcentrationUnits enumeration.
	Unit  MassConcentrationUnits `json:"unit"`
}

// MassConcentrationDtoFactory groups methods for creating and serializing MassConcentrationDto objects.
type MassConcentrationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassConcentrationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassConcentrationDtoFactory) FromJSON(data []byte) (*MassConcentrationDto, error) {
	a := MassConcentrationDto{}

    // Parse JSON into MassConcentrationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MassConcentrationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassConcentrationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MassConcentration represents a measurement in a of MassConcentration.
//
// In chemistry, the mass concentration ρi (or γi) is defined as the mass of a constituent mi divided by the volume of the mixture V
type MassConcentration struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_cubic_millimeterLazy *float64 
    grams_per_cubic_centimeterLazy *float64 
    grams_per_cubic_meterLazy *float64 
    grams_per_microliterLazy *float64 
    grams_per_milliliterLazy *float64 
    grams_per_deciliterLazy *float64 
    grams_per_literLazy *float64 
    tonnes_per_cubic_millimeterLazy *float64 
    tonnes_per_cubic_centimeterLazy *float64 
    tonnes_per_cubic_meterLazy *float64 
    pounds_per_cubic_inchLazy *float64 
    pounds_per_cubic_footLazy *float64 
    slugs_per_cubic_footLazy *float64 
    pounds_per_us_gallonLazy *float64 
    ounces_per_us_gallonLazy *float64 
    ounces_per_imperial_gallonLazy *float64 
    pounds_per_imperial_gallonLazy *float64 
    kilograms_per_cubic_millimeterLazy *float64 
    kilograms_per_cubic_centimeterLazy *float64 
    kilograms_per_cubic_meterLazy *float64 
    milligrams_per_cubic_meterLazy *float64 
    micrograms_per_cubic_meterLazy *float64 
    picograms_per_microliterLazy *float64 
    nanograms_per_microliterLazy *float64 
    micrograms_per_microliterLazy *float64 
    milligrams_per_microliterLazy *float64 
    centigrams_per_microliterLazy *float64 
    decigrams_per_microliterLazy *float64 
    picograms_per_milliliterLazy *float64 
    nanograms_per_milliliterLazy *float64 
    micrograms_per_milliliterLazy *float64 
    milligrams_per_milliliterLazy *float64 
    centigrams_per_milliliterLazy *float64 
    decigrams_per_milliliterLazy *float64 
    picograms_per_deciliterLazy *float64 
    nanograms_per_deciliterLazy *float64 
    micrograms_per_deciliterLazy *float64 
    milligrams_per_deciliterLazy *float64 
    centigrams_per_deciliterLazy *float64 
    decigrams_per_deciliterLazy *float64 
    picograms_per_literLazy *float64 
    nanograms_per_literLazy *float64 
    micrograms_per_literLazy *float64 
    milligrams_per_literLazy *float64 
    centigrams_per_literLazy *float64 
    decigrams_per_literLazy *float64 
    kilograms_per_literLazy *float64 
    kilopounds_per_cubic_inchLazy *float64 
    kilopounds_per_cubic_footLazy *float64 
}

// MassConcentrationFactory groups methods for creating MassConcentration instances.
type MassConcentrationFactory struct{}

// CreateMassConcentration creates a new MassConcentration instance from the given value and unit.
func (uf MassConcentrationFactory) CreateMassConcentration(value float64, unit MassConcentrationUnits) (*MassConcentration, error) {
	return newMassConcentration(value, unit)
}

// FromDto converts a MassConcentrationDto to a MassConcentration instance.
func (uf MassConcentrationFactory) FromDto(dto MassConcentrationDto) (*MassConcentration, error) {
	return newMassConcentration(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MassConcentration instance.
func (uf MassConcentrationFactory) FromDtoJSON(data []byte) (*MassConcentration, error) {
	unitDto, err := MassConcentrationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassConcentrationDto from JSON: %w", err)
	}
	return MassConcentrationFactory{}.FromDto(*unitDto)
}


// FromGramsPerCubicMillimeter creates a new MassConcentration instance from a value in GramsPerCubicMillimeter.
func (uf MassConcentrationFactory) FromGramsPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicMillimeter)
}

// FromGramsPerCubicCentimeter creates a new MassConcentration instance from a value in GramsPerCubicCentimeter.
func (uf MassConcentrationFactory) FromGramsPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicCentimeter)
}

// FromGramsPerCubicMeter creates a new MassConcentration instance from a value in GramsPerCubicMeter.
func (uf MassConcentrationFactory) FromGramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicMeter)
}

// FromGramsPerMicroliter creates a new MassConcentration instance from a value in GramsPerMicroliter.
func (uf MassConcentrationFactory) FromGramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerMicroliter)
}

// FromGramsPerMilliliter creates a new MassConcentration instance from a value in GramsPerMilliliter.
func (uf MassConcentrationFactory) FromGramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerMilliliter)
}

// FromGramsPerDeciliter creates a new MassConcentration instance from a value in GramsPerDeciliter.
func (uf MassConcentrationFactory) FromGramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerDeciliter)
}

// FromGramsPerLiter creates a new MassConcentration instance from a value in GramsPerLiter.
func (uf MassConcentrationFactory) FromGramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerLiter)
}

// FromTonnesPerCubicMillimeter creates a new MassConcentration instance from a value in TonnesPerCubicMillimeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicMillimeter)
}

// FromTonnesPerCubicCentimeter creates a new MassConcentration instance from a value in TonnesPerCubicCentimeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicCentimeter)
}

// FromTonnesPerCubicMeter creates a new MassConcentration instance from a value in TonnesPerCubicMeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicMeter)
}

// FromPoundsPerCubicInch creates a new MassConcentration instance from a value in PoundsPerCubicInch.
func (uf MassConcentrationFactory) FromPoundsPerCubicInch(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerCubicInch)
}

// FromPoundsPerCubicFoot creates a new MassConcentration instance from a value in PoundsPerCubicFoot.
func (uf MassConcentrationFactory) FromPoundsPerCubicFoot(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerCubicFoot)
}

// FromSlugsPerCubicFoot creates a new MassConcentration instance from a value in SlugsPerCubicFoot.
func (uf MassConcentrationFactory) FromSlugsPerCubicFoot(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationSlugPerCubicFoot)
}

// FromPoundsPerUSGallon creates a new MassConcentration instance from a value in PoundsPerUSGallon.
func (uf MassConcentrationFactory) FromPoundsPerUSGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerUSGallon)
}

// FromOuncesPerUSGallon creates a new MassConcentration instance from a value in OuncesPerUSGallon.
func (uf MassConcentrationFactory) FromOuncesPerUSGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationOuncePerUSGallon)
}

// FromOuncesPerImperialGallon creates a new MassConcentration instance from a value in OuncesPerImperialGallon.
func (uf MassConcentrationFactory) FromOuncesPerImperialGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationOuncePerImperialGallon)
}

// FromPoundsPerImperialGallon creates a new MassConcentration instance from a value in PoundsPerImperialGallon.
func (uf MassConcentrationFactory) FromPoundsPerImperialGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerImperialGallon)
}

// FromKilogramsPerCubicMillimeter creates a new MassConcentration instance from a value in KilogramsPerCubicMillimeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicMillimeter)
}

// FromKilogramsPerCubicCentimeter creates a new MassConcentration instance from a value in KilogramsPerCubicCentimeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicCentimeter)
}

// FromKilogramsPerCubicMeter creates a new MassConcentration instance from a value in KilogramsPerCubicMeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicMeter)
}

// FromMilligramsPerCubicMeter creates a new MassConcentration instance from a value in MilligramsPerCubicMeter.
func (uf MassConcentrationFactory) FromMilligramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerCubicMeter)
}

// FromMicrogramsPerCubicMeter creates a new MassConcentration instance from a value in MicrogramsPerCubicMeter.
func (uf MassConcentrationFactory) FromMicrogramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerCubicMeter)
}

// FromPicogramsPerMicroliter creates a new MassConcentration instance from a value in PicogramsPerMicroliter.
func (uf MassConcentrationFactory) FromPicogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerMicroliter)
}

// FromNanogramsPerMicroliter creates a new MassConcentration instance from a value in NanogramsPerMicroliter.
func (uf MassConcentrationFactory) FromNanogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerMicroliter)
}

// FromMicrogramsPerMicroliter creates a new MassConcentration instance from a value in MicrogramsPerMicroliter.
func (uf MassConcentrationFactory) FromMicrogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerMicroliter)
}

// FromMilligramsPerMicroliter creates a new MassConcentration instance from a value in MilligramsPerMicroliter.
func (uf MassConcentrationFactory) FromMilligramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerMicroliter)
}

// FromCentigramsPerMicroliter creates a new MassConcentration instance from a value in CentigramsPerMicroliter.
func (uf MassConcentrationFactory) FromCentigramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerMicroliter)
}

// FromDecigramsPerMicroliter creates a new MassConcentration instance from a value in DecigramsPerMicroliter.
func (uf MassConcentrationFactory) FromDecigramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerMicroliter)
}

// FromPicogramsPerMilliliter creates a new MassConcentration instance from a value in PicogramsPerMilliliter.
func (uf MassConcentrationFactory) FromPicogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerMilliliter)
}

// FromNanogramsPerMilliliter creates a new MassConcentration instance from a value in NanogramsPerMilliliter.
func (uf MassConcentrationFactory) FromNanogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerMilliliter)
}

// FromMicrogramsPerMilliliter creates a new MassConcentration instance from a value in MicrogramsPerMilliliter.
func (uf MassConcentrationFactory) FromMicrogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerMilliliter)
}

// FromMilligramsPerMilliliter creates a new MassConcentration instance from a value in MilligramsPerMilliliter.
func (uf MassConcentrationFactory) FromMilligramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerMilliliter)
}

// FromCentigramsPerMilliliter creates a new MassConcentration instance from a value in CentigramsPerMilliliter.
func (uf MassConcentrationFactory) FromCentigramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerMilliliter)
}

// FromDecigramsPerMilliliter creates a new MassConcentration instance from a value in DecigramsPerMilliliter.
func (uf MassConcentrationFactory) FromDecigramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerMilliliter)
}

// FromPicogramsPerDeciliter creates a new MassConcentration instance from a value in PicogramsPerDeciliter.
func (uf MassConcentrationFactory) FromPicogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerDeciliter)
}

// FromNanogramsPerDeciliter creates a new MassConcentration instance from a value in NanogramsPerDeciliter.
func (uf MassConcentrationFactory) FromNanogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerDeciliter)
}

// FromMicrogramsPerDeciliter creates a new MassConcentration instance from a value in MicrogramsPerDeciliter.
func (uf MassConcentrationFactory) FromMicrogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerDeciliter)
}

// FromMilligramsPerDeciliter creates a new MassConcentration instance from a value in MilligramsPerDeciliter.
func (uf MassConcentrationFactory) FromMilligramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerDeciliter)
}

// FromCentigramsPerDeciliter creates a new MassConcentration instance from a value in CentigramsPerDeciliter.
func (uf MassConcentrationFactory) FromCentigramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerDeciliter)
}

// FromDecigramsPerDeciliter creates a new MassConcentration instance from a value in DecigramsPerDeciliter.
func (uf MassConcentrationFactory) FromDecigramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerDeciliter)
}

// FromPicogramsPerLiter creates a new MassConcentration instance from a value in PicogramsPerLiter.
func (uf MassConcentrationFactory) FromPicogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerLiter)
}

// FromNanogramsPerLiter creates a new MassConcentration instance from a value in NanogramsPerLiter.
func (uf MassConcentrationFactory) FromNanogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerLiter)
}

// FromMicrogramsPerLiter creates a new MassConcentration instance from a value in MicrogramsPerLiter.
func (uf MassConcentrationFactory) FromMicrogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerLiter)
}

// FromMilligramsPerLiter creates a new MassConcentration instance from a value in MilligramsPerLiter.
func (uf MassConcentrationFactory) FromMilligramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerLiter)
}

// FromCentigramsPerLiter creates a new MassConcentration instance from a value in CentigramsPerLiter.
func (uf MassConcentrationFactory) FromCentigramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerLiter)
}

// FromDecigramsPerLiter creates a new MassConcentration instance from a value in DecigramsPerLiter.
func (uf MassConcentrationFactory) FromDecigramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerLiter)
}

// FromKilogramsPerLiter creates a new MassConcentration instance from a value in KilogramsPerLiter.
func (uf MassConcentrationFactory) FromKilogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerLiter)
}

// FromKilopoundsPerCubicInch creates a new MassConcentration instance from a value in KilopoundsPerCubicInch.
func (uf MassConcentrationFactory) FromKilopoundsPerCubicInch(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilopoundPerCubicInch)
}

// FromKilopoundsPerCubicFoot creates a new MassConcentration instance from a value in KilopoundsPerCubicFoot.
func (uf MassConcentrationFactory) FromKilopoundsPerCubicFoot(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilopoundPerCubicFoot)
}


// newMassConcentration creates a new MassConcentration.
func newMassConcentration(value float64, fromUnit MassConcentrationUnits) (*MassConcentration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MassConcentration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassConcentration in KilogramPerCubicMeter unit (the base/default quantity).
func (a *MassConcentration) BaseValue() float64 {
	return a.value
}


// GramsPerCubicMillimeter returns the MassConcentration value in GramsPerCubicMillimeter.
//
// 
func (a *MassConcentration) GramsPerCubicMillimeter() float64 {
	if a.grams_per_cubic_millimeterLazy != nil {
		return *a.grams_per_cubic_millimeterLazy
	}
	grams_per_cubic_millimeter := a.convertFromBase(MassConcentrationGramPerCubicMillimeter)
	a.grams_per_cubic_millimeterLazy = &grams_per_cubic_millimeter
	return grams_per_cubic_millimeter
}

// GramsPerCubicCentimeter returns the MassConcentration value in GramsPerCubicCentimeter.
//
// 
func (a *MassConcentration) GramsPerCubicCentimeter() float64 {
	if a.grams_per_cubic_centimeterLazy != nil {
		return *a.grams_per_cubic_centimeterLazy
	}
	grams_per_cubic_centimeter := a.convertFromBase(MassConcentrationGramPerCubicCentimeter)
	a.grams_per_cubic_centimeterLazy = &grams_per_cubic_centimeter
	return grams_per_cubic_centimeter
}

// GramsPerCubicMeter returns the MassConcentration value in GramsPerCubicMeter.
//
// 
func (a *MassConcentration) GramsPerCubicMeter() float64 {
	if a.grams_per_cubic_meterLazy != nil {
		return *a.grams_per_cubic_meterLazy
	}
	grams_per_cubic_meter := a.convertFromBase(MassConcentrationGramPerCubicMeter)
	a.grams_per_cubic_meterLazy = &grams_per_cubic_meter
	return grams_per_cubic_meter
}

// GramsPerMicroliter returns the MassConcentration value in GramsPerMicroliter.
//
// 
func (a *MassConcentration) GramsPerMicroliter() float64 {
	if a.grams_per_microliterLazy != nil {
		return *a.grams_per_microliterLazy
	}
	grams_per_microliter := a.convertFromBase(MassConcentrationGramPerMicroliter)
	a.grams_per_microliterLazy = &grams_per_microliter
	return grams_per_microliter
}

// GramsPerMilliliter returns the MassConcentration value in GramsPerMilliliter.
//
// 
func (a *MassConcentration) GramsPerMilliliter() float64 {
	if a.grams_per_milliliterLazy != nil {
		return *a.grams_per_milliliterLazy
	}
	grams_per_milliliter := a.convertFromBase(MassConcentrationGramPerMilliliter)
	a.grams_per_milliliterLazy = &grams_per_milliliter
	return grams_per_milliliter
}

// GramsPerDeciliter returns the MassConcentration value in GramsPerDeciliter.
//
// 
func (a *MassConcentration) GramsPerDeciliter() float64 {
	if a.grams_per_deciliterLazy != nil {
		return *a.grams_per_deciliterLazy
	}
	grams_per_deciliter := a.convertFromBase(MassConcentrationGramPerDeciliter)
	a.grams_per_deciliterLazy = &grams_per_deciliter
	return grams_per_deciliter
}

// GramsPerLiter returns the MassConcentration value in GramsPerLiter.
//
// 
func (a *MassConcentration) GramsPerLiter() float64 {
	if a.grams_per_literLazy != nil {
		return *a.grams_per_literLazy
	}
	grams_per_liter := a.convertFromBase(MassConcentrationGramPerLiter)
	a.grams_per_literLazy = &grams_per_liter
	return grams_per_liter
}

// TonnesPerCubicMillimeter returns the MassConcentration value in TonnesPerCubicMillimeter.
//
// 
func (a *MassConcentration) TonnesPerCubicMillimeter() float64 {
	if a.tonnes_per_cubic_millimeterLazy != nil {
		return *a.tonnes_per_cubic_millimeterLazy
	}
	tonnes_per_cubic_millimeter := a.convertFromBase(MassConcentrationTonnePerCubicMillimeter)
	a.tonnes_per_cubic_millimeterLazy = &tonnes_per_cubic_millimeter
	return tonnes_per_cubic_millimeter
}

// TonnesPerCubicCentimeter returns the MassConcentration value in TonnesPerCubicCentimeter.
//
// 
func (a *MassConcentration) TonnesPerCubicCentimeter() float64 {
	if a.tonnes_per_cubic_centimeterLazy != nil {
		return *a.tonnes_per_cubic_centimeterLazy
	}
	tonnes_per_cubic_centimeter := a.convertFromBase(MassConcentrationTonnePerCubicCentimeter)
	a.tonnes_per_cubic_centimeterLazy = &tonnes_per_cubic_centimeter
	return tonnes_per_cubic_centimeter
}

// TonnesPerCubicMeter returns the MassConcentration value in TonnesPerCubicMeter.
//
// 
func (a *MassConcentration) TonnesPerCubicMeter() float64 {
	if a.tonnes_per_cubic_meterLazy != nil {
		return *a.tonnes_per_cubic_meterLazy
	}
	tonnes_per_cubic_meter := a.convertFromBase(MassConcentrationTonnePerCubicMeter)
	a.tonnes_per_cubic_meterLazy = &tonnes_per_cubic_meter
	return tonnes_per_cubic_meter
}

// PoundsPerCubicInch returns the MassConcentration value in PoundsPerCubicInch.
//
// 
func (a *MassConcentration) PoundsPerCubicInch() float64 {
	if a.pounds_per_cubic_inchLazy != nil {
		return *a.pounds_per_cubic_inchLazy
	}
	pounds_per_cubic_inch := a.convertFromBase(MassConcentrationPoundPerCubicInch)
	a.pounds_per_cubic_inchLazy = &pounds_per_cubic_inch
	return pounds_per_cubic_inch
}

// PoundsPerCubicFoot returns the MassConcentration value in PoundsPerCubicFoot.
//
// 
func (a *MassConcentration) PoundsPerCubicFoot() float64 {
	if a.pounds_per_cubic_footLazy != nil {
		return *a.pounds_per_cubic_footLazy
	}
	pounds_per_cubic_foot := a.convertFromBase(MassConcentrationPoundPerCubicFoot)
	a.pounds_per_cubic_footLazy = &pounds_per_cubic_foot
	return pounds_per_cubic_foot
}

// SlugsPerCubicFoot returns the MassConcentration value in SlugsPerCubicFoot.
//
// 
func (a *MassConcentration) SlugsPerCubicFoot() float64 {
	if a.slugs_per_cubic_footLazy != nil {
		return *a.slugs_per_cubic_footLazy
	}
	slugs_per_cubic_foot := a.convertFromBase(MassConcentrationSlugPerCubicFoot)
	a.slugs_per_cubic_footLazy = &slugs_per_cubic_foot
	return slugs_per_cubic_foot
}

// PoundsPerUSGallon returns the MassConcentration value in PoundsPerUSGallon.
//
// 
func (a *MassConcentration) PoundsPerUSGallon() float64 {
	if a.pounds_per_us_gallonLazy != nil {
		return *a.pounds_per_us_gallonLazy
	}
	pounds_per_us_gallon := a.convertFromBase(MassConcentrationPoundPerUSGallon)
	a.pounds_per_us_gallonLazy = &pounds_per_us_gallon
	return pounds_per_us_gallon
}

// OuncesPerUSGallon returns the MassConcentration value in OuncesPerUSGallon.
//
// 
func (a *MassConcentration) OuncesPerUSGallon() float64 {
	if a.ounces_per_us_gallonLazy != nil {
		return *a.ounces_per_us_gallonLazy
	}
	ounces_per_us_gallon := a.convertFromBase(MassConcentrationOuncePerUSGallon)
	a.ounces_per_us_gallonLazy = &ounces_per_us_gallon
	return ounces_per_us_gallon
}

// OuncesPerImperialGallon returns the MassConcentration value in OuncesPerImperialGallon.
//
// 
func (a *MassConcentration) OuncesPerImperialGallon() float64 {
	if a.ounces_per_imperial_gallonLazy != nil {
		return *a.ounces_per_imperial_gallonLazy
	}
	ounces_per_imperial_gallon := a.convertFromBase(MassConcentrationOuncePerImperialGallon)
	a.ounces_per_imperial_gallonLazy = &ounces_per_imperial_gallon
	return ounces_per_imperial_gallon
}

// PoundsPerImperialGallon returns the MassConcentration value in PoundsPerImperialGallon.
//
// 
func (a *MassConcentration) PoundsPerImperialGallon() float64 {
	if a.pounds_per_imperial_gallonLazy != nil {
		return *a.pounds_per_imperial_gallonLazy
	}
	pounds_per_imperial_gallon := a.convertFromBase(MassConcentrationPoundPerImperialGallon)
	a.pounds_per_imperial_gallonLazy = &pounds_per_imperial_gallon
	return pounds_per_imperial_gallon
}

// KilogramsPerCubicMillimeter returns the MassConcentration value in KilogramsPerCubicMillimeter.
//
// 
func (a *MassConcentration) KilogramsPerCubicMillimeter() float64 {
	if a.kilograms_per_cubic_millimeterLazy != nil {
		return *a.kilograms_per_cubic_millimeterLazy
	}
	kilograms_per_cubic_millimeter := a.convertFromBase(MassConcentrationKilogramPerCubicMillimeter)
	a.kilograms_per_cubic_millimeterLazy = &kilograms_per_cubic_millimeter
	return kilograms_per_cubic_millimeter
}

// KilogramsPerCubicCentimeter returns the MassConcentration value in KilogramsPerCubicCentimeter.
//
// 
func (a *MassConcentration) KilogramsPerCubicCentimeter() float64 {
	if a.kilograms_per_cubic_centimeterLazy != nil {
		return *a.kilograms_per_cubic_centimeterLazy
	}
	kilograms_per_cubic_centimeter := a.convertFromBase(MassConcentrationKilogramPerCubicCentimeter)
	a.kilograms_per_cubic_centimeterLazy = &kilograms_per_cubic_centimeter
	return kilograms_per_cubic_centimeter
}

// KilogramsPerCubicMeter returns the MassConcentration value in KilogramsPerCubicMeter.
//
// 
func (a *MassConcentration) KilogramsPerCubicMeter() float64 {
	if a.kilograms_per_cubic_meterLazy != nil {
		return *a.kilograms_per_cubic_meterLazy
	}
	kilograms_per_cubic_meter := a.convertFromBase(MassConcentrationKilogramPerCubicMeter)
	a.kilograms_per_cubic_meterLazy = &kilograms_per_cubic_meter
	return kilograms_per_cubic_meter
}

// MilligramsPerCubicMeter returns the MassConcentration value in MilligramsPerCubicMeter.
//
// 
func (a *MassConcentration) MilligramsPerCubicMeter() float64 {
	if a.milligrams_per_cubic_meterLazy != nil {
		return *a.milligrams_per_cubic_meterLazy
	}
	milligrams_per_cubic_meter := a.convertFromBase(MassConcentrationMilligramPerCubicMeter)
	a.milligrams_per_cubic_meterLazy = &milligrams_per_cubic_meter
	return milligrams_per_cubic_meter
}

// MicrogramsPerCubicMeter returns the MassConcentration value in MicrogramsPerCubicMeter.
//
// 
func (a *MassConcentration) MicrogramsPerCubicMeter() float64 {
	if a.micrograms_per_cubic_meterLazy != nil {
		return *a.micrograms_per_cubic_meterLazy
	}
	micrograms_per_cubic_meter := a.convertFromBase(MassConcentrationMicrogramPerCubicMeter)
	a.micrograms_per_cubic_meterLazy = &micrograms_per_cubic_meter
	return micrograms_per_cubic_meter
}

// PicogramsPerMicroliter returns the MassConcentration value in PicogramsPerMicroliter.
//
// 
func (a *MassConcentration) PicogramsPerMicroliter() float64 {
	if a.picograms_per_microliterLazy != nil {
		return *a.picograms_per_microliterLazy
	}
	picograms_per_microliter := a.convertFromBase(MassConcentrationPicogramPerMicroliter)
	a.picograms_per_microliterLazy = &picograms_per_microliter
	return picograms_per_microliter
}

// NanogramsPerMicroliter returns the MassConcentration value in NanogramsPerMicroliter.
//
// 
func (a *MassConcentration) NanogramsPerMicroliter() float64 {
	if a.nanograms_per_microliterLazy != nil {
		return *a.nanograms_per_microliterLazy
	}
	nanograms_per_microliter := a.convertFromBase(MassConcentrationNanogramPerMicroliter)
	a.nanograms_per_microliterLazy = &nanograms_per_microliter
	return nanograms_per_microliter
}

// MicrogramsPerMicroliter returns the MassConcentration value in MicrogramsPerMicroliter.
//
// 
func (a *MassConcentration) MicrogramsPerMicroliter() float64 {
	if a.micrograms_per_microliterLazy != nil {
		return *a.micrograms_per_microliterLazy
	}
	micrograms_per_microliter := a.convertFromBase(MassConcentrationMicrogramPerMicroliter)
	a.micrograms_per_microliterLazy = &micrograms_per_microliter
	return micrograms_per_microliter
}

// MilligramsPerMicroliter returns the MassConcentration value in MilligramsPerMicroliter.
//
// 
func (a *MassConcentration) MilligramsPerMicroliter() float64 {
	if a.milligrams_per_microliterLazy != nil {
		return *a.milligrams_per_microliterLazy
	}
	milligrams_per_microliter := a.convertFromBase(MassConcentrationMilligramPerMicroliter)
	a.milligrams_per_microliterLazy = &milligrams_per_microliter
	return milligrams_per_microliter
}

// CentigramsPerMicroliter returns the MassConcentration value in CentigramsPerMicroliter.
//
// 
func (a *MassConcentration) CentigramsPerMicroliter() float64 {
	if a.centigrams_per_microliterLazy != nil {
		return *a.centigrams_per_microliterLazy
	}
	centigrams_per_microliter := a.convertFromBase(MassConcentrationCentigramPerMicroliter)
	a.centigrams_per_microliterLazy = &centigrams_per_microliter
	return centigrams_per_microliter
}

// DecigramsPerMicroliter returns the MassConcentration value in DecigramsPerMicroliter.
//
// 
func (a *MassConcentration) DecigramsPerMicroliter() float64 {
	if a.decigrams_per_microliterLazy != nil {
		return *a.decigrams_per_microliterLazy
	}
	decigrams_per_microliter := a.convertFromBase(MassConcentrationDecigramPerMicroliter)
	a.decigrams_per_microliterLazy = &decigrams_per_microliter
	return decigrams_per_microliter
}

// PicogramsPerMilliliter returns the MassConcentration value in PicogramsPerMilliliter.
//
// 
func (a *MassConcentration) PicogramsPerMilliliter() float64 {
	if a.picograms_per_milliliterLazy != nil {
		return *a.picograms_per_milliliterLazy
	}
	picograms_per_milliliter := a.convertFromBase(MassConcentrationPicogramPerMilliliter)
	a.picograms_per_milliliterLazy = &picograms_per_milliliter
	return picograms_per_milliliter
}

// NanogramsPerMilliliter returns the MassConcentration value in NanogramsPerMilliliter.
//
// 
func (a *MassConcentration) NanogramsPerMilliliter() float64 {
	if a.nanograms_per_milliliterLazy != nil {
		return *a.nanograms_per_milliliterLazy
	}
	nanograms_per_milliliter := a.convertFromBase(MassConcentrationNanogramPerMilliliter)
	a.nanograms_per_milliliterLazy = &nanograms_per_milliliter
	return nanograms_per_milliliter
}

// MicrogramsPerMilliliter returns the MassConcentration value in MicrogramsPerMilliliter.
//
// 
func (a *MassConcentration) MicrogramsPerMilliliter() float64 {
	if a.micrograms_per_milliliterLazy != nil {
		return *a.micrograms_per_milliliterLazy
	}
	micrograms_per_milliliter := a.convertFromBase(MassConcentrationMicrogramPerMilliliter)
	a.micrograms_per_milliliterLazy = &micrograms_per_milliliter
	return micrograms_per_milliliter
}

// MilligramsPerMilliliter returns the MassConcentration value in MilligramsPerMilliliter.
//
// 
func (a *MassConcentration) MilligramsPerMilliliter() float64 {
	if a.milligrams_per_milliliterLazy != nil {
		return *a.milligrams_per_milliliterLazy
	}
	milligrams_per_milliliter := a.convertFromBase(MassConcentrationMilligramPerMilliliter)
	a.milligrams_per_milliliterLazy = &milligrams_per_milliliter
	return milligrams_per_milliliter
}

// CentigramsPerMilliliter returns the MassConcentration value in CentigramsPerMilliliter.
//
// 
func (a *MassConcentration) CentigramsPerMilliliter() float64 {
	if a.centigrams_per_milliliterLazy != nil {
		return *a.centigrams_per_milliliterLazy
	}
	centigrams_per_milliliter := a.convertFromBase(MassConcentrationCentigramPerMilliliter)
	a.centigrams_per_milliliterLazy = &centigrams_per_milliliter
	return centigrams_per_milliliter
}

// DecigramsPerMilliliter returns the MassConcentration value in DecigramsPerMilliliter.
//
// 
func (a *MassConcentration) DecigramsPerMilliliter() float64 {
	if a.decigrams_per_milliliterLazy != nil {
		return *a.decigrams_per_milliliterLazy
	}
	decigrams_per_milliliter := a.convertFromBase(MassConcentrationDecigramPerMilliliter)
	a.decigrams_per_milliliterLazy = &decigrams_per_milliliter
	return decigrams_per_milliliter
}

// PicogramsPerDeciliter returns the MassConcentration value in PicogramsPerDeciliter.
//
// 
func (a *MassConcentration) PicogramsPerDeciliter() float64 {
	if a.picograms_per_deciliterLazy != nil {
		return *a.picograms_per_deciliterLazy
	}
	picograms_per_deciliter := a.convertFromBase(MassConcentrationPicogramPerDeciliter)
	a.picograms_per_deciliterLazy = &picograms_per_deciliter
	return picograms_per_deciliter
}

// NanogramsPerDeciliter returns the MassConcentration value in NanogramsPerDeciliter.
//
// 
func (a *MassConcentration) NanogramsPerDeciliter() float64 {
	if a.nanograms_per_deciliterLazy != nil {
		return *a.nanograms_per_deciliterLazy
	}
	nanograms_per_deciliter := a.convertFromBase(MassConcentrationNanogramPerDeciliter)
	a.nanograms_per_deciliterLazy = &nanograms_per_deciliter
	return nanograms_per_deciliter
}

// MicrogramsPerDeciliter returns the MassConcentration value in MicrogramsPerDeciliter.
//
// 
func (a *MassConcentration) MicrogramsPerDeciliter() float64 {
	if a.micrograms_per_deciliterLazy != nil {
		return *a.micrograms_per_deciliterLazy
	}
	micrograms_per_deciliter := a.convertFromBase(MassConcentrationMicrogramPerDeciliter)
	a.micrograms_per_deciliterLazy = &micrograms_per_deciliter
	return micrograms_per_deciliter
}

// MilligramsPerDeciliter returns the MassConcentration value in MilligramsPerDeciliter.
//
// 
func (a *MassConcentration) MilligramsPerDeciliter() float64 {
	if a.milligrams_per_deciliterLazy != nil {
		return *a.milligrams_per_deciliterLazy
	}
	milligrams_per_deciliter := a.convertFromBase(MassConcentrationMilligramPerDeciliter)
	a.milligrams_per_deciliterLazy = &milligrams_per_deciliter
	return milligrams_per_deciliter
}

// CentigramsPerDeciliter returns the MassConcentration value in CentigramsPerDeciliter.
//
// 
func (a *MassConcentration) CentigramsPerDeciliter() float64 {
	if a.centigrams_per_deciliterLazy != nil {
		return *a.centigrams_per_deciliterLazy
	}
	centigrams_per_deciliter := a.convertFromBase(MassConcentrationCentigramPerDeciliter)
	a.centigrams_per_deciliterLazy = &centigrams_per_deciliter
	return centigrams_per_deciliter
}

// DecigramsPerDeciliter returns the MassConcentration value in DecigramsPerDeciliter.
//
// 
func (a *MassConcentration) DecigramsPerDeciliter() float64 {
	if a.decigrams_per_deciliterLazy != nil {
		return *a.decigrams_per_deciliterLazy
	}
	decigrams_per_deciliter := a.convertFromBase(MassConcentrationDecigramPerDeciliter)
	a.decigrams_per_deciliterLazy = &decigrams_per_deciliter
	return decigrams_per_deciliter
}

// PicogramsPerLiter returns the MassConcentration value in PicogramsPerLiter.
//
// 
func (a *MassConcentration) PicogramsPerLiter() float64 {
	if a.picograms_per_literLazy != nil {
		return *a.picograms_per_literLazy
	}
	picograms_per_liter := a.convertFromBase(MassConcentrationPicogramPerLiter)
	a.picograms_per_literLazy = &picograms_per_liter
	return picograms_per_liter
}

// NanogramsPerLiter returns the MassConcentration value in NanogramsPerLiter.
//
// 
func (a *MassConcentration) NanogramsPerLiter() float64 {
	if a.nanograms_per_literLazy != nil {
		return *a.nanograms_per_literLazy
	}
	nanograms_per_liter := a.convertFromBase(MassConcentrationNanogramPerLiter)
	a.nanograms_per_literLazy = &nanograms_per_liter
	return nanograms_per_liter
}

// MicrogramsPerLiter returns the MassConcentration value in MicrogramsPerLiter.
//
// 
func (a *MassConcentration) MicrogramsPerLiter() float64 {
	if a.micrograms_per_literLazy != nil {
		return *a.micrograms_per_literLazy
	}
	micrograms_per_liter := a.convertFromBase(MassConcentrationMicrogramPerLiter)
	a.micrograms_per_literLazy = &micrograms_per_liter
	return micrograms_per_liter
}

// MilligramsPerLiter returns the MassConcentration value in MilligramsPerLiter.
//
// 
func (a *MassConcentration) MilligramsPerLiter() float64 {
	if a.milligrams_per_literLazy != nil {
		return *a.milligrams_per_literLazy
	}
	milligrams_per_liter := a.convertFromBase(MassConcentrationMilligramPerLiter)
	a.milligrams_per_literLazy = &milligrams_per_liter
	return milligrams_per_liter
}

// CentigramsPerLiter returns the MassConcentration value in CentigramsPerLiter.
//
// 
func (a *MassConcentration) CentigramsPerLiter() float64 {
	if a.centigrams_per_literLazy != nil {
		return *a.centigrams_per_literLazy
	}
	centigrams_per_liter := a.convertFromBase(MassConcentrationCentigramPerLiter)
	a.centigrams_per_literLazy = &centigrams_per_liter
	return centigrams_per_liter
}

// DecigramsPerLiter returns the MassConcentration value in DecigramsPerLiter.
//
// 
func (a *MassConcentration) DecigramsPerLiter() float64 {
	if a.decigrams_per_literLazy != nil {
		return *a.decigrams_per_literLazy
	}
	decigrams_per_liter := a.convertFromBase(MassConcentrationDecigramPerLiter)
	a.decigrams_per_literLazy = &decigrams_per_liter
	return decigrams_per_liter
}

// KilogramsPerLiter returns the MassConcentration value in KilogramsPerLiter.
//
// 
func (a *MassConcentration) KilogramsPerLiter() float64 {
	if a.kilograms_per_literLazy != nil {
		return *a.kilograms_per_literLazy
	}
	kilograms_per_liter := a.convertFromBase(MassConcentrationKilogramPerLiter)
	a.kilograms_per_literLazy = &kilograms_per_liter
	return kilograms_per_liter
}

// KilopoundsPerCubicInch returns the MassConcentration value in KilopoundsPerCubicInch.
//
// 
func (a *MassConcentration) KilopoundsPerCubicInch() float64 {
	if a.kilopounds_per_cubic_inchLazy != nil {
		return *a.kilopounds_per_cubic_inchLazy
	}
	kilopounds_per_cubic_inch := a.convertFromBase(MassConcentrationKilopoundPerCubicInch)
	a.kilopounds_per_cubic_inchLazy = &kilopounds_per_cubic_inch
	return kilopounds_per_cubic_inch
}

// KilopoundsPerCubicFoot returns the MassConcentration value in KilopoundsPerCubicFoot.
//
// 
func (a *MassConcentration) KilopoundsPerCubicFoot() float64 {
	if a.kilopounds_per_cubic_footLazy != nil {
		return *a.kilopounds_per_cubic_footLazy
	}
	kilopounds_per_cubic_foot := a.convertFromBase(MassConcentrationKilopoundPerCubicFoot)
	a.kilopounds_per_cubic_footLazy = &kilopounds_per_cubic_foot
	return kilopounds_per_cubic_foot
}


// ToDto creates a MassConcentrationDto representation from the MassConcentration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerCubicMeter by default.
func (a *MassConcentration) ToDto(holdInUnit *MassConcentrationUnits) MassConcentrationDto {
	if holdInUnit == nil {
		defaultUnit := MassConcentrationKilogramPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MassConcentrationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MassConcentration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerCubicMeter by default.
func (a *MassConcentration) ToDtoJSON(holdInUnit *MassConcentrationUnits) ([]byte, error) {
	// Convert to MassConcentrationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MassConcentration to a specific unit value.
// The function uses the provided unit type (MassConcentrationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MassConcentration) Convert(toUnit MassConcentrationUnits) float64 {
	switch toUnit { 
    case MassConcentrationGramPerCubicMillimeter:
		return a.GramsPerCubicMillimeter()
    case MassConcentrationGramPerCubicCentimeter:
		return a.GramsPerCubicCentimeter()
    case MassConcentrationGramPerCubicMeter:
		return a.GramsPerCubicMeter()
    case MassConcentrationGramPerMicroliter:
		return a.GramsPerMicroliter()
    case MassConcentrationGramPerMilliliter:
		return a.GramsPerMilliliter()
    case MassConcentrationGramPerDeciliter:
		return a.GramsPerDeciliter()
    case MassConcentrationGramPerLiter:
		return a.GramsPerLiter()
    case MassConcentrationTonnePerCubicMillimeter:
		return a.TonnesPerCubicMillimeter()
    case MassConcentrationTonnePerCubicCentimeter:
		return a.TonnesPerCubicCentimeter()
    case MassConcentrationTonnePerCubicMeter:
		return a.TonnesPerCubicMeter()
    case MassConcentrationPoundPerCubicInch:
		return a.PoundsPerCubicInch()
    case MassConcentrationPoundPerCubicFoot:
		return a.PoundsPerCubicFoot()
    case MassConcentrationSlugPerCubicFoot:
		return a.SlugsPerCubicFoot()
    case MassConcentrationPoundPerUSGallon:
		return a.PoundsPerUSGallon()
    case MassConcentrationOuncePerUSGallon:
		return a.OuncesPerUSGallon()
    case MassConcentrationOuncePerImperialGallon:
		return a.OuncesPerImperialGallon()
    case MassConcentrationPoundPerImperialGallon:
		return a.PoundsPerImperialGallon()
    case MassConcentrationKilogramPerCubicMillimeter:
		return a.KilogramsPerCubicMillimeter()
    case MassConcentrationKilogramPerCubicCentimeter:
		return a.KilogramsPerCubicCentimeter()
    case MassConcentrationKilogramPerCubicMeter:
		return a.KilogramsPerCubicMeter()
    case MassConcentrationMilligramPerCubicMeter:
		return a.MilligramsPerCubicMeter()
    case MassConcentrationMicrogramPerCubicMeter:
		return a.MicrogramsPerCubicMeter()
    case MassConcentrationPicogramPerMicroliter:
		return a.PicogramsPerMicroliter()
    case MassConcentrationNanogramPerMicroliter:
		return a.NanogramsPerMicroliter()
    case MassConcentrationMicrogramPerMicroliter:
		return a.MicrogramsPerMicroliter()
    case MassConcentrationMilligramPerMicroliter:
		return a.MilligramsPerMicroliter()
    case MassConcentrationCentigramPerMicroliter:
		return a.CentigramsPerMicroliter()
    case MassConcentrationDecigramPerMicroliter:
		return a.DecigramsPerMicroliter()
    case MassConcentrationPicogramPerMilliliter:
		return a.PicogramsPerMilliliter()
    case MassConcentrationNanogramPerMilliliter:
		return a.NanogramsPerMilliliter()
    case MassConcentrationMicrogramPerMilliliter:
		return a.MicrogramsPerMilliliter()
    case MassConcentrationMilligramPerMilliliter:
		return a.MilligramsPerMilliliter()
    case MassConcentrationCentigramPerMilliliter:
		return a.CentigramsPerMilliliter()
    case MassConcentrationDecigramPerMilliliter:
		return a.DecigramsPerMilliliter()
    case MassConcentrationPicogramPerDeciliter:
		return a.PicogramsPerDeciliter()
    case MassConcentrationNanogramPerDeciliter:
		return a.NanogramsPerDeciliter()
    case MassConcentrationMicrogramPerDeciliter:
		return a.MicrogramsPerDeciliter()
    case MassConcentrationMilligramPerDeciliter:
		return a.MilligramsPerDeciliter()
    case MassConcentrationCentigramPerDeciliter:
		return a.CentigramsPerDeciliter()
    case MassConcentrationDecigramPerDeciliter:
		return a.DecigramsPerDeciliter()
    case MassConcentrationPicogramPerLiter:
		return a.PicogramsPerLiter()
    case MassConcentrationNanogramPerLiter:
		return a.NanogramsPerLiter()
    case MassConcentrationMicrogramPerLiter:
		return a.MicrogramsPerLiter()
    case MassConcentrationMilligramPerLiter:
		return a.MilligramsPerLiter()
    case MassConcentrationCentigramPerLiter:
		return a.CentigramsPerLiter()
    case MassConcentrationDecigramPerLiter:
		return a.DecigramsPerLiter()
    case MassConcentrationKilogramPerLiter:
		return a.KilogramsPerLiter()
    case MassConcentrationKilopoundPerCubicInch:
		return a.KilopoundsPerCubicInch()
    case MassConcentrationKilopoundPerCubicFoot:
		return a.KilopoundsPerCubicFoot()
	default:
		return math.NaN()
	}
}

func (a *MassConcentration) convertFromBase(toUnit MassConcentrationUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassConcentrationGramPerCubicMillimeter:
		return (value * 1e-6) 
	case MassConcentrationGramPerCubicCentimeter:
		return (value * 1e-3) 
	case MassConcentrationGramPerCubicMeter:
		return (value * 1e3) 
	case MassConcentrationGramPerMicroliter:
		return (value * 1e-6) 
	case MassConcentrationGramPerMilliliter:
		return (value * 1e-3) 
	case MassConcentrationGramPerDeciliter:
		return (value * 1e-1) 
	case MassConcentrationGramPerLiter:
		return (value) 
	case MassConcentrationTonnePerCubicMillimeter:
		return (value * 1e-12) 
	case MassConcentrationTonnePerCubicCentimeter:
		return (value * 1e-9) 
	case MassConcentrationTonnePerCubicMeter:
		return (value * 0.001) 
	case MassConcentrationPoundPerCubicInch:
		return (value * 3.6127298147753e-5) 
	case MassConcentrationPoundPerCubicFoot:
		return (value * 0.062427961) 
	case MassConcentrationSlugPerCubicFoot:
		return (value * 0.00194032033) 
	case MassConcentrationPoundPerUSGallon:
		return (value / 1.19826427e2) 
	case MassConcentrationOuncePerUSGallon:
		return (value * 0.1335264711843) 
	case MassConcentrationOuncePerImperialGallon:
		return (value * 0.1603586720609) 
	case MassConcentrationPoundPerImperialGallon:
		return (value / 9.9776398e1) 
	case MassConcentrationKilogramPerCubicMillimeter:
		return ((value * 1e-6) / 1000.0) 
	case MassConcentrationKilogramPerCubicCentimeter:
		return ((value * 1e-3) / 1000.0) 
	case MassConcentrationKilogramPerCubicMeter:
		return ((value * 1e3) / 1000.0) 
	case MassConcentrationMilligramPerCubicMeter:
		return ((value * 1e3) / 0.001) 
	case MassConcentrationMicrogramPerCubicMeter:
		return ((value * 1e3) / 1e-06) 
	case MassConcentrationPicogramPerMicroliter:
		return ((value * 1e-6) / 1e-12) 
	case MassConcentrationNanogramPerMicroliter:
		return ((value * 1e-6) / 1e-09) 
	case MassConcentrationMicrogramPerMicroliter:
		return ((value * 1e-6) / 1e-06) 
	case MassConcentrationMilligramPerMicroliter:
		return ((value * 1e-6) / 0.001) 
	case MassConcentrationCentigramPerMicroliter:
		return ((value * 1e-6) / 0.01) 
	case MassConcentrationDecigramPerMicroliter:
		return ((value * 1e-6) / 0.1) 
	case MassConcentrationPicogramPerMilliliter:
		return ((value * 1e-3) / 1e-12) 
	case MassConcentrationNanogramPerMilliliter:
		return ((value * 1e-3) / 1e-09) 
	case MassConcentrationMicrogramPerMilliliter:
		return ((value * 1e-3) / 1e-06) 
	case MassConcentrationMilligramPerMilliliter:
		return ((value * 1e-3) / 0.001) 
	case MassConcentrationCentigramPerMilliliter:
		return ((value * 1e-3) / 0.01) 
	case MassConcentrationDecigramPerMilliliter:
		return ((value * 1e-3) / 0.1) 
	case MassConcentrationPicogramPerDeciliter:
		return ((value * 1e-1) / 1e-12) 
	case MassConcentrationNanogramPerDeciliter:
		return ((value * 1e-1) / 1e-09) 
	case MassConcentrationMicrogramPerDeciliter:
		return ((value * 1e-1) / 1e-06) 
	case MassConcentrationMilligramPerDeciliter:
		return ((value * 1e-1) / 0.001) 
	case MassConcentrationCentigramPerDeciliter:
		return ((value * 1e-1) / 0.01) 
	case MassConcentrationDecigramPerDeciliter:
		return ((value * 1e-1) / 0.1) 
	case MassConcentrationPicogramPerLiter:
		return ((value) / 1e-12) 
	case MassConcentrationNanogramPerLiter:
		return ((value) / 1e-09) 
	case MassConcentrationMicrogramPerLiter:
		return ((value) / 1e-06) 
	case MassConcentrationMilligramPerLiter:
		return ((value) / 0.001) 
	case MassConcentrationCentigramPerLiter:
		return ((value) / 0.01) 
	case MassConcentrationDecigramPerLiter:
		return ((value) / 0.1) 
	case MassConcentrationKilogramPerLiter:
		return ((value) / 1000.0) 
	case MassConcentrationKilopoundPerCubicInch:
		return ((value * 3.6127298147753e-5) / 1000.0) 
	case MassConcentrationKilopoundPerCubicFoot:
		return ((value * 0.062427961) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *MassConcentration) convertToBase(value float64, fromUnit MassConcentrationUnits) float64 {
	switch fromUnit { 
	case MassConcentrationGramPerCubicMillimeter:
		return (value / 1e-6) 
	case MassConcentrationGramPerCubicCentimeter:
		return (value / 1e-3) 
	case MassConcentrationGramPerCubicMeter:
		return (value / 1e3) 
	case MassConcentrationGramPerMicroliter:
		return (value / 1e-6) 
	case MassConcentrationGramPerMilliliter:
		return (value / 1e-3) 
	case MassConcentrationGramPerDeciliter:
		return (value / 1e-1) 
	case MassConcentrationGramPerLiter:
		return (value) 
	case MassConcentrationTonnePerCubicMillimeter:
		return (value / 1e-12) 
	case MassConcentrationTonnePerCubicCentimeter:
		return (value / 1e-9) 
	case MassConcentrationTonnePerCubicMeter:
		return (value / 0.001) 
	case MassConcentrationPoundPerCubicInch:
		return (value / 3.6127298147753e-5) 
	case MassConcentrationPoundPerCubicFoot:
		return (value / 0.062427961) 
	case MassConcentrationSlugPerCubicFoot:
		return (value * 515.378818) 
	case MassConcentrationPoundPerUSGallon:
		return (value * 1.19826427e2) 
	case MassConcentrationOuncePerUSGallon:
		return ( value / 0.1335264711843) 
	case MassConcentrationOuncePerImperialGallon:
		return ( value / 0.1603586720609) 
	case MassConcentrationPoundPerImperialGallon:
		return (value * 9.9776398e1) 
	case MassConcentrationKilogramPerCubicMillimeter:
		return ((value / 1e-6) * 1000.0) 
	case MassConcentrationKilogramPerCubicCentimeter:
		return ((value / 1e-3) * 1000.0) 
	case MassConcentrationKilogramPerCubicMeter:
		return ((value / 1e3) * 1000.0) 
	case MassConcentrationMilligramPerCubicMeter:
		return ((value / 1e3) * 0.001) 
	case MassConcentrationMicrogramPerCubicMeter:
		return ((value / 1e3) * 1e-06) 
	case MassConcentrationPicogramPerMicroliter:
		return ((value / 1e-6) * 1e-12) 
	case MassConcentrationNanogramPerMicroliter:
		return ((value / 1e-6) * 1e-09) 
	case MassConcentrationMicrogramPerMicroliter:
		return ((value / 1e-6) * 1e-06) 
	case MassConcentrationMilligramPerMicroliter:
		return ((value / 1e-6) * 0.001) 
	case MassConcentrationCentigramPerMicroliter:
		return ((value / 1e-6) * 0.01) 
	case MassConcentrationDecigramPerMicroliter:
		return ((value / 1e-6) * 0.1) 
	case MassConcentrationPicogramPerMilliliter:
		return ((value / 1e-3) * 1e-12) 
	case MassConcentrationNanogramPerMilliliter:
		return ((value / 1e-3) * 1e-09) 
	case MassConcentrationMicrogramPerMilliliter:
		return ((value / 1e-3) * 1e-06) 
	case MassConcentrationMilligramPerMilliliter:
		return ((value / 1e-3) * 0.001) 
	case MassConcentrationCentigramPerMilliliter:
		return ((value / 1e-3) * 0.01) 
	case MassConcentrationDecigramPerMilliliter:
		return ((value / 1e-3) * 0.1) 
	case MassConcentrationPicogramPerDeciliter:
		return ((value / 1e-1) * 1e-12) 
	case MassConcentrationNanogramPerDeciliter:
		return ((value / 1e-1) * 1e-09) 
	case MassConcentrationMicrogramPerDeciliter:
		return ((value / 1e-1) * 1e-06) 
	case MassConcentrationMilligramPerDeciliter:
		return ((value / 1e-1) * 0.001) 
	case MassConcentrationCentigramPerDeciliter:
		return ((value / 1e-1) * 0.01) 
	case MassConcentrationDecigramPerDeciliter:
		return ((value / 1e-1) * 0.1) 
	case MassConcentrationPicogramPerLiter:
		return ((value) * 1e-12) 
	case MassConcentrationNanogramPerLiter:
		return ((value) * 1e-09) 
	case MassConcentrationMicrogramPerLiter:
		return ((value) * 1e-06) 
	case MassConcentrationMilligramPerLiter:
		return ((value) * 0.001) 
	case MassConcentrationCentigramPerLiter:
		return ((value) * 0.01) 
	case MassConcentrationDecigramPerLiter:
		return ((value) * 0.1) 
	case MassConcentrationKilogramPerLiter:
		return ((value) * 1000.0) 
	case MassConcentrationKilopoundPerCubicInch:
		return ((value / 3.6127298147753e-5) * 1000.0) 
	case MassConcentrationKilopoundPerCubicFoot:
		return ((value / 0.062427961) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MassConcentration in the default unit (KilogramPerCubicMeter),
// formatted to two decimal places.
func (a MassConcentration) String() string {
	return a.ToString(MassConcentrationKilogramPerCubicMeter, 2)
}

// ToString formats the MassConcentration value as a string with the specified unit and fractional digits.
// It converts the MassConcentration to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MassConcentration value will be converted (e.g., KilogramPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MassConcentration with the unit abbreviation.
func (a *MassConcentration) ToString(unit MassConcentrationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MassConcentration) getUnitAbbreviation(unit MassConcentrationUnits) string {
	switch unit { 
	case MassConcentrationGramPerCubicMillimeter:
		return "g/mm³" 
	case MassConcentrationGramPerCubicCentimeter:
		return "g/cm³" 
	case MassConcentrationGramPerCubicMeter:
		return "g/m³" 
	case MassConcentrationGramPerMicroliter:
		return "g/μL" 
	case MassConcentrationGramPerMilliliter:
		return "g/mL" 
	case MassConcentrationGramPerDeciliter:
		return "g/dL" 
	case MassConcentrationGramPerLiter:
		return "g/L" 
	case MassConcentrationTonnePerCubicMillimeter:
		return "t/mm³" 
	case MassConcentrationTonnePerCubicCentimeter:
		return "t/cm³" 
	case MassConcentrationTonnePerCubicMeter:
		return "t/m³" 
	case MassConcentrationPoundPerCubicInch:
		return "lb/in³" 
	case MassConcentrationPoundPerCubicFoot:
		return "lb/ft³" 
	case MassConcentrationSlugPerCubicFoot:
		return "slug/ft³" 
	case MassConcentrationPoundPerUSGallon:
		return "ppg (U.S.)" 
	case MassConcentrationOuncePerUSGallon:
		return "oz/gal (U.S.)" 
	case MassConcentrationOuncePerImperialGallon:
		return "oz/gal (imp.)" 
	case MassConcentrationPoundPerImperialGallon:
		return "ppg (imp.)" 
	case MassConcentrationKilogramPerCubicMillimeter:
		return "kg/mm³" 
	case MassConcentrationKilogramPerCubicCentimeter:
		return "kg/cm³" 
	case MassConcentrationKilogramPerCubicMeter:
		return "kg/m³" 
	case MassConcentrationMilligramPerCubicMeter:
		return "mg/m³" 
	case MassConcentrationMicrogramPerCubicMeter:
		return "μg/m³" 
	case MassConcentrationPicogramPerMicroliter:
		return "pg/μL" 
	case MassConcentrationNanogramPerMicroliter:
		return "ng/μL" 
	case MassConcentrationMicrogramPerMicroliter:
		return "μg/μL" 
	case MassConcentrationMilligramPerMicroliter:
		return "mg/μL" 
	case MassConcentrationCentigramPerMicroliter:
		return "cg/μL" 
	case MassConcentrationDecigramPerMicroliter:
		return "dg/μL" 
	case MassConcentrationPicogramPerMilliliter:
		return "pg/mL" 
	case MassConcentrationNanogramPerMilliliter:
		return "ng/mL" 
	case MassConcentrationMicrogramPerMilliliter:
		return "μg/mL" 
	case MassConcentrationMilligramPerMilliliter:
		return "mg/mL" 
	case MassConcentrationCentigramPerMilliliter:
		return "cg/mL" 
	case MassConcentrationDecigramPerMilliliter:
		return "dg/mL" 
	case MassConcentrationPicogramPerDeciliter:
		return "pg/dL" 
	case MassConcentrationNanogramPerDeciliter:
		return "ng/dL" 
	case MassConcentrationMicrogramPerDeciliter:
		return "μg/dL" 
	case MassConcentrationMilligramPerDeciliter:
		return "mg/dL" 
	case MassConcentrationCentigramPerDeciliter:
		return "cg/dL" 
	case MassConcentrationDecigramPerDeciliter:
		return "dg/dL" 
	case MassConcentrationPicogramPerLiter:
		return "pg/L" 
	case MassConcentrationNanogramPerLiter:
		return "ng/L" 
	case MassConcentrationMicrogramPerLiter:
		return "μg/L" 
	case MassConcentrationMilligramPerLiter:
		return "mg/L" 
	case MassConcentrationCentigramPerLiter:
		return "cg/L" 
	case MassConcentrationDecigramPerLiter:
		return "dg/L" 
	case MassConcentrationKilogramPerLiter:
		return "kg/L" 
	case MassConcentrationKilopoundPerCubicInch:
		return "klb/in³" 
	case MassConcentrationKilopoundPerCubicFoot:
		return "klb/ft³" 
	default:
		return ""
	}
}

// Equals checks if the given MassConcentration is equal to the current MassConcentration.
//
// Parameters:
//    other: The MassConcentration to compare against.
//
// Returns:
//    bool: Returns true if both MassConcentration are equal, false otherwise.
func (a *MassConcentration) Equals(other *MassConcentration) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MassConcentration with another MassConcentration.
// It returns -1 if the current MassConcentration is less than the other MassConcentration, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MassConcentration to compare against.
//
// Returns:
//    int: -1 if the current MassConcentration is less, 1 if greater, and 0 if equal.
func (a *MassConcentration) CompareTo(other *MassConcentration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MassConcentration to the current MassConcentration and returns the result.
// The result is a new MassConcentration instance with the sum of the values.
//
// Parameters:
//    other: The MassConcentration to add to the current MassConcentration.
//
// Returns:
//    *MassConcentration: A new MassConcentration instance representing the sum of both MassConcentration.
func (a *MassConcentration) Add(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MassConcentration from the current MassConcentration and returns the result.
// The result is a new MassConcentration instance with the difference of the values.
//
// Parameters:
//    other: The MassConcentration to subtract from the current MassConcentration.
//
// Returns:
//    *MassConcentration: A new MassConcentration instance representing the difference of both MassConcentration.
func (a *MassConcentration) Subtract(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MassConcentration by the given MassConcentration and returns the result.
// The result is a new MassConcentration instance with the product of the values.
//
// Parameters:
//    other: The MassConcentration to multiply with the current MassConcentration.
//
// Returns:
//    *MassConcentration: A new MassConcentration instance representing the product of both MassConcentration.
func (a *MassConcentration) Multiply(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value * other.BaseValue()}
}

// Divide divides the current MassConcentration by the given MassConcentration and returns the result.
// The result is a new MassConcentration instance with the quotient of the values.
//
// Parameters:
//    other: The MassConcentration to divide the current MassConcentration by.
//
// Returns:
//    *MassConcentration: A new MassConcentration instance representing the quotient of both MassConcentration.
func (a *MassConcentration) Divide(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value / other.BaseValue()}
}