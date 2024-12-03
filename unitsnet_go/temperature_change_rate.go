package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureChangeRateUnits enumeration
type TemperatureChangeRateUnits string

const (
    
        // 
        TemperatureChangeRateDegreeCelsiusPerSecond TemperatureChangeRateUnits = "DegreeCelsiusPerSecond"
        // 
        TemperatureChangeRateDegreeCelsiusPerMinute TemperatureChangeRateUnits = "DegreeCelsiusPerMinute"
        // 
        TemperatureChangeRateDegreeKelvinPerMinute TemperatureChangeRateUnits = "DegreeKelvinPerMinute"
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

// TemperatureChangeRateDto represents an TemperatureChangeRate
type TemperatureChangeRateDto struct {
	Value float64
	Unit  TemperatureChangeRateUnits
}

// TemperatureChangeRateDtoFactory struct to group related functions
type TemperatureChangeRateDtoFactory struct{}

func (udf TemperatureChangeRateDtoFactory) FromJSON(data []byte) (*TemperatureChangeRateDto, error) {
	a := TemperatureChangeRateDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TemperatureChangeRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// TemperatureChangeRate struct
type TemperatureChangeRate struct {
	value       float64
    
    degrees_celsius_per_secondLazy *float64 
    degrees_celsius_per_minuteLazy *float64 
    degrees_kelvin_per_minuteLazy *float64 
    nanodegrees_celsius_per_secondLazy *float64 
    microdegrees_celsius_per_secondLazy *float64 
    millidegrees_celsius_per_secondLazy *float64 
    centidegrees_celsius_per_secondLazy *float64 
    decidegrees_celsius_per_secondLazy *float64 
    decadegrees_celsius_per_secondLazy *float64 
    hectodegrees_celsius_per_secondLazy *float64 
    kilodegrees_celsius_per_secondLazy *float64 
}

// TemperatureChangeRateFactory struct to group related functions
type TemperatureChangeRateFactory struct{}

func (uf TemperatureChangeRateFactory) CreateTemperatureChangeRate(value float64, unit TemperatureChangeRateUnits) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, unit)
}

func (uf TemperatureChangeRateFactory) FromDto(dto TemperatureChangeRateDto) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(dto.Value, dto.Unit)
}

func (uf TemperatureChangeRateFactory) FromDtoJSON(data []byte) (*TemperatureChangeRate, error) {
	unitDto, err := TemperatureChangeRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TemperatureChangeRateFactory{}.FromDto(*unitDto)
}


// FromDegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from DegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeCelsiusPerSecond)
}

// FromDegreeCelsiusPerMinute creates a new TemperatureChangeRate instance from DegreeCelsiusPerMinute.
func (uf TemperatureChangeRateFactory) FromDegreesCelsiusPerMinute(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeCelsiusPerMinute)
}

// FromDegreeKelvinPerMinute creates a new TemperatureChangeRate instance from DegreeKelvinPerMinute.
func (uf TemperatureChangeRateFactory) FromDegreesKelvinPerMinute(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDegreeKelvinPerMinute)
}

// FromNanodegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from NanodegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromNanodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateNanodegreeCelsiusPerSecond)
}

// FromMicrodegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from MicrodegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromMicrodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateMicrodegreeCelsiusPerSecond)
}

// FromMillidegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from MillidegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromMillidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateMillidegreeCelsiusPerSecond)
}

// FromCentidegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from CentidegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromCentidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateCentidegreeCelsiusPerSecond)
}

// FromDecidegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from DecidegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDecidegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDecidegreeCelsiusPerSecond)
}

// FromDecadegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from DecadegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromDecadegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateDecadegreeCelsiusPerSecond)
}

// FromHectodegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from HectodegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromHectodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateHectodegreeCelsiusPerSecond)
}

// FromKilodegreeCelsiusPerSecond creates a new TemperatureChangeRate instance from KilodegreeCelsiusPerSecond.
func (uf TemperatureChangeRateFactory) FromKilodegreesCelsiusPerSecond(value float64) (*TemperatureChangeRate, error) {
	return newTemperatureChangeRate(value, TemperatureChangeRateKilodegreeCelsiusPerSecond)
}




// newTemperatureChangeRate creates a new TemperatureChangeRate.
func newTemperatureChangeRate(value float64, fromUnit TemperatureChangeRateUnits) (*TemperatureChangeRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TemperatureChangeRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureChangeRate in DegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) BaseValue() float64 {
	return a.value
}


// DegreeCelsiusPerSecond returns the value in DegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) DegreesCelsiusPerSecond() float64 {
	if a.degrees_celsius_per_secondLazy != nil {
		return *a.degrees_celsius_per_secondLazy
	}
	degrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDegreeCelsiusPerSecond)
	a.degrees_celsius_per_secondLazy = &degrees_celsius_per_second
	return degrees_celsius_per_second
}

// DegreeCelsiusPerMinute returns the value in DegreeCelsiusPerMinute.
func (a *TemperatureChangeRate) DegreesCelsiusPerMinute() float64 {
	if a.degrees_celsius_per_minuteLazy != nil {
		return *a.degrees_celsius_per_minuteLazy
	}
	degrees_celsius_per_minute := a.convertFromBase(TemperatureChangeRateDegreeCelsiusPerMinute)
	a.degrees_celsius_per_minuteLazy = &degrees_celsius_per_minute
	return degrees_celsius_per_minute
}

// DegreeKelvinPerMinute returns the value in DegreeKelvinPerMinute.
func (a *TemperatureChangeRate) DegreesKelvinPerMinute() float64 {
	if a.degrees_kelvin_per_minuteLazy != nil {
		return *a.degrees_kelvin_per_minuteLazy
	}
	degrees_kelvin_per_minute := a.convertFromBase(TemperatureChangeRateDegreeKelvinPerMinute)
	a.degrees_kelvin_per_minuteLazy = &degrees_kelvin_per_minute
	return degrees_kelvin_per_minute
}

// NanodegreeCelsiusPerSecond returns the value in NanodegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) NanodegreesCelsiusPerSecond() float64 {
	if a.nanodegrees_celsius_per_secondLazy != nil {
		return *a.nanodegrees_celsius_per_secondLazy
	}
	nanodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateNanodegreeCelsiusPerSecond)
	a.nanodegrees_celsius_per_secondLazy = &nanodegrees_celsius_per_second
	return nanodegrees_celsius_per_second
}

// MicrodegreeCelsiusPerSecond returns the value in MicrodegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) MicrodegreesCelsiusPerSecond() float64 {
	if a.microdegrees_celsius_per_secondLazy != nil {
		return *a.microdegrees_celsius_per_secondLazy
	}
	microdegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateMicrodegreeCelsiusPerSecond)
	a.microdegrees_celsius_per_secondLazy = &microdegrees_celsius_per_second
	return microdegrees_celsius_per_second
}

// MillidegreeCelsiusPerSecond returns the value in MillidegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) MillidegreesCelsiusPerSecond() float64 {
	if a.millidegrees_celsius_per_secondLazy != nil {
		return *a.millidegrees_celsius_per_secondLazy
	}
	millidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateMillidegreeCelsiusPerSecond)
	a.millidegrees_celsius_per_secondLazy = &millidegrees_celsius_per_second
	return millidegrees_celsius_per_second
}

// CentidegreeCelsiusPerSecond returns the value in CentidegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) CentidegreesCelsiusPerSecond() float64 {
	if a.centidegrees_celsius_per_secondLazy != nil {
		return *a.centidegrees_celsius_per_secondLazy
	}
	centidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateCentidegreeCelsiusPerSecond)
	a.centidegrees_celsius_per_secondLazy = &centidegrees_celsius_per_second
	return centidegrees_celsius_per_second
}

// DecidegreeCelsiusPerSecond returns the value in DecidegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) DecidegreesCelsiusPerSecond() float64 {
	if a.decidegrees_celsius_per_secondLazy != nil {
		return *a.decidegrees_celsius_per_secondLazy
	}
	decidegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDecidegreeCelsiusPerSecond)
	a.decidegrees_celsius_per_secondLazy = &decidegrees_celsius_per_second
	return decidegrees_celsius_per_second
}

// DecadegreeCelsiusPerSecond returns the value in DecadegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) DecadegreesCelsiusPerSecond() float64 {
	if a.decadegrees_celsius_per_secondLazy != nil {
		return *a.decadegrees_celsius_per_secondLazy
	}
	decadegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateDecadegreeCelsiusPerSecond)
	a.decadegrees_celsius_per_secondLazy = &decadegrees_celsius_per_second
	return decadegrees_celsius_per_second
}

// HectodegreeCelsiusPerSecond returns the value in HectodegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) HectodegreesCelsiusPerSecond() float64 {
	if a.hectodegrees_celsius_per_secondLazy != nil {
		return *a.hectodegrees_celsius_per_secondLazy
	}
	hectodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateHectodegreeCelsiusPerSecond)
	a.hectodegrees_celsius_per_secondLazy = &hectodegrees_celsius_per_second
	return hectodegrees_celsius_per_second
}

// KilodegreeCelsiusPerSecond returns the value in KilodegreeCelsiusPerSecond.
func (a *TemperatureChangeRate) KilodegreesCelsiusPerSecond() float64 {
	if a.kilodegrees_celsius_per_secondLazy != nil {
		return *a.kilodegrees_celsius_per_secondLazy
	}
	kilodegrees_celsius_per_second := a.convertFromBase(TemperatureChangeRateKilodegreeCelsiusPerSecond)
	a.kilodegrees_celsius_per_secondLazy = &kilodegrees_celsius_per_second
	return kilodegrees_celsius_per_second
}


// ToDto creates an TemperatureChangeRateDto representation.
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

// ToDtoJSON creates an TemperatureChangeRateDto representation.
func (a *TemperatureChangeRate) ToDtoJSON(holdInUnit *TemperatureChangeRateUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts TemperatureChangeRate to a specific unit value.
func (a *TemperatureChangeRate) Convert(toUnit TemperatureChangeRateUnits) float64 {
	switch toUnit { 
    case TemperatureChangeRateDegreeCelsiusPerSecond:
		return a.DegreesCelsiusPerSecond()
    case TemperatureChangeRateDegreeCelsiusPerMinute:
		return a.DegreesCelsiusPerMinute()
    case TemperatureChangeRateDegreeKelvinPerMinute:
		return a.DegreesKelvinPerMinute()
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
		return 0
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
		return ((value + 273.15) * 60) 
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
		return ((value / 60) - 273.15) 
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

// Implement the String() method for AngleDto
func (a TemperatureChangeRate) String() string {
	return a.ToString(TemperatureChangeRateDegreeCelsiusPerSecond, 2)
}

// ToString formats the TemperatureChangeRate to string.
// fractionalDigits -1 for not mention
func (a *TemperatureChangeRate) ToString(unit TemperatureChangeRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TemperatureChangeRate) getUnitAbbreviation(unit TemperatureChangeRateUnits) string {
	switch unit { 
	case TemperatureChangeRateDegreeCelsiusPerSecond:
		return "°C/s" 
	case TemperatureChangeRateDegreeCelsiusPerMinute:
		return "°C/min" 
	case TemperatureChangeRateDegreeKelvinPerMinute:
		return "K/min" 
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

// Check if the given TemperatureChangeRate are equals to the current TemperatureChangeRate
func (a *TemperatureChangeRate) Equals(other *TemperatureChangeRate) bool {
	return a.value == other.BaseValue()
}

// Check if the given TemperatureChangeRate are equals to the current TemperatureChangeRate
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

// Add the given TemperatureChangeRate to the current TemperatureChangeRate.
func (a *TemperatureChangeRate) Add(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value + other.BaseValue()}
}

// Subtract the given TemperatureChangeRate to the current TemperatureChangeRate.
func (a *TemperatureChangeRate) Subtract(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value - other.BaseValue()}
}

// Multiply the given TemperatureChangeRate to the current TemperatureChangeRate.
func (a *TemperatureChangeRate) Multiply(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value * other.BaseValue()}
}

// Divide the given TemperatureChangeRate to the current TemperatureChangeRate.
func (a *TemperatureChangeRate) Divide(other *TemperatureChangeRate) *TemperatureChangeRate {
	return &TemperatureChangeRate{value: a.value / other.BaseValue()}
}