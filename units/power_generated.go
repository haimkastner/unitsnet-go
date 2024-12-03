// Code generated - DO NOT EDIT.

package units

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// PowerUnits enumeration
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

// PowerDto represents an Power
type PowerDto struct {
	Value float64
	Unit  PowerUnits
}

// PowerDtoFactory struct to group related functions
type PowerDtoFactory struct{}

func (udf PowerDtoFactory) FromJSON(data []byte) (*PowerDto, error) {
	a := PowerDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a PowerDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// Power struct
type Power struct {
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

// PowerFactory struct to group related functions
type PowerFactory struct{}

func (uf PowerFactory) CreatePower(value float64, unit PowerUnits) (*Power, error) {
	return newPower(value, unit)
}

func (uf PowerFactory) FromDto(dto PowerDto) (*Power, error) {
	return newPower(dto.Value, dto.Unit)
}

func (uf PowerFactory) FromDtoJSON(data []byte) (*Power, error) {
	unitDto, err := PowerDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return PowerFactory{}.FromDto(*unitDto)
}


// FromWatt creates a new Power instance from Watt.
func (uf PowerFactory) FromWatts(value float64) (*Power, error) {
	return newPower(value, PowerWatt)
}

// FromMechanicalHorsepower creates a new Power instance from MechanicalHorsepower.
func (uf PowerFactory) FromMechanicalHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerMechanicalHorsepower)
}

// FromMetricHorsepower creates a new Power instance from MetricHorsepower.
func (uf PowerFactory) FromMetricHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerMetricHorsepower)
}

// FromElectricalHorsepower creates a new Power instance from ElectricalHorsepower.
func (uf PowerFactory) FromElectricalHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerElectricalHorsepower)
}

// FromBoilerHorsepower creates a new Power instance from BoilerHorsepower.
func (uf PowerFactory) FromBoilerHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerBoilerHorsepower)
}

// FromHydraulicHorsepower creates a new Power instance from HydraulicHorsepower.
func (uf PowerFactory) FromHydraulicHorsepower(value float64) (*Power, error) {
	return newPower(value, PowerHydraulicHorsepower)
}

// FromBritishThermalUnitPerHour creates a new Power instance from BritishThermalUnitPerHour.
func (uf PowerFactory) FromBritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerBritishThermalUnitPerHour)
}

// FromJoulePerHour creates a new Power instance from JoulePerHour.
func (uf PowerFactory) FromJoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerJoulePerHour)
}

// FromTonOfRefrigeration creates a new Power instance from TonOfRefrigeration.
func (uf PowerFactory) FromTonsOfRefrigeration(value float64) (*Power, error) {
	return newPower(value, PowerTonOfRefrigeration)
}

// FromFemtowatt creates a new Power instance from Femtowatt.
func (uf PowerFactory) FromFemtowatts(value float64) (*Power, error) {
	return newPower(value, PowerFemtowatt)
}

// FromPicowatt creates a new Power instance from Picowatt.
func (uf PowerFactory) FromPicowatts(value float64) (*Power, error) {
	return newPower(value, PowerPicowatt)
}

// FromNanowatt creates a new Power instance from Nanowatt.
func (uf PowerFactory) FromNanowatts(value float64) (*Power, error) {
	return newPower(value, PowerNanowatt)
}

// FromMicrowatt creates a new Power instance from Microwatt.
func (uf PowerFactory) FromMicrowatts(value float64) (*Power, error) {
	return newPower(value, PowerMicrowatt)
}

// FromMilliwatt creates a new Power instance from Milliwatt.
func (uf PowerFactory) FromMilliwatts(value float64) (*Power, error) {
	return newPower(value, PowerMilliwatt)
}

// FromDeciwatt creates a new Power instance from Deciwatt.
func (uf PowerFactory) FromDeciwatts(value float64) (*Power, error) {
	return newPower(value, PowerDeciwatt)
}

// FromDecawatt creates a new Power instance from Decawatt.
func (uf PowerFactory) FromDecawatts(value float64) (*Power, error) {
	return newPower(value, PowerDecawatt)
}

// FromKilowatt creates a new Power instance from Kilowatt.
func (uf PowerFactory) FromKilowatts(value float64) (*Power, error) {
	return newPower(value, PowerKilowatt)
}

// FromMegawatt creates a new Power instance from Megawatt.
func (uf PowerFactory) FromMegawatts(value float64) (*Power, error) {
	return newPower(value, PowerMegawatt)
}

// FromGigawatt creates a new Power instance from Gigawatt.
func (uf PowerFactory) FromGigawatts(value float64) (*Power, error) {
	return newPower(value, PowerGigawatt)
}

// FromTerawatt creates a new Power instance from Terawatt.
func (uf PowerFactory) FromTerawatts(value float64) (*Power, error) {
	return newPower(value, PowerTerawatt)
}

// FromPetawatt creates a new Power instance from Petawatt.
func (uf PowerFactory) FromPetawatts(value float64) (*Power, error) {
	return newPower(value, PowerPetawatt)
}

// FromKilobritishThermalUnitPerHour creates a new Power instance from KilobritishThermalUnitPerHour.
func (uf PowerFactory) FromKilobritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerKilobritishThermalUnitPerHour)
}

// FromMegabritishThermalUnitPerHour creates a new Power instance from MegabritishThermalUnitPerHour.
func (uf PowerFactory) FromMegabritishThermalUnitsPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMegabritishThermalUnitPerHour)
}

// FromMillijoulePerHour creates a new Power instance from MillijoulePerHour.
func (uf PowerFactory) FromMillijoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMillijoulePerHour)
}

// FromKilojoulePerHour creates a new Power instance from KilojoulePerHour.
func (uf PowerFactory) FromKilojoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerKilojoulePerHour)
}

// FromMegajoulePerHour creates a new Power instance from MegajoulePerHour.
func (uf PowerFactory) FromMegajoulesPerHour(value float64) (*Power, error) {
	return newPower(value, PowerMegajoulePerHour)
}

// FromGigajoulePerHour creates a new Power instance from GigajoulePerHour.
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

// BaseValue returns the base value of Power in Watt.
func (a *Power) BaseValue() float64 {
	return a.value
}


// Watt returns the value in Watt.
func (a *Power) Watts() float64 {
	if a.wattsLazy != nil {
		return *a.wattsLazy
	}
	watts := a.convertFromBase(PowerWatt)
	a.wattsLazy = &watts
	return watts
}

// MechanicalHorsepower returns the value in MechanicalHorsepower.
func (a *Power) MechanicalHorsepower() float64 {
	if a.mechanical_horsepowerLazy != nil {
		return *a.mechanical_horsepowerLazy
	}
	mechanical_horsepower := a.convertFromBase(PowerMechanicalHorsepower)
	a.mechanical_horsepowerLazy = &mechanical_horsepower
	return mechanical_horsepower
}

// MetricHorsepower returns the value in MetricHorsepower.
func (a *Power) MetricHorsepower() float64 {
	if a.metric_horsepowerLazy != nil {
		return *a.metric_horsepowerLazy
	}
	metric_horsepower := a.convertFromBase(PowerMetricHorsepower)
	a.metric_horsepowerLazy = &metric_horsepower
	return metric_horsepower
}

// ElectricalHorsepower returns the value in ElectricalHorsepower.
func (a *Power) ElectricalHorsepower() float64 {
	if a.electrical_horsepowerLazy != nil {
		return *a.electrical_horsepowerLazy
	}
	electrical_horsepower := a.convertFromBase(PowerElectricalHorsepower)
	a.electrical_horsepowerLazy = &electrical_horsepower
	return electrical_horsepower
}

// BoilerHorsepower returns the value in BoilerHorsepower.
func (a *Power) BoilerHorsepower() float64 {
	if a.boiler_horsepowerLazy != nil {
		return *a.boiler_horsepowerLazy
	}
	boiler_horsepower := a.convertFromBase(PowerBoilerHorsepower)
	a.boiler_horsepowerLazy = &boiler_horsepower
	return boiler_horsepower
}

// HydraulicHorsepower returns the value in HydraulicHorsepower.
func (a *Power) HydraulicHorsepower() float64 {
	if a.hydraulic_horsepowerLazy != nil {
		return *a.hydraulic_horsepowerLazy
	}
	hydraulic_horsepower := a.convertFromBase(PowerHydraulicHorsepower)
	a.hydraulic_horsepowerLazy = &hydraulic_horsepower
	return hydraulic_horsepower
}

// BritishThermalUnitPerHour returns the value in BritishThermalUnitPerHour.
func (a *Power) BritishThermalUnitsPerHour() float64 {
	if a.british_thermal_units_per_hourLazy != nil {
		return *a.british_thermal_units_per_hourLazy
	}
	british_thermal_units_per_hour := a.convertFromBase(PowerBritishThermalUnitPerHour)
	a.british_thermal_units_per_hourLazy = &british_thermal_units_per_hour
	return british_thermal_units_per_hour
}

// JoulePerHour returns the value in JoulePerHour.
func (a *Power) JoulesPerHour() float64 {
	if a.joules_per_hourLazy != nil {
		return *a.joules_per_hourLazy
	}
	joules_per_hour := a.convertFromBase(PowerJoulePerHour)
	a.joules_per_hourLazy = &joules_per_hour
	return joules_per_hour
}

// TonOfRefrigeration returns the value in TonOfRefrigeration.
func (a *Power) TonsOfRefrigeration() float64 {
	if a.tons_of_refrigerationLazy != nil {
		return *a.tons_of_refrigerationLazy
	}
	tons_of_refrigeration := a.convertFromBase(PowerTonOfRefrigeration)
	a.tons_of_refrigerationLazy = &tons_of_refrigeration
	return tons_of_refrigeration
}

// Femtowatt returns the value in Femtowatt.
func (a *Power) Femtowatts() float64 {
	if a.femtowattsLazy != nil {
		return *a.femtowattsLazy
	}
	femtowatts := a.convertFromBase(PowerFemtowatt)
	a.femtowattsLazy = &femtowatts
	return femtowatts
}

// Picowatt returns the value in Picowatt.
func (a *Power) Picowatts() float64 {
	if a.picowattsLazy != nil {
		return *a.picowattsLazy
	}
	picowatts := a.convertFromBase(PowerPicowatt)
	a.picowattsLazy = &picowatts
	return picowatts
}

// Nanowatt returns the value in Nanowatt.
func (a *Power) Nanowatts() float64 {
	if a.nanowattsLazy != nil {
		return *a.nanowattsLazy
	}
	nanowatts := a.convertFromBase(PowerNanowatt)
	a.nanowattsLazy = &nanowatts
	return nanowatts
}

// Microwatt returns the value in Microwatt.
func (a *Power) Microwatts() float64 {
	if a.microwattsLazy != nil {
		return *a.microwattsLazy
	}
	microwatts := a.convertFromBase(PowerMicrowatt)
	a.microwattsLazy = &microwatts
	return microwatts
}

// Milliwatt returns the value in Milliwatt.
func (a *Power) Milliwatts() float64 {
	if a.milliwattsLazy != nil {
		return *a.milliwattsLazy
	}
	milliwatts := a.convertFromBase(PowerMilliwatt)
	a.milliwattsLazy = &milliwatts
	return milliwatts
}

// Deciwatt returns the value in Deciwatt.
func (a *Power) Deciwatts() float64 {
	if a.deciwattsLazy != nil {
		return *a.deciwattsLazy
	}
	deciwatts := a.convertFromBase(PowerDeciwatt)
	a.deciwattsLazy = &deciwatts
	return deciwatts
}

// Decawatt returns the value in Decawatt.
func (a *Power) Decawatts() float64 {
	if a.decawattsLazy != nil {
		return *a.decawattsLazy
	}
	decawatts := a.convertFromBase(PowerDecawatt)
	a.decawattsLazy = &decawatts
	return decawatts
}

// Kilowatt returns the value in Kilowatt.
func (a *Power) Kilowatts() float64 {
	if a.kilowattsLazy != nil {
		return *a.kilowattsLazy
	}
	kilowatts := a.convertFromBase(PowerKilowatt)
	a.kilowattsLazy = &kilowatts
	return kilowatts
}

// Megawatt returns the value in Megawatt.
func (a *Power) Megawatts() float64 {
	if a.megawattsLazy != nil {
		return *a.megawattsLazy
	}
	megawatts := a.convertFromBase(PowerMegawatt)
	a.megawattsLazy = &megawatts
	return megawatts
}

// Gigawatt returns the value in Gigawatt.
func (a *Power) Gigawatts() float64 {
	if a.gigawattsLazy != nil {
		return *a.gigawattsLazy
	}
	gigawatts := a.convertFromBase(PowerGigawatt)
	a.gigawattsLazy = &gigawatts
	return gigawatts
}

// Terawatt returns the value in Terawatt.
func (a *Power) Terawatts() float64 {
	if a.terawattsLazy != nil {
		return *a.terawattsLazy
	}
	terawatts := a.convertFromBase(PowerTerawatt)
	a.terawattsLazy = &terawatts
	return terawatts
}

// Petawatt returns the value in Petawatt.
func (a *Power) Petawatts() float64 {
	if a.petawattsLazy != nil {
		return *a.petawattsLazy
	}
	petawatts := a.convertFromBase(PowerPetawatt)
	a.petawattsLazy = &petawatts
	return petawatts
}

// KilobritishThermalUnitPerHour returns the value in KilobritishThermalUnitPerHour.
func (a *Power) KilobritishThermalUnitsPerHour() float64 {
	if a.kilobritish_thermal_units_per_hourLazy != nil {
		return *a.kilobritish_thermal_units_per_hourLazy
	}
	kilobritish_thermal_units_per_hour := a.convertFromBase(PowerKilobritishThermalUnitPerHour)
	a.kilobritish_thermal_units_per_hourLazy = &kilobritish_thermal_units_per_hour
	return kilobritish_thermal_units_per_hour
}

// MegabritishThermalUnitPerHour returns the value in MegabritishThermalUnitPerHour.
func (a *Power) MegabritishThermalUnitsPerHour() float64 {
	if a.megabritish_thermal_units_per_hourLazy != nil {
		return *a.megabritish_thermal_units_per_hourLazy
	}
	megabritish_thermal_units_per_hour := a.convertFromBase(PowerMegabritishThermalUnitPerHour)
	a.megabritish_thermal_units_per_hourLazy = &megabritish_thermal_units_per_hour
	return megabritish_thermal_units_per_hour
}

// MillijoulePerHour returns the value in MillijoulePerHour.
func (a *Power) MillijoulesPerHour() float64 {
	if a.millijoules_per_hourLazy != nil {
		return *a.millijoules_per_hourLazy
	}
	millijoules_per_hour := a.convertFromBase(PowerMillijoulePerHour)
	a.millijoules_per_hourLazy = &millijoules_per_hour
	return millijoules_per_hour
}

// KilojoulePerHour returns the value in KilojoulePerHour.
func (a *Power) KilojoulesPerHour() float64 {
	if a.kilojoules_per_hourLazy != nil {
		return *a.kilojoules_per_hourLazy
	}
	kilojoules_per_hour := a.convertFromBase(PowerKilojoulePerHour)
	a.kilojoules_per_hourLazy = &kilojoules_per_hour
	return kilojoules_per_hour
}

// MegajoulePerHour returns the value in MegajoulePerHour.
func (a *Power) MegajoulesPerHour() float64 {
	if a.megajoules_per_hourLazy != nil {
		return *a.megajoules_per_hourLazy
	}
	megajoules_per_hour := a.convertFromBase(PowerMegajoulePerHour)
	a.megajoules_per_hourLazy = &megajoules_per_hour
	return megajoules_per_hour
}

// GigajoulePerHour returns the value in GigajoulePerHour.
func (a *Power) GigajoulesPerHour() float64 {
	if a.gigajoules_per_hourLazy != nil {
		return *a.gigajoules_per_hourLazy
	}
	gigajoules_per_hour := a.convertFromBase(PowerGigajoulePerHour)
	a.gigajoules_per_hourLazy = &gigajoules_per_hour
	return gigajoules_per_hour
}


// ToDto creates an PowerDto representation.
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

// ToDtoJSON creates an PowerDto representation.
func (a *Power) ToDtoJSON(holdInUnit *PowerUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts Power to a specific unit value.
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
		return 0
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

// Implement the String() method for AngleDto
func (a Power) String() string {
	return a.ToString(PowerWatt, 2)
}

// ToString formats the Power to string.
// fractionalDigits -1 for not mention
func (a *Power) ToString(unit PowerUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *Power) getUnitAbbreviation(unit PowerUnits) string {
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

// Check if the given Power are equals to the current Power
func (a *Power) Equals(other *Power) bool {
	return a.value == other.BaseValue()
}

// Check if the given Power are equals to the current Power
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

// Add the given Power to the current Power.
func (a *Power) Add(other *Power) *Power {
	return &Power{value: a.value + other.BaseValue()}
}

// Subtract the given Power to the current Power.
func (a *Power) Subtract(other *Power) *Power {
	return &Power{value: a.value - other.BaseValue()}
}

// Multiply the given Power to the current Power.
func (a *Power) Multiply(other *Power) *Power {
	return &Power{value: a.value * other.BaseValue()}
}

// Divide the given Power to the current Power.
func (a *Power) Divide(other *Power) *Power {
	return &Power{value: a.value / other.BaseValue()}
}