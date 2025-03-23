// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricResistivityUnits defines various units of ElectricResistivity.
type ElectricResistivityUnits string

const (
    
        // 
        ElectricResistivityOhmMeter ElectricResistivityUnits = "OhmMeter"
        // 
        ElectricResistivityOhmCentimeter ElectricResistivityUnits = "OhmCentimeter"
        // 
        ElectricResistivityPicoohmMeter ElectricResistivityUnits = "PicoohmMeter"
        // 
        ElectricResistivityNanoohmMeter ElectricResistivityUnits = "NanoohmMeter"
        // 
        ElectricResistivityMicroohmMeter ElectricResistivityUnits = "MicroohmMeter"
        // 
        ElectricResistivityMilliohmMeter ElectricResistivityUnits = "MilliohmMeter"
        // 
        ElectricResistivityKiloohmMeter ElectricResistivityUnits = "KiloohmMeter"
        // 
        ElectricResistivityMegaohmMeter ElectricResistivityUnits = "MegaohmMeter"
        // 
        ElectricResistivityPicoohmCentimeter ElectricResistivityUnits = "PicoohmCentimeter"
        // 
        ElectricResistivityNanoohmCentimeter ElectricResistivityUnits = "NanoohmCentimeter"
        // 
        ElectricResistivityMicroohmCentimeter ElectricResistivityUnits = "MicroohmCentimeter"
        // 
        ElectricResistivityMilliohmCentimeter ElectricResistivityUnits = "MilliohmCentimeter"
        // 
        ElectricResistivityKiloohmCentimeter ElectricResistivityUnits = "KiloohmCentimeter"
        // 
        ElectricResistivityMegaohmCentimeter ElectricResistivityUnits = "MegaohmCentimeter"
)

var internalElectricResistivityUnitsMap = map[ElectricResistivityUnits]bool{
	
	ElectricResistivityOhmMeter: true,
	ElectricResistivityOhmCentimeter: true,
	ElectricResistivityPicoohmMeter: true,
	ElectricResistivityNanoohmMeter: true,
	ElectricResistivityMicroohmMeter: true,
	ElectricResistivityMilliohmMeter: true,
	ElectricResistivityKiloohmMeter: true,
	ElectricResistivityMegaohmMeter: true,
	ElectricResistivityPicoohmCentimeter: true,
	ElectricResistivityNanoohmCentimeter: true,
	ElectricResistivityMicroohmCentimeter: true,
	ElectricResistivityMilliohmCentimeter: true,
	ElectricResistivityKiloohmCentimeter: true,
	ElectricResistivityMegaohmCentimeter: true,
}

// ElectricResistivityDto represents a ElectricResistivity measurement with a numerical value and its corresponding unit.
type ElectricResistivityDto struct {
    // Value is the numerical representation of the ElectricResistivity.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the ElectricResistivity, as defined in the ElectricResistivityUnits enumeration.
	Unit  ElectricResistivityUnits `json:"unit" validate:"required,oneof=OhmMeter OhmCentimeter PicoohmMeter NanoohmMeter MicroohmMeter MilliohmMeter KiloohmMeter MegaohmMeter PicoohmCentimeter NanoohmCentimeter MicroohmCentimeter MilliohmCentimeter KiloohmCentimeter MegaohmCentimeter"`
}

// ElectricResistivityDtoFactory groups methods for creating and serializing ElectricResistivityDto objects.
type ElectricResistivityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricResistivityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricResistivityDtoFactory) FromJSON(data []byte) (*ElectricResistivityDto, error) {
	a := ElectricResistivityDto{}

    // Parse JSON into ElectricResistivityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricResistivityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricResistivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricResistivity represents a measurement in a of ElectricResistivity.
//
// Electrical resistivity (also known as resistivity, specific electrical resistance, or volume resistivity) is a fundamental property that quantifies how strongly a given material opposes the flow of electric current.
type ElectricResistivity struct {
	// value is the base measurement stored internally.
	value       float64
    
    ohm_metersLazy *float64 
    ohms_centimeterLazy *float64 
    picoohm_metersLazy *float64 
    nanoohm_metersLazy *float64 
    microohm_metersLazy *float64 
    milliohm_metersLazy *float64 
    kiloohm_metersLazy *float64 
    megaohm_metersLazy *float64 
    picoohms_centimeterLazy *float64 
    nanoohms_centimeterLazy *float64 
    microohms_centimeterLazy *float64 
    milliohms_centimeterLazy *float64 
    kiloohms_centimeterLazy *float64 
    megaohms_centimeterLazy *float64 
}

// ElectricResistivityFactory groups methods for creating ElectricResistivity instances.
type ElectricResistivityFactory struct{}

// CreateElectricResistivity creates a new ElectricResistivity instance from the given value and unit.
func (uf ElectricResistivityFactory) CreateElectricResistivity(value float64, unit ElectricResistivityUnits) (*ElectricResistivity, error) {
	return newElectricResistivity(value, unit)
}

// FromDto converts a ElectricResistivityDto to a ElectricResistivity instance.
func (uf ElectricResistivityFactory) FromDto(dto ElectricResistivityDto) (*ElectricResistivity, error) {
	return newElectricResistivity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricResistivity instance.
func (uf ElectricResistivityFactory) FromDtoJSON(data []byte) (*ElectricResistivity, error) {
	unitDto, err := ElectricResistivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricResistivityDto from JSON: %w", err)
	}
	return ElectricResistivityFactory{}.FromDto(*unitDto)
}


// FromOhmMeters creates a new ElectricResistivity instance from a value in OhmMeters.
func (uf ElectricResistivityFactory) FromOhmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityOhmMeter)
}

// FromOhmsCentimeter creates a new ElectricResistivity instance from a value in OhmsCentimeter.
func (uf ElectricResistivityFactory) FromOhmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityOhmCentimeter)
}

// FromPicoohmMeters creates a new ElectricResistivity instance from a value in PicoohmMeters.
func (uf ElectricResistivityFactory) FromPicoohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityPicoohmMeter)
}

// FromNanoohmMeters creates a new ElectricResistivity instance from a value in NanoohmMeters.
func (uf ElectricResistivityFactory) FromNanoohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityNanoohmMeter)
}

// FromMicroohmMeters creates a new ElectricResistivity instance from a value in MicroohmMeters.
func (uf ElectricResistivityFactory) FromMicroohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMicroohmMeter)
}

// FromMilliohmMeters creates a new ElectricResistivity instance from a value in MilliohmMeters.
func (uf ElectricResistivityFactory) FromMilliohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMilliohmMeter)
}

// FromKiloohmMeters creates a new ElectricResistivity instance from a value in KiloohmMeters.
func (uf ElectricResistivityFactory) FromKiloohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityKiloohmMeter)
}

// FromMegaohmMeters creates a new ElectricResistivity instance from a value in MegaohmMeters.
func (uf ElectricResistivityFactory) FromMegaohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMegaohmMeter)
}

// FromPicoohmsCentimeter creates a new ElectricResistivity instance from a value in PicoohmsCentimeter.
func (uf ElectricResistivityFactory) FromPicoohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityPicoohmCentimeter)
}

// FromNanoohmsCentimeter creates a new ElectricResistivity instance from a value in NanoohmsCentimeter.
func (uf ElectricResistivityFactory) FromNanoohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityNanoohmCentimeter)
}

// FromMicroohmsCentimeter creates a new ElectricResistivity instance from a value in MicroohmsCentimeter.
func (uf ElectricResistivityFactory) FromMicroohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMicroohmCentimeter)
}

// FromMilliohmsCentimeter creates a new ElectricResistivity instance from a value in MilliohmsCentimeter.
func (uf ElectricResistivityFactory) FromMilliohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMilliohmCentimeter)
}

// FromKiloohmsCentimeter creates a new ElectricResistivity instance from a value in KiloohmsCentimeter.
func (uf ElectricResistivityFactory) FromKiloohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityKiloohmCentimeter)
}

// FromMegaohmsCentimeter creates a new ElectricResistivity instance from a value in MegaohmsCentimeter.
func (uf ElectricResistivityFactory) FromMegaohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMegaohmCentimeter)
}


// newElectricResistivity creates a new ElectricResistivity.
func newElectricResistivity(value float64, fromUnit ElectricResistivityUnits) (*ElectricResistivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricResistivityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricResistivityUnits", fromUnit)
	}
	a := &ElectricResistivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricResistivity in OhmMeter unit (the base/default quantity).
func (a *ElectricResistivity) BaseValue() float64 {
	return a.value
}


// OhmMeters returns the ElectricResistivity value in OhmMeters.
//
// 
func (a *ElectricResistivity) OhmMeters() float64 {
	if a.ohm_metersLazy != nil {
		return *a.ohm_metersLazy
	}
	ohm_meters := a.convertFromBase(ElectricResistivityOhmMeter)
	a.ohm_metersLazy = &ohm_meters
	return ohm_meters
}

// OhmsCentimeter returns the ElectricResistivity value in OhmsCentimeter.
//
// 
func (a *ElectricResistivity) OhmsCentimeter() float64 {
	if a.ohms_centimeterLazy != nil {
		return *a.ohms_centimeterLazy
	}
	ohms_centimeter := a.convertFromBase(ElectricResistivityOhmCentimeter)
	a.ohms_centimeterLazy = &ohms_centimeter
	return ohms_centimeter
}

// PicoohmMeters returns the ElectricResistivity value in PicoohmMeters.
//
// 
func (a *ElectricResistivity) PicoohmMeters() float64 {
	if a.picoohm_metersLazy != nil {
		return *a.picoohm_metersLazy
	}
	picoohm_meters := a.convertFromBase(ElectricResistivityPicoohmMeter)
	a.picoohm_metersLazy = &picoohm_meters
	return picoohm_meters
}

// NanoohmMeters returns the ElectricResistivity value in NanoohmMeters.
//
// 
func (a *ElectricResistivity) NanoohmMeters() float64 {
	if a.nanoohm_metersLazy != nil {
		return *a.nanoohm_metersLazy
	}
	nanoohm_meters := a.convertFromBase(ElectricResistivityNanoohmMeter)
	a.nanoohm_metersLazy = &nanoohm_meters
	return nanoohm_meters
}

// MicroohmMeters returns the ElectricResistivity value in MicroohmMeters.
//
// 
func (a *ElectricResistivity) MicroohmMeters() float64 {
	if a.microohm_metersLazy != nil {
		return *a.microohm_metersLazy
	}
	microohm_meters := a.convertFromBase(ElectricResistivityMicroohmMeter)
	a.microohm_metersLazy = &microohm_meters
	return microohm_meters
}

// MilliohmMeters returns the ElectricResistivity value in MilliohmMeters.
//
// 
func (a *ElectricResistivity) MilliohmMeters() float64 {
	if a.milliohm_metersLazy != nil {
		return *a.milliohm_metersLazy
	}
	milliohm_meters := a.convertFromBase(ElectricResistivityMilliohmMeter)
	a.milliohm_metersLazy = &milliohm_meters
	return milliohm_meters
}

// KiloohmMeters returns the ElectricResistivity value in KiloohmMeters.
//
// 
func (a *ElectricResistivity) KiloohmMeters() float64 {
	if a.kiloohm_metersLazy != nil {
		return *a.kiloohm_metersLazy
	}
	kiloohm_meters := a.convertFromBase(ElectricResistivityKiloohmMeter)
	a.kiloohm_metersLazy = &kiloohm_meters
	return kiloohm_meters
}

// MegaohmMeters returns the ElectricResistivity value in MegaohmMeters.
//
// 
func (a *ElectricResistivity) MegaohmMeters() float64 {
	if a.megaohm_metersLazy != nil {
		return *a.megaohm_metersLazy
	}
	megaohm_meters := a.convertFromBase(ElectricResistivityMegaohmMeter)
	a.megaohm_metersLazy = &megaohm_meters
	return megaohm_meters
}

// PicoohmsCentimeter returns the ElectricResistivity value in PicoohmsCentimeter.
//
// 
func (a *ElectricResistivity) PicoohmsCentimeter() float64 {
	if a.picoohms_centimeterLazy != nil {
		return *a.picoohms_centimeterLazy
	}
	picoohms_centimeter := a.convertFromBase(ElectricResistivityPicoohmCentimeter)
	a.picoohms_centimeterLazy = &picoohms_centimeter
	return picoohms_centimeter
}

// NanoohmsCentimeter returns the ElectricResistivity value in NanoohmsCentimeter.
//
// 
func (a *ElectricResistivity) NanoohmsCentimeter() float64 {
	if a.nanoohms_centimeterLazy != nil {
		return *a.nanoohms_centimeterLazy
	}
	nanoohms_centimeter := a.convertFromBase(ElectricResistivityNanoohmCentimeter)
	a.nanoohms_centimeterLazy = &nanoohms_centimeter
	return nanoohms_centimeter
}

// MicroohmsCentimeter returns the ElectricResistivity value in MicroohmsCentimeter.
//
// 
func (a *ElectricResistivity) MicroohmsCentimeter() float64 {
	if a.microohms_centimeterLazy != nil {
		return *a.microohms_centimeterLazy
	}
	microohms_centimeter := a.convertFromBase(ElectricResistivityMicroohmCentimeter)
	a.microohms_centimeterLazy = &microohms_centimeter
	return microohms_centimeter
}

// MilliohmsCentimeter returns the ElectricResistivity value in MilliohmsCentimeter.
//
// 
func (a *ElectricResistivity) MilliohmsCentimeter() float64 {
	if a.milliohms_centimeterLazy != nil {
		return *a.milliohms_centimeterLazy
	}
	milliohms_centimeter := a.convertFromBase(ElectricResistivityMilliohmCentimeter)
	a.milliohms_centimeterLazy = &milliohms_centimeter
	return milliohms_centimeter
}

// KiloohmsCentimeter returns the ElectricResistivity value in KiloohmsCentimeter.
//
// 
func (a *ElectricResistivity) KiloohmsCentimeter() float64 {
	if a.kiloohms_centimeterLazy != nil {
		return *a.kiloohms_centimeterLazy
	}
	kiloohms_centimeter := a.convertFromBase(ElectricResistivityKiloohmCentimeter)
	a.kiloohms_centimeterLazy = &kiloohms_centimeter
	return kiloohms_centimeter
}

// MegaohmsCentimeter returns the ElectricResistivity value in MegaohmsCentimeter.
//
// 
func (a *ElectricResistivity) MegaohmsCentimeter() float64 {
	if a.megaohms_centimeterLazy != nil {
		return *a.megaohms_centimeterLazy
	}
	megaohms_centimeter := a.convertFromBase(ElectricResistivityMegaohmCentimeter)
	a.megaohms_centimeterLazy = &megaohms_centimeter
	return megaohms_centimeter
}


// ToDto creates a ElectricResistivityDto representation from the ElectricResistivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by OhmMeter by default.
func (a *ElectricResistivity) ToDto(holdInUnit *ElectricResistivityUnits) ElectricResistivityDto {
	if holdInUnit == nil {
		defaultUnit := ElectricResistivityOhmMeter // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricResistivityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricResistivity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by OhmMeter by default.
func (a *ElectricResistivity) ToDtoJSON(holdInUnit *ElectricResistivityUnits) ([]byte, error) {
	// Convert to ElectricResistivityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricResistivity to a specific unit value.
// The function uses the provided unit type (ElectricResistivityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricResistivity) Convert(toUnit ElectricResistivityUnits) float64 {
	switch toUnit { 
    case ElectricResistivityOhmMeter:
		return a.OhmMeters()
    case ElectricResistivityOhmCentimeter:
		return a.OhmsCentimeter()
    case ElectricResistivityPicoohmMeter:
		return a.PicoohmMeters()
    case ElectricResistivityNanoohmMeter:
		return a.NanoohmMeters()
    case ElectricResistivityMicroohmMeter:
		return a.MicroohmMeters()
    case ElectricResistivityMilliohmMeter:
		return a.MilliohmMeters()
    case ElectricResistivityKiloohmMeter:
		return a.KiloohmMeters()
    case ElectricResistivityMegaohmMeter:
		return a.MegaohmMeters()
    case ElectricResistivityPicoohmCentimeter:
		return a.PicoohmsCentimeter()
    case ElectricResistivityNanoohmCentimeter:
		return a.NanoohmsCentimeter()
    case ElectricResistivityMicroohmCentimeter:
		return a.MicroohmsCentimeter()
    case ElectricResistivityMilliohmCentimeter:
		return a.MilliohmsCentimeter()
    case ElectricResistivityKiloohmCentimeter:
		return a.KiloohmsCentimeter()
    case ElectricResistivityMegaohmCentimeter:
		return a.MegaohmsCentimeter()
	default:
		return math.NaN()
	}
}

func (a *ElectricResistivity) convertFromBase(toUnit ElectricResistivityUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricResistivityOhmMeter:
		return (value) 
	case ElectricResistivityOhmCentimeter:
		return (value * 100) 
	case ElectricResistivityPicoohmMeter:
		return ((value) / 1e-12) 
	case ElectricResistivityNanoohmMeter:
		return ((value) / 1e-09) 
	case ElectricResistivityMicroohmMeter:
		return ((value) / 1e-06) 
	case ElectricResistivityMilliohmMeter:
		return ((value) / 0.001) 
	case ElectricResistivityKiloohmMeter:
		return ((value) / 1000.0) 
	case ElectricResistivityMegaohmMeter:
		return ((value) / 1000000.0) 
	case ElectricResistivityPicoohmCentimeter:
		return ((value * 100) / 1e-12) 
	case ElectricResistivityNanoohmCentimeter:
		return ((value * 100) / 1e-09) 
	case ElectricResistivityMicroohmCentimeter:
		return ((value * 100) / 1e-06) 
	case ElectricResistivityMilliohmCentimeter:
		return ((value * 100) / 0.001) 
	case ElectricResistivityKiloohmCentimeter:
		return ((value * 100) / 1000.0) 
	case ElectricResistivityMegaohmCentimeter:
		return ((value * 100) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricResistivity) convertToBase(value float64, fromUnit ElectricResistivityUnits) float64 {
	switch fromUnit { 
	case ElectricResistivityOhmMeter:
		return (value) 
	case ElectricResistivityOhmCentimeter:
		return (value / 100) 
	case ElectricResistivityPicoohmMeter:
		return ((value) * 1e-12) 
	case ElectricResistivityNanoohmMeter:
		return ((value) * 1e-09) 
	case ElectricResistivityMicroohmMeter:
		return ((value) * 1e-06) 
	case ElectricResistivityMilliohmMeter:
		return ((value) * 0.001) 
	case ElectricResistivityKiloohmMeter:
		return ((value) * 1000.0) 
	case ElectricResistivityMegaohmMeter:
		return ((value) * 1000000.0) 
	case ElectricResistivityPicoohmCentimeter:
		return ((value / 100) * 1e-12) 
	case ElectricResistivityNanoohmCentimeter:
		return ((value / 100) * 1e-09) 
	case ElectricResistivityMicroohmCentimeter:
		return ((value / 100) * 1e-06) 
	case ElectricResistivityMilliohmCentimeter:
		return ((value / 100) * 0.001) 
	case ElectricResistivityKiloohmCentimeter:
		return ((value / 100) * 1000.0) 
	case ElectricResistivityMegaohmCentimeter:
		return ((value / 100) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricResistivity in the default unit (OhmMeter),
// formatted to two decimal places.
func (a ElectricResistivity) String() string {
	return a.ToString(ElectricResistivityOhmMeter, 2)
}

// ToString formats the ElectricResistivity value as a string with the specified unit and fractional digits.
// It converts the ElectricResistivity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricResistivity value will be converted (e.g., OhmMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricResistivity with the unit abbreviation.
func (a *ElectricResistivity) ToString(unit ElectricResistivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricResistivityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricResistivityAbbreviation(unit))
}

// Equals checks if the given ElectricResistivity is equal to the current ElectricResistivity.
//
// Parameters:
//    other: The ElectricResistivity to compare against.
//
// Returns:
//    bool: Returns true if both ElectricResistivity are equal, false otherwise.
func (a *ElectricResistivity) Equals(other *ElectricResistivity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricResistivity with another ElectricResistivity.
// It returns -1 if the current ElectricResistivity is less than the other ElectricResistivity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricResistivity to compare against.
//
// Returns:
//    int: -1 if the current ElectricResistivity is less, 1 if greater, and 0 if equal.
func (a *ElectricResistivity) CompareTo(other *ElectricResistivity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricResistivity to the current ElectricResistivity and returns the result.
// The result is a new ElectricResistivity instance with the sum of the values.
//
// Parameters:
//    other: The ElectricResistivity to add to the current ElectricResistivity.
//
// Returns:
//    *ElectricResistivity: A new ElectricResistivity instance representing the sum of both ElectricResistivity.
func (a *ElectricResistivity) Add(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricResistivity from the current ElectricResistivity and returns the result.
// The result is a new ElectricResistivity instance with the difference of the values.
//
// Parameters:
//    other: The ElectricResistivity to subtract from the current ElectricResistivity.
//
// Returns:
//    *ElectricResistivity: A new ElectricResistivity instance representing the difference of both ElectricResistivity.
func (a *ElectricResistivity) Subtract(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricResistivity by the given ElectricResistivity and returns the result.
// The result is a new ElectricResistivity instance with the product of the values.
//
// Parameters:
//    other: The ElectricResistivity to multiply with the current ElectricResistivity.
//
// Returns:
//    *ElectricResistivity: A new ElectricResistivity instance representing the product of both ElectricResistivity.
func (a *ElectricResistivity) Multiply(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricResistivity by the given ElectricResistivity and returns the result.
// The result is a new ElectricResistivity instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricResistivity to divide the current ElectricResistivity by.
//
// Returns:
//    *ElectricResistivity: A new ElectricResistivity instance representing the quotient of both ElectricResistivity.
func (a *ElectricResistivity) Divide(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value / other.BaseValue()}
}

// GetElectricResistivityAbbreviation gets the unit abbreviation.
func GetElectricResistivityAbbreviation(unit ElectricResistivityUnits) string {
	switch unit { 
	case ElectricResistivityOhmMeter:
		return "Ω·m" 
	case ElectricResistivityOhmCentimeter:
		return "Ω·cm" 
	case ElectricResistivityPicoohmMeter:
		return "pΩ·m" 
	case ElectricResistivityNanoohmMeter:
		return "nΩ·m" 
	case ElectricResistivityMicroohmMeter:
		return "μΩ·m" 
	case ElectricResistivityMilliohmMeter:
		return "mΩ·m" 
	case ElectricResistivityKiloohmMeter:
		return "kΩ·m" 
	case ElectricResistivityMegaohmMeter:
		return "MΩ·m" 
	case ElectricResistivityPicoohmCentimeter:
		return "pΩ·cm" 
	case ElectricResistivityNanoohmCentimeter:
		return "nΩ·cm" 
	case ElectricResistivityMicroohmCentimeter:
		return "μΩ·cm" 
	case ElectricResistivityMilliohmCentimeter:
		return "mΩ·cm" 
	case ElectricResistivityKiloohmCentimeter:
		return "kΩ·cm" 
	case ElectricResistivityMegaohmCentimeter:
		return "MΩ·cm" 
	default:
		return ""
	}
}