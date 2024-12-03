package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassConcentrationUnits enumeration
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

// MassConcentrationDto represents an MassConcentration
type MassConcentrationDto struct {
	Value float64
	Unit  MassConcentrationUnits
}

// MassConcentrationDtoFactory struct to group related functions
type MassConcentrationDtoFactory struct{}

func (udf MassConcentrationDtoFactory) FromJSON(data []byte) (*MassConcentrationDto, error) {
	a := MassConcentrationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassConcentrationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MassConcentration struct
type MassConcentration struct {
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

// MassConcentrationFactory struct to group related functions
type MassConcentrationFactory struct{}

func (uf MassConcentrationFactory) CreateMassConcentration(value float64, unit MassConcentrationUnits) (*MassConcentration, error) {
	return newMassConcentration(value, unit)
}

func (uf MassConcentrationFactory) FromDto(dto MassConcentrationDto) (*MassConcentration, error) {
	return newMassConcentration(dto.Value, dto.Unit)
}

func (uf MassConcentrationFactory) FromDtoJSON(data []byte) (*MassConcentration, error) {
	unitDto, err := MassConcentrationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassConcentrationFactory{}.FromDto(*unitDto)
}


// FromGramPerCubicMillimeter creates a new MassConcentration instance from GramPerCubicMillimeter.
func (uf MassConcentrationFactory) FromGramsPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicMillimeter)
}

// FromGramPerCubicCentimeter creates a new MassConcentration instance from GramPerCubicCentimeter.
func (uf MassConcentrationFactory) FromGramsPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicCentimeter)
}

// FromGramPerCubicMeter creates a new MassConcentration instance from GramPerCubicMeter.
func (uf MassConcentrationFactory) FromGramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerCubicMeter)
}

// FromGramPerMicroliter creates a new MassConcentration instance from GramPerMicroliter.
func (uf MassConcentrationFactory) FromGramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerMicroliter)
}

// FromGramPerMilliliter creates a new MassConcentration instance from GramPerMilliliter.
func (uf MassConcentrationFactory) FromGramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerMilliliter)
}

// FromGramPerDeciliter creates a new MassConcentration instance from GramPerDeciliter.
func (uf MassConcentrationFactory) FromGramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerDeciliter)
}

// FromGramPerLiter creates a new MassConcentration instance from GramPerLiter.
func (uf MassConcentrationFactory) FromGramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationGramPerLiter)
}

// FromTonnePerCubicMillimeter creates a new MassConcentration instance from TonnePerCubicMillimeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicMillimeter)
}

// FromTonnePerCubicCentimeter creates a new MassConcentration instance from TonnePerCubicCentimeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicCentimeter)
}

// FromTonnePerCubicMeter creates a new MassConcentration instance from TonnePerCubicMeter.
func (uf MassConcentrationFactory) FromTonnesPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationTonnePerCubicMeter)
}

// FromPoundPerCubicInch creates a new MassConcentration instance from PoundPerCubicInch.
func (uf MassConcentrationFactory) FromPoundsPerCubicInch(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerCubicInch)
}

// FromPoundPerCubicFoot creates a new MassConcentration instance from PoundPerCubicFoot.
func (uf MassConcentrationFactory) FromPoundsPerCubicFoot(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerCubicFoot)
}

// FromSlugPerCubicFoot creates a new MassConcentration instance from SlugPerCubicFoot.
func (uf MassConcentrationFactory) FromSlugsPerCubicFoot(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationSlugPerCubicFoot)
}

// FromPoundPerUSGallon creates a new MassConcentration instance from PoundPerUSGallon.
func (uf MassConcentrationFactory) FromPoundsPerUSGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerUSGallon)
}

// FromOuncePerUSGallon creates a new MassConcentration instance from OuncePerUSGallon.
func (uf MassConcentrationFactory) FromOuncesPerUSGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationOuncePerUSGallon)
}

// FromOuncePerImperialGallon creates a new MassConcentration instance from OuncePerImperialGallon.
func (uf MassConcentrationFactory) FromOuncesPerImperialGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationOuncePerImperialGallon)
}

// FromPoundPerImperialGallon creates a new MassConcentration instance from PoundPerImperialGallon.
func (uf MassConcentrationFactory) FromPoundsPerImperialGallon(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPoundPerImperialGallon)
}

// FromKilogramPerCubicMillimeter creates a new MassConcentration instance from KilogramPerCubicMillimeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicMillimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicMillimeter)
}

// FromKilogramPerCubicCentimeter creates a new MassConcentration instance from KilogramPerCubicCentimeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicCentimeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicCentimeter)
}

// FromKilogramPerCubicMeter creates a new MassConcentration instance from KilogramPerCubicMeter.
func (uf MassConcentrationFactory) FromKilogramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerCubicMeter)
}

// FromMilligramPerCubicMeter creates a new MassConcentration instance from MilligramPerCubicMeter.
func (uf MassConcentrationFactory) FromMilligramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerCubicMeter)
}

// FromMicrogramPerCubicMeter creates a new MassConcentration instance from MicrogramPerCubicMeter.
func (uf MassConcentrationFactory) FromMicrogramsPerCubicMeter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerCubicMeter)
}

// FromPicogramPerMicroliter creates a new MassConcentration instance from PicogramPerMicroliter.
func (uf MassConcentrationFactory) FromPicogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerMicroliter)
}

// FromNanogramPerMicroliter creates a new MassConcentration instance from NanogramPerMicroliter.
func (uf MassConcentrationFactory) FromNanogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerMicroliter)
}

// FromMicrogramPerMicroliter creates a new MassConcentration instance from MicrogramPerMicroliter.
func (uf MassConcentrationFactory) FromMicrogramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerMicroliter)
}

// FromMilligramPerMicroliter creates a new MassConcentration instance from MilligramPerMicroliter.
func (uf MassConcentrationFactory) FromMilligramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerMicroliter)
}

// FromCentigramPerMicroliter creates a new MassConcentration instance from CentigramPerMicroliter.
func (uf MassConcentrationFactory) FromCentigramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerMicroliter)
}

// FromDecigramPerMicroliter creates a new MassConcentration instance from DecigramPerMicroliter.
func (uf MassConcentrationFactory) FromDecigramsPerMicroliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerMicroliter)
}

// FromPicogramPerMilliliter creates a new MassConcentration instance from PicogramPerMilliliter.
func (uf MassConcentrationFactory) FromPicogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerMilliliter)
}

// FromNanogramPerMilliliter creates a new MassConcentration instance from NanogramPerMilliliter.
func (uf MassConcentrationFactory) FromNanogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerMilliliter)
}

// FromMicrogramPerMilliliter creates a new MassConcentration instance from MicrogramPerMilliliter.
func (uf MassConcentrationFactory) FromMicrogramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerMilliliter)
}

// FromMilligramPerMilliliter creates a new MassConcentration instance from MilligramPerMilliliter.
func (uf MassConcentrationFactory) FromMilligramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerMilliliter)
}

// FromCentigramPerMilliliter creates a new MassConcentration instance from CentigramPerMilliliter.
func (uf MassConcentrationFactory) FromCentigramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerMilliliter)
}

// FromDecigramPerMilliliter creates a new MassConcentration instance from DecigramPerMilliliter.
func (uf MassConcentrationFactory) FromDecigramsPerMilliliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerMilliliter)
}

// FromPicogramPerDeciliter creates a new MassConcentration instance from PicogramPerDeciliter.
func (uf MassConcentrationFactory) FromPicogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerDeciliter)
}

// FromNanogramPerDeciliter creates a new MassConcentration instance from NanogramPerDeciliter.
func (uf MassConcentrationFactory) FromNanogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerDeciliter)
}

// FromMicrogramPerDeciliter creates a new MassConcentration instance from MicrogramPerDeciliter.
func (uf MassConcentrationFactory) FromMicrogramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerDeciliter)
}

// FromMilligramPerDeciliter creates a new MassConcentration instance from MilligramPerDeciliter.
func (uf MassConcentrationFactory) FromMilligramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerDeciliter)
}

// FromCentigramPerDeciliter creates a new MassConcentration instance from CentigramPerDeciliter.
func (uf MassConcentrationFactory) FromCentigramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerDeciliter)
}

// FromDecigramPerDeciliter creates a new MassConcentration instance from DecigramPerDeciliter.
func (uf MassConcentrationFactory) FromDecigramsPerDeciliter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerDeciliter)
}

// FromPicogramPerLiter creates a new MassConcentration instance from PicogramPerLiter.
func (uf MassConcentrationFactory) FromPicogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationPicogramPerLiter)
}

// FromNanogramPerLiter creates a new MassConcentration instance from NanogramPerLiter.
func (uf MassConcentrationFactory) FromNanogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationNanogramPerLiter)
}

// FromMicrogramPerLiter creates a new MassConcentration instance from MicrogramPerLiter.
func (uf MassConcentrationFactory) FromMicrogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMicrogramPerLiter)
}

// FromMilligramPerLiter creates a new MassConcentration instance from MilligramPerLiter.
func (uf MassConcentrationFactory) FromMilligramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationMilligramPerLiter)
}

// FromCentigramPerLiter creates a new MassConcentration instance from CentigramPerLiter.
func (uf MassConcentrationFactory) FromCentigramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationCentigramPerLiter)
}

// FromDecigramPerLiter creates a new MassConcentration instance from DecigramPerLiter.
func (uf MassConcentrationFactory) FromDecigramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationDecigramPerLiter)
}

// FromKilogramPerLiter creates a new MassConcentration instance from KilogramPerLiter.
func (uf MassConcentrationFactory) FromKilogramsPerLiter(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilogramPerLiter)
}

// FromKilopoundPerCubicInch creates a new MassConcentration instance from KilopoundPerCubicInch.
func (uf MassConcentrationFactory) FromKilopoundsPerCubicInch(value float64) (*MassConcentration, error) {
	return newMassConcentration(value, MassConcentrationKilopoundPerCubicInch)
}

// FromKilopoundPerCubicFoot creates a new MassConcentration instance from KilopoundPerCubicFoot.
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

// BaseValue returns the base value of MassConcentration in KilogramPerCubicMeter.
func (a *MassConcentration) BaseValue() float64 {
	return a.value
}


// GramPerCubicMillimeter returns the value in GramPerCubicMillimeter.
func (a *MassConcentration) GramsPerCubicMillimeter() float64 {
	if a.grams_per_cubic_millimeterLazy != nil {
		return *a.grams_per_cubic_millimeterLazy
	}
	grams_per_cubic_millimeter := a.convertFromBase(MassConcentrationGramPerCubicMillimeter)
	a.grams_per_cubic_millimeterLazy = &grams_per_cubic_millimeter
	return grams_per_cubic_millimeter
}

// GramPerCubicCentimeter returns the value in GramPerCubicCentimeter.
func (a *MassConcentration) GramsPerCubicCentimeter() float64 {
	if a.grams_per_cubic_centimeterLazy != nil {
		return *a.grams_per_cubic_centimeterLazy
	}
	grams_per_cubic_centimeter := a.convertFromBase(MassConcentrationGramPerCubicCentimeter)
	a.grams_per_cubic_centimeterLazy = &grams_per_cubic_centimeter
	return grams_per_cubic_centimeter
}

// GramPerCubicMeter returns the value in GramPerCubicMeter.
func (a *MassConcentration) GramsPerCubicMeter() float64 {
	if a.grams_per_cubic_meterLazy != nil {
		return *a.grams_per_cubic_meterLazy
	}
	grams_per_cubic_meter := a.convertFromBase(MassConcentrationGramPerCubicMeter)
	a.grams_per_cubic_meterLazy = &grams_per_cubic_meter
	return grams_per_cubic_meter
}

// GramPerMicroliter returns the value in GramPerMicroliter.
func (a *MassConcentration) GramsPerMicroliter() float64 {
	if a.grams_per_microliterLazy != nil {
		return *a.grams_per_microliterLazy
	}
	grams_per_microliter := a.convertFromBase(MassConcentrationGramPerMicroliter)
	a.grams_per_microliterLazy = &grams_per_microliter
	return grams_per_microliter
}

// GramPerMilliliter returns the value in GramPerMilliliter.
func (a *MassConcentration) GramsPerMilliliter() float64 {
	if a.grams_per_milliliterLazy != nil {
		return *a.grams_per_milliliterLazy
	}
	grams_per_milliliter := a.convertFromBase(MassConcentrationGramPerMilliliter)
	a.grams_per_milliliterLazy = &grams_per_milliliter
	return grams_per_milliliter
}

// GramPerDeciliter returns the value in GramPerDeciliter.
func (a *MassConcentration) GramsPerDeciliter() float64 {
	if a.grams_per_deciliterLazy != nil {
		return *a.grams_per_deciliterLazy
	}
	grams_per_deciliter := a.convertFromBase(MassConcentrationGramPerDeciliter)
	a.grams_per_deciliterLazy = &grams_per_deciliter
	return grams_per_deciliter
}

// GramPerLiter returns the value in GramPerLiter.
func (a *MassConcentration) GramsPerLiter() float64 {
	if a.grams_per_literLazy != nil {
		return *a.grams_per_literLazy
	}
	grams_per_liter := a.convertFromBase(MassConcentrationGramPerLiter)
	a.grams_per_literLazy = &grams_per_liter
	return grams_per_liter
}

// TonnePerCubicMillimeter returns the value in TonnePerCubicMillimeter.
func (a *MassConcentration) TonnesPerCubicMillimeter() float64 {
	if a.tonnes_per_cubic_millimeterLazy != nil {
		return *a.tonnes_per_cubic_millimeterLazy
	}
	tonnes_per_cubic_millimeter := a.convertFromBase(MassConcentrationTonnePerCubicMillimeter)
	a.tonnes_per_cubic_millimeterLazy = &tonnes_per_cubic_millimeter
	return tonnes_per_cubic_millimeter
}

// TonnePerCubicCentimeter returns the value in TonnePerCubicCentimeter.
func (a *MassConcentration) TonnesPerCubicCentimeter() float64 {
	if a.tonnes_per_cubic_centimeterLazy != nil {
		return *a.tonnes_per_cubic_centimeterLazy
	}
	tonnes_per_cubic_centimeter := a.convertFromBase(MassConcentrationTonnePerCubicCentimeter)
	a.tonnes_per_cubic_centimeterLazy = &tonnes_per_cubic_centimeter
	return tonnes_per_cubic_centimeter
}

// TonnePerCubicMeter returns the value in TonnePerCubicMeter.
func (a *MassConcentration) TonnesPerCubicMeter() float64 {
	if a.tonnes_per_cubic_meterLazy != nil {
		return *a.tonnes_per_cubic_meterLazy
	}
	tonnes_per_cubic_meter := a.convertFromBase(MassConcentrationTonnePerCubicMeter)
	a.tonnes_per_cubic_meterLazy = &tonnes_per_cubic_meter
	return tonnes_per_cubic_meter
}

// PoundPerCubicInch returns the value in PoundPerCubicInch.
func (a *MassConcentration) PoundsPerCubicInch() float64 {
	if a.pounds_per_cubic_inchLazy != nil {
		return *a.pounds_per_cubic_inchLazy
	}
	pounds_per_cubic_inch := a.convertFromBase(MassConcentrationPoundPerCubicInch)
	a.pounds_per_cubic_inchLazy = &pounds_per_cubic_inch
	return pounds_per_cubic_inch
}

// PoundPerCubicFoot returns the value in PoundPerCubicFoot.
func (a *MassConcentration) PoundsPerCubicFoot() float64 {
	if a.pounds_per_cubic_footLazy != nil {
		return *a.pounds_per_cubic_footLazy
	}
	pounds_per_cubic_foot := a.convertFromBase(MassConcentrationPoundPerCubicFoot)
	a.pounds_per_cubic_footLazy = &pounds_per_cubic_foot
	return pounds_per_cubic_foot
}

// SlugPerCubicFoot returns the value in SlugPerCubicFoot.
func (a *MassConcentration) SlugsPerCubicFoot() float64 {
	if a.slugs_per_cubic_footLazy != nil {
		return *a.slugs_per_cubic_footLazy
	}
	slugs_per_cubic_foot := a.convertFromBase(MassConcentrationSlugPerCubicFoot)
	a.slugs_per_cubic_footLazy = &slugs_per_cubic_foot
	return slugs_per_cubic_foot
}

// PoundPerUSGallon returns the value in PoundPerUSGallon.
func (a *MassConcentration) PoundsPerUSGallon() float64 {
	if a.pounds_per_us_gallonLazy != nil {
		return *a.pounds_per_us_gallonLazy
	}
	pounds_per_us_gallon := a.convertFromBase(MassConcentrationPoundPerUSGallon)
	a.pounds_per_us_gallonLazy = &pounds_per_us_gallon
	return pounds_per_us_gallon
}

// OuncePerUSGallon returns the value in OuncePerUSGallon.
func (a *MassConcentration) OuncesPerUSGallon() float64 {
	if a.ounces_per_us_gallonLazy != nil {
		return *a.ounces_per_us_gallonLazy
	}
	ounces_per_us_gallon := a.convertFromBase(MassConcentrationOuncePerUSGallon)
	a.ounces_per_us_gallonLazy = &ounces_per_us_gallon
	return ounces_per_us_gallon
}

// OuncePerImperialGallon returns the value in OuncePerImperialGallon.
func (a *MassConcentration) OuncesPerImperialGallon() float64 {
	if a.ounces_per_imperial_gallonLazy != nil {
		return *a.ounces_per_imperial_gallonLazy
	}
	ounces_per_imperial_gallon := a.convertFromBase(MassConcentrationOuncePerImperialGallon)
	a.ounces_per_imperial_gallonLazy = &ounces_per_imperial_gallon
	return ounces_per_imperial_gallon
}

// PoundPerImperialGallon returns the value in PoundPerImperialGallon.
func (a *MassConcentration) PoundsPerImperialGallon() float64 {
	if a.pounds_per_imperial_gallonLazy != nil {
		return *a.pounds_per_imperial_gallonLazy
	}
	pounds_per_imperial_gallon := a.convertFromBase(MassConcentrationPoundPerImperialGallon)
	a.pounds_per_imperial_gallonLazy = &pounds_per_imperial_gallon
	return pounds_per_imperial_gallon
}

// KilogramPerCubicMillimeter returns the value in KilogramPerCubicMillimeter.
func (a *MassConcentration) KilogramsPerCubicMillimeter() float64 {
	if a.kilograms_per_cubic_millimeterLazy != nil {
		return *a.kilograms_per_cubic_millimeterLazy
	}
	kilograms_per_cubic_millimeter := a.convertFromBase(MassConcentrationKilogramPerCubicMillimeter)
	a.kilograms_per_cubic_millimeterLazy = &kilograms_per_cubic_millimeter
	return kilograms_per_cubic_millimeter
}

// KilogramPerCubicCentimeter returns the value in KilogramPerCubicCentimeter.
func (a *MassConcentration) KilogramsPerCubicCentimeter() float64 {
	if a.kilograms_per_cubic_centimeterLazy != nil {
		return *a.kilograms_per_cubic_centimeterLazy
	}
	kilograms_per_cubic_centimeter := a.convertFromBase(MassConcentrationKilogramPerCubicCentimeter)
	a.kilograms_per_cubic_centimeterLazy = &kilograms_per_cubic_centimeter
	return kilograms_per_cubic_centimeter
}

// KilogramPerCubicMeter returns the value in KilogramPerCubicMeter.
func (a *MassConcentration) KilogramsPerCubicMeter() float64 {
	if a.kilograms_per_cubic_meterLazy != nil {
		return *a.kilograms_per_cubic_meterLazy
	}
	kilograms_per_cubic_meter := a.convertFromBase(MassConcentrationKilogramPerCubicMeter)
	a.kilograms_per_cubic_meterLazy = &kilograms_per_cubic_meter
	return kilograms_per_cubic_meter
}

// MilligramPerCubicMeter returns the value in MilligramPerCubicMeter.
func (a *MassConcentration) MilligramsPerCubicMeter() float64 {
	if a.milligrams_per_cubic_meterLazy != nil {
		return *a.milligrams_per_cubic_meterLazy
	}
	milligrams_per_cubic_meter := a.convertFromBase(MassConcentrationMilligramPerCubicMeter)
	a.milligrams_per_cubic_meterLazy = &milligrams_per_cubic_meter
	return milligrams_per_cubic_meter
}

// MicrogramPerCubicMeter returns the value in MicrogramPerCubicMeter.
func (a *MassConcentration) MicrogramsPerCubicMeter() float64 {
	if a.micrograms_per_cubic_meterLazy != nil {
		return *a.micrograms_per_cubic_meterLazy
	}
	micrograms_per_cubic_meter := a.convertFromBase(MassConcentrationMicrogramPerCubicMeter)
	a.micrograms_per_cubic_meterLazy = &micrograms_per_cubic_meter
	return micrograms_per_cubic_meter
}

// PicogramPerMicroliter returns the value in PicogramPerMicroliter.
func (a *MassConcentration) PicogramsPerMicroliter() float64 {
	if a.picograms_per_microliterLazy != nil {
		return *a.picograms_per_microliterLazy
	}
	picograms_per_microliter := a.convertFromBase(MassConcentrationPicogramPerMicroliter)
	a.picograms_per_microliterLazy = &picograms_per_microliter
	return picograms_per_microliter
}

// NanogramPerMicroliter returns the value in NanogramPerMicroliter.
func (a *MassConcentration) NanogramsPerMicroliter() float64 {
	if a.nanograms_per_microliterLazy != nil {
		return *a.nanograms_per_microliterLazy
	}
	nanograms_per_microliter := a.convertFromBase(MassConcentrationNanogramPerMicroliter)
	a.nanograms_per_microliterLazy = &nanograms_per_microliter
	return nanograms_per_microliter
}

// MicrogramPerMicroliter returns the value in MicrogramPerMicroliter.
func (a *MassConcentration) MicrogramsPerMicroliter() float64 {
	if a.micrograms_per_microliterLazy != nil {
		return *a.micrograms_per_microliterLazy
	}
	micrograms_per_microliter := a.convertFromBase(MassConcentrationMicrogramPerMicroliter)
	a.micrograms_per_microliterLazy = &micrograms_per_microliter
	return micrograms_per_microliter
}

// MilligramPerMicroliter returns the value in MilligramPerMicroliter.
func (a *MassConcentration) MilligramsPerMicroliter() float64 {
	if a.milligrams_per_microliterLazy != nil {
		return *a.milligrams_per_microliterLazy
	}
	milligrams_per_microliter := a.convertFromBase(MassConcentrationMilligramPerMicroliter)
	a.milligrams_per_microliterLazy = &milligrams_per_microliter
	return milligrams_per_microliter
}

// CentigramPerMicroliter returns the value in CentigramPerMicroliter.
func (a *MassConcentration) CentigramsPerMicroliter() float64 {
	if a.centigrams_per_microliterLazy != nil {
		return *a.centigrams_per_microliterLazy
	}
	centigrams_per_microliter := a.convertFromBase(MassConcentrationCentigramPerMicroliter)
	a.centigrams_per_microliterLazy = &centigrams_per_microliter
	return centigrams_per_microliter
}

// DecigramPerMicroliter returns the value in DecigramPerMicroliter.
func (a *MassConcentration) DecigramsPerMicroliter() float64 {
	if a.decigrams_per_microliterLazy != nil {
		return *a.decigrams_per_microliterLazy
	}
	decigrams_per_microliter := a.convertFromBase(MassConcentrationDecigramPerMicroliter)
	a.decigrams_per_microliterLazy = &decigrams_per_microliter
	return decigrams_per_microliter
}

// PicogramPerMilliliter returns the value in PicogramPerMilliliter.
func (a *MassConcentration) PicogramsPerMilliliter() float64 {
	if a.picograms_per_milliliterLazy != nil {
		return *a.picograms_per_milliliterLazy
	}
	picograms_per_milliliter := a.convertFromBase(MassConcentrationPicogramPerMilliliter)
	a.picograms_per_milliliterLazy = &picograms_per_milliliter
	return picograms_per_milliliter
}

// NanogramPerMilliliter returns the value in NanogramPerMilliliter.
func (a *MassConcentration) NanogramsPerMilliliter() float64 {
	if a.nanograms_per_milliliterLazy != nil {
		return *a.nanograms_per_milliliterLazy
	}
	nanograms_per_milliliter := a.convertFromBase(MassConcentrationNanogramPerMilliliter)
	a.nanograms_per_milliliterLazy = &nanograms_per_milliliter
	return nanograms_per_milliliter
}

// MicrogramPerMilliliter returns the value in MicrogramPerMilliliter.
func (a *MassConcentration) MicrogramsPerMilliliter() float64 {
	if a.micrograms_per_milliliterLazy != nil {
		return *a.micrograms_per_milliliterLazy
	}
	micrograms_per_milliliter := a.convertFromBase(MassConcentrationMicrogramPerMilliliter)
	a.micrograms_per_milliliterLazy = &micrograms_per_milliliter
	return micrograms_per_milliliter
}

// MilligramPerMilliliter returns the value in MilligramPerMilliliter.
func (a *MassConcentration) MilligramsPerMilliliter() float64 {
	if a.milligrams_per_milliliterLazy != nil {
		return *a.milligrams_per_milliliterLazy
	}
	milligrams_per_milliliter := a.convertFromBase(MassConcentrationMilligramPerMilliliter)
	a.milligrams_per_milliliterLazy = &milligrams_per_milliliter
	return milligrams_per_milliliter
}

// CentigramPerMilliliter returns the value in CentigramPerMilliliter.
func (a *MassConcentration) CentigramsPerMilliliter() float64 {
	if a.centigrams_per_milliliterLazy != nil {
		return *a.centigrams_per_milliliterLazy
	}
	centigrams_per_milliliter := a.convertFromBase(MassConcentrationCentigramPerMilliliter)
	a.centigrams_per_milliliterLazy = &centigrams_per_milliliter
	return centigrams_per_milliliter
}

// DecigramPerMilliliter returns the value in DecigramPerMilliliter.
func (a *MassConcentration) DecigramsPerMilliliter() float64 {
	if a.decigrams_per_milliliterLazy != nil {
		return *a.decigrams_per_milliliterLazy
	}
	decigrams_per_milliliter := a.convertFromBase(MassConcentrationDecigramPerMilliliter)
	a.decigrams_per_milliliterLazy = &decigrams_per_milliliter
	return decigrams_per_milliliter
}

// PicogramPerDeciliter returns the value in PicogramPerDeciliter.
func (a *MassConcentration) PicogramsPerDeciliter() float64 {
	if a.picograms_per_deciliterLazy != nil {
		return *a.picograms_per_deciliterLazy
	}
	picograms_per_deciliter := a.convertFromBase(MassConcentrationPicogramPerDeciliter)
	a.picograms_per_deciliterLazy = &picograms_per_deciliter
	return picograms_per_deciliter
}

// NanogramPerDeciliter returns the value in NanogramPerDeciliter.
func (a *MassConcentration) NanogramsPerDeciliter() float64 {
	if a.nanograms_per_deciliterLazy != nil {
		return *a.nanograms_per_deciliterLazy
	}
	nanograms_per_deciliter := a.convertFromBase(MassConcentrationNanogramPerDeciliter)
	a.nanograms_per_deciliterLazy = &nanograms_per_deciliter
	return nanograms_per_deciliter
}

// MicrogramPerDeciliter returns the value in MicrogramPerDeciliter.
func (a *MassConcentration) MicrogramsPerDeciliter() float64 {
	if a.micrograms_per_deciliterLazy != nil {
		return *a.micrograms_per_deciliterLazy
	}
	micrograms_per_deciliter := a.convertFromBase(MassConcentrationMicrogramPerDeciliter)
	a.micrograms_per_deciliterLazy = &micrograms_per_deciliter
	return micrograms_per_deciliter
}

// MilligramPerDeciliter returns the value in MilligramPerDeciliter.
func (a *MassConcentration) MilligramsPerDeciliter() float64 {
	if a.milligrams_per_deciliterLazy != nil {
		return *a.milligrams_per_deciliterLazy
	}
	milligrams_per_deciliter := a.convertFromBase(MassConcentrationMilligramPerDeciliter)
	a.milligrams_per_deciliterLazy = &milligrams_per_deciliter
	return milligrams_per_deciliter
}

// CentigramPerDeciliter returns the value in CentigramPerDeciliter.
func (a *MassConcentration) CentigramsPerDeciliter() float64 {
	if a.centigrams_per_deciliterLazy != nil {
		return *a.centigrams_per_deciliterLazy
	}
	centigrams_per_deciliter := a.convertFromBase(MassConcentrationCentigramPerDeciliter)
	a.centigrams_per_deciliterLazy = &centigrams_per_deciliter
	return centigrams_per_deciliter
}

// DecigramPerDeciliter returns the value in DecigramPerDeciliter.
func (a *MassConcentration) DecigramsPerDeciliter() float64 {
	if a.decigrams_per_deciliterLazy != nil {
		return *a.decigrams_per_deciliterLazy
	}
	decigrams_per_deciliter := a.convertFromBase(MassConcentrationDecigramPerDeciliter)
	a.decigrams_per_deciliterLazy = &decigrams_per_deciliter
	return decigrams_per_deciliter
}

// PicogramPerLiter returns the value in PicogramPerLiter.
func (a *MassConcentration) PicogramsPerLiter() float64 {
	if a.picograms_per_literLazy != nil {
		return *a.picograms_per_literLazy
	}
	picograms_per_liter := a.convertFromBase(MassConcentrationPicogramPerLiter)
	a.picograms_per_literLazy = &picograms_per_liter
	return picograms_per_liter
}

// NanogramPerLiter returns the value in NanogramPerLiter.
func (a *MassConcentration) NanogramsPerLiter() float64 {
	if a.nanograms_per_literLazy != nil {
		return *a.nanograms_per_literLazy
	}
	nanograms_per_liter := a.convertFromBase(MassConcentrationNanogramPerLiter)
	a.nanograms_per_literLazy = &nanograms_per_liter
	return nanograms_per_liter
}

// MicrogramPerLiter returns the value in MicrogramPerLiter.
func (a *MassConcentration) MicrogramsPerLiter() float64 {
	if a.micrograms_per_literLazy != nil {
		return *a.micrograms_per_literLazy
	}
	micrograms_per_liter := a.convertFromBase(MassConcentrationMicrogramPerLiter)
	a.micrograms_per_literLazy = &micrograms_per_liter
	return micrograms_per_liter
}

// MilligramPerLiter returns the value in MilligramPerLiter.
func (a *MassConcentration) MilligramsPerLiter() float64 {
	if a.milligrams_per_literLazy != nil {
		return *a.milligrams_per_literLazy
	}
	milligrams_per_liter := a.convertFromBase(MassConcentrationMilligramPerLiter)
	a.milligrams_per_literLazy = &milligrams_per_liter
	return milligrams_per_liter
}

// CentigramPerLiter returns the value in CentigramPerLiter.
func (a *MassConcentration) CentigramsPerLiter() float64 {
	if a.centigrams_per_literLazy != nil {
		return *a.centigrams_per_literLazy
	}
	centigrams_per_liter := a.convertFromBase(MassConcentrationCentigramPerLiter)
	a.centigrams_per_literLazy = &centigrams_per_liter
	return centigrams_per_liter
}

// DecigramPerLiter returns the value in DecigramPerLiter.
func (a *MassConcentration) DecigramsPerLiter() float64 {
	if a.decigrams_per_literLazy != nil {
		return *a.decigrams_per_literLazy
	}
	decigrams_per_liter := a.convertFromBase(MassConcentrationDecigramPerLiter)
	a.decigrams_per_literLazy = &decigrams_per_liter
	return decigrams_per_liter
}

// KilogramPerLiter returns the value in KilogramPerLiter.
func (a *MassConcentration) KilogramsPerLiter() float64 {
	if a.kilograms_per_literLazy != nil {
		return *a.kilograms_per_literLazy
	}
	kilograms_per_liter := a.convertFromBase(MassConcentrationKilogramPerLiter)
	a.kilograms_per_literLazy = &kilograms_per_liter
	return kilograms_per_liter
}

// KilopoundPerCubicInch returns the value in KilopoundPerCubicInch.
func (a *MassConcentration) KilopoundsPerCubicInch() float64 {
	if a.kilopounds_per_cubic_inchLazy != nil {
		return *a.kilopounds_per_cubic_inchLazy
	}
	kilopounds_per_cubic_inch := a.convertFromBase(MassConcentrationKilopoundPerCubicInch)
	a.kilopounds_per_cubic_inchLazy = &kilopounds_per_cubic_inch
	return kilopounds_per_cubic_inch
}

// KilopoundPerCubicFoot returns the value in KilopoundPerCubicFoot.
func (a *MassConcentration) KilopoundsPerCubicFoot() float64 {
	if a.kilopounds_per_cubic_footLazy != nil {
		return *a.kilopounds_per_cubic_footLazy
	}
	kilopounds_per_cubic_foot := a.convertFromBase(MassConcentrationKilopoundPerCubicFoot)
	a.kilopounds_per_cubic_footLazy = &kilopounds_per_cubic_foot
	return kilopounds_per_cubic_foot
}


// ToDto creates an MassConcentrationDto representation.
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

// ToDtoJSON creates an MassConcentrationDto representation.
func (a *MassConcentration) ToDtoJSON(holdInUnit *MassConcentrationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MassConcentration to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a MassConcentration) String() string {
	return a.ToString(MassConcentrationKilogramPerCubicMeter, 2)
}

// ToString formats the MassConcentration to string.
// fractionalDigits -1 for not mention
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

// Check if the given MassConcentration are equals to the current MassConcentration
func (a *MassConcentration) Equals(other *MassConcentration) bool {
	return a.value == other.BaseValue()
}

// Check if the given MassConcentration are equals to the current MassConcentration
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

// Add the given MassConcentration to the current MassConcentration.
func (a *MassConcentration) Add(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value + other.BaseValue()}
}

// Subtract the given MassConcentration to the current MassConcentration.
func (a *MassConcentration) Subtract(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value - other.BaseValue()}
}

// Multiply the given MassConcentration to the current MassConcentration.
func (a *MassConcentration) Multiply(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value * other.BaseValue()}
}

// Divide the given MassConcentration to the current MassConcentration.
func (a *MassConcentration) Divide(other *MassConcentration) *MassConcentration {
	return &MassConcentration{value: a.value / other.BaseValue()}
}