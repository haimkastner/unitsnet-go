// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ForceUnits defines various units of Force.
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

// ForceDto represents a Force measurement with a numerical value and its corresponding unit.
type ForceDto struct {
    // Value is the numerical representation of the Force.
	Value float64
    // Unit specifies the unit of measurement for the Force, as defined in the ForceUnits enumeration.
	Unit  ForceUnits
}

// ForceDtoFactory groups methods for creating and serializing ForceDto objects.
type ForceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ForceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ForceDtoFactory) FromJSON(data []byte) (*ForceDto, error) {
	a := ForceDto{}

    // Parse JSON into ForceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ForceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ForceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Force represents a measurement in a of Force.
//
// In physics, a force is any influence that causes an object to undergo a certain change, either concerning its movement, direction, or geometrical construction. In other words, a force can cause an object with mass to change its velocity (which includes to begin moving from a state of rest), i.e., to accelerate, or a flexible object to deform, or both. Force can also be described by intuitive concepts such as a push or a pull. A force has both magnitude and direction, making it a vector quantity. It is measured in the SI unit of newtons and represented by the symbol F.
type Force struct {
	// value is the base measurement stored internally.
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

// ForceFactory groups methods for creating Force instances.
type ForceFactory struct{}

// CreateForce creates a new Force instance from the given value and unit.
func (uf ForceFactory) CreateForce(value float64, unit ForceUnits) (*Force, error) {
	return newForce(value, unit)
}

// FromDto converts a ForceDto to a Force instance.
func (uf ForceFactory) FromDto(dto ForceDto) (*Force, error) {
	return newForce(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Force instance.
func (uf ForceFactory) FromDtoJSON(data []byte) (*Force, error) {
	unitDto, err := ForceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ForceDto from JSON: %w", err)
	}
	return ForceFactory{}.FromDto(*unitDto)
}


// FromDyne creates a new Force instance from a value in Dyne.
func (uf ForceFactory) FromDyne(value float64) (*Force, error) {
	return newForce(value, ForceDyn)
}

// FromKilogramsForce creates a new Force instance from a value in KilogramsForce.
func (uf ForceFactory) FromKilogramsForce(value float64) (*Force, error) {
	return newForce(value, ForceKilogramForce)
}

// FromTonnesForce creates a new Force instance from a value in TonnesForce.
func (uf ForceFactory) FromTonnesForce(value float64) (*Force, error) {
	return newForce(value, ForceTonneForce)
}

// FromNewtons creates a new Force instance from a value in Newtons.
func (uf ForceFactory) FromNewtons(value float64) (*Force, error) {
	return newForce(value, ForceNewton)
}

// FromKiloPonds creates a new Force instance from a value in KiloPonds.
func (uf ForceFactory) FromKiloPonds(value float64) (*Force, error) {
	return newForce(value, ForceKiloPond)
}

// FromPoundals creates a new Force instance from a value in Poundals.
func (uf ForceFactory) FromPoundals(value float64) (*Force, error) {
	return newForce(value, ForcePoundal)
}

// FromPoundsForce creates a new Force instance from a value in PoundsForce.
func (uf ForceFactory) FromPoundsForce(value float64) (*Force, error) {
	return newForce(value, ForcePoundForce)
}

// FromOunceForce creates a new Force instance from a value in OunceForce.
func (uf ForceFactory) FromOunceForce(value float64) (*Force, error) {
	return newForce(value, ForceOunceForce)
}

// FromShortTonsForce creates a new Force instance from a value in ShortTonsForce.
func (uf ForceFactory) FromShortTonsForce(value float64) (*Force, error) {
	return newForce(value, ForceShortTonForce)
}

// FromMicronewtons creates a new Force instance from a value in Micronewtons.
func (uf ForceFactory) FromMicronewtons(value float64) (*Force, error) {
	return newForce(value, ForceMicronewton)
}

// FromMillinewtons creates a new Force instance from a value in Millinewtons.
func (uf ForceFactory) FromMillinewtons(value float64) (*Force, error) {
	return newForce(value, ForceMillinewton)
}

// FromDecanewtons creates a new Force instance from a value in Decanewtons.
func (uf ForceFactory) FromDecanewtons(value float64) (*Force, error) {
	return newForce(value, ForceDecanewton)
}

// FromKilonewtons creates a new Force instance from a value in Kilonewtons.
func (uf ForceFactory) FromKilonewtons(value float64) (*Force, error) {
	return newForce(value, ForceKilonewton)
}

// FromMeganewtons creates a new Force instance from a value in Meganewtons.
func (uf ForceFactory) FromMeganewtons(value float64) (*Force, error) {
	return newForce(value, ForceMeganewton)
}

// FromKilopoundsForce creates a new Force instance from a value in KilopoundsForce.
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

// BaseValue returns the base value of Force in Newton unit (the base/default quantity).
func (a *Force) BaseValue() float64 {
	return a.value
}


// Dyne returns the Force value in Dyne.
//
// One dyne is equal to 10 micronewtons, 10e−5 N or to 10 nsn (nanosthenes) in the old metre–tonne–second system of units.
func (a *Force) Dyne() float64 {
	if a.dyneLazy != nil {
		return *a.dyneLazy
	}
	dyne := a.convertFromBase(ForceDyn)
	a.dyneLazy = &dyne
	return dyne
}

// KilogramsForce returns the Force value in KilogramsForce.
//
// The kilogram-force, or kilopond, is equal to the magnitude of the force exerted on one kilogram of mass in a 9.80665 m/s2 gravitational field (standard gravity). Therefore, one kilogram-force is by definition equal to 9.80665 N.
func (a *Force) KilogramsForce() float64 {
	if a.kilograms_forceLazy != nil {
		return *a.kilograms_forceLazy
	}
	kilograms_force := a.convertFromBase(ForceKilogramForce)
	a.kilograms_forceLazy = &kilograms_force
	return kilograms_force
}

// TonnesForce returns the Force value in TonnesForce.
//
// The tonne-force, metric ton-force, megagram-force, and megapond (Mp) are each 1000 kilograms-force.
func (a *Force) TonnesForce() float64 {
	if a.tonnes_forceLazy != nil {
		return *a.tonnes_forceLazy
	}
	tonnes_force := a.convertFromBase(ForceTonneForce)
	a.tonnes_forceLazy = &tonnes_force
	return tonnes_force
}

// Newtons returns the Force value in Newtons.
//
// The newton (symbol: N) is the unit of force in the International System of Units (SI). It is defined as 1 kg⋅m/s2, the force which gives a mass of 1 kilogram an acceleration of 1 metre per second per second.
func (a *Force) Newtons() float64 {
	if a.newtonsLazy != nil {
		return *a.newtonsLazy
	}
	newtons := a.convertFromBase(ForceNewton)
	a.newtonsLazy = &newtons
	return newtons
}

// KiloPonds returns the Force value in KiloPonds.
//
// The kilogram-force, or kilopond, is equal to the magnitude of the force exerted on one kilogram of mass in a 9.80665 m/s2 gravitational field (standard gravity). Therefore, one kilogram-force is by definition equal to 9.80665 N.
func (a *Force) KiloPonds() float64 {
	if a.kilo_pondsLazy != nil {
		return *a.kilo_pondsLazy
	}
	kilo_ponds := a.convertFromBase(ForceKiloPond)
	a.kilo_pondsLazy = &kilo_ponds
	return kilo_ponds
}

// Poundals returns the Force value in Poundals.
//
// The poundal is defined as the force necessary to accelerate 1 pound-mass at 1 foot per second per second. 1 pdl = 0.138254954376 N exactly.
func (a *Force) Poundals() float64 {
	if a.poundalsLazy != nil {
		return *a.poundalsLazy
	}
	poundals := a.convertFromBase(ForcePoundal)
	a.poundalsLazy = &poundals
	return poundals
}

// PoundsForce returns the Force value in PoundsForce.
//
// The standard values of acceleration of the standard gravitational field (gn) and the international avoirdupois pound (lb) result in a pound-force equal to 4.4482216152605 N.
func (a *Force) PoundsForce() float64 {
	if a.pounds_forceLazy != nil {
		return *a.pounds_forceLazy
	}
	pounds_force := a.convertFromBase(ForcePoundForce)
	a.pounds_forceLazy = &pounds_force
	return pounds_force
}

// OunceForce returns the Force value in OunceForce.
//
// An ounce-force is 1⁄16 of a pound-force, or about 0.2780139 newtons.
func (a *Force) OunceForce() float64 {
	if a.ounce_forceLazy != nil {
		return *a.ounce_forceLazy
	}
	ounce_force := a.convertFromBase(ForceOunceForce)
	a.ounce_forceLazy = &ounce_force
	return ounce_force
}

// ShortTonsForce returns the Force value in ShortTonsForce.
//
// The short ton-force is a unit of force equal to 2,000 pounds-force (907.18474 kgf), that is most commonly used in the United States – known there simply as the ton or US ton.
func (a *Force) ShortTonsForce() float64 {
	if a.short_tons_forceLazy != nil {
		return *a.short_tons_forceLazy
	}
	short_tons_force := a.convertFromBase(ForceShortTonForce)
	a.short_tons_forceLazy = &short_tons_force
	return short_tons_force
}

// Micronewtons returns the Force value in Micronewtons.
//
// 
func (a *Force) Micronewtons() float64 {
	if a.micronewtonsLazy != nil {
		return *a.micronewtonsLazy
	}
	micronewtons := a.convertFromBase(ForceMicronewton)
	a.micronewtonsLazy = &micronewtons
	return micronewtons
}

// Millinewtons returns the Force value in Millinewtons.
//
// 
func (a *Force) Millinewtons() float64 {
	if a.millinewtonsLazy != nil {
		return *a.millinewtonsLazy
	}
	millinewtons := a.convertFromBase(ForceMillinewton)
	a.millinewtonsLazy = &millinewtons
	return millinewtons
}

// Decanewtons returns the Force value in Decanewtons.
//
// 
func (a *Force) Decanewtons() float64 {
	if a.decanewtonsLazy != nil {
		return *a.decanewtonsLazy
	}
	decanewtons := a.convertFromBase(ForceDecanewton)
	a.decanewtonsLazy = &decanewtons
	return decanewtons
}

// Kilonewtons returns the Force value in Kilonewtons.
//
// 
func (a *Force) Kilonewtons() float64 {
	if a.kilonewtonsLazy != nil {
		return *a.kilonewtonsLazy
	}
	kilonewtons := a.convertFromBase(ForceKilonewton)
	a.kilonewtonsLazy = &kilonewtons
	return kilonewtons
}

// Meganewtons returns the Force value in Meganewtons.
//
// 
func (a *Force) Meganewtons() float64 {
	if a.meganewtonsLazy != nil {
		return *a.meganewtonsLazy
	}
	meganewtons := a.convertFromBase(ForceMeganewton)
	a.meganewtonsLazy = &meganewtons
	return meganewtons
}

// KilopoundsForce returns the Force value in KilopoundsForce.
//
// 
func (a *Force) KilopoundsForce() float64 {
	if a.kilopounds_forceLazy != nil {
		return *a.kilopounds_forceLazy
	}
	kilopounds_force := a.convertFromBase(ForceKilopoundForce)
	a.kilopounds_forceLazy = &kilopounds_force
	return kilopounds_force
}


// ToDto creates a ForceDto representation from the Force instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Newton by default.
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

// ToDtoJSON creates a JSON representation of the Force instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Newton by default.
func (a *Force) ToDtoJSON(holdInUnit *ForceUnits) ([]byte, error) {
	// Convert to ForceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Force to a specific unit value.
// The function uses the provided unit type (ForceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Force in the default unit (Newton),
// formatted to two decimal places.
func (a Force) String() string {
	return a.ToString(ForceNewton, 2)
}

// ToString formats the Force value as a string with the specified unit and fractional digits.
// It converts the Force to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Force value will be converted (e.g., Newton))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Force with the unit abbreviation.
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

// Equals checks if the given Force is equal to the current Force.
//
// Parameters:
//    other: The Force to compare against.
//
// Returns:
//    bool: Returns true if both Force are equal, false otherwise.
func (a *Force) Equals(other *Force) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Force with another Force.
// It returns -1 if the current Force is less than the other Force, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Force to compare against.
//
// Returns:
//    int: -1 if the current Force is less, 1 if greater, and 0 if equal.
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

// Add adds the given Force to the current Force and returns the result.
// The result is a new Force instance with the sum of the values.
//
// Parameters:
//    other: The Force to add to the current Force.
//
// Returns:
//    *Force: A new Force instance representing the sum of both Force.
func (a *Force) Add(other *Force) *Force {
	return &Force{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Force from the current Force and returns the result.
// The result is a new Force instance with the difference of the values.
//
// Parameters:
//    other: The Force to subtract from the current Force.
//
// Returns:
//    *Force: A new Force instance representing the difference of both Force.
func (a *Force) Subtract(other *Force) *Force {
	return &Force{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Force by the given Force and returns the result.
// The result is a new Force instance with the product of the values.
//
// Parameters:
//    other: The Force to multiply with the current Force.
//
// Returns:
//    *Force: A new Force instance representing the product of both Force.
func (a *Force) Multiply(other *Force) *Force {
	return &Force{value: a.value * other.BaseValue()}
}

// Divide divides the current Force by the given Force and returns the result.
// The result is a new Force instance with the quotient of the values.
//
// Parameters:
//    other: The Force to divide the current Force by.
//
// Returns:
//    *Force: A new Force instance representing the quotient of both Force.
func (a *Force) Divide(other *Force) *Force {
	return &Force{value: a.value / other.BaseValue()}
}