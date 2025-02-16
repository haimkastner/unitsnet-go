// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalSpeedUnits defines various units of RotationalSpeed.
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

// RotationalSpeedDto represents a RotationalSpeed measurement with a numerical value and its corresponding unit.
type RotationalSpeedDto struct {
    // Value is the numerical representation of the RotationalSpeed.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RotationalSpeed, as defined in the RotationalSpeedUnits enumeration.
	Unit  RotationalSpeedUnits `json:"unit"`
}

// RotationalSpeedDtoFactory groups methods for creating and serializing RotationalSpeedDto objects.
type RotationalSpeedDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RotationalSpeedDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RotationalSpeedDtoFactory) FromJSON(data []byte) (*RotationalSpeedDto, error) {
	a := RotationalSpeedDto{}

    // Parse JSON into RotationalSpeedDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a RotationalSpeedDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RotationalSpeedDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RotationalSpeed represents a measurement in a of RotationalSpeed.
//
// Rotational speed (sometimes called speed of revolution) is the number of complete rotations, revolutions, cycles, or turns per time unit. Rotational speed is a cyclic frequency, measured in radians per second or in hertz in the SI System by scientists, or in revolutions per minute (rpm or min-1) or revolutions per second in everyday life. The symbol for rotational speed is ω (the Greek lowercase letter "omega").
type RotationalSpeed struct {
	// value is the base measurement stored internally.
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

// RotationalSpeedFactory groups methods for creating RotationalSpeed instances.
type RotationalSpeedFactory struct{}

// CreateRotationalSpeed creates a new RotationalSpeed instance from the given value and unit.
func (uf RotationalSpeedFactory) CreateRotationalSpeed(value float64, unit RotationalSpeedUnits) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, unit)
}

// FromDto converts a RotationalSpeedDto to a RotationalSpeed instance.
func (uf RotationalSpeedFactory) FromDto(dto RotationalSpeedDto) (*RotationalSpeed, error) {
	return newRotationalSpeed(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RotationalSpeed instance.
func (uf RotationalSpeedFactory) FromDtoJSON(data []byte) (*RotationalSpeed, error) {
	unitDto, err := RotationalSpeedDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RotationalSpeedDto from JSON: %w", err)
	}
	return RotationalSpeedFactory{}.FromDto(*unitDto)
}


// FromRadiansPerSecond creates a new RotationalSpeed instance from a value in RadiansPerSecond.
func (uf RotationalSpeedFactory) FromRadiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRadianPerSecond)
}

// FromDegreesPerSecond creates a new RotationalSpeed instance from a value in DegreesPerSecond.
func (uf RotationalSpeedFactory) FromDegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDegreePerSecond)
}

// FromDegreesPerMinute creates a new RotationalSpeed instance from a value in DegreesPerMinute.
func (uf RotationalSpeedFactory) FromDegreesPerMinute(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDegreePerMinute)
}

// FromRevolutionsPerSecond creates a new RotationalSpeed instance from a value in RevolutionsPerSecond.
func (uf RotationalSpeedFactory) FromRevolutionsPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRevolutionPerSecond)
}

// FromRevolutionsPerMinute creates a new RotationalSpeed instance from a value in RevolutionsPerMinute.
func (uf RotationalSpeedFactory) FromRevolutionsPerMinute(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedRevolutionPerMinute)
}

// FromNanoradiansPerSecond creates a new RotationalSpeed instance from a value in NanoradiansPerSecond.
func (uf RotationalSpeedFactory) FromNanoradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedNanoradianPerSecond)
}

// FromMicroradiansPerSecond creates a new RotationalSpeed instance from a value in MicroradiansPerSecond.
func (uf RotationalSpeedFactory) FromMicroradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMicroradianPerSecond)
}

// FromMilliradiansPerSecond creates a new RotationalSpeed instance from a value in MilliradiansPerSecond.
func (uf RotationalSpeedFactory) FromMilliradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMilliradianPerSecond)
}

// FromCentiradiansPerSecond creates a new RotationalSpeed instance from a value in CentiradiansPerSecond.
func (uf RotationalSpeedFactory) FromCentiradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedCentiradianPerSecond)
}

// FromDeciradiansPerSecond creates a new RotationalSpeed instance from a value in DeciradiansPerSecond.
func (uf RotationalSpeedFactory) FromDeciradiansPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedDeciradianPerSecond)
}

// FromNanodegreesPerSecond creates a new RotationalSpeed instance from a value in NanodegreesPerSecond.
func (uf RotationalSpeedFactory) FromNanodegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedNanodegreePerSecond)
}

// FromMicrodegreesPerSecond creates a new RotationalSpeed instance from a value in MicrodegreesPerSecond.
func (uf RotationalSpeedFactory) FromMicrodegreesPerSecond(value float64) (*RotationalSpeed, error) {
	return newRotationalSpeed(value, RotationalSpeedMicrodegreePerSecond)
}

// FromMillidegreesPerSecond creates a new RotationalSpeed instance from a value in MillidegreesPerSecond.
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

// BaseValue returns the base value of RotationalSpeed in RadianPerSecond unit (the base/default quantity).
func (a *RotationalSpeed) BaseValue() float64 {
	return a.value
}


// RadiansPerSecond returns the RotationalSpeed value in RadiansPerSecond.
//
// 
func (a *RotationalSpeed) RadiansPerSecond() float64 {
	if a.radians_per_secondLazy != nil {
		return *a.radians_per_secondLazy
	}
	radians_per_second := a.convertFromBase(RotationalSpeedRadianPerSecond)
	a.radians_per_secondLazy = &radians_per_second
	return radians_per_second
}

// DegreesPerSecond returns the RotationalSpeed value in DegreesPerSecond.
//
// 
func (a *RotationalSpeed) DegreesPerSecond() float64 {
	if a.degrees_per_secondLazy != nil {
		return *a.degrees_per_secondLazy
	}
	degrees_per_second := a.convertFromBase(RotationalSpeedDegreePerSecond)
	a.degrees_per_secondLazy = &degrees_per_second
	return degrees_per_second
}

// DegreesPerMinute returns the RotationalSpeed value in DegreesPerMinute.
//
// 
func (a *RotationalSpeed) DegreesPerMinute() float64 {
	if a.degrees_per_minuteLazy != nil {
		return *a.degrees_per_minuteLazy
	}
	degrees_per_minute := a.convertFromBase(RotationalSpeedDegreePerMinute)
	a.degrees_per_minuteLazy = &degrees_per_minute
	return degrees_per_minute
}

// RevolutionsPerSecond returns the RotationalSpeed value in RevolutionsPerSecond.
//
// 
func (a *RotationalSpeed) RevolutionsPerSecond() float64 {
	if a.revolutions_per_secondLazy != nil {
		return *a.revolutions_per_secondLazy
	}
	revolutions_per_second := a.convertFromBase(RotationalSpeedRevolutionPerSecond)
	a.revolutions_per_secondLazy = &revolutions_per_second
	return revolutions_per_second
}

// RevolutionsPerMinute returns the RotationalSpeed value in RevolutionsPerMinute.
//
// 
func (a *RotationalSpeed) RevolutionsPerMinute() float64 {
	if a.revolutions_per_minuteLazy != nil {
		return *a.revolutions_per_minuteLazy
	}
	revolutions_per_minute := a.convertFromBase(RotationalSpeedRevolutionPerMinute)
	a.revolutions_per_minuteLazy = &revolutions_per_minute
	return revolutions_per_minute
}

// NanoradiansPerSecond returns the RotationalSpeed value in NanoradiansPerSecond.
//
// 
func (a *RotationalSpeed) NanoradiansPerSecond() float64 {
	if a.nanoradians_per_secondLazy != nil {
		return *a.nanoradians_per_secondLazy
	}
	nanoradians_per_second := a.convertFromBase(RotationalSpeedNanoradianPerSecond)
	a.nanoradians_per_secondLazy = &nanoradians_per_second
	return nanoradians_per_second
}

// MicroradiansPerSecond returns the RotationalSpeed value in MicroradiansPerSecond.
//
// 
func (a *RotationalSpeed) MicroradiansPerSecond() float64 {
	if a.microradians_per_secondLazy != nil {
		return *a.microradians_per_secondLazy
	}
	microradians_per_second := a.convertFromBase(RotationalSpeedMicroradianPerSecond)
	a.microradians_per_secondLazy = &microradians_per_second
	return microradians_per_second
}

// MilliradiansPerSecond returns the RotationalSpeed value in MilliradiansPerSecond.
//
// 
func (a *RotationalSpeed) MilliradiansPerSecond() float64 {
	if a.milliradians_per_secondLazy != nil {
		return *a.milliradians_per_secondLazy
	}
	milliradians_per_second := a.convertFromBase(RotationalSpeedMilliradianPerSecond)
	a.milliradians_per_secondLazy = &milliradians_per_second
	return milliradians_per_second
}

// CentiradiansPerSecond returns the RotationalSpeed value in CentiradiansPerSecond.
//
// 
func (a *RotationalSpeed) CentiradiansPerSecond() float64 {
	if a.centiradians_per_secondLazy != nil {
		return *a.centiradians_per_secondLazy
	}
	centiradians_per_second := a.convertFromBase(RotationalSpeedCentiradianPerSecond)
	a.centiradians_per_secondLazy = &centiradians_per_second
	return centiradians_per_second
}

// DeciradiansPerSecond returns the RotationalSpeed value in DeciradiansPerSecond.
//
// 
func (a *RotationalSpeed) DeciradiansPerSecond() float64 {
	if a.deciradians_per_secondLazy != nil {
		return *a.deciradians_per_secondLazy
	}
	deciradians_per_second := a.convertFromBase(RotationalSpeedDeciradianPerSecond)
	a.deciradians_per_secondLazy = &deciradians_per_second
	return deciradians_per_second
}

// NanodegreesPerSecond returns the RotationalSpeed value in NanodegreesPerSecond.
//
// 
func (a *RotationalSpeed) NanodegreesPerSecond() float64 {
	if a.nanodegrees_per_secondLazy != nil {
		return *a.nanodegrees_per_secondLazy
	}
	nanodegrees_per_second := a.convertFromBase(RotationalSpeedNanodegreePerSecond)
	a.nanodegrees_per_secondLazy = &nanodegrees_per_second
	return nanodegrees_per_second
}

// MicrodegreesPerSecond returns the RotationalSpeed value in MicrodegreesPerSecond.
//
// 
func (a *RotationalSpeed) MicrodegreesPerSecond() float64 {
	if a.microdegrees_per_secondLazy != nil {
		return *a.microdegrees_per_secondLazy
	}
	microdegrees_per_second := a.convertFromBase(RotationalSpeedMicrodegreePerSecond)
	a.microdegrees_per_secondLazy = &microdegrees_per_second
	return microdegrees_per_second
}

// MillidegreesPerSecond returns the RotationalSpeed value in MillidegreesPerSecond.
//
// 
func (a *RotationalSpeed) MillidegreesPerSecond() float64 {
	if a.millidegrees_per_secondLazy != nil {
		return *a.millidegrees_per_secondLazy
	}
	millidegrees_per_second := a.convertFromBase(RotationalSpeedMillidegreePerSecond)
	a.millidegrees_per_secondLazy = &millidegrees_per_second
	return millidegrees_per_second
}


// ToDto creates a RotationalSpeedDto representation from the RotationalSpeed instance.
//
// If the provided holdInUnit is nil, the value will be repesented by RadianPerSecond by default.
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

// ToDtoJSON creates a JSON representation of the RotationalSpeed instance.
//
// If the provided holdInUnit is nil, the value will be repesented by RadianPerSecond by default.
func (a *RotationalSpeed) ToDtoJSON(holdInUnit *RotationalSpeedUnits) ([]byte, error) {
	// Convert to RotationalSpeedDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RotationalSpeed to a specific unit value.
// The function uses the provided unit type (RotationalSpeedUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the RotationalSpeed in the default unit (RadianPerSecond),
// formatted to two decimal places.
func (a RotationalSpeed) String() string {
	return a.ToString(RotationalSpeedRadianPerSecond, 2)
}

// ToString formats the RotationalSpeed value as a string with the specified unit and fractional digits.
// It converts the RotationalSpeed to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RotationalSpeed value will be converted (e.g., RadianPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RotationalSpeed with the unit abbreviation.
func (a *RotationalSpeed) ToString(unit RotationalSpeedUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRotationalSpeedAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRotationalSpeedAbbreviation(unit))
}

// Equals checks if the given RotationalSpeed is equal to the current RotationalSpeed.
//
// Parameters:
//    other: The RotationalSpeed to compare against.
//
// Returns:
//    bool: Returns true if both RotationalSpeed are equal, false otherwise.
func (a *RotationalSpeed) Equals(other *RotationalSpeed) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RotationalSpeed with another RotationalSpeed.
// It returns -1 if the current RotationalSpeed is less than the other RotationalSpeed, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RotationalSpeed to compare against.
//
// Returns:
//    int: -1 if the current RotationalSpeed is less, 1 if greater, and 0 if equal.
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

// Add adds the given RotationalSpeed to the current RotationalSpeed and returns the result.
// The result is a new RotationalSpeed instance with the sum of the values.
//
// Parameters:
//    other: The RotationalSpeed to add to the current RotationalSpeed.
//
// Returns:
//    *RotationalSpeed: A new RotationalSpeed instance representing the sum of both RotationalSpeed.
func (a *RotationalSpeed) Add(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RotationalSpeed from the current RotationalSpeed and returns the result.
// The result is a new RotationalSpeed instance with the difference of the values.
//
// Parameters:
//    other: The RotationalSpeed to subtract from the current RotationalSpeed.
//
// Returns:
//    *RotationalSpeed: A new RotationalSpeed instance representing the difference of both RotationalSpeed.
func (a *RotationalSpeed) Subtract(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RotationalSpeed by the given RotationalSpeed and returns the result.
// The result is a new RotationalSpeed instance with the product of the values.
//
// Parameters:
//    other: The RotationalSpeed to multiply with the current RotationalSpeed.
//
// Returns:
//    *RotationalSpeed: A new RotationalSpeed instance representing the product of both RotationalSpeed.
func (a *RotationalSpeed) Multiply(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value * other.BaseValue()}
}

// Divide divides the current RotationalSpeed by the given RotationalSpeed and returns the result.
// The result is a new RotationalSpeed instance with the quotient of the values.
//
// Parameters:
//    other: The RotationalSpeed to divide the current RotationalSpeed by.
//
// Returns:
//    *RotationalSpeed: A new RotationalSpeed instance representing the quotient of both RotationalSpeed.
func (a *RotationalSpeed) Divide(other *RotationalSpeed) *RotationalSpeed {
	return &RotationalSpeed{value: a.value / other.BaseValue()}
}

// GetRotationalSpeedAbbreviation gets the unit abbreviation.
func GetRotationalSpeedAbbreviation(unit RotationalSpeedUnits) string {
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