// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricImpedanceUnits defines various units of ElectricImpedance.
type ElectricImpedanceUnits string

const (
    
        // 
        ElectricImpedanceOhm ElectricImpedanceUnits = "Ohm"
        // 
        ElectricImpedanceNanoohm ElectricImpedanceUnits = "Nanoohm"
        // 
        ElectricImpedanceMicroohm ElectricImpedanceUnits = "Microohm"
        // 
        ElectricImpedanceMilliohm ElectricImpedanceUnits = "Milliohm"
        // 
        ElectricImpedanceKiloohm ElectricImpedanceUnits = "Kiloohm"
        // 
        ElectricImpedanceMegaohm ElectricImpedanceUnits = "Megaohm"
        // 
        ElectricImpedanceGigaohm ElectricImpedanceUnits = "Gigaohm"
        // 
        ElectricImpedanceTeraohm ElectricImpedanceUnits = "Teraohm"
)

// ElectricImpedanceDto represents a ElectricImpedance measurement with a numerical value and its corresponding unit.
type ElectricImpedanceDto struct {
    // Value is the numerical representation of the ElectricImpedance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricImpedance, as defined in the ElectricImpedanceUnits enumeration.
	Unit  ElectricImpedanceUnits `json:"unit"`
}

// ElectricImpedanceDtoFactory groups methods for creating and serializing ElectricImpedanceDto objects.
type ElectricImpedanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricImpedanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricImpedanceDtoFactory) FromJSON(data []byte) (*ElectricImpedanceDto, error) {
	a := ElectricImpedanceDto{}

    // Parse JSON into ElectricImpedanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricImpedanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricImpedanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricImpedance represents a measurement in a of ElectricImpedance.
//
// Electric impedance is the opposition to alternating current presented by the combined effect of resistance and reactance in a circuit. It is defined as the inverse of admittance. The SI unit of impedance is the ohm (symbol Ω).
type ElectricImpedance struct {
	// value is the base measurement stored internally.
	value       float64
    
    ohmsLazy *float64 
    nanoohmsLazy *float64 
    microohmsLazy *float64 
    milliohmsLazy *float64 
    kiloohmsLazy *float64 
    megaohmsLazy *float64 
    gigaohmsLazy *float64 
    teraohmsLazy *float64 
}

// ElectricImpedanceFactory groups methods for creating ElectricImpedance instances.
type ElectricImpedanceFactory struct{}

// CreateElectricImpedance creates a new ElectricImpedance instance from the given value and unit.
func (uf ElectricImpedanceFactory) CreateElectricImpedance(value float64, unit ElectricImpedanceUnits) (*ElectricImpedance, error) {
	return newElectricImpedance(value, unit)
}

// FromDto converts a ElectricImpedanceDto to a ElectricImpedance instance.
func (uf ElectricImpedanceFactory) FromDto(dto ElectricImpedanceDto) (*ElectricImpedance, error) {
	return newElectricImpedance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricImpedance instance.
func (uf ElectricImpedanceFactory) FromDtoJSON(data []byte) (*ElectricImpedance, error) {
	unitDto, err := ElectricImpedanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricImpedanceDto from JSON: %w", err)
	}
	return ElectricImpedanceFactory{}.FromDto(*unitDto)
}


// FromOhms creates a new ElectricImpedance instance from a value in Ohms.
func (uf ElectricImpedanceFactory) FromOhms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceOhm)
}

// FromNanoohms creates a new ElectricImpedance instance from a value in Nanoohms.
func (uf ElectricImpedanceFactory) FromNanoohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceNanoohm)
}

// FromMicroohms creates a new ElectricImpedance instance from a value in Microohms.
func (uf ElectricImpedanceFactory) FromMicroohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceMicroohm)
}

// FromMilliohms creates a new ElectricImpedance instance from a value in Milliohms.
func (uf ElectricImpedanceFactory) FromMilliohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceMilliohm)
}

// FromKiloohms creates a new ElectricImpedance instance from a value in Kiloohms.
func (uf ElectricImpedanceFactory) FromKiloohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceKiloohm)
}

// FromMegaohms creates a new ElectricImpedance instance from a value in Megaohms.
func (uf ElectricImpedanceFactory) FromMegaohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceMegaohm)
}

// FromGigaohms creates a new ElectricImpedance instance from a value in Gigaohms.
func (uf ElectricImpedanceFactory) FromGigaohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceGigaohm)
}

// FromTeraohms creates a new ElectricImpedance instance from a value in Teraohms.
func (uf ElectricImpedanceFactory) FromTeraohms(value float64) (*ElectricImpedance, error) {
	return newElectricImpedance(value, ElectricImpedanceTeraohm)
}


// newElectricImpedance creates a new ElectricImpedance.
func newElectricImpedance(value float64, fromUnit ElectricImpedanceUnits) (*ElectricImpedance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricImpedance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricImpedance in Ohm unit (the base/default quantity).
func (a *ElectricImpedance) BaseValue() float64 {
	return a.value
}


// Ohms returns the ElectricImpedance value in Ohms.
//
// 
func (a *ElectricImpedance) Ohms() float64 {
	if a.ohmsLazy != nil {
		return *a.ohmsLazy
	}
	ohms := a.convertFromBase(ElectricImpedanceOhm)
	a.ohmsLazy = &ohms
	return ohms
}

// Nanoohms returns the ElectricImpedance value in Nanoohms.
//
// 
func (a *ElectricImpedance) Nanoohms() float64 {
	if a.nanoohmsLazy != nil {
		return *a.nanoohmsLazy
	}
	nanoohms := a.convertFromBase(ElectricImpedanceNanoohm)
	a.nanoohmsLazy = &nanoohms
	return nanoohms
}

// Microohms returns the ElectricImpedance value in Microohms.
//
// 
func (a *ElectricImpedance) Microohms() float64 {
	if a.microohmsLazy != nil {
		return *a.microohmsLazy
	}
	microohms := a.convertFromBase(ElectricImpedanceMicroohm)
	a.microohmsLazy = &microohms
	return microohms
}

// Milliohms returns the ElectricImpedance value in Milliohms.
//
// 
func (a *ElectricImpedance) Milliohms() float64 {
	if a.milliohmsLazy != nil {
		return *a.milliohmsLazy
	}
	milliohms := a.convertFromBase(ElectricImpedanceMilliohm)
	a.milliohmsLazy = &milliohms
	return milliohms
}

// Kiloohms returns the ElectricImpedance value in Kiloohms.
//
// 
func (a *ElectricImpedance) Kiloohms() float64 {
	if a.kiloohmsLazy != nil {
		return *a.kiloohmsLazy
	}
	kiloohms := a.convertFromBase(ElectricImpedanceKiloohm)
	a.kiloohmsLazy = &kiloohms
	return kiloohms
}

// Megaohms returns the ElectricImpedance value in Megaohms.
//
// 
func (a *ElectricImpedance) Megaohms() float64 {
	if a.megaohmsLazy != nil {
		return *a.megaohmsLazy
	}
	megaohms := a.convertFromBase(ElectricImpedanceMegaohm)
	a.megaohmsLazy = &megaohms
	return megaohms
}

// Gigaohms returns the ElectricImpedance value in Gigaohms.
//
// 
func (a *ElectricImpedance) Gigaohms() float64 {
	if a.gigaohmsLazy != nil {
		return *a.gigaohmsLazy
	}
	gigaohms := a.convertFromBase(ElectricImpedanceGigaohm)
	a.gigaohmsLazy = &gigaohms
	return gigaohms
}

// Teraohms returns the ElectricImpedance value in Teraohms.
//
// 
func (a *ElectricImpedance) Teraohms() float64 {
	if a.teraohmsLazy != nil {
		return *a.teraohmsLazy
	}
	teraohms := a.convertFromBase(ElectricImpedanceTeraohm)
	a.teraohmsLazy = &teraohms
	return teraohms
}


// ToDto creates a ElectricImpedanceDto representation from the ElectricImpedance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricImpedance) ToDto(holdInUnit *ElectricImpedanceUnits) ElectricImpedanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricImpedanceOhm // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricImpedanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricImpedance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricImpedance) ToDtoJSON(holdInUnit *ElectricImpedanceUnits) ([]byte, error) {
	// Convert to ElectricImpedanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricImpedance to a specific unit value.
// The function uses the provided unit type (ElectricImpedanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricImpedance) Convert(toUnit ElectricImpedanceUnits) float64 {
	switch toUnit { 
    case ElectricImpedanceOhm:
		return a.Ohms()
    case ElectricImpedanceNanoohm:
		return a.Nanoohms()
    case ElectricImpedanceMicroohm:
		return a.Microohms()
    case ElectricImpedanceMilliohm:
		return a.Milliohms()
    case ElectricImpedanceKiloohm:
		return a.Kiloohms()
    case ElectricImpedanceMegaohm:
		return a.Megaohms()
    case ElectricImpedanceGigaohm:
		return a.Gigaohms()
    case ElectricImpedanceTeraohm:
		return a.Teraohms()
	default:
		return math.NaN()
	}
}

func (a *ElectricImpedance) convertFromBase(toUnit ElectricImpedanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricImpedanceOhm:
		return (value) 
	case ElectricImpedanceNanoohm:
		return ((value) / 1e-09) 
	case ElectricImpedanceMicroohm:
		return ((value) / 1e-06) 
	case ElectricImpedanceMilliohm:
		return ((value) / 0.001) 
	case ElectricImpedanceKiloohm:
		return ((value) / 1000.0) 
	case ElectricImpedanceMegaohm:
		return ((value) / 1000000.0) 
	case ElectricImpedanceGigaohm:
		return ((value) / 1000000000.0) 
	case ElectricImpedanceTeraohm:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricImpedance) convertToBase(value float64, fromUnit ElectricImpedanceUnits) float64 {
	switch fromUnit { 
	case ElectricImpedanceOhm:
		return (value) 
	case ElectricImpedanceNanoohm:
		return ((value) * 1e-09) 
	case ElectricImpedanceMicroohm:
		return ((value) * 1e-06) 
	case ElectricImpedanceMilliohm:
		return ((value) * 0.001) 
	case ElectricImpedanceKiloohm:
		return ((value) * 1000.0) 
	case ElectricImpedanceMegaohm:
		return ((value) * 1000000.0) 
	case ElectricImpedanceGigaohm:
		return ((value) * 1000000000.0) 
	case ElectricImpedanceTeraohm:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricImpedance in the default unit (Ohm),
// formatted to two decimal places.
func (a ElectricImpedance) String() string {
	return a.ToString(ElectricImpedanceOhm, 2)
}

// ToString formats the ElectricImpedance value as a string with the specified unit and fractional digits.
// It converts the ElectricImpedance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricImpedance value will be converted (e.g., Ohm))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricImpedance with the unit abbreviation.
func (a *ElectricImpedance) ToString(unit ElectricImpedanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricImpedanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricImpedanceAbbreviation(unit))
}

// Equals checks if the given ElectricImpedance is equal to the current ElectricImpedance.
//
// Parameters:
//    other: The ElectricImpedance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricImpedance are equal, false otherwise.
func (a *ElectricImpedance) Equals(other *ElectricImpedance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricImpedance with another ElectricImpedance.
// It returns -1 if the current ElectricImpedance is less than the other ElectricImpedance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricImpedance to compare against.
//
// Returns:
//    int: -1 if the current ElectricImpedance is less, 1 if greater, and 0 if equal.
func (a *ElectricImpedance) CompareTo(other *ElectricImpedance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricImpedance to the current ElectricImpedance and returns the result.
// The result is a new ElectricImpedance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricImpedance to add to the current ElectricImpedance.
//
// Returns:
//    *ElectricImpedance: A new ElectricImpedance instance representing the sum of both ElectricImpedance.
func (a *ElectricImpedance) Add(other *ElectricImpedance) *ElectricImpedance {
	return &ElectricImpedance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricImpedance from the current ElectricImpedance and returns the result.
// The result is a new ElectricImpedance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricImpedance to subtract from the current ElectricImpedance.
//
// Returns:
//    *ElectricImpedance: A new ElectricImpedance instance representing the difference of both ElectricImpedance.
func (a *ElectricImpedance) Subtract(other *ElectricImpedance) *ElectricImpedance {
	return &ElectricImpedance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricImpedance by the given ElectricImpedance and returns the result.
// The result is a new ElectricImpedance instance with the product of the values.
//
// Parameters:
//    other: The ElectricImpedance to multiply with the current ElectricImpedance.
//
// Returns:
//    *ElectricImpedance: A new ElectricImpedance instance representing the product of both ElectricImpedance.
func (a *ElectricImpedance) Multiply(other *ElectricImpedance) *ElectricImpedance {
	return &ElectricImpedance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricImpedance by the given ElectricImpedance and returns the result.
// The result is a new ElectricImpedance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricImpedance to divide the current ElectricImpedance by.
//
// Returns:
//    *ElectricImpedance: A new ElectricImpedance instance representing the quotient of both ElectricImpedance.
func (a *ElectricImpedance) Divide(other *ElectricImpedance) *ElectricImpedance {
	return &ElectricImpedance{value: a.value / other.BaseValue()}
}

// GetElectricImpedanceAbbreviation gets the unit abbreviation.
func GetElectricImpedanceAbbreviation(unit ElectricImpedanceUnits) string {
	switch unit { 
	case ElectricImpedanceOhm:
		return "Ω" 
	case ElectricImpedanceNanoohm:
		return "nΩ" 
	case ElectricImpedanceMicroohm:
		return "μΩ" 
	case ElectricImpedanceMilliohm:
		return "mΩ" 
	case ElectricImpedanceKiloohm:
		return "kΩ" 
	case ElectricImpedanceMegaohm:
		return "MΩ" 
	case ElectricImpedanceGigaohm:
		return "GΩ" 
	case ElectricImpedanceTeraohm:
		return "TΩ" 
	default:
		return ""
	}
}