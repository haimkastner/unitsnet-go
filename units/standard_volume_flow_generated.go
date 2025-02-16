// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// StandardVolumeFlowUnits defines various units of StandardVolumeFlow.
type StandardVolumeFlowUnits string

const (
    
        // 
        StandardVolumeFlowStandardCubicMeterPerSecond StandardVolumeFlowUnits = "StandardCubicMeterPerSecond"
        // 
        StandardVolumeFlowStandardCubicMeterPerMinute StandardVolumeFlowUnits = "StandardCubicMeterPerMinute"
        // 
        StandardVolumeFlowStandardCubicMeterPerHour StandardVolumeFlowUnits = "StandardCubicMeterPerHour"
        // 
        StandardVolumeFlowStandardCubicMeterPerDay StandardVolumeFlowUnits = "StandardCubicMeterPerDay"
        // 
        StandardVolumeFlowStandardCubicCentimeterPerMinute StandardVolumeFlowUnits = "StandardCubicCentimeterPerMinute"
        // 
        StandardVolumeFlowStandardLiterPerMinute StandardVolumeFlowUnits = "StandardLiterPerMinute"
        // 
        StandardVolumeFlowStandardCubicFootPerSecond StandardVolumeFlowUnits = "StandardCubicFootPerSecond"
        // 
        StandardVolumeFlowStandardCubicFootPerMinute StandardVolumeFlowUnits = "StandardCubicFootPerMinute"
        // 
        StandardVolumeFlowStandardCubicFootPerHour StandardVolumeFlowUnits = "StandardCubicFootPerHour"
)

// StandardVolumeFlowDto represents a StandardVolumeFlow measurement with a numerical value and its corresponding unit.
type StandardVolumeFlowDto struct {
    // Value is the numerical representation of the StandardVolumeFlow.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the StandardVolumeFlow, as defined in the StandardVolumeFlowUnits enumeration.
	Unit  StandardVolumeFlowUnits `json:"unit"`
}

// StandardVolumeFlowDtoFactory groups methods for creating and serializing StandardVolumeFlowDto objects.
type StandardVolumeFlowDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a StandardVolumeFlowDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf StandardVolumeFlowDtoFactory) FromJSON(data []byte) (*StandardVolumeFlowDto, error) {
	a := StandardVolumeFlowDto{}

    // Parse JSON into StandardVolumeFlowDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a StandardVolumeFlowDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a StandardVolumeFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// StandardVolumeFlow represents a measurement in a of StandardVolumeFlow.
//
// The molar flow rate of a gas corrected to standardized conditions of temperature and pressure thus representing a fixed number of moles of gas regardless of composition and actual flow conditions.
type StandardVolumeFlow struct {
	// value is the base measurement stored internally.
	value       float64
    
    standard_cubic_meters_per_secondLazy *float64 
    standard_cubic_meters_per_minuteLazy *float64 
    standard_cubic_meters_per_hourLazy *float64 
    standard_cubic_meters_per_dayLazy *float64 
    standard_cubic_centimeters_per_minuteLazy *float64 
    standard_liters_per_minuteLazy *float64 
    standard_cubic_feet_per_secondLazy *float64 
    standard_cubic_feet_per_minuteLazy *float64 
    standard_cubic_feet_per_hourLazy *float64 
}

// StandardVolumeFlowFactory groups methods for creating StandardVolumeFlow instances.
type StandardVolumeFlowFactory struct{}

// CreateStandardVolumeFlow creates a new StandardVolumeFlow instance from the given value and unit.
func (uf StandardVolumeFlowFactory) CreateStandardVolumeFlow(value float64, unit StandardVolumeFlowUnits) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, unit)
}

// FromDto converts a StandardVolumeFlowDto to a StandardVolumeFlow instance.
func (uf StandardVolumeFlowFactory) FromDto(dto StandardVolumeFlowDto) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a StandardVolumeFlow instance.
func (uf StandardVolumeFlowFactory) FromDtoJSON(data []byte) (*StandardVolumeFlow, error) {
	unitDto, err := StandardVolumeFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse StandardVolumeFlowDto from JSON: %w", err)
	}
	return StandardVolumeFlowFactory{}.FromDto(*unitDto)
}


// FromStandardCubicMetersPerSecond creates a new StandardVolumeFlow instance from a value in StandardCubicMetersPerSecond.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerSecond(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerSecond)
}

// FromStandardCubicMetersPerMinute creates a new StandardVolumeFlow instance from a value in StandardCubicMetersPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerMinute)
}

// FromStandardCubicMetersPerHour creates a new StandardVolumeFlow instance from a value in StandardCubicMetersPerHour.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerHour(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerHour)
}

// FromStandardCubicMetersPerDay creates a new StandardVolumeFlow instance from a value in StandardCubicMetersPerDay.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerDay(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerDay)
}

// FromStandardCubicCentimetersPerMinute creates a new StandardVolumeFlow instance from a value in StandardCubicCentimetersPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicCentimetersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicCentimeterPerMinute)
}

// FromStandardLitersPerMinute creates a new StandardVolumeFlow instance from a value in StandardLitersPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardLitersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardLiterPerMinute)
}

// FromStandardCubicFeetPerSecond creates a new StandardVolumeFlow instance from a value in StandardCubicFeetPerSecond.
func (uf StandardVolumeFlowFactory) FromStandardCubicFeetPerSecond(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicFootPerSecond)
}

// FromStandardCubicFeetPerMinute creates a new StandardVolumeFlow instance from a value in StandardCubicFeetPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicFeetPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicFootPerMinute)
}

// FromStandardCubicFeetPerHour creates a new StandardVolumeFlow instance from a value in StandardCubicFeetPerHour.
func (uf StandardVolumeFlowFactory) FromStandardCubicFeetPerHour(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicFootPerHour)
}


// newStandardVolumeFlow creates a new StandardVolumeFlow.
func newStandardVolumeFlow(value float64, fromUnit StandardVolumeFlowUnits) (*StandardVolumeFlow, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &StandardVolumeFlow{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of StandardVolumeFlow in StandardCubicMeterPerSecond unit (the base/default quantity).
func (a *StandardVolumeFlow) BaseValue() float64 {
	return a.value
}


// StandardCubicMetersPerSecond returns the StandardVolumeFlow value in StandardCubicMetersPerSecond.
//
// 
func (a *StandardVolumeFlow) StandardCubicMetersPerSecond() float64 {
	if a.standard_cubic_meters_per_secondLazy != nil {
		return *a.standard_cubic_meters_per_secondLazy
	}
	standard_cubic_meters_per_second := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerSecond)
	a.standard_cubic_meters_per_secondLazy = &standard_cubic_meters_per_second
	return standard_cubic_meters_per_second
}

// StandardCubicMetersPerMinute returns the StandardVolumeFlow value in StandardCubicMetersPerMinute.
//
// 
func (a *StandardVolumeFlow) StandardCubicMetersPerMinute() float64 {
	if a.standard_cubic_meters_per_minuteLazy != nil {
		return *a.standard_cubic_meters_per_minuteLazy
	}
	standard_cubic_meters_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerMinute)
	a.standard_cubic_meters_per_minuteLazy = &standard_cubic_meters_per_minute
	return standard_cubic_meters_per_minute
}

// StandardCubicMetersPerHour returns the StandardVolumeFlow value in StandardCubicMetersPerHour.
//
// 
func (a *StandardVolumeFlow) StandardCubicMetersPerHour() float64 {
	if a.standard_cubic_meters_per_hourLazy != nil {
		return *a.standard_cubic_meters_per_hourLazy
	}
	standard_cubic_meters_per_hour := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerHour)
	a.standard_cubic_meters_per_hourLazy = &standard_cubic_meters_per_hour
	return standard_cubic_meters_per_hour
}

// StandardCubicMetersPerDay returns the StandardVolumeFlow value in StandardCubicMetersPerDay.
//
// 
func (a *StandardVolumeFlow) StandardCubicMetersPerDay() float64 {
	if a.standard_cubic_meters_per_dayLazy != nil {
		return *a.standard_cubic_meters_per_dayLazy
	}
	standard_cubic_meters_per_day := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerDay)
	a.standard_cubic_meters_per_dayLazy = &standard_cubic_meters_per_day
	return standard_cubic_meters_per_day
}

// StandardCubicCentimetersPerMinute returns the StandardVolumeFlow value in StandardCubicCentimetersPerMinute.
//
// 
func (a *StandardVolumeFlow) StandardCubicCentimetersPerMinute() float64 {
	if a.standard_cubic_centimeters_per_minuteLazy != nil {
		return *a.standard_cubic_centimeters_per_minuteLazy
	}
	standard_cubic_centimeters_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicCentimeterPerMinute)
	a.standard_cubic_centimeters_per_minuteLazy = &standard_cubic_centimeters_per_minute
	return standard_cubic_centimeters_per_minute
}

// StandardLitersPerMinute returns the StandardVolumeFlow value in StandardLitersPerMinute.
//
// 
func (a *StandardVolumeFlow) StandardLitersPerMinute() float64 {
	if a.standard_liters_per_minuteLazy != nil {
		return *a.standard_liters_per_minuteLazy
	}
	standard_liters_per_minute := a.convertFromBase(StandardVolumeFlowStandardLiterPerMinute)
	a.standard_liters_per_minuteLazy = &standard_liters_per_minute
	return standard_liters_per_minute
}

// StandardCubicFeetPerSecond returns the StandardVolumeFlow value in StandardCubicFeetPerSecond.
//
// 
func (a *StandardVolumeFlow) StandardCubicFeetPerSecond() float64 {
	if a.standard_cubic_feet_per_secondLazy != nil {
		return *a.standard_cubic_feet_per_secondLazy
	}
	standard_cubic_feet_per_second := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerSecond)
	a.standard_cubic_feet_per_secondLazy = &standard_cubic_feet_per_second
	return standard_cubic_feet_per_second
}

// StandardCubicFeetPerMinute returns the StandardVolumeFlow value in StandardCubicFeetPerMinute.
//
// 
func (a *StandardVolumeFlow) StandardCubicFeetPerMinute() float64 {
	if a.standard_cubic_feet_per_minuteLazy != nil {
		return *a.standard_cubic_feet_per_minuteLazy
	}
	standard_cubic_feet_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerMinute)
	a.standard_cubic_feet_per_minuteLazy = &standard_cubic_feet_per_minute
	return standard_cubic_feet_per_minute
}

// StandardCubicFeetPerHour returns the StandardVolumeFlow value in StandardCubicFeetPerHour.
//
// 
func (a *StandardVolumeFlow) StandardCubicFeetPerHour() float64 {
	if a.standard_cubic_feet_per_hourLazy != nil {
		return *a.standard_cubic_feet_per_hourLazy
	}
	standard_cubic_feet_per_hour := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerHour)
	a.standard_cubic_feet_per_hourLazy = &standard_cubic_feet_per_hour
	return standard_cubic_feet_per_hour
}


// ToDto creates a StandardVolumeFlowDto representation from the StandardVolumeFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by StandardCubicMeterPerSecond by default.
func (a *StandardVolumeFlow) ToDto(holdInUnit *StandardVolumeFlowUnits) StandardVolumeFlowDto {
	if holdInUnit == nil {
		defaultUnit := StandardVolumeFlowStandardCubicMeterPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return StandardVolumeFlowDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the StandardVolumeFlow instance.
//
// If the provided holdInUnit is nil, the value will be repesented by StandardCubicMeterPerSecond by default.
func (a *StandardVolumeFlow) ToDtoJSON(holdInUnit *StandardVolumeFlowUnits) ([]byte, error) {
	// Convert to StandardVolumeFlowDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a StandardVolumeFlow to a specific unit value.
// The function uses the provided unit type (StandardVolumeFlowUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *StandardVolumeFlow) Convert(toUnit StandardVolumeFlowUnits) float64 {
	switch toUnit { 
    case StandardVolumeFlowStandardCubicMeterPerSecond:
		return a.StandardCubicMetersPerSecond()
    case StandardVolumeFlowStandardCubicMeterPerMinute:
		return a.StandardCubicMetersPerMinute()
    case StandardVolumeFlowStandardCubicMeterPerHour:
		return a.StandardCubicMetersPerHour()
    case StandardVolumeFlowStandardCubicMeterPerDay:
		return a.StandardCubicMetersPerDay()
    case StandardVolumeFlowStandardCubicCentimeterPerMinute:
		return a.StandardCubicCentimetersPerMinute()
    case StandardVolumeFlowStandardLiterPerMinute:
		return a.StandardLitersPerMinute()
    case StandardVolumeFlowStandardCubicFootPerSecond:
		return a.StandardCubicFeetPerSecond()
    case StandardVolumeFlowStandardCubicFootPerMinute:
		return a.StandardCubicFeetPerMinute()
    case StandardVolumeFlowStandardCubicFootPerHour:
		return a.StandardCubicFeetPerHour()
	default:
		return math.NaN()
	}
}

func (a *StandardVolumeFlow) convertFromBase(toUnit StandardVolumeFlowUnits) float64 {
    value := a.value
	switch toUnit { 
	case StandardVolumeFlowStandardCubicMeterPerSecond:
		return (value) 
	case StandardVolumeFlowStandardCubicMeterPerMinute:
		return (value * 60) 
	case StandardVolumeFlowStandardCubicMeterPerHour:
		return (value * 3600) 
	case StandardVolumeFlowStandardCubicMeterPerDay:
		return (value * 86400) 
	case StandardVolumeFlowStandardCubicCentimeterPerMinute:
		return (value * 6e7) 
	case StandardVolumeFlowStandardLiterPerMinute:
		return (value * 60000) 
	case StandardVolumeFlowStandardCubicFootPerSecond:
		return (value * 35.314666721) 
	case StandardVolumeFlowStandardCubicFootPerMinute:
		return (value * 2118.88000326) 
	case StandardVolumeFlowStandardCubicFootPerHour:
		return (value / 7.8657907199999087346816086183876e-6) 
	default:
		return math.NaN()
	}
}

func (a *StandardVolumeFlow) convertToBase(value float64, fromUnit StandardVolumeFlowUnits) float64 {
	switch fromUnit { 
	case StandardVolumeFlowStandardCubicMeterPerSecond:
		return (value) 
	case StandardVolumeFlowStandardCubicMeterPerMinute:
		return (value / 60) 
	case StandardVolumeFlowStandardCubicMeterPerHour:
		return (value / 3600) 
	case StandardVolumeFlowStandardCubicMeterPerDay:
		return (value / 86400) 
	case StandardVolumeFlowStandardCubicCentimeterPerMinute:
		return (value / 6e7) 
	case StandardVolumeFlowStandardLiterPerMinute:
		return (value / 60000) 
	case StandardVolumeFlowStandardCubicFootPerSecond:
		return (value / 35.314666721) 
	case StandardVolumeFlowStandardCubicFootPerMinute:
		return (value / 2118.88000326) 
	case StandardVolumeFlowStandardCubicFootPerHour:
		return (value * 7.8657907199999087346816086183876e-6) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the StandardVolumeFlow in the default unit (StandardCubicMeterPerSecond),
// formatted to two decimal places.
func (a StandardVolumeFlow) String() string {
	return a.ToString(StandardVolumeFlowStandardCubicMeterPerSecond, 2)
}

// ToString formats the StandardVolumeFlow value as a string with the specified unit and fractional digits.
// It converts the StandardVolumeFlow to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the StandardVolumeFlow value will be converted (e.g., StandardCubicMeterPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the StandardVolumeFlow with the unit abbreviation.
func (a *StandardVolumeFlow) ToString(unit StandardVolumeFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetStandardVolumeFlowAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetStandardVolumeFlowAbbreviation(unit))
}

// Equals checks if the given StandardVolumeFlow is equal to the current StandardVolumeFlow.
//
// Parameters:
//    other: The StandardVolumeFlow to compare against.
//
// Returns:
//    bool: Returns true if both StandardVolumeFlow are equal, false otherwise.
func (a *StandardVolumeFlow) Equals(other *StandardVolumeFlow) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current StandardVolumeFlow with another StandardVolumeFlow.
// It returns -1 if the current StandardVolumeFlow is less than the other StandardVolumeFlow, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The StandardVolumeFlow to compare against.
//
// Returns:
//    int: -1 if the current StandardVolumeFlow is less, 1 if greater, and 0 if equal.
func (a *StandardVolumeFlow) CompareTo(other *StandardVolumeFlow) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given StandardVolumeFlow to the current StandardVolumeFlow and returns the result.
// The result is a new StandardVolumeFlow instance with the sum of the values.
//
// Parameters:
//    other: The StandardVolumeFlow to add to the current StandardVolumeFlow.
//
// Returns:
//    *StandardVolumeFlow: A new StandardVolumeFlow instance representing the sum of both StandardVolumeFlow.
func (a *StandardVolumeFlow) Add(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given StandardVolumeFlow from the current StandardVolumeFlow and returns the result.
// The result is a new StandardVolumeFlow instance with the difference of the values.
//
// Parameters:
//    other: The StandardVolumeFlow to subtract from the current StandardVolumeFlow.
//
// Returns:
//    *StandardVolumeFlow: A new StandardVolumeFlow instance representing the difference of both StandardVolumeFlow.
func (a *StandardVolumeFlow) Subtract(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current StandardVolumeFlow by the given StandardVolumeFlow and returns the result.
// The result is a new StandardVolumeFlow instance with the product of the values.
//
// Parameters:
//    other: The StandardVolumeFlow to multiply with the current StandardVolumeFlow.
//
// Returns:
//    *StandardVolumeFlow: A new StandardVolumeFlow instance representing the product of both StandardVolumeFlow.
func (a *StandardVolumeFlow) Multiply(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value * other.BaseValue()}
}

// Divide divides the current StandardVolumeFlow by the given StandardVolumeFlow and returns the result.
// The result is a new StandardVolumeFlow instance with the quotient of the values.
//
// Parameters:
//    other: The StandardVolumeFlow to divide the current StandardVolumeFlow by.
//
// Returns:
//    *StandardVolumeFlow: A new StandardVolumeFlow instance representing the quotient of both StandardVolumeFlow.
func (a *StandardVolumeFlow) Divide(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value / other.BaseValue()}
}

// GetStandardVolumeFlowAbbreviation gets the unit abbreviation.
func GetStandardVolumeFlowAbbreviation(unit StandardVolumeFlowUnits) string {
	switch unit { 
	case StandardVolumeFlowStandardCubicMeterPerSecond:
		return "Sm³/s" 
	case StandardVolumeFlowStandardCubicMeterPerMinute:
		return "Sm³/min" 
	case StandardVolumeFlowStandardCubicMeterPerHour:
		return "Sm³/h" 
	case StandardVolumeFlowStandardCubicMeterPerDay:
		return "Sm³/d" 
	case StandardVolumeFlowStandardCubicCentimeterPerMinute:
		return "sccm" 
	case StandardVolumeFlowStandardLiterPerMinute:
		return "slm" 
	case StandardVolumeFlowStandardCubicFootPerSecond:
		return "Sft³/s" 
	case StandardVolumeFlowStandardCubicFootPerMinute:
		return "scfm" 
	case StandardVolumeFlowStandardCubicFootPerHour:
		return "scfh" 
	default:
		return ""
	}
}