package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// DensityUnits enumeration
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

// DensityDto represents an Density
type DensityDto struct {
	Value float64
	Unit  DensityUnits
}

// DensityDtoFactory struct to group related functions
type DensityDtoFactory struct{}

func (udf DensityDtoFactory) FromJSON(data []byte) (*DensityDto, error) {
	a := DensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a DensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Density struct
type Density struct {
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

// DensityFactory struct to group related functions
type DensityFactory struct{}

func (uf DensityFactory) CreateDensity(value float64, unit DensityUnits) (*Density, error) {
	return newDensity(value, unit)
}

func (uf DensityFactory) FromDto(dto DensityDto) (*Density, error) {
	return newDensity(dto.Value, dto.Unit)
}

func (uf DensityFactory) FromDtoJSON(data []byte) (*Density, error) {
	unitDto, err := DensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return DensityFactory{}.FromDto(*unitDto)
}


// FromGramPerCubicMillimeter creates a new Density instance from GramPerCubicMillimeter.
func (uf DensityFactory) FromGramsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicMillimeter)
}

// FromGramPerCubicCentimeter creates a new Density instance from GramPerCubicCentimeter.
func (uf DensityFactory) FromGramsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicCentimeter)
}

// FromGramPerCubicMeter creates a new Density instance from GramPerCubicMeter.
func (uf DensityFactory) FromGramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicMeter)
}

// FromPoundPerCubicInch creates a new Density instance from PoundPerCubicInch.
func (uf DensityFactory) FromPoundsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicInch)
}

// FromPoundPerCubicFoot creates a new Density instance from PoundPerCubicFoot.
func (uf DensityFactory) FromPoundsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicFoot)
}

// FromPoundPerCubicYard creates a new Density instance from PoundPerCubicYard.
func (uf DensityFactory) FromPoundsPerCubicYard(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicYard)
}

// FromTonnePerCubicMillimeter creates a new Density instance from TonnePerCubicMillimeter.
func (uf DensityFactory) FromTonnesPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicMillimeter)
}

// FromTonnePerCubicCentimeter creates a new Density instance from TonnePerCubicCentimeter.
func (uf DensityFactory) FromTonnesPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicCentimeter)
}

// FromTonnePerCubicMeter creates a new Density instance from TonnePerCubicMeter.
func (uf DensityFactory) FromTonnesPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicMeter)
}

// FromSlugPerCubicFoot creates a new Density instance from SlugPerCubicFoot.
func (uf DensityFactory) FromSlugsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicFoot)
}

// FromGramPerLiter creates a new Density instance from GramPerLiter.
func (uf DensityFactory) FromGramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerLiter)
}

// FromGramPerDeciliter creates a new Density instance from GramPerDeciliter.
func (uf DensityFactory) FromGramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerDeciliter)
}

// FromGramPerMilliliter creates a new Density instance from GramPerMilliliter.
func (uf DensityFactory) FromGramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerMilliliter)
}

// FromPoundPerUSGallon creates a new Density instance from PoundPerUSGallon.
func (uf DensityFactory) FromPoundsPerUSGallon(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerUSGallon)
}

// FromPoundPerImperialGallon creates a new Density instance from PoundPerImperialGallon.
func (uf DensityFactory) FromPoundsPerImperialGallon(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerImperialGallon)
}

// FromKilogramPerLiter creates a new Density instance from KilogramPerLiter.
func (uf DensityFactory) FromKilogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerLiter)
}

// FromTonnePerCubicFoot creates a new Density instance from TonnePerCubicFoot.
func (uf DensityFactory) FromTonnesPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicFoot)
}

// FromTonnePerCubicInch creates a new Density instance from TonnePerCubicInch.
func (uf DensityFactory) FromTonnesPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityTonnePerCubicInch)
}

// FromGramPerCubicFoot creates a new Density instance from GramPerCubicFoot.
func (uf DensityFactory) FromGramsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicFoot)
}

// FromGramPerCubicInch creates a new Density instance from GramPerCubicInch.
func (uf DensityFactory) FromGramsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityGramPerCubicInch)
}

// FromPoundPerCubicMeter creates a new Density instance from PoundPerCubicMeter.
func (uf DensityFactory) FromPoundsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicMeter)
}

// FromPoundPerCubicCentimeter creates a new Density instance from PoundPerCubicCentimeter.
func (uf DensityFactory) FromPoundsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicCentimeter)
}

// FromPoundPerCubicMillimeter creates a new Density instance from PoundPerCubicMillimeter.
func (uf DensityFactory) FromPoundsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityPoundPerCubicMillimeter)
}

// FromSlugPerCubicMeter creates a new Density instance from SlugPerCubicMeter.
func (uf DensityFactory) FromSlugsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicMeter)
}

// FromSlugPerCubicCentimeter creates a new Density instance from SlugPerCubicCentimeter.
func (uf DensityFactory) FromSlugsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicCentimeter)
}

// FromSlugPerCubicMillimeter creates a new Density instance from SlugPerCubicMillimeter.
func (uf DensityFactory) FromSlugsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicMillimeter)
}

// FromSlugPerCubicInch creates a new Density instance from SlugPerCubicInch.
func (uf DensityFactory) FromSlugsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensitySlugPerCubicInch)
}

// FromKilogramPerCubicMillimeter creates a new Density instance from KilogramPerCubicMillimeter.
func (uf DensityFactory) FromKilogramsPerCubicMillimeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicMillimeter)
}

// FromKilogramPerCubicCentimeter creates a new Density instance from KilogramPerCubicCentimeter.
func (uf DensityFactory) FromKilogramsPerCubicCentimeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicCentimeter)
}

// FromKilogramPerCubicMeter creates a new Density instance from KilogramPerCubicMeter.
func (uf DensityFactory) FromKilogramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityKilogramPerCubicMeter)
}

// FromMilligramPerCubicMeter creates a new Density instance from MilligramPerCubicMeter.
func (uf DensityFactory) FromMilligramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerCubicMeter)
}

// FromMicrogramPerCubicMeter creates a new Density instance from MicrogramPerCubicMeter.
func (uf DensityFactory) FromMicrogramsPerCubicMeter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerCubicMeter)
}

// FromKilopoundPerCubicInch creates a new Density instance from KilopoundPerCubicInch.
func (uf DensityFactory) FromKilopoundsPerCubicInch(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicInch)
}

// FromKilopoundPerCubicFoot creates a new Density instance from KilopoundPerCubicFoot.
func (uf DensityFactory) FromKilopoundsPerCubicFoot(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicFoot)
}

// FromKilopoundPerCubicYard creates a new Density instance from KilopoundPerCubicYard.
func (uf DensityFactory) FromKilopoundsPerCubicYard(value float64) (*Density, error) {
	return newDensity(value, DensityKilopoundPerCubicYard)
}

// FromFemtogramPerLiter creates a new Density instance from FemtogramPerLiter.
func (uf DensityFactory) FromFemtogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerLiter)
}

// FromPicogramPerLiter creates a new Density instance from PicogramPerLiter.
func (uf DensityFactory) FromPicogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerLiter)
}

// FromNanogramPerLiter creates a new Density instance from NanogramPerLiter.
func (uf DensityFactory) FromNanogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerLiter)
}

// FromMicrogramPerLiter creates a new Density instance from MicrogramPerLiter.
func (uf DensityFactory) FromMicrogramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerLiter)
}

// FromMilligramPerLiter creates a new Density instance from MilligramPerLiter.
func (uf DensityFactory) FromMilligramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerLiter)
}

// FromCentigramPerLiter creates a new Density instance from CentigramPerLiter.
func (uf DensityFactory) FromCentigramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerLiter)
}

// FromDecigramPerLiter creates a new Density instance from DecigramPerLiter.
func (uf DensityFactory) FromDecigramsPerLiter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerLiter)
}

// FromFemtogramPerDeciliter creates a new Density instance from FemtogramPerDeciliter.
func (uf DensityFactory) FromFemtogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerDeciliter)
}

// FromPicogramPerDeciliter creates a new Density instance from PicogramPerDeciliter.
func (uf DensityFactory) FromPicogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerDeciliter)
}

// FromNanogramPerDeciliter creates a new Density instance from NanogramPerDeciliter.
func (uf DensityFactory) FromNanogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerDeciliter)
}

// FromMicrogramPerDeciliter creates a new Density instance from MicrogramPerDeciliter.
func (uf DensityFactory) FromMicrogramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerDeciliter)
}

// FromMilligramPerDeciliter creates a new Density instance from MilligramPerDeciliter.
func (uf DensityFactory) FromMilligramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerDeciliter)
}

// FromCentigramPerDeciliter creates a new Density instance from CentigramPerDeciliter.
func (uf DensityFactory) FromCentigramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerDeciliter)
}

// FromDecigramPerDeciliter creates a new Density instance from DecigramPerDeciliter.
func (uf DensityFactory) FromDecigramsPerDeciLiter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerDeciliter)
}

// FromFemtogramPerMilliliter creates a new Density instance from FemtogramPerMilliliter.
func (uf DensityFactory) FromFemtogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityFemtogramPerMilliliter)
}

// FromPicogramPerMilliliter creates a new Density instance from PicogramPerMilliliter.
func (uf DensityFactory) FromPicogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityPicogramPerMilliliter)
}

// FromNanogramPerMilliliter creates a new Density instance from NanogramPerMilliliter.
func (uf DensityFactory) FromNanogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityNanogramPerMilliliter)
}

// FromMicrogramPerMilliliter creates a new Density instance from MicrogramPerMilliliter.
func (uf DensityFactory) FromMicrogramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityMicrogramPerMilliliter)
}

// FromMilligramPerMilliliter creates a new Density instance from MilligramPerMilliliter.
func (uf DensityFactory) FromMilligramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityMilligramPerMilliliter)
}

// FromCentigramPerMilliliter creates a new Density instance from CentigramPerMilliliter.
func (uf DensityFactory) FromCentigramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityCentigramPerMilliliter)
}

// FromDecigramPerMilliliter creates a new Density instance from DecigramPerMilliliter.
func (uf DensityFactory) FromDecigramsPerMilliliter(value float64) (*Density, error) {
	return newDensity(value, DensityDecigramPerMilliliter)
}




// newDensity creates a new Density.
func newDensity(value float64, fromUnit DensityUnits) (*Density, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Density{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Density in KilogramPerCubicMeter.
func (a *Density) BaseValue() float64 {
	return a.value
}


// GramPerCubicMillimeter returns the value in GramPerCubicMillimeter.
func (a *Density) GramsPerCubicMillimeter() float64 {
	if a.grams_per_cubic_millimeterLazy != nil {
		return *a.grams_per_cubic_millimeterLazy
	}
	grams_per_cubic_millimeter := a.convertFromBase(DensityGramPerCubicMillimeter)
	a.grams_per_cubic_millimeterLazy = &grams_per_cubic_millimeter
	return grams_per_cubic_millimeter
}

// GramPerCubicCentimeter returns the value in GramPerCubicCentimeter.
func (a *Density) GramsPerCubicCentimeter() float64 {
	if a.grams_per_cubic_centimeterLazy != nil {
		return *a.grams_per_cubic_centimeterLazy
	}
	grams_per_cubic_centimeter := a.convertFromBase(DensityGramPerCubicCentimeter)
	a.grams_per_cubic_centimeterLazy = &grams_per_cubic_centimeter
	return grams_per_cubic_centimeter
}

// GramPerCubicMeter returns the value in GramPerCubicMeter.
func (a *Density) GramsPerCubicMeter() float64 {
	if a.grams_per_cubic_meterLazy != nil {
		return *a.grams_per_cubic_meterLazy
	}
	grams_per_cubic_meter := a.convertFromBase(DensityGramPerCubicMeter)
	a.grams_per_cubic_meterLazy = &grams_per_cubic_meter
	return grams_per_cubic_meter
}

// PoundPerCubicInch returns the value in PoundPerCubicInch.
func (a *Density) PoundsPerCubicInch() float64 {
	if a.pounds_per_cubic_inchLazy != nil {
		return *a.pounds_per_cubic_inchLazy
	}
	pounds_per_cubic_inch := a.convertFromBase(DensityPoundPerCubicInch)
	a.pounds_per_cubic_inchLazy = &pounds_per_cubic_inch
	return pounds_per_cubic_inch
}

// PoundPerCubicFoot returns the value in PoundPerCubicFoot.
func (a *Density) PoundsPerCubicFoot() float64 {
	if a.pounds_per_cubic_footLazy != nil {
		return *a.pounds_per_cubic_footLazy
	}
	pounds_per_cubic_foot := a.convertFromBase(DensityPoundPerCubicFoot)
	a.pounds_per_cubic_footLazy = &pounds_per_cubic_foot
	return pounds_per_cubic_foot
}

// PoundPerCubicYard returns the value in PoundPerCubicYard.
func (a *Density) PoundsPerCubicYard() float64 {
	if a.pounds_per_cubic_yardLazy != nil {
		return *a.pounds_per_cubic_yardLazy
	}
	pounds_per_cubic_yard := a.convertFromBase(DensityPoundPerCubicYard)
	a.pounds_per_cubic_yardLazy = &pounds_per_cubic_yard
	return pounds_per_cubic_yard
}

// TonnePerCubicMillimeter returns the value in TonnePerCubicMillimeter.
func (a *Density) TonnesPerCubicMillimeter() float64 {
	if a.tonnes_per_cubic_millimeterLazy != nil {
		return *a.tonnes_per_cubic_millimeterLazy
	}
	tonnes_per_cubic_millimeter := a.convertFromBase(DensityTonnePerCubicMillimeter)
	a.tonnes_per_cubic_millimeterLazy = &tonnes_per_cubic_millimeter
	return tonnes_per_cubic_millimeter
}

// TonnePerCubicCentimeter returns the value in TonnePerCubicCentimeter.
func (a *Density) TonnesPerCubicCentimeter() float64 {
	if a.tonnes_per_cubic_centimeterLazy != nil {
		return *a.tonnes_per_cubic_centimeterLazy
	}
	tonnes_per_cubic_centimeter := a.convertFromBase(DensityTonnePerCubicCentimeter)
	a.tonnes_per_cubic_centimeterLazy = &tonnes_per_cubic_centimeter
	return tonnes_per_cubic_centimeter
}

// TonnePerCubicMeter returns the value in TonnePerCubicMeter.
func (a *Density) TonnesPerCubicMeter() float64 {
	if a.tonnes_per_cubic_meterLazy != nil {
		return *a.tonnes_per_cubic_meterLazy
	}
	tonnes_per_cubic_meter := a.convertFromBase(DensityTonnePerCubicMeter)
	a.tonnes_per_cubic_meterLazy = &tonnes_per_cubic_meter
	return tonnes_per_cubic_meter
}

// SlugPerCubicFoot returns the value in SlugPerCubicFoot.
func (a *Density) SlugsPerCubicFoot() float64 {
	if a.slugs_per_cubic_footLazy != nil {
		return *a.slugs_per_cubic_footLazy
	}
	slugs_per_cubic_foot := a.convertFromBase(DensitySlugPerCubicFoot)
	a.slugs_per_cubic_footLazy = &slugs_per_cubic_foot
	return slugs_per_cubic_foot
}

// GramPerLiter returns the value in GramPerLiter.
func (a *Density) GramsPerLiter() float64 {
	if a.grams_per_literLazy != nil {
		return *a.grams_per_literLazy
	}
	grams_per_liter := a.convertFromBase(DensityGramPerLiter)
	a.grams_per_literLazy = &grams_per_liter
	return grams_per_liter
}

// GramPerDeciliter returns the value in GramPerDeciliter.
func (a *Density) GramsPerDeciLiter() float64 {
	if a.grams_per_deci_literLazy != nil {
		return *a.grams_per_deci_literLazy
	}
	grams_per_deci_liter := a.convertFromBase(DensityGramPerDeciliter)
	a.grams_per_deci_literLazy = &grams_per_deci_liter
	return grams_per_deci_liter
}

// GramPerMilliliter returns the value in GramPerMilliliter.
func (a *Density) GramsPerMilliliter() float64 {
	if a.grams_per_milliliterLazy != nil {
		return *a.grams_per_milliliterLazy
	}
	grams_per_milliliter := a.convertFromBase(DensityGramPerMilliliter)
	a.grams_per_milliliterLazy = &grams_per_milliliter
	return grams_per_milliliter
}

// PoundPerUSGallon returns the value in PoundPerUSGallon.
func (a *Density) PoundsPerUSGallon() float64 {
	if a.pounds_per_us_gallonLazy != nil {
		return *a.pounds_per_us_gallonLazy
	}
	pounds_per_us_gallon := a.convertFromBase(DensityPoundPerUSGallon)
	a.pounds_per_us_gallonLazy = &pounds_per_us_gallon
	return pounds_per_us_gallon
}

// PoundPerImperialGallon returns the value in PoundPerImperialGallon.
func (a *Density) PoundsPerImperialGallon() float64 {
	if a.pounds_per_imperial_gallonLazy != nil {
		return *a.pounds_per_imperial_gallonLazy
	}
	pounds_per_imperial_gallon := a.convertFromBase(DensityPoundPerImperialGallon)
	a.pounds_per_imperial_gallonLazy = &pounds_per_imperial_gallon
	return pounds_per_imperial_gallon
}

// KilogramPerLiter returns the value in KilogramPerLiter.
func (a *Density) KilogramsPerLiter() float64 {
	if a.kilograms_per_literLazy != nil {
		return *a.kilograms_per_literLazy
	}
	kilograms_per_liter := a.convertFromBase(DensityKilogramPerLiter)
	a.kilograms_per_literLazy = &kilograms_per_liter
	return kilograms_per_liter
}

// TonnePerCubicFoot returns the value in TonnePerCubicFoot.
func (a *Density) TonnesPerCubicFoot() float64 {
	if a.tonnes_per_cubic_footLazy != nil {
		return *a.tonnes_per_cubic_footLazy
	}
	tonnes_per_cubic_foot := a.convertFromBase(DensityTonnePerCubicFoot)
	a.tonnes_per_cubic_footLazy = &tonnes_per_cubic_foot
	return tonnes_per_cubic_foot
}

// TonnePerCubicInch returns the value in TonnePerCubicInch.
func (a *Density) TonnesPerCubicInch() float64 {
	if a.tonnes_per_cubic_inchLazy != nil {
		return *a.tonnes_per_cubic_inchLazy
	}
	tonnes_per_cubic_inch := a.convertFromBase(DensityTonnePerCubicInch)
	a.tonnes_per_cubic_inchLazy = &tonnes_per_cubic_inch
	return tonnes_per_cubic_inch
}

// GramPerCubicFoot returns the value in GramPerCubicFoot.
func (a *Density) GramsPerCubicFoot() float64 {
	if a.grams_per_cubic_footLazy != nil {
		return *a.grams_per_cubic_footLazy
	}
	grams_per_cubic_foot := a.convertFromBase(DensityGramPerCubicFoot)
	a.grams_per_cubic_footLazy = &grams_per_cubic_foot
	return grams_per_cubic_foot
}

// GramPerCubicInch returns the value in GramPerCubicInch.
func (a *Density) GramsPerCubicInch() float64 {
	if a.grams_per_cubic_inchLazy != nil {
		return *a.grams_per_cubic_inchLazy
	}
	grams_per_cubic_inch := a.convertFromBase(DensityGramPerCubicInch)
	a.grams_per_cubic_inchLazy = &grams_per_cubic_inch
	return grams_per_cubic_inch
}

// PoundPerCubicMeter returns the value in PoundPerCubicMeter.
func (a *Density) PoundsPerCubicMeter() float64 {
	if a.pounds_per_cubic_meterLazy != nil {
		return *a.pounds_per_cubic_meterLazy
	}
	pounds_per_cubic_meter := a.convertFromBase(DensityPoundPerCubicMeter)
	a.pounds_per_cubic_meterLazy = &pounds_per_cubic_meter
	return pounds_per_cubic_meter
}

// PoundPerCubicCentimeter returns the value in PoundPerCubicCentimeter.
func (a *Density) PoundsPerCubicCentimeter() float64 {
	if a.pounds_per_cubic_centimeterLazy != nil {
		return *a.pounds_per_cubic_centimeterLazy
	}
	pounds_per_cubic_centimeter := a.convertFromBase(DensityPoundPerCubicCentimeter)
	a.pounds_per_cubic_centimeterLazy = &pounds_per_cubic_centimeter
	return pounds_per_cubic_centimeter
}

// PoundPerCubicMillimeter returns the value in PoundPerCubicMillimeter.
func (a *Density) PoundsPerCubicMillimeter() float64 {
	if a.pounds_per_cubic_millimeterLazy != nil {
		return *a.pounds_per_cubic_millimeterLazy
	}
	pounds_per_cubic_millimeter := a.convertFromBase(DensityPoundPerCubicMillimeter)
	a.pounds_per_cubic_millimeterLazy = &pounds_per_cubic_millimeter
	return pounds_per_cubic_millimeter
}

// SlugPerCubicMeter returns the value in SlugPerCubicMeter.
func (a *Density) SlugsPerCubicMeter() float64 {
	if a.slugs_per_cubic_meterLazy != nil {
		return *a.slugs_per_cubic_meterLazy
	}
	slugs_per_cubic_meter := a.convertFromBase(DensitySlugPerCubicMeter)
	a.slugs_per_cubic_meterLazy = &slugs_per_cubic_meter
	return slugs_per_cubic_meter
}

// SlugPerCubicCentimeter returns the value in SlugPerCubicCentimeter.
func (a *Density) SlugsPerCubicCentimeter() float64 {
	if a.slugs_per_cubic_centimeterLazy != nil {
		return *a.slugs_per_cubic_centimeterLazy
	}
	slugs_per_cubic_centimeter := a.convertFromBase(DensitySlugPerCubicCentimeter)
	a.slugs_per_cubic_centimeterLazy = &slugs_per_cubic_centimeter
	return slugs_per_cubic_centimeter
}

// SlugPerCubicMillimeter returns the value in SlugPerCubicMillimeter.
func (a *Density) SlugsPerCubicMillimeter() float64 {
	if a.slugs_per_cubic_millimeterLazy != nil {
		return *a.slugs_per_cubic_millimeterLazy
	}
	slugs_per_cubic_millimeter := a.convertFromBase(DensitySlugPerCubicMillimeter)
	a.slugs_per_cubic_millimeterLazy = &slugs_per_cubic_millimeter
	return slugs_per_cubic_millimeter
}

// SlugPerCubicInch returns the value in SlugPerCubicInch.
func (a *Density) SlugsPerCubicInch() float64 {
	if a.slugs_per_cubic_inchLazy != nil {
		return *a.slugs_per_cubic_inchLazy
	}
	slugs_per_cubic_inch := a.convertFromBase(DensitySlugPerCubicInch)
	a.slugs_per_cubic_inchLazy = &slugs_per_cubic_inch
	return slugs_per_cubic_inch
}

// KilogramPerCubicMillimeter returns the value in KilogramPerCubicMillimeter.
func (a *Density) KilogramsPerCubicMillimeter() float64 {
	if a.kilograms_per_cubic_millimeterLazy != nil {
		return *a.kilograms_per_cubic_millimeterLazy
	}
	kilograms_per_cubic_millimeter := a.convertFromBase(DensityKilogramPerCubicMillimeter)
	a.kilograms_per_cubic_millimeterLazy = &kilograms_per_cubic_millimeter
	return kilograms_per_cubic_millimeter
}

// KilogramPerCubicCentimeter returns the value in KilogramPerCubicCentimeter.
func (a *Density) KilogramsPerCubicCentimeter() float64 {
	if a.kilograms_per_cubic_centimeterLazy != nil {
		return *a.kilograms_per_cubic_centimeterLazy
	}
	kilograms_per_cubic_centimeter := a.convertFromBase(DensityKilogramPerCubicCentimeter)
	a.kilograms_per_cubic_centimeterLazy = &kilograms_per_cubic_centimeter
	return kilograms_per_cubic_centimeter
}

// KilogramPerCubicMeter returns the value in KilogramPerCubicMeter.
func (a *Density) KilogramsPerCubicMeter() float64 {
	if a.kilograms_per_cubic_meterLazy != nil {
		return *a.kilograms_per_cubic_meterLazy
	}
	kilograms_per_cubic_meter := a.convertFromBase(DensityKilogramPerCubicMeter)
	a.kilograms_per_cubic_meterLazy = &kilograms_per_cubic_meter
	return kilograms_per_cubic_meter
}

// MilligramPerCubicMeter returns the value in MilligramPerCubicMeter.
func (a *Density) MilligramsPerCubicMeter() float64 {
	if a.milligrams_per_cubic_meterLazy != nil {
		return *a.milligrams_per_cubic_meterLazy
	}
	milligrams_per_cubic_meter := a.convertFromBase(DensityMilligramPerCubicMeter)
	a.milligrams_per_cubic_meterLazy = &milligrams_per_cubic_meter
	return milligrams_per_cubic_meter
}

// MicrogramPerCubicMeter returns the value in MicrogramPerCubicMeter.
func (a *Density) MicrogramsPerCubicMeter() float64 {
	if a.micrograms_per_cubic_meterLazy != nil {
		return *a.micrograms_per_cubic_meterLazy
	}
	micrograms_per_cubic_meter := a.convertFromBase(DensityMicrogramPerCubicMeter)
	a.micrograms_per_cubic_meterLazy = &micrograms_per_cubic_meter
	return micrograms_per_cubic_meter
}

// KilopoundPerCubicInch returns the value in KilopoundPerCubicInch.
func (a *Density) KilopoundsPerCubicInch() float64 {
	if a.kilopounds_per_cubic_inchLazy != nil {
		return *a.kilopounds_per_cubic_inchLazy
	}
	kilopounds_per_cubic_inch := a.convertFromBase(DensityKilopoundPerCubicInch)
	a.kilopounds_per_cubic_inchLazy = &kilopounds_per_cubic_inch
	return kilopounds_per_cubic_inch
}

// KilopoundPerCubicFoot returns the value in KilopoundPerCubicFoot.
func (a *Density) KilopoundsPerCubicFoot() float64 {
	if a.kilopounds_per_cubic_footLazy != nil {
		return *a.kilopounds_per_cubic_footLazy
	}
	kilopounds_per_cubic_foot := a.convertFromBase(DensityKilopoundPerCubicFoot)
	a.kilopounds_per_cubic_footLazy = &kilopounds_per_cubic_foot
	return kilopounds_per_cubic_foot
}

// KilopoundPerCubicYard returns the value in KilopoundPerCubicYard.
func (a *Density) KilopoundsPerCubicYard() float64 {
	if a.kilopounds_per_cubic_yardLazy != nil {
		return *a.kilopounds_per_cubic_yardLazy
	}
	kilopounds_per_cubic_yard := a.convertFromBase(DensityKilopoundPerCubicYard)
	a.kilopounds_per_cubic_yardLazy = &kilopounds_per_cubic_yard
	return kilopounds_per_cubic_yard
}

// FemtogramPerLiter returns the value in FemtogramPerLiter.
func (a *Density) FemtogramsPerLiter() float64 {
	if a.femtograms_per_literLazy != nil {
		return *a.femtograms_per_literLazy
	}
	femtograms_per_liter := a.convertFromBase(DensityFemtogramPerLiter)
	a.femtograms_per_literLazy = &femtograms_per_liter
	return femtograms_per_liter
}

// PicogramPerLiter returns the value in PicogramPerLiter.
func (a *Density) PicogramsPerLiter() float64 {
	if a.picograms_per_literLazy != nil {
		return *a.picograms_per_literLazy
	}
	picograms_per_liter := a.convertFromBase(DensityPicogramPerLiter)
	a.picograms_per_literLazy = &picograms_per_liter
	return picograms_per_liter
}

// NanogramPerLiter returns the value in NanogramPerLiter.
func (a *Density) NanogramsPerLiter() float64 {
	if a.nanograms_per_literLazy != nil {
		return *a.nanograms_per_literLazy
	}
	nanograms_per_liter := a.convertFromBase(DensityNanogramPerLiter)
	a.nanograms_per_literLazy = &nanograms_per_liter
	return nanograms_per_liter
}

// MicrogramPerLiter returns the value in MicrogramPerLiter.
func (a *Density) MicrogramsPerLiter() float64 {
	if a.micrograms_per_literLazy != nil {
		return *a.micrograms_per_literLazy
	}
	micrograms_per_liter := a.convertFromBase(DensityMicrogramPerLiter)
	a.micrograms_per_literLazy = &micrograms_per_liter
	return micrograms_per_liter
}

// MilligramPerLiter returns the value in MilligramPerLiter.
func (a *Density) MilligramsPerLiter() float64 {
	if a.milligrams_per_literLazy != nil {
		return *a.milligrams_per_literLazy
	}
	milligrams_per_liter := a.convertFromBase(DensityMilligramPerLiter)
	a.milligrams_per_literLazy = &milligrams_per_liter
	return milligrams_per_liter
}

// CentigramPerLiter returns the value in CentigramPerLiter.
func (a *Density) CentigramsPerLiter() float64 {
	if a.centigrams_per_literLazy != nil {
		return *a.centigrams_per_literLazy
	}
	centigrams_per_liter := a.convertFromBase(DensityCentigramPerLiter)
	a.centigrams_per_literLazy = &centigrams_per_liter
	return centigrams_per_liter
}

// DecigramPerLiter returns the value in DecigramPerLiter.
func (a *Density) DecigramsPerLiter() float64 {
	if a.decigrams_per_literLazy != nil {
		return *a.decigrams_per_literLazy
	}
	decigrams_per_liter := a.convertFromBase(DensityDecigramPerLiter)
	a.decigrams_per_literLazy = &decigrams_per_liter
	return decigrams_per_liter
}

// FemtogramPerDeciliter returns the value in FemtogramPerDeciliter.
func (a *Density) FemtogramsPerDeciLiter() float64 {
	if a.femtograms_per_deci_literLazy != nil {
		return *a.femtograms_per_deci_literLazy
	}
	femtograms_per_deci_liter := a.convertFromBase(DensityFemtogramPerDeciliter)
	a.femtograms_per_deci_literLazy = &femtograms_per_deci_liter
	return femtograms_per_deci_liter
}

// PicogramPerDeciliter returns the value in PicogramPerDeciliter.
func (a *Density) PicogramsPerDeciLiter() float64 {
	if a.picograms_per_deci_literLazy != nil {
		return *a.picograms_per_deci_literLazy
	}
	picograms_per_deci_liter := a.convertFromBase(DensityPicogramPerDeciliter)
	a.picograms_per_deci_literLazy = &picograms_per_deci_liter
	return picograms_per_deci_liter
}

// NanogramPerDeciliter returns the value in NanogramPerDeciliter.
func (a *Density) NanogramsPerDeciLiter() float64 {
	if a.nanograms_per_deci_literLazy != nil {
		return *a.nanograms_per_deci_literLazy
	}
	nanograms_per_deci_liter := a.convertFromBase(DensityNanogramPerDeciliter)
	a.nanograms_per_deci_literLazy = &nanograms_per_deci_liter
	return nanograms_per_deci_liter
}

// MicrogramPerDeciliter returns the value in MicrogramPerDeciliter.
func (a *Density) MicrogramsPerDeciLiter() float64 {
	if a.micrograms_per_deci_literLazy != nil {
		return *a.micrograms_per_deci_literLazy
	}
	micrograms_per_deci_liter := a.convertFromBase(DensityMicrogramPerDeciliter)
	a.micrograms_per_deci_literLazy = &micrograms_per_deci_liter
	return micrograms_per_deci_liter
}

// MilligramPerDeciliter returns the value in MilligramPerDeciliter.
func (a *Density) MilligramsPerDeciLiter() float64 {
	if a.milligrams_per_deci_literLazy != nil {
		return *a.milligrams_per_deci_literLazy
	}
	milligrams_per_deci_liter := a.convertFromBase(DensityMilligramPerDeciliter)
	a.milligrams_per_deci_literLazy = &milligrams_per_deci_liter
	return milligrams_per_deci_liter
}

// CentigramPerDeciliter returns the value in CentigramPerDeciliter.
func (a *Density) CentigramsPerDeciLiter() float64 {
	if a.centigrams_per_deci_literLazy != nil {
		return *a.centigrams_per_deci_literLazy
	}
	centigrams_per_deci_liter := a.convertFromBase(DensityCentigramPerDeciliter)
	a.centigrams_per_deci_literLazy = &centigrams_per_deci_liter
	return centigrams_per_deci_liter
}

// DecigramPerDeciliter returns the value in DecigramPerDeciliter.
func (a *Density) DecigramsPerDeciLiter() float64 {
	if a.decigrams_per_deci_literLazy != nil {
		return *a.decigrams_per_deci_literLazy
	}
	decigrams_per_deci_liter := a.convertFromBase(DensityDecigramPerDeciliter)
	a.decigrams_per_deci_literLazy = &decigrams_per_deci_liter
	return decigrams_per_deci_liter
}

// FemtogramPerMilliliter returns the value in FemtogramPerMilliliter.
func (a *Density) FemtogramsPerMilliliter() float64 {
	if a.femtograms_per_milliliterLazy != nil {
		return *a.femtograms_per_milliliterLazy
	}
	femtograms_per_milliliter := a.convertFromBase(DensityFemtogramPerMilliliter)
	a.femtograms_per_milliliterLazy = &femtograms_per_milliliter
	return femtograms_per_milliliter
}

// PicogramPerMilliliter returns the value in PicogramPerMilliliter.
func (a *Density) PicogramsPerMilliliter() float64 {
	if a.picograms_per_milliliterLazy != nil {
		return *a.picograms_per_milliliterLazy
	}
	picograms_per_milliliter := a.convertFromBase(DensityPicogramPerMilliliter)
	a.picograms_per_milliliterLazy = &picograms_per_milliliter
	return picograms_per_milliliter
}

// NanogramPerMilliliter returns the value in NanogramPerMilliliter.
func (a *Density) NanogramsPerMilliliter() float64 {
	if a.nanograms_per_milliliterLazy != nil {
		return *a.nanograms_per_milliliterLazy
	}
	nanograms_per_milliliter := a.convertFromBase(DensityNanogramPerMilliliter)
	a.nanograms_per_milliliterLazy = &nanograms_per_milliliter
	return nanograms_per_milliliter
}

// MicrogramPerMilliliter returns the value in MicrogramPerMilliliter.
func (a *Density) MicrogramsPerMilliliter() float64 {
	if a.micrograms_per_milliliterLazy != nil {
		return *a.micrograms_per_milliliterLazy
	}
	micrograms_per_milliliter := a.convertFromBase(DensityMicrogramPerMilliliter)
	a.micrograms_per_milliliterLazy = &micrograms_per_milliliter
	return micrograms_per_milliliter
}

// MilligramPerMilliliter returns the value in MilligramPerMilliliter.
func (a *Density) MilligramsPerMilliliter() float64 {
	if a.milligrams_per_milliliterLazy != nil {
		return *a.milligrams_per_milliliterLazy
	}
	milligrams_per_milliliter := a.convertFromBase(DensityMilligramPerMilliliter)
	a.milligrams_per_milliliterLazy = &milligrams_per_milliliter
	return milligrams_per_milliliter
}

// CentigramPerMilliliter returns the value in CentigramPerMilliliter.
func (a *Density) CentigramsPerMilliliter() float64 {
	if a.centigrams_per_milliliterLazy != nil {
		return *a.centigrams_per_milliliterLazy
	}
	centigrams_per_milliliter := a.convertFromBase(DensityCentigramPerMilliliter)
	a.centigrams_per_milliliterLazy = &centigrams_per_milliliter
	return centigrams_per_milliliter
}

// DecigramPerMilliliter returns the value in DecigramPerMilliliter.
func (a *Density) DecigramsPerMilliliter() float64 {
	if a.decigrams_per_milliliterLazy != nil {
		return *a.decigrams_per_milliliterLazy
	}
	decigrams_per_milliliter := a.convertFromBase(DensityDecigramPerMilliliter)
	a.decigrams_per_milliliterLazy = &decigrams_per_milliliter
	return decigrams_per_milliliter
}


// ToDto creates an DensityDto representation.
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

// ToDtoJSON creates an DensityDto representation.
func (a *Density) ToDtoJSON(holdInUnit *DensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Density to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Density) String() string {
	return a.ToString(DensityKilogramPerCubicMeter, 2)
}

// ToString formats the Density to string.
// fractionalDigits -1 for not mention
func (a *Density) ToString(unit DensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Density) getUnitAbbreviation(unit DensityUnits) string {
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

// Check if the given Density are equals to the current Density
func (a *Density) Equals(other *Density) bool {
	return a.value == other.BaseValue()
}

// Check if the given Density are equals to the current Density
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

// Add the given Density to the current Density.
func (a *Density) Add(other *Density) *Density {
	return &Density{value: a.value + other.BaseValue()}
}

// Subtract the given Density to the current Density.
func (a *Density) Subtract(other *Density) *Density {
	return &Density{value: a.value - other.BaseValue()}
}

// Multiply the given Density to the current Density.
func (a *Density) Multiply(other *Density) *Density {
	return &Density{value: a.value * other.BaseValue()}
}

// Divide the given Density to the current Density.
func (a *Density) Divide(other *Density) *Density {
	return &Density{value: a.value / other.BaseValue()}
}