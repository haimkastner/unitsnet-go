package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricPotentialDcUnits enumeration
type ElectricPotentialDcUnits string

const (
    
        // 
        ElectricPotentialDcVoltDc ElectricPotentialDcUnits = "VoltDc"
        // 
        ElectricPotentialDcMicrovoltDc ElectricPotentialDcUnits = "MicrovoltDc"
        // 
        ElectricPotentialDcMillivoltDc ElectricPotentialDcUnits = "MillivoltDc"
        // 
        ElectricPotentialDcKilovoltDc ElectricPotentialDcUnits = "KilovoltDc"
        // 
        ElectricPotentialDcMegavoltDc ElectricPotentialDcUnits = "MegavoltDc"
)

// ElectricPotentialDcDto represents an ElectricPotentialDc
type ElectricPotentialDcDto struct {
	Value float64
	Unit  ElectricPotentialDcUnits
}

// ElectricPotentialDcDtoFactory struct to group related functions
type ElectricPotentialDcDtoFactory struct{}

func (udf ElectricPotentialDcDtoFactory) FromJSON(data []byte) (*ElectricPotentialDcDto, error) {
	a := ElectricPotentialDcDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricPotentialDcDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricPotentialDc struct
type ElectricPotentialDc struct {
	value       float64
    
    volts_dcLazy *float64 
    microvolts_dcLazy *float64 
    millivolts_dcLazy *float64 
    kilovolts_dcLazy *float64 
    megavolts_dcLazy *float64 
}

// ElectricPotentialDcFactory struct to group related functions
type ElectricPotentialDcFactory struct{}

func (uf ElectricPotentialDcFactory) CreateElectricPotentialDc(value float64, unit ElectricPotentialDcUnits) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, unit)
}

func (uf ElectricPotentialDcFactory) FromDto(dto ElectricPotentialDcDto) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(dto.Value, dto.Unit)
}

func (uf ElectricPotentialDcFactory) FromDtoJSON(data []byte) (*ElectricPotentialDc, error) {
	unitDto, err := ElectricPotentialDcDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricPotentialDcFactory{}.FromDto(*unitDto)
}


// FromVoltDc creates a new ElectricPotentialDc instance from VoltDc.
func (uf ElectricPotentialDcFactory) FromVoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcVoltDc)
}

// FromMicrovoltDc creates a new ElectricPotentialDc instance from MicrovoltDc.
func (uf ElectricPotentialDcFactory) FromMicrovoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMicrovoltDc)
}

// FromMillivoltDc creates a new ElectricPotentialDc instance from MillivoltDc.
func (uf ElectricPotentialDcFactory) FromMillivoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMillivoltDc)
}

// FromKilovoltDc creates a new ElectricPotentialDc instance from KilovoltDc.
func (uf ElectricPotentialDcFactory) FromKilovoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcKilovoltDc)
}

// FromMegavoltDc creates a new ElectricPotentialDc instance from MegavoltDc.
func (uf ElectricPotentialDcFactory) FromMegavoltsDc(value float64) (*ElectricPotentialDc, error) {
	return newElectricPotentialDc(value, ElectricPotentialDcMegavoltDc)
}




// newElectricPotentialDc creates a new ElectricPotentialDc.
func newElectricPotentialDc(value float64, fromUnit ElectricPotentialDcUnits) (*ElectricPotentialDc, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricPotentialDc{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricPotentialDc in VoltDc.
func (a *ElectricPotentialDc) BaseValue() float64 {
	return a.value
}


// VoltDc returns the value in VoltDc.
func (a *ElectricPotentialDc) VoltsDc() float64 {
	if a.volts_dcLazy != nil {
		return *a.volts_dcLazy
	}
	volts_dc := a.convertFromBase(ElectricPotentialDcVoltDc)
	a.volts_dcLazy = &volts_dc
	return volts_dc
}

// MicrovoltDc returns the value in MicrovoltDc.
func (a *ElectricPotentialDc) MicrovoltsDc() float64 {
	if a.microvolts_dcLazy != nil {
		return *a.microvolts_dcLazy
	}
	microvolts_dc := a.convertFromBase(ElectricPotentialDcMicrovoltDc)
	a.microvolts_dcLazy = &microvolts_dc
	return microvolts_dc
}

// MillivoltDc returns the value in MillivoltDc.
func (a *ElectricPotentialDc) MillivoltsDc() float64 {
	if a.millivolts_dcLazy != nil {
		return *a.millivolts_dcLazy
	}
	millivolts_dc := a.convertFromBase(ElectricPotentialDcMillivoltDc)
	a.millivolts_dcLazy = &millivolts_dc
	return millivolts_dc
}

// KilovoltDc returns the value in KilovoltDc.
func (a *ElectricPotentialDc) KilovoltsDc() float64 {
	if a.kilovolts_dcLazy != nil {
		return *a.kilovolts_dcLazy
	}
	kilovolts_dc := a.convertFromBase(ElectricPotentialDcKilovoltDc)
	a.kilovolts_dcLazy = &kilovolts_dc
	return kilovolts_dc
}

// MegavoltDc returns the value in MegavoltDc.
func (a *ElectricPotentialDc) MegavoltsDc() float64 {
	if a.megavolts_dcLazy != nil {
		return *a.megavolts_dcLazy
	}
	megavolts_dc := a.convertFromBase(ElectricPotentialDcMegavoltDc)
	a.megavolts_dcLazy = &megavolts_dc
	return megavolts_dc
}


// ToDto creates an ElectricPotentialDcDto representation.
func (a *ElectricPotentialDc) ToDto(holdInUnit *ElectricPotentialDcUnits) ElectricPotentialDcDto {
	if holdInUnit == nil {
		defaultUnit := ElectricPotentialDcVoltDc // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricPotentialDcDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricPotentialDcDto representation.
func (a *ElectricPotentialDc) ToDtoJSON(holdInUnit *ElectricPotentialDcUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricPotentialDc to a specific unit value.
func (a *ElectricPotentialDc) Convert(toUnit ElectricPotentialDcUnits) float64 {
	switch toUnit { 
    case ElectricPotentialDcVoltDc:
		return a.VoltsDc()
    case ElectricPotentialDcMicrovoltDc:
		return a.MicrovoltsDc()
    case ElectricPotentialDcMillivoltDc:
		return a.MillivoltsDc()
    case ElectricPotentialDcKilovoltDc:
		return a.KilovoltsDc()
    case ElectricPotentialDcMegavoltDc:
		return a.MegavoltsDc()
	default:
		return 0
	}
}

func (a *ElectricPotentialDc) convertFromBase(toUnit ElectricPotentialDcUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricPotentialDcVoltDc:
		return (value) 
	case ElectricPotentialDcMicrovoltDc:
		return ((value) / 1e-06) 
	case ElectricPotentialDcMillivoltDc:
		return ((value) / 0.001) 
	case ElectricPotentialDcKilovoltDc:
		return ((value) / 1000.0) 
	case ElectricPotentialDcMegavoltDc:
		return ((value) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricPotentialDc) convertToBase(value float64, fromUnit ElectricPotentialDcUnits) float64 {
	switch fromUnit { 
	case ElectricPotentialDcVoltDc:
		return (value) 
	case ElectricPotentialDcMicrovoltDc:
		return ((value) * 1e-06) 
	case ElectricPotentialDcMillivoltDc:
		return ((value) * 0.001) 
	case ElectricPotentialDcKilovoltDc:
		return ((value) * 1000.0) 
	case ElectricPotentialDcMegavoltDc:
		return ((value) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricPotentialDc) String() string {
	return a.ToString(ElectricPotentialDcVoltDc, 2)
}

// ToString formats the ElectricPotentialDc to string.
// fractionalDigits -1 for not mention
func (a *ElectricPotentialDc) ToString(unit ElectricPotentialDcUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricPotentialDc) getUnitAbbreviation(unit ElectricPotentialDcUnits) string {
	switch unit { 
	case ElectricPotentialDcVoltDc:
		return "Vdc" 
	case ElectricPotentialDcMicrovoltDc:
		return "μVdc" 
	case ElectricPotentialDcMillivoltDc:
		return "mVdc" 
	case ElectricPotentialDcKilovoltDc:
		return "kVdc" 
	case ElectricPotentialDcMegavoltDc:
		return "MVdc" 
	default:
		return ""
	}
}

// Check if the given ElectricPotentialDc are equals to the current ElectricPotentialDc
func (a *ElectricPotentialDc) Equals(other *ElectricPotentialDc) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricPotentialDc are equals to the current ElectricPotentialDc
func (a *ElectricPotentialDc) CompareTo(other *ElectricPotentialDc) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricPotentialDc to the current ElectricPotentialDc.
func (a *ElectricPotentialDc) Add(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricPotentialDc to the current ElectricPotentialDc.
func (a *ElectricPotentialDc) Subtract(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricPotentialDc to the current ElectricPotentialDc.
func (a *ElectricPotentialDc) Multiply(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value * other.BaseValue()}
}

// Divide the given ElectricPotentialDc to the current ElectricPotentialDc.
func (a *ElectricPotentialDc) Divide(other *ElectricPotentialDc) *ElectricPotentialDc {
	return &ElectricPotentialDc{value: a.value / other.BaseValue()}
}