package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialUnits enumeration
type ElectricPotentialUnits string

const (
    
        // 
        ElectricPotentialVolt ElectricPotentialUnits = "Volt"
        // 
        ElectricPotentialNanovolt ElectricPotentialUnits = "Nanovolt"
        // 
        ElectricPotentialMicrovolt ElectricPotentialUnits = "Microvolt"
        // 
        ElectricPotentialMillivolt ElectricPotentialUnits = "Millivolt"
        // 
        ElectricPotentialKilovolt ElectricPotentialUnits = "Kilovolt"
        // 
        ElectricPotentialMegavolt ElectricPotentialUnits = "Megavolt"
)

// ElectricPotentialDto represents an ElectricPotential
type ElectricPotentialDto struct {
	Value float64
	Unit  ElectricPotentialUnits
}

// ElectricPotentialDtoFactory struct to group related functions
type ElectricPotentialDtoFactory struct{}

func (udf ElectricPotentialDtoFactory) FromJSON(data []byte) (*ElectricPotentialDto, error) {
	a := ElectricPotentialDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricPotentialDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricPotential struct
type ElectricPotential struct {
	value       float64
    
    voltsLazy *float64 
    nanovoltsLazy *float64 
    microvoltsLazy *float64 
    millivoltsLazy *float64 
    kilovoltsLazy *float64 
    megavoltsLazy *float64 
}

// ElectricPotentialFactory struct to group related functions
type ElectricPotentialFactory struct{}

func (uf ElectricPotentialFactory) CreateElectricPotential(value float64, unit ElectricPotentialUnits) (*ElectricPotential, error) {
	return newElectricPotential(value, unit)
}

func (uf ElectricPotentialFactory) FromDto(dto ElectricPotentialDto) (*ElectricPotential, error) {
	return newElectricPotential(dto.Value, dto.Unit)
}

func (uf ElectricPotentialFactory) FromDtoJSON(data []byte) (*ElectricPotential, error) {
	unitDto, err := ElectricPotentialDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricPotentialFactory{}.FromDto(*unitDto)
}


// FromVolt creates a new ElectricPotential instance from Volt.
func (uf ElectricPotentialFactory) FromVolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialVolt)
}

// FromNanovolt creates a new ElectricPotential instance from Nanovolt.
func (uf ElectricPotentialFactory) FromNanovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialNanovolt)
}

// FromMicrovolt creates a new ElectricPotential instance from Microvolt.
func (uf ElectricPotentialFactory) FromMicrovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMicrovolt)
}

// FromMillivolt creates a new ElectricPotential instance from Millivolt.
func (uf ElectricPotentialFactory) FromMillivolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMillivolt)
}

// FromKilovolt creates a new ElectricPotential instance from Kilovolt.
func (uf ElectricPotentialFactory) FromKilovolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialKilovolt)
}

// FromMegavolt creates a new ElectricPotential instance from Megavolt.
func (uf ElectricPotentialFactory) FromMegavolts(value float64) (*ElectricPotential, error) {
	return newElectricPotential(value, ElectricPotentialMegavolt)
}




// newElectricPotential creates a new ElectricPotential.
func newElectricPotential(value float64, fromUnit ElectricPotentialUnits) (*ElectricPotential, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotential{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotential in Volt.
func (a *ElectricPotential) BaseValue() float64 {
	return a.value
}


// Volt returns the value in Volt.
func (a *ElectricPotential) Volts() float64 {
	if a.voltsLazy != nil {
		return *a.voltsLazy
	}
	volts := a.convertFromBase(ElectricPotentialVolt)
	a.voltsLazy = &volts
	return volts
}

// Nanovolt returns the value in Nanovolt.
func (a *ElectricPotential) Nanovolts() float64 {
	if a.nanovoltsLazy != nil {
		return *a.nanovoltsLazy
	}
	nanovolts := a.convertFromBase(ElectricPotentialNanovolt)
	a.nanovoltsLazy = &nanovolts
	return nanovolts
}

// Microvolt returns the value in Microvolt.
func (a *ElectricPotential) Microvolts() float64 {
	if a.microvoltsLazy != nil {
		return *a.microvoltsLazy
	}
	microvolts := a.convertFromBase(ElectricPotentialMicrovolt)
	a.microvoltsLazy = &microvolts
	return microvolts
}

// Millivolt returns the value in Millivolt.
func (a *ElectricPotential) Millivolts() float64 {
	if a.millivoltsLazy != nil {
		return *a.millivoltsLazy
	}
	millivolts := a.convertFromBase(ElectricPotentialMillivolt)
	a.millivoltsLazy = &millivolts
	return millivolts
}

// Kilovolt returns the value in Kilovolt.
func (a *ElectricPotential) Kilovolts() float64 {
	if a.kilovoltsLazy != nil {
		return *a.kilovoltsLazy
	}
	kilovolts := a.convertFromBase(ElectricPotentialKilovolt)
	a.kilovoltsLazy = &kilovolts
	return kilovolts
}

// Megavolt returns the value in Megavolt.
func (a *ElectricPotential) Megavolts() float64 {
	if a.megavoltsLazy != nil {
		return *a.megavoltsLazy
	}
	megavolts := a.convertFromBase(ElectricPotentialMegavolt)
	a.megavoltsLazy = &megavolts
	return megavolts
}


// ToDto creates an ElectricPotentialDto representation.
func (a *ElectricPotential) ToDto(holdInUnit *ElectricPotentialUnits) ElectricPotentialDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialVolt // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricPotentialDto representation.
func (a *ElectricPotential) ToDtoJSON(holdInUnit *ElectricPotentialUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricPotential to a specific unit value.
func (a *ElectricPotential) Convert(toUnit ElectricPotentialUnits) float64 {
	switch toUnit { 
    case ElectricPotentialVolt:
		return a.Volts()
    case ElectricPotentialNanovolt:
		return a.Nanovolts()
    case ElectricPotentialMicrovolt:
		return a.Microvolts()
    case ElectricPotentialMillivolt:
		return a.Millivolts()
    case ElectricPotentialKilovolt:
		return a.Kilovolts()
    case ElectricPotentialMegavolt:
		return a.Megavolts()
	default:
		return 0
	}
}

func (a *ElectricPotential) convertFromBase(toUnit ElectricPotentialUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialVolt:
		return (value) 
	case ElectricPotentialNanovolt:
		return ((value) / 1e-09) 
	case ElectricPotentialMicrovolt:
		return ((value) / 1e-06) 
	case ElectricPotentialMillivolt:
		return ((value) / 0.001) 
	case ElectricPotentialKilovolt:
		return ((value) / 1000.0) 
	case ElectricPotentialMegavolt:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotential) convertToBase(value float64, fromUnit ElectricPotentialUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialVolt:
		return (value) 
	case ElectricPotentialNanovolt:
		return ((value) * 1e-09) 
	case ElectricPotentialMicrovolt:
		return ((value) * 1e-06) 
	case ElectricPotentialMillivolt:
		return ((value) * 0.001) 
	case ElectricPotentialKilovolt:
		return ((value) * 1000.0) 
	case ElectricPotentialMegavolt:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricPotential) String() string {
	return a.ToString(ElectricPotentialVolt, 2)
}

// ToString formats the ElectricPotential to string.
// fractionalDigits -1 for not mention
func (a *ElectricPotential) ToString(unit ElectricPotentialUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricPotential) getUnitAbbreviation(unit ElectricPotentialUnits) string {
	switch unit { 
	case ElectricPotentialVolt:
		return "V" 
	case ElectricPotentialNanovolt:
		return "nV" 
	case ElectricPotentialMicrovolt:
		return "μV" 
	case ElectricPotentialMillivolt:
		return "mV" 
	case ElectricPotentialKilovolt:
		return "kV" 
	case ElectricPotentialMegavolt:
		return "MV" 
	default:
		return ""
	}
}

// Check if the given ElectricPotential are equals to the current ElectricPotential
func (a *ElectricPotential) Equals(other *ElectricPotential) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricPotential are equals to the current ElectricPotential
func (a *ElectricPotential) CompareTo(other *ElectricPotential) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricPotential to the current ElectricPotential.
func (a *ElectricPotential) Add(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricPotential to the current ElectricPotential.
func (a *ElectricPotential) Subtract(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricPotential to the current ElectricPotential.
func (a *ElectricPotential) Multiply(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value * other.BaseValue()}
}

// Divide the given ElectricPotential to the current ElectricPotential.
func (a *ElectricPotential) Divide(other *ElectricPotential) *ElectricPotential {
	return &ElectricPotential{value: a.value / other.BaseValue()}
}