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

// ElectricPotentialChangeRateDto represents a ElectricPotentialChangeRate measurement with a numerical value and its corresponding unit.
type ElectricPotentialChangeRateDto struct {
    // Value is the numerical representation of the ElectricPotentialChangeRate.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricPotentialChangeRate, as defined in the ElectricPotentialChangeRateUnits enumeration.
	Unit  ElectricPotentialChangeRateUnits `json:"unit"`
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
    
    volts_per_secondsLazy *float64 
    volts_per_microsecondsLazy *float64 
    volts_per_minutesLazy *float64 
    volts_per_hoursLazy *float64 
    microvolts_per_secondsLazy *float64 
    millivolts_per_secondsLazy *float64 
    kilovolts_per_secondsLazy *float64 
    megavolts_per_secondsLazy *float64 
    microvolts_per_microsecondsLazy *float64 
    millivolts_per_microsecondsLazy *float64 
    kilovolts_per_microsecondsLazy *float64 
    megavolts_per_microsecondsLazy *float64 
    microvolts_per_minutesLazy *float64 
    millivolts_per_minutesLazy *float64 
    kilovolts_per_minutesLazy *float64 
    megavolts_per_minutesLazy *float64 
    microvolts_per_hoursLazy *float64 
    millivolts_per_hoursLazy *float64 
    kilovolts_per_hoursLazy *float64 
    megavolts_per_hoursLazy *float64 
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


// FromVoltsPerSeconds creates a new ElectricPotentialChangeRate instance from a value in VoltsPerSeconds.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerSecond)
}

// FromVoltsPerMicroseconds creates a new ElectricPotentialChangeRate instance from a value in VoltsPerMicroseconds.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMicrosecond)
}

// FromVoltsPerMinutes creates a new ElectricPotentialChangeRate instance from a value in VoltsPerMinutes.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMinute)
}

// FromVoltsPerHours creates a new ElectricPotentialChangeRate instance from a value in VoltsPerHours.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerHour)
}

// FromMicrovoltsPerSeconds creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerSeconds.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerSecond)
}

// FromMillivoltsPerSeconds creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerSeconds.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerSecond)
}

// FromKilovoltsPerSeconds creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerSeconds.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerSecond)
}

// FromMegavoltsPerSeconds creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerSeconds.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerSecond)
}

// FromMicrovoltsPerMicroseconds creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerMicroseconds.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMicrosecond)
}

// FromMillivoltsPerMicroseconds creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerMicroseconds.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMicrosecond)
}

// FromKilovoltsPerMicroseconds creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerMicroseconds.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMicrosecond)
}

// FromMegavoltsPerMicroseconds creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerMicroseconds.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMicrosecond)
}

// FromMicrovoltsPerMinutes creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerMinutes.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMinute)
}

// FromMillivoltsPerMinutes creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerMinutes.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMinute)
}

// FromKilovoltsPerMinutes creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerMinutes.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMinute)
}

// FromMegavoltsPerMinutes creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerMinutes.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMinute)
}

// FromMicrovoltsPerHours creates a new ElectricPotentialChangeRate instance from a value in MicrovoltsPerHours.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerHour)
}

// FromMillivoltsPerHours creates a new ElectricPotentialChangeRate instance from a value in MillivoltsPerHours.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerHour)
}

// FromKilovoltsPerHours creates a new ElectricPotentialChangeRate instance from a value in KilovoltsPerHours.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerHour)
}

// FromMegavoltsPerHours creates a new ElectricPotentialChangeRate instance from a value in MegavoltsPerHours.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerHour)
}


// newElectricPotentialChangeRate creates a new ElectricPotentialChangeRate.
func newElectricPotentialChangeRate(value float64, fromUnit ElectricPotentialChangeRateUnits) (*ElectricPotentialChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotentialChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialChangeRate in VoltPerSecond unit (the base/default quantity).
func (a *ElectricPotentialChangeRate) BaseValue() float64 {
	return a.value
}


// VoltsPerSeconds returns the ElectricPotentialChangeRate value in VoltsPerSeconds.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerSeconds() float64 {
	if a.volts_per_secondsLazy != nil {
		return *a.volts_per_secondsLazy
	}
	volts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateVoltPerSecond)
	a.volts_per_secondsLazy = &volts_per_seconds
	return volts_per_seconds
}

// VoltsPerMicroseconds returns the ElectricPotentialChangeRate value in VoltsPerMicroseconds.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerMicroseconds() float64 {
	if a.volts_per_microsecondsLazy != nil {
		return *a.volts_per_microsecondsLazy
	}
	volts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateVoltPerMicrosecond)
	a.volts_per_microsecondsLazy = &volts_per_microseconds
	return volts_per_microseconds
}

// VoltsPerMinutes returns the ElectricPotentialChangeRate value in VoltsPerMinutes.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerMinutes() float64 {
	if a.volts_per_minutesLazy != nil {
		return *a.volts_per_minutesLazy
	}
	volts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateVoltPerMinute)
	a.volts_per_minutesLazy = &volts_per_minutes
	return volts_per_minutes
}

// VoltsPerHours returns the ElectricPotentialChangeRate value in VoltsPerHours.
//
// 
func (a *ElectricPotentialChangeRate) VoltsPerHours() float64 {
	if a.volts_per_hoursLazy != nil {
		return *a.volts_per_hoursLazy
	}
	volts_per_hours := a.convertFromBase(ElectricPotentialChangeRateVoltPerHour)
	a.volts_per_hoursLazy = &volts_per_hours
	return volts_per_hours
}

// MicrovoltsPerSeconds returns the ElectricPotentialChangeRate value in MicrovoltsPerSeconds.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerSeconds() float64 {
	if a.microvolts_per_secondsLazy != nil {
		return *a.microvolts_per_secondsLazy
	}
	microvolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerSecond)
	a.microvolts_per_secondsLazy = &microvolts_per_seconds
	return microvolts_per_seconds
}

// MillivoltsPerSeconds returns the ElectricPotentialChangeRate value in MillivoltsPerSeconds.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerSeconds() float64 {
	if a.millivolts_per_secondsLazy != nil {
		return *a.millivolts_per_secondsLazy
	}
	millivolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerSecond)
	a.millivolts_per_secondsLazy = &millivolts_per_seconds
	return millivolts_per_seconds
}

// KilovoltsPerSeconds returns the ElectricPotentialChangeRate value in KilovoltsPerSeconds.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerSeconds() float64 {
	if a.kilovolts_per_secondsLazy != nil {
		return *a.kilovolts_per_secondsLazy
	}
	kilovolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerSecond)
	a.kilovolts_per_secondsLazy = &kilovolts_per_seconds
	return kilovolts_per_seconds
}

// MegavoltsPerSeconds returns the ElectricPotentialChangeRate value in MegavoltsPerSeconds.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerSeconds() float64 {
	if a.megavolts_per_secondsLazy != nil {
		return *a.megavolts_per_secondsLazy
	}
	megavolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerSecond)
	a.megavolts_per_secondsLazy = &megavolts_per_seconds
	return megavolts_per_seconds
}

// MicrovoltsPerMicroseconds returns the ElectricPotentialChangeRate value in MicrovoltsPerMicroseconds.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerMicroseconds() float64 {
	if a.microvolts_per_microsecondsLazy != nil {
		return *a.microvolts_per_microsecondsLazy
	}
	microvolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMicrosecond)
	a.microvolts_per_microsecondsLazy = &microvolts_per_microseconds
	return microvolts_per_microseconds
}

// MillivoltsPerMicroseconds returns the ElectricPotentialChangeRate value in MillivoltsPerMicroseconds.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerMicroseconds() float64 {
	if a.millivolts_per_microsecondsLazy != nil {
		return *a.millivolts_per_microsecondsLazy
	}
	millivolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMicrosecond)
	a.millivolts_per_microsecondsLazy = &millivolts_per_microseconds
	return millivolts_per_microseconds
}

// KilovoltsPerMicroseconds returns the ElectricPotentialChangeRate value in KilovoltsPerMicroseconds.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerMicroseconds() float64 {
	if a.kilovolts_per_microsecondsLazy != nil {
		return *a.kilovolts_per_microsecondsLazy
	}
	kilovolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMicrosecond)
	a.kilovolts_per_microsecondsLazy = &kilovolts_per_microseconds
	return kilovolts_per_microseconds
}

// MegavoltsPerMicroseconds returns the ElectricPotentialChangeRate value in MegavoltsPerMicroseconds.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerMicroseconds() float64 {
	if a.megavolts_per_microsecondsLazy != nil {
		return *a.megavolts_per_microsecondsLazy
	}
	megavolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMicrosecond)
	a.megavolts_per_microsecondsLazy = &megavolts_per_microseconds
	return megavolts_per_microseconds
}

// MicrovoltsPerMinutes returns the ElectricPotentialChangeRate value in MicrovoltsPerMinutes.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerMinutes() float64 {
	if a.microvolts_per_minutesLazy != nil {
		return *a.microvolts_per_minutesLazy
	}
	microvolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMinute)
	a.microvolts_per_minutesLazy = &microvolts_per_minutes
	return microvolts_per_minutes
}

// MillivoltsPerMinutes returns the ElectricPotentialChangeRate value in MillivoltsPerMinutes.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerMinutes() float64 {
	if a.millivolts_per_minutesLazy != nil {
		return *a.millivolts_per_minutesLazy
	}
	millivolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMinute)
	a.millivolts_per_minutesLazy = &millivolts_per_minutes
	return millivolts_per_minutes
}

// KilovoltsPerMinutes returns the ElectricPotentialChangeRate value in KilovoltsPerMinutes.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerMinutes() float64 {
	if a.kilovolts_per_minutesLazy != nil {
		return *a.kilovolts_per_minutesLazy
	}
	kilovolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMinute)
	a.kilovolts_per_minutesLazy = &kilovolts_per_minutes
	return kilovolts_per_minutes
}

// MegavoltsPerMinutes returns the ElectricPotentialChangeRate value in MegavoltsPerMinutes.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerMinutes() float64 {
	if a.megavolts_per_minutesLazy != nil {
		return *a.megavolts_per_minutesLazy
	}
	megavolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMinute)
	a.megavolts_per_minutesLazy = &megavolts_per_minutes
	return megavolts_per_minutes
}

// MicrovoltsPerHours returns the ElectricPotentialChangeRate value in MicrovoltsPerHours.
//
// 
func (a *ElectricPotentialChangeRate) MicrovoltsPerHours() float64 {
	if a.microvolts_per_hoursLazy != nil {
		return *a.microvolts_per_hoursLazy
	}
	microvolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerHour)
	a.microvolts_per_hoursLazy = &microvolts_per_hours
	return microvolts_per_hours
}

// MillivoltsPerHours returns the ElectricPotentialChangeRate value in MillivoltsPerHours.
//
// 
func (a *ElectricPotentialChangeRate) MillivoltsPerHours() float64 {
	if a.millivolts_per_hoursLazy != nil {
		return *a.millivolts_per_hoursLazy
	}
	millivolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerHour)
	a.millivolts_per_hoursLazy = &millivolts_per_hours
	return millivolts_per_hours
}

// KilovoltsPerHours returns the ElectricPotentialChangeRate value in KilovoltsPerHours.
//
// 
func (a *ElectricPotentialChangeRate) KilovoltsPerHours() float64 {
	if a.kilovolts_per_hoursLazy != nil {
		return *a.kilovolts_per_hoursLazy
	}
	kilovolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerHour)
	a.kilovolts_per_hoursLazy = &kilovolts_per_hours
	return kilovolts_per_hours
}

// MegavoltsPerHours returns the ElectricPotentialChangeRate value in MegavoltsPerHours.
//
// 
func (a *ElectricPotentialChangeRate) MegavoltsPerHours() float64 {
	if a.megavolts_per_hoursLazy != nil {
		return *a.megavolts_per_hoursLazy
	}
	megavolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerHour)
	a.megavolts_per_hoursLazy = &megavolts_per_hours
	return megavolts_per_hours
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
		return a.VoltsPerSeconds()
    case ElectricPotentialChangeRateVoltPerMicrosecond:
		return a.VoltsPerMicroseconds()
    case ElectricPotentialChangeRateVoltPerMinute:
		return a.VoltsPerMinutes()
    case ElectricPotentialChangeRateVoltPerHour:
		return a.VoltsPerHours()
    case ElectricPotentialChangeRateMicrovoltPerSecond:
		return a.MicrovoltsPerSeconds()
    case ElectricPotentialChangeRateMillivoltPerSecond:
		return a.MillivoltsPerSeconds()
    case ElectricPotentialChangeRateKilovoltPerSecond:
		return a.KilovoltsPerSeconds()
    case ElectricPotentialChangeRateMegavoltPerSecond:
		return a.MegavoltsPerSeconds()
    case ElectricPotentialChangeRateMicrovoltPerMicrosecond:
		return a.MicrovoltsPerMicroseconds()
    case ElectricPotentialChangeRateMillivoltPerMicrosecond:
		return a.MillivoltsPerMicroseconds()
    case ElectricPotentialChangeRateKilovoltPerMicrosecond:
		return a.KilovoltsPerMicroseconds()
    case ElectricPotentialChangeRateMegavoltPerMicrosecond:
		return a.MegavoltsPerMicroseconds()
    case ElectricPotentialChangeRateMicrovoltPerMinute:
		return a.MicrovoltsPerMinutes()
    case ElectricPotentialChangeRateMillivoltPerMinute:
		return a.MillivoltsPerMinutes()
    case ElectricPotentialChangeRateKilovoltPerMinute:
		return a.KilovoltsPerMinutes()
    case ElectricPotentialChangeRateMegavoltPerMinute:
		return a.MegavoltsPerMinutes()
    case ElectricPotentialChangeRateMicrovoltPerHour:
		return a.MicrovoltsPerHours()
    case ElectricPotentialChangeRateMillivoltPerHour:
		return a.MillivoltsPerHours()
    case ElectricPotentialChangeRateKilovoltPerHour:
		return a.KilovoltsPerHours()
    case ElectricPotentialChangeRateMegavoltPerHour:
		return a.MegavoltsPerHours()
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