package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialChangeRateUnits enumeration
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

// ElectricPotentialChangeRateDto represents an ElectricPotentialChangeRate
type ElectricPotentialChangeRateDto struct {
	Value float64
	Unit  ElectricPotentialChangeRateUnits
}

// ElectricPotentialChangeRateDtoFactory struct to group related functions
type ElectricPotentialChangeRateDtoFactory struct{}

func (udf ElectricPotentialChangeRateDtoFactory) FromJSON(data []byte) (*ElectricPotentialChangeRateDto, error) {
	a := ElectricPotentialChangeRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricPotentialChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricPotentialChangeRate struct
type ElectricPotentialChangeRate struct {
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

// ElectricPotentialChangeRateFactory struct to group related functions
type ElectricPotentialChangeRateFactory struct{}

func (uf ElectricPotentialChangeRateFactory) CreateElectricPotentialChangeRate(value float64, unit ElectricPotentialChangeRateUnits) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, unit)
}

func (uf ElectricPotentialChangeRateFactory) FromDto(dto ElectricPotentialChangeRateDto) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(dto.Value, dto.Unit)
}

func (uf ElectricPotentialChangeRateFactory) FromDtoJSON(data []byte) (*ElectricPotentialChangeRate, error) {
	unitDto, err := ElectricPotentialChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricPotentialChangeRateFactory{}.FromDto(*unitDto)
}


// FromVoltPerSecond creates a new ElectricPotentialChangeRate instance from VoltPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerSecond)
}

// FromVoltPerMicrosecond creates a new ElectricPotentialChangeRate instance from VoltPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMicrosecond)
}

// FromVoltPerMinute creates a new ElectricPotentialChangeRate instance from VoltPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerMinute)
}

// FromVoltPerHour creates a new ElectricPotentialChangeRate instance from VoltPerHour.
func (uf ElectricPotentialChangeRateFactory) FromVoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateVoltPerHour)
}

// FromMicrovoltPerSecond creates a new ElectricPotentialChangeRate instance from MicrovoltPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerSecond)
}

// FromMillivoltPerSecond creates a new ElectricPotentialChangeRate instance from MillivoltPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerSecond)
}

// FromKilovoltPerSecond creates a new ElectricPotentialChangeRate instance from KilovoltPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerSecond)
}

// FromMegavoltPerSecond creates a new ElectricPotentialChangeRate instance from MegavoltPerSecond.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerSeconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerSecond)
}

// FromMicrovoltPerMicrosecond creates a new ElectricPotentialChangeRate instance from MicrovoltPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMicrosecond)
}

// FromMillivoltPerMicrosecond creates a new ElectricPotentialChangeRate instance from MillivoltPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMicrosecond)
}

// FromKilovoltPerMicrosecond creates a new ElectricPotentialChangeRate instance from KilovoltPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMicrosecond)
}

// FromMegavoltPerMicrosecond creates a new ElectricPotentialChangeRate instance from MegavoltPerMicrosecond.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMicroseconds(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMicrosecond)
}

// FromMicrovoltPerMinute creates a new ElectricPotentialChangeRate instance from MicrovoltPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerMinute)
}

// FromMillivoltPerMinute creates a new ElectricPotentialChangeRate instance from MillivoltPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerMinute)
}

// FromKilovoltPerMinute creates a new ElectricPotentialChangeRate instance from KilovoltPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerMinute)
}

// FromMegavoltPerMinute creates a new ElectricPotentialChangeRate instance from MegavoltPerMinute.
func (uf ElectricPotentialChangeRateFactory) FromMegavoltsPerMinutes(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMegavoltPerMinute)
}

// FromMicrovoltPerHour creates a new ElectricPotentialChangeRate instance from MicrovoltPerHour.
func (uf ElectricPotentialChangeRateFactory) FromMicrovoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMicrovoltPerHour)
}

// FromMillivoltPerHour creates a new ElectricPotentialChangeRate instance from MillivoltPerHour.
func (uf ElectricPotentialChangeRateFactory) FromMillivoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateMillivoltPerHour)
}

// FromKilovoltPerHour creates a new ElectricPotentialChangeRate instance from KilovoltPerHour.
func (uf ElectricPotentialChangeRateFactory) FromKilovoltsPerHours(value float64) (*ElectricPotentialChangeRate, error) {
	return newElectricPotentialChangeRate(value, ElectricPotentialChangeRateKilovoltPerHour)
}

// FromMegavoltPerHour creates a new ElectricPotentialChangeRate instance from MegavoltPerHour.
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

// BaseValue returns the base value of ElectricPotentialChangeRate in VoltPerSecond.
func (a *ElectricPotentialChangeRate) BaseValue() float64 {
	return a.value
}


// VoltPerSecond returns the value in VoltPerSecond.
func (a *ElectricPotentialChangeRate) VoltsPerSeconds() float64 {
	if a.volts_per_secondsLazy != nil {
		return *a.volts_per_secondsLazy
	}
	volts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateVoltPerSecond)
	a.volts_per_secondsLazy = &volts_per_seconds
	return volts_per_seconds
}

// VoltPerMicrosecond returns the value in VoltPerMicrosecond.
func (a *ElectricPotentialChangeRate) VoltsPerMicroseconds() float64 {
	if a.volts_per_microsecondsLazy != nil {
		return *a.volts_per_microsecondsLazy
	}
	volts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateVoltPerMicrosecond)
	a.volts_per_microsecondsLazy = &volts_per_microseconds
	return volts_per_microseconds
}

// VoltPerMinute returns the value in VoltPerMinute.
func (a *ElectricPotentialChangeRate) VoltsPerMinutes() float64 {
	if a.volts_per_minutesLazy != nil {
		return *a.volts_per_minutesLazy
	}
	volts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateVoltPerMinute)
	a.volts_per_minutesLazy = &volts_per_minutes
	return volts_per_minutes
}

// VoltPerHour returns the value in VoltPerHour.
func (a *ElectricPotentialChangeRate) VoltsPerHours() float64 {
	if a.volts_per_hoursLazy != nil {
		return *a.volts_per_hoursLazy
	}
	volts_per_hours := a.convertFromBase(ElectricPotentialChangeRateVoltPerHour)
	a.volts_per_hoursLazy = &volts_per_hours
	return volts_per_hours
}

// MicrovoltPerSecond returns the value in MicrovoltPerSecond.
func (a *ElectricPotentialChangeRate) MicrovoltsPerSeconds() float64 {
	if a.microvolts_per_secondsLazy != nil {
		return *a.microvolts_per_secondsLazy
	}
	microvolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerSecond)
	a.microvolts_per_secondsLazy = &microvolts_per_seconds
	return microvolts_per_seconds
}

// MillivoltPerSecond returns the value in MillivoltPerSecond.
func (a *ElectricPotentialChangeRate) MillivoltsPerSeconds() float64 {
	if a.millivolts_per_secondsLazy != nil {
		return *a.millivolts_per_secondsLazy
	}
	millivolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerSecond)
	a.millivolts_per_secondsLazy = &millivolts_per_seconds
	return millivolts_per_seconds
}

// KilovoltPerSecond returns the value in KilovoltPerSecond.
func (a *ElectricPotentialChangeRate) KilovoltsPerSeconds() float64 {
	if a.kilovolts_per_secondsLazy != nil {
		return *a.kilovolts_per_secondsLazy
	}
	kilovolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerSecond)
	a.kilovolts_per_secondsLazy = &kilovolts_per_seconds
	return kilovolts_per_seconds
}

// MegavoltPerSecond returns the value in MegavoltPerSecond.
func (a *ElectricPotentialChangeRate) MegavoltsPerSeconds() float64 {
	if a.megavolts_per_secondsLazy != nil {
		return *a.megavolts_per_secondsLazy
	}
	megavolts_per_seconds := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerSecond)
	a.megavolts_per_secondsLazy = &megavolts_per_seconds
	return megavolts_per_seconds
}

// MicrovoltPerMicrosecond returns the value in MicrovoltPerMicrosecond.
func (a *ElectricPotentialChangeRate) MicrovoltsPerMicroseconds() float64 {
	if a.microvolts_per_microsecondsLazy != nil {
		return *a.microvolts_per_microsecondsLazy
	}
	microvolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMicrosecond)
	a.microvolts_per_microsecondsLazy = &microvolts_per_microseconds
	return microvolts_per_microseconds
}

// MillivoltPerMicrosecond returns the value in MillivoltPerMicrosecond.
func (a *ElectricPotentialChangeRate) MillivoltsPerMicroseconds() float64 {
	if a.millivolts_per_microsecondsLazy != nil {
		return *a.millivolts_per_microsecondsLazy
	}
	millivolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMicrosecond)
	a.millivolts_per_microsecondsLazy = &millivolts_per_microseconds
	return millivolts_per_microseconds
}

// KilovoltPerMicrosecond returns the value in KilovoltPerMicrosecond.
func (a *ElectricPotentialChangeRate) KilovoltsPerMicroseconds() float64 {
	if a.kilovolts_per_microsecondsLazy != nil {
		return *a.kilovolts_per_microsecondsLazy
	}
	kilovolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMicrosecond)
	a.kilovolts_per_microsecondsLazy = &kilovolts_per_microseconds
	return kilovolts_per_microseconds
}

// MegavoltPerMicrosecond returns the value in MegavoltPerMicrosecond.
func (a *ElectricPotentialChangeRate) MegavoltsPerMicroseconds() float64 {
	if a.megavolts_per_microsecondsLazy != nil {
		return *a.megavolts_per_microsecondsLazy
	}
	megavolts_per_microseconds := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMicrosecond)
	a.megavolts_per_microsecondsLazy = &megavolts_per_microseconds
	return megavolts_per_microseconds
}

// MicrovoltPerMinute returns the value in MicrovoltPerMinute.
func (a *ElectricPotentialChangeRate) MicrovoltsPerMinutes() float64 {
	if a.microvolts_per_minutesLazy != nil {
		return *a.microvolts_per_minutesLazy
	}
	microvolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerMinute)
	a.microvolts_per_minutesLazy = &microvolts_per_minutes
	return microvolts_per_minutes
}

// MillivoltPerMinute returns the value in MillivoltPerMinute.
func (a *ElectricPotentialChangeRate) MillivoltsPerMinutes() float64 {
	if a.millivolts_per_minutesLazy != nil {
		return *a.millivolts_per_minutesLazy
	}
	millivolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerMinute)
	a.millivolts_per_minutesLazy = &millivolts_per_minutes
	return millivolts_per_minutes
}

// KilovoltPerMinute returns the value in KilovoltPerMinute.
func (a *ElectricPotentialChangeRate) KilovoltsPerMinutes() float64 {
	if a.kilovolts_per_minutesLazy != nil {
		return *a.kilovolts_per_minutesLazy
	}
	kilovolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerMinute)
	a.kilovolts_per_minutesLazy = &kilovolts_per_minutes
	return kilovolts_per_minutes
}

// MegavoltPerMinute returns the value in MegavoltPerMinute.
func (a *ElectricPotentialChangeRate) MegavoltsPerMinutes() float64 {
	if a.megavolts_per_minutesLazy != nil {
		return *a.megavolts_per_minutesLazy
	}
	megavolts_per_minutes := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerMinute)
	a.megavolts_per_minutesLazy = &megavolts_per_minutes
	return megavolts_per_minutes
}

// MicrovoltPerHour returns the value in MicrovoltPerHour.
func (a *ElectricPotentialChangeRate) MicrovoltsPerHours() float64 {
	if a.microvolts_per_hoursLazy != nil {
		return *a.microvolts_per_hoursLazy
	}
	microvolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMicrovoltPerHour)
	a.microvolts_per_hoursLazy = &microvolts_per_hours
	return microvolts_per_hours
}

// MillivoltPerHour returns the value in MillivoltPerHour.
func (a *ElectricPotentialChangeRate) MillivoltsPerHours() float64 {
	if a.millivolts_per_hoursLazy != nil {
		return *a.millivolts_per_hoursLazy
	}
	millivolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMillivoltPerHour)
	a.millivolts_per_hoursLazy = &millivolts_per_hours
	return millivolts_per_hours
}

// KilovoltPerHour returns the value in KilovoltPerHour.
func (a *ElectricPotentialChangeRate) KilovoltsPerHours() float64 {
	if a.kilovolts_per_hoursLazy != nil {
		return *a.kilovolts_per_hoursLazy
	}
	kilovolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateKilovoltPerHour)
	a.kilovolts_per_hoursLazy = &kilovolts_per_hours
	return kilovolts_per_hours
}

// MegavoltPerHour returns the value in MegavoltPerHour.
func (a *ElectricPotentialChangeRate) MegavoltsPerHours() float64 {
	if a.megavolts_per_hoursLazy != nil {
		return *a.megavolts_per_hoursLazy
	}
	megavolts_per_hours := a.convertFromBase(ElectricPotentialChangeRateMegavoltPerHour)
	a.megavolts_per_hoursLazy = &megavolts_per_hours
	return megavolts_per_hours
}


// ToDto creates an ElectricPotentialChangeRateDto representation.
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

// ToDtoJSON creates an ElectricPotentialChangeRateDto representation.
func (a *ElectricPotentialChangeRate) ToDtoJSON(holdInUnit *ElectricPotentialChangeRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricPotentialChangeRate to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ElectricPotentialChangeRate) String() string {
	return a.ToString(ElectricPotentialChangeRateVoltPerSecond, 2)
}

// ToString formats the ElectricPotentialChangeRate to string.
// fractionalDigits -1 for not mention
func (a *ElectricPotentialChangeRate) ToString(unit ElectricPotentialChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricPotentialChangeRate) getUnitAbbreviation(unit ElectricPotentialChangeRateUnits) string {
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

// Check if the given ElectricPotentialChangeRate are equals to the current ElectricPotentialChangeRate
func (a *ElectricPotentialChangeRate) Equals(other *ElectricPotentialChangeRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricPotentialChangeRate are equals to the current ElectricPotentialChangeRate
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

// Add the given ElectricPotentialChangeRate to the current ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Add(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricPotentialChangeRate to the current ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Subtract(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricPotentialChangeRate to the current ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Multiply(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value * other.BaseValue()}
}

// Divide the given ElectricPotentialChangeRate to the current ElectricPotentialChangeRate.
func (a *ElectricPotentialChangeRate) Divide(other *ElectricPotentialChangeRate) *ElectricPotentialChangeRate {
	return &ElectricPotentialChangeRate{value: a.value / other.BaseValue()}
}