// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RatioUnits enumeration
type RatioUnits string

const (
    
        // 
        RatioDecimalFraction RatioUnits = "DecimalFraction"
        // 
        RatioPercent RatioUnits = "Percent"
        // 
        RatioPartPerThousand RatioUnits = "PartPerThousand"
        // 
        RatioPartPerMillion RatioUnits = "PartPerMillion"
        // 
        RatioPartPerBillion RatioUnits = "PartPerBillion"
        // 
        RatioPartPerTrillion RatioUnits = "PartPerTrillion"
)

// RatioDto represents an Ratio
type RatioDto struct {
	Value float64
	Unit  RatioUnits
}

// RatioDtoFactory struct to group related functions
type RatioDtoFactory struct{}

func (udf RatioDtoFactory) FromJSON(data []byte) (*RatioDto, error) {
	a := RatioDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Ratio struct
type Ratio struct {
	value       float64
    
    decimal_fractionsLazy *float64 
    percentLazy *float64 
    parts_per_thousandLazy *float64 
    parts_per_millionLazy *float64 
    parts_per_billionLazy *float64 
    parts_per_trillionLazy *float64 
}

// RatioFactory struct to group related functions
type RatioFactory struct{}

func (uf RatioFactory) CreateRatio(value float64, unit RatioUnits) (*Ratio, error) {
	return newRatio(value, unit)
}

func (uf RatioFactory) FromDto(dto RatioDto) (*Ratio, error) {
	return newRatio(dto.Value, dto.Unit)
}

func (uf RatioFactory) FromDtoJSON(data []byte) (*Ratio, error) {
	unitDto, err := RatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RatioFactory{}.FromDto(*unitDto)
}


// FromDecimalFraction creates a new Ratio instance from DecimalFraction.
func (uf RatioFactory) FromDecimalFractions(value float64) (*Ratio, error) {
	return newRatio(value, RatioDecimalFraction)
}

// FromPercent creates a new Ratio instance from Percent.
func (uf RatioFactory) FromPercent(value float64) (*Ratio, error) {
	return newRatio(value, RatioPercent)
}

// FromPartPerThousand creates a new Ratio instance from PartPerThousand.
func (uf RatioFactory) FromPartsPerThousand(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerThousand)
}

// FromPartPerMillion creates a new Ratio instance from PartPerMillion.
func (uf RatioFactory) FromPartsPerMillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerMillion)
}

// FromPartPerBillion creates a new Ratio instance from PartPerBillion.
func (uf RatioFactory) FromPartsPerBillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerBillion)
}

// FromPartPerTrillion creates a new Ratio instance from PartPerTrillion.
func (uf RatioFactory) FromPartsPerTrillion(value float64) (*Ratio, error) {
	return newRatio(value, RatioPartPerTrillion)
}




// newRatio creates a new Ratio.
func newRatio(value float64, fromUnit RatioUnits) (*Ratio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Ratio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Ratio in DecimalFraction.
func (a *Ratio) BaseValue() float64 {
	return a.value
}


// DecimalFraction returns the value in DecimalFraction.
func (a *Ratio) DecimalFractions() float64 {
	if a.decimal_fractionsLazy != nil {
		return *a.decimal_fractionsLazy
	}
	decimal_fractions := a.convertFromBase(RatioDecimalFraction)
	a.decimal_fractionsLazy = &decimal_fractions
	return decimal_fractions
}

// Percent returns the value in Percent.
func (a *Ratio) Percent() float64 {
	if a.percentLazy != nil {
		return *a.percentLazy
	}
	percent := a.convertFromBase(RatioPercent)
	a.percentLazy = &percent
	return percent
}

// PartPerThousand returns the value in PartPerThousand.
func (a *Ratio) PartsPerThousand() float64 {
	if a.parts_per_thousandLazy != nil {
		return *a.parts_per_thousandLazy
	}
	parts_per_thousand := a.convertFromBase(RatioPartPerThousand)
	a.parts_per_thousandLazy = &parts_per_thousand
	return parts_per_thousand
}

// PartPerMillion returns the value in PartPerMillion.
func (a *Ratio) PartsPerMillion() float64 {
	if a.parts_per_millionLazy != nil {
		return *a.parts_per_millionLazy
	}
	parts_per_million := a.convertFromBase(RatioPartPerMillion)
	a.parts_per_millionLazy = &parts_per_million
	return parts_per_million
}

// PartPerBillion returns the value in PartPerBillion.
func (a *Ratio) PartsPerBillion() float64 {
	if a.parts_per_billionLazy != nil {
		return *a.parts_per_billionLazy
	}
	parts_per_billion := a.convertFromBase(RatioPartPerBillion)
	a.parts_per_billionLazy = &parts_per_billion
	return parts_per_billion
}

// PartPerTrillion returns the value in PartPerTrillion.
func (a *Ratio) PartsPerTrillion() float64 {
	if a.parts_per_trillionLazy != nil {
		return *a.parts_per_trillionLazy
	}
	parts_per_trillion := a.convertFromBase(RatioPartPerTrillion)
	a.parts_per_trillionLazy = &parts_per_trillion
	return parts_per_trillion
}


// ToDto creates an RatioDto representation.
func (a *Ratio) ToDto(holdInUnit *RatioUnits) RatioDto {
	if holdInUnit == nil {
		defaultUnit := RatioDecimalFraction // Default value
		holdInUnit = &defaultUnit
	}

	return RatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RatioDto representation.
func (a *Ratio) ToDtoJSON(holdInUnit *RatioUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Ratio to a specific unit value.
func (a *Ratio) Convert(toUnit RatioUnits) float64 {
	switch toUnit { 
    case RatioDecimalFraction:
		return a.DecimalFractions()
    case RatioPercent:
		return a.Percent()
    case RatioPartPerThousand:
		return a.PartsPerThousand()
    case RatioPartPerMillion:
		return a.PartsPerMillion()
    case RatioPartPerBillion:
		return a.PartsPerBillion()
    case RatioPartPerTrillion:
		return a.PartsPerTrillion()
	default:
		return 0
	}
}

func (a *Ratio) convertFromBase(toUnit RatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case RatioDecimalFraction:
		return (value) 
	case RatioPercent:
		return (value * 1e2) 
	case RatioPartPerThousand:
		return (value * 1e3) 
	case RatioPartPerMillion:
		return (value * 1e6) 
	case RatioPartPerBillion:
		return (value * 1e9) 
	case RatioPartPerTrillion:
		return (value * 1e12) 
	default:
		return math.NaN()
	}
}

func (a *Ratio) convertToBase(value float64, fromUnit RatioUnits) float64 {
	switch fromUnit { 
	case RatioDecimalFraction:
		return (value) 
	case RatioPercent:
		return (value / 1e2) 
	case RatioPartPerThousand:
		return (value / 1e3) 
	case RatioPartPerMillion:
		return (value / 1e6) 
	case RatioPartPerBillion:
		return (value / 1e9) 
	case RatioPartPerTrillion:
		return (value / 1e12) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Ratio) String() string {
	return a.ToString(RatioDecimalFraction, 2)
}

// ToString formats the Ratio to string.
// fractionalDigits -1 for not mention
func (a *Ratio) ToString(unit RatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Ratio) getUnitAbbreviation(unit RatioUnits) string {
	switch unit { 
	case RatioDecimalFraction:
		return "" 
	case RatioPercent:
		return "%" 
	case RatioPartPerThousand:
		return "‰" 
	case RatioPartPerMillion:
		return "ppm" 
	case RatioPartPerBillion:
		return "ppb" 
	case RatioPartPerTrillion:
		return "ppt" 
	default:
		return ""
	}
}

// Check if the given Ratio are equals to the current Ratio
func (a *Ratio) Equals(other *Ratio) bool {
	return a.value == other.BaseValue()
}

// Check if the given Ratio are equals to the current Ratio
func (a *Ratio) CompareTo(other *Ratio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Ratio to the current Ratio.
func (a *Ratio) Add(other *Ratio) *Ratio {
	return &Ratio{value: a.value + other.BaseValue()}
}

// Subtract the given Ratio to the current Ratio.
func (a *Ratio) Subtract(other *Ratio) *Ratio {
	return &Ratio{value: a.value - other.BaseValue()}
}

// Multiply the given Ratio to the current Ratio.
func (a *Ratio) Multiply(other *Ratio) *Ratio {
	return &Ratio{value: a.value * other.BaseValue()}
}

// Divide the given Ratio to the current Ratio.
func (a *Ratio) Divide(other *Ratio) *Ratio {
	return &Ratio{value: a.value / other.BaseValue()}
}