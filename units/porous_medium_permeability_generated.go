// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PorousMediumPermeabilityUnits enumeration
type PorousMediumPermeabilityUnits string

const (
    
        // 
        PorousMediumPermeabilityDarcy PorousMediumPermeabilityUnits = "Darcy"
        // 
        PorousMediumPermeabilitySquareMeter PorousMediumPermeabilityUnits = "SquareMeter"
        // 
        PorousMediumPermeabilitySquareCentimeter PorousMediumPermeabilityUnits = "SquareCentimeter"
        // 
        PorousMediumPermeabilityMicrodarcy PorousMediumPermeabilityUnits = "Microdarcy"
        // 
        PorousMediumPermeabilityMillidarcy PorousMediumPermeabilityUnits = "Millidarcy"
)

// PorousMediumPermeabilityDto represents an PorousMediumPermeability
type PorousMediumPermeabilityDto struct {
	Value float64
	Unit  PorousMediumPermeabilityUnits
}

// PorousMediumPermeabilityDtoFactory struct to group related functions
type PorousMediumPermeabilityDtoFactory struct{}

func (udf PorousMediumPermeabilityDtoFactory) FromJSON(data []byte) (*PorousMediumPermeabilityDto, error) {
	a := PorousMediumPermeabilityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PorousMediumPermeabilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// PorousMediumPermeability struct
type PorousMediumPermeability struct {
	value       float64
    
    darcysLazy *float64 
    square_metersLazy *float64 
    square_centimetersLazy *float64 
    microdarcysLazy *float64 
    millidarcysLazy *float64 
}

// PorousMediumPermeabilityFactory struct to group related functions
type PorousMediumPermeabilityFactory struct{}

func (uf PorousMediumPermeabilityFactory) CreatePorousMediumPermeability(value float64, unit PorousMediumPermeabilityUnits) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, unit)
}

func (uf PorousMediumPermeabilityFactory) FromDto(dto PorousMediumPermeabilityDto) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(dto.Value, dto.Unit)
}

func (uf PorousMediumPermeabilityFactory) FromDtoJSON(data []byte) (*PorousMediumPermeability, error) {
	unitDto, err := PorousMediumPermeabilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PorousMediumPermeabilityFactory{}.FromDto(*unitDto)
}


// FromDarcy creates a new PorousMediumPermeability instance from Darcy.
func (uf PorousMediumPermeabilityFactory) FromDarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityDarcy)
}

// FromSquareMeter creates a new PorousMediumPermeability instance from SquareMeter.
func (uf PorousMediumPermeabilityFactory) FromSquareMeters(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilitySquareMeter)
}

// FromSquareCentimeter creates a new PorousMediumPermeability instance from SquareCentimeter.
func (uf PorousMediumPermeabilityFactory) FromSquareCentimeters(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilitySquareCentimeter)
}

// FromMicrodarcy creates a new PorousMediumPermeability instance from Microdarcy.
func (uf PorousMediumPermeabilityFactory) FromMicrodarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityMicrodarcy)
}

// FromMillidarcy creates a new PorousMediumPermeability instance from Millidarcy.
func (uf PorousMediumPermeabilityFactory) FromMillidarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityMillidarcy)
}




// newPorousMediumPermeability creates a new PorousMediumPermeability.
func newPorousMediumPermeability(value float64, fromUnit PorousMediumPermeabilityUnits) (*PorousMediumPermeability, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &PorousMediumPermeability{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PorousMediumPermeability in SquareMeter.
func (a *PorousMediumPermeability) BaseValue() float64 {
	return a.value
}


// Darcy returns the value in Darcy.
func (a *PorousMediumPermeability) Darcys() float64 {
	if a.darcysLazy != nil {
		return *a.darcysLazy
	}
	darcys := a.convertFromBase(PorousMediumPermeabilityDarcy)
	a.darcysLazy = &darcys
	return darcys
}

// SquareMeter returns the value in SquareMeter.
func (a *PorousMediumPermeability) SquareMeters() float64 {
	if a.square_metersLazy != nil {
		return *a.square_metersLazy
	}
	square_meters := a.convertFromBase(PorousMediumPermeabilitySquareMeter)
	a.square_metersLazy = &square_meters
	return square_meters
}

// SquareCentimeter returns the value in SquareCentimeter.
func (a *PorousMediumPermeability) SquareCentimeters() float64 {
	if a.square_centimetersLazy != nil {
		return *a.square_centimetersLazy
	}
	square_centimeters := a.convertFromBase(PorousMediumPermeabilitySquareCentimeter)
	a.square_centimetersLazy = &square_centimeters
	return square_centimeters
}

// Microdarcy returns the value in Microdarcy.
func (a *PorousMediumPermeability) Microdarcys() float64 {
	if a.microdarcysLazy != nil {
		return *a.microdarcysLazy
	}
	microdarcys := a.convertFromBase(PorousMediumPermeabilityMicrodarcy)
	a.microdarcysLazy = &microdarcys
	return microdarcys
}

// Millidarcy returns the value in Millidarcy.
func (a *PorousMediumPermeability) Millidarcys() float64 {
	if a.millidarcysLazy != nil {
		return *a.millidarcysLazy
	}
	millidarcys := a.convertFromBase(PorousMediumPermeabilityMillidarcy)
	a.millidarcysLazy = &millidarcys
	return millidarcys
}


// ToDto creates an PorousMediumPermeabilityDto representation.
func (a *PorousMediumPermeability) ToDto(holdInUnit *PorousMediumPermeabilityUnits) PorousMediumPermeabilityDto {
	if holdInUnit == nil {
		defaultUnit := PorousMediumPermeabilitySquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return PorousMediumPermeabilityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an PorousMediumPermeabilityDto representation.
func (a *PorousMediumPermeability) ToDtoJSON(holdInUnit *PorousMediumPermeabilityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts PorousMediumPermeability to a specific unit value.
func (a *PorousMediumPermeability) Convert(toUnit PorousMediumPermeabilityUnits) float64 {
	switch toUnit { 
    case PorousMediumPermeabilityDarcy:
		return a.Darcys()
    case PorousMediumPermeabilitySquareMeter:
		return a.SquareMeters()
    case PorousMediumPermeabilitySquareCentimeter:
		return a.SquareCentimeters()
    case PorousMediumPermeabilityMicrodarcy:
		return a.Microdarcys()
    case PorousMediumPermeabilityMillidarcy:
		return a.Millidarcys()
	default:
		return 0
	}
}

func (a *PorousMediumPermeability) convertFromBase(toUnit PorousMediumPermeabilityUnits) float64 {
    value := a.value
	switch toUnit { 
	case PorousMediumPermeabilityDarcy:
		return (value / 9.869233e-13) 
	case PorousMediumPermeabilitySquareMeter:
		return (value) 
	case PorousMediumPermeabilitySquareCentimeter:
		return (value / 1e-4) 
	case PorousMediumPermeabilityMicrodarcy:
		return ((value / 9.869233e-13) / 1e-06) 
	case PorousMediumPermeabilityMillidarcy:
		return ((value / 9.869233e-13) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *PorousMediumPermeability) convertToBase(value float64, fromUnit PorousMediumPermeabilityUnits) float64 {
	switch fromUnit { 
	case PorousMediumPermeabilityDarcy:
		return (value * 9.869233e-13) 
	case PorousMediumPermeabilitySquareMeter:
		return (value) 
	case PorousMediumPermeabilitySquareCentimeter:
		return (value * 1e-4) 
	case PorousMediumPermeabilityMicrodarcy:
		return ((value * 9.869233e-13) * 1e-06) 
	case PorousMediumPermeabilityMillidarcy:
		return ((value * 9.869233e-13) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a PorousMediumPermeability) String() string {
	return a.ToString(PorousMediumPermeabilitySquareMeter, 2)
}

// ToString formats the PorousMediumPermeability to string.
// fractionalDigits -1 for not mention
func (a *PorousMediumPermeability) ToString(unit PorousMediumPermeabilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *PorousMediumPermeability) getUnitAbbreviation(unit PorousMediumPermeabilityUnits) string {
	switch unit { 
	case PorousMediumPermeabilityDarcy:
		return "D" 
	case PorousMediumPermeabilitySquareMeter:
		return "m²" 
	case PorousMediumPermeabilitySquareCentimeter:
		return "cm²" 
	case PorousMediumPermeabilityMicrodarcy:
		return "μD" 
	case PorousMediumPermeabilityMillidarcy:
		return "mD" 
	default:
		return ""
	}
}

// Check if the given PorousMediumPermeability are equals to the current PorousMediumPermeability
func (a *PorousMediumPermeability) Equals(other *PorousMediumPermeability) bool {
	return a.value == other.BaseValue()
}

// Check if the given PorousMediumPermeability are equals to the current PorousMediumPermeability
func (a *PorousMediumPermeability) CompareTo(other *PorousMediumPermeability) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given PorousMediumPermeability to the current PorousMediumPermeability.
func (a *PorousMediumPermeability) Add(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value + other.BaseValue()}
}

// Subtract the given PorousMediumPermeability to the current PorousMediumPermeability.
func (a *PorousMediumPermeability) Subtract(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value - other.BaseValue()}
}

// Multiply the given PorousMediumPermeability to the current PorousMediumPermeability.
func (a *PorousMediumPermeability) Multiply(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value * other.BaseValue()}
}

// Divide the given PorousMediumPermeability to the current PorousMediumPermeability.
func (a *PorousMediumPermeability) Divide(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value / other.BaseValue()}
}