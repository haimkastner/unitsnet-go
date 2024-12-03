// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AreaDensityUnits enumeration
type AreaDensityUnits string

const (
    
        // 
        AreaDensityKilogramPerSquareMeter AreaDensityUnits = "KilogramPerSquareMeter"
        // Also known as grammage for paper industry. In fiber industry used with abbreviation 'gsm'.
        AreaDensityGramPerSquareMeter AreaDensityUnits = "GramPerSquareMeter"
        // 
        AreaDensityMilligramPerSquareMeter AreaDensityUnits = "MilligramPerSquareMeter"
)

// AreaDensityDto represents an AreaDensity
type AreaDensityDto struct {
	Value float64
	Unit  AreaDensityUnits
}

// AreaDensityDtoFactory struct to group related functions
type AreaDensityDtoFactory struct{}

func (udf AreaDensityDtoFactory) FromJSON(data []byte) (*AreaDensityDto, error) {
	a := AreaDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AreaDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// AreaDensity struct
type AreaDensity struct {
	value       float64
    
    kilograms_per_square_meterLazy *float64 
    grams_per_square_meterLazy *float64 
    milligrams_per_square_meterLazy *float64 
}

// AreaDensityFactory struct to group related functions
type AreaDensityFactory struct{}

func (uf AreaDensityFactory) CreateAreaDensity(value float64, unit AreaDensityUnits) (*AreaDensity, error) {
	return newAreaDensity(value, unit)
}

func (uf AreaDensityFactory) FromDto(dto AreaDensityDto) (*AreaDensity, error) {
	return newAreaDensity(dto.Value, dto.Unit)
}

func (uf AreaDensityFactory) FromDtoJSON(data []byte) (*AreaDensity, error) {
	unitDto, err := AreaDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AreaDensityFactory{}.FromDto(*unitDto)
}


// FromKilogramPerSquareMeter creates a new AreaDensity instance from KilogramPerSquareMeter.
func (uf AreaDensityFactory) FromKilogramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityKilogramPerSquareMeter)
}

// FromGramPerSquareMeter creates a new AreaDensity instance from GramPerSquareMeter.
func (uf AreaDensityFactory) FromGramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityGramPerSquareMeter)
}

// FromMilligramPerSquareMeter creates a new AreaDensity instance from MilligramPerSquareMeter.
func (uf AreaDensityFactory) FromMilligramsPerSquareMeter(value float64) (*AreaDensity, error) {
	return newAreaDensity(value, AreaDensityMilligramPerSquareMeter)
}




// newAreaDensity creates a new AreaDensity.
func newAreaDensity(value float64, fromUnit AreaDensityUnits) (*AreaDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AreaDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AreaDensity in KilogramPerSquareMeter.
func (a *AreaDensity) BaseValue() float64 {
	return a.value
}


// KilogramPerSquareMeter returns the value in KilogramPerSquareMeter.
func (a *AreaDensity) KilogramsPerSquareMeter() float64 {
	if a.kilograms_per_square_meterLazy != nil {
		return *a.kilograms_per_square_meterLazy
	}
	kilograms_per_square_meter := a.convertFromBase(AreaDensityKilogramPerSquareMeter)
	a.kilograms_per_square_meterLazy = &kilograms_per_square_meter
	return kilograms_per_square_meter
}

// GramPerSquareMeter returns the value in GramPerSquareMeter.
func (a *AreaDensity) GramsPerSquareMeter() float64 {
	if a.grams_per_square_meterLazy != nil {
		return *a.grams_per_square_meterLazy
	}
	grams_per_square_meter := a.convertFromBase(AreaDensityGramPerSquareMeter)
	a.grams_per_square_meterLazy = &grams_per_square_meter
	return grams_per_square_meter
}

// MilligramPerSquareMeter returns the value in MilligramPerSquareMeter.
func (a *AreaDensity) MilligramsPerSquareMeter() float64 {
	if a.milligrams_per_square_meterLazy != nil {
		return *a.milligrams_per_square_meterLazy
	}
	milligrams_per_square_meter := a.convertFromBase(AreaDensityMilligramPerSquareMeter)
	a.milligrams_per_square_meterLazy = &milligrams_per_square_meter
	return milligrams_per_square_meter
}


// ToDto creates an AreaDensityDto representation.
func (a *AreaDensity) ToDto(holdInUnit *AreaDensityUnits) AreaDensityDto {
	if holdInUnit == nil {
		defaultUnit := AreaDensityKilogramPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return AreaDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an AreaDensityDto representation.
func (a *AreaDensity) ToDtoJSON(holdInUnit *AreaDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts AreaDensity to a specific unit value.
func (a *AreaDensity) Convert(toUnit AreaDensityUnits) float64 {
	switch toUnit { 
    case AreaDensityKilogramPerSquareMeter:
		return a.KilogramsPerSquareMeter()
    case AreaDensityGramPerSquareMeter:
		return a.GramsPerSquareMeter()
    case AreaDensityMilligramPerSquareMeter:
		return a.MilligramsPerSquareMeter()
	default:
		return 0
	}
}

func (a *AreaDensity) convertFromBase(toUnit AreaDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaDensityKilogramPerSquareMeter:
		return (value) 
	case AreaDensityGramPerSquareMeter:
		return (value * 1000) 
	case AreaDensityMilligramPerSquareMeter:
		return (value * 1000000) 
	default:
		return math.NaN()
	}
}

func (a *AreaDensity) convertToBase(value float64, fromUnit AreaDensityUnits) float64 {
	switch fromUnit { 
	case AreaDensityKilogramPerSquareMeter:
		return (value) 
	case AreaDensityGramPerSquareMeter:
		return (value / 1000) 
	case AreaDensityMilligramPerSquareMeter:
		return (value / 1000000) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a AreaDensity) String() string {
	return a.ToString(AreaDensityKilogramPerSquareMeter, 2)
}

// ToString formats the AreaDensity to string.
// fractionalDigits -1 for not mention
func (a *AreaDensity) ToString(unit AreaDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AreaDensity) getUnitAbbreviation(unit AreaDensityUnits) string {
	switch unit { 
	case AreaDensityKilogramPerSquareMeter:
		return "kg/m²" 
	case AreaDensityGramPerSquareMeter:
		return "g/m²" 
	case AreaDensityMilligramPerSquareMeter:
		return "mg/m²" 
	default:
		return ""
	}
}

// Check if the given AreaDensity are equals to the current AreaDensity
func (a *AreaDensity) Equals(other *AreaDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given AreaDensity are equals to the current AreaDensity
func (a *AreaDensity) CompareTo(other *AreaDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given AreaDensity to the current AreaDensity.
func (a *AreaDensity) Add(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value + other.BaseValue()}
}

// Subtract the given AreaDensity to the current AreaDensity.
func (a *AreaDensity) Subtract(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value - other.BaseValue()}
}

// Multiply the given AreaDensity to the current AreaDensity.
func (a *AreaDensity) Multiply(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value * other.BaseValue()}
}

// Divide the given AreaDensity to the current AreaDensity.
func (a *AreaDensity) Divide(other *AreaDensity) *AreaDensity {
	return &AreaDensity{value: a.value / other.BaseValue()}
}