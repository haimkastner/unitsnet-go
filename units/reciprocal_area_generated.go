// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ReciprocalAreaUnits defines various units of ReciprocalArea.
type ReciprocalAreaUnits string

const (
    
        // 
        ReciprocalAreaInverseSquareMeter ReciprocalAreaUnits = "InverseSquareMeter"
        // 
        ReciprocalAreaInverseSquareKilometer ReciprocalAreaUnits = "InverseSquareKilometer"
        // 
        ReciprocalAreaInverseSquareDecimeter ReciprocalAreaUnits = "InverseSquareDecimeter"
        // 
        ReciprocalAreaInverseSquareCentimeter ReciprocalAreaUnits = "InverseSquareCentimeter"
        // 
        ReciprocalAreaInverseSquareMillimeter ReciprocalAreaUnits = "InverseSquareMillimeter"
        // 
        ReciprocalAreaInverseSquareMicrometer ReciprocalAreaUnits = "InverseSquareMicrometer"
        // 
        ReciprocalAreaInverseSquareMile ReciprocalAreaUnits = "InverseSquareMile"
        // 
        ReciprocalAreaInverseSquareYard ReciprocalAreaUnits = "InverseSquareYard"
        // 
        ReciprocalAreaInverseSquareFoot ReciprocalAreaUnits = "InverseSquareFoot"
        // 
        ReciprocalAreaInverseUsSurveySquareFoot ReciprocalAreaUnits = "InverseUsSurveySquareFoot"
        // 
        ReciprocalAreaInverseSquareInch ReciprocalAreaUnits = "InverseSquareInch"
)

// ReciprocalAreaDto represents a ReciprocalArea measurement with a numerical value and its corresponding unit.
type ReciprocalAreaDto struct {
    // Value is the numerical representation of the ReciprocalArea.
	Value float64
    // Unit specifies the unit of measurement for the ReciprocalArea, as defined in the ReciprocalAreaUnits enumeration.
	Unit  ReciprocalAreaUnits
}

// ReciprocalAreaDtoFactory groups methods for creating and serializing ReciprocalAreaDto objects.
type ReciprocalAreaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ReciprocalAreaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ReciprocalAreaDtoFactory) FromJSON(data []byte) (*ReciprocalAreaDto, error) {
	a := ReciprocalAreaDto{}

    // Parse JSON into ReciprocalAreaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ReciprocalAreaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ReciprocalAreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ReciprocalArea represents a measurement in a of ReciprocalArea.
//
// Reciprocal area (Inverse-square) quantity is used to specify a physical quantity inversely proportional to the square of the distance.
type ReciprocalArea struct {
	// value is the base measurement stored internally.
	value       float64
    
    inverse_square_metersLazy *float64 
    inverse_square_kilometersLazy *float64 
    inverse_square_decimetersLazy *float64 
    inverse_square_centimetersLazy *float64 
    inverse_square_millimetersLazy *float64 
    inverse_square_micrometersLazy *float64 
    inverse_square_milesLazy *float64 
    inverse_square_yardsLazy *float64 
    inverse_square_feetLazy *float64 
    inverse_us_survey_square_feetLazy *float64 
    inverse_square_inchesLazy *float64 
}

// ReciprocalAreaFactory groups methods for creating ReciprocalArea instances.
type ReciprocalAreaFactory struct{}

// CreateReciprocalArea creates a new ReciprocalArea instance from the given value and unit.
func (uf ReciprocalAreaFactory) CreateReciprocalArea(value float64, unit ReciprocalAreaUnits) (*ReciprocalArea, error) {
	return newReciprocalArea(value, unit)
}

// FromDto converts a ReciprocalAreaDto to a ReciprocalArea instance.
func (uf ReciprocalAreaFactory) FromDto(dto ReciprocalAreaDto) (*ReciprocalArea, error) {
	return newReciprocalArea(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ReciprocalArea instance.
func (uf ReciprocalAreaFactory) FromDtoJSON(data []byte) (*ReciprocalArea, error) {
	unitDto, err := ReciprocalAreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ReciprocalAreaDto from JSON: %w", err)
	}
	return ReciprocalAreaFactory{}.FromDto(*unitDto)
}


// FromInverseSquareMeters creates a new ReciprocalArea instance from a value in InverseSquareMeters.
func (uf ReciprocalAreaFactory) FromInverseSquareMeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMeter)
}

// FromInverseSquareKilometers creates a new ReciprocalArea instance from a value in InverseSquareKilometers.
func (uf ReciprocalAreaFactory) FromInverseSquareKilometers(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareKilometer)
}

// FromInverseSquareDecimeters creates a new ReciprocalArea instance from a value in InverseSquareDecimeters.
func (uf ReciprocalAreaFactory) FromInverseSquareDecimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareDecimeter)
}

// FromInverseSquareCentimeters creates a new ReciprocalArea instance from a value in InverseSquareCentimeters.
func (uf ReciprocalAreaFactory) FromInverseSquareCentimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareCentimeter)
}

// FromInverseSquareMillimeters creates a new ReciprocalArea instance from a value in InverseSquareMillimeters.
func (uf ReciprocalAreaFactory) FromInverseSquareMillimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMillimeter)
}

// FromInverseSquareMicrometers creates a new ReciprocalArea instance from a value in InverseSquareMicrometers.
func (uf ReciprocalAreaFactory) FromInverseSquareMicrometers(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMicrometer)
}

// FromInverseSquareMiles creates a new ReciprocalArea instance from a value in InverseSquareMiles.
func (uf ReciprocalAreaFactory) FromInverseSquareMiles(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMile)
}

// FromInverseSquareYards creates a new ReciprocalArea instance from a value in InverseSquareYards.
func (uf ReciprocalAreaFactory) FromInverseSquareYards(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareYard)
}

// FromInverseSquareFeet creates a new ReciprocalArea instance from a value in InverseSquareFeet.
func (uf ReciprocalAreaFactory) FromInverseSquareFeet(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareFoot)
}

// FromInverseUsSurveySquareFeet creates a new ReciprocalArea instance from a value in InverseUsSurveySquareFeet.
func (uf ReciprocalAreaFactory) FromInverseUsSurveySquareFeet(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseUsSurveySquareFoot)
}

// FromInverseSquareInches creates a new ReciprocalArea instance from a value in InverseSquareInches.
func (uf ReciprocalAreaFactory) FromInverseSquareInches(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareInch)
}


// newReciprocalArea creates a new ReciprocalArea.
func newReciprocalArea(value float64, fromUnit ReciprocalAreaUnits) (*ReciprocalArea, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ReciprocalArea{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReciprocalArea in InverseSquareMeter unit (the base/default quantity).
func (a *ReciprocalArea) BaseValue() float64 {
	return a.value
}


// InverseSquareMeters returns the ReciprocalArea value in InverseSquareMeters.
//
// 
func (a *ReciprocalArea) InverseSquareMeters() float64 {
	if a.inverse_square_metersLazy != nil {
		return *a.inverse_square_metersLazy
	}
	inverse_square_meters := a.convertFromBase(ReciprocalAreaInverseSquareMeter)
	a.inverse_square_metersLazy = &inverse_square_meters
	return inverse_square_meters
}

// InverseSquareKilometers returns the ReciprocalArea value in InverseSquareKilometers.
//
// 
func (a *ReciprocalArea) InverseSquareKilometers() float64 {
	if a.inverse_square_kilometersLazy != nil {
		return *a.inverse_square_kilometersLazy
	}
	inverse_square_kilometers := a.convertFromBase(ReciprocalAreaInverseSquareKilometer)
	a.inverse_square_kilometersLazy = &inverse_square_kilometers
	return inverse_square_kilometers
}

// InverseSquareDecimeters returns the ReciprocalArea value in InverseSquareDecimeters.
//
// 
func (a *ReciprocalArea) InverseSquareDecimeters() float64 {
	if a.inverse_square_decimetersLazy != nil {
		return *a.inverse_square_decimetersLazy
	}
	inverse_square_decimeters := a.convertFromBase(ReciprocalAreaInverseSquareDecimeter)
	a.inverse_square_decimetersLazy = &inverse_square_decimeters
	return inverse_square_decimeters
}

// InverseSquareCentimeters returns the ReciprocalArea value in InverseSquareCentimeters.
//
// 
func (a *ReciprocalArea) InverseSquareCentimeters() float64 {
	if a.inverse_square_centimetersLazy != nil {
		return *a.inverse_square_centimetersLazy
	}
	inverse_square_centimeters := a.convertFromBase(ReciprocalAreaInverseSquareCentimeter)
	a.inverse_square_centimetersLazy = &inverse_square_centimeters
	return inverse_square_centimeters
}

// InverseSquareMillimeters returns the ReciprocalArea value in InverseSquareMillimeters.
//
// 
func (a *ReciprocalArea) InverseSquareMillimeters() float64 {
	if a.inverse_square_millimetersLazy != nil {
		return *a.inverse_square_millimetersLazy
	}
	inverse_square_millimeters := a.convertFromBase(ReciprocalAreaInverseSquareMillimeter)
	a.inverse_square_millimetersLazy = &inverse_square_millimeters
	return inverse_square_millimeters
}

// InverseSquareMicrometers returns the ReciprocalArea value in InverseSquareMicrometers.
//
// 
func (a *ReciprocalArea) InverseSquareMicrometers() float64 {
	if a.inverse_square_micrometersLazy != nil {
		return *a.inverse_square_micrometersLazy
	}
	inverse_square_micrometers := a.convertFromBase(ReciprocalAreaInverseSquareMicrometer)
	a.inverse_square_micrometersLazy = &inverse_square_micrometers
	return inverse_square_micrometers
}

// InverseSquareMiles returns the ReciprocalArea value in InverseSquareMiles.
//
// 
func (a *ReciprocalArea) InverseSquareMiles() float64 {
	if a.inverse_square_milesLazy != nil {
		return *a.inverse_square_milesLazy
	}
	inverse_square_miles := a.convertFromBase(ReciprocalAreaInverseSquareMile)
	a.inverse_square_milesLazy = &inverse_square_miles
	return inverse_square_miles
}

// InverseSquareYards returns the ReciprocalArea value in InverseSquareYards.
//
// 
func (a *ReciprocalArea) InverseSquareYards() float64 {
	if a.inverse_square_yardsLazy != nil {
		return *a.inverse_square_yardsLazy
	}
	inverse_square_yards := a.convertFromBase(ReciprocalAreaInverseSquareYard)
	a.inverse_square_yardsLazy = &inverse_square_yards
	return inverse_square_yards
}

// InverseSquareFeet returns the ReciprocalArea value in InverseSquareFeet.
//
// 
func (a *ReciprocalArea) InverseSquareFeet() float64 {
	if a.inverse_square_feetLazy != nil {
		return *a.inverse_square_feetLazy
	}
	inverse_square_feet := a.convertFromBase(ReciprocalAreaInverseSquareFoot)
	a.inverse_square_feetLazy = &inverse_square_feet
	return inverse_square_feet
}

// InverseUsSurveySquareFeet returns the ReciprocalArea value in InverseUsSurveySquareFeet.
//
// 
func (a *ReciprocalArea) InverseUsSurveySquareFeet() float64 {
	if a.inverse_us_survey_square_feetLazy != nil {
		return *a.inverse_us_survey_square_feetLazy
	}
	inverse_us_survey_square_feet := a.convertFromBase(ReciprocalAreaInverseUsSurveySquareFoot)
	a.inverse_us_survey_square_feetLazy = &inverse_us_survey_square_feet
	return inverse_us_survey_square_feet
}

// InverseSquareInches returns the ReciprocalArea value in InverseSquareInches.
//
// 
func (a *ReciprocalArea) InverseSquareInches() float64 {
	if a.inverse_square_inchesLazy != nil {
		return *a.inverse_square_inchesLazy
	}
	inverse_square_inches := a.convertFromBase(ReciprocalAreaInverseSquareInch)
	a.inverse_square_inchesLazy = &inverse_square_inches
	return inverse_square_inches
}


// ToDto creates a ReciprocalAreaDto representation from the ReciprocalArea instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InverseSquareMeter by default.
func (a *ReciprocalArea) ToDto(holdInUnit *ReciprocalAreaUnits) ReciprocalAreaDto {
	if holdInUnit == nil {
		defaultUnit := ReciprocalAreaInverseSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ReciprocalAreaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ReciprocalArea instance.
//
// If the provided holdInUnit is nil, the value will be repesented by InverseSquareMeter by default.
func (a *ReciprocalArea) ToDtoJSON(holdInUnit *ReciprocalAreaUnits) ([]byte, error) {
	// Convert to ReciprocalAreaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ReciprocalArea to a specific unit value.
// The function uses the provided unit type (ReciprocalAreaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ReciprocalArea) Convert(toUnit ReciprocalAreaUnits) float64 {
	switch toUnit { 
    case ReciprocalAreaInverseSquareMeter:
		return a.InverseSquareMeters()
    case ReciprocalAreaInverseSquareKilometer:
		return a.InverseSquareKilometers()
    case ReciprocalAreaInverseSquareDecimeter:
		return a.InverseSquareDecimeters()
    case ReciprocalAreaInverseSquareCentimeter:
		return a.InverseSquareCentimeters()
    case ReciprocalAreaInverseSquareMillimeter:
		return a.InverseSquareMillimeters()
    case ReciprocalAreaInverseSquareMicrometer:
		return a.InverseSquareMicrometers()
    case ReciprocalAreaInverseSquareMile:
		return a.InverseSquareMiles()
    case ReciprocalAreaInverseSquareYard:
		return a.InverseSquareYards()
    case ReciprocalAreaInverseSquareFoot:
		return a.InverseSquareFeet()
    case ReciprocalAreaInverseUsSurveySquareFoot:
		return a.InverseUsSurveySquareFeet()
    case ReciprocalAreaInverseSquareInch:
		return a.InverseSquareInches()
	default:
		return math.NaN()
	}
}

func (a *ReciprocalArea) convertFromBase(toUnit ReciprocalAreaUnits) float64 {
    value := a.value
	switch toUnit { 
	case ReciprocalAreaInverseSquareMeter:
		return (value) 
	case ReciprocalAreaInverseSquareKilometer:
		return (value * 1e6) 
	case ReciprocalAreaInverseSquareDecimeter:
		return (value * 1e-2) 
	case ReciprocalAreaInverseSquareCentimeter:
		return (value * 1e-4) 
	case ReciprocalAreaInverseSquareMillimeter:
		return (value * 1e-6) 
	case ReciprocalAreaInverseSquareMicrometer:
		return (value * 1e-12) 
	case ReciprocalAreaInverseSquareMile:
		return (value * 2.59e6) 
	case ReciprocalAreaInverseSquareYard:
		return (value * 0.836127) 
	case ReciprocalAreaInverseSquareFoot:
		return (value * 0.092903) 
	case ReciprocalAreaInverseUsSurveySquareFoot:
		return (value * 0.09290341161) 
	case ReciprocalAreaInverseSquareInch:
		return (value * 0.00064516) 
	default:
		return math.NaN()
	}
}

func (a *ReciprocalArea) convertToBase(value float64, fromUnit ReciprocalAreaUnits) float64 {
	switch fromUnit { 
	case ReciprocalAreaInverseSquareMeter:
		return (value) 
	case ReciprocalAreaInverseSquareKilometer:
		return (value / 1e6) 
	case ReciprocalAreaInverseSquareDecimeter:
		return (value / 1e-2) 
	case ReciprocalAreaInverseSquareCentimeter:
		return (value / 1e-4) 
	case ReciprocalAreaInverseSquareMillimeter:
		return (value / 1e-6) 
	case ReciprocalAreaInverseSquareMicrometer:
		return (value / 1e-12) 
	case ReciprocalAreaInverseSquareMile:
		return (value / 2.59e6) 
	case ReciprocalAreaInverseSquareYard:
		return (value / 0.836127) 
	case ReciprocalAreaInverseSquareFoot:
		return (value / 0.092903) 
	case ReciprocalAreaInverseUsSurveySquareFoot:
		return (value / 0.09290341161) 
	case ReciprocalAreaInverseSquareInch:
		return (value / 0.00064516) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ReciprocalArea in the default unit (InverseSquareMeter),
// formatted to two decimal places.
func (a ReciprocalArea) String() string {
	return a.ToString(ReciprocalAreaInverseSquareMeter, 2)
}

// ToString formats the ReciprocalArea value as a string with the specified unit and fractional digits.
// It converts the ReciprocalArea to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ReciprocalArea value will be converted (e.g., InverseSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ReciprocalArea with the unit abbreviation.
func (a *ReciprocalArea) ToString(unit ReciprocalAreaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ReciprocalArea) getUnitAbbreviation(unit ReciprocalAreaUnits) string {
	switch unit { 
	case ReciprocalAreaInverseSquareMeter:
		return "m⁻²" 
	case ReciprocalAreaInverseSquareKilometer:
		return "km⁻²" 
	case ReciprocalAreaInverseSquareDecimeter:
		return "dm⁻²" 
	case ReciprocalAreaInverseSquareCentimeter:
		return "cm⁻²" 
	case ReciprocalAreaInverseSquareMillimeter:
		return "mm⁻²" 
	case ReciprocalAreaInverseSquareMicrometer:
		return "µm⁻²" 
	case ReciprocalAreaInverseSquareMile:
		return "mi⁻²" 
	case ReciprocalAreaInverseSquareYard:
		return "yd⁻²" 
	case ReciprocalAreaInverseSquareFoot:
		return "ft⁻²" 
	case ReciprocalAreaInverseUsSurveySquareFoot:
		return "ft⁻² (US)" 
	case ReciprocalAreaInverseSquareInch:
		return "in⁻²" 
	default:
		return ""
	}
}

// Equals checks if the given ReciprocalArea is equal to the current ReciprocalArea.
//
// Parameters:
//    other: The ReciprocalArea to compare against.
//
// Returns:
//    bool: Returns true if both ReciprocalArea are equal, false otherwise.
func (a *ReciprocalArea) Equals(other *ReciprocalArea) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ReciprocalArea with another ReciprocalArea.
// It returns -1 if the current ReciprocalArea is less than the other ReciprocalArea, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ReciprocalArea to compare against.
//
// Returns:
//    int: -1 if the current ReciprocalArea is less, 1 if greater, and 0 if equal.
func (a *ReciprocalArea) CompareTo(other *ReciprocalArea) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ReciprocalArea to the current ReciprocalArea and returns the result.
// The result is a new ReciprocalArea instance with the sum of the values.
//
// Parameters:
//    other: The ReciprocalArea to add to the current ReciprocalArea.
//
// Returns:
//    *ReciprocalArea: A new ReciprocalArea instance representing the sum of both ReciprocalArea.
func (a *ReciprocalArea) Add(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ReciprocalArea from the current ReciprocalArea and returns the result.
// The result is a new ReciprocalArea instance with the difference of the values.
//
// Parameters:
//    other: The ReciprocalArea to subtract from the current ReciprocalArea.
//
// Returns:
//    *ReciprocalArea: A new ReciprocalArea instance representing the difference of both ReciprocalArea.
func (a *ReciprocalArea) Subtract(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ReciprocalArea by the given ReciprocalArea and returns the result.
// The result is a new ReciprocalArea instance with the product of the values.
//
// Parameters:
//    other: The ReciprocalArea to multiply with the current ReciprocalArea.
//
// Returns:
//    *ReciprocalArea: A new ReciprocalArea instance representing the product of both ReciprocalArea.
func (a *ReciprocalArea) Multiply(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value * other.BaseValue()}
}

// Divide divides the current ReciprocalArea by the given ReciprocalArea and returns the result.
// The result is a new ReciprocalArea instance with the quotient of the values.
//
// Parameters:
//    other: The ReciprocalArea to divide the current ReciprocalArea by.
//
// Returns:
//    *ReciprocalArea: A new ReciprocalArea instance representing the quotient of both ReciprocalArea.
func (a *ReciprocalArea) Divide(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value / other.BaseValue()}
}