// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// LengthUnits enumeration
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

// LengthDto represents an Length
type LengthDto struct {
	Value float64
	Unit  LengthUnits
}

// LengthDtoFactory struct to group related functions
type LengthDtoFactory struct{}

func (udf LengthDtoFactory) FromJSON(data []byte) (*LengthDto, error) {
	a := LengthDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a LengthDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Length struct
type Length struct {
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

// LengthFactory struct to group related functions
type LengthFactory struct{}

func (uf LengthFactory) CreateLength(value float64, unit LengthUnits) (*Length, error) {
	return newLength(value, unit)
}

func (uf LengthFactory) FromDto(dto LengthDto) (*Length, error) {
	return newLength(dto.Value, dto.Unit)
}

func (uf LengthFactory) FromDtoJSON(data []byte) (*Length, error) {
	unitDto, err := LengthDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return LengthFactory{}.FromDto(*unitDto)
}


// FromMeter creates a new Length instance from Meter.
func (uf LengthFactory) FromMeters(value float64) (*Length, error) {
	return newLength(value, LengthMeter)
}

// FromMile creates a new Length instance from Mile.
func (uf LengthFactory) FromMiles(value float64) (*Length, error) {
	return newLength(value, LengthMile)
}

// FromYard creates a new Length instance from Yard.
func (uf LengthFactory) FromYards(value float64) (*Length, error) {
	return newLength(value, LengthYard)
}

// FromFoot creates a new Length instance from Foot.
func (uf LengthFactory) FromFeet(value float64) (*Length, error) {
	return newLength(value, LengthFoot)
}

// FromUsSurveyFoot creates a new Length instance from UsSurveyFoot.
func (uf LengthFactory) FromUsSurveyFeet(value float64) (*Length, error) {
	return newLength(value, LengthUsSurveyFoot)
}

// FromInch creates a new Length instance from Inch.
func (uf LengthFactory) FromInches(value float64) (*Length, error) {
	return newLength(value, LengthInch)
}

// FromMil creates a new Length instance from Mil.
func (uf LengthFactory) FromMils(value float64) (*Length, error) {
	return newLength(value, LengthMil)
}

// FromNauticalMile creates a new Length instance from NauticalMile.
func (uf LengthFactory) FromNauticalMiles(value float64) (*Length, error) {
	return newLength(value, LengthNauticalMile)
}

// FromFathom creates a new Length instance from Fathom.
func (uf LengthFactory) FromFathoms(value float64) (*Length, error) {
	return newLength(value, LengthFathom)
}

// FromShackle creates a new Length instance from Shackle.
func (uf LengthFactory) FromShackles(value float64) (*Length, error) {
	return newLength(value, LengthShackle)
}

// FromMicroinch creates a new Length instance from Microinch.
func (uf LengthFactory) FromMicroinches(value float64) (*Length, error) {
	return newLength(value, LengthMicroinch)
}

// FromPrinterPoint creates a new Length instance from PrinterPoint.
func (uf LengthFactory) FromPrinterPoints(value float64) (*Length, error) {
	return newLength(value, LengthPrinterPoint)
}

// FromDtpPoint creates a new Length instance from DtpPoint.
func (uf LengthFactory) FromDtpPoints(value float64) (*Length, error) {
	return newLength(value, LengthDtpPoint)
}

// FromPrinterPica creates a new Length instance from PrinterPica.
func (uf LengthFactory) FromPrinterPicas(value float64) (*Length, error) {
	return newLength(value, LengthPrinterPica)
}

// FromDtpPica creates a new Length instance from DtpPica.
func (uf LengthFactory) FromDtpPicas(value float64) (*Length, error) {
	return newLength(value, LengthDtpPica)
}

// FromTwip creates a new Length instance from Twip.
func (uf LengthFactory) FromTwips(value float64) (*Length, error) {
	return newLength(value, LengthTwip)
}

// FromHand creates a new Length instance from Hand.
func (uf LengthFactory) FromHands(value float64) (*Length, error) {
	return newLength(value, LengthHand)
}

// FromAstronomicalUnit creates a new Length instance from AstronomicalUnit.
func (uf LengthFactory) FromAstronomicalUnits(value float64) (*Length, error) {
	return newLength(value, LengthAstronomicalUnit)
}

// FromParsec creates a new Length instance from Parsec.
func (uf LengthFactory) FromParsecs(value float64) (*Length, error) {
	return newLength(value, LengthParsec)
}

// FromLightYear creates a new Length instance from LightYear.
func (uf LengthFactory) FromLightYears(value float64) (*Length, error) {
	return newLength(value, LengthLightYear)
}

// FromSolarRadius creates a new Length instance from SolarRadius.
func (uf LengthFactory) FromSolarRadiuses(value float64) (*Length, error) {
	return newLength(value, LengthSolarRadius)
}

// FromChain creates a new Length instance from Chain.
func (uf LengthFactory) FromChains(value float64) (*Length, error) {
	return newLength(value, LengthChain)
}

// FromAngstrom creates a new Length instance from Angstrom.
func (uf LengthFactory) FromAngstroms(value float64) (*Length, error) {
	return newLength(value, LengthAngstrom)
}

// FromDataMile creates a new Length instance from DataMile.
func (uf LengthFactory) FromDataMiles(value float64) (*Length, error) {
	return newLength(value, LengthDataMile)
}

// FromFemtometer creates a new Length instance from Femtometer.
func (uf LengthFactory) FromFemtometers(value float64) (*Length, error) {
	return newLength(value, LengthFemtometer)
}

// FromPicometer creates a new Length instance from Picometer.
func (uf LengthFactory) FromPicometers(value float64) (*Length, error) {
	return newLength(value, LengthPicometer)
}

// FromNanometer creates a new Length instance from Nanometer.
func (uf LengthFactory) FromNanometers(value float64) (*Length, error) {
	return newLength(value, LengthNanometer)
}

// FromMicrometer creates a new Length instance from Micrometer.
func (uf LengthFactory) FromMicrometers(value float64) (*Length, error) {
	return newLength(value, LengthMicrometer)
}

// FromMillimeter creates a new Length instance from Millimeter.
func (uf LengthFactory) FromMillimeters(value float64) (*Length, error) {
	return newLength(value, LengthMillimeter)
}

// FromCentimeter creates a new Length instance from Centimeter.
func (uf LengthFactory) FromCentimeters(value float64) (*Length, error) {
	return newLength(value, LengthCentimeter)
}

// FromDecimeter creates a new Length instance from Decimeter.
func (uf LengthFactory) FromDecimeters(value float64) (*Length, error) {
	return newLength(value, LengthDecimeter)
}

// FromDecameter creates a new Length instance from Decameter.
func (uf LengthFactory) FromDecameters(value float64) (*Length, error) {
	return newLength(value, LengthDecameter)
}

// FromHectometer creates a new Length instance from Hectometer.
func (uf LengthFactory) FromHectometers(value float64) (*Length, error) {
	return newLength(value, LengthHectometer)
}

// FromKilometer creates a new Length instance from Kilometer.
func (uf LengthFactory) FromKilometers(value float64) (*Length, error) {
	return newLength(value, LengthKilometer)
}

// FromMegameter creates a new Length instance from Megameter.
func (uf LengthFactory) FromMegameters(value float64) (*Length, error) {
	return newLength(value, LengthMegameter)
}

// FromGigameter creates a new Length instance from Gigameter.
func (uf LengthFactory) FromGigameters(value float64) (*Length, error) {
	return newLength(value, LengthGigameter)
}

// FromKiloyard creates a new Length instance from Kiloyard.
func (uf LengthFactory) FromKiloyards(value float64) (*Length, error) {
	return newLength(value, LengthKiloyard)
}

// FromKilofoot creates a new Length instance from Kilofoot.
func (uf LengthFactory) FromKilofeet(value float64) (*Length, error) {
	return newLength(value, LengthKilofoot)
}

// FromKiloparsec creates a new Length instance from Kiloparsec.
func (uf LengthFactory) FromKiloparsecs(value float64) (*Length, error) {
	return newLength(value, LengthKiloparsec)
}

// FromMegaparsec creates a new Length instance from Megaparsec.
func (uf LengthFactory) FromMegaparsecs(value float64) (*Length, error) {
	return newLength(value, LengthMegaparsec)
}

// FromKilolightYear creates a new Length instance from KilolightYear.
func (uf LengthFactory) FromKilolightYears(value float64) (*Length, error) {
	return newLength(value, LengthKilolightYear)
}

// FromMegalightYear creates a new Length instance from MegalightYear.
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

// BaseValue returns the base value of Length in Meter.
func (a *Length) BaseValue() float64 {
	return a.value
}


// Meter returns the value in Meter.
func (a *Length) Meters() float64 {
	if a.metersLazy != nil {
		return *a.metersLazy
	}
	meters := a.convertFromBase(LengthMeter)
	a.metersLazy = &meters
	return meters
}

// Mile returns the value in Mile.
func (a *Length) Miles() float64 {
	if a.milesLazy != nil {
		return *a.milesLazy
	}
	miles := a.convertFromBase(LengthMile)
	a.milesLazy = &miles
	return miles
}

// Yard returns the value in Yard.
func (a *Length) Yards() float64 {
	if a.yardsLazy != nil {
		return *a.yardsLazy
	}
	yards := a.convertFromBase(LengthYard)
	a.yardsLazy = &yards
	return yards
}

// Foot returns the value in Foot.
func (a *Length) Feet() float64 {
	if a.feetLazy != nil {
		return *a.feetLazy
	}
	feet := a.convertFromBase(LengthFoot)
	a.feetLazy = &feet
	return feet
}

// UsSurveyFoot returns the value in UsSurveyFoot.
func (a *Length) UsSurveyFeet() float64 {
	if a.us_survey_feetLazy != nil {
		return *a.us_survey_feetLazy
	}
	us_survey_feet := a.convertFromBase(LengthUsSurveyFoot)
	a.us_survey_feetLazy = &us_survey_feet
	return us_survey_feet
}

// Inch returns the value in Inch.
func (a *Length) Inches() float64 {
	if a.inchesLazy != nil {
		return *a.inchesLazy
	}
	inches := a.convertFromBase(LengthInch)
	a.inchesLazy = &inches
	return inches
}

// Mil returns the value in Mil.
func (a *Length) Mils() float64 {
	if a.milsLazy != nil {
		return *a.milsLazy
	}
	mils := a.convertFromBase(LengthMil)
	a.milsLazy = &mils
	return mils
}

// NauticalMile returns the value in NauticalMile.
func (a *Length) NauticalMiles() float64 {
	if a.nautical_milesLazy != nil {
		return *a.nautical_milesLazy
	}
	nautical_miles := a.convertFromBase(LengthNauticalMile)
	a.nautical_milesLazy = &nautical_miles
	return nautical_miles
}

// Fathom returns the value in Fathom.
func (a *Length) Fathoms() float64 {
	if a.fathomsLazy != nil {
		return *a.fathomsLazy
	}
	fathoms := a.convertFromBase(LengthFathom)
	a.fathomsLazy = &fathoms
	return fathoms
}

// Shackle returns the value in Shackle.
func (a *Length) Shackles() float64 {
	if a.shacklesLazy != nil {
		return *a.shacklesLazy
	}
	shackles := a.convertFromBase(LengthShackle)
	a.shacklesLazy = &shackles
	return shackles
}

// Microinch returns the value in Microinch.
func (a *Length) Microinches() float64 {
	if a.microinchesLazy != nil {
		return *a.microinchesLazy
	}
	microinches := a.convertFromBase(LengthMicroinch)
	a.microinchesLazy = &microinches
	return microinches
}

// PrinterPoint returns the value in PrinterPoint.
func (a *Length) PrinterPoints() float64 {
	if a.printer_pointsLazy != nil {
		return *a.printer_pointsLazy
	}
	printer_points := a.convertFromBase(LengthPrinterPoint)
	a.printer_pointsLazy = &printer_points
	return printer_points
}

// DtpPoint returns the value in DtpPoint.
func (a *Length) DtpPoints() float64 {
	if a.dtp_pointsLazy != nil {
		return *a.dtp_pointsLazy
	}
	dtp_points := a.convertFromBase(LengthDtpPoint)
	a.dtp_pointsLazy = &dtp_points
	return dtp_points
}

// PrinterPica returns the value in PrinterPica.
func (a *Length) PrinterPicas() float64 {
	if a.printer_picasLazy != nil {
		return *a.printer_picasLazy
	}
	printer_picas := a.convertFromBase(LengthPrinterPica)
	a.printer_picasLazy = &printer_picas
	return printer_picas
}

// DtpPica returns the value in DtpPica.
func (a *Length) DtpPicas() float64 {
	if a.dtp_picasLazy != nil {
		return *a.dtp_picasLazy
	}
	dtp_picas := a.convertFromBase(LengthDtpPica)
	a.dtp_picasLazy = &dtp_picas
	return dtp_picas
}

// Twip returns the value in Twip.
func (a *Length) Twips() float64 {
	if a.twipsLazy != nil {
		return *a.twipsLazy
	}
	twips := a.convertFromBase(LengthTwip)
	a.twipsLazy = &twips
	return twips
}

// Hand returns the value in Hand.
func (a *Length) Hands() float64 {
	if a.handsLazy != nil {
		return *a.handsLazy
	}
	hands := a.convertFromBase(LengthHand)
	a.handsLazy = &hands
	return hands
}

// AstronomicalUnit returns the value in AstronomicalUnit.
func (a *Length) AstronomicalUnits() float64 {
	if a.astronomical_unitsLazy != nil {
		return *a.astronomical_unitsLazy
	}
	astronomical_units := a.convertFromBase(LengthAstronomicalUnit)
	a.astronomical_unitsLazy = &astronomical_units
	return astronomical_units
}

// Parsec returns the value in Parsec.
func (a *Length) Parsecs() float64 {
	if a.parsecsLazy != nil {
		return *a.parsecsLazy
	}
	parsecs := a.convertFromBase(LengthParsec)
	a.parsecsLazy = &parsecs
	return parsecs
}

// LightYear returns the value in LightYear.
func (a *Length) LightYears() float64 {
	if a.light_yearsLazy != nil {
		return *a.light_yearsLazy
	}
	light_years := a.convertFromBase(LengthLightYear)
	a.light_yearsLazy = &light_years
	return light_years
}

// SolarRadius returns the value in SolarRadius.
func (a *Length) SolarRadiuses() float64 {
	if a.solar_radiusesLazy != nil {
		return *a.solar_radiusesLazy
	}
	solar_radiuses := a.convertFromBase(LengthSolarRadius)
	a.solar_radiusesLazy = &solar_radiuses
	return solar_radiuses
}

// Chain returns the value in Chain.
func (a *Length) Chains() float64 {
	if a.chainsLazy != nil {
		return *a.chainsLazy
	}
	chains := a.convertFromBase(LengthChain)
	a.chainsLazy = &chains
	return chains
}

// Angstrom returns the value in Angstrom.
func (a *Length) Angstroms() float64 {
	if a.angstromsLazy != nil {
		return *a.angstromsLazy
	}
	angstroms := a.convertFromBase(LengthAngstrom)
	a.angstromsLazy = &angstroms
	return angstroms
}

// DataMile returns the value in DataMile.
func (a *Length) DataMiles() float64 {
	if a.data_milesLazy != nil {
		return *a.data_milesLazy
	}
	data_miles := a.convertFromBase(LengthDataMile)
	a.data_milesLazy = &data_miles
	return data_miles
}

// Femtometer returns the value in Femtometer.
func (a *Length) Femtometers() float64 {
	if a.femtometersLazy != nil {
		return *a.femtometersLazy
	}
	femtometers := a.convertFromBase(LengthFemtometer)
	a.femtometersLazy = &femtometers
	return femtometers
}

// Picometer returns the value in Picometer.
func (a *Length) Picometers() float64 {
	if a.picometersLazy != nil {
		return *a.picometersLazy
	}
	picometers := a.convertFromBase(LengthPicometer)
	a.picometersLazy = &picometers
	return picometers
}

// Nanometer returns the value in Nanometer.
func (a *Length) Nanometers() float64 {
	if a.nanometersLazy != nil {
		return *a.nanometersLazy
	}
	nanometers := a.convertFromBase(LengthNanometer)
	a.nanometersLazy = &nanometers
	return nanometers
}

// Micrometer returns the value in Micrometer.
func (a *Length) Micrometers() float64 {
	if a.micrometersLazy != nil {
		return *a.micrometersLazy
	}
	micrometers := a.convertFromBase(LengthMicrometer)
	a.micrometersLazy = &micrometers
	return micrometers
}

// Millimeter returns the value in Millimeter.
func (a *Length) Millimeters() float64 {
	if a.millimetersLazy != nil {
		return *a.millimetersLazy
	}
	millimeters := a.convertFromBase(LengthMillimeter)
	a.millimetersLazy = &millimeters
	return millimeters
}

// Centimeter returns the value in Centimeter.
func (a *Length) Centimeters() float64 {
	if a.centimetersLazy != nil {
		return *a.centimetersLazy
	}
	centimeters := a.convertFromBase(LengthCentimeter)
	a.centimetersLazy = &centimeters
	return centimeters
}

// Decimeter returns the value in Decimeter.
func (a *Length) Decimeters() float64 {
	if a.decimetersLazy != nil {
		return *a.decimetersLazy
	}
	decimeters := a.convertFromBase(LengthDecimeter)
	a.decimetersLazy = &decimeters
	return decimeters
}

// Decameter returns the value in Decameter.
func (a *Length) Decameters() float64 {
	if a.decametersLazy != nil {
		return *a.decametersLazy
	}
	decameters := a.convertFromBase(LengthDecameter)
	a.decametersLazy = &decameters
	return decameters
}

// Hectometer returns the value in Hectometer.
func (a *Length) Hectometers() float64 {
	if a.hectometersLazy != nil {
		return *a.hectometersLazy
	}
	hectometers := a.convertFromBase(LengthHectometer)
	a.hectometersLazy = &hectometers
	return hectometers
}

// Kilometer returns the value in Kilometer.
func (a *Length) Kilometers() float64 {
	if a.kilometersLazy != nil {
		return *a.kilometersLazy
	}
	kilometers := a.convertFromBase(LengthKilometer)
	a.kilometersLazy = &kilometers
	return kilometers
}

// Megameter returns the value in Megameter.
func (a *Length) Megameters() float64 {
	if a.megametersLazy != nil {
		return *a.megametersLazy
	}
	megameters := a.convertFromBase(LengthMegameter)
	a.megametersLazy = &megameters
	return megameters
}

// Gigameter returns the value in Gigameter.
func (a *Length) Gigameters() float64 {
	if a.gigametersLazy != nil {
		return *a.gigametersLazy
	}
	gigameters := a.convertFromBase(LengthGigameter)
	a.gigametersLazy = &gigameters
	return gigameters
}

// Kiloyard returns the value in Kiloyard.
func (a *Length) Kiloyards() float64 {
	if a.kiloyardsLazy != nil {
		return *a.kiloyardsLazy
	}
	kiloyards := a.convertFromBase(LengthKiloyard)
	a.kiloyardsLazy = &kiloyards
	return kiloyards
}

// Kilofoot returns the value in Kilofoot.
func (a *Length) Kilofeet() float64 {
	if a.kilofeetLazy != nil {
		return *a.kilofeetLazy
	}
	kilofeet := a.convertFromBase(LengthKilofoot)
	a.kilofeetLazy = &kilofeet
	return kilofeet
}

// Kiloparsec returns the value in Kiloparsec.
func (a *Length) Kiloparsecs() float64 {
	if a.kiloparsecsLazy != nil {
		return *a.kiloparsecsLazy
	}
	kiloparsecs := a.convertFromBase(LengthKiloparsec)
	a.kiloparsecsLazy = &kiloparsecs
	return kiloparsecs
}

// Megaparsec returns the value in Megaparsec.
func (a *Length) Megaparsecs() float64 {
	if a.megaparsecsLazy != nil {
		return *a.megaparsecsLazy
	}
	megaparsecs := a.convertFromBase(LengthMegaparsec)
	a.megaparsecsLazy = &megaparsecs
	return megaparsecs
}

// KilolightYear returns the value in KilolightYear.
func (a *Length) KilolightYears() float64 {
	if a.kilolight_yearsLazy != nil {
		return *a.kilolight_yearsLazy
	}
	kilolight_years := a.convertFromBase(LengthKilolightYear)
	a.kilolight_yearsLazy = &kilolight_years
	return kilolight_years
}

// MegalightYear returns the value in MegalightYear.
func (a *Length) MegalightYears() float64 {
	if a.megalight_yearsLazy != nil {
		return *a.megalight_yearsLazy
	}
	megalight_years := a.convertFromBase(LengthMegalightYear)
	a.megalight_yearsLazy = &megalight_years
	return megalight_years
}


// ToDto creates an LengthDto representation.
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

// ToDtoJSON creates an LengthDto representation.
func (a *Length) ToDtoJSON(holdInUnit *LengthUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Length to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Length) String() string {
	return a.ToString(LengthMeter, 2)
}

// ToString formats the Length to string.
// fractionalDigits -1 for not mention
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

// Check if the given Length are equals to the current Length
func (a *Length) Equals(other *Length) bool {
	return a.value == other.BaseValue()
}

// Check if the given Length are equals to the current Length
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

// Add the given Length to the current Length.
func (a *Length) Add(other *Length) *Length {
	return &Length{value: a.value + other.BaseValue()}
}

// Subtract the given Length to the current Length.
func (a *Length) Subtract(other *Length) *Length {
	return &Length{value: a.value - other.BaseValue()}
}

// Multiply the given Length to the current Length.
func (a *Length) Multiply(other *Length) *Length {
	return &Length{value: a.value * other.BaseValue()}
}

// Divide the given Length to the current Length.
func (a *Length) Divide(other *Length) *Length {
	return &Length{value: a.value / other.BaseValue()}
}