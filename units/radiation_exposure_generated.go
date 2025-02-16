// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RadiationExposureUnits defines various units of RadiationExposure.
type RadiationExposureUnits string

const (
    
        // 
        RadiationExposureCoulombPerKilogram RadiationExposureUnits = "CoulombPerKilogram"
        // 
        RadiationExposureRoentgen RadiationExposureUnits = "Roentgen"
        // 
        RadiationExposurePicocoulombPerKilogram RadiationExposureUnits = "PicocoulombPerKilogram"
        // 
        RadiationExposureNanocoulombPerKilogram RadiationExposureUnits = "NanocoulombPerKilogram"
        // 
        RadiationExposureMicrocoulombPerKilogram RadiationExposureUnits = "MicrocoulombPerKilogram"
        // 
        RadiationExposureMillicoulombPerKilogram RadiationExposureUnits = "MillicoulombPerKilogram"
        // 
        RadiationExposureMicroroentgen RadiationExposureUnits = "Microroentgen"
        // 
        RadiationExposureMilliroentgen RadiationExposureUnits = "Milliroentgen"
)

// RadiationExposureDto represents a RadiationExposure measurement with a numerical value and its corresponding unit.
type RadiationExposureDto struct {
    // Value is the numerical representation of the RadiationExposure.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RadiationExposure, as defined in the RadiationExposureUnits enumeration.
	Unit  RadiationExposureUnits `json:"unit"`
}

// RadiationExposureDtoFactory groups methods for creating and serializing RadiationExposureDto objects.
type RadiationExposureDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RadiationExposureDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RadiationExposureDtoFactory) FromJSON(data []byte) (*RadiationExposureDto, error) {
	a := RadiationExposureDto{}

    // Parse JSON into RadiationExposureDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a RadiationExposureDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RadiationExposureDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RadiationExposure represents a measurement in a of RadiationExposure.
//
// Radiation exposure is a measure of the ionization of air due to ionizing radiation from photons.
type RadiationExposure struct {
	// value is the base measurement stored internally.
	value       float64
    
    coulombs_per_kilogramLazy *float64 
    roentgensLazy *float64 
    picocoulombs_per_kilogramLazy *float64 
    nanocoulombs_per_kilogramLazy *float64 
    microcoulombs_per_kilogramLazy *float64 
    millicoulombs_per_kilogramLazy *float64 
    microroentgensLazy *float64 
    milliroentgensLazy *float64 
}

// RadiationExposureFactory groups methods for creating RadiationExposure instances.
type RadiationExposureFactory struct{}

// CreateRadiationExposure creates a new RadiationExposure instance from the given value and unit.
func (uf RadiationExposureFactory) CreateRadiationExposure(value float64, unit RadiationExposureUnits) (*RadiationExposure, error) {
	return newRadiationExposure(value, unit)
}

// FromDto converts a RadiationExposureDto to a RadiationExposure instance.
func (uf RadiationExposureFactory) FromDto(dto RadiationExposureDto) (*RadiationExposure, error) {
	return newRadiationExposure(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RadiationExposure instance.
func (uf RadiationExposureFactory) FromDtoJSON(data []byte) (*RadiationExposure, error) {
	unitDto, err := RadiationExposureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RadiationExposureDto from JSON: %w", err)
	}
	return RadiationExposureFactory{}.FromDto(*unitDto)
}


// FromCoulombsPerKilogram creates a new RadiationExposure instance from a value in CoulombsPerKilogram.
func (uf RadiationExposureFactory) FromCoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureCoulombPerKilogram)
}

// FromRoentgens creates a new RadiationExposure instance from a value in Roentgens.
func (uf RadiationExposureFactory) FromRoentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureRoentgen)
}

// FromPicocoulombsPerKilogram creates a new RadiationExposure instance from a value in PicocoulombsPerKilogram.
func (uf RadiationExposureFactory) FromPicocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposurePicocoulombPerKilogram)
}

// FromNanocoulombsPerKilogram creates a new RadiationExposure instance from a value in NanocoulombsPerKilogram.
func (uf RadiationExposureFactory) FromNanocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureNanocoulombPerKilogram)
}

// FromMicrocoulombsPerKilogram creates a new RadiationExposure instance from a value in MicrocoulombsPerKilogram.
func (uf RadiationExposureFactory) FromMicrocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMicrocoulombPerKilogram)
}

// FromMillicoulombsPerKilogram creates a new RadiationExposure instance from a value in MillicoulombsPerKilogram.
func (uf RadiationExposureFactory) FromMillicoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMillicoulombPerKilogram)
}

// FromMicroroentgens creates a new RadiationExposure instance from a value in Microroentgens.
func (uf RadiationExposureFactory) FromMicroroentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMicroroentgen)
}

// FromMilliroentgens creates a new RadiationExposure instance from a value in Milliroentgens.
func (uf RadiationExposureFactory) FromMilliroentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMilliroentgen)
}


// newRadiationExposure creates a new RadiationExposure.
func newRadiationExposure(value float64, fromUnit RadiationExposureUnits) (*RadiationExposure, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RadiationExposure{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RadiationExposure in CoulombPerKilogram unit (the base/default quantity).
func (a *RadiationExposure) BaseValue() float64 {
	return a.value
}


// CoulombsPerKilogram returns the RadiationExposure value in CoulombsPerKilogram.
//
// 
func (a *RadiationExposure) CoulombsPerKilogram() float64 {
	if a.coulombs_per_kilogramLazy != nil {
		return *a.coulombs_per_kilogramLazy
	}
	coulombs_per_kilogram := a.convertFromBase(RadiationExposureCoulombPerKilogram)
	a.coulombs_per_kilogramLazy = &coulombs_per_kilogram
	return coulombs_per_kilogram
}

// Roentgens returns the RadiationExposure value in Roentgens.
//
// 
func (a *RadiationExposure) Roentgens() float64 {
	if a.roentgensLazy != nil {
		return *a.roentgensLazy
	}
	roentgens := a.convertFromBase(RadiationExposureRoentgen)
	a.roentgensLazy = &roentgens
	return roentgens
}

// PicocoulombsPerKilogram returns the RadiationExposure value in PicocoulombsPerKilogram.
//
// 
func (a *RadiationExposure) PicocoulombsPerKilogram() float64 {
	if a.picocoulombs_per_kilogramLazy != nil {
		return *a.picocoulombs_per_kilogramLazy
	}
	picocoulombs_per_kilogram := a.convertFromBase(RadiationExposurePicocoulombPerKilogram)
	a.picocoulombs_per_kilogramLazy = &picocoulombs_per_kilogram
	return picocoulombs_per_kilogram
}

// NanocoulombsPerKilogram returns the RadiationExposure value in NanocoulombsPerKilogram.
//
// 
func (a *RadiationExposure) NanocoulombsPerKilogram() float64 {
	if a.nanocoulombs_per_kilogramLazy != nil {
		return *a.nanocoulombs_per_kilogramLazy
	}
	nanocoulombs_per_kilogram := a.convertFromBase(RadiationExposureNanocoulombPerKilogram)
	a.nanocoulombs_per_kilogramLazy = &nanocoulombs_per_kilogram
	return nanocoulombs_per_kilogram
}

// MicrocoulombsPerKilogram returns the RadiationExposure value in MicrocoulombsPerKilogram.
//
// 
func (a *RadiationExposure) MicrocoulombsPerKilogram() float64 {
	if a.microcoulombs_per_kilogramLazy != nil {
		return *a.microcoulombs_per_kilogramLazy
	}
	microcoulombs_per_kilogram := a.convertFromBase(RadiationExposureMicrocoulombPerKilogram)
	a.microcoulombs_per_kilogramLazy = &microcoulombs_per_kilogram
	return microcoulombs_per_kilogram
}

// MillicoulombsPerKilogram returns the RadiationExposure value in MillicoulombsPerKilogram.
//
// 
func (a *RadiationExposure) MillicoulombsPerKilogram() float64 {
	if a.millicoulombs_per_kilogramLazy != nil {
		return *a.millicoulombs_per_kilogramLazy
	}
	millicoulombs_per_kilogram := a.convertFromBase(RadiationExposureMillicoulombPerKilogram)
	a.millicoulombs_per_kilogramLazy = &millicoulombs_per_kilogram
	return millicoulombs_per_kilogram
}

// Microroentgens returns the RadiationExposure value in Microroentgens.
//
// 
func (a *RadiationExposure) Microroentgens() float64 {
	if a.microroentgensLazy != nil {
		return *a.microroentgensLazy
	}
	microroentgens := a.convertFromBase(RadiationExposureMicroroentgen)
	a.microroentgensLazy = &microroentgens
	return microroentgens
}

// Milliroentgens returns the RadiationExposure value in Milliroentgens.
//
// 
func (a *RadiationExposure) Milliroentgens() float64 {
	if a.milliroentgensLazy != nil {
		return *a.milliroentgensLazy
	}
	milliroentgens := a.convertFromBase(RadiationExposureMilliroentgen)
	a.milliroentgensLazy = &milliroentgens
	return milliroentgens
}


// ToDto creates a RadiationExposureDto representation from the RadiationExposure instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerKilogram by default.
func (a *RadiationExposure) ToDto(holdInUnit *RadiationExposureUnits) RadiationExposureDto {
	if holdInUnit == nil {
		defaultUnit := RadiationExposureCoulombPerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return RadiationExposureDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RadiationExposure instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CoulombPerKilogram by default.
func (a *RadiationExposure) ToDtoJSON(holdInUnit *RadiationExposureUnits) ([]byte, error) {
	// Convert to RadiationExposureDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RadiationExposure to a specific unit value.
// The function uses the provided unit type (RadiationExposureUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RadiationExposure) Convert(toUnit RadiationExposureUnits) float64 {
	switch toUnit { 
    case RadiationExposureCoulombPerKilogram:
		return a.CoulombsPerKilogram()
    case RadiationExposureRoentgen:
		return a.Roentgens()
    case RadiationExposurePicocoulombPerKilogram:
		return a.PicocoulombsPerKilogram()
    case RadiationExposureNanocoulombPerKilogram:
		return a.NanocoulombsPerKilogram()
    case RadiationExposureMicrocoulombPerKilogram:
		return a.MicrocoulombsPerKilogram()
    case RadiationExposureMillicoulombPerKilogram:
		return a.MillicoulombsPerKilogram()
    case RadiationExposureMicroroentgen:
		return a.Microroentgens()
    case RadiationExposureMilliroentgen:
		return a.Milliroentgens()
	default:
		return math.NaN()
	}
}

func (a *RadiationExposure) convertFromBase(toUnit RadiationExposureUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadiationExposureCoulombPerKilogram:
		return (value) 
	case RadiationExposureRoentgen:
		return (value / 2.58e-4) 
	case RadiationExposurePicocoulombPerKilogram:
		return ((value) / 1e-12) 
	case RadiationExposureNanocoulombPerKilogram:
		return ((value) / 1e-09) 
	case RadiationExposureMicrocoulombPerKilogram:
		return ((value) / 1e-06) 
	case RadiationExposureMillicoulombPerKilogram:
		return ((value) / 0.001) 
	case RadiationExposureMicroroentgen:
		return ((value / 2.58e-4) / 1e-06) 
	case RadiationExposureMilliroentgen:
		return ((value / 2.58e-4) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RadiationExposure) convertToBase(value float64, fromUnit RadiationExposureUnits) float64 {
	switch fromUnit { 
	case RadiationExposureCoulombPerKilogram:
		return (value) 
	case RadiationExposureRoentgen:
		return (value * 2.58e-4) 
	case RadiationExposurePicocoulombPerKilogram:
		return ((value) * 1e-12) 
	case RadiationExposureNanocoulombPerKilogram:
		return ((value) * 1e-09) 
	case RadiationExposureMicrocoulombPerKilogram:
		return ((value) * 1e-06) 
	case RadiationExposureMillicoulombPerKilogram:
		return ((value) * 0.001) 
	case RadiationExposureMicroroentgen:
		return ((value * 2.58e-4) * 1e-06) 
	case RadiationExposureMilliroentgen:
		return ((value * 2.58e-4) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RadiationExposure in the default unit (CoulombPerKilogram),
// formatted to two decimal places.
func (a RadiationExposure) String() string {
	return a.ToString(RadiationExposureCoulombPerKilogram, 2)
}

// ToString formats the RadiationExposure value as a string with the specified unit and fractional digits.
// It converts the RadiationExposure to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RadiationExposure value will be converted (e.g., CoulombPerKilogram))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RadiationExposure with the unit abbreviation.
func (a *RadiationExposure) ToString(unit RadiationExposureUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRadiationExposureAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRadiationExposureAbbreviation(unit))
}

// Equals checks if the given RadiationExposure is equal to the current RadiationExposure.
//
// Parameters:
//    other: The RadiationExposure to compare against.
//
// Returns:
//    bool: Returns true if both RadiationExposure are equal, false otherwise.
func (a *RadiationExposure) Equals(other *RadiationExposure) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RadiationExposure with another RadiationExposure.
// It returns -1 if the current RadiationExposure is less than the other RadiationExposure, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RadiationExposure to compare against.
//
// Returns:
//    int: -1 if the current RadiationExposure is less, 1 if greater, and 0 if equal.
func (a *RadiationExposure) CompareTo(other *RadiationExposure) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RadiationExposure to the current RadiationExposure and returns the result.
// The result is a new RadiationExposure instance with the sum of the values.
//
// Parameters:
//    other: The RadiationExposure to add to the current RadiationExposure.
//
// Returns:
//    *RadiationExposure: A new RadiationExposure instance representing the sum of both RadiationExposure.
func (a *RadiationExposure) Add(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RadiationExposure from the current RadiationExposure and returns the result.
// The result is a new RadiationExposure instance with the difference of the values.
//
// Parameters:
//    other: The RadiationExposure to subtract from the current RadiationExposure.
//
// Returns:
//    *RadiationExposure: A new RadiationExposure instance representing the difference of both RadiationExposure.
func (a *RadiationExposure) Subtract(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RadiationExposure by the given RadiationExposure and returns the result.
// The result is a new RadiationExposure instance with the product of the values.
//
// Parameters:
//    other: The RadiationExposure to multiply with the current RadiationExposure.
//
// Returns:
//    *RadiationExposure: A new RadiationExposure instance representing the product of both RadiationExposure.
func (a *RadiationExposure) Multiply(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value * other.BaseValue()}
}

// Divide divides the current RadiationExposure by the given RadiationExposure and returns the result.
// The result is a new RadiationExposure instance with the quotient of the values.
//
// Parameters:
//    other: The RadiationExposure to divide the current RadiationExposure by.
//
// Returns:
//    *RadiationExposure: A new RadiationExposure instance representing the quotient of both RadiationExposure.
func (a *RadiationExposure) Divide(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value / other.BaseValue()}
}

// GetRadiationExposureAbbreviation gets the unit abbreviation.
func GetRadiationExposureAbbreviation(unit RadiationExposureUnits) string {
	switch unit { 
	case RadiationExposureCoulombPerKilogram:
		return "C/kg" 
	case RadiationExposureRoentgen:
		return "R" 
	case RadiationExposurePicocoulombPerKilogram:
		return "pC/kg" 
	case RadiationExposureNanocoulombPerKilogram:
		return "nC/kg" 
	case RadiationExposureMicrocoulombPerKilogram:
		return "μC/kg" 
	case RadiationExposureMillicoulombPerKilogram:
		return "mC/kg" 
	case RadiationExposureMicroroentgen:
		return "μR" 
	case RadiationExposureMilliroentgen:
		return "mR" 
	default:
		return ""
	}
}