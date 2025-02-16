// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// LuminosityUnits defines various units of Luminosity.
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

// LuminosityDto represents a Luminosity measurement with a numerical value and its corresponding unit.
type LuminosityDto struct {
    // Value is the numerical representation of the Luminosity.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Luminosity, as defined in the LuminosityUnits enumeration.
	Unit  LuminosityUnits `json:"unit"`
}

// LuminosityDtoFactory groups methods for creating and serializing LuminosityDto objects.
type LuminosityDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LuminosityDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LuminosityDtoFactory) FromJSON(data []byte) (*LuminosityDto, error) {
	a := LuminosityDto{}

    // Parse JSON into LuminosityDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LuminosityDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LuminosityDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Luminosity represents a measurement in a of Luminosity.
//
// Luminosity is an absolute measure of radiated electromagnetic power (light), the radiant power emitted by a light-emitting object.
type Luminosity struct {
	// value is the base measurement stored internally.
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

// LuminosityFactory groups methods for creating Luminosity instances.
type LuminosityFactory struct{}

// CreateLuminosity creates a new Luminosity instance from the given value and unit.
func (uf LuminosityFactory) CreateLuminosity(value float64, unit LuminosityUnits) (*Luminosity, error) {
	return newLuminosity(value, unit)
}

// FromDto converts a LuminosityDto to a Luminosity instance.
func (uf LuminosityFactory) FromDto(dto LuminosityDto) (*Luminosity, error) {
	return newLuminosity(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Luminosity instance.
func (uf LuminosityFactory) FromDtoJSON(data []byte) (*Luminosity, error) {
	unitDto, err := LuminosityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LuminosityDto from JSON: %w", err)
	}
	return LuminosityFactory{}.FromDto(*unitDto)
}


// FromWatts creates a new Luminosity instance from a value in Watts.
func (uf LuminosityFactory) FromWatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityWatt)
}

// FromSolarLuminosities creates a new Luminosity instance from a value in SolarLuminosities.
func (uf LuminosityFactory) FromSolarLuminosities(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminositySolarLuminosity)
}

// FromFemtowatts creates a new Luminosity instance from a value in Femtowatts.
func (uf LuminosityFactory) FromFemtowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityFemtowatt)
}

// FromPicowatts creates a new Luminosity instance from a value in Picowatts.
func (uf LuminosityFactory) FromPicowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityPicowatt)
}

// FromNanowatts creates a new Luminosity instance from a value in Nanowatts.
func (uf LuminosityFactory) FromNanowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityNanowatt)
}

// FromMicrowatts creates a new Luminosity instance from a value in Microwatts.
func (uf LuminosityFactory) FromMicrowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMicrowatt)
}

// FromMilliwatts creates a new Luminosity instance from a value in Milliwatts.
func (uf LuminosityFactory) FromMilliwatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMilliwatt)
}

// FromDeciwatts creates a new Luminosity instance from a value in Deciwatts.
func (uf LuminosityFactory) FromDeciwatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityDeciwatt)
}

// FromDecawatts creates a new Luminosity instance from a value in Decawatts.
func (uf LuminosityFactory) FromDecawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityDecawatt)
}

// FromKilowatts creates a new Luminosity instance from a value in Kilowatts.
func (uf LuminosityFactory) FromKilowatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityKilowatt)
}

// FromMegawatts creates a new Luminosity instance from a value in Megawatts.
func (uf LuminosityFactory) FromMegawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityMegawatt)
}

// FromGigawatts creates a new Luminosity instance from a value in Gigawatts.
func (uf LuminosityFactory) FromGigawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityGigawatt)
}

// FromTerawatts creates a new Luminosity instance from a value in Terawatts.
func (uf LuminosityFactory) FromTerawatts(value float64) (*Luminosity, error) {
	return newLuminosity(value, LuminosityTerawatt)
}

// FromPetawatts creates a new Luminosity instance from a value in Petawatts.
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

// BaseValue returns the base value of Luminosity in Watt unit (the base/default quantity).
func (a *Luminosity) BaseValue() float64 {
	return a.value
}


// Watts returns the Luminosity value in Watts.
//
// 
func (a *Luminosity) Watts() float64 {
	if a.wattsLazy != nil {
		return *a.wattsLazy
	}
	watts := a.convertFromBase(LuminosityWatt)
	a.wattsLazy = &watts
	return watts
}

// SolarLuminosities returns the Luminosity value in SolarLuminosities.
//
// 
func (a *Luminosity) SolarLuminosities() float64 {
	if a.solar_luminositiesLazy != nil {
		return *a.solar_luminositiesLazy
	}
	solar_luminosities := a.convertFromBase(LuminositySolarLuminosity)
	a.solar_luminositiesLazy = &solar_luminosities
	return solar_luminosities
}

// Femtowatts returns the Luminosity value in Femtowatts.
//
// 
func (a *Luminosity) Femtowatts() float64 {
	if a.femtowattsLazy != nil {
		return *a.femtowattsLazy
	}
	femtowatts := a.convertFromBase(LuminosityFemtowatt)
	a.femtowattsLazy = &femtowatts
	return femtowatts
}

// Picowatts returns the Luminosity value in Picowatts.
//
// 
func (a *Luminosity) Picowatts() float64 {
	if a.picowattsLazy != nil {
		return *a.picowattsLazy
	}
	picowatts := a.convertFromBase(LuminosityPicowatt)
	a.picowattsLazy = &picowatts
	return picowatts
}

// Nanowatts returns the Luminosity value in Nanowatts.
//
// 
func (a *Luminosity) Nanowatts() float64 {
	if a.nanowattsLazy != nil {
		return *a.nanowattsLazy
	}
	nanowatts := a.convertFromBase(LuminosityNanowatt)
	a.nanowattsLazy = &nanowatts
	return nanowatts
}

// Microwatts returns the Luminosity value in Microwatts.
//
// 
func (a *Luminosity) Microwatts() float64 {
	if a.microwattsLazy != nil {
		return *a.microwattsLazy
	}
	microwatts := a.convertFromBase(LuminosityMicrowatt)
	a.microwattsLazy = &microwatts
	return microwatts
}

// Milliwatts returns the Luminosity value in Milliwatts.
//
// 
func (a *Luminosity) Milliwatts() float64 {
	if a.milliwattsLazy != nil {
		return *a.milliwattsLazy
	}
	milliwatts := a.convertFromBase(LuminosityMilliwatt)
	a.milliwattsLazy = &milliwatts
	return milliwatts
}

// Deciwatts returns the Luminosity value in Deciwatts.
//
// 
func (a *Luminosity) Deciwatts() float64 {
	if a.deciwattsLazy != nil {
		return *a.deciwattsLazy
	}
	deciwatts := a.convertFromBase(LuminosityDeciwatt)
	a.deciwattsLazy = &deciwatts
	return deciwatts
}

// Decawatts returns the Luminosity value in Decawatts.
//
// 
func (a *Luminosity) Decawatts() float64 {
	if a.decawattsLazy != nil {
		return *a.decawattsLazy
	}
	decawatts := a.convertFromBase(LuminosityDecawatt)
	a.decawattsLazy = &decawatts
	return decawatts
}

// Kilowatts returns the Luminosity value in Kilowatts.
//
// 
func (a *Luminosity) Kilowatts() float64 {
	if a.kilowattsLazy != nil {
		return *a.kilowattsLazy
	}
	kilowatts := a.convertFromBase(LuminosityKilowatt)
	a.kilowattsLazy = &kilowatts
	return kilowatts
}

// Megawatts returns the Luminosity value in Megawatts.
//
// 
func (a *Luminosity) Megawatts() float64 {
	if a.megawattsLazy != nil {
		return *a.megawattsLazy
	}
	megawatts := a.convertFromBase(LuminosityMegawatt)
	a.megawattsLazy = &megawatts
	return megawatts
}

// Gigawatts returns the Luminosity value in Gigawatts.
//
// 
func (a *Luminosity) Gigawatts() float64 {
	if a.gigawattsLazy != nil {
		return *a.gigawattsLazy
	}
	gigawatts := a.convertFromBase(LuminosityGigawatt)
	a.gigawattsLazy = &gigawatts
	return gigawatts
}

// Terawatts returns the Luminosity value in Terawatts.
//
// 
func (a *Luminosity) Terawatts() float64 {
	if a.terawattsLazy != nil {
		return *a.terawattsLazy
	}
	terawatts := a.convertFromBase(LuminosityTerawatt)
	a.terawattsLazy = &terawatts
	return terawatts
}

// Petawatts returns the Luminosity value in Petawatts.
//
// 
func (a *Luminosity) Petawatts() float64 {
	if a.petawattsLazy != nil {
		return *a.petawattsLazy
	}
	petawatts := a.convertFromBase(LuminosityPetawatt)
	a.petawattsLazy = &petawatts
	return petawatts
}


// ToDto creates a LuminosityDto representation from the Luminosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Watt by default.
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

// ToDtoJSON creates a JSON representation of the Luminosity instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Watt by default.
func (a *Luminosity) ToDtoJSON(holdInUnit *LuminosityUnits) ([]byte, error) {
	// Convert to LuminosityDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Luminosity to a specific unit value.
// The function uses the provided unit type (LuminosityUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
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
		return math.NaN()
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

// String returns a string representation of the Luminosity in the default unit (Watt),
// formatted to two decimal places.
func (a Luminosity) String() string {
	return a.ToString(LuminosityWatt, 2)
}

// ToString formats the Luminosity value as a string with the specified unit and fractional digits.
// It converts the Luminosity to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Luminosity value will be converted (e.g., Watt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Luminosity with the unit abbreviation.
func (a *Luminosity) ToString(unit LuminosityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetLuminosityAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetLuminosityAbbreviation(unit))
}

// Equals checks if the given Luminosity is equal to the current Luminosity.
//
// Parameters:
//    other: The Luminosity to compare against.
//
// Returns:
//    bool: Returns true if both Luminosity are equal, false otherwise.
func (a *Luminosity) Equals(other *Luminosity) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Luminosity with another Luminosity.
// It returns -1 if the current Luminosity is less than the other Luminosity, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Luminosity to compare against.
//
// Returns:
//    int: -1 if the current Luminosity is less, 1 if greater, and 0 if equal.
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

// Add adds the given Luminosity to the current Luminosity and returns the result.
// The result is a new Luminosity instance with the sum of the values.
//
// Parameters:
//    other: The Luminosity to add to the current Luminosity.
//
// Returns:
//    *Luminosity: A new Luminosity instance representing the sum of both Luminosity.
func (a *Luminosity) Add(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Luminosity from the current Luminosity and returns the result.
// The result is a new Luminosity instance with the difference of the values.
//
// Parameters:
//    other: The Luminosity to subtract from the current Luminosity.
//
// Returns:
//    *Luminosity: A new Luminosity instance representing the difference of both Luminosity.
func (a *Luminosity) Subtract(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Luminosity by the given Luminosity and returns the result.
// The result is a new Luminosity instance with the product of the values.
//
// Parameters:
//    other: The Luminosity to multiply with the current Luminosity.
//
// Returns:
//    *Luminosity: A new Luminosity instance representing the product of both Luminosity.
func (a *Luminosity) Multiply(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value * other.BaseValue()}
}

// Divide divides the current Luminosity by the given Luminosity and returns the result.
// The result is a new Luminosity instance with the quotient of the values.
//
// Parameters:
//    other: The Luminosity to divide the current Luminosity by.
//
// Returns:
//    *Luminosity: A new Luminosity instance representing the quotient of both Luminosity.
func (a *Luminosity) Divide(other *Luminosity) *Luminosity {
	return &Luminosity{value: a.value / other.BaseValue()}
}

// GetLuminosityAbbreviation gets the unit abbreviation.
func GetLuminosityAbbreviation(unit LuminosityUnits) string {
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