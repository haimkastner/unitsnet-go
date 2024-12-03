package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricResistanceUnits enumeration
type ElectricResistanceUnits string

const (
    
        // 
        ElectricResistanceOhm ElectricResistanceUnits = "Ohm"
        // 
        ElectricResistanceMicroohm ElectricResistanceUnits = "Microohm"
        // 
        ElectricResistanceMilliohm ElectricResistanceUnits = "Milliohm"
        // 
        ElectricResistanceKiloohm ElectricResistanceUnits = "Kiloohm"
        // 
        ElectricResistanceMegaohm ElectricResistanceUnits = "Megaohm"
        // 
        ElectricResistanceGigaohm ElectricResistanceUnits = "Gigaohm"
        // 
        ElectricResistanceTeraohm ElectricResistanceUnits = "Teraohm"
)

// ElectricResistanceDto represents an ElectricResistance
type ElectricResistanceDto struct {
	Value float64
	Unit  ElectricResistanceUnits
}

// ElectricResistanceDtoFactory struct to group related functions
type ElectricResistanceDtoFactory struct{}

func (udf ElectricResistanceDtoFactory) FromJSON(data []byte) (*ElectricResistanceDto, error) {
	a := ElectricResistanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricResistanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricResistance struct
type ElectricResistance struct {
	value       float64
    
    ohmsLazy *float64 
    microohmsLazy *float64 
    milliohmsLazy *float64 
    kiloohmsLazy *float64 
    megaohmsLazy *float64 
    gigaohmsLazy *float64 
    teraohmsLazy *float64 
}

// ElectricResistanceFactory struct to group related functions
type ElectricResistanceFactory struct{}

func (uf ElectricResistanceFactory) CreateElectricResistance(value float64, unit ElectricResistanceUnits) (*ElectricResistance, error) {
	return newElectricResistance(value, unit)
}

func (uf ElectricResistanceFactory) FromDto(dto ElectricResistanceDto) (*ElectricResistance, error) {
	return newElectricResistance(dto.Value, dto.Unit)
}

func (uf ElectricResistanceFactory) FromDtoJSON(data []byte) (*ElectricResistance, error) {
	unitDto, err := ElectricResistanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricResistanceFactory{}.FromDto(*unitDto)
}


// FromOhm creates a new ElectricResistance instance from Ohm.
func (uf ElectricResistanceFactory) FromOhms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceOhm)
}

// FromMicroohm creates a new ElectricResistance instance from Microohm.
func (uf ElectricResistanceFactory) FromMicroohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMicroohm)
}

// FromMilliohm creates a new ElectricResistance instance from Milliohm.
func (uf ElectricResistanceFactory) FromMilliohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMilliohm)
}

// FromKiloohm creates a new ElectricResistance instance from Kiloohm.
func (uf ElectricResistanceFactory) FromKiloohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceKiloohm)
}

// FromMegaohm creates a new ElectricResistance instance from Megaohm.
func (uf ElectricResistanceFactory) FromMegaohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceMegaohm)
}

// FromGigaohm creates a new ElectricResistance instance from Gigaohm.
func (uf ElectricResistanceFactory) FromGigaohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceGigaohm)
}

// FromTeraohm creates a new ElectricResistance instance from Teraohm.
func (uf ElectricResistanceFactory) FromTeraohms(value float64) (*ElectricResistance, error) {
	return newElectricResistance(value, ElectricResistanceTeraohm)
}




// newElectricResistance creates a new ElectricResistance.
func newElectricResistance(value float64, fromUnit ElectricResistanceUnits) (*ElectricResistance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricResistance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricResistance in Ohm.
func (a *ElectricResistance) BaseValue() float64 {
	return a.value
}


// Ohm returns the value in Ohm.
func (a *ElectricResistance) Ohms() float64 {
	if a.ohmsLazy != nil {
		return *a.ohmsLazy
	}
	ohms := a.convertFromBase(ElectricResistanceOhm)
	a.ohmsLazy = &ohms
	return ohms
}

// Microohm returns the value in Microohm.
func (a *ElectricResistance) Microohms() float64 {
	if a.microohmsLazy != nil {
		return *a.microohmsLazy
	}
	microohms := a.convertFromBase(ElectricResistanceMicroohm)
	a.microohmsLazy = &microohms
	return microohms
}

// Milliohm returns the value in Milliohm.
func (a *ElectricResistance) Milliohms() float64 {
	if a.milliohmsLazy != nil {
		return *a.milliohmsLazy
	}
	milliohms := a.convertFromBase(ElectricResistanceMilliohm)
	a.milliohmsLazy = &milliohms
	return milliohms
}

// Kiloohm returns the value in Kiloohm.
func (a *ElectricResistance) Kiloohms() float64 {
	if a.kiloohmsLazy != nil {
		return *a.kiloohmsLazy
	}
	kiloohms := a.convertFromBase(ElectricResistanceKiloohm)
	a.kiloohmsLazy = &kiloohms
	return kiloohms
}

// Megaohm returns the value in Megaohm.
func (a *ElectricResistance) Megaohms() float64 {
	if a.megaohmsLazy != nil {
		return *a.megaohmsLazy
	}
	megaohms := a.convertFromBase(ElectricResistanceMegaohm)
	a.megaohmsLazy = &megaohms
	return megaohms
}

// Gigaohm returns the value in Gigaohm.
func (a *ElectricResistance) Gigaohms() float64 {
	if a.gigaohmsLazy != nil {
		return *a.gigaohmsLazy
	}
	gigaohms := a.convertFromBase(ElectricResistanceGigaohm)
	a.gigaohmsLazy = &gigaohms
	return gigaohms
}

// Teraohm returns the value in Teraohm.
func (a *ElectricResistance) Teraohms() float64 {
	if a.teraohmsLazy != nil {
		return *a.teraohmsLazy
	}
	teraohms := a.convertFromBase(ElectricResistanceTeraohm)
	a.teraohmsLazy = &teraohms
	return teraohms
}


// ToDto creates an ElectricResistanceDto representation.
func (a *ElectricResistance) ToDto(holdInUnit *ElectricResistanceUnits) ElectricResistanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricResistanceOhm // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricResistanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricResistanceDto representation.
func (a *ElectricResistance) ToDtoJSON(holdInUnit *ElectricResistanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricResistance to a specific unit value.
func (a *ElectricResistance) Convert(toUnit ElectricResistanceUnits) float64 {
	switch toUnit { 
    case ElectricResistanceOhm:
		return a.Ohms()
    case ElectricResistanceMicroohm:
		return a.Microohms()
    case ElectricResistanceMilliohm:
		return a.Milliohms()
    case ElectricResistanceKiloohm:
		return a.Kiloohms()
    case ElectricResistanceMegaohm:
		return a.Megaohms()
    case ElectricResistanceGigaohm:
		return a.Gigaohms()
    case ElectricResistanceTeraohm:
		return a.Teraohms()
	default:
		return 0
	}
}

func (a *ElectricResistance) convertFromBase(toUnit ElectricResistanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricResistanceOhm:
		return (value) 
	case ElectricResistanceMicroohm:
		return ((value) / 1e-06) 
	case ElectricResistanceMilliohm:
		return ((value) / 0.001) 
	case ElectricResistanceKiloohm:
		return ((value) / 1000.0) 
	case ElectricResistanceMegaohm:
		return ((value) / 1000000.0) 
	case ElectricResistanceGigaohm:
		return ((value) / 1000000000.0) 
	case ElectricResistanceTeraohm:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricResistance) convertToBase(value float64, fromUnit ElectricResistanceUnits) float64 {
	switch fromUnit { 
	case ElectricResistanceOhm:
		return (value) 
	case ElectricResistanceMicroohm:
		return ((value) * 1e-06) 
	case ElectricResistanceMilliohm:
		return ((value) * 0.001) 
	case ElectricResistanceKiloohm:
		return ((value) * 1000.0) 
	case ElectricResistanceMegaohm:
		return ((value) * 1000000.0) 
	case ElectricResistanceGigaohm:
		return ((value) * 1000000000.0) 
	case ElectricResistanceTeraohm:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricResistance) String() string {
	return a.ToString(ElectricResistanceOhm, 2)
}

// ToString formats the ElectricResistance to string.
// fractionalDigits -1 for not mention
func (a *ElectricResistance) ToString(unit ElectricResistanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricResistance) getUnitAbbreviation(unit ElectricResistanceUnits) string {
	switch unit { 
	case ElectricResistanceOhm:
		return "Ω" 
	case ElectricResistanceMicroohm:
		return "μΩ" 
	case ElectricResistanceMilliohm:
		return "mΩ" 
	case ElectricResistanceKiloohm:
		return "kΩ" 
	case ElectricResistanceMegaohm:
		return "MΩ" 
	case ElectricResistanceGigaohm:
		return "GΩ" 
	case ElectricResistanceTeraohm:
		return "TΩ" 
	default:
		return ""
	}
}

// Check if the given ElectricResistance are equals to the current ElectricResistance
func (a *ElectricResistance) Equals(other *ElectricResistance) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricResistance are equals to the current ElectricResistance
func (a *ElectricResistance) CompareTo(other *ElectricResistance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricResistance to the current ElectricResistance.
func (a *ElectricResistance) Add(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricResistance to the current ElectricResistance.
func (a *ElectricResistance) Subtract(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricResistance to the current ElectricResistance.
func (a *ElectricResistance) Multiply(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value * other.BaseValue()}
}

// Divide the given ElectricResistance to the current ElectricResistance.
func (a *ElectricResistance) Divide(other *ElectricResistance) *ElectricResistance {
	return &ElectricResistance{value: a.value / other.BaseValue()}
}