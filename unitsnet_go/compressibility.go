package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// CompressibilityUnits enumeration
type CompressibilityUnits string

const (
    
        // 
        CompressibilityInversePascal CompressibilityUnits = "InversePascal"
        // 
        CompressibilityInverseKilopascal CompressibilityUnits = "InverseKilopascal"
        // 
        CompressibilityInverseMegapascal CompressibilityUnits = "InverseMegapascal"
        // 
        CompressibilityInverseAtmosphere CompressibilityUnits = "InverseAtmosphere"
        // 
        CompressibilityInverseMillibar CompressibilityUnits = "InverseMillibar"
        // 
        CompressibilityInverseBar CompressibilityUnits = "InverseBar"
        // 
        CompressibilityInversePoundForcePerSquareInch CompressibilityUnits = "InversePoundForcePerSquareInch"
)

// CompressibilityDto represents an Compressibility
type CompressibilityDto struct {
	Value float64
	Unit  CompressibilityUnits
}

// CompressibilityDtoFactory struct to group related functions
type CompressibilityDtoFactory struct{}

func (udf CompressibilityDtoFactory) FromJSON(data []byte) (*CompressibilityDto, error) {
	a := CompressibilityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a CompressibilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Compressibility struct
type Compressibility struct {
	value       float64
    
    inverse_pascalsLazy *float64 
    inverse_kilopascalsLazy *float64 
    inverse_megapascalsLazy *float64 
    inverse_atmospheresLazy *float64 
    inverse_millibarsLazy *float64 
    inverse_barsLazy *float64 
    inverse_pounds_force_per_square_inchLazy *float64 
}

// CompressibilityFactory struct to group related functions
type CompressibilityFactory struct{}

func (uf CompressibilityFactory) CreateCompressibility(value float64, unit CompressibilityUnits) (*Compressibility, error) {
	return newCompressibility(value, unit)
}

func (uf CompressibilityFactory) FromDto(dto CompressibilityDto) (*Compressibility, error) {
	return newCompressibility(dto.Value, dto.Unit)
}

func (uf CompressibilityFactory) FromDtoJSON(data []byte) (*Compressibility, error) {
	unitDto, err := CompressibilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return CompressibilityFactory{}.FromDto(*unitDto)
}


// FromInversePascal creates a new Compressibility instance from InversePascal.
func (uf CompressibilityFactory) FromInversePascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInversePascal)
}

// FromInverseKilopascal creates a new Compressibility instance from InverseKilopascal.
func (uf CompressibilityFactory) FromInverseKilopascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseKilopascal)
}

// FromInverseMegapascal creates a new Compressibility instance from InverseMegapascal.
func (uf CompressibilityFactory) FromInverseMegapascals(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseMegapascal)
}

// FromInverseAtmosphere creates a new Compressibility instance from InverseAtmosphere.
func (uf CompressibilityFactory) FromInverseAtmospheres(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseAtmosphere)
}

// FromInverseMillibar creates a new Compressibility instance from InverseMillibar.
func (uf CompressibilityFactory) FromInverseMillibars(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseMillibar)
}

// FromInverseBar creates a new Compressibility instance from InverseBar.
func (uf CompressibilityFactory) FromInverseBars(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInverseBar)
}

// FromInversePoundForcePerSquareInch creates a new Compressibility instance from InversePoundForcePerSquareInch.
func (uf CompressibilityFactory) FromInversePoundsForcePerSquareInch(value float64) (*Compressibility, error) {
	return newCompressibility(value, CompressibilityInversePoundForcePerSquareInch)
}




// newCompressibility creates a new Compressibility.
func newCompressibility(value float64, fromUnit CompressibilityUnits) (*Compressibility, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Compressibility{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Compressibility in InversePascal.
func (a *Compressibility) BaseValue() float64 {
	return a.value
}


// InversePascal returns the value in InversePascal.
func (a *Compressibility) InversePascals() float64 {
	if a.inverse_pascalsLazy != nil {
		return *a.inverse_pascalsLazy
	}
	inverse_pascals := a.convertFromBase(CompressibilityInversePascal)
	a.inverse_pascalsLazy = &inverse_pascals
	return inverse_pascals
}

// InverseKilopascal returns the value in InverseKilopascal.
func (a *Compressibility) InverseKilopascals() float64 {
	if a.inverse_kilopascalsLazy != nil {
		return *a.inverse_kilopascalsLazy
	}
	inverse_kilopascals := a.convertFromBase(CompressibilityInverseKilopascal)
	a.inverse_kilopascalsLazy = &inverse_kilopascals
	return inverse_kilopascals
}

// InverseMegapascal returns the value in InverseMegapascal.
func (a *Compressibility) InverseMegapascals() float64 {
	if a.inverse_megapascalsLazy != nil {
		return *a.inverse_megapascalsLazy
	}
	inverse_megapascals := a.convertFromBase(CompressibilityInverseMegapascal)
	a.inverse_megapascalsLazy = &inverse_megapascals
	return inverse_megapascals
}

// InverseAtmosphere returns the value in InverseAtmosphere.
func (a *Compressibility) InverseAtmospheres() float64 {
	if a.inverse_atmospheresLazy != nil {
		return *a.inverse_atmospheresLazy
	}
	inverse_atmospheres := a.convertFromBase(CompressibilityInverseAtmosphere)
	a.inverse_atmospheresLazy = &inverse_atmospheres
	return inverse_atmospheres
}

// InverseMillibar returns the value in InverseMillibar.
func (a *Compressibility) InverseMillibars() float64 {
	if a.inverse_millibarsLazy != nil {
		return *a.inverse_millibarsLazy
	}
	inverse_millibars := a.convertFromBase(CompressibilityInverseMillibar)
	a.inverse_millibarsLazy = &inverse_millibars
	return inverse_millibars
}

// InverseBar returns the value in InverseBar.
func (a *Compressibility) InverseBars() float64 {
	if a.inverse_barsLazy != nil {
		return *a.inverse_barsLazy
	}
	inverse_bars := a.convertFromBase(CompressibilityInverseBar)
	a.inverse_barsLazy = &inverse_bars
	return inverse_bars
}

// InversePoundForcePerSquareInch returns the value in InversePoundForcePerSquareInch.
func (a *Compressibility) InversePoundsForcePerSquareInch() float64 {
	if a.inverse_pounds_force_per_square_inchLazy != nil {
		return *a.inverse_pounds_force_per_square_inchLazy
	}
	inverse_pounds_force_per_square_inch := a.convertFromBase(CompressibilityInversePoundForcePerSquareInch)
	a.inverse_pounds_force_per_square_inchLazy = &inverse_pounds_force_per_square_inch
	return inverse_pounds_force_per_square_inch
}


// ToDto creates an CompressibilityDto representation.
func (a *Compressibility) ToDto(holdInUnit *CompressibilityUnits) CompressibilityDto {
	if holdInUnit == nil {
		defaultUnit := CompressibilityInversePascal // Default value
		holdInUnit = &defaultUnit
	}

	return CompressibilityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an CompressibilityDto representation.
func (a *Compressibility) ToDtoJSON(holdInUnit *CompressibilityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Compressibility to a specific unit value.
func (a *Compressibility) Convert(toUnit CompressibilityUnits) float64 {
	switch toUnit { 
    case CompressibilityInversePascal:
		return a.InversePascals()
    case CompressibilityInverseKilopascal:
		return a.InverseKilopascals()
    case CompressibilityInverseMegapascal:
		return a.InverseMegapascals()
    case CompressibilityInverseAtmosphere:
		return a.InverseAtmospheres()
    case CompressibilityInverseMillibar:
		return a.InverseMillibars()
    case CompressibilityInverseBar:
		return a.InverseBars()
    case CompressibilityInversePoundForcePerSquareInch:
		return a.InversePoundsForcePerSquareInch()
	default:
		return 0
	}
}

func (a *Compressibility) convertFromBase(toUnit CompressibilityUnits) float64 {
    value := a.value
	switch toUnit { 
	case CompressibilityInversePascal:
		return (value) 
	case CompressibilityInverseKilopascal:
		return (value / 1e3) 
	case CompressibilityInverseMegapascal:
		return (value / 1e6) 
	case CompressibilityInverseAtmosphere:
		return (value / 101325) 
	case CompressibilityInverseMillibar:
		return (value / 100) 
	case CompressibilityInverseBar:
		return (value / 1e5) 
	case CompressibilityInversePoundForcePerSquareInch:
		return (value / 6.894757293168361e3) 
	default:
		return math.NaN()
	}
}

func (a *Compressibility) convertToBase(value float64, fromUnit CompressibilityUnits) float64 {
	switch fromUnit { 
	case CompressibilityInversePascal:
		return (value) 
	case CompressibilityInverseKilopascal:
		return (value * 1e3) 
	case CompressibilityInverseMegapascal:
		return (value * 1e6) 
	case CompressibilityInverseAtmosphere:
		return (value * 101325) 
	case CompressibilityInverseMillibar:
		return (value * 100) 
	case CompressibilityInverseBar:
		return (value * 1e5) 
	case CompressibilityInversePoundForcePerSquareInch:
		return (value * 6.894757293168361e3) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Compressibility) String() string {
	return a.ToString(CompressibilityInversePascal, 2)
}

// ToString formats the Compressibility to string.
// fractionalDigits -1 for not mention
func (a *Compressibility) ToString(unit CompressibilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Compressibility) getUnitAbbreviation(unit CompressibilityUnits) string {
	switch unit { 
	case CompressibilityInversePascal:
		return "Pa⁻¹" 
	case CompressibilityInverseKilopascal:
		return "kPa⁻¹" 
	case CompressibilityInverseMegapascal:
		return "MPa⁻¹" 
	case CompressibilityInverseAtmosphere:
		return "atm⁻¹" 
	case CompressibilityInverseMillibar:
		return "mbar⁻¹" 
	case CompressibilityInverseBar:
		return "bar⁻¹" 
	case CompressibilityInversePoundForcePerSquareInch:
		return "psi⁻¹" 
	default:
		return ""
	}
}

// Check if the given Compressibility are equals to the current Compressibility
func (a *Compressibility) Equals(other *Compressibility) bool {
	return a.value == other.BaseValue()
}

// Check if the given Compressibility are equals to the current Compressibility
func (a *Compressibility) CompareTo(other *Compressibility) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Compressibility to the current Compressibility.
func (a *Compressibility) Add(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value + other.BaseValue()}
}

// Subtract the given Compressibility to the current Compressibility.
func (a *Compressibility) Subtract(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value - other.BaseValue()}
}

// Multiply the given Compressibility to the current Compressibility.
func (a *Compressibility) Multiply(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value * other.BaseValue()}
}

// Divide the given Compressibility to the current Compressibility.
func (a *Compressibility) Divide(other *Compressibility) *Compressibility {
	return &Compressibility{value: a.value / other.BaseValue()}
}