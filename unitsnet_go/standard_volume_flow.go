package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// StandardVolumeFlowUnits enumeration
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

// StandardVolumeFlowDto represents an StandardVolumeFlow
type StandardVolumeFlowDto struct {
	Value float64
	Unit  StandardVolumeFlowUnits
}

// StandardVolumeFlowDtoFactory struct to group related functions
type StandardVolumeFlowDtoFactory struct{}

func (udf StandardVolumeFlowDtoFactory) FromJSON(data []byte) (*StandardVolumeFlowDto, error) {
	a := StandardVolumeFlowDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a StandardVolumeFlowDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// StandardVolumeFlow struct
type StandardVolumeFlow struct {
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

// StandardVolumeFlowFactory struct to group related functions
type StandardVolumeFlowFactory struct{}

func (uf StandardVolumeFlowFactory) CreateStandardVolumeFlow(value float64, unit StandardVolumeFlowUnits) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, unit)
}

func (uf StandardVolumeFlowFactory) FromDto(dto StandardVolumeFlowDto) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(dto.Value, dto.Unit)
}

func (uf StandardVolumeFlowFactory) FromDtoJSON(data []byte) (*StandardVolumeFlow, error) {
	unitDto, err := StandardVolumeFlowDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return StandardVolumeFlowFactory{}.FromDto(*unitDto)
}


// FromStandardCubicMeterPerSecond creates a new StandardVolumeFlow instance from StandardCubicMeterPerSecond.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerSecond(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerSecond)
}

// FromStandardCubicMeterPerMinute creates a new StandardVolumeFlow instance from StandardCubicMeterPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerMinute)
}

// FromStandardCubicMeterPerHour creates a new StandardVolumeFlow instance from StandardCubicMeterPerHour.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerHour(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerHour)
}

// FromStandardCubicMeterPerDay creates a new StandardVolumeFlow instance from StandardCubicMeterPerDay.
func (uf StandardVolumeFlowFactory) FromStandardCubicMetersPerDay(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicMeterPerDay)
}

// FromStandardCubicCentimeterPerMinute creates a new StandardVolumeFlow instance from StandardCubicCentimeterPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicCentimetersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicCentimeterPerMinute)
}

// FromStandardLiterPerMinute creates a new StandardVolumeFlow instance from StandardLiterPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardLitersPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardLiterPerMinute)
}

// FromStandardCubicFootPerSecond creates a new StandardVolumeFlow instance from StandardCubicFootPerSecond.
func (uf StandardVolumeFlowFactory) FromStandardCubicFeetPerSecond(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicFootPerSecond)
}

// FromStandardCubicFootPerMinute creates a new StandardVolumeFlow instance from StandardCubicFootPerMinute.
func (uf StandardVolumeFlowFactory) FromStandardCubicFeetPerMinute(value float64) (*StandardVolumeFlow, error) {
	return newStandardVolumeFlow(value, StandardVolumeFlowStandardCubicFootPerMinute)
}

// FromStandardCubicFootPerHour creates a new StandardVolumeFlow instance from StandardCubicFootPerHour.
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

// BaseValue returns the base value of StandardVolumeFlow in StandardCubicMeterPerSecond.
func (a *StandardVolumeFlow) BaseValue() float64 {
	return a.value
}


// StandardCubicMeterPerSecond returns the value in StandardCubicMeterPerSecond.
func (a *StandardVolumeFlow) StandardCubicMetersPerSecond() float64 {
	if a.standard_cubic_meters_per_secondLazy != nil {
		return *a.standard_cubic_meters_per_secondLazy
	}
	standard_cubic_meters_per_second := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerSecond)
	a.standard_cubic_meters_per_secondLazy = &standard_cubic_meters_per_second
	return standard_cubic_meters_per_second
}

// StandardCubicMeterPerMinute returns the value in StandardCubicMeterPerMinute.
func (a *StandardVolumeFlow) StandardCubicMetersPerMinute() float64 {
	if a.standard_cubic_meters_per_minuteLazy != nil {
		return *a.standard_cubic_meters_per_minuteLazy
	}
	standard_cubic_meters_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerMinute)
	a.standard_cubic_meters_per_minuteLazy = &standard_cubic_meters_per_minute
	return standard_cubic_meters_per_minute
}

// StandardCubicMeterPerHour returns the value in StandardCubicMeterPerHour.
func (a *StandardVolumeFlow) StandardCubicMetersPerHour() float64 {
	if a.standard_cubic_meters_per_hourLazy != nil {
		return *a.standard_cubic_meters_per_hourLazy
	}
	standard_cubic_meters_per_hour := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerHour)
	a.standard_cubic_meters_per_hourLazy = &standard_cubic_meters_per_hour
	return standard_cubic_meters_per_hour
}

// StandardCubicMeterPerDay returns the value in StandardCubicMeterPerDay.
func (a *StandardVolumeFlow) StandardCubicMetersPerDay() float64 {
	if a.standard_cubic_meters_per_dayLazy != nil {
		return *a.standard_cubic_meters_per_dayLazy
	}
	standard_cubic_meters_per_day := a.convertFromBase(StandardVolumeFlowStandardCubicMeterPerDay)
	a.standard_cubic_meters_per_dayLazy = &standard_cubic_meters_per_day
	return standard_cubic_meters_per_day
}

// StandardCubicCentimeterPerMinute returns the value in StandardCubicCentimeterPerMinute.
func (a *StandardVolumeFlow) StandardCubicCentimetersPerMinute() float64 {
	if a.standard_cubic_centimeters_per_minuteLazy != nil {
		return *a.standard_cubic_centimeters_per_minuteLazy
	}
	standard_cubic_centimeters_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicCentimeterPerMinute)
	a.standard_cubic_centimeters_per_minuteLazy = &standard_cubic_centimeters_per_minute
	return standard_cubic_centimeters_per_minute
}

// StandardLiterPerMinute returns the value in StandardLiterPerMinute.
func (a *StandardVolumeFlow) StandardLitersPerMinute() float64 {
	if a.standard_liters_per_minuteLazy != nil {
		return *a.standard_liters_per_minuteLazy
	}
	standard_liters_per_minute := a.convertFromBase(StandardVolumeFlowStandardLiterPerMinute)
	a.standard_liters_per_minuteLazy = &standard_liters_per_minute
	return standard_liters_per_minute
}

// StandardCubicFootPerSecond returns the value in StandardCubicFootPerSecond.
func (a *StandardVolumeFlow) StandardCubicFeetPerSecond() float64 {
	if a.standard_cubic_feet_per_secondLazy != nil {
		return *a.standard_cubic_feet_per_secondLazy
	}
	standard_cubic_feet_per_second := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerSecond)
	a.standard_cubic_feet_per_secondLazy = &standard_cubic_feet_per_second
	return standard_cubic_feet_per_second
}

// StandardCubicFootPerMinute returns the value in StandardCubicFootPerMinute.
func (a *StandardVolumeFlow) StandardCubicFeetPerMinute() float64 {
	if a.standard_cubic_feet_per_minuteLazy != nil {
		return *a.standard_cubic_feet_per_minuteLazy
	}
	standard_cubic_feet_per_minute := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerMinute)
	a.standard_cubic_feet_per_minuteLazy = &standard_cubic_feet_per_minute
	return standard_cubic_feet_per_minute
}

// StandardCubicFootPerHour returns the value in StandardCubicFootPerHour.
func (a *StandardVolumeFlow) StandardCubicFeetPerHour() float64 {
	if a.standard_cubic_feet_per_hourLazy != nil {
		return *a.standard_cubic_feet_per_hourLazy
	}
	standard_cubic_feet_per_hour := a.convertFromBase(StandardVolumeFlowStandardCubicFootPerHour)
	a.standard_cubic_feet_per_hourLazy = &standard_cubic_feet_per_hour
	return standard_cubic_feet_per_hour
}


// ToDto creates an StandardVolumeFlowDto representation.
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

// ToDtoJSON creates an StandardVolumeFlowDto representation.
func (a *StandardVolumeFlow) ToDtoJSON(holdInUnit *StandardVolumeFlowUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts StandardVolumeFlow to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a StandardVolumeFlow) String() string {
	return a.ToString(StandardVolumeFlowStandardCubicMeterPerSecond, 2)
}

// ToString formats the StandardVolumeFlow to string.
// fractionalDigits -1 for not mention
func (a *StandardVolumeFlow) ToString(unit StandardVolumeFlowUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *StandardVolumeFlow) getUnitAbbreviation(unit StandardVolumeFlowUnits) string {
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

// Check if the given StandardVolumeFlow are equals to the current StandardVolumeFlow
func (a *StandardVolumeFlow) Equals(other *StandardVolumeFlow) bool {
	return a.value == other.BaseValue()
}

// Check if the given StandardVolumeFlow are equals to the current StandardVolumeFlow
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

// Add the given StandardVolumeFlow to the current StandardVolumeFlow.
func (a *StandardVolumeFlow) Add(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value + other.BaseValue()}
}

// Subtract the given StandardVolumeFlow to the current StandardVolumeFlow.
func (a *StandardVolumeFlow) Subtract(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value - other.BaseValue()}
}

// Multiply the given StandardVolumeFlow to the current StandardVolumeFlow.
func (a *StandardVolumeFlow) Multiply(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value * other.BaseValue()}
}

// Divide the given StandardVolumeFlow to the current StandardVolumeFlow.
func (a *StandardVolumeFlow) Divide(other *StandardVolumeFlow) *StandardVolumeFlow {
	return &StandardVolumeFlow{value: a.value / other.BaseValue()}
}