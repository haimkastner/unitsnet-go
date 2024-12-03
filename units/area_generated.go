// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AreaUnits enumeration
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

// AreaDto represents an Area
type AreaDto struct {
	Value float64
	Unit  AreaUnits
}

// AreaDtoFactory struct to group related functions
type AreaDtoFactory struct{}

func (udf AreaDtoFactory) FromJSON(data []byte) (*AreaDto, error) {
	a := AreaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Area struct
type Area struct {
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

// AreaFactory struct to group related functions
type AreaFactory struct{}

func (uf AreaFactory) CreateArea(value float64, unit AreaUnits) (*Area, error) {
	return newArea(value, unit)
}

func (uf AreaFactory) FromDto(dto AreaDto) (*Area, error) {
	return newArea(dto.Value, dto.Unit)
}

func (uf AreaFactory) FromDtoJSON(data []byte) (*Area, error) {
	unitDto, err := AreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AreaFactory{}.FromDto(*unitDto)
}


// FromSquareKilometer creates a new Area instance from SquareKilometer.
func (uf AreaFactory) FromSquareKilometers(value float64) (*Area, error) {
	return newArea(value, AreaSquareKilometer)
}

// FromSquareMeter creates a new Area instance from SquareMeter.
func (uf AreaFactory) FromSquareMeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareMeter)
}

// FromSquareDecimeter creates a new Area instance from SquareDecimeter.
func (uf AreaFactory) FromSquareDecimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareDecimeter)
}

// FromSquareCentimeter creates a new Area instance from SquareCentimeter.
func (uf AreaFactory) FromSquareCentimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareCentimeter)
}

// FromSquareMillimeter creates a new Area instance from SquareMillimeter.
func (uf AreaFactory) FromSquareMillimeters(value float64) (*Area, error) {
	return newArea(value, AreaSquareMillimeter)
}

// FromSquareMicrometer creates a new Area instance from SquareMicrometer.
func (uf AreaFactory) FromSquareMicrometers(value float64) (*Area, error) {
	return newArea(value, AreaSquareMicrometer)
}

// FromSquareMile creates a new Area instance from SquareMile.
func (uf AreaFactory) FromSquareMiles(value float64) (*Area, error) {
	return newArea(value, AreaSquareMile)
}

// FromSquareYard creates a new Area instance from SquareYard.
func (uf AreaFactory) FromSquareYards(value float64) (*Area, error) {
	return newArea(value, AreaSquareYard)
}

// FromSquareFoot creates a new Area instance from SquareFoot.
func (uf AreaFactory) FromSquareFeet(value float64) (*Area, error) {
	return newArea(value, AreaSquareFoot)
}

// FromUsSurveySquareFoot creates a new Area instance from UsSurveySquareFoot.
func (uf AreaFactory) FromUsSurveySquareFeet(value float64) (*Area, error) {
	return newArea(value, AreaUsSurveySquareFoot)
}

// FromSquareInch creates a new Area instance from SquareInch.
func (uf AreaFactory) FromSquareInches(value float64) (*Area, error) {
	return newArea(value, AreaSquareInch)
}

// FromAcre creates a new Area instance from Acre.
func (uf AreaFactory) FromAcres(value float64) (*Area, error) {
	return newArea(value, AreaAcre)
}

// FromHectare creates a new Area instance from Hectare.
func (uf AreaFactory) FromHectares(value float64) (*Area, error) {
	return newArea(value, AreaHectare)
}

// FromSquareNauticalMile creates a new Area instance from SquareNauticalMile.
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

// BaseValue returns the base value of Area in SquareMeter.
func (a *Area) BaseValue() float64 {
	return a.value
}


// SquareKilometer returns the value in SquareKilometer.
func (a *Area) SquareKilometers() float64 {
	if a.square_kilometersLazy != nil {
		return *a.square_kilometersLazy
	}
	square_kilometers := a.convertFromBase(AreaSquareKilometer)
	a.square_kilometersLazy = &square_kilometers
	return square_kilometers
}

// SquareMeter returns the value in SquareMeter.
func (a *Area) SquareMeters() float64 {
	if a.square_metersLazy != nil {
		return *a.square_metersLazy
	}
	square_meters := a.convertFromBase(AreaSquareMeter)
	a.square_metersLazy = &square_meters
	return square_meters
}

// SquareDecimeter returns the value in SquareDecimeter.
func (a *Area) SquareDecimeters() float64 {
	if a.square_decimetersLazy != nil {
		return *a.square_decimetersLazy
	}
	square_decimeters := a.convertFromBase(AreaSquareDecimeter)
	a.square_decimetersLazy = &square_decimeters
	return square_decimeters
}

// SquareCentimeter returns the value in SquareCentimeter.
func (a *Area) SquareCentimeters() float64 {
	if a.square_centimetersLazy != nil {
		return *a.square_centimetersLazy
	}
	square_centimeters := a.convertFromBase(AreaSquareCentimeter)
	a.square_centimetersLazy = &square_centimeters
	return square_centimeters
}

// SquareMillimeter returns the value in SquareMillimeter.
func (a *Area) SquareMillimeters() float64 {
	if a.square_millimetersLazy != nil {
		return *a.square_millimetersLazy
	}
	square_millimeters := a.convertFromBase(AreaSquareMillimeter)
	a.square_millimetersLazy = &square_millimeters
	return square_millimeters
}

// SquareMicrometer returns the value in SquareMicrometer.
func (a *Area) SquareMicrometers() float64 {
	if a.square_micrometersLazy != nil {
		return *a.square_micrometersLazy
	}
	square_micrometers := a.convertFromBase(AreaSquareMicrometer)
	a.square_micrometersLazy = &square_micrometers
	return square_micrometers
}

// SquareMile returns the value in SquareMile.
func (a *Area) SquareMiles() float64 {
	if a.square_milesLazy != nil {
		return *a.square_milesLazy
	}
	square_miles := a.convertFromBase(AreaSquareMile)
	a.square_milesLazy = &square_miles
	return square_miles
}

// SquareYard returns the value in SquareYard.
func (a *Area) SquareYards() float64 {
	if a.square_yardsLazy != nil {
		return *a.square_yardsLazy
	}
	square_yards := a.convertFromBase(AreaSquareYard)
	a.square_yardsLazy = &square_yards
	return square_yards
}

// SquareFoot returns the value in SquareFoot.
func (a *Area) SquareFeet() float64 {
	if a.square_feetLazy != nil {
		return *a.square_feetLazy
	}
	square_feet := a.convertFromBase(AreaSquareFoot)
	a.square_feetLazy = &square_feet
	return square_feet
}

// UsSurveySquareFoot returns the value in UsSurveySquareFoot.
func (a *Area) UsSurveySquareFeet() float64 {
	if a.us_survey_square_feetLazy != nil {
		return *a.us_survey_square_feetLazy
	}
	us_survey_square_feet := a.convertFromBase(AreaUsSurveySquareFoot)
	a.us_survey_square_feetLazy = &us_survey_square_feet
	return us_survey_square_feet
}

// SquareInch returns the value in SquareInch.
func (a *Area) SquareInches() float64 {
	if a.square_inchesLazy != nil {
		return *a.square_inchesLazy
	}
	square_inches := a.convertFromBase(AreaSquareInch)
	a.square_inchesLazy = &square_inches
	return square_inches
}

// Acre returns the value in Acre.
func (a *Area) Acres() float64 {
	if a.acresLazy != nil {
		return *a.acresLazy
	}
	acres := a.convertFromBase(AreaAcre)
	a.acresLazy = &acres
	return acres
}

// Hectare returns the value in Hectare.
func (a *Area) Hectares() float64 {
	if a.hectaresLazy != nil {
		return *a.hectaresLazy
	}
	hectares := a.convertFromBase(AreaHectare)
	a.hectaresLazy = &hectares
	return hectares
}

// SquareNauticalMile returns the value in SquareNauticalMile.
func (a *Area) SquareNauticalMiles() float64 {
	if a.square_nautical_milesLazy != nil {
		return *a.square_nautical_milesLazy
	}
	square_nautical_miles := a.convertFromBase(AreaSquareNauticalMile)
	a.square_nautical_milesLazy = &square_nautical_miles
	return square_nautical_miles
}


// ToDto creates an AreaDto representation.
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

// ToDtoJSON creates an AreaDto representation.
func (a *Area) ToDtoJSON(holdInUnit *AreaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Area to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Area) String() string {
	return a.ToString(AreaSquareMeter, 2)
}

// ToString formats the Area to string.
// fractionalDigits -1 for not mention
func (a *Area) ToString(unit AreaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Area) getUnitAbbreviation(unit AreaUnits) string {
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

// Check if the given Area are equals to the current Area
func (a *Area) Equals(other *Area) bool {
	return a.value == other.BaseValue()
}

// Check if the given Area are equals to the current Area
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

// Add the given Area to the current Area.
func (a *Area) Add(other *Area) *Area {
	return &Area{value: a.value + other.BaseValue()}
}

// Subtract the given Area to the current Area.
func (a *Area) Subtract(other *Area) *Area {
	return &Area{value: a.value - other.BaseValue()}
}

// Multiply the given Area to the current Area.
func (a *Area) Multiply(other *Area) *Area {
	return &Area{value: a.value * other.BaseValue()}
}

// Divide the given Area to the current Area.
func (a *Area) Divide(other *Area) *Area {
	return &Area{value: a.value / other.BaseValue()}
}