// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricApparentPowerUnits defines various units of ElectricApparentPower.
type ElectricApparentPowerUnits string

const (
    
        // 
        ElectricApparentPowerVoltampere ElectricApparentPowerUnits = "Voltampere"
        // 
        ElectricApparentPowerMicrovoltampere ElectricApparentPowerUnits = "Microvoltampere"
        // 
        ElectricApparentPowerMillivoltampere ElectricApparentPowerUnits = "Millivoltampere"
        // 
        ElectricApparentPowerKilovoltampere ElectricApparentPowerUnits = "Kilovoltampere"
        // 
        ElectricApparentPowerMegavoltampere ElectricApparentPowerUnits = "Megavoltampere"
        // 
        ElectricApparentPowerGigavoltampere ElectricApparentPowerUnits = "Gigavoltampere"
)

var internalElectricApparentPowerUnitsMap = map[ElectricApparentPowerUnits]bool{
	
	ElectricApparentPowerVoltampere: true,
	ElectricApparentPowerMicrovoltampere: true,
	ElectricApparentPowerMillivoltampere: true,
	ElectricApparentPowerKilovoltampere: true,
	ElectricApparentPowerMegavoltampere: true,
	ElectricApparentPowerGigavoltampere: true,
}

// ElectricApparentPowerDto represents a ElectricApparentPower measurement with a numerical value and its corresponding unit.
type ElectricApparentPowerDto struct {
    // Value is the numerical representation of the ElectricApparentPower.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricApparentPower, as defined in the ElectricApparentPowerUnits enumeration.
	Unit  ElectricApparentPowerUnits `json:"unit" validate:"required,oneof=Voltampere Microvoltampere Millivoltampere Kilovoltampere Megavoltampere Gigavoltampere"`
}

// ElectricApparentPowerDtoFactory groups methods for creating and serializing ElectricApparentPowerDto objects.
type ElectricApparentPowerDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricApparentPowerDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricApparentPowerDtoFactory) FromJSON(data []byte) (*ElectricApparentPowerDto, error) {
	a := ElectricApparentPowerDto{}

    // Parse JSON into ElectricApparentPowerDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricApparentPowerDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricApparentPowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricApparentPower represents a measurement in a of ElectricApparentPower.
//
// Power engineers measure apparent power as the magnitude of the vector sum of active and reactive power. It is the product of the root mean square voltage (in volts) and the root mean square current (in amperes).
type ElectricApparentPower struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltamperesLazy *float64 
    microvoltamperesLazy *float64 
    millivoltamperesLazy *float64 
    kilovoltamperesLazy *float64 
    megavoltamperesLazy *float64 
    gigavoltamperesLazy *float64 
}

// ElectricApparentPowerFactory groups methods for creating ElectricApparentPower instances.
type ElectricApparentPowerFactory struct{}

// CreateElectricApparentPower creates a new ElectricApparentPower instance from the given value and unit.
func (uf ElectricApparentPowerFactory) CreateElectricApparentPower(value float64, unit ElectricApparentPowerUnits) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, unit)
}

// FromDto converts a ElectricApparentPowerDto to a ElectricApparentPower instance.
func (uf ElectricApparentPowerFactory) FromDto(dto ElectricApparentPowerDto) (*ElectricApparentPower, error) {
	return newElectricApparentPower(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricApparentPower instance.
func (uf ElectricApparentPowerFactory) FromDtoJSON(data []byte) (*ElectricApparentPower, error) {
	unitDto, err := ElectricApparentPowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricApparentPowerDto from JSON: %w", err)
	}
	return ElectricApparentPowerFactory{}.FromDto(*unitDto)
}


// FromVoltamperes creates a new ElectricApparentPower instance from a value in Voltamperes.
func (uf ElectricApparentPowerFactory) FromVoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerVoltampere)
}

// FromMicrovoltamperes creates a new ElectricApparentPower instance from a value in Microvoltamperes.
func (uf ElectricApparentPowerFactory) FromMicrovoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerMicrovoltampere)
}

// FromMillivoltamperes creates a new ElectricApparentPower instance from a value in Millivoltamperes.
func (uf ElectricApparentPowerFactory) FromMillivoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerMillivoltampere)
}

// FromKilovoltamperes creates a new ElectricApparentPower instance from a value in Kilovoltamperes.
func (uf ElectricApparentPowerFactory) FromKilovoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerKilovoltampere)
}

// FromMegavoltamperes creates a new ElectricApparentPower instance from a value in Megavoltamperes.
func (uf ElectricApparentPowerFactory) FromMegavoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerMegavoltampere)
}

// FromGigavoltamperes creates a new ElectricApparentPower instance from a value in Gigavoltamperes.
func (uf ElectricApparentPowerFactory) FromGigavoltamperes(value float64) (*ElectricApparentPower, error) {
	return newElectricApparentPower(value, ElectricApparentPowerGigavoltampere)
}


// newElectricApparentPower creates a new ElectricApparentPower.
func newElectricApparentPower(value float64, fromUnit ElectricApparentPowerUnits) (*ElectricApparentPower, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricApparentPowerUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricApparentPowerUnits", fromUnit)
	}
	a := &ElectricApparentPower{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricApparentPower in Voltampere unit (the base/default quantity).
func (a *ElectricApparentPower) BaseValue() float64 {
	return a.value
}


// Voltamperes returns the ElectricApparentPower value in Voltamperes.
//
// 
func (a *ElectricApparentPower) Voltamperes() float64 {
	if a.voltamperesLazy != nil {
		return *a.voltamperesLazy
	}
	voltamperes := a.convertFromBase(ElectricApparentPowerVoltampere)
	a.voltamperesLazy = &voltamperes
	return voltamperes
}

// Microvoltamperes returns the ElectricApparentPower value in Microvoltamperes.
//
// 
func (a *ElectricApparentPower) Microvoltamperes() float64 {
	if a.microvoltamperesLazy != nil {
		return *a.microvoltamperesLazy
	}
	microvoltamperes := a.convertFromBase(ElectricApparentPowerMicrovoltampere)
	a.microvoltamperesLazy = &microvoltamperes
	return microvoltamperes
}

// Millivoltamperes returns the ElectricApparentPower value in Millivoltamperes.
//
// 
func (a *ElectricApparentPower) Millivoltamperes() float64 {
	if a.millivoltamperesLazy != nil {
		return *a.millivoltamperesLazy
	}
	millivoltamperes := a.convertFromBase(ElectricApparentPowerMillivoltampere)
	a.millivoltamperesLazy = &millivoltamperes
	return millivoltamperes
}

// Kilovoltamperes returns the ElectricApparentPower value in Kilovoltamperes.
//
// 
func (a *ElectricApparentPower) Kilovoltamperes() float64 {
	if a.kilovoltamperesLazy != nil {
		return *a.kilovoltamperesLazy
	}
	kilovoltamperes := a.convertFromBase(ElectricApparentPowerKilovoltampere)
	a.kilovoltamperesLazy = &kilovoltamperes
	return kilovoltamperes
}

// Megavoltamperes returns the ElectricApparentPower value in Megavoltamperes.
//
// 
func (a *ElectricApparentPower) Megavoltamperes() float64 {
	if a.megavoltamperesLazy != nil {
		return *a.megavoltamperesLazy
	}
	megavoltamperes := a.convertFromBase(ElectricApparentPowerMegavoltampere)
	a.megavoltamperesLazy = &megavoltamperes
	return megavoltamperes
}

// Gigavoltamperes returns the ElectricApparentPower value in Gigavoltamperes.
//
// 
func (a *ElectricApparentPower) Gigavoltamperes() float64 {
	if a.gigavoltamperesLazy != nil {
		return *a.gigavoltamperesLazy
	}
	gigavoltamperes := a.convertFromBase(ElectricApparentPowerGigavoltampere)
	a.gigavoltamperesLazy = &gigavoltamperes
	return gigavoltamperes
}


// ToDto creates a ElectricApparentPowerDto representation from the ElectricApparentPower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Voltampere by default.
func (a *ElectricApparentPower) ToDto(holdInUnit *ElectricApparentPowerUnits) ElectricApparentPowerDto {
	if holdInUnit == nil {
		defaultUnit := ElectricApparentPowerVoltampere // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricApparentPowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricApparentPower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Voltampere by default.
func (a *ElectricApparentPower) ToDtoJSON(holdInUnit *ElectricApparentPowerUnits) ([]byte, error) {
	// Convert to ElectricApparentPowerDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricApparentPower to a specific unit value.
// The function uses the provided unit type (ElectricApparentPowerUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricApparentPower) Convert(toUnit ElectricApparentPowerUnits) float64 {
	switch toUnit { 
    case ElectricApparentPowerVoltampere:
		return a.Voltamperes()
    case ElectricApparentPowerMicrovoltampere:
		return a.Microvoltamperes()
    case ElectricApparentPowerMillivoltampere:
		return a.Millivoltamperes()
    case ElectricApparentPowerKilovoltampere:
		return a.Kilovoltamperes()
    case ElectricApparentPowerMegavoltampere:
		return a.Megavoltamperes()
    case ElectricApparentPowerGigavoltampere:
		return a.Gigavoltamperes()
	default:
		return math.NaN()
	}
}

func (a *ElectricApparentPower) convertFromBase(toUnit ElectricApparentPowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricApparentPowerVoltampere:
		return (value) 
	case ElectricApparentPowerMicrovoltampere:
		return ((value) / 1e-06) 
	case ElectricApparentPowerMillivoltampere:
		return ((value) / 0.001) 
	case ElectricApparentPowerKilovoltampere:
		return ((value) / 1000.0) 
	case ElectricApparentPowerMegavoltampere:
		return ((value) / 1000000.0) 
	case ElectricApparentPowerGigavoltampere:
		return ((value) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricApparentPower) convertToBase(value float64, fromUnit ElectricApparentPowerUnits) float64 {
	switch fromUnit { 
	case ElectricApparentPowerVoltampere:
		return (value) 
	case ElectricApparentPowerMicrovoltampere:
		return ((value) * 1e-06) 
	case ElectricApparentPowerMillivoltampere:
		return ((value) * 0.001) 
	case ElectricApparentPowerKilovoltampere:
		return ((value) * 1000.0) 
	case ElectricApparentPowerMegavoltampere:
		return ((value) * 1000000.0) 
	case ElectricApparentPowerGigavoltampere:
		return ((value) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricApparentPower in the default unit (Voltampere),
// formatted to two decimal places.
func (a ElectricApparentPower) String() string {
	return a.ToString(ElectricApparentPowerVoltampere, 2)
}

// ToString formats the ElectricApparentPower value as a string with the specified unit and fractional digits.
// It converts the ElectricApparentPower to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricApparentPower value will be converted (e.g., Voltampere))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricApparentPower with the unit abbreviation.
func (a *ElectricApparentPower) ToString(unit ElectricApparentPowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricApparentPowerAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricApparentPowerAbbreviation(unit))
}

// Equals checks if the given ElectricApparentPower is equal to the current ElectricApparentPower.
//
// Parameters:
//    other: The ElectricApparentPower to compare against.
//
// Returns:
//    bool: Returns true if both ElectricApparentPower are equal, false otherwise.
func (a *ElectricApparentPower) Equals(other *ElectricApparentPower) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricApparentPower with another ElectricApparentPower.
// It returns -1 if the current ElectricApparentPower is less than the other ElectricApparentPower, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricApparentPower to compare against.
//
// Returns:
//    int: -1 if the current ElectricApparentPower is less, 1 if greater, and 0 if equal.
func (a *ElectricApparentPower) CompareTo(other *ElectricApparentPower) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricApparentPower to the current ElectricApparentPower and returns the result.
// The result is a new ElectricApparentPower instance with the sum of the values.
//
// Parameters:
//    other: The ElectricApparentPower to add to the current ElectricApparentPower.
//
// Returns:
//    *ElectricApparentPower: A new ElectricApparentPower instance representing the sum of both ElectricApparentPower.
func (a *ElectricApparentPower) Add(other *ElectricApparentPower) *ElectricApparentPower {
	return &ElectricApparentPower{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricApparentPower from the current ElectricApparentPower and returns the result.
// The result is a new ElectricApparentPower instance with the difference of the values.
//
// Parameters:
//    other: The ElectricApparentPower to subtract from the current ElectricApparentPower.
//
// Returns:
//    *ElectricApparentPower: A new ElectricApparentPower instance representing the difference of both ElectricApparentPower.
func (a *ElectricApparentPower) Subtract(other *ElectricApparentPower) *ElectricApparentPower {
	return &ElectricApparentPower{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricApparentPower by the given ElectricApparentPower and returns the result.
// The result is a new ElectricApparentPower instance with the product of the values.
//
// Parameters:
//    other: The ElectricApparentPower to multiply with the current ElectricApparentPower.
//
// Returns:
//    *ElectricApparentPower: A new ElectricApparentPower instance representing the product of both ElectricApparentPower.
func (a *ElectricApparentPower) Multiply(other *ElectricApparentPower) *ElectricApparentPower {
	return &ElectricApparentPower{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricApparentPower by the given ElectricApparentPower and returns the result.
// The result is a new ElectricApparentPower instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricApparentPower to divide the current ElectricApparentPower by.
//
// Returns:
//    *ElectricApparentPower: A new ElectricApparentPower instance representing the quotient of both ElectricApparentPower.
func (a *ElectricApparentPower) Divide(other *ElectricApparentPower) *ElectricApparentPower {
	return &ElectricApparentPower{value: a.value / other.BaseValue()}
}

// GetElectricApparentPowerAbbreviation gets the unit abbreviation.
func GetElectricApparentPowerAbbreviation(unit ElectricApparentPowerUnits) string {
	switch unit { 
	case ElectricApparentPowerVoltampere:
		return "VA" 
	case ElectricApparentPowerMicrovoltampere:
		return "μVA" 
	case ElectricApparentPowerMillivoltampere:
		return "mVA" 
	case ElectricApparentPowerKilovoltampere:
		return "kVA" 
	case ElectricApparentPowerMegavoltampere:
		return "MVA" 
	case ElectricApparentPowerGigavoltampere:
		return "GVA" 
	default:
		return ""
	}
}