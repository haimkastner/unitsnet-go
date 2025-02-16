// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// DynamicViscosityUnits defines various units of DynamicViscosity.
type DynamicViscosityUnits string

const (
    
        // 
        DynamicViscosityNewtonSecondPerMeterSquared DynamicViscosityUnits = "NewtonSecondPerMeterSquared"
        // 
        DynamicViscosityPascalSecond DynamicViscosityUnits = "PascalSecond"
        // 
        DynamicViscosityPoise DynamicViscosityUnits = "Poise"
        // 
        DynamicViscosityReyn DynamicViscosityUnits = "Reyn"
        // 
        DynamicViscosityPoundForceSecondPerSquareInch DynamicViscosityUnits = "PoundForceSecondPerSquareInch"
        // 
        DynamicViscosityPoundForceSecondPerSquareFoot DynamicViscosityUnits = "PoundForceSecondPerSquareFoot"
        // 
        DynamicViscosityPoundPerFootSecond DynamicViscosityUnits = "PoundPerFootSecond"
        // 
        DynamicViscosityMillipascalSecond DynamicViscosityUnits = "MillipascalSecond"
        // 
        DynamicViscosityMicropascalSecond DynamicViscosityUnits = "MicropascalSecond"
        // 
        DynamicViscosityCentipoise DynamicViscosityUnits = "Centipoise"
)

// DynamicViscosityDto represents a DynamicViscosity measurement with a numerical value and its corresponding unit.
type DynamicViscosityDto struct {
    // Value is the numerical representation of the DynamicViscosity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the DynamicViscosity, as defined in the DynamicViscosityUnits enumeration.
	Unit  DynamicViscosityUnits `json:"unit"`
}

// DynamicViscosityDtoFactory groups methods for creating and serializing DynamicViscosityDto objects.
type DynamicViscosityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a DynamicViscosityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf DynamicViscosityDtoFactory) FromJSON(data []byte) (*DynamicViscosityDto, error) {
	a := DynamicViscosityDto{}

    // Parse JSON into DynamicViscosityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a DynamicViscosityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a DynamicViscosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// DynamicViscosity represents a measurement in a of DynamicViscosity.
//
// The dynamic (shear) viscosity of a fluid expresses its resistance to shearing flows, where adjacent layers move parallel to each other with different speeds
type DynamicViscosity struct {
	// value is the base measurement stored internally.
	value       float64
    
    newton_seconds_per_meter_squaredLazy *float64 
    pascal_secondsLazy *float64 
    poiseLazy *float64 
    reynsLazy *float64 
    pounds_force_second_per_square_inchLazy *float64 
    pounds_force_second_per_square_footLazy *float64 
    pounds_per_foot_secondLazy *float64 
    millipascal_secondsLazy *float64 
    micropascal_secondsLazy *float64 
    centipoiseLazy *float64 
}

// DynamicViscosityFactory groups methods for creating DynamicViscosity instances.
type DynamicViscosityFactory struct{}

// CreateDynamicViscosity creates a new DynamicViscosity instance from the given value and unit.
func (uf DynamicViscosityFactory) CreateDynamicViscosity(value float64, unit DynamicViscosityUnits) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, unit)
}

// FromDto converts a DynamicViscosityDto to a DynamicViscosity instance.
func (uf DynamicViscosityFactory) FromDto(dto DynamicViscosityDto) (*DynamicViscosity, error) {
	return newDynamicViscosity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a DynamicViscosity instance.
func (uf DynamicViscosityFactory) FromDtoJSON(data []byte) (*DynamicViscosity, error) {
	unitDto, err := DynamicViscosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DynamicViscosityDto from JSON: %w", err)
	}
	return DynamicViscosityFactory{}.FromDto(*unitDto)
}


// FromNewtonSecondsPerMeterSquared creates a new DynamicViscosity instance from a value in NewtonSecondsPerMeterSquared.
func (uf DynamicViscosityFactory) FromNewtonSecondsPerMeterSquared(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityNewtonSecondPerMeterSquared)
}

// FromPascalSeconds creates a new DynamicViscosity instance from a value in PascalSeconds.
func (uf DynamicViscosityFactory) FromPascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPascalSecond)
}

// FromPoise creates a new DynamicViscosity instance from a value in Poise.
func (uf DynamicViscosityFactory) FromPoise(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoise)
}

// FromReyns creates a new DynamicViscosity instance from a value in Reyns.
func (uf DynamicViscosityFactory) FromReyns(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityReyn)
}

// FromPoundsForceSecondPerSquareInch creates a new DynamicViscosity instance from a value in PoundsForceSecondPerSquareInch.
func (uf DynamicViscosityFactory) FromPoundsForceSecondPerSquareInch(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundForceSecondPerSquareInch)
}

// FromPoundsForceSecondPerSquareFoot creates a new DynamicViscosity instance from a value in PoundsForceSecondPerSquareFoot.
func (uf DynamicViscosityFactory) FromPoundsForceSecondPerSquareFoot(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundForceSecondPerSquareFoot)
}

// FromPoundsPerFootSecond creates a new DynamicViscosity instance from a value in PoundsPerFootSecond.
func (uf DynamicViscosityFactory) FromPoundsPerFootSecond(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundPerFootSecond)
}

// FromMillipascalSeconds creates a new DynamicViscosity instance from a value in MillipascalSeconds.
func (uf DynamicViscosityFactory) FromMillipascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityMillipascalSecond)
}

// FromMicropascalSeconds creates a new DynamicViscosity instance from a value in MicropascalSeconds.
func (uf DynamicViscosityFactory) FromMicropascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityMicropascalSecond)
}

// FromCentipoise creates a new DynamicViscosity instance from a value in Centipoise.
func (uf DynamicViscosityFactory) FromCentipoise(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityCentipoise)
}


// newDynamicViscosity creates a new DynamicViscosity.
func newDynamicViscosity(value float64, fromUnit DynamicViscosityUnits) (*DynamicViscosity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &DynamicViscosity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of DynamicViscosity in NewtonSecondPerMeterSquared unit (the base/default quantity).
func (a *DynamicViscosity) BaseValue() float64 {
	return a.value
}


// NewtonSecondsPerMeterSquared returns the DynamicViscosity value in NewtonSecondsPerMeterSquared.
//
// 
func (a *DynamicViscosity) NewtonSecondsPerMeterSquared() float64 {
	if a.newton_seconds_per_meter_squaredLazy != nil {
		return *a.newton_seconds_per_meter_squaredLazy
	}
	newton_seconds_per_meter_squared := a.convertFromBase(DynamicViscosityNewtonSecondPerMeterSquared)
	a.newton_seconds_per_meter_squaredLazy = &newton_seconds_per_meter_squared
	return newton_seconds_per_meter_squared
}

// PascalSeconds returns the DynamicViscosity value in PascalSeconds.
//
// 
func (a *DynamicViscosity) PascalSeconds() float64 {
	if a.pascal_secondsLazy != nil {
		return *a.pascal_secondsLazy
	}
	pascal_seconds := a.convertFromBase(DynamicViscosityPascalSecond)
	a.pascal_secondsLazy = &pascal_seconds
	return pascal_seconds
}

// Poise returns the DynamicViscosity value in Poise.
//
// 
func (a *DynamicViscosity) Poise() float64 {
	if a.poiseLazy != nil {
		return *a.poiseLazy
	}
	poise := a.convertFromBase(DynamicViscosityPoise)
	a.poiseLazy = &poise
	return poise
}

// Reyns returns the DynamicViscosity value in Reyns.
//
// 
func (a *DynamicViscosity) Reyns() float64 {
	if a.reynsLazy != nil {
		return *a.reynsLazy
	}
	reyns := a.convertFromBase(DynamicViscosityReyn)
	a.reynsLazy = &reyns
	return reyns
}

// PoundsForceSecondPerSquareInch returns the DynamicViscosity value in PoundsForceSecondPerSquareInch.
//
// 
func (a *DynamicViscosity) PoundsForceSecondPerSquareInch() float64 {
	if a.pounds_force_second_per_square_inchLazy != nil {
		return *a.pounds_force_second_per_square_inchLazy
	}
	pounds_force_second_per_square_inch := a.convertFromBase(DynamicViscosityPoundForceSecondPerSquareInch)
	a.pounds_force_second_per_square_inchLazy = &pounds_force_second_per_square_inch
	return pounds_force_second_per_square_inch
}

// PoundsForceSecondPerSquareFoot returns the DynamicViscosity value in PoundsForceSecondPerSquareFoot.
//
// 
func (a *DynamicViscosity) PoundsForceSecondPerSquareFoot() float64 {
	if a.pounds_force_second_per_square_footLazy != nil {
		return *a.pounds_force_second_per_square_footLazy
	}
	pounds_force_second_per_square_foot := a.convertFromBase(DynamicViscosityPoundForceSecondPerSquareFoot)
	a.pounds_force_second_per_square_footLazy = &pounds_force_second_per_square_foot
	return pounds_force_second_per_square_foot
}

// PoundsPerFootSecond returns the DynamicViscosity value in PoundsPerFootSecond.
//
// 
func (a *DynamicViscosity) PoundsPerFootSecond() float64 {
	if a.pounds_per_foot_secondLazy != nil {
		return *a.pounds_per_foot_secondLazy
	}
	pounds_per_foot_second := a.convertFromBase(DynamicViscosityPoundPerFootSecond)
	a.pounds_per_foot_secondLazy = &pounds_per_foot_second
	return pounds_per_foot_second
}

// MillipascalSeconds returns the DynamicViscosity value in MillipascalSeconds.
//
// 
func (a *DynamicViscosity) MillipascalSeconds() float64 {
	if a.millipascal_secondsLazy != nil {
		return *a.millipascal_secondsLazy
	}
	millipascal_seconds := a.convertFromBase(DynamicViscosityMillipascalSecond)
	a.millipascal_secondsLazy = &millipascal_seconds
	return millipascal_seconds
}

// MicropascalSeconds returns the DynamicViscosity value in MicropascalSeconds.
//
// 
func (a *DynamicViscosity) MicropascalSeconds() float64 {
	if a.micropascal_secondsLazy != nil {
		return *a.micropascal_secondsLazy
	}
	micropascal_seconds := a.convertFromBase(DynamicViscosityMicropascalSecond)
	a.micropascal_secondsLazy = &micropascal_seconds
	return micropascal_seconds
}

// Centipoise returns the DynamicViscosity value in Centipoise.
//
// 
func (a *DynamicViscosity) Centipoise() float64 {
	if a.centipoiseLazy != nil {
		return *a.centipoiseLazy
	}
	centipoise := a.convertFromBase(DynamicViscosityCentipoise)
	a.centipoiseLazy = &centipoise
	return centipoise
}


// ToDto creates a DynamicViscosityDto representation from the DynamicViscosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonSecondPerMeterSquared by default.
func (a *DynamicViscosity) ToDto(holdInUnit *DynamicViscosityUnits) DynamicViscosityDto {
	if holdInUnit == nil {
		defaultUnit := DynamicViscosityNewtonSecondPerMeterSquared // Default value
		holdInUnit = &defaultUnit
	}

	return DynamicViscosityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the DynamicViscosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by NewtonSecondPerMeterSquared by default.
func (a *DynamicViscosity) ToDtoJSON(holdInUnit *DynamicViscosityUnits) ([]byte, error) {
	// Convert to DynamicViscosityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a DynamicViscosity to a specific unit value.
// The function uses the provided unit type (DynamicViscosityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *DynamicViscosity) Convert(toUnit DynamicViscosityUnits) float64 {
	switch toUnit { 
    case DynamicViscosityNewtonSecondPerMeterSquared:
		return a.NewtonSecondsPerMeterSquared()
    case DynamicViscosityPascalSecond:
		return a.PascalSeconds()
    case DynamicViscosityPoise:
		return a.Poise()
    case DynamicViscosityReyn:
		return a.Reyns()
    case DynamicViscosityPoundForceSecondPerSquareInch:
		return a.PoundsForceSecondPerSquareInch()
    case DynamicViscosityPoundForceSecondPerSquareFoot:
		return a.PoundsForceSecondPerSquareFoot()
    case DynamicViscosityPoundPerFootSecond:
		return a.PoundsPerFootSecond()
    case DynamicViscosityMillipascalSecond:
		return a.MillipascalSeconds()
    case DynamicViscosityMicropascalSecond:
		return a.MicropascalSeconds()
    case DynamicViscosityCentipoise:
		return a.Centipoise()
	default:
		return math.NaN()
	}
}

func (a *DynamicViscosity) convertFromBase(toUnit DynamicViscosityUnits) float64 {
    value := a.value
	switch toUnit { 
	case DynamicViscosityNewtonSecondPerMeterSquared:
		return (value) 
	case DynamicViscosityPascalSecond:
		return (value) 
	case DynamicViscosityPoise:
		return (value * 10) 
	case DynamicViscosityReyn:
		return (value / 6.8947572931683613e3) 
	case DynamicViscosityPoundForceSecondPerSquareInch:
		return (value / 6.8947572931683613e3) 
	case DynamicViscosityPoundForceSecondPerSquareFoot:
		return (value / 4.7880258980335843e1) 
	case DynamicViscosityPoundPerFootSecond:
		return (value / 1.4881639) 
	case DynamicViscosityMillipascalSecond:
		return ((value) / 0.001) 
	case DynamicViscosityMicropascalSecond:
		return ((value) / 1e-06) 
	case DynamicViscosityCentipoise:
		return ((value * 10) / 0.01) 
	default:
		return math.NaN()
	}
}

func (a *DynamicViscosity) convertToBase(value float64, fromUnit DynamicViscosityUnits) float64 {
	switch fromUnit { 
	case DynamicViscosityNewtonSecondPerMeterSquared:
		return (value) 
	case DynamicViscosityPascalSecond:
		return (value) 
	case DynamicViscosityPoise:
		return (value / 10) 
	case DynamicViscosityReyn:
		return (value * 6.8947572931683613e3) 
	case DynamicViscosityPoundForceSecondPerSquareInch:
		return (value * 6.8947572931683613e3) 
	case DynamicViscosityPoundForceSecondPerSquareFoot:
		return (value * 4.7880258980335843e1) 
	case DynamicViscosityPoundPerFootSecond:
		return (value * 1.4881639) 
	case DynamicViscosityMillipascalSecond:
		return ((value) * 0.001) 
	case DynamicViscosityMicropascalSecond:
		return ((value) * 1e-06) 
	case DynamicViscosityCentipoise:
		return ((value / 10) * 0.01) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the DynamicViscosity in the default unit (NewtonSecondPerMeterSquared),
// formatted to two decimal places.
func (a DynamicViscosity) String() string {
	return a.ToString(DynamicViscosityNewtonSecondPerMeterSquared, 2)
}

// ToString formats the DynamicViscosity value as a string with the specified unit and fractional digits.
// It converts the DynamicViscosity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the DynamicViscosity value will be converted (e.g., NewtonSecondPerMeterSquared))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the DynamicViscosity with the unit abbreviation.
func (a *DynamicViscosity) ToString(unit DynamicViscosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetDynamicViscosityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetDynamicViscosityAbbreviation(unit))
}

// Equals checks if the given DynamicViscosity is equal to the current DynamicViscosity.
//
// Parameters:
//    other: The DynamicViscosity to compare against.
//
// Returns:
//    bool: Returns true if both DynamicViscosity are equal, false otherwise.
func (a *DynamicViscosity) Equals(other *DynamicViscosity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current DynamicViscosity with another DynamicViscosity.
// It returns -1 if the current DynamicViscosity is less than the other DynamicViscosity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The DynamicViscosity to compare against.
//
// Returns:
//    int: -1 if the current DynamicViscosity is less, 1 if greater, and 0 if equal.
func (a *DynamicViscosity) CompareTo(other *DynamicViscosity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given DynamicViscosity to the current DynamicViscosity and returns the result.
// The result is a new DynamicViscosity instance with the sum of the values.
//
// Parameters:
//    other: The DynamicViscosity to add to the current DynamicViscosity.
//
// Returns:
//    *DynamicViscosity: A new DynamicViscosity instance representing the sum of both DynamicViscosity.
func (a *DynamicViscosity) Add(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given DynamicViscosity from the current DynamicViscosity and returns the result.
// The result is a new DynamicViscosity instance with the difference of the values.
//
// Parameters:
//    other: The DynamicViscosity to subtract from the current DynamicViscosity.
//
// Returns:
//    *DynamicViscosity: A new DynamicViscosity instance representing the difference of both DynamicViscosity.
func (a *DynamicViscosity) Subtract(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current DynamicViscosity by the given DynamicViscosity and returns the result.
// The result is a new DynamicViscosity instance with the product of the values.
//
// Parameters:
//    other: The DynamicViscosity to multiply with the current DynamicViscosity.
//
// Returns:
//    *DynamicViscosity: A new DynamicViscosity instance representing the product of both DynamicViscosity.
func (a *DynamicViscosity) Multiply(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value * other.BaseValue()}
}

// Divide divides the current DynamicViscosity by the given DynamicViscosity and returns the result.
// The result is a new DynamicViscosity instance with the quotient of the values.
//
// Parameters:
//    other: The DynamicViscosity to divide the current DynamicViscosity by.
//
// Returns:
//    *DynamicViscosity: A new DynamicViscosity instance representing the quotient of both DynamicViscosity.
func (a *DynamicViscosity) Divide(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value / other.BaseValue()}
}

// GetDynamicViscosityAbbreviation gets the unit abbreviation.
func GetDynamicViscosityAbbreviation(unit DynamicViscosityUnits) string {
	switch unit { 
	case DynamicViscosityNewtonSecondPerMeterSquared:
		return "Ns/m²" 
	case DynamicViscosityPascalSecond:
		return "Pa·s" 
	case DynamicViscosityPoise:
		return "P" 
	case DynamicViscosityReyn:
		return "reyn" 
	case DynamicViscosityPoundForceSecondPerSquareInch:
		return "lbf·s/in²" 
	case DynamicViscosityPoundForceSecondPerSquareFoot:
		return "lbf·s/ft²" 
	case DynamicViscosityPoundPerFootSecond:
		return "lb/ft·s" 
	case DynamicViscosityMillipascalSecond:
		return "mPa·s" 
	case DynamicViscosityMicropascalSecond:
		return "μPa·s" 
	case DynamicViscosityCentipoise:
		return "cP" 
	default:
		return ""
	}
}