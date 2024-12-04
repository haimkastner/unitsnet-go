// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AngleUnits defines various units of Angle.
type AngleUnits string

const (
    
        // 
        AngleRadian AngleUnits = "Radian"
        // 
        AngleDegree AngleUnits = "Degree"
        // 
        AngleArcminute AngleUnits = "Arcminute"
        // 
        AngleArcsecond AngleUnits = "Arcsecond"
        // 
        AngleGradian AngleUnits = "Gradian"
        // 
        AngleNatoMil AngleUnits = "NatoMil"
        // 
        AngleRevolution AngleUnits = "Revolution"
        // 
        AngleTilt AngleUnits = "Tilt"
        // 
        AngleNanoradian AngleUnits = "Nanoradian"
        // 
        AngleMicroradian AngleUnits = "Microradian"
        // 
        AngleMilliradian AngleUnits = "Milliradian"
        // 
        AngleCentiradian AngleUnits = "Centiradian"
        // 
        AngleDeciradian AngleUnits = "Deciradian"
        // 
        AngleNanodegree AngleUnits = "Nanodegree"
        // 
        AngleMicrodegree AngleUnits = "Microdegree"
        // 
        AngleMillidegree AngleUnits = "Millidegree"
)

// AngleDto represents a Angle measurement with a numerical value and its corresponding unit.
type AngleDto struct {
    // Value is the numerical representation of the Angle.
	Value float64
    // Unit specifies the unit of measurement for the Angle, as defined in the AngleUnits enumeration.
	Unit  AngleUnits
}

// AngleDtoFactory groups methods for creating and serializing AngleDto objects.
type AngleDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AngleDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AngleDtoFactory) FromJSON(data []byte) (*AngleDto, error) {
	a := AngleDto{}

    // Parse JSON into AngleDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a AngleDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AngleDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Angle represents a measurement in a of Angle.
//
// In geometry, an angle is the figure formed by two rays, called the sides of the angle, sharing a common endpoint, called the vertex of the angle.
type Angle struct {
	// value is the base measurement stored internally.
	value       float64
    
    radiansLazy *float64 
    degreesLazy *float64 
    arcminutesLazy *float64 
    arcsecondsLazy *float64 
    gradiansLazy *float64 
    nato_milsLazy *float64 
    revolutionsLazy *float64 
    tiltLazy *float64 
    nanoradiansLazy *float64 
    microradiansLazy *float64 
    milliradiansLazy *float64 
    centiradiansLazy *float64 
    deciradiansLazy *float64 
    nanodegreesLazy *float64 
    microdegreesLazy *float64 
    millidegreesLazy *float64 
}

// AngleFactory groups methods for creating Angle instances.
type AngleFactory struct{}

// CreateAngle creates a new Angle instance from the given value and unit.
func (uf AngleFactory) CreateAngle(value float64, unit AngleUnits) (*Angle, error) {
	return newAngle(value, unit)
}

// FromDto converts a AngleDto to a Angle instance.
func (uf AngleFactory) FromDto(dto AngleDto) (*Angle, error) {
	return newAngle(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Angle instance.
func (uf AngleFactory) FromDtoJSON(data []byte) (*Angle, error) {
	unitDto, err := AngleDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AngleDto from JSON: %w", err)
	}
	return AngleFactory{}.FromDto(*unitDto)
}


// FromRadians creates a new Angle instance from a value in Radians.
func (uf AngleFactory) FromRadians(value float64) (*Angle, error) {
	return newAngle(value, AngleRadian)
}

// FromDegrees creates a new Angle instance from a value in Degrees.
func (uf AngleFactory) FromDegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleDegree)
}

// FromArcminutes creates a new Angle instance from a value in Arcminutes.
func (uf AngleFactory) FromArcminutes(value float64) (*Angle, error) {
	return newAngle(value, AngleArcminute)
}

// FromArcseconds creates a new Angle instance from a value in Arcseconds.
func (uf AngleFactory) FromArcseconds(value float64) (*Angle, error) {
	return newAngle(value, AngleArcsecond)
}

// FromGradians creates a new Angle instance from a value in Gradians.
func (uf AngleFactory) FromGradians(value float64) (*Angle, error) {
	return newAngle(value, AngleGradian)
}

// FromNatoMils creates a new Angle instance from a value in NatoMils.
func (uf AngleFactory) FromNatoMils(value float64) (*Angle, error) {
	return newAngle(value, AngleNatoMil)
}

// FromRevolutions creates a new Angle instance from a value in Revolutions.
func (uf AngleFactory) FromRevolutions(value float64) (*Angle, error) {
	return newAngle(value, AngleRevolution)
}

// FromTilt creates a new Angle instance from a value in Tilt.
func (uf AngleFactory) FromTilt(value float64) (*Angle, error) {
	return newAngle(value, AngleTilt)
}

// FromNanoradians creates a new Angle instance from a value in Nanoradians.
func (uf AngleFactory) FromNanoradians(value float64) (*Angle, error) {
	return newAngle(value, AngleNanoradian)
}

// FromMicroradians creates a new Angle instance from a value in Microradians.
func (uf AngleFactory) FromMicroradians(value float64) (*Angle, error) {
	return newAngle(value, AngleMicroradian)
}

// FromMilliradians creates a new Angle instance from a value in Milliradians.
func (uf AngleFactory) FromMilliradians(value float64) (*Angle, error) {
	return newAngle(value, AngleMilliradian)
}

// FromCentiradians creates a new Angle instance from a value in Centiradians.
func (uf AngleFactory) FromCentiradians(value float64) (*Angle, error) {
	return newAngle(value, AngleCentiradian)
}

// FromDeciradians creates a new Angle instance from a value in Deciradians.
func (uf AngleFactory) FromDeciradians(value float64) (*Angle, error) {
	return newAngle(value, AngleDeciradian)
}

// FromNanodegrees creates a new Angle instance from a value in Nanodegrees.
func (uf AngleFactory) FromNanodegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleNanodegree)
}

// FromMicrodegrees creates a new Angle instance from a value in Microdegrees.
func (uf AngleFactory) FromMicrodegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleMicrodegree)
}

// FromMillidegrees creates a new Angle instance from a value in Millidegrees.
func (uf AngleFactory) FromMillidegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleMillidegree)
}


// newAngle creates a new Angle.
func newAngle(value float64, fromUnit AngleUnits) (*Angle, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Angle{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Angle in Degree unit (the base/default quantity).
func (a *Angle) BaseValue() float64 {
	return a.value
}


// Radians returns the Angle value in Radians.
//
// 
func (a *Angle) Radians() float64 {
	if a.radiansLazy != nil {
		return *a.radiansLazy
	}
	radians := a.convertFromBase(AngleRadian)
	a.radiansLazy = &radians
	return radians
}

// Degrees returns the Angle value in Degrees.
//
// 
func (a *Angle) Degrees() float64 {
	if a.degreesLazy != nil {
		return *a.degreesLazy
	}
	degrees := a.convertFromBase(AngleDegree)
	a.degreesLazy = &degrees
	return degrees
}

// Arcminutes returns the Angle value in Arcminutes.
//
// 
func (a *Angle) Arcminutes() float64 {
	if a.arcminutesLazy != nil {
		return *a.arcminutesLazy
	}
	arcminutes := a.convertFromBase(AngleArcminute)
	a.arcminutesLazy = &arcminutes
	return arcminutes
}

// Arcseconds returns the Angle value in Arcseconds.
//
// 
func (a *Angle) Arcseconds() float64 {
	if a.arcsecondsLazy != nil {
		return *a.arcsecondsLazy
	}
	arcseconds := a.convertFromBase(AngleArcsecond)
	a.arcsecondsLazy = &arcseconds
	return arcseconds
}

// Gradians returns the Angle value in Gradians.
//
// 
func (a *Angle) Gradians() float64 {
	if a.gradiansLazy != nil {
		return *a.gradiansLazy
	}
	gradians := a.convertFromBase(AngleGradian)
	a.gradiansLazy = &gradians
	return gradians
}

// NatoMils returns the Angle value in NatoMils.
//
// 
func (a *Angle) NatoMils() float64 {
	if a.nato_milsLazy != nil {
		return *a.nato_milsLazy
	}
	nato_mils := a.convertFromBase(AngleNatoMil)
	a.nato_milsLazy = &nato_mils
	return nato_mils
}

// Revolutions returns the Angle value in Revolutions.
//
// 
func (a *Angle) Revolutions() float64 {
	if a.revolutionsLazy != nil {
		return *a.revolutionsLazy
	}
	revolutions := a.convertFromBase(AngleRevolution)
	a.revolutionsLazy = &revolutions
	return revolutions
}

// Tilt returns the Angle value in Tilt.
//
// 
func (a *Angle) Tilt() float64 {
	if a.tiltLazy != nil {
		return *a.tiltLazy
	}
	tilt := a.convertFromBase(AngleTilt)
	a.tiltLazy = &tilt
	return tilt
}

// Nanoradians returns the Angle value in Nanoradians.
//
// 
func (a *Angle) Nanoradians() float64 {
	if a.nanoradiansLazy != nil {
		return *a.nanoradiansLazy
	}
	nanoradians := a.convertFromBase(AngleNanoradian)
	a.nanoradiansLazy = &nanoradians
	return nanoradians
}

// Microradians returns the Angle value in Microradians.
//
// 
func (a *Angle) Microradians() float64 {
	if a.microradiansLazy != nil {
		return *a.microradiansLazy
	}
	microradians := a.convertFromBase(AngleMicroradian)
	a.microradiansLazy = &microradians
	return microradians
}

// Milliradians returns the Angle value in Milliradians.
//
// 
func (a *Angle) Milliradians() float64 {
	if a.milliradiansLazy != nil {
		return *a.milliradiansLazy
	}
	milliradians := a.convertFromBase(AngleMilliradian)
	a.milliradiansLazy = &milliradians
	return milliradians
}

// Centiradians returns the Angle value in Centiradians.
//
// 
func (a *Angle) Centiradians() float64 {
	if a.centiradiansLazy != nil {
		return *a.centiradiansLazy
	}
	centiradians := a.convertFromBase(AngleCentiradian)
	a.centiradiansLazy = &centiradians
	return centiradians
}

// Deciradians returns the Angle value in Deciradians.
//
// 
func (a *Angle) Deciradians() float64 {
	if a.deciradiansLazy != nil {
		return *a.deciradiansLazy
	}
	deciradians := a.convertFromBase(AngleDeciradian)
	a.deciradiansLazy = &deciradians
	return deciradians
}

// Nanodegrees returns the Angle value in Nanodegrees.
//
// 
func (a *Angle) Nanodegrees() float64 {
	if a.nanodegreesLazy != nil {
		return *a.nanodegreesLazy
	}
	nanodegrees := a.convertFromBase(AngleNanodegree)
	a.nanodegreesLazy = &nanodegrees
	return nanodegrees
}

// Microdegrees returns the Angle value in Microdegrees.
//
// 
func (a *Angle) Microdegrees() float64 {
	if a.microdegreesLazy != nil {
		return *a.microdegreesLazy
	}
	microdegrees := a.convertFromBase(AngleMicrodegree)
	a.microdegreesLazy = &microdegrees
	return microdegrees
}

// Millidegrees returns the Angle value in Millidegrees.
//
// 
func (a *Angle) Millidegrees() float64 {
	if a.millidegreesLazy != nil {
		return *a.millidegreesLazy
	}
	millidegrees := a.convertFromBase(AngleMillidegree)
	a.millidegreesLazy = &millidegrees
	return millidegrees
}


// ToDto creates a AngleDto representation from the Angle instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Degree by default.
func (a *Angle) ToDto(holdInUnit *AngleUnits) AngleDto {
	if holdInUnit == nil {
		defaultUnit := AngleDegree // Default value
		holdInUnit = &defaultUnit
	}

	return AngleDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Angle instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Degree by default.
func (a *Angle) ToDtoJSON(holdInUnit *AngleUnits) ([]byte, error) {
	// Convert to AngleDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Angle to a specific unit value.
// The function uses the provided unit type (AngleUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Angle) Convert(toUnit AngleUnits) float64 {
	switch toUnit { 
    case AngleRadian:
		return a.Radians()
    case AngleDegree:
		return a.Degrees()
    case AngleArcminute:
		return a.Arcminutes()
    case AngleArcsecond:
		return a.Arcseconds()
    case AngleGradian:
		return a.Gradians()
    case AngleNatoMil:
		return a.NatoMils()
    case AngleRevolution:
		return a.Revolutions()
    case AngleTilt:
		return a.Tilt()
    case AngleNanoradian:
		return a.Nanoradians()
    case AngleMicroradian:
		return a.Microradians()
    case AngleMilliradian:
		return a.Milliradians()
    case AngleCentiradian:
		return a.Centiradians()
    case AngleDeciradian:
		return a.Deciradians()
    case AngleNanodegree:
		return a.Nanodegrees()
    case AngleMicrodegree:
		return a.Microdegrees()
    case AngleMillidegree:
		return a.Millidegrees()
	default:
		return math.NaN()
	}
}

func (a *Angle) convertFromBase(toUnit AngleUnits) float64 {
    value := a.value
	switch toUnit { 
	case AngleRadian:
		return (value / 180 * math.Pi) 
	case AngleDegree:
		return (value) 
	case AngleArcminute:
		return (value * 60) 
	case AngleArcsecond:
		return (value * 3600) 
	case AngleGradian:
		return (value / 0.9) 
	case AngleNatoMil:
		return (value * 160 / 9) 
	case AngleRevolution:
		return (value / 360) 
	case AngleTilt:
		return (math.Sin(value / 180 * math.Pi)) 
	case AngleNanoradian:
		return ((value / 180 * math.Pi) / 1e-09) 
	case AngleMicroradian:
		return ((value / 180 * math.Pi) / 1e-06) 
	case AngleMilliradian:
		return ((value / 180 * math.Pi) / 0.001) 
	case AngleCentiradian:
		return ((value / 180 * math.Pi) / 0.01) 
	case AngleDeciradian:
		return ((value / 180 * math.Pi) / 0.1) 
	case AngleNanodegree:
		return ((value) / 1e-09) 
	case AngleMicrodegree:
		return ((value) / 1e-06) 
	case AngleMillidegree:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Angle) convertToBase(value float64, fromUnit AngleUnits) float64 {
	switch fromUnit { 
	case AngleRadian:
		return (value * 180 / math.Pi) 
	case AngleDegree:
		return (value) 
	case AngleArcminute:
		return (value / 60) 
	case AngleArcsecond:
		return (value / 3600) 
	case AngleGradian:
		return (value * 0.9) 
	case AngleNatoMil:
		return (value * 9 / 160) 
	case AngleRevolution:
		return (value * 360) 
	case AngleTilt:
		return (math.Asin(value) * 180 / math.Pi) 
	case AngleNanoradian:
		return ((value * 180 / math.Pi) * 1e-09) 
	case AngleMicroradian:
		return ((value * 180 / math.Pi) * 1e-06) 
	case AngleMilliradian:
		return ((value * 180 / math.Pi) * 0.001) 
	case AngleCentiradian:
		return ((value * 180 / math.Pi) * 0.01) 
	case AngleDeciradian:
		return ((value * 180 / math.Pi) * 0.1) 
	case AngleNanodegree:
		return ((value) * 1e-09) 
	case AngleMicrodegree:
		return ((value) * 1e-06) 
	case AngleMillidegree:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Angle in the default unit (Degree),
// formatted to two decimal places.
func (a Angle) String() string {
	return a.ToString(AngleDegree, 2)
}

// ToString formats the Angle value as a string with the specified unit and fractional digits.
// It converts the Angle to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Angle value will be converted (e.g., Degree))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Angle with the unit abbreviation.
func (a *Angle) ToString(unit AngleUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Angle) getUnitAbbreviation(unit AngleUnits) string {
	switch unit { 
	case AngleRadian:
		return "rad" 
	case AngleDegree:
		return "°" 
	case AngleArcminute:
		return "'" 
	case AngleArcsecond:
		return "″" 
	case AngleGradian:
		return "g" 
	case AngleNatoMil:
		return "mil" 
	case AngleRevolution:
		return "r" 
	case AngleTilt:
		return "sin(θ)" 
	case AngleNanoradian:
		return "nrad" 
	case AngleMicroradian:
		return "μrad" 
	case AngleMilliradian:
		return "mrad" 
	case AngleCentiradian:
		return "crad" 
	case AngleDeciradian:
		return "drad" 
	case AngleNanodegree:
		return "n°" 
	case AngleMicrodegree:
		return "μ°" 
	case AngleMillidegree:
		return "m°" 
	default:
		return ""
	}
}

// Equals checks if the given Angle is equal to the current Angle.
//
// Parameters:
//    other: The Angle to compare against.
//
// Returns:
//    bool: Returns true if both Angle are equal, false otherwise.
func (a *Angle) Equals(other *Angle) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Angle with another Angle.
// It returns -1 if the current Angle is less than the other Angle, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Angle to compare against.
//
// Returns:
//    int: -1 if the current Angle is less, 1 if greater, and 0 if equal.
func (a *Angle) CompareTo(other *Angle) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Angle to the current Angle and returns the result.
// The result is a new Angle instance with the sum of the values.
//
// Parameters:
//    other: The Angle to add to the current Angle.
//
// Returns:
//    *Angle: A new Angle instance representing the sum of both Angle.
func (a *Angle) Add(other *Angle) *Angle {
	return &Angle{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Angle from the current Angle and returns the result.
// The result is a new Angle instance with the difference of the values.
//
// Parameters:
//    other: The Angle to subtract from the current Angle.
//
// Returns:
//    *Angle: A new Angle instance representing the difference of both Angle.
func (a *Angle) Subtract(other *Angle) *Angle {
	return &Angle{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Angle by the given Angle and returns the result.
// The result is a new Angle instance with the product of the values.
//
// Parameters:
//    other: The Angle to multiply with the current Angle.
//
// Returns:
//    *Angle: A new Angle instance representing the product of both Angle.
func (a *Angle) Multiply(other *Angle) *Angle {
	return &Angle{value: a.value * other.BaseValue()}
}

// Divide divides the current Angle by the given Angle and returns the result.
// The result is a new Angle instance with the quotient of the values.
//
// Parameters:
//    other: The Angle to divide the current Angle by.
//
// Returns:
//    *Angle: A new Angle instance representing the quotient of both Angle.
func (a *Angle) Divide(other *Angle) *Angle {
	return &Angle{value: a.value / other.BaseValue()}
}