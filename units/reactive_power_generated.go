// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ReactivePowerUnits enumeration
type ReactivePowerUnits string

const (
    
        // 
        ReactivePowerVoltampereReactive ReactivePowerUnits = "VoltampereReactive"
        // 
        ReactivePowerKilovoltampereReactive ReactivePowerUnits = "KilovoltampereReactive"
        // 
        ReactivePowerMegavoltampereReactive ReactivePowerUnits = "MegavoltampereReactive"
        // 
        ReactivePowerGigavoltampereReactive ReactivePowerUnits = "GigavoltampereReactive"
)

// ReactivePowerDto represents an ReactivePower
type ReactivePowerDto struct {
	Value float64
	Unit  ReactivePowerUnits
}

// ReactivePowerDtoFactory struct to group related functions
type ReactivePowerDtoFactory struct{}

func (udf ReactivePowerDtoFactory) FromJSON(data []byte) (*ReactivePowerDto, error) {
	a := ReactivePowerDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ReactivePowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ReactivePower struct
type ReactivePower struct {
	value       float64
    
    voltamperes_reactiveLazy *float64 
    kilovoltamperes_reactiveLazy *float64 
    megavoltamperes_reactiveLazy *float64 
    gigavoltamperes_reactiveLazy *float64 
}

// ReactivePowerFactory struct to group related functions
type ReactivePowerFactory struct{}

func (uf ReactivePowerFactory) CreateReactivePower(value float64, unit ReactivePowerUnits) (*ReactivePower, error) {
	return newReactivePower(value, unit)
}

func (uf ReactivePowerFactory) FromDto(dto ReactivePowerDto) (*ReactivePower, error) {
	return newReactivePower(dto.Value, dto.Unit)
}

func (uf ReactivePowerFactory) FromDtoJSON(data []byte) (*ReactivePower, error) {
	unitDto, err := ReactivePowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ReactivePowerFactory{}.FromDto(*unitDto)
}


// FromVoltampereReactive creates a new ReactivePower instance from VoltampereReactive.
func (uf ReactivePowerFactory) FromVoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerVoltampereReactive)
}

// FromKilovoltampereReactive creates a new ReactivePower instance from KilovoltampereReactive.
func (uf ReactivePowerFactory) FromKilovoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerKilovoltampereReactive)
}

// FromMegavoltampereReactive creates a new ReactivePower instance from MegavoltampereReactive.
func (uf ReactivePowerFactory) FromMegavoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerMegavoltampereReactive)
}

// FromGigavoltampereReactive creates a new ReactivePower instance from GigavoltampereReactive.
func (uf ReactivePowerFactory) FromGigavoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerGigavoltampereReactive)
}




// newReactivePower creates a new ReactivePower.
func newReactivePower(value float64, fromUnit ReactivePowerUnits) (*ReactivePower, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ReactivePower{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReactivePower in VoltampereReactive.
func (a *ReactivePower) BaseValue() float64 {
	return a.value
}


// VoltampereReactive returns the value in VoltampereReactive.
func (a *ReactivePower) VoltamperesReactive() float64 {
	if a.voltamperes_reactiveLazy != nil {
		return *a.voltamperes_reactiveLazy
	}
	voltamperes_reactive := a.convertFromBase(ReactivePowerVoltampereReactive)
	a.voltamperes_reactiveLazy = &voltamperes_reactive
	return voltamperes_reactive
}

// KilovoltampereReactive returns the value in KilovoltampereReactive.
func (a *ReactivePower) KilovoltamperesReactive() float64 {
	if a.kilovoltamperes_reactiveLazy != nil {
		return *a.kilovoltamperes_reactiveLazy
	}
	kilovoltamperes_reactive := a.convertFromBase(ReactivePowerKilovoltampereReactive)
	a.kilovoltamperes_reactiveLazy = &kilovoltamperes_reactive
	return kilovoltamperes_reactive
}

// MegavoltampereReactive returns the value in MegavoltampereReactive.
func (a *ReactivePower) MegavoltamperesReactive() float64 {
	if a.megavoltamperes_reactiveLazy != nil {
		return *a.megavoltamperes_reactiveLazy
	}
	megavoltamperes_reactive := a.convertFromBase(ReactivePowerMegavoltampereReactive)
	a.megavoltamperes_reactiveLazy = &megavoltamperes_reactive
	return megavoltamperes_reactive
}

// GigavoltampereReactive returns the value in GigavoltampereReactive.
func (a *ReactivePower) GigavoltamperesReactive() float64 {
	if a.gigavoltamperes_reactiveLazy != nil {
		return *a.gigavoltamperes_reactiveLazy
	}
	gigavoltamperes_reactive := a.convertFromBase(ReactivePowerGigavoltampereReactive)
	a.gigavoltamperes_reactiveLazy = &gigavoltamperes_reactive
	return gigavoltamperes_reactive
}


// ToDto creates an ReactivePowerDto representation.
func (a *ReactivePower) ToDto(holdInUnit *ReactivePowerUnits) ReactivePowerDto {
	if holdInUnit == nil {
		defaultUnit := ReactivePowerVoltampereReactive // Default value
		holdInUnit = &defaultUnit
	}

	return ReactivePowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ReactivePowerDto representation.
func (a *ReactivePower) ToDtoJSON(holdInUnit *ReactivePowerUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ReactivePower to a specific unit value.
func (a *ReactivePower) Convert(toUnit ReactivePowerUnits) float64 {
	switch toUnit { 
    case ReactivePowerVoltampereReactive:
		return a.VoltamperesReactive()
    case ReactivePowerKilovoltampereReactive:
		return a.KilovoltamperesReactive()
    case ReactivePowerMegavoltampereReactive:
		return a.MegavoltamperesReactive()
    case ReactivePowerGigavoltampereReactive:
		return a.GigavoltamperesReactive()
	default:
		return 0
	}
}

func (a *ReactivePower) convertFromBase(toUnit ReactivePowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case ReactivePowerVoltampereReactive:
		return (value) 
	case ReactivePowerKilovoltampereReactive:
		return ((value) / 1000.0) 
	case ReactivePowerMegavoltampereReactive:
		return ((value) / 1000000.0) 
	case ReactivePowerGigavoltampereReactive:
		return ((value) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ReactivePower) convertToBase(value float64, fromUnit ReactivePowerUnits) float64 {
	switch fromUnit { 
	case ReactivePowerVoltampereReactive:
		return (value) 
	case ReactivePowerKilovoltampereReactive:
		return ((value) * 1000.0) 
	case ReactivePowerMegavoltampereReactive:
		return ((value) * 1000000.0) 
	case ReactivePowerGigavoltampereReactive:
		return ((value) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ReactivePower) String() string {
	return a.ToString(ReactivePowerVoltampereReactive, 2)
}

// ToString formats the ReactivePower to string.
// fractionalDigits -1 for not mention
func (a *ReactivePower) ToString(unit ReactivePowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ReactivePower) getUnitAbbreviation(unit ReactivePowerUnits) string {
	switch unit { 
	case ReactivePowerVoltampereReactive:
		return "var" 
	case ReactivePowerKilovoltampereReactive:
		return "kvar" 
	case ReactivePowerMegavoltampereReactive:
		return "Mvar" 
	case ReactivePowerGigavoltampereReactive:
		return "Gvar" 
	default:
		return ""
	}
}

// Check if the given ReactivePower are equals to the current ReactivePower
func (a *ReactivePower) Equals(other *ReactivePower) bool {
	return a.value == other.BaseValue()
}

// Check if the given ReactivePower are equals to the current ReactivePower
func (a *ReactivePower) CompareTo(other *ReactivePower) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ReactivePower to the current ReactivePower.
func (a *ReactivePower) Add(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value + other.BaseValue()}
}

// Subtract the given ReactivePower to the current ReactivePower.
func (a *ReactivePower) Subtract(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value - other.BaseValue()}
}

// Multiply the given ReactivePower to the current ReactivePower.
func (a *ReactivePower) Multiply(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value * other.BaseValue()}
}

// Divide the given ReactivePower to the current ReactivePower.
func (a *ReactivePower) Divide(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value / other.BaseValue()}
}