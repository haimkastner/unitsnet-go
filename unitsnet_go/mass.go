package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// MassUnits enumeration
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

// MassDto represents an Mass
type MassDto struct {
	Value float64
	Unit  MassUnits
}

// MassDtoFactory struct to group related functions
type MassDtoFactory struct{}

func (udf MassDtoFactory) FromJSON(data []byte) (*MassDto, error) {
	a := MassDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a MassDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Mass struct
type Mass struct {
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

// MassFactory struct to group related functions
type MassFactory struct{}

func (uf MassFactory) CreateMass(value float64, unit MassUnits) (*Mass, error) {
	return newMass(value, unit)
}

func (uf MassFactory) FromDto(dto MassDto) (*Mass, error) {
	return newMass(dto.Value, dto.Unit)
}

func (uf MassFactory) FromDtoJSON(data []byte) (*Mass, error) {
	unitDto, err := MassDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return MassFactory{}.FromDto(*unitDto)
}


// FromGram creates a new Mass instance from Gram.
func (uf MassFactory) FromGrams(value float64) (*Mass, error) {
	return newMass(value, MassGram)
}

// FromTonne creates a new Mass instance from Tonne.
func (uf MassFactory) FromTonnes(value float64) (*Mass, error) {
	return newMass(value, MassTonne)
}

// FromShortTon creates a new Mass instance from ShortTon.
func (uf MassFactory) FromShortTons(value float64) (*Mass, error) {
	return newMass(value, MassShortTon)
}

// FromLongTon creates a new Mass instance from LongTon.
func (uf MassFactory) FromLongTons(value float64) (*Mass, error) {
	return newMass(value, MassLongTon)
}

// FromPound creates a new Mass instance from Pound.
func (uf MassFactory) FromPounds(value float64) (*Mass, error) {
	return newMass(value, MassPound)
}

// FromOunce creates a new Mass instance from Ounce.
func (uf MassFactory) FromOunces(value float64) (*Mass, error) {
	return newMass(value, MassOunce)
}

// FromSlug creates a new Mass instance from Slug.
func (uf MassFactory) FromSlugs(value float64) (*Mass, error) {
	return newMass(value, MassSlug)
}

// FromStone creates a new Mass instance from Stone.
func (uf MassFactory) FromStone(value float64) (*Mass, error) {
	return newMass(value, MassStone)
}

// FromShortHundredweight creates a new Mass instance from ShortHundredweight.
func (uf MassFactory) FromShortHundredweight(value float64) (*Mass, error) {
	return newMass(value, MassShortHundredweight)
}

// FromLongHundredweight creates a new Mass instance from LongHundredweight.
func (uf MassFactory) FromLongHundredweight(value float64) (*Mass, error) {
	return newMass(value, MassLongHundredweight)
}

// FromGrain creates a new Mass instance from Grain.
func (uf MassFactory) FromGrains(value float64) (*Mass, error) {
	return newMass(value, MassGrain)
}

// FromSolarMass creates a new Mass instance from SolarMass.
func (uf MassFactory) FromSolarMasses(value float64) (*Mass, error) {
	return newMass(value, MassSolarMass)
}

// FromEarthMass creates a new Mass instance from EarthMass.
func (uf MassFactory) FromEarthMasses(value float64) (*Mass, error) {
	return newMass(value, MassEarthMass)
}

// FromFemtogram creates a new Mass instance from Femtogram.
func (uf MassFactory) FromFemtograms(value float64) (*Mass, error) {
	return newMass(value, MassFemtogram)
}

// FromPicogram creates a new Mass instance from Picogram.
func (uf MassFactory) FromPicograms(value float64) (*Mass, error) {
	return newMass(value, MassPicogram)
}

// FromNanogram creates a new Mass instance from Nanogram.
func (uf MassFactory) FromNanograms(value float64) (*Mass, error) {
	return newMass(value, MassNanogram)
}

// FromMicrogram creates a new Mass instance from Microgram.
func (uf MassFactory) FromMicrograms(value float64) (*Mass, error) {
	return newMass(value, MassMicrogram)
}

// FromMilligram creates a new Mass instance from Milligram.
func (uf MassFactory) FromMilligrams(value float64) (*Mass, error) {
	return newMass(value, MassMilligram)
}

// FromCentigram creates a new Mass instance from Centigram.
func (uf MassFactory) FromCentigrams(value float64) (*Mass, error) {
	return newMass(value, MassCentigram)
}

// FromDecigram creates a new Mass instance from Decigram.
func (uf MassFactory) FromDecigrams(value float64) (*Mass, error) {
	return newMass(value, MassDecigram)
}

// FromDecagram creates a new Mass instance from Decagram.
func (uf MassFactory) FromDecagrams(value float64) (*Mass, error) {
	return newMass(value, MassDecagram)
}

// FromHectogram creates a new Mass instance from Hectogram.
func (uf MassFactory) FromHectograms(value float64) (*Mass, error) {
	return newMass(value, MassHectogram)
}

// FromKilogram creates a new Mass instance from Kilogram.
func (uf MassFactory) FromKilograms(value float64) (*Mass, error) {
	return newMass(value, MassKilogram)
}

// FromKilotonne creates a new Mass instance from Kilotonne.
func (uf MassFactory) FromKilotonnes(value float64) (*Mass, error) {
	return newMass(value, MassKilotonne)
}

// FromMegatonne creates a new Mass instance from Megatonne.
func (uf MassFactory) FromMegatonnes(value float64) (*Mass, error) {
	return newMass(value, MassMegatonne)
}

// FromKilopound creates a new Mass instance from Kilopound.
func (uf MassFactory) FromKilopounds(value float64) (*Mass, error) {
	return newMass(value, MassKilopound)
}

// FromMegapound creates a new Mass instance from Megapound.
func (uf MassFactory) FromMegapounds(value float64) (*Mass, error) {
	return newMass(value, MassMegapound)
}




// newMass creates a new Mass.
func newMass(value float64, fromUnit MassUnits) (*Mass, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Mass{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Mass in Kilogram.
func (a *Mass) BaseValue() float64 {
	return a.value
}


// Gram returns the value in Gram.
func (a *Mass) Grams() float64 {
	if a.gramsLazy != nil {
		return *a.gramsLazy
	}
	grams := a.convertFromBase(MassGram)
	a.gramsLazy = &grams
	return grams
}

// Tonne returns the value in Tonne.
func (a *Mass) Tonnes() float64 {
	if a.tonnesLazy != nil {
		return *a.tonnesLazy
	}
	tonnes := a.convertFromBase(MassTonne)
	a.tonnesLazy = &tonnes
	return tonnes
}

// ShortTon returns the value in ShortTon.
func (a *Mass) ShortTons() float64 {
	if a.short_tonsLazy != nil {
		return *a.short_tonsLazy
	}
	short_tons := a.convertFromBase(MassShortTon)
	a.short_tonsLazy = &short_tons
	return short_tons
}

// LongTon returns the value in LongTon.
func (a *Mass) LongTons() float64 {
	if a.long_tonsLazy != nil {
		return *a.long_tonsLazy
	}
	long_tons := a.convertFromBase(MassLongTon)
	a.long_tonsLazy = &long_tons
	return long_tons
}

// Pound returns the value in Pound.
func (a *Mass) Pounds() float64 {
	if a.poundsLazy != nil {
		return *a.poundsLazy
	}
	pounds := a.convertFromBase(MassPound)
	a.poundsLazy = &pounds
	return pounds
}

// Ounce returns the value in Ounce.
func (a *Mass) Ounces() float64 {
	if a.ouncesLazy != nil {
		return *a.ouncesLazy
	}
	ounces := a.convertFromBase(MassOunce)
	a.ouncesLazy = &ounces
	return ounces
}

// Slug returns the value in Slug.
func (a *Mass) Slugs() float64 {
	if a.slugsLazy != nil {
		return *a.slugsLazy
	}
	slugs := a.convertFromBase(MassSlug)
	a.slugsLazy = &slugs
	return slugs
}

// Stone returns the value in Stone.
func (a *Mass) Stone() float64 {
	if a.stoneLazy != nil {
		return *a.stoneLazy
	}
	stone := a.convertFromBase(MassStone)
	a.stoneLazy = &stone
	return stone
}

// ShortHundredweight returns the value in ShortHundredweight.
func (a *Mass) ShortHundredweight() float64 {
	if a.short_hundredweightLazy != nil {
		return *a.short_hundredweightLazy
	}
	short_hundredweight := a.convertFromBase(MassShortHundredweight)
	a.short_hundredweightLazy = &short_hundredweight
	return short_hundredweight
}

// LongHundredweight returns the value in LongHundredweight.
func (a *Mass) LongHundredweight() float64 {
	if a.long_hundredweightLazy != nil {
		return *a.long_hundredweightLazy
	}
	long_hundredweight := a.convertFromBase(MassLongHundredweight)
	a.long_hundredweightLazy = &long_hundredweight
	return long_hundredweight
}

// Grain returns the value in Grain.
func (a *Mass) Grains() float64 {
	if a.grainsLazy != nil {
		return *a.grainsLazy
	}
	grains := a.convertFromBase(MassGrain)
	a.grainsLazy = &grains
	return grains
}

// SolarMass returns the value in SolarMass.
func (a *Mass) SolarMasses() float64 {
	if a.solar_massesLazy != nil {
		return *a.solar_massesLazy
	}
	solar_masses := a.convertFromBase(MassSolarMass)
	a.solar_massesLazy = &solar_masses
	return solar_masses
}

// EarthMass returns the value in EarthMass.
func (a *Mass) EarthMasses() float64 {
	if a.earth_massesLazy != nil {
		return *a.earth_massesLazy
	}
	earth_masses := a.convertFromBase(MassEarthMass)
	a.earth_massesLazy = &earth_masses
	return earth_masses
}

// Femtogram returns the value in Femtogram.
func (a *Mass) Femtograms() float64 {
	if a.femtogramsLazy != nil {
		return *a.femtogramsLazy
	}
	femtograms := a.convertFromBase(MassFemtogram)
	a.femtogramsLazy = &femtograms
	return femtograms
}

// Picogram returns the value in Picogram.
func (a *Mass) Picograms() float64 {
	if a.picogramsLazy != nil {
		return *a.picogramsLazy
	}
	picograms := a.convertFromBase(MassPicogram)
	a.picogramsLazy = &picograms
	return picograms
}

// Nanogram returns the value in Nanogram.
func (a *Mass) Nanograms() float64 {
	if a.nanogramsLazy != nil {
		return *a.nanogramsLazy
	}
	nanograms := a.convertFromBase(MassNanogram)
	a.nanogramsLazy = &nanograms
	return nanograms
}

// Microgram returns the value in Microgram.
func (a *Mass) Micrograms() float64 {
	if a.microgramsLazy != nil {
		return *a.microgramsLazy
	}
	micrograms := a.convertFromBase(MassMicrogram)
	a.microgramsLazy = &micrograms
	return micrograms
}

// Milligram returns the value in Milligram.
func (a *Mass) Milligrams() float64 {
	if a.milligramsLazy != nil {
		return *a.milligramsLazy
	}
	milligrams := a.convertFromBase(MassMilligram)
	a.milligramsLazy = &milligrams
	return milligrams
}

// Centigram returns the value in Centigram.
func (a *Mass) Centigrams() float64 {
	if a.centigramsLazy != nil {
		return *a.centigramsLazy
	}
	centigrams := a.convertFromBase(MassCentigram)
	a.centigramsLazy = &centigrams
	return centigrams
}

// Decigram returns the value in Decigram.
func (a *Mass) Decigrams() float64 {
	if a.decigramsLazy != nil {
		return *a.decigramsLazy
	}
	decigrams := a.convertFromBase(MassDecigram)
	a.decigramsLazy = &decigrams
	return decigrams
}

// Decagram returns the value in Decagram.
func (a *Mass) Decagrams() float64 {
	if a.decagramsLazy != nil {
		return *a.decagramsLazy
	}
	decagrams := a.convertFromBase(MassDecagram)
	a.decagramsLazy = &decagrams
	return decagrams
}

// Hectogram returns the value in Hectogram.
func (a *Mass) Hectograms() float64 {
	if a.hectogramsLazy != nil {
		return *a.hectogramsLazy
	}
	hectograms := a.convertFromBase(MassHectogram)
	a.hectogramsLazy = &hectograms
	return hectograms
}

// Kilogram returns the value in Kilogram.
func (a *Mass) Kilograms() float64 {
	if a.kilogramsLazy != nil {
		return *a.kilogramsLazy
	}
	kilograms := a.convertFromBase(MassKilogram)
	a.kilogramsLazy = &kilograms
	return kilograms
}

// Kilotonne returns the value in Kilotonne.
func (a *Mass) Kilotonnes() float64 {
	if a.kilotonnesLazy != nil {
		return *a.kilotonnesLazy
	}
	kilotonnes := a.convertFromBase(MassKilotonne)
	a.kilotonnesLazy = &kilotonnes
	return kilotonnes
}

// Megatonne returns the value in Megatonne.
func (a *Mass) Megatonnes() float64 {
	if a.megatonnesLazy != nil {
		return *a.megatonnesLazy
	}
	megatonnes := a.convertFromBase(MassMegatonne)
	a.megatonnesLazy = &megatonnes
	return megatonnes
}

// Kilopound returns the value in Kilopound.
func (a *Mass) Kilopounds() float64 {
	if a.kilopoundsLazy != nil {
		return *a.kilopoundsLazy
	}
	kilopounds := a.convertFromBase(MassKilopound)
	a.kilopoundsLazy = &kilopounds
	return kilopounds
}

// Megapound returns the value in Megapound.
func (a *Mass) Megapounds() float64 {
	if a.megapoundsLazy != nil {
		return *a.megapoundsLazy
	}
	megapounds := a.convertFromBase(MassMegapound)
	a.megapoundsLazy = &megapounds
	return megapounds
}


// ToDto creates an MassDto representation.
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

// ToDtoJSON creates an MassDto representation.
func (a *Mass) ToDtoJSON(holdInUnit *MassUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Mass to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Mass) String() string {
	return a.ToString(MassKilogram, 2)
}

// ToString formats the Mass to string.
// fractionalDigits -1 for not mention
func (a *Mass) ToString(unit MassUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Mass) getUnitAbbreviation(unit MassUnits) string {
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

// Check if the given Mass are equals to the current Mass
func (a *Mass) Equals(other *Mass) bool {
	return a.value == other.BaseValue()
}

// Check if the given Mass are equals to the current Mass
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

// Add the given Mass to the current Mass.
func (a *Mass) Add(other *Mass) *Mass {
	return &Mass{value: a.value + other.BaseValue()}
}

// Subtract the given Mass to the current Mass.
func (a *Mass) Subtract(other *Mass) *Mass {
	return &Mass{value: a.value - other.BaseValue()}
}

// Multiply the given Mass to the current Mass.
func (a *Mass) Multiply(other *Mass) *Mass {
	return &Mass{value: a.value * other.BaseValue()}
}

// Divide the given Mass to the current Mass.
func (a *Mass) Divide(other *Mass) *Mass {
	return &Mass{value: a.value / other.BaseValue()}
}