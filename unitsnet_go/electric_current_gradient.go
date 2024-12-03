package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentGradientUnits enumeration
type ElectricCurrentGradientUnits string

const (
    
        // 
        ElectricCurrentGradientAmperePerSecond ElectricCurrentGradientUnits = "AmperePerSecond"
        // 
        ElectricCurrentGradientAmperePerMinute ElectricCurrentGradientUnits = "AmperePerMinute"
        // 
        ElectricCurrentGradientAmperePerMillisecond ElectricCurrentGradientUnits = "AmperePerMillisecond"
        // 
        ElectricCurrentGradientAmperePerMicrosecond ElectricCurrentGradientUnits = "AmperePerMicrosecond"
        // 
        ElectricCurrentGradientAmperePerNanosecond ElectricCurrentGradientUnits = "AmperePerNanosecond"
        // 
        ElectricCurrentGradientMilliamperePerSecond ElectricCurrentGradientUnits = "MilliamperePerSecond"
        // 
        ElectricCurrentGradientMilliamperePerMinute ElectricCurrentGradientUnits = "MilliamperePerMinute"
)

// ElectricCurrentGradientDto represents an ElectricCurrentGradient
type ElectricCurrentGradientDto struct {
	Value float64
	Unit  ElectricCurrentGradientUnits
}

// ElectricCurrentGradientDtoFactory struct to group related functions
type ElectricCurrentGradientDtoFactory struct{}

func (udf ElectricCurrentGradientDtoFactory) FromJSON(data []byte) (*ElectricCurrentGradientDto, error) {
	a := ElectricCurrentGradientDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricCurrentGradientDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricCurrentGradient struct
type ElectricCurrentGradient struct {
	value       float64
    
    amperes_per_secondLazy *float64 
    amperes_per_minuteLazy *float64 
    amperes_per_millisecondLazy *float64 
    amperes_per_microsecondLazy *float64 
    amperes_per_nanosecondLazy *float64 
    milliamperes_per_secondLazy *float64 
    milliamperes_per_minuteLazy *float64 
}

// ElectricCurrentGradientFactory struct to group related functions
type ElectricCurrentGradientFactory struct{}

func (uf ElectricCurrentGradientFactory) CreateElectricCurrentGradient(value float64, unit ElectricCurrentGradientUnits) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, unit)
}

func (uf ElectricCurrentGradientFactory) FromDto(dto ElectricCurrentGradientDto) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(dto.Value, dto.Unit)
}

func (uf ElectricCurrentGradientFactory) FromDtoJSON(data []byte) (*ElectricCurrentGradient, error) {
	unitDto, err := ElectricCurrentGradientDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricCurrentGradientFactory{}.FromDto(*unitDto)
}


// FromAmperePerSecond creates a new ElectricCurrentGradient instance from AmperePerSecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerSecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerSecond)
}

// FromAmperePerMinute creates a new ElectricCurrentGradient instance from AmperePerMinute.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMinute(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMinute)
}

// FromAmperePerMillisecond creates a new ElectricCurrentGradient instance from AmperePerMillisecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMillisecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMillisecond)
}

// FromAmperePerMicrosecond creates a new ElectricCurrentGradient instance from AmperePerMicrosecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerMicrosecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerMicrosecond)
}

// FromAmperePerNanosecond creates a new ElectricCurrentGradient instance from AmperePerNanosecond.
func (uf ElectricCurrentGradientFactory) FromAmperesPerNanosecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientAmperePerNanosecond)
}

// FromMilliamperePerSecond creates a new ElectricCurrentGradient instance from MilliamperePerSecond.
func (uf ElectricCurrentGradientFactory) FromMilliamperesPerSecond(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientMilliamperePerSecond)
}

// FromMilliamperePerMinute creates a new ElectricCurrentGradient instance from MilliamperePerMinute.
func (uf ElectricCurrentGradientFactory) FromMilliamperesPerMinute(value float64) (*ElectricCurrentGradient, error) {
	return newElectricCurrentGradient(value, ElectricCurrentGradientMilliamperePerMinute)
}




// newElectricCurrentGradient creates a new ElectricCurrentGradient.
func newElectricCurrentGradient(value float64, fromUnit ElectricCurrentGradientUnits) (*ElectricCurrentGradient, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrentGradient{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrentGradient in AmperePerSecond.
func (a *ElectricCurrentGradient) BaseValue() float64 {
	return a.value
}


// AmperePerSecond returns the value in AmperePerSecond.
func (a *ElectricCurrentGradient) AmperesPerSecond() float64 {
	if a.amperes_per_secondLazy != nil {
		return *a.amperes_per_secondLazy
	}
	amperes_per_second := a.convertFromBase(ElectricCurrentGradientAmperePerSecond)
	a.amperes_per_secondLazy = &amperes_per_second
	return amperes_per_second
}

// AmperePerMinute returns the value in AmperePerMinute.
func (a *ElectricCurrentGradient) AmperesPerMinute() float64 {
	if a.amperes_per_minuteLazy != nil {
		return *a.amperes_per_minuteLazy
	}
	amperes_per_minute := a.convertFromBase(ElectricCurrentGradientAmperePerMinute)
	a.amperes_per_minuteLazy = &amperes_per_minute
	return amperes_per_minute
}

// AmperePerMillisecond returns the value in AmperePerMillisecond.
func (a *ElectricCurrentGradient) AmperesPerMillisecond() float64 {
	if a.amperes_per_millisecondLazy != nil {
		return *a.amperes_per_millisecondLazy
	}
	amperes_per_millisecond := a.convertFromBase(ElectricCurrentGradientAmperePerMillisecond)
	a.amperes_per_millisecondLazy = &amperes_per_millisecond
	return amperes_per_millisecond
}

// AmperePerMicrosecond returns the value in AmperePerMicrosecond.
func (a *ElectricCurrentGradient) AmperesPerMicrosecond() float64 {
	if a.amperes_per_microsecondLazy != nil {
		return *a.amperes_per_microsecondLazy
	}
	amperes_per_microsecond := a.convertFromBase(ElectricCurrentGradientAmperePerMicrosecond)
	a.amperes_per_microsecondLazy = &amperes_per_microsecond
	return amperes_per_microsecond
}

// AmperePerNanosecond returns the value in AmperePerNanosecond.
func (a *ElectricCurrentGradient) AmperesPerNanosecond() float64 {
	if a.amperes_per_nanosecondLazy != nil {
		return *a.amperes_per_nanosecondLazy
	}
	amperes_per_nanosecond := a.convertFromBase(ElectricCurrentGradientAmperePerNanosecond)
	a.amperes_per_nanosecondLazy = &amperes_per_nanosecond
	return amperes_per_nanosecond
}

// MilliamperePerSecond returns the value in MilliamperePerSecond.
func (a *ElectricCurrentGradient) MilliamperesPerSecond() float64 {
	if a.milliamperes_per_secondLazy != nil {
		return *a.milliamperes_per_secondLazy
	}
	milliamperes_per_second := a.convertFromBase(ElectricCurrentGradientMilliamperePerSecond)
	a.milliamperes_per_secondLazy = &milliamperes_per_second
	return milliamperes_per_second
}

// MilliamperePerMinute returns the value in MilliamperePerMinute.
func (a *ElectricCurrentGradient) MilliamperesPerMinute() float64 {
	if a.milliamperes_per_minuteLazy != nil {
		return *a.milliamperes_per_minuteLazy
	}
	milliamperes_per_minute := a.convertFromBase(ElectricCurrentGradientMilliamperePerMinute)
	a.milliamperes_per_minuteLazy = &milliamperes_per_minute
	return milliamperes_per_minute
}


// ToDto creates an ElectricCurrentGradientDto representation.
func (a *ElectricCurrentGradient) ToDto(holdInUnit *ElectricCurrentGradientUnits) ElectricCurrentGradientDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentGradientAmperePerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentGradientDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricCurrentGradientDto representation.
func (a *ElectricCurrentGradient) ToDtoJSON(holdInUnit *ElectricCurrentGradientUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricCurrentGradient to a specific unit value.
func (a *ElectricCurrentGradient) Convert(toUnit ElectricCurrentGradientUnits) float64 {
	switch toUnit { 
    case ElectricCurrentGradientAmperePerSecond:
		return a.AmperesPerSecond()
    case ElectricCurrentGradientAmperePerMinute:
		return a.AmperesPerMinute()
    case ElectricCurrentGradientAmperePerMillisecond:
		return a.AmperesPerMillisecond()
    case ElectricCurrentGradientAmperePerMicrosecond:
		return a.AmperesPerMicrosecond()
    case ElectricCurrentGradientAmperePerNanosecond:
		return a.AmperesPerNanosecond()
    case ElectricCurrentGradientMilliamperePerSecond:
		return a.MilliamperesPerSecond()
    case ElectricCurrentGradientMilliamperePerMinute:
		return a.MilliamperesPerMinute()
	default:
		return 0
	}
}

func (a *ElectricCurrentGradient) convertFromBase(toUnit ElectricCurrentGradientUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentGradientAmperePerSecond:
		return (value) 
	case ElectricCurrentGradientAmperePerMinute:
		return (value * 60) 
	case ElectricCurrentGradientAmperePerMillisecond:
		return (value / 1e3) 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return (value / 1e6) 
	case ElectricCurrentGradientAmperePerNanosecond:
		return (value / 1e9) 
	case ElectricCurrentGradientMilliamperePerSecond:
		return ((value) / 0.001) 
	case ElectricCurrentGradientMilliamperePerMinute:
		return ((value * 60) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrentGradient) convertToBase(value float64, fromUnit ElectricCurrentGradientUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentGradientAmperePerSecond:
		return (value) 
	case ElectricCurrentGradientAmperePerMinute:
		return (value / 60) 
	case ElectricCurrentGradientAmperePerMillisecond:
		return (value * 1e3) 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return (value * 1e6) 
	case ElectricCurrentGradientAmperePerNanosecond:
		return (value * 1e9) 
	case ElectricCurrentGradientMilliamperePerSecond:
		return ((value) * 0.001) 
	case ElectricCurrentGradientMilliamperePerMinute:
		return ((value / 60) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricCurrentGradient) String() string {
	return a.ToString(ElectricCurrentGradientAmperePerSecond, 2)
}

// ToString formats the ElectricCurrentGradient to string.
// fractionalDigits -1 for not mention
func (a *ElectricCurrentGradient) ToString(unit ElectricCurrentGradientUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCurrentGradient) getUnitAbbreviation(unit ElectricCurrentGradientUnits) string {
	switch unit { 
	case ElectricCurrentGradientAmperePerSecond:
		return "A/s" 
	case ElectricCurrentGradientAmperePerMinute:
		return "A/min" 
	case ElectricCurrentGradientAmperePerMillisecond:
		return "A/ms" 
	case ElectricCurrentGradientAmperePerMicrosecond:
		return "A/μs" 
	case ElectricCurrentGradientAmperePerNanosecond:
		return "A/ns" 
	case ElectricCurrentGradientMilliamperePerSecond:
		return "mA/s" 
	case ElectricCurrentGradientMilliamperePerMinute:
		return "mA/min" 
	default:
		return ""
	}
}

// Check if the given ElectricCurrentGradient are equals to the current ElectricCurrentGradient
func (a *ElectricCurrentGradient) Equals(other *ElectricCurrentGradient) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricCurrentGradient are equals to the current ElectricCurrentGradient
func (a *ElectricCurrentGradient) CompareTo(other *ElectricCurrentGradient) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricCurrentGradient to the current ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Add(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricCurrentGradient to the current ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Subtract(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricCurrentGradient to the current ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Multiply(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value * other.BaseValue()}
}

// Divide the given ElectricCurrentGradient to the current ElectricCurrentGradient.
func (a *ElectricCurrentGradient) Divide(other *ElectricCurrentGradient) *ElectricCurrentGradient {
	return &ElectricCurrentGradient{value: a.value / other.BaseValue()}
}