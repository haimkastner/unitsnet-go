// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricConductivityUnits defines various units of ElectricConductivity.
type ElectricConductivityUnits string

const (
    
        // 
        ElectricConductivitySiemensPerMeter ElectricConductivityUnits = "SiemensPerMeter"
        // 
        ElectricConductivitySiemensPerInch ElectricConductivityUnits = "SiemensPerInch"
        // 
        ElectricConductivitySiemensPerFoot ElectricConductivityUnits = "SiemensPerFoot"
        // 
        ElectricConductivitySiemensPerCentimeter ElectricConductivityUnits = "SiemensPerCentimeter"
        // 
        ElectricConductivityMicrosiemensPerCentimeter ElectricConductivityUnits = "MicrosiemensPerCentimeter"
        // 
        ElectricConductivityMillisiemensPerCentimeter ElectricConductivityUnits = "MillisiemensPerCentimeter"
)

// ElectricConductivityDto represents a ElectricConductivity measurement with a numerical value and its corresponding unit.
type ElectricConductivityDto struct {
    // Value is the numerical representation of the ElectricConductivity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricConductivity, as defined in the ElectricConductivityUnits enumeration.
	Unit  ElectricConductivityUnits `json:"unit"`
}

// ElectricConductivityDtoFactory groups methods for creating and serializing ElectricConductivityDto objects.
type ElectricConductivityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricConductivityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricConductivityDtoFactory) FromJSON(data []byte) (*ElectricConductivityDto, error) {
	a := ElectricConductivityDto{}

    // Parse JSON into ElectricConductivityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricConductivityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricConductivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricConductivity represents a measurement in a of ElectricConductivity.
//
// Electrical conductivity or specific conductance is the reciprocal of electrical resistivity, and measures a material's ability to conduct an electric current.
type ElectricConductivity struct {
	// value is the base measurement stored internally.
	value       float64
    
    siemens_per_meterLazy *float64 
    siemens_per_inchLazy *float64 
    siemens_per_footLazy *float64 
    siemens_per_centimeterLazy *float64 
    microsiemens_per_centimeterLazy *float64 
    millisiemens_per_centimeterLazy *float64 
}

// ElectricConductivityFactory groups methods for creating ElectricConductivity instances.
type ElectricConductivityFactory struct{}

// CreateElectricConductivity creates a new ElectricConductivity instance from the given value and unit.
func (uf ElectricConductivityFactory) CreateElectricConductivity(value float64, unit ElectricConductivityUnits) (*ElectricConductivity, error) {
	return newElectricConductivity(value, unit)
}

// FromDto converts a ElectricConductivityDto to a ElectricConductivity instance.
func (uf ElectricConductivityFactory) FromDto(dto ElectricConductivityDto) (*ElectricConductivity, error) {
	return newElectricConductivity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricConductivity instance.
func (uf ElectricConductivityFactory) FromDtoJSON(data []byte) (*ElectricConductivity, error) {
	unitDto, err := ElectricConductivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricConductivityDto from JSON: %w", err)
	}
	return ElectricConductivityFactory{}.FromDto(*unitDto)
}


// FromSiemensPerMeter creates a new ElectricConductivity instance from a value in SiemensPerMeter.
func (uf ElectricConductivityFactory) FromSiemensPerMeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerMeter)
}

// FromSiemensPerInch creates a new ElectricConductivity instance from a value in SiemensPerInch.
func (uf ElectricConductivityFactory) FromSiemensPerInch(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerInch)
}

// FromSiemensPerFoot creates a new ElectricConductivity instance from a value in SiemensPerFoot.
func (uf ElectricConductivityFactory) FromSiemensPerFoot(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerFoot)
}

// FromSiemensPerCentimeter creates a new ElectricConductivity instance from a value in SiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromSiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerCentimeter)
}

// FromMicrosiemensPerCentimeter creates a new ElectricConductivity instance from a value in MicrosiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromMicrosiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivityMicrosiemensPerCentimeter)
}

// FromMillisiemensPerCentimeter creates a new ElectricConductivity instance from a value in MillisiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromMillisiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivityMillisiemensPerCentimeter)
}


// newElectricConductivity creates a new ElectricConductivity.
func newElectricConductivity(value float64, fromUnit ElectricConductivityUnits) (*ElectricConductivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricConductivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricConductivity in SiemensPerMeter unit (the base/default quantity).
func (a *ElectricConductivity) BaseValue() float64 {
	return a.value
}


// SiemensPerMeter returns the ElectricConductivity value in SiemensPerMeter.
//
// 
func (a *ElectricConductivity) SiemensPerMeter() float64 {
	if a.siemens_per_meterLazy != nil {
		return *a.siemens_per_meterLazy
	}
	siemens_per_meter := a.convertFromBase(ElectricConductivitySiemensPerMeter)
	a.siemens_per_meterLazy = &siemens_per_meter
	return siemens_per_meter
}

// SiemensPerInch returns the ElectricConductivity value in SiemensPerInch.
//
// 
func (a *ElectricConductivity) SiemensPerInch() float64 {
	if a.siemens_per_inchLazy != nil {
		return *a.siemens_per_inchLazy
	}
	siemens_per_inch := a.convertFromBase(ElectricConductivitySiemensPerInch)
	a.siemens_per_inchLazy = &siemens_per_inch
	return siemens_per_inch
}

// SiemensPerFoot returns the ElectricConductivity value in SiemensPerFoot.
//
// 
func (a *ElectricConductivity) SiemensPerFoot() float64 {
	if a.siemens_per_footLazy != nil {
		return *a.siemens_per_footLazy
	}
	siemens_per_foot := a.convertFromBase(ElectricConductivitySiemensPerFoot)
	a.siemens_per_footLazy = &siemens_per_foot
	return siemens_per_foot
}

// SiemensPerCentimeter returns the ElectricConductivity value in SiemensPerCentimeter.
//
// 
func (a *ElectricConductivity) SiemensPerCentimeter() float64 {
	if a.siemens_per_centimeterLazy != nil {
		return *a.siemens_per_centimeterLazy
	}
	siemens_per_centimeter := a.convertFromBase(ElectricConductivitySiemensPerCentimeter)
	a.siemens_per_centimeterLazy = &siemens_per_centimeter
	return siemens_per_centimeter
}

// MicrosiemensPerCentimeter returns the ElectricConductivity value in MicrosiemensPerCentimeter.
//
// 
func (a *ElectricConductivity) MicrosiemensPerCentimeter() float64 {
	if a.microsiemens_per_centimeterLazy != nil {
		return *a.microsiemens_per_centimeterLazy
	}
	microsiemens_per_centimeter := a.convertFromBase(ElectricConductivityMicrosiemensPerCentimeter)
	a.microsiemens_per_centimeterLazy = &microsiemens_per_centimeter
	return microsiemens_per_centimeter
}

// MillisiemensPerCentimeter returns the ElectricConductivity value in MillisiemensPerCentimeter.
//
// 
func (a *ElectricConductivity) MillisiemensPerCentimeter() float64 {
	if a.millisiemens_per_centimeterLazy != nil {
		return *a.millisiemens_per_centimeterLazy
	}
	millisiemens_per_centimeter := a.convertFromBase(ElectricConductivityMillisiemensPerCentimeter)
	a.millisiemens_per_centimeterLazy = &millisiemens_per_centimeter
	return millisiemens_per_centimeter
}


// ToDto creates a ElectricConductivityDto representation from the ElectricConductivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SiemensPerMeter by default.
func (a *ElectricConductivity) ToDto(holdInUnit *ElectricConductivityUnits) ElectricConductivityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricConductivitySiemensPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricConductivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricConductivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SiemensPerMeter by default.
func (a *ElectricConductivity) ToDtoJSON(holdInUnit *ElectricConductivityUnits) ([]byte, error) {
	// Convert to ElectricConductivityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricConductivity to a specific unit value.
// The function uses the provided unit type (ElectricConductivityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricConductivity) Convert(toUnit ElectricConductivityUnits) float64 {
	switch toUnit { 
    case ElectricConductivitySiemensPerMeter:
		return a.SiemensPerMeter()
    case ElectricConductivitySiemensPerInch:
		return a.SiemensPerInch()
    case ElectricConductivitySiemensPerFoot:
		return a.SiemensPerFoot()
    case ElectricConductivitySiemensPerCentimeter:
		return a.SiemensPerCentimeter()
    case ElectricConductivityMicrosiemensPerCentimeter:
		return a.MicrosiemensPerCentimeter()
    case ElectricConductivityMillisiemensPerCentimeter:
		return a.MillisiemensPerCentimeter()
	default:
		return math.NaN()
	}
}

func (a *ElectricConductivity) convertFromBase(toUnit ElectricConductivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricConductivitySiemensPerMeter:
		return (value) 
	case ElectricConductivitySiemensPerInch:
		return (value / 3.937007874015748e1) 
	case ElectricConductivitySiemensPerFoot:
		return (value / 3.2808398950131234) 
	case ElectricConductivitySiemensPerCentimeter:
		return (value / 1e2) 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return ((value / 1e2) / 1e-06) 
	case ElectricConductivityMillisiemensPerCentimeter:
		return ((value / 1e2) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricConductivity) convertToBase(value float64, fromUnit ElectricConductivityUnits) float64 {
	switch fromUnit { 
	case ElectricConductivitySiemensPerMeter:
		return (value) 
	case ElectricConductivitySiemensPerInch:
		return (value * 3.937007874015748e1) 
	case ElectricConductivitySiemensPerFoot:
		return (value * 3.2808398950131234) 
	case ElectricConductivitySiemensPerCentimeter:
		return (value * 1e2) 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return ((value * 1e2) * 1e-06) 
	case ElectricConductivityMillisiemensPerCentimeter:
		return ((value * 1e2) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricConductivity in the default unit (SiemensPerMeter),
// formatted to two decimal places.
func (a ElectricConductivity) String() string {
	return a.ToString(ElectricConductivitySiemensPerMeter, 2)
}

// ToString formats the ElectricConductivity value as a string with the specified unit and fractional digits.
// It converts the ElectricConductivity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricConductivity value will be converted (e.g., SiemensPerMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricConductivity with the unit abbreviation.
func (a *ElectricConductivity) ToString(unit ElectricConductivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricConductivityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricConductivityAbbreviation(unit))
}

// Equals checks if the given ElectricConductivity is equal to the current ElectricConductivity.
//
// Parameters:
//    other: The ElectricConductivity to compare against.
//
// Returns:
//    bool: Returns true if both ElectricConductivity are equal, false otherwise.
func (a *ElectricConductivity) Equals(other *ElectricConductivity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricConductivity with another ElectricConductivity.
// It returns -1 if the current ElectricConductivity is less than the other ElectricConductivity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricConductivity to compare against.
//
// Returns:
//    int: -1 if the current ElectricConductivity is less, 1 if greater, and 0 if equal.
func (a *ElectricConductivity) CompareTo(other *ElectricConductivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricConductivity to the current ElectricConductivity and returns the result.
// The result is a new ElectricConductivity instance with the sum of the values.
//
// Parameters:
//    other: The ElectricConductivity to add to the current ElectricConductivity.
//
// Returns:
//    *ElectricConductivity: A new ElectricConductivity instance representing the sum of both ElectricConductivity.
func (a *ElectricConductivity) Add(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricConductivity from the current ElectricConductivity and returns the result.
// The result is a new ElectricConductivity instance with the difference of the values.
//
// Parameters:
//    other: The ElectricConductivity to subtract from the current ElectricConductivity.
//
// Returns:
//    *ElectricConductivity: A new ElectricConductivity instance representing the difference of both ElectricConductivity.
func (a *ElectricConductivity) Subtract(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricConductivity by the given ElectricConductivity and returns the result.
// The result is a new ElectricConductivity instance with the product of the values.
//
// Parameters:
//    other: The ElectricConductivity to multiply with the current ElectricConductivity.
//
// Returns:
//    *ElectricConductivity: A new ElectricConductivity instance representing the product of both ElectricConductivity.
func (a *ElectricConductivity) Multiply(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricConductivity by the given ElectricConductivity and returns the result.
// The result is a new ElectricConductivity instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricConductivity to divide the current ElectricConductivity by.
//
// Returns:
//    *ElectricConductivity: A new ElectricConductivity instance representing the quotient of both ElectricConductivity.
func (a *ElectricConductivity) Divide(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value / other.BaseValue()}
}

// GetElectricConductivityAbbreviation gets the unit abbreviation.
func GetElectricConductivityAbbreviation(unit ElectricConductivityUnits) string {
	switch unit { 
	case ElectricConductivitySiemensPerMeter:
		return "S/m" 
	case ElectricConductivitySiemensPerInch:
		return "S/in" 
	case ElectricConductivitySiemensPerFoot:
		return "S/ft" 
	case ElectricConductivitySiemensPerCentimeter:
		return "S/cm" 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return "μS/cm" 
	case ElectricConductivityMillisiemensPerCentimeter:
		return "mS/cm" 
	default:
		return ""
	}
}