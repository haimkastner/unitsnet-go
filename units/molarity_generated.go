// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolarityUnits defines various units of Molarity.
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

var internalMolarityUnitsMap = map[MolarityUnits]bool{
	
	MolarityMolePerCubicMeter: true,
	MolarityMolePerLiter: true,
	MolarityPoundMolePerCubicFoot: true,
	MolarityKilomolePerCubicMeter: true,
	MolarityFemtomolePerLiter: true,
	MolarityPicomolePerLiter: true,
	MolarityNanomolePerLiter: true,
	MolarityMicromolePerLiter: true,
	MolarityMillimolePerLiter: true,
	MolarityCentimolePerLiter: true,
	MolarityDecimolePerLiter: true,
}

// MolarityDto represents a Molarity measurement with a numerical value and its corresponding unit.
type MolarityDto struct {
    // Value is the numerical representation of the Molarity.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Molarity, as defined in the MolarityUnits enumeration.
	Unit  MolarityUnits `json:"unit" validate:"required,oneof=MolePerCubicMeter MolePerLiter PoundMolePerCubicFoot KilomolePerCubicMeter FemtomolePerLiter PicomolePerLiter NanomolePerLiter MicromolePerLiter MillimolePerLiter CentimolePerLiter DecimolePerLiter"`
}

// MolarityDtoFactory groups methods for creating and serializing MolarityDto objects.
type MolarityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolarityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolarityDtoFactory) FromJSON(data []byte) (*MolarityDto, error) {
	a := MolarityDto{}

    // Parse JSON into MolarityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MolarityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolarityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Molarity represents a measurement in a of Molarity.
//
// Molar concentration, also called molarity, amount concentration or substance concentration, is a measure of the concentration of a solute in a solution, or of any chemical species, in terms of amount of substance in a given volume.
type Molarity struct {
	// value is the base measurement stored internally.
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

// MolarityFactory groups methods for creating Molarity instances.
type MolarityFactory struct{}

// CreateMolarity creates a new Molarity instance from the given value and unit.
func (uf MolarityFactory) CreateMolarity(value float64, unit MolarityUnits) (*Molarity, error) {
	return newMolarity(value, unit)
}

// FromDto converts a MolarityDto to a Molarity instance.
func (uf MolarityFactory) FromDto(dto MolarityDto) (*Molarity, error) {
	return newMolarity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Molarity instance.
func (uf MolarityFactory) FromDtoJSON(data []byte) (*Molarity, error) {
	unitDto, err := MolarityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolarityDto from JSON: %w", err)
	}
	return MolarityFactory{}.FromDto(*unitDto)
}


// FromMolesPerCubicMeter creates a new Molarity instance from a value in MolesPerCubicMeter.
func (uf MolarityFactory) FromMolesPerCubicMeter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMolePerCubicMeter)
}

// FromMolesPerLiter creates a new Molarity instance from a value in MolesPerLiter.
func (uf MolarityFactory) FromMolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMolePerLiter)
}

// FromPoundMolesPerCubicFoot creates a new Molarity instance from a value in PoundMolesPerCubicFoot.
func (uf MolarityFactory) FromPoundMolesPerCubicFoot(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityPoundMolePerCubicFoot)
}

// FromKilomolesPerCubicMeter creates a new Molarity instance from a value in KilomolesPerCubicMeter.
func (uf MolarityFactory) FromKilomolesPerCubicMeter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityKilomolePerCubicMeter)
}

// FromFemtomolesPerLiter creates a new Molarity instance from a value in FemtomolesPerLiter.
func (uf MolarityFactory) FromFemtomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityFemtomolePerLiter)
}

// FromPicomolesPerLiter creates a new Molarity instance from a value in PicomolesPerLiter.
func (uf MolarityFactory) FromPicomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityPicomolePerLiter)
}

// FromNanomolesPerLiter creates a new Molarity instance from a value in NanomolesPerLiter.
func (uf MolarityFactory) FromNanomolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityNanomolePerLiter)
}

// FromMicromolesPerLiter creates a new Molarity instance from a value in MicromolesPerLiter.
func (uf MolarityFactory) FromMicromolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMicromolePerLiter)
}

// FromMillimolesPerLiter creates a new Molarity instance from a value in MillimolesPerLiter.
func (uf MolarityFactory) FromMillimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityMillimolePerLiter)
}

// FromCentimolesPerLiter creates a new Molarity instance from a value in CentimolesPerLiter.
func (uf MolarityFactory) FromCentimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityCentimolePerLiter)
}

// FromDecimolesPerLiter creates a new Molarity instance from a value in DecimolesPerLiter.
func (uf MolarityFactory) FromDecimolesPerLiter(value float64) (*Molarity, error) {
	return newMolarity(value, MolarityDecimolePerLiter)
}


// newMolarity creates a new Molarity.
func newMolarity(value float64, fromUnit MolarityUnits) (*Molarity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMolarityUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MolarityUnits", fromUnit)
	}
	a := &Molarity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Molarity in MolePerCubicMeter unit (the base/default quantity).
func (a *Molarity) BaseValue() float64 {
	return a.value
}


// MolesPerCubicMeter returns the Molarity value in MolesPerCubicMeter.
//
// 
func (a *Molarity) MolesPerCubicMeter() float64 {
	if a.moles_per_cubic_meterLazy != nil {
		return *a.moles_per_cubic_meterLazy
	}
	moles_per_cubic_meter := a.convertFromBase(MolarityMolePerCubicMeter)
	a.moles_per_cubic_meterLazy = &moles_per_cubic_meter
	return moles_per_cubic_meter
}

// MolesPerLiter returns the Molarity value in MolesPerLiter.
//
// 
func (a *Molarity) MolesPerLiter() float64 {
	if a.moles_per_literLazy != nil {
		return *a.moles_per_literLazy
	}
	moles_per_liter := a.convertFromBase(MolarityMolePerLiter)
	a.moles_per_literLazy = &moles_per_liter
	return moles_per_liter
}

// PoundMolesPerCubicFoot returns the Molarity value in PoundMolesPerCubicFoot.
//
// 
func (a *Molarity) PoundMolesPerCubicFoot() float64 {
	if a.pound_moles_per_cubic_footLazy != nil {
		return *a.pound_moles_per_cubic_footLazy
	}
	pound_moles_per_cubic_foot := a.convertFromBase(MolarityPoundMolePerCubicFoot)
	a.pound_moles_per_cubic_footLazy = &pound_moles_per_cubic_foot
	return pound_moles_per_cubic_foot
}

// KilomolesPerCubicMeter returns the Molarity value in KilomolesPerCubicMeter.
//
// 
func (a *Molarity) KilomolesPerCubicMeter() float64 {
	if a.kilomoles_per_cubic_meterLazy != nil {
		return *a.kilomoles_per_cubic_meterLazy
	}
	kilomoles_per_cubic_meter := a.convertFromBase(MolarityKilomolePerCubicMeter)
	a.kilomoles_per_cubic_meterLazy = &kilomoles_per_cubic_meter
	return kilomoles_per_cubic_meter
}

// FemtomolesPerLiter returns the Molarity value in FemtomolesPerLiter.
//
// 
func (a *Molarity) FemtomolesPerLiter() float64 {
	if a.femtomoles_per_literLazy != nil {
		return *a.femtomoles_per_literLazy
	}
	femtomoles_per_liter := a.convertFromBase(MolarityFemtomolePerLiter)
	a.femtomoles_per_literLazy = &femtomoles_per_liter
	return femtomoles_per_liter
}

// PicomolesPerLiter returns the Molarity value in PicomolesPerLiter.
//
// 
func (a *Molarity) PicomolesPerLiter() float64 {
	if a.picomoles_per_literLazy != nil {
		return *a.picomoles_per_literLazy
	}
	picomoles_per_liter := a.convertFromBase(MolarityPicomolePerLiter)
	a.picomoles_per_literLazy = &picomoles_per_liter
	return picomoles_per_liter
}

// NanomolesPerLiter returns the Molarity value in NanomolesPerLiter.
//
// 
func (a *Molarity) NanomolesPerLiter() float64 {
	if a.nanomoles_per_literLazy != nil {
		return *a.nanomoles_per_literLazy
	}
	nanomoles_per_liter := a.convertFromBase(MolarityNanomolePerLiter)
	a.nanomoles_per_literLazy = &nanomoles_per_liter
	return nanomoles_per_liter
}

// MicromolesPerLiter returns the Molarity value in MicromolesPerLiter.
//
// 
func (a *Molarity) MicromolesPerLiter() float64 {
	if a.micromoles_per_literLazy != nil {
		return *a.micromoles_per_literLazy
	}
	micromoles_per_liter := a.convertFromBase(MolarityMicromolePerLiter)
	a.micromoles_per_literLazy = &micromoles_per_liter
	return micromoles_per_liter
}

// MillimolesPerLiter returns the Molarity value in MillimolesPerLiter.
//
// 
func (a *Molarity) MillimolesPerLiter() float64 {
	if a.millimoles_per_literLazy != nil {
		return *a.millimoles_per_literLazy
	}
	millimoles_per_liter := a.convertFromBase(MolarityMillimolePerLiter)
	a.millimoles_per_literLazy = &millimoles_per_liter
	return millimoles_per_liter
}

// CentimolesPerLiter returns the Molarity value in CentimolesPerLiter.
//
// 
func (a *Molarity) CentimolesPerLiter() float64 {
	if a.centimoles_per_literLazy != nil {
		return *a.centimoles_per_literLazy
	}
	centimoles_per_liter := a.convertFromBase(MolarityCentimolePerLiter)
	a.centimoles_per_literLazy = &centimoles_per_liter
	return centimoles_per_liter
}

// DecimolesPerLiter returns the Molarity value in DecimolesPerLiter.
//
// 
func (a *Molarity) DecimolesPerLiter() float64 {
	if a.decimoles_per_literLazy != nil {
		return *a.decimoles_per_literLazy
	}
	decimoles_per_liter := a.convertFromBase(MolarityDecimolePerLiter)
	a.decimoles_per_literLazy = &decimoles_per_liter
	return decimoles_per_liter
}


// ToDto creates a MolarityDto representation from the Molarity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerCubicMeter by default.
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

// ToDtoJSON creates a JSON representation of the Molarity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by MolePerCubicMeter by default.
func (a *Molarity) ToDtoJSON(holdInUnit *MolarityUnits) ([]byte, error) {
	// Convert to MolarityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Molarity to a specific unit value.
// The function uses the provided unit type (MolarityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Molarity in the default unit (MolePerCubicMeter),
// formatted to two decimal places.
func (a Molarity) String() string {
	return a.ToString(MolarityMolePerCubicMeter, 2)
}

// ToString formats the Molarity value as a string with the specified unit and fractional digits.
// It converts the Molarity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Molarity value will be converted (e.g., MolePerCubicMeter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Molarity with the unit abbreviation.
func (a *Molarity) ToString(unit MolarityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMolarityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMolarityAbbreviation(unit))
}

// Equals checks if the given Molarity is equal to the current Molarity.
//
// Parameters:
//    other: The Molarity to compare against.
//
// Returns:
//    bool: Returns true if both Molarity are equal, false otherwise.
func (a *Molarity) Equals(other *Molarity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Molarity with another Molarity.
// It returns -1 if the current Molarity is less than the other Molarity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Molarity to compare against.
//
// Returns:
//    int: -1 if the current Molarity is less, 1 if greater, and 0 if equal.
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

// Add adds the given Molarity to the current Molarity and returns the result.
// The result is a new Molarity instance with the sum of the values.
//
// Parameters:
//    other: The Molarity to add to the current Molarity.
//
// Returns:
//    *Molarity: A new Molarity instance representing the sum of both Molarity.
func (a *Molarity) Add(other *Molarity) *Molarity {
	return &Molarity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Molarity from the current Molarity and returns the result.
// The result is a new Molarity instance with the difference of the values.
//
// Parameters:
//    other: The Molarity to subtract from the current Molarity.
//
// Returns:
//    *Molarity: A new Molarity instance representing the difference of both Molarity.
func (a *Molarity) Subtract(other *Molarity) *Molarity {
	return &Molarity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Molarity by the given Molarity and returns the result.
// The result is a new Molarity instance with the product of the values.
//
// Parameters:
//    other: The Molarity to multiply with the current Molarity.
//
// Returns:
//    *Molarity: A new Molarity instance representing the product of both Molarity.
func (a *Molarity) Multiply(other *Molarity) *Molarity {
	return &Molarity{value: a.value * other.BaseValue()}
}

// Divide divides the current Molarity by the given Molarity and returns the result.
// The result is a new Molarity instance with the quotient of the values.
//
// Parameters:
//    other: The Molarity to divide the current Molarity by.
//
// Returns:
//    *Molarity: A new Molarity instance representing the quotient of both Molarity.
func (a *Molarity) Divide(other *Molarity) *Molarity {
	return &Molarity{value: a.value / other.BaseValue()}
}

// GetMolarityAbbreviation gets the unit abbreviation.
func GetMolarityAbbreviation(unit MolarityUnits) string {
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