// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AmplitudeRatioUnits defines various units of AmplitudeRatio.
type AmplitudeRatioUnits string

const (
    
        // 
        AmplitudeRatioDecibelVolt AmplitudeRatioUnits = "DecibelVolt"
        // 
        AmplitudeRatioDecibelMicrovolt AmplitudeRatioUnits = "DecibelMicrovolt"
        // 
        AmplitudeRatioDecibelMillivolt AmplitudeRatioUnits = "DecibelMillivolt"
        // 
        AmplitudeRatioDecibelUnloaded AmplitudeRatioUnits = "DecibelUnloaded"
)

// AmplitudeRatioDto represents a AmplitudeRatio measurement with a numerical value and its corresponding unit.
type AmplitudeRatioDto struct {
    // Value is the numerical representation of the AmplitudeRatio.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the AmplitudeRatio, as defined in the AmplitudeRatioUnits enumeration.
	Unit  AmplitudeRatioUnits `json:"unit"`
}

// AmplitudeRatioDtoFactory groups methods for creating and serializing AmplitudeRatioDto objects.
type AmplitudeRatioDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AmplitudeRatioDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AmplitudeRatioDtoFactory) FromJSON(data []byte) (*AmplitudeRatioDto, error) {
	a := AmplitudeRatioDto{}

    // Parse JSON into AmplitudeRatioDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AmplitudeRatioDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AmplitudeRatioDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// AmplitudeRatio represents a measurement in a of AmplitudeRatio.
//
// The strength of a signal expressed in decibels (dB) relative to one volt RMS.
type AmplitudeRatio struct {
	// value is the base measurement stored internally.
	value       float64
    
    decibel_voltsLazy *float64 
    decibel_microvoltsLazy *float64 
    decibel_millivoltsLazy *float64 
    decibels_unloadedLazy *float64 
}

// AmplitudeRatioFactory groups methods for creating AmplitudeRatio instances.
type AmplitudeRatioFactory struct{}

// CreateAmplitudeRatio creates a new AmplitudeRatio instance from the given value and unit.
func (uf AmplitudeRatioFactory) CreateAmplitudeRatio(value float64, unit AmplitudeRatioUnits) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, unit)
}

// FromDto converts a AmplitudeRatioDto to a AmplitudeRatio instance.
func (uf AmplitudeRatioFactory) FromDto(dto AmplitudeRatioDto) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AmplitudeRatio instance.
func (uf AmplitudeRatioFactory) FromDtoJSON(data []byte) (*AmplitudeRatio, error) {
	unitDto, err := AmplitudeRatioDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AmplitudeRatioDto from JSON: %w", err)
	}
	return AmplitudeRatioFactory{}.FromDto(*unitDto)
}


// FromDecibelVolts creates a new AmplitudeRatio instance from a value in DecibelVolts.
func (uf AmplitudeRatioFactory) FromDecibelVolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelVolt)
}

// FromDecibelMicrovolts creates a new AmplitudeRatio instance from a value in DecibelMicrovolts.
func (uf AmplitudeRatioFactory) FromDecibelMicrovolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelMicrovolt)
}

// FromDecibelMillivolts creates a new AmplitudeRatio instance from a value in DecibelMillivolts.
func (uf AmplitudeRatioFactory) FromDecibelMillivolts(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelMillivolt)
}

// FromDecibelsUnloaded creates a new AmplitudeRatio instance from a value in DecibelsUnloaded.
func (uf AmplitudeRatioFactory) FromDecibelsUnloaded(value float64) (*AmplitudeRatio, error) {
	return newAmplitudeRatio(value, AmplitudeRatioDecibelUnloaded)
}


// newAmplitudeRatio creates a new AmplitudeRatio.
func newAmplitudeRatio(value float64, fromUnit AmplitudeRatioUnits) (*AmplitudeRatio, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AmplitudeRatio{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AmplitudeRatio in DecibelVolt unit (the base/default quantity).
func (a *AmplitudeRatio) BaseValue() float64 {
	return a.value
}


// DecibelVolts returns the AmplitudeRatio value in DecibelVolts.
//
// 
func (a *AmplitudeRatio) DecibelVolts() float64 {
	if a.decibel_voltsLazy != nil {
		return *a.decibel_voltsLazy
	}
	decibel_volts := a.convertFromBase(AmplitudeRatioDecibelVolt)
	a.decibel_voltsLazy = &decibel_volts
	return decibel_volts
}

// DecibelMicrovolts returns the AmplitudeRatio value in DecibelMicrovolts.
//
// 
func (a *AmplitudeRatio) DecibelMicrovolts() float64 {
	if a.decibel_microvoltsLazy != nil {
		return *a.decibel_microvoltsLazy
	}
	decibel_microvolts := a.convertFromBase(AmplitudeRatioDecibelMicrovolt)
	a.decibel_microvoltsLazy = &decibel_microvolts
	return decibel_microvolts
}

// DecibelMillivolts returns the AmplitudeRatio value in DecibelMillivolts.
//
// 
func (a *AmplitudeRatio) DecibelMillivolts() float64 {
	if a.decibel_millivoltsLazy != nil {
		return *a.decibel_millivoltsLazy
	}
	decibel_millivolts := a.convertFromBase(AmplitudeRatioDecibelMillivolt)
	a.decibel_millivoltsLazy = &decibel_millivolts
	return decibel_millivolts
}

// DecibelsUnloaded returns the AmplitudeRatio value in DecibelsUnloaded.
//
// 
func (a *AmplitudeRatio) DecibelsUnloaded() float64 {
	if a.decibels_unloadedLazy != nil {
		return *a.decibels_unloadedLazy
	}
	decibels_unloaded := a.convertFromBase(AmplitudeRatioDecibelUnloaded)
	a.decibels_unloadedLazy = &decibels_unloaded
	return decibels_unloaded
}


// ToDto creates a AmplitudeRatioDto representation from the AmplitudeRatio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecibelVolt by default.
func (a *AmplitudeRatio) ToDto(holdInUnit *AmplitudeRatioUnits) AmplitudeRatioDto {
	if holdInUnit == nil {
		defaultUnit := AmplitudeRatioDecibelVolt // Default value
		holdInUnit = &defaultUnit
	}

	return AmplitudeRatioDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AmplitudeRatio instance.
//
// If the provided holdInUnit is nil, the value will be repesented by DecibelVolt by default.
func (a *AmplitudeRatio) ToDtoJSON(holdInUnit *AmplitudeRatioUnits) ([]byte, error) {
	// Convert to AmplitudeRatioDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AmplitudeRatio to a specific unit value.
// The function uses the provided unit type (AmplitudeRatioUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AmplitudeRatio) Convert(toUnit AmplitudeRatioUnits) float64 {
	switch toUnit { 
    case AmplitudeRatioDecibelVolt:
		return a.DecibelVolts()
    case AmplitudeRatioDecibelMicrovolt:
		return a.DecibelMicrovolts()
    case AmplitudeRatioDecibelMillivolt:
		return a.DecibelMillivolts()
    case AmplitudeRatioDecibelUnloaded:
		return a.DecibelsUnloaded()
	default:
		return math.NaN()
	}
}

func (a *AmplitudeRatio) convertFromBase(toUnit AmplitudeRatioUnits) float64 {
    value := a.value
	switch toUnit { 
	case AmplitudeRatioDecibelVolt:
		return (value) 
	case AmplitudeRatioDecibelMicrovolt:
		return (value + 120) 
	case AmplitudeRatioDecibelMillivolt:
		return (value + 60) 
	case AmplitudeRatioDecibelUnloaded:
		return (value + 2.218487499) 
	default:
		return math.NaN()
	}
}

func (a *AmplitudeRatio) convertToBase(value float64, fromUnit AmplitudeRatioUnits) float64 {
	switch fromUnit { 
	case AmplitudeRatioDecibelVolt:
		return (value) 
	case AmplitudeRatioDecibelMicrovolt:
		return (value - 120) 
	case AmplitudeRatioDecibelMillivolt:
		return (value - 60) 
	case AmplitudeRatioDecibelUnloaded:
		return (value - 2.218487499) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AmplitudeRatio in the default unit (DecibelVolt),
// formatted to two decimal places.
func (a AmplitudeRatio) String() string {
	return a.ToString(AmplitudeRatioDecibelVolt, 2)
}

// ToString formats the AmplitudeRatio value as a string with the specified unit and fractional digits.
// It converts the AmplitudeRatio to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AmplitudeRatio value will be converted (e.g., DecibelVolt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AmplitudeRatio with the unit abbreviation.
func (a *AmplitudeRatio) ToString(unit AmplitudeRatioUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAmplitudeRatioAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAmplitudeRatioAbbreviation(unit))
}

// Equals checks if the given AmplitudeRatio is equal to the current AmplitudeRatio.
//
// Parameters:
//    other: The AmplitudeRatio to compare against.
//
// Returns:
//    bool: Returns true if both AmplitudeRatio are equal, false otherwise.
func (a *AmplitudeRatio) Equals(other *AmplitudeRatio) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AmplitudeRatio with another AmplitudeRatio.
// It returns -1 if the current AmplitudeRatio is less than the other AmplitudeRatio, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AmplitudeRatio to compare against.
//
// Returns:
//    int: -1 if the current AmplitudeRatio is less, 1 if greater, and 0 if equal.
func (a *AmplitudeRatio) CompareTo(other *AmplitudeRatio) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AmplitudeRatio to the current AmplitudeRatio and returns the result.
// The result is a new AmplitudeRatio instance with the sum of the values.
//
// Parameters:
//    other: The AmplitudeRatio to add to the current AmplitudeRatio.
//
// Returns:
//    *AmplitudeRatio: A new AmplitudeRatio instance representing the sum of both AmplitudeRatio.
func (a *AmplitudeRatio) Add(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AmplitudeRatio from the current AmplitudeRatio and returns the result.
// The result is a new AmplitudeRatio instance with the difference of the values.
//
// Parameters:
//    other: The AmplitudeRatio to subtract from the current AmplitudeRatio.
//
// Returns:
//    *AmplitudeRatio: A new AmplitudeRatio instance representing the difference of both AmplitudeRatio.
func (a *AmplitudeRatio) Subtract(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AmplitudeRatio by the given AmplitudeRatio and returns the result.
// The result is a new AmplitudeRatio instance with the product of the values.
//
// Parameters:
//    other: The AmplitudeRatio to multiply with the current AmplitudeRatio.
//
// Returns:
//    *AmplitudeRatio: A new AmplitudeRatio instance representing the product of both AmplitudeRatio.
func (a *AmplitudeRatio) Multiply(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value * other.BaseValue()}
}

// Divide divides the current AmplitudeRatio by the given AmplitudeRatio and returns the result.
// The result is a new AmplitudeRatio instance with the quotient of the values.
//
// Parameters:
//    other: The AmplitudeRatio to divide the current AmplitudeRatio by.
//
// Returns:
//    *AmplitudeRatio: A new AmplitudeRatio instance representing the quotient of both AmplitudeRatio.
func (a *AmplitudeRatio) Divide(other *AmplitudeRatio) *AmplitudeRatio {
	return &AmplitudeRatio{value: a.value / other.BaseValue()}
}

// GetAmplitudeRatioAbbreviation gets the unit abbreviation.
func GetAmplitudeRatioAbbreviation(unit AmplitudeRatioUnits) string {
	switch unit { 
	case AmplitudeRatioDecibelVolt:
		return "dBV" 
	case AmplitudeRatioDecibelMicrovolt:
		return "dBµV" 
	case AmplitudeRatioDecibelMillivolt:
		return "dBmV" 
	case AmplitudeRatioDecibelUnloaded:
		return "dBu" 
	default:
		return ""
	}
}