package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TorqueUnits enumeration
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

// TorqueDto represents an Torque
type TorqueDto struct {
	Value float64
	Unit  TorqueUnits
}

// TorqueDtoFactory struct to group related functions
type TorqueDtoFactory struct{}

func (udf TorqueDtoFactory) FromJSON(data []byte) (*TorqueDto, error) {
	a := TorqueDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TorqueDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Torque struct
type Torque struct {
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

// TorqueFactory struct to group related functions
type TorqueFactory struct{}

func (uf TorqueFactory) CreateTorque(value float64, unit TorqueUnits) (*Torque, error) {
	return newTorque(value, unit)
}

func (uf TorqueFactory) FromDto(dto TorqueDto) (*Torque, error) {
	return newTorque(dto.Value, dto.Unit)
}

func (uf TorqueFactory) FromDtoJSON(data []byte) (*Torque, error) {
	unitDto, err := TorqueDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TorqueFactory{}.FromDto(*unitDto)
}


// FromNewtonMillimeter creates a new Torque instance from NewtonMillimeter.
func (uf TorqueFactory) FromNewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonMillimeter)
}

// FromNewtonCentimeter creates a new Torque instance from NewtonCentimeter.
func (uf TorqueFactory) FromNewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonCentimeter)
}

// FromNewtonMeter creates a new Torque instance from NewtonMeter.
func (uf TorqueFactory) FromNewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueNewtonMeter)
}

// FromPoundalFoot creates a new Torque instance from PoundalFoot.
func (uf TorqueFactory) FromPoundalFeet(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundalFoot)
}

// FromPoundForceInch creates a new Torque instance from PoundForceInch.
func (uf TorqueFactory) FromPoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundForceInch)
}

// FromPoundForceFoot creates a new Torque instance from PoundForceFoot.
func (uf TorqueFactory) FromPoundForceFeet(value float64) (*Torque, error) {
	return newTorque(value, TorquePoundForceFoot)
}

// FromGramForceMillimeter creates a new Torque instance from GramForceMillimeter.
func (uf TorqueFactory) FromGramForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceMillimeter)
}

// FromGramForceCentimeter creates a new Torque instance from GramForceCentimeter.
func (uf TorqueFactory) FromGramForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceCentimeter)
}

// FromGramForceMeter creates a new Torque instance from GramForceMeter.
func (uf TorqueFactory) FromGramForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueGramForceMeter)
}

// FromKilogramForceMillimeter creates a new Torque instance from KilogramForceMillimeter.
func (uf TorqueFactory) FromKilogramForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceMillimeter)
}

// FromKilogramForceCentimeter creates a new Torque instance from KilogramForceCentimeter.
func (uf TorqueFactory) FromKilogramForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceCentimeter)
}

// FromKilogramForceMeter creates a new Torque instance from KilogramForceMeter.
func (uf TorqueFactory) FromKilogramForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilogramForceMeter)
}

// FromTonneForceMillimeter creates a new Torque instance from TonneForceMillimeter.
func (uf TorqueFactory) FromTonneForceMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceMillimeter)
}

// FromTonneForceCentimeter creates a new Torque instance from TonneForceCentimeter.
func (uf TorqueFactory) FromTonneForceCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceCentimeter)
}

// FromTonneForceMeter creates a new Torque instance from TonneForceMeter.
func (uf TorqueFactory) FromTonneForceMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueTonneForceMeter)
}

// FromKilonewtonMillimeter creates a new Torque instance from KilonewtonMillimeter.
func (uf TorqueFactory) FromKilonewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonMillimeter)
}

// FromMeganewtonMillimeter creates a new Torque instance from MeganewtonMillimeter.
func (uf TorqueFactory) FromMeganewtonMillimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonMillimeter)
}

// FromKilonewtonCentimeter creates a new Torque instance from KilonewtonCentimeter.
func (uf TorqueFactory) FromKilonewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonCentimeter)
}

// FromMeganewtonCentimeter creates a new Torque instance from MeganewtonCentimeter.
func (uf TorqueFactory) FromMeganewtonCentimeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonCentimeter)
}

// FromKilonewtonMeter creates a new Torque instance from KilonewtonMeter.
func (uf TorqueFactory) FromKilonewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilonewtonMeter)
}

// FromMeganewtonMeter creates a new Torque instance from MeganewtonMeter.
func (uf TorqueFactory) FromMeganewtonMeters(value float64) (*Torque, error) {
	return newTorque(value, TorqueMeganewtonMeter)
}

// FromKilopoundForceInch creates a new Torque instance from KilopoundForceInch.
func (uf TorqueFactory) FromKilopoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilopoundForceInch)
}

// FromMegapoundForceInch creates a new Torque instance from MegapoundForceInch.
func (uf TorqueFactory) FromMegapoundForceInches(value float64) (*Torque, error) {
	return newTorque(value, TorqueMegapoundForceInch)
}

// FromKilopoundForceFoot creates a new Torque instance from KilopoundForceFoot.
func (uf TorqueFactory) FromKilopoundForceFeet(value float64) (*Torque, error) {
	return newTorque(value, TorqueKilopoundForceFoot)
}

// FromMegapoundForceFoot creates a new Torque instance from MegapoundForceFoot.
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

// BaseValue returns the base value of Torque in NewtonMeter.
func (a *Torque) BaseValue() float64 {
	return a.value
}


// NewtonMillimeter returns the value in NewtonMillimeter.
func (a *Torque) NewtonMillimeters() float64 {
	if a.newton_millimetersLazy != nil {
		return *a.newton_millimetersLazy
	}
	newton_millimeters := a.convertFromBase(TorqueNewtonMillimeter)
	a.newton_millimetersLazy = &newton_millimeters
	return newton_millimeters
}

// NewtonCentimeter returns the value in NewtonCentimeter.
func (a *Torque) NewtonCentimeters() float64 {
	if a.newton_centimetersLazy != nil {
		return *a.newton_centimetersLazy
	}
	newton_centimeters := a.convertFromBase(TorqueNewtonCentimeter)
	a.newton_centimetersLazy = &newton_centimeters
	return newton_centimeters
}

// NewtonMeter returns the value in NewtonMeter.
func (a *Torque) NewtonMeters() float64 {
	if a.newton_metersLazy != nil {
		return *a.newton_metersLazy
	}
	newton_meters := a.convertFromBase(TorqueNewtonMeter)
	a.newton_metersLazy = &newton_meters
	return newton_meters
}

// PoundalFoot returns the value in PoundalFoot.
func (a *Torque) PoundalFeet() float64 {
	if a.poundal_feetLazy != nil {
		return *a.poundal_feetLazy
	}
	poundal_feet := a.convertFromBase(TorquePoundalFoot)
	a.poundal_feetLazy = &poundal_feet
	return poundal_feet
}

// PoundForceInch returns the value in PoundForceInch.
func (a *Torque) PoundForceInches() float64 {
	if a.pound_force_inchesLazy != nil {
		return *a.pound_force_inchesLazy
	}
	pound_force_inches := a.convertFromBase(TorquePoundForceInch)
	a.pound_force_inchesLazy = &pound_force_inches
	return pound_force_inches
}

// PoundForceFoot returns the value in PoundForceFoot.
func (a *Torque) PoundForceFeet() float64 {
	if a.pound_force_feetLazy != nil {
		return *a.pound_force_feetLazy
	}
	pound_force_feet := a.convertFromBase(TorquePoundForceFoot)
	a.pound_force_feetLazy = &pound_force_feet
	return pound_force_feet
}

// GramForceMillimeter returns the value in GramForceMillimeter.
func (a *Torque) GramForceMillimeters() float64 {
	if a.gram_force_millimetersLazy != nil {
		return *a.gram_force_millimetersLazy
	}
	gram_force_millimeters := a.convertFromBase(TorqueGramForceMillimeter)
	a.gram_force_millimetersLazy = &gram_force_millimeters
	return gram_force_millimeters
}

// GramForceCentimeter returns the value in GramForceCentimeter.
func (a *Torque) GramForceCentimeters() float64 {
	if a.gram_force_centimetersLazy != nil {
		return *a.gram_force_centimetersLazy
	}
	gram_force_centimeters := a.convertFromBase(TorqueGramForceCentimeter)
	a.gram_force_centimetersLazy = &gram_force_centimeters
	return gram_force_centimeters
}

// GramForceMeter returns the value in GramForceMeter.
func (a *Torque) GramForceMeters() float64 {
	if a.gram_force_metersLazy != nil {
		return *a.gram_force_metersLazy
	}
	gram_force_meters := a.convertFromBase(TorqueGramForceMeter)
	a.gram_force_metersLazy = &gram_force_meters
	return gram_force_meters
}

// KilogramForceMillimeter returns the value in KilogramForceMillimeter.
func (a *Torque) KilogramForceMillimeters() float64 {
	if a.kilogram_force_millimetersLazy != nil {
		return *a.kilogram_force_millimetersLazy
	}
	kilogram_force_millimeters := a.convertFromBase(TorqueKilogramForceMillimeter)
	a.kilogram_force_millimetersLazy = &kilogram_force_millimeters
	return kilogram_force_millimeters
}

// KilogramForceCentimeter returns the value in KilogramForceCentimeter.
func (a *Torque) KilogramForceCentimeters() float64 {
	if a.kilogram_force_centimetersLazy != nil {
		return *a.kilogram_force_centimetersLazy
	}
	kilogram_force_centimeters := a.convertFromBase(TorqueKilogramForceCentimeter)
	a.kilogram_force_centimetersLazy = &kilogram_force_centimeters
	return kilogram_force_centimeters
}

// KilogramForceMeter returns the value in KilogramForceMeter.
func (a *Torque) KilogramForceMeters() float64 {
	if a.kilogram_force_metersLazy != nil {
		return *a.kilogram_force_metersLazy
	}
	kilogram_force_meters := a.convertFromBase(TorqueKilogramForceMeter)
	a.kilogram_force_metersLazy = &kilogram_force_meters
	return kilogram_force_meters
}

// TonneForceMillimeter returns the value in TonneForceMillimeter.
func (a *Torque) TonneForceMillimeters() float64 {
	if a.tonne_force_millimetersLazy != nil {
		return *a.tonne_force_millimetersLazy
	}
	tonne_force_millimeters := a.convertFromBase(TorqueTonneForceMillimeter)
	a.tonne_force_millimetersLazy = &tonne_force_millimeters
	return tonne_force_millimeters
}

// TonneForceCentimeter returns the value in TonneForceCentimeter.
func (a *Torque) TonneForceCentimeters() float64 {
	if a.tonne_force_centimetersLazy != nil {
		return *a.tonne_force_centimetersLazy
	}
	tonne_force_centimeters := a.convertFromBase(TorqueTonneForceCentimeter)
	a.tonne_force_centimetersLazy = &tonne_force_centimeters
	return tonne_force_centimeters
}

// TonneForceMeter returns the value in TonneForceMeter.
func (a *Torque) TonneForceMeters() float64 {
	if a.tonne_force_metersLazy != nil {
		return *a.tonne_force_metersLazy
	}
	tonne_force_meters := a.convertFromBase(TorqueTonneForceMeter)
	a.tonne_force_metersLazy = &tonne_force_meters
	return tonne_force_meters
}

// KilonewtonMillimeter returns the value in KilonewtonMillimeter.
func (a *Torque) KilonewtonMillimeters() float64 {
	if a.kilonewton_millimetersLazy != nil {
		return *a.kilonewton_millimetersLazy
	}
	kilonewton_millimeters := a.convertFromBase(TorqueKilonewtonMillimeter)
	a.kilonewton_millimetersLazy = &kilonewton_millimeters
	return kilonewton_millimeters
}

// MeganewtonMillimeter returns the value in MeganewtonMillimeter.
func (a *Torque) MeganewtonMillimeters() float64 {
	if a.meganewton_millimetersLazy != nil {
		return *a.meganewton_millimetersLazy
	}
	meganewton_millimeters := a.convertFromBase(TorqueMeganewtonMillimeter)
	a.meganewton_millimetersLazy = &meganewton_millimeters
	return meganewton_millimeters
}

// KilonewtonCentimeter returns the value in KilonewtonCentimeter.
func (a *Torque) KilonewtonCentimeters() float64 {
	if a.kilonewton_centimetersLazy != nil {
		return *a.kilonewton_centimetersLazy
	}
	kilonewton_centimeters := a.convertFromBase(TorqueKilonewtonCentimeter)
	a.kilonewton_centimetersLazy = &kilonewton_centimeters
	return kilonewton_centimeters
}

// MeganewtonCentimeter returns the value in MeganewtonCentimeter.
func (a *Torque) MeganewtonCentimeters() float64 {
	if a.meganewton_centimetersLazy != nil {
		return *a.meganewton_centimetersLazy
	}
	meganewton_centimeters := a.convertFromBase(TorqueMeganewtonCentimeter)
	a.meganewton_centimetersLazy = &meganewton_centimeters
	return meganewton_centimeters
}

// KilonewtonMeter returns the value in KilonewtonMeter.
func (a *Torque) KilonewtonMeters() float64 {
	if a.kilonewton_metersLazy != nil {
		return *a.kilonewton_metersLazy
	}
	kilonewton_meters := a.convertFromBase(TorqueKilonewtonMeter)
	a.kilonewton_metersLazy = &kilonewton_meters
	return kilonewton_meters
}

// MeganewtonMeter returns the value in MeganewtonMeter.
func (a *Torque) MeganewtonMeters() float64 {
	if a.meganewton_metersLazy != nil {
		return *a.meganewton_metersLazy
	}
	meganewton_meters := a.convertFromBase(TorqueMeganewtonMeter)
	a.meganewton_metersLazy = &meganewton_meters
	return meganewton_meters
}

// KilopoundForceInch returns the value in KilopoundForceInch.
func (a *Torque) KilopoundForceInches() float64 {
	if a.kilopound_force_inchesLazy != nil {
		return *a.kilopound_force_inchesLazy
	}
	kilopound_force_inches := a.convertFromBase(TorqueKilopoundForceInch)
	a.kilopound_force_inchesLazy = &kilopound_force_inches
	return kilopound_force_inches
}

// MegapoundForceInch returns the value in MegapoundForceInch.
func (a *Torque) MegapoundForceInches() float64 {
	if a.megapound_force_inchesLazy != nil {
		return *a.megapound_force_inchesLazy
	}
	megapound_force_inches := a.convertFromBase(TorqueMegapoundForceInch)
	a.megapound_force_inchesLazy = &megapound_force_inches
	return megapound_force_inches
}

// KilopoundForceFoot returns the value in KilopoundForceFoot.
func (a *Torque) KilopoundForceFeet() float64 {
	if a.kilopound_force_feetLazy != nil {
		return *a.kilopound_force_feetLazy
	}
	kilopound_force_feet := a.convertFromBase(TorqueKilopoundForceFoot)
	a.kilopound_force_feetLazy = &kilopound_force_feet
	return kilopound_force_feet
}

// MegapoundForceFoot returns the value in MegapoundForceFoot.
func (a *Torque) MegapoundForceFeet() float64 {
	if a.megapound_force_feetLazy != nil {
		return *a.megapound_force_feetLazy
	}
	megapound_force_feet := a.convertFromBase(TorqueMegapoundForceFoot)
	a.megapound_force_feetLazy = &megapound_force_feet
	return megapound_force_feet
}


// ToDto creates an TorqueDto representation.
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

// ToDtoJSON creates an TorqueDto representation.
func (a *Torque) ToDtoJSON(holdInUnit *TorqueUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Torque to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Torque) String() string {
	return a.ToString(TorqueNewtonMeter, 2)
}

// ToString formats the Torque to string.
// fractionalDigits -1 for not mention
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

// Check if the given Torque are equals to the current Torque
func (a *Torque) Equals(other *Torque) bool {
	return a.value == other.BaseValue()
}

// Check if the given Torque are equals to the current Torque
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

// Add the given Torque to the current Torque.
func (a *Torque) Add(other *Torque) *Torque {
	return &Torque{value: a.value + other.BaseValue()}
}

// Subtract the given Torque to the current Torque.
func (a *Torque) Subtract(other *Torque) *Torque {
	return &Torque{value: a.value - other.BaseValue()}
}

// Multiply the given Torque to the current Torque.
func (a *Torque) Multiply(other *Torque) *Torque {
	return &Torque{value: a.value * other.BaseValue()}
}

// Divide the given Torque to the current Torque.
func (a *Torque) Divide(other *Torque) *Torque {
	return &Torque{value: a.value / other.BaseValue()}
}