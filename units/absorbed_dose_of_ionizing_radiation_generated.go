// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// AbsorbedDoseOfIonizingRadiationUnits defines various units of AbsorbedDoseOfIonizingRadiation.
type AbsorbedDoseOfIonizingRadiationUnits string

const (
    
        // The gray is the unit of ionizing radiation dose in the SI, defined as the absorption of one joule of radiation energy per kilogram of matter.
        AbsorbedDoseOfIonizingRadiationGray AbsorbedDoseOfIonizingRadiationUnits = "Gray"
        // The rad is a unit of absorbed radiation dose, defined as 1 rad = 0.01 Gy = 0.01 J/kg.
        AbsorbedDoseOfIonizingRadiationRad AbsorbedDoseOfIonizingRadiationUnits = "Rad"
        // 
        AbsorbedDoseOfIonizingRadiationFemtogray AbsorbedDoseOfIonizingRadiationUnits = "Femtogray"
        // 
        AbsorbedDoseOfIonizingRadiationPicogray AbsorbedDoseOfIonizingRadiationUnits = "Picogray"
        // 
        AbsorbedDoseOfIonizingRadiationNanogray AbsorbedDoseOfIonizingRadiationUnits = "Nanogray"
        // 
        AbsorbedDoseOfIonizingRadiationMicrogray AbsorbedDoseOfIonizingRadiationUnits = "Microgray"
        // 
        AbsorbedDoseOfIonizingRadiationMilligray AbsorbedDoseOfIonizingRadiationUnits = "Milligray"
        // 
        AbsorbedDoseOfIonizingRadiationCentigray AbsorbedDoseOfIonizingRadiationUnits = "Centigray"
        // 
        AbsorbedDoseOfIonizingRadiationKilogray AbsorbedDoseOfIonizingRadiationUnits = "Kilogray"
        // 
        AbsorbedDoseOfIonizingRadiationMegagray AbsorbedDoseOfIonizingRadiationUnits = "Megagray"
        // 
        AbsorbedDoseOfIonizingRadiationGigagray AbsorbedDoseOfIonizingRadiationUnits = "Gigagray"
        // 
        AbsorbedDoseOfIonizingRadiationTeragray AbsorbedDoseOfIonizingRadiationUnits = "Teragray"
        // 
        AbsorbedDoseOfIonizingRadiationPetagray AbsorbedDoseOfIonizingRadiationUnits = "Petagray"
        // 
        AbsorbedDoseOfIonizingRadiationMillirad AbsorbedDoseOfIonizingRadiationUnits = "Millirad"
        // 
        AbsorbedDoseOfIonizingRadiationKilorad AbsorbedDoseOfIonizingRadiationUnits = "Kilorad"
        // 
        AbsorbedDoseOfIonizingRadiationMegarad AbsorbedDoseOfIonizingRadiationUnits = "Megarad"
)

var internalAbsorbedDoseOfIonizingRadiationUnitsMap = map[AbsorbedDoseOfIonizingRadiationUnits]bool{
	
	AbsorbedDoseOfIonizingRadiationGray: true,
	AbsorbedDoseOfIonizingRadiationRad: true,
	AbsorbedDoseOfIonizingRadiationFemtogray: true,
	AbsorbedDoseOfIonizingRadiationPicogray: true,
	AbsorbedDoseOfIonizingRadiationNanogray: true,
	AbsorbedDoseOfIonizingRadiationMicrogray: true,
	AbsorbedDoseOfIonizingRadiationMilligray: true,
	AbsorbedDoseOfIonizingRadiationCentigray: true,
	AbsorbedDoseOfIonizingRadiationKilogray: true,
	AbsorbedDoseOfIonizingRadiationMegagray: true,
	AbsorbedDoseOfIonizingRadiationGigagray: true,
	AbsorbedDoseOfIonizingRadiationTeragray: true,
	AbsorbedDoseOfIonizingRadiationPetagray: true,
	AbsorbedDoseOfIonizingRadiationMillirad: true,
	AbsorbedDoseOfIonizingRadiationKilorad: true,
	AbsorbedDoseOfIonizingRadiationMegarad: true,
}

// AbsorbedDoseOfIonizingRadiationDto represents a AbsorbedDoseOfIonizingRadiation measurement with a numerical value and its corresponding unit.
type AbsorbedDoseOfIonizingRadiationDto struct {
    // Value is the numerical representation of the AbsorbedDoseOfIonizingRadiation.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the AbsorbedDoseOfIonizingRadiation, as defined in the AbsorbedDoseOfIonizingRadiationUnits enumeration.
	Unit  AbsorbedDoseOfIonizingRadiationUnits `json:"unit" validate:"required,oneof=Gray Rad Femtogray Picogray Nanogray Microgray Milligray Centigray Kilogray Megagray Gigagray Teragray Petagray Millirad Kilorad Megarad"`
}

// AbsorbedDoseOfIonizingRadiationDtoFactory groups methods for creating and serializing AbsorbedDoseOfIonizingRadiationDto objects.
type AbsorbedDoseOfIonizingRadiationDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a AbsorbedDoseOfIonizingRadiationDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf AbsorbedDoseOfIonizingRadiationDtoFactory) FromJSON(data []byte) (*AbsorbedDoseOfIonizingRadiationDto, error) {
	a := AbsorbedDoseOfIonizingRadiationDto{}

    // Parse JSON into AbsorbedDoseOfIonizingRadiationDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a AbsorbedDoseOfIonizingRadiationDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a AbsorbedDoseOfIonizingRadiationDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// AbsorbedDoseOfIonizingRadiation represents a measurement in a of AbsorbedDoseOfIonizingRadiation.
//
// Absorbed dose is a dose quantity which is the measure of the energy deposited in matter by ionizing radiation per unit mass.
type AbsorbedDoseOfIonizingRadiation struct {
	// value is the base measurement stored internally.
	value       float64
    
    graysLazy *float64 
    radsLazy *float64 
    femtograysLazy *float64 
    picograysLazy *float64 
    nanograysLazy *float64 
    micrograysLazy *float64 
    milligraysLazy *float64 
    centigraysLazy *float64 
    kilograysLazy *float64 
    megagraysLazy *float64 
    gigagraysLazy *float64 
    teragraysLazy *float64 
    petagraysLazy *float64 
    milliradsLazy *float64 
    kiloradsLazy *float64 
    megaradsLazy *float64 
}

// AbsorbedDoseOfIonizingRadiationFactory groups methods for creating AbsorbedDoseOfIonizingRadiation instances.
type AbsorbedDoseOfIonizingRadiationFactory struct{}

// CreateAbsorbedDoseOfIonizingRadiation creates a new AbsorbedDoseOfIonizingRadiation instance from the given value and unit.
func (uf AbsorbedDoseOfIonizingRadiationFactory) CreateAbsorbedDoseOfIonizingRadiation(value float64, unit AbsorbedDoseOfIonizingRadiationUnits) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, unit)
}

// FromDto converts a AbsorbedDoseOfIonizingRadiationDto to a AbsorbedDoseOfIonizingRadiation instance.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromDto(dto AbsorbedDoseOfIonizingRadiationDto) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a AbsorbedDoseOfIonizingRadiation instance.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromDtoJSON(data []byte) (*AbsorbedDoseOfIonizingRadiation, error) {
	unitDto, err := AbsorbedDoseOfIonizingRadiationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AbsorbedDoseOfIonizingRadiationDto from JSON: %w", err)
	}
	return AbsorbedDoseOfIonizingRadiationFactory{}.FromDto(*unitDto)
}


// FromGrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Grays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromGrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationGray)
}

// FromRads creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Rads.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromRads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationRad)
}

// FromFemtograys creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Femtograys.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromFemtograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationFemtogray)
}

// FromPicograys creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Picograys.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromPicograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationPicogray)
}

// FromNanograys creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Nanograys.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromNanograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationNanogray)
}

// FromMicrograys creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Micrograys.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMicrograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMicrogray)
}

// FromMilligrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Milligrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMilligrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMilligray)
}

// FromCentigrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Centigrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromCentigrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationCentigray)
}

// FromKilograys creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Kilograys.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromKilograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationKilogray)
}

// FromMegagrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Megagrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMegagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMegagray)
}

// FromGigagrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Gigagrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromGigagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationGigagray)
}

// FromTeragrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Teragrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromTeragrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationTeragray)
}

// FromPetagrays creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Petagrays.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromPetagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationPetagray)
}

// FromMillirads creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Millirads.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMillirads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMillirad)
}

// FromKilorads creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Kilorads.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromKilorads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationKilorad)
}

// FromMegarads creates a new AbsorbedDoseOfIonizingRadiation instance from a value in Megarads.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMegarads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMegarad)
}


// newAbsorbedDoseOfIonizingRadiation creates a new AbsorbedDoseOfIonizingRadiation.
func newAbsorbedDoseOfIonizingRadiation(value float64, fromUnit AbsorbedDoseOfIonizingRadiationUnits) (*AbsorbedDoseOfIonizingRadiation, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalAbsorbedDoseOfIonizingRadiationUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in AbsorbedDoseOfIonizingRadiationUnits", fromUnit)
	}
	a := &AbsorbedDoseOfIonizingRadiation{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AbsorbedDoseOfIonizingRadiation in Gray unit (the base/default quantity).
func (a *AbsorbedDoseOfIonizingRadiation) BaseValue() float64 {
	return a.value
}


// Grays returns the AbsorbedDoseOfIonizingRadiation value in Grays.
//
// The gray is the unit of ionizing radiation dose in the SI, defined as the absorption of one joule of radiation energy per kilogram of matter.
func (a *AbsorbedDoseOfIonizingRadiation) Grays() float64 {
	if a.graysLazy != nil {
		return *a.graysLazy
	}
	grays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationGray)
	a.graysLazy = &grays
	return grays
}

// Rads returns the AbsorbedDoseOfIonizingRadiation value in Rads.
//
// The rad is a unit of absorbed radiation dose, defined as 1 rad = 0.01 Gy = 0.01 J/kg.
func (a *AbsorbedDoseOfIonizingRadiation) Rads() float64 {
	if a.radsLazy != nil {
		return *a.radsLazy
	}
	rads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationRad)
	a.radsLazy = &rads
	return rads
}

// Femtograys returns the AbsorbedDoseOfIonizingRadiation value in Femtograys.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Femtograys() float64 {
	if a.femtograysLazy != nil {
		return *a.femtograysLazy
	}
	femtograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationFemtogray)
	a.femtograysLazy = &femtograys
	return femtograys
}

// Picograys returns the AbsorbedDoseOfIonizingRadiation value in Picograys.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Picograys() float64 {
	if a.picograysLazy != nil {
		return *a.picograysLazy
	}
	picograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationPicogray)
	a.picograysLazy = &picograys
	return picograys
}

// Nanograys returns the AbsorbedDoseOfIonizingRadiation value in Nanograys.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Nanograys() float64 {
	if a.nanograysLazy != nil {
		return *a.nanograysLazy
	}
	nanograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationNanogray)
	a.nanograysLazy = &nanograys
	return nanograys
}

// Micrograys returns the AbsorbedDoseOfIonizingRadiation value in Micrograys.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Micrograys() float64 {
	if a.micrograysLazy != nil {
		return *a.micrograysLazy
	}
	micrograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMicrogray)
	a.micrograysLazy = &micrograys
	return micrograys
}

// Milligrays returns the AbsorbedDoseOfIonizingRadiation value in Milligrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Milligrays() float64 {
	if a.milligraysLazy != nil {
		return *a.milligraysLazy
	}
	milligrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMilligray)
	a.milligraysLazy = &milligrays
	return milligrays
}

// Centigrays returns the AbsorbedDoseOfIonizingRadiation value in Centigrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Centigrays() float64 {
	if a.centigraysLazy != nil {
		return *a.centigraysLazy
	}
	centigrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationCentigray)
	a.centigraysLazy = &centigrays
	return centigrays
}

// Kilograys returns the AbsorbedDoseOfIonizingRadiation value in Kilograys.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Kilograys() float64 {
	if a.kilograysLazy != nil {
		return *a.kilograysLazy
	}
	kilograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationKilogray)
	a.kilograysLazy = &kilograys
	return kilograys
}

// Megagrays returns the AbsorbedDoseOfIonizingRadiation value in Megagrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Megagrays() float64 {
	if a.megagraysLazy != nil {
		return *a.megagraysLazy
	}
	megagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMegagray)
	a.megagraysLazy = &megagrays
	return megagrays
}

// Gigagrays returns the AbsorbedDoseOfIonizingRadiation value in Gigagrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Gigagrays() float64 {
	if a.gigagraysLazy != nil {
		return *a.gigagraysLazy
	}
	gigagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationGigagray)
	a.gigagraysLazy = &gigagrays
	return gigagrays
}

// Teragrays returns the AbsorbedDoseOfIonizingRadiation value in Teragrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Teragrays() float64 {
	if a.teragraysLazy != nil {
		return *a.teragraysLazy
	}
	teragrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationTeragray)
	a.teragraysLazy = &teragrays
	return teragrays
}

// Petagrays returns the AbsorbedDoseOfIonizingRadiation value in Petagrays.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Petagrays() float64 {
	if a.petagraysLazy != nil {
		return *a.petagraysLazy
	}
	petagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationPetagray)
	a.petagraysLazy = &petagrays
	return petagrays
}

// Millirads returns the AbsorbedDoseOfIonizingRadiation value in Millirads.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Millirads() float64 {
	if a.milliradsLazy != nil {
		return *a.milliradsLazy
	}
	millirads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMillirad)
	a.milliradsLazy = &millirads
	return millirads
}

// Kilorads returns the AbsorbedDoseOfIonizingRadiation value in Kilorads.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Kilorads() float64 {
	if a.kiloradsLazy != nil {
		return *a.kiloradsLazy
	}
	kilorads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationKilorad)
	a.kiloradsLazy = &kilorads
	return kilorads
}

// Megarads returns the AbsorbedDoseOfIonizingRadiation value in Megarads.
//
// 
func (a *AbsorbedDoseOfIonizingRadiation) Megarads() float64 {
	if a.megaradsLazy != nil {
		return *a.megaradsLazy
	}
	megarads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMegarad)
	a.megaradsLazy = &megarads
	return megarads
}


// ToDto creates a AbsorbedDoseOfIonizingRadiationDto representation from the AbsorbedDoseOfIonizingRadiation instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Gray by default.
func (a *AbsorbedDoseOfIonizingRadiation) ToDto(holdInUnit *AbsorbedDoseOfIonizingRadiationUnits) AbsorbedDoseOfIonizingRadiationDto {
	if holdInUnit == nil {
		defaultUnit := AbsorbedDoseOfIonizingRadiationGray // Default value
		holdInUnit = &defaultUnit
	}

	return AbsorbedDoseOfIonizingRadiationDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the AbsorbedDoseOfIonizingRadiation instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Gray by default.
func (a *AbsorbedDoseOfIonizingRadiation) ToDtoJSON(holdInUnit *AbsorbedDoseOfIonizingRadiationUnits) ([]byte, error) {
	// Convert to AbsorbedDoseOfIonizingRadiationDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a AbsorbedDoseOfIonizingRadiation to a specific unit value.
// The function uses the provided unit type (AbsorbedDoseOfIonizingRadiationUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *AbsorbedDoseOfIonizingRadiation) Convert(toUnit AbsorbedDoseOfIonizingRadiationUnits) float64 {
	switch toUnit { 
    case AbsorbedDoseOfIonizingRadiationGray:
		return a.Grays()
    case AbsorbedDoseOfIonizingRadiationRad:
		return a.Rads()
    case AbsorbedDoseOfIonizingRadiationFemtogray:
		return a.Femtograys()
    case AbsorbedDoseOfIonizingRadiationPicogray:
		return a.Picograys()
    case AbsorbedDoseOfIonizingRadiationNanogray:
		return a.Nanograys()
    case AbsorbedDoseOfIonizingRadiationMicrogray:
		return a.Micrograys()
    case AbsorbedDoseOfIonizingRadiationMilligray:
		return a.Milligrays()
    case AbsorbedDoseOfIonizingRadiationCentigray:
		return a.Centigrays()
    case AbsorbedDoseOfIonizingRadiationKilogray:
		return a.Kilograys()
    case AbsorbedDoseOfIonizingRadiationMegagray:
		return a.Megagrays()
    case AbsorbedDoseOfIonizingRadiationGigagray:
		return a.Gigagrays()
    case AbsorbedDoseOfIonizingRadiationTeragray:
		return a.Teragrays()
    case AbsorbedDoseOfIonizingRadiationPetagray:
		return a.Petagrays()
    case AbsorbedDoseOfIonizingRadiationMillirad:
		return a.Millirads()
    case AbsorbedDoseOfIonizingRadiationKilorad:
		return a.Kilorads()
    case AbsorbedDoseOfIonizingRadiationMegarad:
		return a.Megarads()
	default:
		return math.NaN()
	}
}

func (a *AbsorbedDoseOfIonizingRadiation) convertFromBase(toUnit AbsorbedDoseOfIonizingRadiationUnits) float64 {
    value := a.value
	switch toUnit { 
	case AbsorbedDoseOfIonizingRadiationGray:
		return (value) 
	case AbsorbedDoseOfIonizingRadiationRad:
		return (value * 100) 
	case AbsorbedDoseOfIonizingRadiationFemtogray:
		return ((value) / 1e-15) 
	case AbsorbedDoseOfIonizingRadiationPicogray:
		return ((value) / 1e-12) 
	case AbsorbedDoseOfIonizingRadiationNanogray:
		return ((value) / 1e-09) 
	case AbsorbedDoseOfIonizingRadiationMicrogray:
		return ((value) / 1e-06) 
	case AbsorbedDoseOfIonizingRadiationMilligray:
		return ((value) / 0.001) 
	case AbsorbedDoseOfIonizingRadiationCentigray:
		return ((value) / 0.01) 
	case AbsorbedDoseOfIonizingRadiationKilogray:
		return ((value) / 1000.0) 
	case AbsorbedDoseOfIonizingRadiationMegagray:
		return ((value) / 1000000.0) 
	case AbsorbedDoseOfIonizingRadiationGigagray:
		return ((value) / 1000000000.0) 
	case AbsorbedDoseOfIonizingRadiationTeragray:
		return ((value) / 1000000000000.0) 
	case AbsorbedDoseOfIonizingRadiationPetagray:
		return ((value) / 1000000000000000.0) 
	case AbsorbedDoseOfIonizingRadiationMillirad:
		return ((value * 100) / 0.001) 
	case AbsorbedDoseOfIonizingRadiationKilorad:
		return ((value * 100) / 1000.0) 
	case AbsorbedDoseOfIonizingRadiationMegarad:
		return ((value * 100) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *AbsorbedDoseOfIonizingRadiation) convertToBase(value float64, fromUnit AbsorbedDoseOfIonizingRadiationUnits) float64 {
	switch fromUnit { 
	case AbsorbedDoseOfIonizingRadiationGray:
		return (value) 
	case AbsorbedDoseOfIonizingRadiationRad:
		return (value / 100) 
	case AbsorbedDoseOfIonizingRadiationFemtogray:
		return ((value) * 1e-15) 
	case AbsorbedDoseOfIonizingRadiationPicogray:
		return ((value) * 1e-12) 
	case AbsorbedDoseOfIonizingRadiationNanogray:
		return ((value) * 1e-09) 
	case AbsorbedDoseOfIonizingRadiationMicrogray:
		return ((value) * 1e-06) 
	case AbsorbedDoseOfIonizingRadiationMilligray:
		return ((value) * 0.001) 
	case AbsorbedDoseOfIonizingRadiationCentigray:
		return ((value) * 0.01) 
	case AbsorbedDoseOfIonizingRadiationKilogray:
		return ((value) * 1000.0) 
	case AbsorbedDoseOfIonizingRadiationMegagray:
		return ((value) * 1000000.0) 
	case AbsorbedDoseOfIonizingRadiationGigagray:
		return ((value) * 1000000000.0) 
	case AbsorbedDoseOfIonizingRadiationTeragray:
		return ((value) * 1000000000000.0) 
	case AbsorbedDoseOfIonizingRadiationPetagray:
		return ((value) * 1000000000000000.0) 
	case AbsorbedDoseOfIonizingRadiationMillirad:
		return ((value / 100) * 0.001) 
	case AbsorbedDoseOfIonizingRadiationKilorad:
		return ((value / 100) * 1000.0) 
	case AbsorbedDoseOfIonizingRadiationMegarad:
		return ((value / 100) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the AbsorbedDoseOfIonizingRadiation in the default unit (Gray),
// formatted to two decimal places.
func (a AbsorbedDoseOfIonizingRadiation) String() string {
	return a.ToString(AbsorbedDoseOfIonizingRadiationGray, 2)
}

// ToString formats the AbsorbedDoseOfIonizingRadiation value as a string with the specified unit and fractional digits.
// It converts the AbsorbedDoseOfIonizingRadiation to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the AbsorbedDoseOfIonizingRadiation value will be converted (e.g., Gray))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the AbsorbedDoseOfIonizingRadiation with the unit abbreviation.
func (a *AbsorbedDoseOfIonizingRadiation) ToString(unit AbsorbedDoseOfIonizingRadiationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetAbsorbedDoseOfIonizingRadiationAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetAbsorbedDoseOfIonizingRadiationAbbreviation(unit))
}

// Equals checks if the given AbsorbedDoseOfIonizingRadiation is equal to the current AbsorbedDoseOfIonizingRadiation.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to compare against.
//
// Returns:
//    bool: Returns true if both AbsorbedDoseOfIonizingRadiation are equal, false otherwise.
func (a *AbsorbedDoseOfIonizingRadiation) Equals(other *AbsorbedDoseOfIonizingRadiation) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current AbsorbedDoseOfIonizingRadiation with another AbsorbedDoseOfIonizingRadiation.
// It returns -1 if the current AbsorbedDoseOfIonizingRadiation is less than the other AbsorbedDoseOfIonizingRadiation, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to compare against.
//
// Returns:
//    int: -1 if the current AbsorbedDoseOfIonizingRadiation is less, 1 if greater, and 0 if equal.
func (a *AbsorbedDoseOfIonizingRadiation) CompareTo(other *AbsorbedDoseOfIonizingRadiation) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given AbsorbedDoseOfIonizingRadiation to the current AbsorbedDoseOfIonizingRadiation and returns the result.
// The result is a new AbsorbedDoseOfIonizingRadiation instance with the sum of the values.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to add to the current AbsorbedDoseOfIonizingRadiation.
//
// Returns:
//    *AbsorbedDoseOfIonizingRadiation: A new AbsorbedDoseOfIonizingRadiation instance representing the sum of both AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Add(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given AbsorbedDoseOfIonizingRadiation from the current AbsorbedDoseOfIonizingRadiation and returns the result.
// The result is a new AbsorbedDoseOfIonizingRadiation instance with the difference of the values.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to subtract from the current AbsorbedDoseOfIonizingRadiation.
//
// Returns:
//    *AbsorbedDoseOfIonizingRadiation: A new AbsorbedDoseOfIonizingRadiation instance representing the difference of both AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Subtract(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current AbsorbedDoseOfIonizingRadiation by the given AbsorbedDoseOfIonizingRadiation and returns the result.
// The result is a new AbsorbedDoseOfIonizingRadiation instance with the product of the values.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to multiply with the current AbsorbedDoseOfIonizingRadiation.
//
// Returns:
//    *AbsorbedDoseOfIonizingRadiation: A new AbsorbedDoseOfIonizingRadiation instance representing the product of both AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Multiply(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value * other.BaseValue()}
}

// Divide divides the current AbsorbedDoseOfIonizingRadiation by the given AbsorbedDoseOfIonizingRadiation and returns the result.
// The result is a new AbsorbedDoseOfIonizingRadiation instance with the quotient of the values.
//
// Parameters:
//    other: The AbsorbedDoseOfIonizingRadiation to divide the current AbsorbedDoseOfIonizingRadiation by.
//
// Returns:
//    *AbsorbedDoseOfIonizingRadiation: A new AbsorbedDoseOfIonizingRadiation instance representing the quotient of both AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Divide(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value / other.BaseValue()}
}

// GetAbsorbedDoseOfIonizingRadiationAbbreviation gets the unit abbreviation.
func GetAbsorbedDoseOfIonizingRadiationAbbreviation(unit AbsorbedDoseOfIonizingRadiationUnits) string {
	switch unit { 
	case AbsorbedDoseOfIonizingRadiationGray:
		return "Gy" 
	case AbsorbedDoseOfIonizingRadiationRad:
		return "rad" 
	case AbsorbedDoseOfIonizingRadiationFemtogray:
		return "fGy" 
	case AbsorbedDoseOfIonizingRadiationPicogray:
		return "pGy" 
	case AbsorbedDoseOfIonizingRadiationNanogray:
		return "nGy" 
	case AbsorbedDoseOfIonizingRadiationMicrogray:
		return "μGy" 
	case AbsorbedDoseOfIonizingRadiationMilligray:
		return "mGy" 
	case AbsorbedDoseOfIonizingRadiationCentigray:
		return "cGy" 
	case AbsorbedDoseOfIonizingRadiationKilogray:
		return "kGy" 
	case AbsorbedDoseOfIonizingRadiationMegagray:
		return "MGy" 
	case AbsorbedDoseOfIonizingRadiationGigagray:
		return "GGy" 
	case AbsorbedDoseOfIonizingRadiationTeragray:
		return "TGy" 
	case AbsorbedDoseOfIonizingRadiationPetagray:
		return "PGy" 
	case AbsorbedDoseOfIonizingRadiationMillirad:
		return "mrad" 
	case AbsorbedDoseOfIonizingRadiationKilorad:
		return "krad" 
	case AbsorbedDoseOfIonizingRadiationMegarad:
		return "Mrad" 
	default:
		return ""
	}
}