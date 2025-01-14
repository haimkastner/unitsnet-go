// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LinearDensityUnits defines various units of LinearDensity.
type LinearDensityUnits string

const (
    
        // 
        LinearDensityGramPerMillimeter LinearDensityUnits = "GramPerMillimeter"
        // 
        LinearDensityGramPerCentimeter LinearDensityUnits = "GramPerCentimeter"
        // 
        LinearDensityGramPerMeter LinearDensityUnits = "GramPerMeter"
        // 
        LinearDensityPoundPerInch LinearDensityUnits = "PoundPerInch"
        // 
        LinearDensityPoundPerFoot LinearDensityUnits = "PoundPerFoot"
        // 
        LinearDensityGramPerFoot LinearDensityUnits = "GramPerFoot"
        // 
        LinearDensityMicrogramPerMillimeter LinearDensityUnits = "MicrogramPerMillimeter"
        // 
        LinearDensityMilligramPerMillimeter LinearDensityUnits = "MilligramPerMillimeter"
        // 
        LinearDensityKilogramPerMillimeter LinearDensityUnits = "KilogramPerMillimeter"
        // 
        LinearDensityMicrogramPerCentimeter LinearDensityUnits = "MicrogramPerCentimeter"
        // 
        LinearDensityMilligramPerCentimeter LinearDensityUnits = "MilligramPerCentimeter"
        // 
        LinearDensityKilogramPerCentimeter LinearDensityUnits = "KilogramPerCentimeter"
        // 
        LinearDensityMicrogramPerMeter LinearDensityUnits = "MicrogramPerMeter"
        // 
        LinearDensityMilligramPerMeter LinearDensityUnits = "MilligramPerMeter"
        // 
        LinearDensityKilogramPerMeter LinearDensityUnits = "KilogramPerMeter"
        // 
        LinearDensityMicrogramPerFoot LinearDensityUnits = "MicrogramPerFoot"
        // 
        LinearDensityMilligramPerFoot LinearDensityUnits = "MilligramPerFoot"
        // 
        LinearDensityKilogramPerFoot LinearDensityUnits = "KilogramPerFoot"
)

// LinearDensityDto represents a LinearDensity measurement with a numerical value and its corresponding unit.
type LinearDensityDto struct {
    // Value is the numerical representation of the LinearDensity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the LinearDensity, as defined in the LinearDensityUnits enumeration.
	Unit  LinearDensityUnits `json:"unit"`
}

// LinearDensityDtoFactory groups methods for creating and serializing LinearDensityDto objects.
type LinearDensityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LinearDensityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LinearDensityDtoFactory) FromJSON(data []byte) (*LinearDensityDto, error) {
	a := LinearDensityDto{}

    // Parse JSON into LinearDensityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LinearDensityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LinearDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// LinearDensity represents a measurement in a of LinearDensity.
//
// The Linear Density, or more precisely, the linear mass density, of a substance is its mass per unit length.  The term linear density is most often used when describing the characteristics of one-dimensional objects, although linear density can also be used to describe the density of a three-dimensional quantity along one particular dimension.
type LinearDensity struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_millimeterLazy *float64 
    grams_per_centimeterLazy *float64 
    grams_per_meterLazy *float64 
    pounds_per_inchLazy *float64 
    pounds_per_footLazy *float64 
    grams_per_footLazy *float64 
    micrograms_per_millimeterLazy *float64 
    milligrams_per_millimeterLazy *float64 
    kilograms_per_millimeterLazy *float64 
    micrograms_per_centimeterLazy *float64 
    milligrams_per_centimeterLazy *float64 
    kilograms_per_centimeterLazy *float64 
    micrograms_per_meterLazy *float64 
    milligrams_per_meterLazy *float64 
    kilograms_per_meterLazy *float64 
    micrograms_per_footLazy *float64 
    milligrams_per_footLazy *float64 
    kilograms_per_footLazy *float64 
}

// LinearDensityFactory groups methods for creating LinearDensity instances.
type LinearDensityFactory struct{}

// CreateLinearDensity creates a new LinearDensity instance from the given value and unit.
func (uf LinearDensityFactory) CreateLinearDensity(value float64, unit LinearDensityUnits) (*LinearDensity, error) {
	return newLinearDensity(value, unit)
}

// FromDto converts a LinearDensityDto to a LinearDensity instance.
func (uf LinearDensityFactory) FromDto(dto LinearDensityDto) (*LinearDensity, error) {
	return newLinearDensity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a LinearDensity instance.
func (uf LinearDensityFactory) FromDtoJSON(data []byte) (*LinearDensity, error) {
	unitDto, err := LinearDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LinearDensityDto from JSON: %w", err)
	}
	return LinearDensityFactory{}.FromDto(*unitDto)
}


// FromGramsPerMillimeter creates a new LinearDensity instance from a value in GramsPerMillimeter.
func (uf LinearDensityFactory) FromGramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerMillimeter)
}

// FromGramsPerCentimeter creates a new LinearDensity instance from a value in GramsPerCentimeter.
func (uf LinearDensityFactory) FromGramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerCentimeter)
}

// FromGramsPerMeter creates a new LinearDensity instance from a value in GramsPerMeter.
func (uf LinearDensityFactory) FromGramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerMeter)
}

// FromPoundsPerInch creates a new LinearDensity instance from a value in PoundsPerInch.
func (uf LinearDensityFactory) FromPoundsPerInch(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityPoundPerInch)
}

// FromPoundsPerFoot creates a new LinearDensity instance from a value in PoundsPerFoot.
func (uf LinearDensityFactory) FromPoundsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityPoundPerFoot)
}

// FromGramsPerFoot creates a new LinearDensity instance from a value in GramsPerFoot.
func (uf LinearDensityFactory) FromGramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityGramPerFoot)
}

// FromMicrogramsPerMillimeter creates a new LinearDensity instance from a value in MicrogramsPerMillimeter.
func (uf LinearDensityFactory) FromMicrogramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerMillimeter)
}

// FromMilligramsPerMillimeter creates a new LinearDensity instance from a value in MilligramsPerMillimeter.
func (uf LinearDensityFactory) FromMilligramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerMillimeter)
}

// FromKilogramsPerMillimeter creates a new LinearDensity instance from a value in KilogramsPerMillimeter.
func (uf LinearDensityFactory) FromKilogramsPerMillimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerMillimeter)
}

// FromMicrogramsPerCentimeter creates a new LinearDensity instance from a value in MicrogramsPerCentimeter.
func (uf LinearDensityFactory) FromMicrogramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerCentimeter)
}

// FromMilligramsPerCentimeter creates a new LinearDensity instance from a value in MilligramsPerCentimeter.
func (uf LinearDensityFactory) FromMilligramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerCentimeter)
}

// FromKilogramsPerCentimeter creates a new LinearDensity instance from a value in KilogramsPerCentimeter.
func (uf LinearDensityFactory) FromKilogramsPerCentimeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerCentimeter)
}

// FromMicrogramsPerMeter creates a new LinearDensity instance from a value in MicrogramsPerMeter.
func (uf LinearDensityFactory) FromMicrogramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerMeter)
}

// FromMilligramsPerMeter creates a new LinearDensity instance from a value in MilligramsPerMeter.
func (uf LinearDensityFactory) FromMilligramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerMeter)
}

// FromKilogramsPerMeter creates a new LinearDensity instance from a value in KilogramsPerMeter.
func (uf LinearDensityFactory) FromKilogramsPerMeter(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerMeter)
}

// FromMicrogramsPerFoot creates a new LinearDensity instance from a value in MicrogramsPerFoot.
func (uf LinearDensityFactory) FromMicrogramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMicrogramPerFoot)
}

// FromMilligramsPerFoot creates a new LinearDensity instance from a value in MilligramsPerFoot.
func (uf LinearDensityFactory) FromMilligramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityMilligramPerFoot)
}

// FromKilogramsPerFoot creates a new LinearDensity instance from a value in KilogramsPerFoot.
func (uf LinearDensityFactory) FromKilogramsPerFoot(value float64) (*LinearDensity, error) {
	return newLinearDensity(value, LinearDensityKilogramPerFoot)
}


// newLinearDensity creates a new LinearDensity.
func newLinearDensity(value float64, fromUnit LinearDensityUnits) (*LinearDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &LinearDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of LinearDensity in KilogramPerMeter unit (the base/default quantity).
func (a *LinearDensity) BaseValue() float64 {
	return a.value
}


// GramsPerMillimeter returns the LinearDensity value in GramsPerMillimeter.
//
// 
func (a *LinearDensity) GramsPerMillimeter() float64 {
	if a.grams_per_millimeterLazy != nil {
		return *a.grams_per_millimeterLazy
	}
	grams_per_millimeter := a.convertFromBase(LinearDensityGramPerMillimeter)
	a.grams_per_millimeterLazy = &grams_per_millimeter
	return grams_per_millimeter
}

// GramsPerCentimeter returns the LinearDensity value in GramsPerCentimeter.
//
// 
func (a *LinearDensity) GramsPerCentimeter() float64 {
	if a.grams_per_centimeterLazy != nil {
		return *a.grams_per_centimeterLazy
	}
	grams_per_centimeter := a.convertFromBase(LinearDensityGramPerCentimeter)
	a.grams_per_centimeterLazy = &grams_per_centimeter
	return grams_per_centimeter
}

// GramsPerMeter returns the LinearDensity value in GramsPerMeter.
//
// 
func (a *LinearDensity) GramsPerMeter() float64 {
	if a.grams_per_meterLazy != nil {
		return *a.grams_per_meterLazy
	}
	grams_per_meter := a.convertFromBase(LinearDensityGramPerMeter)
	a.grams_per_meterLazy = &grams_per_meter
	return grams_per_meter
}

// PoundsPerInch returns the LinearDensity value in PoundsPerInch.
//
// 
func (a *LinearDensity) PoundsPerInch() float64 {
	if a.pounds_per_inchLazy != nil {
		return *a.pounds_per_inchLazy
	}
	pounds_per_inch := a.convertFromBase(LinearDensityPoundPerInch)
	a.pounds_per_inchLazy = &pounds_per_inch
	return pounds_per_inch
}

// PoundsPerFoot returns the LinearDensity value in PoundsPerFoot.
//
// 
func (a *LinearDensity) PoundsPerFoot() float64 {
	if a.pounds_per_footLazy != nil {
		return *a.pounds_per_footLazy
	}
	pounds_per_foot := a.convertFromBase(LinearDensityPoundPerFoot)
	a.pounds_per_footLazy = &pounds_per_foot
	return pounds_per_foot
}

// GramsPerFoot returns the LinearDensity value in GramsPerFoot.
//
// 
func (a *LinearDensity) GramsPerFoot() float64 {
	if a.grams_per_footLazy != nil {
		return *a.grams_per_footLazy
	}
	grams_per_foot := a.convertFromBase(LinearDensityGramPerFoot)
	a.grams_per_footLazy = &grams_per_foot
	return grams_per_foot
}

// MicrogramsPerMillimeter returns the LinearDensity value in MicrogramsPerMillimeter.
//
// 
func (a *LinearDensity) MicrogramsPerMillimeter() float64 {
	if a.micrograms_per_millimeterLazy != nil {
		return *a.micrograms_per_millimeterLazy
	}
	micrograms_per_millimeter := a.convertFromBase(LinearDensityMicrogramPerMillimeter)
	a.micrograms_per_millimeterLazy = &micrograms_per_millimeter
	return micrograms_per_millimeter
}

// MilligramsPerMillimeter returns the LinearDensity value in MilligramsPerMillimeter.
//
// 
func (a *LinearDensity) MilligramsPerMillimeter() float64 {
	if a.milligrams_per_millimeterLazy != nil {
		return *a.milligrams_per_millimeterLazy
	}
	milligrams_per_millimeter := a.convertFromBase(LinearDensityMilligramPerMillimeter)
	a.milligrams_per_millimeterLazy = &milligrams_per_millimeter
	return milligrams_per_millimeter
}

// KilogramsPerMillimeter returns the LinearDensity value in KilogramsPerMillimeter.
//
// 
func (a *LinearDensity) KilogramsPerMillimeter() float64 {
	if a.kilograms_per_millimeterLazy != nil {
		return *a.kilograms_per_millimeterLazy
	}
	kilograms_per_millimeter := a.convertFromBase(LinearDensityKilogramPerMillimeter)
	a.kilograms_per_millimeterLazy = &kilograms_per_millimeter
	return kilograms_per_millimeter
}

// MicrogramsPerCentimeter returns the LinearDensity value in MicrogramsPerCentimeter.
//
// 
func (a *LinearDensity) MicrogramsPerCentimeter() float64 {
	if a.micrograms_per_centimeterLazy != nil {
		return *a.micrograms_per_centimeterLazy
	}
	micrograms_per_centimeter := a.convertFromBase(LinearDensityMicrogramPerCentimeter)
	a.micrograms_per_centimeterLazy = &micrograms_per_centimeter
	return micrograms_per_centimeter
}

// MilligramsPerCentimeter returns the LinearDensity value in MilligramsPerCentimeter.
//
// 
func (a *LinearDensity) MilligramsPerCentimeter() float64 {
	if a.milligrams_per_centimeterLazy != nil {
		return *a.milligrams_per_centimeterLazy
	}
	milligrams_per_centimeter := a.convertFromBase(LinearDensityMilligramPerCentimeter)
	a.milligrams_per_centimeterLazy = &milligrams_per_centimeter
	return milligrams_per_centimeter
}

// KilogramsPerCentimeter returns the LinearDensity value in KilogramsPerCentimeter.
//
// 
func (a *LinearDensity) KilogramsPerCentimeter() float64 {
	if a.kilograms_per_centimeterLazy != nil {
		return *a.kilograms_per_centimeterLazy
	}
	kilograms_per_centimeter := a.convertFromBase(LinearDensityKilogramPerCentimeter)
	a.kilograms_per_centimeterLazy = &kilograms_per_centimeter
	return kilograms_per_centimeter
}

// MicrogramsPerMeter returns the LinearDensity value in MicrogramsPerMeter.
//
// 
func (a *LinearDensity) MicrogramsPerMeter() float64 {
	if a.micrograms_per_meterLazy != nil {
		return *a.micrograms_per_meterLazy
	}
	micrograms_per_meter := a.convertFromBase(LinearDensityMicrogramPerMeter)
	a.micrograms_per_meterLazy = &micrograms_per_meter
	return micrograms_per_meter
}

// MilligramsPerMeter returns the LinearDensity value in MilligramsPerMeter.
//
// 
func (a *LinearDensity) MilligramsPerMeter() float64 {
	if a.milligrams_per_meterLazy != nil {
		return *a.milligrams_per_meterLazy
	}
	milligrams_per_meter := a.convertFromBase(LinearDensityMilligramPerMeter)
	a.milligrams_per_meterLazy = &milligrams_per_meter
	return milligrams_per_meter
}

// KilogramsPerMeter returns the LinearDensity value in KilogramsPerMeter.
//
// 
func (a *LinearDensity) KilogramsPerMeter() float64 {
	if a.kilograms_per_meterLazy != nil {
		return *a.kilograms_per_meterLazy
	}
	kilograms_per_meter := a.convertFromBase(LinearDensityKilogramPerMeter)
	a.kilograms_per_meterLazy = &kilograms_per_meter
	return kilograms_per_meter
}

// MicrogramsPerFoot returns the LinearDensity value in MicrogramsPerFoot.
//
// 
func (a *LinearDensity) MicrogramsPerFoot() float64 {
	if a.micrograms_per_footLazy != nil {
		return *a.micrograms_per_footLazy
	}
	micrograms_per_foot := a.convertFromBase(LinearDensityMicrogramPerFoot)
	a.micrograms_per_footLazy = &micrograms_per_foot
	return micrograms_per_foot
}

// MilligramsPerFoot returns the LinearDensity value in MilligramsPerFoot.
//
// 
func (a *LinearDensity) MilligramsPerFoot() float64 {
	if a.milligrams_per_footLazy != nil {
		return *a.milligrams_per_footLazy
	}
	milligrams_per_foot := a.convertFromBase(LinearDensityMilligramPerFoot)
	a.milligrams_per_footLazy = &milligrams_per_foot
	return milligrams_per_foot
}

// KilogramsPerFoot returns the LinearDensity value in KilogramsPerFoot.
//
// 
func (a *LinearDensity) KilogramsPerFoot() float64 {
	if a.kilograms_per_footLazy != nil {
		return *a.kilograms_per_footLazy
	}
	kilograms_per_foot := a.convertFromBase(LinearDensityKilogramPerFoot)
	a.kilograms_per_footLazy = &kilograms_per_foot
	return kilograms_per_foot
}


// ToDto creates a LinearDensityDto representation from the LinearDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerMeter by default.
func (a *LinearDensity) ToDto(holdInUnit *LinearDensityUnits) LinearDensityDto {
	if holdInUnit == nil {
		defaultUnit := LinearDensityKilogramPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LinearDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the LinearDensity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerMeter by default.
func (a *LinearDensity) ToDtoJSON(holdInUnit *LinearDensityUnits) ([]byte, error) {
	// Convert to LinearDensityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a LinearDensity to a specific unit value.
// The function uses the provided unit type (LinearDensityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *LinearDensity) Convert(toUnit LinearDensityUnits) float64 {
	switch toUnit { 
    case LinearDensityGramPerMillimeter:
		return a.GramsPerMillimeter()
    case LinearDensityGramPerCentimeter:
		return a.GramsPerCentimeter()
    case LinearDensityGramPerMeter:
		return a.GramsPerMeter()
    case LinearDensityPoundPerInch:
		return a.PoundsPerInch()
    case LinearDensityPoundPerFoot:
		return a.PoundsPerFoot()
    case LinearDensityGramPerFoot:
		return a.GramsPerFoot()
    case LinearDensityMicrogramPerMillimeter:
		return a.MicrogramsPerMillimeter()
    case LinearDensityMilligramPerMillimeter:
		return a.MilligramsPerMillimeter()
    case LinearDensityKilogramPerMillimeter:
		return a.KilogramsPerMillimeter()
    case LinearDensityMicrogramPerCentimeter:
		return a.MicrogramsPerCentimeter()
    case LinearDensityMilligramPerCentimeter:
		return a.MilligramsPerCentimeter()
    case LinearDensityKilogramPerCentimeter:
		return a.KilogramsPerCentimeter()
    case LinearDensityMicrogramPerMeter:
		return a.MicrogramsPerMeter()
    case LinearDensityMilligramPerMeter:
		return a.MilligramsPerMeter()
    case LinearDensityKilogramPerMeter:
		return a.KilogramsPerMeter()
    case LinearDensityMicrogramPerFoot:
		return a.MicrogramsPerFoot()
    case LinearDensityMilligramPerFoot:
		return a.MilligramsPerFoot()
    case LinearDensityKilogramPerFoot:
		return a.KilogramsPerFoot()
	default:
		return math.NaN()
	}
}

func (a *LinearDensity) convertFromBase(toUnit LinearDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LinearDensityGramPerMillimeter:
		return (value) 
	case LinearDensityGramPerCentimeter:
		return (value / 1e-1) 
	case LinearDensityGramPerMeter:
		return (value / 1e-3) 
	case LinearDensityPoundPerInch:
		return (value * 5.5997415e-2) 
	case LinearDensityPoundPerFoot:
		return (value / 1.48816394) 
	case LinearDensityGramPerFoot:
		return (value / ( 1e-3 / 0.3048 )) 
	case LinearDensityMicrogramPerMillimeter:
		return ((value) / 1e-06) 
	case LinearDensityMilligramPerMillimeter:
		return ((value) / 0.001) 
	case LinearDensityKilogramPerMillimeter:
		return ((value) / 1000.0) 
	case LinearDensityMicrogramPerCentimeter:
		return ((value / 1e-1) / 1e-06) 
	case LinearDensityMilligramPerCentimeter:
		return ((value / 1e-1) / 0.001) 
	case LinearDensityKilogramPerCentimeter:
		return ((value / 1e-1) / 1000.0) 
	case LinearDensityMicrogramPerMeter:
		return ((value / 1e-3) / 1e-06) 
	case LinearDensityMilligramPerMeter:
		return ((value / 1e-3) / 0.001) 
	case LinearDensityKilogramPerMeter:
		return ((value / 1e-3) / 1000.0) 
	case LinearDensityMicrogramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 1e-06) 
	case LinearDensityMilligramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 0.001) 
	case LinearDensityKilogramPerFoot:
		return ((value / ( 1e-3 / 0.3048 )) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *LinearDensity) convertToBase(value float64, fromUnit LinearDensityUnits) float64 {
	switch fromUnit { 
	case LinearDensityGramPerMillimeter:
		return (value) 
	case LinearDensityGramPerCentimeter:
		return (value * 1e-1) 
	case LinearDensityGramPerMeter:
		return (value * 1e-3) 
	case LinearDensityPoundPerInch:
		return (value / 5.5997415e-2) 
	case LinearDensityPoundPerFoot:
		return (value * 1.48816394) 
	case LinearDensityGramPerFoot:
		return (value * ( 1e-3 / 0.3048 )) 
	case LinearDensityMicrogramPerMillimeter:
		return ((value) * 1e-06) 
	case LinearDensityMilligramPerMillimeter:
		return ((value) * 0.001) 
	case LinearDensityKilogramPerMillimeter:
		return ((value) * 1000.0) 
	case LinearDensityMicrogramPerCentimeter:
		return ((value * 1e-1) * 1e-06) 
	case LinearDensityMilligramPerCentimeter:
		return ((value * 1e-1) * 0.001) 
	case LinearDensityKilogramPerCentimeter:
		return ((value * 1e-1) * 1000.0) 
	case LinearDensityMicrogramPerMeter:
		return ((value * 1e-3) * 1e-06) 
	case LinearDensityMilligramPerMeter:
		return ((value * 1e-3) * 0.001) 
	case LinearDensityKilogramPerMeter:
		return ((value * 1e-3) * 1000.0) 
	case LinearDensityMicrogramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 1e-06) 
	case LinearDensityMilligramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 0.001) 
	case LinearDensityKilogramPerFoot:
		return ((value * ( 1e-3 / 0.3048 )) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the LinearDensity in the default unit (KilogramPerMeter),
// formatted to two decimal places.
func (a LinearDensity) String() string {
	return a.ToString(LinearDensityKilogramPerMeter, 2)
}

// ToString formats the LinearDensity value as a string with the specified unit and fractional digits.
// It converts the LinearDensity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the LinearDensity value will be converted (e.g., KilogramPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the LinearDensity with the unit abbreviation.
func (a *LinearDensity) ToString(unit LinearDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *LinearDensity) getUnitAbbreviation(unit LinearDensityUnits) string {
	switch unit { 
	case LinearDensityGramPerMillimeter:
		return "g/mm" 
	case LinearDensityGramPerCentimeter:
		return "g/cm" 
	case LinearDensityGramPerMeter:
		return "g/m" 
	case LinearDensityPoundPerInch:
		return "lb/in" 
	case LinearDensityPoundPerFoot:
		return "lb/ft" 
	case LinearDensityGramPerFoot:
		return "g/ft" 
	case LinearDensityMicrogramPerMillimeter:
		return "μg/mm" 
	case LinearDensityMilligramPerMillimeter:
		return "mg/mm" 
	case LinearDensityKilogramPerMillimeter:
		return "kg/mm" 
	case LinearDensityMicrogramPerCentimeter:
		return "μg/cm" 
	case LinearDensityMilligramPerCentimeter:
		return "mg/cm" 
	case LinearDensityKilogramPerCentimeter:
		return "kg/cm" 
	case LinearDensityMicrogramPerMeter:
		return "μg/m" 
	case LinearDensityMilligramPerMeter:
		return "mg/m" 
	case LinearDensityKilogramPerMeter:
		return "kg/m" 
	case LinearDensityMicrogramPerFoot:
		return "μg/ft" 
	case LinearDensityMilligramPerFoot:
		return "mg/ft" 
	case LinearDensityKilogramPerFoot:
		return "kg/ft" 
	default:
		return ""
	}
}

// Equals checks if the given LinearDensity is equal to the current LinearDensity.
//
// Parameters:
//    other: The LinearDensity to compare against.
//
// Returns:
//    bool: Returns true if both LinearDensity are equal, false otherwise.
func (a *LinearDensity) Equals(other *LinearDensity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current LinearDensity with another LinearDensity.
// It returns -1 if the current LinearDensity is less than the other LinearDensity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The LinearDensity to compare against.
//
// Returns:
//    int: -1 if the current LinearDensity is less, 1 if greater, and 0 if equal.
func (a *LinearDensity) CompareTo(other *LinearDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given LinearDensity to the current LinearDensity and returns the result.
// The result is a new LinearDensity instance with the sum of the values.
//
// Parameters:
//    other: The LinearDensity to add to the current LinearDensity.
//
// Returns:
//    *LinearDensity: A new LinearDensity instance representing the sum of both LinearDensity.
func (a *LinearDensity) Add(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given LinearDensity from the current LinearDensity and returns the result.
// The result is a new LinearDensity instance with the difference of the values.
//
// Parameters:
//    other: The LinearDensity to subtract from the current LinearDensity.
//
// Returns:
//    *LinearDensity: A new LinearDensity instance representing the difference of both LinearDensity.
func (a *LinearDensity) Subtract(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current LinearDensity by the given LinearDensity and returns the result.
// The result is a new LinearDensity instance with the product of the values.
//
// Parameters:
//    other: The LinearDensity to multiply with the current LinearDensity.
//
// Returns:
//    *LinearDensity: A new LinearDensity instance representing the product of both LinearDensity.
func (a *LinearDensity) Multiply(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value * other.BaseValue()}
}

// Divide divides the current LinearDensity by the given LinearDensity and returns the result.
// The result is a new LinearDensity instance with the quotient of the values.
//
// Parameters:
//    other: The LinearDensity to divide the current LinearDensity by.
//
// Returns:
//    *LinearDensity: A new LinearDensity instance representing the quotient of both LinearDensity.
func (a *LinearDensity) Divide(other *LinearDensity) *LinearDensity {
	return &LinearDensity{value: a.value / other.BaseValue()}
}