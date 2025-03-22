// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ApparentPowerUnits defines various units of ApparentPower.
type ApparentPowerUnits string

const (
    
        // 
        ApparentPowerVoltampere ApparentPowerUnits = "Voltampere"
        // 
        ApparentPowerMicrovoltampere ApparentPowerUnits = "Microvoltampere"
        // 
        ApparentPowerMillivoltampere ApparentPowerUnits = "Millivoltampere"
        // 
        ApparentPowerKilovoltampere ApparentPowerUnits = "Kilovoltampere"
        // 
        ApparentPowerMegavoltampere ApparentPowerUnits = "Megavoltampere"
        // 
        ApparentPowerGigavoltampere ApparentPowerUnits = "Gigavoltampere"
)

var internalApparentPowerUnitsMap = map[ApparentPowerUnits]bool{
	
	ApparentPowerVoltampere: true,
	ApparentPowerMicrovoltampere: true,
	ApparentPowerMillivoltampere: true,
	ApparentPowerKilovoltampere: true,
	ApparentPowerMegavoltampere: true,
	ApparentPowerGigavoltampere: true,
}

// ApparentPowerDto represents a ApparentPower measurement with a numerical value and its corresponding unit.
type ApparentPowerDto struct {
    // Value is the numerical representation of the ApparentPower.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ApparentPower, as defined in the ApparentPowerUnits enumeration.
	Unit  ApparentPowerUnits `json:"unit" validate:"required,oneof=Voltampere,Microvoltampere,Millivoltampere,Kilovoltampere,Megavoltampere,Gigavoltampere"`
}

// ApparentPowerDtoFactory groups methods for creating and serializing ApparentPowerDto objects.
type ApparentPowerDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ApparentPowerDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ApparentPowerDtoFactory) FromJSON(data []byte) (*ApparentPowerDto, error) {
	a := ApparentPowerDto{}

    // Parse JSON into ApparentPowerDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ApparentPowerDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ApparentPowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ApparentPower represents a measurement in a of ApparentPower.
//
// Power engineers measure apparent power as the magnitude of the vector sum of active and reactive power. Apparent power is the product of the root-mean-square of voltage and current.
type ApparentPower struct {
	// value is the base measurement stored internally.
	value       float64
    
    voltamperesLazy *float64 
    microvoltamperesLazy *float64 
    millivoltamperesLazy *float64 
    kilovoltamperesLazy *float64 
    megavoltamperesLazy *float64 
    gigavoltamperesLazy *float64 
}

// ApparentPowerFactory groups methods for creating ApparentPower instances.
type ApparentPowerFactory struct{}

// CreateApparentPower creates a new ApparentPower instance from the given value and unit.
func (uf ApparentPowerFactory) CreateApparentPower(value float64, unit ApparentPowerUnits) (*ApparentPower, error) {
	return newApparentPower(value, unit)
}

// FromDto converts a ApparentPowerDto to a ApparentPower instance.
func (uf ApparentPowerFactory) FromDto(dto ApparentPowerDto) (*ApparentPower, error) {
	return newApparentPower(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ApparentPower instance.
func (uf ApparentPowerFactory) FromDtoJSON(data []byte) (*ApparentPower, error) {
	unitDto, err := ApparentPowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ApparentPowerDto from JSON: %w", err)
	}
	return ApparentPowerFactory{}.FromDto(*unitDto)
}


// FromVoltamperes creates a new ApparentPower instance from a value in Voltamperes.
func (uf ApparentPowerFactory) FromVoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerVoltampere)
}

// FromMicrovoltamperes creates a new ApparentPower instance from a value in Microvoltamperes.
func (uf ApparentPowerFactory) FromMicrovoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMicrovoltampere)
}

// FromMillivoltamperes creates a new ApparentPower instance from a value in Millivoltamperes.
func (uf ApparentPowerFactory) FromMillivoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMillivoltampere)
}

// FromKilovoltamperes creates a new ApparentPower instance from a value in Kilovoltamperes.
func (uf ApparentPowerFactory) FromKilovoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerKilovoltampere)
}

// FromMegavoltamperes creates a new ApparentPower instance from a value in Megavoltamperes.
func (uf ApparentPowerFactory) FromMegavoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerMegavoltampere)
}

// FromGigavoltamperes creates a new ApparentPower instance from a value in Gigavoltamperes.
func (uf ApparentPowerFactory) FromGigavoltamperes(value float64) (*ApparentPower, error) {
	return newApparentPower(value, ApparentPowerGigavoltampere)
}


// newApparentPower creates a new ApparentPower.
func newApparentPower(value float64, fromUnit ApparentPowerUnits) (*ApparentPower, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalApparentPowerUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ApparentPowerUnits", fromUnit)
	}
	a := &ApparentPower{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ApparentPower in Voltampere unit (the base/default quantity).
func (a *ApparentPower) BaseValue() float64 {
	return a.value
}


// Voltamperes returns the ApparentPower value in Voltamperes.
//
// 
func (a *ApparentPower) Voltamperes() float64 {
	if a.voltamperesLazy != nil {
		return *a.voltamperesLazy
	}
	voltamperes := a.convertFromBase(ApparentPowerVoltampere)
	a.voltamperesLazy = &voltamperes
	return voltamperes
}

// Microvoltamperes returns the ApparentPower value in Microvoltamperes.
//
// 
func (a *ApparentPower) Microvoltamperes() float64 {
	if a.microvoltamperesLazy != nil {
		return *a.microvoltamperesLazy
	}
	microvoltamperes := a.convertFromBase(ApparentPowerMicrovoltampere)
	a.microvoltamperesLazy = &microvoltamperes
	return microvoltamperes
}

// Millivoltamperes returns the ApparentPower value in Millivoltamperes.
//
// 
func (a *ApparentPower) Millivoltamperes() float64 {
	if a.millivoltamperesLazy != nil {
		return *a.millivoltamperesLazy
	}
	millivoltamperes := a.convertFromBase(ApparentPowerMillivoltampere)
	a.millivoltamperesLazy = &millivoltamperes
	return millivoltamperes
}

// Kilovoltamperes returns the ApparentPower value in Kilovoltamperes.
//
// 
func (a *ApparentPower) Kilovoltamperes() float64 {
	if a.kilovoltamperesLazy != nil {
		return *a.kilovoltamperesLazy
	}
	kilovoltamperes := a.convertFromBase(ApparentPowerKilovoltampere)
	a.kilovoltamperesLazy = &kilovoltamperes
	return kilovoltamperes
}

// Megavoltamperes returns the ApparentPower value in Megavoltamperes.
//
// 
func (a *ApparentPower) Megavoltamperes() float64 {
	if a.megavoltamperesLazy != nil {
		return *a.megavoltamperesLazy
	}
	megavoltamperes := a.convertFromBase(ApparentPowerMegavoltampere)
	a.megavoltamperesLazy = &megavoltamperes
	return megavoltamperes
}

// Gigavoltamperes returns the ApparentPower value in Gigavoltamperes.
//
// 
func (a *ApparentPower) Gigavoltamperes() float64 {
	if a.gigavoltamperesLazy != nil {
		return *a.gigavoltamperesLazy
	}
	gigavoltamperes := a.convertFromBase(ApparentPowerGigavoltampere)
	a.gigavoltamperesLazy = &gigavoltamperes
	return gigavoltamperes
}


// ToDto creates a ApparentPowerDto representation from the ApparentPower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Voltampere by default.
func (a *ApparentPower) ToDto(holdInUnit *ApparentPowerUnits) ApparentPowerDto {
	if holdInUnit == nil {
		defaultUnit := ApparentPowerVoltampere // Default value
		holdInUnit = &defaultUnit
	}

	return ApparentPowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ApparentPower instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Voltampere by default.
func (a *ApparentPower) ToDtoJSON(holdInUnit *ApparentPowerUnits) ([]byte, error) {
	// Convert to ApparentPowerDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ApparentPower to a specific unit value.
// The function uses the provided unit type (ApparentPowerUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ApparentPower) Convert(toUnit ApparentPowerUnits) float64 {
	switch toUnit { 
    case ApparentPowerVoltampere:
		return a.Voltamperes()
    case ApparentPowerMicrovoltampere:
		return a.Microvoltamperes()
    case ApparentPowerMillivoltampere:
		return a.Millivoltamperes()
    case ApparentPowerKilovoltampere:
		return a.Kilovoltamperes()
    case ApparentPowerMegavoltampere:
		return a.Megavoltamperes()
    case ApparentPowerGigavoltampere:
		return a.Gigavoltamperes()
	default:
		return math.NaN()
	}
}

func (a *ApparentPower) convertFromBase(toUnit ApparentPowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case ApparentPowerVoltampere:
		return (value) 
	case ApparentPowerMicrovoltampere:
		return ((value) / 1e-06) 
	case ApparentPowerMillivoltampere:
		return ((value) / 0.001) 
	case ApparentPowerKilovoltampere:
		return ((value) / 1000.0) 
	case ApparentPowerMegavoltampere:
		return ((value) / 1000000.0) 
	case ApparentPowerGigavoltampere:
		return ((value) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ApparentPower) convertToBase(value float64, fromUnit ApparentPowerUnits) float64 {
	switch fromUnit { 
	case ApparentPowerVoltampere:
		return (value) 
	case ApparentPowerMicrovoltampere:
		return ((value) * 1e-06) 
	case ApparentPowerMillivoltampere:
		return ((value) * 0.001) 
	case ApparentPowerKilovoltampere:
		return ((value) * 1000.0) 
	case ApparentPowerMegavoltampere:
		return ((value) * 1000000.0) 
	case ApparentPowerGigavoltampere:
		return ((value) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ApparentPower in the default unit (Voltampere),
// formatted to two decimal places.
func (a ApparentPower) String() string {
	return a.ToString(ApparentPowerVoltampere, 2)
}

// ToString formats the ApparentPower value as a string with the specified unit and fractional digits.
// It converts the ApparentPower to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ApparentPower value will be converted (e.g., Voltampere))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ApparentPower with the unit abbreviation.
func (a *ApparentPower) ToString(unit ApparentPowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetApparentPowerAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetApparentPowerAbbreviation(unit))
}

// Equals checks if the given ApparentPower is equal to the current ApparentPower.
//
// Parameters:
//    other: The ApparentPower to compare against.
//
// Returns:
//    bool: Returns true if both ApparentPower are equal, false otherwise.
func (a *ApparentPower) Equals(other *ApparentPower) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ApparentPower with another ApparentPower.
// It returns -1 if the current ApparentPower is less than the other ApparentPower, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ApparentPower to compare against.
//
// Returns:
//    int: -1 if the current ApparentPower is less, 1 if greater, and 0 if equal.
func (a *ApparentPower) CompareTo(other *ApparentPower) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ApparentPower to the current ApparentPower and returns the result.
// The result is a new ApparentPower instance with the sum of the values.
//
// Parameters:
//    other: The ApparentPower to add to the current ApparentPower.
//
// Returns:
//    *ApparentPower: A new ApparentPower instance representing the sum of both ApparentPower.
func (a *ApparentPower) Add(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ApparentPower from the current ApparentPower and returns the result.
// The result is a new ApparentPower instance with the difference of the values.
//
// Parameters:
//    other: The ApparentPower to subtract from the current ApparentPower.
//
// Returns:
//    *ApparentPower: A new ApparentPower instance representing the difference of both ApparentPower.
func (a *ApparentPower) Subtract(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ApparentPower by the given ApparentPower and returns the result.
// The result is a new ApparentPower instance with the product of the values.
//
// Parameters:
//    other: The ApparentPower to multiply with the current ApparentPower.
//
// Returns:
//    *ApparentPower: A new ApparentPower instance representing the product of both ApparentPower.
func (a *ApparentPower) Multiply(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value * other.BaseValue()}
}

// Divide divides the current ApparentPower by the given ApparentPower and returns the result.
// The result is a new ApparentPower instance with the quotient of the values.
//
// Parameters:
//    other: The ApparentPower to divide the current ApparentPower by.
//
// Returns:
//    *ApparentPower: A new ApparentPower instance representing the quotient of both ApparentPower.
func (a *ApparentPower) Divide(other *ApparentPower) *ApparentPower {
	return &ApparentPower{value: a.value / other.BaseValue()}
}

// GetApparentPowerAbbreviation gets the unit abbreviation.
func GetApparentPowerAbbreviation(unit ApparentPowerUnits) string {
	switch unit { 
	case ApparentPowerVoltampere:
		return "VA" 
	case ApparentPowerMicrovoltampere:
		return "μVA" 
	case ApparentPowerMillivoltampere:
		return "mVA" 
	case ApparentPowerKilovoltampere:
		return "kVA" 
	case ApparentPowerMegavoltampere:
		return "MVA" 
	case ApparentPowerGigavoltampere:
		return "GVA" 
	default:
		return ""
	}
}