// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricInductanceUnits defines various units of ElectricInductance.
type ElectricInductanceUnits string

const (
    
        // 
        ElectricInductanceHenry ElectricInductanceUnits = "Henry"
        // 
        ElectricInductancePicohenry ElectricInductanceUnits = "Picohenry"
        // 
        ElectricInductanceNanohenry ElectricInductanceUnits = "Nanohenry"
        // 
        ElectricInductanceMicrohenry ElectricInductanceUnits = "Microhenry"
        // 
        ElectricInductanceMillihenry ElectricInductanceUnits = "Millihenry"
)

var internalElectricInductanceUnitsMap = map[ElectricInductanceUnits]bool{
	
	ElectricInductanceHenry: true,
	ElectricInductancePicohenry: true,
	ElectricInductanceNanohenry: true,
	ElectricInductanceMicrohenry: true,
	ElectricInductanceMillihenry: true,
}

// ElectricInductanceDto represents a ElectricInductance measurement with a numerical value and its corresponding unit.
type ElectricInductanceDto struct {
    // Value is the numerical representation of the ElectricInductance.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricInductance, as defined in the ElectricInductanceUnits enumeration.
	Unit  ElectricInductanceUnits `json:"unit" validate:"required,oneof=Henry Picohenry Nanohenry Microhenry Millihenry"`
}

// ElectricInductanceDtoFactory groups methods for creating and serializing ElectricInductanceDto objects.
type ElectricInductanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricInductanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricInductanceDtoFactory) FromJSON(data []byte) (*ElectricInductanceDto, error) {
	a := ElectricInductanceDto{}

    // Parse JSON into ElectricInductanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricInductanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricInductanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricInductance represents a measurement in a of ElectricInductance.
//
// Inductance is a property of an electrical conductor which opposes a change in current.
type ElectricInductance struct {
	// value is the base measurement stored internally.
	value       float64
    
    henriesLazy *float64 
    picohenriesLazy *float64 
    nanohenriesLazy *float64 
    microhenriesLazy *float64 
    millihenriesLazy *float64 
}

// ElectricInductanceFactory groups methods for creating ElectricInductance instances.
type ElectricInductanceFactory struct{}

// CreateElectricInductance creates a new ElectricInductance instance from the given value and unit.
func (uf ElectricInductanceFactory) CreateElectricInductance(value float64, unit ElectricInductanceUnits) (*ElectricInductance, error) {
	return newElectricInductance(value, unit)
}

// FromDto converts a ElectricInductanceDto to a ElectricInductance instance.
func (uf ElectricInductanceFactory) FromDto(dto ElectricInductanceDto) (*ElectricInductance, error) {
	return newElectricInductance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricInductance instance.
func (uf ElectricInductanceFactory) FromDtoJSON(data []byte) (*ElectricInductance, error) {
	unitDto, err := ElectricInductanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricInductanceDto from JSON: %w", err)
	}
	return ElectricInductanceFactory{}.FromDto(*unitDto)
}


// FromHenries creates a new ElectricInductance instance from a value in Henries.
func (uf ElectricInductanceFactory) FromHenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceHenry)
}

// FromPicohenries creates a new ElectricInductance instance from a value in Picohenries.
func (uf ElectricInductanceFactory) FromPicohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductancePicohenry)
}

// FromNanohenries creates a new ElectricInductance instance from a value in Nanohenries.
func (uf ElectricInductanceFactory) FromNanohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceNanohenry)
}

// FromMicrohenries creates a new ElectricInductance instance from a value in Microhenries.
func (uf ElectricInductanceFactory) FromMicrohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceMicrohenry)
}

// FromMillihenries creates a new ElectricInductance instance from a value in Millihenries.
func (uf ElectricInductanceFactory) FromMillihenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceMillihenry)
}


// newElectricInductance creates a new ElectricInductance.
func newElectricInductance(value float64, fromUnit ElectricInductanceUnits) (*ElectricInductance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricInductanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricInductanceUnits", fromUnit)
	}
	a := &ElectricInductance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricInductance in Henry unit (the base/default quantity).
func (a *ElectricInductance) BaseValue() float64 {
	return a.value
}


// Henries returns the ElectricInductance value in Henries.
//
// 
func (a *ElectricInductance) Henries() float64 {
	if a.henriesLazy != nil {
		return *a.henriesLazy
	}
	henries := a.convertFromBase(ElectricInductanceHenry)
	a.henriesLazy = &henries
	return henries
}

// Picohenries returns the ElectricInductance value in Picohenries.
//
// 
func (a *ElectricInductance) Picohenries() float64 {
	if a.picohenriesLazy != nil {
		return *a.picohenriesLazy
	}
	picohenries := a.convertFromBase(ElectricInductancePicohenry)
	a.picohenriesLazy = &picohenries
	return picohenries
}

// Nanohenries returns the ElectricInductance value in Nanohenries.
//
// 
func (a *ElectricInductance) Nanohenries() float64 {
	if a.nanohenriesLazy != nil {
		return *a.nanohenriesLazy
	}
	nanohenries := a.convertFromBase(ElectricInductanceNanohenry)
	a.nanohenriesLazy = &nanohenries
	return nanohenries
}

// Microhenries returns the ElectricInductance value in Microhenries.
//
// 
func (a *ElectricInductance) Microhenries() float64 {
	if a.microhenriesLazy != nil {
		return *a.microhenriesLazy
	}
	microhenries := a.convertFromBase(ElectricInductanceMicrohenry)
	a.microhenriesLazy = &microhenries
	return microhenries
}

// Millihenries returns the ElectricInductance value in Millihenries.
//
// 
func (a *ElectricInductance) Millihenries() float64 {
	if a.millihenriesLazy != nil {
		return *a.millihenriesLazy
	}
	millihenries := a.convertFromBase(ElectricInductanceMillihenry)
	a.millihenriesLazy = &millihenries
	return millihenries
}


// ToDto creates a ElectricInductanceDto representation from the ElectricInductance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Henry by default.
func (a *ElectricInductance) ToDto(holdInUnit *ElectricInductanceUnits) ElectricInductanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricInductanceHenry // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricInductanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricInductance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Henry by default.
func (a *ElectricInductance) ToDtoJSON(holdInUnit *ElectricInductanceUnits) ([]byte, error) {
	// Convert to ElectricInductanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricInductance to a specific unit value.
// The function uses the provided unit type (ElectricInductanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricInductance) Convert(toUnit ElectricInductanceUnits) float64 {
	switch toUnit { 
    case ElectricInductanceHenry:
		return a.Henries()
    case ElectricInductancePicohenry:
		return a.Picohenries()
    case ElectricInductanceNanohenry:
		return a.Nanohenries()
    case ElectricInductanceMicrohenry:
		return a.Microhenries()
    case ElectricInductanceMillihenry:
		return a.Millihenries()
	default:
		return math.NaN()
	}
}

func (a *ElectricInductance) convertFromBase(toUnit ElectricInductanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricInductanceHenry:
		return (value) 
	case ElectricInductancePicohenry:
		return ((value) / 1e-12) 
	case ElectricInductanceNanohenry:
		return ((value) / 1e-09) 
	case ElectricInductanceMicrohenry:
		return ((value) / 1e-06) 
	case ElectricInductanceMillihenry:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricInductance) convertToBase(value float64, fromUnit ElectricInductanceUnits) float64 {
	switch fromUnit { 
	case ElectricInductanceHenry:
		return (value) 
	case ElectricInductancePicohenry:
		return ((value) * 1e-12) 
	case ElectricInductanceNanohenry:
		return ((value) * 1e-09) 
	case ElectricInductanceMicrohenry:
		return ((value) * 1e-06) 
	case ElectricInductanceMillihenry:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricInductance in the default unit (Henry),
// formatted to two decimal places.
func (a ElectricInductance) String() string {
	return a.ToString(ElectricInductanceHenry, 2)
}

// ToString formats the ElectricInductance value as a string with the specified unit and fractional digits.
// It converts the ElectricInductance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricInductance value will be converted (e.g., Henry))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricInductance with the unit abbreviation.
func (a *ElectricInductance) ToString(unit ElectricInductanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricInductanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricInductanceAbbreviation(unit))
}

// Equals checks if the given ElectricInductance is equal to the current ElectricInductance.
//
// Parameters:
//    other: The ElectricInductance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricInductance are equal, false otherwise.
func (a *ElectricInductance) Equals(other *ElectricInductance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricInductance with another ElectricInductance.
// It returns -1 if the current ElectricInductance is less than the other ElectricInductance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricInductance to compare against.
//
// Returns:
//    int: -1 if the current ElectricInductance is less, 1 if greater, and 0 if equal.
func (a *ElectricInductance) CompareTo(other *ElectricInductance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricInductance to the current ElectricInductance and returns the result.
// The result is a new ElectricInductance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricInductance to add to the current ElectricInductance.
//
// Returns:
//    *ElectricInductance: A new ElectricInductance instance representing the sum of both ElectricInductance.
func (a *ElectricInductance) Add(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricInductance from the current ElectricInductance and returns the result.
// The result is a new ElectricInductance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricInductance to subtract from the current ElectricInductance.
//
// Returns:
//    *ElectricInductance: A new ElectricInductance instance representing the difference of both ElectricInductance.
func (a *ElectricInductance) Subtract(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricInductance by the given ElectricInductance and returns the result.
// The result is a new ElectricInductance instance with the product of the values.
//
// Parameters:
//    other: The ElectricInductance to multiply with the current ElectricInductance.
//
// Returns:
//    *ElectricInductance: A new ElectricInductance instance representing the product of both ElectricInductance.
func (a *ElectricInductance) Multiply(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricInductance by the given ElectricInductance and returns the result.
// The result is a new ElectricInductance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricInductance to divide the current ElectricInductance by.
//
// Returns:
//    *ElectricInductance: A new ElectricInductance instance representing the quotient of both ElectricInductance.
func (a *ElectricInductance) Divide(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value / other.BaseValue()}
}

// GetElectricInductanceAbbreviation gets the unit abbreviation.
func GetElectricInductanceAbbreviation(unit ElectricInductanceUnits) string {
	switch unit { 
	case ElectricInductanceHenry:
		return "H" 
	case ElectricInductancePicohenry:
		return "pH" 
	case ElectricInductanceNanohenry:
		return "nH" 
	case ElectricInductanceMicrohenry:
		return "μH" 
	case ElectricInductanceMillihenry:
		return "mH" 
	default:
		return ""
	}
}