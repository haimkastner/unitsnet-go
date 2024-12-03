// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LuminosityUnits enumeration
type LuminosityUnits string

const (
    
        // 
        LuminosityWatt LuminosityUnits = "Watt"
        // 
        LuminositySolarLuminosity LuminosityUnits = "SolarLuminosity"
        // 
        LuminosityFemtowatt LuminosityUnits = "Femtowatt"
        // 
        LuminosityPicowatt LuminosityUnits = "Picowatt"
        // 
        LuminosityNanowatt LuminosityUnits = "Nanowatt"
        // 
        LuminosityMicrowatt LuminosityUnits = "Microwatt"
        // 
        LuminosityMilliwatt LuminosityUnits = "Milliwatt"
        // 
        LuminosityDeciwatt LuminosityUnits = "Deciwatt"
        // 
        LuminosityDecawatt LuminosityUnits = "Decawatt"
        // 
        LuminosityKilowatt LuminosityUnits = "Kilowatt"
        // 
        LuminosityMegawatt LuminosityUnits = "Megawatt"
        // 
        LuminosityGigawatt LuminosityUnits = "Gigawatt"
        // 
        LuminosityTerawatt LuminosityUnits = "Terawatt"
        // 
        LuminosityPetawatt LuminosityUnits = "Petawatt"
)

// LuminosityDto represents an Luminosity
type LuminosityDto struct {
	Value float64
	Unit  LuminosityUnits
}

// LuminosityDtoFactory struct to group related functions
type LuminosityDtoFactory struct{}

func (udf LuminosityDtoFactory) FromJSON(data []byte) (*LuminosityDto, error) {
	a := LuminosityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LuminosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Luminosity struct
type Luminosity struct {
	value       float64
    
    wattsLazy *float64 
    solar_luminositiesLazy *float64 
    femtowattsLazy *float64 
    picowattsLazy *float64 
    nanowattsLazy *float64 
    microwattsLazy *float64 
    milliwattsLazy *float64 
    deciwattsLazy *float64 
    decawattsLazy *float64 
    kilowattsLazy *float64 
    megawattsLazy *float64 
    gigawattsLazy *float64 
    terawattsLazy *float64 
    petawattsLazy *float64 
}

// LuminosityFactory struct to group related functions
type LuminosityFactory struct{}

func (uf LuminosityFactory) CreateLuminosity(value float64, unit LuminosityUnits) (*Luminosity, error) {
	return newLuminosity(value, unit)
}

func (uf LuminosityFactory) FromDto(dto LuminosityDto) (*Luminosity, error) {
	return newLuminosity(dto.Value, dto.Unit)
}

func (uf LuminosityFactory) FromDtoJSON(data []byte) (*Luminosity, error) {
	unitDto, err := LuminosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LuminosityFactory{}.FromDto(*unitDto)
}


// FromWatt creates a new Luminosity instance from Watt.
func (uf LuminosityFactory) FromWatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityWatt)
}

// FromSolarLuminosity creates a new Luminosity instance from SolarLuminosity.
func (uf LuminosityFactory) FromSolarLuminosities(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminositySolarLuminosity)
}

// FromFemtowatt creates a new Luminosity instance from Femtowatt.
func (uf LuminosityFactory) FromFemtowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityFemtowatt)
}

// FromPicowatt creates a new Luminosity instance from Picowatt.
func (uf LuminosityFactory) FromPicowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityPicowatt)
}

// FromNanowatt creates a new Luminosity instance from Nanowatt.
func (uf LuminosityFactory) FromNanowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityNanowatt)
}

// FromMicrowatt creates a new Luminosity instance from Microwatt.
func (uf LuminosityFactory) FromMicrowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMicrowatt)
}

// FromMilliwatt creates a new Luminosity instance from Milliwatt.
func (uf LuminosityFactory) FromMilliwatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMilliwatt)
}

// FromDeciwatt creates a new Luminosity instance from Deciwatt.
func (uf LuminosityFactory) FromDeciwatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityDeciwatt)
}

// FromDecawatt creates a new Luminosity instance from Decawatt.
func (uf LuminosityFactory) FromDecawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityDecawatt)
}

// FromKilowatt creates a new Luminosity instance from Kilowatt.
func (uf LuminosityFactory) FromKilowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityKilowatt)
}

// FromMegawatt creates a new Luminosity instance from Megawatt.
func (uf LuminosityFactory) FromMegawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMegawatt)
}

// FromGigawatt creates a new Luminosity instance from Gigawatt.
func (uf LuminosityFactory) FromGigawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityGigawatt)
}

// FromTerawatt creates a new Luminosity instance from Terawatt.
func (uf LuminosityFactory) FromTerawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityTerawatt)
}

// FromPetawatt creates a new Luminosity instance from Petawatt.
func (uf LuminosityFactory) FromPetawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityPetawatt)
}




// newLuminosity creates a new Luminosity.
func newLuminosity(value float64, fromUnit LuminosityUnits) (*Luminosity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Luminosity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Luminosity in Watt.
func (a *Luminosity) BaseValue() float64 {
	return a.value
}


// Watt returns the value in Watt.
func (a *Luminosity) Watts() float64 {
	if a.wattsLazy != nil {
		return *a.wattsLazy
	}
	watts := a.convertFromBase(LuminosityWatt)
	a.wattsLazy = &watts
	return watts
}

// SolarLuminosity returns the value in SolarLuminosity.
func (a *Luminosity) SolarLuminosities() float64 {
	if a.solar_luminositiesLazy != nil {
		return *a.solar_luminositiesLazy
	}
	solar_luminosities := a.convertFromBase(LuminositySolarLuminosity)
	a.solar_luminositiesLazy = &solar_luminosities
	return solar_luminosities
}

// Femtowatt returns the value in Femtowatt.
func (a *Luminosity) Femtowatts() float64 {
	if a.femtowattsLazy != nil {
		return *a.femtowattsLazy
	}
	femtowatts := a.convertFromBase(LuminosityFemtowatt)
	a.femtowattsLazy = &femtowatts
	return femtowatts
}

// Picowatt returns the value in Picowatt.
func (a *Luminosity) Picowatts() float64 {
	if a.picowattsLazy != nil {
		return *a.picowattsLazy
	}
	picowatts := a.convertFromBase(LuminosityPicowatt)
	a.picowattsLazy = &picowatts
	return picowatts
}

// Nanowatt returns the value in Nanowatt.
func (a *Luminosity) Nanowatts() float64 {
	if a.nanowattsLazy != nil {
		return *a.nanowattsLazy
	}
	nanowatts := a.convertFromBase(LuminosityNanowatt)
	a.nanowattsLazy = &nanowatts
	return nanowatts
}

// Microwatt returns the value in Microwatt.
func (a *Luminosity) Microwatts() float64 {
	if a.microwattsLazy != nil {
		return *a.microwattsLazy
	}
	microwatts := a.convertFromBase(LuminosityMicrowatt)
	a.microwattsLazy = &microwatts
	return microwatts
}

// Milliwatt returns the value in Milliwatt.
func (a *Luminosity) Milliwatts() float64 {
	if a.milliwattsLazy != nil {
		return *a.milliwattsLazy
	}
	milliwatts := a.convertFromBase(LuminosityMilliwatt)
	a.milliwattsLazy = &milliwatts
	return milliwatts
}

// Deciwatt returns the value in Deciwatt.
func (a *Luminosity) Deciwatts() float64 {
	if a.deciwattsLazy != nil {
		return *a.deciwattsLazy
	}
	deciwatts := a.convertFromBase(LuminosityDeciwatt)
	a.deciwattsLazy = &deciwatts
	return deciwatts
}

// Decawatt returns the value in Decawatt.
func (a *Luminosity) Decawatts() float64 {
	if a.decawattsLazy != nil {
		return *a.decawattsLazy
	}
	decawatts := a.convertFromBase(LuminosityDecawatt)
	a.decawattsLazy = &decawatts
	return decawatts
}

// Kilowatt returns the value in Kilowatt.
func (a *Luminosity) Kilowatts() float64 {
	if a.kilowattsLazy != nil {
		return *a.kilowattsLazy
	}
	kilowatts := a.convertFromBase(LuminosityKilowatt)
	a.kilowattsLazy = &kilowatts
	return kilowatts
}

// Megawatt returns the value in Megawatt.
func (a *Luminosity) Megawatts() float64 {
	if a.megawattsLazy != nil {
		return *a.megawattsLazy
	}
	megawatts := a.convertFromBase(LuminosityMegawatt)
	a.megawattsLazy = &megawatts
	return megawatts
}

// Gigawatt returns the value in Gigawatt.
func (a *Luminosity) Gigawatts() float64 {
	if a.gigawattsLazy != nil {
		return *a.gigawattsLazy
	}
	gigawatts := a.convertFromBase(LuminosityGigawatt)
	a.gigawattsLazy = &gigawatts
	return gigawatts
}

// Terawatt returns the value in Terawatt.
func (a *Luminosity) Terawatts() float64 {
	if a.terawattsLazy != nil {
		return *a.terawattsLazy
	}
	terawatts := a.convertFromBase(LuminosityTerawatt)
	a.terawattsLazy = &terawatts
	return terawatts
}

// Petawatt returns the value in Petawatt.
func (a *Luminosity) Petawatts() float64 {
	if a.petawattsLazy != nil {
		return *a.petawattsLazy
	}
	petawatts := a.convertFromBase(LuminosityPetawatt)
	a.petawattsLazy = &petawatts
	return petawatts
}


// ToDto creates an LuminosityDto representation.
func (a *Luminosity) ToDto(holdInUnit *LuminosityUnits) LuminosityDto {
	if holdInUnit == nil {
		defaultUnit := LuminosityWatt // Default value
		holdInUnit = &defaultUnit
	}

	return LuminosityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an LuminosityDto representation.
func (a *Luminosity) ToDtoJSON(holdInUnit *LuminosityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Luminosity to a specific unit value.
func (a *Luminosity) Convert(toUnit LuminosityUnits) float64 {
	switch toUnit { 
    case LuminosityWatt:
		return a.Watts()
    case LuminositySolarLuminosity:
		return a.SolarLuminosities()
    case LuminosityFemtowatt:
		return a.Femtowatts()
    case LuminosityPicowatt:
		return a.Picowatts()
    case LuminosityNanowatt:
		return a.Nanowatts()
    case LuminosityMicrowatt:
		return a.Microwatts()
    case LuminosityMilliwatt:
		return a.Milliwatts()
    case LuminosityDeciwatt:
		return a.Deciwatts()
    case LuminosityDecawatt:
		return a.Decawatts()
    case LuminosityKilowatt:
		return a.Kilowatts()
    case LuminosityMegawatt:
		return a.Megawatts()
    case LuminosityGigawatt:
		return a.Gigawatts()
    case LuminosityTerawatt:
		return a.Terawatts()
    case LuminosityPetawatt:
		return a.Petawatts()
	default:
		return 0
	}
}

func (a *Luminosity) convertFromBase(toUnit LuminosityUnits) float64 {
    value := a.value
	switch toUnit { 
	case LuminosityWatt:
		return (value) 
	case LuminositySolarLuminosity:
		return (value / 3.846e26) 
	case LuminosityFemtowatt:
		return ((value) / 1e-15) 
	case LuminosityPicowatt:
		return ((value) / 1e-12) 
	case LuminosityNanowatt:
		return ((value) / 1e-09) 
	case LuminosityMicrowatt:
		return ((value) / 1e-06) 
	case LuminosityMilliwatt:
		return ((value) / 0.001) 
	case LuminosityDeciwatt:
		return ((value) / 0.1) 
	case LuminosityDecawatt:
		return ((value) / 10.0) 
	case LuminosityKilowatt:
		return ((value) / 1000.0) 
	case LuminosityMegawatt:
		return ((value) / 1000000.0) 
	case LuminosityGigawatt:
		return ((value) / 1000000000.0) 
	case LuminosityTerawatt:
		return ((value) / 1000000000000.0) 
	case LuminosityPetawatt:
		return ((value) / 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Luminosity) convertToBase(value float64, fromUnit LuminosityUnits) float64 {
	switch fromUnit { 
	case LuminosityWatt:
		return (value) 
	case LuminositySolarLuminosity:
		return (value * 3.846e26) 
	case LuminosityFemtowatt:
		return ((value) * 1e-15) 
	case LuminosityPicowatt:
		return ((value) * 1e-12) 
	case LuminosityNanowatt:
		return ((value) * 1e-09) 
	case LuminosityMicrowatt:
		return ((value) * 1e-06) 
	case LuminosityMilliwatt:
		return ((value) * 0.001) 
	case LuminosityDeciwatt:
		return ((value) * 0.1) 
	case LuminosityDecawatt:
		return ((value) * 10.0) 
	case LuminosityKilowatt:
		return ((value) * 1000.0) 
	case LuminosityMegawatt:
		return ((value) * 1000000.0) 
	case LuminosityGigawatt:
		return ((value) * 1000000000.0) 
	case LuminosityTerawatt:
		return ((value) * 1000000000000.0) 
	case LuminosityPetawatt:
		return ((value) * 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a Luminosity) String() string {
	return a.ToString(LuminosityWatt, 2)
}

// ToString formats the Luminosity to string.
// fractionalDigits -1 for not mention
func (a *Luminosity) ToString(unit LuminosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Luminosity) getUnitAbbreviation(unit LuminosityUnits) string {
	switch unit { 
	case LuminosityWatt:
		return "W" 
	case LuminositySolarLuminosity:
		return "L⊙" 
	case LuminosityFemtowatt:
		return "fW" 
	case LuminosityPicowatt:
		return "pW" 
	case LuminosityNanowatt:
		return "nW" 
	case LuminosityMicrowatt:
		return "μW" 
	case LuminosityMilliwatt:
		return "mW" 
	case LuminosityDeciwatt:
		return "dW" 
	case LuminosityDecawatt:
		return "daW" 
	case LuminosityKilowatt:
		return "kW" 
	case LuminosityMegawatt:
		return "MW" 
	case LuminosityGigawatt:
		return "GW" 
	case LuminosityTerawatt:
		return "TW" 
	case LuminosityPetawatt:
		return "PW" 
	default:
		return ""
	}
}

// Check if the given Luminosity are equals to the current Luminosity
func (a *Luminosity) Equals(other *Luminosity) bool {
	return a.value == other.BaseValue()
}

// Check if the given Luminosity are equals to the current Luminosity
func (a *Luminosity) CompareTo(other *Luminosity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given Luminosity to the current Luminosity.
func (a *Luminosity) Add(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value + other.BaseValue()}
}

// Subtract the given Luminosity to the current Luminosity.
func (a *Luminosity) Subtract(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value - other.BaseValue()}
}

// Multiply the given Luminosity to the current Luminosity.
func (a *Luminosity) Multiply(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value * other.BaseValue()}
}

// Divide the given Luminosity to the current Luminosity.
func (a *Luminosity) Divide(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value / other.BaseValue()}
}