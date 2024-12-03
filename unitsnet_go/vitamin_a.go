package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// VitaminAUnits enumeration
type VitaminAUnits string

const (
    
        // 
        VitaminAInternationalUnit VitaminAUnits = "InternationalUnit"
)

// VitaminADto represents an VitaminA
type VitaminADto struct {
	Value float64
	Unit  VitaminAUnits
}

// VitaminADtoFactory struct to group related functions
type VitaminADtoFactory struct{}

func (udf VitaminADtoFactory) FromJSON(data []byte) (*VitaminADto, error) {
	a := VitaminADto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a VitaminADto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// VitaminA struct
type VitaminA struct {
	value       float64
    
    international_unitsLazy *float64 
}

// VitaminAFactory struct to group related functions
type VitaminAFactory struct{}

func (uf VitaminAFactory) CreateVitaminA(value float64, unit VitaminAUnits) (*VitaminA, error) {
	return newVitaminA(value, unit)
}

func (uf VitaminAFactory) FromDto(dto VitaminADto) (*VitaminA, error) {
	return newVitaminA(dto.Value, dto.Unit)
}

func (uf VitaminAFactory) FromDtoJSON(data []byte) (*VitaminA, error) {
	unitDto, err := VitaminADtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return VitaminAFactory{}.FromDto(*unitDto)
}


// FromInternationalUnit creates a new VitaminA instance from InternationalUnit.
func (uf VitaminAFactory) FromInternationalUnits(value float64) (*VitaminA, error) {
	return newVitaminA(value, VitaminAInternationalUnit)
}




// newVitaminA creates a new VitaminA.
func newVitaminA(value float64, fromUnit VitaminAUnits) (*VitaminA, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &VitaminA{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of VitaminA in InternationalUnit.
func (a *VitaminA) BaseValue() float64 {
	return a.value
}


// InternationalUnit returns the value in InternationalUnit.
func (a *VitaminA) InternationalUnits() float64 {
	if a.international_unitsLazy != nil {
		return *a.international_unitsLazy
	}
	international_units := a.convertFromBase(VitaminAInternationalUnit)
	a.international_unitsLazy = &international_units
	return international_units
}


// ToDto creates an VitaminADto representation.
func (a *VitaminA) ToDto(holdInUnit *VitaminAUnits) VitaminADto {
	if holdInUnit == nil {
		defaultUnit := VitaminAInternationalUnit // Default value
		holdInUnit = &defaultUnit
	}

	return VitaminADto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an VitaminADto representation.
func (a *VitaminA) ToDtoJSON(holdInUnit *VitaminAUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts VitaminA to a specific unit value.
func (a *VitaminA) Convert(toUnit VitaminAUnits) float64 {
	switch toUnit { 
    case VitaminAInternationalUnit:
		return a.InternationalUnits()
	default:
		return 0
	}
}

func (a *VitaminA) convertFromBase(toUnit VitaminAUnits) float64 {
    value := a.value
	switch toUnit { 
	case VitaminAInternationalUnit:
		return (value) 
	default:
		return math.NaN()
	}
}

func (a *VitaminA) convertToBase(value float64, fromUnit VitaminAUnits) float64 {
	switch fromUnit { 
	case VitaminAInternationalUnit:
		return (value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a VitaminA) String() string {
	return a.ToString(VitaminAInternationalUnit, 2)
}

// ToString formats the VitaminA to string.
// fractionalDigits -1 for not mention
func (a *VitaminA) ToString(unit VitaminAUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *VitaminA) getUnitAbbreviation(unit VitaminAUnits) string {
	switch unit { 
	case VitaminAInternationalUnit:
		return "IU" 
	default:
		return ""
	}
}

// Check if the given VitaminA are equals to the current VitaminA
func (a *VitaminA) Equals(other *VitaminA) bool {
	return a.value == other.BaseValue()
}

// Check if the given VitaminA are equals to the current VitaminA
func (a *VitaminA) CompareTo(other *VitaminA) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given VitaminA to the current VitaminA.
func (a *VitaminA) Add(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value + other.BaseValue()}
}

// Subtract the given VitaminA to the current VitaminA.
func (a *VitaminA) Subtract(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value - other.BaseValue()}
}

// Multiply the given VitaminA to the current VitaminA.
func (a *VitaminA) Multiply(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value * other.BaseValue()}
}

// Divide the given VitaminA to the current VitaminA.
func (a *VitaminA) Divide(other *VitaminA) *VitaminA {
	return &VitaminA{value: a.value / other.BaseValue()}
}