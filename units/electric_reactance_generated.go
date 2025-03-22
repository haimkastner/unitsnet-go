// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricReactanceUnits defines various units of ElectricReactance.
type ElectricReactanceUnits string

const (
    
        // 
        ElectricReactanceOhm ElectricReactanceUnits = "Ohm"
        // 
        ElectricReactanceNanoohm ElectricReactanceUnits = "Nanoohm"
        // 
        ElectricReactanceMicroohm ElectricReactanceUnits = "Microohm"
        // 
        ElectricReactanceMilliohm ElectricReactanceUnits = "Milliohm"
        // 
        ElectricReactanceKiloohm ElectricReactanceUnits = "Kiloohm"
        // 
        ElectricReactanceMegaohm ElectricReactanceUnits = "Megaohm"
        // 
        ElectricReactanceGigaohm ElectricReactanceUnits = "Gigaohm"
        // 
        ElectricReactanceTeraohm ElectricReactanceUnits = "Teraohm"
)

var internalElectricReactanceUnitsMap = map[ElectricReactanceUnits]bool{
	
	ElectricReactanceOhm: true,
	ElectricReactanceNanoohm: true,
	ElectricReactanceMicroohm: true,
	ElectricReactanceMilliohm: true,
	ElectricReactanceKiloohm: true,
	ElectricReactanceMegaohm: true,
	ElectricReactanceGigaohm: true,
	ElectricReactanceTeraohm: true,
}

// ElectricReactanceDto represents a ElectricReactance measurement with a numerical value and its corresponding unit.
type ElectricReactanceDto struct {
    // Value is the numerical representation of the ElectricReactance.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricReactance, as defined in the ElectricReactanceUnits enumeration.
	Unit  ElectricReactanceUnits `json:"unit" validate:"required,oneof=Ohm,Nanoohm,Microohm,Milliohm,Kiloohm,Megaohm,Gigaohm,Teraohm"`
}

// ElectricReactanceDtoFactory groups methods for creating and serializing ElectricReactanceDto objects.
type ElectricReactanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricReactanceDtoFactory) FromJSON(data []byte) (*ElectricReactanceDto, error) {
	a := ElectricReactanceDto{}

    // Parse JSON into ElectricReactanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricReactanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricReactanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricReactance represents a measurement in a of ElectricReactance.
//
// In electrical circuits, reactance is the opposition presented to alternating current by inductance and capacitance. Along with resistance, it is one of two elements of impedance.
type ElectricReactance struct {
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

// ElectricReactanceFactory groups methods for creating ElectricReactance instances.
type ElectricReactanceFactory struct{}

// CreateElectricReactance creates a new ElectricReactance instance from the given value and unit.
func (uf ElectricReactanceFactory) CreateElectricReactance(value float64, unit ElectricReactanceUnits) (*ElectricReactance, error) {
	return newElectricReactance(value, unit)
}

// FromDto converts a ElectricReactanceDto to a ElectricReactance instance.
func (uf ElectricReactanceFactory) FromDto(dto ElectricReactanceDto) (*ElectricReactance, error) {
	return newElectricReactance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricReactance instance.
func (uf ElectricReactanceFactory) FromDtoJSON(data []byte) (*ElectricReactance, error) {
	unitDto, err := ElectricReactanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricReactanceDto from JSON: %w", err)
	}
	return ElectricReactanceFactory{}.FromDto(*unitDto)
}


// FromOhms creates a new ElectricReactance instance from a value in Ohms.
func (uf ElectricReactanceFactory) FromOhms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceOhm)
}

// FromNanoohms creates a new ElectricReactance instance from a value in Nanoohms.
func (uf ElectricReactanceFactory) FromNanoohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceNanoohm)
}

// FromMicroohms creates a new ElectricReactance instance from a value in Microohms.
func (uf ElectricReactanceFactory) FromMicroohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceMicroohm)
}

// FromMilliohms creates a new ElectricReactance instance from a value in Milliohms.
func (uf ElectricReactanceFactory) FromMilliohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceMilliohm)
}

// FromKiloohms creates a new ElectricReactance instance from a value in Kiloohms.
func (uf ElectricReactanceFactory) FromKiloohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceKiloohm)
}

// FromMegaohms creates a new ElectricReactance instance from a value in Megaohms.
func (uf ElectricReactanceFactory) FromMegaohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceMegaohm)
}

// FromGigaohms creates a new ElectricReactance instance from a value in Gigaohms.
func (uf ElectricReactanceFactory) FromGigaohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceGigaohm)
}

// FromTeraohms creates a new ElectricReactance instance from a value in Teraohms.
func (uf ElectricReactanceFactory) FromTeraohms(value float64) (*ElectricReactance, error) {
	return newElectricReactance(value, ElectricReactanceTeraohm)
}


// newElectricReactance creates a new ElectricReactance.
func newElectricReactance(value float64, fromUnit ElectricReactanceUnits) (*ElectricReactance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricReactanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricReactanceUnits", fromUnit)
	}
	a := &ElectricReactance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricReactance in Ohm unit (the base/default quantity).
func (a *ElectricReactance) BaseValue() float64 {
	return a.value
}


// Ohms returns the ElectricReactance value in Ohms.
//
// 
func (a *ElectricReactance) Ohms() float64 {
	if a.ohmsLazy != nil {
		return *a.ohmsLazy
	}
	ohms := a.convertFromBase(ElectricReactanceOhm)
	a.ohmsLazy = &ohms
	return ohms
}

// Nanoohms returns the ElectricReactance value in Nanoohms.
//
// 
func (a *ElectricReactance) Nanoohms() float64 {
	if a.nanoohmsLazy != nil {
		return *a.nanoohmsLazy
	}
	nanoohms := a.convertFromBase(ElectricReactanceNanoohm)
	a.nanoohmsLazy = &nanoohms
	return nanoohms
}

// Microohms returns the ElectricReactance value in Microohms.
//
// 
func (a *ElectricReactance) Microohms() float64 {
	if a.microohmsLazy != nil {
		return *a.microohmsLazy
	}
	microohms := a.convertFromBase(ElectricReactanceMicroohm)
	a.microohmsLazy = &microohms
	return microohms
}

// Milliohms returns the ElectricReactance value in Milliohms.
//
// 
func (a *ElectricReactance) Milliohms() float64 {
	if a.milliohmsLazy != nil {
		return *a.milliohmsLazy
	}
	milliohms := a.convertFromBase(ElectricReactanceMilliohm)
	a.milliohmsLazy = &milliohms
	return milliohms
}

// Kiloohms returns the ElectricReactance value in Kiloohms.
//
// 
func (a *ElectricReactance) Kiloohms() float64 {
	if a.kiloohmsLazy != nil {
		return *a.kiloohmsLazy
	}
	kiloohms := a.convertFromBase(ElectricReactanceKiloohm)
	a.kiloohmsLazy = &kiloohms
	return kiloohms
}

// Megaohms returns the ElectricReactance value in Megaohms.
//
// 
func (a *ElectricReactance) Megaohms() float64 {
	if a.megaohmsLazy != nil {
		return *a.megaohmsLazy
	}
	megaohms := a.convertFromBase(ElectricReactanceMegaohm)
	a.megaohmsLazy = &megaohms
	return megaohms
}

// Gigaohms returns the ElectricReactance value in Gigaohms.
//
// 
func (a *ElectricReactance) Gigaohms() float64 {
	if a.gigaohmsLazy != nil {
		return *a.gigaohmsLazy
	}
	gigaohms := a.convertFromBase(ElectricReactanceGigaohm)
	a.gigaohmsLazy = &gigaohms
	return gigaohms
}

// Teraohms returns the ElectricReactance value in Teraohms.
//
// 
func (a *ElectricReactance) Teraohms() float64 {
	if a.teraohmsLazy != nil {
		return *a.teraohmsLazy
	}
	teraohms := a.convertFromBase(ElectricReactanceTeraohm)
	a.teraohmsLazy = &teraohms
	return teraohms
}


// ToDto creates a ElectricReactanceDto representation from the ElectricReactance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricReactance) ToDto(holdInUnit *ElectricReactanceUnits) ElectricReactanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricReactanceOhm // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricReactanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricReactance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Ohm by default.
func (a *ElectricReactance) ToDtoJSON(holdInUnit *ElectricReactanceUnits) ([]byte, error) {
	// Convert to ElectricReactanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricReactance to a specific unit value.
// The function uses the provided unit type (ElectricReactanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricReactance) Convert(toUnit ElectricReactanceUnits) float64 {
	switch toUnit { 
    case ElectricReactanceOhm:
		return a.Ohms()
    case ElectricReactanceNanoohm:
		return a.Nanoohms()
    case ElectricReactanceMicroohm:
		return a.Microohms()
    case ElectricReactanceMilliohm:
		return a.Milliohms()
    case ElectricReactanceKiloohm:
		return a.Kiloohms()
    case ElectricReactanceMegaohm:
		return a.Megaohms()
    case ElectricReactanceGigaohm:
		return a.Gigaohms()
    case ElectricReactanceTeraohm:
		return a.Teraohms()
	default:
		return math.NaN()
	}
}

func (a *ElectricReactance) convertFromBase(toUnit ElectricReactanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricReactanceOhm:
		return (value) 
	case ElectricReactanceNanoohm:
		return ((value) / 1e-09) 
	case ElectricReactanceMicroohm:
		return ((value) / 1e-06) 
	case ElectricReactanceMilliohm:
		return ((value) / 0.001) 
	case ElectricReactanceKiloohm:
		return ((value) / 1000.0) 
	case ElectricReactanceMegaohm:
		return ((value) / 1000000.0) 
	case ElectricReactanceGigaohm:
		return ((value) / 1000000000.0) 
	case ElectricReactanceTeraohm:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricReactance) convertToBase(value float64, fromUnit ElectricReactanceUnits) float64 {
	switch fromUnit { 
	case ElectricReactanceOhm:
		return (value) 
	case ElectricReactanceNanoohm:
		return ((value) * 1e-09) 
	case ElectricReactanceMicroohm:
		return ((value) * 1e-06) 
	case ElectricReactanceMilliohm:
		return ((value) * 0.001) 
	case ElectricReactanceKiloohm:
		return ((value) * 1000.0) 
	case ElectricReactanceMegaohm:
		return ((value) * 1000000.0) 
	case ElectricReactanceGigaohm:
		return ((value) * 1000000000.0) 
	case ElectricReactanceTeraohm:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricReactance in the default unit (Ohm),
// formatted to two decimal places.
func (a ElectricReactance) String() string {
	return a.ToString(ElectricReactanceOhm, 2)
}

// ToString formats the ElectricReactance value as a string with the specified unit and fractional digits.
// It converts the ElectricReactance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricReactance value will be converted (e.g., Ohm))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricReactance with the unit abbreviation.
func (a *ElectricReactance) ToString(unit ElectricReactanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricReactanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricReactanceAbbreviation(unit))
}

// Equals checks if the given ElectricReactance is equal to the current ElectricReactance.
//
// Parameters:
//    other: The ElectricReactance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricReactance are equal, false otherwise.
func (a *ElectricReactance) Equals(other *ElectricReactance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricReactance with another ElectricReactance.
// It returns -1 if the current ElectricReactance is less than the other ElectricReactance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricReactance to compare against.
//
// Returns:
//    int: -1 if the current ElectricReactance is less, 1 if greater, and 0 if equal.
func (a *ElectricReactance) CompareTo(other *ElectricReactance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricReactance to the current ElectricReactance and returns the result.
// The result is a new ElectricReactance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricReactance to add to the current ElectricReactance.
//
// Returns:
//    *ElectricReactance: A new ElectricReactance instance representing the sum of both ElectricReactance.
func (a *ElectricReactance) Add(other *ElectricReactance) *ElectricReactance {
	return &ElectricReactance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricReactance from the current ElectricReactance and returns the result.
// The result is a new ElectricReactance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricReactance to subtract from the current ElectricReactance.
//
// Returns:
//    *ElectricReactance: A new ElectricReactance instance representing the difference of both ElectricReactance.
func (a *ElectricReactance) Subtract(other *ElectricReactance) *ElectricReactance {
	return &ElectricReactance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricReactance by the given ElectricReactance and returns the result.
// The result is a new ElectricReactance instance with the product of the values.
//
// Parameters:
//    other: The ElectricReactance to multiply with the current ElectricReactance.
//
// Returns:
//    *ElectricReactance: A new ElectricReactance instance representing the product of both ElectricReactance.
func (a *ElectricReactance) Multiply(other *ElectricReactance) *ElectricReactance {
	return &ElectricReactance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricReactance by the given ElectricReactance and returns the result.
// The result is a new ElectricReactance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricReactance to divide the current ElectricReactance by.
//
// Returns:
//    *ElectricReactance: A new ElectricReactance instance representing the quotient of both ElectricReactance.
func (a *ElectricReactance) Divide(other *ElectricReactance) *ElectricReactance {
	return &ElectricReactance{value: a.value / other.BaseValue()}
}

// GetElectricReactanceAbbreviation gets the unit abbreviation.
func GetElectricReactanceAbbreviation(unit ElectricReactanceUnits) string {
	switch unit { 
	case ElectricReactanceOhm:
		return "Ω" 
	case ElectricReactanceNanoohm:
		return "nΩ" 
	case ElectricReactanceMicroohm:
		return "μΩ" 
	case ElectricReactanceMilliohm:
		return "mΩ" 
	case ElectricReactanceKiloohm:
		return "kΩ" 
	case ElectricReactanceMegaohm:
		return "MΩ" 
	case ElectricReactanceGigaohm:
		return "GΩ" 
	case ElectricReactanceTeraohm:
		return "TΩ" 
	default:
		return ""
	}
}