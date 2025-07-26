// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeConcentrationUnits defines various units of VolumeConcentration.
type VolumeConcentrationUnits string

const (
    
        // 
        VolumeConcentrationDecimalFraction VolumeConcentrationUnits = "DecimalFraction"
        // 
        VolumeConcentrationLiterPerLiter VolumeConcentrationUnits = "LiterPerLiter"
        // 
        VolumeConcentrationLiterPerMilliliter VolumeConcentrationUnits = "LiterPerMilliliter"
        // 
        VolumeConcentrationPercent VolumeConcentrationUnits = "Percent"
        // 
        VolumeConcentrationPartPerThousand VolumeConcentrationUnits = "PartPerThousand"
        // 
        VolumeConcentrationPartPerMillion VolumeConcentrationUnits = "PartPerMillion"
        // 
        VolumeConcentrationPartPerBillion VolumeConcentrationUnits = "PartPerBillion"
        // 
        VolumeConcentrationPartPerTrillion VolumeConcentrationUnits = "PartPerTrillion"
        // 
        VolumeConcentrationPicoliterPerLiter VolumeConcentrationUnits = "PicoliterPerLiter"
        // 
        VolumeConcentrationNanoliterPerLiter VolumeConcentrationUnits = "NanoliterPerLiter"
        // 
        VolumeConcentrationMicroliterPerLiter VolumeConcentrationUnits = "MicroliterPerLiter"
        // 
        VolumeConcentrationMilliliterPerLiter VolumeConcentrationUnits = "MilliliterPerLiter"
        // 
        VolumeConcentrationCentiliterPerLiter VolumeConcentrationUnits = "CentiliterPerLiter"
        // 
        VolumeConcentrationDeciliterPerLiter VolumeConcentrationUnits = "DeciliterPerLiter"
        // 
        VolumeConcentrationPicoliterPerMilliliter VolumeConcentrationUnits = "PicoliterPerMilliliter"
        // 
        VolumeConcentrationNanoliterPerMilliliter VolumeConcentrationUnits = "NanoliterPerMilliliter"
        // 
        VolumeConcentrationMicroliterPerMilliliter VolumeConcentrationUnits = "MicroliterPerMilliliter"
        // 
        VolumeConcentrationMilliliterPerMilliliter VolumeConcentrationUnits = "MilliliterPerMilliliter"
        // 
        VolumeConcentrationCentiliterPerMilliliter VolumeConcentrationUnits = "CentiliterPerMilliliter"
        // 
        VolumeConcentrationDeciliterPerMilliliter VolumeConcentrationUnits = "DeciliterPerMilliliter"
)

var internalVolumeConcentrationUnitsMap = map[VolumeConcentrationUnits]bool{
	
	VolumeConcentrationDecimalFraction: true,
	VolumeConcentrationLiterPerLiter: true,
	VolumeConcentrationLiterPerMilliliter: true,
	VolumeConcentrationPercent: true,
	VolumeConcentrationPartPerThousand: true,
	VolumeConcentrationPartPerMillion: true,
	VolumeConcentrationPartPerBillion: true,
	VolumeConcentrationPartPerTrillion: true,
	VolumeConcentrationPicoliterPerLiter: true,
	VolumeConcentrationNanoliterPerLiter: true,
	VolumeConcentrationMicroliterPerLiter: true,
	VolumeConcentrationMilliliterPerLiter: true,
	VolumeConcentrationCentiliterPerLiter: true,
	VolumeConcentrationDeciliterPerLiter: true,
	VolumeConcentrationPicoliterPerMilliliter: true,
	VolumeConcentrationNanoliterPerMilliliter: true,
	VolumeConcentrationMicroliterPerMilliliter: true,
	VolumeConcentrationMilliliterPerMilliliter: true,
	VolumeConcentrationCentiliterPerMilliliter: true,
	VolumeConcentrationDeciliterPerMilliliter: true,
}

// VolumeConcentrationDto represents a VolumeConcentration measurement with a numerical value and its corresponding unit.
type VolumeConcentrationDto struct {
    // Value is the numerical representation of the VolumeConcentration.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the VolumeConcentration, as defined in the VolumeConcentrationUnits enumeration.
	Unit  VolumeConcentrationUnits `json:"unit" validate:"required,oneof=DecimalFraction LiterPerLiter LiterPerMilliliter Percent PartPerThousand PartPerMillion PartPerBillion PartPerTrillion PicoliterPerLiter NanoliterPerLiter MicroliterPerLiter MilliliterPerLiter CentiliterPerLiter DeciliterPerLiter PicoliterPerMilliliter NanoliterPerMilliliter MicroliterPerMilliliter MilliliterPerMilliliter CentiliterPerMilliliter DeciliterPerMilliliter"`
}

// VolumeConcentrationDtoFactory groups methods for creating and serializing VolumeConcentrationDto objects.
type VolumeConcentrationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumeConcentrationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumeConcentrationDtoFactory) FromJSON(data []byte) (*VolumeConcentrationDto, error) {
	a := VolumeConcentrationDto{}

    // Parse JSON into VolumeConcentrationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a VolumeConcentrationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumeConcentrationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// VolumeConcentration represents a measurement in a of VolumeConcentration.
//
// The volume concentration (not to be confused with volume fraction) is defined as the volume of a constituent divided by the total volume of the mixture.
type VolumeConcentration struct {
	// value is the base measurement stored internally.
	value       float64
    
    decimal_fractionsLazy *float64 
    liters_per_literLazy *float64 
    liters_per_milliliterLazy *float64 
    percentLazy *float64 
    parts_per_thousandLazy *float64 
    parts_per_millionLazy *float64 
    parts_per_billionLazy *float64 
    parts_per_trillionLazy *float64 
    picoliters_per_literLazy *float64 
    nanoliters_per_literLazy *float64 
    microliters_per_literLazy *float64 
    milliliters_per_literLazy *float64 
    centiliters_per_literLazy *float64 
    deciliters_per_literLazy *float64 
    picoliters_per_milliliterLazy *float64 
    nanoliters_per_milliliterLazy *float64 
    microliters_per_milliliterLazy *float64 
    milliliters_per_milliliterLazy *float64 
    centiliters_per_milliliterLazy *float64 
    deciliters_per_milliliterLazy *float64 
}

// VolumeConcentrationFactory groups methods for creating VolumeConcentration instances.
type VolumeConcentrationFactory struct{}

// CreateVolumeConcentration creates a new VolumeConcentration instance from the given value and unit.
func (uf VolumeConcentrationFactory) CreateVolumeConcentration(value float64, unit VolumeConcentrationUnits) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, unit)
}

// FromDto converts a VolumeConcentrationDto to a VolumeConcentration instance.
func (uf VolumeConcentrationFactory) FromDto(dto VolumeConcentrationDto) (*VolumeConcentration, error) {
	return newVolumeConcentration(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VolumeConcentration instance.
func (uf VolumeConcentrationFactory) FromDtoJSON(data []byte) (*VolumeConcentration, error) {
	unitDto, err := VolumeConcentrationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumeConcentrationDto from JSON: %w", err)
	}
	return VolumeConcentrationFactory{}.FromDto(*unitDto)
}


// FromDecimalFractions creates a new VolumeConcentration instance from a value in DecimalFractions.
func (uf VolumeConcentrationFactory) FromDecimalFractions(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDecimalFraction)
}

// FromLitersPerLiter creates a new VolumeConcentration instance from a value in LitersPerLiter.
func (uf VolumeConcentrationFactory) FromLitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationLiterPerLiter)
}

// FromLitersPerMilliliter creates a new VolumeConcentration instance from a value in LitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromLitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationLiterPerMilliliter)
}

// FromPercent creates a new VolumeConcentration instance from a value in Percent.
func (uf VolumeConcentrationFactory) FromPercent(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPercent)
}

// FromPartsPerThousand creates a new VolumeConcentration instance from a value in PartsPerThousand.
func (uf VolumeConcentrationFactory) FromPartsPerThousand(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerThousand)
}

// FromPartsPerMillion creates a new VolumeConcentration instance from a value in PartsPerMillion.
func (uf VolumeConcentrationFactory) FromPartsPerMillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerMillion)
}

// FromPartsPerBillion creates a new VolumeConcentration instance from a value in PartsPerBillion.
func (uf VolumeConcentrationFactory) FromPartsPerBillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerBillion)
}

// FromPartsPerTrillion creates a new VolumeConcentration instance from a value in PartsPerTrillion.
func (uf VolumeConcentrationFactory) FromPartsPerTrillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerTrillion)
}

// FromPicolitersPerLiter creates a new VolumeConcentration instance from a value in PicolitersPerLiter.
func (uf VolumeConcentrationFactory) FromPicolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPicoliterPerLiter)
}

// FromNanolitersPerLiter creates a new VolumeConcentration instance from a value in NanolitersPerLiter.
func (uf VolumeConcentrationFactory) FromNanolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationNanoliterPerLiter)
}

// FromMicrolitersPerLiter creates a new VolumeConcentration instance from a value in MicrolitersPerLiter.
func (uf VolumeConcentrationFactory) FromMicrolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMicroliterPerLiter)
}

// FromMillilitersPerLiter creates a new VolumeConcentration instance from a value in MillilitersPerLiter.
func (uf VolumeConcentrationFactory) FromMillilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMilliliterPerLiter)
}

// FromCentilitersPerLiter creates a new VolumeConcentration instance from a value in CentilitersPerLiter.
func (uf VolumeConcentrationFactory) FromCentilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationCentiliterPerLiter)
}

// FromDecilitersPerLiter creates a new VolumeConcentration instance from a value in DecilitersPerLiter.
func (uf VolumeConcentrationFactory) FromDecilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDeciliterPerLiter)
}

// FromPicolitersPerMilliliter creates a new VolumeConcentration instance from a value in PicolitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromPicolitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPicoliterPerMilliliter)
}

// FromNanolitersPerMilliliter creates a new VolumeConcentration instance from a value in NanolitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromNanolitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationNanoliterPerMilliliter)
}

// FromMicrolitersPerMilliliter creates a new VolumeConcentration instance from a value in MicrolitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromMicrolitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMicroliterPerMilliliter)
}

// FromMillilitersPerMilliliter creates a new VolumeConcentration instance from a value in MillilitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromMillilitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMilliliterPerMilliliter)
}

// FromCentilitersPerMilliliter creates a new VolumeConcentration instance from a value in CentilitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromCentilitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationCentiliterPerMilliliter)
}

// FromDecilitersPerMilliliter creates a new VolumeConcentration instance from a value in DecilitersPerMilliliter.
func (uf VolumeConcentrationFactory) FromDecilitersPerMilliliter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDeciliterPerMilliliter)
}


// newVolumeConcentration creates a new VolumeConcentration.
func newVolumeConcentration(value float64, fromUnit VolumeConcentrationUnits) (*VolumeConcentration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalVolumeConcentrationUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in VolumeConcentrationUnits", fromUnit)
	}
	a := &VolumeConcentration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumeConcentration in DecimalFraction unit (the base/default quantity).
func (a *VolumeConcentration) BaseValue() float64 {
	return a.value
}


// DecimalFractions returns the VolumeConcentration value in DecimalFractions.
//
// 
func (a *VolumeConcentration) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(VolumeConcentrationDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// LitersPerLiter returns the VolumeConcentration value in LitersPerLiter.
//
// 
func (a *VolumeConcentration) LitersPerLiter() float64 {
	if a.liters_per_literLazy != nil {
		return *a.liters_per_literLazy
	}
	liters_per_liter := a.convertFromBase(VolumeConcentrationLiterPerLiter)
	a.liters_per_literLazy = &liters_per_liter
	return liters_per_liter
}

// LitersPerMilliliter returns the VolumeConcentration value in LitersPerMilliliter.
//
// 
func (a *VolumeConcentration) LitersPerMilliliter() float64 {
	if a.liters_per_milliliterLazy != nil {
		return *a.liters_per_milliliterLazy
	}
	liters_per_milliliter := a.convertFromBase(VolumeConcentrationLiterPerMilliliter)
	a.liters_per_milliliterLazy = &liters_per_milliliter
	return liters_per_milliliter
}

// Percent returns the VolumeConcentration value in Percent.
//
// 
func (a *VolumeConcentration) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(VolumeConcentrationPercent)
	a.percentLazy = &percent
	return percent
}

// PartsPerThousand returns the VolumeConcentration value in PartsPerThousand.
//
// 
func (a *VolumeConcentration) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(VolumeConcentrationPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartsPerMillion returns the VolumeConcentration value in PartsPerMillion.
//
// 
func (a *VolumeConcentration) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(VolumeConcentrationPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartsPerBillion returns the VolumeConcentration value in PartsPerBillion.
//
// 
func (a *VolumeConcentration) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(VolumeConcentrationPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartsPerTrillion returns the VolumeConcentration value in PartsPerTrillion.
//
// 
func (a *VolumeConcentration) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(VolumeConcentrationPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}

// PicolitersPerLiter returns the VolumeConcentration value in PicolitersPerLiter.
//
// 
func (a *VolumeConcentration) PicolitersPerLiter() float64 {
	if a.picoliters_per_literLazy != nil {
		return *a.picoliters_per_literLazy
	}
	picoliters_per_liter := a.convertFromBase(VolumeConcentrationPicoliterPerLiter)
	a.picoliters_per_literLazy = &picoliters_per_liter
	return picoliters_per_liter
}

// NanolitersPerLiter returns the VolumeConcentration value in NanolitersPerLiter.
//
// 
func (a *VolumeConcentration) NanolitersPerLiter() float64 {
	if a.nanoliters_per_literLazy != nil {
		return *a.nanoliters_per_literLazy
	}
	nanoliters_per_liter := a.convertFromBase(VolumeConcentrationNanoliterPerLiter)
	a.nanoliters_per_literLazy = &nanoliters_per_liter
	return nanoliters_per_liter
}

// MicrolitersPerLiter returns the VolumeConcentration value in MicrolitersPerLiter.
//
// 
func (a *VolumeConcentration) MicrolitersPerLiter() float64 {
	if a.microliters_per_literLazy != nil {
		return *a.microliters_per_literLazy
	}
	microliters_per_liter := a.convertFromBase(VolumeConcentrationMicroliterPerLiter)
	a.microliters_per_literLazy = &microliters_per_liter
	return microliters_per_liter
}

// MillilitersPerLiter returns the VolumeConcentration value in MillilitersPerLiter.
//
// 
func (a *VolumeConcentration) MillilitersPerLiter() float64 {
	if a.milliliters_per_literLazy != nil {
		return *a.milliliters_per_literLazy
	}
	milliliters_per_liter := a.convertFromBase(VolumeConcentrationMilliliterPerLiter)
	a.milliliters_per_literLazy = &milliliters_per_liter
	return milliliters_per_liter
}

// CentilitersPerLiter returns the VolumeConcentration value in CentilitersPerLiter.
//
// 
func (a *VolumeConcentration) CentilitersPerLiter() float64 {
	if a.centiliters_per_literLazy != nil {
		return *a.centiliters_per_literLazy
	}
	centiliters_per_liter := a.convertFromBase(VolumeConcentrationCentiliterPerLiter)
	a.centiliters_per_literLazy = &centiliters_per_liter
	return centiliters_per_liter
}

// DecilitersPerLiter returns the VolumeConcentration value in DecilitersPerLiter.
//
// 
func (a *VolumeConcentration) DecilitersPerLiter() float64 {
	if a.deciliters_per_literLazy != nil {
		return *a.deciliters_per_literLazy
	}
	deciliters_per_liter := a.convertFromBase(VolumeConcentrationDeciliterPerLiter)
	a.deciliters_per_literLazy = &deciliters_per_liter
	return deciliters_per_liter
}

// PicolitersPerMilliliter returns the VolumeConcentration value in PicolitersPerMilliliter.
//
// 
func (a *VolumeConcentration) PicolitersPerMilliliter() float64 {
	if a.picoliters_per_milliliterLazy != nil {
		return *a.picoliters_per_milliliterLazy
	}
	picoliters_per_milliliter := a.convertFromBase(VolumeConcentrationPicoliterPerMilliliter)
	a.picoliters_per_milliliterLazy = &picoliters_per_milliliter
	return picoliters_per_milliliter
}

// NanolitersPerMilliliter returns the VolumeConcentration value in NanolitersPerMilliliter.
//
// 
func (a *VolumeConcentration) NanolitersPerMilliliter() float64 {
	if a.nanoliters_per_milliliterLazy != nil {
		return *a.nanoliters_per_milliliterLazy
	}
	nanoliters_per_milliliter := a.convertFromBase(VolumeConcentrationNanoliterPerMilliliter)
	a.nanoliters_per_milliliterLazy = &nanoliters_per_milliliter
	return nanoliters_per_milliliter
}

// MicrolitersPerMilliliter returns the VolumeConcentration value in MicrolitersPerMilliliter.
//
// 
func (a *VolumeConcentration) MicrolitersPerMilliliter() float64 {
	if a.microliters_per_milliliterLazy != nil {
		return *a.microliters_per_milliliterLazy
	}
	microliters_per_milliliter := a.convertFromBase(VolumeConcentrationMicroliterPerMilliliter)
	a.microliters_per_milliliterLazy = &microliters_per_milliliter
	return microliters_per_milliliter
}

// MillilitersPerMilliliter returns the VolumeConcentration value in MillilitersPerMilliliter.
//
// 
func (a *VolumeConcentration) MillilitersPerMilliliter() float64 {
	if a.milliliters_per_milliliterLazy != nil {
		return *a.milliliters_per_milliliterLazy
	}
	milliliters_per_milliliter := a.convertFromBase(VolumeConcentrationMilliliterPerMilliliter)
	a.milliliters_per_milliliterLazy = &milliliters_per_milliliter
	return milliliters_per_milliliter
}

// CentilitersPerMilliliter returns the VolumeConcentration value in CentilitersPerMilliliter.
//
// 
func (a *VolumeConcentration) CentilitersPerMilliliter() float64 {
	if a.centiliters_per_milliliterLazy != nil {
		return *a.centiliters_per_milliliterLazy
	}
	centiliters_per_milliliter := a.convertFromBase(VolumeConcentrationCentiliterPerMilliliter)
	a.centiliters_per_milliliterLazy = &centiliters_per_milliliter
	return centiliters_per_milliliter
}

// DecilitersPerMilliliter returns the VolumeConcentration value in DecilitersPerMilliliter.
//
// 
func (a *VolumeConcentration) DecilitersPerMilliliter() float64 {
	if a.deciliters_per_milliliterLazy != nil {
		return *a.deciliters_per_milliliterLazy
	}
	deciliters_per_milliliter := a.convertFromBase(VolumeConcentrationDeciliterPerMilliliter)
	a.deciliters_per_milliliterLazy = &deciliters_per_milliliter
	return deciliters_per_milliliter
}


// ToDto creates a VolumeConcentrationDto representation from the VolumeConcentration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *VolumeConcentration) ToDto(holdInUnit *VolumeConcentrationUnits) VolumeConcentrationDto {
	if holdInUnit == nil {
		defaultUnit := VolumeConcentrationDecimalFraction // Default value
		holdInUnit = &defaultUnit
	}

	return VolumeConcentrationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VolumeConcentration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecimalFraction by default.
func (a *VolumeConcentration) ToDtoJSON(holdInUnit *VolumeConcentrationUnits) ([]byte, error) {
	// Convert to VolumeConcentrationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VolumeConcentration to a specific unit value.
// The function uses the provided unit type (VolumeConcentrationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VolumeConcentration) Convert(toUnit VolumeConcentrationUnits) float64 {
	switch toUnit { 
    case VolumeConcentrationDecimalFraction:
		return a.DecimalFractions()
    case VolumeConcentrationLiterPerLiter:
		return a.LitersPerLiter()
    case VolumeConcentrationLiterPerMilliliter:
		return a.LitersPerMilliliter()
    case VolumeConcentrationPercent:
		return a.Percent()
    case VolumeConcentrationPartPerThousand:
		return a.PartsPerThousand()
    case VolumeConcentrationPartPerMillion:
		return a.PartsPerMillion()
    case VolumeConcentrationPartPerBillion:
		return a.PartsPerBillion()
    case VolumeConcentrationPartPerTrillion:
		return a.PartsPerTrillion()
    case VolumeConcentrationPicoliterPerLiter:
		return a.PicolitersPerLiter()
    case VolumeConcentrationNanoliterPerLiter:
		return a.NanolitersPerLiter()
    case VolumeConcentrationMicroliterPerLiter:
		return a.MicrolitersPerLiter()
    case VolumeConcentrationMilliliterPerLiter:
		return a.MillilitersPerLiter()
    case VolumeConcentrationCentiliterPerLiter:
		return a.CentilitersPerLiter()
    case VolumeConcentrationDeciliterPerLiter:
		return a.DecilitersPerLiter()
    case VolumeConcentrationPicoliterPerMilliliter:
		return a.PicolitersPerMilliliter()
    case VolumeConcentrationNanoliterPerMilliliter:
		return a.NanolitersPerMilliliter()
    case VolumeConcentrationMicroliterPerMilliliter:
		return a.MicrolitersPerMilliliter()
    case VolumeConcentrationMilliliterPerMilliliter:
		return a.MillilitersPerMilliliter()
    case VolumeConcentrationCentiliterPerMilliliter:
		return a.CentilitersPerMilliliter()
    case VolumeConcentrationDeciliterPerMilliliter:
		return a.DecilitersPerMilliliter()
	default:
		return math.NaN()
	}
}

func (a *VolumeConcentration) convertFromBase(toUnit VolumeConcentrationUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumeConcentrationDecimalFraction:
		return (value) 
	case VolumeConcentrationLiterPerLiter:
		return (value) 
	case VolumeConcentrationLiterPerMilliliter:
		return (value * 1e-3) 
	case VolumeConcentrationPercent:
		return (value * 1e2) 
	case VolumeConcentrationPartPerThousand:
		return (value * 1e3) 
	case VolumeConcentrationPartPerMillion:
		return (value * 1e6) 
	case VolumeConcentrationPartPerBillion:
		return (value * 1e9) 
	case VolumeConcentrationPartPerTrillion:
		return (value * 1e12) 
	case VolumeConcentrationPicoliterPerLiter:
		return ((value) / 1e-12) 
	case VolumeConcentrationNanoliterPerLiter:
		return ((value) / 1e-09) 
	case VolumeConcentrationMicroliterPerLiter:
		return ((value) / 1e-06) 
	case VolumeConcentrationMilliliterPerLiter:
		return ((value) / 0.001) 
	case VolumeConcentrationCentiliterPerLiter:
		return ((value) / 0.01) 
	case VolumeConcentrationDeciliterPerLiter:
		return ((value) / 0.1) 
	case VolumeConcentrationPicoliterPerMilliliter:
		return ((value * 1e-3) / 1e-12) 
	case VolumeConcentrationNanoliterPerMilliliter:
		return ((value * 1e-3) / 1e-09) 
	case VolumeConcentrationMicroliterPerMilliliter:
		return ((value * 1e-3) / 1e-06) 
	case VolumeConcentrationMilliliterPerMilliliter:
		return ((value * 1e-3) / 0.001) 
	case VolumeConcentrationCentiliterPerMilliliter:
		return ((value * 1e-3) / 0.01) 
	case VolumeConcentrationDeciliterPerMilliliter:
		return ((value * 1e-3) / 0.1) 
	default:
		return math.NaN()
	}
}

func (a *VolumeConcentration) convertToBase(value float64, fromUnit VolumeConcentrationUnits) float64 {
	switch fromUnit { 
	case VolumeConcentrationDecimalFraction:
		return (value) 
	case VolumeConcentrationLiterPerLiter:
		return (value) 
	case VolumeConcentrationLiterPerMilliliter:
		return (value / 1e-3) 
	case VolumeConcentrationPercent:
		return (value / 1e2) 
	case VolumeConcentrationPartPerThousand:
		return (value / 1e3) 
	case VolumeConcentrationPartPerMillion:
		return (value / 1e6) 
	case VolumeConcentrationPartPerBillion:
		return (value / 1e9) 
	case VolumeConcentrationPartPerTrillion:
		return (value / 1e12) 
	case VolumeConcentrationPicoliterPerLiter:
		return ((value) * 1e-12) 
	case VolumeConcentrationNanoliterPerLiter:
		return ((value) * 1e-09) 
	case VolumeConcentrationMicroliterPerLiter:
		return ((value) * 1e-06) 
	case VolumeConcentrationMilliliterPerLiter:
		return ((value) * 0.001) 
	case VolumeConcentrationCentiliterPerLiter:
		return ((value) * 0.01) 
	case VolumeConcentrationDeciliterPerLiter:
		return ((value) * 0.1) 
	case VolumeConcentrationPicoliterPerMilliliter:
		return ((value / 1e-3) * 1e-12) 
	case VolumeConcentrationNanoliterPerMilliliter:
		return ((value / 1e-3) * 1e-09) 
	case VolumeConcentrationMicroliterPerMilliliter:
		return ((value / 1e-3) * 1e-06) 
	case VolumeConcentrationMilliliterPerMilliliter:
		return ((value / 1e-3) * 0.001) 
	case VolumeConcentrationCentiliterPerMilliliter:
		return ((value / 1e-3) * 0.01) 
	case VolumeConcentrationDeciliterPerMilliliter:
		return ((value / 1e-3) * 0.1) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VolumeConcentration in the default unit (DecimalFraction),
// formatted to two decimal places.
func (a VolumeConcentration) String() string {
	return a.ToString(VolumeConcentrationDecimalFraction, 2)
}

// ToString formats the VolumeConcentration value as a string with the specified unit and fractional digits.
// It converts the VolumeConcentration to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VolumeConcentration value will be converted (e.g., DecimalFraction))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VolumeConcentration with the unit abbreviation.
func (a *VolumeConcentration) ToString(unit VolumeConcentrationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetVolumeConcentrationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetVolumeConcentrationAbbreviation(unit))
}

// Equals checks if the given VolumeConcentration is equal to the current VolumeConcentration.
//
// Parameters:
//    other: The VolumeConcentration to compare against.
//
// Returns:
//    bool: Returns true if both VolumeConcentration are equal, false otherwise.
func (a *VolumeConcentration) Equals(other *VolumeConcentration) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VolumeConcentration with another VolumeConcentration.
// It returns -1 if the current VolumeConcentration is less than the other VolumeConcentration, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VolumeConcentration to compare against.
//
// Returns:
//    int: -1 if the current VolumeConcentration is less, 1 if greater, and 0 if equal.
func (a *VolumeConcentration) CompareTo(other *VolumeConcentration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VolumeConcentration to the current VolumeConcentration and returns the result.
// The result is a new VolumeConcentration instance with the sum of the values.
//
// Parameters:
//    other: The VolumeConcentration to add to the current VolumeConcentration.
//
// Returns:
//    *VolumeConcentration: A new VolumeConcentration instance representing the sum of both VolumeConcentration.
func (a *VolumeConcentration) Add(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VolumeConcentration from the current VolumeConcentration and returns the result.
// The result is a new VolumeConcentration instance with the difference of the values.
//
// Parameters:
//    other: The VolumeConcentration to subtract from the current VolumeConcentration.
//
// Returns:
//    *VolumeConcentration: A new VolumeConcentration instance representing the difference of both VolumeConcentration.
func (a *VolumeConcentration) Subtract(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VolumeConcentration by the given VolumeConcentration and returns the result.
// The result is a new VolumeConcentration instance with the product of the values.
//
// Parameters:
//    other: The VolumeConcentration to multiply with the current VolumeConcentration.
//
// Returns:
//    *VolumeConcentration: A new VolumeConcentration instance representing the product of both VolumeConcentration.
func (a *VolumeConcentration) Multiply(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value * other.BaseValue()}
}

// Divide divides the current VolumeConcentration by the given VolumeConcentration and returns the result.
// The result is a new VolumeConcentration instance with the quotient of the values.
//
// Parameters:
//    other: The VolumeConcentration to divide the current VolumeConcentration by.
//
// Returns:
//    *VolumeConcentration: A new VolumeConcentration instance representing the quotient of both VolumeConcentration.
func (a *VolumeConcentration) Divide(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value / other.BaseValue()}
}

// GetVolumeConcentrationAbbreviation gets the unit abbreviation.
func GetVolumeConcentrationAbbreviation(unit VolumeConcentrationUnits) string {
	switch unit { 
	case VolumeConcentrationDecimalFraction:
		return "" 
	case VolumeConcentrationLiterPerLiter:
		return "l/l" 
	case VolumeConcentrationLiterPerMilliliter:
		return "l/ml" 
	case VolumeConcentrationPercent:
		return "%" 
	case VolumeConcentrationPartPerThousand:
		return "‰" 
	case VolumeConcentrationPartPerMillion:
		return "ppm" 
	case VolumeConcentrationPartPerBillion:
		return "ppb" 
	case VolumeConcentrationPartPerTrillion:
		return "ppt" 
	case VolumeConcentrationPicoliterPerLiter:
		return "pl/l" 
	case VolumeConcentrationNanoliterPerLiter:
		return "nl/l" 
	case VolumeConcentrationMicroliterPerLiter:
		return "μl/l" 
	case VolumeConcentrationMilliliterPerLiter:
		return "ml/l" 
	case VolumeConcentrationCentiliterPerLiter:
		return "cl/l" 
	case VolumeConcentrationDeciliterPerLiter:
		return "dl/l" 
	case VolumeConcentrationPicoliterPerMilliliter:
		return "pl/ml" 
	case VolumeConcentrationNanoliterPerMilliliter:
		return "nl/ml" 
	case VolumeConcentrationMicroliterPerMilliliter:
		return "μl/ml" 
	case VolumeConcentrationMilliliterPerMilliliter:
		return "ml/ml" 
	case VolumeConcentrationCentiliterPerMilliliter:
		return "cl/ml" 
	case VolumeConcentrationDeciliterPerMilliliter:
		return "dl/ml" 
	default:
		return ""
	}
}