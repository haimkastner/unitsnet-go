// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AreaUnits defines various units of Area.
type AreaUnits string

const (
    
        // 
        AreaSquareKilometer AreaUnits = "SquareKilometer"
        // 
        AreaSquareMeter AreaUnits = "SquareMeter"
        // 
        AreaSquareDecimeter AreaUnits = "SquareDecimeter"
        // 
        AreaSquareCentimeter AreaUnits = "SquareCentimeter"
        // 
        AreaSquareMillimeter AreaUnits = "SquareMillimeter"
        // 
        AreaSquareMicrometer AreaUnits = "SquareMicrometer"
        // The statute mile was standardised between the British Commonwealth and the United States by an international agreement in 1959, when it was formally redefined with respect to SI units as exactly 1,609.344 metres.
        AreaSquareMile AreaUnits = "SquareMile"
        // The yard (symbol: yd) is an English unit of length in both the British imperial and US customary systems of measurement equalling 3 feet (or 36 inches). Since 1959 the yard has been by international agreement standardized as exactly 0.9144 meter. A distance of 1,760 yards is equal to 1 mile.
        AreaSquareYard AreaUnits = "SquareYard"
        // 
        AreaSquareFoot AreaUnits = "SquareFoot"
        // In the United States, the foot was defined as 12 inches, with the inch being defined by the Mendenhall Order of 1893 as 39.37 inches = 1 m. This makes a U.S. survey foot exactly 1200/3937 meters.
        AreaUsSurveySquareFoot AreaUnits = "UsSurveySquareFoot"
        // 
        AreaSquareInch AreaUnits = "SquareInch"
        // Based upon the international yard and pound agreement of 1959, an acre may be declared as exactly 4,046.8564224 square metres.
        AreaAcre AreaUnits = "Acre"
        // 
        AreaHectare AreaUnits = "Hectare"
        // 
        AreaSquareNauticalMile AreaUnits = "SquareNauticalMile"
)

// AreaDto represents a Area measurement with a numerical value and its corresponding unit.
type AreaDto struct {
    // Value is the numerical representation of the Area.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Area, as defined in the AreaUnits enumeration.
	Unit  AreaUnits `json:"unit"`
}

// AreaDtoFactory groups methods for creating and serializing AreaDto objects.
type AreaDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AreaDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AreaDtoFactory) FromJSON(data []byte) (*AreaDto, error) {
	a := AreaDto{}

    // Parse JSON into AreaDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AreaDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Area represents a measurement in a of Area.
//
// Area is a quantity that expresses the extent of a two-dimensional surface or shape, or planar lamina, in the plane. Area can be understood as the amount of material with a given thickness that would be necessary to fashion a model of the shape, or the amount of paint necessary to cover the surface with a single coat.[1] It is the two-dimensional analog of the length of a curve (a one-dimensional concept) or the volume of a solid (a three-dimensional concept).
type Area struct {
	// value is the base measurement stored internally.
	value       float64
    
    square_kilometersLazy *float64 
    square_metersLazy *float64 
    square_decimetersLazy *float64 
    square_centimetersLazy *float64 
    square_millimetersLazy *float64 
    square_micrometersLazy *float64 
    square_milesLazy *float64 
    square_yardsLazy *float64 
    square_feetLazy *float64 
    us_survey_square_feetLazy *float64 
    square_inchesLazy *float64 
    acresLazy *float64 
    hectaresLazy *float64 
    square_nautical_milesLazy *float64 
}

// AreaFactory groups methods for creating Area instances.
type AreaFactory struct{}

// CreateArea creates a new Area instance from the given value and unit.
func (uf AreaFactory) CreateArea(value float64, unit AreaUnits) (*Area, error) {
	return newArea(value, unit)
}

// FromDto converts a AreaDto to a Area instance.
func (uf AreaFactory) FromDto(dto AreaDto) (*Area, error) {
	return newArea(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Area instance.
func (uf AreaFactory) FromDtoJSON(data []byte) (*Area, error) {
	unitDto, err := AreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AreaDto from JSON: %w", err)
	}
	return AreaFactory{}.FromDto(*unitDto)
}


// FromSquareKilometers creates a new Area instance from a value in SquareKilometers.
func (uf AreaFactory) FromSquareKilometers(value float64) (*Area, error) {
	return newArea(value, AreaSquareKilometer)
}

// FromSquareMeters creates a new Area instance from a value in SquareMeters.
func (uf AreaFactory) FromSquareMeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareMeter)
}

// FromSquareDecimeters creates a new Area instance from a value in SquareDecimeters.
func (uf AreaFactory) FromSquareDecimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareDecimeter)
}

// FromSquareCentimeters creates a new Area instance from a value in SquareCentimeters.
func (uf AreaFactory) FromSquareCentimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareCentimeter)
}

// FromSquareMillimeters creates a new Area instance from a value in SquareMillimeters.
func (uf AreaFactory) FromSquareMillimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareMillimeter)
}

// FromSquareMicrometers creates a new Area instance from a value in SquareMicrometers.
func (uf AreaFactory) FromSquareMicrometers(value float64) (*Area, error) {
	return newArea(value, AreaSquareMicrometer)
}

// FromSquareMiles creates a new Area instance from a value in SquareMiles.
func (uf AreaFactory) FromSquareMiles(value float64) (*Area, error) {
	return newArea(value, AreaSquareMile)
}

// FromSquareYards creates a new Area instance from a value in SquareYards.
func (uf AreaFactory) FromSquareYards(value float64) (*Area, error) {
	return newArea(value, AreaSquareYard)
}

// FromSquareFeet creates a new Area instance from a value in SquareFeet.
func (uf AreaFactory) FromSquareFeet(value float64) (*Area, error) {
	return newArea(value, AreaSquareFoot)
}

// FromUsSurveySquareFeet creates a new Area instance from a value in UsSurveySquareFeet.
func (uf AreaFactory) FromUsSurveySquareFeet(value float64) (*Area, error) {
	return newArea(value, AreaUsSurveySquareFoot)
}

// FromSquareInches creates a new Area instance from a value in SquareInches.
func (uf AreaFactory) FromSquareInches(value float64) (*Area, error) {
	return newArea(value, AreaSquareInch)
}

// FromAcres creates a new Area instance from a value in Acres.
func (uf AreaFactory) FromAcres(value float64) (*Area, error) {
	return newArea(value, AreaAcre)
}

// FromHectares creates a new Area instance from a value in Hectares.
func (uf AreaFactory) FromHectares(value float64) (*Area, error) {
	return newArea(value, AreaHectare)
}

// FromSquareNauticalMiles creates a new Area instance from a value in SquareNauticalMiles.
func (uf AreaFactory) FromSquareNauticalMiles(value float64) (*Area, error) {
	return newArea(value, AreaSquareNauticalMile)
}


// newArea creates a new Area.
func newArea(value float64, fromUnit AreaUnits) (*Area, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Area{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Area in SquareMeter unit (the base/default quantity).
func (a *Area) BaseValue() float64 {
	return a.value
}


// SquareKilometers returns the Area value in SquareKilometers.
//
// 
func (a *Area) SquareKilometers() float64 {
	if a.square_kilometersLazy != nil {
		return *a.square_kilometersLazy
	}
	square_kilometers := a.convertFromBase(AreaSquareKilometer)
	a.square_kilometersLazy = &square_kilometers
	return square_kilometers
}

// SquareMeters returns the Area value in SquareMeters.
//
// 
func (a *Area) SquareMeters() float64 {
	if a.square_metersLazy != nil {
		return *a.square_metersLazy
	}
	square_meters := a.convertFromBase(AreaSquareMeter)
	a.square_metersLazy = &square_meters
	return square_meters
}

// SquareDecimeters returns the Area value in SquareDecimeters.
//
// 
func (a *Area) SquareDecimeters() float64 {
	if a.square_decimetersLazy != nil {
		return *a.square_decimetersLazy
	}
	square_decimeters := a.convertFromBase(AreaSquareDecimeter)
	a.square_decimetersLazy = &square_decimeters
	return square_decimeters
}

// SquareCentimeters returns the Area value in SquareCentimeters.
//
// 
func (a *Area) SquareCentimeters() float64 {
	if a.square_centimetersLazy != nil {
		return *a.square_centimetersLazy
	}
	square_centimeters := a.convertFromBase(AreaSquareCentimeter)
	a.square_centimetersLazy = &square_centimeters
	return square_centimeters
}

// SquareMillimeters returns the Area value in SquareMillimeters.
//
// 
func (a *Area) SquareMillimeters() float64 {
	if a.square_millimetersLazy != nil {
		return *a.square_millimetersLazy
	}
	square_millimeters := a.convertFromBase(AreaSquareMillimeter)
	a.square_millimetersLazy = &square_millimeters
	return square_millimeters
}

// SquareMicrometers returns the Area value in SquareMicrometers.
//
// 
func (a *Area) SquareMicrometers() float64 {
	if a.square_micrometersLazy != nil {
		return *a.square_micrometersLazy
	}
	square_micrometers := a.convertFromBase(AreaSquareMicrometer)
	a.square_micrometersLazy = &square_micrometers
	return square_micrometers
}

// SquareMiles returns the Area value in SquareMiles.
//
// The statute mile was standardised between the British Commonwealth and the United States by an international agreement in 1959, when it was formally redefined with respect to SI units as exactly 1,609.344 metres.
func (a *Area) SquareMiles() float64 {
	if a.square_milesLazy != nil {
		return *a.square_milesLazy
	}
	square_miles := a.convertFromBase(AreaSquareMile)
	a.square_milesLazy = &square_miles
	return square_miles
}

// SquareYards returns the Area value in SquareYards.
//
// The yard (symbol: yd) is an English unit of length in both the British imperial and US customary systems of measurement equalling 3 feet (or 36 inches). Since 1959 the yard has been by international agreement standardized as exactly 0.9144 meter. A distance of 1,760 yards is equal to 1 mile.
func (a *Area) SquareYards() float64 {
	if a.square_yardsLazy != nil {
		return *a.square_yardsLazy
	}
	square_yards := a.convertFromBase(AreaSquareYard)
	a.square_yardsLazy = &square_yards
	return square_yards
}

// SquareFeet returns the Area value in SquareFeet.
//
// 
func (a *Area) SquareFeet() float64 {
	if a.square_feetLazy != nil {
		return *a.square_feetLazy
	}
	square_feet := a.convertFromBase(AreaSquareFoot)
	a.square_feetLazy = &square_feet
	return square_feet
}

// UsSurveySquareFeet returns the Area value in UsSurveySquareFeet.
//
// In the United States, the foot was defined as 12 inches, with the inch being defined by the Mendenhall Order of 1893 as 39.37 inches = 1 m. This makes a U.S. survey foot exactly 1200/3937 meters.
func (a *Area) UsSurveySquareFeet() float64 {
	if a.us_survey_square_feetLazy != nil {
		return *a.us_survey_square_feetLazy
	}
	us_survey_square_feet := a.convertFromBase(AreaUsSurveySquareFoot)
	a.us_survey_square_feetLazy = &us_survey_square_feet
	return us_survey_square_feet
}

// SquareInches returns the Area value in SquareInches.
//
// 
func (a *Area) SquareInches() float64 {
	if a.square_inchesLazy != nil {
		return *a.square_inchesLazy
	}
	square_inches := a.convertFromBase(AreaSquareInch)
	a.square_inchesLazy = &square_inches
	return square_inches
}

// Acres returns the Area value in Acres.
//
// Based upon the international yard and pound agreement of 1959, an acre may be declared as exactly 4,046.8564224 square metres.
func (a *Area) Acres() float64 {
	if a.acresLazy != nil {
		return *a.acresLazy
	}
	acres := a.convertFromBase(AreaAcre)
	a.acresLazy = &acres
	return acres
}

// Hectares returns the Area value in Hectares.
//
// 
func (a *Area) Hectares() float64 {
	if a.hectaresLazy != nil {
		return *a.hectaresLazy
	}
	hectares := a.convertFromBase(AreaHectare)
	a.hectaresLazy = &hectares
	return hectares
}

// SquareNauticalMiles returns the Area value in SquareNauticalMiles.
//
// 
func (a *Area) SquareNauticalMiles() float64 {
	if a.square_nautical_milesLazy != nil {
		return *a.square_nautical_milesLazy
	}
	square_nautical_miles := a.convertFromBase(AreaSquareNauticalMile)
	a.square_nautical_milesLazy = &square_nautical_miles
	return square_nautical_miles
}


// ToDto creates a AreaDto representation from the Area instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeter by default.
func (a *Area) ToDto(holdInUnit *AreaUnits) AreaDto {
	if holdInUnit == nil {
		defaultUnit := AreaSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return AreaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Area instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeter by default.
func (a *Area) ToDtoJSON(holdInUnit *AreaUnits) ([]byte, error) {
	// Convert to AreaDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Area to a specific unit value.
// The function uses the provided unit type (AreaUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Area) Convert(toUnit AreaUnits) float64 {
	switch toUnit { 
    case AreaSquareKilometer:
		return a.SquareKilometers()
    case AreaSquareMeter:
		return a.SquareMeters()
    case AreaSquareDecimeter:
		return a.SquareDecimeters()
    case AreaSquareCentimeter:
		return a.SquareCentimeters()
    case AreaSquareMillimeter:
		return a.SquareMillimeters()
    case AreaSquareMicrometer:
		return a.SquareMicrometers()
    case AreaSquareMile:
		return a.SquareMiles()
    case AreaSquareYard:
		return a.SquareYards()
    case AreaSquareFoot:
		return a.SquareFeet()
    case AreaUsSurveySquareFoot:
		return a.UsSurveySquareFeet()
    case AreaSquareInch:
		return a.SquareInches()
    case AreaAcre:
		return a.Acres()
    case AreaHectare:
		return a.Hectares()
    case AreaSquareNauticalMile:
		return a.SquareNauticalMiles()
	default:
		return math.NaN()
	}
}

func (a *Area) convertFromBase(toUnit AreaUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaSquareKilometer:
		return (value / 1e6) 
	case AreaSquareMeter:
		return (value) 
	case AreaSquareDecimeter:
		return (value / 1e-2) 
	case AreaSquareCentimeter:
		return (value / 1e-4) 
	case AreaSquareMillimeter:
		return (value / 1e-6) 
	case AreaSquareMicrometer:
		return (value / 1e-12) 
	case AreaSquareMile:
		return (value / 1609.344 / 1609.344) 
	case AreaSquareYard:
		return (value / 0.9144 / 0.9144) 
	case AreaSquareFoot:
		return (value / 9.290304e-2) 
	case AreaUsSurveySquareFoot:
		return (value / (1200.0 / 3937.0) / (1200.0 / 3937.0)) 
	case AreaSquareInch:
		return (value / 0.00064516) 
	case AreaAcre:
		return (value / 4046.8564224) 
	case AreaHectare:
		return (value / 1e4) 
	case AreaSquareNauticalMile:
		return (value / 3429904) 
	default:
		return math.NaN()
	}
}

func (a *Area) convertToBase(value float64, fromUnit AreaUnits) float64 {
	switch fromUnit { 
	case AreaSquareKilometer:
		return (value * 1e6) 
	case AreaSquareMeter:
		return (value) 
	case AreaSquareDecimeter:
		return (value * 1e-2) 
	case AreaSquareCentimeter:
		return (value * 1e-4) 
	case AreaSquareMillimeter:
		return (value * 1e-6) 
	case AreaSquareMicrometer:
		return (value * 1e-12) 
	case AreaSquareMile:
		return (value * 1609.344 * 1609.344) 
	case AreaSquareYard:
		return (value * 0.9144 * 0.9144) 
	case AreaSquareFoot:
		return (value * 9.290304e-2) 
	case AreaUsSurveySquareFoot:
		return (value * (1200.0 / 3937.0) * (1200.0 / 3937.0)) 
	case AreaSquareInch:
		return (value * 0.00064516) 
	case AreaAcre:
		return (value * 4046.8564224) 
	case AreaHectare:
		return (value * 1e4) 
	case AreaSquareNauticalMile:
		return (value * 3429904) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Area in the default unit (SquareMeter),
// formatted to two decimal places.
func (a Area) String() string {
	return a.ToString(AreaSquareMeter, 2)
}

// ToString formats the Area value as a string with the specified unit and fractional digits.
// It converts the Area to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Area value will be converted (e.g., SquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Area with the unit abbreviation.
func (a *Area) ToString(unit AreaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAreaAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAreaAbbreviation(unit))
}

// Equals checks if the given Area is equal to the current Area.
//
// Parameters:
//    other: The Area to compare against.
//
// Returns:
//    bool: Returns true if both Area are equal, false otherwise.
func (a *Area) Equals(other *Area) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Area with another Area.
// It returns -1 if the current Area is less than the other Area, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Area to compare against.
//
// Returns:
//    int: -1 if the current Area is less, 1 if greater, and 0 if equal.
func (a *Area) CompareTo(other *Area) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Area to the current Area and returns the result.
// The result is a new Area instance with the sum of the values.
//
// Parameters:
//    other: The Area to add to the current Area.
//
// Returns:
//    *Area: A new Area instance representing the sum of both Area.
func (a *Area) Add(other *Area) *Area {
	return &Area{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Area from the current Area and returns the result.
// The result is a new Area instance with the difference of the values.
//
// Parameters:
//    other: The Area to subtract from the current Area.
//
// Returns:
//    *Area: A new Area instance representing the difference of both Area.
func (a *Area) Subtract(other *Area) *Area {
	return &Area{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Area by the given Area and returns the result.
// The result is a new Area instance with the product of the values.
//
// Parameters:
//    other: The Area to multiply with the current Area.
//
// Returns:
//    *Area: A new Area instance representing the product of both Area.
func (a *Area) Multiply(other *Area) *Area {
	return &Area{value: a.value * other.BaseValue()}
}

// Divide divides the current Area by the given Area and returns the result.
// The result is a new Area instance with the quotient of the values.
//
// Parameters:
//    other: The Area to divide the current Area by.
//
// Returns:
//    *Area: A new Area instance representing the quotient of both Area.
func (a *Area) Divide(other *Area) *Area {
	return &Area{value: a.value / other.BaseValue()}
}

// GetAreaAbbreviation gets the unit abbreviation.
func GetAreaAbbreviation(unit AreaUnits) string {
	switch unit { 
	case AreaSquareKilometer:
		return "km²" 
	case AreaSquareMeter:
		return "m²" 
	case AreaSquareDecimeter:
		return "dm²" 
	case AreaSquareCentimeter:
		return "cm²" 
	case AreaSquareMillimeter:
		return "mm²" 
	case AreaSquareMicrometer:
		return "µm²" 
	case AreaSquareMile:
		return "mi²" 
	case AreaSquareYard:
		return "yd²" 
	case AreaSquareFoot:
		return "ft²" 
	case AreaUsSurveySquareFoot:
		return "ft² (US)" 
	case AreaSquareInch:
		return "in²" 
	case AreaAcre:
		return "ac" 
	case AreaHectare:
		return "ha" 
	case AreaSquareNauticalMile:
		return "nmi²" 
	default:
		return ""
	}
}