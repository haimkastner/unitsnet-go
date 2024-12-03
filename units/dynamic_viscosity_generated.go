// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// DynamicViscosityUnits enumeration
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

// DynamicViscosityDto represents an DynamicViscosity
type DynamicViscosityDto struct {
	Value float64
	Unit  DynamicViscosityUnits
}

// DynamicViscosityDtoFactory struct to group related functions
type DynamicViscosityDtoFactory struct{}

func (udf DynamicViscosityDtoFactory) FromJSON(data []byte) (*DynamicViscosityDto, error) {
	a := DynamicViscosityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a DynamicViscosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// DynamicViscosity struct
type DynamicViscosity struct {
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

// DynamicViscosityFactory struct to group related functions
type DynamicViscosityFactory struct{}

func (uf DynamicViscosityFactory) CreateDynamicViscosity(value float64, unit DynamicViscosityUnits) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, unit)
}

func (uf DynamicViscosityFactory) FromDto(dto DynamicViscosityDto) (*DynamicViscosity, error) {
	return newDynamicViscosity(dto.Value, dto.Unit)
}

func (uf DynamicViscosityFactory) FromDtoJSON(data []byte) (*DynamicViscosity, error) {
	unitDto, err := DynamicViscosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return DynamicViscosityFactory{}.FromDto(*unitDto)
}


// FromNewtonSecondPerMeterSquared creates a new DynamicViscosity instance from NewtonSecondPerMeterSquared.
func (uf DynamicViscosityFactory) FromNewtonSecondsPerMeterSquared(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityNewtonSecondPerMeterSquared)
}

// FromPascalSecond creates a new DynamicViscosity instance from PascalSecond.
func (uf DynamicViscosityFactory) FromPascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPascalSecond)
}

// FromPoise creates a new DynamicViscosity instance from Poise.
func (uf DynamicViscosityFactory) FromPoise(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoise)
}

// FromReyn creates a new DynamicViscosity instance from Reyn.
func (uf DynamicViscosityFactory) FromReyns(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityReyn)
}

// FromPoundForceSecondPerSquareInch creates a new DynamicViscosity instance from PoundForceSecondPerSquareInch.
func (uf DynamicViscosityFactory) FromPoundsForceSecondPerSquareInch(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundForceSecondPerSquareInch)
}

// FromPoundForceSecondPerSquareFoot creates a new DynamicViscosity instance from PoundForceSecondPerSquareFoot.
func (uf DynamicViscosityFactory) FromPoundsForceSecondPerSquareFoot(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundForceSecondPerSquareFoot)
}

// FromPoundPerFootSecond creates a new DynamicViscosity instance from PoundPerFootSecond.
func (uf DynamicViscosityFactory) FromPoundsPerFootSecond(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityPoundPerFootSecond)
}

// FromMillipascalSecond creates a new DynamicViscosity instance from MillipascalSecond.
func (uf DynamicViscosityFactory) FromMillipascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityMillipascalSecond)
}

// FromMicropascalSecond creates a new DynamicViscosity instance from MicropascalSecond.
func (uf DynamicViscosityFactory) FromMicropascalSeconds(value float64) (*DynamicViscosity, error) {
	return newDynamicViscosity(value, DynamicViscosityMicropascalSecond)
}

// FromCentipoise creates a new DynamicViscosity instance from Centipoise.
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

// BaseValue returns the base value of DynamicViscosity in NewtonSecondPerMeterSquared.
func (a *DynamicViscosity) BaseValue() float64 {
	return a.value
}


// NewtonSecondPerMeterSquared returns the value in NewtonSecondPerMeterSquared.
func (a *DynamicViscosity) NewtonSecondsPerMeterSquared() float64 {
	if a.newton_seconds_per_meter_squaredLazy != nil {
		return *a.newton_seconds_per_meter_squaredLazy
	}
	newton_seconds_per_meter_squared := a.convertFromBase(DynamicViscosityNewtonSecondPerMeterSquared)
	a.newton_seconds_per_meter_squaredLazy = &newton_seconds_per_meter_squared
	return newton_seconds_per_meter_squared
}

// PascalSecond returns the value in PascalSecond.
func (a *DynamicViscosity) PascalSeconds() float64 {
	if a.pascal_secondsLazy != nil {
		return *a.pascal_secondsLazy
	}
	pascal_seconds := a.convertFromBase(DynamicViscosityPascalSecond)
	a.pascal_secondsLazy = &pascal_seconds
	return pascal_seconds
}

// Poise returns the value in Poise.
func (a *DynamicViscosity) Poise() float64 {
	if a.poiseLazy != nil {
		return *a.poiseLazy
	}
	poise := a.convertFromBase(DynamicViscosityPoise)
	a.poiseLazy = &poise
	return poise
}

// Reyn returns the value in Reyn.
func (a *DynamicViscosity) Reyns() float64 {
	if a.reynsLazy != nil {
		return *a.reynsLazy
	}
	reyns := a.convertFromBase(DynamicViscosityReyn)
	a.reynsLazy = &reyns
	return reyns
}

// PoundForceSecondPerSquareInch returns the value in PoundForceSecondPerSquareInch.
func (a *DynamicViscosity) PoundsForceSecondPerSquareInch() float64 {
	if a.pounds_force_second_per_square_inchLazy != nil {
		return *a.pounds_force_second_per_square_inchLazy
	}
	pounds_force_second_per_square_inch := a.convertFromBase(DynamicViscosityPoundForceSecondPerSquareInch)
	a.pounds_force_second_per_square_inchLazy = &pounds_force_second_per_square_inch
	return pounds_force_second_per_square_inch
}

// PoundForceSecondPerSquareFoot returns the value in PoundForceSecondPerSquareFoot.
func (a *DynamicViscosity) PoundsForceSecondPerSquareFoot() float64 {
	if a.pounds_force_second_per_square_footLazy != nil {
		return *a.pounds_force_second_per_square_footLazy
	}
	pounds_force_second_per_square_foot := a.convertFromBase(DynamicViscosityPoundForceSecondPerSquareFoot)
	a.pounds_force_second_per_square_footLazy = &pounds_force_second_per_square_foot
	return pounds_force_second_per_square_foot
}

// PoundPerFootSecond returns the value in PoundPerFootSecond.
func (a *DynamicViscosity) PoundsPerFootSecond() float64 {
	if a.pounds_per_foot_secondLazy != nil {
		return *a.pounds_per_foot_secondLazy
	}
	pounds_per_foot_second := a.convertFromBase(DynamicViscosityPoundPerFootSecond)
	a.pounds_per_foot_secondLazy = &pounds_per_foot_second
	return pounds_per_foot_second
}

// MillipascalSecond returns the value in MillipascalSecond.
func (a *DynamicViscosity) MillipascalSeconds() float64 {
	if a.millipascal_secondsLazy != nil {
		return *a.millipascal_secondsLazy
	}
	millipascal_seconds := a.convertFromBase(DynamicViscosityMillipascalSecond)
	a.millipascal_secondsLazy = &millipascal_seconds
	return millipascal_seconds
}

// MicropascalSecond returns the value in MicropascalSecond.
func (a *DynamicViscosity) MicropascalSeconds() float64 {
	if a.micropascal_secondsLazy != nil {
		return *a.micropascal_secondsLazy
	}
	micropascal_seconds := a.convertFromBase(DynamicViscosityMicropascalSecond)
	a.micropascal_secondsLazy = &micropascal_seconds
	return micropascal_seconds
}

// Centipoise returns the value in Centipoise.
func (a *DynamicViscosity) Centipoise() float64 {
	if a.centipoiseLazy != nil {
		return *a.centipoiseLazy
	}
	centipoise := a.convertFromBase(DynamicViscosityCentipoise)
	a.centipoiseLazy = &centipoise
	return centipoise
}


// ToDto creates an DynamicViscosityDto representation.
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

// ToDtoJSON creates an DynamicViscosityDto representation.
func (a *DynamicViscosity) ToDtoJSON(holdInUnit *DynamicViscosityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts DynamicViscosity to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a DynamicViscosity) String() string {
	return a.ToString(DynamicViscosityNewtonSecondPerMeterSquared, 2)
}

// ToString formats the DynamicViscosity to string.
// fractionalDigits -1 for not mention
func (a *DynamicViscosity) ToString(unit DynamicViscosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *DynamicViscosity) getUnitAbbreviation(unit DynamicViscosityUnits) string {
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

// Check if the given DynamicViscosity are equals to the current DynamicViscosity
func (a *DynamicViscosity) Equals(other *DynamicViscosity) bool {
	return a.value == other.BaseValue()
}

// Check if the given DynamicViscosity are equals to the current DynamicViscosity
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

// Add the given DynamicViscosity to the current DynamicViscosity.
func (a *DynamicViscosity) Add(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value + other.BaseValue()}
}

// Subtract the given DynamicViscosity to the current DynamicViscosity.
func (a *DynamicViscosity) Subtract(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value - other.BaseValue()}
}

// Multiply the given DynamicViscosity to the current DynamicViscosity.
func (a *DynamicViscosity) Multiply(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value * other.BaseValue()}
}

// Divide the given DynamicViscosity to the current DynamicViscosity.
func (a *DynamicViscosity) Divide(other *DynamicViscosity) *DynamicViscosity {
	return &DynamicViscosity{value: a.value / other.BaseValue()}
}