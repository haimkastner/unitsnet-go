// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalSpeedUnits enumeration
type RotationalSpeedUnits string

const (
    
        // 
        RotationalSpeedRadianPerSecond RotationalSpeedUnits = "RadianPerSecond"
        // 
        RotationalSpeedDegreePerSecond RotationalSpeedUnits = "DegreePerSecond"
        // 
        RotationalSpeedDegreePerMinute RotationalSpeedUnits = "DegreePerMinute"
        // 
        RotationalSpeedRevolutionPerSecond RotationalSpeedUnits = "RevolutionPerSecond"
        // 
        RotationalSpeedRevolutionPerMinute RotationalSpeedUnits = "RevolutionPerMinute"
        // 
        RotationalSpeedNanoradianPerSecond RotationalSpeedUnits = "NanoradianPerSecond"
        // 
        RotationalSpeedMicroradianPerSecond RotationalSpeedUnits = "MicroradianPerSecond"
        // 
        RotationalSpeedMilliradianPerSecond RotationalSpeedUnits = "MilliradianPerSecond"
        // 
        RotationalSpeedCentiradianPerSecond RotationalSpeedUnits = "CentiradianPerSecond"
        // 
        RotationalSpeedDeciradianPerSecond RotationalSpeedUnits = "DeciradianPerSecond"
        // 
        RotationalSpeedNanodegreePerSecond RotationalSpeedUnits = "NanodegreePerSecond"
        // 
        RotationalSpeedMicrodegreePerSecond RotationalSpeedUnits = "MicrodegreePerSecond"
        // 
        RotationalSpeedMillidegreePerSecond RotationalSpeedUnits = "MillidegreePerSecond"
)

// RotationalSpeedDto represents an RotationalSpeed
type RotationalSpeedDto struct {
	Value float64
	Unit  RotationalSpeedUnits
}

// RotationalSpeedDtoFactory struct to group related functions
type RotationalSpeedDtoFactory struct{}

func (udf RotationalSpeedDtoFactory) FromJSON(data []byte) (*RotationalSpeedDto, error) {
	a := RotationalSpeedDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RotationalSpeedDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RotationalSpeed struct
type RotationalSpeed struct {
	value       float64
    
    radians_per_secondLazy *float64 
    degrees_per_secondLazy *float64 
    degrees_per_minuteLazy *float64 
    revolutions_per_secondLazy *float64 
    revolutions_per_minuteLazy *float64 
    nanoradians_per_secondLazy *float64 
    microradians_per_secondLazy *float64 
    milliradians_per_secondLazy *float64 
    centiradians_per_secondLazy *float64 
    deciradians_per_secondLazy *float64 
    nanodegrees_per_secondLazy *float64 
    microdegrees_per_secondLazy *float64 
    millidegrees_per_secondLazy *float64 
}

// RotationalSpeedFactory struct to group related functions
type RotationalSpeedFactory struct{}

func (uf RotationalSpeedFactory) CreateRotationalSpeed(value float64, unit RotationalSpeedUnits) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, unit)
}

func (uf RotationalSpeedFactory) FromDto(dto RotationalSpeedDto) (*RotationalSpeed, error) {
	return newRotationalSpeed(dto.Value, dto.Unit)
}

func (uf RotationalSpeedFactory) FromDtoJSON(data []byte) (*RotationalSpeed, error) {
	unitDto, err := RotationalSpeedDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RotationalSpeedFactory{}.FromDto(*unitDto)
}


// FromRadianPerSecond creates a new RotationalSpeed instance from RadianPerSecond.
func (uf RotationalSpeedFactory) FromRadiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRadianPerSecond)
}

// FromDegreePerSecond creates a new RotationalSpeed instance from DegreePerSecond.
func (uf RotationalSpeedFactory) FromDegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDegreePerSecond)
}

// FromDegreePerMinute creates a new RotationalSpeed instance from DegreePerMinute.
func (uf RotationalSpeedFactory) FromDegreesPerMinute(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDegreePerMinute)
}

// FromRevolutionPerSecond creates a new RotationalSpeed instance from RevolutionPerSecond.
func (uf RotationalSpeedFactory) FromRevolutionsPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRevolutionPerSecond)
}

// FromRevolutionPerMinute creates a new RotationalSpeed instance from RevolutionPerMinute.
func (uf RotationalSpeedFactory) FromRevolutionsPerMinute(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRevolutionPerMinute)
}

// FromNanoradianPerSecond creates a new RotationalSpeed instance from NanoradianPerSecond.
func (uf RotationalSpeedFactory) FromNanoradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedNanoradianPerSecond)
}

// FromMicroradianPerSecond creates a new RotationalSpeed instance from MicroradianPerSecond.
func (uf RotationalSpeedFactory) FromMicroradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMicroradianPerSecond)
}

// FromMilliradianPerSecond creates a new RotationalSpeed instance from MilliradianPerSecond.
func (uf RotationalSpeedFactory) FromMilliradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMilliradianPerSecond)
}

// FromCentiradianPerSecond creates a new RotationalSpeed instance from CentiradianPerSecond.
func (uf RotationalSpeedFactory) FromCentiradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedCentiradianPerSecond)
}

// FromDeciradianPerSecond creates a new RotationalSpeed instance from DeciradianPerSecond.
func (uf RotationalSpeedFactory) FromDeciradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDeciradianPerSecond)
}

// FromNanodegreePerSecond creates a new RotationalSpeed instance from NanodegreePerSecond.
func (uf RotationalSpeedFactory) FromNanodegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedNanodegreePerSecond)
}

// FromMicrodegreePerSecond creates a new RotationalSpeed instance from MicrodegreePerSecond.
func (uf RotationalSpeedFactory) FromMicrodegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMicrodegreePerSecond)
}

// FromMillidegreePerSecond creates a new RotationalSpeed instance from MillidegreePerSecond.
func (uf RotationalSpeedFactory) FromMillidegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMillidegreePerSecond)
}




// newRotationalSpeed creates a new RotationalSpeed.
func newRotationalSpeed(value float64, fromUnit RotationalSpeedUnits) (*RotationalSpeed, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalSpeed{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalSpeed in RadianPerSecond.
func (a *RotationalSpeed) BaseValue() float64 {
	return a.value
}


// RadianPerSecond returns the value in RadianPerSecond.
func (a *RotationalSpeed) RadiansPerSecond() float64 {
	if a.radians_per_secondLazy != nil {
		return *a.radians_per_secondLazy
	}
	radians_per_second := a.convertFromBase(RotationalSpeedRadianPerSecond)
	a.radians_per_secondLazy = &radians_per_second
	return radians_per_second
}

// DegreePerSecond returns the value in DegreePerSecond.
func (a *RotationalSpeed) DegreesPerSecond() float64 {
	if a.degrees_per_secondLazy != nil {
		return *a.degrees_per_secondLazy
	}
	degrees_per_second := a.convertFromBase(RotationalSpeedDegreePerSecond)
	a.degrees_per_secondLazy = &degrees_per_second
	return degrees_per_second
}

// DegreePerMinute returns the value in DegreePerMinute.
func (a *RotationalSpeed) DegreesPerMinute() float64 {
	if a.degrees_per_minuteLazy != nil {
		return *a.degrees_per_minuteLazy
	}
	degrees_per_minute := a.convertFromBase(RotationalSpeedDegreePerMinute)
	a.degrees_per_minuteLazy = &degrees_per_minute
	return degrees_per_minute
}

// RevolutionPerSecond returns the value in RevolutionPerSecond.
func (a *RotationalSpeed) RevolutionsPerSecond() float64 {
	if a.revolutions_per_secondLazy != nil {
		return *a.revolutions_per_secondLazy
	}
	revolutions_per_second := a.convertFromBase(RotationalSpeedRevolutionPerSecond)
	a.revolutions_per_secondLazy = &revolutions_per_second
	return revolutions_per_second
}

// RevolutionPerMinute returns the value in RevolutionPerMinute.
func (a *RotationalSpeed) RevolutionsPerMinute() float64 {
	if a.revolutions_per_minuteLazy != nil {
		return *a.revolutions_per_minuteLazy
	}
	revolutions_per_minute := a.convertFromBase(RotationalSpeedRevolutionPerMinute)
	a.revolutions_per_minuteLazy = &revolutions_per_minute
	return revolutions_per_minute
}

// NanoradianPerSecond returns the value in NanoradianPerSecond.
func (a *RotationalSpeed) NanoradiansPerSecond() float64 {
	if a.nanoradians_per_secondLazy != nil {
		return *a.nanoradians_per_secondLazy
	}
	nanoradians_per_second := a.convertFromBase(RotationalSpeedNanoradianPerSecond)
	a.nanoradians_per_secondLazy = &nanoradians_per_second
	return nanoradians_per_second
}

// MicroradianPerSecond returns the value in MicroradianPerSecond.
func (a *RotationalSpeed) MicroradiansPerSecond() float64 {
	if a.microradians_per_secondLazy != nil {
		return *a.microradians_per_secondLazy
	}
	microradians_per_second := a.convertFromBase(RotationalSpeedMicroradianPerSecond)
	a.microradians_per_secondLazy = &microradians_per_second
	return microradians_per_second
}

// MilliradianPerSecond returns the value in MilliradianPerSecond.
func (a *RotationalSpeed) MilliradiansPerSecond() float64 {
	if a.milliradians_per_secondLazy != nil {
		return *a.milliradians_per_secondLazy
	}
	milliradians_per_second := a.convertFromBase(RotationalSpeedMilliradianPerSecond)
	a.milliradians_per_secondLazy = &milliradians_per_second
	return milliradians_per_second
}

// CentiradianPerSecond returns the value in CentiradianPerSecond.
func (a *RotationalSpeed) CentiradiansPerSecond() float64 {
	if a.centiradians_per_secondLazy != nil {
		return *a.centiradians_per_secondLazy
	}
	centiradians_per_second := a.convertFromBase(RotationalSpeedCentiradianPerSecond)
	a.centiradians_per_secondLazy = &centiradians_per_second
	return centiradians_per_second
}

// DeciradianPerSecond returns the value in DeciradianPerSecond.
func (a *RotationalSpeed) DeciradiansPerSecond() float64 {
	if a.deciradians_per_secondLazy != nil {
		return *a.deciradians_per_secondLazy
	}
	deciradians_per_second := a.convertFromBase(RotationalSpeedDeciradianPerSecond)
	a.deciradians_per_secondLazy = &deciradians_per_second
	return deciradians_per_second
}

// NanodegreePerSecond returns the value in NanodegreePerSecond.
func (a *RotationalSpeed) NanodegreesPerSecond() float64 {
	if a.nanodegrees_per_secondLazy != nil {
		return *a.nanodegrees_per_secondLazy
	}
	nanodegrees_per_second := a.convertFromBase(RotationalSpeedNanodegreePerSecond)
	a.nanodegrees_per_secondLazy = &nanodegrees_per_second
	return nanodegrees_per_second
}

// MicrodegreePerSecond returns the value in MicrodegreePerSecond.
func (a *RotationalSpeed) MicrodegreesPerSecond() float64 {
	if a.microdegrees_per_secondLazy != nil {
		return *a.microdegrees_per_secondLazy
	}
	microdegrees_per_second := a.convertFromBase(RotationalSpeedMicrodegreePerSecond)
	a.microdegrees_per_secondLazy = &microdegrees_per_second
	return microdegrees_per_second
}

// MillidegreePerSecond returns the value in MillidegreePerSecond.
func (a *RotationalSpeed) MillidegreesPerSecond() float64 {
	if a.millidegrees_per_secondLazy != nil {
		return *a.millidegrees_per_secondLazy
	}
	millidegrees_per_second := a.convertFromBase(RotationalSpeedMillidegreePerSecond)
	a.millidegrees_per_secondLazy = &millidegrees_per_second
	return millidegrees_per_second
}


// ToDto creates an RotationalSpeedDto representation.
func (a *RotationalSpeed) ToDto(holdInUnit *RotationalSpeedUnits) RotationalSpeedDto {
	if holdInUnit == nil {
		defaultUnit := RotationalSpeedRadianPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalSpeedDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RotationalSpeedDto representation.
func (a *RotationalSpeed) ToDtoJSON(holdInUnit *RotationalSpeedUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RotationalSpeed to a specific unit value.
func (a *RotationalSpeed) Convert(toUnit RotationalSpeedUnits) float64 {
	switch toUnit { 
    case RotationalSpeedRadianPerSecond:
		return a.RadiansPerSecond()
    case RotationalSpeedDegreePerSecond:
		return a.DegreesPerSecond()
    case RotationalSpeedDegreePerMinute:
		return a.DegreesPerMinute()
    case RotationalSpeedRevolutionPerSecond:
		return a.RevolutionsPerSecond()
    case RotationalSpeedRevolutionPerMinute:
		return a.RevolutionsPerMinute()
    case RotationalSpeedNanoradianPerSecond:
		return a.NanoradiansPerSecond()
    case RotationalSpeedMicroradianPerSecond:
		return a.MicroradiansPerSecond()
    case RotationalSpeedMilliradianPerSecond:
		return a.MilliradiansPerSecond()
    case RotationalSpeedCentiradianPerSecond:
		return a.CentiradiansPerSecond()
    case RotationalSpeedDeciradianPerSecond:
		return a.DeciradiansPerSecond()
    case RotationalSpeedNanodegreePerSecond:
		return a.NanodegreesPerSecond()
    case RotationalSpeedMicrodegreePerSecond:
		return a.MicrodegreesPerSecond()
    case RotationalSpeedMillidegreePerSecond:
		return a.MillidegreesPerSecond()
	default:
		return 0
	}
}

func (a *RotationalSpeed) convertFromBase(toUnit RotationalSpeedUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalSpeedRadianPerSecond:
		return (value) 
	case RotationalSpeedDegreePerSecond:
		return ((180 / math.Pi) * value) 
	case RotationalSpeedDegreePerMinute:
		return ((180 * 60 / math.Pi) * value) 
	case RotationalSpeedRevolutionPerSecond:
		return (value / 6.2831853072) 
	case RotationalSpeedRevolutionPerMinute:
		return ((value / 6.2831853072) * 60) 
	case RotationalSpeedNanoradianPerSecond:
		return ((value) / 1e-09) 
	case RotationalSpeedMicroradianPerSecond:
		return ((value) / 1e-06) 
	case RotationalSpeedMilliradianPerSecond:
		return ((value) / 0.001) 
	case RotationalSpeedCentiradianPerSecond:
		return ((value) / 0.01) 
	case RotationalSpeedDeciradianPerSecond:
		return ((value) / 0.1) 
	case RotationalSpeedNanodegreePerSecond:
		return (((180 / math.Pi) * value) / 1e-09) 
	case RotationalSpeedMicrodegreePerSecond:
		return (((180 / math.Pi) * value) / 1e-06) 
	case RotationalSpeedMillidegreePerSecond:
		return (((180 / math.Pi) * value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RotationalSpeed) convertToBase(value float64, fromUnit RotationalSpeedUnits) float64 {
	switch fromUnit { 
	case RotationalSpeedRadianPerSecond:
		return (value) 
	case RotationalSpeedDegreePerSecond:
		return ((math.Pi / 180) * value) 
	case RotationalSpeedDegreePerMinute:
		return ((math.Pi / (180 * 60)) * value) 
	case RotationalSpeedRevolutionPerSecond:
		return (value * 6.2831853072) 
	case RotationalSpeedRevolutionPerMinute:
		return ((value * 6.2831853072) / 60) 
	case RotationalSpeedNanoradianPerSecond:
		return ((value) * 1e-09) 
	case RotationalSpeedMicroradianPerSecond:
		return ((value) * 1e-06) 
	case RotationalSpeedMilliradianPerSecond:
		return ((value) * 0.001) 
	case RotationalSpeedCentiradianPerSecond:
		return ((value) * 0.01) 
	case RotationalSpeedDeciradianPerSecond:
		return ((value) * 0.1) 
	case RotationalSpeedNanodegreePerSecond:
		return (((math.Pi / 180) * value) * 1e-09) 
	case RotationalSpeedMicrodegreePerSecond:
		return (((math.Pi / 180) * value) * 1e-06) 
	case RotationalSpeedMillidegreePerSecond:
		return (((math.Pi / 180) * value) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RotationalSpeed) String() string {
	return a.ToString(RotationalSpeedRadianPerSecond, 2)
}

// ToString formats the RotationalSpeed to string.
// fractionalDigits -1 for not mention
func (a *RotationalSpeed) ToString(unit RotationalSpeedUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalSpeed) getUnitAbbreviation(unit RotationalSpeedUnits) string {
	switch unit { 
	case RotationalSpeedRadianPerSecond:
		return "rad/s" 
	case RotationalSpeedDegreePerSecond:
		return "°/s" 
	case RotationalSpeedDegreePerMinute:
		return "°/min" 
	case RotationalSpeedRevolutionPerSecond:
		return "r/s" 
	case RotationalSpeedRevolutionPerMinute:
		return "rpm" 
	case RotationalSpeedNanoradianPerSecond:
		return "nrad/s" 
	case RotationalSpeedMicroradianPerSecond:
		return "μrad/s" 
	case RotationalSpeedMilliradianPerSecond:
		return "mrad/s" 
	case RotationalSpeedCentiradianPerSecond:
		return "crad/s" 
	case RotationalSpeedDeciradianPerSecond:
		return "drad/s" 
	case RotationalSpeedNanodegreePerSecond:
		return "n°/s" 
	case RotationalSpeedMicrodegreePerSecond:
		return "μ°/s" 
	case RotationalSpeedMillidegreePerSecond:
		return "m°/s" 
	default:
		return ""
	}
}

// Check if the given RotationalSpeed are equals to the current RotationalSpeed
func (a *RotationalSpeed) Equals(other *RotationalSpeed) bool {
	return a.value == other.BaseValue()
}

// Check if the given RotationalSpeed are equals to the current RotationalSpeed
func (a *RotationalSpeed) CompareTo(other *RotationalSpeed) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RotationalSpeed to the current RotationalSpeed.
func (a *RotationalSpeed) Add(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value + other.BaseValue()}
}

// Subtract the given RotationalSpeed to the current RotationalSpeed.
func (a *RotationalSpeed) Subtract(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value - other.BaseValue()}
}

// Multiply the given RotationalSpeed to the current RotationalSpeed.
func (a *RotationalSpeed) Multiply(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value * other.BaseValue()}
}

// Divide the given RotationalSpeed to the current RotationalSpeed.
func (a *RotationalSpeed) Divide(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value / other.BaseValue()}
}