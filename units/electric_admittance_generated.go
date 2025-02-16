// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// ElectricAdmittanceUnits defines various units of ElectricAdmittance.
type ElectricAdmittanceUnits string

const (
    
        // 
        ElectricAdmittanceSiemens ElectricAdmittanceUnits = "Siemens"
        // 
        ElectricAdmittanceMho ElectricAdmittanceUnits = "Mho"
        // 
        ElectricAdmittanceNanosiemens ElectricAdmittanceUnits = "Nanosiemens"
        // 
        ElectricAdmittanceMicrosiemens ElectricAdmittanceUnits = "Microsiemens"
        // 
        ElectricAdmittanceMillisiemens ElectricAdmittanceUnits = "Millisiemens"
        // 
        ElectricAdmittanceKilosiemens ElectricAdmittanceUnits = "Kilosiemens"
        // 
        ElectricAdmittanceMegasiemens ElectricAdmittanceUnits = "Megasiemens"
        // 
        ElectricAdmittanceGigasiemens ElectricAdmittanceUnits = "Gigasiemens"
        // 
        ElectricAdmittanceTerasiemens ElectricAdmittanceUnits = "Terasiemens"
        // 
        ElectricAdmittanceNanomho ElectricAdmittanceUnits = "Nanomho"
        // 
        ElectricAdmittanceMicromho ElectricAdmittanceUnits = "Micromho"
        // 
        ElectricAdmittanceMillimho ElectricAdmittanceUnits = "Millimho"
        // 
        ElectricAdmittanceKilomho ElectricAdmittanceUnits = "Kilomho"
        // 
        ElectricAdmittanceMegamho ElectricAdmittanceUnits = "Megamho"
        // 
        ElectricAdmittanceGigamho ElectricAdmittanceUnits = "Gigamho"
        // 
        ElectricAdmittanceTeramho ElectricAdmittanceUnits = "Teramho"
)

// ElectricAdmittanceDto represents a ElectricAdmittance measurement with a numerical value and its corresponding unit.
type ElectricAdmittanceDto struct {
    // Value is the numerical representation of the ElectricAdmittance.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the ElectricAdmittance, as defined in the ElectricAdmittanceUnits enumeration.
	Unit  ElectricAdmittanceUnits `json:"unit"`
}

// ElectricAdmittanceDtoFactory groups methods for creating and serializing ElectricAdmittanceDto objects.
type ElectricAdmittanceDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a ElectricAdmittanceDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf ElectricAdmittanceDtoFactory) FromJSON(data []byte) (*ElectricAdmittanceDto, error) {
	a := ElectricAdmittanceDto{}

    // Parse JSON into ElectricAdmittanceDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a ElectricAdmittanceDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a ElectricAdmittanceDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// ElectricAdmittance represents a measurement in a of ElectricAdmittance.
//
// Electric admittance is a measure of how easily a circuit or device will allow a current to flow by the combined effect of conductance and susceptance in a circuit. It is defined as the inverse of impedance. The SI unit of admittance is the siemens (symbol S).
type ElectricAdmittance struct {
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

// ElectricAdmittanceFactory groups methods for creating ElectricAdmittance instances.
type ElectricAdmittanceFactory struct{}

// CreateElectricAdmittance creates a new ElectricAdmittance instance from the given value and unit.
func (uf ElectricAdmittanceFactory) CreateElectricAdmittance(value float64, unit ElectricAdmittanceUnits) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, unit)
}

// FromDto converts a ElectricAdmittanceDto to a ElectricAdmittance instance.
func (uf ElectricAdmittanceFactory) FromDto(dto ElectricAdmittanceDto) (*ElectricAdmittance, error) {
	return newElectricAdmittance(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a ElectricAdmittance instance.
func (uf ElectricAdmittanceFactory) FromDtoJSON(data []byte) (*ElectricAdmittance, error) {
	unitDto, err := ElectricAdmittanceDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ElectricAdmittanceDto from JSON: %w", err)
	}
	return ElectricAdmittanceFactory{}.FromDto(*unitDto)
}


// FromSiemens creates a new ElectricAdmittance instance from a value in Siemens.
func (uf ElectricAdmittanceFactory) FromSiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceSiemens)
}

// FromMhos creates a new ElectricAdmittance instance from a value in Mhos.
func (uf ElectricAdmittanceFactory) FromMhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMho)
}

// FromNanosiemens creates a new ElectricAdmittance instance from a value in Nanosiemens.
func (uf ElectricAdmittanceFactory) FromNanosiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceNanosiemens)
}

// FromMicrosiemens creates a new ElectricAdmittance instance from a value in Microsiemens.
func (uf ElectricAdmittanceFactory) FromMicrosiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMicrosiemens)
}

// FromMillisiemens creates a new ElectricAdmittance instance from a value in Millisiemens.
func (uf ElectricAdmittanceFactory) FromMillisiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMillisiemens)
}

// FromKilosiemens creates a new ElectricAdmittance instance from a value in Kilosiemens.
func (uf ElectricAdmittanceFactory) FromKilosiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceKilosiemens)
}

// FromMegasiemens creates a new ElectricAdmittance instance from a value in Megasiemens.
func (uf ElectricAdmittanceFactory) FromMegasiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMegasiemens)
}

// FromGigasiemens creates a new ElectricAdmittance instance from a value in Gigasiemens.
func (uf ElectricAdmittanceFactory) FromGigasiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceGigasiemens)
}

// FromTerasiemens creates a new ElectricAdmittance instance from a value in Terasiemens.
func (uf ElectricAdmittanceFactory) FromTerasiemens(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceTerasiemens)
}

// FromNanomhos creates a new ElectricAdmittance instance from a value in Nanomhos.
func (uf ElectricAdmittanceFactory) FromNanomhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceNanomho)
}

// FromMicromhos creates a new ElectricAdmittance instance from a value in Micromhos.
func (uf ElectricAdmittanceFactory) FromMicromhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMicromho)
}

// FromMillimhos creates a new ElectricAdmittance instance from a value in Millimhos.
func (uf ElectricAdmittanceFactory) FromMillimhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMillimho)
}

// FromKilomhos creates a new ElectricAdmittance instance from a value in Kilomhos.
func (uf ElectricAdmittanceFactory) FromKilomhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceKilomho)
}

// FromMegamhos creates a new ElectricAdmittance instance from a value in Megamhos.
func (uf ElectricAdmittanceFactory) FromMegamhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceMegamho)
}

// FromGigamhos creates a new ElectricAdmittance instance from a value in Gigamhos.
func (uf ElectricAdmittanceFactory) FromGigamhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceGigamho)
}

// FromTeramhos creates a new ElectricAdmittance instance from a value in Teramhos.
func (uf ElectricAdmittanceFactory) FromTeramhos(value float64) (*ElectricAdmittance, error) {
	return newElectricAdmittance(value, ElectricAdmittanceTeramho)
}


// newElectricAdmittance creates a new ElectricAdmittance.
func newElectricAdmittance(value float64, fromUnit ElectricAdmittanceUnits) (*ElectricAdmittance, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &ElectricAdmittance{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of ElectricAdmittance in Siemens unit (the base/default quantity).
func (a *ElectricAdmittance) BaseValue() float64 {
	return a.value
}


// Siemens returns the ElectricAdmittance value in Siemens.
//
// 
func (a *ElectricAdmittance) Siemens() float64 {
	if a.siemensLazy != nil {
		return *a.siemensLazy
	}
	siemens := a.convertFromBase(ElectricAdmittanceSiemens)
	a.siemensLazy = &siemens
	return siemens
}

// Mhos returns the ElectricAdmittance value in Mhos.
//
// 
func (a *ElectricAdmittance) Mhos() float64 {
	if a.mhosLazy != nil {
		return *a.mhosLazy
	}
	mhos := a.convertFromBase(ElectricAdmittanceMho)
	a.mhosLazy = &mhos
	return mhos
}

// Nanosiemens returns the ElectricAdmittance value in Nanosiemens.
//
// 
func (a *ElectricAdmittance) Nanosiemens() float64 {
	if a.nanosiemensLazy != nil {
		return *a.nanosiemensLazy
	}
	nanosiemens := a.convertFromBase(ElectricAdmittanceNanosiemens)
	a.nanosiemensLazy = &nanosiemens
	return nanosiemens
}

// Microsiemens returns the ElectricAdmittance value in Microsiemens.
//
// 
func (a *ElectricAdmittance) Microsiemens() float64 {
	if a.microsiemensLazy != nil {
		return *a.microsiemensLazy
	}
	microsiemens := a.convertFromBase(ElectricAdmittanceMicrosiemens)
	a.microsiemensLazy = &microsiemens
	return microsiemens
}

// Millisiemens returns the ElectricAdmittance value in Millisiemens.
//
// 
func (a *ElectricAdmittance) Millisiemens() float64 {
	if a.millisiemensLazy != nil {
		return *a.millisiemensLazy
	}
	millisiemens := a.convertFromBase(ElectricAdmittanceMillisiemens)
	a.millisiemensLazy = &millisiemens
	return millisiemens
}

// Kilosiemens returns the ElectricAdmittance value in Kilosiemens.
//
// 
func (a *ElectricAdmittance) Kilosiemens() float64 {
	if a.kilosiemensLazy != nil {
		return *a.kilosiemensLazy
	}
	kilosiemens := a.convertFromBase(ElectricAdmittanceKilosiemens)
	a.kilosiemensLazy = &kilosiemens
	return kilosiemens
}

// Megasiemens returns the ElectricAdmittance value in Megasiemens.
//
// 
func (a *ElectricAdmittance) Megasiemens() float64 {
	if a.megasiemensLazy != nil {
		return *a.megasiemensLazy
	}
	megasiemens := a.convertFromBase(ElectricAdmittanceMegasiemens)
	a.megasiemensLazy = &megasiemens
	return megasiemens
}

// Gigasiemens returns the ElectricAdmittance value in Gigasiemens.
//
// 
func (a *ElectricAdmittance) Gigasiemens() float64 {
	if a.gigasiemensLazy != nil {
		return *a.gigasiemensLazy
	}
	gigasiemens := a.convertFromBase(ElectricAdmittanceGigasiemens)
	a.gigasiemensLazy = &gigasiemens
	return gigasiemens
}

// Terasiemens returns the ElectricAdmittance value in Terasiemens.
//
// 
func (a *ElectricAdmittance) Terasiemens() float64 {
	if a.terasiemensLazy != nil {
		return *a.terasiemensLazy
	}
	terasiemens := a.convertFromBase(ElectricAdmittanceTerasiemens)
	a.terasiemensLazy = &terasiemens
	return terasiemens
}

// Nanomhos returns the ElectricAdmittance value in Nanomhos.
//
// 
func (a *ElectricAdmittance) Nanomhos() float64 {
	if a.nanomhosLazy != nil {
		return *a.nanomhosLazy
	}
	nanomhos := a.convertFromBase(ElectricAdmittanceNanomho)
	a.nanomhosLazy = &nanomhos
	return nanomhos
}

// Micromhos returns the ElectricAdmittance value in Micromhos.
//
// 
func (a *ElectricAdmittance) Micromhos() float64 {
	if a.micromhosLazy != nil {
		return *a.micromhosLazy
	}
	micromhos := a.convertFromBase(ElectricAdmittanceMicromho)
	a.micromhosLazy = &micromhos
	return micromhos
}

// Millimhos returns the ElectricAdmittance value in Millimhos.
//
// 
func (a *ElectricAdmittance) Millimhos() float64 {
	if a.millimhosLazy != nil {
		return *a.millimhosLazy
	}
	millimhos := a.convertFromBase(ElectricAdmittanceMillimho)
	a.millimhosLazy = &millimhos
	return millimhos
}

// Kilomhos returns the ElectricAdmittance value in Kilomhos.
//
// 
func (a *ElectricAdmittance) Kilomhos() float64 {
	if a.kilomhosLazy != nil {
		return *a.kilomhosLazy
	}
	kilomhos := a.convertFromBase(ElectricAdmittanceKilomho)
	a.kilomhosLazy = &kilomhos
	return kilomhos
}

// Megamhos returns the ElectricAdmittance value in Megamhos.
//
// 
func (a *ElectricAdmittance) Megamhos() float64 {
	if a.megamhosLazy != nil {
		return *a.megamhosLazy
	}
	megamhos := a.convertFromBase(ElectricAdmittanceMegamho)
	a.megamhosLazy = &megamhos
	return megamhos
}

// Gigamhos returns the ElectricAdmittance value in Gigamhos.
//
// 
func (a *ElectricAdmittance) Gigamhos() float64 {
	if a.gigamhosLazy != nil {
		return *a.gigamhosLazy
	}
	gigamhos := a.convertFromBase(ElectricAdmittanceGigamho)
	a.gigamhosLazy = &gigamhos
	return gigamhos
}

// Teramhos returns the ElectricAdmittance value in Teramhos.
//
// 
func (a *ElectricAdmittance) Teramhos() float64 {
	if a.teramhosLazy != nil {
		return *a.teramhosLazy
	}
	teramhos := a.convertFromBase(ElectricAdmittanceTeramho)
	a.teramhosLazy = &teramhos
	return teramhos
}


// ToDto creates a ElectricAdmittanceDto representation from the ElectricAdmittance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricAdmittance) ToDto(holdInUnit *ElectricAdmittanceUnits) ElectricAdmittanceDto {
	if holdInUnit == nil {
		defaultUnit := ElectricAdmittanceSiemens // Default value
		holdInUnit = &defaultUnit
	}

	return ElectricAdmittanceDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the ElectricAdmittance instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Siemens by default.
func (a *ElectricAdmittance) ToDtoJSON(holdInUnit *ElectricAdmittanceUnits) ([]byte, error) {
	// Convert to ElectricAdmittanceDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a ElectricAdmittance to a specific unit value.
// The function uses the provided unit type (ElectricAdmittanceUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *ElectricAdmittance) Convert(toUnit ElectricAdmittanceUnits) float64 {
	switch toUnit { 
    case ElectricAdmittanceSiemens:
		return a.Siemens()
    case ElectricAdmittanceMho:
		return a.Mhos()
    case ElectricAdmittanceNanosiemens:
		return a.Nanosiemens()
    case ElectricAdmittanceMicrosiemens:
		return a.Microsiemens()
    case ElectricAdmittanceMillisiemens:
		return a.Millisiemens()
    case ElectricAdmittanceKilosiemens:
		return a.Kilosiemens()
    case ElectricAdmittanceMegasiemens:
		return a.Megasiemens()
    case ElectricAdmittanceGigasiemens:
		return a.Gigasiemens()
    case ElectricAdmittanceTerasiemens:
		return a.Terasiemens()
    case ElectricAdmittanceNanomho:
		return a.Nanomhos()
    case ElectricAdmittanceMicromho:
		return a.Micromhos()
    case ElectricAdmittanceMillimho:
		return a.Millimhos()
    case ElectricAdmittanceKilomho:
		return a.Kilomhos()
    case ElectricAdmittanceMegamho:
		return a.Megamhos()
    case ElectricAdmittanceGigamho:
		return a.Gigamhos()
    case ElectricAdmittanceTeramho:
		return a.Teramhos()
	default:
		return math.NaN()
	}
}

func (a *ElectricAdmittance) convertFromBase(toUnit ElectricAdmittanceUnits) float64 {
    value := a.value
	switch toUnit { 
	case ElectricAdmittanceSiemens:
		return (value) 
	case ElectricAdmittanceMho:
		return (value) 
	case ElectricAdmittanceNanosiemens:
		return ((value) / 1e-09) 
	case ElectricAdmittanceMicrosiemens:
		return ((value) / 1e-06) 
	case ElectricAdmittanceMillisiemens:
		return ((value) / 0.001) 
	case ElectricAdmittanceKilosiemens:
		return ((value) / 1000.0) 
	case ElectricAdmittanceMegasiemens:
		return ((value) / 1000000.0) 
	case ElectricAdmittanceGigasiemens:
		return ((value) / 1000000000.0) 
	case ElectricAdmittanceTerasiemens:
		return ((value) / 1000000000000.0) 
	case ElectricAdmittanceNanomho:
		return ((value) / 1e-09) 
	case ElectricAdmittanceMicromho:
		return ((value) / 1e-06) 
	case ElectricAdmittanceMillimho:
		return ((value) / 0.001) 
	case ElectricAdmittanceKilomho:
		return ((value) / 1000.0) 
	case ElectricAdmittanceMegamho:
		return ((value) / 1000000.0) 
	case ElectricAdmittanceGigamho:
		return ((value) / 1000000000.0) 
	case ElectricAdmittanceTeramho:
		return ((value) / 1000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *ElectricAdmittance) convertToBase(value float64, fromUnit ElectricAdmittanceUnits) float64 {
	switch fromUnit { 
	case ElectricAdmittanceSiemens:
		return (value) 
	case ElectricAdmittanceMho:
		return (value) 
	case ElectricAdmittanceNanosiemens:
		return ((value) * 1e-09) 
	case ElectricAdmittanceMicrosiemens:
		return ((value) * 1e-06) 
	case ElectricAdmittanceMillisiemens:
		return ((value) * 0.001) 
	case ElectricAdmittanceKilosiemens:
		return ((value) * 1000.0) 
	case ElectricAdmittanceMegasiemens:
		return ((value) * 1000000.0) 
	case ElectricAdmittanceGigasiemens:
		return ((value) * 1000000000.0) 
	case ElectricAdmittanceTerasiemens:
		return ((value) * 1000000000000.0) 
	case ElectricAdmittanceNanomho:
		return ((value) * 1e-09) 
	case ElectricAdmittanceMicromho:
		return ((value) * 1e-06) 
	case ElectricAdmittanceMillimho:
		return ((value) * 0.001) 
	case ElectricAdmittanceKilomho:
		return ((value) * 1000.0) 
	case ElectricAdmittanceMegamho:
		return ((value) * 1000000.0) 
	case ElectricAdmittanceGigamho:
		return ((value) * 1000000000.0) 
	case ElectricAdmittanceTeramho:
		return ((value) * 1000000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the ElectricAdmittance in the default unit (Siemens),
// formatted to two decimal places.
func (a ElectricAdmittance) String() string {
	return a.ToString(ElectricAdmittanceSiemens, 2)
}

// ToString formats the ElectricAdmittance value as a string with the specified unit and fractional digits.
// It converts the ElectricAdmittance to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the ElectricAdmittance value will be converted (e.g., Siemens))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the ElectricAdmittance with the unit abbreviation.
func (a *ElectricAdmittance) ToString(unit ElectricAdmittanceUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetElectricAdmittanceAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetElectricAdmittanceAbbreviation(unit))
}

// Equals checks if the given ElectricAdmittance is equal to the current ElectricAdmittance.
//
// Parameters:
//    other: The ElectricAdmittance to compare against.
//
// Returns:
//    bool: Returns true if both ElectricAdmittance are equal, false otherwise.
func (a *ElectricAdmittance) Equals(other *ElectricAdmittance) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current ElectricAdmittance with another ElectricAdmittance.
// It returns -1 if the current ElectricAdmittance is less than the other ElectricAdmittance, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The ElectricAdmittance to compare against.
//
// Returns:
//    int: -1 if the current ElectricAdmittance is less, 1 if greater, and 0 if equal.
func (a *ElectricAdmittance) CompareTo(other *ElectricAdmittance) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given ElectricAdmittance to the current ElectricAdmittance and returns the result.
// The result is a new ElectricAdmittance instance with the sum of the values.
//
// Parameters:
//    other: The ElectricAdmittance to add to the current ElectricAdmittance.
//
// Returns:
//    *ElectricAdmittance: A new ElectricAdmittance instance representing the sum of both ElectricAdmittance.
func (a *ElectricAdmittance) Add(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given ElectricAdmittance from the current ElectricAdmittance and returns the result.
// The result is a new ElectricAdmittance instance with the difference of the values.
//
// Parameters:
//    other: The ElectricAdmittance to subtract from the current ElectricAdmittance.
//
// Returns:
//    *ElectricAdmittance: A new ElectricAdmittance instance representing the difference of both ElectricAdmittance.
func (a *ElectricAdmittance) Subtract(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current ElectricAdmittance by the given ElectricAdmittance and returns the result.
// The result is a new ElectricAdmittance instance with the product of the values.
//
// Parameters:
//    other: The ElectricAdmittance to multiply with the current ElectricAdmittance.
//
// Returns:
//    *ElectricAdmittance: A new ElectricAdmittance instance representing the product of both ElectricAdmittance.
func (a *ElectricAdmittance) Multiply(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value * other.BaseValue()}
}

// Divide divides the current ElectricAdmittance by the given ElectricAdmittance and returns the result.
// The result is a new ElectricAdmittance instance with the quotient of the values.
//
// Parameters:
//    other: The ElectricAdmittance to divide the current ElectricAdmittance by.
//
// Returns:
//    *ElectricAdmittance: A new ElectricAdmittance instance representing the quotient of both ElectricAdmittance.
func (a *ElectricAdmittance) Divide(other *ElectricAdmittance) *ElectricAdmittance {
	return &ElectricAdmittance{value: a.value / other.BaseValue()}
}

// GetElectricAdmittanceAbbreviation gets the unit abbreviation.
func GetElectricAdmittanceAbbreviation(unit ElectricAdmittanceUnits) string {
	switch unit { 
	case ElectricAdmittanceSiemens:
		return "S" 
	case ElectricAdmittanceMho:
		return "℧" 
	case ElectricAdmittanceNanosiemens:
		return "nS" 
	case ElectricAdmittanceMicrosiemens:
		return "μS" 
	case ElectricAdmittanceMillisiemens:
		return "mS" 
	case ElectricAdmittanceKilosiemens:
		return "kS" 
	case ElectricAdmittanceMegasiemens:
		return "MS" 
	case ElectricAdmittanceGigasiemens:
		return "GS" 
	case ElectricAdmittanceTerasiemens:
		return "TS" 
	case ElectricAdmittanceNanomho:
		return "n℧" 
	case ElectricAdmittanceMicromho:
		return "μ℧" 
	case ElectricAdmittanceMillimho:
		return "m℧" 
	case ElectricAdmittanceKilomho:
		return "k℧" 
	case ElectricAdmittanceMegamho:
		return "M℧" 
	case ElectricAdmittanceGigamho:
		return "G℧" 
	case ElectricAdmittanceTeramho:
		return "T℧" 
	default:
		return ""
	}
}