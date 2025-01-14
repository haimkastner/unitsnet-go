// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// FrequencyUnits defines various units of Frequency.
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

// FrequencyDto represents a Frequency measurement with a numerical value and its corresponding unit.
type FrequencyDto struct {
    // Value is the numerical representation of the Frequency.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Frequency, as defined in the FrequencyUnits enumeration.
	Unit  FrequencyUnits `json:"unit"`
}

// FrequencyDtoFactory groups methods for creating and serializing FrequencyDto objects.
type FrequencyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a FrequencyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf FrequencyDtoFactory) FromJSON(data []byte) (*FrequencyDto, error) {
	a := FrequencyDto{}

    // Parse JSON into FrequencyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a FrequencyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a FrequencyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Frequency represents a measurement in a of Frequency.
//
// The number of occurrences of a repeating event per unit time.
type Frequency struct {
	// value is the base measurement stored internally.
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

// FrequencyFactory groups methods for creating Frequency instances.
type FrequencyFactory struct{}

// CreateFrequency creates a new Frequency instance from the given value and unit.
func (uf FrequencyFactory) CreateFrequency(value float64, unit FrequencyUnits) (*Frequency, error) {
	return newFrequency(value, unit)
}

// FromDto converts a FrequencyDto to a Frequency instance.
func (uf FrequencyFactory) FromDto(dto FrequencyDto) (*Frequency, error) {
	return newFrequency(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Frequency instance.
func (uf FrequencyFactory) FromDtoJSON(data []byte) (*Frequency, error) {
	unitDto, err := FrequencyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FrequencyDto from JSON: %w", err)
	}
	return FrequencyFactory{}.FromDto(*unitDto)
}


// FromHertz creates a new Frequency instance from a value in Hertz.
func (uf FrequencyFactory) FromHertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyHertz)
}

// FromRadiansPerSecond creates a new Frequency instance from a value in RadiansPerSecond.
func (uf FrequencyFactory) FromRadiansPerSecond(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyRadianPerSecond)
}

// FromCyclesPerMinute creates a new Frequency instance from a value in CyclesPerMinute.
func (uf FrequencyFactory) FromCyclesPerMinute(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyCyclePerMinute)
}

// FromCyclesPerHour creates a new Frequency instance from a value in CyclesPerHour.
func (uf FrequencyFactory) FromCyclesPerHour(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyCyclePerHour)
}

// FromBeatsPerMinute creates a new Frequency instance from a value in BeatsPerMinute.
func (uf FrequencyFactory) FromBeatsPerMinute(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyBeatPerMinute)
}

// FromPerSecond creates a new Frequency instance from a value in PerSecond.
func (uf FrequencyFactory) FromPerSecond(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyPerSecond)
}

// FromBUnits creates a new Frequency instance from a value in BUnits.
func (uf FrequencyFactory) FromBUnits(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyBUnit)
}

// FromMicrohertz creates a new Frequency instance from a value in Microhertz.
func (uf FrequencyFactory) FromMicrohertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMicrohertz)
}

// FromMillihertz creates a new Frequency instance from a value in Millihertz.
func (uf FrequencyFactory) FromMillihertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMillihertz)
}

// FromKilohertz creates a new Frequency instance from a value in Kilohertz.
func (uf FrequencyFactory) FromKilohertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyKilohertz)
}

// FromMegahertz creates a new Frequency instance from a value in Megahertz.
func (uf FrequencyFactory) FromMegahertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyMegahertz)
}

// FromGigahertz creates a new Frequency instance from a value in Gigahertz.
func (uf FrequencyFactory) FromGigahertz(value float64) (*Frequency, error) {
	return newFrequency(value, FrequencyGigahertz)
}

// FromTerahertz creates a new Frequency instance from a value in Terahertz.
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

// BaseValue returns the base value of Frequency in Hertz unit (the base/default quantity).
func (a *Frequency) BaseValue() float64 {
	return a.value
}


// Hertz returns the Frequency value in Hertz.
//
// 
func (a *Frequency) Hertz() float64 {
	if a.hertzLazy != nil {
		return *a.hertzLazy
	}
	hertz := a.convertFromBase(FrequencyHertz)
	a.hertzLazy = &hertz
	return hertz
}

// RadiansPerSecond returns the Frequency value in RadiansPerSecond.
//
// 
func (a *Frequency) RadiansPerSecond() float64 {
	if a.radians_per_secondLazy != nil {
		return *a.radians_per_secondLazy
	}
	radians_per_second := a.convertFromBase(FrequencyRadianPerSecond)
	a.radians_per_secondLazy = &radians_per_second
	return radians_per_second
}

// CyclesPerMinute returns the Frequency value in CyclesPerMinute.
//
// 
func (a *Frequency) CyclesPerMinute() float64 {
	if a.cycles_per_minuteLazy != nil {
		return *a.cycles_per_minuteLazy
	}
	cycles_per_minute := a.convertFromBase(FrequencyCyclePerMinute)
	a.cycles_per_minuteLazy = &cycles_per_minute
	return cycles_per_minute
}

// CyclesPerHour returns the Frequency value in CyclesPerHour.
//
// 
func (a *Frequency) CyclesPerHour() float64 {
	if a.cycles_per_hourLazy != nil {
		return *a.cycles_per_hourLazy
	}
	cycles_per_hour := a.convertFromBase(FrequencyCyclePerHour)
	a.cycles_per_hourLazy = &cycles_per_hour
	return cycles_per_hour
}

// BeatsPerMinute returns the Frequency value in BeatsPerMinute.
//
// 
func (a *Frequency) BeatsPerMinute() float64 {
	if a.beats_per_minuteLazy != nil {
		return *a.beats_per_minuteLazy
	}
	beats_per_minute := a.convertFromBase(FrequencyBeatPerMinute)
	a.beats_per_minuteLazy = &beats_per_minute
	return beats_per_minute
}

// PerSecond returns the Frequency value in PerSecond.
//
// 
func (a *Frequency) PerSecond() float64 {
	if a.per_secondLazy != nil {
		return *a.per_secondLazy
	}
	per_second := a.convertFromBase(FrequencyPerSecond)
	a.per_secondLazy = &per_second
	return per_second
}

// BUnits returns the Frequency value in BUnits.
//
// 
func (a *Frequency) BUnits() float64 {
	if a.b_unitsLazy != nil {
		return *a.b_unitsLazy
	}
	b_units := a.convertFromBase(FrequencyBUnit)
	a.b_unitsLazy = &b_units
	return b_units
}

// Microhertz returns the Frequency value in Microhertz.
//
// 
func (a *Frequency) Microhertz() float64 {
	if a.microhertzLazy != nil {
		return *a.microhertzLazy
	}
	microhertz := a.convertFromBase(FrequencyMicrohertz)
	a.microhertzLazy = &microhertz
	return microhertz
}

// Millihertz returns the Frequency value in Millihertz.
//
// 
func (a *Frequency) Millihertz() float64 {
	if a.millihertzLazy != nil {
		return *a.millihertzLazy
	}
	millihertz := a.convertFromBase(FrequencyMillihertz)
	a.millihertzLazy = &millihertz
	return millihertz
}

// Kilohertz returns the Frequency value in Kilohertz.
//
// 
func (a *Frequency) Kilohertz() float64 {
	if a.kilohertzLazy != nil {
		return *a.kilohertzLazy
	}
	kilohertz := a.convertFromBase(FrequencyKilohertz)
	a.kilohertzLazy = &kilohertz
	return kilohertz
}

// Megahertz returns the Frequency value in Megahertz.
//
// 
func (a *Frequency) Megahertz() float64 {
	if a.megahertzLazy != nil {
		return *a.megahertzLazy
	}
	megahertz := a.convertFromBase(FrequencyMegahertz)
	a.megahertzLazy = &megahertz
	return megahertz
}

// Gigahertz returns the Frequency value in Gigahertz.
//
// 
func (a *Frequency) Gigahertz() float64 {
	if a.gigahertzLazy != nil {
		return *a.gigahertzLazy
	}
	gigahertz := a.convertFromBase(FrequencyGigahertz)
	a.gigahertzLazy = &gigahertz
	return gigahertz
}

// Terahertz returns the Frequency value in Terahertz.
//
// 
func (a *Frequency) Terahertz() float64 {
	if a.terahertzLazy != nil {
		return *a.terahertzLazy
	}
	terahertz := a.convertFromBase(FrequencyTerahertz)
	a.terahertzLazy = &terahertz
	return terahertz
}


// ToDto creates a FrequencyDto representation from the Frequency instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Hertz by default.
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

// ToDtoJSON creates a JSON representation of the Frequency instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Hertz by default.
func (a *Frequency) ToDtoJSON(holdInUnit *FrequencyUnits) ([]byte, error) {
	// Convert to FrequencyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Frequency to a specific unit value.
// The function uses the provided unit type (FrequencyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Frequency in the default unit (Hertz),
// formatted to two decimal places.
func (a Frequency) String() string {
	return a.ToString(FrequencyHertz, 2)
}

// ToString formats the Frequency value as a string with the specified unit and fractional digits.
// It converts the Frequency to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Frequency value will be converted (e.g., Hertz))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Frequency with the unit abbreviation.
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

// Equals checks if the given Frequency is equal to the current Frequency.
//
// Parameters:
//    other: The Frequency to compare against.
//
// Returns:
//    bool: Returns true if both Frequency are equal, false otherwise.
func (a *Frequency) Equals(other *Frequency) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Frequency with another Frequency.
// It returns -1 if the current Frequency is less than the other Frequency, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Frequency to compare against.
//
// Returns:
//    int: -1 if the current Frequency is less, 1 if greater, and 0 if equal.
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

// Add adds the given Frequency to the current Frequency and returns the result.
// The result is a new Frequency instance with the sum of the values.
//
// Parameters:
//    other: The Frequency to add to the current Frequency.
//
// Returns:
//    *Frequency: A new Frequency instance representing the sum of both Frequency.
func (a *Frequency) Add(other *Frequency) *Frequency {
	return &Frequency{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Frequency from the current Frequency and returns the result.
// The result is a new Frequency instance with the difference of the values.
//
// Parameters:
//    other: The Frequency to subtract from the current Frequency.
//
// Returns:
//    *Frequency: A new Frequency instance representing the difference of both Frequency.
func (a *Frequency) Subtract(other *Frequency) *Frequency {
	return &Frequency{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Frequency by the given Frequency and returns the result.
// The result is a new Frequency instance with the product of the values.
//
// Parameters:
//    other: The Frequency to multiply with the current Frequency.
//
// Returns:
//    *Frequency: A new Frequency instance representing the product of both Frequency.
func (a *Frequency) Multiply(other *Frequency) *Frequency {
	return &Frequency{value: a.value * other.BaseValue()}
}

// Divide divides the current Frequency by the given Frequency and returns the result.
// The result is a new Frequency instance with the quotient of the values.
//
// Parameters:
//    other: The Frequency to divide the current Frequency by.
//
// Returns:
//    *Frequency: A new Frequency instance representing the quotient of both Frequency.
func (a *Frequency) Divide(other *Frequency) *Frequency {
	return &Frequency{value: a.value / other.BaseValue()}
}