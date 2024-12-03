// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// FrequencyUnits enumeration
type FrequencyUnits string

const (
    
        // 
        FrequencyHertz FrequencyUnits = "Hertz"
        // 
        FrequencyRadianPerSecond FrequencyUnits = "RadianPerSecond"
        // 
        FrequencyCyclePerMinute FrequencyUnits = "CyclePerMinute"
        // 
        FrequencyCyclePerHour FrequencyUnits = "CyclePerHour"
        // 
        FrequencyBeatPerMinute FrequencyUnits = "BeatPerMinute"
        // 
        FrequencyPerSecond FrequencyUnits = "PerSecond"
        // 
        FrequencyBUnit FrequencyUnits = "BUnit"
        // 
        FrequencyMicrohertz FrequencyUnits = "Microhertz"
        // 
        FrequencyMillihertz FrequencyUnits = "Millihertz"
        // 
        FrequencyKilohertz FrequencyUnits = "Kilohertz"
        // 
        FrequencyMegahertz FrequencyUnits = "Megahertz"
        // 
        FrequencyGigahertz FrequencyUnits = "Gigahertz"
        // 
        FrequencyTerahertz FrequencyUnits = "Terahertz"
)

// FrequencyDto represents an Frequency
type FrequencyDto struct {
	Value float64
	Unit  FrequencyUnits
}

// FrequencyDtoFactory struct to group related functions
type FrequencyDtoFactory struct{}

func (udf FrequencyDtoFactory) FromJSON(data []byte) (*FrequencyDto, error) {
	a := FrequencyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a FrequencyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Frequency struct
type Frequency struct {
	value       float64
    
    hertzLazy *float64 
    radians_per_secondLazy *float64 
    cycles_per_minuteLazy *float64 
    cycles_per_hourLazy *float64 
    beats_per_minuteLazy *float64 
    per_secondLazy *float64 
    b_unitsLazy *float64 
    microhertzLazy *float64 
    millihertzLazy *float64 
    kilohertzLazy *float64 
    megahertzLazy *float64 
    gigahertzLazy *float64 
    terahertzLazy *float64 
}

// FrequencyFactory struct to group related functions
type FrequencyFactory struct{}

func (uf FrequencyFactory) CreateFrequency(value float64, unit FrequencyUnits) (*Frequency, error) {
	return newFrequency(value, unit)
}

func (uf FrequencyFactory) FromDto(dto FrequencyDto) (*Frequency, error) {
	return newFrequency(dto.Value, dto.Unit)
}

func (uf FrequencyFactory) FromDtoJSON(data []byte) (*Frequency, error) {
	unitDto, err := FrequencyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return FrequencyFactory{}.FromDto(*unitDto)
}


// FromHertz creates a new Frequency instance from Hertz.
func (uf FrequencyFactory) FromHertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyHertz)
}

// FromRadianPerSecond creates a new Frequency instance from RadianPerSecond.
func (uf FrequencyFactory) FromRadiansPerSecond(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyRadianPerSecond)
}

// FromCyclePerMinute creates a new Frequency instance from CyclePerMinute.
func (uf FrequencyFactory) FromCyclesPerMinute(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyCyclePerMinute)
}

// FromCyclePerHour creates a new Frequency instance from CyclePerHour.
func (uf FrequencyFactory) FromCyclesPerHour(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyCyclePerHour)
}

// FromBeatPerMinute creates a new Frequency instance from BeatPerMinute.
func (uf FrequencyFactory) FromBeatsPerMinute(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyBeatPerMinute)
}

// FromPerSecond creates a new Frequency instance from PerSecond.
func (uf FrequencyFactory) FromPerSecond(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyPerSecond)
}

// FromBUnit creates a new Frequency instance from BUnit.
func (uf FrequencyFactory) FromBUnits(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyBUnit)
}

// FromMicrohertz creates a new Frequency instance from Microhertz.
func (uf FrequencyFactory) FromMicrohertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMicrohertz)
}

// FromMillihertz creates a new Frequency instance from Millihertz.
func (uf FrequencyFactory) FromMillihertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMillihertz)
}

// FromKilohertz creates a new Frequency instance from Kilohertz.
func (uf FrequencyFactory) FromKilohertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyKilohertz)
}

// FromMegahertz creates a new Frequency instance from Megahertz.
func (uf FrequencyFactory) FromMegahertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMegahertz)
}

// FromGigahertz creates a new Frequency instance from Gigahertz.
func (uf FrequencyFactory) FromGigahertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyGigahertz)
}

// FromTerahertz creates a new Frequency instance from Terahertz.
func (uf FrequencyFactory) FromTerahertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyTerahertz)
}




// newFrequency creates a new Frequency.
func newFrequency(value float64, fromUnit FrequencyUnits) (*Frequency, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Frequency{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Frequency in Hertz.
func (a *Frequency) BaseValue() float64 {
	return a.value
}


// Hertz returns the value in Hertz.
func (a *Frequency) Hertz() float64 {
	if a.hertzLazy != nil {
		return *a.hertzLazy
	}
	hertz := a.convertFromBase(FrequencyHertz)
	a.hertzLazy = &hertz
	return hertz
}

// RadianPerSecond returns the value in RadianPerSecond.
func (a *Frequency) RadiansPerSecond() float64 {
	if a.radians_per_secondLazy != nil {
		return *a.radians_per_secondLazy
	}
	radians_per_second := a.convertFromBase(FrequencyRadianPerSecond)
	a.radians_per_secondLazy = &radians_per_second
	return radians_per_second
}

// CyclePerMinute returns the value in CyclePerMinute.
func (a *Frequency) CyclesPerMinute() float64 {
	if a.cycles_per_minuteLazy != nil {
		return *a.cycles_per_minuteLazy
	}
	cycles_per_minute := a.convertFromBase(FrequencyCyclePerMinute)
	a.cycles_per_minuteLazy = &cycles_per_minute
	return cycles_per_minute
}

// CyclePerHour returns the value in CyclePerHour.
func (a *Frequency) CyclesPerHour() float64 {
	if a.cycles_per_hourLazy != nil {
		return *a.cycles_per_hourLazy
	}
	cycles_per_hour := a.convertFromBase(FrequencyCyclePerHour)
	a.cycles_per_hourLazy = &cycles_per_hour
	return cycles_per_hour
}

// BeatPerMinute returns the value in BeatPerMinute.
func (a *Frequency) BeatsPerMinute() float64 {
	if a.beats_per_minuteLazy != nil {
		return *a.beats_per_minuteLazy
	}
	beats_per_minute := a.convertFromBase(FrequencyBeatPerMinute)
	a.beats_per_minuteLazy = &beats_per_minute
	return beats_per_minute
}

// PerSecond returns the value in PerSecond.
func (a *Frequency) PerSecond() float64 {
	if a.per_secondLazy != nil {
		return *a.per_secondLazy
	}
	per_second := a.convertFromBase(FrequencyPerSecond)
	a.per_secondLazy = &per_second
	return per_second
}

// BUnit returns the value in BUnit.
func (a *Frequency) BUnits() float64 {
	if a.b_unitsLazy != nil {
		return *a.b_unitsLazy
	}
	b_units := a.convertFromBase(FrequencyBUnit)
	a.b_unitsLazy = &b_units
	return b_units
}

// Microhertz returns the value in Microhertz.
func (a *Frequency) Microhertz() float64 {
	if a.microhertzLazy != nil {
		return *a.microhertzLazy
	}
	microhertz := a.convertFromBase(FrequencyMicrohertz)
	a.microhertzLazy = &microhertz
	return microhertz
}

// Millihertz returns the value in Millihertz.
func (a *Frequency) Millihertz() float64 {
	if a.millihertzLazy != nil {
		return *a.millihertzLazy
	}
	millihertz := a.convertFromBase(FrequencyMillihertz)
	a.millihertzLazy = &millihertz
	return millihertz
}

// Kilohertz returns the value in Kilohertz.
func (a *Frequency) Kilohertz() float64 {
	if a.kilohertzLazy != nil {
		return *a.kilohertzLazy
	}
	kilohertz := a.convertFromBase(FrequencyKilohertz)
	a.kilohertzLazy = &kilohertz
	return kilohertz
}

// Megahertz returns the value in Megahertz.
func (a *Frequency) Megahertz() float64 {
	if a.megahertzLazy != nil {
		return *a.megahertzLazy
	}
	megahertz := a.convertFromBase(FrequencyMegahertz)
	a.megahertzLazy = &megahertz
	return megahertz
}

// Gigahertz returns the value in Gigahertz.
func (a *Frequency) Gigahertz() float64 {
	if a.gigahertzLazy != nil {
		return *a.gigahertzLazy
	}
	gigahertz := a.convertFromBase(FrequencyGigahertz)
	a.gigahertzLazy = &gigahertz
	return gigahertz
}

// Terahertz returns the value in Terahertz.
func (a *Frequency) Terahertz() float64 {
	if a.terahertzLazy != nil {
		return *a.terahertzLazy
	}
	terahertz := a.convertFromBase(FrequencyTerahertz)
	a.terahertzLazy = &terahertz
	return terahertz
}


// ToDto creates an FrequencyDto representation.
func (a *Frequency) ToDto(holdInUnit *FrequencyUnits) FrequencyDto {
	if holdInUnit == nil {
		defaultUnit := FrequencyHertz // Default value
		holdInUnit = &defaultUnit
	}

	return FrequencyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an FrequencyDto representation.
func (a *Frequency) ToDtoJSON(holdInUnit *FrequencyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Frequency to a specific unit value.
func (a *Frequency) Convert(toUnit FrequencyUnits) float64 {
	switch toUnit { 
    case FrequencyHertz:
		return a.Hertz()
    case FrequencyRadianPerSecond:
		return a.RadiansPerSecond()
    case FrequencyCyclePerMinute:
		return a.CyclesPerMinute()
    case FrequencyCyclePerHour:
		return a.CyclesPerHour()
    case FrequencyBeatPerMinute:
		return a.BeatsPerMinute()
    case FrequencyPerSecond:
		return a.PerSecond()
    case FrequencyBUnit:
		return a.BUnits()
    case FrequencyMicrohertz:
		return a.Microhertz()
    case FrequencyMillihertz:
		return a.Millihertz()
    case FrequencyKilohertz:
		return a.Kilohertz()
    case FrequencyMegahertz:
		return a.Megahertz()
    case FrequencyGigahertz:
		return a.Gigahertz()
    case FrequencyTerahertz:
		return a.Terahertz()
	default:
		return 0
	}
}

func (a *Frequency) convertFromBase(toUnit FrequencyUnits) float64 {
    value := a.value
	switch toUnit { 
	case FrequencyHertz:
		return (value) 
	case FrequencyRadianPerSecond:
		return (value * 6.2831853072) 
	case FrequencyCyclePerMinute:
		return (value * 60) 
	case FrequencyCyclePerHour:
		return (value * 3600) 
	case FrequencyBeatPerMinute:
		return (value * 60) 
	case FrequencyPerSecond:
		return (value) 
	case FrequencyBUnit:
		return (value * value * 1e-3) 
	case FrequencyMicrohertz:
		return ((value) / 1e-06) 
	case FrequencyMillihertz:
		return ((value) / 0.001) 
	case FrequencyKilohertz:
		return ((value) / 1000.0) 
	case FrequencyMegahertz:
		return ((value) / 1000000.0) 
	case FrequencyGigahertz:
		return ((value) / 1000000000.0) 
	case FrequencyTerahertz:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Frequency) convertToBase(value float64, fromUnit FrequencyUnits) float64 {
	switch fromUnit { 
	case FrequencyHertz:
		return (value) 
	case FrequencyRadianPerSecond:
		return (value / 6.2831853072) 
	case FrequencyCyclePerMinute:
		return (value / 60) 
	case FrequencyCyclePerHour:
		return (value / 3600) 
	case FrequencyBeatPerMinute:
		return (value / 60) 
	case FrequencyPerSecond:
		return (value) 
	case FrequencyBUnit:
		return (math.Sqrt(value * 1e3)) 
	case FrequencyMicrohertz:
		return ((value) * 1e-06) 
	case FrequencyMillihertz:
		return ((value) * 0.001) 
	case FrequencyKilohertz:
		return ((value) * 1000.0) 
	case FrequencyMegahertz:
		return ((value) * 1000000.0) 
	case FrequencyGigahertz:
		return ((value) * 1000000000.0) 
	case FrequencyTerahertz:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Frequency) String() string {
	return a.ToString(FrequencyHertz, 2)
}

// ToString formats the Frequency to string.
// fractionalDigits -1 for not mention
func (a *Frequency) ToString(unit FrequencyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Frequency) getUnitAbbreviation(unit FrequencyUnits) string {
	switch unit { 
	case FrequencyHertz:
		return "Hz" 
	case FrequencyRadianPerSecond:
		return "rad/s" 
	case FrequencyCyclePerMinute:
		return "cpm" 
	case FrequencyCyclePerHour:
		return "cph" 
	case FrequencyBeatPerMinute:
		return "bpm" 
	case FrequencyPerSecond:
		return "s⁻¹" 
	case FrequencyBUnit:
		return "B Units" 
	case FrequencyMicrohertz:
		return "μHz" 
	case FrequencyMillihertz:
		return "mHz" 
	case FrequencyKilohertz:
		return "kHz" 
	case FrequencyMegahertz:
		return "MHz" 
	case FrequencyGigahertz:
		return "GHz" 
	case FrequencyTerahertz:
		return "THz" 
	default:
		return ""
	}
}

// Check if the given Frequency are equals to the current Frequency
func (a *Frequency) Equals(other *Frequency) bool {
	return a.value == other.BaseValue()
}

// Check if the given Frequency are equals to the current Frequency
func (a *Frequency) CompareTo(other *Frequency) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Frequency to the current Frequency.
func (a *Frequency) Add(other *Frequency) *Frequency {
	return &Frequency{value: a.value + other.BaseValue()}
}

// Subtract the given Frequency to the current Frequency.
func (a *Frequency) Subtract(other *Frequency) *Frequency {
	return &Frequency{value: a.value - other.BaseValue()}
}

// Multiply the given Frequency to the current Frequency.
func (a *Frequency) Multiply(other *Frequency) *Frequency {
	return &Frequency{value: a.value * other.BaseValue()}
}

// Divide the given Frequency to the current Frequency.
func (a *Frequency) Divide(other *Frequency) *Frequency {
	return &Frequency{value: a.value / other.BaseValue()}
}