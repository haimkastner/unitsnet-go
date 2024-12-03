// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpeedUnits enumeration
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

// SpeedDto represents an Speed
type SpeedDto struct {
	Value float64
	Unit  SpeedUnits
}

// SpeedDtoFactory struct to group related functions
type SpeedDtoFactory struct{}

func (udf SpeedDtoFactory) FromJSON(data []byte) (*SpeedDto, error) {
	a := SpeedDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpeedDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Speed struct
type Speed struct {
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

// SpeedFactory struct to group related functions
type SpeedFactory struct{}

func (uf SpeedFactory) CreateSpeed(value float64, unit SpeedUnits) (*Speed, error) {
	return newSpeed(value, unit)
}

func (uf SpeedFactory) FromDto(dto SpeedDto) (*Speed, error) {
	return newSpeed(dto.Value, dto.Unit)
}

func (uf SpeedFactory) FromDtoJSON(data []byte) (*Speed, error) {
	unitDto, err := SpeedDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpeedFactory{}.FromDto(*unitDto)
}


// FromMeterPerSecond creates a new Speed instance from MeterPerSecond.
func (uf SpeedFactory) FromMetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerSecond)
}

// FromMeterPerMinute creates a new Speed instance from MeterPerMinute.
func (uf SpeedFactory) FromMetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerMinute)
}

// FromMeterPerHour creates a new Speed instance from MeterPerHour.
func (uf SpeedFactory) FromMetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMeterPerHour)
}

// FromFootPerSecond creates a new Speed instance from FootPerSecond.
func (uf SpeedFactory) FromFeetPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerSecond)
}

// FromFootPerMinute creates a new Speed instance from FootPerMinute.
func (uf SpeedFactory) FromFeetPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerMinute)
}

// FromFootPerHour creates a new Speed instance from FootPerHour.
func (uf SpeedFactory) FromFeetPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedFootPerHour)
}

// FromUsSurveyFootPerSecond creates a new Speed instance from UsSurveyFootPerSecond.
func (uf SpeedFactory) FromUsSurveyFeetPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerSecond)
}

// FromUsSurveyFootPerMinute creates a new Speed instance from UsSurveyFootPerMinute.
func (uf SpeedFactory) FromUsSurveyFeetPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerMinute)
}

// FromUsSurveyFootPerHour creates a new Speed instance from UsSurveyFootPerHour.
func (uf SpeedFactory) FromUsSurveyFeetPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedUsSurveyFootPerHour)
}

// FromInchPerSecond creates a new Speed instance from InchPerSecond.
func (uf SpeedFactory) FromInchesPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerSecond)
}

// FromInchPerMinute creates a new Speed instance from InchPerMinute.
func (uf SpeedFactory) FromInchesPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerMinute)
}

// FromInchPerHour creates a new Speed instance from InchPerHour.
func (uf SpeedFactory) FromInchesPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedInchPerHour)
}

// FromYardPerSecond creates a new Speed instance from YardPerSecond.
func (uf SpeedFactory) FromYardsPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerSecond)
}

// FromYardPerMinute creates a new Speed instance from YardPerMinute.
func (uf SpeedFactory) FromYardsPerMinute(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerMinute)
}

// FromYardPerHour creates a new Speed instance from YardPerHour.
func (uf SpeedFactory) FromYardsPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedYardPerHour)
}

// FromKnot creates a new Speed instance from Knot.
func (uf SpeedFactory) FromKnots(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKnot)
}

// FromMilePerHour creates a new Speed instance from MilePerHour.
func (uf SpeedFactory) FromMilesPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMilePerHour)
}

// FromMach creates a new Speed instance from Mach.
func (uf SpeedFactory) FromMach(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMach)
}

// FromNanometerPerSecond creates a new Speed instance from NanometerPerSecond.
func (uf SpeedFactory) FromNanometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedNanometerPerSecond)
}

// FromMicrometerPerSecond creates a new Speed instance from MicrometerPerSecond.
func (uf SpeedFactory) FromMicrometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMicrometerPerSecond)
}

// FromMillimeterPerSecond creates a new Speed instance from MillimeterPerSecond.
func (uf SpeedFactory) FromMillimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerSecond)
}

// FromCentimeterPerSecond creates a new Speed instance from CentimeterPerSecond.
func (uf SpeedFactory) FromCentimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerSecond)
}

// FromDecimeterPerSecond creates a new Speed instance from DecimeterPerSecond.
func (uf SpeedFactory) FromDecimetersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedDecimeterPerSecond)
}

// FromKilometerPerSecond creates a new Speed instance from KilometerPerSecond.
func (uf SpeedFactory) FromKilometersPerSecond(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKilometerPerSecond)
}

// FromNanometerPerMinute creates a new Speed instance from NanometerPerMinute.
func (uf SpeedFactory) FromNanometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedNanometerPerMinute)
}

// FromMicrometerPerMinute creates a new Speed instance from MicrometerPerMinute.
func (uf SpeedFactory) FromMicrometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMicrometerPerMinute)
}

// FromMillimeterPerMinute creates a new Speed instance from MillimeterPerMinute.
func (uf SpeedFactory) FromMillimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerMinute)
}

// FromCentimeterPerMinute creates a new Speed instance from CentimeterPerMinute.
func (uf SpeedFactory) FromCentimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerMinute)
}

// FromDecimeterPerMinute creates a new Speed instance from DecimeterPerMinute.
func (uf SpeedFactory) FromDecimetersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedDecimeterPerMinute)
}

// FromKilometerPerMinute creates a new Speed instance from KilometerPerMinute.
func (uf SpeedFactory) FromKilometersPerMinutes(value float64) (*Speed, error) {
	return newSpeed(value, SpeedKilometerPerMinute)
}

// FromMillimeterPerHour creates a new Speed instance from MillimeterPerHour.
func (uf SpeedFactory) FromMillimetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedMillimeterPerHour)
}

// FromCentimeterPerHour creates a new Speed instance from CentimeterPerHour.
func (uf SpeedFactory) FromCentimetersPerHour(value float64) (*Speed, error) {
	return newSpeed(value, SpeedCentimeterPerHour)
}

// FromKilometerPerHour creates a new Speed instance from KilometerPerHour.
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

// BaseValue returns the base value of Speed in MeterPerSecond.
func (a *Speed) BaseValue() float64 {
	return a.value
}


// MeterPerSecond returns the value in MeterPerSecond.
func (a *Speed) MetersPerSecond() float64 {
	if a.meters_per_secondLazy != nil {
		return *a.meters_per_secondLazy
	}
	meters_per_second := a.convertFromBase(SpeedMeterPerSecond)
	a.meters_per_secondLazy = &meters_per_second
	return meters_per_second
}

// MeterPerMinute returns the value in MeterPerMinute.
func (a *Speed) MetersPerMinutes() float64 {
	if a.meters_per_minutesLazy != nil {
		return *a.meters_per_minutesLazy
	}
	meters_per_minutes := a.convertFromBase(SpeedMeterPerMinute)
	a.meters_per_minutesLazy = &meters_per_minutes
	return meters_per_minutes
}

// MeterPerHour returns the value in MeterPerHour.
func (a *Speed) MetersPerHour() float64 {
	if a.meters_per_hourLazy != nil {
		return *a.meters_per_hourLazy
	}
	meters_per_hour := a.convertFromBase(SpeedMeterPerHour)
	a.meters_per_hourLazy = &meters_per_hour
	return meters_per_hour
}

// FootPerSecond returns the value in FootPerSecond.
func (a *Speed) FeetPerSecond() float64 {
	if a.feet_per_secondLazy != nil {
		return *a.feet_per_secondLazy
	}
	feet_per_second := a.convertFromBase(SpeedFootPerSecond)
	a.feet_per_secondLazy = &feet_per_second
	return feet_per_second
}

// FootPerMinute returns the value in FootPerMinute.
func (a *Speed) FeetPerMinute() float64 {
	if a.feet_per_minuteLazy != nil {
		return *a.feet_per_minuteLazy
	}
	feet_per_minute := a.convertFromBase(SpeedFootPerMinute)
	a.feet_per_minuteLazy = &feet_per_minute
	return feet_per_minute
}

// FootPerHour returns the value in FootPerHour.
func (a *Speed) FeetPerHour() float64 {
	if a.feet_per_hourLazy != nil {
		return *a.feet_per_hourLazy
	}
	feet_per_hour := a.convertFromBase(SpeedFootPerHour)
	a.feet_per_hourLazy = &feet_per_hour
	return feet_per_hour
}

// UsSurveyFootPerSecond returns the value in UsSurveyFootPerSecond.
func (a *Speed) UsSurveyFeetPerSecond() float64 {
	if a.us_survey_feet_per_secondLazy != nil {
		return *a.us_survey_feet_per_secondLazy
	}
	us_survey_feet_per_second := a.convertFromBase(SpeedUsSurveyFootPerSecond)
	a.us_survey_feet_per_secondLazy = &us_survey_feet_per_second
	return us_survey_feet_per_second
}

// UsSurveyFootPerMinute returns the value in UsSurveyFootPerMinute.
func (a *Speed) UsSurveyFeetPerMinute() float64 {
	if a.us_survey_feet_per_minuteLazy != nil {
		return *a.us_survey_feet_per_minuteLazy
	}
	us_survey_feet_per_minute := a.convertFromBase(SpeedUsSurveyFootPerMinute)
	a.us_survey_feet_per_minuteLazy = &us_survey_feet_per_minute
	return us_survey_feet_per_minute
}

// UsSurveyFootPerHour returns the value in UsSurveyFootPerHour.
func (a *Speed) UsSurveyFeetPerHour() float64 {
	if a.us_survey_feet_per_hourLazy != nil {
		return *a.us_survey_feet_per_hourLazy
	}
	us_survey_feet_per_hour := a.convertFromBase(SpeedUsSurveyFootPerHour)
	a.us_survey_feet_per_hourLazy = &us_survey_feet_per_hour
	return us_survey_feet_per_hour
}

// InchPerSecond returns the value in InchPerSecond.
func (a *Speed) InchesPerSecond() float64 {
	if a.inches_per_secondLazy != nil {
		return *a.inches_per_secondLazy
	}
	inches_per_second := a.convertFromBase(SpeedInchPerSecond)
	a.inches_per_secondLazy = &inches_per_second
	return inches_per_second
}

// InchPerMinute returns the value in InchPerMinute.
func (a *Speed) InchesPerMinute() float64 {
	if a.inches_per_minuteLazy != nil {
		return *a.inches_per_minuteLazy
	}
	inches_per_minute := a.convertFromBase(SpeedInchPerMinute)
	a.inches_per_minuteLazy = &inches_per_minute
	return inches_per_minute
}

// InchPerHour returns the value in InchPerHour.
func (a *Speed) InchesPerHour() float64 {
	if a.inches_per_hourLazy != nil {
		return *a.inches_per_hourLazy
	}
	inches_per_hour := a.convertFromBase(SpeedInchPerHour)
	a.inches_per_hourLazy = &inches_per_hour
	return inches_per_hour
}

// YardPerSecond returns the value in YardPerSecond.
func (a *Speed) YardsPerSecond() float64 {
	if a.yards_per_secondLazy != nil {
		return *a.yards_per_secondLazy
	}
	yards_per_second := a.convertFromBase(SpeedYardPerSecond)
	a.yards_per_secondLazy = &yards_per_second
	return yards_per_second
}

// YardPerMinute returns the value in YardPerMinute.
func (a *Speed) YardsPerMinute() float64 {
	if a.yards_per_minuteLazy != nil {
		return *a.yards_per_minuteLazy
	}
	yards_per_minute := a.convertFromBase(SpeedYardPerMinute)
	a.yards_per_minuteLazy = &yards_per_minute
	return yards_per_minute
}

// YardPerHour returns the value in YardPerHour.
func (a *Speed) YardsPerHour() float64 {
	if a.yards_per_hourLazy != nil {
		return *a.yards_per_hourLazy
	}
	yards_per_hour := a.convertFromBase(SpeedYardPerHour)
	a.yards_per_hourLazy = &yards_per_hour
	return yards_per_hour
}

// Knot returns the value in Knot.
func (a *Speed) Knots() float64 {
	if a.knotsLazy != nil {
		return *a.knotsLazy
	}
	knots := a.convertFromBase(SpeedKnot)
	a.knotsLazy = &knots
	return knots
}

// MilePerHour returns the value in MilePerHour.
func (a *Speed) MilesPerHour() float64 {
	if a.miles_per_hourLazy != nil {
		return *a.miles_per_hourLazy
	}
	miles_per_hour := a.convertFromBase(SpeedMilePerHour)
	a.miles_per_hourLazy = &miles_per_hour
	return miles_per_hour
}

// Mach returns the value in Mach.
func (a *Speed) Mach() float64 {
	if a.machLazy != nil {
		return *a.machLazy
	}
	mach := a.convertFromBase(SpeedMach)
	a.machLazy = &mach
	return mach
}

// NanometerPerSecond returns the value in NanometerPerSecond.
func (a *Speed) NanometersPerSecond() float64 {
	if a.nanometers_per_secondLazy != nil {
		return *a.nanometers_per_secondLazy
	}
	nanometers_per_second := a.convertFromBase(SpeedNanometerPerSecond)
	a.nanometers_per_secondLazy = &nanometers_per_second
	return nanometers_per_second
}

// MicrometerPerSecond returns the value in MicrometerPerSecond.
func (a *Speed) MicrometersPerSecond() float64 {
	if a.micrometers_per_secondLazy != nil {
		return *a.micrometers_per_secondLazy
	}
	micrometers_per_second := a.convertFromBase(SpeedMicrometerPerSecond)
	a.micrometers_per_secondLazy = &micrometers_per_second
	return micrometers_per_second
}

// MillimeterPerSecond returns the value in MillimeterPerSecond.
func (a *Speed) MillimetersPerSecond() float64 {
	if a.millimeters_per_secondLazy != nil {
		return *a.millimeters_per_secondLazy
	}
	millimeters_per_second := a.convertFromBase(SpeedMillimeterPerSecond)
	a.millimeters_per_secondLazy = &millimeters_per_second
	return millimeters_per_second
}

// CentimeterPerSecond returns the value in CentimeterPerSecond.
func (a *Speed) CentimetersPerSecond() float64 {
	if a.centimeters_per_secondLazy != nil {
		return *a.centimeters_per_secondLazy
	}
	centimeters_per_second := a.convertFromBase(SpeedCentimeterPerSecond)
	a.centimeters_per_secondLazy = &centimeters_per_second
	return centimeters_per_second
}

// DecimeterPerSecond returns the value in DecimeterPerSecond.
func (a *Speed) DecimetersPerSecond() float64 {
	if a.decimeters_per_secondLazy != nil {
		return *a.decimeters_per_secondLazy
	}
	decimeters_per_second := a.convertFromBase(SpeedDecimeterPerSecond)
	a.decimeters_per_secondLazy = &decimeters_per_second
	return decimeters_per_second
}

// KilometerPerSecond returns the value in KilometerPerSecond.
func (a *Speed) KilometersPerSecond() float64 {
	if a.kilometers_per_secondLazy != nil {
		return *a.kilometers_per_secondLazy
	}
	kilometers_per_second := a.convertFromBase(SpeedKilometerPerSecond)
	a.kilometers_per_secondLazy = &kilometers_per_second
	return kilometers_per_second
}

// NanometerPerMinute returns the value in NanometerPerMinute.
func (a *Speed) NanometersPerMinutes() float64 {
	if a.nanometers_per_minutesLazy != nil {
		return *a.nanometers_per_minutesLazy
	}
	nanometers_per_minutes := a.convertFromBase(SpeedNanometerPerMinute)
	a.nanometers_per_minutesLazy = &nanometers_per_minutes
	return nanometers_per_minutes
}

// MicrometerPerMinute returns the value in MicrometerPerMinute.
func (a *Speed) MicrometersPerMinutes() float64 {
	if a.micrometers_per_minutesLazy != nil {
		return *a.micrometers_per_minutesLazy
	}
	micrometers_per_minutes := a.convertFromBase(SpeedMicrometerPerMinute)
	a.micrometers_per_minutesLazy = &micrometers_per_minutes
	return micrometers_per_minutes
}

// MillimeterPerMinute returns the value in MillimeterPerMinute.
func (a *Speed) MillimetersPerMinutes() float64 {
	if a.millimeters_per_minutesLazy != nil {
		return *a.millimeters_per_minutesLazy
	}
	millimeters_per_minutes := a.convertFromBase(SpeedMillimeterPerMinute)
	a.millimeters_per_minutesLazy = &millimeters_per_minutes
	return millimeters_per_minutes
}

// CentimeterPerMinute returns the value in CentimeterPerMinute.
func (a *Speed) CentimetersPerMinutes() float64 {
	if a.centimeters_per_minutesLazy != nil {
		return *a.centimeters_per_minutesLazy
	}
	centimeters_per_minutes := a.convertFromBase(SpeedCentimeterPerMinute)
	a.centimeters_per_minutesLazy = &centimeters_per_minutes
	return centimeters_per_minutes
}

// DecimeterPerMinute returns the value in DecimeterPerMinute.
func (a *Speed) DecimetersPerMinutes() float64 {
	if a.decimeters_per_minutesLazy != nil {
		return *a.decimeters_per_minutesLazy
	}
	decimeters_per_minutes := a.convertFromBase(SpeedDecimeterPerMinute)
	a.decimeters_per_minutesLazy = &decimeters_per_minutes
	return decimeters_per_minutes
}

// KilometerPerMinute returns the value in KilometerPerMinute.
func (a *Speed) KilometersPerMinutes() float64 {
	if a.kilometers_per_minutesLazy != nil {
		return *a.kilometers_per_minutesLazy
	}
	kilometers_per_minutes := a.convertFromBase(SpeedKilometerPerMinute)
	a.kilometers_per_minutesLazy = &kilometers_per_minutes
	return kilometers_per_minutes
}

// MillimeterPerHour returns the value in MillimeterPerHour.
func (a *Speed) MillimetersPerHour() float64 {
	if a.millimeters_per_hourLazy != nil {
		return *a.millimeters_per_hourLazy
	}
	millimeters_per_hour := a.convertFromBase(SpeedMillimeterPerHour)
	a.millimeters_per_hourLazy = &millimeters_per_hour
	return millimeters_per_hour
}

// CentimeterPerHour returns the value in CentimeterPerHour.
func (a *Speed) CentimetersPerHour() float64 {
	if a.centimeters_per_hourLazy != nil {
		return *a.centimeters_per_hourLazy
	}
	centimeters_per_hour := a.convertFromBase(SpeedCentimeterPerHour)
	a.centimeters_per_hourLazy = &centimeters_per_hour
	return centimeters_per_hour
}

// KilometerPerHour returns the value in KilometerPerHour.
func (a *Speed) KilometersPerHour() float64 {
	if a.kilometers_per_hourLazy != nil {
		return *a.kilometers_per_hourLazy
	}
	kilometers_per_hour := a.convertFromBase(SpeedKilometerPerHour)
	a.kilometers_per_hourLazy = &kilometers_per_hour
	return kilometers_per_hour
}


// ToDto creates an SpeedDto representation.
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

// ToDtoJSON creates an SpeedDto representation.
func (a *Speed) ToDtoJSON(holdInUnit *SpeedUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Speed to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Speed) String() string {
	return a.ToString(SpeedMeterPerSecond, 2)
}

// ToString formats the Speed to string.
// fractionalDigits -1 for not mention
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

// Check if the given Speed are equals to the current Speed
func (a *Speed) Equals(other *Speed) bool {
	return a.value == other.BaseValue()
}

// Check if the given Speed are equals to the current Speed
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

// Add the given Speed to the current Speed.
func (a *Speed) Add(other *Speed) *Speed {
	return &Speed{value: a.value + other.BaseValue()}
}

// Subtract the given Speed to the current Speed.
func (a *Speed) Subtract(other *Speed) *Speed {
	return &Speed{value: a.value - other.BaseValue()}
}

// Multiply the given Speed to the current Speed.
func (a *Speed) Multiply(other *Speed) *Speed {
	return &Speed{value: a.value * other.BaseValue()}
}

// Divide the given Speed to the current Speed.
func (a *Speed) Divide(other *Speed) *Speed {
	return &Speed{value: a.value / other.BaseValue()}
}