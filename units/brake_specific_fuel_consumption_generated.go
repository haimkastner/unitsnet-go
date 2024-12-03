// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// BrakeSpecificFuelConsumptionUnits enumeration
type BrakeSpecificFuelConsumptionUnits string

const (
    
        // 
        BrakeSpecificFuelConsumptionGramPerKiloWattHour BrakeSpecificFuelConsumptionUnits = "GramPerKiloWattHour"
        // 
        BrakeSpecificFuelConsumptionKilogramPerJoule BrakeSpecificFuelConsumptionUnits = "KilogramPerJoule"
        // The pound per horse power hour uses mechanical horse power and the imperial pound
        BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour BrakeSpecificFuelConsumptionUnits = "PoundPerMechanicalHorsepowerHour"
)

// BrakeSpecificFuelConsumptionDto represents an BrakeSpecificFuelConsumption
type BrakeSpecificFuelConsumptionDto struct {
	Value float64
	Unit  BrakeSpecificFuelConsumptionUnits
}

// BrakeSpecificFuelConsumptionDtoFactory struct to group related functions
type BrakeSpecificFuelConsumptionDtoFactory struct{}

func (udf BrakeSpecificFuelConsumptionDtoFactory) FromJSON(data []byte) (*BrakeSpecificFuelConsumptionDto, error) {
	a := BrakeSpecificFuelConsumptionDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a BrakeSpecificFuelConsumptionDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// BrakeSpecificFuelConsumption struct
type BrakeSpecificFuelConsumption struct {
	value       float64
    
    grams_per_kilo_watt_hourLazy *float64 
    kilograms_per_jouleLazy *float64 
    pounds_per_mechanical_horsepower_hourLazy *float64 
}

// BrakeSpecificFuelConsumptionFactory struct to group related functions
type BrakeSpecificFuelConsumptionFactory struct{}

func (uf BrakeSpecificFuelConsumptionFactory) CreateBrakeSpecificFuelConsumption(value float64, unit BrakeSpecificFuelConsumptionUnits) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, unit)
}

func (uf BrakeSpecificFuelConsumptionFactory) FromDto(dto BrakeSpecificFuelConsumptionDto) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(dto.Value, dto.Unit)
}

func (uf BrakeSpecificFuelConsumptionFactory) FromDtoJSON(data []byte) (*BrakeSpecificFuelConsumption, error) {
	unitDto, err := BrakeSpecificFuelConsumptionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return BrakeSpecificFuelConsumptionFactory{}.FromDto(*unitDto)
}


// FromGramPerKiloWattHour creates a new BrakeSpecificFuelConsumption instance from GramPerKiloWattHour.
func (uf BrakeSpecificFuelConsumptionFactory) FromGramsPerKiloWattHour(value float64) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, BrakeSpecificFuelConsumptionGramPerKiloWattHour)
}

// FromKilogramPerJoule creates a new BrakeSpecificFuelConsumption instance from KilogramPerJoule.
func (uf BrakeSpecificFuelConsumptionFactory) FromKilogramsPerJoule(value float64) (*BrakeSpecificFuelConsumption, error) {
	return newBrakeSpecificFuelConsumption(value, BrakeSpecificFuelConsumptionKilogramPerJoule)
}

// FromPoundPerMechanicalHorsepowerHour creates a new BrakeSpecificFuelConsumption instance from PoundPerMechanicalHorsepowerHour.
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

// BaseValue returns the base value of BrakeSpecificFuelConsumption in KilogramPerJoule.
func (a *BrakeSpecificFuelConsumption) BaseValue() float64 {
	return a.value
}


// GramPerKiloWattHour returns the value in GramPerKiloWattHour.
func (a *BrakeSpecificFuelConsumption) GramsPerKiloWattHour() float64 {
	if a.grams_per_kilo_watt_hourLazy != nil {
		return *a.grams_per_kilo_watt_hourLazy
	}
	grams_per_kilo_watt_hour := a.convertFromBase(BrakeSpecificFuelConsumptionGramPerKiloWattHour)
	a.grams_per_kilo_watt_hourLazy = &grams_per_kilo_watt_hour
	return grams_per_kilo_watt_hour
}

// KilogramPerJoule returns the value in KilogramPerJoule.
func (a *BrakeSpecificFuelConsumption) KilogramsPerJoule() float64 {
	if a.kilograms_per_jouleLazy != nil {
		return *a.kilograms_per_jouleLazy
	}
	kilograms_per_joule := a.convertFromBase(BrakeSpecificFuelConsumptionKilogramPerJoule)
	a.kilograms_per_jouleLazy = &kilograms_per_joule
	return kilograms_per_joule
}

// PoundPerMechanicalHorsepowerHour returns the value in PoundPerMechanicalHorsepowerHour.
func (a *BrakeSpecificFuelConsumption) PoundsPerMechanicalHorsepowerHour() float64 {
	if a.pounds_per_mechanical_horsepower_hourLazy != nil {
		return *a.pounds_per_mechanical_horsepower_hourLazy
	}
	pounds_per_mechanical_horsepower_hour := a.convertFromBase(BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour)
	a.pounds_per_mechanical_horsepower_hourLazy = &pounds_per_mechanical_horsepower_hour
	return pounds_per_mechanical_horsepower_hour
}


// ToDto creates an BrakeSpecificFuelConsumptionDto representation.
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

// ToDtoJSON creates an BrakeSpecificFuelConsumptionDto representation.
func (a *BrakeSpecificFuelConsumption) ToDtoJSON(holdInUnit *BrakeSpecificFuelConsumptionUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts BrakeSpecificFuelConsumption to a specific unit value.
func (a *BrakeSpecificFuelConsumption) Convert(toUnit BrakeSpecificFuelConsumptionUnits) float64 {
	switch toUnit { 
    case BrakeSpecificFuelConsumptionGramPerKiloWattHour:
		return a.GramsPerKiloWattHour()
    case BrakeSpecificFuelConsumptionKilogramPerJoule:
		return a.KilogramsPerJoule()
    case BrakeSpecificFuelConsumptionPoundPerMechanicalHorsepowerHour:
		return a.PoundsPerMechanicalHorsepowerHour()
	default:
		return 0
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

// Implement the String() method for AngleDto
func (a BrakeSpecificFuelConsumption) String() string {
	return a.ToString(BrakeSpecificFuelConsumptionKilogramPerJoule, 2)
}

// ToString formats the BrakeSpecificFuelConsumption to string.
// fractionalDigits -1 for not mention
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

// Check if the given BrakeSpecificFuelConsumption are equals to the current BrakeSpecificFuelConsumption
func (a *BrakeSpecificFuelConsumption) Equals(other *BrakeSpecificFuelConsumption) bool {
	return a.value == other.BaseValue()
}

// Check if the given BrakeSpecificFuelConsumption are equals to the current BrakeSpecificFuelConsumption
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

// Add the given BrakeSpecificFuelConsumption to the current BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Add(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value + other.BaseValue()}
}

// Subtract the given BrakeSpecificFuelConsumption to the current BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Subtract(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value - other.BaseValue()}
}

// Multiply the given BrakeSpecificFuelConsumption to the current BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Multiply(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value * other.BaseValue()}
}

// Divide the given BrakeSpecificFuelConsumption to the current BrakeSpecificFuelConsumption.
func (a *BrakeSpecificFuelConsumption) Divide(other *BrakeSpecificFuelConsumption) *BrakeSpecificFuelConsumption {
	return &BrakeSpecificFuelConsumption{value: a.value / other.BaseValue()}
}