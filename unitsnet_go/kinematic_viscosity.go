package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// KinematicViscosityUnits enumeration
type KinematicViscosityUnits string

const (
    
        // 
        KinematicViscositySquareMeterPerSecond KinematicViscosityUnits = "SquareMeterPerSecond"
        // 
        KinematicViscosityStokes KinematicViscosityUnits = "Stokes"
        // 
        KinematicViscositySquareFootPerSecond KinematicViscosityUnits = "SquareFootPerSecond"
        // 
        KinematicViscosityNanostokes KinematicViscosityUnits = "Nanostokes"
        // 
        KinematicViscosityMicrostokes KinematicViscosityUnits = "Microstokes"
        // 
        KinematicViscosityMillistokes KinematicViscosityUnits = "Millistokes"
        // 
        KinematicViscosityCentistokes KinematicViscosityUnits = "Centistokes"
        // 
        KinematicViscosityDecistokes KinematicViscosityUnits = "Decistokes"
        // 
        KinematicViscosityKilostokes KinematicViscosityUnits = "Kilostokes"
)

// KinematicViscosityDto represents an KinematicViscosity
type KinematicViscosityDto struct {
	Value float64
	Unit  KinematicViscosityUnits
}

// KinematicViscosityDtoFactory struct to group related functions
type KinematicViscosityDtoFactory struct{}

func (udf KinematicViscosityDtoFactory) FromJSON(data []byte) (*KinematicViscosityDto, error) {
	a := KinematicViscosityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a KinematicViscosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// KinematicViscosity struct
type KinematicViscosity struct {
	value       float64
    
    square_meters_per_secondLazy *float64 
    stokesLazy *float64 
    square_feet_per_secondLazy *float64 
    nanostokesLazy *float64 
    microstokesLazy *float64 
    millistokesLazy *float64 
    centistokesLazy *float64 
    decistokesLazy *float64 
    kilostokesLazy *float64 
}

// KinematicViscosityFactory struct to group related functions
type KinematicViscosityFactory struct{}

func (uf KinematicViscosityFactory) CreateKinematicViscosity(value float64, unit KinematicViscosityUnits) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, unit)
}

func (uf KinematicViscosityFactory) FromDto(dto KinematicViscosityDto) (*KinematicViscosity, error) {
	return newKinematicViscosity(dto.Value, dto.Unit)
}

func (uf KinematicViscosityFactory) FromDtoJSON(data []byte) (*KinematicViscosity, error) {
	unitDto, err := KinematicViscosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return KinematicViscosityFactory{}.FromDto(*unitDto)
}


// FromSquareMeterPerSecond creates a new KinematicViscosity instance from SquareMeterPerSecond.
func (uf KinematicViscosityFactory) FromSquareMetersPerSecond(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscositySquareMeterPerSecond)
}

// FromStokes creates a new KinematicViscosity instance from Stokes.
func (uf KinematicViscosityFactory) FromStokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityStokes)
}

// FromSquareFootPerSecond creates a new KinematicViscosity instance from SquareFootPerSecond.
func (uf KinematicViscosityFactory) FromSquareFeetPerSecond(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscositySquareFootPerSecond)
}

// FromNanostokes creates a new KinematicViscosity instance from Nanostokes.
func (uf KinematicViscosityFactory) FromNanostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityNanostokes)
}

// FromMicrostokes creates a new KinematicViscosity instance from Microstokes.
func (uf KinematicViscosityFactory) FromMicrostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityMicrostokes)
}

// FromMillistokes creates a new KinematicViscosity instance from Millistokes.
func (uf KinematicViscosityFactory) FromMillistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityMillistokes)
}

// FromCentistokes creates a new KinematicViscosity instance from Centistokes.
func (uf KinematicViscosityFactory) FromCentistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityCentistokes)
}

// FromDecistokes creates a new KinematicViscosity instance from Decistokes.
func (uf KinematicViscosityFactory) FromDecistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityDecistokes)
}

// FromKilostokes creates a new KinematicViscosity instance from Kilostokes.
func (uf KinematicViscosityFactory) FromKilostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityKilostokes)
}




// newKinematicViscosity creates a new KinematicViscosity.
func newKinematicViscosity(value float64, fromUnit KinematicViscosityUnits) (*KinematicViscosity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &KinematicViscosity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of KinematicViscosity in SquareMeterPerSecond.
func (a *KinematicViscosity) BaseValue() float64 {
	return a.value
}


// SquareMeterPerSecond returns the value in SquareMeterPerSecond.
func (a *KinematicViscosity) SquareMetersPerSecond() float64 {
	if a.square_meters_per_secondLazy != nil {
		return *a.square_meters_per_secondLazy
	}
	square_meters_per_second := a.convertFromBase(KinematicViscositySquareMeterPerSecond)
	a.square_meters_per_secondLazy = &square_meters_per_second
	return square_meters_per_second
}

// Stokes returns the value in Stokes.
func (a *KinematicViscosity) Stokes() float64 {
	if a.stokesLazy != nil {
		return *a.stokesLazy
	}
	stokes := a.convertFromBase(KinematicViscosityStokes)
	a.stokesLazy = &stokes
	return stokes
}

// SquareFootPerSecond returns the value in SquareFootPerSecond.
func (a *KinematicViscosity) SquareFeetPerSecond() float64 {
	if a.square_feet_per_secondLazy != nil {
		return *a.square_feet_per_secondLazy
	}
	square_feet_per_second := a.convertFromBase(KinematicViscositySquareFootPerSecond)
	a.square_feet_per_secondLazy = &square_feet_per_second
	return square_feet_per_second
}

// Nanostokes returns the value in Nanostokes.
func (a *KinematicViscosity) Nanostokes() float64 {
	if a.nanostokesLazy != nil {
		return *a.nanostokesLazy
	}
	nanostokes := a.convertFromBase(KinematicViscosityNanostokes)
	a.nanostokesLazy = &nanostokes
	return nanostokes
}

// Microstokes returns the value in Microstokes.
func (a *KinematicViscosity) Microstokes() float64 {
	if a.microstokesLazy != nil {
		return *a.microstokesLazy
	}
	microstokes := a.convertFromBase(KinematicViscosityMicrostokes)
	a.microstokesLazy = &microstokes
	return microstokes
}

// Millistokes returns the value in Millistokes.
func (a *KinematicViscosity) Millistokes() float64 {
	if a.millistokesLazy != nil {
		return *a.millistokesLazy
	}
	millistokes := a.convertFromBase(KinematicViscosityMillistokes)
	a.millistokesLazy = &millistokes
	return millistokes
}

// Centistokes returns the value in Centistokes.
func (a *KinematicViscosity) Centistokes() float64 {
	if a.centistokesLazy != nil {
		return *a.centistokesLazy
	}
	centistokes := a.convertFromBase(KinematicViscosityCentistokes)
	a.centistokesLazy = &centistokes
	return centistokes
}

// Decistokes returns the value in Decistokes.
func (a *KinematicViscosity) Decistokes() float64 {
	if a.decistokesLazy != nil {
		return *a.decistokesLazy
	}
	decistokes := a.convertFromBase(KinematicViscosityDecistokes)
	a.decistokesLazy = &decistokes
	return decistokes
}

// Kilostokes returns the value in Kilostokes.
func (a *KinematicViscosity) Kilostokes() float64 {
	if a.kilostokesLazy != nil {
		return *a.kilostokesLazy
	}
	kilostokes := a.convertFromBase(KinematicViscosityKilostokes)
	a.kilostokesLazy = &kilostokes
	return kilostokes
}


// ToDto creates an KinematicViscosityDto representation.
func (a *KinematicViscosity) ToDto(holdInUnit *KinematicViscosityUnits) KinematicViscosityDto {
	if holdInUnit == nil {
		defaultUnit := KinematicViscositySquareMeterPerSecond // Default value
		holdInUnit = &defaultUnit
	}

	return KinematicViscosityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an KinematicViscosityDto representation.
func (a *KinematicViscosity) ToDtoJSON(holdInUnit *KinematicViscosityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts KinematicViscosity to a specific unit value.
func (a *KinematicViscosity) Convert(toUnit KinematicViscosityUnits) float64 {
	switch toUnit { 
    case KinematicViscositySquareMeterPerSecond:
		return a.SquareMetersPerSecond()
    case KinematicViscosityStokes:
		return a.Stokes()
    case KinematicViscositySquareFootPerSecond:
		return a.SquareFeetPerSecond()
    case KinematicViscosityNanostokes:
		return a.Nanostokes()
    case KinematicViscosityMicrostokes:
		return a.Microstokes()
    case KinematicViscosityMillistokes:
		return a.Millistokes()
    case KinematicViscosityCentistokes:
		return a.Centistokes()
    case KinematicViscosityDecistokes:
		return a.Decistokes()
    case KinematicViscosityKilostokes:
		return a.Kilostokes()
	default:
		return 0
	}
}

func (a *KinematicViscosity) convertFromBase(toUnit KinematicViscosityUnits) float64 {
    value := a.value
	switch toUnit { 
	case KinematicViscositySquareMeterPerSecond:
		return (value) 
	case KinematicViscosityStokes:
		return (value * 1e4) 
	case KinematicViscositySquareFootPerSecond:
		return (value * 10.7639) 
	case KinematicViscosityNanostokes:
		return ((value * 1e4) / 1e-09) 
	case KinematicViscosityMicrostokes:
		return ((value * 1e4) / 1e-06) 
	case KinematicViscosityMillistokes:
		return ((value * 1e4) / 0.001) 
	case KinematicViscosityCentistokes:
		return ((value * 1e4) / 0.01) 
	case KinematicViscosityDecistokes:
		return ((value * 1e4) / 0.1) 
	case KinematicViscosityKilostokes:
		return ((value * 1e4) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *KinematicViscosity) convertToBase(value float64, fromUnit KinematicViscosityUnits) float64 {
	switch fromUnit { 
	case KinematicViscositySquareMeterPerSecond:
		return (value) 
	case KinematicViscosityStokes:
		return (value / 1e4) 
	case KinematicViscositySquareFootPerSecond:
		return (value / 10.7639) 
	case KinematicViscosityNanostokes:
		return ((value / 1e4) * 1e-09) 
	case KinematicViscosityMicrostokes:
		return ((value / 1e4) * 1e-06) 
	case KinematicViscosityMillistokes:
		return ((value / 1e4) * 0.001) 
	case KinematicViscosityCentistokes:
		return ((value / 1e4) * 0.01) 
	case KinematicViscosityDecistokes:
		return ((value / 1e4) * 0.1) 
	case KinematicViscosityKilostokes:
		return ((value / 1e4) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a KinematicViscosity) String() string {
	return a.ToString(KinematicViscositySquareMeterPerSecond, 2)
}

// ToString formats the KinematicViscosity to string.
// fractionalDigits -1 for not mention
func (a *KinematicViscosity) ToString(unit KinematicViscosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *KinematicViscosity) getUnitAbbreviation(unit KinematicViscosityUnits) string {
	switch unit { 
	case KinematicViscositySquareMeterPerSecond:
		return "m²/s" 
	case KinematicViscosityStokes:
		return "St" 
	case KinematicViscositySquareFootPerSecond:
		return "ft²/s" 
	case KinematicViscosityNanostokes:
		return "nSt" 
	case KinematicViscosityMicrostokes:
		return "μSt" 
	case KinematicViscosityMillistokes:
		return "mSt" 
	case KinematicViscosityCentistokes:
		return "cSt" 
	case KinematicViscosityDecistokes:
		return "dSt" 
	case KinematicViscosityKilostokes:
		return "kSt" 
	default:
		return ""
	}
}

// Check if the given KinematicViscosity are equals to the current KinematicViscosity
func (a *KinematicViscosity) Equals(other *KinematicViscosity) bool {
	return a.value == other.BaseValue()
}

// Check if the given KinematicViscosity are equals to the current KinematicViscosity
func (a *KinematicViscosity) CompareTo(other *KinematicViscosity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given KinematicViscosity to the current KinematicViscosity.
func (a *KinematicViscosity) Add(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value + other.BaseValue()}
}

// Subtract the given KinematicViscosity to the current KinematicViscosity.
func (a *KinematicViscosity) Subtract(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value - other.BaseValue()}
}

// Multiply the given KinematicViscosity to the current KinematicViscosity.
func (a *KinematicViscosity) Multiply(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value * other.BaseValue()}
}

// Divide the given KinematicViscosity to the current KinematicViscosity.
func (a *KinematicViscosity) Divide(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value / other.BaseValue()}
}