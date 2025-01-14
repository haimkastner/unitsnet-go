// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// BrakeSpecificFuelConsumptionUnits defines various units of BrakeSpecificFuelConsumption.
type BrakeSpecificFuelConsumptionUnits string

const (
    
        // 
        BrakeSpecificFuelConsumptionGramPerKiloWattHour BrakeSpecificFuelConsumptionUnits = "GramPerKiloWattHour"
        // 
        BrakeSpecificFuelConsumptionKilogramPerJoule BrakeSpecificFuelConsumptionUnits = "KilogramPerJoule"
        // The pound per horse power hour uses mechanical horse power and the imperial pound
        BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour BrakeSpecificFuelConsumptionUnits = "PoundPerMechanicalHorsepowerHour"
)

// BrakeSpecificFuelConsumptionDto represents a BrakeSpecificFuelConsumption measurement with a numerical value and its corresponding unit.
type BrakeSpecificFuelConsumptionDto struct {
    // Value is the numerical representation of the BrakeSpecificFuelConsumption.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the BrakeSpecificFuelConsumption, as defined in the BrakeSpecificFuelConsumptionUnits enumeration.
	Unit  BrakeSpecificFuelConsumptionUnits `json:"unit"`
}

// BrakeSpecificFuelConsumptionDtoFactory groups methods for creating and serializing BrakeSpecificFuelConsumptionDto objects.
type BrakeSpecificFuelConsumptionDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a BrakeSpecificFuelConsumptionDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf BrakeSpecificFuelConsumptionDtoFactory) FromJSON(data []byte) (*BrakeSpecificFuelConsumptionDto, error) {
	a := BrakeSpecificFuelConsumptionDto{}

    // Parse JSON into BrakeSpecificFuelConsumptionDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a BrakeSpecificFuelConsumptionDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a BrakeSpecificFuelConsumptionDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// BrakeSpecificFuelConsumption represents a measurement in a of BrakeSpecificFuelConsumption.
//
// Brake specific fuel consumption (BSFC) is a measure of the fuel efficiency of any prime mover that burns fuel and produces rotational, or shaft, power. It is typically used for comparing the efficiency of internal combustion engines with a shaft output.
type BrakeSpecificFuelConsumption struct {
	// value is the base measurement stored internally.
	value       float64
    
    grams_per_kilo_watt_hourLazy *float64 
    kilograms_per_jouleLazy *float64 
    pounds_per_mechanical_horsepower_hourLazy *float64 
}

// BrakeSpecificFuelConsumptionFactory groups methods for creating BrakeSpecificFuelConsumption instances.
type BrakeSpecificFuelConsumptionFactory struct{}

// CreateBrakeSpecificFuelConsumption creates a new BrakeSpecificFuelConsumption instance from the given value and unit.
func (uf BrakeSpecificFuelConsumptionFactory) CreateBrakeSpecificFuelConsumption(value float64, unit BrakeSpecificFuelConsumptionUnits) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, unit)
}

// FromDto converts a BrakeSpecificFuelConsumptionDto to a BrakeSpecificFuelConsumption instance.
func (uf BrakeSpecificFuelConsumptionFactory) FromDto(dto BrakeSpecificFuelConsumptionDto) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a BrakeSpecificFuelConsumption instance.
func (uf BrakeSpecificFuelConsumptionFactory) FromDtoJSON(data []byte) (*BrakeSpecificFuelConsumption, error) {
	unitDto, err := BrakeSpecificFuelConsumptionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse BrakeSpecificFuelConsumptionDto from JSON: %w", err)
	}
	return BrakeSpecificFuelConsumptionFactory{}.FromDto(*unitDto)
}


// FromGramsPerKiloWattHour creates a new BrakeSpecificFuelConsumption instance from a value in GramsPerKiloWattHour.
func (uf BrakeSpecificFuelConsumptionFactory) FromGramsPerKiloWattHour(value float64) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, BrakeSpecificFuelConsumptionGramPerKiloWattHour)
}

// FromKilogramsPerJoule creates a new BrakeSpecificFuelConsumption instance from a value in KilogramsPerJoule.
func (uf BrakeSpecificFuelConsumptionFactory) FromKilogramsPerJoule(value float64) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, BrakeSpecificFuelConsumptionKilogramPerJoule)
}

// FromPoundsPerMechanicalHorsepowerHour creates a new BrakeSpecificFuelConsumption instance from a value in PoundsPerMechanicalHorsepowerHour.
func (uf BrakeSpecificFuelConsumptionFactory) FromPoundsPerMechanicalHorsepowerHour(value float64) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
}


// newBrakeSpecificFuelConsumption creates a new BrakeSpecificFuelConsumption.
func newBrakeSpecificFuelConsumption(value float64, fromUnit BrakeSpecificFuelConsumptionUnits) (*BrakeSpecificFuelConsumption, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &BrakeSpecificFuelConsumption{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of BrakeSpecificFuelConsumption in KilogramPerJoule unit (the base/default quantity).
func (a *BrakeSpecificFuelConsumption) BaseValue() float64 {
	return a.value
}


// GramsPerKiloWattHour returns the BrakeSpecificFuelConsumption value in GramsPerKiloWattHour.
//
// 
func (a *BrakeSpecificFuelConsumption) GramsPerKiloWattHour() float64 {
	if a.grams_per_kilo_watt_hourLazy != nil {
		return *a.grams_per_kilo_watt_hourLazy
	}
	grams_per_kilo_watt_hour := a.convertFromBase(BrakeSpecificFuelConsumptionGramPerKiloWattHour)
	a.grams_per_kilo_watt_hourLazy = &grams_per_kilo_watt_hour
	return grams_per_kilo_watt_hour
}

// KilogramsPerJoule returns the BrakeSpecificFuelConsumption value in KilogramsPerJoule.
//
// 
func (a *BrakeSpecificFuelConsumption) KilogramsPerJoule() float64 {
	if a.kilograms_per_jouleLazy != nil {
		return *a.kilograms_per_jouleLazy
	}
	kilograms_per_joule := a.convertFromBase(BrakeSpecificFuelConsumptionKilogramPerJoule)
	a.kilograms_per_jouleLazy = &kilograms_per_joule
	return kilograms_per_joule
}

// PoundsPerMechanicalHorsepowerHour returns the BrakeSpecificFuelConsumption value in PoundsPerMechanicalHorsepowerHour.
//
// The pound per horse power hour uses mechanical horse power and the imperial pound
func (a *BrakeSpecificFuelConsumption) PoundsPerMechanicalHorsepowerHour() float64 {
	if a.pounds_per_mechanical_horsepower_hourLazy != nil {
		return *a.pounds_per_mechanical_horsepower_hourLazy
	}
	pounds_per_mechanical_horsepower_hour := a.convertFromBase(BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
	a.pounds_per_mechanical_horsepower_hourLazy = &pounds_per_mechanical_horsepower_hour
	return pounds_per_mechanical_horsepower_hour
}


// ToDto creates a BrakeSpecificFuelConsumptionDto representation from the BrakeSpecificFuelConsumption instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerJoule by default.
func (a *BrakeSpecificFuelConsumption) ToDto(holdInUnit *BrakeSpecificFuelConsumptionUnits) BrakeSpecificFuelConsumptionDto {
	if holdInUnit == nil {
		defaultUnit := BrakeSpecificFuelConsumptionKilogramPerJoule // Default value
		holdInUnit = &defaultUnit
	}

	return BrakeSpecificFuelConsumptionDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the BrakeSpecificFuelConsumption instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerJoule by default.
func (a *BrakeSpecificFuelConsumption) ToDtoJSON(holdInUnit *BrakeSpecificFuelConsumptionUnits) ([]byte, error) {
	// Convert to BrakeSpecificFuelConsumptionDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a BrakeSpecificFuelConsumption to a specific unit value.
// The function uses the provided unit type (BrakeSpecificFuelConsumptionUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *BrakeSpecificFuelConsumption) Convert(toUnit BrakeSpecificFuelConsumptionUnits) float64 {
	switch toUnit { 
    case BrakeSpecificFuelConsumptionGramPerKiloWattHour:
		return a.GramsPerKiloWattHour()
    case BrakeSpecificFuelConsumptionKilogramPerJoule:
		return a.KilogramsPerJoule()
    case BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour:
		return a.PoundsPerMechanicalHorsepowerHour()
	default:
		return math.NaN()
	}
}

func (a *BrakeSpecificFuelConsumption) convertFromBase(toUnit BrakeSpecificFuelConsumptionUnits) float64 {
    value := a.value
	switch toUnit { 
	case BrakeSpecificFuelConsumptionGramPerKiloWattHour:
		return (value * 3.6e9) 
	case BrakeSpecificFuelConsumptionKilogramPerJoule:
		return (value) 
	case BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour:
		return (value / 1.689659410672e-7) 
	default:
		return math.NaN()
	}
}

func (a *BrakeSpecificFuelConsumption) convertToBase(value float64, fromUnit BrakeSpecificFuelConsumptionUnits) float64 {
	switch fromUnit { 
	case BrakeSpecificFuelConsumptionGramPerKiloWattHour:
		return (value / 3.6e9) 
	case BrakeSpecificFuelConsumptionKilogramPerJoule:
		return (value) 
	case BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour:
		return (value * 1.689659410672e-7) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the BrakeSpecificFuelConsumption in the default unit (KilogramPerJoule),
// formatted to two decimal places.
func (a BrakeSpecificFuelConsumption) String() string {
	return a.ToString(BrakeSpecificFuelConsumptionKilogramPerJoule, 2)
}

// ToString formats the BrakeSpecificFuelConsumption value as a string with the specified unit and fractional digits.
// It converts the BrakeSpecificFuelConsumption to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the BrakeSpecificFuelConsumption value will be converted (e.g., KilogramPerJoule))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the BrakeSpecificFuelConsumption with the unit abbreviation.
func (a *BrakeSpecificFuelConsumption) ToString(unit BrakeSpecificFuelConsumptionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *BrakeSpecificFuelConsumption) getUnitAbbreviation(unit BrakeSpecificFuelConsumptionUnits) string {
	switch unit { 
	case BrakeSpecificFuelConsumptionGramPerKiloWattHour:
		return "g/kWh" 
	case BrakeSpecificFuelConsumptionKilogramPerJoule:
		return "kg/J" 
	case BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour:
		return "lb/hph" 
	default:
		return ""
	}
}

// Equals checks if the given BrakeSpecificFuelConsumption is equal to the current BrakeSpecificFuelConsumption.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to compare against.
//
// Returns:
//    bool: Returns true if both BrakeSpecificFuelConsumption are equal, false otherwise.
func (a *BrakeSpecificFuelConsumption) Equals(other *BrakeSpecificFuelConsumption) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current BrakeSpecificFuelConsumption with another BrakeSpecificFuelConsumption.
// It returns -1 if the current BrakeSpecificFuelConsumption is less than the other BrakeSpecificFuelConsumption, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to compare against.
//
// Returns:
//    int: -1 if the current BrakeSpecificFuelConsumption is less, 1 if greater, and 0 if equal.
func (a *BrakeSpecificFuelConsumption) CompareTo(other *BrakeSpecificFuelConsumption) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given BrakeSpecificFuelConsumption to the current BrakeSpecificFuelConsumption and returns the result.
// The result is a new BrakeSpecificFuelConsumption instance with the sum of the values.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to add to the current BrakeSpecificFuelConsumption.
//
// Returns:
//    *BrakeSpecificFuelConsumption: A new BrakeSpecificFuelConsumption instance representing the sum of both BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Add(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given BrakeSpecificFuelConsumption from the current BrakeSpecificFuelConsumption and returns the result.
// The result is a new BrakeSpecificFuelConsumption instance with the difference of the values.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to subtract from the current BrakeSpecificFuelConsumption.
//
// Returns:
//    *BrakeSpecificFuelConsumption: A new BrakeSpecificFuelConsumption instance representing the difference of both BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Subtract(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current BrakeSpecificFuelConsumption by the given BrakeSpecificFuelConsumption and returns the result.
// The result is a new BrakeSpecificFuelConsumption instance with the product of the values.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to multiply with the current BrakeSpecificFuelConsumption.
//
// Returns:
//    *BrakeSpecificFuelConsumption: A new BrakeSpecificFuelConsumption instance representing the product of both BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Multiply(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value * other.BaseValue()}
}

// Divide divides the current BrakeSpecificFuelConsumption by the given BrakeSpecificFuelConsumption and returns the result.
// The result is a new BrakeSpecificFuelConsumption instance with the quotient of the values.
//
// Parameters:
//    other: The BrakeSpecificFuelConsumption to divide the current BrakeSpecificFuelConsumption by.
//
// Returns:
//    *BrakeSpecificFuelConsumption: A new BrakeSpecificFuelConsumption instance representing the quotient of both BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Divide(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value / other.BaseValue()}
}