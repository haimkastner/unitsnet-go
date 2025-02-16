// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// DurationUnits defines various units of Duration.
type DurationUnits string

const (
    
        // 
        DurationYear365 DurationUnits = "Year365"
        // 
        DurationMonth30 DurationUnits = "Month30"
        // 
        DurationWeek DurationUnits = "Week"
        // 
        DurationDay DurationUnits = "Day"
        // 
        DurationHour DurationUnits = "Hour"
        // 
        DurationMinute DurationUnits = "Minute"
        // 
        DurationSecond DurationUnits = "Second"
        // 
        DurationJulianYear DurationUnits = "JulianYear"
        // 
        DurationSol DurationUnits = "Sol"
        // 
        DurationNanosecond DurationUnits = "Nanosecond"
        // 
        DurationMicrosecond DurationUnits = "Microsecond"
        // 
        DurationMillisecond DurationUnits = "Millisecond"
)

// DurationDto represents a Duration measurement with a numerical value and its corresponding unit.
type DurationDto struct {
    // Value is the numerical representation of the Duration.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Duration, as defined in the DurationUnits enumeration.
	Unit  DurationUnits `json:"unit"`
}

// DurationDtoFactory groups methods for creating and serializing DurationDto objects.
type DurationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a DurationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf DurationDtoFactory) FromJSON(data []byte) (*DurationDto, error) {
	a := DurationDto{}

    // Parse JSON into DurationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a DurationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a DurationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Duration represents a measurement in a of Duration.
//
// Time is a dimension in which events can be ordered from the past through the present into the future, and also the measure of durations of events and the intervals between them.
type Duration struct {
	// value is the base measurement stored internally.
	value       float64
    
    years365Lazy *float64 
    months30Lazy *float64 
    weeksLazy *float64 
    daysLazy *float64 
    hoursLazy *float64 
    minutesLazy *float64 
    secondsLazy *float64 
    julian_yearsLazy *float64 
    solsLazy *float64 
    nanosecondsLazy *float64 
    microsecondsLazy *float64 
    millisecondsLazy *float64 
}

// DurationFactory groups methods for creating Duration instances.
type DurationFactory struct{}

// CreateDuration creates a new Duration instance from the given value and unit.
func (uf DurationFactory) CreateDuration(value float64, unit DurationUnits) (*Duration, error) {
	return newDuration(value, unit)
}

// FromDto converts a DurationDto to a Duration instance.
func (uf DurationFactory) FromDto(dto DurationDto) (*Duration, error) {
	return newDuration(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Duration instance.
func (uf DurationFactory) FromDtoJSON(data []byte) (*Duration, error) {
	unitDto, err := DurationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DurationDto from JSON: %w", err)
	}
	return DurationFactory{}.FromDto(*unitDto)
}


// FromYears365 creates a new Duration instance from a value in Years365.
func (uf DurationFactory) FromYears365(value float64) (*Duration, error) {
	return newDuration(value, DurationYear365)
}

// FromMonths30 creates a new Duration instance from a value in Months30.
func (uf DurationFactory) FromMonths30(value float64) (*Duration, error) {
	return newDuration(value, DurationMonth30)
}

// FromWeeks creates a new Duration instance from a value in Weeks.
func (uf DurationFactory) FromWeeks(value float64) (*Duration, error) {
	return newDuration(value, DurationWeek)
}

// FromDays creates a new Duration instance from a value in Days.
func (uf DurationFactory) FromDays(value float64) (*Duration, error) {
	return newDuration(value, DurationDay)
}

// FromHours creates a new Duration instance from a value in Hours.
func (uf DurationFactory) FromHours(value float64) (*Duration, error) {
	return newDuration(value, DurationHour)
}

// FromMinutes creates a new Duration instance from a value in Minutes.
func (uf DurationFactory) FromMinutes(value float64) (*Duration, error) {
	return newDuration(value, DurationMinute)
}

// FromSeconds creates a new Duration instance from a value in Seconds.
func (uf DurationFactory) FromSeconds(value float64) (*Duration, error) {
	return newDuration(value, DurationSecond)
}

// FromJulianYears creates a new Duration instance from a value in JulianYears.
func (uf DurationFactory) FromJulianYears(value float64) (*Duration, error) {
	return newDuration(value, DurationJulianYear)
}

// FromSols creates a new Duration instance from a value in Sols.
func (uf DurationFactory) FromSols(value float64) (*Duration, error) {
	return newDuration(value, DurationSol)
}

// FromNanoseconds creates a new Duration instance from a value in Nanoseconds.
func (uf DurationFactory) FromNanoseconds(value float64) (*Duration, error) {
	return newDuration(value, DurationNanosecond)
}

// FromMicroseconds creates a new Duration instance from a value in Microseconds.
func (uf DurationFactory) FromMicroseconds(value float64) (*Duration, error) {
	return newDuration(value, DurationMicrosecond)
}

// FromMilliseconds creates a new Duration instance from a value in Milliseconds.
func (uf DurationFactory) FromMilliseconds(value float64) (*Duration, error) {
	return newDuration(value, DurationMillisecond)
}


// newDuration creates a new Duration.
func newDuration(value float64, fromUnit DurationUnits) (*Duration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Duration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Duration in Second unit (the base/default quantity).
func (a *Duration) BaseValue() float64 {
	return a.value
}


// Years365 returns the Duration value in Years365.
//
// 
func (a *Duration) Years365() float64 {
	if a.years365Lazy != nil {
		return *a.years365Lazy
	}
	years365 := a.convertFromBase(DurationYear365)
	a.years365Lazy = &years365
	return years365
}

// Months30 returns the Duration value in Months30.
//
// 
func (a *Duration) Months30() float64 {
	if a.months30Lazy != nil {
		return *a.months30Lazy
	}
	months30 := a.convertFromBase(DurationMonth30)
	a.months30Lazy = &months30
	return months30
}

// Weeks returns the Duration value in Weeks.
//
// 
func (a *Duration) Weeks() float64 {
	if a.weeksLazy != nil {
		return *a.weeksLazy
	}
	weeks := a.convertFromBase(DurationWeek)
	a.weeksLazy = &weeks
	return weeks
}

// Days returns the Duration value in Days.
//
// 
func (a *Duration) Days() float64 {
	if a.daysLazy != nil {
		return *a.daysLazy
	}
	days := a.convertFromBase(DurationDay)
	a.daysLazy = &days
	return days
}

// Hours returns the Duration value in Hours.
//
// 
func (a *Duration) Hours() float64 {
	if a.hoursLazy != nil {
		return *a.hoursLazy
	}
	hours := a.convertFromBase(DurationHour)
	a.hoursLazy = &hours
	return hours
}

// Minutes returns the Duration value in Minutes.
//
// 
func (a *Duration) Minutes() float64 {
	if a.minutesLazy != nil {
		return *a.minutesLazy
	}
	minutes := a.convertFromBase(DurationMinute)
	a.minutesLazy = &minutes
	return minutes
}

// Seconds returns the Duration value in Seconds.
//
// 
func (a *Duration) Seconds() float64 {
	if a.secondsLazy != nil {
		return *a.secondsLazy
	}
	seconds := a.convertFromBase(DurationSecond)
	a.secondsLazy = &seconds
	return seconds
}

// JulianYears returns the Duration value in JulianYears.
//
// 
func (a *Duration) JulianYears() float64 {
	if a.julian_yearsLazy != nil {
		return *a.julian_yearsLazy
	}
	julian_years := a.convertFromBase(DurationJulianYear)
	a.julian_yearsLazy = &julian_years
	return julian_years
}

// Sols returns the Duration value in Sols.
//
// 
func (a *Duration) Sols() float64 {
	if a.solsLazy != nil {
		return *a.solsLazy
	}
	sols := a.convertFromBase(DurationSol)
	a.solsLazy = &sols
	return sols
}

// Nanoseconds returns the Duration value in Nanoseconds.
//
// 
func (a *Duration) Nanoseconds() float64 {
	if a.nanosecondsLazy != nil {
		return *a.nanosecondsLazy
	}
	nanoseconds := a.convertFromBase(DurationNanosecond)
	a.nanosecondsLazy = &nanoseconds
	return nanoseconds
}

// Microseconds returns the Duration value in Microseconds.
//
// 
func (a *Duration) Microseconds() float64 {
	if a.microsecondsLazy != nil {
		return *a.microsecondsLazy
	}
	microseconds := a.convertFromBase(DurationMicrosecond)
	a.microsecondsLazy = &microseconds
	return microseconds
}

// Milliseconds returns the Duration value in Milliseconds.
//
// 
func (a *Duration) Milliseconds() float64 {
	if a.millisecondsLazy != nil {
		return *a.millisecondsLazy
	}
	milliseconds := a.convertFromBase(DurationMillisecond)
	a.millisecondsLazy = &milliseconds
	return milliseconds
}


// ToDto creates a DurationDto representation from the Duration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Second by default.
func (a *Duration) ToDto(holdInUnit *DurationUnits) DurationDto {
	if holdInUnit == nil {
		defaultUnit := DurationSecond // Default value
		holdInUnit = &defaultUnit
	}

	return DurationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Duration instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Second by default.
func (a *Duration) ToDtoJSON(holdInUnit *DurationUnits) ([]byte, error) {
	// Convert to DurationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Duration to a specific unit value.
// The function uses the provided unit type (DurationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Duration) Convert(toUnit DurationUnits) float64 {
	switch toUnit { 
    case DurationYear365:
		return a.Years365()
    case DurationMonth30:
		return a.Months30()
    case DurationWeek:
		return a.Weeks()
    case DurationDay:
		return a.Days()
    case DurationHour:
		return a.Hours()
    case DurationMinute:
		return a.Minutes()
    case DurationSecond:
		return a.Seconds()
    case DurationJulianYear:
		return a.JulianYears()
    case DurationSol:
		return a.Sols()
    case DurationNanosecond:
		return a.Nanoseconds()
    case DurationMicrosecond:
		return a.Microseconds()
    case DurationMillisecond:
		return a.Milliseconds()
	default:
		return math.NaN()
	}
}

func (a *Duration) convertFromBase(toUnit DurationUnits) float64 {
    value := a.value
	switch toUnit { 
	case DurationYear365:
		return (value / (365 * 24 * 3600)) 
	case DurationMonth30:
		return (value / (30 * 24 * 3600)) 
	case DurationWeek:
		return (value / (7 * 24 * 3600)) 
	case DurationDay:
		return (value / (24 * 3600)) 
	case DurationHour:
		return (value / 3600) 
	case DurationMinute:
		return (value / 60) 
	case DurationSecond:
		return (value) 
	case DurationJulianYear:
		return (value / (365.25 * 24 * 3600)) 
	case DurationSol:
		return (value / 88775.244) 
	case DurationNanosecond:
		return ((value) / 1e-09) 
	case DurationMicrosecond:
		return ((value) / 1e-06) 
	case DurationMillisecond:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Duration) convertToBase(value float64, fromUnit DurationUnits) float64 {
	switch fromUnit { 
	case DurationYear365:
		return (value * 365 * 24 * 3600) 
	case DurationMonth30:
		return (value * 30 * 24 * 3600) 
	case DurationWeek:
		return (value * 7 * 24 * 3600) 
	case DurationDay:
		return (value * 24 * 3600) 
	case DurationHour:
		return (value * 3600) 
	case DurationMinute:
		return (value * 60) 
	case DurationSecond:
		return (value) 
	case DurationJulianYear:
		return (value * 365.25 * 24 * 3600) 
	case DurationSol:
		return (value * 88775.244) 
	case DurationNanosecond:
		return ((value) * 1e-09) 
	case DurationMicrosecond:
		return ((value) * 1e-06) 
	case DurationMillisecond:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Duration in the default unit (Second),
// formatted to two decimal places.
func (a Duration) String() string {
	return a.ToString(DurationSecond, 2)
}

// ToString formats the Duration value as a string with the specified unit and fractional digits.
// It converts the Duration to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Duration value will be converted (e.g., Second))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Duration with the unit abbreviation.
func (a *Duration) ToString(unit DurationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetDurationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetDurationAbbreviation(unit))
}

// Equals checks if the given Duration is equal to the current Duration.
//
// Parameters:
//    other: The Duration to compare against.
//
// Returns:
//    bool: Returns true if both Duration are equal, false otherwise.
func (a *Duration) Equals(other *Duration) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Duration with another Duration.
// It returns -1 if the current Duration is less than the other Duration, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Duration to compare against.
//
// Returns:
//    int: -1 if the current Duration is less, 1 if greater, and 0 if equal.
func (a *Duration) CompareTo(other *Duration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Duration to the current Duration and returns the result.
// The result is a new Duration instance with the sum of the values.
//
// Parameters:
//    other: The Duration to add to the current Duration.
//
// Returns:
//    *Duration: A new Duration instance representing the sum of both Duration.
func (a *Duration) Add(other *Duration) *Duration {
	return &Duration{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Duration from the current Duration and returns the result.
// The result is a new Duration instance with the difference of the values.
//
// Parameters:
//    other: The Duration to subtract from the current Duration.
//
// Returns:
//    *Duration: A new Duration instance representing the difference of both Duration.
func (a *Duration) Subtract(other *Duration) *Duration {
	return &Duration{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Duration by the given Duration and returns the result.
// The result is a new Duration instance with the product of the values.
//
// Parameters:
//    other: The Duration to multiply with the current Duration.
//
// Returns:
//    *Duration: A new Duration instance representing the product of both Duration.
func (a *Duration) Multiply(other *Duration) *Duration {
	return &Duration{value: a.value * other.BaseValue()}
}

// Divide divides the current Duration by the given Duration and returns the result.
// The result is a new Duration instance with the quotient of the values.
//
// Parameters:
//    other: The Duration to divide the current Duration by.
//
// Returns:
//    *Duration: A new Duration instance representing the quotient of both Duration.
func (a *Duration) Divide(other *Duration) *Duration {
	return &Duration{value: a.value / other.BaseValue()}
}

// GetDurationAbbreviation gets the unit abbreviation.
func GetDurationAbbreviation(unit DurationUnits) string {
	switch unit { 
	case DurationYear365:
		return "yr" 
	case DurationMonth30:
		return "mo" 
	case DurationWeek:
		return "wk" 
	case DurationDay:
		return "d" 
	case DurationHour:
		return "h" 
	case DurationMinute:
		return "m" 
	case DurationSecond:
		return "s" 
	case DurationJulianYear:
		return "jyr" 
	case DurationSol:
		return "sol" 
	case DurationNanosecond:
		return "ns" 
	case DurationMicrosecond:
		return "μs" 
	case DurationMillisecond:
		return "ms" 
	default:
		return ""
	}
}