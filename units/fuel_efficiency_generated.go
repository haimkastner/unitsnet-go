// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// FuelEfficiencyUnits enumeration
type FuelEfficiencyUnits string

const (
    
        // 
        FuelEfficiencyLiterPer100Kilometers FuelEfficiencyUnits = "LiterPer100Kilometers"
        // 
        FuelEfficiencyMilePerUsGallon FuelEfficiencyUnits = "MilePerUsGallon"
        // 
        FuelEfficiencyMilePerUkGallon FuelEfficiencyUnits = "MilePerUkGallon"
        // 
        FuelEfficiencyKilometerPerLiter FuelEfficiencyUnits = "KilometerPerLiter"
)

// FuelEfficiencyDto represents an FuelEfficiency
type FuelEfficiencyDto struct {
	Value float64
	Unit  FuelEfficiencyUnits
}

// FuelEfficiencyDtoFactory struct to group related functions
type FuelEfficiencyDtoFactory struct{}

func (udf FuelEfficiencyDtoFactory) FromJSON(data []byte) (*FuelEfficiencyDto, error) {
	a := FuelEfficiencyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a FuelEfficiencyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// FuelEfficiency struct
type FuelEfficiency struct {
	value       float64
    
    liters_per100_kilometersLazy *float64 
    miles_per_us_gallonLazy *float64 
    miles_per_uk_gallonLazy *float64 
    kilometers_per_litersLazy *float64 
}

// FuelEfficiencyFactory struct to group related functions
type FuelEfficiencyFactory struct{}

func (uf FuelEfficiencyFactory) CreateFuelEfficiency(value float64, unit FuelEfficiencyUnits) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, unit)
}

func (uf FuelEfficiencyFactory) FromDto(dto FuelEfficiencyDto) (*FuelEfficiency, error) {
	return newFuelEfficiency(dto.Value, dto.Unit)
}

func (uf FuelEfficiencyFactory) FromDtoJSON(data []byte) (*FuelEfficiency, error) {
	unitDto, err := FuelEfficiencyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return FuelEfficiencyFactory{}.FromDto(*unitDto)
}


// FromLiterPer100Kilometers creates a new FuelEfficiency instance from LiterPer100Kilometers.
func (uf FuelEfficiencyFactory) FromLitersPer100Kilometers(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyLiterPer100Kilometers)
}

// FromMilePerUsGallon creates a new FuelEfficiency instance from MilePerUsGallon.
func (uf FuelEfficiencyFactory) FromMilesPerUsGallon(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyMilePerUsGallon)
}

// FromMilePerUkGallon creates a new FuelEfficiency instance from MilePerUkGallon.
func (uf FuelEfficiencyFactory) FromMilesPerUkGallon(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyMilePerUkGallon)
}

// FromKilometerPerLiter creates a new FuelEfficiency instance from KilometerPerLiter.
func (uf FuelEfficiencyFactory) FromKilometersPerLiters(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyKilometerPerLiter)
}




// newFuelEfficiency creates a new FuelEfficiency.
func newFuelEfficiency(value float64, fromUnit FuelEfficiencyUnits) (*FuelEfficiency, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &FuelEfficiency{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of FuelEfficiency in LiterPer100Kilometers.
func (a *FuelEfficiency) BaseValue() float64 {
	return a.value
}


// LiterPer100Kilometers returns the value in LiterPer100Kilometers.
func (a *FuelEfficiency) LitersPer100Kilometers() float64 {
	if a.liters_per100_kilometersLazy != nil {
		return *a.liters_per100_kilometersLazy
	}
	liters_per100_kilometers := a.convertFromBase(FuelEfficiencyLiterPer100Kilometers)
	a.liters_per100_kilometersLazy = &liters_per100_kilometers
	return liters_per100_kilometers
}

// MilePerUsGallon returns the value in MilePerUsGallon.
func (a *FuelEfficiency) MilesPerUsGallon() float64 {
	if a.miles_per_us_gallonLazy != nil {
		return *a.miles_per_us_gallonLazy
	}
	miles_per_us_gallon := a.convertFromBase(FuelEfficiencyMilePerUsGallon)
	a.miles_per_us_gallonLazy = &miles_per_us_gallon
	return miles_per_us_gallon
}

// MilePerUkGallon returns the value in MilePerUkGallon.
func (a *FuelEfficiency) MilesPerUkGallon() float64 {
	if a.miles_per_uk_gallonLazy != nil {
		return *a.miles_per_uk_gallonLazy
	}
	miles_per_uk_gallon := a.convertFromBase(FuelEfficiencyMilePerUkGallon)
	a.miles_per_uk_gallonLazy = &miles_per_uk_gallon
	return miles_per_uk_gallon
}

// KilometerPerLiter returns the value in KilometerPerLiter.
func (a *FuelEfficiency) KilometersPerLiters() float64 {
	if a.kilometers_per_litersLazy != nil {
		return *a.kilometers_per_litersLazy
	}
	kilometers_per_liters := a.convertFromBase(FuelEfficiencyKilometerPerLiter)
	a.kilometers_per_litersLazy = &kilometers_per_liters
	return kilometers_per_liters
}


// ToDto creates an FuelEfficiencyDto representation.
func (a *FuelEfficiency) ToDto(holdInUnit *FuelEfficiencyUnits) FuelEfficiencyDto {
	if holdInUnit == nil {
		defaultUnit := FuelEfficiencyLiterPer100Kilometers // Default value
		holdInUnit = &defaultUnit
	}

	return FuelEfficiencyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an FuelEfficiencyDto representation.
func (a *FuelEfficiency) ToDtoJSON(holdInUnit *FuelEfficiencyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts FuelEfficiency to a specific unit value.
func (a *FuelEfficiency) Convert(toUnit FuelEfficiencyUnits) float64 {
	switch toUnit { 
    case FuelEfficiencyLiterPer100Kilometers:
		return a.LitersPer100Kilometers()
    case FuelEfficiencyMilePerUsGallon:
		return a.MilesPerUsGallon()
    case FuelEfficiencyMilePerUkGallon:
		return a.MilesPerUkGallon()
    case FuelEfficiencyKilometerPerLiter:
		return a.KilometersPerLiters()
	default:
		return 0
	}
}

func (a *FuelEfficiency) convertFromBase(toUnit FuelEfficiencyUnits) float64 {
    value := a.value
	switch toUnit { 
	case FuelEfficiencyLiterPer100Kilometers:
		return (value) 
	case FuelEfficiencyMilePerUsGallon:
		return ((100 * 3.785411784) / (1.609344 * value)) 
	case FuelEfficiencyMilePerUkGallon:
		return ((100 * 4.54609188) / (1.609344 * value)) 
	case FuelEfficiencyKilometerPerLiter:
		return (100 / value) 
	default:
		return math.NaN()
	}
}

func (a *FuelEfficiency) convertToBase(value float64, fromUnit FuelEfficiencyUnits) float64 {
	switch fromUnit { 
	case FuelEfficiencyLiterPer100Kilometers:
		return (value) 
	case FuelEfficiencyMilePerUsGallon:
		return ((100 * 3.785411784) / (1.609344 * value)) 
	case FuelEfficiencyMilePerUkGallon:
		return ((100 * 4.54609188) / (1.609344 * value)) 
	case FuelEfficiencyKilometerPerLiter:
		return (100 / value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a FuelEfficiency) String() string {
	return a.ToString(FuelEfficiencyLiterPer100Kilometers, 2)
}

// ToString formats the FuelEfficiency to string.
// fractionalDigits -1 for not mention
func (a *FuelEfficiency) ToString(unit FuelEfficiencyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *FuelEfficiency) getUnitAbbreviation(unit FuelEfficiencyUnits) string {
	switch unit { 
	case FuelEfficiencyLiterPer100Kilometers:
		return "L/100km" 
	case FuelEfficiencyMilePerUsGallon:
		return "mpg (U.S.)" 
	case FuelEfficiencyMilePerUkGallon:
		return "mpg (imp.)" 
	case FuelEfficiencyKilometerPerLiter:
		return "km/L" 
	default:
		return ""
	}
}

// Check if the given FuelEfficiency are equals to the current FuelEfficiency
func (a *FuelEfficiency) Equals(other *FuelEfficiency) bool {
	return a.value == other.BaseValue()
}

// Check if the given FuelEfficiency are equals to the current FuelEfficiency
func (a *FuelEfficiency) CompareTo(other *FuelEfficiency) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given FuelEfficiency to the current FuelEfficiency.
func (a *FuelEfficiency) Add(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value + other.BaseValue()}
}

// Subtract the given FuelEfficiency to the current FuelEfficiency.
func (a *FuelEfficiency) Subtract(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value - other.BaseValue()}
}

// Multiply the given FuelEfficiency to the current FuelEfficiency.
func (a *FuelEfficiency) Multiply(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value * other.BaseValue()}
}

// Divide the given FuelEfficiency to the current FuelEfficiency.
func (a *FuelEfficiency) Divide(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value / other.BaseValue()}
}