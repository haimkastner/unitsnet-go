package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// WarpingMomentOfInertiaUnits enumeration
type WarpingMomentOfInertiaUnits string

const (
    
        // 
        WarpingMomentOfInertiaMeterToTheSixth WarpingMomentOfInertiaUnits = "MeterToTheSixth"
        // 
        WarpingMomentOfInertiaDecimeterToTheSixth WarpingMomentOfInertiaUnits = "DecimeterToTheSixth"
        // 
        WarpingMomentOfInertiaCentimeterToTheSixth WarpingMomentOfInertiaUnits = "CentimeterToTheSixth"
        // 
        WarpingMomentOfInertiaMillimeterToTheSixth WarpingMomentOfInertiaUnits = "MillimeterToTheSixth"
        // 
        WarpingMomentOfInertiaFootToTheSixth WarpingMomentOfInertiaUnits = "FootToTheSixth"
        // 
        WarpingMomentOfInertiaInchToTheSixth WarpingMomentOfInertiaUnits = "InchToTheSixth"
)

// WarpingMomentOfInertiaDto represents an WarpingMomentOfInertia
type WarpingMomentOfInertiaDto struct {
	Value float64
	Unit  WarpingMomentOfInertiaUnits
}

// WarpingMomentOfInertiaDtoFactory struct to group related functions
type WarpingMomentOfInertiaDtoFactory struct{}

func (udf WarpingMomentOfInertiaDtoFactory) FromJSON(data []byte) (*WarpingMomentOfInertiaDto, error) {
	a := WarpingMomentOfInertiaDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a WarpingMomentOfInertiaDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// WarpingMomentOfInertia struct
type WarpingMomentOfInertia struct {
	value       float64
    
    meters_to_the_sixthLazy *float64 
    decimeters_to_the_sixthLazy *float64 
    centimeters_to_the_sixthLazy *float64 
    millimeters_to_the_sixthLazy *float64 
    feet_to_the_sixthLazy *float64 
    inches_to_the_sixthLazy *float64 
}

// WarpingMomentOfInertiaFactory struct to group related functions
type WarpingMomentOfInertiaFactory struct{}

func (uf WarpingMomentOfInertiaFactory) CreateWarpingMomentOfInertia(value float64, unit WarpingMomentOfInertiaUnits) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, unit)
}

func (uf WarpingMomentOfInertiaFactory) FromDto(dto WarpingMomentOfInertiaDto) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(dto.Value, dto.Unit)
}

func (uf WarpingMomentOfInertiaFactory) FromDtoJSON(data []byte) (*WarpingMomentOfInertia, error) {
	unitDto, err := WarpingMomentOfInertiaDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return WarpingMomentOfInertiaFactory{}.FromDto(*unitDto)
}


// FromMeterToTheSixth creates a new WarpingMomentOfInertia instance from MeterToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromMetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaMeterToTheSixth)
}

// FromDecimeterToTheSixth creates a new WarpingMomentOfInertia instance from DecimeterToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromDecimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaDecimeterToTheSixth)
}

// FromCentimeterToTheSixth creates a new WarpingMomentOfInertia instance from CentimeterToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromCentimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaCentimeterToTheSixth)
}

// FromMillimeterToTheSixth creates a new WarpingMomentOfInertia instance from MillimeterToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromMillimetersToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaMillimeterToTheSixth)
}

// FromFootToTheSixth creates a new WarpingMomentOfInertia instance from FootToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromFeetToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaFootToTheSixth)
}

// FromInchToTheSixth creates a new WarpingMomentOfInertia instance from InchToTheSixth.
func (uf WarpingMomentOfInertiaFactory) FromInchesToTheSixth(value float64) (*WarpingMomentOfInertia, error) {
	return newWarpingMomentOfInertia(value, WarpingMomentOfInertiaInchToTheSixth)
}




// newWarpingMomentOfInertia creates a new WarpingMomentOfInertia.
func newWarpingMomentOfInertia(value float64, fromUnit WarpingMomentOfInertiaUnits) (*WarpingMomentOfInertia, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &WarpingMomentOfInertia{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of WarpingMomentOfInertia in MeterToTheSixth.
func (a *WarpingMomentOfInertia) BaseValue() float64 {
	return a.value
}


// MeterToTheSixth returns the value in MeterToTheSixth.
func (a *WarpingMomentOfInertia) MetersToTheSixth() float64 {
	if a.meters_to_the_sixthLazy != nil {
		return *a.meters_to_the_sixthLazy
	}
	meters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaMeterToTheSixth)
	a.meters_to_the_sixthLazy = &meters_to_the_sixth
	return meters_to_the_sixth
}

// DecimeterToTheSixth returns the value in DecimeterToTheSixth.
func (a *WarpingMomentOfInertia) DecimetersToTheSixth() float64 {
	if a.decimeters_to_the_sixthLazy != nil {
		return *a.decimeters_to_the_sixthLazy
	}
	decimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaDecimeterToTheSixth)
	a.decimeters_to_the_sixthLazy = &decimeters_to_the_sixth
	return decimeters_to_the_sixth
}

// CentimeterToTheSixth returns the value in CentimeterToTheSixth.
func (a *WarpingMomentOfInertia) CentimetersToTheSixth() float64 {
	if a.centimeters_to_the_sixthLazy != nil {
		return *a.centimeters_to_the_sixthLazy
	}
	centimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaCentimeterToTheSixth)
	a.centimeters_to_the_sixthLazy = &centimeters_to_the_sixth
	return centimeters_to_the_sixth
}

// MillimeterToTheSixth returns the value in MillimeterToTheSixth.
func (a *WarpingMomentOfInertia) MillimetersToTheSixth() float64 {
	if a.millimeters_to_the_sixthLazy != nil {
		return *a.millimeters_to_the_sixthLazy
	}
	millimeters_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaMillimeterToTheSixth)
	a.millimeters_to_the_sixthLazy = &millimeters_to_the_sixth
	return millimeters_to_the_sixth
}

// FootToTheSixth returns the value in FootToTheSixth.
func (a *WarpingMomentOfInertia) FeetToTheSixth() float64 {
	if a.feet_to_the_sixthLazy != nil {
		return *a.feet_to_the_sixthLazy
	}
	feet_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaFootToTheSixth)
	a.feet_to_the_sixthLazy = &feet_to_the_sixth
	return feet_to_the_sixth
}

// InchToTheSixth returns the value in InchToTheSixth.
func (a *WarpingMomentOfInertia) InchesToTheSixth() float64 {
	if a.inches_to_the_sixthLazy != nil {
		return *a.inches_to_the_sixthLazy
	}
	inches_to_the_sixth := a.convertFromBase(WarpingMomentOfInertiaInchToTheSixth)
	a.inches_to_the_sixthLazy = &inches_to_the_sixth
	return inches_to_the_sixth
}


// ToDto creates an WarpingMomentOfInertiaDto representation.
func (a *WarpingMomentOfInertia) ToDto(holdInUnit *WarpingMomentOfInertiaUnits) WarpingMomentOfInertiaDto {
	if holdInUnit == nil {
		defaultUnit := WarpingMomentOfInertiaMeterToTheSixth // Default value
		holdInUnit = &defaultUnit
	}

	return WarpingMomentOfInertiaDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an WarpingMomentOfInertiaDto representation.
func (a *WarpingMomentOfInertia) ToDtoJSON(holdInUnit *WarpingMomentOfInertiaUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts WarpingMomentOfInertia to a specific unit value.
func (a *WarpingMomentOfInertia) Convert(toUnit WarpingMomentOfInertiaUnits) float64 {
	switch toUnit { 
    case WarpingMomentOfInertiaMeterToTheSixth:
		return a.MetersToTheSixth()
    case WarpingMomentOfInertiaDecimeterToTheSixth:
		return a.DecimetersToTheSixth()
    case WarpingMomentOfInertiaCentimeterToTheSixth:
		return a.CentimetersToTheSixth()
    case WarpingMomentOfInertiaMillimeterToTheSixth:
		return a.MillimetersToTheSixth()
    case WarpingMomentOfInertiaFootToTheSixth:
		return a.FeetToTheSixth()
    case WarpingMomentOfInertiaInchToTheSixth:
		return a.InchesToTheSixth()
	default:
		return 0
	}
}

func (a *WarpingMomentOfInertia) convertFromBase(toUnit WarpingMomentOfInertiaUnits) float64 {
    value := a.value
	switch toUnit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return (value) 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return (value * 1e6) 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return (value * 1e12) 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return (value * 1e18) 
	case WarpingMomentOfInertiaFootToTheSixth:
		return (value / math.Pow(0.3048, 6)) 
	case WarpingMomentOfInertiaInchToTheSixth:
		return (value / math.Pow(2.54e-2, 6)) 
	default:
		return math.NaN()
	}
}

func (a *WarpingMomentOfInertia) convertToBase(value float64, fromUnit WarpingMomentOfInertiaUnits) float64 {
	switch fromUnit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return (value) 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return (value / 1e6) 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return (value / 1e12) 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return (value / 1e18) 
	case WarpingMomentOfInertiaFootToTheSixth:
		return (value * math.Pow(0.3048, 6)) 
	case WarpingMomentOfInertiaInchToTheSixth:
		return (value * math.Pow(2.54e-2, 6)) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a WarpingMomentOfInertia) String() string {
	return a.ToString(WarpingMomentOfInertiaMeterToTheSixth, 2)
}

// ToString formats the WarpingMomentOfInertia to string.
// fractionalDigits -1 for not mention
func (a *WarpingMomentOfInertia) ToString(unit WarpingMomentOfInertiaUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *WarpingMomentOfInertia) getUnitAbbreviation(unit WarpingMomentOfInertiaUnits) string {
	switch unit { 
	case WarpingMomentOfInertiaMeterToTheSixth:
		return "m⁶" 
	case WarpingMomentOfInertiaDecimeterToTheSixth:
		return "dm⁶" 
	case WarpingMomentOfInertiaCentimeterToTheSixth:
		return "cm⁶" 
	case WarpingMomentOfInertiaMillimeterToTheSixth:
		return "mm⁶" 
	case WarpingMomentOfInertiaFootToTheSixth:
		return "ft⁶" 
	case WarpingMomentOfInertiaInchToTheSixth:
		return "in⁶" 
	default:
		return ""
	}
}

// Check if the given WarpingMomentOfInertia are equals to the current WarpingMomentOfInertia
func (a *WarpingMomentOfInertia) Equals(other *WarpingMomentOfInertia) bool {
	return a.value == other.BaseValue()
}

// Check if the given WarpingMomentOfInertia are equals to the current WarpingMomentOfInertia
func (a *WarpingMomentOfInertia) CompareTo(other *WarpingMomentOfInertia) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given WarpingMomentOfInertia to the current WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Add(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value + other.BaseValue()}
}

// Subtract the given WarpingMomentOfInertia to the current WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Subtract(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value - other.BaseValue()}
}

// Multiply the given WarpingMomentOfInertia to the current WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Multiply(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value * other.BaseValue()}
}

// Divide the given WarpingMomentOfInertia to the current WarpingMomentOfInertia.
func (a *WarpingMomentOfInertia) Divide(other *WarpingMomentOfInertia) *WarpingMomentOfInertia {
	return &WarpingMomentOfInertia{value: a.value / other.BaseValue()}
}