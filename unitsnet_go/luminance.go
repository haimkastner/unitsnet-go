package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LuminanceUnits enumeration
type LuminanceUnits string

const (
    
        // 
        LuminanceCandelaPerSquareMeter LuminanceUnits = "CandelaPerSquareMeter"
        // 
        LuminanceCandelaPerSquareFoot LuminanceUnits = "CandelaPerSquareFoot"
        // 
        LuminanceCandelaPerSquareInch LuminanceUnits = "CandelaPerSquareInch"
        // 
        LuminanceNit LuminanceUnits = "Nit"
        // 
        LuminanceNanocandelaPerSquareMeter LuminanceUnits = "NanocandelaPerSquareMeter"
        // 
        LuminanceMicrocandelaPerSquareMeter LuminanceUnits = "MicrocandelaPerSquareMeter"
        // 
        LuminanceMillicandelaPerSquareMeter LuminanceUnits = "MillicandelaPerSquareMeter"
        // 
        LuminanceCenticandelaPerSquareMeter LuminanceUnits = "CenticandelaPerSquareMeter"
        // 
        LuminanceDecicandelaPerSquareMeter LuminanceUnits = "DecicandelaPerSquareMeter"
        // 
        LuminanceKilocandelaPerSquareMeter LuminanceUnits = "KilocandelaPerSquareMeter"
)

// LuminanceDto represents an Luminance
type LuminanceDto struct {
	Value float64
	Unit  LuminanceUnits
}

// LuminanceDtoFactory struct to group related functions
type LuminanceDtoFactory struct{}

func (udf LuminanceDtoFactory) FromJSON(data []byte) (*LuminanceDto, error) {
	a := LuminanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LuminanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Luminance struct
type Luminance struct {
	value       float64
    
    candelas_per_square_meterLazy *float64 
    candelas_per_square_footLazy *float64 
    candelas_per_square_inchLazy *float64 
    nitsLazy *float64 
    nanocandelas_per_square_meterLazy *float64 
    microcandelas_per_square_meterLazy *float64 
    millicandelas_per_square_meterLazy *float64 
    centicandelas_per_square_meterLazy *float64 
    decicandelas_per_square_meterLazy *float64 
    kilocandelas_per_square_meterLazy *float64 
}

// LuminanceFactory struct to group related functions
type LuminanceFactory struct{}

func (uf LuminanceFactory) CreateLuminance(value float64, unit LuminanceUnits) (*Luminance, error) {
	return newLuminance(value, unit)
}

func (uf LuminanceFactory) FromDto(dto LuminanceDto) (*Luminance, error) {
	return newLuminance(dto.Value, dto.Unit)
}

func (uf LuminanceFactory) FromDtoJSON(data []byte) (*Luminance, error) {
	unitDto, err := LuminanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LuminanceFactory{}.FromDto(*unitDto)
}


// FromCandelaPerSquareMeter creates a new Luminance instance from CandelaPerSquareMeter.
func (uf LuminanceFactory) FromCandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareMeter)
}

// FromCandelaPerSquareFoot creates a new Luminance instance from CandelaPerSquareFoot.
func (uf LuminanceFactory) FromCandelasPerSquareFoot(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareFoot)
}

// FromCandelaPerSquareInch creates a new Luminance instance from CandelaPerSquareInch.
func (uf LuminanceFactory) FromCandelasPerSquareInch(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCandelaPerSquareInch)
}

// FromNit creates a new Luminance instance from Nit.
func (uf LuminanceFactory) FromNits(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceNit)
}

// FromNanocandelaPerSquareMeter creates a new Luminance instance from NanocandelaPerSquareMeter.
func (uf LuminanceFactory) FromNanocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceNanocandelaPerSquareMeter)
}

// FromMicrocandelaPerSquareMeter creates a new Luminance instance from MicrocandelaPerSquareMeter.
func (uf LuminanceFactory) FromMicrocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceMicrocandelaPerSquareMeter)
}

// FromMillicandelaPerSquareMeter creates a new Luminance instance from MillicandelaPerSquareMeter.
func (uf LuminanceFactory) FromMillicandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceMillicandelaPerSquareMeter)
}

// FromCenticandelaPerSquareMeter creates a new Luminance instance from CenticandelaPerSquareMeter.
func (uf LuminanceFactory) FromCenticandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceCenticandelaPerSquareMeter)
}

// FromDecicandelaPerSquareMeter creates a new Luminance instance from DecicandelaPerSquareMeter.
func (uf LuminanceFactory) FromDecicandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceDecicandelaPerSquareMeter)
}

// FromKilocandelaPerSquareMeter creates a new Luminance instance from KilocandelaPerSquareMeter.
func (uf LuminanceFactory) FromKilocandelasPerSquareMeter(value float64) (*Luminance, error) {
	return newLuminance(value, LuminanceKilocandelaPerSquareMeter)
}




// newLuminance creates a new Luminance.
func newLuminance(value float64, fromUnit LuminanceUnits) (*Luminance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Luminance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Luminance in CandelaPerSquareMeter.
func (a *Luminance) BaseValue() float64 {
	return a.value
}


// CandelaPerSquareMeter returns the value in CandelaPerSquareMeter.
func (a *Luminance) CandelasPerSquareMeter() float64 {
	if a.candelas_per_square_meterLazy != nil {
		return *a.candelas_per_square_meterLazy
	}
	candelas_per_square_meter := a.convertFromBase(LuminanceCandelaPerSquareMeter)
	a.candelas_per_square_meterLazy = &candelas_per_square_meter
	return candelas_per_square_meter
}

// CandelaPerSquareFoot returns the value in CandelaPerSquareFoot.
func (a *Luminance) CandelasPerSquareFoot() float64 {
	if a.candelas_per_square_footLazy != nil {
		return *a.candelas_per_square_footLazy
	}
	candelas_per_square_foot := a.convertFromBase(LuminanceCandelaPerSquareFoot)
	a.candelas_per_square_footLazy = &candelas_per_square_foot
	return candelas_per_square_foot
}

// CandelaPerSquareInch returns the value in CandelaPerSquareInch.
func (a *Luminance) CandelasPerSquareInch() float64 {
	if a.candelas_per_square_inchLazy != nil {
		return *a.candelas_per_square_inchLazy
	}
	candelas_per_square_inch := a.convertFromBase(LuminanceCandelaPerSquareInch)
	a.candelas_per_square_inchLazy = &candelas_per_square_inch
	return candelas_per_square_inch
}

// Nit returns the value in Nit.
func (a *Luminance) Nits() float64 {
	if a.nitsLazy != nil {
		return *a.nitsLazy
	}
	nits := a.convertFromBase(LuminanceNit)
	a.nitsLazy = &nits
	return nits
}

// NanocandelaPerSquareMeter returns the value in NanocandelaPerSquareMeter.
func (a *Luminance) NanocandelasPerSquareMeter() float64 {
	if a.nanocandelas_per_square_meterLazy != nil {
		return *a.nanocandelas_per_square_meterLazy
	}
	nanocandelas_per_square_meter := a.convertFromBase(LuminanceNanocandelaPerSquareMeter)
	a.nanocandelas_per_square_meterLazy = &nanocandelas_per_square_meter
	return nanocandelas_per_square_meter
}

// MicrocandelaPerSquareMeter returns the value in MicrocandelaPerSquareMeter.
func (a *Luminance) MicrocandelasPerSquareMeter() float64 {
	if a.microcandelas_per_square_meterLazy != nil {
		return *a.microcandelas_per_square_meterLazy
	}
	microcandelas_per_square_meter := a.convertFromBase(LuminanceMicrocandelaPerSquareMeter)
	a.microcandelas_per_square_meterLazy = &microcandelas_per_square_meter
	return microcandelas_per_square_meter
}

// MillicandelaPerSquareMeter returns the value in MillicandelaPerSquareMeter.
func (a *Luminance) MillicandelasPerSquareMeter() float64 {
	if a.millicandelas_per_square_meterLazy != nil {
		return *a.millicandelas_per_square_meterLazy
	}
	millicandelas_per_square_meter := a.convertFromBase(LuminanceMillicandelaPerSquareMeter)
	a.millicandelas_per_square_meterLazy = &millicandelas_per_square_meter
	return millicandelas_per_square_meter
}

// CenticandelaPerSquareMeter returns the value in CenticandelaPerSquareMeter.
func (a *Luminance) CenticandelasPerSquareMeter() float64 {
	if a.centicandelas_per_square_meterLazy != nil {
		return *a.centicandelas_per_square_meterLazy
	}
	centicandelas_per_square_meter := a.convertFromBase(LuminanceCenticandelaPerSquareMeter)
	a.centicandelas_per_square_meterLazy = &centicandelas_per_square_meter
	return centicandelas_per_square_meter
}

// DecicandelaPerSquareMeter returns the value in DecicandelaPerSquareMeter.
func (a *Luminance) DecicandelasPerSquareMeter() float64 {
	if a.decicandelas_per_square_meterLazy != nil {
		return *a.decicandelas_per_square_meterLazy
	}
	decicandelas_per_square_meter := a.convertFromBase(LuminanceDecicandelaPerSquareMeter)
	a.decicandelas_per_square_meterLazy = &decicandelas_per_square_meter
	return decicandelas_per_square_meter
}

// KilocandelaPerSquareMeter returns the value in KilocandelaPerSquareMeter.
func (a *Luminance) KilocandelasPerSquareMeter() float64 {
	if a.kilocandelas_per_square_meterLazy != nil {
		return *a.kilocandelas_per_square_meterLazy
	}
	kilocandelas_per_square_meter := a.convertFromBase(LuminanceKilocandelaPerSquareMeter)
	a.kilocandelas_per_square_meterLazy = &kilocandelas_per_square_meter
	return kilocandelas_per_square_meter
}


// ToDto creates an LuminanceDto representation.
func (a *Luminance) ToDto(holdInUnit *LuminanceUnits) LuminanceDto {
	if holdInUnit == nil {
		defaultUnit := LuminanceCandelaPerSquareMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LuminanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LuminanceDto representation.
func (a *Luminance) ToDtoJSON(holdInUnit *LuminanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Luminance to a specific unit value.
func (a *Luminance) Convert(toUnit LuminanceUnits) float64 {
	switch toUnit { 
    case LuminanceCandelaPerSquareMeter:
		return a.CandelasPerSquareMeter()
    case LuminanceCandelaPerSquareFoot:
		return a.CandelasPerSquareFoot()
    case LuminanceCandelaPerSquareInch:
		return a.CandelasPerSquareInch()
    case LuminanceNit:
		return a.Nits()
    case LuminanceNanocandelaPerSquareMeter:
		return a.NanocandelasPerSquareMeter()
    case LuminanceMicrocandelaPerSquareMeter:
		return a.MicrocandelasPerSquareMeter()
    case LuminanceMillicandelaPerSquareMeter:
		return a.MillicandelasPerSquareMeter()
    case LuminanceCenticandelaPerSquareMeter:
		return a.CenticandelasPerSquareMeter()
    case LuminanceDecicandelaPerSquareMeter:
		return a.DecicandelasPerSquareMeter()
    case LuminanceKilocandelaPerSquareMeter:
		return a.KilocandelasPerSquareMeter()
	default:
		return 0
	}
}

func (a *Luminance) convertFromBase(toUnit LuminanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminanceCandelaPerSquareMeter:
		return (value) 
	case LuminanceCandelaPerSquareFoot:
		return (value/ 1.07639e1) 
	case LuminanceCandelaPerSquareInch:
		return (value/ 1.5500031e3) 
	case LuminanceNit:
		return (value) 
	case LuminanceNanocandelaPerSquareMeter:
		return ((value) / 1e-09) 
	case LuminanceMicrocandelaPerSquareMeter:
		return ((value) / 1e-06) 
	case LuminanceMillicandelaPerSquareMeter:
		return ((value) / 0.001) 
	case LuminanceCenticandelaPerSquareMeter:
		return ((value) / 0.01) 
	case LuminanceDecicandelaPerSquareMeter:
		return ((value) / 0.1) 
	case LuminanceKilocandelaPerSquareMeter:
		return ((value) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *Luminance) convertToBase(value float64, fromUnit LuminanceUnits) float64 {
	switch fromUnit { 
	case LuminanceCandelaPerSquareMeter:
		return (value) 
	case LuminanceCandelaPerSquareFoot:
		return (value* 1.07639e1) 
	case LuminanceCandelaPerSquareInch:
		return (value* 1.5500031e3) 
	case LuminanceNit:
		return (value) 
	case LuminanceNanocandelaPerSquareMeter:
		return ((value) * 1e-09) 
	case LuminanceMicrocandelaPerSquareMeter:
		return ((value) * 1e-06) 
	case LuminanceMillicandelaPerSquareMeter:
		return ((value) * 0.001) 
	case LuminanceCenticandelaPerSquareMeter:
		return ((value) * 0.01) 
	case LuminanceDecicandelaPerSquareMeter:
		return ((value) * 0.1) 
	case LuminanceKilocandelaPerSquareMeter:
		return ((value) * 1000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Luminance) String() string {
	return a.ToString(LuminanceCandelaPerSquareMeter, 2)
}

// ToString formats the Luminance to string.
// fractionalDigits -1 for not mention
func (a *Luminance) ToString(unit LuminanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Luminance) getUnitAbbreviation(unit LuminanceUnits) string {
	switch unit { 
	case LuminanceCandelaPerSquareMeter:
		return "Cd/m²" 
	case LuminanceCandelaPerSquareFoot:
		return "Cd/ft²" 
	case LuminanceCandelaPerSquareInch:
		return "Cd/in²" 
	case LuminanceNit:
		return "nt" 
	case LuminanceNanocandelaPerSquareMeter:
		return "nCd/m²" 
	case LuminanceMicrocandelaPerSquareMeter:
		return "μCd/m²" 
	case LuminanceMillicandelaPerSquareMeter:
		return "mCd/m²" 
	case LuminanceCenticandelaPerSquareMeter:
		return "cCd/m²" 
	case LuminanceDecicandelaPerSquareMeter:
		return "dCd/m²" 
	case LuminanceKilocandelaPerSquareMeter:
		return "kCd/m²" 
	default:
		return ""
	}
}

// Check if the given Luminance are equals to the current Luminance
func (a *Luminance) Equals(other *Luminance) bool {
	return a.value == other.BaseValue()
}

// Check if the given Luminance are equals to the current Luminance
func (a *Luminance) CompareTo(other *Luminance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Luminance to the current Luminance.
func (a *Luminance) Add(other *Luminance) *Luminance {
	return &Luminance{value: a.value + other.BaseValue()}
}

// Subtract the given Luminance to the current Luminance.
func (a *Luminance) Subtract(other *Luminance) *Luminance {
	return &Luminance{value: a.value - other.BaseValue()}
}

// Multiply the given Luminance to the current Luminance.
func (a *Luminance) Multiply(other *Luminance) *Luminance {
	return &Luminance{value: a.value * other.BaseValue()}
}

// Divide the given Luminance to the current Luminance.
func (a *Luminance) Divide(other *Luminance) *Luminance {
	return &Luminance{value: a.value / other.BaseValue()}
}