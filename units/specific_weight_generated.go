// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificWeightUnits enumeration
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

// SpecificWeightDto represents an SpecificWeight
type SpecificWeightDto struct {
	Value float64
	Unit  SpecificWeightUnits
}

// SpecificWeightDtoFactory struct to group related functions
type SpecificWeightDtoFactory struct{}

func (udf SpecificWeightDtoFactory) FromJSON(data []byte) (*SpecificWeightDto, error) {
	a := SpecificWeightDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpecificWeightDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SpecificWeight struct
type SpecificWeight struct {
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

// SpecificWeightFactory struct to group related functions
type SpecificWeightFactory struct{}

func (uf SpecificWeightFactory) CreateSpecificWeight(value float64, unit SpecificWeightUnits) (*SpecificWeight, error) {
	return newSpecificWeight(value, unit)
}

func (uf SpecificWeightFactory) FromDto(dto SpecificWeightDto) (*SpecificWeight, error) {
	return newSpecificWeight(dto.Value, dto.Unit)
}

func (uf SpecificWeightFactory) FromDtoJSON(data []byte) (*SpecificWeight, error) {
	unitDto, err := SpecificWeightDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpecificWeightFactory{}.FromDto(*unitDto)
}


// FromNewtonPerCubicMillimeter creates a new SpecificWeight instance from NewtonPerCubicMillimeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicMillimeter)
}

// FromNewtonPerCubicCentimeter creates a new SpecificWeight instance from NewtonPerCubicCentimeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicCentimeter)
}

// FromNewtonPerCubicMeter creates a new SpecificWeight instance from NewtonPerCubicMeter.
func (uf SpecificWeightFactory) FromNewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightNewtonPerCubicMeter)
}

// FromKilogramForcePerCubicMillimeter creates a new SpecificWeight instance from KilogramForcePerCubicMillimeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicMillimeter)
}

// FromKilogramForcePerCubicCentimeter creates a new SpecificWeight instance from KilogramForcePerCubicCentimeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicCentimeter)
}

// FromKilogramForcePerCubicMeter creates a new SpecificWeight instance from KilogramForcePerCubicMeter.
func (uf SpecificWeightFactory) FromKilogramsForcePerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilogramForcePerCubicMeter)
}

// FromPoundForcePerCubicInch creates a new SpecificWeight instance from PoundForcePerCubicInch.
func (uf SpecificWeightFactory) FromPoundsForcePerCubicInch(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightPoundForcePerCubicInch)
}

// FromPoundForcePerCubicFoot creates a new SpecificWeight instance from PoundForcePerCubicFoot.
func (uf SpecificWeightFactory) FromPoundsForcePerCubicFoot(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightPoundForcePerCubicFoot)
}

// FromTonneForcePerCubicMillimeter creates a new SpecificWeight instance from TonneForcePerCubicMillimeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicMillimeter)
}

// FromTonneForcePerCubicCentimeter creates a new SpecificWeight instance from TonneForcePerCubicCentimeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicCentimeter)
}

// FromTonneForcePerCubicMeter creates a new SpecificWeight instance from TonneForcePerCubicMeter.
func (uf SpecificWeightFactory) FromTonnesForcePerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightTonneForcePerCubicMeter)
}

// FromKilonewtonPerCubicMillimeter creates a new SpecificWeight instance from KilonewtonPerCubicMillimeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicMillimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicMillimeter)
}

// FromKilonewtonPerCubicCentimeter creates a new SpecificWeight instance from KilonewtonPerCubicCentimeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicCentimeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicCentimeter)
}

// FromKilonewtonPerCubicMeter creates a new SpecificWeight instance from KilonewtonPerCubicMeter.
func (uf SpecificWeightFactory) FromKilonewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilonewtonPerCubicMeter)
}

// FromMeganewtonPerCubicMeter creates a new SpecificWeight instance from MeganewtonPerCubicMeter.
func (uf SpecificWeightFactory) FromMeganewtonsPerCubicMeter(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightMeganewtonPerCubicMeter)
}

// FromKilopoundForcePerCubicInch creates a new SpecificWeight instance from KilopoundForcePerCubicInch.
func (uf SpecificWeightFactory) FromKilopoundsForcePerCubicInch(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilopoundForcePerCubicInch)
}

// FromKilopoundForcePerCubicFoot creates a new SpecificWeight instance from KilopoundForcePerCubicFoot.
func (uf SpecificWeightFactory) FromKilopoundsForcePerCubicFoot(value float64) (*SpecificWeight, error) {
	return newSpecificWeight(value, SpecificWeightKilopoundForcePerCubicFoot)
}




// newSpecificWeight creates a new SpecificWeight.
func newSpecificWeight(value float64, fromUnit SpecificWeightUnits) (*SpecificWeight, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificWeight{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificWeight in NewtonPerCubicMeter.
func (a *SpecificWeight) BaseValue() float64 {
	return a.value
}


// NewtonPerCubicMillimeter returns the value in NewtonPerCubicMillimeter.
func (a *SpecificWeight) NewtonsPerCubicMillimeter() float64 {
	if a.newtons_per_cubic_millimeterLazy != nil {
		return *a.newtons_per_cubic_millimeterLazy
	}
	newtons_per_cubic_millimeter := a.convertFromBase(SpecificWeightNewtonPerCubicMillimeter)
	a.newtons_per_cubic_millimeterLazy = &newtons_per_cubic_millimeter
	return newtons_per_cubic_millimeter
}

// NewtonPerCubicCentimeter returns the value in NewtonPerCubicCentimeter.
func (a *SpecificWeight) NewtonsPerCubicCentimeter() float64 {
	if a.newtons_per_cubic_centimeterLazy != nil {
		return *a.newtons_per_cubic_centimeterLazy
	}
	newtons_per_cubic_centimeter := a.convertFromBase(SpecificWeightNewtonPerCubicCentimeter)
	a.newtons_per_cubic_centimeterLazy = &newtons_per_cubic_centimeter
	return newtons_per_cubic_centimeter
}

// NewtonPerCubicMeter returns the value in NewtonPerCubicMeter.
func (a *SpecificWeight) NewtonsPerCubicMeter() float64 {
	if a.newtons_per_cubic_meterLazy != nil {
		return *a.newtons_per_cubic_meterLazy
	}
	newtons_per_cubic_meter := a.convertFromBase(SpecificWeightNewtonPerCubicMeter)
	a.newtons_per_cubic_meterLazy = &newtons_per_cubic_meter
	return newtons_per_cubic_meter
}

// KilogramForcePerCubicMillimeter returns the value in KilogramForcePerCubicMillimeter.
func (a *SpecificWeight) KilogramsForcePerCubicMillimeter() float64 {
	if a.kilograms_force_per_cubic_millimeterLazy != nil {
		return *a.kilograms_force_per_cubic_millimeterLazy
	}
	kilograms_force_per_cubic_millimeter := a.convertFromBase(SpecificWeightKilogramForcePerCubicMillimeter)
	a.kilograms_force_per_cubic_millimeterLazy = &kilograms_force_per_cubic_millimeter
	return kilograms_force_per_cubic_millimeter
}

// KilogramForcePerCubicCentimeter returns the value in KilogramForcePerCubicCentimeter.
func (a *SpecificWeight) KilogramsForcePerCubicCentimeter() float64 {
	if a.kilograms_force_per_cubic_centimeterLazy != nil {
		return *a.kilograms_force_per_cubic_centimeterLazy
	}
	kilograms_force_per_cubic_centimeter := a.convertFromBase(SpecificWeightKilogramForcePerCubicCentimeter)
	a.kilograms_force_per_cubic_centimeterLazy = &kilograms_force_per_cubic_centimeter
	return kilograms_force_per_cubic_centimeter
}

// KilogramForcePerCubicMeter returns the value in KilogramForcePerCubicMeter.
func (a *SpecificWeight) KilogramsForcePerCubicMeter() float64 {
	if a.kilograms_force_per_cubic_meterLazy != nil {
		return *a.kilograms_force_per_cubic_meterLazy
	}
	kilograms_force_per_cubic_meter := a.convertFromBase(SpecificWeightKilogramForcePerCubicMeter)
	a.kilograms_force_per_cubic_meterLazy = &kilograms_force_per_cubic_meter
	return kilograms_force_per_cubic_meter
}

// PoundForcePerCubicInch returns the value in PoundForcePerCubicInch.
func (a *SpecificWeight) PoundsForcePerCubicInch() float64 {
	if a.pounds_force_per_cubic_inchLazy != nil {
		return *a.pounds_force_per_cubic_inchLazy
	}
	pounds_force_per_cubic_inch := a.convertFromBase(SpecificWeightPoundForcePerCubicInch)
	a.pounds_force_per_cubic_inchLazy = &pounds_force_per_cubic_inch
	return pounds_force_per_cubic_inch
}

// PoundForcePerCubicFoot returns the value in PoundForcePerCubicFoot.
func (a *SpecificWeight) PoundsForcePerCubicFoot() float64 {
	if a.pounds_force_per_cubic_footLazy != nil {
		return *a.pounds_force_per_cubic_footLazy
	}
	pounds_force_per_cubic_foot := a.convertFromBase(SpecificWeightPoundForcePerCubicFoot)
	a.pounds_force_per_cubic_footLazy = &pounds_force_per_cubic_foot
	return pounds_force_per_cubic_foot
}

// TonneForcePerCubicMillimeter returns the value in TonneForcePerCubicMillimeter.
func (a *SpecificWeight) TonnesForcePerCubicMillimeter() float64 {
	if a.tonnes_force_per_cubic_millimeterLazy != nil {
		return *a.tonnes_force_per_cubic_millimeterLazy
	}
	tonnes_force_per_cubic_millimeter := a.convertFromBase(SpecificWeightTonneForcePerCubicMillimeter)
	a.tonnes_force_per_cubic_millimeterLazy = &tonnes_force_per_cubic_millimeter
	return tonnes_force_per_cubic_millimeter
}

// TonneForcePerCubicCentimeter returns the value in TonneForcePerCubicCentimeter.
func (a *SpecificWeight) TonnesForcePerCubicCentimeter() float64 {
	if a.tonnes_force_per_cubic_centimeterLazy != nil {
		return *a.tonnes_force_per_cubic_centimeterLazy
	}
	tonnes_force_per_cubic_centimeter := a.convertFromBase(SpecificWeightTonneForcePerCubicCentimeter)
	a.tonnes_force_per_cubic_centimeterLazy = &tonnes_force_per_cubic_centimeter
	return tonnes_force_per_cubic_centimeter
}

// TonneForcePerCubicMeter returns the value in TonneForcePerCubicMeter.
func (a *SpecificWeight) TonnesForcePerCubicMeter() float64 {
	if a.tonnes_force_per_cubic_meterLazy != nil {
		return *a.tonnes_force_per_cubic_meterLazy
	}
	tonnes_force_per_cubic_meter := a.convertFromBase(SpecificWeightTonneForcePerCubicMeter)
	a.tonnes_force_per_cubic_meterLazy = &tonnes_force_per_cubic_meter
	return tonnes_force_per_cubic_meter
}

// KilonewtonPerCubicMillimeter returns the value in KilonewtonPerCubicMillimeter.
func (a *SpecificWeight) KilonewtonsPerCubicMillimeter() float64 {
	if a.kilonewtons_per_cubic_millimeterLazy != nil {
		return *a.kilonewtons_per_cubic_millimeterLazy
	}
	kilonewtons_per_cubic_millimeter := a.convertFromBase(SpecificWeightKilonewtonPerCubicMillimeter)
	a.kilonewtons_per_cubic_millimeterLazy = &kilonewtons_per_cubic_millimeter
	return kilonewtons_per_cubic_millimeter
}

// KilonewtonPerCubicCentimeter returns the value in KilonewtonPerCubicCentimeter.
func (a *SpecificWeight) KilonewtonsPerCubicCentimeter() float64 {
	if a.kilonewtons_per_cubic_centimeterLazy != nil {
		return *a.kilonewtons_per_cubic_centimeterLazy
	}
	kilonewtons_per_cubic_centimeter := a.convertFromBase(SpecificWeightKilonewtonPerCubicCentimeter)
	a.kilonewtons_per_cubic_centimeterLazy = &kilonewtons_per_cubic_centimeter
	return kilonewtons_per_cubic_centimeter
}

// KilonewtonPerCubicMeter returns the value in KilonewtonPerCubicMeter.
func (a *SpecificWeight) KilonewtonsPerCubicMeter() float64 {
	if a.kilonewtons_per_cubic_meterLazy != nil {
		return *a.kilonewtons_per_cubic_meterLazy
	}
	kilonewtons_per_cubic_meter := a.convertFromBase(SpecificWeightKilonewtonPerCubicMeter)
	a.kilonewtons_per_cubic_meterLazy = &kilonewtons_per_cubic_meter
	return kilonewtons_per_cubic_meter
}

// MeganewtonPerCubicMeter returns the value in MeganewtonPerCubicMeter.
func (a *SpecificWeight) MeganewtonsPerCubicMeter() float64 {
	if a.meganewtons_per_cubic_meterLazy != nil {
		return *a.meganewtons_per_cubic_meterLazy
	}
	meganewtons_per_cubic_meter := a.convertFromBase(SpecificWeightMeganewtonPerCubicMeter)
	a.meganewtons_per_cubic_meterLazy = &meganewtons_per_cubic_meter
	return meganewtons_per_cubic_meter
}

// KilopoundForcePerCubicInch returns the value in KilopoundForcePerCubicInch.
func (a *SpecificWeight) KilopoundsForcePerCubicInch() float64 {
	if a.kilopounds_force_per_cubic_inchLazy != nil {
		return *a.kilopounds_force_per_cubic_inchLazy
	}
	kilopounds_force_per_cubic_inch := a.convertFromBase(SpecificWeightKilopoundForcePerCubicInch)
	a.kilopounds_force_per_cubic_inchLazy = &kilopounds_force_per_cubic_inch
	return kilopounds_force_per_cubic_inch
}

// KilopoundForcePerCubicFoot returns the value in KilopoundForcePerCubicFoot.
func (a *SpecificWeight) KilopoundsForcePerCubicFoot() float64 {
	if a.kilopounds_force_per_cubic_footLazy != nil {
		return *a.kilopounds_force_per_cubic_footLazy
	}
	kilopounds_force_per_cubic_foot := a.convertFromBase(SpecificWeightKilopoundForcePerCubicFoot)
	a.kilopounds_force_per_cubic_footLazy = &kilopounds_force_per_cubic_foot
	return kilopounds_force_per_cubic_foot
}


// ToDto creates an SpecificWeightDto representation.
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

// ToDtoJSON creates an SpecificWeightDto representation.
func (a *SpecificWeight) ToDtoJSON(holdInUnit *SpecificWeightUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SpecificWeight to a specific unit value.
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
		return 0
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
		return (value / 2.714471375263134e5) 
	case SpecificWeightPoundForcePerCubicFoot:
		return (value / 1.570874638462462e2) 
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
		return ((value / 2.714471375263134e5) / 1000.0) 
	case SpecificWeightKilopoundForcePerCubicFoot:
		return ((value / 1.570874638462462e2) / 1000.0) 
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
		return (value * 2.714471375263134e5) 
	case SpecificWeightPoundForcePerCubicFoot:
		return (value * 1.570874638462462e2) 
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
		return ((value * 2.714471375263134e5) * 1000.0) 
	case SpecificWeightKilopoundForcePerCubicFoot:
		return ((value * 1.570874638462462e2) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a SpecificWeight) String() string {
	return a.ToString(SpecificWeightNewtonPerCubicMeter, 2)
}

// ToString formats the SpecificWeight to string.
// fractionalDigits -1 for not mention
func (a *SpecificWeight) ToString(unit SpecificWeightUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SpecificWeight) getUnitAbbreviation(unit SpecificWeightUnits) string {
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

// Check if the given SpecificWeight are equals to the current SpecificWeight
func (a *SpecificWeight) Equals(other *SpecificWeight) bool {
	return a.value == other.BaseValue()
}

// Check if the given SpecificWeight are equals to the current SpecificWeight
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

// Add the given SpecificWeight to the current SpecificWeight.
func (a *SpecificWeight) Add(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value + other.BaseValue()}
}

// Subtract the given SpecificWeight to the current SpecificWeight.
func (a *SpecificWeight) Subtract(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value - other.BaseValue()}
}

// Multiply the given SpecificWeight to the current SpecificWeight.
func (a *SpecificWeight) Multiply(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value * other.BaseValue()}
}

// Divide the given SpecificWeight to the current SpecificWeight.
func (a *SpecificWeight) Divide(other *SpecificWeight) *SpecificWeight {
	return &SpecificWeight{value: a.value / other.BaseValue()}
}