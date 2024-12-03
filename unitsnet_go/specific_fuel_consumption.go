package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// SpecificFuelConsumptionUnits enumeration
type SpecificFuelConsumptionUnits string

const (
    
        // 
        SpecificFuelConsumptionPoundMassPerPoundForceHour SpecificFuelConsumptionUnits = "PoundMassPerPoundForceHour"
        // 
        SpecificFuelConsumptionKilogramPerKilogramForceHour SpecificFuelConsumptionUnits = "KilogramPerKilogramForceHour"
        // 
        SpecificFuelConsumptionGramPerKiloNewtonSecond SpecificFuelConsumptionUnits = "GramPerKiloNewtonSecond"
        // 
        SpecificFuelConsumptionKilogramPerKiloNewtonSecond SpecificFuelConsumptionUnits = "KilogramPerKiloNewtonSecond"
)

// SpecificFuelConsumptionDto represents an SpecificFuelConsumption
type SpecificFuelConsumptionDto struct {
	Value float64
	Unit  SpecificFuelConsumptionUnits
}

// SpecificFuelConsumptionDtoFactory struct to group related functions
type SpecificFuelConsumptionDtoFactory struct{}

func (udf SpecificFuelConsumptionDtoFactory) FromJSON(data []byte) (*SpecificFuelConsumptionDto, error) {
	a := SpecificFuelConsumptionDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a SpecificFuelConsumptionDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// SpecificFuelConsumption struct
type SpecificFuelConsumption struct {
	value       float64
    
    pounds_mass_per_pound_force_hourLazy *float64 
    kilograms_per_kilogram_force_hourLazy *float64 
    grams_per_kilo_newton_secondLazy *float64 
    kilograms_per_kilo_newton_secondLazy *float64 
}

// SpecificFuelConsumptionFactory struct to group related functions
type SpecificFuelConsumptionFactory struct{}

func (uf SpecificFuelConsumptionFactory) CreateSpecificFuelConsumption(value float64, unit SpecificFuelConsumptionUnits) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, unit)
}

func (uf SpecificFuelConsumptionFactory) FromDto(dto SpecificFuelConsumptionDto) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(dto.Value, dto.Unit)
}

func (uf SpecificFuelConsumptionFactory) FromDtoJSON(data []byte) (*SpecificFuelConsumption, error) {
	unitDto, err := SpecificFuelConsumptionDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return SpecificFuelConsumptionFactory{}.FromDto(*unitDto)
}


// FromPoundMassPerPoundForceHour creates a new SpecificFuelConsumption instance from PoundMassPerPoundForceHour.
func (uf SpecificFuelConsumptionFactory) FromPoundsMassPerPoundForceHour(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionPoundMassPerPoundForceHour)
}

// FromKilogramPerKilogramForceHour creates a new SpecificFuelConsumption instance from KilogramPerKilogramForceHour.
func (uf SpecificFuelConsumptionFactory) FromKilogramsPerKilogramForceHour(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionKilogramPerKilogramForceHour)
}

// FromGramPerKiloNewtonSecond creates a new SpecificFuelConsumption instance from GramPerKiloNewtonSecond.
func (uf SpecificFuelConsumptionFactory) FromGramsPerKiloNewtonSecond(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionGramPerKiloNewtonSecond)
}

// FromKilogramPerKiloNewtonSecond creates a new SpecificFuelConsumption instance from KilogramPerKiloNewtonSecond.
func (uf SpecificFuelConsumptionFactory) FromKilogramsPerKiloNewtonSecond(value float64) (*SpecificFuelConsumption, error) {
	return newSpecificFuelConsumption(value, SpecificFuelConsumptionKilogramPerKiloNewtonSecond)
}




// newSpecificFuelConsumption creates a new SpecificFuelConsumption.
func newSpecificFuelConsumption(value float64, fromUnit SpecificFuelConsumptionUnits) (*SpecificFuelConsumption, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &SpecificFuelConsumption{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of SpecificFuelConsumption in GramPerKiloNewtonSecond.
func (a *SpecificFuelConsumption) BaseValue() float64 {
	return a.value
}


// PoundMassPerPoundForceHour returns the value in PoundMassPerPoundForceHour.
func (a *SpecificFuelConsumption) PoundsMassPerPoundForceHour() float64 {
	if a.pounds_mass_per_pound_force_hourLazy != nil {
		return *a.pounds_mass_per_pound_force_hourLazy
	}
	pounds_mass_per_pound_force_hour := a.convertFromBase(SpecificFuelConsumptionPoundMassPerPoundForceHour)
	a.pounds_mass_per_pound_force_hourLazy = &pounds_mass_per_pound_force_hour
	return pounds_mass_per_pound_force_hour
}

// KilogramPerKilogramForceHour returns the value in KilogramPerKilogramForceHour.
func (a *SpecificFuelConsumption) KilogramsPerKilogramForceHour() float64 {
	if a.kilograms_per_kilogram_force_hourLazy != nil {
		return *a.kilograms_per_kilogram_force_hourLazy
	}
	kilograms_per_kilogram_force_hour := a.convertFromBase(SpecificFuelConsumptionKilogramPerKilogramForceHour)
	a.kilograms_per_kilogram_force_hourLazy = &kilograms_per_kilogram_force_hour
	return kilograms_per_kilogram_force_hour
}

// GramPerKiloNewtonSecond returns the value in GramPerKiloNewtonSecond.
func (a *SpecificFuelConsumption) GramsPerKiloNewtonSecond() float64 {
	if a.grams_per_kilo_newton_secondLazy != nil {
		return *a.grams_per_kilo_newton_secondLazy
	}
	grams_per_kilo_newton_second := a.convertFromBase(SpecificFuelConsumptionGramPerKiloNewtonSecond)
	a.grams_per_kilo_newton_secondLazy = &grams_per_kilo_newton_second
	return grams_per_kilo_newton_second
}

// KilogramPerKiloNewtonSecond returns the value in KilogramPerKiloNewtonSecond.
func (a *SpecificFuelConsumption) KilogramsPerKiloNewtonSecond() float64 {
	if a.kilograms_per_kilo_newton_secondLazy != nil {
		return *a.kilograms_per_kilo_newton_secondLazy
	}
	kilograms_per_kilo_newton_second := a.convertFromBase(SpecificFuelConsumptionKilogramPerKiloNewtonSecond)
	a.kilograms_per_kilo_newton_secondLazy = &kilograms_per_kilo_newton_second
	return kilograms_per_kilo_newton_second
}


// ToDto creates an SpecificFuelConsumptionDto representation.
func (a *SpecificFuelConsumption) ToDto(holdInUnit *SpecificFuelConsumptionUnits) SpecificFuelConsumptionDto {
	if holdInUnit == nil {
		defaultUnit := SpecificFuelConsumptionGramPerKiloNewtonSecond // Default value
		holdInUnit = &defaultUnit
	}

	return SpecificFuelConsumptionDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an SpecificFuelConsumptionDto representation.
func (a *SpecificFuelConsumption) ToDtoJSON(holdInUnit *SpecificFuelConsumptionUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts SpecificFuelConsumption to a specific unit value.
func (a *SpecificFuelConsumption) Convert(toUnit SpecificFuelConsumptionUnits) float64 {
	switch toUnit { 
    case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return a.PoundsMassPerPoundForceHour()
    case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return a.KilogramsPerKilogramForceHour()
    case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return a.GramsPerKiloNewtonSecond()
    case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return a.KilogramsPerKiloNewtonSecond()
	default:
		return 0
	}
}

func (a *SpecificFuelConsumption) convertFromBase(toUnit SpecificFuelConsumptionUnits) float64 {
    value := a.value
	switch toUnit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return (value / 28.33) 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return (value / 28.33) 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return (value) 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *SpecificFuelConsumption) convertToBase(value float64, fromUnit SpecificFuelConsumptionUnits) float64 {
	switch fromUnit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return (value * 28.33) 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return (value * 28.33) 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return (value) 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a SpecificFuelConsumption) String() string {
	return a.ToString(SpecificFuelConsumptionGramPerKiloNewtonSecond, 2)
}

// ToString formats the SpecificFuelConsumption to string.
// fractionalDigits -1 for not mention
func (a *SpecificFuelConsumption) ToString(unit SpecificFuelConsumptionUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *SpecificFuelConsumption) getUnitAbbreviation(unit SpecificFuelConsumptionUnits) string {
	switch unit { 
	case SpecificFuelConsumptionPoundMassPerPoundForceHour:
		return "lb/(lbf·h)" 
	case SpecificFuelConsumptionKilogramPerKilogramForceHour:
		return "kg/(kgf�h)" 
	case SpecificFuelConsumptionGramPerKiloNewtonSecond:
		return "g/(kN�s)" 
	case SpecificFuelConsumptionKilogramPerKiloNewtonSecond:
		return "kg/(kN�s)" 
	default:
		return ""
	}
}

// Check if the given SpecificFuelConsumption are equals to the current SpecificFuelConsumption
func (a *SpecificFuelConsumption) Equals(other *SpecificFuelConsumption) bool {
	return a.value == other.BaseValue()
}

// Check if the given SpecificFuelConsumption are equals to the current SpecificFuelConsumption
func (a *SpecificFuelConsumption) CompareTo(other *SpecificFuelConsumption) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given SpecificFuelConsumption to the current SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Add(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value + other.BaseValue()}
}

// Subtract the given SpecificFuelConsumption to the current SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Subtract(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value - other.BaseValue()}
}

// Multiply the given SpecificFuelConsumption to the current SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Multiply(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value * other.BaseValue()}
}

// Divide the given SpecificFuelConsumption to the current SpecificFuelConsumption.
func (a *SpecificFuelConsumption) Divide(other *SpecificFuelConsumption) *SpecificFuelConsumption {
	return &SpecificFuelConsumption{value: a.value / other.BaseValue()}
}