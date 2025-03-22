// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RadiationEquivalentDoseRateUnits defines various units of RadiationEquivalentDoseRate.
type RadiationEquivalentDoseRateUnits string

const (
    
        // 
        RadiationEquivalentDoseRateSievertPerHour RadiationEquivalentDoseRateUnits = "SievertPerHour"
        // 
        RadiationEquivalentDoseRateSievertPerSecond RadiationEquivalentDoseRateUnits = "SievertPerSecond"
        // 
        RadiationEquivalentDoseRateRoentgenEquivalentManPerHour RadiationEquivalentDoseRateUnits = "RoentgenEquivalentManPerHour"
        // 
        RadiationEquivalentDoseRateNanosievertPerHour RadiationEquivalentDoseRateUnits = "NanosievertPerHour"
        // 
        RadiationEquivalentDoseRateMicrosievertPerHour RadiationEquivalentDoseRateUnits = "MicrosievertPerHour"
        // 
        RadiationEquivalentDoseRateMillisievertPerHour RadiationEquivalentDoseRateUnits = "MillisievertPerHour"
        // 
        RadiationEquivalentDoseRateNanosievertPerSecond RadiationEquivalentDoseRateUnits = "NanosievertPerSecond"
        // 
        RadiationEquivalentDoseRateMicrosievertPerSecond RadiationEquivalentDoseRateUnits = "MicrosievertPerSecond"
        // 
        RadiationEquivalentDoseRateMillisievertPerSecond RadiationEquivalentDoseRateUnits = "MillisievertPerSecond"
        // 
        RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour RadiationEquivalentDoseRateUnits = "MilliroentgenEquivalentManPerHour"
)

var internalRadiationEquivalentDoseRateUnitsMap = map[RadiationEquivalentDoseRateUnits]bool{
	
	RadiationEquivalentDoseRateSievertPerHour: true,
	RadiationEquivalentDoseRateSievertPerSecond: true,
	RadiationEquivalentDoseRateRoentgenEquivalentManPerHour: true,
	RadiationEquivalentDoseRateNanosievertPerHour: true,
	RadiationEquivalentDoseRateMicrosievertPerHour: true,
	RadiationEquivalentDoseRateMillisievertPerHour: true,
	RadiationEquivalentDoseRateNanosievertPerSecond: true,
	RadiationEquivalentDoseRateMicrosievertPerSecond: true,
	RadiationEquivalentDoseRateMillisievertPerSecond: true,
	RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour: true,
}

// RadiationEquivalentDoseRateDto represents a RadiationEquivalentDoseRate measurement with a numerical value and its corresponding unit.
type RadiationEquivalentDoseRateDto struct {
    // Value is the numerical representation of the RadiationEquivalentDoseRate.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the RadiationEquivalentDoseRate, as defined in the RadiationEquivalentDoseRateUnits enumeration.
	Unit  RadiationEquivalentDoseRateUnits `json:"unit" validate:"required,oneof=SievertPerHour,SievertPerSecond,RoentgenEquivalentManPerHour,NanosievertPerHour,MicrosievertPerHour,MillisievertPerHour,NanosievertPerSecond,MicrosievertPerSecond,MillisievertPerSecond,MilliroentgenEquivalentManPerHour"`
}

// RadiationEquivalentDoseRateDtoFactory groups methods for creating and serializing RadiationEquivalentDoseRateDto objects.
type RadiationEquivalentDoseRateDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RadiationEquivalentDoseRateDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RadiationEquivalentDoseRateDtoFactory) FromJSON(data []byte) (*RadiationEquivalentDoseRateDto, error) {
	a := RadiationEquivalentDoseRateDto{}

    // Parse JSON into RadiationEquivalentDoseRateDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a RadiationEquivalentDoseRateDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RadiationEquivalentDoseRateDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RadiationEquivalentDoseRate represents a measurement in a of RadiationEquivalentDoseRate.
//
// A dose rate is quantity of radiation absorbed or delivered per unit time.
type RadiationEquivalentDoseRate struct {
	// value is the base measurement stored internally.
	value       float64
    
    sieverts_per_hourLazy *float64 
    sieverts_per_secondLazy *float64 
    roentgens_equivalent_man_per_hourLazy *float64 
    nanosieverts_per_hourLazy *float64 
    microsieverts_per_hourLazy *float64 
    millisieverts_per_hourLazy *float64 
    nanosieverts_per_secondLazy *float64 
    microsieverts_per_secondLazy *float64 
    millisieverts_per_secondLazy *float64 
    milliroentgens_equivalent_man_per_hourLazy *float64 
}

// RadiationEquivalentDoseRateFactory groups methods for creating RadiationEquivalentDoseRate instances.
type RadiationEquivalentDoseRateFactory struct{}

// CreateRadiationEquivalentDoseRate creates a new RadiationEquivalentDoseRate instance from the given value and unit.
func (uf RadiationEquivalentDoseRateFactory) CreateRadiationEquivalentDoseRate(value float64, unit RadiationEquivalentDoseRateUnits) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, unit)
}

// FromDto converts a RadiationEquivalentDoseRateDto to a RadiationEquivalentDoseRate instance.
func (uf RadiationEquivalentDoseRateFactory) FromDto(dto RadiationEquivalentDoseRateDto) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RadiationEquivalentDoseRate instance.
func (uf RadiationEquivalentDoseRateFactory) FromDtoJSON(data []byte) (*RadiationEquivalentDoseRate, error) {
	unitDto, err := RadiationEquivalentDoseRateDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RadiationEquivalentDoseRateDto from JSON: %w", err)
	}
	return RadiationEquivalentDoseRateFactory{}.FromDto(*unitDto)
}


// FromSievertsPerHour creates a new RadiationEquivalentDoseRate instance from a value in SievertsPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromSievertsPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateSievertPerHour)
}

// FromSievertsPerSecond creates a new RadiationEquivalentDoseRate instance from a value in SievertsPerSecond.
func (uf RadiationEquivalentDoseRateFactory) FromSievertsPerSecond(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateSievertPerSecond)
}

// FromRoentgensEquivalentManPerHour creates a new RadiationEquivalentDoseRate instance from a value in RoentgensEquivalentManPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromRoentgensEquivalentManPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
}

// FromNanosievertsPerHour creates a new RadiationEquivalentDoseRate instance from a value in NanosievertsPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromNanosievertsPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateNanosievertPerHour)
}

// FromMicrosievertsPerHour creates a new RadiationEquivalentDoseRate instance from a value in MicrosievertsPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromMicrosievertsPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateMicrosievertPerHour)
}

// FromMillisievertsPerHour creates a new RadiationEquivalentDoseRate instance from a value in MillisievertsPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromMillisievertsPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateMillisievertPerHour)
}

// FromNanosievertsPerSecond creates a new RadiationEquivalentDoseRate instance from a value in NanosievertsPerSecond.
func (uf RadiationEquivalentDoseRateFactory) FromNanosievertsPerSecond(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateNanosievertPerSecond)
}

// FromMicrosievertsPerSecond creates a new RadiationEquivalentDoseRate instance from a value in MicrosievertsPerSecond.
func (uf RadiationEquivalentDoseRateFactory) FromMicrosievertsPerSecond(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateMicrosievertPerSecond)
}

// FromMillisievertsPerSecond creates a new RadiationEquivalentDoseRate instance from a value in MillisievertsPerSecond.
func (uf RadiationEquivalentDoseRateFactory) FromMillisievertsPerSecond(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateMillisievertPerSecond)
}

// FromMilliroentgensEquivalentManPerHour creates a new RadiationEquivalentDoseRate instance from a value in MilliroentgensEquivalentManPerHour.
func (uf RadiationEquivalentDoseRateFactory) FromMilliroentgensEquivalentManPerHour(value float64) (*RadiationEquivalentDoseRate, error) {
	return newRadiationEquivalentDoseRate(value, RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
}


// newRadiationEquivalentDoseRate creates a new RadiationEquivalentDoseRate.
func newRadiationEquivalentDoseRate(value float64, fromUnit RadiationEquivalentDoseRateUnits) (*RadiationEquivalentDoseRate, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalRadiationEquivalentDoseRateUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in RadiationEquivalentDoseRateUnits", fromUnit)
	}
	a := &RadiationEquivalentDoseRate{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RadiationEquivalentDoseRate in SievertPerSecond unit (the base/default quantity).
func (a *RadiationEquivalentDoseRate) BaseValue() float64 {
	return a.value
}


// SievertsPerHour returns the RadiationEquivalentDoseRate value in SievertsPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) SievertsPerHour() float64 {
	if a.sieverts_per_hourLazy != nil {
		return *a.sieverts_per_hourLazy
	}
	sieverts_per_hour := a.convertFromBase(RadiationEquivalentDoseRateSievertPerHour)
	a.sieverts_per_hourLazy = &sieverts_per_hour
	return sieverts_per_hour
}

// SievertsPerSecond returns the RadiationEquivalentDoseRate value in SievertsPerSecond.
//
// 
func (a *RadiationEquivalentDoseRate) SievertsPerSecond() float64 {
	if a.sieverts_per_secondLazy != nil {
		return *a.sieverts_per_secondLazy
	}
	sieverts_per_second := a.convertFromBase(RadiationEquivalentDoseRateSievertPerSecond)
	a.sieverts_per_secondLazy = &sieverts_per_second
	return sieverts_per_second
}

// RoentgensEquivalentManPerHour returns the RadiationEquivalentDoseRate value in RoentgensEquivalentManPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) RoentgensEquivalentManPerHour() float64 {
	if a.roentgens_equivalent_man_per_hourLazy != nil {
		return *a.roentgens_equivalent_man_per_hourLazy
	}
	roentgens_equivalent_man_per_hour := a.convertFromBase(RadiationEquivalentDoseRateRoentgenEquivalentManPerHour)
	a.roentgens_equivalent_man_per_hourLazy = &roentgens_equivalent_man_per_hour
	return roentgens_equivalent_man_per_hour
}

// NanosievertsPerHour returns the RadiationEquivalentDoseRate value in NanosievertsPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) NanosievertsPerHour() float64 {
	if a.nanosieverts_per_hourLazy != nil {
		return *a.nanosieverts_per_hourLazy
	}
	nanosieverts_per_hour := a.convertFromBase(RadiationEquivalentDoseRateNanosievertPerHour)
	a.nanosieverts_per_hourLazy = &nanosieverts_per_hour
	return nanosieverts_per_hour
}

// MicrosievertsPerHour returns the RadiationEquivalentDoseRate value in MicrosievertsPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) MicrosievertsPerHour() float64 {
	if a.microsieverts_per_hourLazy != nil {
		return *a.microsieverts_per_hourLazy
	}
	microsieverts_per_hour := a.convertFromBase(RadiationEquivalentDoseRateMicrosievertPerHour)
	a.microsieverts_per_hourLazy = &microsieverts_per_hour
	return microsieverts_per_hour
}

// MillisievertsPerHour returns the RadiationEquivalentDoseRate value in MillisievertsPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) MillisievertsPerHour() float64 {
	if a.millisieverts_per_hourLazy != nil {
		return *a.millisieverts_per_hourLazy
	}
	millisieverts_per_hour := a.convertFromBase(RadiationEquivalentDoseRateMillisievertPerHour)
	a.millisieverts_per_hourLazy = &millisieverts_per_hour
	return millisieverts_per_hour
}

// NanosievertsPerSecond returns the RadiationEquivalentDoseRate value in NanosievertsPerSecond.
//
// 
func (a *RadiationEquivalentDoseRate) NanosievertsPerSecond() float64 {
	if a.nanosieverts_per_secondLazy != nil {
		return *a.nanosieverts_per_secondLazy
	}
	nanosieverts_per_second := a.convertFromBase(RadiationEquivalentDoseRateNanosievertPerSecond)
	a.nanosieverts_per_secondLazy = &nanosieverts_per_second
	return nanosieverts_per_second
}

// MicrosievertsPerSecond returns the RadiationEquivalentDoseRate value in MicrosievertsPerSecond.
//
// 
func (a *RadiationEquivalentDoseRate) MicrosievertsPerSecond() float64 {
	if a.microsieverts_per_secondLazy != nil {
		return *a.microsieverts_per_secondLazy
	}
	microsieverts_per_second := a.convertFromBase(RadiationEquivalentDoseRateMicrosievertPerSecond)
	a.microsieverts_per_secondLazy = &microsieverts_per_second
	return microsieverts_per_second
}

// MillisievertsPerSecond returns the RadiationEquivalentDoseRate value in MillisievertsPerSecond.
//
// 
func (a *RadiationEquivalentDoseRate) MillisievertsPerSecond() float64 {
	if a.millisieverts_per_secondLazy != nil {
		return *a.millisieverts_per_secondLazy
	}
	millisieverts_per_second := a.convertFromBase(RadiationEquivalentDoseRateMillisievertPerSecond)
	a.millisieverts_per_secondLazy = &millisieverts_per_second
	return millisieverts_per_second
}

// MilliroentgensEquivalentManPerHour returns the RadiationEquivalentDoseRate value in MilliroentgensEquivalentManPerHour.
//
// 
func (a *RadiationEquivalentDoseRate) MilliroentgensEquivalentManPerHour() float64 {
	if a.milliroentgens_equivalent_man_per_hourLazy != nil {
		return *a.milliroentgens_equivalent_man_per_hourLazy
	}
	milliroentgens_equivalent_man_per_hour := a.convertFromBase(RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour)
	a.milliroentgens_equivalent_man_per_hourLazy = &milliroentgens_equivalent_man_per_hour
	return milliroentgens_equivalent_man_per_hour
}


// ToDto creates a RadiationEquivalentDoseRateDto representation from the RadiationEquivalentDoseRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SievertPerSecond by default.
func (a *RadiationEquivalentDoseRate) ToDto(holdInUnit *RadiationEquivalentDoseRateUnits) RadiationEquivalentDoseRateDto {
	if holdInUnit == nil {
		defaultUnit := RadiationEquivalentDoseRateSievertPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return RadiationEquivalentDoseRateDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RadiationEquivalentDoseRate instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SievertPerSecond by default.
func (a *RadiationEquivalentDoseRate) ToDtoJSON(holdInUnit *RadiationEquivalentDoseRateUnits) ([]byte, error) {
	// Convert to RadiationEquivalentDoseRateDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RadiationEquivalentDoseRate to a specific unit value.
// The function uses the provided unit type (RadiationEquivalentDoseRateUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RadiationEquivalentDoseRate) Convert(toUnit RadiationEquivalentDoseRateUnits) float64 {
	switch toUnit { 
    case RadiationEquivalentDoseRateSievertPerHour:
		return a.SievertsPerHour()
    case RadiationEquivalentDoseRateSievertPerSecond:
		return a.SievertsPerSecond()
    case RadiationEquivalentDoseRateRoentgenEquivalentManPerHour:
		return a.RoentgensEquivalentManPerHour()
    case RadiationEquivalentDoseRateNanosievertPerHour:
		return a.NanosievertsPerHour()
    case RadiationEquivalentDoseRateMicrosievertPerHour:
		return a.MicrosievertsPerHour()
    case RadiationEquivalentDoseRateMillisievertPerHour:
		return a.MillisievertsPerHour()
    case RadiationEquivalentDoseRateNanosievertPerSecond:
		return a.NanosievertsPerSecond()
    case RadiationEquivalentDoseRateMicrosievertPerSecond:
		return a.MicrosievertsPerSecond()
    case RadiationEquivalentDoseRateMillisievertPerSecond:
		return a.MillisievertsPerSecond()
    case RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour:
		return a.MilliroentgensEquivalentManPerHour()
	default:
		return math.NaN()
	}
}

func (a *RadiationEquivalentDoseRate) convertFromBase(toUnit RadiationEquivalentDoseRateUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadiationEquivalentDoseRateSievertPerHour:
		return (value*3600) 
	case RadiationEquivalentDoseRateSievertPerSecond:
		return (value) 
	case RadiationEquivalentDoseRateRoentgenEquivalentManPerHour:
		return (value * 100 * 3600) 
	case RadiationEquivalentDoseRateNanosievertPerHour:
		return ((value*3600) / 1e-09) 
	case RadiationEquivalentDoseRateMicrosievertPerHour:
		return ((value*3600) / 1e-06) 
	case RadiationEquivalentDoseRateMillisievertPerHour:
		return ((value*3600) / 0.001) 
	case RadiationEquivalentDoseRateNanosievertPerSecond:
		return ((value) / 1e-09) 
	case RadiationEquivalentDoseRateMicrosievertPerSecond:
		return ((value) / 1e-06) 
	case RadiationEquivalentDoseRateMillisievertPerSecond:
		return ((value) / 0.001) 
	case RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour:
		return ((value * 100 * 3600) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RadiationEquivalentDoseRate) convertToBase(value float64, fromUnit RadiationEquivalentDoseRateUnits) float64 {
	switch fromUnit { 
	case RadiationEquivalentDoseRateSievertPerHour:
		return (value/3600) 
	case RadiationEquivalentDoseRateSievertPerSecond:
		return (value) 
	case RadiationEquivalentDoseRateRoentgenEquivalentManPerHour:
		return (value / 100 / 3600) 
	case RadiationEquivalentDoseRateNanosievertPerHour:
		return ((value/3600) * 1e-09) 
	case RadiationEquivalentDoseRateMicrosievertPerHour:
		return ((value/3600) * 1e-06) 
	case RadiationEquivalentDoseRateMillisievertPerHour:
		return ((value/3600) * 0.001) 
	case RadiationEquivalentDoseRateNanosievertPerSecond:
		return ((value) * 1e-09) 
	case RadiationEquivalentDoseRateMicrosievertPerSecond:
		return ((value) * 1e-06) 
	case RadiationEquivalentDoseRateMillisievertPerSecond:
		return ((value) * 0.001) 
	case RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour:
		return ((value / 100 / 3600) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RadiationEquivalentDoseRate in the default unit (SievertPerSecond),
// formatted to two decimal places.
func (a RadiationEquivalentDoseRate) String() string {
	return a.ToString(RadiationEquivalentDoseRateSievertPerSecond, 2)
}

// ToString formats the RadiationEquivalentDoseRate value as a string with the specified unit and fractional digits.
// It converts the RadiationEquivalentDoseRate to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RadiationEquivalentDoseRate value will be converted (e.g., SievertPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RadiationEquivalentDoseRate with the unit abbreviation.
func (a *RadiationEquivalentDoseRate) ToString(unit RadiationEquivalentDoseRateUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRadiationEquivalentDoseRateAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRadiationEquivalentDoseRateAbbreviation(unit))
}

// Equals checks if the given RadiationEquivalentDoseRate is equal to the current RadiationEquivalentDoseRate.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to compare against.
//
// Returns:
//    bool: Returns true if both RadiationEquivalentDoseRate are equal, false otherwise.
func (a *RadiationEquivalentDoseRate) Equals(other *RadiationEquivalentDoseRate) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RadiationEquivalentDoseRate with another RadiationEquivalentDoseRate.
// It returns -1 if the current RadiationEquivalentDoseRate is less than the other RadiationEquivalentDoseRate, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to compare against.
//
// Returns:
//    int: -1 if the current RadiationEquivalentDoseRate is less, 1 if greater, and 0 if equal.
func (a *RadiationEquivalentDoseRate) CompareTo(other *RadiationEquivalentDoseRate) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RadiationEquivalentDoseRate to the current RadiationEquivalentDoseRate and returns the result.
// The result is a new RadiationEquivalentDoseRate instance with the sum of the values.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to add to the current RadiationEquivalentDoseRate.
//
// Returns:
//    *RadiationEquivalentDoseRate: A new RadiationEquivalentDoseRate instance representing the sum of both RadiationEquivalentDoseRate.
func (a *RadiationEquivalentDoseRate) Add(other *RadiationEquivalentDoseRate) *RadiationEquivalentDoseRate {
	return &RadiationEquivalentDoseRate{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RadiationEquivalentDoseRate from the current RadiationEquivalentDoseRate and returns the result.
// The result is a new RadiationEquivalentDoseRate instance with the difference of the values.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to subtract from the current RadiationEquivalentDoseRate.
//
// Returns:
//    *RadiationEquivalentDoseRate: A new RadiationEquivalentDoseRate instance representing the difference of both RadiationEquivalentDoseRate.
func (a *RadiationEquivalentDoseRate) Subtract(other *RadiationEquivalentDoseRate) *RadiationEquivalentDoseRate {
	return &RadiationEquivalentDoseRate{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RadiationEquivalentDoseRate by the given RadiationEquivalentDoseRate and returns the result.
// The result is a new RadiationEquivalentDoseRate instance with the product of the values.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to multiply with the current RadiationEquivalentDoseRate.
//
// Returns:
//    *RadiationEquivalentDoseRate: A new RadiationEquivalentDoseRate instance representing the product of both RadiationEquivalentDoseRate.
func (a *RadiationEquivalentDoseRate) Multiply(other *RadiationEquivalentDoseRate) *RadiationEquivalentDoseRate {
	return &RadiationEquivalentDoseRate{value: a.value * other.BaseValue()}
}

// Divide divides the current RadiationEquivalentDoseRate by the given RadiationEquivalentDoseRate and returns the result.
// The result is a new RadiationEquivalentDoseRate instance with the quotient of the values.
//
// Parameters:
//    other: The RadiationEquivalentDoseRate to divide the current RadiationEquivalentDoseRate by.
//
// Returns:
//    *RadiationEquivalentDoseRate: A new RadiationEquivalentDoseRate instance representing the quotient of both RadiationEquivalentDoseRate.
func (a *RadiationEquivalentDoseRate) Divide(other *RadiationEquivalentDoseRate) *RadiationEquivalentDoseRate {
	return &RadiationEquivalentDoseRate{value: a.value / other.BaseValue()}
}

// GetRadiationEquivalentDoseRateAbbreviation gets the unit abbreviation.
func GetRadiationEquivalentDoseRateAbbreviation(unit RadiationEquivalentDoseRateUnits) string {
	switch unit { 
	case RadiationEquivalentDoseRateSievertPerHour:
		return "Sv/h" 
	case RadiationEquivalentDoseRateSievertPerSecond:
		return "Sv/s" 
	case RadiationEquivalentDoseRateRoentgenEquivalentManPerHour:
		return "rem/h" 
	case RadiationEquivalentDoseRateNanosievertPerHour:
		return "nSv/h" 
	case RadiationEquivalentDoseRateMicrosievertPerHour:
		return "μSv/h" 
	case RadiationEquivalentDoseRateMillisievertPerHour:
		return "mSv/h" 
	case RadiationEquivalentDoseRateNanosievertPerSecond:
		return "nSv/s" 
	case RadiationEquivalentDoseRateMicrosievertPerSecond:
		return "μSv/s" 
	case RadiationEquivalentDoseRateMillisievertPerSecond:
		return "mSv/s" 
	case RadiationEquivalentDoseRateMilliroentgenEquivalentManPerHour:
		return "mrem/h" 
	default:
		return ""
	}
}