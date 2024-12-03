package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AngleUnits enumeration
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

// AngleDto represents an Angle
type AngleDto struct {
	Value float64
	Unit  AngleUnits
}

// AngleDtoFactory struct to group related functions
type AngleDtoFactory struct{}

func (udf AngleDtoFactory) FromJSON(data []byte) (*AngleDto, error) {
	a := AngleDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AngleDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Angle struct
type Angle struct {
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

// AngleFactory struct to group related functions
type AngleFactory struct{}

func (uf AngleFactory) CreateAngle(value float64, unit AngleUnits) (*Angle, error) {
	return newAngle(value, unit)
}

func (uf AngleFactory) FromDto(dto AngleDto) (*Angle, error) {
	return newAngle(dto.Value, dto.Unit)
}

func (uf AngleFactory) FromDtoJSON(data []byte) (*Angle, error) {
	unitDto, err := AngleDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AngleFactory{}.FromDto(*unitDto)
}


// FromRadian creates a new Angle instance from Radian.
func (uf AngleFactory) FromRadians(value float64) (*Angle, error) {
	return newAngle(value, AngleRadian)
}

// FromDegree creates a new Angle instance from Degree.
func (uf AngleFactory) FromDegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleDegree)
}

// FromArcminute creates a new Angle instance from Arcminute.
func (uf AngleFactory) FromArcminutes(value float64) (*Angle, error) {
	return newAngle(value, AngleArcminute)
}

// FromArcsecond creates a new Angle instance from Arcsecond.
func (uf AngleFactory) FromArcseconds(value float64) (*Angle, error) {
	return newAngle(value, AngleArcsecond)
}

// FromGradian creates a new Angle instance from Gradian.
func (uf AngleFactory) FromGradians(value float64) (*Angle, error) {
	return newAngle(value, AngleGradian)
}

// FromNatoMil creates a new Angle instance from NatoMil.
func (uf AngleFactory) FromNatoMils(value float64) (*Angle, error) {
	return newAngle(value, AngleNatoMil)
}

// FromRevolution creates a new Angle instance from Revolution.
func (uf AngleFactory) FromRevolutions(value float64) (*Angle, error) {
	return newAngle(value, AngleRevolution)
}

// FromTilt creates a new Angle instance from Tilt.
func (uf AngleFactory) FromTilt(value float64) (*Angle, error) {
	return newAngle(value, AngleTilt)
}

// FromNanoradian creates a new Angle instance from Nanoradian.
func (uf AngleFactory) FromNanoradians(value float64) (*Angle, error) {
	return newAngle(value, AngleNanoradian)
}

// FromMicroradian creates a new Angle instance from Microradian.
func (uf AngleFactory) FromMicroradians(value float64) (*Angle, error) {
	return newAngle(value, AngleMicroradian)
}

// FromMilliradian creates a new Angle instance from Milliradian.
func (uf AngleFactory) FromMilliradians(value float64) (*Angle, error) {
	return newAngle(value, AngleMilliradian)
}

// FromCentiradian creates a new Angle instance from Centiradian.
func (uf AngleFactory) FromCentiradians(value float64) (*Angle, error) {
	return newAngle(value, AngleCentiradian)
}

// FromDeciradian creates a new Angle instance from Deciradian.
func (uf AngleFactory) FromDeciradians(value float64) (*Angle, error) {
	return newAngle(value, AngleDeciradian)
}

// FromNanodegree creates a new Angle instance from Nanodegree.
func (uf AngleFactory) FromNanodegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleNanodegree)
}

// FromMicrodegree creates a new Angle instance from Microdegree.
func (uf AngleFactory) FromMicrodegrees(value float64) (*Angle, error) {
	return newAngle(value, AngleMicrodegree)
}

// FromMillidegree creates a new Angle instance from Millidegree.
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

// BaseValue returns the base value of Angle in Degree.
func (a *Angle) BaseValue() float64 {
	return a.value
}


// Radian returns the value in Radian.
func (a *Angle) Radians() float64 {
	if a.radiansLazy != nil {
		return *a.radiansLazy
	}
	radians := a.convertFromBase(AngleRadian)
	a.radiansLazy = &radians
	return radians
}

// Degree returns the value in Degree.
func (a *Angle) Degrees() float64 {
	if a.degreesLazy != nil {
		return *a.degreesLazy
	}
	degrees := a.convertFromBase(AngleDegree)
	a.degreesLazy = &degrees
	return degrees
}

// Arcminute returns the value in Arcminute.
func (a *Angle) Arcminutes() float64 {
	if a.arcminutesLazy != nil {
		return *a.arcminutesLazy
	}
	arcminutes := a.convertFromBase(AngleArcminute)
	a.arcminutesLazy = &arcminutes
	return arcminutes
}

// Arcsecond returns the value in Arcsecond.
func (a *Angle) Arcseconds() float64 {
	if a.arcsecondsLazy != nil {
		return *a.arcsecondsLazy
	}
	arcseconds := a.convertFromBase(AngleArcsecond)
	a.arcsecondsLazy = &arcseconds
	return arcseconds
}

// Gradian returns the value in Gradian.
func (a *Angle) Gradians() float64 {
	if a.gradiansLazy != nil {
		return *a.gradiansLazy
	}
	gradians := a.convertFromBase(AngleGradian)
	a.gradiansLazy = &gradians
	return gradians
}

// NatoMil returns the value in NatoMil.
func (a *Angle) NatoMils() float64 {
	if a.nato_milsLazy != nil {
		return *a.nato_milsLazy
	}
	nato_mils := a.convertFromBase(AngleNatoMil)
	a.nato_milsLazy = &nato_mils
	return nato_mils
}

// Revolution returns the value in Revolution.
func (a *Angle) Revolutions() float64 {
	if a.revolutionsLazy != nil {
		return *a.revolutionsLazy
	}
	revolutions := a.convertFromBase(AngleRevolution)
	a.revolutionsLazy = &revolutions
	return revolutions
}

// Tilt returns the value in Tilt.
func (a *Angle) Tilt() float64 {
	if a.tiltLazy != nil {
		return *a.tiltLazy
	}
	tilt := a.convertFromBase(AngleTilt)
	a.tiltLazy = &tilt
	return tilt
}

// Nanoradian returns the value in Nanoradian.
func (a *Angle) Nanoradians() float64 {
	if a.nanoradiansLazy != nil {
		return *a.nanoradiansLazy
	}
	nanoradians := a.convertFromBase(AngleNanoradian)
	a.nanoradiansLazy = &nanoradians
	return nanoradians
}

// Microradian returns the value in Microradian.
func (a *Angle) Microradians() float64 {
	if a.microradiansLazy != nil {
		return *a.microradiansLazy
	}
	microradians := a.convertFromBase(AngleMicroradian)
	a.microradiansLazy = &microradians
	return microradians
}

// Milliradian returns the value in Milliradian.
func (a *Angle) Milliradians() float64 {
	if a.milliradiansLazy != nil {
		return *a.milliradiansLazy
	}
	milliradians := a.convertFromBase(AngleMilliradian)
	a.milliradiansLazy = &milliradians
	return milliradians
}

// Centiradian returns the value in Centiradian.
func (a *Angle) Centiradians() float64 {
	if a.centiradiansLazy != nil {
		return *a.centiradiansLazy
	}
	centiradians := a.convertFromBase(AngleCentiradian)
	a.centiradiansLazy = &centiradians
	return centiradians
}

// Deciradian returns the value in Deciradian.
func (a *Angle) Deciradians() float64 {
	if a.deciradiansLazy != nil {
		return *a.deciradiansLazy
	}
	deciradians := a.convertFromBase(AngleDeciradian)
	a.deciradiansLazy = &deciradians
	return deciradians
}

// Nanodegree returns the value in Nanodegree.
func (a *Angle) Nanodegrees() float64 {
	if a.nanodegreesLazy != nil {
		return *a.nanodegreesLazy
	}
	nanodegrees := a.convertFromBase(AngleNanodegree)
	a.nanodegreesLazy = &nanodegrees
	return nanodegrees
}

// Microdegree returns the value in Microdegree.
func (a *Angle) Microdegrees() float64 {
	if a.microdegreesLazy != nil {
		return *a.microdegreesLazy
	}
	microdegrees := a.convertFromBase(AngleMicrodegree)
	a.microdegreesLazy = &microdegrees
	return microdegrees
}

// Millidegree returns the value in Millidegree.
func (a *Angle) Millidegrees() float64 {
	if a.millidegreesLazy != nil {
		return *a.millidegreesLazy
	}
	millidegrees := a.convertFromBase(AngleMillidegree)
	a.millidegreesLazy = &millidegrees
	return millidegrees
}


// ToDto creates an AngleDto representation.
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

// ToDtoJSON creates an AngleDto representation.
func (a *Angle) ToDtoJSON(holdInUnit *AngleUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Angle to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Angle) String() string {
	return a.ToString(AngleDegree, 2)
}

// ToString formats the Angle to string.
// fractionalDigits -1 for not mention
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

// Check if the given Angle are equals to the current Angle
func (a *Angle) Equals(other *Angle) bool {
	return a.value == other.BaseValue()
}

// Check if the given Angle are equals to the current Angle
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

// Add the given Angle to the current Angle.
func (a *Angle) Add(other *Angle) *Angle {
	return &Angle{value: a.value + other.BaseValue()}
}

// Subtract the given Angle to the current Angle.
func (a *Angle) Subtract(other *Angle) *Angle {
	return &Angle{value: a.value - other.BaseValue()}
}

// Multiply the given Angle to the current Angle.
func (a *Angle) Multiply(other *Angle) *Angle {
	return &Angle{value: a.value * other.BaseValue()}
}

// Divide the given Angle to the current Angle.
func (a *Angle) Divide(other *Angle) *Angle {
	return &Angle{value: a.value / other.BaseValue()}
}