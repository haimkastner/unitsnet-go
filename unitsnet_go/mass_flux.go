package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassFluxUnits enumeration
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

// MassFluxDto represents an MassFlux
type MassFluxDto struct {
	Value float64
	Unit  MassFluxUnits
}

// MassFluxDtoFactory struct to group related functions
type MassFluxDtoFactory struct{}

func (udf MassFluxDtoFactory) FromJSON(data []byte) (*MassFluxDto, error) {
	a := MassFluxDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassFluxDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MassFlux struct
type MassFlux struct {
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

// MassFluxFactory struct to group related functions
type MassFluxFactory struct{}

func (uf MassFluxFactory) CreateMassFlux(value float64, unit MassFluxUnits) (*MassFlux, error) {
	return newMassFlux(value, unit)
}

func (uf MassFluxFactory) FromDto(dto MassFluxDto) (*MassFlux, error) {
	return newMassFlux(dto.Value, dto.Unit)
}

func (uf MassFluxFactory) FromDtoJSON(data []byte) (*MassFlux, error) {
	unitDto, err := MassFluxDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassFluxFactory{}.FromDto(*unitDto)
}


// FromGramPerSecondPerSquareMeter creates a new MassFlux instance from GramPerSecondPerSquareMeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareMeter)
}

// FromGramPerSecondPerSquareCentimeter creates a new MassFlux instance from GramPerSecondPerSquareCentimeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareCentimeter)
}

// FromGramPerSecondPerSquareMillimeter creates a new MassFlux instance from GramPerSecondPerSquareMillimeter.
func (uf MassFluxFactory) FromGramsPerSecondPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerSecondPerSquareMillimeter)
}

// FromGramPerHourPerSquareMeter creates a new MassFlux instance from GramPerHourPerSquareMeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareMeter)
}

// FromGramPerHourPerSquareCentimeter creates a new MassFlux instance from GramPerHourPerSquareCentimeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareCentimeter)
}

// FromGramPerHourPerSquareMillimeter creates a new MassFlux instance from GramPerHourPerSquareMillimeter.
func (uf MassFluxFactory) FromGramsPerHourPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxGramPerHourPerSquareMillimeter)
}

// FromKilogramPerSecondPerSquareMeter creates a new MassFlux instance from KilogramPerSecondPerSquareMeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareMeter)
}

// FromKilogramPerSecondPerSquareCentimeter creates a new MassFlux instance from KilogramPerSecondPerSquareCentimeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareCentimeter)
}

// FromKilogramPerSecondPerSquareMillimeter creates a new MassFlux instance from KilogramPerSecondPerSquareMillimeter.
func (uf MassFluxFactory) FromKilogramsPerSecondPerSquareMillimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerSecondPerSquareMillimeter)
}

// FromKilogramPerHourPerSquareMeter creates a new MassFlux instance from KilogramPerHourPerSquareMeter.
func (uf MassFluxFactory) FromKilogramsPerHourPerSquareMeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerHourPerSquareMeter)
}

// FromKilogramPerHourPerSquareCentimeter creates a new MassFlux instance from KilogramPerHourPerSquareCentimeter.
func (uf MassFluxFactory) FromKilogramsPerHourPerSquareCentimeter(value float64) (*MassFlux, error) {
	return newMassFlux(value, MassFluxKilogramPerHourPerSquareCentimeter)
}

// FromKilogramPerHourPerSquareMillimeter creates a new MassFlux instance from KilogramPerHourPerSquareMillimeter.
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

// BaseValue returns the base value of MassFlux in KilogramPerSecondPerSquareMeter.
func (a *MassFlux) BaseValue() float64 {
	return a.value
}


// GramPerSecondPerSquareMeter returns the value in GramPerSecondPerSquareMeter.
func (a *MassFlux) GramsPerSecondPerSquareMeter() float64 {
	if a.grams_per_second_per_square_meterLazy != nil {
		return *a.grams_per_second_per_square_meterLazy
	}
	grams_per_second_per_square_meter := a.convertFromBase(MassFluxGramPerSecondPerSquareMeter)
	a.grams_per_second_per_square_meterLazy = &grams_per_second_per_square_meter
	return grams_per_second_per_square_meter
}

// GramPerSecondPerSquareCentimeter returns the value in GramPerSecondPerSquareCentimeter.
func (a *MassFlux) GramsPerSecondPerSquareCentimeter() float64 {
	if a.grams_per_second_per_square_centimeterLazy != nil {
		return *a.grams_per_second_per_square_centimeterLazy
	}
	grams_per_second_per_square_centimeter := a.convertFromBase(MassFluxGramPerSecondPerSquareCentimeter)
	a.grams_per_second_per_square_centimeterLazy = &grams_per_second_per_square_centimeter
	return grams_per_second_per_square_centimeter
}

// GramPerSecondPerSquareMillimeter returns the value in GramPerSecondPerSquareMillimeter.
func (a *MassFlux) GramsPerSecondPerSquareMillimeter() float64 {
	if a.grams_per_second_per_square_millimeterLazy != nil {
		return *a.grams_per_second_per_square_millimeterLazy
	}
	grams_per_second_per_square_millimeter := a.convertFromBase(MassFluxGramPerSecondPerSquareMillimeter)
	a.grams_per_second_per_square_millimeterLazy = &grams_per_second_per_square_millimeter
	return grams_per_second_per_square_millimeter
}

// GramPerHourPerSquareMeter returns the value in GramPerHourPerSquareMeter.
func (a *MassFlux) GramsPerHourPerSquareMeter() float64 {
	if a.grams_per_hour_per_square_meterLazy != nil {
		return *a.grams_per_hour_per_square_meterLazy
	}
	grams_per_hour_per_square_meter := a.convertFromBase(MassFluxGramPerHourPerSquareMeter)
	a.grams_per_hour_per_square_meterLazy = &grams_per_hour_per_square_meter
	return grams_per_hour_per_square_meter
}

// GramPerHourPerSquareCentimeter returns the value in GramPerHourPerSquareCentimeter.
func (a *MassFlux) GramsPerHourPerSquareCentimeter() float64 {
	if a.grams_per_hour_per_square_centimeterLazy != nil {
		return *a.grams_per_hour_per_square_centimeterLazy
	}
	grams_per_hour_per_square_centimeter := a.convertFromBase(MassFluxGramPerHourPerSquareCentimeter)
	a.grams_per_hour_per_square_centimeterLazy = &grams_per_hour_per_square_centimeter
	return grams_per_hour_per_square_centimeter
}

// GramPerHourPerSquareMillimeter returns the value in GramPerHourPerSquareMillimeter.
func (a *MassFlux) GramsPerHourPerSquareMillimeter() float64 {
	if a.grams_per_hour_per_square_millimeterLazy != nil {
		return *a.grams_per_hour_per_square_millimeterLazy
	}
	grams_per_hour_per_square_millimeter := a.convertFromBase(MassFluxGramPerHourPerSquareMillimeter)
	a.grams_per_hour_per_square_millimeterLazy = &grams_per_hour_per_square_millimeter
	return grams_per_hour_per_square_millimeter
}

// KilogramPerSecondPerSquareMeter returns the value in KilogramPerSecondPerSquareMeter.
func (a *MassFlux) KilogramsPerSecondPerSquareMeter() float64 {
	if a.kilograms_per_second_per_square_meterLazy != nil {
		return *a.kilograms_per_second_per_square_meterLazy
	}
	kilograms_per_second_per_square_meter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareMeter)
	a.kilograms_per_second_per_square_meterLazy = &kilograms_per_second_per_square_meter
	return kilograms_per_second_per_square_meter
}

// KilogramPerSecondPerSquareCentimeter returns the value in KilogramPerSecondPerSquareCentimeter.
func (a *MassFlux) KilogramsPerSecondPerSquareCentimeter() float64 {
	if a.kilograms_per_second_per_square_centimeterLazy != nil {
		return *a.kilograms_per_second_per_square_centimeterLazy
	}
	kilograms_per_second_per_square_centimeter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareCentimeter)
	a.kilograms_per_second_per_square_centimeterLazy = &kilograms_per_second_per_square_centimeter
	return kilograms_per_second_per_square_centimeter
}

// KilogramPerSecondPerSquareMillimeter returns the value in KilogramPerSecondPerSquareMillimeter.
func (a *MassFlux) KilogramsPerSecondPerSquareMillimeter() float64 {
	if a.kilograms_per_second_per_square_millimeterLazy != nil {
		return *a.kilograms_per_second_per_square_millimeterLazy
	}
	kilograms_per_second_per_square_millimeter := a.convertFromBase(MassFluxKilogramPerSecondPerSquareMillimeter)
	a.kilograms_per_second_per_square_millimeterLazy = &kilograms_per_second_per_square_millimeter
	return kilograms_per_second_per_square_millimeter
}

// KilogramPerHourPerSquareMeter returns the value in KilogramPerHourPerSquareMeter.
func (a *MassFlux) KilogramsPerHourPerSquareMeter() float64 {
	if a.kilograms_per_hour_per_square_meterLazy != nil {
		return *a.kilograms_per_hour_per_square_meterLazy
	}
	kilograms_per_hour_per_square_meter := a.convertFromBase(MassFluxKilogramPerHourPerSquareMeter)
	a.kilograms_per_hour_per_square_meterLazy = &kilograms_per_hour_per_square_meter
	return kilograms_per_hour_per_square_meter
}

// KilogramPerHourPerSquareCentimeter returns the value in KilogramPerHourPerSquareCentimeter.
func (a *MassFlux) KilogramsPerHourPerSquareCentimeter() float64 {
	if a.kilograms_per_hour_per_square_centimeterLazy != nil {
		return *a.kilograms_per_hour_per_square_centimeterLazy
	}
	kilograms_per_hour_per_square_centimeter := a.convertFromBase(MassFluxKilogramPerHourPerSquareCentimeter)
	a.kilograms_per_hour_per_square_centimeterLazy = &kilograms_per_hour_per_square_centimeter
	return kilograms_per_hour_per_square_centimeter
}

// KilogramPerHourPerSquareMillimeter returns the value in KilogramPerHourPerSquareMillimeter.
func (a *MassFlux) KilogramsPerHourPerSquareMillimeter() float64 {
	if a.kilograms_per_hour_per_square_millimeterLazy != nil {
		return *a.kilograms_per_hour_per_square_millimeterLazy
	}
	kilograms_per_hour_per_square_millimeter := a.convertFromBase(MassFluxKilogramPerHourPerSquareMillimeter)
	a.kilograms_per_hour_per_square_millimeterLazy = &kilograms_per_hour_per_square_millimeter
	return kilograms_per_hour_per_square_millimeter
}


// ToDto creates an MassFluxDto representation.
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

// ToDtoJSON creates an MassFluxDto representation.
func (a *MassFlux) ToDtoJSON(holdInUnit *MassFluxUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MassFlux to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a MassFlux) String() string {
	return a.ToString(MassFluxKilogramPerSecondPerSquareMeter, 2)
}

// ToString formats the MassFlux to string.
// fractionalDigits -1 for not mention
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

// Check if the given MassFlux are equals to the current MassFlux
func (a *MassFlux) Equals(other *MassFlux) bool {
	return a.value == other.BaseValue()
}

// Check if the given MassFlux are equals to the current MassFlux
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

// Add the given MassFlux to the current MassFlux.
func (a *MassFlux) Add(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value + other.BaseValue()}
}

// Subtract the given MassFlux to the current MassFlux.
func (a *MassFlux) Subtract(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value - other.BaseValue()}
}

// Multiply the given MassFlux to the current MassFlux.
func (a *MassFlux) Multiply(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value * other.BaseValue()}
}

// Divide the given MassFlux to the current MassFlux.
func (a *MassFlux) Divide(other *MassFlux) *MassFlux {
	return &MassFlux{value: a.value / other.BaseValue()}
}