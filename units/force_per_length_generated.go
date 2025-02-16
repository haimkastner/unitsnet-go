// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ForcePerLengthUnits defines various units of ForcePerLength.
type ForcePerLengthUnits string

const (
    
        // 
        ForcePerLengthNewtonPerMeter ForcePerLengthUnits = "NewtonPerMeter"
        // 
        ForcePerLengthNewtonPerCentimeter ForcePerLengthUnits = "NewtonPerCentimeter"
        // 
        ForcePerLengthNewtonPerMillimeter ForcePerLengthUnits = "NewtonPerMillimeter"
        // 
        ForcePerLengthKilogramForcePerMeter ForcePerLengthUnits = "KilogramForcePerMeter"
        // 
        ForcePerLengthKilogramForcePerCentimeter ForcePerLengthUnits = "KilogramForcePerCentimeter"
        // 
        ForcePerLengthKilogramForcePerMillimeter ForcePerLengthUnits = "KilogramForcePerMillimeter"
        // 
        ForcePerLengthTonneForcePerMeter ForcePerLengthUnits = "TonneForcePerMeter"
        // 
        ForcePerLengthTonneForcePerCentimeter ForcePerLengthUnits = "TonneForcePerCentimeter"
        // 
        ForcePerLengthTonneForcePerMillimeter ForcePerLengthUnits = "TonneForcePerMillimeter"
        // 
        ForcePerLengthPoundForcePerFoot ForcePerLengthUnits = "PoundForcePerFoot"
        // 
        ForcePerLengthPoundForcePerInch ForcePerLengthUnits = "PoundForcePerInch"
        // 
        ForcePerLengthPoundForcePerYard ForcePerLengthUnits = "PoundForcePerYard"
        // 
        ForcePerLengthKilopoundForcePerFoot ForcePerLengthUnits = "KilopoundForcePerFoot"
        // 
        ForcePerLengthKilopoundForcePerInch ForcePerLengthUnits = "KilopoundForcePerInch"
        // 
        ForcePerLengthNanonewtonPerMeter ForcePerLengthUnits = "NanonewtonPerMeter"
        // 
        ForcePerLengthMicronewtonPerMeter ForcePerLengthUnits = "MicronewtonPerMeter"
        // 
        ForcePerLengthMillinewtonPerMeter ForcePerLengthUnits = "MillinewtonPerMeter"
        // 
        ForcePerLengthCentinewtonPerMeter ForcePerLengthUnits = "CentinewtonPerMeter"
        // 
        ForcePerLengthDecinewtonPerMeter ForcePerLengthUnits = "DecinewtonPerMeter"
        // 
        ForcePerLengthDecanewtonPerMeter ForcePerLengthUnits = "DecanewtonPerMeter"
        // 
        ForcePerLengthKilonewtonPerMeter ForcePerLengthUnits = "KilonewtonPerMeter"
        // 
        ForcePerLengthMeganewtonPerMeter ForcePerLengthUnits = "MeganewtonPerMeter"
        // 
        ForcePerLengthNanonewtonPerCentimeter ForcePerLengthUnits = "NanonewtonPerCentimeter"
        // 
        ForcePerLengthMicronewtonPerCentimeter ForcePerLengthUnits = "MicronewtonPerCentimeter"
        // 
        ForcePerLengthMillinewtonPerCentimeter ForcePerLengthUnits = "MillinewtonPerCentimeter"
        // 
        ForcePerLengthCentinewtonPerCentimeter ForcePerLengthUnits = "CentinewtonPerCentimeter"
        // 
        ForcePerLengthDecinewtonPerCentimeter ForcePerLengthUnits = "DecinewtonPerCentimeter"
        // 
        ForcePerLengthDecanewtonPerCentimeter ForcePerLengthUnits = "DecanewtonPerCentimeter"
        // 
        ForcePerLengthKilonewtonPerCentimeter ForcePerLengthUnits = "KilonewtonPerCentimeter"
        // 
        ForcePerLengthMeganewtonPerCentimeter ForcePerLengthUnits = "MeganewtonPerCentimeter"
        // 
        ForcePerLengthNanonewtonPerMillimeter ForcePerLengthUnits = "NanonewtonPerMillimeter"
        // 
        ForcePerLengthMicronewtonPerMillimeter ForcePerLengthUnits = "MicronewtonPerMillimeter"
        // 
        ForcePerLengthMillinewtonPerMillimeter ForcePerLengthUnits = "MillinewtonPerMillimeter"
        // 
        ForcePerLengthCentinewtonPerMillimeter ForcePerLengthUnits = "CentinewtonPerMillimeter"
        // 
        ForcePerLengthDecinewtonPerMillimeter ForcePerLengthUnits = "DecinewtonPerMillimeter"
        // 
        ForcePerLengthDecanewtonPerMillimeter ForcePerLengthUnits = "DecanewtonPerMillimeter"
        // 
        ForcePerLengthKilonewtonPerMillimeter ForcePerLengthUnits = "KilonewtonPerMillimeter"
        // 
        ForcePerLengthMeganewtonPerMillimeter ForcePerLengthUnits = "MeganewtonPerMillimeter"
)

// ForcePerLengthDto represents a ForcePerLength measurement with a numerical value and its corresponding unit.
type ForcePerLengthDto struct {
    // Value is the numerical representation of the ForcePerLength.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ForcePerLength, as defined in the ForcePerLengthUnits enumeration.
	Unit  ForcePerLengthUnits `json:"unit"`
}

// ForcePerLengthDtoFactory groups methods for creating and serializing ForcePerLengthDto objects.
type ForcePerLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ForcePerLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ForcePerLengthDtoFactory) FromJSON(data []byte) (*ForcePerLengthDto, error) {
	a := ForcePerLengthDto{}

    // Parse JSON into ForcePerLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ForcePerLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ForcePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ForcePerLength represents a measurement in a of ForcePerLength.
//
// The magnitude of force per unit length.
type ForcePerLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    newtons_per_meterLazy *float64 
    newtons_per_centimeterLazy *float64 
    newtons_per_millimeterLazy *float64 
    kilograms_force_per_meterLazy *float64 
    kilograms_force_per_centimeterLazy *float64 
    kilograms_force_per_millimeterLazy *float64 
    tonnes_force_per_meterLazy *float64 
    tonnes_force_per_centimeterLazy *float64 
    tonnes_force_per_millimeterLazy *float64 
    pounds_force_per_footLazy *float64 
    pounds_force_per_inchLazy *float64 
    pounds_force_per_yardLazy *float64 
    kilopounds_force_per_footLazy *float64 
    kilopounds_force_per_inchLazy *float64 
    nanonewtons_per_meterLazy *float64 
    micronewtons_per_meterLazy *float64 
    millinewtons_per_meterLazy *float64 
    centinewtons_per_meterLazy *float64 
    decinewtons_per_meterLazy *float64 
    decanewtons_per_meterLazy *float64 
    kilonewtons_per_meterLazy *float64 
    meganewtons_per_meterLazy *float64 
    nanonewtons_per_centimeterLazy *float64 
    micronewtons_per_centimeterLazy *float64 
    millinewtons_per_centimeterLazy *float64 
    centinewtons_per_centimeterLazy *float64 
    decinewtons_per_centimeterLazy *float64 
    decanewtons_per_centimeterLazy *float64 
    kilonewtons_per_centimeterLazy *float64 
    meganewtons_per_centimeterLazy *float64 
    nanonewtons_per_millimeterLazy *float64 
    micronewtons_per_millimeterLazy *float64 
    millinewtons_per_millimeterLazy *float64 
    centinewtons_per_millimeterLazy *float64 
    decinewtons_per_millimeterLazy *float64 
    decanewtons_per_millimeterLazy *float64 
    kilonewtons_per_millimeterLazy *float64 
    meganewtons_per_millimeterLazy *float64 
}

// ForcePerLengthFactory groups methods for creating ForcePerLength instances.
type ForcePerLengthFactory struct{}

// CreateForcePerLength creates a new ForcePerLength instance from the given value and unit.
func (uf ForcePerLengthFactory) CreateForcePerLength(value float64, unit ForcePerLengthUnits) (*ForcePerLength, error) {
	return newForcePerLength(value, unit)
}

// FromDto converts a ForcePerLengthDto to a ForcePerLength instance.
func (uf ForcePerLengthFactory) FromDto(dto ForcePerLengthDto) (*ForcePerLength, error) {
	return newForcePerLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ForcePerLength instance.
func (uf ForcePerLengthFactory) FromDtoJSON(data []byte) (*ForcePerLength, error) {
	unitDto, err := ForcePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ForcePerLengthDto from JSON: %w", err)
	}
	return ForcePerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonsPerMeter creates a new ForcePerLength instance from a value in NewtonsPerMeter.
func (uf ForcePerLengthFactory) FromNewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerMeter)
}

// FromNewtonsPerCentimeter creates a new ForcePerLength instance from a value in NewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromNewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerCentimeter)
}

// FromNewtonsPerMillimeter creates a new ForcePerLength instance from a value in NewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromNewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerMillimeter)
}

// FromKilogramsForcePerMeter creates a new ForcePerLength instance from a value in KilogramsForcePerMeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerMeter)
}

// FromKilogramsForcePerCentimeter creates a new ForcePerLength instance from a value in KilogramsForcePerCentimeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerCentimeter)
}

// FromKilogramsForcePerMillimeter creates a new ForcePerLength instance from a value in KilogramsForcePerMillimeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerMillimeter)
}

// FromTonnesForcePerMeter creates a new ForcePerLength instance from a value in TonnesForcePerMeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerMeter)
}

// FromTonnesForcePerCentimeter creates a new ForcePerLength instance from a value in TonnesForcePerCentimeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerCentimeter)
}

// FromTonnesForcePerMillimeter creates a new ForcePerLength instance from a value in TonnesForcePerMillimeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerMillimeter)
}

// FromPoundsForcePerFoot creates a new ForcePerLength instance from a value in PoundsForcePerFoot.
func (uf ForcePerLengthFactory) FromPoundsForcePerFoot(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerFoot)
}

// FromPoundsForcePerInch creates a new ForcePerLength instance from a value in PoundsForcePerInch.
func (uf ForcePerLengthFactory) FromPoundsForcePerInch(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerInch)
}

// FromPoundsForcePerYard creates a new ForcePerLength instance from a value in PoundsForcePerYard.
func (uf ForcePerLengthFactory) FromPoundsForcePerYard(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerYard)
}

// FromKilopoundsForcePerFoot creates a new ForcePerLength instance from a value in KilopoundsForcePerFoot.
func (uf ForcePerLengthFactory) FromKilopoundsForcePerFoot(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilopoundForcePerFoot)
}

// FromKilopoundsForcePerInch creates a new ForcePerLength instance from a value in KilopoundsForcePerInch.
func (uf ForcePerLengthFactory) FromKilopoundsForcePerInch(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilopoundForcePerInch)
}

// FromNanonewtonsPerMeter creates a new ForcePerLength instance from a value in NanonewtonsPerMeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerMeter)
}

// FromMicronewtonsPerMeter creates a new ForcePerLength instance from a value in MicronewtonsPerMeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerMeter)
}

// FromMillinewtonsPerMeter creates a new ForcePerLength instance from a value in MillinewtonsPerMeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerMeter)
}

// FromCentinewtonsPerMeter creates a new ForcePerLength instance from a value in CentinewtonsPerMeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerMeter)
}

// FromDecinewtonsPerMeter creates a new ForcePerLength instance from a value in DecinewtonsPerMeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerMeter)
}

// FromDecanewtonsPerMeter creates a new ForcePerLength instance from a value in DecanewtonsPerMeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerMeter)
}

// FromKilonewtonsPerMeter creates a new ForcePerLength instance from a value in KilonewtonsPerMeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerMeter)
}

// FromMeganewtonsPerMeter creates a new ForcePerLength instance from a value in MeganewtonsPerMeter.
func (uf ForcePerLengthFactory) FromMeganewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMeganewtonPerMeter)
}

// FromNanonewtonsPerCentimeter creates a new ForcePerLength instance from a value in NanonewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerCentimeter)
}

// FromMicronewtonsPerCentimeter creates a new ForcePerLength instance from a value in MicronewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerCentimeter)
}

// FromMillinewtonsPerCentimeter creates a new ForcePerLength instance from a value in MillinewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerCentimeter)
}

// FromCentinewtonsPerCentimeter creates a new ForcePerLength instance from a value in CentinewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerCentimeter)
}

// FromDecinewtonsPerCentimeter creates a new ForcePerLength instance from a value in DecinewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerCentimeter)
}

// FromDecanewtonsPerCentimeter creates a new ForcePerLength instance from a value in DecanewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerCentimeter)
}

// FromKilonewtonsPerCentimeter creates a new ForcePerLength instance from a value in KilonewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerCentimeter)
}

// FromMeganewtonsPerCentimeter creates a new ForcePerLength instance from a value in MeganewtonsPerCentimeter.
func (uf ForcePerLengthFactory) FromMeganewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMeganewtonPerCentimeter)
}

// FromNanonewtonsPerMillimeter creates a new ForcePerLength instance from a value in NanonewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerMillimeter)
}

// FromMicronewtonsPerMillimeter creates a new ForcePerLength instance from a value in MicronewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerMillimeter)
}

// FromMillinewtonsPerMillimeter creates a new ForcePerLength instance from a value in MillinewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerMillimeter)
}

// FromCentinewtonsPerMillimeter creates a new ForcePerLength instance from a value in CentinewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerMillimeter)
}

// FromDecinewtonsPerMillimeter creates a new ForcePerLength instance from a value in DecinewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerMillimeter)
}

// FromDecanewtonsPerMillimeter creates a new ForcePerLength instance from a value in DecanewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerMillimeter)
}

// FromKilonewtonsPerMillimeter creates a new ForcePerLength instance from a value in KilonewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerMillimeter)
}

// FromMeganewtonsPerMillimeter creates a new ForcePerLength instance from a value in MeganewtonsPerMillimeter.
func (uf ForcePerLengthFactory) FromMeganewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMeganewtonPerMillimeter)
}


// newForcePerLength creates a new ForcePerLength.
func newForcePerLength(value float64, fromUnit ForcePerLengthUnits) (*ForcePerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ForcePerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ForcePerLength in NewtonPerMeter unit (the base/default quantity).
func (a *ForcePerLength) BaseValue() float64 {
	return a.value
}


// NewtonsPerMeter returns the ForcePerLength value in NewtonsPerMeter.
//
// 
func (a *ForcePerLength) NewtonsPerMeter() float64 {
	if a.newtons_per_meterLazy != nil {
		return *a.newtons_per_meterLazy
	}
	newtons_per_meter := a.convertFromBase(ForcePerLengthNewtonPerMeter)
	a.newtons_per_meterLazy = &newtons_per_meter
	return newtons_per_meter
}

// NewtonsPerCentimeter returns the ForcePerLength value in NewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) NewtonsPerCentimeter() float64 {
	if a.newtons_per_centimeterLazy != nil {
		return *a.newtons_per_centimeterLazy
	}
	newtons_per_centimeter := a.convertFromBase(ForcePerLengthNewtonPerCentimeter)
	a.newtons_per_centimeterLazy = &newtons_per_centimeter
	return newtons_per_centimeter
}

// NewtonsPerMillimeter returns the ForcePerLength value in NewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) NewtonsPerMillimeter() float64 {
	if a.newtons_per_millimeterLazy != nil {
		return *a.newtons_per_millimeterLazy
	}
	newtons_per_millimeter := a.convertFromBase(ForcePerLengthNewtonPerMillimeter)
	a.newtons_per_millimeterLazy = &newtons_per_millimeter
	return newtons_per_millimeter
}

// KilogramsForcePerMeter returns the ForcePerLength value in KilogramsForcePerMeter.
//
// 
func (a *ForcePerLength) KilogramsForcePerMeter() float64 {
	if a.kilograms_force_per_meterLazy != nil {
		return *a.kilograms_force_per_meterLazy
	}
	kilograms_force_per_meter := a.convertFromBase(ForcePerLengthKilogramForcePerMeter)
	a.kilograms_force_per_meterLazy = &kilograms_force_per_meter
	return kilograms_force_per_meter
}

// KilogramsForcePerCentimeter returns the ForcePerLength value in KilogramsForcePerCentimeter.
//
// 
func (a *ForcePerLength) KilogramsForcePerCentimeter() float64 {
	if a.kilograms_force_per_centimeterLazy != nil {
		return *a.kilograms_force_per_centimeterLazy
	}
	kilograms_force_per_centimeter := a.convertFromBase(ForcePerLengthKilogramForcePerCentimeter)
	a.kilograms_force_per_centimeterLazy = &kilograms_force_per_centimeter
	return kilograms_force_per_centimeter
}

// KilogramsForcePerMillimeter returns the ForcePerLength value in KilogramsForcePerMillimeter.
//
// 
func (a *ForcePerLength) KilogramsForcePerMillimeter() float64 {
	if a.kilograms_force_per_millimeterLazy != nil {
		return *a.kilograms_force_per_millimeterLazy
	}
	kilograms_force_per_millimeter := a.convertFromBase(ForcePerLengthKilogramForcePerMillimeter)
	a.kilograms_force_per_millimeterLazy = &kilograms_force_per_millimeter
	return kilograms_force_per_millimeter
}

// TonnesForcePerMeter returns the ForcePerLength value in TonnesForcePerMeter.
//
// 
func (a *ForcePerLength) TonnesForcePerMeter() float64 {
	if a.tonnes_force_per_meterLazy != nil {
		return *a.tonnes_force_per_meterLazy
	}
	tonnes_force_per_meter := a.convertFromBase(ForcePerLengthTonneForcePerMeter)
	a.tonnes_force_per_meterLazy = &tonnes_force_per_meter
	return tonnes_force_per_meter
}

// TonnesForcePerCentimeter returns the ForcePerLength value in TonnesForcePerCentimeter.
//
// 
func (a *ForcePerLength) TonnesForcePerCentimeter() float64 {
	if a.tonnes_force_per_centimeterLazy != nil {
		return *a.tonnes_force_per_centimeterLazy
	}
	tonnes_force_per_centimeter := a.convertFromBase(ForcePerLengthTonneForcePerCentimeter)
	a.tonnes_force_per_centimeterLazy = &tonnes_force_per_centimeter
	return tonnes_force_per_centimeter
}

// TonnesForcePerMillimeter returns the ForcePerLength value in TonnesForcePerMillimeter.
//
// 
func (a *ForcePerLength) TonnesForcePerMillimeter() float64 {
	if a.tonnes_force_per_millimeterLazy != nil {
		return *a.tonnes_force_per_millimeterLazy
	}
	tonnes_force_per_millimeter := a.convertFromBase(ForcePerLengthTonneForcePerMillimeter)
	a.tonnes_force_per_millimeterLazy = &tonnes_force_per_millimeter
	return tonnes_force_per_millimeter
}

// PoundsForcePerFoot returns the ForcePerLength value in PoundsForcePerFoot.
//
// 
func (a *ForcePerLength) PoundsForcePerFoot() float64 {
	if a.pounds_force_per_footLazy != nil {
		return *a.pounds_force_per_footLazy
	}
	pounds_force_per_foot := a.convertFromBase(ForcePerLengthPoundForcePerFoot)
	a.pounds_force_per_footLazy = &pounds_force_per_foot
	return pounds_force_per_foot
}

// PoundsForcePerInch returns the ForcePerLength value in PoundsForcePerInch.
//
// 
func (a *ForcePerLength) PoundsForcePerInch() float64 {
	if a.pounds_force_per_inchLazy != nil {
		return *a.pounds_force_per_inchLazy
	}
	pounds_force_per_inch := a.convertFromBase(ForcePerLengthPoundForcePerInch)
	a.pounds_force_per_inchLazy = &pounds_force_per_inch
	return pounds_force_per_inch
}

// PoundsForcePerYard returns the ForcePerLength value in PoundsForcePerYard.
//
// 
func (a *ForcePerLength) PoundsForcePerYard() float64 {
	if a.pounds_force_per_yardLazy != nil {
		return *a.pounds_force_per_yardLazy
	}
	pounds_force_per_yard := a.convertFromBase(ForcePerLengthPoundForcePerYard)
	a.pounds_force_per_yardLazy = &pounds_force_per_yard
	return pounds_force_per_yard
}

// KilopoundsForcePerFoot returns the ForcePerLength value in KilopoundsForcePerFoot.
//
// 
func (a *ForcePerLength) KilopoundsForcePerFoot() float64 {
	if a.kilopounds_force_per_footLazy != nil {
		return *a.kilopounds_force_per_footLazy
	}
	kilopounds_force_per_foot := a.convertFromBase(ForcePerLengthKilopoundForcePerFoot)
	a.kilopounds_force_per_footLazy = &kilopounds_force_per_foot
	return kilopounds_force_per_foot
}

// KilopoundsForcePerInch returns the ForcePerLength value in KilopoundsForcePerInch.
//
// 
func (a *ForcePerLength) KilopoundsForcePerInch() float64 {
	if a.kilopounds_force_per_inchLazy != nil {
		return *a.kilopounds_force_per_inchLazy
	}
	kilopounds_force_per_inch := a.convertFromBase(ForcePerLengthKilopoundForcePerInch)
	a.kilopounds_force_per_inchLazy = &kilopounds_force_per_inch
	return kilopounds_force_per_inch
}

// NanonewtonsPerMeter returns the ForcePerLength value in NanonewtonsPerMeter.
//
// 
func (a *ForcePerLength) NanonewtonsPerMeter() float64 {
	if a.nanonewtons_per_meterLazy != nil {
		return *a.nanonewtons_per_meterLazy
	}
	nanonewtons_per_meter := a.convertFromBase(ForcePerLengthNanonewtonPerMeter)
	a.nanonewtons_per_meterLazy = &nanonewtons_per_meter
	return nanonewtons_per_meter
}

// MicronewtonsPerMeter returns the ForcePerLength value in MicronewtonsPerMeter.
//
// 
func (a *ForcePerLength) MicronewtonsPerMeter() float64 {
	if a.micronewtons_per_meterLazy != nil {
		return *a.micronewtons_per_meterLazy
	}
	micronewtons_per_meter := a.convertFromBase(ForcePerLengthMicronewtonPerMeter)
	a.micronewtons_per_meterLazy = &micronewtons_per_meter
	return micronewtons_per_meter
}

// MillinewtonsPerMeter returns the ForcePerLength value in MillinewtonsPerMeter.
//
// 
func (a *ForcePerLength) MillinewtonsPerMeter() float64 {
	if a.millinewtons_per_meterLazy != nil {
		return *a.millinewtons_per_meterLazy
	}
	millinewtons_per_meter := a.convertFromBase(ForcePerLengthMillinewtonPerMeter)
	a.millinewtons_per_meterLazy = &millinewtons_per_meter
	return millinewtons_per_meter
}

// CentinewtonsPerMeter returns the ForcePerLength value in CentinewtonsPerMeter.
//
// 
func (a *ForcePerLength) CentinewtonsPerMeter() float64 {
	if a.centinewtons_per_meterLazy != nil {
		return *a.centinewtons_per_meterLazy
	}
	centinewtons_per_meter := a.convertFromBase(ForcePerLengthCentinewtonPerMeter)
	a.centinewtons_per_meterLazy = &centinewtons_per_meter
	return centinewtons_per_meter
}

// DecinewtonsPerMeter returns the ForcePerLength value in DecinewtonsPerMeter.
//
// 
func (a *ForcePerLength) DecinewtonsPerMeter() float64 {
	if a.decinewtons_per_meterLazy != nil {
		return *a.decinewtons_per_meterLazy
	}
	decinewtons_per_meter := a.convertFromBase(ForcePerLengthDecinewtonPerMeter)
	a.decinewtons_per_meterLazy = &decinewtons_per_meter
	return decinewtons_per_meter
}

// DecanewtonsPerMeter returns the ForcePerLength value in DecanewtonsPerMeter.
//
// 
func (a *ForcePerLength) DecanewtonsPerMeter() float64 {
	if a.decanewtons_per_meterLazy != nil {
		return *a.decanewtons_per_meterLazy
	}
	decanewtons_per_meter := a.convertFromBase(ForcePerLengthDecanewtonPerMeter)
	a.decanewtons_per_meterLazy = &decanewtons_per_meter
	return decanewtons_per_meter
}

// KilonewtonsPerMeter returns the ForcePerLength value in KilonewtonsPerMeter.
//
// 
func (a *ForcePerLength) KilonewtonsPerMeter() float64 {
	if a.kilonewtons_per_meterLazy != nil {
		return *a.kilonewtons_per_meterLazy
	}
	kilonewtons_per_meter := a.convertFromBase(ForcePerLengthKilonewtonPerMeter)
	a.kilonewtons_per_meterLazy = &kilonewtons_per_meter
	return kilonewtons_per_meter
}

// MeganewtonsPerMeter returns the ForcePerLength value in MeganewtonsPerMeter.
//
// 
func (a *ForcePerLength) MeganewtonsPerMeter() float64 {
	if a.meganewtons_per_meterLazy != nil {
		return *a.meganewtons_per_meterLazy
	}
	meganewtons_per_meter := a.convertFromBase(ForcePerLengthMeganewtonPerMeter)
	a.meganewtons_per_meterLazy = &meganewtons_per_meter
	return meganewtons_per_meter
}

// NanonewtonsPerCentimeter returns the ForcePerLength value in NanonewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) NanonewtonsPerCentimeter() float64 {
	if a.nanonewtons_per_centimeterLazy != nil {
		return *a.nanonewtons_per_centimeterLazy
	}
	nanonewtons_per_centimeter := a.convertFromBase(ForcePerLengthNanonewtonPerCentimeter)
	a.nanonewtons_per_centimeterLazy = &nanonewtons_per_centimeter
	return nanonewtons_per_centimeter
}

// MicronewtonsPerCentimeter returns the ForcePerLength value in MicronewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) MicronewtonsPerCentimeter() float64 {
	if a.micronewtons_per_centimeterLazy != nil {
		return *a.micronewtons_per_centimeterLazy
	}
	micronewtons_per_centimeter := a.convertFromBase(ForcePerLengthMicronewtonPerCentimeter)
	a.micronewtons_per_centimeterLazy = &micronewtons_per_centimeter
	return micronewtons_per_centimeter
}

// MillinewtonsPerCentimeter returns the ForcePerLength value in MillinewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) MillinewtonsPerCentimeter() float64 {
	if a.millinewtons_per_centimeterLazy != nil {
		return *a.millinewtons_per_centimeterLazy
	}
	millinewtons_per_centimeter := a.convertFromBase(ForcePerLengthMillinewtonPerCentimeter)
	a.millinewtons_per_centimeterLazy = &millinewtons_per_centimeter
	return millinewtons_per_centimeter
}

// CentinewtonsPerCentimeter returns the ForcePerLength value in CentinewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) CentinewtonsPerCentimeter() float64 {
	if a.centinewtons_per_centimeterLazy != nil {
		return *a.centinewtons_per_centimeterLazy
	}
	centinewtons_per_centimeter := a.convertFromBase(ForcePerLengthCentinewtonPerCentimeter)
	a.centinewtons_per_centimeterLazy = &centinewtons_per_centimeter
	return centinewtons_per_centimeter
}

// DecinewtonsPerCentimeter returns the ForcePerLength value in DecinewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) DecinewtonsPerCentimeter() float64 {
	if a.decinewtons_per_centimeterLazy != nil {
		return *a.decinewtons_per_centimeterLazy
	}
	decinewtons_per_centimeter := a.convertFromBase(ForcePerLengthDecinewtonPerCentimeter)
	a.decinewtons_per_centimeterLazy = &decinewtons_per_centimeter
	return decinewtons_per_centimeter
}

// DecanewtonsPerCentimeter returns the ForcePerLength value in DecanewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) DecanewtonsPerCentimeter() float64 {
	if a.decanewtons_per_centimeterLazy != nil {
		return *a.decanewtons_per_centimeterLazy
	}
	decanewtons_per_centimeter := a.convertFromBase(ForcePerLengthDecanewtonPerCentimeter)
	a.decanewtons_per_centimeterLazy = &decanewtons_per_centimeter
	return decanewtons_per_centimeter
}

// KilonewtonsPerCentimeter returns the ForcePerLength value in KilonewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) KilonewtonsPerCentimeter() float64 {
	if a.kilonewtons_per_centimeterLazy != nil {
		return *a.kilonewtons_per_centimeterLazy
	}
	kilonewtons_per_centimeter := a.convertFromBase(ForcePerLengthKilonewtonPerCentimeter)
	a.kilonewtons_per_centimeterLazy = &kilonewtons_per_centimeter
	return kilonewtons_per_centimeter
}

// MeganewtonsPerCentimeter returns the ForcePerLength value in MeganewtonsPerCentimeter.
//
// 
func (a *ForcePerLength) MeganewtonsPerCentimeter() float64 {
	if a.meganewtons_per_centimeterLazy != nil {
		return *a.meganewtons_per_centimeterLazy
	}
	meganewtons_per_centimeter := a.convertFromBase(ForcePerLengthMeganewtonPerCentimeter)
	a.meganewtons_per_centimeterLazy = &meganewtons_per_centimeter
	return meganewtons_per_centimeter
}

// NanonewtonsPerMillimeter returns the ForcePerLength value in NanonewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) NanonewtonsPerMillimeter() float64 {
	if a.nanonewtons_per_millimeterLazy != nil {
		return *a.nanonewtons_per_millimeterLazy
	}
	nanonewtons_per_millimeter := a.convertFromBase(ForcePerLengthNanonewtonPerMillimeter)
	a.nanonewtons_per_millimeterLazy = &nanonewtons_per_millimeter
	return nanonewtons_per_millimeter
}

// MicronewtonsPerMillimeter returns the ForcePerLength value in MicronewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) MicronewtonsPerMillimeter() float64 {
	if a.micronewtons_per_millimeterLazy != nil {
		return *a.micronewtons_per_millimeterLazy
	}
	micronewtons_per_millimeter := a.convertFromBase(ForcePerLengthMicronewtonPerMillimeter)
	a.micronewtons_per_millimeterLazy = &micronewtons_per_millimeter
	return micronewtons_per_millimeter
}

// MillinewtonsPerMillimeter returns the ForcePerLength value in MillinewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) MillinewtonsPerMillimeter() float64 {
	if a.millinewtons_per_millimeterLazy != nil {
		return *a.millinewtons_per_millimeterLazy
	}
	millinewtons_per_millimeter := a.convertFromBase(ForcePerLengthMillinewtonPerMillimeter)
	a.millinewtons_per_millimeterLazy = &millinewtons_per_millimeter
	return millinewtons_per_millimeter
}

// CentinewtonsPerMillimeter returns the ForcePerLength value in CentinewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) CentinewtonsPerMillimeter() float64 {
	if a.centinewtons_per_millimeterLazy != nil {
		return *a.centinewtons_per_millimeterLazy
	}
	centinewtons_per_millimeter := a.convertFromBase(ForcePerLengthCentinewtonPerMillimeter)
	a.centinewtons_per_millimeterLazy = &centinewtons_per_millimeter
	return centinewtons_per_millimeter
}

// DecinewtonsPerMillimeter returns the ForcePerLength value in DecinewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) DecinewtonsPerMillimeter() float64 {
	if a.decinewtons_per_millimeterLazy != nil {
		return *a.decinewtons_per_millimeterLazy
	}
	decinewtons_per_millimeter := a.convertFromBase(ForcePerLengthDecinewtonPerMillimeter)
	a.decinewtons_per_millimeterLazy = &decinewtons_per_millimeter
	return decinewtons_per_millimeter
}

// DecanewtonsPerMillimeter returns the ForcePerLength value in DecanewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) DecanewtonsPerMillimeter() float64 {
	if a.decanewtons_per_millimeterLazy != nil {
		return *a.decanewtons_per_millimeterLazy
	}
	decanewtons_per_millimeter := a.convertFromBase(ForcePerLengthDecanewtonPerMillimeter)
	a.decanewtons_per_millimeterLazy = &decanewtons_per_millimeter
	return decanewtons_per_millimeter
}

// KilonewtonsPerMillimeter returns the ForcePerLength value in KilonewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) KilonewtonsPerMillimeter() float64 {
	if a.kilonewtons_per_millimeterLazy != nil {
		return *a.kilonewtons_per_millimeterLazy
	}
	kilonewtons_per_millimeter := a.convertFromBase(ForcePerLengthKilonewtonPerMillimeter)
	a.kilonewtons_per_millimeterLazy = &kilonewtons_per_millimeter
	return kilonewtons_per_millimeter
}

// MeganewtonsPerMillimeter returns the ForcePerLength value in MeganewtonsPerMillimeter.
//
// 
func (a *ForcePerLength) MeganewtonsPerMillimeter() float64 {
	if a.meganewtons_per_millimeterLazy != nil {
		return *a.meganewtons_per_millimeterLazy
	}
	meganewtons_per_millimeter := a.convertFromBase(ForcePerLengthMeganewtonPerMillimeter)
	a.meganewtons_per_millimeterLazy = &meganewtons_per_millimeter
	return meganewtons_per_millimeter
}


// ToDto creates a ForcePerLengthDto representation from the ForcePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerMeter by default.
func (a *ForcePerLength) ToDto(holdInUnit *ForcePerLengthUnits) ForcePerLengthDto {
	if holdInUnit == nil {
		defaultUnit := ForcePerLengthNewtonPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ForcePerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ForcePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerMeter by default.
func (a *ForcePerLength) ToDtoJSON(holdInUnit *ForcePerLengthUnits) ([]byte, error) {
	// Convert to ForcePerLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ForcePerLength to a specific unit value.
// The function uses the provided unit type (ForcePerLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ForcePerLength) Convert(toUnit ForcePerLengthUnits) float64 {
	switch toUnit { 
    case ForcePerLengthNewtonPerMeter:
		return a.NewtonsPerMeter()
    case ForcePerLengthNewtonPerCentimeter:
		return a.NewtonsPerCentimeter()
    case ForcePerLengthNewtonPerMillimeter:
		return a.NewtonsPerMillimeter()
    case ForcePerLengthKilogramForcePerMeter:
		return a.KilogramsForcePerMeter()
    case ForcePerLengthKilogramForcePerCentimeter:
		return a.KilogramsForcePerCentimeter()
    case ForcePerLengthKilogramForcePerMillimeter:
		return a.KilogramsForcePerMillimeter()
    case ForcePerLengthTonneForcePerMeter:
		return a.TonnesForcePerMeter()
    case ForcePerLengthTonneForcePerCentimeter:
		return a.TonnesForcePerCentimeter()
    case ForcePerLengthTonneForcePerMillimeter:
		return a.TonnesForcePerMillimeter()
    case ForcePerLengthPoundForcePerFoot:
		return a.PoundsForcePerFoot()
    case ForcePerLengthPoundForcePerInch:
		return a.PoundsForcePerInch()
    case ForcePerLengthPoundForcePerYard:
		return a.PoundsForcePerYard()
    case ForcePerLengthKilopoundForcePerFoot:
		return a.KilopoundsForcePerFoot()
    case ForcePerLengthKilopoundForcePerInch:
		return a.KilopoundsForcePerInch()
    case ForcePerLengthNanonewtonPerMeter:
		return a.NanonewtonsPerMeter()
    case ForcePerLengthMicronewtonPerMeter:
		return a.MicronewtonsPerMeter()
    case ForcePerLengthMillinewtonPerMeter:
		return a.MillinewtonsPerMeter()
    case ForcePerLengthCentinewtonPerMeter:
		return a.CentinewtonsPerMeter()
    case ForcePerLengthDecinewtonPerMeter:
		return a.DecinewtonsPerMeter()
    case ForcePerLengthDecanewtonPerMeter:
		return a.DecanewtonsPerMeter()
    case ForcePerLengthKilonewtonPerMeter:
		return a.KilonewtonsPerMeter()
    case ForcePerLengthMeganewtonPerMeter:
		return a.MeganewtonsPerMeter()
    case ForcePerLengthNanonewtonPerCentimeter:
		return a.NanonewtonsPerCentimeter()
    case ForcePerLengthMicronewtonPerCentimeter:
		return a.MicronewtonsPerCentimeter()
    case ForcePerLengthMillinewtonPerCentimeter:
		return a.MillinewtonsPerCentimeter()
    case ForcePerLengthCentinewtonPerCentimeter:
		return a.CentinewtonsPerCentimeter()
    case ForcePerLengthDecinewtonPerCentimeter:
		return a.DecinewtonsPerCentimeter()
    case ForcePerLengthDecanewtonPerCentimeter:
		return a.DecanewtonsPerCentimeter()
    case ForcePerLengthKilonewtonPerCentimeter:
		return a.KilonewtonsPerCentimeter()
    case ForcePerLengthMeganewtonPerCentimeter:
		return a.MeganewtonsPerCentimeter()
    case ForcePerLengthNanonewtonPerMillimeter:
		return a.NanonewtonsPerMillimeter()
    case ForcePerLengthMicronewtonPerMillimeter:
		return a.MicronewtonsPerMillimeter()
    case ForcePerLengthMillinewtonPerMillimeter:
		return a.MillinewtonsPerMillimeter()
    case ForcePerLengthCentinewtonPerMillimeter:
		return a.CentinewtonsPerMillimeter()
    case ForcePerLengthDecinewtonPerMillimeter:
		return a.DecinewtonsPerMillimeter()
    case ForcePerLengthDecanewtonPerMillimeter:
		return a.DecanewtonsPerMillimeter()
    case ForcePerLengthKilonewtonPerMillimeter:
		return a.KilonewtonsPerMillimeter()
    case ForcePerLengthMeganewtonPerMillimeter:
		return a.MeganewtonsPerMillimeter()
	default:
		return math.NaN()
	}
}

func (a *ForcePerLength) convertFromBase(toUnit ForcePerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case ForcePerLengthNewtonPerMeter:
		return (value) 
	case ForcePerLengthNewtonPerCentimeter:
		return (value / 1e2) 
	case ForcePerLengthNewtonPerMillimeter:
		return (value / 1e3) 
	case ForcePerLengthKilogramForcePerMeter:
		return (value / 9.80665002864) 
	case ForcePerLengthKilogramForcePerCentimeter:
		return (value / 980.665002864) 
	case ForcePerLengthKilogramForcePerMillimeter:
		return (value / 9.80665002864e3) 
	case ForcePerLengthTonneForcePerMeter:
		return (value / 9.80665002864e3) 
	case ForcePerLengthTonneForcePerCentimeter:
		return (value / 9.80665002864e5) 
	case ForcePerLengthTonneForcePerMillimeter:
		return (value / 9.80665002864e6) 
	case ForcePerLengthPoundForcePerFoot:
		return (value / 14.59390292) 
	case ForcePerLengthPoundForcePerInch:
		return (value / 1.75126835e2) 
	case ForcePerLengthPoundForcePerYard:
		return (value / 4.864634307) 
	case ForcePerLengthKilopoundForcePerFoot:
		return (value / 14593.90292) 
	case ForcePerLengthKilopoundForcePerInch:
		return (value / 1.75126835e5) 
	case ForcePerLengthNanonewtonPerMeter:
		return ((value) / 1e-09) 
	case ForcePerLengthMicronewtonPerMeter:
		return ((value) / 1e-06) 
	case ForcePerLengthMillinewtonPerMeter:
		return ((value) / 0.001) 
	case ForcePerLengthCentinewtonPerMeter:
		return ((value) / 0.01) 
	case ForcePerLengthDecinewtonPerMeter:
		return ((value) / 0.1) 
	case ForcePerLengthDecanewtonPerMeter:
		return ((value) / 10.0) 
	case ForcePerLengthKilonewtonPerMeter:
		return ((value) / 1000.0) 
	case ForcePerLengthMeganewtonPerMeter:
		return ((value) / 1000000.0) 
	case ForcePerLengthNanonewtonPerCentimeter:
		return ((value / 1e2) / 1e-09) 
	case ForcePerLengthMicronewtonPerCentimeter:
		return ((value / 1e2) / 1e-06) 
	case ForcePerLengthMillinewtonPerCentimeter:
		return ((value / 1e2) / 0.001) 
	case ForcePerLengthCentinewtonPerCentimeter:
		return ((value / 1e2) / 0.01) 
	case ForcePerLengthDecinewtonPerCentimeter:
		return ((value / 1e2) / 0.1) 
	case ForcePerLengthDecanewtonPerCentimeter:
		return ((value / 1e2) / 10.0) 
	case ForcePerLengthKilonewtonPerCentimeter:
		return ((value / 1e2) / 1000.0) 
	case ForcePerLengthMeganewtonPerCentimeter:
		return ((value / 1e2) / 1000000.0) 
	case ForcePerLengthNanonewtonPerMillimeter:
		return ((value / 1e3) / 1e-09) 
	case ForcePerLengthMicronewtonPerMillimeter:
		return ((value / 1e3) / 1e-06) 
	case ForcePerLengthMillinewtonPerMillimeter:
		return ((value / 1e3) / 0.001) 
	case ForcePerLengthCentinewtonPerMillimeter:
		return ((value / 1e3) / 0.01) 
	case ForcePerLengthDecinewtonPerMillimeter:
		return ((value / 1e3) / 0.1) 
	case ForcePerLengthDecanewtonPerMillimeter:
		return ((value / 1e3) / 10.0) 
	case ForcePerLengthKilonewtonPerMillimeter:
		return ((value / 1e3) / 1000.0) 
	case ForcePerLengthMeganewtonPerMillimeter:
		return ((value / 1e3) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ForcePerLength) convertToBase(value float64, fromUnit ForcePerLengthUnits) float64 {
	switch fromUnit { 
	case ForcePerLengthNewtonPerMeter:
		return (value) 
	case ForcePerLengthNewtonPerCentimeter:
		return (value * 1e2) 
	case ForcePerLengthNewtonPerMillimeter:
		return (value * 1e3) 
	case ForcePerLengthKilogramForcePerMeter:
		return (value * 9.80665002864) 
	case ForcePerLengthKilogramForcePerCentimeter:
		return (value * 980.665002864) 
	case ForcePerLengthKilogramForcePerMillimeter:
		return (value * 9.80665002864e3) 
	case ForcePerLengthTonneForcePerMeter:
		return (value * 9.80665002864e3) 
	case ForcePerLengthTonneForcePerCentimeter:
		return (value * 9.80665002864e5) 
	case ForcePerLengthTonneForcePerMillimeter:
		return (value * 9.80665002864e6) 
	case ForcePerLengthPoundForcePerFoot:
		return (value * 14.59390292) 
	case ForcePerLengthPoundForcePerInch:
		return (value * 1.75126835e2) 
	case ForcePerLengthPoundForcePerYard:
		return (value * 4.864634307) 
	case ForcePerLengthKilopoundForcePerFoot:
		return (value * 14593.90292) 
	case ForcePerLengthKilopoundForcePerInch:
		return (value * 1.75126835e5) 
	case ForcePerLengthNanonewtonPerMeter:
		return ((value) * 1e-09) 
	case ForcePerLengthMicronewtonPerMeter:
		return ((value) * 1e-06) 
	case ForcePerLengthMillinewtonPerMeter:
		return ((value) * 0.001) 
	case ForcePerLengthCentinewtonPerMeter:
		return ((value) * 0.01) 
	case ForcePerLengthDecinewtonPerMeter:
		return ((value) * 0.1) 
	case ForcePerLengthDecanewtonPerMeter:
		return ((value) * 10.0) 
	case ForcePerLengthKilonewtonPerMeter:
		return ((value) * 1000.0) 
	case ForcePerLengthMeganewtonPerMeter:
		return ((value) * 1000000.0) 
	case ForcePerLengthNanonewtonPerCentimeter:
		return ((value * 1e2) * 1e-09) 
	case ForcePerLengthMicronewtonPerCentimeter:
		return ((value * 1e2) * 1e-06) 
	case ForcePerLengthMillinewtonPerCentimeter:
		return ((value * 1e2) * 0.001) 
	case ForcePerLengthCentinewtonPerCentimeter:
		return ((value * 1e2) * 0.01) 
	case ForcePerLengthDecinewtonPerCentimeter:
		return ((value * 1e2) * 0.1) 
	case ForcePerLengthDecanewtonPerCentimeter:
		return ((value * 1e2) * 10.0) 
	case ForcePerLengthKilonewtonPerCentimeter:
		return ((value * 1e2) * 1000.0) 
	case ForcePerLengthMeganewtonPerCentimeter:
		return ((value * 1e2) * 1000000.0) 
	case ForcePerLengthNanonewtonPerMillimeter:
		return ((value * 1e3) * 1e-09) 
	case ForcePerLengthMicronewtonPerMillimeter:
		return ((value * 1e3) * 1e-06) 
	case ForcePerLengthMillinewtonPerMillimeter:
		return ((value * 1e3) * 0.001) 
	case ForcePerLengthCentinewtonPerMillimeter:
		return ((value * 1e3) * 0.01) 
	case ForcePerLengthDecinewtonPerMillimeter:
		return ((value * 1e3) * 0.1) 
	case ForcePerLengthDecanewtonPerMillimeter:
		return ((value * 1e3) * 10.0) 
	case ForcePerLengthKilonewtonPerMillimeter:
		return ((value * 1e3) * 1000.0) 
	case ForcePerLengthMeganewtonPerMillimeter:
		return ((value * 1e3) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ForcePerLength in the default unit (NewtonPerMeter),
// formatted to two decimal places.
func (a ForcePerLength) String() string {
	return a.ToString(ForcePerLengthNewtonPerMeter, 2)
}

// ToString formats the ForcePerLength value as a string with the specified unit and fractional digits.
// It converts the ForcePerLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ForcePerLength value will be converted (e.g., NewtonPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ForcePerLength with the unit abbreviation.
func (a *ForcePerLength) ToString(unit ForcePerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetForcePerLengthAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetForcePerLengthAbbreviation(unit))
}

// Equals checks if the given ForcePerLength is equal to the current ForcePerLength.
//
// Parameters:
//    other: The ForcePerLength to compare against.
//
// Returns:
//    bool: Returns true if both ForcePerLength are equal, false otherwise.
func (a *ForcePerLength) Equals(other *ForcePerLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ForcePerLength with another ForcePerLength.
// It returns -1 if the current ForcePerLength is less than the other ForcePerLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ForcePerLength to compare against.
//
// Returns:
//    int: -1 if the current ForcePerLength is less, 1 if greater, and 0 if equal.
func (a *ForcePerLength) CompareTo(other *ForcePerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ForcePerLength to the current ForcePerLength and returns the result.
// The result is a new ForcePerLength instance with the sum of the values.
//
// Parameters:
//    other: The ForcePerLength to add to the current ForcePerLength.
//
// Returns:
//    *ForcePerLength: A new ForcePerLength instance representing the sum of both ForcePerLength.
func (a *ForcePerLength) Add(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ForcePerLength from the current ForcePerLength and returns the result.
// The result is a new ForcePerLength instance with the difference of the values.
//
// Parameters:
//    other: The ForcePerLength to subtract from the current ForcePerLength.
//
// Returns:
//    *ForcePerLength: A new ForcePerLength instance representing the difference of both ForcePerLength.
func (a *ForcePerLength) Subtract(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ForcePerLength by the given ForcePerLength and returns the result.
// The result is a new ForcePerLength instance with the product of the values.
//
// Parameters:
//    other: The ForcePerLength to multiply with the current ForcePerLength.
//
// Returns:
//    *ForcePerLength: A new ForcePerLength instance representing the product of both ForcePerLength.
func (a *ForcePerLength) Multiply(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value * other.BaseValue()}
}

// Divide divides the current ForcePerLength by the given ForcePerLength and returns the result.
// The result is a new ForcePerLength instance with the quotient of the values.
//
// Parameters:
//    other: The ForcePerLength to divide the current ForcePerLength by.
//
// Returns:
//    *ForcePerLength: A new ForcePerLength instance representing the quotient of both ForcePerLength.
func (a *ForcePerLength) Divide(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value / other.BaseValue()}
}

// GetForcePerLengthAbbreviation gets the unit abbreviation.
func GetForcePerLengthAbbreviation(unit ForcePerLengthUnits) string {
	switch unit { 
	case ForcePerLengthNewtonPerMeter:
		return "N/m" 
	case ForcePerLengthNewtonPerCentimeter:
		return "N/cm" 
	case ForcePerLengthNewtonPerMillimeter:
		return "N/mm" 
	case ForcePerLengthKilogramForcePerMeter:
		return "kgf/m" 
	case ForcePerLengthKilogramForcePerCentimeter:
		return "kgf/cm" 
	case ForcePerLengthKilogramForcePerMillimeter:
		return "kgf/mm" 
	case ForcePerLengthTonneForcePerMeter:
		return "tf/m" 
	case ForcePerLengthTonneForcePerCentimeter:
		return "tf/cm" 
	case ForcePerLengthTonneForcePerMillimeter:
		return "tf/mm" 
	case ForcePerLengthPoundForcePerFoot:
		return "lbf/ft" 
	case ForcePerLengthPoundForcePerInch:
		return "lbf/in" 
	case ForcePerLengthPoundForcePerYard:
		return "lbf/yd" 
	case ForcePerLengthKilopoundForcePerFoot:
		return "kipf/ft" 
	case ForcePerLengthKilopoundForcePerInch:
		return "kipf/in" 
	case ForcePerLengthNanonewtonPerMeter:
		return "nN/m" 
	case ForcePerLengthMicronewtonPerMeter:
		return "μN/m" 
	case ForcePerLengthMillinewtonPerMeter:
		return "mN/m" 
	case ForcePerLengthCentinewtonPerMeter:
		return "cN/m" 
	case ForcePerLengthDecinewtonPerMeter:
		return "dN/m" 
	case ForcePerLengthDecanewtonPerMeter:
		return "daN/m" 
	case ForcePerLengthKilonewtonPerMeter:
		return "kN/m" 
	case ForcePerLengthMeganewtonPerMeter:
		return "MN/m" 
	case ForcePerLengthNanonewtonPerCentimeter:
		return "nN/cm" 
	case ForcePerLengthMicronewtonPerCentimeter:
		return "μN/cm" 
	case ForcePerLengthMillinewtonPerCentimeter:
		return "mN/cm" 
	case ForcePerLengthCentinewtonPerCentimeter:
		return "cN/cm" 
	case ForcePerLengthDecinewtonPerCentimeter:
		return "dN/cm" 
	case ForcePerLengthDecanewtonPerCentimeter:
		return "daN/cm" 
	case ForcePerLengthKilonewtonPerCentimeter:
		return "kN/cm" 
	case ForcePerLengthMeganewtonPerCentimeter:
		return "MN/cm" 
	case ForcePerLengthNanonewtonPerMillimeter:
		return "nN/mm" 
	case ForcePerLengthMicronewtonPerMillimeter:
		return "μN/mm" 
	case ForcePerLengthMillinewtonPerMillimeter:
		return "mN/mm" 
	case ForcePerLengthCentinewtonPerMillimeter:
		return "cN/mm" 
	case ForcePerLengthDecinewtonPerMillimeter:
		return "dN/mm" 
	case ForcePerLengthDecanewtonPerMillimeter:
		return "daN/mm" 
	case ForcePerLengthKilonewtonPerMillimeter:
		return "kN/mm" 
	case ForcePerLengthMeganewtonPerMillimeter:
		return "MN/mm" 
	default:
		return ""
	}
}