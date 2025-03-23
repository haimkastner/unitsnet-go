// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// KinematicViscosityUnits defines various units of KinematicViscosity.
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

var internalKinematicViscosityUnitsMap = map[KinematicViscosityUnits]bool{
	
	KinematicViscositySquareMeterPerSecond: true,
	KinematicViscosityStokes: true,
	KinematicViscositySquareFootPerSecond: true,
	KinematicViscosityNanostokes: true,
	KinematicViscosityMicrostokes: true,
	KinematicViscosityMillistokes: true,
	KinematicViscosityCentistokes: true,
	KinematicViscosityDecistokes: true,
	KinematicViscosityKilostokes: true,
}

// KinematicViscosityDto represents a KinematicViscosity measurement with a numerical value and its corresponding unit.
type KinematicViscosityDto struct {
    // Value is the numerical representation of the KinematicViscosity.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the KinematicViscosity, as defined in the KinematicViscosityUnits enumeration.
	Unit  KinematicViscosityUnits `json:"unit" validate:"required,oneof=SquareMeterPerSecond Stokes SquareFootPerSecond Nanostokes Microstokes Millistokes Centistokes Decistokes Kilostokes"`
}

// KinematicViscosityDtoFactory groups methods for creating and serializing KinematicViscosityDto objects.
type KinematicViscosityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a KinematicViscosityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf KinematicViscosityDtoFactory) FromJSON(data []byte) (*KinematicViscosityDto, error) {
	a := KinematicViscosityDto{}

    // Parse JSON into KinematicViscosityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a KinematicViscosityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a KinematicViscosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// KinematicViscosity represents a measurement in a of KinematicViscosity.
//
// The viscosity of a fluid is a measure of its resistance to gradual deformation by shear stress or tensile stress.
type KinematicViscosity struct {
	// value is the base measurement stored internally.
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

// KinematicViscosityFactory groups methods for creating KinematicViscosity instances.
type KinematicViscosityFactory struct{}

// CreateKinematicViscosity creates a new KinematicViscosity instance from the given value and unit.
func (uf KinematicViscosityFactory) CreateKinematicViscosity(value float64, unit KinematicViscosityUnits) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, unit)
}

// FromDto converts a KinematicViscosityDto to a KinematicViscosity instance.
func (uf KinematicViscosityFactory) FromDto(dto KinematicViscosityDto) (*KinematicViscosity, error) {
	return newKinematicViscosity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a KinematicViscosity instance.
func (uf KinematicViscosityFactory) FromDtoJSON(data []byte) (*KinematicViscosity, error) {
	unitDto, err := KinematicViscosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse KinematicViscosityDto from JSON: %w", err)
	}
	return KinematicViscosityFactory{}.FromDto(*unitDto)
}


// FromSquareMetersPerSecond creates a new KinematicViscosity instance from a value in SquareMetersPerSecond.
func (uf KinematicViscosityFactory) FromSquareMetersPerSecond(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscositySquareMeterPerSecond)
}

// FromStokes creates a new KinematicViscosity instance from a value in Stokes.
func (uf KinematicViscosityFactory) FromStokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityStokes)
}

// FromSquareFeetPerSecond creates a new KinematicViscosity instance from a value in SquareFeetPerSecond.
func (uf KinematicViscosityFactory) FromSquareFeetPerSecond(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscositySquareFootPerSecond)
}

// FromNanostokes creates a new KinematicViscosity instance from a value in Nanostokes.
func (uf KinematicViscosityFactory) FromNanostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityNanostokes)
}

// FromMicrostokes creates a new KinematicViscosity instance from a value in Microstokes.
func (uf KinematicViscosityFactory) FromMicrostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityMicrostokes)
}

// FromMillistokes creates a new KinematicViscosity instance from a value in Millistokes.
func (uf KinematicViscosityFactory) FromMillistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityMillistokes)
}

// FromCentistokes creates a new KinematicViscosity instance from a value in Centistokes.
func (uf KinematicViscosityFactory) FromCentistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityCentistokes)
}

// FromDecistokes creates a new KinematicViscosity instance from a value in Decistokes.
func (uf KinematicViscosityFactory) FromDecistokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityDecistokes)
}

// FromKilostokes creates a new KinematicViscosity instance from a value in Kilostokes.
func (uf KinematicViscosityFactory) FromKilostokes(value float64) (*KinematicViscosity, error) {
	return newKinematicViscosity(value, KinematicViscosityKilostokes)
}


// newKinematicViscosity creates a new KinematicViscosity.
func newKinematicViscosity(value float64, fromUnit KinematicViscosityUnits) (*KinematicViscosity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalKinematicViscosityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in KinematicViscosityUnits", fromUnit)
	}
	a := &KinematicViscosity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of KinematicViscosity in SquareMeterPerSecond unit (the base/default quantity).
func (a *KinematicViscosity) BaseValue() float64 {
	return a.value
}


// SquareMetersPerSecond returns the KinematicViscosity value in SquareMetersPerSecond.
//
// 
func (a *KinematicViscosity) SquareMetersPerSecond() float64 {
	if a.square_meters_per_secondLazy != nil {
		return *a.square_meters_per_secondLazy
	}
	square_meters_per_second := a.convertFromBase(KinematicViscositySquareMeterPerSecond)
	a.square_meters_per_secondLazy = &square_meters_per_second
	return square_meters_per_second
}

// Stokes returns the KinematicViscosity value in Stokes.
//
// 
func (a *KinematicViscosity) Stokes() float64 {
	if a.stokesLazy != nil {
		return *a.stokesLazy
	}
	stokes := a.convertFromBase(KinematicViscosityStokes)
	a.stokesLazy = &stokes
	return stokes
}

// SquareFeetPerSecond returns the KinematicViscosity value in SquareFeetPerSecond.
//
// 
func (a *KinematicViscosity) SquareFeetPerSecond() float64 {
	if a.square_feet_per_secondLazy != nil {
		return *a.square_feet_per_secondLazy
	}
	square_feet_per_second := a.convertFromBase(KinematicViscositySquareFootPerSecond)
	a.square_feet_per_secondLazy = &square_feet_per_second
	return square_feet_per_second
}

// Nanostokes returns the KinematicViscosity value in Nanostokes.
//
// 
func (a *KinematicViscosity) Nanostokes() float64 {
	if a.nanostokesLazy != nil {
		return *a.nanostokesLazy
	}
	nanostokes := a.convertFromBase(KinematicViscosityNanostokes)
	a.nanostokesLazy = &nanostokes
	return nanostokes
}

// Microstokes returns the KinematicViscosity value in Microstokes.
//
// 
func (a *KinematicViscosity) Microstokes() float64 {
	if a.microstokesLazy != nil {
		return *a.microstokesLazy
	}
	microstokes := a.convertFromBase(KinematicViscosityMicrostokes)
	a.microstokesLazy = &microstokes
	return microstokes
}

// Millistokes returns the KinematicViscosity value in Millistokes.
//
// 
func (a *KinematicViscosity) Millistokes() float64 {
	if a.millistokesLazy != nil {
		return *a.millistokesLazy
	}
	millistokes := a.convertFromBase(KinematicViscosityMillistokes)
	a.millistokesLazy = &millistokes
	return millistokes
}

// Centistokes returns the KinematicViscosity value in Centistokes.
//
// 
func (a *KinematicViscosity) Centistokes() float64 {
	if a.centistokesLazy != nil {
		return *a.centistokesLazy
	}
	centistokes := a.convertFromBase(KinematicViscosityCentistokes)
	a.centistokesLazy = &centistokes
	return centistokes
}

// Decistokes returns the KinematicViscosity value in Decistokes.
//
// 
func (a *KinematicViscosity) Decistokes() float64 {
	if a.decistokesLazy != nil {
		return *a.decistokesLazy
	}
	decistokes := a.convertFromBase(KinematicViscosityDecistokes)
	a.decistokesLazy = &decistokes
	return decistokes
}

// Kilostokes returns the KinematicViscosity value in Kilostokes.
//
// 
func (a *KinematicViscosity) Kilostokes() float64 {
	if a.kilostokesLazy != nil {
		return *a.kilostokesLazy
	}
	kilostokes := a.convertFromBase(KinematicViscosityKilostokes)
	a.kilostokesLazy = &kilostokes
	return kilostokes
}


// ToDto creates a KinematicViscosityDto representation from the KinematicViscosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterPerSecond by default.
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

// ToDtoJSON creates a JSON representation of the KinematicViscosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeterPerSecond by default.
func (a *KinematicViscosity) ToDtoJSON(holdInUnit *KinematicViscosityUnits) ([]byte, error) {
	// Convert to KinematicViscosityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a KinematicViscosity to a specific unit value.
// The function uses the provided unit type (KinematicViscosityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the KinematicViscosity in the default unit (SquareMeterPerSecond),
// formatted to two decimal places.
func (a KinematicViscosity) String() string {
	return a.ToString(KinematicViscositySquareMeterPerSecond, 2)
}

// ToString formats the KinematicViscosity value as a string with the specified unit and fractional digits.
// It converts the KinematicViscosity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the KinematicViscosity value will be converted (e.g., SquareMeterPerSecond))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the KinematicViscosity with the unit abbreviation.
func (a *KinematicViscosity) ToString(unit KinematicViscosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetKinematicViscosityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetKinematicViscosityAbbreviation(unit))
}

// Equals checks if the given KinematicViscosity is equal to the current KinematicViscosity.
//
// Parameters:
//    other: The KinematicViscosity to compare against.
//
// Returns:
//    bool: Returns true if both KinematicViscosity are equal, false otherwise.
func (a *KinematicViscosity) Equals(other *KinematicViscosity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current KinematicViscosity with another KinematicViscosity.
// It returns -1 if the current KinematicViscosity is less than the other KinematicViscosity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The KinematicViscosity to compare against.
//
// Returns:
//    int: -1 if the current KinematicViscosity is less, 1 if greater, and 0 if equal.
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

// Add adds the given KinematicViscosity to the current KinematicViscosity and returns the result.
// The result is a new KinematicViscosity instance with the sum of the values.
//
// Parameters:
//    other: The KinematicViscosity to add to the current KinematicViscosity.
//
// Returns:
//    *KinematicViscosity: A new KinematicViscosity instance representing the sum of both KinematicViscosity.
func (a *KinematicViscosity) Add(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given KinematicViscosity from the current KinematicViscosity and returns the result.
// The result is a new KinematicViscosity instance with the difference of the values.
//
// Parameters:
//    other: The KinematicViscosity to subtract from the current KinematicViscosity.
//
// Returns:
//    *KinematicViscosity: A new KinematicViscosity instance representing the difference of both KinematicViscosity.
func (a *KinematicViscosity) Subtract(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current KinematicViscosity by the given KinematicViscosity and returns the result.
// The result is a new KinematicViscosity instance with the product of the values.
//
// Parameters:
//    other: The KinematicViscosity to multiply with the current KinematicViscosity.
//
// Returns:
//    *KinematicViscosity: A new KinematicViscosity instance representing the product of both KinematicViscosity.
func (a *KinematicViscosity) Multiply(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value * other.BaseValue()}
}

// Divide divides the current KinematicViscosity by the given KinematicViscosity and returns the result.
// The result is a new KinematicViscosity instance with the quotient of the values.
//
// Parameters:
//    other: The KinematicViscosity to divide the current KinematicViscosity by.
//
// Returns:
//    *KinematicViscosity: A new KinematicViscosity instance representing the quotient of both KinematicViscosity.
func (a *KinematicViscosity) Divide(other *KinematicViscosity) *KinematicViscosity {
	return &KinematicViscosity{value: a.value / other.BaseValue()}
}

// GetKinematicViscosityAbbreviation gets the unit abbreviation.
func GetKinematicViscosityAbbreviation(unit KinematicViscosityUnits) string {
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