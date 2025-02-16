// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ReactivePowerUnits defines various units of ReactivePower.
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

// ReactivePowerDto represents a ReactivePower measurement with a numerical value and its corresponding unit.
type ReactivePowerDto struct {
    // Value is the numerical representation of the ReactivePower.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ReactivePower, as defined in the ReactivePowerUnits enumeration.
	Unit  ReactivePowerUnits `json:"unit"`
}

// ReactivePowerDtoFactory groups methods for creating and serializing ReactivePowerDto objects.
type ReactivePowerDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ReactivePowerDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ReactivePowerDtoFactory) FromJSON(data []byte) (*ReactivePowerDto, error) {
	a := ReactivePowerDto{}

    // Parse JSON into ReactivePowerDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ReactivePowerDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ReactivePowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ReactivePower represents a measurement in a of ReactivePower.
//
// Volt-ampere reactive (var) is a unit by which reactive power is expressed in an AC electric power system. Reactive power exists in an AC circuit when the current and voltage are not in phase.
type ReactivePower struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltamperes_reactiveLazy *float64 
    kilovoltamperes_reactiveLazy *float64 
    megavoltamperes_reactiveLazy *float64 
    gigavoltamperes_reactiveLazy *float64 
}

// ReactivePowerFactory groups methods for creating ReactivePower instances.
type ReactivePowerFactory struct{}

// CreateReactivePower creates a new ReactivePower instance from the given value and unit.
func (uf ReactivePowerFactory) CreateReactivePower(value float64, unit ReactivePowerUnits) (*ReactivePower, error) {
	return newReactivePower(value, unit)
}

// FromDto converts a ReactivePowerDto to a ReactivePower instance.
func (uf ReactivePowerFactory) FromDto(dto ReactivePowerDto) (*ReactivePower, error) {
	return newReactivePower(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ReactivePower instance.
func (uf ReactivePowerFactory) FromDtoJSON(data []byte) (*ReactivePower, error) {
	unitDto, err := ReactivePowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ReactivePowerDto from JSON: %w", err)
	}
	return ReactivePowerFactory{}.FromDto(*unitDto)
}


// FromVoltamperesReactive creates a new ReactivePower instance from a value in VoltamperesReactive.
func (uf ReactivePowerFactory) FromVoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerVoltampereReactive)
}

// FromKilovoltamperesReactive creates a new ReactivePower instance from a value in KilovoltamperesReactive.
func (uf ReactivePowerFactory) FromKilovoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerKilovoltampereReactive)
}

// FromMegavoltamperesReactive creates a new ReactivePower instance from a value in MegavoltamperesReactive.
func (uf ReactivePowerFactory) FromMegavoltamperesReactive(value float64) (*ReactivePower, error) {
	return newReactivePower(value, ReactivePowerMegavoltampereReactive)
}

// FromGigavoltamperesReactive creates a new ReactivePower instance from a value in GigavoltamperesReactive.
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

// BaseValue returns the base value of ReactivePower in VoltampereReactive unit (the base/default quantity).
func (a *ReactivePower) BaseValue() float64 {
	return a.value
}


// VoltamperesReactive returns the ReactivePower value in VoltamperesReactive.
//
// 
func (a *ReactivePower) VoltamperesReactive() float64 {
	if a.voltamperes_reactiveLazy != nil {
		return *a.voltamperes_reactiveLazy
	}
	voltamperes_reactive := a.convertFromBase(ReactivePowerVoltampereReactive)
	a.voltamperes_reactiveLazy = &voltamperes_reactive
	return voltamperes_reactive
}

// KilovoltamperesReactive returns the ReactivePower value in KilovoltamperesReactive.
//
// 
func (a *ReactivePower) KilovoltamperesReactive() float64 {
	if a.kilovoltamperes_reactiveLazy != nil {
		return *a.kilovoltamperes_reactiveLazy
	}
	kilovoltamperes_reactive := a.convertFromBase(ReactivePowerKilovoltampereReactive)
	a.kilovoltamperes_reactiveLazy = &kilovoltamperes_reactive
	return kilovoltamperes_reactive
}

// MegavoltamperesReactive returns the ReactivePower value in MegavoltamperesReactive.
//
// 
func (a *ReactivePower) MegavoltamperesReactive() float64 {
	if a.megavoltamperes_reactiveLazy != nil {
		return *a.megavoltamperes_reactiveLazy
	}
	megavoltamperes_reactive := a.convertFromBase(ReactivePowerMegavoltampereReactive)
	a.megavoltamperes_reactiveLazy = &megavoltamperes_reactive
	return megavoltamperes_reactive
}

// GigavoltamperesReactive returns the ReactivePower value in GigavoltamperesReactive.
//
// 
func (a *ReactivePower) GigavoltamperesReactive() float64 {
	if a.gigavoltamperes_reactiveLazy != nil {
		return *a.gigavoltamperes_reactiveLazy
	}
	gigavoltamperes_reactive := a.convertFromBase(ReactivePowerGigavoltampereReactive)
	a.gigavoltamperes_reactiveLazy = &gigavoltamperes_reactive
	return gigavoltamperes_reactive
}


// ToDto creates a ReactivePowerDto representation from the ReactivePower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactive by default.
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

// ToDtoJSON creates a JSON representation of the ReactivePower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactive by default.
func (a *ReactivePower) ToDtoJSON(holdInUnit *ReactivePowerUnits) ([]byte, error) {
	// Convert to ReactivePowerDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ReactivePower to a specific unit value.
// The function uses the provided unit type (ReactivePowerUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the ReactivePower in the default unit (VoltampereReactive),
// formatted to two decimal places.
func (a ReactivePower) String() string {
	return a.ToString(ReactivePowerVoltampereReactive, 2)
}

// ToString formats the ReactivePower value as a string with the specified unit and fractional digits.
// It converts the ReactivePower to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ReactivePower value will be converted (e.g., VoltampereReactive))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ReactivePower with the unit abbreviation.
func (a *ReactivePower) ToString(unit ReactivePowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetReactivePowerAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetReactivePowerAbbreviation(unit))
}

// Equals checks if the given ReactivePower is equal to the current ReactivePower.
//
// Parameters:
//    other: The ReactivePower to compare against.
//
// Returns:
//    bool: Returns true if both ReactivePower are equal, false otherwise.
func (a *ReactivePower) Equals(other *ReactivePower) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ReactivePower with another ReactivePower.
// It returns -1 if the current ReactivePower is less than the other ReactivePower, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ReactivePower to compare against.
//
// Returns:
//    int: -1 if the current ReactivePower is less, 1 if greater, and 0 if equal.
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

// Add adds the given ReactivePower to the current ReactivePower and returns the result.
// The result is a new ReactivePower instance with the sum of the values.
//
// Parameters:
//    other: The ReactivePower to add to the current ReactivePower.
//
// Returns:
//    *ReactivePower: A new ReactivePower instance representing the sum of both ReactivePower.
func (a *ReactivePower) Add(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ReactivePower from the current ReactivePower and returns the result.
// The result is a new ReactivePower instance with the difference of the values.
//
// Parameters:
//    other: The ReactivePower to subtract from the current ReactivePower.
//
// Returns:
//    *ReactivePower: A new ReactivePower instance representing the difference of both ReactivePower.
func (a *ReactivePower) Subtract(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ReactivePower by the given ReactivePower and returns the result.
// The result is a new ReactivePower instance with the product of the values.
//
// Parameters:
//    other: The ReactivePower to multiply with the current ReactivePower.
//
// Returns:
//    *ReactivePower: A new ReactivePower instance representing the product of both ReactivePower.
func (a *ReactivePower) Multiply(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value * other.BaseValue()}
}

// Divide divides the current ReactivePower by the given ReactivePower and returns the result.
// The result is a new ReactivePower instance with the quotient of the values.
//
// Parameters:
//    other: The ReactivePower to divide the current ReactivePower by.
//
// Returns:
//    *ReactivePower: A new ReactivePower instance representing the quotient of both ReactivePower.
func (a *ReactivePower) Divide(other *ReactivePower) *ReactivePower {
	return &ReactivePower{value: a.value / other.BaseValue()}
}

// GetReactivePowerAbbreviation gets the unit abbreviation.
func GetReactivePowerAbbreviation(unit ReactivePowerUnits) string {
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