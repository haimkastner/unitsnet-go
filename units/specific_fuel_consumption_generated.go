// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificFuelConsumptionUnits defines various units of SpecificFuelConsumption.
type SpecificFuelConsumptionUnits string

const (
    
        // 
        SpecificFuelConsumptionPoundMassPerPoundForceHour SpecificFuelConsumptionUnits = "PoundMassPerPoundForceHour"
        // 
        SpecificFuelConsumptionKilogramPerKilogramForceHour SpecificFuelConsumptionUnits = "KilogramPerKilogramForceHour"
        // 
        SpecificFuelConsumptionGramPerKiloNewtonSecond SpecificFuelConsumptionUnits = "GramPerKiloNewtonSecond"
        // 
        SpecificFuelConsumptionKilogramPerKiloNewtonSecond SpecificFuelConsumptionUnits = "KilogramPerKiloNewtonSecond"
)

// SpecificFuelConsumptionDto represents a SpecificFuelConsumption measurement with a numerical value and its corresponding unit.
type SpecificFuelConsumptionDto struct {
    // Value is the numerical representation of the SpecificFuelConsumption.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the SpecificFuelConsumption, as defined in the SpecificFuelConsumptionUnits enumeration.
	Unit  SpecificFuelConsumptionUnits `json:"unit"`
}

// SpecificFuelConsumptionDtoFactory groups methods for creating and serializing SpecificFuelConsumptionDto objects.
type SpecificFuelConsumptionDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a SpecificFuelConsumptionDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf SpecificFuelConsumptionDtoFactory) FromJSON(data []byte) (*SpecificFuelConsumptionDto, error) {
	a := SpecificFuelConsumptionDto{}

    // Parse JSON into SpecificFuelConsumptionDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a SpecificFuelConsumptionDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a SpecificFuelConsumptionDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// SpecificFuelConsumption represents a measurement in a of SpecificFuelConsumption.
//
// SFC is the fuel efficiency of an engine design with respect to thrust output
type SpecificFuelConsumption struct {
	// value is the base measurement stored internally.
	value       float64
    
    pounds_mass_per_pound_force_hourLazy *float64 
    kilograms_per_kilogram_force_hourLazy *float64 
    grams_per_kilo_newton_secondLazy *float64 
    kilograms_per_kilo_newton_secondLazy *float64 
}

// SpecificFuelConsumptionFactory groups methods for creating SpecificFuelConsumption instances.
type SpecificFuelConsumptionFactory struct{}

// CreateSpecificFuelConsumption creates a new SpecificFuelConsumption instance from the given value and unit.
func (uf SpecificFuelConsumptionFactory) CreateSpecificFuelConsumption(value float64, unit SpecificFuelConsumptionUnits) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, unit)
}

// FromDto converts a SpecificFuelConsumptionDto to a SpecificFuelConsumption instance.
func (uf SpecificFuelConsumptionFactory) FromDto(dto SpecificFuelConsumptionDto) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a SpecificFuelConsumption instance.
func (uf SpecificFuelConsumptionFactory) FromDtoJSON(data []byte) (*SpecificFuelConsumption, error) {
	unitDto, err := SpecificFuelConsumptionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SpecificFuelConsumptionDto from JSON: %w", err)
	}
	return SpecificFuelConsumptionFactory{}.FromDto(*unitDto)
}


// FromPoundsMassPerPoundForceHour creates a new SpecificFuelConsumption instance from a value in PoundsMassPerPoundForceHour.
func (uf SpecificFuelConsumptionFactory) FromPoundsMassPerPoundForceHour(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionPoundMassPerPoundForceHour)
}

// FromKilogramsPerKilogramForceHour creates a new SpecificFuelConsumption instance from a value in KilogramsPerKilogramForceHour.
func (uf SpecificFuelConsumptionFactory) FromKilogramsPerKilogramForceHour(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionKilogramPerKilogramForceHour)
}

// FromGramsPerKiloNewtonSecond creates a new SpecificFuelConsumption instance from a value in GramsPerKiloNewtonSecond.
func (uf SpecificFuelConsumptionFactory) FromGramsPerKiloNewtonSecond(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionGramPerKiloNewtonSecond)
}

// FromKilogramsPerKiloNewtonSecond creates a new SpecificFuelConsumption instance from a value in KilogramsPerKiloNewtonSecond.
func (uf SpecificFuelConsumptionFactory) FromKilogramsPerKiloNewtonSecond(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionKilogramPerKiloNewtonSecond)
}


// newSpecificFuelConsumption creates a new SpecificFuelConsumption.
func newSpecificFuelConsumption(value float64, fromUnit SpecificFuelConsumptionUnits) (*SpecificFuelConsumption, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificFuelConsumption{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificFuelConsumption in GramPerKiloNewtonSecond unit (the base/default quantity).
func (a *SpecificFuelConsumption) BaseValue() float64 {
	return a.value
}


// PoundsMassPerPoundForceHour returns the SpecificFuelConsumption value in PoundsMassPerPoundForceHour.
//
// 
func (a *SpecificFuelConsumption) PoundsMassPerPoundForceHour() float64 {
	if a.pounds_mass_per_pound_force_hourLazy != nil {
		return *a.pounds_mass_per_pound_force_hourLazy
	}
	pounds_mass_per_pound_force_hour := a.convertFromBase(SpecificFuelConsumptionPoundMassPerPoundForceHour)
	a.pounds_mass_per_pound_force_hourLazy = &pounds_mass_per_pound_force_hour
	return pounds_mass_per_pound_force_hour
}

// KilogramsPerKilogramForceHour returns the SpecificFuelConsumption value in KilogramsPerKilogramForceHour.
//
// 
func (a *SpecificFuelConsumption) KilogramsPerKilogramForceHour() float64 {
	if a.kilograms_per_kilogram_force_hourLazy != nil {
		return *a.kilograms_per_kilogram_force_hourLazy
	}
	kilograms_per_kilogram_force_hour := a.convertFromBase(SpecificFuelConsumptionKilogramPerKilogramForceHour)
	a.kilograms_per_kilogram_force_hourLazy = &kilograms_per_kilogram_force_hour
	return kilograms_per_kilogram_force_hour
}

// GramsPerKiloNewtonSecond returns the SpecificFuelConsumption value in GramsPerKiloNewtonSecond.
//
// 
func (a *SpecificFuelConsumption) GramsPerKiloNewtonSecond() float64 {
	if a.grams_per_kilo_newton_secondLazy != nil {
		return *a.grams_per_kilo_newton_secondLazy
	}
	grams_per_kilo_newton_second := a.convertFromBase(SpecificFuelConsumptionGramPerKiloNewtonSecond)
	a.grams_per_kilo_newton_secondLazy = &grams_per_kilo_newton_second
	return grams_per_kilo_newton_second
}

// KilogramsPerKiloNewtonSecond returns the SpecificFuelConsumption value in KilogramsPerKiloNewtonSecond.
//
// 
func (a *SpecificFuelConsumption) KilogramsPerKiloNewtonSecond() float64 {
	if a.kilograms_per_kilo_newton_secondLazy != nil {
		return *a.kilograms_per_kilo_newton_secondLazy
	}
	kilograms_per_kilo_newton_second := a.convertFromBase(SpecificFuelConsumptionKilogramPerKiloNewtonSecond)
	a.kilograms_per_kilo_newton_secondLazy = &kilograms_per_kilo_newton_second
	return kilograms_per_kilo_newton_second
}


// ToDto creates a SpecificFuelConsumptionDto representation from the SpecificFuelConsumption instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GramPerKiloNewtonSecond by default.
func (a *SpecificFuelConsumption) ToDto(holdInUnit *SpecificFuelConsumptionUnits) SpecificFuelConsumptionDto {
	if holdInUnit == nil {
		defaultUnit := SpecificFuelConsumptionGramPerKiloNewtonSecond // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificFuelConsumptionDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the SpecificFuelConsumption instance.
//
// If the provided holdInUnit is nil, the value will be repesented by GramPerKiloNewtonSecond by default.
func (a *SpecificFuelConsumption) ToDtoJSON(holdInUnit *SpecificFuelConsumptionUnits) ([]byte, error) {
	// Convert to SpecificFuelConsumptionDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a SpecificFuelConsumption to a specific unit value.
// The function uses the provided unit type (SpecificFuelConsumptionUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *SpecificFuelConsumption) Convert(toUnit SpecificFuelConsumptionUnits) float64 {
	switch toUnit { 
    case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return a.PoundsMassPerPoundForceHour()
    case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return a.KilogramsPerKilogramForceHour()
    case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return a.GramsPerKiloNewtonSecond()
    case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return a.KilogramsPerKiloNewtonSecond()
	default:
		return math.NaN()
	}
}

func (a *SpecificFuelConsumption) convertFromBase(toUnit SpecificFuelConsumptionUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return (value / 28.33) 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return (value / 28.33) 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return (value) 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *SpecificFuelConsumption) convertToBase(value float64, fromUnit SpecificFuelConsumptionUnits) float64 {
	switch fromUnit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return (value * 28.33) 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return (value * 28.33) 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return (value) 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the SpecificFuelConsumption in the default unit (GramPerKiloNewtonSecond),
// formatted to two decimal places.
func (a SpecificFuelConsumption) String() string {
	return a.ToString(SpecificFuelConsumptionGramPerKiloNewtonSecond, 2)
}

// ToString formats the SpecificFuelConsumption value as a string with the specified unit and fractional digits.
// It converts the SpecificFuelConsumption to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the SpecificFuelConsumption value will be converted (e.g., GramPerKiloNewtonSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the SpecificFuelConsumption with the unit abbreviation.
func (a *SpecificFuelConsumption) ToString(unit SpecificFuelConsumptionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetSpecificFuelConsumptionAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetSpecificFuelConsumptionAbbreviation(unit))
}

// Equals checks if the given SpecificFuelConsumption is equal to the current SpecificFuelConsumption.
//
// Parameters:
//    other: The SpecificFuelConsumption to compare against.
//
// Returns:
//    bool: Returns true if both SpecificFuelConsumption are equal, false otherwise.
func (a *SpecificFuelConsumption) Equals(other *SpecificFuelConsumption) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current SpecificFuelConsumption with another SpecificFuelConsumption.
// It returns -1 if the current SpecificFuelConsumption is less than the other SpecificFuelConsumption, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The SpecificFuelConsumption to compare against.
//
// Returns:
//    int: -1 if the current SpecificFuelConsumption is less, 1 if greater, and 0 if equal.
func (a *SpecificFuelConsumption) CompareTo(other *SpecificFuelConsumption) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given SpecificFuelConsumption to the current SpecificFuelConsumption and returns the result.
// The result is a new SpecificFuelConsumption instance with the sum of the values.
//
// Parameters:
//    other: The SpecificFuelConsumption to add to the current SpecificFuelConsumption.
//
// Returns:
//    *SpecificFuelConsumption: A new SpecificFuelConsumption instance representing the sum of both SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Add(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given SpecificFuelConsumption from the current SpecificFuelConsumption and returns the result.
// The result is a new SpecificFuelConsumption instance with the difference of the values.
//
// Parameters:
//    other: The SpecificFuelConsumption to subtract from the current SpecificFuelConsumption.
//
// Returns:
//    *SpecificFuelConsumption: A new SpecificFuelConsumption instance representing the difference of both SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Subtract(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current SpecificFuelConsumption by the given SpecificFuelConsumption and returns the result.
// The result is a new SpecificFuelConsumption instance with the product of the values.
//
// Parameters:
//    other: The SpecificFuelConsumption to multiply with the current SpecificFuelConsumption.
//
// Returns:
//    *SpecificFuelConsumption: A new SpecificFuelConsumption instance representing the product of both SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Multiply(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value * other.BaseValue()}
}

// Divide divides the current SpecificFuelConsumption by the given SpecificFuelConsumption and returns the result.
// The result is a new SpecificFuelConsumption instance with the quotient of the values.
//
// Parameters:
//    other: The SpecificFuelConsumption to divide the current SpecificFuelConsumption by.
//
// Returns:
//    *SpecificFuelConsumption: A new SpecificFuelConsumption instance representing the quotient of both SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Divide(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value / other.BaseValue()}
}

// GetSpecificFuelConsumptionAbbreviation gets the unit abbreviation.
func GetSpecificFuelConsumptionAbbreviation(unit SpecificFuelConsumptionUnits) string {
	switch unit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return "lb/(lbf·h)" 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return "kg/(kgf�h)" 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return "g/(kN�s)" 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return "kg/(kN�s)" 
	default:
		return ""
	}
}