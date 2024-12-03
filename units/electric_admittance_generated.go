// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricAdmittanceUnits enumeration
type ElectricAdmittanceUnits string

const (
    
        // 
        ElectricAdmittanceSiemens ElectricAdmittanceUnits = "Siemens"
        // 
        ElectricAdmittanceNanosiemens ElectricAdmittanceUnits = "Nanosiemens"
        // 
        ElectricAdmittanceMicrosiemens ElectricAdmittanceUnits = "Microsiemens"
        // 
        ElectricAdmittanceMillisiemens ElectricAdmittanceUnits = "Millisiemens"
)

// ElectricAdmittanceDto represents an ElectricAdmittance
type ElectricAdmittanceDto struct {
	Value float64
	Unit  ElectricAdmittanceUnits
}

// ElectricAdmittanceDtoFactory struct to group related functions
type ElectricAdmittanceDtoFactory struct{}

func (udf ElectricAdmittanceDtoFactory) FromJSON(data []byte) (*ElectricAdmittanceDto, error) {
	a := ElectricAdmittanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricAdmittanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricAdmittance struct
type ElectricAdmittance struct {
	value       float64
    
    siemensLazy *float64 
    nanosiemensLazy *float64 
    microsiemensLazy *float64 
    millisiemensLazy *float64 
}

// ElectricAdmittanceFactory struct to group related functions
type ElectricAdmittanceFactory struct{}

func (uf ElectricAdmittanceFactory) CreateElectricAdmittance(value float64, unit ElectricAdmittanceUnits) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, unit)
}

func (uf ElectricAdmittanceFactory) FromDto(dto ElectricAdmittanceDto) (*ElectricAdmittance, error) {
	return newElectricAdmittance(dto.Value, dto.Unit)
}

func (uf ElectricAdmittanceFactory) FromDtoJSON(data []byte) (*ElectricAdmittance, error) {
	unitDto, err := ElectricAdmittanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricAdmittanceFactory{}.FromDto(*unitDto)
}


// FromSiemens creates a new ElectricAdmittance instance from Siemens.
func (uf ElectricAdmittanceFactory) FromSiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceSiemens)
}

// FromNanosiemens creates a new ElectricAdmittance instance from Nanosiemens.
func (uf ElectricAdmittanceFactory) FromNanosiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceNanosiemens)
}

// FromMicrosiemens creates a new ElectricAdmittance instance from Microsiemens.
func (uf ElectricAdmittanceFactory) FromMicrosiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMicrosiemens)
}

// FromMillisiemens creates a new ElectricAdmittance instance from Millisiemens.
func (uf ElectricAdmittanceFactory) FromMillisiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMillisiemens)
}




// newElectricAdmittance creates a new ElectricAdmittance.
func newElectricAdmittance(value float64, fromUnit ElectricAdmittanceUnits) (*ElectricAdmittance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricAdmittance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricAdmittance in Siemens.
func (a *ElectricAdmittance) BaseValue() float64 {
	return a.value
}


// Siemens returns the value in Siemens.
func (a *ElectricAdmittance) Siemens() float64 {
	if a.siemensLazy != nil {
		return *a.siemensLazy
	}
	siemens := a.convertFromBase(ElectricAdmittanceSiemens)
	a.siemensLazy = &siemens
	return siemens
}

// Nanosiemens returns the value in Nanosiemens.
func (a *ElectricAdmittance) Nanosiemens() float64 {
	if a.nanosiemensLazy != nil {
		return *a.nanosiemensLazy
	}
	nanosiemens := a.convertFromBase(ElectricAdmittanceNanosiemens)
	a.nanosiemensLazy = &nanosiemens
	return nanosiemens
}

// Microsiemens returns the value in Microsiemens.
func (a *ElectricAdmittance) Microsiemens() float64 {
	if a.microsiemensLazy != nil {
		return *a.microsiemensLazy
	}
	microsiemens := a.convertFromBase(ElectricAdmittanceMicrosiemens)
	a.microsiemensLazy = &microsiemens
	return microsiemens
}

// Millisiemens returns the value in Millisiemens.
func (a *ElectricAdmittance) Millisiemens() float64 {
	if a.millisiemensLazy != nil {
		return *a.millisiemensLazy
	}
	millisiemens := a.convertFromBase(ElectricAdmittanceMillisiemens)
	a.millisiemensLazy = &millisiemens
	return millisiemens
}


// ToDto creates an ElectricAdmittanceDto representation.
func (a *ElectricAdmittance) ToDto(holdInUnit *ElectricAdmittanceUnits) ElectricAdmittanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricAdmittanceSiemens // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricAdmittanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricAdmittanceDto representation.
func (a *ElectricAdmittance) ToDtoJSON(holdInUnit *ElectricAdmittanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricAdmittance to a specific unit value.
func (a *ElectricAdmittance) Convert(toUnit ElectricAdmittanceUnits) float64 {
	switch toUnit { 
    case ElectricAdmittanceSiemens:
		return a.Siemens()
    case ElectricAdmittanceNanosiemens:
		return a.Nanosiemens()
    case ElectricAdmittanceMicrosiemens:
		return a.Microsiemens()
    case ElectricAdmittanceMillisiemens:
		return a.Millisiemens()
	default:
		return 0
	}
}

func (a *ElectricAdmittance) convertFromBase(toUnit ElectricAdmittanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricAdmittanceSiemens:
		return (value) 
	case ElectricAdmittanceNanosiemens:
		return ((value) / 1e-09) 
	case ElectricAdmittanceMicrosiemens:
		return ((value) / 1e-06) 
	case ElectricAdmittanceMillisiemens:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricAdmittance) convertToBase(value float64, fromUnit ElectricAdmittanceUnits) float64 {
	switch fromUnit { 
	case ElectricAdmittanceSiemens:
		return (value) 
	case ElectricAdmittanceNanosiemens:
		return ((value) * 1e-09) 
	case ElectricAdmittanceMicrosiemens:
		return ((value) * 1e-06) 
	case ElectricAdmittanceMillisiemens:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricAdmittance) String() string {
	return a.ToString(ElectricAdmittanceSiemens, 2)
}

// ToString formats the ElectricAdmittance to string.
// fractionalDigits -1 for not mention
func (a *ElectricAdmittance) ToString(unit ElectricAdmittanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricAdmittance) getUnitAbbreviation(unit ElectricAdmittanceUnits) string {
	switch unit { 
	case ElectricAdmittanceSiemens:
		return "S" 
	case ElectricAdmittanceNanosiemens:
		return "nS" 
	case ElectricAdmittanceMicrosiemens:
		return "μS" 
	case ElectricAdmittanceMillisiemens:
		return "mS" 
	default:
		return ""
	}
}

// Check if the given ElectricAdmittance are equals to the current ElectricAdmittance
func (a *ElectricAdmittance) Equals(other *ElectricAdmittance) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricAdmittance are equals to the current ElectricAdmittance
func (a *ElectricAdmittance) CompareTo(other *ElectricAdmittance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricAdmittance to the current ElectricAdmittance.
func (a *ElectricAdmittance) Add(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricAdmittance to the current ElectricAdmittance.
func (a *ElectricAdmittance) Subtract(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricAdmittance to the current ElectricAdmittance.
func (a *ElectricAdmittance) Multiply(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value * other.BaseValue()}
}

// Divide the given ElectricAdmittance to the current ElectricAdmittance.
func (a *ElectricAdmittance) Divide(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value / other.BaseValue()}
}