// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// EnergyUnits enumeration
type EnergyUnits string

const (
    
        // 
        EnergyJoule EnergyUnits = "Joule"
        // 
        EnergyCalorie EnergyUnits = "Calorie"
        // 
        EnergyBritishThermalUnit EnergyUnits = "BritishThermalUnit"
        // 
        EnergyElectronVolt EnergyUnits = "ElectronVolt"
        // 
        EnergyFootPound EnergyUnits = "FootPound"
        // 
        EnergyErg EnergyUnits = "Erg"
        // 
        EnergyWattHour EnergyUnits = "WattHour"
        // 
        EnergyWattDay EnergyUnits = "WattDay"
        // 
        EnergyThermEc EnergyUnits = "ThermEc"
        // 
        EnergyThermUs EnergyUnits = "ThermUs"
        // 
        EnergyThermImperial EnergyUnits = "ThermImperial"
        // 
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

// EnergyDto represents an Energy
type EnergyDto struct {
	Value float64
	Unit  EnergyUnits
}

// EnergyDtoFactory struct to group related functions
type EnergyDtoFactory struct{}

func (udf EnergyDtoFactory) FromJSON(data []byte) (*EnergyDto, error) {
	a := EnergyDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a EnergyDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Energy struct
type Energy struct {
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

// EnergyFactory struct to group related functions
type EnergyFactory struct{}

func (uf EnergyFactory) CreateEnergy(value float64, unit EnergyUnits) (*Energy, error) {
	return newEnergy(value, unit)
}

func (uf EnergyFactory) FromDto(dto EnergyDto) (*Energy, error) {
	return newEnergy(dto.Value, dto.Unit)
}

func (uf EnergyFactory) FromDtoJSON(data []byte) (*Energy, error) {
	unitDto, err := EnergyDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return EnergyFactory{}.FromDto(*unitDto)
}


// FromJoule creates a new Energy instance from Joule.
func (uf EnergyFactory) FromJoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyJoule)
}

// FromCalorie creates a new Energy instance from Calorie.
func (uf EnergyFactory) FromCalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyCalorie)
}

// FromBritishThermalUnit creates a new Energy instance from BritishThermalUnit.
func (uf EnergyFactory) FromBritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyBritishThermalUnit)
}

// FromElectronVolt creates a new Energy instance from ElectronVolt.
func (uf EnergyFactory) FromElectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyElectronVolt)
}

// FromFootPound creates a new Energy instance from FootPound.
func (uf EnergyFactory) FromFootPounds(value float64) (*Energy, error) {
	return newEnergy(value, EnergyFootPound)
}

// FromErg creates a new Energy instance from Erg.
func (uf EnergyFactory) FromErgs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyErg)
}

// FromWattHour creates a new Energy instance from WattHour.
func (uf EnergyFactory) FromWattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyWattHour)
}

// FromWattDay creates a new Energy instance from WattDay.
func (uf EnergyFactory) FromWattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyWattDay)
}

// FromThermEc creates a new Energy instance from ThermEc.
func (uf EnergyFactory) FromThermsEc(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermEc)
}

// FromThermUs creates a new Energy instance from ThermUs.
func (uf EnergyFactory) FromThermsUs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermUs)
}

// FromThermImperial creates a new Energy instance from ThermImperial.
func (uf EnergyFactory) FromThermsImperial(value float64) (*Energy, error) {
	return newEnergy(value, EnergyThermImperial)
}

// FromHorsepowerHour creates a new Energy instance from HorsepowerHour.
func (uf EnergyFactory) FromHorsepowerHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyHorsepowerHour)
}

// FromNanojoule creates a new Energy instance from Nanojoule.
func (uf EnergyFactory) FromNanojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyNanojoule)
}

// FromMicrojoule creates a new Energy instance from Microjoule.
func (uf EnergyFactory) FromMicrojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMicrojoule)
}

// FromMillijoule creates a new Energy instance from Millijoule.
func (uf EnergyFactory) FromMillijoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMillijoule)
}

// FromKilojoule creates a new Energy instance from Kilojoule.
func (uf EnergyFactory) FromKilojoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilojoule)
}

// FromMegajoule creates a new Energy instance from Megajoule.
func (uf EnergyFactory) FromMegajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegajoule)
}

// FromGigajoule creates a new Energy instance from Gigajoule.
func (uf EnergyFactory) FromGigajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigajoule)
}

// FromTerajoule creates a new Energy instance from Terajoule.
func (uf EnergyFactory) FromTerajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerajoule)
}

// FromPetajoule creates a new Energy instance from Petajoule.
func (uf EnergyFactory) FromPetajoules(value float64) (*Energy, error) {
	return newEnergy(value, EnergyPetajoule)
}

// FromKilocalorie creates a new Energy instance from Kilocalorie.
func (uf EnergyFactory) FromKilocalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilocalorie)
}

// FromMegacalorie creates a new Energy instance from Megacalorie.
func (uf EnergyFactory) FromMegacalories(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegacalorie)
}

// FromKilobritishThermalUnit creates a new Energy instance from KilobritishThermalUnit.
func (uf EnergyFactory) FromKilobritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilobritishThermalUnit)
}

// FromMegabritishThermalUnit creates a new Energy instance from MegabritishThermalUnit.
func (uf EnergyFactory) FromMegabritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegabritishThermalUnit)
}

// FromGigabritishThermalUnit creates a new Energy instance from GigabritishThermalUnit.
func (uf EnergyFactory) FromGigabritishThermalUnits(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigabritishThermalUnit)
}

// FromKiloelectronVolt creates a new Energy instance from KiloelectronVolt.
func (uf EnergyFactory) FromKiloelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKiloelectronVolt)
}

// FromMegaelectronVolt creates a new Energy instance from MegaelectronVolt.
func (uf EnergyFactory) FromMegaelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegaelectronVolt)
}

// FromGigaelectronVolt creates a new Energy instance from GigaelectronVolt.
func (uf EnergyFactory) FromGigaelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigaelectronVolt)
}

// FromTeraelectronVolt creates a new Energy instance from TeraelectronVolt.
func (uf EnergyFactory) FromTeraelectronVolts(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTeraelectronVolt)
}

// FromKilowattHour creates a new Energy instance from KilowattHour.
func (uf EnergyFactory) FromKilowattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilowattHour)
}

// FromMegawattHour creates a new Energy instance from MegawattHour.
func (uf EnergyFactory) FromMegawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegawattHour)
}

// FromGigawattHour creates a new Energy instance from GigawattHour.
func (uf EnergyFactory) FromGigawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigawattHour)
}

// FromTerawattHour creates a new Energy instance from TerawattHour.
func (uf EnergyFactory) FromTerawattHours(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerawattHour)
}

// FromKilowattDay creates a new Energy instance from KilowattDay.
func (uf EnergyFactory) FromKilowattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyKilowattDay)
}

// FromMegawattDay creates a new Energy instance from MegawattDay.
func (uf EnergyFactory) FromMegawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyMegawattDay)
}

// FromGigawattDay creates a new Energy instance from GigawattDay.
func (uf EnergyFactory) FromGigawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyGigawattDay)
}

// FromTerawattDay creates a new Energy instance from TerawattDay.
func (uf EnergyFactory) FromTerawattDays(value float64) (*Energy, error) {
	return newEnergy(value, EnergyTerawattDay)
}

// FromDecathermEc creates a new Energy instance from DecathermEc.
func (uf EnergyFactory) FromDecathermsEc(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermEc)
}

// FromDecathermUs creates a new Energy instance from DecathermUs.
func (uf EnergyFactory) FromDecathermsUs(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermUs)
}

// FromDecathermImperial creates a new Energy instance from DecathermImperial.
func (uf EnergyFactory) FromDecathermsImperial(value float64) (*Energy, error) {
	return newEnergy(value, EnergyDecathermImperial)
}




// newEnergy creates a new Energy.
func newEnergy(value float64, fromUnit EnergyUnits) (*Energy, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Energy{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Energy in Joule.
func (a *Energy) BaseValue() float64 {
	return a.value
}


// Joule returns the value in Joule.
func (a *Energy) Joules() float64 {
	if a.joulesLazy != nil {
		return *a.joulesLazy
	}
	joules := a.convertFromBase(EnergyJoule)
	a.joulesLazy = &joules
	return joules
}

// Calorie returns the value in Calorie.
func (a *Energy) Calories() float64 {
	if a.caloriesLazy != nil {
		return *a.caloriesLazy
	}
	calories := a.convertFromBase(EnergyCalorie)
	a.caloriesLazy = &calories
	return calories
}

// BritishThermalUnit returns the value in BritishThermalUnit.
func (a *Energy) BritishThermalUnits() float64 {
	if a.british_thermal_unitsLazy != nil {
		return *a.british_thermal_unitsLazy
	}
	british_thermal_units := a.convertFromBase(EnergyBritishThermalUnit)
	a.british_thermal_unitsLazy = &british_thermal_units
	return british_thermal_units
}

// ElectronVolt returns the value in ElectronVolt.
func (a *Energy) ElectronVolts() float64 {
	if a.electron_voltsLazy != nil {
		return *a.electron_voltsLazy
	}
	electron_volts := a.convertFromBase(EnergyElectronVolt)
	a.electron_voltsLazy = &electron_volts
	return electron_volts
}

// FootPound returns the value in FootPound.
func (a *Energy) FootPounds() float64 {
	if a.foot_poundsLazy != nil {
		return *a.foot_poundsLazy
	}
	foot_pounds := a.convertFromBase(EnergyFootPound)
	a.foot_poundsLazy = &foot_pounds
	return foot_pounds
}

// Erg returns the value in Erg.
func (a *Energy) Ergs() float64 {
	if a.ergsLazy != nil {
		return *a.ergsLazy
	}
	ergs := a.convertFromBase(EnergyErg)
	a.ergsLazy = &ergs
	return ergs
}

// WattHour returns the value in WattHour.
func (a *Energy) WattHours() float64 {
	if a.watt_hoursLazy != nil {
		return *a.watt_hoursLazy
	}
	watt_hours := a.convertFromBase(EnergyWattHour)
	a.watt_hoursLazy = &watt_hours
	return watt_hours
}

// WattDay returns the value in WattDay.
func (a *Energy) WattDays() float64 {
	if a.watt_daysLazy != nil {
		return *a.watt_daysLazy
	}
	watt_days := a.convertFromBase(EnergyWattDay)
	a.watt_daysLazy = &watt_days
	return watt_days
}

// ThermEc returns the value in ThermEc.
func (a *Energy) ThermsEc() float64 {
	if a.therms_ecLazy != nil {
		return *a.therms_ecLazy
	}
	therms_ec := a.convertFromBase(EnergyThermEc)
	a.therms_ecLazy = &therms_ec
	return therms_ec
}

// ThermUs returns the value in ThermUs.
func (a *Energy) ThermsUs() float64 {
	if a.therms_usLazy != nil {
		return *a.therms_usLazy
	}
	therms_us := a.convertFromBase(EnergyThermUs)
	a.therms_usLazy = &therms_us
	return therms_us
}

// ThermImperial returns the value in ThermImperial.
func (a *Energy) ThermsImperial() float64 {
	if a.therms_imperialLazy != nil {
		return *a.therms_imperialLazy
	}
	therms_imperial := a.convertFromBase(EnergyThermImperial)
	a.therms_imperialLazy = &therms_imperial
	return therms_imperial
}

// HorsepowerHour returns the value in HorsepowerHour.
func (a *Energy) HorsepowerHours() float64 {
	if a.horsepower_hoursLazy != nil {
		return *a.horsepower_hoursLazy
	}
	horsepower_hours := a.convertFromBase(EnergyHorsepowerHour)
	a.horsepower_hoursLazy = &horsepower_hours
	return horsepower_hours
}

// Nanojoule returns the value in Nanojoule.
func (a *Energy) Nanojoules() float64 {
	if a.nanojoulesLazy != nil {
		return *a.nanojoulesLazy
	}
	nanojoules := a.convertFromBase(EnergyNanojoule)
	a.nanojoulesLazy = &nanojoules
	return nanojoules
}

// Microjoule returns the value in Microjoule.
func (a *Energy) Microjoules() float64 {
	if a.microjoulesLazy != nil {
		return *a.microjoulesLazy
	}
	microjoules := a.convertFromBase(EnergyMicrojoule)
	a.microjoulesLazy = &microjoules
	return microjoules
}

// Millijoule returns the value in Millijoule.
func (a *Energy) Millijoules() float64 {
	if a.millijoulesLazy != nil {
		return *a.millijoulesLazy
	}
	millijoules := a.convertFromBase(EnergyMillijoule)
	a.millijoulesLazy = &millijoules
	return millijoules
}

// Kilojoule returns the value in Kilojoule.
func (a *Energy) Kilojoules() float64 {
	if a.kilojoulesLazy != nil {
		return *a.kilojoulesLazy
	}
	kilojoules := a.convertFromBase(EnergyKilojoule)
	a.kilojoulesLazy = &kilojoules
	return kilojoules
}

// Megajoule returns the value in Megajoule.
func (a *Energy) Megajoules() float64 {
	if a.megajoulesLazy != nil {
		return *a.megajoulesLazy
	}
	megajoules := a.convertFromBase(EnergyMegajoule)
	a.megajoulesLazy = &megajoules
	return megajoules
}

// Gigajoule returns the value in Gigajoule.
func (a *Energy) Gigajoules() float64 {
	if a.gigajoulesLazy != nil {
		return *a.gigajoulesLazy
	}
	gigajoules := a.convertFromBase(EnergyGigajoule)
	a.gigajoulesLazy = &gigajoules
	return gigajoules
}

// Terajoule returns the value in Terajoule.
func (a *Energy) Terajoules() float64 {
	if a.terajoulesLazy != nil {
		return *a.terajoulesLazy
	}
	terajoules := a.convertFromBase(EnergyTerajoule)
	a.terajoulesLazy = &terajoules
	return terajoules
}

// Petajoule returns the value in Petajoule.
func (a *Energy) Petajoules() float64 {
	if a.petajoulesLazy != nil {
		return *a.petajoulesLazy
	}
	petajoules := a.convertFromBase(EnergyPetajoule)
	a.petajoulesLazy = &petajoules
	return petajoules
}

// Kilocalorie returns the value in Kilocalorie.
func (a *Energy) Kilocalories() float64 {
	if a.kilocaloriesLazy != nil {
		return *a.kilocaloriesLazy
	}
	kilocalories := a.convertFromBase(EnergyKilocalorie)
	a.kilocaloriesLazy = &kilocalories
	return kilocalories
}

// Megacalorie returns the value in Megacalorie.
func (a *Energy) Megacalories() float64 {
	if a.megacaloriesLazy != nil {
		return *a.megacaloriesLazy
	}
	megacalories := a.convertFromBase(EnergyMegacalorie)
	a.megacaloriesLazy = &megacalories
	return megacalories
}

// KilobritishThermalUnit returns the value in KilobritishThermalUnit.
func (a *Energy) KilobritishThermalUnits() float64 {
	if a.kilobritish_thermal_unitsLazy != nil {
		return *a.kilobritish_thermal_unitsLazy
	}
	kilobritish_thermal_units := a.convertFromBase(EnergyKilobritishThermalUnit)
	a.kilobritish_thermal_unitsLazy = &kilobritish_thermal_units
	return kilobritish_thermal_units
}

// MegabritishThermalUnit returns the value in MegabritishThermalUnit.
func (a *Energy) MegabritishThermalUnits() float64 {
	if a.megabritish_thermal_unitsLazy != nil {
		return *a.megabritish_thermal_unitsLazy
	}
	megabritish_thermal_units := a.convertFromBase(EnergyMegabritishThermalUnit)
	a.megabritish_thermal_unitsLazy = &megabritish_thermal_units
	return megabritish_thermal_units
}

// GigabritishThermalUnit returns the value in GigabritishThermalUnit.
func (a *Energy) GigabritishThermalUnits() float64 {
	if a.gigabritish_thermal_unitsLazy != nil {
		return *a.gigabritish_thermal_unitsLazy
	}
	gigabritish_thermal_units := a.convertFromBase(EnergyGigabritishThermalUnit)
	a.gigabritish_thermal_unitsLazy = &gigabritish_thermal_units
	return gigabritish_thermal_units
}

// KiloelectronVolt returns the value in KiloelectronVolt.
func (a *Energy) KiloelectronVolts() float64 {
	if a.kiloelectron_voltsLazy != nil {
		return *a.kiloelectron_voltsLazy
	}
	kiloelectron_volts := a.convertFromBase(EnergyKiloelectronVolt)
	a.kiloelectron_voltsLazy = &kiloelectron_volts
	return kiloelectron_volts
}

// MegaelectronVolt returns the value in MegaelectronVolt.
func (a *Energy) MegaelectronVolts() float64 {
	if a.megaelectron_voltsLazy != nil {
		return *a.megaelectron_voltsLazy
	}
	megaelectron_volts := a.convertFromBase(EnergyMegaelectronVolt)
	a.megaelectron_voltsLazy = &megaelectron_volts
	return megaelectron_volts
}

// GigaelectronVolt returns the value in GigaelectronVolt.
func (a *Energy) GigaelectronVolts() float64 {
	if a.gigaelectron_voltsLazy != nil {
		return *a.gigaelectron_voltsLazy
	}
	gigaelectron_volts := a.convertFromBase(EnergyGigaelectronVolt)
	a.gigaelectron_voltsLazy = &gigaelectron_volts
	return gigaelectron_volts
}

// TeraelectronVolt returns the value in TeraelectronVolt.
func (a *Energy) TeraelectronVolts() float64 {
	if a.teraelectron_voltsLazy != nil {
		return *a.teraelectron_voltsLazy
	}
	teraelectron_volts := a.convertFromBase(EnergyTeraelectronVolt)
	a.teraelectron_voltsLazy = &teraelectron_volts
	return teraelectron_volts
}

// KilowattHour returns the value in KilowattHour.
func (a *Energy) KilowattHours() float64 {
	if a.kilowatt_hoursLazy != nil {
		return *a.kilowatt_hoursLazy
	}
	kilowatt_hours := a.convertFromBase(EnergyKilowattHour)
	a.kilowatt_hoursLazy = &kilowatt_hours
	return kilowatt_hours
}

// MegawattHour returns the value in MegawattHour.
func (a *Energy) MegawattHours() float64 {
	if a.megawatt_hoursLazy != nil {
		return *a.megawatt_hoursLazy
	}
	megawatt_hours := a.convertFromBase(EnergyMegawattHour)
	a.megawatt_hoursLazy = &megawatt_hours
	return megawatt_hours
}

// GigawattHour returns the value in GigawattHour.
func (a *Energy) GigawattHours() float64 {
	if a.gigawatt_hoursLazy != nil {
		return *a.gigawatt_hoursLazy
	}
	gigawatt_hours := a.convertFromBase(EnergyGigawattHour)
	a.gigawatt_hoursLazy = &gigawatt_hours
	return gigawatt_hours
}

// TerawattHour returns the value in TerawattHour.
func (a *Energy) TerawattHours() float64 {
	if a.terawatt_hoursLazy != nil {
		return *a.terawatt_hoursLazy
	}
	terawatt_hours := a.convertFromBase(EnergyTerawattHour)
	a.terawatt_hoursLazy = &terawatt_hours
	return terawatt_hours
}

// KilowattDay returns the value in KilowattDay.
func (a *Energy) KilowattDays() float64 {
	if a.kilowatt_daysLazy != nil {
		return *a.kilowatt_daysLazy
	}
	kilowatt_days := a.convertFromBase(EnergyKilowattDay)
	a.kilowatt_daysLazy = &kilowatt_days
	return kilowatt_days
}

// MegawattDay returns the value in MegawattDay.
func (a *Energy) MegawattDays() float64 {
	if a.megawatt_daysLazy != nil {
		return *a.megawatt_daysLazy
	}
	megawatt_days := a.convertFromBase(EnergyMegawattDay)
	a.megawatt_daysLazy = &megawatt_days
	return megawatt_days
}

// GigawattDay returns the value in GigawattDay.
func (a *Energy) GigawattDays() float64 {
	if a.gigawatt_daysLazy != nil {
		return *a.gigawatt_daysLazy
	}
	gigawatt_days := a.convertFromBase(EnergyGigawattDay)
	a.gigawatt_daysLazy = &gigawatt_days
	return gigawatt_days
}

// TerawattDay returns the value in TerawattDay.
func (a *Energy) TerawattDays() float64 {
	if a.terawatt_daysLazy != nil {
		return *a.terawatt_daysLazy
	}
	terawatt_days := a.convertFromBase(EnergyTerawattDay)
	a.terawatt_daysLazy = &terawatt_days
	return terawatt_days
}

// DecathermEc returns the value in DecathermEc.
func (a *Energy) DecathermsEc() float64 {
	if a.decatherms_ecLazy != nil {
		return *a.decatherms_ecLazy
	}
	decatherms_ec := a.convertFromBase(EnergyDecathermEc)
	a.decatherms_ecLazy = &decatherms_ec
	return decatherms_ec
}

// DecathermUs returns the value in DecathermUs.
func (a *Energy) DecathermsUs() float64 {
	if a.decatherms_usLazy != nil {
		return *a.decatherms_usLazy
	}
	decatherms_us := a.convertFromBase(EnergyDecathermUs)
	a.decatherms_usLazy = &decatherms_us
	return decatherms_us
}

// DecathermImperial returns the value in DecathermImperial.
func (a *Energy) DecathermsImperial() float64 {
	if a.decatherms_imperialLazy != nil {
		return *a.decatherms_imperialLazy
	}
	decatherms_imperial := a.convertFromBase(EnergyDecathermImperial)
	a.decatherms_imperialLazy = &decatherms_imperial
	return decatherms_imperial
}


// ToDto creates an EnergyDto representation.
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

// ToDtoJSON creates an EnergyDto representation.
func (a *Energy) ToDtoJSON(holdInUnit *EnergyUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Energy to a specific unit value.
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
		return 0
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
		return (value / 1.602176565e-19) 
	case EnergyFootPound:
		return (value / 1.355817948) 
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
		return (value / 2.6845195377e6) 
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
		return ((value / 1.602176565e-19) / 1000.0) 
	case EnergyMegaelectronVolt:
		return ((value / 1.602176565e-19) / 1000000.0) 
	case EnergyGigaelectronVolt:
		return ((value / 1.602176565e-19) / 1000000000.0) 
	case EnergyTeraelectronVolt:
		return ((value / 1.602176565e-19) / 1000000000000.0) 
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
		return (value * 1.602176565e-19) 
	case EnergyFootPound:
		return (value * 1.355817948) 
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
		return (value * 2.6845195377e6) 
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
		return ((value * 1.602176565e-19) * 1000.0) 
	case EnergyMegaelectronVolt:
		return ((value * 1.602176565e-19) * 1000000.0) 
	case EnergyGigaelectronVolt:
		return ((value * 1.602176565e-19) * 1000000000.0) 
	case EnergyTeraelectronVolt:
		return ((value * 1.602176565e-19) * 1000000000000.0) 
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

// Implement the String() method for AngleDto
func (a Energy) String() string {
	return a.ToString(EnergyJoule, 2)
}

// ToString formats the Energy to string.
// fractionalDigits -1 for not mention
func (a *Energy) ToString(unit EnergyUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Energy) getUnitAbbreviation(unit EnergyUnits) string {
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

// Check if the given Energy are equals to the current Energy
func (a *Energy) Equals(other *Energy) bool {
	return a.value == other.BaseValue()
}

// Check if the given Energy are equals to the current Energy
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

// Add the given Energy to the current Energy.
func (a *Energy) Add(other *Energy) *Energy {
	return &Energy{value: a.value + other.BaseValue()}
}

// Subtract the given Energy to the current Energy.
func (a *Energy) Subtract(other *Energy) *Energy {
	return &Energy{value: a.value - other.BaseValue()}
}

// Multiply the given Energy to the current Energy.
func (a *Energy) Multiply(other *Energy) *Energy {
	return &Energy{value: a.value * other.BaseValue()}
}

// Divide the given Energy to the current Energy.
func (a *Energy) Divide(other *Energy) *Energy {
	return &Energy{value: a.value / other.BaseValue()}
}