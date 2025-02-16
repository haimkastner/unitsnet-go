// Generated Code - DO NOT EDIT.

package units

import (
	"encoding/json"
	"math"
	"errors"
	"fmt"
	"strconv"
)



// PowerUnits defines various units of Power.
type PowerUnits string

const (
    
        // 
        PowerWatt PowerUnits = "Watt"
        // 
        PowerMechanicalHorsepower PowerUnits = "MechanicalHorsepower"
        // 
        PowerMetricHorsepower PowerUnits = "MetricHorsepower"
        // 
        PowerElectricalHorsepower PowerUnits = "ElectricalHorsepower"
        // 
        PowerBoilerHorsepower PowerUnits = "BoilerHorsepower"
        // 
        PowerHydraulicHorsepower PowerUnits = "HydraulicHorsepower"
        // 
        PowerBritishThermalUnitPerHour PowerUnits = "BritishThermalUnitPerHour"
        // 
        PowerJoulePerHour PowerUnits = "JoulePerHour"
        // 
        PowerTonOfRefrigeration PowerUnits = "TonOfRefrigeration"
        // 
        PowerFemtowatt PowerUnits = "Femtowatt"
        // 
        PowerPicowatt PowerUnits = "Picowatt"
        // 
        PowerNanowatt PowerUnits = "Nanowatt"
        // 
        PowerMicrowatt PowerUnits = "Microwatt"
        // 
        PowerMilliwatt PowerUnits = "Milliwatt"
        // 
        PowerDeciwatt PowerUnits = "Deciwatt"
        // 
        PowerDecawatt PowerUnits = "Decawatt"
        // 
        PowerKilowatt PowerUnits = "Kilowatt"
        // 
        PowerMegawatt PowerUnits = "Megawatt"
        // 
        PowerGigawatt PowerUnits = "Gigawatt"
        // 
        PowerTerawatt PowerUnits = "Terawatt"
        // 
        PowerPetawatt PowerUnits = "Petawatt"
        // 
        PowerKilobritishThermalUnitPerHour PowerUnits = "KilobritishThermalUnitPerHour"
        // 
        PowerMegabritishThermalUnitPerHour PowerUnits = "MegabritishThermalUnitPerHour"
        // 
        PowerMillijoulePerHour PowerUnits = "MillijoulePerHour"
        // 
        PowerKilojoulePerHour PowerUnits = "KilojoulePerHour"
        // 
        PowerMegajoulePerHour PowerUnits = "MegajoulePerHour"
        // 
        PowerGigajoulePerHour PowerUnits = "GigajoulePerHour"
)

// PowerDto represents a Power measurement with a numerical value and its corresponding unit.
type PowerDto struct {
    // Value is the numerical representation of the Power.
	Value float64 `json:"value"`
    // Unit specifies the unit of measurement for the Power, as defined in the PowerUnits enumeration.
	Unit  PowerUnits `json:"unit"`
}

// PowerDtoFactory groups methods for creating and serializing PowerDto objects.
type PowerDtoFactory struct{}

// FromJSON parses a JSON-encoded byte slice into a PowerDto object.
//
// Returns an error if the JSON cannot be parsed.
func (udf PowerDtoFactory) FromJSON(data []byte) (*PowerDto, error) {
	a := PowerDto{}

    // Parse JSON into PowerDto
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	if a.Unit == "" {
		return nil, errors.New("unit is required")
	} 
	
	return &a, nil
}

// ToJSON serializes a PowerDto into a JSON-encoded byte slice.
//
// Returns an error if the serialization fails.
func (a PowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}


// Power represents a measurement in a of Power.
//
// In physics, power is the rate of doing work. It is equivalent to an amount of energy consumed per unit time.
type Power struct {
	// value is the base measurement stored internally.
	value       float64
    
    wattsLazy *float64 
    mechanical_horsepowerLazy *float64 
    metric_horsepowerLazy *float64 
    electrical_horsepowerLazy *float64 
    boiler_horsepowerLazy *float64 
    hydraulic_horsepowerLazy *float64 
    british_thermal_units_per_hourLazy *float64 
    joules_per_hourLazy *float64 
    tons_of_refrigerationLazy *float64 
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
    kilobritish_thermal_units_per_hourLazy *float64 
    megabritish_thermal_units_per_hourLazy *float64 
    millijoules_per_hourLazy *float64 
    kilojoules_per_hourLazy *float64 
    megajoules_per_hourLazy *float64 
    gigajoules_per_hourLazy *float64 
}

// PowerFactory groups methods for creating Power instances.
type PowerFactory struct{}

// CreatePower creates a new Power instance from the given value and unit.
func (uf PowerFactory) CreatePower(value float64, unit PowerUnits) (*Power, error) {
	return newPower(value, unit)
}

// FromDto converts a PowerDto to a Power instance.
func (uf PowerFactory) FromDto(dto PowerDto) (*Power, error) {
	return newPower(dto.Value, dto.Unit)
}

// FromJSON parses a JSON-encoded byte slice into a Power instance.
func (uf PowerFactory) FromDtoJSON(data []byte) (*Power, error) {
	unitDto, err := PowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PowerDto from JSON: %w", err)
	}
	return PowerFactory{}.FromDto(*unitDto)
}


// FromWatts creates a new Power instance from a value in Watts.
func (uf PowerFactory) FromWatts(value float64) (*Power, error) {
	return newPower(value, PowerWatt)
}

// FromMechanicalHorsepower creates a new Power instance from a value in MechanicalHorsepower.
func (uf PowerFactory) FromMechanicalHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerMechanicalHorsepower)
}

// FromMetricHorsepower creates a new Power instance from a value in MetricHorsepower.
func (uf PowerFactory) FromMetricHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerMetricHorsepower)
}

// FromElectricalHorsepower creates a new Power instance from a value in ElectricalHorsepower.
func (uf PowerFactory) FromElectricalHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerElectricalHorsepower)
}

// FromBoilerHorsepower creates a new Power instance from a value in BoilerHorsepower.
func (uf PowerFactory) FromBoilerHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerBoilerHorsepower)
}

// FromHydraulicHorsepower creates a new Power instance from a value in HydraulicHorsepower.
func (uf PowerFactory) FromHydraulicHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerHydraulicHorsepower)
}

// FromBritishThermalUnitsPerHour creates a new Power instance from a value in BritishThermalUnitsPerHour.
func (uf PowerFactory) FromBritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerBritishThermalUnitPerHour)
}

// FromJoulesPerHour creates a new Power instance from a value in JoulesPerHour.
func (uf PowerFactory) FromJoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerJoulePerHour)
}

// FromTonsOfRefrigeration creates a new Power instance from a value in TonsOfRefrigeration.
func (uf PowerFactory) FromTonsOfRefrigeration(value float64) (*Power, error) {
	return newPower(value, PowerTonOfRefrigeration)
}

// FromFemtowatts creates a new Power instance from a value in Femtowatts.
func (uf PowerFactory) FromFemtowatts(value float64) (*Power, error) {
	return newPower(value, PowerFemtowatt)
}

// FromPicowatts creates a new Power instance from a value in Picowatts.
func (uf PowerFactory) FromPicowatts(value float64) (*Power, error) {
	return newPower(value, PowerPicowatt)
}

// FromNanowatts creates a new Power instance from a value in Nanowatts.
func (uf PowerFactory) FromNanowatts(value float64) (*Power, error) {
	return newPower(value, PowerNanowatt)
}

// FromMicrowatts creates a new Power instance from a value in Microwatts.
func (uf PowerFactory) FromMicrowatts(value float64) (*Power, error) {
	return newPower(value, PowerMicrowatt)
}

// FromMilliwatts creates a new Power instance from a value in Milliwatts.
func (uf PowerFactory) FromMilliwatts(value float64) (*Power, error) {
	return newPower(value, PowerMilliwatt)
}

// FromDeciwatts creates a new Power instance from a value in Deciwatts.
func (uf PowerFactory) FromDeciwatts(value float64) (*Power, error) {
	return newPower(value, PowerDeciwatt)
}

// FromDecawatts creates a new Power instance from a value in Decawatts.
func (uf PowerFactory) FromDecawatts(value float64) (*Power, error) {
	return newPower(value, PowerDecawatt)
}

// FromKilowatts creates a new Power instance from a value in Kilowatts.
func (uf PowerFactory) FromKilowatts(value float64) (*Power, error) {
	return newPower(value, PowerKilowatt)
}

// FromMegawatts creates a new Power instance from a value in Megawatts.
func (uf PowerFactory) FromMegawatts(value float64) (*Power, error) {
	return newPower(value, PowerMegawatt)
}

// FromGigawatts creates a new Power instance from a value in Gigawatts.
func (uf PowerFactory) FromGigawatts(value float64) (*Power, error) {
	return newPower(value, PowerGigawatt)
}

// FromTerawatts creates a new Power instance from a value in Terawatts.
func (uf PowerFactory) FromTerawatts(value float64) (*Power, error) {
	return newPower(value, PowerTerawatt)
}

// FromPetawatts creates a new Power instance from a value in Petawatts.
func (uf PowerFactory) FromPetawatts(value float64) (*Power, error) {
	return newPower(value, PowerPetawatt)
}

// FromKilobritishThermalUnitsPerHour creates a new Power instance from a value in KilobritishThermalUnitsPerHour.
func (uf PowerFactory) FromKilobritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerKilobritishThermalUnitPerHour)
}

// FromMegabritishThermalUnitsPerHour creates a new Power instance from a value in MegabritishThermalUnitsPerHour.
func (uf PowerFactory) FromMegabritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMegabritishThermalUnitPerHour)
}

// FromMillijoulesPerHour creates a new Power instance from a value in MillijoulesPerHour.
func (uf PowerFactory) FromMillijoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMillijoulePerHour)
}

// FromKilojoulesPerHour creates a new Power instance from a value in KilojoulesPerHour.
func (uf PowerFactory) FromKilojoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerKilojoulePerHour)
}

// FromMegajoulesPerHour creates a new Power instance from a value in MegajoulesPerHour.
func (uf PowerFactory) FromMegajoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMegajoulePerHour)
}

// FromGigajoulesPerHour creates a new Power instance from a value in GigajoulesPerHour.
func (uf PowerFactory) FromGigajoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerGigajoulePerHour)
}


// newPower creates a new Power.
func newPower(value float64, fromUnit PowerUnits) (*Power, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &Power{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of Power in Watt unit (the base/default quantity).
func (a *Power) BaseValue() float64 {
	return a.value
}


// Watts returns the Power value in Watts.
//
// 
func (a *Power) Watts() float64 {
	if a.wattsLazy != nil {
		return *a.wattsLazy
	}
	watts := a.convertFromBase(PowerWatt)
	a.wattsLazy = &watts
	return watts
}

// MechanicalHorsepower returns the Power value in MechanicalHorsepower.
//
// 
func (a *Power) MechanicalHorsepower() float64 {
	if a.mechanical_horsepowerLazy != nil {
		return *a.mechanical_horsepowerLazy
	}
	mechanical_horsepower := a.convertFromBase(PowerMechanicalHorsepower)
	a.mechanical_horsepowerLazy = &mechanical_horsepower
	return mechanical_horsepower
}

// MetricHorsepower returns the Power value in MetricHorsepower.
//
// 
func (a *Power) MetricHorsepower() float64 {
	if a.metric_horsepowerLazy != nil {
		return *a.metric_horsepowerLazy
	}
	metric_horsepower := a.convertFromBase(PowerMetricHorsepower)
	a.metric_horsepowerLazy = &metric_horsepower
	return metric_horsepower
}

// ElectricalHorsepower returns the Power value in ElectricalHorsepower.
//
// 
func (a *Power) ElectricalHorsepower() float64 {
	if a.electrical_horsepowerLazy != nil {
		return *a.electrical_horsepowerLazy
	}
	electrical_horsepower := a.convertFromBase(PowerElectricalHorsepower)
	a.electrical_horsepowerLazy = &electrical_horsepower
	return electrical_horsepower
}

// BoilerHorsepower returns the Power value in BoilerHorsepower.
//
// 
func (a *Power) BoilerHorsepower() float64 {
	if a.boiler_horsepowerLazy != nil {
		return *a.boiler_horsepowerLazy
	}
	boiler_horsepower := a.convertFromBase(PowerBoilerHorsepower)
	a.boiler_horsepowerLazy = &boiler_horsepower
	return boiler_horsepower
}

// HydraulicHorsepower returns the Power value in HydraulicHorsepower.
//
// 
func (a *Power) HydraulicHorsepower() float64 {
	if a.hydraulic_horsepowerLazy != nil {
		return *a.hydraulic_horsepowerLazy
	}
	hydraulic_horsepower := a.convertFromBase(PowerHydraulicHorsepower)
	a.hydraulic_horsepowerLazy = &hydraulic_horsepower
	return hydraulic_horsepower
}

// BritishThermalUnitsPerHour returns the Power value in BritishThermalUnitsPerHour.
//
// 
func (a *Power) BritishThermalUnitsPerHour() float64 {
	if a.british_thermal_units_per_hourLazy != nil {
		return *a.british_thermal_units_per_hourLazy
	}
	british_thermal_units_per_hour := a.convertFromBase(PowerBritishThermalUnitPerHour)
	a.british_thermal_units_per_hourLazy = &british_thermal_units_per_hour
	return british_thermal_units_per_hour
}

// JoulesPerHour returns the Power value in JoulesPerHour.
//
// 
func (a *Power) JoulesPerHour() float64 {
	if a.joules_per_hourLazy != nil {
		return *a.joules_per_hourLazy
	}
	joules_per_hour := a.convertFromBase(PowerJoulePerHour)
	a.joules_per_hourLazy = &joules_per_hour
	return joules_per_hour
}

// TonsOfRefrigeration returns the Power value in TonsOfRefrigeration.
//
// 
func (a *Power) TonsOfRefrigeration() float64 {
	if a.tons_of_refrigerationLazy != nil {
		return *a.tons_of_refrigerationLazy
	}
	tons_of_refrigeration := a.convertFromBase(PowerTonOfRefrigeration)
	a.tons_of_refrigerationLazy = &tons_of_refrigeration
	return tons_of_refrigeration
}

// Femtowatts returns the Power value in Femtowatts.
//
// 
func (a *Power) Femtowatts() float64 {
	if a.femtowattsLazy != nil {
		return *a.femtowattsLazy
	}
	femtowatts := a.convertFromBase(PowerFemtowatt)
	a.femtowattsLazy = &femtowatts
	return femtowatts
}

// Picowatts returns the Power value in Picowatts.
//
// 
func (a *Power) Picowatts() float64 {
	if a.picowattsLazy != nil {
		return *a.picowattsLazy
	}
	picowatts := a.convertFromBase(PowerPicowatt)
	a.picowattsLazy = &picowatts
	return picowatts
}

// Nanowatts returns the Power value in Nanowatts.
//
// 
func (a *Power) Nanowatts() float64 {
	if a.nanowattsLazy != nil {
		return *a.nanowattsLazy
	}
	nanowatts := a.convertFromBase(PowerNanowatt)
	a.nanowattsLazy = &nanowatts
	return nanowatts
}

// Microwatts returns the Power value in Microwatts.
//
// 
func (a *Power) Microwatts() float64 {
	if a.microwattsLazy != nil {
		return *a.microwattsLazy
	}
	microwatts := a.convertFromBase(PowerMicrowatt)
	a.microwattsLazy = &microwatts
	return microwatts
}

// Milliwatts returns the Power value in Milliwatts.
//
// 
func (a *Power) Milliwatts() float64 {
	if a.milliwattsLazy != nil {
		return *a.milliwattsLazy
	}
	milliwatts := a.convertFromBase(PowerMilliwatt)
	a.milliwattsLazy = &milliwatts
	return milliwatts
}

// Deciwatts returns the Power value in Deciwatts.
//
// 
func (a *Power) Deciwatts() float64 {
	if a.deciwattsLazy != nil {
		return *a.deciwattsLazy
	}
	deciwatts := a.convertFromBase(PowerDeciwatt)
	a.deciwattsLazy = &deciwatts
	return deciwatts
}

// Decawatts returns the Power value in Decawatts.
//
// 
func (a *Power) Decawatts() float64 {
	if a.decawattsLazy != nil {
		return *a.decawattsLazy
	}
	decawatts := a.convertFromBase(PowerDecawatt)
	a.decawattsLazy = &decawatts
	return decawatts
}

// Kilowatts returns the Power value in Kilowatts.
//
// 
func (a *Power) Kilowatts() float64 {
	if a.kilowattsLazy != nil {
		return *a.kilowattsLazy
	}
	kilowatts := a.convertFromBase(PowerKilowatt)
	a.kilowattsLazy = &kilowatts
	return kilowatts
}

// Megawatts returns the Power value in Megawatts.
//
// 
func (a *Power) Megawatts() float64 {
	if a.megawattsLazy != nil {
		return *a.megawattsLazy
	}
	megawatts := a.convertFromBase(PowerMegawatt)
	a.megawattsLazy = &megawatts
	return megawatts
}

// Gigawatts returns the Power value in Gigawatts.
//
// 
func (a *Power) Gigawatts() float64 {
	if a.gigawattsLazy != nil {
		return *a.gigawattsLazy
	}
	gigawatts := a.convertFromBase(PowerGigawatt)
	a.gigawattsLazy = &gigawatts
	return gigawatts
}

// Terawatts returns the Power value in Terawatts.
//
// 
func (a *Power) Terawatts() float64 {
	if a.terawattsLazy != nil {
		return *a.terawattsLazy
	}
	terawatts := a.convertFromBase(PowerTerawatt)
	a.terawattsLazy = &terawatts
	return terawatts
}

// Petawatts returns the Power value in Petawatts.
//
// 
func (a *Power) Petawatts() float64 {
	if a.petawattsLazy != nil {
		return *a.petawattsLazy
	}
	petawatts := a.convertFromBase(PowerPetawatt)
	a.petawattsLazy = &petawatts
	return petawatts
}

// KilobritishThermalUnitsPerHour returns the Power value in KilobritishThermalUnitsPerHour.
//
// 
func (a *Power) KilobritishThermalUnitsPerHour() float64 {
	if a.kilobritish_thermal_units_per_hourLazy != nil {
		return *a.kilobritish_thermal_units_per_hourLazy
	}
	kilobritish_thermal_units_per_hour := a.convertFromBase(PowerKilobritishThermalUnitPerHour)
	a.kilobritish_thermal_units_per_hourLazy = &kilobritish_thermal_units_per_hour
	return kilobritish_thermal_units_per_hour
}

// MegabritishThermalUnitsPerHour returns the Power value in MegabritishThermalUnitsPerHour.
//
// 
func (a *Power) MegabritishThermalUnitsPerHour() float64 {
	if a.megabritish_thermal_units_per_hourLazy != nil {
		return *a.megabritish_thermal_units_per_hourLazy
	}
	megabritish_thermal_units_per_hour := a.convertFromBase(PowerMegabritishThermalUnitPerHour)
	a.megabritish_thermal_units_per_hourLazy = &megabritish_thermal_units_per_hour
	return megabritish_thermal_units_per_hour
}

// MillijoulesPerHour returns the Power value in MillijoulesPerHour.
//
// 
func (a *Power) MillijoulesPerHour() float64 {
	if a.millijoules_per_hourLazy != nil {
		return *a.millijoules_per_hourLazy
	}
	millijoules_per_hour := a.convertFromBase(PowerMillijoulePerHour)
	a.millijoules_per_hourLazy = &millijoules_per_hour
	return millijoules_per_hour
}

// KilojoulesPerHour returns the Power value in KilojoulesPerHour.
//
// 
func (a *Power) KilojoulesPerHour() float64 {
	if a.kilojoules_per_hourLazy != nil {
		return *a.kilojoules_per_hourLazy
	}
	kilojoules_per_hour := a.convertFromBase(PowerKilojoulePerHour)
	a.kilojoules_per_hourLazy = &kilojoules_per_hour
	return kilojoules_per_hour
}

// MegajoulesPerHour returns the Power value in MegajoulesPerHour.
//
// 
func (a *Power) MegajoulesPerHour() float64 {
	if a.megajoules_per_hourLazy != nil {
		return *a.megajoules_per_hourLazy
	}
	megajoules_per_hour := a.convertFromBase(PowerMegajoulePerHour)
	a.megajoules_per_hourLazy = &megajoules_per_hour
	return megajoules_per_hour
}

// GigajoulesPerHour returns the Power value in GigajoulesPerHour.
//
// 
func (a *Power) GigajoulesPerHour() float64 {
	if a.gigajoules_per_hourLazy != nil {
		return *a.gigajoules_per_hourLazy
	}
	gigajoules_per_hour := a.convertFromBase(PowerGigajoulePerHour)
	a.gigajoules_per_hourLazy = &gigajoules_per_hour
	return gigajoules_per_hour
}


// ToDto creates a PowerDto representation from the Power instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Watt by default.
func (a *Power) ToDto(holdInUnit *PowerUnits) PowerDto {
	if holdInUnit == nil {
		defaultUnit := PowerWatt // Default value
		holdInUnit = &defaultUnit
	}

	return PowerDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates a JSON representation of the Power instance.
//
// If the provided holdInUnit is nil, the value will be repesented by Watt by default.
func (a *Power) ToDtoJSON(holdInUnit *PowerUnits) ([]byte, error) {
	// Convert to PowerDto and then serialize to JSON
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts a Power to a specific unit value.
// The function uses the provided unit type (PowerUnits) to return the corresponding value in the target unit.
// 
// Returns:
//    float64: The converted value in the target unit.
func (a *Power) Convert(toUnit PowerUnits) float64 {
	switch toUnit { 
    case PowerWatt:
		return a.Watts()
    case PowerMechanicalHorsepower:
		return a.MechanicalHorsepower()
    case PowerMetricHorsepower:
		return a.MetricHorsepower()
    case PowerElectricalHorsepower:
		return a.ElectricalHorsepower()
    case PowerBoilerHorsepower:
		return a.BoilerHorsepower()
    case PowerHydraulicHorsepower:
		return a.HydraulicHorsepower()
    case PowerBritishThermalUnitPerHour:
		return a.BritishThermalUnitsPerHour()
    case PowerJoulePerHour:
		return a.JoulesPerHour()
    case PowerTonOfRefrigeration:
		return a.TonsOfRefrigeration()
    case PowerFemtowatt:
		return a.Femtowatts()
    case PowerPicowatt:
		return a.Picowatts()
    case PowerNanowatt:
		return a.Nanowatts()
    case PowerMicrowatt:
		return a.Microwatts()
    case PowerMilliwatt:
		return a.Milliwatts()
    case PowerDeciwatt:
		return a.Deciwatts()
    case PowerDecawatt:
		return a.Decawatts()
    case PowerKilowatt:
		return a.Kilowatts()
    case PowerMegawatt:
		return a.Megawatts()
    case PowerGigawatt:
		return a.Gigawatts()
    case PowerTerawatt:
		return a.Terawatts()
    case PowerPetawatt:
		return a.Petawatts()
    case PowerKilobritishThermalUnitPerHour:
		return a.KilobritishThermalUnitsPerHour()
    case PowerMegabritishThermalUnitPerHour:
		return a.MegabritishThermalUnitsPerHour()
    case PowerMillijoulePerHour:
		return a.MillijoulesPerHour()
    case PowerKilojoulePerHour:
		return a.KilojoulesPerHour()
    case PowerMegajoulePerHour:
		return a.MegajoulesPerHour()
    case PowerGigajoulePerHour:
		return a.GigajoulesPerHour()
	default:
		return math.NaN()
	}
}

func (a *Power) convertFromBase(toUnit PowerUnits) float64 {
    value := a.value
	switch toUnit { 
	case PowerWatt:
		return (value) 
	case PowerMechanicalHorsepower:
		return (value / 745.69) 
	case PowerMetricHorsepower:
		return (value / 735.49875) 
	case PowerElectricalHorsepower:
		return (value / 746) 
	case PowerBoilerHorsepower:
		return (value / 9812.5) 
	case PowerHydraulicHorsepower:
		return (value / 745.69988145) 
	case PowerBritishThermalUnitPerHour:
		return (value / 0.29307107017) 
	case PowerJoulePerHour:
		return (value * 3600) 
	case PowerTonOfRefrigeration:
		return (value / 3516.853) 
	case PowerFemtowatt:
		return ((value) / 1e-15) 
	case PowerPicowatt:
		return ((value) / 1e-12) 
	case PowerNanowatt:
		return ((value) / 1e-09) 
	case PowerMicrowatt:
		return ((value) / 1e-06) 
	case PowerMilliwatt:
		return ((value) / 0.001) 
	case PowerDeciwatt:
		return ((value) / 0.1) 
	case PowerDecawatt:
		return ((value) / 10.0) 
	case PowerKilowatt:
		return ((value) / 1000.0) 
	case PowerMegawatt:
		return ((value) / 1000000.0) 
	case PowerGigawatt:
		return ((value) / 1000000000.0) 
	case PowerTerawatt:
		return ((value) / 1000000000000.0) 
	case PowerPetawatt:
		return ((value) / 1000000000000000.0) 
	case PowerKilobritishThermalUnitPerHour:
		return ((value / 0.29307107017) / 1000.0) 
	case PowerMegabritishThermalUnitPerHour:
		return ((value / 0.29307107017) / 1000000.0) 
	case PowerMillijoulePerHour:
		return ((value * 3600) / 0.001) 
	case PowerKilojoulePerHour:
		return ((value * 3600) / 1000.0) 
	case PowerMegajoulePerHour:
		return ((value * 3600) / 1000000.0) 
	case PowerGigajoulePerHour:
		return ((value * 3600) / 1000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *Power) convertToBase(value float64, fromUnit PowerUnits) float64 {
	switch fromUnit { 
	case PowerWatt:
		return (value) 
	case PowerMechanicalHorsepower:
		return (value * 745.69) 
	case PowerMetricHorsepower:
		return (value * 735.49875) 
	case PowerElectricalHorsepower:
		return (value * 746) 
	case PowerBoilerHorsepower:
		return (value * 9812.5) 
	case PowerHydraulicHorsepower:
		return (value * 745.69988145) 
	case PowerBritishThermalUnitPerHour:
		return (value * 0.29307107017) 
	case PowerJoulePerHour:
		return (value / 3600) 
	case PowerTonOfRefrigeration:
		return (value * 3516.853) 
	case PowerFemtowatt:
		return ((value) * 1e-15) 
	case PowerPicowatt:
		return ((value) * 1e-12) 
	case PowerNanowatt:
		return ((value) * 1e-09) 
	case PowerMicrowatt:
		return ((value) * 1e-06) 
	case PowerMilliwatt:
		return ((value) * 0.001) 
	case PowerDeciwatt:
		return ((value) * 0.1) 
	case PowerDecawatt:
		return ((value) * 10.0) 
	case PowerKilowatt:
		return ((value) * 1000.0) 
	case PowerMegawatt:
		return ((value) * 1000000.0) 
	case PowerGigawatt:
		return ((value) * 1000000000.0) 
	case PowerTerawatt:
		return ((value) * 1000000000000.0) 
	case PowerPetawatt:
		return ((value) * 1000000000000000.0) 
	case PowerKilobritishThermalUnitPerHour:
		return ((value * 0.29307107017) * 1000.0) 
	case PowerMegabritishThermalUnitPerHour:
		return ((value * 0.29307107017) * 1000000.0) 
	case PowerMillijoulePerHour:
		return ((value / 3600) * 0.001) 
	case PowerKilojoulePerHour:
		return ((value / 3600) * 1000.0) 
	case PowerMegajoulePerHour:
		return ((value / 3600) * 1000000.0) 
	case PowerGigajoulePerHour:
		return ((value / 3600) * 1000000000.0) 
	default:
		return math.NaN()
	}
}

// String returns a string representation of the Power in the default unit (Watt),
// formatted to two decimal places.
func (a Power) String() string {
	return a.ToString(PowerWatt, 2)
}

// ToString formats the Power value as a string with the specified unit and fractional digits.
// It converts the Power to the specified unit and returns the formatted value with the appropriate unit abbreviation.
// 
// Parameters:
//    unit: The unit to which the Power value will be converted (e.g., Watt))
//    fractionalDigits: The number of digits to show after the decimal point. 
//                       If fractionalDigits is -1, it uses the most compact format without rounding or padding.
// 
// Returns:
//    string: The formatted string representing the Power with the unit abbreviation.
func (a *Power) ToString(unit PowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf("%s %s", formatted ,GetPowerAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, GetPowerAbbreviation(unit))
}

// Equals checks if the given Power is equal to the current Power.
//
// Parameters:
//    other: The Power to compare against.
//
// Returns:
//    bool: Returns true if both Power are equal, false otherwise.
func (a *Power) Equals(other *Power) bool {
	return a.value == other.BaseValue()
}

// CompareTo compares the current Power with another Power.
// It returns -1 if the current Power is less than the other Power, 
// 1 if it is greater, and 0 if they are equal.
//
// Parameters:
//    other: The Power to compare against.
//
// Returns:
//    int: -1 if the current Power is less, 1 if greater, and 0 if equal.
func (a *Power) CompareTo(other *Power) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add adds the given Power to the current Power and returns the result.
// The result is a new Power instance with the sum of the values.
//
// Parameters:
//    other: The Power to add to the current Power.
//
// Returns:
//    *Power: A new Power instance representing the sum of both Power.
func (a *Power) Add(other *Power) *Power {
	return &Power{value: a.value + other.BaseValue()}
}

// Subtract subtracts the given Power from the current Power and returns the result.
// The result is a new Power instance with the difference of the values.
//
// Parameters:
//    other: The Power to subtract from the current Power.
//
// Returns:
//    *Power: A new Power instance representing the difference of both Power.
func (a *Power) Subtract(other *Power) *Power {
	return &Power{value: a.value - other.BaseValue()}
}

// Multiply multiplies the current Power by the given Power and returns the result.
// The result is a new Power instance with the product of the values.
//
// Parameters:
//    other: The Power to multiply with the current Power.
//
// Returns:
//    *Power: A new Power instance representing the product of both Power.
func (a *Power) Multiply(other *Power) *Power {
	return &Power{value: a.value * other.BaseValue()}
}

// Divide divides the current Power by the given Power and returns the result.
// The result is a new Power instance with the quotient of the values.
//
// Parameters:
//    other: The Power to divide the current Power by.
//
// Returns:
//    *Power: A new Power instance representing the quotient of both Power.
func (a *Power) Divide(other *Power) *Power {
	return &Power{value: a.value / other.BaseValue()}
}

// GetPowerAbbreviation gets the unit abbreviation.
func GetPowerAbbreviation(unit PowerUnits) string {
	switch unit { 
	case PowerWatt:
		return "W" 
	case PowerMechanicalHorsepower:
		return "hp(I)" 
	case PowerMetricHorsepower:
		return "hp(M)" 
	case PowerElectricalHorsepower:
		return "hp(E)" 
	case PowerBoilerHorsepower:
		return "hp(S)" 
	case PowerHydraulicHorsepower:
		return "hp(H)" 
	case PowerBritishThermalUnitPerHour:
		return "Btu/h" 
	case PowerJoulePerHour:
		return "J/h" 
	case PowerTonOfRefrigeration:
		return "TR" 
	case PowerFemtowatt:
		return "fW" 
	case PowerPicowatt:
		return "pW" 
	case PowerNanowatt:
		return "nW" 
	case PowerMicrowatt:
		return "μW" 
	case PowerMilliwatt:
		return "mW" 
	case PowerDeciwatt:
		return "dW" 
	case PowerDecawatt:
		return "daW" 
	case PowerKilowatt:
		return "kW" 
	case PowerMegawatt:
		return "MW" 
	case PowerGigawatt:
		return "GW" 
	case PowerTerawatt:
		return "TW" 
	case PowerPetawatt:
		return "PW" 
	case PowerKilobritishThermalUnitPerHour:
		return "kBtu/h" 
	case PowerMegabritishThermalUnitPerHour:
		return "MBtu/h" 
	case PowerMillijoulePerHour:
		return "mJ/h" 
	case PowerKilojoulePerHour:
		return "kJ/h" 
	case PowerMegajoulePerHour:
		return "MJ/h" 
	case PowerGigajoulePerHour:
		return "GJ/h" 
	default:
		return ""
	}
}