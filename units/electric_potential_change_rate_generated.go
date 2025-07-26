// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialChangeRateUnits defines various units of ElectricPotentialChangeRate.
type ElectricPotentialChangeRateUnits string

const (
    
        // 
        ElectricPotentialChangeRateVoltPerSecond ElectricPotentialChangeRateUnits = "VoltPerSecond"
        // 
        ElectricPotentialChangeRateVoltPerMicrosecond ElectricPotentialChangeRateUnits = "VoltPerMicrosecond"
        // 
        ElectricPotentialChangeRateVoltPerMinute ElectricPotentialChangeRateUnits = "VoltPerMinute"
        // 
        ElectricPotentialChangeRateVoltPerHour ElectricPotentialChangeRateUnits = "VoltPerHour"
        // 
        ElectricPotentialChangeRateMicrovoltPerSecond ElectricPotentialChangeRateUnits = "MicrovoltPerSecond"
        // 
        ElectricPotentialChangeRateMillivoltPerSecond ElectricPotentialChangeRateUnits = "MillivoltPerSecond"
        // 
        ElectricPotentialChangeRateKilovoltPerSecond ElectricPotentialChangeRateUnits = "KilovoltPerSecond"
        // 
        ElectricPotentialChangeRateMegavoltPerSecond ElectricPotentialChangeRateUnits = "MegavoltPerSecond"
        // 
        ElectricPotentialChangeRateMicrovoltPerMicrosecond ElectricPotentialChangeRateUnits = "MicrovoltPerMicrosecond"
        // 
        ElectricPotentialChangeRateMillivoltPerMicrosecond ElectricPotentialChangeRateUnits = "MillivoltPerMicrosecond"
        // 
        ElectricPotentialChangeRateKilovoltPerMicrosecond ElectricPotentialChangeRateUnits = "KilovoltPerMicrosecond"
        // 
        ElectricPotentialChangeRateMegavoltPerMicrosecond ElectricPotentialChangeRateUnits = "MegavoltPerMicrosecond"
        // 
        ElectricPotentialChangeRateMicrovoltPerMinute ElectricPotentialChangeRateUnits = "MicrovoltPerMinute"
        // 
        ElectricPotentialChangeRateMillivoltPerMinute ElectricPotentialChangeRateUnits = "MillivoltPerMinute"
        // 
        ElectricPotentialChangeRateKilovoltPerMinute ElectricPotentialChangeRateUnits = "KilovoltPerMinute"
        // 
        ElectricPotentialChangeRateMegavoltPerMinute ElectricPotentialChangeRateUnits = "MegavoltPerMinute"
        // 
        ElectricPotentialChangeRateMicrovoltPerHour ElectricPotentialChangeRateUnits = "MicrovoltPerHour"
        // 
        ElectricPotentialChangeRateMillivoltPerHour ElectricPotentialChangeRateUnits = "MillivoltPerHour"
        // 
        ElectricPotentialChangeRateKilovoltPerHour ElectricPotentialChangeRateUnits = "KilovoltPerHour"
        // 
        ElectricPotentialChangeRateMegavoltPerHour ElectricPotentialChangeRateUnits = "MegavoltPerHour"
)

var internalElectricPotentialChangeRateUnitsMap = map[ElectricPotentialChangeRateUnits]bool{
	
	ElectricPotentialChangeRateVoltPerSecond: true,
	ElectricPotentialChangeRateVoltPerMicrosecond: true,
	ElectricPotentialChangeRateVoltPerMinute: true,
	ElectricPotentialChangeRateVoltPerHour: true,
	ElectricPotentialChangeRateMicrovoltPerSecond: true,
	ElectricPotentialChangeRateMillivoltPerSecond: true,
	ElectricPotentialChangeRateKilovoltPerSecond: true,
	ElectricPotentialChangeRateMegavoltPerSecond: true,
	ElectricPotentialChangeRateMicrovoltPerMicrosecond: true,
	ElectricPotentialChangeRateMillivoltPerMicrosecond: true,
	ElectricPotentialChangeRateKilovoltPerMicrosecond: true,
	ElectricPotentialChangeRateMegavoltPerMicrosecond: true,
	ElectricPotentialChangeRateMicrovoltPerMinute: true,
	ElectricPotentialChangeRateMillivoltPerMinute: true,
	ElectricPotentialChangeRateKilovoltPerMinute: true,
	ElectricPotentialChangeRateMegavoltPerMinute: true,
	ElectricPotentialChangeRateMicrovoltPerHour: true,
	ElectricPotentialChangeRateMillivoltPerHour: true,
	ElectricPotentialChangeRateKilovoltPerHour: true,
	ElectricPotentialChangeRateMegavoltPerHour: true,
}

// ElectricPotentialChangeRateDto represents a ElectricPotentialChangeRate measurement with a numerical value and its corresponding unit.
type ElectricPotentialChangeRateDto struct {
    // Value is the numerical representation of the ElectricPotentialChangeRate.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricPotentialChangeRate, as defined in the ElectricPotentialChangeRateUnits enumeration.
	Unit  ElectricPotentialChangeRateUnits `json:"unit" validate:"required,oneof=VoltPerSecond VoltPerMicrosecond VoltPerMinute VoltPerHour MicrovoltPerSecond MillivoltPerSecond KilovoltPerSecond MegavoltPerSecond MicrovoltPerMicrosecond MillivoltPerMicrosecond KilovoltPerMicrosecond MegavoltPerMicrosecond MicrovoltPerMinute MillivoltPerMinute KilovoltPerMinute MegavoltPerMinute MicrovoltPerHour MillivoltPerHour KilovoltPerHour MegavoltPerHour"`
}

// ElectricPotentialChangeRateDtoFactory groups methods for creating and serializing ElectricPotentialChangeRateDto objects.
type ElectricPotentialChangeRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialChangeRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricPotentialChangeRateDtoFactory) FromJSON(data []byte) (*ElectricPotentialChangeRateDto, error) {
	a := ElectricPotentialChangeRateDto{}

    // Parse JSON into ElectricPotentialChangeRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricPotentialChangeRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricPotentialChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricPotentialChangeRate represents a measurement in a of ElectricPotentialChangeRate.
//
// ElectricPotential change rate is the ratio of the electric potential change to the time during which the change occurred (value of electric potential changes per unit time).
type ElectricPotentialChangeRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    volts_per_secondLazy *float64 
    volts_per_microsecondLazy *float64 
    volts_per_minuteLazy *float64 
    volts_per_hourLazy *float64 
    microvolts_per_secondLazy *float64 
    millivolts_per_secondLazy *float64 
    kilovolts_per_secondLazy *float64 
    megavolts_per_secondLazy *float64 
    microvolts_per_microsecondLazy *float64 
    millivolts_per_microsecondLazy *float64 
    kilovolts_per_microsecondLazy *float64 
    megavolts_per_microsecondLazy *float64 
    microvolts_per_minuteLazy *float64 
    millivolts_per_minuteLazy *float64 
    kilovolts_per_minuteLazy *float64 
    megavolts_per_minuteLazy *float64 
    microvolts_per_hourLazy *float64 
    millivolts_per_hourLazy *float64 
    kilovolts_per_hourLazy *float64 
    megavolts_per_hourLazy *float64 
}

// ElectricPotentialChangeRateFactory groups methods for creating ElectricPotentialChangeRate instances.
type ElectricPotentialChangeRateFactory struct{}

// CreateElectricPotentialChangeRate creates a new ElectricPotentialChangeRate instance from the given value and unit.
func (uf ElectricPotentialChangeRateFactory) CreateElectricPotentialChangeRate(value float64, unit ElectricPotentialChangeRateUnits) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, unit)
}

// FromDto converts a ElectricPotentialChangeRateDto to a ElectricPotentialChangeRate instance.
func (uf ElectricPotentialChangeRateFactory) FromDto(dto ElectricPotentialChangeRateDto) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricPotentialChangeRate instance.
func (uf ElectricPotentialChangeRateFactory) FromDtoJSON(data []byte) (*ElectricPotentialChangeRate, error) {
	unitDto, err := ElectricPotentialChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricPotentialChangeRateDto from JSON: %w", err)
	}
	return ElectricPotentialChangeRateFactory{}.FromDto(*unitDto)
}


// FromVoltsPerSecond creates a new ElectricPotentialChangeRate instance from a value in VoltsPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerSecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerSecond)
}

// FromVoltsPerMicrosecond creates a new ElectricPotentialChangeRate instance from a value in VoltsPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMicrosecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMicrosecond)
}

// FromVoltsPerMinute creates a new ElectricPotentialChangeRate instance from a value in VoltsPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMinute(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMinute)
}

// FromVoltsPerHour creates a new ElectricPotentialChangeRate instance from a value in VoltsPerHour.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerHour(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerHour)
}

// FromMicrovoltsPerSecond creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerSecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerSecond)
}

// FromMillivoltsPerSecond creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerSecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerSecond)
}

// FromKilovoltsPerSecond creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerSecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerSecond)
}

// FromMegavoltsPerSecond creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerSecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerSecond)
}

// FromMicrovoltsPerMicrosecond creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMicrosecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMicrosecond)
}

// FromMillivoltsPerMicrosecond creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMicrosecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMicrosecond)
}

// FromKilovoltsPerMicrosecond creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMicrosecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMicrosecond)
}

// FromMegavoltsPerMicrosecond creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMicrosecond(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMicrosecond)
}

// FromMicrovoltsPerMinute creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMinute(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMinute)
}

// FromMillivoltsPerMinute creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMinute(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMinute)
}

// FromKilovoltsPerMinute creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMinute(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMinute)
}

// FromMegavoltsPerMinute creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMinute(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMinute)
}

// FromMicrovoltsPerHour creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerHour.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerHour(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerHour)
}

// FromMillivoltsPerHour creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerHour.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerHour(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerHour)
}

// FromKilovoltsPerHour creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerHour.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerHour(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerHour)
}

// FromMegavoltsPerHour creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerHour.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerHour(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerHour)
}


// newElectricPotentialChangeRate creates a new ElectricPotentialChangeRate.
func newElectricPotentialChangeRate(value float64, fromUnit ElectricPotentialChangeRateUnits) (*ElectricPotentialChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricPotentialChangeRateUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricPotentialChangeRateUnits", fromUnit)
	}
	a := &ElectricPotentialChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialChangeRate in VoltPerSecond unit (the base/default quantity).
func (a *ElectricPotentialChangeRate) BaseValue() float64 {
	return a.value
}


// VoltsPerSecond returns the ElectricPotentialChangeRate value in VoltsPerSecond.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerSecond() float64 {
	if a.volts_per_secondLazy != nil {
		return *a.volts_per_secondLazy
	}
	volts_per_second := a.convertFromBase(ElectricPotentialChangeRateVoltPerSecond)
	a.volts_per_secondLazy = &volts_per_second
	return volts_per_second
}

// VoltsPerMicrosecond returns the ElectricPotentialChangeRate value in VoltsPerMicrosecond.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerMicrosecond() float64 {
	if a.volts_per_microsecondLazy != nil {
		return *a.volts_per_microsecondLazy
	}
	volts_per_microsecond := a.convertFromBase(ElectricPotentialChangeRateVoltPerMicrosecond)
	a.volts_per_microsecondLazy = &volts_per_microsecond
	return volts_per_microsecond
}

// VoltsPerMinute returns the ElectricPotentialChangeRate value in VoltsPerMinute.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerMinute() float64 {
	if a.volts_per_minuteLazy != nil {
		return *a.volts_per_minuteLazy
	}
	volts_per_minute := a.convertFromBase(ElectricPotentialChangeRateVoltPerMinute)
	a.volts_per_minuteLazy = &volts_per_minute
	return volts_per_minute
}

// VoltsPerHour returns the ElectricPotentialChangeRate value in VoltsPerHour.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerHour() float64 {
	if a.volts_per_hourLazy != nil {
		return *a.volts_per_hourLazy
	}
	volts_per_hour := a.convertFromBase(ElectricPotentialChangeRateVoltPerHour)
	a.volts_per_hourLazy = &volts_per_hour
	return volts_per_hour
}

// MicrovoltsPerSecond returns the ElectricPotentialChangeRate value in MicrovoltsPerSecond.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerSecond() float64 {
	if a.microvolts_per_secondLazy != nil {
		return *a.microvolts_per_secondLazy
	}
	microvolts_per_second := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerSecond)
	a.microvolts_per_secondLazy = &microvolts_per_second
	return microvolts_per_second
}

// MillivoltsPerSecond returns the ElectricPotentialChangeRate value in MillivoltsPerSecond.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerSecond() float64 {
	if a.millivolts_per_secondLazy != nil {
		return *a.millivolts_per_secondLazy
	}
	millivolts_per_second := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerSecond)
	a.millivolts_per_secondLazy = &millivolts_per_second
	return millivolts_per_second
}

// KilovoltsPerSecond returns the ElectricPotentialChangeRate value in KilovoltsPerSecond.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerSecond() float64 {
	if a.kilovolts_per_secondLazy != nil {
		return *a.kilovolts_per_secondLazy
	}
	kilovolts_per_second := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerSecond)
	a.kilovolts_per_secondLazy = &kilovolts_per_second
	return kilovolts_per_second
}

// MegavoltsPerSecond returns the ElectricPotentialChangeRate value in MegavoltsPerSecond.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerSecond() float64 {
	if a.megavolts_per_secondLazy != nil {
		return *a.megavolts_per_secondLazy
	}
	megavolts_per_second := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerSecond)
	a.megavolts_per_secondLazy = &megavolts_per_second
	return megavolts_per_second
}

// MicrovoltsPerMicrosecond returns the ElectricPotentialChangeRate value in MicrovoltsPerMicrosecond.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerMicrosecond() float64 {
	if a.microvolts_per_microsecondLazy != nil {
		return *a.microvolts_per_microsecondLazy
	}
	microvolts_per_microsecond := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMicrosecond)
	a.microvolts_per_microsecondLazy = &microvolts_per_microsecond
	return microvolts_per_microsecond
}

// MillivoltsPerMicrosecond returns the ElectricPotentialChangeRate value in MillivoltsPerMicrosecond.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerMicrosecond() float64 {
	if a.millivolts_per_microsecondLazy != nil {
		return *a.millivolts_per_microsecondLazy
	}
	millivolts_per_microsecond := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMicrosecond)
	a.millivolts_per_microsecondLazy = &millivolts_per_microsecond
	return millivolts_per_microsecond
}

// KilovoltsPerMicrosecond returns the ElectricPotentialChangeRate value in KilovoltsPerMicrosecond.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerMicrosecond() float64 {
	if a.kilovolts_per_microsecondLazy != nil {
		return *a.kilovolts_per_microsecondLazy
	}
	kilovolts_per_microsecond := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMicrosecond)
	a.kilovolts_per_microsecondLazy = &kilovolts_per_microsecond
	return kilovolts_per_microsecond
}

// MegavoltsPerMicrosecond returns the ElectricPotentialChangeRate value in MegavoltsPerMicrosecond.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerMicrosecond() float64 {
	if a.megavolts_per_microsecondLazy != nil {
		return *a.megavolts_per_microsecondLazy
	}
	megavolts_per_microsecond := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMicrosecond)
	a.megavolts_per_microsecondLazy = &megavolts_per_microsecond
	return megavolts_per_microsecond
}

// MicrovoltsPerMinute returns the ElectricPotentialChangeRate value in MicrovoltsPerMinute.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerMinute() float64 {
	if a.microvolts_per_minuteLazy != nil {
		return *a.microvolts_per_minuteLazy
	}
	microvolts_per_minute := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMinute)
	a.microvolts_per_minuteLazy = &microvolts_per_minute
	return microvolts_per_minute
}

// MillivoltsPerMinute returns the ElectricPotentialChangeRate value in MillivoltsPerMinute.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerMinute() float64 {
	if a.millivolts_per_minuteLazy != nil {
		return *a.millivolts_per_minuteLazy
	}
	millivolts_per_minute := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMinute)
	a.millivolts_per_minuteLazy = &millivolts_per_minute
	return millivolts_per_minute
}

// KilovoltsPerMinute returns the ElectricPotentialChangeRate value in KilovoltsPerMinute.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerMinute() float64 {
	if a.kilovolts_per_minuteLazy != nil {
		return *a.kilovolts_per_minuteLazy
	}
	kilovolts_per_minute := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMinute)
	a.kilovolts_per_minuteLazy = &kilovolts_per_minute
	return kilovolts_per_minute
}

// MegavoltsPerMinute returns the ElectricPotentialChangeRate value in MegavoltsPerMinute.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerMinute() float64 {
	if a.megavolts_per_minuteLazy != nil {
		return *a.megavolts_per_minuteLazy
	}
	megavolts_per_minute := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMinute)
	a.megavolts_per_minuteLazy = &megavolts_per_minute
	return megavolts_per_minute
}

// MicrovoltsPerHour returns the ElectricPotentialChangeRate value in MicrovoltsPerHour.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerHour() float64 {
	if a.microvolts_per_hourLazy != nil {
		return *a.microvolts_per_hourLazy
	}
	microvolts_per_hour := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerHour)
	a.microvolts_per_hourLazy = &microvolts_per_hour
	return microvolts_per_hour
}

// MillivoltsPerHour returns the ElectricPotentialChangeRate value in MillivoltsPerHour.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerHour() float64 {
	if a.millivolts_per_hourLazy != nil {
		return *a.millivolts_per_hourLazy
	}
	millivolts_per_hour := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerHour)
	a.millivolts_per_hourLazy = &millivolts_per_hour
	return millivolts_per_hour
}

// KilovoltsPerHour returns the ElectricPotentialChangeRate value in KilovoltsPerHour.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerHour() float64 {
	if a.kilovolts_per_hourLazy != nil {
		return *a.kilovolts_per_hourLazy
	}
	kilovolts_per_hour := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerHour)
	a.kilovolts_per_hourLazy = &kilovolts_per_hour
	return kilovolts_per_hour
}

// MegavoltsPerHour returns the ElectricPotentialChangeRate value in MegavoltsPerHour.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerHour() float64 {
	if a.megavolts_per_hourLazy != nil {
		return *a.megavolts_per_hourLazy
	}
	megavolts_per_hour := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerHour)
	a.megavolts_per_hourLazy = &megavolts_per_hour
	return megavolts_per_hour
}


// ToDto creates a ElectricPotentialChangeRateDto representation from the ElectricPotentialChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltPerSecond by default.
func (a *ElectricPotentialChangeRate) ToDto(holdInUnit *ElectricPotentialChangeRateUnits) ElectricPotentialChangeRateDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialChangeRateVoltPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialChangeRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricPotentialChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltPerSecond by default.
func (a *ElectricPotentialChangeRate) ToDtoJSON(holdInUnit *ElectricPotentialChangeRateUnits) ([]byte, error) {
	// Convert to ElectricPotentialChangeRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricPotentialChangeRate to a specific unit value.
// The function uses the provided unit type (ElectricPotentialChangeRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricPotentialChangeRate) Convert(toUnit ElectricPotentialChangeRateUnits) float64 {
	switch toUnit { 
    case ElectricPotentialChangeRateVoltPerSecond:
		return a.VoltsPerSecond()
    case ElectricPotentialChangeRateVoltPerMicrosecond:
		return a.VoltsPerMicrosecond()
    case ElectricPotentialChangeRateVoltPerMinute:
		return a.VoltsPerMinute()
    case ElectricPotentialChangeRateVoltPerHour:
		return a.VoltsPerHour()
    case ElectricPotentialChangeRateMicrovoltPerSecond:
		return a.MicrovoltsPerSecond()
    case ElectricPotentialChangeRateMillivoltPerSecond:
		return a.MillivoltsPerSecond()
    case ElectricPotentialChangeRateKilovoltPerSecond:
		return a.KilovoltsPerSecond()
    case ElectricPotentialChangeRateMegavoltPerSecond:
		return a.MegavoltsPerSecond()
    case ElectricPotentialChangeRateMicrovoltPerMicrosecond:
		return a.MicrovoltsPerMicrosecond()
    case ElectricPotentialChangeRateMillivoltPerMicrosecond:
		return a.MillivoltsPerMicrosecond()
    case ElectricPotentialChangeRateKilovoltPerMicrosecond:
		return a.KilovoltsPerMicrosecond()
    case ElectricPotentialChangeRateMegavoltPerMicrosecond:
		return a.MegavoltsPerMicrosecond()
    case ElectricPotentialChangeRateMicrovoltPerMinute:
		return a.MicrovoltsPerMinute()
    case ElectricPotentialChangeRateMillivoltPerMinute:
		return a.MillivoltsPerMinute()
    case ElectricPotentialChangeRateKilovoltPerMinute:
		return a.KilovoltsPerMinute()
    case ElectricPotentialChangeRateMegavoltPerMinute:
		return a.MegavoltsPerMinute()
    case ElectricPotentialChangeRateMicrovoltPerHour:
		return a.MicrovoltsPerHour()
    case ElectricPotentialChangeRateMillivoltPerHour:
		return a.MillivoltsPerHour()
    case ElectricPotentialChangeRateKilovoltPerHour:
		return a.KilovoltsPerHour()
    case ElectricPotentialChangeRateMegavoltPerHour:
		return a.MegavoltsPerHour()
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialChangeRate) convertFromBase(toUnit ElectricPotentialChangeRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialChangeRateVoltPerSecond:
		return (value) 
	case ElectricPotentialChangeRateVoltPerMicrosecond:
		return (value / 1e6) 
	case ElectricPotentialChangeRateVoltPerMinute:
		return (value * 60) 
	case ElectricPotentialChangeRateVoltPerHour:
		return (value * 3600) 
	case ElectricPotentialChangeRateMicrovoltPerSecond:
		return ((value) / 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerSecond:
		return ((value) / 0.001) 
	case ElectricPotentialChangeRateKilovoltPerSecond:
		return ((value) / 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerSecond:
		return ((value) / 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerMicrosecond:
		return ((value / 1e6) / 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerMicrosecond:
		return ((value / 1e6) / 0.001) 
	case ElectricPotentialChangeRateKilovoltPerMicrosecond:
		return ((value / 1e6) / 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerMicrosecond:
		return ((value / 1e6) / 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerMinute:
		return ((value * 60) / 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerMinute:
		return ((value * 60) / 0.001) 
	case ElectricPotentialChangeRateKilovoltPerMinute:
		return ((value * 60) / 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerMinute:
		return ((value * 60) / 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerHour:
		return ((value * 3600) / 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerHour:
		return ((value * 3600) / 0.001) 
	case ElectricPotentialChangeRateKilovoltPerHour:
		return ((value * 3600) / 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerHour:
		return ((value * 3600) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialChangeRate) convertToBase(value float64, fromUnit ElectricPotentialChangeRateUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialChangeRateVoltPerSecond:
		return (value) 
	case ElectricPotentialChangeRateVoltPerMicrosecond:
		return (value * 1e6) 
	case ElectricPotentialChangeRateVoltPerMinute:
		return (value / 60) 
	case ElectricPotentialChangeRateVoltPerHour:
		return (value / 3600) 
	case ElectricPotentialChangeRateMicrovoltPerSecond:
		return ((value) * 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerSecond:
		return ((value) * 0.001) 
	case ElectricPotentialChangeRateKilovoltPerSecond:
		return ((value) * 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerSecond:
		return ((value) * 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerMicrosecond:
		return ((value * 1e6) * 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerMicrosecond:
		return ((value * 1e6) * 0.001) 
	case ElectricPotentialChangeRateKilovoltPerMicrosecond:
		return ((value * 1e6) * 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerMicrosecond:
		return ((value * 1e6) * 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerMinute:
		return ((value / 60) * 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerMinute:
		return ((value / 60) * 0.001) 
	case ElectricPotentialChangeRateKilovoltPerMinute:
		return ((value / 60) * 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerMinute:
		return ((value / 60) * 1000000.0) 
	case ElectricPotentialChangeRateMicrovoltPerHour:
		return ((value / 3600) * 1e-06) 
	case ElectricPotentialChangeRateMillivoltPerHour:
		return ((value / 3600) * 0.001) 
	case ElectricPotentialChangeRateKilovoltPerHour:
		return ((value / 3600) * 1000.0) 
	case ElectricPotentialChangeRateMegavoltPerHour:
		return ((value / 3600) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricPotentialChangeRate in the default unit (VoltPerSecond),
// formatted to two decimal places.
func (a ElectricPotentialChangeRate) String() string {
	return a.ToString(ElectricPotentialChangeRateVoltPerSecond, 2)
}

// ToString formats the ElectricPotentialChangeRate value as a string with the specified unit and fractional digits.
// It converts the ElectricPotentialChangeRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricPotentialChangeRate value will be converted (e.g., VoltPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricPotentialChangeRate with the unit abbreviation.
func (a *ElectricPotentialChangeRate) ToString(unit ElectricPotentialChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricPotentialChangeRateAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricPotentialChangeRateAbbreviation(unit))
}

// Equals checks if the given ElectricPotentialChangeRate is equal to the current ElectricPotentialChangeRate.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to compare against.
//
// Returns:
//    bool: Returns true if both ElectricPotentialChangeRate are equal, false otherwise.
func (a *ElectricPotentialChangeRate) Equals(other *ElectricPotentialChangeRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricPotentialChangeRate with another ElectricPotentialChangeRate.
// It returns -1 if the current ElectricPotentialChangeRate is less than the other ElectricPotentialChangeRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to compare against.
//
// Returns:
//    int: -1 if the current ElectricPotentialChangeRate is less, 1 if greater, and 0 if equal.
func (a *ElectricPotentialChangeRate) CompareTo(other *ElectricPotentialChangeRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricPotentialChangeRate to the current ElectricPotentialChangeRate and returns the result.
// The result is a new ElectricPotentialChangeRate instance with the sum of the values.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to add to the current ElectricPotentialChangeRate.
//
// Returns:
//    *ElectricPotentialChangeRate: A new ElectricPotentialChangeRate instance representing the sum of both ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Add(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricPotentialChangeRate from the current ElectricPotentialChangeRate and returns the result.
// The result is a new ElectricPotentialChangeRate instance with the difference of the values.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to subtract from the current ElectricPotentialChangeRate.
//
// Returns:
//    *ElectricPotentialChangeRate: A new ElectricPotentialChangeRate instance representing the difference of both ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Subtract(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricPotentialChangeRate by the given ElectricPotentialChangeRate and returns the result.
// The result is a new ElectricPotentialChangeRate instance with the product of the values.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to multiply with the current ElectricPotentialChangeRate.
//
// Returns:
//    *ElectricPotentialChangeRate: A new ElectricPotentialChangeRate instance representing the product of both ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Multiply(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricPotentialChangeRate by the given ElectricPotentialChangeRate and returns the result.
// The result is a new ElectricPotentialChangeRate instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricPotentialChangeRate to divide the current ElectricPotentialChangeRate by.
//
// Returns:
//    *ElectricPotentialChangeRate: A new ElectricPotentialChangeRate instance representing the quotient of both ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Divide(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value / other.BaseValue()}
}

// GetElectricPotentialChangeRateAbbreviation gets the unit abbreviation.
func GetElectricPotentialChangeRateAbbreviation(unit ElectricPotentialChangeRateUnits) string {
	switch unit { 
	case ElectricPotentialChangeRateVoltPerSecond:
		return "V/s" 
	case ElectricPotentialChangeRateVoltPerMicrosecond:
		return "V/μs" 
	case ElectricPotentialChangeRateVoltPerMinute:
		return "V/min" 
	case ElectricPotentialChangeRateVoltPerHour:
		return "V/h" 
	case ElectricPotentialChangeRateMicrovoltPerSecond:
		return "μV/s" 
	case ElectricPotentialChangeRateMillivoltPerSecond:
		return "mV/s" 
	case ElectricPotentialChangeRateKilovoltPerSecond:
		return "kV/s" 
	case ElectricPotentialChangeRateMegavoltPerSecond:
		return "MV/s" 
	case ElectricPotentialChangeRateMicrovoltPerMicrosecond:
		return "μV/μs" 
	case ElectricPotentialChangeRateMillivoltPerMicrosecond:
		return "mV/μs" 
	case ElectricPotentialChangeRateKilovoltPerMicrosecond:
		return "kV/μs" 
	case ElectricPotentialChangeRateMegavoltPerMicrosecond:
		return "MV/μs" 
	case ElectricPotentialChangeRateMicrovoltPerMinute:
		return "μV/min" 
	case ElectricPotentialChangeRateMillivoltPerMinute:
		return "mV/min" 
	case ElectricPotentialChangeRateKilovoltPerMinute:
		return "kV/min" 
	case ElectricPotentialChangeRateMegavoltPerMinute:
		return "MV/min" 
	case ElectricPotentialChangeRateMicrovoltPerHour:
		return "μV/h" 
	case ElectricPotentialChangeRateMillivoltPerHour:
		return "mV/h" 
	case ElectricPotentialChangeRateKilovoltPerHour:
		return "kV/h" 
	case ElectricPotentialChangeRateMegavoltPerHour:
		return "MV/h" 
	default:
		return ""
	}
}