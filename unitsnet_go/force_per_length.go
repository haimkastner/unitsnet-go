package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ForcePerLengthUnits enumeration
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

// ForcePerLengthDto represents an ForcePerLength
type ForcePerLengthDto struct {
	Value float64
	Unit  ForcePerLengthUnits
}

// ForcePerLengthDtoFactory struct to group related functions
type ForcePerLengthDtoFactory struct{}

func (udf ForcePerLengthDtoFactory) FromJSON(data []byte) (*ForcePerLengthDto, error) {
	a := ForcePerLengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ForcePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ForcePerLength struct
type ForcePerLength struct {
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

// ForcePerLengthFactory struct to group related functions
type ForcePerLengthFactory struct{}

func (uf ForcePerLengthFactory) CreateForcePerLength(value float64, unit ForcePerLengthUnits) (*ForcePerLength, error) {
	return newForcePerLength(value, unit)
}

func (uf ForcePerLengthFactory) FromDto(dto ForcePerLengthDto) (*ForcePerLength, error) {
	return newForcePerLength(dto.Value, dto.Unit)
}

func (uf ForcePerLengthFactory) FromDtoJSON(data []byte) (*ForcePerLength, error) {
	unitDto, err := ForcePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ForcePerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonPerMeter creates a new ForcePerLength instance from NewtonPerMeter.
func (uf ForcePerLengthFactory) FromNewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerMeter)
}

// FromNewtonPerCentimeter creates a new ForcePerLength instance from NewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromNewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerCentimeter)
}

// FromNewtonPerMillimeter creates a new ForcePerLength instance from NewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromNewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNewtonPerMillimeter)
}

// FromKilogramForcePerMeter creates a new ForcePerLength instance from KilogramForcePerMeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerMeter)
}

// FromKilogramForcePerCentimeter creates a new ForcePerLength instance from KilogramForcePerCentimeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerCentimeter)
}

// FromKilogramForcePerMillimeter creates a new ForcePerLength instance from KilogramForcePerMillimeter.
func (uf ForcePerLengthFactory) FromKilogramsForcePerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilogramForcePerMillimeter)
}

// FromTonneForcePerMeter creates a new ForcePerLength instance from TonneForcePerMeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerMeter)
}

// FromTonneForcePerCentimeter creates a new ForcePerLength instance from TonneForcePerCentimeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerCentimeter)
}

// FromTonneForcePerMillimeter creates a new ForcePerLength instance from TonneForcePerMillimeter.
func (uf ForcePerLengthFactory) FromTonnesForcePerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthTonneForcePerMillimeter)
}

// FromPoundForcePerFoot creates a new ForcePerLength instance from PoundForcePerFoot.
func (uf ForcePerLengthFactory) FromPoundsForcePerFoot(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerFoot)
}

// FromPoundForcePerInch creates a new ForcePerLength instance from PoundForcePerInch.
func (uf ForcePerLengthFactory) FromPoundsForcePerInch(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerInch)
}

// FromPoundForcePerYard creates a new ForcePerLength instance from PoundForcePerYard.
func (uf ForcePerLengthFactory) FromPoundsForcePerYard(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthPoundForcePerYard)
}

// FromKilopoundForcePerFoot creates a new ForcePerLength instance from KilopoundForcePerFoot.
func (uf ForcePerLengthFactory) FromKilopoundsForcePerFoot(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilopoundForcePerFoot)
}

// FromKilopoundForcePerInch creates a new ForcePerLength instance from KilopoundForcePerInch.
func (uf ForcePerLengthFactory) FromKilopoundsForcePerInch(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilopoundForcePerInch)
}

// FromNanonewtonPerMeter creates a new ForcePerLength instance from NanonewtonPerMeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerMeter)
}

// FromMicronewtonPerMeter creates a new ForcePerLength instance from MicronewtonPerMeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerMeter)
}

// FromMillinewtonPerMeter creates a new ForcePerLength instance from MillinewtonPerMeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerMeter)
}

// FromCentinewtonPerMeter creates a new ForcePerLength instance from CentinewtonPerMeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerMeter)
}

// FromDecinewtonPerMeter creates a new ForcePerLength instance from DecinewtonPerMeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerMeter)
}

// FromDecanewtonPerMeter creates a new ForcePerLength instance from DecanewtonPerMeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerMeter)
}

// FromKilonewtonPerMeter creates a new ForcePerLength instance from KilonewtonPerMeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerMeter)
}

// FromMeganewtonPerMeter creates a new ForcePerLength instance from MeganewtonPerMeter.
func (uf ForcePerLengthFactory) FromMeganewtonsPerMeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMeganewtonPerMeter)
}

// FromNanonewtonPerCentimeter creates a new ForcePerLength instance from NanonewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerCentimeter)
}

// FromMicronewtonPerCentimeter creates a new ForcePerLength instance from MicronewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerCentimeter)
}

// FromMillinewtonPerCentimeter creates a new ForcePerLength instance from MillinewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerCentimeter)
}

// FromCentinewtonPerCentimeter creates a new ForcePerLength instance from CentinewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerCentimeter)
}

// FromDecinewtonPerCentimeter creates a new ForcePerLength instance from DecinewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerCentimeter)
}

// FromDecanewtonPerCentimeter creates a new ForcePerLength instance from DecanewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerCentimeter)
}

// FromKilonewtonPerCentimeter creates a new ForcePerLength instance from KilonewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerCentimeter)
}

// FromMeganewtonPerCentimeter creates a new ForcePerLength instance from MeganewtonPerCentimeter.
func (uf ForcePerLengthFactory) FromMeganewtonsPerCentimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMeganewtonPerCentimeter)
}

// FromNanonewtonPerMillimeter creates a new ForcePerLength instance from NanonewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromNanonewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthNanonewtonPerMillimeter)
}

// FromMicronewtonPerMillimeter creates a new ForcePerLength instance from MicronewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromMicronewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMicronewtonPerMillimeter)
}

// FromMillinewtonPerMillimeter creates a new ForcePerLength instance from MillinewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromMillinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthMillinewtonPerMillimeter)
}

// FromCentinewtonPerMillimeter creates a new ForcePerLength instance from CentinewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromCentinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthCentinewtonPerMillimeter)
}

// FromDecinewtonPerMillimeter creates a new ForcePerLength instance from DecinewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromDecinewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecinewtonPerMillimeter)
}

// FromDecanewtonPerMillimeter creates a new ForcePerLength instance from DecanewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromDecanewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthDecanewtonPerMillimeter)
}

// FromKilonewtonPerMillimeter creates a new ForcePerLength instance from KilonewtonPerMillimeter.
func (uf ForcePerLengthFactory) FromKilonewtonsPerMillimeter(value float64) (*ForcePerLength, error) {
	return newForcePerLength(value, ForcePerLengthKilonewtonPerMillimeter)
}

// FromMeganewtonPerMillimeter creates a new ForcePerLength instance from MeganewtonPerMillimeter.
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

// BaseValue returns the base value of ForcePerLength in NewtonPerMeter.
func (a *ForcePerLength) BaseValue() float64 {
	return a.value
}


// NewtonPerMeter returns the value in NewtonPerMeter.
func (a *ForcePerLength) NewtonsPerMeter() float64 {
	if a.newtons_per_meterLazy != nil {
		return *a.newtons_per_meterLazy
	}
	newtons_per_meter := a.convertFromBase(ForcePerLengthNewtonPerMeter)
	a.newtons_per_meterLazy = &newtons_per_meter
	return newtons_per_meter
}

// NewtonPerCentimeter returns the value in NewtonPerCentimeter.
func (a *ForcePerLength) NewtonsPerCentimeter() float64 {
	if a.newtons_per_centimeterLazy != nil {
		return *a.newtons_per_centimeterLazy
	}
	newtons_per_centimeter := a.convertFromBase(ForcePerLengthNewtonPerCentimeter)
	a.newtons_per_centimeterLazy = &newtons_per_centimeter
	return newtons_per_centimeter
}

// NewtonPerMillimeter returns the value in NewtonPerMillimeter.
func (a *ForcePerLength) NewtonsPerMillimeter() float64 {
	if a.newtons_per_millimeterLazy != nil {
		return *a.newtons_per_millimeterLazy
	}
	newtons_per_millimeter := a.convertFromBase(ForcePerLengthNewtonPerMillimeter)
	a.newtons_per_millimeterLazy = &newtons_per_millimeter
	return newtons_per_millimeter
}

// KilogramForcePerMeter returns the value in KilogramForcePerMeter.
func (a *ForcePerLength) KilogramsForcePerMeter() float64 {
	if a.kilograms_force_per_meterLazy != nil {
		return *a.kilograms_force_per_meterLazy
	}
	kilograms_force_per_meter := a.convertFromBase(ForcePerLengthKilogramForcePerMeter)
	a.kilograms_force_per_meterLazy = &kilograms_force_per_meter
	return kilograms_force_per_meter
}

// KilogramForcePerCentimeter returns the value in KilogramForcePerCentimeter.
func (a *ForcePerLength) KilogramsForcePerCentimeter() float64 {
	if a.kilograms_force_per_centimeterLazy != nil {
		return *a.kilograms_force_per_centimeterLazy
	}
	kilograms_force_per_centimeter := a.convertFromBase(ForcePerLengthKilogramForcePerCentimeter)
	a.kilograms_force_per_centimeterLazy = &kilograms_force_per_centimeter
	return kilograms_force_per_centimeter
}

// KilogramForcePerMillimeter returns the value in KilogramForcePerMillimeter.
func (a *ForcePerLength) KilogramsForcePerMillimeter() float64 {
	if a.kilograms_force_per_millimeterLazy != nil {
		return *a.kilograms_force_per_millimeterLazy
	}
	kilograms_force_per_millimeter := a.convertFromBase(ForcePerLengthKilogramForcePerMillimeter)
	a.kilograms_force_per_millimeterLazy = &kilograms_force_per_millimeter
	return kilograms_force_per_millimeter
}

// TonneForcePerMeter returns the value in TonneForcePerMeter.
func (a *ForcePerLength) TonnesForcePerMeter() float64 {
	if a.tonnes_force_per_meterLazy != nil {
		return *a.tonnes_force_per_meterLazy
	}
	tonnes_force_per_meter := a.convertFromBase(ForcePerLengthTonneForcePerMeter)
	a.tonnes_force_per_meterLazy = &tonnes_force_per_meter
	return tonnes_force_per_meter
}

// TonneForcePerCentimeter returns the value in TonneForcePerCentimeter.
func (a *ForcePerLength) TonnesForcePerCentimeter() float64 {
	if a.tonnes_force_per_centimeterLazy != nil {
		return *a.tonnes_force_per_centimeterLazy
	}
	tonnes_force_per_centimeter := a.convertFromBase(ForcePerLengthTonneForcePerCentimeter)
	a.tonnes_force_per_centimeterLazy = &tonnes_force_per_centimeter
	return tonnes_force_per_centimeter
}

// TonneForcePerMillimeter returns the value in TonneForcePerMillimeter.
func (a *ForcePerLength) TonnesForcePerMillimeter() float64 {
	if a.tonnes_force_per_millimeterLazy != nil {
		return *a.tonnes_force_per_millimeterLazy
	}
	tonnes_force_per_millimeter := a.convertFromBase(ForcePerLengthTonneForcePerMillimeter)
	a.tonnes_force_per_millimeterLazy = &tonnes_force_per_millimeter
	return tonnes_force_per_millimeter
}

// PoundForcePerFoot returns the value in PoundForcePerFoot.
func (a *ForcePerLength) PoundsForcePerFoot() float64 {
	if a.pounds_force_per_footLazy != nil {
		return *a.pounds_force_per_footLazy
	}
	pounds_force_per_foot := a.convertFromBase(ForcePerLengthPoundForcePerFoot)
	a.pounds_force_per_footLazy = &pounds_force_per_foot
	return pounds_force_per_foot
}

// PoundForcePerInch returns the value in PoundForcePerInch.
func (a *ForcePerLength) PoundsForcePerInch() float64 {
	if a.pounds_force_per_inchLazy != nil {
		return *a.pounds_force_per_inchLazy
	}
	pounds_force_per_inch := a.convertFromBase(ForcePerLengthPoundForcePerInch)
	a.pounds_force_per_inchLazy = &pounds_force_per_inch
	return pounds_force_per_inch
}

// PoundForcePerYard returns the value in PoundForcePerYard.
func (a *ForcePerLength) PoundsForcePerYard() float64 {
	if a.pounds_force_per_yardLazy != nil {
		return *a.pounds_force_per_yardLazy
	}
	pounds_force_per_yard := a.convertFromBase(ForcePerLengthPoundForcePerYard)
	a.pounds_force_per_yardLazy = &pounds_force_per_yard
	return pounds_force_per_yard
}

// KilopoundForcePerFoot returns the value in KilopoundForcePerFoot.
func (a *ForcePerLength) KilopoundsForcePerFoot() float64 {
	if a.kilopounds_force_per_footLazy != nil {
		return *a.kilopounds_force_per_footLazy
	}
	kilopounds_force_per_foot := a.convertFromBase(ForcePerLengthKilopoundForcePerFoot)
	a.kilopounds_force_per_footLazy = &kilopounds_force_per_foot
	return kilopounds_force_per_foot
}

// KilopoundForcePerInch returns the value in KilopoundForcePerInch.
func (a *ForcePerLength) KilopoundsForcePerInch() float64 {
	if a.kilopounds_force_per_inchLazy != nil {
		return *a.kilopounds_force_per_inchLazy
	}
	kilopounds_force_per_inch := a.convertFromBase(ForcePerLengthKilopoundForcePerInch)
	a.kilopounds_force_per_inchLazy = &kilopounds_force_per_inch
	return kilopounds_force_per_inch
}

// NanonewtonPerMeter returns the value in NanonewtonPerMeter.
func (a *ForcePerLength) NanonewtonsPerMeter() float64 {
	if a.nanonewtons_per_meterLazy != nil {
		return *a.nanonewtons_per_meterLazy
	}
	nanonewtons_per_meter := a.convertFromBase(ForcePerLengthNanonewtonPerMeter)
	a.nanonewtons_per_meterLazy = &nanonewtons_per_meter
	return nanonewtons_per_meter
}

// MicronewtonPerMeter returns the value in MicronewtonPerMeter.
func (a *ForcePerLength) MicronewtonsPerMeter() float64 {
	if a.micronewtons_per_meterLazy != nil {
		return *a.micronewtons_per_meterLazy
	}
	micronewtons_per_meter := a.convertFromBase(ForcePerLengthMicronewtonPerMeter)
	a.micronewtons_per_meterLazy = &micronewtons_per_meter
	return micronewtons_per_meter
}

// MillinewtonPerMeter returns the value in MillinewtonPerMeter.
func (a *ForcePerLength) MillinewtonsPerMeter() float64 {
	if a.millinewtons_per_meterLazy != nil {
		return *a.millinewtons_per_meterLazy
	}
	millinewtons_per_meter := a.convertFromBase(ForcePerLengthMillinewtonPerMeter)
	a.millinewtons_per_meterLazy = &millinewtons_per_meter
	return millinewtons_per_meter
}

// CentinewtonPerMeter returns the value in CentinewtonPerMeter.
func (a *ForcePerLength) CentinewtonsPerMeter() float64 {
	if a.centinewtons_per_meterLazy != nil {
		return *a.centinewtons_per_meterLazy
	}
	centinewtons_per_meter := a.convertFromBase(ForcePerLengthCentinewtonPerMeter)
	a.centinewtons_per_meterLazy = &centinewtons_per_meter
	return centinewtons_per_meter
}

// DecinewtonPerMeter returns the value in DecinewtonPerMeter.
func (a *ForcePerLength) DecinewtonsPerMeter() float64 {
	if a.decinewtons_per_meterLazy != nil {
		return *a.decinewtons_per_meterLazy
	}
	decinewtons_per_meter := a.convertFromBase(ForcePerLengthDecinewtonPerMeter)
	a.decinewtons_per_meterLazy = &decinewtons_per_meter
	return decinewtons_per_meter
}

// DecanewtonPerMeter returns the value in DecanewtonPerMeter.
func (a *ForcePerLength) DecanewtonsPerMeter() float64 {
	if a.decanewtons_per_meterLazy != nil {
		return *a.decanewtons_per_meterLazy
	}
	decanewtons_per_meter := a.convertFromBase(ForcePerLengthDecanewtonPerMeter)
	a.decanewtons_per_meterLazy = &decanewtons_per_meter
	return decanewtons_per_meter
}

// KilonewtonPerMeter returns the value in KilonewtonPerMeter.
func (a *ForcePerLength) KilonewtonsPerMeter() float64 {
	if a.kilonewtons_per_meterLazy != nil {
		return *a.kilonewtons_per_meterLazy
	}
	kilonewtons_per_meter := a.convertFromBase(ForcePerLengthKilonewtonPerMeter)
	a.kilonewtons_per_meterLazy = &kilonewtons_per_meter
	return kilonewtons_per_meter
}

// MeganewtonPerMeter returns the value in MeganewtonPerMeter.
func (a *ForcePerLength) MeganewtonsPerMeter() float64 {
	if a.meganewtons_per_meterLazy != nil {
		return *a.meganewtons_per_meterLazy
	}
	meganewtons_per_meter := a.convertFromBase(ForcePerLengthMeganewtonPerMeter)
	a.meganewtons_per_meterLazy = &meganewtons_per_meter
	return meganewtons_per_meter
}

// NanonewtonPerCentimeter returns the value in NanonewtonPerCentimeter.
func (a *ForcePerLength) NanonewtonsPerCentimeter() float64 {
	if a.nanonewtons_per_centimeterLazy != nil {
		return *a.nanonewtons_per_centimeterLazy
	}
	nanonewtons_per_centimeter := a.convertFromBase(ForcePerLengthNanonewtonPerCentimeter)
	a.nanonewtons_per_centimeterLazy = &nanonewtons_per_centimeter
	return nanonewtons_per_centimeter
}

// MicronewtonPerCentimeter returns the value in MicronewtonPerCentimeter.
func (a *ForcePerLength) MicronewtonsPerCentimeter() float64 {
	if a.micronewtons_per_centimeterLazy != nil {
		return *a.micronewtons_per_centimeterLazy
	}
	micronewtons_per_centimeter := a.convertFromBase(ForcePerLengthMicronewtonPerCentimeter)
	a.micronewtons_per_centimeterLazy = &micronewtons_per_centimeter
	return micronewtons_per_centimeter
}

// MillinewtonPerCentimeter returns the value in MillinewtonPerCentimeter.
func (a *ForcePerLength) MillinewtonsPerCentimeter() float64 {
	if a.millinewtons_per_centimeterLazy != nil {
		return *a.millinewtons_per_centimeterLazy
	}
	millinewtons_per_centimeter := a.convertFromBase(ForcePerLengthMillinewtonPerCentimeter)
	a.millinewtons_per_centimeterLazy = &millinewtons_per_centimeter
	return millinewtons_per_centimeter
}

// CentinewtonPerCentimeter returns the value in CentinewtonPerCentimeter.
func (a *ForcePerLength) CentinewtonsPerCentimeter() float64 {
	if a.centinewtons_per_centimeterLazy != nil {
		return *a.centinewtons_per_centimeterLazy
	}
	centinewtons_per_centimeter := a.convertFromBase(ForcePerLengthCentinewtonPerCentimeter)
	a.centinewtons_per_centimeterLazy = &centinewtons_per_centimeter
	return centinewtons_per_centimeter
}

// DecinewtonPerCentimeter returns the value in DecinewtonPerCentimeter.
func (a *ForcePerLength) DecinewtonsPerCentimeter() float64 {
	if a.decinewtons_per_centimeterLazy != nil {
		return *a.decinewtons_per_centimeterLazy
	}
	decinewtons_per_centimeter := a.convertFromBase(ForcePerLengthDecinewtonPerCentimeter)
	a.decinewtons_per_centimeterLazy = &decinewtons_per_centimeter
	return decinewtons_per_centimeter
}

// DecanewtonPerCentimeter returns the value in DecanewtonPerCentimeter.
func (a *ForcePerLength) DecanewtonsPerCentimeter() float64 {
	if a.decanewtons_per_centimeterLazy != nil {
		return *a.decanewtons_per_centimeterLazy
	}
	decanewtons_per_centimeter := a.convertFromBase(ForcePerLengthDecanewtonPerCentimeter)
	a.decanewtons_per_centimeterLazy = &decanewtons_per_centimeter
	return decanewtons_per_centimeter
}

// KilonewtonPerCentimeter returns the value in KilonewtonPerCentimeter.
func (a *ForcePerLength) KilonewtonsPerCentimeter() float64 {
	if a.kilonewtons_per_centimeterLazy != nil {
		return *a.kilonewtons_per_centimeterLazy
	}
	kilonewtons_per_centimeter := a.convertFromBase(ForcePerLengthKilonewtonPerCentimeter)
	a.kilonewtons_per_centimeterLazy = &kilonewtons_per_centimeter
	return kilonewtons_per_centimeter
}

// MeganewtonPerCentimeter returns the value in MeganewtonPerCentimeter.
func (a *ForcePerLength) MeganewtonsPerCentimeter() float64 {
	if a.meganewtons_per_centimeterLazy != nil {
		return *a.meganewtons_per_centimeterLazy
	}
	meganewtons_per_centimeter := a.convertFromBase(ForcePerLengthMeganewtonPerCentimeter)
	a.meganewtons_per_centimeterLazy = &meganewtons_per_centimeter
	return meganewtons_per_centimeter
}

// NanonewtonPerMillimeter returns the value in NanonewtonPerMillimeter.
func (a *ForcePerLength) NanonewtonsPerMillimeter() float64 {
	if a.nanonewtons_per_millimeterLazy != nil {
		return *a.nanonewtons_per_millimeterLazy
	}
	nanonewtons_per_millimeter := a.convertFromBase(ForcePerLengthNanonewtonPerMillimeter)
	a.nanonewtons_per_millimeterLazy = &nanonewtons_per_millimeter
	return nanonewtons_per_millimeter
}

// MicronewtonPerMillimeter returns the value in MicronewtonPerMillimeter.
func (a *ForcePerLength) MicronewtonsPerMillimeter() float64 {
	if a.micronewtons_per_millimeterLazy != nil {
		return *a.micronewtons_per_millimeterLazy
	}
	micronewtons_per_millimeter := a.convertFromBase(ForcePerLengthMicronewtonPerMillimeter)
	a.micronewtons_per_millimeterLazy = &micronewtons_per_millimeter
	return micronewtons_per_millimeter
}

// MillinewtonPerMillimeter returns the value in MillinewtonPerMillimeter.
func (a *ForcePerLength) MillinewtonsPerMillimeter() float64 {
	if a.millinewtons_per_millimeterLazy != nil {
		return *a.millinewtons_per_millimeterLazy
	}
	millinewtons_per_millimeter := a.convertFromBase(ForcePerLengthMillinewtonPerMillimeter)
	a.millinewtons_per_millimeterLazy = &millinewtons_per_millimeter
	return millinewtons_per_millimeter
}

// CentinewtonPerMillimeter returns the value in CentinewtonPerMillimeter.
func (a *ForcePerLength) CentinewtonsPerMillimeter() float64 {
	if a.centinewtons_per_millimeterLazy != nil {
		return *a.centinewtons_per_millimeterLazy
	}
	centinewtons_per_millimeter := a.convertFromBase(ForcePerLengthCentinewtonPerMillimeter)
	a.centinewtons_per_millimeterLazy = &centinewtons_per_millimeter
	return centinewtons_per_millimeter
}

// DecinewtonPerMillimeter returns the value in DecinewtonPerMillimeter.
func (a *ForcePerLength) DecinewtonsPerMillimeter() float64 {
	if a.decinewtons_per_millimeterLazy != nil {
		return *a.decinewtons_per_millimeterLazy
	}
	decinewtons_per_millimeter := a.convertFromBase(ForcePerLengthDecinewtonPerMillimeter)
	a.decinewtons_per_millimeterLazy = &decinewtons_per_millimeter
	return decinewtons_per_millimeter
}

// DecanewtonPerMillimeter returns the value in DecanewtonPerMillimeter.
func (a *ForcePerLength) DecanewtonsPerMillimeter() float64 {
	if a.decanewtons_per_millimeterLazy != nil {
		return *a.decanewtons_per_millimeterLazy
	}
	decanewtons_per_millimeter := a.convertFromBase(ForcePerLengthDecanewtonPerMillimeter)
	a.decanewtons_per_millimeterLazy = &decanewtons_per_millimeter
	return decanewtons_per_millimeter
}

// KilonewtonPerMillimeter returns the value in KilonewtonPerMillimeter.
func (a *ForcePerLength) KilonewtonsPerMillimeter() float64 {
	if a.kilonewtons_per_millimeterLazy != nil {
		return *a.kilonewtons_per_millimeterLazy
	}
	kilonewtons_per_millimeter := a.convertFromBase(ForcePerLengthKilonewtonPerMillimeter)
	a.kilonewtons_per_millimeterLazy = &kilonewtons_per_millimeter
	return kilonewtons_per_millimeter
}

// MeganewtonPerMillimeter returns the value in MeganewtonPerMillimeter.
func (a *ForcePerLength) MeganewtonsPerMillimeter() float64 {
	if a.meganewtons_per_millimeterLazy != nil {
		return *a.meganewtons_per_millimeterLazy
	}
	meganewtons_per_millimeter := a.convertFromBase(ForcePerLengthMeganewtonPerMillimeter)
	a.meganewtons_per_millimeterLazy = &meganewtons_per_millimeter
	return meganewtons_per_millimeter
}


// ToDto creates an ForcePerLengthDto representation.
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

// ToDtoJSON creates an ForcePerLengthDto representation.
func (a *ForcePerLength) ToDtoJSON(holdInUnit *ForcePerLengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ForcePerLength to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ForcePerLength) String() string {
	return a.ToString(ForcePerLengthNewtonPerMeter, 2)
}

// ToString formats the ForcePerLength to string.
// fractionalDigits -1 for not mention
func (a *ForcePerLength) ToString(unit ForcePerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ForcePerLength) getUnitAbbreviation(unit ForcePerLengthUnits) string {
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

// Check if the given ForcePerLength are equals to the current ForcePerLength
func (a *ForcePerLength) Equals(other *ForcePerLength) bool {
	return a.value == other.BaseValue()
}

// Check if the given ForcePerLength are equals to the current ForcePerLength
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

// Add the given ForcePerLength to the current ForcePerLength.
func (a *ForcePerLength) Add(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value + other.BaseValue()}
}

// Subtract the given ForcePerLength to the current ForcePerLength.
func (a *ForcePerLength) Subtract(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value - other.BaseValue()}
}

// Multiply the given ForcePerLength to the current ForcePerLength.
func (a *ForcePerLength) Multiply(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value * other.BaseValue()}
}

// Divide the given ForcePerLength to the current ForcePerLength.
func (a *ForcePerLength) Divide(other *ForcePerLength) *ForcePerLength {
	return &ForcePerLength{value: a.value / other.BaseValue()}
}