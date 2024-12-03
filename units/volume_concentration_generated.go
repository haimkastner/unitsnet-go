// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeConcentrationUnits enumeration
type VolumeConcentrationUnits string

const (
    
        // 
        VolumeConcentrationDecimalFraction VolumeConcentrationUnits = "DecimalFraction"
        // 
        VolumeConcentrationLitersPerLiter VolumeConcentrationUnits = "LitersPerLiter"
        // 
        VolumeConcentrationLitersPerMililiter VolumeConcentrationUnits = "LitersPerMililiter"
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
        VolumeConcentrationPicolitersPerLiter VolumeConcentrationUnits = "PicolitersPerLiter"
        // 
        VolumeConcentrationNanolitersPerLiter VolumeConcentrationUnits = "NanolitersPerLiter"
        // 
        VolumeConcentrationMicrolitersPerLiter VolumeConcentrationUnits = "MicrolitersPerLiter"
        // 
        VolumeConcentrationMillilitersPerLiter VolumeConcentrationUnits = "MillilitersPerLiter"
        // 
        VolumeConcentrationCentilitersPerLiter VolumeConcentrationUnits = "CentilitersPerLiter"
        // 
        VolumeConcentrationDecilitersPerLiter VolumeConcentrationUnits = "DecilitersPerLiter"
        // 
        VolumeConcentrationPicolitersPerMililiter VolumeConcentrationUnits = "PicolitersPerMililiter"
        // 
        VolumeConcentrationNanolitersPerMililiter VolumeConcentrationUnits = "NanolitersPerMililiter"
        // 
        VolumeConcentrationMicrolitersPerMililiter VolumeConcentrationUnits = "MicrolitersPerMililiter"
        // 
        VolumeConcentrationMillilitersPerMililiter VolumeConcentrationUnits = "MillilitersPerMililiter"
        // 
        VolumeConcentrationCentilitersPerMililiter VolumeConcentrationUnits = "CentilitersPerMililiter"
        // 
        VolumeConcentrationDecilitersPerMililiter VolumeConcentrationUnits = "DecilitersPerMililiter"
)

// VolumeConcentrationDto represents an VolumeConcentration
type VolumeConcentrationDto struct {
	Value float64
	Unit  VolumeConcentrationUnits
}

// VolumeConcentrationDtoFactory struct to group related functions
type VolumeConcentrationDtoFactory struct{}

func (udf VolumeConcentrationDtoFactory) FromJSON(data []byte) (*VolumeConcentrationDto, error) {
	a := VolumeConcentrationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumeConcentrationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VolumeConcentration struct
type VolumeConcentration struct {
	value       float64
    
    decimal_fractionsLazy *float64 
    liters_per_literLazy *float64 
    liters_per_mililiterLazy *float64 
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
    picoliters_per_mililiterLazy *float64 
    nanoliters_per_mililiterLazy *float64 
    microliters_per_mililiterLazy *float64 
    milliliters_per_mililiterLazy *float64 
    centiliters_per_mililiterLazy *float64 
    deciliters_per_mililiterLazy *float64 
}

// VolumeConcentrationFactory struct to group related functions
type VolumeConcentrationFactory struct{}

func (uf VolumeConcentrationFactory) CreateVolumeConcentration(value float64, unit VolumeConcentrationUnits) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, unit)
}

func (uf VolumeConcentrationFactory) FromDto(dto VolumeConcentrationDto) (*VolumeConcentration, error) {
	return newVolumeConcentration(dto.Value, dto.Unit)
}

func (uf VolumeConcentrationFactory) FromDtoJSON(data []byte) (*VolumeConcentration, error) {
	unitDto, err := VolumeConcentrationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumeConcentrationFactory{}.FromDto(*unitDto)
}


// FromDecimalFraction creates a new VolumeConcentration instance from DecimalFraction.
func (uf VolumeConcentrationFactory) FromDecimalFractions(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDecimalFraction)
}

// FromLitersPerLiter creates a new VolumeConcentration instance from LitersPerLiter.
func (uf VolumeConcentrationFactory) FromLitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationLitersPerLiter)
}

// FromLitersPerMililiter creates a new VolumeConcentration instance from LitersPerMililiter.
func (uf VolumeConcentrationFactory) FromLitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationLitersPerMililiter)
}

// FromPercent creates a new VolumeConcentration instance from Percent.
func (uf VolumeConcentrationFactory) FromPercent(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPercent)
}

// FromPartPerThousand creates a new VolumeConcentration instance from PartPerThousand.
func (uf VolumeConcentrationFactory) FromPartsPerThousand(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerThousand)
}

// FromPartPerMillion creates a new VolumeConcentration instance from PartPerMillion.
func (uf VolumeConcentrationFactory) FromPartsPerMillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerMillion)
}

// FromPartPerBillion creates a new VolumeConcentration instance from PartPerBillion.
func (uf VolumeConcentrationFactory) FromPartsPerBillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerBillion)
}

// FromPartPerTrillion creates a new VolumeConcentration instance from PartPerTrillion.
func (uf VolumeConcentrationFactory) FromPartsPerTrillion(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPartPerTrillion)
}

// FromPicolitersPerLiter creates a new VolumeConcentration instance from PicolitersPerLiter.
func (uf VolumeConcentrationFactory) FromPicolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPicolitersPerLiter)
}

// FromNanolitersPerLiter creates a new VolumeConcentration instance from NanolitersPerLiter.
func (uf VolumeConcentrationFactory) FromNanolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationNanolitersPerLiter)
}

// FromMicrolitersPerLiter creates a new VolumeConcentration instance from MicrolitersPerLiter.
func (uf VolumeConcentrationFactory) FromMicrolitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMicrolitersPerLiter)
}

// FromMillilitersPerLiter creates a new VolumeConcentration instance from MillilitersPerLiter.
func (uf VolumeConcentrationFactory) FromMillilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMillilitersPerLiter)
}

// FromCentilitersPerLiter creates a new VolumeConcentration instance from CentilitersPerLiter.
func (uf VolumeConcentrationFactory) FromCentilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationCentilitersPerLiter)
}

// FromDecilitersPerLiter creates a new VolumeConcentration instance from DecilitersPerLiter.
func (uf VolumeConcentrationFactory) FromDecilitersPerLiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDecilitersPerLiter)
}

// FromPicolitersPerMililiter creates a new VolumeConcentration instance from PicolitersPerMililiter.
func (uf VolumeConcentrationFactory) FromPicolitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationPicolitersPerMililiter)
}

// FromNanolitersPerMililiter creates a new VolumeConcentration instance from NanolitersPerMililiter.
func (uf VolumeConcentrationFactory) FromNanolitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationNanolitersPerMililiter)
}

// FromMicrolitersPerMililiter creates a new VolumeConcentration instance from MicrolitersPerMililiter.
func (uf VolumeConcentrationFactory) FromMicrolitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMicrolitersPerMililiter)
}

// FromMillilitersPerMililiter creates a new VolumeConcentration instance from MillilitersPerMililiter.
func (uf VolumeConcentrationFactory) FromMillilitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationMillilitersPerMililiter)
}

// FromCentilitersPerMililiter creates a new VolumeConcentration instance from CentilitersPerMililiter.
func (uf VolumeConcentrationFactory) FromCentilitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationCentilitersPerMililiter)
}

// FromDecilitersPerMililiter creates a new VolumeConcentration instance from DecilitersPerMililiter.
func (uf VolumeConcentrationFactory) FromDecilitersPerMililiter(value float64) (*VolumeConcentration, error) {
	return newVolumeConcentration(value, VolumeConcentrationDecilitersPerMililiter)
}




// newVolumeConcentration creates a new VolumeConcentration.
func newVolumeConcentration(value float64, fromUnit VolumeConcentrationUnits) (*VolumeConcentration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VolumeConcentration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumeConcentration in DecimalFraction.
func (a *VolumeConcentration) BaseValue() float64 {
	return a.value
}


// DecimalFraction returns the value in DecimalFraction.
func (a *VolumeConcentration) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(VolumeConcentrationDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// LitersPerLiter returns the value in LitersPerLiter.
func (a *VolumeConcentration) LitersPerLiter() float64 {
	if a.liters_per_literLazy != nil {
		return *a.liters_per_literLazy
	}
	liters_per_liter := a.convertFromBase(VolumeConcentrationLitersPerLiter)
	a.liters_per_literLazy = &liters_per_liter
	return liters_per_liter
}

// LitersPerMililiter returns the value in LitersPerMililiter.
func (a *VolumeConcentration) LitersPerMililiter() float64 {
	if a.liters_per_mililiterLazy != nil {
		return *a.liters_per_mililiterLazy
	}
	liters_per_mililiter := a.convertFromBase(VolumeConcentrationLitersPerMililiter)
	a.liters_per_mililiterLazy = &liters_per_mililiter
	return liters_per_mililiter
}

// Percent returns the value in Percent.
func (a *VolumeConcentration) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(VolumeConcentrationPercent)
	a.percentLazy = &percent
	return percent
}

// PartPerThousand returns the value in PartPerThousand.
func (a *VolumeConcentration) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(VolumeConcentrationPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartPerMillion returns the value in PartPerMillion.
func (a *VolumeConcentration) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(VolumeConcentrationPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartPerBillion returns the value in PartPerBillion.
func (a *VolumeConcentration) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(VolumeConcentrationPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartPerTrillion returns the value in PartPerTrillion.
func (a *VolumeConcentration) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(VolumeConcentrationPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}

// PicolitersPerLiter returns the value in PicolitersPerLiter.
func (a *VolumeConcentration) PicolitersPerLiter() float64 {
	if a.picoliters_per_literLazy != nil {
		return *a.picoliters_per_literLazy
	}
	picoliters_per_liter := a.convertFromBase(VolumeConcentrationPicolitersPerLiter)
	a.picoliters_per_literLazy = &picoliters_per_liter
	return picoliters_per_liter
}

// NanolitersPerLiter returns the value in NanolitersPerLiter.
func (a *VolumeConcentration) NanolitersPerLiter() float64 {
	if a.nanoliters_per_literLazy != nil {
		return *a.nanoliters_per_literLazy
	}
	nanoliters_per_liter := a.convertFromBase(VolumeConcentrationNanolitersPerLiter)
	a.nanoliters_per_literLazy = &nanoliters_per_liter
	return nanoliters_per_liter
}

// MicrolitersPerLiter returns the value in MicrolitersPerLiter.
func (a *VolumeConcentration) MicrolitersPerLiter() float64 {
	if a.microliters_per_literLazy != nil {
		return *a.microliters_per_literLazy
	}
	microliters_per_liter := a.convertFromBase(VolumeConcentrationMicrolitersPerLiter)
	a.microliters_per_literLazy = &microliters_per_liter
	return microliters_per_liter
}

// MillilitersPerLiter returns the value in MillilitersPerLiter.
func (a *VolumeConcentration) MillilitersPerLiter() float64 {
	if a.milliliters_per_literLazy != nil {
		return *a.milliliters_per_literLazy
	}
	milliliters_per_liter := a.convertFromBase(VolumeConcentrationMillilitersPerLiter)
	a.milliliters_per_literLazy = &milliliters_per_liter
	return milliliters_per_liter
}

// CentilitersPerLiter returns the value in CentilitersPerLiter.
func (a *VolumeConcentration) CentilitersPerLiter() float64 {
	if a.centiliters_per_literLazy != nil {
		return *a.centiliters_per_literLazy
	}
	centiliters_per_liter := a.convertFromBase(VolumeConcentrationCentilitersPerLiter)
	a.centiliters_per_literLazy = &centiliters_per_liter
	return centiliters_per_liter
}

// DecilitersPerLiter returns the value in DecilitersPerLiter.
func (a *VolumeConcentration) DecilitersPerLiter() float64 {
	if a.deciliters_per_literLazy != nil {
		return *a.deciliters_per_literLazy
	}
	deciliters_per_liter := a.convertFromBase(VolumeConcentrationDecilitersPerLiter)
	a.deciliters_per_literLazy = &deciliters_per_liter
	return deciliters_per_liter
}

// PicolitersPerMililiter returns the value in PicolitersPerMililiter.
func (a *VolumeConcentration) PicolitersPerMililiter() float64 {
	if a.picoliters_per_mililiterLazy != nil {
		return *a.picoliters_per_mililiterLazy
	}
	picoliters_per_mililiter := a.convertFromBase(VolumeConcentrationPicolitersPerMililiter)
	a.picoliters_per_mililiterLazy = &picoliters_per_mililiter
	return picoliters_per_mililiter
}

// NanolitersPerMililiter returns the value in NanolitersPerMililiter.
func (a *VolumeConcentration) NanolitersPerMililiter() float64 {
	if a.nanoliters_per_mililiterLazy != nil {
		return *a.nanoliters_per_mililiterLazy
	}
	nanoliters_per_mililiter := a.convertFromBase(VolumeConcentrationNanolitersPerMililiter)
	a.nanoliters_per_mililiterLazy = &nanoliters_per_mililiter
	return nanoliters_per_mililiter
}

// MicrolitersPerMililiter returns the value in MicrolitersPerMililiter.
func (a *VolumeConcentration) MicrolitersPerMililiter() float64 {
	if a.microliters_per_mililiterLazy != nil {
		return *a.microliters_per_mililiterLazy
	}
	microliters_per_mililiter := a.convertFromBase(VolumeConcentrationMicrolitersPerMililiter)
	a.microliters_per_mililiterLazy = &microliters_per_mililiter
	return microliters_per_mililiter
}

// MillilitersPerMililiter returns the value in MillilitersPerMililiter.
func (a *VolumeConcentration) MillilitersPerMililiter() float64 {
	if a.milliliters_per_mililiterLazy != nil {
		return *a.milliliters_per_mililiterLazy
	}
	milliliters_per_mililiter := a.convertFromBase(VolumeConcentrationMillilitersPerMililiter)
	a.milliliters_per_mililiterLazy = &milliliters_per_mililiter
	return milliliters_per_mililiter
}

// CentilitersPerMililiter returns the value in CentilitersPerMililiter.
func (a *VolumeConcentration) CentilitersPerMililiter() float64 {
	if a.centiliters_per_mililiterLazy != nil {
		return *a.centiliters_per_mililiterLazy
	}
	centiliters_per_mililiter := a.convertFromBase(VolumeConcentrationCentilitersPerMililiter)
	a.centiliters_per_mililiterLazy = &centiliters_per_mililiter
	return centiliters_per_mililiter
}

// DecilitersPerMililiter returns the value in DecilitersPerMililiter.
func (a *VolumeConcentration) DecilitersPerMililiter() float64 {
	if a.deciliters_per_mililiterLazy != nil {
		return *a.deciliters_per_mililiterLazy
	}
	deciliters_per_mililiter := a.convertFromBase(VolumeConcentrationDecilitersPerMililiter)
	a.deciliters_per_mililiterLazy = &deciliters_per_mililiter
	return deciliters_per_mililiter
}


// ToDto creates an VolumeConcentrationDto representation.
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

// ToDtoJSON creates an VolumeConcentrationDto representation.
func (a *VolumeConcentration) ToDtoJSON(holdInUnit *VolumeConcentrationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VolumeConcentration to a specific unit value.
func (a *VolumeConcentration) Convert(toUnit VolumeConcentrationUnits) float64 {
	switch toUnit { 
    case VolumeConcentrationDecimalFraction:
		return a.DecimalFractions()
    case VolumeConcentrationLitersPerLiter:
		return a.LitersPerLiter()
    case VolumeConcentrationLitersPerMililiter:
		return a.LitersPerMililiter()
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
    case VolumeConcentrationPicolitersPerLiter:
		return a.PicolitersPerLiter()
    case VolumeConcentrationNanolitersPerLiter:
		return a.NanolitersPerLiter()
    case VolumeConcentrationMicrolitersPerLiter:
		return a.MicrolitersPerLiter()
    case VolumeConcentrationMillilitersPerLiter:
		return a.MillilitersPerLiter()
    case VolumeConcentrationCentilitersPerLiter:
		return a.CentilitersPerLiter()
    case VolumeConcentrationDecilitersPerLiter:
		return a.DecilitersPerLiter()
    case VolumeConcentrationPicolitersPerMililiter:
		return a.PicolitersPerMililiter()
    case VolumeConcentrationNanolitersPerMililiter:
		return a.NanolitersPerMililiter()
    case VolumeConcentrationMicrolitersPerMililiter:
		return a.MicrolitersPerMililiter()
    case VolumeConcentrationMillilitersPerMililiter:
		return a.MillilitersPerMililiter()
    case VolumeConcentrationCentilitersPerMililiter:
		return a.CentilitersPerMililiter()
    case VolumeConcentrationDecilitersPerMililiter:
		return a.DecilitersPerMililiter()
	default:
		return 0
	}
}

func (a *VolumeConcentration) convertFromBase(toUnit VolumeConcentrationUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumeConcentrationDecimalFraction:
		return (value) 
	case VolumeConcentrationLitersPerLiter:
		return (value) 
	case VolumeConcentrationLitersPerMililiter:
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
	case VolumeConcentrationPicolitersPerLiter:
		return ((value) / 1e-12) 
	case VolumeConcentrationNanolitersPerLiter:
		return ((value) / 1e-09) 
	case VolumeConcentrationMicrolitersPerLiter:
		return ((value) / 1e-06) 
	case VolumeConcentrationMillilitersPerLiter:
		return ((value) / 0.001) 
	case VolumeConcentrationCentilitersPerLiter:
		return ((value) / 0.01) 
	case VolumeConcentrationDecilitersPerLiter:
		return ((value) / 0.1) 
	case VolumeConcentrationPicolitersPerMililiter:
		return ((value * 1e-3) / 1e-12) 
	case VolumeConcentrationNanolitersPerMililiter:
		return ((value * 1e-3) / 1e-09) 
	case VolumeConcentrationMicrolitersPerMililiter:
		return ((value * 1e-3) / 1e-06) 
	case VolumeConcentrationMillilitersPerMililiter:
		return ((value * 1e-3) / 0.001) 
	case VolumeConcentrationCentilitersPerMililiter:
		return ((value * 1e-3) / 0.01) 
	case VolumeConcentrationDecilitersPerMililiter:
		return ((value * 1e-3) / 0.1) 
	default:
		return math.NaN()
	}
}

func (a *VolumeConcentration) convertToBase(value float64, fromUnit VolumeConcentrationUnits) float64 {
	switch fromUnit { 
	case VolumeConcentrationDecimalFraction:
		return (value) 
	case VolumeConcentrationLitersPerLiter:
		return (value) 
	case VolumeConcentrationLitersPerMililiter:
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
	case VolumeConcentrationPicolitersPerLiter:
		return ((value) * 1e-12) 
	case VolumeConcentrationNanolitersPerLiter:
		return ((value) * 1e-09) 
	case VolumeConcentrationMicrolitersPerLiter:
		return ((value) * 1e-06) 
	case VolumeConcentrationMillilitersPerLiter:
		return ((value) * 0.001) 
	case VolumeConcentrationCentilitersPerLiter:
		return ((value) * 0.01) 
	case VolumeConcentrationDecilitersPerLiter:
		return ((value) * 0.1) 
	case VolumeConcentrationPicolitersPerMililiter:
		return ((value / 1e-3) * 1e-12) 
	case VolumeConcentrationNanolitersPerMililiter:
		return ((value / 1e-3) * 1e-09) 
	case VolumeConcentrationMicrolitersPerMililiter:
		return ((value / 1e-3) * 1e-06) 
	case VolumeConcentrationMillilitersPerMililiter:
		return ((value / 1e-3) * 0.001) 
	case VolumeConcentrationCentilitersPerMililiter:
		return ((value / 1e-3) * 0.01) 
	case VolumeConcentrationDecilitersPerMililiter:
		return ((value / 1e-3) * 0.1) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a VolumeConcentration) String() string {
	return a.ToString(VolumeConcentrationDecimalFraction, 2)
}

// ToString formats the VolumeConcentration to string.
// fractionalDigits -1 for not mention
func (a *VolumeConcentration) ToString(unit VolumeConcentrationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VolumeConcentration) getUnitAbbreviation(unit VolumeConcentrationUnits) string {
	switch unit { 
	case VolumeConcentrationDecimalFraction:
		return "" 
	case VolumeConcentrationLitersPerLiter:
		return "L/L" 
	case VolumeConcentrationLitersPerMililiter:
		return "L/mL" 
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
	case VolumeConcentrationPicolitersPerLiter:
		return "pL/L" 
	case VolumeConcentrationNanolitersPerLiter:
		return "nL/L" 
	case VolumeConcentrationMicrolitersPerLiter:
		return "μL/L" 
	case VolumeConcentrationMillilitersPerLiter:
		return "mL/L" 
	case VolumeConcentrationCentilitersPerLiter:
		return "cL/L" 
	case VolumeConcentrationDecilitersPerLiter:
		return "dL/L" 
	case VolumeConcentrationPicolitersPerMililiter:
		return "pL/mL" 
	case VolumeConcentrationNanolitersPerMililiter:
		return "nL/mL" 
	case VolumeConcentrationMicrolitersPerMililiter:
		return "μL/mL" 
	case VolumeConcentrationMillilitersPerMililiter:
		return "mL/mL" 
	case VolumeConcentrationCentilitersPerMililiter:
		return "cL/mL" 
	case VolumeConcentrationDecilitersPerMililiter:
		return "dL/mL" 
	default:
		return ""
	}
}

// Check if the given VolumeConcentration are equals to the current VolumeConcentration
func (a *VolumeConcentration) Equals(other *VolumeConcentration) bool {
	return a.value == other.BaseValue()
}

// Check if the given VolumeConcentration are equals to the current VolumeConcentration
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

// Add the given VolumeConcentration to the current VolumeConcentration.
func (a *VolumeConcentration) Add(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value + other.BaseValue()}
}

// Subtract the given VolumeConcentration to the current VolumeConcentration.
func (a *VolumeConcentration) Subtract(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value - other.BaseValue()}
}

// Multiply the given VolumeConcentration to the current VolumeConcentration.
func (a *VolumeConcentration) Multiply(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value * other.BaseValue()}
}

// Divide the given VolumeConcentration to the current VolumeConcentration.
func (a *VolumeConcentration) Divide(other *VolumeConcentration) *VolumeConcentration {
	return &VolumeConcentration{value: a.value / other.BaseValue()}
}