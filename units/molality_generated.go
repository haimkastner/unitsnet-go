// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolalityUnits defines various units of Molality.
type MolalityUnits string

const (
    
        // 
        MolalityMolePerKilogram MolalityUnits = "MolePerKilogram"
        // 
        MolalityMolePerGram MolalityUnits = "MolePerGram"
        // 
        MolalityMillimolePerKilogram MolalityUnits = "MillimolePerKilogram"
)

var internalMolalityUnitsMap = map[MolalityUnits]bool{
	
	MolalityMolePerKilogram: true,
	MolalityMolePerGram: true,
	MolalityMillimolePerKilogram: true,
}

// MolalityDto represents a Molality measurement with a numerical value and its corresponding unit.
type MolalityDto struct {
    // Value is the numerical representation of the Molality.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Molality, as defined in the MolalityUnits enumeration.
	Unit  MolalityUnits `json:"unit" validate:"required,oneof=MolePerKilogram,MolePerGram,MillimolePerKilogram"`
}

// MolalityDtoFactory groups methods for creating and serializing MolalityDto objects.
type MolalityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolalityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolalityDtoFactory) FromJSON(data []byte) (*MolalityDto, error) {
	a := MolalityDto{}

    // Parse JSON into MolalityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MolalityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolalityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Molality represents a measurement in a of Molality.
//
// Molality is a measure of the amount of solute in a solution relative to a given mass of solvent.
type Molality struct {
	// value is the base measurement stored internally.
	value       float64
    
    moles_per_kilogramLazy *float64 
    moles_per_gramLazy *float64 
    millimoles_per_kilogramLazy *float64 
}

// MolalityFactory groups methods for creating Molality instances.
type MolalityFactory struct{}

// CreateMolality creates a new Molality instance from the given value and unit.
func (uf MolalityFactory) CreateMolality(value float64, unit MolalityUnits) (*Molality, error) {
	return newMolality(value, unit)
}

// FromDto converts a MolalityDto to a Molality instance.
func (uf MolalityFactory) FromDto(dto MolalityDto) (*Molality, error) {
	return newMolality(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Molality instance.
func (uf MolalityFactory) FromDtoJSON(data []byte) (*Molality, error) {
	unitDto, err := MolalityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolalityDto from JSON: %w", err)
	}
	return MolalityFactory{}.FromDto(*unitDto)
}


// FromMolesPerKilogram creates a new Molality instance from a value in MolesPerKilogram.
func (uf MolalityFactory) FromMolesPerKilogram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMolePerKilogram)
}

// FromMolesPerGram creates a new Molality instance from a value in MolesPerGram.
func (uf MolalityFactory) FromMolesPerGram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMolePerGram)
}

// FromMillimolesPerKilogram creates a new Molality instance from a value in MillimolesPerKilogram.
func (uf MolalityFactory) FromMillimolesPerKilogram(value float64) (*Molality, error) {
	return newMolality(value, MolalityMillimolePerKilogram)
}


// newMolality creates a new Molality.
func newMolality(value float64, fromUnit MolalityUnits) (*Molality, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMolalityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MolalityUnits", fromUnit)
	}
	a := &Molality{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Molality in MolePerKilogram unit (the base/default quantity).
func (a *Molality) BaseValue() float64 {
	return a.value
}


// MolesPerKilogram returns the Molality value in MolesPerKilogram.
//
// 
func (a *Molality) MolesPerKilogram() float64 {
	if a.moles_per_kilogramLazy != nil {
		return *a.moles_per_kilogramLazy
	}
	moles_per_kilogram := a.convertFromBase(MolalityMolePerKilogram)
	a.moles_per_kilogramLazy = &moles_per_kilogram
	return moles_per_kilogram
}

// MolesPerGram returns the Molality value in MolesPerGram.
//
// 
func (a *Molality) MolesPerGram() float64 {
	if a.moles_per_gramLazy != nil {
		return *a.moles_per_gramLazy
	}
	moles_per_gram := a.convertFromBase(MolalityMolePerGram)
	a.moles_per_gramLazy = &moles_per_gram
	return moles_per_gram
}

// MillimolesPerKilogram returns the Molality value in MillimolesPerKilogram.
//
// 
func (a *Molality) MillimolesPerKilogram() float64 {
	if a.millimoles_per_kilogramLazy != nil {
		return *a.millimoles_per_kilogramLazy
	}
	millimoles_per_kilogram := a.convertFromBase(MolalityMillimolePerKilogram)
	a.millimoles_per_kilogramLazy = &millimoles_per_kilogram
	return millimoles_per_kilogram
}


// ToDto creates a MolalityDto representation from the Molality instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerKilogram by default.
func (a *Molality) ToDto(holdInUnit *MolalityUnits) MolalityDto {
	if holdInUnit == nil {
		defaultUnit := MolalityMolePerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return MolalityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Molality instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerKilogram by default.
func (a *Molality) ToDtoJSON(holdInUnit *MolalityUnits) ([]byte, error) {
	// Convert to MolalityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Molality to a specific unit value.
// The function uses the provided unit type (MolalityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Molality) Convert(toUnit MolalityUnits) float64 {
	switch toUnit { 
    case MolalityMolePerKilogram:
		return a.MolesPerKilogram()
    case MolalityMolePerGram:
		return a.MolesPerGram()
    case MolalityMillimolePerKilogram:
		return a.MillimolesPerKilogram()
	default:
		return math.NaN()
	}
}

func (a *Molality) convertFromBase(toUnit MolalityUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolalityMolePerKilogram:
		return (value) 
	case MolalityMolePerGram:
		return (value * 1e-3) 
	case MolalityMillimolePerKilogram:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *Molality) convertToBase(value float64, fromUnit MolalityUnits) float64 {
	switch fromUnit { 
	case MolalityMolePerKilogram:
		return (value) 
	case MolalityMolePerGram:
		return (value / 1e-3) 
	case MolalityMillimolePerKilogram:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Molality in the default unit (MolePerKilogram),
// formatted to two decimal places.
func (a Molality) String() string {
	return a.ToString(MolalityMolePerKilogram, 2)
}

// ToString formats the Molality value as a string with the specified unit and fractional digits.
// It converts the Molality to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Molality value will be converted (e.g., MolePerKilogram))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Molality with the unit abbreviation.
func (a *Molality) ToString(unit MolalityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMolalityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMolalityAbbreviation(unit))
}

// Equals checks if the given Molality is equal to the current Molality.
//
// Parameters:
//    other: The Molality to compare against.
//
// Returns:
//    bool: Returns true if both Molality are equal, false otherwise.
func (a *Molality) Equals(other *Molality) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Molality with another Molality.
// It returns -1 if the current Molality is less than the other Molality, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Molality to compare against.
//
// Returns:
//    int: -1 if the current Molality is less, 1 if greater, and 0 if equal.
func (a *Molality) CompareTo(other *Molality) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Molality to the current Molality and returns the result.
// The result is a new Molality instance with the sum of the values.
//
// Parameters:
//    other: The Molality to add to the current Molality.
//
// Returns:
//    *Molality: A new Molality instance representing the sum of both Molality.
func (a *Molality) Add(other *Molality) *Molality {
	return &Molality{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Molality from the current Molality and returns the result.
// The result is a new Molality instance with the difference of the values.
//
// Parameters:
//    other: The Molality to subtract from the current Molality.
//
// Returns:
//    *Molality: A new Molality instance representing the difference of both Molality.
func (a *Molality) Subtract(other *Molality) *Molality {
	return &Molality{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Molality by the given Molality and returns the result.
// The result is a new Molality instance with the product of the values.
//
// Parameters:
//    other: The Molality to multiply with the current Molality.
//
// Returns:
//    *Molality: A new Molality instance representing the product of both Molality.
func (a *Molality) Multiply(other *Molality) *Molality {
	return &Molality{value: a.value * other.BaseValue()}
}

// Divide divides the current Molality by the given Molality and returns the result.
// The result is a new Molality instance with the quotient of the values.
//
// Parameters:
//    other: The Molality to divide the current Molality by.
//
// Returns:
//    *Molality: A new Molality instance representing the quotient of both Molality.
func (a *Molality) Divide(other *Molality) *Molality {
	return &Molality{value: a.value / other.BaseValue()}
}

// GetMolalityAbbreviation gets the unit abbreviation.
func GetMolalityAbbreviation(unit MolalityUnits) string {
	switch unit { 
	case MolalityMolePerKilogram:
		return "mol/kg" 
	case MolalityMolePerGram:
		return "mol/g" 
	case MolalityMillimolePerKilogram:
		return "mmol/kg" 
	default:
		return ""
	}
}