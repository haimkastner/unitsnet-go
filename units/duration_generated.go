// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// DurationUnits enumeration
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

// DurationDto represents an Duration
type DurationDto struct {
	Value float64
	Unit  DurationUnits
}

// DurationDtoFactory struct to group related functions
type DurationDtoFactory struct{}

func (udf DurationDtoFactory) FromJSON(data []byte) (*DurationDto, error) {
	a := DurationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a DurationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Duration struct
type Duration struct {
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

// DurationFactory struct to group related functions
type DurationFactory struct{}

func (uf DurationFactory) CreateDuration(value float64, unit DurationUnits) (*Duration, error) {
	return newDuration(value, unit)
}

func (uf DurationFactory) FromDto(dto DurationDto) (*Duration, error) {
	return newDuration(dto.Value, dto.Unit)
}

func (uf DurationFactory) FromDtoJSON(data []byte) (*Duration, error) {
	unitDto, err := DurationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return DurationFactory{}.FromDto(*unitDto)
}


// FromYear365 creates a new Duration instance from Year365.
func (uf DurationFactory) FromYears365(value float64) (*Duration, error) {
	return newDuration(value, DurationYear365)
}

// FromMonth30 creates a new Duration instance from Month30.
func (uf DurationFactory) FromMonths30(value float64) (*Duration, error) {
	return newDuration(value, DurationMonth30)
}

// FromWeek creates a new Duration instance from Week.
func (uf DurationFactory) FromWeeks(value float64) (*Duration, error) {
	return newDuration(value, DurationWeek)
}

// FromDay creates a new Duration instance from Day.
func (uf DurationFactory) FromDays(value float64) (*Duration, error) {
	return newDuration(value, DurationDay)
}

// FromHour creates a new Duration instance from Hour.
func (uf DurationFactory) FromHours(value float64) (*Duration, error) {
	return newDuration(value, DurationHour)
}

// FromMinute creates a new Duration instance from Minute.
func (uf DurationFactory) FromMinutes(value float64) (*Duration, error) {
	return newDuration(value, DurationMinute)
}

// FromSecond creates a new Duration instance from Second.
func (uf DurationFactory) FromSeconds(value float64) (*Duration, error) {
	return newDuration(value, DurationSecond)
}

// FromJulianYear creates a new Duration instance from JulianYear.
func (uf DurationFactory) FromJulianYears(value float64) (*Duration, error) {
	return newDuration(value, DurationJulianYear)
}

// FromSol creates a new Duration instance from Sol.
func (uf DurationFactory) FromSols(value float64) (*Duration, error) {
	return newDuration(value, DurationSol)
}

// FromNanosecond creates a new Duration instance from Nanosecond.
func (uf DurationFactory) FromNanoseconds(value float64) (*Duration, error) {
	return newDuration(value, DurationNanosecond)
}

// FromMicrosecond creates a new Duration instance from Microsecond.
func (uf DurationFactory) FromMicroseconds(value float64) (*Duration, error) {
	return newDuration(value, DurationMicrosecond)
}

// FromMillisecond creates a new Duration instance from Millisecond.
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

// BaseValue returns the base value of Duration in Second.
func (a *Duration) BaseValue() float64 {
	return a.value
}


// Year365 returns the value in Year365.
func (a *Duration) Years365() float64 {
	if a.years365Lazy != nil {
		return *a.years365Lazy
	}
	years365 := a.convertFromBase(DurationYear365)
	a.years365Lazy = &years365
	return years365
}

// Month30 returns the value in Month30.
func (a *Duration) Months30() float64 {
	if a.months30Lazy != nil {
		return *a.months30Lazy
	}
	months30 := a.convertFromBase(DurationMonth30)
	a.months30Lazy = &months30
	return months30
}

// Week returns the value in Week.
func (a *Duration) Weeks() float64 {
	if a.weeksLazy != nil {
		return *a.weeksLazy
	}
	weeks := a.convertFromBase(DurationWeek)
	a.weeksLazy = &weeks
	return weeks
}

// Day returns the value in Day.
func (a *Duration) Days() float64 {
	if a.daysLazy != nil {
		return *a.daysLazy
	}
	days := a.convertFromBase(DurationDay)
	a.daysLazy = &days
	return days
}

// Hour returns the value in Hour.
func (a *Duration) Hours() float64 {
	if a.hoursLazy != nil {
		return *a.hoursLazy
	}
	hours := a.convertFromBase(DurationHour)
	a.hoursLazy = &hours
	return hours
}

// Minute returns the value in Minute.
func (a *Duration) Minutes() float64 {
	if a.minutesLazy != nil {
		return *a.minutesLazy
	}
	minutes := a.convertFromBase(DurationMinute)
	a.minutesLazy = &minutes
	return minutes
}

// Second returns the value in Second.
func (a *Duration) Seconds() float64 {
	if a.secondsLazy != nil {
		return *a.secondsLazy
	}
	seconds := a.convertFromBase(DurationSecond)
	a.secondsLazy = &seconds
	return seconds
}

// JulianYear returns the value in JulianYear.
func (a *Duration) JulianYears() float64 {
	if a.julian_yearsLazy != nil {
		return *a.julian_yearsLazy
	}
	julian_years := a.convertFromBase(DurationJulianYear)
	a.julian_yearsLazy = &julian_years
	return julian_years
}

// Sol returns the value in Sol.
func (a *Duration) Sols() float64 {
	if a.solsLazy != nil {
		return *a.solsLazy
	}
	sols := a.convertFromBase(DurationSol)
	a.solsLazy = &sols
	return sols
}

// Nanosecond returns the value in Nanosecond.
func (a *Duration) Nanoseconds() float64 {
	if a.nanosecondsLazy != nil {
		return *a.nanosecondsLazy
	}
	nanoseconds := a.convertFromBase(DurationNanosecond)
	a.nanosecondsLazy = &nanoseconds
	return nanoseconds
}

// Microsecond returns the value in Microsecond.
func (a *Duration) Microseconds() float64 {
	if a.microsecondsLazy != nil {
		return *a.microsecondsLazy
	}
	microseconds := a.convertFromBase(DurationMicrosecond)
	a.microsecondsLazy = &microseconds
	return microseconds
}

// Millisecond returns the value in Millisecond.
func (a *Duration) Milliseconds() float64 {
	if a.millisecondsLazy != nil {
		return *a.millisecondsLazy
	}
	milliseconds := a.convertFromBase(DurationMillisecond)
	a.millisecondsLazy = &milliseconds
	return milliseconds
}


// ToDto creates an DurationDto representation.
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

// ToDtoJSON creates an DurationDto representation.
func (a *Duration) ToDtoJSON(holdInUnit *DurationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Duration to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Duration) String() string {
	return a.ToString(DurationSecond, 2)
}

// ToString formats the Duration to string.
// fractionalDigits -1 for not mention
func (a *Duration) ToString(unit DurationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Duration) getUnitAbbreviation(unit DurationUnits) string {
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

// Check if the given Duration are equals to the current Duration
func (a *Duration) Equals(other *Duration) bool {
	return a.value == other.BaseValue()
}

// Check if the given Duration are equals to the current Duration
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

// Add the given Duration to the current Duration.
func (a *Duration) Add(other *Duration) *Duration {
	return &Duration{value: a.value + other.BaseValue()}
}

// Subtract the given Duration to the current Duration.
func (a *Duration) Subtract(other *Duration) *Duration {
	return &Duration{value: a.value - other.BaseValue()}
}

// Multiply the given Duration to the current Duration.
func (a *Duration) Multiply(other *Duration) *Duration {
	return &Duration{value: a.value * other.BaseValue()}
}

// Divide the given Duration to the current Duration.
func (a *Duration) Divide(other *Duration) *Duration {
	return &Duration{value: a.value / other.BaseValue()}
}