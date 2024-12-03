// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SolidAngleUnits enumeration
type SolidAngleUnits string

const (
    
        // 
        SolidAngleSteradian SolidAngleUnits = "Steradian"
)

// SolidAngleDto represents an SolidAngle
type SolidAngleDto struct {
	Value float64
	Unit  SolidAngleUnits
}

// SolidAngleDtoFactory struct to group related functions
type SolidAngleDtoFactory struct{}

func (udf SolidAngleDtoFactory) FromJSON(data []byte) (*SolidAngleDto, error) {
	a := SolidAngleDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SolidAngleDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SolidAngle struct
type SolidAngle struct {
	value       float64
    
    steradiansLazy *float64 
}

// SolidAngleFactory struct to group related functions
type SolidAngleFactory struct{}

func (uf SolidAngleFactory) CreateSolidAngle(value float64, unit SolidAngleUnits) (*SolidAngle, error) {
	return newSolidAngle(value, unit)
}

func (uf SolidAngleFactory) FromDto(dto SolidAngleDto) (*SolidAngle, error) {
	return newSolidAngle(dto.Value, dto.Unit)
}

func (uf SolidAngleFactory) FromDtoJSON(data []byte) (*SolidAngle, error) {
	unitDto, err := SolidAngleDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SolidAngleFactory{}.FromDto(*unitDto)
}


// FromSteradian creates a new SolidAngle instance from Steradian.
func (uf SolidAngleFactory) FromSteradians(value float64) (*SolidAngle, error) {
	return newSolidAngle(value, SolidAngleSteradian)
}




// newSolidAngle creates a new SolidAngle.
func newSolidAngle(value float64, fromUnit SolidAngleUnits) (*SolidAngle, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SolidAngle{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SolidAngle in Steradian.
func (a *SolidAngle) BaseValue() float64 {
	return a.value
}


// Steradian returns the value in Steradian.
func (a *SolidAngle) Steradians() float64 {
	if a.steradiansLazy != nil {
		return *a.steradiansLazy
	}
	steradians := a.convertFromBase(SolidAngleSteradian)
	a.steradiansLazy = &steradians
	return steradians
}


// ToDto creates an SolidAngleDto representation.
func (a *SolidAngle) ToDto(holdInUnit *SolidAngleUnits) SolidAngleDto {
	if holdInUnit == nil {
		defaultUnit := SolidAngleSteradian // Default value
		holdInUnit = &defaultUnit
	}

	return SolidAngleDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an SolidAngleDto representation.
func (a *SolidAngle) ToDtoJSON(holdInUnit *SolidAngleUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SolidAngle to a specific unit value.
func (a *SolidAngle) Convert(toUnit SolidAngleUnits) float64 {
	switch toUnit { 
    case SolidAngleSteradian:
		return a.Steradians()
	default:
		return 0
	}
}

func (a *SolidAngle) convertFromBase(toUnit SolidAngleUnits) float64 {
    value := a.value
	switch toUnit { 
	case SolidAngleSteradian:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *SolidAngle) convertToBase(value float64, fromUnit SolidAngleUnits) float64 {
	switch fromUnit { 
	case SolidAngleSteradian:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a SolidAngle) String() string {
	return a.ToString(SolidAngleSteradian, 2)
}

// ToString formats the SolidAngle to string.
// fractionalDigits -1 for not mention
func (a *SolidAngle) ToString(unit SolidAngleUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SolidAngle) getUnitAbbreviation(unit SolidAngleUnits) string {
	switch unit { 
	case SolidAngleSteradian:
		return "sr" 
	default:
		return ""
	}
}

// Check if the given SolidAngle are equals to the current SolidAngle
func (a *SolidAngle) Equals(other *SolidAngle) bool {
	return a.value == other.BaseValue()
}

// Check if the given SolidAngle are equals to the current SolidAngle
func (a *SolidAngle) CompareTo(other *SolidAngle) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given SolidAngle to the current SolidAngle.
func (a *SolidAngle) Add(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value + other.BaseValue()}
}

// Subtract the given SolidAngle to the current SolidAngle.
func (a *SolidAngle) Subtract(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value - other.BaseValue()}
}

// Multiply the given SolidAngle to the current SolidAngle.
func (a *SolidAngle) Multiply(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value * other.BaseValue()}
}

// Divide the given SolidAngle to the current SolidAngle.
func (a *SolidAngle) Divide(other *SolidAngle) *SolidAngle {
	return &SolidAngle{value: a.value / other.BaseValue()}
}