// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricResistanceUnits defines various units of ElectricResistance.
type ElectricResistanceUnits string

const (
    
        // 
        ElectricResistanceOhm ElectricResistanceUnits = "Ohm"
        // 
        ElectricResistanceMicroohm ElectricResistanceUnits = "Microohm"
        // 
        ElectricResistanceMilliohm ElectricResistanceUnits = "Milliohm"
        // 
        ElectricResistanceKiloohm ElectricResistanceUnits = "Kiloohm"
        // 
        ElectricResistanceMegaohm ElectricResistanceUnits = "Megaohm"
        // 
        ElectricResistanceGigaohm ElectricResistanceUnits = "Gigaohm"
        // 
        ElectricResistanceTeraohm ElectricResistanceUnits = "Teraohm"
)

// ElectricResistanceDto represents a ElectricResistance measurement with a numerical value and its corresponding unit.
type ElectricResistanceDto struct {
    // Value is the numerical representation of the ElectricResistance.
	Value float64
    // Unit specifies the unit of measurement for the ElectricResistance, as defined in the ElectricResistanceUnits enumeration.
	Unit  ElectricResistanceUnits
}

// ElectricResistanceDtoFactory groups methods for creating and serializing ElectricResistanceDto objects.
type ElectricResistanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricResistanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricResistanceDtoFactory) FromJSON(data []byte) (*ElectricResistanceDto, error) {
	a := ElectricResistanceDto{}

    // Parse JSON into ElectricResistanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricResistanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricResistanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// ElectricResistance represents a measurement in a of ElectricResistance.
//
// The electrical resistance of an electrical conductor is the opposition to the passage of an electric current through that conductor.
type ElectricResistance struct {
	// value is the base measurement stored internally.
	value       float64
    
    ohmsLazy *float64 
    microohmsLazy *float64 
    milliohmsLazy *float64 
    kiloohmsLazy *float64 
    megaohmsLazy *float64 
    gigaohmsLazy *float64 
    teraohmsLazy *float64 
}

// ElectricResistanceFactory groups methods for creating ElectricResistance instances.
type ElectricResistanceFactory struct{}

// CreateElectricResistance creates a new ElectricResistance instance from the given value and unit.
func (uf ElectricResistanceFactory) CreateElectricResistance(value float64, unit ElectricResistanceUnits) (*ElectricResistance, error) {
	return newElectricResistance(value, unit)
}

// FromDto converts a ElectricResistanceDto to a ElectricResistance instance.
func (uf ElectricResistanceFactory) FromDto(dto ElectricResistanceDto) (*ElectricResistance, error) {
	return newElectricResistance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricResistance instance.
func (uf ElectricResistanceFactory) FromDtoJSON(data []byte) (*ElectricResistance, error) {
	unitDto, err := ElectricResistanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricResistanceDto from JSON: %w", err)
	}
	return ElectricResistanceFactory{}.FromDto(*unitDto)
}


// FromOhms creates a new ElectricResistance instance from a value in Ohms.
func (uf ElectricResistanceFactory) FromOhms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceOhm)
}

// FromMicroohms creates a new ElectricResistance instance from a value in Microohms.
func (uf ElectricResistanceFactory) FromMicroohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMicroohm)
}

// FromMilliohms creates a new ElectricResistance instance from a value in Milliohms.
func (uf ElectricResistanceFactory) FromMilliohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMilliohm)
}

// FromKiloohms creates a new ElectricResistance instance from a value in Kiloohms.
func (uf ElectricResistanceFactory) FromKiloohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceKiloohm)
}

// FromMegaohms creates a new ElectricResistance instance from a value in Megaohms.
func (uf ElectricResistanceFactory) FromMegaohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMegaohm)
}

// FromGigaohms creates a new ElectricResistance instance from a value in Gigaohms.
func (uf ElectricResistanceFactory) FromGigaohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceGigaohm)
}

// FromTeraohms creates a new ElectricResistance instance from a value in Teraohms.
func (uf ElectricResistanceFactory) FromTeraohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceTeraohm)
}


// newElectricResistance creates a new ElectricResistance.
func newElectricResistance(value float64, fromUnit ElectricResistanceUnits) (*ElectricResistance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricResistance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricResistance in Ohm unit (the base/default quantity).
func (a *ElectricResistance) BaseValue() float64 {
	return a.value
}


// Ohms returns the ElectricResistance value in Ohms.
//
// 
func (a *ElectricResistance) Ohms() float64 {
	if a.ohmsLazy != nil {
		return *a.ohmsLazy
	}
	ohms := a.convertFromBase(ElectricResistanceOhm)
	a.ohmsLazy = &ohms
	return ohms
}

// Microohms returns the ElectricResistance value in Microohms.
//
// 
func (a *ElectricResistance) Microohms() float64 {
	if a.microohmsLazy != nil {
		return *a.microohmsLazy
	}
	microohms := a.convertFromBase(ElectricResistanceMicroohm)
	a.microohmsLazy = &microohms
	return microohms
}

// Milliohms returns the ElectricResistance value in Milliohms.
//
// 
func (a *ElectricResistance) Milliohms() float64 {
	if a.milliohmsLazy != nil {
		return *a.milliohmsLazy
	}
	milliohms := a.convertFromBase(ElectricResistanceMilliohm)
	a.milliohmsLazy = &milliohms
	return milliohms
}

// Kiloohms returns the ElectricResistance value in Kiloohms.
//
// 
func (a *ElectricResistance) Kiloohms() float64 {
	if a.kiloohmsLazy != nil {
		return *a.kiloohmsLazy
	}
	kiloohms := a.convertFromBase(ElectricResistanceKiloohm)
	a.kiloohmsLazy = &kiloohms
	return kiloohms
}

// Megaohms returns the ElectricResistance value in Megaohms.
//
// 
func (a *ElectricResistance) Megaohms() float64 {
	if a.megaohmsLazy != nil {
		return *a.megaohmsLazy
	}
	megaohms := a.convertFromBase(ElectricResistanceMegaohm)
	a.megaohmsLazy = &megaohms
	return megaohms
}

// Gigaohms returns the ElectricResistance value in Gigaohms.
//
// 
func (a *ElectricResistance) Gigaohms() float64 {
	if a.gigaohmsLazy != nil {
		return *a.gigaohmsLazy
	}
	gigaohms := a.convertFromBase(ElectricResistanceGigaohm)
	a.gigaohmsLazy = &gigaohms
	return gigaohms
}

// Teraohms returns the ElectricResistance value in Teraohms.
//
// 
func (a *ElectricResistance) Teraohms() float64 {
	if a.teraohmsLazy != nil {
		return *a.teraohmsLazy
	}
	teraohms := a.convertFromBase(ElectricResistanceTeraohm)
	a.teraohmsLazy = &teraohms
	return teraohms
}


// ToDto creates a ElectricResistanceDto representation from the ElectricResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricResistance) ToDto(holdInUnit *ElectricResistanceUnits) ElectricResistanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricResistanceOhm // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricResistanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricResistance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricResistance) ToDtoJSON(holdInUnit *ElectricResistanceUnits) ([]byte, error) {
	// Convert to ElectricResistanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricResistance to a specific unit value.
// The function uses the provided unit type (ElectricResistanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricResistance) Convert(toUnit ElectricResistanceUnits) float64 {
	switch toUnit { 
    case ElectricResistanceOhm:
		return a.Ohms()
    case ElectricResistanceMicroohm:
		return a.Microohms()
    case ElectricResistanceMilliohm:
		return a.Milliohms()
    case ElectricResistanceKiloohm:
		return a.Kiloohms()
    case ElectricResistanceMegaohm:
		return a.Megaohms()
    case ElectricResistanceGigaohm:
		return a.Gigaohms()
    case ElectricResistanceTeraohm:
		return a.Teraohms()
	default:
		return math.NaN()
	}
}

func (a *ElectricResistance) convertFromBase(toUnit ElectricResistanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricResistanceOhm:
		return (value) 
	case ElectricResistanceMicroohm:
		return ((value) / 1e-06) 
	case ElectricResistanceMilliohm:
		return ((value) / 0.001) 
	case ElectricResistanceKiloohm:
		return ((value) / 1000.0) 
	case ElectricResistanceMegaohm:
		return ((value) / 1000000.0) 
	case ElectricResistanceGigaohm:
		return ((value) / 1000000000.0) 
	case ElectricResistanceTeraohm:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricResistance) convertToBase(value float64, fromUnit ElectricResistanceUnits) float64 {
	switch fromUnit { 
	case ElectricResistanceOhm:
		return (value) 
	case ElectricResistanceMicroohm:
		return ((value) * 1e-06) 
	case ElectricResistanceMilliohm:
		return ((value) * 0.001) 
	case ElectricResistanceKiloohm:
		return ((value) * 1000.0) 
	case ElectricResistanceMegaohm:
		return ((value) * 1000000.0) 
	case ElectricResistanceGigaohm:
		return ((value) * 1000000000.0) 
	case ElectricResistanceTeraohm:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricResistance in the default unit (Ohm),
// formatted to two decimal places.
func (a ElectricResistance) String() string {
	return a.ToString(ElectricResistanceOhm, 2)
}

// ToString formats the ElectricResistance value as a string with the specified unit and fractional digits.
// It converts the ElectricResistance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricResistance value will be converted (e.g., Ohm))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricResistance with the unit abbreviation.
func (a *ElectricResistance) ToString(unit ElectricResistanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricResistance) getUnitAbbreviation(unit ElectricResistanceUnits) string {
	switch unit { 
	case ElectricResistanceOhm:
		return "Ω" 
	case ElectricResistanceMicroohm:
		return "μΩ" 
	case ElectricResistanceMilliohm:
		return "mΩ" 
	case ElectricResistanceKiloohm:
		return "kΩ" 
	case ElectricResistanceMegaohm:
		return "MΩ" 
	case ElectricResistanceGigaohm:
		return "GΩ" 
	case ElectricResistanceTeraohm:
		return "TΩ" 
	default:
		return ""
	}
}

// Equals checks if the given ElectricResistance is equal to the current ElectricResistance.
//
// Parameters:
//    other: The ElectricResistance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricResistance are equal, false otherwise.
func (a *ElectricResistance) Equals(other *ElectricResistance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricResistance with another ElectricResistance.
// It returns -1 if the current ElectricResistance is less than the other ElectricResistance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricResistance to compare against.
//
// Returns:
//    int: -1 if the current ElectricResistance is less, 1 if greater, and 0 if equal.
func (a *ElectricResistance) CompareTo(other *ElectricResistance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricResistance to the current ElectricResistance and returns the result.
// The result is a new ElectricResistance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricResistance to add to the current ElectricResistance.
//
// Returns:
//    *ElectricResistance: A new ElectricResistance instance representing the sum of both ElectricResistance.
func (a *ElectricResistance) Add(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricResistance from the current ElectricResistance and returns the result.
// The result is a new ElectricResistance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricResistance to subtract from the current ElectricResistance.
//
// Returns:
//    *ElectricResistance: A new ElectricResistance instance representing the difference of both ElectricResistance.
func (a *ElectricResistance) Subtract(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricResistance by the given ElectricResistance and returns the result.
// The result is a new ElectricResistance instance with the product of the values.
//
// Parameters:
//    other: The ElectricResistance to multiply with the current ElectricResistance.
//
// Returns:
//    *ElectricResistance: A new ElectricResistance instance representing the product of both ElectricResistance.
func (a *ElectricResistance) Multiply(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricResistance by the given ElectricResistance and returns the result.
// The result is a new ElectricResistance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricResistance to divide the current ElectricResistance by.
//
// Returns:
//    *ElectricResistance: A new ElectricResistance instance representing the quotient of both ElectricResistance.
func (a *ElectricResistance) Divide(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value / other.BaseValue()}
}