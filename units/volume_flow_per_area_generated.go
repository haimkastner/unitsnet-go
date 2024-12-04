// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeFlowPerAreaUnits defines various units of VolumeFlowPerArea.
type VolumeFlowPerAreaUnits string

const (
    
        // 
        VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter VolumeFlowPerAreaUnits = "CubicMeterPerSecondPerSquareMeter"
        // 
        VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot VolumeFlowPerAreaUnits = "CubicFootPerMinutePerSquareFoot"
)

// VolumeFlowPerAreaDto represents a VolumeFlowPerArea measurement with a numerical value and its corresponding unit.
type VolumeFlowPerAreaDto struct {
    // Value is the numerical representation of the VolumeFlowPerArea.
	Value float64
    // Unit specifies the unit of measurement for the VolumeFlowPerArea, as defined in the VolumeFlowPerAreaUnits enumeration.
	Unit  VolumeFlowPerAreaUnits
}

// VolumeFlowPerAreaDtoFactory groups methods for creating and serializing VolumeFlowPerAreaDto objects.
type VolumeFlowPerAreaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a VolumeFlowPerAreaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf VolumeFlowPerAreaDtoFactory) FromJSON(data []byte) (*VolumeFlowPerAreaDto, error) {
	a := VolumeFlowPerAreaDto{}

    // Parse JSON into VolumeFlowPerAreaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a VolumeFlowPerAreaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a VolumeFlowPerAreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// VolumeFlowPerArea represents a measurement in a of VolumeFlowPerArea.
//
// None
type VolumeFlowPerArea struct {
	// value is the base measurement stored internally.
	value       float64
    
    cubic_meters_per_second_per_square_meterLazy *float64 
    cubic_feet_per_minute_per_square_footLazy *float64 
}

// VolumeFlowPerAreaFactory groups methods for creating VolumeFlowPerArea instances.
type VolumeFlowPerAreaFactory struct{}

// CreateVolumeFlowPerArea creates a new VolumeFlowPerArea instance from the given value and unit.
func (uf VolumeFlowPerAreaFactory) CreateVolumeFlowPerArea(value float64, unit VolumeFlowPerAreaUnits) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(value, unit)
}

// FromDto converts a VolumeFlowPerAreaDto to a VolumeFlowPerArea instance.
func (uf VolumeFlowPerAreaFactory) FromDto(dto VolumeFlowPerAreaDto) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a VolumeFlowPerArea instance.
func (uf VolumeFlowPerAreaFactory) FromDtoJSON(data []byte) (*VolumeFlowPerArea, error) {
	unitDto, err := VolumeFlowPerAreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse VolumeFlowPerAreaDto from JSON: %w", err)
	}
	return VolumeFlowPerAreaFactory{}.FromDto(*unitDto)
}


// FromCubicMetersPerSecondPerSquareMeter creates a new VolumeFlowPerArea instance from a value in CubicMetersPerSecondPerSquareMeter.
func (uf VolumeFlowPerAreaFactory) FromCubicMetersPerSecondPerSquareMeter(value float64) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(value, VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
}

// FromCubicFeetPerMinutePerSquareFoot creates a new VolumeFlowPerArea instance from a value in CubicFeetPerMinutePerSquareFoot.
func (uf VolumeFlowPerAreaFactory) FromCubicFeetPerMinutePerSquareFoot(value float64) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(value, VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
}


// newVolumeFlowPerArea creates a new VolumeFlowPerArea.
func newVolumeFlowPerArea(value float64, fromUnit VolumeFlowPerAreaUnits) (*VolumeFlowPerArea, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VolumeFlowPerArea{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VolumeFlowPerArea in CubicMeterPerSecondPerSquareMeter unit (the base/default quantity).
func (a *VolumeFlowPerArea) BaseValue() float64 {
	return a.value
}


// CubicMetersPerSecondPerSquareMeter returns the VolumeFlowPerArea value in CubicMetersPerSecondPerSquareMeter.
//
// 
func (a *VolumeFlowPerArea) CubicMetersPerSecondPerSquareMeter() float64 {
	if a.cubic_meters_per_second_per_square_meterLazy != nil {
		return *a.cubic_meters_per_second_per_square_meterLazy
	}
	cubic_meters_per_second_per_square_meter := a.convertFromBase(VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	a.cubic_meters_per_second_per_square_meterLazy = &cubic_meters_per_second_per_square_meter
	return cubic_meters_per_second_per_square_meter
}

// CubicFeetPerMinutePerSquareFoot returns the VolumeFlowPerArea value in CubicFeetPerMinutePerSquareFoot.
//
// 
func (a *VolumeFlowPerArea) CubicFeetPerMinutePerSquareFoot() float64 {
	if a.cubic_feet_per_minute_per_square_footLazy != nil {
		return *a.cubic_feet_per_minute_per_square_footLazy
	}
	cubic_feet_per_minute_per_square_foot := a.convertFromBase(VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
	a.cubic_feet_per_minute_per_square_footLazy = &cubic_feet_per_minute_per_square_foot
	return cubic_feet_per_minute_per_square_foot
}


// ToDto creates a VolumeFlowPerAreaDto representation from the VolumeFlowPerArea instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerSecondPerSquareMeter by default.
func (a *VolumeFlowPerArea) ToDto(holdInUnit *VolumeFlowPerAreaUnits) VolumeFlowPerAreaDto {
	if holdInUnit == nil {
		defaultUnit := VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return VolumeFlowPerAreaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the VolumeFlowPerArea instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CubicMeterPerSecondPerSquareMeter by default.
func (a *VolumeFlowPerArea) ToDtoJSON(holdInUnit *VolumeFlowPerAreaUnits) ([]byte, error) {
	// Convert to VolumeFlowPerAreaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a VolumeFlowPerArea to a specific unit value.
// The function uses the provided unit type (VolumeFlowPerAreaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *VolumeFlowPerArea) Convert(toUnit VolumeFlowPerAreaUnits) float64 {
	switch toUnit { 
    case VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter:
		return a.CubicMetersPerSecondPerSquareMeter()
    case VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot:
		return a.CubicFeetPerMinutePerSquareFoot()
	default:
		return math.NaN()
	}
}

func (a *VolumeFlowPerArea) convertFromBase(toUnit VolumeFlowPerAreaUnits) float64 {
    value := a.value
	switch toUnit { 
	case VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter:
		return (value) 
	case VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot:
		return (value * 196.850394) 
	default:
		return math.NaN()
	}
}

func (a *VolumeFlowPerArea) convertToBase(value float64, fromUnit VolumeFlowPerAreaUnits) float64 {
	switch fromUnit { 
	case VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter:
		return (value) 
	case VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot:
		return (value / 196.850394) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the VolumeFlowPerArea in the default unit (CubicMeterPerSecondPerSquareMeter),
// formatted to two decimal places.
func (a VolumeFlowPerArea) String() string {
	return a.ToString(VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, 2)
}

// ToString formats the VolumeFlowPerArea value as a string with the specified unit and fractional digits.
// It converts the VolumeFlowPerArea to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the VolumeFlowPerArea value will be converted (e.g., CubicMeterPerSecondPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the VolumeFlowPerArea with the unit abbreviation.
func (a *VolumeFlowPerArea) ToString(unit VolumeFlowPerAreaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VolumeFlowPerArea) getUnitAbbreviation(unit VolumeFlowPerAreaUnits) string {
	switch unit { 
	case VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter:
		return "m³/(s·m²)" 
	case VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot:
		return "CFM/ft²" 
	default:
		return ""
	}
}

// Equals checks if the given VolumeFlowPerArea is equal to the current VolumeFlowPerArea.
//
// Parameters:
//    other: The VolumeFlowPerArea to compare against.
//
// Returns:
//    bool: Returns true if both VolumeFlowPerArea are equal, false otherwise.
func (a *VolumeFlowPerArea) Equals(other *VolumeFlowPerArea) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current VolumeFlowPerArea with another VolumeFlowPerArea.
// It returns -1 if the current VolumeFlowPerArea is less than the other VolumeFlowPerArea, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The VolumeFlowPerArea to compare against.
//
// Returns:
//    int: -1 if the current VolumeFlowPerArea is less, 1 if greater, and 0 if equal.
func (a *VolumeFlowPerArea) CompareTo(other *VolumeFlowPerArea) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given VolumeFlowPerArea to the current VolumeFlowPerArea and returns the result.
// The result is a new VolumeFlowPerArea instance with the sum of the values.
//
// Parameters:
//    other: The VolumeFlowPerArea to add to the current VolumeFlowPerArea.
//
// Returns:
//    *VolumeFlowPerArea: A new VolumeFlowPerArea instance representing the sum of both VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Add(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given VolumeFlowPerArea from the current VolumeFlowPerArea and returns the result.
// The result is a new VolumeFlowPerArea instance with the difference of the values.
//
// Parameters:
//    other: The VolumeFlowPerArea to subtract from the current VolumeFlowPerArea.
//
// Returns:
//    *VolumeFlowPerArea: A new VolumeFlowPerArea instance representing the difference of both VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Subtract(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current VolumeFlowPerArea by the given VolumeFlowPerArea and returns the result.
// The result is a new VolumeFlowPerArea instance with the product of the values.
//
// Parameters:
//    other: The VolumeFlowPerArea to multiply with the current VolumeFlowPerArea.
//
// Returns:
//    *VolumeFlowPerArea: A new VolumeFlowPerArea instance representing the product of both VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Multiply(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value * other.BaseValue()}
}

// Divide divides the current VolumeFlowPerArea by the given VolumeFlowPerArea and returns the result.
// The result is a new VolumeFlowPerArea instance with the quotient of the values.
//
// Parameters:
//    other: The VolumeFlowPerArea to divide the current VolumeFlowPerArea by.
//
// Returns:
//    *VolumeFlowPerArea: A new VolumeFlowPerArea instance representing the quotient of both VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Divide(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value / other.BaseValue()}
}