// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ForceChangeRateUnits defines various units of ForceChangeRate.
type ForceChangeRateUnits string

const (
    
        // 
        ForceChangeRateNewtonPerMinute ForceChangeRateUnits = "NewtonPerMinute"
        // 
        ForceChangeRateNewtonPerSecond ForceChangeRateUnits = "NewtonPerSecond"
        // 
        ForceChangeRatePoundForcePerMinute ForceChangeRateUnits = "PoundForcePerMinute"
        // 
        ForceChangeRatePoundForcePerSecond ForceChangeRateUnits = "PoundForcePerSecond"
        // 
        ForceChangeRateDecanewtonPerMinute ForceChangeRateUnits = "DecanewtonPerMinute"
        // 
        ForceChangeRateKilonewtonPerMinute ForceChangeRateUnits = "KilonewtonPerMinute"
        // 
        ForceChangeRateNanonewtonPerSecond ForceChangeRateUnits = "NanonewtonPerSecond"
        // 
        ForceChangeRateMicronewtonPerSecond ForceChangeRateUnits = "MicronewtonPerSecond"
        // 
        ForceChangeRateMillinewtonPerSecond ForceChangeRateUnits = "MillinewtonPerSecond"
        // 
        ForceChangeRateCentinewtonPerSecond ForceChangeRateUnits = "CentinewtonPerSecond"
        // 
        ForceChangeRateDecinewtonPerSecond ForceChangeRateUnits = "DecinewtonPerSecond"
        // 
        ForceChangeRateDecanewtonPerSecond ForceChangeRateUnits = "DecanewtonPerSecond"
        // 
        ForceChangeRateKilonewtonPerSecond ForceChangeRateUnits = "KilonewtonPerSecond"
        // 
        ForceChangeRateKilopoundForcePerMinute ForceChangeRateUnits = "KilopoundForcePerMinute"
        // 
        ForceChangeRateKilopoundForcePerSecond ForceChangeRateUnits = "KilopoundForcePerSecond"
)

// ForceChangeRateDto represents a ForceChangeRate measurement with a numerical value and its corresponding unit.
type ForceChangeRateDto struct {
    // Value is the numerical representation of the ForceChangeRate.
	Value float64
    // Unit specifies the unit of measurement for the ForceChangeRate, as defined in the ForceChangeRateUnits enumeration.
	Unit  ForceChangeRateUnits
}

// ForceChangeRateDtoFactory groups methods for creating and serializing ForceChangeRateDto objects.
type ForceChangeRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ForceChangeRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ForceChangeRateDtoFactory) FromJSON(data []byte) (*ForceChangeRateDto, error) {
	a := ForceChangeRateDto{}

    // Parse JSON into ForceChangeRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ForceChangeRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ForceChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ForceChangeRate represents a measurement in a of ForceChangeRate.
//
// Force change rate is the ratio of the force change to the time during which the change occurred (value of force changes per unit time).
type ForceChangeRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    newtons_per_minuteLazy *float64 
    newtons_per_secondLazy *float64 
    pounds_force_per_minuteLazy *float64 
    pounds_force_per_secondLazy *float64 
    decanewtons_per_minuteLazy *float64 
    kilonewtons_per_minuteLazy *float64 
    nanonewtons_per_secondLazy *float64 
    micronewtons_per_secondLazy *float64 
    millinewtons_per_secondLazy *float64 
    centinewtons_per_secondLazy *float64 
    decinewtons_per_secondLazy *float64 
    decanewtons_per_secondLazy *float64 
    kilonewtons_per_secondLazy *float64 
    kilopounds_force_per_minuteLazy *float64 
    kilopounds_force_per_secondLazy *float64 
}

// ForceChangeRateFactory groups methods for creating ForceChangeRate instances.
type ForceChangeRateFactory struct{}

// CreateForceChangeRate creates a new ForceChangeRate instance from the given value and unit.
func (uf ForceChangeRateFactory) CreateForceChangeRate(value float64, unit ForceChangeRateUnits) (*ForceChangeRate, error) {
	return newForceChangeRate(value, unit)
}

// FromDto converts a ForceChangeRateDto to a ForceChangeRate instance.
func (uf ForceChangeRateFactory) FromDto(dto ForceChangeRateDto) (*ForceChangeRate, error) {
	return newForceChangeRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ForceChangeRate instance.
func (uf ForceChangeRateFactory) FromDtoJSON(data []byte) (*ForceChangeRate, error) {
	unitDto, err := ForceChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ForceChangeRateDto from JSON: %w", err)
	}
	return ForceChangeRateFactory{}.FromDto(*unitDto)
}


// FromNewtonsPerMinute creates a new ForceChangeRate instance from a value in NewtonsPerMinute.
func (uf ForceChangeRateFactory) FromNewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNewtonPerMinute)
}

// FromNewtonsPerSecond creates a new ForceChangeRate instance from a value in NewtonsPerSecond.
func (uf ForceChangeRateFactory) FromNewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNewtonPerSecond)
}

// FromPoundsForcePerMinute creates a new ForceChangeRate instance from a value in PoundsForcePerMinute.
func (uf ForceChangeRateFactory) FromPoundsForcePerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRatePoundForcePerMinute)
}

// FromPoundsForcePerSecond creates a new ForceChangeRate instance from a value in PoundsForcePerSecond.
func (uf ForceChangeRateFactory) FromPoundsForcePerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRatePoundForcePerSecond)
}

// FromDecanewtonsPerMinute creates a new ForceChangeRate instance from a value in DecanewtonsPerMinute.
func (uf ForceChangeRateFactory) FromDecanewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecanewtonPerMinute)
}

// FromKilonewtonsPerMinute creates a new ForceChangeRate instance from a value in KilonewtonsPerMinute.
func (uf ForceChangeRateFactory) FromKilonewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilonewtonPerMinute)
}

// FromNanonewtonsPerSecond creates a new ForceChangeRate instance from a value in NanonewtonsPerSecond.
func (uf ForceChangeRateFactory) FromNanonewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNanonewtonPerSecond)
}

// FromMicronewtonsPerSecond creates a new ForceChangeRate instance from a value in MicronewtonsPerSecond.
func (uf ForceChangeRateFactory) FromMicronewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateMicronewtonPerSecond)
}

// FromMillinewtonsPerSecond creates a new ForceChangeRate instance from a value in MillinewtonsPerSecond.
func (uf ForceChangeRateFactory) FromMillinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateMillinewtonPerSecond)
}

// FromCentinewtonsPerSecond creates a new ForceChangeRate instance from a value in CentinewtonsPerSecond.
func (uf ForceChangeRateFactory) FromCentinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateCentinewtonPerSecond)
}

// FromDecinewtonsPerSecond creates a new ForceChangeRate instance from a value in DecinewtonsPerSecond.
func (uf ForceChangeRateFactory) FromDecinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecinewtonPerSecond)
}

// FromDecanewtonsPerSecond creates a new ForceChangeRate instance from a value in DecanewtonsPerSecond.
func (uf ForceChangeRateFactory) FromDecanewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecanewtonPerSecond)
}

// FromKilonewtonsPerSecond creates a new ForceChangeRate instance from a value in KilonewtonsPerSecond.
func (uf ForceChangeRateFactory) FromKilonewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilonewtonPerSecond)
}

// FromKilopoundsForcePerMinute creates a new ForceChangeRate instance from a value in KilopoundsForcePerMinute.
func (uf ForceChangeRateFactory) FromKilopoundsForcePerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilopoundForcePerMinute)
}

// FromKilopoundsForcePerSecond creates a new ForceChangeRate instance from a value in KilopoundsForcePerSecond.
func (uf ForceChangeRateFactory) FromKilopoundsForcePerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilopoundForcePerSecond)
}


// newForceChangeRate creates a new ForceChangeRate.
func newForceChangeRate(value float64, fromUnit ForceChangeRateUnits) (*ForceChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ForceChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ForceChangeRate in NewtonPerSecond unit (the base/default quantity).
func (a *ForceChangeRate) BaseValue() float64 {
	return a.value
}


// NewtonsPerMinute returns the ForceChangeRate value in NewtonsPerMinute.
//
// 
func (a *ForceChangeRate) NewtonsPerMinute() float64 {
	if a.newtons_per_minuteLazy != nil {
		return *a.newtons_per_minuteLazy
	}
	newtons_per_minute := a.convertFromBase(ForceChangeRateNewtonPerMinute)
	a.newtons_per_minuteLazy = &newtons_per_minute
	return newtons_per_minute
}

// NewtonsPerSecond returns the ForceChangeRate value in NewtonsPerSecond.
//
// 
func (a *ForceChangeRate) NewtonsPerSecond() float64 {
	if a.newtons_per_secondLazy != nil {
		return *a.newtons_per_secondLazy
	}
	newtons_per_second := a.convertFromBase(ForceChangeRateNewtonPerSecond)
	a.newtons_per_secondLazy = &newtons_per_second
	return newtons_per_second
}

// PoundsForcePerMinute returns the ForceChangeRate value in PoundsForcePerMinute.
//
// 
func (a *ForceChangeRate) PoundsForcePerMinute() float64 {
	if a.pounds_force_per_minuteLazy != nil {
		return *a.pounds_force_per_minuteLazy
	}
	pounds_force_per_minute := a.convertFromBase(ForceChangeRatePoundForcePerMinute)
	a.pounds_force_per_minuteLazy = &pounds_force_per_minute
	return pounds_force_per_minute
}

// PoundsForcePerSecond returns the ForceChangeRate value in PoundsForcePerSecond.
//
// 
func (a *ForceChangeRate) PoundsForcePerSecond() float64 {
	if a.pounds_force_per_secondLazy != nil {
		return *a.pounds_force_per_secondLazy
	}
	pounds_force_per_second := a.convertFromBase(ForceChangeRatePoundForcePerSecond)
	a.pounds_force_per_secondLazy = &pounds_force_per_second
	return pounds_force_per_second
}

// DecanewtonsPerMinute returns the ForceChangeRate value in DecanewtonsPerMinute.
//
// 
func (a *ForceChangeRate) DecanewtonsPerMinute() float64 {
	if a.decanewtons_per_minuteLazy != nil {
		return *a.decanewtons_per_minuteLazy
	}
	decanewtons_per_minute := a.convertFromBase(ForceChangeRateDecanewtonPerMinute)
	a.decanewtons_per_minuteLazy = &decanewtons_per_minute
	return decanewtons_per_minute
}

// KilonewtonsPerMinute returns the ForceChangeRate value in KilonewtonsPerMinute.
//
// 
func (a *ForceChangeRate) KilonewtonsPerMinute() float64 {
	if a.kilonewtons_per_minuteLazy != nil {
		return *a.kilonewtons_per_minuteLazy
	}
	kilonewtons_per_minute := a.convertFromBase(ForceChangeRateKilonewtonPerMinute)
	a.kilonewtons_per_minuteLazy = &kilonewtons_per_minute
	return kilonewtons_per_minute
}

// NanonewtonsPerSecond returns the ForceChangeRate value in NanonewtonsPerSecond.
//
// 
func (a *ForceChangeRate) NanonewtonsPerSecond() float64 {
	if a.nanonewtons_per_secondLazy != nil {
		return *a.nanonewtons_per_secondLazy
	}
	nanonewtons_per_second := a.convertFromBase(ForceChangeRateNanonewtonPerSecond)
	a.nanonewtons_per_secondLazy = &nanonewtons_per_second
	return nanonewtons_per_second
}

// MicronewtonsPerSecond returns the ForceChangeRate value in MicronewtonsPerSecond.
//
// 
func (a *ForceChangeRate) MicronewtonsPerSecond() float64 {
	if a.micronewtons_per_secondLazy != nil {
		return *a.micronewtons_per_secondLazy
	}
	micronewtons_per_second := a.convertFromBase(ForceChangeRateMicronewtonPerSecond)
	a.micronewtons_per_secondLazy = &micronewtons_per_second
	return micronewtons_per_second
}

// MillinewtonsPerSecond returns the ForceChangeRate value in MillinewtonsPerSecond.
//
// 
func (a *ForceChangeRate) MillinewtonsPerSecond() float64 {
	if a.millinewtons_per_secondLazy != nil {
		return *a.millinewtons_per_secondLazy
	}
	millinewtons_per_second := a.convertFromBase(ForceChangeRateMillinewtonPerSecond)
	a.millinewtons_per_secondLazy = &millinewtons_per_second
	return millinewtons_per_second
}

// CentinewtonsPerSecond returns the ForceChangeRate value in CentinewtonsPerSecond.
//
// 
func (a *ForceChangeRate) CentinewtonsPerSecond() float64 {
	if a.centinewtons_per_secondLazy != nil {
		return *a.centinewtons_per_secondLazy
	}
	centinewtons_per_second := a.convertFromBase(ForceChangeRateCentinewtonPerSecond)
	a.centinewtons_per_secondLazy = &centinewtons_per_second
	return centinewtons_per_second
}

// DecinewtonsPerSecond returns the ForceChangeRate value in DecinewtonsPerSecond.
//
// 
func (a *ForceChangeRate) DecinewtonsPerSecond() float64 {
	if a.decinewtons_per_secondLazy != nil {
		return *a.decinewtons_per_secondLazy
	}
	decinewtons_per_second := a.convertFromBase(ForceChangeRateDecinewtonPerSecond)
	a.decinewtons_per_secondLazy = &decinewtons_per_second
	return decinewtons_per_second
}

// DecanewtonsPerSecond returns the ForceChangeRate value in DecanewtonsPerSecond.
//
// 
func (a *ForceChangeRate) DecanewtonsPerSecond() float64 {
	if a.decanewtons_per_secondLazy != nil {
		return *a.decanewtons_per_secondLazy
	}
	decanewtons_per_second := a.convertFromBase(ForceChangeRateDecanewtonPerSecond)
	a.decanewtons_per_secondLazy = &decanewtons_per_second
	return decanewtons_per_second
}

// KilonewtonsPerSecond returns the ForceChangeRate value in KilonewtonsPerSecond.
//
// 
func (a *ForceChangeRate) KilonewtonsPerSecond() float64 {
	if a.kilonewtons_per_secondLazy != nil {
		return *a.kilonewtons_per_secondLazy
	}
	kilonewtons_per_second := a.convertFromBase(ForceChangeRateKilonewtonPerSecond)
	a.kilonewtons_per_secondLazy = &kilonewtons_per_second
	return kilonewtons_per_second
}

// KilopoundsForcePerMinute returns the ForceChangeRate value in KilopoundsForcePerMinute.
//
// 
func (a *ForceChangeRate) KilopoundsForcePerMinute() float64 {
	if a.kilopounds_force_per_minuteLazy != nil {
		return *a.kilopounds_force_per_minuteLazy
	}
	kilopounds_force_per_minute := a.convertFromBase(ForceChangeRateKilopoundForcePerMinute)
	a.kilopounds_force_per_minuteLazy = &kilopounds_force_per_minute
	return kilopounds_force_per_minute
}

// KilopoundsForcePerSecond returns the ForceChangeRate value in KilopoundsForcePerSecond.
//
// 
func (a *ForceChangeRate) KilopoundsForcePerSecond() float64 {
	if a.kilopounds_force_per_secondLazy != nil {
		return *a.kilopounds_force_per_secondLazy
	}
	kilopounds_force_per_second := a.convertFromBase(ForceChangeRateKilopoundForcePerSecond)
	a.kilopounds_force_per_secondLazy = &kilopounds_force_per_second
	return kilopounds_force_per_second
}


// ToDto creates a ForceChangeRateDto representation from the ForceChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerSecond by default.
func (a *ForceChangeRate) ToDto(holdInUnit *ForceChangeRateUnits) ForceChangeRateDto {
	if holdInUnit == nil {
		defaultUnit := ForceChangeRateNewtonPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return ForceChangeRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ForceChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonPerSecond by default.
func (a *ForceChangeRate) ToDtoJSON(holdInUnit *ForceChangeRateUnits) ([]byte, error) {
	// Convert to ForceChangeRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ForceChangeRate to a specific unit value.
// The function uses the provided unit type (ForceChangeRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ForceChangeRate) Convert(toUnit ForceChangeRateUnits) float64 {
	switch toUnit { 
    case ForceChangeRateNewtonPerMinute:
		return a.NewtonsPerMinute()
    case ForceChangeRateNewtonPerSecond:
		return a.NewtonsPerSecond()
    case ForceChangeRatePoundForcePerMinute:
		return a.PoundsForcePerMinute()
    case ForceChangeRatePoundForcePerSecond:
		return a.PoundsForcePerSecond()
    case ForceChangeRateDecanewtonPerMinute:
		return a.DecanewtonsPerMinute()
    case ForceChangeRateKilonewtonPerMinute:
		return a.KilonewtonsPerMinute()
    case ForceChangeRateNanonewtonPerSecond:
		return a.NanonewtonsPerSecond()
    case ForceChangeRateMicronewtonPerSecond:
		return a.MicronewtonsPerSecond()
    case ForceChangeRateMillinewtonPerSecond:
		return a.MillinewtonsPerSecond()
    case ForceChangeRateCentinewtonPerSecond:
		return a.CentinewtonsPerSecond()
    case ForceChangeRateDecinewtonPerSecond:
		return a.DecinewtonsPerSecond()
    case ForceChangeRateDecanewtonPerSecond:
		return a.DecanewtonsPerSecond()
    case ForceChangeRateKilonewtonPerSecond:
		return a.KilonewtonsPerSecond()
    case ForceChangeRateKilopoundForcePerMinute:
		return a.KilopoundsForcePerMinute()
    case ForceChangeRateKilopoundForcePerSecond:
		return a.KilopoundsForcePerSecond()
	default:
		return math.NaN()
	}
}

func (a *ForceChangeRate) convertFromBase(toUnit ForceChangeRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case ForceChangeRateNewtonPerMinute:
		return (value * 60) 
	case ForceChangeRateNewtonPerSecond:
		return (value) 
	case ForceChangeRatePoundForcePerMinute:
		return (value / 4.4482216152605095551842641431421 * 60) 
	case ForceChangeRatePoundForcePerSecond:
		return (value / 4.4482216152605095551842641431421) 
	case ForceChangeRateDecanewtonPerMinute:
		return ((value * 60) / 10.0) 
	case ForceChangeRateKilonewtonPerMinute:
		return ((value * 60) / 1000.0) 
	case ForceChangeRateNanonewtonPerSecond:
		return ((value) / 1e-09) 
	case ForceChangeRateMicronewtonPerSecond:
		return ((value) / 1e-06) 
	case ForceChangeRateMillinewtonPerSecond:
		return ((value) / 0.001) 
	case ForceChangeRateCentinewtonPerSecond:
		return ((value) / 0.01) 
	case ForceChangeRateDecinewtonPerSecond:
		return ((value) / 0.1) 
	case ForceChangeRateDecanewtonPerSecond:
		return ((value) / 10.0) 
	case ForceChangeRateKilonewtonPerSecond:
		return ((value) / 1000.0) 
	case ForceChangeRateKilopoundForcePerMinute:
		return ((value / 4.4482216152605095551842641431421 * 60) / 1000.0) 
	case ForceChangeRateKilopoundForcePerSecond:
		return ((value / 4.4482216152605095551842641431421) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *ForceChangeRate) convertToBase(value float64, fromUnit ForceChangeRateUnits) float64 {
	switch fromUnit { 
	case ForceChangeRateNewtonPerMinute:
		return (value / 60) 
	case ForceChangeRateNewtonPerSecond:
		return (value) 
	case ForceChangeRatePoundForcePerMinute:
		return (value * 4.4482216152605095551842641431421 / 60) 
	case ForceChangeRatePoundForcePerSecond:
		return (value * 4.4482216152605095551842641431421) 
	case ForceChangeRateDecanewtonPerMinute:
		return ((value / 60) * 10.0) 
	case ForceChangeRateKilonewtonPerMinute:
		return ((value / 60) * 1000.0) 
	case ForceChangeRateNanonewtonPerSecond:
		return ((value) * 1e-09) 
	case ForceChangeRateMicronewtonPerSecond:
		return ((value) * 1e-06) 
	case ForceChangeRateMillinewtonPerSecond:
		return ((value) * 0.001) 
	case ForceChangeRateCentinewtonPerSecond:
		return ((value) * 0.01) 
	case ForceChangeRateDecinewtonPerSecond:
		return ((value) * 0.1) 
	case ForceChangeRateDecanewtonPerSecond:
		return ((value) * 10.0) 
	case ForceChangeRateKilonewtonPerSecond:
		return ((value) * 1000.0) 
	case ForceChangeRateKilopoundForcePerMinute:
		return ((value * 4.4482216152605095551842641431421 / 60) * 1000.0) 
	case ForceChangeRateKilopoundForcePerSecond:
		return ((value * 4.4482216152605095551842641431421) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ForceChangeRate in the default unit (NewtonPerSecond),
// formatted to two decimal places.
func (a ForceChangeRate) String() string {
	return a.ToString(ForceChangeRateNewtonPerSecond, 2)
}

// ToString formats the ForceChangeRate value as a string with the specified unit and fractional digits.
// It converts the ForceChangeRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ForceChangeRate value will be converted (e.g., NewtonPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ForceChangeRate with the unit abbreviation.
func (a *ForceChangeRate) ToString(unit ForceChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ForceChangeRate) getUnitAbbreviation(unit ForceChangeRateUnits) string {
	switch unit { 
	case ForceChangeRateNewtonPerMinute:
		return "N/min" 
	case ForceChangeRateNewtonPerSecond:
		return "N/s" 
	case ForceChangeRatePoundForcePerMinute:
		return "lbf/min" 
	case ForceChangeRatePoundForcePerSecond:
		return "lbf/s" 
	case ForceChangeRateDecanewtonPerMinute:
		return "daN/min" 
	case ForceChangeRateKilonewtonPerMinute:
		return "kN/min" 
	case ForceChangeRateNanonewtonPerSecond:
		return "nN/s" 
	case ForceChangeRateMicronewtonPerSecond:
		return "μN/s" 
	case ForceChangeRateMillinewtonPerSecond:
		return "mN/s" 
	case ForceChangeRateCentinewtonPerSecond:
		return "cN/s" 
	case ForceChangeRateDecinewtonPerSecond:
		return "dN/s" 
	case ForceChangeRateDecanewtonPerSecond:
		return "daN/s" 
	case ForceChangeRateKilonewtonPerSecond:
		return "kN/s" 
	case ForceChangeRateKilopoundForcePerMinute:
		return "klbf/min" 
	case ForceChangeRateKilopoundForcePerSecond:
		return "klbf/s" 
	default:
		return ""
	}
}

// Equals checks if the given ForceChangeRate is equal to the current ForceChangeRate.
//
// Parameters:
//    other: The ForceChangeRate to compare against.
//
// Returns:
//    bool: Returns true if both ForceChangeRate are equal, false otherwise.
func (a *ForceChangeRate) Equals(other *ForceChangeRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ForceChangeRate with another ForceChangeRate.
// It returns -1 if the current ForceChangeRate is less than the other ForceChangeRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ForceChangeRate to compare against.
//
// Returns:
//    int: -1 if the current ForceChangeRate is less, 1 if greater, and 0 if equal.
func (a *ForceChangeRate) CompareTo(other *ForceChangeRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ForceChangeRate to the current ForceChangeRate and returns the result.
// The result is a new ForceChangeRate instance with the sum of the values.
//
// Parameters:
//    other: The ForceChangeRate to add to the current ForceChangeRate.
//
// Returns:
//    *ForceChangeRate: A new ForceChangeRate instance representing the sum of both ForceChangeRate.
func (a *ForceChangeRate) Add(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ForceChangeRate from the current ForceChangeRate and returns the result.
// The result is a new ForceChangeRate instance with the difference of the values.
//
// Parameters:
//    other: The ForceChangeRate to subtract from the current ForceChangeRate.
//
// Returns:
//    *ForceChangeRate: A new ForceChangeRate instance representing the difference of both ForceChangeRate.
func (a *ForceChangeRate) Subtract(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ForceChangeRate by the given ForceChangeRate and returns the result.
// The result is a new ForceChangeRate instance with the product of the values.
//
// Parameters:
//    other: The ForceChangeRate to multiply with the current ForceChangeRate.
//
// Returns:
//    *ForceChangeRate: A new ForceChangeRate instance representing the product of both ForceChangeRate.
func (a *ForceChangeRate) Multiply(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value * other.BaseValue()}
}

// Divide divides the current ForceChangeRate by the given ForceChangeRate and returns the result.
// The result is a new ForceChangeRate instance with the quotient of the values.
//
// Parameters:
//    other: The ForceChangeRate to divide the current ForceChangeRate by.
//
// Returns:
//    *ForceChangeRate: A new ForceChangeRate instance representing the quotient of both ForceChangeRate.
func (a *ForceChangeRate) Divide(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value / other.BaseValue()}
}