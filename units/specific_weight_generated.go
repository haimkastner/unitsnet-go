// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificWeightUnits defines various units of SpecificWeight.
type SpecificWeightUnits string

const (
    
        // 
        SpecificWeightNewtonPerCubicMillimeter SpecificWeightUnits = "NewtonPerCubicMillimeter"
        // 
        SpecificWeightNewtonPerCubicCentimeter SpecificWeightUnits = "NewtonPerCubicCentimeter"
        // 
        SpecificWeightNewtonPerCubicMeter SpecificWeightUnits = "NewtonPerCubicMeter"
        // 
        SpecificWeightKilogramForcePerCubicMillimeter SpecificWeightUnits = "KilogramForcePerCubicMillimeter"
        // 
        SpecificWeightKilogramForcePerCubicCentimeter SpecificWeightUnits = "KilogramForcePerCubicCentimeter"
        // 
        SpecificWeightKilogramForcePerCubicMeter SpecificWeightUnits = "KilogramForcePerCubicMeter"
        // 
        SpecificWeightPoundForcePerCubicInch SpecificWeightUnits = "PoundForcePerCubicInch"
        // 
        SpecificWeightPoundForcePerCubicFoot SpecificWeightUnits = "PoundForcePerCubicFoot"
        // 
        SpecificWeightTonneForcePerCubicMillimeter SpecificWeightUnits = "TonneForcePerCubicMillimeter"
        // 
        SpecificWeightTonneForcePerCubicCentimeter SpecificWeightUnits = "TonneForcePerCubicCentimeter"
        // 
        SpecificWeightTonneForcePerCubicMeter SpecificWeightUnits = "TonneForcePerCubicMeter"
        // 
        SpecificWeightKilonewtonPerCubicMillimeter SpecificWeightUnits = "KilonewtonPerCubicMillimeter"
        // 
        SpecificWeightKilonewtonPerCubicCentimeter SpecificWeightUnits = "KilonewtonPerCubicCentimeter"
        // 
        SpecificWeightKilonewtonPerCubicMeter SpecificWeightUnits = "KilonewtonPerCubicMeter"
        // 
        SpecificWeightMeganewtonPerCubicMeter SpecificWeightUnits = "MeganewtonPerCubicMeter"
        // 
        SpecificWeightKilopoundForcePerCubicInch SpecificWeightUnits = "KilopoundForcePerCubicInch"
        // 
        SpecificWeightKilopoundForcePerCubicFoot SpecificWeightUnits = "KilopoundForcePerCubicFoot"
)

var internalSpecificWeightUnitsMap = map[SpecificWeightUnits]bool{
	
	SpecificWeightNewtonPerCubicMillimeter: true,
	SpecificWeightNewtonPerCubicCentimeter: true,
	SpecificWeightNewtonPerCubicMeter: true,
	SpecificWeightKilogramForcePerCubicMillimeter: true,
	SpecificWeightKilogramForcePerCubicCentimeter: true,
	SpecificWeightKilogramForcePerCubicMeter: true,
	SpecificWeightPoundForcePerCubicInch: true,
	SpecificWeightPoundForcePerCubicFoot: true,
	SpecificWeightTonneForcePerCubicMillimeter: true,
	SpecificWeightTonneForcePerCubicCentimeter: true,
	SpecificWeightTonneForcePerCubicMeter: true,
	SpecificWeightKilonewtonPerCubicMillimeter: true,
	SpecificWeightKilonewtonPerCubicCentimeter: true,
	SpecificWeightKilonewtonPerCubicMeter: true,
	SpecificWeightMeganewtonPerCubicMeter: true,
	SpecificWeightKilopoundForcePerCubicInch: true,
	SpecificWeightKilopoundForcePerCubicFoot: true,
}

// SpecificWeightDto represents a SpecificWeight measurement with a numerical value and its corresponding unit.
type SpecificWeightDto struct {
    // Value is the numerical representation of the SpecificWeight.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the SpecificWeight, as defined in the SpecificWeightUnits enumeration.
	Unit  SpecificWeightUnits `json:"unit" validate:"required,oneof=NewtonPerCubicMillimeter NewtonPerCubicCentimeter NewtonPerCubicMeter KilogramForcePerCubicMillimeter KilogramForcePerCubicCentimeter KilogramForcePerCubicMeter PoundForcePerCubicInch PoundForcePerCubicFoot TonneForcePerCubicMillimeter TonneForcePerCubicCentimeter TonneForcePerCubicMeter KilonewtonPerCubicMillimeter KilonewtonPerCubicCentimeter KilonewtonPerCubicMeter MeganewtonPerCubicMeter KilopoundForcePerCubicInch KilopoundForcePerCubicFoot"`
}

// SpecificWeightDtoFactory groups methods for creating and serializing SpecificWeightDto objects.
type SpecificWeightDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpecificWeightDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpecificWeightDtoFactory) FromJSON(data []byte) (*SpecificWeightDto, error) {
	a := SpecificWeightDto{}

    // Parse JSON into SpecificWeightDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a SpecificWeightDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpecificWeightDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SpecificWeight represents a measurement in a of SpecificWeight.
//
// The SpecificWeight, or more precisely, the volumetric weight density, of a substance is its weight per unit volume.
type SpecificWeight struct {
	// value is the base measurement stored internally.
	value       float64
    
    newtons_per_cubic_millimeterLazy *float64 
    newtons_per_cubic_centimeterLazy *float64 
    newtons_per_cubic_meterLazy *float64 
    kilograms_force_per_cubic_millimeterLazy *float64 
    kilograms_force_per_cubic_centimeterLazy *float64 
    kilograms_force_per_cubic_meterLazy *float64 
    pounds_force_per_cubic_inchLazy *float64 
    pounds_force_per_cubic_footLazy *float64 
    tonnes_force_per_cubic_millimeterLazy *float64 
    tonnes_force_per_cubic_centimeterLazy *float64 
    tonnes_force_per_cubic_meterLazy *float64 
    kilonewtons_per_cubic_millimeterLazy *float64 
    kilonewtons_per_cubic_centimeterLazy *float64 
    kilonewtons_per_cubic_meterLazy *float64 
    meganewtons_per_cubic_meterLazy *float64 
    kilopounds_force_per_cubic_inchLazy *float64 
    kilopounds_force_per_cubic_footLazy *float64 
}

// SpecificWeightFactory groups methods for creating SpecificWeight instances.
type SpecificWeightFactory struct{}

// CreateSpecificWeight creates a new SpecificWeight instance from the given value and unit.
func (uf SpecificWeightFactory) CreateSpecificWeight(value float64, unit SpecificWeightUnits) (*SpecificWeight, error) {
	return newSpecificWeight(value, unit)
}

// FromDto converts a SpecificWeightDto to a SpecificWeight instance.
func (uf SpecificWeightFactory) FromDto(dto SpecificWeightDto) (*SpecificWeight, error) {
	return newSpecificWeight(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SpecificWeight instance.
func (uf SpecificWeightFactory) FromDtoJSON(data []byte) (*SpecificWeight, error) {
	unitDto, err := SpecificWeightDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpecificWeightDto from JSON: %w", err)
	}
	return SpecificWeightFactory{}.FromDto(*unitDto)
}


// FromNewtonsPerCubicMillimeter creates a new SpecificWeight instance from a value in NewtonsPerCubicMillimeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicMillimeter)
}

// FromNewtonsPerCubicCentimeter creates a new SpecificWeight instance from a value in NewtonsPerCubicCentimeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicCentimeter)
}

// FromNewtonsPerCubicMeter creates a new SpecificWeight instance from a value in NewtonsPerCubicMeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicMeter)
}

// FromKilogramsForcePerCubicMillimeter creates a new SpecificWeight instance from a value in KilogramsForcePerCubicMillimeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicMillimeter)
}

// FromKilogramsForcePerCubicCentimeter creates a new SpecificWeight instance from a value in KilogramsForcePerCubicCentimeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicCentimeter)
}

// FromKilogramsForcePerCubicMeter creates a new SpecificWeight instance from a value in KilogramsForcePerCubicMeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicMeter)
}

// FromPoundsForcePerCubicInch creates a new SpecificWeight instance from a value in PoundsForcePerCubicInch.
func (uf SpecificWeightFactory) FromPoundsForcePerCubicInch(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightPoundForcePerCubicInch)
}

// FromPoundsForcePerCubicFoot creates a new SpecificWeight instance from a value in PoundsForcePerCubicFoot.
func (uf SpecificWeightFactory) FromPoundsForcePerCubicFoot(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightPoundForcePerCubicFoot)
}

// FromTonnesForcePerCubicMillimeter creates a new SpecificWeight instance from a value in TonnesForcePerCubicMillimeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicMillimeter)
}

// FromTonnesForcePerCubicCentimeter creates a new SpecificWeight instance from a value in TonnesForcePerCubicCentimeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicCentimeter)
}

// FromTonnesForcePerCubicMeter creates a new SpecificWeight instance from a value in TonnesForcePerCubicMeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicMeter)
}

// FromKilonewtonsPerCubicMillimeter creates a new SpecificWeight instance from a value in KilonewtonsPerCubicMillimeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicMillimeter)
}

// FromKilonewtonsPerCubicCentimeter creates a new SpecificWeight instance from a value in KilonewtonsPerCubicCentimeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicCentimeter)
}

// FromKilonewtonsPerCubicMeter creates a new SpecificWeight instance from a value in KilonewtonsPerCubicMeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicMeter)
}

// FromMeganewtonsPerCubicMeter creates a new SpecificWeight instance from a value in MeganewtonsPerCubicMeter.
func (uf SpecificWeightFactory) FromMeganewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightMeganewtonPerCubicMeter)
}

// FromKilopoundsForcePerCubicInch creates a new SpecificWeight instance from a value in KilopoundsForcePerCubicInch.
func (uf SpecificWeightFactory) FromKilopoundsForcePerCubicInch(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilopoundForcePerCubicInch)
}

// FromKilopoundsForcePerCubicFoot creates a new SpecificWeight instance from a value in KilopoundsForcePerCubicFoot.
func (uf SpecificWeightFactory) FromKilopoundsForcePerCubicFoot(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilopoundForcePerCubicFoot)
}


// newSpecificWeight creates a new SpecificWeight.
func newSpecificWeight(value float64, fromUnit SpecificWeightUnits) (*SpecificWeight, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalSpecificWeightUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in SpecificWeightUnits", fromUnit)
	}
	a := &SpecificWeight{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificWeight in NewtonPerCubicMeter unit (the base/default quantity).
func (a *SpecificWeight) BaseValue() float64 {
	return a.value
}


// NewtonsPerCubicMillimeter returns the SpecificWeight value in NewtonsPerCubicMillimeter.
//
// 
func (a *SpecificWeight) NewtonsPerCubicMillimeter() float64 {
	if a.newtons_per_cubic_millimeterLazy != nil {
		return *a.newtons_per_cubic_millimeterLazy
	}
	newtons_per_cubic_millimeter := a.convertFromBase(SpecificWeightNewtonPerCubicMillimeter)
	a.newtons_per_cubic_millimeterLazy = &newtons_per_cubic_millimeter
	return newtons_per_cubic_millimeter
}

// NewtonsPerCubicCentimeter returns the SpecificWeight value in NewtonsPerCubicCentimeter.
//
// 
func (a *SpecificWeight) NewtonsPerCubicCentimeter() float64 {
	if a.newtons_per_cubic_centimeterLazy != nil {
		return *a.newtons_per_cubic_centimeterLazy
	}
	newtons_per_cubic_centimeter := a.convertFromBase(SpecificWeightNewtonPerCubicCentimeter)
	a.newtons_per_cubic_centimeterLazy = &newtons_per_cubic_centimeter
	return newtons_per_cubic_centimeter
}

// NewtonsPerCubicMeter returns the SpecificWeight value in NewtonsPerCubicMeter.
//
// 
func (a *SpecificWeight) NewtonsPerCubicMeter() float64 {
	if a.newtons_per_cubic_meterLazy != nil {
		return *a.newtons_per_cubic_meterLazy
	}
	newtons_per_cubic_meter := a.convertFromBase(SpecificWeightNewtonPerCubicMeter)
	a.newtons_per_cubic_meterLazy = &newtons_per_cubic_meter
	return newtons_per_cubic_meter
}

// KilogramsForcePerCubicMillimeter returns the SpecificWeight value in KilogramsForcePerCubicMillimeter.
//
// 
func (a *SpecificWeight) KilogramsForcePerCubicMillimeter() float64 {
	if a.kilograms_force_per_cubic_millimeterLazy != nil {
		return *a.kilograms_force_per_cubic_millimeterLazy
	}
	kilograms_force_per_cubic_millimeter := a.convertFromBase(SpecificWeightKilogramForcePerCubicMillimeter)
	a.kilograms_force_per_cubic_millimeterLazy = &kilograms_force_per_cubic_millimeter
	return kilograms_force_per_cubic_millimeter
}

// KilogramsForcePerCubicCentimeter returns the SpecificWeight value in KilogramsForcePerCubicCentimeter.
//
// 
func (a *SpecificWeight) KilogramsForcePerCubicCentimeter() float64 {
	if a.kilograms_force_per_cubic_centimeterLazy != nil {
		return *a.kilograms_force_per_cubic_centimeterLazy
	}
	kilograms_force_per_cubic_centimeter := a.convertFromBase(SpecificWeightKilogramForcePerCubicCentimeter)
	a.kilograms_force_per_cubic_centimeterLazy = &kilograms_force_per_cubic_centimeter
	return kilograms_force_per_cubic_centimeter
}

// KilogramsForcePerCubicMeter returns the SpecificWeight value in KilogramsForcePerCubicMeter.
//
// 
func (a *SpecificWeight) KilogramsForcePerCubicMeter() float64 {
	if a.kilograms_force_per_cubic_meterLazy != nil {
		return *a.kilograms_force_per_cubic_meterLazy
	}
	kilograms_force_per_cubic_meter := a.convertFromBase(SpecificWeightKilogramForcePerCubicMeter)
	a.kilograms_force_per_cubic_meterLazy = &kilograms_force_per_cubic_meter
	return kilograms_force_per_cubic_meter
}

// PoundsForcePerCubicInch returns the SpecificWeight value in PoundsForcePerCubicInch.
//
// 
func (a *SpecificWeight) PoundsForcePerCubicInch() float64 {
	if a.pounds_force_per_cubic_inchLazy != nil {
		return *a.pounds_force_per_cubic_inchLazy
	}
	pounds_force_per_cubic_inch := a.convertFromBase(SpecificWeightPoundForcePerCubicInch)
	a.pounds_force_per_cubic_inchLazy = &pounds_force_per_cubic_inch
	return pounds_force_per_cubic_inch
}

// PoundsForcePerCubicFoot returns the SpecificWeight value in PoundsForcePerCubicFoot.
//
// 
func (a *SpecificWeight) PoundsForcePerCubicFoot() float64 {
	if a.pounds_force_per_cubic_footLazy != nil {
		return *a.pounds_force_per_cubic_footLazy
	}
	pounds_force_per_cubic_foot := a.convertFromBase(SpecificWeightPoundForcePerCubicFoot)
	a.pounds_force_per_cubic_footLazy = &pounds_force_per_cubic_foot
	return pounds_force_per_cubic_foot
}

// TonnesForcePerCubicMillimeter returns the SpecificWeight value in TonnesForcePerCubicMillimeter.
//
// 
func (a *SpecificWeight) TonnesForcePerCubicMillimeter() float64 {
	if a.tonnes_force_per_cubic_millimeterLazy != nil {
		return *a.tonnes_force_per_cubic_millimeterLazy
	}
	tonnes_force_per_cubic_millimeter := a.convertFromBase(SpecificWeightTonneForcePerCubicMillimeter)
	a.tonnes_force_per_cubic_millimeterLazy = &tonnes_force_per_cubic_millimeter
	return tonnes_force_per_cubic_millimeter
}

// TonnesForcePerCubicCentimeter returns the SpecificWeight value in TonnesForcePerCubicCentimeter.
//
// 
func (a *SpecificWeight) TonnesForcePerCubicCentimeter() float64 {
	if a.tonnes_force_per_cubic_centimeterLazy != nil {
		return *a.tonnes_force_per_cubic_centimeterLazy
	}
	tonnes_force_per_cubic_centimeter := a.convertFromBase(SpecificWeightTonneForcePerCubicCentimeter)
	a.tonnes_force_per_cubic_centimeterLazy = &tonnes_force_per_cubic_centimeter
	return tonnes_force_per_cubic_centimeter
}

// TonnesForcePerCubicMeter returns the SpecificWeight value in TonnesForcePerCubicMeter.
//
// 
func (a *SpecificWeight) TonnesForcePerCubicMeter() float64 {
	if a.tonnes_force_per_cubic_meterLazy != nil {
		return *a.tonnes_force_per_cubic_meterLazy
	}
	tonnes_force_per_cubic_meter := a.convertFromBase(SpecificWeightTonneForcePerCubicMeter)
	a.tonnes_force_per_cubic_meterLazy = &tonnes_force_per_cubic_meter
	return tonnes_force_per_cubic_meter
}

// KilonewtonsPerCubicMillimeter returns the SpecificWeight value in KilonewtonsPerCubicMillimeter.
//
// 
func (a *SpecificWeight) KilonewtonsPerCubicMillimeter() float64 {
	if a.kilonewtons_per_cubic_millimeterLazy != nil {
		return *a.kilonewtons_per_cubic_millimeterLazy
	}
	kilonewtons_per_cubic_millimeter := a.convertFromBase(SpecificWeightKilonewtonPerCubicMillimeter)
	a.kilonewtons_per_cubic_millimeterLazy = &kilonewtons_per_cubic_millimeter
	return kilonewtons_per_cubic_millimeter
}

// KilonewtonsPerCubicCentimeter returns the SpecificWeight value in KilonewtonsPerCubicCentimeter.
//
// 
func (a *SpecificWeight) KilonewtonsPerCubicCentimeter() float64 {
	if a.kilonewtons_per_cubic_centimeterLazy != nil {
		return *a.kilonewtons_per_cubic_centimeterLazy
	}
	kilonewtons_per_cubic_centimeter := a.convertFromBase(SpecificWeightKilonewtonPerCubicCentimeter)
	a.kilonewtons_per_cubic_centimeterLazy = &kilonewtons_per_cubic_centimeter
	return kilonewtons_per_cubic_centimeter
}

// KilonewtonsPerCubicMeter returns the SpecificWeight value in KilonewtonsPerCubicMeter.
//
// 
func (a *SpecificWeight) KilonewtonsPerCubicMeter() float64 {
	if a.kilonewtons_per_cubic_meterLazy != nil {
		return *a.kilonewtons_per_cubic_meterLazy
	}
	kilonewtons_per_cubic_meter := a.convertFromBase(SpecificWeightKilonewtonPerCubicMeter)
	a.kilonewtons_per_cubic_meterLazy = &kilonewtons_per_cubic_meter
	return kilonewtons_per_cubic_meter
}

// MeganewtonsPerCubicMeter returns the SpecificWeight value in MeganewtonsPerCubicMeter.
//
// 
func (a *SpecificWeight) MeganewtonsPerCubicMeter() float64 {
	if a.meganewtons_per_cubic_meterLazy != nil {
		return *a.meganewtons_per_cubic_meterLazy
	}
	meganewtons_per_cubic_meter := a.convertFromBase(SpecificWeightMeganewtonPerCubicMeter)
	a.meganewtons_per_cubic_meterLazy = &meganewtons_per_cubic_meter
	return meganewtons_per_cubic_meter
}

// KilopoundsForcePerCubicInch returns the SpecificWeight value in KilopoundsForcePerCubicInch.
//
// 
func (a *SpecificWeight) KilopoundsForcePerCubicInch() float64 {
	if a.kilopounds_force_per_cubic_inchLazy != nil {
		return *a.kilopounds_force_per_cubic_inchLazy
	}
	kilopounds_force_per_cubic_inch := a.convertFromBase(SpecificWeightKilopoundForcePerCubicInch)
	a.kilopounds_force_per_cubic_inchLazy = &kilopounds_force_per_cubic_inch
	return kilopounds_force_per_cubic_inch
}

// KilopoundsForcePerCubicFoot returns the SpecificWeight value in KilopoundsForcePerCubicFoot.
//
// 
func (a *SpecificWeight) KilopoundsForcePerCubicFoot() float64 {
	if a.kilopounds_force_per_cubic_footLazy != nil {
		return *a.kilopounds_force_per_cubic_footLazy
	}
	kilopounds_force_per_cubic_foot := a.convertFromBase(SpecificWeightKilopoundForcePerCubicFoot)
	a.kilopounds_force_per_cubic_footLazy = &kilopounds_force_per_cubic_foot
	return kilopounds_force_per_cubic_foot
}


// ToDto creates a SpecificWeightDto representation from the SpecificWeight instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerCubicMeter by default.
func (a *SpecificWeight) ToDto(holdInUnit *SpecificWeightUnits) SpecificWeightDto {
	if holdInUnit == nil {
		defaultUnit := SpecificWeightNewtonPerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificWeightDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the SpecificWeight instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerCubicMeter by default.
func (a *SpecificWeight) ToDtoJSON(holdInUnit *SpecificWeightUnits) ([]byte, error) {
	// Convert to SpecificWeightDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SpecificWeight to a specific unit value.
// The function uses the provided unit type (SpecificWeightUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *SpecificWeight) Convert(toUnit SpecificWeightUnits) float64 {
	switch toUnit { 
    case SpecificWeightNewtonPerCubicMillimeter:
		return a.NewtonsPerCubicMillimeter()
    case SpecificWeightNewtonPerCubicCentimeter:
		return a.NewtonsPerCubicCentimeter()
    case SpecificWeightNewtonPerCubicMeter:
		return a.NewtonsPerCubicMeter()
    case SpecificWeightKilogramForcePerCubicMillimeter:
		return a.KilogramsForcePerCubicMillimeter()
    case SpecificWeightKilogramForcePerCubicCentimeter:
		return a.KilogramsForcePerCubicCentimeter()
    case SpecificWeightKilogramForcePerCubicMeter:
		return a.KilogramsForcePerCubicMeter()
    case SpecificWeightPoundForcePerCubicInch:
		return a.PoundsForcePerCubicInch()
    case SpecificWeightPoundForcePerCubicFoot:
		return a.PoundsForcePerCubicFoot()
    case SpecificWeightTonneForcePerCubicMillimeter:
		return a.TonnesForcePerCubicMillimeter()
    case SpecificWeightTonneForcePerCubicCentimeter:
		return a.TonnesForcePerCubicCentimeter()
    case SpecificWeightTonneForcePerCubicMeter:
		return a.TonnesForcePerCubicMeter()
    case SpecificWeightKilonewtonPerCubicMillimeter:
		return a.KilonewtonsPerCubicMillimeter()
    case SpecificWeightKilonewtonPerCubicCentimeter:
		return a.KilonewtonsPerCubicCentimeter()
    case SpecificWeightKilonewtonPerCubicMeter:
		return a.KilonewtonsPerCubicMeter()
    case SpecificWeightMeganewtonPerCubicMeter:
		return a.MeganewtonsPerCubicMeter()
    case SpecificWeightKilopoundForcePerCubicInch:
		return a.KilopoundsForcePerCubicInch()
    case SpecificWeightKilopoundForcePerCubicFoot:
		return a.KilopoundsForcePerCubicFoot()
	default:
		return math.NaN()
	}
}

func (a *SpecificWeight) convertFromBase(toUnit SpecificWeightUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificWeightNewtonPerCubicMillimeter:
		return (value * 0.000000001) 
	case SpecificWeightNewtonPerCubicCentimeter:
		return (value * 0.000001) 
	case SpecificWeightNewtonPerCubicMeter:
		return (value) 
	case SpecificWeightKilogramForcePerCubicMillimeter:
		return (value / 9.80665e9) 
	case SpecificWeightKilogramForcePerCubicCentimeter:
		return (value / 9.80665e6) 
	case SpecificWeightKilogramForcePerCubicMeter:
		return (value / 9.80665) 
	case SpecificWeightPoundForcePerCubicInch:
		return (value * 1.6387064e-5 / 4.4482216152605) 
	case SpecificWeightPoundForcePerCubicFoot:
		return (value * 0.028316846592 / 4.4482216152605) 
	case SpecificWeightTonneForcePerCubicMillimeter:
		return (value / 9.80665e12) 
	case SpecificWeightTonneForcePerCubicCentimeter:
		return (value / 9.80665e9) 
	case SpecificWeightTonneForcePerCubicMeter:
		return (value / 9.80665e3) 
	case SpecificWeightKilonewtonPerCubicMillimeter:
		return ((value * 0.000000001) / 1000.0) 
	case SpecificWeightKilonewtonPerCubicCentimeter:
		return ((value * 0.000001) / 1000.0) 
	case SpecificWeightKilonewtonPerCubicMeter:
		return ((value) / 1000.0) 
	case SpecificWeightMeganewtonPerCubicMeter:
		return ((value) / 1000000.0) 
	case SpecificWeightKilopoundForcePerCubicInch:
		return ((value * 1.6387064e-5 / 4.4482216152605) / 1000.0) 
	case SpecificWeightKilopoundForcePerCubicFoot:
		return ((value * 0.028316846592 / 4.4482216152605) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *SpecificWeight) convertToBase(value float64, fromUnit SpecificWeightUnits) float64 {
	switch fromUnit { 
	case SpecificWeightNewtonPerCubicMillimeter:
		return (value * 1000000000) 
	case SpecificWeightNewtonPerCubicCentimeter:
		return (value * 1000000) 
	case SpecificWeightNewtonPerCubicMeter:
		return (value) 
	case SpecificWeightKilogramForcePerCubicMillimeter:
		return (value * 9.80665e9) 
	case SpecificWeightKilogramForcePerCubicCentimeter:
		return (value * 9.80665e6) 
	case SpecificWeightKilogramForcePerCubicMeter:
		return (value * 9.80665) 
	case SpecificWeightPoundForcePerCubicInch:
		return (value * 4.4482216152605 / 1.6387064e-5) 
	case SpecificWeightPoundForcePerCubicFoot:
		return (value * 4.4482216152605 / 0.028316846592) 
	case SpecificWeightTonneForcePerCubicMillimeter:
		return (value * 9.80665e12) 
	case SpecificWeightTonneForcePerCubicCentimeter:
		return (value * 9.80665e9) 
	case SpecificWeightTonneForcePerCubicMeter:
		return (value * 9.80665e3) 
	case SpecificWeightKilonewtonPerCubicMillimeter:
		return ((value * 1000000000) * 1000.0) 
	case SpecificWeightKilonewtonPerCubicCentimeter:
		return ((value * 1000000) * 1000.0) 
	case SpecificWeightKilonewtonPerCubicMeter:
		return ((value) * 1000.0) 
	case SpecificWeightMeganewtonPerCubicMeter:
		return ((value) * 1000000.0) 
	case SpecificWeightKilopoundForcePerCubicInch:
		return ((value * 4.4482216152605 / 1.6387064e-5) * 1000.0) 
	case SpecificWeightKilopoundForcePerCubicFoot:
		return ((value * 4.4482216152605 / 0.028316846592) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the SpecificWeight in the default unit (NewtonPerCubicMeter),
// formatted to two decimal places.
func (a SpecificWeight) String() string {
	return a.ToString(SpecificWeightNewtonPerCubicMeter, 2)
}

// ToString formats the SpecificWeight value as a string with the specified unit and fractional digits.
// It converts the SpecificWeight to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SpecificWeight value will be converted (e.g., NewtonPerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SpecificWeight with the unit abbreviation.
func (a *SpecificWeight) ToString(unit SpecificWeightUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSpecificWeightAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSpecificWeightAbbreviation(unit))
}

// Equals checks if the given SpecificWeight is equal to the current SpecificWeight.
//
// Parameters:
//    other: The SpecificWeight to compare against.
//
// Returns:
//    bool: Returns true if both SpecificWeight are equal, false otherwise.
func (a *SpecificWeight) Equals(other *SpecificWeight) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SpecificWeight with another SpecificWeight.
// It returns -1 if the current SpecificWeight is less than the other SpecificWeight, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SpecificWeight to compare against.
//
// Returns:
//    int: -1 if the current SpecificWeight is less, 1 if greater, and 0 if equal.
func (a *SpecificWeight) CompareTo(other *SpecificWeight) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given SpecificWeight to the current SpecificWeight and returns the result.
// The result is a new SpecificWeight instance with the sum of the values.
//
// Parameters:
//    other: The SpecificWeight to add to the current SpecificWeight.
//
// Returns:
//    *SpecificWeight: A new SpecificWeight instance representing the sum of both SpecificWeight.
func (a *SpecificWeight) Add(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SpecificWeight from the current SpecificWeight and returns the result.
// The result is a new SpecificWeight instance with the difference of the values.
//
// Parameters:
//    other: The SpecificWeight to subtract from the current SpecificWeight.
//
// Returns:
//    *SpecificWeight: A new SpecificWeight instance representing the difference of both SpecificWeight.
func (a *SpecificWeight) Subtract(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SpecificWeight by the given SpecificWeight and returns the result.
// The result is a new SpecificWeight instance with the product of the values.
//
// Parameters:
//    other: The SpecificWeight to multiply with the current SpecificWeight.
//
// Returns:
//    *SpecificWeight: A new SpecificWeight instance representing the product of both SpecificWeight.
func (a *SpecificWeight) Multiply(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value * other.BaseValue()}
}

// Divide divides the current SpecificWeight by the given SpecificWeight and returns the result.
// The result is a new SpecificWeight instance with the quotient of the values.
//
// Parameters:
//    other: The SpecificWeight to divide the current SpecificWeight by.
//
// Returns:
//    *SpecificWeight: A new SpecificWeight instance representing the quotient of both SpecificWeight.
func (a *SpecificWeight) Divide(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value / other.BaseValue()}
}

// GetSpecificWeightAbbreviation gets the unit abbreviation.
func GetSpecificWeightAbbreviation(unit SpecificWeightUnits) string {
	switch unit { 
	case SpecificWeightNewtonPerCubicMillimeter:
		return "N/mm³" 
	case SpecificWeightNewtonPerCubicCentimeter:
		return "N/cm³" 
	case SpecificWeightNewtonPerCubicMeter:
		return "N/m³" 
	case SpecificWeightKilogramForcePerCubicMillimeter:
		return "kgf/mm³" 
	case SpecificWeightKilogramForcePerCubicCentimeter:
		return "kgf/cm³" 
	case SpecificWeightKilogramForcePerCubicMeter:
		return "kgf/m³" 
	case SpecificWeightPoundForcePerCubicInch:
		return "lbf/in³" 
	case SpecificWeightPoundForcePerCubicFoot:
		return "lbf/ft³" 
	case SpecificWeightTonneForcePerCubicMillimeter:
		return "tf/mm³" 
	case SpecificWeightTonneForcePerCubicCentimeter:
		return "tf/cm³" 
	case SpecificWeightTonneForcePerCubicMeter:
		return "tf/m³" 
	case SpecificWeightKilonewtonPerCubicMillimeter:
		return "kN/mm³" 
	case SpecificWeightKilonewtonPerCubicCentimeter:
		return "kN/cm³" 
	case SpecificWeightKilonewtonPerCubicMeter:
		return "kN/m³" 
	case SpecificWeightMeganewtonPerCubicMeter:
		return "MN/m³" 
	case SpecificWeightKilopoundForcePerCubicInch:
		return "klbf/in³" 
	case SpecificWeightKilopoundForcePerCubicFoot:
		return "klbf/ft³" 
	default:
		return ""
	}
}