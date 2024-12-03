package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PressureUnits enumeration
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

// PressureDto represents an Pressure
type PressureDto struct {
	Value float64
	Unit  PressureUnits
}

// PressureDtoFactory struct to group related functions
type PressureDtoFactory struct{}

func (udf PressureDtoFactory) FromJSON(data []byte) (*PressureDto, error) {
	a := PressureDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PressureDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Pressure struct
type Pressure struct {
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

// PressureFactory struct to group related functions
type PressureFactory struct{}

func (uf PressureFactory) CreatePressure(value float64, unit PressureUnits) (*Pressure, error) {
	return newPressure(value, unit)
}

func (uf PressureFactory) FromDto(dto PressureDto) (*Pressure, error) {
	return newPressure(dto.Value, dto.Unit)
}

func (uf PressureFactory) FromDtoJSON(data []byte) (*Pressure, error) {
	unitDto, err := PressureDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PressureFactory{}.FromDto(*unitDto)
}


// FromPascal creates a new Pressure instance from Pascal.
func (uf PressureFactory) FromPascals(value float64) (*Pressure, error) {
	return newPressure(value, PressurePascal)
}

// FromAtmosphere creates a new Pressure instance from Atmosphere.
func (uf PressureFactory) FromAtmospheres(value float64) (*Pressure, error) {
	return newPressure(value, PressureAtmosphere)
}

// FromBar creates a new Pressure instance from Bar.
func (uf PressureFactory) FromBars(value float64) (*Pressure, error) {
	return newPressure(value, PressureBar)
}

// FromKilogramForcePerSquareMeter creates a new Pressure instance from KilogramForcePerSquareMeter.
func (uf PressureFactory) FromKilogramsForcePerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareMeter)
}

// FromKilogramForcePerSquareCentimeter creates a new Pressure instance from KilogramForcePerSquareCentimeter.
func (uf PressureFactory) FromKilogramsForcePerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareCentimeter)
}

// FromKilogramForcePerSquareMillimeter creates a new Pressure instance from KilogramForcePerSquareMillimeter.
func (uf PressureFactory) FromKilogramsForcePerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilogramForcePerSquareMillimeter)
}

// FromNewtonPerSquareMeter creates a new Pressure instance from NewtonPerSquareMeter.
func (uf PressureFactory) FromNewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareMeter)
}

// FromNewtonPerSquareCentimeter creates a new Pressure instance from NewtonPerSquareCentimeter.
func (uf PressureFactory) FromNewtonsPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareCentimeter)
}

// FromNewtonPerSquareMillimeter creates a new Pressure instance from NewtonPerSquareMillimeter.
func (uf PressureFactory) FromNewtonsPerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureNewtonPerSquareMillimeter)
}

// FromTechnicalAtmosphere creates a new Pressure instance from TechnicalAtmosphere.
func (uf PressureFactory) FromTechnicalAtmospheres(value float64) (*Pressure, error) {
	return newPressure(value, PressureTechnicalAtmosphere)
}

// FromTorr creates a new Pressure instance from Torr.
func (uf PressureFactory) FromTorrs(value float64) (*Pressure, error) {
	return newPressure(value, PressureTorr)
}

// FromPoundForcePerSquareInch creates a new Pressure instance from PoundForcePerSquareInch.
func (uf PressureFactory) FromPoundsForcePerSquareInch(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareInch)
}

// FromPoundForcePerSquareMil creates a new Pressure instance from PoundForcePerSquareMil.
func (uf PressureFactory) FromPoundsForcePerSquareMil(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareMil)
}

// FromPoundForcePerSquareFoot creates a new Pressure instance from PoundForcePerSquareFoot.
func (uf PressureFactory) FromPoundsForcePerSquareFoot(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundForcePerSquareFoot)
}

// FromTonneForcePerSquareMillimeter creates a new Pressure instance from TonneForcePerSquareMillimeter.
func (uf PressureFactory) FromTonnesForcePerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareMillimeter)
}

// FromTonneForcePerSquareMeter creates a new Pressure instance from TonneForcePerSquareMeter.
func (uf PressureFactory) FromTonnesForcePerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareMeter)
}

// FromMeterOfHead creates a new Pressure instance from MeterOfHead.
func (uf PressureFactory) FromMetersOfHead(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfHead)
}

// FromTonneForcePerSquareCentimeter creates a new Pressure instance from TonneForcePerSquareCentimeter.
func (uf PressureFactory) FromTonnesForcePerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureTonneForcePerSquareCentimeter)
}

// FromFootOfHead creates a new Pressure instance from FootOfHead.
func (uf PressureFactory) FromFeetOfHead(value float64) (*Pressure, error) {
	return newPressure(value, PressureFootOfHead)
}

// FromMillimeterOfMercury creates a new Pressure instance from MillimeterOfMercury.
func (uf PressureFactory) FromMillimetersOfMercury(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillimeterOfMercury)
}

// FromInchOfMercury creates a new Pressure instance from InchOfMercury.
func (uf PressureFactory) FromInchesOfMercury(value float64) (*Pressure, error) {
	return newPressure(value, PressureInchOfMercury)
}

// FromDynePerSquareCentimeter creates a new Pressure instance from DynePerSquareCentimeter.
func (uf PressureFactory) FromDynesPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureDynePerSquareCentimeter)
}

// FromPoundPerInchSecondSquared creates a new Pressure instance from PoundPerInchSecondSquared.
func (uf PressureFactory) FromPoundsPerInchSecondSquared(value float64) (*Pressure, error) {
	return newPressure(value, PressurePoundPerInchSecondSquared)
}

// FromMeterOfWaterColumn creates a new Pressure instance from MeterOfWaterColumn.
func (uf PressureFactory) FromMetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfWaterColumn)
}

// FromInchOfWaterColumn creates a new Pressure instance from InchOfWaterColumn.
func (uf PressureFactory) FromInchesOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureInchOfWaterColumn)
}

// FromMeterOfElevation creates a new Pressure instance from MeterOfElevation.
func (uf PressureFactory) FromMetersOfElevation(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeterOfElevation)
}

// FromFootOfElevation creates a new Pressure instance from FootOfElevation.
func (uf PressureFactory) FromFeetOfElevation(value float64) (*Pressure, error) {
	return newPressure(value, PressureFootOfElevation)
}

// FromMicropascal creates a new Pressure instance from Micropascal.
func (uf PressureFactory) FromMicropascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMicropascal)
}

// FromMillipascal creates a new Pressure instance from Millipascal.
func (uf PressureFactory) FromMillipascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillipascal)
}

// FromDecapascal creates a new Pressure instance from Decapascal.
func (uf PressureFactory) FromDecapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureDecapascal)
}

// FromHectopascal creates a new Pressure instance from Hectopascal.
func (uf PressureFactory) FromHectopascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureHectopascal)
}

// FromKilopascal creates a new Pressure instance from Kilopascal.
func (uf PressureFactory) FromKilopascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopascal)
}

// FromMegapascal creates a new Pressure instance from Megapascal.
func (uf PressureFactory) FromMegapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureMegapascal)
}

// FromGigapascal creates a new Pressure instance from Gigapascal.
func (uf PressureFactory) FromGigapascals(value float64) (*Pressure, error) {
	return newPressure(value, PressureGigapascal)
}

// FromMicrobar creates a new Pressure instance from Microbar.
func (uf PressureFactory) FromMicrobars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMicrobar)
}

// FromMillibar creates a new Pressure instance from Millibar.
func (uf PressureFactory) FromMillibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillibar)
}

// FromCentibar creates a new Pressure instance from Centibar.
func (uf PressureFactory) FromCentibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureCentibar)
}

// FromDecibar creates a new Pressure instance from Decibar.
func (uf PressureFactory) FromDecibars(value float64) (*Pressure, error) {
	return newPressure(value, PressureDecibar)
}

// FromKilobar creates a new Pressure instance from Kilobar.
func (uf PressureFactory) FromKilobars(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilobar)
}

// FromMegabar creates a new Pressure instance from Megabar.
func (uf PressureFactory) FromMegabars(value float64) (*Pressure, error) {
	return newPressure(value, PressureMegabar)
}

// FromKilonewtonPerSquareMeter creates a new Pressure instance from KilonewtonPerSquareMeter.
func (uf PressureFactory) FromKilonewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareMeter)
}

// FromMeganewtonPerSquareMeter creates a new Pressure instance from MeganewtonPerSquareMeter.
func (uf PressureFactory) FromMeganewtonsPerSquareMeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureMeganewtonPerSquareMeter)
}

// FromKilonewtonPerSquareCentimeter creates a new Pressure instance from KilonewtonPerSquareCentimeter.
func (uf PressureFactory) FromKilonewtonsPerSquareCentimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareCentimeter)
}

// FromKilonewtonPerSquareMillimeter creates a new Pressure instance from KilonewtonPerSquareMillimeter.
func (uf PressureFactory) FromKilonewtonsPerSquareMillimeter(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilonewtonPerSquareMillimeter)
}

// FromKilopoundForcePerSquareInch creates a new Pressure instance from KilopoundForcePerSquareInch.
func (uf PressureFactory) FromKilopoundsForcePerSquareInch(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareInch)
}

// FromKilopoundForcePerSquareMil creates a new Pressure instance from KilopoundForcePerSquareMil.
func (uf PressureFactory) FromKilopoundsForcePerSquareMil(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareMil)
}

// FromKilopoundForcePerSquareFoot creates a new Pressure instance from KilopoundForcePerSquareFoot.
func (uf PressureFactory) FromKilopoundsForcePerSquareFoot(value float64) (*Pressure, error) {
	return newPressure(value, PressureKilopoundForcePerSquareFoot)
}

// FromMillimeterOfWaterColumn creates a new Pressure instance from MillimeterOfWaterColumn.
func (uf PressureFactory) FromMillimetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureMillimeterOfWaterColumn)
}

// FromCentimeterOfWaterColumn creates a new Pressure instance from CentimeterOfWaterColumn.
func (uf PressureFactory) FromCentimetersOfWaterColumn(value float64) (*Pressure, error) {
	return newPressure(value, PressureCentimeterOfWaterColumn)
}




// newPressure creates a new Pressure.
func newPressure(value float64, fromUnit PressureUnits) (*Pressure, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Pressure{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Pressure in Pascal.
func (a *Pressure) BaseValue() float64 {
	return a.value
}


// Pascal returns the value in Pascal.
func (a *Pressure) Pascals() float64 {
	if a.pascalsLazy != nil {
		return *a.pascalsLazy
	}
	pascals := a.convertFromBase(PressurePascal)
	a.pascalsLazy = &pascals
	return pascals
}

// Atmosphere returns the value in Atmosphere.
func (a *Pressure) Atmospheres() float64 {
	if a.atmospheresLazy != nil {
		return *a.atmospheresLazy
	}
	atmospheres := a.convertFromBase(PressureAtmosphere)
	a.atmospheresLazy = &atmospheres
	return atmospheres
}

// Bar returns the value in Bar.
func (a *Pressure) Bars() float64 {
	if a.barsLazy != nil {
		return *a.barsLazy
	}
	bars := a.convertFromBase(PressureBar)
	a.barsLazy = &bars
	return bars
}

// KilogramForcePerSquareMeter returns the value in KilogramForcePerSquareMeter.
func (a *Pressure) KilogramsForcePerSquareMeter() float64 {
	if a.kilograms_force_per_square_meterLazy != nil {
		return *a.kilograms_force_per_square_meterLazy
	}
	kilograms_force_per_square_meter := a.convertFromBase(PressureKilogramForcePerSquareMeter)
	a.kilograms_force_per_square_meterLazy = &kilograms_force_per_square_meter
	return kilograms_force_per_square_meter
}

// KilogramForcePerSquareCentimeter returns the value in KilogramForcePerSquareCentimeter.
func (a *Pressure) KilogramsForcePerSquareCentimeter() float64 {
	if a.kilograms_force_per_square_centimeterLazy != nil {
		return *a.kilograms_force_per_square_centimeterLazy
	}
	kilograms_force_per_square_centimeter := a.convertFromBase(PressureKilogramForcePerSquareCentimeter)
	a.kilograms_force_per_square_centimeterLazy = &kilograms_force_per_square_centimeter
	return kilograms_force_per_square_centimeter
}

// KilogramForcePerSquareMillimeter returns the value in KilogramForcePerSquareMillimeter.
func (a *Pressure) KilogramsForcePerSquareMillimeter() float64 {
	if a.kilograms_force_per_square_millimeterLazy != nil {
		return *a.kilograms_force_per_square_millimeterLazy
	}
	kilograms_force_per_square_millimeter := a.convertFromBase(PressureKilogramForcePerSquareMillimeter)
	a.kilograms_force_per_square_millimeterLazy = &kilograms_force_per_square_millimeter
	return kilograms_force_per_square_millimeter
}

// NewtonPerSquareMeter returns the value in NewtonPerSquareMeter.
func (a *Pressure) NewtonsPerSquareMeter() float64 {
	if a.newtons_per_square_meterLazy != nil {
		return *a.newtons_per_square_meterLazy
	}
	newtons_per_square_meter := a.convertFromBase(PressureNewtonPerSquareMeter)
	a.newtons_per_square_meterLazy = &newtons_per_square_meter
	return newtons_per_square_meter
}

// NewtonPerSquareCentimeter returns the value in NewtonPerSquareCentimeter.
func (a *Pressure) NewtonsPerSquareCentimeter() float64 {
	if a.newtons_per_square_centimeterLazy != nil {
		return *a.newtons_per_square_centimeterLazy
	}
	newtons_per_square_centimeter := a.convertFromBase(PressureNewtonPerSquareCentimeter)
	a.newtons_per_square_centimeterLazy = &newtons_per_square_centimeter
	return newtons_per_square_centimeter
}

// NewtonPerSquareMillimeter returns the value in NewtonPerSquareMillimeter.
func (a *Pressure) NewtonsPerSquareMillimeter() float64 {
	if a.newtons_per_square_millimeterLazy != nil {
		return *a.newtons_per_square_millimeterLazy
	}
	newtons_per_square_millimeter := a.convertFromBase(PressureNewtonPerSquareMillimeter)
	a.newtons_per_square_millimeterLazy = &newtons_per_square_millimeter
	return newtons_per_square_millimeter
}

// TechnicalAtmosphere returns the value in TechnicalAtmosphere.
func (a *Pressure) TechnicalAtmospheres() float64 {
	if a.technical_atmospheresLazy != nil {
		return *a.technical_atmospheresLazy
	}
	technical_atmospheres := a.convertFromBase(PressureTechnicalAtmosphere)
	a.technical_atmospheresLazy = &technical_atmospheres
	return technical_atmospheres
}

// Torr returns the value in Torr.
func (a *Pressure) Torrs() float64 {
	if a.torrsLazy != nil {
		return *a.torrsLazy
	}
	torrs := a.convertFromBase(PressureTorr)
	a.torrsLazy = &torrs
	return torrs
}

// PoundForcePerSquareInch returns the value in PoundForcePerSquareInch.
func (a *Pressure) PoundsForcePerSquareInch() float64 {
	if a.pounds_force_per_square_inchLazy != nil {
		return *a.pounds_force_per_square_inchLazy
	}
	pounds_force_per_square_inch := a.convertFromBase(PressurePoundForcePerSquareInch)
	a.pounds_force_per_square_inchLazy = &pounds_force_per_square_inch
	return pounds_force_per_square_inch
}

// PoundForcePerSquareMil returns the value in PoundForcePerSquareMil.
func (a *Pressure) PoundsForcePerSquareMil() float64 {
	if a.pounds_force_per_square_milLazy != nil {
		return *a.pounds_force_per_square_milLazy
	}
	pounds_force_per_square_mil := a.convertFromBase(PressurePoundForcePerSquareMil)
	a.pounds_force_per_square_milLazy = &pounds_force_per_square_mil
	return pounds_force_per_square_mil
}

// PoundForcePerSquareFoot returns the value in PoundForcePerSquareFoot.
func (a *Pressure) PoundsForcePerSquareFoot() float64 {
	if a.pounds_force_per_square_footLazy != nil {
		return *a.pounds_force_per_square_footLazy
	}
	pounds_force_per_square_foot := a.convertFromBase(PressurePoundForcePerSquareFoot)
	a.pounds_force_per_square_footLazy = &pounds_force_per_square_foot
	return pounds_force_per_square_foot
}

// TonneForcePerSquareMillimeter returns the value in TonneForcePerSquareMillimeter.
func (a *Pressure) TonnesForcePerSquareMillimeter() float64 {
	if a.tonnes_force_per_square_millimeterLazy != nil {
		return *a.tonnes_force_per_square_millimeterLazy
	}
	tonnes_force_per_square_millimeter := a.convertFromBase(PressureTonneForcePerSquareMillimeter)
	a.tonnes_force_per_square_millimeterLazy = &tonnes_force_per_square_millimeter
	return tonnes_force_per_square_millimeter
}

// TonneForcePerSquareMeter returns the value in TonneForcePerSquareMeter.
func (a *Pressure) TonnesForcePerSquareMeter() float64 {
	if a.tonnes_force_per_square_meterLazy != nil {
		return *a.tonnes_force_per_square_meterLazy
	}
	tonnes_force_per_square_meter := a.convertFromBase(PressureTonneForcePerSquareMeter)
	a.tonnes_force_per_square_meterLazy = &tonnes_force_per_square_meter
	return tonnes_force_per_square_meter
}

// MeterOfHead returns the value in MeterOfHead.
func (a *Pressure) MetersOfHead() float64 {
	if a.meters_of_headLazy != nil {
		return *a.meters_of_headLazy
	}
	meters_of_head := a.convertFromBase(PressureMeterOfHead)
	a.meters_of_headLazy = &meters_of_head
	return meters_of_head
}

// TonneForcePerSquareCentimeter returns the value in TonneForcePerSquareCentimeter.
func (a *Pressure) TonnesForcePerSquareCentimeter() float64 {
	if a.tonnes_force_per_square_centimeterLazy != nil {
		return *a.tonnes_force_per_square_centimeterLazy
	}
	tonnes_force_per_square_centimeter := a.convertFromBase(PressureTonneForcePerSquareCentimeter)
	a.tonnes_force_per_square_centimeterLazy = &tonnes_force_per_square_centimeter
	return tonnes_force_per_square_centimeter
}

// FootOfHead returns the value in FootOfHead.
func (a *Pressure) FeetOfHead() float64 {
	if a.feet_of_headLazy != nil {
		return *a.feet_of_headLazy
	}
	feet_of_head := a.convertFromBase(PressureFootOfHead)
	a.feet_of_headLazy = &feet_of_head
	return feet_of_head
}

// MillimeterOfMercury returns the value in MillimeterOfMercury.
func (a *Pressure) MillimetersOfMercury() float64 {
	if a.millimeters_of_mercuryLazy != nil {
		return *a.millimeters_of_mercuryLazy
	}
	millimeters_of_mercury := a.convertFromBase(PressureMillimeterOfMercury)
	a.millimeters_of_mercuryLazy = &millimeters_of_mercury
	return millimeters_of_mercury
}

// InchOfMercury returns the value in InchOfMercury.
func (a *Pressure) InchesOfMercury() float64 {
	if a.inches_of_mercuryLazy != nil {
		return *a.inches_of_mercuryLazy
	}
	inches_of_mercury := a.convertFromBase(PressureInchOfMercury)
	a.inches_of_mercuryLazy = &inches_of_mercury
	return inches_of_mercury
}

// DynePerSquareCentimeter returns the value in DynePerSquareCentimeter.
func (a *Pressure) DynesPerSquareCentimeter() float64 {
	if a.dynes_per_square_centimeterLazy != nil {
		return *a.dynes_per_square_centimeterLazy
	}
	dynes_per_square_centimeter := a.convertFromBase(PressureDynePerSquareCentimeter)
	a.dynes_per_square_centimeterLazy = &dynes_per_square_centimeter
	return dynes_per_square_centimeter
}

// PoundPerInchSecondSquared returns the value in PoundPerInchSecondSquared.
func (a *Pressure) PoundsPerInchSecondSquared() float64 {
	if a.pounds_per_inch_second_squaredLazy != nil {
		return *a.pounds_per_inch_second_squaredLazy
	}
	pounds_per_inch_second_squared := a.convertFromBase(PressurePoundPerInchSecondSquared)
	a.pounds_per_inch_second_squaredLazy = &pounds_per_inch_second_squared
	return pounds_per_inch_second_squared
}

// MeterOfWaterColumn returns the value in MeterOfWaterColumn.
func (a *Pressure) MetersOfWaterColumn() float64 {
	if a.meters_of_water_columnLazy != nil {
		return *a.meters_of_water_columnLazy
	}
	meters_of_water_column := a.convertFromBase(PressureMeterOfWaterColumn)
	a.meters_of_water_columnLazy = &meters_of_water_column
	return meters_of_water_column
}

// InchOfWaterColumn returns the value in InchOfWaterColumn.
func (a *Pressure) InchesOfWaterColumn() float64 {
	if a.inches_of_water_columnLazy != nil {
		return *a.inches_of_water_columnLazy
	}
	inches_of_water_column := a.convertFromBase(PressureInchOfWaterColumn)
	a.inches_of_water_columnLazy = &inches_of_water_column
	return inches_of_water_column
}

// MeterOfElevation returns the value in MeterOfElevation.
func (a *Pressure) MetersOfElevation() float64 {
	if a.meters_of_elevationLazy != nil {
		return *a.meters_of_elevationLazy
	}
	meters_of_elevation := a.convertFromBase(PressureMeterOfElevation)
	a.meters_of_elevationLazy = &meters_of_elevation
	return meters_of_elevation
}

// FootOfElevation returns the value in FootOfElevation.
func (a *Pressure) FeetOfElevation() float64 {
	if a.feet_of_elevationLazy != nil {
		return *a.feet_of_elevationLazy
	}
	feet_of_elevation := a.convertFromBase(PressureFootOfElevation)
	a.feet_of_elevationLazy = &feet_of_elevation
	return feet_of_elevation
}

// Micropascal returns the value in Micropascal.
func (a *Pressure) Micropascals() float64 {
	if a.micropascalsLazy != nil {
		return *a.micropascalsLazy
	}
	micropascals := a.convertFromBase(PressureMicropascal)
	a.micropascalsLazy = &micropascals
	return micropascals
}

// Millipascal returns the value in Millipascal.
func (a *Pressure) Millipascals() float64 {
	if a.millipascalsLazy != nil {
		return *a.millipascalsLazy
	}
	millipascals := a.convertFromBase(PressureMillipascal)
	a.millipascalsLazy = &millipascals
	return millipascals
}

// Decapascal returns the value in Decapascal.
func (a *Pressure) Decapascals() float64 {
	if a.decapascalsLazy != nil {
		return *a.decapascalsLazy
	}
	decapascals := a.convertFromBase(PressureDecapascal)
	a.decapascalsLazy = &decapascals
	return decapascals
}

// Hectopascal returns the value in Hectopascal.
func (a *Pressure) Hectopascals() float64 {
	if a.hectopascalsLazy != nil {
		return *a.hectopascalsLazy
	}
	hectopascals := a.convertFromBase(PressureHectopascal)
	a.hectopascalsLazy = &hectopascals
	return hectopascals
}

// Kilopascal returns the value in Kilopascal.
func (a *Pressure) Kilopascals() float64 {
	if a.kilopascalsLazy != nil {
		return *a.kilopascalsLazy
	}
	kilopascals := a.convertFromBase(PressureKilopascal)
	a.kilopascalsLazy = &kilopascals
	return kilopascals
}

// Megapascal returns the value in Megapascal.
func (a *Pressure) Megapascals() float64 {
	if a.megapascalsLazy != nil {
		return *a.megapascalsLazy
	}
	megapascals := a.convertFromBase(PressureMegapascal)
	a.megapascalsLazy = &megapascals
	return megapascals
}

// Gigapascal returns the value in Gigapascal.
func (a *Pressure) Gigapascals() float64 {
	if a.gigapascalsLazy != nil {
		return *a.gigapascalsLazy
	}
	gigapascals := a.convertFromBase(PressureGigapascal)
	a.gigapascalsLazy = &gigapascals
	return gigapascals
}

// Microbar returns the value in Microbar.
func (a *Pressure) Microbars() float64 {
	if a.microbarsLazy != nil {
		return *a.microbarsLazy
	}
	microbars := a.convertFromBase(PressureMicrobar)
	a.microbarsLazy = &microbars
	return microbars
}

// Millibar returns the value in Millibar.
func (a *Pressure) Millibars() float64 {
	if a.millibarsLazy != nil {
		return *a.millibarsLazy
	}
	millibars := a.convertFromBase(PressureMillibar)
	a.millibarsLazy = &millibars
	return millibars
}

// Centibar returns the value in Centibar.
func (a *Pressure) Centibars() float64 {
	if a.centibarsLazy != nil {
		return *a.centibarsLazy
	}
	centibars := a.convertFromBase(PressureCentibar)
	a.centibarsLazy = &centibars
	return centibars
}

// Decibar returns the value in Decibar.
func (a *Pressure) Decibars() float64 {
	if a.decibarsLazy != nil {
		return *a.decibarsLazy
	}
	decibars := a.convertFromBase(PressureDecibar)
	a.decibarsLazy = &decibars
	return decibars
}

// Kilobar returns the value in Kilobar.
func (a *Pressure) Kilobars() float64 {
	if a.kilobarsLazy != nil {
		return *a.kilobarsLazy
	}
	kilobars := a.convertFromBase(PressureKilobar)
	a.kilobarsLazy = &kilobars
	return kilobars
}

// Megabar returns the value in Megabar.
func (a *Pressure) Megabars() float64 {
	if a.megabarsLazy != nil {
		return *a.megabarsLazy
	}
	megabars := a.convertFromBase(PressureMegabar)
	a.megabarsLazy = &megabars
	return megabars
}

// KilonewtonPerSquareMeter returns the value in KilonewtonPerSquareMeter.
func (a *Pressure) KilonewtonsPerSquareMeter() float64 {
	if a.kilonewtons_per_square_meterLazy != nil {
		return *a.kilonewtons_per_square_meterLazy
	}
	kilonewtons_per_square_meter := a.convertFromBase(PressureKilonewtonPerSquareMeter)
	a.kilonewtons_per_square_meterLazy = &kilonewtons_per_square_meter
	return kilonewtons_per_square_meter
}

// MeganewtonPerSquareMeter returns the value in MeganewtonPerSquareMeter.
func (a *Pressure) MeganewtonsPerSquareMeter() float64 {
	if a.meganewtons_per_square_meterLazy != nil {
		return *a.meganewtons_per_square_meterLazy
	}
	meganewtons_per_square_meter := a.convertFromBase(PressureMeganewtonPerSquareMeter)
	a.meganewtons_per_square_meterLazy = &meganewtons_per_square_meter
	return meganewtons_per_square_meter
}

// KilonewtonPerSquareCentimeter returns the value in KilonewtonPerSquareCentimeter.
func (a *Pressure) KilonewtonsPerSquareCentimeter() float64 {
	if a.kilonewtons_per_square_centimeterLazy != nil {
		return *a.kilonewtons_per_square_centimeterLazy
	}
	kilonewtons_per_square_centimeter := a.convertFromBase(PressureKilonewtonPerSquareCentimeter)
	a.kilonewtons_per_square_centimeterLazy = &kilonewtons_per_square_centimeter
	return kilonewtons_per_square_centimeter
}

// KilonewtonPerSquareMillimeter returns the value in KilonewtonPerSquareMillimeter.
func (a *Pressure) KilonewtonsPerSquareMillimeter() float64 {
	if a.kilonewtons_per_square_millimeterLazy != nil {
		return *a.kilonewtons_per_square_millimeterLazy
	}
	kilonewtons_per_square_millimeter := a.convertFromBase(PressureKilonewtonPerSquareMillimeter)
	a.kilonewtons_per_square_millimeterLazy = &kilonewtons_per_square_millimeter
	return kilonewtons_per_square_millimeter
}

// KilopoundForcePerSquareInch returns the value in KilopoundForcePerSquareInch.
func (a *Pressure) KilopoundsForcePerSquareInch() float64 {
	if a.kilopounds_force_per_square_inchLazy != nil {
		return *a.kilopounds_force_per_square_inchLazy
	}
	kilopounds_force_per_square_inch := a.convertFromBase(PressureKilopoundForcePerSquareInch)
	a.kilopounds_force_per_square_inchLazy = &kilopounds_force_per_square_inch
	return kilopounds_force_per_square_inch
}

// KilopoundForcePerSquareMil returns the value in KilopoundForcePerSquareMil.
func (a *Pressure) KilopoundsForcePerSquareMil() float64 {
	if a.kilopounds_force_per_square_milLazy != nil {
		return *a.kilopounds_force_per_square_milLazy
	}
	kilopounds_force_per_square_mil := a.convertFromBase(PressureKilopoundForcePerSquareMil)
	a.kilopounds_force_per_square_milLazy = &kilopounds_force_per_square_mil
	return kilopounds_force_per_square_mil
}

// KilopoundForcePerSquareFoot returns the value in KilopoundForcePerSquareFoot.
func (a *Pressure) KilopoundsForcePerSquareFoot() float64 {
	if a.kilopounds_force_per_square_footLazy != nil {
		return *a.kilopounds_force_per_square_footLazy
	}
	kilopounds_force_per_square_foot := a.convertFromBase(PressureKilopoundForcePerSquareFoot)
	a.kilopounds_force_per_square_footLazy = &kilopounds_force_per_square_foot
	return kilopounds_force_per_square_foot
}

// MillimeterOfWaterColumn returns the value in MillimeterOfWaterColumn.
func (a *Pressure) MillimetersOfWaterColumn() float64 {
	if a.millimeters_of_water_columnLazy != nil {
		return *a.millimeters_of_water_columnLazy
	}
	millimeters_of_water_column := a.convertFromBase(PressureMillimeterOfWaterColumn)
	a.millimeters_of_water_columnLazy = &millimeters_of_water_column
	return millimeters_of_water_column
}

// CentimeterOfWaterColumn returns the value in CentimeterOfWaterColumn.
func (a *Pressure) CentimetersOfWaterColumn() float64 {
	if a.centimeters_of_water_columnLazy != nil {
		return *a.centimeters_of_water_columnLazy
	}
	centimeters_of_water_column := a.convertFromBase(PressureCentimeterOfWaterColumn)
	a.centimeters_of_water_columnLazy = &centimeters_of_water_column
	return centimeters_of_water_column
}


// ToDto creates an PressureDto representation.
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

// ToDtoJSON creates an PressureDto representation.
func (a *Pressure) ToDtoJSON(holdInUnit *PressureUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Pressure to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Pressure) String() string {
	return a.ToString(PressurePascal, 2)
}

// ToString formats the Pressure to string.
// fractionalDigits -1 for not mention
func (a *Pressure) ToString(unit PressureUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Pressure) getUnitAbbreviation(unit PressureUnits) string {
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

// Check if the given Pressure are equals to the current Pressure
func (a *Pressure) Equals(other *Pressure) bool {
	return a.value == other.BaseValue()
}

// Check if the given Pressure are equals to the current Pressure
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

// Add the given Pressure to the current Pressure.
func (a *Pressure) Add(other *Pressure) *Pressure {
	return &Pressure{value: a.value + other.BaseValue()}
}

// Subtract the given Pressure to the current Pressure.
func (a *Pressure) Subtract(other *Pressure) *Pressure {
	return &Pressure{value: a.value - other.BaseValue()}
}

// Multiply the given Pressure to the current Pressure.
func (a *Pressure) Multiply(other *Pressure) *Pressure {
	return &Pressure{value: a.value * other.BaseValue()}
}

// Divide the given Pressure to the current Pressure.
func (a *Pressure) Divide(other *Pressure) *Pressure {
	return &Pressure{value: a.value / other.BaseValue()}
}