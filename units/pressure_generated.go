// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PressureUnits defines various units of Pressure.
type PressureUnits string

const (
    
        // 
        PressurePascal PressureUnits = "Pascal"
        // 
        PressureAtmosphere PressureUnits = "Atmosphere"
        // 
        PressureBar PressureUnits = "Bar"
        // 
        PressureKilogramForcePerSquareMeter PressureUnits = "KilogramForcePerSquareMeter"
        // 
        PressureKilogramForcePerSquareCentimeter PressureUnits = "KilogramForcePerSquareCentimeter"
        // 
        PressureKilogramForcePerSquareMillimeter PressureUnits = "KilogramForcePerSquareMillimeter"
        // 
        PressureNewtonPerSquareMeter PressureUnits = "NewtonPerSquareMeter"
        // 
        PressureNewtonPerSquareCentimeter PressureUnits = "NewtonPerSquareCentimeter"
        // 
        PressureNewtonPerSquareMillimeter PressureUnits = "NewtonPerSquareMillimeter"
        // 
        PressureTechnicalAtmosphere PressureUnits = "TechnicalAtmosphere"
        // 
        PressureTorr PressureUnits = "Torr"
        // 
        PressurePoundForcePerSquareInch PressureUnits = "PoundForcePerSquareInch"
        // 
        PressurePoundForcePerSquareMil PressureUnits = "PoundForcePerSquareMil"
        // 
        PressurePoundForcePerSquareFoot PressureUnits = "PoundForcePerSquareFoot"
        // 
        PressureTonneForcePerSquareMillimeter PressureUnits = "TonneForcePerSquareMillimeter"
        // 
        PressureTonneForcePerSquareMeter PressureUnits = "TonneForcePerSquareMeter"
        // 
        PressureMeterOfHead PressureUnits = "MeterOfHead"
        // 
        PressureTonneForcePerSquareCentimeter PressureUnits = "TonneForcePerSquareCentimeter"
        // 
        PressureFootOfHead PressureUnits = "FootOfHead"
        // 
        PressureMillimeterOfMercury PressureUnits = "MillimeterOfMercury"
        // 
        PressureInchOfMercury PressureUnits = "InchOfMercury"
        // 
        PressureDynePerSquareCentimeter PressureUnits = "DynePerSquareCentimeter"
        // 
        PressurePoundPerInchSecondSquared PressureUnits = "PoundPerInchSecondSquared"
        // 
        PressureMeterOfWaterColumn PressureUnits = "MeterOfWaterColumn"
        // 
        PressureInchOfWaterColumn PressureUnits = "InchOfWaterColumn"
        // 
        PressureMeterOfElevation PressureUnits = "MeterOfElevation"
        // 
        PressureFootOfElevation PressureUnits = "FootOfElevation"
        // 
        PressureMicropascal PressureUnits = "Micropascal"
        // 
        PressureMillipascal PressureUnits = "Millipascal"
        // 
        PressureDecapascal PressureUnits = "Decapascal"
        // 
        PressureHectopascal PressureUnits = "Hectopascal"
        // 
        PressureKilopascal PressureUnits = "Kilopascal"
        // 
        PressureMegapascal PressureUnits = "Megapascal"
        // 
        PressureGigapascal PressureUnits = "Gigapascal"
        // 
        PressureMicrobar PressureUnits = "Microbar"
        // 
        PressureMillibar PressureUnits = "Millibar"
        // 
        PressureCentibar PressureUnits = "Centibar"
        // 
        PressureDecibar PressureUnits = "Decibar"
        // 
        PressureKilobar PressureUnits = "Kilobar"
        // 
        PressureMegabar PressureUnits = "Megabar"
        // 
        PressureKilonewtonPerSquareMeter PressureUnits = "KilonewtonPerSquareMeter"
        // 
        PressureMeganewtonPerSquareMeter PressureUnits = "MeganewtonPerSquareMeter"
        // 
        PressureKilonewtonPerSquareCentimeter PressureUnits = "KilonewtonPerSquareCentimeter"
        // 
        PressureKilonewtonPerSquareMillimeter PressureUnits = "KilonewtonPerSquareMillimeter"
        // 
        PressureKilopoundForcePerSquareInch PressureUnits = "KilopoundForcePerSquareInch"
        // 
        PressureKilopoundForcePerSquareMil PressureUnits = "KilopoundForcePerSquareMil"
        // 
        PressureKilopoundForcePerSquareFoot PressureUnits = "KilopoundForcePerSquareFoot"
        // 
        PressureMillimeterOfWaterColumn PressureUnits = "MillimeterOfWaterColumn"
        // 
        PressureCentimeterOfWaterColumn PressureUnits = "CentimeterOfWaterColumn"
)

var internalPressureUnitsMap = map[PressureUnits]bool{
	
	PressurePascal: true,
	PressureAtmosphere: true,
	PressureBar: true,
	PressureKilogramForcePerSquareMeter: true,
	PressureKilogramForcePerSquareCentimeter: true,
	PressureKilogramForcePerSquareMillimeter: true,
	PressureNewtonPerSquareMeter: true,
	PressureNewtonPerSquareCentimeter: true,
	PressureNewtonPerSquareMillimeter: true,
	PressureTechnicalAtmosphere: true,
	PressureTorr: true,
	PressurePoundForcePerSquareInch: true,
	PressurePoundForcePerSquareMil: true,
	PressurePoundForcePerSquareFoot: true,
	PressureTonneForcePerSquareMillimeter: true,
	PressureTonneForcePerSquareMeter: true,
	PressureMeterOfHead: true,
	PressureTonneForcePerSquareCentimeter: true,
	PressureFootOfHead: true,
	PressureMillimeterOfMercury: true,
	PressureInchOfMercury: true,
	PressureDynePerSquareCentimeter: true,
	PressurePoundPerInchSecondSquared: true,
	PressureMeterOfWaterColumn: true,
	PressureInchOfWaterColumn: true,
	PressureMeterOfElevation: true,
	PressureFootOfElevation: true,
	PressureMicropascal: true,
	PressureMillipascal: true,
	PressureDecapascal: true,
	PressureHectopascal: true,
	PressureKilopascal: true,
	PressureMegapascal: true,
	PressureGigapascal: true,
	PressureMicrobar: true,
	PressureMillibar: true,
	PressureCentibar: true,
	PressureDecibar: true,
	PressureKilobar: true,
	PressureMegabar: true,
	PressureKilonewtonPerSquareMeter: true,
	PressureMeganewtonPerSquareMeter: true,
	PressureKilonewtonPerSquareCentimeter: true,
	PressureKilonewtonPerSquareMillimeter: true,
	PressureKilopoundForcePerSquareInch: true,
	PressureKilopoundForcePerSquareMil: true,
	PressureKilopoundForcePerSquareFoot: true,
	PressureMillimeterOfWaterColumn: true,
	PressureCentimeterOfWaterColumn: true,
}

// PressureDto represents a Pressure measurement with a numerical value and its corresponding unit.
type PressureDto struct {
    // Value is the numerical representation of the Pressure.
	Value float64 `json:"value" validate:"required"`
    // Unit specifies the unit of measurement for the Pressure, as defined in the PressureUnits enumeration.
	Unit  PressureUnits `json:"unit" validate:"required,oneof=Pascal Atmosphere Bar KilogramForcePerSquareMeter KilogramForcePerSquareCentimeter KilogramForcePerSquareMillimeter NewtonPerSquareMeter NewtonPerSquareCentimeter NewtonPerSquareMillimeter TechnicalAtmosphere Torr PoundForcePerSquareInch PoundForcePerSquareMil PoundForcePerSquareFoot TonneForcePerSquareMillimeter TonneForcePerSquareMeter MeterOfHead TonneForcePerSquareCentimeter FootOfHead MillimeterOfMercury InchOfMercury DynePerSquareCentimeter PoundPerInchSecondSquared MeterOfWaterColumn InchOfWaterColumn MeterOfElevation FootOfElevation Micropascal Millipascal Decapascal Hectopascal Kilopascal Megapascal Gigapascal Microbar Millibar Centibar Decibar Kilobar Megabar KilonewtonPerSquareMeter MeganewtonPerSquareMeter KilonewtonPerSquareCentimeter KilonewtonPerSquareMillimeter KilopoundForcePerSquareInch KilopoundForcePerSquareMil KilopoundForcePerSquareFoot MillimeterOfWaterColumn CentimeterOfWaterColumn"`
}

// PressureDtoFactory groups methods for creating and serializing PressureDto objects.
type PressureDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PressureDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PressureDtoFactory) FromJSON(data []byte) (*PressureDto, error) {
	a := PressureDto{}

    // Parse JSON into PressureDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a PressureDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PressureDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Pressure represents a measurement in a of Pressure.
//
// Pressure (symbol: P or p) is the ratio of force to the area over which that force is distributed. Pressure is force per unit area applied in a direction perpendicular to the surface of an object. Gauge pressure (also spelled gage pressure)[a] is the pressure relative to the local atmospheric or ambient pressure. Pressure is measured in any unit of force divided by any unit of area. The SI unit of pressure is the newton per square metre, which is called the pascal (Pa) after the seventeenth-century philosopher and scientist Blaise Pascal. A pressure of 1 Pa is small; it approximately equals the pressure exerted by a dollar bill resting flat on a table. Everyday pressures are often stated in kilopascals (1 kPa = 1000 Pa).
type Pressure struct {
	// value is the base measurement stored internally.
	value       float64
    
    pascalsLazy *float64 
    atmospheresLazy *float64 
    barsLazy *float64 
    kilograms_force_per_square_meterLazy *float64 
    kilograms_force_per_square_centimeterLazy *float64 
    kilograms_force_per_square_millimeterLazy *float64 
    newtons_per_square_meterLazy *float64 
    newtons_per_square_centimeterLazy *float64 
    newtons_per_square_millimeterLazy *float64 
    technical_atmospheresLazy *float64 
    torrsLazy *float64 
    pounds_force_per_square_inchLazy *float64 
    pounds_force_per_square_milLazy *float64 
    pounds_force_per_square_footLazy *float64 
    tonnes_force_per_square_millimeterLazy *float64 
    tonnes_force_per_square_meterLazy *float64 
    meters_of_headLazy *float64 
    tonnes_force_per_square_centimeterLazy *float64 
    feet_of_headLazy *float64 
    millimeters_of_mercuryLazy *float64 
    inches_of_mercuryLazy *float64 
    dynes_per_square_centimeterLazy *float64 
    pounds_per_inch_second_squaredLazy *float64 
    meters_of_water_columnLazy *float64 
    inches_of_water_columnLazy *float64 
    meters_of_elevationLazy *float64 
    feet_of_elevationLazy *float64 
    micropascalsLazy *float64 
    millipascalsLazy *float64 
    decapascalsLazy *float64 
    hectopascalsLazy *float64 
    kilopascalsLazy *float64 
    megapascalsLazy *float64 
    gigapascalsLazy *float64 
    microbarsLazy *float64 
    millibarsLazy *float64 
    centibarsLazy *float64 
    decibarsLazy *float64 
    kilobarsLazy *float64 
    megabarsLazy *float64 
    kilonewtons_per_square_meterLazy *float64 
    meganewtons_per_square_meterLazy *float64 
    kilonewtons_per_square_centimeterLazy *float64 
    kilonewtons_per_square_millimeterLazy *float64 
    kilopounds_force_per_square_inchLazy *float64 
    kilopounds_force_per_square_milLazy *float64 
    kilopounds_force_per_square_footLazy *float64 
    millimeters_of_water_columnLazy *float64 
    centimeters_of_water_columnLazy *float64 
}

// PressureFactory groups methods for creating Pressure instances.
type PressureFactory struct{}

// CreatePressure creates a new Pressure instance from the given value and unit.
func (uf PressureFactory) CreatePressure(value float64, unit PressureUnits) (*Pressure, error) {
	return newPressure(value, unit)
}

// FromDto converts a PressureDto to a Pressure instance.
func (uf PressureFactory) FromDto(dto PressureDto) (*Pressure, error) {
	return newPressure(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Pressure instance.
func (uf PressureFactory) FromDtoJSON(data []byte) (*Pressure, error) {
	unitDto, err := PressureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PressureDto from JSON: %w", err)
	}
	return PressureFactory{}.FromDto(*unitDto)
}


// FromPascals creates a new Pressure instance from a value in Pascals.
func (uf PressureFactory) FromPascals(value float64) (*Pressure, error) {
	return newPressure(value, PressurePascal)
}

// FromAtmospheres creates a new Pressure instance from a value in Atmospheres.
func (uf PressureFactory) FromAtmospheres(value float64) (*Pressure, error) {
	return newPressure(value, PressureAtmosphere)
}

// FromBars creates a new Pressure instance from a value in Bars.
func (uf PressureFactory) FromBars(value float64) (*Pressure, error) {
	return newPressure(value, PressureBar)
}

// FromKilogramsForcePerSquareMeter creates a new Pressure instance from a value in KilogramsForcePerSquareMeter.
func (uf PressureFactory) FromKilogramsForcePerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareMeter)
}

// FromKilogramsForcePerSquareCentimeter creates a new Pressure instance from a value in KilogramsForcePerSquareCentimeter.
func (uf PressureFactory) FromKilogramsForcePerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareCentimeter)
}

// FromKilogramsForcePerSquareMillimeter creates a new Pressure instance from a value in KilogramsForcePerSquareMillimeter.
func (uf PressureFactory) FromKilogramsForcePerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareMillimeter)
}

// FromNewtonsPerSquareMeter creates a new Pressure instance from a value in NewtonsPerSquareMeter.
func (uf PressureFactory) FromNewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareMeter)
}

// FromNewtonsPerSquareCentimeter creates a new Pressure instance from a value in NewtonsPerSquareCentimeter.
func (uf PressureFactory) FromNewtonsPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareCentimeter)
}

// FromNewtonsPerSquareMillimeter creates a new Pressure instance from a value in NewtonsPerSquareMillimeter.
func (uf PressureFactory) FromNewtonsPerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareMillimeter)
}

// FromTechnicalAtmospheres creates a new Pressure instance from a value in TechnicalAtmospheres.
func (uf PressureFactory) FromTechnicalAtmospheres(value float64) (*Pressure, error) {
	return newPressure(value, PressureTechnicalAtmosphere)
}

// FromTorrs creates a new Pressure instance from a value in Torrs.
func (uf PressureFactory) FromTorrs(value float64) (*Pressure, error) {
	return newPressure(value, PressureTorr)
}

// FromPoundsForcePerSquareInch creates a new Pressure instance from a value in PoundsForcePerSquareInch.
func (uf PressureFactory) FromPoundsForcePerSquareInch(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareInch)
}

// FromPoundsForcePerSquareMil creates a new Pressure instance from a value in PoundsForcePerSquareMil.
func (uf PressureFactory) FromPoundsForcePerSquareMil(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareMil)
}

// FromPoundsForcePerSquareFoot creates a new Pressure instance from a value in PoundsForcePerSquareFoot.
func (uf PressureFactory) FromPoundsForcePerSquareFoot(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareFoot)
}

// FromTonnesForcePerSquareMillimeter creates a new Pressure instance from a value in TonnesForcePerSquareMillimeter.
func (uf PressureFactory) FromTonnesForcePerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareMillimeter)
}

// FromTonnesForcePerSquareMeter creates a new Pressure instance from a value in TonnesForcePerSquareMeter.
func (uf PressureFactory) FromTonnesForcePerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareMeter)
}

// FromMetersOfHead creates a new Pressure instance from a value in MetersOfHead.
func (uf PressureFactory) FromMetersOfHead(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfHead)
}

// FromTonnesForcePerSquareCentimeter creates a new Pressure instance from a value in TonnesForcePerSquareCentimeter.
func (uf PressureFactory) FromTonnesForcePerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareCentimeter)
}

// FromFeetOfHead creates a new Pressure instance from a value in FeetOfHead.
func (uf PressureFactory) FromFeetOfHead(value float64) (*Pressure, error) {
	return newPressure(value, PressureFootOfHead)
}

// FromMillimetersOfMercury creates a new Pressure instance from a value in MillimetersOfMercury.
func (uf PressureFactory) FromMillimetersOfMercury(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillimeterOfMercury)
}

// FromInchesOfMercury creates a new Pressure instance from a value in InchesOfMercury.
func (uf PressureFactory) FromInchesOfMercury(value float64) (*Pressure, error) {
	return newPressure(value, PressureInchOfMercury)
}

// FromDynesPerSquareCentimeter creates a new Pressure instance from a value in DynesPerSquareCentimeter.
func (uf PressureFactory) FromDynesPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureDynePerSquareCentimeter)
}

// FromPoundsPerInchSecondSquared creates a new Pressure instance from a value in PoundsPerInchSecondSquared.
func (uf PressureFactory) FromPoundsPerInchSecondSquared(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundPerInchSecondSquared)
}

// FromMetersOfWaterColumn creates a new Pressure instance from a value in MetersOfWaterColumn.
func (uf PressureFactory) FromMetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfWaterColumn)
}

// FromInchesOfWaterColumn creates a new Pressure instance from a value in InchesOfWaterColumn.
func (uf PressureFactory) FromInchesOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureInchOfWaterColumn)
}

// FromMetersOfElevation creates a new Pressure instance from a value in MetersOfElevation.
func (uf PressureFactory) FromMetersOfElevation(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfElevation)
}

// FromFeetOfElevation creates a new Pressure instance from a value in FeetOfElevation.
func (uf PressureFactory) FromFeetOfElevation(value float64) (*Pressure, error) {
	return newPressure(value, PressureFootOfElevation)
}

// FromMicropascals creates a new Pressure instance from a value in Micropascals.
func (uf PressureFactory) FromMicropascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMicropascal)
}

// FromMillipascals creates a new Pressure instance from a value in Millipascals.
func (uf PressureFactory) FromMillipascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillipascal)
}

// FromDecapascals creates a new Pressure instance from a value in Decapascals.
func (uf PressureFactory) FromDecapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureDecapascal)
}

// FromHectopascals creates a new Pressure instance from a value in Hectopascals.
func (uf PressureFactory) FromHectopascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureHectopascal)
}

// FromKilopascals creates a new Pressure instance from a value in Kilopascals.
func (uf PressureFactory) FromKilopascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopascal)
}

// FromMegapascals creates a new Pressure instance from a value in Megapascals.
func (uf PressureFactory) FromMegapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMegapascal)
}

// FromGigapascals creates a new Pressure instance from a value in Gigapascals.
func (uf PressureFactory) FromGigapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureGigapascal)
}

// FromMicrobars creates a new Pressure instance from a value in Microbars.
func (uf PressureFactory) FromMicrobars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMicrobar)
}

// FromMillibars creates a new Pressure instance from a value in Millibars.
func (uf PressureFactory) FromMillibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillibar)
}

// FromCentibars creates a new Pressure instance from a value in Centibars.
func (uf PressureFactory) FromCentibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureCentibar)
}

// FromDecibars creates a new Pressure instance from a value in Decibars.
func (uf PressureFactory) FromDecibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureDecibar)
}

// FromKilobars creates a new Pressure instance from a value in Kilobars.
func (uf PressureFactory) FromKilobars(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilobar)
}

// FromMegabars creates a new Pressure instance from a value in Megabars.
func (uf PressureFactory) FromMegabars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMegabar)
}

// FromKilonewtonsPerSquareMeter creates a new Pressure instance from a value in KilonewtonsPerSquareMeter.
func (uf PressureFactory) FromKilonewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareMeter)
}

// FromMeganewtonsPerSquareMeter creates a new Pressure instance from a value in MeganewtonsPerSquareMeter.
func (uf PressureFactory) FromMeganewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeganewtonPerSquareMeter)
}

// FromKilonewtonsPerSquareCentimeter creates a new Pressure instance from a value in KilonewtonsPerSquareCentimeter.
func (uf PressureFactory) FromKilonewtonsPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareCentimeter)
}

// FromKilonewtonsPerSquareMillimeter creates a new Pressure instance from a value in KilonewtonsPerSquareMillimeter.
func (uf PressureFactory) FromKilonewtonsPerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareMillimeter)
}

// FromKilopoundsForcePerSquareInch creates a new Pressure instance from a value in KilopoundsForcePerSquareInch.
func (uf PressureFactory) FromKilopoundsForcePerSquareInch(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareInch)
}

// FromKilopoundsForcePerSquareMil creates a new Pressure instance from a value in KilopoundsForcePerSquareMil.
func (uf PressureFactory) FromKilopoundsForcePerSquareMil(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareMil)
}

// FromKilopoundsForcePerSquareFoot creates a new Pressure instance from a value in KilopoundsForcePerSquareFoot.
func (uf PressureFactory) FromKilopoundsForcePerSquareFoot(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareFoot)
}

// FromMillimetersOfWaterColumn creates a new Pressure instance from a value in MillimetersOfWaterColumn.
func (uf PressureFactory) FromMillimetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillimeterOfWaterColumn)
}

// FromCentimetersOfWaterColumn creates a new Pressure instance from a value in CentimetersOfWaterColumn.
func (uf PressureFactory) FromCentimetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureCentimeterOfWaterColumn)
}


// newPressure creates a new Pressure.
func newPressure(value float64, fromUnit PressureUnits) (*Pressure, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalPressureUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in PressureUnits", fromUnit)
	}
	a := &Pressure{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Pressure in Pascal unit (the base/default quantity).
func (a *Pressure) BaseValue() float64 {
	return a.value
}


// Pascals returns the Pressure value in Pascals.
//
// 
func (a *Pressure) Pascals() float64 {
	if a.pascalsLazy != nil {
		return *a.pascalsLazy
	}
	pascals := a.convertFromBase(PressurePascal)
	a.pascalsLazy = &pascals
	return pascals
}

// Atmospheres returns the Pressure value in Atmospheres.
//
// 
func (a *Pressure) Atmospheres() float64 {
	if a.atmospheresLazy != nil {
		return *a.atmospheresLazy
	}
	atmospheres := a.convertFromBase(PressureAtmosphere)
	a.atmospheresLazy = &atmospheres
	return atmospheres
}

// Bars returns the Pressure value in Bars.
//
// 
func (a *Pressure) Bars() float64 {
	if a.barsLazy != nil {
		return *a.barsLazy
	}
	bars := a.convertFromBase(PressureBar)
	a.barsLazy = &bars
	return bars
}

// KilogramsForcePerSquareMeter returns the Pressure value in KilogramsForcePerSquareMeter.
//
// 
func (a *Pressure) KilogramsForcePerSquareMeter() float64 {
	if a.kilograms_force_per_square_meterLazy != nil {
		return *a.kilograms_force_per_square_meterLazy
	}
	kilograms_force_per_square_meter := a.convertFromBase(PressureKilogramForcePerSquareMeter)
	a.kilograms_force_per_square_meterLazy = &kilograms_force_per_square_meter
	return kilograms_force_per_square_meter
}

// KilogramsForcePerSquareCentimeter returns the Pressure value in KilogramsForcePerSquareCentimeter.
//
// 
func (a *Pressure) KilogramsForcePerSquareCentimeter() float64 {
	if a.kilograms_force_per_square_centimeterLazy != nil {
		return *a.kilograms_force_per_square_centimeterLazy
	}
	kilograms_force_per_square_centimeter := a.convertFromBase(PressureKilogramForcePerSquareCentimeter)
	a.kilograms_force_per_square_centimeterLazy = &kilograms_force_per_square_centimeter
	return kilograms_force_per_square_centimeter
}

// KilogramsForcePerSquareMillimeter returns the Pressure value in KilogramsForcePerSquareMillimeter.
//
// 
func (a *Pressure) KilogramsForcePerSquareMillimeter() float64 {
	if a.kilograms_force_per_square_millimeterLazy != nil {
		return *a.kilograms_force_per_square_millimeterLazy
	}
	kilograms_force_per_square_millimeter := a.convertFromBase(PressureKilogramForcePerSquareMillimeter)
	a.kilograms_force_per_square_millimeterLazy = &kilograms_force_per_square_millimeter
	return kilograms_force_per_square_millimeter
}

// NewtonsPerSquareMeter returns the Pressure value in NewtonsPerSquareMeter.
//
// 
func (a *Pressure) NewtonsPerSquareMeter() float64 {
	if a.newtons_per_square_meterLazy != nil {
		return *a.newtons_per_square_meterLazy
	}
	newtons_per_square_meter := a.convertFromBase(PressureNewtonPerSquareMeter)
	a.newtons_per_square_meterLazy = &newtons_per_square_meter
	return newtons_per_square_meter
}

// NewtonsPerSquareCentimeter returns the Pressure value in NewtonsPerSquareCentimeter.
//
// 
func (a *Pressure) NewtonsPerSquareCentimeter() float64 {
	if a.newtons_per_square_centimeterLazy != nil {
		return *a.newtons_per_square_centimeterLazy
	}
	newtons_per_square_centimeter := a.convertFromBase(PressureNewtonPerSquareCentimeter)
	a.newtons_per_square_centimeterLazy = &newtons_per_square_centimeter
	return newtons_per_square_centimeter
}

// NewtonsPerSquareMillimeter returns the Pressure value in NewtonsPerSquareMillimeter.
//
// 
func (a *Pressure) NewtonsPerSquareMillimeter() float64 {
	if a.newtons_per_square_millimeterLazy != nil {
		return *a.newtons_per_square_millimeterLazy
	}
	newtons_per_square_millimeter := a.convertFromBase(PressureNewtonPerSquareMillimeter)
	a.newtons_per_square_millimeterLazy = &newtons_per_square_millimeter
	return newtons_per_square_millimeter
}

// TechnicalAtmospheres returns the Pressure value in TechnicalAtmospheres.
//
// 
func (a *Pressure) TechnicalAtmospheres() float64 {
	if a.technical_atmospheresLazy != nil {
		return *a.technical_atmospheresLazy
	}
	technical_atmospheres := a.convertFromBase(PressureTechnicalAtmosphere)
	a.technical_atmospheresLazy = &technical_atmospheres
	return technical_atmospheres
}

// Torrs returns the Pressure value in Torrs.
//
// 
func (a *Pressure) Torrs() float64 {
	if a.torrsLazy != nil {
		return *a.torrsLazy
	}
	torrs := a.convertFromBase(PressureTorr)
	a.torrsLazy = &torrs
	return torrs
}

// PoundsForcePerSquareInch returns the Pressure value in PoundsForcePerSquareInch.
//
// 
func (a *Pressure) PoundsForcePerSquareInch() float64 {
	if a.pounds_force_per_square_inchLazy != nil {
		return *a.pounds_force_per_square_inchLazy
	}
	pounds_force_per_square_inch := a.convertFromBase(PressurePoundForcePerSquareInch)
	a.pounds_force_per_square_inchLazy = &pounds_force_per_square_inch
	return pounds_force_per_square_inch
}

// PoundsForcePerSquareMil returns the Pressure value in PoundsForcePerSquareMil.
//
// 
func (a *Pressure) PoundsForcePerSquareMil() float64 {
	if a.pounds_force_per_square_milLazy != nil {
		return *a.pounds_force_per_square_milLazy
	}
	pounds_force_per_square_mil := a.convertFromBase(PressurePoundForcePerSquareMil)
	a.pounds_force_per_square_milLazy = &pounds_force_per_square_mil
	return pounds_force_per_square_mil
}

// PoundsForcePerSquareFoot returns the Pressure value in PoundsForcePerSquareFoot.
//
// 
func (a *Pressure) PoundsForcePerSquareFoot() float64 {
	if a.pounds_force_per_square_footLazy != nil {
		return *a.pounds_force_per_square_footLazy
	}
	pounds_force_per_square_foot := a.convertFromBase(PressurePoundForcePerSquareFoot)
	a.pounds_force_per_square_footLazy = &pounds_force_per_square_foot
	return pounds_force_per_square_foot
}

// TonnesForcePerSquareMillimeter returns the Pressure value in TonnesForcePerSquareMillimeter.
//
// 
func (a *Pressure) TonnesForcePerSquareMillimeter() float64 {
	if a.tonnes_force_per_square_millimeterLazy != nil {
		return *a.tonnes_force_per_square_millimeterLazy
	}
	tonnes_force_per_square_millimeter := a.convertFromBase(PressureTonneForcePerSquareMillimeter)
	a.tonnes_force_per_square_millimeterLazy = &tonnes_force_per_square_millimeter
	return tonnes_force_per_square_millimeter
}

// TonnesForcePerSquareMeter returns the Pressure value in TonnesForcePerSquareMeter.
//
// 
func (a *Pressure) TonnesForcePerSquareMeter() float64 {
	if a.tonnes_force_per_square_meterLazy != nil {
		return *a.tonnes_force_per_square_meterLazy
	}
	tonnes_force_per_square_meter := a.convertFromBase(PressureTonneForcePerSquareMeter)
	a.tonnes_force_per_square_meterLazy = &tonnes_force_per_square_meter
	return tonnes_force_per_square_meter
}

// MetersOfHead returns the Pressure value in MetersOfHead.
//
// 
func (a *Pressure) MetersOfHead() float64 {
	if a.meters_of_headLazy != nil {
		return *a.meters_of_headLazy
	}
	meters_of_head := a.convertFromBase(PressureMeterOfHead)
	a.meters_of_headLazy = &meters_of_head
	return meters_of_head
}

// TonnesForcePerSquareCentimeter returns the Pressure value in TonnesForcePerSquareCentimeter.
//
// 
func (a *Pressure) TonnesForcePerSquareCentimeter() float64 {
	if a.tonnes_force_per_square_centimeterLazy != nil {
		return *a.tonnes_force_per_square_centimeterLazy
	}
	tonnes_force_per_square_centimeter := a.convertFromBase(PressureTonneForcePerSquareCentimeter)
	a.tonnes_force_per_square_centimeterLazy = &tonnes_force_per_square_centimeter
	return tonnes_force_per_square_centimeter
}

// FeetOfHead returns the Pressure value in FeetOfHead.
//
// 
func (a *Pressure) FeetOfHead() float64 {
	if a.feet_of_headLazy != nil {
		return *a.feet_of_headLazy
	}
	feet_of_head := a.convertFromBase(PressureFootOfHead)
	a.feet_of_headLazy = &feet_of_head
	return feet_of_head
}

// MillimetersOfMercury returns the Pressure value in MillimetersOfMercury.
//
// 
func (a *Pressure) MillimetersOfMercury() float64 {
	if a.millimeters_of_mercuryLazy != nil {
		return *a.millimeters_of_mercuryLazy
	}
	millimeters_of_mercury := a.convertFromBase(PressureMillimeterOfMercury)
	a.millimeters_of_mercuryLazy = &millimeters_of_mercury
	return millimeters_of_mercury
}

// InchesOfMercury returns the Pressure value in InchesOfMercury.
//
// 
func (a *Pressure) InchesOfMercury() float64 {
	if a.inches_of_mercuryLazy != nil {
		return *a.inches_of_mercuryLazy
	}
	inches_of_mercury := a.convertFromBase(PressureInchOfMercury)
	a.inches_of_mercuryLazy = &inches_of_mercury
	return inches_of_mercury
}

// DynesPerSquareCentimeter returns the Pressure value in DynesPerSquareCentimeter.
//
// 
func (a *Pressure) DynesPerSquareCentimeter() float64 {
	if a.dynes_per_square_centimeterLazy != nil {
		return *a.dynes_per_square_centimeterLazy
	}
	dynes_per_square_centimeter := a.convertFromBase(PressureDynePerSquareCentimeter)
	a.dynes_per_square_centimeterLazy = &dynes_per_square_centimeter
	return dynes_per_square_centimeter
}

// PoundsPerInchSecondSquared returns the Pressure value in PoundsPerInchSecondSquared.
//
// 
func (a *Pressure) PoundsPerInchSecondSquared() float64 {
	if a.pounds_per_inch_second_squaredLazy != nil {
		return *a.pounds_per_inch_second_squaredLazy
	}
	pounds_per_inch_second_squared := a.convertFromBase(PressurePoundPerInchSecondSquared)
	a.pounds_per_inch_second_squaredLazy = &pounds_per_inch_second_squared
	return pounds_per_inch_second_squared
}

// MetersOfWaterColumn returns the Pressure value in MetersOfWaterColumn.
//
// 
func (a *Pressure) MetersOfWaterColumn() float64 {
	if a.meters_of_water_columnLazy != nil {
		return *a.meters_of_water_columnLazy
	}
	meters_of_water_column := a.convertFromBase(PressureMeterOfWaterColumn)
	a.meters_of_water_columnLazy = &meters_of_water_column
	return meters_of_water_column
}

// InchesOfWaterColumn returns the Pressure value in InchesOfWaterColumn.
//
// 
func (a *Pressure) InchesOfWaterColumn() float64 {
	if a.inches_of_water_columnLazy != nil {
		return *a.inches_of_water_columnLazy
	}
	inches_of_water_column := a.convertFromBase(PressureInchOfWaterColumn)
	a.inches_of_water_columnLazy = &inches_of_water_column
	return inches_of_water_column
}

// MetersOfElevation returns the Pressure value in MetersOfElevation.
//
// 
func (a *Pressure) MetersOfElevation() float64 {
	if a.meters_of_elevationLazy != nil {
		return *a.meters_of_elevationLazy
	}
	meters_of_elevation := a.convertFromBase(PressureMeterOfElevation)
	a.meters_of_elevationLazy = &meters_of_elevation
	return meters_of_elevation
}

// FeetOfElevation returns the Pressure value in FeetOfElevation.
//
// 
func (a *Pressure) FeetOfElevation() float64 {
	if a.feet_of_elevationLazy != nil {
		return *a.feet_of_elevationLazy
	}
	feet_of_elevation := a.convertFromBase(PressureFootOfElevation)
	a.feet_of_elevationLazy = &feet_of_elevation
	return feet_of_elevation
}

// Micropascals returns the Pressure value in Micropascals.
//
// 
func (a *Pressure) Micropascals() float64 {
	if a.micropascalsLazy != nil {
		return *a.micropascalsLazy
	}
	micropascals := a.convertFromBase(PressureMicropascal)
	a.micropascalsLazy = &micropascals
	return micropascals
}

// Millipascals returns the Pressure value in Millipascals.
//
// 
func (a *Pressure) Millipascals() float64 {
	if a.millipascalsLazy != nil {
		return *a.millipascalsLazy
	}
	millipascals := a.convertFromBase(PressureMillipascal)
	a.millipascalsLazy = &millipascals
	return millipascals
}

// Decapascals returns the Pressure value in Decapascals.
//
// 
func (a *Pressure) Decapascals() float64 {
	if a.decapascalsLazy != nil {
		return *a.decapascalsLazy
	}
	decapascals := a.convertFromBase(PressureDecapascal)
	a.decapascalsLazy = &decapascals
	return decapascals
}

// Hectopascals returns the Pressure value in Hectopascals.
//
// 
func (a *Pressure) Hectopascals() float64 {
	if a.hectopascalsLazy != nil {
		return *a.hectopascalsLazy
	}
	hectopascals := a.convertFromBase(PressureHectopascal)
	a.hectopascalsLazy = &hectopascals
	return hectopascals
}

// Kilopascals returns the Pressure value in Kilopascals.
//
// 
func (a *Pressure) Kilopascals() float64 {
	if a.kilopascalsLazy != nil {
		return *a.kilopascalsLazy
	}
	kilopascals := a.convertFromBase(PressureKilopascal)
	a.kilopascalsLazy = &kilopascals
	return kilopascals
}

// Megapascals returns the Pressure value in Megapascals.
//
// 
func (a *Pressure) Megapascals() float64 {
	if a.megapascalsLazy != nil {
		return *a.megapascalsLazy
	}
	megapascals := a.convertFromBase(PressureMegapascal)
	a.megapascalsLazy = &megapascals
	return megapascals
}

// Gigapascals returns the Pressure value in Gigapascals.
//
// 
func (a *Pressure) Gigapascals() float64 {
	if a.gigapascalsLazy != nil {
		return *a.gigapascalsLazy
	}
	gigapascals := a.convertFromBase(PressureGigapascal)
	a.gigapascalsLazy = &gigapascals
	return gigapascals
}

// Microbars returns the Pressure value in Microbars.
//
// 
func (a *Pressure) Microbars() float64 {
	if a.microbarsLazy != nil {
		return *a.microbarsLazy
	}
	microbars := a.convertFromBase(PressureMicrobar)
	a.microbarsLazy = &microbars
	return microbars
}

// Millibars returns the Pressure value in Millibars.
//
// 
func (a *Pressure) Millibars() float64 {
	if a.millibarsLazy != nil {
		return *a.millibarsLazy
	}
	millibars := a.convertFromBase(PressureMillibar)
	a.millibarsLazy = &millibars
	return millibars
}

// Centibars returns the Pressure value in Centibars.
//
// 
func (a *Pressure) Centibars() float64 {
	if a.centibarsLazy != nil {
		return *a.centibarsLazy
	}
	centibars := a.convertFromBase(PressureCentibar)
	a.centibarsLazy = &centibars
	return centibars
}

// Decibars returns the Pressure value in Decibars.
//
// 
func (a *Pressure) Decibars() float64 {
	if a.decibarsLazy != nil {
		return *a.decibarsLazy
	}
	decibars := a.convertFromBase(PressureDecibar)
	a.decibarsLazy = &decibars
	return decibars
}

// Kilobars returns the Pressure value in Kilobars.
//
// 
func (a *Pressure) Kilobars() float64 {
	if a.kilobarsLazy != nil {
		return *a.kilobarsLazy
	}
	kilobars := a.convertFromBase(PressureKilobar)
	a.kilobarsLazy = &kilobars
	return kilobars
}

// Megabars returns the Pressure value in Megabars.
//
// 
func (a *Pressure) Megabars() float64 {
	if a.megabarsLazy != nil {
		return *a.megabarsLazy
	}
	megabars := a.convertFromBase(PressureMegabar)
	a.megabarsLazy = &megabars
	return megabars
}

// KilonewtonsPerSquareMeter returns the Pressure value in KilonewtonsPerSquareMeter.
//
// 
func (a *Pressure) KilonewtonsPerSquareMeter() float64 {
	if a.kilonewtons_per_square_meterLazy != nil {
		return *a.kilonewtons_per_square_meterLazy
	}
	kilonewtons_per_square_meter := a.convertFromBase(PressureKilonewtonPerSquareMeter)
	a.kilonewtons_per_square_meterLazy = &kilonewtons_per_square_meter
	return kilonewtons_per_square_meter
}

// MeganewtonsPerSquareMeter returns the Pressure value in MeganewtonsPerSquareMeter.
//
// 
func (a *Pressure) MeganewtonsPerSquareMeter() float64 {
	if a.meganewtons_per_square_meterLazy != nil {
		return *a.meganewtons_per_square_meterLazy
	}
	meganewtons_per_square_meter := a.convertFromBase(PressureMeganewtonPerSquareMeter)
	a.meganewtons_per_square_meterLazy = &meganewtons_per_square_meter
	return meganewtons_per_square_meter
}

// KilonewtonsPerSquareCentimeter returns the Pressure value in KilonewtonsPerSquareCentimeter.
//
// 
func (a *Pressure) KilonewtonsPerSquareCentimeter() float64 {
	if a.kilonewtons_per_square_centimeterLazy != nil {
		return *a.kilonewtons_per_square_centimeterLazy
	}
	kilonewtons_per_square_centimeter := a.convertFromBase(PressureKilonewtonPerSquareCentimeter)
	a.kilonewtons_per_square_centimeterLazy = &kilonewtons_per_square_centimeter
	return kilonewtons_per_square_centimeter
}

// KilonewtonsPerSquareMillimeter returns the Pressure value in KilonewtonsPerSquareMillimeter.
//
// 
func (a *Pressure) KilonewtonsPerSquareMillimeter() float64 {
	if a.kilonewtons_per_square_millimeterLazy != nil {
		return *a.kilonewtons_per_square_millimeterLazy
	}
	kilonewtons_per_square_millimeter := a.convertFromBase(PressureKilonewtonPerSquareMillimeter)
	a.kilonewtons_per_square_millimeterLazy = &kilonewtons_per_square_millimeter
	return kilonewtons_per_square_millimeter
}

// KilopoundsForcePerSquareInch returns the Pressure value in KilopoundsForcePerSquareInch.
//
// 
func (a *Pressure) KilopoundsForcePerSquareInch() float64 {
	if a.kilopounds_force_per_square_inchLazy != nil {
		return *a.kilopounds_force_per_square_inchLazy
	}
	kilopounds_force_per_square_inch := a.convertFromBase(PressureKilopoundForcePerSquareInch)
	a.kilopounds_force_per_square_inchLazy = &kilopounds_force_per_square_inch
	return kilopounds_force_per_square_inch
}

// KilopoundsForcePerSquareMil returns the Pressure value in KilopoundsForcePerSquareMil.
//
// 
func (a *Pressure) KilopoundsForcePerSquareMil() float64 {
	if a.kilopounds_force_per_square_milLazy != nil {
		return *a.kilopounds_force_per_square_milLazy
	}
	kilopounds_force_per_square_mil := a.convertFromBase(PressureKilopoundForcePerSquareMil)
	a.kilopounds_force_per_square_milLazy = &kilopounds_force_per_square_mil
	return kilopounds_force_per_square_mil
}

// KilopoundsForcePerSquareFoot returns the Pressure value in KilopoundsForcePerSquareFoot.
//
// 
func (a *Pressure) KilopoundsForcePerSquareFoot() float64 {
	if a.kilopounds_force_per_square_footLazy != nil {
		return *a.kilopounds_force_per_square_footLazy
	}
	kilopounds_force_per_square_foot := a.convertFromBase(PressureKilopoundForcePerSquareFoot)
	a.kilopounds_force_per_square_footLazy = &kilopounds_force_per_square_foot
	return kilopounds_force_per_square_foot
}

// MillimetersOfWaterColumn returns the Pressure value in MillimetersOfWaterColumn.
//
// 
func (a *Pressure) MillimetersOfWaterColumn() float64 {
	if a.millimeters_of_water_columnLazy != nil {
		return *a.millimeters_of_water_columnLazy
	}
	millimeters_of_water_column := a.convertFromBase(PressureMillimeterOfWaterColumn)
	a.millimeters_of_water_columnLazy = &millimeters_of_water_column
	return millimeters_of_water_column
}

// CentimetersOfWaterColumn returns the Pressure value in CentimetersOfWaterColumn.
//
// 
func (a *Pressure) CentimetersOfWaterColumn() float64 {
	if a.centimeters_of_water_columnLazy != nil {
		return *a.centimeters_of_water_columnLazy
	}
	centimeters_of_water_column := a.convertFromBase(PressureCentimeterOfWaterColumn)
	a.centimeters_of_water_columnLazy = &centimeters_of_water_column
	return centimeters_of_water_column
}


// ToDto creates a PressureDto representation from the Pressure instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Pascal by default.
func (a *Pressure) ToDto(holdInUnit *PressureUnits) PressureDto {
	if holdInUnit == nil {
		defaultUnit := PressurePascal // Default value
		holdInUnit = &defaultUnit
	}

	return PressureDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Pressure instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Pascal by default.
func (a *Pressure) ToDtoJSON(holdInUnit *PressureUnits) ([]byte, error) {
	// Convert to PressureDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Pressure to a specific unit value.
// The function uses the provided unit type (PressureUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Pressure) Convert(toUnit PressureUnits) float64 {
	switch toUnit { 
    case PressurePascal:
		return a.Pascals()
    case PressureAtmosphere:
		return a.Atmospheres()
    case PressureBar:
		return a.Bars()
    case PressureKilogramForcePerSquareMeter:
		return a.KilogramsForcePerSquareMeter()
    case PressureKilogramForcePerSquareCentimeter:
		return a.KilogramsForcePerSquareCentimeter()
    case PressureKilogramForcePerSquareMillimeter:
		return a.KilogramsForcePerSquareMillimeter()
    case PressureNewtonPerSquareMeter:
		return a.NewtonsPerSquareMeter()
    case PressureNewtonPerSquareCentimeter:
		return a.NewtonsPerSquareCentimeter()
    case PressureNewtonPerSquareMillimeter:
		return a.NewtonsPerSquareMillimeter()
    case PressureTechnicalAtmosphere:
		return a.TechnicalAtmospheres()
    case PressureTorr:
		return a.Torrs()
    case PressurePoundForcePerSquareInch:
		return a.PoundsForcePerSquareInch()
    case PressurePoundForcePerSquareMil:
		return a.PoundsForcePerSquareMil()
    case PressurePoundForcePerSquareFoot:
		return a.PoundsForcePerSquareFoot()
    case PressureTonneForcePerSquareMillimeter:
		return a.TonnesForcePerSquareMillimeter()
    case PressureTonneForcePerSquareMeter:
		return a.TonnesForcePerSquareMeter()
    case PressureMeterOfHead:
		return a.MetersOfHead()
    case PressureTonneForcePerSquareCentimeter:
		return a.TonnesForcePerSquareCentimeter()
    case PressureFootOfHead:
		return a.FeetOfHead()
    case PressureMillimeterOfMercury:
		return a.MillimetersOfMercury()
    case PressureInchOfMercury:
		return a.InchesOfMercury()
    case PressureDynePerSquareCentimeter:
		return a.DynesPerSquareCentimeter()
    case PressurePoundPerInchSecondSquared:
		return a.PoundsPerInchSecondSquared()
    case PressureMeterOfWaterColumn:
		return a.MetersOfWaterColumn()
    case PressureInchOfWaterColumn:
		return a.InchesOfWaterColumn()
    case PressureMeterOfElevation:
		return a.MetersOfElevation()
    case PressureFootOfElevation:
		return a.FeetOfElevation()
    case PressureMicropascal:
		return a.Micropascals()
    case PressureMillipascal:
		return a.Millipascals()
    case PressureDecapascal:
		return a.Decapascals()
    case PressureHectopascal:
		return a.Hectopascals()
    case PressureKilopascal:
		return a.Kilopascals()
    case PressureMegapascal:
		return a.Megapascals()
    case PressureGigapascal:
		return a.Gigapascals()
    case PressureMicrobar:
		return a.Microbars()
    case PressureMillibar:
		return a.Millibars()
    case PressureCentibar:
		return a.Centibars()
    case PressureDecibar:
		return a.Decibars()
    case PressureKilobar:
		return a.Kilobars()
    case PressureMegabar:
		return a.Megabars()
    case PressureKilonewtonPerSquareMeter:
		return a.KilonewtonsPerSquareMeter()
    case PressureMeganewtonPerSquareMeter:
		return a.MeganewtonsPerSquareMeter()
    case PressureKilonewtonPerSquareCentimeter:
		return a.KilonewtonsPerSquareCentimeter()
    case PressureKilonewtonPerSquareMillimeter:
		return a.KilonewtonsPerSquareMillimeter()
    case PressureKilopoundForcePerSquareInch:
		return a.KilopoundsForcePerSquareInch()
    case PressureKilopoundForcePerSquareMil:
		return a.KilopoundsForcePerSquareMil()
    case PressureKilopoundForcePerSquareFoot:
		return a.KilopoundsForcePerSquareFoot()
    case PressureMillimeterOfWaterColumn:
		return a.MillimetersOfWaterColumn()
    case PressureCentimeterOfWaterColumn:
		return a.CentimetersOfWaterColumn()
	default:
		return math.NaN()
	}
}

func (a *Pressure) convertFromBase(toUnit PressureUnits) float64 {
    value := a.value
	switch toUnit { 
	case PressurePascal:
		return (value) 
	case PressureAtmosphere:
		return (value / (1.01325 * 1e5)) 
	case PressureBar:
		return (value / 1e5) 
	case PressureKilogramForcePerSquareMeter:
		return (value * 0.101971619222242) 
	case PressureKilogramForcePerSquareCentimeter:
		return (value / 9.80665e4) 
	case PressureKilogramForcePerSquareMillimeter:
		return (value / 9.80665e6) 
	case PressureNewtonPerSquareMeter:
		return (value) 
	case PressureNewtonPerSquareCentimeter:
		return (value / 1e4) 
	case PressureNewtonPerSquareMillimeter:
		return (value / 1e6) 
	case PressureTechnicalAtmosphere:
		return (value / (9.80680592331 * 1e4)) 
	case PressureTorr:
		return (value / (1.3332266752 * 1e2)) 
	case PressurePoundForcePerSquareInch:
		return (value / 6.894757293168361e3) 
	case PressurePoundForcePerSquareMil:
		return (value / 6.894757293168361e9) 
	case PressurePoundForcePerSquareFoot:
		return (value / 4.788025898033584e1) 
	case PressureTonneForcePerSquareMillimeter:
		return (value / 9.80665e9) 
	case PressureTonneForcePerSquareMeter:
		return (value / 9.80665e3) 
	case PressureMeterOfHead:
		return (value * 0.0001019977334) 
	case PressureTonneForcePerSquareCentimeter:
		return (value / 9.80665e7) 
	case PressureFootOfHead:
		return (value * 0.000334552565551) 
	case PressureMillimeterOfMercury:
		return (value * 7.50061561302643e-3) 
	case PressureInchOfMercury:
		return (value * 2.95299830714159e-4) 
	case PressureDynePerSquareCentimeter:
		return (value / 1.0e-1) 
	case PressurePoundPerInchSecondSquared:
		return (value / 1.785796732283465e1) 
	case PressureMeterOfWaterColumn:
		return (value / 9.806650000000272e3) 
	case PressureInchOfWaterColumn:
		return (value / 249.08890833333) 
	case PressureMeterOfElevation:
		return ((1.0 - math.Pow(value / 101325.0, 0.190284)) * 44307.69396) 
	case PressureFootOfElevation:
		return ((1.0 - math.Pow(value / 101325.0, 0.190284)) * 145366.45) 
	case PressureMicropascal:
		return ((value) / 1e-06) 
	case PressureMillipascal:
		return ((value) / 0.001) 
	case PressureDecapascal:
		return ((value) / 10.0) 
	case PressureHectopascal:
		return ((value) / 100.0) 
	case PressureKilopascal:
		return ((value) / 1000.0) 
	case PressureMegapascal:
		return ((value) / 1000000.0) 
	case PressureGigapascal:
		return ((value) / 1000000000.0) 
	case PressureMicrobar:
		return ((value / 1e5) / 1e-06) 
	case PressureMillibar:
		return ((value / 1e5) / 0.001) 
	case PressureCentibar:
		return ((value / 1e5) / 0.01) 
	case PressureDecibar:
		return ((value / 1e5) / 0.1) 
	case PressureKilobar:
		return ((value / 1e5) / 1000.0) 
	case PressureMegabar:
		return ((value / 1e5) / 1000000.0) 
	case PressureKilonewtonPerSquareMeter:
		return ((value) / 1000.0) 
	case PressureMeganewtonPerSquareMeter:
		return ((value) / 1000000.0) 
	case PressureKilonewtonPerSquareCentimeter:
		return ((value / 1e4) / 1000.0) 
	case PressureKilonewtonPerSquareMillimeter:
		return ((value / 1e6) / 1000.0) 
	case PressureKilopoundForcePerSquareInch:
		return ((value / 6.894757293168361e3) / 1000.0) 
	case PressureKilopoundForcePerSquareMil:
		return ((value / 6.894757293168361e9) / 1000.0) 
	case PressureKilopoundForcePerSquareFoot:
		return ((value / 4.788025898033584e1) / 1000.0) 
	case PressureMillimeterOfWaterColumn:
		return ((value / 9.806650000000272e3) / 0.001) 
	case PressureCentimeterOfWaterColumn:
		return ((value / 9.806650000000272e3) / 0.01) 
	default:
		return math.NaN()
	}
}

func (a *Pressure) convertToBase(value float64, fromUnit PressureUnits) float64 {
	switch fromUnit { 
	case PressurePascal:
		return (value) 
	case PressureAtmosphere:
		return (value * 1.01325 * 1e5) 
	case PressureBar:
		return (value * 1e5) 
	case PressureKilogramForcePerSquareMeter:
		return (value * 9.80665019960652) 
	case PressureKilogramForcePerSquareCentimeter:
		return (value * 9.80665e4) 
	case PressureKilogramForcePerSquareMillimeter:
		return (value * 9.80665e6) 
	case PressureNewtonPerSquareMeter:
		return (value) 
	case PressureNewtonPerSquareCentimeter:
		return (value * 1e4) 
	case PressureNewtonPerSquareMillimeter:
		return (value * 1e6) 
	case PressureTechnicalAtmosphere:
		return (value * 9.80680592331 * 1e4) 
	case PressureTorr:
		return (value * 1.3332266752 * 1e2) 
	case PressurePoundForcePerSquareInch:
		return (value * 6.894757293168361e3) 
	case PressurePoundForcePerSquareMil:
		return (value * 6.894757293168361e9) 
	case PressurePoundForcePerSquareFoot:
		return (value * 4.788025898033584e1) 
	case PressureTonneForcePerSquareMillimeter:
		return (value * 9.80665e9) 
	case PressureTonneForcePerSquareMeter:
		return (value * 9.80665e3) 
	case PressureMeterOfHead:
		return (value * 9804.139432) 
	case PressureTonneForcePerSquareCentimeter:
		return (value * 9.80665e7) 
	case PressureFootOfHead:
		return (value * 2989.0669) 
	case PressureMillimeterOfMercury:
		return (value / 7.50061561302643e-3) 
	case PressureInchOfMercury:
		return (value / 2.95299830714159e-4) 
	case PressureDynePerSquareCentimeter:
		return (value * 1.0e-1) 
	case PressurePoundPerInchSecondSquared:
		return (value * 1.785796732283465e1) 
	case PressureMeterOfWaterColumn:
		return (value * 9.806650000000272e3) 
	case PressureInchOfWaterColumn:
		return (value * 249.08890833333) 
	case PressureMeterOfElevation:
		return (math.Pow(1.0 - (value / 44307.69396), 5.2553026003237266401799415610351) * 101325.0) 
	case PressureFootOfElevation:
		return (math.Pow(1.0 - (value / 145366.45), 5.2553026003237266401799415610351) * 101325.0) 
	case PressureMicropascal:
		return ((value) * 1e-06) 
	case PressureMillipascal:
		return ((value) * 0.001) 
	case PressureDecapascal:
		return ((value) * 10.0) 
	case PressureHectopascal:
		return ((value) * 100.0) 
	case PressureKilopascal:
		return ((value) * 1000.0) 
	case PressureMegapascal:
		return ((value) * 1000000.0) 
	case PressureGigapascal:
		return ((value) * 1000000000.0) 
	case PressureMicrobar:
		return ((value * 1e5) * 1e-06) 
	case PressureMillibar:
		return ((value * 1e5) * 0.001) 
	case PressureCentibar:
		return ((value * 1e5) * 0.01) 
	case PressureDecibar:
		return ((value * 1e5) * 0.1) 
	case PressureKilobar:
		return ((value * 1e5) * 1000.0) 
	case PressureMegabar:
		return ((value * 1e5) * 1000000.0) 
	case PressureKilonewtonPerSquareMeter:
		return ((value) * 1000.0) 
	case PressureMeganewtonPerSquareMeter:
		return ((value) * 1000000.0) 
	case PressureKilonewtonPerSquareCentimeter:
		return ((value * 1e4) * 1000.0) 
	case PressureKilonewtonPerSquareMillimeter:
		return ((value * 1e6) * 1000.0) 
	case PressureKilopoundForcePerSquareInch:
		return ((value * 6.894757293168361e3) * 1000.0) 
	case PressureKilopoundForcePerSquareMil:
		return ((value * 6.894757293168361e9) * 1000.0) 
	case PressureKilopoundForcePerSquareFoot:
		return ((value * 4.788025898033584e1) * 1000.0) 
	case PressureMillimeterOfWaterColumn:
		return ((value * 9.806650000000272e3) * 0.001) 
	case PressureCentimeterOfWaterColumn:
		return ((value * 9.806650000000272e3) * 0.01) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Pressure in the default unit (Pascal),
// formatted to two decimal places.
func (a Pressure) String() string {
	return a.ToString(PressurePascal, 2)
}

// ToString formats the Pressure value as a string with the specified unit and fractional digits.
// It converts the Pressure to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Pressure value will be converted (e.g., Pascal))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Pressure with the unit abbreviation.
func (a *Pressure) ToString(unit PressureUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetPressureAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetPressureAbbreviation(unit))
}

// Equals checks if the given Pressure is equal to the current Pressure.
//
// Parameters:
//    other: The Pressure to compare against.
//
// Returns:
//    bool: Returns true if both Pressure are equal, false otherwise.
func (a *Pressure) Equals(other *Pressure) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Pressure with another Pressure.
// It returns -1 if the current Pressure is less than the other Pressure, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Pressure to compare against.
//
// Returns:
//    int: -1 if the current Pressure is less, 1 if greater, and 0 if equal.
func (a *Pressure) CompareTo(other *Pressure) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Pressure to the current Pressure and returns the result.
// The result is a new Pressure instance with the sum of the values.
//
// Parameters:
//    other: The Pressure to add to the current Pressure.
//
// Returns:
//    *Pressure: A new Pressure instance representing the sum of both Pressure.
func (a *Pressure) Add(other *Pressure) *Pressure {
	return &Pressure{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Pressure from the current Pressure and returns the result.
// The result is a new Pressure instance with the difference of the values.
//
// Parameters:
//    other: The Pressure to subtract from the current Pressure.
//
// Returns:
//    *Pressure: A new Pressure instance representing the difference of both Pressure.
func (a *Pressure) Subtract(other *Pressure) *Pressure {
	return &Pressure{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Pressure by the given Pressure and returns the result.
// The result is a new Pressure instance with the product of the values.
//
// Parameters:
//    other: The Pressure to multiply with the current Pressure.
//
// Returns:
//    *Pressure: A new Pressure instance representing the product of both Pressure.
func (a *Pressure) Multiply(other *Pressure) *Pressure {
	return &Pressure{value: a.value * other.BaseValue()}
}

// Divide divides the current Pressure by the given Pressure and returns the result.
// The result is a new Pressure instance with the quotient of the values.
//
// Parameters:
//    other: The Pressure to divide the current Pressure by.
//
// Returns:
//    *Pressure: A new Pressure instance representing the quotient of both Pressure.
func (a *Pressure) Divide(other *Pressure) *Pressure {
	return &Pressure{value: a.value / other.BaseValue()}
}

// GetPressureAbbreviation gets the unit abbreviation.
func GetPressureAbbreviation(unit PressureUnits) string {
	switch unit { 
	case PressurePascal:
		return "Pa" 
	case PressureAtmosphere:
		return "atm" 
	case PressureBar:
		return "bar" 
	case PressureKilogramForcePerSquareMeter:
		return "kgf/m²" 
	case PressureKilogramForcePerSquareCentimeter:
		return "kgf/cm²" 
	case PressureKilogramForcePerSquareMillimeter:
		return "kgf/mm²" 
	case PressureNewtonPerSquareMeter:
		return "N/m²" 
	case PressureNewtonPerSquareCentimeter:
		return "N/cm²" 
	case PressureNewtonPerSquareMillimeter:
		return "N/mm²" 
	case PressureTechnicalAtmosphere:
		return "at" 
	case PressureTorr:
		return "torr" 
	case PressurePoundForcePerSquareInch:
		return "psi" 
	case PressurePoundForcePerSquareMil:
		return "lb/mil²" 
	case PressurePoundForcePerSquareFoot:
		return "lb/ft²" 
	case PressureTonneForcePerSquareMillimeter:
		return "tf/mm²" 
	case PressureTonneForcePerSquareMeter:
		return "tf/m²" 
	case PressureMeterOfHead:
		return "m of head" 
	case PressureTonneForcePerSquareCentimeter:
		return "tf/cm²" 
	case PressureFootOfHead:
		return "ft of head" 
	case PressureMillimeterOfMercury:
		return "mmHg" 
	case PressureInchOfMercury:
		return "inHg" 
	case PressureDynePerSquareCentimeter:
		return "dyn/cm²" 
	case PressurePoundPerInchSecondSquared:
		return "lbm/(in·s²)" 
	case PressureMeterOfWaterColumn:
		return "mH₂O" 
	case PressureInchOfWaterColumn:
		return "inH2O" 
	case PressureMeterOfElevation:
		return "m of elevation" 
	case PressureFootOfElevation:
		return "ft of elevation" 
	case PressureMicropascal:
		return "μPa" 
	case PressureMillipascal:
		return "mPa" 
	case PressureDecapascal:
		return "daPa" 
	case PressureHectopascal:
		return "hPa" 
	case PressureKilopascal:
		return "kPa" 
	case PressureMegapascal:
		return "MPa" 
	case PressureGigapascal:
		return "GPa" 
	case PressureMicrobar:
		return "μbar" 
	case PressureMillibar:
		return "mbar" 
	case PressureCentibar:
		return "cbar" 
	case PressureDecibar:
		return "dbar" 
	case PressureKilobar:
		return "kbar" 
	case PressureMegabar:
		return "Mbar" 
	case PressureKilonewtonPerSquareMeter:
		return "kN/m²" 
	case PressureMeganewtonPerSquareMeter:
		return "MN/m²" 
	case PressureKilonewtonPerSquareCentimeter:
		return "kN/cm²" 
	case PressureKilonewtonPerSquareMillimeter:
		return "kN/mm²" 
	case PressureKilopoundForcePerSquareInch:
		return "kpsi" 
	case PressureKilopoundForcePerSquareMil:
		return "klb/mil²" 
	case PressureKilopoundForcePerSquareFoot:
		return "klb/ft²" 
	case PressureMillimeterOfWaterColumn:
		return "mmH₂O" 
	case PressureCentimeterOfWaterColumn:
		return "cmH₂O" 
	default:
		return ""
	}
}