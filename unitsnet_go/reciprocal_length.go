package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ReciprocalLengthUnits enumeration
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

// ReciprocalLengthDto represents an ReciprocalLength
type ReciprocalLengthDto struct {
	Value float64
	Unit  ReciprocalLengthUnits
}

// ReciprocalLengthDtoFactory struct to group related functions
type ReciprocalLengthDtoFactory struct{}

func (udf ReciprocalLengthDtoFactory) FromJSON(data []byte) (*ReciprocalLengthDto, error) {
	a := ReciprocalLengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ReciprocalLengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ReciprocalLength struct
type ReciprocalLength struct {
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

// ReciprocalLengthFactory struct to group related functions
type ReciprocalLengthFactory struct{}

func (uf ReciprocalLengthFactory) CreateReciprocalLength(value float64, unit ReciprocalLengthUnits) (*ReciprocalLength, error) {
	return newReciprocalLength(value, unit)
}

func (uf ReciprocalLengthFactory) FromDto(dto ReciprocalLengthDto) (*ReciprocalLength, error) {
	return newReciprocalLength(dto.Value, dto.Unit)
}

func (uf ReciprocalLengthFactory) FromDtoJSON(data []byte) (*ReciprocalLength, error) {
	unitDto, err := ReciprocalLengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ReciprocalLengthFactory{}.FromDto(*unitDto)
}


// FromInverseMeter creates a new ReciprocalLength instance from InverseMeter.
func (uf ReciprocalLengthFactory) FromInverseMeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMeter)
}

// FromInverseCentimeter creates a new ReciprocalLength instance from InverseCentimeter.
func (uf ReciprocalLengthFactory) FromInverseCentimeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseCentimeter)
}

// FromInverseMillimeter creates a new ReciprocalLength instance from InverseMillimeter.
func (uf ReciprocalLengthFactory) FromInverseMillimeters(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMillimeter)
}

// FromInverseMile creates a new ReciprocalLength instance from InverseMile.
func (uf ReciprocalLengthFactory) FromInverseMiles(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMile)
}

// FromInverseYard creates a new ReciprocalLength instance from InverseYard.
func (uf ReciprocalLengthFactory) FromInverseYards(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseYard)
}

// FromInverseFoot creates a new ReciprocalLength instance from InverseFoot.
func (uf ReciprocalLengthFactory) FromInverseFeet(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseFoot)
}

// FromInverseUsSurveyFoot creates a new ReciprocalLength instance from InverseUsSurveyFoot.
func (uf ReciprocalLengthFactory) FromInverseUsSurveyFeet(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseUsSurveyFoot)
}

// FromInverseInch creates a new ReciprocalLength instance from InverseInch.
func (uf ReciprocalLengthFactory) FromInverseInches(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseInch)
}

// FromInverseMil creates a new ReciprocalLength instance from InverseMil.
func (uf ReciprocalLengthFactory) FromInverseMils(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMil)
}

// FromInverseMicroinch creates a new ReciprocalLength instance from InverseMicroinch.
func (uf ReciprocalLengthFactory) FromInverseMicroinches(value float64) (*ReciprocalLength, error) {
	return newReciprocalLength(value, ReciprocalLengthInverseMicroinch)
}




// newReciprocalLength creates a new ReciprocalLength.
func newReciprocalLength(value float64, fromUnit ReciprocalLengthUnits) (*ReciprocalLength, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ReciprocalLength{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ReciprocalLength in InverseMeter.
func (a *ReciprocalLength) BaseValue() float64 {
	return a.value
}


// InverseMeter returns the value in InverseMeter.
func (a *ReciprocalLength) InverseMeters() float64 {
	if a.inverse_metersLazy != nil {
		return *a.inverse_metersLazy
	}
	inverse_meters := a.convertFromBase(ReciprocalLengthInverseMeter)
	a.inverse_metersLazy = &inverse_meters
	return inverse_meters
}

// InverseCentimeter returns the value in InverseCentimeter.
func (a *ReciprocalLength) InverseCentimeters() float64 {
	if a.inverse_centimetersLazy != nil {
		return *a.inverse_centimetersLazy
	}
	inverse_centimeters := a.convertFromBase(ReciprocalLengthInverseCentimeter)
	a.inverse_centimetersLazy = &inverse_centimeters
	return inverse_centimeters
}

// InverseMillimeter returns the value in InverseMillimeter.
func (a *ReciprocalLength) InverseMillimeters() float64 {
	if a.inverse_millimetersLazy != nil {
		return *a.inverse_millimetersLazy
	}
	inverse_millimeters := a.convertFromBase(ReciprocalLengthInverseMillimeter)
	a.inverse_millimetersLazy = &inverse_millimeters
	return inverse_millimeters
}

// InverseMile returns the value in InverseMile.
func (a *ReciprocalLength) InverseMiles() float64 {
	if a.inverse_milesLazy != nil {
		return *a.inverse_milesLazy
	}
	inverse_miles := a.convertFromBase(ReciprocalLengthInverseMile)
	a.inverse_milesLazy = &inverse_miles
	return inverse_miles
}

// InverseYard returns the value in InverseYard.
func (a *ReciprocalLength) InverseYards() float64 {
	if a.inverse_yardsLazy != nil {
		return *a.inverse_yardsLazy
	}
	inverse_yards := a.convertFromBase(ReciprocalLengthInverseYard)
	a.inverse_yardsLazy = &inverse_yards
	return inverse_yards
}

// InverseFoot returns the value in InverseFoot.
func (a *ReciprocalLength) InverseFeet() float64 {
	if a.inverse_feetLazy != nil {
		return *a.inverse_feetLazy
	}
	inverse_feet := a.convertFromBase(ReciprocalLengthInverseFoot)
	a.inverse_feetLazy = &inverse_feet
	return inverse_feet
}

// InverseUsSurveyFoot returns the value in InverseUsSurveyFoot.
func (a *ReciprocalLength) InverseUsSurveyFeet() float64 {
	if a.inverse_us_survey_feetLazy != nil {
		return *a.inverse_us_survey_feetLazy
	}
	inverse_us_survey_feet := a.convertFromBase(ReciprocalLengthInverseUsSurveyFoot)
	a.inverse_us_survey_feetLazy = &inverse_us_survey_feet
	return inverse_us_survey_feet
}

// InverseInch returns the value in InverseInch.
func (a *ReciprocalLength) InverseInches() float64 {
	if a.inverse_inchesLazy != nil {
		return *a.inverse_inchesLazy
	}
	inverse_inches := a.convertFromBase(ReciprocalLengthInverseInch)
	a.inverse_inchesLazy = &inverse_inches
	return inverse_inches
}

// InverseMil returns the value in InverseMil.
func (a *ReciprocalLength) InverseMils() float64 {
	if a.inverse_milsLazy != nil {
		return *a.inverse_milsLazy
	}
	inverse_mils := a.convertFromBase(ReciprocalLengthInverseMil)
	a.inverse_milsLazy = &inverse_mils
	return inverse_mils
}

// InverseMicroinch returns the value in InverseMicroinch.
func (a *ReciprocalLength) InverseMicroinches() float64 {
	if a.inverse_microinchesLazy != nil {
		return *a.inverse_microinchesLazy
	}
	inverse_microinches := a.convertFromBase(ReciprocalLengthInverseMicroinch)
	a.inverse_microinchesLazy = &inverse_microinches
	return inverse_microinches
}


// ToDto creates an ReciprocalLengthDto representation.
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

// ToDtoJSON creates an ReciprocalLengthDto representation.
func (a *ReciprocalLength) ToDtoJSON(holdInUnit *ReciprocalLengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ReciprocalLength to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ReciprocalLength) String() string {
	return a.ToString(ReciprocalLengthInverseMeter, 2)
}

// ToString formats the ReciprocalLength to string.
// fractionalDigits -1 for not mention
func (a *ReciprocalLength) ToString(unit ReciprocalLengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ReciprocalLength) getUnitAbbreviation(unit ReciprocalLengthUnits) string {
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

// Check if the given ReciprocalLength are equals to the current ReciprocalLength
func (a *ReciprocalLength) Equals(other *ReciprocalLength) bool {
	return a.value == other.BaseValue()
}

// Check if the given ReciprocalLength are equals to the current ReciprocalLength
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

// Add the given ReciprocalLength to the current ReciprocalLength.
func (a *ReciprocalLength) Add(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value + other.BaseValue()}
}

// Subtract the given ReciprocalLength to the current ReciprocalLength.
func (a *ReciprocalLength) Subtract(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value - other.BaseValue()}
}

// Multiply the given ReciprocalLength to the current ReciprocalLength.
func (a *ReciprocalLength) Multiply(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value * other.BaseValue()}
}

// Divide the given ReciprocalLength to the current ReciprocalLength.
func (a *ReciprocalLength) Divide(other *ReciprocalLength) *ReciprocalLength {
	return &ReciprocalLength{value: a.value / other.BaseValue()}
}