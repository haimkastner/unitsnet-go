package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VolumeFlowPerAreaUnits enumeration
type VolumeFlowPerAreaUnits string

const (
    
        // 
        VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter VolumeFlowPerAreaUnits = "CubicMeterPerSecondPerSquareMeter"
        // 
        VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot VolumeFlowPerAreaUnits = "CubicFootPerMinutePerSquareFoot"
)

// VolumeFlowPerAreaDto represents an VolumeFlowPerArea
type VolumeFlowPerAreaDto struct {
	Value float64
	Unit  VolumeFlowPerAreaUnits
}

// VolumeFlowPerAreaDtoFactory struct to group related functions
type VolumeFlowPerAreaDtoFactory struct{}

func (udf VolumeFlowPerAreaDtoFactory) FromJSON(data []byte) (*VolumeFlowPerAreaDto, error) {
	a := VolumeFlowPerAreaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VolumeFlowPerAreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VolumeFlowPerArea struct
type VolumeFlowPerArea struct {
	value       float64
    
    cubic_meters_per_second_per_square_meterLazy *float64 
    cubic_feet_per_minute_per_square_footLazy *float64 
}

// VolumeFlowPerAreaFactory struct to group related functions
type VolumeFlowPerAreaFactory struct{}

func (uf VolumeFlowPerAreaFactory) CreateVolumeFlowPerArea(value float64, unit VolumeFlowPerAreaUnits) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(value, unit)
}

func (uf VolumeFlowPerAreaFactory) FromDto(dto VolumeFlowPerAreaDto) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(dto.Value, dto.Unit)
}

func (uf VolumeFlowPerAreaFactory) FromDtoJSON(data []byte) (*VolumeFlowPerArea, error) {
	unitDto, err := VolumeFlowPerAreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VolumeFlowPerAreaFactory{}.FromDto(*unitDto)
}


// FromCubicMeterPerSecondPerSquareMeter creates a new VolumeFlowPerArea instance from CubicMeterPerSecondPerSquareMeter.
func (uf VolumeFlowPerAreaFactory) FromCubicMetersPerSecondPerSquareMeter(value float64) (*VolumeFlowPerArea, error) {
	return newVolumeFlowPerArea(value, VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
}

// FromCubicFootPerMinutePerSquareFoot creates a new VolumeFlowPerArea instance from CubicFootPerMinutePerSquareFoot.
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

// BaseValue returns the base value of VolumeFlowPerArea in CubicMeterPerSecondPerSquareMeter.
func (a *VolumeFlowPerArea) BaseValue() float64 {
	return a.value
}


// CubicMeterPerSecondPerSquareMeter returns the value in CubicMeterPerSecondPerSquareMeter.
func (a *VolumeFlowPerArea) CubicMetersPerSecondPerSquareMeter() float64 {
	if a.cubic_meters_per_second_per_square_meterLazy != nil {
		return *a.cubic_meters_per_second_per_square_meterLazy
	}
	cubic_meters_per_second_per_square_meter := a.convertFromBase(VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter)
	a.cubic_meters_per_second_per_square_meterLazy = &cubic_meters_per_second_per_square_meter
	return cubic_meters_per_second_per_square_meter
}

// CubicFootPerMinutePerSquareFoot returns the value in CubicFootPerMinutePerSquareFoot.
func (a *VolumeFlowPerArea) CubicFeetPerMinutePerSquareFoot() float64 {
	if a.cubic_feet_per_minute_per_square_footLazy != nil {
		return *a.cubic_feet_per_minute_per_square_footLazy
	}
	cubic_feet_per_minute_per_square_foot := a.convertFromBase(VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot)
	a.cubic_feet_per_minute_per_square_footLazy = &cubic_feet_per_minute_per_square_foot
	return cubic_feet_per_minute_per_square_foot
}


// ToDto creates an VolumeFlowPerAreaDto representation.
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

// ToDtoJSON creates an VolumeFlowPerAreaDto representation.
func (a *VolumeFlowPerArea) ToDtoJSON(holdInUnit *VolumeFlowPerAreaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VolumeFlowPerArea to a specific unit value.
func (a *VolumeFlowPerArea) Convert(toUnit VolumeFlowPerAreaUnits) float64 {
	switch toUnit { 
    case VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter:
		return a.CubicMetersPerSecondPerSquareMeter()
    case VolumeFlowPerAreaCubicFootPerMinutePerSquareFoot:
		return a.CubicFeetPerMinutePerSquareFoot()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a VolumeFlowPerArea) String() string {
	return a.ToString(VolumeFlowPerAreaCubicMeterPerSecondPerSquareMeter, 2)
}

// ToString formats the VolumeFlowPerArea to string.
// fractionalDigits -1 for not mention
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

// Check if the given VolumeFlowPerArea are equals to the current VolumeFlowPerArea
func (a *VolumeFlowPerArea) Equals(other *VolumeFlowPerArea) bool {
	return a.value == other.BaseValue()
}

// Check if the given VolumeFlowPerArea are equals to the current VolumeFlowPerArea
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

// Add the given VolumeFlowPerArea to the current VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Add(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value + other.BaseValue()}
}

// Subtract the given VolumeFlowPerArea to the current VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Subtract(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value - other.BaseValue()}
}

// Multiply the given VolumeFlowPerArea to the current VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Multiply(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value * other.BaseValue()}
}

// Divide the given VolumeFlowPerArea to the current VolumeFlowPerArea.
func (a *VolumeFlowPerArea) Divide(other *VolumeFlowPerArea) *VolumeFlowPerArea {
	return &VolumeFlowPerArea{value: a.value / other.BaseValue()}
}