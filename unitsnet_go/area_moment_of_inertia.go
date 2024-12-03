package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AreaMomentOfInertiaUnits enumeration
type AreaMomentOfInertiaUnits string

const (
    
        // 
        AreaMomentOfInertiaMeterToTheFourth AreaMomentOfInertiaUnits = "MeterToTheFourth"
        // 
        AreaMomentOfInertiaDecimeterToTheFourth AreaMomentOfInertiaUnits = "DecimeterToTheFourth"
        // 
        AreaMomentOfInertiaCentimeterToTheFourth AreaMomentOfInertiaUnits = "CentimeterToTheFourth"
        // 
        AreaMomentOfInertiaMillimeterToTheFourth AreaMomentOfInertiaUnits = "MillimeterToTheFourth"
        // 
        AreaMomentOfInertiaFootToTheFourth AreaMomentOfInertiaUnits = "FootToTheFourth"
        // 
        AreaMomentOfInertiaInchToTheFourth AreaMomentOfInertiaUnits = "InchToTheFourth"
)

// AreaMomentOfInertiaDto represents an AreaMomentOfInertia
type AreaMomentOfInertiaDto struct {
	Value float64
	Unit  AreaMomentOfInertiaUnits
}

// AreaMomentOfInertiaDtoFactory struct to group related functions
type AreaMomentOfInertiaDtoFactory struct{}

func (udf AreaMomentOfInertiaDtoFactory) FromJSON(data []byte) (*AreaMomentOfInertiaDto, error) {
	a := AreaMomentOfInertiaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AreaMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// AreaMomentOfInertia struct
type AreaMomentOfInertia struct {
	value       float64
    
    meters_to_the_fourthLazy *float64 
    decimeters_to_the_fourthLazy *float64 
    centimeters_to_the_fourthLazy *float64 
    millimeters_to_the_fourthLazy *float64 
    feet_to_the_fourthLazy *float64 
    inches_to_the_fourthLazy *float64 
}

// AreaMomentOfInertiaFactory struct to group related functions
type AreaMomentOfInertiaFactory struct{}

func (uf AreaMomentOfInertiaFactory) CreateAreaMomentOfInertia(value float64, unit AreaMomentOfInertiaUnits) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, unit)
}

func (uf AreaMomentOfInertiaFactory) FromDto(dto AreaMomentOfInertiaDto) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(dto.Value, dto.Unit)
}

func (uf AreaMomentOfInertiaFactory) FromDtoJSON(data []byte) (*AreaMomentOfInertia, error) {
	unitDto, err := AreaMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AreaMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromMeterToTheFourth creates a new AreaMomentOfInertia instance from MeterToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromMetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaMeterToTheFourth)
}

// FromDecimeterToTheFourth creates a new AreaMomentOfInertia instance from DecimeterToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromDecimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaDecimeterToTheFourth)
}

// FromCentimeterToTheFourth creates a new AreaMomentOfInertia instance from CentimeterToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromCentimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaCentimeterToTheFourth)
}

// FromMillimeterToTheFourth creates a new AreaMomentOfInertia instance from MillimeterToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromMillimetersToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaMillimeterToTheFourth)
}

// FromFootToTheFourth creates a new AreaMomentOfInertia instance from FootToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromFeetToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaFootToTheFourth)
}

// FromInchToTheFourth creates a new AreaMomentOfInertia instance from InchToTheFourth.
func (uf AreaMomentOfInertiaFactory) FromInchesToTheFourth(value float64) (*AreaMomentOfInertia, error) {
	return newAreaMomentOfInertia(value, AreaMomentOfInertiaInchToTheFourth)
}




// newAreaMomentOfInertia creates a new AreaMomentOfInertia.
func newAreaMomentOfInertia(value float64, fromUnit AreaMomentOfInertiaUnits) (*AreaMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AreaMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AreaMomentOfInertia in MeterToTheFourth.
func (a *AreaMomentOfInertia) BaseValue() float64 {
	return a.value
}


// MeterToTheFourth returns the value in MeterToTheFourth.
func (a *AreaMomentOfInertia) MetersToTheFourth() float64 {
	if a.meters_to_the_fourthLazy != nil {
		return *a.meters_to_the_fourthLazy
	}
	meters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaMeterToTheFourth)
	a.meters_to_the_fourthLazy = &meters_to_the_fourth
	return meters_to_the_fourth
}

// DecimeterToTheFourth returns the value in DecimeterToTheFourth.
func (a *AreaMomentOfInertia) DecimetersToTheFourth() float64 {
	if a.decimeters_to_the_fourthLazy != nil {
		return *a.decimeters_to_the_fourthLazy
	}
	decimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaDecimeterToTheFourth)
	a.decimeters_to_the_fourthLazy = &decimeters_to_the_fourth
	return decimeters_to_the_fourth
}

// CentimeterToTheFourth returns the value in CentimeterToTheFourth.
func (a *AreaMomentOfInertia) CentimetersToTheFourth() float64 {
	if a.centimeters_to_the_fourthLazy != nil {
		return *a.centimeters_to_the_fourthLazy
	}
	centimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaCentimeterToTheFourth)
	a.centimeters_to_the_fourthLazy = &centimeters_to_the_fourth
	return centimeters_to_the_fourth
}

// MillimeterToTheFourth returns the value in MillimeterToTheFourth.
func (a *AreaMomentOfInertia) MillimetersToTheFourth() float64 {
	if a.millimeters_to_the_fourthLazy != nil {
		return *a.millimeters_to_the_fourthLazy
	}
	millimeters_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaMillimeterToTheFourth)
	a.millimeters_to_the_fourthLazy = &millimeters_to_the_fourth
	return millimeters_to_the_fourth
}

// FootToTheFourth returns the value in FootToTheFourth.
func (a *AreaMomentOfInertia) FeetToTheFourth() float64 {
	if a.feet_to_the_fourthLazy != nil {
		return *a.feet_to_the_fourthLazy
	}
	feet_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaFootToTheFourth)
	a.feet_to_the_fourthLazy = &feet_to_the_fourth
	return feet_to_the_fourth
}

// InchToTheFourth returns the value in InchToTheFourth.
func (a *AreaMomentOfInertia) InchesToTheFourth() float64 {
	if a.inches_to_the_fourthLazy != nil {
		return *a.inches_to_the_fourthLazy
	}
	inches_to_the_fourth := a.convertFromBase(AreaMomentOfInertiaInchToTheFourth)
	a.inches_to_the_fourthLazy = &inches_to_the_fourth
	return inches_to_the_fourth
}


// ToDto creates an AreaMomentOfInertiaDto representation.
func (a *AreaMomentOfInertia) ToDto(holdInUnit *AreaMomentOfInertiaUnits) AreaMomentOfInertiaDto {
	if holdInUnit == nil {
		defaultUnit := AreaMomentOfInertiaMeterToTheFourth // Default value
		holdInUnit = &defaultUnit
	}

	return AreaMomentOfInertiaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an AreaMomentOfInertiaDto representation.
func (a *AreaMomentOfInertia) ToDtoJSON(holdInUnit *AreaMomentOfInertiaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts AreaMomentOfInertia to a specific unit value.
func (a *AreaMomentOfInertia) Convert(toUnit AreaMomentOfInertiaUnits) float64 {
	switch toUnit { 
    case AreaMomentOfInertiaMeterToTheFourth:
		return a.MetersToTheFourth()
    case AreaMomentOfInertiaDecimeterToTheFourth:
		return a.DecimetersToTheFourth()
    case AreaMomentOfInertiaCentimeterToTheFourth:
		return a.CentimetersToTheFourth()
    case AreaMomentOfInertiaMillimeterToTheFourth:
		return a.MillimetersToTheFourth()
    case AreaMomentOfInertiaFootToTheFourth:
		return a.FeetToTheFourth()
    case AreaMomentOfInertiaInchToTheFourth:
		return a.InchesToTheFourth()
	default:
		return 0
	}
}

func (a *AreaMomentOfInertia) convertFromBase(toUnit AreaMomentOfInertiaUnits) float64 {
    value := a.value
	switch toUnit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return (value) 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return (value * 1e4) 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return (value * 1e8) 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return (value * 1e12) 
	case AreaMomentOfInertiaFootToTheFourth:
		return (value / math.Pow(0.3048, 4)) 
	case AreaMomentOfInertiaInchToTheFourth:
		return (value / math.Pow(2.54e-2, 4)) 
	default:
		return math.NaN()
	}
}

func (a *AreaMomentOfInertia) convertToBase(value float64, fromUnit AreaMomentOfInertiaUnits) float64 {
	switch fromUnit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return (value) 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return (value / 1e4) 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return (value / 1e8) 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return (value / 1e12) 
	case AreaMomentOfInertiaFootToTheFourth:
		return (value * math.Pow(0.3048, 4)) 
	case AreaMomentOfInertiaInchToTheFourth:
		return (value * math.Pow(2.54e-2, 4)) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a AreaMomentOfInertia) String() string {
	return a.ToString(AreaMomentOfInertiaMeterToTheFourth, 2)
}

// ToString formats the AreaMomentOfInertia to string.
// fractionalDigits -1 for not mention
func (a *AreaMomentOfInertia) ToString(unit AreaMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AreaMomentOfInertia) getUnitAbbreviation(unit AreaMomentOfInertiaUnits) string {
	switch unit { 
	case AreaMomentOfInertiaMeterToTheFourth:
		return "m⁴" 
	case AreaMomentOfInertiaDecimeterToTheFourth:
		return "dm⁴" 
	case AreaMomentOfInertiaCentimeterToTheFourth:
		return "cm⁴" 
	case AreaMomentOfInertiaMillimeterToTheFourth:
		return "mm⁴" 
	case AreaMomentOfInertiaFootToTheFourth:
		return "ft⁴" 
	case AreaMomentOfInertiaInchToTheFourth:
		return "in⁴" 
	default:
		return ""
	}
}

// Check if the given AreaMomentOfInertia are equals to the current AreaMomentOfInertia
func (a *AreaMomentOfInertia) Equals(other *AreaMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// Check if the given AreaMomentOfInertia are equals to the current AreaMomentOfInertia
func (a *AreaMomentOfInertia) CompareTo(other *AreaMomentOfInertia) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given AreaMomentOfInertia to the current AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Add(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract the given AreaMomentOfInertia to the current AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Subtract(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply the given AreaMomentOfInertia to the current AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Multiply(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide the given AreaMomentOfInertia to the current AreaMomentOfInertia.
func (a *AreaMomentOfInertia) Divide(other *AreaMomentOfInertia) *AreaMomentOfInertia {
	return &AreaMomentOfInertia{value: a.value / other.BaseValue()}
}