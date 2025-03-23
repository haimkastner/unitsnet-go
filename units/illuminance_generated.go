// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// IlluminanceUnits defines various units of Illuminance.
type IlluminanceUnits string

const (
    
        // 
        IlluminanceLux IlluminanceUnits = "Lux"
        // 
        IlluminanceMillilux IlluminanceUnits = "Millilux"
        // 
        IlluminanceKilolux IlluminanceUnits = "Kilolux"
        // 
        IlluminanceMegalux IlluminanceUnits = "Megalux"
)

var internalIlluminanceUnitsMap = map[IlluminanceUnits]bool{
	
	IlluminanceLux: true,
	IlluminanceMillilux: true,
	IlluminanceKilolux: true,
	IlluminanceMegalux: true,
}

// IlluminanceDto represents a Illuminance measurement with a numerical value and its corresponding unit.
type IlluminanceDto struct {
    // Value is the numerical representation of the Illuminance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Illuminance, as defined in the IlluminanceUnits enumeration.
	Unit  IlluminanceUnits `json:"unit" validate:"required,oneof=Lux Millilux Kilolux Megalux"`
}

// IlluminanceDtoFactory groups methods for creating and serializing IlluminanceDto objects.
type IlluminanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a IlluminanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf IlluminanceDtoFactory) FromJSON(data []byte) (*IlluminanceDto, error) {
	a := IlluminanceDto{}

    // Parse JSON into IlluminanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a IlluminanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a IlluminanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Illuminance represents a measurement in a of Illuminance.
//
// In photometry, illuminance is the total luminous flux incident on a surface, per unit area.
type Illuminance struct {
	// value is the base measurement stored internally.
	value       float64
    
    luxLazy *float64 
    milliluxLazy *float64 
    kiloluxLazy *float64 
    megaluxLazy *float64 
}

// IlluminanceFactory groups methods for creating Illuminance instances.
type IlluminanceFactory struct{}

// CreateIlluminance creates a new Illuminance instance from the given value and unit.
func (uf IlluminanceFactory) CreateIlluminance(value float64, unit IlluminanceUnits) (*Illuminance, error) {
	return newIlluminance(value, unit)
}

// FromDto converts a IlluminanceDto to a Illuminance instance.
func (uf IlluminanceFactory) FromDto(dto IlluminanceDto) (*Illuminance, error) {
	return newIlluminance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Illuminance instance.
func (uf IlluminanceFactory) FromDtoJSON(data []byte) (*Illuminance, error) {
	unitDto, err := IlluminanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IlluminanceDto from JSON: %w", err)
	}
	return IlluminanceFactory{}.FromDto(*unitDto)
}


// FromLux creates a new Illuminance instance from a value in Lux.
func (uf IlluminanceFactory) FromLux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceLux)
}

// FromMillilux creates a new Illuminance instance from a value in Millilux.
func (uf IlluminanceFactory) FromMillilux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceMillilux)
}

// FromKilolux creates a new Illuminance instance from a value in Kilolux.
func (uf IlluminanceFactory) FromKilolux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceKilolux)
}

// FromMegalux creates a new Illuminance instance from a value in Megalux.
func (uf IlluminanceFactory) FromMegalux(value float64) (*Illuminance, error) {
	return newIlluminance(value, IlluminanceMegalux)
}


// newIlluminance creates a new Illuminance.
func newIlluminance(value float64, fromUnit IlluminanceUnits) (*Illuminance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalIlluminanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in IlluminanceUnits", fromUnit)
	}
	a := &Illuminance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Illuminance in Lux unit (the base/default quantity).
func (a *Illuminance) BaseValue() float64 {
	return a.value
}


// Lux returns the Illuminance value in Lux.
//
// 
func (a *Illuminance) Lux() float64 {
	if a.luxLazy != nil {
		return *a.luxLazy
	}
	lux := a.convertFromBase(IlluminanceLux)
	a.luxLazy = &lux
	return lux
}

// Millilux returns the Illuminance value in Millilux.
//
// 
func (a *Illuminance) Millilux() float64 {
	if a.milliluxLazy != nil {
		return *a.milliluxLazy
	}
	millilux := a.convertFromBase(IlluminanceMillilux)
	a.milliluxLazy = &millilux
	return millilux
}

// Kilolux returns the Illuminance value in Kilolux.
//
// 
func (a *Illuminance) Kilolux() float64 {
	if a.kiloluxLazy != nil {
		return *a.kiloluxLazy
	}
	kilolux := a.convertFromBase(IlluminanceKilolux)
	a.kiloluxLazy = &kilolux
	return kilolux
}

// Megalux returns the Illuminance value in Megalux.
//
// 
func (a *Illuminance) Megalux() float64 {
	if a.megaluxLazy != nil {
		return *a.megaluxLazy
	}
	megalux := a.convertFromBase(IlluminanceMegalux)
	a.megaluxLazy = &megalux
	return megalux
}


// ToDto creates a IlluminanceDto representation from the Illuminance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Lux by default.
func (a *Illuminance) ToDto(holdInUnit *IlluminanceUnits) IlluminanceDto {
	if holdInUnit == nil {
		defaultUnit := IlluminanceLux // Default value
		holdInUnit = &defaultUnit
	}

	return IlluminanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Illuminance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Lux by default.
func (a *Illuminance) ToDtoJSON(holdInUnit *IlluminanceUnits) ([]byte, error) {
	// Convert to IlluminanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Illuminance to a specific unit value.
// The function uses the provided unit type (IlluminanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Illuminance) Convert(toUnit IlluminanceUnits) float64 {
	switch toUnit { 
    case IlluminanceLux:
		return a.Lux()
    case IlluminanceMillilux:
		return a.Millilux()
    case IlluminanceKilolux:
		return a.Kilolux()
    case IlluminanceMegalux:
		return a.Megalux()
	default:
		return math.NaN()
	}
}

func (a *Illuminance) convertFromBase(toUnit IlluminanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case IlluminanceLux:
		return (value) 
	case IlluminanceMillilux:
		return ((value) / 0.001) 
	case IlluminanceKilolux:
		return ((value) / 1000.0) 
	case IlluminanceMegalux:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Illuminance) convertToBase(value float64, fromUnit IlluminanceUnits) float64 {
	switch fromUnit { 
	case IlluminanceLux:
		return (value) 
	case IlluminanceMillilux:
		return ((value) * 0.001) 
	case IlluminanceKilolux:
		return ((value) * 1000.0) 
	case IlluminanceMegalux:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Illuminance in the default unit (Lux),
// formatted to two decimal places.
func (a Illuminance) String() string {
	return a.ToString(IlluminanceLux, 2)
}

// ToString formats the Illuminance value as a string with the specified unit and fractional digits.
// It converts the Illuminance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Illuminance value will be converted (e.g., Lux))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Illuminance with the unit abbreviation.
func (a *Illuminance) ToString(unit IlluminanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetIlluminanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetIlluminanceAbbreviation(unit))
}

// Equals checks if the given Illuminance is equal to the current Illuminance.
//
// Parameters:
//    other: The Illuminance to compare against.
//
// Returns:
//    bool: Returns true if both Illuminance are equal, false otherwise.
func (a *Illuminance) Equals(other *Illuminance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Illuminance with another Illuminance.
// It returns -1 if the current Illuminance is less than the other Illuminance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Illuminance to compare against.
//
// Returns:
//    int: -1 if the current Illuminance is less, 1 if greater, and 0 if equal.
func (a *Illuminance) CompareTo(other *Illuminance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Illuminance to the current Illuminance and returns the result.
// The result is a new Illuminance instance with the sum of the values.
//
// Parameters:
//    other: The Illuminance to add to the current Illuminance.
//
// Returns:
//    *Illuminance: A new Illuminance instance representing the sum of both Illuminance.
func (a *Illuminance) Add(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Illuminance from the current Illuminance and returns the result.
// The result is a new Illuminance instance with the difference of the values.
//
// Parameters:
//    other: The Illuminance to subtract from the current Illuminance.
//
// Returns:
//    *Illuminance: A new Illuminance instance representing the difference of both Illuminance.
func (a *Illuminance) Subtract(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Illuminance by the given Illuminance and returns the result.
// The result is a new Illuminance instance with the product of the values.
//
// Parameters:
//    other: The Illuminance to multiply with the current Illuminance.
//
// Returns:
//    *Illuminance: A new Illuminance instance representing the product of both Illuminance.
func (a *Illuminance) Multiply(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value * other.BaseValue()}
}

// Divide divides the current Illuminance by the given Illuminance and returns the result.
// The result is a new Illuminance instance with the quotient of the values.
//
// Parameters:
//    other: The Illuminance to divide the current Illuminance by.
//
// Returns:
//    *Illuminance: A new Illuminance instance representing the quotient of both Illuminance.
func (a *Illuminance) Divide(other *Illuminance) *Illuminance {
	return &Illuminance{value: a.value / other.BaseValue()}
}

// GetIlluminanceAbbreviation gets the unit abbreviation.
func GetIlluminanceAbbreviation(unit IlluminanceUnits) string {
	switch unit { 
	case IlluminanceLux:
		return "lx" 
	case IlluminanceMillilux:
		return "mlx" 
	case IlluminanceKilolux:
		return "klx" 
	case IlluminanceMegalux:
		return "Mlx" 
	default:
		return ""
	}
}