// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MolarMassUnits defines various units of MolarMass.
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

// MolarMassDto represents a MolarMass measurement with a numerical value and its corresponding unit.
type MolarMassDto struct {
    // Value is the numerical representation of the MolarMass.
	Value float64
    // Unit specifies the unit of measurement for the MolarMass, as defined in the MolarMassUnits enumeration.
	Unit  MolarMassUnits
}

// MolarMassDtoFactory groups methods for creating and serializing MolarMassDto objects.
type MolarMassDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MolarMassDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MolarMassDtoFactory) FromJSON(data []byte) (*MolarMassDto, error) {
	a := MolarMassDto{}

    // Parse JSON into MolarMassDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a MolarMassDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MolarMassDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// MolarMass represents a measurement in a of MolarMass.
//
// In chemistry, the molar mass M is a physical property defined as the mass of a given substance (chemical element or chemical compound) divided by the amount of substance.
type MolarMass struct {
	// value is the base measurement stored internally.
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

// MolarMassFactory groups methods for creating MolarMass instances.
type MolarMassFactory struct{}

// CreateMolarMass creates a new MolarMass instance from the given value and unit.
func (uf MolarMassFactory) CreateMolarMass(value float64, unit MolarMassUnits) (*MolarMass, error) {
	return newMolarMass(value, unit)
}

// FromDto converts a MolarMassDto to a MolarMass instance.
func (uf MolarMassFactory) FromDto(dto MolarMassDto) (*MolarMass, error) {
	return newMolarMass(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a MolarMass instance.
func (uf MolarMassFactory) FromDtoJSON(data []byte) (*MolarMass, error) {
	unitDto, err := MolarMassDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MolarMassDto from JSON: %w", err)
	}
	return MolarMassFactory{}.FromDto(*unitDto)
}


// FromGramsPerMole creates a new MolarMass instance from a value in GramsPerMole.
func (uf MolarMassFactory) FromGramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassGramPerMole)
}

// FromKilogramsPerKilomole creates a new MolarMass instance from a value in KilogramsPerKilomole.
func (uf MolarMassFactory) FromKilogramsPerKilomole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilogramPerKilomole)
}

// FromPoundsPerMole creates a new MolarMass instance from a value in PoundsPerMole.
func (uf MolarMassFactory) FromPoundsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassPoundPerMole)
}

// FromNanogramsPerMole creates a new MolarMass instance from a value in NanogramsPerMole.
func (uf MolarMassFactory) FromNanogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassNanogramPerMole)
}

// FromMicrogramsPerMole creates a new MolarMass instance from a value in MicrogramsPerMole.
func (uf MolarMassFactory) FromMicrogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassMicrogramPerMole)
}

// FromMilligramsPerMole creates a new MolarMass instance from a value in MilligramsPerMole.
func (uf MolarMassFactory) FromMilligramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassMilligramPerMole)
}

// FromCentigramsPerMole creates a new MolarMass instance from a value in CentigramsPerMole.
func (uf MolarMassFactory) FromCentigramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassCentigramPerMole)
}

// FromDecigramsPerMole creates a new MolarMass instance from a value in DecigramsPerMole.
func (uf MolarMassFactory) FromDecigramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassDecigramPerMole)
}

// FromDecagramsPerMole creates a new MolarMass instance from a value in DecagramsPerMole.
func (uf MolarMassFactory) FromDecagramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassDecagramPerMole)
}

// FromHectogramsPerMole creates a new MolarMass instance from a value in HectogramsPerMole.
func (uf MolarMassFactory) FromHectogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassHectogramPerMole)
}

// FromKilogramsPerMole creates a new MolarMass instance from a value in KilogramsPerMole.
func (uf MolarMassFactory) FromKilogramsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilogramPerMole)
}

// FromKilopoundsPerMole creates a new MolarMass instance from a value in KilopoundsPerMole.
func (uf MolarMassFactory) FromKilopoundsPerMole(value float64) (*MolarMass, error) {
	return newMolarMass(value, MolarMassKilopoundPerMole)
}

// FromMegapoundsPerMole creates a new MolarMass instance from a value in MegapoundsPerMole.
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

// BaseValue returns the base value of MolarMass in KilogramPerMole unit (the base/default quantity).
func (a *MolarMass) BaseValue() float64 {
	return a.value
}


// GramsPerMole returns the MolarMass value in GramsPerMole.
//
// 
func (a *MolarMass) GramsPerMole() float64 {
	if a.grams_per_moleLazy != nil {
		return *a.grams_per_moleLazy
	}
	grams_per_mole := a.convertFromBase(MolarMassGramPerMole)
	a.grams_per_moleLazy = &grams_per_mole
	return grams_per_mole
}

// KilogramsPerKilomole returns the MolarMass value in KilogramsPerKilomole.
//
// 
func (a *MolarMass) KilogramsPerKilomole() float64 {
	if a.kilograms_per_kilomoleLazy != nil {
		return *a.kilograms_per_kilomoleLazy
	}
	kilograms_per_kilomole := a.convertFromBase(MolarMassKilogramPerKilomole)
	a.kilograms_per_kilomoleLazy = &kilograms_per_kilomole
	return kilograms_per_kilomole
}

// PoundsPerMole returns the MolarMass value in PoundsPerMole.
//
// 
func (a *MolarMass) PoundsPerMole() float64 {
	if a.pounds_per_moleLazy != nil {
		return *a.pounds_per_moleLazy
	}
	pounds_per_mole := a.convertFromBase(MolarMassPoundPerMole)
	a.pounds_per_moleLazy = &pounds_per_mole
	return pounds_per_mole
}

// NanogramsPerMole returns the MolarMass value in NanogramsPerMole.
//
// 
func (a *MolarMass) NanogramsPerMole() float64 {
	if a.nanograms_per_moleLazy != nil {
		return *a.nanograms_per_moleLazy
	}
	nanograms_per_mole := a.convertFromBase(MolarMassNanogramPerMole)
	a.nanograms_per_moleLazy = &nanograms_per_mole
	return nanograms_per_mole
}

// MicrogramsPerMole returns the MolarMass value in MicrogramsPerMole.
//
// 
func (a *MolarMass) MicrogramsPerMole() float64 {
	if a.micrograms_per_moleLazy != nil {
		return *a.micrograms_per_moleLazy
	}
	micrograms_per_mole := a.convertFromBase(MolarMassMicrogramPerMole)
	a.micrograms_per_moleLazy = &micrograms_per_mole
	return micrograms_per_mole
}

// MilligramsPerMole returns the MolarMass value in MilligramsPerMole.
//
// 
func (a *MolarMass) MilligramsPerMole() float64 {
	if a.milligrams_per_moleLazy != nil {
		return *a.milligrams_per_moleLazy
	}
	milligrams_per_mole := a.convertFromBase(MolarMassMilligramPerMole)
	a.milligrams_per_moleLazy = &milligrams_per_mole
	return milligrams_per_mole
}

// CentigramsPerMole returns the MolarMass value in CentigramsPerMole.
//
// 
func (a *MolarMass) CentigramsPerMole() float64 {
	if a.centigrams_per_moleLazy != nil {
		return *a.centigrams_per_moleLazy
	}
	centigrams_per_mole := a.convertFromBase(MolarMassCentigramPerMole)
	a.centigrams_per_moleLazy = &centigrams_per_mole
	return centigrams_per_mole
}

// DecigramsPerMole returns the MolarMass value in DecigramsPerMole.
//
// 
func (a *MolarMass) DecigramsPerMole() float64 {
	if a.decigrams_per_moleLazy != nil {
		return *a.decigrams_per_moleLazy
	}
	decigrams_per_mole := a.convertFromBase(MolarMassDecigramPerMole)
	a.decigrams_per_moleLazy = &decigrams_per_mole
	return decigrams_per_mole
}

// DecagramsPerMole returns the MolarMass value in DecagramsPerMole.
//
// 
func (a *MolarMass) DecagramsPerMole() float64 {
	if a.decagrams_per_moleLazy != nil {
		return *a.decagrams_per_moleLazy
	}
	decagrams_per_mole := a.convertFromBase(MolarMassDecagramPerMole)
	a.decagrams_per_moleLazy = &decagrams_per_mole
	return decagrams_per_mole
}

// HectogramsPerMole returns the MolarMass value in HectogramsPerMole.
//
// 
func (a *MolarMass) HectogramsPerMole() float64 {
	if a.hectograms_per_moleLazy != nil {
		return *a.hectograms_per_moleLazy
	}
	hectograms_per_mole := a.convertFromBase(MolarMassHectogramPerMole)
	a.hectograms_per_moleLazy = &hectograms_per_mole
	return hectograms_per_mole
}

// KilogramsPerMole returns the MolarMass value in KilogramsPerMole.
//
// 
func (a *MolarMass) KilogramsPerMole() float64 {
	if a.kilograms_per_moleLazy != nil {
		return *a.kilograms_per_moleLazy
	}
	kilograms_per_mole := a.convertFromBase(MolarMassKilogramPerMole)
	a.kilograms_per_moleLazy = &kilograms_per_mole
	return kilograms_per_mole
}

// KilopoundsPerMole returns the MolarMass value in KilopoundsPerMole.
//
// 
func (a *MolarMass) KilopoundsPerMole() float64 {
	if a.kilopounds_per_moleLazy != nil {
		return *a.kilopounds_per_moleLazy
	}
	kilopounds_per_mole := a.convertFromBase(MolarMassKilopoundPerMole)
	a.kilopounds_per_moleLazy = &kilopounds_per_mole
	return kilopounds_per_mole
}

// MegapoundsPerMole returns the MolarMass value in MegapoundsPerMole.
//
// 
func (a *MolarMass) MegapoundsPerMole() float64 {
	if a.megapounds_per_moleLazy != nil {
		return *a.megapounds_per_moleLazy
	}
	megapounds_per_mole := a.convertFromBase(MolarMassMegapoundPerMole)
	a.megapounds_per_moleLazy = &megapounds_per_mole
	return megapounds_per_mole
}


// ToDto creates a MolarMassDto representation from the MolarMass instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerMole by default.
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

// ToDtoJSON creates a JSON representation of the MolarMass instance.
//
// If the provided holdInUnit is nil, the value will be repesented by KilogramPerMole by default.
func (a *MolarMass) ToDtoJSON(holdInUnit *MolarMassUnits) ([]byte, error) {
	// Convert to MolarMassDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a MolarMass to a specific unit value.
// The function uses the provided unit type (MolarMassUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the MolarMass in the default unit (KilogramPerMole),
// formatted to two decimal places.
func (a MolarMass) String() string {
	return a.ToString(MolarMassKilogramPerMole, 2)
}

// ToString formats the MolarMass value as a string with the specified unit and fractional digits.
// It converts the MolarMass to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the MolarMass value will be converted (e.g., KilogramPerMole))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the MolarMass with the unit abbreviation.
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

// Equals checks if the given MolarMass is equal to the current MolarMass.
//
// Parameters:
//    other: The MolarMass to compare against.
//
// Returns:
//    bool: Returns true if both MolarMass are equal, false otherwise.
func (a *MolarMass) Equals(other *MolarMass) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current MolarMass with another MolarMass.
// It returns -1 if the current MolarMass is less than the other MolarMass, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The MolarMass to compare against.
//
// Returns:
//    int: -1 if the current MolarMass is less, 1 if greater, and 0 if equal.
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

// Add adds the given MolarMass to the current MolarMass and returns the result.
// The result is a new MolarMass instance with the sum of the values.
//
// Parameters:
//    other: The MolarMass to add to the current MolarMass.
//
// Returns:
//    *MolarMass: A new MolarMass instance representing the sum of both MolarMass.
func (a *MolarMass) Add(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given MolarMass from the current MolarMass and returns the result.
// The result is a new MolarMass instance with the difference of the values.
//
// Parameters:
//    other: The MolarMass to subtract from the current MolarMass.
//
// Returns:
//    *MolarMass: A new MolarMass instance representing the difference of both MolarMass.
func (a *MolarMass) Subtract(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current MolarMass by the given MolarMass and returns the result.
// The result is a new MolarMass instance with the product of the values.
//
// Parameters:
//    other: The MolarMass to multiply with the current MolarMass.
//
// Returns:
//    *MolarMass: A new MolarMass instance representing the product of both MolarMass.
func (a *MolarMass) Multiply(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value * other.BaseValue()}
}

// Divide divides the current MolarMass by the given MolarMass and returns the result.
// The result is a new MolarMass instance with the quotient of the values.
//
// Parameters:
//    other: The MolarMass to divide the current MolarMass by.
//
// Returns:
//    *MolarMass: A new MolarMass instance representing the quotient of both MolarMass.
func (a *MolarMass) Divide(other *MolarMass) *MolarMass {
	return &MolarMass{value: a.value / other.BaseValue()}
}