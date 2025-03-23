// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureChangeRateUnits defines various units of TemperatureChangeRate.
type TemperatureChangeRateUnits string

const (
    
        // 
        TemperatureChangeRateDegreeCelsiusPerSecond TemperatureChangeRateUnits = "DegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateDegreeCelsiusPerMinute TemperatureChangeRateUnits = "DegreeCelsiusPerMinute"
        // 
        TemperatureChangeRateDegreeKelvinPerMinute TemperatureChangeRateUnits = "DegreeKelvinPerMinute"
        // 
        TemperatureChangeRateDegreeFahrenheitPerMinute TemperatureChangeRateUnits = "DegreeFahrenheitPerMinute"
        // 
        TemperatureChangeRateDegreeFahrenheitPerSecond TemperatureChangeRateUnits = "DegreeFahrenheitPerSecond"
        // 
        TemperatureChangeRateDegreeKelvinPerSecond TemperatureChangeRateUnits = "DegreeKelvinPerSecond"
        // 
        TemperatureChangeRateDegreeCelsiusPerHour TemperatureChangeRateUnits = "DegreeCelsiusPerHour"
        // 
        TemperatureChangeRateDegreeKelvinPerHour TemperatureChangeRateUnits = "DegreeKelvinPerHour"
        // 
        TemperatureChangeRateDegreeFahrenheitPerHour TemperatureChangeRateUnits = "DegreeFahrenheitPerHour"
        // 
        TemperatureChangeRateNanodegreeCelsiusPerSecond TemperatureChangeRateUnits = "NanodegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateMicrodegreeCelsiusPerSecond TemperatureChangeRateUnits = "MicrodegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateMillidegreeCelsiusPerSecond TemperatureChangeRateUnits = "MillidegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateCentidegreeCelsiusPerSecond TemperatureChangeRateUnits = "CentidegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateDecidegreeCelsiusPerSecond TemperatureChangeRateUnits = "DecidegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateDecadegreeCelsiusPerSecond TemperatureChangeRateUnits = "DecadegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateHectodegreeCelsiusPerSecond TemperatureChangeRateUnits = "HectodegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateKilodegreeCelsiusPerSecond TemperatureChangeRateUnits = "KilodegreeCelsiusPerSecond"
)

var internalTemperatureChangeRateUnitsMap = map[TemperatureChangeRateUnits]bool{
	
	TemperatureChangeRateDegreeCelsiusPerSecond: true,
	TemperatureChangeRateDegreeCelsiusPerMinute: true,
	TemperatureChangeRateDegreeKelvinPerMinute: true,
	TemperatureChangeRateDegreeFahrenheitPerMinute: true,
	TemperatureChangeRateDegreeFahrenheitPerSecond: true,
	TemperatureChangeRateDegreeKelvinPerSecond: true,
	TemperatureChangeRateDegreeCelsiusPerHour: true,
	TemperatureChangeRateDegreeKelvinPerHour: true,
	TemperatureChangeRateDegreeFahrenheitPerHour: true,
	TemperatureChangeRateNanodegreeCelsiusPerSecond: true,
	TemperatureChangeRateMicrodegreeCelsiusPerSecond: true,
	TemperatureChangeRateMillidegreeCelsiusPerSecond: true,
	TemperatureChangeRateCentidegreeCelsiusPerSecond: true,
	TemperatureChangeRateDecidegreeCelsiusPerSecond: true,
	TemperatureChangeRateDecadegreeCelsiusPerSecond: true,
	TemperatureChangeRateHectodegreeCelsiusPerSecond: true,
	TemperatureChangeRateKilodegreeCelsiusPerSecond: true,
}

// TemperatureChangeRateDto represents a TemperatureChangeRate measurement with a numerical value and its corresponding unit.
type TemperatureChangeRateDto struct {
    // Value is the numerical representation of the TemperatureChangeRate.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the TemperatureChangeRate, as defined in the TemperatureChangeRateUnits enumeration.
	Unit  TemperatureChangeRateUnits `json:"unit" validate:"required,oneof=DegreeCelsiusPerSecond DegreeCelsiusPerMinute DegreeKelvinPerMinute DegreeFahrenheitPerMinute DegreeFahrenheitPerSecond DegreeKelvinPerSecond DegreeCelsiusPerHour DegreeKelvinPerHour DegreeFahrenheitPerHour NanodegreeCelsiusPerSecond MicrodegreeCelsiusPerSecond MillidegreeCelsiusPerSecond CentidegreeCelsiusPerSecond DecidegreeCelsiusPerSecond DecadegreeCelsiusPerSecond HectodegreeCelsiusPerSecond KilodegreeCelsiusPerSecond"`
}

// TemperatureChangeRateDtoFactory groups methods for creating and serializing TemperatureChangeRateDto objects.
type TemperatureChangeRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a TemperatureChangeRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf TemperatureChangeRateDtoFactory) FromJSON(data []byte) (*TemperatureChangeRateDto, error) {
	a := TemperatureChangeRateDto{}

    // Parse JSON into TemperatureChangeRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a TemperatureChangeRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a TemperatureChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// TemperatureChangeRate represents a measurement in a of TemperatureChangeRate.
//
// Temperature change rate is the ratio of the temperature change to the time during which the change occurred (value of temperature changes per unit time).
type TemperatureChangeRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    degrees_celsius_per_secondLazy *float64 
    degrees_celsius_per_minuteLazy *float64 
    degrees_kelvin_per_minuteLazy *float64 
    degrees_fahrenheit_per_minuteLazy *float64 
    degrees_fahrenheit_per_secondLazy *float64 
    degrees_kelvin_per_secondLazy *float64 
    degrees_celsius_per_hourLazy *float64 
    degrees_kelvin_per_hourLazy *float64 
    degrees_fahrenheit_per_hourLazy *float64 
    nanodegrees_celsius_per_secondLazy *float64 
    microdegrees_celsius_per_secondLazy *float64 
    millidegrees_celsius_per_secondLazy *float64 
    centidegrees_celsius_per_secondLazy *float64 
    decidegrees_celsius_per_secondLazy *float64 
    decadegrees_celsius_per_secondLazy *float64 
    hectodegrees_celsius_per_secondLazy *float64 
    kilodegrees_celsius_per_secondLazy *float64 
}

// TemperatureChangeRateFactory groups methods for creating TemperatureChangeRate instances.
type TemperatureChangeRateFactory struct{}

// CreateTemperatureChangeRate creates a new TemperatureChangeRate instance from the given value and unit.
func (uf TemperatureChangeRateFactory) CreateTemperatureChangeRate(value float64, unit TemperatureChangeRateUnits) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, unit)
}

// FromDto converts a TemperatureChangeRateDto to a TemperatureChangeRate instance.
func (uf TemperatureChangeRateFactory) FromDto(dto TemperatureChangeRateDto) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a TemperatureChangeRate instance.
func (uf TemperatureChangeRateFactory) FromDtoJSON(data []byte) (*TemperatureChangeRate, error) {
	unitDto, err := TemperatureChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TemperatureChangeRateDto from JSON: %w", err)
	}
	return TemperatureChangeRateFactory{}.FromDto(*unitDto)
}


// FromDegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in DegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeCelsiusPerSecond)
}

// FromDegreesCelsiusPerMinute creates a new TemperatureChangeRate instance from a value in DegreesCelsiusPerMinute.
func (uf TemperatureChangeRateFactory) FromDegreesCelsiusPerMinute(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeCelsiusPerMinute)
}

// FromDegreesKelvinPerMinute creates a new TemperatureChangeRate instance from a value in DegreesKelvinPerMinute.
func (uf TemperatureChangeRateFactory) FromDegreesKelvinPerMinute(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeKelvinPerMinute)
}

// FromDegreesFahrenheitPerMinute creates a new TemperatureChangeRate instance from a value in DegreesFahrenheitPerMinute.
func (uf TemperatureChangeRateFactory) FromDegreesFahrenheitPerMinute(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeFahrenheitPerMinute)
}

// FromDegreesFahrenheitPerSecond creates a new TemperatureChangeRate instance from a value in DegreesFahrenheitPerSecond.
func (uf TemperatureChangeRateFactory) FromDegreesFahrenheitPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeFahrenheitPerSecond)
}

// FromDegreesKelvinPerSecond creates a new TemperatureChangeRate instance from a value in DegreesKelvinPerSecond.
func (uf TemperatureChangeRateFactory) FromDegreesKelvinPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeKelvinPerSecond)
}

// FromDegreesCelsiusPerHour creates a new TemperatureChangeRate instance from a value in DegreesCelsiusPerHour.
func (uf TemperatureChangeRateFactory) FromDegreesCelsiusPerHour(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeCelsiusPerHour)
}

// FromDegreesKelvinPerHour creates a new TemperatureChangeRate instance from a value in DegreesKelvinPerHour.
func (uf TemperatureChangeRateFactory) FromDegreesKelvinPerHour(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeKelvinPerHour)
}

// FromDegreesFahrenheitPerHour creates a new TemperatureChangeRate instance from a value in DegreesFahrenheitPerHour.
func (uf TemperatureChangeRateFactory) FromDegreesFahrenheitPerHour(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeFahrenheitPerHour)
}

// FromNanodegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in NanodegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromNanodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateNanodegreeCelsiusPerSecond)
}

// FromMicrodegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in MicrodegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromMicrodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateMicrodegreeCelsiusPerSecond)
}

// FromMillidegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in MillidegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromMillidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateMillidegreeCelsiusPerSecond)
}

// FromCentidegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in CentidegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromCentidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateCentidegreeCelsiusPerSecond)
}

// FromDecidegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in DecidegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDecidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDecidegreeCelsiusPerSecond)
}

// FromDecadegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in DecadegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDecadegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDecadegreeCelsiusPerSecond)
}

// FromHectodegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in HectodegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromHectodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateHectodegreeCelsiusPerSecond)
}

// FromKilodegreesCelsiusPerSecond creates a new TemperatureChangeRate instance from a value in KilodegreesCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromKilodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateKilodegreeCelsiusPerSecond)
}


// newTemperatureChangeRate creates a new TemperatureChangeRate.
func newTemperatureChangeRate(value float64, fromUnit TemperatureChangeRateUnits) (*TemperatureChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalTemperatureChangeRateUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in TemperatureChangeRateUnits", fromUnit)
	}
	a := &TemperatureChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureChangeRate in DegreeCelsiusPerSecond unit (the base/default quantity).
func (a *TemperatureChangeRate) BaseValue() float64 {
	return a.value
}


// DegreesCelsiusPerSecond returns the TemperatureChangeRate value in DegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) DegreesCelsiusPerSecond() float64 {
	if a.degrees_celsius_per_secondLazy != nil {
		return *a.degrees_celsius_per_secondLazy
	}
	degrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDegreeCelsiusPerSecond)
	a.degrees_celsius_per_secondLazy = &degrees_celsius_per_second
	return degrees_celsius_per_second
}

// DegreesCelsiusPerMinute returns the TemperatureChangeRate value in DegreesCelsiusPerMinute.
//
// 
func (a *TemperatureChangeRate) DegreesCelsiusPerMinute() float64 {
	if a.degrees_celsius_per_minuteLazy != nil {
		return *a.degrees_celsius_per_minuteLazy
	}
	degrees_celsius_per_minute := a.convertFromBase(TemperatureChangeRateDegreeCelsiusPerMinute)
	a.degrees_celsius_per_minuteLazy = &degrees_celsius_per_minute
	return degrees_celsius_per_minute
}

// DegreesKelvinPerMinute returns the TemperatureChangeRate value in DegreesKelvinPerMinute.
//
// 
func (a *TemperatureChangeRate) DegreesKelvinPerMinute() float64 {
	if a.degrees_kelvin_per_minuteLazy != nil {
		return *a.degrees_kelvin_per_minuteLazy
	}
	degrees_kelvin_per_minute := a.convertFromBase(TemperatureChangeRateDegreeKelvinPerMinute)
	a.degrees_kelvin_per_minuteLazy = &degrees_kelvin_per_minute
	return degrees_kelvin_per_minute
}

// DegreesFahrenheitPerMinute returns the TemperatureChangeRate value in DegreesFahrenheitPerMinute.
//
// 
func (a *TemperatureChangeRate) DegreesFahrenheitPerMinute() float64 {
	if a.degrees_fahrenheit_per_minuteLazy != nil {
		return *a.degrees_fahrenheit_per_minuteLazy
	}
	degrees_fahrenheit_per_minute := a.convertFromBase(TemperatureChangeRateDegreeFahrenheitPerMinute)
	a.degrees_fahrenheit_per_minuteLazy = &degrees_fahrenheit_per_minute
	return degrees_fahrenheit_per_minute
}

// DegreesFahrenheitPerSecond returns the TemperatureChangeRate value in DegreesFahrenheitPerSecond.
//
// 
func (a *TemperatureChangeRate) DegreesFahrenheitPerSecond() float64 {
	if a.degrees_fahrenheit_per_secondLazy != nil {
		return *a.degrees_fahrenheit_per_secondLazy
	}
	degrees_fahrenheit_per_second := a.convertFromBase(TemperatureChangeRateDegreeFahrenheitPerSecond)
	a.degrees_fahrenheit_per_secondLazy = &degrees_fahrenheit_per_second
	return degrees_fahrenheit_per_second
}

// DegreesKelvinPerSecond returns the TemperatureChangeRate value in DegreesKelvinPerSecond.
//
// 
func (a *TemperatureChangeRate) DegreesKelvinPerSecond() float64 {
	if a.degrees_kelvin_per_secondLazy != nil {
		return *a.degrees_kelvin_per_secondLazy
	}
	degrees_kelvin_per_second := a.convertFromBase(TemperatureChangeRateDegreeKelvinPerSecond)
	a.degrees_kelvin_per_secondLazy = &degrees_kelvin_per_second
	return degrees_kelvin_per_second
}

// DegreesCelsiusPerHour returns the TemperatureChangeRate value in DegreesCelsiusPerHour.
//
// 
func (a *TemperatureChangeRate) DegreesCelsiusPerHour() float64 {
	if a.degrees_celsius_per_hourLazy != nil {
		return *a.degrees_celsius_per_hourLazy
	}
	degrees_celsius_per_hour := a.convertFromBase(TemperatureChangeRateDegreeCelsiusPerHour)
	a.degrees_celsius_per_hourLazy = &degrees_celsius_per_hour
	return degrees_celsius_per_hour
}

// DegreesKelvinPerHour returns the TemperatureChangeRate value in DegreesKelvinPerHour.
//
// 
func (a *TemperatureChangeRate) DegreesKelvinPerHour() float64 {
	if a.degrees_kelvin_per_hourLazy != nil {
		return *a.degrees_kelvin_per_hourLazy
	}
	degrees_kelvin_per_hour := a.convertFromBase(TemperatureChangeRateDegreeKelvinPerHour)
	a.degrees_kelvin_per_hourLazy = &degrees_kelvin_per_hour
	return degrees_kelvin_per_hour
}

// DegreesFahrenheitPerHour returns the TemperatureChangeRate value in DegreesFahrenheitPerHour.
//
// 
func (a *TemperatureChangeRate) DegreesFahrenheitPerHour() float64 {
	if a.degrees_fahrenheit_per_hourLazy != nil {
		return *a.degrees_fahrenheit_per_hourLazy
	}
	degrees_fahrenheit_per_hour := a.convertFromBase(TemperatureChangeRateDegreeFahrenheitPerHour)
	a.degrees_fahrenheit_per_hourLazy = &degrees_fahrenheit_per_hour
	return degrees_fahrenheit_per_hour
}

// NanodegreesCelsiusPerSecond returns the TemperatureChangeRate value in NanodegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) NanodegreesCelsiusPerSecond() float64 {
	if a.nanodegrees_celsius_per_secondLazy != nil {
		return *a.nanodegrees_celsius_per_secondLazy
	}
	nanodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateNanodegreeCelsiusPerSecond)
	a.nanodegrees_celsius_per_secondLazy = &nanodegrees_celsius_per_second
	return nanodegrees_celsius_per_second
}

// MicrodegreesCelsiusPerSecond returns the TemperatureChangeRate value in MicrodegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) MicrodegreesCelsiusPerSecond() float64 {
	if a.microdegrees_celsius_per_secondLazy != nil {
		return *a.microdegrees_celsius_per_secondLazy
	}
	microdegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateMicrodegreeCelsiusPerSecond)
	a.microdegrees_celsius_per_secondLazy = &microdegrees_celsius_per_second
	return microdegrees_celsius_per_second
}

// MillidegreesCelsiusPerSecond returns the TemperatureChangeRate value in MillidegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) MillidegreesCelsiusPerSecond() float64 {
	if a.millidegrees_celsius_per_secondLazy != nil {
		return *a.millidegrees_celsius_per_secondLazy
	}
	millidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateMillidegreeCelsiusPerSecond)
	a.millidegrees_celsius_per_secondLazy = &millidegrees_celsius_per_second
	return millidegrees_celsius_per_second
}

// CentidegreesCelsiusPerSecond returns the TemperatureChangeRate value in CentidegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) CentidegreesCelsiusPerSecond() float64 {
	if a.centidegrees_celsius_per_secondLazy != nil {
		return *a.centidegrees_celsius_per_secondLazy
	}
	centidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateCentidegreeCelsiusPerSecond)
	a.centidegrees_celsius_per_secondLazy = &centidegrees_celsius_per_second
	return centidegrees_celsius_per_second
}

// DecidegreesCelsiusPerSecond returns the TemperatureChangeRate value in DecidegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) DecidegreesCelsiusPerSecond() float64 {
	if a.decidegrees_celsius_per_secondLazy != nil {
		return *a.decidegrees_celsius_per_secondLazy
	}
	decidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDecidegreeCelsiusPerSecond)
	a.decidegrees_celsius_per_secondLazy = &decidegrees_celsius_per_second
	return decidegrees_celsius_per_second
}

// DecadegreesCelsiusPerSecond returns the TemperatureChangeRate value in DecadegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) DecadegreesCelsiusPerSecond() float64 {
	if a.decadegrees_celsius_per_secondLazy != nil {
		return *a.decadegrees_celsius_per_secondLazy
	}
	decadegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDecadegreeCelsiusPerSecond)
	a.decadegrees_celsius_per_secondLazy = &decadegrees_celsius_per_second
	return decadegrees_celsius_per_second
}

// HectodegreesCelsiusPerSecond returns the TemperatureChangeRate value in HectodegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) HectodegreesCelsiusPerSecond() float64 {
	if a.hectodegrees_celsius_per_secondLazy != nil {
		return *a.hectodegrees_celsius_per_secondLazy
	}
	hectodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateHectodegreeCelsiusPerSecond)
	a.hectodegrees_celsius_per_secondLazy = &hectodegrees_celsius_per_second
	return hectodegrees_celsius_per_second
}

// KilodegreesCelsiusPerSecond returns the TemperatureChangeRate value in KilodegreesCelsiusPerSecond.
//
// 
func (a *TemperatureChangeRate) KilodegreesCelsiusPerSecond() float64 {
	if a.kilodegrees_celsius_per_secondLazy != nil {
		return *a.kilodegrees_celsius_per_secondLazy
	}
	kilodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateKilodegreeCelsiusPerSecond)
	a.kilodegrees_celsius_per_secondLazy = &kilodegrees_celsius_per_second
	return kilodegrees_celsius_per_second
}


// ToDto creates a TemperatureChangeRateDto representation from the TemperatureChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DegreeCelsiusPerSecond by default.
func (a *TemperatureChangeRate) ToDto(holdInUnit *TemperatureChangeRateUnits) TemperatureChangeRateDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureChangeRateDegreeCelsiusPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureChangeRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the TemperatureChangeRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DegreeCelsiusPerSecond by default.
func (a *TemperatureChangeRate) ToDtoJSON(holdInUnit *TemperatureChangeRateUnits) ([]byte, error) {
	// Convert to TemperatureChangeRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a TemperatureChangeRate to a specific unit value.
// The function uses the provided unit type (TemperatureChangeRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *TemperatureChangeRate) Convert(toUnit TemperatureChangeRateUnits) float64 {
	switch toUnit { 
    case TemperatureChangeRateDegreeCelsiusPerSecond:
		return a.DegreesCelsiusPerSecond()
    case TemperatureChangeRateDegreeCelsiusPerMinute:
		return a.DegreesCelsiusPerMinute()
    case TemperatureChangeRateDegreeKelvinPerMinute:
		return a.DegreesKelvinPerMinute()
    case TemperatureChangeRateDegreeFahrenheitPerMinute:
		return a.DegreesFahrenheitPerMinute()
    case TemperatureChangeRateDegreeFahrenheitPerSecond:
		return a.DegreesFahrenheitPerSecond()
    case TemperatureChangeRateDegreeKelvinPerSecond:
		return a.DegreesKelvinPerSecond()
    case TemperatureChangeRateDegreeCelsiusPerHour:
		return a.DegreesCelsiusPerHour()
    case TemperatureChangeRateDegreeKelvinPerHour:
		return a.DegreesKelvinPerHour()
    case TemperatureChangeRateDegreeFahrenheitPerHour:
		return a.DegreesFahrenheitPerHour()
    case TemperatureChangeRateNanodegreeCelsiusPerSecond:
		return a.NanodegreesCelsiusPerSecond()
    case TemperatureChangeRateMicrodegreeCelsiusPerSecond:
		return a.MicrodegreesCelsiusPerSecond()
    case TemperatureChangeRateMillidegreeCelsiusPerSecond:
		return a.MillidegreesCelsiusPerSecond()
    case TemperatureChangeRateCentidegreeCelsiusPerSecond:
		return a.CentidegreesCelsiusPerSecond()
    case TemperatureChangeRateDecidegreeCelsiusPerSecond:
		return a.DecidegreesCelsiusPerSecond()
    case TemperatureChangeRateDecadegreeCelsiusPerSecond:
		return a.DecadegreesCelsiusPerSecond()
    case TemperatureChangeRateHectodegreeCelsiusPerSecond:
		return a.HectodegreesCelsiusPerSecond()
    case TemperatureChangeRateKilodegreeCelsiusPerSecond:
		return a.KilodegreesCelsiusPerSecond()
	default:
		return math.NaN()
	}
}

func (a *TemperatureChangeRate) convertFromBase(toUnit TemperatureChangeRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureChangeRateDegreeCelsiusPerSecond:
		return (value) 
	case TemperatureChangeRateDegreeCelsiusPerMinute:
		return (value * 60) 
	case TemperatureChangeRateDegreeKelvinPerMinute:
		return (value * 60) 
	case TemperatureChangeRateDegreeFahrenheitPerMinute:
		return (value * 9 / 5 * 60) 
	case TemperatureChangeRateDegreeFahrenheitPerSecond:
		return (value * 9 / 5) 
	case TemperatureChangeRateDegreeKelvinPerSecond:
		return (value) 
	case TemperatureChangeRateDegreeCelsiusPerHour:
		return (value * 3600) 
	case TemperatureChangeRateDegreeKelvinPerHour:
		return (value * 3600) 
	case TemperatureChangeRateDegreeFahrenheitPerHour:
		return (value * 9 / 5 * 3600) 
	case TemperatureChangeRateNanodegreeCelsiusPerSecond:
		return ((value) / 1e-09) 
	case TemperatureChangeRateMicrodegreeCelsiusPerSecond:
		return ((value) / 1e-06) 
	case TemperatureChangeRateMillidegreeCelsiusPerSecond:
		return ((value) / 0.001) 
	case TemperatureChangeRateCentidegreeCelsiusPerSecond:
		return ((value) / 0.01) 
	case TemperatureChangeRateDecidegreeCelsiusPerSecond:
		return ((value) / 0.1) 
	case TemperatureChangeRateDecadegreeCelsiusPerSecond:
		return ((value) / 10.0) 
	case TemperatureChangeRateHectodegreeCelsiusPerSecond:
		return ((value) / 100.0) 
	case TemperatureChangeRateKilodegreeCelsiusPerSecond:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *TemperatureChangeRate) convertToBase(value float64, fromUnit TemperatureChangeRateUnits) float64 {
	switch fromUnit { 
	case TemperatureChangeRateDegreeCelsiusPerSecond:
		return (value) 
	case TemperatureChangeRateDegreeCelsiusPerMinute:
		return (value / 60) 
	case TemperatureChangeRateDegreeKelvinPerMinute:
		return (value / 60) 
	case TemperatureChangeRateDegreeFahrenheitPerMinute:
		return (value * 5 / 9 / 60) 
	case TemperatureChangeRateDegreeFahrenheitPerSecond:
		return (value * 5 / 9) 
	case TemperatureChangeRateDegreeKelvinPerSecond:
		return (value) 
	case TemperatureChangeRateDegreeCelsiusPerHour:
		return (value / 3600) 
	case TemperatureChangeRateDegreeKelvinPerHour:
		return (value / 3600) 
	case TemperatureChangeRateDegreeFahrenheitPerHour:
		return (value * 5 / 9 / 3600) 
	case TemperatureChangeRateNanodegreeCelsiusPerSecond:
		return ((value) * 1e-09) 
	case TemperatureChangeRateMicrodegreeCelsiusPerSecond:
		return ((value) * 1e-06) 
	case TemperatureChangeRateMillidegreeCelsiusPerSecond:
		return ((value) * 0.001) 
	case TemperatureChangeRateCentidegreeCelsiusPerSecond:
		return ((value) * 0.01) 
	case TemperatureChangeRateDecidegreeCelsiusPerSecond:
		return ((value) * 0.1) 
	case TemperatureChangeRateDecadegreeCelsiusPerSecond:
		return ((value) * 10.0) 
	case TemperatureChangeRateHectodegreeCelsiusPerSecond:
		return ((value) * 100.0) 
	case TemperatureChangeRateKilodegreeCelsiusPerSecond:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the TemperatureChangeRate in the default unit (DegreeCelsiusPerSecond),
// formatted to two decimal places.
func (a TemperatureChangeRate) String() string {
	return a.ToString(TemperatureChangeRateDegreeCelsiusPerSecond, 2)
}

// ToString formats the TemperatureChangeRate value as a string with the specified unit and fractional digits.
// It converts the TemperatureChangeRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the TemperatureChangeRate value will be converted (e.g., DegreeCelsiusPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the TemperatureChangeRate with the unit abbreviation.
func (a *TemperatureChangeRate) ToString(unit TemperatureChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetTemperatureChangeRateAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetTemperatureChangeRateAbbreviation(unit))
}

// Equals checks if the given TemperatureChangeRate is equal to the current TemperatureChangeRate.
//
// Parameters:
//    other: The TemperatureChangeRate to compare against.
//
// Returns:
//    bool: Returns true if both TemperatureChangeRate are equal, false otherwise.
func (a *TemperatureChangeRate) Equals(other *TemperatureChangeRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current TemperatureChangeRate with another TemperatureChangeRate.
// It returns -1 if the current TemperatureChangeRate is less than the other TemperatureChangeRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The TemperatureChangeRate to compare against.
//
// Returns:
//    int: -1 if the current TemperatureChangeRate is less, 1 if greater, and 0 if equal.
func (a *TemperatureChangeRate) CompareTo(other *TemperatureChangeRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given TemperatureChangeRate to the current TemperatureChangeRate and returns the result.
// The result is a new TemperatureChangeRate instance with the sum of the values.
//
// Parameters:
//    other: The TemperatureChangeRate to add to the current TemperatureChangeRate.
//
// Returns:
//    *TemperatureChangeRate: A new TemperatureChangeRate instance representing the sum of both TemperatureChangeRate.
func (a *TemperatureChangeRate) Add(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given TemperatureChangeRate from the current TemperatureChangeRate and returns the result.
// The result is a new TemperatureChangeRate instance with the difference of the values.
//
// Parameters:
//    other: The TemperatureChangeRate to subtract from the current TemperatureChangeRate.
//
// Returns:
//    *TemperatureChangeRate: A new TemperatureChangeRate instance representing the difference of both TemperatureChangeRate.
func (a *TemperatureChangeRate) Subtract(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current TemperatureChangeRate by the given TemperatureChangeRate and returns the result.
// The result is a new TemperatureChangeRate instance with the product of the values.
//
// Parameters:
//    other: The TemperatureChangeRate to multiply with the current TemperatureChangeRate.
//
// Returns:
//    *TemperatureChangeRate: A new TemperatureChangeRate instance representing the product of both TemperatureChangeRate.
func (a *TemperatureChangeRate) Multiply(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value * other.BaseValue()}
}

// Divide divides the current TemperatureChangeRate by the given TemperatureChangeRate and returns the result.
// The result is a new TemperatureChangeRate instance with the quotient of the values.
//
// Parameters:
//    other: The TemperatureChangeRate to divide the current TemperatureChangeRate by.
//
// Returns:
//    *TemperatureChangeRate: A new TemperatureChangeRate instance representing the quotient of both TemperatureChangeRate.
func (a *TemperatureChangeRate) Divide(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value / other.BaseValue()}
}

// GetTemperatureChangeRateAbbreviation gets the unit abbreviation.
func GetTemperatureChangeRateAbbreviation(unit TemperatureChangeRateUnits) string {
	switch unit { 
	case TemperatureChangeRateDegreeCelsiusPerSecond:
		return "°C/s" 
	case TemperatureChangeRateDegreeCelsiusPerMinute:
		return "°C/min" 
	case TemperatureChangeRateDegreeKelvinPerMinute:
		return "K/min" 
	case TemperatureChangeRateDegreeFahrenheitPerMinute:
		return "°F/min" 
	case TemperatureChangeRateDegreeFahrenheitPerSecond:
		return "°F/s" 
	case TemperatureChangeRateDegreeKelvinPerSecond:
		return "K/s" 
	case TemperatureChangeRateDegreeCelsiusPerHour:
		return "°C/h" 
	case TemperatureChangeRateDegreeKelvinPerHour:
		return "K/h" 
	case TemperatureChangeRateDegreeFahrenheitPerHour:
		return "°F/h" 
	case TemperatureChangeRateNanodegreeCelsiusPerSecond:
		return "n°C/s" 
	case TemperatureChangeRateMicrodegreeCelsiusPerSecond:
		return "μ°C/s" 
	case TemperatureChangeRateMillidegreeCelsiusPerSecond:
		return "m°C/s" 
	case TemperatureChangeRateCentidegreeCelsiusPerSecond:
		return "c°C/s" 
	case TemperatureChangeRateDecidegreeCelsiusPerSecond:
		return "d°C/s" 
	case TemperatureChangeRateDecadegreeCelsiusPerSecond:
		return "da°C/s" 
	case TemperatureChangeRateHectodegreeCelsiusPerSecond:
		return "h°C/s" 
	case TemperatureChangeRateKilodegreeCelsiusPerSecond:
		return "k°C/s" 
	default:
		return ""
	}
}