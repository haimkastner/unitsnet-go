// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MagneticFieldUnits defines various units of MagneticField.
type MagneticFieldUnits string

const (
    
        // 
        MagneticFieldTesla MagneticFieldUnits = "Tesla"
        // 
        MagneticFieldGauss MagneticFieldUnits = "Gauss"
        // 
        MagneticFieldNanotesla MagneticFieldUnits = "Nanotesla"
        // 
        MagneticFieldMicrotesla MagneticFieldUnits = "Microtesla"
        // 
        MagneticFieldMillitesla MagneticFieldUnits = "Millitesla"
        // 
        MagneticFieldMilligauss MagneticFieldUnits = "Milligauss"
)

var internalMagneticFieldUnitsMap = map[MagneticFieldUnits]bool{
	
	MagneticFieldTesla: true,
	MagneticFieldGauss: true,
	MagneticFieldNanotesla: true,
	MagneticFieldMicrotesla: true,
	MagneticFieldMillitesla: true,
	MagneticFieldMilligauss: true,
}

// MagneticFieldDto represents a MagneticField measurement with a numerical value and its corresponding unit.
type MagneticFieldDto struct {
    // Value is the numerical representation of the MagneticField.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the MagneticField, as defined in the MagneticFieldUnits enumeration.
	Unit  MagneticFieldUnits `json:"unit" validate:"required,oneof=Tesla Gauss Nanotesla Microtesla Millitesla Milligauss"`
}

// MagneticFieldDtoFactory groups methods for creating and serializing MagneticFieldDto objects.
type MagneticFieldDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MagneticFieldDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MagneticFieldDtoFactory) FromJSON(data []byte) (*MagneticFieldDto, error) {
	a := MagneticFieldDto{}

    // Parse JSON into MagneticFieldDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MagneticFieldDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MagneticFieldDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// MagneticField represents a measurement in a of MagneticField.
//
// A magnetic field is a force field that is created by moving electric charges (electric currents) and magnetic dipoles, and exerts a force on other nearby moving charges and magnetic dipoles.
type MagneticField struct {
	// value is the base measurement stored internally.
	value       float64
    
    teslasLazy *float64 
    gaussesLazy *float64 
    nanoteslasLazy *float64 
    microteslasLazy *float64 
    milliteslasLazy *float64 
    milligaussesLazy *float64 
}

// MagneticFieldFactory groups methods for creating MagneticField instances.
type MagneticFieldFactory struct{}

// CreateMagneticField creates a new MagneticField instance from the given value and unit.
func (uf MagneticFieldFactory) CreateMagneticField(value float64, unit MagneticFieldUnits) (*MagneticField, error) {
	return newMagneticField(value, unit)
}

// FromDto converts a MagneticFieldDto to a MagneticField instance.
func (uf MagneticFieldFactory) FromDto(dto MagneticFieldDto) (*MagneticField, error) {
	return newMagneticField(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MagneticField instance.
func (uf MagneticFieldFactory) FromDtoJSON(data []byte) (*MagneticField, error) {
	unitDto, err := MagneticFieldDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MagneticFieldDto from JSON: %w", err)
	}
	return MagneticFieldFactory{}.FromDto(*unitDto)
}


// FromTeslas creates a new MagneticField instance from a value in Teslas.
func (uf MagneticFieldFactory) FromTeslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldTesla)
}

// FromGausses creates a new MagneticField instance from a value in Gausses.
func (uf MagneticFieldFactory) FromGausses(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldGauss)
}

// FromNanoteslas creates a new MagneticField instance from a value in Nanoteslas.
func (uf MagneticFieldFactory) FromNanoteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldNanotesla)
}

// FromMicroteslas creates a new MagneticField instance from a value in Microteslas.
func (uf MagneticFieldFactory) FromMicroteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMicrotesla)
}

// FromMilliteslas creates a new MagneticField instance from a value in Milliteslas.
func (uf MagneticFieldFactory) FromMilliteslas(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMillitesla)
}

// FromMilligausses creates a new MagneticField instance from a value in Milligausses.
func (uf MagneticFieldFactory) FromMilligausses(value float64) (*MagneticField, error) {
	return newMagneticField(value, MagneticFieldMilligauss)
}


// newMagneticField creates a new MagneticField.
func newMagneticField(value float64, fromUnit MagneticFieldUnits) (*MagneticField, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMagneticFieldUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MagneticFieldUnits", fromUnit)
	}
	a := &MagneticField{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MagneticField in Tesla unit (the base/default quantity).
func (a *MagneticField) BaseValue() float64 {
	return a.value
}


// Teslas returns the MagneticField value in Teslas.
//
// 
func (a *MagneticField) Teslas() float64 {
	if a.teslasLazy != nil {
		return *a.teslasLazy
	}
	teslas := a.convertFromBase(MagneticFieldTesla)
	a.teslasLazy = &teslas
	return teslas
}

// Gausses returns the MagneticField value in Gausses.
//
// 
func (a *MagneticField) Gausses() float64 {
	if a.gaussesLazy != nil {
		return *a.gaussesLazy
	}
	gausses := a.convertFromBase(MagneticFieldGauss)
	a.gaussesLazy = &gausses
	return gausses
}

// Nanoteslas returns the MagneticField value in Nanoteslas.
//
// 
func (a *MagneticField) Nanoteslas() float64 {
	if a.nanoteslasLazy != nil {
		return *a.nanoteslasLazy
	}
	nanoteslas := a.convertFromBase(MagneticFieldNanotesla)
	a.nanoteslasLazy = &nanoteslas
	return nanoteslas
}

// Microteslas returns the MagneticField value in Microteslas.
//
// 
func (a *MagneticField) Microteslas() float64 {
	if a.microteslasLazy != nil {
		return *a.microteslasLazy
	}
	microteslas := a.convertFromBase(MagneticFieldMicrotesla)
	a.microteslasLazy = &microteslas
	return microteslas
}

// Milliteslas returns the MagneticField value in Milliteslas.
//
// 
func (a *MagneticField) Milliteslas() float64 {
	if a.milliteslasLazy != nil {
		return *a.milliteslasLazy
	}
	milliteslas := a.convertFromBase(MagneticFieldMillitesla)
	a.milliteslasLazy = &milliteslas
	return milliteslas
}

// Milligausses returns the MagneticField value in Milligausses.
//
// 
func (a *MagneticField) Milligausses() float64 {
	if a.milligaussesLazy != nil {
		return *a.milligaussesLazy
	}
	milligausses := a.convertFromBase(MagneticFieldMilligauss)
	a.milligaussesLazy = &milligausses
	return milligausses
}


// ToDto creates a MagneticFieldDto representation from the MagneticField instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Tesla by default.
func (a *MagneticField) ToDto(holdInUnit *MagneticFieldUnits) MagneticFieldDto {
	if holdInUnit == nil {
		defaultUnit := MagneticFieldTesla // Default value
		holdInUnit = &defaultUnit
	}

	return MagneticFieldDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MagneticField instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Tesla by default.
func (a *MagneticField) ToDtoJSON(holdInUnit *MagneticFieldUnits) ([]byte, error) {
	// Convert to MagneticFieldDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MagneticField to a specific unit value.
// The function uses the provided unit type (MagneticFieldUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MagneticField) Convert(toUnit MagneticFieldUnits) float64 {
	switch toUnit { 
    case MagneticFieldTesla:
		return a.Teslas()
    case MagneticFieldGauss:
		return a.Gausses()
    case MagneticFieldNanotesla:
		return a.Nanoteslas()
    case MagneticFieldMicrotesla:
		return a.Microteslas()
    case MagneticFieldMillitesla:
		return a.Milliteslas()
    case MagneticFieldMilligauss:
		return a.Milligausses()
	default:
		return math.NaN()
	}
}

func (a *MagneticField) convertFromBase(toUnit MagneticFieldUnits) float64 {
    value := a.value
	switch toUnit { 
	case MagneticFieldTesla:
		return (value) 
	case MagneticFieldGauss:
		return (value * 1e4) 
	case MagneticFieldNanotesla:
		return ((value) / 1e-09) 
	case MagneticFieldMicrotesla:
		return ((value) / 1e-06) 
	case MagneticFieldMillitesla:
		return ((value) / 0.001) 
	case MagneticFieldMilligauss:
		return ((value * 1e4) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *MagneticField) convertToBase(value float64, fromUnit MagneticFieldUnits) float64 {
	switch fromUnit { 
	case MagneticFieldTesla:
		return (value) 
	case MagneticFieldGauss:
		return (value / 1e4) 
	case MagneticFieldNanotesla:
		return ((value) * 1e-09) 
	case MagneticFieldMicrotesla:
		return ((value) * 1e-06) 
	case MagneticFieldMillitesla:
		return ((value) * 0.001) 
	case MagneticFieldMilligauss:
		return ((value / 1e4) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MagneticField in the default unit (Tesla),
// formatted to two decimal places.
func (a MagneticField) String() string {
	return a.ToString(MagneticFieldTesla, 2)
}

// ToString formats the MagneticField value as a string with the specified unit and fractional digits.
// It converts the MagneticField to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MagneticField value will be converted (e.g., Tesla))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MagneticField with the unit abbreviation.
func (a *MagneticField) ToString(unit MagneticFieldUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMagneticFieldAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMagneticFieldAbbreviation(unit))
}

// Equals checks if the given MagneticField is equal to the current MagneticField.
//
// Parameters:
//    other: The MagneticField to compare against.
//
// Returns:
//    bool: Returns true if both MagneticField are equal, false otherwise.
func (a *MagneticField) Equals(other *MagneticField) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MagneticField with another MagneticField.
// It returns -1 if the current MagneticField is less than the other MagneticField, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MagneticField to compare against.
//
// Returns:
//    int: -1 if the current MagneticField is less, 1 if greater, and 0 if equal.
func (a *MagneticField) CompareTo(other *MagneticField) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MagneticField to the current MagneticField and returns the result.
// The result is a new MagneticField instance with the sum of the values.
//
// Parameters:
//    other: The MagneticField to add to the current MagneticField.
//
// Returns:
//    *MagneticField: A new MagneticField instance representing the sum of both MagneticField.
func (a *MagneticField) Add(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MagneticField from the current MagneticField and returns the result.
// The result is a new MagneticField instance with the difference of the values.
//
// Parameters:
//    other: The MagneticField to subtract from the current MagneticField.
//
// Returns:
//    *MagneticField: A new MagneticField instance representing the difference of both MagneticField.
func (a *MagneticField) Subtract(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MagneticField by the given MagneticField and returns the result.
// The result is a new MagneticField instance with the product of the values.
//
// Parameters:
//    other: The MagneticField to multiply with the current MagneticField.
//
// Returns:
//    *MagneticField: A new MagneticField instance representing the product of both MagneticField.
func (a *MagneticField) Multiply(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value * other.BaseValue()}
}

// Divide divides the current MagneticField by the given MagneticField and returns the result.
// The result is a new MagneticField instance with the quotient of the values.
//
// Parameters:
//    other: The MagneticField to divide the current MagneticField by.
//
// Returns:
//    *MagneticField: A new MagneticField instance representing the quotient of both MagneticField.
func (a *MagneticField) Divide(other *MagneticField) *MagneticField {
	return &MagneticField{value: a.value / other.BaseValue()}
}

// GetMagneticFieldAbbreviation gets the unit abbreviation.
func GetMagneticFieldAbbreviation(unit MagneticFieldUnits) string {
	switch unit { 
	case MagneticFieldTesla:
		return "T" 
	case MagneticFieldGauss:
		return "G" 
	case MagneticFieldNanotesla:
		return "nT" 
	case MagneticFieldMicrotesla:
		return "μT" 
	case MagneticFieldMillitesla:
		return "mT" 
	case MagneticFieldMilligauss:
		return "mG" 
	default:
		return ""
	}
}