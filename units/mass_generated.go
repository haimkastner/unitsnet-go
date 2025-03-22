// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// MassUnits defines various units of Mass.
type MassUnits string

const (
    
        // 
        MassGram MassUnits = "Gram"
        // 
        MassTonne MassUnits = "Tonne"
        // The short ton is a unit of mass equal to 2,000 pounds (907.18474 kg), that is most commonly used in the United States – known there simply as the ton.
        MassShortTon MassUnits = "ShortTon"
        // Long ton (weight ton or Imperial ton) is a unit of mass equal to 2,240 pounds (1,016 kg) and is the name for the unit called the "ton" in the avoirdupois or Imperial system of measurements that was used in the United Kingdom and several other Commonwealth countries before metrication.
        MassLongTon MassUnits = "LongTon"
        // The pound or pound-mass (abbreviations: lb, lbm) is a unit of mass used in the imperial, United States customary and other systems of measurement. A number of different definitions have been used, the most common today being the international avoirdupois pound which is legally defined as exactly 0.45359237 kilograms, and which is divided into 16 avoirdupois ounces.
        MassPound MassUnits = "Pound"
        // The international avoirdupois ounce (abbreviated oz) is defined as exactly 28.349523125 g under the international yard and pound agreement of 1959, signed by the United States and countries of the Commonwealth of Nations. 16 oz make up an avoirdupois pound.
        MassOunce MassUnits = "Ounce"
        // The slug (abbreviation slug) is a unit of mass that is accelerated by 1 ft/s² when a force of one pound (lbf) is exerted on it.
        MassSlug MassUnits = "Slug"
        // The stone (abbreviation st) is a unit of mass equal to 14 pounds avoirdupois (about 6.35 kilograms) used in Great Britain and Ireland for measuring human body weight.
        MassStone MassUnits = "Stone"
        // The short hundredweight (abbreviation cwt) is a unit of mass equal to 100 pounds in US and Canada. In British English, the short hundredweight is referred to as the "cental".
        MassShortHundredweight MassUnits = "ShortHundredweight"
        // The long or imperial hundredweight (abbreviation cwt) is a unit of mass equal to 112 pounds in US and Canada.
        MassLongHundredweight MassUnits = "LongHundredweight"
        // A grain is a unit of measurement of mass, and in the troy weight, avoirdupois, and Apothecaries' system, equal to exactly 64.79891 milligrams.
        MassGrain MassUnits = "Grain"
        // Solar mass is a ratio unit to the mass of the solar system star, the sun.
        MassSolarMass MassUnits = "SolarMass"
        // Earth mass is a ratio unit to the mass of planet Earth.
        MassEarthMass MassUnits = "EarthMass"
        // 
        MassFemtogram MassUnits = "Femtogram"
        // 
        MassPicogram MassUnits = "Picogram"
        // 
        MassNanogram MassUnits = "Nanogram"
        // 
        MassMicrogram MassUnits = "Microgram"
        // 
        MassMilligram MassUnits = "Milligram"
        // 
        MassCentigram MassUnits = "Centigram"
        // 
        MassDecigram MassUnits = "Decigram"
        // 
        MassDecagram MassUnits = "Decagram"
        // 
        MassHectogram MassUnits = "Hectogram"
        // 
        MassKilogram MassUnits = "Kilogram"
        // 
        MassKilotonne MassUnits = "Kilotonne"
        // 
        MassMegatonne MassUnits = "Megatonne"
        // 
        MassKilopound MassUnits = "Kilopound"
        // 
        MassMegapound MassUnits = "Megapound"
)

var internalMassUnitsMap = map[MassUnits]bool{
	
	MassGram: true,
	MassTonne: true,
	MassShortTon: true,
	MassLongTon: true,
	MassPound: true,
	MassOunce: true,
	MassSlug: true,
	MassStone: true,
	MassShortHundredweight: true,
	MassLongHundredweight: true,
	MassGrain: true,
	MassSolarMass: true,
	MassEarthMass: true,
	MassFemtogram: true,
	MassPicogram: true,
	MassNanogram: true,
	MassMicrogram: true,
	MassMilligram: true,
	MassCentigram: true,
	MassDecigram: true,
	MassDecagram: true,
	MassHectogram: true,
	MassKilogram: true,
	MassKilotonne: true,
	MassMegatonne: true,
	MassKilopound: true,
	MassMegapound: true,
}

// MassDto represents a Mass measurement with a numerical value and its corresponding unit.
type MassDto struct {
    // Value is the numerical representation of the Mass.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Mass, as defined in the MassUnits enumeration.
	Unit  MassUnits `json:"unit" validate:"required,oneof=Gram,Tonne,ShortTon,LongTon,Pound,Ounce,Slug,Stone,ShortHundredweight,LongHundredweight,Grain,SolarMass,EarthMass,Femtogram,Picogram,Nanogram,Microgram,Milligram,Centigram,Decigram,Decagram,Hectogram,Kilogram,Kilotonne,Megatonne,Kilopound,Megapound"`
}

// MassDtoFactory groups methods for creating and serializing MassDto objects.
type MassDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a MassDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf MassDtoFactory) FromJSON(data []byte) (*MassDto, error) {
	a := MassDto{}

    // Parse JSON into MassDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a MassDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a MassDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Mass represents a measurement in a of Mass.
//
// In physics, mass (from Greek μᾶζα "barley cake, lump [of dough]") is a property of a physical system or body, giving rise to the phenomena of the body's resistance to being accelerated by a force and the strength of its mutual gravitational attraction with other bodies. Instruments such as mass balances or scales use those phenomena to measure mass. The SI unit of mass is the kilogram (kg).
type Mass struct {
	// value is the base measurement stored internally.
	value       float64
    
    gramsLazy *float64 
    tonnesLazy *float64 
    short_tonsLazy *float64 
    long_tonsLazy *float64 
    poundsLazy *float64 
    ouncesLazy *float64 
    slugsLazy *float64 
    stoneLazy *float64 
    short_hundredweightLazy *float64 
    long_hundredweightLazy *float64 
    grainsLazy *float64 
    solar_massesLazy *float64 
    earth_massesLazy *float64 
    femtogramsLazy *float64 
    picogramsLazy *float64 
    nanogramsLazy *float64 
    microgramsLazy *float64 
    milligramsLazy *float64 
    centigramsLazy *float64 
    decigramsLazy *float64 
    decagramsLazy *float64 
    hectogramsLazy *float64 
    kilogramsLazy *float64 
    kilotonnesLazy *float64 
    megatonnesLazy *float64 
    kilopoundsLazy *float64 
    megapoundsLazy *float64 
}

// MassFactory groups methods for creating Mass instances.
type MassFactory struct{}

// CreateMass creates a new Mass instance from the given value and unit.
func (uf MassFactory) CreateMass(value float64, unit MassUnits) (*Mass, error) {
	return newMass(value, unit)
}

// FromDto converts a MassDto to a Mass instance.
func (uf MassFactory) FromDto(dto MassDto) (*Mass, error) {
	return newMass(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Mass instance.
func (uf MassFactory) FromDtoJSON(data []byte) (*Mass, error) {
	unitDto, err := MassDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse MassDto from JSON: %w", err)
	}
	return MassFactory{}.FromDto(*unitDto)
}


// FromGrams creates a new Mass instance from a value in Grams.
func (uf MassFactory) FromGrams(value float64) (*Mass, error) {
	return newMass(value, MassGram)
}

// FromTonnes creates a new Mass instance from a value in Tonnes.
func (uf MassFactory) FromTonnes(value float64) (*Mass, error) {
	return newMass(value, MassTonne)
}

// FromShortTons creates a new Mass instance from a value in ShortTons.
func (uf MassFactory) FromShortTons(value float64) (*Mass, error) {
	return newMass(value, MassShortTon)
}

// FromLongTons creates a new Mass instance from a value in LongTons.
func (uf MassFactory) FromLongTons(value float64) (*Mass, error) {
	return newMass(value, MassLongTon)
}

// FromPounds creates a new Mass instance from a value in Pounds.
func (uf MassFactory) FromPounds(value float64) (*Mass, error) {
	return newMass(value, MassPound)
}

// FromOunces creates a new Mass instance from a value in Ounces.
func (uf MassFactory) FromOunces(value float64) (*Mass, error) {
	return newMass(value, MassOunce)
}

// FromSlugs creates a new Mass instance from a value in Slugs.
func (uf MassFactory) FromSlugs(value float64) (*Mass, error) {
	return newMass(value, MassSlug)
}

// FromStone creates a new Mass instance from a value in Stone.
func (uf MassFactory) FromStone(value float64) (*Mass, error) {
	return newMass(value, MassStone)
}

// FromShortHundredweight creates a new Mass instance from a value in ShortHundredweight.
func (uf MassFactory) FromShortHundredweight(value float64) (*Mass, error) {
	return newMass(value, MassShortHundredweight)
}

// FromLongHundredweight creates a new Mass instance from a value in LongHundredweight.
func (uf MassFactory) FromLongHundredweight(value float64) (*Mass, error) {
	return newMass(value, MassLongHundredweight)
}

// FromGrains creates a new Mass instance from a value in Grains.
func (uf MassFactory) FromGrains(value float64) (*Mass, error) {
	return newMass(value, MassGrain)
}

// FromSolarMasses creates a new Mass instance from a value in SolarMasses.
func (uf MassFactory) FromSolarMasses(value float64) (*Mass, error) {
	return newMass(value, MassSolarMass)
}

// FromEarthMasses creates a new Mass instance from a value in EarthMasses.
func (uf MassFactory) FromEarthMasses(value float64) (*Mass, error) {
	return newMass(value, MassEarthMass)
}

// FromFemtograms creates a new Mass instance from a value in Femtograms.
func (uf MassFactory) FromFemtograms(value float64) (*Mass, error) {
	return newMass(value, MassFemtogram)
}

// FromPicograms creates a new Mass instance from a value in Picograms.
func (uf MassFactory) FromPicograms(value float64) (*Mass, error) {
	return newMass(value, MassPicogram)
}

// FromNanograms creates a new Mass instance from a value in Nanograms.
func (uf MassFactory) FromNanograms(value float64) (*Mass, error) {
	return newMass(value, MassNanogram)
}

// FromMicrograms creates a new Mass instance from a value in Micrograms.
func (uf MassFactory) FromMicrograms(value float64) (*Mass, error) {
	return newMass(value, MassMicrogram)
}

// FromMilligrams creates a new Mass instance from a value in Milligrams.
func (uf MassFactory) FromMilligrams(value float64) (*Mass, error) {
	return newMass(value, MassMilligram)
}

// FromCentigrams creates a new Mass instance from a value in Centigrams.
func (uf MassFactory) FromCentigrams(value float64) (*Mass, error) {
	return newMass(value, MassCentigram)
}

// FromDecigrams creates a new Mass instance from a value in Decigrams.
func (uf MassFactory) FromDecigrams(value float64) (*Mass, error) {
	return newMass(value, MassDecigram)
}

// FromDecagrams creates a new Mass instance from a value in Decagrams.
func (uf MassFactory) FromDecagrams(value float64) (*Mass, error) {
	return newMass(value, MassDecagram)
}

// FromHectograms creates a new Mass instance from a value in Hectograms.
func (uf MassFactory) FromHectograms(value float64) (*Mass, error) {
	return newMass(value, MassHectogram)
}

// FromKilograms creates a new Mass instance from a value in Kilograms.
func (uf MassFactory) FromKilograms(value float64) (*Mass, error) {
	return newMass(value, MassKilogram)
}

// FromKilotonnes creates a new Mass instance from a value in Kilotonnes.
func (uf MassFactory) FromKilotonnes(value float64) (*Mass, error) {
	return newMass(value, MassKilotonne)
}

// FromMegatonnes creates a new Mass instance from a value in Megatonnes.
func (uf MassFactory) FromMegatonnes(value float64) (*Mass, error) {
	return newMass(value, MassMegatonne)
}

// FromKilopounds creates a new Mass instance from a value in Kilopounds.
func (uf MassFactory) FromKilopounds(value float64) (*Mass, error) {
	return newMass(value, MassKilopound)
}

// FromMegapounds creates a new Mass instance from a value in Megapounds.
func (uf MassFactory) FromMegapounds(value float64) (*Mass, error) {
	return newMass(value, MassMegapound)
}


// newMass creates a new Mass.
func newMass(value float64, fromUnit MassUnits) (*Mass, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalMassUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in MassUnits", fromUnit)
	}
	a := &Mass{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Mass in Kilogram unit (the base/default quantity).
func (a *Mass) BaseValue() float64 {
	return a.value
}


// Grams returns the Mass value in Grams.
//
// 
func (a *Mass) Grams() float64 {
	if a.gramsLazy != nil {
		return *a.gramsLazy
	}
	grams := a.convertFromBase(MassGram)
	a.gramsLazy = &grams
	return grams
}

// Tonnes returns the Mass value in Tonnes.
//
// 
func (a *Mass) Tonnes() float64 {
	if a.tonnesLazy != nil {
		return *a.tonnesLazy
	}
	tonnes := a.convertFromBase(MassTonne)
	a.tonnesLazy = &tonnes
	return tonnes
}

// ShortTons returns the Mass value in ShortTons.
//
// The short ton is a unit of mass equal to 2,000 pounds (907.18474 kg), that is most commonly used in the United States – known there simply as the ton.
func (a *Mass) ShortTons() float64 {
	if a.short_tonsLazy != nil {
		return *a.short_tonsLazy
	}
	short_tons := a.convertFromBase(MassShortTon)
	a.short_tonsLazy = &short_tons
	return short_tons
}

// LongTons returns the Mass value in LongTons.
//
// Long ton (weight ton or Imperial ton) is a unit of mass equal to 2,240 pounds (1,016 kg) and is the name for the unit called the "ton" in the avoirdupois or Imperial system of measurements that was used in the United Kingdom and several other Commonwealth countries before metrication.
func (a *Mass) LongTons() float64 {
	if a.long_tonsLazy != nil {
		return *a.long_tonsLazy
	}
	long_tons := a.convertFromBase(MassLongTon)
	a.long_tonsLazy = &long_tons
	return long_tons
}

// Pounds returns the Mass value in Pounds.
//
// The pound or pound-mass (abbreviations: lb, lbm) is a unit of mass used in the imperial, United States customary and other systems of measurement. A number of different definitions have been used, the most common today being the international avoirdupois pound which is legally defined as exactly 0.45359237 kilograms, and which is divided into 16 avoirdupois ounces.
func (a *Mass) Pounds() float64 {
	if a.poundsLazy != nil {
		return *a.poundsLazy
	}
	pounds := a.convertFromBase(MassPound)
	a.poundsLazy = &pounds
	return pounds
}

// Ounces returns the Mass value in Ounces.
//
// The international avoirdupois ounce (abbreviated oz) is defined as exactly 28.349523125 g under the international yard and pound agreement of 1959, signed by the United States and countries of the Commonwealth of Nations. 16 oz make up an avoirdupois pound.
func (a *Mass) Ounces() float64 {
	if a.ouncesLazy != nil {
		return *a.ouncesLazy
	}
	ounces := a.convertFromBase(MassOunce)
	a.ouncesLazy = &ounces
	return ounces
}

// Slugs returns the Mass value in Slugs.
//
// The slug (abbreviation slug) is a unit of mass that is accelerated by 1 ft/s² when a force of one pound (lbf) is exerted on it.
func (a *Mass) Slugs() float64 {
	if a.slugsLazy != nil {
		return *a.slugsLazy
	}
	slugs := a.convertFromBase(MassSlug)
	a.slugsLazy = &slugs
	return slugs
}

// Stone returns the Mass value in Stone.
//
// The stone (abbreviation st) is a unit of mass equal to 14 pounds avoirdupois (about 6.35 kilograms) used in Great Britain and Ireland for measuring human body weight.
func (a *Mass) Stone() float64 {
	if a.stoneLazy != nil {
		return *a.stoneLazy
	}
	stone := a.convertFromBase(MassStone)
	a.stoneLazy = &stone
	return stone
}

// ShortHundredweight returns the Mass value in ShortHundredweight.
//
// The short hundredweight (abbreviation cwt) is a unit of mass equal to 100 pounds in US and Canada. In British English, the short hundredweight is referred to as the "cental".
func (a *Mass) ShortHundredweight() float64 {
	if a.short_hundredweightLazy != nil {
		return *a.short_hundredweightLazy
	}
	short_hundredweight := a.convertFromBase(MassShortHundredweight)
	a.short_hundredweightLazy = &short_hundredweight
	return short_hundredweight
}

// LongHundredweight returns the Mass value in LongHundredweight.
//
// The long or imperial hundredweight (abbreviation cwt) is a unit of mass equal to 112 pounds in US and Canada.
func (a *Mass) LongHundredweight() float64 {
	if a.long_hundredweightLazy != nil {
		return *a.long_hundredweightLazy
	}
	long_hundredweight := a.convertFromBase(MassLongHundredweight)
	a.long_hundredweightLazy = &long_hundredweight
	return long_hundredweight
}

// Grains returns the Mass value in Grains.
//
// A grain is a unit of measurement of mass, and in the troy weight, avoirdupois, and Apothecaries' system, equal to exactly 64.79891 milligrams.
func (a *Mass) Grains() float64 {
	if a.grainsLazy != nil {
		return *a.grainsLazy
	}
	grains := a.convertFromBase(MassGrain)
	a.grainsLazy = &grains
	return grains
}

// SolarMasses returns the Mass value in SolarMasses.
//
// Solar mass is a ratio unit to the mass of the solar system star, the sun.
func (a *Mass) SolarMasses() float64 {
	if a.solar_massesLazy != nil {
		return *a.solar_massesLazy
	}
	solar_masses := a.convertFromBase(MassSolarMass)
	a.solar_massesLazy = &solar_masses
	return solar_masses
}

// EarthMasses returns the Mass value in EarthMasses.
//
// Earth mass is a ratio unit to the mass of planet Earth.
func (a *Mass) EarthMasses() float64 {
	if a.earth_massesLazy != nil {
		return *a.earth_massesLazy
	}
	earth_masses := a.convertFromBase(MassEarthMass)
	a.earth_massesLazy = &earth_masses
	return earth_masses
}

// Femtograms returns the Mass value in Femtograms.
//
// 
func (a *Mass) Femtograms() float64 {
	if a.femtogramsLazy != nil {
		return *a.femtogramsLazy
	}
	femtograms := a.convertFromBase(MassFemtogram)
	a.femtogramsLazy = &femtograms
	return femtograms
}

// Picograms returns the Mass value in Picograms.
//
// 
func (a *Mass) Picograms() float64 {
	if a.picogramsLazy != nil {
		return *a.picogramsLazy
	}
	picograms := a.convertFromBase(MassPicogram)
	a.picogramsLazy = &picograms
	return picograms
}

// Nanograms returns the Mass value in Nanograms.
//
// 
func (a *Mass) Nanograms() float64 {
	if a.nanogramsLazy != nil {
		return *a.nanogramsLazy
	}
	nanograms := a.convertFromBase(MassNanogram)
	a.nanogramsLazy = &nanograms
	return nanograms
}

// Micrograms returns the Mass value in Micrograms.
//
// 
func (a *Mass) Micrograms() float64 {
	if a.microgramsLazy != nil {
		return *a.microgramsLazy
	}
	micrograms := a.convertFromBase(MassMicrogram)
	a.microgramsLazy = &micrograms
	return micrograms
}

// Milligrams returns the Mass value in Milligrams.
//
// 
func (a *Mass) Milligrams() float64 {
	if a.milligramsLazy != nil {
		return *a.milligramsLazy
	}
	milligrams := a.convertFromBase(MassMilligram)
	a.milligramsLazy = &milligrams
	return milligrams
}

// Centigrams returns the Mass value in Centigrams.
//
// 
func (a *Mass) Centigrams() float64 {
	if a.centigramsLazy != nil {
		return *a.centigramsLazy
	}
	centigrams := a.convertFromBase(MassCentigram)
	a.centigramsLazy = &centigrams
	return centigrams
}

// Decigrams returns the Mass value in Decigrams.
//
// 
func (a *Mass) Decigrams() float64 {
	if a.decigramsLazy != nil {
		return *a.decigramsLazy
	}
	decigrams := a.convertFromBase(MassDecigram)
	a.decigramsLazy = &decigrams
	return decigrams
}

// Decagrams returns the Mass value in Decagrams.
//
// 
func (a *Mass) Decagrams() float64 {
	if a.decagramsLazy != nil {
		return *a.decagramsLazy
	}
	decagrams := a.convertFromBase(MassDecagram)
	a.decagramsLazy = &decagrams
	return decagrams
}

// Hectograms returns the Mass value in Hectograms.
//
// 
func (a *Mass) Hectograms() float64 {
	if a.hectogramsLazy != nil {
		return *a.hectogramsLazy
	}
	hectograms := a.convertFromBase(MassHectogram)
	a.hectogramsLazy = &hectograms
	return hectograms
}

// Kilograms returns the Mass value in Kilograms.
//
// 
func (a *Mass) Kilograms() float64 {
	if a.kilogramsLazy != nil {
		return *a.kilogramsLazy
	}
	kilograms := a.convertFromBase(MassKilogram)
	a.kilogramsLazy = &kilograms
	return kilograms
}

// Kilotonnes returns the Mass value in Kilotonnes.
//
// 
func (a *Mass) Kilotonnes() float64 {
	if a.kilotonnesLazy != nil {
		return *a.kilotonnesLazy
	}
	kilotonnes := a.convertFromBase(MassKilotonne)
	a.kilotonnesLazy = &kilotonnes
	return kilotonnes
}

// Megatonnes returns the Mass value in Megatonnes.
//
// 
func (a *Mass) Megatonnes() float64 {
	if a.megatonnesLazy != nil {
		return *a.megatonnesLazy
	}
	megatonnes := a.convertFromBase(MassMegatonne)
	a.megatonnesLazy = &megatonnes
	return megatonnes
}

// Kilopounds returns the Mass value in Kilopounds.
//
// 
func (a *Mass) Kilopounds() float64 {
	if a.kilopoundsLazy != nil {
		return *a.kilopoundsLazy
	}
	kilopounds := a.convertFromBase(MassKilopound)
	a.kilopoundsLazy = &kilopounds
	return kilopounds
}

// Megapounds returns the Mass value in Megapounds.
//
// 
func (a *Mass) Megapounds() float64 {
	if a.megapoundsLazy != nil {
		return *a.megapoundsLazy
	}
	megapounds := a.convertFromBase(MassMegapound)
	a.megapoundsLazy = &megapounds
	return megapounds
}


// ToDto creates a MassDto representation from the Mass instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kilogram by default.
func (a *Mass) ToDto(holdInUnit *MassUnits) MassDto {
	if holdInUnit == nil {
		defaultUnit := MassKilogram // Default value
		holdInUnit = &defaultUnit
	}

	return MassDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Mass instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Kilogram by default.
func (a *Mass) ToDtoJSON(holdInUnit *MassUnits) ([]byte, error) {
	// Convert to MassDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Mass to a specific unit value.
// The function uses the provided unit type (MassUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Mass) Convert(toUnit MassUnits) float64 {
	switch toUnit { 
    case MassGram:
		return a.Grams()
    case MassTonne:
		return a.Tonnes()
    case MassShortTon:
		return a.ShortTons()
    case MassLongTon:
		return a.LongTons()
    case MassPound:
		return a.Pounds()
    case MassOunce:
		return a.Ounces()
    case MassSlug:
		return a.Slugs()
    case MassStone:
		return a.Stone()
    case MassShortHundredweight:
		return a.ShortHundredweight()
    case MassLongHundredweight:
		return a.LongHundredweight()
    case MassGrain:
		return a.Grains()
    case MassSolarMass:
		return a.SolarMasses()
    case MassEarthMass:
		return a.EarthMasses()
    case MassFemtogram:
		return a.Femtograms()
    case MassPicogram:
		return a.Picograms()
    case MassNanogram:
		return a.Nanograms()
    case MassMicrogram:
		return a.Micrograms()
    case MassMilligram:
		return a.Milligrams()
    case MassCentigram:
		return a.Centigrams()
    case MassDecigram:
		return a.Decigrams()
    case MassDecagram:
		return a.Decagrams()
    case MassHectogram:
		return a.Hectograms()
    case MassKilogram:
		return a.Kilograms()
    case MassKilotonne:
		return a.Kilotonnes()
    case MassMegatonne:
		return a.Megatonnes()
    case MassKilopound:
		return a.Kilopounds()
    case MassMegapound:
		return a.Megapounds()
	default:
		return math.NaN()
	}
}

func (a *Mass) convertFromBase(toUnit MassUnits) float64 {
    value := a.value
	switch toUnit { 
	case MassGram:
		return (value * 1e3) 
	case MassTonne:
		return (value / 1e3) 
	case MassShortTon:
		return (value / 9.0718474e2) 
	case MassLongTon:
		return (value / 1.0160469088e3) 
	case MassPound:
		return (value / 0.45359237) 
	case MassOunce:
		return (value / 0.028349523125) 
	case MassSlug:
		return (value * 6.852176556196105e-2) 
	case MassStone:
		return (value * 0.1574731728702698) 
	case MassShortHundredweight:
		return (value * 0.022046226218487758) 
	case MassLongHundredweight:
		return (value * 0.01968413055222121) 
	case MassGrain:
		return (value * 15432.358352941431) 
	case MassSolarMass:
		return (value / 1.98947e30) 
	case MassEarthMass:
		return (value / 5.9722e+24) 
	case MassFemtogram:
		return ((value * 1e3) / 1e-15) 
	case MassPicogram:
		return ((value * 1e3) / 1e-12) 
	case MassNanogram:
		return ((value * 1e3) / 1e-09) 
	case MassMicrogram:
		return ((value * 1e3) / 1e-06) 
	case MassMilligram:
		return ((value * 1e3) / 0.001) 
	case MassCentigram:
		return ((value * 1e3) / 0.01) 
	case MassDecigram:
		return ((value * 1e3) / 0.1) 
	case MassDecagram:
		return ((value * 1e3) / 10.0) 
	case MassHectogram:
		return ((value * 1e3) / 100.0) 
	case MassKilogram:
		return ((value * 1e3) / 1000.0) 
	case MassKilotonne:
		return ((value / 1e3) / 1000.0) 
	case MassMegatonne:
		return ((value / 1e3) / 1000000.0) 
	case MassKilopound:
		return ((value / 0.45359237) / 1000.0) 
	case MassMegapound:
		return ((value / 0.45359237) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Mass) convertToBase(value float64, fromUnit MassUnits) float64 {
	switch fromUnit { 
	case MassGram:
		return (value / 1e3) 
	case MassTonne:
		return (value * 1e3) 
	case MassShortTon:
		return (value * 9.0718474e2) 
	case MassLongTon:
		return (value * 1.0160469088e3) 
	case MassPound:
		return (value * 0.45359237) 
	case MassOunce:
		return (value * 0.028349523125) 
	case MassSlug:
		return (value / 6.852176556196105e-2) 
	case MassStone:
		return (value / 0.1574731728702698) 
	case MassShortHundredweight:
		return (value / 0.022046226218487758) 
	case MassLongHundredweight:
		return (value / 0.01968413055222121) 
	case MassGrain:
		return (value / 15432.358352941431) 
	case MassSolarMass:
		return (value * 1.98947e30) 
	case MassEarthMass:
		return (value * 5.9722e+24) 
	case MassFemtogram:
		return ((value / 1e3) * 1e-15) 
	case MassPicogram:
		return ((value / 1e3) * 1e-12) 
	case MassNanogram:
		return ((value / 1e3) * 1e-09) 
	case MassMicrogram:
		return ((value / 1e3) * 1e-06) 
	case MassMilligram:
		return ((value / 1e3) * 0.001) 
	case MassCentigram:
		return ((value / 1e3) * 0.01) 
	case MassDecigram:
		return ((value / 1e3) * 0.1) 
	case MassDecagram:
		return ((value / 1e3) * 10.0) 
	case MassHectogram:
		return ((value / 1e3) * 100.0) 
	case MassKilogram:
		return ((value / 1e3) * 1000.0) 
	case MassKilotonne:
		return ((value * 1e3) * 1000.0) 
	case MassMegatonne:
		return ((value * 1e3) * 1000000.0) 
	case MassKilopound:
		return ((value * 0.45359237) * 1000.0) 
	case MassMegapound:
		return ((value * 0.45359237) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Mass in the default unit (Kilogram),
// formatted to two decimal places.
func (a Mass) String() string {
	return a.ToString(MassKilogram, 2)
}

// ToString formats the Mass value as a string with the specified unit and fractional digits.
// It converts the Mass to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Mass value will be converted (e.g., Kilogram))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Mass with the unit abbreviation.
func (a *Mass) ToString(unit MassUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetMassAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetMassAbbreviation(unit))
}

// Equals checks if the given Mass is equal to the current Mass.
//
// Parameters:
//    other: The Mass to compare against.
//
// Returns:
//    bool: Returns true if both Mass are equal, false otherwise.
func (a *Mass) Equals(other *Mass) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Mass with another Mass.
// It returns -1 if the current Mass is less than the other Mass, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Mass to compare against.
//
// Returns:
//    int: -1 if the current Mass is less, 1 if greater, and 0 if equal.
func (a *Mass) CompareTo(other *Mass) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Mass to the current Mass and returns the result.
// The result is a new Mass instance with the sum of the values.
//
// Parameters:
//    other: The Mass to add to the current Mass.
//
// Returns:
//    *Mass: A new Mass instance representing the sum of both Mass.
func (a *Mass) Add(other *Mass) *Mass {
	return &Mass{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Mass from the current Mass and returns the result.
// The result is a new Mass instance with the difference of the values.
//
// Parameters:
//    other: The Mass to subtract from the current Mass.
//
// Returns:
//    *Mass: A new Mass instance representing the difference of both Mass.
func (a *Mass) Subtract(other *Mass) *Mass {
	return &Mass{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Mass by the given Mass and returns the result.
// The result is a new Mass instance with the product of the values.
//
// Parameters:
//    other: The Mass to multiply with the current Mass.
//
// Returns:
//    *Mass: A new Mass instance representing the product of both Mass.
func (a *Mass) Multiply(other *Mass) *Mass {
	return &Mass{value: a.value * other.BaseValue()}
}

// Divide divides the current Mass by the given Mass and returns the result.
// The result is a new Mass instance with the quotient of the values.
//
// Parameters:
//    other: The Mass to divide the current Mass by.
//
// Returns:
//    *Mass: A new Mass instance representing the quotient of both Mass.
func (a *Mass) Divide(other *Mass) *Mass {
	return &Mass{value: a.value / other.BaseValue()}
}

// GetMassAbbreviation gets the unit abbreviation.
func GetMassAbbreviation(unit MassUnits) string {
	switch unit { 
	case MassGram:
		return "g" 
	case MassTonne:
		return "t" 
	case MassShortTon:
		return "t (short)" 
	case MassLongTon:
		return "long tn" 
	case MassPound:
		return "lb" 
	case MassOunce:
		return "oz" 
	case MassSlug:
		return "slug" 
	case MassStone:
		return "st" 
	case MassShortHundredweight:
		return "cwt" 
	case MassLongHundredweight:
		return "cwt" 
	case MassGrain:
		return "gr" 
	case MassSolarMass:
		return "M☉" 
	case MassEarthMass:
		return "em" 
	case MassFemtogram:
		return "fg" 
	case MassPicogram:
		return "pg" 
	case MassNanogram:
		return "ng" 
	case MassMicrogram:
		return "μg" 
	case MassMilligram:
		return "mg" 
	case MassCentigram:
		return "cg" 
	case MassDecigram:
		return "dg" 
	case MassDecagram:
		return "dag" 
	case MassHectogram:
		return "hg" 
	case MassKilogram:
		return "kg" 
	case MassKilotonne:
		return "kt" 
	case MassMegatonne:
		return "Mt" 
	case MassKilopound:
		return "klb" 
	case MassMegapound:
		return "Mlb" 
	default:
		return ""
	}
}