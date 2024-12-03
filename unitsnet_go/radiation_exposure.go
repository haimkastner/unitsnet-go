package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// RadiationExposureUnits enumeration
type RadiationExposureUnits string

const (
    
        // 
        RadiationExposureCoulombPerKilogram RadiationExposureUnits = "CoulombPerKilogram"
        // 
        RadiationExposureRoentgen RadiationExposureUnits = "Roentgen"
        // 
        RadiationExposurePicocoulombPerKilogram RadiationExposureUnits = "PicocoulombPerKilogram"
        // 
        RadiationExposureNanocoulombPerKilogram RadiationExposureUnits = "NanocoulombPerKilogram"
        // 
        RadiationExposureMicrocoulombPerKilogram RadiationExposureUnits = "MicrocoulombPerKilogram"
        // 
        RadiationExposureMillicoulombPerKilogram RadiationExposureUnits = "MillicoulombPerKilogram"
        // 
        RadiationExposureMicroroentgen RadiationExposureUnits = "Microroentgen"
        // 
        RadiationExposureMilliroentgen RadiationExposureUnits = "Milliroentgen"
)

// RadiationExposureDto represents an RadiationExposure
type RadiationExposureDto struct {
	Value float64
	Unit  RadiationExposureUnits
}

// RadiationExposureDtoFactory struct to group related functions
type RadiationExposureDtoFactory struct{}

func (udf RadiationExposureDtoFactory) FromJSON(data []byte) (*RadiationExposureDto, error) {
	a := RadiationExposureDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a RadiationExposureDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// RadiationExposure struct
type RadiationExposure struct {
	value       float64
    
    coulombs_per_kilogramLazy *float64 
    roentgensLazy *float64 
    picocoulombs_per_kilogramLazy *float64 
    nanocoulombs_per_kilogramLazy *float64 
    microcoulombs_per_kilogramLazy *float64 
    millicoulombs_per_kilogramLazy *float64 
    microroentgensLazy *float64 
    milliroentgensLazy *float64 
}

// RadiationExposureFactory struct to group related functions
type RadiationExposureFactory struct{}

func (uf RadiationExposureFactory) CreateRadiationExposure(value float64, unit RadiationExposureUnits) (*RadiationExposure, error) {
	return newRadiationExposure(value, unit)
}

func (uf RadiationExposureFactory) FromDto(dto RadiationExposureDto) (*RadiationExposure, error) {
	return newRadiationExposure(dto.Value, dto.Unit)
}

func (uf RadiationExposureFactory) FromDtoJSON(data []byte) (*RadiationExposure, error) {
	unitDto, err := RadiationExposureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return RadiationExposureFactory{}.FromDto(*unitDto)
}


// FromCoulombPerKilogram creates a new RadiationExposure instance from CoulombPerKilogram.
func (uf RadiationExposureFactory) FromCoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureCoulombPerKilogram)
}

// FromRoentgen creates a new RadiationExposure instance from Roentgen.
func (uf RadiationExposureFactory) FromRoentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureRoentgen)
}

// FromPicocoulombPerKilogram creates a new RadiationExposure instance from PicocoulombPerKilogram.
func (uf RadiationExposureFactory) FromPicocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposurePicocoulombPerKilogram)
}

// FromNanocoulombPerKilogram creates a new RadiationExposure instance from NanocoulombPerKilogram.
func (uf RadiationExposureFactory) FromNanocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureNanocoulombPerKilogram)
}

// FromMicrocoulombPerKilogram creates a new RadiationExposure instance from MicrocoulombPerKilogram.
func (uf RadiationExposureFactory) FromMicrocoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMicrocoulombPerKilogram)
}

// FromMillicoulombPerKilogram creates a new RadiationExposure instance from MillicoulombPerKilogram.
func (uf RadiationExposureFactory) FromMillicoulombsPerKilogram(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMillicoulombPerKilogram)
}

// FromMicroroentgen creates a new RadiationExposure instance from Microroentgen.
func (uf RadiationExposureFactory) FromMicroroentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMicroroentgen)
}

// FromMilliroentgen creates a new RadiationExposure instance from Milliroentgen.
func (uf RadiationExposureFactory) FromMilliroentgens(value float64) (*RadiationExposure, error) {
	return newRadiationExposure(value, RadiationExposureMilliroentgen)
}




// newRadiationExposure creates a new RadiationExposure.
func newRadiationExposure(value float64, fromUnit RadiationExposureUnits) (*RadiationExposure, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &RadiationExposure{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of RadiationExposure in CoulombPerKilogram.
func (a *RadiationExposure) BaseValue() float64 {
	return a.value
}


// CoulombPerKilogram returns the value in CoulombPerKilogram.
func (a *RadiationExposure) CoulombsPerKilogram() float64 {
	if a.coulombs_per_kilogramLazy != nil {
		return *a.coulombs_per_kilogramLazy
	}
	coulombs_per_kilogram := a.convertFromBase(RadiationExposureCoulombPerKilogram)
	a.coulombs_per_kilogramLazy = &coulombs_per_kilogram
	return coulombs_per_kilogram
}

// Roentgen returns the value in Roentgen.
func (a *RadiationExposure) Roentgens() float64 {
	if a.roentgensLazy != nil {
		return *a.roentgensLazy
	}
	roentgens := a.convertFromBase(RadiationExposureRoentgen)
	a.roentgensLazy = &roentgens
	return roentgens
}

// PicocoulombPerKilogram returns the value in PicocoulombPerKilogram.
func (a *RadiationExposure) PicocoulombsPerKilogram() float64 {
	if a.picocoulombs_per_kilogramLazy != nil {
		return *a.picocoulombs_per_kilogramLazy
	}
	picocoulombs_per_kilogram := a.convertFromBase(RadiationExposurePicocoulombPerKilogram)
	a.picocoulombs_per_kilogramLazy = &picocoulombs_per_kilogram
	return picocoulombs_per_kilogram
}

// NanocoulombPerKilogram returns the value in NanocoulombPerKilogram.
func (a *RadiationExposure) NanocoulombsPerKilogram() float64 {
	if a.nanocoulombs_per_kilogramLazy != nil {
		return *a.nanocoulombs_per_kilogramLazy
	}
	nanocoulombs_per_kilogram := a.convertFromBase(RadiationExposureNanocoulombPerKilogram)
	a.nanocoulombs_per_kilogramLazy = &nanocoulombs_per_kilogram
	return nanocoulombs_per_kilogram
}

// MicrocoulombPerKilogram returns the value in MicrocoulombPerKilogram.
func (a *RadiationExposure) MicrocoulombsPerKilogram() float64 {
	if a.microcoulombs_per_kilogramLazy != nil {
		return *a.microcoulombs_per_kilogramLazy
	}
	microcoulombs_per_kilogram := a.convertFromBase(RadiationExposureMicrocoulombPerKilogram)
	a.microcoulombs_per_kilogramLazy = &microcoulombs_per_kilogram
	return microcoulombs_per_kilogram
}

// MillicoulombPerKilogram returns the value in MillicoulombPerKilogram.
func (a *RadiationExposure) MillicoulombsPerKilogram() float64 {
	if a.millicoulombs_per_kilogramLazy != nil {
		return *a.millicoulombs_per_kilogramLazy
	}
	millicoulombs_per_kilogram := a.convertFromBase(RadiationExposureMillicoulombPerKilogram)
	a.millicoulombs_per_kilogramLazy = &millicoulombs_per_kilogram
	return millicoulombs_per_kilogram
}

// Microroentgen returns the value in Microroentgen.
func (a *RadiationExposure) Microroentgens() float64 {
	if a.microroentgensLazy != nil {
		return *a.microroentgensLazy
	}
	microroentgens := a.convertFromBase(RadiationExposureMicroroentgen)
	a.microroentgensLazy = &microroentgens
	return microroentgens
}

// Milliroentgen returns the value in Milliroentgen.
func (a *RadiationExposure) Milliroentgens() float64 {
	if a.milliroentgensLazy != nil {
		return *a.milliroentgensLazy
	}
	milliroentgens := a.convertFromBase(RadiationExposureMilliroentgen)
	a.milliroentgensLazy = &milliroentgens
	return milliroentgens
}


// ToDto creates an RadiationExposureDto representation.
func (a *RadiationExposure) ToDto(holdInUnit *RadiationExposureUnits) RadiationExposureDto {
	if holdInUnit == nil {
		defaultUnit := RadiationExposureCoulombPerKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return RadiationExposureDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an RadiationExposureDto representation.
func (a *RadiationExposure) ToDtoJSON(holdInUnit *RadiationExposureUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts RadiationExposure to a specific unit value.
func (a *RadiationExposure) Convert(toUnit RadiationExposureUnits) float64 {
	switch toUnit { 
    case RadiationExposureCoulombPerKilogram:
		return a.CoulombsPerKilogram()
    case RadiationExposureRoentgen:
		return a.Roentgens()
    case RadiationExposurePicocoulombPerKilogram:
		return a.PicocoulombsPerKilogram()
    case RadiationExposureNanocoulombPerKilogram:
		return a.NanocoulombsPerKilogram()
    case RadiationExposureMicrocoulombPerKilogram:
		return a.MicrocoulombsPerKilogram()
    case RadiationExposureMillicoulombPerKilogram:
		return a.MillicoulombsPerKilogram()
    case RadiationExposureMicroroentgen:
		return a.Microroentgens()
    case RadiationExposureMilliroentgen:
		return a.Milliroentgens()
	default:
		return 0
	}
}

func (a *RadiationExposure) convertFromBase(toUnit RadiationExposureUnits) float64 {
    value := a.value
	switch toUnit { 
	case RadiationExposureCoulombPerKilogram:
		return (value) 
	case RadiationExposureRoentgen:
		return (value / 2.58e-4) 
	case RadiationExposurePicocoulombPerKilogram:
		return ((value) / 1e-12) 
	case RadiationExposureNanocoulombPerKilogram:
		return ((value) / 1e-09) 
	case RadiationExposureMicrocoulombPerKilogram:
		return ((value) / 1e-06) 
	case RadiationExposureMillicoulombPerKilogram:
		return ((value) / 0.001) 
	case RadiationExposureMicroroentgen:
		return ((value / 2.58e-4) / 1e-06) 
	case RadiationExposureMilliroentgen:
		return ((value / 2.58e-4) / 0.001) 
	default:
		return math.NaN()
	}
}

func (a *RadiationExposure) convertToBase(value float64, fromUnit RadiationExposureUnits) float64 {
	switch fromUnit { 
	case RadiationExposureCoulombPerKilogram:
		return (value) 
	case RadiationExposureRoentgen:
		return (value * 2.58e-4) 
	case RadiationExposurePicocoulombPerKilogram:
		return ((value) * 1e-12) 
	case RadiationExposureNanocoulombPerKilogram:
		return ((value) * 1e-09) 
	case RadiationExposureMicrocoulombPerKilogram:
		return ((value) * 1e-06) 
	case RadiationExposureMillicoulombPerKilogram:
		return ((value) * 0.001) 
	case RadiationExposureMicroroentgen:
		return ((value * 2.58e-4) * 1e-06) 
	case RadiationExposureMilliroentgen:
		return ((value * 2.58e-4) * 0.001) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a RadiationExposure) String() string {
	return a.ToString(RadiationExposureCoulombPerKilogram, 2)
}

// ToString formats the RadiationExposure to string.
// fractionalDigits -1 for not mention
func (a *RadiationExposure) ToString(unit RadiationExposureUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *RadiationExposure) getUnitAbbreviation(unit RadiationExposureUnits) string {
	switch unit { 
	case RadiationExposureCoulombPerKilogram:
		return "C/kg" 
	case RadiationExposureRoentgen:
		return "R" 
	case RadiationExposurePicocoulombPerKilogram:
		return "pC/kg" 
	case RadiationExposureNanocoulombPerKilogram:
		return "nC/kg" 
	case RadiationExposureMicrocoulombPerKilogram:
		return "μC/kg" 
	case RadiationExposureMillicoulombPerKilogram:
		return "mC/kg" 
	case RadiationExposureMicroroentgen:
		return "μR" 
	case RadiationExposureMilliroentgen:
		return "mR" 
	default:
		return ""
	}
}

// Check if the given RadiationExposure are equals to the current RadiationExposure
func (a *RadiationExposure) Equals(other *RadiationExposure) bool {
	return a.value == other.BaseValue()
}

// Check if the given RadiationExposure are equals to the current RadiationExposure
func (a *RadiationExposure) CompareTo(other *RadiationExposure) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given RadiationExposure to the current RadiationExposure.
func (a *RadiationExposure) Add(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value + other.BaseValue()}
}

// Subtract the given RadiationExposure to the current RadiationExposure.
func (a *RadiationExposure) Subtract(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value - other.BaseValue()}
}

// Multiply the given RadiationExposure to the current RadiationExposure.
func (a *RadiationExposure) Multiply(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value * other.BaseValue()}
}

// Divide the given RadiationExposure to the current RadiationExposure.
func (a *RadiationExposure) Divide(other *RadiationExposure) *RadiationExposure {
	return &RadiationExposure{value: a.value / other.BaseValue()}
}