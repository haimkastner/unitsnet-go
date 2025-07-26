// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ImpulseUnits defines various units of Impulse.
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

var internalImpulseUnitsMap = map[ImpulseUnits]bool{
	
	ImpulseKilogramMeterPerSecond: true,
	ImpulseNewtonSecond: true,
	ImpulsePoundFootPerSecond: true,
	ImpulsePoundForceSecond: true,
	ImpulseSlugFootPerSecond: true,
	ImpulseNanonewtonSecond: true,
	ImpulseMicronewtonSecond: true,
	ImpulseMillinewtonSecond: true,
	ImpulseCentinewtonSecond: true,
	ImpulseDecinewtonSecond: true,
	ImpulseDecanewtonSecond: true,
	ImpulseKilonewtonSecond: true,
	ImpulseMeganewtonSecond: true,
}

// ImpulseDto represents a Impulse measurement with a numerical value and its corresponding unit.
type ImpulseDto struct {
    // Value is the numerical representation of the Impulse.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Impulse, as defined in the ImpulseUnits enumeration.
	Unit  ImpulseUnits `json:"unit" validate:"required,oneof=KilogramMeterPerSecond NewtonSecond PoundFootPerSecond PoundForceSecond SlugFootPerSecond NanonewtonSecond MicronewtonSecond MillinewtonSecond CentinewtonSecond DecinewtonSecond DecanewtonSecond KilonewtonSecond MeganewtonSecond"`
}

// ImpulseDtoFactory groups methods for creating and serializing ImpulseDto objects.
type ImpulseDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ImpulseDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ImpulseDtoFactory) FromJSON(data []byte) (*ImpulseDto, error) {
	a := ImpulseDto{}

    // Parse JSON into ImpulseDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ImpulseDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ImpulseDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Impulse represents a measurement in a of Impulse.
//
// In classical mechanics, impulse is the integral of a force, F, over the time interval, t, for which it acts. Impulse applied to an object produces an equivalent vector change in its linear momentum, also in the resultant direction.
type Impulse struct {
	// value is the base measurement stored internally.
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

// ImpulseFactory groups methods for creating Impulse instances.
type ImpulseFactory struct{}

// CreateImpulse creates a new Impulse instance from the given value and unit.
func (uf ImpulseFactory) CreateImpulse(value float64, unit ImpulseUnits) (*Impulse, error) {
	return newImpulse(value, unit)
}

// FromDto converts a ImpulseDto to a Impulse instance.
func (uf ImpulseFactory) FromDto(dto ImpulseDto) (*Impulse, error) {
	return newImpulse(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Impulse instance.
func (uf ImpulseFactory) FromDtoJSON(data []byte) (*Impulse, error) {
	unitDto, err := ImpulseDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ImpulseDto from JSON: %w", err)
	}
	return ImpulseFactory{}.FromDto(*unitDto)
}


// FromKilogramMetersPerSecond creates a new Impulse instance from a value in KilogramMetersPerSecond.
func (uf ImpulseFactory) FromKilogramMetersPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseKilogramMeterPerSecond)
}

// FromNewtonSeconds creates a new Impulse instance from a value in NewtonSeconds.
func (uf ImpulseFactory) FromNewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseNewtonSecond)
}

// FromPoundFeetPerSecond creates a new Impulse instance from a value in PoundFeetPerSecond.
func (uf ImpulseFactory) FromPoundFeetPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulsePoundFootPerSecond)
}

// FromPoundForceSeconds creates a new Impulse instance from a value in PoundForceSeconds.
func (uf ImpulseFactory) FromPoundForceSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulsePoundForceSecond)
}

// FromSlugFeetPerSecond creates a new Impulse instance from a value in SlugFeetPerSecond.
func (uf ImpulseFactory) FromSlugFeetPerSecond(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseSlugFootPerSecond)
}

// FromNanonewtonSeconds creates a new Impulse instance from a value in NanonewtonSeconds.
func (uf ImpulseFactory) FromNanonewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseNanonewtonSecond)
}

// FromMicronewtonSeconds creates a new Impulse instance from a value in MicronewtonSeconds.
func (uf ImpulseFactory) FromMicronewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMicronewtonSecond)
}

// FromMillinewtonSeconds creates a new Impulse instance from a value in MillinewtonSeconds.
func (uf ImpulseFactory) FromMillinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMillinewtonSecond)
}

// FromCentinewtonSeconds creates a new Impulse instance from a value in CentinewtonSeconds.
func (uf ImpulseFactory) FromCentinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseCentinewtonSecond)
}

// FromDecinewtonSeconds creates a new Impulse instance from a value in DecinewtonSeconds.
func (uf ImpulseFactory) FromDecinewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseDecinewtonSecond)
}

// FromDecanewtonSeconds creates a new Impulse instance from a value in DecanewtonSeconds.
func (uf ImpulseFactory) FromDecanewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseDecanewtonSecond)
}

// FromKilonewtonSeconds creates a new Impulse instance from a value in KilonewtonSeconds.
func (uf ImpulseFactory) FromKilonewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseKilonewtonSecond)
}

// FromMeganewtonSeconds creates a new Impulse instance from a value in MeganewtonSeconds.
func (uf ImpulseFactory) FromMeganewtonSeconds(value float64) (*Impulse, error) {
	return newImpulse(value, ImpulseMeganewtonSecond)
}


// newImpulse creates a new Impulse.
func newImpulse(value float64, fromUnit ImpulseUnits) (*Impulse, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalImpulseUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ImpulseUnits", fromUnit)
	}
	a := &Impulse{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Impulse in NewtonSecond unit (the base/default quantity).
func (a *Impulse) BaseValue() float64 {
	return a.value
}


// KilogramMetersPerSecond returns the Impulse value in KilogramMetersPerSecond.
//
// 
func (a *Impulse) KilogramMetersPerSecond() float64 {
	if a.kilogram_meters_per_secondLazy != nil {
		return *a.kilogram_meters_per_secondLazy
	}
	kilogram_meters_per_second := a.convertFromBase(ImpulseKilogramMeterPerSecond)
	a.kilogram_meters_per_secondLazy = &kilogram_meters_per_second
	return kilogram_meters_per_second
}

// NewtonSeconds returns the Impulse value in NewtonSeconds.
//
// 
func (a *Impulse) NewtonSeconds() float64 {
	if a.newton_secondsLazy != nil {
		return *a.newton_secondsLazy
	}
	newton_seconds := a.convertFromBase(ImpulseNewtonSecond)
	a.newton_secondsLazy = &newton_seconds
	return newton_seconds
}

// PoundFeetPerSecond returns the Impulse value in PoundFeetPerSecond.
//
// 
func (a *Impulse) PoundFeetPerSecond() float64 {
	if a.pound_feet_per_secondLazy != nil {
		return *a.pound_feet_per_secondLazy
	}
	pound_feet_per_second := a.convertFromBase(ImpulsePoundFootPerSecond)
	a.pound_feet_per_secondLazy = &pound_feet_per_second
	return pound_feet_per_second
}

// PoundForceSeconds returns the Impulse value in PoundForceSeconds.
//
// 
func (a *Impulse) PoundForceSeconds() float64 {
	if a.pound_force_secondsLazy != nil {
		return *a.pound_force_secondsLazy
	}
	pound_force_seconds := a.convertFromBase(ImpulsePoundForceSecond)
	a.pound_force_secondsLazy = &pound_force_seconds
	return pound_force_seconds
}

// SlugFeetPerSecond returns the Impulse value in SlugFeetPerSecond.
//
// 
func (a *Impulse) SlugFeetPerSecond() float64 {
	if a.slug_feet_per_secondLazy != nil {
		return *a.slug_feet_per_secondLazy
	}
	slug_feet_per_second := a.convertFromBase(ImpulseSlugFootPerSecond)
	a.slug_feet_per_secondLazy = &slug_feet_per_second
	return slug_feet_per_second
}

// NanonewtonSeconds returns the Impulse value in NanonewtonSeconds.
//
// 
func (a *Impulse) NanonewtonSeconds() float64 {
	if a.nanonewton_secondsLazy != nil {
		return *a.nanonewton_secondsLazy
	}
	nanonewton_seconds := a.convertFromBase(ImpulseNanonewtonSecond)
	a.nanonewton_secondsLazy = &nanonewton_seconds
	return nanonewton_seconds
}

// MicronewtonSeconds returns the Impulse value in MicronewtonSeconds.
//
// 
func (a *Impulse) MicronewtonSeconds() float64 {
	if a.micronewton_secondsLazy != nil {
		return *a.micronewton_secondsLazy
	}
	micronewton_seconds := a.convertFromBase(ImpulseMicronewtonSecond)
	a.micronewton_secondsLazy = &micronewton_seconds
	return micronewton_seconds
}

// MillinewtonSeconds returns the Impulse value in MillinewtonSeconds.
//
// 
func (a *Impulse) MillinewtonSeconds() float64 {
	if a.millinewton_secondsLazy != nil {
		return *a.millinewton_secondsLazy
	}
	millinewton_seconds := a.convertFromBase(ImpulseMillinewtonSecond)
	a.millinewton_secondsLazy = &millinewton_seconds
	return millinewton_seconds
}

// CentinewtonSeconds returns the Impulse value in CentinewtonSeconds.
//
// 
func (a *Impulse) CentinewtonSeconds() float64 {
	if a.centinewton_secondsLazy != nil {
		return *a.centinewton_secondsLazy
	}
	centinewton_seconds := a.convertFromBase(ImpulseCentinewtonSecond)
	a.centinewton_secondsLazy = &centinewton_seconds
	return centinewton_seconds
}

// DecinewtonSeconds returns the Impulse value in DecinewtonSeconds.
//
// 
func (a *Impulse) DecinewtonSeconds() float64 {
	if a.decinewton_secondsLazy != nil {
		return *a.decinewton_secondsLazy
	}
	decinewton_seconds := a.convertFromBase(ImpulseDecinewtonSecond)
	a.decinewton_secondsLazy = &decinewton_seconds
	return decinewton_seconds
}

// DecanewtonSeconds returns the Impulse value in DecanewtonSeconds.
//
// 
func (a *Impulse) DecanewtonSeconds() float64 {
	if a.decanewton_secondsLazy != nil {
		return *a.decanewton_secondsLazy
	}
	decanewton_seconds := a.convertFromBase(ImpulseDecanewtonSecond)
	a.decanewton_secondsLazy = &decanewton_seconds
	return decanewton_seconds
}

// KilonewtonSeconds returns the Impulse value in KilonewtonSeconds.
//
// 
func (a *Impulse) KilonewtonSeconds() float64 {
	if a.kilonewton_secondsLazy != nil {
		return *a.kilonewton_secondsLazy
	}
	kilonewton_seconds := a.convertFromBase(ImpulseKilonewtonSecond)
	a.kilonewton_secondsLazy = &kilonewton_seconds
	return kilonewton_seconds
}

// MeganewtonSeconds returns the Impulse value in MeganewtonSeconds.
//
// 
func (a *Impulse) MeganewtonSeconds() float64 {
	if a.meganewton_secondsLazy != nil {
		return *a.meganewton_secondsLazy
	}
	meganewton_seconds := a.convertFromBase(ImpulseMeganewtonSecond)
	a.meganewton_secondsLazy = &meganewton_seconds
	return meganewton_seconds
}


// ToDto creates a ImpulseDto representation from the Impulse instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonSecond by default.
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

// ToDtoJSON creates a JSON representation of the Impulse instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonSecond by default.
func (a *Impulse) ToDtoJSON(holdInUnit *ImpulseUnits) ([]byte, error) {
	// Convert to ImpulseDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Impulse to a specific unit value.
// The function uses the provided unit type (ImpulseUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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
		return (value / (0.45359237 * 0.3048)) 
	case ImpulsePoundForceSecond:
		return (value / (0.45359237 * 9.80665)) 
	case ImpulseSlugFootPerSecond:
		return (value / (0.45359237 * 9.80665)) 
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
		return (value * (0.45359237 * 0.3048)) 
	case ImpulsePoundForceSecond:
		return (value * 0.45359237 * 9.80665) 
	case ImpulseSlugFootPerSecond:
		return (value * (0.45359237 * 9.80665)) 
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

// String returns a string representation of the Impulse in the default unit (NewtonSecond),
// formatted to two decimal places.
func (a Impulse) String() string {
	return a.ToString(ImpulseNewtonSecond, 2)
}

// ToString formats the Impulse value as a string with the specified unit and fractional digits.
// It converts the Impulse to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Impulse value will be converted (e.g., NewtonSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Impulse with the unit abbreviation.
func (a *Impulse) ToString(unit ImpulseUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetImpulseAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetImpulseAbbreviation(unit))
}

// Equals checks if the given Impulse is equal to the current Impulse.
//
// Parameters:
//    other: The Impulse to compare against.
//
// Returns:
//    bool: Returns true if both Impulse are equal, false otherwise.
func (a *Impulse) Equals(other *Impulse) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Impulse with another Impulse.
// It returns -1 if the current Impulse is less than the other Impulse, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Impulse to compare against.
//
// Returns:
//    int: -1 if the current Impulse is less, 1 if greater, and 0 if equal.
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

// Add adds the given Impulse to the current Impulse and returns the result.
// The result is a new Impulse instance with the sum of the values.
//
// Parameters:
//    other: The Impulse to add to the current Impulse.
//
// Returns:
//    *Impulse: A new Impulse instance representing the sum of both Impulse.
func (a *Impulse) Add(other *Impulse) *Impulse {
	return &Impulse{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Impulse from the current Impulse and returns the result.
// The result is a new Impulse instance with the difference of the values.
//
// Parameters:
//    other: The Impulse to subtract from the current Impulse.
//
// Returns:
//    *Impulse: A new Impulse instance representing the difference of both Impulse.
func (a *Impulse) Subtract(other *Impulse) *Impulse {
	return &Impulse{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Impulse by the given Impulse and returns the result.
// The result is a new Impulse instance with the product of the values.
//
// Parameters:
//    other: The Impulse to multiply with the current Impulse.
//
// Returns:
//    *Impulse: A new Impulse instance representing the product of both Impulse.
func (a *Impulse) Multiply(other *Impulse) *Impulse {
	return &Impulse{value: a.value * other.BaseValue()}
}

// Divide divides the current Impulse by the given Impulse and returns the result.
// The result is a new Impulse instance with the quotient of the values.
//
// Parameters:
//    other: The Impulse to divide the current Impulse by.
//
// Returns:
//    *Impulse: A new Impulse instance representing the quotient of both Impulse.
func (a *Impulse) Divide(other *Impulse) *Impulse {
	return &Impulse{value: a.value / other.BaseValue()}
}

// GetImpulseAbbreviation gets the unit abbreviation.
func GetImpulseAbbreviation(unit ImpulseUnits) string {
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