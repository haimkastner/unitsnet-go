// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// CapacitanceUnits enumeration
type CapacitanceUnits string

const (
    
        // 
        CapacitanceFarad CapacitanceUnits = "Farad"
        // 
        CapacitancePicofarad CapacitanceUnits = "Picofarad"
        // 
        CapacitanceNanofarad CapacitanceUnits = "Nanofarad"
        // 
        CapacitanceMicrofarad CapacitanceUnits = "Microfarad"
        // 
        CapacitanceMillifarad CapacitanceUnits = "Millifarad"
        // 
        CapacitanceKilofarad CapacitanceUnits = "Kilofarad"
        // 
        CapacitanceMegafarad CapacitanceUnits = "Megafarad"
)

// CapacitanceDto represents an Capacitance
type CapacitanceDto struct {
	Value float64
	Unit  CapacitanceUnits
}

// CapacitanceDtoFactory struct to group related functions
type CapacitanceDtoFactory struct{}

func (udf CapacitanceDtoFactory) FromJSON(data []byte) (*CapacitanceDto, error) {
	a := CapacitanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a CapacitanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Capacitance struct
type Capacitance struct {
	value       float64
    
    faradsLazy *float64 
    picofaradsLazy *float64 
    nanofaradsLazy *float64 
    microfaradsLazy *float64 
    millifaradsLazy *float64 
    kilofaradsLazy *float64 
    megafaradsLazy *float64 
}

// CapacitanceFactory struct to group related functions
type CapacitanceFactory struct{}

func (uf CapacitanceFactory) CreateCapacitance(value float64, unit CapacitanceUnits) (*Capacitance, error) {
	return newCapacitance(value, unit)
}

func (uf CapacitanceFactory) FromDto(dto CapacitanceDto) (*Capacitance, error) {
	return newCapacitance(dto.Value, dto.Unit)
}

func (uf CapacitanceFactory) FromDtoJSON(data []byte) (*Capacitance, error) {
	unitDto, err := CapacitanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return CapacitanceFactory{}.FromDto(*unitDto)
}


// FromFarad creates a new Capacitance instance from Farad.
func (uf CapacitanceFactory) FromFarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceFarad)
}

// FromPicofarad creates a new Capacitance instance from Picofarad.
func (uf CapacitanceFactory) FromPicofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitancePicofarad)
}

// FromNanofarad creates a new Capacitance instance from Nanofarad.
func (uf CapacitanceFactory) FromNanofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceNanofarad)
}

// FromMicrofarad creates a new Capacitance instance from Microfarad.
func (uf CapacitanceFactory) FromMicrofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMicrofarad)
}

// FromMillifarad creates a new Capacitance instance from Millifarad.
func (uf CapacitanceFactory) FromMillifarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMillifarad)
}

// FromKilofarad creates a new Capacitance instance from Kilofarad.
func (uf CapacitanceFactory) FromKilofarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceKilofarad)
}

// FromMegafarad creates a new Capacitance instance from Megafarad.
func (uf CapacitanceFactory) FromMegafarads(value float64) (*Capacitance, error) {
	return newCapacitance(value, CapacitanceMegafarad)
}




// newCapacitance creates a new Capacitance.
func newCapacitance(value float64, fromUnit CapacitanceUnits) (*Capacitance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Capacitance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Capacitance in Farad.
func (a *Capacitance) BaseValue() float64 {
	return a.value
}


// Farad returns the value in Farad.
func (a *Capacitance) Farads() float64 {
	if a.faradsLazy != nil {
		return *a.faradsLazy
	}
	farads := a.convertFromBase(CapacitanceFarad)
	a.faradsLazy = &farads
	return farads
}

// Picofarad returns the value in Picofarad.
func (a *Capacitance) Picofarads() float64 {
	if a.picofaradsLazy != nil {
		return *a.picofaradsLazy
	}
	picofarads := a.convertFromBase(CapacitancePicofarad)
	a.picofaradsLazy = &picofarads
	return picofarads
}

// Nanofarad returns the value in Nanofarad.
func (a *Capacitance) Nanofarads() float64 {
	if a.nanofaradsLazy != nil {
		return *a.nanofaradsLazy
	}
	nanofarads := a.convertFromBase(CapacitanceNanofarad)
	a.nanofaradsLazy = &nanofarads
	return nanofarads
}

// Microfarad returns the value in Microfarad.
func (a *Capacitance) Microfarads() float64 {
	if a.microfaradsLazy != nil {
		return *a.microfaradsLazy
	}
	microfarads := a.convertFromBase(CapacitanceMicrofarad)
	a.microfaradsLazy = &microfarads
	return microfarads
}

// Millifarad returns the value in Millifarad.
func (a *Capacitance) Millifarads() float64 {
	if a.millifaradsLazy != nil {
		return *a.millifaradsLazy
	}
	millifarads := a.convertFromBase(CapacitanceMillifarad)
	a.millifaradsLazy = &millifarads
	return millifarads
}

// Kilofarad returns the value in Kilofarad.
func (a *Capacitance) Kilofarads() float64 {
	if a.kilofaradsLazy != nil {
		return *a.kilofaradsLazy
	}
	kilofarads := a.convertFromBase(CapacitanceKilofarad)
	a.kilofaradsLazy = &kilofarads
	return kilofarads
}

// Megafarad returns the value in Megafarad.
func (a *Capacitance) Megafarads() float64 {
	if a.megafaradsLazy != nil {
		return *a.megafaradsLazy
	}
	megafarads := a.convertFromBase(CapacitanceMegafarad)
	a.megafaradsLazy = &megafarads
	return megafarads
}


// ToDto creates an CapacitanceDto representation.
func (a *Capacitance) ToDto(holdInUnit *CapacitanceUnits) CapacitanceDto {
	if holdInUnit == nil {
		defaultUnit := CapacitanceFarad // Default value
		holdInUnit = &defaultUnit
	}

	return CapacitanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an CapacitanceDto representation.
func (a *Capacitance) ToDtoJSON(holdInUnit *CapacitanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Capacitance to a specific unit value.
func (a *Capacitance) Convert(toUnit CapacitanceUnits) float64 {
	switch toUnit { 
    case CapacitanceFarad:
		return a.Farads()
    case CapacitancePicofarad:
		return a.Picofarads()
    case CapacitanceNanofarad:
		return a.Nanofarads()
    case CapacitanceMicrofarad:
		return a.Microfarads()
    case CapacitanceMillifarad:
		return a.Millifarads()
    case CapacitanceKilofarad:
		return a.Kilofarads()
    case CapacitanceMegafarad:
		return a.Megafarads()
	default:
		return 0
	}
}

func (a *Capacitance) convertFromBase(toUnit CapacitanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case CapacitanceFarad:
		return (value) 
	case CapacitancePicofarad:
		return ((value) / 1e-12) 
	case CapacitanceNanofarad:
		return ((value) / 1e-09) 
	case CapacitanceMicrofarad:
		return ((value) / 1e-06) 
	case CapacitanceMillifarad:
		return ((value) / 0.001) 
	case CapacitanceKilofarad:
		return ((value) / 1000.0) 
	case CapacitanceMegafarad:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Capacitance) convertToBase(value float64, fromUnit CapacitanceUnits) float64 {
	switch fromUnit { 
	case CapacitanceFarad:
		return (value) 
	case CapacitancePicofarad:
		return ((value) * 1e-12) 
	case CapacitanceNanofarad:
		return ((value) * 1e-09) 
	case CapacitanceMicrofarad:
		return ((value) * 1e-06) 
	case CapacitanceMillifarad:
		return ((value) * 0.001) 
	case CapacitanceKilofarad:
		return ((value) * 1000.0) 
	case CapacitanceMegafarad:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Capacitance) String() string {
	return a.ToString(CapacitanceFarad, 2)
}

// ToString formats the Capacitance to string.
// fractionalDigits -1 for not mention
func (a *Capacitance) ToString(unit CapacitanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Capacitance) getUnitAbbreviation(unit CapacitanceUnits) string {
	switch unit { 
	case CapacitanceFarad:
		return "F" 
	case CapacitancePicofarad:
		return "pF" 
	case CapacitanceNanofarad:
		return "nF" 
	case CapacitanceMicrofarad:
		return "μF" 
	case CapacitanceMillifarad:
		return "mF" 
	case CapacitanceKilofarad:
		return "kF" 
	case CapacitanceMegafarad:
		return "MF" 
	default:
		return ""
	}
}

// Check if the given Capacitance are equals to the current Capacitance
func (a *Capacitance) Equals(other *Capacitance) bool {
	return a.value == other.BaseValue()
}

// Check if the given Capacitance are equals to the current Capacitance
func (a *Capacitance) CompareTo(other *Capacitance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Capacitance to the current Capacitance.
func (a *Capacitance) Add(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value + other.BaseValue()}
}

// Subtract the given Capacitance to the current Capacitance.
func (a *Capacitance) Subtract(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value - other.BaseValue()}
}

// Multiply the given Capacitance to the current Capacitance.
func (a *Capacitance) Multiply(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value * other.BaseValue()}
}

// Divide the given Capacitance to the current Capacitance.
func (a *Capacitance) Divide(other *Capacitance) *Capacitance {
	return &Capacitance{value: a.value / other.BaseValue()}
}