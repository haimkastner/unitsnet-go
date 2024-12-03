package unitsnet_go

import (
	"encoding/json"
    "math"
	"errors"
	"fmt"
	"strconv"
)



// EnergyDensityUnits enumeration
type EnergyDensityUnits string

const (
    
        // 
        EnergyDensityJoulePerCubicMeter EnergyDensityUnits = "JoulePerCubicMeter"
        // 
        EnergyDensityWattHourPerCubicMeter EnergyDensityUnits = "WattHourPerCubicMeter"
        // 
        EnergyDensityKilojoulePerCubicMeter EnergyDensityUnits = "KilojoulePerCubicMeter"
        // 
        EnergyDensityMegajoulePerCubicMeter EnergyDensityUnits = "MegajoulePerCubicMeter"
        // 
        EnergyDensityGigajoulePerCubicMeter EnergyDensityUnits = "GigajoulePerCubicMeter"
        // 
        EnergyDensityTerajoulePerCubicMeter EnergyDensityUnits = "TerajoulePerCubicMeter"
        // 
        EnergyDensityPetajoulePerCubicMeter EnergyDensityUnits = "PetajoulePerCubicMeter"
        // 
        EnergyDensityKilowattHourPerCubicMeter EnergyDensityUnits = "KilowattHourPerCubicMeter"
        // 
        EnergyDensityMegawattHourPerCubicMeter EnergyDensityUnits = "MegawattHourPerCubicMeter"
        // 
        EnergyDensityGigawattHourPerCubicMeter EnergyDensityUnits = "GigawattHourPerCubicMeter"
        // 
        EnergyDensityTerawattHourPerCubicMeter EnergyDensityUnits = "TerawattHourPerCubicMeter"
        // 
        EnergyDensityPetawattHourPerCubicMeter EnergyDensityUnits = "PetawattHourPerCubicMeter"
)

// EnergyDensityDto represents an EnergyDensity
type EnergyDensityDto struct {
	Value float64
	Unit  EnergyDensityUnits
}

// EnergyDensityDtoFactory struct to group related functions
type EnergyDensityDtoFactory struct{}

func (udf EnergyDensityDtoFactory) FromJSON(data []byte) (*EnergyDensityDto, error) {
	a := EnergyDensityDto{}

	// Parse JSON into the temporary structure
	if err := json.Unmarshal(data, &a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a EnergyDensityDto) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}{
		Value: a.Value,
		Unit:  string(a.Unit),
	})
}




// EnergyDensity struct
type EnergyDensity struct {
	value       float64
    
    joules_per_cubic_meterLazy *float64 
    watt_hours_per_cubic_meterLazy *float64 
    kilojoules_per_cubic_meterLazy *float64 
    megajoules_per_cubic_meterLazy *float64 
    gigajoules_per_cubic_meterLazy *float64 
    terajoules_per_cubic_meterLazy *float64 
    petajoules_per_cubic_meterLazy *float64 
    kilowatt_hours_per_cubic_meterLazy *float64 
    megawatt_hours_per_cubic_meterLazy *float64 
    gigawatt_hours_per_cubic_meterLazy *float64 
    terawatt_hours_per_cubic_meterLazy *float64 
    petawatt_hours_per_cubic_meterLazy *float64 
}

// EnergyDensityFactory struct to group related functions
type EnergyDensityFactory struct{}

func (uf EnergyDensityFactory) CreateEnergyDensity(value float64, unit EnergyDensityUnits) (*EnergyDensity, error) {
	return newEnergyDensity(value, unit)
}

func (uf EnergyDensityFactory) FromDto(dto EnergyDensityDto) (*EnergyDensity, error) {
	return newEnergyDensity(dto.Value, dto.Unit)
}

func (uf EnergyDensityFactory) FromDtoJSON(data []byte) (*EnergyDensity, error) {
	unitDto, err := EnergyDensityDtoFactory{}.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return EnergyDensityFactory{}.FromDto(*unitDto)
}


// FromJoulePerCubicMeter creates a new EnergyDensity instance from JoulePerCubicMeter.
func (uf EnergyDensityFactory) FromJoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityJoulePerCubicMeter)
}

// FromWattHourPerCubicMeter creates a new EnergyDensity instance from WattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromWattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityWattHourPerCubicMeter)
}

// FromKilojoulePerCubicMeter creates a new EnergyDensity instance from KilojoulePerCubicMeter.
func (uf EnergyDensityFactory) FromKilojoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityKilojoulePerCubicMeter)
}

// FromMegajoulePerCubicMeter creates a new EnergyDensity instance from MegajoulePerCubicMeter.
func (uf EnergyDensityFactory) FromMegajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityMegajoulePerCubicMeter)
}

// FromGigajoulePerCubicMeter creates a new EnergyDensity instance from GigajoulePerCubicMeter.
func (uf EnergyDensityFactory) FromGigajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityGigajoulePerCubicMeter)
}

// FromTerajoulePerCubicMeter creates a new EnergyDensity instance from TerajoulePerCubicMeter.
func (uf EnergyDensityFactory) FromTerajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityTerajoulePerCubicMeter)
}

// FromPetajoulePerCubicMeter creates a new EnergyDensity instance from PetajoulePerCubicMeter.
func (uf EnergyDensityFactory) FromPetajoulesPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityPetajoulePerCubicMeter)
}

// FromKilowattHourPerCubicMeter creates a new EnergyDensity instance from KilowattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromKilowattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityKilowattHourPerCubicMeter)
}

// FromMegawattHourPerCubicMeter creates a new EnergyDensity instance from MegawattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromMegawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityMegawattHourPerCubicMeter)
}

// FromGigawattHourPerCubicMeter creates a new EnergyDensity instance from GigawattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromGigawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityGigawattHourPerCubicMeter)
}

// FromTerawattHourPerCubicMeter creates a new EnergyDensity instance from TerawattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromTerawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityTerawattHourPerCubicMeter)
}

// FromPetawattHourPerCubicMeter creates a new EnergyDensity instance from PetawattHourPerCubicMeter.
func (uf EnergyDensityFactory) FromPetawattHoursPerCubicMeter(value float64) (*EnergyDensity, error) {
	return newEnergyDensity(value, EnergyDensityPetawattHourPerCubicMeter)
}




// newEnergyDensity creates a new EnergyDensity.
func newEnergyDensity(value float64, fromUnit EnergyDensityUnits) (*EnergyDensity, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, errors.New("invalid unit value number")
	}
	a := &EnergyDensity{}
	a.value = a.convertToBase(value, fromUnit)
	return a, nil
}

// BaseValue returns the base value of EnergyDensity in JoulePerCubicMeter.
func (a *EnergyDensity) BaseValue() float64 {
	return a.value
}


// JoulePerCubicMeter returns the value in JoulePerCubicMeter.
func (a *EnergyDensity) JoulesPerCubicMeter() float64 {
	if a.joules_per_cubic_meterLazy != nil {
		return *a.joules_per_cubic_meterLazy
	}
	joules_per_cubic_meter := a.convertFromBase(EnergyDensityJoulePerCubicMeter)
	a.joules_per_cubic_meterLazy = &joules_per_cubic_meter
	return joules_per_cubic_meter
}

// WattHourPerCubicMeter returns the value in WattHourPerCubicMeter.
func (a *EnergyDensity) WattHoursPerCubicMeter() float64 {
	if a.watt_hours_per_cubic_meterLazy != nil {
		return *a.watt_hours_per_cubic_meterLazy
	}
	watt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityWattHourPerCubicMeter)
	a.watt_hours_per_cubic_meterLazy = &watt_hours_per_cubic_meter
	return watt_hours_per_cubic_meter
}

// KilojoulePerCubicMeter returns the value in KilojoulePerCubicMeter.
func (a *EnergyDensity) KilojoulesPerCubicMeter() float64 {
	if a.kilojoules_per_cubic_meterLazy != nil {
		return *a.kilojoules_per_cubic_meterLazy
	}
	kilojoules_per_cubic_meter := a.convertFromBase(EnergyDensityKilojoulePerCubicMeter)
	a.kilojoules_per_cubic_meterLazy = &kilojoules_per_cubic_meter
	return kilojoules_per_cubic_meter
}

// MegajoulePerCubicMeter returns the value in MegajoulePerCubicMeter.
func (a *EnergyDensity) MegajoulesPerCubicMeter() float64 {
	if a.megajoules_per_cubic_meterLazy != nil {
		return *a.megajoules_per_cubic_meterLazy
	}
	megajoules_per_cubic_meter := a.convertFromBase(EnergyDensityMegajoulePerCubicMeter)
	a.megajoules_per_cubic_meterLazy = &megajoules_per_cubic_meter
	return megajoules_per_cubic_meter
}

// GigajoulePerCubicMeter returns the value in GigajoulePerCubicMeter.
func (a *EnergyDensity) GigajoulesPerCubicMeter() float64 {
	if a.gigajoules_per_cubic_meterLazy != nil {
		return *a.gigajoules_per_cubic_meterLazy
	}
	gigajoules_per_cubic_meter := a.convertFromBase(EnergyDensityGigajoulePerCubicMeter)
	a.gigajoules_per_cubic_meterLazy = &gigajoules_per_cubic_meter
	return gigajoules_per_cubic_meter
}

// TerajoulePerCubicMeter returns the value in TerajoulePerCubicMeter.
func (a *EnergyDensity) TerajoulesPerCubicMeter() float64 {
	if a.terajoules_per_cubic_meterLazy != nil {
		return *a.terajoules_per_cubic_meterLazy
	}
	terajoules_per_cubic_meter := a.convertFromBase(EnergyDensityTerajoulePerCubicMeter)
	a.terajoules_per_cubic_meterLazy = &terajoules_per_cubic_meter
	return terajoules_per_cubic_meter
}

// PetajoulePerCubicMeter returns the value in PetajoulePerCubicMeter.
func (a *EnergyDensity) PetajoulesPerCubicMeter() float64 {
	if a.petajoules_per_cubic_meterLazy != nil {
		return *a.petajoules_per_cubic_meterLazy
	}
	petajoules_per_cubic_meter := a.convertFromBase(EnergyDensityPetajoulePerCubicMeter)
	a.petajoules_per_cubic_meterLazy = &petajoules_per_cubic_meter
	return petajoules_per_cubic_meter
}

// KilowattHourPerCubicMeter returns the value in KilowattHourPerCubicMeter.
func (a *EnergyDensity) KilowattHoursPerCubicMeter() float64 {
	if a.kilowatt_hours_per_cubic_meterLazy != nil {
		return *a.kilowatt_hours_per_cubic_meterLazy
	}
	kilowatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityKilowattHourPerCubicMeter)
	a.kilowatt_hours_per_cubic_meterLazy = &kilowatt_hours_per_cubic_meter
	return kilowatt_hours_per_cubic_meter
}

// MegawattHourPerCubicMeter returns the value in MegawattHourPerCubicMeter.
func (a *EnergyDensity) MegawattHoursPerCubicMeter() float64 {
	if a.megawatt_hours_per_cubic_meterLazy != nil {
		return *a.megawatt_hours_per_cubic_meterLazy
	}
	megawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityMegawattHourPerCubicMeter)
	a.megawatt_hours_per_cubic_meterLazy = &megawatt_hours_per_cubic_meter
	return megawatt_hours_per_cubic_meter
}

// GigawattHourPerCubicMeter returns the value in GigawattHourPerCubicMeter.
func (a *EnergyDensity) GigawattHoursPerCubicMeter() float64 {
	if a.gigawatt_hours_per_cubic_meterLazy != nil {
		return *a.gigawatt_hours_per_cubic_meterLazy
	}
	gigawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityGigawattHourPerCubicMeter)
	a.gigawatt_hours_per_cubic_meterLazy = &gigawatt_hours_per_cubic_meter
	return gigawatt_hours_per_cubic_meter
}

// TerawattHourPerCubicMeter returns the value in TerawattHourPerCubicMeter.
func (a *EnergyDensity) TerawattHoursPerCubicMeter() float64 {
	if a.terawatt_hours_per_cubic_meterLazy != nil {
		return *a.terawatt_hours_per_cubic_meterLazy
	}
	terawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityTerawattHourPerCubicMeter)
	a.terawatt_hours_per_cubic_meterLazy = &terawatt_hours_per_cubic_meter
	return terawatt_hours_per_cubic_meter
}

// PetawattHourPerCubicMeter returns the value in PetawattHourPerCubicMeter.
func (a *EnergyDensity) PetawattHoursPerCubicMeter() float64 {
	if a.petawatt_hours_per_cubic_meterLazy != nil {
		return *a.petawatt_hours_per_cubic_meterLazy
	}
	petawatt_hours_per_cubic_meter := a.convertFromBase(EnergyDensityPetawattHourPerCubicMeter)
	a.petawatt_hours_per_cubic_meterLazy = &petawatt_hours_per_cubic_meter
	return petawatt_hours_per_cubic_meter
}


// ToDto creates an EnergyDensityDto representation.
func (a *EnergyDensity) ToDto(holdInUnit *EnergyDensityUnits) EnergyDensityDto {
	if holdInUnit == nil {
		defaultUnit := EnergyDensityJoulePerCubicMeter // Default value
		holdInUnit = &defaultUnit
	}

	return EnergyDensityDto{
		Value: a.Convert(*holdInUnit),
		Unit:  *holdInUnit,
	}
}

// ToDtoJSON creates an EnergyDensityDto representation.
func (a *EnergyDensity) ToDtoJSON(holdInUnit *EnergyDensityUnits) ([]byte, error) {
	return a.ToDto(holdInUnit).ToJSON()
}

// Convert converts EnergyDensity to a specific unit value.
func (a *EnergyDensity) Convert(toUnit EnergyDensityUnits) float64 {
	switch toUnit { 
    case EnergyDensityJoulePerCubicMeter:
		return a.JoulesPerCubicMeter()
    case EnergyDensityWattHourPerCubicMeter:
		return a.WattHoursPerCubicMeter()
    case EnergyDensityKilojoulePerCubicMeter:
		return a.KilojoulesPerCubicMeter()
    case EnergyDensityMegajoulePerCubicMeter:
		return a.MegajoulesPerCubicMeter()
    case EnergyDensityGigajoulePerCubicMeter:
		return a.GigajoulesPerCubicMeter()
    case EnergyDensityTerajoulePerCubicMeter:
		return a.TerajoulesPerCubicMeter()
    case EnergyDensityPetajoulePerCubicMeter:
		return a.PetajoulesPerCubicMeter()
    case EnergyDensityKilowattHourPerCubicMeter:
		return a.KilowattHoursPerCubicMeter()
    case EnergyDensityMegawattHourPerCubicMeter:
		return a.MegawattHoursPerCubicMeter()
    case EnergyDensityGigawattHourPerCubicMeter:
		return a.GigawattHoursPerCubicMeter()
    case EnergyDensityTerawattHourPerCubicMeter:
		return a.TerawattHoursPerCubicMeter()
    case EnergyDensityPetawattHourPerCubicMeter:
		return a.PetawattHoursPerCubicMeter()
	default:
		return 0
	}
}

func (a *EnergyDensity) convertFromBase(toUnit EnergyDensityUnits) float64 {
    value := a.value
	switch toUnit { 
	case EnergyDensityJoulePerCubicMeter:
		return (value) 
	case EnergyDensityWattHourPerCubicMeter:
		return (value / 3.6e+3) 
	case EnergyDensityKilojoulePerCubicMeter:
		return ((value) / 1000.0) 
	case EnergyDensityMegajoulePerCubicMeter:
		return ((value) / 1000000.0) 
	case EnergyDensityGigajoulePerCubicMeter:
		return ((value) / 1000000000.0) 
	case EnergyDensityTerajoulePerCubicMeter:
		return ((value) / 1000000000000.0) 
	case EnergyDensityPetajoulePerCubicMeter:
		return ((value) / 1000000000000000.0) 
	case EnergyDensityKilowattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000.0) 
	case EnergyDensityMegawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000.0) 
	case EnergyDensityGigawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000.0) 
	case EnergyDensityTerawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000000.0) 
	case EnergyDensityPetawattHourPerCubicMeter:
		return ((value / 3.6e+3) / 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

func (a *EnergyDensity) convertToBase(value float64, fromUnit EnergyDensityUnits) float64 {
	switch fromUnit { 
	case EnergyDensityJoulePerCubicMeter:
		return (value) 
	case EnergyDensityWattHourPerCubicMeter:
		return (value * 3.6e+3) 
	case EnergyDensityKilojoulePerCubicMeter:
		return ((value) * 1000.0) 
	case EnergyDensityMegajoulePerCubicMeter:
		return ((value) * 1000000.0) 
	case EnergyDensityGigajoulePerCubicMeter:
		return ((value) * 1000000000.0) 
	case EnergyDensityTerajoulePerCubicMeter:
		return ((value) * 1000000000000.0) 
	case EnergyDensityPetajoulePerCubicMeter:
		return ((value) * 1000000000000000.0) 
	case EnergyDensityKilowattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000.0) 
	case EnergyDensityMegawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000.0) 
	case EnergyDensityGigawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000.0) 
	case EnergyDensityTerawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000000.0) 
	case EnergyDensityPetawattHourPerCubicMeter:
		return ((value * 3.6e+3) * 1000000000000000.0) 
	default:
		return math.NaN()
	}
}

// Implement the String() method for AngleDto
func (a EnergyDensity) String() string {
	return a.ToString(EnergyDensityJoulePerCubicMeter, 2)
}

// ToString formats the EnergyDensity to string.
// fractionalDigits -1 for not mention
func (a *EnergyDensity) ToString(unit EnergyDensityUnits, fractionalDigits int) string {
	value := a.Convert(unit)
	if fractionalDigits < 0 {
		formatted := strconv.FormatFloat(value, 'g', -1, 64)
		return fmt.Sprintf(formatted + " " + a.getUnitAbbreviation(unit))
	}
	return fmt.Sprintf("%.*f %s", fractionalDigits, value, a.getUnitAbbreviation(unit))
}

// GetUnitAbbreviation gets the unit abbreviation.
func (a *EnergyDensity) getUnitAbbreviation(unit EnergyDensityUnits) string {
	switch unit { 
	case EnergyDensityJoulePerCubicMeter:
		return "J/m³" 
	case EnergyDensityWattHourPerCubicMeter:
		return "Wh/m³" 
	case EnergyDensityKilojoulePerCubicMeter:
		return "kJ/m³" 
	case EnergyDensityMegajoulePerCubicMeter:
		return "MJ/m³" 
	case EnergyDensityGigajoulePerCubicMeter:
		return "GJ/m³" 
	case EnergyDensityTerajoulePerCubicMeter:
		return "TJ/m³" 
	case EnergyDensityPetajoulePerCubicMeter:
		return "PJ/m³" 
	case EnergyDensityKilowattHourPerCubicMeter:
		return "kWh/m³" 
	case EnergyDensityMegawattHourPerCubicMeter:
		return "MWh/m³" 
	case EnergyDensityGigawattHourPerCubicMeter:
		return "GWh/m³" 
	case EnergyDensityTerawattHourPerCubicMeter:
		return "TWh/m³" 
	case EnergyDensityPetawattHourPerCubicMeter:
		return "PWh/m³" 
	default:
		return ""
	}
}

// Check if the given EnergyDensity are equals to the current EnergyDensity
func (a *EnergyDensity) Equals(other *EnergyDensity) bool {
	return a.value == other.BaseValue()
}

// Check if the given EnergyDensity are equals to the current EnergyDensity
func (a *EnergyDensity) CompareTo(other *EnergyDensity) int {
	otherValue := other.BaseValue()
	if a.value < otherValue {
		return -1
	} else if a.value > otherValue {
		return 1
	}

	// If both are equal
	return 0
}

// Add the given EnergyDensity to the current EnergyDensity.
func (a *EnergyDensity) Add(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value + other.BaseValue()}
}

// Subtract the given EnergyDensity to the current EnergyDensity.
func (a *EnergyDensity) Subtract(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value - other.BaseValue()}
}

// Multiply the given EnergyDensity to the current EnergyDensity.
func (a *EnergyDensity) Multiply(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value * other.BaseValue()}
}

// Divide the given EnergyDensity to the current EnergyDensity.
func (a *EnergyDensity) Divide(other *EnergyDensity) *EnergyDensity {
	return &EnergyDensity{value: a.value / other.BaseValue()}
}