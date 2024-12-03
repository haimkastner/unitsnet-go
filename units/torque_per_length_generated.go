// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TorquePerLengthUnits enumeration
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

// TorquePerLengthDto represents an TorquePerLength
type TorquePerLengthDto struct {
	Value float64
	Unit  TorquePerLengthUnits
}

// TorquePerLengthDtoFactory struct to group related functions
type TorquePerLengthDtoFactory struct{}

func (udf TorquePerLengthDtoFactory) FromJSON(data []byte) (*TorquePerLengthDto, error) {
	a := TorquePerLengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TorquePerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// TorquePerLength struct
type TorquePerLength struct {
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

// TorquePerLengthFactory struct to group related functions
type TorquePerLengthFactory struct{}

func (uf TorquePerLengthFactory) CreateTorquePerLength(value float64, unit TorquePerLengthUnits) (*TorquePerLength, error) {
	return newTorquePerLength(value, unit)
}

func (uf TorquePerLengthFactory) FromDto(dto TorquePerLengthDto) (*TorquePerLength, error) {
	return newTorquePerLength(dto.Value, dto.Unit)
}

func (uf TorquePerLengthFactory) FromDtoJSON(data []byte) (*TorquePerLength, error) {
	unitDto, err := TorquePerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TorquePerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonMillimeterPerMeter creates a new TorquePerLength instance from NewtonMillimeterPerMeter.
func (uf TorquePerLengthFactory) FromNewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonMillimeterPerMeter)
}

// FromNewtonCentimeterPerMeter creates a new TorquePerLength instance from NewtonCentimeterPerMeter.
func (uf TorquePerLengthFactory) FromNewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonCentimeterPerMeter)
}

// FromNewtonMeterPerMeter creates a new TorquePerLength instance from NewtonMeterPerMeter.
func (uf TorquePerLengthFactory) FromNewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthNewtonMeterPerMeter)
}

// FromPoundForceInchPerFoot creates a new TorquePerLength instance from PoundForceInchPerFoot.
func (uf TorquePerLengthFactory) FromPoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthPoundForceInchPerFoot)
}

// FromPoundForceFootPerFoot creates a new TorquePerLength instance from PoundForceFootPerFoot.
func (uf TorquePerLengthFactory) FromPoundForceFeetPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthPoundForceFootPerFoot)
}

// FromKilogramForceMillimeterPerMeter creates a new TorquePerLength instance from KilogramForceMillimeterPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceMillimeterPerMeter)
}

// FromKilogramForceCentimeterPerMeter creates a new TorquePerLength instance from KilogramForceCentimeterPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceCentimeterPerMeter)
}

// FromKilogramForceMeterPerMeter creates a new TorquePerLength instance from KilogramForceMeterPerMeter.
func (uf TorquePerLengthFactory) FromKilogramForceMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilogramForceMeterPerMeter)
}

// FromTonneForceMillimeterPerMeter creates a new TorquePerLength instance from TonneForceMillimeterPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceMillimeterPerMeter)
}

// FromTonneForceCentimeterPerMeter creates a new TorquePerLength instance from TonneForceCentimeterPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceCentimeterPerMeter)
}

// FromTonneForceMeterPerMeter creates a new TorquePerLength instance from TonneForceMeterPerMeter.
func (uf TorquePerLengthFactory) FromTonneForceMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthTonneForceMeterPerMeter)
}

// FromKilonewtonMillimeterPerMeter creates a new TorquePerLength instance from KilonewtonMillimeterPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonMillimeterPerMeter)
}

// FromMeganewtonMillimeterPerMeter creates a new TorquePerLength instance from MeganewtonMillimeterPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonMillimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonMillimeterPerMeter)
}

// FromKilonewtonCentimeterPerMeter creates a new TorquePerLength instance from KilonewtonCentimeterPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonCentimeterPerMeter)
}

// FromMeganewtonCentimeterPerMeter creates a new TorquePerLength instance from MeganewtonCentimeterPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonCentimetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonCentimeterPerMeter)
}

// FromKilonewtonMeterPerMeter creates a new TorquePerLength instance from KilonewtonMeterPerMeter.
func (uf TorquePerLengthFactory) FromKilonewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilonewtonMeterPerMeter)
}

// FromMeganewtonMeterPerMeter creates a new TorquePerLength instance from MeganewtonMeterPerMeter.
func (uf TorquePerLengthFactory) FromMeganewtonMetersPerMeter(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMeganewtonMeterPerMeter)
}

// FromKilopoundForceInchPerFoot creates a new TorquePerLength instance from KilopoundForceInchPerFoot.
func (uf TorquePerLengthFactory) FromKilopoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilopoundForceInchPerFoot)
}

// FromMegapoundForceInchPerFoot creates a new TorquePerLength instance from MegapoundForceInchPerFoot.
func (uf TorquePerLengthFactory) FromMegapoundForceInchesPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthMegapoundForceInchPerFoot)
}

// FromKilopoundForceFootPerFoot creates a new TorquePerLength instance from KilopoundForceFootPerFoot.
func (uf TorquePerLengthFactory) FromKilopoundForceFeetPerFoot(value float64) (*TorquePerLength, error) {
	return newTorquePerLength(value, TorquePerLengthKilopoundForceFootPerFoot)
}

// FromMegapoundForceFootPerFoot creates a new TorquePerLength instance from MegapoundForceFootPerFoot.
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

// BaseValue returns the base value of TorquePerLength in NewtonMeterPerMeter.
func (a *TorquePerLength) BaseValue() float64 {
	return a.value
}


// NewtonMillimeterPerMeter returns the value in NewtonMillimeterPerMeter.
func (a *TorquePerLength) NewtonMillimetersPerMeter() float64 {
	if a.newton_millimeters_per_meterLazy != nil {
		return *a.newton_millimeters_per_meterLazy
	}
	newton_millimeters_per_meter := a.convertFromBase(TorquePerLengthNewtonMillimeterPerMeter)
	a.newton_millimeters_per_meterLazy = &newton_millimeters_per_meter
	return newton_millimeters_per_meter
}

// NewtonCentimeterPerMeter returns the value in NewtonCentimeterPerMeter.
func (a *TorquePerLength) NewtonCentimetersPerMeter() float64 {
	if a.newton_centimeters_per_meterLazy != nil {
		return *a.newton_centimeters_per_meterLazy
	}
	newton_centimeters_per_meter := a.convertFromBase(TorquePerLengthNewtonCentimeterPerMeter)
	a.newton_centimeters_per_meterLazy = &newton_centimeters_per_meter
	return newton_centimeters_per_meter
}

// NewtonMeterPerMeter returns the value in NewtonMeterPerMeter.
func (a *TorquePerLength) NewtonMetersPerMeter() float64 {
	if a.newton_meters_per_meterLazy != nil {
		return *a.newton_meters_per_meterLazy
	}
	newton_meters_per_meter := a.convertFromBase(TorquePerLengthNewtonMeterPerMeter)
	a.newton_meters_per_meterLazy = &newton_meters_per_meter
	return newton_meters_per_meter
}

// PoundForceInchPerFoot returns the value in PoundForceInchPerFoot.
func (a *TorquePerLength) PoundForceInchesPerFoot() float64 {
	if a.pound_force_inches_per_footLazy != nil {
		return *a.pound_force_inches_per_footLazy
	}
	pound_force_inches_per_foot := a.convertFromBase(TorquePerLengthPoundForceInchPerFoot)
	a.pound_force_inches_per_footLazy = &pound_force_inches_per_foot
	return pound_force_inches_per_foot
}

// PoundForceFootPerFoot returns the value in PoundForceFootPerFoot.
func (a *TorquePerLength) PoundForceFeetPerFoot() float64 {
	if a.pound_force_feet_per_footLazy != nil {
		return *a.pound_force_feet_per_footLazy
	}
	pound_force_feet_per_foot := a.convertFromBase(TorquePerLengthPoundForceFootPerFoot)
	a.pound_force_feet_per_footLazy = &pound_force_feet_per_foot
	return pound_force_feet_per_foot
}

// KilogramForceMillimeterPerMeter returns the value in KilogramForceMillimeterPerMeter.
func (a *TorquePerLength) KilogramForceMillimetersPerMeter() float64 {
	if a.kilogram_force_millimeters_per_meterLazy != nil {
		return *a.kilogram_force_millimeters_per_meterLazy
	}
	kilogram_force_millimeters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceMillimeterPerMeter)
	a.kilogram_force_millimeters_per_meterLazy = &kilogram_force_millimeters_per_meter
	return kilogram_force_millimeters_per_meter
}

// KilogramForceCentimeterPerMeter returns the value in KilogramForceCentimeterPerMeter.
func (a *TorquePerLength) KilogramForceCentimetersPerMeter() float64 {
	if a.kilogram_force_centimeters_per_meterLazy != nil {
		return *a.kilogram_force_centimeters_per_meterLazy
	}
	kilogram_force_centimeters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceCentimeterPerMeter)
	a.kilogram_force_centimeters_per_meterLazy = &kilogram_force_centimeters_per_meter
	return kilogram_force_centimeters_per_meter
}

// KilogramForceMeterPerMeter returns the value in KilogramForceMeterPerMeter.
func (a *TorquePerLength) KilogramForceMetersPerMeter() float64 {
	if a.kilogram_force_meters_per_meterLazy != nil {
		return *a.kilogram_force_meters_per_meterLazy
	}
	kilogram_force_meters_per_meter := a.convertFromBase(TorquePerLengthKilogramForceMeterPerMeter)
	a.kilogram_force_meters_per_meterLazy = &kilogram_force_meters_per_meter
	return kilogram_force_meters_per_meter
}

// TonneForceMillimeterPerMeter returns the value in TonneForceMillimeterPerMeter.
func (a *TorquePerLength) TonneForceMillimetersPerMeter() float64 {
	if a.tonne_force_millimeters_per_meterLazy != nil {
		return *a.tonne_force_millimeters_per_meterLazy
	}
	tonne_force_millimeters_per_meter := a.convertFromBase(TorquePerLengthTonneForceMillimeterPerMeter)
	a.tonne_force_millimeters_per_meterLazy = &tonne_force_millimeters_per_meter
	return tonne_force_millimeters_per_meter
}

// TonneForceCentimeterPerMeter returns the value in TonneForceCentimeterPerMeter.
func (a *TorquePerLength) TonneForceCentimetersPerMeter() float64 {
	if a.tonne_force_centimeters_per_meterLazy != nil {
		return *a.tonne_force_centimeters_per_meterLazy
	}
	tonne_force_centimeters_per_meter := a.convertFromBase(TorquePerLengthTonneForceCentimeterPerMeter)
	a.tonne_force_centimeters_per_meterLazy = &tonne_force_centimeters_per_meter
	return tonne_force_centimeters_per_meter
}

// TonneForceMeterPerMeter returns the value in TonneForceMeterPerMeter.
func (a *TorquePerLength) TonneForceMetersPerMeter() float64 {
	if a.tonne_force_meters_per_meterLazy != nil {
		return *a.tonne_force_meters_per_meterLazy
	}
	tonne_force_meters_per_meter := a.convertFromBase(TorquePerLengthTonneForceMeterPerMeter)
	a.tonne_force_meters_per_meterLazy = &tonne_force_meters_per_meter
	return tonne_force_meters_per_meter
}

// KilonewtonMillimeterPerMeter returns the value in KilonewtonMillimeterPerMeter.
func (a *TorquePerLength) KilonewtonMillimetersPerMeter() float64 {
	if a.kilonewton_millimeters_per_meterLazy != nil {
		return *a.kilonewton_millimeters_per_meterLazy
	}
	kilonewton_millimeters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonMillimeterPerMeter)
	a.kilonewton_millimeters_per_meterLazy = &kilonewton_millimeters_per_meter
	return kilonewton_millimeters_per_meter
}

// MeganewtonMillimeterPerMeter returns the value in MeganewtonMillimeterPerMeter.
func (a *TorquePerLength) MeganewtonMillimetersPerMeter() float64 {
	if a.meganewton_millimeters_per_meterLazy != nil {
		return *a.meganewton_millimeters_per_meterLazy
	}
	meganewton_millimeters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonMillimeterPerMeter)
	a.meganewton_millimeters_per_meterLazy = &meganewton_millimeters_per_meter
	return meganewton_millimeters_per_meter
}

// KilonewtonCentimeterPerMeter returns the value in KilonewtonCentimeterPerMeter.
func (a *TorquePerLength) KilonewtonCentimetersPerMeter() float64 {
	if a.kilonewton_centimeters_per_meterLazy != nil {
		return *a.kilonewton_centimeters_per_meterLazy
	}
	kilonewton_centimeters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonCentimeterPerMeter)
	a.kilonewton_centimeters_per_meterLazy = &kilonewton_centimeters_per_meter
	return kilonewton_centimeters_per_meter
}

// MeganewtonCentimeterPerMeter returns the value in MeganewtonCentimeterPerMeter.
func (a *TorquePerLength) MeganewtonCentimetersPerMeter() float64 {
	if a.meganewton_centimeters_per_meterLazy != nil {
		return *a.meganewton_centimeters_per_meterLazy
	}
	meganewton_centimeters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonCentimeterPerMeter)
	a.meganewton_centimeters_per_meterLazy = &meganewton_centimeters_per_meter
	return meganewton_centimeters_per_meter
}

// KilonewtonMeterPerMeter returns the value in KilonewtonMeterPerMeter.
func (a *TorquePerLength) KilonewtonMetersPerMeter() float64 {
	if a.kilonewton_meters_per_meterLazy != nil {
		return *a.kilonewton_meters_per_meterLazy
	}
	kilonewton_meters_per_meter := a.convertFromBase(TorquePerLengthKilonewtonMeterPerMeter)
	a.kilonewton_meters_per_meterLazy = &kilonewton_meters_per_meter
	return kilonewton_meters_per_meter
}

// MeganewtonMeterPerMeter returns the value in MeganewtonMeterPerMeter.
func (a *TorquePerLength) MeganewtonMetersPerMeter() float64 {
	if a.meganewton_meters_per_meterLazy != nil {
		return *a.meganewton_meters_per_meterLazy
	}
	meganewton_meters_per_meter := a.convertFromBase(TorquePerLengthMeganewtonMeterPerMeter)
	a.meganewton_meters_per_meterLazy = &meganewton_meters_per_meter
	return meganewton_meters_per_meter
}

// KilopoundForceInchPerFoot returns the value in KilopoundForceInchPerFoot.
func (a *TorquePerLength) KilopoundForceInchesPerFoot() float64 {
	if a.kilopound_force_inches_per_footLazy != nil {
		return *a.kilopound_force_inches_per_footLazy
	}
	kilopound_force_inches_per_foot := a.convertFromBase(TorquePerLengthKilopoundForceInchPerFoot)
	a.kilopound_force_inches_per_footLazy = &kilopound_force_inches_per_foot
	return kilopound_force_inches_per_foot
}

// MegapoundForceInchPerFoot returns the value in MegapoundForceInchPerFoot.
func (a *TorquePerLength) MegapoundForceInchesPerFoot() float64 {
	if a.megapound_force_inches_per_footLazy != nil {
		return *a.megapound_force_inches_per_footLazy
	}
	megapound_force_inches_per_foot := a.convertFromBase(TorquePerLengthMegapoundForceInchPerFoot)
	a.megapound_force_inches_per_footLazy = &megapound_force_inches_per_foot
	return megapound_force_inches_per_foot
}

// KilopoundForceFootPerFoot returns the value in KilopoundForceFootPerFoot.
func (a *TorquePerLength) KilopoundForceFeetPerFoot() float64 {
	if a.kilopound_force_feet_per_footLazy != nil {
		return *a.kilopound_force_feet_per_footLazy
	}
	kilopound_force_feet_per_foot := a.convertFromBase(TorquePerLengthKilopoundForceFootPerFoot)
	a.kilopound_force_feet_per_footLazy = &kilopound_force_feet_per_foot
	return kilopound_force_feet_per_foot
}

// MegapoundForceFootPerFoot returns the value in MegapoundForceFootPerFoot.
func (a *TorquePerLength) MegapoundForceFeetPerFoot() float64 {
	if a.megapound_force_feet_per_footLazy != nil {
		return *a.megapound_force_feet_per_footLazy
	}
	megapound_force_feet_per_foot := a.convertFromBase(TorquePerLengthMegapoundForceFootPerFoot)
	a.megapound_force_feet_per_footLazy = &megapound_force_feet_per_foot
	return megapound_force_feet_per_foot
}


// ToDto creates an TorquePerLengthDto representation.
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

// ToDtoJSON creates an TorquePerLengthDto representation.
func (a *TorquePerLength) ToDtoJSON(holdInUnit *TorquePerLengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts TorquePerLength to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a TorquePerLength) String() string {
	return a.ToString(TorquePerLengthNewtonMeterPerMeter, 2)
}

// ToString formats the TorquePerLength to string.
// fractionalDigits -1 for not mention
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

// Check if the given TorquePerLength are equals to the current TorquePerLength
func (a *TorquePerLength) Equals(other *TorquePerLength) bool {
	return a.value == other.BaseValue()
}

// Check if the given TorquePerLength are equals to the current TorquePerLength
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

// Add the given TorquePerLength to the current TorquePerLength.
func (a *TorquePerLength) Add(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value + other.BaseValue()}
}

// Subtract the given TorquePerLength to the current TorquePerLength.
func (a *TorquePerLength) Subtract(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value - other.BaseValue()}
}

// Multiply the given TorquePerLength to the current TorquePerLength.
func (a *TorquePerLength) Multiply(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value * other.BaseValue()}
}

// Divide the given TorquePerLength to the current TorquePerLength.
func (a *TorquePerLength) Divide(other *TorquePerLength) *TorquePerLength {
	return &TorquePerLength{value: a.value / other.BaseValue()}
}