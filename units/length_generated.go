// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LengthUnits defines various units of Length.
type LengthUnits string

const (
    
        // 
        LengthMeter LengthUnits = "Meter"
        // The statute mile was standardised between the British Commonwealth and the United States by an international agreement in 1959, when it was formally redefined with respect to SI units as exactly 1,609.344 metres.
        LengthMile LengthUnits = "Mile"
        // The yard (symbol: yd) is an English unit of length in both the British imperial and US customary systems of measurement equalling 3 feet (or 36 inches). Since 1959 the yard has been by international agreement standardized as exactly 0.9144 meter. A distance of 1,760 yards is equal to 1 mile.
        LengthYard LengthUnits = "Yard"
        // 
        LengthFoot LengthUnits = "Foot"
        // In the United States, the foot was defined as 12 inches, with the inch being defined by the Mendenhall Order of 1893 as 39.37 inches = 1 m. This makes a U.S. survey foot exactly 1200/3937 meters.
        LengthUsSurveyFoot LengthUnits = "UsSurveyFoot"
        // 
        LengthInch LengthUnits = "Inch"
        // 
        LengthMil LengthUnits = "Mil"
        // 
        LengthNauticalMile LengthUnits = "NauticalMile"
        // 
        LengthFathom LengthUnits = "Fathom"
        // 
        LengthShackle LengthUnits = "Shackle"
        // 
        LengthMicroinch LengthUnits = "Microinch"
        // 
        LengthPrinterPoint LengthUnits = "PrinterPoint"
        // 
        LengthDtpPoint LengthUnits = "DtpPoint"
        // 
        LengthPrinterPica LengthUnits = "PrinterPica"
        // 
        LengthDtpPica LengthUnits = "DtpPica"
        // 
        LengthTwip LengthUnits = "Twip"
        // 
        LengthHand LengthUnits = "Hand"
        // One Astronomical Unit is the distance from the solar system Star, the sun, to planet Earth.
        LengthAstronomicalUnit LengthUnits = "AstronomicalUnit"
        // A parsec is defined as the distance at which one astronomical unit (AU) subtends an angle of one arcsecond.
        LengthParsec LengthUnits = "Parsec"
        // A Light Year (ly) is the distance that light travel during an Earth year, ie 365 days.
        LengthLightYear LengthUnits = "LightYear"
        // Solar radius is a ratio unit to the radius of the solar system star, the sun.
        LengthSolarRadius LengthUnits = "SolarRadius"
        // 
        LengthChain LengthUnits = "Chain"
        // Angstrom is a metric unit of length equal to 1e-10 meter
        LengthAngstrom LengthUnits = "Angstrom"
        // In radar-related subjects and in JTIDS, a data mile is a unit of distance equal to 6000 feet (1.8288 kilometres or 0.987 nautical miles).
        LengthDataMile LengthUnits = "DataMile"
        // 
        LengthFemtometer LengthUnits = "Femtometer"
        // 
        LengthPicometer LengthUnits = "Picometer"
        // 
        LengthNanometer LengthUnits = "Nanometer"
        // 
        LengthMicrometer LengthUnits = "Micrometer"
        // 
        LengthMillimeter LengthUnits = "Millimeter"
        // 
        LengthCentimeter LengthUnits = "Centimeter"
        // 
        LengthDecimeter LengthUnits = "Decimeter"
        // 
        LengthDecameter LengthUnits = "Decameter"
        // 
        LengthHectometer LengthUnits = "Hectometer"
        // 
        LengthKilometer LengthUnits = "Kilometer"
        // 
        LengthMegameter LengthUnits = "Megameter"
        // 
        LengthGigameter LengthUnits = "Gigameter"
        // 
        LengthKiloyard LengthUnits = "Kiloyard"
        // 
        LengthKilofoot LengthUnits = "Kilofoot"
        // 
        LengthKiloparsec LengthUnits = "Kiloparsec"
        // 
        LengthMegaparsec LengthUnits = "Megaparsec"
        // 
        LengthKilolightYear LengthUnits = "KilolightYear"
        // 
        LengthMegalightYear LengthUnits = "MegalightYear"
)

// LengthDto represents a Length measurement with a numerical value and its corresponding unit.
type LengthDto struct {
    // Value is the numerical representation of the Length.
	Value float64
    // Unit specifies the unit of measurement for the Length, as defined in the LengthUnits enumeration.
	Unit  LengthUnits
}

// LengthDtoFactory groups methods for creating and serializing LengthDto objects.
type LengthDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a LengthDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf LengthDtoFactory) FromJSON(data []byte) (*LengthDto, error) {
	a := LengthDto{}

    // Parse JSON into LengthDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

// ToJSON serializes a LengthDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a LengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}


// Length represents a measurement in a of Length.
//
// Many different units of length have been used around the world. The main units in modern use are U.S. customary units in the United States and the Metric system elsewhere. British Imperial units are still used for some purposes in the United Kingdom and some other countries. The metric system is sub-divided into SI and non-SI units.
type Length struct {
	// value is the base measurement stored internally.
	value       float64
    
    metersLazy *float64 
    milesLazy *float64 
    yardsLazy *float64 
    feetLazy *float64 
    us_survey_feetLazy *float64 
    inchesLazy *float64 
    milsLazy *float64 
    nautical_milesLazy *float64 
    fathomsLazy *float64 
    shacklesLazy *float64 
    microinchesLazy *float64 
    printer_pointsLazy *float64 
    dtp_pointsLazy *float64 
    printer_picasLazy *float64 
    dtp_picasLazy *float64 
    twipsLazy *float64 
    handsLazy *float64 
    astronomical_unitsLazy *float64 
    parsecsLazy *float64 
    light_yearsLazy *float64 
    solar_radiusesLazy *float64 
    chainsLazy *float64 
    angstromsLazy *float64 
    data_milesLazy *float64 
    femtometersLazy *float64 
    picometersLazy *float64 
    nanometersLazy *float64 
    micrometersLazy *float64 
    millimetersLazy *float64 
    centimetersLazy *float64 
    decimetersLazy *float64 
    decametersLazy *float64 
    hectometersLazy *float64 
    kilometersLazy *float64 
    megametersLazy *float64 
    gigametersLazy *float64 
    kiloyardsLazy *float64 
    kilofeetLazy *float64 
    kiloparsecsLazy *float64 
    megaparsecsLazy *float64 
    kilolight_yearsLazy *float64 
    megalight_yearsLazy *float64 
}

// LengthFactory groups methods for creating Length instances.
type LengthFactory struct{}

// CreateLength creates a new Length instance from the given value and unit.
func (uf LengthFactory) CreateLength(value float64, unit LengthUnits) (*Length, error) {
	return newLength(value, unit)
}

// FromDto converts a LengthDto to a Length instance.
func (uf LengthFactory) FromDto(dto LengthDto) (*Length, error) {
	return newLength(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Length instance.
func (uf LengthFactory) FromDtoJSON(data []byte) (*Length, error) {
	unitDto, err := LengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LengthDto from JSON: %w", err)
	}
	return LengthFactory{}.FromDto(*unitDto)
}


// FromMeters creates a new Length instance from a value in Meters.
func (uf LengthFactory) FromMeters(value float64) (*Length, error) {
	return newLength(value, LengthMeter)
}

// FromMiles creates a new Length instance from a value in Miles.
func (uf LengthFactory) FromMiles(value float64) (*Length, error) {
	return newLength(value, LengthMile)
}

// FromYards creates a new Length instance from a value in Yards.
func (uf LengthFactory) FromYards(value float64) (*Length, error) {
	return newLength(value, LengthYard)
}

// FromFeet creates a new Length instance from a value in Feet.
func (uf LengthFactory) FromFeet(value float64) (*Length, error) {
	return newLength(value, LengthFoot)
}

// FromUsSurveyFeet creates a new Length instance from a value in UsSurveyFeet.
func (uf LengthFactory) FromUsSurveyFeet(value float64) (*Length, error) {
	return newLength(value, LengthUsSurveyFoot)
}

// FromInches creates a new Length instance from a value in Inches.
func (uf LengthFactory) FromInches(value float64) (*Length, error) {
	return newLength(value, LengthInch)
}

// FromMils creates a new Length instance from a value in Mils.
func (uf LengthFactory) FromMils(value float64) (*Length, error) {
	return newLength(value, LengthMil)
}

// FromNauticalMiles creates a new Length instance from a value in NauticalMiles.
func (uf LengthFactory) FromNauticalMiles(value float64) (*Length, error) {
	return newLength(value, LengthNauticalMile)
}

// FromFathoms creates a new Length instance from a value in Fathoms.
func (uf LengthFactory) FromFathoms(value float64) (*Length, error) {
	return newLength(value, LengthFathom)
}

// FromShackles creates a new Length instance from a value in Shackles.
func (uf LengthFactory) FromShackles(value float64) (*Length, error) {
	return newLength(value, LengthShackle)
}

// FromMicroinches creates a new Length instance from a value in Microinches.
func (uf LengthFactory) FromMicroinches(value float64) (*Length, error) {
	return newLength(value, LengthMicroinch)
}

// FromPrinterPoints creates a new Length instance from a value in PrinterPoints.
func (uf LengthFactory) FromPrinterPoints(value float64) (*Length, error) {
	return newLength(value, LengthPrinterPoint)
}

// FromDtpPoints creates a new Length instance from a value in DtpPoints.
func (uf LengthFactory) FromDtpPoints(value float64) (*Length, error) {
	return newLength(value, LengthDtpPoint)
}

// FromPrinterPicas creates a new Length instance from a value in PrinterPicas.
func (uf LengthFactory) FromPrinterPicas(value float64) (*Length, error) {
	return newLength(value, LengthPrinterPica)
}

// FromDtpPicas creates a new Length instance from a value in DtpPicas.
func (uf LengthFactory) FromDtpPicas(value float64) (*Length, error) {
	return newLength(value, LengthDtpPica)
}

// FromTwips creates a new Length instance from a value in Twips.
func (uf LengthFactory) FromTwips(value float64) (*Length, error) {
	return newLength(value, LengthTwip)
}

// FromHands creates a new Length instance from a value in Hands.
func (uf LengthFactory) FromHands(value float64) (*Length, error) {
	return newLength(value, LengthHand)
}

// FromAstronomicalUnits creates a new Length instance from a value in AstronomicalUnits.
func (uf LengthFactory) FromAstronomicalUnits(value float64) (*Length, error) {
	return newLength(value, LengthAstronomicalUnit)
}

// FromParsecs creates a new Length instance from a value in Parsecs.
func (uf LengthFactory) FromParsecs(value float64) (*Length, error) {
	return newLength(value, LengthParsec)
}

// FromLightYears creates a new Length instance from a value in LightYears.
func (uf LengthFactory) FromLightYears(value float64) (*Length, error) {
	return newLength(value, LengthLightYear)
}

// FromSolarRadiuses creates a new Length instance from a value in SolarRadiuses.
func (uf LengthFactory) FromSolarRadiuses(value float64) (*Length, error) {
	return newLength(value, LengthSolarRadius)
}

// FromChains creates a new Length instance from a value in Chains.
func (uf LengthFactory) FromChains(value float64) (*Length, error) {
	return newLength(value, LengthChain)
}

// FromAngstroms creates a new Length instance from a value in Angstroms.
func (uf LengthFactory) FromAngstroms(value float64) (*Length, error) {
	return newLength(value, LengthAngstrom)
}

// FromDataMiles creates a new Length instance from a value in DataMiles.
func (uf LengthFactory) FromDataMiles(value float64) (*Length, error) {
	return newLength(value, LengthDataMile)
}

// FromFemtometers creates a new Length instance from a value in Femtometers.
func (uf LengthFactory) FromFemtometers(value float64) (*Length, error) {
	return newLength(value, LengthFemtometer)
}

// FromPicometers creates a new Length instance from a value in Picometers.
func (uf LengthFactory) FromPicometers(value float64) (*Length, error) {
	return newLength(value, LengthPicometer)
}

// FromNanometers creates a new Length instance from a value in Nanometers.
func (uf LengthFactory) FromNanometers(value float64) (*Length, error) {
	return newLength(value, LengthNanometer)
}

// FromMicrometers creates a new Length instance from a value in Micrometers.
func (uf LengthFactory) FromMicrometers(value float64) (*Length, error) {
	return newLength(value, LengthMicrometer)
}

// FromMillimeters creates a new Length instance from a value in Millimeters.
func (uf LengthFactory) FromMillimeters(value float64) (*Length, error) {
	return newLength(value, LengthMillimeter)
}

// FromCentimeters creates a new Length instance from a value in Centimeters.
func (uf LengthFactory) FromCentimeters(value float64) (*Length, error) {
	return newLength(value, LengthCentimeter)
}

// FromDecimeters creates a new Length instance from a value in Decimeters.
func (uf LengthFactory) FromDecimeters(value float64) (*Length, error) {
	return newLength(value, LengthDecimeter)
}

// FromDecameters creates a new Length instance from a value in Decameters.
func (uf LengthFactory) FromDecameters(value float64) (*Length, error) {
	return newLength(value, LengthDecameter)
}

// FromHectometers creates a new Length instance from a value in Hectometers.
func (uf LengthFactory) FromHectometers(value float64) (*Length, error) {
	return newLength(value, LengthHectometer)
}

// FromKilometers creates a new Length instance from a value in Kilometers.
func (uf LengthFactory) FromKilometers(value float64) (*Length, error) {
	return newLength(value, LengthKilometer)
}

// FromMegameters creates a new Length instance from a value in Megameters.
func (uf LengthFactory) FromMegameters(value float64) (*Length, error) {
	return newLength(value, LengthMegameter)
}

// FromGigameters creates a new Length instance from a value in Gigameters.
func (uf LengthFactory) FromGigameters(value float64) (*Length, error) {
	return newLength(value, LengthGigameter)
}

// FromKiloyards creates a new Length instance from a value in Kiloyards.
func (uf LengthFactory) FromKiloyards(value float64) (*Length, error) {
	return newLength(value, LengthKiloyard)
}

// FromKilofeet creates a new Length instance from a value in Kilofeet.
func (uf LengthFactory) FromKilofeet(value float64) (*Length, error) {
	return newLength(value, LengthKilofoot)
}

// FromKiloparsecs creates a new Length instance from a value in Kiloparsecs.
func (uf LengthFactory) FromKiloparsecs(value float64) (*Length, error) {
	return newLength(value, LengthKiloparsec)
}

// FromMegaparsecs creates a new Length instance from a value in Megaparsecs.
func (uf LengthFactory) FromMegaparsecs(value float64) (*Length, error) {
	return newLength(value, LengthMegaparsec)
}

// FromKilolightYears creates a new Length instance from a value in KilolightYears.
func (uf LengthFactory) FromKilolightYears(value float64) (*Length, error) {
	return newLength(value, LengthKilolightYear)
}

// FromMegalightYears creates a new Length instance from a value in MegalightYears.
func (uf LengthFactory) FromMegalightYears(value float64) (*Length, error) {
	return newLength(value, LengthMegalightYear)
}


// newLength creates a new Length.
func newLength(value float64, fromUnit LengthUnits) (*Length, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Length{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Length in Meter unit (the base/default quantity).
func (a *Length) BaseValue() float64 {
	return a.value
}


// Meters returns the Length value in Meters.
//
// 
func (a *Length) Meters() float64 {
	if a.metersLazy != nil {
		return *a.metersLazy
	}
	meters := a.convertFromBase(LengthMeter)
	a.metersLazy = &meters
	return meters
}

// Miles returns the Length value in Miles.
//
// The statute mile was standardised between the British Commonwealth and the United States by an international agreement in 1959, when it was formally redefined with respect to SI units as exactly 1,609.344 metres.
func (a *Length) Miles() float64 {
	if a.milesLazy != nil {
		return *a.milesLazy
	}
	miles := a.convertFromBase(LengthMile)
	a.milesLazy = &miles
	return miles
}

// Yards returns the Length value in Yards.
//
// The yard (symbol: yd) is an English unit of length in both the British imperial and US customary systems of measurement equalling 3 feet (or 36 inches). Since 1959 the yard has been by international agreement standardized as exactly 0.9144 meter. A distance of 1,760 yards is equal to 1 mile.
func (a *Length) Yards() float64 {
	if a.yardsLazy != nil {
		return *a.yardsLazy
	}
	yards := a.convertFromBase(LengthYard)
	a.yardsLazy = &yards
	return yards
}

// Feet returns the Length value in Feet.
//
// 
func (a *Length) Feet() float64 {
	if a.feetLazy != nil {
		return *a.feetLazy
	}
	feet := a.convertFromBase(LengthFoot)
	a.feetLazy = &feet
	return feet
}

// UsSurveyFeet returns the Length value in UsSurveyFeet.
//
// In the United States, the foot was defined as 12 inches, with the inch being defined by the Mendenhall Order of 1893 as 39.37 inches = 1 m. This makes a U.S. survey foot exactly 1200/3937 meters.
func (a *Length) UsSurveyFeet() float64 {
	if a.us_survey_feetLazy != nil {
		return *a.us_survey_feetLazy
	}
	us_survey_feet := a.convertFromBase(LengthUsSurveyFoot)
	a.us_survey_feetLazy = &us_survey_feet
	return us_survey_feet
}

// Inches returns the Length value in Inches.
//
// 
func (a *Length) Inches() float64 {
	if a.inchesLazy != nil {
		return *a.inchesLazy
	}
	inches := a.convertFromBase(LengthInch)
	a.inchesLazy = &inches
	return inches
}

// Mils returns the Length value in Mils.
//
// 
func (a *Length) Mils() float64 {
	if a.milsLazy != nil {
		return *a.milsLazy
	}
	mils := a.convertFromBase(LengthMil)
	a.milsLazy = &mils
	return mils
}

// NauticalMiles returns the Length value in NauticalMiles.
//
// 
func (a *Length) NauticalMiles() float64 {
	if a.nautical_milesLazy != nil {
		return *a.nautical_milesLazy
	}
	nautical_miles := a.convertFromBase(LengthNauticalMile)
	a.nautical_milesLazy = &nautical_miles
	return nautical_miles
}

// Fathoms returns the Length value in Fathoms.
//
// 
func (a *Length) Fathoms() float64 {
	if a.fathomsLazy != nil {
		return *a.fathomsLazy
	}
	fathoms := a.convertFromBase(LengthFathom)
	a.fathomsLazy = &fathoms
	return fathoms
}

// Shackles returns the Length value in Shackles.
//
// 
func (a *Length) Shackles() float64 {
	if a.shacklesLazy != nil {
		return *a.shacklesLazy
	}
	shackles := a.convertFromBase(LengthShackle)
	a.shacklesLazy = &shackles
	return shackles
}

// Microinches returns the Length value in Microinches.
//
// 
func (a *Length) Microinches() float64 {
	if a.microinchesLazy != nil {
		return *a.microinchesLazy
	}
	microinches := a.convertFromBase(LengthMicroinch)
	a.microinchesLazy = &microinches
	return microinches
}

// PrinterPoints returns the Length value in PrinterPoints.
//
// 
func (a *Length) PrinterPoints() float64 {
	if a.printer_pointsLazy != nil {
		return *a.printer_pointsLazy
	}
	printer_points := a.convertFromBase(LengthPrinterPoint)
	a.printer_pointsLazy = &printer_points
	return printer_points
}

// DtpPoints returns the Length value in DtpPoints.
//
// 
func (a *Length) DtpPoints() float64 {
	if a.dtp_pointsLazy != nil {
		return *a.dtp_pointsLazy
	}
	dtp_points := a.convertFromBase(LengthDtpPoint)
	a.dtp_pointsLazy = &dtp_points
	return dtp_points
}

// PrinterPicas returns the Length value in PrinterPicas.
//
// 
func (a *Length) PrinterPicas() float64 {
	if a.printer_picasLazy != nil {
		return *a.printer_picasLazy
	}
	printer_picas := a.convertFromBase(LengthPrinterPica)
	a.printer_picasLazy = &printer_picas
	return printer_picas
}

// DtpPicas returns the Length value in DtpPicas.
//
// 
func (a *Length) DtpPicas() float64 {
	if a.dtp_picasLazy != nil {
		return *a.dtp_picasLazy
	}
	dtp_picas := a.convertFromBase(LengthDtpPica)
	a.dtp_picasLazy = &dtp_picas
	return dtp_picas
}

// Twips returns the Length value in Twips.
//
// 
func (a *Length) Twips() float64 {
	if a.twipsLazy != nil {
		return *a.twipsLazy
	}
	twips := a.convertFromBase(LengthTwip)
	a.twipsLazy = &twips
	return twips
}

// Hands returns the Length value in Hands.
//
// 
func (a *Length) Hands() float64 {
	if a.handsLazy != nil {
		return *a.handsLazy
	}
	hands := a.convertFromBase(LengthHand)
	a.handsLazy = &hands
	return hands
}

// AstronomicalUnits returns the Length value in AstronomicalUnits.
//
// One Astronomical Unit is the distance from the solar system Star, the sun, to planet Earth.
func (a *Length) AstronomicalUnits() float64 {
	if a.astronomical_unitsLazy != nil {
		return *a.astronomical_unitsLazy
	}
	astronomical_units := a.convertFromBase(LengthAstronomicalUnit)
	a.astronomical_unitsLazy = &astronomical_units
	return astronomical_units
}

// Parsecs returns the Length value in Parsecs.
//
// A parsec is defined as the distance at which one astronomical unit (AU) subtends an angle of one arcsecond.
func (a *Length) Parsecs() float64 {
	if a.parsecsLazy != nil {
		return *a.parsecsLazy
	}
	parsecs := a.convertFromBase(LengthParsec)
	a.parsecsLazy = &parsecs
	return parsecs
}

// LightYears returns the Length value in LightYears.
//
// A Light Year (ly) is the distance that light travel during an Earth year, ie 365 days.
func (a *Length) LightYears() float64 {
	if a.light_yearsLazy != nil {
		return *a.light_yearsLazy
	}
	light_years := a.convertFromBase(LengthLightYear)
	a.light_yearsLazy = &light_years
	return light_years
}

// SolarRadiuses returns the Length value in SolarRadiuses.
//
// Solar radius is a ratio unit to the radius of the solar system star, the sun.
func (a *Length) SolarRadiuses() float64 {
	if a.solar_radiusesLazy != nil {
		return *a.solar_radiusesLazy
	}
	solar_radiuses := a.convertFromBase(LengthSolarRadius)
	a.solar_radiusesLazy = &solar_radiuses
	return solar_radiuses
}

// Chains returns the Length value in Chains.
//
// 
func (a *Length) Chains() float64 {
	if a.chainsLazy != nil {
		return *a.chainsLazy
	}
	chains := a.convertFromBase(LengthChain)
	a.chainsLazy = &chains
	return chains
}

// Angstroms returns the Length value in Angstroms.
//
// Angstrom is a metric unit of length equal to 1e-10 meter
func (a *Length) Angstroms() float64 {
	if a.angstromsLazy != nil {
		return *a.angstromsLazy
	}
	angstroms := a.convertFromBase(LengthAngstrom)
	a.angstromsLazy = &angstroms
	return angstroms
}

// DataMiles returns the Length value in DataMiles.
//
// In radar-related subjects and in JTIDS, a data mile is a unit of distance equal to 6000 feet (1.8288 kilometres or 0.987 nautical miles).
func (a *Length) DataMiles() float64 {
	if a.data_milesLazy != nil {
		return *a.data_milesLazy
	}
	data_miles := a.convertFromBase(LengthDataMile)
	a.data_milesLazy = &data_miles
	return data_miles
}

// Femtometers returns the Length value in Femtometers.
//
// 
func (a *Length) Femtometers() float64 {
	if a.femtometersLazy != nil {
		return *a.femtometersLazy
	}
	femtometers := a.convertFromBase(LengthFemtometer)
	a.femtometersLazy = &femtometers
	return femtometers
}

// Picometers returns the Length value in Picometers.
//
// 
func (a *Length) Picometers() float64 {
	if a.picometersLazy != nil {
		return *a.picometersLazy
	}
	picometers := a.convertFromBase(LengthPicometer)
	a.picometersLazy = &picometers
	return picometers
}

// Nanometers returns the Length value in Nanometers.
//
// 
func (a *Length) Nanometers() float64 {
	if a.nanometersLazy != nil {
		return *a.nanometersLazy
	}
	nanometers := a.convertFromBase(LengthNanometer)
	a.nanometersLazy = &nanometers
	return nanometers
}

// Micrometers returns the Length value in Micrometers.
//
// 
func (a *Length) Micrometers() float64 {
	if a.micrometersLazy != nil {
		return *a.micrometersLazy
	}
	micrometers := a.convertFromBase(LengthMicrometer)
	a.micrometersLazy = &micrometers
	return micrometers
}

// Millimeters returns the Length value in Millimeters.
//
// 
func (a *Length) Millimeters() float64 {
	if a.millimetersLazy != nil {
		return *a.millimetersLazy
	}
	millimeters := a.convertFromBase(LengthMillimeter)
	a.millimetersLazy = &millimeters
	return millimeters
}

// Centimeters returns the Length value in Centimeters.
//
// 
func (a *Length) Centimeters() float64 {
	if a.centimetersLazy != nil {
		return *a.centimetersLazy
	}
	centimeters := a.convertFromBase(LengthCentimeter)
	a.centimetersLazy = &centimeters
	return centimeters
}

// Decimeters returns the Length value in Decimeters.
//
// 
func (a *Length) Decimeters() float64 {
	if a.decimetersLazy != nil {
		return *a.decimetersLazy
	}
	decimeters := a.convertFromBase(LengthDecimeter)
	a.decimetersLazy = &decimeters
	return decimeters
}

// Decameters returns the Length value in Decameters.
//
// 
func (a *Length) Decameters() float64 {
	if a.decametersLazy != nil {
		return *a.decametersLazy
	}
	decameters := a.convertFromBase(LengthDecameter)
	a.decametersLazy = &decameters
	return decameters
}

// Hectometers returns the Length value in Hectometers.
//
// 
func (a *Length) Hectometers() float64 {
	if a.hectometersLazy != nil {
		return *a.hectometersLazy
	}
	hectometers := a.convertFromBase(LengthHectometer)
	a.hectometersLazy = &hectometers
	return hectometers
}

// Kilometers returns the Length value in Kilometers.
//
// 
func (a *Length) Kilometers() float64 {
	if a.kilometersLazy != nil {
		return *a.kilometersLazy
	}
	kilometers := a.convertFromBase(LengthKilometer)
	a.kilometersLazy = &kilometers
	return kilometers
}

// Megameters returns the Length value in Megameters.
//
// 
func (a *Length) Megameters() float64 {
	if a.megametersLazy != nil {
		return *a.megametersLazy
	}
	megameters := a.convertFromBase(LengthMegameter)
	a.megametersLazy = &megameters
	return megameters
}

// Gigameters returns the Length value in Gigameters.
//
// 
func (a *Length) Gigameters() float64 {
	if a.gigametersLazy != nil {
		return *a.gigametersLazy
	}
	gigameters := a.convertFromBase(LengthGigameter)
	a.gigametersLazy = &gigameters
	return gigameters
}

// Kiloyards returns the Length value in Kiloyards.
//
// 
func (a *Length) Kiloyards() float64 {
	if a.kiloyardsLazy != nil {
		return *a.kiloyardsLazy
	}
	kiloyards := a.convertFromBase(LengthKiloyard)
	a.kiloyardsLazy = &kiloyards
	return kiloyards
}

// Kilofeet returns the Length value in Kilofeet.
//
// 
func (a *Length) Kilofeet() float64 {
	if a.kilofeetLazy != nil {
		return *a.kilofeetLazy
	}
	kilofeet := a.convertFromBase(LengthKilofoot)
	a.kilofeetLazy = &kilofeet
	return kilofeet
}

// Kiloparsecs returns the Length value in Kiloparsecs.
//
// 
func (a *Length) Kiloparsecs() float64 {
	if a.kiloparsecsLazy != nil {
		return *a.kiloparsecsLazy
	}
	kiloparsecs := a.convertFromBase(LengthKiloparsec)
	a.kiloparsecsLazy = &kiloparsecs
	return kiloparsecs
}

// Megaparsecs returns the Length value in Megaparsecs.
//
// 
func (a *Length) Megaparsecs() float64 {
	if a.megaparsecsLazy != nil {
		return *a.megaparsecsLazy
	}
	megaparsecs := a.convertFromBase(LengthMegaparsec)
	a.megaparsecsLazy = &megaparsecs
	return megaparsecs
}

// KilolightYears returns the Length value in KilolightYears.
//
// 
func (a *Length) KilolightYears() float64 {
	if a.kilolight_yearsLazy != nil {
		return *a.kilolight_yearsLazy
	}
	kilolight_years := a.convertFromBase(LengthKilolightYear)
	a.kilolight_yearsLazy = &kilolight_years
	return kilolight_years
}

// MegalightYears returns the Length value in MegalightYears.
//
// 
func (a *Length) MegalightYears() float64 {
	if a.megalight_yearsLazy != nil {
		return *a.megalight_yearsLazy
	}
	megalight_years := a.convertFromBase(LengthMegalightYear)
	a.megalight_yearsLazy = &megalight_years
	return megalight_years
}


// ToDto creates a LengthDto representation from the Length instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Meter by default.
func (a *Length) ToDto(holdInUnit *LengthUnits) LengthDto {
	if holdInUnit == nil {
		defaultUnit := LengthMeter // Default value
		holdInUnit = &defaultUnit
	}

	return LengthDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Length instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Meter by default.
func (a *Length) ToDtoJSON(holdInUnit *LengthUnits) ([]byte, error) {
	// Convert to LengthDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Length to a specific unit value.
// The function uses the provided unit type (LengthUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Length) Convert(toUnit LengthUnits) float64 {
	switch toUnit { 
    case LengthMeter:
		return a.Meters()
    case LengthMile:
		return a.Miles()
    case LengthYard:
		return a.Yards()
    case LengthFoot:
		return a.Feet()
    case LengthUsSurveyFoot:
		return a.UsSurveyFeet()
    case LengthInch:
		return a.Inches()
    case LengthMil:
		return a.Mils()
    case LengthNauticalMile:
		return a.NauticalMiles()
    case LengthFathom:
		return a.Fathoms()
    case LengthShackle:
		return a.Shackles()
    case LengthMicroinch:
		return a.Microinches()
    case LengthPrinterPoint:
		return a.PrinterPoints()
    case LengthDtpPoint:
		return a.DtpPoints()
    case LengthPrinterPica:
		return a.PrinterPicas()
    case LengthDtpPica:
		return a.DtpPicas()
    case LengthTwip:
		return a.Twips()
    case LengthHand:
		return a.Hands()
    case LengthAstronomicalUnit:
		return a.AstronomicalUnits()
    case LengthParsec:
		return a.Parsecs()
    case LengthLightYear:
		return a.LightYears()
    case LengthSolarRadius:
		return a.SolarRadiuses()
    case LengthChain:
		return a.Chains()
    case LengthAngstrom:
		return a.Angstroms()
    case LengthDataMile:
		return a.DataMiles()
    case LengthFemtometer:
		return a.Femtometers()
    case LengthPicometer:
		return a.Picometers()
    case LengthNanometer:
		return a.Nanometers()
    case LengthMicrometer:
		return a.Micrometers()
    case LengthMillimeter:
		return a.Millimeters()
    case LengthCentimeter:
		return a.Centimeters()
    case LengthDecimeter:
		return a.Decimeters()
    case LengthDecameter:
		return a.Decameters()
    case LengthHectometer:
		return a.Hectometers()
    case LengthKilometer:
		return a.Kilometers()
    case LengthMegameter:
		return a.Megameters()
    case LengthGigameter:
		return a.Gigameters()
    case LengthKiloyard:
		return a.Kiloyards()
    case LengthKilofoot:
		return a.Kilofeet()
    case LengthKiloparsec:
		return a.Kiloparsecs()
    case LengthMegaparsec:
		return a.Megaparsecs()
    case LengthKilolightYear:
		return a.KilolightYears()
    case LengthMegalightYear:
		return a.MegalightYears()
	default:
		return math.NaN()
	}
}

func (a *Length) convertFromBase(toUnit LengthUnits) float64 {
    value := a.value
	switch toUnit { 
	case LengthMeter:
		return (value) 
	case LengthMile:
		return (value / 1609.344) 
	case LengthYard:
		return (value / 0.9144) 
	case LengthFoot:
		return (value / 0.3048) 
	case LengthUsSurveyFoot:
		return (value * 3937 / 1200) 
	case LengthInch:
		return (value / 2.54e-2) 
	case LengthMil:
		return (value / 2.54e-5) 
	case LengthNauticalMile:
		return (value / 1852) 
	case LengthFathom:
		return (value / 1.8288) 
	case LengthShackle:
		return (value / 27.432) 
	case LengthMicroinch:
		return (value / 2.54e-8) 
	case LengthPrinterPoint:
		return ((value / 2.54e-2) * 72.27) 
	case LengthDtpPoint:
		return ((value / 2.54e-2) * 72) 
	case LengthPrinterPica:
		return (value * 237.106301584) 
	case LengthDtpPica:
		return (value * 236.220472441) 
	case LengthTwip:
		return (value * 56692.913385826) 
	case LengthHand:
		return (value / 1.016e-1) 
	case LengthAstronomicalUnit:
		return (value / 1.4959787070e11) 
	case LengthParsec:
		return (value / 3.08567758128e16) 
	case LengthLightYear:
		return (value / 9.46073047258e15) 
	case LengthSolarRadius:
		return (value / 6.95700e8) 
	case LengthChain:
		return (value / 20.1168) 
	case LengthAngstrom:
		return (value / 1e-10) 
	case LengthDataMile:
		return (value / 1828.8) 
	case LengthFemtometer:
		return ((value) / 1e-15) 
	case LengthPicometer:
		return ((value) / 1e-12) 
	case LengthNanometer:
		return ((value) / 1e-09) 
	case LengthMicrometer:
		return ((value) / 1e-06) 
	case LengthMillimeter:
		return ((value) / 0.001) 
	case LengthCentimeter:
		return ((value) / 0.01) 
	case LengthDecimeter:
		return ((value) / 0.1) 
	case LengthDecameter:
		return ((value) / 10.0) 
	case LengthHectometer:
		return ((value) / 100.0) 
	case LengthKilometer:
		return ((value) / 1000.0) 
	case LengthMegameter:
		return ((value) / 1000000.0) 
	case LengthGigameter:
		return ((value) / 1000000000.0) 
	case LengthKiloyard:
		return ((value / 0.9144) / 1000.0) 
	case LengthKilofoot:
		return ((value / 0.3048) / 1000.0) 
	case LengthKiloparsec:
		return ((value / 3.08567758128e16) / 1000.0) 
	case LengthMegaparsec:
		return ((value / 3.08567758128e16) / 1000000.0) 
	case LengthKilolightYear:
		return ((value / 9.46073047258e15) / 1000.0) 
	case LengthMegalightYear:
		return ((value / 9.46073047258e15) / 1000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Length) convertToBase(value float64, fromUnit LengthUnits) float64 {
	switch fromUnit { 
	case LengthMeter:
		return (value) 
	case LengthMile:
		return (value * 1609.344) 
	case LengthYard:
		return (value * 0.9144) 
	case LengthFoot:
		return (value * 0.3048) 
	case LengthUsSurveyFoot:
		return (value * 1200 / 3937) 
	case LengthInch:
		return (value * 2.54e-2) 
	case LengthMil:
		return (value * 2.54e-5) 
	case LengthNauticalMile:
		return (value * 1852) 
	case LengthFathom:
		return (value * 1.8288) 
	case LengthShackle:
		return (value * 27.432) 
	case LengthMicroinch:
		return (value * 2.54e-8) 
	case LengthPrinterPoint:
		return ((value / 72.27) * 2.54e-2) 
	case LengthDtpPoint:
		return ((value / 72) * 2.54e-2) 
	case LengthPrinterPica:
		return (value / 237.106301584) 
	case LengthDtpPica:
		return (value / 236.220472441) 
	case LengthTwip:
		return (value / 56692.913385826) 
	case LengthHand:
		return (value * 1.016e-1) 
	case LengthAstronomicalUnit:
		return (value * 1.4959787070e11) 
	case LengthParsec:
		return (value * 3.08567758128e16) 
	case LengthLightYear:
		return (value * 9.46073047258e15) 
	case LengthSolarRadius:
		return (value * 6.95700e8) 
	case LengthChain:
		return (value * 20.1168) 
	case LengthAngstrom:
		return (value * 1e-10) 
	case LengthDataMile:
		return (value * 1828.8) 
	case LengthFemtometer:
		return ((value) * 1e-15) 
	case LengthPicometer:
		return ((value) * 1e-12) 
	case LengthNanometer:
		return ((value) * 1e-09) 
	case LengthMicrometer:
		return ((value) * 1e-06) 
	case LengthMillimeter:
		return ((value) * 0.001) 
	case LengthCentimeter:
		return ((value) * 0.01) 
	case LengthDecimeter:
		return ((value) * 0.1) 
	case LengthDecameter:
		return ((value) * 10.0) 
	case LengthHectometer:
		return ((value) * 100.0) 
	case LengthKilometer:
		return ((value) * 1000.0) 
	case LengthMegameter:
		return ((value) * 1000000.0) 
	case LengthGigameter:
		return ((value) * 1000000000.0) 
	case LengthKiloyard:
		return ((value * 0.9144) * 1000.0) 
	case LengthKilofoot:
		return ((value * 0.3048) * 1000.0) 
	case LengthKiloparsec:
		return ((value * 3.08567758128e16) * 1000.0) 
	case LengthMegaparsec:
		return ((value * 3.08567758128e16) * 1000000.0) 
	case LengthKilolightYear:
		return ((value * 9.46073047258e15) * 1000.0) 
	case LengthMegalightYear:
		return ((value * 9.46073047258e15) * 1000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Length in the default unit (Meter),
// formatted to two decimal places.
func (a Length) String() string {
	return a.ToString(LengthMeter, 2)
}

// ToString formats the Length value as a string with the specified unit and fractional digits.
// It converts the Length to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Length value will be converted (e.g., Meter))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Length with the unit abbreviation.
func (a *Length) ToString(unit LengthUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Length) getUnitAbbreviation(unit LengthUnits) string {
	switch unit { 
	case LengthMeter:
		return "m" 
	case LengthMile:
		return "mi" 
	case LengthYard:
		return "yd" 
	case LengthFoot:
		return "ft" 
	case LengthUsSurveyFoot:
		return "ftUS" 
	case LengthInch:
		return "in" 
	case LengthMil:
		return "mil" 
	case LengthNauticalMile:
		return "NM" 
	case LengthFathom:
		return "fathom" 
	case LengthShackle:
		return "shackle" 
	case LengthMicroinch:
		return "µin" 
	case LengthPrinterPoint:
		return "pt" 
	case LengthDtpPoint:
		return "pt" 
	case LengthPrinterPica:
		return "pica" 
	case LengthDtpPica:
		return "pica" 
	case LengthTwip:
		return "twip" 
	case LengthHand:
		return "h" 
	case LengthAstronomicalUnit:
		return "au" 
	case LengthParsec:
		return "pc" 
	case LengthLightYear:
		return "ly" 
	case LengthSolarRadius:
		return "R⊙" 
	case LengthChain:
		return "ch" 
	case LengthAngstrom:
		return "Å" 
	case LengthDataMile:
		return "DM" 
	case LengthFemtometer:
		return "fm" 
	case LengthPicometer:
		return "pm" 
	case LengthNanometer:
		return "nm" 
	case LengthMicrometer:
		return "μm" 
	case LengthMillimeter:
		return "mm" 
	case LengthCentimeter:
		return "cm" 
	case LengthDecimeter:
		return "dm" 
	case LengthDecameter:
		return "dam" 
	case LengthHectometer:
		return "hm" 
	case LengthKilometer:
		return "km" 
	case LengthMegameter:
		return "Mm" 
	case LengthGigameter:
		return "Gm" 
	case LengthKiloyard:
		return "kyd" 
	case LengthKilofoot:
		return "kft" 
	case LengthKiloparsec:
		return "kpc" 
	case LengthMegaparsec:
		return "Mpc" 
	case LengthKilolightYear:
		return "kly" 
	case LengthMegalightYear:
		return "Mly" 
	default:
		return ""
	}
}

// Equals checks if the given Length is equal to the current Length.
//
// Parameters:
//    other: The Length to compare against.
//
// Returns:
//    bool: Returns true if both Length are equal, false otherwise.
func (a *Length) Equals(other *Length) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Length with another Length.
// It returns -1 if the current Length is less than the other Length, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Length to compare against.
//
// Returns:
//    int: -1 if the current Length is less, 1 if greater, and 0 if equal.
func (a *Length) CompareTo(other *Length) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Length to the current Length and returns the result.
// The result is a new Length instance with the sum of the values.
//
// Parameters:
//    other: The Length to add to the current Length.
//
// Returns:
//    *Length: A new Length instance representing the sum of both Length.
func (a *Length) Add(other *Length) *Length {
	return &Length{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Length from the current Length and returns the result.
// The result is a new Length instance with the difference of the values.
//
// Parameters:
//    other: The Length to subtract from the current Length.
//
// Returns:
//    *Length: A new Length instance representing the difference of both Length.
func (a *Length) Subtract(other *Length) *Length {
	return &Length{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Length by the given Length and returns the result.
// The result is a new Length instance with the product of the values.
//
// Parameters:
//    other: The Length to multiply with the current Length.
//
// Returns:
//    *Length: A new Length instance representing the product of both Length.
func (a *Length) Multiply(other *Length) *Length {
	return &Length{value: a.value * other.BaseValue()}
}

// Divide divides the current Length by the given Length and returns the result.
// The result is a new Length instance with the quotient of the values.
//
// Parameters:
//    other: The Length to divide the current Length by.
//
// Returns:
//    *Length: A new Length instance representing the quotient of both Length.
func (a *Length) Divide(other *Length) *Length {
	return &Length{value: a.value / other.BaseValue()}
}