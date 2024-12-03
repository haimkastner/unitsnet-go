package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricConductivityUnits enumeration
type ElectricConductivityUnits string

const (
    
        // 
        ElectricConductivitySiemensPerMeter ElectricConductivityUnits = "SiemensPerMeter"
        // 
        ElectricConductivitySiemensPerInch ElectricConductivityUnits = "SiemensPerInch"
        // 
        ElectricConductivitySiemensPerFoot ElectricConductivityUnits = "SiemensPerFoot"
        // 
        ElectricConductivitySiemensPerCentimeter ElectricConductivityUnits = "SiemensPerCentimeter"
        // 
        ElectricConductivityMicrosiemensPerCentimeter ElectricConductivityUnits = "MicrosiemensPerCentimeter"
        // 
        ElectricConductivityMillisiemensPerCentimeter ElectricConductivityUnits = "MillisiemensPerCentimeter"
)

// ElectricConductivityDto represents an ElectricConductivity
type ElectricConductivityDto struct {
	Value float64
	Unit  ElectricConductivityUnits
}

// ElectricConductivityDtoFactory struct to group related functions
type ElectricConductivityDtoFactory struct{}

func (udf ElectricConductivityDtoFactory) FromJSON(data []byte) (*ElectricConductivityDto, error) {
	a := ElectricConductivityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricConductivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricConductivity struct
type ElectricConductivity struct {
	value       float64
    
    siemens_per_meterLazy *float64 
    siemens_per_inchLazy *float64 
    siemens_per_footLazy *float64 
    siemens_per_centimeterLazy *float64 
    microsiemens_per_centimeterLazy *float64 
    millisiemens_per_centimeterLazy *float64 
}

// ElectricConductivityFactory struct to group related functions
type ElectricConductivityFactory struct{}

func (uf ElectricConductivityFactory) CreateElectricConductivity(value float64, unit ElectricConductivityUnits) (*ElectricConductivity, error) {
	return newElectricConductivity(value, unit)
}

func (uf ElectricConductivityFactory) FromDto(dto ElectricConductivityDto) (*ElectricConductivity, error) {
	return newElectricConductivity(dto.Value, dto.Unit)
}

func (uf ElectricConductivityFactory) FromDtoJSON(data []byte) (*ElectricConductivity, error) {
	unitDto, err := ElectricConductivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricConductivityFactory{}.FromDto(*unitDto)
}


// FromSiemensPerMeter creates a new ElectricConductivity instance from SiemensPerMeter.
func (uf ElectricConductivityFactory) FromSiemensPerMeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerMeter)
}

// FromSiemensPerInch creates a new ElectricConductivity instance from SiemensPerInch.
func (uf ElectricConductivityFactory) FromSiemensPerInch(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerInch)
}

// FromSiemensPerFoot creates a new ElectricConductivity instance from SiemensPerFoot.
func (uf ElectricConductivityFactory) FromSiemensPerFoot(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerFoot)
}

// FromSiemensPerCentimeter creates a new ElectricConductivity instance from SiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromSiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivitySiemensPerCentimeter)
}

// FromMicrosiemensPerCentimeter creates a new ElectricConductivity instance from MicrosiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromMicrosiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivityMicrosiemensPerCentimeter)
}

// FromMillisiemensPerCentimeter creates a new ElectricConductivity instance from MillisiemensPerCentimeter.
func (uf ElectricConductivityFactory) FromMillisiemensPerCentimeter(value float64) (*ElectricConductivity, error) {
	return newElectricConductivity(value, ElectricConductivityMillisiemensPerCentimeter)
}




// newElectricConductivity creates a new ElectricConductivity.
func newElectricConductivity(value float64, fromUnit ElectricConductivityUnits) (*ElectricConductivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricConductivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricConductivity in SiemensPerMeter.
func (a *ElectricConductivity) BaseValue() float64 {
	return a.value
}


// SiemensPerMeter returns the value in SiemensPerMeter.
func (a *ElectricConductivity) SiemensPerMeter() float64 {
	if a.siemens_per_meterLazy != nil {
		return *a.siemens_per_meterLazy
	}
	siemens_per_meter := a.convertFromBase(ElectricConductivitySiemensPerMeter)
	a.siemens_per_meterLazy = &siemens_per_meter
	return siemens_per_meter
}

// SiemensPerInch returns the value in SiemensPerInch.
func (a *ElectricConductivity) SiemensPerInch() float64 {
	if a.siemens_per_inchLazy != nil {
		return *a.siemens_per_inchLazy
	}
	siemens_per_inch := a.convertFromBase(ElectricConductivitySiemensPerInch)
	a.siemens_per_inchLazy = &siemens_per_inch
	return siemens_per_inch
}

// SiemensPerFoot returns the value in SiemensPerFoot.
func (a *ElectricConductivity) SiemensPerFoot() float64 {
	if a.siemens_per_footLazy != nil {
		return *a.siemens_per_footLazy
	}
	siemens_per_foot := a.convertFromBase(ElectricConductivitySiemensPerFoot)
	a.siemens_per_footLazy = &siemens_per_foot
	return siemens_per_foot
}

// SiemensPerCentimeter returns the value in SiemensPerCentimeter.
func (a *ElectricConductivity) SiemensPerCentimeter() float64 {
	if a.siemens_per_centimeterLazy != nil {
		return *a.siemens_per_centimeterLazy
	}
	siemens_per_centimeter := a.convertFromBase(ElectricConductivitySiemensPerCentimeter)
	a.siemens_per_centimeterLazy = &siemens_per_centimeter
	return siemens_per_centimeter
}

// MicrosiemensPerCentimeter returns the value in MicrosiemensPerCentimeter.
func (a *ElectricConductivity) MicrosiemensPerCentimeter() float64 {
	if a.microsiemens_per_centimeterLazy != nil {
		return *a.microsiemens_per_centimeterLazy
	}
	microsiemens_per_centimeter := a.convertFromBase(ElectricConductivityMicrosiemensPerCentimeter)
	a.microsiemens_per_centimeterLazy = &microsiemens_per_centimeter
	return microsiemens_per_centimeter
}

// MillisiemensPerCentimeter returns the value in MillisiemensPerCentimeter.
func (a *ElectricConductivity) MillisiemensPerCentimeter() float64 {
	if a.millisiemens_per_centimeterLazy != nil {
		return *a.millisiemens_per_centimeterLazy
	}
	millisiemens_per_centimeter := a.convertFromBase(ElectricConductivityMillisiemensPerCentimeter)
	a.millisiemens_per_centimeterLazy = &millisiemens_per_centimeter
	return millisiemens_per_centimeter
}


// ToDto creates an ElectricConductivityDto representation.
func (a *ElectricConductivity) ToDto(holdInUnit *ElectricConductivityUnits) ElectricConductivityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricConductivitySiemensPerMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricConductivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricConductivityDto representation.
func (a *ElectricConductivity) ToDtoJSON(holdInUnit *ElectricConductivityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricConductivity to a specific unit value.
func (a *ElectricConductivity) Convert(toUnit ElectricConductivityUnits) float64 {
	switch toUnit { 
    case ElectricConductivitySiemensPerMeter:
		return a.SiemensPerMeter()
    case ElectricConductivitySiemensPerInch:
		return a.SiemensPerInch()
    case ElectricConductivitySiemensPerFoot:
		return a.SiemensPerFoot()
    case ElectricConductivitySiemensPerCentimeter:
		return a.SiemensPerCentimeter()
    case ElectricConductivityMicrosiemensPerCentimeter:
		return a.MicrosiemensPerCentimeter()
    case ElectricConductivityMillisiemensPerCentimeter:
		return a.MillisiemensPerCentimeter()
	default:
		return 0
	}
}

func (a *ElectricConductivity) convertFromBase(toUnit ElectricConductivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricConductivitySiemensPerMeter:
		return (value) 
	case ElectricConductivitySiemensPerInch:
		return (value / 3.937007874015748e1) 
	case ElectricConductivitySiemensPerFoot:
		return (value / 3.2808398950131234) 
	case ElectricConductivitySiemensPerCentimeter:
		return (value / 1e2) 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return ((value / 1e2) / 1e-06) 
	case ElectricConductivityMillisiemensPerCentimeter:
		return ((value / 1e2) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricConductivity) convertToBase(value float64, fromUnit ElectricConductivityUnits) float64 {
	switch fromUnit { 
	case ElectricConductivitySiemensPerMeter:
		return (value) 
	case ElectricConductivitySiemensPerInch:
		return (value * 3.937007874015748e1) 
	case ElectricConductivitySiemensPerFoot:
		return (value * 3.2808398950131234) 
	case ElectricConductivitySiemensPerCentimeter:
		return (value * 1e2) 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return ((value * 1e2) * 1e-06) 
	case ElectricConductivityMillisiemensPerCentimeter:
		return ((value * 1e2) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricConductivity) String() string {
	return a.ToString(ElectricConductivitySiemensPerMeter, 2)
}

// ToString formats the ElectricConductivity to string.
// fractionalDigits -1 for not mention
func (a *ElectricConductivity) ToString(unit ElectricConductivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricConductivity) getUnitAbbreviation(unit ElectricConductivityUnits) string {
	switch unit { 
	case ElectricConductivitySiemensPerMeter:
		return "S/m" 
	case ElectricConductivitySiemensPerInch:
		return "S/in" 
	case ElectricConductivitySiemensPerFoot:
		return "S/ft" 
	case ElectricConductivitySiemensPerCentimeter:
		return "S/cm" 
	case ElectricConductivityMicrosiemensPerCentimeter:
		return "μS/cm" 
	case ElectricConductivityMillisiemensPerCentimeter:
		return "mS/cm" 
	default:
		return ""
	}
}

// Check if the given ElectricConductivity are equals to the current ElectricConductivity
func (a *ElectricConductivity) Equals(other *ElectricConductivity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricConductivity are equals to the current ElectricConductivity
func (a *ElectricConductivity) CompareTo(other *ElectricConductivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricConductivity to the current ElectricConductivity.
func (a *ElectricConductivity) Add(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricConductivity to the current ElectricConductivity.
func (a *ElectricConductivity) Subtract(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricConductivity to the current ElectricConductivity.
func (a *ElectricConductivity) Multiply(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value * other.BaseValue()}
}

// Divide the given ElectricConductivity to the current ElectricConductivity.
func (a *ElectricConductivity) Divide(other *ElectricConductivity) *ElectricConductivity {
	return &ElectricConductivity{value: a.value / other.BaseValue()}
}