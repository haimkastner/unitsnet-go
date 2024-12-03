package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricCurrentUnits enumeration
type ElectricCurrentUnits string

const (
    
        // 
        ElectricCurrentAmpere ElectricCurrentUnits = "Ampere"
        // 
        ElectricCurrentFemtoampere ElectricCurrentUnits = "Femtoampere"
        // 
        ElectricCurrentPicoampere ElectricCurrentUnits = "Picoampere"
        // 
        ElectricCurrentNanoampere ElectricCurrentUnits = "Nanoampere"
        // 
        ElectricCurrentMicroampere ElectricCurrentUnits = "Microampere"
        // 
        ElectricCurrentMilliampere ElectricCurrentUnits = "Milliampere"
        // 
        ElectricCurrentCentiampere ElectricCurrentUnits = "Centiampere"
        // 
        ElectricCurrentKiloampere ElectricCurrentUnits = "Kiloampere"
        // 
        ElectricCurrentMegaampere ElectricCurrentUnits = "Megaampere"
)

// ElectricCurrentDto represents an ElectricCurrent
type ElectricCurrentDto struct {
	Value float64
	Unit  ElectricCurrentUnits
}

// ElectricCurrentDtoFactory struct to group related functions
type ElectricCurrentDtoFactory struct{}

func (udf ElectricCurrentDtoFactory) FromJSON(data []byte) (*ElectricCurrentDto, error) {
	a := ElectricCurrentDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricCurrentDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricCurrent struct
type ElectricCurrent struct {
	value       float64
    
    amperesLazy *float64 
    femtoamperesLazy *float64 
    picoamperesLazy *float64 
    nanoamperesLazy *float64 
    microamperesLazy *float64 
    milliamperesLazy *float64 
    centiamperesLazy *float64 
    kiloamperesLazy *float64 
    megaamperesLazy *float64 
}

// ElectricCurrentFactory struct to group related functions
type ElectricCurrentFactory struct{}

func (uf ElectricCurrentFactory) CreateElectricCurrent(value float64, unit ElectricCurrentUnits) (*ElectricCurrent, error) {
	return newElectricCurrent(value, unit)
}

func (uf ElectricCurrentFactory) FromDto(dto ElectricCurrentDto) (*ElectricCurrent, error) {
	return newElectricCurrent(dto.Value, dto.Unit)
}

func (uf ElectricCurrentFactory) FromDtoJSON(data []byte) (*ElectricCurrent, error) {
	unitDto, err := ElectricCurrentDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricCurrentFactory{}.FromDto(*unitDto)
}


// FromAmpere creates a new ElectricCurrent instance from Ampere.
func (uf ElectricCurrentFactory) FromAmperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentAmpere)
}

// FromFemtoampere creates a new ElectricCurrent instance from Femtoampere.
func (uf ElectricCurrentFactory) FromFemtoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentFemtoampere)
}

// FromPicoampere creates a new ElectricCurrent instance from Picoampere.
func (uf ElectricCurrentFactory) FromPicoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentPicoampere)
}

// FromNanoampere creates a new ElectricCurrent instance from Nanoampere.
func (uf ElectricCurrentFactory) FromNanoamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentNanoampere)
}

// FromMicroampere creates a new ElectricCurrent instance from Microampere.
func (uf ElectricCurrentFactory) FromMicroamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMicroampere)
}

// FromMilliampere creates a new ElectricCurrent instance from Milliampere.
func (uf ElectricCurrentFactory) FromMilliamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMilliampere)
}

// FromCentiampere creates a new ElectricCurrent instance from Centiampere.
func (uf ElectricCurrentFactory) FromCentiamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentCentiampere)
}

// FromKiloampere creates a new ElectricCurrent instance from Kiloampere.
func (uf ElectricCurrentFactory) FromKiloamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentKiloampere)
}

// FromMegaampere creates a new ElectricCurrent instance from Megaampere.
func (uf ElectricCurrentFactory) FromMegaamperes(value float64) (*ElectricCurrent, error) {
	return newElectricCurrent(value, ElectricCurrentMegaampere)
}




// newElectricCurrent creates a new ElectricCurrent.
func newElectricCurrent(value float64, fromUnit ElectricCurrentUnits) (*ElectricCurrent, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricCurrent{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricCurrent in Ampere.
func (a *ElectricCurrent) BaseValue() float64 {
	return a.value
}


// Ampere returns the value in Ampere.
func (a *ElectricCurrent) Amperes() float64 {
	if a.amperesLazy != nil {
		return *a.amperesLazy
	}
	amperes := a.convertFromBase(ElectricCurrentAmpere)
	a.amperesLazy = &amperes
	return amperes
}

// Femtoampere returns the value in Femtoampere.
func (a *ElectricCurrent) Femtoamperes() float64 {
	if a.femtoamperesLazy != nil {
		return *a.femtoamperesLazy
	}
	femtoamperes := a.convertFromBase(ElectricCurrentFemtoampere)
	a.femtoamperesLazy = &femtoamperes
	return femtoamperes
}

// Picoampere returns the value in Picoampere.
func (a *ElectricCurrent) Picoamperes() float64 {
	if a.picoamperesLazy != nil {
		return *a.picoamperesLazy
	}
	picoamperes := a.convertFromBase(ElectricCurrentPicoampere)
	a.picoamperesLazy = &picoamperes
	return picoamperes
}

// Nanoampere returns the value in Nanoampere.
func (a *ElectricCurrent) Nanoamperes() float64 {
	if a.nanoamperesLazy != nil {
		return *a.nanoamperesLazy
	}
	nanoamperes := a.convertFromBase(ElectricCurrentNanoampere)
	a.nanoamperesLazy = &nanoamperes
	return nanoamperes
}

// Microampere returns the value in Microampere.
func (a *ElectricCurrent) Microamperes() float64 {
	if a.microamperesLazy != nil {
		return *a.microamperesLazy
	}
	microamperes := a.convertFromBase(ElectricCurrentMicroampere)
	a.microamperesLazy = &microamperes
	return microamperes
}

// Milliampere returns the value in Milliampere.
func (a *ElectricCurrent) Milliamperes() float64 {
	if a.milliamperesLazy != nil {
		return *a.milliamperesLazy
	}
	milliamperes := a.convertFromBase(ElectricCurrentMilliampere)
	a.milliamperesLazy = &milliamperes
	return milliamperes
}

// Centiampere returns the value in Centiampere.
func (a *ElectricCurrent) Centiamperes() float64 {
	if a.centiamperesLazy != nil {
		return *a.centiamperesLazy
	}
	centiamperes := a.convertFromBase(ElectricCurrentCentiampere)
	a.centiamperesLazy = &centiamperes
	return centiamperes
}

// Kiloampere returns the value in Kiloampere.
func (a *ElectricCurrent) Kiloamperes() float64 {
	if a.kiloamperesLazy != nil {
		return *a.kiloamperesLazy
	}
	kiloamperes := a.convertFromBase(ElectricCurrentKiloampere)
	a.kiloamperesLazy = &kiloamperes
	return kiloamperes
}

// Megaampere returns the value in Megaampere.
func (a *ElectricCurrent) Megaamperes() float64 {
	if a.megaamperesLazy != nil {
		return *a.megaamperesLazy
	}
	megaamperes := a.convertFromBase(ElectricCurrentMegaampere)
	a.megaamperesLazy = &megaamperes
	return megaamperes
}


// ToDto creates an ElectricCurrentDto representation.
func (a *ElectricCurrent) ToDto(holdInUnit *ElectricCurrentUnits) ElectricCurrentDto {
	if holdInUnit == nil {
		defaultUnit := ElectricCurrentAmpere // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricCurrentDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricCurrentDto representation.
func (a *ElectricCurrent) ToDtoJSON(holdInUnit *ElectricCurrentUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricCurrent to a specific unit value.
func (a *ElectricCurrent) Convert(toUnit ElectricCurrentUnits) float64 {
	switch toUnit { 
    case ElectricCurrentAmpere:
		return a.Amperes()
    case ElectricCurrentFemtoampere:
		return a.Femtoamperes()
    case ElectricCurrentPicoampere:
		return a.Picoamperes()
    case ElectricCurrentNanoampere:
		return a.Nanoamperes()
    case ElectricCurrentMicroampere:
		return a.Microamperes()
    case ElectricCurrentMilliampere:
		return a.Milliamperes()
    case ElectricCurrentCentiampere:
		return a.Centiamperes()
    case ElectricCurrentKiloampere:
		return a.Kiloamperes()
    case ElectricCurrentMegaampere:
		return a.Megaamperes()
	default:
		return 0
	}
}

func (a *ElectricCurrent) convertFromBase(toUnit ElectricCurrentUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricCurrentAmpere:
		return (value) 
	case ElectricCurrentFemtoampere:
		return ((value) / 1e-15) 
	case ElectricCurrentPicoampere:
		return ((value) / 1e-12) 
	case ElectricCurrentNanoampere:
		return ((value) / 1e-09) 
	case ElectricCurrentMicroampere:
		return ((value) / 1e-06) 
	case ElectricCurrentMilliampere:
		return ((value) / 0.001) 
	case ElectricCurrentCentiampere:
		return ((value) / 0.01) 
	case ElectricCurrentKiloampere:
		return ((value) / 1000.0) 
	case ElectricCurrentMegaampere:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricCurrent) convertToBase(value float64, fromUnit ElectricCurrentUnits) float64 {
	switch fromUnit { 
	case ElectricCurrentAmpere:
		return (value) 
	case ElectricCurrentFemtoampere:
		return ((value) * 1e-15) 
	case ElectricCurrentPicoampere:
		return ((value) * 1e-12) 
	case ElectricCurrentNanoampere:
		return ((value) * 1e-09) 
	case ElectricCurrentMicroampere:
		return ((value) * 1e-06) 
	case ElectricCurrentMilliampere:
		return ((value) * 0.001) 
	case ElectricCurrentCentiampere:
		return ((value) * 0.01) 
	case ElectricCurrentKiloampere:
		return ((value) * 1000.0) 
	case ElectricCurrentMegaampere:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricCurrent) String() string {
	return a.ToString(ElectricCurrentAmpere, 2)
}

// ToString formats the ElectricCurrent to string.
// fractionalDigits -1 for not mention
func (a *ElectricCurrent) ToString(unit ElectricCurrentUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricCurrent) getUnitAbbreviation(unit ElectricCurrentUnits) string {
	switch unit { 
	case ElectricCurrentAmpere:
		return "A" 
	case ElectricCurrentFemtoampere:
		return "fA" 
	case ElectricCurrentPicoampere:
		return "pA" 
	case ElectricCurrentNanoampere:
		return "nA" 
	case ElectricCurrentMicroampere:
		return "μA" 
	case ElectricCurrentMilliampere:
		return "mA" 
	case ElectricCurrentCentiampere:
		return "cA" 
	case ElectricCurrentKiloampere:
		return "kA" 
	case ElectricCurrentMegaampere:
		return "MA" 
	default:
		return ""
	}
}

// Check if the given ElectricCurrent are equals to the current ElectricCurrent
func (a *ElectricCurrent) Equals(other *ElectricCurrent) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricCurrent are equals to the current ElectricCurrent
func (a *ElectricCurrent) CompareTo(other *ElectricCurrent) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricCurrent to the current ElectricCurrent.
func (a *ElectricCurrent) Add(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricCurrent to the current ElectricCurrent.
func (a *ElectricCurrent) Subtract(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricCurrent to the current ElectricCurrent.
func (a *ElectricCurrent) Multiply(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value * other.BaseValue()}
}

// Divide the given ElectricCurrent to the current ElectricCurrent.
func (a *ElectricCurrent) Divide(other *ElectricCurrent) *ElectricCurrent {
	return &ElectricCurrent{value: a.value / other.BaseValue()}
}