package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AmountOfSubstanceUnits enumeration
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

// AmountOfSubstanceDto represents an AmountOfSubstance
type AmountOfSubstanceDto struct {
	Value float64
	Unit  AmountOfSubstanceUnits
}

// AmountOfSubstanceDtoFactory struct to group related functions
type AmountOfSubstanceDtoFactory struct{}

func (udf AmountOfSubstanceDtoFactory) FromJSON(data []byte) (*AmountOfSubstanceDto, error) {
	a := AmountOfSubstanceDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AmountOfSubstanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// AmountOfSubstance struct
type AmountOfSubstance struct {
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

// AmountOfSubstanceFactory struct to group related functions
type AmountOfSubstanceFactory struct{}

func (uf AmountOfSubstanceFactory) CreateAmountOfSubstance(value float64, unit AmountOfSubstanceUnits) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, unit)
}

func (uf AmountOfSubstanceFactory) FromDto(dto AmountOfSubstanceDto) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(dto.Value, dto.Unit)
}

func (uf AmountOfSubstanceFactory) FromDtoJSON(data []byte) (*AmountOfSubstance, error) {
	unitDto, err := AmountOfSubstanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AmountOfSubstanceFactory{}.FromDto(*unitDto)
}


// FromMole creates a new AmountOfSubstance instance from Mole.
func (uf AmountOfSubstanceFactory) FromMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMole)
}

// FromPoundMole creates a new AmountOfSubstance instance from PoundMole.
func (uf AmountOfSubstanceFactory) FromPoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstancePoundMole)
}

// FromFemtomole creates a new AmountOfSubstance instance from Femtomole.
func (uf AmountOfSubstanceFactory) FromFemtomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceFemtomole)
}

// FromPicomole creates a new AmountOfSubstance instance from Picomole.
func (uf AmountOfSubstanceFactory) FromPicomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstancePicomole)
}

// FromNanomole creates a new AmountOfSubstance instance from Nanomole.
func (uf AmountOfSubstanceFactory) FromNanomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceNanomole)
}

// FromMicromole creates a new AmountOfSubstance instance from Micromole.
func (uf AmountOfSubstanceFactory) FromMicromoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMicromole)
}

// FromMillimole creates a new AmountOfSubstance instance from Millimole.
func (uf AmountOfSubstanceFactory) FromMillimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMillimole)
}

// FromCentimole creates a new AmountOfSubstance instance from Centimole.
func (uf AmountOfSubstanceFactory) FromCentimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceCentimole)
}

// FromDecimole creates a new AmountOfSubstance instance from Decimole.
func (uf AmountOfSubstanceFactory) FromDecimoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceDecimole)
}

// FromKilomole creates a new AmountOfSubstance instance from Kilomole.
func (uf AmountOfSubstanceFactory) FromKilomoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceKilomole)
}

// FromMegamole creates a new AmountOfSubstance instance from Megamole.
func (uf AmountOfSubstanceFactory) FromMegamoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMegamole)
}

// FromNanopoundMole creates a new AmountOfSubstance instance from NanopoundMole.
func (uf AmountOfSubstanceFactory) FromNanopoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceNanopoundMole)
}

// FromMicropoundMole creates a new AmountOfSubstance instance from MicropoundMole.
func (uf AmountOfSubstanceFactory) FromMicropoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMicropoundMole)
}

// FromMillipoundMole creates a new AmountOfSubstance instance from MillipoundMole.
func (uf AmountOfSubstanceFactory) FromMillipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceMillipoundMole)
}

// FromCentipoundMole creates a new AmountOfSubstance instance from CentipoundMole.
func (uf AmountOfSubstanceFactory) FromCentipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceCentipoundMole)
}

// FromDecipoundMole creates a new AmountOfSubstance instance from DecipoundMole.
func (uf AmountOfSubstanceFactory) FromDecipoundMoles(value float64) (*AmountOfSubstance, error) {
	return newAmountOfSubstance(value, AmountOfSubstanceDecipoundMole)
}

// FromKilopoundMole creates a new AmountOfSubstance instance from KilopoundMole.
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

// BaseValue returns the base value of AmountOfSubstance in Mole.
func (a *AmountOfSubstance) BaseValue() float64 {
	return a.value
}


// Mole returns the value in Mole.
func (a *AmountOfSubstance) Moles() float64 {
	if a.molesLazy != nil {
		return *a.molesLazy
	}
	moles := a.convertFromBase(AmountOfSubstanceMole)
	a.molesLazy = &moles
	return moles
}

// PoundMole returns the value in PoundMole.
func (a *AmountOfSubstance) PoundMoles() float64 {
	if a.pound_molesLazy != nil {
		return *a.pound_molesLazy
	}
	pound_moles := a.convertFromBase(AmountOfSubstancePoundMole)
	a.pound_molesLazy = &pound_moles
	return pound_moles
}

// Femtomole returns the value in Femtomole.
func (a *AmountOfSubstance) Femtomoles() float64 {
	if a.femtomolesLazy != nil {
		return *a.femtomolesLazy
	}
	femtomoles := a.convertFromBase(AmountOfSubstanceFemtomole)
	a.femtomolesLazy = &femtomoles
	return femtomoles
}

// Picomole returns the value in Picomole.
func (a *AmountOfSubstance) Picomoles() float64 {
	if a.picomolesLazy != nil {
		return *a.picomolesLazy
	}
	picomoles := a.convertFromBase(AmountOfSubstancePicomole)
	a.picomolesLazy = &picomoles
	return picomoles
}

// Nanomole returns the value in Nanomole.
func (a *AmountOfSubstance) Nanomoles() float64 {
	if a.nanomolesLazy != nil {
		return *a.nanomolesLazy
	}
	nanomoles := a.convertFromBase(AmountOfSubstanceNanomole)
	a.nanomolesLazy = &nanomoles
	return nanomoles
}

// Micromole returns the value in Micromole.
func (a *AmountOfSubstance) Micromoles() float64 {
	if a.micromolesLazy != nil {
		return *a.micromolesLazy
	}
	micromoles := a.convertFromBase(AmountOfSubstanceMicromole)
	a.micromolesLazy = &micromoles
	return micromoles
}

// Millimole returns the value in Millimole.
func (a *AmountOfSubstance) Millimoles() float64 {
	if a.millimolesLazy != nil {
		return *a.millimolesLazy
	}
	millimoles := a.convertFromBase(AmountOfSubstanceMillimole)
	a.millimolesLazy = &millimoles
	return millimoles
}

// Centimole returns the value in Centimole.
func (a *AmountOfSubstance) Centimoles() float64 {
	if a.centimolesLazy != nil {
		return *a.centimolesLazy
	}
	centimoles := a.convertFromBase(AmountOfSubstanceCentimole)
	a.centimolesLazy = &centimoles
	return centimoles
}

// Decimole returns the value in Decimole.
func (a *AmountOfSubstance) Decimoles() float64 {
	if a.decimolesLazy != nil {
		return *a.decimolesLazy
	}
	decimoles := a.convertFromBase(AmountOfSubstanceDecimole)
	a.decimolesLazy = &decimoles
	return decimoles
}

// Kilomole returns the value in Kilomole.
func (a *AmountOfSubstance) Kilomoles() float64 {
	if a.kilomolesLazy != nil {
		return *a.kilomolesLazy
	}
	kilomoles := a.convertFromBase(AmountOfSubstanceKilomole)
	a.kilomolesLazy = &kilomoles
	return kilomoles
}

// Megamole returns the value in Megamole.
func (a *AmountOfSubstance) Megamoles() float64 {
	if a.megamolesLazy != nil {
		return *a.megamolesLazy
	}
	megamoles := a.convertFromBase(AmountOfSubstanceMegamole)
	a.megamolesLazy = &megamoles
	return megamoles
}

// NanopoundMole returns the value in NanopoundMole.
func (a *AmountOfSubstance) NanopoundMoles() float64 {
	if a.nanopound_molesLazy != nil {
		return *a.nanopound_molesLazy
	}
	nanopound_moles := a.convertFromBase(AmountOfSubstanceNanopoundMole)
	a.nanopound_molesLazy = &nanopound_moles
	return nanopound_moles
}

// MicropoundMole returns the value in MicropoundMole.
func (a *AmountOfSubstance) MicropoundMoles() float64 {
	if a.micropound_molesLazy != nil {
		return *a.micropound_molesLazy
	}
	micropound_moles := a.convertFromBase(AmountOfSubstanceMicropoundMole)
	a.micropound_molesLazy = &micropound_moles
	return micropound_moles
}

// MillipoundMole returns the value in MillipoundMole.
func (a *AmountOfSubstance) MillipoundMoles() float64 {
	if a.millipound_molesLazy != nil {
		return *a.millipound_molesLazy
	}
	millipound_moles := a.convertFromBase(AmountOfSubstanceMillipoundMole)
	a.millipound_molesLazy = &millipound_moles
	return millipound_moles
}

// CentipoundMole returns the value in CentipoundMole.
func (a *AmountOfSubstance) CentipoundMoles() float64 {
	if a.centipound_molesLazy != nil {
		return *a.centipound_molesLazy
	}
	centipound_moles := a.convertFromBase(AmountOfSubstanceCentipoundMole)
	a.centipound_molesLazy = &centipound_moles
	return centipound_moles
}

// DecipoundMole returns the value in DecipoundMole.
func (a *AmountOfSubstance) DecipoundMoles() float64 {
	if a.decipound_molesLazy != nil {
		return *a.decipound_molesLazy
	}
	decipound_moles := a.convertFromBase(AmountOfSubstanceDecipoundMole)
	a.decipound_molesLazy = &decipound_moles
	return decipound_moles
}

// KilopoundMole returns the value in KilopoundMole.
func (a *AmountOfSubstance) KilopoundMoles() float64 {
	if a.kilopound_molesLazy != nil {
		return *a.kilopound_molesLazy
	}
	kilopound_moles := a.convertFromBase(AmountOfSubstanceKilopoundMole)
	a.kilopound_molesLazy = &kilopound_moles
	return kilopound_moles
}


// ToDto creates an AmountOfSubstanceDto representation.
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

// ToDtoJSON creates an AmountOfSubstanceDto representation.
func (a *AmountOfSubstance) ToDtoJSON(holdInUnit *AmountOfSubstanceUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts AmountOfSubstance to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a AmountOfSubstance) String() string {
	return a.ToString(AmountOfSubstanceMole, 2)
}

// ToString formats the AmountOfSubstance to string.
// fractionalDigits -1 for not mention
func (a *AmountOfSubstance) ToString(unit AmountOfSubstanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AmountOfSubstance) getUnitAbbreviation(unit AmountOfSubstanceUnits) string {
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

// Check if the given AmountOfSubstance are equals to the current AmountOfSubstance
func (a *AmountOfSubstance) Equals(other *AmountOfSubstance) bool {
	return a.value == other.BaseValue()
}

// Check if the given AmountOfSubstance are equals to the current AmountOfSubstance
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

// Add the given AmountOfSubstance to the current AmountOfSubstance.
func (a *AmountOfSubstance) Add(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value + other.BaseValue()}
}

// Subtract the given AmountOfSubstance to the current AmountOfSubstance.
func (a *AmountOfSubstance) Subtract(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value - other.BaseValue()}
}

// Multiply the given AmountOfSubstance to the current AmountOfSubstance.
func (a *AmountOfSubstance) Multiply(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value * other.BaseValue()}
}

// Divide the given AmountOfSubstance to the current AmountOfSubstance.
func (a *AmountOfSubstance) Divide(other *AmountOfSubstance) *AmountOfSubstance {
	return &AmountOfSubstance{value: a.value / other.BaseValue()}
}