package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// TemperatureDeltaUnits enumeration
type TemperatureDeltaUnits string

const (
    
        // 
        TemperatureDeltaKelvin TemperatureDeltaUnits = "Kelvin"
        // 
        TemperatureDeltaDegreeCelsius TemperatureDeltaUnits = "DegreeCelsius"
        // 
        TemperatureDeltaDegreeDelisle TemperatureDeltaUnits = "DegreeDelisle"
        // 
        TemperatureDeltaDegreeFahrenheit TemperatureDeltaUnits = "DegreeFahrenheit"
        // 
        TemperatureDeltaDegreeNewton TemperatureDeltaUnits = "DegreeNewton"
        // 
        TemperatureDeltaDegreeRankine TemperatureDeltaUnits = "DegreeRankine"
        // 
        TemperatureDeltaDegreeReaumur TemperatureDeltaUnits = "DegreeReaumur"
        // 
        TemperatureDeltaDegreeRoemer TemperatureDeltaUnits = "DegreeRoemer"
        // 
        TemperatureDeltaMillidegreeCelsius TemperatureDeltaUnits = "MillidegreeCelsius"
)

// TemperatureDeltaDto represents an TemperatureDelta
type TemperatureDeltaDto struct {
	Value float64
	Unit  TemperatureDeltaUnits
}

// TemperatureDeltaDtoFactory struct to group related functions
type TemperatureDeltaDtoFactory struct{}

func (udf TemperatureDeltaDtoFactory) FromJSON(data []byte) (*TemperatureDeltaDto, error) {
	a := TemperatureDeltaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a TemperatureDeltaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// TemperatureDelta struct
type TemperatureDelta struct {
	value       float64
    
    kelvinsLazy *float64 
    degrees_celsiusLazy *float64 
    degrees_delisleLazy *float64 
    degrees_fahrenheitLazy *float64 
    degrees_newtonLazy *float64 
    degrees_rankineLazy *float64 
    degrees_reaumurLazy *float64 
    degrees_roemerLazy *float64 
    millidegrees_celsiusLazy *float64 
}

// TemperatureDeltaFactory struct to group related functions
type TemperatureDeltaFactory struct{}

func (uf TemperatureDeltaFactory) CreateTemperatureDelta(value float64, unit TemperatureDeltaUnits) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, unit)
}

func (uf TemperatureDeltaFactory) FromDto(dto TemperatureDeltaDto) (*TemperatureDelta, error) {
	return newTemperatureDelta(dto.Value, dto.Unit)
}

func (uf TemperatureDeltaFactory) FromDtoJSON(data []byte) (*TemperatureDelta, error) {
	unitDto, err := TemperatureDeltaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return TemperatureDeltaFactory{}.FromDto(*unitDto)
}


// FromKelvin creates a new TemperatureDelta instance from Kelvin.
func (uf TemperatureDeltaFactory) FromKelvins(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaKelvin)
}

// FromDegreeCelsius creates a new TemperatureDelta instance from DegreeCelsius.
func (uf TemperatureDeltaFactory) FromDegreesCelsius(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeCelsius)
}

// FromDegreeDelisle creates a new TemperatureDelta instance from DegreeDelisle.
func (uf TemperatureDeltaFactory) FromDegreesDelisle(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeDelisle)
}

// FromDegreeFahrenheit creates a new TemperatureDelta instance from DegreeFahrenheit.
func (uf TemperatureDeltaFactory) FromDegreesFahrenheit(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeFahrenheit)
}

// FromDegreeNewton creates a new TemperatureDelta instance from DegreeNewton.
func (uf TemperatureDeltaFactory) FromDegreesNewton(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeNewton)
}

// FromDegreeRankine creates a new TemperatureDelta instance from DegreeRankine.
func (uf TemperatureDeltaFactory) FromDegreesRankine(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeRankine)
}

// FromDegreeReaumur creates a new TemperatureDelta instance from DegreeReaumur.
func (uf TemperatureDeltaFactory) FromDegreesReaumur(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeReaumur)
}

// FromDegreeRoemer creates a new TemperatureDelta instance from DegreeRoemer.
func (uf TemperatureDeltaFactory) FromDegreesRoemer(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaDegreeRoemer)
}

// FromMillidegreeCelsius creates a new TemperatureDelta instance from MillidegreeCelsius.
func (uf TemperatureDeltaFactory) FromMillidegreesCelsius(value float64) (*TemperatureDelta, error) {
	return newTemperatureDelta(value, TemperatureDeltaMillidegreeCelsius)
}




// newTemperatureDelta creates a new TemperatureDelta.
func newTemperatureDelta(value float64, fromUnit TemperatureDeltaUnits) (*TemperatureDelta, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &TemperatureDelta{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of TemperatureDelta in Kelvin.
func (a *TemperatureDelta) BaseValue() float64 {
	return a.value
}


// Kelvin returns the value in Kelvin.
func (a *TemperatureDelta) Kelvins() float64 {
	if a.kelvinsLazy != nil {
		return *a.kelvinsLazy
	}
	kelvins := a.convertFromBase(TemperatureDeltaKelvin)
	a.kelvinsLazy = &kelvins
	return kelvins
}

// DegreeCelsius returns the value in DegreeCelsius.
func (a *TemperatureDelta) DegreesCelsius() float64 {
	if a.degrees_celsiusLazy != nil {
		return *a.degrees_celsiusLazy
	}
	degrees_celsius := a.convertFromBase(TemperatureDeltaDegreeCelsius)
	a.degrees_celsiusLazy = &degrees_celsius
	return degrees_celsius
}

// DegreeDelisle returns the value in DegreeDelisle.
func (a *TemperatureDelta) DegreesDelisle() float64 {
	if a.degrees_delisleLazy != nil {
		return *a.degrees_delisleLazy
	}
	degrees_delisle := a.convertFromBase(TemperatureDeltaDegreeDelisle)
	a.degrees_delisleLazy = &degrees_delisle
	return degrees_delisle
}

// DegreeFahrenheit returns the value in DegreeFahrenheit.
func (a *TemperatureDelta) DegreesFahrenheit() float64 {
	if a.degrees_fahrenheitLazy != nil {
		return *a.degrees_fahrenheitLazy
	}
	degrees_fahrenheit := a.convertFromBase(TemperatureDeltaDegreeFahrenheit)
	a.degrees_fahrenheitLazy = &degrees_fahrenheit
	return degrees_fahrenheit
}

// DegreeNewton returns the value in DegreeNewton.
func (a *TemperatureDelta) DegreesNewton() float64 {
	if a.degrees_newtonLazy != nil {
		return *a.degrees_newtonLazy
	}
	degrees_newton := a.convertFromBase(TemperatureDeltaDegreeNewton)
	a.degrees_newtonLazy = &degrees_newton
	return degrees_newton
}

// DegreeRankine returns the value in DegreeRankine.
func (a *TemperatureDelta) DegreesRankine() float64 {
	if a.degrees_rankineLazy != nil {
		return *a.degrees_rankineLazy
	}
	degrees_rankine := a.convertFromBase(TemperatureDeltaDegreeRankine)
	a.degrees_rankineLazy = &degrees_rankine
	return degrees_rankine
}

// DegreeReaumur returns the value in DegreeReaumur.
func (a *TemperatureDelta) DegreesReaumur() float64 {
	if a.degrees_reaumurLazy != nil {
		return *a.degrees_reaumurLazy
	}
	degrees_reaumur := a.convertFromBase(TemperatureDeltaDegreeReaumur)
	a.degrees_reaumurLazy = &degrees_reaumur
	return degrees_reaumur
}

// DegreeRoemer returns the value in DegreeRoemer.
func (a *TemperatureDelta) DegreesRoemer() float64 {
	if a.degrees_roemerLazy != nil {
		return *a.degrees_roemerLazy
	}
	degrees_roemer := a.convertFromBase(TemperatureDeltaDegreeRoemer)
	a.degrees_roemerLazy = &degrees_roemer
	return degrees_roemer
}

// MillidegreeCelsius returns the value in MillidegreeCelsius.
func (a *TemperatureDelta) MillidegreesCelsius() float64 {
	if a.millidegrees_celsiusLazy != nil {
		return *a.millidegrees_celsiusLazy
	}
	millidegrees_celsius := a.convertFromBase(TemperatureDeltaMillidegreeCelsius)
	a.millidegrees_celsiusLazy = &millidegrees_celsius
	return millidegrees_celsius
}


// ToDto creates an TemperatureDeltaDto representation.
func (a *TemperatureDelta) ToDto(holdInUnit *TemperatureDeltaUnits) TemperatureDeltaDto {
	if holdInUnit == nil {
		defaultUnit := TemperatureDeltaKelvin // Default value
		holdInUnit = &defaultUnit
	}

	return TemperatureDeltaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an TemperatureDeltaDto representation.
func (a *TemperatureDelta) ToDtoJSON(holdInUnit *TemperatureDeltaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts TemperatureDelta to a specific unit value.
func (a *TemperatureDelta) Convert(toUnit TemperatureDeltaUnits) float64 {
	switch toUnit { 
    case TemperatureDeltaKelvin:
		return a.Kelvins()
    case TemperatureDeltaDegreeCelsius:
		return a.DegreesCelsius()
    case TemperatureDeltaDegreeDelisle:
		return a.DegreesDelisle()
    case TemperatureDeltaDegreeFahrenheit:
		return a.DegreesFahrenheit()
    case TemperatureDeltaDegreeNewton:
		return a.DegreesNewton()
    case TemperatureDeltaDegreeRankine:
		return a.DegreesRankine()
    case TemperatureDeltaDegreeReaumur:
		return a.DegreesReaumur()
    case TemperatureDeltaDegreeRoemer:
		return a.DegreesRoemer()
    case TemperatureDeltaMillidegreeCelsius:
		return a.MillidegreesCelsius()
	default:
		return 0
	}
}

func (a *TemperatureDelta) convertFromBase(toUnit TemperatureDeltaUnits) float64 {
    value := a.value
	switch toUnit { 
	case TemperatureDeltaKelvin:
		return (value) 
	case TemperatureDeltaDegreeCelsius:
		return (value) 
	case TemperatureDeltaDegreeDelisle:
		return (value * -3 / 2) 
	case TemperatureDeltaDegreeFahrenheit:
		return (value * 9 / 5) 
	case TemperatureDeltaDegreeNewton:
		return (value * 33 / 100) 
	case TemperatureDeltaDegreeRankine:
		return (value * 9 / 5) 
	case TemperatureDeltaDegreeReaumur:
		return (value * 4 / 5) 
	case TemperatureDeltaDegreeRoemer:
		return (value * 21 / 40) 
	case TemperatureDeltaMillidegreeCelsius:
		return ((value) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *TemperatureDelta) convertToBase(value float64, fromUnit TemperatureDeltaUnits) float64 {
	switch fromUnit { 
	case TemperatureDeltaKelvin:
		return (value) 
	case TemperatureDeltaDegreeCelsius:
		return (value) 
	case TemperatureDeltaDegreeDelisle:
		return (value * -2 / 3) 
	case TemperatureDeltaDegreeFahrenheit:
		return (value * 5 / 9) 
	case TemperatureDeltaDegreeNewton:
		return (value * 100 / 33) 
	case TemperatureDeltaDegreeRankine:
		return (value * 5 / 9) 
	case TemperatureDeltaDegreeReaumur:
		return (value * 5 / 4) 
	case TemperatureDeltaDegreeRoemer:
		return (value * 40 / 21) 
	case TemperatureDeltaMillidegreeCelsius:
		return ((value) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a TemperatureDelta) String() string {
	return a.ToString(TemperatureDeltaKelvin, 2)
}

// ToString formats the TemperatureDelta to string.
// fractionalDigits -1 for not mention
func (a *TemperatureDelta) ToString(unit TemperatureDeltaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *TemperatureDelta) getUnitAbbreviation(unit TemperatureDeltaUnits) string {
	switch unit { 
	case TemperatureDeltaKelvin:
		return "∆K" 
	case TemperatureDeltaDegreeCelsius:
		return "∆°C" 
	case TemperatureDeltaDegreeDelisle:
		return "∆°De" 
	case TemperatureDeltaDegreeFahrenheit:
		return "∆°F" 
	case TemperatureDeltaDegreeNewton:
		return "∆°N" 
	case TemperatureDeltaDegreeRankine:
		return "∆°R" 
	case TemperatureDeltaDegreeReaumur:
		return "∆°Ré" 
	case TemperatureDeltaDegreeRoemer:
		return "∆°Rø" 
	case TemperatureDeltaMillidegreeCelsius:
		return "m∆°C" 
	default:
		return ""
	}
}

// Check if the given TemperatureDelta are equals to the current TemperatureDelta
func (a *TemperatureDelta) Equals(other *TemperatureDelta) bool {
	return a.value == other.BaseValue()
}

// Check if the given TemperatureDelta are equals to the current TemperatureDelta
func (a *TemperatureDelta) CompareTo(other *TemperatureDelta) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given TemperatureDelta to the current TemperatureDelta.
func (a *TemperatureDelta) Add(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value + other.BaseValue()}
}

// Subtract the given TemperatureDelta to the current TemperatureDelta.
func (a *TemperatureDelta) Subtract(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value - other.BaseValue()}
}

// Multiply the given TemperatureDelta to the current TemperatureDelta.
func (a *TemperatureDelta) Multiply(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value * other.BaseValue()}
}

// Divide the given TemperatureDelta to the current TemperatureDelta.
func (a *TemperatureDelta) Divide(other *TemperatureDelta) *TemperatureDelta {
	return &TemperatureDelta{value: a.value / other.BaseValue()}
}