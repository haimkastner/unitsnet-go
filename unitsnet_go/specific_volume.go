package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificVolumeUnits enumeration
type SpecificVolumeUnits string

const (
    
        // 
        SpecificVolumeCubicMeterPerKilogram SpecificVolumeUnits = "CubicMeterPerKilogram"
        // 
        SpecificVolumeCubicFootPerPound SpecificVolumeUnits = "CubicFootPerPound"
        // 
        SpecificVolumeMillicubicMeterPerKilogram SpecificVolumeUnits = "MillicubicMeterPerKilogram"
)

// SpecificVolumeDto represents an SpecificVolume
type SpecificVolumeDto struct {
	Value float64
	Unit  SpecificVolumeUnits
}

// SpecificVolumeDtoFactory struct to group related functions
type SpecificVolumeDtoFactory struct{}

func (udf SpecificVolumeDtoFactory) FromJSON(data []byte) (*SpecificVolumeDto, error) {
	a := SpecificVolumeDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpecificVolumeDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SpecificVolume struct
type SpecificVolume struct {
	value       float64
    
    cubic_meters_per_kilogramLazy *float64 
    cubic_feet_per_poundLazy *float64 
    millicubic_meters_per_kilogramLazy *float64 
}

// SpecificVolumeFactory struct to group related functions
type SpecificVolumeFactory struct{}

func (uf SpecificVolumeFactory) CreateSpecificVolume(value float64, unit SpecificVolumeUnits) (*SpecificVolume, error) {
	return newSpecificVolume(value, unit)
}

func (uf SpecificVolumeFactory) FromDto(dto SpecificVolumeDto) (*SpecificVolume, error) {
	return newSpecificVolume(dto.Value, dto.Unit)
}

func (uf SpecificVolumeFactory) FromDtoJSON(data []byte) (*SpecificVolume, error) {
	unitDto, err := SpecificVolumeDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpecificVolumeFactory{}.FromDto(*unitDto)
}


// FromCubicMeterPerKilogram creates a new SpecificVolume instance from CubicMeterPerKilogram.
func (uf SpecificVolumeFactory) FromCubicMetersPerKilogram(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeCubicMeterPerKilogram)
}

// FromCubicFootPerPound creates a new SpecificVolume instance from CubicFootPerPound.
func (uf SpecificVolumeFactory) FromCubicFeetPerPound(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeCubicFootPerPound)
}

// FromMillicubicMeterPerKilogram creates a new SpecificVolume instance from MillicubicMeterPerKilogram.
func (uf SpecificVolumeFactory) FromMillicubicMetersPerKilogram(value float64) (*SpecificVolume, error) {
	return newSpecificVolume(value, SpecificVolumeMillicubicMeterPerKilogram)
}




// newSpecificVolume creates a new SpecificVolume.
func newSpecificVolume(value float64, fromUnit SpecificVolumeUnits) (*SpecificVolume, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificVolume{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificVolume in CubicMeterPerKilogram.
func (a *SpecificVolume) BaseValue() float64 {
	return a.value
}


// CubicMeterPerKilogram returns the value in CubicMeterPerKilogram.
func (a *SpecificVolume) CubicMetersPerKilogram() float64 {
	if a.cubic_meters_per_kilogramLazy != nil {
		return *a.cubic_meters_per_kilogramLazy
	}
	cubic_meters_per_kilogram := a.convertFromBase(SpecificVolumeCubicMeterPerKilogram)
	a.cubic_meters_per_kilogramLazy = &cubic_meters_per_kilogram
	return cubic_meters_per_kilogram
}

// CubicFootPerPound returns the value in CubicFootPerPound.
func (a *SpecificVolume) CubicFeetPerPound() float64 {
	if a.cubic_feet_per_poundLazy != nil {
		return *a.cubic_feet_per_poundLazy
	}
	cubic_feet_per_pound := a.convertFromBase(SpecificVolumeCubicFootPerPound)
	a.cubic_feet_per_poundLazy = &cubic_feet_per_pound
	return cubic_feet_per_pound
}

// MillicubicMeterPerKilogram returns the value in MillicubicMeterPerKilogram.
func (a *SpecificVolume) MillicubicMetersPerKilogram() float64 {
	if a.millicubic_meters_per_kilogramLazy != nil {
		return *a.millicubic_meters_per_kilogramLazy
	}
	millicubic_meters_per_kilogram := a.convertFromBase(SpecificVolumeMillicubicMeterPerKilogram)
	a.millicubic_meters_per_kilogramLazy = &millicubic_meters_per_kilogram
	return millicubic_meters_per_kilogram
}


// ToDto creates an SpecificVolumeDto representation.
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

// ToDtoJSON creates an SpecificVolumeDto representation.
func (a *SpecificVolume) ToDtoJSON(holdInUnit *SpecificVolumeUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SpecificVolume to a specific unit value.
func (a *SpecificVolume) Convert(toUnit SpecificVolumeUnits) float64 {
	switch toUnit { 
    case SpecificVolumeCubicMeterPerKilogram:
		return a.CubicMetersPerKilogram()
    case SpecificVolumeCubicFootPerPound:
		return a.CubicFeetPerPound()
    case SpecificVolumeMillicubicMeterPerKilogram:
		return a.MillicubicMetersPerKilogram()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a SpecificVolume) String() string {
	return a.ToString(SpecificVolumeCubicMeterPerKilogram, 2)
}

// ToString formats the SpecificVolume to string.
// fractionalDigits -1 for not mention
func (a *SpecificVolume) ToString(unit SpecificVolumeUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SpecificVolume) getUnitAbbreviation(unit SpecificVolumeUnits) string {
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

// Check if the given SpecificVolume are equals to the current SpecificVolume
func (a *SpecificVolume) Equals(other *SpecificVolume) bool {
	return a.value == other.BaseValue()
}

// Check if the given SpecificVolume are equals to the current SpecificVolume
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

// Add the given SpecificVolume to the current SpecificVolume.
func (a *SpecificVolume) Add(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value + other.BaseValue()}
}

// Subtract the given SpecificVolume to the current SpecificVolume.
func (a *SpecificVolume) Subtract(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value - other.BaseValue()}
}

// Multiply the given SpecificVolume to the current SpecificVolume.
func (a *SpecificVolume) Multiply(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value * other.BaseValue()}
}

// Divide the given SpecificVolume to the current SpecificVolume.
func (a *SpecificVolume) Divide(other *SpecificVolume) *SpecificVolume {
	return &SpecificVolume{value: a.value / other.BaseValue()}
}