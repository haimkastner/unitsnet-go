package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricInductanceUnits enumeration
type ElectricInductanceUnits string

const (
    
        // 
        ElectricInductanceHenry ElectricInductanceUnits = "Henry"
        // 
        ElectricInductancePicohenry ElectricInductanceUnits = "Picohenry"
        // 
        ElectricInductanceNanohenry ElectricInductanceUnits = "Nanohenry"
        // 
        ElectricInductanceMicrohenry ElectricInductanceUnits = "Microhenry"
        // 
        ElectricInductanceMillihenry ElectricInductanceUnits = "Millihenry"
)

// ElectricInductanceDto represents an ElectricInductance
type ElectricInductanceDto struct {
	Value float64
	Unit  ElectricInductanceUnits
}

// ElectricInductanceDtoFactory struct to group related functions
type ElectricInductanceDtoFactory struct{}

func (udf ElectricInductanceDtoFactory) FromJSON(data []byte) (*ElectricInductanceDto, error) {
	a := ElectricInductanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricInductanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricInductance struct
type ElectricInductance struct {
	value       float64
    
    henriesLazy *float64 
    picohenriesLazy *float64 
    nanohenriesLazy *float64 
    microhenriesLazy *float64 
    millihenriesLazy *float64 
}

// ElectricInductanceFactory struct to group related functions
type ElectricInductanceFactory struct{}

func (uf ElectricInductanceFactory) CreateElectricInductance(value float64, unit ElectricInductanceUnits) (*ElectricInductance, error) {
	return newElectricInductance(value, unit)
}

func (uf ElectricInductanceFactory) FromDto(dto ElectricInductanceDto) (*ElectricInductance, error) {
	return newElectricInductance(dto.Value, dto.Unit)
}

func (uf ElectricInductanceFactory) FromDtoJSON(data []byte) (*ElectricInductance, error) {
	unitDto, err := ElectricInductanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricInductanceFactory{}.FromDto(*unitDto)
}


// FromHenry creates a new ElectricInductance instance from Henry.
func (uf ElectricInductanceFactory) FromHenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceHenry)
}

// FromPicohenry creates a new ElectricInductance instance from Picohenry.
func (uf ElectricInductanceFactory) FromPicohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductancePicohenry)
}

// FromNanohenry creates a new ElectricInductance instance from Nanohenry.
func (uf ElectricInductanceFactory) FromNanohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceNanohenry)
}

// FromMicrohenry creates a new ElectricInductance instance from Microhenry.
func (uf ElectricInductanceFactory) FromMicrohenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceMicrohenry)
}

// FromMillihenry creates a new ElectricInductance instance from Millihenry.
func (uf ElectricInductanceFactory) FromMillihenries(value float64) (*ElectricInductance, error) {
	return newElectricInductance(value, ElectricInductanceMillihenry)
}




// newElectricInductance creates a new ElectricInductance.
func newElectricInductance(value float64, fromUnit ElectricInductanceUnits) (*ElectricInductance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricInductance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricInductance in Henry.
func (a *ElectricInductance) BaseValue() float64 {
	return a.value
}


// Henry returns the value in Henry.
func (a *ElectricInductance) Henries() float64 {
	if a.henriesLazy != nil {
		return *a.henriesLazy
	}
	henries := a.convertFromBase(ElectricInductanceHenry)
	a.henriesLazy = &henries
	return henries
}

// Picohenry returns the value in Picohenry.
func (a *ElectricInductance) Picohenries() float64 {
	if a.picohenriesLazy != nil {
		return *a.picohenriesLazy
	}
	picohenries := a.convertFromBase(ElectricInductancePicohenry)
	a.picohenriesLazy = &picohenries
	return picohenries
}

// Nanohenry returns the value in Nanohenry.
func (a *ElectricInductance) Nanohenries() float64 {
	if a.nanohenriesLazy != nil {
		return *a.nanohenriesLazy
	}
	nanohenries := a.convertFromBase(ElectricInductanceNanohenry)
	a.nanohenriesLazy = &nanohenries
	return nanohenries
}

// Microhenry returns the value in Microhenry.
func (a *ElectricInductance) Microhenries() float64 {
	if a.microhenriesLazy != nil {
		return *a.microhenriesLazy
	}
	microhenries := a.convertFromBase(ElectricInductanceMicrohenry)
	a.microhenriesLazy = &microhenries
	return microhenries
}

// Millihenry returns the value in Millihenry.
func (a *ElectricInductance) Millihenries() float64 {
	if a.millihenriesLazy != nil {
		return *a.millihenriesLazy
	}
	millihenries := a.convertFromBase(ElectricInductanceMillihenry)
	a.millihenriesLazy = &millihenries
	return millihenries
}


// ToDto creates an ElectricInductanceDto representation.
func (a *ElectricInductance) ToDto(holdInUnit *ElectricInductanceUnits) ElectricInductanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricInductanceHenry // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricInductanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an ElectricInductanceDto representation.
func (a *ElectricInductance) ToDtoJSON(holdInUnit *ElectricInductanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricInductance to a specific unit value.
func (a *ElectricInductance) Convert(toUnit ElectricInductanceUnits) float64 {
	switch toUnit { 
    case ElectricInductanceHenry:
		return a.Henries()
    case ElectricInductancePicohenry:
		return a.Picohenries()
    case ElectricInductanceNanohenry:
		return a.Nanohenries()
    case ElectricInductanceMicrohenry:
		return a.Microhenries()
    case ElectricInductanceMillihenry:
		return a.Millihenries()
	default:
		return 0
	}
}

func (a *ElectricInductance) convertFromBase(toUnit ElectricInductanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricInductanceHenry:
		return (value) 
	case ElectricInductancePicohenry:
		return ((value) / 1e-12) 
	case ElectricInductanceNanohenry:
		return ((value) / 1e-09) 
	case ElectricInductanceMicrohenry:
		return ((value) / 1e-06) 
	case ElectricInductanceMillihenry:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *ElectricInductance) convertToBase(value float64, fromUnit ElectricInductanceUnits) float64 {
	switch fromUnit { 
	case ElectricInductanceHenry:
		return (value) 
	case ElectricInductancePicohenry:
		return ((value) * 1e-12) 
	case ElectricInductanceNanohenry:
		return ((value) * 1e-09) 
	case ElectricInductanceMicrohenry:
		return ((value) * 1e-06) 
	case ElectricInductanceMillihenry:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a ElectricInductance) String() string {
	return a.ToString(ElectricInductanceHenry, 2)
}

// ToString formats the ElectricInductance to string.
// fractionalDigits -1 for not mention
func (a *ElectricInductance) ToString(unit ElectricInductanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricInductance) getUnitAbbreviation(unit ElectricInductanceUnits) string {
	switch unit { 
	case ElectricInductanceHenry:
		return "H" 
	case ElectricInductancePicohenry:
		return "pH" 
	case ElectricInductanceNanohenry:
		return "nH" 
	case ElectricInductanceMicrohenry:
		return "μH" 
	case ElectricInductanceMillihenry:
		return "mH" 
	default:
		return ""
	}
}

// Check if the given ElectricInductance are equals to the current ElectricInductance
func (a *ElectricInductance) Equals(other *ElectricInductance) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricInductance are equals to the current ElectricInductance
func (a *ElectricInductance) CompareTo(other *ElectricInductance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given ElectricInductance to the current ElectricInductance.
func (a *ElectricInductance) Add(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricInductance to the current ElectricInductance.
func (a *ElectricInductance) Subtract(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricInductance to the current ElectricInductance.
func (a *ElectricInductance) Multiply(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value * other.BaseValue()}
}

// Divide the given ElectricInductance to the current ElectricInductance.
func (a *ElectricInductance) Divide(other *ElectricInductance) *ElectricInductance {
	return &ElectricInductance{value: a.value / other.BaseValue()}
}