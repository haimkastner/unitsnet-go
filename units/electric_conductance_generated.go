// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricConductanceUnits defines various units of ElectricConductance.
type ElectricConductanceUnits string

const (
    
        // 
        ElectricConductanceSiemens ElectricConductanceUnits = "Siemens"
        // 
        ElectricConductanceNanosiemens ElectricConductanceUnits = "Nanosiemens"
        // 
        ElectricConductanceMicrosiemens ElectricConductanceUnits = "Microsiemens"
        // 
        ElectricConductanceMillisiemens ElectricConductanceUnits = "Millisiemens"
        // 
        ElectricConductanceKilosiemens ElectricConductanceUnits = "Kilosiemens"
)

// ElectricConductanceDto represents a ElectricConductance measurement with a numerical value and its corresponding unit.
type ElectricConductanceDto struct {
    // Value is the numerical representation of the ElectricConductance.
	Value float64
    // Unit specifies the unit of measurement for the ElectricConductance, as defined in the ElectricConductanceUnits enumeration.
	Unit  ElectricConductanceUnits
}

// ElectricConductanceDtoFactory groups methods for creating and serializing ElectricConductanceDto objects.
type ElectricConductanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricConductanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricConductanceDtoFactory) FromJSON(data []byte) (*ElectricConductanceDto, error) {
	a := ElectricConductanceDto{}

    // Parse JSON into ElectricConductanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricConductanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricConductanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ElectricConductance represents a measurement in a of ElectricConductance.
//
// The electrical conductance of an electrical conductor is a measure of the easeness to pass an electric current through that conductor.
type ElectricConductance struct {
	// value is the base measurement stored internally.
	value       float64
    
    siemensLazy *float64 
    nanosiemensLazy *float64 
    microsiemensLazy *float64 
    millisiemensLazy *float64 
    kilosiemensLazy *float64 
}

// ElectricConductanceFactory groups methods for creating ElectricConductance instances.
type ElectricConductanceFactory struct{}

// CreateElectricConductance creates a new ElectricConductance instance from the given value and unit.
func (uf ElectricConductanceFactory) CreateElectricConductance(value float64, unit ElectricConductanceUnits) (*ElectricConductance, error) {
	return newElectricConductance(value, unit)
}

// FromDto converts a ElectricConductanceDto to a ElectricConductance instance.
func (uf ElectricConductanceFactory) FromDto(dto ElectricConductanceDto) (*ElectricConductance, error) {
	return newElectricConductance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricConductance instance.
func (uf ElectricConductanceFactory) FromDtoJSON(data []byte) (*ElectricConductance, error) {
	unitDto, err := ElectricConductanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricConductanceDto from JSON: %w", err)
	}
	return ElectricConductanceFactory{}.FromDto(*unitDto)
}


// FromSiemens creates a new ElectricConductance instance from a value in Siemens.
func (uf ElectricConductanceFactory) FromSiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceSiemens)
}

// FromNanosiemens creates a new ElectricConductance instance from a value in Nanosiemens.
func (uf ElectricConductanceFactory) FromNanosiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceNanosiemens)
}

// FromMicrosiemens creates a new ElectricConductance instance from a value in Microsiemens.
func (uf ElectricConductanceFactory) FromMicrosiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceMicrosiemens)
}

// FromMillisiemens creates a new ElectricConductance instance from a value in Millisiemens.
func (uf ElectricConductanceFactory) FromMillisiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceMillisiemens)
}

// FromKilosiemens creates a new ElectricConductance instance from a value in Kilosiemens.
func (uf ElectricConductanceFactory) FromKilosiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceKilosiemens)
}


// newElectricConductance creates a new ElectricConductance.
func newElectricConductance(value float64, fromUnit ElectricConductanceUnits) (*ElectricConductance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricConductance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricConductance in Siemens unit (the base/default quantity).
func (a *ElectricConductance) BaseValue() float64 {
	return a.value
}


// Siemens returns the ElectricConductance value in Siemens.
//
// 
func (a *ElectricConductance) Siemens() float64 {
	if a.siemensLazy != nil {
		return *a.siemensLazy
	}
	siemens := a.convertFromBase(ElectricConductanceSiemens)
	a.siemensLazy = &siemens
	return siemens
}

// Nanosiemens returns the ElectricConductance value in Nanosiemens.
//
// 
func (a *ElectricConductance) Nanosiemens() float64 {
	if a.nanosiemensLazy != nil {
		return *a.nanosiemensLazy
	}
	nanosiemens := a.convertFromBase(ElectricConductanceNanosiemens)
	a.nanosiemensLazy = &nanosiemens
	return nanosiemens
}

// Microsiemens returns the ElectricConductance value in Microsiemens.
//
// 
func (a *ElectricConductance) Microsiemens() float64 {
	if a.microsiemensLazy != nil {
		return *a.microsiemensLazy
	}
	microsiemens := a.convertFromBase(ElectricConductanceMicrosiemens)
	a.microsiemensLazy = &microsiemens
	return microsiemens
}

// Millisiemens returns the ElectricConductance value in Millisiemens.
//
// 
func (a *ElectricConductance) Millisiemens() float64 {
	if a.millisiemensLazy != nil {
		return *a.millisiemensLazy
	}
	millisiemens := a.convertFromBase(ElectricConductanceMillisiemens)
	a.millisiemensLazy = &millisiemens
	return millisiemens
}

// Kilosiemens returns the ElectricConductance value in Kilosiemens.
//
// 
func (a *ElectricConductance) Kilosiemens() float64 {
	if a.kilosiemensLazy != nil {
		return *a.kilosiemensLazy
	}
	kilosiemens := a.convertFromBase(ElectricConductanceKilosiemens)
	a.kilosiemensLazy = &kilosiemens
	return kilosiemens
}


// ToDto creates a ElectricConductanceDto representation from the ElectricConductance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricConductance) ToDto(holdInUnit *ElectricConductanceUnits) ElectricConductanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricConductanceSiemens // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricConductanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricConductance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricConductance) ToDtoJSON(holdInUnit *ElectricConductanceUnits) ([]byte, error) {
	// Convert to ElectricConductanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricConductance to a specific unit value.
// The function uses the provided unit type (ElectricConductanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricConductance) Convert(toUnit ElectricConductanceUnits) float64 {
	switch toUnit { 
    case ElectricConductanceSiemens:
		return a.Siemens()
    case ElectricConductanceNanosiemens:
		return a.Nanosiemens()
    case ElectricConductanceMicrosiemens:
		return a.Microsiemens()
    case ElectricConductanceMillisiemens:
		return a.Millisiemens()
    case ElectricConductanceKilosiemens:
		return a.Kilosiemens()
	default:
		return math.NaN()
	}
}

func (a *ElectricConductance) convertFromBase(toUnit ElectricConductanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricConductanceSiemens:
		return (value) 
	case ElectricConductanceNanosiemens:
		return ((value) / 1e-09) 
	case ElectricConductanceMicrosiemens:
		return ((value) / 1e-06) 
	case ElectricConductanceMillisiemens:
		return ((value) / 0.001) 
	case ElectricConductanceKilosiemens:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricConductance) convertToBase(value float64, fromUnit ElectricConductanceUnits) float64 {
	switch fromUnit { 
	case ElectricConductanceSiemens:
		return (value) 
	case ElectricConductanceNanosiemens:
		return ((value) * 1e-09) 
	case ElectricConductanceMicrosiemens:
		return ((value) * 1e-06) 
	case ElectricConductanceMillisiemens:
		return ((value) * 0.001) 
	case ElectricConductanceKilosiemens:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricConductance in the default unit (Siemens),
// formatted to two decimal places.
func (a ElectricConductance) String() string {
	return a.ToString(ElectricConductanceSiemens, 2)
}

// ToString formats the ElectricConductance value as a string with the specified unit and fractional digits.
// It converts the ElectricConductance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricConductance value will be converted (e.g., Siemens))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricConductance with the unit abbreviation.
func (a *ElectricConductance) ToString(unit ElectricConductanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricConductance) getUnitAbbreviation(unit ElectricConductanceUnits) string {
	switch unit { 
	case ElectricConductanceSiemens:
		return "S" 
	case ElectricConductanceNanosiemens:
		return "nS" 
	case ElectricConductanceMicrosiemens:
		return "μS" 
	case ElectricConductanceMillisiemens:
		return "mS" 
	case ElectricConductanceKilosiemens:
		return "kS" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricConductance is equal to the current ElectricConductance.
//
// Parameters:
//    other: The ElectricConductance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricConductance are equal, false otherwise.
func (a *ElectricConductance) Equals(other *ElectricConductance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricConductance with another ElectricConductance.
// It returns -1 if the current ElectricConductance is less than the other ElectricConductance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricConductance to compare against.
//
// Returns:
//    int: -1 if the current ElectricConductance is less, 1 if greater, and 0 if equal.
func (a *ElectricConductance) CompareTo(other *ElectricConductance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricConductance to the current ElectricConductance and returns the result.
// The result is a new ElectricConductance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricConductance to add to the current ElectricConductance.
//
// Returns:
//    *ElectricConductance: A new ElectricConductance instance representing the sum of both ElectricConductance.
func (a *ElectricConductance) Add(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricConductance from the current ElectricConductance and returns the result.
// The result is a new ElectricConductance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricConductance to subtract from the current ElectricConductance.
//
// Returns:
//    *ElectricConductance: A new ElectricConductance instance representing the difference of both ElectricConductance.
func (a *ElectricConductance) Subtract(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricConductance by the given ElectricConductance and returns the result.
// The result is a new ElectricConductance instance with the product of the values.
//
// Parameters:
//    other: The ElectricConductance to multiply with the current ElectricConductance.
//
// Returns:
//    *ElectricConductance: A new ElectricConductance instance representing the product of both ElectricConductance.
func (a *ElectricConductance) Multiply(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricConductance by the given ElectricConductance and returns the result.
// The result is a new ElectricConductance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricConductance to divide the current ElectricConductance by.
//
// Returns:
//    *ElectricConductance: A new ElectricConductance instance representing the quotient of both ElectricConductance.
func (a *ElectricConductance) Divide(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value / other.BaseValue()}
}