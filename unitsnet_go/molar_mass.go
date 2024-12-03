package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MolarMassUnits enumeration
type MolarMassUnits string

const (
    
        // 
        MolarMassGramPerMole MolarMassUnits = "GramPerMole"
        // 
        MolarMassKilogramPerKilomole MolarMassUnits = "KilogramPerKilomole"
        // 
        MolarMassPoundPerMole MolarMassUnits = "PoundPerMole"
        // 
        MolarMassNanogramPerMole MolarMassUnits = "NanogramPerMole"
        // 
        MolarMassMicrogramPerMole MolarMassUnits = "MicrogramPerMole"
        // 
        MolarMassMilligramPerMole MolarMassUnits = "MilligramPerMole"
        // 
        MolarMassCentigramPerMole MolarMassUnits = "CentigramPerMole"
        // 
        MolarMassDecigramPerMole MolarMassUnits = "DecigramPerMole"
        // 
        MolarMassDecagramPerMole MolarMassUnits = "DecagramPerMole"
        // 
        MolarMassHectogramPerMole MolarMassUnits = "HectogramPerMole"
        // 
        MolarMassKilogramPerMole MolarMassUnits = "KilogramPerMole"
        // 
        MolarMassKilopoundPerMole MolarMassUnits = "KilopoundPerMole"
        // 
        MolarMassMegapoundPerMole MolarMassUnits = "MegapoundPerMole"
)

// MolarMassDto represents an MolarMass
type MolarMassDto struct {
	Value float64
	Unit  MolarMassUnits
}

// MolarMassDtoFactory struct to group related functions
type MolarMassDtoFactory struct{}

func (udf MolarMassDtoFactory) FromJSON(data []byte) (*MolarMassDto, error) {
	a := MolarMassDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MolarMassDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// MolarMass struct
type MolarMass struct {
	value       float64
    
    grams_per_moleLazy *float64 
    kilograms_per_kilomoleLazy *float64 
    pounds_per_moleLazy *float64 
    nanograms_per_moleLazy *float64 
    micrograms_per_moleLazy *float64 
    milligrams_per_moleLazy *float64 
    centigrams_per_moleLazy *float64 
    decigrams_per_moleLazy *float64 
    decagrams_per_moleLazy *float64 
    hectograms_per_moleLazy *float64 
    kilograms_per_moleLazy *float64 
    kilopounds_per_moleLazy *float64 
    megapounds_per_moleLazy *float64 
}

// MolarMassFactory struct to group related functions
type MolarMassFactory struct{}

func (uf MolarMassFactory) CreateMolarMass(value float64, unit MolarMassUnits) (*MolarMass, error) {
	return newMolarMass(value, unit)
}

func (uf MolarMassFactory) FromDto(dto MolarMassDto) (*MolarMass, error) {
	return newMolarMass(dto.Value, dto.Unit)
}

func (uf MolarMassFactory) FromDtoJSON(data []byte) (*MolarMass, error) {
	unitDto, err := MolarMassDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MolarMassFactory{}.FromDto(*unitDto)
}


// FromGramPerMole creates a new MolarMass instance from GramPerMole.
func (uf MolarMassFactory) FromGramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassGramPerMole)
}

// FromKilogramPerKilomole creates a new MolarMass instance from KilogramPerKilomole.
func (uf MolarMassFactory) FromKilogramsPerKilomole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilogramPerKilomole)
}

// FromPoundPerMole creates a new MolarMass instance from PoundPerMole.
func (uf MolarMassFactory) FromPoundsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassPoundPerMole)
}

// FromNanogramPerMole creates a new MolarMass instance from NanogramPerMole.
func (uf MolarMassFactory) FromNanogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassNanogramPerMole)
}

// FromMicrogramPerMole creates a new MolarMass instance from MicrogramPerMole.
func (uf MolarMassFactory) FromMicrogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassMicrogramPerMole)
}

// FromMilligramPerMole creates a new MolarMass instance from MilligramPerMole.
func (uf MolarMassFactory) FromMilligramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassMilligramPerMole)
}

// FromCentigramPerMole creates a new MolarMass instance from CentigramPerMole.
func (uf MolarMassFactory) FromCentigramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassCentigramPerMole)
}

// FromDecigramPerMole creates a new MolarMass instance from DecigramPerMole.
func (uf MolarMassFactory) FromDecigramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassDecigramPerMole)
}

// FromDecagramPerMole creates a new MolarMass instance from DecagramPerMole.
func (uf MolarMassFactory) FromDecagramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassDecagramPerMole)
}

// FromHectogramPerMole creates a new MolarMass instance from HectogramPerMole.
func (uf MolarMassFactory) FromHectogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassHectogramPerMole)
}

// FromKilogramPerMole creates a new MolarMass instance from KilogramPerMole.
func (uf MolarMassFactory) FromKilogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilogramPerMole)
}

// FromKilopoundPerMole creates a new MolarMass instance from KilopoundPerMole.
func (uf MolarMassFactory) FromKilopoundsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilopoundPerMole)
}

// FromMegapoundPerMole creates a new MolarMass instance from MegapoundPerMole.
func (uf MolarMassFactory) FromMegapoundsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassMegapoundPerMole)
}




// newMolarMass creates a new MolarMass.
func newMolarMass(value float64, fromUnit MolarMassUnits) (*MolarMass, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &MolarMass{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of MolarMass in KilogramPerMole.
func (a *MolarMass) BaseValue() float64 {
	return a.value
}


// GramPerMole returns the value in GramPerMole.
func (a *MolarMass) GramsPerMole() float64 {
	if a.grams_per_moleLazy != nil {
		return *a.grams_per_moleLazy
	}
	grams_per_mole := a.convertFromBase(MolarMassGramPerMole)
	a.grams_per_moleLazy = &grams_per_mole
	return grams_per_mole
}

// KilogramPerKilomole returns the value in KilogramPerKilomole.
func (a *MolarMass) KilogramsPerKilomole() float64 {
	if a.kilograms_per_kilomoleLazy != nil {
		return *a.kilograms_per_kilomoleLazy
	}
	kilograms_per_kilomole := a.convertFromBase(MolarMassKilogramPerKilomole)
	a.kilograms_per_kilomoleLazy = &kilograms_per_kilomole
	return kilograms_per_kilomole
}

// PoundPerMole returns the value in PoundPerMole.
func (a *MolarMass) PoundsPerMole() float64 {
	if a.pounds_per_moleLazy != nil {
		return *a.pounds_per_moleLazy
	}
	pounds_per_mole := a.convertFromBase(MolarMassPoundPerMole)
	a.pounds_per_moleLazy = &pounds_per_mole
	return pounds_per_mole
}

// NanogramPerMole returns the value in NanogramPerMole.
func (a *MolarMass) NanogramsPerMole() float64 {
	if a.nanograms_per_moleLazy != nil {
		return *a.nanograms_per_moleLazy
	}
	nanograms_per_mole := a.convertFromBase(MolarMassNanogramPerMole)
	a.nanograms_per_moleLazy = &nanograms_per_mole
	return nanograms_per_mole
}

// MicrogramPerMole returns the value in MicrogramPerMole.
func (a *MolarMass) MicrogramsPerMole() float64 {
	if a.micrograms_per_moleLazy != nil {
		return *a.micrograms_per_moleLazy
	}
	micrograms_per_mole := a.convertFromBase(MolarMassMicrogramPerMole)
	a.micrograms_per_moleLazy = &micrograms_per_mole
	return micrograms_per_mole
}

// MilligramPerMole returns the value in MilligramPerMole.
func (a *MolarMass) MilligramsPerMole() float64 {
	if a.milligrams_per_moleLazy != nil {
		return *a.milligrams_per_moleLazy
	}
	milligrams_per_mole := a.convertFromBase(MolarMassMilligramPerMole)
	a.milligrams_per_moleLazy = &milligrams_per_mole
	return milligrams_per_mole
}

// CentigramPerMole returns the value in CentigramPerMole.
func (a *MolarMass) CentigramsPerMole() float64 {
	if a.centigrams_per_moleLazy != nil {
		return *a.centigrams_per_moleLazy
	}
	centigrams_per_mole := a.convertFromBase(MolarMassCentigramPerMole)
	a.centigrams_per_moleLazy = &centigrams_per_mole
	return centigrams_per_mole
}

// DecigramPerMole returns the value in DecigramPerMole.
func (a *MolarMass) DecigramsPerMole() float64 {
	if a.decigrams_per_moleLazy != nil {
		return *a.decigrams_per_moleLazy
	}
	decigrams_per_mole := a.convertFromBase(MolarMassDecigramPerMole)
	a.decigrams_per_moleLazy = &decigrams_per_mole
	return decigrams_per_mole
}

// DecagramPerMole returns the value in DecagramPerMole.
func (a *MolarMass) DecagramsPerMole() float64 {
	if a.decagrams_per_moleLazy != nil {
		return *a.decagrams_per_moleLazy
	}
	decagrams_per_mole := a.convertFromBase(MolarMassDecagramPerMole)
	a.decagrams_per_moleLazy = &decagrams_per_mole
	return decagrams_per_mole
}

// HectogramPerMole returns the value in HectogramPerMole.
func (a *MolarMass) HectogramsPerMole() float64 {
	if a.hectograms_per_moleLazy != nil {
		return *a.hectograms_per_moleLazy
	}
	hectograms_per_mole := a.convertFromBase(MolarMassHectogramPerMole)
	a.hectograms_per_moleLazy = &hectograms_per_mole
	return hectograms_per_mole
}

// KilogramPerMole returns the value in KilogramPerMole.
func (a *MolarMass) KilogramsPerMole() float64 {
	if a.kilograms_per_moleLazy != nil {
		return *a.kilograms_per_moleLazy
	}
	kilograms_per_mole := a.convertFromBase(MolarMassKilogramPerMole)
	a.kilograms_per_moleLazy = &kilograms_per_mole
	return kilograms_per_mole
}

// KilopoundPerMole returns the value in KilopoundPerMole.
func (a *MolarMass) KilopoundsPerMole() float64 {
	if a.kilopounds_per_moleLazy != nil {
		return *a.kilopounds_per_moleLazy
	}
	kilopounds_per_mole := a.convertFromBase(MolarMassKilopoundPerMole)
	a.kilopounds_per_moleLazy = &kilopounds_per_mole
	return kilopounds_per_mole
}

// MegapoundPerMole returns the value in MegapoundPerMole.
func (a *MolarMass) MegapoundsPerMole() float64 {
	if a.megapounds_per_moleLazy != nil {
		return *a.megapounds_per_moleLazy
	}
	megapounds_per_mole := a.convertFromBase(MolarMassMegapoundPerMole)
	a.megapounds_per_moleLazy = &megapounds_per_mole
	return megapounds_per_mole
}


// ToDto creates an MolarMassDto representation.
func (a *MolarMass) ToDto(holdInUnit *MolarMassUnits) MolarMassDto {
	if holdInUnit == nil {
		defaultUnit := MolarMassKilogramPerMole // Default value
		holdInUnit = &defaultUnit
	}

	return MolarMassDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an MolarMassDto representation.
func (a *MolarMass) ToDtoJSON(holdInUnit *MolarMassUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts MolarMass to a specific unit value.
func (a *MolarMass) Convert(toUnit MolarMassUnits) float64 {
	switch toUnit { 
    case MolarMassGramPerMole:
		return a.GramsPerMole()
    case MolarMassKilogramPerKilomole:
		return a.KilogramsPerKilomole()
    case MolarMassPoundPerMole:
		return a.PoundsPerMole()
    case MolarMassNanogramPerMole:
		return a.NanogramsPerMole()
    case MolarMassMicrogramPerMole:
		return a.MicrogramsPerMole()
    case MolarMassMilligramPerMole:
		return a.MilligramsPerMole()
    case MolarMassCentigramPerMole:
		return a.CentigramsPerMole()
    case MolarMassDecigramPerMole:
		return a.DecigramsPerMole()
    case MolarMassDecagramPerMole:
		return a.DecagramsPerMole()
    case MolarMassHectogramPerMole:
		return a.HectogramsPerMole()
    case MolarMassKilogramPerMole:
		return a.KilogramsPerMole()
    case MolarMassKilopoundPerMole:
		return a.KilopoundsPerMole()
    case MolarMassMegapoundPerMole:
		return a.MegapoundsPerMole()
	default:
		return 0
	}
}

func (a *MolarMass) convertFromBase(toUnit MolarMassUnits) float64 {
    value := a.value
	switch toUnit { 
	case MolarMassGramPerMole:
		return (value * 1e3) 
	case MolarMassKilogramPerKilomole:
		return (value * 1e3) 
	case MolarMassPoundPerMole:
		return (value / 0.45359237) 
	case MolarMassNanogramPerMole:
		return ((value * 1e3) / 1e-09) 
	case MolarMassMicrogramPerMole:
		return ((value * 1e3) / 1e-06) 
	case MolarMassMilligramPerMole:
		return ((value * 1e3) / 0.001) 
	case MolarMassCentigramPerMole:
		return ((value * 1e3) / 0.01) 
	case MolarMassDecigramPerMole:
		return ((value * 1e3) / 0.1) 
	case MolarMassDecagramPerMole:
		return ((value * 1e3) / 10.0) 
	case MolarMassHectogramPerMole:
		return ((value * 1e3) / 100.0) 
	case MolarMassKilogramPerMole:
		return ((value * 1e3) / 1000.0) 
	case MolarMassKilopoundPerMole:
		return ((value / 0.45359237) / 1000.0) 
	case MolarMassMegapoundPerMole:
		return ((value / 0.45359237) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *MolarMass) convertToBase(value float64, fromUnit MolarMassUnits) float64 {
	switch fromUnit { 
	case MolarMassGramPerMole:
		return (value / 1e3) 
	case MolarMassKilogramPerKilomole:
		return (value / 1e3) 
	case MolarMassPoundPerMole:
		return (value * 0.45359237) 
	case MolarMassNanogramPerMole:
		return ((value / 1e3) * 1e-09) 
	case MolarMassMicrogramPerMole:
		return ((value / 1e3) * 1e-06) 
	case MolarMassMilligramPerMole:
		return ((value / 1e3) * 0.001) 
	case MolarMassCentigramPerMole:
		return ((value / 1e3) * 0.01) 
	case MolarMassDecigramPerMole:
		return ((value / 1e3) * 0.1) 
	case MolarMassDecagramPerMole:
		return ((value / 1e3) * 10.0) 
	case MolarMassHectogramPerMole:
		return ((value / 1e3) * 100.0) 
	case MolarMassKilogramPerMole:
		return ((value / 1e3) * 1000.0) 
	case MolarMassKilopoundPerMole:
		return ((value * 0.45359237) * 1000.0) 
	case MolarMassMegapoundPerMole:
		return ((value * 0.45359237) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a MolarMass) String() string {
	return a.ToString(MolarMassKilogramPerMole, 2)
}

// ToString formats the MolarMass to string.
// fractionalDigits -1 for not mention
func (a *MolarMass) ToString(unit MolarMassUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *MolarMass) getUnitAbbreviation(unit MolarMassUnits) string {
	switch unit { 
	case MolarMassGramPerMole:
		return "g/mol" 
	case MolarMassKilogramPerKilomole:
		return "kg/kmol" 
	case MolarMassPoundPerMole:
		return "lb/mol" 
	case MolarMassNanogramPerMole:
		return "ng/mol" 
	case MolarMassMicrogramPerMole:
		return "μg/mol" 
	case MolarMassMilligramPerMole:
		return "mg/mol" 
	case MolarMassCentigramPerMole:
		return "cg/mol" 
	case MolarMassDecigramPerMole:
		return "dg/mol" 
	case MolarMassDecagramPerMole:
		return "dag/mol" 
	case MolarMassHectogramPerMole:
		return "hg/mol" 
	case MolarMassKilogramPerMole:
		return "kg/mol" 
	case MolarMassKilopoundPerMole:
		return "klb/mol" 
	case MolarMassMegapoundPerMole:
		return "Mlb/mol" 
	default:
		return ""
	}
}

// Check if the given MolarMass are equals to the current MolarMass
func (a *MolarMass) Equals(other *MolarMass) bool {
	return a.value == other.BaseValue()
}

// Check if the given MolarMass are equals to the current MolarMass
func (a *MolarMass) CompareTo(other *MolarMass) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given MolarMass to the current MolarMass.
func (a *MolarMass) Add(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value + other.BaseValue()}
}

// Subtract the given MolarMass to the current MolarMass.
func (a *MolarMass) Subtract(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value - other.BaseValue()}
}

// Multiply the given MolarMass to the current MolarMass.
func (a *MolarMass) Multiply(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value * other.BaseValue()}
}

// Divide the given MolarMass to the current MolarMass.
func (a *MolarMass) Divide(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value / other.BaseValue()}
}