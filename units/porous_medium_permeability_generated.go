// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PorousMediumPermeabilityUnits defines various units of PorousMediumPermeability.
type PorousMediumPermeabilityUnits string

const (
    
        // The darcy (or darcy unit) and millidarcy (md or mD) are units of permeability, named after Henry Darcy. They are not SI units, but they are widely used in petroleum engineering and geology.
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

var internalPorousMediumPermeabilityUnitsMap = map[PorousMediumPermeabilityUnits]bool{
	
	PorousMediumPermeabilityDarcy: true,
	PorousMediumPermeabilitySquareMeter: true,
	PorousMediumPermeabilitySquareCentimeter: true,
	PorousMediumPermeabilityMicrodarcy: true,
	PorousMediumPermeabilityMillidarcy: true,
}

// PorousMediumPermeabilityDto represents a PorousMediumPermeability measurement with a numerical value and its corresponding unit.
type PorousMediumPermeabilityDto struct {
    // Value is the numerical representation of the PorousMediumPermeability.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the PorousMediumPermeability, as defined in the PorousMediumPermeabilityUnits enumeration.
	Unit  PorousMediumPermeabilityUnits `json:"unit" validate:"required,oneof=Darcy SquareMeter SquareCentimeter Microdarcy Millidarcy"`
}

// PorousMediumPermeabilityDtoFactory groups methods for creating and serializing PorousMediumPermeabilityDto objects.
type PorousMediumPermeabilityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PorousMediumPermeabilityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PorousMediumPermeabilityDtoFactory) FromJSON(data []byte) (*PorousMediumPermeabilityDto, error) {
	a := PorousMediumPermeabilityDto{}

    // Parse JSON into PorousMediumPermeabilityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a PorousMediumPermeabilityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PorousMediumPermeabilityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// PorousMediumPermeability represents a measurement in a of PorousMediumPermeability.
//
// None
type PorousMediumPermeability struct {
	// value is the base measurement stored internally.
	value       float64
    
    darcysLazy *float64 
    square_metersLazy *float64 
    square_centimetersLazy *float64 
    microdarcysLazy *float64 
    millidarcysLazy *float64 
}

// PorousMediumPermeabilityFactory groups methods for creating PorousMediumPermeability instances.
type PorousMediumPermeabilityFactory struct{}

// CreatePorousMediumPermeability creates a new PorousMediumPermeability instance from the given value and unit.
func (uf PorousMediumPermeabilityFactory) CreatePorousMediumPermeability(value float64, unit PorousMediumPermeabilityUnits) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, unit)
}

// FromDto converts a PorousMediumPermeabilityDto to a PorousMediumPermeability instance.
func (uf PorousMediumPermeabilityFactory) FromDto(dto PorousMediumPermeabilityDto) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a PorousMediumPermeability instance.
func (uf PorousMediumPermeabilityFactory) FromDtoJSON(data []byte) (*PorousMediumPermeability, error) {
	unitDto, err := PorousMediumPermeabilityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PorousMediumPermeabilityDto from JSON: %w", err)
	}
	return PorousMediumPermeabilityFactory{}.FromDto(*unitDto)
}


// FromDarcys creates a new PorousMediumPermeability instance from a value in Darcys.
func (uf PorousMediumPermeabilityFactory) FromDarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityDarcy)
}

// FromSquareMeters creates a new PorousMediumPermeability instance from a value in SquareMeters.
func (uf PorousMediumPermeabilityFactory) FromSquareMeters(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilitySquareMeter)
}

// FromSquareCentimeters creates a new PorousMediumPermeability instance from a value in SquareCentimeters.
func (uf PorousMediumPermeabilityFactory) FromSquareCentimeters(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilitySquareCentimeter)
}

// FromMicrodarcys creates a new PorousMediumPermeability instance from a value in Microdarcys.
func (uf PorousMediumPermeabilityFactory) FromMicrodarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityMicrodarcy)
}

// FromMillidarcys creates a new PorousMediumPermeability instance from a value in Millidarcys.
func (uf PorousMediumPermeabilityFactory) FromMillidarcys(value float64) (*PorousMediumPermeability, error) {
	return newPorousMediumPermeability(value, PorousMediumPermeabilityMillidarcy)
}


// newPorousMediumPermeability creates a new PorousMediumPermeability.
func newPorousMediumPermeability(value float64, fromUnit PorousMediumPermeabilityUnits) (*PorousMediumPermeability, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalPorousMediumPermeabilityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in PorousMediumPermeabilityUnits", fromUnit)
	}
	a := &PorousMediumPermeability{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of PorousMediumPermeability in SquareMeter unit (the base/default quantity).
func (a *PorousMediumPermeability) BaseValue() float64 {
	return a.value
}


// Darcys returns the PorousMediumPermeability value in Darcys.
//
// The darcy (or darcy unit) and millidarcy (md or mD) are units of permeability, named after Henry Darcy. They are not SI units, but they are widely used in petroleum engineering and geology.
func (a *PorousMediumPermeability) Darcys() float64 {
	if a.darcysLazy != nil {
		return *a.darcysLazy
	}
	darcys := a.convertFromBase(PorousMediumPermeabilityDarcy)
	a.darcysLazy = &darcys
	return darcys
}

// SquareMeters returns the PorousMediumPermeability value in SquareMeters.
//
// 
func (a *PorousMediumPermeability) SquareMeters() float64 {
	if a.square_metersLazy != nil {
		return *a.square_metersLazy
	}
	square_meters := a.convertFromBase(PorousMediumPermeabilitySquareMeter)
	a.square_metersLazy = &square_meters
	return square_meters
}

// SquareCentimeters returns the PorousMediumPermeability value in SquareCentimeters.
//
// 
func (a *PorousMediumPermeability) SquareCentimeters() float64 {
	if a.square_centimetersLazy != nil {
		return *a.square_centimetersLazy
	}
	square_centimeters := a.convertFromBase(PorousMediumPermeabilitySquareCentimeter)
	a.square_centimetersLazy = &square_centimeters
	return square_centimeters
}

// Microdarcys returns the PorousMediumPermeability value in Microdarcys.
//
// 
func (a *PorousMediumPermeability) Microdarcys() float64 {
	if a.microdarcysLazy != nil {
		return *a.microdarcysLazy
	}
	microdarcys := a.convertFromBase(PorousMediumPermeabilityMicrodarcy)
	a.microdarcysLazy = &microdarcys
	return microdarcys
}

// Millidarcys returns the PorousMediumPermeability value in Millidarcys.
//
// 
func (a *PorousMediumPermeability) Millidarcys() float64 {
	if a.millidarcysLazy != nil {
		return *a.millidarcysLazy
	}
	millidarcys := a.convertFromBase(PorousMediumPermeabilityMillidarcy)
	a.millidarcysLazy = &millidarcys
	return millidarcys
}


// ToDto creates a PorousMediumPermeabilityDto representation from the PorousMediumPermeability instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeter by default.
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

// ToDtoJSON creates a JSON representation of the PorousMediumPermeability instance.
//
// If the provided holdInUnit is nil, the value will be repesented by SquareMeter by default.
func (a *PorousMediumPermeability) ToDtoJSON(holdInUnit *PorousMediumPermeabilityUnits) ([]byte, error) {
	// Convert to PorousMediumPermeabilityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a PorousMediumPermeability to a specific unit value.
// The function uses the provided unit type (PorousMediumPermeabilityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the PorousMediumPermeability in the default unit (SquareMeter),
// formatted to two decimal places.
func (a PorousMediumPermeability) String() string {
	return a.ToString(PorousMediumPermeabilitySquareMeter, 2)
}

// ToString formats the PorousMediumPermeability value as a string with the specified unit and fractional digits.
// It converts the PorousMediumPermeability to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the PorousMediumPermeability value will be converted (e.g., SquareMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the PorousMediumPermeability with the unit abbreviation.
func (a *PorousMediumPermeability) ToString(unit PorousMediumPermeabilityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetPorousMediumPermeabilityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetPorousMediumPermeabilityAbbreviation(unit))
}

// Equals checks if the given PorousMediumPermeability is equal to the current PorousMediumPermeability.
//
// Parameters:
//    other: The PorousMediumPermeability to compare against.
//
// Returns:
//    bool: Returns true if both PorousMediumPermeability are equal, false otherwise.
func (a *PorousMediumPermeability) Equals(other *PorousMediumPermeability) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current PorousMediumPermeability with another PorousMediumPermeability.
// It returns -1 if the current PorousMediumPermeability is less than the other PorousMediumPermeability, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The PorousMediumPermeability to compare against.
//
// Returns:
//    int: -1 if the current PorousMediumPermeability is less, 1 if greater, and 0 if equal.
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

// Add adds the given PorousMediumPermeability to the current PorousMediumPermeability and returns the result.
// The result is a new PorousMediumPermeability instance with the sum of the values.
//
// Parameters:
//    other: The PorousMediumPermeability to add to the current PorousMediumPermeability.
//
// Returns:
//    *PorousMediumPermeability: A new PorousMediumPermeability instance representing the sum of both PorousMediumPermeability.
func (a *PorousMediumPermeability) Add(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given PorousMediumPermeability from the current PorousMediumPermeability and returns the result.
// The result is a new PorousMediumPermeability instance with the difference of the values.
//
// Parameters:
//    other: The PorousMediumPermeability to subtract from the current PorousMediumPermeability.
//
// Returns:
//    *PorousMediumPermeability: A new PorousMediumPermeability instance representing the difference of both PorousMediumPermeability.
func (a *PorousMediumPermeability) Subtract(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current PorousMediumPermeability by the given PorousMediumPermeability and returns the result.
// The result is a new PorousMediumPermeability instance with the product of the values.
//
// Parameters:
//    other: The PorousMediumPermeability to multiply with the current PorousMediumPermeability.
//
// Returns:
//    *PorousMediumPermeability: A new PorousMediumPermeability instance representing the product of both PorousMediumPermeability.
func (a *PorousMediumPermeability) Multiply(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value * other.BaseValue()}
}

// Divide divides the current PorousMediumPermeability by the given PorousMediumPermeability and returns the result.
// The result is a new PorousMediumPermeability instance with the quotient of the values.
//
// Parameters:
//    other: The PorousMediumPermeability to divide the current PorousMediumPermeability by.
//
// Returns:
//    *PorousMediumPermeability: A new PorousMediumPermeability instance representing the quotient of both PorousMediumPermeability.
func (a *PorousMediumPermeability) Divide(other *PorousMediumPermeability) *PorousMediumPermeability {
	return &PorousMediumPermeability{value: a.value / other.BaseValue()}
}

// GetPorousMediumPermeabilityAbbreviation gets the unit abbreviation.
func GetPorousMediumPermeabilityAbbreviation(unit PorousMediumPermeabilityUnits) string {
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