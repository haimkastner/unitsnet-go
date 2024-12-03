// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ForceChangeRateUnits enumeration
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

// ForceChangeRateDto represents an ForceChangeRate
type ForceChangeRateDto struct {
	Value float64
	Unit  ForceChangeRateUnits
}

// ForceChangeRateDtoFactory struct to group related functions
type ForceChangeRateDtoFactory struct{}

func (udf ForceChangeRateDtoFactory) FromJSON(data []byte) (*ForceChangeRateDto, error) {
	a := ForceChangeRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ForceChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ForceChangeRate struct
type ForceChangeRate struct {
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

// ForceChangeRateFactory struct to group related functions
type ForceChangeRateFactory struct{}

func (uf ForceChangeRateFactory) CreateForceChangeRate(value float64, unit ForceChangeRateUnits) (*ForceChangeRate, error) {
	return newForceChangeRate(value, unit)
}

func (uf ForceChangeRateFactory) FromDto(dto ForceChangeRateDto) (*ForceChangeRate, error) {
	return newForceChangeRate(dto.Value, dto.Unit)
}

func (uf ForceChangeRateFactory) FromDtoJSON(data []byte) (*ForceChangeRate, error) {
	unitDto, err := ForceChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ForceChangeRateFactory{}.FromDto(*unitDto)
}


// FromNewtonPerMinute creates a new ForceChangeRate instance from NewtonPerMinute.
func (uf ForceChangeRateFactory) FromNewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNewtonPerMinute)
}

// FromNewtonPerSecond creates a new ForceChangeRate instance from NewtonPerSecond.
func (uf ForceChangeRateFactory) FromNewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNewtonPerSecond)
}

// FromPoundForcePerMinute creates a new ForceChangeRate instance from PoundForcePerMinute.
func (uf ForceChangeRateFactory) FromPoundsForcePerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRatePoundForcePerMinute)
}

// FromPoundForcePerSecond creates a new ForceChangeRate instance from PoundForcePerSecond.
func (uf ForceChangeRateFactory) FromPoundsForcePerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRatePoundForcePerSecond)
}

// FromDecanewtonPerMinute creates a new ForceChangeRate instance from DecanewtonPerMinute.
func (uf ForceChangeRateFactory) FromDecanewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecanewtonPerMinute)
}

// FromKilonewtonPerMinute creates a new ForceChangeRate instance from KilonewtonPerMinute.
func (uf ForceChangeRateFactory) FromKilonewtonsPerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilonewtonPerMinute)
}

// FromNanonewtonPerSecond creates a new ForceChangeRate instance from NanonewtonPerSecond.
func (uf ForceChangeRateFactory) FromNanonewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateNanonewtonPerSecond)
}

// FromMicronewtonPerSecond creates a new ForceChangeRate instance from MicronewtonPerSecond.
func (uf ForceChangeRateFactory) FromMicronewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateMicronewtonPerSecond)
}

// FromMillinewtonPerSecond creates a new ForceChangeRate instance from MillinewtonPerSecond.
func (uf ForceChangeRateFactory) FromMillinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateMillinewtonPerSecond)
}

// FromCentinewtonPerSecond creates a new ForceChangeRate instance from CentinewtonPerSecond.
func (uf ForceChangeRateFactory) FromCentinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateCentinewtonPerSecond)
}

// FromDecinewtonPerSecond creates a new ForceChangeRate instance from DecinewtonPerSecond.
func (uf ForceChangeRateFactory) FromDecinewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecinewtonPerSecond)
}

// FromDecanewtonPerSecond creates a new ForceChangeRate instance from DecanewtonPerSecond.
func (uf ForceChangeRateFactory) FromDecanewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateDecanewtonPerSecond)
}

// FromKilonewtonPerSecond creates a new ForceChangeRate instance from KilonewtonPerSecond.
func (uf ForceChangeRateFactory) FromKilonewtonsPerSecond(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilonewtonPerSecond)
}

// FromKilopoundForcePerMinute creates a new ForceChangeRate instance from KilopoundForcePerMinute.
func (uf ForceChangeRateFactory) FromKilopoundsForcePerMinute(value float64) (*ForceChangeRate, error) {
	return newForceChangeRate(value, ForceChangeRateKilopoundForcePerMinute)
}

// FromKilopoundForcePerSecond creates a new ForceChangeRate instance from KilopoundForcePerSecond.
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

// BaseValue returns the base value of ForceChangeRate in NewtonPerSecond.
func (a *ForceChangeRate) BaseValue() float64 {
	return a.value
}


// NewtonPerMinute returns the value in NewtonPerMinute.
func (a *ForceChangeRate) NewtonsPerMinute() float64 {
	if a.newtons_per_minuteLazy != nil {
		return *a.newtons_per_minuteLazy
	}
	newtons_per_minute := a.convertFromBase(ForceChangeRateNewtonPerMinute)
	a.newtons_per_minuteLazy = &newtons_per_minute
	return newtons_per_minute
}

// NewtonPerSecond returns the value in NewtonPerSecond.
func (a *ForceChangeRate) NewtonsPerSecond() float64 {
	if a.newtons_per_secondLazy != nil {
		return *a.newtons_per_secondLazy
	}
	newtons_per_second := a.convertFromBase(ForceChangeRateNewtonPerSecond)
	a.newtons_per_secondLazy = &newtons_per_second
	return newtons_per_second
}

// PoundForcePerMinute returns the value in PoundForcePerMinute.
func (a *ForceChangeRate) PoundsForcePerMinute() float64 {
	if a.pounds_force_per_minuteLazy != nil {
		return *a.pounds_force_per_minuteLazy
	}
	pounds_force_per_minute := a.convertFromBase(ForceChangeRatePoundForcePerMinute)
	a.pounds_force_per_minuteLazy = &pounds_force_per_minute
	return pounds_force_per_minute
}

// PoundForcePerSecond returns the value in PoundForcePerSecond.
func (a *ForceChangeRate) PoundsForcePerSecond() float64 {
	if a.pounds_force_per_secondLazy != nil {
		return *a.pounds_force_per_secondLazy
	}
	pounds_force_per_second := a.convertFromBase(ForceChangeRatePoundForcePerSecond)
	a.pounds_force_per_secondLazy = &pounds_force_per_second
	return pounds_force_per_second
}

// DecanewtonPerMinute returns the value in DecanewtonPerMinute.
func (a *ForceChangeRate) DecanewtonsPerMinute() float64 {
	if a.decanewtons_per_minuteLazy != nil {
		return *a.decanewtons_per_minuteLazy
	}
	decanewtons_per_minute := a.convertFromBase(ForceChangeRateDecanewtonPerMinute)
	a.decanewtons_per_minuteLazy = &decanewtons_per_minute
	return decanewtons_per_minute
}

// KilonewtonPerMinute returns the value in KilonewtonPerMinute.
func (a *ForceChangeRate) KilonewtonsPerMinute() float64 {
	if a.kilonewtons_per_minuteLazy != nil {
		return *a.kilonewtons_per_minuteLazy
	}
	kilonewtons_per_minute := a.convertFromBase(ForceChangeRateKilonewtonPerMinute)
	a.kilonewtons_per_minuteLazy = &kilonewtons_per_minute
	return kilonewtons_per_minute
}

// NanonewtonPerSecond returns the value in NanonewtonPerSecond.
func (a *ForceChangeRate) NanonewtonsPerSecond() float64 {
	if a.nanonewtons_per_secondLazy != nil {
		return *a.nanonewtons_per_secondLazy
	}
	nanonewtons_per_second := a.convertFromBase(ForceChangeRateNanonewtonPerSecond)
	a.nanonewtons_per_secondLazy = &nanonewtons_per_second
	return nanonewtons_per_second
}

// MicronewtonPerSecond returns the value in MicronewtonPerSecond.
func (a *ForceChangeRate) MicronewtonsPerSecond() float64 {
	if a.micronewtons_per_secondLazy != nil {
		return *a.micronewtons_per_secondLazy
	}
	micronewtons_per_second := a.convertFromBase(ForceChangeRateMicronewtonPerSecond)
	a.micronewtons_per_secondLazy = &micronewtons_per_second
	return micronewtons_per_second
}

// MillinewtonPerSecond returns the value in MillinewtonPerSecond.
func (a *ForceChangeRate) MillinewtonsPerSecond() float64 {
	if a.millinewtons_per_secondLazy != nil {
		return *a.millinewtons_per_secondLazy
	}
	millinewtons_per_second := a.convertFromBase(ForceChangeRateMillinewtonPerSecond)
	a.millinewtons_per_secondLazy = &millinewtons_per_second
	return millinewtons_per_second
}

// CentinewtonPerSecond returns the value in CentinewtonPerSecond.
func (a *ForceChangeRate) CentinewtonsPerSecond() float64 {
	if a.centinewtons_per_secondLazy != nil {
		return *a.centinewtons_per_secondLazy
	}
	centinewtons_per_second := a.convertFromBase(ForceChangeRateCentinewtonPerSecond)
	a.centinewtons_per_secondLazy = &centinewtons_per_second
	return centinewtons_per_second
}

// DecinewtonPerSecond returns the value in DecinewtonPerSecond.
func (a *ForceChangeRate) DecinewtonsPerSecond() float64 {
	if a.decinewtons_per_secondLazy != nil {
		return *a.decinewtons_per_secondLazy
	}
	decinewtons_per_second := a.convertFromBase(ForceChangeRateDecinewtonPerSecond)
	a.decinewtons_per_secondLazy = &decinewtons_per_second
	return decinewtons_per_second
}

// DecanewtonPerSecond returns the value in DecanewtonPerSecond.
func (a *ForceChangeRate) DecanewtonsPerSecond() float64 {
	if a.decanewtons_per_secondLazy != nil {
		return *a.decanewtons_per_secondLazy
	}
	decanewtons_per_second := a.convertFromBase(ForceChangeRateDecanewtonPerSecond)
	a.decanewtons_per_secondLazy = &decanewtons_per_second
	return decanewtons_per_second
}

// KilonewtonPerSecond returns the value in KilonewtonPerSecond.
func (a *ForceChangeRate) KilonewtonsPerSecond() float64 {
	if a.kilonewtons_per_secondLazy != nil {
		return *a.kilonewtons_per_secondLazy
	}
	kilonewtons_per_second := a.convertFromBase(ForceChangeRateKilonewtonPerSecond)
	a.kilonewtons_per_secondLazy = &kilonewtons_per_second
	return kilonewtons_per_second
}

// KilopoundForcePerMinute returns the value in KilopoundForcePerMinute.
func (a *ForceChangeRate) KilopoundsForcePerMinute() float64 {
	if a.kilopounds_force_per_minuteLazy != nil {
		return *a.kilopounds_force_per_minuteLazy
	}
	kilopounds_force_per_minute := a.convertFromBase(ForceChangeRateKilopoundForcePerMinute)
	a.kilopounds_force_per_minuteLazy = &kilopounds_force_per_minute
	return kilopounds_force_per_minute
}

// KilopoundForcePerSecond returns the value in KilopoundForcePerSecond.
func (a *ForceChangeRate) KilopoundsForcePerSecond() float64 {
	if a.kilopounds_force_per_secondLazy != nil {
		return *a.kilopounds_force_per_secondLazy
	}
	kilopounds_force_per_second := a.convertFromBase(ForceChangeRateKilopoundForcePerSecond)
	a.kilopounds_force_per_secondLazy = &kilopounds_force_per_second
	return kilopounds_force_per_second
}


// ToDto creates an ForceChangeRateDto representation.
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

// ToDtoJSON creates an ForceChangeRateDto representation.
func (a *ForceChangeRate) ToDtoJSON(holdInUnit *ForceChangeRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ForceChangeRate to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ForceChangeRate) String() string {
	return a.ToString(ForceChangeRateNewtonPerSecond, 2)
}

// ToString formats the ForceChangeRate to string.
// fractionalDigits -1 for not mention
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

// Check if the given ForceChangeRate are equals to the current ForceChangeRate
func (a *ForceChangeRate) Equals(other *ForceChangeRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given ForceChangeRate are equals to the current ForceChangeRate
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

// Add the given ForceChangeRate to the current ForceChangeRate.
func (a *ForceChangeRate) Add(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value + other.BaseValue()}
}

// Subtract the given ForceChangeRate to the current ForceChangeRate.
func (a *ForceChangeRate) Subtract(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value - other.BaseValue()}
}

// Multiply the given ForceChangeRate to the current ForceChangeRate.
func (a *ForceChangeRate) Multiply(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value * other.BaseValue()}
}

// Divide the given ForceChangeRate to the current ForceChangeRate.
func (a *ForceChangeRate) Divide(other *ForceChangeRate) *ForceChangeRate {
	return &ForceChangeRate{value: a.value / other.BaseValue()}
}