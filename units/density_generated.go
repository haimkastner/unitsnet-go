// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// DensityUnits defines various units of Density.
type DensityUnits string

const (
    
        // 
        DensityGramPerCubicMillimeter DensityUnits = "GramPerCubicMillimeter"
        // 
        DensityGramPerCubicCentimeter DensityUnits = "GramPerCubicCentimeter"
        // 
        DensityGramPerCubicMeter DensityUnits = "GramPerCubicMeter"
        // 
        DensityPoundPerCubicInch DensityUnits = "PoundPerCubicInch"
        // 
        DensityPoundPerCubicFoot DensityUnits = "PoundPerCubicFoot"
        // Calculated from the definition of <a href="https://en.wikipedia.org/wiki/Pound_(mass)">pound</a> and <a href="https://en.wikipedia.org/wiki/Yard">yard</a> compared to metric kilogram and meter.
        DensityPoundPerCubicYard DensityUnits = "PoundPerCubicYard"
        // 
        DensityTonnePerCubicMillimeter DensityUnits = "TonnePerCubicMillimeter"
        // 
        DensityTonnePerCubicCentimeter DensityUnits = "TonnePerCubicCentimeter"
        // 
        DensityTonnePerCubicMeter DensityUnits = "TonnePerCubicMeter"
        // 
        DensitySlugPerCubicFoot DensityUnits = "SlugPerCubicFoot"
        // 
        DensityGramPerLiter DensityUnits = "GramPerLiter"
        // 
        DensityGramPerDeciliter DensityUnits = "GramPerDeciliter"
        // 
        DensityGramPerMilliliter DensityUnits = "GramPerMilliliter"
        // 
        DensityPoundPerUSGallon DensityUnits = "PoundPerUSGallon"
        // 
        DensityPoundPerImperialGallon DensityUnits = "PoundPerImperialGallon"
        // 
        DensityKilogramPerLiter DensityUnits = "KilogramPerLiter"
        // 
        DensityTonnePerCubicFoot DensityUnits = "TonnePerCubicFoot"
        // 
        DensityTonnePerCubicInch DensityUnits = "TonnePerCubicInch"
        // 
        DensityGramPerCubicFoot DensityUnits = "GramPerCubicFoot"
        // 
        DensityGramPerCubicInch DensityUnits = "GramPerCubicInch"
        // 
        DensityPoundPerCubicMeter DensityUnits = "PoundPerCubicMeter"
        // 
        DensityPoundPerCubicCentimeter DensityUnits = "PoundPerCubicCentimeter"
        // 
        DensityPoundPerCubicMillimeter DensityUnits = "PoundPerCubicMillimeter"
        // 
        DensitySlugPerCubicMeter DensityUnits = "SlugPerCubicMeter"
        // 
        DensitySlugPerCubicCentimeter DensityUnits = "SlugPerCubicCentimeter"
        // 
        DensitySlugPerCubicMillimeter DensityUnits = "SlugPerCubicMillimeter"
        // 
        DensitySlugPerCubicInch DensityUnits = "SlugPerCubicInch"
        // 
        DensityKilogramPerCubicMillimeter DensityUnits = "KilogramPerCubicMillimeter"
        // 
        DensityKilogramPerCubicCentimeter DensityUnits = "KilogramPerCubicCentimeter"
        // 
        DensityKilogramPerCubicMeter DensityUnits = "KilogramPerCubicMeter"
        // 
        DensityMilligramPerCubicMeter DensityUnits = "MilligramPerCubicMeter"
        // 
        DensityMicrogramPerCubicMeter DensityUnits = "MicrogramPerCubicMeter"
        // 
        DensityKilopoundPerCubicInch DensityUnits = "KilopoundPerCubicInch"
        // 
        DensityKilopoundPerCubicFoot DensityUnits = "KilopoundPerCubicFoot"
        // 
        DensityKilopoundPerCubicYard DensityUnits = "KilopoundPerCubicYard"
        // 
        DensityFemtogramPerLiter DensityUnits = "FemtogramPerLiter"
        // 
        DensityPicogramPerLiter DensityUnits = "PicogramPerLiter"
        // 
        DensityNanogramPerLiter DensityUnits = "NanogramPerLiter"
        // 
        DensityMicrogramPerLiter DensityUnits = "MicrogramPerLiter"
        // 
        DensityMilligramPerLiter DensityUnits = "MilligramPerLiter"
        // 
        DensityCentigramPerLiter DensityUnits = "CentigramPerLiter"
        // 
        DensityDecigramPerLiter DensityUnits = "DecigramPerLiter"
        // 
        DensityFemtogramPerDeciliter DensityUnits = "FemtogramPerDeciliter"
        // 
        DensityPicogramPerDeciliter DensityUnits = "PicogramPerDeciliter"
        // 
        DensityNanogramPerDeciliter DensityUnits = "NanogramPerDeciliter"
        // 
        DensityMicrogramPerDeciliter DensityUnits = "MicrogramPerDeciliter"
        // 
        DensityMilligramPerDeciliter DensityUnits = "MilligramPerDeciliter"
        // 
        DensityCentigramPerDeciliter DensityUnits = "CentigramPerDeciliter"
        // 
        DensityDecigramPerDeciliter DensityUnits = "DecigramPerDeciliter"
        // 
        DensityFemtogramPerMilliliter DensityUnits = "FemtogramPerMilliliter"
        // 
        DensityPicogramPerMilliliter DensityUnits = "PicogramPerMilliliter"
        // 
        DensityNanogramPerMilliliter DensityUnits = "NanogramPerMilliliter"
        // 
        DensityMicrogramPerMilliliter DensityUnits = "MicrogramPerMilliliter"
        // 
        DensityMilligramPerMilliliter DensityUnits = "MilligramPerMilliliter"
        // 
        DensityCentigramPerMilliliter DensityUnits = "CentigramPerMilliliter"
        // 
        DensityDecigramPerMilliliter DensityUnits = "DecigramPerMilliliter"
)

var internalDensityUnitsMap = map[DensityUnits]bool{
	
	DensityGramPerCubicMillimeter: true,
	DensityGramPerCubicCentimeter: true,
	DensityGramPerCubicMeter: true,
	DensityPoundPerCubicInch: true,
	DensityPoundPerCubicFoot: true,
	DensityPoundPerCubicYard: true,
	DensityTonnePerCubicMillimeter: true,
	DensityTonnePerCubicCentimeter: true,
	DensityTonnePerCubicMeter: true,
	DensitySlugPerCubicFoot: true,
	DensityGramPerLiter: true,
	DensityGramPerDeciliter: true,
	DensityGramPerMilliliter: true,
	DensityPoundPerUSGallon: true,
	DensityPoundPerImperialGallon: true,
	DensityKilogramPerLiter: true,
	DensityTonnePerCubicFoot: true,
	DensityTonnePerCubicInch: true,
	DensityGramPerCubicFoot: true,
	DensityGramPerCubicInch: true,
	DensityPoundPerCubicMeter: true,
	DensityPoundPerCubicCentimeter: true,
	DensityPoundPerCubicMillimeter: true,
	DensitySlugPerCubicMeter: true,
	DensitySlugPerCubicCentimeter: true,
	DensitySlugPerCubicMillimeter: true,
	DensitySlugPerCubicInch: true,
	DensityKilogramPerCubicMillimeter: true,
	DensityKilogramPerCubicCentimeter: true,
	DensityKilogramPerCubicMeter: true,
	DensityMilligramPerCubicMeter: true,
	DensityMicrogramPerCubicMeter: true,
	DensityKilopoundPerCubicInch: true,
	DensityKilopoundPerCubicFoot: true,
	DensityKilopoundPerCubicYard: true,
	DensityFemtogramPerLiter: true,
	DensityPicogramPerLiter: true,
	DensityNanogramPerLiter: true,
	DensityMicrogramPerLiter: true,
	DensityMilligramPerLiter: true,
	DensityCentigramPerLiter: true,
	DensityDecigramPerLiter: true,
	DensityFemtogramPerDeciliter: true,
	DensityPicogramPerDeciliter: true,
	DensityNanogramPerDeciliter: true,
	DensityMicrogramPerDeciliter: true,
	DensityMilligramPerDeciliter: true,
	DensityCentigramPerDeciliter: true,
	DensityDecigramPerDeciliter: true,
	DensityFemtogramPerMilliliter: true,
	DensityPicogramPerMilliliter: true,
	DensityNanogramPerMilliliter: true,
	DensityMicrogramPerMilliliter: true,
	DensityMilligramPerMilliliter: true,
	DensityCentigramPerMilliliter: true,
	DensityDecigramPerMilliliter: true,
}

// DensityDto represents a Density measurement with a numerical value and its corresponding unit.
type DensityDto struct {
    // Value is the numerical representation of the Density.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Density, as defined in the DensityUnits enumeration.
	Unit  DensityUnits `json:"unit" validate:"required,oneof=GramPerCubicMillimeter GramPerCubicCentimeter GramPerCubicMeter PoundPerCubicInch PoundPerCubicFoot PoundPerCubicYard TonnePerCubicMillimeter TonnePerCubicCentimeter TonnePerCubicMeter SlugPerCubicFoot GramPerLiter GramPerDeciliter GramPerMilliliter PoundPerUSGallon PoundPerImperialGallon KilogramPerLiter TonnePerCubicFoot TonnePerCubicInch GramPerCubicFoot GramPerCubicInch PoundPerCubicMeter PoundPerCubicCentimeter PoundPerCubicMillimeter SlugPerCubicMeter SlugPerCubicCentimeter SlugPerCubicMillimeter SlugPerCubicInch KilogramPerCubicMillimeter KilogramPerCubicCentimeter KilogramPerCubicMeter MilligramPerCubicMeter MicrogramPerCubicMeter KilopoundPerCubicInch KilopoundPerCubicFoot KilopoundPerCubicYard FemtogramPerLiter PicogramPerLiter NanogramPerLiter MicrogramPerLiter MilligramPerLiter CentigramPerLiter DecigramPerLiter FemtogramPerDeciliter PicogramPerDeciliter NanogramPerDeciliter MicrogramPerDeciliter MilligramPerDeciliter CentigramPerDeciliter DecigramPerDeciliter FemtogramPerMilliliter PicogramPerMilliliter NanogramPerMilliliter MicrogramPerMilliliter MilligramPerMilliliter CentigramPerMilliliter DecigramPerMilliliter"`
}

// DensityDtoFactory groups methods for creating and serializing DensityDto objects.
type DensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a DensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf DensityDtoFactory) FromJSON(data []byte) (*DensityDto, error) {
	a := DensityDto{}

    // Parse JSON into DensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a DensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a DensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Density represents a measurement in a of Density.
//
// The density, or more precisely, the volumetric mass density, of a substance is its mass per unit volume.
type Density struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_cubic_millimeterLazy *float64 
    grams_per_cubic_centimeterLazy *float64 
    grams_per_cubic_meterLazy *float64 
    pounds_per_cubic_inchLazy *float64 
    pounds_per_cubic_footLazy *float64 
    pounds_per_cubic_yardLazy *float64 
    tonnes_per_cubic_millimeterLazy *float64 
    tonnes_per_cubic_centimeterLazy *float64 
    tonnes_per_cubic_meterLazy *float64 
    slugs_per_cubic_footLazy *float64 
    grams_per_literLazy *float64 
    grams_per_deci_literLazy *float64 
    grams_per_milliliterLazy *float64 
    pounds_per_us_gallonLazy *float64 
    pounds_per_imperial_gallonLazy *float64 
    kilograms_per_literLazy *float64 
    tonnes_per_cubic_footLazy *float64 
    tonnes_per_cubic_inchLazy *float64 
    grams_per_cubic_footLazy *float64 
    grams_per_cubic_inchLazy *float64 
    pounds_per_cubic_meterLazy *float64 
    pounds_per_cubic_centimeterLazy *float64 
    pounds_per_cubic_millimeterLazy *float64 
    slugs_per_cubic_meterLazy *float64 
    slugs_per_cubic_centimeterLazy *float64 
    slugs_per_cubic_millimeterLazy *float64 
    slugs_per_cubic_inchLazy *float64 
    kilograms_per_cubic_millimeterLazy *float64 
    kilograms_per_cubic_centimeterLazy *float64 
    kilograms_per_cubic_meterLazy *float64 
    milligrams_per_cubic_meterLazy *float64 
    micrograms_per_cubic_meterLazy *float64 
    kilopounds_per_cubic_inchLazy *float64 
    kilopounds_per_cubic_footLazy *float64 
    kilopounds_per_cubic_yardLazy *float64 
    femtograms_per_literLazy *float64 
    picograms_per_literLazy *float64 
    nanograms_per_literLazy *float64 
    micrograms_per_literLazy *float64 
    milligrams_per_literLazy *float64 
    centigrams_per_literLazy *float64 
    decigrams_per_literLazy *float64 
    femtograms_per_deci_literLazy *float64 
    picograms_per_deci_literLazy *float64 
    nanograms_per_deci_literLazy *float64 
    micrograms_per_deci_literLazy *float64 
    milligrams_per_deci_literLazy *float64 
    centigrams_per_deci_literLazy *float64 
    decigrams_per_deci_literLazy *float64 
    femtograms_per_milliliterLazy *float64 
    picograms_per_milliliterLazy *float64 
    nanograms_per_milliliterLazy *float64 
    micrograms_per_milliliterLazy *float64 
    milligrams_per_milliliterLazy *float64 
    centigrams_per_milliliterLazy *float64 
    decigrams_per_milliliterLazy *float64 
}

// DensityFactory groups methods for creating Density instances.
type DensityFactory struct{}

// CreateDensity creates a new Density instance from the given value and unit.
func (uf DensityFactory) CreateDensity(value float64, unit DensityUnits) (*Density, error) {
	return newDensity(value, unit)
}

// FromDto converts a DensityDto to a Density instance.
func (uf DensityFactory) FromDto(dto DensityDto) (*Density, error) {
	return newDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Density instance.
func (uf DensityFactory) FromDtoJSON(data []byte) (*Density, error) {
	unitDto, err := DensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DensityDto from JSON: %w", err)
	}
	return DensityFactory{}.FromDto(*unitDto)
}


// FromGramsPerCubicMillimeter creates a new Density instance from a value in GramsPerCubicMillimeter.
func (uf DensityFactory) FromGramsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicMillimeter)
}

// FromGramsPerCubicCentimeter creates a new Density instance from a value in GramsPerCubicCentimeter.
func (uf DensityFactory) FromGramsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicCentimeter)
}

// FromGramsPerCubicMeter creates a new Density instance from a value in GramsPerCubicMeter.
func (uf DensityFactory) FromGramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicMeter)
}

// FromPoundsPerCubicInch creates a new Density instance from a value in PoundsPerCubicInch.
func (uf DensityFactory) FromPoundsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicInch)
}

// FromPoundsPerCubicFoot creates a new Density instance from a value in PoundsPerCubicFoot.
func (uf DensityFactory) FromPoundsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicFoot)
}

// FromPoundsPerCubicYard creates a new Density instance from a value in PoundsPerCubicYard.
func (uf DensityFactory) FromPoundsPerCubicYard(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicYard)
}

// FromTonnesPerCubicMillimeter creates a new Density instance from a value in TonnesPerCubicMillimeter.
func (uf DensityFactory) FromTonnesPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicMillimeter)
}

// FromTonnesPerCubicCentimeter creates a new Density instance from a value in TonnesPerCubicCentimeter.
func (uf DensityFactory) FromTonnesPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicCentimeter)
}

// FromTonnesPerCubicMeter creates a new Density instance from a value in TonnesPerCubicMeter.
func (uf DensityFactory) FromTonnesPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicMeter)
}

// FromSlugsPerCubicFoot creates a new Density instance from a value in SlugsPerCubicFoot.
func (uf DensityFactory) FromSlugsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicFoot)
}

// FromGramsPerLiter creates a new Density instance from a value in GramsPerLiter.
func (uf DensityFactory) FromGramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerLiter)
}

// FromGramsPerDeciLiter creates a new Density instance from a value in GramsPerDeciLiter.
func (uf DensityFactory) FromGramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerDeciliter)
}

// FromGramsPerMilliliter creates a new Density instance from a value in GramsPerMilliliter.
func (uf DensityFactory) FromGramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerMilliliter)
}

// FromPoundsPerUSGallon creates a new Density instance from a value in PoundsPerUSGallon.
func (uf DensityFactory) FromPoundsPerUSGallon(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerUSGallon)
}

// FromPoundsPerImperialGallon creates a new Density instance from a value in PoundsPerImperialGallon.
func (uf DensityFactory) FromPoundsPerImperialGallon(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerImperialGallon)
}

// FromKilogramsPerLiter creates a new Density instance from a value in KilogramsPerLiter.
func (uf DensityFactory) FromKilogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerLiter)
}

// FromTonnesPerCubicFoot creates a new Density instance from a value in TonnesPerCubicFoot.
func (uf DensityFactory) FromTonnesPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicFoot)
}

// FromTonnesPerCubicInch creates a new Density instance from a value in TonnesPerCubicInch.
func (uf DensityFactory) FromTonnesPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicInch)
}

// FromGramsPerCubicFoot creates a new Density instance from a value in GramsPerCubicFoot.
func (uf DensityFactory) FromGramsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicFoot)
}

// FromGramsPerCubicInch creates a new Density instance from a value in GramsPerCubicInch.
func (uf DensityFactory) FromGramsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicInch)
}

// FromPoundsPerCubicMeter creates a new Density instance from a value in PoundsPerCubicMeter.
func (uf DensityFactory) FromPoundsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicMeter)
}

// FromPoundsPerCubicCentimeter creates a new Density instance from a value in PoundsPerCubicCentimeter.
func (uf DensityFactory) FromPoundsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicCentimeter)
}

// FromPoundsPerCubicMillimeter creates a new Density instance from a value in PoundsPerCubicMillimeter.
func (uf DensityFactory) FromPoundsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicMillimeter)
}

// FromSlugsPerCubicMeter creates a new Density instance from a value in SlugsPerCubicMeter.
func (uf DensityFactory) FromSlugsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicMeter)
}

// FromSlugsPerCubicCentimeter creates a new Density instance from a value in SlugsPerCubicCentimeter.
func (uf DensityFactory) FromSlugsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicCentimeter)
}

// FromSlugsPerCubicMillimeter creates a new Density instance from a value in SlugsPerCubicMillimeter.
func (uf DensityFactory) FromSlugsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicMillimeter)
}

// FromSlugsPerCubicInch creates a new Density instance from a value in SlugsPerCubicInch.
func (uf DensityFactory) FromSlugsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicInch)
}

// FromKilogramsPerCubicMillimeter creates a new Density instance from a value in KilogramsPerCubicMillimeter.
func (uf DensityFactory) FromKilogramsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicMillimeter)
}

// FromKilogramsPerCubicCentimeter creates a new Density instance from a value in KilogramsPerCubicCentimeter.
func (uf DensityFactory) FromKilogramsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicCentimeter)
}

// FromKilogramsPerCubicMeter creates a new Density instance from a value in KilogramsPerCubicMeter.
func (uf DensityFactory) FromKilogramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicMeter)
}

// FromMilligramsPerCubicMeter creates a new Density instance from a value in MilligramsPerCubicMeter.
func (uf DensityFactory) FromMilligramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerCubicMeter)
}

// FromMicrogramsPerCubicMeter creates a new Density instance from a value in MicrogramsPerCubicMeter.
func (uf DensityFactory) FromMicrogramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerCubicMeter)
}

// FromKilopoundsPerCubicInch creates a new Density instance from a value in KilopoundsPerCubicInch.
func (uf DensityFactory) FromKilopoundsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicInch)
}

// FromKilopoundsPerCubicFoot creates a new Density instance from a value in KilopoundsPerCubicFoot.
func (uf DensityFactory) FromKilopoundsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicFoot)
}

// FromKilopoundsPerCubicYard creates a new Density instance from a value in KilopoundsPerCubicYard.
func (uf DensityFactory) FromKilopoundsPerCubicYard(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicYard)
}

// FromFemtogramsPerLiter creates a new Density instance from a value in FemtogramsPerLiter.
func (uf DensityFactory) FromFemtogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerLiter)
}

// FromPicogramsPerLiter creates a new Density instance from a value in PicogramsPerLiter.
func (uf DensityFactory) FromPicogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerLiter)
}

// FromNanogramsPerLiter creates a new Density instance from a value in NanogramsPerLiter.
func (uf DensityFactory) FromNanogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerLiter)
}

// FromMicrogramsPerLiter creates a new Density instance from a value in MicrogramsPerLiter.
func (uf DensityFactory) FromMicrogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerLiter)
}

// FromMilligramsPerLiter creates a new Density instance from a value in MilligramsPerLiter.
func (uf DensityFactory) FromMilligramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerLiter)
}

// FromCentigramsPerLiter creates a new Density instance from a value in CentigramsPerLiter.
func (uf DensityFactory) FromCentigramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerLiter)
}

// FromDecigramsPerLiter creates a new Density instance from a value in DecigramsPerLiter.
func (uf DensityFactory) FromDecigramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerLiter)
}

// FromFemtogramsPerDeciLiter creates a new Density instance from a value in FemtogramsPerDeciLiter.
func (uf DensityFactory) FromFemtogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerDeciliter)
}

// FromPicogramsPerDeciLiter creates a new Density instance from a value in PicogramsPerDeciLiter.
func (uf DensityFactory) FromPicogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerDeciliter)
}

// FromNanogramsPerDeciLiter creates a new Density instance from a value in NanogramsPerDeciLiter.
func (uf DensityFactory) FromNanogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerDeciliter)
}

// FromMicrogramsPerDeciLiter creates a new Density instance from a value in MicrogramsPerDeciLiter.
func (uf DensityFactory) FromMicrogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerDeciliter)
}

// FromMilligramsPerDeciLiter creates a new Density instance from a value in MilligramsPerDeciLiter.
func (uf DensityFactory) FromMilligramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerDeciliter)
}

// FromCentigramsPerDeciLiter creates a new Density instance from a value in CentigramsPerDeciLiter.
func (uf DensityFactory) FromCentigramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerDeciliter)
}

// FromDecigramsPerDeciLiter creates a new Density instance from a value in DecigramsPerDeciLiter.
func (uf DensityFactory) FromDecigramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerDeciliter)
}

// FromFemtogramsPerMilliliter creates a new Density instance from a value in FemtogramsPerMilliliter.
func (uf DensityFactory) FromFemtogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerMilliliter)
}

// FromPicogramsPerMilliliter creates a new Density instance from a value in PicogramsPerMilliliter.
func (uf DensityFactory) FromPicogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerMilliliter)
}

// FromNanogramsPerMilliliter creates a new Density instance from a value in NanogramsPerMilliliter.
func (uf DensityFactory) FromNanogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerMilliliter)
}

// FromMicrogramsPerMilliliter creates a new Density instance from a value in MicrogramsPerMilliliter.
func (uf DensityFactory) FromMicrogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerMilliliter)
}

// FromMilligramsPerMilliliter creates a new Density instance from a value in MilligramsPerMilliliter.
func (uf DensityFactory) FromMilligramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerMilliliter)
}

// FromCentigramsPerMilliliter creates a new Density instance from a value in CentigramsPerMilliliter.
func (uf DensityFactory) FromCentigramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerMilliliter)
}

// FromDecigramsPerMilliliter creates a new Density instance from a value in DecigramsPerMilliliter.
func (uf DensityFactory) FromDecigramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerMilliliter)
}


// newDensity creates a new Density.
func newDensity(value float64, fromUnit DensityUnits) (*Density, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalDensityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in DensityUnits", fromUnit)
	}
	a := &Density{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Density in KilogramPerCubicMeter unit (the base/default quantity).
func (a *Density) BaseValue() float64 {
	return a.value
}


// GramsPerCubicMillimeter returns the Density value in GramsPerCubicMillimeter.
//
// 
func (a *Density) GramsPerCubicMillimeter() float64 {
	if a.grams_per_cubic_millimeterLazy != nil {
		return *a.grams_per_cubic_millimeterLazy
	}
	grams_per_cubic_millimeter := a.convertFromBase(DensityGramPerCubicMillimeter)
	a.grams_per_cubic_millimeterLazy = &grams_per_cubic_millimeter
	return grams_per_cubic_millimeter
}

// GramsPerCubicCentimeter returns the Density value in GramsPerCubicCentimeter.
//
// 
func (a *Density) GramsPerCubicCentimeter() float64 {
	if a.grams_per_cubic_centimeterLazy != nil {
		return *a.grams_per_cubic_centimeterLazy
	}
	grams_per_cubic_centimeter := a.convertFromBase(DensityGramPerCubicCentimeter)
	a.grams_per_cubic_centimeterLazy = &grams_per_cubic_centimeter
	return grams_per_cubic_centimeter
}

// GramsPerCubicMeter returns the Density value in GramsPerCubicMeter.
//
// 
func (a *Density) GramsPerCubicMeter() float64 {
	if a.grams_per_cubic_meterLazy != nil {
		return *a.grams_per_cubic_meterLazy
	}
	grams_per_cubic_meter := a.convertFromBase(DensityGramPerCubicMeter)
	a.grams_per_cubic_meterLazy = &grams_per_cubic_meter
	return grams_per_cubic_meter
}

// PoundsPerCubicInch returns the Density value in PoundsPerCubicInch.
//
// 
func (a *Density) PoundsPerCubicInch() float64 {
	if a.pounds_per_cubic_inchLazy != nil {
		return *a.pounds_per_cubic_inchLazy
	}
	pounds_per_cubic_inch := a.convertFromBase(DensityPoundPerCubicInch)
	a.pounds_per_cubic_inchLazy = &pounds_per_cubic_inch
	return pounds_per_cubic_inch
}

// PoundsPerCubicFoot returns the Density value in PoundsPerCubicFoot.
//
// 
func (a *Density) PoundsPerCubicFoot() float64 {
	if a.pounds_per_cubic_footLazy != nil {
		return *a.pounds_per_cubic_footLazy
	}
	pounds_per_cubic_foot := a.convertFromBase(DensityPoundPerCubicFoot)
	a.pounds_per_cubic_footLazy = &pounds_per_cubic_foot
	return pounds_per_cubic_foot
}

// PoundsPerCubicYard returns the Density value in PoundsPerCubicYard.
//
// Calculated from the definition of <a href="https://en.wikipedia.org/wiki/Pound_(mass)">pound</a> and <a href="https://en.wikipedia.org/wiki/Yard">yard</a> compared to metric kilogram and meter.
func (a *Density) PoundsPerCubicYard() float64 {
	if a.pounds_per_cubic_yardLazy != nil {
		return *a.pounds_per_cubic_yardLazy
	}
	pounds_per_cubic_yard := a.convertFromBase(DensityPoundPerCubicYard)
	a.pounds_per_cubic_yardLazy = &pounds_per_cubic_yard
	return pounds_per_cubic_yard
}

// TonnesPerCubicMillimeter returns the Density value in TonnesPerCubicMillimeter.
//
// 
func (a *Density) TonnesPerCubicMillimeter() float64 {
	if a.tonnes_per_cubic_millimeterLazy != nil {
		return *a.tonnes_per_cubic_millimeterLazy
	}
	tonnes_per_cubic_millimeter := a.convertFromBase(DensityTonnePerCubicMillimeter)
	a.tonnes_per_cubic_millimeterLazy = &tonnes_per_cubic_millimeter
	return tonnes_per_cubic_millimeter
}

// TonnesPerCubicCentimeter returns the Density value in TonnesPerCubicCentimeter.
//
// 
func (a *Density) TonnesPerCubicCentimeter() float64 {
	if a.tonnes_per_cubic_centimeterLazy != nil {
		return *a.tonnes_per_cubic_centimeterLazy
	}
	tonnes_per_cubic_centimeter := a.convertFromBase(DensityTonnePerCubicCentimeter)
	a.tonnes_per_cubic_centimeterLazy = &tonnes_per_cubic_centimeter
	return tonnes_per_cubic_centimeter
}

// TonnesPerCubicMeter returns the Density value in TonnesPerCubicMeter.
//
// 
func (a *Density) TonnesPerCubicMeter() float64 {
	if a.tonnes_per_cubic_meterLazy != nil {
		return *a.tonnes_per_cubic_meterLazy
	}
	tonnes_per_cubic_meter := a.convertFromBase(DensityTonnePerCubicMeter)
	a.tonnes_per_cubic_meterLazy = &tonnes_per_cubic_meter
	return tonnes_per_cubic_meter
}

// SlugsPerCubicFoot returns the Density value in SlugsPerCubicFoot.
//
// 
func (a *Density) SlugsPerCubicFoot() float64 {
	if a.slugs_per_cubic_footLazy != nil {
		return *a.slugs_per_cubic_footLazy
	}
	slugs_per_cubic_foot := a.convertFromBase(DensitySlugPerCubicFoot)
	a.slugs_per_cubic_footLazy = &slugs_per_cubic_foot
	return slugs_per_cubic_foot
}

// GramsPerLiter returns the Density value in GramsPerLiter.
//
// 
func (a *Density) GramsPerLiter() float64 {
	if a.grams_per_literLazy != nil {
		return *a.grams_per_literLazy
	}
	grams_per_liter := a.convertFromBase(DensityGramPerLiter)
	a.grams_per_literLazy = &grams_per_liter
	return grams_per_liter
}

// GramsPerDeciLiter returns the Density value in GramsPerDeciLiter.
//
// 
func (a *Density) GramsPerDeciLiter() float64 {
	if a.grams_per_deci_literLazy != nil {
		return *a.grams_per_deci_literLazy
	}
	grams_per_deci_liter := a.convertFromBase(DensityGramPerDeciliter)
	a.grams_per_deci_literLazy = &grams_per_deci_liter
	return grams_per_deci_liter
}

// GramsPerMilliliter returns the Density value in GramsPerMilliliter.
//
// 
func (a *Density) GramsPerMilliliter() float64 {
	if a.grams_per_milliliterLazy != nil {
		return *a.grams_per_milliliterLazy
	}
	grams_per_milliliter := a.convertFromBase(DensityGramPerMilliliter)
	a.grams_per_milliliterLazy = &grams_per_milliliter
	return grams_per_milliliter
}

// PoundsPerUSGallon returns the Density value in PoundsPerUSGallon.
//
// 
func (a *Density) PoundsPerUSGallon() float64 {
	if a.pounds_per_us_gallonLazy != nil {
		return *a.pounds_per_us_gallonLazy
	}
	pounds_per_us_gallon := a.convertFromBase(DensityPoundPerUSGallon)
	a.pounds_per_us_gallonLazy = &pounds_per_us_gallon
	return pounds_per_us_gallon
}

// PoundsPerImperialGallon returns the Density value in PoundsPerImperialGallon.
//
// 
func (a *Density) PoundsPerImperialGallon() float64 {
	if a.pounds_per_imperial_gallonLazy != nil {
		return *a.pounds_per_imperial_gallonLazy
	}
	pounds_per_imperial_gallon := a.convertFromBase(DensityPoundPerImperialGallon)
	a.pounds_per_imperial_gallonLazy = &pounds_per_imperial_gallon
	return pounds_per_imperial_gallon
}

// KilogramsPerLiter returns the Density value in KilogramsPerLiter.
//
// 
func (a *Density) KilogramsPerLiter() float64 {
	if a.kilograms_per_literLazy != nil {
		return *a.kilograms_per_literLazy
	}
	kilograms_per_liter := a.convertFromBase(DensityKilogramPerLiter)
	a.kilograms_per_literLazy = &kilograms_per_liter
	return kilograms_per_liter
}

// TonnesPerCubicFoot returns the Density value in TonnesPerCubicFoot.
//
// 
func (a *Density) TonnesPerCubicFoot() float64 {
	if a.tonnes_per_cubic_footLazy != nil {
		return *a.tonnes_per_cubic_footLazy
	}
	tonnes_per_cubic_foot := a.convertFromBase(DensityTonnePerCubicFoot)
	a.tonnes_per_cubic_footLazy = &tonnes_per_cubic_foot
	return tonnes_per_cubic_foot
}

// TonnesPerCubicInch returns the Density value in TonnesPerCubicInch.
//
// 
func (a *Density) TonnesPerCubicInch() float64 {
	if a.tonnes_per_cubic_inchLazy != nil {
		return *a.tonnes_per_cubic_inchLazy
	}
	tonnes_per_cubic_inch := a.convertFromBase(DensityTonnePerCubicInch)
	a.tonnes_per_cubic_inchLazy = &tonnes_per_cubic_inch
	return tonnes_per_cubic_inch
}

// GramsPerCubicFoot returns the Density value in GramsPerCubicFoot.
//
// 
func (a *Density) GramsPerCubicFoot() float64 {
	if a.grams_per_cubic_footLazy != nil {
		return *a.grams_per_cubic_footLazy
	}
	grams_per_cubic_foot := a.convertFromBase(DensityGramPerCubicFoot)
	a.grams_per_cubic_footLazy = &grams_per_cubic_foot
	return grams_per_cubic_foot
}

// GramsPerCubicInch returns the Density value in GramsPerCubicInch.
//
// 
func (a *Density) GramsPerCubicInch() float64 {
	if a.grams_per_cubic_inchLazy != nil {
		return *a.grams_per_cubic_inchLazy
	}
	grams_per_cubic_inch := a.convertFromBase(DensityGramPerCubicInch)
	a.grams_per_cubic_inchLazy = &grams_per_cubic_inch
	return grams_per_cubic_inch
}

// PoundsPerCubicMeter returns the Density value in PoundsPerCubicMeter.
//
// 
func (a *Density) PoundsPerCubicMeter() float64 {
	if a.pounds_per_cubic_meterLazy != nil {
		return *a.pounds_per_cubic_meterLazy
	}
	pounds_per_cubic_meter := a.convertFromBase(DensityPoundPerCubicMeter)
	a.pounds_per_cubic_meterLazy = &pounds_per_cubic_meter
	return pounds_per_cubic_meter
}

// PoundsPerCubicCentimeter returns the Density value in PoundsPerCubicCentimeter.
//
// 
func (a *Density) PoundsPerCubicCentimeter() float64 {
	if a.pounds_per_cubic_centimeterLazy != nil {
		return *a.pounds_per_cubic_centimeterLazy
	}
	pounds_per_cubic_centimeter := a.convertFromBase(DensityPoundPerCubicCentimeter)
	a.pounds_per_cubic_centimeterLazy = &pounds_per_cubic_centimeter
	return pounds_per_cubic_centimeter
}

// PoundsPerCubicMillimeter returns the Density value in PoundsPerCubicMillimeter.
//
// 
func (a *Density) PoundsPerCubicMillimeter() float64 {
	if a.pounds_per_cubic_millimeterLazy != nil {
		return *a.pounds_per_cubic_millimeterLazy
	}
	pounds_per_cubic_millimeter := a.convertFromBase(DensityPoundPerCubicMillimeter)
	a.pounds_per_cubic_millimeterLazy = &pounds_per_cubic_millimeter
	return pounds_per_cubic_millimeter
}

// SlugsPerCubicMeter returns the Density value in SlugsPerCubicMeter.
//
// 
func (a *Density) SlugsPerCubicMeter() float64 {
	if a.slugs_per_cubic_meterLazy != nil {
		return *a.slugs_per_cubic_meterLazy
	}
	slugs_per_cubic_meter := a.convertFromBase(DensitySlugPerCubicMeter)
	a.slugs_per_cubic_meterLazy = &slugs_per_cubic_meter
	return slugs_per_cubic_meter
}

// SlugsPerCubicCentimeter returns the Density value in SlugsPerCubicCentimeter.
//
// 
func (a *Density) SlugsPerCubicCentimeter() float64 {
	if a.slugs_per_cubic_centimeterLazy != nil {
		return *a.slugs_per_cubic_centimeterLazy
	}
	slugs_per_cubic_centimeter := a.convertFromBase(DensitySlugPerCubicCentimeter)
	a.slugs_per_cubic_centimeterLazy = &slugs_per_cubic_centimeter
	return slugs_per_cubic_centimeter
}

// SlugsPerCubicMillimeter returns the Density value in SlugsPerCubicMillimeter.
//
// 
func (a *Density) SlugsPerCubicMillimeter() float64 {
	if a.slugs_per_cubic_millimeterLazy != nil {
		return *a.slugs_per_cubic_millimeterLazy
	}
	slugs_per_cubic_millimeter := a.convertFromBase(DensitySlugPerCubicMillimeter)
	a.slugs_per_cubic_millimeterLazy = &slugs_per_cubic_millimeter
	return slugs_per_cubic_millimeter
}

// SlugsPerCubicInch returns the Density value in SlugsPerCubicInch.
//
// 
func (a *Density) SlugsPerCubicInch() float64 {
	if a.slugs_per_cubic_inchLazy != nil {
		return *a.slugs_per_cubic_inchLazy
	}
	slugs_per_cubic_inch := a.convertFromBase(DensitySlugPerCubicInch)
	a.slugs_per_cubic_inchLazy = &slugs_per_cubic_inch
	return slugs_per_cubic_inch
}

// KilogramsPerCubicMillimeter returns the Density value in KilogramsPerCubicMillimeter.
//
// 
func (a *Density) KilogramsPerCubicMillimeter() float64 {
	if a.kilograms_per_cubic_millimeterLazy != nil {
		return *a.kilograms_per_cubic_millimeterLazy
	}
	kilograms_per_cubic_millimeter := a.convertFromBase(DensityKilogramPerCubicMillimeter)
	a.kilograms_per_cubic_millimeterLazy = &kilograms_per_cubic_millimeter
	return kilograms_per_cubic_millimeter
}

// KilogramsPerCubicCentimeter returns the Density value in KilogramsPerCubicCentimeter.
//
// 
func (a *Density) KilogramsPerCubicCentimeter() float64 {
	if a.kilograms_per_cubic_centimeterLazy != nil {
		return *a.kilograms_per_cubic_centimeterLazy
	}
	kilograms_per_cubic_centimeter := a.convertFromBase(DensityKilogramPerCubicCentimeter)
	a.kilograms_per_cubic_centimeterLazy = &kilograms_per_cubic_centimeter
	return kilograms_per_cubic_centimeter
}

// KilogramsPerCubicMeter returns the Density value in KilogramsPerCubicMeter.
//
// 
func (a *Density) KilogramsPerCubicMeter() float64 {
	if a.kilograms_per_cubic_meterLazy != nil {
		return *a.kilograms_per_cubic_meterLazy
	}
	kilograms_per_cubic_meter := a.convertFromBase(DensityKilogramPerCubicMeter)
	a.kilograms_per_cubic_meterLazy = &kilograms_per_cubic_meter
	return kilograms_per_cubic_meter
}

// MilligramsPerCubicMeter returns the Density value in MilligramsPerCubicMeter.
//
// 
func (a *Density) MilligramsPerCubicMeter() float64 {
	if a.milligrams_per_cubic_meterLazy != nil {
		return *a.milligrams_per_cubic_meterLazy
	}
	milligrams_per_cubic_meter := a.convertFromBase(DensityMilligramPerCubicMeter)
	a.milligrams_per_cubic_meterLazy = &milligrams_per_cubic_meter
	return milligrams_per_cubic_meter
}

// MicrogramsPerCubicMeter returns the Density value in MicrogramsPerCubicMeter.
//
// 
func (a *Density) MicrogramsPerCubicMeter() float64 {
	if a.micrograms_per_cubic_meterLazy != nil {
		return *a.micrograms_per_cubic_meterLazy
	}
	micrograms_per_cubic_meter := a.convertFromBase(DensityMicrogramPerCubicMeter)
	a.micrograms_per_cubic_meterLazy = &micrograms_per_cubic_meter
	return micrograms_per_cubic_meter
}

// KilopoundsPerCubicInch returns the Density value in KilopoundsPerCubicInch.
//
// 
func (a *Density) KilopoundsPerCubicInch() float64 {
	if a.kilopounds_per_cubic_inchLazy != nil {
		return *a.kilopounds_per_cubic_inchLazy
	}
	kilopounds_per_cubic_inch := a.convertFromBase(DensityKilopoundPerCubicInch)
	a.kilopounds_per_cubic_inchLazy = &kilopounds_per_cubic_inch
	return kilopounds_per_cubic_inch
}

// KilopoundsPerCubicFoot returns the Density value in KilopoundsPerCubicFoot.
//
// 
func (a *Density) KilopoundsPerCubicFoot() float64 {
	if a.kilopounds_per_cubic_footLazy != nil {
		return *a.kilopounds_per_cubic_footLazy
	}
	kilopounds_per_cubic_foot := a.convertFromBase(DensityKilopoundPerCubicFoot)
	a.kilopounds_per_cubic_footLazy = &kilopounds_per_cubic_foot
	return kilopounds_per_cubic_foot
}

// KilopoundsPerCubicYard returns the Density value in KilopoundsPerCubicYard.
//
// 
func (a *Density) KilopoundsPerCubicYard() float64 {
	if a.kilopounds_per_cubic_yardLazy != nil {
		return *a.kilopounds_per_cubic_yardLazy
	}
	kilopounds_per_cubic_yard := a.convertFromBase(DensityKilopoundPerCubicYard)
	a.kilopounds_per_cubic_yardLazy = &kilopounds_per_cubic_yard
	return kilopounds_per_cubic_yard
}

// FemtogramsPerLiter returns the Density value in FemtogramsPerLiter.
//
// 
func (a *Density) FemtogramsPerLiter() float64 {
	if a.femtograms_per_literLazy != nil {
		return *a.femtograms_per_literLazy
	}
	femtograms_per_liter := a.convertFromBase(DensityFemtogramPerLiter)
	a.femtograms_per_literLazy = &femtograms_per_liter
	return femtograms_per_liter
}

// PicogramsPerLiter returns the Density value in PicogramsPerLiter.
//
// 
func (a *Density) PicogramsPerLiter() float64 {
	if a.picograms_per_literLazy != nil {
		return *a.picograms_per_literLazy
	}
	picograms_per_liter := a.convertFromBase(DensityPicogramPerLiter)
	a.picograms_per_literLazy = &picograms_per_liter
	return picograms_per_liter
}

// NanogramsPerLiter returns the Density value in NanogramsPerLiter.
//
// 
func (a *Density) NanogramsPerLiter() float64 {
	if a.nanograms_per_literLazy != nil {
		return *a.nanograms_per_literLazy
	}
	nanograms_per_liter := a.convertFromBase(DensityNanogramPerLiter)
	a.nanograms_per_literLazy = &nanograms_per_liter
	return nanograms_per_liter
}

// MicrogramsPerLiter returns the Density value in MicrogramsPerLiter.
//
// 
func (a *Density) MicrogramsPerLiter() float64 {
	if a.micrograms_per_literLazy != nil {
		return *a.micrograms_per_literLazy
	}
	micrograms_per_liter := a.convertFromBase(DensityMicrogramPerLiter)
	a.micrograms_per_literLazy = &micrograms_per_liter
	return micrograms_per_liter
}

// MilligramsPerLiter returns the Density value in MilligramsPerLiter.
//
// 
func (a *Density) MilligramsPerLiter() float64 {
	if a.milligrams_per_literLazy != nil {
		return *a.milligrams_per_literLazy
	}
	milligrams_per_liter := a.convertFromBase(DensityMilligramPerLiter)
	a.milligrams_per_literLazy = &milligrams_per_liter
	return milligrams_per_liter
}

// CentigramsPerLiter returns the Density value in CentigramsPerLiter.
//
// 
func (a *Density) CentigramsPerLiter() float64 {
	if a.centigrams_per_literLazy != nil {
		return *a.centigrams_per_literLazy
	}
	centigrams_per_liter := a.convertFromBase(DensityCentigramPerLiter)
	a.centigrams_per_literLazy = &centigrams_per_liter
	return centigrams_per_liter
}

// DecigramsPerLiter returns the Density value in DecigramsPerLiter.
//
// 
func (a *Density) DecigramsPerLiter() float64 {
	if a.decigrams_per_literLazy != nil {
		return *a.decigrams_per_literLazy
	}
	decigrams_per_liter := a.convertFromBase(DensityDecigramPerLiter)
	a.decigrams_per_literLazy = &decigrams_per_liter
	return decigrams_per_liter
}

// FemtogramsPerDeciLiter returns the Density value in FemtogramsPerDeciLiter.
//
// 
func (a *Density) FemtogramsPerDeciLiter() float64 {
	if a.femtograms_per_deci_literLazy != nil {
		return *a.femtograms_per_deci_literLazy
	}
	femtograms_per_deci_liter := a.convertFromBase(DensityFemtogramPerDeciliter)
	a.femtograms_per_deci_literLazy = &femtograms_per_deci_liter
	return femtograms_per_deci_liter
}

// PicogramsPerDeciLiter returns the Density value in PicogramsPerDeciLiter.
//
// 
func (a *Density) PicogramsPerDeciLiter() float64 {
	if a.picograms_per_deci_literLazy != nil {
		return *a.picograms_per_deci_literLazy
	}
	picograms_per_deci_liter := a.convertFromBase(DensityPicogramPerDeciliter)
	a.picograms_per_deci_literLazy = &picograms_per_deci_liter
	return picograms_per_deci_liter
}

// NanogramsPerDeciLiter returns the Density value in NanogramsPerDeciLiter.
//
// 
func (a *Density) NanogramsPerDeciLiter() float64 {
	if a.nanograms_per_deci_literLazy != nil {
		return *a.nanograms_per_deci_literLazy
	}
	nanograms_per_deci_liter := a.convertFromBase(DensityNanogramPerDeciliter)
	a.nanograms_per_deci_literLazy = &nanograms_per_deci_liter
	return nanograms_per_deci_liter
}

// MicrogramsPerDeciLiter returns the Density value in MicrogramsPerDeciLiter.
//
// 
func (a *Density) MicrogramsPerDeciLiter() float64 {
	if a.micrograms_per_deci_literLazy != nil {
		return *a.micrograms_per_deci_literLazy
	}
	micrograms_per_deci_liter := a.convertFromBase(DensityMicrogramPerDeciliter)
	a.micrograms_per_deci_literLazy = &micrograms_per_deci_liter
	return micrograms_per_deci_liter
}

// MilligramsPerDeciLiter returns the Density value in MilligramsPerDeciLiter.
//
// 
func (a *Density) MilligramsPerDeciLiter() float64 {
	if a.milligrams_per_deci_literLazy != nil {
		return *a.milligrams_per_deci_literLazy
	}
	milligrams_per_deci_liter := a.convertFromBase(DensityMilligramPerDeciliter)
	a.milligrams_per_deci_literLazy = &milligrams_per_deci_liter
	return milligrams_per_deci_liter
}

// CentigramsPerDeciLiter returns the Density value in CentigramsPerDeciLiter.
//
// 
func (a *Density) CentigramsPerDeciLiter() float64 {
	if a.centigrams_per_deci_literLazy != nil {
		return *a.centigrams_per_deci_literLazy
	}
	centigrams_per_deci_liter := a.convertFromBase(DensityCentigramPerDeciliter)
	a.centigrams_per_deci_literLazy = &centigrams_per_deci_liter
	return centigrams_per_deci_liter
}

// DecigramsPerDeciLiter returns the Density value in DecigramsPerDeciLiter.
//
// 
func (a *Density) DecigramsPerDeciLiter() float64 {
	if a.decigrams_per_deci_literLazy != nil {
		return *a.decigrams_per_deci_literLazy
	}
	decigrams_per_deci_liter := a.convertFromBase(DensityDecigramPerDeciliter)
	a.decigrams_per_deci_literLazy = &decigrams_per_deci_liter
	return decigrams_per_deci_liter
}

// FemtogramsPerMilliliter returns the Density value in FemtogramsPerMilliliter.
//
// 
func (a *Density) FemtogramsPerMilliliter() float64 {
	if a.femtograms_per_milliliterLazy != nil {
		return *a.femtograms_per_milliliterLazy
	}
	femtograms_per_milliliter := a.convertFromBase(DensityFemtogramPerMilliliter)
	a.femtograms_per_milliliterLazy = &femtograms_per_milliliter
	return femtograms_per_milliliter
}

// PicogramsPerMilliliter returns the Density value in PicogramsPerMilliliter.
//
// 
func (a *Density) PicogramsPerMilliliter() float64 {
	if a.picograms_per_milliliterLazy != nil {
		return *a.picograms_per_milliliterLazy
	}
	picograms_per_milliliter := a.convertFromBase(DensityPicogramPerMilliliter)
	a.picograms_per_milliliterLazy = &picograms_per_milliliter
	return picograms_per_milliliter
}

// NanogramsPerMilliliter returns the Density value in NanogramsPerMilliliter.
//
// 
func (a *Density) NanogramsPerMilliliter() float64 {
	if a.nanograms_per_milliliterLazy != nil {
		return *a.nanograms_per_milliliterLazy
	}
	nanograms_per_milliliter := a.convertFromBase(DensityNanogramPerMilliliter)
	a.nanograms_per_milliliterLazy = &nanograms_per_milliliter
	return nanograms_per_milliliter
}

// MicrogramsPerMilliliter returns the Density value in MicrogramsPerMilliliter.
//
// 
func (a *Density) MicrogramsPerMilliliter() float64 {
	if a.micrograms_per_milliliterLazy != nil {
		return *a.micrograms_per_milliliterLazy
	}
	micrograms_per_milliliter := a.convertFromBase(DensityMicrogramPerMilliliter)
	a.micrograms_per_milliliterLazy = &micrograms_per_milliliter
	return micrograms_per_milliliter
}

// MilligramsPerMilliliter returns the Density value in MilligramsPerMilliliter.
//
// 
func (a *Density) MilligramsPerMilliliter() float64 {
	if a.milligrams_per_milliliterLazy != nil {
		return *a.milligrams_per_milliliterLazy
	}
	milligrams_per_milliliter := a.convertFromBase(DensityMilligramPerMilliliter)
	a.milligrams_per_milliliterLazy = &milligrams_per_milliliter
	return milligrams_per_milliliter
}

// CentigramsPerMilliliter returns the Density value in CentigramsPerMilliliter.
//
// 
func (a *Density) CentigramsPerMilliliter() float64 {
	if a.centigrams_per_milliliterLazy != nil {
		return *a.centigrams_per_milliliterLazy
	}
	centigrams_per_milliliter := a.convertFromBase(DensityCentigramPerMilliliter)
	a.centigrams_per_milliliterLazy = &centigrams_per_milliliter
	return centigrams_per_milliliter
}

// DecigramsPerMilliliter returns the Density value in DecigramsPerMilliliter.
//
// 
func (a *Density) DecigramsPerMilliliter() float64 {
	if a.decigrams_per_milliliterLazy != nil {
		return *a.decigrams_per_milliliterLazy
	}
	decigrams_per_milliliter := a.convertFromBase(DensityDecigramPerMilliliter)
	a.decigrams_per_milliliterLazy = &decigrams_per_milliliter
	return decigrams_per_milliliter
}


// ToDto creates a DensityDto representation from the Density instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerCubicMeter by default.
func (a *Density) ToDto(holdInUnit *DensityUnits) DensityDto {
	if holdInUnit == nil {
		defaultUnit := DensityKilogramPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return DensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Density instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerCubicMeter by default.
func (a *Density) ToDtoJSON(holdInUnit *DensityUnits) ([]byte, error) {
	// Convert to DensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Density to a specific unit value.
// The function uses the provided unit type (DensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Density) Convert(toUnit DensityUnits) float64 {
	switch toUnit { 
    case DensityGramPerCubicMillimeter:
		return a.GramsPerCubicMillimeter()
    case DensityGramPerCubicCentimeter:
		return a.GramsPerCubicCentimeter()
    case DensityGramPerCubicMeter:
		return a.GramsPerCubicMeter()
    case DensityPoundPerCubicInch:
		return a.PoundsPerCubicInch()
    case DensityPoundPerCubicFoot:
		return a.PoundsPerCubicFoot()
    case DensityPoundPerCubicYard:
		return a.PoundsPerCubicYard()
    case DensityTonnePerCubicMillimeter:
		return a.TonnesPerCubicMillimeter()
    case DensityTonnePerCubicCentimeter:
		return a.TonnesPerCubicCentimeter()
    case DensityTonnePerCubicMeter:
		return a.TonnesPerCubicMeter()
    case DensitySlugPerCubicFoot:
		return a.SlugsPerCubicFoot()
    case DensityGramPerLiter:
		return a.GramsPerLiter()
    case DensityGramPerDeciliter:
		return a.GramsPerDeciLiter()
    case DensityGramPerMilliliter:
		return a.GramsPerMilliliter()
    case DensityPoundPerUSGallon:
		return a.PoundsPerUSGallon()
    case DensityPoundPerImperialGallon:
		return a.PoundsPerImperialGallon()
    case DensityKilogramPerLiter:
		return a.KilogramsPerLiter()
    case DensityTonnePerCubicFoot:
		return a.TonnesPerCubicFoot()
    case DensityTonnePerCubicInch:
		return a.TonnesPerCubicInch()
    case DensityGramPerCubicFoot:
		return a.GramsPerCubicFoot()
    case DensityGramPerCubicInch:
		return a.GramsPerCubicInch()
    case DensityPoundPerCubicMeter:
		return a.PoundsPerCubicMeter()
    case DensityPoundPerCubicCentimeter:
		return a.PoundsPerCubicCentimeter()
    case DensityPoundPerCubicMillimeter:
		return a.PoundsPerCubicMillimeter()
    case DensitySlugPerCubicMeter:
		return a.SlugsPerCubicMeter()
    case DensitySlugPerCubicCentimeter:
		return a.SlugsPerCubicCentimeter()
    case DensitySlugPerCubicMillimeter:
		return a.SlugsPerCubicMillimeter()
    case DensitySlugPerCubicInch:
		return a.SlugsPerCubicInch()
    case DensityKilogramPerCubicMillimeter:
		return a.KilogramsPerCubicMillimeter()
    case DensityKilogramPerCubicCentimeter:
		return a.KilogramsPerCubicCentimeter()
    case DensityKilogramPerCubicMeter:
		return a.KilogramsPerCubicMeter()
    case DensityMilligramPerCubicMeter:
		return a.MilligramsPerCubicMeter()
    case DensityMicrogramPerCubicMeter:
		return a.MicrogramsPerCubicMeter()
    case DensityKilopoundPerCubicInch:
		return a.KilopoundsPerCubicInch()
    case DensityKilopoundPerCubicFoot:
		return a.KilopoundsPerCubicFoot()
    case DensityKilopoundPerCubicYard:
		return a.KilopoundsPerCubicYard()
    case DensityFemtogramPerLiter:
		return a.FemtogramsPerLiter()
    case DensityPicogramPerLiter:
		return a.PicogramsPerLiter()
    case DensityNanogramPerLiter:
		return a.NanogramsPerLiter()
    case DensityMicrogramPerLiter:
		return a.MicrogramsPerLiter()
    case DensityMilligramPerLiter:
		return a.MilligramsPerLiter()
    case DensityCentigramPerLiter:
		return a.CentigramsPerLiter()
    case DensityDecigramPerLiter:
		return a.DecigramsPerLiter()
    case DensityFemtogramPerDeciliter:
		return a.FemtogramsPerDeciLiter()
    case DensityPicogramPerDeciliter:
		return a.PicogramsPerDeciLiter()
    case DensityNanogramPerDeciliter:
		return a.NanogramsPerDeciLiter()
    case DensityMicrogramPerDeciliter:
		return a.MicrogramsPerDeciLiter()
    case DensityMilligramPerDeciliter:
		return a.MilligramsPerDeciLiter()
    case DensityCentigramPerDeciliter:
		return a.CentigramsPerDeciLiter()
    case DensityDecigramPerDeciliter:
		return a.DecigramsPerDeciLiter()
    case DensityFemtogramPerMilliliter:
		return a.FemtogramsPerMilliliter()
    case DensityPicogramPerMilliliter:
		return a.PicogramsPerMilliliter()
    case DensityNanogramPerMilliliter:
		return a.NanogramsPerMilliliter()
    case DensityMicrogramPerMilliliter:
		return a.MicrogramsPerMilliliter()
    case DensityMilligramPerMilliliter:
		return a.MilligramsPerMilliliter()
    case DensityCentigramPerMilliliter:
		return a.CentigramsPerMilliliter()
    case DensityDecigramPerMilliliter:
		return a.DecigramsPerMilliliter()
	default:
		return math.NaN()
	}
}

func (a *Density) convertFromBase(toUnit DensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case DensityGramPerCubicMillimeter:
		return (value * 1e-6) 
	case DensityGramPerCubicCentimeter:
		return (value * 1e-3) 
	case DensityGramPerCubicMeter:
		return (value * 1e3) 
	case DensityPoundPerCubicInch:
		return (value * 3.6127298147753e-5) 
	case DensityPoundPerCubicFoot:
		return (value * 0.062427961) 
	case DensityPoundPerCubicYard:
		return (value / (0.45359237 / 0.9144 / 0.9144 / 0.9144)) 
	case DensityTonnePerCubicMillimeter:
		return (value * 1e-12) 
	case DensityTonnePerCubicCentimeter:
		return (value * 1e-9) 
	case DensityTonnePerCubicMeter:
		return (value * 0.001) 
	case DensitySlugPerCubicFoot:
		return (value * 0.00194032033) 
	case DensityGramPerLiter:
		return (value * 1) 
	case DensityGramPerDeciliter:
		return (value * 1e-1) 
	case DensityGramPerMilliliter:
		return (value * 1e-3) 
	case DensityPoundPerUSGallon:
		return (value / 1.19826427e2) 
	case DensityPoundPerImperialGallon:
		return (value / 9.9776398e1) 
	case DensityKilogramPerLiter:
		return (value / 1e3) 
	case DensityTonnePerCubicFoot:
		return (value / 3.53146667214886e4) 
	case DensityTonnePerCubicInch:
		return (value / 6.10237440947323e7) 
	case DensityGramPerCubicFoot:
		return (value / 0.0353146667214886) 
	case DensityGramPerCubicInch:
		return (value / 61.0237440947323) 
	case DensityPoundPerCubicMeter:
		return (value * 2.204622621848775) 
	case DensityPoundPerCubicCentimeter:
		return (value * 2.204622621848775e-6) 
	case DensityPoundPerCubicMillimeter:
		return (value * 2.204622621848775e-9) 
	case DensitySlugPerCubicMeter:
		return (value / 14.5939) 
	case DensitySlugPerCubicCentimeter:
		return (value / 14593903) 
	case DensitySlugPerCubicMillimeter:
		return (value / 14593903000) 
	case DensitySlugPerCubicInch:
		return (value / 890574.60201535) 
	case DensityKilogramPerCubicMillimeter:
		return ((value * 1e-6) / 1000.0) 
	case DensityKilogramPerCubicCentimeter:
		return ((value * 1e-3) / 1000.0) 
	case DensityKilogramPerCubicMeter:
		return ((value * 1e3) / 1000.0) 
	case DensityMilligramPerCubicMeter:
		return ((value * 1e3) / 0.001) 
	case DensityMicrogramPerCubicMeter:
		return ((value * 1e3) / 1e-06) 
	case DensityKilopoundPerCubicInch:
		return ((value * 3.6127298147753e-5) / 1000.0) 
	case DensityKilopoundPerCubicFoot:
		return ((value * 0.062427961) / 1000.0) 
	case DensityKilopoundPerCubicYard:
		return ((value / (0.45359237 / 0.9144 / 0.9144 / 0.9144)) / 1000.0) 
	case DensityFemtogramPerLiter:
		return ((value * 1) / 1e-15) 
	case DensityPicogramPerLiter:
		return ((value * 1) / 1e-12) 
	case DensityNanogramPerLiter:
		return ((value * 1) / 1e-09) 
	case DensityMicrogramPerLiter:
		return ((value * 1) / 1e-06) 
	case DensityMilligramPerLiter:
		return ((value * 1) / 0.001) 
	case DensityCentigramPerLiter:
		return ((value * 1) / 0.01) 
	case DensityDecigramPerLiter:
		return ((value * 1) / 0.1) 
	case DensityFemtogramPerDeciliter:
		return ((value * 1e-1) / 1e-15) 
	case DensityPicogramPerDeciliter:
		return ((value * 1e-1) / 1e-12) 
	case DensityNanogramPerDeciliter:
		return ((value * 1e-1) / 1e-09) 
	case DensityMicrogramPerDeciliter:
		return ((value * 1e-1) / 1e-06) 
	case DensityMilligramPerDeciliter:
		return ((value * 1e-1) / 0.001) 
	case DensityCentigramPerDeciliter:
		return ((value * 1e-1) / 0.01) 
	case DensityDecigramPerDeciliter:
		return ((value * 1e-1) / 0.1) 
	case DensityFemtogramPerMilliliter:
		return ((value * 1e-3) / 1e-15) 
	case DensityPicogramPerMilliliter:
		return ((value * 1e-3) / 1e-12) 
	case DensityNanogramPerMilliliter:
		return ((value * 1e-3) / 1e-09) 
	case DensityMicrogramPerMilliliter:
		return ((value * 1e-3) / 1e-06) 
	case DensityMilligramPerMilliliter:
		return ((value * 1e-3) / 0.001) 
	case DensityCentigramPerMilliliter:
		return ((value * 1e-3) / 0.01) 
	case DensityDecigramPerMilliliter:
		return ((value * 1e-3) / 0.1) 
	default:
		return math.NaN()
	}
}

func (a *Density) convertToBase(value float64, fromUnit DensityUnits) float64 {
	switch fromUnit { 
	case DensityGramPerCubicMillimeter:
		return (value / 1e-6) 
	case DensityGramPerCubicCentimeter:
		return (value / 1e-3) 
	case DensityGramPerCubicMeter:
		return (value / 1e3) 
	case DensityPoundPerCubicInch:
		return (value / 3.6127298147753e-5) 
	case DensityPoundPerCubicFoot:
		return (value / 0.062427961) 
	case DensityPoundPerCubicYard:
		return (value * (0.45359237 / 0.9144 / 0.9144 / 0.9144)) 
	case DensityTonnePerCubicMillimeter:
		return (value / 1e-12) 
	case DensityTonnePerCubicCentimeter:
		return (value / 1e-9) 
	case DensityTonnePerCubicMeter:
		return (value / 0.001) 
	case DensitySlugPerCubicFoot:
		return (value * 515.378818) 
	case DensityGramPerLiter:
		return (value / 1) 
	case DensityGramPerDeciliter:
		return (value / 1e-1) 
	case DensityGramPerMilliliter:
		return (value / 1e-3) 
	case DensityPoundPerUSGallon:
		return (value * 1.19826427e2) 
	case DensityPoundPerImperialGallon:
		return (value * 9.9776398e1) 
	case DensityKilogramPerLiter:
		return (value * 1e3) 
	case DensityTonnePerCubicFoot:
		return (value * 3.53146667214886e4) 
	case DensityTonnePerCubicInch:
		return (value * 6.10237440947323e7) 
	case DensityGramPerCubicFoot:
		return (value * 0.0353146667214886) 
	case DensityGramPerCubicInch:
		return (value * 61.0237440947323) 
	case DensityPoundPerCubicMeter:
		return (value / 2.204622621848775) 
	case DensityPoundPerCubicCentimeter:
		return (value / 2.204622621848775e-6) 
	case DensityPoundPerCubicMillimeter:
		return (value / 2.204622621848775e-9) 
	case DensitySlugPerCubicMeter:
		return (value * 14.5939) 
	case DensitySlugPerCubicCentimeter:
		return (value * 14593903) 
	case DensitySlugPerCubicMillimeter:
		return (value * 14593903000) 
	case DensitySlugPerCubicInch:
		return (value * 890574.60201535) 
	case DensityKilogramPerCubicMillimeter:
		return ((value / 1e-6) * 1000.0) 
	case DensityKilogramPerCubicCentimeter:
		return ((value / 1e-3) * 1000.0) 
	case DensityKilogramPerCubicMeter:
		return ((value / 1e3) * 1000.0) 
	case DensityMilligramPerCubicMeter:
		return ((value / 1e3) * 0.001) 
	case DensityMicrogramPerCubicMeter:
		return ((value / 1e3) * 1e-06) 
	case DensityKilopoundPerCubicInch:
		return ((value / 3.6127298147753e-5) * 1000.0) 
	case DensityKilopoundPerCubicFoot:
		return ((value / 0.062427961) * 1000.0) 
	case DensityKilopoundPerCubicYard:
		return ((value * (0.45359237 / 0.9144 / 0.9144 / 0.9144)) * 1000.0) 
	case DensityFemtogramPerLiter:
		return ((value / 1) * 1e-15) 
	case DensityPicogramPerLiter:
		return ((value / 1) * 1e-12) 
	case DensityNanogramPerLiter:
		return ((value / 1) * 1e-09) 
	case DensityMicrogramPerLiter:
		return ((value / 1) * 1e-06) 
	case DensityMilligramPerLiter:
		return ((value / 1) * 0.001) 
	case DensityCentigramPerLiter:
		return ((value / 1) * 0.01) 
	case DensityDecigramPerLiter:
		return ((value / 1) * 0.1) 
	case DensityFemtogramPerDeciliter:
		return ((value / 1e-1) * 1e-15) 
	case DensityPicogramPerDeciliter:
		return ((value / 1e-1) * 1e-12) 
	case DensityNanogramPerDeciliter:
		return ((value / 1e-1) * 1e-09) 
	case DensityMicrogramPerDeciliter:
		return ((value / 1e-1) * 1e-06) 
	case DensityMilligramPerDeciliter:
		return ((value / 1e-1) * 0.001) 
	case DensityCentigramPerDeciliter:
		return ((value / 1e-1) * 0.01) 
	case DensityDecigramPerDeciliter:
		return ((value / 1e-1) * 0.1) 
	case DensityFemtogramPerMilliliter:
		return ((value / 1e-3) * 1e-15) 
	case DensityPicogramPerMilliliter:
		return ((value / 1e-3) * 1e-12) 
	case DensityNanogramPerMilliliter:
		return ((value / 1e-3) * 1e-09) 
	case DensityMicrogramPerMilliliter:
		return ((value / 1e-3) * 1e-06) 
	case DensityMilligramPerMilliliter:
		return ((value / 1e-3) * 0.001) 
	case DensityCentigramPerMilliliter:
		return ((value / 1e-3) * 0.01) 
	case DensityDecigramPerMilliliter:
		return ((value / 1e-3) * 0.1) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Density in the default unit (KilogramPerCubicMeter),
// formatted to two decimal places.
func (a Density) String() string {
	return a.ToString(DensityKilogramPerCubicMeter, 2)
}

// ToString formats the Density value as a string with the specified unit and fractional digits.
// It converts the Density to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Density value will be converted (e.g., KilogramPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Density with the unit abbreviation.
func (a *Density) ToString(unit DensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetDensityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetDensityAbbreviation(unit))
}

// Equals checks if the given Density is equal to the current Density.
//
// Parameters:
//    other: The Density to compare against.
//
// Returns:
//    bool: Returns true if both Density are equal, false otherwise.
func (a *Density) Equals(other *Density) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Density with another Density.
// It returns -1 if the current Density is less than the other Density, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Density to compare against.
//
// Returns:
//    int: -1 if the current Density is less, 1 if greater, and 0 if equal.
func (a *Density) CompareTo(other *Density) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Density to the current Density and returns the result.
// The result is a new Density instance with the sum of the values.
//
// Parameters:
//    other: The Density to add to the current Density.
//
// Returns:
//    *Density: A new Density instance representing the sum of both Density.
func (a *Density) Add(other *Density) *Density {
	return &Density{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Density from the current Density and returns the result.
// The result is a new Density instance with the difference of the values.
//
// Parameters:
//    other: The Density to subtract from the current Density.
//
// Returns:
//    *Density: A new Density instance representing the difference of both Density.
func (a *Density) Subtract(other *Density) *Density {
	return &Density{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Density by the given Density and returns the result.
// The result is a new Density instance with the product of the values.
//
// Parameters:
//    other: The Density to multiply with the current Density.
//
// Returns:
//    *Density: A new Density instance representing the product of both Density.
func (a *Density) Multiply(other *Density) *Density {
	return &Density{value: a.value * other.BaseValue()}
}

// Divide divides the current Density by the given Density and returns the result.
// The result is a new Density instance with the quotient of the values.
//
// Parameters:
//    other: The Density to divide the current Density by.
//
// Returns:
//    *Density: A new Density instance representing the quotient of both Density.
func (a *Density) Divide(other *Density) *Density {
	return &Density{value: a.value / other.BaseValue()}
}

// GetDensityAbbreviation gets the unit abbreviation.
func GetDensityAbbreviation(unit DensityUnits) string {
	switch unit { 
	case DensityGramPerCubicMillimeter:
		return "g/mm³" 
	case DensityGramPerCubicCentimeter:
		return "g/cm³" 
	case DensityGramPerCubicMeter:
		return "g/m³" 
	case DensityPoundPerCubicInch:
		return "lb/in³" 
	case DensityPoundPerCubicFoot:
		return "lb/ft³" 
	case DensityPoundPerCubicYard:
		return "lb/yd³" 
	case DensityTonnePerCubicMillimeter:
		return "t/mm³" 
	case DensityTonnePerCubicCentimeter:
		return "t/cm³" 
	case DensityTonnePerCubicMeter:
		return "t/m³" 
	case DensitySlugPerCubicFoot:
		return "slug/ft³" 
	case DensityGramPerLiter:
		return "g/L" 
	case DensityGramPerDeciliter:
		return "g/dl" 
	case DensityGramPerMilliliter:
		return "g/ml" 
	case DensityPoundPerUSGallon:
		return "ppg (U.S.)" 
	case DensityPoundPerImperialGallon:
		return "ppg (imp.)" 
	case DensityKilogramPerLiter:
		return "kg/l" 
	case DensityTonnePerCubicFoot:
		return "t/ft³" 
	case DensityTonnePerCubicInch:
		return "t/in³" 
	case DensityGramPerCubicFoot:
		return "g/ft³" 
	case DensityGramPerCubicInch:
		return "g/in³" 
	case DensityPoundPerCubicMeter:
		return "lb/m³" 
	case DensityPoundPerCubicCentimeter:
		return "lb/cm³" 
	case DensityPoundPerCubicMillimeter:
		return "lb/mm³" 
	case DensitySlugPerCubicMeter:
		return "slug/m³" 
	case DensitySlugPerCubicCentimeter:
		return "slug/cm³" 
	case DensitySlugPerCubicMillimeter:
		return "slug/mm³" 
	case DensitySlugPerCubicInch:
		return "slug/in³" 
	case DensityKilogramPerCubicMillimeter:
		return "kg/mm³" 
	case DensityKilogramPerCubicCentimeter:
		return "kg/cm³" 
	case DensityKilogramPerCubicMeter:
		return "kg/m³" 
	case DensityMilligramPerCubicMeter:
		return "mg/m³" 
	case DensityMicrogramPerCubicMeter:
		return "μg/m³" 
	case DensityKilopoundPerCubicInch:
		return "klb/in³" 
	case DensityKilopoundPerCubicFoot:
		return "klb/ft³" 
	case DensityKilopoundPerCubicYard:
		return "klb/yd³" 
	case DensityFemtogramPerLiter:
		return "fg/L" 
	case DensityPicogramPerLiter:
		return "pg/L" 
	case DensityNanogramPerLiter:
		return "ng/L" 
	case DensityMicrogramPerLiter:
		return "μg/L" 
	case DensityMilligramPerLiter:
		return "mg/L" 
	case DensityCentigramPerLiter:
		return "cg/L" 
	case DensityDecigramPerLiter:
		return "dg/L" 
	case DensityFemtogramPerDeciliter:
		return "fg/dl" 
	case DensityPicogramPerDeciliter:
		return "pg/dl" 
	case DensityNanogramPerDeciliter:
		return "ng/dl" 
	case DensityMicrogramPerDeciliter:
		return "μg/dl" 
	case DensityMilligramPerDeciliter:
		return "mg/dl" 
	case DensityCentigramPerDeciliter:
		return "cg/dl" 
	case DensityDecigramPerDeciliter:
		return "dg/dl" 
	case DensityFemtogramPerMilliliter:
		return "fg/ml" 
	case DensityPicogramPerMilliliter:
		return "pg/ml" 
	case DensityNanogramPerMilliliter:
		return "ng/ml" 
	case DensityMicrogramPerMilliliter:
		return "μg/ml" 
	case DensityMilligramPerMilliliter:
		return "mg/ml" 
	case DensityCentigramPerMilliliter:
		return "cg/ml" 
	case DensityDecigramPerMilliliter:
		return "dg/ml" 
	default:
		return ""
	}
}