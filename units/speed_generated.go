// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpeedUnits defines various units of Speed.
type SpeedUnits string

const (
    
        // 
        SpeedMeterPerSecond SpeedUnits = "MeterPerSecond"
        // 
        SpeedMeterPerMinute SpeedUnits = "MeterPerMinute"
        // 
        SpeedMeterPerHour SpeedUnits = "MeterPerHour"
        // 
        SpeedFootPerSecond SpeedUnits = "FootPerSecond"
        // 
        SpeedFootPerMinute SpeedUnits = "FootPerMinute"
        // 
        SpeedFootPerHour SpeedUnits = "FootPerHour"
        // 
        SpeedUsSurveyFootPerSecond SpeedUnits = "UsSurveyFootPerSecond"
        // 
        SpeedUsSurveyFootPerMinute SpeedUnits = "UsSurveyFootPerMinute"
        // 
        SpeedUsSurveyFootPerHour SpeedUnits = "UsSurveyFootPerHour"
        // 
        SpeedInchPerSecond SpeedUnits = "InchPerSecond"
        // 
        SpeedInchPerMinute SpeedUnits = "InchPerMinute"
        // 
        SpeedInchPerHour SpeedUnits = "InchPerHour"
        // 
        SpeedYardPerSecond SpeedUnits = "YardPerSecond"
        // 
        SpeedYardPerMinute SpeedUnits = "YardPerMinute"
        // 
        SpeedYardPerHour SpeedUnits = "YardPerHour"
        // The knot, by definition, is a unit of speed equals to 1 nautical mile per hour, which is exactly 1852.000 metres per hour. The length of the internationally agreed nautical mile is 1852 m. The US adopted the international definition in 1954, the UK adopted the international nautical mile definition in 1970.
        SpeedKnot SpeedUnits = "Knot"
        // 
        SpeedMilePerHour SpeedUnits = "MilePerHour"
        // 
        SpeedMach SpeedUnits = "Mach"
        // 
        SpeedNanometerPerSecond SpeedUnits = "NanometerPerSecond"
        // 
        SpeedMicrometerPerSecond SpeedUnits = "MicrometerPerSecond"
        // 
        SpeedMillimeterPerSecond SpeedUnits = "MillimeterPerSecond"
        // 
        SpeedCentimeterPerSecond SpeedUnits = "CentimeterPerSecond"
        // 
        SpeedDecimeterPerSecond SpeedUnits = "DecimeterPerSecond"
        // 
        SpeedKilometerPerSecond SpeedUnits = "KilometerPerSecond"
        // 
        SpeedNanometerPerMinute SpeedUnits = "NanometerPerMinute"
        // 
        SpeedMicrometerPerMinute SpeedUnits = "MicrometerPerMinute"
        // 
        SpeedMillimeterPerMinute SpeedUnits = "MillimeterPerMinute"
        // 
        SpeedCentimeterPerMinute SpeedUnits = "CentimeterPerMinute"
        // 
        SpeedDecimeterPerMinute SpeedUnits = "DecimeterPerMinute"
        // 
        SpeedKilometerPerMinute SpeedUnits = "KilometerPerMinute"
        // 
        SpeedMillimeterPerHour SpeedUnits = "MillimeterPerHour"
        // 
        SpeedCentimeterPerHour SpeedUnits = "CentimeterPerHour"
        // 
        SpeedKilometerPerHour SpeedUnits = "KilometerPerHour"
)

// SpeedDto represents a Speed measurement with a numerical value and its corresponding unit.
type SpeedDto struct {
    // Value is the numerical representation of the Speed.
	Value float64
    // Unit specifies the unit of measurement for the Speed, as defined in the SpeedUnits enumeration.
	Unit  SpeedUnits
}

// SpeedDtoFactory groups methods for creating and serializing SpeedDto objects.
type SpeedDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpeedDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpeedDtoFactory) FromJSON(data []byte) (*SpeedDto, error) {
	a := SpeedDto{}

    // Parse JSON into SpeedDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a SpeedDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpeedDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Speed represents a measurement in a of Speed.
//
// In everyday use and in kinematics, the speed of an object is the magnitude of its velocity (the rate of change of its position); it is thus a scalar quantity.[1] The average speed of an object in an interval of time is the distance travelled by the object divided by the duration of the interval;[2] the instantaneous speed is the limit of the average speed as the duration of the time interval approaches zero.
type Speed struct {
	// value is the base measurement stored internally.
	value       float64
    
    meters_per_secondLazy *float64 
    meters_per_minutesLazy *float64 
    meters_per_hourLazy *float64 
    feet_per_secondLazy *float64 
    feet_per_minuteLazy *float64 
    feet_per_hourLazy *float64 
    us_survey_feet_per_secondLazy *float64 
    us_survey_feet_per_minuteLazy *float64 
    us_survey_feet_per_hourLazy *float64 
    inches_per_secondLazy *float64 
    inches_per_minuteLazy *float64 
    inches_per_hourLazy *float64 
    yards_per_secondLazy *float64 
    yards_per_minuteLazy *float64 
    yards_per_hourLazy *float64 
    knotsLazy *float64 
    miles_per_hourLazy *float64 
    machLazy *float64 
    nanometers_per_secondLazy *float64 
    micrometers_per_secondLazy *float64 
    millimeters_per_secondLazy *float64 
    centimeters_per_secondLazy *float64 
    decimeters_per_secondLazy *float64 
    kilometers_per_secondLazy *float64 
    nanometers_per_minutesLazy *float64 
    micrometers_per_minutesLazy *float64 
    millimeters_per_minutesLazy *float64 
    centimeters_per_minutesLazy *float64 
    decimeters_per_minutesLazy *float64 
    kilometers_per_minutesLazy *float64 
    millimeters_per_hourLazy *float64 
    centimeters_per_hourLazy *float64 
    kilometers_per_hourLazy *float64 
}

// SpeedFactory groups methods for creating Speed instances.
type SpeedFactory struct{}

// CreateSpeed creates a new Speed instance from the given value and unit.
func (uf SpeedFactory) CreateSpeed(value float64, unit SpeedUnits) (*Speed, error) {
	return newSpeed(value, unit)
}

// FromDto converts a SpeedDto to a Speed instance.
func (uf SpeedFactory) FromDto(dto SpeedDto) (*Speed, error) {
	return newSpeed(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Speed instance.
func (uf SpeedFactory) FromDtoJSON(data []byte) (*Speed, error) {
	unitDto, err := SpeedDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpeedDto from JSON: %w", err)
	}
	return SpeedFactory{}.FromDto(*unitDto)
}


// FromMetersPerSecond creates a new Speed instance from a value in MetersPerSecond.
func (uf SpeedFactory) FromMetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerSecond)
}

// FromMetersPerMinutes creates a new Speed instance from a value in MetersPerMinutes.
func (uf SpeedFactory) FromMetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerMinute)
}

// FromMetersPerHour creates a new Speed instance from a value in MetersPerHour.
func (uf SpeedFactory) FromMetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerHour)
}

// FromFeetPerSecond creates a new Speed instance from a value in FeetPerSecond.
func (uf SpeedFactory) FromFeetPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerSecond)
}

// FromFeetPerMinute creates a new Speed instance from a value in FeetPerMinute.
func (uf SpeedFactory) FromFeetPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerMinute)
}

// FromFeetPerHour creates a new Speed instance from a value in FeetPerHour.
func (uf SpeedFactory) FromFeetPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerHour)
}

// FromUsSurveyFeetPerSecond creates a new Speed instance from a value in UsSurveyFeetPerSecond.
func (uf SpeedFactory) FromUsSurveyFeetPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerSecond)
}

// FromUsSurveyFeetPerMinute creates a new Speed instance from a value in UsSurveyFeetPerMinute.
func (uf SpeedFactory) FromUsSurveyFeetPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerMinute)
}

// FromUsSurveyFeetPerHour creates a new Speed instance from a value in UsSurveyFeetPerHour.
func (uf SpeedFactory) FromUsSurveyFeetPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerHour)
}

// FromInchesPerSecond creates a new Speed instance from a value in InchesPerSecond.
func (uf SpeedFactory) FromInchesPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerSecond)
}

// FromInchesPerMinute creates a new Speed instance from a value in InchesPerMinute.
func (uf SpeedFactory) FromInchesPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerMinute)
}

// FromInchesPerHour creates a new Speed instance from a value in InchesPerHour.
func (uf SpeedFactory) FromInchesPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerHour)
}

// FromYardsPerSecond creates a new Speed instance from a value in YardsPerSecond.
func (uf SpeedFactory) FromYardsPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerSecond)
}

// FromYardsPerMinute creates a new Speed instance from a value in YardsPerMinute.
func (uf SpeedFactory) FromYardsPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerMinute)
}

// FromYardsPerHour creates a new Speed instance from a value in YardsPerHour.
func (uf SpeedFactory) FromYardsPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerHour)
}

// FromKnots creates a new Speed instance from a value in Knots.
func (uf SpeedFactory) FromKnots(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKnot)
}

// FromMilesPerHour creates a new Speed instance from a value in MilesPerHour.
func (uf SpeedFactory) FromMilesPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMilePerHour)
}

// FromMach creates a new Speed instance from a value in Mach.
func (uf SpeedFactory) FromMach(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMach)
}

// FromNanometersPerSecond creates a new Speed instance from a value in NanometersPerSecond.
func (uf SpeedFactory) FromNanometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedNanometerPerSecond)
}

// FromMicrometersPerSecond creates a new Speed instance from a value in MicrometersPerSecond.
func (uf SpeedFactory) FromMicrometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMicrometerPerSecond)
}

// FromMillimetersPerSecond creates a new Speed instance from a value in MillimetersPerSecond.
func (uf SpeedFactory) FromMillimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerSecond)
}

// FromCentimetersPerSecond creates a new Speed instance from a value in CentimetersPerSecond.
func (uf SpeedFactory) FromCentimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerSecond)
}

// FromDecimetersPerSecond creates a new Speed instance from a value in DecimetersPerSecond.
func (uf SpeedFactory) FromDecimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedDecimeterPerSecond)
}

// FromKilometersPerSecond creates a new Speed instance from a value in KilometersPerSecond.
func (uf SpeedFactory) FromKilometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKilometerPerSecond)
}

// FromNanometersPerMinutes creates a new Speed instance from a value in NanometersPerMinutes.
func (uf SpeedFactory) FromNanometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedNanometerPerMinute)
}

// FromMicrometersPerMinutes creates a new Speed instance from a value in MicrometersPerMinutes.
func (uf SpeedFactory) FromMicrometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMicrometerPerMinute)
}

// FromMillimetersPerMinutes creates a new Speed instance from a value in MillimetersPerMinutes.
func (uf SpeedFactory) FromMillimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerMinute)
}

// FromCentimetersPerMinutes creates a new Speed instance from a value in CentimetersPerMinutes.
func (uf SpeedFactory) FromCentimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerMinute)
}

// FromDecimetersPerMinutes creates a new Speed instance from a value in DecimetersPerMinutes.
func (uf SpeedFactory) FromDecimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedDecimeterPerMinute)
}

// FromKilometersPerMinutes creates a new Speed instance from a value in KilometersPerMinutes.
func (uf SpeedFactory) FromKilometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKilometerPerMinute)
}

// FromMillimetersPerHour creates a new Speed instance from a value in MillimetersPerHour.
func (uf SpeedFactory) FromMillimetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerHour)
}

// FromCentimetersPerHour creates a new Speed instance from a value in CentimetersPerHour.
func (uf SpeedFactory) FromCentimetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerHour)
}

// FromKilometersPerHour creates a new Speed instance from a value in KilometersPerHour.
func (uf SpeedFactory) FromKilometersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKilometerPerHour)
}


// newSpeed creates a new Speed.
func newSpeed(value float64, fromUnit SpeedUnits) (*Speed, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Speed{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Speed in MeterPerSecond unit (the base/default quantity).
func (a *Speed) BaseValue() float64 {
	return a.value
}


// MetersPerSecond returns the Speed value in MetersPerSecond.
//
// 
func (a *Speed) MetersPerSecond() float64 {
	if a.meters_per_secondLazy != nil {
		return *a.meters_per_secondLazy
	}
	meters_per_second := a.convertFromBase(SpeedMeterPerSecond)
	a.meters_per_secondLazy = &meters_per_second
	return meters_per_second
}

// MetersPerMinutes returns the Speed value in MetersPerMinutes.
//
// 
func (a *Speed) MetersPerMinutes() float64 {
	if a.meters_per_minutesLazy != nil {
		return *a.meters_per_minutesLazy
	}
	meters_per_minutes := a.convertFromBase(SpeedMeterPerMinute)
	a.meters_per_minutesLazy = &meters_per_minutes
	return meters_per_minutes
}

// MetersPerHour returns the Speed value in MetersPerHour.
//
// 
func (a *Speed) MetersPerHour() float64 {
	if a.meters_per_hourLazy != nil {
		return *a.meters_per_hourLazy
	}
	meters_per_hour := a.convertFromBase(SpeedMeterPerHour)
	a.meters_per_hourLazy = &meters_per_hour
	return meters_per_hour
}

// FeetPerSecond returns the Speed value in FeetPerSecond.
//
// 
func (a *Speed) FeetPerSecond() float64 {
	if a.feet_per_secondLazy != nil {
		return *a.feet_per_secondLazy
	}
	feet_per_second := a.convertFromBase(SpeedFootPerSecond)
	a.feet_per_secondLazy = &feet_per_second
	return feet_per_second
}

// FeetPerMinute returns the Speed value in FeetPerMinute.
//
// 
func (a *Speed) FeetPerMinute() float64 {
	if a.feet_per_minuteLazy != nil {
		return *a.feet_per_minuteLazy
	}
	feet_per_minute := a.convertFromBase(SpeedFootPerMinute)
	a.feet_per_minuteLazy = &feet_per_minute
	return feet_per_minute
}

// FeetPerHour returns the Speed value in FeetPerHour.
//
// 
func (a *Speed) FeetPerHour() float64 {
	if a.feet_per_hourLazy != nil {
		return *a.feet_per_hourLazy
	}
	feet_per_hour := a.convertFromBase(SpeedFootPerHour)
	a.feet_per_hourLazy = &feet_per_hour
	return feet_per_hour
}

// UsSurveyFeetPerSecond returns the Speed value in UsSurveyFeetPerSecond.
//
// 
func (a *Speed) UsSurveyFeetPerSecond() float64 {
	if a.us_survey_feet_per_secondLazy != nil {
		return *a.us_survey_feet_per_secondLazy
	}
	us_survey_feet_per_second := a.convertFromBase(SpeedUsSurveyFootPerSecond)
	a.us_survey_feet_per_secondLazy = &us_survey_feet_per_second
	return us_survey_feet_per_second
}

// UsSurveyFeetPerMinute returns the Speed value in UsSurveyFeetPerMinute.
//
// 
func (a *Speed) UsSurveyFeetPerMinute() float64 {
	if a.us_survey_feet_per_minuteLazy != nil {
		return *a.us_survey_feet_per_minuteLazy
	}
	us_survey_feet_per_minute := a.convertFromBase(SpeedUsSurveyFootPerMinute)
	a.us_survey_feet_per_minuteLazy = &us_survey_feet_per_minute
	return us_survey_feet_per_minute
}

// UsSurveyFeetPerHour returns the Speed value in UsSurveyFeetPerHour.
//
// 
func (a *Speed) UsSurveyFeetPerHour() float64 {
	if a.us_survey_feet_per_hourLazy != nil {
		return *a.us_survey_feet_per_hourLazy
	}
	us_survey_feet_per_hour := a.convertFromBase(SpeedUsSurveyFootPerHour)
	a.us_survey_feet_per_hourLazy = &us_survey_feet_per_hour
	return us_survey_feet_per_hour
}

// InchesPerSecond returns the Speed value in InchesPerSecond.
//
// 
func (a *Speed) InchesPerSecond() float64 {
	if a.inches_per_secondLazy != nil {
		return *a.inches_per_secondLazy
	}
	inches_per_second := a.convertFromBase(SpeedInchPerSecond)
	a.inches_per_secondLazy = &inches_per_second
	return inches_per_second
}

// InchesPerMinute returns the Speed value in InchesPerMinute.
//
// 
func (a *Speed) InchesPerMinute() float64 {
	if a.inches_per_minuteLazy != nil {
		return *a.inches_per_minuteLazy
	}
	inches_per_minute := a.convertFromBase(SpeedInchPerMinute)
	a.inches_per_minuteLazy = &inches_per_minute
	return inches_per_minute
}

// InchesPerHour returns the Speed value in InchesPerHour.
//
// 
func (a *Speed) InchesPerHour() float64 {
	if a.inches_per_hourLazy != nil {
		return *a.inches_per_hourLazy
	}
	inches_per_hour := a.convertFromBase(SpeedInchPerHour)
	a.inches_per_hourLazy = &inches_per_hour
	return inches_per_hour
}

// YardsPerSecond returns the Speed value in YardsPerSecond.
//
// 
func (a *Speed) YardsPerSecond() float64 {
	if a.yards_per_secondLazy != nil {
		return *a.yards_per_secondLazy
	}
	yards_per_second := a.convertFromBase(SpeedYardPerSecond)
	a.yards_per_secondLazy = &yards_per_second
	return yards_per_second
}

// YardsPerMinute returns the Speed value in YardsPerMinute.
//
// 
func (a *Speed) YardsPerMinute() float64 {
	if a.yards_per_minuteLazy != nil {
		return *a.yards_per_minuteLazy
	}
	yards_per_minute := a.convertFromBase(SpeedYardPerMinute)
	a.yards_per_minuteLazy = &yards_per_minute
	return yards_per_minute
}

// YardsPerHour returns the Speed value in YardsPerHour.
//
// 
func (a *Speed) YardsPerHour() float64 {
	if a.yards_per_hourLazy != nil {
		return *a.yards_per_hourLazy
	}
	yards_per_hour := a.convertFromBase(SpeedYardPerHour)
	a.yards_per_hourLazy = &yards_per_hour
	return yards_per_hour
}

// Knots returns the Speed value in Knots.
//
// The knot, by definition, is a unit of speed equals to 1 nautical mile per hour, which is exactly 1852.000 metres per hour. The length of the internationally agreed nautical mile is 1852 m. The US adopted the international definition in 1954, the UK adopted the international nautical mile definition in 1970.
func (a *Speed) Knots() float64 {
	if a.knotsLazy != nil {
		return *a.knotsLazy
	}
	knots := a.convertFromBase(SpeedKnot)
	a.knotsLazy = &knots
	return knots
}

// MilesPerHour returns the Speed value in MilesPerHour.
//
// 
func (a *Speed) MilesPerHour() float64 {
	if a.miles_per_hourLazy != nil {
		return *a.miles_per_hourLazy
	}
	miles_per_hour := a.convertFromBase(SpeedMilePerHour)
	a.miles_per_hourLazy = &miles_per_hour
	return miles_per_hour
}

// Mach returns the Speed value in Mach.
//
// 
func (a *Speed) Mach() float64 {
	if a.machLazy != nil {
		return *a.machLazy
	}
	mach := a.convertFromBase(SpeedMach)
	a.machLazy = &mach
	return mach
}

// NanometersPerSecond returns the Speed value in NanometersPerSecond.
//
// 
func (a *Speed) NanometersPerSecond() float64 {
	if a.nanometers_per_secondLazy != nil {
		return *a.nanometers_per_secondLazy
	}
	nanometers_per_second := a.convertFromBase(SpeedNanometerPerSecond)
	a.nanometers_per_secondLazy = &nanometers_per_second
	return nanometers_per_second
}

// MicrometersPerSecond returns the Speed value in MicrometersPerSecond.
//
// 
func (a *Speed) MicrometersPerSecond() float64 {
	if a.micrometers_per_secondLazy != nil {
		return *a.micrometers_per_secondLazy
	}
	micrometers_per_second := a.convertFromBase(SpeedMicrometerPerSecond)
	a.micrometers_per_secondLazy = &micrometers_per_second
	return micrometers_per_second
}

// MillimetersPerSecond returns the Speed value in MillimetersPerSecond.
//
// 
func (a *Speed) MillimetersPerSecond() float64 {
	if a.millimeters_per_secondLazy != nil {
		return *a.millimeters_per_secondLazy
	}
	millimeters_per_second := a.convertFromBase(SpeedMillimeterPerSecond)
	a.millimeters_per_secondLazy = &millimeters_per_second
	return millimeters_per_second
}

// CentimetersPerSecond returns the Speed value in CentimetersPerSecond.
//
// 
func (a *Speed) CentimetersPerSecond() float64 {
	if a.centimeters_per_secondLazy != nil {
		return *a.centimeters_per_secondLazy
	}
	centimeters_per_second := a.convertFromBase(SpeedCentimeterPerSecond)
	a.centimeters_per_secondLazy = &centimeters_per_second
	return centimeters_per_second
}

// DecimetersPerSecond returns the Speed value in DecimetersPerSecond.
//
// 
func (a *Speed) DecimetersPerSecond() float64 {
	if a.decimeters_per_secondLazy != nil {
		return *a.decimeters_per_secondLazy
	}
	decimeters_per_second := a.convertFromBase(SpeedDecimeterPerSecond)
	a.decimeters_per_secondLazy = &decimeters_per_second
	return decimeters_per_second
}

// KilometersPerSecond returns the Speed value in KilometersPerSecond.
//
// 
func (a *Speed) KilometersPerSecond() float64 {
	if a.kilometers_per_secondLazy != nil {
		return *a.kilometers_per_secondLazy
	}
	kilometers_per_second := a.convertFromBase(SpeedKilometerPerSecond)
	a.kilometers_per_secondLazy = &kilometers_per_second
	return kilometers_per_second
}

// NanometersPerMinutes returns the Speed value in NanometersPerMinutes.
//
// 
func (a *Speed) NanometersPerMinutes() float64 {
	if a.nanometers_per_minutesLazy != nil {
		return *a.nanometers_per_minutesLazy
	}
	nanometers_per_minutes := a.convertFromBase(SpeedNanometerPerMinute)
	a.nanometers_per_minutesLazy = &nanometers_per_minutes
	return nanometers_per_minutes
}

// MicrometersPerMinutes returns the Speed value in MicrometersPerMinutes.
//
// 
func (a *Speed) MicrometersPerMinutes() float64 {
	if a.micrometers_per_minutesLazy != nil {
		return *a.micrometers_per_minutesLazy
	}
	micrometers_per_minutes := a.convertFromBase(SpeedMicrometerPerMinute)
	a.micrometers_per_minutesLazy = &micrometers_per_minutes
	return micrometers_per_minutes
}

// MillimetersPerMinutes returns the Speed value in MillimetersPerMinutes.
//
// 
func (a *Speed) MillimetersPerMinutes() float64 {
	if a.millimeters_per_minutesLazy != nil {
		return *a.millimeters_per_minutesLazy
	}
	millimeters_per_minutes := a.convertFromBase(SpeedMillimeterPerMinute)
	a.millimeters_per_minutesLazy = &millimeters_per_minutes
	return millimeters_per_minutes
}

// CentimetersPerMinutes returns the Speed value in CentimetersPerMinutes.
//
// 
func (a *Speed) CentimetersPerMinutes() float64 {
	if a.centimeters_per_minutesLazy != nil {
		return *a.centimeters_per_minutesLazy
	}
	centimeters_per_minutes := a.convertFromBase(SpeedCentimeterPerMinute)
	a.centimeters_per_minutesLazy = &centimeters_per_minutes
	return centimeters_per_minutes
}

// DecimetersPerMinutes returns the Speed value in DecimetersPerMinutes.
//
// 
func (a *Speed) DecimetersPerMinutes() float64 {
	if a.decimeters_per_minutesLazy != nil {
		return *a.decimeters_per_minutesLazy
	}
	decimeters_per_minutes := a.convertFromBase(SpeedDecimeterPerMinute)
	a.decimeters_per_minutesLazy = &decimeters_per_minutes
	return decimeters_per_minutes
}

// KilometersPerMinutes returns the Speed value in KilometersPerMinutes.
//
// 
func (a *Speed) KilometersPerMinutes() float64 {
	if a.kilometers_per_minutesLazy != nil {
		return *a.kilometers_per_minutesLazy
	}
	kilometers_per_minutes := a.convertFromBase(SpeedKilometerPerMinute)
	a.kilometers_per_minutesLazy = &kilometers_per_minutes
	return kilometers_per_minutes
}

// MillimetersPerHour returns the Speed value in MillimetersPerHour.
//
// 
func (a *Speed) MillimetersPerHour() float64 {
	if a.millimeters_per_hourLazy != nil {
		return *a.millimeters_per_hourLazy
	}
	millimeters_per_hour := a.convertFromBase(SpeedMillimeterPerHour)
	a.millimeters_per_hourLazy = &millimeters_per_hour
	return millimeters_per_hour
}

// CentimetersPerHour returns the Speed value in CentimetersPerHour.
//
// 
func (a *Speed) CentimetersPerHour() float64 {
	if a.centimeters_per_hourLazy != nil {
		return *a.centimeters_per_hourLazy
	}
	centimeters_per_hour := a.convertFromBase(SpeedCentimeterPerHour)
	a.centimeters_per_hourLazy = &centimeters_per_hour
	return centimeters_per_hour
}

// KilometersPerHour returns the Speed value in KilometersPerHour.
//
// 
func (a *Speed) KilometersPerHour() float64 {
	if a.kilometers_per_hourLazy != nil {
		return *a.kilometers_per_hourLazy
	}
	kilometers_per_hour := a.convertFromBase(SpeedKilometerPerHour)
	a.kilometers_per_hourLazy = &kilometers_per_hour
	return kilometers_per_hour
}


// ToDto creates a SpeedDto representation from the Speed instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecond by default.
func (a *Speed) ToDto(holdInUnit *SpeedUnits) SpeedDto {
	if holdInUnit == nil {
		defaultUnit := SpeedMeterPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return SpeedDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Speed instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MeterPerSecond by default.
func (a *Speed) ToDtoJSON(holdInUnit *SpeedUnits) ([]byte, error) {
	// Convert to SpeedDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Speed to a specific unit value.
// The function uses the provided unit type (SpeedUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Speed) Convert(toUnit SpeedUnits) float64 {
	switch toUnit { 
    case SpeedMeterPerSecond:
		return a.MetersPerSecond()
    case SpeedMeterPerMinute:
		return a.MetersPerMinutes()
    case SpeedMeterPerHour:
		return a.MetersPerHour()
    case SpeedFootPerSecond:
		return a.FeetPerSecond()
    case SpeedFootPerMinute:
		return a.FeetPerMinute()
    case SpeedFootPerHour:
		return a.FeetPerHour()
    case SpeedUsSurveyFootPerSecond:
		return a.UsSurveyFeetPerSecond()
    case SpeedUsSurveyFootPerMinute:
		return a.UsSurveyFeetPerMinute()
    case SpeedUsSurveyFootPerHour:
		return a.UsSurveyFeetPerHour()
    case SpeedInchPerSecond:
		return a.InchesPerSecond()
    case SpeedInchPerMinute:
		return a.InchesPerMinute()
    case SpeedInchPerHour:
		return a.InchesPerHour()
    case SpeedYardPerSecond:
		return a.YardsPerSecond()
    case SpeedYardPerMinute:
		return a.YardsPerMinute()
    case SpeedYardPerHour:
		return a.YardsPerHour()
    case SpeedKnot:
		return a.Knots()
    case SpeedMilePerHour:
		return a.MilesPerHour()
    case SpeedMach:
		return a.Mach()
    case SpeedNanometerPerSecond:
		return a.NanometersPerSecond()
    case SpeedMicrometerPerSecond:
		return a.MicrometersPerSecond()
    case SpeedMillimeterPerSecond:
		return a.MillimetersPerSecond()
    case SpeedCentimeterPerSecond:
		return a.CentimetersPerSecond()
    case SpeedDecimeterPerSecond:
		return a.DecimetersPerSecond()
    case SpeedKilometerPerSecond:
		return a.KilometersPerSecond()
    case SpeedNanometerPerMinute:
		return a.NanometersPerMinutes()
    case SpeedMicrometerPerMinute:
		return a.MicrometersPerMinutes()
    case SpeedMillimeterPerMinute:
		return a.MillimetersPerMinutes()
    case SpeedCentimeterPerMinute:
		return a.CentimetersPerMinutes()
    case SpeedDecimeterPerMinute:
		return a.DecimetersPerMinutes()
    case SpeedKilometerPerMinute:
		return a.KilometersPerMinutes()
    case SpeedMillimeterPerHour:
		return a.MillimetersPerHour()
    case SpeedCentimeterPerHour:
		return a.CentimetersPerHour()
    case SpeedKilometerPerHour:
		return a.KilometersPerHour()
	default:
		return math.NaN()
	}
}

func (a *Speed) convertFromBase(toUnit SpeedUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpeedMeterPerSecond:
		return (value) 
	case SpeedMeterPerMinute:
		return (value * 60) 
	case SpeedMeterPerHour:
		return (value * 3600) 
	case SpeedFootPerSecond:
		return (value / 0.3048) 
	case SpeedFootPerMinute:
		return (value / 0.3048 * 60) 
	case SpeedFootPerHour:
		return (value / 0.3048 * 3600) 
	case SpeedUsSurveyFootPerSecond:
		return (value * 3937 / 1200) 
	case SpeedUsSurveyFootPerMinute:
		return ((value * 3937 / 1200) * 60) 
	case SpeedUsSurveyFootPerHour:
		return ((value * 3937 / 1200) * 3600) 
	case SpeedInchPerSecond:
		return (value / 2.54e-2) 
	case SpeedInchPerMinute:
		return ((value / 2.54e-2) * 60) 
	case SpeedInchPerHour:
		return ((value / 2.54e-2) * 3600) 
	case SpeedYardPerSecond:
		return (value / 0.9144) 
	case SpeedYardPerMinute:
		return (value / 0.9144 * 60) 
	case SpeedYardPerHour:
		return (value / 0.9144 * 3600) 
	case SpeedKnot:
		return (value / (1852.0 / 3600.0)) 
	case SpeedMilePerHour:
		return (value / 0.44704) 
	case SpeedMach:
		return (value / 340.29) 
	case SpeedNanometerPerSecond:
		return ((value) / 1e-09) 
	case SpeedMicrometerPerSecond:
		return ((value) / 1e-06) 
	case SpeedMillimeterPerSecond:
		return ((value) / 0.001) 
	case SpeedCentimeterPerSecond:
		return ((value) / 0.01) 
	case SpeedDecimeterPerSecond:
		return ((value) / 0.1) 
	case SpeedKilometerPerSecond:
		return ((value) / 1000.0) 
	case SpeedNanometerPerMinute:
		return ((value * 60) / 1e-09) 
	case SpeedMicrometerPerMinute:
		return ((value * 60) / 1e-06) 
	case SpeedMillimeterPerMinute:
		return ((value * 60) / 0.001) 
	case SpeedCentimeterPerMinute:
		return ((value * 60) / 0.01) 
	case SpeedDecimeterPerMinute:
		return ((value * 60) / 0.1) 
	case SpeedKilometerPerMinute:
		return ((value * 60) / 1000.0) 
	case SpeedMillimeterPerHour:
		return ((value * 3600) / 0.001) 
	case SpeedCentimeterPerHour:
		return ((value * 3600) / 0.01) 
	case SpeedKilometerPerHour:
		return ((value * 3600) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Speed) convertToBase(value float64, fromUnit SpeedUnits) float64 {
	switch fromUnit { 
	case SpeedMeterPerSecond:
		return (value) 
	case SpeedMeterPerMinute:
		return (value / 60) 
	case SpeedMeterPerHour:
		return (value / 3600) 
	case SpeedFootPerSecond:
		return (value * 0.3048) 
	case SpeedFootPerMinute:
		return (value * 0.3048 / 60) 
	case SpeedFootPerHour:
		return (value * 0.3048 / 3600) 
	case SpeedUsSurveyFootPerSecond:
		return (value * 1200 / 3937) 
	case SpeedUsSurveyFootPerMinute:
		return ((value * 1200 / 3937) / 60) 
	case SpeedUsSurveyFootPerHour:
		return ((value * 1200 / 3937) / 3600) 
	case SpeedInchPerSecond:
		return (value * 2.54e-2) 
	case SpeedInchPerMinute:
		return ((value / 60) * 2.54e-2) 
	case SpeedInchPerHour:
		return ((value / 3600) * 2.54e-2) 
	case SpeedYardPerSecond:
		return (value * 0.9144) 
	case SpeedYardPerMinute:
		return (value * 0.9144 / 60) 
	case SpeedYardPerHour:
		return (value * 0.9144 / 3600) 
	case SpeedKnot:
		return (value * (1852.0 / 3600.0)) 
	case SpeedMilePerHour:
		return (value * 0.44704) 
	case SpeedMach:
		return (value * 340.29) 
	case SpeedNanometerPerSecond:
		return ((value) * 1e-09) 
	case SpeedMicrometerPerSecond:
		return ((value) * 1e-06) 
	case SpeedMillimeterPerSecond:
		return ((value) * 0.001) 
	case SpeedCentimeterPerSecond:
		return ((value) * 0.01) 
	case SpeedDecimeterPerSecond:
		return ((value) * 0.1) 
	case SpeedKilometerPerSecond:
		return ((value) * 1000.0) 
	case SpeedNanometerPerMinute:
		return ((value / 60) * 1e-09) 
	case SpeedMicrometerPerMinute:
		return ((value / 60) * 1e-06) 
	case SpeedMillimeterPerMinute:
		return ((value / 60) * 0.001) 
	case SpeedCentimeterPerMinute:
		return ((value / 60) * 0.01) 
	case SpeedDecimeterPerMinute:
		return ((value / 60) * 0.1) 
	case SpeedKilometerPerMinute:
		return ((value / 60) * 1000.0) 
	case SpeedMillimeterPerHour:
		return ((value / 3600) * 0.001) 
	case SpeedCentimeterPerHour:
		return ((value / 3600) * 0.01) 
	case SpeedKilometerPerHour:
		return ((value / 3600) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Speed in the default unit (MeterPerSecond),
// formatted to two decimal places.
func (a Speed) String() string {
	return a.ToString(SpeedMeterPerSecond, 2)
}

// ToString formats the Speed value as a string with the specified unit and fractional digits.
// It converts the Speed to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Speed value will be converted (e.g., MeterPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Speed with the unit abbreviation.
func (a *Speed) ToString(unit SpeedUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Speed) getUnitAbbreviation(unit SpeedUnits) string {
	switch unit { 
	case SpeedMeterPerSecond:
		return "m/s" 
	case SpeedMeterPerMinute:
		return "m/min" 
	case SpeedMeterPerHour:
		return "m/h" 
	case SpeedFootPerSecond:
		return "ft/s" 
	case SpeedFootPerMinute:
		return "ft/min" 
	case SpeedFootPerHour:
		return "ft/h" 
	case SpeedUsSurveyFootPerSecond:
		return "ftUS/s" 
	case SpeedUsSurveyFootPerMinute:
		return "ftUS/min" 
	case SpeedUsSurveyFootPerHour:
		return "ftUS/h" 
	case SpeedInchPerSecond:
		return "in/s" 
	case SpeedInchPerMinute:
		return "in/min" 
	case SpeedInchPerHour:
		return "in/h" 
	case SpeedYardPerSecond:
		return "yd/s" 
	case SpeedYardPerMinute:
		return "yd/min" 
	case SpeedYardPerHour:
		return "yd/h" 
	case SpeedKnot:
		return "kn" 
	case SpeedMilePerHour:
		return "mph" 
	case SpeedMach:
		return "M" 
	case SpeedNanometerPerSecond:
		return "nm/s" 
	case SpeedMicrometerPerSecond:
		return "μm/s" 
	case SpeedMillimeterPerSecond:
		return "mm/s" 
	case SpeedCentimeterPerSecond:
		return "cm/s" 
	case SpeedDecimeterPerSecond:
		return "dm/s" 
	case SpeedKilometerPerSecond:
		return "km/s" 
	case SpeedNanometerPerMinute:
		return "nm/min" 
	case SpeedMicrometerPerMinute:
		return "μm/min" 
	case SpeedMillimeterPerMinute:
		return "mm/min" 
	case SpeedCentimeterPerMinute:
		return "cm/min" 
	case SpeedDecimeterPerMinute:
		return "dm/min" 
	case SpeedKilometerPerMinute:
		return "km/min" 
	case SpeedMillimeterPerHour:
		return "mm/h" 
	case SpeedCentimeterPerHour:
		return "cm/h" 
	case SpeedKilometerPerHour:
		return "km/h" 
	default:
		return ""
	}
}

// Equals checks if the given Speed is equal to the current Speed.
//
// Parameters:
//    other: The Speed to compare against.
//
// Returns:
//    bool: Returns true if both Speed are equal, false otherwise.
func (a *Speed) Equals(other *Speed) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Speed with another Speed.
// It returns -1 if the current Speed is less than the other Speed, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Speed to compare against.
//
// Returns:
//    int: -1 if the current Speed is less, 1 if greater, and 0 if equal.
func (a *Speed) CompareTo(other *Speed) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Speed to the current Speed and returns the result.
// The result is a new Speed instance with the sum of the values.
//
// Parameters:
//    other: The Speed to add to the current Speed.
//
// Returns:
//    *Speed: A new Speed instance representing the sum of both Speed.
func (a *Speed) Add(other *Speed) *Speed {
	return &Speed{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Speed from the current Speed and returns the result.
// The result is a new Speed instance with the difference of the values.
//
// Parameters:
//    other: The Speed to subtract from the current Speed.
//
// Returns:
//    *Speed: A new Speed instance representing the difference of both Speed.
func (a *Speed) Subtract(other *Speed) *Speed {
	return &Speed{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Speed by the given Speed and returns the result.
// The result is a new Speed instance with the product of the values.
//
// Parameters:
//    other: The Speed to multiply with the current Speed.
//
// Returns:
//    *Speed: A new Speed instance representing the product of both Speed.
func (a *Speed) Multiply(other *Speed) *Speed {
	return &Speed{value: a.value * other.BaseValue()}
}

// Divide divides the current Speed by the given Speed and returns the result.
// The result is a new Speed instance with the quotient of the values.
//
// Parameters:
//    other: The Speed to divide the current Speed by.
//
// Returns:
//    *Speed: A new Speed instance representing the quotient of both Speed.
func (a *Speed) Divide(other *Speed) *Speed {
	return &Speed{value: a.value / other.BaseValue()}
}