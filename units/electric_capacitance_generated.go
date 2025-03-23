// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCapacitanceUnits defines various units of ElectricCapacitance.
type ElectricCapacitanceUnits string

const (
    
        // 
        ElectricCapacitanceFarad ElectricCapacitanceUnits = "Farad"
        // 
        ElectricCapacitancePicofarad ElectricCapacitanceUnits = "Picofarad"
        // 
        ElectricCapacitanceNanofarad ElectricCapacitanceUnits = "Nanofarad"
        // 
        ElectricCapacitanceMicrofarad ElectricCapacitanceUnits = "Microfarad"
        // 
        ElectricCapacitanceMillifarad ElectricCapacitanceUnits = "Millifarad"
        // 
        ElectricCapacitanceKilofarad ElectricCapacitanceUnits = "Kilofarad"
        // 
        ElectricCapacitanceMegafarad ElectricCapacitanceUnits = "Megafarad"
)

var internalElectricCapacitanceUnitsMap = map[ElectricCapacitanceUnits]bool{
	
	ElectricCapacitanceFarad: true,
	ElectricCapacitancePicofarad: true,
	ElectricCapacitanceNanofarad: true,
	ElectricCapacitanceMicrofarad: true,
	ElectricCapacitanceMillifarad: true,
	ElectricCapacitanceKilofarad: true,
	ElectricCapacitanceMegafarad: true,
}

// ElectricCapacitanceDto represents a ElectricCapacitance measurement with a numerical value and its corresponding unit.
type ElectricCapacitanceDto struct {
    // Value is the numerical representation of the ElectricCapacitance.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricCapacitance, as defined in the ElectricCapacitanceUnits enumeration.
	Unit  ElectricCapacitanceUnits `json:"unit" validate:"required,oneof=Farad Picofarad Nanofarad Microfarad Millifarad Kilofarad Megafarad"`
}

// ElectricCapacitanceDtoFactory groups methods for creating and serializing ElectricCapacitanceDto objects.
type ElectricCapacitanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricCapacitanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricCapacitanceDtoFactory) FromJSON(data []byte) (*ElectricCapacitanceDto, error) {
	a := ElectricCapacitanceDto{}

    // Parse JSON into ElectricCapacitanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricCapacitanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricCapacitanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricCapacitance represents a measurement in a of ElectricCapacitance.
//
// Capacitance is the capacity of a material object or device to store electric charge.
type ElectricCapacitance struct {
	// value is the base measurement stored internally.
	value       float64
    
    faradsLazy *float64 
    picofaradsLazy *float64 
    nanofaradsLazy *float64 
    microfaradsLazy *float64 
    millifaradsLazy *float64 
    kilofaradsLazy *float64 
    megafaradsLazy *float64 
}

// ElectricCapacitanceFactory groups methods for creating ElectricCapacitance instances.
type ElectricCapacitanceFactory struct{}

// CreateElectricCapacitance creates a new ElectricCapacitance instance from the given value and unit.
func (uf ElectricCapacitanceFactory) CreateElectricCapacitance(value float64, unit ElectricCapacitanceUnits) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, unit)
}

// FromDto converts a ElectricCapacitanceDto to a ElectricCapacitance instance.
func (uf ElectricCapacitanceFactory) FromDto(dto ElectricCapacitanceDto) (*ElectricCapacitance, error) {
	return newElectricCapacitance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricCapacitance instance.
func (uf ElectricCapacitanceFactory) FromDtoJSON(data []byte) (*ElectricCapacitance, error) {
	unitDto, err := ElectricCapacitanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricCapacitanceDto from JSON: %w", err)
	}
	return ElectricCapacitanceFactory{}.FromDto(*unitDto)
}


// FromFarads creates a new ElectricCapacitance instance from a value in Farads.
func (uf ElectricCapacitanceFactory) FromFarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceFarad)
}

// FromPicofarads creates a new ElectricCapacitance instance from a value in Picofarads.
func (uf ElectricCapacitanceFactory) FromPicofarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitancePicofarad)
}

// FromNanofarads creates a new ElectricCapacitance instance from a value in Nanofarads.
func (uf ElectricCapacitanceFactory) FromNanofarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceNanofarad)
}

// FromMicrofarads creates a new ElectricCapacitance instance from a value in Microfarads.
func (uf ElectricCapacitanceFactory) FromMicrofarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceMicrofarad)
}

// FromMillifarads creates a new ElectricCapacitance instance from a value in Millifarads.
func (uf ElectricCapacitanceFactory) FromMillifarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceMillifarad)
}

// FromKilofarads creates a new ElectricCapacitance instance from a value in Kilofarads.
func (uf ElectricCapacitanceFactory) FromKilofarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceKilofarad)
}

// FromMegafarads creates a new ElectricCapacitance instance from a value in Megafarads.
func (uf ElectricCapacitanceFactory) FromMegafarads(value float64) (*ElectricCapacitance, error) {
	return newElectricCapacitance(value, ElectricCapacitanceMegafarad)
}


// newElectricCapacitance creates a new ElectricCapacitance.
func newElectricCapacitance(value float64, fromUnit ElectricCapacitanceUnits) (*ElectricCapacitance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricCapacitanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricCapacitanceUnits", fromUnit)
	}
	a := &ElectricCapacitance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCapacitance in Farad unit (the base/default quantity).
func (a *ElectricCapacitance) BaseValue() float64 {
	return a.value
}


// Farads returns the ElectricCapacitance value in Farads.
//
// 
func (a *ElectricCapacitance) Farads() float64 {
	if a.faradsLazy != nil {
		return *a.faradsLazy
	}
	farads := a.convertFromBase(ElectricCapacitanceFarad)
	a.faradsLazy = &farads
	return farads
}

// Picofarads returns the ElectricCapacitance value in Picofarads.
//
// 
func (a *ElectricCapacitance) Picofarads() float64 {
	if a.picofaradsLazy != nil {
		return *a.picofaradsLazy
	}
	picofarads := a.convertFromBase(ElectricCapacitancePicofarad)
	a.picofaradsLazy = &picofarads
	return picofarads
}

// Nanofarads returns the ElectricCapacitance value in Nanofarads.
//
// 
func (a *ElectricCapacitance) Nanofarads() float64 {
	if a.nanofaradsLazy != nil {
		return *a.nanofaradsLazy
	}
	nanofarads := a.convertFromBase(ElectricCapacitanceNanofarad)
	a.nanofaradsLazy = &nanofarads
	return nanofarads
}

// Microfarads returns the ElectricCapacitance value in Microfarads.
//
// 
func (a *ElectricCapacitance) Microfarads() float64 {
	if a.microfaradsLazy != nil {
		return *a.microfaradsLazy
	}
	microfarads := a.convertFromBase(ElectricCapacitanceMicrofarad)
	a.microfaradsLazy = &microfarads
	return microfarads
}

// Millifarads returns the ElectricCapacitance value in Millifarads.
//
// 
func (a *ElectricCapacitance) Millifarads() float64 {
	if a.millifaradsLazy != nil {
		return *a.millifaradsLazy
	}
	millifarads := a.convertFromBase(ElectricCapacitanceMillifarad)
	a.millifaradsLazy = &millifarads
	return millifarads
}

// Kilofarads returns the ElectricCapacitance value in Kilofarads.
//
// 
func (a *ElectricCapacitance) Kilofarads() float64 {
	if a.kilofaradsLazy != nil {
		return *a.kilofaradsLazy
	}
	kilofarads := a.convertFromBase(ElectricCapacitanceKilofarad)
	a.kilofaradsLazy = &kilofarads
	return kilofarads
}

// Megafarads returns the ElectricCapacitance value in Megafarads.
//
// 
func (a *ElectricCapacitance) Megafarads() float64 {
	if a.megafaradsLazy != nil {
		return *a.megafaradsLazy
	}
	megafarads := a.convertFromBase(ElectricCapacitanceMegafarad)
	a.megafaradsLazy = &megafarads
	return megafarads
}


// ToDto creates a ElectricCapacitanceDto representation from the ElectricCapacitance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Farad by default.
func (a *ElectricCapacitance) ToDto(holdInUnit *ElectricCapacitanceUnits) ElectricCapacitanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCapacitanceFarad // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCapacitanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricCapacitance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Farad by default.
func (a *ElectricCapacitance) ToDtoJSON(holdInUnit *ElectricCapacitanceUnits) ([]byte, error) {
	// Convert to ElectricCapacitanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricCapacitance to a specific unit value.
// The function uses the provided unit type (ElectricCapacitanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricCapacitance) Convert(toUnit ElectricCapacitanceUnits) float64 {
	switch toUnit { 
    case ElectricCapacitanceFarad:
		return a.Farads()
    case ElectricCapacitancePicofarad:
		return a.Picofarads()
    case ElectricCapacitanceNanofarad:
		return a.Nanofarads()
    case ElectricCapacitanceMicrofarad:
		return a.Microfarads()
    case ElectricCapacitanceMillifarad:
		return a.Millifarads()
    case ElectricCapacitanceKilofarad:
		return a.Kilofarads()
    case ElectricCapacitanceMegafarad:
		return a.Megafarads()
	default:
		return math.NaN()
	}
}

func (a *ElectricCapacitance) convertFromBase(toUnit ElectricCapacitanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCapacitanceFarad:
		return (value) 
	case ElectricCapacitancePicofarad:
		return ((value) / 1e-12) 
	case ElectricCapacitanceNanofarad:
		return ((value) / 1e-09) 
	case ElectricCapacitanceMicrofarad:
		return ((value) / 1e-06) 
	case ElectricCapacitanceMillifarad:
		return ((value) / 0.001) 
	case ElectricCapacitanceKilofarad:
		return ((value) / 1000.0) 
	case ElectricCapacitanceMegafarad:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCapacitance) convertToBase(value float64, fromUnit ElectricCapacitanceUnits) float64 {
	switch fromUnit { 
	case ElectricCapacitanceFarad:
		return (value) 
	case ElectricCapacitancePicofarad:
		return ((value) * 1e-12) 
	case ElectricCapacitanceNanofarad:
		return ((value) * 1e-09) 
	case ElectricCapacitanceMicrofarad:
		return ((value) * 1e-06) 
	case ElectricCapacitanceMillifarad:
		return ((value) * 0.001) 
	case ElectricCapacitanceKilofarad:
		return ((value) * 1000.0) 
	case ElectricCapacitanceMegafarad:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricCapacitance in the default unit (Farad),
// formatted to two decimal places.
func (a ElectricCapacitance) String() string {
	return a.ToString(ElectricCapacitanceFarad, 2)
}

// ToString formats the ElectricCapacitance value as a string with the specified unit and fractional digits.
// It converts the ElectricCapacitance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricCapacitance value will be converted (e.g., Farad))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricCapacitance with the unit abbreviation.
func (a *ElectricCapacitance) ToString(unit ElectricCapacitanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricCapacitanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricCapacitanceAbbreviation(unit))
}

// Equals checks if the given ElectricCapacitance is equal to the current ElectricCapacitance.
//
// Parameters:
//    other: The ElectricCapacitance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricCapacitance are equal, false otherwise.
func (a *ElectricCapacitance) Equals(other *ElectricCapacitance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricCapacitance with another ElectricCapacitance.
// It returns -1 if the current ElectricCapacitance is less than the other ElectricCapacitance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricCapacitance to compare against.
//
// Returns:
//    int: -1 if the current ElectricCapacitance is less, 1 if greater, and 0 if equal.
func (a *ElectricCapacitance) CompareTo(other *ElectricCapacitance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricCapacitance to the current ElectricCapacitance and returns the result.
// The result is a new ElectricCapacitance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricCapacitance to add to the current ElectricCapacitance.
//
// Returns:
//    *ElectricCapacitance: A new ElectricCapacitance instance representing the sum of both ElectricCapacitance.
func (a *ElectricCapacitance) Add(other *ElectricCapacitance) *ElectricCapacitance {
	return &ElectricCapacitance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricCapacitance from the current ElectricCapacitance and returns the result.
// The result is a new ElectricCapacitance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricCapacitance to subtract from the current ElectricCapacitance.
//
// Returns:
//    *ElectricCapacitance: A new ElectricCapacitance instance representing the difference of both ElectricCapacitance.
func (a *ElectricCapacitance) Subtract(other *ElectricCapacitance) *ElectricCapacitance {
	return &ElectricCapacitance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricCapacitance by the given ElectricCapacitance and returns the result.
// The result is a new ElectricCapacitance instance with the product of the values.
//
// Parameters:
//    other: The ElectricCapacitance to multiply with the current ElectricCapacitance.
//
// Returns:
//    *ElectricCapacitance: A new ElectricCapacitance instance representing the product of both ElectricCapacitance.
func (a *ElectricCapacitance) Multiply(other *ElectricCapacitance) *ElectricCapacitance {
	return &ElectricCapacitance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricCapacitance by the given ElectricCapacitance and returns the result.
// The result is a new ElectricCapacitance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricCapacitance to divide the current ElectricCapacitance by.
//
// Returns:
//    *ElectricCapacitance: A new ElectricCapacitance instance representing the quotient of both ElectricCapacitance.
func (a *ElectricCapacitance) Divide(other *ElectricCapacitance) *ElectricCapacitance {
	return &ElectricCapacitance{value: a.value / other.BaseValue()}
}

// GetElectricCapacitanceAbbreviation gets the unit abbreviation.
func GetElectricCapacitanceAbbreviation(unit ElectricCapacitanceUnits) string {
	switch unit { 
	case ElectricCapacitanceFarad:
		return "F" 
	case ElectricCapacitancePicofarad:
		return "pF" 
	case ElectricCapacitanceNanofarad:
		return "nF" 
	case ElectricCapacitanceMicrofarad:
		return "μF" 
	case ElectricCapacitanceMillifarad:
		return "mF" 
	case ElectricCapacitanceKilofarad:
		return "kF" 
	case ElectricCapacitanceMegafarad:
		return "MF" 
	default:
		return ""
	}
}