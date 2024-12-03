package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RotationalAccelerationUnits enumeration
type RotationalAccelerationUnits string

const (
    
        // 
        RotationalAccelerationRadianPerSecondSquared RotationalAccelerationUnits = "RadianPerSecondSquared"
        // 
        RotationalAccelerationDegreePerSecondSquared RotationalAccelerationUnits = "DegreePerSecondSquared"
        // 
        RotationalAccelerationRevolutionPerMinutePerSecond RotationalAccelerationUnits = "RevolutionPerMinutePerSecond"
        // 
        RotationalAccelerationRevolutionPerSecondSquared RotationalAccelerationUnits = "RevolutionPerSecondSquared"
)

// RotationalAccelerationDto represents an RotationalAcceleration
type RotationalAccelerationDto struct {
	Value float64
	Unit  RotationalAccelerationUnits
}

// RotationalAccelerationDtoFactory struct to group related functions
type RotationalAccelerationDtoFactory struct{}

func (udf RotationalAccelerationDtoFactory) FromJSON(data []byte) (*RotationalAccelerationDto, error) {
	a := RotationalAccelerationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RotationalAccelerationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RotationalAcceleration struct
type RotationalAcceleration struct {
	value       float64
    
    radians_per_second_squaredLazy *float64 
    degrees_per_second_squaredLazy *float64 
    revolutions_per_minute_per_secondLazy *float64 
    revolutions_per_second_squaredLazy *float64 
}

// RotationalAccelerationFactory struct to group related functions
type RotationalAccelerationFactory struct{}

func (uf RotationalAccelerationFactory) CreateRotationalAcceleration(value float64, unit RotationalAccelerationUnits) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, unit)
}

func (uf RotationalAccelerationFactory) FromDto(dto RotationalAccelerationDto) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(dto.Value, dto.Unit)
}

func (uf RotationalAccelerationFactory) FromDtoJSON(data []byte) (*RotationalAcceleration, error) {
	unitDto, err := RotationalAccelerationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RotationalAccelerationFactory{}.FromDto(*unitDto)
}


// FromRadianPerSecondSquared creates a new RotationalAcceleration instance from RadianPerSecondSquared.
func (uf RotationalAccelerationFactory) FromRadiansPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRadianPerSecondSquared)
}

// FromDegreePerSecondSquared creates a new RotationalAcceleration instance from DegreePerSecondSquared.
func (uf RotationalAccelerationFactory) FromDegreesPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationDegreePerSecondSquared)
}

// FromRevolutionPerMinutePerSecond creates a new RotationalAcceleration instance from RevolutionPerMinutePerSecond.
func (uf RotationalAccelerationFactory) FromRevolutionsPerMinutePerSecond(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRevolutionPerMinutePerSecond)
}

// FromRevolutionPerSecondSquared creates a new RotationalAcceleration instance from RevolutionPerSecondSquared.
func (uf RotationalAccelerationFactory) FromRevolutionsPerSecondSquared(value float64) (*RotationalAcceleration, error) {
	return newRotationalAcceleration(value, RotationalAccelerationRevolutionPerSecondSquared)
}




// newRotationalAcceleration creates a new RotationalAcceleration.
func newRotationalAcceleration(value float64, fromUnit RotationalAccelerationUnits) (*RotationalAcceleration, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RotationalAcceleration{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RotationalAcceleration in RadianPerSecondSquared.
func (a *RotationalAcceleration) BaseValue() float64 {
	return a.value
}


// RadianPerSecondSquared returns the value in RadianPerSecondSquared.
func (a *RotationalAcceleration) RadiansPerSecondSquared() float64 {
	if a.radians_per_second_squaredLazy != nil {
		return *a.radians_per_second_squaredLazy
	}
	radians_per_second_squared := a.convertFromBase(RotationalAccelerationRadianPerSecondSquared)
	a.radians_per_second_squaredLazy = &radians_per_second_squared
	return radians_per_second_squared
}

// DegreePerSecondSquared returns the value in DegreePerSecondSquared.
func (a *RotationalAcceleration) DegreesPerSecondSquared() float64 {
	if a.degrees_per_second_squaredLazy != nil {
		return *a.degrees_per_second_squaredLazy
	}
	degrees_per_second_squared := a.convertFromBase(RotationalAccelerationDegreePerSecondSquared)
	a.degrees_per_second_squaredLazy = &degrees_per_second_squared
	return degrees_per_second_squared
}

// RevolutionPerMinutePerSecond returns the value in RevolutionPerMinutePerSecond.
func (a *RotationalAcceleration) RevolutionsPerMinutePerSecond() float64 {
	if a.revolutions_per_minute_per_secondLazy != nil {
		return *a.revolutions_per_minute_per_secondLazy
	}
	revolutions_per_minute_per_second := a.convertFromBase(RotationalAccelerationRevolutionPerMinutePerSecond)
	a.revolutions_per_minute_per_secondLazy = &revolutions_per_minute_per_second
	return revolutions_per_minute_per_second
}

// RevolutionPerSecondSquared returns the value in RevolutionPerSecondSquared.
func (a *RotationalAcceleration) RevolutionsPerSecondSquared() float64 {
	if a.revolutions_per_second_squaredLazy != nil {
		return *a.revolutions_per_second_squaredLazy
	}
	revolutions_per_second_squared := a.convertFromBase(RotationalAccelerationRevolutionPerSecondSquared)
	a.revolutions_per_second_squaredLazy = &revolutions_per_second_squared
	return revolutions_per_second_squared
}


// ToDto creates an RotationalAccelerationDto representation.
func (a *RotationalAcceleration) ToDto(holdInUnit *RotationalAccelerationUnits) RotationalAccelerationDto {
	if holdInUnit == nil {
		defaultUnit := RotationalAccelerationRadianPerSecondSquared // Default value
		holdInUnit = &defaultUnit
	}

	return RotationalAccelerationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RotationalAccelerationDto representation.
func (a *RotationalAcceleration) ToDtoJSON(holdInUnit *RotationalAccelerationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RotationalAcceleration to a specific unit value.
func (a *RotationalAcceleration) Convert(toUnit RotationalAccelerationUnits) float64 {
	switch toUnit { 
    case RotationalAccelerationRadianPerSecondSquared:
		return a.RadiansPerSecondSquared()
    case RotationalAccelerationDegreePerSecondSquared:
		return a.DegreesPerSecondSquared()
    case RotationalAccelerationRevolutionPerMinutePerSecond:
		return a.RevolutionsPerMinutePerSecond()
    case RotationalAccelerationRevolutionPerSecondSquared:
		return a.RevolutionsPerSecondSquared()
	default:
		return 0
	}
}

func (a *RotationalAcceleration) convertFromBase(toUnit RotationalAccelerationUnits) float64 {
    value := a.value
	switch toUnit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return (value) 
	case RotationalAccelerationDegreePerSecondSquared:
		return ((180 / math.Pi) * value) 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return ((60 / (2 * math.Pi)) * value) 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return ((1 / (2 * math.Pi)) * value) 
	default:
		return math.NaN()
	}
}

func (a *RotationalAcceleration) convertToBase(value float64, fromUnit RotationalAccelerationUnits) float64 {
	switch fromUnit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return (value) 
	case RotationalAccelerationDegreePerSecondSquared:
		return ((math.Pi / 180) * value) 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return (((2 * math.Pi) / 60) * value) 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return ((2 * math.Pi) * value) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RotationalAcceleration) String() string {
	return a.ToString(RotationalAccelerationRadianPerSecondSquared, 2)
}

// ToString formats the RotationalAcceleration to string.
// fractionalDigits -1 for not mention
func (a *RotationalAcceleration) ToString(unit RotationalAccelerationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RotationalAcceleration) getUnitAbbreviation(unit RotationalAccelerationUnits) string {
	switch unit { 
	case RotationalAccelerationRadianPerSecondSquared:
		return "rad/s²" 
	case RotationalAccelerationDegreePerSecondSquared:
		return "°/s²" 
	case RotationalAccelerationRevolutionPerMinutePerSecond:
		return "rpm/s" 
	case RotationalAccelerationRevolutionPerSecondSquared:
		return "r/s²" 
	default:
		return ""
	}
}

// Check if the given RotationalAcceleration are equals to the current RotationalAcceleration
func (a *RotationalAcceleration) Equals(other *RotationalAcceleration) bool {
	return a.value == other.BaseValue()
}

// Check if the given RotationalAcceleration are equals to the current RotationalAcceleration
func (a *RotationalAcceleration) CompareTo(other *RotationalAcceleration) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RotationalAcceleration to the current RotationalAcceleration.
func (a *RotationalAcceleration) Add(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value + other.BaseValue()}
}

// Subtract the given RotationalAcceleration to the current RotationalAcceleration.
func (a *RotationalAcceleration) Subtract(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value - other.BaseValue()}
}

// Multiply the given RotationalAcceleration to the current RotationalAcceleration.
func (a *RotationalAcceleration) Multiply(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value * other.BaseValue()}
}

// Divide the given RotationalAcceleration to the current RotationalAcceleration.
func (a *RotationalAcceleration) Divide(other *RotationalAcceleration) *RotationalAcceleration {
	return &RotationalAcceleration{value: a.value / other.BaseValue()}
}