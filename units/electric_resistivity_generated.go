// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricResistivityUnits enumeration
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

// ElectricResistivityDto represents an ElectricResistivity
type ElectricResistivityDto struct {
	Value float64
	Unit  ElectricResistivityUnits
}

// ElectricResistivityDtoFactory struct to group related functions
type ElectricResistivityDtoFactory struct{}

func (udf ElectricResistivityDtoFactory) FromJSON(data []byte) (*ElectricResistivityDto, error) {
	a := ElectricResistivityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a ElectricResistivityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// ElectricResistivity struct
type ElectricResistivity struct {
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

// ElectricResistivityFactory struct to group related functions
type ElectricResistivityFactory struct{}

func (uf ElectricResistivityFactory) CreateElectricResistivity(value float64, unit ElectricResistivityUnits) (*ElectricResistivity, error) {
	return newElectricResistivity(value, unit)
}

func (uf ElectricResistivityFactory) FromDto(dto ElectricResistivityDto) (*ElectricResistivity, error) {
	return newElectricResistivity(dto.Value, dto.Unit)
}

func (uf ElectricResistivityFactory) FromDtoJSON(data []byte) (*ElectricResistivity, error) {
	unitDto, err := ElectricResistivityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return ElectricResistivityFactory{}.FromDto(*unitDto)
}


// FromOhmMeter creates a new ElectricResistivity instance from OhmMeter.
func (uf ElectricResistivityFactory) FromOhmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityOhmMeter)
}

// FromOhmCentimeter creates a new ElectricResistivity instance from OhmCentimeter.
func (uf ElectricResistivityFactory) FromOhmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityOhmCentimeter)
}

// FromPicoohmMeter creates a new ElectricResistivity instance from PicoohmMeter.
func (uf ElectricResistivityFactory) FromPicoohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityPicoohmMeter)
}

// FromNanoohmMeter creates a new ElectricResistivity instance from NanoohmMeter.
func (uf ElectricResistivityFactory) FromNanoohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityNanoohmMeter)
}

// FromMicroohmMeter creates a new ElectricResistivity instance from MicroohmMeter.
func (uf ElectricResistivityFactory) FromMicroohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMicroohmMeter)
}

// FromMilliohmMeter creates a new ElectricResistivity instance from MilliohmMeter.
func (uf ElectricResistivityFactory) FromMilliohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMilliohmMeter)
}

// FromKiloohmMeter creates a new ElectricResistivity instance from KiloohmMeter.
func (uf ElectricResistivityFactory) FromKiloohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityKiloohmMeter)
}

// FromMegaohmMeter creates a new ElectricResistivity instance from MegaohmMeter.
func (uf ElectricResistivityFactory) FromMegaohmMeters(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMegaohmMeter)
}

// FromPicoohmCentimeter creates a new ElectricResistivity instance from PicoohmCentimeter.
func (uf ElectricResistivityFactory) FromPicoohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityPicoohmCentimeter)
}

// FromNanoohmCentimeter creates a new ElectricResistivity instance from NanoohmCentimeter.
func (uf ElectricResistivityFactory) FromNanoohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityNanoohmCentimeter)
}

// FromMicroohmCentimeter creates a new ElectricResistivity instance from MicroohmCentimeter.
func (uf ElectricResistivityFactory) FromMicroohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMicroohmCentimeter)
}

// FromMilliohmCentimeter creates a new ElectricResistivity instance from MilliohmCentimeter.
func (uf ElectricResistivityFactory) FromMilliohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMilliohmCentimeter)
}

// FromKiloohmCentimeter creates a new ElectricResistivity instance from KiloohmCentimeter.
func (uf ElectricResistivityFactory) FromKiloohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityKiloohmCentimeter)
}

// FromMegaohmCentimeter creates a new ElectricResistivity instance from MegaohmCentimeter.
func (uf ElectricResistivityFactory) FromMegaohmsCentimeter(value float64) (*ElectricResistivity, error) {
	return newElectricResistivity(value, ElectricResistivityMegaohmCentimeter)
}




// newElectricResistivity creates a new ElectricResistivity.
func newElectricResistivity(value float64, fromUnit ElectricResistivityUnits) (*ElectricResistivity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricResistivity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricResistivity in OhmMeter.
func (a *ElectricResistivity) BaseValue() float64 {
	return a.value
}


// OhmMeter returns the value in OhmMeter.
func (a *ElectricResistivity) OhmMeters() float64 {
	if a.ohm_metersLazy != nil {
		return *a.ohm_metersLazy
	}
	ohm_meters := a.convertFromBase(ElectricResistivityOhmMeter)
	a.ohm_metersLazy = &ohm_meters
	return ohm_meters
}

// OhmCentimeter returns the value in OhmCentimeter.
func (a *ElectricResistivity) OhmsCentimeter() float64 {
	if a.ohms_centimeterLazy != nil {
		return *a.ohms_centimeterLazy
	}
	ohms_centimeter := a.convertFromBase(ElectricResistivityOhmCentimeter)
	a.ohms_centimeterLazy = &ohms_centimeter
	return ohms_centimeter
}

// PicoohmMeter returns the value in PicoohmMeter.
func (a *ElectricResistivity) PicoohmMeters() float64 {
	if a.picoohm_metersLazy != nil {
		return *a.picoohm_metersLazy
	}
	picoohm_meters := a.convertFromBase(ElectricResistivityPicoohmMeter)
	a.picoohm_metersLazy = &picoohm_meters
	return picoohm_meters
}

// NanoohmMeter returns the value in NanoohmMeter.
func (a *ElectricResistivity) NanoohmMeters() float64 {
	if a.nanoohm_metersLazy != nil {
		return *a.nanoohm_metersLazy
	}
	nanoohm_meters := a.convertFromBase(ElectricResistivityNanoohmMeter)
	a.nanoohm_metersLazy = &nanoohm_meters
	return nanoohm_meters
}

// MicroohmMeter returns the value in MicroohmMeter.
func (a *ElectricResistivity) MicroohmMeters() float64 {
	if a.microohm_metersLazy != nil {
		return *a.microohm_metersLazy
	}
	microohm_meters := a.convertFromBase(ElectricResistivityMicroohmMeter)
	a.microohm_metersLazy = &microohm_meters
	return microohm_meters
}

// MilliohmMeter returns the value in MilliohmMeter.
func (a *ElectricResistivity) MilliohmMeters() float64 {
	if a.milliohm_metersLazy != nil {
		return *a.milliohm_metersLazy
	}
	milliohm_meters := a.convertFromBase(ElectricResistivityMilliohmMeter)
	a.milliohm_metersLazy = &milliohm_meters
	return milliohm_meters
}

// KiloohmMeter returns the value in KiloohmMeter.
func (a *ElectricResistivity) KiloohmMeters() float64 {
	if a.kiloohm_metersLazy != nil {
		return *a.kiloohm_metersLazy
	}
	kiloohm_meters := a.convertFromBase(ElectricResistivityKiloohmMeter)
	a.kiloohm_metersLazy = &kiloohm_meters
	return kiloohm_meters
}

// MegaohmMeter returns the value in MegaohmMeter.
func (a *ElectricResistivity) MegaohmMeters() float64 {
	if a.megaohm_metersLazy != nil {
		return *a.megaohm_metersLazy
	}
	megaohm_meters := a.convertFromBase(ElectricResistivityMegaohmMeter)
	a.megaohm_metersLazy = &megaohm_meters
	return megaohm_meters
}

// PicoohmCentimeter returns the value in PicoohmCentimeter.
func (a *ElectricResistivity) PicoohmsCentimeter() float64 {
	if a.picoohms_centimeterLazy != nil {
		return *a.picoohms_centimeterLazy
	}
	picoohms_centimeter := a.convertFromBase(ElectricResistivityPicoohmCentimeter)
	a.picoohms_centimeterLazy = &picoohms_centimeter
	return picoohms_centimeter
}

// NanoohmCentimeter returns the value in NanoohmCentimeter.
func (a *ElectricResistivity) NanoohmsCentimeter() float64 {
	if a.nanoohms_centimeterLazy != nil {
		return *a.nanoohms_centimeterLazy
	}
	nanoohms_centimeter := a.convertFromBase(ElectricResistivityNanoohmCentimeter)
	a.nanoohms_centimeterLazy = &nanoohms_centimeter
	return nanoohms_centimeter
}

// MicroohmCentimeter returns the value in MicroohmCentimeter.
func (a *ElectricResistivity) MicroohmsCentimeter() float64 {
	if a.microohms_centimeterLazy != nil {
		return *a.microohms_centimeterLazy
	}
	microohms_centimeter := a.convertFromBase(ElectricResistivityMicroohmCentimeter)
	a.microohms_centimeterLazy = &microohms_centimeter
	return microohms_centimeter
}

// MilliohmCentimeter returns the value in MilliohmCentimeter.
func (a *ElectricResistivity) MilliohmsCentimeter() float64 {
	if a.milliohms_centimeterLazy != nil {
		return *a.milliohms_centimeterLazy
	}
	milliohms_centimeter := a.convertFromBase(ElectricResistivityMilliohmCentimeter)
	a.milliohms_centimeterLazy = &milliohms_centimeter
	return milliohms_centimeter
}

// KiloohmCentimeter returns the value in KiloohmCentimeter.
func (a *ElectricResistivity) KiloohmsCentimeter() float64 {
	if a.kiloohms_centimeterLazy != nil {
		return *a.kiloohms_centimeterLazy
	}
	kiloohms_centimeter := a.convertFromBase(ElectricResistivityKiloohmCentimeter)
	a.kiloohms_centimeterLazy = &kiloohms_centimeter
	return kiloohms_centimeter
}

// MegaohmCentimeter returns the value in MegaohmCentimeter.
func (a *ElectricResistivity) MegaohmsCentimeter() float64 {
	if a.megaohms_centimeterLazy != nil {
		return *a.megaohms_centimeterLazy
	}
	megaohms_centimeter := a.convertFromBase(ElectricResistivityMegaohmCentimeter)
	a.megaohms_centimeterLazy = &megaohms_centimeter
	return megaohms_centimeter
}


// ToDto creates an ElectricResistivityDto representation.
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

// ToDtoJSON creates an ElectricResistivityDto representation.
func (a *ElectricResistivity) ToDtoJSON(holdInUnit *ElectricResistivityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts ElectricResistivity to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a ElectricResistivity) String() string {
	return a.ToString(ElectricResistivityOhmMeter, 2)
}

// ToString formats the ElectricResistivity to string.
// fractionalDigits -1 for not mention
func (a *ElectricResistivity) ToString(unit ElectricResistivityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *ElectricResistivity) getUnitAbbreviation(unit ElectricResistivityUnits) string {
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

// Check if the given ElectricResistivity are equals to the current ElectricResistivity
func (a *ElectricResistivity) Equals(other *ElectricResistivity) bool {
	return a.value == other.BaseValue()
}

// Check if the given ElectricResistivity are equals to the current ElectricResistivity
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

// Add the given ElectricResistivity to the current ElectricResistivity.
func (a *ElectricResistivity) Add(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value + other.BaseValue()}
}

// Subtract the given ElectricResistivity to the current ElectricResistivity.
func (a *ElectricResistivity) Subtract(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value - other.BaseValue()}
}

// Multiply the given ElectricResistivity to the current ElectricResistivity.
func (a *ElectricResistivity) Multiply(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value * other.BaseValue()}
}

// Divide the given ElectricResistivity to the current ElectricResistivity.
func (a *ElectricResistivity) Divide(other *ElectricResistivity) *ElectricResistivity {
	return &ElectricResistivity{value: a.value / other.BaseValue()}
}