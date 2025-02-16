// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AmountOfSubstanceUnits defines various units of AmountOfSubstance.
type AmountOfSubstanceUnits string

const (
    
        // 
        AmountOfSubstanceMole AmountOfSubstanceUnits = "Mole"
        // 
        AmountOfSubstancePoundMole AmountOfSubstanceUnits = "PoundMole"
        // 
        AmountOfSubstanceFemtomole AmountOfSubstanceUnits = "Femtomole"
        // 
        AmountOfSubstancePicomole AmountOfSubstanceUnits = "Picomole"
        // 
        AmountOfSubstanceNanomole AmountOfSubstanceUnits = "Nanomole"
        // 
        AmountOfSubstanceMicromole AmountOfSubstanceUnits = "Micromole"
        // 
        AmountOfSubstanceMillimole AmountOfSubstanceUnits = "Millimole"
        // 
        AmountOfSubstanceCentimole AmountOfSubstanceUnits = "Centimole"
        // 
        AmountOfSubstanceDecimole AmountOfSubstanceUnits = "Decimole"
        // 
        AmountOfSubstanceKilomole AmountOfSubstanceUnits = "Kilomole"
        // 
        AmountOfSubstanceMegamole AmountOfSubstanceUnits = "Megamole"
        // 
        AmountOfSubstanceNanopoundMole AmountOfSubstanceUnits = "NanopoundMole"
        // 
        AmountOfSubstanceMicropoundMole AmountOfSubstanceUnits = "MicropoundMole"
        // 
        AmountOfSubstanceMillipoundMole AmountOfSubstanceUnits = "MillipoundMole"
        // 
        AmountOfSubstanceCentipoundMole AmountOfSubstanceUnits = "CentipoundMole"
        // 
        AmountOfSubstanceDecipoundMole AmountOfSubstanceUnits = "DecipoundMole"
        // 
        AmountOfSubstanceKilopoundMole AmountOfSubstanceUnits = "KilopoundMole"
)

// AmountOfSubstanceDto represents a AmountOfSubstance measurement with a numerical value and its corresponding unit.
type AmountOfSubstanceDto struct {
    // Value is the numerical representation of the AmountOfSubstance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the AmountOfSubstance, as defined in the AmountOfSubstanceUnits enumeration.
	Unit  AmountOfSubstanceUnits `json:"unit"`
}

// AmountOfSubstanceDtoFactory groups methods for creating and serializing AmountOfSubstanceDto objects.
type AmountOfSubstanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AmountOfSubstanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AmountOfSubstanceDtoFactory) FromJSON(data []byte) (*AmountOfSubstanceDto, error) {
	a := AmountOfSubstanceDto{}

    // Parse JSON into AmountOfSubstanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AmountOfSubstanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AmountOfSubstanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// AmountOfSubstance represents a measurement in a of AmountOfSubstance.
//
// Mole is the amount of substance containing Avagadro's Number (6.02 x 10 ^ 23) of real particles such as molecules,atoms, ions or radicals.
type AmountOfSubstance struct {
	// value is the base measurement stored internally.
	value       float64
    
    molesLazy *float64 
    pound_molesLazy *float64 
    femtomolesLazy *float64 
    picomolesLazy *float64 
    nanomolesLazy *float64 
    micromolesLazy *float64 
    millimolesLazy *float64 
    centimolesLazy *float64 
    decimolesLazy *float64 
    kilomolesLazy *float64 
    megamolesLazy *float64 
    nanopound_molesLazy *float64 
    micropound_molesLazy *float64 
    millipound_molesLazy *float64 
    centipound_molesLazy *float64 
    decipound_molesLazy *float64 
    kilopound_molesLazy *float64 
}

// AmountOfSubstanceFactory groups methods for creating AmountOfSubstance instances.
type AmountOfSubstanceFactory struct{}

// CreateAmountOfSubstance creates a new AmountOfSubstance instance from the given value and unit.
func (uf AmountOfSubstanceFactory) CreateAmountOfSubstance(value float64, unit AmountOfSubstanceUnits) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, unit)
}

// FromDto converts a AmountOfSubstanceDto to a AmountOfSubstance instance.
func (uf AmountOfSubstanceFactory) FromDto(dto AmountOfSubstanceDto) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AmountOfSubstance instance.
func (uf AmountOfSubstanceFactory) FromDtoJSON(data []byte) (*AmountOfSubstance, error) {
	unitDto, err := AmountOfSubstanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AmountOfSubstanceDto from JSON: %w", err)
	}
	return AmountOfSubstanceFactory{}.FromDto(*unitDto)
}


// FromMoles creates a new AmountOfSubstance instance from a value in Moles.
func (uf AmountOfSubstanceFactory) FromMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMole)
}

// FromPoundMoles creates a new AmountOfSubstance instance from a value in PoundMoles.
func (uf AmountOfSubstanceFactory) FromPoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstancePoundMole)
}

// FromFemtomoles creates a new AmountOfSubstance instance from a value in Femtomoles.
func (uf AmountOfSubstanceFactory) FromFemtomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceFemtomole)
}

// FromPicomoles creates a new AmountOfSubstance instance from a value in Picomoles.
func (uf AmountOfSubstanceFactory) FromPicomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstancePicomole)
}

// FromNanomoles creates a new AmountOfSubstance instance from a value in Nanomoles.
func (uf AmountOfSubstanceFactory) FromNanomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceNanomole)
}

// FromMicromoles creates a new AmountOfSubstance instance from a value in Micromoles.
func (uf AmountOfSubstanceFactory) FromMicromoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMicromole)
}

// FromMillimoles creates a new AmountOfSubstance instance from a value in Millimoles.
func (uf AmountOfSubstanceFactory) FromMillimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMillimole)
}

// FromCentimoles creates a new AmountOfSubstance instance from a value in Centimoles.
func (uf AmountOfSubstanceFactory) FromCentimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceCentimole)
}

// FromDecimoles creates a new AmountOfSubstance instance from a value in Decimoles.
func (uf AmountOfSubstanceFactory) FromDecimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceDecimole)
}

// FromKilomoles creates a new AmountOfSubstance instance from a value in Kilomoles.
func (uf AmountOfSubstanceFactory) FromKilomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceKilomole)
}

// FromMegamoles creates a new AmountOfSubstance instance from a value in Megamoles.
func (uf AmountOfSubstanceFactory) FromMegamoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMegamole)
}

// FromNanopoundMoles creates a new AmountOfSubstance instance from a value in NanopoundMoles.
func (uf AmountOfSubstanceFactory) FromNanopoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceNanopoundMole)
}

// FromMicropoundMoles creates a new AmountOfSubstance instance from a value in MicropoundMoles.
func (uf AmountOfSubstanceFactory) FromMicropoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMicropoundMole)
}

// FromMillipoundMoles creates a new AmountOfSubstance instance from a value in MillipoundMoles.
func (uf AmountOfSubstanceFactory) FromMillipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMillipoundMole)
}

// FromCentipoundMoles creates a new AmountOfSubstance instance from a value in CentipoundMoles.
func (uf AmountOfSubstanceFactory) FromCentipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceCentipoundMole)
}

// FromDecipoundMoles creates a new AmountOfSubstance instance from a value in DecipoundMoles.
func (uf AmountOfSubstanceFactory) FromDecipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceDecipoundMole)
}

// FromKilopoundMoles creates a new AmountOfSubstance instance from a value in KilopoundMoles.
func (uf AmountOfSubstanceFactory) FromKilopoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceKilopoundMole)
}


// newAmountOfSubstance creates a new AmountOfSubstance.
func newAmountOfSubstance(value float64, fromUnit AmountOfSubstanceUnits) (*AmountOfSubstance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AmountOfSubstance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AmountOfSubstance in Mole unit (the base/default quantity).
func (a *AmountOfSubstance) BaseValue() float64 {
	return a.value
}


// Moles returns the AmountOfSubstance value in Moles.
//
// 
func (a *AmountOfSubstance) Moles() float64 {
	if a.molesLazy != nil {
		return *a.molesLazy
	}
	moles := a.convertFromBase(AmountOfSubstanceMole)
	a.molesLazy = &moles
	return moles
}

// PoundMoles returns the AmountOfSubstance value in PoundMoles.
//
// 
func (a *AmountOfSubstance) PoundMoles() float64 {
	if a.pound_molesLazy != nil {
		return *a.pound_molesLazy
	}
	pound_moles := a.convertFromBase(AmountOfSubstancePoundMole)
	a.pound_molesLazy = &pound_moles
	return pound_moles
}

// Femtomoles returns the AmountOfSubstance value in Femtomoles.
//
// 
func (a *AmountOfSubstance) Femtomoles() float64 {
	if a.femtomolesLazy != nil {
		return *a.femtomolesLazy
	}
	femtomoles := a.convertFromBase(AmountOfSubstanceFemtomole)
	a.femtomolesLazy = &femtomoles
	return femtomoles
}

// Picomoles returns the AmountOfSubstance value in Picomoles.
//
// 
func (a *AmountOfSubstance) Picomoles() float64 {
	if a.picomolesLazy != nil {
		return *a.picomolesLazy
	}
	picomoles := a.convertFromBase(AmountOfSubstancePicomole)
	a.picomolesLazy = &picomoles
	return picomoles
}

// Nanomoles returns the AmountOfSubstance value in Nanomoles.
//
// 
func (a *AmountOfSubstance) Nanomoles() float64 {
	if a.nanomolesLazy != nil {
		return *a.nanomolesLazy
	}
	nanomoles := a.convertFromBase(AmountOfSubstanceNanomole)
	a.nanomolesLazy = &nanomoles
	return nanomoles
}

// Micromoles returns the AmountOfSubstance value in Micromoles.
//
// 
func (a *AmountOfSubstance) Micromoles() float64 {
	if a.micromolesLazy != nil {
		return *a.micromolesLazy
	}
	micromoles := a.convertFromBase(AmountOfSubstanceMicromole)
	a.micromolesLazy = &micromoles
	return micromoles
}

// Millimoles returns the AmountOfSubstance value in Millimoles.
//
// 
func (a *AmountOfSubstance) Millimoles() float64 {
	if a.millimolesLazy != nil {
		return *a.millimolesLazy
	}
	millimoles := a.convertFromBase(AmountOfSubstanceMillimole)
	a.millimolesLazy = &millimoles
	return millimoles
}

// Centimoles returns the AmountOfSubstance value in Centimoles.
//
// 
func (a *AmountOfSubstance) Centimoles() float64 {
	if a.centimolesLazy != nil {
		return *a.centimolesLazy
	}
	centimoles := a.convertFromBase(AmountOfSubstanceCentimole)
	a.centimolesLazy = &centimoles
	return centimoles
}

// Decimoles returns the AmountOfSubstance value in Decimoles.
//
// 
func (a *AmountOfSubstance) Decimoles() float64 {
	if a.decimolesLazy != nil {
		return *a.decimolesLazy
	}
	decimoles := a.convertFromBase(AmountOfSubstanceDecimole)
	a.decimolesLazy = &decimoles
	return decimoles
}

// Kilomoles returns the AmountOfSubstance value in Kilomoles.
//
// 
func (a *AmountOfSubstance) Kilomoles() float64 {
	if a.kilomolesLazy != nil {
		return *a.kilomolesLazy
	}
	kilomoles := a.convertFromBase(AmountOfSubstanceKilomole)
	a.kilomolesLazy = &kilomoles
	return kilomoles
}

// Megamoles returns the AmountOfSubstance value in Megamoles.
//
// 
func (a *AmountOfSubstance) Megamoles() float64 {
	if a.megamolesLazy != nil {
		return *a.megamolesLazy
	}
	megamoles := a.convertFromBase(AmountOfSubstanceMegamole)
	a.megamolesLazy = &megamoles
	return megamoles
}

// NanopoundMoles returns the AmountOfSubstance value in NanopoundMoles.
//
// 
func (a *AmountOfSubstance) NanopoundMoles() float64 {
	if a.nanopound_molesLazy != nil {
		return *a.nanopound_molesLazy
	}
	nanopound_moles := a.convertFromBase(AmountOfSubstanceNanopoundMole)
	a.nanopound_molesLazy = &nanopound_moles
	return nanopound_moles
}

// MicropoundMoles returns the AmountOfSubstance value in MicropoundMoles.
//
// 
func (a *AmountOfSubstance) MicropoundMoles() float64 {
	if a.micropound_molesLazy != nil {
		return *a.micropound_molesLazy
	}
	micropound_moles := a.convertFromBase(AmountOfSubstanceMicropoundMole)
	a.micropound_molesLazy = &micropound_moles
	return micropound_moles
}

// MillipoundMoles returns the AmountOfSubstance value in MillipoundMoles.
//
// 
func (a *AmountOfSubstance) MillipoundMoles() float64 {
	if a.millipound_molesLazy != nil {
		return *a.millipound_molesLazy
	}
	millipound_moles := a.convertFromBase(AmountOfSubstanceMillipoundMole)
	a.millipound_molesLazy = &millipound_moles
	return millipound_moles
}

// CentipoundMoles returns the AmountOfSubstance value in CentipoundMoles.
//
// 
func (a *AmountOfSubstance) CentipoundMoles() float64 {
	if a.centipound_molesLazy != nil {
		return *a.centipound_molesLazy
	}
	centipound_moles := a.convertFromBase(AmountOfSubstanceCentipoundMole)
	a.centipound_molesLazy = &centipound_moles
	return centipound_moles
}

// DecipoundMoles returns the AmountOfSubstance value in DecipoundMoles.
//
// 
func (a *AmountOfSubstance) DecipoundMoles() float64 {
	if a.decipound_molesLazy != nil {
		return *a.decipound_molesLazy
	}
	decipound_moles := a.convertFromBase(AmountOfSubstanceDecipoundMole)
	a.decipound_molesLazy = &decipound_moles
	return decipound_moles
}

// KilopoundMoles returns the AmountOfSubstance value in KilopoundMoles.
//
// 
func (a *AmountOfSubstance) KilopoundMoles() float64 {
	if a.kilopound_molesLazy != nil {
		return *a.kilopound_molesLazy
	}
	kilopound_moles := a.convertFromBase(AmountOfSubstanceKilopoundMole)
	a.kilopound_molesLazy = &kilopound_moles
	return kilopound_moles
}


// ToDto creates a AmountOfSubstanceDto representation from the AmountOfSubstance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Mole by default.
func (a *AmountOfSubstance) ToDto(holdInUnit *AmountOfSubstanceUnits) AmountOfSubstanceDto {
	if holdInUnit == nil {
		defaultUnit := AmountOfSubstanceMole // Default value
		holdInUnit = &defaultUnit
	}

	return AmountOfSubstanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AmountOfSubstance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Mole by default.
func (a *AmountOfSubstance) ToDtoJSON(holdInUnit *AmountOfSubstanceUnits) ([]byte, error) {
	// Convert to AmountOfSubstanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AmountOfSubstance to a specific unit value.
// The function uses the provided unit type (AmountOfSubstanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AmountOfSubstance) Convert(toUnit AmountOfSubstanceUnits) float64 {
	switch toUnit { 
    case AmountOfSubstanceMole:
		return a.Moles()
    case AmountOfSubstancePoundMole:
		return a.PoundMoles()
    case AmountOfSubstanceFemtomole:
		return a.Femtomoles()
    case AmountOfSubstancePicomole:
		return a.Picomoles()
    case AmountOfSubstanceNanomole:
		return a.Nanomoles()
    case AmountOfSubstanceMicromole:
		return a.Micromoles()
    case AmountOfSubstanceMillimole:
		return a.Millimoles()
    case AmountOfSubstanceCentimole:
		return a.Centimoles()
    case AmountOfSubstanceDecimole:
		return a.Decimoles()
    case AmountOfSubstanceKilomole:
		return a.Kilomoles()
    case AmountOfSubstanceMegamole:
		return a.Megamoles()
    case AmountOfSubstanceNanopoundMole:
		return a.NanopoundMoles()
    case AmountOfSubstanceMicropoundMole:
		return a.MicropoundMoles()
    case AmountOfSubstanceMillipoundMole:
		return a.MillipoundMoles()
    case AmountOfSubstanceCentipoundMole:
		return a.CentipoundMoles()
    case AmountOfSubstanceDecipoundMole:
		return a.DecipoundMoles()
    case AmountOfSubstanceKilopoundMole:
		return a.KilopoundMoles()
	default:
		return math.NaN()
	}
}

func (a *AmountOfSubstance) convertFromBase(toUnit AmountOfSubstanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case AmountOfSubstanceMole:
		return (value) 
	case AmountOfSubstancePoundMole:
		return (value / 453.59237) 
	case AmountOfSubstanceFemtomole:
		return ((value) / 1e-15) 
	case AmountOfSubstancePicomole:
		return ((value) / 1e-12) 
	case AmountOfSubstanceNanomole:
		return ((value) / 1e-09) 
	case AmountOfSubstanceMicromole:
		return ((value) / 1e-06) 
	case AmountOfSubstanceMillimole:
		return ((value) / 0.001) 
	case AmountOfSubstanceCentimole:
		return ((value) / 0.01) 
	case AmountOfSubstanceDecimole:
		return ((value) / 0.1) 
	case AmountOfSubstanceKilomole:
		return ((value) / 1000.0) 
	case AmountOfSubstanceMegamole:
		return ((value) / 1000000.0) 
	case AmountOfSubstanceNanopoundMole:
		return ((value / 453.59237) / 1e-09) 
	case AmountOfSubstanceMicropoundMole:
		return ((value / 453.59237) / 1e-06) 
	case AmountOfSubstanceMillipoundMole:
		return ((value / 453.59237) / 0.001) 
	case AmountOfSubstanceCentipoundMole:
		return ((value / 453.59237) / 0.01) 
	case AmountOfSubstanceDecipoundMole:
		return ((value / 453.59237) / 0.1) 
	case AmountOfSubstanceKilopoundMole:
		return ((value / 453.59237) / 1000.0) 
	default:
		return math.NaN()
	}
}

func (a *AmountOfSubstance) convertToBase(value float64, fromUnit AmountOfSubstanceUnits) float64 {
	switch fromUnit { 
	case AmountOfSubstanceMole:
		return (value) 
	case AmountOfSubstancePoundMole:
		return (value * 453.59237) 
	case AmountOfSubstanceFemtomole:
		return ((value) * 1e-15) 
	case AmountOfSubstancePicomole:
		return ((value) * 1e-12) 
	case AmountOfSubstanceNanomole:
		return ((value) * 1e-09) 
	case AmountOfSubstanceMicromole:
		return ((value) * 1e-06) 
	case AmountOfSubstanceMillimole:
		return ((value) * 0.001) 
	case AmountOfSubstanceCentimole:
		return ((value) * 0.01) 
	case AmountOfSubstanceDecimole:
		return ((value) * 0.1) 
	case AmountOfSubstanceKilomole:
		return ((value) * 1000.0) 
	case AmountOfSubstanceMegamole:
		return ((value) * 1000000.0) 
	case AmountOfSubstanceNanopoundMole:
		return ((value * 453.59237) * 1e-09) 
	case AmountOfSubstanceMicropoundMole:
		return ((value * 453.59237) * 1e-06) 
	case AmountOfSubstanceMillipoundMole:
		return ((value * 453.59237) * 0.001) 
	case AmountOfSubstanceCentipoundMole:
		return ((value * 453.59237) * 0.01) 
	case AmountOfSubstanceDecipoundMole:
		return ((value * 453.59237) * 0.1) 
	case AmountOfSubstanceKilopoundMole:
		return ((value * 453.59237) * 1000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AmountOfSubstance in the default unit (Mole),
// formatted to two decimal places.
func (a AmountOfSubstance) String() string {
	return a.ToString(AmountOfSubstanceMole, 2)
}

// ToString formats the AmountOfSubstance value as a string with the specified unit and fractional digits.
// It converts the AmountOfSubstance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AmountOfSubstance value will be converted (e.g., Mole))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AmountOfSubstance with the unit abbreviation.
func (a *AmountOfSubstance) ToString(unit AmountOfSubstanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAmountOfSubstanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAmountOfSubstanceAbbreviation(unit))
}

// Equals checks if the given AmountOfSubstance is equal to the current AmountOfSubstance.
//
// Parameters:
//    other: The AmountOfSubstance to compare against.
//
// Returns:
//    bool: Returns true if both AmountOfSubstance are equal, false otherwise.
func (a *AmountOfSubstance) Equals(other *AmountOfSubstance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AmountOfSubstance with another AmountOfSubstance.
// It returns -1 if the current AmountOfSubstance is less than the other AmountOfSubstance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AmountOfSubstance to compare against.
//
// Returns:
//    int: -1 if the current AmountOfSubstance is less, 1 if greater, and 0 if equal.
func (a *AmountOfSubstance) CompareTo(other *AmountOfSubstance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AmountOfSubstance to the current AmountOfSubstance and returns the result.
// The result is a new AmountOfSubstance instance with the sum of the values.
//
// Parameters:
//    other: The AmountOfSubstance to add to the current AmountOfSubstance.
//
// Returns:
//    *AmountOfSubstance: A new AmountOfSubstance instance representing the sum of both AmountOfSubstance.
func (a *AmountOfSubstance) Add(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AmountOfSubstance from the current AmountOfSubstance and returns the result.
// The result is a new AmountOfSubstance instance with the difference of the values.
//
// Parameters:
//    other: The AmountOfSubstance to subtract from the current AmountOfSubstance.
//
// Returns:
//    *AmountOfSubstance: A new AmountOfSubstance instance representing the difference of both AmountOfSubstance.
func (a *AmountOfSubstance) Subtract(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AmountOfSubstance by the given AmountOfSubstance and returns the result.
// The result is a new AmountOfSubstance instance with the product of the values.
//
// Parameters:
//    other: The AmountOfSubstance to multiply with the current AmountOfSubstance.
//
// Returns:
//    *AmountOfSubstance: A new AmountOfSubstance instance representing the product of both AmountOfSubstance.
func (a *AmountOfSubstance) Multiply(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value * other.BaseValue()}
}

// Divide divides the current AmountOfSubstance by the given AmountOfSubstance and returns the result.
// The result is a new AmountOfSubstance instance with the quotient of the values.
//
// Parameters:
//    other: The AmountOfSubstance to divide the current AmountOfSubstance by.
//
// Returns:
//    *AmountOfSubstance: A new AmountOfSubstance instance representing the quotient of both AmountOfSubstance.
func (a *AmountOfSubstance) Divide(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value / other.BaseValue()}
}

// GetAmountOfSubstanceAbbreviation gets the unit abbreviation.
func GetAmountOfSubstanceAbbreviation(unit AmountOfSubstanceUnits) string {
	switch unit { 
	case AmountOfSubstanceMole:
		return "mol" 
	case AmountOfSubstancePoundMole:
		return "lbmol" 
	case AmountOfSubstanceFemtomole:
		return "fmol" 
	case AmountOfSubstancePicomole:
		return "pmol" 
	case AmountOfSubstanceNanomole:
		return "nmol" 
	case AmountOfSubstanceMicromole:
		return "μmol" 
	case AmountOfSubstanceMillimole:
		return "mmol" 
	case AmountOfSubstanceCentimole:
		return "cmol" 
	case AmountOfSubstanceDecimole:
		return "dmol" 
	case AmountOfSubstanceKilomole:
		return "kmol" 
	case AmountOfSubstanceMegamole:
		return "Mmol" 
	case AmountOfSubstanceNanopoundMole:
		return "nlbmol" 
	case AmountOfSubstanceMicropoundMole:
		return "μlbmol" 
	case AmountOfSubstanceMillipoundMole:
		return "mlbmol" 
	case AmountOfSubstanceCentipoundMole:
		return "clbmol" 
	case AmountOfSubstanceDecipoundMole:
		return "dlbmol" 
	case AmountOfSubstanceKilopoundMole:
		return "klbmol" 
	default:
		return ""
	}
}