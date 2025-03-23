// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricSusceptanceUnits defines various units of ElectricSusceptance.
type ElectricSusceptanceUnits string

const (
    
        // 
        ElectricSusceptanceSiemens ElectricSusceptanceUnits = "Siemens"
        // 
        ElectricSusceptanceMho ElectricSusceptanceUnits = "Mho"
        // 
        ElectricSusceptanceNanosiemens ElectricSusceptanceUnits = "Nanosiemens"
        // 
        ElectricSusceptanceMicrosiemens ElectricSusceptanceUnits = "Microsiemens"
        // 
        ElectricSusceptanceMillisiemens ElectricSusceptanceUnits = "Millisiemens"
        // 
        ElectricSusceptanceKilosiemens ElectricSusceptanceUnits = "Kilosiemens"
        // 
        ElectricSusceptanceMegasiemens ElectricSusceptanceUnits = "Megasiemens"
        // 
        ElectricSusceptanceGigasiemens ElectricSusceptanceUnits = "Gigasiemens"
        // 
        ElectricSusceptanceTerasiemens ElectricSusceptanceUnits = "Terasiemens"
        // 
        ElectricSusceptanceNanomho ElectricSusceptanceUnits = "Nanomho"
        // 
        ElectricSusceptanceMicromho ElectricSusceptanceUnits = "Micromho"
        // 
        ElectricSusceptanceMillimho ElectricSusceptanceUnits = "Millimho"
        // 
        ElectricSusceptanceKilomho ElectricSusceptanceUnits = "Kilomho"
        // 
        ElectricSusceptanceMegamho ElectricSusceptanceUnits = "Megamho"
        // 
        ElectricSusceptanceGigamho ElectricSusceptanceUnits = "Gigamho"
        // 
        ElectricSusceptanceTeramho ElectricSusceptanceUnits = "Teramho"
)

var internalElectricSusceptanceUnitsMap = map[ElectricSusceptanceUnits]bool{
	
	ElectricSusceptanceSiemens: true,
	ElectricSusceptanceMho: true,
	ElectricSusceptanceNanosiemens: true,
	ElectricSusceptanceMicrosiemens: true,
	ElectricSusceptanceMillisiemens: true,
	ElectricSusceptanceKilosiemens: true,
	ElectricSusceptanceMegasiemens: true,
	ElectricSusceptanceGigasiemens: true,
	ElectricSusceptanceTerasiemens: true,
	ElectricSusceptanceNanomho: true,
	ElectricSusceptanceMicromho: true,
	ElectricSusceptanceMillimho: true,
	ElectricSusceptanceKilomho: true,
	ElectricSusceptanceMegamho: true,
	ElectricSusceptanceGigamho: true,
	ElectricSusceptanceTeramho: true,
}

// ElectricSusceptanceDto represents a ElectricSusceptance measurement with a numerical value and its corresponding unit.
type ElectricSusceptanceDto struct {
    // Value is the numerical representation of the ElectricSusceptance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricSusceptance, as defined in the ElectricSusceptanceUnits enumeration.
	Unit  ElectricSusceptanceUnits `json:"unit" validate:"required,oneof=Siemens Mho Nanosiemens Microsiemens Millisiemens Kilosiemens Megasiemens Gigasiemens Terasiemens Nanomho Micromho Millimho Kilomho Megamho Gigamho Teramho"`
}

// ElectricSusceptanceDtoFactory groups methods for creating and serializing ElectricSusceptanceDto objects.
type ElectricSusceptanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricSusceptanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricSusceptanceDtoFactory) FromJSON(data []byte) (*ElectricSusceptanceDto, error) {
	a := ElectricSusceptanceDto{}

    // Parse JSON into ElectricSusceptanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a ElectricSusceptanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricSusceptanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricSusceptance represents a measurement in a of ElectricSusceptance.
//
// Electrical susceptance is the imaginary part of admittance, where the real part is conductance.
type ElectricSusceptance struct {
	// value is the base measurement stored internally.
	value       float64
    
    siemensLazy *float64 
    mhosLazy *float64 
    nanosiemensLazy *float64 
    microsiemensLazy *float64 
    millisiemensLazy *float64 
    kilosiemensLazy *float64 
    megasiemensLazy *float64 
    gigasiemensLazy *float64 
    terasiemensLazy *float64 
    nanomhosLazy *float64 
    micromhosLazy *float64 
    millimhosLazy *float64 
    kilomhosLazy *float64 
    megamhosLazy *float64 
    gigamhosLazy *float64 
    teramhosLazy *float64 
}

// ElectricSusceptanceFactory groups methods for creating ElectricSusceptance instances.
type ElectricSusceptanceFactory struct{}

// CreateElectricSusceptance creates a new ElectricSusceptance instance from the given value and unit.
func (uf ElectricSusceptanceFactory) CreateElectricSusceptance(value float64, unit ElectricSusceptanceUnits) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, unit)
}

// FromDto converts a ElectricSusceptanceDto to a ElectricSusceptance instance.
func (uf ElectricSusceptanceFactory) FromDto(dto ElectricSusceptanceDto) (*ElectricSusceptance, error) {
	return newElectricSusceptance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricSusceptance instance.
func (uf ElectricSusceptanceFactory) FromDtoJSON(data []byte) (*ElectricSusceptance, error) {
	unitDto, err := ElectricSusceptanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricSusceptanceDto from JSON: %w", err)
	}
	return ElectricSusceptanceFactory{}.FromDto(*unitDto)
}


// FromSiemens creates a new ElectricSusceptance instance from a value in Siemens.
func (uf ElectricSusceptanceFactory) FromSiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceSiemens)
}

// FromMhos creates a new ElectricSusceptance instance from a value in Mhos.
func (uf ElectricSusceptanceFactory) FromMhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMho)
}

// FromNanosiemens creates a new ElectricSusceptance instance from a value in Nanosiemens.
func (uf ElectricSusceptanceFactory) FromNanosiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceNanosiemens)
}

// FromMicrosiemens creates a new ElectricSusceptance instance from a value in Microsiemens.
func (uf ElectricSusceptanceFactory) FromMicrosiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMicrosiemens)
}

// FromMillisiemens creates a new ElectricSusceptance instance from a value in Millisiemens.
func (uf ElectricSusceptanceFactory) FromMillisiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMillisiemens)
}

// FromKilosiemens creates a new ElectricSusceptance instance from a value in Kilosiemens.
func (uf ElectricSusceptanceFactory) FromKilosiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceKilosiemens)
}

// FromMegasiemens creates a new ElectricSusceptance instance from a value in Megasiemens.
func (uf ElectricSusceptanceFactory) FromMegasiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMegasiemens)
}

// FromGigasiemens creates a new ElectricSusceptance instance from a value in Gigasiemens.
func (uf ElectricSusceptanceFactory) FromGigasiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceGigasiemens)
}

// FromTerasiemens creates a new ElectricSusceptance instance from a value in Terasiemens.
func (uf ElectricSusceptanceFactory) FromTerasiemens(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceTerasiemens)
}

// FromNanomhos creates a new ElectricSusceptance instance from a value in Nanomhos.
func (uf ElectricSusceptanceFactory) FromNanomhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceNanomho)
}

// FromMicromhos creates a new ElectricSusceptance instance from a value in Micromhos.
func (uf ElectricSusceptanceFactory) FromMicromhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMicromho)
}

// FromMillimhos creates a new ElectricSusceptance instance from a value in Millimhos.
func (uf ElectricSusceptanceFactory) FromMillimhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMillimho)
}

// FromKilomhos creates a new ElectricSusceptance instance from a value in Kilomhos.
func (uf ElectricSusceptanceFactory) FromKilomhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceKilomho)
}

// FromMegamhos creates a new ElectricSusceptance instance from a value in Megamhos.
func (uf ElectricSusceptanceFactory) FromMegamhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceMegamho)
}

// FromGigamhos creates a new ElectricSusceptance instance from a value in Gigamhos.
func (uf ElectricSusceptanceFactory) FromGigamhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceGigamho)
}

// FromTeramhos creates a new ElectricSusceptance instance from a value in Teramhos.
func (uf ElectricSusceptanceFactory) FromTeramhos(value float64) (*ElectricSusceptance, error) {
	return newElectricSusceptance(value, ElectricSusceptanceTeramho)
}


// newElectricSusceptance creates a new ElectricSusceptance.
func newElectricSusceptance(value float64, fromUnit ElectricSusceptanceUnits) (*ElectricSusceptance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalElectricSusceptanceUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in ElectricSusceptanceUnits", fromUnit)
	}
	a := &ElectricSusceptance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricSusceptance in Siemens unit (the base/default quantity).
func (a *ElectricSusceptance) BaseValue() float64 {
	return a.value
}


// Siemens returns the ElectricSusceptance value in Siemens.
//
// 
func (a *ElectricSusceptance) Siemens() float64 {
	if a.siemensLazy != nil {
		return *a.siemensLazy
	}
	siemens := a.convertFromBase(ElectricSusceptanceSiemens)
	a.siemensLazy = &siemens
	return siemens
}

// Mhos returns the ElectricSusceptance value in Mhos.
//
// 
func (a *ElectricSusceptance) Mhos() float64 {
	if a.mhosLazy != nil {
		return *a.mhosLazy
	}
	mhos := a.convertFromBase(ElectricSusceptanceMho)
	a.mhosLazy = &mhos
	return mhos
}

// Nanosiemens returns the ElectricSusceptance value in Nanosiemens.
//
// 
func (a *ElectricSusceptance) Nanosiemens() float64 {
	if a.nanosiemensLazy != nil {
		return *a.nanosiemensLazy
	}
	nanosiemens := a.convertFromBase(ElectricSusceptanceNanosiemens)
	a.nanosiemensLazy = &nanosiemens
	return nanosiemens
}

// Microsiemens returns the ElectricSusceptance value in Microsiemens.
//
// 
func (a *ElectricSusceptance) Microsiemens() float64 {
	if a.microsiemensLazy != nil {
		return *a.microsiemensLazy
	}
	microsiemens := a.convertFromBase(ElectricSusceptanceMicrosiemens)
	a.microsiemensLazy = &microsiemens
	return microsiemens
}

// Millisiemens returns the ElectricSusceptance value in Millisiemens.
//
// 
func (a *ElectricSusceptance) Millisiemens() float64 {
	if a.millisiemensLazy != nil {
		return *a.millisiemensLazy
	}
	millisiemens := a.convertFromBase(ElectricSusceptanceMillisiemens)
	a.millisiemensLazy = &millisiemens
	return millisiemens
}

// Kilosiemens returns the ElectricSusceptance value in Kilosiemens.
//
// 
func (a *ElectricSusceptance) Kilosiemens() float64 {
	if a.kilosiemensLazy != nil {
		return *a.kilosiemensLazy
	}
	kilosiemens := a.convertFromBase(ElectricSusceptanceKilosiemens)
	a.kilosiemensLazy = &kilosiemens
	return kilosiemens
}

// Megasiemens returns the ElectricSusceptance value in Megasiemens.
//
// 
func (a *ElectricSusceptance) Megasiemens() float64 {
	if a.megasiemensLazy != nil {
		return *a.megasiemensLazy
	}
	megasiemens := a.convertFromBase(ElectricSusceptanceMegasiemens)
	a.megasiemensLazy = &megasiemens
	return megasiemens
}

// Gigasiemens returns the ElectricSusceptance value in Gigasiemens.
//
// 
func (a *ElectricSusceptance) Gigasiemens() float64 {
	if a.gigasiemensLazy != nil {
		return *a.gigasiemensLazy
	}
	gigasiemens := a.convertFromBase(ElectricSusceptanceGigasiemens)
	a.gigasiemensLazy = &gigasiemens
	return gigasiemens
}

// Terasiemens returns the ElectricSusceptance value in Terasiemens.
//
// 
func (a *ElectricSusceptance) Terasiemens() float64 {
	if a.terasiemensLazy != nil {
		return *a.terasiemensLazy
	}
	terasiemens := a.convertFromBase(ElectricSusceptanceTerasiemens)
	a.terasiemensLazy = &terasiemens
	return terasiemens
}

// Nanomhos returns the ElectricSusceptance value in Nanomhos.
//
// 
func (a *ElectricSusceptance) Nanomhos() float64 {
	if a.nanomhosLazy != nil {
		return *a.nanomhosLazy
	}
	nanomhos := a.convertFromBase(ElectricSusceptanceNanomho)
	a.nanomhosLazy = &nanomhos
	return nanomhos
}

// Micromhos returns the ElectricSusceptance value in Micromhos.
//
// 
func (a *ElectricSusceptance) Micromhos() float64 {
	if a.micromhosLazy != nil {
		return *a.micromhosLazy
	}
	micromhos := a.convertFromBase(ElectricSusceptanceMicromho)
	a.micromhosLazy = &micromhos
	return micromhos
}

// Millimhos returns the ElectricSusceptance value in Millimhos.
//
// 
func (a *ElectricSusceptance) Millimhos() float64 {
	if a.millimhosLazy != nil {
		return *a.millimhosLazy
	}
	millimhos := a.convertFromBase(ElectricSusceptanceMillimho)
	a.millimhosLazy = &millimhos
	return millimhos
}

// Kilomhos returns the ElectricSusceptance value in Kilomhos.
//
// 
func (a *ElectricSusceptance) Kilomhos() float64 {
	if a.kilomhosLazy != nil {
		return *a.kilomhosLazy
	}
	kilomhos := a.convertFromBase(ElectricSusceptanceKilomho)
	a.kilomhosLazy = &kilomhos
	return kilomhos
}

// Megamhos returns the ElectricSusceptance value in Megamhos.
//
// 
func (a *ElectricSusceptance) Megamhos() float64 {
	if a.megamhosLazy != nil {
		return *a.megamhosLazy
	}
	megamhos := a.convertFromBase(ElectricSusceptanceMegamho)
	a.megamhosLazy = &megamhos
	return megamhos
}

// Gigamhos returns the ElectricSusceptance value in Gigamhos.
//
// 
func (a *ElectricSusceptance) Gigamhos() float64 {
	if a.gigamhosLazy != nil {
		return *a.gigamhosLazy
	}
	gigamhos := a.convertFromBase(ElectricSusceptanceGigamho)
	a.gigamhosLazy = &gigamhos
	return gigamhos
}

// Teramhos returns the ElectricSusceptance value in Teramhos.
//
// 
func (a *ElectricSusceptance) Teramhos() float64 {
	if a.teramhosLazy != nil {
		return *a.teramhosLazy
	}
	teramhos := a.convertFromBase(ElectricSusceptanceTeramho)
	a.teramhosLazy = &teramhos
	return teramhos
}


// ToDto creates a ElectricSusceptanceDto representation from the ElectricSusceptance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricSusceptance) ToDto(holdInUnit *ElectricSusceptanceUnits) ElectricSusceptanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricSusceptanceSiemens // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricSusceptanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricSusceptance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricSusceptance) ToDtoJSON(holdInUnit *ElectricSusceptanceUnits) ([]byte, error) {
	// Convert to ElectricSusceptanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricSusceptance to a specific unit value.
// The function uses the provided unit type (ElectricSusceptanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricSusceptance) Convert(toUnit ElectricSusceptanceUnits) float64 {
	switch toUnit { 
    case ElectricSusceptanceSiemens:
		return a.Siemens()
    case ElectricSusceptanceMho:
		return a.Mhos()
    case ElectricSusceptanceNanosiemens:
		return a.Nanosiemens()
    case ElectricSusceptanceMicrosiemens:
		return a.Microsiemens()
    case ElectricSusceptanceMillisiemens:
		return a.Millisiemens()
    case ElectricSusceptanceKilosiemens:
		return a.Kilosiemens()
    case ElectricSusceptanceMegasiemens:
		return a.Megasiemens()
    case ElectricSusceptanceGigasiemens:
		return a.Gigasiemens()
    case ElectricSusceptanceTerasiemens:
		return a.Terasiemens()
    case ElectricSusceptanceNanomho:
		return a.Nanomhos()
    case ElectricSusceptanceMicromho:
		return a.Micromhos()
    case ElectricSusceptanceMillimho:
		return a.Millimhos()
    case ElectricSusceptanceKilomho:
		return a.Kilomhos()
    case ElectricSusceptanceMegamho:
		return a.Megamhos()
    case ElectricSusceptanceGigamho:
		return a.Gigamhos()
    case ElectricSusceptanceTeramho:
		return a.Teramhos()
	default:
		return math.NaN()
	}
}

func (a *ElectricSusceptance) convertFromBase(toUnit ElectricSusceptanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricSusceptanceSiemens:
		return (value) 
	case ElectricSusceptanceMho:
		return (value) 
	case ElectricSusceptanceNanosiemens:
		return ((value) / 1e-09) 
	case ElectricSusceptanceMicrosiemens:
		return ((value) / 1e-06) 
	case ElectricSusceptanceMillisiemens:
		return ((value) / 0.001) 
	case ElectricSusceptanceKilosiemens:
		return ((value) / 1000.0) 
	case ElectricSusceptanceMegasiemens:
		return ((value) / 1000000.0) 
	case ElectricSusceptanceGigasiemens:
		return ((value) / 1000000000.0) 
	case ElectricSusceptanceTerasiemens:
		return ((value) / 1000000000000.0) 
	case ElectricSusceptanceNanomho:
		return ((value) / 1e-09) 
	case ElectricSusceptanceMicromho:
		return ((value) / 1e-06) 
	case ElectricSusceptanceMillimho:
		return ((value) / 0.001) 
	case ElectricSusceptanceKilomho:
		return ((value) / 1000.0) 
	case ElectricSusceptanceMegamho:
		return ((value) / 1000000.0) 
	case ElectricSusceptanceGigamho:
		return ((value) / 1000000000.0) 
	case ElectricSusceptanceTeramho:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricSusceptance) convertToBase(value float64, fromUnit ElectricSusceptanceUnits) float64 {
	switch fromUnit { 
	case ElectricSusceptanceSiemens:
		return (value) 
	case ElectricSusceptanceMho:
		return (value) 
	case ElectricSusceptanceNanosiemens:
		return ((value) * 1e-09) 
	case ElectricSusceptanceMicrosiemens:
		return ((value) * 1e-06) 
	case ElectricSusceptanceMillisiemens:
		return ((value) * 0.001) 
	case ElectricSusceptanceKilosiemens:
		return ((value) * 1000.0) 
	case ElectricSusceptanceMegasiemens:
		return ((value) * 1000000.0) 
	case ElectricSusceptanceGigasiemens:
		return ((value) * 1000000000.0) 
	case ElectricSusceptanceTerasiemens:
		return ((value) * 1000000000000.0) 
	case ElectricSusceptanceNanomho:
		return ((value) * 1e-09) 
	case ElectricSusceptanceMicromho:
		return ((value) * 1e-06) 
	case ElectricSusceptanceMillimho:
		return ((value) * 0.001) 
	case ElectricSusceptanceKilomho:
		return ((value) * 1000.0) 
	case ElectricSusceptanceMegamho:
		return ((value) * 1000000.0) 
	case ElectricSusceptanceGigamho:
		return ((value) * 1000000000.0) 
	case ElectricSusceptanceTeramho:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricSusceptance in the default unit (Siemens),
// formatted to two decimal places.
func (a ElectricSusceptance) String() string {
	return a.ToString(ElectricSusceptanceSiemens, 2)
}

// ToString formats the ElectricSusceptance value as a string with the specified unit and fractional digits.
// It converts the ElectricSusceptance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricSusceptance value will be converted (e.g., Siemens))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricSusceptance with the unit abbreviation.
func (a *ElectricSusceptance) ToString(unit ElectricSusceptanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricSusceptanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricSusceptanceAbbreviation(unit))
}

// Equals checks if the given ElectricSusceptance is equal to the current ElectricSusceptance.
//
// Parameters:
//    other: The ElectricSusceptance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricSusceptance are equal, false otherwise.
func (a *ElectricSusceptance) Equals(other *ElectricSusceptance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricSusceptance with another ElectricSusceptance.
// It returns -1 if the current ElectricSusceptance is less than the other ElectricSusceptance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricSusceptance to compare against.
//
// Returns:
//    int: -1 if the current ElectricSusceptance is less, 1 if greater, and 0 if equal.
func (a *ElectricSusceptance) CompareTo(other *ElectricSusceptance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricSusceptance to the current ElectricSusceptance and returns the result.
// The result is a new ElectricSusceptance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricSusceptance to add to the current ElectricSusceptance.
//
// Returns:
//    *ElectricSusceptance: A new ElectricSusceptance instance representing the sum of both ElectricSusceptance.
func (a *ElectricSusceptance) Add(other *ElectricSusceptance) *ElectricSusceptance {
	return &ElectricSusceptance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricSusceptance from the current ElectricSusceptance and returns the result.
// The result is a new ElectricSusceptance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricSusceptance to subtract from the current ElectricSusceptance.
//
// Returns:
//    *ElectricSusceptance: A new ElectricSusceptance instance representing the difference of both ElectricSusceptance.
func (a *ElectricSusceptance) Subtract(other *ElectricSusceptance) *ElectricSusceptance {
	return &ElectricSusceptance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricSusceptance by the given ElectricSusceptance and returns the result.
// The result is a new ElectricSusceptance instance with the product of the values.
//
// Parameters:
//    other: The ElectricSusceptance to multiply with the current ElectricSusceptance.
//
// Returns:
//    *ElectricSusceptance: A new ElectricSusceptance instance representing the product of both ElectricSusceptance.
func (a *ElectricSusceptance) Multiply(other *ElectricSusceptance) *ElectricSusceptance {
	return &ElectricSusceptance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricSusceptance by the given ElectricSusceptance and returns the result.
// The result is a new ElectricSusceptance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricSusceptance to divide the current ElectricSusceptance by.
//
// Returns:
//    *ElectricSusceptance: A new ElectricSusceptance instance representing the quotient of both ElectricSusceptance.
func (a *ElectricSusceptance) Divide(other *ElectricSusceptance) *ElectricSusceptance {
	return &ElectricSusceptance{value: a.value / other.BaseValue()}
}

// GetElectricSusceptanceAbbreviation gets the unit abbreviation.
func GetElectricSusceptanceAbbreviation(unit ElectricSusceptanceUnits) string {
	switch unit { 
	case ElectricSusceptanceSiemens:
		return "S" 
	case ElectricSusceptanceMho:
		return "℧" 
	case ElectricSusceptanceNanosiemens:
		return "nS" 
	case ElectricSusceptanceMicrosiemens:
		return "μS" 
	case ElectricSusceptanceMillisiemens:
		return "mS" 
	case ElectricSusceptanceKilosiemens:
		return "kS" 
	case ElectricSusceptanceMegasiemens:
		return "MS" 
	case ElectricSusceptanceGigasiemens:
		return "GS" 
	case ElectricSusceptanceTerasiemens:
		return "TS" 
	case ElectricSusceptanceNanomho:
		return "n℧" 
	case ElectricSusceptanceMicromho:
		return "μ℧" 
	case ElectricSusceptanceMillimho:
		return "m℧" 
	case ElectricSusceptanceKilomho:
		return "k℧" 
	case ElectricSusceptanceMegamho:
		return "M℧" 
	case ElectricSusceptanceGigamho:
		return "G℧" 
	case ElectricSusceptanceTeramho:
		return "T℧" 
	default:
		return ""
	}
}