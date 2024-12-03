// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricConductanceUnits enumeration
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

// ElectricConductanceDto represents an ElectricConductance
type ElectricConductanceDto struct {
	Value float64
	Unit  ElectricConductanceUnits
}

// ElectricConductanceDtoFactory struct to group related functions
type ElectricConductanceDtoFactory struct{}

func (udf ElectricConductanceDtoFactory) FromJSON(data []byte) (*ElectricConductanceDto, error) {
	a := ElectricConductanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricConductanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricConductance struct
type ElectricConductance struct {
	value       float64
    
    siemensLazy *float64 
    nanosiemensLazy *float64 
    microsiemensLazy *float64 
    millisiemensLazy *float64 
    kilosiemensLazy *float64 
}

// ElectricConductanceFactory struct to group related functions
type ElectricConductanceFactory struct{}

func (uf ElectricConductanceFactory) CreateElectricConductance(value float64, unit ElectricConductanceUnits) (*ElectricConductance, error) {
	return newElectricConductance(value, unit)
}

func (uf ElectricConductanceFactory) FromDto(dto ElectricConductanceDto) (*ElectricConductance, error) {
	return newElectricConductance(dto.Value, dto.Unit)
}

func (uf ElectricConductanceFactory) FromDtoJSON(data []byte) (*ElectricConductance, error) {
	unitDto, err := ElectricConductanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricConductanceFactory{}.FromDto(*unitDto)
}


// FromSiemens creates a new ElectricConductance instance from Siemens.
func (uf ElectricConductanceFactory) FromSiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceSiemens)
}

// FromNanosiemens creates a new ElectricConductance instance from Nanosiemens.
func (uf ElectricConductanceFactory) FromNanosiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceNanosiemens)
}

// FromMicrosiemens creates a new ElectricConductance instance from Microsiemens.
func (uf ElectricConductanceFactory) FromMicrosiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceMicrosiemens)
}

// FromMillisiemens creates a new ElectricConductance instance from Millisiemens.
func (uf ElectricConductanceFactory) FromMillisiemens(value float64) (*ElectricConductance, error) {
	return newElectricConductance(value, ElectricConductanceMillisiemens)
}

// FromKilosiemens creates a new ElectricConductance instance from Kilosiemens.
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

// BaseValue returns the base value of ElectricConductance in Siemens.
func (a *ElectricConductance) BaseValue() float64 {
	return a.value
}


// Siemens returns the value in Siemens.
func (a *ElectricConductance) Siemens() float64 {
	if a.siemensLazy != nil {
		return *a.siemensLazy
	}
	siemens := a.convertFromBase(ElectricConductanceSiemens)
	a.siemensLazy = &siemens
	return siemens
}

// Nanosiemens returns the value in Nanosiemens.
func (a *ElectricConductance) Nanosiemens() float64 {
	if a.nanosiemensLazy != nil {
		return *a.nanosiemensLazy
	}
	nanosiemens := a.convertFromBase(ElectricConductanceNanosiemens)
	a.nanosiemensLazy = &nanosiemens
	return nanosiemens
}

// Microsiemens returns the value in Microsiemens.
func (a *ElectricConductance) Microsiemens() float64 {
	if a.microsiemensLazy != nil {
		return *a.microsiemensLazy
	}
	microsiemens := a.convertFromBase(ElectricConductanceMicrosiemens)
	a.microsiemensLazy = &microsiemens
	return microsiemens
}

// Millisiemens returns the value in Millisiemens.
func (a *ElectricConductance) Millisiemens() float64 {
	if a.millisiemensLazy != nil {
		return *a.millisiemensLazy
	}
	millisiemens := a.convertFromBase(ElectricConductanceMillisiemens)
	a.millisiemensLazy = &millisiemens
	return millisiemens
}

// Kilosiemens returns the value in Kilosiemens.
func (a *ElectricConductance) Kilosiemens() float64 {
	if a.kilosiemensLazy != nil {
		return *a.kilosiemensLazy
	}
	kilosiemens := a.convertFromBase(ElectricConductanceKilosiemens)
	a.kilosiemensLazy = &kilosiemens
	return kilosiemens
}


// ToDto creates an ElectricConductanceDto representation.
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

// ToDtoJSON creates an ElectricConductanceDto representation.
func (a *ElectricConductance) ToDtoJSON(holdInUnit *ElectricConductanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricConductance to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ElectricConductance) String() string {
	return a.ToString(ElectricConductanceSiemens, 2)
}

// ToString formats the ElectricConductance to string.
// fractionalDigits -1 for not mention
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

// Check if the given ElectricConductance are equals to the current ElectricConductance
func (a *ElectricConductance) Equals(other *ElectricConductance) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricConductance are equals to the current ElectricConductance
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

// Add the given ElectricConductance to the current ElectricConductance.
func (a *ElectricConductance) Add(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricConductance to the current ElectricConductance.
func (a *ElectricConductance) Subtract(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricConductance to the current ElectricConductance.
func (a *ElectricConductance) Multiply(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value * other.BaseValue()}
}

// Divide the given ElectricConductance to the current ElectricConductance.
func (a *ElectricConductance) Divide(other *ElectricConductance) *ElectricConductance {
	return &ElectricConductance{value: a.value / other.BaseValue()}
}