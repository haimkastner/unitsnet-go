package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialAcUnits enumeration
type ElectricPotentialAcUnits string

const (
    
        // 
        ElectricPotentialAcVoltAc ElectricPotentialAcUnits = "VoltAc"
        // 
        ElectricPotentialAcMicrovoltAc ElectricPotentialAcUnits = "MicrovoltAc"
        // 
        ElectricPotentialAcMillivoltAc ElectricPotentialAcUnits = "MillivoltAc"
        // 
        ElectricPotentialAcKilovoltAc ElectricPotentialAcUnits = "KilovoltAc"
        // 
        ElectricPotentialAcMegavoltAc ElectricPotentialAcUnits = "MegavoltAc"
)

// ElectricPotentialAcDto represents an ElectricPotentialAc
type ElectricPotentialAcDto struct {
	Value float64
	Unit  ElectricPotentialAcUnits
}

// ElectricPotentialAcDtoFactory struct to group related functions
type ElectricPotentialAcDtoFactory struct{}

func (udf ElectricPotentialAcDtoFactory) FromJSON(data []byte) (*ElectricPotentialAcDto, error) {
	a := ElectricPotentialAcDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricPotentialAcDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricPotentialAc struct
type ElectricPotentialAc struct {
	value       float64
    
    volts_acLazy *float64 
    microvolts_acLazy *float64 
    millivolts_acLazy *float64 
    kilovolts_acLazy *float64 
    megavolts_acLazy *float64 
}

// ElectricPotentialAcFactory struct to group related functions
type ElectricPotentialAcFactory struct{}

func (uf ElectricPotentialAcFactory) CreateElectricPotentialAc(value float64, unit ElectricPotentialAcUnits) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, unit)
}

func (uf ElectricPotentialAcFactory) FromDto(dto ElectricPotentialAcDto) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(dto.Value, dto.Unit)
}

func (uf ElectricPotentialAcFactory) FromDtoJSON(data []byte) (*ElectricPotentialAc, error) {
	unitDto, err := ElectricPotentialAcDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricPotentialAcFactory{}.FromDto(*unitDto)
}


// FromVoltAc creates a new ElectricPotentialAc instance from VoltAc.
func (uf ElectricPotentialAcFactory) FromVoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcVoltAc)
}

// FromMicrovoltAc creates a new ElectricPotentialAc instance from MicrovoltAc.
func (uf ElectricPotentialAcFactory) FromMicrovoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMicrovoltAc)
}

// FromMillivoltAc creates a new ElectricPotentialAc instance from MillivoltAc.
func (uf ElectricPotentialAcFactory) FromMillivoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMillivoltAc)
}

// FromKilovoltAc creates a new ElectricPotentialAc instance from KilovoltAc.
func (uf ElectricPotentialAcFactory) FromKilovoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcKilovoltAc)
}

// FromMegavoltAc creates a new ElectricPotentialAc instance from MegavoltAc.
func (uf ElectricPotentialAcFactory) FromMegavoltsAc(value float64) (*ElectricPotentialAc, error) {
	return newElectricPotentialAc(value, ElectricPotentialAcMegavoltAc)
}




// newElectricPotentialAc creates a new ElectricPotentialAc.
func newElectricPotentialAc(value float64, fromUnit ElectricPotentialAcUnits) (*ElectricPotentialAc, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotentialAc{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialAc in VoltAc.
func (a *ElectricPotentialAc) BaseValue() float64 {
	return a.value
}


// VoltAc returns the value in VoltAc.
func (a *ElectricPotentialAc) VoltsAc() float64 {
	if a.volts_acLazy != nil {
		return *a.volts_acLazy
	}
	volts_ac := a.convertFromBase(ElectricPotentialAcVoltAc)
	a.volts_acLazy = &volts_ac
	return volts_ac
}

// MicrovoltAc returns the value in MicrovoltAc.
func (a *ElectricPotentialAc) MicrovoltsAc() float64 {
	if a.microvolts_acLazy != nil {
		return *a.microvolts_acLazy
	}
	microvolts_ac := a.convertFromBase(ElectricPotentialAcMicrovoltAc)
	a.microvolts_acLazy = &microvolts_ac
	return microvolts_ac
}

// MillivoltAc returns the value in MillivoltAc.
func (a *ElectricPotentialAc) MillivoltsAc() float64 {
	if a.millivolts_acLazy != nil {
		return *a.millivolts_acLazy
	}
	millivolts_ac := a.convertFromBase(ElectricPotentialAcMillivoltAc)
	a.millivolts_acLazy = &millivolts_ac
	return millivolts_ac
}

// KilovoltAc returns the value in KilovoltAc.
func (a *ElectricPotentialAc) KilovoltsAc() float64 {
	if a.kilovolts_acLazy != nil {
		return *a.kilovolts_acLazy
	}
	kilovolts_ac := a.convertFromBase(ElectricPotentialAcKilovoltAc)
	a.kilovolts_acLazy = &kilovolts_ac
	return kilovolts_ac
}

// MegavoltAc returns the value in MegavoltAc.
func (a *ElectricPotentialAc) MegavoltsAc() float64 {
	if a.megavolts_acLazy != nil {
		return *a.megavolts_acLazy
	}
	megavolts_ac := a.convertFromBase(ElectricPotentialAcMegavoltAc)
	a.megavolts_acLazy = &megavolts_ac
	return megavolts_ac
}


// ToDto creates an ElectricPotentialAcDto representation.
func (a *ElectricPotentialAc) ToDto(holdInUnit *ElectricPotentialAcUnits) ElectricPotentialAcDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialAcVoltAc // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialAcDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricPotentialAcDto representation.
func (a *ElectricPotentialAc) ToDtoJSON(holdInUnit *ElectricPotentialAcUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricPotentialAc to a specific unit value.
func (a *ElectricPotentialAc) Convert(toUnit ElectricPotentialAcUnits) float64 {
	switch toUnit { 
    case ElectricPotentialAcVoltAc:
		return a.VoltsAc()
    case ElectricPotentialAcMicrovoltAc:
		return a.MicrovoltsAc()
    case ElectricPotentialAcMillivoltAc:
		return a.MillivoltsAc()
    case ElectricPotentialAcKilovoltAc:
		return a.KilovoltsAc()
    case ElectricPotentialAcMegavoltAc:
		return a.MegavoltsAc()
	default:
		return 0
	}
}

func (a *ElectricPotentialAc) convertFromBase(toUnit ElectricPotentialAcUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialAcVoltAc:
		return (value) 
	case ElectricPotentialAcMicrovoltAc:
		return ((value) / 1e-06) 
	case ElectricPotentialAcMillivoltAc:
		return ((value) / 0.001) 
	case ElectricPotentialAcKilovoltAc:
		return ((value) / 1000.0) 
	case ElectricPotentialAcMegavoltAc:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialAc) convertToBase(value float64, fromUnit ElectricPotentialAcUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialAcVoltAc:
		return (value) 
	case ElectricPotentialAcMicrovoltAc:
		return ((value) * 1e-06) 
	case ElectricPotentialAcMillivoltAc:
		return ((value) * 0.001) 
	case ElectricPotentialAcKilovoltAc:
		return ((value) * 1000.0) 
	case ElectricPotentialAcMegavoltAc:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricPotentialAc) String() string {
	return a.ToString(ElectricPotentialAcVoltAc, 2)
}

// ToString formats the ElectricPotentialAc to string.
// fractionalDigits -1 for not mention
func (a *ElectricPotentialAc) ToString(unit ElectricPotentialAcUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricPotentialAc) getUnitAbbreviation(unit ElectricPotentialAcUnits) string {
	switch unit { 
	case ElectricPotentialAcVoltAc:
		return "Vac" 
	case ElectricPotentialAcMicrovoltAc:
		return "μVac" 
	case ElectricPotentialAcMillivoltAc:
		return "mVac" 
	case ElectricPotentialAcKilovoltAc:
		return "kVac" 
	case ElectricPotentialAcMegavoltAc:
		return "MVac" 
	default:
		return ""
	}
}

// Check if the given ElectricPotentialAc are equals to the current ElectricPotentialAc
func (a *ElectricPotentialAc) Equals(other *ElectricPotentialAc) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricPotentialAc are equals to the current ElectricPotentialAc
func (a *ElectricPotentialAc) CompareTo(other *ElectricPotentialAc) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricPotentialAc to the current ElectricPotentialAc.
func (a *ElectricPotentialAc) Add(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricPotentialAc to the current ElectricPotentialAc.
func (a *ElectricPotentialAc) Subtract(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricPotentialAc to the current ElectricPotentialAc.
func (a *ElectricPotentialAc) Multiply(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value * other.BaseValue()}
}

// Divide the given ElectricPotentialAc to the current ElectricPotentialAc.
func (a *ElectricPotentialAc) Divide(other *ElectricPotentialAc) *ElectricPotentialAc {
	return &ElectricPotentialAc{value: a.value / other.BaseValue()}
}