// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolarityUnits enumeration
type MolarityUnits string

const (
    
        // 
        MolarityMolePerCubicMeter MolarityUnits = "MolePerCubicMeter"
        // 
        MolarityMolePerLiter MolarityUnits = "MolePerLiter"
        // 
        MolarityPoundMolePerCubicFoot MolarityUnits = "PoundMolePerCubicFoot"
        // 
        MolarityKilomolePerCubicMeter MolarityUnits = "KilomolePerCubicMeter"
        // 
        MolarityFemtomolePerLiter MolarityUnits = "FemtomolePerLiter"
        // 
        MolarityPicomolePerLiter MolarityUnits = "PicomolePerLiter"
        // 
        MolarityNanomolePerLiter MolarityUnits = "NanomolePerLiter"
        // 
        MolarityMicromolePerLiter MolarityUnits = "MicromolePerLiter"
        // 
        MolarityMillimolePerLiter MolarityUnits = "MillimolePerLiter"
        // 
        MolarityCentimolePerLiter MolarityUnits = "CentimolePerLiter"
        // 
        MolarityDecimolePerLiter MolarityUnits = "DecimolePerLiter"
)

// MolarityDto represents an Molarity
type MolarityDto struct {
	Value float64
	Unit  MolarityUnits
}

// MolarityDtoFactory struct to group related functions
type MolarityDtoFactory struct{}

func (udf MolarityDtoFactory) FromJSON(data []byte) (*MolarityDto, error) {
	a := MolarityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolarityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Molarity struct
type Molarity struct {
	value       float64
    
    moles_per_cubic_meterLazy *float64 
    moles_per_literLazy *float64 
    pound_moles_per_cubic_footLazy *float64 
    kilomoles_per_cubic_meterLazy *float64 
    femtomoles_per_literLazy *float64 
    picomoles_per_literLazy *float64 
    nanomoles_per_literLazy *float64 
    micromoles_per_literLazy *float64 
    millimoles_per_literLazy *float64 
    centimoles_per_literLazy *float64 
    decimoles_per_literLazy *float64 
}

// MolarityFactory struct to group related functions
type MolarityFactory struct{}

func (uf MolarityFactory) CreateMolarity(value float64, unit MolarityUnits) (*Molarity, error) {
	return newMolarity(value, unit)
}

func (uf MolarityFactory) FromDto(dto MolarityDto) (*Molarity, error) {
	return newMolarity(dto.Value, dto.Unit)
}

func (uf MolarityFactory) FromDtoJSON(data []byte) (*Molarity, error) {
	unitDto, err := MolarityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolarityFactory{}.FromDto(*unitDto)
}


// FromMolePerCubicMeter creates a new Molarity instance from MolePerCubicMeter.
func (uf MolarityFactory) FromMolesPerCubicMeter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMolePerCubicMeter)
}

// FromMolePerLiter creates a new Molarity instance from MolePerLiter.
func (uf MolarityFactory) FromMolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMolePerLiter)
}

// FromPoundMolePerCubicFoot creates a new Molarity instance from PoundMolePerCubicFoot.
func (uf MolarityFactory) FromPoundMolesPerCubicFoot(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityPoundMolePerCubicFoot)
}

// FromKilomolePerCubicMeter creates a new Molarity instance from KilomolePerCubicMeter.
func (uf MolarityFactory) FromKilomolesPerCubicMeter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityKilomolePerCubicMeter)
}

// FromFemtomolePerLiter creates a new Molarity instance from FemtomolePerLiter.
func (uf MolarityFactory) FromFemtomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityFemtomolePerLiter)
}

// FromPicomolePerLiter creates a new Molarity instance from PicomolePerLiter.
func (uf MolarityFactory) FromPicomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityPicomolePerLiter)
}

// FromNanomolePerLiter creates a new Molarity instance from NanomolePerLiter.
func (uf MolarityFactory) FromNanomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityNanomolePerLiter)
}

// FromMicromolePerLiter creates a new Molarity instance from MicromolePerLiter.
func (uf MolarityFactory) FromMicromolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMicromolePerLiter)
}

// FromMillimolePerLiter creates a new Molarity instance from MillimolePerLiter.
func (uf MolarityFactory) FromMillimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMillimolePerLiter)
}

// FromCentimolePerLiter creates a new Molarity instance from CentimolePerLiter.
func (uf MolarityFactory) FromCentimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityCentimolePerLiter)
}

// FromDecimolePerLiter creates a new Molarity instance from DecimolePerLiter.
func (uf MolarityFactory) FromDecimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityDecimolePerLiter)
}




// newMolarity creates a new Molarity.
func newMolarity(value float64, fromUnit MolarityUnits) (*Molarity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Molarity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Molarity in MolePerCubicMeter.
func (a *Molarity) BaseValue() float64 {
	return a.value
}


// MolePerCubicMeter returns the value in MolePerCubicMeter.
func (a *Molarity) MolesPerCubicMeter() float64 {
	if a.moles_per_cubic_meterLazy != nil {
		return *a.moles_per_cubic_meterLazy
	}
	moles_per_cubic_meter := a.convertFromBase(MolarityMolePerCubicMeter)
	a.moles_per_cubic_meterLazy = &moles_per_cubic_meter
	return moles_per_cubic_meter
}

// MolePerLiter returns the value in MolePerLiter.
func (a *Molarity) MolesPerLiter() float64 {
	if a.moles_per_literLazy != nil {
		return *a.moles_per_literLazy
	}
	moles_per_liter := a.convertFromBase(MolarityMolePerLiter)
	a.moles_per_literLazy = &moles_per_liter
	return moles_per_liter
}

// PoundMolePerCubicFoot returns the value in PoundMolePerCubicFoot.
func (a *Molarity) PoundMolesPerCubicFoot() float64 {
	if a.pound_moles_per_cubic_footLazy != nil {
		return *a.pound_moles_per_cubic_footLazy
	}
	pound_moles_per_cubic_foot := a.convertFromBase(MolarityPoundMolePerCubicFoot)
	a.pound_moles_per_cubic_footLazy = &pound_moles_per_cubic_foot
	return pound_moles_per_cubic_foot
}

// KilomolePerCubicMeter returns the value in KilomolePerCubicMeter.
func (a *Molarity) KilomolesPerCubicMeter() float64 {
	if a.kilomoles_per_cubic_meterLazy != nil {
		return *a.kilomoles_per_cubic_meterLazy
	}
	kilomoles_per_cubic_meter := a.convertFromBase(MolarityKilomolePerCubicMeter)
	a.kilomoles_per_cubic_meterLazy = &kilomoles_per_cubic_meter
	return kilomoles_per_cubic_meter
}

// FemtomolePerLiter returns the value in FemtomolePerLiter.
func (a *Molarity) FemtomolesPerLiter() float64 {
	if a.femtomoles_per_literLazy != nil {
		return *a.femtomoles_per_literLazy
	}
	femtomoles_per_liter := a.convertFromBase(MolarityFemtomolePerLiter)
	a.femtomoles_per_literLazy = &femtomoles_per_liter
	return femtomoles_per_liter
}

// PicomolePerLiter returns the value in PicomolePerLiter.
func (a *Molarity) PicomolesPerLiter() float64 {
	if a.picomoles_per_literLazy != nil {
		return *a.picomoles_per_literLazy
	}
	picomoles_per_liter := a.convertFromBase(MolarityPicomolePerLiter)
	a.picomoles_per_literLazy = &picomoles_per_liter
	return picomoles_per_liter
}

// NanomolePerLiter returns the value in NanomolePerLiter.
func (a *Molarity) NanomolesPerLiter() float64 {
	if a.nanomoles_per_literLazy != nil {
		return *a.nanomoles_per_literLazy
	}
	nanomoles_per_liter := a.convertFromBase(MolarityNanomolePerLiter)
	a.nanomoles_per_literLazy = &nanomoles_per_liter
	return nanomoles_per_liter
}

// MicromolePerLiter returns the value in MicromolePerLiter.
func (a *Molarity) MicromolesPerLiter() float64 {
	if a.micromoles_per_literLazy != nil {
		return *a.micromoles_per_literLazy
	}
	micromoles_per_liter := a.convertFromBase(MolarityMicromolePerLiter)
	a.micromoles_per_literLazy = &micromoles_per_liter
	return micromoles_per_liter
}

// MillimolePerLiter returns the value in MillimolePerLiter.
func (a *Molarity) MillimolesPerLiter() float64 {
	if a.millimoles_per_literLazy != nil {
		return *a.millimoles_per_literLazy
	}
	millimoles_per_liter := a.convertFromBase(MolarityMillimolePerLiter)
	a.millimoles_per_literLazy = &millimoles_per_liter
	return millimoles_per_liter
}

// CentimolePerLiter returns the value in CentimolePerLiter.
func (a *Molarity) CentimolesPerLiter() float64 {
	if a.centimoles_per_literLazy != nil {
		return *a.centimoles_per_literLazy
	}
	centimoles_per_liter := a.convertFromBase(MolarityCentimolePerLiter)
	a.centimoles_per_literLazy = &centimoles_per_liter
	return centimoles_per_liter
}

// DecimolePerLiter returns the value in DecimolePerLiter.
func (a *Molarity) DecimolesPerLiter() float64 {
	if a.decimoles_per_literLazy != nil {
		return *a.decimoles_per_literLazy
	}
	decimoles_per_liter := a.convertFromBase(MolarityDecimolePerLiter)
	a.decimoles_per_literLazy = &decimoles_per_liter
	return decimoles_per_liter
}


// ToDto creates an MolarityDto representation.
func (a *Molarity) ToDto(holdInUnit *MolarityUnits) MolarityDto {
	if holdInUnit == nil {
		defaultUnit := MolarityMolePerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return MolarityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MolarityDto representation.
func (a *Molarity) ToDtoJSON(holdInUnit *MolarityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Molarity to a specific unit value.
func (a *Molarity) Convert(toUnit MolarityUnits) float64 {
	switch toUnit { 
    case MolarityMolePerCubicMeter:
		return a.MolesPerCubicMeter()
    case MolarityMolePerLiter:
		return a.MolesPerLiter()
    case MolarityPoundMolePerCubicFoot:
		return a.PoundMolesPerCubicFoot()
    case MolarityKilomolePerCubicMeter:
		return a.KilomolesPerCubicMeter()
    case MolarityFemtomolePerLiter:
		return a.FemtomolesPerLiter()
    case MolarityPicomolePerLiter:
		return a.PicomolesPerLiter()
    case MolarityNanomolePerLiter:
		return a.NanomolesPerLiter()
    case MolarityMicromolePerLiter:
		return a.MicromolesPerLiter()
    case MolarityMillimolePerLiter:
		return a.MillimolesPerLiter()
    case MolarityCentimolePerLiter:
		return a.CentimolesPerLiter()
    case MolarityDecimolePerLiter:
		return a.DecimolesPerLiter()
	default:
		return 0
	}
}

func (a *Molarity) convertFromBase(toUnit MolarityUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarityMolePerCubicMeter:
		return (value) 
	case MolarityMolePerLiter:
		return (value * 1e-3) 
	case MolarityPoundMolePerCubicFoot:
		return (value * 6.2427960576144611956325455827221e-5) 
	case MolarityKilomolePerCubicMeter:
		return ((value) / 1000.0) 
	case MolarityFemtomolePerLiter:
		return ((value * 1e-3) / 1e-15) 
	case MolarityPicomolePerLiter:
		return ((value * 1e-3) / 1e-12) 
	case MolarityNanomolePerLiter:
		return ((value * 1e-3) / 1e-09) 
	case MolarityMicromolePerLiter:
		return ((value * 1e-3) / 1e-06) 
	case MolarityMillimolePerLiter:
		return ((value * 1e-3) / 0.001) 
	case MolarityCentimolePerLiter:
		return ((value * 1e-3) / 0.01) 
	case MolarityDecimolePerLiter:
		return ((value * 1e-3) / 0.1) 
	default:
		return math.NaN()
	}
}

func (a *Molarity) convertToBase(value float64, fromUnit MolarityUnits) float64 {
	switch fromUnit { 
	case MolarityMolePerCubicMeter:
		return (value) 
	case MolarityMolePerLiter:
		return (value / 1e-3) 
	case MolarityPoundMolePerCubicFoot:
		return (value / 6.2427960576144611956325455827221e-5) 
	case MolarityKilomolePerCubicMeter:
		return ((value) * 1000.0) 
	case MolarityFemtomolePerLiter:
		return ((value / 1e-3) * 1e-15) 
	case MolarityPicomolePerLiter:
		return ((value / 1e-3) * 1e-12) 
	case MolarityNanomolePerLiter:
		return ((value / 1e-3) * 1e-09) 
	case MolarityMicromolePerLiter:
		return ((value / 1e-3) * 1e-06) 
	case MolarityMillimolePerLiter:
		return ((value / 1e-3) * 0.001) 
	case MolarityCentimolePerLiter:
		return ((value / 1e-3) * 0.01) 
	case MolarityDecimolePerLiter:
		return ((value / 1e-3) * 0.1) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Molarity) String() string {
	return a.ToString(MolarityMolePerCubicMeter, 2)
}

// ToString formats the Molarity to string.
// fractionalDigits -1 for not mention
func (a *Molarity) ToString(unit MolarityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Molarity) getUnitAbbreviation(unit MolarityUnits) string {
	switch unit { 
	case MolarityMolePerCubicMeter:
		return "mol/m³" 
	case MolarityMolePerLiter:
		return "mol/L" 
	case MolarityPoundMolePerCubicFoot:
		return "lbmol/ft³" 
	case MolarityKilomolePerCubicMeter:
		return "kmol/m³" 
	case MolarityFemtomolePerLiter:
		return "fmol/L" 
	case MolarityPicomolePerLiter:
		return "pmol/L" 
	case MolarityNanomolePerLiter:
		return "nmol/L" 
	case MolarityMicromolePerLiter:
		return "μmol/L" 
	case MolarityMillimolePerLiter:
		return "mmol/L" 
	case MolarityCentimolePerLiter:
		return "cmol/L" 
	case MolarityDecimolePerLiter:
		return "dmol/L" 
	default:
		return ""
	}
}

// Check if the given Molarity are equals to the current Molarity
func (a *Molarity) Equals(other *Molarity) bool {
	return a.value == other.BaseValue()
}

// Check if the given Molarity are equals to the current Molarity
func (a *Molarity) CompareTo(other *Molarity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Molarity to the current Molarity.
func (a *Molarity) Add(other *Molarity) *Molarity {
	return &Molarity{value: a.value + other.BaseValue()}
}

// Subtract the given Molarity to the current Molarity.
func (a *Molarity) Subtract(other *Molarity) *Molarity {
	return &Molarity{value: a.value - other.BaseValue()}
}

// Multiply the given Molarity to the current Molarity.
func (a *Molarity) Multiply(other *Molarity) *Molarity {
	return &Molarity{value: a.value * other.BaseValue()}
}

// Divide the given Molarity to the current Molarity.
func (a *Molarity) Divide(other *Molarity) *Molarity {
	return &Molarity{value: a.value / other.BaseValue()}
}