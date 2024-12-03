package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ReciprocalAreaUnits enumeration
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

// ReciprocalAreaDto represents an ReciprocalArea
type ReciprocalAreaDto struct {
	Value float64
	Unit  ReciprocalAreaUnits
}

// ReciprocalAreaDtoFactory struct to group related functions
type ReciprocalAreaDtoFactory struct{}

func (udf ReciprocalAreaDtoFactory) FromJSON(data []byte) (*ReciprocalAreaDto, error) {
	a := ReciprocalAreaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ReciprocalAreaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ReciprocalArea struct
type ReciprocalArea struct {
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

// ReciprocalAreaFactory struct to group related functions
type ReciprocalAreaFactory struct{}

func (uf ReciprocalAreaFactory) CreateReciprocalArea(value float64, unit ReciprocalAreaUnits) (*ReciprocalArea, error) {
	return newReciprocalArea(value, unit)
}

func (uf ReciprocalAreaFactory) FromDto(dto ReciprocalAreaDto) (*ReciprocalArea, error) {
	return newReciprocalArea(dto.Value, dto.Unit)
}

func (uf ReciprocalAreaFactory) FromDtoJSON(data []byte) (*ReciprocalArea, error) {
	unitDto, err := ReciprocalAreaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ReciprocalAreaFactory{}.FromDto(*unitDto)
}


// FromInverseSquareMeter creates a new ReciprocalArea instance from InverseSquareMeter.
func (uf ReciprocalAreaFactory) FromInverseSquareMeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMeter)
}

// FromInverseSquareKilometer creates a new ReciprocalArea instance from InverseSquareKilometer.
func (uf ReciprocalAreaFactory) FromInverseSquareKilometers(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareKilometer)
}

// FromInverseSquareDecimeter creates a new ReciprocalArea instance from InverseSquareDecimeter.
func (uf ReciprocalAreaFactory) FromInverseSquareDecimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareDecimeter)
}

// FromInverseSquareCentimeter creates a new ReciprocalArea instance from InverseSquareCentimeter.
func (uf ReciprocalAreaFactory) FromInverseSquareCentimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareCentimeter)
}

// FromInverseSquareMillimeter creates a new ReciprocalArea instance from InverseSquareMillimeter.
func (uf ReciprocalAreaFactory) FromInverseSquareMillimeters(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMillimeter)
}

// FromInverseSquareMicrometer creates a new ReciprocalArea instance from InverseSquareMicrometer.
func (uf ReciprocalAreaFactory) FromInverseSquareMicrometers(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMicrometer)
}

// FromInverseSquareMile creates a new ReciprocalArea instance from InverseSquareMile.
func (uf ReciprocalAreaFactory) FromInverseSquareMiles(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareMile)
}

// FromInverseSquareYard creates a new ReciprocalArea instance from InverseSquareYard.
func (uf ReciprocalAreaFactory) FromInverseSquareYards(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareYard)
}

// FromInverseSquareFoot creates a new ReciprocalArea instance from InverseSquareFoot.
func (uf ReciprocalAreaFactory) FromInverseSquareFeet(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseSquareFoot)
}

// FromInverseUsSurveySquareFoot creates a new ReciprocalArea instance from InverseUsSurveySquareFoot.
func (uf ReciprocalAreaFactory) FromInverseUsSurveySquareFeet(value float64) (*ReciprocalArea, error) {
	return newReciprocalArea(value, ReciprocalAreaInverseUsSurveySquareFoot)
}

// FromInverseSquareInch creates a new ReciprocalArea instance from InverseSquareInch.
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

// BaseValue returns the base value of ReciprocalArea in InverseSquareMeter.
func (a *ReciprocalArea) BaseValue() float64 {
	return a.value
}


// InverseSquareMeter returns the value in InverseSquareMeter.
func (a *ReciprocalArea) InverseSquareMeters() float64 {
	if a.inverse_square_metersLazy != nil {
		return *a.inverse_square_metersLazy
	}
	inverse_square_meters := a.convertFromBase(ReciprocalAreaInverseSquareMeter)
	a.inverse_square_metersLazy = &inverse_square_meters
	return inverse_square_meters
}

// InverseSquareKilometer returns the value in InverseSquareKilometer.
func (a *ReciprocalArea) InverseSquareKilometers() float64 {
	if a.inverse_square_kilometersLazy != nil {
		return *a.inverse_square_kilometersLazy
	}
	inverse_square_kilometers := a.convertFromBase(ReciprocalAreaInverseSquareKilometer)
	a.inverse_square_kilometersLazy = &inverse_square_kilometers
	return inverse_square_kilometers
}

// InverseSquareDecimeter returns the value in InverseSquareDecimeter.
func (a *ReciprocalArea) InverseSquareDecimeters() float64 {
	if a.inverse_square_decimetersLazy != nil {
		return *a.inverse_square_decimetersLazy
	}
	inverse_square_decimeters := a.convertFromBase(ReciprocalAreaInverseSquareDecimeter)
	a.inverse_square_decimetersLazy = &inverse_square_decimeters
	return inverse_square_decimeters
}

// InverseSquareCentimeter returns the value in InverseSquareCentimeter.
func (a *ReciprocalArea) InverseSquareCentimeters() float64 {
	if a.inverse_square_centimetersLazy != nil {
		return *a.inverse_square_centimetersLazy
	}
	inverse_square_centimeters := a.convertFromBase(ReciprocalAreaInverseSquareCentimeter)
	a.inverse_square_centimetersLazy = &inverse_square_centimeters
	return inverse_square_centimeters
}

// InverseSquareMillimeter returns the value in InverseSquareMillimeter.
func (a *ReciprocalArea) InverseSquareMillimeters() float64 {
	if a.inverse_square_millimetersLazy != nil {
		return *a.inverse_square_millimetersLazy
	}
	inverse_square_millimeters := a.convertFromBase(ReciprocalAreaInverseSquareMillimeter)
	a.inverse_square_millimetersLazy = &inverse_square_millimeters
	return inverse_square_millimeters
}

// InverseSquareMicrometer returns the value in InverseSquareMicrometer.
func (a *ReciprocalArea) InverseSquareMicrometers() float64 {
	if a.inverse_square_micrometersLazy != nil {
		return *a.inverse_square_micrometersLazy
	}
	inverse_square_micrometers := a.convertFromBase(ReciprocalAreaInverseSquareMicrometer)
	a.inverse_square_micrometersLazy = &inverse_square_micrometers
	return inverse_square_micrometers
}

// InverseSquareMile returns the value in InverseSquareMile.
func (a *ReciprocalArea) InverseSquareMiles() float64 {
	if a.inverse_square_milesLazy != nil {
		return *a.inverse_square_milesLazy
	}
	inverse_square_miles := a.convertFromBase(ReciprocalAreaInverseSquareMile)
	a.inverse_square_milesLazy = &inverse_square_miles
	return inverse_square_miles
}

// InverseSquareYard returns the value in InverseSquareYard.
func (a *ReciprocalArea) InverseSquareYards() float64 {
	if a.inverse_square_yardsLazy != nil {
		return *a.inverse_square_yardsLazy
	}
	inverse_square_yards := a.convertFromBase(ReciprocalAreaInverseSquareYard)
	a.inverse_square_yardsLazy = &inverse_square_yards
	return inverse_square_yards
}

// InverseSquareFoot returns the value in InverseSquareFoot.
func (a *ReciprocalArea) InverseSquareFeet() float64 {
	if a.inverse_square_feetLazy != nil {
		return *a.inverse_square_feetLazy
	}
	inverse_square_feet := a.convertFromBase(ReciprocalAreaInverseSquareFoot)
	a.inverse_square_feetLazy = &inverse_square_feet
	return inverse_square_feet
}

// InverseUsSurveySquareFoot returns the value in InverseUsSurveySquareFoot.
func (a *ReciprocalArea) InverseUsSurveySquareFeet() float64 {
	if a.inverse_us_survey_square_feetLazy != nil {
		return *a.inverse_us_survey_square_feetLazy
	}
	inverse_us_survey_square_feet := a.convertFromBase(ReciprocalAreaInverseUsSurveySquareFoot)
	a.inverse_us_survey_square_feetLazy = &inverse_us_survey_square_feet
	return inverse_us_survey_square_feet
}

// InverseSquareInch returns the value in InverseSquareInch.
func (a *ReciprocalArea) InverseSquareInches() float64 {
	if a.inverse_square_inchesLazy != nil {
		return *a.inverse_square_inchesLazy
	}
	inverse_square_inches := a.convertFromBase(ReciprocalAreaInverseSquareInch)
	a.inverse_square_inchesLazy = &inverse_square_inches
	return inverse_square_inches
}


// ToDto creates an ReciprocalAreaDto representation.
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

// ToDtoJSON creates an ReciprocalAreaDto representation.
func (a *ReciprocalArea) ToDtoJSON(holdInUnit *ReciprocalAreaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ReciprocalArea to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ReciprocalArea) String() string {
	return a.ToString(ReciprocalAreaInverseSquareMeter, 2)
}

// ToString formats the ReciprocalArea to string.
// fractionalDigits -1 for not mention
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

// Check if the given ReciprocalArea are equals to the current ReciprocalArea
func (a *ReciprocalArea) Equals(other *ReciprocalArea) bool {
	return a.value == other.BaseValue()
}

// Check if the given ReciprocalArea are equals to the current ReciprocalArea
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

// Add the given ReciprocalArea to the current ReciprocalArea.
func (a *ReciprocalArea) Add(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value + other.BaseValue()}
}

// Subtract the given ReciprocalArea to the current ReciprocalArea.
func (a *ReciprocalArea) Subtract(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value - other.BaseValue()}
}

// Multiply the given ReciprocalArea to the current ReciprocalArea.
func (a *ReciprocalArea) Multiply(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value * other.BaseValue()}
}

// Divide the given ReciprocalArea to the current ReciprocalArea.
func (a *ReciprocalArea) Divide(other *ReciprocalArea) *ReciprocalArea {
	return &ReciprocalArea{value: a.value / other.BaseValue()}
}