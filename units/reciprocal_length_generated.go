// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ReciprocalLengthUnits defines various units of ReciprocalLength.
type ReciprocalLengthUnits string

const (
    
        // 
        ReciprocalLengthInverseMeter ReciprocalLengthUnits = "InverseMeter"
        // 
        ReciprocalLengthInverseCentimeter ReciprocalLengthUnits = "InverseCentimeter"
        // 
        ReciprocalLengthInverseMillimeter ReciprocalLengthUnits = "InverseMillimeter"
        // 
        ReciprocalLengthInverseMile ReciprocalLengthUnits = "InverseMile"
        // 
        ReciprocalLengthInverseYard ReciprocalLengthUnits = "InverseYard"
        // 
        ReciprocalLengthInverseFoot ReciprocalLengthUnits = "InverseFoot"
        // 
        ReciprocalLengthInverseUsSurveyFoot ReciprocalLengthUnits = "InverseUsSurveyFoot"
        // 
        ReciprocalLengthInverseInch ReciprocalLengthUnits = "InverseInch"
        // 
        ReciprocalLengthInverseMil ReciprocalLengthUnits = "InverseMil"
        // 
        ReciprocalLengthInverseMicroinch ReciprocalLengthUnits = "InverseMicroinch"
)

var internalReciprocalLengthUnitsMap = map[ReciprocalLengthUnits]bool{
	
	ReciprocalLengthInverseMeter: true,
	ReciprocalLengthInverseCentimeter: true,
	ReciprocalLengthInverseMillimeter: true,
	ReciprocalLengthInverseMile: true,
	ReciprocalLengthInverseYard: true,
	ReciprocalLengthInverseFoot: true,
	ReciprocalLengthInverseUsSurveyFoot: true,
	ReciprocalLengthInverseInch: true,
	ReciprocalLengthInverseMil: true,
	ReciprocalLengthInverseMicroinch: true,
}

// ReciprocalLengthDto represents a ReciprocalLength measurement with a numerical value and its corresponding unit.
type ReciprocalLengthDto struct {
    // Value is the numerical representation of the ReciprocalLength.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ReciprocalLength, as defined in the ReciprocalLengthUnits enumeration.
	Unit  ReciprocalLengthUnits `json:"unit" validate:"required,oneof=InverseMeter,InverseCentimeter,InverseMillimeter,InverseMile,InverseYard,InverseFoot,InverseUsSurveyFoot,InverseInch,InverseMil,InverseMicroinch"`
}

// ReciprocalLengthDtoFactory groups methods for creating and serializing ReciprocalLengthDto objects.
type ReciprocalLengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ReciprocalLengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ReciprocalLengthDtoFactory) FromJSON(data []byte) (*ReciprocalLengthDto, error) {
	a := ReciprocalLengthDto{}

    // Parse JSON into ReciprocalLengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ReciprocalLengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ReciprocalLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ReciprocalLength represents a measurement in a of ReciprocalLength.
//
// Reciprocal (Inverse) Length is used in various fields of science and mathematics. It is defined as the inverse value of a length unit.
type ReciprocalLength struct {
	// value is the base measurement stored internally.
	value       float64
    
    inverse_metersLazy *float64 
    inverse_centimetersLazy *float64 
    inverse_millimetersLazy *float64 
    inverse_milesLazy *float64 
    inverse_yardsLazy *float64 
    inverse_feetLazy *float64 
    inverse_us_survey_feetLazy *float64 
    inverse_inchesLazy *float64 
    inverse_milsLazy *float64 
    inverse_microinchesLazy *float64 
}

// ReciprocalLengthFactory groups methods for creating ReciprocalLength instances.
type ReciprocalLengthFactory struct{}

// CreateReciprocalLength creates a new ReciprocalLength instance from the given value and unit.
func (uf ReciprocalLengthFactory) CreateReciprocalLength(value float64, unit ReciprocalLengthUnits) (*ReciprocalLength, error) {
	return newReciprocalLength(value, unit)
}

// FromDto converts a ReciprocalLengthDto to a ReciprocalLength instance.
func (uf ReciprocalLengthFactory) FromDto(dto ReciprocalLengthDto) (*ReciprocalLength, error) {
	return newReciprocalLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ReciprocalLength instance.
func (uf ReciprocalLengthFactory) FromDtoJSON(data []byte) (*ReciprocalLength, error) {
	unitDto, err := ReciprocalLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ReciprocalLengthDto from JSON: %w", err)
	}
	return ReciprocalLengthFactory{}.FromDto(*unitDto)
}


// FromInverseMeters creates a new ReciprocalLength instance from a value in InverseMeters.
func (uf ReciprocalLengthFactory) FromInverseMeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMeter)
}

// FromInverseCentimeters creates a new ReciprocalLength instance from a value in InverseCentimeters.
func (uf ReciprocalLengthFactory) FromInverseCentimeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseCentimeter)
}

// FromInverseMillimeters creates a new ReciprocalLength instance from a value in InverseMillimeters.
func (uf ReciprocalLengthFactory) FromInverseMillimeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMillimeter)
}

// FromInverseMiles creates a new ReciprocalLength instance from a value in InverseMiles.
func (uf ReciprocalLengthFactory) FromInverseMiles(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMile)
}

// FromInverseYards creates a new ReciprocalLength instance from a value in InverseYards.
func (uf ReciprocalLengthFactory) FromInverseYards(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseYard)
}

// FromInverseFeet creates a new ReciprocalLength instance from a value in InverseFeet.
func (uf ReciprocalLengthFactory) FromInverseFeet(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseFoot)
}

// FromInverseUsSurveyFeet creates a new ReciprocalLength instance from a value in InverseUsSurveyFeet.
func (uf ReciprocalLengthFactory) FromInverseUsSurveyFeet(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseUsSurveyFoot)
}

// FromInverseInches creates a new ReciprocalLength instance from a value in InverseInches.
func (uf ReciprocalLengthFactory) FromInverseInches(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseInch)
}

// FromInverseMils creates a new ReciprocalLength instance from a value in InverseMils.
func (uf ReciprocalLengthFactory) FromInverseMils(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMil)
}

// FromInverseMicroinches creates a new ReciprocalLength instance from a value in InverseMicroinches.
func (uf ReciprocalLengthFactory) FromInverseMicroinches(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMicroinch)
}


// newReciprocalLength creates a new ReciprocalLength.
func newReciprocalLength(value float64, fromUnit ReciprocalLengthUnits) (*ReciprocalLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalReciprocalLengthUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ReciprocalLengthUnits", fromUnit)
	}
	a := &ReciprocalLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReciprocalLength in InverseMeter unit (the base/default quantity).
func (a *ReciprocalLength) BaseValue() float64 {
	return a.value
}


// InverseMeters returns the ReciprocalLength value in InverseMeters.
//
// 
func (a *ReciprocalLength) InverseMeters() float64 {
	if a.inverse_metersLazy != nil {
		return *a.inverse_metersLazy
	}
	inverse_meters := a.convertFromBase(ReciprocalLengthInverseMeter)
	a.inverse_metersLazy = &inverse_meters
	return inverse_meters
}

// InverseCentimeters returns the ReciprocalLength value in InverseCentimeters.
//
// 
func (a *ReciprocalLength) InverseCentimeters() float64 {
	if a.inverse_centimetersLazy != nil {
		return *a.inverse_centimetersLazy
	}
	inverse_centimeters := a.convertFromBase(ReciprocalLengthInverseCentimeter)
	a.inverse_centimetersLazy = &inverse_centimeters
	return inverse_centimeters
}

// InverseMillimeters returns the ReciprocalLength value in InverseMillimeters.
//
// 
func (a *ReciprocalLength) InverseMillimeters() float64 {
	if a.inverse_millimetersLazy != nil {
		return *a.inverse_millimetersLazy
	}
	inverse_millimeters := a.convertFromBase(ReciprocalLengthInverseMillimeter)
	a.inverse_millimetersLazy = &inverse_millimeters
	return inverse_millimeters
}

// InverseMiles returns the ReciprocalLength value in InverseMiles.
//
// 
func (a *ReciprocalLength) InverseMiles() float64 {
	if a.inverse_milesLazy != nil {
		return *a.inverse_milesLazy
	}
	inverse_miles := a.convertFromBase(ReciprocalLengthInverseMile)
	a.inverse_milesLazy = &inverse_miles
	return inverse_miles
}

// InverseYards returns the ReciprocalLength value in InverseYards.
//
// 
func (a *ReciprocalLength) InverseYards() float64 {
	if a.inverse_yardsLazy != nil {
		return *a.inverse_yardsLazy
	}
	inverse_yards := a.convertFromBase(ReciprocalLengthInverseYard)
	a.inverse_yardsLazy = &inverse_yards
	return inverse_yards
}

// InverseFeet returns the ReciprocalLength value in InverseFeet.
//
// 
func (a *ReciprocalLength) InverseFeet() float64 {
	if a.inverse_feetLazy != nil {
		return *a.inverse_feetLazy
	}
	inverse_feet := a.convertFromBase(ReciprocalLengthInverseFoot)
	a.inverse_feetLazy = &inverse_feet
	return inverse_feet
}

// InverseUsSurveyFeet returns the ReciprocalLength value in InverseUsSurveyFeet.
//
// 
func (a *ReciprocalLength) InverseUsSurveyFeet() float64 {
	if a.inverse_us_survey_feetLazy != nil {
		return *a.inverse_us_survey_feetLazy
	}
	inverse_us_survey_feet := a.convertFromBase(ReciprocalLengthInverseUsSurveyFoot)
	a.inverse_us_survey_feetLazy = &inverse_us_survey_feet
	return inverse_us_survey_feet
}

// InverseInches returns the ReciprocalLength value in InverseInches.
//
// 
func (a *ReciprocalLength) InverseInches() float64 {
	if a.inverse_inchesLazy != nil {
		return *a.inverse_inchesLazy
	}
	inverse_inches := a.convertFromBase(ReciprocalLengthInverseInch)
	a.inverse_inchesLazy = &inverse_inches
	return inverse_inches
}

// InverseMils returns the ReciprocalLength value in InverseMils.
//
// 
func (a *ReciprocalLength) InverseMils() float64 {
	if a.inverse_milsLazy != nil {
		return *a.inverse_milsLazy
	}
	inverse_mils := a.convertFromBase(ReciprocalLengthInverseMil)
	a.inverse_milsLazy = &inverse_mils
	return inverse_mils
}

// InverseMicroinches returns the ReciprocalLength value in InverseMicroinches.
//
// 
func (a *ReciprocalLength) InverseMicroinches() float64 {
	if a.inverse_microinchesLazy != nil {
		return *a.inverse_microinchesLazy
	}
	inverse_microinches := a.convertFromBase(ReciprocalLengthInverseMicroinch)
	a.inverse_microinchesLazy = &inverse_microinches
	return inverse_microinches
}


// ToDto creates a ReciprocalLengthDto representation from the ReciprocalLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InverseMeter by default.
func (a *ReciprocalLength) ToDto(holdInUnit *ReciprocalLengthUnits) ReciprocalLengthDto {
	if holdInUnit == nil {
		defaultUnit := ReciprocalLengthInverseMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ReciprocalLengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ReciprocalLength instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InverseMeter by default.
func (a *ReciprocalLength) ToDtoJSON(holdInUnit *ReciprocalLengthUnits) ([]byte, error) {
	// Convert to ReciprocalLengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ReciprocalLength to a specific unit value.
// The function uses the provided unit type (ReciprocalLengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ReciprocalLength) Convert(toUnit ReciprocalLengthUnits) float64 {
	switch toUnit { 
    case ReciprocalLengthInverseMeter:
		return a.InverseMeters()
    case ReciprocalLengthInverseCentimeter:
		return a.InverseCentimeters()
    case ReciprocalLengthInverseMillimeter:
		return a.InverseMillimeters()
    case ReciprocalLengthInverseMile:
		return a.InverseMiles()
    case ReciprocalLengthInverseYard:
		return a.InverseYards()
    case ReciprocalLengthInverseFoot:
		return a.InverseFeet()
    case ReciprocalLengthInverseUsSurveyFoot:
		return a.InverseUsSurveyFeet()
    case ReciprocalLengthInverseInch:
		return a.InverseInches()
    case ReciprocalLengthInverseMil:
		return a.InverseMils()
    case ReciprocalLengthInverseMicroinch:
		return a.InverseMicroinches()
	default:
		return math.NaN()
	}
}

func (a *ReciprocalLength) convertFromBase(toUnit ReciprocalLengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case ReciprocalLengthInverseMeter:
		return (value) 
	case ReciprocalLengthInverseCentimeter:
		return (value / 1e2) 
	case ReciprocalLengthInverseMillimeter:
		return (value / 1e3) 
	case ReciprocalLengthInverseMile:
		return (value * 1609.344) 
	case ReciprocalLengthInverseYard:
		return (value * 0.9144) 
	case ReciprocalLengthInverseFoot:
		return (value * 0.3048) 
	case ReciprocalLengthInverseUsSurveyFoot:
		return (value * 1200 / 3937) 
	case ReciprocalLengthInverseInch:
		return (value * 2.54e-2) 
	case ReciprocalLengthInverseMil:
		return (value * 2.54e-5) 
	case ReciprocalLengthInverseMicroinch:
		return (value * 2.54e-8) 
	default:
		return math.NaN()
	}
}

func (a *ReciprocalLength) convertToBase(value float64, fromUnit ReciprocalLengthUnits) float64 {
	switch fromUnit { 
	case ReciprocalLengthInverseMeter:
		return (value) 
	case ReciprocalLengthInverseCentimeter:
		return (value * 1e2) 
	case ReciprocalLengthInverseMillimeter:
		return (value * 1e3) 
	case ReciprocalLengthInverseMile:
		return (value / 1609.344) 
	case ReciprocalLengthInverseYard:
		return (value / 0.9144) 
	case ReciprocalLengthInverseFoot:
		return (value / 0.3048) 
	case ReciprocalLengthInverseUsSurveyFoot:
		return (value * 3937 / 1200) 
	case ReciprocalLengthInverseInch:
		return (value / 2.54e-2) 
	case ReciprocalLengthInverseMil:
		return (value / 2.54e-5) 
	case ReciprocalLengthInverseMicroinch:
		return (value / 2.54e-8) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ReciprocalLength in the default unit (InverseMeter),
// formatted to two decimal places.
func (a ReciprocalLength) String() string {
	return a.ToString(ReciprocalLengthInverseMeter, 2)
}

// ToString formats the ReciprocalLength value as a string with the specified unit and fractional digits.
// It converts the ReciprocalLength to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ReciprocalLength value will be converted (e.g., InverseMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ReciprocalLength with the unit abbreviation.
func (a *ReciprocalLength) ToString(unit ReciprocalLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetReciprocalLengthAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetReciprocalLengthAbbreviation(unit))
}

// Equals checks if the given ReciprocalLength is equal to the current ReciprocalLength.
//
// Parameters:
//    other: The ReciprocalLength to compare against.
//
// Returns:
//    bool: Returns true if both ReciprocalLength are equal, false otherwise.
func (a *ReciprocalLength) Equals(other *ReciprocalLength) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ReciprocalLength with another ReciprocalLength.
// It returns -1 if the current ReciprocalLength is less than the other ReciprocalLength, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ReciprocalLength to compare against.
//
// Returns:
//    int: -1 if the current ReciprocalLength is less, 1 if greater, and 0 if equal.
func (a *ReciprocalLength) CompareTo(other *ReciprocalLength) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ReciprocalLength to the current ReciprocalLength and returns the result.
// The result is a new ReciprocalLength instance with the sum of the values.
//
// Parameters:
//    other: The ReciprocalLength to add to the current ReciprocalLength.
//
// Returns:
//    *ReciprocalLength: A new ReciprocalLength instance representing the sum of both ReciprocalLength.
func (a *ReciprocalLength) Add(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ReciprocalLength from the current ReciprocalLength and returns the result.
// The result is a new ReciprocalLength instance with the difference of the values.
//
// Parameters:
//    other: The ReciprocalLength to subtract from the current ReciprocalLength.
//
// Returns:
//    *ReciprocalLength: A new ReciprocalLength instance representing the difference of both ReciprocalLength.
func (a *ReciprocalLength) Subtract(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ReciprocalLength by the given ReciprocalLength and returns the result.
// The result is a new ReciprocalLength instance with the product of the values.
//
// Parameters:
//    other: The ReciprocalLength to multiply with the current ReciprocalLength.
//
// Returns:
//    *ReciprocalLength: A new ReciprocalLength instance representing the product of both ReciprocalLength.
func (a *ReciprocalLength) Multiply(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value * other.BaseValue()}
}

// Divide divides the current ReciprocalLength by the given ReciprocalLength and returns the result.
// The result is a new ReciprocalLength instance with the quotient of the values.
//
// Parameters:
//    other: The ReciprocalLength to divide the current ReciprocalLength by.
//
// Returns:
//    *ReciprocalLength: A new ReciprocalLength instance representing the quotient of both ReciprocalLength.
func (a *ReciprocalLength) Divide(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value / other.BaseValue()}
}

// GetReciprocalLengthAbbreviation gets the unit abbreviation.
func GetReciprocalLengthAbbreviation(unit ReciprocalLengthUnits) string {
	switch unit { 
	case ReciprocalLengthInverseMeter:
		return "m⁻¹" 
	case ReciprocalLengthInverseCentimeter:
		return "cm⁻¹" 
	case ReciprocalLengthInverseMillimeter:
		return "mm⁻¹" 
	case ReciprocalLengthInverseMile:
		return "mi⁻¹" 
	case ReciprocalLengthInverseYard:
		return "yd⁻¹" 
	case ReciprocalLengthInverseFoot:
		return "ft⁻¹" 
	case ReciprocalLengthInverseUsSurveyFoot:
		return "ftUS⁻¹" 
	case ReciprocalLengthInverseInch:
		return "in⁻¹" 
	case ReciprocalLengthInverseMil:
		return "mil⁻¹" 
	case ReciprocalLengthInverseMicroinch:
		return "µin⁻¹" 
	default:
		return ""
	}
}