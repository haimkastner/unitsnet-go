// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassFluxUnits defines various units of MassFlux.
type MassFluxUnits string

const (
    
        // 
        MassFluxGramPerSecondPerSquareMeter MassFluxUnits = "GramPerSecondPerSquareMeter"
        // 
        MassFluxGramPerSecondPerSquareCentimeter MassFluxUnits = "GramPerSecondPerSquareCentimeter"
        // 
        MassFluxGramPerSecondPerSquareMillimeter MassFluxUnits = "GramPerSecondPerSquareMillimeter"
        // 
        MassFluxGramPerHourPerSquareMeter MassFluxUnits = "GramPerHourPerSquareMeter"
        // 
        MassFluxGramPerHourPerSquareCentimeter MassFluxUnits = "GramPerHourPerSquareCentimeter"
        // 
        MassFluxGramPerHourPerSquareMillimeter MassFluxUnits = "GramPerHourPerSquareMillimeter"
        // 
        MassFluxKilogramPerSecondPerSquareMeter MassFluxUnits = "KilogramPerSecondPerSquareMeter"
        // 
        MassFluxKilogramPerSecondPerSquareCentimeter MassFluxUnits = "KilogramPerSecondPerSquareCentimeter"
        // 
        MassFluxKilogramPerSecondPerSquareMillimeter MassFluxUnits = "KilogramPerSecondPerSquareMillimeter"
        // 
        MassFluxKilogramPerHourPerSquareMeter MassFluxUnits = "KilogramPerHourPerSquareMeter"
        // 
        MassFluxKilogramPerHourPerSquareCentimeter MassFluxUnits = "KilogramPerHourPerSquareCentimeter"
        // 
        MassFluxKilogramPerHourPerSquareMillimeter MassFluxUnits = "KilogramPerHourPerSquareMillimeter"
)

// MassFluxDto represents a MassFlux measurement with a numerical value and its corresponding unit.
type MassFluxDto struct {
    // Value is the numerical representation of the MassFlux.
	Value float64
    // Unit specifies the unit of measurement for the MassFlux, as defined in the MassFluxUnits enumeration.
	Unit  MassFluxUnits
}

// MassFluxDtoFactory groups methods for creating and serializing MassFluxDto objects.
type MassFluxDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassFluxDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassFluxDtoFactory) FromJSON(data []byte) (*MassFluxDto, error) {
	a := MassFluxDto{}

    // Parse JSON into MassFluxDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MassFluxDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// MassFlux represents a measurement in a of MassFlux.
//
// Mass flux is the mass flow rate per unit area.
type MassFlux struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_second_per_square_meterLazy *float64 
    grams_per_second_per_square_centimeterLazy *float64 
    grams_per_second_per_square_millimeterLazy *float64 
    grams_per_hour_per_square_meterLazy *float64 
    grams_per_hour_per_square_centimeterLazy *float64 
    grams_per_hour_per_square_millimeterLazy *float64 
    kilograms_per_second_per_square_meterLazy *float64 
    kilograms_per_second_per_square_centimeterLazy *float64 
    kilograms_per_second_per_square_millimeterLazy *float64 
    kilograms_per_hour_per_square_meterLazy *float64 
    kilograms_per_hour_per_square_centimeterLazy *float64 
    kilograms_per_hour_per_square_millimeterLazy *float64 
}

// MassFluxFactory groups methods for creating MassFlux instances.
type MassFluxFactory struct{}

// CreateMassFlux creates a new MassFlux instance from the given value and unit.
func (uf MassFluxFactory) CreateMassFlux(value float64, unit MassFluxUnits) (*MassFlux, error) {
	return newMassFlux(value, unit)
}

// FromDto converts a MassFluxDto to a MassFlux instance.
func (uf MassFluxFactory) FromDto(dto MassFluxDto) (*MassFlux, error) {
	return newMassFlux(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MassFlux instance.
func (uf MassFluxFactory) FromDtoJSON(data []byte) (*MassFlux, error) {
	unitDto, err := MassFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassFluxDto from JSON: %w", err)
	}
	return MassFluxFactory{}.FromDto(*unitDto)
}


// FromGramsPerSecondPerSquareMeter creates a new MassFlux instance from a value in GramsPerSecondPerSquareMeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareMeter)
}

// FromGramsPerSecondPerSquareCentimeter creates a new MassFlux instance from a value in GramsPerSecondPerSquareCentimeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareCentimeter)
}

// FromGramsPerSecondPerSquareMillimeter creates a new MassFlux instance from a value in GramsPerSecondPerSquareMillimeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareMillimeter)
}

// FromGramsPerHourPerSquareMeter creates a new MassFlux instance from a value in GramsPerHourPerSquareMeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareMeter)
}

// FromGramsPerHourPerSquareCentimeter creates a new MassFlux instance from a value in GramsPerHourPerSquareCentimeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareCentimeter)
}

// FromGramsPerHourPerSquareMillimeter creates a new MassFlux instance from a value in GramsPerHourPerSquareMillimeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareMillimeter)
}

// FromKilogramsPerSecondPerSquareMeter creates a new MassFlux instance from a value in KilogramsPerSecondPerSquareMeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareMeter)
}

// FromKilogramsPerSecondPerSquareCentimeter creates a new MassFlux instance from a value in KilogramsPerSecondPerSquareCentimeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareCentimeter)
}

// FromKilogramsPerSecondPerSquareMillimeter creates a new MassFlux instance from a value in KilogramsPerSecondPerSquareMillimeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareMillimeter)
}

// FromKilogramsPerHourPerSquareMeter creates a new MassFlux instance from a value in KilogramsPerHourPerSquareMeter.
func (uf MassFluxFactory) FromKilogramsPerHourPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerHourPerSquareMeter)
}

// FromKilogramsPerHourPerSquareCentimeter creates a new MassFlux instance from a value in KilogramsPerHourPerSquareCentimeter.
func (uf MassFluxFactory) FromKilogramsPerHourPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerHourPerSquareCentimeter)
}

// FromKilogramsPerHourPerSquareMillimeter creates a new MassFlux instance from a value in KilogramsPerHourPerSquareMillimeter.
func (uf MassFluxFactory) FromKilogramsPerHourPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerHourPerSquareMillimeter)
}


// newMassFlux creates a new MassFlux.
func newMassFlux(value float64, fromUnit MassFluxUnits) (*MassFlux, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MassFlux{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MassFlux in KilogramPerSecondPerSquareMeter unit (the base/default quantity).
func (a *MassFlux) BaseValue() float64 {
	return a.value
}


// GramsPerSecondPerSquareMeter returns the MassFlux value in GramsPerSecondPerSquareMeter.
//
// 
func (a *MassFlux) GramsPerSecondPerSquareMeter() float64 {
	if a.grams_per_second_per_square_meterLazy != nil {
		return *a.grams_per_second_per_square_meterLazy
	}
	grams_per_second_per_square_meter := a.convertFromBase(MassFluxGramPerSecondPerSquareMeter)
	a.grams_per_second_per_square_meterLazy = &grams_per_second_per_square_meter
	return grams_per_second_per_square_meter
}

// GramsPerSecondPerSquareCentimeter returns the MassFlux value in GramsPerSecondPerSquareCentimeter.
//
// 
func (a *MassFlux) GramsPerSecondPerSquareCentimeter() float64 {
	if a.grams_per_second_per_square_centimeterLazy != nil {
		return *a.grams_per_second_per_square_centimeterLazy
	}
	grams_per_second_per_square_centimeter := a.convertFromBase(MassFluxGramPerSecondPerSquareCentimeter)
	a.grams_per_second_per_square_centimeterLazy = &grams_per_second_per_square_centimeter
	return grams_per_second_per_square_centimeter
}

// GramsPerSecondPerSquareMillimeter returns the MassFlux value in GramsPerSecondPerSquareMillimeter.
//
// 
func (a *MassFlux) GramsPerSecondPerSquareMillimeter() float64 {
	if a.grams_per_second_per_square_millimeterLazy != nil {
		return *a.grams_per_second_per_square_millimeterLazy
	}
	grams_per_second_per_square_millimeter := a.convertFromBase(MassFluxGramPerSecondPerSquareMillimeter)
	a.grams_per_second_per_square_millimeterLazy = &grams_per_second_per_square_millimeter
	return grams_per_second_per_square_millimeter
}

// GramsPerHourPerSquareMeter returns the MassFlux value in GramsPerHourPerSquareMeter.
//
// 
func (a *MassFlux) GramsPerHourPerSquareMeter() float64 {
	if a.grams_per_hour_per_square_meterLazy != nil {
		return *a.grams_per_hour_per_square_meterLazy
	}
	grams_per_hour_per_square_meter := a.convertFromBase(MassFluxGramPerHourPerSquareMeter)
	a.grams_per_hour_per_square_meterLazy = &grams_per_hour_per_square_meter
	return grams_per_hour_per_square_meter
}

// GramsPerHourPerSquareCentimeter returns the MassFlux value in GramsPerHourPerSquareCentimeter.
//
// 
func (a *MassFlux) GramsPerHourPerSquareCentimeter() float64 {
	if a.grams_per_hour_per_square_centimeterLazy != nil {
		return *a.grams_per_hour_per_square_centimeterLazy
	}
	grams_per_hour_per_square_centimeter := a.convertFromBase(MassFluxGramPerHourPerSquareCentimeter)
	a.grams_per_hour_per_square_centimeterLazy = &grams_per_hour_per_square_centimeter
	return grams_per_hour_per_square_centimeter
}

// GramsPerHourPerSquareMillimeter returns the MassFlux value in GramsPerHourPerSquareMillimeter.
//
// 
func (a *MassFlux) GramsPerHourPerSquareMillimeter() float64 {
	if a.grams_per_hour_per_square_millimeterLazy != nil {
		return *a.grams_per_hour_per_square_millimeterLazy
	}
	grams_per_hour_per_square_millimeter := a.convertFromBase(MassFluxGramPerHourPerSquareMillimeter)
	a.grams_per_hour_per_square_millimeterLazy = &grams_per_hour_per_square_millimeter
	return grams_per_hour_per_square_millimeter
}

// KilogramsPerSecondPerSquareMeter returns the MassFlux value in KilogramsPerSecondPerSquareMeter.
//
// 
func (a *MassFlux) KilogramsPerSecondPerSquareMeter() float64 {
	if a.kilograms_per_second_per_square_meterLazy != nil {
		return *a.kilograms_per_second_per_square_meterLazy
	}
	kilograms_per_second_per_square_meter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareMeter)
	a.kilograms_per_second_per_square_meterLazy = &kilograms_per_second_per_square_meter
	return kilograms_per_second_per_square_meter
}

// KilogramsPerSecondPerSquareCentimeter returns the MassFlux value in KilogramsPerSecondPerSquareCentimeter.
//
// 
func (a *MassFlux) KilogramsPerSecondPerSquareCentimeter() float64 {
	if a.kilograms_per_second_per_square_centimeterLazy != nil {
		return *a.kilograms_per_second_per_square_centimeterLazy
	}
	kilograms_per_second_per_square_centimeter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareCentimeter)
	a.kilograms_per_second_per_square_centimeterLazy = &kilograms_per_second_per_square_centimeter
	return kilograms_per_second_per_square_centimeter
}

// KilogramsPerSecondPerSquareMillimeter returns the MassFlux value in KilogramsPerSecondPerSquareMillimeter.
//
// 
func (a *MassFlux) KilogramsPerSecondPerSquareMillimeter() float64 {
	if a.kilograms_per_second_per_square_millimeterLazy != nil {
		return *a.kilograms_per_second_per_square_millimeterLazy
	}
	kilograms_per_second_per_square_millimeter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareMillimeter)
	a.kilograms_per_second_per_square_millimeterLazy = &kilograms_per_second_per_square_millimeter
	return kilograms_per_second_per_square_millimeter
}

// KilogramsPerHourPerSquareMeter returns the MassFlux value in KilogramsPerHourPerSquareMeter.
//
// 
func (a *MassFlux) KilogramsPerHourPerSquareMeter() float64 {
	if a.kilograms_per_hour_per_square_meterLazy != nil {
		return *a.kilograms_per_hour_per_square_meterLazy
	}
	kilograms_per_hour_per_square_meter := a.convertFromBase(MassFluxKilogramPerHourPerSquareMeter)
	a.kilograms_per_hour_per_square_meterLazy = &kilograms_per_hour_per_square_meter
	return kilograms_per_hour_per_square_meter
}

// KilogramsPerHourPerSquareCentimeter returns the MassFlux value in KilogramsPerHourPerSquareCentimeter.
//
// 
func (a *MassFlux) KilogramsPerHourPerSquareCentimeter() float64 {
	if a.kilograms_per_hour_per_square_centimeterLazy != nil {
		return *a.kilograms_per_hour_per_square_centimeterLazy
	}
	kilograms_per_hour_per_square_centimeter := a.convertFromBase(MassFluxKilogramPerHourPerSquareCentimeter)
	a.kilograms_per_hour_per_square_centimeterLazy = &kilograms_per_hour_per_square_centimeter
	return kilograms_per_hour_per_square_centimeter
}

// KilogramsPerHourPerSquareMillimeter returns the MassFlux value in KilogramsPerHourPerSquareMillimeter.
//
// 
func (a *MassFlux) KilogramsPerHourPerSquareMillimeter() float64 {
	if a.kilograms_per_hour_per_square_millimeterLazy != nil {
		return *a.kilograms_per_hour_per_square_millimeterLazy
	}
	kilograms_per_hour_per_square_millimeter := a.convertFromBase(MassFluxKilogramPerHourPerSquareMillimeter)
	a.kilograms_per_hour_per_square_millimeterLazy = &kilograms_per_hour_per_square_millimeter
	return kilograms_per_hour_per_square_millimeter
}


// ToDto creates a MassFluxDto representation from the MassFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerSecondPerSquareMeter by default.
func (a *MassFlux) ToDto(holdInUnit *MassFluxUnits) MassFluxDto {
	if holdInUnit == nil {
		defaultUnit := MassFluxKilogramPerSecondPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MassFluxDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the MassFlux instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerSecondPerSquareMeter by default.
func (a *MassFlux) ToDtoJSON(holdInUnit *MassFluxUnits) ([]byte, error) {
	// Convert to MassFluxDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MassFlux to a specific unit value.
// The function uses the provided unit type (MassFluxUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *MassFlux) Convert(toUnit MassFluxUnits) float64 {
	switch toUnit { 
    case MassFluxGramPerSecondPerSquareMeter:
		return a.GramsPerSecondPerSquareMeter()
    case MassFluxGramPerSecondPerSquareCentimeter:
		return a.GramsPerSecondPerSquareCentimeter()
    case MassFluxGramPerSecondPerSquareMillimeter:
		return a.GramsPerSecondPerSquareMillimeter()
    case MassFluxGramPerHourPerSquareMeter:
		return a.GramsPerHourPerSquareMeter()
    case MassFluxGramPerHourPerSquareCentimeter:
		return a.GramsPerHourPerSquareCentimeter()
    case MassFluxGramPerHourPerSquareMillimeter:
		return a.GramsPerHourPerSquareMillimeter()
    case MassFluxKilogramPerSecondPerSquareMeter:
		return a.KilogramsPerSecondPerSquareMeter()
    case MassFluxKilogramPerSecondPerSquareCentimeter:
		return a.KilogramsPerSecondPerSquareCentimeter()
    case MassFluxKilogramPerSecondPerSquareMillimeter:
		return a.KilogramsPerSecondPerSquareMillimeter()
    case MassFluxKilogramPerHourPerSquareMeter:
		return a.KilogramsPerHourPerSquareMeter()
    case MassFluxKilogramPerHourPerSquareCentimeter:
		return a.KilogramsPerHourPerSquareCentimeter()
    case MassFluxKilogramPerHourPerSquareMillimeter:
		return a.KilogramsPerHourPerSquareMillimeter()
	default:
		return math.NaN()
	}
}

func (a *MassFlux) convertFromBase(toUnit MassFluxUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassFluxGramPerSecondPerSquareMeter:
		return (value * 1e3) 
	case MassFluxGramPerSecondPerSquareCentimeter:
		return (value * 1e-1) 
	case MassFluxGramPerSecondPerSquareMillimeter:
		return (value * 1e-3) 
	case MassFluxGramPerHourPerSquareMeter:
		return (value * 3.6e6) 
	case MassFluxGramPerHourPerSquareCentimeter:
		return (value * 3.6e2) 
	case MassFluxGramPerHourPerSquareMillimeter:
		return (value * 3.6e0) 
	case MassFluxKilogramPerSecondPerSquareMeter:
		return ((value * 1e3) / 1000.0) 
	case MassFluxKilogramPerSecondPerSquareCentimeter:
		return ((value * 1e-1) / 1000.0) 
	case MassFluxKilogramPerSecondPerSquareMillimeter:
		return ((value * 1e-3) / 1000.0) 
	case MassFluxKilogramPerHourPerSquareMeter:
		return ((value * 3.6e6) / 1000.0) 
	case MassFluxKilogramPerHourPerSquareCentimeter:
		return ((value * 3.6e2) / 1000.0) 
	case MassFluxKilogramPerHourPerSquareMillimeter:
		return ((value * 3.6e0) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *MassFlux) convertToBase(value float64, fromUnit MassFluxUnits) float64 {
	switch fromUnit { 
	case MassFluxGramPerSecondPerSquareMeter:
		return (value / 1e3) 
	case MassFluxGramPerSecondPerSquareCentimeter:
		return (value / 1e-1) 
	case MassFluxGramPerSecondPerSquareMillimeter:
		return (value / 1e-3) 
	case MassFluxGramPerHourPerSquareMeter:
		return (value / 3.6e6) 
	case MassFluxGramPerHourPerSquareCentimeter:
		return (value / 3.6e2) 
	case MassFluxGramPerHourPerSquareMillimeter:
		return (value / 3.6e0) 
	case MassFluxKilogramPerSecondPerSquareMeter:
		return ((value / 1e3) * 1000.0) 
	case MassFluxKilogramPerSecondPerSquareCentimeter:
		return ((value / 1e-1) * 1000.0) 
	case MassFluxKilogramPerSecondPerSquareMillimeter:
		return ((value / 1e-3) * 1000.0) 
	case MassFluxKilogramPerHourPerSquareMeter:
		return ((value / 3.6e6) * 1000.0) 
	case MassFluxKilogramPerHourPerSquareCentimeter:
		return ((value / 3.6e2) * 1000.0) 
	case MassFluxKilogramPerHourPerSquareMillimeter:
		return ((value / 3.6e0) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the MassFlux in the default unit (KilogramPerSecondPerSquareMeter),
// formatted to two decimal places.
func (a MassFlux) String() string {
	return a.ToString(MassFluxKilogramPerSecondPerSquareMeter, 2)
}

// ToString formats the MassFlux value as a string with the specified unit and fractional digits.
// It converts the MassFlux to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MassFlux value will be converted (e.g., KilogramPerSecondPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MassFlux with the unit abbreviation.
func (a *MassFlux) ToString(unit MassFluxUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MassFlux) getUnitAbbreviation(unit MassFluxUnits) string {
	switch unit { 
	case MassFluxGramPerSecondPerSquareMeter:
		return "g·s⁻¹·m⁻²" 
	case MassFluxGramPerSecondPerSquareCentimeter:
		return "g·s⁻¹·cm⁻²" 
	case MassFluxGramPerSecondPerSquareMillimeter:
		return "g·s⁻¹·mm⁻²" 
	case MassFluxGramPerHourPerSquareMeter:
		return "g·h⁻¹·m⁻²" 
	case MassFluxGramPerHourPerSquareCentimeter:
		return "g·h⁻¹·cm⁻²" 
	case MassFluxGramPerHourPerSquareMillimeter:
		return "g·h⁻¹·mm⁻²" 
	case MassFluxKilogramPerSecondPerSquareMeter:
		return "kg·s⁻¹·m⁻²" 
	case MassFluxKilogramPerSecondPerSquareCentimeter:
		return "kg·s⁻¹·cm⁻²" 
	case MassFluxKilogramPerSecondPerSquareMillimeter:
		return "kg·s⁻¹·mm⁻²" 
	case MassFluxKilogramPerHourPerSquareMeter:
		return "kg·h⁻¹·m⁻²" 
	case MassFluxKilogramPerHourPerSquareCentimeter:
		return "kg·h⁻¹·cm⁻²" 
	case MassFluxKilogramPerHourPerSquareMillimeter:
		return "kg·h⁻¹·mm⁻²" 
	default:
		return ""
	}
}

// Equals checks if the given MassFlux is equal to the current MassFlux.
//
// Parameters:
//    other: The MassFlux to compare against.
//
// Returns:
//    bool: Returns true if both MassFlux are equal, false otherwise.
func (a *MassFlux) Equals(other *MassFlux) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MassFlux with another MassFlux.
// It returns -1 if the current MassFlux is less than the other MassFlux, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MassFlux to compare against.
//
// Returns:
//    int: -1 if the current MassFlux is less, 1 if greater, and 0 if equal.
func (a *MassFlux) CompareTo(other *MassFlux) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given MassFlux to the current MassFlux and returns the result.
// The result is a new MassFlux instance with the sum of the values.
//
// Parameters:
//    other: The MassFlux to add to the current MassFlux.
//
// Returns:
//    *MassFlux: A new MassFlux instance representing the sum of both MassFlux.
func (a *MassFlux) Add(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MassFlux from the current MassFlux and returns the result.
// The result is a new MassFlux instance with the difference of the values.
//
// Parameters:
//    other: The MassFlux to subtract from the current MassFlux.
//
// Returns:
//    *MassFlux: A new MassFlux instance representing the difference of both MassFlux.
func (a *MassFlux) Subtract(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MassFlux by the given MassFlux and returns the result.
// The result is a new MassFlux instance with the product of the values.
//
// Parameters:
//    other: The MassFlux to multiply with the current MassFlux.
//
// Returns:
//    *MassFlux: A new MassFlux instance representing the product of both MassFlux.
func (a *MassFlux) Multiply(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value * other.BaseValue()}
}

// Divide divides the current MassFlux by the given MassFlux and returns the result.
// The result is a new MassFlux instance with the quotient of the values.
//
// Parameters:
//    other: The MassFlux to divide the current MassFlux by.
//
// Returns:
//    *MassFlux: A new MassFlux instance representing the quotient of both MassFlux.
func (a *MassFlux) Divide(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value / other.BaseValue()}
}