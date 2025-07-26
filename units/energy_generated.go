// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// EnergyUnits defines various units of Energy.
type EnergyUnits string

const (
    
        // 
        EnergyJoule EnergyUnits = "Joule"
        // 
        EnergyCalorie EnergyUnits = "Calorie"
        // 
        EnergyBritishThermalUnit EnergyUnits = "BritishThermalUnit"
        // In physics, an electronvolt (symbol eV, also written electron-volt and electron volt) is the measure of an amount of kinetic energy gained by a single electron accelerating from rest through an electric potential difference of one volt in vacuum. When used as a unit of energy, the numerical value of 1 eV in joules (symbol J) is equivalent to the numerical value of the charge of an electron in coulombs (symbol C). Under the 2019 redefinition of the SI base units, this sets 1 eV equal to the exact value 1.602176634×10−19 J.
        EnergyElectronVolt EnergyUnits = "ElectronVolt"
        // A pound-foot (lb⋅ft), abbreviated from pound-force foot (lbf · ft), is a unit of torque representing one pound of force acting at a perpendicular distance of one foot from a pivot point. Conversely one foot pound-force (ft · lbf) is the moment about an axis that applies one pound-force at a radius of one foot.
        EnergyFootPound EnergyUnits = "FootPound"
        // The erg is a unit of energy equal to 10−7 joules (100 nJ). It originated in the Centimetre–gram–second system of units (CGS). It has the symbol erg. The erg is not an SI unit. Its name is derived from ergon (ἔργον), a Greek word meaning 'work' or 'task'.
        EnergyErg EnergyUnits = "Erg"
        // 
        EnergyWattHour EnergyUnits = "WattHour"
        // 
        EnergyWattDay EnergyUnits = "WattDay"
        // The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
        EnergyThermEc EnergyUnits = "ThermEc"
        // The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
        EnergyThermUs EnergyUnits = "ThermUs"
        // The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
        EnergyThermImperial EnergyUnits = "ThermImperial"
        // A horsepower-hour (symbol: hp⋅h) is an outdated unit of energy, not used in the International System of Units. The unit represents an amount of work a horse is supposed capable of delivering during an hour (1 horsepower integrated over a time interval of an hour).
        EnergyHorsepowerHour EnergyUnits = "HorsepowerHour"
        // 
        EnergyNanojoule EnergyUnits = "Nanojoule"
        // 
        EnergyMicrojoule EnergyUnits = "Microjoule"
        // 
        EnergyMillijoule EnergyUnits = "Millijoule"
        // 
        EnergyKilojoule EnergyUnits = "Kilojoule"
        // 
        EnergyMegajoule EnergyUnits = "Megajoule"
        // 
        EnergyGigajoule EnergyUnits = "Gigajoule"
        // 
        EnergyTerajoule EnergyUnits = "Terajoule"
        // 
        EnergyPetajoule EnergyUnits = "Petajoule"
        // 
        EnergyKilocalorie EnergyUnits = "Kilocalorie"
        // 
        EnergyMegacalorie EnergyUnits = "Megacalorie"
        // 
        EnergyKilobritishThermalUnit EnergyUnits = "KilobritishThermalUnit"
        // 
        EnergyMegabritishThermalUnit EnergyUnits = "MegabritishThermalUnit"
        // 
        EnergyGigabritishThermalUnit EnergyUnits = "GigabritishThermalUnit"
        // 
        EnergyKiloelectronVolt EnergyUnits = "KiloelectronVolt"
        // 
        EnergyMegaelectronVolt EnergyUnits = "MegaelectronVolt"
        // 
        EnergyGigaelectronVolt EnergyUnits = "GigaelectronVolt"
        // 
        EnergyTeraelectronVolt EnergyUnits = "TeraelectronVolt"
        // 
        EnergyKilowattHour EnergyUnits = "KilowattHour"
        // 
        EnergyMegawattHour EnergyUnits = "MegawattHour"
        // 
        EnergyGigawattHour EnergyUnits = "GigawattHour"
        // 
        EnergyTerawattHour EnergyUnits = "TerawattHour"
        // 
        EnergyKilowattDay EnergyUnits = "KilowattDay"
        // 
        EnergyMegawattDay EnergyUnits = "MegawattDay"
        // 
        EnergyGigawattDay EnergyUnits = "GigawattDay"
        // 
        EnergyTerawattDay EnergyUnits = "TerawattDay"
        // 
        EnergyDecathermEc EnergyUnits = "DecathermEc"
        // 
        EnergyDecathermUs EnergyUnits = "DecathermUs"
        // 
        EnergyDecathermImperial EnergyUnits = "DecathermImperial"
)

var internalEnergyUnitsMap = map[EnergyUnits]bool{
	
	EnergyJoule: true,
	EnergyCalorie: true,
	EnergyBritishThermalUnit: true,
	EnergyElectronVolt: true,
	EnergyFootPound: true,
	EnergyErg: true,
	EnergyWattHour: true,
	EnergyWattDay: true,
	EnergyThermEc: true,
	EnergyThermUs: true,
	EnergyThermImperial: true,
	EnergyHorsepowerHour: true,
	EnergyNanojoule: true,
	EnergyMicrojoule: true,
	EnergyMillijoule: true,
	EnergyKilojoule: true,
	EnergyMegajoule: true,
	EnergyGigajoule: true,
	EnergyTerajoule: true,
	EnergyPetajoule: true,
	EnergyKilocalorie: true,
	EnergyMegacalorie: true,
	EnergyKilobritishThermalUnit: true,
	EnergyMegabritishThermalUnit: true,
	EnergyGigabritishThermalUnit: true,
	EnergyKiloelectronVolt: true,
	EnergyMegaelectronVolt: true,
	EnergyGigaelectronVolt: true,
	EnergyTeraelectronVolt: true,
	EnergyKilowattHour: true,
	EnergyMegawattHour: true,
	EnergyGigawattHour: true,
	EnergyTerawattHour: true,
	EnergyKilowattDay: true,
	EnergyMegawattDay: true,
	EnergyGigawattDay: true,
	EnergyTerawattDay: true,
	EnergyDecathermEc: true,
	EnergyDecathermUs: true,
	EnergyDecathermImperial: true,
}

// EnergyDto represents a Energy measurement with a numerical value and its corresponding unit.
type EnergyDto struct {
    // Value is the numerical representation of the Energy.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Energy, as defined in the EnergyUnits enumeration.
	Unit  EnergyUnits `json:"unit" validate:"required,oneof=Joule Calorie BritishThermalUnit ElectronVolt FootPound Erg WattHour WattDay ThermEc ThermUs ThermImperial HorsepowerHour Nanojoule Microjoule Millijoule Kilojoule Megajoule Gigajoule Terajoule Petajoule Kilocalorie Megacalorie KilobritishThermalUnit MegabritishThermalUnit GigabritishThermalUnit KiloelectronVolt MegaelectronVolt GigaelectronVolt TeraelectronVolt KilowattHour MegawattHour GigawattHour TerawattHour KilowattDay MegawattDay GigawattDay TerawattDay DecathermEc DecathermUs DecathermImperial"`
}

// EnergyDtoFactory groups methods for creating and serializing EnergyDto objects.
type EnergyDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a EnergyDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf EnergyDtoFactory) FromJSON(data []byte) (*EnergyDto, error) {
	a := EnergyDto{}

    // Parse JSON into EnergyDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a EnergyDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a EnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Energy represents a measurement in a of Energy.
//
// The joule, symbol J, is a derived unit of energy, work, or amount of heat in the International System of Units. It is equal to the energy transferred (or work done) when applying a force of one newton through a distance of one metre (1 newton metre or N·m), or in passing an electric current of one ampere through a resistance of one ohm for one second. Many other units of energy are included. Please do not confuse this definition of the calorie with the one colloquially used by the food industry, the large calorie, which is equivalent to 1 kcal. Thermochemical definition of the calorie is used. For BTU, the IT definition is used.
type Energy struct {
	// value is the base measurement stored internally.
	value       float64
    
    joulesLazy *float64 
    caloriesLazy *float64 
    british_thermal_unitsLazy *float64 
    electron_voltsLazy *float64 
    foot_poundsLazy *float64 
    ergsLazy *float64 
    watt_hoursLazy *float64 
    watt_daysLazy *float64 
    therms_ecLazy *float64 
    therms_usLazy *float64 
    therms_imperialLazy *float64 
    horsepower_hoursLazy *float64 
    nanojoulesLazy *float64 
    microjoulesLazy *float64 
    millijoulesLazy *float64 
    kilojoulesLazy *float64 
    megajoulesLazy *float64 
    gigajoulesLazy *float64 
    terajoulesLazy *float64 
    petajoulesLazy *float64 
    kilocaloriesLazy *float64 
    megacaloriesLazy *float64 
    kilobritish_thermal_unitsLazy *float64 
    megabritish_thermal_unitsLazy *float64 
    gigabritish_thermal_unitsLazy *float64 
    kiloelectron_voltsLazy *float64 
    megaelectron_voltsLazy *float64 
    gigaelectron_voltsLazy *float64 
    teraelectron_voltsLazy *float64 
    kilowatt_hoursLazy *float64 
    megawatt_hoursLazy *float64 
    gigawatt_hoursLazy *float64 
    terawatt_hoursLazy *float64 
    kilowatt_daysLazy *float64 
    megawatt_daysLazy *float64 
    gigawatt_daysLazy *float64 
    terawatt_daysLazy *float64 
    decatherms_ecLazy *float64 
    decatherms_usLazy *float64 
    decatherms_imperialLazy *float64 
}

// EnergyFactory groups methods for creating Energy instances.
type EnergyFactory struct{}

// CreateEnergy creates a new Energy instance from the given value and unit.
func (uf EnergyFactory) CreateEnergy(value float64, unit EnergyUnits) (*Energy, error) {
	return newEnergy(value, unit)
}

// FromDto converts a EnergyDto to a Energy instance.
func (uf EnergyFactory) FromDto(dto EnergyDto) (*Energy, error) {
	return newEnergy(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Energy instance.
func (uf EnergyFactory) FromDtoJSON(data []byte) (*Energy, error) {
	unitDto, err := EnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EnergyDto from JSON: %w", err)
	}
	return EnergyFactory{}.FromDto(*unitDto)
}


// FromJoules creates a new Energy instance from a value in Joules.
func (uf EnergyFactory) FromJoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyJoule)
}

// FromCalories creates a new Energy instance from a value in Calories.
func (uf EnergyFactory) FromCalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyCalorie)
}

// FromBritishThermalUnits creates a new Energy instance from a value in BritishThermalUnits.
func (uf EnergyFactory) FromBritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyBritishThermalUnit)
}

// FromElectronVolts creates a new Energy instance from a value in ElectronVolts.
func (uf EnergyFactory) FromElectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyElectronVolt)
}

// FromFootPounds creates a new Energy instance from a value in FootPounds.
func (uf EnergyFactory) FromFootPounds(value float64) (*Energy, error) {
	return newEnergy(value, EnergyFootPound)
}

// FromErgs creates a new Energy instance from a value in Ergs.
func (uf EnergyFactory) FromErgs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyErg)
}

// FromWattHours creates a new Energy instance from a value in WattHours.
func (uf EnergyFactory) FromWattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyWattHour)
}

// FromWattDays creates a new Energy instance from a value in WattDays.
func (uf EnergyFactory) FromWattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyWattDay)
}

// FromThermsEc creates a new Energy instance from a value in ThermsEc.
func (uf EnergyFactory) FromThermsEc(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermEc)
}

// FromThermsUs creates a new Energy instance from a value in ThermsUs.
func (uf EnergyFactory) FromThermsUs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermUs)
}

// FromThermsImperial creates a new Energy instance from a value in ThermsImperial.
func (uf EnergyFactory) FromThermsImperial(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermImperial)
}

// FromHorsepowerHours creates a new Energy instance from a value in HorsepowerHours.
func (uf EnergyFactory) FromHorsepowerHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyHorsepowerHour)
}

// FromNanojoules creates a new Energy instance from a value in Nanojoules.
func (uf EnergyFactory) FromNanojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyNanojoule)
}

// FromMicrojoules creates a new Energy instance from a value in Microjoules.
func (uf EnergyFactory) FromMicrojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMicrojoule)
}

// FromMillijoules creates a new Energy instance from a value in Millijoules.
func (uf EnergyFactory) FromMillijoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMillijoule)
}

// FromKilojoules creates a new Energy instance from a value in Kilojoules.
func (uf EnergyFactory) FromKilojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilojoule)
}

// FromMegajoules creates a new Energy instance from a value in Megajoules.
func (uf EnergyFactory) FromMegajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegajoule)
}

// FromGigajoules creates a new Energy instance from a value in Gigajoules.
func (uf EnergyFactory) FromGigajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigajoule)
}

// FromTerajoules creates a new Energy instance from a value in Terajoules.
func (uf EnergyFactory) FromTerajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerajoule)
}

// FromPetajoules creates a new Energy instance from a value in Petajoules.
func (uf EnergyFactory) FromPetajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyPetajoule)
}

// FromKilocalories creates a new Energy instance from a value in Kilocalories.
func (uf EnergyFactory) FromKilocalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilocalorie)
}

// FromMegacalories creates a new Energy instance from a value in Megacalories.
func (uf EnergyFactory) FromMegacalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegacalorie)
}

// FromKilobritishThermalUnits creates a new Energy instance from a value in KilobritishThermalUnits.
func (uf EnergyFactory) FromKilobritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilobritishThermalUnit)
}

// FromMegabritishThermalUnits creates a new Energy instance from a value in MegabritishThermalUnits.
func (uf EnergyFactory) FromMegabritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegabritishThermalUnit)
}

// FromGigabritishThermalUnits creates a new Energy instance from a value in GigabritishThermalUnits.
func (uf EnergyFactory) FromGigabritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigabritishThermalUnit)
}

// FromKiloelectronVolts creates a new Energy instance from a value in KiloelectronVolts.
func (uf EnergyFactory) FromKiloelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKiloelectronVolt)
}

// FromMegaelectronVolts creates a new Energy instance from a value in MegaelectronVolts.
func (uf EnergyFactory) FromMegaelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegaelectronVolt)
}

// FromGigaelectronVolts creates a new Energy instance from a value in GigaelectronVolts.
func (uf EnergyFactory) FromGigaelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigaelectronVolt)
}

// FromTeraelectronVolts creates a new Energy instance from a value in TeraelectronVolts.
func (uf EnergyFactory) FromTeraelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTeraelectronVolt)
}

// FromKilowattHours creates a new Energy instance from a value in KilowattHours.
func (uf EnergyFactory) FromKilowattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilowattHour)
}

// FromMegawattHours creates a new Energy instance from a value in MegawattHours.
func (uf EnergyFactory) FromMegawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegawattHour)
}

// FromGigawattHours creates a new Energy instance from a value in GigawattHours.
func (uf EnergyFactory) FromGigawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigawattHour)
}

// FromTerawattHours creates a new Energy instance from a value in TerawattHours.
func (uf EnergyFactory) FromTerawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerawattHour)
}

// FromKilowattDays creates a new Energy instance from a value in KilowattDays.
func (uf EnergyFactory) FromKilowattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilowattDay)
}

// FromMegawattDays creates a new Energy instance from a value in MegawattDays.
func (uf EnergyFactory) FromMegawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegawattDay)
}

// FromGigawattDays creates a new Energy instance from a value in GigawattDays.
func (uf EnergyFactory) FromGigawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigawattDay)
}

// FromTerawattDays creates a new Energy instance from a value in TerawattDays.
func (uf EnergyFactory) FromTerawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerawattDay)
}

// FromDecathermsEc creates a new Energy instance from a value in DecathermsEc.
func (uf EnergyFactory) FromDecathermsEc(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermEc)
}

// FromDecathermsUs creates a new Energy instance from a value in DecathermsUs.
func (uf EnergyFactory) FromDecathermsUs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermUs)
}

// FromDecathermsImperial creates a new Energy instance from a value in DecathermsImperial.
func (uf EnergyFactory) FromDecathermsImperial(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermImperial)
}


// newEnergy creates a new Energy.
func newEnergy(value float64, fromUnit EnergyUnits) (*Energy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	if _, ok := internalEnergyUnitsMap[fromUnit]; !ok {
		return nil, fmt.Errorf("unknown unit %s in EnergyUnits", fromUnit)
	}
	a := &Energy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Energy in Joule unit (the base/default quantity).
func (a *Energy) BaseValue() float64 {
	return a.value
}


// Joules returns the Energy value in Joules.
//
// 
func (a *Energy) Joules() float64 {
	if a.joulesLazy != nil {
		return *a.joulesLazy
	}
	joules := a.convertFromBase(EnergyJoule)
	a.joulesLazy = &joules
	return joules
}

// Calories returns the Energy value in Calories.
//
// 
func (a *Energy) Calories() float64 {
	if a.caloriesLazy != nil {
		return *a.caloriesLazy
	}
	calories := a.convertFromBase(EnergyCalorie)
	a.caloriesLazy = &calories
	return calories
}

// BritishThermalUnits returns the Energy value in BritishThermalUnits.
//
// 
func (a *Energy) BritishThermalUnits() float64 {
	if a.british_thermal_unitsLazy != nil {
		return *a.british_thermal_unitsLazy
	}
	british_thermal_units := a.convertFromBase(EnergyBritishThermalUnit)
	a.british_thermal_unitsLazy = &british_thermal_units
	return british_thermal_units
}

// ElectronVolts returns the Energy value in ElectronVolts.
//
// In physics, an electronvolt (symbol eV, also written electron-volt and electron volt) is the measure of an amount of kinetic energy gained by a single electron accelerating from rest through an electric potential difference of one volt in vacuum. When used as a unit of energy, the numerical value of 1 eV in joules (symbol J) is equivalent to the numerical value of the charge of an electron in coulombs (symbol C). Under the 2019 redefinition of the SI base units, this sets 1 eV equal to the exact value 1.602176634×10−19 J.
func (a *Energy) ElectronVolts() float64 {
	if a.electron_voltsLazy != nil {
		return *a.electron_voltsLazy
	}
	electron_volts := a.convertFromBase(EnergyElectronVolt)
	a.electron_voltsLazy = &electron_volts
	return electron_volts
}

// FootPounds returns the Energy value in FootPounds.
//
// A pound-foot (lb⋅ft), abbreviated from pound-force foot (lbf · ft), is a unit of torque representing one pound of force acting at a perpendicular distance of one foot from a pivot point. Conversely one foot pound-force (ft · lbf) is the moment about an axis that applies one pound-force at a radius of one foot.
func (a *Energy) FootPounds() float64 {
	if a.foot_poundsLazy != nil {
		return *a.foot_poundsLazy
	}
	foot_pounds := a.convertFromBase(EnergyFootPound)
	a.foot_poundsLazy = &foot_pounds
	return foot_pounds
}

// Ergs returns the Energy value in Ergs.
//
// The erg is a unit of energy equal to 10−7 joules (100 nJ). It originated in the Centimetre–gram–second system of units (CGS). It has the symbol erg. The erg is not an SI unit. Its name is derived from ergon (ἔργον), a Greek word meaning 'work' or 'task'.
func (a *Energy) Ergs() float64 {
	if a.ergsLazy != nil {
		return *a.ergsLazy
	}
	ergs := a.convertFromBase(EnergyErg)
	a.ergsLazy = &ergs
	return ergs
}

// WattHours returns the Energy value in WattHours.
//
// 
func (a *Energy) WattHours() float64 {
	if a.watt_hoursLazy != nil {
		return *a.watt_hoursLazy
	}
	watt_hours := a.convertFromBase(EnergyWattHour)
	a.watt_hoursLazy = &watt_hours
	return watt_hours
}

// WattDays returns the Energy value in WattDays.
//
// 
func (a *Energy) WattDays() float64 {
	if a.watt_daysLazy != nil {
		return *a.watt_daysLazy
	}
	watt_days := a.convertFromBase(EnergyWattDay)
	a.watt_daysLazy = &watt_days
	return watt_days
}

// ThermsEc returns the Energy value in ThermsEc.
//
// The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
func (a *Energy) ThermsEc() float64 {
	if a.therms_ecLazy != nil {
		return *a.therms_ecLazy
	}
	therms_ec := a.convertFromBase(EnergyThermEc)
	a.therms_ecLazy = &therms_ec
	return therms_ec
}

// ThermsUs returns the Energy value in ThermsUs.
//
// The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
func (a *Energy) ThermsUs() float64 {
	if a.therms_usLazy != nil {
		return *a.therms_usLazy
	}
	therms_us := a.convertFromBase(EnergyThermUs)
	a.therms_usLazy = &therms_us
	return therms_us
}

// ThermsImperial returns the Energy value in ThermsImperial.
//
// The therm (symbol, thm) is a non-SI unit of heat energy equal to 100,000 British thermal units (BTU), and approximately 105 megajoules, 29.3 kilowatt-hours, 25,200 kilocalories and 25.2 thermies. One therm is the energy content of approximately 100 cubic feet (2.83 cubic metres) of natural gas at standard temperature and pressure. However, the BTU is not standardised worldwide, with slightly different values in the EU, UK, and United States, meaning that the energy content of the therm also varies by territory.
func (a *Energy) ThermsImperial() float64 {
	if a.therms_imperialLazy != nil {
		return *a.therms_imperialLazy
	}
	therms_imperial := a.convertFromBase(EnergyThermImperial)
	a.therms_imperialLazy = &therms_imperial
	return therms_imperial
}

// HorsepowerHours returns the Energy value in HorsepowerHours.
//
// A horsepower-hour (symbol: hp⋅h) is an outdated unit of energy, not used in the International System of Units. The unit represents an amount of work a horse is supposed capable of delivering during an hour (1 horsepower integrated over a time interval of an hour).
func (a *Energy) HorsepowerHours() float64 {
	if a.horsepower_hoursLazy != nil {
		return *a.horsepower_hoursLazy
	}
	horsepower_hours := a.convertFromBase(EnergyHorsepowerHour)
	a.horsepower_hoursLazy = &horsepower_hours
	return horsepower_hours
}

// Nanojoules returns the Energy value in Nanojoules.
//
// 
func (a *Energy) Nanojoules() float64 {
	if a.nanojoulesLazy != nil {
		return *a.nanojoulesLazy
	}
	nanojoules := a.convertFromBase(EnergyNanojoule)
	a.nanojoulesLazy = &nanojoules
	return nanojoules
}

// Microjoules returns the Energy value in Microjoules.
//
// 
func (a *Energy) Microjoules() float64 {
	if a.microjoulesLazy != nil {
		return *a.microjoulesLazy
	}
	microjoules := a.convertFromBase(EnergyMicrojoule)
	a.microjoulesLazy = &microjoules
	return microjoules
}

// Millijoules returns the Energy value in Millijoules.
//
// 
func (a *Energy) Millijoules() float64 {
	if a.millijoulesLazy != nil {
		return *a.millijoulesLazy
	}
	millijoules := a.convertFromBase(EnergyMillijoule)
	a.millijoulesLazy = &millijoules
	return millijoules
}

// Kilojoules returns the Energy value in Kilojoules.
//
// 
func (a *Energy) Kilojoules() float64 {
	if a.kilojoulesLazy != nil {
		return *a.kilojoulesLazy
	}
	kilojoules := a.convertFromBase(EnergyKilojoule)
	a.kilojoulesLazy = &kilojoules
	return kilojoules
}

// Megajoules returns the Energy value in Megajoules.
//
// 
func (a *Energy) Megajoules() float64 {
	if a.megajoulesLazy != nil {
		return *a.megajoulesLazy
	}
	megajoules := a.convertFromBase(EnergyMegajoule)
	a.megajoulesLazy = &megajoules
	return megajoules
}

// Gigajoules returns the Energy value in Gigajoules.
//
// 
func (a *Energy) Gigajoules() float64 {
	if a.gigajoulesLazy != nil {
		return *a.gigajoulesLazy
	}
	gigajoules := a.convertFromBase(EnergyGigajoule)
	a.gigajoulesLazy = &gigajoules
	return gigajoules
}

// Terajoules returns the Energy value in Terajoules.
//
// 
func (a *Energy) Terajoules() float64 {
	if a.terajoulesLazy != nil {
		return *a.terajoulesLazy
	}
	terajoules := a.convertFromBase(EnergyTerajoule)
	a.terajoulesLazy = &terajoules
	return terajoules
}

// Petajoules returns the Energy value in Petajoules.
//
// 
func (a *Energy) Petajoules() float64 {
	if a.petajoulesLazy != nil {
		return *a.petajoulesLazy
	}
	petajoules := a.convertFromBase(EnergyPetajoule)
	a.petajoulesLazy = &petajoules
	return petajoules
}

// Kilocalories returns the Energy value in Kilocalories.
//
// 
func (a *Energy) Kilocalories() float64 {
	if a.kilocaloriesLazy != nil {
		return *a.kilocaloriesLazy
	}
	kilocalories := a.convertFromBase(EnergyKilocalorie)
	a.kilocaloriesLazy = &kilocalories
	return kilocalories
}

// Megacalories returns the Energy value in Megacalories.
//
// 
func (a *Energy) Megacalories() float64 {
	if a.megacaloriesLazy != nil {
		return *a.megacaloriesLazy
	}
	megacalories := a.convertFromBase(EnergyMegacalorie)
	a.megacaloriesLazy = &megacalories
	return megacalories
}

// KilobritishThermalUnits returns the Energy value in KilobritishThermalUnits.
//
// 
func (a *Energy) KilobritishThermalUnits() float64 {
	if a.kilobritish_thermal_unitsLazy != nil {
		return *a.kilobritish_thermal_unitsLazy
	}
	kilobritish_thermal_units := a.convertFromBase(EnergyKilobritishThermalUnit)
	a.kilobritish_thermal_unitsLazy = &kilobritish_thermal_units
	return kilobritish_thermal_units
}

// MegabritishThermalUnits returns the Energy value in MegabritishThermalUnits.
//
// 
func (a *Energy) MegabritishThermalUnits() float64 {
	if a.megabritish_thermal_unitsLazy != nil {
		return *a.megabritish_thermal_unitsLazy
	}
	megabritish_thermal_units := a.convertFromBase(EnergyMegabritishThermalUnit)
	a.megabritish_thermal_unitsLazy = &megabritish_thermal_units
	return megabritish_thermal_units
}

// GigabritishThermalUnits returns the Energy value in GigabritishThermalUnits.
//
// 
func (a *Energy) GigabritishThermalUnits() float64 {
	if a.gigabritish_thermal_unitsLazy != nil {
		return *a.gigabritish_thermal_unitsLazy
	}
	gigabritish_thermal_units := a.convertFromBase(EnergyGigabritishThermalUnit)
	a.gigabritish_thermal_unitsLazy = &gigabritish_thermal_units
	return gigabritish_thermal_units
}

// KiloelectronVolts returns the Energy value in KiloelectronVolts.
//
// 
func (a *Energy) KiloelectronVolts() float64 {
	if a.kiloelectron_voltsLazy != nil {
		return *a.kiloelectron_voltsLazy
	}
	kiloelectron_volts := a.convertFromBase(EnergyKiloelectronVolt)
	a.kiloelectron_voltsLazy = &kiloelectron_volts
	return kiloelectron_volts
}

// MegaelectronVolts returns the Energy value in MegaelectronVolts.
//
// 
func (a *Energy) MegaelectronVolts() float64 {
	if a.megaelectron_voltsLazy != nil {
		return *a.megaelectron_voltsLazy
	}
	megaelectron_volts := a.convertFromBase(EnergyMegaelectronVolt)
	a.megaelectron_voltsLazy = &megaelectron_volts
	return megaelectron_volts
}

// GigaelectronVolts returns the Energy value in GigaelectronVolts.
//
// 
func (a *Energy) GigaelectronVolts() float64 {
	if a.gigaelectron_voltsLazy != nil {
		return *a.gigaelectron_voltsLazy
	}
	gigaelectron_volts := a.convertFromBase(EnergyGigaelectronVolt)
	a.gigaelectron_voltsLazy = &gigaelectron_volts
	return gigaelectron_volts
}

// TeraelectronVolts returns the Energy value in TeraelectronVolts.
//
// 
func (a *Energy) TeraelectronVolts() float64 {
	if a.teraelectron_voltsLazy != nil {
		return *a.teraelectron_voltsLazy
	}
	teraelectron_volts := a.convertFromBase(EnergyTeraelectronVolt)
	a.teraelectron_voltsLazy = &teraelectron_volts
	return teraelectron_volts
}

// KilowattHours returns the Energy value in KilowattHours.
//
// 
func (a *Energy) KilowattHours() float64 {
	if a.kilowatt_hoursLazy != nil {
		return *a.kilowatt_hoursLazy
	}
	kilowatt_hours := a.convertFromBase(EnergyKilowattHour)
	a.kilowatt_hoursLazy = &kilowatt_hours
	return kilowatt_hours
}

// MegawattHours returns the Energy value in MegawattHours.
//
// 
func (a *Energy) MegawattHours() float64 {
	if a.megawatt_hoursLazy != nil {
		return *a.megawatt_hoursLazy
	}
	megawatt_hours := a.convertFromBase(EnergyMegawattHour)
	a.megawatt_hoursLazy = &megawatt_hours
	return megawatt_hours
}

// GigawattHours returns the Energy value in GigawattHours.
//
// 
func (a *Energy) GigawattHours() float64 {
	if a.gigawatt_hoursLazy != nil {
		return *a.gigawatt_hoursLazy
	}
	gigawatt_hours := a.convertFromBase(EnergyGigawattHour)
	a.gigawatt_hoursLazy = &gigawatt_hours
	return gigawatt_hours
}

// TerawattHours returns the Energy value in TerawattHours.
//
// 
func (a *Energy) TerawattHours() float64 {
	if a.terawatt_hoursLazy != nil {
		return *a.terawatt_hoursLazy
	}
	terawatt_hours := a.convertFromBase(EnergyTerawattHour)
	a.terawatt_hoursLazy = &terawatt_hours
	return terawatt_hours
}

// KilowattDays returns the Energy value in KilowattDays.
//
// 
func (a *Energy) KilowattDays() float64 {
	if a.kilowatt_daysLazy != nil {
		return *a.kilowatt_daysLazy
	}
	kilowatt_days := a.convertFromBase(EnergyKilowattDay)
	a.kilowatt_daysLazy = &kilowatt_days
	return kilowatt_days
}

// MegawattDays returns the Energy value in MegawattDays.
//
// 
func (a *Energy) MegawattDays() float64 {
	if a.megawatt_daysLazy != nil {
		return *a.megawatt_daysLazy
	}
	megawatt_days := a.convertFromBase(EnergyMegawattDay)
	a.megawatt_daysLazy = &megawatt_days
	return megawatt_days
}

// GigawattDays returns the Energy value in GigawattDays.
//
// 
func (a *Energy) GigawattDays() float64 {
	if a.gigawatt_daysLazy != nil {
		return *a.gigawatt_daysLazy
	}
	gigawatt_days := a.convertFromBase(EnergyGigawattDay)
	a.gigawatt_daysLazy = &gigawatt_days
	return gigawatt_days
}

// TerawattDays returns the Energy value in TerawattDays.
//
// 
func (a *Energy) TerawattDays() float64 {
	if a.terawatt_daysLazy != nil {
		return *a.terawatt_daysLazy
	}
	terawatt_days := a.convertFromBase(EnergyTerawattDay)
	a.terawatt_daysLazy = &terawatt_days
	return terawatt_days
}

// DecathermsEc returns the Energy value in DecathermsEc.
//
// 
func (a *Energy) DecathermsEc() float64 {
	if a.decatherms_ecLazy != nil {
		return *a.decatherms_ecLazy
	}
	decatherms_ec := a.convertFromBase(EnergyDecathermEc)
	a.decatherms_ecLazy = &decatherms_ec
	return decatherms_ec
}

// DecathermsUs returns the Energy value in DecathermsUs.
//
// 
func (a *Energy) DecathermsUs() float64 {
	if a.decatherms_usLazy != nil {
		return *a.decatherms_usLazy
	}
	decatherms_us := a.convertFromBase(EnergyDecathermUs)
	a.decatherms_usLazy = &decatherms_us
	return decatherms_us
}

// DecathermsImperial returns the Energy value in DecathermsImperial.
//
// 
func (a *Energy) DecathermsImperial() float64 {
	if a.decatherms_imperialLazy != nil {
		return *a.decatherms_imperialLazy
	}
	decatherms_imperial := a.convertFromBase(EnergyDecathermImperial)
	a.decatherms_imperialLazy = &decatherms_imperial
	return decatherms_imperial
}


// ToDto creates a EnergyDto representation from the Energy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Joule by default.
func (a *Energy) ToDto(holdInUnit *EnergyUnits) EnergyDto {
	if holdInUnit == nil {
		defaultUnit := EnergyJoule // Default value
		holdInUnit = &defaultUnit
	}

	return EnergyDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Energy instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Joule by default.
func (a *Energy) ToDtoJSON(holdInUnit *EnergyUnits) ([]byte, error) {
	// Convert to EnergyDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Energy to a specific unit value.
// The function uses the provided unit type (EnergyUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Energy) Convert(toUnit EnergyUnits) float64 {
	switch toUnit { 
    case EnergyJoule:
		return a.Joules()
    case EnergyCalorie:
		return a.Calories()
    case EnergyBritishThermalUnit:
		return a.BritishThermalUnits()
    case EnergyElectronVolt:
		return a.ElectronVolts()
    case EnergyFootPound:
		return a.FootPounds()
    case EnergyErg:
		return a.Ergs()
    case EnergyWattHour:
		return a.WattHours()
    case EnergyWattDay:
		return a.WattDays()
    case EnergyThermEc:
		return a.ThermsEc()
    case EnergyThermUs:
		return a.ThermsUs()
    case EnergyThermImperial:
		return a.ThermsImperial()
    case EnergyHorsepowerHour:
		return a.HorsepowerHours()
    case EnergyNanojoule:
		return a.Nanojoules()
    case EnergyMicrojoule:
		return a.Microjoules()
    case EnergyMillijoule:
		return a.Millijoules()
    case EnergyKilojoule:
		return a.Kilojoules()
    case EnergyMegajoule:
		return a.Megajoules()
    case EnergyGigajoule:
		return a.Gigajoules()
    case EnergyTerajoule:
		return a.Terajoules()
    case EnergyPetajoule:
		return a.Petajoules()
    case EnergyKilocalorie:
		return a.Kilocalories()
    case EnergyMegacalorie:
		return a.Megacalories()
    case EnergyKilobritishThermalUnit:
		return a.KilobritishThermalUnits()
    case EnergyMegabritishThermalUnit:
		return a.MegabritishThermalUnits()
    case EnergyGigabritishThermalUnit:
		return a.GigabritishThermalUnits()
    case EnergyKiloelectronVolt:
		return a.KiloelectronVolts()
    case EnergyMegaelectronVolt:
		return a.MegaelectronVolts()
    case EnergyGigaelectronVolt:
		return a.GigaelectronVolts()
    case EnergyTeraelectronVolt:
		return a.TeraelectronVolts()
    case EnergyKilowattHour:
		return a.KilowattHours()
    case EnergyMegawattHour:
		return a.MegawattHours()
    case EnergyGigawattHour:
		return a.GigawattHours()
    case EnergyTerawattHour:
		return a.TerawattHours()
    case EnergyKilowattDay:
		return a.KilowattDays()
    case EnergyMegawattDay:
		return a.MegawattDays()
    case EnergyGigawattDay:
		return a.GigawattDays()
    case EnergyTerawattDay:
		return a.TerawattDays()
    case EnergyDecathermEc:
		return a.DecathermsEc()
    case EnergyDecathermUs:
		return a.DecathermsUs()
    case EnergyDecathermImperial:
		return a.DecathermsImperial()
	default:
		return math.NaN()
	}
}

func (a *Energy) convertFromBase(toUnit EnergyUnits) float64 {
    value := a.value
	switch toUnit { 
	case EnergyJoule:
		return (value) 
	case EnergyCalorie:
		return (value / 4.184) 
	case EnergyBritishThermalUnit:
		return (value / 1055.05585262) 
	case EnergyElectronVolt:
		return (value / 1.602176634e-19) 
	case EnergyFootPound:
		return (value / 1.3558179483314004) 
	case EnergyErg:
		return (value / 1e-7) 
	case EnergyWattHour:
		return (value / 3600) 
	case EnergyWattDay:
		return (value / (24 * 3600)) 
	case EnergyThermEc:
		return (value / 1.05505585262e8) 
	case EnergyThermUs:
		return (value / 1.054804e8) 
	case EnergyThermImperial:
		return (value / 1.05505585257348e8) 
	case EnergyHorsepowerHour:
		return (value / (76.0402249 * 9.80665 * 3600)) 
	case EnergyNanojoule:
		return ((value) / 1e-09) 
	case EnergyMicrojoule:
		return ((value) / 1e-06) 
	case EnergyMillijoule:
		return ((value) / 0.001) 
	case EnergyKilojoule:
		return ((value) / 1000.0) 
	case EnergyMegajoule:
		return ((value) / 1000000.0) 
	case EnergyGigajoule:
		return ((value) / 1000000000.0) 
	case EnergyTerajoule:
		return ((value) / 1000000000000.0) 
	case EnergyPetajoule:
		return ((value) / 1000000000000000.0) 
	case EnergyKilocalorie:
		return ((value / 4.184) / 1000.0) 
	case EnergyMegacalorie:
		return ((value / 4.184) / 1000000.0) 
	case EnergyKilobritishThermalUnit:
		return ((value / 1055.05585262) / 1000.0) 
	case EnergyMegabritishThermalUnit:
		return ((value / 1055.05585262) / 1000000.0) 
	case EnergyGigabritishThermalUnit:
		return ((value / 1055.05585262) / 1000000000.0) 
	case EnergyKiloelectronVolt:
		return ((value / 1.602176634e-19) / 1000.0) 
	case EnergyMegaelectronVolt:
		return ((value / 1.602176634e-19) / 1000000.0) 
	case EnergyGigaelectronVolt:
		return ((value / 1.602176634e-19) / 1000000000.0) 
	case EnergyTeraelectronVolt:
		return ((value / 1.602176634e-19) / 1000000000000.0) 
	case EnergyKilowattHour:
		return ((value / 3600) / 1000.0) 
	case EnergyMegawattHour:
		return ((value / 3600) / 1000000.0) 
	case EnergyGigawattHour:
		return ((value / 3600) / 1000000000.0) 
	case EnergyTerawattHour:
		return ((value / 3600) / 1000000000000.0) 
	case EnergyKilowattDay:
		return ((value / (24 * 3600)) / 1000.0) 
	case EnergyMegawattDay:
		return ((value / (24 * 3600)) / 1000000.0) 
	case EnergyGigawattDay:
		return ((value / (24 * 3600)) / 1000000000.0) 
	case EnergyTerawattDay:
		return ((value / (24 * 3600)) / 1000000000000.0) 
	case EnergyDecathermEc:
		return ((value / 1.05505585262e8) / 10.0) 
	case EnergyDecathermUs:
		return ((value / 1.054804e8) / 10.0) 
	case EnergyDecathermImperial:
		return ((value / 1.05505585257348e8) / 10.0) 
	default:
		return math.NaN()
	}
}

func (a *Energy) convertToBase(value float64, fromUnit EnergyUnits) float64 {
	switch fromUnit { 
	case EnergyJoule:
		return (value) 
	case EnergyCalorie:
		return (value * 4.184) 
	case EnergyBritishThermalUnit:
		return (value * 1055.05585262) 
	case EnergyElectronVolt:
		return (value * 1.602176634e-19) 
	case EnergyFootPound:
		return (value * 1.3558179483314004) 
	case EnergyErg:
		return (value * 1e-7) 
	case EnergyWattHour:
		return (value * 3600) 
	case EnergyWattDay:
		return (value * 24 * 3600) 
	case EnergyThermEc:
		return (value * 1.05505585262e8) 
	case EnergyThermUs:
		return (value * 1.054804e8) 
	case EnergyThermImperial:
		return (value * 1.05505585257348e8) 
	case EnergyHorsepowerHour:
		return (value * 76.0402249 * 9.80665 * 3600) 
	case EnergyNanojoule:
		return ((value) * 1e-09) 
	case EnergyMicrojoule:
		return ((value) * 1e-06) 
	case EnergyMillijoule:
		return ((value) * 0.001) 
	case EnergyKilojoule:
		return ((value) * 1000.0) 
	case EnergyMegajoule:
		return ((value) * 1000000.0) 
	case EnergyGigajoule:
		return ((value) * 1000000000.0) 
	case EnergyTerajoule:
		return ((value) * 1000000000000.0) 
	case EnergyPetajoule:
		return ((value) * 1000000000000000.0) 
	case EnergyKilocalorie:
		return ((value * 4.184) * 1000.0) 
	case EnergyMegacalorie:
		return ((value * 4.184) * 1000000.0) 
	case EnergyKilobritishThermalUnit:
		return ((value * 1055.05585262) * 1000.0) 
	case EnergyMegabritishThermalUnit:
		return ((value * 1055.05585262) * 1000000.0) 
	case EnergyGigabritishThermalUnit:
		return ((value * 1055.05585262) * 1000000000.0) 
	case EnergyKiloelectronVolt:
		return ((value * 1.602176634e-19) * 1000.0) 
	case EnergyMegaelectronVolt:
		return ((value * 1.602176634e-19) * 1000000.0) 
	case EnergyGigaelectronVolt:
		return ((value * 1.602176634e-19) * 1000000000.0) 
	case EnergyTeraelectronVolt:
		return ((value * 1.602176634e-19) * 1000000000000.0) 
	case EnergyKilowattHour:
		return ((value * 3600) * 1000.0) 
	case EnergyMegawattHour:
		return ((value * 3600) * 1000000.0) 
	case EnergyGigawattHour:
		return ((value * 3600) * 1000000000.0) 
	case EnergyTerawattHour:
		return ((value * 3600) * 1000000000000.0) 
	case EnergyKilowattDay:
		return ((value * 24 * 3600) * 1000.0) 
	case EnergyMegawattDay:
		return ((value * 24 * 3600) * 1000000.0) 
	case EnergyGigawattDay:
		return ((value * 24 * 3600) * 1000000000.0) 
	case EnergyTerawattDay:
		return ((value * 24 * 3600) * 1000000000000.0) 
	case EnergyDecathermEc:
		return ((value * 1.05505585262e8) * 10.0) 
	case EnergyDecathermUs:
		return ((value * 1.054804e8) * 10.0) 
	case EnergyDecathermImperial:
		return ((value * 1.05505585257348e8) * 10.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Energy in the default unit (Joule),
// formatted to two decimal places.
func (a Energy) String() string {
	return a.ToString(EnergyJoule, 2)
}

// ToString formats the Energy value as a string with the specified unit and fractional digits.
// It converts the Energy to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Energy value will be converted (e.g., Joule))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Energy with the unit abbreviation.
func (a *Energy) ToString(unit EnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetEnergyAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetEnergyAbbreviation(unit))
}

// Equals checks if the given Energy is equal to the current Energy.
//
// Parameters:
//    other: The Energy to compare against.
//
// Returns:
//    bool: Returns true if both Energy are equal, false otherwise.
func (a *Energy) Equals(other *Energy) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Energy with another Energy.
// It returns -1 if the current Energy is less than the other Energy, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Energy to compare against.
//
// Returns:
//    int: -1 if the current Energy is less, 1 if greater, and 0 if equal.
func (a *Energy) CompareTo(other *Energy) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Energy to the current Energy and returns the result.
// The result is a new Energy instance with the sum of the values.
//
// Parameters:
//    other: The Energy to add to the current Energy.
//
// Returns:
//    *Energy: A new Energy instance representing the sum of both Energy.
func (a *Energy) Add(other *Energy) *Energy {
	return &Energy{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Energy from the current Energy and returns the result.
// The result is a new Energy instance with the difference of the values.
//
// Parameters:
//    other: The Energy to subtract from the current Energy.
//
// Returns:
//    *Energy: A new Energy instance representing the difference of both Energy.
func (a *Energy) Subtract(other *Energy) *Energy {
	return &Energy{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Energy by the given Energy and returns the result.
// The result is a new Energy instance with the product of the values.
//
// Parameters:
//    other: The Energy to multiply with the current Energy.
//
// Returns:
//    *Energy: A new Energy instance representing the product of both Energy.
func (a *Energy) Multiply(other *Energy) *Energy {
	return &Energy{value: a.value * other.BaseValue()}
}

// Divide divides the current Energy by the given Energy and returns the result.
// The result is a new Energy instance with the quotient of the values.
//
// Parameters:
//    other: The Energy to divide the current Energy by.
//
// Returns:
//    *Energy: A new Energy instance representing the quotient of both Energy.
func (a *Energy) Divide(other *Energy) *Energy {
	return &Energy{value: a.value / other.BaseValue()}
}

// GetEnergyAbbreviation gets the unit abbreviation.
func GetEnergyAbbreviation(unit EnergyUnits) string {
	switch unit { 
	case EnergyJoule:
		return "J" 
	case EnergyCalorie:
		return "cal" 
	case EnergyBritishThermalUnit:
		return "BTU" 
	case EnergyElectronVolt:
		return "eV" 
	case EnergyFootPound:
		return "ft·lb" 
	case EnergyErg:
		return "erg" 
	case EnergyWattHour:
		return "Wh" 
	case EnergyWattDay:
		return "Wd" 
	case EnergyThermEc:
		return "th (E.C.)" 
	case EnergyThermUs:
		return "th (U.S.)" 
	case EnergyThermImperial:
		return "th (imp.)" 
	case EnergyHorsepowerHour:
		return "hp·h" 
	case EnergyNanojoule:
		return "nJ" 
	case EnergyMicrojoule:
		return "μJ" 
	case EnergyMillijoule:
		return "mJ" 
	case EnergyKilojoule:
		return "kJ" 
	case EnergyMegajoule:
		return "MJ" 
	case EnergyGigajoule:
		return "GJ" 
	case EnergyTerajoule:
		return "TJ" 
	case EnergyPetajoule:
		return "PJ" 
	case EnergyKilocalorie:
		return "kcal" 
	case EnergyMegacalorie:
		return "Mcal" 
	case EnergyKilobritishThermalUnit:
		return "kBTU" 
	case EnergyMegabritishThermalUnit:
		return "MBTU" 
	case EnergyGigabritishThermalUnit:
		return "GBTU" 
	case EnergyKiloelectronVolt:
		return "keV" 
	case EnergyMegaelectronVolt:
		return "MeV" 
	case EnergyGigaelectronVolt:
		return "GeV" 
	case EnergyTeraelectronVolt:
		return "TeV" 
	case EnergyKilowattHour:
		return "kWh" 
	case EnergyMegawattHour:
		return "MWh" 
	case EnergyGigawattHour:
		return "GWh" 
	case EnergyTerawattHour:
		return "TWh" 
	case EnergyKilowattDay:
		return "kWd" 
	case EnergyMegawattDay:
		return "MWd" 
	case EnergyGigawattDay:
		return "GWd" 
	case EnergyTerawattDay:
		return "TWd" 
	case EnergyDecathermEc:
		return "dath (E.C.)" 
	case EnergyDecathermUs:
		return "dath (U.S.)" 
	case EnergyDecathermImperial:
		return "dath (imp.)" 
	default:
		return ""
	}
}