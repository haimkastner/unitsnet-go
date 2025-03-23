// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificVolumeUnits defines various units of SpecificVolume.
type SpecificVolumeUnits string

const (
    
        // 
        SpecificVolumeCubicMeterPerKilogram SpecificVolumeUnits = "CubicMeterPerKilogram"
        // 
        SpecificVolumeCubicFootPerPound SpecificVolumeUnits = "CubicFootPerPound"
        // 
        SpecificVolumeMillicubicMeterPerKilogram SpecificVolumeUnits = "MillicubicMeterPerKilogram"
)

var internalSpecificVolumeUnitsMap = map[SpecificVolumeUnits]bool{
	
	SpecificVolumeCubicMeterPerKilogram: true,
	SpecificVolumeCubicFootPerPound: true,
	SpecificVolumeMillicubicMeterPerKilogram: true,
}

// SpecificVolumeDto represents a SpecificVolume measurement with a numerical value and its corresponding unit.
type SpecificVolumeDto struct {
    // Value is the numerical representation of the SpecificVolume.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the SpecificVolume, as defined in the SpecificVolumeUnits enumeration.
	Unit  SpecificVolumeUnits `json:"unit" validate:"required,oneof=CubicMeterPerKilogram CubicFootPerPound MillicubicMeterPerKilogram"`
}

// SpecificVolumeDtoFactory groups methods for creating and serializing SpecificVolumeDto objects.
type SpecificVolumeDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpecificVolumeDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpecificVolumeDtoFactory) FromJSON(data []byte) (*SpecificVolumeDto, error) {
	a := SpecificVolumeDto{}

    // Parse JSON into SpecificVolumeDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a SpecificVolumeDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpecificVolumeDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SpecificVolume represents a measurement in a of SpecificVolume.
//
// In thermodynamics, the specific volume of a substance is the ratio of the substance's volume to its mass. It is the reciprocal of density and an intrinsic property of matter as well.
type SpecificVolume struct {
	// value is the base measurement stored internally.
	value       float64
    
    cubic_meters_per_kilogramLazy *float64 
    cubic_feet_per_poundLazy *float64 
    millicubic_meters_per_kilogramLazy *float64 
}

// SpecificVolumeFactory groups methods for creating SpecificVolume instances.
type SpecificVolumeFactory struct{}

// CreateSpecificVolume creates a new SpecificVolume instance from the given value and unit.
func (uf SpecificVolumeFactory) CreateSpecificVolume(value float64, unit SpecificVolumeUnits) (*SpecificVolume, error) {
	return newSpecificVolume(value, unit)
}

// FromDto converts a SpecificVolumeDto to a SpecificVolume instance.
func (uf SpecificVolumeFactory) FromDto(dto SpecificVolumeDto) (*SpecificVolume, error) {
	return newSpecificVolume(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SpecificVolume instance.
func (uf SpecificVolumeFactory) FromDtoJSON(data []byte) (*SpecificVolume, error) {
	unitDto, err := SpecificVolumeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpecificVolumeDto from JSON: %w", err)
	}
	return SpecificVolumeFactory{}.FromDto(*unitDto)
}


// FromCubicMetersPerKilogram creates a new SpecificVolume instance from a value in CubicMetersPerKilogram.
func (uf SpecificVolumeFactory) FromCubicMetersPerKilogram(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeCubicMeterPerKilogram)
}

// FromCubicFeetPerPound creates a new SpecificVolume instance from a value in CubicFeetPerPound.
func (uf SpecificVolumeFactory) FromCubicFeetPerPound(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeCubicFootPerPound)
}

// FromMillicubicMetersPerKilogram creates a new SpecificVolume instance from a value in MillicubicMetersPerKilogram.
func (uf SpecificVolumeFactory) FromMillicubicMetersPerKilogram(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeMillicubicMeterPerKilogram)
}


// newSpecificVolume creates a new SpecificVolume.
func newSpecificVolume(value float64, fromUnit SpecificVolumeUnits) (*SpecificVolume, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalSpecificVolumeUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in SpecificVolumeUnits", fromUnit)
	}
	a := &SpecificVolume{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificVolume in CubicMeterPerKilogram unit (the base/default quantity).
func (a *SpecificVolume) BaseValue() float64 {
	return a.value
}


// CubicMetersPerKilogram returns the SpecificVolume value in CubicMetersPerKilogram.
//
// 
func (a *SpecificVolume) CubicMetersPerKilogram() float64 {
	if a.cubic_meters_per_kilogramLazy != nil {
		return *a.cubic_meters_per_kilogramLazy
	}
	cubic_meters_per_kilogram := a.convertFromBase(SpecificVolumeCubicMeterPerKilogram)
	a.cubic_meters_per_kilogramLazy = &cubic_meters_per_kilogram
	return cubic_meters_per_kilogram
}

// CubicFeetPerPound returns the SpecificVolume value in CubicFeetPerPound.
//
// 
func (a *SpecificVolume) CubicFeetPerPound() float64 {
	if a.cubic_feet_per_poundLazy != nil {
		return *a.cubic_feet_per_poundLazy
	}
	cubic_feet_per_pound := a.convertFromBase(SpecificVolumeCubicFootPerPound)
	a.cubic_feet_per_poundLazy = &cubic_feet_per_pound
	return cubic_feet_per_pound
}

// MillicubicMetersPerKilogram returns the SpecificVolume value in MillicubicMetersPerKilogram.
//
// 
func (a *SpecificVolume) MillicubicMetersPerKilogram() float64 {
	if a.millicubic_meters_per_kilogramLazy != nil {
		return *a.millicubic_meters_per_kilogramLazy
	}
	millicubic_meters_per_kilogram := a.convertFromBase(SpecificVolumeMillicubicMeterPerKilogram)
	a.millicubic_meters_per_kilogramLazy = &millicubic_meters_per_kilogram
	return millicubic_meters_per_kilogram
}


// ToDto creates a SpecificVolumeDto representation from the SpecificVolume instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerKilogram by default.
func (a *SpecificVolume) ToDto(holdInUnit *SpecificVolumeUnits) SpecificVolumeDto {
	if holdInUnit == nil {
		defaultUnit := SpecificVolumeCubicMeterPerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificVolumeDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the SpecificVolume instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerKilogram by default.
func (a *SpecificVolume) ToDtoJSON(holdInUnit *SpecificVolumeUnits) ([]byte, error) {
	// Convert to SpecificVolumeDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SpecificVolume to a specific unit value.
// The function uses the provided unit type (SpecificVolumeUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *SpecificVolume) Convert(toUnit SpecificVolumeUnits) float64 {
	switch toUnit { 
    case SpecificVolumeCubicMeterPerKilogram:
		return a.CubicMetersPerKilogram()
    case SpecificVolumeCubicFootPerPound:
		return a.CubicFeetPerPound()
    case SpecificVolumeMillicubicMeterPerKilogram:
		return a.MillicubicMetersPerKilogram()
	default:
		return math.NaN()
	}
}

func (a *SpecificVolume) convertFromBase(toUnit SpecificVolumeUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificVolumeCubicMeterPerKilogram:
		return (value) 
	case SpecificVolumeCubicFootPerPound:
		return (value * 16.01846353) 
	case SpecificVolumeMillicubicMeterPerKilogram:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *SpecificVolume) convertToBase(value float64, fromUnit SpecificVolumeUnits) float64 {
	switch fromUnit { 
	case SpecificVolumeCubicMeterPerKilogram:
		return (value) 
	case SpecificVolumeCubicFootPerPound:
		return (value / 16.01846353) 
	case SpecificVolumeMillicubicMeterPerKilogram:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the SpecificVolume in the default unit (CubicMeterPerKilogram),
// formatted to two decimal places.
func (a SpecificVolume) String() string {
	return a.ToString(SpecificVolumeCubicMeterPerKilogram, 2)
}

// ToString formats the SpecificVolume value as a string with the specified unit and fractional digits.
// It converts the SpecificVolume to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SpecificVolume value will be converted (e.g., CubicMeterPerKilogram))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SpecificVolume with the unit abbreviation.
func (a *SpecificVolume) ToString(unit SpecificVolumeUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSpecificVolumeAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSpecificVolumeAbbreviation(unit))
}

// Equals checks if the given SpecificVolume is equal to the current SpecificVolume.
//
// Parameters:
//    other: The SpecificVolume to compare against.
//
// Returns:
//    bool: Returns true if both SpecificVolume are equal, false otherwise.
func (a *SpecificVolume) Equals(other *SpecificVolume) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SpecificVolume with another SpecificVolume.
// It returns -1 if the current SpecificVolume is less than the other SpecificVolume, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SpecificVolume to compare against.
//
// Returns:
//    int: -1 if the current SpecificVolume is less, 1 if greater, and 0 if equal.
func (a *SpecificVolume) CompareTo(other *SpecificVolume) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given SpecificVolume to the current SpecificVolume and returns the result.
// The result is a new SpecificVolume instance with the sum of the values.
//
// Parameters:
//    other: The SpecificVolume to add to the current SpecificVolume.
//
// Returns:
//    *SpecificVolume: A new SpecificVolume instance representing the sum of both SpecificVolume.
func (a *SpecificVolume) Add(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SpecificVolume from the current SpecificVolume and returns the result.
// The result is a new SpecificVolume instance with the difference of the values.
//
// Parameters:
//    other: The SpecificVolume to subtract from the current SpecificVolume.
//
// Returns:
//    *SpecificVolume: A new SpecificVolume instance representing the difference of both SpecificVolume.
func (a *SpecificVolume) Subtract(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SpecificVolume by the given SpecificVolume and returns the result.
// The result is a new SpecificVolume instance with the product of the values.
//
// Parameters:
//    other: The SpecificVolume to multiply with the current SpecificVolume.
//
// Returns:
//    *SpecificVolume: A new SpecificVolume instance representing the product of both SpecificVolume.
func (a *SpecificVolume) Multiply(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value * other.BaseValue()}
}

// Divide divides the current SpecificVolume by the given SpecificVolume and returns the result.
// The result is a new SpecificVolume instance with the quotient of the values.
//
// Parameters:
//    other: The SpecificVolume to divide the current SpecificVolume by.
//
// Returns:
//    *SpecificVolume: A new SpecificVolume instance representing the quotient of both SpecificVolume.
func (a *SpecificVolume) Divide(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value / other.BaseValue()}
}

// GetSpecificVolumeAbbreviation gets the unit abbreviation.
func GetSpecificVolumeAbbreviation(unit SpecificVolumeUnits) string {
	switch unit { 
	case SpecificVolumeCubicMeterPerKilogram:
		return "m³/kg" 
	case SpecificVolumeCubicFootPerPound:
		return "ft³/lb" 
	case SpecificVolumeMillicubicMeterPerKilogram:
		return "mm³/kg" 
	default:
		return ""
	}
}