// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// AbsorbedDoseOfIonizingRadiationUnits enumeration
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

// AbsorbedDoseOfIonizingRadiationDto represents an AbsorbedDoseOfIonizingRadiation
type AbsorbedDoseOfIonizingRadiationDto struct {
	Value float64
	Unit  AbsorbedDoseOfIonizingRadiationUnits
}

// AbsorbedDoseOfIonizingRadiationDtoFactory struct to group related functions
type AbsorbedDoseOfIonizingRadiationDtoFactory struct{}

func (udf AbsorbedDoseOfIonizingRadiationDtoFactory) FromJSON(data []byte) (*AbsorbedDoseOfIonizingRadiationDto, error) {
	a := AbsorbedDoseOfIonizingRadiationDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a AbsorbedDoseOfIonizingRadiationDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// AbsorbedDoseOfIonizingRadiation struct
type AbsorbedDoseOfIonizingRadiation struct {
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

// AbsorbedDoseOfIonizingRadiationFactory struct to group related functions
type AbsorbedDoseOfIonizingRadiationFactory struct{}

func (uf AbsorbedDoseOfIonizingRadiationFactory) CreateAbsorbedDoseOfIonizingRadiation(value float64, unit AbsorbedDoseOfIonizingRadiationUnits) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, unit)
}

func (uf AbsorbedDoseOfIonizingRadiationFactory) FromDto(dto AbsorbedDoseOfIonizingRadiationDto) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(dto.Value, dto.Unit)
}

func (uf AbsorbedDoseOfIonizingRadiationFactory) FromDtoJSON(data []byte) (*AbsorbedDoseOfIonizingRadiation, error) {
	unitDto, err := AbsorbedDoseOfIonizingRadiationDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return AbsorbedDoseOfIonizingRadiationFactory{}.FromDto(*unitDto)
}


// FromGray creates a new AbsorbedDoseOfIonizingRadiation instance from Gray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromGrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationGray)
}

// FromRad creates a new AbsorbedDoseOfIonizingRadiation instance from Rad.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromRads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationRad)
}

// FromFemtogray creates a new AbsorbedDoseOfIonizingRadiation instance from Femtogray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromFemtograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationFemtogray)
}

// FromPicogray creates a new AbsorbedDoseOfIonizingRadiation instance from Picogray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromPicograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationPicogray)
}

// FromNanogray creates a new AbsorbedDoseOfIonizingRadiation instance from Nanogray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromNanograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationNanogray)
}

// FromMicrogray creates a new AbsorbedDoseOfIonizingRadiation instance from Microgray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMicrograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMicrogray)
}

// FromMilligray creates a new AbsorbedDoseOfIonizingRadiation instance from Milligray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMilligrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMilligray)
}

// FromCentigray creates a new AbsorbedDoseOfIonizingRadiation instance from Centigray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromCentigrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationCentigray)
}

// FromKilogray creates a new AbsorbedDoseOfIonizingRadiation instance from Kilogray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromKilograys(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationKilogray)
}

// FromMegagray creates a new AbsorbedDoseOfIonizingRadiation instance from Megagray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMegagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMegagray)
}

// FromGigagray creates a new AbsorbedDoseOfIonizingRadiation instance from Gigagray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromGigagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationGigagray)
}

// FromTeragray creates a new AbsorbedDoseOfIonizingRadiation instance from Teragray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromTeragrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationTeragray)
}

// FromPetagray creates a new AbsorbedDoseOfIonizingRadiation instance from Petagray.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromPetagrays(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationPetagray)
}

// FromMillirad creates a new AbsorbedDoseOfIonizingRadiation instance from Millirad.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMillirads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMillirad)
}

// FromKilorad creates a new AbsorbedDoseOfIonizingRadiation instance from Kilorad.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromKilorads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationKilorad)
}

// FromMegarad creates a new AbsorbedDoseOfIonizingRadiation instance from Megarad.
func (uf AbsorbedDoseOfIonizingRadiationFactory) FromMegarads(value float64) (*AbsorbedDoseOfIonizingRadiation, error) {
	return newAbsorbedDoseOfIonizingRadiation(value, AbsorbedDoseOfIonizingRadiationMegarad)
}




// newAbsorbedDoseOfIonizingRadiation creates a new AbsorbedDoseOfIonizingRadiation.
func newAbsorbedDoseOfIonizingRadiation(value float64, fromUnit AbsorbedDoseOfIonizingRadiationUnits) (*AbsorbedDoseOfIonizingRadiation, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &AbsorbedDoseOfIonizingRadiation{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of AbsorbedDoseOfIonizingRadiation in Gray.
func (a *AbsorbedDoseOfIonizingRadiation) BaseValue() float64 {
	return a.value
}


// Gray returns the value in Gray.
func (a *AbsorbedDoseOfIonizingRadiation) Grays() float64 {
	if a.graysLazy != nil {
		return *a.graysLazy
	}
	grays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationGray)
	a.graysLazy = &grays
	return grays
}

// Rad returns the value in Rad.
func (a *AbsorbedDoseOfIonizingRadiation) Rads() float64 {
	if a.radsLazy != nil {
		return *a.radsLazy
	}
	rads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationRad)
	a.radsLazy = &rads
	return rads
}

// Femtogray returns the value in Femtogray.
func (a *AbsorbedDoseOfIonizingRadiation) Femtograys() float64 {
	if a.femtograysLazy != nil {
		return *a.femtograysLazy
	}
	femtograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationFemtogray)
	a.femtograysLazy = &femtograys
	return femtograys
}

// Picogray returns the value in Picogray.
func (a *AbsorbedDoseOfIonizingRadiation) Picograys() float64 {
	if a.picograysLazy != nil {
		return *a.picograysLazy
	}
	picograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationPicogray)
	a.picograysLazy = &picograys
	return picograys
}

// Nanogray returns the value in Nanogray.
func (a *AbsorbedDoseOfIonizingRadiation) Nanograys() float64 {
	if a.nanograysLazy != nil {
		return *a.nanograysLazy
	}
	nanograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationNanogray)
	a.nanograysLazy = &nanograys
	return nanograys
}

// Microgray returns the value in Microgray.
func (a *AbsorbedDoseOfIonizingRadiation) Micrograys() float64 {
	if a.micrograysLazy != nil {
		return *a.micrograysLazy
	}
	micrograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMicrogray)
	a.micrograysLazy = &micrograys
	return micrograys
}

// Milligray returns the value in Milligray.
func (a *AbsorbedDoseOfIonizingRadiation) Milligrays() float64 {
	if a.milligraysLazy != nil {
		return *a.milligraysLazy
	}
	milligrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMilligray)
	a.milligraysLazy = &milligrays
	return milligrays
}

// Centigray returns the value in Centigray.
func (a *AbsorbedDoseOfIonizingRadiation) Centigrays() float64 {
	if a.centigraysLazy != nil {
		return *a.centigraysLazy
	}
	centigrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationCentigray)
	a.centigraysLazy = &centigrays
	return centigrays
}

// Kilogray returns the value in Kilogray.
func (a *AbsorbedDoseOfIonizingRadiation) Kilograys() float64 {
	if a.kilograysLazy != nil {
		return *a.kilograysLazy
	}
	kilograys := a.convertFromBase(AbsorbedDoseOfIonizingRadiationKilogray)
	a.kilograysLazy = &kilograys
	return kilograys
}

// Megagray returns the value in Megagray.
func (a *AbsorbedDoseOfIonizingRadiation) Megagrays() float64 {
	if a.megagraysLazy != nil {
		return *a.megagraysLazy
	}
	megagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMegagray)
	a.megagraysLazy = &megagrays
	return megagrays
}

// Gigagray returns the value in Gigagray.
func (a *AbsorbedDoseOfIonizingRadiation) Gigagrays() float64 {
	if a.gigagraysLazy != nil {
		return *a.gigagraysLazy
	}
	gigagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationGigagray)
	a.gigagraysLazy = &gigagrays
	return gigagrays
}

// Teragray returns the value in Teragray.
func (a *AbsorbedDoseOfIonizingRadiation) Teragrays() float64 {
	if a.teragraysLazy != nil {
		return *a.teragraysLazy
	}
	teragrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationTeragray)
	a.teragraysLazy = &teragrays
	return teragrays
}

// Petagray returns the value in Petagray.
func (a *AbsorbedDoseOfIonizingRadiation) Petagrays() float64 {
	if a.petagraysLazy != nil {
		return *a.petagraysLazy
	}
	petagrays := a.convertFromBase(AbsorbedDoseOfIonizingRadiationPetagray)
	a.petagraysLazy = &petagrays
	return petagrays
}

// Millirad returns the value in Millirad.
func (a *AbsorbedDoseOfIonizingRadiation) Millirads() float64 {
	if a.milliradsLazy != nil {
		return *a.milliradsLazy
	}
	millirads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMillirad)
	a.milliradsLazy = &millirads
	return millirads
}

// Kilorad returns the value in Kilorad.
func (a *AbsorbedDoseOfIonizingRadiation) Kilorads() float64 {
	if a.kiloradsLazy != nil {
		return *a.kiloradsLazy
	}
	kilorads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationKilorad)
	a.kiloradsLazy = &kilorads
	return kilorads
}

// Megarad returns the value in Megarad.
func (a *AbsorbedDoseOfIonizingRadiation) Megarads() float64 {
	if a.megaradsLazy != nil {
		return *a.megaradsLazy
	}
	megarads := a.convertFromBase(AbsorbedDoseOfIonizingRadiationMegarad)
	a.megaradsLazy = &megarads
	return megarads
}


// ToDto creates an AbsorbedDoseOfIonizingRadiationDto representation.
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

// ToDtoJSON creates an AbsorbedDoseOfIonizingRadiationDto representation.
func (a *AbsorbedDoseOfIonizingRadiation) ToDtoJSON(holdInUnit *AbsorbedDoseOfIonizingRadiationUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts AbsorbedDoseOfIonizingRadiation to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a AbsorbedDoseOfIonizingRadiation) String() string {
	return a.ToString(AbsorbedDoseOfIonizingRadiationGray, 2)
}

// ToString formats the AbsorbedDoseOfIonizingRadiation to string.
// fractionalDigits -1 for not mention
func (a *AbsorbedDoseOfIonizingRadiation) ToString(unit AbsorbedDoseOfIonizingRadiationUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *AbsorbedDoseOfIonizingRadiation) getUnitAbbreviation(unit AbsorbedDoseOfIonizingRadiationUnits) string {
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

// Check if the given AbsorbedDoseOfIonizingRadiation are equals to the current AbsorbedDoseOfIonizingRadiation
func (a *AbsorbedDoseOfIonizingRadiation) Equals(other *AbsorbedDoseOfIonizingRadiation) bool {
	return a.value == other.BaseValue()
}

// Check if the given AbsorbedDoseOfIonizingRadiation are equals to the current AbsorbedDoseOfIonizingRadiation
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

// Add the given AbsorbedDoseOfIonizingRadiation to the current AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Add(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value + other.BaseValue()}
}

// Subtract the given AbsorbedDoseOfIonizingRadiation to the current AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Subtract(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value - other.BaseValue()}
}

// Multiply the given AbsorbedDoseOfIonizingRadiation to the current AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Multiply(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value * other.BaseValue()}
}

// Divide the given AbsorbedDoseOfIonizingRadiation to the current AbsorbedDoseOfIonizingRadiation.
func (a *AbsorbedDoseOfIonizingRadiation) Divide(other *AbsorbedDoseOfIonizingRadiation) *AbsorbedDoseOfIonizingRadiation {
	return &AbsorbedDoseOfIonizingRadiation{value: a.value / other.BaseValue()}
}