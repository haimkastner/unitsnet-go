// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TorquePerLengthUnits defines various units of TorquePerLength.
type TorquePerLengthUnits string

const (
    
        // 
        TorquePerLengthNewtonMillimeterPerMeter TorquePerLengthUnits = "NewtonMillimeterPerMeter"
        // 
        TorquePerLengthNewtonCentimeterPerMeter TorquePerLengthUnits = "NewtonCentimeterPerMeter"
        // 
        TorquePerLengthNewtonMeterPerMeter TorquePerLengthUnits = "NewtonMeterPerMeter"
        // 
        TorquePerLengthPoundForceInchPerFoot TorquePerLengthUnits = "PoundForceInchPerFoot"
        // 
        TorquePerLengthPoundForceFootPerFoot TorquePerLengthUnits = "PoundForceFootPerFoot"
        // 
        TorquePerLengthKilogramForceMillimeterPerMeter TorquePerLengthUnits = "KilogramForceMillimeterPerMeter"
        // 
        TorquePerLengthKilogramForceCentimeterPerMeter TorquePerLengthUnits = "KilogramForceCentimeterPerMeter"
        // 
        TorquePerLengthKilogramForceMeterPerMeter TorquePerLengthUnits = "KilogramForceMeterPerMeter"
        // 
        TorquePerLengthTonneForceMillimeterPerMeter TorquePerLengthUnits = "TonneForceMillimeterPerMeter"
        // 
        TorquePerLengthTonneForceCentimeterPerMeter TorquePerLengthUnits = "TonneForceCentimeterPerMeter"
        // 
        TorquePerLengthTonneForceMeterPerMeter TorquePerLengthUnits = "TonneForceMeterPerMeter"
        // 
        TorquePerLengthKilonewtonMillimeterPerMeter TorquePerLengthUnits = "KilonewtonMillimeterPerMeter"
        // 
        TorquePerLengthMeganewtonMillimeterPerMeter TorquePerLengthUnits = "MeganewtonMillimeterPerMeter"
        // 
        TorquePerLengthKilonewtonCentimeterPerMeter TorquePerLengthUnits = "KilonewtonCentimeterPerMeter"
        // 
        TorquePerLengthMeganewtonCentimeterPerMeter TorquePerLengthUnits = "MeganewtonCentimeterPerMeter"
        // 
        TorquePerLengthKilonewtonMeterPerMeter TorquePerLengthUnits = "KilonewtonMeterPerMeter"
        // 
        TorquePerLengthMeganewtonMeterPerMeter TorquePerLengthUnits = "MeganewtonMeterPerMeter"
        // 
        TorquePerLengthKilopoundForceInchPerFoot TorquePerLengthUnits = "KilopoundForceInchPerFoot"
        // 
        TorquePerLengthMegapoundForceInchPerFoot TorquePerLengthUnits = "MegapoundForceInchPerFoot"
        // 
        TorquePerLengthKilopoundForceFootPerFoot TorquePerLengthUnits = "KilopoundForceFootPerFoot"
        // 
        TorquePerLengthMegapoundForceFootPerFoot TorquePerLengthUnits = "MegapoundForceFootPerFoot"
)

// TorquePerLengthDto represents a TorquePerLength measurement with a numerical value and its corresponding unit.
type TorquePerLengthDto struct {
    // Value is the numerical representation of the TorquePerLength.
	Value float64
    // Unit specifies the unit of measurement for the TorquePerLength, as defined in the TorquePerLengthUnits enumeration.
	Unit  TorquePerLengthUnits
}

// TorquePerLengthDtoFactory groups methods for creating and serializing TorquePerLengthDto objects.
type TorquePerLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TorquePerLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TorquePerLengthDtoFactory) FromJSON(data []byte) (*TorquePerLengthDto, error) {
	a := TorquePerLengthDto{}

    // Parse JSON into TorquePerLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a TorquePerLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TorquePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// TorquePerLength represents a measurement in a of TorquePerLength.
//
// The magnitude of torque per unit length.
type TorquePerLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    newton_millimeters_per_meterLazy *float64 
    newton_centimeters_per_meterLazy *float64 
    newton_meters_per_meterLazy *float64 
    pound_force_inches_per_footLazy *float64 
    pound_force_feet_per_footLazy *float64 
    kilogram_force_millimeters_per_meterLazy *float64 
    kilogram_force_centimeters_per_meterLazy *float64 
    kilogram_force_meters_per_meterLazy *float64 
    tonne_force_millimeters_per_meterLazy *float64 
    tonne_force_centimeters_per_meterLazy *float64 
    tonne_force_meters_per_meterLazy *float64 
    kilonewton_millimeters_per_meterLazy *float64 
    meganewton_millimeters_per_meterLazy *float64 
    kilonewton_centimeters_per_meterLazy *float64 
    meganewton_centimeters_per_meterLazy *float64 
    kilonewton_meters_per_meterLazy *float64 
    meganewton_meters_per_meterLazy *float64 
    kilopound_force_inches_per_footLazy *float64 
    megapound_force_inches_per_footLazy *float64 
    kilopound_force_feet_per_footLazy *float64 
    megapound_force_feet_per_footLazy *float64 
}

// TorquePerLengthFactory groups methods for creating TorquePerLength instances.
type TorquePerLengthFactory struct{}

// CreateTorquePerLength creates a new TorquePerLength instance from the given value and unit.
func (uf TorquePerLengthFactory) CreateTorquePerLength(value float64, unit TorquePerLengthUnits) (*TorquePerLength, error) {
	return newTorquePerLength(value, unit)
}

// FromDto converts a TorquePerLengthDto to a TorquePerLength instance.
func (uf TorquePerLengthFactory) FromDto(dto TorquePerLengthDto) (*TorquePerLength, error) {
	return newTorquePerLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a TorquePerLength instance.
func (uf TorquePerLengthFactory) FromDtoJSON(data []byte) (*TorquePerLength, error) {
	unitDto, err := TorquePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TorquePerLengthDto from JSON: %w", err)
	}
	return TorquePerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonMillimetersPerMeter creates a new TorquePerLength instance from a value in NewtonMillimetersPerMeter.
func (uf TorquePerLengthFactory) FromNewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonMillimeterPerMeter)
}

// FromNewtonCentimetersPerMeter creates a new TorquePerLength instance from a value in NewtonCentimetersPerMeter.
func (uf TorquePerLengthFactory) FromNewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonCentimeterPerMeter)
}

// FromNewtonMetersPerMeter creates a new TorquePerLength instance from a value in NewtonMetersPerMeter.
func (uf TorquePerLengthFactory) FromNewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonMeterPerMeter)
}

// FromPoundForceInchesPerFoot creates a new TorquePerLength instance from a value in PoundForceInchesPerFoot.
func (uf TorquePerLengthFactory) FromPoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthPoundForceInchPerFoot)
}

// FromPoundForceFeetPerFoot creates a new TorquePerLength instance from a value in PoundForceFeetPerFoot.
func (uf TorquePerLengthFactory) FromPoundForceFeetPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthPoundForceFootPerFoot)
}

// FromKilogramForceMillimetersPerMeter creates a new TorquePerLength instance from a value in KilogramForceMillimetersPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceMillimeterPerMeter)
}

// FromKilogramForceCentimetersPerMeter creates a new TorquePerLength instance from a value in KilogramForceCentimetersPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceCentimeterPerMeter)
}

// FromKilogramForceMetersPerMeter creates a new TorquePerLength instance from a value in KilogramForceMetersPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceMeterPerMeter)
}

// FromTonneForceMillimetersPerMeter creates a new TorquePerLength instance from a value in TonneForceMillimetersPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceMillimeterPerMeter)
}

// FromTonneForceCentimetersPerMeter creates a new TorquePerLength instance from a value in TonneForceCentimetersPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceCentimeterPerMeter)
}

// FromTonneForceMetersPerMeter creates a new TorquePerLength instance from a value in TonneForceMetersPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceMeterPerMeter)
}

// FromKilonewtonMillimetersPerMeter creates a new TorquePerLength instance from a value in KilonewtonMillimetersPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonMillimeterPerMeter)
}

// FromMeganewtonMillimetersPerMeter creates a new TorquePerLength instance from a value in MeganewtonMillimetersPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonMillimeterPerMeter)
}

// FromKilonewtonCentimetersPerMeter creates a new TorquePerLength instance from a value in KilonewtonCentimetersPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonCentimeterPerMeter)
}

// FromMeganewtonCentimetersPerMeter creates a new TorquePerLength instance from a value in MeganewtonCentimetersPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonCentimeterPerMeter)
}

// FromKilonewtonMetersPerMeter creates a new TorquePerLength instance from a value in KilonewtonMetersPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonMeterPerMeter)
}

// FromMeganewtonMetersPerMeter creates a new TorquePerLength instance from a value in MeganewtonMetersPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonMeterPerMeter)
}

// FromKilopoundForceInchesPerFoot creates a new TorquePerLength instance from a value in KilopoundForceInchesPerFoot.
func (uf TorquePerLengthFactory) FromKilopoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilopoundForceInchPerFoot)
}

// FromMegapoundForceInchesPerFoot creates a new TorquePerLength instance from a value in MegapoundForceInchesPerFoot.
func (uf TorquePerLengthFactory) FromMegapoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMegapoundForceInchPerFoot)
}

// FromKilopoundForceFeetPerFoot creates a new TorquePerLength instance from a value in KilopoundForceFeetPerFoot.
func (uf TorquePerLengthFactory) FromKilopoundForceFeetPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilopoundForceFootPerFoot)
}

// FromMegapoundForceFeetPerFoot creates a new TorquePerLength instance from a value in MegapoundForceFeetPerFoot.
func (uf TorquePerLengthFactory) FromMegapoundForceFeetPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMegapoundForceFootPerFoot)
}


// newTorquePerLength creates a new TorquePerLength.
func newTorquePerLength(value float64, fromUnit TorquePerLengthUnits) (*TorquePerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TorquePerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TorquePerLength in NewtonMeterPerMeter unit (the base/default quantity).
func (a *TorquePerLength) BaseValue() float64 {
	return a.value
}


// NewtonMillimetersPerMeter returns the TorquePerLength value in NewtonMillimetersPerMeter.
//
// 
func (a *TorquePerLength) NewtonMillimetersPerMeter() float64 {
	if a.newton_millimeters_per_meterLazy != nil {
		return *a.newton_millimeters_per_meterLazy
	}
	newton_millimeters_per_meter := a.convertFromBase(TorquePerLengthNewtonMillimeterPerMeter)
	a.newton_millimeters_per_meterLazy = &newton_millimeters_per_meter
	return newton_millimeters_per_meter
}

// NewtonCentimetersPerMeter returns the TorquePerLength value in NewtonCentimetersPerMeter.
//
// 
func (a *TorquePerLength) NewtonCentimetersPerMeter() float64 {
	if a.newton_centimeters_per_meterLazy != nil {
		return *a.newton_centimeters_per_meterLazy
	}
	newton_centimeters_per_meter := a.convertFromBase(TorquePerLengthNewtonCentimeterPerMeter)
	a.newton_centimeters_per_meterLazy = &newton_centimeters_per_meter
	return newton_centimeters_per_meter
}

// NewtonMetersPerMeter returns the TorquePerLength value in NewtonMetersPerMeter.
//
// 
func (a *TorquePerLength) NewtonMetersPerMeter() float64 {
	if a.newton_meters_per_meterLazy != nil {
		return *a.newton_meters_per_meterLazy
	}
	newton_meters_per_meter := a.convertFromBase(TorquePerLengthNewtonMeterPerMeter)
	a.newton_meters_per_meterLazy = &newton_meters_per_meter
	return newton_meters_per_meter
}

// PoundForceInchesPerFoot returns the TorquePerLength value in PoundForceInchesPerFoot.
//
// 
func (a *TorquePerLength) PoundForceInchesPerFoot() float64 {
	if a.pound_force_inches_per_footLazy != nil {
		return *a.pound_force_inches_per_footLazy
	}
	pound_force_inches_per_foot := a.convertFromBase(TorquePerLengthPoundForceInchPerFoot)
	a.pound_force_inches_per_footLazy = &pound_force_inches_per_foot
	return pound_force_inches_per_foot
}

// PoundForceFeetPerFoot returns the TorquePerLength value in PoundForceFeetPerFoot.
//
// 
func (a *TorquePerLength) PoundForceFeetPerFoot() float64 {
	if a.pound_force_feet_per_footLazy != nil {
		return *a.pound_force_feet_per_footLazy
	}
	pound_force_feet_per_foot := a.convertFromBase(TorquePerLengthPoundForceFootPerFoot)
	a.pound_force_feet_per_footLazy = &pound_force_feet_per_foot
	return pound_force_feet_per_foot
}

// KilogramForceMillimetersPerMeter returns the TorquePerLength value in KilogramForceMillimetersPerMeter.
//
// 
func (a *TorquePerLength) KilogramForceMillimetersPerMeter() float64 {
	if a.kilogram_force_millimeters_per_meterLazy != nil {
		return *a.kilogram_force_millimeters_per_meterLazy
	}
	kilogram_force_millimeters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceMillimeterPerMeter)
	a.kilogram_force_millimeters_per_meterLazy = &kilogram_force_millimeters_per_meter
	return kilogram_force_millimeters_per_meter
}

// KilogramForceCentimetersPerMeter returns the TorquePerLength value in KilogramForceCentimetersPerMeter.
//
// 
func (a *TorquePerLength) KilogramForceCentimetersPerMeter() float64 {
	if a.kilogram_force_centimeters_per_meterLazy != nil {
		return *a.kilogram_force_centimeters_per_meterLazy
	}
	kilogram_force_centimeters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceCentimeterPerMeter)
	a.kilogram_force_centimeters_per_meterLazy = &kilogram_force_centimeters_per_meter
	return kilogram_force_centimeters_per_meter
}

// KilogramForceMetersPerMeter returns the TorquePerLength value in KilogramForceMetersPerMeter.
//
// 
func (a *TorquePerLength) KilogramForceMetersPerMeter() float64 {
	if a.kilogram_force_meters_per_meterLazy != nil {
		return *a.kilogram_force_meters_per_meterLazy
	}
	kilogram_force_meters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceMeterPerMeter)
	a.kilogram_force_meters_per_meterLazy = &kilogram_force_meters_per_meter
	return kilogram_force_meters_per_meter
}

// TonneForceMillimetersPerMeter returns the TorquePerLength value in TonneForceMillimetersPerMeter.
//
// 
func (a *TorquePerLength) TonneForceMillimetersPerMeter() float64 {
	if a.tonne_force_millimeters_per_meterLazy != nil {
		return *a.tonne_force_millimeters_per_meterLazy
	}
	tonne_force_millimeters_per_meter := a.convertFromBase(TorquePerLengthTonneForceMillimeterPerMeter)
	a.tonne_force_millimeters_per_meterLazy = &tonne_force_millimeters_per_meter
	return tonne_force_millimeters_per_meter
}

// TonneForceCentimetersPerMeter returns the TorquePerLength value in TonneForceCentimetersPerMeter.
//
// 
func (a *TorquePerLength) TonneForceCentimetersPerMeter() float64 {
	if a.tonne_force_centimeters_per_meterLazy != nil {
		return *a.tonne_force_centimeters_per_meterLazy
	}
	tonne_force_centimeters_per_meter := a.convertFromBase(TorquePerLengthTonneForceCentimeterPerMeter)
	a.tonne_force_centimeters_per_meterLazy = &tonne_force_centimeters_per_meter
	return tonne_force_centimeters_per_meter
}

// TonneForceMetersPerMeter returns the TorquePerLength value in TonneForceMetersPerMeter.
//
// 
func (a *TorquePerLength) TonneForceMetersPerMeter() float64 {
	if a.tonne_force_meters_per_meterLazy != nil {
		return *a.tonne_force_meters_per_meterLazy
	}
	tonne_force_meters_per_meter := a.convertFromBase(TorquePerLengthTonneForceMeterPerMeter)
	a.tonne_force_meters_per_meterLazy = &tonne_force_meters_per_meter
	return tonne_force_meters_per_meter
}

// KilonewtonMillimetersPerMeter returns the TorquePerLength value in KilonewtonMillimetersPerMeter.
//
// 
func (a *TorquePerLength) KilonewtonMillimetersPerMeter() float64 {
	if a.kilonewton_millimeters_per_meterLazy != nil {
		return *a.kilonewton_millimeters_per_meterLazy
	}
	kilonewton_millimeters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonMillimeterPerMeter)
	a.kilonewton_millimeters_per_meterLazy = &kilonewton_millimeters_per_meter
	return kilonewton_millimeters_per_meter
}

// MeganewtonMillimetersPerMeter returns the TorquePerLength value in MeganewtonMillimetersPerMeter.
//
// 
func (a *TorquePerLength) MeganewtonMillimetersPerMeter() float64 {
	if a.meganewton_millimeters_per_meterLazy != nil {
		return *a.meganewton_millimeters_per_meterLazy
	}
	meganewton_millimeters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonMillimeterPerMeter)
	a.meganewton_millimeters_per_meterLazy = &meganewton_millimeters_per_meter
	return meganewton_millimeters_per_meter
}

// KilonewtonCentimetersPerMeter returns the TorquePerLength value in KilonewtonCentimetersPerMeter.
//
// 
func (a *TorquePerLength) KilonewtonCentimetersPerMeter() float64 {
	if a.kilonewton_centimeters_per_meterLazy != nil {
		return *a.kilonewton_centimeters_per_meterLazy
	}
	kilonewton_centimeters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonCentimeterPerMeter)
	a.kilonewton_centimeters_per_meterLazy = &kilonewton_centimeters_per_meter
	return kilonewton_centimeters_per_meter
}

// MeganewtonCentimetersPerMeter returns the TorquePerLength value in MeganewtonCentimetersPerMeter.
//
// 
func (a *TorquePerLength) MeganewtonCentimetersPerMeter() float64 {
	if a.meganewton_centimeters_per_meterLazy != nil {
		return *a.meganewton_centimeters_per_meterLazy
	}
	meganewton_centimeters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonCentimeterPerMeter)
	a.meganewton_centimeters_per_meterLazy = &meganewton_centimeters_per_meter
	return meganewton_centimeters_per_meter
}

// KilonewtonMetersPerMeter returns the TorquePerLength value in KilonewtonMetersPerMeter.
//
// 
func (a *TorquePerLength) KilonewtonMetersPerMeter() float64 {
	if a.kilonewton_meters_per_meterLazy != nil {
		return *a.kilonewton_meters_per_meterLazy
	}
	kilonewton_meters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonMeterPerMeter)
	a.kilonewton_meters_per_meterLazy = &kilonewton_meters_per_meter
	return kilonewton_meters_per_meter
}

// MeganewtonMetersPerMeter returns the TorquePerLength value in MeganewtonMetersPerMeter.
//
// 
func (a *TorquePerLength) MeganewtonMetersPerMeter() float64 {
	if a.meganewton_meters_per_meterLazy != nil {
		return *a.meganewton_meters_per_meterLazy
	}
	meganewton_meters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonMeterPerMeter)
	a.meganewton_meters_per_meterLazy = &meganewton_meters_per_meter
	return meganewton_meters_per_meter
}

// KilopoundForceInchesPerFoot returns the TorquePerLength value in KilopoundForceInchesPerFoot.
//
// 
func (a *TorquePerLength) KilopoundForceInchesPerFoot() float64 {
	if a.kilopound_force_inches_per_footLazy != nil {
		return *a.kilopound_force_inches_per_footLazy
	}
	kilopound_force_inches_per_foot := a.convertFromBase(TorquePerLengthKilopoundForceInchPerFoot)
	a.kilopound_force_inches_per_footLazy = &kilopound_force_inches_per_foot
	return kilopound_force_inches_per_foot
}

// MegapoundForceInchesPerFoot returns the TorquePerLength value in MegapoundForceInchesPerFoot.
//
// 
func (a *TorquePerLength) MegapoundForceInchesPerFoot() float64 {
	if a.megapound_force_inches_per_footLazy != nil {
		return *a.megapound_force_inches_per_footLazy
	}
	megapound_force_inches_per_foot := a.convertFromBase(TorquePerLengthMegapoundForceInchPerFoot)
	a.megapound_force_inches_per_footLazy = &megapound_force_inches_per_foot
	return megapound_force_inches_per_foot
}

// KilopoundForceFeetPerFoot returns the TorquePerLength value in KilopoundForceFeetPerFoot.
//
// 
func (a *TorquePerLength) KilopoundForceFeetPerFoot() float64 {
	if a.kilopound_force_feet_per_footLazy != nil {
		return *a.kilopound_force_feet_per_footLazy
	}
	kilopound_force_feet_per_foot := a.convertFromBase(TorquePerLengthKilopoundForceFootPerFoot)
	a.kilopound_force_feet_per_footLazy = &kilopound_force_feet_per_foot
	return kilopound_force_feet_per_foot
}

// MegapoundForceFeetPerFoot returns the TorquePerLength value in MegapoundForceFeetPerFoot.
//
// 
func (a *TorquePerLength) MegapoundForceFeetPerFoot() float64 {
	if a.megapound_force_feet_per_footLazy != nil {
		return *a.megapound_force_feet_per_footLazy
	}
	megapound_force_feet_per_foot := a.convertFromBase(TorquePerLengthMegapoundForceFootPerFoot)
	a.megapound_force_feet_per_footLazy = &megapound_force_feet_per_foot
	return megapound_force_feet_per_foot
}


// ToDto creates a TorquePerLengthDto representation from the TorquePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerMeter by default.
func (a *TorquePerLength) ToDto(holdInUnit *TorquePerLengthUnits) TorquePerLengthDto {
	if holdInUnit == nil {
		defaultUnit := TorquePerLengthNewtonMeterPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return TorquePerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the TorquePerLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeterPerMeter by default.
func (a *TorquePerLength) ToDtoJSON(holdInUnit *TorquePerLengthUnits) ([]byte, error) {
	// Convert to TorquePerLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a TorquePerLength to a specific unit value.
// The function uses the provided unit type (TorquePerLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *TorquePerLength) Convert(toUnit TorquePerLengthUnits) float64 {
	switch toUnit { 
    case TorquePerLengthNewtonMillimeterPerMeter:
		return a.NewtonMillimetersPerMeter()
    case TorquePerLengthNewtonCentimeterPerMeter:
		return a.NewtonCentimetersPerMeter()
    case TorquePerLengthNewtonMeterPerMeter:
		return a.NewtonMetersPerMeter()
    case TorquePerLengthPoundForceInchPerFoot:
		return a.PoundForceInchesPerFoot()
    case TorquePerLengthPoundForceFootPerFoot:
		return a.PoundForceFeetPerFoot()
    case TorquePerLengthKilogramForceMillimeterPerMeter:
		return a.KilogramForceMillimetersPerMeter()
    case TorquePerLengthKilogramForceCentimeterPerMeter:
		return a.KilogramForceCentimetersPerMeter()
    case TorquePerLengthKilogramForceMeterPerMeter:
		return a.KilogramForceMetersPerMeter()
    case TorquePerLengthTonneForceMillimeterPerMeter:
		return a.TonneForceMillimetersPerMeter()
    case TorquePerLengthTonneForceCentimeterPerMeter:
		return a.TonneForceCentimetersPerMeter()
    case TorquePerLengthTonneForceMeterPerMeter:
		return a.TonneForceMetersPerMeter()
    case TorquePerLengthKilonewtonMillimeterPerMeter:
		return a.KilonewtonMillimetersPerMeter()
    case TorquePerLengthMeganewtonMillimeterPerMeter:
		return a.MeganewtonMillimetersPerMeter()
    case TorquePerLengthKilonewtonCentimeterPerMeter:
		return a.KilonewtonCentimetersPerMeter()
    case TorquePerLengthMeganewtonCentimeterPerMeter:
		return a.MeganewtonCentimetersPerMeter()
    case TorquePerLengthKilonewtonMeterPerMeter:
		return a.KilonewtonMetersPerMeter()
    case TorquePerLengthMeganewtonMeterPerMeter:
		return a.MeganewtonMetersPerMeter()
    case TorquePerLengthKilopoundForceInchPerFoot:
		return a.KilopoundForceInchesPerFoot()
    case TorquePerLengthMegapoundForceInchPerFoot:
		return a.MegapoundForceInchesPerFoot()
    case TorquePerLengthKilopoundForceFootPerFoot:
		return a.KilopoundForceFeetPerFoot()
    case TorquePerLengthMegapoundForceFootPerFoot:
		return a.MegapoundForceFeetPerFoot()
	default:
		return math.NaN()
	}
}

func (a *TorquePerLength) convertFromBase(toUnit TorquePerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case TorquePerLengthNewtonMillimeterPerMeter:
		return (value * 1000) 
	case TorquePerLengthNewtonCentimeterPerMeter:
		return (value * 100) 
	case TorquePerLengthNewtonMeterPerMeter:
		return (value) 
	case TorquePerLengthPoundForceInchPerFoot:
		return (value / 0.370685147638) 
	case TorquePerLengthPoundForceFootPerFoot:
		return (value / 4.44822161526) 
	case TorquePerLengthKilogramForceMillimeterPerMeter:
		return (value * 101.971619222242) 
	case TorquePerLengthKilogramForceCentimeterPerMeter:
		return (value * 10.1971619222242) 
	case TorquePerLengthKilogramForceMeterPerMeter:
		return (value * 0.101971619222242) 
	case TorquePerLengthTonneForceMillimeterPerMeter:
		return (value * 0.101971619222242) 
	case TorquePerLengthTonneForceCentimeterPerMeter:
		return (value * 0.0101971619222242) 
	case TorquePerLengthTonneForceMeterPerMeter:
		return (value * 0.000101971619222242) 
	case TorquePerLengthKilonewtonMillimeterPerMeter:
		return ((value * 1000) / 1000.0) 
	case TorquePerLengthMeganewtonMillimeterPerMeter:
		return ((value * 1000) / 1000000.0) 
	case TorquePerLengthKilonewtonCentimeterPerMeter:
		return ((value * 100) / 1000.0) 
	case TorquePerLengthMeganewtonCentimeterPerMeter:
		return ((value * 100) / 1000000.0) 
	case TorquePerLengthKilonewtonMeterPerMeter:
		return ((value) / 1000.0) 
	case TorquePerLengthMeganewtonMeterPerMeter:
		return ((value) / 1000000.0) 
	case TorquePerLengthKilopoundForceInchPerFoot:
		return ((value / 0.370685147638) / 1000.0) 
	case TorquePerLengthMegapoundForceInchPerFoot:
		return ((value / 0.370685147638) / 1000000.0) 
	case TorquePerLengthKilopoundForceFootPerFoot:
		return ((value / 4.44822161526) / 1000.0) 
	case TorquePerLengthMegapoundForceFootPerFoot:
		return ((value / 4.44822161526) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *TorquePerLength) convertToBase(value float64, fromUnit TorquePerLengthUnits) float64 {
	switch fromUnit { 
	case TorquePerLengthNewtonMillimeterPerMeter:
		return (value * 0.001) 
	case TorquePerLengthNewtonCentimeterPerMeter:
		return (value * 0.01) 
	case TorquePerLengthNewtonMeterPerMeter:
		return (value) 
	case TorquePerLengthPoundForceInchPerFoot:
		return (value * 0.370685147638) 
	case TorquePerLengthPoundForceFootPerFoot:
		return (value * 4.44822161526) 
	case TorquePerLengthKilogramForceMillimeterPerMeter:
		return (value * 0.00980665019960652) 
	case TorquePerLengthKilogramForceCentimeterPerMeter:
		return (value * 0.0980665019960652) 
	case TorquePerLengthKilogramForceMeterPerMeter:
		return (value * 9.80665019960652) 
	case TorquePerLengthTonneForceMillimeterPerMeter:
		return (value * 9.80665019960652) 
	case TorquePerLengthTonneForceCentimeterPerMeter:
		return (value * 98.0665019960652) 
	case TorquePerLengthTonneForceMeterPerMeter:
		return (value * 9806.65019960653) 
	case TorquePerLengthKilonewtonMillimeterPerMeter:
		return ((value * 0.001) * 1000.0) 
	case TorquePerLengthMeganewtonMillimeterPerMeter:
		return ((value * 0.001) * 1000000.0) 
	case TorquePerLengthKilonewtonCentimeterPerMeter:
		return ((value * 0.01) * 1000.0) 
	case TorquePerLengthMeganewtonCentimeterPerMeter:
		return ((value * 0.01) * 1000000.0) 
	case TorquePerLengthKilonewtonMeterPerMeter:
		return ((value) * 1000.0) 
	case TorquePerLengthMeganewtonMeterPerMeter:
		return ((value) * 1000000.0) 
	case TorquePerLengthKilopoundForceInchPerFoot:
		return ((value * 0.370685147638) * 1000.0) 
	case TorquePerLengthMegapoundForceInchPerFoot:
		return ((value * 0.370685147638) * 1000000.0) 
	case TorquePerLengthKilopoundForceFootPerFoot:
		return ((value * 4.44822161526) * 1000.0) 
	case TorquePerLengthMegapoundForceFootPerFoot:
		return ((value * 4.44822161526) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the TorquePerLength in the default unit (NewtonMeterPerMeter),
// formatted to two decimal places.
func (a TorquePerLength) String() string {
	return a.ToString(TorquePerLengthNewtonMeterPerMeter, 2)
}

// ToString formats the TorquePerLength value as a string with the specified unit and fractional digits.
// It converts the TorquePerLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the TorquePerLength value will be converted (e.g., NewtonMeterPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the TorquePerLength with the unit abbreviation.
func (a *TorquePerLength) ToString(unit TorquePerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TorquePerLength) getUnitAbbreviation(unit TorquePerLengthUnits) string {
	switch unit { 
	case TorquePerLengthNewtonMillimeterPerMeter:
		return "N·mm/m" 
	case TorquePerLengthNewtonCentimeterPerMeter:
		return "N·cm/m" 
	case TorquePerLengthNewtonMeterPerMeter:
		return "N·m/m" 
	case TorquePerLengthPoundForceInchPerFoot:
		return "lbf·in/ft" 
	case TorquePerLengthPoundForceFootPerFoot:
		return "lbf·ft/ft" 
	case TorquePerLengthKilogramForceMillimeterPerMeter:
		return "kgf·mm/m" 
	case TorquePerLengthKilogramForceCentimeterPerMeter:
		return "kgf·cm/m" 
	case TorquePerLengthKilogramForceMeterPerMeter:
		return "kgf·m/m" 
	case TorquePerLengthTonneForceMillimeterPerMeter:
		return "tf·mm/m" 
	case TorquePerLengthTonneForceCentimeterPerMeter:
		return "tf·cm/m" 
	case TorquePerLengthTonneForceMeterPerMeter:
		return "tf·m/m" 
	case TorquePerLengthKilonewtonMillimeterPerMeter:
		return "kN·mm/m" 
	case TorquePerLengthMeganewtonMillimeterPerMeter:
		return "MN·mm/m" 
	case TorquePerLengthKilonewtonCentimeterPerMeter:
		return "kN·cm/m" 
	case TorquePerLengthMeganewtonCentimeterPerMeter:
		return "MN·cm/m" 
	case TorquePerLengthKilonewtonMeterPerMeter:
		return "kN·m/m" 
	case TorquePerLengthMeganewtonMeterPerMeter:
		return "MN·m/m" 
	case TorquePerLengthKilopoundForceInchPerFoot:
		return "klbf·in/ft" 
	case TorquePerLengthMegapoundForceInchPerFoot:
		return "Mlbf·in/ft" 
	case TorquePerLengthKilopoundForceFootPerFoot:
		return "klbf·ft/ft" 
	case TorquePerLengthMegapoundForceFootPerFoot:
		return "Mlbf·ft/ft" 
	default:
		return ""
	}
}

// Equals checks if the given TorquePerLength is equal to the current TorquePerLength.
//
// Parameters:
//    other: The TorquePerLength to compare against.
//
// Returns:
//    bool: Returns true if both TorquePerLength are equal, false otherwise.
func (a *TorquePerLength) Equals(other *TorquePerLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current TorquePerLength with another TorquePerLength.
// It returns -1 if the current TorquePerLength is less than the other TorquePerLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The TorquePerLength to compare against.
//
// Returns:
//    int: -1 if the current TorquePerLength is less, 1 if greater, and 0 if equal.
func (a *TorquePerLength) CompareTo(other *TorquePerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given TorquePerLength to the current TorquePerLength and returns the result.
// The result is a new TorquePerLength instance with the sum of the values.
//
// Parameters:
//    other: The TorquePerLength to add to the current TorquePerLength.
//
// Returns:
//    *TorquePerLength: A new TorquePerLength instance representing the sum of both TorquePerLength.
func (a *TorquePerLength) Add(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given TorquePerLength from the current TorquePerLength and returns the result.
// The result is a new TorquePerLength instance with the difference of the values.
//
// Parameters:
//    other: The TorquePerLength to subtract from the current TorquePerLength.
//
// Returns:
//    *TorquePerLength: A new TorquePerLength instance representing the difference of both TorquePerLength.
func (a *TorquePerLength) Subtract(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current TorquePerLength by the given TorquePerLength and returns the result.
// The result is a new TorquePerLength instance with the product of the values.
//
// Parameters:
//    other: The TorquePerLength to multiply with the current TorquePerLength.
//
// Returns:
//    *TorquePerLength: A new TorquePerLength instance representing the product of both TorquePerLength.
func (a *TorquePerLength) Multiply(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value * other.BaseValue()}
}

// Divide divides the current TorquePerLength by the given TorquePerLength and returns the result.
// The result is a new TorquePerLength instance with the quotient of the values.
//
// Parameters:
//    other: The TorquePerLength to divide the current TorquePerLength by.
//
// Returns:
//    *TorquePerLength: A new TorquePerLength instance representing the quotient of both TorquePerLength.
func (a *TorquePerLength) Divide(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value / other.BaseValue()}
}