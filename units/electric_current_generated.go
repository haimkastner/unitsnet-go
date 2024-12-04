// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentUnits defines various units of ElectricCurrent.
type ElectricCurrentUnits string

const (
    
        // 
        ElectricCurrentAmpere ElectricCurrentUnits = "Ampere"
        // 
        ElectricCurrentFemtoampere ElectricCurrentUnits = "Femtoampere"
        // 
        ElectricCurrentPicoampere ElectricCurrentUnits = "Picoampere"
        // 
        ElectricCurrentNanoampere ElectricCurrentUnits = "Nanoampere"
        // 
        ElectricCurrentMicroampere ElectricCurrentUnits = "Microampere"
        // 
        ElectricCurrentMilliampere ElectricCurrentUnits = "Milliampere"
        // 
        ElectricCurrentCentiampere ElectricCurrentUnits = "Centiampere"
        // 
        ElectricCurrentKiloampere ElectricCurrentUnits = "Kiloampere"
        // 
        ElectricCurrentMegaampere ElectricCurrentUnits = "Megaampere"
)

// ElectricCurrentDto represents a ElectricCurrent measurement with a numerical value and its corresponding unit.
type ElectricCurrentDto struct {
    // Value is the numerical representation of the ElectricCurrent.
	Value float64
    // Unit specifies the unit of measurement for the ElectricCurrent, as defined in the ElectricCurrentUnits enumeration.
	Unit  ElectricCurrentUnits
}

// ElectricCurrentDtoFactory groups methods for creating and serializing ElectricCurrentDto objects.
type ElectricCurrentDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrentDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricCurrentDtoFactory) FromJSON(data []byte) (*ElectricCurrentDto, error) {
	a := ElectricCurrentDto{}

    // Parse JSON into ElectricCurrentDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricCurrentDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricCurrentDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ElectricCurrent represents a measurement in a of ElectricCurrent.
//
// An electric current is a flow of electric charge. In electric circuits this charge is often carried by moving electrons in a wire. It can also be carried by ions in an electrolyte, or by both ions and electrons such as in a plasma.
type ElectricCurrent struct {
	// value is the base measurement stored internally.
	value       float64
    
    amperesLazy *float64 
    femtoamperesLazy *float64 
    picoamperesLazy *float64 
    nanoamperesLazy *float64 
    microamperesLazy *float64 
    milliamperesLazy *float64 
    centiamperesLazy *float64 
    kiloamperesLazy *float64 
    megaamperesLazy *float64 
}

// ElectricCurrentFactory groups methods for creating ElectricCurrent instances.
type ElectricCurrentFactory struct{}

// CreateElectricCurrent creates a new ElectricCurrent instance from the given value and unit.
func (uf ElectricCurrentFactory) CreateElectricCurrent(value float64, unit ElectricCurrentUnits) (*ElectricCurrent, error) {
	return newElectricCurrent(value, unit)
}

// FromDto converts a ElectricCurrentDto to a ElectricCurrent instance.
func (uf ElectricCurrentFactory) FromDto(dto ElectricCurrentDto) (*ElectricCurrent, error) {
	return newElectricCurrent(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricCurrent instance.
func (uf ElectricCurrentFactory) FromDtoJSON(data []byte) (*ElectricCurrent, error) {
	unitDto, err := ElectricCurrentDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricCurrentDto from JSON: %w", err)
	}
	return ElectricCurrentFactory{}.FromDto(*unitDto)
}


// FromAmperes creates a new ElectricCurrent instance from a value in Amperes.
func (uf ElectricCurrentFactory) FromAmperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentAmpere)
}

// FromFemtoamperes creates a new ElectricCurrent instance from a value in Femtoamperes.
func (uf ElectricCurrentFactory) FromFemtoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentFemtoampere)
}

// FromPicoamperes creates a new ElectricCurrent instance from a value in Picoamperes.
func (uf ElectricCurrentFactory) FromPicoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentPicoampere)
}

// FromNanoamperes creates a new ElectricCurrent instance from a value in Nanoamperes.
func (uf ElectricCurrentFactory) FromNanoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentNanoampere)
}

// FromMicroamperes creates a new ElectricCurrent instance from a value in Microamperes.
func (uf ElectricCurrentFactory) FromMicroamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMicroampere)
}

// FromMilliamperes creates a new ElectricCurrent instance from a value in Milliamperes.
func (uf ElectricCurrentFactory) FromMilliamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMilliampere)
}

// FromCentiamperes creates a new ElectricCurrent instance from a value in Centiamperes.
func (uf ElectricCurrentFactory) FromCentiamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentCentiampere)
}

// FromKiloamperes creates a new ElectricCurrent instance from a value in Kiloamperes.
func (uf ElectricCurrentFactory) FromKiloamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentKiloampere)
}

// FromMegaamperes creates a new ElectricCurrent instance from a value in Megaamperes.
func (uf ElectricCurrentFactory) FromMegaamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMegaampere)
}


// newElectricCurrent creates a new ElectricCurrent.
func newElectricCurrent(value float64, fromUnit ElectricCurrentUnits) (*ElectricCurrent, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrent{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrent in Ampere unit (the base/default quantity).
func (a *ElectricCurrent) BaseValue() float64 {
	return a.value
}


// Amperes returns the ElectricCurrent value in Amperes.
//
// 
func (a *ElectricCurrent) Amperes() float64 {
	if a.amperesLazy != nil {
		return *a.amperesLazy
	}
	amperes := a.convertFromBase(ElectricCurrentAmpere)
	a.amperesLazy = &amperes
	return amperes
}

// Femtoamperes returns the ElectricCurrent value in Femtoamperes.
//
// 
func (a *ElectricCurrent) Femtoamperes() float64 {
	if a.femtoamperesLazy != nil {
		return *a.femtoamperesLazy
	}
	femtoamperes := a.convertFromBase(ElectricCurrentFemtoampere)
	a.femtoamperesLazy = &femtoamperes
	return femtoamperes
}

// Picoamperes returns the ElectricCurrent value in Picoamperes.
//
// 
func (a *ElectricCurrent) Picoamperes() float64 {
	if a.picoamperesLazy != nil {
		return *a.picoamperesLazy
	}
	picoamperes := a.convertFromBase(ElectricCurrentPicoampere)
	a.picoamperesLazy = &picoamperes
	return picoamperes
}

// Nanoamperes returns the ElectricCurrent value in Nanoamperes.
//
// 
func (a *ElectricCurrent) Nanoamperes() float64 {
	if a.nanoamperesLazy != nil {
		return *a.nanoamperesLazy
	}
	nanoamperes := a.convertFromBase(ElectricCurrentNanoampere)
	a.nanoamperesLazy = &nanoamperes
	return nanoamperes
}

// Microamperes returns the ElectricCurrent value in Microamperes.
//
// 
func (a *ElectricCurrent) Microamperes() float64 {
	if a.microamperesLazy != nil {
		return *a.microamperesLazy
	}
	microamperes := a.convertFromBase(ElectricCurrentMicroampere)
	a.microamperesLazy = &microamperes
	return microamperes
}

// Milliamperes returns the ElectricCurrent value in Milliamperes.
//
// 
func (a *ElectricCurrent) Milliamperes() float64 {
	if a.milliamperesLazy != nil {
		return *a.milliamperesLazy
	}
	milliamperes := a.convertFromBase(ElectricCurrentMilliampere)
	a.milliamperesLazy = &milliamperes
	return milliamperes
}

// Centiamperes returns the ElectricCurrent value in Centiamperes.
//
// 
func (a *ElectricCurrent) Centiamperes() float64 {
	if a.centiamperesLazy != nil {
		return *a.centiamperesLazy
	}
	centiamperes := a.convertFromBase(ElectricCurrentCentiampere)
	a.centiamperesLazy = &centiamperes
	return centiamperes
}

// Kiloamperes returns the ElectricCurrent value in Kiloamperes.
//
// 
func (a *ElectricCurrent) Kiloamperes() float64 {
	if a.kiloamperesLazy != nil {
		return *a.kiloamperesLazy
	}
	kiloamperes := a.convertFromBase(ElectricCurrentKiloampere)
	a.kiloamperesLazy = &kiloamperes
	return kiloamperes
}

// Megaamperes returns the ElectricCurrent value in Megaamperes.
//
// 
func (a *ElectricCurrent) Megaamperes() float64 {
	if a.megaamperesLazy != nil {
		return *a.megaamperesLazy
	}
	megaamperes := a.convertFromBase(ElectricCurrentMegaampere)
	a.megaamperesLazy = &megaamperes
	return megaamperes
}


// ToDto creates a ElectricCurrentDto representation from the ElectricCurrent instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ampere by default.
func (a *ElectricCurrent) ToDto(holdInUnit *ElectricCurrentUnits) ElectricCurrentDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentAmpere // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricCurrent instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ampere by default.
func (a *ElectricCurrent) ToDtoJSON(holdInUnit *ElectricCurrentUnits) ([]byte, error) {
	// Convert to ElectricCurrentDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricCurrent to a specific unit value.
// The function uses the provided unit type (ElectricCurrentUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricCurrent) Convert(toUnit ElectricCurrentUnits) float64 {
	switch toUnit { 
    case ElectricCurrentAmpere:
		return a.Amperes()
    case ElectricCurrentFemtoampere:
		return a.Femtoamperes()
    case ElectricCurrentPicoampere:
		return a.Picoamperes()
    case ElectricCurrentNanoampere:
		return a.Nanoamperes()
    case ElectricCurrentMicroampere:
		return a.Microamperes()
    case ElectricCurrentMilliampere:
		return a.Milliamperes()
    case ElectricCurrentCentiampere:
		return a.Centiamperes()
    case ElectricCurrentKiloampere:
		return a.Kiloamperes()
    case ElectricCurrentMegaampere:
		return a.Megaamperes()
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrent) convertFromBase(toUnit ElectricCurrentUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentAmpere:
		return (value) 
	case ElectricCurrentFemtoampere:
		return ((value) / 1e-15) 
	case ElectricCurrentPicoampere:
		return ((value) / 1e-12) 
	case ElectricCurrentNanoampere:
		return ((value) / 1e-09) 
	case ElectricCurrentMicroampere:
		return ((value) / 1e-06) 
	case ElectricCurrentMilliampere:
		return ((value) / 0.001) 
	case ElectricCurrentCentiampere:
		return ((value) / 0.01) 
	case ElectricCurrentKiloampere:
		return ((value) / 1000.0) 
	case ElectricCurrentMegaampere:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrent) convertToBase(value float64, fromUnit ElectricCurrentUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentAmpere:
		return (value) 
	case ElectricCurrentFemtoampere:
		return ((value) * 1e-15) 
	case ElectricCurrentPicoampere:
		return ((value) * 1e-12) 
	case ElectricCurrentNanoampere:
		return ((value) * 1e-09) 
	case ElectricCurrentMicroampere:
		return ((value) * 1e-06) 
	case ElectricCurrentMilliampere:
		return ((value) * 0.001) 
	case ElectricCurrentCentiampere:
		return ((value) * 0.01) 
	case ElectricCurrentKiloampere:
		return ((value) * 1000.0) 
	case ElectricCurrentMegaampere:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricCurrent in the default unit (Ampere),
// formatted to two decimal places.
func (a ElectricCurrent) String() string {
	return a.ToString(ElectricCurrentAmpere, 2)
}

// ToString formats the ElectricCurrent value as a string with the specified unit and fractional digits.
// It converts the ElectricCurrent to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricCurrent value will be converted (e.g., Ampere))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricCurrent with the unit abbreviation.
func (a *ElectricCurrent) ToString(unit ElectricCurrentUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCurrent) getUnitAbbreviation(unit ElectricCurrentUnits) string {
	switch unit { 
	case ElectricCurrentAmpere:
		return "A" 
	case ElectricCurrentFemtoampere:
		return "fA" 
	case ElectricCurrentPicoampere:
		return "pA" 
	case ElectricCurrentNanoampere:
		return "nA" 
	case ElectricCurrentMicroampere:
		return "μA" 
	case ElectricCurrentMilliampere:
		return "mA" 
	case ElectricCurrentCentiampere:
		return "cA" 
	case ElectricCurrentKiloampere:
		return "kA" 
	case ElectricCurrentMegaampere:
		return "MA" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricCurrent is equal to the current ElectricCurrent.
//
// Parameters:
//    other: The ElectricCurrent to compare against.
//
// Returns:
//    bool: Returns true if both ElectricCurrent are equal, false otherwise.
func (a *ElectricCurrent) Equals(other *ElectricCurrent) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricCurrent with another ElectricCurrent.
// It returns -1 if the current ElectricCurrent is less than the other ElectricCurrent, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricCurrent to compare against.
//
// Returns:
//    int: -1 if the current ElectricCurrent is less, 1 if greater, and 0 if equal.
func (a *ElectricCurrent) CompareTo(other *ElectricCurrent) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricCurrent to the current ElectricCurrent and returns the result.
// The result is a new ElectricCurrent instance with the sum of the values.
//
// Parameters:
//    other: The ElectricCurrent to add to the current ElectricCurrent.
//
// Returns:
//    *ElectricCurrent: A new ElectricCurrent instance representing the sum of both ElectricCurrent.
func (a *ElectricCurrent) Add(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricCurrent from the current ElectricCurrent and returns the result.
// The result is a new ElectricCurrent instance with the difference of the values.
//
// Parameters:
//    other: The ElectricCurrent to subtract from the current ElectricCurrent.
//
// Returns:
//    *ElectricCurrent: A new ElectricCurrent instance representing the difference of both ElectricCurrent.
func (a *ElectricCurrent) Subtract(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricCurrent by the given ElectricCurrent and returns the result.
// The result is a new ElectricCurrent instance with the product of the values.
//
// Parameters:
//    other: The ElectricCurrent to multiply with the current ElectricCurrent.
//
// Returns:
//    *ElectricCurrent: A new ElectricCurrent instance representing the product of both ElectricCurrent.
func (a *ElectricCurrent) Multiply(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricCurrent by the given ElectricCurrent and returns the result.
// The result is a new ElectricCurrent instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricCurrent to divide the current ElectricCurrent by.
//
// Returns:
//    *ElectricCurrent: A new ElectricCurrent instance representing the quotient of both ElectricCurrent.
func (a *ElectricCurrent) Divide(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value / other.BaseValue()}
}