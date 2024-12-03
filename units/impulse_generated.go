// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ImpulseUnits enumeration
type ImpulseUnits string

const (
    
        // 
        ImpulseKilogramMeterPerSecond ImpulseUnits = "KilogramMeterPerSecond"
        // 
        ImpulseNewtonSecond ImpulseUnits = "NewtonSecond"
        // 
        ImpulsePoundFootPerSecond ImpulseUnits = "PoundFootPerSecond"
        // 
        ImpulsePoundForceSecond ImpulseUnits = "PoundForceSecond"
        // 
        ImpulseSlugFootPerSecond ImpulseUnits = "SlugFootPerSecond"
        // 
        ImpulseNanonewtonSecond ImpulseUnits = "NanonewtonSecond"
        // 
        ImpulseMicronewtonSecond ImpulseUnits = "MicronewtonSecond"
        // 
        ImpulseMillinewtonSecond ImpulseUnits = "MillinewtonSecond"
        // 
        ImpulseCentinewtonSecond ImpulseUnits = "CentinewtonSecond"
        // 
        ImpulseDecinewtonSecond ImpulseUnits = "DecinewtonSecond"
        // 
        ImpulseDecanewtonSecond ImpulseUnits = "DecanewtonSecond"
        // 
        ImpulseKilonewtonSecond ImpulseUnits = "KilonewtonSecond"
        // 
        ImpulseMeganewtonSecond ImpulseUnits = "MeganewtonSecond"
)

// ImpulseDto represents an Impulse
type ImpulseDto struct {
	Value float64
	Unit  ImpulseUnits
}

// ImpulseDtoFactory struct to group related functions
type ImpulseDtoFactory struct{}

func (udf ImpulseDtoFactory) FromJSON(data []byte) (*ImpulseDto, error) {
	a := ImpulseDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ImpulseDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Impulse struct
type Impulse struct {
	value       float64
    
    kilogram_meters_per_secondLazy *float64 
    newton_secondsLazy *float64 
    pound_feet_per_secondLazy *float64 
    pound_force_secondsLazy *float64 
    slug_feet_per_secondLazy *float64 
    nanonewton_secondsLazy *float64 
    micronewton_secondsLazy *float64 
    millinewton_secondsLazy *float64 
    centinewton_secondsLazy *float64 
    decinewton_secondsLazy *float64 
    decanewton_secondsLazy *float64 
    kilonewton_secondsLazy *float64 
    meganewton_secondsLazy *float64 
}

// ImpulseFactory struct to group related functions
type ImpulseFactory struct{}

func (uf ImpulseFactory) CreateImpulse(value float64, unit ImpulseUnits) (*Impulse, error) {
	return newImpulse(value, unit)
}

func (uf ImpulseFactory) FromDto(dto ImpulseDto) (*Impulse, error) {
	return newImpulse(dto.Value, dto.Unit)
}

func (uf ImpulseFactory) FromDtoJSON(data []byte) (*Impulse, error) {
	unitDto, err := ImpulseDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ImpulseFactory{}.FromDto(*unitDto)
}


// FromKilogramMeterPerSecond creates a new Impulse instance from KilogramMeterPerSecond.
func (uf ImpulseFactory) FromKilogramMetersPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseKilogramMeterPerSecond)
}

// FromNewtonSecond creates a new Impulse instance from NewtonSecond.
func (uf ImpulseFactory) FromNewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseNewtonSecond)
}

// FromPoundFootPerSecond creates a new Impulse instance from PoundFootPerSecond.
func (uf ImpulseFactory) FromPoundFeetPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulsePoundFootPerSecond)
}

// FromPoundForceSecond creates a new Impulse instance from PoundForceSecond.
func (uf ImpulseFactory) FromPoundForceSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulsePoundForceSecond)
}

// FromSlugFootPerSecond creates a new Impulse instance from SlugFootPerSecond.
func (uf ImpulseFactory) FromSlugFeetPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseSlugFootPerSecond)
}

// FromNanonewtonSecond creates a new Impulse instance from NanonewtonSecond.
func (uf ImpulseFactory) FromNanonewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseNanonewtonSecond)
}

// FromMicronewtonSecond creates a new Impulse instance from MicronewtonSecond.
func (uf ImpulseFactory) FromMicronewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMicronewtonSecond)
}

// FromMillinewtonSecond creates a new Impulse instance from MillinewtonSecond.
func (uf ImpulseFactory) FromMillinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMillinewtonSecond)
}

// FromCentinewtonSecond creates a new Impulse instance from CentinewtonSecond.
func (uf ImpulseFactory) FromCentinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseCentinewtonSecond)
}

// FromDecinewtonSecond creates a new Impulse instance from DecinewtonSecond.
func (uf ImpulseFactory) FromDecinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseDecinewtonSecond)
}

// FromDecanewtonSecond creates a new Impulse instance from DecanewtonSecond.
func (uf ImpulseFactory) FromDecanewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseDecanewtonSecond)
}

// FromKilonewtonSecond creates a new Impulse instance from KilonewtonSecond.
func (uf ImpulseFactory) FromKilonewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseKilonewtonSecond)
}

// FromMeganewtonSecond creates a new Impulse instance from MeganewtonSecond.
func (uf ImpulseFactory) FromMeganewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMeganewtonSecond)
}




// newImpulse creates a new Impulse.
func newImpulse(value float64, fromUnit ImpulseUnits) (*Impulse, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Impulse{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Impulse in NewtonSecond.
func (a *Impulse) BaseValue() float64 {
	return a.value
}


// KilogramMeterPerSecond returns the value in KilogramMeterPerSecond.
func (a *Impulse) KilogramMetersPerSecond() float64 {
	if a.kilogram_meters_per_secondLazy != nil {
		return *a.kilogram_meters_per_secondLazy
	}
	kilogram_meters_per_second := a.convertFromBase(ImpulseKilogramMeterPerSecond)
	a.kilogram_meters_per_secondLazy = &kilogram_meters_per_second
	return kilogram_meters_per_second
}

// NewtonSecond returns the value in NewtonSecond.
func (a *Impulse) NewtonSeconds() float64 {
	if a.newton_secondsLazy != nil {
		return *a.newton_secondsLazy
	}
	newton_seconds := a.convertFromBase(ImpulseNewtonSecond)
	a.newton_secondsLazy = &newton_seconds
	return newton_seconds
}

// PoundFootPerSecond returns the value in PoundFootPerSecond.
func (a *Impulse) PoundFeetPerSecond() float64 {
	if a.pound_feet_per_secondLazy != nil {
		return *a.pound_feet_per_secondLazy
	}
	pound_feet_per_second := a.convertFromBase(ImpulsePoundFootPerSecond)
	a.pound_feet_per_secondLazy = &pound_feet_per_second
	return pound_feet_per_second
}

// PoundForceSecond returns the value in PoundForceSecond.
func (a *Impulse) PoundForceSeconds() float64 {
	if a.pound_force_secondsLazy != nil {
		return *a.pound_force_secondsLazy
	}
	pound_force_seconds := a.convertFromBase(ImpulsePoundForceSecond)
	a.pound_force_secondsLazy = &pound_force_seconds
	return pound_force_seconds
}

// SlugFootPerSecond returns the value in SlugFootPerSecond.
func (a *Impulse) SlugFeetPerSecond() float64 {
	if a.slug_feet_per_secondLazy != nil {
		return *a.slug_feet_per_secondLazy
	}
	slug_feet_per_second := a.convertFromBase(ImpulseSlugFootPerSecond)
	a.slug_feet_per_secondLazy = &slug_feet_per_second
	return slug_feet_per_second
}

// NanonewtonSecond returns the value in NanonewtonSecond.
func (a *Impulse) NanonewtonSeconds() float64 {
	if a.nanonewton_secondsLazy != nil {
		return *a.nanonewton_secondsLazy
	}
	nanonewton_seconds := a.convertFromBase(ImpulseNanonewtonSecond)
	a.nanonewton_secondsLazy = &nanonewton_seconds
	return nanonewton_seconds
}

// MicronewtonSecond returns the value in MicronewtonSecond.
func (a *Impulse) MicronewtonSeconds() float64 {
	if a.micronewton_secondsLazy != nil {
		return *a.micronewton_secondsLazy
	}
	micronewton_seconds := a.convertFromBase(ImpulseMicronewtonSecond)
	a.micronewton_secondsLazy = &micronewton_seconds
	return micronewton_seconds
}

// MillinewtonSecond returns the value in MillinewtonSecond.
func (a *Impulse) MillinewtonSeconds() float64 {
	if a.millinewton_secondsLazy != nil {
		return *a.millinewton_secondsLazy
	}
	millinewton_seconds := a.convertFromBase(ImpulseMillinewtonSecond)
	a.millinewton_secondsLazy = &millinewton_seconds
	return millinewton_seconds
}

// CentinewtonSecond returns the value in CentinewtonSecond.
func (a *Impulse) CentinewtonSeconds() float64 {
	if a.centinewton_secondsLazy != nil {
		return *a.centinewton_secondsLazy
	}
	centinewton_seconds := a.convertFromBase(ImpulseCentinewtonSecond)
	a.centinewton_secondsLazy = &centinewton_seconds
	return centinewton_seconds
}

// DecinewtonSecond returns the value in DecinewtonSecond.
func (a *Impulse) DecinewtonSeconds() float64 {
	if a.decinewton_secondsLazy != nil {
		return *a.decinewton_secondsLazy
	}
	decinewton_seconds := a.convertFromBase(ImpulseDecinewtonSecond)
	a.decinewton_secondsLazy = &decinewton_seconds
	return decinewton_seconds
}

// DecanewtonSecond returns the value in DecanewtonSecond.
func (a *Impulse) DecanewtonSeconds() float64 {
	if a.decanewton_secondsLazy != nil {
		return *a.decanewton_secondsLazy
	}
	decanewton_seconds := a.convertFromBase(ImpulseDecanewtonSecond)
	a.decanewton_secondsLazy = &decanewton_seconds
	return decanewton_seconds
}

// KilonewtonSecond returns the value in KilonewtonSecond.
func (a *Impulse) KilonewtonSeconds() float64 {
	if a.kilonewton_secondsLazy != nil {
		return *a.kilonewton_secondsLazy
	}
	kilonewton_seconds := a.convertFromBase(ImpulseKilonewtonSecond)
	a.kilonewton_secondsLazy = &kilonewton_seconds
	return kilonewton_seconds
}

// MeganewtonSecond returns the value in MeganewtonSecond.
func (a *Impulse) MeganewtonSeconds() float64 {
	if a.meganewton_secondsLazy != nil {
		return *a.meganewton_secondsLazy
	}
	meganewton_seconds := a.convertFromBase(ImpulseMeganewtonSecond)
	a.meganewton_secondsLazy = &meganewton_seconds
	return meganewton_seconds
}


// ToDto creates an ImpulseDto representation.
func (a *Impulse) ToDto(holdInUnit *ImpulseUnits) ImpulseDto {
	if holdInUnit == nil {
		defaultUnit := ImpulseNewtonSecond // Default value
		holdInUnit = &defaultUnit
	}

	return ImpulseDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ImpulseDto representation.
func (a *Impulse) ToDtoJSON(holdInUnit *ImpulseUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Impulse to a specific unit value.
func (a *Impulse) Convert(toUnit ImpulseUnits) float64 {
	switch toUnit { 
    case ImpulseKilogramMeterPerSecond:
		return a.KilogramMetersPerSecond()
    case ImpulseNewtonSecond:
		return a.NewtonSeconds()
    case ImpulsePoundFootPerSecond:
		return a.PoundFeetPerSecond()
    case ImpulsePoundForceSecond:
		return a.PoundForceSeconds()
    case ImpulseSlugFootPerSecond:
		return a.SlugFeetPerSecond()
    case ImpulseNanonewtonSecond:
		return a.NanonewtonSeconds()
    case ImpulseMicronewtonSecond:
		return a.MicronewtonSeconds()
    case ImpulseMillinewtonSecond:
		return a.MillinewtonSeconds()
    case ImpulseCentinewtonSecond:
		return a.CentinewtonSeconds()
    case ImpulseDecinewtonSecond:
		return a.DecinewtonSeconds()
    case ImpulseDecanewtonSecond:
		return a.DecanewtonSeconds()
    case ImpulseKilonewtonSecond:
		return a.KilonewtonSeconds()
    case ImpulseMeganewtonSecond:
		return a.MeganewtonSeconds()
	default:
		return 0
	}
}

func (a *Impulse) convertFromBase(toUnit ImpulseUnits) float64 {
    value := a.value
	switch toUnit { 
	case ImpulseKilogramMeterPerSecond:
		return (value) 
	case ImpulseNewtonSecond:
		return (value) 
	case ImpulsePoundFootPerSecond:
		return (value * 7.230657989877) 
	case ImpulsePoundForceSecond:
		return (value * 0.2248089430997) 
	case ImpulseSlugFootPerSecond:
		return (value * 0.224735720691) 
	case ImpulseNanonewtonSecond:
		return ((value) / 1e-09) 
	case ImpulseMicronewtonSecond:
		return ((value) / 1e-06) 
	case ImpulseMillinewtonSecond:
		return ((value) / 0.001) 
	case ImpulseCentinewtonSecond:
		return ((value) / 0.01) 
	case ImpulseDecinewtonSecond:
		return ((value) / 0.1) 
	case ImpulseDecanewtonSecond:
		return ((value) / 10.0) 
	case ImpulseKilonewtonSecond:
		return ((value) / 1000.0) 
	case ImpulseMeganewtonSecond:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Impulse) convertToBase(value float64, fromUnit ImpulseUnits) float64 {
	switch fromUnit { 
	case ImpulseKilogramMeterPerSecond:
		return (value) 
	case ImpulseNewtonSecond:
		return (value) 
	case ImpulsePoundFootPerSecond:
		return (value / 7.230657989877) 
	case ImpulsePoundForceSecond:
		return (value / 0.2248089430997) 
	case ImpulseSlugFootPerSecond:
		return (value / 0.224735720691) 
	case ImpulseNanonewtonSecond:
		return ((value) * 1e-09) 
	case ImpulseMicronewtonSecond:
		return ((value) * 1e-06) 
	case ImpulseMillinewtonSecond:
		return ((value) * 0.001) 
	case ImpulseCentinewtonSecond:
		return ((value) * 0.01) 
	case ImpulseDecinewtonSecond:
		return ((value) * 0.1) 
	case ImpulseDecanewtonSecond:
		return ((value) * 10.0) 
	case ImpulseKilonewtonSecond:
		return ((value) * 1000.0) 
	case ImpulseMeganewtonSecond:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Impulse) String() string {
	return a.ToString(ImpulseNewtonSecond, 2)
}

// ToString formats the Impulse to string.
// fractionalDigits -1 for not mention
func (a *Impulse) ToString(unit ImpulseUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Impulse) getUnitAbbreviation(unit ImpulseUnits) string {
	switch unit { 
	case ImpulseKilogramMeterPerSecond:
		return "kg·m/s" 
	case ImpulseNewtonSecond:
		return "N·s" 
	case ImpulsePoundFootPerSecond:
		return "lb·ft/s" 
	case ImpulsePoundForceSecond:
		return "lbf·s" 
	case ImpulseSlugFootPerSecond:
		return "slug·ft/s" 
	case ImpulseNanonewtonSecond:
		return "nN·s" 
	case ImpulseMicronewtonSecond:
		return "μN·s" 
	case ImpulseMillinewtonSecond:
		return "mN·s" 
	case ImpulseCentinewtonSecond:
		return "cN·s" 
	case ImpulseDecinewtonSecond:
		return "dN·s" 
	case ImpulseDecanewtonSecond:
		return "daN·s" 
	case ImpulseKilonewtonSecond:
		return "kN·s" 
	case ImpulseMeganewtonSecond:
		return "MN·s" 
	default:
		return ""
	}
}

// Check if the given Impulse are equals to the current Impulse
func (a *Impulse) Equals(other *Impulse) bool {
	return a.value == other.BaseValue()
}

// Check if the given Impulse are equals to the current Impulse
func (a *Impulse) CompareTo(other *Impulse) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Impulse to the current Impulse.
func (a *Impulse) Add(other *Impulse) *Impulse {
	return &Impulse{value: a.value + other.BaseValue()}
}

// Subtract the given Impulse to the current Impulse.
func (a *Impulse) Subtract(other *Impulse) *Impulse {
	return &Impulse{value: a.value - other.BaseValue()}
}

// Multiply the given Impulse to the current Impulse.
func (a *Impulse) Multiply(other *Impulse) *Impulse {
	return &Impulse{value: a.value * other.BaseValue()}
}

// Divide the given Impulse to the current Impulse.
func (a *Impulse) Divide(other *Impulse) *Impulse {
	return &Impulse{value: a.value / other.BaseValue()}
}