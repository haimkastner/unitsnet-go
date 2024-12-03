package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ForceUnits enumeration
type ForceUnits string

const (
    
        // One dyne is equal to 10 micronewtons, 10e−5 N or to 10 nsn (nanosthenes) in the old metre–tonne–second system of units.
        ForceDyn ForceUnits = "Dyn"
        // The kilogram-force, or kilopond, is equal to the magnitude of the force exerted on one kilogram of mass in a 9.80665 m/s2 gravitational field (standard gravity). Therefore, one kilogram-force is by definition equal to 9.80665 N.
        ForceKilogramForce ForceUnits = "KilogramForce"
        // The tonne-force, metric ton-force, megagram-force, and megapond (Mp) are each 1000 kilograms-force.
        ForceTonneForce ForceUnits = "TonneForce"
        // The newton (symbol: N) is the unit of force in the International System of Units (SI). It is defined as 1 kg⋅m/s2, the force which gives a mass of 1 kilogram an acceleration of 1 metre per second per second.
        ForceNewton ForceUnits = "Newton"
        // The kilogram-force, or kilopond, is equal to the magnitude of the force exerted on one kilogram of mass in a 9.80665 m/s2 gravitational field (standard gravity). Therefore, one kilogram-force is by definition equal to 9.80665 N.
        ForceKiloPond ForceUnits = "KiloPond"
        // The poundal is defined as the force necessary to accelerate 1 pound-mass at 1 foot per second per second. 1 pdl = 0.138254954376 N exactly.
        ForcePoundal ForceUnits = "Poundal"
        // The standard values of acceleration of the standard gravitational field (gn) and the international avoirdupois pound (lb) result in a pound-force equal to 4.4482216152605 N.
        ForcePoundForce ForceUnits = "PoundForce"
        // An ounce-force is 1⁄16 of a pound-force, or about 0.2780139 newtons.
        ForceOunceForce ForceUnits = "OunceForce"
        // The short ton-force is a unit of force equal to 2,000 pounds-force (907.18474 kgf), that is most commonly used in the United States – known there simply as the ton or US ton.
        ForceShortTonForce ForceUnits = "ShortTonForce"
        // 
        ForceMicronewton ForceUnits = "Micronewton"
        // 
        ForceMillinewton ForceUnits = "Millinewton"
        // 
        ForceDecanewton ForceUnits = "Decanewton"
        // 
        ForceKilonewton ForceUnits = "Kilonewton"
        // 
        ForceMeganewton ForceUnits = "Meganewton"
        // 
        ForceKilopoundForce ForceUnits = "KilopoundForce"
)

// ForceDto represents an Force
type ForceDto struct {
	Value float64
	Unit  ForceUnits
}

// ForceDtoFactory struct to group related functions
type ForceDtoFactory struct{}

func (udf ForceDtoFactory) FromJSON(data []byte) (*ForceDto, error) {
	a := ForceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ForceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Force struct
type Force struct {
	value       float64
    
    dyneLazy *float64 
    kilograms_forceLazy *float64 
    tonnes_forceLazy *float64 
    newtonsLazy *float64 
    kilo_pondsLazy *float64 
    poundalsLazy *float64 
    pounds_forceLazy *float64 
    ounce_forceLazy *float64 
    short_tons_forceLazy *float64 
    micronewtonsLazy *float64 
    millinewtonsLazy *float64 
    decanewtonsLazy *float64 
    kilonewtonsLazy *float64 
    meganewtonsLazy *float64 
    kilopounds_forceLazy *float64 
}

// ForceFactory struct to group related functions
type ForceFactory struct{}

func (uf ForceFactory) CreateForce(value float64, unit ForceUnits) (*Force, error) {
	return newForce(value, unit)
}

func (uf ForceFactory) FromDto(dto ForceDto) (*Force, error) {
	return newForce(dto.Value, dto.Unit)
}

func (uf ForceFactory) FromDtoJSON(data []byte) (*Force, error) {
	unitDto, err := ForceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ForceFactory{}.FromDto(*unitDto)
}


// FromDyn creates a new Force instance from Dyn.
func (uf ForceFactory) FromDyne(value float64) (*Force, error) {
	return newForce(value, ForceDyn)
}

// FromKilogramForce creates a new Force instance from KilogramForce.
func (uf ForceFactory) FromKilogramsForce(value float64) (*Force, error) {
	return newForce(value, ForceKilogramForce)
}

// FromTonneForce creates a new Force instance from TonneForce.
func (uf ForceFactory) FromTonnesForce(value float64) (*Force, error) {
	return newForce(value, ForceTonneForce)
}

// FromNewton creates a new Force instance from Newton.
func (uf ForceFactory) FromNewtons(value float64) (*Force, error) {
	return newForce(value, ForceNewton)
}

// FromKiloPond creates a new Force instance from KiloPond.
func (uf ForceFactory) FromKiloPonds(value float64) (*Force, error) {
	return newForce(value, ForceKiloPond)
}

// FromPoundal creates a new Force instance from Poundal.
func (uf ForceFactory) FromPoundals(value float64) (*Force, error) {
	return newForce(value, ForcePoundal)
}

// FromPoundForce creates a new Force instance from PoundForce.
func (uf ForceFactory) FromPoundsForce(value float64) (*Force, error) {
	return newForce(value, ForcePoundForce)
}

// FromOunceForce creates a new Force instance from OunceForce.
func (uf ForceFactory) FromOunceForce(value float64) (*Force, error) {
	return newForce(value, ForceOunceForce)
}

// FromShortTonForce creates a new Force instance from ShortTonForce.
func (uf ForceFactory) FromShortTonsForce(value float64) (*Force, error) {
	return newForce(value, ForceShortTonForce)
}

// FromMicronewton creates a new Force instance from Micronewton.
func (uf ForceFactory) FromMicronewtons(value float64) (*Force, error) {
	return newForce(value, ForceMicronewton)
}

// FromMillinewton creates a new Force instance from Millinewton.
func (uf ForceFactory) FromMillinewtons(value float64) (*Force, error) {
	return newForce(value, ForceMillinewton)
}

// FromDecanewton creates a new Force instance from Decanewton.
func (uf ForceFactory) FromDecanewtons(value float64) (*Force, error) {
	return newForce(value, ForceDecanewton)
}

// FromKilonewton creates a new Force instance from Kilonewton.
func (uf ForceFactory) FromKilonewtons(value float64) (*Force, error) {
	return newForce(value, ForceKilonewton)
}

// FromMeganewton creates a new Force instance from Meganewton.
func (uf ForceFactory) FromMeganewtons(value float64) (*Force, error) {
	return newForce(value, ForceMeganewton)
}

// FromKilopoundForce creates a new Force instance from KilopoundForce.
func (uf ForceFactory) FromKilopoundsForce(value float64) (*Force, error) {
	return newForce(value, ForceKilopoundForce)
}




// newForce creates a new Force.
func newForce(value float64, fromUnit ForceUnits) (*Force, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Force{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Force in Newton.
func (a *Force) BaseValue() float64 {
	return a.value
}


// Dyn returns the value in Dyn.
func (a *Force) Dyne() float64 {
	if a.dyneLazy != nil {
		return *a.dyneLazy
	}
	dyne := a.convertFromBase(ForceDyn)
	a.dyneLazy = &dyne
	return dyne
}

// KilogramForce returns the value in KilogramForce.
func (a *Force) KilogramsForce() float64 {
	if a.kilograms_forceLazy != nil {
		return *a.kilograms_forceLazy
	}
	kilograms_force := a.convertFromBase(ForceKilogramForce)
	a.kilograms_forceLazy = &kilograms_force
	return kilograms_force
}

// TonneForce returns the value in TonneForce.
func (a *Force) TonnesForce() float64 {
	if a.tonnes_forceLazy != nil {
		return *a.tonnes_forceLazy
	}
	tonnes_force := a.convertFromBase(ForceTonneForce)
	a.tonnes_forceLazy = &tonnes_force
	return tonnes_force
}

// Newton returns the value in Newton.
func (a *Force) Newtons() float64 {
	if a.newtonsLazy != nil {
		return *a.newtonsLazy
	}
	newtons := a.convertFromBase(ForceNewton)
	a.newtonsLazy = &newtons
	return newtons
}

// KiloPond returns the value in KiloPond.
func (a *Force) KiloPonds() float64 {
	if a.kilo_pondsLazy != nil {
		return *a.kilo_pondsLazy
	}
	kilo_ponds := a.convertFromBase(ForceKiloPond)
	a.kilo_pondsLazy = &kilo_ponds
	return kilo_ponds
}

// Poundal returns the value in Poundal.
func (a *Force) Poundals() float64 {
	if a.poundalsLazy != nil {
		return *a.poundalsLazy
	}
	poundals := a.convertFromBase(ForcePoundal)
	a.poundalsLazy = &poundals
	return poundals
}

// PoundForce returns the value in PoundForce.
func (a *Force) PoundsForce() float64 {
	if a.pounds_forceLazy != nil {
		return *a.pounds_forceLazy
	}
	pounds_force := a.convertFromBase(ForcePoundForce)
	a.pounds_forceLazy = &pounds_force
	return pounds_force
}

// OunceForce returns the value in OunceForce.
func (a *Force) OunceForce() float64 {
	if a.ounce_forceLazy != nil {
		return *a.ounce_forceLazy
	}
	ounce_force := a.convertFromBase(ForceOunceForce)
	a.ounce_forceLazy = &ounce_force
	return ounce_force
}

// ShortTonForce returns the value in ShortTonForce.
func (a *Force) ShortTonsForce() float64 {
	if a.short_tons_forceLazy != nil {
		return *a.short_tons_forceLazy
	}
	short_tons_force := a.convertFromBase(ForceShortTonForce)
	a.short_tons_forceLazy = &short_tons_force
	return short_tons_force
}

// Micronewton returns the value in Micronewton.
func (a *Force) Micronewtons() float64 {
	if a.micronewtonsLazy != nil {
		return *a.micronewtonsLazy
	}
	micronewtons := a.convertFromBase(ForceMicronewton)
	a.micronewtonsLazy = &micronewtons
	return micronewtons
}

// Millinewton returns the value in Millinewton.
func (a *Force) Millinewtons() float64 {
	if a.millinewtonsLazy != nil {
		return *a.millinewtonsLazy
	}
	millinewtons := a.convertFromBase(ForceMillinewton)
	a.millinewtonsLazy = &millinewtons
	return millinewtons
}

// Decanewton returns the value in Decanewton.
func (a *Force) Decanewtons() float64 {
	if a.decanewtonsLazy != nil {
		return *a.decanewtonsLazy
	}
	decanewtons := a.convertFromBase(ForceDecanewton)
	a.decanewtonsLazy = &decanewtons
	return decanewtons
}

// Kilonewton returns the value in Kilonewton.
func (a *Force) Kilonewtons() float64 {
	if a.kilonewtonsLazy != nil {
		return *a.kilonewtonsLazy
	}
	kilonewtons := a.convertFromBase(ForceKilonewton)
	a.kilonewtonsLazy = &kilonewtons
	return kilonewtons
}

// Meganewton returns the value in Meganewton.
func (a *Force) Meganewtons() float64 {
	if a.meganewtonsLazy != nil {
		return *a.meganewtonsLazy
	}
	meganewtons := a.convertFromBase(ForceMeganewton)
	a.meganewtonsLazy = &meganewtons
	return meganewtons
}

// KilopoundForce returns the value in KilopoundForce.
func (a *Force) KilopoundsForce() float64 {
	if a.kilopounds_forceLazy != nil {
		return *a.kilopounds_forceLazy
	}
	kilopounds_force := a.convertFromBase(ForceKilopoundForce)
	a.kilopounds_forceLazy = &kilopounds_force
	return kilopounds_force
}


// ToDto creates an ForceDto representation.
func (a *Force) ToDto(holdInUnit *ForceUnits) ForceDto {
	if holdInUnit == nil {
		defaultUnit := ForceNewton // Default value
		holdInUnit = &defaultUnit
	}

	return ForceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ForceDto representation.
func (a *Force) ToDtoJSON(holdInUnit *ForceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Force to a specific unit value.
func (a *Force) Convert(toUnit ForceUnits) float64 {
	switch toUnit { 
    case ForceDyn:
		return a.Dyne()
    case ForceKilogramForce:
		return a.KilogramsForce()
    case ForceTonneForce:
		return a.TonnesForce()
    case ForceNewton:
		return a.Newtons()
    case ForceKiloPond:
		return a.KiloPonds()
    case ForcePoundal:
		return a.Poundals()
    case ForcePoundForce:
		return a.PoundsForce()
    case ForceOunceForce:
		return a.OunceForce()
    case ForceShortTonForce:
		return a.ShortTonsForce()
    case ForceMicronewton:
		return a.Micronewtons()
    case ForceMillinewton:
		return a.Millinewtons()
    case ForceDecanewton:
		return a.Decanewtons()
    case ForceKilonewton:
		return a.Kilonewtons()
    case ForceMeganewton:
		return a.Meganewtons()
    case ForceKilopoundForce:
		return a.KilopoundsForce()
	default:
		return 0
	}
}

func (a *Force) convertFromBase(toUnit ForceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ForceDyn:
		return (value * 1e5) 
	case ForceKilogramForce:
		return (value / 9.80665) 
	case ForceTonneForce:
		return (value / (9.80665 * 1000)) 
	case ForceNewton:
		return (value) 
	case ForceKiloPond:
		return (value / 9.80665) 
	case ForcePoundal:
		return (value / 0.138254954376) 
	case ForcePoundForce:
		return (value / 4.4482216152605) 
	case ForceOunceForce:
		return (value / (4.4482216152605 / 16)) 
	case ForceShortTonForce:
		return (value / (4.4482216152605 * 2000)) 
	case ForceMicronewton:
		return ((value) / 1e-06) 
	case ForceMillinewton:
		return ((value) / 0.001) 
	case ForceDecanewton:
		return ((value) / 10.0) 
	case ForceKilonewton:
		return ((value) / 1000.0) 
	case ForceMeganewton:
		return ((value) / 1000000.0) 
	case ForceKilopoundForce:
		return ((value / 4.4482216152605) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Force) convertToBase(value float64, fromUnit ForceUnits) float64 {
	switch fromUnit { 
	case ForceDyn:
		return (value / 1e5) 
	case ForceKilogramForce:
		return (value * 9.80665) 
	case ForceTonneForce:
		return (value * (9.80665 * 1000)) 
	case ForceNewton:
		return (value) 
	case ForceKiloPond:
		return (value * 9.80665) 
	case ForcePoundal:
		return (value * 0.138254954376) 
	case ForcePoundForce:
		return (value * 4.4482216152605) 
	case ForceOunceForce:
		return (value * (4.4482216152605 / 16)) 
	case ForceShortTonForce:
		return (value * (4.4482216152605 * 2000)) 
	case ForceMicronewton:
		return ((value) * 1e-06) 
	case ForceMillinewton:
		return ((value) * 0.001) 
	case ForceDecanewton:
		return ((value) * 10.0) 
	case ForceKilonewton:
		return ((value) * 1000.0) 
	case ForceMeganewton:
		return ((value) * 1000000.0) 
	case ForceKilopoundForce:
		return ((value * 4.4482216152605) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Force) String() string {
	return a.ToString(ForceNewton, 2)
}

// ToString formats the Force to string.
// fractionalDigits -1 for not mention
func (a *Force) ToString(unit ForceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Force) getUnitAbbreviation(unit ForceUnits) string {
	switch unit { 
	case ForceDyn:
		return "dyn" 
	case ForceKilogramForce:
		return "kgf" 
	case ForceTonneForce:
		return "tf" 
	case ForceNewton:
		return "N" 
	case ForceKiloPond:
		return "kp" 
	case ForcePoundal:
		return "pdl" 
	case ForcePoundForce:
		return "lbf" 
	case ForceOunceForce:
		return "ozf" 
	case ForceShortTonForce:
		return "tf (short)" 
	case ForceMicronewton:
		return "μN" 
	case ForceMillinewton:
		return "mN" 
	case ForceDecanewton:
		return "daN" 
	case ForceKilonewton:
		return "kN" 
	case ForceMeganewton:
		return "MN" 
	case ForceKilopoundForce:
		return "klbf" 
	default:
		return ""
	}
}

// Check if the given Force are equals to the current Force
func (a *Force) Equals(other *Force) bool {
	return a.value == other.BaseValue()
}

// Check if the given Force are equals to the current Force
func (a *Force) CompareTo(other *Force) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Force to the current Force.
func (a *Force) Add(other *Force) *Force {
	return &Force{value: a.value + other.BaseValue()}
}

// Subtract the given Force to the current Force.
func (a *Force) Subtract(other *Force) *Force {
	return &Force{value: a.value - other.BaseValue()}
}

// Multiply the given Force to the current Force.
func (a *Force) Multiply(other *Force) *Force {
	return &Force{value: a.value * other.BaseValue()}
}

// Divide the given Force to the current Force.
func (a *Force) Divide(other *Force) *Force {
	return &Force{value: a.value / other.BaseValue()}
}