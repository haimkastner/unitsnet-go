// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// FuelEfficiencyUnits defines various units of FuelEfficiency.
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

// FuelEfficiencyDto represents a FuelEfficiency measurement with a numerical value and its corresponding unit.
type FuelEfficiencyDto struct {
    // Value is the numerical representation of the FuelEfficiency.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the FuelEfficiency, as defined in the FuelEfficiencyUnits enumeration.
	Unit  FuelEfficiencyUnits `json:"unit"`
}

// FuelEfficiencyDtoFactory groups methods for creating and serializing FuelEfficiencyDto objects.
type FuelEfficiencyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a FuelEfficiencyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf FuelEfficiencyDtoFactory) FromJSON(data []byte) (*FuelEfficiencyDto, error) {
	a := FuelEfficiencyDto{}

    // Parse JSON into FuelEfficiencyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a FuelEfficiencyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a FuelEfficiencyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// FuelEfficiency represents a measurement in a of FuelEfficiency.
//
// Fuel efficiency is a form of thermal efficiency, meaning the ratio from effort to result of a process that converts chemical potential energy contained in a carrier (fuel) into kinetic energy or work. Fuel economy is stated as "fuel consumption" in liters per 100 kilometers (L/100 km). In countries using non-metric system, fuel economy is expressed in miles per gallon (mpg) (imperial galon or US galon).
type FuelEfficiency struct {
	// value is the base measurement stored internally.
	value       float64
    
    liters_per100_kilometersLazy *float64 
    miles_per_us_gallonLazy *float64 
    miles_per_uk_gallonLazy *float64 
    kilometers_per_litersLazy *float64 
}

// FuelEfficiencyFactory groups methods for creating FuelEfficiency instances.
type FuelEfficiencyFactory struct{}

// CreateFuelEfficiency creates a new FuelEfficiency instance from the given value and unit.
func (uf FuelEfficiencyFactory) CreateFuelEfficiency(value float64, unit FuelEfficiencyUnits) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, unit)
}

// FromDto converts a FuelEfficiencyDto to a FuelEfficiency instance.
func (uf FuelEfficiencyFactory) FromDto(dto FuelEfficiencyDto) (*FuelEfficiency, error) {
	return newFuelEfficiency(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a FuelEfficiency instance.
func (uf FuelEfficiencyFactory) FromDtoJSON(data []byte) (*FuelEfficiency, error) {
	unitDto, err := FuelEfficiencyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FuelEfficiencyDto from JSON: %w", err)
	}
	return FuelEfficiencyFactory{}.FromDto(*unitDto)
}


// FromLitersPer100Kilometers creates a new FuelEfficiency instance from a value in LitersPer100Kilometers.
func (uf FuelEfficiencyFactory) FromLitersPer100Kilometers(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyLiterPer100Kilometers)
}

// FromMilesPerUsGallon creates a new FuelEfficiency instance from a value in MilesPerUsGallon.
func (uf FuelEfficiencyFactory) FromMilesPerUsGallon(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyMilePerUsGallon)
}

// FromMilesPerUkGallon creates a new FuelEfficiency instance from a value in MilesPerUkGallon.
func (uf FuelEfficiencyFactory) FromMilesPerUkGallon(value float64) (*FuelEfficiency, error) {
	return newFuelEfficiency(value, FuelEfficiencyMilePerUkGallon)
}

// FromKilometersPerLiters creates a new FuelEfficiency instance from a value in KilometersPerLiters.
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

// BaseValue returns the base value of FuelEfficiency in LiterPer100Kilometers unit (the base/default quantity).
func (a *FuelEfficiency) BaseValue() float64 {
	return a.value
}


// LitersPer100Kilometers returns the FuelEfficiency value in LitersPer100Kilometers.
//
// 
func (a *FuelEfficiency) LitersPer100Kilometers() float64 {
	if a.liters_per100_kilometersLazy != nil {
		return *a.liters_per100_kilometersLazy
	}
	liters_per100_kilometers := a.convertFromBase(FuelEfficiencyLiterPer100Kilometers)
	a.liters_per100_kilometersLazy = &liters_per100_kilometers
	return liters_per100_kilometers
}

// MilesPerUsGallon returns the FuelEfficiency value in MilesPerUsGallon.
//
// 
func (a *FuelEfficiency) MilesPerUsGallon() float64 {
	if a.miles_per_us_gallonLazy != nil {
		return *a.miles_per_us_gallonLazy
	}
	miles_per_us_gallon := a.convertFromBase(FuelEfficiencyMilePerUsGallon)
	a.miles_per_us_gallonLazy = &miles_per_us_gallon
	return miles_per_us_gallon
}

// MilesPerUkGallon returns the FuelEfficiency value in MilesPerUkGallon.
//
// 
func (a *FuelEfficiency) MilesPerUkGallon() float64 {
	if a.miles_per_uk_gallonLazy != nil {
		return *a.miles_per_uk_gallonLazy
	}
	miles_per_uk_gallon := a.convertFromBase(FuelEfficiencyMilePerUkGallon)
	a.miles_per_uk_gallonLazy = &miles_per_uk_gallon
	return miles_per_uk_gallon
}

// KilometersPerLiters returns the FuelEfficiency value in KilometersPerLiters.
//
// 
func (a *FuelEfficiency) KilometersPerLiters() float64 {
	if a.kilometers_per_litersLazy != nil {
		return *a.kilometers_per_litersLazy
	}
	kilometers_per_liters := a.convertFromBase(FuelEfficiencyKilometerPerLiter)
	a.kilometers_per_litersLazy = &kilometers_per_liters
	return kilometers_per_liters
}


// ToDto creates a FuelEfficiencyDto representation from the FuelEfficiency instance.
//
// If the provided holdInUnit is nil, the value will be repesented by LiterPer100Kilometers by default.
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

// ToDtoJSON creates a JSON representation of the FuelEfficiency instance.
//
// If the provided holdInUnit is nil, the value will be repesented by LiterPer100Kilometers by default.
func (a *FuelEfficiency) ToDtoJSON(holdInUnit *FuelEfficiencyUnits) ([]byte, error) {
	// Convert to FuelEfficiencyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a FuelEfficiency to a specific unit value.
// The function uses the provided unit type (FuelEfficiencyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the FuelEfficiency in the default unit (LiterPer100Kilometers),
// formatted to two decimal places.
func (a FuelEfficiency) String() string {
	return a.ToString(FuelEfficiencyLiterPer100Kilometers, 2)
}

// ToString formats the FuelEfficiency value as a string with the specified unit and fractional digits.
// It converts the FuelEfficiency to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the FuelEfficiency value will be converted (e.g., LiterPer100Kilometers))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the FuelEfficiency with the unit abbreviation.
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

// Equals checks if the given FuelEfficiency is equal to the current FuelEfficiency.
//
// Parameters:
//    other: The FuelEfficiency to compare against.
//
// Returns:
//    bool: Returns true if both FuelEfficiency are equal, false otherwise.
func (a *FuelEfficiency) Equals(other *FuelEfficiency) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current FuelEfficiency with another FuelEfficiency.
// It returns -1 if the current FuelEfficiency is less than the other FuelEfficiency, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The FuelEfficiency to compare against.
//
// Returns:
//    int: -1 if the current FuelEfficiency is less, 1 if greater, and 0 if equal.
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

// Add adds the given FuelEfficiency to the current FuelEfficiency and returns the result.
// The result is a new FuelEfficiency instance with the sum of the values.
//
// Parameters:
//    other: The FuelEfficiency to add to the current FuelEfficiency.
//
// Returns:
//    *FuelEfficiency: A new FuelEfficiency instance representing the sum of both FuelEfficiency.
func (a *FuelEfficiency) Add(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given FuelEfficiency from the current FuelEfficiency and returns the result.
// The result is a new FuelEfficiency instance with the difference of the values.
//
// Parameters:
//    other: The FuelEfficiency to subtract from the current FuelEfficiency.
//
// Returns:
//    *FuelEfficiency: A new FuelEfficiency instance representing the difference of both FuelEfficiency.
func (a *FuelEfficiency) Subtract(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current FuelEfficiency by the given FuelEfficiency and returns the result.
// The result is a new FuelEfficiency instance with the product of the values.
//
// Parameters:
//    other: The FuelEfficiency to multiply with the current FuelEfficiency.
//
// Returns:
//    *FuelEfficiency: A new FuelEfficiency instance representing the product of both FuelEfficiency.
func (a *FuelEfficiency) Multiply(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value * other.BaseValue()}
}

// Divide divides the current FuelEfficiency by the given FuelEfficiency and returns the result.
// The result is a new FuelEfficiency instance with the quotient of the values.
//
// Parameters:
//    other: The FuelEfficiency to divide the current FuelEfficiency by.
//
// Returns:
//    *FuelEfficiency: A new FuelEfficiency instance representing the quotient of both FuelEfficiency.
func (a *FuelEfficiency) Divide(other *FuelEfficiency) *FuelEfficiency {
	return &FuelEfficiency{value: a.value / other.BaseValue()}
}