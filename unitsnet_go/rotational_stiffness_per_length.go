package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalStiffnessPerLengthUnits enumeration
type RotationalStiffnessPerLengthUnits string

const (
    
        // 
        RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "NewtonMeterPerRadianPerMeter"
        // 
        RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot RotationalStiffnessPerLengthUnits = "PoundForceFootPerDegreesPerFoot"
        // 
        RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot RotationalStiffnessPerLengthUnits = "KilopoundForceFootPerDegreesPerFoot"
        // 
        RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "KilonewtonMeterPerRadianPerMeter"
        // 
        RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter RotationalStiffnessPerLengthUnits = "MeganewtonMeterPerRadianPerMeter"
)

// RotationalStiffnessPerLengthDto represents an RotationalStiffnessPerLength
type RotationalStiffnessPerLengthDto struct {
	Value float64
	Unit  RotationalStiffnessPerLengthUnits
}

// RotationalStiffnessPerLengthDtoFactory struct to group related functions
type RotationalStiffnessPerLengthDtoFactory struct{}

func (udf RotationalStiffnessPerLengthDtoFactory) FromJSON(data []byte) (*RotationalStiffnessPerLengthDto, error) {
	a := RotationalStiffnessPerLengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RotationalStiffnessPerLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RotationalStiffnessPerLength struct
type RotationalStiffnessPerLength struct {
	value       float64
    
    newton_meters_per_radian_per_meterLazy *float64 
    pound_force_feet_per_degrees_per_feetLazy *float64 
    kilopound_force_feet_per_degrees_per_feetLazy *float64 
    kilonewton_meters_per_radian_per_meterLazy *float64 
    meganewton_meters_per_radian_per_meterLazy *float64 
}

// RotationalStiffnessPerLengthFactory struct to group related functions
type RotationalStiffnessPerLengthFactory struct{}

func (uf RotationalStiffnessPerLengthFactory) CreateRotationalStiffnessPerLength(value float64, unit RotationalStiffnessPerLengthUnits) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, unit)
}

func (uf RotationalStiffnessPerLengthFactory) FromDto(dto RotationalStiffnessPerLengthDto) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(dto.Value, dto.Unit)
}

func (uf RotationalStiffnessPerLengthFactory) FromDtoJSON(data []byte) (*RotationalStiffnessPerLength, error) {
	unitDto, err := RotationalStiffnessPerLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RotationalStiffnessPerLengthFactory{}.FromDto(*unitDto)
}


// FromNewtonMeterPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from NewtonMeterPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromNewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
}

// FromPoundForceFootPerDegreesPerFoot creates a new RotationalStiffnessPerLength instance from PoundForceFootPerDegreesPerFoot.
func (uf RotationalStiffnessPerLengthFactory) FromPoundForceFeetPerDegreesPerFeet(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
}

// FromKilopoundForceFootPerDegreesPerFoot creates a new RotationalStiffnessPerLength instance from KilopoundForceFootPerDegreesPerFoot.
func (uf RotationalStiffnessPerLengthFactory) FromKilopoundForceFeetPerDegreesPerFeet(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
}

// FromKilonewtonMeterPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from KilonewtonMeterPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromKilonewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
}

// FromMeganewtonMeterPerRadianPerMeter creates a new RotationalStiffnessPerLength instance from MeganewtonMeterPerRadianPerMeter.
func (uf RotationalStiffnessPerLengthFactory) FromMeganewtonMetersPerRadianPerMeter(value float64) (*RotationalStiffnessPerLength, error) {
	return newRotationalStiffnessPerLength(value, RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
}




// newRotationalStiffnessPerLength creates a new RotationalStiffnessPerLength.
func newRotationalStiffnessPerLength(value float64, fromUnit RotationalStiffnessPerLengthUnits) (*RotationalStiffnessPerLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalStiffnessPerLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalStiffnessPerLength in NewtonMeterPerRadianPerMeter.
func (a *RotationalStiffnessPerLength) BaseValue() float64 {
	return a.value
}


// NewtonMeterPerRadianPerMeter returns the value in NewtonMeterPerRadianPerMeter.
func (a *RotationalStiffnessPerLength) NewtonMetersPerRadianPerMeter() float64 {
	if a.newton_meters_per_radian_per_meterLazy != nil {
		return *a.newton_meters_per_radian_per_meterLazy
	}
	newton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter)
	a.newton_meters_per_radian_per_meterLazy = &newton_meters_per_radian_per_meter
	return newton_meters_per_radian_per_meter
}

// PoundForceFootPerDegreesPerFoot returns the value in PoundForceFootPerDegreesPerFoot.
func (a *RotationalStiffnessPerLength) PoundForceFeetPerDegreesPerFeet() float64 {
	if a.pound_force_feet_per_degrees_per_feetLazy != nil {
		return *a.pound_force_feet_per_degrees_per_feetLazy
	}
	pound_force_feet_per_degrees_per_feet := a.convertFromBase(RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot)
	a.pound_force_feet_per_degrees_per_feetLazy = &pound_force_feet_per_degrees_per_feet
	return pound_force_feet_per_degrees_per_feet
}

// KilopoundForceFootPerDegreesPerFoot returns the value in KilopoundForceFootPerDegreesPerFoot.
func (a *RotationalStiffnessPerLength) KilopoundForceFeetPerDegreesPerFeet() float64 {
	if a.kilopound_force_feet_per_degrees_per_feetLazy != nil {
		return *a.kilopound_force_feet_per_degrees_per_feetLazy
	}
	kilopound_force_feet_per_degrees_per_feet := a.convertFromBase(RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot)
	a.kilopound_force_feet_per_degrees_per_feetLazy = &kilopound_force_feet_per_degrees_per_feet
	return kilopound_force_feet_per_degrees_per_feet
}

// KilonewtonMeterPerRadianPerMeter returns the value in KilonewtonMeterPerRadianPerMeter.
func (a *RotationalStiffnessPerLength) KilonewtonMetersPerRadianPerMeter() float64 {
	if a.kilonewton_meters_per_radian_per_meterLazy != nil {
		return *a.kilonewton_meters_per_radian_per_meterLazy
	}
	kilonewton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter)
	a.kilonewton_meters_per_radian_per_meterLazy = &kilonewton_meters_per_radian_per_meter
	return kilonewton_meters_per_radian_per_meter
}

// MeganewtonMeterPerRadianPerMeter returns the value in MeganewtonMeterPerRadianPerMeter.
func (a *RotationalStiffnessPerLength) MeganewtonMetersPerRadianPerMeter() float64 {
	if a.meganewton_meters_per_radian_per_meterLazy != nil {
		return *a.meganewton_meters_per_radian_per_meterLazy
	}
	meganewton_meters_per_radian_per_meter := a.convertFromBase(RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter)
	a.meganewton_meters_per_radian_per_meterLazy = &meganewton_meters_per_radian_per_meter
	return meganewton_meters_per_radian_per_meter
}


// ToDto creates an RotationalStiffnessPerLengthDto representation.
func (a *RotationalStiffnessPerLength) ToDto(holdInUnit *RotationalStiffnessPerLengthUnits) RotationalStiffnessPerLengthDto {
	if holdInUnit == nil {
		defaultUnit := RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalStiffnessPerLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RotationalStiffnessPerLengthDto representation.
func (a *RotationalStiffnessPerLength) ToDtoJSON(holdInUnit *RotationalStiffnessPerLengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RotationalStiffnessPerLength to a specific unit value.
func (a *RotationalStiffnessPerLength) Convert(toUnit RotationalStiffnessPerLengthUnits) float64 {
	switch toUnit { 
    case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return a.NewtonMetersPerRadianPerMeter()
    case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return a.PoundForceFeetPerDegreesPerFeet()
    case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return a.KilopoundForceFeetPerDegreesPerFeet()
    case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return a.KilonewtonMetersPerRadianPerMeter()
    case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return a.MeganewtonMetersPerRadianPerMeter()
	default:
		return 0
	}
}

func (a *RotationalStiffnessPerLength) convertFromBase(toUnit RotationalStiffnessPerLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return (value) 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return (value / 254.864324570) 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return (value / 254864.324570) 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return ((value) / 1000.0) 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *RotationalStiffnessPerLength) convertToBase(value float64, fromUnit RotationalStiffnessPerLengthUnits) float64 {
	switch fromUnit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return (value) 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return (value * 254.864324570) 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return (value * 254864.324570) 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return ((value) * 1000.0) 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RotationalStiffnessPerLength) String() string {
	return a.ToString(RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter, 2)
}

// ToString formats the RotationalStiffnessPerLength to string.
// fractionalDigits -1 for not mention
func (a *RotationalStiffnessPerLength) ToString(unit RotationalStiffnessPerLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalStiffnessPerLength) getUnitAbbreviation(unit RotationalStiffnessPerLengthUnits) string {
	switch unit { 
	case RotationalStiffnessPerLengthNewtonMeterPerRadianPerMeter:
		return "N·m/rad/m" 
	case RotationalStiffnessPerLengthPoundForceFootPerDegreesPerFoot:
		return "lbf·ft/deg/ft" 
	case RotationalStiffnessPerLengthKilopoundForceFootPerDegreesPerFoot:
		return "kipf·ft/°/ft" 
	case RotationalStiffnessPerLengthKilonewtonMeterPerRadianPerMeter:
		return "kN·m/rad/m" 
	case RotationalStiffnessPerLengthMeganewtonMeterPerRadianPerMeter:
		return "MN·m/rad/m" 
	default:
		return ""
	}
}

// Check if the given RotationalStiffnessPerLength are equals to the current RotationalStiffnessPerLength
func (a *RotationalStiffnessPerLength) Equals(other *RotationalStiffnessPerLength) bool {
	return a.value == other.BaseValue()
}

// Check if the given RotationalStiffnessPerLength are equals to the current RotationalStiffnessPerLength
func (a *RotationalStiffnessPerLength) CompareTo(other *RotationalStiffnessPerLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RotationalStiffnessPerLength to the current RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Add(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value + other.BaseValue()}
}

// Subtract the given RotationalStiffnessPerLength to the current RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Subtract(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value - other.BaseValue()}
}

// Multiply the given RotationalStiffnessPerLength to the current RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Multiply(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value * other.BaseValue()}
}

// Divide the given RotationalStiffnessPerLength to the current RotationalStiffnessPerLength.
func (a *RotationalStiffnessPerLength) Divide(other *RotationalStiffnessPerLength) *RotationalStiffnessPerLength {
	return &RotationalStiffnessPerLength{value: a.value / other.BaseValue()}
}