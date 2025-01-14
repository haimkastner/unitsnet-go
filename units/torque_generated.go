// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// TorqueUnits defines various units of Torque.
type TorqueUnits string

const (
    
        // 
        TorqueNewtonMillimeter TorqueUnits = "NewtonMillimeter"
        // 
        TorqueNewtonCentimeter TorqueUnits = "NewtonCentimeter"
        // 
        TorqueNewtonMeter TorqueUnits = "NewtonMeter"
        // 
        TorquePoundalFoot TorqueUnits = "PoundalFoot"
        // 
        TorquePoundForceInch TorqueUnits = "PoundForceInch"
        // 
        TorquePoundForceFoot TorqueUnits = "PoundForceFoot"
        // 
        TorqueGramForceMillimeter TorqueUnits = "GramForceMillimeter"
        // 
        TorqueGramForceCentimeter TorqueUnits = "GramForceCentimeter"
        // 
        TorqueGramForceMeter TorqueUnits = "GramForceMeter"
        // 
        TorqueKilogramForceMillimeter TorqueUnits = "KilogramForceMillimeter"
        // 
        TorqueKilogramForceCentimeter TorqueUnits = "KilogramForceCentimeter"
        // 
        TorqueKilogramForceMeter TorqueUnits = "KilogramForceMeter"
        // 
        TorqueTonneForceMillimeter TorqueUnits = "TonneForceMillimeter"
        // 
        TorqueTonneForceCentimeter TorqueUnits = "TonneForceCentimeter"
        // 
        TorqueTonneForceMeter TorqueUnits = "TonneForceMeter"
        // 
        TorqueKilonewtonMillimeter TorqueUnits = "KilonewtonMillimeter"
        // 
        TorqueMeganewtonMillimeter TorqueUnits = "MeganewtonMillimeter"
        // 
        TorqueKilonewtonCentimeter TorqueUnits = "KilonewtonCentimeter"
        // 
        TorqueMeganewtonCentimeter TorqueUnits = "MeganewtonCentimeter"
        // 
        TorqueKilonewtonMeter TorqueUnits = "KilonewtonMeter"
        // 
        TorqueMeganewtonMeter TorqueUnits = "MeganewtonMeter"
        // 
        TorqueKilopoundForceInch TorqueUnits = "KilopoundForceInch"
        // 
        TorqueMegapoundForceInch TorqueUnits = "MegapoundForceInch"
        // 
        TorqueKilopoundForceFoot TorqueUnits = "KilopoundForceFoot"
        // 
        TorqueMegapoundForceFoot TorqueUnits = "MegapoundForceFoot"
)

// TorqueDto represents a Torque measurement with a numerical value and its corresponding unit.
type TorqueDto struct {
    // Value is the numerical representation of the Torque.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Torque, as defined in the TorqueUnits enumeration.
	Unit  TorqueUnits `json:"unit"`
}

// TorqueDtoFactory groups methods for creating and serializing TorqueDto objects.
type TorqueDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TorqueDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TorqueDtoFactory) FromJSON(data []byte) (*TorqueDto, error) {
	a := TorqueDto{}

    // Parse JSON into TorqueDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a TorqueDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TorqueDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Torque represents a measurement in a of Torque.
//
// Torque, moment or moment of force (see the terminology below), is the tendency of a force to rotate an object about an axis,[1] fulcrum, or pivot. Just as a force is a push or a pull, a torque can be thought of as a twist to an object. Mathematically, torque is defined as the cross product of the lever-arm distance and force, which tends to produce rotation. Loosely speaking, torque is a measure of the turning force on an object such as a bolt or a flywheel. For example, pushing or pulling the handle of a wrench connected to a nut or bolt produces a torque (turning force) that loosens or tightens the nut or bolt.
type Torque struct {
	// value is the base measurement stored internally.
	value       float64
    
    newton_millimetersLazy *float64 
    newton_centimetersLazy *float64 
    newton_metersLazy *float64 
    poundal_feetLazy *float64 
    pound_force_inchesLazy *float64 
    pound_force_feetLazy *float64 
    gram_force_millimetersLazy *float64 
    gram_force_centimetersLazy *float64 
    gram_force_metersLazy *float64 
    kilogram_force_millimetersLazy *float64 
    kilogram_force_centimetersLazy *float64 
    kilogram_force_metersLazy *float64 
    tonne_force_millimetersLazy *float64 
    tonne_force_centimetersLazy *float64 
    tonne_force_metersLazy *float64 
    kilonewton_millimetersLazy *float64 
    meganewton_millimetersLazy *float64 
    kilonewton_centimetersLazy *float64 
    meganewton_centimetersLazy *float64 
    kilonewton_metersLazy *float64 
    meganewton_metersLazy *float64 
    kilopound_force_inchesLazy *float64 
    megapound_force_inchesLazy *float64 
    kilopound_force_feetLazy *float64 
    megapound_force_feetLazy *float64 
}

// TorqueFactory groups methods for creating Torque instances.
type TorqueFactory struct{}

// CreateTorque creates a new Torque instance from the given value and unit.
func (uf TorqueFactory) CreateTorque(value float64, unit TorqueUnits) (*Torque, error) {
	return newTorque(value, unit)
}

// FromDto converts a TorqueDto to a Torque instance.
func (uf TorqueFactory) FromDto(dto TorqueDto) (*Torque, error) {
	return newTorque(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Torque instance.
func (uf TorqueFactory) FromDtoJSON(data []byte) (*Torque, error) {
	unitDto, err := TorqueDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TorqueDto from JSON: %w", err)
	}
	return TorqueFactory{}.FromDto(*unitDto)
}


// FromNewtonMillimeters creates a new Torque instance from a value in NewtonMillimeters.
func (uf TorqueFactory) FromNewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonMillimeter)
}

// FromNewtonCentimeters creates a new Torque instance from a value in NewtonCentimeters.
func (uf TorqueFactory) FromNewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonCentimeter)
}

// FromNewtonMeters creates a new Torque instance from a value in NewtonMeters.
func (uf TorqueFactory) FromNewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonMeter)
}

// FromPoundalFeet creates a new Torque instance from a value in PoundalFeet.
func (uf TorqueFactory) FromPoundalFeet(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundalFoot)
}

// FromPoundForceInches creates a new Torque instance from a value in PoundForceInches.
func (uf TorqueFactory) FromPoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundForceInch)
}

// FromPoundForceFeet creates a new Torque instance from a value in PoundForceFeet.
func (uf TorqueFactory) FromPoundForceFeet(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundForceFoot)
}

// FromGramForceMillimeters creates a new Torque instance from a value in GramForceMillimeters.
func (uf TorqueFactory) FromGramForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceMillimeter)
}

// FromGramForceCentimeters creates a new Torque instance from a value in GramForceCentimeters.
func (uf TorqueFactory) FromGramForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceCentimeter)
}

// FromGramForceMeters creates a new Torque instance from a value in GramForceMeters.
func (uf TorqueFactory) FromGramForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceMeter)
}

// FromKilogramForceMillimeters creates a new Torque instance from a value in KilogramForceMillimeters.
func (uf TorqueFactory) FromKilogramForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceMillimeter)
}

// FromKilogramForceCentimeters creates a new Torque instance from a value in KilogramForceCentimeters.
func (uf TorqueFactory) FromKilogramForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceCentimeter)
}

// FromKilogramForceMeters creates a new Torque instance from a value in KilogramForceMeters.
func (uf TorqueFactory) FromKilogramForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceMeter)
}

// FromTonneForceMillimeters creates a new Torque instance from a value in TonneForceMillimeters.
func (uf TorqueFactory) FromTonneForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceMillimeter)
}

// FromTonneForceCentimeters creates a new Torque instance from a value in TonneForceCentimeters.
func (uf TorqueFactory) FromTonneForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceCentimeter)
}

// FromTonneForceMeters creates a new Torque instance from a value in TonneForceMeters.
func (uf TorqueFactory) FromTonneForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceMeter)
}

// FromKilonewtonMillimeters creates a new Torque instance from a value in KilonewtonMillimeters.
func (uf TorqueFactory) FromKilonewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonMillimeter)
}

// FromMeganewtonMillimeters creates a new Torque instance from a value in MeganewtonMillimeters.
func (uf TorqueFactory) FromMeganewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonMillimeter)
}

// FromKilonewtonCentimeters creates a new Torque instance from a value in KilonewtonCentimeters.
func (uf TorqueFactory) FromKilonewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonCentimeter)
}

// FromMeganewtonCentimeters creates a new Torque instance from a value in MeganewtonCentimeters.
func (uf TorqueFactory) FromMeganewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonCentimeter)
}

// FromKilonewtonMeters creates a new Torque instance from a value in KilonewtonMeters.
func (uf TorqueFactory) FromKilonewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonMeter)
}

// FromMeganewtonMeters creates a new Torque instance from a value in MeganewtonMeters.
func (uf TorqueFactory) FromMeganewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonMeter)
}

// FromKilopoundForceInches creates a new Torque instance from a value in KilopoundForceInches.
func (uf TorqueFactory) FromKilopoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilopoundForceInch)
}

// FromMegapoundForceInches creates a new Torque instance from a value in MegapoundForceInches.
func (uf TorqueFactory) FromMegapoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorqueMegapoundForceInch)
}

// FromKilopoundForceFeet creates a new Torque instance from a value in KilopoundForceFeet.
func (uf TorqueFactory) FromKilopoundForceFeet(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilopoundForceFoot)
}

// FromMegapoundForceFeet creates a new Torque instance from a value in MegapoundForceFeet.
func (uf TorqueFactory) FromMegapoundForceFeet(value float64) (*Torque, error) {
	return newTorque(value, TorqueMegapoundForceFoot)
}


// newTorque creates a new Torque.
func newTorque(value float64, fromUnit TorqueUnits) (*Torque, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Torque{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Torque in NewtonMeter unit (the base/default quantity).
func (a *Torque) BaseValue() float64 {
	return a.value
}


// NewtonMillimeters returns the Torque value in NewtonMillimeters.
//
// 
func (a *Torque) NewtonMillimeters() float64 {
	if a.newton_millimetersLazy != nil {
		return *a.newton_millimetersLazy
	}
	newton_millimeters := a.convertFromBase(TorqueNewtonMillimeter)
	a.newton_millimetersLazy = &newton_millimeters
	return newton_millimeters
}

// NewtonCentimeters returns the Torque value in NewtonCentimeters.
//
// 
func (a *Torque) NewtonCentimeters() float64 {
	if a.newton_centimetersLazy != nil {
		return *a.newton_centimetersLazy
	}
	newton_centimeters := a.convertFromBase(TorqueNewtonCentimeter)
	a.newton_centimetersLazy = &newton_centimeters
	return newton_centimeters
}

// NewtonMeters returns the Torque value in NewtonMeters.
//
// 
func (a *Torque) NewtonMeters() float64 {
	if a.newton_metersLazy != nil {
		return *a.newton_metersLazy
	}
	newton_meters := a.convertFromBase(TorqueNewtonMeter)
	a.newton_metersLazy = &newton_meters
	return newton_meters
}

// PoundalFeet returns the Torque value in PoundalFeet.
//
// 
func (a *Torque) PoundalFeet() float64 {
	if a.poundal_feetLazy != nil {
		return *a.poundal_feetLazy
	}
	poundal_feet := a.convertFromBase(TorquePoundalFoot)
	a.poundal_feetLazy = &poundal_feet
	return poundal_feet
}

// PoundForceInches returns the Torque value in PoundForceInches.
//
// 
func (a *Torque) PoundForceInches() float64 {
	if a.pound_force_inchesLazy != nil {
		return *a.pound_force_inchesLazy
	}
	pound_force_inches := a.convertFromBase(TorquePoundForceInch)
	a.pound_force_inchesLazy = &pound_force_inches
	return pound_force_inches
}

// PoundForceFeet returns the Torque value in PoundForceFeet.
//
// 
func (a *Torque) PoundForceFeet() float64 {
	if a.pound_force_feetLazy != nil {
		return *a.pound_force_feetLazy
	}
	pound_force_feet := a.convertFromBase(TorquePoundForceFoot)
	a.pound_force_feetLazy = &pound_force_feet
	return pound_force_feet
}

// GramForceMillimeters returns the Torque value in GramForceMillimeters.
//
// 
func (a *Torque) GramForceMillimeters() float64 {
	if a.gram_force_millimetersLazy != nil {
		return *a.gram_force_millimetersLazy
	}
	gram_force_millimeters := a.convertFromBase(TorqueGramForceMillimeter)
	a.gram_force_millimetersLazy = &gram_force_millimeters
	return gram_force_millimeters
}

// GramForceCentimeters returns the Torque value in GramForceCentimeters.
//
// 
func (a *Torque) GramForceCentimeters() float64 {
	if a.gram_force_centimetersLazy != nil {
		return *a.gram_force_centimetersLazy
	}
	gram_force_centimeters := a.convertFromBase(TorqueGramForceCentimeter)
	a.gram_force_centimetersLazy = &gram_force_centimeters
	return gram_force_centimeters
}

// GramForceMeters returns the Torque value in GramForceMeters.
//
// 
func (a *Torque) GramForceMeters() float64 {
	if a.gram_force_metersLazy != nil {
		return *a.gram_force_metersLazy
	}
	gram_force_meters := a.convertFromBase(TorqueGramForceMeter)
	a.gram_force_metersLazy = &gram_force_meters
	return gram_force_meters
}

// KilogramForceMillimeters returns the Torque value in KilogramForceMillimeters.
//
// 
func (a *Torque) KilogramForceMillimeters() float64 {
	if a.kilogram_force_millimetersLazy != nil {
		return *a.kilogram_force_millimetersLazy
	}
	kilogram_force_millimeters := a.convertFromBase(TorqueKilogramForceMillimeter)
	a.kilogram_force_millimetersLazy = &kilogram_force_millimeters
	return kilogram_force_millimeters
}

// KilogramForceCentimeters returns the Torque value in KilogramForceCentimeters.
//
// 
func (a *Torque) KilogramForceCentimeters() float64 {
	if a.kilogram_force_centimetersLazy != nil {
		return *a.kilogram_force_centimetersLazy
	}
	kilogram_force_centimeters := a.convertFromBase(TorqueKilogramForceCentimeter)
	a.kilogram_force_centimetersLazy = &kilogram_force_centimeters
	return kilogram_force_centimeters
}

// KilogramForceMeters returns the Torque value in KilogramForceMeters.
//
// 
func (a *Torque) KilogramForceMeters() float64 {
	if a.kilogram_force_metersLazy != nil {
		return *a.kilogram_force_metersLazy
	}
	kilogram_force_meters := a.convertFromBase(TorqueKilogramForceMeter)
	a.kilogram_force_metersLazy = &kilogram_force_meters
	return kilogram_force_meters
}

// TonneForceMillimeters returns the Torque value in TonneForceMillimeters.
//
// 
func (a *Torque) TonneForceMillimeters() float64 {
	if a.tonne_force_millimetersLazy != nil {
		return *a.tonne_force_millimetersLazy
	}
	tonne_force_millimeters := a.convertFromBase(TorqueTonneForceMillimeter)
	a.tonne_force_millimetersLazy = &tonne_force_millimeters
	return tonne_force_millimeters
}

// TonneForceCentimeters returns the Torque value in TonneForceCentimeters.
//
// 
func (a *Torque) TonneForceCentimeters() float64 {
	if a.tonne_force_centimetersLazy != nil {
		return *a.tonne_force_centimetersLazy
	}
	tonne_force_centimeters := a.convertFromBase(TorqueTonneForceCentimeter)
	a.tonne_force_centimetersLazy = &tonne_force_centimeters
	return tonne_force_centimeters
}

// TonneForceMeters returns the Torque value in TonneForceMeters.
//
// 
func (a *Torque) TonneForceMeters() float64 {
	if a.tonne_force_metersLazy != nil {
		return *a.tonne_force_metersLazy
	}
	tonne_force_meters := a.convertFromBase(TorqueTonneForceMeter)
	a.tonne_force_metersLazy = &tonne_force_meters
	return tonne_force_meters
}

// KilonewtonMillimeters returns the Torque value in KilonewtonMillimeters.
//
// 
func (a *Torque) KilonewtonMillimeters() float64 {
	if a.kilonewton_millimetersLazy != nil {
		return *a.kilonewton_millimetersLazy
	}
	kilonewton_millimeters := a.convertFromBase(TorqueKilonewtonMillimeter)
	a.kilonewton_millimetersLazy = &kilonewton_millimeters
	return kilonewton_millimeters
}

// MeganewtonMillimeters returns the Torque value in MeganewtonMillimeters.
//
// 
func (a *Torque) MeganewtonMillimeters() float64 {
	if a.meganewton_millimetersLazy != nil {
		return *a.meganewton_millimetersLazy
	}
	meganewton_millimeters := a.convertFromBase(TorqueMeganewtonMillimeter)
	a.meganewton_millimetersLazy = &meganewton_millimeters
	return meganewton_millimeters
}

// KilonewtonCentimeters returns the Torque value in KilonewtonCentimeters.
//
// 
func (a *Torque) KilonewtonCentimeters() float64 {
	if a.kilonewton_centimetersLazy != nil {
		return *a.kilonewton_centimetersLazy
	}
	kilonewton_centimeters := a.convertFromBase(TorqueKilonewtonCentimeter)
	a.kilonewton_centimetersLazy = &kilonewton_centimeters
	return kilonewton_centimeters
}

// MeganewtonCentimeters returns the Torque value in MeganewtonCentimeters.
//
// 
func (a *Torque) MeganewtonCentimeters() float64 {
	if a.meganewton_centimetersLazy != nil {
		return *a.meganewton_centimetersLazy
	}
	meganewton_centimeters := a.convertFromBase(TorqueMeganewtonCentimeter)
	a.meganewton_centimetersLazy = &meganewton_centimeters
	return meganewton_centimeters
}

// KilonewtonMeters returns the Torque value in KilonewtonMeters.
//
// 
func (a *Torque) KilonewtonMeters() float64 {
	if a.kilonewton_metersLazy != nil {
		return *a.kilonewton_metersLazy
	}
	kilonewton_meters := a.convertFromBase(TorqueKilonewtonMeter)
	a.kilonewton_metersLazy = &kilonewton_meters
	return kilonewton_meters
}

// MeganewtonMeters returns the Torque value in MeganewtonMeters.
//
// 
func (a *Torque) MeganewtonMeters() float64 {
	if a.meganewton_metersLazy != nil {
		return *a.meganewton_metersLazy
	}
	meganewton_meters := a.convertFromBase(TorqueMeganewtonMeter)
	a.meganewton_metersLazy = &meganewton_meters
	return meganewton_meters
}

// KilopoundForceInches returns the Torque value in KilopoundForceInches.
//
// 
func (a *Torque) KilopoundForceInches() float64 {
	if a.kilopound_force_inchesLazy != nil {
		return *a.kilopound_force_inchesLazy
	}
	kilopound_force_inches := a.convertFromBase(TorqueKilopoundForceInch)
	a.kilopound_force_inchesLazy = &kilopound_force_inches
	return kilopound_force_inches
}

// MegapoundForceInches returns the Torque value in MegapoundForceInches.
//
// 
func (a *Torque) MegapoundForceInches() float64 {
	if a.megapound_force_inchesLazy != nil {
		return *a.megapound_force_inchesLazy
	}
	megapound_force_inches := a.convertFromBase(TorqueMegapoundForceInch)
	a.megapound_force_inchesLazy = &megapound_force_inches
	return megapound_force_inches
}

// KilopoundForceFeet returns the Torque value in KilopoundForceFeet.
//
// 
func (a *Torque) KilopoundForceFeet() float64 {
	if a.kilopound_force_feetLazy != nil {
		return *a.kilopound_force_feetLazy
	}
	kilopound_force_feet := a.convertFromBase(TorqueKilopoundForceFoot)
	a.kilopound_force_feetLazy = &kilopound_force_feet
	return kilopound_force_feet
}

// MegapoundForceFeet returns the Torque value in MegapoundForceFeet.
//
// 
func (a *Torque) MegapoundForceFeet() float64 {
	if a.megapound_force_feetLazy != nil {
		return *a.megapound_force_feetLazy
	}
	megapound_force_feet := a.convertFromBase(TorqueMegapoundForceFoot)
	a.megapound_force_feetLazy = &megapound_force_feet
	return megapound_force_feet
}


// ToDto creates a TorqueDto representation from the Torque instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeter by default.
func (a *Torque) ToDto(holdInUnit *TorqueUnits) TorqueDto {
	if holdInUnit == nil {
		defaultUnit := TorqueNewtonMeter // Default value
		holdInUnit = &defaultUnit
	}

	return TorqueDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Torque instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonMeter by default.
func (a *Torque) ToDtoJSON(holdInUnit *TorqueUnits) ([]byte, error) {
	// Convert to TorqueDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Torque to a specific unit value.
// The function uses the provided unit type (TorqueUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Torque) Convert(toUnit TorqueUnits) float64 {
	switch toUnit { 
    case TorqueNewtonMillimeter:
		return a.NewtonMillimeters()
    case TorqueNewtonCentimeter:
		return a.NewtonCentimeters()
    case TorqueNewtonMeter:
		return a.NewtonMeters()
    case TorquePoundalFoot:
		return a.PoundalFeet()
    case TorquePoundForceInch:
		return a.PoundForceInches()
    case TorquePoundForceFoot:
		return a.PoundForceFeet()
    case TorqueGramForceMillimeter:
		return a.GramForceMillimeters()
    case TorqueGramForceCentimeter:
		return a.GramForceCentimeters()
    case TorqueGramForceMeter:
		return a.GramForceMeters()
    case TorqueKilogramForceMillimeter:
		return a.KilogramForceMillimeters()
    case TorqueKilogramForceCentimeter:
		return a.KilogramForceCentimeters()
    case TorqueKilogramForceMeter:
		return a.KilogramForceMeters()
    case TorqueTonneForceMillimeter:
		return a.TonneForceMillimeters()
    case TorqueTonneForceCentimeter:
		return a.TonneForceCentimeters()
    case TorqueTonneForceMeter:
		return a.TonneForceMeters()
    case TorqueKilonewtonMillimeter:
		return a.KilonewtonMillimeters()
    case TorqueMeganewtonMillimeter:
		return a.MeganewtonMillimeters()
    case TorqueKilonewtonCentimeter:
		return a.KilonewtonCentimeters()
    case TorqueMeganewtonCentimeter:
		return a.MeganewtonCentimeters()
    case TorqueKilonewtonMeter:
		return a.KilonewtonMeters()
    case TorqueMeganewtonMeter:
		return a.MeganewtonMeters()
    case TorqueKilopoundForceInch:
		return a.KilopoundForceInches()
    case TorqueMegapoundForceInch:
		return a.MegapoundForceInches()
    case TorqueKilopoundForceFoot:
		return a.KilopoundForceFeet()
    case TorqueMegapoundForceFoot:
		return a.MegapoundForceFeet()
	default:
		return math.NaN()
	}
}

func (a *Torque) convertFromBase(toUnit TorqueUnits) float64 {
    value := a.value
	switch toUnit { 
	case TorqueNewtonMillimeter:
		return (value * 1000) 
	case TorqueNewtonCentimeter:
		return (value * 100) 
	case TorqueNewtonMeter:
		return (value) 
	case TorquePoundalFoot:
		return (value / 4.21401100938048e-2) 
	case TorquePoundForceInch:
		return (value / 1.129848290276167e-1) 
	case TorquePoundForceFoot:
		return (value / 1.3558179483314) 
	case TorqueGramForceMillimeter:
		return (value / 9.80665e-6) 
	case TorqueGramForceCentimeter:
		return (value / 9.80665e-5) 
	case TorqueGramForceMeter:
		return (value / 9.80665e-3) 
	case TorqueKilogramForceMillimeter:
		return (value / 9.80665e-3) 
	case TorqueKilogramForceCentimeter:
		return (value / 9.80665e-2) 
	case TorqueKilogramForceMeter:
		return (value / 9.80665) 
	case TorqueTonneForceMillimeter:
		return (value / 9.80665) 
	case TorqueTonneForceCentimeter:
		return (value / 9.80665e1) 
	case TorqueTonneForceMeter:
		return (value / 9.80665e3) 
	case TorqueKilonewtonMillimeter:
		return ((value * 1000) / 1000.0) 
	case TorqueMeganewtonMillimeter:
		return ((value * 1000) / 1000000.0) 
	case TorqueKilonewtonCentimeter:
		return ((value * 100) / 1000.0) 
	case TorqueMeganewtonCentimeter:
		return ((value * 100) / 1000000.0) 
	case TorqueKilonewtonMeter:
		return ((value) / 1000.0) 
	case TorqueMeganewtonMeter:
		return ((value) / 1000000.0) 
	case TorqueKilopoundForceInch:
		return ((value / 1.129848290276167e-1) / 1000.0) 
	case TorqueMegapoundForceInch:
		return ((value / 1.129848290276167e-1) / 1000000.0) 
	case TorqueKilopoundForceFoot:
		return ((value / 1.3558179483314) / 1000.0) 
	case TorqueMegapoundForceFoot:
		return ((value / 1.3558179483314) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Torque) convertToBase(value float64, fromUnit TorqueUnits) float64 {
	switch fromUnit { 
	case TorqueNewtonMillimeter:
		return (value * 0.001) 
	case TorqueNewtonCentimeter:
		return (value * 0.01) 
	case TorqueNewtonMeter:
		return (value) 
	case TorquePoundalFoot:
		return (value * 4.21401100938048e-2) 
	case TorquePoundForceInch:
		return (value * 1.129848290276167e-1) 
	case TorquePoundForceFoot:
		return (value * 1.3558179483314) 
	case TorqueGramForceMillimeter:
		return (value * 9.80665e-6) 
	case TorqueGramForceCentimeter:
		return (value * 9.80665e-5) 
	case TorqueGramForceMeter:
		return (value * 9.80665e-3) 
	case TorqueKilogramForceMillimeter:
		return (value * 9.80665e-3) 
	case TorqueKilogramForceCentimeter:
		return (value * 9.80665e-2) 
	case TorqueKilogramForceMeter:
		return (value * 9.80665) 
	case TorqueTonneForceMillimeter:
		return (value * 9.80665) 
	case TorqueTonneForceCentimeter:
		return (value * 9.80665e1) 
	case TorqueTonneForceMeter:
		return (value * 9.80665e3) 
	case TorqueKilonewtonMillimeter:
		return ((value * 0.001) * 1000.0) 
	case TorqueMeganewtonMillimeter:
		return ((value * 0.001) * 1000000.0) 
	case TorqueKilonewtonCentimeter:
		return ((value * 0.01) * 1000.0) 
	case TorqueMeganewtonCentimeter:
		return ((value * 0.01) * 1000000.0) 
	case TorqueKilonewtonMeter:
		return ((value) * 1000.0) 
	case TorqueMeganewtonMeter:
		return ((value) * 1000000.0) 
	case TorqueKilopoundForceInch:
		return ((value * 1.129848290276167e-1) * 1000.0) 
	case TorqueMegapoundForceInch:
		return ((value * 1.129848290276167e-1) * 1000000.0) 
	case TorqueKilopoundForceFoot:
		return ((value * 1.3558179483314) * 1000.0) 
	case TorqueMegapoundForceFoot:
		return ((value * 1.3558179483314) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Torque in the default unit (NewtonMeter),
// formatted to two decimal places.
func (a Torque) String() string {
	return a.ToString(TorqueNewtonMeter, 2)
}

// ToString formats the Torque value as a string with the specified unit and fractional digits.
// It converts the Torque to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Torque value will be converted (e.g., NewtonMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Torque with the unit abbreviation.
func (a *Torque) ToString(unit TorqueUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Torque) getUnitAbbreviation(unit TorqueUnits) string {
	switch unit { 
	case TorqueNewtonMillimeter:
		return "N·mm" 
	case TorqueNewtonCentimeter:
		return "N·cm" 
	case TorqueNewtonMeter:
		return "N·m" 
	case TorquePoundalFoot:
		return "pdl·ft" 
	case TorquePoundForceInch:
		return "lbf·in" 
	case TorquePoundForceFoot:
		return "lbf·ft" 
	case TorqueGramForceMillimeter:
		return "gf·mm" 
	case TorqueGramForceCentimeter:
		return "gf·cm" 
	case TorqueGramForceMeter:
		return "gf·m" 
	case TorqueKilogramForceMillimeter:
		return "kgf·mm" 
	case TorqueKilogramForceCentimeter:
		return "kgf·cm" 
	case TorqueKilogramForceMeter:
		return "kgf·m" 
	case TorqueTonneForceMillimeter:
		return "tf·mm" 
	case TorqueTonneForceCentimeter:
		return "tf·cm" 
	case TorqueTonneForceMeter:
		return "tf·m" 
	case TorqueKilonewtonMillimeter:
		return "kN·mm" 
	case TorqueMeganewtonMillimeter:
		return "MN·mm" 
	case TorqueKilonewtonCentimeter:
		return "kN·cm" 
	case TorqueMeganewtonCentimeter:
		return "MN·cm" 
	case TorqueKilonewtonMeter:
		return "kN·m" 
	case TorqueMeganewtonMeter:
		return "MN·m" 
	case TorqueKilopoundForceInch:
		return "klbf·in" 
	case TorqueMegapoundForceInch:
		return "Mlbf·in" 
	case TorqueKilopoundForceFoot:
		return "klbf·ft" 
	case TorqueMegapoundForceFoot:
		return "Mlbf·ft" 
	default:
		return ""
	}
}

// Equals checks if the given Torque is equal to the current Torque.
//
// Parameters:
//    other: The Torque to compare against.
//
// Returns:
//    bool: Returns true if both Torque are equal, false otherwise.
func (a *Torque) Equals(other *Torque) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Torque with another Torque.
// It returns -1 if the current Torque is less than the other Torque, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Torque to compare against.
//
// Returns:
//    int: -1 if the current Torque is less, 1 if greater, and 0 if equal.
func (a *Torque) CompareTo(other *Torque) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Torque to the current Torque and returns the result.
// The result is a new Torque instance with the sum of the values.
//
// Parameters:
//    other: The Torque to add to the current Torque.
//
// Returns:
//    *Torque: A new Torque instance representing the sum of both Torque.
func (a *Torque) Add(other *Torque) *Torque {
	return &Torque{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Torque from the current Torque and returns the result.
// The result is a new Torque instance with the difference of the values.
//
// Parameters:
//    other: The Torque to subtract from the current Torque.
//
// Returns:
//    *Torque: A new Torque instance representing the difference of both Torque.
func (a *Torque) Subtract(other *Torque) *Torque {
	return &Torque{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Torque by the given Torque and returns the result.
// The result is a new Torque instance with the product of the values.
//
// Parameters:
//    other: The Torque to multiply with the current Torque.
//
// Returns:
//    *Torque: A new Torque instance representing the product of both Torque.
func (a *Torque) Multiply(other *Torque) *Torque {
	return &Torque{value: a.value * other.BaseValue()}
}

// Divide divides the current Torque by the given Torque and returns the result.
// The result is a new Torque instance with the quotient of the values.
//
// Parameters:
//    other: The Torque to divide the current Torque by.
//
// Returns:
//    *Torque: A new Torque instance representing the quotient of both Torque.
func (a *Torque) Divide(other *Torque) *Torque {
	return &Torque{value: a.value / other.BaseValue()}
}