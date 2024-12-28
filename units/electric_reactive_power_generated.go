// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricReactivePowerUnits defines various units of ElectricReactivePower.
type ElectricReactivePowerUnits string

const (
    
        // 
        ElectricReactivePowerVoltampereReactive ElectricReactivePowerUnits = "VoltampereReactive"
        // 
        ElectricReactivePowerKilovoltampereReactive ElectricReactivePowerUnits = "KilovoltampereReactive"
        // 
        ElectricReactivePowerMegavoltampereReactive ElectricReactivePowerUnits = "MegavoltampereReactive"
        // 
        ElectricReactivePowerGigavoltampereReactive ElectricReactivePowerUnits = "GigavoltampereReactive"
)

// ElectricReactivePowerDto represents a ElectricReactivePower measurement with a numerical value and its corresponding unit.
type ElectricReactivePowerDto struct {
    // Value is the numerical representation of the ElectricReactivePower.
	Value float64
    // Unit specifies the unit of measurement for the ElectricReactivePower, as defined in the ElectricReactivePowerUnits enumeration.
	Unit  ElectricReactivePowerUnits
}

// ElectricReactivePowerDtoFactory groups methods for creating and serializing ElectricReactivePowerDto objects.
type ElectricReactivePowerDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactivePowerDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricReactivePowerDtoFactory) FromJSON(data []byte) (*ElectricReactivePowerDto, error) {
	a := ElectricReactivePowerDto{}

    // Parse JSON into ElectricReactivePowerDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricReactivePowerDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricReactivePowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ElectricReactivePower represents a measurement in a of ElectricReactivePower.
//
// In electric power transmission and distribution, volt-ampere reactive (var) is a unit of measurement of reactive power. Reactive power exists in an AC circuit when the current and voltage are not in phase.
type ElectricReactivePower struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltamperes_reactiveLazy *float64 
    kilovoltamperes_reactiveLazy *float64 
    megavoltamperes_reactiveLazy *float64 
    gigavoltamperes_reactiveLazy *float64 
}

// ElectricReactivePowerFactory groups methods for creating ElectricReactivePower instances.
type ElectricReactivePowerFactory struct{}

// CreateElectricReactivePower creates a new ElectricReactivePower instance from the given value and unit.
func (uf ElectricReactivePowerFactory) CreateElectricReactivePower(value float64, unit ElectricReactivePowerUnits) (*ElectricReactivePower, error) {
	return newElectricReactivePower(value, unit)
}

// FromDto converts a ElectricReactivePowerDto to a ElectricReactivePower instance.
func (uf ElectricReactivePowerFactory) FromDto(dto ElectricReactivePowerDto) (*ElectricReactivePower, error) {
	return newElectricReactivePower(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactivePower instance.
func (uf ElectricReactivePowerFactory) FromDtoJSON(data []byte) (*ElectricReactivePower, error) {
	unitDto, err := ElectricReactivePowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricReactivePowerDto from JSON: %w", err)
	}
	return ElectricReactivePowerFactory{}.FromDto(*unitDto)
}


// FromVoltamperesReactive creates a new ElectricReactivePower instance from a value in VoltamperesReactive.
func (uf ElectricReactivePowerFactory) FromVoltamperesReactive(value float64) (*ElectricReactivePower, error) {
	return newElectricReactivePower(value, ElectricReactivePowerVoltampereReactive)
}

// FromKilovoltamperesReactive creates a new ElectricReactivePower instance from a value in KilovoltamperesReactive.
func (uf ElectricReactivePowerFactory) FromKilovoltamperesReactive(value float64) (*ElectricReactivePower, error) {
	return newElectricReactivePower(value, ElectricReactivePowerKilovoltampereReactive)
}

// FromMegavoltamperesReactive creates a new ElectricReactivePower instance from a value in MegavoltamperesReactive.
func (uf ElectricReactivePowerFactory) FromMegavoltamperesReactive(value float64) (*ElectricReactivePower, error) {
	return newElectricReactivePower(value, ElectricReactivePowerMegavoltampereReactive)
}

// FromGigavoltamperesReactive creates a new ElectricReactivePower instance from a value in GigavoltamperesReactive.
func (uf ElectricReactivePowerFactory) FromGigavoltamperesReactive(value float64) (*ElectricReactivePower, error) {
	return newElectricReactivePower(value, ElectricReactivePowerGigavoltampereReactive)
}


// newElectricReactivePower creates a new ElectricReactivePower.
func newElectricReactivePower(value float64, fromUnit ElectricReactivePowerUnits) (*ElectricReactivePower, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricReactivePower{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricReactivePower in VoltampereReactive unit (the base/default quantity).
func (a *ElectricReactivePower) BaseValue() float64 {
	return a.value
}


// VoltamperesReactive returns the ElectricReactivePower value in VoltamperesReactive.
//
// 
func (a *ElectricReactivePower) VoltamperesReactive() float64 {
	if a.voltamperes_reactiveLazy != nil {
		return *a.voltamperes_reactiveLazy
	}
	voltamperes_reactive := a.convertFromBase(ElectricReactivePowerVoltampereReactive)
	a.voltamperes_reactiveLazy = &voltamperes_reactive
	return voltamperes_reactive
}

// KilovoltamperesReactive returns the ElectricReactivePower value in KilovoltamperesReactive.
//
// 
func (a *ElectricReactivePower) KilovoltamperesReactive() float64 {
	if a.kilovoltamperes_reactiveLazy != nil {
		return *a.kilovoltamperes_reactiveLazy
	}
	kilovoltamperes_reactive := a.convertFromBase(ElectricReactivePowerKilovoltampereReactive)
	a.kilovoltamperes_reactiveLazy = &kilovoltamperes_reactive
	return kilovoltamperes_reactive
}

// MegavoltamperesReactive returns the ElectricReactivePower value in MegavoltamperesReactive.
//
// 
func (a *ElectricReactivePower) MegavoltamperesReactive() float64 {
	if a.megavoltamperes_reactiveLazy != nil {
		return *a.megavoltamperes_reactiveLazy
	}
	megavoltamperes_reactive := a.convertFromBase(ElectricReactivePowerMegavoltampereReactive)
	a.megavoltamperes_reactiveLazy = &megavoltamperes_reactive
	return megavoltamperes_reactive
}

// GigavoltamperesReactive returns the ElectricReactivePower value in GigavoltamperesReactive.
//
// 
func (a *ElectricReactivePower) GigavoltamperesReactive() float64 {
	if a.gigavoltamperes_reactiveLazy != nil {
		return *a.gigavoltamperes_reactiveLazy
	}
	gigavoltamperes_reactive := a.convertFromBase(ElectricReactivePowerGigavoltampereReactive)
	a.gigavoltamperes_reactiveLazy = &gigavoltamperes_reactive
	return gigavoltamperes_reactive
}


// ToDto creates a ElectricReactivePowerDto representation from the ElectricReactivePower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactive by default.
func (a *ElectricReactivePower) ToDto(holdInUnit *ElectricReactivePowerUnits) ElectricReactivePowerDto {
	if holdInUnit == nil {
		defaultUnit := ElectricReactivePowerVoltampereReactive // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricReactivePowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricReactivePower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by VoltampereReactive by default.
func (a *ElectricReactivePower) ToDtoJSON(holdInUnit *ElectricReactivePowerUnits) ([]byte, error) {
	// Convert to ElectricReactivePowerDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricReactivePower to a specific unit value.
// The function uses the provided unit type (ElectricReactivePowerUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricReactivePower) Convert(toUnit ElectricReactivePowerUnits) float64 {
	switch toUnit { 
    case ElectricReactivePowerVoltampereReactive:
		return a.VoltamperesReactive()
    case ElectricReactivePowerKilovoltampereReactive:
		return a.KilovoltamperesReactive()
    case ElectricReactivePowerMegavoltampereReactive:
		return a.MegavoltamperesReactive()
    case ElectricReactivePowerGigavoltampereReactive:
		return a.GigavoltamperesReactive()
	default:
		return math.NaN()
	}
}

func (a *ElectricReactivePower) convertFromBase(toUnit ElectricReactivePowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricReactivePowerVoltampereReactive:
		return (value) 
	case ElectricReactivePowerKilovoltampereReactive:
		return ((value) / 1000.0) 
	case ElectricReactivePowerMegavoltampereReactive:
		return ((value) / 1000000.0) 
	case ElectricReactivePowerGigavoltampereReactive:
		return ((value) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricReactivePower) convertToBase(value float64, fromUnit ElectricReactivePowerUnits) float64 {
	switch fromUnit { 
	case ElectricReactivePowerVoltampereReactive:
		return (value) 
	case ElectricReactivePowerKilovoltampereReactive:
		return ((value) * 1000.0) 
	case ElectricReactivePowerMegavoltampereReactive:
		return ((value) * 1000000.0) 
	case ElectricReactivePowerGigavoltampereReactive:
		return ((value) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricReactivePower in the default unit (VoltampereReactive),
// formatted to two decimal places.
func (a ElectricReactivePower) String() string {
	return a.ToString(ElectricReactivePowerVoltampereReactive, 2)
}

// ToString formats the ElectricReactivePower value as a string with the specified unit and fractional digits.
// It converts the ElectricReactivePower to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricReactivePower value will be converted (e.g., VoltampereReactive))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricReactivePower with the unit abbreviation.
func (a *ElectricReactivePower) ToString(unit ElectricReactivePowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricReactivePower) getUnitAbbreviation(unit ElectricReactivePowerUnits) string {
	switch unit { 
	case ElectricReactivePowerVoltampereReactive:
		return "var" 
	case ElectricReactivePowerKilovoltampereReactive:
		return "kvar" 
	case ElectricReactivePowerMegavoltampereReactive:
		return "Mvar" 
	case ElectricReactivePowerGigavoltampereReactive:
		return "Gvar" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricReactivePower is equal to the current ElectricReactivePower.
//
// Parameters:
//    other: The ElectricReactivePower to compare against.
//
// Returns:
//    bool: Returns true if both ElectricReactivePower are equal, false otherwise.
func (a *ElectricReactivePower) Equals(other *ElectricReactivePower) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricReactivePower with another ElectricReactivePower.
// It returns -1 if the current ElectricReactivePower is less than the other ElectricReactivePower, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricReactivePower to compare against.
//
// Returns:
//    int: -1 if the current ElectricReactivePower is less, 1 if greater, and 0 if equal.
func (a *ElectricReactivePower) CompareTo(other *ElectricReactivePower) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricReactivePower to the current ElectricReactivePower and returns the result.
// The result is a new ElectricReactivePower instance with the sum of the values.
//
// Parameters:
//    other: The ElectricReactivePower to add to the current ElectricReactivePower.
//
// Returns:
//    *ElectricReactivePower: A new ElectricReactivePower instance representing the sum of both ElectricReactivePower.
func (a *ElectricReactivePower) Add(other *ElectricReactivePower) *ElectricReactivePower {
	return &ElectricReactivePower{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricReactivePower from the current ElectricReactivePower and returns the result.
// The result is a new ElectricReactivePower instance with the difference of the values.
//
// Parameters:
//    other: The ElectricReactivePower to subtract from the current ElectricReactivePower.
//
// Returns:
//    *ElectricReactivePower: A new ElectricReactivePower instance representing the difference of both ElectricReactivePower.
func (a *ElectricReactivePower) Subtract(other *ElectricReactivePower) *ElectricReactivePower {
	return &ElectricReactivePower{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricReactivePower by the given ElectricReactivePower and returns the result.
// The result is a new ElectricReactivePower instance with the product of the values.
//
// Parameters:
//    other: The ElectricReactivePower to multiply with the current ElectricReactivePower.
//
// Returns:
//    *ElectricReactivePower: A new ElectricReactivePower instance representing the product of both ElectricReactivePower.
func (a *ElectricReactivePower) Multiply(other *ElectricReactivePower) *ElectricReactivePower {
	return &ElectricReactivePower{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricReactivePower by the given ElectricReactivePower and returns the result.
// The result is a new ElectricReactivePower instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricReactivePower to divide the current ElectricReactivePower by.
//
// Returns:
//    *ElectricReactivePower: A new ElectricReactivePower instance representing the quotient of both ElectricReactivePower.
func (a *ElectricReactivePower) Divide(other *ElectricReactivePower) *ElectricReactivePower {
	return &ElectricReactivePower{value: a.value / other.BaseValue()}
}