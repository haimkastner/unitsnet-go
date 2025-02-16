// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolarFlowUnits defines various units of MolarFlow.
type MolarFlowUnits string

const (
    
        // 
        MolarFlowMolePerSecond MolarFlowUnits = "MolePerSecond"
        // 
        MolarFlowMolePerMinute MolarFlowUnits = "MolePerMinute"
        // 
        MolarFlowMolePerHour MolarFlowUnits = "MolePerHour"
        // 
        MolarFlowPoundMolePerSecond MolarFlowUnits = "PoundMolePerSecond"
        // 
        MolarFlowPoundMolePerMinute MolarFlowUnits = "PoundMolePerMinute"
        // 
        MolarFlowPoundMolePerHour MolarFlowUnits = "PoundMolePerHour"
        // 
        MolarFlowKilomolePerSecond MolarFlowUnits = "KilomolePerSecond"
        // 
        MolarFlowKilomolePerMinute MolarFlowUnits = "KilomolePerMinute"
        // 
        MolarFlowKilomolePerHour MolarFlowUnits = "KilomolePerHour"
)

// MolarFlowDto represents a MolarFlow measurement with a numerical value and its corresponding unit.
type MolarFlowDto struct {
    // Value is the numerical representation of the MolarFlow.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MolarFlow, as defined in the MolarFlowUnits enumeration.
	Unit  MolarFlowUnits `json:"unit"`
}

// MolarFlowDtoFactory groups methods for creating and serializing MolarFlowDto objects.
type MolarFlowDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolarFlowDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolarFlowDtoFactory) FromJSON(data []byte) (*MolarFlowDto, error) {
	a := MolarFlowDto{}

    // Parse JSON into MolarFlowDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MolarFlowDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolarFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MolarFlow represents a measurement in a of MolarFlow.
//
// Molar flow is the ratio of the amount of substance change to the time during which the change occurred (value of amount of substance changes per unit time).
type MolarFlow struct {
	// value is the base measurement stored internally.
	value       float64
    
    moles_per_secondLazy *float64 
    moles_per_minuteLazy *float64 
    moles_per_hourLazy *float64 
    pound_moles_per_secondLazy *float64 
    pound_moles_per_minuteLazy *float64 
    pound_moles_per_hourLazy *float64 
    kilomoles_per_secondLazy *float64 
    kilomoles_per_minuteLazy *float64 
    kilomoles_per_hourLazy *float64 
}

// MolarFlowFactory groups methods for creating MolarFlow instances.
type MolarFlowFactory struct{}

// CreateMolarFlow creates a new MolarFlow instance from the given value and unit.
func (uf MolarFlowFactory) CreateMolarFlow(value float64, unit MolarFlowUnits) (*MolarFlow, error) {
	return newMolarFlow(value, unit)
}

// FromDto converts a MolarFlowDto to a MolarFlow instance.
func (uf MolarFlowFactory) FromDto(dto MolarFlowDto) (*MolarFlow, error) {
	return newMolarFlow(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MolarFlow instance.
func (uf MolarFlowFactory) FromDtoJSON(data []byte) (*MolarFlow, error) {
	unitDto, err := MolarFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolarFlowDto from JSON: %w", err)
	}
	return MolarFlowFactory{}.FromDto(*unitDto)
}


// FromMolesPerSecond creates a new MolarFlow instance from a value in MolesPerSecond.
func (uf MolarFlowFactory) FromMolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerSecond)
}

// FromMolesPerMinute creates a new MolarFlow instance from a value in MolesPerMinute.
func (uf MolarFlowFactory) FromMolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerMinute)
}

// FromMolesPerHour creates a new MolarFlow instance from a value in MolesPerHour.
func (uf MolarFlowFactory) FromMolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowMolePerHour)
}

// FromPoundMolesPerSecond creates a new MolarFlow instance from a value in PoundMolesPerSecond.
func (uf MolarFlowFactory) FromPoundMolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerSecond)
}

// FromPoundMolesPerMinute creates a new MolarFlow instance from a value in PoundMolesPerMinute.
func (uf MolarFlowFactory) FromPoundMolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerMinute)
}

// FromPoundMolesPerHour creates a new MolarFlow instance from a value in PoundMolesPerHour.
func (uf MolarFlowFactory) FromPoundMolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowPoundMolePerHour)
}

// FromKilomolesPerSecond creates a new MolarFlow instance from a value in KilomolesPerSecond.
func (uf MolarFlowFactory) FromKilomolesPerSecond(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerSecond)
}

// FromKilomolesPerMinute creates a new MolarFlow instance from a value in KilomolesPerMinute.
func (uf MolarFlowFactory) FromKilomolesPerMinute(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerMinute)
}

// FromKilomolesPerHour creates a new MolarFlow instance from a value in KilomolesPerHour.
func (uf MolarFlowFactory) FromKilomolesPerHour(value float64) (*MolarFlow, error) {
	return newMolarFlow(value, MolarFlowKilomolePerHour)
}


// newMolarFlow creates a new MolarFlow.
func newMolarFlow(value float64, fromUnit MolarFlowUnits) (*MolarFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MolarFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarFlow in MolePerSecond unit (the base/default quantity).
func (a *MolarFlow) BaseValue() float64 {
	return a.value
}


// MolesPerSecond returns the MolarFlow value in MolesPerSecond.
//
// 
func (a *MolarFlow) MolesPerSecond() float64 {
	if a.moles_per_secondLazy != nil {
		return *a.moles_per_secondLazy
	}
	moles_per_second := a.convertFromBase(MolarFlowMolePerSecond)
	a.moles_per_secondLazy = &moles_per_second
	return moles_per_second
}

// MolesPerMinute returns the MolarFlow value in MolesPerMinute.
//
// 
func (a *MolarFlow) MolesPerMinute() float64 {
	if a.moles_per_minuteLazy != nil {
		return *a.moles_per_minuteLazy
	}
	moles_per_minute := a.convertFromBase(MolarFlowMolePerMinute)
	a.moles_per_minuteLazy = &moles_per_minute
	return moles_per_minute
}

// MolesPerHour returns the MolarFlow value in MolesPerHour.
//
// 
func (a *MolarFlow) MolesPerHour() float64 {
	if a.moles_per_hourLazy != nil {
		return *a.moles_per_hourLazy
	}
	moles_per_hour := a.convertFromBase(MolarFlowMolePerHour)
	a.moles_per_hourLazy = &moles_per_hour
	return moles_per_hour
}

// PoundMolesPerSecond returns the MolarFlow value in PoundMolesPerSecond.
//
// 
func (a *MolarFlow) PoundMolesPerSecond() float64 {
	if a.pound_moles_per_secondLazy != nil {
		return *a.pound_moles_per_secondLazy
	}
	pound_moles_per_second := a.convertFromBase(MolarFlowPoundMolePerSecond)
	a.pound_moles_per_secondLazy = &pound_moles_per_second
	return pound_moles_per_second
}

// PoundMolesPerMinute returns the MolarFlow value in PoundMolesPerMinute.
//
// 
func (a *MolarFlow) PoundMolesPerMinute() float64 {
	if a.pound_moles_per_minuteLazy != nil {
		return *a.pound_moles_per_minuteLazy
	}
	pound_moles_per_minute := a.convertFromBase(MolarFlowPoundMolePerMinute)
	a.pound_moles_per_minuteLazy = &pound_moles_per_minute
	return pound_moles_per_minute
}

// PoundMolesPerHour returns the MolarFlow value in PoundMolesPerHour.
//
// 
func (a *MolarFlow) PoundMolesPerHour() float64 {
	if a.pound_moles_per_hourLazy != nil {
		return *a.pound_moles_per_hourLazy
	}
	pound_moles_per_hour := a.convertFromBase(MolarFlowPoundMolePerHour)
	a.pound_moles_per_hourLazy = &pound_moles_per_hour
	return pound_moles_per_hour
}

// KilomolesPerSecond returns the MolarFlow value in KilomolesPerSecond.
//
// 
func (a *MolarFlow) KilomolesPerSecond() float64 {
	if a.kilomoles_per_secondLazy != nil {
		return *a.kilomoles_per_secondLazy
	}
	kilomoles_per_second := a.convertFromBase(MolarFlowKilomolePerSecond)
	a.kilomoles_per_secondLazy = &kilomoles_per_second
	return kilomoles_per_second
}

// KilomolesPerMinute returns the MolarFlow value in KilomolesPerMinute.
//
// 
func (a *MolarFlow) KilomolesPerMinute() float64 {
	if a.kilomoles_per_minuteLazy != nil {
		return *a.kilomoles_per_minuteLazy
	}
	kilomoles_per_minute := a.convertFromBase(MolarFlowKilomolePerMinute)
	a.kilomoles_per_minuteLazy = &kilomoles_per_minute
	return kilomoles_per_minute
}

// KilomolesPerHour returns the MolarFlow value in KilomolesPerHour.
//
// 
func (a *MolarFlow) KilomolesPerHour() float64 {
	if a.kilomoles_per_hourLazy != nil {
		return *a.kilomoles_per_hourLazy
	}
	kilomoles_per_hour := a.convertFromBase(MolarFlowKilomolePerHour)
	a.kilomoles_per_hourLazy = &kilomoles_per_hour
	return kilomoles_per_hour
}


// ToDto creates a MolarFlowDto representation from the MolarFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerSecond by default.
func (a *MolarFlow) ToDto(holdInUnit *MolarFlowUnits) MolarFlowDto {
	if holdInUnit == nil {
		defaultUnit := MolarFlowMolePerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return MolarFlowDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MolarFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerSecond by default.
func (a *MolarFlow) ToDtoJSON(holdInUnit *MolarFlowUnits) ([]byte, error) {
	// Convert to MolarFlowDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MolarFlow to a specific unit value.
// The function uses the provided unit type (MolarFlowUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MolarFlow) Convert(toUnit MolarFlowUnits) float64 {
	switch toUnit { 
    case MolarFlowMolePerSecond:
		return a.MolesPerSecond()
    case MolarFlowMolePerMinute:
		return a.MolesPerMinute()
    case MolarFlowMolePerHour:
		return a.MolesPerHour()
    case MolarFlowPoundMolePerSecond:
		return a.PoundMolesPerSecond()
    case MolarFlowPoundMolePerMinute:
		return a.PoundMolesPerMinute()
    case MolarFlowPoundMolePerHour:
		return a.PoundMolesPerHour()
    case MolarFlowKilomolePerSecond:
		return a.KilomolesPerSecond()
    case MolarFlowKilomolePerMinute:
		return a.KilomolesPerMinute()
    case MolarFlowKilomolePerHour:
		return a.KilomolesPerHour()
	default:
		return math.NaN()
	}
}

func (a *MolarFlow) convertFromBase(toUnit MolarFlowUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarFlowMolePerSecond:
		return (value) 
	case MolarFlowMolePerMinute:
		return (value * 60) 
	case MolarFlowMolePerHour:
		return (value * 3600) 
	case MolarFlowPoundMolePerSecond:
		return (value / 453.59237) 
	case MolarFlowPoundMolePerMinute:
		return ((value / 453.59237) * 60) 
	case MolarFlowPoundMolePerHour:
		return ((value / 453.59237) * 3600) 
	case MolarFlowKilomolePerSecond:
		return ((value) / 1000.0) 
	case MolarFlowKilomolePerMinute:
		return ((value * 60) / 1000.0) 
	case MolarFlowKilomolePerHour:
		return ((value * 3600) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *MolarFlow) convertToBase(value float64, fromUnit MolarFlowUnits) float64 {
	switch fromUnit { 
	case MolarFlowMolePerSecond:
		return (value) 
	case MolarFlowMolePerMinute:
		return (value / 60) 
	case MolarFlowMolePerHour:
		return (value / 3600) 
	case MolarFlowPoundMolePerSecond:
		return (value * 453.59237) 
	case MolarFlowPoundMolePerMinute:
		return ((value * 453.59237) / 60) 
	case MolarFlowPoundMolePerHour:
		return ((value * 453.59237) / 3600) 
	case MolarFlowKilomolePerSecond:
		return ((value) * 1000.0) 
	case MolarFlowKilomolePerMinute:
		return ((value / 60) * 1000.0) 
	case MolarFlowKilomolePerHour:
		return ((value / 3600) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MolarFlow in the default unit (MolePerSecond),
// formatted to two decimal places.
func (a MolarFlow) String() string {
	return a.ToString(MolarFlowMolePerSecond, 2)
}

// ToString formats the MolarFlow value as a string with the specified unit and fractional digits.
// It converts the MolarFlow to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MolarFlow value will be converted (e.g., MolePerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MolarFlow with the unit abbreviation.
func (a *MolarFlow) ToString(unit MolarFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMolarFlowAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMolarFlowAbbreviation(unit))
}

// Equals checks if the given MolarFlow is equal to the current MolarFlow.
//
// Parameters:
//    other: The MolarFlow to compare against.
//
// Returns:
//    bool: Returns true if both MolarFlow are equal, false otherwise.
func (a *MolarFlow) Equals(other *MolarFlow) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MolarFlow with another MolarFlow.
// It returns -1 if the current MolarFlow is less than the other MolarFlow, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MolarFlow to compare against.
//
// Returns:
//    int: -1 if the current MolarFlow is less, 1 if greater, and 0 if equal.
func (a *MolarFlow) CompareTo(other *MolarFlow) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MolarFlow to the current MolarFlow and returns the result.
// The result is a new MolarFlow instance with the sum of the values.
//
// Parameters:
//    other: The MolarFlow to add to the current MolarFlow.
//
// Returns:
//    *MolarFlow: A new MolarFlow instance representing the sum of both MolarFlow.
func (a *MolarFlow) Add(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MolarFlow from the current MolarFlow and returns the result.
// The result is a new MolarFlow instance with the difference of the values.
//
// Parameters:
//    other: The MolarFlow to subtract from the current MolarFlow.
//
// Returns:
//    *MolarFlow: A new MolarFlow instance representing the difference of both MolarFlow.
func (a *MolarFlow) Subtract(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MolarFlow by the given MolarFlow and returns the result.
// The result is a new MolarFlow instance with the product of the values.
//
// Parameters:
//    other: The MolarFlow to multiply with the current MolarFlow.
//
// Returns:
//    *MolarFlow: A new MolarFlow instance representing the product of both MolarFlow.
func (a *MolarFlow) Multiply(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value * other.BaseValue()}
}

// Divide divides the current MolarFlow by the given MolarFlow and returns the result.
// The result is a new MolarFlow instance with the quotient of the values.
//
// Parameters:
//    other: The MolarFlow to divide the current MolarFlow by.
//
// Returns:
//    *MolarFlow: A new MolarFlow instance representing the quotient of both MolarFlow.
func (a *MolarFlow) Divide(other *MolarFlow) *MolarFlow {
	return &MolarFlow{value: a.value / other.BaseValue()}
}

// GetMolarFlowAbbreviation gets the unit abbreviation.
func GetMolarFlowAbbreviation(unit MolarFlowUnits) string {
	switch unit { 
	case MolarFlowMolePerSecond:
		return "mol/s" 
	case MolarFlowMolePerMinute:
		return "mol/min" 
	case MolarFlowMolePerHour:
		return "kmol/h" 
	case MolarFlowPoundMolePerSecond:
		return "lbmol/s" 
	case MolarFlowPoundMolePerMinute:
		return "lbmol/min" 
	case MolarFlowPoundMolePerHour:
		return "lbmol/h" 
	case MolarFlowKilomolePerSecond:
		return "kmol/s" 
	case MolarFlowKilomolePerMinute:
		return "kmol/min" 
	case MolarFlowKilomolePerHour:
		return "kkmol/h" 
	default:
		return ""
	}
}