// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LuminanceUnits defines various units of Luminance.
type LuminanceUnits string

const (
    
        // 
        LuminanceCandelaPerSquareMeter LuminanceUnits = "CandelaPerSquareMeter"
        // 
        LuminanceCandelaPerSquareFoot LuminanceUnits = "CandelaPerSquareFoot"
        // 
        LuminanceCandelaPerSquareInch LuminanceUnits = "CandelaPerSquareInch"
        // 
        LuminanceNit LuminanceUnits = "Nit"
        // 
        LuminanceNanocandelaPerSquareMeter LuminanceUnits = "NanocandelaPerSquareMeter"
        // 
        LuminanceMicrocandelaPerSquareMeter LuminanceUnits = "MicrocandelaPerSquareMeter"
        // 
        LuminanceMillicandelaPerSquareMeter LuminanceUnits = "MillicandelaPerSquareMeter"
        // 
        LuminanceCenticandelaPerSquareMeter LuminanceUnits = "CenticandelaPerSquareMeter"
        // 
        LuminanceDecicandelaPerSquareMeter LuminanceUnits = "DecicandelaPerSquareMeter"
        // 
        LuminanceKilocandelaPerSquareMeter LuminanceUnits = "KilocandelaPerSquareMeter"
)

var internalLuminanceUnitsMap = map[LuminanceUnits]bool{
	
	LuminanceCandelaPerSquareMeter: true,
	LuminanceCandelaPerSquareFoot: true,
	LuminanceCandelaPerSquareInch: true,
	LuminanceNit: true,
	LuminanceNanocandelaPerSquareMeter: true,
	LuminanceMicrocandelaPerSquareMeter: true,
	LuminanceMillicandelaPerSquareMeter: true,
	LuminanceCenticandelaPerSquareMeter: true,
	LuminanceDecicandelaPerSquareMeter: true,
	LuminanceKilocandelaPerSquareMeter: true,
}

// LuminanceDto represents a Luminance measurement with a numerical value and its corresponding unit.
type LuminanceDto struct {
    // Value is the numerical representation of the Luminance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Luminance, as defined in the LuminanceUnits enumeration.
	Unit  LuminanceUnits `json:"unit" validate:"required,oneof=CandelaPerSquareMeter CandelaPerSquareFoot CandelaPerSquareInch Nit NanocandelaPerSquareMeter MicrocandelaPerSquareMeter MillicandelaPerSquareMeter CenticandelaPerSquareMeter DecicandelaPerSquareMeter KilocandelaPerSquareMeter"`
}

// LuminanceDtoFactory groups methods for creating and serializing LuminanceDto objects.
type LuminanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LuminanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LuminanceDtoFactory) FromJSON(data []byte) (*LuminanceDto, error) {
	a := LuminanceDto{}

    // Parse JSON into LuminanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a LuminanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LuminanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Luminance represents a measurement in a of Luminance.
//
// None
type Luminance struct {
	// value is the base measurement stored internally.
	value       float64
    
    candelas_per_square_meterLazy *float64 
    candelas_per_square_footLazy *float64 
    candelas_per_square_inchLazy *float64 
    nitsLazy *float64 
    nanocandelas_per_square_meterLazy *float64 
    microcandelas_per_square_meterLazy *float64 
    millicandelas_per_square_meterLazy *float64 
    centicandelas_per_square_meterLazy *float64 
    decicandelas_per_square_meterLazy *float64 
    kilocandelas_per_square_meterLazy *float64 
}

// LuminanceFactory groups methods for creating Luminance instances.
type LuminanceFactory struct{}

// CreateLuminance creates a new Luminance instance from the given value and unit.
func (uf LuminanceFactory) CreateLuminance(value float64, unit LuminanceUnits) (*Luminance, error) {
	return newLuminance(value, unit)
}

// FromDto converts a LuminanceDto to a Luminance instance.
func (uf LuminanceFactory) FromDto(dto LuminanceDto) (*Luminance, error) {
	return newLuminance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Luminance instance.
func (uf LuminanceFactory) FromDtoJSON(data []byte) (*Luminance, error) {
	unitDto, err := LuminanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LuminanceDto from JSON: %w", err)
	}
	return LuminanceFactory{}.FromDto(*unitDto)
}


// FromCandelasPerSquareMeter creates a new Luminance instance from a value in CandelasPerSquareMeter.
func (uf LuminanceFactory) FromCandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareMeter)
}

// FromCandelasPerSquareFoot creates a new Luminance instance from a value in CandelasPerSquareFoot.
func (uf LuminanceFactory) FromCandelasPerSquareFoot(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareFoot)
}

// FromCandelasPerSquareInch creates a new Luminance instance from a value in CandelasPerSquareInch.
func (uf LuminanceFactory) FromCandelasPerSquareInch(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareInch)
}

// FromNits creates a new Luminance instance from a value in Nits.
func (uf LuminanceFactory) FromNits(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceNit)
}

// FromNanocandelasPerSquareMeter creates a new Luminance instance from a value in NanocandelasPerSquareMeter.
func (uf LuminanceFactory) FromNanocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceNanocandelaPerSquareMeter)
}

// FromMicrocandelasPerSquareMeter creates a new Luminance instance from a value in MicrocandelasPerSquareMeter.
func (uf LuminanceFactory) FromMicrocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceMicrocandelaPerSquareMeter)
}

// FromMillicandelasPerSquareMeter creates a new Luminance instance from a value in MillicandelasPerSquareMeter.
func (uf LuminanceFactory) FromMillicandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceMillicandelaPerSquareMeter)
}

// FromCenticandelasPerSquareMeter creates a new Luminance instance from a value in CenticandelasPerSquareMeter.
func (uf LuminanceFactory) FromCenticandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCenticandelaPerSquareMeter)
}

// FromDecicandelasPerSquareMeter creates a new Luminance instance from a value in DecicandelasPerSquareMeter.
func (uf LuminanceFactory) FromDecicandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceDecicandelaPerSquareMeter)
}

// FromKilocandelasPerSquareMeter creates a new Luminance instance from a value in KilocandelasPerSquareMeter.
func (uf LuminanceFactory) FromKilocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceKilocandelaPerSquareMeter)
}


// newLuminance creates a new Luminance.
func newLuminance(value float64, fromUnit LuminanceUnits) (*Luminance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalLuminanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in LuminanceUnits", fromUnit)
	}
	a := &Luminance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Luminance in CandelaPerSquareMeter unit (the base/default quantity).
func (a *Luminance) BaseValue() float64 {
	return a.value
}


// CandelasPerSquareMeter returns the Luminance value in CandelasPerSquareMeter.
//
// 
func (a *Luminance) CandelasPerSquareMeter() float64 {
	if a.candelas_per_square_meterLazy != nil {
		return *a.candelas_per_square_meterLazy
	}
	candelas_per_square_meter := a.convertFromBase(LuminanceCandelaPerSquareMeter)
	a.candelas_per_square_meterLazy = &candelas_per_square_meter
	return candelas_per_square_meter
}

// CandelasPerSquareFoot returns the Luminance value in CandelasPerSquareFoot.
//
// 
func (a *Luminance) CandelasPerSquareFoot() float64 {
	if a.candelas_per_square_footLazy != nil {
		return *a.candelas_per_square_footLazy
	}
	candelas_per_square_foot := a.convertFromBase(LuminanceCandelaPerSquareFoot)
	a.candelas_per_square_footLazy = &candelas_per_square_foot
	return candelas_per_square_foot
}

// CandelasPerSquareInch returns the Luminance value in CandelasPerSquareInch.
//
// 
func (a *Luminance) CandelasPerSquareInch() float64 {
	if a.candelas_per_square_inchLazy != nil {
		return *a.candelas_per_square_inchLazy
	}
	candelas_per_square_inch := a.convertFromBase(LuminanceCandelaPerSquareInch)
	a.candelas_per_square_inchLazy = &candelas_per_square_inch
	return candelas_per_square_inch
}

// Nits returns the Luminance value in Nits.
//
// 
func (a *Luminance) Nits() float64 {
	if a.nitsLazy != nil {
		return *a.nitsLazy
	}
	nits := a.convertFromBase(LuminanceNit)
	a.nitsLazy = &nits
	return nits
}

// NanocandelasPerSquareMeter returns the Luminance value in NanocandelasPerSquareMeter.
//
// 
func (a *Luminance) NanocandelasPerSquareMeter() float64 {
	if a.nanocandelas_per_square_meterLazy != nil {
		return *a.nanocandelas_per_square_meterLazy
	}
	nanocandelas_per_square_meter := a.convertFromBase(LuminanceNanocandelaPerSquareMeter)
	a.nanocandelas_per_square_meterLazy = &nanocandelas_per_square_meter
	return nanocandelas_per_square_meter
}

// MicrocandelasPerSquareMeter returns the Luminance value in MicrocandelasPerSquareMeter.
//
// 
func (a *Luminance) MicrocandelasPerSquareMeter() float64 {
	if a.microcandelas_per_square_meterLazy != nil {
		return *a.microcandelas_per_square_meterLazy
	}
	microcandelas_per_square_meter := a.convertFromBase(LuminanceMicrocandelaPerSquareMeter)
	a.microcandelas_per_square_meterLazy = &microcandelas_per_square_meter
	return microcandelas_per_square_meter
}

// MillicandelasPerSquareMeter returns the Luminance value in MillicandelasPerSquareMeter.
//
// 
func (a *Luminance) MillicandelasPerSquareMeter() float64 {
	if a.millicandelas_per_square_meterLazy != nil {
		return *a.millicandelas_per_square_meterLazy
	}
	millicandelas_per_square_meter := a.convertFromBase(LuminanceMillicandelaPerSquareMeter)
	a.millicandelas_per_square_meterLazy = &millicandelas_per_square_meter
	return millicandelas_per_square_meter
}

// CenticandelasPerSquareMeter returns the Luminance value in CenticandelasPerSquareMeter.
//
// 
func (a *Luminance) CenticandelasPerSquareMeter() float64 {
	if a.centicandelas_per_square_meterLazy != nil {
		return *a.centicandelas_per_square_meterLazy
	}
	centicandelas_per_square_meter := a.convertFromBase(LuminanceCenticandelaPerSquareMeter)
	a.centicandelas_per_square_meterLazy = &centicandelas_per_square_meter
	return centicandelas_per_square_meter
}

// DecicandelasPerSquareMeter returns the Luminance value in DecicandelasPerSquareMeter.
//
// 
func (a *Luminance) DecicandelasPerSquareMeter() float64 {
	if a.decicandelas_per_square_meterLazy != nil {
		return *a.decicandelas_per_square_meterLazy
	}
	decicandelas_per_square_meter := a.convertFromBase(LuminanceDecicandelaPerSquareMeter)
	a.decicandelas_per_square_meterLazy = &decicandelas_per_square_meter
	return decicandelas_per_square_meter
}

// KilocandelasPerSquareMeter returns the Luminance value in KilocandelasPerSquareMeter.
//
// 
func (a *Luminance) KilocandelasPerSquareMeter() float64 {
	if a.kilocandelas_per_square_meterLazy != nil {
		return *a.kilocandelas_per_square_meterLazy
	}
	kilocandelas_per_square_meter := a.convertFromBase(LuminanceKilocandelaPerSquareMeter)
	a.kilocandelas_per_square_meterLazy = &kilocandelas_per_square_meter
	return kilocandelas_per_square_meter
}


// ToDto creates a LuminanceDto representation from the Luminance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CandelaPerSquareMeter by default.
func (a *Luminance) ToDto(holdInUnit *LuminanceUnits) LuminanceDto {
	if holdInUnit == nil {
		defaultUnit := LuminanceCandelaPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LuminanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Luminance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by CandelaPerSquareMeter by default.
func (a *Luminance) ToDtoJSON(holdInUnit *LuminanceUnits) ([]byte, error) {
	// Convert to LuminanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Luminance to a specific unit value.
// The function uses the provided unit type (LuminanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Luminance) Convert(toUnit LuminanceUnits) float64 {
	switch toUnit { 
    case LuminanceCandelaPerSquareMeter:
		return a.CandelasPerSquareMeter()
    case LuminanceCandelaPerSquareFoot:
		return a.CandelasPerSquareFoot()
    case LuminanceCandelaPerSquareInch:
		return a.CandelasPerSquareInch()
    case LuminanceNit:
		return a.Nits()
    case LuminanceNanocandelaPerSquareMeter:
		return a.NanocandelasPerSquareMeter()
    case LuminanceMicrocandelaPerSquareMeter:
		return a.MicrocandelasPerSquareMeter()
    case LuminanceMillicandelaPerSquareMeter:
		return a.MillicandelasPerSquareMeter()
    case LuminanceCenticandelaPerSquareMeter:
		return a.CenticandelasPerSquareMeter()
    case LuminanceDecicandelaPerSquareMeter:
		return a.DecicandelasPerSquareMeter()
    case LuminanceKilocandelaPerSquareMeter:
		return a.KilocandelasPerSquareMeter()
	default:
		return math.NaN()
	}
}

func (a *Luminance) convertFromBase(toUnit LuminanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminanceCandelaPerSquareMeter:
		return (value) 
	case LuminanceCandelaPerSquareFoot:
		return (value * 9.290304e-2) 
	case LuminanceCandelaPerSquareInch:
		return (value * 0.00064516) 
	case LuminanceNit:
		return (value) 
	case LuminanceNanocandelaPerSquareMeter:
		return ((value) / 1e-09) 
	case LuminanceMicrocandelaPerSquareMeter:
		return ((value) / 1e-06) 
	case LuminanceMillicandelaPerSquareMeter:
		return ((value) / 0.001) 
	case LuminanceCenticandelaPerSquareMeter:
		return ((value) / 0.01) 
	case LuminanceDecicandelaPerSquareMeter:
		return ((value) / 0.1) 
	case LuminanceKilocandelaPerSquareMeter:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Luminance) convertToBase(value float64, fromUnit LuminanceUnits) float64 {
	switch fromUnit { 
	case LuminanceCandelaPerSquareMeter:
		return (value) 
	case LuminanceCandelaPerSquareFoot:
		return (value / 9.290304e-2) 
	case LuminanceCandelaPerSquareInch:
		return (value / 0.00064516) 
	case LuminanceNit:
		return (value) 
	case LuminanceNanocandelaPerSquareMeter:
		return ((value) * 1e-09) 
	case LuminanceMicrocandelaPerSquareMeter:
		return ((value) * 1e-06) 
	case LuminanceMillicandelaPerSquareMeter:
		return ((value) * 0.001) 
	case LuminanceCenticandelaPerSquareMeter:
		return ((value) * 0.01) 
	case LuminanceDecicandelaPerSquareMeter:
		return ((value) * 0.1) 
	case LuminanceKilocandelaPerSquareMeter:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Luminance in the default unit (CandelaPerSquareMeter),
// formatted to two decimal places.
func (a Luminance) String() string {
	return a.ToString(LuminanceCandelaPerSquareMeter, 2)
}

// ToString formats the Luminance value as a string with the specified unit and fractional digits.
// It converts the Luminance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Luminance value will be converted (e.g., CandelaPerSquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Luminance with the unit abbreviation.
func (a *Luminance) ToString(unit LuminanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetLuminanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetLuminanceAbbreviation(unit))
}

// Equals checks if the given Luminance is equal to the current Luminance.
//
// Parameters:
//    other: The Luminance to compare against.
//
// Returns:
//    bool: Returns true if both Luminance are equal, false otherwise.
func (a *Luminance) Equals(other *Luminance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Luminance with another Luminance.
// It returns -1 if the current Luminance is less than the other Luminance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Luminance to compare against.
//
// Returns:
//    int: -1 if the current Luminance is less, 1 if greater, and 0 if equal.
func (a *Luminance) CompareTo(other *Luminance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Luminance to the current Luminance and returns the result.
// The result is a new Luminance instance with the sum of the values.
//
// Parameters:
//    other: The Luminance to add to the current Luminance.
//
// Returns:
//    *Luminance: A new Luminance instance representing the sum of both Luminance.
func (a *Luminance) Add(other *Luminance) *Luminance {
	return &Luminance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Luminance from the current Luminance and returns the result.
// The result is a new Luminance instance with the difference of the values.
//
// Parameters:
//    other: The Luminance to subtract from the current Luminance.
//
// Returns:
//    *Luminance: A new Luminance instance representing the difference of both Luminance.
func (a *Luminance) Subtract(other *Luminance) *Luminance {
	return &Luminance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Luminance by the given Luminance and returns the result.
// The result is a new Luminance instance with the product of the values.
//
// Parameters:
//    other: The Luminance to multiply with the current Luminance.
//
// Returns:
//    *Luminance: A new Luminance instance representing the product of both Luminance.
func (a *Luminance) Multiply(other *Luminance) *Luminance {
	return &Luminance{value: a.value * other.BaseValue()}
}

// Divide divides the current Luminance by the given Luminance and returns the result.
// The result is a new Luminance instance with the quotient of the values.
//
// Parameters:
//    other: The Luminance to divide the current Luminance by.
//
// Returns:
//    *Luminance: A new Luminance instance representing the quotient of both Luminance.
func (a *Luminance) Divide(other *Luminance) *Luminance {
	return &Luminance{value: a.value / other.BaseValue()}
}

// GetLuminanceAbbreviation gets the unit abbreviation.
func GetLuminanceAbbreviation(unit LuminanceUnits) string {
	switch unit { 
	case LuminanceCandelaPerSquareMeter:
		return "Cd/m²" 
	case LuminanceCandelaPerSquareFoot:
		return "Cd/ft²" 
	case LuminanceCandelaPerSquareInch:
		return "Cd/in²" 
	case LuminanceNit:
		return "nt" 
	case LuminanceNanocandelaPerSquareMeter:
		return "nCd/m²" 
	case LuminanceMicrocandelaPerSquareMeter:
		return "μCd/m²" 
	case LuminanceMillicandelaPerSquareMeter:
		return "mCd/m²" 
	case LuminanceCenticandelaPerSquareMeter:
		return "cCd/m²" 
	case LuminanceDecicandelaPerSquareMeter:
		return "dCd/m²" 
	case LuminanceKilocandelaPerSquareMeter:
		return "kCd/m²" 
	default:
		return ""
	}
}