// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// RadiationEquivalentDoseUnits defines various units of RadiationEquivalentDose.
type RadiationEquivalentDoseUnits string

const (
    
        // The sievert is a unit in the International System of Units (SI) intended to represent the stochastic health risk of ionizing radiation, which is defined as the probability of causing radiation-induced cancer and genetic damage.
        RadiationEquivalentDoseSievert RadiationEquivalentDoseUnits = "Sievert"
        // 
        RadiationEquivalentDoseRoentgenEquivalentMan RadiationEquivalentDoseUnits = "RoentgenEquivalentMan"
        // 
        RadiationEquivalentDoseNanosievert RadiationEquivalentDoseUnits = "Nanosievert"
        // 
        RadiationEquivalentDoseMicrosievert RadiationEquivalentDoseUnits = "Microsievert"
        // 
        RadiationEquivalentDoseMillisievert RadiationEquivalentDoseUnits = "Millisievert"
        // 
        RadiationEquivalentDoseMilliroentgenEquivalentMan RadiationEquivalentDoseUnits = "MilliroentgenEquivalentMan"
)

// RadiationEquivalentDoseDto represents a RadiationEquivalentDose measurement with a numerical value and its corresponding unit.
type RadiationEquivalentDoseDto struct {
    // Value is the numerical representation of the RadiationEquivalentDose.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the RadiationEquivalentDose, as defined in the RadiationEquivalentDoseUnits enumeration.
	Unit  RadiationEquivalentDoseUnits `json:"unit"`
}

// RadiationEquivalentDoseDtoFactory groups methods for creating and serializing RadiationEquivalentDoseDto objects.
type RadiationEquivalentDoseDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a RadiationEquivalentDoseDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf RadiationEquivalentDoseDtoFactory) FromJSON(data []byte) (*RadiationEquivalentDoseDto, error) {
	a := RadiationEquivalentDoseDto{}

    // Parse JSON into RadiationEquivalentDoseDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a RadiationEquivalentDoseDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a RadiationEquivalentDoseDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// RadiationEquivalentDose represents a measurement in a of RadiationEquivalentDose.
//
// Equivalent dose is a dose quantity representing the stochastic health effects of low levels of ionizing radiation on the human body which represents the probability of radiation-induced cancer and genetic damage.
type RadiationEquivalentDose struct {
	// value is the base measurement stored internally.
	value       float64
    
    sievertsLazy *float64 
    roentgens_equivalent_manLazy *float64 
    nanosievertsLazy *float64 
    microsievertsLazy *float64 
    millisievertsLazy *float64 
    milliroentgens_equivalent_manLazy *float64 
}

// RadiationEquivalentDoseFactory groups methods for creating RadiationEquivalentDose instances.
type RadiationEquivalentDoseFactory struct{}

// CreateRadiationEquivalentDose creates a new RadiationEquivalentDose instance from the given value and unit.
func (uf RadiationEquivalentDoseFactory) CreateRadiationEquivalentDose(value float64, unit RadiationEquivalentDoseUnits) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, unit)
}

// FromDto converts a RadiationEquivalentDoseDto to a RadiationEquivalentDose instance.
func (uf RadiationEquivalentDoseFactory) FromDto(dto RadiationEquivalentDoseDto) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a RadiationEquivalentDose instance.
func (uf RadiationEquivalentDoseFactory) FromDtoJSON(data []byte) (*RadiationEquivalentDose, error) {
	unitDto, err := RadiationEquivalentDoseDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RadiationEquivalentDoseDto from JSON: %w", err)
	}
	return RadiationEquivalentDoseFactory{}.FromDto(*unitDto)
}


// FromSieverts creates a new RadiationEquivalentDose instance from a value in Sieverts.
func (uf RadiationEquivalentDoseFactory) FromSieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseSievert)
}

// FromRoentgensEquivalentMan creates a new RadiationEquivalentDose instance from a value in RoentgensEquivalentMan.
func (uf RadiationEquivalentDoseFactory) FromRoentgensEquivalentMan(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseRoentgenEquivalentMan)
}

// FromNanosieverts creates a new RadiationEquivalentDose instance from a value in Nanosieverts.
func (uf RadiationEquivalentDoseFactory) FromNanosieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseNanosievert)
}

// FromMicrosieverts creates a new RadiationEquivalentDose instance from a value in Microsieverts.
func (uf RadiationEquivalentDoseFactory) FromMicrosieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMicrosievert)
}

// FromMillisieverts creates a new RadiationEquivalentDose instance from a value in Millisieverts.
func (uf RadiationEquivalentDoseFactory) FromMillisieverts(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMillisievert)
}

// FromMilliroentgensEquivalentMan creates a new RadiationEquivalentDose instance from a value in MilliroentgensEquivalentMan.
func (uf RadiationEquivalentDoseFactory) FromMilliroentgensEquivalentMan(value float64) (*RadiationEquivalentDose, error) {
	return newRadiationEquivalentDose(value, RadiationEquivalentDoseMilliroentgenEquivalentMan)
}


// newRadiationEquivalentDose creates a new RadiationEquivalentDose.
func newRadiationEquivalentDose(value float64, fromUnit RadiationEquivalentDoseUnits) (*RadiationEquivalentDose, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RadiationEquivalentDose{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RadiationEquivalentDose in Sievert unit (the base/default quantity).
func (a *RadiationEquivalentDose) BaseValue() float64 {
	return a.value
}


// Sieverts returns the RadiationEquivalentDose value in Sieverts.
//
// The sievert is a unit in the International System of Units (SI) intended to represent the stochastic health risk of ionizing radiation, which is defined as the probability of causing radiation-induced cancer and genetic damage.
func (a *RadiationEquivalentDose) Sieverts() float64 {
	if a.sievertsLazy != nil {
		return *a.sievertsLazy
	}
	sieverts := a.convertFromBase(RadiationEquivalentDoseSievert)
	a.sievertsLazy = &sieverts
	return sieverts
}

// RoentgensEquivalentMan returns the RadiationEquivalentDose value in RoentgensEquivalentMan.
//
// 
func (a *RadiationEquivalentDose) RoentgensEquivalentMan() float64 {
	if a.roentgens_equivalent_manLazy != nil {
		return *a.roentgens_equivalent_manLazy
	}
	roentgens_equivalent_man := a.convertFromBase(RadiationEquivalentDoseRoentgenEquivalentMan)
	a.roentgens_equivalent_manLazy = &roentgens_equivalent_man
	return roentgens_equivalent_man
}

// Nanosieverts returns the RadiationEquivalentDose value in Nanosieverts.
//
// 
func (a *RadiationEquivalentDose) Nanosieverts() float64 {
	if a.nanosievertsLazy != nil {
		return *a.nanosievertsLazy
	}
	nanosieverts := a.convertFromBase(RadiationEquivalentDoseNanosievert)
	a.nanosievertsLazy = &nanosieverts
	return nanosieverts
}

// Microsieverts returns the RadiationEquivalentDose value in Microsieverts.
//
// 
func (a *RadiationEquivalentDose) Microsieverts() float64 {
	if a.microsievertsLazy != nil {
		return *a.microsievertsLazy
	}
	microsieverts := a.convertFromBase(RadiationEquivalentDoseMicrosievert)
	a.microsievertsLazy = &microsieverts
	return microsieverts
}

// Millisieverts returns the RadiationEquivalentDose value in Millisieverts.
//
// 
func (a *RadiationEquivalentDose) Millisieverts() float64 {
	if a.millisievertsLazy != nil {
		return *a.millisievertsLazy
	}
	millisieverts := a.convertFromBase(RadiationEquivalentDoseMillisievert)
	a.millisievertsLazy = &millisieverts
	return millisieverts
}

// MilliroentgensEquivalentMan returns the RadiationEquivalentDose value in MilliroentgensEquivalentMan.
//
// 
func (a *RadiationEquivalentDose) MilliroentgensEquivalentMan() float64 {
	if a.milliroentgens_equivalent_manLazy != nil {
		return *a.milliroentgens_equivalent_manLazy
	}
	milliroentgens_equivalent_man := a.convertFromBase(RadiationEquivalentDoseMilliroentgenEquivalentMan)
	a.milliroentgens_equivalent_manLazy = &milliroentgens_equivalent_man
	return milliroentgens_equivalent_man
}


// ToDto creates a RadiationEquivalentDoseDto representation from the RadiationEquivalentDose instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Sievert by default.
func (a *RadiationEquivalentDose) ToDto(holdInUnit *RadiationEquivalentDoseUnits) RadiationEquivalentDoseDto {
	if holdInUnit == nil {
		defaultUnit := RadiationEquivalentDoseSievert // Default value
		holdInUnit = &defaultUnit
	}

	return RadiationEquivalentDoseDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the RadiationEquivalentDose instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Sievert by default.
func (a *RadiationEquivalentDose) ToDtoJSON(holdInUnit *RadiationEquivalentDoseUnits) ([]byte, error) {
	// Convert to RadiationEquivalentDoseDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a RadiationEquivalentDose to a specific unit value.
// The function uses the provided unit type (RadiationEquivalentDoseUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *RadiationEquivalentDose) Convert(toUnit RadiationEquivalentDoseUnits) float64 {
	switch toUnit { 
    case RadiationEquivalentDoseSievert:
		return a.Sieverts()
    case RadiationEquivalentDoseRoentgenEquivalentMan:
		return a.RoentgensEquivalentMan()
    case RadiationEquivalentDoseNanosievert:
		return a.Nanosieverts()
    case RadiationEquivalentDoseMicrosievert:
		return a.Microsieverts()
    case RadiationEquivalentDoseMillisievert:
		return a.Millisieverts()
    case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return a.MilliroentgensEquivalentMan()
	default:
		return math.NaN()
	}
}

func (a *RadiationEquivalentDose) convertFromBase(toUnit RadiationEquivalentDoseUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadiationEquivalentDoseSievert:
		return (value) 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return (value * 100) 
	case RadiationEquivalentDoseNanosievert:
		return ((value) / 1e-09) 
	case RadiationEquivalentDoseMicrosievert:
		return ((value) / 1e-06) 
	case RadiationEquivalentDoseMillisievert:
		return ((value) / 0.001) 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return ((value * 100) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RadiationEquivalentDose) convertToBase(value float64, fromUnit RadiationEquivalentDoseUnits) float64 {
	switch fromUnit { 
	case RadiationEquivalentDoseSievert:
		return (value) 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return (value / 100) 
	case RadiationEquivalentDoseNanosievert:
		return ((value) * 1e-09) 
	case RadiationEquivalentDoseMicrosievert:
		return ((value) * 1e-06) 
	case RadiationEquivalentDoseMillisievert:
		return ((value) * 0.001) 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return ((value / 100) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the RadiationEquivalentDose in the default unit (Sievert),
// formatted to two decimal places.
func (a RadiationEquivalentDose) String() string {
	return a.ToString(RadiationEquivalentDoseSievert, 2)
}

// ToString formats the RadiationEquivalentDose value as a string with the specified unit and fractional digits.
// It converts the RadiationEquivalentDose to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the RadiationEquivalentDose value will be converted (e.g., Sievert))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the RadiationEquivalentDose with the unit abbreviation.
func (a *RadiationEquivalentDose) ToString(unit RadiationEquivalentDoseUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetRadiationEquivalentDoseAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetRadiationEquivalentDoseAbbreviation(unit))
}

// Equals checks if the given RadiationEquivalentDose is equal to the current RadiationEquivalentDose.
//
// Parameters:
//    other: The RadiationEquivalentDose to compare against.
//
// Returns:
//    bool: Returns true if both RadiationEquivalentDose are equal, false otherwise.
func (a *RadiationEquivalentDose) Equals(other *RadiationEquivalentDose) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current RadiationEquivalentDose with another RadiationEquivalentDose.
// It returns -1 if the current RadiationEquivalentDose is less than the other RadiationEquivalentDose, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The RadiationEquivalentDose to compare against.
//
// Returns:
//    int: -1 if the current RadiationEquivalentDose is less, 1 if greater, and 0 if equal.
func (a *RadiationEquivalentDose) CompareTo(other *RadiationEquivalentDose) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given RadiationEquivalentDose to the current RadiationEquivalentDose and returns the result.
// The result is a new RadiationEquivalentDose instance with the sum of the values.
//
// Parameters:
//    other: The RadiationEquivalentDose to add to the current RadiationEquivalentDose.
//
// Returns:
//    *RadiationEquivalentDose: A new RadiationEquivalentDose instance representing the sum of both RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Add(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given RadiationEquivalentDose from the current RadiationEquivalentDose and returns the result.
// The result is a new RadiationEquivalentDose instance with the difference of the values.
//
// Parameters:
//    other: The RadiationEquivalentDose to subtract from the current RadiationEquivalentDose.
//
// Returns:
//    *RadiationEquivalentDose: A new RadiationEquivalentDose instance representing the difference of both RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Subtract(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current RadiationEquivalentDose by the given RadiationEquivalentDose and returns the result.
// The result is a new RadiationEquivalentDose instance with the product of the values.
//
// Parameters:
//    other: The RadiationEquivalentDose to multiply with the current RadiationEquivalentDose.
//
// Returns:
//    *RadiationEquivalentDose: A new RadiationEquivalentDose instance representing the product of both RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Multiply(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value * other.BaseValue()}
}

// Divide divides the current RadiationEquivalentDose by the given RadiationEquivalentDose and returns the result.
// The result is a new RadiationEquivalentDose instance with the quotient of the values.
//
// Parameters:
//    other: The RadiationEquivalentDose to divide the current RadiationEquivalentDose by.
//
// Returns:
//    *RadiationEquivalentDose: A new RadiationEquivalentDose instance representing the quotient of both RadiationEquivalentDose.
func (a *RadiationEquivalentDose) Divide(other *RadiationEquivalentDose) *RadiationEquivalentDose {
	return &RadiationEquivalentDose{value: a.value / other.BaseValue()}
}

// GetRadiationEquivalentDoseAbbreviation gets the unit abbreviation.
func GetRadiationEquivalentDoseAbbreviation(unit RadiationEquivalentDoseUnits) string {
	switch unit { 
	case RadiationEquivalentDoseSievert:
		return "Sv" 
	case RadiationEquivalentDoseRoentgenEquivalentMan:
		return "rem" 
	case RadiationEquivalentDoseNanosievert:
		return "nSv" 
	case RadiationEquivalentDoseMicrosievert:
		return "μSv" 
	case RadiationEquivalentDoseMillisievert:
		return "mSv" 
	case RadiationEquivalentDoseMilliroentgenEquivalentMan:
		return "mrem" 
	default:
		return ""
	}
}